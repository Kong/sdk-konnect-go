# UpsertBillingProfileRequestPayment

The payment settings for this workflow


## Supported Types

### BillingWorkflowPaymentChargeAutomaticallySettings

```go
upsertBillingProfileRequestPayment := components.CreateUpsertBillingProfileRequestPaymentChargeAutomatically(metering.BillingWorkflowPaymentChargeAutomaticallySettings{/* values here */})
```

### BillingWorkflowPaymentSendInvoiceSettings

```go
upsertBillingProfileRequestPayment := components.CreateUpsertBillingProfileRequestPaymentSendInvoice(metering.BillingWorkflowPaymentSendInvoiceSettings{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch upsertBillingProfileRequestPayment.Type {
	case components.UpsertBillingProfileRequestPaymentTypeChargeAutomatically:
		// upsertBillingProfileRequestPayment.BillingWorkflowPaymentChargeAutomaticallySettings is populated
	case components.UpsertBillingProfileRequestPaymentTypeSendInvoice:
		// upsertBillingProfileRequestPayment.BillingWorkflowPaymentSendInvoiceSettings is populated
}
```
