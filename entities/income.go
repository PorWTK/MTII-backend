package entities

import (
	"time"
)

type Income struct {
	QuotationIdNumber          int       `json:"quotation_id_number"`
	QuotationIssueDate         time.Time `gorm:"type:timestamp with time zone" json:"quotation_issue_date"`
	QuotationDueDate           time.Time `gorm:"type:timestamp with time zone" json:"quotation_due_date"`
	InvoiceIdNumber            int       `gorm:"primary_key;unique" json:"invoice_id_number"`
	InvoiceIssueDate           time.Time `gorm:"type:timestamp with time zone" json:"invoice_issue_date"`
	InvoiceDueDate             time.Time `gorm:"type:timestamp with time zone" json:"invoice_due_date"`
	ReceiptIssueDate           time.Time `gorm:"type:timestamp with time zone" json:"receipt_issue_date"`
	ReceiptIdNumber            int       `json:"receipt_id_number"`
	AgencyTaxPayerIdNumber     int       `json:"agency_tax_payer_id_number"`
	InfluencerPostingDate      time.Time `gorm:"type:timestamp with time zone" json:"influencer_posting_date"`
	AgencyAgencyName           string    `gorm:"type:varchar(255)" json:"agency_agency_name"`
	AgencyAddress              string    `gorm:"type:varchar(255)" json:"agency_address"`
	AgencyPhoneNumber          string    `gorm:"type:varchar(255)" json:"agency_phone_number"`
	ContactorContactorName     string    `gorm:"type:varchar(255)" json:"contactor_contactor_name"`
	ContactorPhoneNumber       string    `gorm:"type:varchar(255)" json:"contactor_phone_number"`
	ContactorLine              string    `gorm:"type:varchar(255)" json:"contactor_line"`
	ContactorEmail             string    `gorm:"type:varchar(255)" json:"contactor_email"`
	BrandBrandName             string    `gorm:"type:varchar(255)" json:"brand_brand_name"`
	BrandProduct               string    `gorm:"type:varchar(255)" json:"brand_product"`
	TransactionReferenceNumber int       `json:"transaction_reference_number"`
	TermsAndConditions         string    `gorm:"type:varchar(255)" json:"terms_and_conditions"`
	TotalPaymentAmount         int       `json:"total_payment_amount"`
	NotesForTheTotalPayment    string    `gorm:"type:varchar(255)" json:"notes_for_the_total_payment"`
	FirstPayment               int       `json:"first_payment"`
	NotesForTheFirstPayment    string    `gorm:"type:varchar(255)" json:"notes_for_the_first_payment"`
	SecondPayment              int       `json:"second_payment"`
	NotesForTheSecondPayment   string    `gorm:"type:varchar(255)" json:"notes_for_the_second_payment"`
	UnpaidPaymentAmount        int       `json:"unpaid_payment_amount"`
	NotesForTheUnpaidPayment   string    `gorm:"type:varchar(255)" json:"notes_for_the_unpaid_payment"`

	PlatformId      int           `json:"platform_id"`
	Platform        Platform      `gorm:"foreignKey:PlatformId" json:"-"`
	StatusId        int           `json:"status_id"`
	Status          Status        `gorm:"foreignKey:StatusId" json:"-"`
	PaymentMethodId int           `json:"payment_method_id"`
	PaymentMethod   PaymentMethod `gorm:"foreignKey:PaymentMethodId" json:"-"`
	ReceiverId      int           `json:"receiver_id"`
	Receiver        Receiver      `gorm:"foreignKey:ReceiverId" json:"-"`
	SalePersonId    int           `json:"sale_person_id"`
	SalePerson      SalePerson    `gorm:"foreignKey:SalePersonId" json:"-"`
	ChannelId       int           `json:"channel_id"`
	Channel         Channel       `gorm:"foreignKey:ChannelId" json:"-"`
	BankId          int           `json:"bank_id"`
	Bank            Bank          `gorm:"foreignKey:BankId" json:"-"`
}
