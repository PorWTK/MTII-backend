package services

import (
	"context"
	"fmt"
	"mtii-backend/dtos"
	"mtii-backend/entities"
	"mtii-backend/helpers"
	"mtii-backend/repositories"

	"gorm.io/gorm"
)

type IncomeService interface {
	GetAllIncome(ctx context.Context) ([]dtos.Income, error)
	GetIncomeByInvoiceIdNumber(ctx context.Context, incomeInvoiceIdNumber int) (dtos.Income, error)
	CreateIncome(ctx context.Context, req dtos.CreateIncomeRequest) (dtos.IncomeResponse, error)
	UpdateIncome(ctx context.Context, incomeInvoiceIdNumber int, req dtos.UpdateIncomeRequest) (dtos.IncomeResponse, error)
	DeleteIncome(ctx context.Context, incomeInvoiceIdNumber int) error
}

type incomeService struct {
	incomeRepository repositories.IncomeRepository
}

func NewIncomeService(
	incomeRepository repositories.IncomeRepository,
) IncomeService {
	return &incomeService{
		incomeRepository: incomeRepository,
	}
}

func (s *incomeService) GetAllIncome(ctx context.Context) ([]dtos.Income, error) {
	incomes, err := s.incomeRepository.GetAllIncome(ctx)
	if err != nil {
		return []dtos.Income{}, fmt.Errorf("failed to get income: %w", err)
	}

	var incomeDTOs []dtos.Income
	for _, i := range incomes {
		incomeDTOs = append(incomeDTOs, dtos.Income{
			QuotationIdNumber:          i.QuotationIdNumber,
			QuotationIssueDate:         i.QuotationIssueDate,
			QuotationDueDate:           i.QuotationDueDate,
			InvoiceIdNumber:            i.InvoiceIdNumber,
			InvoiceIssueDate:           i.InvoiceIssueDate,
			InvoiceDueDate:             i.InvoiceDueDate,
			ReceiptIssueDate:           i.ReceiptIssueDate,
			ReceiptIdNumber:            i.ReceiptIdNumber,
			AgencyTaxPayerIdNumber:     i.AgencyTaxPayerIdNumber,
			InfluencerPostingDate:      i.InfluencerPostingDate,
			AgencyAgencyName:           i.AgencyAgencyName,
			AgencyAddress:              i.AgencyAddress,
			AgencyPhoneNumber:          i.AgencyPhoneNumber,
			ContactorContactorName:     i.ContactorContactorName,
			ContactorPhoneNumber:       i.ContactorPhoneNumber,
			ContactorLine:              i.ContactorLine,
			ContactorEmail:             i.ContactorEmail,
			BrandBrandName:             i.BrandBrandName,
			BrandProduct:               i.BrandProduct,
			TransactionReferenceNumber: i.TransactionReferenceNumber,
			TermsAndConditions:         i.TermsAndConditions,
			TotalPaymentAmount:         i.TotalPaymentAmount,
			NotesForTheTotalPayment:    i.NotesForTheTotalPayment,
			FirstPayment:               i.FirstPayment,
			NotesForTheFirstPayment:    i.NotesForTheFirstPayment,
			SecondPayment:              i.SecondPayment,
			NotesForTheSecondPayment:   i.NotesForTheSecondPayment,
			UnpaidPaymentAmount:        i.UnpaidPaymentAmount,
			NotesForTheUnpaidPayment:   i.NotesForTheUnpaidPayment,
			Platform: dtos.Platform{
				Id:   i.Platform.Id,
				Name: i.Platform.Name,
			},
			Status: dtos.Status{
				Id:   i.Status.Id,
				Name: i.Status.Name,
			},
			PaymentMethod: dtos.PaymentMethod{
				Id:   i.PaymentMethod.Id,
				Name: i.PaymentMethod.Name,
			},
			Receiver: dtos.Receiver{
				Id:         i.Receiver.Id,
				Name:       i.Receiver.Name,
				Address:    i.Receiver.Address,
				Email:      i.Receiver.Email,
				Phone:      i.Receiver.Phone,
				TaxPayerId: i.Receiver.TaxPayerId,
			},
			SalePerson: dtos.SalePerson{
				Id:   i.SalePerson.Id,
				Name: i.SalePerson.Name,
			},
			Channel: dtos.Channel{
				Id:   i.Channel.Id,
				Name: i.Channel.Name,
			},
			Bank: dtos.Bank{
				Id:   i.Bank.Id,
				Name: i.Bank.Name,
			},
		})
	}

	if len(incomeDTOs) == 0 {
		return []dtos.Income{}, nil
	}

	return incomeDTOs, nil
}

