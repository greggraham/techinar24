package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

// Hold the information for a particular Pokemon
type Pokemon struct {
	Id        int
	Name      string
	Type      []string
	HP        int
	Attack    int
	Defense   int
	SpAttack  int
	SpDefense int
	Speed     int
}

// Read in a JSON Pokemon database and let the user query it.
func main() {
	pArray := readPokemonDb("./pokedex.json")
	pMap := makePokemonMap(pArray)
	query(pMap)
}

// Read in the Pokemon database from a JSON file into an array.
func readPokemonDb(fileName string) []Pokemon {
	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// Unmarshall the data
	var pArray []Pokemon
	err = json.Unmarshal(content, &pArray)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	return pArray
}

// Organize an array of Pokemon entries into a Map keyed by the Pokemon name.
func makePokemonMap(pArray []Pokemon) map[string]Pokemon {
	pMap := make(map[string]Pokemon, 900)
	for _, p := range pArray {
		key := strings.ToLower(p.Name)
		pMap[key] = p
	}
	return pMap
}

// Repeatedly ask the user for a Pokemon name and print out the entry that corresponds to it.
func query(pMap map[string]Pokemon) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the name of a pokemon to look up information about it. Enter 'q' to quit.")
	for {
		fmt.Print("\n> ")
		name := getString(reader)
		if name == "q" {
			break
		}
		key := strings.ToLower(name)
		entry, ok := pMap[key]
		if !ok {
			fmt.Printf("[%s] is not in the database. Please try again", name)
		} else {
			printEntry(entry)
		}
	}
}

// Neatly read in a string from the console
func getString(reader *bufio.Reader) string {
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	name = strings.TrimSpace(name)
	return name
}

// Nicely print out a Pokemon entry.
func printEntry(e Pokemon) {
	fmt.Printf("%s (%d) - ", e.Name, e.Id)
	types := strings.Join(e.Type, ", ")
	fmt.Println(types)
	fmt.Println("HP        :", e.HP)
	fmt.Println("Attack    :", e.Attack)
	fmt.Println("Defense   :", e.Defense)
	fmt.Println("SpAttack  :", e.SpAttack)
	fmt.Println("SpDefense :", e.SpDefense)
	fmt.Println("Speed     :", e.Speed)
}

// Reimplement the String method for the Pokemon type to use a more readable format.
func (p Pokemon) String() string {
	return fmt.Sprintf("%d: %s %v HP:%d Att:%d Def:%d SpAtt:%d SpDef:%d Speed:%d\n",
		p.Id, p.Name, p.Type, p.HP, p.Attack, p.Defense, p.SpAttack, p.SpDefense, p.Speed)
}
