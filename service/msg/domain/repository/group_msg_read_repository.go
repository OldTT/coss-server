package repository

import "github.com/cossim/coss-server/service/msg/domain/entity"

type GroupMsgReadRepository interface {
	// GetGroupMsgReadByDialogID 获取群消息阅读状态
	GetGroupMsgReadByMsgID(msgId uint32) ([]*entity.GroupMessageRead, error)
	// SetGroupMsgReadByDialogID 设置群消息阅读状态
	SetGroupMsgReadByMsgs(read []*entity.GroupMessageRead) error
	GetGroupMsgReadUserIdsByMsgID(msgId uint32) ([]string, error)
	GetGroupMsgReadByMsgIDAndUserID(msgId uint32, userId string) (*entity.GroupMessageRead, error)
}
