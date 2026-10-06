# BillingUpdateAppRequest

Request to update an installed app.


## Supported Types

### UpdateAppStripeRequest

```go
billingUpdateAppRequest := components.CreateBillingUpdateAppRequestStripe(components.UpdateAppStripeRequest{/* values here */})
```

### UpdateAppSandboxRequest

```go
billingUpdateAppRequest := components.CreateBillingUpdateAppRequestSandbox(components.UpdateAppSandboxRequest{/* values here */})
```

### UpdateAppExternalInvoicingRequest

```go
billingUpdateAppRequest := components.CreateBillingUpdateAppRequestExternalInvoicing(components.UpdateAppExternalInvoicingRequest{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch billingUpdateAppRequest.Type {
	case components.BillingUpdateAppRequestTypeStripe:
		// billingUpdateAppRequest.UpdateAppStripeRequest is populated
	case components.BillingUpdateAppRequestTypeSandbox:
		// billingUpdateAppRequest.UpdateAppSandboxRequest is populated
	case components.BillingUpdateAppRequestTypeExternalInvoicing:
		// billingUpdateAppRequest.UpdateAppExternalInvoicingRequest is populated
}
```
