# Golang Dependency Injection Framework🔥

## 🔧 Installation
To install ioc, use the following command:

    go get github.com/Ignaciojeria/einar-ioc@1.4.0

## 🔍 Tutorial : Before start

The dependencies in the framework are registered in a specific sequence, as follows:
1. Installations
2. Outbound Adapters
3. Use Cases
4. Inbound Adapters

This sequence ensures that each component is registered in the correct order for optimal functionality. The diagram below illustrates this dependency registry order:

<p align="center">
  <img src="diagram.drawio.svg" alt="diagram">
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

## 👨‍💻 HTTP Router Installation Registration

Installations are key instances in the framework, used for a variety of infrastructure operations. These tasks include providing route instances for registering HTTP requests, initializing connections for database operations, setting up clients for event publication and reception, among others. Effectively, installations lay the groundwork for the essential services that your application relies on, ensuring their seamless integration and operation as foundational elements.

####  🔍 Create your first Installation File
Navigate to the /chi_router folder. Inside this folder, we will create mux.go
```bash
/chi_router
 - mux.go #Chi Router interface Implementation 
``` 

####  🔍 mux.go Implementation
We will be using the Chi Router (https://github.com/go-chi/chi) and registering the Chi router using IOC.

	go get -u github.com/go-chi/chi/v5 

```go
package chi_router

import (
	ioc "github.com/Ignaciojeria/einar-ioc"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)
// Registered installation
var MuxContainer = ioc.Installation[*chi.Mux](func() (*chi.Mux, error) {
	r := chi.NewRouter()
	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	return r, nil
})

```

####  🔍 Starting The Server With An Registered Installation
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
	// Retrieving the registered installation to start the server
	http.ListenAndServe(":8080", chi_router.MuxContainer.Dependency)
}
```

## 👨‍💻 HTTP Inbound Adapter Registration

An "Inbound Adapter" encompasses all operations that trigger an inbound action in the system. This includes scenarios such as receiving an HTTP request from an external client or consuming an event published by an external producer. These adapters serve as the entry points for external inputs, ensuring that your system can respond and interact with external stimuli, such as client requests or external events, effectively integrating them into your system's workflow.

####  🔍 Create your first HTTP Inbound Adapter File

Navigate to the /controller folder. Inside this folder, we will create post_customer.go
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

type postCustomer struct {
}

var postCustomerContainer = ioc.InboundAdapter[postCustomer](func() (postCustomer, error) {
	controller := postCustomer{}
	 // Installation Dependency retrieved Here!
	chi_router.MuxContainer.Dependency.Post("/customers", controller.handle)
	return controller, nil
})
func (ctrl postCustomer) handle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from post customer controller"))
}
```

#### 🔍 Discovery Controller Package Dependencies

```go
package main

import (
	"example/chi_router"
	"fmt"
	"net/http"
	"os"
	_ "example/controller" // Required to discover all inbound adapters registered inside the controller package
	ioc "github.com/Ignaciojeria/einar-ioc"
)

func main() {
	if err := ioc.LoadDependencies(); err != nil {
		fmt.Println("Error loading dependencies:", err)
		os.Exit(1)
	}
	fmt.Println("Server starting at port 8080")
	http.ListenAndServe(":8080", chi_router.MuxContainer.Dependency)
}
```

## 👨‍💻 Usecase Registration & Injection
Use cases are specialized functions designed to execute business logic. They handle key operations aligning with business goals

####  🔍 Create your first Usecase File

Navigate to the /usecase folder. Inside this folder, we will create create_customer.go
```bash
/controller
 - create_customer.go 
```

####  🔍 create_customer.go Implementation
```go
package usecase

import (
	"fmt"

	ioc "github.com/Ignaciojeria/einar-ioc"
)

type createCustomer struct {
}

// Usecase interface
type ICreateCustomer interface {
	Execute(customerName string)
}

// [ICreateCustomer] is a generic type that is flexible and can be specified by you.
var CreateCustomerContainer = ioc.UseCase[ICreateCustomer](func() (ICreateCustomer, error) {
	usecase := createCustomer{}
	// Returning an exposed structure to be injected, which matches the specified interface.
	return usecase, nil
})

// Implementation of an exposed method that matches the method signature of the interface.
func (u createCustomer) Execute(customerName string) {
	fmt.Println("Customer name to create : " + customerName)
	// TODO : CUSTOMER CREATION PROCESS
}
```

####  🔍 Injecting Usecase inside HTTP inbound adapter

```go
package controller

import (
	"example/chi_router"
	"example/usecase"
	"net/http"

	ioc "github.com/Ignaciojeria/einar-ioc"
	"github.com/go-chi/chi/v5"
)

type postCustomer struct {
	createCustomer usecase.ICreateCustomer // [ICreateCustomer] exposed by the IOC container
}

var postCustomerContainer = ioc.InboundAdapter[postCustomer](func() (postCustomer, error) {
	controller := postCustomer{      
		// [ICreateCustomer] use case dependency is injected here!
		createCustomer: usecase.CreateCustomerContainer.Dependency, 
	}
	// Installation Dependency retrieved Here!
	chi_router.MuxContainer.Dependency.Post("/customers", controller.handle)
	return controller, nil
})

func (ctrl postCustomer) handle(w http.ResponseWriter, r *http.Request) {
	// Here, we are calling the injected dependency.
	ctrl.createCustomer.Execute("juan")
	w.Write([]byte("customer created"))
}
```

## 👨‍💻 Memory Database Installation Registration

We will be using the hashicorp memdb (https://github.com/hashicorp/go-memdb) and registering the memdb using IOC.

	go get github.com/hashicorp/go-memdb

```go
package db

import (
	ioc "github.com/Ignaciojeria/einar-ioc"
	"github.com/hashicorp/go-memdb"
)

var MemDBContainer = ioc.Installation[*memdb.MemDB](func() (*memdb.MemDB, error) {
	// Create the DB schema
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"customer": {
				Name: "customer",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"name": {
						Name:    "name",
						Unique:  false,
						Indexer: &memdb.IntFieldIndex{Field: "Name"},
					},
				},
			},
		},
	}
	return memdb.NewMemDB(schema)
})
```