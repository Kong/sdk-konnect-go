# IntegrationInstanceName


## Supported Types

### StringFieldFilter

```go
integrationInstanceName := components.CreateIntegrationInstanceNameStringFieldFilter(components.StringFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch integrationInstanceName.Type {
	case components.IntegrationInstanceNameTypeStringFieldFilter:
		// integrationInstanceName.StringFieldFilter is populated
}
```
