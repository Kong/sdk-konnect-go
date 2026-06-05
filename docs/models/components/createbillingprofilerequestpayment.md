# CreateBillingProfileRequestPayment

The payment settings for this workflow


## Supported Types

### BillingWorkflowPaymentChargeAutomaticallySettings

```go
createBillingProfileRequestPayment := components.CreateCreateBillingProfileRequestPaymentChargeAutomatically(components.BillingWorkflowPaymentChargeAutomaticallySettings{/* values here */})
```

### BillingWorkflowPaymentSendInvoiceSettings

```go
createBillingProfileRequestPayment := components.CreateCreateBillingProfileRequestPaymentSendInvoice(components.BillingWorkflowPaymentSendInvoiceSettings{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createBillingProfileRequestPayment.Type {
	case components.CreateBillingProfileRequestPaymentTypeChargeAutomatically:
		// createBillingProfileRequestPayment.BillingWorkflowPaymentChargeAutomaticallySettings is populated
	case components.CreateBillingProfileRequestPaymentTypeSendInvoice:
		// createBillingProfileRequestPayment.BillingWorkflowPaymentSendInvoiceSettings is populated
}
```
