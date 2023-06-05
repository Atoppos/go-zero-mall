package logic

import (
	"context"
	"mall/service/order/api/internal/svc"
	"mall/service/order/api/internal/types"
	"mall/service/order/rpc/order"
	"mall/service/product/rpc/product"
	"strings"

	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateRequest) (resp *types.CreateResponse, err error) {
	// 获取 OrderRpc BuildTarget
	orderRpc, err := l.svcCtx.Config.OrderRpc.BuildTarget()
	if err != nil {
		return nil, status.Error(100, "订单创建异常")
	}
	orderRpcBusiServer := strings.Index(orderRpc, "?")
	// 获取 ProductRpc BuildTarget

	productRpc, err := l.svcCtx.Config.ProductRpc.BuildTarget()
	if err != nil {
		return nil, status.Error(100, "订单创建异常")
	}
	productRpcBusiServer := strings.Index(productRpc, "?")

	var dtmServer = "consul://127.0.0.1:18404/dtmservice"

	gid := dtmgrpc.MustGenGid(dtmServer)

	saga := dtmgrpc.NewSagaGrpc(dtmServer, gid).
		Add(orderRpc[:orderRpcBusiServer]+"/order.Order/Create", orderRpc[:orderRpcBusiServer]+"/order.Order/CreateRevert", &order.CreateRequest{
			Uid:    req.Uid,
			Pid:    req.Pid,
			Amount: req.Amount,
			Status: 0,
		}).
		Add(productRpc[:productRpcBusiServer]+"/product.Product/DecrStock", productRpc[:productRpcBusiServer]+"/product.Product/DecrStockRevert", &product.DecrStockRequest{
			Id:  req.Pid,
			Num: 1,
		})
	// 事务提交
	/*saga := dtmgrpc.NewSagaGrpc(dtmServer, gid).
		Add(orderRpc[:orderRpcBusiServer]+"/order.Order/Create/TransOutHeaderYes", orderRpc[:orderRpcBusiServer]+"/order.Order/CreateRevert/TransOutHeaderYes", &order.CreateRequest{
			Uid:    req.Uid,
			Pid:    req.Pid,
			Amount: req.Amount,
			Status: 0,
		}).
		Add(productRpc[:productRpcBusiServer]+"/product.Product/DecrStock/TransOutHeaderYes", productRpc[:productRpcBusiServer]+"/product.Product/DecrStockRevert/TransOutHeaderYes", &product.DecrStockRequest{
			Id:  req.Pid,
			Num: 1,
	    })
	saga.BranchHeaders = map[string]string{
		"token": "f0512db6-76d6-f25e-f344-a98cc3484d42",
	}*/
	err = saga.Submit()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &types.CreateResponse{}, nil
}
