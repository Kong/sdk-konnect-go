# EventGatewayMaskEmailDomainStrategy

The strategy used to redact the domain of an email address.


## Supported Types

### EventGatewayMaskStrategyKeepChars

```go
eventGatewayMaskEmailDomainStrategy := components.CreateEventGatewayMaskEmailDomainStrategyKeepChars(components.EventGatewayMaskStrategyKeepChars{/* values here */})
```

### EventGatewayMaskEmailDomainKeepAll

```go
eventGatewayMaskEmailDomainStrategy := components.CreateEventGatewayMaskEmailDomainStrategyKeepAll(components.EventGatewayMaskEmailDomainKeepAll{/* values here */})
```

### EventGatewayMaskStrategyReplace

```go
eventGatewayMaskEmailDomainStrategy := components.CreateEventGatewayMaskEmailDomainStrategyReplace(components.EventGatewayMaskStrategyReplace{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch eventGatewayMaskEmailDomainStrategy.Type {
	case components.EventGatewayMaskEmailDomainStrategyTypeKeepChars:
		// eventGatewayMaskEmailDomainStrategy.EventGatewayMaskStrategyKeepChars is populated
	case components.EventGatewayMaskEmailDomainStrategyTypeKeepAll:
		// eventGatewayMaskEmailDomainStrategy.EventGatewayMaskEmailDomainKeepAll is populated
	case components.EventGatewayMaskEmailDomainStrategyTypeReplace:
		// eventGatewayMaskEmailDomainStrategy.EventGatewayMaskStrategyReplace is populated
}
```
