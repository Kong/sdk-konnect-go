# Partial


## Supported Types

### PartialEmbeddings

```go
partial := components.CreatePartialEmbeddings(components.PartialEmbeddings{/* values here */})
```

### PartialModel

```go
partial := components.CreatePartialModel(components.PartialModel{/* values here */})
```

### PartialRedisCe

```go
partial := components.CreatePartialRedisCe(components.PartialRedisCe{/* values here */})
```

### PartialRedisEe

```go
partial := components.CreatePartialRedisEe(components.PartialRedisEe{/* values here */})
```

### PartialVectordb

```go
partial := components.CreatePartialVectordb(components.PartialVectordb{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch partial.Type {
	case components.PartialTypeEmbeddings:
		// partial.PartialEmbeddings is populated
	case components.PartialTypeModel:
		// partial.PartialModel is populated
	case components.PartialTypeRedisCe:
		// partial.PartialRedisCe is populated
	case components.PartialTypeRedisEe:
		// partial.PartialRedisEe is populated
	case components.PartialTypeVectordb:
		// partial.PartialVectordb is populated
}
```
