# kongctl SDK compatibility patches

This branch is based on public release `v0.76.0` (`ab1436fa`). It carries
the generated compatibility bundle from `Kong/kongctl@19b8cbbe9ac782da7c9e214ddf6166f83026f8c6`.
The source files and hashes are recorded in `kongctl-patches/provenance.json`.
Maintain patches in kongctl's `tools/sdk-patches`, then prepare a new snapshot.

The nine overlays omit optional request defaults for Portal, identity
providers, email templates, DCR configuration, audit webhooks, and backend
cluster TLS. Portal customization omits unset theme names and Try It UI
audiences. The post-generation replacement preserves explicit empty AI
Gateway `identity_realms` arrays. Upstream overlays and fixes remain enabled.

Regenerate with the bundle's pinned Speakeasy `1.802.0`:

```sh
python3 tools/sdk-patches/apply.py prepare /absolute/path/to/sdk-checkout
# From the prepared SDK checkout:
mise trust
mise install
git fetch origin 15001a933f0154e3177b2f31c07e9e79955e9ea4
CGO_ENABLED=0 GOFLAGS=-buildvcs=false mise exec -- make generate.sdk
go mod edit -go=1.26.10
CGO_ENABLED=0 GOFLAGS=-buildvcs=false go test . ./models/... ./internal/... ./pkg/...
GOFLAGS=-buildvcs=false make test.unit
```

The snapshot fetch supplies the release's pristine generation baseline for
linked worktrees. Restore the release's Go directive after generation because
its persistent edits otherwise select Go 1.26.8.

This branch is consumed by commit through kongctl's Go module replacement;
it is not a separate SDK release.
