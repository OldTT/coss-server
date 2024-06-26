package persistence

import (
	"github.com/cossim/coss-server/internal/group/domain/entity"
	"github.com/cossim/coss-server/pkg/utils/time"
	"gorm.io/gorm"
)

type GroupRepo struct {
	db *gorm.DB
}

func NewGroupRepo(db *gorm.DB) *GroupRepo {
	return &GroupRepo{db: db}
}

func (repo *GroupRepo) GetGroupInfoByGid(gid uint) (*entity.Group, error) {
	var group entity.Group
	result := repo.db.First(&group, gid)
	if result.Error != nil {
		return nil, result.Error
	}
	return &group, nil
}

func (repo *GroupRepo) GetBatchGetGroupInfoByIDs(groupIds []uint) ([]*entity.Group, error) {
	var groups []*entity.Group
	result := repo.db.Find(&groups, groupIds)
	if result.Error != nil {
		return nil, result.Error
	}
	return groups, nil
}

func (repo *GroupRepo) UpdateGroup(group *entity.Group) (*entity.Group, error) {
	result := repo.db.Where("id = ?", group.ID).Updates(group)
	if result.Error != nil {
		return nil, result.Error
	}
	return group, nil
}

func (repo *GroupRepo) InsertGroup(group *entity.Group) (*entity.Group, error) {
	result := repo.db.Create(group)
	if result.Error != nil {
		return nil, result.Error
	}
	return group, nil
}

func (repo *GroupRepo) DeleteGroup(gid uint) error {
	err := repo.db.Model(&entity.Group{}).Where("id = ?", gid).Update("status", entity.GroupStatusDeleted).Update("deleted_at", time.Now()).Error
	return err
}

func (repo *GroupRepo) UpdateGroupByGroupID(gid uint, updateFields map[string]interface{}) error {
	return repo.db.Model(&entity.Group{}).Where("id = ?", gid).Unscoped().Updates(updateFields).Error
}
