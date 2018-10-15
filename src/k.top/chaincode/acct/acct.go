package main

import (
	"errors"
	"k.top/chaincode/comm"
	"log"
	"strings"
	"time"
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

var (
	ErrInValidPayload      = errors.New(`this payload is invalid`)
	ErrAccountAlreadyExist = errors.New(`account already exist`)
	ErrNoSuchAccount = errors.New(`no such account in ledger`)
)

func (ac *Ac) String() string {
	return strings.Join([]string{ac.TxId, ac.Memo, ac.Name}, "|")
}

type Ledger struct {
	admin string
	comm.State
}

func (a *AcPayload) isValid() bool {
	return true
}

func (a *AcPayload) GenKey() string {
	return a.Name
}

func (l *Ledger) createAc(payload *AcPayload) (*Ac, error) {

	if !payload.isValid() {
		return nil, ErrInValidPayload
	}

	if b, err := l.hasAc(payload.Name); err != nil {
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

func (l *Ledger) getAc(name string) (*Ac, error) {

	ac := &Ac{
		AcPayload: &AcPayload{
			Name: name,
		},
	}

	e := l.Get(ac.GenKey(), &Ac{})

	return ac, e
}

func (l *Ledger) hasAc(name string) (bool, error) {
	return l.Has(name)
}

func (l *Ledger) updateAc(payload *AcPayload) (*Ac, error) {

	if !payload.isValid() {
		return nil, ErrInValidPayload
	}

	if b, err := l.hasAc(payload.Name); err != nil {
		return nil, err
	} else if b == true {
		return nil, ErrNoSuchAccount
	}

	return l.createAc(payload)
}
