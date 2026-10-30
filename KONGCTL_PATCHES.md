# kongctl compatibility patches for public SDK v0.79.0

Base: `Kong/sdk-konnect-go` v0.79.0
(`a9855b2dd9ddbb4d5a35aebbc63a157b79e03f62`). Generation uses the release's
public OpenAPI input directly, including the AI Gateway 2.3 API surface.
No internal SDK or replacement OpenAPI snapshot is used.

The canonical patch bundle comes from `Kong/kongctl` commit
`e7e524fe2103a9492d973d1e339356095f895eb2`. Source hashes are recorded in
`kongctl-patches/provenance.json`.

The patches preserve omitted optional request fields, explicit boolean and
format selections, and explicit empty identity realms. Replacements for
internal-only fields skip those fields when absent from this public SDK.
Regression tests cover serialization and omission behavior.

Speakeasy 1.805.0 matches the release's workflow lock. The upstream Makefile
and generation tools are retained, with the canonical overlay validation
and post-generation hooks installed by `apply.py prepare`.

```sh
sdk_snapshot_ref=refs/speakeasy/gen/115eea27-cda1-4a8f-b6ac-a19fea51c61c
git fetch origin "${sdk_snapshot_ref}:${sdk_snapshot_ref}"
mise trust
CGO_ENABLED=0 GOFLAGS=-buildvcs=false mise exec -- make generate.sdk
CGO_ENABLED=0 GOFLAGS=-buildvcs=false go test . ./models/... ./internal/... ./pkg/...
CGO_ENABLED=1 GOFLAGS=-buildvcs=false mise exec -- make test.unit
```

To refresh the bundle, commit it in kongctl first and run
`python3 tools/sdk-patches/apply.py prepare /absolute/path/to/this/checkout`
from kongctl before generation. Consume the SDK commit through an immutable
Go pseudo-version; ordinary kongctl builds do not apply patches.
