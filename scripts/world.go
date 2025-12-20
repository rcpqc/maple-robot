package scripts

import (
	"context"
	"fmt"
	"time"

	"maple-robot/ix"
	"maple-robot/log"
)

// Enter 进入角色
func Enter(ctx context.Context, index int) {

	page, pos := index/100, index%100

	log.Info(ctx, "角色选择")
	// 等待场景
	LabelWait(ctx, "角色选择-服务器", 30*time.Second)

	// 切换角色页
	for !LabelCheck(ctx, fmt.Sprintf("角色选择-选中角色页%d", page)) {
		LabelClick(ctx, "角色选择-下一页")
	}

	// 点击角色
	LabelClick(ctx, fmt.Sprintf("角色选择-%d号角色", pos))

	// 游戏开始
	LabelClick(ctx, "角色选择-游戏开始")

}

func WaitEnter(ctx context.Context) {
	// 等待进入
	LabelWait(ctx, "世界-电量", 30*time.Second)
	LabelClick(ctx, "世界-技能8")
}

func Exit(ctx context.Context) {
	for !LabelCheck(ctx, "通用-游戏结束-否", "通用-游戏结束-返回角色界面", "通用-游戏结束-是") {
		Back(ctx)
	}
	LabelClick(ctx, "通用-游戏结束-返回角色界面")
}

// 返回世界
func BackWorld(ctx context.Context) {
	log.Info(ctx, "返回世界")
	for retry := 0; !LabelCheck(ctx, "世界-导航", "世界-电量"); retry++ {
		Back(ctx)
		if retry >= 10 {
			log.Error(ctx, "返回世界", "retry", retry)
		}
	}
	time.Sleep(2 * time.Second)
	if !LabelCheck(ctx, "世界-导航", "世界-电量") {
		BackWorld(ctx)
	}
}

// 返回
func Back(ctx context.Context) {
	ix.Key(ix.KeyCodeEscape)
	time.Sleep(1500 * time.Millisecond)
}

func NextRole(ctx context.Context) {
	LabelWait(ctx, "世界-电量", 5*time.Second)
	LabelWaitClick(ctx, "世界-导航", 5*time.Second)
	LabelWait(ctx, "导航-关闭", 5*time.Second)
	LabelClick(ctx, "导航-更改角色")
	LabelWait(ctx, "更改角色-选择角色", 5*time.Second)
	if LabelCheck(ctx, "更改角色-上左") {
		LabelClick(ctx, "更改角色-上右")
		LabelWaitClick(ctx, "更改角色-变更", 5*time.Second)
	} else if LabelCheck(ctx, "更改角色-上右") {
		LabelClick(ctx, "更改角色-中左")
		LabelWaitClick(ctx, "更改角色-变更", 5*time.Second)
	} else if LabelCheck(ctx, "更改角色-中左") {
		LabelClick(ctx, "更改角色-中右")
		LabelWaitClick(ctx, "更改角色-变更", 5*time.Second)
	} else if LabelCheck(ctx, "更改角色-中右") {
		LabelClick(ctx, "更改角色-下左")
		LabelWaitClick(ctx, "更改角色-变更", 5*time.Second)
	} else if LabelCheck(ctx, "更改角色-下左") {
		LabelClick(ctx, "更改角色-下右")
		LabelWaitClick(ctx, "更改角色-变更", 5*time.Second)
	}
}

// 背包利用比率
func BackpackUtilizationRatio() float64 {
	empty := ix.Color{R: 161, G: 166, B: 173}
	x := 893
	for ; x >= 859; x-- {
		if !ix.GetPixel(ix.Position{X: int64(x), Y: 45}).Equals(empty) {
			break
		}
	}
	return float64(x-858+1) / 37.0
}

// 经验比率
func ExpRatio() float64 {
	full := ix.Color{R: 255, G: 168, B: 15}
	x := 0
	for ; x < 960; x++ {
		if !ix.GetPixel(ix.Position{X: int64(x), Y: 538}).Equals(full) {
			break
		}
	}
	return float64(x) / 960
}
