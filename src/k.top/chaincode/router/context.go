package router

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/ledger"
)

type context struct {
	stub   shim.ChaincodeStubInterface
	logger *shim.ChaincodeLogger
	path   string
	args   []string
}

type Context interface {
	Stub() shim.ChaincodeStubInterface
	Logger() *shim.ChaincodeLogger
	Path() string
	Args() []string
	Ledger() ledger.BlockAndPvtData
}
