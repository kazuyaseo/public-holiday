package main

import(
    "fmt"
    "strings"
    "net/http"
    "io/ioutil"
    // "encoding/json"
    // "encoding/csv"
    "golang.org/x/text/encoding/japanese"
    "golang.org/x/text/transform"
)

type Holiday struct{
    Date int `json:"date"`
    Name string `json:"name"`
}

type Holidays []Holiday

func main (){
    // 内閣府が提供するCSVファイルのURL
    url := "http://www8.cao.go.jp/chosei/shukujitsu/syukujitsu.csv"

    // GET
    reps, _ := http.Get(url)
    defer reps.Body.Close()
    byteArray, _ := ioutil.ReadAll(reps.Body)
    repsBody := string(byteArray)

    // Shift-jis -> UTF-8 
    strReader := strings.NewReader(repsBody)
    decodedReader := transform.NewReader(strReader, japanese.ShiftJIS.NewDecoder())
    decoded, err := ioutil.ReadAll(decodedReader)
    if err != nil{
       fmt.Print("error")
    }
    csvStr := strings.Replace(string(decoded),"-","",-1)

    fmt.Print(csvStr)
}
