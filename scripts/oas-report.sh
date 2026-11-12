#!/bin/bash
#
# Purpose: print the input OAS report as one markdown block: the change list, then the
# versions table. Run from the repository root.
#
# Usage: oas-report.sh [old-ref]
#
#   oas-report.sh v0.80.0
#
#   old-ref  Optional. Git ref with the openapi.yaml to compare against.
#            An empty value means that no comparison happens.

set -o errexit
set -o nounset
set -o pipefail

# Changes table, then the marker @@TABLE@@, then the versions table.
OUTPUT=$(./scripts/oas-versions.sh "${1:-}")

CHANGES="${OUTPUT%%@@TABLE@@*}"
CHANGES="${CHANGES%$'\n'}"
TABLE="${OUTPUT#*@@TABLE@@}"

printf '## Changes\n\n%s\n\n## Input OAS versions\n\n%s\n' "${CHANGES}" "${TABLE}"
