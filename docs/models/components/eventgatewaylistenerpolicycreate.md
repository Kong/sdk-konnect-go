# EventGatewayListenerPolicyCreate

The typed schema of the listener policy to modify it.


## Supported Types

### EventGatewayTLSListenerPolicy

```go
eventGatewayListenerPolicyCreate := components.CreateEventGatewayListenerPolicyCreateTLSServer(components.EventGatewayTLSListenerPolicy{/* values here */})
```

### ForwardToVirtualClusterPolicy

```go
eventGatewayListenerPolicyCreate := components.CreateEventGatewayListenerPolicyCreateForwardToVirtualCluster(components.ForwardToVirtualClusterPolicy{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch eventGatewayListenerPolicyCreate.Type {
	case components.EventGatewayListenerPolicyCreateTypeTLSServer:
		// eventGatewayListenerPolicyCreate.EventGatewayTLSListenerPolicy is populated
	case components.EventGatewayListenerPolicyCreateTypeForwardToVirtualCluster:
		// eventGatewayListenerPolicyCreate.ForwardToVirtualClusterPolicy is populated
}
```
