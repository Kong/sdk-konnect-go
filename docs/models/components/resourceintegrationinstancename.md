# ResourceIntegrationInstanceName


## Supported Types

### StringFieldFilter

```go
resourceIntegrationInstanceName := components.CreateResourceIntegrationInstanceNameStringFieldFilter(components.StringFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch resourceIntegrationInstanceName.Type {
	case components.ResourceIntegrationInstanceNameTypeStringFieldFilter:
		// resourceIntegrationInstanceName.StringFieldFilter is populated
}
```
