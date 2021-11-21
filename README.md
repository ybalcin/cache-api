# RESTful Cache API

The project is a cache api developed in Go. Currently, only supports memory caching.
It developed in DDD and Ports&Adapters architecture. Follows the SOLID principles and Idiomatic Go.

## Getting Started

If this is your first time encountering Go, please follow [the instructions](https://golang.org/doc/install) to
install Go on your computer.

```shell
# clone repo
git clone https://github.com/ybalcin/cache-api

cd cache-api/cmd

# run api server
go run main.go
```

At this time, you have a RESTful API server running at `http://127.0.0.1:8080`. It provides the following endpoints:

* `POST /v1/cache/set`: sets a key-value pair in cache, post data: {"key": "dummykey", "value": "dummyvalue"}
* `GET /v1/cache/get/:key`: returns key-value pair from cache by key
* `DELETE /v1/cache/flush`: flushes cache

Application also provides to saving memory data to file json under the OS tmp directory. This can be setting
while initialization the inemmorystore.Client by calling inmemorystore.NewClient(interval int). Application save cache
contents by a fixed interval time of minutes. If it is not desired the interval is set to zero. Also, when application started
it is load the last cache file contents to the memory.
(You can see the load and save info on the log.)

## Project Layout

The cache api uses the following project layout:

```
├── cmd                  main applications of the project
│   └── http             the http API server package
│   └── main.go          starts applications
├── internal             private application and library code
│   ├── application      application core
│   ├── common           common library for internal codes
│   ├── infrastructre    provides low level services for the application
│   ├── ports            library for accessing to the application and from applications to the infrastructure
├── pkg                  public library code
    └── inmemorystore    library for memory caching
```
Each package contains its own test. 