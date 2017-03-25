package cdp_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/cdpcmd"
	"github.com/mafredri/cdp/rpcc"
)

func Example() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Connect to Chrome Debugging Protocol target (webSocketDebuggerUrl).
	wsURL := "ws://localhost:9222/devtools/page/45a887ba-c92a-4cff-9194-d9398cc87e2c"
	conn, err := rpcc.DialContext(ctx, wsURL)
	if err != nil {
		panic(err)
	}
	defer conn.Close() // Must be closed when we are done.

	// Create a new CDP Client that uses conn.
	c := cdp.NewClient(conn)

	// Enable events on the Page domain.
	if err = c.Page.Enable(ctx); err != nil {
		panic(err)
	}

	// New DOMContentEventFired client will receive and buffer
	// ContentEventFired events from now on.
	domContentEventFired, err := c.Page.DOMContentEventFired(ctx)
	if err != nil {
		panic(err)
	}
	defer domContentEventFired.Close()

	// Create the Navigate arguments with the optional Referrer field set.
	navArgs := cdpcmd.NewPageNavigateArgs("https://www.google.com").SetReferrer("https://duckduckgo.com")
	nav, err := c.Page.Navigate(ctx, navArgs)
	if err != nil {
		panic(err)
	}

	// Block until a DOM ContentEventFired event is triggered.
	if _, err = domContentEventFired.Recv(); err != nil {
		panic(err)
	}

	fmt.Printf("Page loaded with frame ID: %s\n", nav.FrameID)

	// Fetch the document root node. We can pass nil here
	// since this method only takes optional arguments.
	doc, err := c.DOM.GetDocument(ctx, nil)
	if err != nil {
		panic(err)
	}

	// Get the outer HTML for the page.
	result, err := c.DOM.GetOuterHTML(ctx, cdpcmd.NewDOMGetOuterHTMLArgs(doc.Root.NodeID))
	if err != nil {
		panic(err)
	}

	fmt.Printf("HTML: %s\n", result.OuterHTML)

	// Capture a screenshot of the current page.
	screenshotName := "screenshot.jpg"
	screenshot, err := c.Page.CaptureScreenshot(ctx, cdpcmd.NewPageCaptureScreenshotArgs().SetFormat("jpeg").SetQuality(80))
	if err != nil {
		panic(err)
	}
	if err = ioutil.WriteFile(screenshotName, screenshot.Data, 0644); err != nil {
		panic(err)
	}
	fmt.Printf("Saved screenshot: %s\n", screenshotName)
}
