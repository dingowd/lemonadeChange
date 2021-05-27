package main

import "fmt"

func sortDecrease(mass []int, l int) {
  var change int
  for i:=0; i<(l-1); i++{
    for j:=0; j<(l-i-1); j++ {
      if mass[j]<mass[j+1] {
        change=mass[j]
        mass[j]=mass[j+1]
        mass[j+1]=change
      }
    }
  }
}

func lemonadeChange(bills []int) bool {
  var summ int
  flag:=false
  for i:=1; i<len(bills); i++ {
    sortDecrease(bills, i)
    if bills[i]!=5 {
      innerLoop:
      for j:=0; j<i; j++ {
       summ=summ+bills[j]
       switch {
         case summ==(bills[i]-5):
         flag=true
         bills[j]=0
         break innerLoop
         case summ<(bills[i]-5):
         flag=false
         bills[j]=0
         case summ>(bills[i]-5):
         summ=summ-bills[j]
         flag=false
       }
      }
    }
    summ=0
  }
  return flag
}

func main() {
  var number int
  fmt.Print("Сколько человек стоит в очереди?")
  fmt.Scan(&number)
  fmt.Println("Введите номиналы купюр у покупателей в порядке очереди")
  bills:=make([]int,number)
  for i:=0; i<number; i++ {
    fmt.Scan(&bills[i])
    for bills[i]!=5 && bills[i]!=10 && bills[i]!=20 {
      fmt.Println("Номинал купюры может быть равным 5, 10 или 20, введите заново")
      fmt.Scan(&bills[i])
    }
  }
  if lemonadeChange(bills) {
    fmt.Println("Можно сдать сдачу всем покупателям")
  } else {
    fmt.Println("Нельзя сдать сдачу всем покупателям")
  }
}