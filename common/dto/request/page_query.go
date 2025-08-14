package request

// PageQuery 分页查询请求
type PageQuery struct {
	PageNum  int `form:"pageNum"`  // 页码
	PageSize int `form:"pageSize"` // 每页数量
}
