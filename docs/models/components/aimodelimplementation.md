# AiModelImplementation

AI Model Implementation.


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    | Example                                                        |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ID`                                                           | `string`                                                       | :heavy_check_mark:                                             | The unique identifier of the implementation (link) record.     | d2e1f0a9-8b7c-6d5e-4f3a-2b1c0d9e8f7a                           |
| `AiModelID`                                                    | `string`                                                       | :heavy_check_mark:                                             | The identifier of the parent AI Model.                         | 123e4567-e89b-12d3-a456-426614174000                           |
| `GatewayControlPlaneID`                                        | `string`                                                       | :heavy_check_mark:                                             | The AI Gateway control plane uuid the linked model belongs to. | 223e4567-e89b-12d3-a456-426614174999                           |
| `GatewayModelID`                                               | `string`                                                       | :heavy_check_mark:                                             | The text or uuid identifier of the linked AI Gateway model.    | gw-model-abc                                                   |
| `CreatedAt`                                                    | [time.Time](https://pkg.go.dev/time#Time)                      | :heavy_check_mark:                                             | An ISO-8601 timestamp representation of entity creation date.  | 2022-11-04T20:10:06.927Z                                       |
| `UpdatedAt`                                                    | [time.Time](https://pkg.go.dev/time#Time)                      | :heavy_check_mark:                                             | An ISO-8601 timestamp representation of entity update date.    | 2022-11-04T20:10:06.927Z                                       |