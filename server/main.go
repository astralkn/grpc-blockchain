package main

import (
	context "context"
	"google.golang.org/grpc"
	"grpc-server/proto"
	"grpc-server/server/blockchain"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("unable to listen on 8080 port: %v", err)
	}

	srv := grpc.NewServer()

	proto.RegisterBlockChainServer(srv, &Server{BlockChain: blockchain.NewBlockChain()})
	err = srv.Serve(listener)
	if err != nil {
		log.Fatalf("server encounterd an error: %v", err)
	}
}

type Server struct {
	BlockChain *blockchain.BlockChain
}

func (s *Server) AddBlock(ctx context.Context, request *proto.AddBlockRequest) (*proto.AddBlockResponse, error) {
	block := s.BlockChain.AddBlock(request.Data)
	return &proto.AddBlockResponse{
		Hash: block.Hash,
	}, nil
}

func (s *Server) GetBlockChain(ctx context.Context, request *proto.GetBlockChainRequest) (*proto.GetBlockChainResponse, error) {
	resp := new(proto.GetBlockChainResponse)
	for _, block := range s.BlockChain.Blocks {
		resp.Blocks = append(resp.Blocks, &proto.Block{
			Hash:              block.Hash,
			PreviousBlockHash: block.PrevBlockHash,
			Data:              block.Data,
		})
	}
	return resp, nil
}
