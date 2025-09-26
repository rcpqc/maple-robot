package scripts

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"maple-robot/config"
	"maple-robot/ix"
	"maple-robot/log"
	"strconv"
	"time"
)

func init() {
	config.ProvideTask("天空岛贸易", tkdmy)
	config.ProvideTask("材料副本", clfb)
	config.ProvideTask("精英副本", jyfb)
	config.ProvideTask("周常副本", zcfb)
	config.ProvideTask("特殊周常副本", zcfb)
	config.ProvideTask("奈特的金字塔", ntdjzt)
	config.ProvideTask("武陵道场", wldc)
	config.ProvideTask("金钩海兵王", jghbw)
	config.ProvideTask("怪物乐园", gwly)
	config.ProvideTask("自动战斗时间", zdzdsj)
	config.ProvideTask("公会签到", ghqd)
	config.ProvideTask("领取个人奖励", lqgrjl)
	config.ProvideTask("送人气", srq)
	config.ProvideTask("领取日常奖励", lqrcjl)
	config.ProvideTask("公会聊天", ghlt)
	config.ProvideTask("领取公会奖励", lqghjl)
	config.ProvideTask("委托佣兵", wtyb)
	config.ProvideTask("购买委托书", gmwts)
}

// tkdmy 天空岛贸易
func tkdmy(ctx context.Context) {
	LabelWait(ctx, "世界-电量", 5*time.Second)
	LabelWaitClick(ctx, "世界-贸易", 5*time.Second)
	LabelWaitClick(ctx, "天空岛贸易-前往获取战利品", 5*time.Second)
	LabelWaitClick(ctx, "天空岛贸易-入场", 5*time.Second)
	log.Info(ctx, "任务入场")
	LabelWait(ctx, "太初森林-今日", 5*time.Second)
	LabelClick(ctx, "太初森林-精灵")
	LabelWaitClick(ctx, "副本-退出", 5*time.Second)
	LabelWaitClick(ctx, "太初森林-退出确认", 5*time.Second)
	LabelWaitClick(ctx, "太初森林-结果确认", 5*time.Second)
	BackWorld(ctx)
}

// wldc 武陵道场
func wldc(ctx context.Context) {
	LabelWait(ctx, "世界-电量", 5*time.Second)
	LabelWaitClick(ctx, "世界-日常", 5*time.Second)
	LabelWait(ctx, "日常-进度", 5*time.Second)
	LabelClick(ctx, "日常-简化模式5号")
	time.Sleep(3 * time.Second)
	LabelClick(ctx, "武陵道场-挑战武陵道场")
	LabelWaitClick(ctx, "武陵道场-入场", 5*time.Second)
	LabelWaitClick(ctx, "武陵道场-进入", 5*time.Second)
	log.Info(ctx, "任务入场")
	LabelWaitClick(ctx, "副本-退出", 15*time.Second)
	time.Sleep(6 * time.Second)
	LabelWaitClick(ctx, "武陵道场-退出", 5*time.Second)
	LabelWaitClick(ctx, "武陵道场-离开", 10*time.Second)
	BackWorld(ctx)
}

// zcfb 周常副本
func zcfb(ctx context.Context) {
	LabelWait(ctx, "世界-电量", 5*time.Second)
	LabelWaitClick(ctx, "世界-日常", 5*time.Second)
	LabelWait(ctx, "日常-进度", 5*time.Second)
	LabelClick(ctx, "日常-简化模式3号")

	if mode := config.GetTaskOptions(ctx, "模式"); mode == "特殊" {
		LabelClick(ctx, "周常副本-特殊")
	}
	if difficulty := config.GetTaskOptions(ctx, "难度"); difficulty != "" {
		LabelClick(ctx, "周常副本-"+difficulty)
	}
	LabelWait(ctx, "周常副本-入场", 5*time.Second)
	if time.Now().Weekday() == time.Saturday || time.Now().Weekday() == time.Sunday {
		LabelClick(ctx, "周常副本-星期五")
	}
	LabelClick(ctx, "周常副本-入场")
	LabelWaitClick(ctx, "周常副本-入场-确定", 5*time.Second)
	LabelWait(ctx, "副本-退出", 15*time.Second)
	log.Info(ctx, "任务入场")
	LabelWaitClick(ctx, "周常副本-副本结算-退出", 180*time.Second)
	BackWorld(ctx)
}

