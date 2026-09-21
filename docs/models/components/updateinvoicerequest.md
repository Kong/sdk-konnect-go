# UpdateInvoiceRequest

UpdateInvoiceRequest update request.


## Supported Types

### UpdateInvoiceStandardRequest

```go
updateInvoiceRequest := components.CreateUpdateInvoiceRequestStandard(components.UpdateInvoiceStandardRequest{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch updateInvoiceRequest.Type {
	case components.UpdateInvoiceRequestTypeStandard:
		// updateInvoiceRequest.UpdateInvoiceStandardRequest is populated
}
```
