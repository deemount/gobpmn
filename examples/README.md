# Examples

## Basic Assumption

1. A process model in BPMN format, which is guided from left to right and from top to bottom in the visualization, is only written from top to bottom in Go
2. The process model is described in Go only in a single struct or split, in several structs
3. Each process model points to a field to a specified interface, which is associated with the core package and only in the first position in the single or the first struct of a collection of structs
4. The first struct of the process model must contain the word Process
