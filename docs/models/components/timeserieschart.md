# TimeseriesChart

A chart that can render timeseries data -- data from a query that has `time` as a dimension -- as lines or bars.

This type of chart can support:

- One or more metrics: `{ metrics: ["response_latency_p99", "response_latency_p95"], dimensions: ["time"] }`
- One metric plus one non-time dimension: `{ metrics: ["request_count"], dimensions: ["time", "gateway_service"] }` 

Either way, ensure that `time` is in the list of query dimensions.



## Fields

| Field                                                                                                                             | Type                                                                                                                              | Required                                                                                                                          | Description                                                                                                                       |
| --------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------- |
| `ChartTitle`                                                                                                                      | `*string`                                                                                                                         | :heavy_minus_sign:                                                                                                                | The title of the chart, which is displayed in the tile's header.                                                                  |
| `Type`                                                                                                                            | [components.TimeseriesChartType](../../models/components/timeseriescharttype.md)                                                  | :heavy_check_mark:                                                                                                                | N/A                                                                                                                               |
| `Stacked`                                                                                                                         | `*bool`                                                                                                                           | :heavy_minus_sign:                                                                                                                | Whether to stack the bars or lines (implicitly adding them together to form a total), or leave them independent from each other.<br/> |