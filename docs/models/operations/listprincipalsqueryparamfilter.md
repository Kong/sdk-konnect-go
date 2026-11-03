# ListPrincipalsQueryParamFilter

Filter principals by field. Supported filter fields:
- `display_name`: exact match or contains. Example: `filter[display_name]="service-account"` or `filter[display_name][contains]="service-account"`
- `description`: exact match or contains. Example: `filter[description]="service principal"`
- `id`: exact match. Example: `filter[id]="<uuid>"`
- `labels.<key>`: exact match on a specific label key. Example: `filter[labels.env]="prod"`



## Fields

| Field                                                                                     | Type                                                                                      | Required                                                                                  | Description                                                                               |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `DisplayName`                                                                             | [*components.LegacyStringFieldFilter](../../models/components/legacystringfieldfilter.md) | :heavy_minus_sign:                                                                        | Filter using **one** of the following operators: `eq`, `contains`                         |
| `Description`                                                                             | [*components.LegacyStringFieldFilter](../../models/components/legacystringfieldfilter.md) | :heavy_minus_sign:                                                                        | Filter using **one** of the following operators: `eq`, `contains`                         |
| `ID`                                                                                      | [*components.StringFieldEqualsFilter](../../models/components/stringfieldequalsfilter.md) | :heavy_minus_sign:                                                                        | Filters on the given string field value by exact match.                                   |
| `Labels`                                                                                  | [*components.StringFieldEqualsFilter](../../models/components/stringfieldequalsfilter.md) | :heavy_minus_sign:                                                                        | Filters on the given string field value by exact match.                                   |