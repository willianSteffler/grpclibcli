package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/fullstorydev/grpcui/standalone"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
)

type Conf struct {
	socketPort int
	grpcPort int
	grpcUiPort int
}

var conf Conf

func main() {

	flag.IntVar(&conf.grpcPort, "grpc", 3100, "porta grpc")
	flag.IntVar(&conf.grpcUiPort, "grpcui", 3101, "porta pagina web")


	Usage := func() {
		fmt.Fprintf(os.Stderr, "Sintaxe do comando %s: [opções] \n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Usage = Usage
	flag.Parse()

	cc, err := grpc.Dial(fmt.Sprintf("0.0.0.0:%d", conf.grpcPort),grpc.WithInsecure())
	if err != nil {
		log.Fatalf("falha ao criar cliente: %v", err)
	}

	target := fmt.Sprintf("0.0.0.0:%d", conf.grpcUiPort)
	h, err := standalone.HandlerViaReflection(context.Background(), cc, target)
	if err != nil {
		log.Fatalf("falha ao criar pagina web : %v", err)
	}
	serveMux := http.NewServeMux()

	serveMux.Handle("/", h)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d",conf.grpcUiPort), serveMux))

}