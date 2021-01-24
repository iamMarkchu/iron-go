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
	ironTrainingLogsFieldNames          = builderx.RawFieldNames(&IronTrainingLogs{})
	ironTrainingLogsRows                = strings.Join(ironTrainingLogsFieldNames, ",")
	ironTrainingLogsRowsExpectAutoSet   = strings.Join(stringx.Remove(ironTrainingLogsFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	ironTrainingLogsRowsWithPlaceHolder = strings.Join(stringx.Remove(ironTrainingLogsFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	IronTrainingLogsModel interface {
		Insert(data IronTrainingLogs) (sql.Result, error)
		FindOne(id int64) (*IronTrainingLogs, error)
		Update(data IronTrainingLogs) error
		Delete(id int64) error
	}

	defaultIronTrainingLogsModel struct {
		conn  sqlx.SqlConn
		table string
	}

	IronTrainingLogs struct {
		Id           int64     `db:"id"`
		PlanDetailId int64     `db:"plan_detail_id"` // 计划详情id
		Done         int64     `db:"done"`           // 是否完成    done (0 为完成，2, 部分完成  3, 完成 4,失败)
		Status       int64     `db:"status"`         // 状态
		UserId       int64     `db:"user_id"`        // 用户
		CreatedAt    time.Time `db:"created_at"`     // 创建时间
		UpdatedAt    time.Time `db:"updated_at"`     // 修改时间
	}
)

func NewIronTrainingLogsModel(conn sqlx.SqlConn) IronTrainingLogsModel {
	return &defaultIronTrainingLogsModel{
		conn:  conn,
		table: "`iron_training_logs`",
	}
}

func (m *defaultIronTrainingLogsModel) Insert(data IronTrainingLogs) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, ironTrainingLogsRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.PlanDetailId, data.Done, data.Status, data.UserId, data.CreatedAt, data.UpdatedAt)
	return ret, err
}

func (m *defaultIronTrainingLogsModel) FindOne(id int64) (*IronTrainingLogs, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", ironTrainingLogsRows, m.table)
	var resp IronTrainingLogs
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

func (m *defaultIronTrainingLogsModel) Update(data IronTrainingLogs) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, ironTrainingLogsRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.PlanDetailId, data.Done, data.Status, data.UserId, data.CreatedAt, data.UpdatedAt, data.Id)
	return err
}

func (m *defaultIronTrainingLogsModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
