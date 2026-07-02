# PaymentSettings

Payment settings for this invoice.


## Supported Types

### PaymentSettingsBillingWorkflowPaymentChargeAutomaticallySettings

```go
paymentSettings := components.CreatePaymentSettingsChargeAutomatically(components.PaymentSettingsBillingWorkflowPaymentChargeAutomaticallySettings{/* values here */})
```

### PaymentSettingsBillingWorkflowPaymentSendInvoiceSettings

```go
paymentSettings := components.CreatePaymentSettingsSendInvoice(components.PaymentSettingsBillingWorkflowPaymentSendInvoiceSettings{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch paymentSettings.Type {
	case components.PaymentSettingsTypeChargeAutomatically:
		// paymentSettings.PaymentSettingsBillingWorkflowPaymentChargeAutomaticallySettings is populated
	case components.PaymentSettingsTypeSendInvoice:
		// paymentSettings.PaymentSettingsBillingWorkflowPaymentSendInvoiceSettings is populated
}
```
