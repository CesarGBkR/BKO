package Domain 

import (
  "fmt"
  "os"
  "sync"
  "regexp"
  "errors"

  "Requester/Interfaces"
  "Requester/Controllers"
)


func Worker(id int, Requests <-chan Interfaces.Request, wg *sync.WaitGroup){

  defer wg.Done()

  // Codes for match or filter
  ErrorCodes := [1]int{404}

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
      for _, c := range ErrorCodes {
        if Request.Code == c {
          continue
        }
        // Write response on his respective file
        if err := ResWriter(Request); err != nil {
          fmt.Printf("\nError:%v\n", err)
        }
      }
    }
  } 
}


func RequestAll(Command Interfaces.Command) (Interfaces.Command, error) {
  
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
  URLS, err := Controllers.Reader(Command.Argument)

  if err != nil {
    serr :=  fmt.Sprintf("Error:\n%v", err)
    err = errors.New(serr)
    return Command, err
  }

  // Create Workers to do Request
  for w := 1; w < 3; w++{
    wg.Add(1)
    go Worker(w, cRequests, &wg)
  }
  // Send to Workers URLs to Request
  for _, URL := range URLS{
    var Request Interfaces.Request
    Request.URL = URL
    cRequests <- Request 
  }
  close(cRequests)
  wg.Wait()
  Command.Output = "Requests Done, Saved on ./Results"
  return Command, nil
}

func ExtractSources(Command Interfaces.Command)(Interfaces.Command, error){
  re := regexp.MustCompile(`src=["']([^"']+)["']`)
  elements, err := Controllers.Extract(*re, Command.Argument)
  if err != nil{
    return Command, err
  }else {
    if len(elements) > 1 {
      Command.FileData[0].Sources = elements
      return Command, nil
    }
  }
  err = errors.New("No Matches Found")
  return Command, err
}

func ExtractComments(Command Interfaces.Command)(Interfaces.Command, error){

  re := regexp.MustCompile(`<!--(.*?)-->`)
  elements, err := Controllers.Extract(*re, Command.Argument)
  if err != nil{
    return Command, err
  }else {
    if len(elements) > 1 {
      Command.FileData[0].Comments = elements
      return Command, nil
    }
  }
  err = errors.New("No Matches Found")
  return Command, err
}

func ExtractAllForFile(Command Interfaces.Command) (Interfaces.Command, error){
  content, err := Controllers.FContentReader(Command.Argument)
  if err != nil {
    return Command, err
  }
  Command.Argument = content  
  Command, err1 := ExtractSources(Command)
  Command, err2 := ExtractComments(Command)
  switch {
  case err1 != nil && err2 != nil:
    err := errors.New("No sources and no comments found")
    return Command, err
  case err1 != nil && err2 == nil :
    err := errors.New("No sources found")
    return Command, err 
  case err1 == nil && err2 != nil :
    err := errors.New("No conmments found")
    return Command, err 
  }
  return Command, nil 
}

func ExtractAllForDir(Command Interfaces.Command) (Interfaces.Command, error){

  files, err := os.ReadDir(Command.Argument)
  if err != nil {
    return Command, err
	}
  for _, file := range files {
		if file.IsDir() {
      continue
		} else {
      path := fmt.Sprintf("%s/%v", Command.Argument, file) 
      Command.Argument = path
      Command, err := ExtractAllForFile(Command)
      return Command, err
		}
  }
  return Command, nil 
}
