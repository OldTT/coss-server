package repository

import "github.com/cossim/coss-server/internal/relation/domain/entity"

type GroupRelationRepository interface {
	CreateRelation(ur *entity.GroupRelation) (*entity.GroupRelation, error)
	CreateRelations(urs []*entity.GroupRelation) ([]*entity.GroupRelation, error)
	UpdateRelation(ur *entity.GroupRelation) (*entity.GroupRelation, error)
	UpdateRelationColumnByGroupAndUser(gid uint32, uid string, column string, value interface{}) error
	DeleteRelationByID(gid uint32, uid string) error
	DeleteUserGroupRelationByGroupIDAndUserID(gid uint32, uid string) error

	// GetGroupUserIDs 获取群聊所有的用户ID
	GetGroupUserIDs(gid uint32) ([]string, error)
	GetUserGroupIDs(uid string) ([]uint32, error)
	GetUserGroupByID(gid uint32, uid string) (*entity.GroupRelation, error)
	GetUserGroupByIDs(gid uint32, uids []string) ([]*entity.GroupRelation, error)
	GetUserJoinedGroupIDs(uid string) ([]uint32, error) // 获取用户加入的所有群聊ID
	GetUserManageGroupIDs(uid string) ([]uint32, error) // 获取用户管理的或创建的群聊ID
	GetJoinRequestBatchListByID(gids []uint32) ([]*entity.GroupRelation, error)
	GetJoinRequestListByID(gid uint32) ([]*entity.GroupRelation, error)
	DeleteGroupRelationByID(gid uint32) error
	GetGroupAdminIds(gid uint32) ([]string, error) // 获取群聊管理员
	SetUserGroupRemark(gid uint32, uid string, remark string) error

	// UpdateGroupRelationByGroupID 根据群聊id更新群聊的所有用户信息
	UpdateGroupRelationByGroupID(dialogID uint32, updateFields map[string]interface{}) error
	DeleteRelationByGroupIDAndUserIDs(gid uint32, uid []string) error

	SetUserGroupSilentNotification(gid uint32, uid string, silentNotification entity.SilentNotification) error
	SetUserGroupOpenBurnAfterReading(gid uint32, uid string, openBurnAfterReading entity.OpenBurnAfterReadingType) error
	SetUserGroupOpenBurnAfterReadingTimeOUt(gid uint32, uid string, burnAfterReadingTimeOut int64) error
}
