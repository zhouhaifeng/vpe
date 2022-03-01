// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package vpe

import (
	"context"
	"fmt"
	"io"

	api "git.fd.io/govpp.git/api"
)

// RPCService defines RPC service  vpe.
type RPCService interface {
	AddNodeNext(ctx context.Context, in *AddNodeNext) (*AddNodeNextReply, error)
	Cli(ctx context.Context, in *Cli) (*CliReply, error)
	CliInband(ctx context.Context, in *CliInband) (*CliInbandReply, error)
	ControlPing(ctx context.Context, in *ControlPing) (*ControlPingReply, error)
	GetF64EndianValue(ctx context.Context, in *GetF64EndianValue) (*GetF64EndianValueReply, error)
	GetF64IncrementByOne(ctx context.Context, in *GetF64IncrementByOne) (*GetF64IncrementByOneReply, error)
	GetNextIndex(ctx context.Context, in *GetNextIndex) (*GetNextIndexReply, error)
	GetNodeGraph(ctx context.Context, in *GetNodeGraph) (*GetNodeGraphReply, error)
	GetNodeIndex(ctx context.Context, in *GetNodeIndex) (*GetNodeIndexReply, error)
	LogDump(ctx context.Context, in *LogDump) (RPCService_LogDumpClient, error)
	ShowThreads(ctx context.Context, in *ShowThreads) (*ShowThreadsReply, error)
	ShowVersion(ctx context.Context, in *ShowVersion) (*ShowVersionReply, error)
	ShowVpeSystemTime(ctx context.Context, in *ShowVpeSystemTime) (*ShowVpeSystemTimeReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) AddNodeNext(ctx context.Context, in *AddNodeNext) (*AddNodeNextReply, error) {
	out := new(AddNodeNextReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Cli(ctx context.Context, in *Cli) (*CliReply, error) {
	out := new(CliReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) CliInband(ctx context.Context, in *CliInband) (*CliInbandReply, error) {
	out := new(CliInbandReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ControlPing(ctx context.Context, in *ControlPing) (*ControlPingReply, error) {
	out := new(ControlPingReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetF64EndianValue(ctx context.Context, in *GetF64EndianValue) (*GetF64EndianValueReply, error) {
	out := new(GetF64EndianValueReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetF64IncrementByOne(ctx context.Context, in *GetF64IncrementByOne) (*GetF64IncrementByOneReply, error) {
	out := new(GetF64IncrementByOneReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetNextIndex(ctx context.Context, in *GetNextIndex) (*GetNextIndexReply, error) {
	out := new(GetNextIndexReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetNodeGraph(ctx context.Context, in *GetNodeGraph) (*GetNodeGraphReply, error) {
	out := new(GetNodeGraphReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetNodeIndex(ctx context.Context, in *GetNodeIndex) (*GetNodeIndexReply, error) {
	out := new(GetNodeIndexReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) LogDump(ctx context.Context, in *LogDump) (RPCService_LogDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_LogDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_LogDumpClient interface {
	Recv() (*LogDetails, error)
	api.Stream
}

type serviceClient_LogDumpClient struct {
	api.Stream
}

func (c *serviceClient_LogDumpClient) Recv() (*LogDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *LogDetails:
		return m, nil
	case *ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) ShowThreads(ctx context.Context, in *ShowThreads) (*ShowThreadsReply, error) {
	out := new(ShowThreadsReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ShowVersion(ctx context.Context, in *ShowVersion) (*ShowVersionReply, error) {
	out := new(ShowVersionReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ShowVpeSystemTime(ctx context.Context, in *ShowVpeSystemTime) (*ShowVpeSystemTimeReply, error) {
	out := new(ShowVpeSystemTimeReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
