/**
 * @author 刘荣飞 yes@noxue.com
 * @date 2018/12/26 23:55
 */

package utils

import "noxue/config"

type Error string

func (e Error) Error() string {
	if config.Config.Debug {
		return string(e)
	}
	return string("server error")
}

func CheckErr(err error) {
	if err != nil {
		panic(Error(err.Error()))
	}
}

type ApiError struct {
	Status int
	Code   int
	Data   string
}

func CheckApiError(status, code int, err error) {
	if err == nil {
		return
	}
	panic(ApiError{
		Status: status,
		Code:   code,
		Data:   err.Error(),
	})
}

func (this *ApiError) Error() string {
	return this.Data
}
