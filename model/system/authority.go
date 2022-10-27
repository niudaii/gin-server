package system

type Authority struct {
	AuthorityId   string `json:"authorityId" gorm:"not null;unique;primary_key;"`
	AuthorityName string `json:"authorityName"`
	Menus         []Menu `json:"menus" gorm:"many2many:authority_menus"`
}
