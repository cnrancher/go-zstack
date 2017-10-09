package instance

import (
	"os"
	"testing"
	"time"
)

func Test_createInstance(t *testing.T) {
	account := os.Getenv("ACCOUNT")
	password := os.Getenv("PASSWORD")
	zstackEndpoint := os.Getenv("ENDPOINT")
	if zstackEndpoint == "" || account == "" || password == "" {
		t.Skip("Environment not set.")
	}
	c := &Client{}
	if err := c.Init(account, password, zstackEndpoint); err != nil {
		t.Fatal(err)
	}
	//going to create instance
	request := makeCreateRequest()
	resp, err := c.CreateInstance(request)
	if err != nil {
		t.Fatal(err)
	}
	instanceCreateResponse := Response{}
	if err = resp.QueryRealResponse(&instanceCreateResponse, 60*time.Second); err != nil {
		t.Fatal(err)
	}
	if instanceCreateResponse.Error != nil {
		t.Log("Create instance success.")
	}
	//going to delete instance
	uuid := instanceCreateResponse.Inventory.UUID
	if resp, err = c.DeleteInstance(uuid); err != nil {
		t.Log("Delete instance fail.")
		t.Fatal(err)
	}
	//going to expunge instance
	instanceDeleteResponse := DeleteInstanceResponse{}
	if err = resp.QueryRealResponse(&instanceDeleteResponse, 60*time.Second); err != nil {
		t.Fatal(err)
	}
}

func makeCreateRequest() CreateRequest {
	req := CreateRequest{}
	req.Params.Name = "vm99"
	req.Params.InstanceOfferingUUID = "2039ebb02b7a413ba440ac994b495f31"
	req.Params.ImageUUID = "34d1cf5396bd47c994eb3458b58a3f71"
	req.Params.L3NetworkUuids = []string{
		"8be0e4a12161498c8b9aa13c9e84134c",
	}
	req.Params.ZoneUUID = "8a6200a46bda4091b0571dd98e31e3be"
	req.Params.ClusterUUID = "dee0d44ab5204185a84c7b827845fcd2"
	return req
}
