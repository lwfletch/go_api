package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllUsers(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/users", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAllUsers)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"id":1,"first_name":"Trey","last_name":"Fletcher","email_address":"treyfletcher612@gmail.com","birth_date":"06-12-1986"},
	{"id":2,"first_name":"Lewis","last_name":"Fletcher","email_address":"lwfletcher86@gmail.com","birth_date":"06-12-1986"},
	{"id":3,"first_name":"Abby","last_name":"Fletcher","email_address":"abbygirl716@gmail.com","birth_date":"07-16-2012"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetUser(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"id":1,"first_name":"Trey","last_name":"Fletcher","email_address":"treyfletcher612@gmail.com","birth_date":"06-12-1986"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetUserByIDNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/users/123", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status == http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestCreateUser(t *testing.T) {

	var jsonStr = []byte(`{"id":"4","first_name":"Liz","last_name":"Fletcher","email_address":"lizfletcher@nc.rr.com","birth_date":"10-06-1863"}`)

	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"id":4,"first_name":"Liz","last_name":"Fletcher","email_address":"lizfletcher@nc.rr.com","birth_date":"10-06-1963"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestUpdateUser(t *testing.T) {

	var jsonStr = []byte(`{"id":2,"first_name":"Lewis","last_name":"Fletcher III","email_address":"extremejv@gmail.com","birth_date":"06-11-1986"}`)

	req, err := http.NewRequest("PUT", "/users/2", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"id":2,"first_name":"Lewis","last_name":"Fletcher","email_address":"extremejv@gmail.com","birth_date":"06-11-1986"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
func TestDeleteUser(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/users/2", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `[{"id":1,"first_name":"Trey","last_name":"Fletcher","email_address":"treyfletcher612@gmail.com","birth_date":"06-12-1986"},
	{"id":"3", "first_name":"Abby", "last_name":"Fletcher","email_address":"abbygirl716@gmail.com","birth_date":"07-16-2012"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
