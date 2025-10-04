package utils

const (
	// General
	MsgSuccess                = "Success"
	MsgOperationCompleted     = "Operation completed successfully"
	MsgFailed                 = "Failed"
	MsgMissingMandatoryParams = "Missing Mandatory Parameters"
	MsgInvalidInput           = "Invalid user input"
	MsgValidationFailed       = "Validation failed"
	MsgInternalError          = "Internal server error"
	MsgBadRequest             = "Bad request"
	MsgUnauthorized           = "Unauthorized access"
	MsgForbidden              = "Access forbidden"
	MsgNotFound               = "Resource not found"
	MsgConflict               = "Conflict detected"
	MsgServiceUnavailable     = "Service temporarily unavailable"
	MsgMappingError           = "Mapping error"
	MsgParsingError           = "Parging error"

	// CRUD
	MsgCreateSuccess = "Successfully created"
	MsgUpdateSuccess = "Successfully updated"
	MsgDeleteSuccess = "Successfully deleted"

	// Auth
	MsgInvalidCredentials = "Invalid username or password"
	MsgTokenExpired       = "Token has expired"
	MsgTokenInvalid       = "Invalid or malformed token"
	MsgAccessDenied       = "Access denied"
	MsgLoginSuccess       = "Login successful"
	MsgLogoutSuccess      = "Logout successful"
	MsgAccountLocked      = "Account locked"
	MsgSessionExpired     = "Session expired"

	// Database
	MsgDatabaseError     = "Error in database"
	MsgRecordNotFound    = "Record not found"
	MsgDuplicateRecord   = "Duplicate record"
	MsgDataFetchError    = "Failed to fetch data"
	MsgDataSaveError     = "Failed to save data"
	MsgDataUpdateError   = "Failed to update data"
	MsgDataDeleteError   = "Failed to delete data"
	MsgTransactionFailed = "Transaction failed"
	MsgConnectionError   = "Database connection error"

	// File Upload
	MsgFileUploadSuccess = "File uploaded successfully"
	MsgFileUploadFailed  = "File upload failed"
	MsgFileNotFound      = "File not found"
	MsgInvalidFileFormat = "Invalid file format"
	MsgFileTooLarge      = "File size exceeds limit"

	// External Services
	MsgExternalAPIFailed = "External API request failed"
	MsgTimeout           = "Request timed out"
	MsgNetworkError      = "Network error occurred"
	MsgDependencyError   = "Dependency service error"
	MsgThirdPartyError   = "Third-party service error"

	//Pagination
	MsgNoMoreRecords     = "No more records available"
	MsgInvalidPagination = "Invalid pagination parameters"
)
