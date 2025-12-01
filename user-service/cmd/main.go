package main


import (
	"log"
	"net"
	"os"
	"user-service/internal/config"
	"user-service/internal/handler"
	"user-service/internal/repostroy"
	"user-service/internal/service"

	userpb "github.com/khbdev/arena-proto-files/proto/user"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	// ENV load
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	config.InitDB()
	config.InitRedis()

	userRepo := repostroy.NewUserRepository(config.DB)
	userSvc := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userSvc)

	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, userHandler)

	// GRPC port from env
	port := os.Getenv("GRPC_PORT")
	if port == "" {
		port = "50051" // default
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("gRPC server listening on port", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
