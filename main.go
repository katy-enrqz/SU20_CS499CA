package main

import ("fmt"
        
)
type rectangle struct {
  width, height float64
    
}

func (r * rectangle) vals(){
  if r.width > 0 && r.width < 20.0{
      r.width = r.width
  } else {
    r.width = 1.0
  }

  if r.height > 0 && r.height < 20.0{
      r.height = r.height
  } else {
    r.height = 1.0
  }

}


//area
func (r * rectangle) area () float64{
  return r.width * r.height
} 

//perimeter
func (r rectangle) perimeter() float64{
  return 2*r.width + 2*r.height
}



func main() {
  //new instance for rectangle
  a := new (rectangle)
  a.vals()
  fmt.Println("a: length = ", a.height, " width = ", a.width, " perimeter = ", a.perimeter(), " area = ", a.area())

  b := rectangle{width: 4.0, height: 5.0}
  b.vals()
  fmt.Println("b: length = ", b.height, " width = ", b.width, " perimeter = ", b.perimeter(), " area = ", b.area())

  c := rectangle{width: 67.0, height: 888.0}
  c.vals()
  fmt.Println("c: length = ", c.height, " width = ", c.width, " perimeter = ", c.perimeter(), " area = ", c.area())

  


}