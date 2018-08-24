// Code generated by cdpgen. DO NOT EDIT.

package browser

// WindowID
//
// Note: This type is experimental.
type WindowID int

// WindowState The state of the browser window.
//
// Note: This type is experimental.
type WindowState string

// WindowState as enums.
const (
	WindowStateNotSet     WindowState = ""
	WindowStateNormal     WindowState = "normal"
	WindowStateMinimized  WindowState = "minimized"
	WindowStateMaximized  WindowState = "maximized"
	WindowStateFullscreen WindowState = "fullscreen"
)

func (e WindowState) Valid() bool {
	switch e {
	case "normal", "minimized", "maximized", "fullscreen":
		return true
	default:
		return false
	}
}

func (e WindowState) String() string {
	return string(e)
}

// Bounds Browser window bounds information
//
// Note: This type is experimental.
type Bounds struct {
	Left        *int        `json:"left,omitempty"`        // The offset from the left edge of the screen to the window in pixels.
	Top         *int        `json:"top,omitempty"`         // The offset from the top edge of the screen to the window in pixels.
	Width       *int        `json:"width,omitempty"`       // The window width in pixels.
	Height      *int        `json:"height,omitempty"`      // The window height in pixels.
	WindowState WindowState `json:"windowState,omitempty"` // The window state. Default to normal.
}

// PermissionType
//
// Note: This type is experimental.
type PermissionType string

// PermissionType as enums.
const (
	PermissionTypeNotSet                   PermissionType = ""
	PermissionTypeAccessibilityEvents      PermissionType = "accessibilityEvents"
	PermissionTypeAudioCapture             PermissionType = "audioCapture"
	PermissionTypeBackgroundSync           PermissionType = "backgroundSync"
	PermissionTypeClipboardRead            PermissionType = "clipboardRead"
	PermissionTypeClipboardWrite           PermissionType = "clipboardWrite"
	PermissionTypeDurableStorage           PermissionType = "durableStorage"
	PermissionTypeFlash                    PermissionType = "flash"
	PermissionTypeGeolocation              PermissionType = "geolocation"
	PermissionTypeMidi                     PermissionType = "midi"
	PermissionTypeMidiSysex                PermissionType = "midiSysex"
	PermissionTypeNotifications            PermissionType = "notifications"
	PermissionTypePaymentHandler           PermissionType = "paymentHandler"
	PermissionTypeProtectedMediaIdentifier PermissionType = "protectedMediaIdentifier"
	PermissionTypeSensors                  PermissionType = "sensors"
	PermissionTypeVideoCapture             PermissionType = "videoCapture"
)

func (e PermissionType) Valid() bool {
	switch e {
	case "accessibilityEvents", "audioCapture", "backgroundSync", "clipboardRead", "clipboardWrite", "durableStorage", "flash", "geolocation", "midi", "midiSysex", "notifications", "paymentHandler", "protectedMediaIdentifier", "sensors", "videoCapture":
		return true
	default:
		return false
	}
}

func (e PermissionType) String() string {
	return string(e)
}

// Bucket Chrome histogram bucket.
//
// Note: This type is experimental.
type Bucket struct {
	Low   int `json:"low"`   // Minimum value (inclusive).
	High  int `json:"high"`  // Maximum value (exclusive).
	Count int `json:"count"` // Number of samples.
}

// Histogram Chrome histogram.
//
// Note: This type is experimental.
type Histogram struct {
	Name    string   `json:"name"`    // Name.
	Sum     int      `json:"sum"`     // Sum of sample values.
	Count   int      `json:"count"`   // Total number of samples.
	Buckets []Bucket `json:"buckets"` // Buckets.
}
