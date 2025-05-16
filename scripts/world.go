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

func WaitEnter() {
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
	for {
		for !LabelCheck("世界-导航", "世界-电量") {
			Back()
		}
		time.Sleep(2 * time.Second)
		if LabelCheck("世界-导航", "世界-电量") {
			break
		}
	}
}

func Back() {
	ix.Key(ix.KeyCodeEscape)
	time.Sleep(1500 * time.Millisecond)
}

func NextRole() bool {
	LabelWait("世界-电量", 5*time.Second)
	LabelWaitClick("世界-角色", 5*time.Second)
	LabelWait("更改角色-选择角色", 5*time.Second)
	if LabelCheck("更改角色-上左") {
		LabelClick("更改角色-上右")
		LabelWaitClick("更改角色-变更", 5*time.Second)
		return true
	} else if LabelCheck("更改角色-上右") {
		LabelClick("更改角色-中左")
		LabelWaitClick("更改角色-变更", 5*time.Second)
		return true
	} else if LabelCheck("更改角色-中左") {
		LabelClick("更改角色-中右")
		LabelWaitClick("更改角色-变更", 5*time.Second)
		return true
	} else if LabelCheck("更改角色-中右") {
		LabelClick("更改角色-下左")
		LabelWaitClick("更改角色-变更", 5*time.Second)
		return true
	} else if LabelCheck("更改角色-下左") {
		LabelClick("更改角色-下右")
		LabelWaitClick("更改角色-变更", 5*time.Second)
		return true
	} else if LabelCheck("更改角色-下右") {
		LabelClick("更改角色-选择角色")
		return false
	}
	return false
}
