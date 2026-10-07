# BillingInstallAppRequest

Request to install an app from the catalog.


## Supported Types

### BillingInstallAppStripeWithAPIKey

```go
billingInstallAppRequest := components.CreateBillingInstallAppRequestStripe(components.BillingInstallAppStripeWithAPIKey{/* values here */})
```

### BillingInstallAppSandbox

```go
billingInstallAppRequest := components.CreateBillingInstallAppRequestSandbox(components.BillingInstallAppSandbox{/* values here */})
```

### BillingInstallAppExternalInvoicing

```go
billingInstallAppRequest := components.CreateBillingInstallAppRequestExternalInvoicing(components.BillingInstallAppExternalInvoicing{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch billingInstallAppRequest.Type {
	case components.BillingInstallAppRequestTypeStripe:
		// billingInstallAppRequest.BillingInstallAppStripeWithAPIKey is populated
	case components.BillingInstallAppRequestTypeSandbox:
		// billingInstallAppRequest.BillingInstallAppSandbox is populated
	case components.BillingInstallAppRequestTypeExternalInvoicing:
		// billingInstallAppRequest.BillingInstallAppExternalInvoicing is populated
}
```
