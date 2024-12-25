package Objects

type Command struct {
    Name     string
    Argument string
    Output interface{}
    Error error
}
