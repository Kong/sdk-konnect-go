# GetPortalAssetFaviconRawResponse


## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `ContentType`                                           | *string*                                                | :heavy_check_mark:                                      | HTTP response content type for this operation           |
| `StatusCode`                                            | *int*                                                   | :heavy_check_mark:                                      | HTTP response status code for this operation            |
| `RawResponse`                                           | [*http.Response](https://pkg.go.dev/net/http#Response)  | :heavy_check_mark:                                      | Raw HTTP response; suitable for custom response parsing |
| `TwoHundredImagePngPortalImageAssetBlob`                | *io.ReadCloser*                                         | :heavy_minus_sign:                                      | Logo of the portal. Can be either png, jpeg or svg      |
| `TwoHundredImageJpegPortalImageAssetBlob`               | *io.ReadCloser*                                         | :heavy_minus_sign:                                      | Logo of the portal. Can be either png, jpeg or svg      |
| `TwoHundredImageSvgPlusXMLPortalImageAssetBlob`         | *io.ReadCloser*                                         | :heavy_minus_sign:                                      | Logo of the portal. Can be either png, jpeg or svg      |