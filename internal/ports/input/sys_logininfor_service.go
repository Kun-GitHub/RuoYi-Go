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
type SysLogininforService interface {
	QueryLogininforByID(id int64) (*model.SysLogininfor, error)
	QueryLogininforList(request *model.SysLogininforRequest) ([]*model.SysLogininfor, error)
	QueryLogininforPage(pageReq common.PageRequest, r *model.SysLogininforRequest) ([]*model.SysLogininfor, int64, error)
	AddLogininfor(post *model.SysLogininfor) (*model.SysLogininfor, error)
	EditLogininfor(post *model.SysLogininfor) (*model.SysLogininfor, int64, error)
	DeleteLogininforById(id int64) (int64, error)
}
