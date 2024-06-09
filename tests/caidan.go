package main

import (
	"fmt"
)

// SysMenu represents a system menu.
type SysMenu struct {
	ID       int
	Name     string
	Path     string
	ParentID int
	Children []*SysMenu
}

// BuildMenuTree builds the menu tree from a flat list of SysMenu.
func BuildMenuTree(menus []*SysMenu) []*SysMenu {
	menuMap := make(map[int]*SysMenu)
	rootMenus := make([]*SysMenu, 0)

	// Fill the map with all menus
	for _, menu := range menus {
		menuMap[menu.ID] = menu
	}

	// Construct the tree
	for _, menu := range menus {
		if menu.ParentID == 0 {
			rootMenus = append(rootMenus, menu)
		} else {
			parent, exists := menuMap[menu.ParentID]
			if exists {
				parent.Children = append(parent.Children, menu)
				// Concatenate the path with parent's path
				menu.Path = parent.Path + "/" + menu.Path
			}
		}
	}

	return rootMenus
}

// Example usage
func main() {
	menus := []*SysMenu{
		{ID: 1, Name: "Home", Path: "", ParentID: 0},
		{ID: 2, Name: "Admin", Path: "admin", ParentID: 0},
		{ID: 3, Name: "Users", Path: "users", ParentID: 2},
		{ID: 4, Name: "Settings", Path: "settings", ParentID: 2},
		{ID: 5, Name: "Profile", Path: "profile", ParentID: 3},
		{ID: 6, Name: "Preferences", Path: "preferences", ParentID: 4},
	}

	tree := BuildMenuTree(menus)
	printMenuTree(tree, "")
}

// printMenuTree prints the menu tree in a readable format.
func printMenuTree(menus []*SysMenu, prefix string) {
	for _, menu := range menus {
		fmt.Printf("%s%s (Path: %s)\n", prefix, menu.Name, menu.Path)
		printMenuTree(menu.Children, prefix+"  ")
	}
}
