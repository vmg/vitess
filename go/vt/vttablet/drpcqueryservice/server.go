package drpcqueryservice

import (
	"context"

	"storj.io/drpc"

	binlogdatapb "vitess.io/vitess/go/vt/proto/binlogdata"
	querypb "vitess.io/vitess/go/vt/proto/query"
	queryservicepb "vitess.io/vitess/go/vt/proto/queryservice"
	"vitess.io/vitess/go/vt/vttablet/queryservice"
)

type queryServer struct {
	server queryservice.QueryService
}

func (q *queryServer) Execute(ctx context.Context, request *querypb.ExecuteRequest) (*querypb.ExecuteResponse, error) {
	panic("implement me")
}

func (q *queryServer) ExecuteBatch(ctx context.Context, request *querypb.ExecuteBatchRequest) (*querypb.ExecuteBatchResponse, error) {
	panic("implement me")
}

func (q *queryServer) StreamExecute(request *querypb.StreamExecuteRequest, stream queryservicepb.DRPCQuery_StreamExecuteStream) error {
	panic("implement me")
}

func (q *queryServer) Begin(ctx context.Context, request *querypb.BeginRequest) (*querypb.BeginResponse, error) {
	panic("implement me")
}

func (q *queryServer) Commit(ctx context.Context, request *querypb.CommitRequest) (*querypb.CommitResponse, error) {
	panic("implement me")
}

func (q *queryServer) Rollback(ctx context.Context, request *querypb.RollbackRequest) (*querypb.RollbackResponse, error) {
	panic("implement me")
}

func (q *queryServer) Prepare(ctx context.Context, request *querypb.PrepareRequest) (*querypb.PrepareResponse, error) {
	panic("implement me")
}

func (q *queryServer) CommitPrepared(ctx context.Context, request *querypb.CommitPreparedRequest) (*querypb.CommitPreparedResponse, error) {
	panic("implement me")
}

func (q *queryServer) RollbackPrepared(ctx context.Context, request *querypb.RollbackPreparedRequest) (*querypb.RollbackPreparedResponse, error) {
	panic("implement me")
}

func (q *queryServer) CreateTransaction(ctx context.Context, request *querypb.CreateTransactionRequest) (*querypb.CreateTransactionResponse, error) {
	panic("implement me")
}

func (q *queryServer) StartCommit(ctx context.Context, request *querypb.StartCommitRequest) (*querypb.StartCommitResponse, error) {
	panic("implement me")
}

func (q *queryServer) SetRollback(ctx context.Context, request *querypb.SetRollbackRequest) (*querypb.SetRollbackResponse, error) {
	panic("implement me")
}

func (q *queryServer) ConcludeTransaction(ctx context.Context, request *querypb.ConcludeTransactionRequest) (*querypb.ConcludeTransactionResponse, error) {
	panic("implement me")
}

func (q *queryServer) ReadTransaction(ctx context.Context, request *querypb.ReadTransactionRequest) (*querypb.ReadTransactionResponse, error) {
	panic("implement me")
}

func (q *queryServer) BeginExecute(ctx context.Context, request *querypb.BeginExecuteRequest) (*querypb.BeginExecuteResponse, error) {
	panic("implement me")
}

func (q *queryServer) BeginExecuteBatch(ctx context.Context, request *querypb.BeginExecuteBatchRequest) (*querypb.BeginExecuteBatchResponse, error) {
	panic("implement me")
}

func (q *queryServer) MessageStream(request *querypb.MessageStreamRequest, stream queryservicepb.DRPCQuery_MessageStreamStream) error {
	panic("implement me")
}

func (q *queryServer) MessageAck(ctx context.Context, request *querypb.MessageAckRequest) (*querypb.MessageAckResponse, error) {
	panic("implement me")
}

func (q *queryServer) ReserveExecute(ctx context.Context, request *querypb.ReserveExecuteRequest) (*querypb.ReserveExecuteResponse, error) {
	panic("implement me")
}

func (q *queryServer) ReserveBeginExecute(ctx context.Context, request *querypb.ReserveBeginExecuteRequest) (*querypb.ReserveBeginExecuteResponse, error) {
	panic("implement me")
}

func (q *queryServer) Release(ctx context.Context, request *querypb.ReleaseRequest) (*querypb.ReleaseResponse, error) {
	panic("implement me")
}

func (q *queryServer) StreamHealth(request *querypb.StreamHealthRequest, stream queryservicepb.DRPCQuery_StreamHealthStream) error {
	panic("implement me")
}

func (q *queryServer) VStream(request *binlogdatapb.VStreamRequest, stream queryservicepb.DRPCQuery_VStreamStream) error {
	panic("implement me")
}

func (q *queryServer) VStreamRows(request *binlogdatapb.VStreamRowsRequest, stream queryservicepb.DRPCQuery_VStreamRowsStream) error {
	panic("implement me")
}

func (q *queryServer) VStreamResults(request *binlogdatapb.VStreamResultsRequest, stream queryservicepb.DRPCQuery_VStreamResultsStream) error {
	panic("implement me")
}

var _ queryservicepb.DRPCQueryServer = (*queryServer)(nil)

func Register(m drpc.Mux, server queryservice.QueryService) {
	queryservicepb.DRPCRegisterQuery(m, &queryServer{server: server})
}
