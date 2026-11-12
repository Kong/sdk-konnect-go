# EventGatewayMaskStrategy

The strategy used to redact a matched field value.


## Supported Types

### EventGatewayMaskStrategyKeepChars

```go
eventGatewayMaskStrategy := components.CreateEventGatewayMaskStrategyKeepChars(components.EventGatewayMaskStrategyKeepChars{/* values here */})
```

### EventGatewayMaskStrategyEmail

```go
eventGatewayMaskStrategy := components.CreateEventGatewayMaskStrategyEmail(components.EventGatewayMaskStrategyEmail{/* values here */})
```

### EventGatewayMaskStrategyReplace

```go
eventGatewayMaskStrategy := components.CreateEventGatewayMaskStrategyReplace(components.EventGatewayMaskStrategyReplace{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch eventGatewayMaskStrategy.Type {
	case components.EventGatewayMaskStrategyTypeKeepChars:
		// eventGatewayMaskStrategy.EventGatewayMaskStrategyKeepChars is populated
	case components.EventGatewayMaskStrategyTypeEmail:
		// eventGatewayMaskStrategy.EventGatewayMaskStrategyEmail is populated
	case components.EventGatewayMaskStrategyTypeReplace:
		// eventGatewayMaskStrategy.EventGatewayMaskStrategyReplace is populated
}
```
