package Views

import (
  "fmt"
  "errors"
  "Requester/Interfaces"
  "Requester/Domain"
	"github.com/manifoldco/promptui"
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

var Commands = []string{"RequestAll", "ExtractAll","ExtractSrc", "Help", "SetArg"}

func PrintBannerLogo() {
  lines, err := Domain.Reader("./Views/bannerLogo.txt")
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
  lines, err := Domain.Reader("./Views/bannerLet.txt")
  if err != nil {
    fmt.Printf("\nError:\n%v\n", err)
    return
  }
  for _, line := range lines {
    println(Magenta + line + Reset)
  }
  fmt.Printf("\n")

}


func Shell() {
  var Command Interfaces.Command
  validate := func(input string) error {
    for _, Command  := range Commands {
      if Command == input {
        return nil
      } 
    }
    err := errors.New("Command Not Found")
    return err
    }
    // Render Current Prompt
    templates := &promptui.PromptTemplates{
      Prompt:  "{{ . }} ",
      Valid:   "{{ . | green }} ",
      Invalid: "{{ . | red }} ",
      Success: "{{ . | bold }} ",
    }

    prompt := promptui.Prompt{
      Label:     "BKO>",
      Templates: templates,
      Validate:  validate,
    }

    result, err := prompt.Run()

    if err != nil {
      fmt.Printf("Prompt failed %v\n", err)
      return
    }
  Command.Name = result
  Command.Argument = "./TestAliveScope.txt"
  Domain.CommandSwitcher(Command)
}
