package Controllers
import (
  "os"
  "bufio"
  "fmt"
  "errors" 
)

func ValidateFileArguments(Args []string) error {
  if len(Args) < 0 {
    err := errors.New("\n [i] No Arguments To eval")
    return err 
  }
  for i, Arg := range Args {
    if (Arg == "-d" ) || (Arg == "-w" ) || (Arg == "-f" ){
      if Args[i+1] == "" {
        err := errors.New("\n[i] No File Specified in Argument")
        return err
      }else if FileExists(Args[i+1]) == false {
        sErr := fmt.Sprintf("\n[i] File For Argument: %s Not Found", Arg) 
        err := errors.New(sErr)
        return err
      } 
    }
  } 
  return nil
}

// Create File or Dir
func CreateFile(name string) error {
  fPath := fmt.Sprintf("./%s", name)
  err := os.Mkdir(fPath, 0755)
  if err != nil {
    return err
  }
  return nil
}

// Eval if DirectoryExists
func DirectoryExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// Eval if FileExists
func FileExists(path string) bool {
	info, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func FContentReader(Path string) (string, error) {
  content, err := os.ReadFile(Path)
  if err != nil {
    return "", err
  }
  return string(content), nil
}

// Read lines of File
func Reader(fPath string) ([]string, error){
  var Lines []string
  file, err := os.Open(fPath)
  if err != nil {
    return Lines, err
  }
  

  defer file.Close()

  fileScanner := bufio.NewScanner(file)
  fileScanner.Split(bufio.ScanLines)

  for fileScanner.Scan() {
      Lines = append(Lines, fileScanner.Text())
  }

  file.Close()
  return Lines, nil
}

func Writer(filePath string, content []byte) error {
  if exist := FileExists(filePath); exist == true {
    serr := fmt.Sprintf("File: %s exist", filePath)
    err := errors.New(serr)
    return err
  }
  file, err := os.Create(filePath)
  if err != nil {
      return err
    }
  if _,err := file.Write(content); err != nil {
     file.Close()
     return err
  }
  file.Close()
  return nil
}

