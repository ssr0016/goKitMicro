package transport

// Requests and Responses for RPC endpoints
// The service methods expose as RPC endpoints. So we need to define message types to be used for send and receive messages over RPC endpoints. Let’s define structs for request and response types to be used for RPC endpoints on Order service:
// Listing 4. Message types for Request and Response for RPC endpoints

import "github.com/go-kit/kit/examples/addsvc/order"

// CreateRequest holds the request parameters values for the Create method.
type CreateRequest struct {
	Order order.Order
}

// CreateResponse holds the response values for the Create method.
type CreateResponse struct {
	ID  string `json:"id"`
	Err error  `json:"error,omitempty"`
}

// GetByIDRequest holds the request parameters for the GetByID method.
type GetByIDRequest struct {
	ID string
}

type GetByIDResponse struct {
	Order order.Order `json:"order"`
	Err   error       `json:"error,omitempty"`
}

// ChangeStatusRequest holds the request parameters for the ChangeStatus method.
type ChangeStatusRequest struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

// ChangeStatusResponse holds the response values for the ChangeStatus method.
type ChangeStatusResponse struct {
	Err error `json:"error,omitempty"`
}
