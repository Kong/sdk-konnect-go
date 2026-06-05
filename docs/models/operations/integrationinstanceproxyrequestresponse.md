# IntegrationInstanceProxyRequestResponse


## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `ContentType`                                           | `string`                                                | :heavy_check_mark:                                      | HTTP response content type for this operation           |
| `StatusCode`                                            | `int`                                                   | :heavy_check_mark:                                      | HTTP response status code for this operation            |
| `RawResponse`                                           | [*http.Response](https://pkg.go.dev/net/http#Response)  | :heavy_check_mark:                                      | Raw HTTP response; suitable for custom response parsing |
| `TwoXXApplicationJSONObject`                            | map[string]`any`                                        | :heavy_minus_sign:                                      | 2XX response code returned from the proxied API         |
| `ThreeXXApplicationJSONObject`                          | map[string]`any`                                        | :heavy_minus_sign:                                      | 3XX response code returned from the proxied API         |
| `Headers`                                               | map[string][]`string`                                   | :heavy_check_mark:                                      | N/A                                                     |