package obsws

// This file is automatically generated.
// https://github.com/christopher-dG/go-obs-websocket/blob/master/codegen/protocol.py

// SourceOrderChangedEvent : Scene items have been reordered.
// Since obs-websocket version: 4.0.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#sourceorderchanged
type SourceOrderChangedEvent struct {
	// Name of the scene where items have been reordered.
	// Required: Yes.
	Name string `json:"name"`
	// Array of sources.
	// Required: Yes.
	Sources []string `json:"sources"`
	_event  `json:",squash"`
}

// Type returns the event's update type.
func (e SourceOrderChangedEvent) Type() string { return e.UpdateType }

// StreamTC returns the event's stream timecode.
func (e SourceOrderChangedEvent) StreamTC() string { return e.StreamTimecode }

// RecTC returns the event's recording timecode.
func (e SourceOrderChangedEvent) RecTC() string { return e.RecTimecode }

// SceneItemAddedEvent : An item has been added to the current scene.
// Since obs-websocket version: 4.0.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#sceneitemadded
type SceneItemAddedEvent struct {
	// Name of the scene.
	// Required: Yes.
	SceneName string `json:"scene-name"`
	// Name of the item added to the scene.
	// Required: Yes.
	ItemName string `json:"item-name"`
	_event   `json:",squash"`
}

// Type returns the event's update type.
func (e SceneItemAddedEvent) Type() string { return e.UpdateType }

// StreamTC returns the event's stream timecode.
func (e SceneItemAddedEvent) StreamTC() string { return e.StreamTimecode }

// RecTC returns the event's recording timecode.
func (e SceneItemAddedEvent) RecTC() string { return e.RecTimecode }

// SceneItemRemovedEvent : An item has been removed from the current scene.
// Since obs-websocket version: 4.0.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#sceneitemremoved
type SceneItemRemovedEvent struct {
	// Name of the scene.
	// Required: Yes.
	SceneName string `json:"scene-name"`
	// Name of the item removed from the scene.
	// Required: Yes.
	ItemName string `json:"item-name"`
	_event   `json:",squash"`
}

// Type returns the event's update type.
func (e SceneItemRemovedEvent) Type() string { return e.UpdateType }

// StreamTC returns the event's stream timecode.
func (e SceneItemRemovedEvent) StreamTC() string { return e.StreamTimecode }

// RecTC returns the event's recording timecode.
func (e SceneItemRemovedEvent) RecTC() string { return e.RecTimecode }

// SceneItemVisibilityChangedEvent : An item's visibility has been toggled.
// Since obs-websocket version: 4.0.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#sceneitemvisibilitychanged
type SceneItemVisibilityChangedEvent struct {
	// Name of the scene.
	// Required: Yes.
	SceneName string `json:"scene-name"`
	// Name of the item in the scene.
	// Required: Yes.
	ItemName string `json:"item-name"`
	// New visibility state of the item.
	// Required: Yes.
	ItemVisible bool `json:"item-visible"`
	_event      `json:",squash"`
}

// Type returns the event's update type.
func (e SceneItemVisibilityChangedEvent) Type() string { return e.UpdateType }

// StreamTC returns the event's stream timecode.
func (e SceneItemVisibilityChangedEvent) StreamTC() string { return e.StreamTimecode }

// RecTC returns the event's recording timecode.
func (e SceneItemVisibilityChangedEvent) RecTC() string { return e.RecTimecode }
