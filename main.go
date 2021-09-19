package main

import (
	"bufio"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

const (
	input_file = "../pkms/STL-evaluation.md"

	output_file = "./out.json"
)

type SyncPair struct {
	Id        string `json:"id,omitempty"`      // 生成的唯一id
	Picture   string `json:"picture,omitempty"` // 图片地址
	Name      string `json:"name,omitempty"`    // 名称
	Source      string `json:"source,omitempty"`    // 来源
	Role      string `json:"role,omitempty"`    // 定位
	Stats     *Attr  `json:"stats,omitempty"`   // 白值
	Skill     *Attr  `json:"skill,omitempty"`   // 技能
	Passive   *Attr  `json:"passive,omitempty"` // 被动技能
	Grid      *Attr  `json:"grid,omitempty"`    // 石盘
	Focus     *Attr  `json:"focus,omitempty"`   // 专精度
	Depend    *Attr  `json:"depend,omitempty"`  // 依赖度
	Common    *Attr  `json:"common,omitempty"`  // 通用性
	Score     string `json:"score,omitempty"`   // 分数
	FinishNum int    `json:"-"`                 // 完成数量，用来标记已经完成的栏目数量
}

// 栏目通用模板，有等级评定和原因分析
type Attr struct {
	Level  string `json:"level,omitempty"`  // 等级评定
	Reason string `json:"reason,omitempty"` // 原因分析
}

func parseSTLMarkdown(filePth string) ([]*SyncPair, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	//解析正则表达式
	// 匹配拍组名称
	remove0 := regexp.MustCompile(`^\s*$`)
	remove1 := regexp.MustCompile(`^\[回到目录]\(#目录\)$`)
	reg1 := regexp.MustCompile(`^#+\s(.+&.+)`)
	reg2 := regexp.MustCompile(`^来源\S(.+)`)
	reg3 := regexp.MustCompile(`^定位\S(.+)`)
	reg4 := regexp.MustCompile(`^#+\s白值\S(.+)`)
	reg5 := regexp.MustCompile(`^#+\s技能组\S(.+)`)
	reg6 := regexp.MustCompile(`^#+\s被动\S(.+)`)
	reg7 := regexp.MustCompile(`^#+\s石盘\S(.+)`)
	reg8 := regexp.MustCompile(`^#+\s专精度\S(.+)`)
	reg9 := regexp.MustCompile(`^#+\s依赖度\S(.+)`)
	reg10 := regexp.MustCompile(`^#+\s通用性\S(.+)`)
	reg11 := regexp.MustCompile(`^#+\s综合得分\S(.+)`)

	var syncPairList []*SyncPair
	var nextSyncPair *SyncPair

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if res := reg1.FindAllStringSubmatch(line, -1); res != nil {
			if nextSyncPair != nil {
				syncPairList = append(syncPairList, nextSyncPair)
			}
			nextSyncPair = &SyncPair{
				Id:   fmt.Sprintf("%x", md5.Sum([]byte(res[0][1])))[0:8],
				Name: res[0][1],
			}
		} else if nextSyncPair != nil {
			if res := remove0.Find([]byte(line)); res != nil {
				continue
			}
			if res := remove1.Find([]byte(line)); res != nil {
				continue
			}
			if res := reg2.FindAllStringSubmatch(line, -1); res != nil {
				nextSyncPair.Source = res[0][1]
				continue
			}
			if res := reg3.FindAllStringSubmatch(line, -1); res != nil {
				nextSyncPair.Role = res[0][1]
				continue
			}
			if res := reg4.FindAllStringSubmatch(line, -1); res != nil {
				nextSyncPair.Stats = &Attr{
					Level: res[0][1],
				}
				nextSyncPair.FinishNum++
				continue
			}
			if res := reg5.FindAllStringSubmatch(line, -1); res != nil {
				nextSyncPair.Skill = &Attr{
					Level: res[0][1],
				}
				nextSyncPair.FinishNum++
				continue
			}
			if res := reg6.FindAllStringSubmatch(line, -1); res != nil {
				nextSyncPair.Passive = &Attr{
					Level: res[0][1],
				}
				nextSyncPair.FinishNum++
				continue
			}
			if res := reg7.FindAllStringSubmatch(line, -1); res != nil {
				nextSyncPair.Grid = &Attr{
					Level: res[0][1],
				}
				nextSyncPair.FinishNum++
				continue
			}
			if res := reg8.FindAllStringSubmatch(line, -1); res != nil {
				nextSyncPair.Focus = &Attr{
					Level: res[0][1],
				}
				nextSyncPair.FinishNum++
				continue
			}
			if res := reg9.FindAllStringSubmatch(line, -1); res != nil {
				nextSyncPair.Depend = &Attr{
					Level: res[0][1],
				}
				nextSyncPair.FinishNum++
				continue
			}
			if res := reg10.FindAllStringSubmatch(line, -1); res != nil {
				nextSyncPair.Common = &Attr{
					Level: res[0][1],
				}
				nextSyncPair.FinishNum++
				continue
			}
			if res := reg11.FindAllStringSubmatch(line, -1); res != nil {
				nextSyncPair.Score = res[0][1]
				continue
			}
			if nextSyncPair.FinishNum == 1 {
				nextSyncPair.Stats.Reason += line + "\n"
				continue
			}
			if nextSyncPair.FinishNum == 2 {
				nextSyncPair.Skill.Reason += line + "\n"
				continue
			}
			if nextSyncPair.FinishNum == 3 {
				nextSyncPair.Passive.Reason += line + "\n"
				continue
			}
			if nextSyncPair.FinishNum == 4 {
				nextSyncPair.Grid.Reason += line + "\n"
				continue
			}
			if nextSyncPair.FinishNum == 5 {
				nextSyncPair.Focus.Reason += line + "\n"
				continue
			}
			if nextSyncPair.FinishNum == 6 {
				nextSyncPair.Depend.Reason += line + "\n"
				continue
			}
			if nextSyncPair.FinishNum == 7 {
				nextSyncPair.Common.Reason += line + "\n"
				continue
			}
		}
	}
	//for _, syncPair := range syncPairList {
	//	fmt.Printf("%v\n", syncPair)
	//	fmt.Printf("%v\n", syncPair.Stats)
	//	fmt.Printf("%v\n", syncPair.Skill)
	//	fmt.Printf("%v\n", syncPair.Passive)
	//	fmt.Printf("%v\n", syncPair.Grid)
	//	fmt.Printf("%v\n", syncPair.Focus)
	//	fmt.Printf("%v\n", syncPair.Depend)
	//	fmt.Printf("%v\n", syncPair.Common)
	//	fmt.Printf("%v\n", syncPair.Score)
	//}

	return syncPairList, nil
}

func main() {
	arr, err := parseSTLMarkdown(input_file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	_res, err := json.Marshal(arr)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	err = ioutil.WriteFile(output_file, _res, 0666)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
}
