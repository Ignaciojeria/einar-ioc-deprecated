# Golang Dependency Injection Framework🔥

## 🔧 Installation
To install ioc, use the following command:

    go get github.com/Ignaciojeria/einar-ioc@v1.3.0

## 🔍 Tutorial : Before start

The dependencies in the framework are loaded in a specific sequence, as follows:
1. Installations
2. Outbound Adapters
3. Use Cases
4. Inbound Adapters

This sequence ensures that each component is loaded in the correct order for optimal functionality. The diagram below illustrates this dependency loading process:

<p align="center">
  <img src="diagram.drawio.svg" alt="Diagram">
</p>

## 👨‍💻 Setup

As a first step, we'll make sure that the `main` function loads all the dependencies we will inject later on. This initial loading of dependencies is crucial for setting up our Dependency Injection framework.

```go
package main

import (
	"os"
	ioc "github.com/Ignaciojeria/einar-ioc"
)
func main() {
	if err := ioc.LoadDependencies(); err != nil {
		os.Exit(1)
	}
}
```

## 👨‍💻 Injecting your first Installation

Installations are key instances in the framework, used for a variety of infrastructure operations. These tasks include providing route instances for registering HTTP requests, initializing connections for database operations, setting up brokers for event publication and reception, among others. Effectively, installations lay the groundwork for the essential services that your application relies on, ensuring their seamless integration and operation as foundational elements.

####  🔍 Create your first Installation File
Navigate to the /chi_router folder. Inside this folder, we will create mux.go.
```bash
/chi_router
 - mux.go #Chi Router interface Implementation 
```

####  🔍 mux.go Implementation
We will be using the Chi Framework (https://github.com/go-chi/chi) and injecting the Chi router using IOC.
```go
package chi_router

import (
	ioc "github.com/Ignaciojeria/einar-ioc"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

var Mux = ioc.InjectInstallation[*chi.Mux](func() (*chi.Mux, error) {
	r := chi.NewRouter()
	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	return r, nil
})

```

####  🔍 Starting the Server with an Installation
```go
package main

import (
	"example/chi_router"
	"fmt"
	"net/http"
	"os"

	ioc "github.com/Ignaciojeria/einar-ioc"
)

func main() {
	if err := ioc.LoadDependencies(); err != nil {
		os.Exit(0)
	}
	fmt.Println("server starting at port 8080")
	http.ListenAndServe(":8080", chi_router.Mux.Dependency)
}
```

## 👨‍💻 Injecting your first Inbound Adapter

An "Inbound Adapter" encompasses all operations that trigger an inbound action in the system. This includes scenarios such as receiving an HTTP request from an external client or consuming an event published by an external producer. These adapters serve as the entry points for external inputs, ensuring that your system can respond and interact with external stimuli, such as client requests or external events, effectively integrating them into your system's workflow.

####  🔍 Create your first Inbound Adapter File

Navigate to the /controller folder. Inside this folder, we will create post_customer.go.
```bash
/controller
 - post_customer.go 
```

####  🔍 post_customer.go Implementation

```go
package controller

import (
	"example/chi_router"
	"net/http"

	ioc "github.com/Ignaciojeria/einar-ioc"
	"github.com/go-chi/chi/v5"
)

type postCustomers struct {
	router *chi.Mux
}

var postCustomersDependency = ioc.InjectInboundAdapter[postCustomers](func() (postCustomers, error) {
	controller := postCustomers{
		router: chi_router.Mux.Dependency, // Installation Dependency Injected Here!
	}
	controller.router.Post("/person", controller.handle)
	return controller, nil
})

func (ctrl postCustomers) handle(w http.ResponseWriter, r *http.Request) {
	w.Write("hello from post customer controller")
}
```

#### 🔍 Registering Controller Package Dependencies

```go
package main

import (
	"example/chi_router"
	"fmt"
	"net/http"
	"os"
	_ "example/controller" // Automatically injects all inbound adapters registered inside the controller package
	ioc "github.com/Ignaciojeria/einar-ioc"
)

func main() {
	if err := ioc.LoadDependencies(); err != nil {
		fmt.Println("Error loading dependencies:", err)
		os.Exit(1)
	}
	fmt.Println("Server starting at port 8080")
	http.ListenAndServe(":8080", chi_router.Mux.Dependency)
}
```