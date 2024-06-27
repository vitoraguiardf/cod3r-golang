package main

import "runtime/debug"

func main() {
	f1()
}

func f3() {
	debug.PrintStack()
}

func f2() {
	f3()
}

func f1() {
	f2()
}
