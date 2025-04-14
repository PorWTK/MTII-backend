package dtos

type (
	Detail struct {
		Id          int    `json:"id"`
		Description string `json:"description"`
		Notes       string `json:"notes"`
		Quantity    int    `json:"quantity"`
		UnitPrice   int    `json:"unit_price"`
		Income      Income `json:"income"`
	}

	CreateDetailRequest struct {
		Description           string `json:"description" binding:"required"`
		Notes                 string `json:"notes" binding:"required"`
		Quantity              int    `json:"quantity" binding:"required"`
		UnitPrice             int    `json:"unit_price" binding:"required"`
		IncomeInvoiceIdNumber int    `json:"income_invoice_id_number" binding:"required"`
	}

	UpdateDetailRequest struct {
		Description           string `json:"description"`
		Notes                 string `json:"notes"`
		Quantity              int    `json:"quantity"`
		UnitPrice             int    `json:"unit_price"`
		IncomeInvoiceIdNumber int    `json:"income_invoice_id_number"`
	}

	DetailResponse struct {
		Id int `json:"id"`
	}
)
