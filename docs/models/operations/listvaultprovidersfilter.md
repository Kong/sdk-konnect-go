# ListVaultProvidersFilter

Filter vault providers by field. Supported filter fields:
- `name`: exact match or contains. Example: `filter[name]="my-provider"` or `filter[name][contains]="my-provider"`
- `template_name`: exact match on one or more template names. Example: `filter[template_name]="github"` or `filter[template_name][oeq]="github,slack"`



## Fields

| Field                                                                                     | Type                                                                                      | Required                                                                                  | Description                                                                               |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `Name`                                                                                    | [*components.LegacyStringFieldFilter](../../models/components/legacystringfieldfilter.md) | :heavy_minus_sign:                                                                        | Filter using **one** of the following operators: `eq`, `contains`                         |
| `TemplateName`                                                                            | [*operations.TemplateName](../../models/operations/templatename.md)                       | :heavy_minus_sign:                                                                        | Filter using **one** of the following operators: `eq`, `oeq`                              |