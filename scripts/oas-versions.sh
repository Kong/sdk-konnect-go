#!/bin/bash
#
# Purpose: compare the input OAS versions of the working tree with a git ref.
# Prints the changes table, the marker @@TABLE@@, then the versions table.
# The caller splits the output on @@TABLE@@. Run from the repository root.
#
# Usage: oas-versions.sh [old-ref]
#
#   oas-versions.sh v0.80.0
#
#   old-ref  Optional. Git ref with the openapi.yaml to compare against.
#            An empty value means that no comparison happens.

set -o errexit
set -o nounset
set -o pipefail

OLD_REF="${1:-}"

# Input OAS title, version and path at old-ref. Empty when there is no ref or no openapi.yaml there.
PREV=""
if [[ -n "${OLD_REF}" ]]; then
  if ! PREV=$(git show "${OLD_REF}:openapi.yaml" 2>/dev/null | yq -r '.x-input-schemas[] | [.title, .version, .path] | @tsv'); then
    echo "No openapi.yaml with x-input-schemas at ${OLD_REF}. There is no change list." >&2
    PREV=""
  fi
fi
export PREV

# Entries match by path, which is stable when a title changes.
# ENVIRON keeps the newlines out of awk escaping.
yq -r '.x-input-schemas[] | [.title, .version, .path] | @tsv' openapi.yaml | awk -F'\t' '
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
  }'
