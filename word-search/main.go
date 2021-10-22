package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type LineInfo struct {
	lineNo int
	line   string
}

type FindInfo struct {
	fileName  string
	lineInfos []LineInfo
}

func main() {

	if len(os.Args) < 3 {
		fmt.Println("2개 이상의 인수가 필요합니다. ex) 단어 파일경로")
		return
	}

	word := os.Args[1]
	files := os.Args[2:]
	findInfos := []FindInfo{}

	for _, path := range files {
		findInfos = append(findInfos, FindWordInAllFiles(word, path)...)
	}

	for _, findInfo := range findInfos {
		PrintFindInfo(findInfo)
	}

}

func GetFileList(path string) ([]string, error) {
	return filepath.Glob(path)
}

func FindWordInAllFiles(word, path string) []FindInfo {
	findInfos := []FindInfo{}

	fileList, err := GetFileList(path)
	if err != nil {
		fmt.Println("파일 경로가 잘못되었습니다. err:", err, "path:", path)
		return findInfos
	}

	channel := make(chan FindInfo)
	fileCount := len(fileList)
	resultReceiveCount := 0

	for _, fileName := range fileList {
		go FindWordInFile(word, fileName, channel)
	}

	for findInfo := range channel {
		findInfos = append(findInfos, findInfo)
		resultReceiveCount++

		if resultReceiveCount == fileCount {
			break
		}
	}

	return findInfos
}

func FindWordInFile(word, fileName string, channel chan FindInfo) {
	findInfo := FindInfo{fileName, []LineInfo{}}
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다. fileName:", fileName)
		channel <- findInfo
		return
	}
	defer file.Close()

	FindWord(file, word, &findInfo)

	channel <- findInfo
}

func FindWord(file *os.File, word string, findInfo *FindInfo) {
	lineNo := 1
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, word) {
			findInfo.lineInfos = append(findInfo.lineInfos, LineInfo{lineNo, line})
		}
		lineNo++
	}
}

func PrintFindInfo(findInfo FindInfo) {
	fmt.Println(findInfo.fileName)
	fmt.Println("=====================")
	for _, lineInfo := range findInfo.lineInfos {
		fmt.Printf("\t%d\t%s\n", lineInfo.lineNo, lineInfo.line)
	}
	fmt.Println("=====================")
	fmt.Println()
}
