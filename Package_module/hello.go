package hello

//A module is a collection of Go packages stored in a file tree with a go.mod file at its root
//As of Go 1.11, the go command enables the use of modules when the current directory or any parent
//directory has a go.mod, provided the directory is outside $GOPATH/src. (Inside $GOPATH/src, for
//compatibility, the go command still runs in the old GOPATH mode, even if a go.mod is found. See the go command documentation for details.)

// //(1)

//CREATING MODULE
//go mod init GO_WORK

//TEST THE CODE
//go test

// func Hello() string {
// 	return "Hello, world."
// }

// //(2)Adding a dependency
////The primary motivation for Go modules was to improve the experience of using (that is, adding a dependency on) code written by other developers.

//CHECK ALL THE DEPENDENCY
//go list -m all

//UPGRADING THE DEPENDENCIES
//go get golang.org/x/text
//go get rsc.io/sampler

// BUT ABOVE VERSION  NOT COMPATIBLE WITH OUR USAGE
//$ go list -m -versions rsc.io/sampler
// rsc.io/sampler v1.0.0 v1.2.0 v1.2.1 v1.3.0 v1.3.1 v1.99.99
// $

//We had been using v1.3.0; v1.99.99 is clearly no good. Maybe we can try using v1.3.1 instead:
// $ go get rsc.io/sampler@v1.3.1

import "rsc.io/quote"

func Hello() string {
	return quote.Hello()
}
