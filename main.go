package odeskidb

import "errors"

type Client struct {
	database map[string]string
}

func (c *Client) Get(key string) (string, error) {
	val, ok := c.database[key]
	if !ok {
		return "", errors.New("key not found")
	}

	return val, nil
}

func (c *Client) Set(key string, value string) {
	c.database[key] = value
}

func (c *Client) Clear(key string) {
	delete(c.database, key)
}
