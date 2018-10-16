package router

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type handler func(args []string) peer.Response

type Router struct {
	name         string
	handlerGroup map[string]handler
	stub         shim.ChaincodeStubInterface
}

func (r *Router) bontext(path string, stub shim.ChaincodeStubInterface) {
	//return &context{
	//	stub:   nil,
	//	logger: nil,
	//	path:   "",
	//	args:   nil,
	//}
}

func New(name string) *Router {
	return &Router{
		name:         name,
		handlerGroup: make(map[string]handler),
	}
}

func (r *Router) invokeAll() peer.Response {
	fmt.Printf(`invoke all the handlers`)
	return shim.Error(``)
}

func (r *Router) HandleQuery(path string, h handler) *Router {
	r.handlerGroup[path] = h
	return r
}

func (r *Router) HandleInvoke(path string, h handler) *Router {
	r.handlerGroup[path] = h
	return r
}

func (r *Router) Build() *Router {
	return r
}

func (r *Router) Handle(stub shim.ChaincodeStubInterface) peer.Response {
	fn, args := stub.GetFunctionAndParameters()
	handler := r.handlerGroup[fn]
	//set up the stub
	r.stub = stub
	//
	return handler(args)
}
