package scripts

import (
	"fmt"
	"maple-robot/ix"
	"maple-robot/log"
	"time"
)

// Enter 进入角色
func Enter(index int) {

	page, pos := index/100, index%100

	log.Infof("选择角色-%d\n", index)
	// 等待场景
	LabelWait("角色选择-服务器", 30*time.Second)

	// 切换角色页
	for !LabelCheck(fmt.Sprintf("角色选择-选中角色页%d", page)) {
		LabelClick("角色选择-下一页")
	}

	// 点击角色
	LabelClick(fmt.Sprintf("角色选择-%d号角色", pos))

	// 游戏开始
	LabelClick("角色选择-游戏开始")

	// 等待进入
	LabelWait("世界-电量", 30*time.Second)

	LabelClick("世界-技能8")
}

func Exit() {
	log.Infof("返回角色界面\n")
	for !LabelCheck("通用-游戏结束-否", "通用-游戏结束-返回角色界面", "通用-游戏结束-是") {
		Back()
	}
	LabelClick("通用-游戏结束-返回角色界面")
}

func BackWorld() {
	log.Infof("返回世界界面\n")
	for !LabelCheck("世界-导航", "世界-电量") {
		Back()
	}
}

func Back() {
	ix.Key(ix.KeyCodeEscape)
	time.Sleep(1500 * time.Millisecond)
}
