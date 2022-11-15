package payload

import "mongodb-trx/shared/gogen"

type Payload struct {
	Data      interface{}           `json:"data"`
	Publisher gogen.ApplicationData `json:"publisher"`
	TraceID   string                `json:"traceId"`
}
