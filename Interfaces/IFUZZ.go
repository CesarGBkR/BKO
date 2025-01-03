package Interfaces

type FUZZConfig struct {
  Threats int
  Type string
  Wordlist []string
  Match bool
  Filter bool
  Filters Filters
  Matchs Matchs
}

type Filters struct {
  FCodes []int
  FLengths []int 
}

type Matchs struct {
  MCodes []int
  MLengths []int 
}