func (s *incomeService) GetIncomeByInvoiceIdNumber(ctx context.Context, incomeId int) (dtos.Income, error) {
	income, err := s.incomeRepository.GetIncomeByInvoiceIdNumber(ctx, incomeId)
	if err != nil {
		return dtos.Income{}, fmt.Errorf("failed to get income: %w", err)
	}

	return dtos.Income{
		QuotationIdNumber:          income.QuotationIdNumber,
		QuotationIssueDate:         income.QuotationIssueDate,
		QuotationDueDate:           income.QuotationDueDate,
		InvoiceIdNumber:            income.InvoiceIdNumber,
		InvoiceIssueDate:           income.InvoiceIssueDate,
		InvoiceDueDate:             income.InvoiceDueDate,
		ReceiptIssueDate:           income.ReceiptIssueDate,
		ReceiptIdNumber:            income.ReceiptIdNumber,
		AgencyTaxPayerIdNumber:     income.AgencyTaxPayerIdNumber,
		InfluencerPostingDate:      income.InfluencerPostingDate,
		AgencyAgencyName:           income.AgencyAgencyName,
		AgencyAddress:              income.AgencyAddress,
		AgencyPhoneNumber:          income.AgencyPhoneNumber,
		ContactorContactorName:     income.ContactorContactorName,
		ContactorPhoneNumber:       income.ContactorPhoneNumber,
		ContactorLine:              income.ContactorLine,
		ContactorEmail:             income.ContactorEmail,
		BrandBrandName:             income.BrandBrandName,
		BrandProduct:               income.BrandProduct,
		TransactionReferenceNumber: income.TransactionReferenceNumber,
		TermsAndConditions:         income.TermsAndConditions,
		TotalPaymentAmount:         income.TotalPaymentAmount,
		NotesForTheTotalPayment:    income.NotesForTheTotalPayment,
		FirstPayment:               income.FirstPayment,
		NotesForTheFirstPayment:    income.NotesForTheFirstPayment,
		SecondPayment:              income.SecondPayment,
		NotesForTheSecondPayment:   income.NotesForTheSecondPayment,
		UnpaidPaymentAmount:        income.UnpaidPaymentAmount,
		NotesForTheUnpaidPayment:   income.NotesForTheUnpaidPayment,
		Platform: dtos.Platform{
			Id:   income.Platform.Id,
			Name: income.Platform.Name,
		},
		Status: dtos.Status{
			Id:   income.Status.Id,
			Name: income.Status.Name,
		},
		PaymentMethod: dtos.PaymentMethod{
			Id:   income.PaymentMethod.Id,
			Name: income.PaymentMethod.Name,
		},
		Receiver: dtos.Receiver{
			Id:         income.Receiver.Id,
			Name:       income.Receiver.Name,
			Address:    income.Receiver.Address,
			Email:      income.Receiver.Email,
			Phone:      income.Receiver.Phone,
			TaxPayerId: income.Receiver.TaxPayerId,
		},
		SalePerson: dtos.SalePerson{
			Id:   income.SalePerson.Id,
			Name: income.SalePerson.Name,
		},
		Channel: dtos.Channel{
			Id:   income.Channel.Id,
			Name: income.Channel.Name,
		},
		Bank: dtos.Bank{
			Id:   income.Bank.Id,
			Name: income.Bank.Name,
		},
	}, nil
}

