# JSONPathStrategy

Strategy using JSONPath expression to locate content for sanitization.


## Fields

| Field                                    | Type                                     | Required                                 | Description                              | Example                                  |
| ---------------------------------------- | ---------------------------------------- | ---------------------------------------- | ---------------------------------------- | ---------------------------------------- |
| `JSONPath`                               | `string`                                 | :heavy_check_mark:                       | JSONPath expression to locate the value. | $.user.email                             |