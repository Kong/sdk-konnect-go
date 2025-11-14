# ListTeamUsersQueryParamFilter

Filter users returned in the response.


## Fields

| Field                                                                                     | Type                                                                                      | Required                                                                                  | Description                                                                               |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `ID`                                                                                      | [*components.StringFieldEqualsFilter](../../models/components/stringfieldequalsfilter.md) | :heavy_minus_sign:                                                                        | Filters on the given string field value by exact match.                                   |
| `Email`                                                                                   | [*components.LegacyStringFieldFilter](../../models/components/legacystringfieldfilter.md) | :heavy_minus_sign:                                                                        | Filters on the given string field value by fuzzy match.                                   |
| `FullName`                                                                                | [*components.LegacyStringFieldFilter](../../models/components/legacystringfieldfilter.md) | :heavy_minus_sign:                                                                        | Filters on the given string field value by fuzzy match.                                   |
| `Active`                                                                                  | **bool*                                                                                   | :heavy_minus_sign:                                                                        | Filter by a boolean value (true/false).                                                   |