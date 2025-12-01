package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	userpb "github.com/khbdev/arena-proto-files/proto/user"
	"google.golang.org/grpc"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: go run main.go <command> [args...]")
	}

	command := os.Args[1]

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := userpb.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	switch command {

	// ========== CREATE USER ==========
	case "create":
		if len(os.Args) != 6 {
			log.Fatalf("Usage: create <telegram_id> <firstname> <lastname> <role>")
		}
		tgID, _ := strconv.ParseInt(os.Args[2], 10, 64)
		req := &userpb.CreateUserRequest{
			TelegramId: tgID,
			Firstname:  os.Args[3],
			Lastname:   os.Args[4],
			Role:       os.Args[5],
		}
		res, err := client.CreateUser(ctx, req)
		if err != nil {
			log.Fatalf("CreateUser error: %v", err)
		}
		fmt.Printf("Created User: %+v\n", res.User)

	// ========== UPDATE USER ==========
	case "update":
		if len(os.Args) != 6 {
			log.Fatalf("Usage: update <id> <firstname> <lastname> <role>")
		}
		id, _ := strconv.ParseInt(os.Args[2], 10, 64)
		req := &userpb.UpdateUserRequest{
			Id:        id,
			Firstname: os.Args[3],
			Lastname:  os.Args[4],
			Role:      os.Args[5],
		}
		res, err := client.UpdateUser(ctx, req)
		if err != nil {
			log.Fatalf("UpdateUser error: %v", err)
		}
		fmt.Printf("Updated User: %+v\n", res.User)

	// ========== GET USER BY ID ==========
	case "get":
		if len(os.Args) != 3 {
			log.Fatalf("Usage: get <id>")
		}
		id, _ := strconv.ParseInt(os.Args[2], 10, 64)
		req := &userpb.GetUserByIDRequest{Id: id}
		res, err := client.GetUserByID(ctx, req)
		if err != nil {
			log.Fatalf("GetUserByID error: %v", err)
		}
		fmt.Printf("User: %+v\n", res.User)

	// ========== LIST USERS ==========
	case "list":
		req := &userpb.ListUsersRequest{}
		res, err := client.ListUsers(ctx, req)
		if err != nil {
			log.Fatalf("ListUsers error: %v", err)
		}
		for _, u := range res.Users {
			fmt.Printf("User: %+v\n", u)
		}

	// ========== GET BY TELEGRAM ID ==========
	case "getbytg":
		if len(os.Args) != 3 {
			log.Fatalf("Usage: getbytg <telegram_id>")
		}
		tgID, _ := strconv.ParseInt(os.Args[2], 10, 64)
		req := &userpb.GetUserByTelegramIDRequest{TelegramId: tgID}
		res, err := client.GetUserByTelegramID(ctx, req)
		if err != nil {
			log.Fatalf("GetUserByTelegramID error: %v", err)
		}
		fmt.Printf("User: %+v\n", res.User)

	// ========== GET TELEGRAM IDS BY USER IDS ==========
	case "tgids":
		if len(os.Args) < 3 {
			log.Fatalf("Usage: tgids <id1> <id2> ...")
		}
		ids := make([]int64, 0, len(os.Args)-2)
		for _, a := range os.Args[2:] {
			v, _ := strconv.ParseInt(a, 10, 64)
			ids = append(ids, v)
		}
		req := &userpb.GetTelegramIDsByUserIDsRequest{UserIds: ids}
		res, err := client.GetTelegramIDsByUserIDs(ctx, req)
		if err != nil {
			log.Fatalf("GetTelegramIDsByUserIDs error: %v", err)
		}
		fmt.Printf("Telegram IDs: %v\n", res.TelegramIds)

	default:
		log.Fatalf("Unknown command: %s", command)
	}
}
