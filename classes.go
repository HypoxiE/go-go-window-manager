package main

type Hotkey struct {
	//Modifier  uint16
	//Key       string
	Hotkey    string
	Action    func(Args)
	Arguments Args
}
type Args struct {
	Strings  []string
	Integers []int64
	Floats   []float64
	Booleans []bool
}
