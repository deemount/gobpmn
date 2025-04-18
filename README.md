# gobpmn

![gobpmn](https://github.com/deemount/gobpmn/blob/main/docs/img/header.webp "a gopher in front of business model")

## Description

gobpmn is an experimental library that makes it possible to reflect

* structs,
* anonymous fields,
* typed and
* generic maps

in Go and create business process models with it.

***STILL IN DEVELOPMENT***

## Requirements

* Go >= 1.24

Optional:

* Saxon >= 12.5 (Requires Java >= 23)

## Install

1. Clone the repository or the download the .zip-File of the repository
2. To run the app, copy and paste this into your terminal

```bash
go run examples/generic_map_simple_process/main.go
go run examples/typed_map_simple_process/main.go
go run examples/renting_process/main.go
go run examples/simple_process/main.go
```

1. Install Saxon XSL Transformer (MacOS) with Homebrew
2. Run the main file in the cmd/saxon/ directory

```bash
brew install saxon
go run cmd/saxon/main.go
```

## Wiki

Read the [documentation](https://github.com/deemount/gobpmn/wiki)

## Further Links

* [BPMN Specification](https://www.omg.org/spec/BPMN)
* [Camunda BPMN](https://camunda.com/bpmn/)
* [bpmn.io](https://bpmn.io/)

Checkout the other projects for BPMN written in Go

* [gobpm](https://github.com/dr-dobermann/gobpm)
* [lib-bpmn-engine](https://github.com/nitram509/lib-bpmn-engine)

Check my articles for this project on medium.com

* [Two-Pass-Matching with Go](https://medium.com/@salvatoregonda/two-pass-matching-with-go-480faffe88fa)
* [Go & BPMN: A modelling approach with reflect, Part 1](https://medium.com/@salvatoregonda/go-bpmn-a-modelling-approach-with-reflect-part-1-6f572adeac79)
* [Go & BPMN: A programmatic approach](https://medium.com/@salvatoregonda/go-bpmn-a-programmatic-approach-c25cbef45cc6)
