package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	ironUsersFieldNames          = builderx.RawFieldNames(&IronUsers{})
	ironUsersRows                = strings.Join(ironUsersFieldNames, ",")
	ironUsersRowsExpectAutoSet   = strings.Join(stringx.Remove(ironUsersFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	ironUsersRowsWithPlaceHolder = strings.Join(stringx.Remove(ironUsersFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	IronUsersModel interface {
		Insert(data IronUsers) (sql.Result, error)
		FindOne(id int64) (*IronUsers, error)
		FindOneByUserName(userName string) (*IronUsers, error)
		Update(data IronUsers) error
		Delete(id int64) error
	}

	defaultIronUsersModel struct {
		conn  sqlx.SqlConn
		table string
	}

	IronUsers struct {
		Id        int64     `db:"id"`
		UserName  string    `db:"user_name"` // 用户名
		Password  string    `db:"password"`
		NickName  string    `db:"nick_name"`
		Mobile    string    `db:"mobile"` // 手机号
		WxId      string    `db:"wx_id"`  // 微信号
		Status    int64     `db:"status"` // 状态
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
	}
)

func NewIronUsersModel(conn sqlx.SqlConn) IronUsersModel {
	return &defaultIronUsersModel{
		conn:  conn,
		table: "`iron_users`",
	}
}

func (m *defaultIronUsersModel) Insert(data IronUsers) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, ironUsersRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.UserName, data.Password, data.NickName, data.Mobile, data.WxId, data.Status, data.CreatedAt, data.UpdatedAt)
	return ret, err
}

func (m *defaultIronUsersModel) FindOne(id int64) (*IronUsers, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", ironUsersRows, m.table)
	var resp IronUsers
	err := m.conn.QueryRow(&resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultIronUsersModel) FindOneByUserName(userName string) (*IronUsers, error) {
	var resp IronUsers
	query := fmt.Sprintf("select %s from %s where `user_name` = ? limit 1", ironUsersRows, m.table)
	err := m.conn.QueryRow(&resp, query, userName)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultIronUsersModel) Update(data IronUsers) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, ironUsersRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.UserName, data.Password, data.NickName, data.Mobile, data.WxId, data.Status, data.CreatedAt, data.UpdatedAt, data.Id)
	return err
}

func (m *defaultIronUsersModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
