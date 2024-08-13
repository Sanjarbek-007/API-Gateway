package models

type Success struct {
	Message string `json:"message"`
}

type Error struct {
    Message string `json:"message"`
}

type GetProfileRes struct {
	Email string `json:"email"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	DateOfBirth string `json:"date_of_birth"`
	Gender string `json:"gender"`
	Role string `json:"role"`
}

type Update struct {
	Message string `json:"message"`
}

type UpdateProfileReq struct {
	Email string `json:"email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	DateOfBirth string `json:"date_of_birth"`
	Gender      string `json:"gender"`
}

type GetRealtimeHealthMonitoringRes struct {
    FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	RecommendationType string `json:"recommendation_type"`
	Description        string `json:"description"`
	Priority           int32  `json:"priority"`
}

type GetDailyHealthSummaryReq struct {
	UserId      string `json:"user_id"`
	Date        string `json:"date"`
}

type GetLifeStyle struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	DataType     string `json:"data_type"`
	DataValue    string `json:"data_value"`
	RecordedDate string `json:"recorded_date"`
}

type MedicalReport struct {
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	RecordType   string `json:"record_type"`
	RecordedDate string `json:"recorded_date"`
	Description string `json:"description"`
	DoctorName  string `json:"doctor_name"`
	Attachments   []string `json:"attachments"`
}

type Warable struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	DeviceType string `json:"device_type"`
	DataType string `json:"data_type"`
	DataValue string `json:"data_value"`
	RecordedTimestamp string `json:"recorded_timestamp"`
}