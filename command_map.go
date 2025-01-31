package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, args ...string) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaPage)
	if err != nil {
		return err
	}
	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf("%s\n", area.Name)
	}

	cfg.nextLocationAreaPage = resp.Next
	cfg.prevLocationAreaPage = resp.Previous
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.prevLocationAreaPage == nil {
		return errors.New("you're on the first page")
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreaPage)
	if err != nil {
		return err
	}
	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf("%s\n", area.Name)
	}

	cfg.nextLocationAreaPage = resp.Next
	cfg.prevLocationAreaPage = resp.Previous
	return nil
}
