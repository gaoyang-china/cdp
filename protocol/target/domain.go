// Code generated by cdpgen. DO NOT EDIT.

// Package target implements the Target domain. Supports additional targets discovery and allows to attach to them.
package target

import (
	"context"

	"github.com/mafredri/cdp/protocol/internal"
	"github.com/mafredri/cdp/rpcc"
)

// domainClient is a client for the Target domain. Supports additional targets discovery and allows to attach to them.
type domainClient struct{ conn *rpcc.Conn }

// NewClient returns a client for the Target domain with the connection set to conn.
func NewClient(conn *rpcc.Conn) *domainClient {
	return &domainClient{conn: conn}
}

// SetDiscoverTargets invokes the Target method. Controls whether to discover available targets and notify via targetCreated/targetInfoChanged/targetDestroyed events.
func (d *domainClient) SetDiscoverTargets(ctx context.Context, args *SetDiscoverTargetsArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Target.setDiscoverTargets", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Target.setDiscoverTargets", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Target", Op: "SetDiscoverTargets", Err: err}
	}
	return
}

// SetAutoAttach invokes the Target method. Controls whether to automatically attach to new targets which are considered to be related to this one. When turned on, attaches to all existing related targets as well. When turned off, automatically detaches from all currently attached targets.
func (d *domainClient) SetAutoAttach(ctx context.Context, args *SetAutoAttachArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Target.setAutoAttach", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Target.setAutoAttach", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Target", Op: "SetAutoAttach", Err: err}
	}
	return
}

// SetAttachToFrames invokes the Target method.
func (d *domainClient) SetAttachToFrames(ctx context.Context, args *SetAttachToFramesArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Target.setAttachToFrames", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Target.setAttachToFrames", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Target", Op: "SetAttachToFrames", Err: err}
	}
	return
}

// SetRemoteLocations invokes the Target method. Enables target discovery for the specified locations, when setDiscoverTargets was set to true.
func (d *domainClient) SetRemoteLocations(ctx context.Context, args *SetRemoteLocationsArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Target.setRemoteLocations", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Target.setRemoteLocations", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Target", Op: "SetRemoteLocations", Err: err}
	}
	return
}

// SendMessageToTarget invokes the Target method. Sends protocol message over session with given id.
func (d *domainClient) SendMessageToTarget(ctx context.Context, args *SendMessageToTargetArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Target.sendMessageToTarget", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Target.sendMessageToTarget", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Target", Op: "SendMessageToTarget", Err: err}
	}
	return
}

// GetTargetInfo invokes the Target method. Returns information about a target.
func (d *domainClient) GetTargetInfo(ctx context.Context, args *GetTargetInfoArgs) (reply *GetTargetInfoReply, err error) {
	reply = new(GetTargetInfoReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "Target.getTargetInfo", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Target.getTargetInfo", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Target", Op: "GetTargetInfo", Err: err}
	}
	return
}

// ActivateTarget invokes the Target method. Activates (focuses) the target.
func (d *domainClient) ActivateTarget(ctx context.Context, args *ActivateTargetArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Target.activateTarget", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Target.activateTarget", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Target", Op: "ActivateTarget", Err: err}
	}
	return
}

// CloseTarget invokes the Target method. Closes the target. If the target is a page that gets closed too.
func (d *domainClient) CloseTarget(ctx context.Context, args *CloseTargetArgs) (reply *CloseTargetReply, err error) {
	reply = new(CloseTargetReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "Target.closeTarget", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Target.closeTarget", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Target", Op: "CloseTarget", Err: err}
	}
	return
}

// AttachToTarget invokes the Target method. Attaches to the target with given id.
func (d *domainClient) AttachToTarget(ctx context.Context, args *AttachToTargetArgs) (reply *AttachToTargetReply, err error) {
	reply = new(AttachToTargetReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "Target.attachToTarget", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Target.attachToTarget", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Target", Op: "AttachToTarget", Err: err}
	}
	return
}

// DetachFromTarget invokes the Target method. Detaches session with given id.
func (d *domainClient) DetachFromTarget(ctx context.Context, args *DetachFromTargetArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Target.detachFromTarget", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Target.detachFromTarget", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Target", Op: "DetachFromTarget", Err: err}
	}
	return
}

// CreateBrowserContext invokes the Target method. Creates a new empty BrowserContext. Similar to an incognito profile but you can have more than one.
func (d *domainClient) CreateBrowserContext(ctx context.Context) (reply *CreateBrowserContextReply, err error) {
	reply = new(CreateBrowserContextReply)
	err = rpcc.Invoke(ctx, "Target.createBrowserContext", nil, reply, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Target", Op: "CreateBrowserContext", Err: err}
	}
	return
}

