package drpctabletconn

import (
	"context"
	"io"
	"net"
	"sync"

	"vitess.io/vitess/go/vt/callerid"

	"storj.io/drpc/drpcconn"

	"vitess.io/vitess/go/netutil"
	"vitess.io/vitess/go/sqltypes"
	"vitess.io/vitess/go/vt/grpcclient"
	binlogdatapb "vitess.io/vitess/go/vt/proto/binlogdata"
	querypb "vitess.io/vitess/go/vt/proto/query"
	queryservicepb "vitess.io/vitess/go/vt/proto/queryservice"
	topodatapb "vitess.io/vitess/go/vt/proto/topodata"
	"vitess.io/vitess/go/vt/vttablet/queryservice"
	"vitess.io/vitess/go/vt/vttablet/tabletconn"
)

func init() {
	tabletconn.RegisterDialer("drpc", DialTablet)
}

type dRPCQueryClient struct {
	// tablet is set at construction time, and never changed
	tablet *topodatapb.Tablet

	// mu protects the next fields
	mu sync.RWMutex
	cc *drpcconn.Conn
	c  queryservicepb.DRPCQueryClient
}

func (conn *dRPCQueryClient) Begin(ctx context.Context, target *querypb.Target, options *querypb.ExecuteOptions) (int64, *topodatapb.TabletAlias, error) {
	panic("implement me")
}

func (conn *dRPCQueryClient) Commit(ctx context.Context, target *querypb.Target, transactionID int64) (int64, error) {
	panic("implement me")
}

func (conn *dRPCQueryClient) Rollback(ctx context.Context, target *querypb.Target, transactionID int64) (int64, error) {
	panic("implement me")
}

func (conn *dRPCQueryClient) Prepare(ctx context.Context, target *querypb.Target, transactionID int64, dtid string) (err error) {
	panic("implement me")
}

func (conn *dRPCQueryClient) CommitPrepared(ctx context.Context, target *querypb.Target, dtid string) (err error) {
	panic("implement me")
}

func (conn *dRPCQueryClient) RollbackPrepared(ctx context.Context, target *querypb.Target, dtid string, originalID int64) (err error) {
	panic("implement me")
}

func (conn *dRPCQueryClient) CreateTransaction(ctx context.Context, target *querypb.Target, dtid string, participants []*querypb.Target) (err error) {
	panic("implement me")
}

func (conn *dRPCQueryClient) StartCommit(ctx context.Context, target *querypb.Target, transactionID int64, dtid string) (err error) {
	panic("implement me")
}

func (conn *dRPCQueryClient) SetRollback(ctx context.Context, target *querypb.Target, dtid string, transactionID int64) (err error) {
	panic("implement me")
}

func (conn *dRPCQueryClient) ConcludeTransaction(ctx context.Context, target *querypb.Target, dtid string) (err error) {
	panic("implement me")
}

func (conn *dRPCQueryClient) ReadTransaction(ctx context.Context, target *querypb.Target, dtid string) (metadata *querypb.TransactionMetadata, err error) {
	panic("implement me")
}

func (conn *dRPCQueryClient) Execute(ctx context.Context, target *querypb.Target, sql string, bindVariables map[string]*querypb.BindVariable, transactionID, reservedID int64, options *querypb.ExecuteOptions) (*sqltypes.Result, error) {
	panic("implement me")
}

func (conn *dRPCQueryClient) StreamExecute(ctx context.Context, target *querypb.Target, sql string, bindVariables map[string]*querypb.BindVariable, transactionID int64, options *querypb.ExecuteOptions, callback func(*sqltypes.Result) error) error {
	panic("implement me")
}

func (conn *dRPCQueryClient) ExecuteBatch(ctx context.Context, target *querypb.Target, queries []*querypb.BoundQuery, asTransaction bool, transactionID int64, options *querypb.ExecuteOptions) ([]sqltypes.Result, error) {
	panic("implement me")
}

func (conn *dRPCQueryClient) BeginExecute(ctx context.Context, target *querypb.Target, preQueries []string, sql string, bindVariables map[string]*querypb.BindVariable, reservedID int64, options *querypb.ExecuteOptions) (*sqltypes.Result, int64, *topodatapb.TabletAlias, error) {
	panic("implement me")
}

func (conn *dRPCQueryClient) BeginExecuteBatch(ctx context.Context, target *querypb.Target, queries []*querypb.BoundQuery, asTransaction bool, options *querypb.ExecuteOptions) ([]sqltypes.Result, int64, *topodatapb.TabletAlias, error) {
	panic("implement me")
}

func (conn *dRPCQueryClient) MessageStream(ctx context.Context, target *querypb.Target, name string, callback func(*sqltypes.Result) error) error {
	panic("implement me")
}

