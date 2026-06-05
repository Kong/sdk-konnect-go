# ControlPlaneFilterParameters


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ID`                                                                     | [*components.ID](../../models/components/id.md)                          | :heavy_minus_sign:                                                       | Filter using **one** of the following operators: `eq`, `oeq`             |
| `Name`                                                                   | [*components.Name](../../models/components/name.md)                      | :heavy_minus_sign:                                                       | Filter using **one** of the following operators: `eq`, `neq`, `contains` |
| `ClusterType`                                                            | [*components.ClusterType](../../models/components/clustertype.md)        | :heavy_minus_sign:                                                       | Filter using **one** of the following operators: `eq`, `oeq`, `neq`      |
| `CloudGateway`                                                           | `*bool`                                                                  | :heavy_minus_sign:                                                       | Filter by a boolean value (true/false).                                  |