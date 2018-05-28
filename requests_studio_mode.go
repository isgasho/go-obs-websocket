package obsws

// This file is automatically generated.
// https://github.com/christopher-dG/go-obs-websocket/blob/master/codegen/protocol.py

// GetStudioModeStatusRequest : Indicates if Studio Mode is currently enabled.
// Since obs-websocket version: 4.1.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getstudiomodestatus
type GetStudioModeStatusRequest _request

// NewGetStudioModeStatusRequest returns a new GetStudioModeStatusRequest.
func (c *Client) NewGetStudioModeStatusRequest() GetStudioModeStatusRequest {
	return GetStudioModeStatusRequest{MessageID: c.getMessageID(), RequestType: "GetStudioModeStatus"}
}

// ID returns the request's message ID.
func (r GetStudioModeStatusRequest) ID() string { return r.MessageID }

// Type returns the request's message type.
func (r GetStudioModeStatusRequest) Type() string { return r.RequestType }

// GetStudioModeStatusResponse : Response for GetStudioModeStatusRequest.
// Since obs-websocket version: 4.1.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getstudiomodestatus
type GetStudioModeStatusResponse struct {
	// Indicates if Studio Mode is enabled.
	// Required: Yes.
	StudioMode bool `mapstructure:"studio-mode"`
	_response  `mapstructure:",squash"`
}

// ID returns the response's message ID.
func (r GetStudioModeStatusResponse) ID() string { return r.MessageID }

// Stat returns the response's status.
func (r GetStudioModeStatusResponse) Stat() string { return r.Status }

// Err returns the response's error.
func (r GetStudioModeStatusResponse) Err() string { return r.Error }

// GetPreviewSceneRequest : Get the name of the currently previewed scene and its list of sources.
// Will return an `error` if Studio Mode is not enabled.
// Since obs-websocket version: 4.1.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getpreviewscene
type GetPreviewSceneRequest _request

// NewGetPreviewSceneRequest returns a new GetPreviewSceneRequest.
func (c *Client) NewGetPreviewSceneRequest() GetPreviewSceneRequest {
	return GetPreviewSceneRequest{MessageID: c.getMessageID(), RequestType: "GetPreviewScene"}
}

// ID returns the request's message ID.
func (r GetPreviewSceneRequest) ID() string { return r.MessageID }

// Type returns the request's message type.
func (r GetPreviewSceneRequest) Type() string { return r.RequestType }

// GetPreviewSceneResponse : Response for GetPreviewSceneRequest.
// Since obs-websocket version: 4.1.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getpreviewscene
type GetPreviewSceneResponse struct {
	// The name of the active preview scene.
	// Required: Yes.
	Name string `mapstructure:"name"`
	// Required: Yes.
	// TODO: Unknown type (Source|Array).
	Sources   interface{} `mapstructure:"sources"`
	_response `mapstructure:",squash"`
}

// ID returns the response's message ID.
func (r GetPreviewSceneResponse) ID() string { return r.MessageID }

// Stat returns the response's status.
func (r GetPreviewSceneResponse) Stat() string { return r.Status }

// Err returns the response's error.
func (r GetPreviewSceneResponse) Err() string { return r.Error }

// SetPreviewSceneRequest : Set the active preview scene.
// Will return an `error` if Studio Mode is not enabled.
// Since obs-websocket version: 4.1.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setpreviewscene
type SetPreviewSceneRequest struct {
	// The name of the scene to preview.
	// Required: Yes.
	SceneName string `json:"scene-name"`
	_request
}

// NewSetPreviewSceneRequest returns a new SetPreviewSceneRequest.
func (c *Client) NewSetPreviewSceneRequest(sceneName string) SetPreviewSceneRequest {
	return SetPreviewSceneRequest{
		sceneName,
		_request{
			MessageID:   c.getMessageID(),
			RequestType: "SetPreviewScene",
		},
	}

}

// ID returns the request's message ID.
func (r SetPreviewSceneRequest) ID() string { return r.MessageID }

// Type returns the request's message type.
func (r SetPreviewSceneRequest) Type() string { return r.RequestType }

// SetPreviewSceneResponse : Response for SetPreviewSceneRequest.
// Since obs-websocket version: 4.1.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setpreviewscene
type SetPreviewSceneResponse _response

// ID returns the response's message ID.
func (r SetPreviewSceneResponse) ID() string { return r.MessageID }

// Stat returns the response's status.
func (r SetPreviewSceneResponse) Stat() string { return r.Status }

// Err returns the response's error.
func (r SetPreviewSceneResponse) Err() string { return r.Error }

// TransitionToProgramRequest : Transitions the currently previewed scene to the main output.
// Will return an `error` if Studio Mode is not enabled.
// Since obs-websocket version: 4.1.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#transitiontoprogram
type TransitionToProgramRequest struct {
	// Change the active transition before switching scenes.
	// Defaults to the active transition.
	// Required: No.
	WithTransition map[string]interface{} `json:"with-transition"`
	// Name of the transition.
	// Required: Yes.
	WithTransitionName string `json:"with-transition.name"`
	// Transition duration (in milliseconds).
	// Required: No.
	WithTransitionDuration int `json:"with-transition.duration"`
	_request
}

