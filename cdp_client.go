// Code generated by cdpgen. DO NOT EDIT.

package cdp

import (
	"github.com/mafredri/cdp/protocol/accessibility"
	"github.com/mafredri/cdp/protocol/animation"
	"github.com/mafredri/cdp/protocol/applicationcache"
	"github.com/mafredri/cdp/protocol/audits"
	"github.com/mafredri/cdp/protocol/backgroundservice"
	"github.com/mafredri/cdp/protocol/browser"
	"github.com/mafredri/cdp/protocol/cachestorage"
	"github.com/mafredri/cdp/protocol/cast"
	"github.com/mafredri/cdp/protocol/console"
	"github.com/mafredri/cdp/protocol/css"
	"github.com/mafredri/cdp/protocol/database"
	"github.com/mafredri/cdp/protocol/debugger"
	"github.com/mafredri/cdp/protocol/deviceorientation"
	"github.com/mafredri/cdp/protocol/dom"
	"github.com/mafredri/cdp/protocol/domdebugger"
	"github.com/mafredri/cdp/protocol/domsnapshot"
	"github.com/mafredri/cdp/protocol/domstorage"
	"github.com/mafredri/cdp/protocol/emulation"
	"github.com/mafredri/cdp/protocol/fetch"
	"github.com/mafredri/cdp/protocol/headlessexperimental"
	"github.com/mafredri/cdp/protocol/heapprofiler"
	"github.com/mafredri/cdp/protocol/indexeddb"
	"github.com/mafredri/cdp/protocol/input"
	"github.com/mafredri/cdp/protocol/inspector"
	"github.com/mafredri/cdp/protocol/io"
	"github.com/mafredri/cdp/protocol/layertree"
	"github.com/mafredri/cdp/protocol/log"
	"github.com/mafredri/cdp/protocol/memory"
	"github.com/mafredri/cdp/protocol/network"
	"github.com/mafredri/cdp/protocol/overlay"
	"github.com/mafredri/cdp/protocol/page"
	"github.com/mafredri/cdp/protocol/performance"
	"github.com/mafredri/cdp/protocol/profiler"
	"github.com/mafredri/cdp/protocol/runtime"
	"github.com/mafredri/cdp/protocol/schema"
	"github.com/mafredri/cdp/protocol/security"
	"github.com/mafredri/cdp/protocol/serviceworker"
	"github.com/mafredri/cdp/protocol/storage"
	"github.com/mafredri/cdp/protocol/systeminfo"
	"github.com/mafredri/cdp/protocol/target"
	"github.com/mafredri/cdp/protocol/tethering"
	"github.com/mafredri/cdp/protocol/tracing"
	"github.com/mafredri/cdp/protocol/webaudio"
	"github.com/mafredri/cdp/rpcc"
)

// Client represents a Chrome DevTools Protocol client that can be used to
// invoke methods or listen to events in every CDP domain. The Client consumes
// a rpcc connection, used to invoke the methods.
type Client struct {
	Accessibility        Accessibility
	Animation            Animation
	ApplicationCache     ApplicationCache
	Audits               Audits
	BackgroundService    BackgroundService
	Browser              Browser
	CSS                  CSS
	CacheStorage         CacheStorage
	Cast                 Cast
	Console              Console
	DOM                  DOM
	DOMDebugger          DOMDebugger
	DOMSnapshot          DOMSnapshot
	DOMStorage           DOMStorage
	Database             Database
	Debugger             Debugger
	DeviceOrientation    DeviceOrientation
	Emulation            Emulation
	Fetch                Fetch
	HeadlessExperimental HeadlessExperimental
	HeapProfiler         HeapProfiler
	IO                   IO
	IndexedDB            IndexedDB
	Input                Input
	Inspector            Inspector
	LayerTree            LayerTree
	Log                  Log
	Memory               Memory
	Network              Network
	Overlay              Overlay
	Page                 Page
	Performance          Performance
	Profiler             Profiler
	Runtime              Runtime
	Schema               Schema
	Security             Security
	ServiceWorker        ServiceWorker
	Storage              Storage
	SystemInfo           SystemInfo
	Target               Target
	Tethering            Tethering
	Tracing              Tracing
	WebAudio             WebAudio
}

// NewClient returns a new Client that uses conn
// for communication with the debugging target.
func NewClient(conn *rpcc.Conn) *Client {
	return &Client{
		Accessibility:        accessibility.NewClient(conn),
		Animation:            animation.NewClient(conn),
		ApplicationCache:     applicationcache.NewClient(conn),
		Audits:               audits.NewClient(conn),
		BackgroundService:    backgroundservice.NewClient(conn),
		Browser:              browser.NewClient(conn),
		CSS:                  css.NewClient(conn),
		CacheStorage:         cachestorage.NewClient(conn),
		Cast:                 cast.NewClient(conn),
		Console:              console.NewClient(conn),
		DOM:                  dom.NewClient(conn),
		DOMDebugger:          domdebugger.NewClient(conn),
		DOMSnapshot:          domsnapshot.NewClient(conn),
		DOMStorage:           domstorage.NewClient(conn),
		Database:             database.NewClient(conn),
		Debugger:             debugger.NewClient(conn),
		DeviceOrientation:    deviceorientation.NewClient(conn),
		Emulation:            emulation.NewClient(conn),
		Fetch:                fetch.NewClient(conn),
		HeadlessExperimental: headlessexperimental.NewClient(conn),
		HeapProfiler:         heapprofiler.NewClient(conn),
		IO:                   io.NewClient(conn),
		IndexedDB:            indexeddb.NewClient(conn),
		Input:                input.NewClient(conn),
		Inspector:            inspector.NewClient(conn),
		LayerTree:            layertree.NewClient(conn),
		Log:                  log.NewClient(conn),
		Memory:               memory.NewClient(conn),
		Network:              network.NewClient(conn),
		Overlay:              overlay.NewClient(conn),
		Page:                 page.NewClient(conn),
		Performance:          performance.NewClient(conn),
		Profiler:             profiler.NewClient(conn),
		Runtime:              runtime.NewClient(conn),
		Schema:               schema.NewClient(conn),
		Security:             security.NewClient(conn),
		ServiceWorker:        serviceworker.NewClient(conn),
		Storage:              storage.NewClient(conn),
		SystemInfo:           systeminfo.NewClient(conn),
		Target:               target.NewClient(conn),
		Tethering:            tethering.NewClient(conn),
		Tracing:              tracing.NewClient(conn),
		WebAudio:             webaudio.NewClient(conn),
	}
}
