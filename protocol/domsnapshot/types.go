// Code generated by cdpgen. DO NOT EDIT.

package domsnapshot

import (
	"github.com/mafredri/cdp/protocol/dom"
)

// InlineTextBox Details of post layout rendered text positions. The exact
// layout should not be regarded as stable and may change between versions.
type InlineTextBox struct {
	BoundingBox         dom.Rect `json:"boundingBox"`         // The absolute position bounding box.
	StartCharacterIndex int      `json:"startCharacterIndex"` // The starting index in characters, for this post layout textbox substring. Characters that would be represented as a surrogate pair in UTF-16 have length 2.
	NumCharacters       int      `json:"numCharacters"`       // The number of characters in this post layout textbox substring. Characters that would be represented as a surrogate pair in UTF-16 have length 2.
}

// LayoutTreeNode Details of an element in the DOM tree with a LayoutObject.
type LayoutTreeNode struct {
	DOMNodeIndex      int             `json:"domNodeIndex"`                // The index of the related DOM node in the `domNodes` array returned by `getSnapshot`.
	BoundingBox       dom.Rect        `json:"boundingBox"`                 // The absolute position bounding box.
	LayoutText        *string         `json:"layoutText,omitempty"`        // Contents of the LayoutText, if any.
	InlineTextNodes   []InlineTextBox `json:"inlineTextNodes,omitempty"`   // The post-layout inline text nodes, if any.
	StyleIndex        *int            `json:"styleIndex,omitempty"`        // Index into the `computedStyles` array returned by `getSnapshot`.
	PaintOrder        *int            `json:"paintOrder,omitempty"`        // Global paint order index, which is determined by the stacking order of the nodes. Nodes that are painted together will have the same index. Only provided if includePaintOrder in getSnapshot was true.
	IsStackingContext *bool           `json:"isStackingContext,omitempty"` // Set to true to indicate the element begins a new stacking context.
}

// ComputedStyle A subset of the full ComputedStyle as defined by the request
// whitelist.
type ComputedStyle struct {
	Properties []NameValue `json:"properties"` // Name/value pairs of computed style properties.
}

// NameValue A name/value pair.
type NameValue struct {
	Name  string `json:"name"`  // Attribute/property name.
	Value string `json:"value"` // Attribute/property value.
}

// StringIndex Index of the string in the strings table.
type StringIndex int

// ArrayOfStrings Index of the string in the strings table.
type ArrayOfStrings []StringIndex

// RareStringData Data that is only present on rare nodes.
type RareStringData struct {
	Index []int         `json:"index"` // No description.
	Value []StringIndex `json:"value"` // No description.
}

// RareBooleanData
type RareBooleanData struct {
	Index []int `json:"index"` // No description.
}

// RareIntegerData
type RareIntegerData struct {
	Index []int `json:"index"` // No description.
	Value []int `json:"value"` // No description.
}

// Rectangle
type Rectangle []float64

// DocumentSnapshot Document snapshot.
type DocumentSnapshot struct {
	DocumentURL     StringIndex        `json:"documentURL"`     // Document URL that `Document` or `FrameOwner` node points to.
	BaseURL         StringIndex        `json:"baseURL"`         // Base URL that `Document` or `FrameOwner` node uses for URL completion.
	ContentLanguage StringIndex        `json:"contentLanguage"` // Contains the document's content language.
	EncodingName    StringIndex        `json:"encodingName"`    // Contains the document's character set encoding.
	PublicID        StringIndex        `json:"publicId"`        // `DocumentType` node's publicId.
	SystemID        StringIndex        `json:"systemId"`        // `DocumentType` node's systemId.
	FrameID         StringIndex        `json:"frameId"`         // Frame ID for frame owner elements and also for the document node.
	Nodes           NodeTreeSnapshot   `json:"nodes"`           // A table with dom nodes.
	Layout          LayoutTreeSnapshot `json:"layout"`          // The nodes in the layout tree.
	TextBoxes       TextBoxSnapshot    `json:"textBoxes"`       // The post-layout inline text nodes.
}

// NodeTreeSnapshot Table containing nodes.
type NodeTreeSnapshot struct {
	ParentIndex          []int               `json:"parentIndex,omitempty"`          // Parent node index.
	NodeType             []int               `json:"nodeType,omitempty"`             // `Node`'s nodeType.
	NodeName             []StringIndex       `json:"nodeName,omitempty"`             // `Node`'s nodeName.
	NodeValue            []StringIndex       `json:"nodeValue,omitempty"`            // `Node`'s nodeValue.
	BackendNodeID        []dom.BackendNodeID `json:"backendNodeId,omitempty"`        // `Node`'s id, corresponds to DOM.Node.backendNodeId.
	Attributes           []ArrayOfStrings    `json:"attributes,omitempty"`           // Attributes of an `Element` node. Flatten name, value pairs.
	TextValue            *RareStringData     `json:"textValue,omitempty"`            // Only set for textarea elements, contains the text value.
	InputValue           *RareStringData     `json:"inputValue,omitempty"`           // Only set for input elements, contains the input's associated text value.
	InputChecked         *RareBooleanData    `json:"inputChecked,omitempty"`         // Only set for radio and checkbox input elements, indicates if the element has been checked
	OptionSelected       *RareBooleanData    `json:"optionSelected,omitempty"`       // Only set for option elements, indicates if the element has been selected
	ContentDocumentIndex *RareIntegerData    `json:"contentDocumentIndex,omitempty"` // The index of the document in the list of the snapshot documents.
	PseudoType           *RareStringData     `json:"pseudoType,omitempty"`           // Type of a pseudo element node.
	IsClickable          *RareBooleanData    `json:"isClickable,omitempty"`          // Whether this DOM node responds to mouse clicks. This includes nodes that have had click event listeners attached via JavaScript as well as anchor tags that naturally navigate when clicked.
	CurrentSourceURL     *RareStringData     `json:"currentSourceURL,omitempty"`     // The selected url for nodes with a srcset attribute.
	OriginURL            *RareStringData     `json:"originURL,omitempty"`            // The url of the script (if any) that generates this node.
}

// LayoutTreeSnapshot Details of an element in the DOM tree with a
// LayoutObject.
type LayoutTreeSnapshot struct {
	NodeIndex        []int            `json:"nodeIndex"`        // The index of the related DOM node in the `domNodes` array returned by `getSnapshot`.
	Styles           []ArrayOfStrings `json:"styles"`           // Index into the `computedStyles` array returned by `captureSnapshot`.
	Bounds           []Rectangle      `json:"bounds"`           // The absolute position bounding box.
	Text             []StringIndex    `json:"text"`             // Contents of the LayoutText, if any.
	StackingContexts RareBooleanData  `json:"stackingContexts"` // Stacking context information.
}

// TextBoxSnapshot Details of post layout rendered text positions. The exact
// layout should not be regarded as stable and may change between versions.
type TextBoxSnapshot struct {
	LayoutIndex []int       `json:"layoutIndex"` // Intex of th elayout tree node that owns this box collection.
	Bounds      []Rectangle `json:"bounds"`      // The absolute position bounding box.
	Start       []int       `json:"start"`       // The starting index in characters, for this post layout textbox substring. Characters that would be represented as a surrogate pair in UTF-16 have length 2.
	Length      []int       `json:"length"`      // The number of characters in this post layout textbox substring. Characters that would be represented as a surrogate pair in UTF-16 have length 2.
}
