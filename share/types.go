package share

// API API 配置数据结构
type API struct {
	Name    string        `json:"-"`
	Source  string        `json:"-"`
	Process string        `json:"process,omitempty"`
	Guard   string        `json:"guard,omitempty"`
	Default []interface{} `json:"default,omitempty"`
}

// Column 字段呈现方式
type Column struct {
	Label string `json:"label"`
	View  Render `json:"view,omitempty"`
	Edit  Render `json:"edit,omitempty"`
	Form  Render `json:"form,omitempty"`
}

// Filter 查询过滤器
type Filter struct {
	Label string `json:"label"`
	Bind  string `json:"bind,omitempty"`
	Input Render `json:"input,omitempty"`
}

// Page 页面
type Page struct {
	Primary string                 `json:"primary"`
	Layout  map[string]interface{} `json:"layout"`
	Actions map[string]Render      `json:"actions,omitempty"`
	Option  map[string]interface{} `json:"option,omitempty"`
}

// Render 组件渲染方式
type Render struct {
	Type       string                 `json:"type,omitempty"`
	Props      map[string]interface{} `json:"props,omitempty"`
	Components map[string]interface{} `json:"components,omitempty"`
}