func (s *incomeService) CreateIncome(ctx context.Context, req dtos.CreateIncomeRequest) (dtos.IncomeResponse, error) {
	_, err := s.incomeRepository.GetIncomeByInvoiceIdNumber(ctx, req.InvoiceIdNumber)
	if err == nil {
		return dtos.IncomeResponse{}, fmt.Errorf("invoice id number must be unique")
	} else if err != gorm.ErrRecordNotFound {
		return dtos.IncomeResponse{}, fmt.Errorf("failed to get income: %w", err)
	}

	data := entities.Income{
		QuotationIdNumber:          req.QuotationIdNumber,
		QuotationIssueDate:         req.QuotationIssueDate,
		QuotationDueDate:           req.QuotationDueDate,
		InvoiceIdNumber:            req.InvoiceIdNumber,
		InvoiceIssueDate:           req.InvoiceIssueDate,
		InvoiceDueDate:             req.InvoiceDueDate,
		ReceiptIssueDate:           req.ReceiptIssueDate,
		ReceiptIdNumber:            req.ReceiptIdNumber,
		AgencyTaxPayerIdNumber:     req.AgencyTaxPayerIdNumber,
		InfluencerPostingDate:      req.InfluencerPostingDate,
		AgencyAgencyName:           req.AgencyAgencyName,
		AgencyAddress:              req.AgencyAddress,
		AgencyPhoneNumber:          req.AgencyPhoneNumber,
		ContactorContactorName:     req.ContactorContactorName,
		ContactorPhoneNumber:       req.ContactorPhoneNumber,
		ContactorLine:              req.ContactorLine,
		ContactorEmail:             req.ContactorEmail,
		BrandBrandName:             req.BrandBrandName,
		BrandProduct:               req.BrandProduct,
		TransactionReferenceNumber: req.TransactionReferenceNumber,
		TermsAndConditions:         req.TermsAndConditions,
		TotalPaymentAmount:         req.TotalPaymentAmount,
		NotesForTheTotalPayment:    req.NotesForTheTotalPayment,
		FirstPayment:               req.FirstPayment,
		NotesForTheFirstPayment:    req.NotesForTheFirstPayment,
		SecondPayment:              req.SecondPayment,
		NotesForTheSecondPayment:   req.NotesForTheSecondPayment,
		UnpaidPaymentAmount:        req.UnpaidPaymentAmount,
		NotesForTheUnpaidPayment:   req.NotesForTheUnpaidPayment,
		PlatformId:                 req.PlatformId,
		StatusId:                   req.StatusId,
		PaymentMethodId:            req.PaymentMethodId,
		ReceiverId:                 req.ReceiverId,
		SalePersonId:               req.SalePersonId,
		ChannelId:                  req.ChannelId,
		BankId:                     req.BankId,
	}

	income, err := s.incomeRepository.CreateIncome(ctx, data)
	if err != nil {
		return dtos.IncomeResponse{}, fmt.Errorf("failed to save income: %w", err)
	}

	return dtos.IncomeResponse{
		InvoiceIdNumber: income.InvoiceIdNumber,
	}, nil
}

