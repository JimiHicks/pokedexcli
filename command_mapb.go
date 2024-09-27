package main

import (
	"fmt"

	"github.com/Jimihicks/pokedexcli/internal"
)

func commandMapb(c *config) error {
	resp := &internal.GetLocationAreasResponse{}
	var err error
	if c.previous != nil {
		resp, err = internal.GetLocationAreasFromURL(*c.previous)
	} else {
		return fmt.Errorf("no previous page available")
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
