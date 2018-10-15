package comm

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/pkg/errors"
	"strings"
)

type State struct {
	Stub          shim.ChaincodeStubInterface
	Bucket        string
	RecordCreator string
}

var (
	ErrNoSuchRecord = errors.New(`no such record in ledger`)
)

type KeyGen func(string) string

func (kg *KeyGen) gen(s State) string {
	return strings.Join([]string{s.RecordCreator, s.Bucket, ""}, "|")
}

func (l *State) withBucketKey(k string) string {
	return strings.Join([]string{l.RecordCreator, l.Bucket, k}, "|")
}

func (s *State) Put(k string, values interface{}) error {
	return s.Stub.PutState(s.withBucketKey(k), Marshal(values))
}

func (s *State) Get(k string, target interface{}) error {

	bytes, e := s.Stub.GetState(s.withBucketKey(k))

	if e != nil {
		return e
	}

	if bytes == nil {
		return ErrNoSuchRecord
	}

	return Unmarshal(bytes, target)

}

func (s *State) Has(k string) (bool, error) {

	bytes, e := s.Stub.GetState(s.withBucketKey(k))

	if e != nil || bytes == nil {
		return false, e
	}

	return true, nil
}
