# UpsertBillingProfileRequestWorkflowInvoiceSettings

The invoicing settings for this workflow


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  | Example                                                                      |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `AutoAdvance`                                                                | **bool*                                                                      | :heavy_minus_sign:                                                           | Whether to automatically issue the invoice after the draftPeriod has passed. |                                                                              |
| `DraftPeriod`                                                                | **string*                                                                    | :heavy_minus_sign:                                                           | The period for the invoice to be kept in draft status for manual reviews.    | P1D                                                                          |
| `ProgressiveBilling`                                                         | **bool*                                                                      | :heavy_minus_sign:                                                           | Should progressive billing be allowed for this workflow?                     |                                                                              |