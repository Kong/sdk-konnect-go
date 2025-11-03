# Pre-release Version Implementation Plan

## TLDR

This SDK currently supports semantic versioning via PR labels (`patch`, `minor`, `major`). We want to add support for 
**release candidates (RC)** and other pre-release versions (alpha, beta).

**Key Challenge:** Go's `@latest` tag **intentionally skips pre-release versions**. 
Users must explicitly request them (e.g., `go get github.com/Kong/sdk-konnect-go@v0.13.0-rc.1`). 
This is correct behavior - we don't want unstable RCs auto-installed.

**Solution:** Add new PR labels (`rc`, `alpha`, `beta`) and a manual "graduate" workflow to promote RCs to stable releases. 
Since graduation happens after an RC is already merged (no more code changes), we'll use GitHub Actions' manual 
workflow dispatch feature instead of requiring dummy PRs.

**Key Behaviors:**
- Pre-release labels (`rc`, `alpha`, `beta`) alone on stable versions default to patch bump (e.g., `v0.12.0` + `rc` → `v0.12.1-rc.1`)
- Strict validation prevents invalid combinations like `patch` + `minor` or `rc` + `alpha`
- Manual workflow dispatch enables graduating RCs without creating dummy PRs

**Changes Required:**
1. Modify `.github/workflows/bump-version.yaml` - add pre-release label support and manual workflow dispatch
2. Update `.github/workflows/release.yaml` - mark GitHub releases as "pre-release" for RC/alpha/beta
3. Create/update `.goreleaser.yaml` - configure GoReleaser pre-release detection
4. Add new labels in GitHub repo settings

**Estimated Effort:** 4-6 hours implementation + testing

---

## Table of Contents

