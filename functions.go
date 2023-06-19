package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
    hotelName    = "Gopher Paris Inn"
    totalRooms   = 134
    firstRoomNum = 110
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	roomsOccupied := rand.Intn(totalRooms)

	fmt.Println("Hotel:", hotelName)
	fmt.Println("Number of rooms", totalRooms)

	var availableRooms = vacantRooms(totalRooms, roomsOccupied)
	fmt.Println("Rooms available", availableRooms)

	var occRate = occupancyRate(roomsOccupied, totalRooms)

	var occLevel = checkOccupancyLevel(occRate)
	fmt.Println("                  Occupancy Level:", occLevel)

	fmt.Printf("Occupancy rate : %0.2f \n", occRate)
	fmt.Printf("                  Occupancy Rate: %0.2f %%\n", occRate)

	printRooms(availableRooms)
}

func vacantRooms(totalRooms int, roomsOccupied int) int {
	return totalRooms - roomsOccupied
}
func occupancyRate(roomsOccupied int, totalRooms int) float32 {
	return (float32(roomsOccupied) / float32(totalRooms)) * 100
}

func checkOccupancyLevel(occupancyRate float32) string {
	var occupancyLevel string
	if occupancyRate > 70 {
		occupancyLevel = "High"
	} else if occupancyRate > 20 {
		occupancyLevel = "Medium"
	} else {
		occupancyLevel = "Low"
	}
	return occupancyLevel
}

func printRooms(roomsAvailable int) {
	if roomsAvailable > 0 {
		fmt.Println("Rooms:")
		for i := 0; roomsAvailable > i; i++ {
			roomNumber := firstRoomNum + i
			size := rand.Intn(6) + 1
			nights := rand.Intn(10) + 1
			fmt.Println(roomNumber, ":", size, "people /", nights, " nights ")
		}
	} else {
		fmt.Println("No rooms available for tonight")
	}
}
