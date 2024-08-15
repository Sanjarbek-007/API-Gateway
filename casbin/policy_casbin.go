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
	host     = "postgres_container"
	port     = "5432"
	dbname   = "casbin"
	username = "macbookpro"
	password = "1111"
)

func CasbinEnforcer(logger *slog.Logger) (*casbin.Enforcer, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, username, dbname, password)
	
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Error("Error connecting to database", "error", err.Error())
		return nil, err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		logger.Error("Error pinging the database", "error", err.Error())
		return nil, err
	}
	query := `DROP TABLE IF EXISTS "casbin_rule";`
	db.Exec(query)

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
		//user
		{"admin", "api/user/profile/:id", "GET"},
		{"admin", "api/user/updateUser/:id", "PUT"},
		{"admin", "api/user/email/:email", "GET"},

		{"patient", "api/user/profile/:id", "GET"},
		{"patient", "api/user/updateUser/:id", "PUT"},
		{"patient", "api/user/email/:email", "GET"},

		{"doctor", "api/user/profile/:id", "GET"},
		{"doctor", "api/user/updateUser/:id", "PUT"},
		{"doctor", "api/user/email/:email", "GET"},

		//health
		{"admin", "/api/health/generate", "POST"},
		{"admin", "/api/health/getRealtimeHealthMonitoring/:user_id", "GET"},
		{"admin", "/api/health/getDailyHealthSummary/:date", "GET"},
        {"admin", "/api/health/getWeeklyHealthSummary/:start_date/:end_date", "GET"},

		{"patient", "/api/health/getRealtimeHealthMonitoring/:user_id", "GET"},
		{"patient", "/api/health/getDailyHealthSummary/:date", "GET"},
        {"patient", "/api/health/getWeeklyHealthSummary/:start_date/:end_date", "GET"},

		{"doctor", "/api/health/generate", "POST"},
		{"doctor", "/api/health/getRealtimeHealthMonitoring/:user_id", "GET"},
		{"doctor", "/api/health/getDailyHealthSummary/:date", "GET"},
        {"doctor", "/api/health/getWeeklyHealthSummary/:start_date/:end_date", "GET"},

		//lifestyle
		{"admin", "/api/lifestyle/addLifestyleData", "POST"},
		{"admin", "/api/lifestyle/getAllLifestyleData", "GET"},
        {"admin", "/api/lifestyle/getLifestyleById/:id", "GET"},
        {"admin", "/api/lifestyle/updateLifestyleData", "PUT"},
		{"admin", "/api/lifestyle/deleteLifestyleData/:id", "DELETE"},

		{"patient", "/api/lifestyle/addLifestyleData", "POST"},
		{"patient", "/api/lifestyle/getAllLifestyleData", "GET"},

		{"doctor", "/api/lifestyle/addLifestyleData", "POST"},
		{"doctor", "/api/lifestyle/getAllLifestyleData", "GET"},
        {"doctor", "/api/lifestyle/getLifestyleById/:id", "GET"},
        {"doctor", "/api/lifestyle/updateLifestyleData", "PUT"},
		{"doctor", "/api/lifestyle/deleteLifestyleData/:id", "DELETE"},
		
	    //medical report
		{"admin", "/api/medicalReport/add", "POST"},
		{"admin", "/api/medicalReport/get", "GET"},
        {"admin", "/api/medicalReport/getById/:id", "GET"},
        {"admin", "/api/medicalReport/update", "PUT"},
        {"admin", "/api/medicalReport/delete/:id", "DELETE"},

		{"patient", "/api/medicalReport/add", "POST"},
		{"patient", "/api/medicalReport/get", "GET"},

		{"doctor", "/api/medicalReport/add", "POST"},
		{"doctor", "/api/medicalReport/get", "GET"},
        {"doctor", "/api/medicalReport/getById/:id", "GET"},
        {"doctor", "/api/medicalReport/update", "PUT"},
        {"doctor", "/api/medicalReport/delete/:id", "DELETE"},

		//wearable
		{"admin", "/api/wearable/add", "POST"},
        {"admin", "/api/wearable/get", "GET"},
        {"admin", "/api/wearable/getById/:id", "GET"},
        {"admin", "/api/wearable/update", "PUT"},
        {"admin", "/api/wearable/delete/:id", "DELETE"},

		{"patient", "/api/wearable/add", "POST"},
        {"patient", "/api/wearable/get", "GET"},

		{"doctor", "/api/wearable/add", "POST"},
        {"doctor", "/api/wearable/get", "GET"},
        {"doctor", "/api/wearable/getById/:id", "GET"},
        {"doctor", "/api/wearable/update", "PUT"},
        {"doctor", "/api/wearable/delete/:id", "DELETE"},

		// notification
		{"admin", "/api/notifications/getAll", "GET"},
		{"admin", "/api/notifications/new", "GET"},

		{"patient", "/api/notifications/getAll", "GET"},
        {"patient", "/api/notifications/new", "GET"},

		{"doctor", "/api/notifications/getAll", "GET"},
        {"doctor", "/api/notifications/new", "GET"},
	}

	_, err = enforcer.AddPolicies(policies)
	if err != nil {
		logger.Error("Error adding Casbin policy", "error", err.Error())
		return nil, err
	}

	err = enforcer.SavePolicy()
	if err != nil {
		logger.Error("Error saving Casbin policy", "error", err.Error())
		return nil, err
	}

	return enforcer, nil
}
