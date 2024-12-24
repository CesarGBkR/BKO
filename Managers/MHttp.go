package Managers 

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


func RequestAll(file string) error {
  // Verify Results Directory and Responses Directory
   ok := Controllers.DirectoryExists("./Results"); 
   if ok == false {
     err := Controllers.CreateFile("./Results")
     if err != nil {
       return err
     }
   }

  ok = Controllers.DirectoryExists("./Results/Responses"); 
   if ok == false {
     err := Controllers.CreateFile("./Results/Responses")
     if err != nil {
       return err
     }
  }

  var wg sync.WaitGroup 
  cRequests:=  make(chan Interfaces.Request) 
  
  // Read file with URLs to Request
  URLS, err := Controllers.Reader(file)

  if err != nil {
    fmt.Printf("Error:\n%v", err)
    os.Exit(1)
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
  fmt.Printf("\nWorkersDone")
  return nil
}

func ExtractSources(findOn string)([][]string, error){
  re := regexp.MustCompile(`src=["']([^"']+)["']`)
  elements, err := Controllers.Extract(*re, findOn)
  if err != nil{
    return nil, err
  }else {
    if len(elements) > 1 {
      return elements, nil
    }
  }
  err = errors.New("No Matches Found")
  return nil, err
}

func ExtractComments(findOn string)([][]string, error){
  
  re := regexp.MustCompile(`<!--(.*?)-->`)
  elements, err := Controllers.Extract(*re, findOn)
  if err != nil{
    return nil, err
  }else {
    if len(elements) > 1 {
      return elements, nil
    }
  }
  err = errors.New("No Matches Found")
  return nil, err
}

func ExtractAllForFile(Path string) (Interfaces.FileData, error){
  var FileData Interfaces.FileData
  var src []string
  var cmt []string

  content, err := Controllers.FContentReader(Path)
  if err != nil {
    return FileData, err
  }

  sources, err1 := ExtractSources(content)
  comments, err2 := ExtractComments(content)
  switch {
  case err1 != nil && err2 != nil:
    err := errors.New("No sources and no contents found")
    return FileData, err
  case err1 != nil && err2 == nil :
    for _, element := range comments{
      cmt = append(cmt, element[1])
    }
  case err1 == nil && err2 != nil :

    for _, element := range sources{
      src = append(src, element[1])
    }
  }
  for _, element := range sources {
    src = append(src, element[1])
  }
  for _, element := range comments{
    cmt = append(cmt, element[1])
  } 
  // case err1 == nil && err2 == nil :
  FileData.Sources = src 
  FileData. Comments = cmt 
  return FileData, nil 
}

func ExtractAllForDir(Request Interfaces.Request) ([]Interfaces.FileData, error){
  var FilesData []Interfaces.FileData

  files, err := os.ReadDir(Request.Dir)
  if err != nil {
    return FilesData, err
	}
  for _, file := range files {
		if file.IsDir() {
      continue
		} else {
      path := fmt.Sprintf("%s/%v", Request.Dir, file) 
      FileData, err := ExtractAllForFile(path)
      if err != nil {
        continue
      }
      FileData.Dir = Request.Dir 
      FilesData = append(FilesData, FileData) 
		}
  }
  return FilesData, nil 
}
