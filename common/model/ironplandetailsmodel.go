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
	ironPlanDetailsFieldNames          = builderx.RawFieldNames(&IronPlanDetails{})
	ironPlanDetailsRows                = strings.Join(ironPlanDetailsFieldNames, ",")
	ironPlanDetailsRowsExpectAutoSet   = strings.Join(stringx.Remove(ironPlanDetailsFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	ironPlanDetailsRowsWithPlaceHolder = strings.Join(stringx.Remove(ironPlanDetailsFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	IronPlanDetailsModel interface {
		Insert(data IronPlanDetails) (sql.Result, error)
		FindOne(id int64) (*IronPlanDetails, error)
		Update(data IronPlanDetails) error
		Delete(id int64) error
		BatchGetDetailsMap(pidS []int64) ([]*IronPlanDetails, error)
	}

	defaultIronPlanDetailsModel struct {
		conn  sqlx.SqlConn
		table string
	}

	IronPlanDetails struct {
		Id         int64     `db:"id"`
		PlanId     int64     `db:"plan_id"`     // 属于哪个计划
		MovementId int64     `db:"movement_id"` // 使用哪个动作
		Weight     int64     `db:"weight"`      // 重量
		Count      int64     `db:"count"`       // 次数
		Break      int64     `db:"break"`       // 间歇
		Status     int64     `db:"status"`      // 状态
		UserId     int64     `db:"user_id"`     // 用户
		CreatedAt  time.Time `db:"created_at"`
		UpdatedAt  time.Time `db:"updated_at"`
	}
)

func NewIronPlanDetailsModel(conn sqlx.SqlConn) IronPlanDetailsModel {
	return &defaultIronPlanDetailsModel{
		conn:  conn,
		table: "`iron_plan_details`",
	}
}

func (m *defaultIronPlanDetailsModel) Insert(data IronPlanDetails) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, ironPlanDetailsRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.PlanId, data.MovementId, data.Weight, data.Count, data.Break, data.Status, data.UserId, data.CreatedAt, data.UpdatedAt)
	return ret, err
}

func (m *defaultIronPlanDetailsModel) FindOne(id int64) (*IronPlanDetails, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", ironPlanDetailsRows, m.table)
	var resp IronPlanDetails
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

func (m *defaultIronPlanDetailsModel) Update(data IronPlanDetails) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, ironPlanDetailsRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.PlanId, data.MovementId, data.Weight, data.Count, data.Break, data.Status, data.UserId, data.CreatedAt, data.UpdatedAt, data.Id)
	return err
}

func (m *defaultIronPlanDetailsModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}

func (m *defaultIronPlanDetailsModel) BatchGetDetailsMap(pidS []int64) ([]*IronPlanDetails, error) {
	//qs := make()
	//query := fmt.Sprintf("select %s from %s where plan_id in (%s) and status = ?", ironPlanDetailsRows)
	return nil, nil
}
