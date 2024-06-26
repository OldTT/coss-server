package grpc

import (
	"context"
	"fmt"
	v1 "github.com/cossim/coss-server/internal/relation/api/grpc/v1"
	"github.com/cossim/coss-server/internal/relation/domain/entity"
	"github.com/cossim/coss-server/internal/relation/domain/repository"
	"github.com/cossim/coss-server/internal/relation/infrastructure/persistence"
	"github.com/cossim/coss-server/pkg/cache"
	"github.com/cossim/coss-server/pkg/code"
	ptime "github.com/cossim/coss-server/pkg/utils/time"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

var _ v1.DialogServiceServer = &dialogServiceServer{}

type dialogServiceServer struct {
	db          *gorm.DB
	cache       cache.RelationUserCache
	cacheEnable bool
	dr          repository.DialogRepository
}

func (s *dialogServiceServer) CreateDialog(ctx context.Context, in *v1.CreateDialogRequest) (*v1.CreateDialogResponse, error) {
	resp := &v1.CreateDialogResponse{}

	dialog, err := s.dr.CreateDialog(in.OwnerId, entity.DialogType(in.Type), uint(in.GroupId))
	if err != nil {
		return resp, status.Error(codes.Code(code.DialogErrCreateDialogFailed.Code()), err.Error())
	}
	return &v1.CreateDialogResponse{
		Id:      uint32(dialog.ID),
		OwnerId: dialog.OwnerId,
		GroupId: uint32(dialog.ID),
		Type:    uint32(dialog.Type),
	}, nil
}

func (s *dialogServiceServer) CreateAndJoinDialogWithGroup(ctx context.Context, request *v1.CreateAndJoinDialogWithGroupRequest) (*v1.CreateAndJoinDialogWithGroupResponse, error) {
	resp := &v1.CreateAndJoinDialogWithGroupResponse{}

	//return resp, status.Error(codes.Aborted, formatErrorMessage(errors.New("测试回滚")))

	ids := []string{request.OwnerId}
	for _, v := range request.UserIds {
		ids = append(ids, v)
	}

	if err := s.db.Transaction(func(tx *gorm.DB) error {
		npo := persistence.NewRepositories(tx)
		dialog, err := npo.Dr.CreateDialog(request.OwnerId, entity.DialogType(request.Type), uint(request.GroupId))
		if err != nil {
			return status.Error(codes.Code(code.DialogErrCreateDialogFailed.Code()), fmt.Sprintf("failed to create dialog: %s", err.Error()))
		}
		_, err = npo.Dr.JoinBatchDialog(dialog.ID, ids)
		if err != nil {
			return status.Error(codes.Code(code.DialogErrJoinDialogFailed.Code()), fmt.Sprintf("failed to join dialog: %s", err.Error()))
		}
		resp.Id = uint32(dialog.ID)
		resp.OwnerId = dialog.OwnerId
		resp.GroupId = uint32(dialog.ID)
		resp.Type = uint32(dialog.Type)
		return nil
	}); err != nil {
		return resp, status.Error(codes.Aborted, fmt.Sprintf("failed to create dialog: %s", err.Error()))
	}

	return resp, nil
}

func (s *dialogServiceServer) CreateAndJoinDialogWithGroupRevert(ctx context.Context, request *v1.CreateAndJoinDialogWithGroupRequest) (*v1.CreateAndJoinDialogWithGroupResponse, error) {
	resp := &v1.CreateAndJoinDialogWithGroupResponse{}

	fmt.Println("CreateAndJoinDialogWithGroupRevert req => ", request)

	ids := []string{request.OwnerId}
	for _, id := range request.UserIds {
		ids = append(ids, id)
	}

	if err := s.db.Transaction(func(tx *gorm.DB) error {
		npo := persistence.NewRepositories(tx)
		if err := npo.Dr.DeleteDialogUserByDialogIDAndUserID(uint(request.DialogId), ids); err != nil {
			return status.Error(codes.Code(code.DialogErrDeleteDialogUsersFailed.Code()), fmt.Sprintf("failed to delete dialog user revert : %s", err.Error()))
		}
		if err := npo.Dr.DeleteDialogByDialogID(uint(request.DialogId)); err != nil {
			return status.Error(codes.Code(code.DialogErrDeleteDialogFailed.Code()), fmt.Sprintf("failed to delete dialog revert : %s", err.Error()))
		}

		return nil
	}); err != nil {
		return resp, status.Error(codes.Aborted, fmt.Sprintf("failed to create dialog revert: %s", err.Error()))
	}

	return resp, nil
}

func (s *dialogServiceServer) ConfirmFriendAndJoinDialog(ctx context.Context, request *v1.ConfirmFriendAndJoinDialogRequest) (*v1.ConfirmFriendAndJoinDialogResponse, error) {
	resp := &v1.ConfirmFriendAndJoinDialogResponse{}

	if err := s.db.Transaction(func(tx *gorm.DB) error {
		npo := persistence.NewRepositories(tx)
		dialog, err := npo.Dr.CreateDialog(request.OwnerId, entity.DialogType(request.Type), uint(request.GroupId))
		if err != nil {
			return status.Error(codes.Code(code.DialogErrCreateDialogFailed.Code()), err.Error())
		}

		_, err = npo.Dr.JoinDialog(dialog.ID, request.OwnerId)
		if err != nil {
			return status.Error(codes.Code(code.DialogErrJoinDialogFailed.Code()), err.Error())
		}

		_, err = s.dr.JoinDialog(dialog.ID, request.UserId)
		if err != nil {
			return status.Error(codes.Code(code.DialogErrJoinDialogFailed.Code()), err.Error())
		}
		resp.Id = uint32(dialog.ID)
		resp.OwnerId = request.OwnerId
		resp.Type = v1.DialogType(dialog.Type)
		resp.GroupId = uint32(dialog.GroupId)

		return nil
	}); err != nil {
		return resp, err
	}

	return resp, nil
}

func (s *dialogServiceServer) ConfirmFriendAndJoinDialogRevert(ctx context.Context, request *v1.ConfirmFriendAndJoinDialogRevertRequest) (*v1.ConfirmFriendAndJoinDialogRevertResponse, error) {
	resp := &v1.ConfirmFriendAndJoinDialogRevertResponse{}

	if err := s.dr.DeleteDialogUserByDialogIDAndUserID(uint(request.DialogId), []string{request.UserId}); err != nil {
		return resp, status.Error(codes.Code(code.DialogErrCreateDialogFailed.Code()), err.Error())
	}

	if err := s.dr.DeleteDialogUserByDialogIDAndUserID(uint(request.DialogId), []string{request.OwnerId}); err != nil {
		return resp, status.Error(codes.Code(code.DialogErrCreateDialogFailed.Code()), err.Error())
	}

	if err := s.dr.DeleteDialogByDialogID(uint(request.DialogId)); err != nil {
		return resp, status.Error(codes.Code(code.DialogErrCreateDialogFailed.Code()), err.Error())
	}

	return resp, nil
}

func (s *dialogServiceServer) JoinDialog(ctx context.Context, in *v1.JoinDialogRequest) (*v1.JoinDialogResponse, error) {
	resp := &v1.JoinDialogResponse{}
	_, err := s.dr.JoinDialog(uint(in.DialogId), in.UserId)
	if err != nil {
		return resp, status.Error(codes.Code(code.DialogErrJoinDialogFailed.Code()), err.Error())
	}
	return resp, nil
}

func (s *dialogServiceServer) JoinDialogRevert(ctx context.Context, request *v1.JoinDialogRequest) (*v1.JoinDialogResponse, error) {
	resp := &v1.JoinDialogResponse{}
	if err := s.dr.DeleteDialogUserByDialogIDAndUserID(uint(request.DialogId), []string{request.UserId}); err != nil {
		return resp, status.Error(codes.Code(code.DialogErrJoinDialogFailed.Code()), err.Error())
	}
	return resp, nil
}

func (s *dialogServiceServer) GetUserDialogList(ctx context.Context, in *v1.GetUserDialogListRequest) (*v1.GetUserDialogListResponse, error) {
	resp := &v1.GetUserDialogListResponse{}
	ids, total, err := s.dr.GetUserDialogs(in.UserId, int(in.PageSize), int(in.PageNum))
	if err != nil {
		return resp, status.Error(codes.Code(code.DialogErrGetUserDialogListFailed.Code()), err.Error())
	}
	nids := make([]uint32, 0)
	if len(ids) > 0 {
		for _, id := range ids {
			nids = append(nids, uint32(id))
		}
	}
	resp.Total = uint64(total)
	resp.DialogIds = nids
	return resp, nil
}

func (s *dialogServiceServer) GetDialogByIds(ctx context.Context, in *v1.GetDialogByIdsRequest) (*v1.GetDialogByIdsResponse, error) {
	resp := &v1.GetDialogByIdsResponse{}
	nids := make([]uint, 0)
	if len(in.DialogIds) > 0 {
		for _, id := range in.DialogIds {
			nids = append(nids, uint(id))
		}
	}
	infos, err := s.dr.GetDialogsByIDs(nids)
	if err != nil {
		return resp, status.Error(codes.Code(code.DialogErrGetUserDialogListFailed.Code()), err.Error())
	}
	var dialogInfos []*v1.Dialog
	if len(infos) > 0 {
		for _, info := range infos {
			dialogInfos = append(dialogInfos, &v1.Dialog{
				Id:       uint32(info.ID),
				OwnerId:  info.OwnerId,
				GroupId:  uint32(info.GroupId),
				Type:     uint32(info.Type),
				CreateAt: info.CreatedAt,
			})
		}
	}
	resp.Dialogs = dialogInfos
	return resp, nil
}

func (s *dialogServiceServer) GetDialogUsersByDialogID(ctx context.Context, in *v1.GetDialogUsersByDialogIDRequest) (*v1.GetDialogUsersByDialogIDResponse, error) {
	resp := &v1.GetDialogUsersByDialogIDResponse{}
	users, err := s.dr.GetDialogUsersByDialogID(uint(in.DialogId))
	if err != nil {
		return resp, status.Error(codes.Code(code.DialogErrGetUserDialogListFailed.Code()), err.Error())
	}
	var ids []string
	for _, id := range users {
		ids = append(ids, id.UserId)
	}
	resp.UserIds = ids
	return resp, nil
}

func (s *dialogServiceServer) DeleteDialogByIds(ctx context.Context, in *v1.DeleteDialogByIdsRequest) (*v1.DeleteDialogByIdsResponse, error) {
	var resp = &v1.DeleteDialogByIdsResponse{}
	uintIds := make([]uint, 0)
	if len(in.DialogIds) > 0 {
		for _, id := range in.DialogIds {
			uintIds = append(uintIds, uint(id))
		}
	}
	if err := s.dr.DeleteDialogByIds(uintIds); err != nil {
		return resp, status.Error(codes.Code(code.DialogErrDeleteDialogFailed.Code()), err.Error())
	}
	return resp, nil
}

func (s *dialogServiceServer) DeleteDialogById(ctx context.Context, in *v1.DeleteDialogByIdRequest) (*v1.DeleteDialogByIdResponse, error) {
	resp := &v1.DeleteDialogByIdResponse{}

	//return resp, status.Error(codes.Aborted, formatErrorMessage(errors.New("测试回滚")))

	if err := s.dr.DeleteDialogByDialogID(uint(in.DialogId)); err != nil {
		return resp, status.Error(codes.Code(code.DialogErrDeleteDialogFailed.Code()), err.Error())
	}
	return resp, nil
}

func (s *dialogServiceServer) DeleteDialogByIdRevert(ctx context.Context, request *v1.DeleteDialogByIdRequest) (*v1.DeleteDialogByIdResponse, error) {
	var resp = &v1.DeleteDialogByIdResponse{}
	fmt.Println("DeleteDialogByIdRevert req => ", request)
	if err := s.dr.UpdateDialogByDialogID(uint(request.DialogId), map[string]interface{}{
		"deleted_at": 0,
	}); err != nil {
		return resp, status.Error(codes.Code(code.DialogErrDeleteDialogFailed.Code()), err.Error())
	}
	return resp, nil
}

func (s *dialogServiceServer) DeleteDialogUsersByDialogID(ctx context.Context, in *v1.DeleteDialogUsersByDialogIDRequest) (*v1.DeleteDialogUsersByDialogIDResponse, error) {
	var resp = &v1.DeleteDialogUsersByDialogIDResponse{}
	if err := s.dr.DeleteDialogUserByDialogID(uint(in.DialogId)); err != nil {
		return resp, status.Error(codes.Code(code.DialogErrDeleteDialogUsersFailed.Code()), err.Error())
	}
	return resp, nil
}

func (s *dialogServiceServer) DeleteDialogUsersByDialogIDRevert(ctx context.Context, request *v1.DeleteDialogUsersByDialogIDRequest) (*v1.DeleteDialogUsersByDialogIDResponse, error) {
	var resp = &v1.DeleteDialogUsersByDialogIDResponse{}
	if err := s.dr.UpdateDialogUserByDialogID(uint(request.DialogId), map[string]interface{}{
		"deleted_at": 0,
	}); err != nil {
		return resp, status.Error(codes.Code(code.DialogErrDeleteDialogUsersFailed.Code()), err.Error())
	}
	return resp, nil
}

func (s *dialogServiceServer) GetDialogUserByDialogIDAndUserID(ctx context.Context, in *v1.GetDialogUserByDialogIDAndUserIdRequest) (*v1.GetDialogUserByDialogIDAndUserIdResponse, error) {
	var resp = &v1.GetDialogUserByDialogIDAndUserIdResponse{}
	user, err := s.dr.GetDialogUserByDialogIDAndUserID(uint(in.DialogId), in.UserId)
	if err != nil {
		return resp, status.Error(codes.Code(code.DialogErrGetDialogUserByDialogIDAndUserIDFailed.Code()), err.Error())
	}
	resp.DialogId = uint32(user.DialogId)
	resp.UserId = user.UserId
	resp.IsShow = uint32(user.IsShow)
	resp.TopAt = uint64(user.TopAt)

	return resp, nil
}

func (s *dialogServiceServer) DeleteDialogUserByDialogIDAndUserID(ctx context.Context, request *v1.DeleteDialogUserByDialogIDAndUserIDRequest) (*v1.DeleteDialogUserByDialogIDAndUserIDResponse, error) {
	var resp = &v1.DeleteDialogUserByDialogIDAndUserIDResponse{}

	err := s.dr.DeleteDialogUserByDialogIDAndUserID(uint(request.DialogId), []string{request.UserId})
	if err != nil {
		return resp, status.Error(codes.Code(code.DialogErrGetDialogUserByDialogIDAndUserIDFailed.Code()), err.Error())
	}
	return resp, nil
}

func (s *dialogServiceServer) DeleteDialogUserByDialogIDAndUserIDRevert(ctx context.Context, request *v1.DeleteDialogUserByDialogIDAndUserIDRequest) (*v1.DeleteDialogUserByDialogIDAndUserIDResponse, error) {
	resp := &v1.DeleteDialogUserByDialogIDAndUserIDResponse{}

	if err := s.dr.UpdateDialogUserColumnByDialogIDAndUserId(uint(request.DialogId), request.UserId, "deleted_at", 0); err != nil {
		return resp, status.Error(codes.Code(code.DialogErrGetDialogUserByDialogIDAndUserIDFailed.Code()), err.Error())
	}

	return resp, nil
}

func (s *dialogServiceServer) GetDialogByGroupId(ctx context.Context, in *v1.GetDialogByGroupIdRequest) (*v1.GetDialogByGroupIdResponse, error) {
	var resp = &v1.GetDialogByGroupIdResponse{}
	dialog, err := s.dr.GetDialogByGroupId(uint(in.GroupId))
	if err != nil {
		return resp, err
	}
	resp.DialogId = uint32(dialog.ID)
	resp.GroupId = uint32(dialog.GroupId)
	resp.Type = uint32(dialog.Type)
	resp.CreateAt = dialog.CreatedAt
	return resp, nil
}

func (s *dialogServiceServer) GetDialogByGroupIds(ctx context.Context, in *v1.GetDialogByGroupIdsRequest) (*v1.GetDialogByGroupIdsResponse, error) {
	var resp = &v1.GetDialogByGroupIdsResponse{}
	var idlist []uint
	if len(in.GroupId) > 0 {
		for _, id := range in.GroupId {
			idlist = append(idlist, uint(id))
		}
	}

	ids, err := s.dr.GetDialogByGroupIds(idlist)
	if err != nil {
		return resp, err
	}

	if len(ids) > 0 {
		for _, id := range ids {
			resp.Dialogs = append(resp.Dialogs, &v1.GetDialogByGroupIdResponse{
				DialogId: uint32(id.ID),
				GroupId:  uint32(id.GroupId),
			})
		}
	}
	return resp, nil
}

func (s *dialogServiceServer) CloseOrOpenDialog(ctx context.Context, in *v1.CloseOrOpenDialogRequest) (*emptypb.Empty, error) {
	var resp = &emptypb.Empty{}
	if err := s.dr.UpdateDialogUserColumnByDialogIDAndUserId(uint(in.DialogId), in.UserId, "is_show", in.Action); err != nil {
		return resp, status.Error(codes.Code(code.DialogErrCloseOrOpenDialogFailed.Code()), err.Error())
	}
	return resp, nil
}

func (s *dialogServiceServer) TopOrCancelTopDialog(ctx context.Context, in *v1.TopOrCancelTopDialogRequest) (*emptypb.Empty, error) {
	var resp = &emptypb.Empty{}
	switch in.Action {
	case v1.TopOrCancelTopDialogType_CANCEL_TOP:
		if err := s.dr.UpdateDialogUserColumnByDialogIDAndUserId(uint(in.DialogId), in.UserId, "top_at", 0); err != nil {
			return resp, status.Error(codes.Code(code.DialogErrCloseOrOpenDialogFailed.Code()), err.Error())
		}
	case v1.TopOrCancelTopDialogType_TOP:
		if err := s.dr.UpdateDialogUserColumnByDialogIDAndUserId(uint(in.DialogId), in.UserId, "top_at", ptime.Now()); err != nil {
			return resp, status.Error(codes.Code(code.DialogErrCloseOrOpenDialogFailed.Code()), err.Error())
		}
	}

	return resp, nil
}

func (s *dialogServiceServer) GetDialogById(ctx context.Context, in *v1.GetDialogByIdRequest) (*v1.Dialog, error) {
	var resp = &v1.Dialog{}
	dialog, err := s.dr.GetDialogById(uint(in.DialogId))
	if err != nil {
		return resp, status.Error(codes.Code(code.DialogErrGetDialogByIdFailed.Code()), err.Error())
	}
	resp.Id = uint32(dialog.ID)
	resp.GroupId = uint32(dialog.GroupId)
	resp.Type = uint32(dialog.Type)
	resp.CreateAt = dialog.CreatedAt
	resp.OwnerId = dialog.OwnerId

	return resp, nil
}

func (s *dialogServiceServer) GetAllUsersInConversation(ctx context.Context, in *v1.GetAllUsersInConversationRequest) (*v1.GetAllUsersInConversationResponse, error) {
	resp := &v1.GetAllUsersInConversationResponse{}
	users, err := s.dr.GetDialogAllUsers(uint(in.DialogId))
	if err != nil {
		return resp, status.Error(codes.Code(code.DialogErrGetUserDialogListFailed.Code()), err.Error())
	}
	var ids []string
	for _, id := range users {
		ids = append(ids, id.UserId)
	}
	resp.UserIds = ids
	return resp, nil
}

func (s *dialogServiceServer) BatchCloseOrOpenDialog(ctx context.Context, request *v1.BatchCloseOrOpenDialogRequest) (*emptypb.Empty, error) {
	resp := &emptypb.Empty{}
	err := s.dr.UpdateDialogUserByDialogIDAndUserIds(uint(request.DialogId), request.UserIds, "is_show", request.Action)
	if err != nil {
		return resp, status.Error(codes.Code(code.DialogErrCloseOrOpenDialogFailed.Code()), err.Error())
	}
	return resp, nil
}

func (s *dialogServiceServer) GetDialogTargetUserId(ctx context.Context, request *v1.GetDialogTargetUserIdRequest) (*v1.GetDialogTargetUserIdResponse, error) {
	resp := &v1.GetDialogTargetUserIdResponse{}
	userIDs, err := s.dr.GetDialogTargetUserId(uint(request.DialogId), request.UserId)
	if err != nil {
		return resp, status.Error(codes.Code(code.DialogErrGetTargetIdFailed.Code()), err.Error())
	}
	resp.UserIds = userIDs
	return resp, nil
}
