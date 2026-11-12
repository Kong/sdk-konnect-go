# ListCreditGrantsParamsFilterFeatureKeyStringFieldNEQFilter

Select grants spendable on any requested feature, including grants without
feature restrictions. Keys are case-sensitive. Supports `eq` (or shorthand) and
`oeq` with comma-separated keys. Omit to include all features; use
`exists=false` for grants without feature restrictions.


## Fields

| Field              | Type               | Required           | Description        |
| ------------------ | ------------------ | ------------------ | ------------------ |
| `Eq`               | `*string`          | :heavy_minus_sign: | N/A                |
| `Contains`         | `string`           | :heavy_check_mark: | N/A                |
| `Ocontains`        | `string`           | :heavy_check_mark: | N/A                |
| `Oeq`              | `string`           | :heavy_check_mark: | N/A                |
| `Neq`              | `string`           | :heavy_check_mark: | N/A                |