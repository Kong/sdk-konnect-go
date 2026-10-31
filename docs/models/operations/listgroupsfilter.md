# ListGroupsFilter

Filter groups by field. Supported filter fields:
- `name`: exact match or contains. Example: `filter[name]="PlatformEngineers"` or `filter[name][contains]="Platform"`
- `description`: exact match or contains. Example: `filter[description]="platform team"`
- `id`: exact match. Example: `filter[id]="<uuid>"`
- `labels.<key>`: exact match on a specific label key. Example: `filter[labels.env]="prod"`
- `type`: exact match on one or more group types. Accepted values: `group`, `team`, `project`. Example: `filter[type]="team"` or `filter[type][oeq]="group,team"`



## Fields

| Field                                                                                     | Type                                                                                      | Required                                                                                  | Description                                                                               |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `Name`                                                                                    | [*components.LegacyStringFieldFilter](../../models/components/legacystringfieldfilter.md) | :heavy_minus_sign:                                                                        | Filter using **one** of the following operators: `eq`, `contains`                         |
| `Description`                                                                             | [*components.LegacyStringFieldFilter](../../models/components/legacystringfieldfilter.md) | :heavy_minus_sign:                                                                        | Filter using **one** of the following operators: `eq`, `contains`                         |
| `ID`                                                                                      | [*components.StringFieldEqualsFilter](../../models/components/stringfieldequalsfilter.md) | :heavy_minus_sign:                                                                        | Filters on the given string field value by exact match.                                   |
| `Labels`                                                                                  | [*components.StringFieldEqualsFilter](../../models/components/stringfieldequalsfilter.md) | :heavy_minus_sign:                                                                        | Filters on the given string field value by exact match.                                   |
| `Type`                                                                                    | [*operations.Type](../../models/operations/type.md)                                       | :heavy_minus_sign:                                                                        | Filter using **one** of the following operators: `eq`, `oeq`                              |