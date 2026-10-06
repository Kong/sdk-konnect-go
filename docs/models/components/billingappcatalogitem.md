# BillingAppCatalogItem

Available apps for billing integrations to connect with third-party services.
Apps can have various capabilities like syncing data from or to external
systems, integrating with third-party services for tax calculation, delivery of
invoices, collection of payments, etc.


## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `Type`                                                                                       | [components.BillingAppCatalogItemType](../../models/components/billingappcatalogitemtype.md) | :heavy_check_mark:                                                                           | Type of the app.                                                                             |
| `Name`                                                                                       | `string`                                                                                     | :heavy_check_mark:                                                                           | Name of the app.                                                                             |
| `Description`                                                                                | `string`                                                                                     | :heavy_check_mark:                                                                           | Description of the app.                                                                      |
| `Capabilities`                                                                               | [][components.BillingAppCapability](../../models/components/billingappcapability.md)         | :heavy_check_mark:                                                                           | Capabilities of the app.                                                                     |
| `InstallMethods`                                                                             | [][components.BillingAppInstallMethods](../../models/components/billingappinstallmethods.md) | :heavy_check_mark:                                                                           | Available install methods of the app.                                                        |