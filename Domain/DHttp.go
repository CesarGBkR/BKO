package Domain 

import (
  "fmt"
  "sync"
  "regexp"
  "errors"
  "strings"
  "strconv"

  "Requester/Interfaces"
  "Requester/Controllers"
)

func Filter(){

}

func Worker(id int, Requests <-chan Interfaces.Request, wg *sync.WaitGroup, Filter bool, Match bool, FCodes []int, MCodes []int){

  defer wg.Done()

  // Executes a Function for URL passed
  for Request := range Requests {
    // Eval if URL contains a SubDir
    re := regexp.MustCompile(`^https?://(.*)/?`)
    match := re.FindStringSubmatch(Request.URL)

    // Eval if URL is valid
    if len(match) < 1 {
      fmt.Printf("\n%s\nError Finding Valir URL on Creation Dir", Request.URL)
      continue
    }
    // Create a dir for every valid URL to save data related
    dir := string(match[1])
    dirPath := fmt.Sprintf("./Results/Responses/%s", dir)

    // Eval if already exist the Directory
    dirExist := Controllers.DirectoryExists(dirPath)
    if dirExist == false {
      // If dont exist try to create the Directory
      if err := Controllers.CreateFile(dirPath); err != nil{
        fmt.Printf("\nError:%v\n", err)
        continue
      }
    }
    // If exist continue with the request URL
    // Do the request
    Request.RequestURL()

    // Eval if the response is valid  
    if Request.Err == nil && Request.Code != 0 && Request.Response.Body != nil{
      // Filter by Code
      if Filter == true {
        for _, FCode := range FCodes {
          if Request.Code == FCode {
            continue
          }
          // Write response on his respective file
          if err := ResWriter(Request); err != nil {
            fmt.Printf("\nError:%v\n", err)
          }
        }
      }
      if Match == true {
        for _, MCode := range MCodes {
          if Request.Code == MCode {
            // Write response on his respective file
            if err := ResWriter(Request); err != nil {
              fmt.Printf("\nError:%v\n", err)
            }
          }
          continue
        }
      }
    }
  } 
}


func RequestAll(Command Interfaces.Command) (Interfaces.Command, error) {
  Filter := false
  Match := false
  var FCodes []int
  var MCodes []int  

  // Verify Argument
  if Command.Arguments[0] == "" {
    err := errors.New("No File Specified in Argument")
    return Command, err
  }

  for i, Arg := range Command.Arguments {
    if Arg == "fc" {
      Filter = true
      contains := strings.Contains(Command.Arguments[i+1], ",")
      if contains == true {
        sFCodes := strings.Split(Command.Arguments[i+1], ",")
        for _, FCode := range sFCodes {
          iFCode, err := strconv.Atoi(FCode)
          if err == nil {
            FCodes = append(FCodes, iFCode)
          }
        }
      }

      iFCode, err := strconv.Atoi(Command.Arguments[i+1])
      if err != nil {
        err = errors.New("\nNo Codes To Filter Specified")
      }
      FCodes = append(FCodes, iFCode)

    }
    if Arg == "mc" {
      Match = true
      contains := strings.Contains(Command.Arguments[i+1], ",")
      if contains == true {
        sMCodes := strings.Split(Command.Arguments[i+1], ",")
        for _, MCode := range sMCodes {
          iMCode, err := strconv.Atoi(MCode)
          if err == nil {
            MCodes = append(MCodes, iMCode)
          }
        }
      }
      iMCode, err := strconv.Atoi(Command.Arguments[i+1])
      if err != nil {
        err = errors.New("\nNo Codes To Match Specified")
      }
      MCodes = append(MCodes, iMCode) 
    }
  }
  
  // Verify Results Directory and Responses Directory
   ok := Controllers.DirectoryExists("./Results"); 
   if ok == false {
     err := Controllers.CreateFile("./Results")
     if err != nil {
       return Command, err
     }
   }

  ok = Controllers.DirectoryExists("./Results/Responses"); 
   if ok == false {
     err := Controllers.CreateFile("./Results/Responses")
     if err != nil {
       return Command, err
     }
  }

  var wg sync.WaitGroup 
  cRequests:=  make(chan Interfaces.Request) 
  
  // Read file with URLs to Request
  URLS, err := Controllers.Reader(Command.Arguments[0])

  if err != nil {
    serr :=  fmt.Sprintf("Error:\n%v", err)
    err = errors.New(serr)
    return Command, err
  }

  // Create Workers to do Request
  for w := 1; w < 3; w++{
    wg.Add(1)
    go Worker(w, cRequests, &wg, Filter, Match, FCodes, MCodes)
  }
  // Send to Workers URLs to Request
  for _, URL := range URLS{
    var Request Interfaces.Request
    Request.URL = URL
    cRequests <- Request 
  }
  close(cRequests)
  wg.Wait()
  fmt.Println("Requests Done, Saved on ./Results")
  return Command, nil
}
