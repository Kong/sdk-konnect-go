#!/bin/bash
#
# Purpose: build the Slack payload for a release announcement.
# Prints the JSON payload to stdout. Run from the repository root.
#
# Usage: slack-release-payload.sh [previous-tag] [version]
#
#   slack-release-payload.sh v0.80.0
#
# Arguments, in the order that git takes them: old first, then new.
#   previous-tag  Optional. Git tag of the previous release. The input OAS versions
#                 in the message compare against the input OAS versions at this tag.
#                 An empty value means that no comparison happens.
#   version       Optional. Release version without the leading v. The default is the
#                 version in .speakeasy/gen.lock.
#
# Environment:
#   GITHUB_SERVER_URL    Set by GitHub Actions.
#   GITHUB_REPOSITORY    Set by GitHub Actions.
#   GITHUB_RUN_ID        Set by GitHub Actions.
#   GITHUB_RUN_NUMBER    Set by GitHub Actions.

set -o errexit
set -o nounset
set -o pipefail

PREV_TAG="${1:-}"
VERSION="${2:-$(yq '.management.releaseVersion' .speakeasy/gen.lock)}"

# Input OAS title, version and path for the previous release. Empty when there is no previous tag.
PREV=""
if [[ -n "${PREV_TAG}" ]]; then
  if ! PREV=$(git show "${PREV_TAG}:openapi.yaml" 2>/dev/null | yq -r '.x-input-schemas[] | [.title, .version, .path] | @tsv'); then
    echo "No openapi.yaml at ${PREV_TAG}. The message has no change list." >&2
    PREV=""
  fi
fi
export PREV

# Entries match by path, which is stable when a title changes.
# The awk output holds the changes table, then a marker, then the versions table.
# ENVIRON keeps the newlines out of awk escaping.
# shortcut: all markdown blocks together hold 12000 characters; both tables are about 4000 today.
OUTPUT=$(yq -r '.x-input-schemas[] | [.title, .version, .path] | @tsv' openapi.yaml | awk -F'\t' '
  BEGIN {
    prev = ENVIRON["PREV"]
    n = split(prev, rows, "\n")
    for (i = 1; i <= n; i++) {
      split(rows[i], f, "\t")
      oldtitle[f[3]] = f[1]; oldver[f[3]] = f[2]
    }
  }
  {
    seen[$3] = 1
    versions = versions sprintf("| %s | %s | %s |\n", $1, $2, $3)
    if (prev != "") {
      if (!($3 in oldver)) changes = changes sprintf("| %s | (new) | %s | %s |\n", $1, $2, $3)
      else if (oldver[$3] != $2) changes = changes sprintf("| %s | ~~%s~~ | %s | %s |\n", $1, oldver[$3], $2, $3)
    }
  }
  END {
    for (p in oldver) if (!(p in seen)) changes = changes sprintf("| %s | ~~%s~~ | (removed) | %s |\n", oldtitle[p], oldver[p], p)
    if (changes == "") changes = "No input OAS changes\n"
    else changes = "| Input OAS | Old version | Version | Path |\n| --- | --- | --- | --- |\n" changes
    printf "%s@@TABLE@@| Input OAS | Version | Path |\n| --- | --- | --- |\n%s", changes, versions
  }')

CHANGES="${OUTPUT%%@@TABLE@@*}"
CHANGES="${CHANGES%$'\n'}"
TABLE="${OUTPUT#*@@TABLE@@}"

if [[ -n "${PREV_TAG}" ]]; then
  CHANGES_HEADER="Changes since ${PREV_TAG}"
else
  CHANGES_HEADER="No previous release to compare with"
  CHANGES=""
fi

RELEASE_URL="${GITHUB_SERVER_URL}/${GITHUB_REPOSITORY}/releases/tag/v${VERSION}"
RUN_URL="${GITHUB_SERVER_URL}/${GITHUB_REPOSITORY}/actions/runs/${GITHUB_RUN_ID}"

jq -n \
  --arg channel "project-konnect-sdks" \
  --arg version "v${VERSION}" \
  --arg release_url "${RELEASE_URL}" \
  --arg run_number "${GITHUB_RUN_NUMBER}" \
  --arg run_url "${RUN_URL}" \
  --arg changes_header "${CHANGES_HEADER}" \
  --arg changes "${CHANGES}" \
  --arg table "${TABLE}" \
  '{
    channel: $channel,
    text: "🎉 sdk-konnect-go \($version) released: \($release_url)",
    blocks: [
      { type: "section", text: { type: "mrkdwn", text: "🎉 *sdk-konnect-go <\($release_url)|\($version)>* released 🎉" } },
      { type: "section", text: { type: "mrkdwn", text: "*Workflow run:* <\($run_url)|Run #\($run_number)>" } },
      { type: "markdown", text: "**\($changes_header)**\n\n\($changes)" },
      { type: "markdown", text: "**Input OAS versions**\n\n\($table)" }
    ]
  }'
