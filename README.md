# DSP API Go implementation

Simple DSP(Demand Side Platform) API written in golang compliant with [OpenRTB v2.3](https://github.com/openrtb/OpenRTB/blob/master/OpenRTB-API-Specification-Version-2-3-FINAL.pdf)

# Required middleware 

- [Redis](http://redis.io/download)
- [fluentd](http://www.fluentd.org/download)

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

# End Points

## RTB Request [/v1/bid] POST

Receive RTB request. RTB request must be compliant with [OpenRTB v2.3](https://github.com/openrtb/OpenRTB/blob/master/OpenRTB-API-Specification-Version-2-3-FINAL.pdf)

+ Response 200 (application/json)
  + Headers

  ```
  x-openrtb-version: 2.3
  ```

  + Body

  ```
  {
      "cur": "JPY", 
      "id": "IxexyLDIIk", 
      "seatbid": [
          {
              "bid": [
                  {
                      "adid": "1234", 
                      "adm": "<a href=\"http://test.noadnolife.com/v1/click/12345?impid=${AUCTION_IMP_ID}&price=${AUCTION_PRICE}\"><img src=\"http://test.noadnolife.com/img/12345?impid=${AUCTION_IMP_ID}&price=${AUCTION_PRICE}\" width=\"728\" height=\"90\" border=\"0\" alt=\"Advertisement\" /></a>", 
                      "cid": "1234", 
                      "crid": "12345", 
                      "id": "d8aa5114-cd86-4484-9947-b907c8d12daa", 
                      "impid": "1", 
                      "iurl": "http://test.noadnolife.com/img/12345.png", 
                      "nurl": "http://test.noadnolife.com/v1/win/12345?impid=${AUCTION_IMP_ID}&price=${AUCTION_PRICE}", 
                      "price": 225
                  }
              ]
          }
      ]
  }
  ```

+ Response 204 (application/json)
  + Headers

  ```
  x-openrtb-version: 2.3
  ```

## Win Notice [/v1/win/:crid] GET

Receive win notice. In order to track logs, price and impression ID are required.

+ Parameters
  + price (number, required) - Won price of RTB Auction
  + impid (number, required) - An impression ID of win notice


+ Response 200 (application/json)
  + Headers

  ```
  x-openrtb-version: 2.3
  ```

  + Body

  ```
  { "message": "ok"}
  ```

## Click /v1/click/:crid GET

 Receive click action log. In order to track logs, price and impression ID are required.

+ Parameters
  + price (number, required) - Won price of RTB Auction
  + impid (number, required) - An impression ID of win notice

+ Response 200 (application/json)
  + Headers

  ```
  x-openrtb-version: 2.3
  ```

  + Body

  ```
  { "message": "ok"}
  ```

# Run with mock data

Put mock data to Redis.

```
$ cd command
$ go run create_mock.go
```

Send RTB Request via [httpie](https://github.com/jkbrzt/httpie)

```
$ http POST http://<host>/v1/bid < examples/request.mobile.json
```

Send Win Notice

```
$ http GET http://<host>/v1/win/XXXX impid==YYYYYYYY price==100
```

If you want to add/modify mock data, please edit data file.

```
$ vim command/data.yml
```

# License

See [Licence](https://github.com/satoshi03/go-dsp-api/blob/master/LICENSE)
