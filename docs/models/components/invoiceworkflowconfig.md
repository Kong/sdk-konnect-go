# InvoiceWorkflowConfig

The workflow configuration that was active when the invoice was created.

Only the fields that are meaningful at the per-invoice level are included:
invoicing behaviour (auto-advance, draft period) and payment settings
(collection method, due date). Profile-wide settings such as collection
alignment, progressive billing, and tax policy are omitted.


## Fields

| Field                                                                                       | Type                                                                                        | Required                                                                                    | Description                                                                                 |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `Invoicing`                                                                                 | [*components.InvoiceInvoicingSettings](../../models/components/invoiceinvoicingsettings.md) | :heavy_minus_sign:                                                                          | Invoicing settings for this invoice.                                                        |
| `Payment`                                                                                   | [*components.InvoicePaymentSettings](../../models/components/invoicepaymentsettings.md)     | :heavy_minus_sign:                                                                          | Payment settings for this invoice.                                                          |