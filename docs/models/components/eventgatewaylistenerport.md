# EventGatewayListenerPort


## Supported Types

### 

```go
eventGatewayListenerPort := components.CreateEventGatewayListenerPortStr(string{/* values here */})
```

### 

```go
eventGatewayListenerPort := components.CreateEventGatewayListenerPortInteger(int64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch eventGatewayListenerPort.Type {
	case components.EventGatewayListenerPortTypeStr:
		// eventGatewayListenerPort.Str is populated
	case components.EventGatewayListenerPortTypeInteger:
		// eventGatewayListenerPort.Integer is populated
}
```
