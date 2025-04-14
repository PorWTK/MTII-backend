package entities

type Detail struct {
	Id                    int    `gorm:"primary_key;auto_increment" json:"id"`
	Description           string `gorm:"type:varchar(255)" json:"description"`
	Notes                 string `gorm:"type:varchar(255)" json:"notes"`
	Quantity              int    `json:"quantity"`
	UnitPrice             int    `json:"unit_price"`
	IncomeInvoiceIdNumber int    `json:"income_invoice_id_number"`
	Income                Income `gorm:"foreignKey:IncomeInvoiceIdNumber" json:"-"`
}
