// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// CheckTplDetail is the golang structure for table check_tpl_detail.
type CheckTplDetail struct {
	Id     int     `json:"id"     ` //
	Tid    string  `json:"tid"    ` // 模板id
	Cid    string  `json:"cid"    ` // 检查项id
	Sort   int     `json:"sort"   ` // 排序
	Weight float64 `json:"weight" ` // 权重
}
