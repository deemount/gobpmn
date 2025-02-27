package main

// Cleanup interface for components that need cleanup
type Cleanup interface {
	Cleanup()
}
