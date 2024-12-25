package Domain

import (
  "errors"
  // "strconv"
  // "fmt"
  "Requester/Interfaces"
)

// Argument Validators
// func argTValidation(Type string, arg interface{}) error{
//   switch Type{
//   case "string":
//     _, ok := arg.(string)
//     if !ok {
//       return errors.New("expected a string argument")
//     }
//   case "Request":
//     _, ok := arg.(Interfaces.Request)
//     if !ok {
//       return errors.New("expected a Request argument for ExtractAllForDir")
//     }
//   }
//   return nil
// }
//
// func argLValidation(args interface{}, lenArg int) error {
//   if len(args) != lenArg {
// 		return errors.New("invalid number of arguments")
// 	}
//   return nil
// }
// // Command Mngmtn
// func setArg(args ...interface{})(interface{}, error){
//
// }
// // Adapters
//
// func wrappedRequestAll(args ...interface{}) (interface{}, error) {
//   err := argLValidation(args, 1)
//   if err != nil {
//     return nil, err
//   }
//   err = argTValidation("string", args[0])
// 	if err != nil {
//     return nil, err
//   }
// 	res, err := RequestAll(strconv(args[0]))
// 	return res, err
// }
//
//func wrappedExtractAllForFile(args ...interface{}) (interface{}, error) {
//   err := argLValidation(args, 1)
//   if err != nil {
//     return nil, err
//   }
//   err = argTValidation("string", args[0])
// 	if err != nil {
//     return nil, err
//   }
// 	result, err := ExtractAllForFile(strconv(args[0]))
// 	return result, err
// }
//
// func wrappedExtractAllForDir(args ...interface{}) (interface{}, error) {
// 	err := argLValidation(args, 1)
//   if err != nil {
//     return nil, err
//   }
//   err = argTValidation(args[0], "Request")
// 	if err != nil {
//     return nil, err
//   }
//
//   result, err := ExtractAllForDir(args[0])
// 	return result, err
// }
//
// func wrappedExtractComments(args ...interface{}) (interface{}, error) {
//   err := argLValidation(args, 1)
//   if err != nil {
//     return nil, err
//   }
//   err = argTValidation(args[0], "string")
// 	if err != nil {
//     return nil, err
//   }
//   sarg := fmt.Sprintf(args[0])
//   res, err := ExtractComments(sarg) 
//   if err != nil {
//     return nil, err
//   }
//   return res, nil
// }
//
// func wrappedExtractSources(args ...interface{}) (interface{}, error) {
//   err := argLValidation(args, 1)
//   if err != nil {
//     return nil, err
//   }
//   err = argTValidation(args[0], "string")
// 	if err != nil {
//     return nil, err
//   }
//
//   res, err := wrappedExtractSources(strconv(args[0])) 
//   if err != nil {
//     return nil, err
//   }
//   return res, nil
// }
//


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


