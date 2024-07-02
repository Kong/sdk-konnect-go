# PostSystemAccountsResponse


## Fields

| Field                                                                                                                                                                                                                                                   | Type                                                                                                                                                                                                                                                    | Required                                                                                                                                                                                                                                                | Description                                                                                                                                                                                                                                             | Example                                                                                                                                                                                                                                                 |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ContentType`                                                                                                                                                                                                                                           | *string*                                                                                                                                                                                                                                                | :heavy_check_mark:                                                                                                                                                                                                                                      | HTTP response content type for this operation                                                                                                                                                                                                           |                                                                                                                                                                                                                                                         |
| `StatusCode`                                                                                                                                                                                                                                            | *int*                                                                                                                                                                                                                                                   | :heavy_check_mark:                                                                                                                                                                                                                                      | HTTP response status code for this operation                                                                                                                                                                                                            |                                                                                                                                                                                                                                                         |
| `RawResponse`                                                                                                                                                                                                                                           | [*http.Response](https://pkg.go.dev/net/http#Response)                                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                                                                      | Raw HTTP response; suitable for custom response parsing                                                                                                                                                                                                 |                                                                                                                                                                                                                                                         |
| `SystemAccount`                                                                                                                                                                                                                                         | [*components.SystemAccount](../../models/components/systemaccount.md)                                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                                                                      | A response including a single system account.                                                                                                                                                                                                           | {<br/>"id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",<br/>"name": "Example System Account",<br/>"description": "This is a sample system account description.",<br/>"created_at": "2022-08-24T14:15:22Z",<br/>"updated_at": "2022-10-05T10:33:49Z",<br/>"konnect_managed": false<br/>} |