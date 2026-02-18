# Partial


## Supported Types

### PartialRedisCe

```go
partial := components.CreatePartialRedisCe(components.PartialRedisCe{/* values here */})
```

### PartialRedisEe

```go
partial := components.CreatePartialRedisEe(components.PartialRedisEe{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch partial.Type {
	case components.PartialTypeRedisCe:
		// partial.PartialRedisCe is populated
	case components.PartialTypeRedisEe:
		// partial.PartialRedisEe is populated
}
```
