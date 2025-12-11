# CatalogIntegrationWorkflowProvider

Defines how an integration behaves as a source provider of Catalog Service Workflows.
Workflows are entities that can be attached to Catalog Services via an integration's Resource mapping.
When an integration implements this capability, it can act as a source type for Workflow data.
In this role, the integration becomes the source of truth for the workflow data.
When a workflow is attached to a Catalog Service using this source type, the platform relies on the external system to provide and update the workflow data.
A null value indicates the given integration does not act as a source provider of workflows.



## Fields

| Field       | Type        | Required    | Description |
| ----------- | ----------- | ----------- | ----------- |