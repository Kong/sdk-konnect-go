# ListPortalTeamDevelopersQueryParamFilter

Filter developers returned in the response.


## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `Email`                                                                                      | [*components.StringFieldFilter](../../models/components/stringfieldfilter.md)                | :heavy_minus_sign:                                                                           | Filter using **one** of the following operators: `eq`, `oeq`, `neq`, `contains`, `ocontains` |
| `FullName`                                                                                   | [*components.StringFieldFilter](../../models/components/stringfieldfilter.md)                | :heavy_minus_sign:                                                                           | Filter using **one** of the following operators: `eq`, `oeq`, `neq`, `contains`, `ocontains` |
| `Attributes`                                                                                 | [*components.StringFieldFilter](../../models/components/stringfieldfilter.md)                | :heavy_minus_sign:                                                                           | Filter using **one** of the following operators: `eq`, `oeq`, `neq`, `contains`, `ocontains` |
| `Active`                                                                                     | **bool*                                                                                      | :heavy_minus_sign:                                                                           | Filter by a boolean value (true/false).                                                      |