# APISlugConflict

Conflict - `slug` property must be unique


## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `Status`                                                | `float64`                                               | :heavy_check_mark:                                      | N/A                                                     |
| `Title`                                                 | `string`                                                | :heavy_check_mark:                                      | N/A                                                     |
| `Type`                                                  | `*string`                                               | :heavy_minus_sign:                                      | N/A                                                     |
| `Instance`                                              | `string`                                                | :heavy_check_mark:                                      | N/A                                                     |
| `RawResponse`                                           | [*http.Response](https://pkg.go.dev/net/http#Response)  | :heavy_minus_sign:                                      | Raw HTTP response; suitable for custom response parsing |