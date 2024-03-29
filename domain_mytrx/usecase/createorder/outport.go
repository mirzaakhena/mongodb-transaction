package createorder

import (
	"mongodb-trx/domain_mytrx/model/repository"
	sharedrepository "mongodb-trx/shared/model/repository"
)

// Outport of usecase
type Outport interface {
	sharedrepository.WithTransactionDB
	repository.SaveOrderRepo
	repository.SavePersonRepo
}
