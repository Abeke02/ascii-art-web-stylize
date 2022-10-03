package ascii

import (
	"errors"
	"os"
	"strings"
)

func Art(text string, style string) (string, error) {
	text = strings.ReplaceAll(text, "\r\n", "\n")
	arr := strings.Split(text, "\n")

	var slice [][]string
	nameFile := "ascii-art/" + style + ".txt"
	temp, err := os.ReadFile(nameFile)
	if err != nil {
		return "", errors.New("Error")
	}

	template := strings.Split(string(temp), "\n\n")

	n := []string{"n"}
	for j := 0; j < len(arr); j++ {
		for i := 0; i < len(arr[j]); i++ {
			slice = append(slice, Str(template[arr[j][i]-32]))
		}
		slice = append(slice, n)
	}
	res := findNewLine(slice)
	result := ""
	for _, i := range res {
		result += i
	}
	return result, nil
}

func Str(str string) []string {
	temp := strings.Split(string(str), "\n")
	return temp
}

func asciiFinal(slice [][]string) string {
	result := ""
	var arr []string
	nLine := []string{"\n"}
	for j := 0; j < 8; j++ {
		for i := 0; i < len(slice); i++ {
			if slice[i][0] == "" {
				slice[i][0] = "      "
			}
			if slice[i][0] == "n" {
				arr = append(arr, nLine...)
			} else {
				result += slice[i][j]
			}

		}
		result += "\n"
		arr = append(arr, result)
		result = ""
	}

	for _, i := range arr {
		result += i
	}
	return result
}

func findNewLine(slice [][]string) []string {
	var arr [][]string
	var res []string

	for i := 0; i < len(slice); i++ {
		if slice[i][0] == "n" {
			res = append(res, asciiFinal(arr))
			arr = nil
		} else {
			arr = append(arr, slice[i])
		}
	}
	return res
}
