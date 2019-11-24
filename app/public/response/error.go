/*
@Time : 2019/3/7 11:31 AM
@Author : Tenlu
@File : error
@Software: GoLand
*/
package response

const (
	// common error

	//Success
	RequertSuccess = 0
	// token err
	MissingToken = 10
	InvalidToken = 11
	TokenExpired = 12

	ParamsError   = 10001
	MethodError   = 10002
	NotFoundError = 10003
	UnknownError  = 10004

	// db error
	DBError   = 11000
	DataEmpty = 11001
)

const (
	// common error
	//Success
	OK = "Success"
	// token err
	ErrMissingToken = "Missing token"
	ErrInvalidToken = "Invalid token"
	ErrTokenExpired = "Token expired"

	ErrParamsError   = "Request params error"
	ErrMethodError   = "Request method error"
	ErrNotFoundError = "Request not found"
	ErrUnknownError  = "Unknow error"

	// db error
	ErrDatabaseError = "Database error"
	ErrDataEmpty     = "Data empty"
)

// 对应的提示语
var errorMsg = map[int]string{
	RequertSuccess: OK,
	MissingToken:   ErrMissingToken,
	InvalidToken:   ErrInvalidToken,
	TokenExpired:   ErrTokenExpired,
	ParamsError:    ErrParamsError,
	MethodError:    ErrMethodError,
	NotFoundError:  ErrNotFoundError,
	UnknownError:   ErrUnknownError,

	// db error
	DBError:   ErrDatabaseError,
	DataEmpty: ErrDataEmpty,
}
