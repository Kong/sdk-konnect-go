package hooks

import (
	"net/http"
	"strings"
)

type GlobalAPIURLRequestHook struct{}

var _ beforeRequestHook = (*GlobalAPIURLRequestHook)(nil)

func (i *GlobalAPIURLRequestHook) BeforeRequest(hookCtx BeforeRequestContext, req *http.Request) (*http.Request, error) {
	// NOTE: the list below was generated with
	// for op in $(rg --no-line-number -o "operationId: (.*)" -r '$1' ../platform-api/src/konnect/definitions/identity/v3/openapi.yaml); do printf "case \"$op\":\n\tfallthrough\n"; done
	switch hookCtx.OperationID {
	case "post-auth0-register-internal":
		fallthrough
	case "get-auth0-organizations-internal":
		fallthrough
	case "get-users-internal":
		fallthrough
	case "post-oauth-device-authorize":
		fallthrough
	case "post-oauth-device-token":
		fallthrough
	case "post-oauth-device-authorize-user":
		fallthrough
	case "patch-oauth-device-confirm":
		fallthrough
	case "get-impersonation-settings":
		fallthrough
	case "update-impersonation-settings":
		fallthrough
	case "get-authentication-settings":
		fallthrough
	case "update-authentication-settings":
		fallthrough
	case "invite-user":
		fallthrough
	case "get-identity-providers":
		fallthrough
	case "create-identity-provider":
		fallthrough
	case "get-identity-provider":
		fallthrough
	case "update-identity-provider":
		fallthrough
	case "delete-identity-provider":
		fallthrough
	case "get-idp-configuration":
		fallthrough
	case "update-idp-configuration":
		fallthrough
	case "update-idp-team-mappings":
		fallthrough
	case "get-idp-team-mappings":
		fallthrough
	case "get-team-group-mappings":
		fallthrough
	case "patch-team-group-mappings":
		fallthrough
	case "get-predefined-roles":
		fallthrough
	case "list-teams":
		fallthrough
	case "create-team":
		fallthrough
	case "list-team-users":
		fallthrough
	case "add-user-to-team":
		fallthrough
	case "get-team":
		fallthrough
	case "update-team":
		fallthrough
	case "delete-team":
		fallthrough
	case "remove-user-from-team":
		fallthrough
	case "list-team-roles":
		fallthrough
	case "teams-assign-role":
		fallthrough
	case "teams-remove-role":
		fallthrough
	case "list-users":
		fallthrough
	case "get-user":
		fallthrough
	case "update-user":
		fallthrough
	case "delete-user":
		fallthrough
	case "list-user-teams":
		fallthrough
	case "list-user-roles":
		fallthrough
	case "users-assign-role":
		fallthrough
	case "users-remove-role":
		fallthrough
	case "get-system-accounts":
		fallthrough
	case "post-system-accounts":
		fallthrough
	case "get-system-accounts-id":
		fallthrough
	case "patch-system-accounts-id":
		fallthrough
	case "delete-system-accounts-id":
		fallthrough
	case "get-system-account-id-access-tokens":
		fallthrough
	case "post-system-accounts-id-access-tokens":
		fallthrough
	case "get-system-accounts-id-access-tokens-id":
		fallthrough
	case "patch-system-accounts-id-access-tokens-id":
		fallthrough
	case "delete-system-accounts-id-access-tokens-id":
		fallthrough
	case "get-system-accounts-assigned-roles-internal":
		fallthrough
	case "create-system-accounts-assigned-roles-internal":
		fallthrough
	case "get-system-accounts-accountId-assigned-roles":
		fallthrough
	case "post-system-accounts-accountId-assigned-roles":
		fallthrough
	case "delete-system-accounts-accountId-assigned-roles-roleId":
		fallthrough
	case "get-teams-teamId-system-accounts":
		fallthrough
	case "post-teams-teamId-system-accounts":
		fallthrough
	case "delete-teams-teamId-system-accounts-accountId":
		fallthrough
	case "get-system-accounts-accountId-teams":
		fallthrough
	case "get-users-me":
		fallthrough
	case "delete-users-me":
		fallthrough
	case "patch-users-me":
		fallthrough
	case "get-users-me-permissions":
		fallthrough
	case "get-organizations-me":
		fallthrough
	case "update-organizations-me":
		fallthrough
	case "refresh-token":
		fallthrough
	case "logout":
		fallthrough
	case "resolveCustomer":
		fallthrough
	case "authenticate-sso":
		// NOTE(pmalek): This is because we merge OpenAPI specs and /organizations/me
		// is only served by the global API.
		// @mheap mentioned that we can add operation specific URLs to do away with this.
		if strings.HasSuffix(req.URL.Host, "api.konghq.tech") {
			req.URL.Host = "global.api.konghq.tech"
			req.Host = "global.api.konghq.tech"
		} else {
			req.URL.Host = "global.api.konghq.com"
			req.Host = "global.api.konghq.com"
		}

	default:

	}
	return req, nil
}
