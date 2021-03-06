package main

import (
	"fmt"
	"testing"
)

func TestRegister(t *testing.T) {
	// client := resty.New()

	// TODO(ralphg6) mock the client

	// resp, err := register(client, "http://fakeaddress", "fakeaccount")

	// if err != nil {
	// 	t.Errorf("Not expected error: %s", err)
	// }

	// if resp.Code == "" {
	// 	t.Errorf("The registration not returns the code")
	// }
}

func TestRegisterWithoutClientHTTP(t *testing.T) {
	_, err := register("http://fakeaddress", "fakeaccount")

	if err == nil || err.Error() != "Resty http client not informed" {
		t.Errorf("The not informed address handler is not working")
	}
}

func TestRegisterWithoutAddress(t *testing.T) {
	_, err := register("", "fakeaccount")

	if err == nil || err.Error() != "Address not informed" {
		t.Errorf("The not informed address handler is not working")
	}
}

func TestRegisterWithoutAccount(t *testing.T) {
	_, err := register("http://fakeaddress", "")

	if err == nil || err.Error() != "Account not informed" {
		t.Errorf("The not informed account handler is not working")
	}
}

func TestRegisterUsingJWT(t *testing.T) {
	result, err := register("http://localhost:9999", "00000000000000")
	if err == nil {
		t.Errorf("Error at get content: %s", err)
	}

	fmt.Println(result)
	fmt.Println("End of test!")
}
