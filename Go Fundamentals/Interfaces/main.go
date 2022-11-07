package main

import "fmt"



type Baseball struct {
	mass int
	acceleration int
}

type Football  struct{

}

func  (b Baseball)  getForce()int {
	  return b.mass * b.acceleration
}


func (f Football) getForce()int{
	return 50
}

type Force interface {
	getForce()int 
}

func  compareForce(a,b Force) bool {
	return a.getForce() > b.getForce()
}

type Rectangle struct{
	length int
	 width int
}

type Square struct {
	side int
}

func (r Rectangle) getArea()int{
	return r.length * r.width
}

func (s Square) getArea()int{
	return s.side* s.side
}

type Area interface{
	getArea()int
}

func  compareArea(a, b Area) bool{
   return a.getArea() > b.getArea()
}


func main(){
  

	b1 := Baseball{mass:2, acceleration: 20}
	f := Football{}

	fmt.Println(compareForce(b1,f))


	//Exercise
	r := Rectangle{length: 2, width: 4}
	s:= Square{side: 2}
    

	fmt.Println(compareArea(r, s))

}