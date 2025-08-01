package response

const (
	ErrCodeSuccess      = 20001 // Success
	ErrCodeParamInvalid = 20003 // Invalid email
	ErrCodeUnauthorized = 20004 // Unauthorized
)

var msg = map[int]string{
	ErrCodeSuccess:      "Success",
	ErrCodeParamInvalid: "Invalid email",
	ErrCodeUnauthorized: "Unauthorized",
}
