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


}
func NewCreepsFactory(name string, ready bool) func(group string,value int) *Creep{
	return func(group string,value int)*Creep{
		return &Creep{name,group,value,ready}
	}
}
func creeps(){

	direCreepsFactory:= NewCreepsFactory("Dire",true)
	direCreepsBlyshniki:=direCreepsFactory("ближний",3)
	direCreepsDalniki:=direCreepsFactory("дальний",1)
	fmt.Println(direCreepsBlyshniki)
	fmt.Println(direCreepsDalniki)

	radiantCreepsFactory:= NewCreepsFactory("Radiant",true)
	radiantCreepsBlyshniki:=radiantCreepsFactory("ближний",3)
	radiantCreepsDalniki:=radiantCreepsFactory("дальний",1)
	fmt.Println(radiantCreepsBlyshniki)
	fmt.Println(radiantCreepsDalniki)

	neutralCreepsFactory1:= NewCreepsFactory("Neutrals",true)
	neutralCreepsWolfs:=neutralCreepsFactory1("волк",2)
	neutralCreepsVozhak:=neutralCreepsFactory1("вожак",1)
	fmt.Println(neutralCreepsWolfs)
	fmt.Println(neutralCreepsVozhak)



}
