image: daocloud/ci-golang:1.4

services:
    - mongodb

env:
    - MYENV = /gopath/app
    - GOPATH = /go
    - MY_ENV = /gopath/app
    - echo uMNEZzBVHh8l9YfP:pPLpK1lkMS82cq95e@10.10.72.139:27017/lEyTj8hYrUIKgMfi
    - MONGODB_PORT = tcp://10.10.72.139:27017
    - MONGODB_PORT_27017_TCP = tcp://10.10.72.139:27017
    - MONGODB_PORT_27017_TCP_ADDR = 10.10.72.139
    - MONGODB_PORT_27017_TCP_PROTO = tcp
    - MONGODB_PORT_27017_TCP_PORT = 27017
    - MONGODB_INSTANCE_NAME = lEyTj8hYrUIKgMfi
    - MONGODB_PASSWORD = pPLpK1lkMS82cq95e
    - MONGODB_USERNAME = uMNEZzBVHh8l9YfP

install:
    - echo $MYENV

before_script:
    - echo $MYENV

script:
    - echo $MYENV
    - echo "This is an script segment"
    - echo "Run test cases here"
    - ping -c 2 mongodb
    - export GOPATH=$PWD
    - export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
    - go get github.com/shaalx/membership
    - go test -v ./db