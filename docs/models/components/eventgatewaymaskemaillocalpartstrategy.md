# EventGatewayMaskEmailLocalPartStrategy

The strategy used to redact the local part of an email address.


## Supported Types

### EventGatewayMaskStrategyKeepChars

```go
eventGatewayMaskEmailLocalPartStrategy := components.CreateEventGatewayMaskEmailLocalPartStrategyKeepChars(components.EventGatewayMaskStrategyKeepChars{/* values here */})
```

### EventGatewayMaskStrategyReplace

```go
eventGatewayMaskEmailLocalPartStrategy := components.CreateEventGatewayMaskEmailLocalPartStrategyReplace(components.EventGatewayMaskStrategyReplace{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch eventGatewayMaskEmailLocalPartStrategy.Type {
	case components.EventGatewayMaskEmailLocalPartStrategyTypeKeepChars:
		// eventGatewayMaskEmailLocalPartStrategy.EventGatewayMaskStrategyKeepChars is populated
	case components.EventGatewayMaskEmailLocalPartStrategyTypeReplace:
		// eventGatewayMaskEmailLocalPartStrategy.EventGatewayMaskStrategyReplace is populated
}
```
