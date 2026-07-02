# PreviewAPISpecResponse


## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `ContentType`                                           | `string`                                                | :heavy_check_mark:                                      | HTTP response content type for this operation           |
| `StatusCode`                                            | `int`                                                   | :heavy_check_mark:                                      | HTTP response status code for this operation            |
| `RawResponse`                                           | [*http.Response](https://pkg.go.dev/net/http#Response)  | :heavy_check_mark:                                      | Raw HTTP response; suitable for custom response parsing |
| `APISpecContentsPreview`                                | `*string`                                               | :heavy_minus_sign:                                      | Response object containing raw API spec string contents |
| `APISpecContentsPreviewYaml`                            | `*string`                                               | :heavy_minus_sign:                                      | Response object containing raw API spec string contents |