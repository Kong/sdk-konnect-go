# ServiceSelector


## Supported Types

### AllEntitiesSelector

```go
serviceSelector := components.CreateServiceSelectorAllEntitiesSelector(components.AllEntitiesSelector{/* values here */})
```

### ByIDsSelector

```go
serviceSelector := components.CreateServiceSelectorByIDsSelector(components.ByIDsSelector{/* values here */})
```

### ByNameSelector

```go
serviceSelector := components.CreateServiceSelectorByNameSelector(components.ByNameSelector{/* values here */})
```

### ByDisplayNameSelector

```go
serviceSelector := components.CreateServiceSelectorByDisplayNameSelector(components.ByDisplayNameSelector{/* values here */})
```

### ByCustomFieldSelector

```go
serviceSelector := components.CreateServiceSelectorByCustomFieldSelector(components.ByCustomFieldSelector{/* values here */})
```

### ByLabelSelector

```go
serviceSelector := components.CreateServiceSelectorByLabelSelector(components.ByLabelSelector{/* values here */})
```

### HasLabelKeySelector

```go
serviceSelector := components.CreateServiceSelectorHasLabelKeySelector(components.HasLabelKeySelector{/* values here */})
```

### HasDocsSelector

```go
serviceSelector := components.CreateServiceSelectorHasDocsSelector(components.HasDocsSelector{/* values here */})
```

### HasApisSelector

```go
serviceSelector := components.CreateServiceSelectorHasApisSelector(components.HasApisSelector{/* values here */})
```

### HasResourcesSelector

```go
serviceSelector := components.CreateServiceSelectorHasResourcesSelector(components.HasResourcesSelector{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch serviceSelector.Type {
	case components.ServiceSelectorTypeAllEntitiesSelector:
		// serviceSelector.AllEntitiesSelector is populated
	case components.ServiceSelectorTypeByIDsSelector:
		// serviceSelector.ByIDsSelector is populated
	case components.ServiceSelectorTypeByNameSelector:
		// serviceSelector.ByNameSelector is populated
	case components.ServiceSelectorTypeByDisplayNameSelector:
		// serviceSelector.ByDisplayNameSelector is populated
	case components.ServiceSelectorTypeByCustomFieldSelector:
		// serviceSelector.ByCustomFieldSelector is populated
	case components.ServiceSelectorTypeByLabelSelector:
		// serviceSelector.ByLabelSelector is populated
	case components.ServiceSelectorTypeHasLabelKeySelector:
		// serviceSelector.HasLabelKeySelector is populated
	case components.ServiceSelectorTypeHasDocsSelector:
		// serviceSelector.HasDocsSelector is populated
	case components.ServiceSelectorTypeHasApisSelector:
		// serviceSelector.HasApisSelector is populated
	case components.ServiceSelectorTypeHasResourcesSelector:
		// serviceSelector.HasResourcesSelector is populated
}
```
