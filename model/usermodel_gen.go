// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userFieldNames          = builder.RawFieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, "`userName`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheUserUserNamePrefix = "cache:user:userName:"
)

type (
	userModel interface {
		Insert(ctx context.Context, data *User) (sql.Result, error)
		FindOne(ctx context.Context, userName string) (*User, error)
		Update(ctx context.Context, data *User) error
		Delete(ctx context.Context, userName string) error
	}

	defaultUserModel struct {
		sqlc.CachedConn
		table string
	}

	User struct {
		UserName string  `db:"userName"` // 用户名
		Pwd      string  `db:"pwd"`      // 密码
		Name     string  `db:"name"`
		Type     string  `db:"type"`
		Max      float64 `db:"Max"`
		Min      float64 `db:"Min"`
		Mail     string  `db:"mail"`
		Phone    string  `db:"phone"`
	}
)

func newUserModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUserModel {
	return &defaultUserModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`User`",
	}
}

func (m *defaultUserModel) Delete(ctx context.Context, userName string) error {
	userUserNameKey := fmt.Sprintf("%s%v", cacheUserUserNamePrefix, userName)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `userName` = ?", m.table)
		return conn.ExecCtx(ctx, query, userName)
	}, userUserNameKey)
	return err
}

func (m *defaultUserModel) FindOne(ctx context.Context, userName string) (*User, error) {
	userUserNameKey := fmt.Sprintf("%s%v", cacheUserUserNamePrefix, userName)
	var resp User
	err := m.QueryRowCtx(ctx, &resp, userUserNameKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `userName` = ? limit 1", userRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, userName)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) Insert(ctx context.Context, data *User) (sql.Result, error) {
	userUserNameKey := fmt.Sprintf("%s%v", cacheUserUserNamePrefix, data.UserName)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, userRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserName, data.Pwd, data.Name, data.Type, data.Max, data.Min, data.Mail, data.Phone)
	}, userUserNameKey)
	return ret, err
}

func (m *defaultUserModel) Update(ctx context.Context, data *User) error {
	userUserNameKey := fmt.Sprintf("%s%v", cacheUserUserNamePrefix, data.UserName)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `userName` = ?", m.table, userRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Pwd, data.Name, data.Type, data.Max, data.Min, data.Mail, data.Phone, data.UserName)
	}, userUserNameKey)
	return err
}

func (m *defaultUserModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheUserUserNamePrefix, primary)
}

func (m *defaultUserModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `userName` = ? limit 1", userRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserModel) tableName() string {
	return m.table
}
