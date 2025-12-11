# Display

Instructs how to interpret identifiers in the query result.
The structure consists of: dimension -> ID -> entity metadata.
Since entity names within Konnect may overlap, ids are returned and used for exact grouping.
This blob may be used to join and display results in a more human-friendly manner.
Since IDs represent entities in the system at all points in time, historical data will still exist even if an entity is deleted.
These such entities are marked with the `deleted: true` value.
Another special case is the `____OTHER____` entity.
When the limit of cardinality is reached for a second dimension, the system will group all other entities into the ____OTHER____ grouping.
This prevents returning an overwhelming number of long tail entities.



## Fields

| Field                                                                                                       | Type                                                                                                        | Required                                                                                                    | Description                                                                                                 |
| ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- |
| `Service`                                                                                                   | map[string][components.MetricDimensionDisplayEntry](../../models/components/metricdimensiondisplayentry.md) | :heavy_minus_sign:                                                                                          | N/A                                                                                                         |
| `Severity`                                                                                                  | map[string][components.MetricDimensionDisplayEntry](../../models/components/metricdimensiondisplayentry.md) | :heavy_minus_sign:                                                                                          | N/A                                                                                                         |
| `State`                                                                                                     | map[string][components.MetricDimensionDisplayEntry](../../models/components/metricdimensiondisplayentry.md) | :heavy_minus_sign:                                                                                          | N/A                                                                                                         |
| `Type`                                                                                                      | map[string][components.MetricDimensionDisplayEntry](../../models/components/metricdimensiondisplayentry.md) | :heavy_minus_sign:                                                                                          | N/A                                                                                                         |
| `Source`                                                                                                    | map[string][components.MetricDimensionDisplayEntry](../../models/components/metricdimensiondisplayentry.md) | :heavy_minus_sign:                                                                                          | N/A                                                                                                         |
| `Environment`                                                                                               | map[string][components.MetricDimensionDisplayEntry](../../models/components/metricdimensiondisplayentry.md) | :heavy_minus_sign:                                                                                          | N/A                                                                                                         |
| `Region`                                                                                                    | map[string][components.MetricDimensionDisplayEntry](../../models/components/metricdimensiondisplayentry.md) | :heavy_minus_sign:                                                                                          | N/A                                                                                                         |