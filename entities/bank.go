package entities

type Bank struct {
	Id   int    `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(255)" json:"name"`
}
