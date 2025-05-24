package scripts

import (
	"fmt"
	"maple-robot/context"
	"maple-robot/ix"
	"strconv"
	"time"
)

func init() {
	context.ProvideTask("天空岛贸易", tkdmy)
	context.ProvideTask("材料副本", clfb)
	context.ProvideTask("精英副本", jyfb)
	context.ProvideTask("周常副本", zcfb)
	context.ProvideTask("特殊周常副本", zcfb)
	context.ProvideTask("奈特的金字塔", ntdjzt)
	context.ProvideTask("武陵道场", wldc)
	context.ProvideTask("金钩海兵王", jghbw)
	context.ProvideTask("怪物乐园", gwly)
	context.ProvideTask("自动战斗时间", zdzdsj)
	context.ProvideTask("公会签到", ghqd)
	context.ProvideTask("领取个人奖励", lqgrjl)
	context.ProvideTask("送人气", srq)
	context.ProvideTask("领取日常奖励", lqrcjl)
	context.ProvideTask("公会聊天", ghlt)
	context.ProvideTask("领取公会奖励", lqghjl)
	context.ProvideTask("怪物乐园跳关", gwlytg)
	context.ProvideTask("委托佣兵", wtyb)
}

// tkdmy 天空岛贸易
func tkdmy(ctx *context.Context) {
	LabelWait("世界-电量", 5*time.Second)
	LabelWaitClick("世界-贸易", 5*time.Second)
	LabelWaitClick("天空岛贸易-前往获取战利品", 5*time.Second)
	LabelWaitClick("天空岛贸易-入场", 5*time.Second)
	ctx.Schedule()
	LabelWait("太初森林-今日", 5*time.Second)
	LabelClick("太初森林-精灵")
	LabelWaitClick("副本-退出", 5*time.Second)
	LabelWaitClick("太初森林-退出确认", 5*time.Second)
	LabelWaitClick("太初森林-结果确认", 5*time.Second)
	BackWorld()
}

// wldc 武陵道场
func wldc(ctx *context.Context) {
	LabelWait("世界-电量", 5*time.Second)
	LabelWaitClick("世界-日常", 5*time.Second)
	LabelWait("日常-进度", 5*time.Second)
	LabelClick("日常-简化模式5号")
	time.Sleep(3 * time.Second)
	LabelClick("武陵道场-挑战武陵道场")
	LabelWaitClick("武陵道场-入场", 5*time.Second)
	LabelWaitClick("武陵道场-进入", 5*time.Second)
	ctx.Schedule()
	LabelWaitClick("副本-退出", 15*time.Second)
	time.Sleep(6 * time.Second)
	LabelWaitClick("武陵道场-退出", 5*time.Second)
	LabelWaitClick("武陵道场-离开", 10*time.Second)
	BackWorld()
}

// zcfb 周常副本
func zcfb(ctx *context.Context) {
	LabelWait("世界-电量", 5*time.Second)
	LabelWaitClick("世界-日常", 5*time.Second)
	LabelWait("日常-进度", 5*time.Second)
	LabelClick("日常-简化模式3号")
	if mode := ctx.GetOption("模式"); mode == "特殊" {
		LabelClick("周常副本-特殊")
	}
	LabelWait("周常副本-入场", 5*time.Second)
	if time.Now().Weekday() == time.Saturday || time.Now().Weekday() == time.Sunday {
		LabelClick("周常副本-星期五")
	}
	LabelClick("周常副本-入场")
	LabelWaitClick("周常副本-入场-确定", 5*time.Second)
	LabelWait("副本-退出", 15*time.Second)
	ctx.Schedule()
	LabelWaitClick("周常副本-副本结算-退出", 90*time.Second)
	BackWorld()
}

// jyfb 精英副本
func jyfb(ctx *context.Context) {
	LabelWait("世界-电量", 5*time.Second)
	LabelWaitClick("世界-日常", 5*time.Second)
	LabelWait("日常-进度", 5*time.Second)
	LabelClick("日常-简化模式2号")
	LabelWait("精英副本-普通", 5*time.Second)
	if mode := ctx.GetOption("模式"); mode == "单人" {
		ix.Swipe(ix.Position{X: 100, Y: 120}, ix.Position{X: 100, Y: 520}, 1500)
		time.Sleep(time.Second)
		LabelClick("精英副本-鬼怪蘑菇王")
		LabelWaitClick("精英副本-创建房间", 5*time.Second)
		LabelWaitClick("精英副本-入场-确定", 5*time.Second)
		LabelWaitClick("精英副本-集结地-开始", 5*time.Second)
		ctx.Schedule()
		LabelWaitClick("精英副本-副本结算-单人离开", 90*time.Second)
	} else {
		LabelWaitClick("精英副本-快速组队", 5*time.Second)
		LabelWaitClick("精英副本-入场-确定", 5*time.Second)
		LabelWait("副本-退出", 60*time.Second)
		LabelWait("副本-麦克风", 15*time.Second)
		ctx.Schedule()
		LabelWaitClick("精英副本-副本结算-离开", 180*time.Second)
	}
	BackWorld()
}

