// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package input

import (
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/domain/model"
)

// SysNoticeService 输入端口接口
type SysNoticeService interface {
	QueryNoticeByID(postId int64) (*model.SysNotice, error)
	QueryNoticeList(post *model.SysNoticeRequest) ([]*model.SysNotice, error)
	QueryNoticePage(pageReq common.PageRequest, r *model.SysNoticeRequest) ([]*model.SysNotice, int64, error)
	AddNotice(post *model.SysNotice) (*model.SysNotice, error)
	EditNotice(post *model.SysNotice) (*model.SysNotice, int64, error)
	DeleteNoticeById(id int64) (int64, error)
}
