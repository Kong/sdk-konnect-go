# Filters

Filters to apply to the query.


## Fields

| Field                                                                                                                                | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `Dimensions`                                                                                                                         | map[string][metering.QueryFilterStringMapItem](../../models/metering/queryfilterstringmapitem.md)                                    | :heavy_minus_sign:                                                                                                                   | Filters to apply to the dimensions of the query. For `subject` and `customer_id`<br/>only equals ("eq", "in") comparisons are supported. |