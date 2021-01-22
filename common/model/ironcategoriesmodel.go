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
	ironCategoriesFieldNames          = builderx.RawFieldNames(&IronCategories{})
	ironCategoriesRows                = strings.Join(ironCategoriesFieldNames, ",")
	ironCategoriesRowsExpectAutoSet   = strings.Join(stringx.Remove(ironCategoriesFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	ironCategoriesRowsWithPlaceHolder = strings.Join(stringx.Remove(ironCategoriesFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	IronCategoriesModel interface {
		Insert(data IronCategories) (sql.Result, error)
		FindOne(id int64) (*IronCategories, error)
		FindOneByName(name string) (*IronCategories, error)
		Update(data IronCategories) error
		Delete(id int64) error
		GetAll() ([]*IronCategories, error)
	}

	defaultIronCategoriesModel struct {
		conn  sqlx.SqlConn
		table string
	}

	IronCategories struct {
		Id          int64     `db:"id"`
		ParentId    int64     `db:"parent_id"`   // 父类id
		Description string    `db:"description"` // 类别描述
		Status      int64     `db:"status"`      // 类别状态
		UserId      int64     `db:"user_id"`     // 创建人
		CreatedAt   time.Time `db:"created_at"`  // 创建时间
		UpdatedAt   time.Time `db:"updated_at"`  // 修改时间
		Name        string    `db:"name"`        // 类别名称
	}
)

func NewIronCategoriesModel(conn sqlx.SqlConn) IronCategoriesModel {
	return &defaultIronCategoriesModel{
		conn:  conn,
		table: "`iron_categories`",
	}
}

func (m *defaultIronCategoriesModel) Insert(data IronCategories) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, ironCategoriesRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.ParentId, data.Description, data.Status, data.UserId, data.CreatedAt, data.UpdatedAt, data.Name)
	return ret, err
}

func (m *defaultIronCategoriesModel) FindOne(id int64) (*IronCategories, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", ironCategoriesRows, m.table)
	var resp IronCategories
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

func (m *defaultIronCategoriesModel) FindOneByName(name string) (*IronCategories, error) {
	var resp IronCategories
	query := fmt.Sprintf("select %s from %s where `name` = ? limit 1", ironCategoriesRows, m.table)
	err := m.conn.QueryRow(&resp, query, name)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultIronCategoriesModel) Update(data IronCategories) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, ironCategoriesRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.ParentId, data.Description, data.Status, data.UserId, data.CreatedAt, data.UpdatedAt, data.Name, data.Id)
	return err
}

func (m *defaultIronCategoriesModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}

func (m *defaultIronCategoriesModel) GetAll() ([]*IronCategories, error) {
	var resp = make([]*IronCategories, 0)
	query := fmt.Sprintf("select * from %s", m.table)
	err := m.conn.QueryRows(&resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
