package main

import(
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

var ch4cpu chan uint64
var chTimer chan struct{}
var memMap map[int]interface{}

func init(){
	ch4cpu = make(chan uint64,10000)
	chTimer = make(chan struct{},20)
	memMap = make(map[int]interface{})
}