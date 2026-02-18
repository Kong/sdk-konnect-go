# ServiceName


## Supported Types

### StringFieldFilter

```go
serviceName := components.CreateServiceNameStringFieldFilter(components.StringFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch serviceName.Type {
	case components.ServiceNameTypeStringFieldFilter:
		// serviceName.StringFieldFilter is populated
}
```
