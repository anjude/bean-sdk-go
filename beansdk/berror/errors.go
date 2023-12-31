package berror

import "github.com/anjude/bcore/pkg/berr"

type BizErr berr.BizError

var NewBizError = berr.NewBizError

var (
	InternalErr = NewBizError(10000, "internal err")
	ParamErr    = NewBizError(10001, "param err")
)
