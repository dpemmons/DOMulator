package dom

import (
	"testing"
)

func TestNewDOMException(t *testing.T) {
	tests := []struct {
		name            string
		message         string
		exceptionName   string
		expectedName    string
		expectedMessage string
		expectedCode    int
	}{
		{
			name:            "Basic exception with message and name",
			message:         "Test message",
			exceptionName:   "TestError",
			expectedName:    "TestError",
			expectedMessage: "Test message",
			expectedCode:    0, // Unknown name, should return 0
		},
		{
			name:            "Exception with empty name defaults to Error",
			message:         "Test message",
			exceptionName:   "",
			expectedName:    "Error",
			expectedMessage: "Test message",
			expectedCode:    0,
		},
		{
			name:            "Exception with known name returns correct code",
			message:         "Test message",
			exceptionName:   "HierarchyRequestError",
			expectedName:    "HierarchyRequestError",
			expectedMessage: "Test message",
			expectedCode:    HIERARCHY_REQUEST_ERR,
		},
		{
			name:            "Exception with empty message",
			message:         "",
			exceptionName:   "NotFoundError",
			expectedName:    "NotFoundError",
			expectedMessage: "",
			expectedCode:    NOT_FOUND_ERR,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exc := NewDOMException(tt.message, tt.exceptionName)

			if exc.Name() != tt.expectedName {
				t.Errorf("Name() = %v, want %v", exc.Name(), tt.expectedName)
			}

			if exc.Message() != tt.expectedMessage {
				t.Errorf("Message() = %v, want %v", exc.Message(), tt.expectedMessage)
			}

			if exc.Code() != tt.expectedCode {
				t.Errorf("Code() = %v, want %v", exc.Code(), tt.expectedCode)
			}
		})
	}
}

func TestDOMExceptionError(t *testing.T) {
	tests := []struct {
		name          string
		message       string
		exceptionName string
		expectedError string
	}{
		{
			name:          "Error with message",
			message:       "Test message",
			exceptionName: "TestError",
			expectedError: "TestError: Test message",
		},
		{
			name:          "Error without message",
			message:       "",
			exceptionName: "TestError",
			expectedError: "TestError",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exc := NewDOMException(tt.message, tt.exceptionName)
			if exc.Error() != tt.expectedError {
				t.Errorf("Error() = %v, want %v", exc.Error(), tt.expectedError)
			}
		})
	}
}

func TestDOMExceptionCodes(t *testing.T) {
	tests := []struct {
		name          string
		exceptionName string
		expectedCode  int
	}{
		{"IndexSizeError", "IndexSizeError", INDEX_SIZE_ERR},
		{"HierarchyRequestError", "HierarchyRequestError", HIERARCHY_REQUEST_ERR},
		{"WrongDocumentError", "WrongDocumentError", WRONG_DOCUMENT_ERR},
		{"InvalidCharacterError", "InvalidCharacterError", INVALID_CHARACTER_ERR},
		{"NoModificationAllowedError", "NoModificationAllowedError", NO_MODIFICATION_ALLOWED_ERR},
		{"NotFoundError", "NotFoundError", NOT_FOUND_ERR},
		{"NotSupportedError", "NotSupportedError", NOT_SUPPORTED_ERR},
		{"InUseAttributeError", "InUseAttributeError", INUSE_ATTRIBUTE_ERR},
		{"InvalidStateError", "InvalidStateError", INVALID_STATE_ERR},
		{"SyntaxError", "SyntaxError", SYNTAX_ERR},
		{"InvalidModificationError", "InvalidModificationError", INVALID_MODIFICATION_ERR},
		{"NamespaceError", "NamespaceError", NAMESPACE_ERR},
		{"InvalidAccessError", "InvalidAccessError", INVALID_ACCESS_ERR},
		{"SecurityError", "SecurityError", SECURITY_ERR},
		{"NetworkError", "NetworkError", NETWORK_ERR},
		{"AbortError", "AbortError", ABORT_ERR},
		{"URLMismatchError", "URLMismatchError", URL_MISMATCH_ERR},
		{"QuotaExceededError", "QuotaExceededError", QUOTA_EXCEEDED_ERR},
		{"TimeoutError", "TimeoutError", TIMEOUT_ERR},
		{"InvalidNodeTypeError", "InvalidNodeTypeError", INVALID_NODE_TYPE_ERR},
		{"DataCloneError", "DataCloneError", DATA_CLONE_ERR},
		{"Unknown error", "UnknownError", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exc := NewDOMException("test", tt.exceptionName)
			if exc.Code() != tt.expectedCode {
				t.Errorf("Code() = %v, want %v", exc.Code(), tt.expectedCode)
			}
		})
	}
}

