package comm

import (
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"testing"
)

type AcPayload struct {
	TxId   string
	TxTime int64
	Memo   string
	//
	Name         string
	Organization string
}

type Ac struct {
	*AcPayload
	UpdatedAt int64
	UpdatedBy string
}

func TestMarshal(t *testing.T) {
	ac := Ac{
		AcPayload: &AcPayload{
			TxId:         "tx001",
			TxTime:       0,
			Memo:         "memo1",
			Name:         "name1",
			Organization: "org1",
		},
		UpdatedAt: 0,
		UpdatedBy: "admin",
	}
	bytes := Marshal(&ac)
	fmt.Printf("%v \n", string(bytes))

	result := Ac{}

	e := Unmarshal(bytes, &result)

	if e != nil {
		log.Fatal("%v", e)
	}

	fmt.Printf("%v\n", result)
}
