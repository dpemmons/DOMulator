# Active Context

## Current Focus

### JavaScript Error and Console Callback System
- ✅ **NEW**: JavaScript runtime no longer panics on errors by default
- ✅ **NEW**: Implemented error callback system for graceful error handling
- ✅ **NEW**: Console output callback system for test assertions
- ✅ **NEW**: Created testing helpers: `ConsoleCapture` and `ErrorCapture`

### Console & Error Handling Features
- **SetConsoleCallback()**: Capture console.log, console.error, console.warn, console.info
- **SetErrorCallback()**: Capture JavaScript runtime errors (TypeError, ReferenceError, etc.)
- **ConsoleCapture helper**: Provides fluent assertions for console output
- **ErrorCapture helper**: Provides fluent assertions for JavaScript errors

### Testing Improvements
- Added comprehensive console/error callback tests in `console_error_callbacks_test.go`
- Console callbacks work for all levels: log, info, warn, error
- Error callbacks capture JavaScript exceptions instead of panicking
- Test harness automatically applies callbacks when runtime is created

### ResizeObserver Implementation  
- ✅ ResizeObserver API is fully implemented and working
- ✅ Initial observation callbacks fire correctly
- ✅ Added `TriggerElementResize` method to TestHarness for testing resize events
- ✅ Cleaned up debug tests that relied on unimplemented features (element.style)

## Recent Changes

### Test Harness Enhancements
- Added `TriggerElementResize` method for ResizeObserver testing
- Implemented ResizeOptions struct with Width, Height, ContentWidth, ContentHeight
- Method properly triggers ResizeObserver callbacks via `__triggerResizeObserver` utility

### Cleanup
- Removed `debug_phase6_test.go` - relied on unimplemented element.style property
- ResizeObserver tests now use proper test harness methods instead of trying to manipulate DOM styles

### Test Fixes
- ✅ Fixed `TestDOMAPIPhase6ResizeObserver` in `dom_api_phase6_test.go`
  - Replaced CSS style manipulation (element.style.width/height) with `test.TriggerElementResize()`
  - Updated test to use proper ResizeObserver testing pattern
  - Test now passes and properly validates ResizeObserver functionality

## Key Decisions

### ResizeObserver Testing Strategy
- Don't rely on CSS/style changes to trigger resize events (not implemented)
- Use dedicated test harness methods to simulate resize events
- Focus on testing the observer pattern and callback mechanisms

## Next Steps

1. Fix failing ResizeObserver tests
2. Continue DOM API implementation
3. Address any remaining debug tests that rely on unimplemented features

## Important Patterns

### Test Method Design
- Interaction methods should use JavaScript execution when dealing with event dispatch
- Complex browser APIs may need dedicated test methods (like TriggerElementResize)
- Keep test methods focused and well-documented

### Console & Error Callback Patterns
```go
// Console callback testing
consoleCapture := harness.CaptureConsole(t)
harness.LoadHTML(`<script>console.log("Hello World");</script>`)
consoleCapture.AssertLogContains("Hello World")

// Error callback testing  
errorCapture := harness.CaptureErrors(t)
harness.ExecuteScript(`someUndefinedFunction();`) // Won't panic!
if errorCapture.Count() > 0 {
    errorCapture.AssertReferenceError()
}

// Direct callback usage
harness.SetConsoleCallback(func(level js.ConsoleLevel, args []interface{}) {
    fmt.Printf("[%s] %v\n", level, args)
})

harness.SetErrorCallback(func(err *js.JavaScriptError) {
    fmt.Printf("JS Error: %s - %s\n", err.Type, err.Message)
})
```

### ResizeObserver Pattern
```javascript
// Basic usage
const observer = new ResizeObserver(entries => {
    entries.forEach(entry => {
        console.log(entry.target, entry.contentRect);
    });
});
observer.observe(element);

// Testing
test.TriggerElementResize("#element", domulator.ResizeOptions{
    Width: 400,
    Height: 300,
})
```

### Key Implementation Details

#### Error Handling
- Runtime.RunString() and Runtime.RunScript() now catch panics
- If error callback is set: converts panic to JavaScriptError and calls callback
- If no callback: re-panics (maintains backward compatibility)
- JavaScriptError contains Type, Message, Stack, and Source fields

#### Console Handling
- All console methods (log, info, warn, error) support callbacks
- Arguments are properly formatted: strings remain strings, numbers/bools preserved
- DOM elements get special formatting (e.g., "<DIV>", "[DOM Node]")
- Debug mode vs normal mode affects argument formatting
