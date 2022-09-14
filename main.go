package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/common-nighthawk/go-figure"
	"github.com/libp2p/go-libp2p-core/host"
)

type Config struct {
	Port int
	Seed int64
}

func main() {
	config := Config{}

	flag.Int64Var(&config.Seed, "seed", 0, "Seed value for generating a PeerID, 0 is random")
	flag.IntVar(&config.Port, "port", 0, "")
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())

	h, err := NewHost(ctx, config.Seed, config.Port)
	if err != nil {
		log.Fatal(err)
	}

	art()

	for _, addr := range h.Addrs() {
		log.Printf("  %s/p2p/%s", addr, h.ID().Pretty())
	}

	_, err = NewKDHT(ctx, h)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("")
	log.Println(" ✅️ Kademlia DHT bootstrap node active.. use one of the above address to connect to it")

	waitSignal(h, cancel)
}

func waitSignal(h host.Host, cancel func()) {
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	<-c

	fmt.Printf("\r👋️ stopped...\n")

	cancel()

	if err := h.Close(); err != nil {
		panic(err)
	}
	os.Exit(0)
}

func art() {
	fmt.Println()
	fmt.Println("dedicated Kademlia DHT bootstrap node")
	myFigure := figure.NewColorFigure("KAD DHT boot", "", "yellow", true)
	myFigure.Print()
	fmt.Println()
}
