package request

//分页查询请求
type PageQuery struct {
	PageIndex int `json:"pageIndex"form:"pageIndex"` // 页码
	PageSize  int `json:"pageSize"form:"pageSize"`   // 每页数量
}
