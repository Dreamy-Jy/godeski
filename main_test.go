package odeskidb

import "testing"

func TestClient_Get(t *testing.T) {
	c := Client{
		database: map[string]string{
			"hello": "world",
		},
	}

	// test for key not found
	_, err := c.Get("notfound")
	if err == nil || err.Error() != "key not found" {
		t.Error("Expected key not found error")
	}

	// test for key found
	val, err := c.Get("hello")
	if err != nil {
		t.Error("Expected no error")
	}
	if val != "world" {
		t.Error("Expected the correct value")
	}

}

func TestClient_Set(t *testing.T) {
	c := Client{
		database: map[string]string{},
	}

	c.Set("hello", "world")
	if c.database["hello"] != "world" {
		t.Error("Expected the value to be set correctly")
	}
}

func TestClient_Clear(t *testing.T) {
	c := Client{
		database: map[string]string{
			"hello": "world",
		},
	}

	c.Clear("hello")
	if _, ok := c.database["hello"]; ok {
		t.Error("Expected the key to be deleted")
	}
}
