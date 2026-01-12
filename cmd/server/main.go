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
	_ = db.InitDb()
	// defer database.Close() if we had a handle here, but InitDb currently manages its own lifecycle or we can return it

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