// clfb 材料副本
func clfb(ctx *context.Context) {
	LabelWait("世界-电量", 5*time.Second)
	LabelWaitClick("世界-日常", 5*time.Second)
	LabelWait("日常-进度", 5*time.Second)
	LabelClick("日常-简化模式1号")
	LabelWait("材料副本-入场", 5*time.Second)
	if fbName := ctx.GetOption("副本名"); fbName != "" {
		LabelClick("材料副本-" + fbName)
	}
	LabelClick("材料副本-入场")
	LabelWaitClick("材料副本-入场-确定", 5*time.Second)
	if flag := ctx.GetOption("全部解除"); flag != "" {
		LabelWaitClick("材料副本-进入副本-全部解除", 5*time.Second)
	}
	LabelWaitClick("材料副本-进入副本-确定", 5*time.Second)
	LabelWait("副本-退出", 15*time.Second)
	ctx.Schedule()
	LabelWaitClick("材料副本-副本结算-退出", 300*time.Second)
	BackWorld()
}

// ntdjzt 奈特的金字塔
func ntdjzt(ctx *context.Context) {
	LabelWait("世界-电量", 5*time.Second)
	LabelWaitClick("世界-日常", 5*time.Second)
	LabelWait("日常-进度", 5*time.Second)
	LabelClick("日常-简化模式4号")
	LabelWaitClick("奈特的金字塔-快速组队", 5*time.Second)
	LabelWaitClick("奈特的金字塔-快速组队-入场确定", 5*time.Second)
	LabelWait("副本-退出", 60*time.Second)
	LabelWait("副本-麦克风", 30*time.Second)
	ctx.Schedule()
	LabelWaitClick("奈特的金字塔-副本结算-退出", 180*time.Second)
	BackWorld()
}

// jghbw 金钩海兵王
func jghbw(ctx *context.Context) {
	LabelWait("世界-电量", 5*time.Second)
	LabelWaitClick("世界-日常", 5*time.Second)
	LabelWait("日常-进度", 5*time.Second)
	LabelClick("日常-简化模式6号")
	LabelWaitClick("金钩海兵王-快速组队", 5*time.Second)
	LabelWaitClick("金钩海兵王-快速组队-入场确认", 5*time.Second)
	LabelWait("副本-退出", 60*time.Second)
	ctx.Schedule()
	LabelWaitClick("金钩海兵王-副本结算-退出", 120*time.Second)
	BackWorld()
}

// gwly 怪物乐园
func gwly(ctx *context.Context) {
	LabelWait("世界-电量", 5*time.Second)
	LabelWaitClick("世界-日常", 5*time.Second)
	LabelWait("日常-进度", 5*time.Second)
	LabelClick("日常-简化模式7号")
	buff := ctx.GetOption("经验增益")
	LabelWaitClick("怪物乐园-入场", 5*time.Second)
	LabelWaitClick("怪物乐园-入场-票券确认", 5*time.Second)
	if buff != "" {
		LabelWaitClick("怪物乐园-入场-经验增益", 5*time.Second)
	}
	LabelWaitClick("怪物乐园-入场-入场", 5*time.Second)
	LabelWait("副本-退出", 10*time.Second)
	ctx.Schedule()
	extraBonus := ctx.GetOption("追加奖励")
	if extraBonus != "" {
		LabelWaitClick("怪物乐园-副本结算-追加奖励", 360*time.Second)
		LabelWaitClick("怪物乐园-副本结算-追加奖励-确定", 5*time.Second)
		time.Sleep(2 * time.Second)
		LabelWaitClick("怪物乐园-副本结算-退出", 5*time.Second)
	} else {
		LabelWaitClick("怪物乐园-副本结算-退出", 360*time.Second)
	}
	BackWorld()
}

// zdzdsj 自动战斗时间
func zdzdsj(ctx *context.Context) {
	LabelClick("世界-自动战斗")
	LabelClick("自动战斗-使用")
	ctx.Schedule()
	// LabelClick("自动战斗-关闭")
	BackWorld()
}

// ghqd 公会签到
func ghqd(ctx *context.Context) {
	LabelWait("世界-电量", 5*time.Second)
	LabelWaitClick("世界-公会", 5*time.Second)
	LabelWait("公会-领地树", 5*time.Second)

	if flag := ctx.GetOption("领地树祝福"); flag != "NO" {
		LabelClick("公会-领地树")
		LabelWait("公会-领地树-祝福领取取消", 5*time.Second)
		LabelClick("公会-领地树-祝福领取确定")
		LabelClick("公会-领地树-祝福领取取消")
	}

	// LabelClick("公会-领取")
	ctx.Schedule()
	BackWorld()
}

// lqgrjl 领取个人奖励
func lqgrjl(ctx *context.Context) {
	LabelWait("世界-电量", 5*time.Second)
	LabelWaitClick("世界-邮箱", 5*time.Second)
	LabelWait("邮箱-通用", 5*time.Second)
	LabelClick("邮箱-个人")
	LabelClick("邮箱-全部领取")
	LabelClick("邮箱-全部领取-确定")
	ctx.Schedule()
	BackWorld()
}

