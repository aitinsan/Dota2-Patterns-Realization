package main

import "fmt"

func main(){
	//builder for hero creation
	//axehero()

	//creep factory
	//creeps()
	//startgame with bounty and creeps
	fmt.Println("\ngame started:\n")
	direCreepsFactory:= NewCreepsFactory("Dire",true)
	direCreepsBlyshniki:=direCreepsFactory("ближний",3,10,2,150)
	direCreepsDalniki:=direCreepsFactory("дальний",1,12,0,100)
	//radiant
	radiantCreepsFactory:= NewCreepsFactory("Radiant",true)
	radiantCreepsBlyshniki:=radiantCreepsFactory("ближний",3,10,2,150)
	radiantCreepsDalniki:=radiantCreepsFactory("дальний",1,12,0,100)
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

	//visitor for places in dota
	fmt.Println("\nplaces:\n")
	//create axe
	axe:=heroBuilder{}

	action:=func(axe *heroBuilder) {
		axe.Name("Axe").
			MainAttribute("Strength").
			Tank(true).
			Skills([]string{"Berserker's call", "Battle hunger", "Counter helix", "Culling blade"}).
			Items([]string{"DESOLATOR", "Devine Rapier", "Devine Rapier", "Travel boots", "Devine Rapier", "Devine Rapier"}).
			IsRadiant("Radiant").
			Health(1000).
			Attack(50).
			Defence(10)


	}
	fmt.Println(axe)
	CreateHero(action)
	//create io
	io:=heroBuilder{}

	action2:=func(io *heroBuilder) {
		io.Name("IO").
			MainAttribute("Strength").
			Tank(true).
			Skills([]string{"Tether", "Spirits", "Spirits Movement", "Overcharge"}).
			Items([]string{"DESOLATOR", "Devine Rapier", "Devine Rapier", "Travel boots", "Devine Rapier", "Devine Rapier"}).
			IsRadiant("Dire").
			Health(1000).
			Attack(50).
			Defence(10)

	}

	CreateHero(action2)



	dotaland := new(DotaLand)

	dotaland.Add(&MainShop{"Radiant"})
	dotaland.Add(&HiddenShop{"Radiant"})
	//dotaland.Add(&SideShop{"Radiant"})
	dotaland.Add(&MainShop{"Dire"})
	dotaland.Add(&HiddenShop{"Dire"})
	//dotaland.Add(&SideShop{"Dire"})
	dotaland.Add(&Outpost{"Radiant"})
	dotaland.Add(&Outpost{"Dire"})
	dotaland.Add(&Line{"Radiant","Mid",direCreepsBlyshniki.health*3+direCreepsDalniki.health})
	Axe:=dotaland.Accept(&axe)
	fmt.Println("Axe\n"+Axe)
	Io:=dotaland.Accept(&io)
	fmt.Println("Io\n"+Io)





}