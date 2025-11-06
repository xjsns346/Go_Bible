package main

type Writer interface {
	Write(p []byte) (n int, err error)
}

func count(p []byte) (n int, err error) {}

type a int

func (aa *a) fun(p []byte) (n int, err error) {}

type b struct {
	bond int
}

func (b *b) joy(p []byte) (n int, err error) {}
