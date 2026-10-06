# InvoiceApplicationThatGeneratedTheInvoice

The apps that will be used to orchestrate the invoice's workflow.


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `Tax`                                                                      | [components.InvoiceTax](../../models/components/invoicetax.md)             | :heavy_check_mark:                                                         | The tax app used for this workflow                                         |
| `Invoicing`                                                                | [components.InvoiceInvoicing](../../models/components/invoiceinvoicing.md) | :heavy_check_mark:                                                         | The invoicing app used for this workflow                                   |
| `Payment`                                                                  | [components.InvoicePayment](../../models/components/invoicepayment.md)     | :heavy_check_mark:                                                         | The payment app used for this workflow                                     |