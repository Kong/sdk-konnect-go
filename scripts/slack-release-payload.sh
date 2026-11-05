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

# Changes table, then the marker @@TABLE@@, then the versions table.
# shortcut: all markdown blocks together hold 12000 characters; both tables are about 4000 today.
OUTPUT=$(./scripts/oas-versions.sh "${PREV_TAG}")

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
