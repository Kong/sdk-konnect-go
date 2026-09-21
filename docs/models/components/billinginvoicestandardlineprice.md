# BillingInvoiceStandardLinePrice

The price definition used to calculate charges for this line.


## Supported Types

### BillingPriceFree

```go
billingInvoiceStandardLinePrice := components.CreateBillingInvoiceStandardLinePriceFree(components.BillingPriceFree{/* values here */})
```

### BillingPriceFlat

```go
billingInvoiceStandardLinePrice := components.CreateBillingInvoiceStandardLinePriceFlat(components.BillingPriceFlat{/* values here */})
```

### BillingPriceUnit

```go
billingInvoiceStandardLinePrice := components.CreateBillingInvoiceStandardLinePriceUnit(components.BillingPriceUnit{/* values here */})
```

### BillingPriceGraduated

```go
billingInvoiceStandardLinePrice := components.CreateBillingInvoiceStandardLinePriceGraduated(components.BillingPriceGraduated{/* values here */})
```

### BillingPriceVolume

```go
billingInvoiceStandardLinePrice := components.CreateBillingInvoiceStandardLinePriceVolume(components.BillingPriceVolume{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch billingInvoiceStandardLinePrice.Type {
	case components.BillingInvoiceStandardLinePriceTypeFree:
		// billingInvoiceStandardLinePrice.BillingPriceFree is populated
	case components.BillingInvoiceStandardLinePriceTypeFlat:
		// billingInvoiceStandardLinePrice.BillingPriceFlat is populated
	case components.BillingInvoiceStandardLinePriceTypeUnit:
		// billingInvoiceStandardLinePrice.BillingPriceUnit is populated
	case components.BillingInvoiceStandardLinePriceTypeGraduated:
		// billingInvoiceStandardLinePrice.BillingPriceGraduated is populated
	case components.BillingInvoiceStandardLinePriceTypeVolume:
		// billingInvoiceStandardLinePrice.BillingPriceVolume is populated
}
```
