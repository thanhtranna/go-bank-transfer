package presenter

import (
	"go-bank-transfer/domain"
	"go-bank-transfer/usecase"
	"time"
)

type findAllTransferPresenter struct{}

func NewFindAllTransferPresenter() usecase.FindAllTransferPresenter {
	return findAllTransferPresenter{}
}

func (a findAllTransferPresenter) Output(transfers []domain.Transfer) []usecase.FindAllTransferOutput {
	var o = make([]usecase.FindAllTransferOutput, 0)

	for _, transfer := range transfers {
		o = append(o, usecase.FindAllTransferOutput{
			ID:                   transfer.ID().String(),
			AccountOriginID:      transfer.AccountOriginID().String(),
			AccountDestinationID: transfer.AccountDestinationID().String(),
			Amount:               transfer.Amount().Float64(),
			CreatedAt:            transfer.CreatedAt().Format(time.RFC3339),
		})
	}

	return o
}