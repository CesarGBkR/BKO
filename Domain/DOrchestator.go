package Domain

import (
  "errors"
  "Requester/Interfaces"
)

func CommandSwitcher(Command Interfaces.Command)error{

 var Commands = map[string]func(Command Interfaces.Command) (Interfaces.Command, error) {
  // "SetArg": SetArg, 
  "RequestAll": RequestAll, 
  "ExtractAllForFile": ExtractAllForFile,
  "ExtractAllForDir": ExtractAllForDir,
  "ExtractComments": ExtractComments,
  "wrappedExtractSources": ExtractSources,
  }

  if fn, found := Commands[Command.Name]; found{
    fn(Command)
    
  } else {
    err := errors.New("Command Not Found")
    return err
  }
  return nil
}


