package Interfaces


type Command struct {
    Name     string
    Help string
    LongHelp string
    Arguments []string
    Func func()
    Error error
    FileData []FileData 
}

var Commands = []Command{

  {Name: "ShowCommand", Help: "Show Actuall Command Structure", LongHelp : "Usage: ShowCommand"},
  {Name: "RequestAll", Help: "Request All URLS in File", LongHelp : "Usage: RequestAll ./'urls.txt'"},
  {Name: "ExtractAllForFile", Help: "Extract Comments and Sources by File", LongHelp : "Usage: ExtractAllForFile ./'index.html'"},
  {Name: "ExtractAllForDir", Help: "Extract Comments and Sources by every File in Dir", LongHelp : "Usage: ExtractAllForDir ./Responses"},
  // {Name: "ExtractAll", Help: "Extract Comments and Sources in Argument", LongHelp : "Usage: ExtractAll ./"},
  // {Name: "ExtractSources", Help: "Extract Sources in Argument", LongHelp : "Usage"},
  // {Name: "ExtractComments", Help: "Extract Comments in Argument"},

}

