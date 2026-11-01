# RegistryReferenceAgentNameFilter

Filters on the given string field value by either exact or fuzzy match.


## Supported Types

### RegistryReferenceAgentNameFilterStringFieldEqualsFilter

```go
registryReferenceAgentNameFilter := components.CreateRegistryReferenceAgentNameFilterRegistryReferenceAgentNameFilterStringFieldEqualsFilter(components.RegistryReferenceAgentNameFilterStringFieldEqualsFilter{/* values here */})
```

### RegistryReferenceAgentNameFilterStringFieldContainsFilter

```go
registryReferenceAgentNameFilter := components.CreateRegistryReferenceAgentNameFilterRegistryReferenceAgentNameFilterStringFieldContainsFilter(components.RegistryReferenceAgentNameFilterStringFieldContainsFilter{/* values here */})
```

### StringFieldOContainsFilter

```go
registryReferenceAgentNameFilter := components.CreateRegistryReferenceAgentNameFilterStringFieldOContainsFilter(components.StringFieldOContainsFilter{/* values here */})
```

### RegistryReferenceAgentNameFilterStringFieldOEQFilter

```go
registryReferenceAgentNameFilter := components.CreateRegistryReferenceAgentNameFilterRegistryReferenceAgentNameFilterStringFieldOEQFilter(components.RegistryReferenceAgentNameFilterStringFieldOEQFilter{/* values here */})
```

### RegistryReferenceAgentNameFilterStringFieldNEQFilter

```go
registryReferenceAgentNameFilter := components.CreateRegistryReferenceAgentNameFilterRegistryReferenceAgentNameFilterStringFieldNEQFilter(components.RegistryReferenceAgentNameFilterStringFieldNEQFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch registryReferenceAgentNameFilter.Type {
	case components.RegistryReferenceAgentNameFilterTypeRegistryReferenceAgentNameFilterStringFieldEqualsFilter:
		// registryReferenceAgentNameFilter.RegistryReferenceAgentNameFilterStringFieldEqualsFilter is populated
	case components.RegistryReferenceAgentNameFilterTypeRegistryReferenceAgentNameFilterStringFieldContainsFilter:
		// registryReferenceAgentNameFilter.RegistryReferenceAgentNameFilterStringFieldContainsFilter is populated
	case components.RegistryReferenceAgentNameFilterTypeStringFieldOContainsFilter:
		// registryReferenceAgentNameFilter.StringFieldOContainsFilter is populated
	case components.RegistryReferenceAgentNameFilterTypeRegistryReferenceAgentNameFilterStringFieldOEQFilter:
		// registryReferenceAgentNameFilter.RegistryReferenceAgentNameFilterStringFieldOEQFilter is populated
	case components.RegistryReferenceAgentNameFilterTypeRegistryReferenceAgentNameFilterStringFieldNEQFilter:
		// registryReferenceAgentNameFilter.RegistryReferenceAgentNameFilterStringFieldNEQFilter is populated
}
```
