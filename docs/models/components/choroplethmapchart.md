# ChoroplethMapChart

A chart that displays data on a world map. Each region on the map is colored based on the metric value.
This chart works only with the `api_usage` datasource and requires a single metric and a single dimension of `country_code`.
No additional dimensions are supported.



## Fields

| Field                                                                                  | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ChartTitle`                                                                           | `*string`                                                                              | :heavy_minus_sign:                                                                     | The title of the chart, which is displayed in the tile's header.                       |
| `Type`                                                                                 | [components.ChoroplethMapChartType](../../models/components/choroplethmapcharttype.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |