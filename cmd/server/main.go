package main

import (
	"fmt"
	"github.com/toufiq-austcse/golang-grpc-gateway-test/pkg/cmd"
	"os"
)

func main() {
	if err := cmd.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
