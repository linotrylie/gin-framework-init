package response

// NodeTree 节点树形结构
type NodeTree struct {
	Id         int32       `json:"id"`
	Pid        int32       `json:"pid"`
	Name       string      `json:"name"`
	Icon       string      `json:"icon"`
	Url        string      `json:"url"`
	Type       int8        `json:"type"`
	Sort       int32       `json:"sort"`
	CreateTime int32       `json:"createTime"`
	UpdateTime int32       `json:"updateTime"`
	Children   []*NodeTree `json:"children"`
}

// UserNodeTree 用户拥有的节点树,button 和menu,dir分开存
type UserNodeTree struct {
	NodeTree  []*NodeTree `json:"nodeTree"`
	ButtonUrl []*string   `json:"buttonUrl"`
}

// NodeDetail 节点详情
type NodeDetail struct {
	Id         int32  `json:"id"`
	Pid        int32  `json:"pid"`
	Name       string `json:"name"`
	Icon       string `json:"icon"`
	Url        string `json:"url"`
	Type       int8   `json:"type"`
	Sort       int32  `json:"sort"`
	CreateTime int32  `json:"createTime"`
	UpdateTime int32  `json:"updateTime"`
}
