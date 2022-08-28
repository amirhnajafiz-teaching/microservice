<h1 align="center">
Micro Go
</h1>

**Micro-Go** is a golang project that implements Micro-service architecture 
using **go-kit**. In this project I implemented a microservice application that does string
operations like concat, count, split, ...

The base idea behind this project was to work with **go-kit**, and learn how
to implement a golang application using microservice architecture.

## What do you learn from this project?
- Microservice in Golang
- Go-kit

## What is microservice?
Microservices - also known as the microservice 
architecture - is an architectural style 
that structures an application as a collection 
of services that are:
- Highly maintainable and testable
- Loosely coupled
- Independently deployable
- Organized around business capabilities
- Owned by a small team
- 
The microservice architecture enables the rapid, 
frequent and reliable delivery of large, 
complex applications. 

It also enables an organization to evolve 
its technology stack.

To see how does this microservice work, check the following schema:

<p align="center">
<img src="./assets/microservice-architecture.png" width="400" alt="demo" />
</p>

## What is Go-kit?
Go kit is a programming toolkit for building microservices 
(or elegant monoliths) in Go. 
It solves common problems in distributed systems and application architecture,
so you can focus on delivering business value.

Go is a great general-purpose language, 
but microservices require a certain amount of specialized support. 
RPC safety, system observability, infrastructure integration, 
even program design — Go kit fills in the gaps left by the standard library, 
and makes Go a first-class language for writing microservices in any organization.

## How to use this project?
Clone into repository and set up the services:
```shell
go run main.go
```

#### uppercase string service
url:
```shell
[HOST]:[PORT]/uppercase
```

method:
```shell
POST
```

request:
````go
type UppercaseRequest struct {
	S string `json:"s"`
}
````

response:
```go
type UppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}
```

#### lowercase string service
url:
```shell
[HOST]:[PORT]/lowercase
```

method:
```shell
POST
```

request:
````go
type LowercaseRequest struct {
    S string `json:"s"`
}
````

response:
```go
type LowercaseResponse struct {
    V   string `json:"v"`
    Err string `json:"err,omitempty"`
}
```

#### string length service
url:
```shell
[HOST]:[PORT]/count
```

method:
```shell
POST
```

request:
````go
type CountRequest struct {
    S string `json:"s"`
}
````

response:
```go
type CountResponse struct {
    V int `json:"v"`
}
```

#### concatenate string service
url:
```shell
[HOST]:[PORT]/concatenate
```

method:
```shell
POST
```

request:
````go
type ConcatenateRequest struct {
    S string `json:"s"`
    C string `json:"c"`
}
````

response:
```go
type ConcatenateResponse struct {
    V   string `json:"v"`
    Err string `json:"err,omitempty"`
}
```

#### split string service
url:
```shell
[HOST]:[PORT]/split
```

method:
```shell
POST
```

request:
````go
type SplitRequest struct {
    S string `json:"s"`
    K string `json:"k"`
}
````

response:
```go
type SplitResponse struct {
    V   []string `json:"v"`
    Err string   `json:"err,omitempty"`
}
```
