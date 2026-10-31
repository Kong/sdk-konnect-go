# AnyChartTileDefinition

The tile's definition, which consists of a query to fetch data and a visualization to render the data.
Charts and tables expect certain query types to render properly. The documentation for the individual visualization types has more information.



## Supported Types

### ChartTileDefinition

```go
anyChartTileDefinition := components.CreateAnyChartTileDefinitionChartTileDefinition(components.ChartTileDefinition{/* values here */})
```

### TableChartTileDefinition

```go
anyChartTileDefinition := components.CreateAnyChartTileDefinitionTableChartTileDefinition(components.TableChartTileDefinition{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch anyChartTileDefinition.Type {
	case components.AnyChartTileDefinitionTypeChartTileDefinition:
		// anyChartTileDefinition.ChartTileDefinition is populated
	case components.AnyChartTileDefinitionTypeTableChartTileDefinition:
		// anyChartTileDefinition.TableChartTileDefinition is populated
}
```
