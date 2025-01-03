package Domain 

import (
  "fmt"
  "sync"
  "regexp"
  "errors"
  "strings"
  "strconv"
  //
  "Requester/Interfaces"
  "Requester/Controllers"
)

// UTILS

func RequestErrorManagement(Err error) error {
  contains := strings.Contains(Err.Error(), "timeout")
  if  contains == true {
    return errors.New("Request Timeout")
  }
  return Err 
}

// Workers

func RequestAndSaveResponse(id int, Requests <-chan Interfaces.Request, wg *sync.WaitGroup, Filter bool, Match bool, FCodes []int, MCodes []int){

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
    if Request.Err == nil && Request.Response.Code != 0 && Request.Response.ContentLength < 0 {

      // Write Without Filters And Mathces
      if (Filter == false) && (Match == false) {
        if err := ResWriter(Request); err != nil {
            fmt.Printf("\nError:%v\n", err)
        }
      }

      // Filter By Code
      if Filter == true {
        for _, FCode := range FCodes {
          if Request.Response.Code == FCode {
            continue
          }
          // Write response on his respective file
          if err := ResWriter(Request); err != nil {
            fmt.Printf("\nError:%v\n", err)
          }
        }
      }

      // Filter By Match
      if Match == true {
        for _, MCode := range MCodes {
          if Request.Response.Code == MCode {
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



func RequestFUZZ(id int, Requests <-chan Interfaces.Request, wg *sync.WaitGroup, FUZZConfig Interfaces.FUZZConfig ){

  Filter := FUZZConfig.Filter 
  Match := FUZZConfig.Match

  defer wg.Done()

  // Executes a Function for URL passed
  for Request := range Requests {
    // Do the request
    Request.RequestURL()
    URL := Request.URL
    Code := int(Request.Response.Code)
    ContentLength := int(Request.Response.ContentLength)
    // Eval if the response is valid  
    if Request.Err == nil && Request.Response.Code != 0 && Request.Response.ContentLength > 0{


      // Print if no exist Filters or Matchs 
      if (Filter == false) && (Match == false){
        fmt.Printf("\n[i] %s Code:%d Length: %d", URL, Code, ContentLength)
      }

      // Filter by Code
      if Filter == true {

        FCodes := FUZZConfig.Filters.FCodes
        FLengths := FUZZConfig.Filters.FLengths

        if Pass := Controllers.EvalFiltersAndMatchs(Code, FCodes, "Filter"); Pass == true {
          fmt.Printf("\n[i] %s Code:%d Length: %d", URL, Code, ContentLength)
        }

        if Pass := Controllers.EvalFiltersAndMatchs(ContentLength, FLengths, "Filter"); Pass == true {
          fmt.Printf("\n[i] %s Code:%d Length: %d", URL, Code, ContentLength)
        }
        
      }

      // Filter by Match
      if Match == true {

        MCodes := FUZZConfig.Matchs.MCodes
        MLengths := FUZZConfig.Matchs.MLengths

        if Pass := Controllers.EvalFiltersAndMatchs(Code, MCodes, "Match"); Pass == true {
          fmt.Printf("\n[i] %s Code:%d Length: %d", URL, Code, ContentLength)
        }
        if Pass := Controllers.EvalFiltersAndMatchs(ContentLength, MLengths, "Match"); Pass == true {
          fmt.Printf("\n[i] %s Code:%d Length: %d", URL, Code, ContentLength)
        }

      }

      // Default Case Without Filters or Matchs
    }else {
      // Response Error Format
      if Request.Err != nil {
        err := RequestErrorManagement(Request.Err)
        fmt.Printf("\n%s Request With Error: %v", Request.URL, err)
      }

    }
  } 
}




func RequestAll(Command Interfaces.Command) (Interfaces.Command, error) {
  
  Arguments := Command.Arguments
  // Verify Argument
  // TODO: Add ContentLength Filter And Request Config
  Filter, Match, FCodes, MCodes, _, _, err := Controllers.ValidateFilterAndMatchArgs(Command.Arguments)
  if err != nil {
    return Command, err
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
  URLS, err := Controllers.Reader(Arguments["-f"])

  if err != nil {
    serr :=  fmt.Sprintf("Error:\n%v", err)
    err = errors.New(serr)
    return Command, err
  }

  // Create Workers to do Request
  for w := 1; w < 3; w++{
    wg.Add(1)
    go RequestAndSaveResponse(w, cRequests, &wg, Filter, Match, FCodes, MCodes)
  }

  // Send to Workers URLs to Request
  for _, URL := range URLS{
    var Request Interfaces.Request
    Request.URL = URL
    Request.Method = "GET"
    cRequests <- Request 
  }
  close(cRequests)
  wg.Wait()
  fmt.Println("Requests Done, Saved on ./Results")
  return Command, nil
}

func FUZZ(Command Interfaces.Command) (Interfaces.Command, error) {
  var URLS []string  
  var Wordlist []string

  Arguments := Command.Arguments
  Threats := 3

  Method := "GET" 

  // Argument Validation
  Filter, Match, FCodes, MCodes, FLength, MLength, err := Controllers.ValidateFilterAndMatchArgs(Arguments)

  if err != nil {
    return Command, err
  }

  Filters := &Interfaces.Filters{
    FCodes: FCodes,
    FLengths: FLength, 
  } 
  Mathchs := &Interfaces.Matchs{
    MCodes: MCodes,
    MLengths: MLength,
  }

  FUZZConfig := &Interfaces.FUZZConfig{
    Threats: 3,
    Filter: Filter,
    Match: Match,
    Filters: *Filters,
    Matchs: *Mathchs,
  }
   
  var wg sync.WaitGroup 
  cRequests:=  make(chan Interfaces.Request) 

  // Argument Management
  for Flag, Details := range Arguments {

    if Flag == "-u" {
      URLS = append(URLS, Details)
    }
    if Flag == "-d" {
      URLS, err = Controllers.Reader(Details)
      if err != nil {
        return Command, err
      }
    }
    if Flag == "-T"{
      Threats, err = strconv.Atoi(Details)
      if err != nil {
        return Command, err
      }
      FUZZConfig.Threats = Threats
    }
    if Flag == "-X" {
      Method = Details 
    }
  }
  
  Command.FUZZConfig = *FUZZConfig

  Wordlist, err = Controllers.Reader(Arguments["-w"])
  if err != nil {
    return Command, err
  }

  // Create Workers to do Request
  for w := 1; w < Threats; w++{
    wg.Add(1)
    go RequestFUZZ(w, cRequests, &wg, *FUZZConfig)
  }
         
  // Send to Workers URLs to Request
  for _, URL := range URLS{
    if strings.Contains(URL, "FUZZ") == true {
      for _, Word := range Wordlist{
        URL = strings.Replace(URL, "FUZZ", Word, 1)
        var Request Interfaces.Request
        Request.URL = URL
        Request.Method = Method 
        cRequests <- Request
      } 
    }else{
      fmt.Printf("\n[i] URL: %s Not Contain FUZZ", URL)
    }
  }
  close(cRequests)
  wg.Wait()
  return Command, nil
}
