package main

import (
	"maple-robot/log"
	"maple-robot/scripts"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Roles []*RoleConfig `yaml:"roles"`
}

type RoleConfig struct {
	Index int                  `yaml:"index"`
	Skip  bool                 `yaml:"skip,omitempty"`
	Flags map[string]string    `yaml:"flags,omitempty"`
	Times map[string]time.Time `yaml:"times,omitempty"`
}

func (o *RoleConfig) Record(name string) {
	if o.Times == nil {
		o.Times = map[string]time.Time{}
	}
	o.Times[name] = time.Now()
}

func (o *RoleConfig) Already(name string) bool {
	return o.Times[name].Format("2006-01-02") == time.Now().Format("2006-01-02")
}

func loadConfig() *Config {
	bytes, _ := os.ReadFile("config.yaml")
	cfg := &Config{}
	yaml.Unmarshal(bytes, cfg)
	return cfg
}

func saveConfig(cfg *Config) {
	bytes, _ := yaml.Marshal(cfg)
	os.WriteFile("config.yaml", bytes, 0644)
}
func main() {

	cfg := loadConfig()
	for _, role := range cfg.Roles {
		if role.Skip || role.Already("完成") {
			log.Infof("角色%d skip\n", role.Index)
			continue
		}
		role.Record("开始")
		saveConfig(cfg)
		scripts.RoleSelect(role.Index/100, role.Index%100)
		scripts.LabelClick("世界-技能8")

		// 自动战斗世界
		scripts.LabelClick("世界-自动战斗")
		scripts.LabelClick("自动战斗-使用")
		scripts.LabelClick("自动战斗-关闭")

		// 公会签到
		scripts.LabelWaitClick("世界-导航", 5*time.Second)
		scripts.LabelWait("导航-枫小喵", 5*time.Second)
		scripts.LabelClick("导航-公会")
		scripts.LabelWaitClick("公会-领地树", 5*time.Second)
		scripts.LabelWait("公会-领地树-祝福领取取消", 5*time.Second)
		scripts.LabelClick("公会-领地树-祝福领取确定")
		scripts.LabelClick("公会-领地树-祝福领取取消")
		scripts.LabelClick("公会-领取")
		scripts.BackWorld()

		// 领取签到奖励
		scripts.LabelWaitClick("世界-导航", 5*time.Second)
		scripts.LabelWait("导航-枫小喵", 5*time.Second)
		scripts.LabelClick("导航-邮箱")
		scripts.LabelWait("邮箱-全部领取", 5*time.Second)
		scripts.LabelClick("邮箱-个人")
		scripts.LabelClick("邮箱-全部领取")
		scripts.LabelClick("邮箱-全部领取-确定")
		scripts.BackWorld()

		// 天空岛贸易
		if tkdmy := role.Flags["天空岛贸易"]; tkdmy != "" && !role.Already("天空岛贸易") {
			scripts.LabelWaitClick("世界-导航", 5*time.Second)
			scripts.LabelWait("导航-枫小喵", 5*time.Second)
			scripts.LabelClick("导航-天空岛贸易")
			scripts.LabelWaitClick("天空岛贸易-前往获取战利品", 5*time.Second)
			scripts.LabelWaitClick("天空岛贸易-入场", 5*time.Second)
			role.Record("天空岛贸易")
			saveConfig(cfg)
			scripts.LabelWait("太初森林-今日", 5*time.Second)
			scripts.LabelClick("太初森林-精灵")
			scripts.LabelWaitClick("副本-退出", 5*time.Second)
			scripts.LabelWaitClick("太初森林-退出确认", 5*time.Second)
			scripts.LabelWaitClick("太初森林-结果确认", 5*time.Second)
		}

		if wldc := role.Flags["武陵道场"]; wldc != "" && !role.Already("武陵道场") {
			scripts.LabelWaitClick("世界-导航", 5*time.Second)
			scripts.LabelWait("导航-枫小喵", 5*time.Second)
			scripts.LabelClick("导航-日常")
			scripts.LabelWait("日常-进度", 5*time.Second)
			scripts.LabelClick("日常-简化模式5号")
			scripts.LabelWaitClick("武陵道场-入场", 5*time.Second)
			scripts.LabelWaitClick("武陵道场-进入", 5*time.Second)
			role.Record("武陵道场")
			saveConfig(cfg)
			scripts.LabelWaitClick("副本-退出", 15*time.Second)
			time.Sleep(8 * time.Second)
			scripts.LabelWaitClick("武陵道场-退出", 5*time.Second)
			scripts.LabelWaitClick("武陵道场-离开", 10*time.Second)
		}

		// 周常副本
		if zcfb := role.Flags["周常副本"]; zcfb != "" && !role.Already("周常副本") {
			scripts.LabelWaitClick("世界-导航", 5*time.Second)
			scripts.LabelWait("导航-枫小喵", 5*time.Second)
			scripts.LabelClick("导航-日常")
			scripts.LabelWait("日常-进度", 5*time.Second)
			scripts.LabelClick("日常-简化模式3号")
			scripts.LabelWait("周常副本-入场", 5*time.Second)
			if time.Now().Weekday() == time.Saturday || time.Now().Weekday() == time.Sunday {
				scripts.LabelClick("周常副本-星期五")
			}
			scripts.LabelClick("周常副本-入场")
			scripts.LabelWaitClick("周常副本-入场-确定", 5*time.Second)
			role.Record("周常副本")
			saveConfig(cfg)
			scripts.LabelWait("副本-退出", 15*time.Second)
			scripts.LabelWaitClick("周常副本-副本结算-退出", 90*time.Second)
		}

		// 精英副本
		if jyfb := role.Flags["精英副本"]; jyfb != "" && !role.Already("精英副本") {
			scripts.LabelWaitClick("世界-导航", 5*time.Second)
			scripts.LabelWait("导航-枫小喵", 5*time.Second)
			scripts.LabelClick("导航-日常")
			scripts.LabelWait("日常-进度", 5*time.Second)
			scripts.LabelClick("日常-简化模式2号")
			scripts.LabelWaitClick("精英副本-快速组队", 5*time.Second)
			scripts.LabelWaitClick("精英副本-入场-确定", 5*time.Second)
			scripts.LabelWait("副本-退出", 60*time.Second)
			scripts.LabelWait("副本-麦克风", 15*time.Second)
			role.Record("精英副本")
			saveConfig(cfg)
			scripts.LabelWaitClick("精英副本-副本结算-离开", 180*time.Second)
		}

		// 材料副本
		if clfb := role.Flags["材料副本"]; clfb != "" && !role.Already("材料副本") {
			scripts.LabelWaitClick("世界-导航", 5*time.Second)
			scripts.LabelWait("导航-枫小喵", 5*time.Second)
			scripts.LabelClick("导航-日常")
			scripts.LabelWait("日常-进度", 5*time.Second)
			scripts.LabelClick("日常-简化模式1号")
			scripts.LabelWait("材料副本-入场", 5*time.Second)
			scripts.LabelClick("材料副本-" + clfb)
			scripts.LabelClick("材料副本-入场")
			scripts.LabelWaitClick("材料副本-入场-确定", 5*time.Second)
			if _, ok := role.Flags["材料副本-全部解除"]; ok {
				scripts.LabelWaitClick("材料副本-进入副本-全部解除", 5*time.Second)
			}
			scripts.LabelWaitClick("材料副本-进入副本-确定", 5*time.Second)
			role.Record("材料副本")
			saveConfig(cfg)
			scripts.LabelClick("世界-技能8")
			scripts.LabelWait("副本-退出", 15*time.Second)

			// 送人气
			{
				scripts.LabelWaitClick("世界-导航", 5*time.Second)
				scripts.LabelWait("导航-枫小喵", 5*time.Second)
				scripts.LabelClick("导航-社交")
				scripts.LabelWaitClick("社交-微信好友", 5*time.Second)
				scripts.LabelWait("社交-微信好友-背景", 5*time.Second)
				scripts.LabelClick("社交-微信好友-战斗力")
				scripts.LabelWaitClick("社交-微信好友-送人气给3号", 5*time.Second)
				scripts.LabelClick("社交-送人气")
				scripts.LabelClick("社交-送人气确定")
				scripts.LabelClick("社交-送人气不提醒")
				scripts.LabelClick("社交-送人气关闭")
				scripts.BackWorld()
			}

			// 领取日常奖励
			{
				scripts.LabelWaitClick("世界-导航", 5*time.Second)
				scripts.LabelWait("导航-枫小喵", 5*time.Second)
				scripts.LabelClick("导航-成长")
				scripts.LabelClick("成长-每日任务")
				scripts.LabelClick("成长-每日任务-全部领取")
				scripts.LabelClick("成长-领取确认")
				scripts.LabelClick("成长-每日任务-全部领取")
				scripts.LabelClick("成长-领取确认")
				scripts.LabelClick("成长-每周任务")
				scripts.LabelClick("成长-每周任务-全部领取")
				scripts.LabelClick("成长-领取确认")
				scripts.LabelClick("成长-每周任务-全部领取")
				scripts.LabelClick("成长-领取确认")
				scripts.BackWorld()
			}

			// 补充训练时间
			{
				scripts.LabelWaitClick("世界-导航", 5*time.Second)
				scripts.LabelWait("导航-枫小喵", 5*time.Second)
				scripts.LabelClick("导航-角色训练场")
				scripts.LabelWaitClick("角色训练场-补充时间", 5*time.Second)
				scripts.LabelClick("角色训练场-补充时间-使用")
				scripts.LabelClick("角色训练场-补充时间-使用")
				scripts.LabelWaitClick("角色训练场-补充时间-确定", 5*time.Second)
				scripts.BackWorld()
			}

			// 公会聊天
			// {
			// 	scripts.LabelClick("世界-聊天栏")
			// 	scripts.LabelWaitClick("聊天栏-公会", 5*time.Second)
			// 	scripts.LabelWaitClick("聊天栏-表情", 5*time.Second)
			// 	scripts.LabelWaitClick("聊天栏-表情-害羞", 5*time.Second)
			// 	scripts.LabelWaitClick("聊天栏-发送", 5*time.Second)
			// 	scripts.LabelClick("聊天栏-发送")
			// 	scripts.BackWorld()
			// }

			// 领取公会奖励
			{
				scripts.LabelWaitClick("世界-导航", 5*time.Second)
				scripts.LabelWait("导航-枫小喵", 5*time.Second)
				scripts.LabelClick("导航-公会")
				scripts.LabelClick("公会-公会任务")
				scripts.LabelClick("公会-公会任务-每日任务")
				scripts.LabelClick("公会-公会任务-全部领取")
				scripts.LabelClick("公会-公会任务-每周任务")
				scripts.LabelClick("公会-公会任务-全部领取")
				scripts.BackWorld()
			}

			scripts.LabelWaitClick("材料副本-副本结算-退出", 300*time.Second)
		}

		// 奈特的金字塔
		if jzt := role.Flags["奈特的金字塔"]; jzt != "" && !role.Already("奈特的金字塔") {
			scripts.LabelWaitClick("世界-导航", 5*time.Second)
			scripts.LabelWait("导航-枫小喵", 5*time.Second)
			scripts.LabelClick("导航-日常")
			scripts.LabelWait("日常-进度", 5*time.Second)
			scripts.LabelClick("日常-简化模式4号")
			scripts.LabelWaitClick("奈特的金字塔-快速组队", 5*time.Second)
			scripts.LabelWaitClick("奈特的金字塔-快速组队-入场确定", 5*time.Second)
			scripts.LabelWait("副本-退出", 60*time.Second)
			scripts.LabelWait("副本-麦克风", 15*time.Second)
			role.Record("奈特的金字塔")
			saveConfig(cfg)
			scripts.LabelWaitClick("奈特的金字塔-副本结算-退出", 180*time.Second)
		}

		scripts.WorldReturnRole()
		role.Record("完成")
		saveConfig(cfg)
	}
}
