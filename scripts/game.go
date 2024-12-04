package scripts

import (
	"maple-robot/context"
	"time"
)

func init() {
	context.ProvideScript("天空岛贸易", tkdmy)
	context.ProvideScript("武陵道场", wldc)
	context.ProvideScript("周常副本", zcfb)
	context.ProvideScript("精英副本", jyfb)
	context.ProvideScript("材料副本", clfb)
	context.ProvideScript("奈特的金字塔", ntdjzt)
	context.ProvideScript("自动战斗时间", zdzdsj)
	context.ProvideScript("公会签到", ghqd)
	context.ProvideScript("领取个人奖励", lqgrjl)
	context.ProvideScript("送人气", srq)
	context.ProvideScript("领取日常奖励", lqrcjl)
	context.ProvideScript("补充训练时间", bcxlsj)
	context.ProvideScript("公会聊天", ghlt)
	context.ProvideScript("领取公会奖励", lqghjl)
}

// tkdmy 天空岛贸易
func tkdmy(ctx *context.Context) {
	LabelWaitClick("世界-导航", 5*time.Second)
	LabelWait("导航-枫小喵", 5*time.Second)
	LabelClick("导航-天空岛贸易")
	LabelWaitClick("天空岛贸易-前往获取战利品", 5*time.Second)
	LabelWaitClick("天空岛贸易-入场", 5*time.Second)
	ctx.Schedule()
	LabelWait("太初森林-今日", 5*time.Second)
	LabelClick("太初森林-精灵")
	LabelWaitClick("副本-退出", 5*time.Second)
	LabelWaitClick("太初森林-退出确认", 5*time.Second)
	LabelWaitClick("太初森林-结果确认", 5*time.Second)
}

// wldc 武陵道场
func wldc(ctx *context.Context) {
	LabelWaitClick("世界-导航", 5*time.Second)
	LabelWait("导航-枫小喵", 5*time.Second)
	LabelClick("导航-日常")
	LabelWait("日常-进度", 5*time.Second)
	LabelClick("日常-简化模式5号")
	LabelWaitClick("武陵道场-入场", 5*time.Second)
	LabelWaitClick("武陵道场-进入", 5*time.Second)
	ctx.Schedule()
	LabelWaitClick("副本-退出", 15*time.Second)
	time.Sleep(8 * time.Second)
	LabelWaitClick("武陵道场-退出", 5*time.Second)
	LabelWaitClick("武陵道场-离开", 10*time.Second)
}

// zcfb 周常副本
func zcfb(ctx *context.Context) {
	LabelWaitClick("世界-导航", 5*time.Second)
	LabelWait("导航-枫小喵", 5*time.Second)
	LabelClick("导航-日常")
	LabelWait("日常-进度", 5*time.Second)
	LabelClick("日常-简化模式3号")
	LabelWait("周常副本-入场", 5*time.Second)
	if time.Now().Weekday() == time.Saturday || time.Now().Weekday() == time.Sunday {
		LabelClick("周常副本-星期五")
	}
	LabelClick("周常副本-入场")
	LabelWaitClick("周常副本-入场-确定", 5*time.Second)
	ctx.Schedule()
	LabelWait("副本-退出", 15*time.Second)
	LabelWaitClick("周常副本-副本结算-退出", 90*time.Second)
}

// jyfb 精英副本
func jyfb(ctx *context.Context) {
	LabelWaitClick("世界-导航", 5*time.Second)
	LabelWait("导航-枫小喵", 5*time.Second)
	LabelClick("导航-日常")
	LabelWait("日常-进度", 5*time.Second)
	LabelClick("日常-简化模式2号")
	LabelWaitClick("精英副本-快速组队", 5*time.Second)
	LabelWaitClick("精英副本-入场-确定", 5*time.Second)
	LabelWait("副本-退出", 60*time.Second)
	LabelWait("副本-麦克风", 15*time.Second)
	ctx.Schedule()
	LabelWaitClick("精英副本-副本结算-离开", 180*time.Second)
}

// clfb 材料副本
func clfb(ctx *context.Context) {
	LabelWaitClick("世界-导航", 5*time.Second)
	LabelWait("导航-枫小喵", 5*time.Second)
	LabelClick("导航-日常")
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
}

