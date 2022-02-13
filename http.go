package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
)


func main() {
  url := "http://localhost"
  resp, _ := http.Get(url)
  defer resp.Body.Close()
  req, _ := http.NewRequest("GET", url, nil)
  req.Header.Set("applicatiom/json")
  client := new(http.Client)
  response, err := client.Do(req)
  dump, _ := httputil.DumpRequestOut(req, true)
  fmt.Printf("%s", dump)
  client := new(http.Client)
  resp, err := client.Do(req)
  byteArray, _ := ioutil.ReadAll(resp.Body)
  fmt.Println(string(byteArray))
}
