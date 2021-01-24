package model

import (
	"database/sql"
	"fmt"
	"iron-go/plan/rpc/plan"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	ironTrainingsFieldNames          = builderx.RawFieldNames(&IronTrainings{})
	ironTrainingsRows                = strings.Join(ironTrainingsFieldNames, ",")
	ironTrainingsRowsExpectAutoSet   = strings.Join(stringx.Remove(ironTrainingsFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	ironTrainingsRowsWithPlaceHolder = strings.Join(stringx.Remove(ironTrainingsFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	IronTrainingsModel interface {
		Insert(data IronTrainings) (sql.Result, error)
		FindOne(id int64) (*IronTrainings, error)
		Update(data IronTrainings) error
		Delete(id int64) error
	}

	defaultIronTrainingsModel struct {
		conn          sqlx.SqlConn
		table         string
		PlanRpcClient plan.PlanClient
	}

	IronTrainings struct {
		Id           int64     `db:"id"`
		TrainingDate time.Time `db:"training_date"` // 训练日期
		PlanId       int64     `db:"plan_id"`       // 使用计划
		StartTime    time.Time `db:"start_time"`    // 开始时间
		EndTime      time.Time `db:"end_time"`      // 结束时间
		Description  string    `db:"description"`   // 训练小记
		Status       int64     `db:"status"`
		UserId       int64     `db:"user_id"` // 用户id
		CreatedAt    time.Time `db:"created_at"`
		UpdatedAt    time.Time `db:"updated_at"`
	}
)

func NewIronTrainingsModel(conn sqlx.SqlConn, planRpcClient plan.PlanClient) IronTrainingsModel {
	return &defaultIronTrainingsModel{
		conn:          conn,
		table:         "`iron_trainings`",
		PlanRpcClient: planRpcClient,
	}
}

func (m *defaultIronTrainingsModel) Insert(data IronTrainings) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, ironTrainingsRowsExpectAutoSet)
	m.conn.Transact(func(session sqlx.Session) error {
		ret, err := session.Exec(query, data.TrainingDate, data.PlanId, data.StartTime, data.EndTime, data.Description, data.Status, data.UserId, data.CreatedAt, data.UpdatedAt)
		if err != nil {
			return err
		}
		trainingId, err := ret.LastInsertId()
		if err != nil {
			return err
		}
		fmt.Println(trainingId)
		// 查询所有的计划细节id
		return nil
	})
	return nil, nil
}

func (m *defaultIronTrainingsModel) FindOne(id int64) (*IronTrainings, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", ironTrainingsRows, m.table)
	var resp IronTrainings
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

func (m *defaultIronTrainingsModel) Update(data IronTrainings) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, ironTrainingsRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.TrainingDate, data.PlanId, data.StartTime, data.EndTime, data.Description, data.Status, data.UserId, data.CreatedAt, data.UpdatedAt, data.Id)
	return err
}

func (m *defaultIronTrainingsModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
