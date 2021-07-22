package main

import (
  "github.com/ungerik/go3d/vec3"
  "fmt"
  "math"
  "log"
 )

var G float32 =6.6743e-11 // Nm^2/kg^2

type Planet struct {
  Name string
  Mass float32
  Pos vec3.T // in m
  Vel vec3.T // in m/s
}

var planets []Planet = make([]Planet, 0)

func solar() {
  sun := Planet{"sun",
    1e30,
    vec3.T{0,0,0},
    vec3.T{0,0,0},
  }
  earth := Planet{
    "earth",
    6e24,
    vec3.T{1.5e11,0.0},
    vec3.T{0, 29885.7,0},
 }

  planets = append(planets, sun)
  planets = append(planets, earth)
  fmt.Printf("solar %v\n", planets)
}

func speed(radius float32, revtime float32) float32 {
  U := 2*math.Pi * radius
  return U/revtime
}

func timestep(dt float32) {
  for i, p := range(planets) {
    F := vec3.T{0,0,0}
    for j, pp := range(planets) {
      if j==i { continue}
      R := p.Pos.Sub(&(pp.Pos))
      d := R.Length()
      Rn := R.Normalized()
      log.Printf("%s->%s:R=%v,d=%f, Rn=%v",p.Name,pp.Name,
         R, d, Rn)
      log.Printf("m1=%f, m2=%f G=%f\n", p.Mass, pp.Mass, G)
     // m := (p.Mass*pp.Mass)
      m1d := p.Mass / d
      m2d := pp.Mass /d
      strength := G*m1d*m2d
      F.Add(Rn.Scale(strength))
      log.Printf("strength=%f, F=%v\n", strength, F)
   }
   log.Printf("F=%v\n",F)
   a := F.Scale(1. /p.Mass)
   log.Println("a=",a.String())
   dv := a.Scale(dt)
   log.Printf("dv=%v, v(%s) =%v... ", dv, p.Name, p.Vel)
   p.Vel.Add(dv)
   log.Printf("v=%v\n", p.Vel)
 }
  for _, p := range(planets) {
    dr := p.Vel.Scale(dt)
    log.Printf("%s: dr = %v, v=%v, dt=%f\n", p.Name, *dr,
       p.Vel, dt)
    p.Pos.Add(dr)
    log.Println("pos=", p.Pos)
 }
}

func main() {
  solar()
  a := vec3.UnitX
  fmt.Printf("%v\n", a)
  fmt.Printf("%v\n", planets)
  fmt.Printf("speed earth %f\n",
    speed(1.5e11, 365*24*3600))
  for i:=0;i<3;i++ {
    timestep(3600)
    fmt.Printf("%v\n", planets)
 }
}

