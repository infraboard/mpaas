package route

import "context"

type Service interface {
	CreateRoute(context.Context, *CreateRouteRequest) (*Route, error)
	QueryRoute(context.Context, *QueryRouteRequest) (*RouteList, error)
	DescribeRoute(context.Context, *DescribeRouteRequest) (*Route, error)
	UpdateRoute(context.Context, *UpdateRouteRequest) (*Route, error)
	DeleteRoute(context.Context, *DeleteRouteRequest) (*Route, error)
}

func NewQueryRouteRequest() *QueryRouteRequest {
	return &QueryRouteRequest{}
}

type QueryRouteRequest struct {
}

func NewDescribeRouteRequest(routeId string) *DescribeRouteRequest {
	return &DescribeRouteRequest{
		RouteId: routeId,
	}
}

type DescribeRouteRequest struct {
	RouteId string
}

type UpdateRouteRequest struct {
}

type DeleteRouteRequest struct {
	RouteId string
}
