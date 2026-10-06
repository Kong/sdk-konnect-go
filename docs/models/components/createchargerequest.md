# CreateChargeRequest

Customer charge.


## Supported Types

### CreateChargeFlatFeeRequest

```go
createChargeRequest := components.CreateCreateChargeRequestFlatFee(components.CreateChargeFlatFeeRequest{/* values here */})
```

### CreateChargeUsageBasedRequest

```go
createChargeRequest := components.CreateCreateChargeRequestUsageBased(components.CreateChargeUsageBasedRequest{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createChargeRequest.Type {
	case components.CreateChargeRequestTypeFlatFee:
		// createChargeRequest.CreateChargeFlatFeeRequest is populated
	case components.CreateChargeRequestTypeUsageBased:
		// createChargeRequest.CreateChargeUsageBasedRequest is populated
}
```
