# Rank


## Supported Types

### RankAfterPayload

```go
rank := components.CreateRankRankAfterPayload(components.RankAfterPayload{/* values here */})
```

### RankBeforePayload

```go
rank := components.CreateRankRankBeforePayload(components.RankBeforePayload{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch rank.Type {
	case components.RankTypeRankAfterPayload:
		// rank.RankAfterPayload is populated
	case components.RankTypeRankBeforePayload:
		// rank.RankBeforePayload is populated
}
```
