package main

import "fmt"
/*
	creeps created:
		dire creeps:
			Blyshniki
            Dalniki
		radiant creeps:
			Blyshniki
            Dalniki
		neutral creeps:
			волки:
				обычные(волки)
				вожак
 */
type Creep struct{
	NameArmy,Group string
	Value int
	OnWar bool
	Attack int
	defense int
	health int


}
func NewCreepsFactory(name string, ready bool) func(group string,value int,attack int, defence int, health int) *Creep{
	return func(group string,value int,attack int, defence int, health int)*Creep{
		return &Creep{name,group,value,ready, attack, defence, health}
	}
}
func creeps(){
	//dire
	direCreepsFactory:= NewCreepsFactory("Dire",true)
	direCreepsBlyshniki:=direCreepsFactory("ближний",3,10,2,150)
	direCreepsDalniki:=direCreepsFactory("дальний",1,12,0,100)
	fmt.Println(direCreepsBlyshniki)
	fmt.Println(direCreepsDalniki)
	//radiant
	radiantCreepsFactory:= NewCreepsFactory("Radiant",true)
	radiantCreepsBlyshniki:=radiantCreepsFactory("ближний",3,10,2,150)
	radiantCreepsDalniki:=radiantCreepsFactory("дальний",1,12,0,100)
	fmt.Println(radiantCreepsBlyshniki)
	fmt.Println(radiantCreepsDalniki)
	//neutral
	neutralCreepsFactory1:= NewCreepsFactory("Neutrals",true)
	neutralCreepsWolfs:=neutralCreepsFactory1("волк",2,8,2,150)
	neutralCreepsVozhak:=neutralCreepsFactory1("вожак",1,13,3,200)
	fmt.Println(neutralCreepsWolfs)
	fmt.Println(neutralCreepsVozhak)
}
