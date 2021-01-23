package error

import "errors"

var (
	ErrGrpcErr                = errors.New("grpc错误")
	ErrNoUid                  = errors.New("请提供uid")
	ErrNoCatId                = errors.New("请提供种类id")
	ErrQueryError             = errors.New("查询信息失败")
	ErrQueryEmpty             = errors.New("查询为空")
	ErrDeleteError            = errors.New("删除失败")
	ErrStatusInvalidError     = errors.New("状态非法")
	ErrNoCId                  = errors.New("请提供类别id")
	ErrNotFoundCate           = errors.New("未找到类别信息")
	ErrCateStatusInvalidError = errors.New("类别信息非法")
	ErrNameError              = errors.New("名字不能为空")
	ErrDuplicate              = errors.New("查询去重失败")
	ErrNoPlanDetails          = errors.New("未提供计划细节")
	ErrWeight                 = errors.New("重量非法")
	ErrCount                  = errors.New("次数非法")
	ErrBreak                  = errors.New("间歇非法")
)
