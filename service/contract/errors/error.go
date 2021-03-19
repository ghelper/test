package error
import "github.com/pkg/errors"

var (
	ErrNotData             error = newError(errNotData)
	ErrBalanceInsufficient error = newError(errBalanceInsufficient)
	ErrBalanceLess         error = newError(errBalanceLess)
	ErrRepeatRequire       error = newError(errRepeatRequire)
	ErrAmount              error = newError(errAmount)
	ErrAmountZero          error = newError(errAmountZero)
	ErrFee                 error = newError(errFee)
)

var (
	errorMap = map[code]string{
		success:                "success",
		defaultError:           "error",
		errNotData:             "not data",
		errBalanceInsufficient: "balance insufficient",
		errBalanceLess:         "the amount cannot be less than or equal to zero",
		errRepeatRequire:       "repeat require",
		errAmount:              "amount of error",
		errAmountZero:          "the amount cannot be less than 0",
		errFee:                 "fee of error",
	}
)

type code int

const (
	success                     = 0
	defaultError                = 1
	errNotData             code = 100100
	errBalanceInsufficient code = 100200
	errBalanceLess         code = 100300
	errRepeatRequire       code = 100400
	errAmount              code = 100500
	errAmountZero          code = 100600
	errFee                 code = 100700
)

type mErrors struct {
	Code    code
	Message string
}

func (e *mErrors) Error() string {
	return e.Message
}

func newError(code code) *mErrors {
	return &mErrors{
		Code:    code,
		Message: errorMap[code],
	}
}

func Error(err error) (int, string) {
	if err == nil {
		return success, errorMap[success]
	}
	switch e := errors.Cause(err).(type) {
	case *mErrors:
		return int(e.Code), e.Message
	}
	return defaultError, errorMap[defaultError]
}