// jyfb 精英副本
func jyfb(ctx context.Context) {
	LabelWait(ctx, "世界-电量", 5*time.Second)
	LabelWaitClick(ctx, "世界-日常", 5*time.Second)
	LabelWait(ctx, "日常-进度", 5*time.Second)
	LabelClick(ctx, "日常-简化模式2号")
	LabelWait(ctx, "精英副本-普通", 5*time.Second)
	if mode := config.GetTaskOptions(ctx, "模式"); mode == "单人" {
		ix.Swipe(ix.Position{X: 100, Y: 120}, ix.Position{X: 100, Y: 520}, 1500)
		time.Sleep(time.Second)
		LabelClick(ctx, "精英副本-鬼怪蘑菇王")
		LabelWaitClick(ctx, "精英副本-创建队伍", 5*time.Second)
		LabelWaitClick(ctx, "精英副本-入场-确定", 5*time.Second)
		LabelWaitClick(ctx, "精英副本-集结地-开始", 5*time.Second)
		log.Info(ctx, "任务入场")
		LabelWaitClick(ctx, "精英副本-副本结算-单人离开", 90*time.Second)
	} else {
		LabelWaitClick(ctx, "精英副本-快速组队", 5*time.Second)
		LabelWaitClick(ctx, "精英副本-入场-确定", 5*time.Second)
		LabelWait(ctx, "副本-退出", 60*time.Second)
		LabelWait(ctx, "副本-麦克风", 15*time.Second)
		log.Info(ctx, "任务入场")
		LabelWaitClick(ctx, "精英副本-副本结算-离开", 180*time.Second)
	}
	BackWorld(ctx)
}

// clfb 材料副本
func clfb(ctx context.Context) {
	LabelWait(ctx, "世界-电量", 5*time.Second)
	LabelWaitClick(ctx, "世界-日常", 5*time.Second)
	LabelWait(ctx, "日常-进度", 5*time.Second)
	LabelClick(ctx, "日常-简化模式1号")
	LabelWait(ctx, "材料副本-入场", 5*time.Second)

	if fbName := config.GetTaskOptions(ctx, "副本名"); fbName != "" {
		LabelClick(ctx, "材料副本-"+fbName)
	}
	LabelClick(ctx, "材料副本-入场")
	LabelWaitClick(ctx, "材料副本-入场-确定", 5*time.Second)

	if flag := config.GetTaskOptions(ctx, "全部解除"); flag != "" {
		LabelWaitClick(ctx, "材料副本-进入副本-全部解除", 5*time.Second)
	}
	LabelWaitClick(ctx, "材料副本-进入副本-确定", 5*time.Second)
	LabelWait(ctx, "副本-退出", 15*time.Second)
	log.Info(ctx, "任务入场")
	LabelWaitClick(ctx, "材料副本-副本结算-退出", 300*time.Second)
	BackWorld(ctx)
}

// ntdjzt 奈特的金字塔
func ntdjzt(ctx context.Context) {
	LabelWait(ctx, "世界-电量", 5*time.Second)
	LabelWaitClick(ctx, "世界-日常", 5*time.Second)
	LabelWait(ctx, "日常-进度", 5*time.Second)
	LabelClick(ctx, "日常-简化模式4号")
	LabelWaitClick(ctx, "奈特的金字塔-快速组队", 5*time.Second)
	LabelWaitClick(ctx, "奈特的金字塔-快速组队-入场确定", 5*time.Second)
	LabelWait(ctx, "副本-退出", 60*time.Second)
	LabelWait(ctx, "副本-麦克风", 30*time.Second)
	log.Info(ctx, "任务入场")
	LabelWaitClick(ctx, "奈特的金字塔-副本结算-退出", 180*time.Second)
	BackWorld(ctx)
}

// jghbw 金钩海兵王
func jghbw(ctx context.Context) {
	LabelWait(ctx, "世界-电量", 5*time.Second)
	LabelWaitClick(ctx, "世界-日常", 5*time.Second)
	LabelWait(ctx, "日常-进度", 5*time.Second)
	LabelClick(ctx, "日常-简化模式6号")
	LabelWaitClick(ctx, "金钩海兵王-快速组队", 5*time.Second)
	LabelWaitClick(ctx, "金钩海兵王-快速组队-入场确认", 5*time.Second)
	LabelWait(ctx, "副本-退出", 60*time.Second)
	log.Info(ctx, "任务入场")
	LabelWaitClick(ctx, "金钩海兵王-副本结算-退出", 120*time.Second)
	BackWorld(ctx)
}

