package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var router = Router()

func Test_crystalpalacerestservice(t *testing.T) {

	server := httptest.NewServer(Router())
	defer server.Close()
	f(t)

}

func f(t *testing.T) {

	var expected string

	//read test data from file
	if data, err := ioutil.ReadFile("testoneresultcase.txt"); err != nil {

		t.Error("failed to read test file")

	} else {

		expected = string(data)
	}

	//test success and failure : todo
	req, _ := http.NewRequest("GET", "/results/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	//success
	r := assert.NotEqual(t, resp.Body.String(), expected)

	if r != true {

		t.Error("falied")

	} else {

		fmt.Printf("\n \n Got : %s \n\n AND \n \n expected : \n\n %s", resp.Body.String(), expected)

	}

}
