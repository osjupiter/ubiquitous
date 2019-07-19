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


func TestStatus(t *testing.T){
	p:=Player{ownCards: []Card{
		{name:"a",cardType:Weapon,tags: []Tag{PlusPower2,Shining}},
		{name:"a",cardType:Wearable,tags: []Tag{Magic2,Shining}},
		{name:"a",cardType:Accessory,tags: []Tag{Shining}},
	}}
	s:=p.calcTag()

	if s[Shining]!=3{
		t.Errorf("おかしい")
	}
	if s[Magic2]!=1{
		t.Errorf("おかしい")
	}
	if s[PlusPower2]!=1{
		t.Errorf("おかしい")
	}

	buffs:=p.calcBuff(s)

	if !contains(buffs,Power2){
		t.Errorf("おかしい")
	}
	if !contains(buffs,M2){
		t.Errorf("おかしい")
	}
	if !contains(buffs,HealOnAtk){
		t.Errorf("おかしい")
	}

}
func contains(s []Buff, e Buff) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}
