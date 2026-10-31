# KongImportReport

Result of an import run.


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `Counts`                                                                         | [components.Counts](../../models/components/counts.md)                           | :heavy_check_mark:                                                               | N/A                                                                              |
| `Warnings`                                                                       | [][components.KongImportWarning](../../models/components/kongimportwarning.md)   | :heavy_check_mark:                                                               | The run succeeded, but a human should look at these rows.                        |
| `Errors`                                                                         | [][components.KongImportRowError](../../models/components/kongimportrowerror.md) | :heavy_check_mark:                                                               | N/A                                                                              |