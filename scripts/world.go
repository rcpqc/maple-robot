package scripts

import (
	"maple-robot/ix"
	"maple-robot/log"
	"time"
)

func WorldReturnRole() {
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
