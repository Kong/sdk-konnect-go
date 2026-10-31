# BillingSubscriptionEditOperation

A single customization to apply to a running subscription. The `type` field
discriminates which operation is performed.


## Supported Types

### BillingSubscriptionEditAddItem

```go
billingSubscriptionEditOperation := components.CreateBillingSubscriptionEditOperationAddItem(components.BillingSubscriptionEditAddItem{/* values here */})
```

### BillingSubscriptionEditRemoveItem

```go
billingSubscriptionEditOperation := components.CreateBillingSubscriptionEditOperationRemoveItem(components.BillingSubscriptionEditRemoveItem{/* values here */})
```

### BillingSubscriptionEditAddPhase

```go
billingSubscriptionEditOperation := components.CreateBillingSubscriptionEditOperationAddPhase(components.BillingSubscriptionEditAddPhase{/* values here */})
```

### BillingSubscriptionEditRemovePhase

```go
billingSubscriptionEditOperation := components.CreateBillingSubscriptionEditOperationRemovePhase(components.BillingSubscriptionEditRemovePhase{/* values here */})
```

### BillingSubscriptionEditStretchPhase

```go
billingSubscriptionEditOperation := components.CreateBillingSubscriptionEditOperationStretchPhase(components.BillingSubscriptionEditStretchPhase{/* values here */})
```

### BillingSubscriptionEditUnscheduleEdit

```go
billingSubscriptionEditOperation := components.CreateBillingSubscriptionEditOperationUnscheduleEdit(components.BillingSubscriptionEditUnscheduleEdit{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch billingSubscriptionEditOperation.Type {
	case components.BillingSubscriptionEditOperationTypeAddItem:
		// billingSubscriptionEditOperation.BillingSubscriptionEditAddItem is populated
	case components.BillingSubscriptionEditOperationTypeRemoveItem:
		// billingSubscriptionEditOperation.BillingSubscriptionEditRemoveItem is populated
	case components.BillingSubscriptionEditOperationTypeAddPhase:
		// billingSubscriptionEditOperation.BillingSubscriptionEditAddPhase is populated
	case components.BillingSubscriptionEditOperationTypeRemovePhase:
		// billingSubscriptionEditOperation.BillingSubscriptionEditRemovePhase is populated
	case components.BillingSubscriptionEditOperationTypeStretchPhase:
		// billingSubscriptionEditOperation.BillingSubscriptionEditStretchPhase is populated
	case components.BillingSubscriptionEditOperationTypeUnscheduleEdit:
		// billingSubscriptionEditOperation.BillingSubscriptionEditUnscheduleEdit is populated
}
```
