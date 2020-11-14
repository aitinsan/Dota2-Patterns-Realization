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

		fmt.Println(creep.NameArmy + " " + creep.Group + " go to line\n")
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
	direCreepsBlyshniki:=direCreepsFactory("ближний крип",3,10,2,150)
	direCreepsDalniki:=direCreepsFactory("дальний крип",1,12,0,100)

	//radiant
	radiantCreepsFactory:= NewCreepsFactory("Radiant",true)
	radiantCreepsBlyshniki:=radiantCreepsFactory("ближний крип",3,10,2,150)
	radiantCreepsDalniki:=radiantCreepsFactory("дальний крип",1,12,0,100)
	//neutrals
	neutralCreepsFactory1:= NewCreepsFactory("Neutrals",true)
	neutralCreepsWolfs:=neutralCreepsFactory1("волк",2,8,2,150)
	neutralCreepsVozhak:=neutralCreepsFactory1("вожак",1,13,3,200)

	start := GameStartFacade()
	start.CreepsActivity(direCreepsBlyshniki)
	start.CreepsActivity(direCreepsDalniki)
	start.CreepsActivity(radiantCreepsBlyshniki)
	start.CreepsActivity(radiantCreepsDalniki)
	start.CreepsActivity(neutralCreepsWolfs)
	start.CreepsActivity(neutralCreepsVozhak)
	start.BountyActivity()
}
