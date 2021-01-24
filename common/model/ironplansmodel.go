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
	ironPlansFieldNames          = builderx.RawFieldNames(&IronPlans{})
	ironPlansRows                = strings.Join(ironPlansFieldNames, ",")
	ironPlansRowsExpectAutoSet   = strings.Join(stringx.Remove(ironPlansFieldNames, "`id`", "`create_at`", "`update_time`"), ",")
	ironPlansRowsWithPlaceHolder = strings.Join(stringx.Remove(ironPlansFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	IronPlansModel interface {
		Insert(data IronPlans) (sql.Result, error)
		FindOne(id int64) (*IronPlans, error)
		Update(data IronPlans) error
		Delete(id int64) error
		Create(data IronPlans, details []IronPlanDetails) (int64, error)
		GetListByUid(uid uint64) ([]*IronPlans, error)
	}

	defaultIronPlansModel struct {
		conn     sqlx.SqlConn
		table    string
		subTable string
	}

	IronPlans struct {
		Id        int64     `db:"id"`
		PlanName  string    `db:"plan_name"`  // 计划名称
		Status    int64     `db:"status"`     // 状态
		UserId    int64     `db:"user_id"`    // 用户id
		CreatedAt time.Time `db:"created_at"` // 创建时间
		UpdatedAt time.Time `db:"updated_at"` // 修改时间
	}
)

func NewIronPlansModel(conn sqlx.SqlConn) IronPlansModel {
	return &defaultIronPlansModel{
		conn:     conn,
		table:    "`iron_plans`",
		subTable: "`iron_plan_details`",
	}
}

func (m *defaultIronPlansModel) Insert(data IronPlans) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, ironPlansRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.PlanName, data.Status, data.UserId, data.CreatedAt, data.UpdatedAt)
	return ret, err
}

func (m *defaultIronPlansModel) FindOne(id int64) (*IronPlans, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", ironPlansRows, m.table)
	var resp IronPlans
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

func (m *defaultIronPlansModel) Update(data IronPlans) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, ironPlansRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.PlanName, data.Status, data.UserId, data.CreatedAt, data.UpdatedAt, data.Id)
	return err
}

func (m *defaultIronPlansModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}

func (m *defaultIronPlansModel) Create(data IronPlans, details []IronPlanDetails) (int64, error) {
	planId := int64(0)
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, ironPlansRowsExpectAutoSet)
	query2 := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.subTable, ironPlanDetailsRowsExpectAutoSet)
	err := m.conn.Transact(func(session sqlx.Session) error {
		ret, err := session.Exec(query, data.PlanName, data.Status, data.UserId, data.CreatedAt, data.UpdatedAt)
		if err != nil {
			return err
		}
		planId, err = ret.LastInsertId()
		if err != nil {
			return err
		}
		for _, detail := range details {
			ret, err = session.Exec(query2, planId, detail.MovementId, detail.Weight, detail.Count, detail.Break, StatusOk, detail.UserId, detail.CreatedAt, detail.UpdatedAt)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return planId, err
}

func (m *defaultIronPlansModel) GetListByUid(uid uint64) ([]*IronPlans, error) {
	resp := make([]*IronPlans, 0)
	query := fmt.Sprintf("select %s from %s where user_id = ? and status = ?", ironPlansRows, m.table)
	err := m.conn.QueryRows(&resp, query, uid, StatusOk)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
