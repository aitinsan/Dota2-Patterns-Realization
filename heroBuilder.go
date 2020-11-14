package main

import "fmt"

type hero struct {
	name, mainAttribute        string
	
	skills                     []string
	items                      []string
	backPack                   [2]string
	hiddenItems                [5]string
	farm, support, tank, carry bool
	isRadiant                  string
	Attack					   int
	Defence					   int
	Health					   int
}
type heroBuilder struct {
	hero   hero

}



func (h *heroBuilder) Name(name string) *heroBuilder{
	fmt.Println(name)

	h.hero.name= name
	return h
}
func (h *heroBuilder) IsRadiant(isRadiant string) *heroBuilder{
	fmt.Println(isRadiant)
	h.hero.isRadiant= isRadiant
	return h

}
func (h *heroBuilder) MainAttribute(mainAttribute string) *heroBuilder{
	fmt.Println(mainAttribute)
	h.hero.mainAttribute= mainAttribute
	return h

}
func (h *heroBuilder) Skills(skills []string) *heroBuilder{
	fmt.Println(skills)
	h.hero.skills= skills
	return h

}
func (h *heroBuilder) Items(items []string) *heroBuilder{

	fmt.Println(items)
	h.hero.items= items
	return h

}
func (h *heroBuilder) HiddenItems(hiddenItems [5]string) *heroBuilder{
	fmt.Println(hiddenItems)
	h.hero.hiddenItems= hiddenItems
	return h

}
func (h *heroBuilder) BackPack(backPack [2]string) *heroBuilder{
	fmt.Println(backPack)
	h.hero.backPack= backPack
	return h

}
func (h *heroBuilder) Farm(farm bool) *heroBuilder{
	fmt.Println(farm)
	h.hero.farm= farm
	return h

}
func (h *heroBuilder) Carry(carry bool) *heroBuilder{
	fmt.Println(carry)
	h.hero.carry= carry
	return h

}
func (h *heroBuilder) Support(support bool) *heroBuilder{
	fmt.Println(support)
	h.hero.support = support
	return h

}
func (h *heroBuilder) Tank(tank bool) *heroBuilder{
	fmt.Println(tank)
	h.hero.tank = tank
	return h
}
func (h *heroBuilder) Attack(attack int) *heroBuilder{
	fmt.Println(attack)
	h.hero.Attack = attack
	return h
}
func (h *heroBuilder) Defence(defence int) *heroBuilder{
	fmt.Println(defence)
	fmt.Println("\n")
	h.hero.Defence = defence
	return h
}
func (h *heroBuilder) Health(health int) *heroBuilder{
	fmt.Println(health)
	h.hero.Health = health
	return h
}
func createHero(hero *hero) {

}
type build func(builder *heroBuilder)
func CreateHero(action build)  {
	builder :=heroBuilder{}
	action(&builder)
	createHero(&builder.hero)

}

func axehero(){



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

	CreateHero(action)
}