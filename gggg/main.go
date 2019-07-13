package main

type Effect struct {
}
type Event struct {
	name    string
	effects []Effect
}
type EventType int

const (
	Rock EventType = iota
	Town
	Valley
	Mountain
	Sky
	SeaSide
)

type Status struct {
	power int
	magic int
	gear  int
}

type Champion struct {
	name   string
	events []EventType
	status Status
}

var events = map[EventType]Event{
	Rock: {
		name:    "荒野",
		effects: []Effect{{}},
	},
	Town: {
		name:    "荒野",
		effects: []Effect{{}},
	},
	Valley: {
		name:    "荒野",
		effects: []Effect{{}},
	},
	Mountain: {
		name:    "荒野",
		effects: []Effect{{}},
	},
}
var champions = []Champion{
	{
		name: "転生努力型",
		status: Status{
			power: 3,
			magic: 10,
			gear:  5,
		},
	},
	{
		name: "悪役令嬢",
		status: Status{
			power: 7,
			magic: 7,
			gear:  7,
		},
	},
	{
		name: "異世界探偵",
		status: Status{
			power: 5,
			magic: 7,
			gear:  10,
		},
	},
	{
		name: "追放系勇者",
		status: Status{
			power: 8,
			magic: 8,
			gear:  3,
		},
	},
	{
		name: "一族の末裔",
		status: Status{
			power: 6,
			magic: 10,
			gear:  1,
		},
	},
	{
		name: "一般人（生き残り）",
		status: Status{
			power: 8,
			magic: 3,
			gear:  8,
		},
	},
	{
		name: "怪物（勇者）",
		status: Status{
			power: 9,
			magic: 9,
			gear:  3,
		},
	},
	{
		name: "雑魚モンスター（無天井）",
		status: Status{
			power: 1,
			magic: 1,
			gear:  0,
		},
	},
}

type Origin int
const(
	NEET Origin=iota
	ELITE
	HARDWOKER
	MOTHER
	PET
	SAMURAI
	VILLAN
)
type Job int
const(
	WIZARD Job =iota
	MONSTER
	TAKER
	EXPERT
	INNOCENCE
	MERCHANT
	HEALER
)
type CardType int
const(
	Weapon CardType =iota
	Wearable
	Accessory
	Consume
)

type Tag int
const(
	PlusPower2 Tag =iota
	PlusPower5
	PlusPower8
	Magic2
	Magic5
	Magic8
	Shining
	Dark
	Kingdom
)

type  Card struct{
	name string
	cardType CardType
	tags []Tag

}

type Deck struct {
	pool []Card
}

func (d *Deck) shuffle(){

}

func (d *Deck) draw(num int)[]Card{
	return d.pool[0:num]
}