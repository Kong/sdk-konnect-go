# Payment

The payment settings for this workflow


## Supported Types

### BillingWorkflowPaymentChargeAutomaticallySettings

```go
payment := components.CreatePaymentChargeAutomatically(metering.BillingWorkflowPaymentChargeAutomaticallySettings{/* values here */})
```

### BillingWorkflowPaymentSendInvoiceSettings

```go
payment := components.CreatePaymentSendInvoice(metering.BillingWorkflowPaymentSendInvoiceSettings{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch payment.Type {
	case components.PaymentTypeChargeAutomatically:
		// payment.BillingWorkflowPaymentChargeAutomaticallySettings is populated
	case components.PaymentTypeSendInvoice:
		// payment.BillingWorkflowPaymentSendInvoiceSettings is populated
}
```
