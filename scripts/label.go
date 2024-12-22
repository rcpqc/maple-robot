package scripts

import (
	"maple-robot/ix"
	"maple-robot/log"
	"strings"
	"time"
)

type Label struct {
	Position   ix.Position
	Color      ix.Color
	ClickDelay time.Duration
}

var labels = map[string]*Label{
	// 世界
	"世界-导航":   {ix.Position{X: 939, Y: 19}, ix.Color{R: 255, G: 255, B: 255}, time.Second},
	"世界-经验条":  {ix.Position{X: 2, Y: 538}, ix.Color{R: 254, G: 167, B: 14}, 0},
	"世界-电量":   {ix.Position{X: 204, Y: 528}, ix.Color{R: 122, G: 195, B: 0}, 0},
	"世界-技能8":  {ix.Position{X: 820, Y: 500}, ix.Color{}, time.Second},
	"世界-自动战斗": {ix.Position{X: 316, Y: 495}, ix.Color{}, time.Second},
	"世界-聊天栏":  {ix.Position{X: 475, Y: 520}, ix.Color{}, 2 * time.Second},

	"聊天栏-公会":    {ix.Position{X: 24, Y: 120}, ix.Color{R: 11, G: 13, B: 17}, time.Second},
	"聊天栏-表情":    {ix.Position{X: 284, Y: 506}, ix.ColorWhite, time.Second},
	"聊天栏-表情-害羞": {ix.Position{X: 745, Y: 95}, ix.Color{R: 255, G: 219, B: 85}, time.Second},
	"聊天栏-发送":    {ix.Position{X: 351, Y: 507}, ix.ColorButtonGray, time.Second},

	"通用-关闭":          {ix.Position{X: 931, Y: 31}, ix.Color{}, time.Second},
	"通用-游戏结束-否":      {ix.Position{X: 333, Y: 385}, ix.ColorButtonBlue, time.Second},
	"通用-游戏结束-返回角色界面": {ix.Position{X: 478, Y: 385}, ix.ColorButtonBlue, time.Second},
	"通用-游戏结束-是":      {ix.Position{X: 628, Y: 385}, ix.ColorButtonOrange, time.Second},

	// 副本
	"副本-麦克风": {ix.Position{X: 874, Y: 204}, ix.ColorWhite, time.Second},
	"副本-退出":  {ix.Position{X: 930, Y: 214}, ix.ColorButtonExitDungeon, time.Second},

	// 导航
	"导航-枫小喵":   {ix.Position{X: 707, Y: 501}, ix.Color{R: 253, G: 235, B: 218}, time.Second},
	"导航-公会":    {ix.Position{X: 645, Y: 230}, ix.Color{}, 3 * time.Second},
	"导航-成长":    {ix.Position{X: 820, Y: 90}, ix.Color{}, 3 * time.Second},
	"导航-日常":    {ix.Position{X: 860, Y: 230}, ix.Color{}, 3 * time.Second},
	"导航-社交":    {ix.Position{X: 645, Y: 360}, ix.Color{}, 5 * time.Second},
	"导航-邮箱":    {ix.Position{X: 794, Y: 25}, ix.Color{}, 3 * time.Second},
	"导航-关闭":    {ix.Position{X: 925, Y: 21}, ix.Color{}, 3 * time.Second},
	"导航-更改角色":  {ix.Position{X: 881, Y: 497}, ix.Color{}, 3 * time.Second},
	"导航-角色训练场": {ix.Position{X: 815, Y: 424}, ix.Color{}, 3 * time.Second},
	"导航-天空岛贸易": {ix.Position{X: 900, Y: 424}, ix.Color{}, 3 * time.Second},

	// 成长

	"成长-标题":        {ix.Position{X: 75, Y: 32}, ix.ColorTitleGray, time.Second},
	"成长-每日任务":      {ix.Position{X: 80, Y: 100}, ix.Color{}, time.Second},
	"成长-每周任务":      {ix.Position{X: 80, Y: 160}, ix.Color{}, time.Second},
	"成长-每日任务-全部领取": {ix.Position{X: 880, Y: 500}, ix.ColorButtonOrange, 3 * time.Second},
	"成长-每周任务-全部领取": {ix.Position{X: 880, Y: 500}, ix.ColorButtonOrange, 3 * time.Second},
	"成长-领取确认":      {ix.Position{X: 560, Y: 383}, ix.ColorButtonOrange, 3 * time.Second},

	// 日常
	"日常-进度":                {ix.Position{X: 706, Y: 93}, ix.ColorButtonGray, time.Second},
	"日常-寻找队伍":              {ix.Position{X: 853, Y: 91}, ix.ColorButtonOrange, time.Second},
	"日常-进度-金币":             {ix.Position{X: 716, Y: 138}, ix.Color{R: 255, G: 232, B: 70}, time.Second},
	"日常-进度-关闭":             {ix.Position{X: 858, Y: 139}, ix.ColorWhite, time.Second},
	"日常-进度-怪物乐园跳关":         {ix.Position{X: 808, Y: 425}, ix.Color{R: 255, G: 123, B: 80}, time.Second},
	"日常-进度-怪物乐园跳关-标题":      {ix.Position{X: 161, Y: 68}, ix.ColorTitleGray, time.Second},
	"日常-进度-怪物乐园跳关-隐匿痕迹":    {ix.Position{X: 200, Y: 250}, ix.Color{R: 187, G: 194, B: 202}, time.Second},
	"日常-进度-怪物乐园跳关-大木林":     {ix.Position{X: 400, Y: 250}, ix.Color{R: 187, G: 194, B: 202}, time.Second},
	"日常-进度-怪物乐园跳关-小蘑菇丛林":   {ix.Position{X: 600, Y: 250}, ix.Color{R: 187, G: 194, B: 202}, time.Second},
	"日常-进度-怪物乐园跳关-停止施工的工地": {ix.Position{X: 200, Y: 310}, ix.Color{R: 187, G: 194, B: 202}, time.Second},
	"日常-进度-怪物乐园跳关-怪物乐园沙滩":  {ix.Position{X: 400, Y: 310}, ix.Color{R: 187, G: 194, B: 202}, time.Second},
	"日常-进度-怪物乐园跳关-使用战斗跳关券": {ix.Position{X: 480, Y: 480}, ix.ColorButtonOrange, time.Second},
	"日常-进度-怪物乐园跳关-入场确认":    {ix.Position{X: 540, Y: 400}, ix.ColorButtonOrange, time.Second},
	"日常-进度-怪物乐园跳关-结算确认":    {ix.Position{X: 430, Y: 430}, ix.ColorButtonBlue, time.Second},

	"日常-简化模式1号":  {ix.Position{X: 93, Y: 235}, ix.Color{}, time.Second},
	"日常-简化模式2号":  {ix.Position{X: 93, Y: 438}, ix.Color{}, time.Second},
	"日常-简化模式3号":  {ix.Position{X: 260, Y: 235}, ix.Color{}, time.Second},
	"日常-简化模式4号":  {ix.Position{X: 260, Y: 438}, ix.Color{}, time.Second},
	"日常-简化模式5号":  {ix.Position{X: 427, Y: 235}, ix.Color{}, time.Second},
	"日常-简化模式6号":  {ix.Position{X: 427, Y: 438}, ix.Color{}, time.Second},
	"日常-简化模式7号":  {ix.Position{X: 594, Y: 235}, ix.Color{}, time.Second},
	"日常-简化模式8号":  {ix.Position{X: 594, Y: 438}, ix.Color{}, time.Second},
	"日常-简化模式9号":  {ix.Position{X: 761, Y: 235}, ix.Color{}, time.Second},
	"日常-简化模式10号": {ix.Position{X: 761, Y: 438}, ix.Color{}, time.Second},
	"日常-简化模式11号": {ix.Position{X: 928, Y: 235}, ix.Color{}, time.Second},
	"日常-简化模式12号": {ix.Position{X: 928, Y: 438}, ix.Color{}, time.Second},

	// 公会
	"公会-领取":         {ix.Position{X: 862, Y: 92}, ix.Color{}, 2 * time.Second},
	"公会-举报":         {ix.Position{X: 456, Y: 80}, ix.Color{R: 188, G: 60, B: 87}, 0},
	"公会-领地树":        {ix.Position{X: 501, Y: 511}, ix.Color{R: 194, G: 205, B: 216}, time.Second},
	"公会-领地树-祝福领取确定": {ix.Position{X: 555, Y: 455}, ix.ColorButtonOrange, time.Second},
	"公会-领地树-祝福领取取消": {ix.Position{X: 326, Y: 455}, ix.ColorButtonBlue, time.Second},
	"公会-公会任务":       {ix.Position{X: 100, Y: 370}, ix.Color{R: 43, G: 54, B: 70}, time.Second},
	"公会-公会任务-每日任务":  {ix.Position{X: 216, Y: 92}, ix.Color{}, time.Second},
	"公会-公会任务-每周任务":  {ix.Position{X: 370, Y: 92}, ix.Color{}, time.Second},
	"公会-公会任务-全部领取":  {ix.Position{X: 830, Y: 510}, ix.ColorButtonOrange, 3 * time.Second},

	// 邮箱
	"邮箱-通用": {ix.Position{X: 143, Y: 101}, ix.Color{R: 238, G: 117, B: 70}, time.Second},
	"邮箱-个人": {ix.Position{X: 328, Y: 98}, ix.Color{}, time.Second},
	// "邮箱-今天不再显示":  {ix.Position{X: 560, Y: 350}, ix.Color{}, time.Second},
	// "邮箱-未领取确认":   {ix.Position{X: 480, Y: 400}, ix.Color{}, time.Second},
	"邮箱-全部领取":    {ix.Position{X: 472, Y: 491}, ix.ColorButtonOrange, 3 * time.Second},
	"邮箱-全部领取-确定": {ix.Position{X: 480, Y: 400}, ix.ColorButtonOrange, time.Second},

	// 社交
	"社交-微信好友":        {ix.Position{X: 650, Y: 90}, ix.Color{R: 132, G: 222, B: 66}, 3 * time.Second},
	"社交-微信好友-背景":     {ix.Position{X: 151, Y: 198}, ix.ColorWhite, time.Second},
	"社交-微信好友-战斗力":    {ix.Position{X: 716, Y: 139}, ix.Color{}, 2 * time.Second},
	"社交-微信好友-送人气给3号": {ix.Position{X: 917, Y: 302}, ix.Color{R: 255, G: 123, B: 80}, time.Second},
	"社交-微信好友-送人气给4号": {ix.Position{X: 917, Y: 363}, ix.Color{R: 255, G: 123, B: 80}, time.Second},
	"社交-送人气":         {ix.Position{X: 379, Y: 268}, ix.Color{}, time.Second},
	"社交-送人气确定":       {ix.Position{X: 480, Y: 400}, ix.Color{}, time.Second},
	"社交-送人气不提醒":      {ix.Position{X: 370, Y: 400}, ix.Color{}, time.Second},
	"社交-送人气关闭":       {ix.Position{X: 693, Y: 110}, ix.Color{}, time.Second},

	// 自动战斗
	"自动战斗-使用": {ix.Position{X: 709, Y: 296}, ix.Color{}, time.Second},
	"自动战斗-关闭": {ix.Position{X: 747, Y: 86}, ix.Color{}, time.Second},

	// 材料副本
	"材料副本-阿里安特":      {ix.Position{X: 72, Y: 195}, ix.Color{}, time.Second},
	"材料副本-卡帕莱特":      {ix.Position{X: 72, Y: 245}, ix.Color{}, time.Second},
	"材料副本-神木村":       {ix.Position{X: 72, Y: 295}, ix.Color{}, time.Second},
	"材料副本-龙之峡谷":      {ix.Position{X: 72, Y: 345}, ix.Color{}, time.Second},
	"材料副本-玩具城":       {ix.Position{X: 72, Y: 395}, ix.Color{}, time.Second},
	"材料副本-武陵桃源":      {ix.Position{X: 72, Y: 445}, ix.Color{}, time.Second},
	"材料副本-入场":        {ix.Position{X: 830, Y: 500}, ix.ColorButtonOrange, time.Second},
	"材料副本-入场-确定":     {ix.Position{X: 560, Y: 400}, ix.ColorButtonOrange, time.Second},
	"材料副本-进入副本-全部解除": {ix.Position{X: 843, Y: 426}, ix.Color{R: 188, G: 60, B: 87}, time.Second},
	"材料副本-进入副本-确定":   {ix.Position{X: 560, Y: 480}, ix.ColorButtonOrange, time.Second},
	"材料副本-副本结算-退出":   {ix.Position{X: 228, Y: 423}, ix.ColorButtonBlue, 3 * time.Second},

	// 精英副本
	"精英副本-普通":        {ix.Position{X: 80, Y: 90}, ix.Color{R: 188, G: 60, B: 87}, time.Second},
	"精英副本-鬼怪蘑菇王":     {ix.Position{X: 94, Y: 147}, ix.Color{}, time.Second},
	"精英副本-创建房间":      {ix.Position{X: 625, Y: 500}, ix.ColorButtonBlue, time.Second},
	"精英副本-快速组队":      {ix.Position{X: 800, Y: 500}, ix.ColorButtonOrange, time.Second},
	"精英副本-入场-确定":     {ix.Position{X: 560, Y: 400}, ix.ColorButtonOrange, time.Second},
	"精英副本-集结地-开始":    {ix.Position{X: 154, Y: 164}, ix.ColorButtonOrange, time.Second},
	"精英副本-副本结算-单人离开": {ix.Position{X: 322, Y: 440}, ix.ColorButtonBlue, 3 * time.Second},
	"精英副本-副本结算-离开":   {ix.Position{X: 322, Y: 505}, ix.ColorButtonBlue, 3 * time.Second},
	// 周常副本
	"周常副本-星期五":     {ix.Position{X: 855, Y: 115}, ix.Color{}, time.Second},
	"周常副本-入场":      {ix.Position{X: 820, Y: 500}, ix.ColorButtonOrange, time.Second},
	"周常副本-入场-确定":   {ix.Position{X: 560, Y: 400}, ix.ColorButtonOrange, time.Second},
	"周常副本-副本结算-退出": {ix.Position{X: 267, Y: 421}, ix.ColorButtonBlue, 3 * time.Second},

	// 奈特的金字塔
	"奈特的金字塔-快速组队":      {ix.Position{X: 810, Y: 500}, ix.ColorButtonOrange, time.Second},
	"奈特的金字塔-快速组队-入场确定": {ix.Position{X: 622, Y: 400}, ix.ColorButtonOrange, time.Second},
	"奈特的金字塔-混沌":        {ix.Position{X: 106, Y: 215}, ix.Color{}, time.Second},
	"奈特的金字塔-副本结算-退出":   {ix.Position{X: 250, Y: 500}, ix.ColorButtonBlue, 3 * time.Second},

	// 天空岛贸易
	"天空岛贸易-船长":      {ix.Position{X: 623, Y: 190}, ix.Color{}, time.Second},
	"天空岛贸易-前往获取战利品": {ix.Position{X: 263, Y: 490}, ix.Color{R: 84, G: 143, B: 186}, time.Second},
	"天空岛贸易-入场":      {ix.Position{X: 809, Y: 483}, ix.ColorButtonOrange, time.Second},
	"天空岛贸易-退出":      {ix.Position{X: 926, Y: 214}, ix.Color{}, time.Second},
	"天空岛贸易-退出确认":    {ix.Position{X: 590, Y: 400}, ix.Color{}, time.Second},

	// 武陵道场
	"武陵道场-入场": {ix.Position{X: 800, Y: 500}, ix.ColorButtonOrange, time.Second},
	"武陵道场-进入": {ix.Position{X: 400, Y: 462}, ix.ColorButtonOrange, time.Second},
	"武陵道场-退出": {ix.Position{X: 555, Y: 406}, ix.ColorButtonOrange, time.Second},
	"武陵道场-离开": {ix.Position{X: 347, Y: 483}, ix.ColorButtonBlue, time.Second},
	// 太初森林
	"太初森林-精灵": {ix.Position{X: 477, Y: 133}, ix.Color{R: 222, G: 156, B: 143}, time.Second},
	// "太初森林-精灵": {ix.Position{X: 477, Y: 133}, ix.Color{R: 206, G: 107, B: 103}, time.Second},
	"太初森林-今日":   {ix.Position{X: 855, Y: 61}, ix.Color{R: 97, G: 220, B: 255}, time.Second},
	"太初森林-退出":   {ix.Position{X: 926, Y: 214}, ix.Color{}, time.Second},
	"太初森林-退出确认": {ix.Position{X: 560, Y: 400}, ix.ColorButtonOrange, 2 * time.Second},
	"太初森林-结果确认": {ix.Position{X: 450, Y: 430}, ix.ColorButtonBlue, 2 * time.Second},

	// 更改角色
	"更改角色-选择角色": {ix.Position{X: 269, Y: 480}, ix.Color{}, time.Second},

	// 角色选择
	"角色选择-服务器":    {ix.Position{X: 890, Y: 125}, ix.Color{R: 229, G: 199, B: 140}, 0},
	"角色选择-游戏开始":   {ix.Position{X: 800, Y: 490}, ix.Color{}, 3 * time.Second},
	"角色选择-上一页":    {ix.Position{X: 40, Y: 270}, ix.Color{}, 2 * time.Second},
	"角色选择-下一页":    {ix.Position{X: 640, Y: 270}, ix.Color{}, 2 * time.Second},
	"角色选择-1号角色":   {ix.Position{X: 200, Y: 180}, ix.Color{}, 100 * time.Millisecond},
	"角色选择-2号角色":   {ix.Position{X: 350, Y: 180}, ix.Color{}, 100 * time.Millisecond},
	"角色选择-3号角色":   {ix.Position{X: 490, Y: 180}, ix.Color{}, 100 * time.Millisecond},
	"角色选择-4号角色":   {ix.Position{X: 130, Y: 410}, ix.Color{}, 100 * time.Millisecond},
	"角色选择-5号角色":   {ix.Position{X: 280, Y: 410}, ix.Color{}, 100 * time.Millisecond},
	"角色选择-6号角色":   {ix.Position{X: 420, Y: 410}, ix.Color{}, 100 * time.Millisecond},
	"角色选择-7号角色":   {ix.Position{X: 570, Y: 410}, ix.Color{}, 100 * time.Millisecond},
	"角色选择-选中角色页1": {ix.Position{X: 315, Y: 495}, ix.ColorRolePageSelected, time.Second},
	"角色选择-选中角色页2": {ix.Position{X: 337, Y: 495}, ix.ColorRolePageSelected, time.Second},
	"角色选择-选中角色页3": {ix.Position{X: 359, Y: 495}, ix.ColorRolePageSelected, time.Second},
	"角色选择-选中角色页4": {ix.Position{X: 382, Y: 495}, ix.ColorRolePageSelected, time.Second},
	"角色选择-选中角色页5": {ix.Position{X: 405, Y: 495}, ix.ColorRolePageSelected, time.Second},
	// 角色训练场
	"角色训练场-补充时间":    {ix.Position{X: 819, Y: 384}, ix.ColorButtonBlue, 2 * time.Second},
	"角色训练场-补充时间-使用": {ix.Position{X: 485, Y: 240}, ix.ColorButtonYellow, time.Second},
	"角色训练场-补充时间-确定": {ix.Position{X: 431, Y: 469}, ix.ColorButtonOrange, time.Second},
}

