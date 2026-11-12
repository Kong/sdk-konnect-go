# PaymentSettings

Payment settings for this invoice.


## Supported Types

### BillingWorkflowPaymentChargeAutomaticallySettings

```go
paymentSettings := components.CreatePaymentSettingsChargeAutomatically(components.BillingWorkflowPaymentChargeAutomaticallySettings{/* values here */})
```

### BillingWorkflowPaymentSendInvoiceSettings

```go
paymentSettings := components.CreatePaymentSettingsSendInvoice(components.BillingWorkflowPaymentSendInvoiceSettings{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch paymentSettings.Type {
	case components.PaymentSettingsTypeChargeAutomatically:
		// paymentSettings.BillingWorkflowPaymentChargeAutomaticallySettings is populated
	case components.PaymentSettingsTypeSendInvoice:
		// paymentSettings.BillingWorkflowPaymentSendInvoiceSettings is populated
}
```
