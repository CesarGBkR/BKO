package Controllers 

func ValidateAllArguments(Arguments []string, Type string) error {
  switch Type {
  case "FUZZ":
    if err := ValidateStructureFUZZ(Arguments); err != nil {
      return err
    }  
    if err := ValidateFileArguments(Arguments); err != nil {
      return  err
    }
    if err := ValidateRequestArgs(Arguments); err != nil {
      return err
    }
  case "RequestAll":
    if err := ValidateFileArguments(Arguments); err != nil {
      return  err
    }
  } 
  return nil
}
