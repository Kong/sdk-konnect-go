# UpsertBillingProfileRequestAlignment

The alignment for collecting the pending line items into an invoice.


## Supported Types

### BillingWorkflowCollectionAlignmentSubscription

```go
upsertBillingProfileRequestAlignment := components.CreateUpsertBillingProfileRequestAlignmentSubscription(components.BillingWorkflowCollectionAlignmentSubscription{/* values here */})
```

### BillingWorkflowCollectionAlignmentAnchored

```go
upsertBillingProfileRequestAlignment := components.CreateUpsertBillingProfileRequestAlignmentAnchored(components.BillingWorkflowCollectionAlignmentAnchored{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch upsertBillingProfileRequestAlignment.Type {
	case components.UpsertBillingProfileRequestAlignmentTypeSubscription:
		// upsertBillingProfileRequestAlignment.BillingWorkflowCollectionAlignmentSubscription is populated
	case components.UpsertBillingProfileRequestAlignmentTypeAnchored:
		// upsertBillingProfileRequestAlignment.BillingWorkflowCollectionAlignmentAnchored is populated
}
```
