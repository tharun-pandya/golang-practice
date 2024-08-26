package main

import (
	"fmt"
	"slices"
)

type appitude interface {
	speed()
}

type train struct {
	len            float64
	tspeed         float64
	distance       float64
	object_len     float64
	object_speed_s float64
	object_speed_o float64
	time           float64
}

func (t train) speed() {

}

func main() {
	var ch, op, d int
	var app appitude
	t := train{}
	var ch1 []int
	fmt.Print("1.train\n2.boat\n3.vechile\n enter choice")
	fmt.Scan(&ch)
	switch ch {
	case 1:
		fmt.Println("Enter your requirements calculating train speed")
		fmt.Print("1.length of train\n2.distance\n3.object length\n4.object speed at same direction\n5.object speed at opposite direction\n6.speed of train\n7.time to over take\n")
		fmt.Println("enter how many data you want to enter ")
		fmt.Scan(&op)
		for i := 0; i < op; i++ {
			fmt.Scan(&d)
			ch1 = append(ch1, d)
		}
		if slices.Equal(ch1, []int{3, 6, 7}) {
			fmt.Println("Enter train speed in minutes")
			fmt.Scan(&t.tspeed)
			fmt.Println("Enter object length in meters")
			fmt.Scan(&t.object_len)
			fmt.Println("Enter over take time in munites")
			fmt.Scan(t.time)
			app = t
			app.speed()
		} else if slices.Equal(ch1, []int{1, 4, 7}) {
			fmt.Println("Enter train length in meters")
			fmt.Scan(&t.len)
			fmt.Println("Enter object speed at same direction in minutes")
			fmt.Scan(&t.object_speed_s)
			fmt.Println("Enter time to overtake that object in minutes")
			fmt.Scan(&t.time)
			app = t
			app.speed()
		} else if slices.Equal(ch1, []int{1, 6, 4}) {
			fmt.Println("Enter train length in meters")
			fmt.Scan(&t.len)
			fmt.Println("Enter train speed in minutes")
			fmt.Scan(&t.tspeed)
			fmt.Println("Enter object speed at same direction in minutes")
			fmt.Scan(&t.object_speed_s)
			app = t
			app.speed()
		}
		break
	default:
		fmt.Println("invalid choice")
		break
	}

}
