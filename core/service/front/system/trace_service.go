package system

import (
	"equity/core/dao/systemdao"
	"equity/core/model/mchmodel"
	"equity/core/model/systemmodel"
	"equity/core/model/systemmodel/request"
	"equity/core/service"
	"github.com/gin-gonic/gin"
	"time"
)

// TraceService 调用记录
type TraceService struct {
	systemTraceDao systemTracedao.SystemTraceDao
}

// CreateOpenApiTrace 记录调用错误信息
func (t *TraceService) CreateOpenApiTrace(c *gin.Context, reqParam, resParam, errMsg, rid, fullPath, appKey string) error {
	req := request.SystemTraceCreate{
		ReqParam: reqParam,
		ResParam: resParam,
		ErrMsg:   errMsg,
		Rid:      rid,
		FullPath: fullPath,
		AppKey:   appKey,
	}
	validateErr := req.Validate()
	if validateErr != nil {
		return validateErr
	}
	var mchId int32 = 0
	if req.AppKey != "" {
		mch, getMchError := getMch(req.AppKey)
		if getMchError == nil && mch != nil {
			mchId = int32(mch.Id)
		}
	}

	reqModel := &systemmodel.TSystemTrace{
		ReqParam:   req.ReqParam,
		ResParam:   req.ResParam,
		CreateTime: int32(time.Now().Unix()),
		ErrMsg:     req.ErrMsg,
		MchId:      mchId,
		Rid:        req.Rid,
		FullPath:   req.FullPath,
	}

	return t.systemTraceDao.SystemCreate(c, reqModel)
}

func getMch(appKey string) (*mchmodel.TMch, error) {
	mchService := service.AllServiceGroupApp.FrontMchServiceGroup
	mchDetail, err := mchService.GetMchDetail(appKey)
	if err != nil {
		return nil, err
	}
	return mchDetail, nil
}
