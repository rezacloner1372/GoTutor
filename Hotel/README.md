# Hotel Room Reservation System

This project is a mini project for reserving rooms in a hotel. It includes functionalities to list rooms, add rooms, reserve rooms, and manage room availability. The project is written in Go and utilizes a Room struct to represent room details such as ID, Type, BedCount, Price, and Status.

## Usage

To use the room reservation system, follow these steps:

1. Run the program by executing the main function.
2. Enter commands as prompted:
   - `1`: List rooms
   - `2`: Add a room
   - `3`: Reserve a room
   - `4`: Remove a room
   - `exit`: Exit the program
3. Follow the instructions provided by the program to list rooms, add new rooms, and reserve rooms for guests.

## Functions

- `GetRoomList()`: Lists all available rooms.
- `ReserveRoom()`: Reserves a room by entering the room ID, number of nights, and person count.
- `AddRoom()`: Adds a new room to the list of available rooms.
- `GetRoom()`: Retrieves a room based on the room ID.
- `CheckPersonCount`: Checks if the number of persons entered is valid for the selected room.
- `RemoveRoom()`: Removes a room from the list based on the room ID.
- `GenerateRooms()`: Generates a list of initial rooms with predefined details.

## Contributing

If you would like to contribute to the development of this project or provide feedback, feel free to make changes to the codebase and submit a pull request.
