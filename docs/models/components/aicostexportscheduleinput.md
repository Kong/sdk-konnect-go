# AICostExportScheduleInput

The body of a request that creates or replaces an export schedule. It carries the full schedule without `id`, `next_run_at`, `created_at`, and `updated_at`, which the server sets.


## Supported Types

### AICostExportScheduleInput1

```go
aiCostExportScheduleInput := components.CreateAICostExportScheduleInputAICostExportScheduleInput1(components.AICostExportScheduleInput1{/* values here */})
```

### AICostExportScheduleInput2

```go
aiCostExportScheduleInput := components.CreateAICostExportScheduleInputAICostExportScheduleInput2(components.AICostExportScheduleInput2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiCostExportScheduleInput.Type {
	case components.AICostExportScheduleInputTypeAICostExportScheduleInput1:
		// aiCostExportScheduleInput.AICostExportScheduleInput1 is populated
	case components.AICostExportScheduleInputTypeAICostExportScheduleInput2:
		// aiCostExportScheduleInput.AICostExportScheduleInput2 is populated
}
```
