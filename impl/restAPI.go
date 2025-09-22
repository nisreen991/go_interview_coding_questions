type Book struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("COntent-Type", "application/json")
	var book Book 
	_ = json.NewEncoder(r.Body).Decode(&book)
	books = append(books,book)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := gin.Vars(r)
	
}