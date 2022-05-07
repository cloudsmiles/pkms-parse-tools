package internal

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func DownloadMarkdown(url string, saveFile string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// 保存文件
	file, err := os.Create(saveFile)
	if err != nil {
		return err
	}
	defer file.Close()

	write := bufio.NewWriter(file)
	defer write.Flush()
	fmt.Fprint(write, string(body))

	return nil
}
