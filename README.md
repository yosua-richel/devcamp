## Getting Started
1. Install pre-requisite apps and tools
  * [Golang](https://golang.org/doc/install)
  * [NSQ](https://nsq.io/overview/quick_start.html)

2. Build and run
  * run nsq using docker
    ```
         $ docker-compose up -d --build
    ```
  * run nsq using binary
    ````
         $ nsqlookupd                                        --->   start the nsqlookupd
         $ nsqd --lookupd-tcp-address=127.0.0.1:4160         --->   start the nsqd
         $ nsqadmin --lookupd-http-address=127.0.0.1:4161    --->   start the nsq admin
    ````
  * cd to `producer-1` or `consumer-1` folder
  * run following command to start for each producer or consumer:
    ```
         $ go run main.go --debug 
    ```
