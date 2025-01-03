package Controllers

import (
  "net/http"
  "regexp"  
  "errors"
  "strings"
  "strconv"
  "fmt"
  "bytes"
)

func MethodValidation(Args []string) error {
  if len(Args) < 0 {
    err := errors.New("\n [i] No Arguments To eval")
    return err 
  }
  return nil  
}

func ValidateStructureFUZZ(Args []string) error {
  // "Usage: FUZZ -u https://www.test.com/FUZZ -w ./wordlist.txt -mc 200"

  var IsSetTarget bool 
  var IsSetWordlist bool 

  if len(Args) < 0 {
    err := errors.New("\n [i] No Arguments To eval")
    return err 
  }
  for _, Arg :=  range Args {
    switch Arg {
    case "-w":
      IsSetWordlist = true   
    case "-d":
      IsSetTarget = true   
    case "-u":
      IsSetTarget = true   
    }
  }
  if IsSetTarget == false {
    return errors.New("[i] Target Not Found On Command")
  }
  if IsSetWordlist == false {
    return errors.New("[i] Wordlist Not Selected")
  }
  return nil
}


func ValidateRequestArgs(Args []string) error {

  if len(Args) < 0 {
    return errors.New("[i] No Arguments To Eval")
  }

  for i, Arg := range Args {
    switch Arg {
    case "-u":
      Details := Args[i+1]
      if Details == "" {
        return errors.New("[i] Not URL To Request")
      } else if strings.Contains(Details, "http://") == false && strings.Contains(Details, "https://") == false {
        return errors.New("[i] No Valid URL To Request")
      }  
    default:
      return nil
    }
  } 

  return nil
}

func ValidateFilterAndMatchArgs(Arguments map[string]string)(bool, bool, []int, []int, error){
  
  Filter := false
  Match := false
  var FCodes []int
  var MCodes []int  

  if len(Arguments) < 0 {
    err := errors.New("[i] No Arguments To eval")
    return Filter, Match, FCodes, MCodes, err 
  }
  for Flag, Details := range Arguments {
    switch Flag {
      case "-fc":
        Filter = true
        contains := strings.Contains(Details, ",")
        if contains == true {
          sFCodes := strings.Split(Details, ",")
          for _, FCode := range sFCodes {
            iFCode, err := strconv.Atoi(FCode)
            if err == nil {
              FCodes = append(FCodes, iFCode)
            }
            fmt.Printf("\n[i] Error Converting S to I Filter Code: %s", FCode )
          }
        }
        iFCode, err := strconv.Atoi(Arguments["-fc"])
        if err != nil {
          fmt.Printf("\n[i] Error Converting S to I Filter Code: %s", Details )
        }
        FCodes = append(FCodes, iFCode)

      case "-mc":
        Match = true
        contains := strings.Contains(Details, ",")
        if contains == true {
          sMCodes := strings.Split(Details, ",")
          for _, MCode := range sMCodes {
            iMCode, err := strconv.Atoi(MCode)
            if err == nil {
              MCodes = append(MCodes, iMCode)
            }
            fmt.Printf("\n[i] Error Converting S to I Filter Code: %s", MCode )
          }
        }
        iMCode, err := strconv.Atoi(Arguments["-mc"])
        if err != nil {
          fmt.Printf("\n[i] Error Converting S to I Filter Code: %s", Details )
        }
        MCodes = append(MCodes, iMCode)
      default:
        return Filter, Match, FCodes, MCodes, nil
    }
  }
  return Filter, Match, FCodes, MCodes, nil
}

func Extract(re regexp.Regexp, findOn string) ([][]string, error){
  match := re.FindAllStringSubmatch(findOn, -1)
  if len(match) == 0 {
    err := errors.New("No Matches Found")
    return match, err
  } 
  return match, nil
}  

func RequestURL (Method, URL string) (http.Response, error) {
	body := bytes.NewBuffer([]byte(`{"key":"value"}`))
  req, err := http.NewRequest(Method, URL, body)

  // if len(Headers) > 0 {
  //   for key, value := range Headers {
  //     req.Header.Add(key, value)
  //   }
  // }
  
	client := &http.Client{}

	res, err := client.Do(req)
  if err != nil {
		return *res, err 
	}
  return *res, nil
}
