# Alignment

The alignment for collecting the pending line items into an invoice.


## Supported Types

### BillingWorkflowCollectionAlignmentSubscription

```go
alignment := components.CreateAlignmentSubscription(components.BillingWorkflowCollectionAlignmentSubscription{/* values here */})
```

### BillingWorkflowCollectionAlignmentAnchored

```go
alignment := components.CreateAlignmentAnchored(components.BillingWorkflowCollectionAlignmentAnchored{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch alignment.Type {
	case components.AlignmentUnionTypeSubscription:
		// alignment.BillingWorkflowCollectionAlignmentSubscription is populated
	case components.AlignmentUnionTypeAnchored:
		// alignment.BillingWorkflowCollectionAlignmentAnchored is populated
}
```
