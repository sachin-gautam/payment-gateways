package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var Db *sql.DB

// User represents a user in the system
// @Description Represents a user entity in the system
type User struct {
	ID        int
	Username  string
	Email     string
	CountryID int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Gateway represents a payment gateway in the system
// @Description Represents a payment gateway entity in the system
type Gateway struct {
	ID                  int
	Name                string
	DataFormatSupported string
	CreatedAt           time.Time
	UpdatedAt           time.Time
	Priority            int
	Status              string
}

// Country represents a country in the system
// @Description Represents a country entity in the system
type Country struct {
	ID        int
	Name      string
	Code      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Transaction represents a financial transaction
// @Description Represents a transaction entity in the system
type Transaction struct {
	ID        int
	Amount    string
	Type      string
	Status    string
	UserID    int
	GatewayID int
	CountryID int
	CreatedAt time.Time
}

func InitializeDB(dataSourceName string) error {
	var err error
	Db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		return fmt.Errorf("could not open the database connection: %v", err)
	}

	err = Db.Ping()
	if err != nil {
		return fmt.Errorf("could not ping the database: %v", err)
	}

	log.Println("Successfully connected to the database.")
	return nil
}

func CreateUser(db *sql.DB, user User) error {
	query := `INSERT INTO users (username, email, country_id, created_at, updated_at) 
			  VALUES ($1, $2, $3, $4, $5) RETURNING id`

	err := db.QueryRow(query, user.Username, user.Email, user.CountryID, time.Now(), time.Now()).Scan(&user.ID)
	if err != nil {
		return fmt.Errorf("failed to insert user: %v", err)
	}
	return nil
}

func GetUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query(`SELECT id, username, email, country_id, created_at, updated_at FROM users`)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch users: %v", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID,
			&user.Username,
			&user.Email,
			&user.CountryID,
			&user.CreatedAt,
			&user.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan user: %v", err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserByID(db *sql.DB, userID int) (User, error) {
	var user User

	query := `SELECT id, username, email, country_id, created_at, updated_at 
			  FROM users WHERE id = $1`

	err := db.QueryRow(query, userID).Scan(&user.ID, &user.Username, &user.Email, &user.CountryID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, fmt.Errorf("no user found with id %d", userID)
		}
		return User{}, fmt.Errorf("failed to fetch user: %v", err)
	}

	return user, nil
}

func CreateGateway(db *sql.DB, gateway Gateway) error {
	query := `INSERT INTO gateways (name, data_format_supported, created_at, updated_at) 
			  VALUES ($1, $2, $3, $4) RETURNING id`

	err := db.QueryRow(query, gateway.Name, gateway.DataFormatSupported, time.Now(), time.Now()).Scan(&gateway.ID)
	if err != nil {
		return fmt.Errorf("failed to insert gateway: %v", err)
	}
	return nil
}

func GetGateways(db *sql.DB) ([]Gateway, error) {
	rows, err := db.Query(`SELECT id, name, data_format_supported, created_at, updated_at FROM gateways`)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch gateways: %v", err)
	}
	defer rows.Close()

	var gateways []Gateway
	for rows.Next() {
		var gateway Gateway
		if err := rows.Scan(&gateway.ID,
			&gateway.Name,
			&gateway.DataFormatSupported,
			&gateway.CreatedAt,
			&gateway.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan gateway: %v", err)
		}
		gateways = append(gateways, gateway)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return gateways, nil
}

func CreateCountry(db *sql.DB, country Country) error {
	query := `INSERT INTO countries (name, code, created_at, updated_at) 
			  VALUES ($1, $2, $3, $4) RETURNING id`

	err := db.QueryRow(query, country.Name, country.Code, time.Now(), time.Now()).Scan(&country.ID)
	if err != nil {
		return fmt.Errorf("failed to insert country: %v", err)
	}
	return nil
}

func GetCountries(db *sql.DB) ([]Country, error) {
	rows, err := db.Query(`SELECT id, name, code, created_at, updated_at FROM countries`)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch countries: %v", err)
	}
	defer rows.Close()

	var countries []Country
	for rows.Next() {
		var country Country
		if err := rows.Scan(&country.ID,
			&country.Name,
			&country.Code,
			&country.CreatedAt,
			&country.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan country: %v", err)
		}
		countries = append(countries, country)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return countries, nil
}

func CreateTransaction(db *sql.DB, transaction *Transaction) error {
	query := `INSERT INTO transactions (amount, type, status, gateway_id, country_id, user_id, created_at) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	err := db.QueryRow(query,
		transaction.Amount,
		transaction.Type,
		transaction.Status,
		transaction.GatewayID,
		transaction.CountryID,
		transaction.UserID,
		time.Now()).Scan(&transaction.ID)
	if err != nil {
		return fmt.Errorf("failed to insert transaction: %v", err)
	}
	return nil
}

func GetTransactions(db *sql.DB) ([]Transaction, error) {
	rows, err := db.Query(`SELECT id, amount, type, status, user_id, gateway_id, country_id, created_at FROM transactions`)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch transactions: %v", err)
	}
	defer rows.Close()

	var transactions []Transaction
	for rows.Next() {
		var transaction Transaction
		if err := rows.Scan(&transaction.ID,
			&transaction.Amount,
			&transaction.Type,
			&transaction.Status,
			&transaction.UserID,
			&transaction.GatewayID,
			&transaction.CountryID,
			&transaction.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan transaction: %v", err)
		}
		transactions = append(transactions, transaction)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return transactions, nil
}

func GetSupportedCountriesByGateway(Db *sql.DB, gatewayID int) ([]Country, error) {
	query := `
		SELECT c.id AS country_id, c.name AS country_name
		FROM countries c
		JOIN gateway_countries gc ON c.id = gc.country_id
		WHERE gc.gateway_id = $1
		ORDER BY c.name
	`

	rows, err := Db.Query(query, gatewayID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch countries for gateway %d: %v", gatewayID, err)
	}
	defer rows.Close()

	var countries []Country
	for rows.Next() {
		var country Country
		if err := rows.Scan(&country.ID, &country.Name); err != nil {
			return nil, fmt.Errorf("failed to scan country: %v", err)
		}
		countries = append(countries, country)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate over rows: %v", err)
	}

	return countries, nil
}

func GetAvailableGateways(db *sql.DB, countryID int) ([]Gateway, error) {
	query := `
		SELECT g.id, g.name, g.data_format_supported, g.priority
		FROM gateways g
		JOIN gateway_countries gc ON g.id = gc.gateway_id
		WHERE gc.country_id = $1 AND g.status = 'active'
		ORDER BY g.priority ASC
	`
	rows, err := db.Query(query, countryID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch gateways: %v", err)
	}
	defer rows.Close()

	var gateways []Gateway
	for rows.Next() {
		var gateway Gateway
		if err := rows.Scan(&gateway.ID, &gateway.Name, &gateway.DataFormatSupported, &gateway.Priority); err != nil {
			return nil, fmt.Errorf("failed to scan gateway: %v", err)
		}
		gateways = append(gateways, gateway)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error scanning rows: %v", err)
	}

	return gateways, nil
}

func UpdateTransactionStatus(db *sql.DB, transactionID int, status string) error {
	query := `UPDATE transactions SET status = $1 WHERE id = $2`
	_, err := db.Exec(query, status, transactionID)
	return err
}
