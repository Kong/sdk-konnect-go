# BillingApp

Installed application.


## Supported Types

### BillingAppStripe

```go
billingApp := components.CreateBillingAppStripe(components.BillingAppStripe{/* values here */})
```

### BillingAppSandbox

```go
billingApp := components.CreateBillingAppSandbox(components.BillingAppSandbox{/* values here */})
```

### BillingAppExternalInvoicing

```go
billingApp := components.CreateBillingAppExternalInvoicing(components.BillingAppExternalInvoicing{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch billingApp.Type {
	case components.BillingAppUnionTypeStripe:
		// billingApp.BillingAppStripe is populated
	case components.BillingAppUnionTypeSandbox:
		// billingApp.BillingAppSandbox is populated
	case components.BillingAppUnionTypeExternalInvoicing:
		// billingApp.BillingAppExternalInvoicing is populated
}
```
