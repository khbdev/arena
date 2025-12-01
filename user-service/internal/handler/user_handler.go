package handler

import (
	"context"
	
	"user-service/internal/model"
	"user-service/internal/service"
)
import userpb "github.com/khbdev/arena-proto-files/proto/user"


// ==========================
//         HANDLER
// ==========================
type UserHandler struct {
	svc *service.UserService
	userpb.UnimplementedUserServiceServer
}

// Constructor
func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

// ==========================
//      CREATE USER
// ==========================
func (h *UserHandler) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	dto := &service.CreateUserDTO{
		TelegramID: req.TelegramId,
		Firstname:  req.Firstname,
		Lastname:   req.Lastname,
		Role:       req.Role,
	}

	user, err := h.svc.CreateUser(dto)
	if err != nil {
		return nil, err
	}

	return &userpb.CreateUserResponse{
		User: mapUserToProto(user),
	}, nil
}

// ==========================
//      UPDATE USER
// ==========================
func (h *UserHandler) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	dto := &service.UpdateUserDTO{
		ID:        uint(req.Id),
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Role:      req.Role,
	}

	user, err := h.svc.UpdateUser(dto)
	if err != nil {
		return nil, err
	}

	return &userpb.UpdateUserResponse{
		User: mapUserToProto(user),
	}, nil
}

// ==========================
//      GET USER BY ID
// ==========================
func (h *UserHandler) GetUserByID(ctx context.Context, req *userpb.GetUserByIDRequest) (*userpb.GetUserByIDResponse, error) {
	user, err := h.svc.GetUserByID(uint(req.Id))
	if err != nil {
		return nil, err
	}

	return &userpb.GetUserByIDResponse{
		User: mapUserToProto(user),
	}, nil
}

// ==========================
//      LIST USERS
// ==========================
func (h *UserHandler) ListUsers(ctx context.Context, req *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	users, err := h.svc.ListUsers()
	if err != nil {
		return nil, err
	}

	protoUsers := []*userpb.User{}
	for _, u := range users {
		protoUsers = append(protoUsers, mapUserToProto(&u))
	}

	return &userpb.ListUsersResponse{Users: protoUsers}, nil
}

// ==========================
//      GET USER BY TELEGRAM ID
// ==========================
func (h *UserHandler) GetUserByTelegramID(ctx context.Context, req *userpb.GetUserByTelegramIDRequest) (*userpb.GetUserByTelegramIDResponse, error) {
	user, err := h.svc.GetUserByTelegramID(req.TelegramId)
	if err != nil {
		return nil, err
	}

	return &userpb.GetUserByTelegramIDResponse{
		User: mapUserToProto(user),
	}, nil
}

// ==========================
//      GET TELEGRAM IDS BY USER IDS
// ==========================
func (h *UserHandler) GetTelegramIDsByUserIDs(ctx context.Context, req *userpb.GetTelegramIDsByUserIDsRequest) (*userpb.GetTelegramIDsByUserIDsResponse, error) {
	ids := []uint{}
	for _, id := range req.UserIds {
		ids = append(ids, uint(id))
	}

	telegramIDs, err := h.svc.GetTelegramIDsByUserIDs(ids)
	if err != nil {
		return nil, err
	}

	return &userpb.GetTelegramIDsByUserIDsResponse{
		TelegramIds: telegramIDs,
	}, nil
}

// ==========================
//       HELPER FUNCTION
// ==========================
func mapUserToProto(user *model.User) *userpb.User {
	return &userpb.User{
		Id:         int64(user.ID),
		TelegramId: user.TelegramID,
		Firstname:  user.Firstname,
		Lastname:   user.Lastname,
		Role:       user.Role,
		CreatedAt:  user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:  user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