// gwly 怪物乐园
func gwly(ctx context.Context) {
	role := GetRole(ctx)
	LabelWait(ctx, "世界-电量", 5*time.Second)
	LabelWaitClick(ctx, "世界-日常", 5*time.Second)
	// excludeList := strings.Split(config.GetTaskOptions(ctx, "跳关排除列表"), ",")
	excludeList := []string{"黑骑士", "箭神", "幻影", "炎术士", "冰雷魔导师"}
	pos := config.GetTaskOptions(ctx, "标签位置")
	stage := config.GetTaskOptions(ctx, "关卡")
	if !slices.Contains(excludeList, role.Class) {
		LabelWaitClick(ctx, "日常-进度", 5*time.Second)
		LabelWait(ctx, "日常-进度-关闭", 5*time.Second)
		ix.Swipe(ix.Position{X: 579, Y: 481}, ix.Position{X: 579, Y: 50}, 1500)
		time.Sleep(time.Second)
		LabelWaitClick(ctx, "日常-进度-"+pos, 5*time.Second)
		LabelWait(ctx, "日常-进度-怪物乐园跳关-标题", 5*time.Second)
		if stage != "" {
			LabelClick(ctx, "日常-进度-怪物乐园跳关-"+stage)
		}
		LabelWaitClick(ctx, "日常-进度-怪物乐园跳关-使用战斗跳关券", 5*time.Second)
		LabelWaitClick(ctx, "日常-进度-怪物乐园跳关-入场确认", 5*time.Second)
		log.Info(ctx, "任务入场")
		LabelWaitClick(ctx, "日常-进度-怪物乐园跳关-结算确认", 5*time.Second)
	} else {
		LabelWait(ctx, "日常-进度", 5*time.Second)
		LabelClick(ctx, "日常-简化模式7号")
		buff := config.GetTaskOptions(ctx, "经验增益")
		LabelWaitClick(ctx, "怪物乐园-入场", 5*time.Second)
		LabelWaitClick(ctx, "怪物乐园-入场-票券确认", 5*time.Second)
		if buff != "" {
			LabelWaitClick(ctx, "怪物乐园-入场-经验增益", 5*time.Second)
		}
		LabelWaitClick(ctx, "怪物乐园-入场-入场", 5*time.Second)
		LabelWait(ctx, "副本-退出", 10*time.Second)
		log.Info(ctx, "任务入场")

		extraBonus := config.GetTaskOptions(ctx, "追加奖励")
		if extraBonus != "" {
			LabelWaitClick(ctx, "怪物乐园-副本结算-追加奖励", 360*time.Second)
			LabelWaitClick(ctx, "怪物乐园-副本结算-追加奖励-确定", 5*time.Second)
			time.Sleep(2 * time.Second)
			LabelWaitClick(ctx, "怪物乐园-副本结算-退出", 5*time.Second)
		} else {
			LabelWaitClick(ctx, "怪物乐园-副本结算-退出", 600*time.Second)
		}
	}
	BackWorld(ctx)
}

// zdzdsj 自动战斗时间
func zdzdsj(ctx context.Context) {
	LabelClick(ctx, "世界-自动战斗")
	LabelClick(ctx, "自动战斗-使用")
	log.Info(ctx, "任务入场")
	// LabelClick(ctx, "自动战斗-关闭")
	BackWorld(ctx)
}

// ghqd 公会签到
func ghqd(ctx context.Context) {
	LabelWait(ctx, "世界-电量", 5*time.Second)
	LabelWaitClick(ctx, "世界-公会", 5*time.Second)
	LabelWait(ctx, "公会-领地树", 5*time.Second)

	if flag := config.GetTaskOptions(ctx, "领地树祝福"); flag != "NO" {
		LabelClick(ctx, "公会-领地树")
		LabelWait(ctx, "公会-领地树-祝福领取取消", 5*time.Second)
		LabelClick(ctx, "公会-领地树-祝福领取确定")
		LabelClick(ctx, "公会-领地树-祝福领取取消")
	}

	// LabelClick(ctx, "公会-领取")
	log.Info(ctx, "任务入场")
	BackWorld(ctx)
}

// lqgrjl 领取个人奖励
func lqgrjl(ctx context.Context) {
	LabelWait(ctx, "世界-电量", 5*time.Second)
	LabelWaitClick(ctx, "世界-邮箱", 5*time.Second)
	LabelWait(ctx, "邮箱-通用", 5*time.Second)
	LabelClick(ctx, "邮箱-个人")
	LabelClick(ctx, "邮箱-全部领取")
	LabelClick(ctx, "邮箱-全部领取-确定")
	log.Info(ctx, "任务入场")
	BackWorld(ctx)
}

// srq 送人气
func srq(ctx context.Context) {
	LabelWait(ctx, "世界-电量", 5*time.Second)
	LabelWaitClick(ctx, "世界-社交", 5*time.Second)
	LabelWaitClick(ctx, "社交-微信好友", 5*time.Second)
	LabelWait(ctx, "社交-微信好友-背景", 5*time.Second)
	LabelClick(ctx, "社交-微信好友-战斗力")
	LabelWaitClick(ctx, "社交-微信好友-送人气给3号", 5*time.Second)
	LabelClick(ctx, "社交-送人气")
	LabelClick(ctx, "社交-送人气确定")
	LabelClick(ctx, "社交-送人气不提醒")
	log.Info(ctx, "任务入场")
	// LabelClick(ctx, "社交-送人气关闭")
	BackWorld(ctx)
}

