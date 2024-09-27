package main

import (
	"fmt"

	"github.com/Jimihicks/pokedexcli/internal"
)

func commandMap(c *config) error {
	resp := &internal.GetLocationAreasResponse{}
	var err error
	if c.next != nil {
		resp, err = internal.GetLocationAreasFromURL(*c.next)
	} else {
		resp, err = internal.GetLocationAreas()
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
