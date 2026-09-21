# PluginScopeFilter

Filters on the given plugin scope value. Set `eq` for an exact match on a single scope, or `oeq` for a match on any of the comma-separated scopes. Exactly one of `eq` or `oeq` must be set.


## Fields

| Field                                                                               | Type                                                                                | Required                                                                            | Description                                                                         | Example                                                                             |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `Eq`                                                                                | [*components.Eq](../../models/components/eq.md)                                     | :heavy_minus_sign:                                                                  | Filters on the given plugin scope value with an exact match.                        | service                                                                             |
| `Oeq`                                                                               | `*string`                                                                           | :heavy_minus_sign:                                                                  | Filters on the given plugin scope values with an exact match for any of the values. | global,service                                                                      |