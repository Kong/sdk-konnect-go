# KongService


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        | Example                                            |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `ID`                                               | `*string`                                          | :heavy_minus_sign:                                 | Service ID                                         |                                                    |
| `Name`                                             | `string`                                           | :heavy_check_mark:                                 | Service name                                       | mcp-test-service                                   |
| `Protocol`                                         | `string`                                           | :heavy_check_mark:                                 | The protocol used to communicate with the Service. | https                                              |
| `Host`                                             | `string`                                           | :heavy_check_mark:                                 | Service host                                       | mcp-test-service.svc.cluster.local                 |
| `Port`                                             | `int64`                                            | :heavy_check_mark:                                 | Service port                                       | 80                                                 |
| `Path`                                             | `string`                                           | :heavy_check_mark:                                 | Service path                                       | /                                                  |