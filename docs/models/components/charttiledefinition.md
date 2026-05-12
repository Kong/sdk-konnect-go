# ChartTileDefinition

The tile's definition, which consists of a query to fetch data and a chart to render the data.
Note that some charts expect certain types of queries to render properly.  The documentation for the individual chart types has more information.



## Fields

| Field                                                | Type                                                 | Required                                             | Description                                          |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| `Query`                                              | [components.Query](../../models/components/query.md) | :heavy_check_mark:                                   | N/A                                                  |
| `Chart`                                              | [components.Chart](../../models/components/chart.md) | :heavy_check_mark:                                   | The type of chart to render.                         |