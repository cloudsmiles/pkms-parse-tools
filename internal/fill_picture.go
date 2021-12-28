package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func FillPicture(syncPairList []*SyncPair, filePath string) {
	fmt.Println("----------------- 图片处理 -----------------")

	const url = "https://www.hualigs.cn/image/xxxx.jpg" // 图片链接
	var picMap = make(map[string]string, len(syncPairList))

	// 读取资源文件
	f, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
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
		if number != "" && number != "xxxx" {
			syncPair.Picture = strings.Replace(url, "xxxx", number, 1)
		} else {
			fmt.Printf("%v\n", syncPair)
		}
	}
}
