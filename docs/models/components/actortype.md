# ActorType


## Supported Types

### StringFieldFilter

```go
actorType := components.CreateActorTypeStringFieldFilter(components.StringFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch actorType.Type {
	case components.ActorTypeTypeStringFieldFilter:
		// actorType.StringFieldFilter is populated
}
```
