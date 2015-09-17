package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/bboreham/coatl/backend"
)

func addService(args []string) {
	if len(args) != 3 {
		log.Fatal("Must supply service name, address and port number")
	}
	serviceName := args[0]
	port, err := strconv.Atoi(args[2])
	if err != nil {
		log.Fatal("Invalid port number:", err)
	}
	backend.AddService(serviceName, args[1], port)
	fmt.Println("Added service:", serviceName)
}

type listServiceOpts struct {
	all bool
}

func (opts *listServiceOpts) listService(args []string) {
	backend.ForeachServiceInstance(opts.all, func(name, value string) {
		fmt.Println(name, value)
	}, func(name, value string) {
		fmt.Println("  ", name, value)
	})
}

func (opts *listServiceOpts) resetService(args []string) {
	var err error
	if opts.all {
		err = backend.RemoveAllServices()
	} else if len(args) == 1 {
		err = backend.RemoveService(args[0])
	} else {
		log.Fatal("Must supply service name or -a")
	}
	if err != nil {
		log.Fatal("Failed to delete:", err)
	}
}
