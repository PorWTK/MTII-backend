package services

import (
	"context"
	"fmt"
	"mtii-backend/dtos"
	"mtii-backend/entities"
	"mtii-backend/helpers"
	"mtii-backend/repositories"
)

type DetailService interface {
	GetAllDetail(ctx context.Context) ([]dtos.Detail, error)
	GetDetailById(ctx context.Context, detailId int) (dtos.Detail, error)
	CreateDetail(ctx context.Context, req dtos.CreateDetailRequest) (dtos.DetailResponse, error)
	UpdateDetail(ctx context.Context, detailId int, req dtos.UpdateDetailRequest) (dtos.DetailResponse, error)
	DeleteDetail(ctx context.Context, detailId int) error
}

type detailService struct {
	detailRepository repositories.DetailRepository
}

func NewDetailService(
	detailRepository repositories.DetailRepository,
) DetailService {
	return &detailService{
		detailRepository: detailRepository,
	}
}

func (s *detailService) GetAllDetail(ctx context.Context) ([]dtos.Detail, error) {
	details, err := s.detailRepository.GetAllDetail(ctx)
	if err != nil {
		return []dtos.Detail{}, fmt.Errorf("failed to get detail: %w", err)
	}

	var detailDTOs []dtos.Detail
	for _, d := range details {
		detailDTOs = append(detailDTOs, dtos.Detail{
			Id:          d.Id,
			Description: d.Description,
			Notes:       d.Notes,
			Quantity:    d.Quantity,
			UnitPrice:   d.UnitPrice,
			Income: dtos.Income{
				QuotationIdNumber:          d.Income.QuotationIdNumber,
				QuotationIssueDate:         d.Income.QuotationIssueDate,
				QuotationDueDate:           d.Income.QuotationDueDate,
				InvoiceIdNumber:            d.Income.InvoiceIdNumber,
				InvoiceIssueDate:           d.Income.InvoiceIssueDate,
				InvoiceDueDate:             d.Income.InvoiceDueDate,
				ReceiptIssueDate:           d.Income.ReceiptIssueDate,
				ReceiptIdNumber:            d.Income.ReceiptIdNumber,
				AgencyTaxPayerIdNumber:     d.Income.AgencyTaxPayerIdNumber,
				InfluencerPostingDate:      d.Income.InfluencerPostingDate,
				AgencyAgencyName:           d.Income.AgencyAgencyName,
				AgencyAddress:              d.Income.AgencyAddress,
				AgencyPhoneNumber:          d.Income.AgencyPhoneNumber,
				ContactorContactorName:     d.Income.ContactorContactorName,
				ContactorPhoneNumber:       d.Income.ContactorPhoneNumber,
				ContactorLine:              d.Income.ContactorLine,
				ContactorEmail:             d.Income.ContactorEmail,
				BrandBrandName:             d.Income.BrandBrandName,
				BrandProduct:               d.Income.BrandProduct,
				TransactionReferenceNumber: d.Income.TransactionReferenceNumber,
				TermsAndConditions:         d.Income.TermsAndConditions,
				TotalPaymentAmount:         d.Income.TotalPaymentAmount,
				NotesForTheTotalPayment:    d.Income.NotesForTheTotalPayment,
				FirstPayment:               d.Income.FirstPayment,
				NotesForTheFirstPayment:    d.Income.NotesForTheFirstPayment,
				SecondPayment:              d.Income.SecondPayment,
				NotesForTheSecondPayment:   d.Income.NotesForTheSecondPayment,
				UnpaidPaymentAmount:        d.Income.UnpaidPaymentAmount,
				NotesForTheUnpaidPayment:   d.Income.NotesForTheUnpaidPayment,
				Platform: dtos.Platform{
					Id:   d.Income.Platform.Id,
					Name: d.Income.Platform.Name,
				},
				Status: dtos.Status{
					Id:   d.Income.Status.Id,
					Name: d.Income.Status.Name,
				},
				PaymentMethod: dtos.PaymentMethod{
					Id:   d.Income.PaymentMethod.Id,
					Name: d.Income.PaymentMethod.Name,
				},
				Receiver: dtos.Receiver{
					Id:         d.Income.Receiver.Id,
					Name:       d.Income.Receiver.Name,
					Address:    d.Income.Receiver.Address,
					Email:      d.Income.Receiver.Email,
					Phone:      d.Income.Receiver.Phone,
					TaxPayerId: d.Income.Receiver.TaxPayerId,
				},
				SalePerson: dtos.SalePerson{
					Id:   d.Income.SalePerson.Id,
					Name: d.Income.SalePerson.Name,
				},
				Channel: dtos.Channel{
					Id:   d.Income.Channel.Id,
					Name: d.Income.Channel.Name,
				},
				Bank: dtos.Bank{
					Id:   d.Income.Bank.Id,
					Name: d.Income.Bank.Name,
				},
			},
		})
	}

	if len(detailDTOs) == 0 {
		return []dtos.Detail{}, nil
	}

	return detailDTOs, nil
}

