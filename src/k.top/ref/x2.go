package main

import "fmt"

type Bouncer interface {
	Bounce()
}


func BounceIt(b Bouncer) {
	b.Bounce()
}


type bs struct{

}

func (b *bs) Bounce(){
	fmt.Printf("hahaha\n")
}

func main(){
	b := &bs{}
	BounceIt(b)
}