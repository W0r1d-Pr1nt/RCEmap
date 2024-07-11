package script

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var cleantext string

func CleanText(url string) string {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	cleantext = RemoveHTMLTags(string(body))
	return cleantext
}

func Fuzz(URL string, m string, i string) string {
	result := ""
	fmt.Println("即将进行fuzz测试")
	client := &http.Client{}

	var req *http.Request
	var err error

	for ch := 32; ch <= 126; ch++ {
		cha := fmt.Sprintf("%c", ch)

		//char = url.QueryEscape(char)
		newURL := fmt.Sprintf("%s?%s=%s", URL, i, cha)
		cleanText := CleanText(URL)
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
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)

		}

		res := RemoveHTMLTags(string(body))
		// 比较回显与题目，并输出差异
		cleanResponse := strings.Replace(res, cleanText, "", -1)
		cleanText = strings.TrimSpace(cleanText)
		cleanResponse = strings.TrimSpace(cleanResponse)

		if cleanResponse == "" {
			//fmt.Printf(inf+"该字符可用: %s\n", char)
			result += cha
			continue
		} else {
			//fmt.Printf(WARNING+"该字符不可用: %s\n", char)
			continue
		}

	}
	return result
}

func Fuzz1(url string, i string, er string, m string) string {

	//对所有数字字母和符号进行遍历

	result := ""
	fmt.Println("即将进行fuzz测试")
	client := &http.Client{}

	var req *http.Request
	var err error

	cleanText := CleanText(url)
	for ch := 32; ch <= 126; ch++ {
		cha := fmt.Sprintf("%c", ch)

		//char = url.QueryEscape(char)
		newURL := fmt.Sprintf("%s?%s=%s", url, i, cha)

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
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)

		}

		res := RemoveHTMLTags(string(body))
		// 比较回显与题目，并输出差异
		cleanResponse := strings.Replace(res, cleanText, "", -1)
		cleanText = strings.TrimSpace(cleanText)
		cleanResponse = strings.TrimSpace(cleanResponse)

		if cleanResponse != er {
			//fmt.Printf(inf+"该字符可用: %s\n", char)
			result += cha
			continue
		} else {
			//fmt.Printf(WARNING+"该字符不可用: %s\n", char)
			continue
		}

	}
	return result
}

func Fuzz2(url string, i string, er string, m string) string {

	//对bash命令遍历

	client := &http.Client{}
	response := fmt.Sprintf("%s?%s", url, i)
	req, err := http.NewRequest("GET", response, nil)
	if err != nil {
		fmt.Println("Error creating GET request:", err)

	}

	// 发送请求
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

	cleanText := RemoveHTMLTags(string(body))
	result := ""

	allbash := []string{"cat", "tac", "nl", "more", "wget", "grep", "tail", "flag", "less", "head", "sed", "cut", "awk", "strings", "od", "curl", "scp", "rm", "xxd", "mv", "cp", "pwd", "ls", "echo", "sed", "sort"}
	for _, keyword := range allbash {

		//char = url.QueryEscape(char)
		newURL := fmt.Sprintf("%s?%s=%s", url, i, keyword)

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
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)

		}

		res := RemoveHTMLTags(string(body))
		// 比较回显与题目，并输出差异
		cleanResponse := strings.Replace(res, cleanText, "", -1)
		cleanText = strings.TrimSpace(cleanText)
		cleanResponse = strings.TrimSpace(cleanResponse)
		if cleanResponse != er {
			//该字符可用
			result += keyword
			result = result + " "
			continue
		} else {
			//该字符不可用
			continue
		}

	}

	return result
}

func Fuzzpro(url string, i string, m string) string {

	if m == "" {
		fmt.Println("[???]请输入请求方式")

	}

	fmt.Printf(
		"\n[INFO]这是一个下午赶出来的能够对所有字符和所有命令进行fuzz的脚本\n" +
			"[INFO]目前只支持GET方式,POST方式会后续更新,可能还会更新支持字典\n" +
			"[+]请选择执行类型\n" +
			"[1]对所有数字字母和符号进行遍历\n" +
			"[2]对bash命令遍历\n\n" +
			"[-]请选择: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "1" {

		fmt.Println("[+]请输入错误回显: ")
		readerer := bufio.NewReader(os.Stdin)
		er, _ := readerer.ReadString('\n')
		er = strings.TrimSpace(er)

		if url == "" || i == "" || er == "" {
			fmt.Println("[?]请输入全部参数")

		}

		fuzz1 := Fuzz1(url, i, er, m)
		fmt.Println("[-]可用字符有：" + fuzz1)
		return fuzz1
	} else if input == "2" {

		fmt.Println("[+]请输入错误回显: ")
		readerer := bufio.NewReader(os.Stdin)
		er, _ := readerer.ReadString('\n')
		er = strings.TrimSpace(er)

		if url == "" || i == "" || er == "" {
			fmt.Println("[?]请输入全部参数")

		}

		fuzz2 := Fuzz2(url, i, er, m)
		fmt.Println("[+]可用bash命令有: " + fuzz2)
		return fuzz2

	} else if input == "3" {

		// 处理无效选项的情况
		fmt.Println("[?]咱暂时还没这功能")

	} else {

		fmt.Println("[?]你输入了个嘛玩意")

	}
	return ""
}
