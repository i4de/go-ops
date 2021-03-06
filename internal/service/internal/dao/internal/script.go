// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ScriptDao is the data access object for table script.
type ScriptDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns ScriptColumns // columns contains all the column names of Table for convenient usage.
}

// ScriptColumns defines and stores column names for table script.
type ScriptColumns struct {
	Id        string //
	Name      string // 命令名称
	Content   string // 脚本内容
	Args      string // 参数信息
	Desc      string // 描述信息
	Type      string // 脚本类型shell或者powershell
	OwnerType string // 命令拥有者类型
	OwnerUid  string // 拥有者uid
	Created   string //
	Updated   string //
	ScriptUid string //
	WaitTime  string // 脚本超时时间
	Cmd       string // 脚本解释器
	Ext       string // 脚本文件扩展名
}

//  scriptColumns holds the columns for table script.
var scriptColumns = ScriptColumns{
	Id:        "id",
	Name:      "name",
	Content:   "content",
	Args:      "args",
	Desc:      "desc",
	Type:      "type",
	OwnerType: "owner_type",
	OwnerUid:  "owner_uid",
	Created:   "created",
	Updated:   "updated",
	ScriptUid: "script_uid",
	WaitTime:  "wait_time",
	Cmd:       "cmd",
	Ext:       "ext",
}

// NewScriptDao creates and returns a new DAO object for table data access.
func NewScriptDao() *ScriptDao {
	return &ScriptDao{
		group:   "default",
		table:   "script",
		columns: scriptColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ScriptDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ScriptDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ScriptDao) Columns() ScriptColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ScriptDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ScriptDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ScriptDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
