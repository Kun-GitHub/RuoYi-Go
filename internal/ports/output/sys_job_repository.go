// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package output

import (
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/domain/model"
)

type SysJobRepository interface {
	QueryJobByID(id int64) (*model.SysJob, error)
	QueryJobList(request *model.SysJobRequest) ([]*model.SysJob, error)
	QueryJobPage(pageReq common.PageRequest, r *model.SysJobRequest) ([]*model.SysJob, int64, error)
	AddJob(post *model.SysJob) (*model.SysJob, error)
	EditJob(post *model.SysJob) (*model.SysJob, int64, error)
	DeleteJobById(id int64) (int64, error)
	ChangeJobStatus(user *model.ChangeJobStatusRequest) (int64, error)
}
