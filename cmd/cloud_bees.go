package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ondbyte/cloud_bees/client"
	"github.com/ondbyte/cloud_bees/server"
	flag "github.com/ondbyte/turbo_flag"
)

func main() {
	log.Println("pid:", os.Getpid())
	flag.MainCmd("cloud_bees", "use this to start either server or client instance", flag.ContinueOnError, os.Args[1:], func(cloudBeesCmd flag.CMD, args []string) {
		cloudBeesCmd.SubCmd("server", "run the server instance", func(serverCmd flag.CMD, args []string) {
			port := serverCmd.Uint("port", 8081, "port where this instance should run")
			err := serverCmd.Parse(args)
			if err != nil {
				panic(err)
			}
			log.Printf("running server on port %d\n", *port)
			err = server.Run(*port)
			if err != nil {
				log.Fatalf("error running server %v", err)
			}
		})
		cloudBeesCmd.SubCmd("client", "run the client instance", func(clientCmd flag.CMD, args []string) {
			port := clientCmd.Uint("port", 8081, "port where this client should connect to")
			err := clientCmd.Parse(args)
			if err != nil {
				log.Fatal(err)
			}
			err = client.Start(*port)
			if err != nil {
				log.Fatal(err)
			}
		})
		help := cloudBeesCmd.Bool("help", false, "print help")
		err := cloudBeesCmd.Parse(args)
		if err != nil {
			log.Fatal(err)
		}
		if *help {
			fmt.Println(cloudBeesCmd.GetDefaultUsageLong())
		} else {
			fmt.Println("cloud_bees exited, run with '-help' to see help")
		}
	})
}

func RunClient() {

}
