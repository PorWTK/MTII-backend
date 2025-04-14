package dtos

import "time"

type (
	Income struct {
		QuotationIdNumber          int       `json:"quotation_id_number"`
		QuotationIssueDate         time.Time `json:"quotation_issue_date"`
		QuotationDueDate           time.Time `json:"quotation_due_date"`
		InvoiceIdNumber            int       `json:"invoice_id_number"`
		InvoiceIssueDate           time.Time `json:"invoice_issue_date"`
		InvoiceDueDate             time.Time `json:"invoice_due_date"`
		ReceiptIssueDate           time.Time `json:"receipt_issue_date"`
		ReceiptIdNumber            int       `json:"receipt_id_number"`
		AgencyTaxPayerIdNumber     int       `json:"agency_tax_payer_id_number"`
		InfluencerPostingDate      time.Time `json:"influencer_posting_date"`
		AgencyAgencyName           string    `json:"agency_agency_name"`
		AgencyAddress              string    `json:"agency_address"`
		AgencyPhoneNumber          string    `json:"agency_phone_number"`
		ContactorContactorName     string    `json:"contactor_contactor_name"`
		ContactorPhoneNumber       string    `json:"contactor_phone_number"`
		ContactorLine              string    `json:"contactor_line"`
		ContactorEmail             string    `json:"contactor_email"`
		BrandBrandName             string    `json:"brand_brand_name"`
		BrandProduct               string    `json:"brand_product"`
		TransactionReferenceNumber int       `json:"transaction_reference_number"`
		TermsAndConditions         string    `json:"terms_and_conditions"`
		TotalPaymentAmount         int       `json:"total_payment_amount"`
		NotesForTheTotalPayment    string    `json:"notes_for_the_total_payment"`
		FirstPayment               int       `json:"first_payment"`
		NotesForTheFirstPayment    string    `json:"notes_for_the_first_payment"`
		SecondPayment              int       `json:"second_payment"`
		NotesForTheSecondPayment   string    `json:"notes_for_the_second_payment"`
		UnpaidPaymentAmount        int       `json:"unpaid_payment_amount"`
		NotesForTheUnpaidPayment   string    `json:"notes_for_the_unpaid_payment"`

		Platform      Platform      `json:"platform"`
		Status        Status        `json:"status"`
		PaymentMethod PaymentMethod `json:"payment_method"`
		Receiver      Receiver      `json:"receiver"`
		SalePerson    SalePerson    `json:"sale_person"`
		Channel       Channel       `json:"channel"`
		Bank          Bank          `json:"bank"`
	}

	CreateIncomeRequest struct {
		QuotationIdNumber          int       `json:"quotation_id_number" binding:"required"`
		QuotationIssueDate         time.Time `json:"quotation_issue_date" binding:"required"`
		QuotationDueDate           time.Time `json:"quotation_due_date" binding:"required"`
		InvoiceIdNumber            int       `json:"invoice_id_number" binding:"required"`
		InvoiceIssueDate           time.Time `json:"invoice_issue_date" binding:"required"`
		InvoiceDueDate             time.Time `json:"invoice_due_date" binding:"required"`
		ReceiptIssueDate           time.Time `json:"receipt_issue_date" binding:"required"`
		ReceiptIdNumber            int       `json:"receipt_id_number" binding:"required"`
		AgencyTaxPayerIdNumber     int       `json:"agency_tax_payer_id_number" binding:"required"`
		InfluencerPostingDate      time.Time `json:"influencer_posting_date" binding:"required"`
		AgencyAgencyName           string    `json:"agency_agency_name" binding:"required"`
		AgencyAddress              string    `json:"agency_address" binding:"required"`
		AgencyPhoneNumber          string    `json:"agency_phone_number" binding:"required"`
		ContactorContactorName     string    `json:"contactor_contactor_name" binding:"required"`
		ContactorPhoneNumber       string    `json:"contactor_phone_number" binding:"required"`
		ContactorLine              string    `json:"contactor_line" binding:"required"`
		ContactorEmail             string    `json:"contactor_email" binding:"required"`
		BrandBrandName             string    `json:"brand_brand_name" binding:"required"`
		BrandProduct               string    `json:"brand_product" binding:"required"`
		TransactionReferenceNumber int       `json:"transaction_reference_number" binding:"required"`
		TermsAndConditions         string    `json:"terms_and_conditions" binding:"required"`
		TotalPaymentAmount         int       `json:"total_payment_amount" binding:"required"`
		NotesForTheTotalPayment    string    `json:"notes_for_the_total_payment" binding:"required"`
		FirstPayment               int       `json:"first_payment" binding:"required"`
		NotesForTheFirstPayment    string    `json:"notes_for_the_first_payment" binding:"required"`
		SecondPayment              int       `json:"second_payment" binding:"required"`
		NotesForTheSecondPayment   string    `json:"notes_for_the_second_payment" binding:"required"`
		UnpaidPaymentAmount        int       `json:"unpaid_payment_amount"`
		NotesForTheUnpaidPayment   string    `json:"notes_for_the_unpaid_payment" binding:"required"`

		PlatformId      int `json:"platform_id" binding:"required"`
		StatusId        int `json:"status_id" binding:"required"`
		PaymentMethodId int `json:"payment_method_id" binding:"required"`
		ReceiverId      int `json:"receiver_id" binding:"required"`
		SalePersonId    int `json:"sale_person_id" binding:"required"`
		ChannelId       int `json:"channel_id" binding:"required"`
		BankId          int `json:"bank_id" binding:"required"`
	}

	UpdateIncomeRequest struct {
		QuotationIdNumber          int       `json:"quotation_id_number"`
		QuotationIssueDate         time.Time `json:"quotation_issue_date"`
		QuotationDueDate           time.Time `json:"quotation_due_date"`
		InvoiceIdNumber            int       `json:"invoice_id_number"`
		InvoiceIssueDate           time.Time `json:"invoice_issue_date"`
		InvoiceDueDate             time.Time `json:"invoice_due_date"`
		ReceiptIssueDate           time.Time `json:"receipt_issue_date"`
		ReceiptIdNumber            int       `json:"receipt_id_number"`
		AgencyTaxPayerIdNumber     int       `json:"agency_tax_payer_id_number"`
		InfluencerPostingDate      time.Time `json:"influencer_posting_date"`
		AgencyAgencyName           string    `json:"agency_agency_name"`
		AgencyAddress              string    `json:"agency_address"`
		AgencyPhoneNumber          string    `json:"agency_phone_number"`
		ContactorContactorName     string    `json:"contactor_contactor_name"`
		ContactorPhoneNumber       string    `json:"contactor_phone_number"`
		ContactorLine              string    `json:"contactor_line"`
		ContactorEmail             string    `json:"contactor_email"`
		BrandBrandName             string    `json:"brand_brand_name"`
		BrandProduct               string    `json:"brand_product"`
		TransactionReferenceNumber int       `json:"transaction_reference_number"`
		TermsAndConditions         string    `json:"terms_and_conditions"`
		TotalPaymentAmount         int       `json:"total_payment_amount"`
		NotesForTheTotalPayment    string    `json:"notes_for_the_total_payment"`
		FirstPayment               int       `json:"first_payment"`
		NotesForTheFirstPayment    string    `json:"notes_for_the_first_payment"`
		SecondPayment              int       `json:"second_payment"`
		NotesForTheSecondPayment   string    `json:"notes_for_the_second_payment"`
		UnpaidPaymentAmount        int       `json:"unpaid_payment_amount"`
		NotesForTheUnpaidPayment   string    `json:"notes_for_the_unpaid_payment"`

		PlatformId      int `json:"platform_id"`
		StatusId        int `json:"status_id"`
		PaymentMethodId int `json:"payment_method_id"`
		ReceiverId      int `json:"receiver_id"`
		SalePersonId    int `json:"sale_person_id"`
		ChannelId       int `json:"channel_id"`
		BankId          int `json:"bank_id"`
	}

	IncomeResponse struct {
		InvoiceIdNumber int `json:"invoice_id_number"`
	}
)
