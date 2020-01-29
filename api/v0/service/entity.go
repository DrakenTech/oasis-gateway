package service

import "github.com/oasislabs/oasis-gateway/rpc"

// RequestType defines the type of the request. May be
// useful for serialization and deserialization
type RequestType uint

const (
	Deploy       RequestType = 0
	Execute      RequestType = 1
	Poll         RequestType = 2
	GetCode      RequestType = 3
	GetExpiry    RequestType = 4
	GetPublicKey RequestType = 5
)

// Request is the type implemented by requests expected
// by the API handlers
type Request interface {
	// Type returns the type of the request
	Type() RequestType
}

// AsyncResponse is the response returned by APIs that are asynchronous
// that return an ID that can be used by the user to receive and identify
// a response to the request when it is ready
type AsyncResponse struct {
	// ID to identify an asynchronous response. It uniquely identifies the
	// event and orders it in the sequence of events expected by the user
	ID uint64 `json:"id"`
}

// ExecuteServiceRequest is is used by the user to trigger a service
// execution. A client is always subscribed to a subscription with
// topic "service" from which the client can retrieve the asynchronous
// results to this request
type ExecuteServiceRequest struct {
	// Data is a blob of data that the user wants to pass to the service
	// as argument
	Data string `json:"data"`

	// Address where the service can be found
	Address string `json:"address"`
}

// Type implementation of Request for ExecuteServiceRequest
func (r ExecuteServiceRequest) Type() RequestType {
	return Execute
}

// ExecuteServiceResponse is an asynchronous response that will be obtained
// using the polling mechanisms
type ExecuteServiceResponse AsyncResponse

// DeployServiceRequest is issued by the user to trigger a service
// execution. A client is always subscribed to a subscription with
// topic "service" from which the client can retrieve the asynchronous
// results to this request
type DeployServiceRequest struct {
	// Data is a blob of data that the user wants to pass as argument for
	// the deployment of a service
	Data string `json:"data"`
}

// Type implementation of Request for DeployServiceRequest
func (r DeployServiceRequest) Type() RequestType {
	return Deploy
}

// DeployServiceResponse is an asynchronous response that will be obtained
// using the polling mechanism
type DeployServiceResponse AsyncResponse

// GetCodeRequest is a request to retrieve the code
// associated with a specific service
type GetCodeRequest struct {
	// Address is the unique address that identifies the service,
	// is generated when a service is deployed and it can be used
	// for service execution
	Address string `json:"address"`
}

// Type implementation of Request for GetCodeRequest
func (r GetCodeRequest) Type() RequestType {
	return GetCode
}

// GetCodeResponse is the response in which the code
// associated with the service is provided
type GetCodeResponse struct {
	// Address is the unique address that identifies the service,
	// is generated when a service is deployed and it can be used
	// for service execution
	Address string `json:"address"`

	// Code associated with the service
	Code string `json:"code"`
}

// GetExpiryRequest is a request to retrieve the expiration timestamp
// associated with a specific service
type GetExpiryRequest struct {
	// Address is the unique address that identifies the service,
	// is generated when a service is deployed and it can be used
	// for service execution
	Address string `json:"address"`
}

// Type implementation of Request for GetExpiryRequest
func (r GetExpiryRequest) Type() RequestType {
	return GetExpiry
}

// GetExpiryResponse is the response in which the expiration timestamp
// associated with the service is provided
type GetExpiryResponse struct {
	// Address is the unique address that identifies the service,
	// is generated when a service is deployed and it can be used
	// for service execution
	Address string `json:"address"`

	// Expiration timestamp associated with the service
	Expiry uint64 `json:"expiry"`
}

// GetPublicKeyRequest is a request to retrieve the public key
// associated with a specific service
type GetPublicKeyRequest struct {
	// Address is the unique address that identifies the service,
	// is generated when a service is deployed and it can be used
	// for service execution
	Address string `json:"address"`
}

// Type implementation of Request for GetPublicKeyRequest
func (r GetPublicKeyRequest) Type() RequestType {
	return GetPublicKey
}

// GetPublicKeyResponse is the response in which the public key
// associated with the service is provided
type GetPublicKeyResponse struct {
	// Timestamp at which the public key expired
	Timestamp uint64 `json:"timestamp"`
	// Address is the unique address that identifies the service,
	// is generated when a service is deployed and it can be used
	// for service execution
	Address string `json:"address"`

	// PublicKey associated with the service
	PublicKey string `json:"publicKey"`

	// Signature generated by the key manager for authentication of the
	// public key
	Signature string `json:"signature"`
}

// PollServiceRequest is a request that allows the user to
// poll for events either from asynchronous responses
type PollServiceRequest struct {
	// Offset at which events need to be provided. Events are all ordered
	// with sequence numbers and it is up to the client to specify which
	// events it wants to receive from an offset in the sequence
	Offset uint64 `json:"offset"`

	// Count for the number of items the client would prefer to receive
	// at most from a single response
	Count uint `json:"count"`

	// DiscardPrevious allows the client to define whether the server should
	// discard all the events that have a sequence number lower than the offer
	DiscardPrevious bool `json:"discardPrevious"`
}

// Type implementation of Request for PollServiceRequest
func (r PollServiceRequest) Type() RequestType {
	return Poll
}

// Event is an interface for types that can be fetched by polling on
// a service
type Event interface {
	// EventID is the ID that uniquely identifies the event and it is found
	// inside a sequence of events
	EventID() uint64
}

// PollServiceResponse returns a list of asynchronous responses
// the client requested
type PollServiceResponse struct {
	// Offset is the base offset the requests were got from
	Offset uint64 `json:"offset"`

	// Events is the list of events that the server has starting from
	// the provided Offset
	Events []Event `json:"events"`
}

// ExecuteServiceEvent is the event that can be polled by the user
// as a result to a ServiceExecutionRequest
type ExecuteServiceEvent struct {
	// ID to identify an asynchronous response. It uniquely identifies the
	// event and orders it in the sequence of events expected by the user
	ID uint64 `json:"id"`

	// Address is the unique address that identifies the service,
	// is generated when a service is deployed and it can be used
	// for service execution
	Address string `json:"address"`

	// Output generated by the service at the end of its execution
	Output string `json:"output"`
}

// DeployServiceEvent is the event that can be polled by the user
// as a result to a ServiceExecutionRequest
type DeployServiceEvent struct {
	// ID to identify an asynchronous response. It uniquely identifies the
	// event and orders it in the sequence of events expected by the user
	ID uint64 `json:"id"`

	// Address is the unique address that identifies the service,
	// is generated when a service is deployed and it can be used
	// for service execution
	Address string `json:"address"`
}

// ErrorEvent is the event that can be polled by the user
// as a result to a request that failed
type ErrorEvent struct {
	// ID to identify an asynchronous response. It uniquely identifies the
	// event and orders it in the sequence of events expected by the user
	ID uint64 `json:"id"`

	// Cause is the error that caused the event to failed
	Cause rpc.Error `json:"cause"`
}

// EventID is the implementation of rpc.Event for ExecuteServiceEvent
func (e ExecuteServiceEvent) EventID() uint64 {
	return e.ID
}

// EventID is the implementation of rpc.Event for DeployServiceEvent
func (e DeployServiceEvent) EventID() uint64 {
	return e.ID
}

// EventID is the implementation of rpc.Event for ErrorEvent
func (e ErrorEvent) EventID() uint64 {
	return e.ID
}
