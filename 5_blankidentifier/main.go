package main

import(
"fmt"
"net/http"
"log"
"io/ioutil")

func main(){
    res,err:=http.Get("http://www.espncricinfo.com/")
    
    if(err!=nil){
        log.Fatal("error returned")
    }
    page,err:= ioutil.ReadAll(res.Body)
    res.Body.Close()
    
    if err!=nil{
        log.Fatal(err)
    }
    fmt.Println(page)
    
    
    //blank identiifer
    //The client must close the response body when finished with it:
    resp,_:= http.Get("http://google.com")
    page,_:= ioutil.ReadAll(res.Body)
    res.Body.Close()
    fmt.Println(page)
    
}