// DisposeBrowserContext invokes the Target method. Deletes a BrowserContext, will fail of any open page uses it.
func (d *domainClient) DisposeBrowserContext(ctx context.Context, args *DisposeBrowserContextArgs) (reply *DisposeBrowserContextReply, err error) {
	reply = new(DisposeBrowserContextReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "Target.disposeBrowserContext", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Target.disposeBrowserContext", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Target", Op: "DisposeBrowserContext", Err: err}
	}
	return
}

// CreateTarget invokes the Target method. Creates a new page.
func (d *domainClient) CreateTarget(ctx context.Context, args *CreateTargetArgs) (reply *CreateTargetReply, err error) {
	reply = new(CreateTargetReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "Target.createTarget", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Target.createTarget", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Target", Op: "CreateTarget", Err: err}
	}
	return
}

// GetTargets invokes the Target method. Retrieves a list of available targets.
func (d *domainClient) GetTargets(ctx context.Context) (reply *GetTargetsReply, err error) {
	reply = new(GetTargetsReply)
	err = rpcc.Invoke(ctx, "Target.getTargets", nil, reply, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Target", Op: "GetTargets", Err: err}
	}
	return
}

func (d *domainClient) TargetCreated(ctx context.Context) (CreatedClient, error) {
	s, err := rpcc.NewStream(ctx, "Target.targetCreated", d.conn)
	if err != nil {
		return nil, err
	}
	return &createdClient{Stream: s}, nil
}

type createdClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *createdClient) GetStream() rpcc.Stream { return c.Stream }

func (c *createdClient) Recv() (*CreatedReply, error) {
	event := new(CreatedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Target", Op: "TargetCreated Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) TargetInfoChanged(ctx context.Context) (InfoChangedClient, error) {
	s, err := rpcc.NewStream(ctx, "Target.targetInfoChanged", d.conn)
	if err != nil {
		return nil, err
	}
	return &infoChangedClient{Stream: s}, nil
}

type infoChangedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *infoChangedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *infoChangedClient) Recv() (*InfoChangedReply, error) {
	event := new(InfoChangedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Target", Op: "TargetInfoChanged Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) TargetDestroyed(ctx context.Context) (DestroyedClient, error) {
	s, err := rpcc.NewStream(ctx, "Target.targetDestroyed", d.conn)
	if err != nil {
		return nil, err
	}
	return &destroyedClient{Stream: s}, nil
}

type destroyedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *destroyedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *destroyedClient) Recv() (*DestroyedReply, error) {
	event := new(DestroyedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Target", Op: "TargetDestroyed Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) AttachedToTarget(ctx context.Context) (AttachedToTargetClient, error) {
	s, err := rpcc.NewStream(ctx, "Target.attachedToTarget", d.conn)
	if err != nil {
		return nil, err
	}
	return &attachedToTargetClient{Stream: s}, nil
}

type attachedToTargetClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *attachedToTargetClient) GetStream() rpcc.Stream { return c.Stream }

func (c *attachedToTargetClient) Recv() (*AttachedToTargetReply, error) {
	event := new(AttachedToTargetReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Target", Op: "AttachedToTarget Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) DetachedFromTarget(ctx context.Context) (DetachedFromTargetClient, error) {
	s, err := rpcc.NewStream(ctx, "Target.detachedFromTarget", d.conn)
	if err != nil {
		return nil, err
	}
	return &detachedFromTargetClient{Stream: s}, nil
}

type detachedFromTargetClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *detachedFromTargetClient) GetStream() rpcc.Stream { return c.Stream }

func (c *detachedFromTargetClient) Recv() (*DetachedFromTargetReply, error) {
	event := new(DetachedFromTargetReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Target", Op: "DetachedFromTarget Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) ReceivedMessageFromTarget(ctx context.Context) (ReceivedMessageFromTargetClient, error) {
	s, err := rpcc.NewStream(ctx, "Target.receivedMessageFromTarget", d.conn)
	if err != nil {
		return nil, err
	}
	return &receivedMessageFromTargetClient{Stream: s}, nil
}

type receivedMessageFromTargetClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *receivedMessageFromTargetClient) GetStream() rpcc.Stream { return c.Stream }

func (c *receivedMessageFromTargetClient) Recv() (*ReceivedMessageFromTargetReply, error) {
	event := new(ReceivedMessageFromTargetReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Target", Op: "ReceivedMessageFromTarget Recv", Err: err}
	}
	return event, nil
}
