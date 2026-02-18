# ServiceSelectorSelectorParameters


## Supported Types

### StringCustomFieldSelectorParams

```go
serviceSelectorSelectorParameters := components.CreateServiceSelectorSelectorParametersStringCustomFieldSelectorParams(components.StringCustomFieldSelectorParams{/* values here */})
```

### NumberCustomFieldSelectorParams

```go
serviceSelectorSelectorParameters := components.CreateServiceSelectorSelectorParametersNumberCustomFieldSelectorParams(components.NumberCustomFieldSelectorParams{/* values here */})
```

### BooleanCustomFieldSelectorParams

```go
serviceSelectorSelectorParameters := components.CreateServiceSelectorSelectorParametersBooleanCustomFieldSelectorParams(components.BooleanCustomFieldSelectorParams{/* values here */})
```

### URLCustomFieldSelectorParams

```go
serviceSelectorSelectorParameters := components.CreateServiceSelectorSelectorParametersURLCustomFieldSelectorParams(components.URLCustomFieldSelectorParams{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch serviceSelectorSelectorParameters.Type {
	case components.ServiceSelectorSelectorParametersTypeStringCustomFieldSelectorParams:
		// serviceSelectorSelectorParameters.StringCustomFieldSelectorParams is populated
	case components.ServiceSelectorSelectorParametersTypeNumberCustomFieldSelectorParams:
		// serviceSelectorSelectorParameters.NumberCustomFieldSelectorParams is populated
	case components.ServiceSelectorSelectorParametersTypeBooleanCustomFieldSelectorParams:
		// serviceSelectorSelectorParameters.BooleanCustomFieldSelectorParams is populated
	case components.ServiceSelectorSelectorParametersTypeURLCustomFieldSelectorParams:
		// serviceSelectorSelectorParameters.URLCustomFieldSelectorParams is populated
}
```
