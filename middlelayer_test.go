package main

import (
	//	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	//	"os"
	//	"strings"
	"encoding/json"
	"fmt"
	"testing"
)

var expected = `{"1":{"Productname":"sample"}}`

func Test(t *testing.T) {

	//mock server
	router := prod.router()
	server := httptest.NewServer(router)
	defer server.Close()

	//3 things can go, reaching the server, communicating with the
	//server and getting the wrong data back

	resp, err := http.Get(server.URL + "/list/")

	fmt.Printf("%v", resp)

	if err != nil {

		t.Fatal("cant finde the server")

	}

	if resp.StatusCode != 200 {

		t.Fatal("not 200 status returned")
	}

	outputinbytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {

		t.Fatal("cant read the body")

	}

	data := new(Product)

	err = json.Unmarshal(outputinbytes, &data)

	if err != nil {

		t.Fatal("cant retrieve json")
	}

	fmt.Println(data)

	outstring := string(outputinbytes)

	if outstring != expected {

		fmt.Printf("%q", outstring)
		t.Fatal("cant read the body")

	}

}
