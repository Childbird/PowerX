package dictionary

import (
	"PowerX/internal/model"
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDictionaryTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateDictionaryTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDictionaryTypeLogic {
	return &UpdateDictionaryTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateDictionaryTypeLogic) UpdateDictionaryType(req *types.UpdateDictionaryTypeRequest) (resp *types.UpdateDictionaryTypeReply, err error) {
	newModel := model.DataDictionaryType{
		Name:        req.Name,
		Description: req.Description,
	}
	if err := l.svcCtx.PowerX.DataDictionaryUserCase.PatchDataDictionaryType(l.ctx, req.Type, &newModel); err != nil {
		return nil, err
	}

	return &types.UpdateDictionaryTypeReply{
		DictionaryType: &types.DictionaryType{
			Type:        newModel.Type,
			Name:        newModel.Name,
			Description: newModel.Description,
		},
	}, nil
}
