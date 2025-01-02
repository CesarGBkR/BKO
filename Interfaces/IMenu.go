package Interfaces

import (
  "Requester/Objects"
)

type Command struct {
    Name     string
    Help string
    LongHelp string
    Arguments map[string]string
    ArgumentsRaw []string
    Func func()
    Error error
    FileData []FileData
    FUZZConfig FUZZConfig
    Request Request  
    Response Objects.Response
}

var Commands = []Command{

  {Name: "ShowCommand", Help: "Show Actuall Command Structure", LongHelp : "Usage: ShowCommand"},
  {Name: "RequestAll", Help: "Request All URLS in File", LongHelp : "Usage: RequestAll -f ./'urls.txt' -mc 200"},
  {Name: "ExtractAllForFile", Help: "Extract Comments and Sources by File", LongHelp : "Usage: ExtractAllForFile -f ./'index.html'"},
  {Name: "ExtractAllForDir", Help: "Extract Comments and Sources by every File in Dir", LongHelp : "Usage: ExtractAllForDir -d ./Responses"},
  {Name: "FUZZ", Help: "FUZZ Files And Files", LongHelp : "Usage: FUZZ https://www.test.com/FUZZ -w ./Dir -mc 200"},
  // {Name: "ExtractAll", Help: "Extract Comments and Sources in Argument", LongHelp : "Usage: ExtractAll ./"},
  // {Name: "ExtractSources", Help: "Extract Sources in Argument", LongHelp : "Usage"},
  // {Name: "ExtractComments", Help: "Extract Comments in Argument"},

}
