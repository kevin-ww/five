package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"k.top/chaincode/comm"
	"k.top/chaincode/router"
)

type ChainCode struct {
	name string
	r    *router.Router
	ctx  *router.Context
}

const CCName = `AcctChainCode`

func (c *ChainCode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success([]byte("Initial ..." + CCName))
}

func (c *ChainCode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	//return c.r.Handle(stub)
	return handle(stub)
}

func handle(stub shim.ChaincodeStubInterface) peer.Response {
	fn, args := stub.GetFunctionAndParameters()
	acl := &AcLedger{
		admin: "",
		State: comm.State{
			Stub:          stub,
			Bucket:        "account",
			RecordCreator: "acctAdmin",
		},
	}
	return process(acl, fn, []byte(args[0]))

}

func process(ledger *AcLedger, funcName string, payload []byte) peer.Response {
	//only 1 arg is allowed
	//handlerFunc := funcProvided[funcName]
	//handlerFunc()
	acPayload := &AcPayload{}

	comm.Unmarshal(payload, acPayload)

	if funcName == `create` {
		return encodeResponse(ledger.create(acPayload))
	} else if funcName == `update` {
		return encodeResponse(ledger.update(acPayload))
	} else if funcName == `has` {
		return encodeResponse(ledger.has(acPayload))
	} else if funcName == `get` {
		return encodeResponse(ledger.has(acPayload))
	}

	return shim.Error(``)
}

type handleFunc func(payload []byte) peer.Response

var funcProvided map[string]handleFunc

func New() *ChainCode {

	funcProvided = make(map[string]handleFunc)

	return &ChainCode{
		name: CCName,
		//r:    router,
		ctx: nil,
	}

}

func encodeResponse(data interface{}, err error) peer.Response {
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(comm.Marshal(data))
}

func main() {

	if err := shim.Start(New()); err != nil {
		fmt.Printf("Error starting %s: %s", CCName, err)
	}
}
