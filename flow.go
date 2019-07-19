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
func (p *Player) calcBuff(tagmap Tagmap) []Buff{
	ret:=[]Buff{}
	if tagmap[PlusPower2]>0{
		ret=append(ret,Power2)
	}
	if tagmap[PlusPower5]>0{
		ret=append(ret,Power5)
	}
	if tagmap[PlusPower5]>0{
		ret=append(ret,Power5)
	}
	if tagmap[PlusPower8]>0{
		ret=append(ret,Power8)
	}
	if tagmap[Magic2]>0{
		ret=append(ret,M2)
	}
	if tagmap[Magic5]>0{
		ret=append(ret,M5)
	}
	if tagmap[Magic8]>0{
		ret=append(ret,M8)
	}
	if tagmap[Shining]>2{
		ret=append(ret,HealOnAtk)
	}
	if tagmap[Dark]>2{
		ret=append(ret,DualAtk)
	}
	if tagmap[Kingdom]>1{
		ret=append(ret,Golden)
	}
	return ret
}