func TestConvenienceConstructors(t *testing.T) {
	tests := []struct {
		name         string
		constructor  func(string) *DOMException
		expectedName string
		expectedCode int
		message      string
	}{
		{
			name:         "NewHierarchyRequestError",
			constructor:  NewHierarchyRequestError,
			expectedName: "HierarchyRequestError",
			expectedCode: HIERARCHY_REQUEST_ERR,
			message:      "Test hierarchy error",
		},
		{
			name:         "NewNotFoundError",
			constructor:  NewNotFoundError,
			expectedName: "NotFoundError",
			expectedCode: NOT_FOUND_ERR,
			message:      "Test not found error",
		},
		{
			name:         "NewIndexSizeError",
			constructor:  NewIndexSizeError,
			expectedName: "IndexSizeError",
			expectedCode: INDEX_SIZE_ERR,
			message:      "Test index size error",
		},
		{
			name:         "NewNamespaceError",
			constructor:  NewNamespaceError,
			expectedName: "NamespaceError",
			expectedCode: NAMESPACE_ERR,
			message:      "Test namespace error",
		},
		{
			name:         "NewSyntaxError",
			constructor:  NewSyntaxError,
			expectedName: "SyntaxError",
			expectedCode: SYNTAX_ERR,
			message:      "Test syntax error",
		},
		{
			name:         "NewInvalidCharacterError",
			constructor:  NewInvalidCharacterError,
			expectedName: "InvalidCharacterError",
			expectedCode: INVALID_CHARACTER_ERR,
			message:      "Test invalid character error",
		},
		{
			name:         "NewWrongDocumentError",
			constructor:  NewWrongDocumentError,
			expectedName: "WrongDocumentError",
			expectedCode: WRONG_DOCUMENT_ERR,
			message:      "Test wrong document error",
		},
		{
			name:         "NewNoModificationAllowedError",
			constructor:  NewNoModificationAllowedError,
			expectedName: "NoModificationAllowedError",
			expectedCode: NO_MODIFICATION_ALLOWED_ERR,
			message:      "Test no modification allowed error",
		},
		{
			name:         "NewNotSupportedError",
			constructor:  NewNotSupportedError,
			expectedName: "NotSupportedError",
			expectedCode: NOT_SUPPORTED_ERR,
			message:      "Test not supported error",
		},
		{
			name:         "NewInvalidStateError",
			constructor:  NewInvalidStateError,
			expectedName: "InvalidStateError",
			expectedCode: INVALID_STATE_ERR,
			message:      "Test invalid state error",
		},
		{
			name:         "NewAbortError",
			constructor:  NewAbortError,
			expectedName: "AbortError",
			expectedCode: ABORT_ERR,
			message:      "Test abort error",
		},
		{
			name:         "NewNetworkError",
			constructor:  NewNetworkError,
			expectedName: "NetworkError",
			expectedCode: NETWORK_ERR,
			message:      "Test network error",
		},
		{
			name:         "NewSecurityError",
			constructor:  NewSecurityError,
			expectedName: "SecurityError",
			expectedCode: SECURITY_ERR,
			message:      "Test security error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exc := tt.constructor(tt.message)

			if exc.Name() != tt.expectedName {
				t.Errorf("Name() = %v, want %v", exc.Name(), tt.expectedName)
			}

			if exc.Code() != tt.expectedCode {
				t.Errorf("Code() = %v, want %v", exc.Code(), tt.expectedCode)
			}

			if exc.Message() != tt.message {
				t.Errorf("Message() = %v, want %v", exc.Message(), tt.message)
			}

			expectedError := tt.expectedName + ": " + tt.message
			if exc.Error() != expectedError {
				t.Errorf("Error() = %v, want %v", exc.Error(), expectedError)
			}
		})
	}
}

func TestDOMExceptionConstants(t *testing.T) {
	// Test that all legacy constants have the correct values as per spec
	expectedConstants := map[int]string{
		1:  "INDEX_SIZE_ERR",
		2:  "DOMSTRING_SIZE_ERR",
		3:  "HIERARCHY_REQUEST_ERR",
		4:  "WRONG_DOCUMENT_ERR",
		5:  "INVALID_CHARACTER_ERR",
		6:  "NO_DATA_ALLOWED_ERR",
		7:  "NO_MODIFICATION_ALLOWED_ERR",
		8:  "NOT_FOUND_ERR",
		9:  "NOT_SUPPORTED_ERR",
		10: "INUSE_ATTRIBUTE_ERR",
		11: "INVALID_STATE_ERR",
		12: "SYNTAX_ERR",
		13: "INVALID_MODIFICATION_ERR",
		14: "NAMESPACE_ERR",
		15: "INVALID_ACCESS_ERR",
		16: "VALIDATION_ERR",
		17: "TYPE_MISMATCH_ERR",
		18: "SECURITY_ERR",
		19: "NETWORK_ERR",
		20: "ABORT_ERR",
		21: "URL_MISMATCH_ERR",
		22: "QUOTA_EXCEEDED_ERR",
		23: "TIMEOUT_ERR",
		24: "INVALID_NODE_TYPE_ERR",
		25: "DATA_CLONE_ERR",
	}

	// Verify constant values
	constants := map[int]int{
		INDEX_SIZE_ERR:              1,
		DOMSTRING_SIZE_ERR:          2,
		HIERARCHY_REQUEST_ERR:       3,
		WRONG_DOCUMENT_ERR:          4,
		INVALID_CHARACTER_ERR:       5,
		NO_DATA_ALLOWED_ERR:         6,
		NO_MODIFICATION_ALLOWED_ERR: 7,
		NOT_FOUND_ERR:               8,
		NOT_SUPPORTED_ERR:           9,
		INUSE_ATTRIBUTE_ERR:         10,
		INVALID_STATE_ERR:           11,
		SYNTAX_ERR:                  12,
		INVALID_MODIFICATION_ERR:    13,
		NAMESPACE_ERR:               14,
		INVALID_ACCESS_ERR:          15,
		VALIDATION_ERR:              16,
		TYPE_MISMATCH_ERR:           17,
		SECURITY_ERR:                18,
		NETWORK_ERR:                 19,
		ABORT_ERR:                   20,
		URL_MISMATCH_ERR:            21,
		QUOTA_EXCEEDED_ERR:          22,
		TIMEOUT_ERR:                 23,
		INVALID_NODE_TYPE_ERR:       24,
		DATA_CLONE_ERR:              25,
	}

	for constant, expectedValue := range constants {
		if constant != expectedValue {
			t.Errorf("Constant value mismatch: expected %d, got %d", expectedValue, constant)
		}
	}

	// Verify we have all expected constants
	if len(expectedConstants) != len(constants) {
		t.Errorf("Expected %d constants, but found %d", len(expectedConstants), len(constants))
	}
}
