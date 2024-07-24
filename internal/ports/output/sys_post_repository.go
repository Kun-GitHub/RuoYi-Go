// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package output

import (
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/domain/model"
)

type SysPostRepository interface {
	QueryPostByPostId(postId int64) (*model.SysPost, error)
	QueryPostByUserId(userId int64) ([]*model.SysPost, error)
	QueryPostList(post *model.SysPostRequest) ([]*model.SysPost, error)
	QueryPostPage(pageReq common.PageRequest, r *model.SysPostRequest) ([]*model.SysPost, int64, error)
	AddPost(post *model.SysPost) (*model.SysPost, error)
	EditPost(post *model.SysPost) (*model.SysPost, int64, error)
	DeletePostById(id int64) (int64, error)
}
