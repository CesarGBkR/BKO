package Domain 

import (
  "fmt"
  "io/ioutil"
  "regexp"
  "strings"
  "errors"

  "Requester/Interfaces"
  "Requester/Controllers"
)

func Reader(fPath string) ([]string, error){
  res, err := Controllers.Reader(fPath)
  return res, err
}

// Write File
func ResWriter(Request Interfaces.Request) (error){
  resBody, err := ioutil.ReadAll(Request.Response.Body)
  if err != nil {
    return err
  }  
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
  err = Controllers.Writer(fileDir, resBody)
  if err != nil {
    return err
  }
  return nil
}
