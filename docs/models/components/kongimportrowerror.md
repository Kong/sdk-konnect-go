# KongImportRowError

One failing row of an import run, located within the uploaded file.


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `Row`                                                              | `int64`                                                            | :heavy_check_mark:                                                 | 1-based data-row number within the uploaded file.                  |
| `Field`                                                            | `*string`                                                          | :heavy_minus_sign:                                                 | The column the row failed on, when the failure is column-specific. |
| `Reason`                                                           | `string`                                                           | :heavy_check_mark:                                                 | N/A                                                                |