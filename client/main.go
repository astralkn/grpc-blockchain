package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-server/proto"
	"log"
	"time"
)

var client proto.BlockChainClient

func main() {
	addFlag := flag.Bool("add", false, "add new block")
	listFlag := flag.Bool("list", false, "get the blockchain")
	flag.Parse()
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("can not dial server: %v", err)
	}
	client = proto.NewBlockChainClient(conn)

	if *addFlag {
		addBlock()
	}
	if *listFlag {
		getBlockChain()
	}
}

func getBlockChain() {
	block, err := client.GetBlockChain(context.Background(), &proto.GetBlockChainRequest{})
	if err != nil {
		log.Fatalf("unable to get blockchain: %v \n", err)
	}
	log.Println("blocks:")
	for _, b := range block.Blocks {
		log.Printf("\n\thash: %s\n\tpre_hash: %s\n\tdata: %s\n", b.Hash, b.PreviousBlockHash, b.Data)
	}
}

func addBlock() {
	block, err := client.AddBlock(context.Background(), &proto.AddBlockRequest{Data: time.Now().String()})
	if err != nil {
		log.Fatalf("unable to add block: %v \n", err)
	}
	log.Printf("new block hash: %s\n", block.Hash)
}
