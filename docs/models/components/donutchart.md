# DonutChart

A chart that can display one-dimensional data in a hollow, segmented circle.  To use this chart, ensure that
the query includes only one dimension (not `time`).



## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ChartTitle`                                                           | `*string`                                                              | :heavy_minus_sign:                                                     | The title of the chart, which is displayed in the tile's header.       |
| `Type`                                                                 | [components.DonutChartType](../../models/components/donutcharttype.md) | :heavy_check_mark:                                                     | N/A                                                                    |