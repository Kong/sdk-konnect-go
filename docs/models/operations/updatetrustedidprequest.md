# UpdateTrustedIdpRequest


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `DirectoryID`                                                                        | `string`                                                                             | :heavy_check_mark:                                                                   | ID of the directory                                                                  |
| `TrustedIdpID`                                                                       | `string`                                                                             | :heavy_check_mark:                                                                   | Unique identifier of the trusted idp for the vault                                   |
| `ReplaceTrustedIdpBody`                                                              | [components.ReplaceTrustedIdpBody](../../models/components/replacetrustedidpbody.md) | :heavy_check_mark:                                                                   | Trusted IdP to be replaced                                                           |