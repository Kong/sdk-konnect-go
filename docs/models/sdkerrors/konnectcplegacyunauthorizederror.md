# KonnectCPLegacyUnauthorizedError

standard error


## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             | Example                                                 |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `Message`                                               | `*string`                                               | :heavy_minus_sign:                                      | A short summary of the problem.<br/>                    | Unauthorized                                            |
| `RawResponse`                                           | [*http.Response](https://pkg.go.dev/net/http#Response)  | :heavy_minus_sign:                                      | Raw HTTP response; suitable for custom response parsing |                                                         |