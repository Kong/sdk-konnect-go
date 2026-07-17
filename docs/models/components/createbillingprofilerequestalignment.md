# CreateBillingProfileRequestAlignment

The alignment for collecting the pending line items into an invoice.


## Supported Types

### BillingWorkflowCollectionAlignmentSubscription

```go
createBillingProfileRequestAlignment := components.CreateCreateBillingProfileRequestAlignmentSubscription(components.BillingWorkflowCollectionAlignmentSubscription{/* values here */})
```

### BillingWorkflowCollectionAlignmentAnchored

```go
createBillingProfileRequestAlignment := components.CreateCreateBillingProfileRequestAlignmentAnchored(components.BillingWorkflowCollectionAlignmentAnchored{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createBillingProfileRequestAlignment.Type {
	case components.CreateBillingProfileRequestAlignmentTypeSubscription:
		// createBillingProfileRequestAlignment.BillingWorkflowCollectionAlignmentSubscription is populated
	case components.CreateBillingProfileRequestAlignmentTypeAnchored:
		// createBillingProfileRequestAlignment.BillingWorkflowCollectionAlignmentAnchored is populated
}
```
