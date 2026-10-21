# Configuration through environment variables

This SDK can be configured via the environment variables.

## Konnect domain

Konnect domain can be customized via `KONG_CUSTOM_DOMAIN` environment variable.
This is useful when testing against a development environment or a custom domain.
This is optional and the default value is `konghq.com`.

## Identity operations and the server URL

On the public API, identity operations are served only by the global API, so a request hook rewrites
their host to `global.api.konghq.com`, or to `global.api.konghq.tech` when the requested host ends
with `api.konghq.tech`. The port is ignored when the host is compared.

The rewrite is skipped when the server URL carries a path: a private entry point such as
`https://us.svc.konghq.tech/api` serves the identity path itself, and the global host does not.
A lone `/` is not a path, so `https://us.api.konghq.com/` keeps the global host.

A private entry point has to be configured with its path prefix: a private host without one, such as
`https://us.svc.konghq.tech`, cannot be told apart from a public host and is still sent to the global
API.

Such a URL has to be set per operation, e.g.
`s.Me.GetOrganizationsMe(ctx, operations.WithServerURL("https://us.svc.konghq.tech/api"))`; the
client-level `WithServerURL` option does not reach identity operations. The `KONG_CUSTOM_DOMAIN` swap
described above still applies afterwards.

## Request and response logging

Request and response logging can be enabled via:

- `KONNECT_SDK_HTTP_DUMP_REQUEST`: Set to `true` to log request bodies.
- `KONNECT_SDK_HTTP_DUMP_RESPONSE`: Set to `true` to log response bodies.
- `KONNECT_SDK_HTTP_DUMP_RESPONSE_ERROR`: Set to `true` to log response bodies only when the response status code is an error.
