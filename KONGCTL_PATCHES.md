# kongctl SDK patches

Baseline: public SDK `v0.73.1`
(`a3bc4bed1e4a1cf90cbb2171fd2089755c332020`). The baseline OpenAPI input is
unchanged; the overlays preserve the existing AI Gateway 2.2 surfaces.

## Patch provenance

The original public patches span `16ace680`, `7e77ff86`, and `4d68939e`
on `patch/v0.68.1/targeted-patch-defaults`. They were subsequently ported
to internal SDK v0.5.0 at `ae595d8c3fcea65864b67cfc76abe8970b2751fe`,
then public SDK v0.71.0 at `288b6e157cf052ca2c6a0a70a90d0ca2aaa0e624`.
This branch carries those seven overlays onto public SDK v0.73.1.

Seven overlays preserve omission of optional request fields:

- Backend-cluster TLS `insecure_skip_verify`.
- Custom portal email-template PATCH `enabled`.
- Portal team create/update `konnect_managed`.
- HTTP DCR PATCH `allow_multiple_credentials`.
- Portal audit-webhook PATCH `enabled`.
- Organization and portal identity-provider PATCH `enabled`.
- Portal PATCH boolean fields.

The internal SDK's organization-team `konnect_managed` overlay and
audit-webhook `include_principal_name` action do not apply: these fields
are absent from the public request schemas. Portal-team membership
management remains supported and patched. The organization-team test
still checks sparse updates and preservation of sibling kongctl metadata.

The `RouteJSON.strip_path` and `GroupMembership.members` post-generation
fixes are already upstream. No duplicate fixes are added. The original
handwritten JSON wrappers are unnecessary with Speakeasy 1.801.0 and the
overlays; regression tests check request omission and explicit values.

## Reproduction

```sh
mise trust
mise install
CGO_ENABLED=0 GOFLAGS=-buildvcs=false mise exec -- make generate.sdk
go test . ./models/... ./internal/... ./pkg/...
make test.unit
```

In a linked Git worktree, prefetch the baseline generation snapshot if
Speakeasy cannot locate its remote:

```sh
git fetch origin refs/speakeasy/gen/9febfc42-561b-4985-8db5-343e34562184
```

The snapshot preserves upstream persistent edits, including import formatting.

The workflow and mise pins match the baseline generation lockfile
(1.801.0). Baseline generator settings are preserved. All overlays must apply strictly before
generation. The normal workflow generates code, DeepCopy methods, applies
upstream fixes, tidies modules, and lints the specification.

AI Gateway tests cover passthrough, skills, multiple aliases, Typesafe
provider/target configuration, decisions, and custom-policy CRUD. Mock
HTTP tests do not establish production service or data-plane availability.
The public module requires no private SDK authentication or replacement
repository credentials.

Live SDK integration tests require `KONNECT_API_PAT`, `KONNECT_API_URL`,
and a test organization. They are separate from the offline checks.
