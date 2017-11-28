// Code generated by cdpgen. DO NOT EDIT.

// Package log implements the Log domain. Provides access to log entries.
package log

import (
	"context"

	"github.com/mafredri/cdp/protocol/internal"
	"github.com/mafredri/cdp/rpcc"
)

// domainClient is a client for the Log domain. Provides access to log entries.
type domainClient struct{ conn *rpcc.Conn }

// NewClient returns a client for the Log domain with the connection set to conn.
func NewClient(conn *rpcc.Conn) *domainClient {
	return &domainClient{conn: conn}
}

// Enable invokes the Log method. Enables log domain, sends the entries collected so far to the client by means of the entryAdded notification.
func (d *domainClient) Enable(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Log.enable", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Log", Op: "Enable", Err: err}
	}
	return
}

// Disable invokes the Log method. Disables log domain, prevents further log entries from being reported to the client.
func (d *domainClient) Disable(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Log.disable", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Log", Op: "Disable", Err: err}
	}
	return
}

// Clear invokes the Log method. Clears the log.
func (d *domainClient) Clear(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Log.clear", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Log", Op: "Clear", Err: err}
	}
	return
}

// StartViolationsReport invokes the Log method. start violation reporting.
func (d *domainClient) StartViolationsReport(ctx context.Context, args *StartViolationsReportArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Log.startViolationsReport", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Log.startViolationsReport", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Log", Op: "StartViolationsReport", Err: err}
	}
	return
}

// StopViolationsReport invokes the Log method. Stop violation reporting.
func (d *domainClient) StopViolationsReport(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Log.stopViolationsReport", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Log", Op: "StopViolationsReport", Err: err}
	}
	return
}

func (d *domainClient) EntryAdded(ctx context.Context) (EntryAddedClient, error) {
	s, err := rpcc.NewStream(ctx, "Log.entryAdded", d.conn)
	if err != nil {
		return nil, err
	}
	return &entryAddedClient{Stream: s}, nil
}

type entryAddedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *entryAddedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *entryAddedClient) Recv() (*EntryAddedReply, error) {
	event := new(EntryAddedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Log", Op: "EntryAdded Recv", Err: err}
	}
	return event, nil
}
