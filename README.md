# Your Application

## Prerequisite
```
https://registry.hub.docker.com/r/candis/mongo-replica-set
```

## Contract
```
shared/model/repository/db_transaction.go
```

## Contract Service Helper
```
shared/model/service/db_transaction_impl.go
```

## Contract Implementation
```
shared/infrastructure/database/mongo_db.go
```

## Contract Client
```
domain_belajar/usecase/createorder/interactor.go
domain_belajar/usecase/createorder/outport.go
```

```text

How to identify Transaction
See a Process a command (Insert, Update or Delete) that appear more than once in Outport

Execute () {
  Transaction {
    validation ... --> 1
    calculation ... -> 2 
    save(a) ---------> 3
    validation ... --> 4
    calculation ... -> 5
    save(b) ---------> 6
    calculation ... -> 7
    save(c) ---------> 8
  }
}

if somehow in 7 (for example) has an error then we will rollback step 6 and 3

  
Step 1 : identify all outport that has more than one command

Step 2 : 

add repository.WithTransactionDB in Outport

call service.WithTransaction function in interactor

res, err = service.WithTransaction(ctx, r.outport, func(dbCtx context.Context) (*InportResponse, error) {

  // put everything here that need to be transaction-ed

  return res, nil
})

remember to adding if err != nil {return nil, err} at the end of function service.WithTransaction

replace T with InportResponse

replace dbCtx with ctx

add *database.MongoWithTransaction in gateway struct

initiate MongoWithTransaction: mwt, in return part of gateway
```
