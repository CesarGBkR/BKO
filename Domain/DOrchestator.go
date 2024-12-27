package Domain

import (
  "fmt"
  "errors"
  "Requester/Interfaces"
)

func ShowCommand(Command Interfaces.Command)(Interfaces.Command, error){
  if Command.Name != "" {
    fmt.Printf("Command: \n%s\n", Command.Name)
  }
  if Command.Arguments[0] != "" {
    fmt.Printf("Argument: \n%s\n", Command.Arguments[0])
  }
  if Command.Output != nil {
    fmt.Printf("Output: \n%v\n", Command.Output)
  }
  if Command.Error != nil{
    fmt.Printf("\n\t%v\n", Command.Error)
  }   
  return Command, nil
}

func CommandSwitcher(Command Interfaces.Command)(Interfaces.Command, error){

 var Commands = map[string]func(Command Interfaces.Command) (Interfaces.Command, error) {
  "SetArg": SetArg, 
  "ShowCommand": ShowCommand,
  "RequestAll": RequestAll, 
  "ExtractAllForFile": ExtractAllForFile,
  "ExtractAllForDir": ExtractAllForDir,
  "ExtractComments": ExtractComments,
  "ExtractSources": ExtractSources,
  }

  if fn, found := Commands[Command.Name]; found{
    Command, err :=  fn(Command)
    if err != nil {
      return Command, err
    }
    return Command, nil
  } else {
    err := errors.New("Command Not Found")
    return Command, err
  }
  return Command, nil 
}


