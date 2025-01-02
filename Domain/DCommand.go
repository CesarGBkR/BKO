package Domain

import (
  "errors"

  "Requester/Controllers"
  "Requester/Interfaces"

)

func ValidateAllArguments(Command Interfaces.Command) error {

  Arguments := Command.ArgumentsRaw

  if len(Arguments) > 0 {
    err := Controllers.ValidateAllArguments(Arguments, Command.Name)
    if err != nil {
      return err
    }
    return nil
  }
  return errors.New("[i] No Arguments To Eval")

}
