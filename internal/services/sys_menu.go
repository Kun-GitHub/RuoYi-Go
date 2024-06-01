package services

import (
	"RuoYi-Go/internal/models"
	"gorm.io/gorm"
)

func GetMenus(db *gorm.DB) ([]models.SysMenu, error) {
	var menus []models.SysMenu
	err := db.Table("sys_menu m").
		Select("distinct m.menu_id, m.parent_id, m.menu_name, m.path, m.component, m.query, m.visible, m.status, IFNULL(m.perms, '') as perms, m.is_frame, m.is_cache, m.menu_type, m.icon, m.order_num, m.create_time").
		Where("m.menu_type IN (?) AND m.status = ?", []string{"M", "C"}, 0).
		Order("m.parent_id, m.order_num").
		Find(&menus).Error
	if err != nil {
		return nil, err
	}
	return menus, nil
}
