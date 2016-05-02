# DSP API Go implementation

Simple DSP(Demand Side Platform) API written in golang.

# Required middleware 

Redis
fluentd

# How to setup

Copy source to your local dierctory.

```
$ git clone git@github.com:satoshi03/go-dsp-api.git
```

Install go libs and compile.

```
$ go get -v -d
$ go build
```

Copy config and edit.

```
$ cp config/config.yml.def config.yml
$ vim config.yml
```

Run.

```
$ ./go-dsp-api
```

# Run with Mock data

Put mock data to Redis.

```
$ cd command
$ go run create_mock.go
```

Send RTB Request via [httpie](https://github.com/jkbrzt/httpie)

```
$ http POST http://<host>/v1/bid examples/request.mobile.json
```

Send Win Notice

```
$ http GET http://<host>/v1/win/XXXX impid==YYYYYYYY price==100
```

# License

See [Licence](https://github.com/satoshi03/go-dsp-api/blob/master/LICENSE)
