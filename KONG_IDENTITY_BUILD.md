# Kong Identity public SDK preview

This branch extends KongCTL's current patched public SDK with the Kong
Identity APIs discussed in the triage session. It is consumed by commit or
local module replacement. It has no release, tag, or pull request.

## Inputs and scope

- Public SDK base: `d76cbad8eb3b07577ac30fda2454be0b33ee74ba`.
- Platform API preview: branch `build/ki-public-sdk`, commit
  `57350ae722c69faa20a740b72549cef6cbc70fee`.
- Identity v2 spec: PR #3877 at
  `98137d7be3772b67693cc1ceb0469e77f8ecf0f5`.
- KongCTL patch bundle: `2489716c3eeb8e437b13336cd682fc8d5580b0db`.
- Speakeasy: `1.805.0`.

The public SDK's `.speakeasy/gen.yaml` is unchanged, including persistent
edits, error handling, method argument limits, and the module path. Other
specs retain their public visibility filtering. This avoids the unrelated
public/internal SDK differences found in the earlier experiment.

`identity-preview/prepare.cjs` gives new auth-server shared types, credential
types, force parameters and list filters distinct Go names. It excludes
five Connected Apps operations included by Sliver. These changes affect only
new SDK APIs. `identity-preview/provenance.json` records the inputs and
`identity-preview/coverage.json` lists the selected additions.

Preview preparation also removes code-sample publishing from the workflow.
The standard workflow uploaded code samples during earlier attempts. The
preview does not publish an SDK or spec release.

The preview adds 112 operations: auth servers, clients, claims, scopes,
secrets and token inspection/revocation; directories; principals, identities
and credentials; groups, memberships and keys; Token Vault providers,
templates, shared credentials and public keys; principal vault credential
inspection/revocation; directory trusted IdPs; org-chart import.

Internal computed endpoints, realm configuration, org-root reads,
directory-wide key listing, quotas, and internal Sliver features remain
excluded. API deployment and declarative ownership/matching behavior require
separate validation.

## Regeneration

Use sibling worktrees named `platform-api` and `sdk-konnect-go`. Check out
the platform preview commit above, install its npm dependencies, and run:

```sh
bash tools/gen_go_sdk.sh
```

This writes this SDK's `openapi.yaml` through the normal public pipeline.
From this SDK, prepare the selected APIs and their Go names:

```sh
node identity-preview/prepare.cjs
```

The KongCTL patch snapshot is already installed. To refresh it, run this from
the desired KongCTL revision:

```sh
python3 tools/sdk-patches/apply.py prepare /absolute/path/to/sdk-konnect-go
```

From this SDK, regenerate with the pinned Speakeasy version:

```sh
mise trust
mise install
CGO_ENABLED=0 GOFLAGS=-buildvcs=false mise exec -- \
  make remove.interfaces remove.mocks remove.field_tests
CGO_ENABLED=0 GOFLAGS=-buildvcs=false mise exec -- make generate.sdk
```

Start generation from a fresh checkout of the public SDK base and copy the
`identity-preview/` directory from this branch. Reusing an intermediate
generation in a linked worktree failed when Speakeasy could not load its
persistent-edit snapshot. The successful build preserves the public SDK's
persistent-edit setting; it does not disable that feature to bypass the issue.

The local mise setup cannot install ifacemaker through the SDK's isolated
tool directory because that directory lacks its configured Go dependency.
Installing the same pinned ifacemaker directly avoids changing SDK settings:

```sh
CGO_ENABLED=0 GOTOOLCHAIN=go1.26.10 GOBIN="$PWD/bin" \
  go install github.com/vburenin/ifacemaker@v1.4.0
CGO_ENABLED=0 GOFLAGS=-buildvcs=false mise exec -- make \
  -o ifacemaker IFACEMAKER="$PWD/bin/ifacemaker" \
  generate.interfaces generate.mocks generate.field_tests
```

Restore the baseline module dependency pins before tidying if regeneration
temporarily removed mock-only dependencies. Then run the SDK checks:

```sh
go mod tidy
CGO_ENABLED=0 GOTOOLCHAIN=go1.26.10 \
  go test . ./models/... ./internal/... ./pkg/...
CGO_ENABLED=0 GOTOOLCHAIN=go1.26.10 go test -run '^$' ./...
```

Generation uses `--skip-versioning --skip-upload-spec`. Live API integration
tests are outside this experiment.

## KongCTL compatibility

Claim PATCH now preserves explicit empty `include_in_scopes` arrays. The
targeted `omitzero` patch keeps nil arrays omitted; its HTTP wire regression
fails on the original serializer and passes after patching. The public preview
keeps all previous KongCTL overlays and post-generation fixes.

Validation passed:

- Platform API full build and public SDK generation, including staticcheck
  and OpenAPI linting (zero errors; existing warnings remain).
- All 990 existing operations retain their request/response shapes after
  resolving references and normalizing description whitespace. Global
  servers, security requirements and security schemes are unchanged.
- An AST audit found no incompatible changes in 37,119 existing exported
  declarations: type members/tags, method signatures and constants. It
  allows additional fields and declarations and ignores comments and order.
- SDK unit/serialization tests, generated field guards, and compilation of
  all packages pass.
- Unchanged KongCTL builds and all unit tests pass with `CGO_ENABLED=0`.

The SDK module dependencies and public generator configuration are unchanged.
The API audit result is saved in `identity-preview/go-api-audit.json`.

The compatibility run uses KongCTL revision
`2beef7beeab6b3abd1e0d39978a1c48c203d83dd`. A copied module file points at
this SDK worktree; KongCTL source, `go.mod`, `go.sum` and vendoring are left
unchanged. Use `-mod=mod -modfile=/path/to/copied.mod` to bypass the existing
vendor snapshot during this check.

Build logs and the temporary module files are saved locally under
`/home/rspurgeon/.cache/kong-identity-public-sdk/`.

## Auth server PATCH array preservation

The preview additionally includes the KongCTL patch bundle from
`b734e481468c76b58b97f7624a53d08793bde8de` (including the origins patch
from `af17290382c5b1f93b76a713085632984609ad65`). The targeted
`UpdateAuthServer.trusted_origins` tag uses `omitzero`, preserving an explicitly
empty array while leaving nil omitted. This enables PATCH clearing without a
custom HTTP client. The wire regression fails before the patch and passes
after it; root, model, internal, and package tests pass. API surface and
visibility selection remain unchanged.
