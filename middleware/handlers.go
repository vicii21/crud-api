package middleware

import (
	"crud-api/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func createConnection() *sql.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to postgres")
	return db
}

func getProduct(id int64) (models.Product, error) {
	db := createConnection()
	defer db.Close()
	var product models.Product
	stmt := `SELECT * FROM product WHERE id=$1`

	row := db.QueryRow(stmt, id)
	err := row.Scan(&product.ProductID, &product.Name, &product.ShortDesc, &product.Desc, &product.Price, &product.Quantity, &product.Created, &product.Updated, &product.CategoryID)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return product, nil
	case nil:
		return product, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	return product, err

}

func getAllProduct() ([]models.Product, error) {
	db := createConnection()
	defer db.Close()
	stmt := `SELECT * FROM product`
	var products []models.Product

	rows, err := db.Query(stmt)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	defer rows.Close()
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ProductID, &product.Name, &product.ShortDesc, &product.Desc, &product.Price, &product.Quantity, &product.Created, &product.Updated, &product.CategoryID)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		products = append(products, product)
	}
	return products, err
}

func createProduct(product models.Product) int64 {
	db := createConnection()
	defer db.Close()
	stmt := `INSERT INTO product(name, short_description, description, price, quantity, category_id) VALUES($1, $2, $3, $4, $5, $6) RETURNING id`

	var id int64

	err := db.QueryRow(stmt, product.Name, product.ShortDesc, product.Desc, product.Price, product.Quantity, product.CategoryID).Scan(&id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	fmt.Printf("Inserted a single record (%v). ", id)
	return id
}

func updateProduct(id int64, product models.Product) int64 {
	db := createConnection()
	defer db.Close()
	stmt := `UPDATE product SET name=$2, short_description=$3, description=$4, price=$5, quantity=$6 WHERE id=$1`

	row, err := db.Exec(stmt, id, product.Name, product.ShortDesc, product.Desc, product.Price, product.Quantity)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := row.RowsAffected()
	if err != nil {
		log.Fatalf("Unable to check affected rows. %v", err)
	}

	fmt.Printf("Rows affected: %v", rowsAffected)
	return rowsAffected
}

func deleteProduct(id int64) int64 {
	db := createConnection()
	defer db.Close()
	stmt := `DELETE FROM product WHERE id=$1`

	row, err := db.Exec(stmt, id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := row.RowsAffected()
	if err != nil {
		log.Fatalf("Unable to check the rows affected. %v", err)
	}

	fmt.Printf("Rows affected: %v", rowsAffected)
	return rowsAffected
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert string to int. %v", err)
	}

	product, err := getProduct(int64(id))
	if err != nil {
		log.Fatalf("Unable to get product. %v", err)
	}

	json.NewEncoder(w).Encode(product)
}

func GetAllProduct(w http.ResponseWriter, r *http.Request) {
	products, err := getAllProduct()
	if err != nil {
		log.Fatalf("Unable to get all products. %v", err)
	}
	json.NewEncoder(w).Encode(products)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Fatalf("Unable to decode body of the request. %v", err)
	}

	createID := createProduct(product)

	res := response{
		ID:      createID,
		Message: "Product successfully added.",
	}

	json.NewEncoder(w).Encode(res)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert string to int. %v", err)
	}

	var product models.Product
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Fatalf("Unable to decode body of the request. %v", err)
	}

	updatedRows := updateProduct(int64(id), product)
	msg := fmt.Sprintf("Product successfully updated. Rows affected: %v", updatedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert string to int. %v", err)
	}

	deletedRows := deleteProduct(int64(id))
	msg := fmt.Sprintf("Product successfully deleted. Rows affected: %v", deletedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func getCategory(id int64) (models.Category, error) {
	db := createConnection()
	defer db.Close()
	stmt := `SELECT * FROM category WHERE id=$1`

	var category models.Category
	row := db.QueryRow(stmt, id)

	err := row.Scan(&category.CategoryID, &category.Name, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return category, err
	case nil:
		return category, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}
	return category, err
}

func getAllCategory() ([]models.Category, error) {
	db := createConnection()
	defer db.Close()
	stmt := `SELECT * FROM category`

	var categories []models.Category
	rows, err := db.Query(stmt)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	defer rows.Close()
	for rows.Next() {
		var category models.Category

		err := rows.Scan(&category.CategoryID, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		categories = append(categories, category)
	}
	return categories, err
}

func createCategory(category models.Category) int64 {
	db := createConnection()
	defer db.Close()
	stmt := `INSERT INTO category(category_name) VALUES($1) RETURNING id`

	var id int64

	err := db.QueryRow(stmt, category.Name).Scan(&id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record. %v", id)
	return id
}

func updateCategory(id int64, category models.Category) int64 {
	db := createConnection()
	defer db.Close()
	stmt := `UPDATE category SET category_name=$2 WHERE id=$1`

	result, err := db.Exec(stmt, id, category.Name)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatalf("Unable to check the rows affected. %v", err)
	}

	fmt.Printf("Rows affected: %v", rowsAffected)
	return rowsAffected
}

func deleteCategory(id int64) int64 {
	db := createConnection()
	defer db.Close()
	stmt := `DELETE FROM category WHERE id=$1`

	result, err := db.Exec(stmt, id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatalf("Unable to check the rows affected. ")
	}

	fmt.Printf("Rows affected: %v", rowsAffected)
	return rowsAffected
}

func GetCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert string to int. %v", err)
	}

	category, err := getCategory(int64(id))
	if err != nil {
		log.Fatalf("Unable to get category. %v", err)
	}

	json.NewEncoder(w).Encode(category)
}

func GetAllCategory(w http.ResponseWriter, r *http.Request) {
	categories, err := getAllCategory()
	if err != nil {
		log.Fatalf("Unable to get all categories. %v", err)
	}

	json.NewEncoder(w).Encode(categories)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category

	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		log.Fatalf("Unable to decode body of the request. %v", err)
	}

	createID := createCategory(category)

	res := response{
		ID:      createID,
		Message: "Category created successfully.",
	}

	json.NewEncoder(w).Encode(res)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert string to int.")
	}

	var category models.Category
	err = json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		log.Fatalf("Unable to decode body of the request. %v", err)
	}
	updatedRows := updateCategory(int64(id), category)
	msg := fmt.Sprintf("Product updated successfully. Rows affected: %v", updatedRows)
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert string to int. %v", err)
	}

	var category models.Category
	err = json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		log.Fatalf("Unable to decode body of the request. %v", err)
	}

	deletedRows := deleteCategory(int64(id))
	msg := fmt.Sprintf("Product deleted successfully. Rows affected: %v", deletedRows)
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}
