package system

type Authority struct {
	AuthorityId   string `json:"-" gorm:"not null;unique;primary_key;"`
	AuthorityName string `json:"authorityName"`
	Menus         []Menu `json:"-" gorm:"many2many:authority_menus"`
}
