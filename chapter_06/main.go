package main

import (
	"fmt"
	countfun "gobible/chapter_06/interfaceContract"
)

func main() {
	var w countfun.WordCount
	var l countfun.LineCount
	w.Write([]byte("Our aim is to explain the enduring concepts underlying all computer systems, and to show you the concrete ways that these ideas affect the correctness, performance, and utility of your application programs. Many systems books are written from a builder's perspective, describing how to implement the hardware or the systems software, including the operating system, compiler, and network interface. This book is written from a programmer's perspective, describing how application programmers can use their knowledge of a system to write better programs. Of course, learning what a system is supposed to do provides a good first step in learning how to build one, so this book also serves as a valuable introduction to those who go on to implement systems hardware and software. Most systems books also tend to focus on just one aspect of the system, for example, the hardware architecture, the operating system, the compiler, or the network. This book spans all of these aspects, with the unifying theme of a programmer's perspective."))
	fmt.Print(w)
	fmt.Printf("\n")
	l.Write([]byte("abc\ndef\nghi\n"))
	fmt.Print(l)

}
