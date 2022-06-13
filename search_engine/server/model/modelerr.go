package model

import (
	"errors"
)

var (
	ERREOR_USER_NOTEXISTS = errors.New("该用户不存在...")
	ERREOR_USER_EXISTS    = errors.New("该用户已经存在...")
	ERREOR_USER_PWD       = errors.New("密码不正确...")
)
