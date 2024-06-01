package services

import (
	"RuoYi-Go/internal/models"
	rydb "RuoYi-Go/pkg/db"
	"gorm.io/gorm"
)

func GetAllMenus() ([]models.SysMenu, error) {
	var menus []models.SysMenu
	err := rydb.DB.Transactional(func(db *gorm.DB) error {
		err := db.Table("sys_menu AS m").Select("m.menu_id, m.parent_id, m.menu_name, m.path, m.component, m.query, m.visible, m.status, COALESCE(m.perms, '') AS perms, m.is_frame, m.is_cache, m.menu_type, m.icon, m.order_num, m.create_time").
			Where("m.menu_type IN (?, ?) AND m.status = '0'", "M", "C").
			Order("m.parent_id, m.order_num").
			Scan(&menus).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return menus, nil
}
