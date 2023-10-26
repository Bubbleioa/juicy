package logic

import (
	"context"

	"juicy/internal/svc"
	"juicy/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type JuicyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJuicyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JuicyLogic {
	return &JuicyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JuicyLogic) Juicy(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
