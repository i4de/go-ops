// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Task is the golang structure of table task for DAO operations like Where/Data.
type Task struct {
	g.Meta   `orm:"table:task, do:true"`
	Id       interface{} //
	TaskId   interface{} // 任务id
	Type     interface{} // 任务类型
	Content  interface{} // 任务内容
	Name     interface{} // 任务名称
	Reqid    interface{} // 请求id
	ParentId interface{} //
	Status   interface{} //
	Created  *gtime.Time //
	Updated  *gtime.Time //
	Creater  interface{} // 创建人
}
