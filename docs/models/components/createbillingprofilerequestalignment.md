# CreateBillingProfileRequestAlignment

The alignment for collecting the pending line items into an invoice.


## Supported Types

### AlignmentBillingWorkflowCollectionAlignmentSubscription

```go
createBillingProfileRequestAlignment := components.CreateCreateBillingProfileRequestAlignmentSubscription(components.AlignmentBillingWorkflowCollectionAlignmentSubscription{/* values here */})
```

### AlignmentBillingWorkflowCollectionAlignmentAnchored

```go
createBillingProfileRequestAlignment := components.CreateCreateBillingProfileRequestAlignmentAnchored(components.AlignmentBillingWorkflowCollectionAlignmentAnchored{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createBillingProfileRequestAlignment.Type {
	case components.CreateBillingProfileRequestAlignmentTypeSubscription:
		// createBillingProfileRequestAlignment.AlignmentBillingWorkflowCollectionAlignmentSubscription is populated
	case components.CreateBillingProfileRequestAlignmentTypeAnchored:
		// createBillingProfileRequestAlignment.AlignmentBillingWorkflowCollectionAlignmentAnchored is populated
}
```
