package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config) error {
	loc := cfg.location
	if loc == nil {
		return errors.New("No location specified")
	}
	locationsResp, err := cfg.pokeapiClient.ExploreLocationArea(loc)
	if err != nil {
		return err
	}

	for _, pok := range locationsResp.PokemonEncounters {
		fmt.Println(pok.Pokemon.Name)
	}
	return nil
}
