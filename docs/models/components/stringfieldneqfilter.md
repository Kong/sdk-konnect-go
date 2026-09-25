# StringFieldNEQFilter

Filter credit balance by currency code. When historical custom currencies reuse
a code, each managed currency is returned as a separate balance row.


## Fields

| Field              | Type               | Required           | Description        |
| ------------------ | ------------------ | ------------------ | ------------------ |
| `Eq`               | `*string`          | :heavy_minus_sign: | N/A                |
| `Oeq`              | `string`           | :heavy_check_mark: | N/A                |
| `Neq`              | `string`           | :heavy_check_mark: | N/A                |