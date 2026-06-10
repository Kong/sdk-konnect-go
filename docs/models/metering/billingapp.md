# BillingApp

Installed application.


## Supported Types

### BillingAppStripe

```go
billingApp := components.CreateBillingAppStripe(metering.BillingAppStripe{/* values here */})
```

### BillingAppSandbox

```go
billingApp := components.CreateBillingAppSandbox(metering.BillingAppSandbox{/* values here */})
```

### BillingAppExternalInvoicing

```go
billingApp := components.CreateBillingAppExternalInvoicing(metering.BillingAppExternalInvoicing{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch billingApp.Type {
	case components.BillingAppTypeStripe:
		// billingApp.BillingAppStripe is populated
	case components.BillingAppTypeSandbox:
		// billingApp.BillingAppSandbox is populated
	case components.BillingAppTypeExternalInvoicing:
		// billingApp.BillingAppExternalInvoicing is populated
}
```
