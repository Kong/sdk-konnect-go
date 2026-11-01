# RegistryReferenceAgentNameFilterStringFieldEqualsFilter

Filters on the given string field value by exact match.


## Supported Types

### 

```go
registryReferenceAgentNameFilterStringFieldEqualsFilter := components.CreateRegistryReferenceAgentNameFilterStringFieldEqualsFilterStr(string{/* values here */})
```

### StringFieldEqualsComparison

```go
registryReferenceAgentNameFilterStringFieldEqualsFilter := components.CreateRegistryReferenceAgentNameFilterStringFieldEqualsFilterStringFieldEqualsComparison(components.StringFieldEqualsComparison{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch registryReferenceAgentNameFilterStringFieldEqualsFilter.Type {
	case components.RegistryReferenceAgentNameFilterStringFieldEqualsFilterTypeStr:
		// registryReferenceAgentNameFilterStringFieldEqualsFilter.Str is populated
	case components.RegistryReferenceAgentNameFilterStringFieldEqualsFilterTypeStringFieldEqualsComparison:
		// registryReferenceAgentNameFilterStringFieldEqualsFilter.StringFieldEqualsComparison is populated
}
```
