package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

type Room struct {
	RoomId string `json:"room_id"`
	Price  string `json:"price"`
	Date   string `json:"date"`
}
type Roomd struct {
	RoomId string `json:"room_id"`
}
type Book struct {
	BookId string `json:"booking_id"`
}
type Bookings struct {
	Booking_id string `json:"booking_id"`
	DateStart  string `json:"data_start"`
	DateEnd    string `json:"data_end"`
}

func create(w http.ResponseWriter, r *http.Request) {
	connStr := "user=postgres password=root dbname=db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var exists bool
	id := r.URL.Query().Get("id")
	dataStart := r.URL.Query().Get("date_start")
	dataEnd := r.URL.Query().Get("date_end")
	_ = db.QueryRow("select exists(select 1 from room where room_id=$1)", id).Scan(&exists)
	if exists {
		_, err = db.Exec("insert into bookings(data_start, data_end, room_id) VALUES($1, $2, $3)", dataStart, dataEnd, id)
		if err != nil {
			log.Fatal(err)
		}
		BId := Book{}
		_ = db.QueryRow("select booking_id from bookings order by booking_id DESC limit 1").Scan(&id)
		BId.BookId = id
		jsonData, _ := json.Marshal(BId)
		fmt.Fprintf(w, "%s", jsonData)
		fmt.Printf("%s", jsonData)
		fmt.Println()
	} else {
		fmt.Fprintf(w, "%s", "Номера не сущесвует")
		fmt.Println("Номера не сущесвует")
	}
}
func list(w http.ResponseWriter, r *http.Request) {
	Id := r.URL.Query().Get("id")
	connStr := "user=postgres password=root dbname=db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("select booking_id, data_start, data_end from bookings where room_id=$1 order by data_start ASC", Id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	Booking := []Bookings{}
	for rows.Next() {
		b := Bookings{}
		err := rows.Scan(&b.Booking_id, &b.DateStart, &b.DateEnd)
		if err != nil {
			log.Fatal(err)
			continue
		}
		Booking = append(Booking, b)
	}
	for idx := range Booking {
		Booking[idx].DateStart = Booking[idx].DateStart[:10]
		Booking[idx].DateEnd = Booking[idx].DateEnd[:10]
	}
	jsonData, err := json.Marshal(&Booking)
	fmt.Fprintf(w, "%s", jsonData)
	fmt.Printf("%s", jsonData)
	fmt.Println()
}
func all(w http.ResponseWriter, r *http.Request) {
	connStr := "user=postgres password=root dbname=db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	date := r.URL.Query().Get("date")
	price := r.URL.Query().Get("price")
	if price == "UP" {
		rows, err := db.Query("select * from room order by price ASC")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		rooms := []Room{}
		for rows.Next() {
			p := Room{}
			err := rows.Scan(&p.RoomId, &p.Price, &p.Date)
			if err != nil {
				log.Fatal(err)
				continue
			}
			rooms = append(rooms, p)
		}
		for idx := range rooms {
			rooms[idx].Date = rooms[idx].Date[:10]
		}
		jsonData, err := json.Marshal(&rooms)
		fmt.Fprintf(w, "%s", jsonData)
		fmt.Printf("%s", jsonData)
		fmt.Println()
	} else if price == "Dn" {
		defer db.Close()
		rows, err := db.Query("select * from room order by price DESC")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		rooms := []Room{}
		for rows.Next() {
			p := Room{}
			err := rows.Scan(&p.RoomId, &p.Price, &p.Date)
			if err != nil {
				log.Fatal(err)
				continue
			}
			rooms = append(rooms, p)
		}
		for idx := range rooms {
			rooms[idx].Date = rooms[idx].Date[:10]
		}
		jsonData, err := json.Marshal(&rooms)
		fmt.Fprintf(w, "%s", jsonData)
		fmt.Printf("%s", jsonData)
		fmt.Println()
	} else if date == "UP" {
		defer db.Close()
		rows, err := db.Query("select * from room order by date ASC")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		rooms := []Room{}
		for rows.Next() {
			p := Room{}
			err := rows.Scan(&p.RoomId, &p.Price, &p.Date)
			if err != nil {
				log.Fatal(err)
				continue
			}
			rooms = append(rooms, p)
		}
		for idx := range rooms {
			rooms[idx].Date = rooms[idx].Date[:10]
		}
		jsonData, err := json.Marshal(&rooms)
		fmt.Fprintf(w, "%s", jsonData)
		fmt.Printf("%s", jsonData)
		fmt.Println()
	} else if date == "Dn" {
		rows, err := db.Query("select * from room order by date DESC")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		rooms := []Room{}
		for rows.Next() {
			p := Room{}
			err := rows.Scan(&p.RoomId, &p.Price, &p.Date)
			if err != nil {
				log.Fatal(err)
				continue
			}
			rooms = append(rooms, p)
		}
		for idx := range rooms {
			rooms[idx].Date = rooms[idx].Date[:10]
		}
		jsonData, err := json.Marshal(&rooms)
		fmt.Fprintf(w, "%s", jsonData)
		fmt.Printf("%s", jsonData)
		fmt.Println()
	}
}
func new(w http.ResponseWriter, r *http.Request) {
	connStr := "user=postgres password=root dbname=db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	dt := time.Now()
	room := r.URL.Query().Get("room")
	price := r.URL.Query().Get("price")
	_, err = db.Exec("insert into room VALUES($1, $2, $3)", room, price, dt.Format("2006-01-02"))
	if err != nil {
		log.Fatal(err)
	}
	id := Roomd{}
	id.RoomId = room
	jsonData, err := json.Marshal(id)
	fmt.Fprintf(w, "%s", jsonData)
	fmt.Printf("%s", jsonData)
	fmt.Println()
}
func delRoom(w http.ResponseWriter, r *http.Request) {
	connStr := "user=postgres password=root dbname=db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	id := r.URL.Query().Get("id")
	_, err = db.Exec("delete from room where room_id=$1", id)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("delete from bookings where room_id=$1", id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Номер " + id + " и его брони удалены")
	w.Write([]byte("Номер " + id + " и его брони удалены"))
}
func deleteBook(w http.ResponseWriter, r *http.Request) {
	connStr := "user=postgres password=root dbname=db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	id := r.URL.Query().Get("id")
	_, err = db.Exec("delete from bookings where booking_id=$1", id)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/new", new)
	mux.HandleFunc("/all", all)
	mux.HandleFunc("/bookings/list", list)
	mux.HandleFunc("/delete", delRoom)
	mux.HandleFunc("/bookings/delete", deleteBook)
	mux.HandleFunc("/bookings/create", create)
	fmt.Println("Server is listening...")
	http.ListenAndServe(":9000", mux)

}
