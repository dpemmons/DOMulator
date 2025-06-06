package testing

import (
	"fmt"
	"strings"
	"testing"

	"github.com/dpemmons/DOMulator/internal/js"
)

// ConsoleCapture provides a testing helper to capture and assert on console output
type ConsoleCapture struct {
	messages []ConsoleMessage
	t        testing.TB // Store test context for assertions
}

// ConsoleMessage represents a single console message
type ConsoleMessage struct {
	Level js.ConsoleLevel
	Args  []interface{}
}

// NewConsoleCapture creates a new console capture instance
func NewConsoleCapture(t testing.TB) *ConsoleCapture {
	return &ConsoleCapture{
		messages: make([]ConsoleMessage, 0),
		t:        t,
	}
}

// Callback is the callback function for the runtime console
func (c *ConsoleCapture) Callback(level js.ConsoleLevel, args []interface{}) {
	c.messages = append(c.messages, ConsoleMessage{Level: level, Args: args})
}

// AssertLogContains checks that at least one log message contains the expected text
func (c *ConsoleCapture) AssertLogContains(expected string) *ConsoleCapture {
	c.assertContains(js.ConsoleLog, expected)
	return c
}

// AssertErrorContains checks that at least one error message contains the expected text
func (c *ConsoleCapture) AssertErrorContains(expected string) *ConsoleCapture {
	c.assertContains(js.ConsoleError, expected)
	return c
}

// AssertWarnContains checks that at least one warn message contains the expected text
func (c *ConsoleCapture) AssertWarnContains(expected string) *ConsoleCapture {
	c.assertContains(js.ConsoleWarn, expected)
	return c
}

// AssertInfoContains checks that at least one info message contains the expected text
func (c *ConsoleCapture) AssertInfoContains(expected string) *ConsoleCapture {
	c.assertContains(js.ConsoleInfo, expected)
	return c
}

// AssertLogCount checks the number of log messages
func (c *ConsoleCapture) AssertLogCount(expected int) *ConsoleCapture {
	actual := c.countLevel(js.ConsoleLog)
	if actual != expected {
		c.t.Errorf("Expected %d log messages, got %d", expected, actual)
	}
	return c
}

// AssertErrorCount checks the number of error messages
func (c *ConsoleCapture) AssertErrorCount(expected int) *ConsoleCapture {
	actual := c.countLevel(js.ConsoleError)
	if actual != expected {
		c.t.Errorf("Expected %d error messages, got %d", expected, actual)
	}
	return c
}

// AssertWarnCount checks the number of warn messages
func (c *ConsoleCapture) AssertWarnCount(expected int) *ConsoleCapture {
	actual := c.countLevel(js.ConsoleWarn)
	if actual != expected {
		c.t.Errorf("Expected %d warn messages, got %d", expected, actual)
	}
	return c
}

// AssertInfoCount checks the number of info messages
func (c *ConsoleCapture) AssertInfoCount(expected int) *ConsoleCapture {
	actual := c.countLevel(js.ConsoleInfo)
	if actual != expected {
		c.t.Errorf("Expected %d info messages, got %d", expected, actual)
	}
	return c
}

// AssertNoErrors checks that no error messages were logged
func (c *ConsoleCapture) AssertNoErrors() *ConsoleCapture {
	if count := c.countLevel(js.ConsoleError); count > 0 {
		c.t.Errorf("Expected no errors, but found %d", count)
		// Show the errors for debugging
		for _, msg := range c.messages {
			if msg.Level == js.ConsoleError {
				c.t.Errorf("  Error: %v", msg.Args)
			}
		}
	}
	return c
}

// AssertNoWarnings checks that no warning messages were logged
func (c *ConsoleCapture) AssertNoWarnings() *ConsoleCapture {
	if count := c.countLevel(js.ConsoleWarn); count > 0 {
		c.t.Errorf("Expected no warnings, but found %d", count)
		// Show the warnings for debugging
		for _, msg := range c.messages {
			if msg.Level == js.ConsoleWarn {
				c.t.Errorf("  Warning: %v", msg.Args)
			}
		}
	}
	return c
}

// GetLogs returns all log messages
func (c *ConsoleCapture) GetLogs() [][]interface{} {
	return c.getMessages(js.ConsoleLog)
}

// GetErrors returns all error messages
func (c *ConsoleCapture) GetErrors() [][]interface{} {
	return c.getMessages(js.ConsoleError)
}

// GetWarnings returns all warning messages
func (c *ConsoleCapture) GetWarnings() [][]interface{} {
	return c.getMessages(js.ConsoleWarn)
}

// GetInfos returns all info messages
func (c *ConsoleCapture) GetInfos() [][]interface{} {
	return c.getMessages(js.ConsoleInfo)
}

// GetAllMessages returns all captured messages
func (c *ConsoleCapture) GetAllMessages() []ConsoleMessage {
	return c.messages
}

// Clear removes all captured messages
func (c *ConsoleCapture) Clear() *ConsoleCapture {
	c.messages = c.messages[:0]
	return c
}

// Count returns the total number of messages at the specified level
func (c *ConsoleCapture) Count(level js.ConsoleLevel) int {
	return c.countLevel(level)
}

// assertContains checks that at least one message of the given level contains the expected text
func (c *ConsoleCapture) assertContains(level js.ConsoleLevel, expected string) {
	found := false
	for _, msg := range c.messages {
		if msg.Level == level {
			// Check each argument for the expected text
			for _, arg := range msg.Args {
				argStr := c.argToString(arg)
				if strings.Contains(strings.ToLower(argStr), strings.ToLower(expected)) {
					found = true
					break
				}
			}
			if found {
				break
			}
		}
	}

	if !found {
		c.t.Errorf("Expected %s message containing '%s', but none found", level, expected)
		// Show available messages for debugging
		messages := c.getMessages(level)
		if len(messages) > 0 {
			c.t.Errorf("Available %s messages:", level)
			for i, msg := range messages {
				c.t.Errorf("  %d: %v", i+1, msg)
			}
		}
	}
}

// argToString converts a console argument to its string representation
func (c *ConsoleCapture) argToString(arg interface{}) string {
	if arg == nil {
		return "null"
	}

	switch v := arg.(type) {
	case string:
		return v
	case int, int32, int64, uint, uint32, uint64:
		return fmt.Sprintf("%v", v)
	case float32, float64:
		return fmt.Sprintf("%v", v)
	case bool:
		if v {
			return "true"
		}
		return "false"
	default:
		return fmt.Sprintf("%v", v)
	}
}

// countLevel counts messages at the specified level
func (c *ConsoleCapture) countLevel(level js.ConsoleLevel) int {
	count := 0
	for _, msg := range c.messages {
		if msg.Level == level {
			count++
		}
	}
	return count
}

// getMessages returns all messages at the specified level
func (c *ConsoleCapture) getMessages(level js.ConsoleLevel) [][]interface{} {
	var result [][]interface{}
	for _, msg := range c.messages {
		if msg.Level == level {
			result = append(result, msg.Args)
		}
	}
	return result
}
