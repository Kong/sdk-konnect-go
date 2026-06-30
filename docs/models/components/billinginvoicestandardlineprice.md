# BillingInvoiceStandardLinePrice

The price definition used to calculate charges for this line.


## Supported Types

### BillingInvoiceStandardLinePriceBillingPriceFree

```go
billingInvoiceStandardLinePrice := components.CreateBillingInvoiceStandardLinePriceFree(components.BillingInvoiceStandardLinePriceBillingPriceFree{/* values here */})
```

### BillingInvoiceStandardLinePriceBillingPriceFlat

```go
billingInvoiceStandardLinePrice := components.CreateBillingInvoiceStandardLinePriceFlat(components.BillingInvoiceStandardLinePriceBillingPriceFlat{/* values here */})
```

### BillingInvoiceStandardLinePriceBillingPriceUnit

```go
billingInvoiceStandardLinePrice := components.CreateBillingInvoiceStandardLinePriceUnit(components.BillingInvoiceStandardLinePriceBillingPriceUnit{/* values here */})
```

### BillingInvoiceStandardLinePriceBillingPriceGraduated

```go
billingInvoiceStandardLinePrice := components.CreateBillingInvoiceStandardLinePriceGraduated(components.BillingInvoiceStandardLinePriceBillingPriceGraduated{/* values here */})
```

### BillingInvoiceStandardLinePriceBillingPriceVolume

```go
billingInvoiceStandardLinePrice := components.CreateBillingInvoiceStandardLinePriceVolume(components.BillingInvoiceStandardLinePriceBillingPriceVolume{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch billingInvoiceStandardLinePrice.Type {
	case components.BillingInvoiceStandardLinePriceTypeFree:
		// billingInvoiceStandardLinePrice.BillingInvoiceStandardLinePriceBillingPriceFree is populated
	case components.BillingInvoiceStandardLinePriceTypeFlat:
		// billingInvoiceStandardLinePrice.BillingInvoiceStandardLinePriceBillingPriceFlat is populated
	case components.BillingInvoiceStandardLinePriceTypeUnit:
		// billingInvoiceStandardLinePrice.BillingInvoiceStandardLinePriceBillingPriceUnit is populated
	case components.BillingInvoiceStandardLinePriceTypeGraduated:
		// billingInvoiceStandardLinePrice.BillingInvoiceStandardLinePriceBillingPriceGraduated is populated
	case components.BillingInvoiceStandardLinePriceTypeVolume:
		// billingInvoiceStandardLinePrice.BillingInvoiceStandardLinePriceBillingPriceVolume is populated
}
```
