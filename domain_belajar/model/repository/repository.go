package repository

import (
	"context"
	"mongodb-trx/domain_belajar/model/entity"
)

type SaveOrderRepo interface {
	SaveOrder(ctx context.Context, obj *entity.Order) error
}

type SavePersonRepo interface {
	SavePerson(ctx context.Context, obj *entity.Person) error
}
