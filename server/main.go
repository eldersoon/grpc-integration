/*
	protoc -I=proto 38tao.proto --js_out=import_style=commonjs:../client/grpc --grpc-web_out=import_style=commonjs,mode=grpcwebtext:../client/grpc

    =================================================================================

    protoc -I=proto 38tao.proto --js_out=import_style=commonjs:./client/src --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./client/src


    protoc -I /proto --go_out=. --go-grpc_out=. proto/*.proto

    go build -o ${BIN_DIR}/$@/${SERVER_BIN} ./$@/${SERVER_DIR}
	go build -o ${BIN_DIR}/$@/${CLIENT_BIN} ./$@/${CLIENT_DIR

    grpcwebproxy.exe --backend_addr=localhost:50051 --run_tls_server=false --allow_all_origins

*/

package main

import (
	"context"
	"log"
	"net"
	"strconv"

	pb "github.com/eldersoon/trezoitao/proto/trezoitao/proto"

	"google.golang.org/grpc"
)

type server struct {
    pb.UnimplementedTrezoitaoServiceServer
}

func (s *server) GetBTC(ctx context.Context, req *pb.GiveMeBTCRequest) (*pb.GiveMeBTCResponse, error) {
    return &pb.GiveMeBTCResponse{
        Message: "Vc fez uma transferência satânica de " + strconv.Itoa(int(req.Amount))  + "$ hein,hein? Eoguê?",
        Status:  "success",
    }, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()
    pb.RegisterTrezoitaoServiceServer(s, &server{})

    log.Println("Server is running on port :50051")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
