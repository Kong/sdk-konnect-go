# AiModelVersionCreate

Request body for creating an AI Model Version.


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        | Example                                                            |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `Version`                                                          | `*string`                                                          | :heavy_minus_sign:                                                 | An optional user-supplied version label.<br/>                      | 1.0.0                                                              |
| `TargetModels`                                                     | [][components.TargetModel](../../models/components/targetmodel.md) | :heavy_check_mark:                                                 | The upstream LLM targets (full set).                               |                                                                    |