func LabelCheck(names ...string) bool {
	log.Infof("LabelCheck - label(%s)\n", strings.Join(names, ", "))
	for _, name := range names {
		lbl, ok := labels[name]
		if !ok {
			log.Errorf("LabelCheck - label(%s) not found\n", name)
			return false
		}
		if !ix.GetPixel(lbl.Position).Equals(lbl.Color) {
			return false
		}
	}
	return true
}

func LabelWait(name string, timeout time.Duration) {
	log.Infof("LabelWait - label(%s) timeout=%v\n", name, timeout)
	lbl, ok := labels[name]
	if !ok {
		log.Errorf("LabelWait - label(%s) not found\n", name)
	}
	st := time.Now()
	ddl := st.Add(timeout)
	for {
		cur := ix.GetPixel(lbl.Position)
		if cur.Equals(lbl.Color) {
			log.Infof("LabelWait - label(%s) timeout=%v cost=%v\n", name, timeout, time.Since(st))
			return
		}
		if time.Now().After(ddl) {
			log.Warnf("LabelWait - label(%s) timeout, pos=%s, current=%s, target=%s\n", name, lbl.Position, cur, lbl.Color)
		}
		time.Sleep(time.Second)
	}
}

func LabelClick(name string) {
	log.Infof("LabelClick - label(%s)\n", name)
	lbl, ok := labels[name]
	if !ok {
		log.Errorf("LabelClick - label(%s) not found\n", name)
		return
	}
	ix.Tap(lbl.Position)
	if lbl.ClickDelay > 0 {
		time.Sleep(lbl.ClickDelay)
	}
}

func LabelWaitClick(name string, timeout time.Duration) {
	LabelWait(name, timeout)
	LabelClick(name)
}
