# ResourceNames

If any of these entries match, the resource name matches for this rule. A maximum of 50 entries are allowed.


## Supported Types

### 

```go
resourceNames := components.CreateResourceNamesArrayOfEventGatewayACLResourceName([]components.EventGatewayACLResourceName{/* values here */})
```

### 

```go
resourceNames := components.CreateResourceNamesStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch resourceNames.Type {
	case components.ResourceNamesTypeArrayOfEventGatewayACLResourceName:
		// resourceNames.ArrayOfEventGatewayACLResourceName is populated
	case components.ResourceNamesTypeStr:
		// resourceNames.Str is populated
}
```
