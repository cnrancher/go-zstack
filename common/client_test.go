package common

import "testing"
import "fmt"
import "os"

func Test_client(t *testing.T) {
	account := os.Getenv("ACCOUNT")
	password := os.Getenv("PASSWORD")
	zstackEndpoint := os.Getenv("ENDPOINT")
	if zstackEndpoint == "" || account == "" || password == "" {
		t.Skip("Environment not set")
	}
	c := &Client{}
	if err := c.Init(account, password, zstackEndpoint); err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v\n", *c)
	if err := c.Cleanup(); err != nil {
		t.Fatal(err)
	}
}