// NewTransitionToProgramRequest returns a new TransitionToProgramRequest.
func (c *Client) NewTransitionToProgramRequest(
	withTransition map[string]interface{},
	withTransitionName string,
	withTransitionDuration int,
) TransitionToProgramRequest {
	return TransitionToProgramRequest{
		withTransition,
		withTransitionName,
		withTransitionDuration,
		_request{
			MessageID:   c.getMessageID(),
			RequestType: "TransitionToProgram",
		},
	}

}

// ID returns the request's message ID.
func (r TransitionToProgramRequest) ID() string { return r.MessageID }

// Type returns the request's message type.
func (r TransitionToProgramRequest) Type() string { return r.RequestType }

// TransitionToProgramResponse : Response for TransitionToProgramRequest.
// Since obs-websocket version: 4.1.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#transitiontoprogram
type TransitionToProgramResponse _response

// ID returns the response's message ID.
func (r TransitionToProgramResponse) ID() string { return r.MessageID }

// Stat returns the response's status.
func (r TransitionToProgramResponse) Stat() string { return r.Status }

// Err returns the response's error.
func (r TransitionToProgramResponse) Err() string { return r.Error }

// EnableStudioModeRequest : Enables Studio Mode.
// Since obs-websocket version: 4.1.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#enablestudiomode
type EnableStudioModeRequest _request

// NewEnableStudioModeRequest returns a new EnableStudioModeRequest.
func (c *Client) NewEnableStudioModeRequest() EnableStudioModeRequest {
	return EnableStudioModeRequest{MessageID: c.getMessageID(), RequestType: "EnableStudioMode"}
}

// ID returns the request's message ID.
func (r EnableStudioModeRequest) ID() string { return r.MessageID }

// Type returns the request's message type.
func (r EnableStudioModeRequest) Type() string { return r.RequestType }

// EnableStudioModeResponse : Response for EnableStudioModeRequest.
// Since obs-websocket version: 4.1.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#enablestudiomode
type EnableStudioModeResponse _response

// ID returns the response's message ID.
func (r EnableStudioModeResponse) ID() string { return r.MessageID }

// Stat returns the response's status.
func (r EnableStudioModeResponse) Stat() string { return r.Status }

// Err returns the response's error.
func (r EnableStudioModeResponse) Err() string { return r.Error }

// DisableStudioModeRequest : Disables Studio Mode.
// Since obs-websocket version: 4.1.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#disablestudiomode
type DisableStudioModeRequest _request

// NewDisableStudioModeRequest returns a new DisableStudioModeRequest.
func (c *Client) NewDisableStudioModeRequest() DisableStudioModeRequest {
	return DisableStudioModeRequest{MessageID: c.getMessageID(), RequestType: "DisableStudioMode"}
}

// ID returns the request's message ID.
func (r DisableStudioModeRequest) ID() string { return r.MessageID }

// Type returns the request's message type.
func (r DisableStudioModeRequest) Type() string { return r.RequestType }

// DisableStudioModeResponse : Response for DisableStudioModeRequest.
// Since obs-websocket version: 4.1.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#disablestudiomode
type DisableStudioModeResponse _response

// ID returns the response's message ID.
func (r DisableStudioModeResponse) ID() string { return r.MessageID }

// Stat returns the response's status.
func (r DisableStudioModeResponse) Stat() string { return r.Status }

// Err returns the response's error.
func (r DisableStudioModeResponse) Err() string { return r.Error }

// ToggleStudioModeRequest : Toggles Studio Mode.
// Since obs-websocket version: 4.1.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#togglestudiomode
type ToggleStudioModeRequest _request

// NewToggleStudioModeRequest returns a new ToggleStudioModeRequest.
func (c *Client) NewToggleStudioModeRequest() ToggleStudioModeRequest {
	return ToggleStudioModeRequest{MessageID: c.getMessageID(), RequestType: "ToggleStudioMode"}
}

// ID returns the request's message ID.
func (r ToggleStudioModeRequest) ID() string { return r.MessageID }

// Type returns the request's message type.
func (r ToggleStudioModeRequest) Type() string { return r.RequestType }

// ToggleStudioModeResponse : Response for ToggleStudioModeRequest.
// Since obs-websocket version: 4.1.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#togglestudiomode
type ToggleStudioModeResponse _response

// ID returns the response's message ID.
func (r ToggleStudioModeResponse) ID() string { return r.MessageID }

// Stat returns the response's status.
func (r ToggleStudioModeResponse) Stat() string { return r.Status }

// Err returns the response's error.
func (r ToggleStudioModeResponse) Err() string { return r.Error }