func (conn *dRPCQueryClient) MessageAck(ctx context.Context, target *querypb.Target, name string, ids []*querypb.Value) (count int64, err error) {
	panic("implement me")
}

func (conn *dRPCQueryClient) VStream(ctx context.Context, target *querypb.Target, startPos string, tableLastPKs []*binlogdatapb.TableLastPK, filter *binlogdatapb.Filter, send func([]*binlogdatapb.VEvent) error) error {
	stream, err := func() (queryservicepb.DRPCQuery_VStreamClient, error) {
		conn.mu.RLock()
		defer conn.mu.RUnlock()
		if conn.cc == nil {
			return nil, tabletconn.ConnClosed
		}

		req := &binlogdatapb.VStreamRequest{
			Target:            target,
			EffectiveCallerId: callerid.EffectiveCallerIDFromContext(ctx),
			ImmediateCallerId: callerid.ImmediateCallerIDFromContext(ctx),
			Position:          startPos,
			Filter:            filter,
			TableLastPKs:      tableLastPKs,
		}
		stream, err := conn.c.VStream(ctx, req)
		if err != nil {
			return nil, tabletconn.ErrorFromDRPC(err)
		}
		return stream, nil
	}()
	if err != nil {
		return err
	}
	for {
		r, err := stream.Recv()
		if err != nil {
			return tabletconn.ErrorFromDRPC(err)
		}
		if ctx.Err() != nil {
			return ctx.Err()
		}
		if err := send(r.Events); err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}

func (conn *dRPCQueryClient) VStreamRows(ctx context.Context, target *querypb.Target, query string, lastpk *querypb.QueryResult, send func(*binlogdatapb.VStreamRowsResponse) error) error {
	stream, err := func() (queryservicepb.DRPCQuery_VStreamRowsClient, error) {
		conn.mu.RLock()
		defer conn.mu.RUnlock()
		if conn.cc == nil {
			return nil, tabletconn.ConnClosed
		}
		req := &binlogdatapb.VStreamRowsRequest{
			Target:            target,
			EffectiveCallerId: callerid.EffectiveCallerIDFromContext(ctx),
			ImmediateCallerId: callerid.ImmediateCallerIDFromContext(ctx),
			Query:             query,
			Lastpk:            lastpk,
		}
		stream, err := conn.c.VStreamRows(ctx, req)
		if err != nil {
			return nil, tabletconn.ErrorFromDRPC(err)
		}
		return stream, nil
	}()
	if err != nil {
		return err
	}
	for {
		r, err := stream.Recv()
		if err != nil {
			return tabletconn.ErrorFromDRPC(err)
		}
		if ctx.Err() != nil {
			return ctx.Err()
		}
		if err := send(r); err != nil {
			return err
		}
	}
}

func (conn *dRPCQueryClient) VStreamResults(ctx context.Context, target *querypb.Target, query string, send func(*binlogdatapb.VStreamResultsResponse) error) error {
	panic("implement me")
}

func (conn *dRPCQueryClient) StreamHealth(ctx context.Context, callback func(*querypb.StreamHealthResponse) error) error {
	panic("implement me")
}

func (conn *dRPCQueryClient) HandlePanic(err *error) {
	panic("implement me")
}

func (conn *dRPCQueryClient) ReserveBeginExecute(ctx context.Context, target *querypb.Target, preQueries []string, sql string, bindVariables map[string]*querypb.BindVariable, options *querypb.ExecuteOptions) (*sqltypes.Result, int64, int64, *topodatapb.TabletAlias, error) {
	panic("implement me")
}

func (conn *dRPCQueryClient) ReserveExecute(ctx context.Context, target *querypb.Target, preQueries []string, sql string, bindVariables map[string]*querypb.BindVariable, transactionID int64, options *querypb.ExecuteOptions) (*sqltypes.Result, int64, *topodatapb.TabletAlias, error) {
	panic("implement me")
}

func (conn *dRPCQueryClient) Release(ctx context.Context, target *querypb.Target, transactionID, reservedID int64) error {
	panic("implement me")
}

func (conn *dRPCQueryClient) Close(_ context.Context) error {
	return conn.cc.Close()
}

func DialTablet(tablet *topodatapb.Tablet, _ grpcclient.FailFast) (queryservice.QueryService, error) {
	var addr string
	if grpcPort, ok := tablet.PortMap["drpc"]; ok {
		addr = netutil.JoinHostPort(tablet.Hostname, grpcPort)
	} else {
		addr = tablet.Hostname
	}

	rawconn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	// TODO: tls

	cc := drpcconn.New(rawconn)
	c := queryservicepb.NewDRPCQueryClient(cc)
	return &dRPCQueryClient{
		tablet: tablet,
		cc:     cc,
		c:      c,
	}, nil
}
