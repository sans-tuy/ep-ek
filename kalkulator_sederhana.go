/*
Created By Anwar Sanusi
*/

package main

import (
   "fmt";
)

var (
   menu int
   m,v,g,h  float64
)

func main() {
   fmt.Println("=================WELCOME================")
   KinetikPoten()
}

//========================EK dan EP==================//
func KinetikPoten(){
   fmt.Println("1. Energi Kinetik")
   fmt.Println("2. Energi Potensial")
   fmt.Println("3. Menu Utama")
   fmt.Print("Pilih salah satu menu diatas ini (ketik nomor menu): ")
   fmt.Scanf("%d", &menu)

   if menu == 1 {
      kinetik()
   }else if menu == 2{
      potensial()
   }else{
      main()
   }
}

func kinetik(){
   fmt.Println("=====Mengukur Energi Kinetik=====")
   fmt.Print("Masukkan nilai massa = ")
   fmt.Scanf("%f", &m)
   fmt.Print("Masukkan nilai kecepatan = ")
   fmt.Scanf("%f", &v)
   kinetik := 0.5*(m*(v*v))
   fmt.Println("Besar Energi Kinetik = \n", kinetik)
   exit()
}

func potensial(){
   fmt.Println("=====Mengukur Energi Potensial=====")
   fmt.Print("Masukkan nilai massa = ")
   fmt.Scanf("%f", &m)
   fmt.Print("Masukkan nilai gravitasi = ")
   fmt.Scanf("%f", &g)
   fmt.Print("Masukkan nilai tinggi = ")
   fmt.Scanf("%f", &h)
   potensial := m*g*h
   fmt.Println("Besar Energi Kinetik = \n", potensial)
   exit()
}

//=========================EXIT=====================//
func exit(){
   fmt.Println("====Selesai===")
}