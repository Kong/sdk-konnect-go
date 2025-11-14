# ACL


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `Consumer`                                                        | [*components.ACLConsumer](../../models/components/aclconsumer.md) | :heavy_minus_sign:                                                | N/A                                                               |
| `CreatedAt`                                                       | **int64*                                                          | :heavy_minus_sign:                                                | Unix epoch when the resource was created.                         |
| `Group`                                                           | *string*                                                          | :heavy_check_mark:                                                | N/A                                                               |
| `ID`                                                              | **string*                                                         | :heavy_minus_sign:                                                | A string representing a UUID (universally unique identifier).     |
| `Tags`                                                            | []*string*                                                        | :heavy_minus_sign:                                                | A set of strings representing tags.                               |