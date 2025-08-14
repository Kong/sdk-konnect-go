# ConsumerGroup


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `CreatedAt`                                                   | **int64*                                                      | :heavy_minus_sign:                                            | Unix epoch when the resource was created.                     |
| `ID`                                                          | **string*                                                     | :heavy_minus_sign:                                            | A string representing a UUID (universally unique identifier). |
| `Name`                                                        | *string*                                                      | :heavy_check_mark:                                            | The name of the consumer group.                               |
| `Tags`                                                        | []*string*                                                    | :heavy_minus_sign:                                            | A set of strings representing tags.                           |
| `UpdatedAt`                                                   | **int64*                                                      | :heavy_minus_sign:                                            | Unix epoch when the resource was last updated.                |