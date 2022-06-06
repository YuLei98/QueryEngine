package models

type UserInfo struct {
	ID       uint   `gorm:"primarykey;" json:"id"`
	Identity string `gorm:"column:identity;type:varchar(255);" json:"identity"`
	Name     string `gorm:"column:name;type:varchar(40);" json:"name"`
	Password string `gorm:"column:password;type:varchar(40);" json:"password"`
}

func (*UserInfo) TableName() string {
	return "users"
}

type FavoriteList struct {
	// ID       uint
	Identity string `gorm:"column:identity; type:varchar(255);"`
	Name     string `gorm:"column:name; type:varchar(255);"`
}

func (*FavoriteList) TableName() string {
	return "favorite_list"
}

type FavoriteInfo struct {
	// ID         uint
	Identity   string `gorm:"column:identity; type:varchar(255);"`
	Name       string `gorm:"column:name; type:varchar(255);"`
	DocumentId int    `gorm:"column:doc_id; type:int(11);"`
}

func (*FavoriteInfo) TableName() string {
	return "favorite_info"
}
