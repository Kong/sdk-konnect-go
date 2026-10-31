# ListGroupMembershipFilter

Filter group membership entries by field. Supported filter fields:
- `id`: exact match on the membership entry id. Example: `filter[id]="<uuid>"`
- `display_name`: exact match or contains on the member's display name, case-insensitive. Example: `filter[display_name]="jane.doe"` or `filter[display_name][contains]="jane"`



## Fields

| Field                                                                                     | Type                                                                                      | Required                                                                                  | Description                                                                               |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `ID`                                                                                      | [*components.StringFieldEqualsFilter](../../models/components/stringfieldequalsfilter.md) | :heavy_minus_sign:                                                                        | Filters on the given string field value by exact match.                                   |
| `DisplayName`                                                                             | [*components.LegacyStringFieldFilter](../../models/components/legacystringfieldfilter.md) | :heavy_minus_sign:                                                                        | Filter using **one** of the following operators: `eq`, `contains`                         |