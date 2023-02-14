package collaborative_process

import "github.com/deemount/gobpmn/factory"

var Builder factory.Builder

type Proxy interface{ Build() Process }
