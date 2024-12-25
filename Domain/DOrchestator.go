package Domain

import (
  "errors"
  "Requester/Interfaces"
)

func CommandSwitcher(Command Interfaces.Command)error{

 var Commands = map[string]func(args string) error {
  "RequestAll": RequestAll, 
  }

  if fn, found := Commands[Command.Name]; found{
    fn(Command.Arguments)
  } else {
    err := errors.New("Command Not Found")
    return err
  }
  return nil
}


