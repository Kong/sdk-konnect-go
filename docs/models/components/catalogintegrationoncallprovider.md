# CatalogIntegrationOnCallProvider

Defines how an integration behaves as a source provider of Catalog Service On-Call schedules.
On-Call schedules are entities that can be attached to Catalog Services via an integration's Resource mapping.
When an integration implements this capability, it can act as a source type for On-Call data.
In this role, the integration becomes the source of truth for the on-call schedule data.
When an on-call schedule is attached to a Catalog Service using this source type, the platform relies on the external system to provide and update the on-call schedule data.
A null value indicates the given integration does not act as a source provider of on-call schedules.



## Fields

| Field       | Type        | Required    | Description |
| ----------- | ----------- | ----------- | ----------- |