# UpsertBillingProfileRequestAlignment

The alignment for collecting the pending line items into an invoice.


## Supported Types

### UpsertBillingProfileRequestAlignmentBillingWorkflowCollectionAlignmentSubscription

```go
upsertBillingProfileRequestAlignment := components.CreateUpsertBillingProfileRequestAlignmentSubscription(components.UpsertBillingProfileRequestAlignmentBillingWorkflowCollectionAlignmentSubscription{/* values here */})
```

### UpsertBillingProfileRequestAlignmentBillingWorkflowCollectionAlignmentAnchored

```go
upsertBillingProfileRequestAlignment := components.CreateUpsertBillingProfileRequestAlignmentAnchored(components.UpsertBillingProfileRequestAlignmentBillingWorkflowCollectionAlignmentAnchored{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch upsertBillingProfileRequestAlignment.Type {
	case components.UpsertBillingProfileRequestAlignmentTypeSubscription:
		// upsertBillingProfileRequestAlignment.UpsertBillingProfileRequestAlignmentBillingWorkflowCollectionAlignmentSubscription is populated
	case components.UpsertBillingProfileRequestAlignmentTypeAnchored:
		// upsertBillingProfileRequestAlignment.UpsertBillingProfileRequestAlignmentBillingWorkflowCollectionAlignmentAnchored is populated
}
```
