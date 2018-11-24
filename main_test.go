package main

import "testing"

func TestRead(t *testing.T){
    var v bool
    var c Computer
    v = c.Read([]int{1,2,3,4})
    //если ввели разные циферки, то все ок
    if v != true {
        t.Error("ERROR", v)
    }
    //если ввели повтроряющиеся, то проблема
    v = c.Read([]int{1,1,2,3})
    if v != false {
        t.Error("ERROR", v)
    }
}
