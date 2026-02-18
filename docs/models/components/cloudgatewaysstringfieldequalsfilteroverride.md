# CloudGatewaysStringFieldEqualsFilterOverride

Filters on the given string field value by exact match.


## Supported Types

### 

```go
cloudGatewaysStringFieldEqualsFilterOverride := components.CreateCloudGatewaysStringFieldEqualsFilterOverrideStr(string{/* values here */})
```

### StringFieldEqualsComparison

```go
cloudGatewaysStringFieldEqualsFilterOverride := components.CreateCloudGatewaysStringFieldEqualsFilterOverrideStringFieldEqualsComparison(components.StringFieldEqualsComparison{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch cloudGatewaysStringFieldEqualsFilterOverride.Type {
	case components.CloudGatewaysStringFieldEqualsFilterOverrideTypeStr:
		// cloudGatewaysStringFieldEqualsFilterOverride.Str is populated
	case components.CloudGatewaysStringFieldEqualsFilterOverrideTypeStringFieldEqualsComparison:
		// cloudGatewaysStringFieldEqualsFilterOverride.StringFieldEqualsComparison is populated
}
```
