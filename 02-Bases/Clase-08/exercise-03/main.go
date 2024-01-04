package main

import (
	"errors"
	"fmt"
)

type Client struct {
	File        string
	Name        string
	ID          int
	PhoneNumber int
	Home        string
}

var Clients = []Client{
	{
		File:        "file-01",
		Name:        "Thiago",
		ID:          1,
		PhoneNumber: 12345678,
		Home:        "False street 123",
	},
}

func (c *Client) validateExistingClient(client Client) error {
	for _, cl := range Clients {
		if cl.ID == client.ID {
			panic("Error: client already exists")
		}
	}
	return nil
}

func (c *Client) validateClientData(client Client) (string, error) {
	for _, c := range Clients {
		if c.ID == 0 || c.Home == "" || c.File == "" || c.Name == "" || c.PhoneNumber == 0 {
			msg := "An error has ocurred"
			err := errors.New("Error: client data cannot be zero or empty")
			return msg, err
		}
	}
	return "", nil
}

func (c *Client) saveClient(client Client) {
	c.validateExistingClient(client)

	c.validateClientData(client)

	Clients = append(Clients, client)
	fmt.Println("Client added")

}

func main() {
	defer func() {
		fmt.Println("End of execution\nSeveral errors were detected at runtime")
	}()
	var newClient = Client{
		File:        "file-01",
		Name:        "Thiago",
		ID:          1,
		PhoneNumber: 12345678,
		Home:        "False street 123",
	}

	newClient.saveClient(newClient)
}
