package main

import "fmt"

type Room struct {
	ID       int
	Type     string
	Status   bool // true for reserved, false for available
	BedCount int
	Price    int
}

var Rooms []Room = GenerateRooms() // خروجی فانکنش GenerateRooms

func main() {
	input := ""
	for input != "exit" {
		fmt.Println("Please select one of the following options:")
		fmt.Println("1: Room list")
		fmt.Println("2: Add room")
		fmt.Println("3: Reserve room")
		fmt.Println("4: Remove room")
		fmt.Scanln(&input)
		switch input {
		case "1":
			GetRoomList()
		case "2":
			AddRoom()
		case "3":
			ReserveRoom()
		case "4":
			RemoveRoom()
		case "exit":
			fmt.Println("Exiting...")
			break
		default:
			fmt.Println("Invalid input")
		}
	}
}

func GetRoomList() {
	for _, room := range Rooms {
		fmt.Printf("%+v\n", room)
	}
}

func GetRoomFromInput() Room {
	var room Room = Room{Status: false}
	fmt.Println("Enter room information line by line(ID, Type, BedCount, Price):")
	fmt.Scanln(&room.ID, &room.Type, &room.BedCount, &room.Price)

	return room
}

func AddRoom() {
	room := GetRoomFromInput()
	Rooms = append(Rooms, room)
	fmt.Printf("Room added successfully: %+v\n", room)

}

func RemoveRoom() {
	id := 0
	fmt.Println("Enter room ID for removal:")
	fmt.Scanln(&id)
	for i := 0; i < len(Rooms); i++ {
		if Rooms[i].ID == id {
			Rooms = append(Rooms[:i], Rooms[i+1:]...)
			return
		}
	}
}

func ReserveRoom() {
	id := 0
	nights := 0
	personCount := 0
	fmt.Println("Enter room ID for reservation:")
	fmt.Scanln(&id)
	room := GetRoom(id)
	if room == nil {
		fmt.Println("Room not found")
		return

	}
	if room.Status == true {
		fmt.Println("Room already reserved")
		return
	}

	fmt.Println("Enter number of nights and personCount line by line:")
	fmt.Scanln(&nights)
	fmt.Scanln(&personCount)
	roomPrice, tax, discountAmount, fincalPrice := CalculateRoomPrice(*room, nights, personCount)
	room.Status = true
	fmt.Printf("Room Price: %f, tax: %f, discountAmount: %f, finalPrice: %f\n", roomPrice, tax, discountAmount, fincalPrice)

}
func GetRoom(id int) *Room {
	for i := 0; i < len(Rooms); i++ {
		if Rooms[i].ID == id {
			return &Rooms[i]
		}
	}
	return nil
}

func CalculateRoomPrice(room Room, nights int, personCount int) (roomPrice float64, tax float64, discountAmount float64, fincalPrice float64) {
	discountPercentage := 0.0
	if (nights >= 7) && (nights <= 15) {
		discountPercentage = 0.1
	} else if (nights >= 16) && (nights <= 30) {
		discountPercentage = 0.15
	} else if nights > 31 {
		discountPercentage = 0.2
	}

	switch room.Type {
	case "Single":
		roomPrice = float64(nights*room.Price*personCount) * 1
	case "Double":
		roomPrice = float64(nights*room.Price*personCount) * 1.2
	case "Standard":
		roomPrice = float64(nights*room.Price*personCount) * 1.4
	case "Suite":
		roomPrice = float64(nights*room.Price*personCount) * 1.6
	}
	tax = float64(roomPrice) * 0.09
	discountAmount = roomPrice * discountPercentage
	fincalPrice = roomPrice + tax - discountAmount
	return

}

func GenerateRooms() []Room { // حروجی یک اسلایس از Room است
	rooms := []Room{}
	rooms = append(rooms, Room{ID: 1, Type: "Single", Status: false, BedCount: 1, Price: 100})
	rooms = append(rooms, Room{ID: 2, Type: "Single", Status: false, BedCount: 1, Price: 120})
	rooms = append(rooms, Room{ID: 3, Type: "Single", Status: false, BedCount: 1, Price: 150})
	rooms = append(rooms, Room{ID: 4, Type: "Double", Status: false, BedCount: 2, Price: 200})
	rooms = append(rooms, Room{ID: 5, Type: "Double", Status: false, BedCount: 2, Price: 220})
	rooms = append(rooms, Room{ID: 6, Type: "Double", Status: false, BedCount: 3, Price: 250})
	rooms = append(rooms, Room{ID: 7, Type: "Double", Status: false, BedCount: 3, Price: 255})
	rooms = append(rooms, Room{ID: 8, Type: "Double", Status: false, BedCount: 3, Price: 265})
	rooms = append(rooms, Room{ID: 9, Type: "Double", Status: false, BedCount: 3, Price: 280})
	rooms = append(rooms, Room{ID: 10, Type: "Standard", Status: false, BedCount: 4, Price: 300})
	rooms = append(rooms, Room{ID: 11, Type: "Standard", Status: false, BedCount: 4, Price: 320})
	rooms = append(rooms, Room{ID: 12, Type: "Standard", Status: false, BedCount: 4, Price: 340})

	return rooms
}
