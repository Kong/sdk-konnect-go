# ActorID


## Supported Types

### StringFieldFilter

```go
actorID := components.CreateActorIDStringFieldFilter(components.StringFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch actorID.Type {
	case components.ActorIDTypeStringFieldFilter:
		// actorID.StringFieldFilter is populated
}
```
