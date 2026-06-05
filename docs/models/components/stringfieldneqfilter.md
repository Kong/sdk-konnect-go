# StringFieldNEQFilter

Filter charges by status.

Supported statuses are:

- `created`
- `active`
- `final`
- `deleted`

If omitted, all statuses are returned except for `deleted`.


## Fields

| Field              | Type               | Required           | Description        |
| ------------------ | ------------------ | ------------------ | ------------------ |
| `Eq`               | `*string`          | :heavy_minus_sign: | N/A                |
| `Oeq`              | `string`           | :heavy_check_mark: | N/A                |
| `Neq`              | `string`           | :heavy_check_mark: | N/A                |