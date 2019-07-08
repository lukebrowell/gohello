package main

import (
  "testing"
  "encoding/json"
  "strconv"
)

func TestFoo(t *testing.T){
  expected := "brilliant"
  actual := foo()
  if actual != expected {
    t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
  }
}

func TestGetSample(t *testing.T){
  expected := `{
  "userId": 1,
  "id": 1,
  "title": "delectus aut autem",
  "completed": false
}
`
  actual := getSample()
  if actual != expected {
    t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
  }
}

func TestPostSample(t *testing.T){
  expected := 456
  jsonData := map[string]string{"firstname": "Luke", "lastname": "Browell"}
  jsonValue, _ := json.Marshal(jsonData)

  actual := len(postSample(jsonValue))

  if actual != expected {
    t.Errorf("Test failed, expected: '%s', got:  '%s'", strconv.FormatInt(int64(expected), 10), strconv.FormatInt(int64(actual), 10))
  }
}