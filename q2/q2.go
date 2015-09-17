package main

import ("math")

type Circle struct{
x,y,r float64
}

type Rectangle struct{
x1,y1,x2,y2 float64
}

type Shape interface{
Area() float64
Perimeter() float64
}


func distance(x1, y1, x2, y2 float64) float64 {
a := x2 - x1
b := y2 - y1
return math.Sqrt(a*a + b*b)
}


func (r *Rectangle) Area() float64 {
l := distance(r.x1, r.y1, r.x1, r.y2)
w := distance(r.x1, r.y1, r.x2, r.y1)
return l * w
}

func (r *Rectangle) Perimeter() float64 {
l := distance(r.x1, r.y1, r.x1, r.y2)
w := distance(r.x1, r.y1, r.x2, r.y1)
return (2*l + 2*w)
}


func (c *Circle) Area() float64 {
return math.Pi * c.r*c.r
}

func (c *Circle) Perimeter() float64{
return 2 * math.Pi * c.r
}

