package main

import (
	"fmt"
	"math/rand"
)

type Guess struct {
}

type GuessInfo struct {
	ID     int    `json:"id"`     // 序号
	Riddle string `json:"riddle"` // 谜题
	Answer string `json:"answer"` // 谜底
}

func (g *Guess) GetRiddle() string {
	randNum := rand.Intn(10)
	list := g.GetGuessRecords()
	RiddleInfo := list[randNum]
	riddleStr := fmt.Sprintf("序号：%d\n%s", RiddleInfo.ID, RiddleInfo.Riddle)
	return riddleStr
}

func (g *Guess) GetAnswerByID(id int) GuessInfo {
	list := g.GetGuessRecords()
	for _, v := range list {
		if v.ID == id {
			return v
		}
	}
	return GuessInfo{Answer: "该序号谜语不存在！"}
}

func (g *Guess) GetGuessRecords() []GuessInfo {
	guessRecord := []GuessInfo{
		GuessInfo{0, "一无所获(打一字)", "控"},
		GuessInfo{1, "人言不作信字猜(打一字)——谜底", "认"},
		GuessInfo{2, "七嘴八舌(打一字)——谜底", "哆"},
		GuessInfo{3, "扫平高家庄(打一字)——谜底", "崽"},
		GuessInfo{4, "轻烟薄雾笼山峰(打一字)", "氙"},
		GuessInfo{5, "仁智之乐(打一字)", "汕"},
		GuessInfo{6, "破镜重圆(打一字)", "锺"},
		GuessInfo{7, "乃正负之比耳(打一字)", "陛"},
		GuessInfo{8, "明月出天上(打一字)", "旦"},
		GuessInfo{9, "个个见了笑(打一字)", "夭"},
	}

	return guessRecord
}
