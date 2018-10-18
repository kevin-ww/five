package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	testcc "github.com/s7techlab/cckit/testing"
	"testing"
)

func TestCCInit(t *testing.T) {
	//
	cc := testcc.NewMockStub(`cars`, New())
	fmt.Printf("%v \n", cc)
	response := cc.Init()
	fmt.Printf("%v \n %v\n", response.Message, string(response.Payload))
	//cc.Invoke()

}

func TestCCInvoke(t *testing.T) {

	mockStub = shim.NewMockStub("ac", New())
	fmt.Printf("%v \n", mockStub)
	//mockStub.

	//mockStub.

}
