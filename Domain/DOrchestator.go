package Domain

import (
  "fmt"
  "errors"
  "Requester/Interfaces"

	"github.com/manifoldco/promptui"
)
func ShowCommand(Command Interfaces.Command)(Interfaces.Command, error){
  if Command.Name != "" {
    fmt.Printf("Command: \n%s\n", Command.Name)
  }
  if Command.Argument != "" {
    fmt.Printf("Argument: \n%s\n", Command.Argument)
  }
  if Command.Output != nil {
    fmt.Printf("Output: \n%v\n", Command.Output)
  }
  if Command.Error != nil{
    fmt.Printf("\n\t%v\n", Command.Error)
  }   
  return Command, nil
}

func SetArg(Command Interfaces.Command)(Interfaces.Command, error){
    
    // Render Current Prompt
    templates := &promptui.PromptTemplates{
      Prompt:  "{{ . }} ",
      Success: "{{ . | bold | green }} ",
    }

    prompt := promptui.Prompt{
      Label:     "ARG>",
      Templates: templates,
    }

    arg, err := prompt.Run()

    if err != nil {
      err := errors.New(fmt.Sprintf("Prompt failed %v\n", err))
      return Command, err
    }
    
    Command.Argument = arg
    ShowCommand(Command)
    return Command, err 

}

func CommandSwitcher(Command Interfaces.Command)(Interfaces.Command, error){

 var Commands = map[string]func(Command Interfaces.Command) (Interfaces.Command, error) {
  "SetArg": SetArg, 
  "ShowCommand": ShowCommand,
  "RequestAll": RequestAll, 
  "ExtractAllForFile": ExtractAllForFile,
  "ExtractAllForDir": ExtractAllForDir,
  "ExtractComments": ExtractComments,
  "wrappedExtractSources": ExtractSources,
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


