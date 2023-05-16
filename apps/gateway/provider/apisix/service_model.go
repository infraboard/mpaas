package apisix

type ServiceList struct {
	Total int        `json:"total"`
	List  []*Service `json:"list"`
}

type Service struct {
	*Meta
	*CreateServiceRequest
}

type CreateServiceRequest struct {
}
