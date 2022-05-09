package createorder

import (
	"context"
	"mongodb-trx/domain_belajar/model/entity"
	"mongodb-trx/shared/model/service"
)

//go:generate mockery --name Outport -output mocks/

type createOrderInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase
func NewUsecase(outputPort Outport) Inport {
	return &createOrderInteractor{
		outport: outputPort,
	}
}

// Execute the usecase
func (r *createOrderInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	res := &InportResponse{}

	err = service.WithTransaction(ctx, r.outport, func(dbCtx context.Context) error {

		order, err := entity.NewOrder("pisang", 12)
		if err != nil {
			return err
		}

		err = r.outport.SaveOrder(dbCtx, order)
		if err != nil {
			return err
		}

		person, err := entity.NewPerson("mirza", 28)
		if err != nil {
			return err
		}

		err = r.outport.SavePerson(ctx, person)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
