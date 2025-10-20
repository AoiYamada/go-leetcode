package main

// Medium: PrintFoobarAlternately
// Solution for print-foobar-alternately (medium)
type FooBar struct {
	n   int
	chF chan bool
	chB chan bool
}

func NewFooBar(n int) *FooBar {
	return &FooBar{
		n,
		// use buffered channel here so I could leave the last value of `fb.chB <- true` unconsumed LOL
		make(chan bool, 1),
		make(chan bool, 1),
	}
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		if i > 0 {
			<-fb.chB
		}

		// printFoo() outputs "foo". Do not change or remove this line.
		printFoo()
		fb.chF <- true
	}
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.chF
		// printBar() outputs "bar". Do not change or remove this line.
		printBar()
		// if i < fb.n-1 {
		fb.chB <- true
		// }
	}
}
