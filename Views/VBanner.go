package Views

import (
  "fmt"
  "Requester/Controllers"
  // "sync"
)
var Reset = "\033[0m" 
var Red = "\033[31m" 
var Green = "\033[32m" 
var Yellow = "\033[33m" 
var Blue = "\033[34m" 
var Magenta = "\033[35m" 
var Cyan = "\033[36m" 
var Gray = "\033[37m" 
var White = "\033[97m"

func PrintBannerLogo() {
  lines, err := Controllers.Reader("./Views/bannerLogo.txt")
  if err != nil {
    fmt.Printf("\nError:\n%v\n", err)
    return
  }
  for _, line := range lines {
    println(Magenta + line + Reset)
  }
  fmt.Printf("\n")
}

func PrintBannerLet(){
  lines, err := Controllers.Reader("./Views/bannerLet.txt")
  if err != nil {
    fmt.Printf("\nError:\n%v\n", err)
    return
  }
  for _, line := range lines {
    println(Magenta + line + Reset)
  }
  fmt.Printf("\n")

}

func PrintMenu() string{

  menu := Controllers.NewMenu("Mode")

  menu.AddItem("RequestAll", "requestAll")
  menu.AddItem("FUZZAll", "fuzzAll")
  
  choice := menu.Display()
  return choice

}
