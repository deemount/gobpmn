# gobpmn

![gobpmn](https://github.com/deemount/gobpmn/blob/main/docs/img/header.webp "a gopher in front of business model")

## Description

gobpmn is an experimental library that makes it possible to reflect

* structs,
* anonymous fields,
* typed and
* generic maps

in Go and create business process models with it.

As a modeler with a CLI, gobpmn is a monorepo with clear encapsulation.

***STILL IN DEVELOPMENT***

## Requirements

* Go >= 1.24

Optional:

* Saxon >= 12.5 (Requires Java >= 23)

## Install

1. Clone the repository or the download the .zip-File of the repository
2. Install the latest gobpmn version as a package within a module with **go get**
3. Install the latest gobpmn version as a package outside of a module with **go install**

```bash
go get github.com/deemount/gobpmn@latest
go install github.com/deemount/gobpmn@latest
```

## Examples

```bash
go run examples/generic_map_simple_process/main.go
go run examples/typed_map_simple_process/main.go
go run examples/renting_process/main.go
go run examples/simple_process/main.go
go run examples/small_process/main.go
```

## How To

```go
package main

import (
    "log"
    "github.com/deemount/gobpmn"
)

type Process struct {
    Def             gobpmn.Repository
    IsExecutable    bool
    Process         gobpmn.BPMN
    StartEvent      gobpmn.BPMN
    FromStartEvent  gobpmn.BPMN
    Task            gobpmn.BPMN
    FromTask        gobpmn.BPMN
    EndEvent        gobpmn.BPMN
}

func (p Process) GetDefinitions() gobpmn.Repository {
    return p.Def
}

func main() {
    _, err := gobpmn.FromStruct(Process{})
    if err != nil {
        log.Fatalf("ERROR: %s", err)
        return
    }
}
```

```go
package main

import (
    "log"
    "github.com/deemount/gobpmn"
)

var Process = map[string]any{
    "Def":                 gobpmn.Definitions(),
    "IsExecutable":        true,
    "Process":             gobpmn.BPMN{Pos: 1},
    "StartEvent":          gobpmn.BPMN{Pos: 2},
    "FromStartEvent":      gobpmn.BPMN{Pos: 3},
    "Task":                gobpmn.BPMN{Pos: 4},
    "FromTask":            gobpmn.BPMN{Pos: 5},
    "EndEvent":            gobpmn.BPMN{Pos: 6},
}

func main() {
    type T any
    _, err := gobpmn.FromMap[T](Process)
    if err != nil {
        log.Fatalf("ERROR: %s", err)
        return
    }
}
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
