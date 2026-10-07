# CreateEntitlementRequest

Entitlement create request.


## Supported Types

### CreateEntitlementMeteredRequest

```go
createEntitlementRequest := components.CreateCreateEntitlementRequestMetered(components.CreateEntitlementMeteredRequest{/* values here */})
```

### CreateEntitlementStaticRequest

```go
createEntitlementRequest := components.CreateCreateEntitlementRequestStatic(components.CreateEntitlementStaticRequest{/* values here */})
```

### CreateEntitlementBooleanRequest

```go
createEntitlementRequest := components.CreateCreateEntitlementRequestBoolean(components.CreateEntitlementBooleanRequest{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createEntitlementRequest.Type {
	case components.CreateEntitlementRequestTypeMetered:
		// createEntitlementRequest.CreateEntitlementMeteredRequest is populated
	case components.CreateEntitlementRequestTypeStatic:
		// createEntitlementRequest.CreateEntitlementStaticRequest is populated
	case components.CreateEntitlementRequestTypeBoolean:
		// createEntitlementRequest.CreateEntitlementBooleanRequest is populated
}
```
