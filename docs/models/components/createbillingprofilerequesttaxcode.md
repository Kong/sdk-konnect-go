# CreateBillingProfileRequestTaxCode

Tax code reference.

When both `tax_code` and `tax_code_id` are provided, `tax_code` takes precedence.
When `stripe.code` is also provided, `tax_code` still wins and `stripe.code` is ignored.


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      | Example                                                          |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ID`                                                             | `string`                                                         | :heavy_check_mark:                                               | ULID (Universally Unique Lexicographically Sortable Identifier). | 01G65Z755AFWAKHE12NY0CQ9FH                                       |