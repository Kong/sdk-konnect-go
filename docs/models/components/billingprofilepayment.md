# BillingProfilePayment

The payment settings for this workflow


## Supported Types

### BillingWorkflowPaymentChargeAutomaticallySettings

```go
billingProfilePayment := components.CreateBillingProfilePaymentChargeAutomatically(components.BillingWorkflowPaymentChargeAutomaticallySettings{/* values here */})
```

### BillingWorkflowPaymentSendInvoiceSettings

```go
billingProfilePayment := components.CreateBillingProfilePaymentSendInvoice(components.BillingWorkflowPaymentSendInvoiceSettings{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch billingProfilePayment.Type {
	case components.BillingProfilePaymentTypeChargeAutomatically:
		// billingProfilePayment.BillingWorkflowPaymentChargeAutomaticallySettings is populated
	case components.BillingProfilePaymentTypeSendInvoice:
		// billingProfilePayment.BillingWorkflowPaymentSendInvoiceSettings is populated
}
```
