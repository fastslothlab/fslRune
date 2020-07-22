package fslRune

import (
	"testing"
)

func TestIsRuneHasPixelNoTofu(ot *testing.T){
	ok(IsRuneNoTofu(32)==false)
	ok(IsRuneNoTofu(888)==false)
	ok(IsRuneNoTofu(55296)==false)
	ok(IsRuneNoTofu(33)==true)
}

func ok(b bool){
	if b==false{
		panic("ok fail")
	}
}
