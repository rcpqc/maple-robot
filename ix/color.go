package ix

import (
	"fmt"
	"math"
)

type Color struct {
	R uint8
	G uint8
	B uint8
}

var (
	ColorRolePageSelected  = Color{255, 123, 82}  // 角色页选中
	ColorButtonOrange      = Color{238, 112, 70}  // 按钮橙色
	ColorButtonBlue        = Color{76, 135, 176}  // 按钮蓝色
	ColorButtonGray        = Color{97, 122, 149}  // 按钮灰色
	ColorButtonDisable     = Color{195, 195, 195} // 失效按钮
	ColorButtonYellow      = Color{255, 215, 65}  // 按钮黄色
	ColorButtonWine        = Color{188, 60, 87}   // 按钮酒红色
	ColorTitleGray         = Color{65, 80, 102}   // 标题栏灰色
	ColorWhite             = Color{255, 255, 255} // 白色
	ColorButtonExitDungeon = Color{255, 204, 0}   // 退出副本按钮橙色
	ColorButtonRedDot      = Color{252, 48, 15}   // 红点
)

func (o Color) Equals(c Color) bool {
	delta := 0.01
	r := math.Abs(float64(o.R)/255.0 - float64(c.R)/255.0)
	g := math.Abs(float64(o.G)/255.0 - float64(c.G)/255.0)
	b := math.Abs(float64(o.B)/255.0 - float64(c.B)/255.0)
	return r <= delta && g <= delta && b <= delta
}

func (o Color) String() string {
	return fmt.Sprintf("(%d,%d,%d)", o.R, o.G, o.B)
}
