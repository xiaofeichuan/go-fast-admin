package response

type PageResult struct {
	List       interface{} `json:"list"`       //数据
	TotalCount int64       `json:"totalCount"` //总条数
}
