package main

import (
	"encoding/json"
	"fmt"
)

//generic
type Payload struct {
	TxId   string
	TxTime int64
	Memo   string //
}

type Model interface {
	IsValid() bool
	Key() string
	Marshal() []byte
	UnMarshal(d []byte) (interface{},error)
}

//inbound
type TknPayload struct {
	Model
	*Payload
	from string
	to   string
	amt  int64
}

//outbound
type Tkn struct {
	*TknPayload
	UpdatedAt int64
	UpdatedBy string
}

func (t *Tkn) String() string{
	bytes, _ := json.Marshal(t)
	return string(bytes)
}


//func (t *Tkn) IsValid() bool{
//	return true
//}

func (p *Payload) IsValid() bool{
	return true
}


func (p *Payload) Key() string{
	return ""
}

func (t *Payload) Marshal() []byte{
	bytes, _ := json.Marshal(t)
	return bytes
}

func (t *Payload) UnMarshal(d []byte) (interface{},error){
	var r =& Tkn{}
	err := json.Unmarshal(d, r)
	if err!=nil{
		return nil,err
	}
	return r,nil
}

func (t *Tkn) Marshal() []byte{
	bytes, _ := json.Marshal(t)
	return bytes
}

//func (p *Payload) Marshal() []byte{
//	bytes, _ := json.Marshal(p)
//	return bytes
//}

func (t *Tkn) UnMarshal(d []byte) (interface{},error){
	var r =&Tkn{}
	err := json.Unmarshal(d, r)
	if err!=nil{
		return nil,err
	}
	return r,nil
}

func main(){

	payload := & TknPayload{
		Model: nil,
		Payload: &Payload{
			TxId:   "tx001",
			TxTime: 0,
			Memo:   "a test payload",
		},
		from: "alice",
		to:   "bob",
		amt:  10,
	}

	//payload.IsValid()IsValid

	tkn := &Tkn{
		TknPayload: payload,
		UpdatedAt:  0,
		UpdatedBy:  "kevin",
	}

	//payload.IsValid()

	fmt.Printf("%v\n",*tkn)

	fmt.Printf("%v\n",tkn.IsValid())

//IsValid	tkn.Marshal()
	tknAsBytes := tkn.Marshal()

	fmt.Printf("%v\n",tknAsBytes)

	aTkn, e := tkn.UnMarshal(tknAsBytes)

	if e!=nil{
		fmt.Printf("%v\n",e)
	}

	//do cast
	fmt.Printf("%T\n",aTkn)

	fmt.Printf("%v\n",aTkn)


	i := aTkn.(*Tkn)

	fmt.Printf("%v \n",i)

	//put(payload,)
	//tkn.IsValid()
}

//func process(fn string , p string) peer.Response{
//	return shim.Error(``)
//}
//
//
//func create(data []byte ,target *Tkn) (*Tkn ,error){
//
//	//Kevin.Put()
//	return nil,nil
//}
//
//
//func put(data []byte,target interface{}) (interface{},error){
//	return nil,nil
//}