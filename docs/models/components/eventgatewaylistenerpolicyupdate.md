# EventGatewayListenerPolicyUpdate

The typed schema of the listener policy to modify it.


## Supported Types

### EventGatewayTLSListenerSensitiveDataAwarePolicy

```go
eventGatewayListenerPolicyUpdate := components.CreateEventGatewayListenerPolicyUpdateTLSServer(components.EventGatewayTLSListenerSensitiveDataAwarePolicy{/* values here */})
```

### ForwardToVirtualClusterPolicy

```go
eventGatewayListenerPolicyUpdate := components.CreateEventGatewayListenerPolicyUpdateForwardToVirtualCluster(components.ForwardToVirtualClusterPolicy{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch eventGatewayListenerPolicyUpdate.Type {
	case components.EventGatewayListenerPolicyUpdateTypeTLSServer:
		// eventGatewayListenerPolicyUpdate.EventGatewayTLSListenerSensitiveDataAwarePolicy is populated
	case components.EventGatewayListenerPolicyUpdateTypeForwardToVirtualCluster:
		// eventGatewayListenerPolicyUpdate.ForwardToVirtualClusterPolicy is populated
}
```
