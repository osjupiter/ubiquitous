package main

type TurnContent struct {
	BuyEquip func()
	Battle func()
	Contest func()
	Reward func()
}
func (t *TurnContent)doTurn(){
	t.BuyEquip()
	t.Battle()
	t.Contest()
	t.Reward()
}

type Flow struct{
	content TurnContent
	turns int
}


func (f *Flow)do() {
	for i := 0; i < f.turns; i++ {
		f.content.doTurn()
	}
}


type Player struct {
	ownCards []Card
}
type Tagmap map[Tag]int
func (p *Player) calcTag() Tagmap{
	tagmap:=make(map[Tag]int)

	for _,v:=range p.ownCards{
		for _,t:=range v.tags{
			if _,ok:=tagmap[t];!ok{
				tagmap[t]=1
			}else{
				tagmap[t]++
			}
		}
	}
	return tagmap
}
