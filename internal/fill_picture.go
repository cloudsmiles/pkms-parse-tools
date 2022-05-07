package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	replaceString = "xxxx"
	url           = "https://www.hualigs.cn/image/" + replaceString + ".jpg" // 图片链接
)

func FillPicture(syncPairList []*SyncPair, filePath string) {
	fmt.Println("----------------- 图片处理 -----------------")

	var picMap = make(map[string]string, len(syncPairList))
	var appendString string // 因为没有图片，补充到txt的拍组信息

	// 读取txt资源文件
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_RDWR, os.ModeAppend|os.ModePerm)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		tmpArr := strings.Split(line, " ")
		if len(tmpArr) == 0 {
			continue
		}
		picMap[tmpArr[0]] = tmpArr[1]
	}

	for _, syncPair := range syncPairList {
		number := picMap[syncPair.Name]
		if number == replaceString {

		} else if number != "" {
			syncPair.Picture = strings.Replace(url, replaceString, number, 1)
		} else {
			// 不符合的情况进行补充
			appendString += fmt.Sprintf("\n%s %s", syncPair.Name, replaceString)
			fmt.Printf("%v\n", syncPair)
		}
	}

	if appendString != "" {
		write := bufio.NewWriter(file)
		defer write.Flush()
		fmt.Fprint(write, appendString)
	}
}
