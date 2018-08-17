# meatmelody

## building/running

### prerequisites
- ensure `$GOPATH` is set
- [`dep`](https://golang.github.io/dep/docs/installation.html)

`dep` can be installed via `go get -u github.com/golang/dep/cmd/dep`

then,

- clone to `$GOPATH/src/github.com/TheOrangeOne/meatmelody`
- install dependencies: `dep ensure`
- install: `go install`
- to run: `meatmelody`
- navigate to `localhost:8080`

## commandline args
- `--host <host>`
- `--port <port>`

Note that the environment variables `HOST` and `PORT` can be specified as well.

## testing

### javascript

to test the frontend code run, use [mocha](https://mochajs.org/):

```sh
$ mocha --watch static/test
```