// lqrcjl 领取日常奖励
func lqrcjl(ctx context.Context) {
	LabelWait(ctx, "世界-电量", 5*time.Second)
	for LabelCheck(ctx, "世界-领取日常") {
		LabelClick(ctx, "世界-领取日常")
		BackWorld(ctx)
	}
	log.Info(ctx, "任务入场")
	BackWorld(ctx)
}

// ghlt 公会聊天
func ghlt(ctx context.Context) {
	LabelClick(ctx, "世界-聊天栏")
	LabelWaitClick(ctx, "聊天栏-公会", 5*time.Second)
	LabelWaitClick(ctx, "聊天栏-表情", 5*time.Second)
	LabelWaitClick(ctx, "聊天栏-表情-害羞", 5*time.Second)
	LabelWaitClick(ctx, "聊天栏-发送", 5*time.Second)
	LabelClick(ctx, "聊天栏-发送")
	log.Info(ctx, "任务入场")
	BackWorld(ctx)
}

// lqghjl 领取公会奖励
func lqghjl(ctx context.Context) {
	LabelWait(ctx, "世界-电量", 5*time.Second)
	LabelWaitClick(ctx, "世界-公会", 5*time.Second)
	LabelClick(ctx, "公会-公会任务")
	LabelClick(ctx, "公会-公会任务-每日任务")
	LabelClick(ctx, "公会-公会任务-全部领取")
	if time.Now().Weekday() == time.Saturday || time.Now().Weekday() == time.Sunday {
		LabelClick(ctx, "公会-公会任务-每周任务")
		LabelClick(ctx, "公会-公会任务-全部领取")
	}
	log.Info(ctx, "任务入场")
	BackWorld(ctx)
}

// 委托佣兵
func wtyb(ctx context.Context) {
	LabelWait(ctx, "世界-电量", 5*time.Second)
	LabelWaitClick(ctx, "世界-委托", 5*time.Second)
	LabelWait(ctx, "委托-切换", 5*time.Second)
	maxChange, _ := strconv.ParseInt(config.GetTaskOptions(ctx, "最大切换"), 10, 64)
	maxChange = min(max(maxChange, 0), 5)
	maxAccept, _ := strconv.ParseInt(config.GetTaskOptions(ctx, "最大接受"), 10, 64)
	maxAccept = min(max(maxAccept, 1), 3)

	// 接受委托
	change, pos := 0, 1
	for LabelColor(ctx, fmt.Sprintf("委托-接受%d号位", maxAccept)) == ix.ColorMissionEmpty {
		if LabelColor(ctx, fmt.Sprintf("委托-发布%d号位", pos)) == ix.ColorMissionHard {
			LabelClick(ctx, fmt.Sprintf("委托-发布%d号位", pos))
			LabelWaitClick(ctx, "委托-发布接受", 5*time.Second)
			continue
		}
		pos++
		// 判定完成一页
		if pos > 5 {
			if change < int(maxChange) {
				LabelClick(ctx, "委托-切换")
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
		if LabelColor(ctx, mission) == ix.ColorMissionExist {
			LabelClick(ctx, mission)
			if LabelCheck(ctx, "委托-接受佣兵团") {
				LabelClick(ctx, "委托-接受佣兵团")
				LabelWaitClick(ctx, "委托-领取奖励", 5*time.Second)
				LabelClick(ctx, "委托-接受1号位")
			}
		}
	}

	clr := ix.Color{R: 56, G: 221, B: 219}
	if LabelColor(ctx, "委托-福利") == clr {
		LabelClick(ctx, "委托-福利")
		LabelWaitClick(ctx, "委托-福利-购买", 5*time.Second)
	}
	log.Info(ctx, "任务入场")
	BackWorld(ctx)
}

// gmwts 购买委托书
func gmwts(ctx context.Context) {
	list := config.GetTaskOptions(ctx, "商品列表")
	if list == "" {
		return
	}
	goods := strings.Split(list, ",")
	LabelWaitClick(ctx, "世界-商店", 5*time.Second)

	for _, good := range goods {
		LabelWaitClick(ctx, "商店-收藏-"+good, 5*time.Second)
		LabelWait(ctx, "商店-商品详情-收藏图标", 5*time.Second)
		if LabelColor(ctx, "商店-商品详情-购买") == ix.ColorButtonOrange {
			LabelClick(ctx, "商店-商品详情-购买")
			LabelWaitClick(ctx, "商店-购买道具-购买", 5*time.Second)
		} else {
			LabelClick(ctx, "商店-商品详情-关闭")
		}
	}

	BackWorld(ctx)
}
