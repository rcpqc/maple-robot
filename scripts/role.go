package scripts

import (
	"fmt"
	"maple-robot/log"
	"time"
)

// RoleSelect 选择角色
func RoleSelect(page, index int) error {
	log.Infof("选择角色 %d-%d\n", page, index)
	// 等待场景
	LabelWait("角色选择-服务器", 15*time.Second)

	// 切换角色页
	for !LabelCheck(fmt.Sprintf("角色选择-选中角色页%d", page)) {
		LabelClick("角色选择-下一页")
	}

	// 点击角色
	LabelClick(fmt.Sprintf("角色选择-%d号角色", index))

	// 游戏开始
	LabelClick("角色选择-游戏开始")

	// 等待进入
	LabelWait("世界-电量", 15*time.Second)

	return nil
}
