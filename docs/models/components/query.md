# Query


## Supported Types

### AdvancedQuery

```go
query := components.CreateQueryAPIUsage(components.AdvancedQuery{/* values here */})
```

### LLMQuery

```go
query := components.CreateQueryLlmUsage(components.LLMQuery{/* values here */})
```

### AgenticQuery

```go
query := components.CreateQueryAgenticUsage(components.AgenticQuery{/* values here */})
```

### PlatformQuery

```go
query := components.CreateQueryPlatformUsage(components.PlatformQuery{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch query.Type {
	case components.QueryTypeAPIUsage:
		// query.AdvancedQuery is populated
	case components.QueryTypeLlmUsage:
		// query.LLMQuery is populated
	case components.QueryTypeAgenticUsage:
		// query.AgenticQuery is populated
	case components.QueryTypePlatformUsage:
		// query.PlatformQuery is populated
}
```
