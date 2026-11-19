# AICostExportInput

The body of a request that creates a one-time export: the dataset, the file format and compression, the time range, and optional filters and labels. The labels can't change after creation.


## Supported Types

### AICostExportInput1

```go
aiCostExportInput := components.CreateAICostExportInputAICostExportInput1(components.AICostExportInput1{/* values here */})
```

### AICostExportInput2

```go
aiCostExportInput := components.CreateAICostExportInputAICostExportInput2(components.AICostExportInput2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiCostExportInput.Type {
	case components.AICostExportInputTypeAICostExportInput1:
		// aiCostExportInput.AICostExportInput1 is populated
	case components.AICostExportInputTypeAICostExportInput2:
		// aiCostExportInput.AICostExportInput2 is populated
}
```
