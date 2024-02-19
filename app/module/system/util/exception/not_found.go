package exception

import "github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen"

func IsNoRecordError(err error) bool {
	if err == nil {
		return false
	}

	return codegen.IsNotFound(err)
}
