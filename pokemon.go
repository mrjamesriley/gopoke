// Usage: go run pokemon.go
package main

import (
  "fmt"
  "strconv"
  "os"
  "os/exec"
)

// structs are the closest Go gets to a Class, grouping data to form records
type player struct {
  name string
  pokemon pokemon
}

type pokemon struct {
  name string
  strength int
  hp int
  fainted bool
}

func (p *pokemon) attack(enemy *pokemon) {
 fmt.Println(p.name, "did", strconv.Itoa(p.strength), "HP of damage")
 enemy.hp = enemy.hp - p.strength
 fmt.Println(enemy.name, "now has", enemy.hp, "HP left")
 if(enemy.hp <= 0) {
   enemy.hp = 0
   enemy.fainted = true
   fmt.Println(enemy.name, "has fainted!")
 }
}

type pokeGym struct {
  area string
  difficulty string
  ticketsSold int
  yearBuilt int
}

// define a function, setting a receiver which effectively adds a method to a struct
func (p *pokeGym) yearsOld() int {
  return 2015 - p.yearBuilt
}

func pokemonIsAvailable(selection string, pokemon []pokemon) bool {
  available := false
  for i := 0; i < 3; i++ {
    if(pokemon[i].name == selection) { available = true }
  }
  return available
}

func pokemonByName(pokemonName string, pokemonSet []pokemon) pokemon {
  var chosenPokemon pokemon
  for i := 0; i < 3; i++ {
    if(pokemonSet[i].name == pokemonName) { chosenPokemon = pokemonSet[i] }
  }
  return chosenPokemon
}

func selectPokemon(selection pokemon, player *player) {
  player.pokemon = selection
  fmt.Println("\n" + player.name, "selected", player.pokemon.name + "!")
  fmt.Println(player.pokemon.name, "has", strconv.Itoa(player.pokemon.hp),
    "HP and a Strength of", player.pokemon.strength)
}

func clearScreen() {
  c := exec.Command("clear")
  c.Stdout = os.Stdout
  c.Run()
}

func main() {

  // maps are Go's associative data type (i.e hash)
  playerOne := player{ name: "James" }
  playerTwo := player{ name: "Gary" }

  charmander := pokemon{name: "Charmander", strength: 4, hp: 20, fainted: false }
  squirtle   := pokemon{name: "Squirtle", strength: 10, hp: 20, fainted: false }
  bulbasaur  := pokemon{name: "Bulbasaur", strength: 7, hp: 20, fainted: false }

  pokedex := make([]pokemon, 3)
  pokedex[0] = charmander
  pokedex[1] = squirtle
  pokedex[2] = bulbasaur

  battleLocation := pokeGym{area: "Lavender Town", difficulty: "Easy", yearBuilt: 1940 }

  fmt.Println("\nWelcome to Kanto,", playerOne.name);
  fmt.Println("You are just in time for the battle with your rival", playerTwo.name);
  fmt.Println("The battle will take place at the:", battleLocation.area, "gym\n")
  fmt.Println("The Gym is", battleLocation.yearsOld(), "years old");

  fmt.Println("You have the choice of 3 Pokemon!\n")
  for i := 0; i < 3; i++ { fmt.Println(pokedex[i].name) }
  fmt.Println("\n")

  // get player to select their starting pokemon
  selectingPokemon := true
  for selectingPokemon {

    var usersChoice string
    fmt.Println("Which starting Pokemon would you like? \n")
    _, err := fmt.Scanln(&usersChoice)
    if(err != nil) { fmt.Println("Err has the value: ", err) }

    if(pokemonIsAvailable(usersChoice, pokedex)) {
      selectPokemon(pokemonByName(usersChoice, pokedex), &playerOne)
      selectingPokemon = false
    } else {
      fmt.Println("Sorry,", usersChoice, "is not available\n");
    }
  }

  // opponents picks the first available weakest pokemon
  if(playerOne.pokemon.name == "Charmander") {
    selectPokemon(pokemonByName("Bulbasaur", pokedex), &playerTwo)
  } else {
    selectPokemon(pokemonByName("Charmander", pokedex), &playerTwo)
  }



  // the battle - repeatedly take turns till either pokemon faints
  // player gets options, rival always attacks
  fmt.Println("\nLet battle commence!")

  battleOngoing := true
  turnOngoing := true
  currentPlayer := playerOne

  for battleOngoing {
    fmt.Println("\nIt is now", currentPlayer.name + "'s turn!\n")

    if(currentPlayer.name == playerOne.name) {

      turnOngoing = true

      for turnOngoing {
        fmt.Println("What would you like to do?")
        fmt.Println("1) Attack\n2) Use Item\n3) Run\n")
        var turnChoice string
        _, err := fmt.Scanln(&turnChoice)
        clearScreen()

        if(err != nil) {
          fmt.Println("Oh no, there was an error!")
          fmt.Println(err)
        } else {
          if(turnChoice == "1") {
            fmt.Println("Your decided to attack!")
            playerOne.pokemon.attack(&playerTwo.pokemon)
          } else if(turnChoice == "2") {
            fmt.Println("You have no items!")
          } else if(turnChoice == "3") {
            fmt.Println("You cant run from a fight with your rival!")
          }

          if(turnChoice == "1") {
            turnOngoing = false
          } else {
            fmt.Println("Please choose another action...")
          }
        }
      }

    } else {
      fmt.Println(playerTwo.name, "chose to attack!")
      playerTwo.pokemon.attack(&playerOne.pokemon)
    }

    // if either pokemon has fainted, the game is over
    if(playerTwo.pokemon.fainted || playerOne.pokemon.fainted) {
      fmt.Println("\n\nThe battle is now over...")

      if(playerOne.pokemon.fainted) {
        fmt.Println("You Lose!")
      } else {
        fmt.Println("You Win!")
      }

      battleOngoing = false
    }

    // switch to the other player for their turn
    if(currentPlayer.name == playerOne.name) {
      currentPlayer = playerTwo
    } else {
      currentPlayer = playerOne
    }

  }

}
