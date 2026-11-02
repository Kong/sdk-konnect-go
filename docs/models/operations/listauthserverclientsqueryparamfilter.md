# ListAuthServerClientsQueryParamFilter

Filter clients returned in the response.


## Fields

| Field                                                                                     | Type                                                                                      | Required                                                                                  | Description                                                                               |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `Name`                                                                                    | [*components.LegacyStringFieldFilter](../../models/components/legacystringfieldfilter.md) | :heavy_minus_sign:                                                                        | Filter using **one** of the following operators: `eq`, `contains`                         |
| `ID`                                                                                      | [*components.StringFieldEqualsFilter](../../models/components/stringfieldequalsfilter.md) | :heavy_minus_sign:                                                                        | Filters on the given string field value by exact match.                                   |
| `Labels`                                                                                  | [*components.StringFieldEqualsFilter](../../models/components/stringfieldequalsfilter.md) | :heavy_minus_sign:                                                                        | Filters on the given string field value by exact match.                                   |