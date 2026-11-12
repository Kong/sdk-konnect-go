# UpdateVaultKeyRequest


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `DirectoryID`                                                                | `string`                                                                     | :heavy_check_mark:                                                           | ID of the directory                                                          |
| `KeyID`                                                                      | `string`                                                                     | :heavy_check_mark:                                                           | ID of the vault key                                                          |
| `PatchVaultKeyBody`                                                          | [components.PatchVaultKeyBody](../../models/components/patchvaultkeybody.md) | :heavy_check_mark:                                                           | Fields to update on a vault key                                              |