# ConflictError

standard error


## Fields

| Field                        | Type                         | Required                     | Description                  | Example                      |
| ---------------------------- | ---------------------------- | ---------------------------- | ---------------------------- | ---------------------------- |
| `Status`                     | *interface{}*                | :heavy_check_mark:           | N/A                          | 409                          |
| `Title`                      | *interface{}*                | :heavy_check_mark:           | N/A                          | Conflict                     |
| `Type`                       | *interface{}*                | :heavy_minus_sign:           | N/A                          | https://httpstatuses.com/409 |
| `Instance`                   | *interface{}*                | :heavy_check_mark:           | N/A                          | kong:trace:1234567890        |
| `Detail`                     | *interface{}*                | :heavy_check_mark:           | N/A                          | Conflict                     |