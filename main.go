package main

import (
	"fmt"
	"image/color"
	"os"
	"rcemap/script"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/flopp/go-findfont"
)

var topWindow fyne.Window

// var swi *widget.RadioGroup
var sel *widget.Select

func init() {
	//设置中文字体
	fontPaths := findfont.List()
	for _, path := range fontPaths {
		if strings.Contains(path, "msyhbd.ttf") || strings.Contains(path, "simhei.ttf") || strings.Contains(path, "simsun.ttc") || strings.Contains(path, "simkai.ttf") || strings.Contains(path, "simfang.ttf") {
			err := os.Setenv("FYNE_FONT", path)
			if err != nil {
				return
			}
			break
		}
	}
}

type CustomTheme struct {
	textSize float32
}

func (c CustomTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(name, variant)
}

func (c CustomTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (c CustomTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (c CustomTheme) Size(name fyne.ThemeSizeName) float32 {
	// 修改默认字体大小
	if name == theme.SizeNameText {
		return c.textSize
	}
	return theme.DefaultTheme().Size(name)
}

/*
func init是设置字体

custom部分是自定义主题

最顶上的var是全局变量省事就放最上面了
*/

func main() {

	a := app.New() // 创建一个具有指定ID的Fyne应用实例
	customTheme := &CustomTheme{textSize: 16}
	a.Settings().SetTheme(customTheme)

	icon, _ := fyne.LoadResourceFromPath("icon.ico") //LOGO图标
	a.SetIcon(icon)

	w := a.NewWindow("RCEmap") // 创建一个名为 "RCEmap" 的窗口

	//列表
	var data = []string{"PHP", "ffuf"} //列表,一个php,一个ffuf
	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i])
		})

	content := container.NewStack()

	// 设置列表项被选中时的回调函数
	list.OnSelected = func(id widget.ListItemID) {
		content.Objects = []fyne.CanvasObject{} // 清空之前的内容
		switch data[id] {
		case "PHP":
			// 创建并添加 PHP 相关的内容

			URLentry := widget.NewEntry()
			URLentry.SetPlaceHolder("请输入url(如https://www.baidu.com)")

			PORTentry := widget.NewEntry()
			PORTentry.SetPlaceHolder("请输入端口(如:80,且80或443一定要输入,这里不会默认,一定要加':')")
			//前两个是必备条件,url和端口的设置

			/*
				autoContent自动模式容器
				manualContent手动模式容器
				sel选择方式,各手动选择更改写进sel
			*/

			canshu := widget.NewEntry()
			canshu.SetPlaceHolder("请输入参数点")

			command := widget.NewEntry()
			command.SetPlaceHolder("请输入要执行的命令")

			automethod := container.NewVBox() //自动模式下容纳GET与POST二选一的容器
			GETPOST := container.NewVBox()    //GET与POST二选一选择后显示的容器

			//选GET之后弹出来的容器
			GET := container.NewVBox(
				canshu,
				command,
			)

			//选POST之后弹出来的容器
			POST := container.NewVBox(
				canshu,
				command,
			)

			var getpost1 string

			automethodradio := widget.NewRadioGroup([]string{"GET", "POST"}, func(getpost string) {
				switch getpost {
				case "GET":
					getpost1 = "GET"
					GETPOST.Objects = []fyne.CanvasObject{GET}
				case "POST":
					getpost1 = "POST"
					GETPOST.Objects = []fyne.CanvasObject{POST}
				}
				GETPOST.Refresh()
			})

			automethod.Add(automethodradio) //automethod容器加个选择器

			var Version string

			Version = "5"

			phpRadio := widget.NewRadioGroup([]string{"php5.x", "php7.x"}, func(v string) {
				switch v {
				case "php5.x":

					Version = "5"

				case "php7.x":

					Version = "7"
				}

			})

			phpVersion := container.NewVBox(phpRadio)

			autoContent := container.NewVBox(
				widget.NewLabel("自动模式(该模式不准确,若没有正确执行请切换至手动模式)\n(黑名单过滤,环境变量,少量字符RCE这三种请切换至手动)"), //输出"自动模式"
				phpVersion,
				automethod,
				GETPOST,
			)

			// 创建手动模式的选择组件和内容显示区域
			manualContentBox := container.NewVBox()

			var exploit string

			var radio *widget.RadioGroup
			var sele *widget.Select
			var label *widget.Label

			guolventry := widget.NewEntry()
			guolventry.SetPlaceHolder("请在此输入过滤")

			sel = widget.NewSelect([]string{"无数字字母RCE", "黑名单绕过", "少量字符RCE", "环境变量构造"}, func(value string) {
				manualContentBox.Objects = nil
				label = widget.NewLabel("")

				four2one := container.NewVBox()

				switch value {

				case "无数字字母RCE":

					label = widget.NewLabel("Tips:当过滤了$时请选择进阶")

					sele = widget.NewSelect([]string{"自增", "异或", "ff", "进阶"}, func(m string) {
						switch m {

						case "自增":

							exploit = "zizeng"

						case "异或":

							exploit = "xor"

						case "ff":

							exploit = "ff"

						case "进阶":

							exploit = "进阶"
						}
					})

					sele.PlaceHolder = "请选择一种方法"
					radiocontent := container.NewVBox()
					radio = widget.NewRadioGroup([]string{"eval", "system"}, func(ti string) {

						switch ti {
						case "eval":

							radiocontent.Objects = []fyne.CanvasObject{sele}

						case "system":

							label = widget.NewLabel("只有一种bashfuck的方法")
							exploit = "bashfuck"

							radiocontent.Objects = []fyne.CanvasObject{label}

						}
						radiocontent.Refresh()
					})

					four2one = container.NewVBox(label, radio, radiocontent)

				case "黑名单绕过":

					label = widget.NewLabel("这是黑名单绕过")

				case "少量字符RCE":

					label = widget.NewLabel("这是少量字符RCE")

				case "环境变量构造":

					label = widget.NewLabel("这是环境变量构造")

				}

				manualContentBox.Objects = []fyne.CanvasObject{

					four2one,
				}
				manualContentBox.Refresh()

			})

			sel.PlaceHolder = "请选择一种题型"

			manualContent := container.NewVBox(
				widget.NewLabel("手动模式"),
				guolventry,
				sel,
				manualContentBox,
			)

			//phpContent是选择自动模式和手动模式后出现的容器
			phpContent := container.NewVBox(autoContent)
			var rrraaadio string
			modeRadio := widget.NewRadioGroup([]string{"自动模式", "手动模式"}, func(selected string) {
				switch selected {
				case "自动模式":

					rrraaadio = "自动模式"
					phpContent.Objects = []fyne.CanvasObject{autoContent}

				case "手动模式":

					rrraaadio = "手动模式"
					phpContent.Objects = []fyne.CanvasObject{manualContent}

				}
				phpContent.Refresh()
			})

			// 初始化回显标签
			huixianlabel := widget.NewLabel("这是回显区")
			newlabel := widget.NewLabel("")

			// 容器，包含回显标签
			huixian := container.NewVBox(

				huixianlabel, //这里是当前环境,题目源码和题目过滤
				newlabel,
			)

			// 按钮,主要功能区
			button := widget.NewButton("START", func() {
				URL := URLentry.Text + PORTentry.Text
				huixianlabel.SetText("访问URL+端口为: " + URL) // 更新回显标签的文本

				if rrraaadio == "自动模式" {

					jianjie := `自动模式
请求模式为` + getpost1 + "\n访问URL+端口为: " + URL + "\nphp版本为" + Version

					if canshu.Text == "" {
						huixianlabel.SetText(jianjie + "\n请输入全部参数")
					} else {
						cleanText, guolv := script.Test(URL, Version, command.Text, canshu.Text, getpost1)
						huixianlabel.SetText(jianjie + "题目源码为:\n" + cleanText + "检测到过滤为:\n" + guolv)

						script.Damn(URL, Version, command.Text, canshu.Text, getpost1, cleanText, guolv, newlabel)
					}

				} else if rrraaadio == "手动模式" {

					jianjie := `手动模式
请求模式为` + getpost1 + "\n访问URL+端口为: " + URL + "\nphp版本为" + Version

					guolv := guolventry.Text

					huixianlabel.SetText(jianjie + "过滤为:\n" + guolv)

					if !strings.Contains(guolv, "$") {

						switch exploit {

						//前三个是eval

						case "zizeng":

							script.Zizeng(URL, command.Text, canshu.Text, newlabel)

						case "xor":

							script.Xor(URL, command.Text, Version, canshu.Text, newlabel)

						case "ff":

							script.Ff(URL, command.Text, guolv, canshu.Text, getpost1, newlabel)

						//后三个是system

						case "bashfuck":

							booL := "0"
							script.Bashfuck(URL, command.Text, guolv, canshu.Text, getpost1, newlabel, booL)

						case "black":

							//黑名单

						case "pwd":

							//环境变量

						case "Fewchar":

							//少字符

						}

					} else {

						//这是ban了'$'的

						switch exploit {

						case "进阶":

							script.Noshuzievaljinjie(URL, command.Text, Version, canshu.Text, getpost1, guolv, newlabel)

						case "Fewchar":

						}

					}
				}

			})

			mainContent := container.NewVBox(
				URLentry,
				PORTentry,
				modeRadio,
				phpContent,
			)

			huixianscroll := container.NewScroll(huixian) //给回显区加个滚动

			//mainsplit是中间容器上下分割
			mainsplit := container.NewVSplit(mainContent, huixianscroll)
			mainsplit.Offset = 0.8

			//phpsplit是php部分右侧content与button的分割线
			phpsplit := container.NewHSplit(mainsplit, button)
			phpsplit.Offset = 0.9
			content.Add(phpsplit)

		case "ffuf":
			// 创建并添加 ffuf 相关的内容
			ffufContent := widget.NewLabel("未完待续")

			content.Add(ffufContent)
		}

		content.Refresh() // 刷新容器以显示新内容

	}

	//themes是list下面的设置字体大小按钮

	fontSizeLabel := widget.NewLabel("字体大小: " + fmt.Sprintf("%.0f", customTheme.textSize))

	themes := container.NewGridWithColumns(2,

		widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
			customTheme.textSize -= 1
			a.Settings().SetTheme(customTheme)
			content.Refresh()
			fontSizeLabel.SetText("字体大小: " + fmt.Sprintf("%.0f", customTheme.textSize)) // 更新标签文本

		}),

		widget.NewButtonWithIcon("", theme.NavigateNextIcon(), func() {

			customTheme.textSize += 1
			a.Settings().SetTheme(customTheme)
			content.Refresh()
			fontSizeLabel.SetText("字体大小: " + fmt.Sprintf("%.0f", customTheme.textSize)) // 更新标签文本

		}),
	)

	down := container.NewVBox(fontSizeLabel, themes)

	LIST := container.NewBorder(nil, down, nil, nil, list) //Border(上,下,左,右,中)

	// 创建一个水平分割容器，将列表和内容容器放置在其中
	split := container.NewHSplit(LIST, content)
	split.Offset = 0.161 // 设置分割容器的初始分割位置

	//后缀级代码,不是很需要更改
	w.SetContent(split)
	w.Resize(fyne.NewSize(840, 600)) //设置窗口大小
	topWindow = w                    // 将 topWindow 设置为当前窗口
	topWindow.ShowAndRun()
}
