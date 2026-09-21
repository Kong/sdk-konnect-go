# PatchAISettingsRequest

Update AI settings for a portal.


## Fields

| Field                                                                                                   | Type                                                                                                    | Required                                                                                                | Description                                                                                             |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| `Enabled`                                                                                               | `bool`                                                                                                  | :heavy_check_mark:                                                                                      | Whether AI is enabled or not                                                                            |
| `Features`                                                                                              | [*components.PatchAISettingsRequestFeatures](../../models/components/patchaisettingsrequestfeatures.md) | :heavy_minus_sign:                                                                                      | AI Features configuration. Only nullable when top-level `enabled` is false                              |