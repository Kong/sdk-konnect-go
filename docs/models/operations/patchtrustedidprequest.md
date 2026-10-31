# PatchTrustedIdpRequest


## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `DirectoryID`                                                                      | `string`                                                                           | :heavy_check_mark:                                                                 | ID of the directory                                                                |
| `TrustedIdpID`                                                                     | `string`                                                                           | :heavy_check_mark:                                                                 | Unique identifier of the trusted idp for the vault                                 |
| `UpdateTrustedIdpBody`                                                             | [components.UpdateTrustedIdpBody](../../models/components/updatetrustedidpbody.md) | :heavy_check_mark:                                                                 | Fields to update on a trusted IdP                                                  |