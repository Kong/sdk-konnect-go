# EventGatewayClusterPolicyModify

The typed schema of the cluster policy to modify it.


## Supported Types

### EventGatewayACLsPolicy

```go
eventGatewayClusterPolicyModify := components.CreateEventGatewayClusterPolicyModifyAcls(components.EventGatewayACLsPolicy{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch eventGatewayClusterPolicyModify.Type {
	case components.EventGatewayClusterPolicyModifyTypeAcls:
		// eventGatewayClusterPolicyModify.EventGatewayACLsPolicy is populated
}
```
