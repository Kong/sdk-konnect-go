# IntegrationInstanceDisplayName


## Supported Types

### StringFieldFilter

```go
integrationInstanceDisplayName := components.CreateIntegrationInstanceDisplayNameStringFieldFilter(components.StringFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch integrationInstanceDisplayName.Type {
	case components.IntegrationInstanceDisplayNameTypeStringFieldFilter:
		// integrationInstanceDisplayName.StringFieldFilter is populated
}
```
