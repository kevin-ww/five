package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
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
	return c.r.Handle(stub)
}

//^^^

func getAc(args []string) peer.Response {
	return shim.Error(``)
}

func hasAc(args []string) peer.Response {
	return shim.Error(``)
}

func createAc(args []string) peer.Response {
	return shim.Error(``)
}


func New() *ChainCode{

	router := router.
		New(CCName).
		HandleQuery(`getAc`, getAc).
		HandleQuery(`hasAc`, hasAc).
		HandleInvoke(`createAc`, createAc).
		Build()

	//router.Context{}

	return &ChainCode{
		name: CCName,
		r:    router,
		ctx:  nil,
	}

}

func main() {

	if err := shim.Start(New()); err != nil {
		fmt.Printf("Error starting %s: %s", CCName, err)
	}
}
