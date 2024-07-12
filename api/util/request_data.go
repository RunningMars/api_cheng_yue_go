package util


type ReqeustQuery struct {
	Parameters map[string]interface{} `form:"-"` // 使用form标签忽略绑定
}