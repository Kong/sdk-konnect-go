# Chart

The type of chart to render.


## Supported Types

### DonutChart

```go
chart := components.CreateChartDonut(components.DonutChart{/* values here */})
```

### TimeseriesChart

```go
chart := components.CreateChartTimeseriesLine(components.TimeseriesChart{/* values here */})
```

### TimeseriesChart

```go
chart := components.CreateChartTimeseriesBar(components.TimeseriesChart{/* values here */})
```

### BarChart

```go
chart := components.CreateChartHorizontalBar(components.BarChart{/* values here */})
```

### BarChart

```go
chart := components.CreateChartVerticalBar(components.BarChart{/* values here */})
```

### TopNChart

```go
chart := components.CreateChartTopN(components.TopNChart{/* values here */})
```

### SingleValueChart

```go
chart := components.CreateChartSingleValue(components.SingleValueChart{/* values here */})
```

### ChoroplethMapChart

```go
chart := components.CreateChartChoroplethMap(components.ChoroplethMapChart{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch chart.Type {
	case components.ChartTypeDonut:
		// chart.DonutChart is populated
	case components.ChartTypeTimeseriesLine:
		// chart.TimeseriesChart is populated
	case components.ChartTypeTimeseriesBar:
		// chart.TimeseriesChart is populated
	case components.ChartTypeHorizontalBar:
		// chart.BarChart is populated
	case components.ChartTypeVerticalBar:
		// chart.BarChart is populated
	case components.ChartTypeTopN:
		// chart.TopNChart is populated
	case components.ChartTypeSingleValue:
		// chart.SingleValueChart is populated
	case components.ChartTypeChoroplethMap:
		// chart.ChoroplethMapChart is populated
}
```
