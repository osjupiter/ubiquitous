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
