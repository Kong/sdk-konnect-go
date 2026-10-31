# Selective Kong Identity public SDK preview

This branch keeps the public SDK baseline and adds the selected Identity
specifications. It is a development dependency, not a published SDK release.
The source revisions and preparation hashes are recorded in `provenance.json`.
KongCTL patches are snapshotted separately in `kongctl-patches/provenance.json`.

After building the selected specification with the sibling platform-api
checkout, run:

```sh
node identity-preview/prepare.cjs
node --test identity-preview/credential-expiration.test.cjs
CGO_ENABLED=0 GOFLAGS=-buildvcs=false mise exec -- make generate.sdk
CGO_ENABLED=0 GOFLAGS=-buildvcs=false go test . ./models/... ./internal/... ./pkg/...
GOFLAGS=-buildvcs=false make test.unit
```

Preparation disambiguates colliding schema names and limits selection to the
intended Identity surface. Credential expiration is also separated from an
unrelated read-only IAM component with the same name. The correction changes
only six credential request references and adds one writable nullable schema;
the tests check that boundary and repeated preparation.

Optional credential expiration uses the SDK's existing
`optionalnullable.OptionalNullable[time.Time]` wrapper: zero omits the field,
`optionalnullable.From(&expiration)` sends a date, and
`optionalnullable.From[time.Time](nil)` sends null. Password replacement requires
the field and uses `*time.Time`, where nil sends null. The patch snapshot
contains the generated model/example corrections and actual HTTP wire tests.
Global nullable-wrapper generation remains disabled, preserving unrelated
public SDK models.

In a linked worktree, Speakeasy pristine-generation objects may be stored in
its private Git directory. Ensure the pristine commit recorded in
`.speakeasy/gen.lock` is available to normal Git before regenerating; do not
disable persistent edits to bypass a missing snapshot. Existing public SDK
customizations must survive generation.
