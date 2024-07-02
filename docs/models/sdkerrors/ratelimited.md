# RateLimited

The error object


## Fields

| Field                                       | Type                                        | Required                                    | Description                                 | Example                                     |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| `Status`                                    | **int64*                                    | :heavy_minus_sign:                          | The HTTP response code                      | 429                                         |
| `Title`                                     | **string*                                   | :heavy_minus_sign:                          | The Error response                          | Rate Limited                                |
| `Instance`                                  | **string*                                   | :heavy_minus_sign:                          | The Konnect traceback ID.                   | konnect:trace:3674017986744198214           |
| `Detail`                                    | **string*                                   | :heavy_minus_sign:                          | Detailed explanation of the error response. | Too many requests                           |