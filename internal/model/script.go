package model

import (
	"osp/pkg/message"
	"reflect"

	"osp/pkg/util"
)

type Script struct {
	Path     string            `json:"path"`     //工作路径
	Cmd      string            `json:"cmd"`      // 解释器
	Env      map[string]string `json:"env"`      // 环境变量
	Content  string            `json:"content"`  // 脚本内容
	ExecWay  ExecWay           `json:"execWay"`  // 执行方式
	FileHash string            `json:"filehash"` // 文件hash  从服务器上下载脚本需要
	User     string            `json:"user"`     // 特定用户
	Timeout  int               `json:"timeout"`  // 超时时间
	Args     []string          `json:"args"`
	Input    string            `json:"input"` // 输入内容
}

type ScriptJob struct {
	Jobid   string `json:"jobid"`
	Script  Script `json:"script"`
	RunMode string `json:"runMode"`
}

type ScriptJobCancel struct {
	Jobid string `json:"jobid"`
}

type ResCode string

const (
	CodeSuccess        ResCode = "SUECCES"         // 脚本成功执行，且正常退出码是0
	CodeFailed         ResCode = "FAILED"          // 脚本运行成功，但是退出码不是0
	CodeNotRun         ResCode = "NOT_RUN"         // 其它错误 脚本没有运行
	CodeTimeOut        ResCode = "TIME_OUT"        // 超时退出
	CodeCancel         ResCode = "CANCEL"          // 脚本被取消运行
	CodeDownloadFailed ResCode = "DOWNLOAD_FAILED" // 下载失败
	CodeHashFailed     ResCode = "HASHERR"         // hash 校验不通过
)

type ResCmd struct {
	Stdout   string  `json:"stdout"`
	Stderr   string  `json:"stderr"`
	Code     ResCode `json:"code"`
	Err      string  `json:"err"`      // 错误信息
	ExitCode int     `json:"exitCode"` // 退出码
}

type ResponseResCmd struct {
	Jobid  string `json:"jobid"`
	PeerId string `json:"peerId"`
	ResCmd ResCmd `json:"resCmd"`
}

type ExecWay int

type TaskInfo struct {
	Req    interface{} `json:"req"`    // 请求参数
	Status string      `json:"status"` // 任务状态
	Value  interface{} `json:"value"`
	Err    string      `json:"err"` // 错误信息
}

type GetTaskInfo struct {
	TaskId string `json:"taskid"`
}

const (
	ExecCmd        ExecWay = iota // 命令执行
	ExecContent                   // 脚本内容执行
	ExecScriptName                // 根据脚本名执行   脚本在本机上
	ExecURL                       // 从服务器上下载脚本执行
)

func init() {
	message.RegisterMessage(&message.MessageMeta{
		Type: reflect.TypeOf((*ScriptJob)(nil)).Elem(),
		ID:   uint32(util.StringHash("model.ScriptJob")),
	})

	message.RegisterMessage(&message.MessageMeta{
		Type: reflect.TypeOf((*ResponseResCmd)(nil)).Elem(),
		ID:   uint32(util.StringHash("model.ResponseResCmd")),
	})

	message.RegisterMessage(&message.MessageMeta{
		Type: reflect.TypeOf((*ScriptJobCancel)(nil)).Elem(),
		ID:   uint32(util.StringHash("model.ScriptJobCancel")),
	})

	message.RegisterMessage(&message.MessageMeta{
		Type: reflect.TypeOf((*TaskInfo)(nil)).Elem(),
		ID:   uint32(util.StringHash("model.TaskInfo")),
	})
	message.RegisterMessage(&message.MessageMeta{
		Type: reflect.TypeOf((*GetTaskInfo)(nil)).Elem(),
		ID:   uint32(util.StringHash("model.GetTaskInfo")),
	})
}
