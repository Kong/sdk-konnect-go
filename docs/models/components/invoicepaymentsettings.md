# InvoicePaymentSettings

Payment settings for this invoice.


## Supported Types

### BillingWorkflowPaymentChargeAutomaticallySettings

```go
invoicePaymentSettings := components.CreateInvoicePaymentSettingsChargeAutomatically(components.BillingWorkflowPaymentChargeAutomaticallySettings{/* values here */})
```

### BillingWorkflowPaymentSendInvoiceSettings

```go
invoicePaymentSettings := components.CreateInvoicePaymentSettingsSendInvoice(components.BillingWorkflowPaymentSendInvoiceSettings{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch invoicePaymentSettings.Type {
	case components.InvoicePaymentSettingsTypeChargeAutomatically:
		// invoicePaymentSettings.BillingWorkflowPaymentChargeAutomaticallySettings is populated
	case components.InvoicePaymentSettingsTypeSendInvoice:
		// invoicePaymentSettings.BillingWorkflowPaymentSendInvoiceSettings is populated
}
```
