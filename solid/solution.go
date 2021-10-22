package main

import "fmt"

type Attacker interface {
	Name() string
}

type Victim interface {
	DealDamage(Attacker, int)
}

type Player struct {
	name string
}

type Monster struct {
	hp int
}

func (p *Player) Name() string {
	return p.name
}

func (p *Player) Attack(v Victim) {
	v.DealDamage(p, 100)
}

func (m *Monster) DealDamage(attacker Attacker, damage int) {
	m.hp -= damage
	if m.hp <= 0 {
		fmt.Println(attacker.Name(), "가 나를 죽였다.")
	}
}

func main() {
	player := Player{"전사"}
	var monster Victim = &Monster{100}

	player.Attack(monster)
}
