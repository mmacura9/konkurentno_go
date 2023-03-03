package main

type Person struct {
	id                string
	name              string
	birthYear         string
	deathYear         string
	primaryProfession string
	knownForTitles    string
}

func new_Person(id string, name string, birthYear string, deathYear string, primaryProfession string, knownForTitles string) Person {
	return Person{id: id, name: name, birthYear: birthYear, deathYear: deathYear, primaryProfession: primaryProfession, knownForTitles: knownForTitles}
}
