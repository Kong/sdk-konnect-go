# IntegrationName


## Supported Types

### StringFieldFilter

```go
integrationName := components.CreateIntegrationNameStringFieldFilter(components.StringFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch integrationName.Type {
	case components.IntegrationNameTypeStringFieldFilter:
		// integrationName.StringFieldFilter is populated
}
```
