package main
import "fmt"

func (c *hero) String() string {
	return fmt.Sprintf("%s (%d/%d)",c.name,c.Attack,c.Defence)
}

// 1 Interface Handler
type Modifier interface {
	Add(m Modifier)
	Apply()
}
// 2 Base Handler
type CreatureModifier struct {
	next Modifier
	creature *hero
}
func (c *CreatureModifier) Add(m Modifier)  {
	if c.next != nil{
		c.next.Add(m)
	} else {
		c.next = m
	}
}
func (c *CreatureModifier) Apply()  {
	if c.next != nil {
		c.next.Apply()
	}
}
func NewCreatureModifier(c *hero) *CreatureModifier {
	return &CreatureModifier{creature: c}
}
// 3 Concrete Handler
type DoubleDamageModifier struct {
	CreatureModifier
}
func (d *DoubleDamageModifier) Apply() {
	fmt.Println("double damage rune applied for",d.creature.name)
	d.creature.Attack *=2
	//!!!
	d.CreatureModifier.Apply()
}

func NewDoubleDamageModifier(c *hero) *DoubleDamageModifier {
	return &DoubleDamageModifier{CreatureModifier{creature: c}}
}

type FrostShieldModifier struct {
	CreatureModifier
}
func (d *FrostShieldModifier) Apply() {
	if d.creature.Attack <=150 {
		fmt.Println("Frost shield was applied")
		d.creature.Defence +=20
	}
	d.CreatureModifier.Apply()
}
func NewFrostShieldModifier(c *hero) *FrostShieldModifier {
	return &FrostShieldModifier{CreatureModifier{creature: c}}
}

func jugger()  {
	//create juggernaut


	action3:=func(jugger *heroBuilder) {
		jugger.Name("juggernaut").
			MainAttribute("Agility").
			Tank(true).
			Skills([]string{"Blade Fury", "Healing Ward", "Blade Dance", "Omnislash"}).
			IsRadiant("Dire").
			Health(1000).
			Attack(50).
			Defence(10)
		Juggernaut=&jugger.hero

	}

	CreateHero(action3)

	fmt.Println(Juggernaut)
	root := NewCreatureModifier(Juggernaut)
	root.Add(NewDoubleDamageModifier(Juggernaut))
	root.Add(NewFrostShieldModifier(Juggernaut))
	root.Add(NewDoubleDamageModifier(Juggernaut))
	root.Add(NewFrostShieldModifier(Juggernaut))
	root.Apply()
	fmt.Println(Juggernaut)
}