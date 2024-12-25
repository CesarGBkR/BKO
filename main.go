package main

import (
  "Requester/Views"
  "Requester/Interfaces"
)

  

func main(){
  var Command Interfaces.Command

  View.PrintBannerLogo()
  View.PrintBannerLet()
  View.Shell(Command)

} 
