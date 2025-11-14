# PreconditionFailed

The error response object.


## Fields

| Field                             | Type                              | Required                          | Description                       | Example                           |
| --------------------------------- | --------------------------------- | --------------------------------- | --------------------------------- | --------------------------------- |
| `Status`                          | **int64*                          | :heavy_minus_sign:                | The HTTP status code.             | 412                               |
| `Title`                           | **string*                         | :heavy_minus_sign:                | The error response code.          | Precondition Failed               |
| `Instance`                        | **string*                         | :heavy_minus_sign:                | The Konnect traceback code.       | konnect:trace:1896611024257578096 |
| `Detail`                          | **string*                         | :heavy_minus_sign:                | Details about the error response. | IdP configuration not found       |