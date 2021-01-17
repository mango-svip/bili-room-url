package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "github.com/atotto/clipboard"
    "io/ioutil"
    "net/http"
    "os"
)


var (
    roomId string
    requestUrl = "https://api.live.bilibili.com/xlive/web-room/v1/playUrl/playUrl?cid=%s&platform=h5&otype=json&quality=0"
)


func init() {
    flag.Parse()
    roomId = flag.Arg(0)
}



func main() {

    if roomId == "" {
        fmt.Print("Please input room number")
        os.Exit(-1)
    }

    realUrl := fmt.Sprintf(requestUrl, roomId)

    resp, err := http.Get(realUrl)

    if err != nil {
        panic(err)
    }

    defer resp.Body.Close()

    bytes, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        panic(err)
    }

    urlMap := make(map[string]interface{})

    err = json.Unmarshal(bytes, &urlMap)

    if err != nil {
        panic(err)
    }

    urls := urlMap["data"].(map[string]interface{})["durl"].([]interface{})


    fmt.Println("roomId ", roomId)
    for _,v := range urls {
        fmt.Println(v.(map[string]interface{})["url"])
    }

    // copy first url into clipboard
    clipboard.WriteAll(urls[0].(map[string]interface{})["url"].(string))






}
