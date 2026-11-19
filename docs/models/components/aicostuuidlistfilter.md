# AICostUUIDListFilter

Filters a list field. The filter object carries exactly one operator:

- `has`, `has_any`: match a list that contains the value, or any of the values.
- `nhas`, `nhas_any`: match a list that doesn't contain the value, or contains none of the values. An empty list matches too.
- `exists`: tests whether the list has at least one value.
- `and`, `or`: combine nested filters.

UUIDs are compared ignoring case.


## Fields

| Field                                                                                           | Type                                                                                            | Required                                                                                        | Description                                                                                     |
| ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- |
| `Has`                                                                                           | `*string`                                                                                       | :heavy_minus_sign:                                                                              | Matches when the list contains this value.                                                      |
| `HasAny`                                                                                        | []`string`                                                                                      | :heavy_minus_sign:                                                                              | Matches when the list contains any of these values.<br/><br/>1 to 100 values.                   |
| `Nhas`                                                                                          | `*string`                                                                                       | :heavy_minus_sign:                                                                              | Matches when the list doesn't contain this value, including an empty list.                      |
| `NhasAny`                                                                                       | []`string`                                                                                      | :heavy_minus_sign:                                                                              | Matches when the list contains none of these values, including an empty list.<br/><br/>1 to 100 values. |
| `Exists`                                                                                        | `*bool`                                                                                         | :heavy_minus_sign:                                                                              | `true` matches a list with at least one value, and `false` matches an empty or missing list.    |
| `And`                                                                                           | [][components.AICostUUIDListFilter](../../models/components/aicostuuidlistfilter.md)            | :heavy_minus_sign:                                                                              | Matches when every nested filter matches.<br/><br/>1 to 10 filters.                             |
| `Or`                                                                                            | [][components.AICostUUIDListFilter](../../models/components/aicostuuidlistfilter.md)            | :heavy_minus_sign:                                                                              | Matches when at least one nested filter matches.<br/><br/>1 to 10 filters.                      |