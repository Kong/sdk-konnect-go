# ServiceUnavailable

Error response for temporary service unavailability.


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     | Example                                                         |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `Status`                                                        | [sdkerrors.Status](../../models/sdkerrors/status.md)            | :heavy_check_mark:                                              | The HTTP status code.                                           | 503                                                             |
| `Title`                                                         | *string*                                                        | :heavy_check_mark:                                              | The error response code.                                        | Service Unavailable                                             |
| `Instance`                                                      | *string*                                                        | :heavy_check_mark:                                              | The Konnect traceback code                                      | konnect:trace:2287285207635123011                               |
| `Detail`                                                        | **string*                                                       | :heavy_minus_sign:                                              | Details about the error.                                        | Could not retrieve permissions to check resource accessibility. |