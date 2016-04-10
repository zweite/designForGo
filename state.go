package main

import (
	"fmt"
)

// 状态模式
//
// 投票场景
// http://www.cnblogs.com/java-my-life/archive/2012/06/08/2538146.html

type State interface {
	Vote(user string, voteItem string, voteManager *VoteManager)
}

type NormalVoteState struct{}

func (n *NormalVoteState) Vote(user string, voteItem string, voteManager *VoteManager) {
	voteManager.VoteMap[user] = voteItem
	fmt.Println("恭喜投票成功")
	return
}

type RepeatVoteState struct{}

func (r *RepeatVoteState) Vote(user string, voteItem string, voteManager *VoteManager) {
	fmt.Println("请不要重复投票")
}

type SpiteVoteState struct{}

func (s *SpiteVoteState) Vote(user string, voteItem string, voteManager *VoteManager) {
	if _, ok := voteManager.VoteMap[user]; ok {
		delete(voteManager.VoteMap, user)
	}
	fmt.Println("你有恶意刷屏行为，取消投票资格")
}

type BlackVoteState struct{}

func (b *BlackVoteState) Vote(user string, voteItem string, voteManager *VoteManager) {
	fmt.Println("进入黑名单，将禁止登录和使用本系统")
}

type VoteManager struct {
	VoteMap   map[string]string
	voteCount map[string]int
}

func NewVoteManager() *VoteManager {
	return &VoteManager{
		VoteMap:   make(map[string]string),
		voteCount: make(map[string]int),
	}
}

func (v *VoteManager) Vote(user string, voteItem string) {
	count := v.voteCount[user]
	if count < 0 {
		count = 0
	}

	count++
	v.voteCount[user] = count

	var state State
	// 投票规则
	if count == 1 {
		state = new(NormalVoteState)
	} else if count > 1 && count <= 5 {
		state = new(RepeatVoteState)
	} else if count > 5 && count <= 8 {
		state = new(SpiteVoteState)
	} else if count > 8 {
		state = new(BlackVoteState)
	}

	state.Vote(user, voteItem, v)
}

func main() {
	voteManager := NewVoteManager()
	for i := 0; i < 9; i++ {
		voteManager.Vote("user", "voteItem")
	}
}
