package main

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
)

// 对应 CSS 变量（直接使用十六进制值）
const (
	TestColorRed    = "#f56c6c"
	TestColorGreen  = "#67c23a"
	TestColorBlue   = "#409eff"
	TestColorPrimary   = TestColorBlue  // 相当于 var(--test-color-blue)
	TestColorSuccess   = TestColorGreen // 相当于 var(--test-color-green)
	TestColorDanger    = TestColorRed   // 相当于 var(--test-color-red)
)

// 对应 SCSS 变量
const (
	Gray0 = "#aaaaaa"
	Gray1 = "#808080"
	Gray2 = "#4a4a4a"
	Gray3 = "#333333"
)

func main() {
	// 使用 lipgloss 创建带颜色的样式
	primaryStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(TestColorPrimary)).
		Bold(true).
		Padding(0, 1)

	successStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(TestColorSuccess))

	dangerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(TestColorDanger))

	grayStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(Gray2))

	// 打印示例
	fmt.Println(primaryStyle.Render("主要颜色 (Primary)"))
	fmt.Println(successStyle.Render("成功颜色 (Success)"))
	fmt.Println(dangerStyle.Render("危险颜色 (Danger)"))
	fmt.Println(grayStyle.Render("灰度2 - 中等文字"))
}