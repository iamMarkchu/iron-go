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
	ironMovementsFieldNames          = builderx.RawFieldNames(&IronMovements{})
	ironMovementsRows                = strings.Join(ironMovementsFieldNames, ",")
	ironMovementsRowsExpectAutoSet   = strings.Join(stringx.Remove(ironMovementsFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	ironMovementsRowsWithPlaceHolder = strings.Join(stringx.Remove(ironMovementsFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	IronMovementsModel interface {
		Insert(data IronMovements) (sql.Result, error)
		FindOne(id int64) (*IronMovements, error)
		FindOneByName(name string) (*IronMovements, error)
		Update(data IronMovements) error
		Delete(id int64) error
		GetList(cId int32) ([]*IronMovements, error)
	}

	defaultIronMovementsModel struct {
		conn  sqlx.SqlConn
		table string
	}

	IronMovements struct {
		Id          int64     `db:"id"`
		CatId       int64     `db:"cat_id"`      // 类别 eg：胸,肩,背,腿,手臂,核心，有氧等
		Name        string    `db:"name"`        // 动作名称 eg: 上斜哑铃卧推
		Description string    `db:"description"` // 动作简介
		Status      int64     `db:"status"`      // 状态
		UserId      int64     `db:"user_id"`     // 创建人
		CreatedAt   time.Time `db:"created_at"`  // 创建时间
		UpdatedAt   time.Time `db:"updated_at"`
	}
)

func NewIronMovementsModel(conn sqlx.SqlConn) IronMovementsModel {
	return &defaultIronMovementsModel{
		conn:  conn,
		table: "`iron_movements`",
	}
}

func (m *defaultIronMovementsModel) Insert(data IronMovements) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, ironMovementsRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.CatId, data.Name, data.Description, data.Status, data.UserId, data.CreatedAt, data.UpdatedAt)
	return ret, err
}

func (m *defaultIronMovementsModel) FindOne(id int64) (*IronMovements, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", ironMovementsRows, m.table)
	var resp IronMovements
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

func (m *defaultIronMovementsModel) Update(data IronMovements) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, ironMovementsRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.CatId, data.Name, data.Description, data.Status, data.UserId, data.CreatedAt, data.UpdatedAt, data.Id)
	return err
}

func (m *defaultIronMovementsModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}

func (m *defaultIronMovementsModel) FindOneByName(name string) (*IronMovements, error) {
	query := fmt.Sprintf("select %s from %s where `name` = ? limit 1", ironMovementsRows, m.table)
	var resp IronMovements
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

func (m *defaultIronMovementsModel) GetList(cId int32) ([]*IronMovements, error) {
	var resp = make([]*IronMovements, 0)
	query := fmt.Sprintf("select %s from %s where cat_id = ? and status = ?", ironMovementsRows, m.table)
	err := m.conn.QueryRows(&resp, query, cId, StatusOk)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
