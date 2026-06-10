# QueryFilterString

A query filter for a string attribute. Operators are mutually exclusive, only
one operator is allowed at a time.


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `Eq`                                                                       | `*string`                                                                  | :heavy_minus_sign:                                                         | The attribute equals the provided value.                                   |
| `Neq`                                                                      | `*string`                                                                  | :heavy_minus_sign:                                                         | The attribute does not equal the provided value.                           |
| `In`                                                                       | []`string`                                                                 | :heavy_minus_sign:                                                         | The attribute is one of the provided values.                               |
| `Nin`                                                                      | []`string`                                                                 | :heavy_minus_sign:                                                         | The attribute is not one of the provided values.                           |
| `Contains`                                                                 | `*string`                                                                  | :heavy_minus_sign:                                                         | The attribute contains the provided value.                                 |
| `Ncontains`                                                                | `*string`                                                                  | :heavy_minus_sign:                                                         | The attribute does not contain the provided value.                         |
| `And`                                                                      | [][metering.QueryFilterString](../../models/metering/queryfilterstring.md) | :heavy_minus_sign:                                                         | Combines the provided filters with a logical AND.                          |
| `Or`                                                                       | [][metering.QueryFilterString](../../models/metering/queryfilterstring.md) | :heavy_minus_sign:                                                         | Combines the provided filters with a logical OR.                           |