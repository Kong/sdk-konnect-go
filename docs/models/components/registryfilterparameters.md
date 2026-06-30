# RegistryFilterParameters

Filter parameters for registry list endpoint


## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ID`                                                                                         | [*components.UUIDFieldFilter](../../models/components/uuidfieldfilter.md)                    | :heavy_minus_sign:                                                                           | Filter using **one** of the following operators: `eq`, `oeq`, `neq`                          |
| `Name`                                                                                       | [*components.StringFieldFilter](../../models/components/stringfieldfilter.md)                | :heavy_minus_sign:                                                                           | Filter using **one** of the following operators: `eq`, `oeq`, `neq`, `contains`, `ocontains` |
| `DisplayName`                                                                                | [*components.StringFieldFilter](../../models/components/stringfieldfilter.md)                | :heavy_minus_sign:                                                                           | Filter using **one** of the following operators: `eq`, `oeq`, `neq`, `contains`, `ocontains` |