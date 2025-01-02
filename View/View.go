package View

import (
  "fmt"
  // "errors"

  "Requester/Interfaces"
  "Requester/Domain"

	// "github.com/manifoldco/promptui"
  "github.com/abiosoft/ishell"
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
  lines, err := Domain.Reader("./View/bannerLogo.txt")
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
  lines, err := Domain.Reader("./View/bannerLet.txt")
  if err != nil {
    fmt.Printf("\nError:\n%v\n", err)
    return
  }
  for _, line := range lines {
    println(Magenta + line + Reset)
  }
  fmt.Printf("\n")

}

func IShell(){

  // create new shell.
  // by default, new shell includes 'exit', 'help' and 'clear' commands.
  shell := ishell.New()

  // register a function for "greet" command.
  for _, Command := range Interfaces.Commands {
    shell.AddCmd(&ishell.Cmd{
      Name: Command.Name,
      Help: Command.Help,
      LongHelp: Command.LongHelp,
      Func: func(c *ishell.Context) {
        var Arguments map[string]string

        if len(c.Args) > 0 {
          Command.ArgumentsRaw := c.Args  
        } else {
          serr := fmt.Sprintf("\n[i]No Arguments For Command")
          err := errors.New(serr)
          informer := fmt.Sprintf("\n\n[!] Error Executing: %s\n%v", Command.Name, err)
          c.Printf(informer)
          return
        }
        if err := Domain.ValidateAllArguments(Command); err != nil {
          informer := fmt.Sprintf("\n\n[!] Error Executing: %s\n%v", Command.Name, err)
          c.Printf(informer)
          return
        }
        for i, Argument := COmmand.ArgumentsRaw {
          Arguments[Argument] = Argument[i+1]
        }
        _, err := Domain.CommandSwitcher(Command) 

        if err != nil {
          informer := fmt.Sprintf("\n\n[!] Error Executing: %s\n%v", Command.Name, err)
          c.Printf(informer)
        } else{
          informer := fmt.Sprintf("\n\n[+] Command: %s Executed Correctly\n", Command.Name)
          c.Printf(informer)
        }
      },
    })
  } 

  // run shell
  shell.Run()
}










