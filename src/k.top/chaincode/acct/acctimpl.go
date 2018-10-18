package main

import (
	"errors"
	"k.top/chaincode/comm"
	"log"
	"strings"
	"time"
)

//generic
type payload struct {
	TxId   string
	TxTime int64
	Memo   string //
}

type payloadType interface {
	IsValid() bool
	Key() string
	Marshal() []byte
	UnMarshal(data []byte) *payload
}

//inbound
type AcPayload struct {
	payload
	Name         string
	Organization string
}

//outbound
type Ac struct {
	*AcPayload
	UpdatedAt int64
	UpdatedBy string
}

func (a *Ac) isValid() bool {
	return true
}

func (a *Ac) Key() string {
	return a.Name
}

var (
	ErrInValidPayload      = errors.New(`this payload is invalid`)
	ErrAccountAlreadyExist = errors.New(`account already exist`)
	ErrNoSuchAccount       = errors.New(`no such account in ledger`)
)

func (ac *Ac) String() string {
	return strings.Join([]string{ac.TxId, ac.Memo, ac.Name}, "|")
}

type AcLedger struct {
	admin string
	comm.State
}

type ledger interface {
}

//biz logic

//func (l *AcLedger) create1(payloadAsBytes []byte) (*Ac, error) {
//
//	var payload = &AcPayload{}
//	e := comm.Unmarshal(payloadAsBytes, *payload)
//	if e != nil {
//		return nil, e
//	}
//	return l.create(payload)
//}

func (l *AcLedger) create(payload *AcPayload) (*Ac, error) {

	if !payload.isValid() {
		return nil, ErrInValidPayload
	}

	if b, err := l.has(payload); err != nil {
		return nil, err
	} else if b == true {
		return nil, ErrAccountAlreadyExist
	}

	ac := &Ac{
		AcPayload: payload,
		UpdatedAt: time.Now().Unix(),
		UpdatedBy: l.admin,
	}

	log.Printf(`%s is creating %v at %s \n`, l.admin, *ac, l)

	return ac, l.Put(ac.GenKey(), ac)
}

func (l *AcLedger) get(payload *AcPayload) (*Ac, error) {
	//copy
	ac := &Ac{}
	e := l.Get(payload.GenKey(), ac)
	return ac, e
}

func (l *AcLedger) has(payload *AcPayload) (bool, error) {
	return l.Has(payload.GenKey())
}

func (l *AcLedger) update(payload *AcPayload) (*Ac, error) {

	if !payload.isValid() {
		return nil, ErrInValidPayload
	}

	if b, err := l.has(payload); err != nil {
		return nil, err
	} else if b == true {
		return nil, ErrNoSuchAccount
	}

	return l.create(payload)
}

func (l *AcLedger) getAc(name string) (*Ac, error) {

	ac := &Ac{
		AcPayload: &AcPayload{
			Name: name,
		},
	}

	e := l.Get(ac.GenKey(), &Ac{})

	return ac, e
}

//func (l *AcLedger) hasAc(name string) (bool, error) {
//	return l.Has(name)
//}
//
//func (l *AcLedger) updateAc(payload *AcPayload) (*Ac, error) {
//
//	if !payload.isValid() {
//		return nil, ErrInValidPayload
//	}
//
//	if b, err := l.hasAc(payload.Name); err != nil {
//		return nil, err
//	} else if b == true {
//		return nil, ErrNoSuchAccount
//	}
//
//	return l.create(payload)
//}
