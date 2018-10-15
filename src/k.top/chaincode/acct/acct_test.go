package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"k.top/chaincode/comm"
	"log"
	"testing"
)

var (
	mockStub *shim.MockStub
	ledger   *Ledger
)

func init() {

}

func TestCreateAc(t *testing.T) {

	mockStub = shim.NewMockStub("", nil)

	ledger := Ledger{
		admin: "admin",
		State: comm.State{
			Stub:          mockStub,
			Bucket:        "account",
			RecordCreator: "kevin",
		},
	}

	payload := &AcPayload{
		TxId:   "tx001",
		TxTime: 0,
		Memo:   "test create account",
		Name:   "acct002",
	}

	fmt.Printf("%v \n", *payload)

	txId := `000001`
	mockStub.MockTransactionStart(txId)
	ac, e := ledger.createAc(payload)
	mockStub.MockTransactionEnd(txId)

	if e != nil {
		log.Fatalln(e)
	}

	fmt.Printf("%v \n", ac)

	a := Ac{}

	e = ledger.Get(payload.Name, &a)
	if e != nil {
		log.Fatalln(e)
	}
	fmt.Printf("%v\n", a)

}

func TestGetAcct(t *testing.T) {

	mockStub = shim.NewMockStub("", nil)

	ledger := Ledger{
		admin: "admin",
		State: comm.State{
			Stub:          mockStub,
			Bucket:        "account",
			RecordCreator: "kevin",
		},
	}

	txId := `000001`
	mockStub.MockTransactionStart(txId)
	//ac, e := ledger.createAc(payload)
	ac, e := ledger.getAc("acct001")
	mockStub.MockTransactionEnd(txId)

	if e != nil {
		log.Fatalln(e)
	}

	fmt.Printf("%v", ac)
}

func TestExist(t *testing.T) {

	mockStub = shim.NewMockStub("", nil)

	ledger := Ledger{
		admin: "admin",
		State: comm.State{
			Stub:          mockStub,
			Bucket:        "account",
			RecordCreator: "kevin",
		},
	}

	fmt.Printf("%v \n", ledger)

	b, e := ledger.hasAc("kevin")
	if e != nil {
		log.Fatalln(e)
	}
	fmt.Printf("%v \n", b)

}

func TestFun(t *testing.T) {
	fmt.Printf("hahaha")
}
