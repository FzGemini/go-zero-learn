// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"hello01/internal/svc"
	"hello01/internal/types"
)

type Hello01Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHello01Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Hello01Logic {
	return &Hello01Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Hello01Logic) Hello01(req *types.Request) (resp *types.Response, err error) {
	// add your logic here and delete this line
	resp = &types.Response{
		Message: "hello " + req.Name,
	}
	return
}
