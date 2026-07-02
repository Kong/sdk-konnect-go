# InvoicingSettings

Invoicing settings for this invoice.


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   | Example                                                                       |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `AutoAdvance`                                                                 | `*bool`                                                                       | :heavy_minus_sign:                                                            | Whether to automatically issue the invoice after the draft_period has passed. |                                                                               |
| `DraftPeriod`                                                                 | `*string`                                                                     | :heavy_minus_sign:                                                            | The period for the invoice to be kept in draft status for manual reviews.     | P1D                                                                           |