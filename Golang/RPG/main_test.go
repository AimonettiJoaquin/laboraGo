package main

import (
	"testing"
)

func TestCreateCreature(t *testing.T) {
	c := &Creatures{}
	c.createCreature("Goblin", small, humanoide, 6, 1, 7, 8, 14, 10, 10, 8, 8)
	if len(c.creatures) != 1 {
		t.Error("Failed to create creature")
	}
	//t.Error("Failed to create creature")
	if c.creatures[0].name != "Goblin" {
		t.Error("Failed to create creature")
	}
}

func TestCreateWeapon(t *testing.T) {
	w := &Weapons{}
	w.createWeapon("Zweihänder", 50, 6, "La Zweihänder es una gran espada a dos manos, usada por infantería medieval para romper formaciones.", 2, 6, false)
	if len(w.weapons) != 1 {
		t.Error("Failed to create weapon")
	}
	if w.weapons[0].name != "Zweihänder" {
		t.Error("Failed to create weapon")
	}
}

func TestFindWeapon(t *testing.T) {
	w := &Weapons{}
	w.createWeapon("Zweihänder", 50, 6, "La Zweihänder es una gran espada a dos manos, usada por infantería medieval para romper formaciones.", 2, 6, false)
	weapon := w.findWeapon("Zweihänder")
	if weapon.name != "Zweihänder" {
		t.Error("Failed to find weapon")
	}
}

func TestEquipAndFindEquipWeapon(t *testing.T) {
	c := &Creatures{}
	c.createCreature("Goblin", small, humanoide, 6, 1, 7, 8, 14, 10, 10, 8, 8)
	w := &Weapons{}
	w.createWeapon("Zweihänder", 50, 6, "La Zweihänder es una gran espada a dos manos, usada por infantería medieval para romper formaciones.", 2, 6, false)
	Zweihänder := w.findWeapon("Zweihänder")
	goblin := c.findCreature("Goblin")
	c.equipWeapon(goblin, Zweihänder)
	equipedWeapon := c.findEquipedWeapon(goblin)
	if equipedWeapon.name != Zweihänder.name {
		t.Error("Failed to equip weapon")
	}
}
