# CloudGatewaysStringFieldFilterOverride


## Supported Types

### CloudGatewaysStringFieldEqualsFilterOverride

```go
cloudGatewaysStringFieldFilterOverride := components.CreateCloudGatewaysStringFieldFilterOverrideCloudGatewaysStringFieldEqualsFilterOverride(components.CloudGatewaysStringFieldEqualsFilterOverride{/* values here */})
```

### StringFieldContainsFilter

```go
cloudGatewaysStringFieldFilterOverride := components.CreateCloudGatewaysStringFieldFilterOverrideStringFieldContainsFilter(components.StringFieldContainsFilter{/* values here */})
```

### StringFieldNEQFilter

```go
cloudGatewaysStringFieldFilterOverride := components.CreateCloudGatewaysStringFieldFilterOverrideStringFieldNEQFilter(components.StringFieldNEQFilter{/* values here */})
```

### StringFieldOEQFilter

```go
cloudGatewaysStringFieldFilterOverride := components.CreateCloudGatewaysStringFieldFilterOverrideStringFieldOEQFilter(components.StringFieldOEQFilter{/* values here */})
```

### StringFieldOContainsFilter

```go
cloudGatewaysStringFieldFilterOverride := components.CreateCloudGatewaysStringFieldFilterOverrideStringFieldOContainsFilter(components.StringFieldOContainsFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch cloudGatewaysStringFieldFilterOverride.Type {
	case components.CloudGatewaysStringFieldFilterOverrideTypeCloudGatewaysStringFieldEqualsFilterOverride:
		// cloudGatewaysStringFieldFilterOverride.CloudGatewaysStringFieldEqualsFilterOverride is populated
	case components.CloudGatewaysStringFieldFilterOverrideTypeStringFieldContainsFilter:
		// cloudGatewaysStringFieldFilterOverride.StringFieldContainsFilter is populated
	case components.CloudGatewaysStringFieldFilterOverrideTypeStringFieldNEQFilter:
		// cloudGatewaysStringFieldFilterOverride.StringFieldNEQFilter is populated
	case components.CloudGatewaysStringFieldFilterOverrideTypeStringFieldOEQFilter:
		// cloudGatewaysStringFieldFilterOverride.StringFieldOEQFilter is populated
	case components.CloudGatewaysStringFieldFilterOverrideTypeStringFieldOContainsFilter:
		// cloudGatewaysStringFieldFilterOverride.StringFieldOContainsFilter is populated
}
```
