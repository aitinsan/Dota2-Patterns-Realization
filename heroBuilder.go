package main

import "fmt"

type hero struct {
	name, mainAttribute string
	skills []string
	items []string
	backPack [2]string
	hiddenItems [5]string
	farm, support, tank,carry bool
}
type heroBuilder struct{
	hero hero
}
func (value *heroBuilder) Name(name string) *heroBuilder{
	fmt.Println(name)

	value.hero.name= name
	return value
}
func (value *heroBuilder) MainAttribute(mainAttribute string) *heroBuilder{
	fmt.Println(mainAttribute)
	value.hero.mainAttribute= mainAttribute
	return value

}
func (value *heroBuilder) Skills(skills []string) *heroBuilder{
	fmt.Println(skills)
	value.hero.skills= skills
	return value

}
func (value *heroBuilder) Items(items []string) *heroBuilder{

	fmt.Println(items)
	value.hero.items= items
	return value

}
func (value *heroBuilder) HiddenItems(hiddenItems [5]string) *heroBuilder{
	fmt.Println(hiddenItems)
	value.hero.hiddenItems= hiddenItems
	return value

}
func (value *heroBuilder) BackPack(backPack [2]string) *heroBuilder{
	fmt.Println(backPack)
	value.hero.backPack= backPack
	return value

}
func (value *heroBuilder) Farm(farm bool) *heroBuilder{
	fmt.Println(farm)
	value.hero.farm= farm
	return value

}
func (value *heroBuilder) Carry(carry bool) *heroBuilder{
	fmt.Println(carry)
	value.hero.carry= carry
	return value

}
func (value *heroBuilder) Support(support bool) *heroBuilder{
	fmt.Println(support)
	value.hero.support = support
	return value

}
func (value *heroBuilder) Tank(tank bool) *heroBuilder{
	fmt.Println(tank)
	value.hero.tank = tank
	return value
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

	CreateHero(func(value *heroBuilder) {
		value.Name("Axe").
			MainAttribute("Strength").
			Tank(true).
			Skills([]string{"Berserker's call", "Battle hunger", "Counter helix", "Culling blade"}).
			Items([]string{"DESOLATOR", "Devine Rapier", "Devine Rapier", "Travel boots", "Devine Rapier"})

	})

}
