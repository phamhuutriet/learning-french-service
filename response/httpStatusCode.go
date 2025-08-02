package response

const (
	ErrCodeSuccess      = 20001 // Success
	ErrCodeParamInvalid = 20003 // Invalid email
	ErrCodeUnauthorized = 20004 // Unauthorized
	ErrCodeUserExists   = 20005 // User already exists
)

var msg = map[int]string{
	ErrCodeSuccess:      "Success",
	ErrCodeParamInvalid: "Invalid email",
	ErrCodeUnauthorized: "Unauthorized",
	ErrCodeUserExists:   "User already exists",
}
