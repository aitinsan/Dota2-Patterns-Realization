package main

import "fmt"
/*
	when game starts creeps: neutrals, dire, radiant
	and bounty rune
	appears
 */
// first look at creepsFactory file
// there look at creeps structure
func(c *Creep) appear(creep *Creep){
	fmt.Println(creep.NameArmy+" "+creep.Group + " appeared")
}
func(c *Creep) creepGo(creep *Creep){
	if creep.NameArmy != "Neutrals" {
		fmt.Println(creep.NameArmy + " " + creep.Group + " go to kill")
	}
}
type BountyRune struct{
	Location string
}
func(b *BountyRune) appear(location string){
	fmt.Println("bounty at " + location+ " appeared")
}
type gameStart struct{
	creeps *Creep
	bounty *BountyRune
}
func GameStartFacade() *gameStart{
	return &gameStart{new(Creep),new(BountyRune)}
}
func (g *gameStart) CreepsActivity(creep *Creep){
	g.creeps.appear(creep)
	g.creeps.creepGo(creep)
}
func (g *gameStart) BountyActivity() {
	g.bounty.appear("dire forest")
	g.bounty.appear("dire triangle")
	g.bounty.appear("radiant forest")
	g.bounty.appear("radiant triangle")
}
func startGame(){
	//dire
	direCreepsFactory:= NewCreepsFactory("Dire",true)
	direCreepsBlyshniki:=direCreepsFactory("ближний",3)
	direCreepsDalniki:=direCreepsFactory("дальний",1)
	//radiant
	radiantCreepsFactory:= NewCreepsFactory("Radiant",true)
	radiantCreepsBlyshniki:=radiantCreepsFactory("ближний",3)
	radiantCreepsDalniki:=radiantCreepsFactory("дальний",1)
	start := GameStartFacade()
	start.CreepsActivity(direCreepsBlyshniki)
	start.CreepsActivity(direCreepsDalniki)
	start.CreepsActivity(radiantCreepsBlyshniki)
	start.CreepsActivity(radiantCreepsDalniki)
	start.BountyActivity()
}
