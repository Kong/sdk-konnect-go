# UpdateInvoiceStandardLinePrice

The price definition used to calculate charges for this line.


## Supported Types

### UpdatePriceFree

```go
updateInvoiceStandardLinePrice := components.CreateUpdateInvoiceStandardLinePriceFree(components.UpdatePriceFree{/* values here */})
```

### UpdatePriceFlat

```go
updateInvoiceStandardLinePrice := components.CreateUpdateInvoiceStandardLinePriceFlat(components.UpdatePriceFlat{/* values here */})
```

### UpdatePriceUnit

```go
updateInvoiceStandardLinePrice := components.CreateUpdateInvoiceStandardLinePriceUnit(components.UpdatePriceUnit{/* values here */})
```

### UpdatePriceGraduated

```go
updateInvoiceStandardLinePrice := components.CreateUpdateInvoiceStandardLinePriceGraduated(components.UpdatePriceGraduated{/* values here */})
```

### UpdatePriceVolume

```go
updateInvoiceStandardLinePrice := components.CreateUpdateInvoiceStandardLinePriceVolume(components.UpdatePriceVolume{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch updateInvoiceStandardLinePrice.Type {
	case components.UpdateInvoiceStandardLinePriceTypeFree:
		// updateInvoiceStandardLinePrice.UpdatePriceFree is populated
	case components.UpdateInvoiceStandardLinePriceTypeFlat:
		// updateInvoiceStandardLinePrice.UpdatePriceFlat is populated
	case components.UpdateInvoiceStandardLinePriceTypeUnit:
		// updateInvoiceStandardLinePrice.UpdatePriceUnit is populated
	case components.UpdateInvoiceStandardLinePriceTypeGraduated:
		// updateInvoiceStandardLinePrice.UpdatePriceGraduated is populated
	case components.UpdateInvoiceStandardLinePriceTypeVolume:
		// updateInvoiceStandardLinePrice.UpdatePriceVolume is populated
}
```
