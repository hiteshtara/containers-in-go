type Ball struct {
	Radius   int
	Material string
}
type Football struct {
	Ball
}

fb := Football{}
fmt.Printf("fb = %+v\n", fb)

fb := Football{Ball{Radius: 5, Material: "leather"}}
fmt.Printf("fb = %+v\n", fb)

func (b Ball) Bounce() {
    fmt.Printf("Bouncing ball %+v\n", b)
}

fb.Bounce()
type Bouncer interface {
    Bounce()
}

func BounceIt(b Bouncer) {
    b.Bounce()
}

type Football struct {
    *Ball
}

type Football struct {
    Bouncer
}

fb := Football{&Ball{Radius: 5, Material: "leather"}}
fb.Bounce()
//Serialize Empty Structs to JSON in Go

type MyStruct struct {
	Data   MyData `json:"data,omitempty"`
	Status string `json:"status,omitempty"`
  }
  str := &MyStruct{
	Status: "some-status"
  }
  
  j, _ := json.Marshal(str)
  
  Println(string(j))
// whic results in {
  "data": {},
  "status": "some-status"
}//
type MyStruct struct {
	Data   *MyData `json:"data,omitempty"`
	Status string `json:"status,omitempty"`
  }
  To yield JSON without a "data":
  str := &MyStruct{
	Status: "some-status"
  }  

  package main
  
 type Jedi interface {
	 HasForce() bool
 }
  
 type Knight struct {
 }
  
 func (k *Knight) HasForce() bool {
	 return true
 }
  
 var _ Jedi = (*Knight)(nil)
  
 func main() {
 }