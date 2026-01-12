package main

import (
	"log"
	"net"
	"sports-betting-helper/internal/db"
	handler "sports-betting-helper/internal/handler/grpc"
	"sports-betting-helper/internal/service"
	pb "sports-betting-helper/proto/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	database := db.InitDb()
	defer database.Close()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	sportsbookService := service.NewSportsbookService()
	bettingHandler := handler.NewBettingHandler(sportsbookService)
	pb.RegisterBettingServiceServer(s, bettingHandler)
	reflection.Register(s)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
