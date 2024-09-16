package main

import (
	"errors"
	"fmt"
)

const (
	AK47 string = "Avtomat Kalashnikov"
	M4A1 string = "M4A1"
)

var (
	ErrInvalidGunType = errors.New("invalid gun type")
)

type Gun struct {
	name   string
	damage int
}

type IGun interface {
	setName(name string)
	setDamage(damage int)
	getName() string
	getDamage() int
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) setDamage(damage int) {
	g.damage = damage
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) getDamage() int {
	return g.damage
}

type Ak47 struct {
	Gun
}

func newAk47() IGun {
	return &Ak47{
		Gun: Gun{
			name:   AK47,
			damage: 120,
		},
	}
}

type M4 struct {
	Gun
}

func newM4() IGun {
	return &M4{
		Gun: Gun{
			name:   M4A1,
			damage: 100,
		},
	}
}

func getGun(gunType string) (IGun, error) {

	if gunType == AK47 {
		return newAk47(), nil
	}

	if gunType == M4A1 {
		return newM4(), nil
	}

	return nil, ErrInvalidGunType

}

func printGun(g IGun) {
	fmt.Printf("GUN TYPE : %s\n", g.getName())
	fmt.Printf("GUN DAMAGE : %d\n", g.getDamage())
}

func main() {

	ak47, _ := getGun(AK47)
	m4, _ := getGun(M4A1)

	printGun(ak47)
	printGun(m4)

}
