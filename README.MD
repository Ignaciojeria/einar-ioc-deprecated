# Golang Dependency Injection Framework🔥

## 🔧 Installation
To install ioc, use the following command:

    go get github.com/Ignaciojeria/einar-ioc@v1.3.0

## :computer: Usage

The dependencies in the framework are loaded in a specific sequence, as follows:
1. Installations
2. Outbound Adapters
3. Use Cases
4. Inbound Adapters

This sequence ensures that each component is loaded in the correct order for optimal functionality. The diagram below illustrates this dependency loading process:

<p align="center">
  <img src="diagram.drawio.svg" alt="Diagram">
</p>

As a first step, we'll make sure that the `main` function loads all the dependencies we will inject later on. This initial loading of dependencies is crucial for setting up our Dependency Injection framework.

```go
package main

import (
	"os"
	ioc "github.com/Ignaciojeria/einar-ioc"
)
func main() {
	if err := ioc.LoadDependencies(); err != nil {
		os.Exit(0)
	}
}
