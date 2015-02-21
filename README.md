# go-cg
A collection of graphics-related packages written in go

### Getting started with go:

Install go with your default/preferred package manager.

NOTE: depending on your distriburion, the package may be named any of the following:
* go
* golang
* golang-go

##### After the installation, make sure your $GOPATH is set:

```sh
# to check if set, execute:
echo $GOPATH

# if you do NOT get anything from the above, do the following:
mkdir ~/.go
export GOPATH=~/.go
```

**NOTE**: you may want to put that last line in your ~/.{bash,csh,ksh,sh,zsh}rc

### Getting and testing everything:

To get this package, simply execute the following:
```sh
go get "github.com/aznashwan/go-cg"
```
Now you can find the package it in $GOPATH/src/github.com/aznashwan/go-cg

##### This package has the following structure:
* /\* -  implementation for the various individual packages
* /applications/\* - small programs to test each of the main packages

##### To test something from /applications, simply do the following:

```sh
cd $GOPATH/src/github.com/aznashwan/go-cg/applications

git pull # to pull in the latest version

# all the applications are self-contained, you can run each of them as follows:
go run the_application.go

# or build a binary file and execute that:
go build the_application.go
./the_application.go
```

**NOTE**: any files an application may write/read will be written/read from the
current working dorectory!

### Additional resources:

* [go main page](http://golang.org/)
* [basic go overview](http://en.wikipedia.org/wiki/Go_%28programming_language%29)
* [official go tutorial](https://tour.golang.org/welcome/1)