1. [Background](#background)
2. [Current State](#current-state)
3. [Goals](#goals)
4. [Go Module Pre-release Semantics](#go-module-pre-release-semantics)
5. [Proposed Solution](#proposed-solution)
6. [Implementation Details](#implementation-details)
7. [Workflow Examples](#workflow-examples)
8. [Testing Plan](#testing-plan)
9. [Documentation Updates](#documentation-updates)

---

## Background

This repository uses an automated build and release process triggered by PR labels:
- Upstream automation (from spec repository) creates PRs with OpenAPI spec changes
- Maintainers add semantic version labels (`patch`, `minor`, `major`) to PRs
- When merged, workflows automatically version, build, tag, and publish the SDK

The current workflow files:
- `.github/workflows/bump-version.yaml` - Reads PR labels and updates version in `.speakeasy/gen.yaml`
- `.github/workflows/generate_on_pr.yaml` - Regenerates SDK when `gen.yaml` changes
- `.github/workflows/release.yaml` - Creates git tags and publishes releases when merged to main

---

## Current State

### Version Management
- Versions stored in `.speakeasy/gen.lock` under `management.releaseVersion`
- Current version: `0.12.0`
- Speakeasy CLI manages version bumping via `speakeasy bump [patch|minor|major]`

### PR Label Workflow
1. PR created (usually by automated tool)
2. Maintainer adds exactly one label: `patch`, `minor`, or `major`
3. `bump-version.yaml` validates label exists (enforces exactly 1)
4. Workflow extracts label: `jq -r '.[]' | grep -E 'patch|minor|major'`
5. Runs: `speakeasy bump $BUMP_TYPE`
6. Commits version bump back to PR using `EndBug/add-and-commit` action
7. `generate_on_pr.yaml` detects change, regenerates SDK, commits again
8. PR merged to main
9. `release.yaml` detects `.speakeasy/gen.lock` change, creates tag and release

### Key Technical Details
- Uses `${{ secrets.PAT }}` token to bypass branch protection and commit to PR branches
- Checks out PR branch: `ref: ${{ github.event.pull_request.head.ref }}`
- Dependabot PRs are exempt from label requirements
- Recent tags: `v0.12.0`, `v0.11.1`, `v0.11.0`, `v0.10.2`

---

## Goals

### Primary Goals
1. Support release candidate (RC) builds tagged as `v1.0.0-rc.1`, `v1.0.0-rc.2`, etc.
2. Support alpha and beta pre-releases
3. Allow promotion of RC to stable release without code changes
4. Maintain compatibility with Go module toolchain expectations

### Non-Goals
- Changing `@latest` behavior (this is controlled by Go, not us)
- Auto-promotion of RCs to stable (manual approval required)
- Supporting non-SemVer version schemes

---

## Go Module Pre-release Semantics

### Critical Information

**Go's `@latest` query skips pre-release versions by design.** This is not a bug - it's intentional behavior to protect users from unstable releases.

### Pre-release Version Format
- Must follow SemVer: `v<major>.<minor>.<patch>-<prerelease>`
- Pre-release identifier follows hyphen: `v1.0.0-rc.1`
- Use dot separator for counters: `-rc.1` not `-rc1` (avoids lexical sort issues)
- Valid examples: `v1.0.0-rc.1`, `v2.3.0-alpha.2`, `v1.0.0-beta.1`

### Version Resolution Behavior

| Command | Stable Releases Exist? | Behavior |
|---------|----------------------|----------|
| `go get example.com/sdk@latest` | Yes | Fetches latest stable, **skips all pre-releases** |
| `go get example.com/sdk@latest` | No | Fetches latest pre-release (no stable available) |
| `go get example.com/sdk@v1.0.0-rc.1` | N/A | Fetches specific pre-release |
| `go get example.com/sdk@upgrade` | Yes | Same as `@latest` - skips pre-releases |

### Documentation Requirements
Users must be informed that:
1. Pre-release versions require explicit version tags
2. `@latest` will never fetch RC/alpha/beta releases (this is correct)
3. Example: `go get github.com/Kong/sdk-konnect-go@v0.13.0-rc.1`

### References
- [Go Module Version Numbers](https://go.dev/doc/modules/version-numbers)
- [Publishing Go Modules](https://go.dev/blog/publishing-go-modules)

---

## Proposed Solution

### Overview
Add pre-release support using a **combination of PR labels** (for creating RCs) and **manual workflow dispatch** (for graduating RCs to stable).

### New PR Labels
Add these labels to the repository:
- `rc` - Creates/increments release candidate
- `alpha` - Creates/increments alpha release
- `beta` - Creates/increments beta release

### Label Combination Policy
Allow **maximum 2 labels** on a PR:
- **Semantic version bump** (optional): `patch`, `minor`, or `major`
- **Pre-release type** (optional): `rc`, `alpha`, or `beta`

Examples:
- `minor` + `rc` → Creates `v0.13.0-rc.1` (from `v0.12.0`)
- `rc` only on stable version → Creates `v0.12.1-rc.1` (from `v0.12.0`, defaults to patch bump)
- `rc` only on RC version → Creates `v0.12.1-rc.2` (from `v0.12.1-rc.1`, increments RC counter)
- `major` only → Creates `v1.0.0` (standard stable release)

**Important:** When a pre-release label (`rc`, `alpha`, `beta`) is used alone on a stable version, the workflow defaults to a `patch` bump. This prevents version conflicts and ensures clear version progression.

### Manual Graduation Workflow
Add `workflow_dispatch` trigger to allow manual actions on the `main` branch:
- **graduate** - Promotes pre-release to stable (e.g., `v0.13.0-rc.2` → `v0.13.0`)
- **increment-rc** - Manually increments RC counter (fallback option)
- **set-version** - Manually sets specific version (emergency override)

### GitHub Release Marking
Configure workflows to mark pre-release versions as "Pre-release" in GitHub:
- Prevents RCs from showing as "Latest Release"
- Provides visual distinction in releases page
- Allows filtering by stable vs pre-release

---

## Implementation Details

### 1. Update bump-version.yaml

**File:** `.github/workflows/bump-version.yaml`

#### Changes Section 1: Add workflow_dispatch Trigger

**Location:** After line 9 (after `types: - unlabeled`)

```yaml
  workflow_dispatch:
    inputs:
      action:
        description: 'Action to perform'
        required: true
        type: choice
        options:
          - graduate
          - increment-rc
          - set-version
      version:
        description: 'Version (required for set-version action only)'
        required: false
        type: string
```

#### Changes Section 2: Update Label Validation

Replace the entire `check-semver-label` job with enhanced validation that prevents invalid label combinations.

**Location:** Lines 15-30 (entire `check-semver-label` job)

**Current:**
```yaml
  check-semver-label:
    runs-on: ubuntu-latest
    permissions:
      issues: write
      pull-requests: write
    steps:
      - name: Harden Runner
        uses: step-security/harden-runner@f4a75cfd619ee5ce8d5b864b0d183aff3c69b55a # v2.13.1
        with:
          egress-policy: audit
      - uses: mheap/github-action-required-labels@8afbe8ae6ab7647d0c9f0cfa7c2f939650d22509 # v5.5
        if: github.actor != 'dependabot[bot]'
        with:
          mode: exactly
          count: 1
          labels: "patch, minor, major"
```

**New:**
```yaml
  check-semver-label:
    runs-on: ubuntu-latest
    permissions:
      issues: write
      pull-requests: write
    steps:
      - name: Harden Runner
        uses: step-security/harden-runner@f4a75cfd619ee5ce8d5b864b0d183aff3c69b55a # v2.13.1
        with:
          egress-policy: audit

      # Basic check: ensure at least one version label exists
      - uses: mheap/github-action-required-labels@8afbe8ae6ab7647d0c9f0cfa7c2f939650d22509 # v5.5
        if: github.actor != 'dependabot[bot]'
        with:
          mode: minimum
          count: 1
          labels: "patch, minor, major, rc, alpha, beta"

      # Custom validation: prevent invalid label combinations
      - name: Validate label combination
        if: github.actor != 'dependabot[bot]'
        run: |
          LABELS='${{ toJson(github.event.pull_request.labels.*.name) }}'

          SEMVER_COUNT=$(echo "$LABELS" | jq -r '.[]' | grep -cE '^(patch|minor|major)$' || echo "0")
          PRERELEASE_COUNT=$(echo "$LABELS" | jq -r '.[]' | grep -cE '^(rc|alpha|beta)$' || echo "0")

          echo "### Label Validation" >> $GITHUB_STEP_SUMMARY
          echo "" >> $GITHUB_STEP_SUMMARY
          echo "- Semver labels found: $SEMVER_COUNT" >> $GITHUB_STEP_SUMMARY
          echo "- Pre-release labels found: $PRERELEASE_COUNT" >> $GITHUB_STEP_SUMMARY
          echo "" >> $GITHUB_STEP_SUMMARY

          # Check for multiple semver labels (e.g., patch + minor)
          if [ "$SEMVER_COUNT" -gt 1 ]; then
            echo "❌ **Error:** Multiple semver labels detected" >> $GITHUB_STEP_SUMMARY
            echo "" >> $GITHUB_STEP_SUMMARY
            echo "Please use only ONE of: \`patch\`, \`minor\`, or \`major\`" >> $GITHUB_STEP_SUMMARY
            exit 1
          fi

          # Check for multiple pre-release labels (e.g., rc + alpha)
          if [ "$PRERELEASE_COUNT" -gt 1 ]; then
            echo "❌ **Error:** Multiple pre-release labels detected" >> $GITHUB_STEP_SUMMARY
            echo "" >> $GITHUB_STEP_SUMMARY
            echo "Please use only ONE of: \`rc\`, \`alpha\`, or \`beta\`" >> $GITHUB_STEP_SUMMARY
            exit 1
          fi

          # Check for at least one label (redundant with mheap check, but explicit)
          if [ "$SEMVER_COUNT" -eq 0 ] && [ "$PRERELEASE_COUNT" -eq 0 ]; then
            echo "❌ **Error:** No version label found" >> $GITHUB_STEP_SUMMARY
            echo "" >> $GITHUB_STEP_SUMMARY
            echo "Please add at least one label:" >> $GITHUB_STEP_SUMMARY
            echo "- For stable release: \`patch\`, \`minor\`, or \`major\`" >> $GITHUB_STEP_SUMMARY
            echo "- For pre-release: \`rc\`, \`alpha\`, or \`beta\` (optionally with semver label)" >> $GITHUB_STEP_SUMMARY
            exit 1
          fi

          echo "✅ **Label combination is valid**" >> $GITHUB_STEP_SUMMARY
          echo "" >> $GITHUB_STEP_SUMMARY

          # Show what will happen
          if [ "$PRERELEASE_COUNT" -eq 1 ] && [ "$SEMVER_COUNT" -eq 0 ]; then
            echo "ℹ️  Pre-release label without semver label will default to patch bump if on stable version" >> $GITHUB_STEP_SUMMARY
          fi
```

**Why this approach:**
- Prevents invalid combinations like `patch` + `minor` or `rc` + `alpha`
- Provides clear error messages in GitHub UI
- Still allows valid combinations: `minor` + `rc`, just `patch`, just `rc`, etc.
- Shows helpful summary in workflow output

#### Changes Section 3: Update Conditional for PR Job

**Location:** Line 32 (add `if` condition to `bump-version` job)

**Add after `bump-version:` job declaration:**
```yaml
bump-version:
  if: github.event_name == 'pull_request'
  needs:
```

#### Changes Section 4: Replace Bump Logic

**Location:** Lines 64-67

**Current:**
```yaml
- name: Set bump (patch)
  run: |
    BUMP_TYPE=$(echo '${{ toJson(github.event.pull_request.labels.*.name) }}' | jq -r '.[]' | grep -E 'patch|minor|major')
    speakeasy bump $BUMP_TYPE
```

**New:**
```yaml
- name: Set bump (with pre-release support)
  run: |
    LABELS='${{ toJson(github.event.pull_request.labels.*.name) }}'

    # Extract label types
    SEMVER_BUMP=$(echo "$LABELS" | jq -r '.[]' | grep -E '^(patch|minor|major)$' || echo "")
    PRERELEASE_TYPE=$(echo "$LABELS" | jq -r '.[]' | grep -E '^(rc|alpha|beta)$' || echo "")

    echo "Semantic bump: '$SEMVER_BUMP'"
    echo "Pre-release type: '$PRERELEASE_TYPE'"

    if [ -n "$PRERELEASE_TYPE" ]; then
      # Pre-release workflow
      CURRENT_VERSION=$(yq '.management.releaseVersion' .speakeasy/gen.lock)
      echo "Current version: $CURRENT_VERSION"

      # DEFAULT TO PATCH: If no semver bump specified and current is stable, use patch
      if [ -z "$SEMVER_BUMP" ] && [[ ! "$CURRENT_VERSION" =~ - ]]; then
        echo "ℹ️  No semver label specified for stable version, defaulting to patch bump"
        SEMVER_BUMP="patch"
      fi

      # Apply semver bump if specified (or defaulted)
      if [ -n "$SEMVER_BUMP" ]; then
        echo "Applying $SEMVER_BUMP bump..."
        speakeasy bump $SEMVER_BUMP
        BASE_VERSION=$(yq '.management.releaseVersion' .speakeasy/gen.lock)
      else
        # Extract base version (remove pre-release suffix if exists)
        if [[ "$CURRENT_VERSION" =~ ^([0-9]+\.[0-9]+\.[0-9]+) ]]; then
          BASE_VERSION="${BASH_REMATCH[1]}"
        else
          BASE_VERSION="$CURRENT_VERSION"
        fi
      fi

      echo "Base version: $BASE_VERSION"

      # Determine pre-release counter
      if [[ "$CURRENT_VERSION" == *"-${PRERELEASE_TYPE}."* ]]; then
        # Increment existing counter for this pre-release type
        COUNTER=$(echo "$CURRENT_VERSION" | grep -oE "${PRERELEASE_TYPE}\\.([0-9]+)" | grep -oE "[0-9]+$")
        NEXT_COUNTER=$((COUNTER + 1))
        echo "Incrementing existing ${PRERELEASE_TYPE} counter: $COUNTER -> $NEXT_COUNTER"
      else
        # Start new pre-release series
        NEXT_COUNTER=1
        echo "Starting new ${PRERELEASE_TYPE} series at .1"
      fi

      # Set version with pre-release suffix
      NEW_VERSION="${BASE_VERSION}-${PRERELEASE_TYPE}.${NEXT_COUNTER}"
      echo "Setting version to: $NEW_VERSION"
      speakeasy bump -v "$NEW_VERSION"

    elif [ -n "$SEMVER_BUMP" ]; then
      # Standard semver bump (stable release)
      echo "Applying standard $SEMVER_BUMP bump"
      speakeasy bump $SEMVER_BUMP

    else
      echo "Error: No valid label found"
      exit 1
    fi

    # Verify the version was set
    FINAL_VERSION=$(yq '.management.releaseVersion' .speakeasy/gen.lock)
    echo "✅ Final version: $FINAL_VERSION"
```

**Key Logic Changes:**
- **Lines 335-339:** Automatic patch bump when pre-release label used on stable version
- **Line 336:** Checks if current version has no `-` (stable) and no semver label provided
- **Line 338:** Sets `SEMVER_BUMP="patch"` automatically
- This ensures `v0.12.0` + `rc` → `v0.12.1-rc.1` (not `v0.12.0-rc.1`)

#### Changes Section 5: Add Manual Workflow Job

**Location:** After line 75 (end of existing `bump-version` job)

```yaml
  manual-bump:
    if: github.event_name == 'workflow_dispatch'
    runs-on: ubuntu-latest
    permissions:
      contents: write
    steps:
      - name: Harden Runner
        uses: step-security/harden-runner@f4a75cfd619ee5ce8d5b864b0d183aff3c69b55a # v2.13.1
        with:
          egress-policy: audit

      - name: Install Speakeasy
        uses: mheap/setup-go-cli@fa9b01cdd4115eac636164f0de43bf7d51c82697 # v1.2.2
        with:
          owner: speakeasy-api
          repo: speakeasy
          cli_name: speakeasy
          package_type: zip
          version: 1.640.0 # renovate: datasource=github-releases depName=speakeasy-api/speakeasy

      - name: Checkout
        uses: actions/checkout@08c6903cd8c0fde910a37f88322edcfb5dd907a8 # v5.0.0
        with:
          ref: main
          token: ${{ secrets.PAT }}

      - name: Configure speakeasy CLI
        run: |
          mkdir -p ~/.speakeasy
          echo 'speakeasy_api_key: ${{ secrets.SPEAKEASY_API_KEY }}' > ~/.speakeasy/config.yaml

      - uses: jdx/mise-action@e3d7b8d67a7958d1207f6ed871e83b1ea780e7b0 # v3.3.1
        with:
          install: false

      - name: Validate current version
        run: |
          CURRENT=$(yq '.management.releaseVersion' .speakeasy/gen.lock)
          echo "Current version: $CURRENT"

          case "${{ github.event.inputs.action }}" in
            graduate)
              if [[ ! "$CURRENT" =~ -rc\.|^alpha\.|^beta\. ]]; then
                echo "Error: Current version '$CURRENT' is not a pre-release"
                echo "Graduate can only be used on RC, alpha, or beta versions"
                exit 1
              fi
              echo "✓ Version is a pre-release, proceeding with graduation"
              ;;
            increment-rc)
              if [[ ! "$CURRENT" =~ -rc\. ]]; then
                echo "Error: Current version '$CURRENT' is not an RC"
                exit 1
              fi
              echo "✓ Version is an RC, proceeding with increment"
              ;;
            set-version)
              if [ -z "${{ github.event.inputs.version }}" ]; then
                echo "Error: version input is required for set-version action"
                exit 1
              fi
              echo "✓ Version specified: ${{ github.event.inputs.version }}"
              ;;
          esac

      - name: Execute action
        run: |
          case "${{ github.event.inputs.action }}" in
            graduate)
              echo "Graduating pre-release to stable version..."
              speakeasy bump graduate
              ;;
            increment-rc)
              CURRENT=$(yq '.management.releaseVersion' .speakeasy/gen.lock)
              if [[ "$CURRENT" =~ (.*-rc\.)([0-9]+) ]]; then
                BASE="${BASH_REMATCH[1]}"
                NUM="${BASH_REMATCH[2]}"
                NEW_VERSION="${BASE}$((NUM + 1))"
                echo "Incrementing RC: $CURRENT -> $NEW_VERSION"
                speakeasy bump -v "$NEW_VERSION"
              else
                echo "Error: Failed to parse RC version from '$CURRENT'"
                exit 1
              fi
              ;;
            set-version)
              echo "Setting version to: ${{ github.event.inputs.version }}"
              speakeasy bump -v "${{ github.event.inputs.version }}"
              ;;
          esac

      - name: Generate SDK
        run: |
          export PATH=${PATH}:$(go env GOPATH)/bin
          make generate.sdk

      - name: Commit SDK changes to main
        uses: EndBug/add-and-commit@a94899bca583c204427a224a7af87c02f9b325d5 # v9.1.4
        with:
          add: .
          default_author: github_actions
          message: |
            ${{ github.event.inputs.action }}: automated version update

            🤖 Generated with [Claude Code](https://claude.com/claude-code)

            Co-Authored-By: Claude <noreply@anthropic.com>

      - name: Report new version
        run: |
          NEW_VERSION=$(yq '.management.releaseVersion' .speakeasy/gen.lock)
          echo "## ✅ Version Updated" >> $GITHUB_STEP_SUMMARY
          echo "" >> $GITHUB_STEP_SUMMARY
          echo "**Action:** ${{ github.event.inputs.action }}" >> $GITHUB_STEP_SUMMARY
          echo "**New Version:** v$NEW_VERSION" >> $GITHUB_STEP_SUMMARY
          echo "" >> $GITHUB_STEP_SUMMARY
          echo "The release will be created automatically by the publish workflow." >> $GITHUB_STEP_SUMMARY
```

---

### 2. Update release.yaml

**File:** `.github/workflows/release.yaml`

#### Change: Detect and Mark Pre-releases

**Location:** After line 48 (after GoReleaser step)

Add this step to detect if the release is a pre-release:

```yaml
      - name: Mark as pre-release if applicable
        run: |
          TAG=$(yq '.management.releaseVersion' .speakeasy/gen.lock)
          echo "Released version: v$TAG"

          if [[ "$TAG" =~ -rc\.|^alpha\.|^beta\. ]]; then
            echo "This is a pre-release version"
            echo "IS_PRERELEASE=true" >> $GITHUB_ENV
          else
            echo "This is a stable release"
            echo "IS_PRERELEASE=false" >> $GITHUB_ENV
          fi
```

**Note:** GoReleaser should automatically detect pre-release versions based on the SemVer format, but we'll explicitly configure it to be sure.

---

### 3. Create/Update .goreleaser.yaml

**File:** `.goreleaser.yaml`

If this file doesn't exist, create it. If it exists, update the `release` section.

```yaml
# GoReleaser configuration for Kong Konnect SDK
version: 2

before:
  hooks:
    - go mod tidy

builds:
  - env:
      - CGO_ENABLED=0
    goos:
      - linux
      - darwin
      - windows
    goarch:
      - amd64
      - arm64
    main: .
    binary: sdk-konnect-go
    ldflags:
      - -s -w
      - -X main.version={{.Version}}
      - -X main.commit={{.Commit}}
      - -X main.date={{.Date}}

archives:
  - format: tar.gz
    name_template: >-
      {{ .ProjectName }}_
      {{- title .Os }}_
      {{- if eq .Arch "amd64" }}x86_64
      {{- else if eq .Arch "386" }}i386
      {{- else }}{{ .Arch }}{{ end }}
    format_overrides:
      - goos: windows
        format: zip

checksum:
  name_template: 'checksums.txt'

snapshot:
  name_template: "{{ incpatch .Version }}-next"

changelog:
  sort: asc
  filters:
    exclude:
      - '^docs:'
      - '^test:'
      - '^chore:'

release:
  # Automatically detect pre-release versions based on SemVer
  prerelease: auto

  # Alternative explicit detection (choose one)
  # prerelease: '{{ if contains .Version "-" }}true{{ else }}false{{ end }}'

  # GitHub release configuration
  github:
    owner: Kong
    name: sdk-konnect-go

  # Release notes
  header: |
    ## Kong Konnect Go SDK {{ .Tag }}

    {{ if contains .Version "-rc." }}
    **This is a Release Candidate** - not recommended for production use.
    {{ else if contains .Version "-alpha." }}
    **This is an Alpha Release** - not recommended for production use.
    {{ else if contains .Version "-beta." }}
    **This is a Beta Release** - use with caution.
    {{ else }}
    **Stable Release**
    {{ end }}

    Install this version:
    ```bash
    go get github.com/Kong/sdk-konnect-go@{{ .Tag }}
    ```

  footer: |
    ---
    🤖 Generated with GoReleaser

    **Note:** Pre-release versions (RC, alpha, beta) are not fetched by `@latest`.
    Users must specify the exact version tag.

  # Don't create draft releases
  draft: false

  # Don't replace existing releases
  replace: false

  # GitHub Actions handles release creation
  skip_upload: false
```

**Note:** If the repository doesn't currently use GoReleaser with a config file, check the existing `release.yaml` workflow to see if it passes command-line flags. The `prerelease: auto` setting is the key configuration.

---

### 4. Add GitHub Repository Labels

Add these labels via GitHub UI or API:

| Label | Color | Description |
|-------|-------|-------------|
| `rc` | `#FFA500` (orange) | Creates or increments a release candidate version |
| `alpha` | `#FF6B6B` (red) | Creates or increments an alpha pre-release version |
| `beta` | `#4ECDC4` (teal) | Creates or increments a beta pre-release version |

**Via GitHub CLI:**
```bash
gh label create rc --color FFA500 --description "Creates or increments a release candidate version"
gh label create alpha --color FF6B6B --description "Creates or increments an alpha pre-release version"
gh label create beta --color 4ECDC4 --description "Creates or increments a beta pre-release version"
```

**Via GitHub UI:**
1. Go to repository → Issues → Labels
2. Click "New label"
3. Enter name, description, color
4. Click "Create label"

---

## Workflow Examples

### Example 1: Create First RC from PR

**Scenario:** New features ready for testing, increment minor version

**Steps:**
1. Upstream automation creates PR with OpenAPI spec updates
2. Maintainer reviews PR
3. Maintainer adds labels: `minor` + `rc`
4. `bump-version.yaml` triggers on `labeled` event
5. Workflow logic:
   - Detects `minor` + `rc` labels
   - Bumps from `0.12.0` → `0.13.0`
   - Adds `-rc.1` suffix → `0.13.0-rc.1`
   - Runs `speakeasy bump -v 0.13.0-rc.1`
6. Workflow commits version change back to PR
7. `generate_on_pr.yaml` detects `gen.yaml` change
8. Regenerates SDK, commits to PR
9. Maintainer merges PR to main
10. `release.yaml` triggers on `gen.lock` change
11. Creates tag `v0.13.0-rc.1`
12. Publishes GitHub Release marked as "Pre-release"

**Result:** Users can install with:
```bash
go get github.com/Kong/sdk-konnect-go@v0.13.0-rc.1
```

---

### Example 2: Create RC from Stable Version (No Semver Label)

**Scenario:** Want to test a small change as RC, current version is stable

**Steps:**
1. Upstream creates PR with changes
2. Current version: `0.12.0` (stable)
3. Maintainer adds label: `rc` only (no `patch`/`minor`/`major`)
4. Workflow logic:
   - Detects `rc` label only
   - Current version is stable (no `-` in version)
   - **Defaults to patch bump**
   - Bumps from `0.12.0` → `0.12.1`
   - Adds `-rc.1` suffix → `0.12.1-rc.1`
   - Runs `speakeasy bump -v 0.12.1-rc.1`
5. Commits and triggers SDK regeneration
6. Merge → Release created: `v0.12.1-rc.1`

**Result:**
```bash
go get github.com/Kong/sdk-konnect-go@v0.12.1-rc.1
```

**Note:** This default behavior prevents version conflicts (can't have both `v0.12.0` stable and `v0.12.0-rc.1`).

---

### Example 3: Increment Existing RC from PR

**Scenario:** Bug found in RC, need another RC iteration

**Steps:**
1. Upstream creates PR with bug fix
2. Current version: `0.12.1-rc.1` (pre-release)
3. Maintainer adds label: `rc` (no semver bump)
4. Workflow logic:
   - Detects `rc` label only
   - Current version is pre-release (has `-rc.`)
   - **Does NOT default to patch** (already a pre-release)
   - Extracts base: `0.12.1`
   - Increments counter: `.1` → `.2`
   - Runs `speakeasy bump -v 0.12.1-rc.2`
5. Commits and triggers SDK regeneration
6. Merge → Release created: `v0.12.1-rc.2`

**Result:**
```bash
go get github.com/Kong/sdk-konnect-go@v0.12.1-rc.2
```

---

### Example 4: Graduate RC to Stable (Manual)

**Scenario:** RC tested and approved, promote to stable release

**Steps:**
1. QA team confirms `v0.13.0-rc.2` is production-ready
2. Maintainer goes to GitHub Actions tab
3. Selects "Bump Version" workflow
4. Clicks "Run workflow" button
5. Selects branch: `main`
6. Chooses action: `graduate`
7. Clicks "Run workflow"
8. Workflow executes:
   - Checks out main branch
   - Validates current version is pre-release (`0.13.0-rc.2`)
   - Runs `speakeasy bump graduate`
   - Version becomes: `0.13.0`
   - Regenerates SDK
   - Commits to main with message: "graduate: automated version update"
   - Pushes to main
9. `release.yaml` triggers automatically
10. Creates tag `v0.13.0`
11. Publishes stable GitHub Release

**Result:**
```bash
# This now works (stable release)
go get github.com/Kong/sdk-konnect-go@latest

# Resolves to v0.13.0
```

---

### Example 5: Create Alpha/Beta Releases

**Scenario:** Early testing of experimental features

**Steps:**
1. PR created with experimental features
2. Add labels: `minor` + `alpha`
3. Workflow creates: `v0.13.0-alpha.1`
4. Later, add `beta` label to another PR
5. Workflow creates: `v0.13.0-beta.1`
6. Graduate to stable when ready

**Result:**
```bash
# Alpha testers
go get github.com/Kong/sdk-konnect-go@v0.13.0-alpha.1

# Beta testers
go get github.com/Kong/sdk-konnect-go@v0.13.0-beta.1
```

---

### Example 6: Emergency Version Override

**Scenario:** Manual intervention needed (version conflict, mistake)

**Steps:**
1. Go to Actions → "Bump Version" → "Run workflow"
2. Choose action: `set-version`
3. Enter version: `0.13.1`
4. Workflow sets exact version and publishes

**Use Cases:**
- Correct versioning mistakes
- Skip versions for compatibility
- Align with other Kong SDKs

---

## Testing Plan

### Phase 1: Branch Testing (Non-destructive)

Create a test branch to validate workflow changes:

```bash
# Create test branch
git checkout -b test/pre-release-workflow

# Make workflow changes
# ... implement changes from this document ...

# Commit changes
git add .github/workflows/
git commit -m "Add pre-release workflow support (testing)"

# Push to test branch
git push origin test/pre-release-workflow

# Create test PR
gh pr create --base test/pre-release-workflow --head test/pre-release-workflow \
  --title "Test PR for pre-release workflow" \
  --body "Testing RC label functionality"
```

#### Test Cases for PR Labels

1. **Test RC Label Alone**
   - Add `rc` label to test PR
   - Verify workflow runs
   - Check that version increments RC counter only
   - Verify commit is added to PR

2. **Test Minor + RC**
   - Add `minor` + `rc` labels
   - Verify version bumps minor and adds `-rc.1`
   - Check commit message

3. **Test Major + RC**
   - Add `major` + `rc` labels
   - Verify version bumps major and adds `-rc.1`

4. **Test Alpha/Beta**
   - Test `alpha` label
   - Test `beta` label
   - Verify `-alpha.1` and `-beta.1` suffixes

5. **Test Invalid Combinations**
   - Add 3+ labels (should fail validation)
   - Add `patch` + `minor` (should fail - can't have 2 semver bumps)

#### Test Cases for Manual Workflows

1. **Test Graduate**
   - Manually set version to `0.99.0-rc.1` on test branch
   - Run manual workflow with `graduate` action
   - Verify version becomes `0.99.0`
   - Check commit is created on branch

2. **Test Increment RC**
   - Set version to `0.99.0-rc.1`
   - Run manual workflow with `increment-rc`
   - Verify version becomes `0.99.0-rc.2`

3. **Test Set Version**
   - Run manual workflow with `set-version`
   - Input: `0.99.5`
   - Verify version is set exactly

4. **Test Validation Errors**
   - Try to graduate a stable version (should fail)
   - Try to increment-rc a non-RC version (should fail)
   - Try set-version without version input (should fail)

---

### Phase 2: Integration Testing (Main Branch)

After branch testing succeeds, test on main branch with real releases:

1. **Create Test RC**
   - Create a real PR with actual changes
   - Add `patch` + `rc` labels
   - Merge and verify release is created
   - Verify GitHub marks it as "Pre-release"
   - Test Go installation: `go get github.com/Kong/sdk-konnect-go@v<rc-version>`

2. **Test Graduate**
   - Use manual workflow to graduate RC
   - Verify stable release is created
   - Test Go installation: `go get github.com/Kong/sdk-konnect-go@latest`
   - Verify `@latest` returns graduated version

3. **Verify Go Module Proxy**
   - Check https://proxy.golang.org/github.com/Kong/sdk-konnect-go/@v/list
   - Verify RC versions appear in list
   - Verify `@latest` query returns stable version
   - Test in a separate Go project to ensure resolution works

---

### Phase 3: Documentation Testing

1. Update README with pre-release instructions
2. Have team members follow documentation to install RC
3. Collect feedback on clarity

---

## Documentation Updates

### Files to Update

#### 1. README.md

Add a "Pre-release Versions" section:

```markdown
## Installation

### Stable Release
```bash
go get github.com/Kong/sdk-konnect-go@latest
```

### Pre-release Versions

Release candidates, alpha, and beta versions are available for testing new features
before they're promoted to stable releases.

**Important:** Pre-release versions are not fetched by `@latest`. You must specify
the exact version tag:

```bash
# Install a specific release candidate
go get github.com/Kong/sdk-konnect-go@v0.13.0-rc.1

# Install a specific alpha version
go get github.com/Kong/sdk-konnect-go@v0.13.0-alpha.1

# Install a specific beta version
go get github.com/Kong/sdk-konnect-go@v0.13.0-beta.1
```

To see all available versions:
```bash
go list -m -versions github.com/Kong/sdk-konnect-go
```

Or visit: https://github.com/Kong/sdk-konnect-go/releases

### Why doesn't `@latest` fetch pre-release versions?

This is intentional Go behavior to protect users from installing unstable versions.
Pre-release versions (RC, alpha, beta) are meant for testing and must be explicitly
requested. Once a pre-release is promoted to stable, it becomes available via `@latest`.
```

#### 2. CONTRIBUTING.md (or similar)

Add a "Release Process" section:

```markdown
## Release Process

### Creating a Pre-release

1. **Create PR** with your changes (usually automated from spec repository)

2. **Add labels** to indicate the release type:
   - For a new RC: Add `minor` + `rc` (or `major`/`patch` + `rc`)
   - For incremental RC: Add `rc` only
   - For alpha/beta: Use `alpha` or `beta` label

3. **Merge PR** - The workflow will:
   - Bump the version
   - Regenerate the SDK
   - Create a GitHub Release marked as "Pre-release"
   - Create a git tag

4. **Test the RC**:
   ```bash
   go get github.com/Kong/sdk-konnect-go@v<rc-version>
   ```

### Promoting RC to Stable Release

When a release candidate has been tested and approved:

1. Go to **Actions** tab in GitHub
2. Select **"Bump Version"** workflow
3. Click **"Run workflow"**
4. Choose action: **graduate**
5. Click **"Run workflow"** button

The workflow will:
- Remove the pre-release suffix
- Regenerate the SDK
- Commit to main
- Trigger a stable release
- Make the version available via `@latest`

### Label Combinations

| Labels | Current Version | Result | Example |
|--------|----------------|--------|---------|
| `patch` | Any stable | Stable patch release | `0.12.0` → `0.12.1` |
| `minor` | Any stable | Stable minor release | `0.12.0` → `0.13.0` |
| `major` | Any stable | Stable major release | `0.12.0` → `1.0.0` |
| `minor` + `rc` | Any stable | New minor RC | `0.12.0` → `0.13.0-rc.1` |
| `rc` only | Stable (e.g., `0.12.0`) | **Default to patch** + RC | `0.12.0` → `0.12.1-rc.1` |
| `rc` only | Pre-release (e.g., `0.12.1-rc.1`) | Increment RC counter | `0.12.1-rc.1` → `0.12.1-rc.2` |
| `patch` + `alpha` | Any stable | New patch alpha | `0.12.0` → `0.12.1-alpha.1` |
| `alpha` only | Stable | **Default to patch** + alpha | `0.12.0` → `0.12.1-alpha.1` |
| `major` + `beta` | Any stable | New major beta | `0.12.0` → `1.0.0-beta.1` |
| `beta` only | Stable | **Default to patch** + beta | `0.12.0` → `0.12.1-beta.1` |

**Invalid Combinations (will fail CI):**
- `patch` + `minor` (multiple semver bumps)
- `minor` + `major` (multiple semver bumps)
- `rc` + `alpha` (multiple pre-release types)
- `alpha` + `beta` (multiple pre-release types)

### Manual Workflow Actions

Available actions in the manual workflow:

- **graduate** - Promote pre-release to stable
- **increment-rc** - Manually increment RC counter
- **set-version** - Set a specific version (emergency use)
```

#### 3. Create docs/releasing.md (New File)

Detailed release process documentation for maintainers:

```markdown
# Release Process Documentation

## Overview

This SDK uses automated workflows triggered by GitHub labels to manage versioning
and releases. This document explains the complete release process for maintainers.

## Version Types

### Stable Releases
- Format: `v<major>.<minor>.<patch>` (e.g., `v0.13.0`)
- Fetched by `go get @latest`
- Marked as "Latest Release" on GitHub
- Recommended for production use

### Pre-release Versions
- Format: `v<major>.<minor>.<patch>-<type>.<number>`
- Types: `rc` (release candidate), `alpha`, `beta`
- Examples: `v0.13.0-rc.1`, `v1.0.0-alpha.2`
- NOT fetched by `@latest` (Go intentional behavior)
- Marked as "Pre-release" on GitHub
- Requires explicit version tag to install

## Standard Release Flow

### 1. PR Creation
- Usually automated from upstream spec repository
- Contains OpenAPI specification updates
- Automatically regenerates SDK code

### 2. Version Labeling
Add ONE or TWO labels to the PR:

**Single Label (Stable Release):**
- `patch` - Bug fixes, no breaking changes
- `minor` - New features, backwards compatible
- `major` - Breaking changes

**Two Labels (Pre-release):**
- Semantic bump + pre-release type
- Examples: `minor` + `rc`, `major` + `alpha`

### 3. Automated Workflow
When labels are added:
1. `bump-version.yaml` validates labels
2. Updates version in `.speakeasy/gen.yaml`
3. Commits version bump to PR
4. `generate_on_pr.yaml` regenerates SDK
5. Commits SDK changes to PR

### 4. Review and Merge
- Review generated code changes
- Verify version is correct
- Merge PR to main

### 5. Publish
When merged to main:
1. `release.yaml` detects `.speakeasy/gen.lock` change
2. Creates git tag
3. Runs GoReleaser
4. Publishes GitHub Release

## Release Candidate Workflow

### When to Use RCs
- Major feature releases
- Potential breaking changes
- Need customer validation before stable release
- Want to test in staging environments

### Creating First RC

**Example:** New v0.13.0 features need testing

1. PR with features arrives (from automation)
2. Add labels: `minor` + `rc`
3. Merge PR
4. Result: `v0.13.0-rc.1` released

Share with testers:
```bash
go get github.com/Kong/sdk-konnect-go@v0.13.0-rc.1
```

### Creating Additional RCs

**Example:** Bug found in RC, need fix

1. PR with fix arrives
2. Add label: `rc` (no semver bump)
3. Merge PR
4. Result: `v0.13.0-rc.2` released

### Graduating RC to Stable

**Example:** RC approved for production

1. Verify RC is tested and approved
2. Go to Actions → "Bump Version"
3. Click "Run workflow"
4. Select: `graduate`
5. Click "Run workflow"

Result: `v0.13.0` stable release

Timeline:
- Workflow runs: ~2-3 minutes
- Commit appears on main
- Release published automatically
- Available via `@latest` immediately

### Rollback Procedure

If a graduated release has critical issues:

**Option 1: Quick Patch**
1. Create hotfix PR
2. Add label: `patch`
3. Merge → new stable version created

**Option 2: Manual Version Control**
1. Actions → "Bump Version" → "Run workflow"
2. Choose: `set-version`
3. Enter previous stable version
4. Creates new tag at old version number

## Alpha/Beta Releases

### When to Use

**Alpha:**
- Experimental features
- API design validation
- Early adopter feedback
- High likelihood of changes

**Beta:**
- Feature-complete but needs testing
- API mostly stable
- Performance validation
- Broader testing audience

### Creating Alpha/Beta

Same process as RC:
- Labels: `minor` + `alpha` or `major` + `beta`
- Merge → pre-release created
- Graduate when ready

## Manual Workflow Reference

Access: Actions → "Bump Version" → "Run workflow"

### Graduate
- **Purpose:** Promote pre-release to stable
- **Requirements:** Current version must be RC/alpha/beta
- **Example:** `0.13.0-rc.2` → `0.13.0`

### Increment RC
- **Purpose:** Manually bump RC counter
- **Requirements:** Current version must be RC
- **Example:** `0.13.0-rc.1` → `0.13.0-rc.2`
- **Note:** Usually not needed (use PR labels instead)

### Set Version
- **Purpose:** Emergency version override
- **Requirements:** Must provide version number
- **Example:** Set to `0.13.5` manually
- **Use Cases:**
  - Fix versioning mistakes
  - Align with external requirements
  - Skip versions for compatibility

## Troubleshooting

### Label Validation Failed - No Labels
**Error:** "Required labels not found"

**Cause:** No version labels on PR

**Fix:**
- Add at least one label: `patch`, `minor`, `major`, `rc`, `alpha`, or `beta`
- For stable release: use one semver label
- For pre-release: use pre-release label (optionally with semver label)

### Label Validation Failed - Multiple Semver Labels
**Error:** "Multiple semver labels detected"

**Cause:** PR has more than one of: `patch`, `minor`, `major`

**Example:** PR labeled with both `patch` and `minor`

**Fix:**
- Remove all but one semver label
- Choose the appropriate bump level
- Note: You can combine one semver + one pre-release label (e.g., `minor` + `rc`)

### Label Validation Failed - Multiple Pre-release Labels
**Error:** "Multiple pre-release labels detected"

**Cause:** PR has more than one of: `rc`, `alpha`, `beta`

**Example:** PR labeled with both `rc` and `alpha`

**Fix:**
- Remove all but one pre-release label
- Choose the appropriate pre-release type
- RC is typically used for near-stable releases
- Alpha for early testing, Beta for broader testing

### Version Conflict
**Error:** Version already exists

**Cause:** Tag already pushed to GitHub

**Fix:**
1. Check existing tags: `git tag -l`
2. Delete local: `git tag -d v<version>`
3. Delete remote: `git push origin :refs/tags/v<version>`
4. Re-run workflow

### Graduate Failed
**Error:** "Current version is not a pre-release"

**Cause:** Trying to graduate stable version

**Fix:**
- Verify current version has `-rc`, `-alpha`, or `-beta`
- Check `.speakeasy/gen.lock` → `management.releaseVersion`
- Use appropriate action instead

### Release Not Appearing
**Issue:** GitHub Release not created

**Check:**
1. Actions tab → "Publish" workflow
2. Verify `.speakeasy/gen.lock` was committed
3. Check GoReleaser logs for errors

### Go Can't Find Version
**Issue:** `go get @v<version>` fails

**Check:**
1. Tag exists: `git tag -l | grep <version>`
2. Tag pushed: `git ls-remote --tags origin`
3. Wait 5-10 minutes for Go proxy to sync
4. Check: https://proxy.golang.org/github.com/Kong/sdk-konnect-go/@v/list

## Best Practices

### For Maintainers
- Always create RC for major versions
- Test RCs in staging before graduation
- Document breaking changes in release notes
- Use semantic versioning strictly
- Don't skip versions without good reason

### For Version Numbers
- Follow SemVer 2.0.0 strictly
- Use `.` separator in pre-release: `-rc.1` not `-rc1`
- Increment RCs sequentially: `.1`, `.2`, `.3`
- Don't graduate until fully tested

### For Communication
- Announce RCs in team channels
- Share install instructions with testers
- Document known issues in RC releases
- Confirm approval before graduation

## FAQ

**Q: Can users accidentally install RCs?**
A: No. Go's `@latest` intentionally skips pre-releases.

**Q: What happens if I only add an `rc` label without a semver label?**
A: Depends on the current version:
- If current version is **stable** (e.g., `v0.12.0`): Defaults to **patch bump** → `v0.12.1-rc.1`
- If current version is **pre-release** (e.g., `v0.12.1-rc.1`): Increments RC counter → `v0.12.1-rc.2`

This prevents version conflicts and ensures clear progression.

**Q: Can I skip RC and go directly to stable?**
A: Yes. Use only `patch`/`minor`/`major` labels (no pre-release label).

**Q: How long does the Go proxy take to sync?**
A: Usually 1-2 minutes, up to 10 minutes maximum.

**Q: Can I delete a bad RC?**
A: Yes, but avoid it. Better to create a new RC with fixes.

**Q: What if automation breaks?**
A: Use manual workflow (`set-version`) to fix state, then resume normal process.

**Q: Can I have multiple RC series?**
A: Yes. Example: `0.13.0-rc.1` and `0.14.0-rc.1` can coexist.

**Q: What if I accidentally add both `patch` and `minor` labels?**
A: The workflow will fail with a clear error message. The validation step prevents invalid label combinations like multiple semver bumps or multiple pre-release types.

**Q: Why does `rc` only on stable default to patch instead of using the same version?**
A: Go modules don't allow having both `v0.12.0` (stable) and `v0.12.0-rc.1` published simultaneously. The patch bump ensures the RC version is distinct and progresses logically.
```

---

## Implementation Checklist

Use this checklist during implementation:

### Pre-Implementation
- [ ] Review this document with team
- [ ] Confirm approach with stakeholders
- [ ] Identify test branch/repository
- [ ] Ensure access to repository settings

### Code Changes
- [ ] Update `.github/workflows/bump-version.yaml`
  - [ ] Add `workflow_dispatch` trigger
  - [ ] Update label validation
  - [ ] Add conditional to `bump-version` job
  - [ ] Replace bump logic with pre-release support
  - [ ] Add `manual-bump` job
- [ ] Update `.github/workflows/release.yaml`
  - [ ] Add pre-release detection step
- [ ] Create/update `.goreleaser.yaml`
  - [ ] Add `prerelease: auto` configuration
  - [ ] Configure release notes templates
- [ ] Create `.github/workflows/TESTING.md` (testing notes)

### Repository Configuration
- [ ] Add `rc` label (orange, #FFA500)
- [ ] Add `alpha` label (red, #FF6B6B)
- [ ] Add `beta` label (teal, #4ECDC4)
- [ ] Verify `PAT` secret exists and has correct permissions
- [ ] Verify `SPEAKEASY_API_KEY` secret exists

### Testing
- [ ] Create test branch
- [ ] Test RC label on test PR
- [ ] Test minor + rc labels
- [ ] Test major + rc labels
- [ ] Test alpha/beta labels
- [ ] Test invalid label combinations
- [ ] Test manual graduate workflow
- [ ] Test manual increment-rc workflow
- [ ] Test manual set-version workflow
- [ ] Test validation error cases

### Integration Testing
- [ ] Create real RC on main branch
- [ ] Verify GitHub marks as pre-release
- [ ] Test Go installation of RC
- [ ] Graduate RC to stable
- [ ] Verify stable release created
- [ ] Test `@latest` fetches stable version
- [ ] Verify Go proxy shows correct versions

### Documentation
- [ ] Update README.md with pre-release instructions
- [ ] Update CONTRIBUTING.md with release process
- [ ] Create docs/releasing.md (maintainer guide)
- [ ] Add comments to workflow files
- [ ] Update any deployment/CI documentation

### Communication
- [ ] Announce new capability to team
- [ ] Share documentation links
- [ ] Conduct walkthrough/demo if needed
- [ ] Update team onboarding materials

### Post-Implementation
- [ ] Monitor first few RCs for issues
- [ ] Collect feedback from maintainers
- [ ] Iterate on documentation based on feedback
- [ ] Update this document with lessons learned

---

## Timeline Estimate

| Phase | Duration | Notes |
|-------|----------|-------|
| Pre-implementation review | 1 hour | Team discussion |
| Code changes | 2-3 hours | Workflow modifications |
| Repository setup | 30 minutes | Labels, secrets verification |
| Branch testing | 1-2 hours | Test all scenarios |
| Documentation | 1-2 hours | README, guides |
| Integration testing | 1 hour | Real releases on main |
| Communication | 30 minutes | Team announcement |
| **Total** | **6-10 hours** | Spread over 2-3 days |

---

## Rollback Plan

If the implementation causes issues:

### Immediate Rollback
1. Revert workflow files to previous versions
2. Remove new labels from repository
3. Continue using stable release process only

### Partial Rollback
- Keep workflow changes but don't use pre-release labels
- System remains backwards compatible
- Re-enable when ready

### Recovery Steps
1. Identify issue (check workflow logs)
2. Fix in separate branch
3. Test thoroughly before re-deployment
4. Document issue and resolution

---

## Support and Maintenance

### Workflow Ownership
- **Primary:** DevOps/Platform team
- **Secondary:** SDK maintainers
- **Escalation:** Kong API team

### Monitoring
- Watch GitHub Actions for workflow failures
- Monitor Go proxy for version availability
- Track team feedback on process

### Iteration
- Quarterly review of process
- Update documentation based on usage
- Refine workflows based on pain points

---

## References

### Go Documentation
- [Module version numbering](https://go.dev/doc/modules/version-numbers)
- [Publishing Go Modules](https://go.dev/blog/publishing-go-modules)
- [Module release workflow](https://go.dev/doc/modules/release-workflow)

### Tools
- [Speakeasy CLI Documentation](https://speakeasy.com/docs)
- [GoReleaser Documentation](https://goreleaser.com/intro/)
- [GitHub Actions Documentation](https://docs.github.com/en/actions)

### Internal
- [Current Repository](https://github.com/Kong/sdk-konnect-go)
- [Upstream Spec Repository](https://github.com/Kong/konnect-openapi)
- [Kong SDK Standards](https://konghq.atlassian.net/wiki/spaces/TECH/pages/sdk-standards)

---

## Version History

| Version | Date | Author | Changes |
|---------|------|--------|---------|
| 1.0 | 2025-01-XX | Kong Team | Initial implementation plan |

---

## Appendix: Go Module Versioning Quick Reference

### Version Format
```
v<major>.<minor>.<patch>[-<prerelease>][+<metadata>]

Examples:
v1.0.0              - Stable release
v1.0.0-rc.1         - Release candidate 1
v1.0.0-alpha.2      - Alpha version 2
v1.0.0-beta.3       - Beta version 3
v2.0.0-rc.1+build.1 - RC with build metadata
```

### Resolution Rules
```
@latest             → Latest stable (skips pre-releases)
@v1.0.0-rc.1        → Specific pre-release
@v1                 → Latest v1.x.x stable
@upgrade            → Same as @latest (skips pre-releases)
```

### Pre-release Ordering
```
v1.0.0-alpha.1
v1.0.0-alpha.2
v1.0.0-beta.1
v1.0.0-beta.2
v1.0.0-rc.1
v1.0.0-rc.2
v1.0.0              ← Stable
```

### Common Patterns
```
Feature development:  v1.0.0 → v1.1.0-alpha.1 → v1.1.0-beta.1 → v1.1.0-rc.1 → v1.1.0
Quick iteration:      v1.0.0 → v1.0.1-rc.1 → v1.0.1
Major release:        v1.9.0 → v2.0.0-rc.1 → v2.0.0-rc.2 → v2.0.0
```

---

**Document Status:** Ready for Implementation
**Last Updated:** 2025-01-XX
**Next Review:** After first production RC graduation
