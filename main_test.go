package main

import (
	"github.com/pkg/errors"
	"strconv"
	"testing"
	//"github.com/stretchr/testify/assert"
)

var inputMas []string = []string{
	"a",
	"",
	"a",
	"ab",
	"ba",
	"\n\n",
}

func TestSimpleSet(t *testing.T) {
	eq := func(s1 string, s2 string)error{
		if(s1 != s2){
			return errors.Errorf("%v != %v", s1, s2) 
		}
		return nil
	}

	set := NewSet()

	// add to set
	for _, item := range inputMas {
		set.Add(item)
	}

	// check size
	if(set.GetSize() != 5){
		t.Error("wrong size" + strconv.Itoa(set.GetSize()))
	}

	//create result array
	outputMap := []string{"a", "", "ab", "ba", "\n\n"}

	//check data by index
	for i:=0; i < set.GetSize(); i++{
		val, err := set.GetByIndex(i)
		if err == nil {
			err = eq(val,outputMap[i])
		} 
		if err != nil{
			t.Errorf("wrong data in set, index = %v, %v", i, err.Error())
		}
	}
	
	//check set.find
	if(!(set.Find("a") && set.Find("") && set.Find("\n\n"))){
		t.Error("can`t find element")
	} 

}
