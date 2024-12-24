package main

import (
  "fmt"
  "flag"
  "os"
  "strings"

  "Requester/Managers"
  "Requester/Views"
)

  

func main(){
  args := os.Args
  if len(args) < 1 {
    Views.PrintBannerLogo()
    Views.PrintBannerLet()
    Views.PrintMenu()
  }

  silentPtr := flag.Bool("s", false, "silent mode")
  interactivePtr := flag.Bool("i", true, "interactive mode")
  interType := flag.String("t", "r", "Type of Interaction")
  reqFilePtr := flag.String("rf", "./TestALiveScope.txt", "file for http request") 
  flag.Parse()

  if *silentPtr == false {
    Views.PrintBannerLogo()
    Views.PrintBannerLet()

  }

  if *interactivePtr == true {
    Views.PrintMenu()
  }
  
  switch{
  case strings.Contains(*interType, "r"):
    err := Managers.RequestAll(*reqFilePtr)
    if err != nil {
      fmt.Printf("Error: %v", err)
    }
  case strings.Contains(*interType, "f"):
    fmt.Printf("F")
  case strings.Contains(*interType, "x"):
    Managers.ExtractAll()
  }

} 
