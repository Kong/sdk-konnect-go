# UpdateInvoiceStandardRequestPaymentSettings

Payment settings for this invoice.


## Supported Types

### UpdateBillingWorkflowPaymentChargeAutomaticallySettings

```go
updateInvoiceStandardRequestPaymentSettings := components.CreateUpdateInvoiceStandardRequestPaymentSettingsChargeAutomatically(components.UpdateBillingWorkflowPaymentChargeAutomaticallySettings{/* values here */})
```

### UpdateBillingWorkflowPaymentSendInvoiceSettings

```go
updateInvoiceStandardRequestPaymentSettings := components.CreateUpdateInvoiceStandardRequestPaymentSettingsSendInvoice(components.UpdateBillingWorkflowPaymentSendInvoiceSettings{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch updateInvoiceStandardRequestPaymentSettings.Type {
	case components.UpdateInvoiceStandardRequestPaymentSettingsTypeChargeAutomatically:
		// updateInvoiceStandardRequestPaymentSettings.UpdateBillingWorkflowPaymentChargeAutomaticallySettings is populated
	case components.UpdateInvoiceStandardRequestPaymentSettingsTypeSendInvoice:
		// updateInvoiceStandardRequestPaymentSettings.UpdateBillingWorkflowPaymentSendInvoiceSettings is populated
}
```
