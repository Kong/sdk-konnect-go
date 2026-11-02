# KongImportWarning

One row of an import run that succeeded but deserves attention.


## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `Row`                                                                          | `*int64`                                                                       | :heavy_minus_sign:                                                             | 1-based data-row number within the uploaded file; null for run-level warnings. |
| `Type`                                                                         | `string`                                                                       | :heavy_check_mark:                                                             | The warning kind (for example, email_changed or quota_shortfall).              |
| `Detail`                                                                       | `string`                                                                       | :heavy_check_mark:                                                             | N/A                                                                            |