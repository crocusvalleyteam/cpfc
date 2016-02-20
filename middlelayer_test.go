package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var testdata = map[string]Product{

	"1": Product{Productname: "tree"},
	"2": Product{Productname: "chair"},
	"3": Product{Productname: "coffee"},
	"4": Product{Productname: "bread"},
	"5": Product{Productname: "suger"},
}

//Test_returnallproducts to test the return the list rest service
//the test is simple. it check, using a mock server, if a known json is returned
//the key to this test in the list router of findallproducts is used
func Test_returnallproducts(t *testing.T) {

	//first covert test data into json string
	testdatabyte, err := json.Marshal(testdata)
	if err != nil {

		t.Fatal("failed to marshal test data")
	}
	expected := string(testdatabyte)

	//mock server with router pointing to func to be tested

	router := prod.router()
	server := httptest.NewServer(router)
	defer server.Close()

	//3 things can go, reaching the server, communicating with the
	//server and getting the wrong data back

	resp, err := http.Get(server.URL + "/list/")

	if err != nil {

		t.Fatal("cant finde the server")
	}

	if resp.StatusCode != 200 {

		t.Fatal("not 200 status returned")
	}

	outputinbytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {

		t.Fatal("cant read the body of the response")
	}

	//if all ok  read the body into json
	data := new(Product)
	err = json.Unmarshal(outputinbytes, &data)
	if err != nil {

		t.Fatal("cant retrieve json")
	}
	outstring := string(outputinbytes)

	//check if input and output are ok
	if outstring != expected {

		fmt.Printf("%q", outstring)
		t.Fatal("cant read the body")

	}
}
