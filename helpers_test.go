package teleBot

import (
	"testing"
)

func TestIsValidPointer(t *testing.T) {
	a := 0
	if IsValidPointer(a) {
		t.Error("failed pointer check for ", a)
	}
	if !IsValidPointer(&a) {
		t.Error("failed pointer check for ", &a)
	}
	if IsValidPointer(nil) {
		t.Error("failed pointer check for nil")
	}
}

func compareMaps(m1, m2 map[int][]Update) bool {
	for id, updates1 := range m1 {
		updates2, ok := m2[id]
		if !ok {
			return false
		}
		if len(updates1) != len(updates2) {
			return false
		}
		for inx, upd := range updates1 {
			if upd != updates2[inx] {
				return false
			}
		}
	}
	return true
}

func TestGroupByChatId(t *testing.T) {
	updates := []Update{
		{Message: Message{Chat: Chat{Id: 0}, Text: "0"}},
		{Message: Message{Chat: Chat{Id: 3}, Text: "1"}},
		{Message: Message{Chat: Chat{Id: 0}, Text: "2"}},
		{Message: Message{Chat: Chat{Id: 2}, Text: "3"}},
		{Message: Message{Chat: Chat{Id: 1}, Text: "4"}},
		{Message: Message{Chat: Chat{Id: 3}, Text: "5"}},
		{Message: Message{Chat: Chat{Id: 2}, Text: "6"}},
		{Message: Message{Chat: Chat{Id: 1}, Text: "7"}},
	}
	ans := map[int][]Update{
		0: {{Message: Message{Chat: Chat{Id: 0}, Text: "0"}},
			{Message: Message{Chat: Chat{Id: 0}, Text: "2"}}},
		1: {{Message: Message{Chat: Chat{Id: 1}, Text: "4"}},
			{Message: Message{Chat: Chat{Id: 1}, Text: "7"}}},
		2: {{Message: Message{Chat: Chat{Id: 2}, Text: "3"}},
			{Message: Message{Chat: Chat{Id: 2}, Text: "6"}}},
		3: {{Message: Message{Chat: Chat{Id: 3}, Text: "1"}},
			{Message: Message{Chat: Chat{Id: 3}, Text: "5"}}},
	}

	resp := GetUpdatesResponse{Ok: true, Res: updates}
	grouped := GroupByChatId(&resp)
	if !compareMaps(grouped, ans) || !compareMaps(ans, grouped) {
		t.Error("Failed grouping: ", grouped, " != ", ans)
	}
}
