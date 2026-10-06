# BillingInstallAppResponse

Response of the app install.


## Supported Types

### BillingInstalledAppStripe

```go
billingInstallAppResponse := components.CreateBillingInstallAppResponseStripe(components.BillingInstalledAppStripe{/* values here */})
```

### BillingInstalledAppSandbox

```go
billingInstallAppResponse := components.CreateBillingInstallAppResponseSandbox(components.BillingInstalledAppSandbox{/* values here */})
```

### BillingInstalledAppExternalInvoicing

```go
billingInstallAppResponse := components.CreateBillingInstallAppResponseExternalInvoicing(components.BillingInstalledAppExternalInvoicing{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch billingInstallAppResponse.Type {
	case components.BillingInstallAppResponseTypeStripe:
		// billingInstallAppResponse.BillingInstalledAppStripe is populated
	case components.BillingInstallAppResponseTypeSandbox:
		// billingInstallAppResponse.BillingInstalledAppSandbox is populated
	case components.BillingInstallAppResponseTypeExternalInvoicing:
		// billingInstallAppResponse.BillingInstalledAppExternalInvoicing is populated
}
```