func (s *incomeService) UpdateIncome(ctx context.Context, incomeInvoiceIdNumber int, req dtos.UpdateIncomeRequest) (dtos.IncomeResponse, error) {
	income, err := s.incomeRepository.GetIncomeByInvoiceIdNumber(ctx, incomeInvoiceIdNumber)
	if err != nil {
		return dtos.IncomeResponse{}, fmt.Errorf("failed to get income: %w", err)
	}

	data := entities.Income{
		InvoiceIdNumber:            helpers.DefaultIfEmpty(req.InvoiceIdNumber, incomeInvoiceIdNumber),
		QuotationIdNumber:          helpers.DefaultIfEmpty(req.QuotationIdNumber, income.QuotationIdNumber),
		QuotationIssueDate:         helpers.DefaultIfEmpty(req.QuotationIssueDate, income.QuotationIssueDate),
		QuotationDueDate:           helpers.DefaultIfEmpty(req.QuotationDueDate, income.QuotationDueDate),
		InvoiceIssueDate:           helpers.DefaultIfEmpty(req.InvoiceIssueDate, income.InvoiceIssueDate),
		InvoiceDueDate:             helpers.DefaultIfEmpty(req.InvoiceDueDate, income.InvoiceDueDate),
		ReceiptIssueDate:           helpers.DefaultIfEmpty(req.ReceiptIssueDate, income.ReceiptIssueDate),
		ReceiptIdNumber:            helpers.DefaultIfEmpty(req.ReceiptIdNumber, income.ReceiptIdNumber),
		AgencyTaxPayerIdNumber:     helpers.DefaultIfEmpty(req.AgencyTaxPayerIdNumber, income.AgencyTaxPayerIdNumber),
		InfluencerPostingDate:      helpers.DefaultIfEmpty(req.InfluencerPostingDate, income.InfluencerPostingDate),
		AgencyAgencyName:           helpers.DefaultIfEmpty(req.AgencyAgencyName, income.AgencyAgencyName),
		AgencyAddress:              helpers.DefaultIfEmpty(req.AgencyAddress, income.AgencyAddress),
		AgencyPhoneNumber:          helpers.DefaultIfEmpty(req.AgencyPhoneNumber, income.AgencyPhoneNumber),
		ContactorContactorName:     helpers.DefaultIfEmpty(req.ContactorContactorName, income.ContactorContactorName),
		ContactorPhoneNumber:       helpers.DefaultIfEmpty(req.ContactorPhoneNumber, income.ContactorPhoneNumber),
		ContactorLine:              helpers.DefaultIfEmpty(req.ContactorLine, income.ContactorLine),
		ContactorEmail:             helpers.DefaultIfEmpty(req.ContactorEmail, income.ContactorEmail),
		BrandBrandName:             helpers.DefaultIfEmpty(req.BrandBrandName, income.BrandBrandName),
		BrandProduct:               helpers.DefaultIfEmpty(req.BrandProduct, income.BrandProduct),
		TransactionReferenceNumber: helpers.DefaultIfEmpty(req.TransactionReferenceNumber, income.TransactionReferenceNumber),
		TermsAndConditions:         helpers.DefaultIfEmpty(req.TermsAndConditions, income.TermsAndConditions),
		TotalPaymentAmount:         helpers.DefaultIfEmpty(req.TotalPaymentAmount, income.TotalPaymentAmount),
		NotesForTheTotalPayment:    helpers.DefaultIfEmpty(req.NotesForTheTotalPayment, income.NotesForTheTotalPayment),
		FirstPayment:               helpers.DefaultIfEmpty(req.FirstPayment, income.FirstPayment),
		NotesForTheFirstPayment:    helpers.DefaultIfEmpty(req.NotesForTheFirstPayment, income.NotesForTheFirstPayment),
		SecondPayment:              helpers.DefaultIfEmpty(req.SecondPayment, income.SecondPayment),
		NotesForTheSecondPayment:   helpers.DefaultIfEmpty(req.NotesForTheSecondPayment, income.NotesForTheSecondPayment),
		UnpaidPaymentAmount:        helpers.DefaultIfEmpty(req.UnpaidPaymentAmount, income.UnpaidPaymentAmount),
		NotesForTheUnpaidPayment:   helpers.DefaultIfEmpty(req.NotesForTheUnpaidPayment, income.NotesForTheUnpaidPayment),
		PlatformId:                 helpers.DefaultIfEmpty(req.PlatformId, income.PlatformId),
		StatusId:                   helpers.DefaultIfEmpty(req.StatusId, income.StatusId),
		PaymentMethodId:            helpers.DefaultIfEmpty(req.PaymentMethodId, income.PaymentMethodId),
		ReceiverId:                 helpers.DefaultIfEmpty(req.ReceiverId, income.ReceiverId),
		SalePersonId:               helpers.DefaultIfEmpty(req.SalePersonId, income.SalePersonId),
		ChannelId:                  helpers.DefaultIfEmpty(req.ChannelId, income.ChannelId),
		BankId:                     helpers.DefaultIfEmpty(req.BankId, income.BankId),
	}

	var updatedIncome entities.Income
	if data.InvoiceIdNumber != incomeInvoiceIdNumber {
		updatedIncome, err = s.incomeRepository.UpdateIncomeWithNewInvoiceIdNumber(ctx, data, incomeInvoiceIdNumber)
		if err != nil {
			return dtos.IncomeResponse{}, fmt.Errorf("failed to update income: %w", err)
		}
	} else {
		updatedIncome, err = s.incomeRepository.UpdateIncome(ctx, data)
		if err != nil {
			return dtos.IncomeResponse{}, fmt.Errorf("failed to update income: %w", err)
		}
	}

	return dtos.IncomeResponse{
		InvoiceIdNumber: updatedIncome.InvoiceIdNumber,
	}, nil
}

func (s *incomeService) DeleteIncome(ctx context.Context, incomeInvoiceIdNumber int) error {
	income, err := s.incomeRepository.GetIncomeByInvoiceIdNumber(ctx, incomeInvoiceIdNumber)
	if err != nil {
		return fmt.Errorf("failed to get income: %w", err)
	}

	err = s.incomeRepository.DeleteIncome(ctx, income.InvoiceIdNumber)
	if err != nil {
		return fmt.Errorf("failed to delete income: %w", err)
	}

	return nil
}
