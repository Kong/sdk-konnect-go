# ActorName


## Supported Types

### StringFieldFilter

```go
actorName := components.CreateActorNameStringFieldFilter(components.StringFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch actorName.Type {
	case components.ActorNameTypeStringFieldFilter:
		// actorName.StringFieldFilter is populated
}
```
