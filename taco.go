package main

var Taco = `
       ╭╯╭╯╭╯
     _________
    ~&@~o^^^~^o\
   o~/          |
  ~~/            |
 /o/_______-------

`

var Taco2 = `
      ╰╮╰╮╰╮
     _________
    ~&@~o^^^~^o\
   o~/          |
  ~~/            |
 /o/_______-------

`

var lastTaco = false

func NextTaco() string {
	lastTaco = !lastTaco
	if lastTaco {
		return Taco
	}
	return Taco2
}
