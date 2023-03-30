package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ShareModel = (*customShareModel)(nil)

type (
	// ShareModel is an interface to be customized, add more methods here,
	// and implement the added methods in customShareModel.
	ShareModel interface {
		shareModel
	}

	customShareModel struct {
		*defaultShareModel
	}
)

// NewShareModel returns a model for the database table.
func NewShareModel(conn sqlx.SqlConn, c cache.CacheConf) ShareModel {
	return &customShareModel{
		defaultShareModel: newShareModel(conn, c),
	}
}