// srq 送人气
func srq(ctx *context.Context) {
	LabelWait("世界-电量", 5*time.Second)
	LabelWaitClick("世界-社交", 5*time.Second)
	LabelWaitClick("社交-微信好友", 5*time.Second)
	LabelWait("社交-微信好友-背景", 5*time.Second)
	LabelClick("社交-微信好友-战斗力")
	LabelWaitClick("社交-微信好友-送人气给3号", 5*time.Second)
	LabelClick("社交-送人气")
	LabelClick("社交-送人气确定")
	LabelClick("社交-送人气不提醒")
	ctx.Schedule()
	// LabelClick("社交-送人气关闭")
	BackWorld()
}

// lqrcjl 领取日常奖励
func lqrcjl(ctx *context.Context) {
	LabelWait("世界-电量", 5*time.Second)
	for LabelCheck("世界-领取日常") {
		LabelClick("世界-领取日常")
		BackWorld()
	}
	ctx.Schedule()
	BackWorld()
}

// ghlt 公会聊天
func ghlt(ctx *context.Context) {
	LabelClick("世界-聊天栏")
	LabelWaitClick("聊天栏-公会", 5*time.Second)
	LabelWaitClick("聊天栏-表情", 5*time.Second)
	LabelWaitClick("聊天栏-表情-害羞", 5*time.Second)
	LabelWaitClick("聊天栏-发送", 5*time.Second)
	LabelClick("聊天栏-发送")
	ctx.Schedule()
	BackWorld()
}

// lqghjl 领取公会奖励
func lqghjl(ctx *context.Context) {
	LabelWait("世界-电量", 5*time.Second)
	LabelWaitClick("世界-公会", 5*time.Second)
	LabelClick("公会-公会任务")
	LabelClick("公会-公会任务-每日任务")
	LabelClick("公会-公会任务-全部领取")
	if time.Now().Weekday() == time.Saturday || time.Now().Weekday() == time.Sunday {
		LabelClick("公会-公会任务-每周任务")
		LabelClick("公会-公会任务-全部领取")
	}
	ctx.Schedule()
	BackWorld()
}

// gwlytg 怪物乐园跳关
func gwlytg(ctx *context.Context) {
	LabelWait("世界-电量", 5*time.Second)
	LabelWaitClick("世界-日常", 5*time.Second)
	LabelWaitClick("日常-进度", 5*time.Second)
	LabelWait("日常-进度-关闭", 5*time.Second)
	ix.Swipe(ix.Position{X: 579, Y: 481}, ix.Position{X: 579, Y: 50}, 1500)
	time.Sleep(time.Second)
	LabelWaitClick("日常-进度-怪物乐园跳关", 5*time.Second)
	LabelWait("日常-进度-怪物乐园跳关-标题", 5*time.Second)
	// LabelWaitClick("日常-进度-怪物乐园跳关-隐匿痕迹", 5*time.Second)
	LabelWaitClick("日常-进度-怪物乐园跳关-使用战斗跳关券", 5*time.Second)
	LabelWaitClick("日常-进度-怪物乐园跳关-入场确认", 5*time.Second)
	ctx.Schedule()
	LabelWaitClick("日常-进度-怪物乐园跳关-结算确认", 5*time.Second)
	BackWorld()
}

// 委托佣兵
func wtyb(ctx *context.Context) {
	LabelWait("世界-电量", 5*time.Second)
	LabelWaitClick("世界-委托", 5*time.Second)
	LabelWait("委托-切换", 5*time.Second)
	maxChange, _ := strconv.ParseInt(ctx.GetOption("最大切换"), 10, 64)
	maxChange = min(max(maxChange, 0), 5)
	maxAccept, _ := strconv.ParseInt(ctx.GetOption("最大接受"), 10, 64)
	maxAccept = min(max(maxAccept, 1), 3)

	// 接受委托
	change, pos := 0, 1
	for LabelColor(fmt.Sprintf("委托-接受%d号位", maxAccept)) == ix.ColorMissionEmpty {
		if LabelColor(fmt.Sprintf("委托-发布%d号位", pos)) == ix.ColorMissionHard {
			LabelClick(fmt.Sprintf("委托-发布%d号位", pos))
			LabelWaitClick("委托-发布接受", 5*time.Second)
		}
		pos++
		// 判定完成一页
		if pos > 5 {
			if change < int(maxChange) {
				LabelClick("委托-切换")
				change++
				pos = 1
			} else {
				break
			}
		}
	}
	// 委托雇佣
	for accept := 1; accept <= int(maxAccept); accept++ {
		mission := fmt.Sprintf("委托-接受%d号位", accept)
		if LabelColor(mission) == ix.ColorMissionExist {
			LabelClick(mission)
			if LabelCheck("委托-接受佣兵团") {
				LabelClick("委托-接受佣兵团")
				LabelWaitClick("委托-领取奖励", 5*time.Second)
				LabelClick("委托-接受1号位")
			} else {
				break
			}
		}
	}

	ctx.Schedule()
	BackWorld()
}
