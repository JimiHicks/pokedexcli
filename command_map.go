package main

import (
	"fmt"

	pokedexapi "github.com/Jimihicks/pokedexcli/internal"
)

func commandMap(c *config) error {
	resp := &pokedexapi.GetLocationAreasResponse{}
	var err error
	if !c.initialized {
		c.initialized = true
		resp, err = pokedexapi.GetLocationAreas()
	} else if c.next != nil {
		resp, err = pokedexapi.GetLocationAreasFromURL(*c.next)
	} else {
		return fmt.Errorf("no next page available")
	}
	if err != nil {
		return err
	}
	c.next = resp.Next
	c.previous = resp.Previous
	for _, result := range resp.Results {
		fmt.Println(result.Name)
	}
	return nil
}
