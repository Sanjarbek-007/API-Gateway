package casbin

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/casbin/casbin/v2"
	xormadapter "github.com/casbin/xorm-adapter/v2"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5432"
	dbname   = "casbin"
	username = "macbookpro"
	password = "1111"
)

func CasbinEnforcer(logger *slog.Logger) (*casbin.Enforcer, error) {
	// Creating the connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, username, dbname, password)
	
	// Open database connection to ensure it's reachable (optional)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Error("Error connecting to database", "error", err.Error())
		return nil, err
	}
	defer db.Close()

	// Validate the database connection
	err = db.Ping()
	if err != nil {
		logger.Error("Error pinging the database", "error", err.Error())
		return nil, err
	}

	// Initialize the Casbin adapter and enforcer
	adapter, err := xormadapter.NewAdapter("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, username, dbname, password))
	if err != nil {
		logger.Error("Error creating Casbin adapter", "error", err.Error())
		return nil, err
	}

	enforcer, err := casbin.NewEnforcer("casbin/model.conf", adapter)
	if err != nil {
		logger.Error("Error creating Casbin enforcer", "error", err.Error())
		return nil, err
	}

	// Load existing policies
	err = enforcer.LoadPolicy()
	if err != nil {
		logger.Error("Error loading Casbin policy", "error", err.Error())
		return nil, err
	}

	// Add policies
	policies := [][]string{
		{"doctor", "/health/medical_recordsAdd", "POST"},
		{"admin", "/health/medical_recordsAdd", "POST"},
		{"doctor", "/health/medical_recordsGet/:id", "GET"},
		{"admin", "/health/medical_recordsGet/:id", "GET"},
		{"patient", "/health/medical_recordsGet/:id", "GET"},
		{"doctor", "/health/medical_recordsUp", "PUT"},
		{"admin", "/health/medical_recordsUp", "PUT"},
		{"patient", "/health/medical_recordsUp", "PUT"},
		{"admin", "/health/medical_recordsDel/:id", "DELETE"},
		{"admin", "/health/medical_records/user/:userId", "GET"},
		{"patient", "/health/lifestyleAdd", "POST"},
		{"patient", "/api/lifestyle/getLifestyleById/:id", "GET"},
		{"patient", "/health/lifestyleGet/:id", "GET"},
		{"patient", "/health/lifestyleUp", "PUT"},
		{"patient", "/health/lifestyleDel/:id", "DELETE"},
		{"patient", "/health/wearable-dataAdd", "POST"},
		{"patient", "/health/wearabledata/:limit/:page", "GET"},
		{"patient", "/health/wearable-dataGet/:id", "GET"},
		{"patient", "/health/wearable-dataUp", "PUT"},
		{"patient", "/health/wearable-dataDel/:id", "DELETE"},
		{"doctor", "/health/recommendationsAdd", "POST"},
		{"patient", "/health/monitoring/:user_id/realtime", "GET"},
		{"patient", "/health/summary/:user_id/daily/:date", "GET"},
		{"patient", "/health/summary/:user_id/weekly/:start_date", "GET"},
	}

	_, err = enforcer.AddPolicies(policies)
	if err != nil {
		logger.Error("Error adding Casbin policy", "error", err.Error())
		return nil, err
	}

	// Save the policies to the database
	err = enforcer.SavePolicy()
	if err != nil {
		logger.Error("Error saving Casbin policy", "error", err.Error())
		return nil, err
	}

	return enforcer, nil
}
