package exception

import (
	"fmt"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen"
)

// Rollback 如果发生错误则调用 tx.Rollback 回滚并返回嵌套的错误信息
func Rollback(tx *codegen.Tx, err error) error {
	if rollbackErr := tx.Rollback(); rollbackErr != nil {
		err = fmt.Errorf("%w: %v", err, rollbackErr)
	}
	return err
}
