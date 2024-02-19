package service

import (
	"context"
	"github.com/pkg/errors"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen"
)

type BaseService struct {
}

func (m *BaseService) WithTx(c context.Context, client *codegen.Client, fn func(tx *codegen.Tx) error) error {
	tx, err := client.Tx(c)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			_ = tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			err = errors.Wrapf(err, "rolling back transaction: %v", rollbackErr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return errors.Wrapf(err, "committing transaction: %v", err)
	}
	return nil
}
