// +build go1.9

// Code generated by cdpgen. DO NOT EDIT.

package network

import (
	"github.com/mafredri/cdp/protocol/internal"
	"github.com/mafredri/cdp/rpcc"
)

// LoadingFailedReply is the reply for LoadingFailed events.
type LoadingFailedReply struct {
	RequestID     RequestID                 `json:"requestId"`               // Request identifier.
	Timestamp     MonotonicTime             `json:"timestamp"`               // Timestamp.
	Type          internal.PageResourceType `json:"type"`                    // Resource type.
	ErrorText     string                    `json:"errorText"`               // User friendly error message.
	Canceled      *bool                     `json:"canceled,omitempty"`      // True if loading was canceled.
	BlockedReason BlockedReason             `json:"blockedReason,omitempty"` // The reason why loading was blocked, if any.
}

// LoadingFinishedClient is a client for LoadingFinished events. Fired when
// HTTP request has finished loading.
type LoadingFinishedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*LoadingFinishedReply, error)
	rpcc.Stream
}

// RequestInterceptedReply is the reply for RequestIntercepted events.
type RequestInterceptedReply struct {
	InterceptionID      InterceptionID            `json:"interceptionId"`                // Each request the page makes will have a unique id, however if any redirects are encountered while processing that fetch, they will be reported with the same id as the original fetch. Likewise if HTTP authentication is needed then the same fetch id will be used.
	Request             Request                   `json:"request"`                       // No description.
	FrameID             internal.PageFrameID      `json:"frameId"`                       // The id of the frame that initiated the request.
	ResourceType        internal.PageResourceType `json:"resourceType"`                  // How the requested resource will be used.
	IsNavigationRequest bool                      `json:"isNavigationRequest"`           // Whether this is a navigation request, which can abort the navigation completely.
	RedirectURL         *string                   `json:"redirectUrl,omitempty"`         // Redirect location, only sent if a redirect was intercepted.
	AuthChallenge       *AuthChallenge            `json:"authChallenge,omitempty"`       // Details of the Authorization Challenge encountered. If this is set then continueInterceptedRequest must contain an authChallengeResponse.
	ResponseErrorReason ErrorReason               `json:"responseErrorReason,omitempty"` // Response error if intercepted at response stage or if redirect occurred while intercepting request.
	ResponseStatusCode  *int                      `json:"responseStatusCode,omitempty"`  // Response code if intercepted at response stage or if redirect occurred while intercepting request or auth retry occurred.
	ResponseHeaders     Headers                   `json:"responseHeaders,omitempty"`     // Response headers if intercepted at the response stage or if redirect occurred while intercepting request or auth retry occurred.
}

// RequestServedFromCacheClient is a client for RequestServedFromCache events.
// Fired if request ended up loading from cache.
type RequestServedFromCacheClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*RequestServedFromCacheReply, error)
	rpcc.Stream
}

// RequestWillBeSentReply is the reply for RequestWillBeSent events.
type RequestWillBeSentReply struct {
	RequestID        RequestID                  `json:"requestId"`                  // Request identifier.
	LoaderID         LoaderID                   `json:"loaderId"`                   // Loader identifier. Empty string if the request is fetched from worker.
	DocumentURL      string                     `json:"documentURL"`                // URL of the document this request is loaded for.
	Request          Request                    `json:"request"`                    // Request data.
	Timestamp        MonotonicTime              `json:"timestamp"`                  // Timestamp.
	WallTime         TimeSinceEpoch             `json:"wallTime"`                   // Timestamp.
	Initiator        Initiator                  `json:"initiator"`                  // Request initiator.
	RedirectResponse *Response                  `json:"redirectResponse,omitempty"` // Redirect response data.
	Type             *internal.PageResourceType `json:"type,omitempty"`             // Type of this resource.
	FrameID          *internal.PageFrameID      `json:"frameId,omitempty"`          // Frame identifier.
	HasUserGesture   *bool                      `json:"hasUserGesture,omitempty"`   // Whether the request is initiated by a user gesture. Defaults to false.
}

// ResourceChangedPriorityClient is a client for ResourceChangedPriority events.
// Fired when resource loading priority is changed
type ResourceChangedPriorityClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*ResourceChangedPriorityReply, error)
	rpcc.Stream
}

// ResponseReceivedReply is the reply for ResponseReceived events.
type ResponseReceivedReply struct {
	RequestID RequestID                 `json:"requestId"`         // Request identifier.
	LoaderID  LoaderID                  `json:"loaderId"`          // Loader identifier. Empty string if the request is fetched from worker.
	Timestamp MonotonicTime             `json:"timestamp"`         // Timestamp.
	Type      internal.PageResourceType `json:"type"`              // Resource type.
	Response  Response                  `json:"response"`          // Response data.
	FrameID   *internal.PageFrameID     `json:"frameId,omitempty"` // Frame identifier.
}

// WebSocketClosedClient is a client for WebSocketClosed events. Fired when
// WebSocket is closed.
type WebSocketClosedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*WebSocketClosedReply, error)
	rpcc.Stream
}
