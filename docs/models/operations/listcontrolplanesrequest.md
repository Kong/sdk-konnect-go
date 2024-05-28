# ListControlPlanesRequest


## Fields

| Field                                                                                                 | Type                                                                                                  | Required                                                                                              | Description                                                                                           | Example                                                                                               |
| ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- |
| `Page`                                                                                                | [*components.OffsetPage](../../models/components/offsetpage.md)                                       | :heavy_minus_sign:                                                                                    | Pagination related query parameters                                                                   |                                                                                                       |
| `FilterNameEq`                                                                                        | **string*                                                                                             | :heavy_minus_sign:                                                                                    | Filter by direct equality comparison of the name property with a supplied value.                      | test                                                                                                  |
| `FilterName`                                                                                          | **string*                                                                                             | :heavy_minus_sign:                                                                                    | Filter by direct equality comparison (short-hand) of the name property with a supplied value.         | test                                                                                                  |
| `FilterNameContains`                                                                                  | **string*                                                                                             | :heavy_minus_sign:                                                                                    | Filter by contains comparison of the name property with a supplied substring.                         | test                                                                                                  |
| `FilterNameNeq`                                                                                       | **string*                                                                                             | :heavy_minus_sign:                                                                                    | Filter by non-equality comparison of the name property with a supplied value.                         | test                                                                                                  |
| `FilterIDEq`                                                                                          | **string*                                                                                             | :heavy_minus_sign:                                                                                    | Filter by direct equality comparison of the id property with a supplied value.                        | 7f9fd312-a987-4628-b4c5-bb4f4fddd5f7                                                                  |
| `FilterID`                                                                                            | **string*                                                                                             | :heavy_minus_sign:                                                                                    | Filter by direct equality comparison (short-hand) of the id property with a supplied value.           | 7f9fd312-a987-4628-b4c5-bb4f4fddd5f7                                                                  |
| `FilterIDOeq`                                                                                         | **string*                                                                                             | :heavy_minus_sign:                                                                                    | Filter by direct equality comparison of id property with multiple supplied values.                    | some-value,some-other-value                                                                           |
| `FilterClusterTypeEq`                                                                                 | **string*                                                                                             | :heavy_minus_sign:                                                                                    | Filter by direct equality comparison of the cluster_type property with a supplied value.              | CLUSTER_TYPE_CONTROL_PLANE                                                                            |
| `FilterClusterType`                                                                                   | **string*                                                                                             | :heavy_minus_sign:                                                                                    | Filter by direct equality comparison (short-hand) of the cluster_type property with a supplied value. | CLUSTER_TYPE_CONTROL_PLANE                                                                            |
| `FilterClusterTypeNeq`                                                                                | **string*                                                                                             | :heavy_minus_sign:                                                                                    | Filter by non-equality comparison of the cluster_type property with a supplied value.                 | test                                                                                                  |
| `Labels`                                                                                              | **string*                                                                                             | :heavy_minus_sign:                                                                                    | Filter control planes in the response by associated labels.                                           | key:value,existCheck                                                                                  |