# UnauthorizedError

standard error


## Fields

| Field                        | Type                         | Required                     | Description                  | Example                      |
| ---------------------------- | ---------------------------- | ---------------------------- | ---------------------------- | ---------------------------- |
| `Status`                     | *interface{}*                | :heavy_check_mark:           | N/A                          | 401                          |
| `Title`                      | *interface{}*                | :heavy_check_mark:           | N/A                          | Unauthorized                 |
| `Type`                       | *interface{}*                | :heavy_minus_sign:           | N/A                          | https://httpstatuses.com/401 |
| `Instance`                   | *interface{}*                | :heavy_check_mark:           | N/A                          | kong:trace:1234567890        |
| `Detail`                     | *interface{}*                | :heavy_check_mark:           | N/A                          | Invalid credentials          |