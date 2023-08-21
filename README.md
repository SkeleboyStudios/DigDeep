# DigDeep
A game made using engo for the Brackey's Jam 2023-2

# Building
The game will be available on itch.io by the end of the week for the contest! I'll add the link here when
available!

If you have go on your system, it should simply run with

```
go run .
```

# Gameplay and Features

Your job is to manage a dwarven mining operation! Use the dwarves to collect resources such as iron,
coal and gold. Then use those resources to craft stuff like mine carts, mine tracks, pickaxes, swords,
armor to help your dwarves delve even deeper! Hire more dwarves, fight monsters, and see just how
deep you can go! Watch out for cave monsters like bats, goblins, zombies, and slimes!

## Mine

The mine is going to be randomly generated with hidden cave networks filled with enemies! The player
will use their dwarves to mine deeper and deeper, uncovering resources, caves, and bad guys! The mountain
has a finite size, but the player can start over in a new mountain. This will unlock after a certain point,
and will have benefits for the player as well.

The tileset for the mine is [Goblin Cave](https://opengameart.org/content/goblin-caves) by
[hyptosis](https://opengameart.org/users/hyptosis).

![Goblin Cave Tileset](/assets/goblin_cave.png)

## Dwarves

The player will start with one dwarf, and can hire more at the crafting table for gold. The dwarves are
randomly generated, with different stats: some are better at mining, some at fighting, some at defense,
and some have more stamina.
The names will be randomly generated from a list created by the 
[Random Dwarf Name Generator](https://thestoryshack.com/tools/dwarf-name-generator/).
The sprites will be one random sprite from 
[Dwarves](https://opengameart.org/content/dwarves-0) by 
[Svetlana Kushnariova (Cabbit)](lana-chan@yandex.ru), diamonddmgirl, GrumpyDiamond, & 
[Jordan Irwin (AntumDeluge)](https://opengameart.org/users/antumdeluge)

![Dwarves](/assets/dwarf_preview.png)

## Mining

Mine by selecting a dwarf and then clicking on an ore node or a wall to expand the cave. The ores will 
be coal, iron, gold, and emerald. The sprites for those are the on the same sheet as the rest of the mine.
The dwarf will walk over and mine out the ore, adding it to the dwarf's inventory. The mining speed is
determined by the dwarf's mining skill, which can be enhanced by equipping a pickaxe or leveling up.
The dwarf has to drop off the ore at the crafting table to make it available to the player. Dwarves have
a limited inventory space so have to return at some point.

## Crafting

The player can access the crafting menu by clicking on the crafting table in the main base. They can use that
to create equipment for the dwarves, as well as beds, torches, and candles.

### Pickaxe

The pickaxe improves the dwarf's mining stat. A higher mining stat means the dwarf mines faster.

### Axe

The axe improves the dwarf's fight stat. A higher fight stat means the dwarf deals more damage to enemies.

### Bow

The bow allows the dwarf to shoot distant enemies.

### Armor

Armor improves the dwarf's defense stat.

### Bed

The bed allows the dwarves to rest and recover stamina.

### Backpack

Allows the dwarf to carry more stuff.

### Red Dwarf

An energy potion that restores the dwarf's stamina at the expense of health.

### Torch

A torch to light the way. Can be equipped in a dwarf's hands or placed on a wall to light an area.

### Candle

Similar to a torch, but can only be placed on walls, not carried by a dwarf.

## Enemies

Enemies will populate the underground caves the dwarves can dig into. The caves make it easier to get ore, but
can be dangerous!

### Slimes

Slimes are the easiest enemies. Lowest hp, fight, and def.

### Bats

Bats are faster, but shouldn't be too much of a threat.

### Zombies

Zombies are slow but strong. Deals poisoning damage that could be deadly.

### Goblins

Goblins are versatile and intelligent critters. Some might even be ranged with bows and arrows!

