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
var RW string
var sel *widget.Select
var blackleixing string
var guolvleixing string
var choose string

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

			//是必备条件,url的设置

			/*
				autoContent自动模式容器
				manualContent手动模式容器
				sel选择方式,各手动选择更改写进sel
			*/

			canshu := widget.NewEntry()
			canshu.SetPlaceHolder("请输入参数点")

			xieru := widget.NewEntry()
			xieru.SetPlaceHolder("当需要第二个参数时,请写在这里,不然请为空")

			command := widget.NewEntry()
			command.SetPlaceHolder("请输入要执行的命令")

			automethod := container.NewVBox() //自动模式下容纳GET与POST二选一的容器
			GETPOST := container.NewVBox()    //GET与POST二选一选择后显示的容器

			//选GET之后弹出来的容器
			GET := container.NewVBox(
				canshu,
				xieru,
				command,
			)

			//选POST之后弹出来的容器
			POST := container.NewVBox(
				canshu,
				xieru,
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
			guolventry.SetPlaceHolder("请在此输入过滤(直接复制正则)")

			sel = widget.NewSelect([]string{"无数字字母RCE", "黑名单绕过", "少量字符RCE", "环境变量构造", "preg_replace/e的利用", "伪协议"}, func(value string) {
				manualContentBox.Objects = nil
				label = widget.NewLabel("")

				four2one := container.NewVBox()

				switch value {

				case "无数字字母RCE":

					label = widget.NewLabel("Tips:当过滤了$时请选择进阶")

					sele = widget.NewSelect([]string{"自增", "异或", "或", "ff", "进阶", "限制字符种类(固定)", "限制字符种类(自定义)"}, func(m string) {
						switch m {

						case "自增":

							exploit = "zizeng"

						case "异或":

							exploit = "xor"

						case "或":

							exploit = "or"

						case "ff":

							exploit = "ff"

						case "进阶":

							exploit = "进阶"

						case "限制字符种类(固定)":

							choose = "guding"
							exploit = "进阶plus"

						case "限制字符种类(自定义)":

							choose = "zidingyi"
							exploit = "进阶plus"
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

					es := widget.NewRadioGroup([]string{"eval", "system"}, func(ti string) {

						switch ti {
						case "eval":
							guolvleixing = "e"
						case "system":
							guolvleixing = "s"

						}

					})

					//one2content := container.NewVBox()
					one23 := widget.NewSelect([]string{"用户传命令执行", "执行固定逻辑", "根据输入的过滤回显出可用的函数自行执行"}, func(ti string) {
						switch ti {
						case "用户传命令执行":
							blackleixing = "chuan"

						case "执行固定逻辑":
							blackleixing = "guding"

						case "根据输入的过滤回显出可用的函数自行执行":
							blackleixing = "shoudong"
						}
					})

					one23.PlaceHolder = "请选择一种方法"

					four2one = container.NewVBox(

						es,
						one23,
						//one2content,
					)
					exploit = "black"

				case "少量字符RCE":

					label = widget.NewLabel("这是少量字符RCE")
					exploit = "Fewchar"

				case "环境变量构造":

					label = widget.NewLabel("这是环境变量构造")
					exploit = "pwd"

				case "preg_replace/e的利用":

					exploit = "replace"

				case "伪协议":

					exploit = "file_put"
					radi := widget.NewRadioGroup([]string{"写入", "读取", "请在前两个都无法执行的时候选择这个,如果这个也无效则是无效"}, func(value string) {
						switch value {
						case "写入":
							RW = "write"
						case "读取":
							RW = "read"
						case "请在前两个都无法执行的时候选择这个,如果这个也无效则是无效":
							RW = "plus"
						}
					})

					four2one = container.NewVBox(radi)
				}

				manualContentBox.Objects = []fyne.CanvasObject{

					four2one, //这里的four2one是在上面sel里面选择过后下面会出现的容器,直接在各case里面赋值就可以
				}
				manualContentBox.Refresh()

			})

			sel.PlaceHolder = "请选择一种题型"

			manualContent := container.NewVBox( //手动模式
				canshu,
				xieru,
				command,
				automethodradio,
				phpVersion,
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
				URL := URLentry.Text
				huixianlabel.SetText("访问URL为: " + URL) // 更新回显标签的文本

				if rrraaadio == "自动模式" {

					jianjie := `自动模式
请求模式为` + getpost1 + "\n访问URL为: " + URL + "\nphp版本为" + Version

					if canshu.Text == "" {
						huixianlabel.SetText(jianjie + "\n请输入全部参数")
					} else {
						cleanText, guolv := script.Test(URL, Version, command.Text, canshu.Text, getpost1)
						huixianlabel.SetText(jianjie + "题目源码为:\n" + cleanText + "检测到过滤为:\n" + guolv)

						script.Damn(URL, Version, command.Text, canshu.Text, getpost1, cleanText, guolv, newlabel)
					}

				} else if rrraaadio == "手动模式" {

					jianjie := `手动模式
请求模式为` + getpost1 + "\n访问URL为: " + URL + "\nphp版本为" + Version

					guolv := guolventry.Text

					huixianlabel.SetText(jianjie + "过滤为:\n" + guolv)

					if !strings.Contains(guolv, "$") {

						switch exploit {

						//前三个是eval

						case "zizeng":

							script.Zizeng(URL, command.Text, canshu.Text, getpost1, newlabel)

						case "xor":

							script.Xor(URL, command.Text, Version, canshu.Text, newlabel)

						case "ff":

							script.Ff(URL, command.Text, guolv, canshu.Text, getpost1, newlabel)

						//后三个是system

						case "bashfuck":

							booL := "0"
							script.Bashfuck(URL, command.Text, guolv, canshu.Text, getpost1, newlabel, booL)

						}

					} else {

						//这是ban了'$'的

						switch exploit {

						case "or":

							script.Or(URL, command.Text, canshu.Text, newlabel, guolv, getpost1)

						case "进阶":

							script.Noshuzievaljinjie(URL, command.Text, Version, canshu.Text, getpost1, guolv, newlabel)

						case "进阶plus":

							script.Xorplus(URL, canshu.Text, getpost1, choose, newlabel)

						}
					}

					switch exploit {

					case "black":
						//黑名单
						com, char := script.Blacktest(guolv, guolvleixing)
						//com是所有可用的命令,char是所有可用的字符

						if blackleixing == "chuan" {

							result := script.Blackchuan(guolv, command.Text, com, char)
							newlabel.SetText(result)

						} else if blackleixing == "guding" {

							result := script.Blackguding(URLentry.Text, canshu.Text, getpost1, guolvleixing, com, char, guolv)
							newlabel.SetText(result)

						} else if blackleixing == "shoudong" {

							result := script.Blackshoudong(com, char)
							newlabel.SetText(result + `
下面是这些符号的各自用法(命令我就不写了你自己查一下吧):

$ - 标示变量、命令替换和参数 例如: echo $HOME  # 输出 HOME 变量的值  或者  echo $(date)  # 命令替换，输出当前日期
* - 通配符,例如f*会匹配一切f开头的文件
? - 通配符,例如f???可匹配f开头四个字的文件,例如flag
&& - 逻辑与，前一命令成功时执行下一命令
|| - 逻辑或，前一命令失败时执行下一命令
; - 命令分隔符
| - 管道，将前一个命令的输出作为下一个命令的输入
> - 重定向输出到文件（覆盖）
>> - 重定向输出到文件（追加)
[] - 字符类，匹配括号内的任意一个字符 例如: ls file[1-3].txt  # 匹配 file1.txt, file2.txt, file3.txt
{} - 扩展符，用于生成多个字符串 例如: echo {A,B,C}  # 输出 A B C  或者  echo {1..3}   # 输出 1 2 3
() - 命令组，创建子 Shell 执行命令
$() - 命令替换
$[] - 算术扩展  例如: echo $((1 + 2))
$0 - 当前脚本名
%0a,%09,{$IFS},$IFS$9,%0d - 可替代空格
反引号 - 可替代命令
`)

						}

					case "pwd":

						result := script.Pwd(URLentry.Text, canshu.Text, getpost1, command.Text)
						newlabel.SetText(result)

					case "Fewchar":
						result := script.Fewchar()
						newlabel.SetText(result)

					case "replace":
						result := script.Replace(URL, canshu.Text, getpost1, command.Text)
						newlabel.SetText(result)
					case "file_put":
						result := script.Weixieyi(URL, guolv, canshu.Text, getpost1, command.Text, RW, xieru.Text)
						newlabel.SetText(result)
					}

				}

			})

			mainContent := container.NewVBox(
				URLentry,
				modeRadio,
				phpContent,
			)

			mainscroll := container.NewScroll(mainContent) //给主部分加滚动
			huixianscroll := container.NewScroll(huixian)  //给回显区加个滚动

			//mainsplit是中间容器上下分割
			mainsplit := container.NewVSplit(mainscroll, huixianscroll)
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

	fontSizeLabel := widget.NewLabel("字体大小: " + fmt.Sprintf("%.0f", customTheme.textSize)) //显示当前字体大小

	themes := container.NewGridWithColumns(2, //设置两个按钮

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
