# DeclarativeConfigHTTPRoute


## Fields

| Field                           | Type                            | Required                        | Description                     |
| ------------------------------- | ------------------------------- | ------------------------------- | ------------------------------- |
| `Name`                          | `string`                        | :heavy_check_mark:              | The name of the route.          |
| `Match`                         | `string`                        | :heavy_check_mark:              | The router match expression.    |
| `ReqScript`                     | `*string`                       | :heavy_minus_sign:              | The pd script for the request.  |
| `RespScript`                    | `*string`                       | :heavy_minus_sign:              | The pd script for the response. |