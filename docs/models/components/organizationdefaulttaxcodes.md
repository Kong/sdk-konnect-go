# OrganizationDefaultTaxCodes

Organization-level default tax code references.

Stores the default tax codes applied to specific billing contexts for this
organization. Provisioned automatically when the organization is created.


## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    | Example                                                                        |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `InvoicingTaxCode`                                                             | [components.InvoicingTaxCode](../../models/components/invoicingtaxcode.md)     | :heavy_check_mark:                                                             | Default tax code for invoicing.                                                |                                                                                |
| `CreditGrantTaxCode`                                                           | [components.CreditGrantTaxCode](../../models/components/creditgranttaxcode.md) | :heavy_check_mark:                                                             | Default tax code for credit grants.                                            |                                                                                |
| `CreatedAt`                                                                    | [time.Time](https://pkg.go.dev/time#Time)                                      | :heavy_check_mark:                                                             | Timestamp of creation.                                                         | 2023-01-01T01:01:01.001Z                                                       |
| `UpdatedAt`                                                                    | [time.Time](https://pkg.go.dev/time#Time)                                      | :heavy_check_mark:                                                             | Timestamp of last update.                                                      | 2023-01-01T01:01:01.001Z                                                       |