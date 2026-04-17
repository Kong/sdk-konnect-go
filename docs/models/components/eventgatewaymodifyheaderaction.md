# EventGatewayModifyHeaderAction

An action that modifies a header.


## Supported Types

### EventGatewayModifyHeaderRemoveAction

```go
eventGatewayModifyHeaderAction := components.CreateEventGatewayModifyHeaderActionRemove(components.EventGatewayModifyHeaderRemoveAction{/* values here */})
```

### EventGatewayModifyHeaderSetAction

```go
eventGatewayModifyHeaderAction := components.CreateEventGatewayModifyHeaderActionSet(components.EventGatewayModifyHeaderSetAction{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch eventGatewayModifyHeaderAction.Type {
	case components.EventGatewayModifyHeaderActionTypeRemove:
		// eventGatewayModifyHeaderAction.EventGatewayModifyHeaderRemoveAction is populated
	case components.EventGatewayModifyHeaderActionTypeSet:
		// eventGatewayModifyHeaderAction.EventGatewayModifyHeaderSetAction is populated
}
```