// ntdjzt 奈特的金字塔
func ntdjzt(ctx *context.Context) {
	LabelWaitClick("世界-导航", 5*time.Second)
	LabelWait("导航-枫小喵", 5*time.Second)
	LabelClick("导航-日常")
	LabelWait("日常-进度", 5*time.Second)
	LabelClick("日常-简化模式4号")
	LabelWaitClick("奈特的金字塔-快速组队", 5*time.Second)
	LabelWaitClick("奈特的金字塔-快速组队-入场确定", 5*time.Second)
	LabelWait("副本-退出", 60*time.Second)
	LabelWait("副本-麦克风", 30*time.Second)
	ctx.Schedule()
	LabelWaitClick("奈特的金字塔-副本结算-退出", 180*time.Second)
}

// zdzdsj 自动战斗时间
func zdzdsj(ctx *context.Context) {
	LabelClick("世界-自动战斗")
	LabelClick("自动战斗-使用")
	ctx.Schedule()
	LabelClick("自动战斗-关闭")
}

// ghqd 公会签到
func ghqd(ctx *context.Context) {
	LabelWaitClick("世界-导航", 5*time.Second)
	LabelWait("导航-枫小喵", 5*time.Second)
	LabelClick("导航-公会")
	LabelWait("公会-领地树", 5*time.Second)

	if flag := ctx.GetOption("领地树祝福"); flag != "NO" {
		LabelClick("公会-领地树")
		LabelWait("公会-领地树-祝福领取取消", 5*time.Second)
		LabelClick("公会-领地树-祝福领取确定")
		LabelClick("公会-领地树-祝福领取取消")
	}

	LabelClick("公会-领取")
	ctx.Schedule()
	BackWorld()
}

// lqgrjl 领取个人奖励
func lqgrjl(ctx *context.Context) {
	LabelWaitClick("世界-导航", 5*time.Second)
	LabelWait("导航-枫小喵", 5*time.Second)
	LabelClick("导航-邮箱")
	LabelWait("邮箱-通用", 5*time.Second)
	LabelClick("邮箱-个人")
	LabelClick("邮箱-全部领取")
	LabelClick("邮箱-全部领取-确定")
	ctx.Schedule()
	BackWorld()
}

// srq 送人气
func srq(ctx *context.Context) {
	LabelWaitClick("世界-导航", 5*time.Second)
	LabelWait("导航-枫小喵", 5*time.Second)
	LabelClick("导航-社交")
	LabelWaitClick("社交-微信好友", 5*time.Second)
	LabelWait("社交-微信好友-背景", 5*time.Second)
	LabelClick("社交-微信好友-战斗力")
	LabelWaitClick("社交-微信好友-送人气给3号", 5*time.Second)
	LabelClick("社交-送人气")
	LabelClick("社交-送人气确定")
	LabelClick("社交-送人气不提醒")
	ctx.Schedule()
	LabelClick("社交-送人气关闭")
	BackWorld()
}

// lqrcjl 领取日常奖励
func lqrcjl(ctx *context.Context) {
	LabelWaitClick("世界-导航", 5*time.Second)
	LabelWait("导航-枫小喵", 5*time.Second)
	LabelClick("导航-成长")
	LabelClick("成长-每日任务")
	LabelClick("成长-每日任务-全部领取")
	LabelClick("成长-领取确认")
	LabelClick("成长-每日任务-全部领取")
	LabelClick("成长-领取确认")
	LabelClick("成长-每周任务")
	LabelClick("成长-每周任务-全部领取")
	LabelClick("成长-领取确认")
	LabelClick("成长-每周任务-全部领取")
	LabelClick("成长-领取确认")
	ctx.Schedule()
	BackWorld()
}

// bcxlsj 补充训练时间
func bcxlsj(ctx *context.Context) {
	LabelWaitClick("世界-导航", 5*time.Second)
	LabelWait("导航-枫小喵", 5*time.Second)
	LabelClick("导航-角色训练场")
	LabelWaitClick("角色训练场-补充时间", 5*time.Second)
	LabelClick("角色训练场-补充时间-使用")
	LabelClick("角色训练场-补充时间-使用")
	ctx.Schedule()
	LabelWaitClick("角色训练场-补充时间-确定", 5*time.Second)
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
	LabelWaitClick("世界-导航", 5*time.Second)
	LabelWait("导航-枫小喵", 5*time.Second)
	LabelClick("导航-公会")
	LabelClick("公会-公会任务")
	LabelClick("公会-公会任务-每日任务")
	LabelClick("公会-公会任务-全部领取")
	LabelClick("公会-公会任务-每周任务")
	LabelClick("公会-公会任务-全部领取")
	ctx.Schedule()
	BackWorld()
}
