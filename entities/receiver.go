package entities

type Receiver struct {
	Id         int    `gorm:"primary_key;auto_increment" json:"id"`
	Name       string `gorm:"type:varchar(255)" json:"name"`
	Address    string `gorm:"type:varchar(255)" json:"address"`
	Email      string `gorm:"type:varchar(255)" json:"email"`
	Phone      string `gorm:"type:varchar(255)" json:"phone"`
	TaxPayerId string `gorm:"type:varchar(255)" json:"tax_payer_id"`
}
