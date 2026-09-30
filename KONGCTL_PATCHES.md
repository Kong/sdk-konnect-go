# kongctl SDK patches

Baseline: public SDK `v0.71.0`
(`e1323b4df55923c0ec7fc922646a37dafebb46bd`). Its unchanged OpenAPI input
includes AI Gateway specification `2.2.0`.

## Patch provenance

The original public patches span `16ace680`, `7e77ff86`, and `4d68939e`
on `patch/v0.68.1/targeted-patch-defaults`. They were subsequently ported
to internal SDK v0.5.0 at `ae595d8c3fcea65864b67cfc76abe8970b2751fe`.

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
handwritten JSON wrappers are unnecessary with Speakeasy 1.799.0 and the
overlays; regression tests check request omission and explicit values.

## Reproduction

```sh
mise trust
mise install
mise exec -- make generate.sdk
go test . ./models/... ./internal/... ./pkg/...
make test.unit
```

The mise pin matches the baseline generation lockfile (1.799.0). Baseline
generator settings are preserved. All overlays must apply strictly before
generation. The normal workflow generates code, DeepCopy methods, applies
upstream fixes, tidies modules, and lints the specification.

AI Gateway tests cover passthrough, skills, multiple aliases, Typesafe
provider/target configuration, decisions, and custom-policy CRUD. Mock
HTTP tests do not establish production service or data-plane availability.
The public module requires no private SDK authentication or replacement
repository credentials.

Live SDK integration tests require `KONNECT_API_PAT`, `KONNECT_API_URL`,
and a test organization. They are separate from the offline checks.
