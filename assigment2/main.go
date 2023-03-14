package main

import (
	c "assigment2/components"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"strconv"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func main() {
	psqlconn := fmt.Sprintf("host = %s port = %d user = %s password = %s dbname = %s sslmode = disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal(err)
	}
	//
	//insertStmt := `insert into "store"("name", "description", "price", "rating") VALUES ('almas','asdas',5.0,6.1)`
	//_, e := db.Exec(insertStmt)
	//if e != nil {
	//	log.Fatal(err)
	//}

	//fmt.Println("successful connection to database")
	//httpHandler := http.FileServer(http.Dir("static"))
	//err = http.ListenAndServe(":8080", httpHandler)
	//if err != nil {
	//	log.Fatal(err)
	//}

	inv := c.NewInventory(db)
	err = inv.CreateUser("almas", "beis")
	if err != nil {
		return
	}
	err = db.Close()
	if err != nil {
		return
	}
	r := mux.NewRouter()
	r.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		// Parse the request parameters
		query := r.URL.Query().Get("query")

		// Search for products
		products, err := SearchProducts(query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Serialize the products to JSON and write the response
		jsonBytes, err := json.Marshal(products)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBytes)
	}).Methods("GET")
	r.HandleFunc("/products/filter", func(w http.ResponseWriter, r *http.Request) {
		// Parse the request parameters
		minPriceStr := r.URL.Query().Get("min_price")
		maxPriceStr := r.URL.Query().Get("max_price")
		minRatingStr := r.URL.Query().Get("min_rating")

		// Convert the request parameters to the appropriate types
		minPrice, err := strconv.ParseFloat(minPriceStr, 64)
		if err != nil {
			http.Error(w, "Invalid min_price parameter", http.StatusBadRequest)
			return
		}
		maxPrice, err := strconv.ParseFloat(maxPriceStr, 64)
		if err != nil {
			http.Error(w, "Invalid max_price parameter", http.StatusBadRequest)
			return
		}
		minRating, err := strconv.Atoi(minRatingStr)
		if err != nil {
			http.Error(w, "Invalid min_rating parameter", http.StatusBadRequest)
			return
		}

		// Filter products
		products, err := FilterProducts(minPrice, maxPrice, minRating)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Serialize the products to JSON and write the response
		jsonBytes, err := json.Marshal(products)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBytes)
	}).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))

	// Registration
	//router.HandleFunc("/registration", func(w http.ResponseWriter, r *http.Request) {
	//	var user c.User
	//	err := json.NewDecoder(r.Body).Decode(&user)
	//	if err != nil {
	//		http.Error(w, err.Error(), http.StatusBadRequest)
	//		return
	//	}
	//
	//	err = inv.CreateUser(&user)
	//	if err != nil {
	//		http.Error(w, err.Error(), http.StatusInternalServerError)
	//		return
	//	}
	//
	//	w.WriteHeader(http.StatusCreated)
	//}).Methods(http.MethodPost)
	//
	//http.Handle("/", router)
	//
	//http.ListenAndServe(":8080", nil)
	//
	//// Authorization
	//router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
	//	username, password, ok := r.BasicAuth()
	//	if !ok {
	//		http.Error(w, "invalid authorization header", http.StatusBadRequest)
	//		return
	//	}
	//
	//	user, err := inv.GetUserByUsername(username)
	//	if err != nil {
	//		http.Error(w, err.Error(), http.StatusUnauthorized)
	//		return
	//	}
	//
	//	if user.Password != password {
	//		http.Error(w, "invalid password", http.StatusUnauthorized)
	//		return
	//	}
	//
	//	w.WriteHeader(http.StatusOK)
	//}).Methods(http.MethodGet)
	//
	//// Searching items based on name
	//router.HandleFunc("/items/search", func(w http.ResponseWriter, r *http.Request) {
	//	name := r.URL.Query().Get("name")
	//	if name == "" {
	//		http.Error(w, "missing name query parameter", http.StatusBadRequest)
	//		return
	//	}
	//
	//	items, err := inv.SearchItemsByName(name)
	//	if err != nil {
	//		http.Error(w, err.Error(), http.StatusInternalServerError)
	//		return
	//	}
	//
	//	json.NewEncoder(w).Encode(items)
	//}).Methods(http.MethodGet)
	//
	//// Filtering items based on price and rating
	//router.HandleFunc("/items/filter", func(w http.ResponseWriter, r *http.Request) {
	//	minPriceStr := r.URL.Query().Get("minPrice")
	//	maxPriceStr := r.URL.Query().Get("maxPrice")
	//	minRatingStr := r.URL.Query().Get("minRating")
	//	maxRatingStr := r.URL.Query().Get("maxRating")
	//
	//	minPrice, err := strconv.ParseFloat(minPriceStr, 64)
	//	if err != nil {
	//		http.Error(w, "invalid minPrice query parameter", http.StatusBadRequest)
	//		return
	//	}
	//
	//	maxPrice, err := strconv.ParseFloat(maxPriceStr, 64)
	//	if err != nil {
	//		http.Error(w, "invalid maxPrice query parameter", http.StatusBadRequest)
	//		return
	//	}
	//
	//	minRating, err := strconv.ParseFloat(minRatingStr, 64)
	//	if err != nil {
	//		http.Error(w, "invalid minRating query parameter", http.StatusBadRequest)
	//		return
	//	}
	//
	//	maxRating, err := strconv.ParseFloat(maxRatingStr, 64)
	//	if err != nil {
	//		http.Error(w, "invalid maxRating query parameter", http.StatusBadRequest)
	//		return
	//	}
	//
	//	items, err := inv.FilterItemsByRatingRange(minRating, maxRating)
	//	if err != nil {
	//		http.Error(w, err.Error(), http.StatusInternalServerError)
	//		return
	//	}
	//	items, err = inv.FilterItemsByPriceRange(minPrice, maxPrice)
	//	if err != nil {
	//		http.Error(w, err.Error(), http.StatusInternalServerError)
	//		return
	//	}
	//
	//	json.NewEncoder(w).Encode(items)
	//}).Methods(http.MethodGet)
	//
	//// Giving rating for an item
	//router.HandleFunc("/items/{id}/rate", func(w http.ResponseWriter, r *http.Request) {
	//	vars := mux.Vars(r)
	//	itemIdStr := vars["id"]
	//	itemId, err := strconv.Atoi(itemIdStr)
	//	if err != nil {
	//		http.Error(w, "invalid item ID", http.StatusBadRequest)
	//		return
	//	}
	//
	//	var rating float64
	//	err = json.NewDecoder(r.Body).Decode(&rating)
	//	if err != nil {
	//		http.Error(w, err.Error(), http.StatusBadRequest)
	//		return
	//	}
	//
	//	err = inv.RateItem(itemId, rating)
	//	if err != nil {
	//		http.Error(w, err.Error(), http.StatusInternalServerError)
	//		return
	//	}
	//
	//	w.WriteHeader(http.StatusOK)
	//}).Methods(http.MethodPut)
	//
	//log.Fatal(http.ListenAndServe(":8080", router))

}
