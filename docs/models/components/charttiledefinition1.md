# ChartTileDefinition1

The tile's definition, which consists of a query to fetch data and a visualization to render the data.
Charts and tables expect certain query types to render properly. The documentation for the individual visualization types has more information.



## Supported Types

### ChartTileDefinition

```go
chartTileDefinition1 := components.CreateChartTileDefinition1ChartTileDefinition(components.ChartTileDefinition{/* values here */})
```

### TableChartTileDefinition

```go
chartTileDefinition1 := components.CreateChartTileDefinition1TableChartTileDefinition(components.TableChartTileDefinition{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch chartTileDefinition1.Type {
	case components.ChartTileDefinition1TypeChartTileDefinition:
		// chartTileDefinition1.ChartTileDefinition is populated
	case components.ChartTileDefinition1TypeTableChartTileDefinition:
		// chartTileDefinition1.TableChartTileDefinition is populated
}
```
