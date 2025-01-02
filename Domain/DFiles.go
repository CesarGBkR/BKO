package Domain 

import (
  "fmt"
  "os"
  // "io/ioutil"
  "io"
  "regexp"
  "strings"
  "errors"
  "time"

  "Requester/Interfaces"
  "Requester/Controllers"
)

func Reader(fPath string) ([]string, error){
  res, err := Controllers.Reader(fPath)
  return res, err
}

// Write File
func ResWriter(Request Interfaces.Request) (error){
  Response := Request.Response 
  re := regexp.MustCompile(`^https?://.*/(.*)`)
  match := re.FindStringSubmatch(Request.URL)

  if len(match) < 1 {
    path := fmt.Sprintf("./Results/Responses/%s/index.html", Request.Dir )
    if fileExists := Controllers.FileExists(path); fileExists == true {
      serr := fmt.Sprintf("File: %s exist", Request.Dir)
      err := errors.New(serr)
      return err 
    }
  }

  subdirectories := match[1]
  converted := strings.ReplaceAll(subdirectories, "/", "\\/")
  fileDir := fmt.Sprintf("./Results/Response/%s/%s", Request.Dir, converted)

  bodyBytes, err := io.ReadAll(Response.RawBody)
  if err != nil {
    return err
  }

  err = Controllers.Writer(fileDir, bodyBytes)
  if err != nil {
    return err
  }
  return nil
}

func ExtractSources(Command Interfaces.Command)(Interfaces.Command, error){
  re := regexp.MustCompile(`src=["']([^"']+)["']`)

  elements, err := Controllers.Extract(*re, Command.Arguments["-f"])
  if err != nil{
    return Command, err
  }else {
  if len(elements) > 1 {
      var Data Interfaces.FileData
      Data.Sources = elements
      Command.FileData = append(Command.FileData, Data) 
      return Command, nil
    }
  }
  err = errors.New("No Matches Found")
  return Command, err
}

func ExtractComments(Command Interfaces.Command)(Interfaces.Command, error){

  re := regexp.MustCompile(`<!--(.*?)-->`)
  elements, err := Controllers.Extract(*re, Command.Arguments["-f"])
  if err != nil {
    return Command, err
  }else {
    if len(elements) > 1 {
      var Data Interfaces.FileData
      Data.Sources = elements
      Command.FileData = append(Command.FileData, Data) 
      return Command, nil    }
  }
  err = errors.New("No Matches Found")
  return Command, err
}

func ExtractAllForFile(Command Interfaces.Command) (Interfaces.Command, error){
  
  // Verify Argument
  if Command.Arguments["-f"] == "" {
    err := errors.New("No File Specified in Argument")
    return Command, err
  }
  DirCommand := Command.Arguments["-f"] 
  content, err := Controllers.FContentReader(Command.Arguments["-f"])
  if err != nil {
    return Command, err
  }
  Command.Arguments["-f"] = content  
  Command, err1 := ExtractSources(Command)
  Command, err2 := ExtractComments(Command)
  
  Command.Arguments["-f"] = DirCommand
  // TODO:  Validation of existence of Dir and file
  Dir :=  fmt.Sprintf("Results/Responses/%s", Command.Arguments["-f"])
  cDir := fmt.Sprintf("./%s/comments.txt", Dir ) 
  sDir := fmt.Sprintf("./%s/sources.txt", Dir) 

  if exist := Controllers.DirectoryExists(Dir); exist == false {
    err = Controllers.CreateFile(Dir)
    if err != nil {
      sErr := fmt.Sprintf("Error Creating Directory ./Results/Response/%s/\nError:\n%v", Command.Arguments["-f"], err)
      err = errors.New(sErr)
      return Command, err
    }
  }
  
  if exist := Controllers.FileExists(cDir); exist == true {
    currentTime := time.Now()
    id := fmt.Sprintf("%d%d%d%d%d", currentTime.Month(), currentTime.Day(), currentTime.Hour(), currentTime.Minute, currentTime.Second()) 
    tmpV := fmt.Sprintf("./%s/%s_comments.txt", Dir, id)
    cDir = tmpV
  }
  
  if exist := Controllers.FileExists(sDir); exist == true {
    currentTime := time.Now()
    id := fmt.Sprintf("%d%d%d%d%d", currentTime.Month(), currentTime.Day(), currentTime.Hour(), currentTime.Minute, currentTime.Second())
    
    tmpV := fmt.Sprintf("./%s/%s_sources.txt", Dir, id)
    sDir = tmpV
  }



  var Cmnts string 
  var Srcs string

  switch {
  case err1 != nil && err2 != nil:
    err := errors.New("[i] No sources and no comments found")
    return Command, err
  
  case err1 != nil && err2 == nil :
    for _, FileData := range Command.FileData {
      for _, Comment := range FileData.Comments {
        Cmnts = fmt.Sprintf("%s\n%s", Cmnts, Comment[1])
      }
    }
    fmt.Printf("[i] No sources found")
  case err1 == nil && err2 != nil :
    for _, FileData := range Command.FileData {
      for _, Sources := range FileData.Sources {
        Srcs = fmt.Sprintf("%s\n%s", Srcs, Sources[1])
      }
    }
    fmt.Printf("[i] No conmments found")
  }
  for _, FileData := range Command.FileData {
    for _, Comment := range FileData.Comments {
        Cmnts = fmt.Sprintf("%s\n%s", Cmnts, Comment[1])
    }
    for _, Sources := range FileData.Sources {
        Srcs = fmt.Sprintf("%s\n%s", Srcs, Sources[1])
    }
  }
  err = Controllers.Writer(cDir, []byte(Cmnts))   
  if err != nil {
    fmt.Printf("[i] Error Saving Comments")
  }
  fmt.Printf("\n[i] Comments Saved On %s\n",cDir)
  
  err = Controllers.Writer(sDir, []byte(Srcs))   
  if err != nil {
    fmt.Printf("[i] Error Saving Sources")
  }
  fmt.Printf("\n[i] Sources Saved On %s\n",sDir)

  return Command, nil 
}

func ExtractAllForDir(Command Interfaces.Command) (Interfaces.Command, error){
  
  // Verify Argument
  if Command.Arguments["-d"] == "" {
    err := errors.New("[i] No Dir Specified in Argument")
    return Command, err
  }
  files, err := os.ReadDir(Command.Arguments["-d"])
  if err != nil {
    return Command, err
	}
  for _, file := range files {
		if file.IsDir() {
      continue
		} else {
      path := fmt.Sprintf("%s/%v", Command.Arguments["-d"], file) 
      Command.Arguments["-f"] = path
      Command, err := ExtractAllForFile(Command)
      return Command, err
		}
  }
  return Command, nil 
}