func (s *detailService) GetDetailById(ctx context.Context, detailId int) (dtos.Detail, error) {
	detail, err := s.detailRepository.GetDetailById(ctx, detailId)
	if err != nil {
		return dtos.Detail{}, fmt.Errorf("failed to get detail: %w", err)
	}

	return dtos.Detail{
		Id:          detail.Id,
		Description: detail.Description,
		Notes:       detail.Notes,
		Quantity:    detail.Quantity,
		UnitPrice:   detail.UnitPrice,
		Income: dtos.Income{
			QuotationIdNumber:          detail.Income.QuotationIdNumber,
			QuotationIssueDate:         detail.Income.QuotationIssueDate,
			QuotationDueDate:           detail.Income.QuotationDueDate,
			InvoiceIdNumber:            detail.Income.InvoiceIdNumber,
			InvoiceIssueDate:           detail.Income.InvoiceIssueDate,
			InvoiceDueDate:             detail.Income.InvoiceDueDate,
			ReceiptIssueDate:           detail.Income.ReceiptIssueDate,
			ReceiptIdNumber:            detail.Income.ReceiptIdNumber,
			AgencyTaxPayerIdNumber:     detail.Income.AgencyTaxPayerIdNumber,
			InfluencerPostingDate:      detail.Income.InfluencerPostingDate,
			AgencyAgencyName:           detail.Income.AgencyAgencyName,
			AgencyAddress:              detail.Income.AgencyAddress,
			AgencyPhoneNumber:          detail.Income.AgencyPhoneNumber,
			ContactorContactorName:     detail.Income.ContactorContactorName,
			ContactorPhoneNumber:       detail.Income.ContactorPhoneNumber,
			ContactorLine:              detail.Income.ContactorLine,
			ContactorEmail:             detail.Income.ContactorEmail,
			BrandBrandName:             detail.Income.BrandBrandName,
			BrandProduct:               detail.Income.BrandProduct,
			TransactionReferenceNumber: detail.Income.TransactionReferenceNumber,
			TermsAndConditions:         detail.Income.TermsAndConditions,
			TotalPaymentAmount:         detail.Income.TotalPaymentAmount,
			NotesForTheTotalPayment:    detail.Income.NotesForTheTotalPayment,
			FirstPayment:               detail.Income.FirstPayment,
			NotesForTheFirstPayment:    detail.Income.NotesForTheFirstPayment,
			SecondPayment:              detail.Income.SecondPayment,
			NotesForTheSecondPayment:   detail.Income.NotesForTheSecondPayment,
			UnpaidPaymentAmount:        detail.Income.UnpaidPaymentAmount,
			NotesForTheUnpaidPayment:   detail.Income.NotesForTheUnpaidPayment,
			Platform: dtos.Platform{
				Id:   detail.Income.Platform.Id,
				Name: detail.Income.Platform.Name,
			},
			Status: dtos.Status{
				Id:   detail.Income.Status.Id,
				Name: detail.Income.Status.Name,
			},
			PaymentMethod: dtos.PaymentMethod{
				Id:   detail.Income.PaymentMethod.Id,
				Name: detail.Income.PaymentMethod.Name,
			},
			Receiver: dtos.Receiver{
				Id:         detail.Income.Receiver.Id,
				Name:       detail.Income.Receiver.Name,
				Address:    detail.Income.Receiver.Address,
				Email:      detail.Income.Receiver.Email,
				Phone:      detail.Income.Receiver.Phone,
				TaxPayerId: detail.Income.Receiver.TaxPayerId,
			},
			SalePerson: dtos.SalePerson{
				Id:   detail.Income.SalePerson.Id,
				Name: detail.Income.SalePerson.Name,
			},
			Channel: dtos.Channel{
				Id:   detail.Income.Channel.Id,
				Name: detail.Income.Channel.Name,
			},
			Bank: dtos.Bank{
				Id:   detail.Income.Bank.Id,
				Name: detail.Income.Bank.Name,
			},
		},
	}, nil
}

func (s *detailService) CreateDetail(ctx context.Context, req dtos.CreateDetailRequest) (dtos.DetailResponse, error) {
	data := entities.Detail{
		Description:           req.Description,
		Notes:                 req.Notes,
		Quantity:              req.Quantity,
		UnitPrice:             req.UnitPrice,
		IncomeInvoiceIdNumber: req.IncomeInvoiceIdNumber,
	}

	detail, err := s.detailRepository.CreateDetail(ctx, data)
	if err != nil {
		return dtos.DetailResponse{}, fmt.Errorf("failed to save detail: %w", err)
	}

	return dtos.DetailResponse{
		Id: detail.Id,
	}, nil
}

func (s *detailService) UpdateDetail(ctx context.Context, detailId int, req dtos.UpdateDetailRequest) (dtos.DetailResponse, error) {
	detail, err := s.detailRepository.GetDetailById(ctx, detailId)
	if err != nil {
		return dtos.DetailResponse{}, fmt.Errorf("failed to get detail: %w", err)
	}

	data := entities.Detail{
		Id:                    detailId,
		Description:           helpers.DefaultIfEmpty(req.Description, detail.Description),
		Notes:                 helpers.DefaultIfEmpty(req.Notes, detail.Notes),
		Quantity:              helpers.DefaultIfEmpty(req.Quantity, detail.Quantity),
		UnitPrice:             helpers.DefaultIfEmpty(req.UnitPrice, detail.UnitPrice),
		IncomeInvoiceIdNumber: helpers.DefaultIfEmpty(req.IncomeInvoiceIdNumber, detail.IncomeInvoiceIdNumber),
	}

	updatedDetail, err := s.detailRepository.UpdateDetail(ctx, data)
	if err != nil {
		return dtos.DetailResponse{}, fmt.Errorf("failed to save detail: %w", err)
	}

	return dtos.DetailResponse{
		Id: updatedDetail.Id,
	}, nil
}

func (s *detailService) DeleteDetail(ctx context.Context, detailId int) error {
	detail, err := s.detailRepository.GetDetailById(ctx, detailId)
	if err != nil {
		return fmt.Errorf("failed to get detail: %w", err)
	}

	err = s.detailRepository.DeleteDetail(ctx, detail.Id)
	if err != nil {
		return fmt.Errorf("failed to delete detail: %w", err)
	}

	return nil
}
