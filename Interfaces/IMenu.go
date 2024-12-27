package Interfaces


type Command struct {
    Name     string
    Help string
    Func func()
    Arguments []string
    Output interface{}
    Error error
    FileData []FileData 
}

var Commands = []Command{

  {Name: "Help", Help: ""},
  {Name: "Exit", Help: ""},
  {Name: "ShowCommand", Help: "Show Actuall Command Structure"},
  {Name: "RequestAll", Help: "Request All URLS in File"},
  {Name: "ExtractAllForFile", Help: "Extract Comments and Sources by File"},
  {Name: "ExtractAllForDir", Help: "Extract Comments and Sources by every File in Dir"},
  {Name: "ExtractAll", Help: "Extract Comments and Sources in Argument"},
  {Name: "ExtractSources", Help: "Extract Sources in Argument"},
  {Name: "ExtractComments", Help: "Extract Comments in Argument"},

}

