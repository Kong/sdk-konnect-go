# GitHubAppInstallationAuth

Defines the GitHub App authorization strategy used by the GitHub integration.
This strategy enables secure access to GitHub APIs using app installation tokens.
It supports both API-based interactions and real-time event delivery via GitHub webhooks.
Unlike standard OAuth flows, this strategy leverages GitHub's custom app installation flow
and token lifecycle, making it ideal for deep, organization-level GitHub integration.



## Fields

| Field                                                                                                    | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `Type`                                                                                                   | *string*                                                                                                 | :heavy_check_mark:                                                                                       | N/A                                                                                                      |
| `Config`                                                                                                 | [components.GitHubAppInstallationAuthConfig](../../models/components/githubappinstallationauthconfig.md) | :heavy_check_mark:                                                                                       | N/A                                                                                                      |