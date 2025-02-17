// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"errors"
	"fmt"
	"github.com/Kong/sdk-konnect-go/internal/utils"
)

type RouteType string

const (
	RouteTypeRouteJSON       RouteType = "RouteJson"
	RouteTypeRouteExpression RouteType = "RouteExpression"
)

type Route struct {
	RouteJSON       *RouteJSON       `queryParam:"inline"`
	RouteExpression *RouteExpression `queryParam:"inline"`

	Type RouteType
}

func CreateRouteRouteJSON(routeJSON RouteJSON) Route {
	typ := RouteTypeRouteJSON

	return Route{
		RouteJSON: &routeJSON,
		Type:      typ,
	}
}

func CreateRouteRouteExpression(routeExpression RouteExpression) Route {
	typ := RouteTypeRouteExpression

	return Route{
		RouteExpression: &routeExpression,
		Type:            typ,
	}
}

func (u *Route) UnmarshalJSON(data []byte) error {

	var routeExpression RouteExpression = RouteExpression{}
	if err := utils.UnmarshalJSON(data, &routeExpression, "", true, true); err == nil {
		u.RouteExpression = &routeExpression
		u.Type = RouteTypeRouteExpression
		return nil
	}

	var routeJSON RouteJSON = RouteJSON{}
	if err := utils.UnmarshalJSON(data, &routeJSON, "", true, true); err == nil {
		u.RouteJSON = &routeJSON
		u.Type = RouteTypeRouteJSON
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for Route", string(data))
}

func (u Route) MarshalJSON() ([]byte, error) {
	if u.RouteJSON != nil {
		return utils.MarshalJSON(u.RouteJSON, "", true)
	}

	if u.RouteExpression != nil {
		return utils.MarshalJSON(u.RouteExpression, "", true)
	}

	return nil, errors.New("could not marshal union type Route: all fields are null")
}
