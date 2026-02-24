# BillingCustomerStripeCreateCustomerPortalSessionRequest

Request to create a Stripe Customer Portal Session for the customer.

Useful to redirect the customer to the Stripe Customer Portal to manage their payment methods,
change their billing address and access their invoice history.
Only returns URL if the customer billing profile is linked to a stripe app and customer.


## Fields

| Field                                                                                                                                                                              | Type                                                                                                                                                                               | Required                                                                                                                                                                           | Description                                                                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `StripeOptions`                                                                                                                                                                    | [components.BillingCustomerStripeCreateCustomerPortalSessionRequestStripeOptions](../../models/components/billingcustomerstripecreatecustomerportalsessionrequeststripeoptions.md) | :heavy_check_mark:                                                                                                                                                                 | Options for configuring the Stripe Customer Portal Session.                                                                                                                        |