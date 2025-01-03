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

func EvalFiltersAndMatchs(Code int, List []int, Type string) bool {
  Pass := false 
  if len(List) > 0 {
    switch Type {
    case "Filter":
      for _, FCode := range List {
        if Code == FCode {
          continue
        }
        // Manage Print Response Info 
        Pass = true
      }
    case "Match":
      for _, MCode := range List {
        if Code == MCode {
          continue
        }
        Pass = true
      }
    }  
  }
  return Pass
}

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

func Foo (Details string) []int {
  var List []int

  contains := strings.Contains(Details, ",")
  if contains == true {
      Elements := strings.Split(Details, ",")
    for _, Element := range Elements {
      iElement, err := strconv.Atoi(Element)
      if err == nil {
        List = append(List, iElement)
      }
      fmt.Printf("\n[i] Error Converting S to I Filter Code: %s", Element )
    }
  }
  iElement, err := strconv.Atoi(Details)
  if err != nil {
    fmt.Printf("\n[i] Error Converting S to I Filter Code: %s", Details )
  }
  List = append(List, iElement)
  return List
}

func ValidateFilterAndMatchArgs(Arguments map[string]string)(bool, bool, []int, []int, []int, []int, error){
  
  Filter := false
  Match := false

  var FCodes []int
  var MCodes []int  
  var FLength []int
  var MLength []int  

  if len(Arguments) < 0 {
    err := errors.New("[i] No Arguments To eval")
    return Filter, Match, FCodes, MCodes, FLength, MLength, err 
  }

  for Flag, Details := range Arguments {
    switch Flag {
      case "-fc":
        Filter = true
        FCodes = Foo(Details)     
      case "-mc":
        Match = true
        MCodes = Foo(Details) 
      case "-fl":
        Filter = true
        FLength = Foo(Details)
      case "-ml":
        Filter = true
        MLength = Foo(Details)
      default:
        return Filter, Match, FCodes, MCodes, FLength, MLength, nil
    }
  }
  return Filter, Match, FCodes, MCodes, FLength, MLength, nil
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
  // TODO: ADD Headers Mngmnt
  // if len(Headers) > 0 {
  //   for key, value := range Headers {
  //     req.Header.Add(key, value)
  //   }
  // } 
	client := &http.Client{}

  res, err := client.Do(req)
  if err != nil {
    res := &http.Response{
      StatusCode: 0,
      ContentLength: 0,
      Body: nil,
    }
		return *res, err 
	}
  return *res, nil
}
