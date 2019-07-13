package main

import "testing"

func TestMainFlow(t *testing.T) {

	e := 0
	b := 0
	c := 0
	r := 0

	f := Flow{
		content: TurnContent{
			BuyEquip: func() {
				e++
			},
			Battle: func() {
				b++
			},
			Contest: func() {
				c++
			},
			Reward: func() {
				r++
			},
		},
		turns: 7,
	}
	f.do()
	if e != 7 || b != 7 || c != 7 || r != 7 {
		t.Errorf("回る回数がおかしい")
	}

}
func TestDeck(t *testing.T) {
	d := Deck{
		pool: []Card{
			{name: "a", cardType: Accessory, tags: []Tag{Magic2}},
			{name: "b", cardType: Consume, tags: []Tag{}},
			{name: "c", cardType: Weapon, tags: []Tag{PlusPower2}},
			{name: "d", cardType: Wearable, tags: []Tag{Magic2}},
			{name: "e", cardType: Wearable, tags: []Tag{Magic2}},
		},
	}
	r:=d.draw(3)
	if r[0].name!="a" ||r[1].name!="b"||r[2].name!="c"{
		t.Errorf("おかしい")
	}
}



