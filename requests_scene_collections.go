package obsws

// This file is automatically generated.
// https://github.com/christopher-dG/go-obs-websocket/blob/master/codegen/protocol.py

// SetCurrentSceneCollectionRequest : Change the active scene collection.
// Since obs-websocket version: 4.0.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setcurrentscenecollection
type SetCurrentSceneCollectionRequest struct {
	// Name of the desired scene collection.
	// Required: Yes.
	ScName string `json:"sc-name"`
	_request
}

// NewSetCurrentSceneCollectionRequest returns a new SetCurrentSceneCollectionRequest.
func (c *Client) NewSetCurrentSceneCollectionRequest(scName string) SetCurrentSceneCollectionRequest {
	return SetCurrentSceneCollectionRequest{
		scName,
		_request{
			MessageID:   c.getMessageID(),
			RequestType: "SetCurrentSceneCollection",
		},
	}

}

// ID returns the request's message ID.
func (r SetCurrentSceneCollectionRequest) ID() string { return r.MessageID }

// Type returns the request's message type.
func (r SetCurrentSceneCollectionRequest) Type() string { return r.RequestType }

// SetCurrentSceneCollectionResponse : Response for SetCurrentSceneCollectionRequest.
// Since obs-websocket version: 4.0.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setcurrentscenecollection
type SetCurrentSceneCollectionResponse _response

// ID returns the response's message ID.
func (r SetCurrentSceneCollectionResponse) ID() string { return r.MessageID }

// Stat returns the response's status.
func (r SetCurrentSceneCollectionResponse) Stat() string { return r.Status }

// Err returns the response's error.
func (r SetCurrentSceneCollectionResponse) Err() string { return r.Error }

// GetCurrentSceneCollectionRequest : Get the name of the current scene collection.
// Since obs-websocket version: 4.0.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getcurrentscenecollection
type GetCurrentSceneCollectionRequest _request

// NewGetCurrentSceneCollectionRequest returns a new GetCurrentSceneCollectionRequest.
func (c *Client) NewGetCurrentSceneCollectionRequest() GetCurrentSceneCollectionRequest {
	return GetCurrentSceneCollectionRequest{MessageID: c.getMessageID(), RequestType: "GetCurrentSceneCollection"}
}

// ID returns the request's message ID.
func (r GetCurrentSceneCollectionRequest) ID() string { return r.MessageID }

// Type returns the request's message type.
func (r GetCurrentSceneCollectionRequest) Type() string { return r.RequestType }

// GetCurrentSceneCollectionResponse : Response for GetCurrentSceneCollectionRequest.
// Since obs-websocket version: 4.0.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getcurrentscenecollection
type GetCurrentSceneCollectionResponse struct {
	// Name of the currently active scene collection.
	// Required: Yes.
	ScName    string `json:"sc-name"`
	_response `mapstructure:",squash"`
}

// ID returns the response's message ID.
func (r GetCurrentSceneCollectionResponse) ID() string { return r.MessageID }

// Stat returns the response's status.
func (r GetCurrentSceneCollectionResponse) Stat() string { return r.Status }

// Err returns the response's error.
func (r GetCurrentSceneCollectionResponse) Err() string { return r.Error }

// ListSceneCollectionsRequest : List available scene collections.
// Since obs-websocket version: 4.0.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#listscenecollections
type ListSceneCollectionsRequest _request

// NewListSceneCollectionsRequest returns a new ListSceneCollectionsRequest.
func (c *Client) NewListSceneCollectionsRequest() ListSceneCollectionsRequest {
	return ListSceneCollectionsRequest{MessageID: c.getMessageID(), RequestType: "ListSceneCollections"}
}

// ID returns the request's message ID.
func (r ListSceneCollectionsRequest) ID() string { return r.MessageID }

// Type returns the request's message type.
func (r ListSceneCollectionsRequest) Type() string { return r.RequestType }

// ListSceneCollectionsResponse : Response for ListSceneCollectionsRequest.
// Since obs-websocket version: 4.0.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#listscenecollections
type ListSceneCollectionsResponse struct {
	// Scene collections list.
	// Required: Yes.
	// TODO: Unknown type (Object|Array).
	SceneCollections interface{} `json:"scene-collections"`
	// Required: Yes.
	SceneCollectionsStar string `json:"scene-collections.*."`
	_response            `mapstructure:",squash"`
}

// ID returns the response's message ID.
func (r ListSceneCollectionsResponse) ID() string { return r.MessageID }

// Stat returns the response's status.
func (r ListSceneCollectionsResponse) Stat() string { return r.Status }

// Err returns the response's error.
func (r ListSceneCollectionsResponse) Err() string { return r.Error }
