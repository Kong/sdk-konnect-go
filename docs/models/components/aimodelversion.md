# AiModelVersion

A version of an AI Model, holding its `target_models`. Currently only exactly one version (`latest`) per model exists.



## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        | Example                                                            |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ID`                                                               | `string`                                                           | :heavy_check_mark:                                                 | The unique identifier of the version.                              | 9c4a1b2c-3d4e-5f60-7182-93a4b5c6d7e8                               |
| `AiModelID`                                                        | `string`                                                           | :heavy_check_mark:                                                 | The identifier of the parent AI Model.                             | 123e4567-e89b-12d3-a456-426614174000                               |
| `Version`                                                          | `*string`                                                          | :heavy_check_mark:                                                 | An optional user-supplied version label. `null` when unset.        | 1.0.0                                                              |
| `TargetModels`                                                     | [][components.TargetModel](../../models/components/targetmodel.md) | :heavy_check_mark:                                                 | The upstream LLM targets for this version.                         |                                                                    |
| `CreatedAt`                                                        | [time.Time](https://pkg.go.dev/time#Time)                          | :heavy_check_mark:                                                 | An ISO-8601 timestamp representation of entity creation date.      | 2022-11-04T20:10:06.927Z                                           |
| `UpdatedAt`                                                        | [time.Time](https://pkg.go.dev/time#Time)                          | :heavy_check_mark:                                                 | An ISO-8601 timestamp representation of entity update date.        | 2022-11-04T20:10:06.927Z                                           |