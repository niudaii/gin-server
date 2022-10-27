package system

type Menu struct {
	MenuId    int    `json:"id" gorm:"not null;unique;primary_key"`
	ParentId  int    `json:"parentId"`
	Path      string `json:"path,omitempty"`
	Name      string `json:"name"`
	Meta      `json:"meta"`
	Redirect  string `json:"redirect,omitempty"`
	Component string `json:"component"`
}

type Meta struct {
	Title        string `json:"title"`
	Show         bool   `json:"show"`
	HideChildren bool   `json:"hideChildren,omitempty"`
}
