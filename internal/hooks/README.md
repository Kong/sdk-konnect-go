# Configuration through environment variables

This SDK can be configured via the environment variables.

## Konnect domain

Konnect domain can be customized via `KONG_CUSTOM_DOMAIN` environment variable.
This is useful when testing against a development environment or a custom domain.
This is optional and the default value is `konghq.com`.

## Request and response logging

Request and response logging can be enabled via:

- `KONNECT_SDK_HTTP_DUMP_REQUEST`: Set to `true` to log request bodies.
- `KONNECT_SDK_HTTP_DUMP_RESPONSE`: Set to `true` to log response bodies.
- `KONNECT_SDK_HTTP_DUMP_RESPONSE_ERROR`: Set to `true` to log response bodies only when the response status code is an error.
