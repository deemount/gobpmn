# gobpmn #

Business Process Model Notation 2 with Go

Author: Salvatore Gonda

## Introduction ##

This is part of my journey through BPMN. To teach myself, I opened the Camunda Modeler, switch to XML tab and modelled the xsd template in Go.

### Status ###

... still in development

### Example ###

There some examples, but actually I have no time to describe it.
You should clone the repository and then start the application by

```go
go run main.go
```

### Docker ###

Reminder:

* RabbitMQ is just in a state of testing
* ElasticSearch is in a state of testing

```bash
docker run -d --hostname my-rabbit --name some-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management
```

```bash
docker run -d --name es01 -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" -it docker.elastic.co/elasticsearch/elasticsearch:7.14.0
```

### To Do's ###

Some To Do's are described inside the file, but not yet explained here

### Links ###

* BPMN Specification [https://www.omg.org/spec/BPMN]
* BPMN [https://camunda.com/bpmn/]

#### Practical ####

* Naming ID's [https://camunda.com/best-practices/naming-technically-relevant-ids/]

### History ###

* 2023-03-06 The Version differs totally to the one's below. Description coming soon
* 2021-02-21: Added more elements and tried a message queue
* 2021-03-15: Added more elements and delved deeper into structure
* 2021-01-26: Update
* 2021-01-24: First Upload/ Stable Realease
* 2021-01-23: Initialize Repository
