package js

import (
	"fmt"
	"time"

	"github.com/dop251/goja"
	"github.com/dpemmons/DOMulator/internal/browser/storage"
	"github.com/dpemmons/DOMulator/internal/dom"
	"github.com/dpemmons/DOMulator/internal/loop"
)

// Runtime represents a JavaScript runtime environment with DOM integration
type Runtime struct {
	vm             *goja.Runtime
	global         *goja.Object
	document       *dom.Document
	window         *goja.Object
	console        *goja.Object
	eventLoop      *loop.EventLoop // Event loop for async operations
	timers         map[int]*Timer  // Legacy timer support (will be replaced)
	nextTimerID    int
	storageManager *storage.StorageManager
	debugMode      bool // Controls console verbosity
}

// Timer represents a JavaScript timer (setTimeout/setInterval)
type Timer struct {
	ID       int
	Callback goja.Callable
	Timeout  time.Duration
	Interval bool
	Active   bool
	timer    *time.Timer
}

// New creates a new JavaScript runtime with DOM bindings
func New(document *dom.Document) *Runtime {
	vm := goja.New()

	runtime := &Runtime{
		vm:             vm,
		document:       document,
		timers:         make(map[int]*Timer),
		nextTimerID:    1,
		storageManager: storage.NewStorageManager(),
	}

	// Set up global objects
	runtime.setupGlobals()
	runtime.setupConsole()
	runtime.setupTimers()
	runtime.setupDOM()
	runtime.setupStorage()

	return runtime
}

// SetupEventLoop initializes the event loop and modern async APIs
func (r *Runtime) SetupEventLoop(useVirtualTime bool) {
	// Create event loop options
	opts := &loop.EventLoopOptions{
		FrameRate:        16667 * time.Microsecond, // 60fps
		RenderingEnabled: true,
		VirtualTime:      useVirtualTime,
	}

	// Initialize the event loop
	r.eventLoop = loop.New(r.vm, r.document, opts)

	// Set up modern async APIs
	r.setupEventLoopAPIs()

	// Replace legacy timers with event loop versions
	r.setupEventLoopTimers()
}

// EventLoop returns the event loop instance
func (r *Runtime) EventLoop() *loop.EventLoop {
	return r.eventLoop
}

// VM returns the underlying Goja virtual machine
func (r *Runtime) VM() *goja.Runtime {
	return r.vm
}

// Document returns the DOM document
func (r *Runtime) Document() *dom.Document {
	return r.document
}

// RunString executes JavaScript code and returns the result
func (r *Runtime) RunString(code string) (goja.Value, error) {
	return r.vm.RunString(code)
}

// RunScript executes a JavaScript program and returns the result
func (r *Runtime) RunScript(name, code string) (goja.Value, error) {
	program, err := goja.Compile(name, code, false)
	if err != nil {
		return nil, fmt.Errorf("compile error: %w", err)
	}

	return r.vm.RunProgram(program)
}

// SetGlobal sets a global variable in the JavaScript runtime
func (r *Runtime) SetGlobal(name string, value interface{}) {
	r.vm.Set(name, value)
}

// GetGlobal gets a global variable from the JavaScript runtime
func (r *Runtime) GetGlobal(name string) goja.Value {
	return r.vm.Get(name)
}

// setupGlobals initializes global JavaScript objects
func (r *Runtime) setupGlobals() {
	r.global = r.vm.GlobalObject()

	// Create window object
	r.window = r.vm.NewObject()
	r.global.Set("window", r.window)
	r.global.Set("self", r.window)       // window.self === window
	r.global.Set("globalThis", r.window) // ES2020 globalThis

	// Basic window properties
	r.window.Set("location", r.vm.NewObject())
	r.window.Set("navigator", r.vm.NewObject())
	r.window.Set("screen", r.vm.NewObject())
	r.window.Set("history", r.vm.NewObject())

	// Set default window dimensions
	r.window.Set("innerWidth", 800)
	r.window.Set("innerHeight", 600)
	r.window.Set("outerWidth", 800)
	r.window.Set("outerHeight", 600)
}

// SetWindowDimensions sets the window dimensions for testing
func (r *Runtime) SetWindowDimensions(width, height int) {
	if r.window != nil {
		r.window.Set("innerWidth", width)
		r.window.Set("innerHeight", height)
		r.window.Set("outerWidth", width)
		r.window.Set("outerHeight", height)
	}
}

// setupConsole initializes the console object
func (r *Runtime) setupConsole() {
	r.console = r.vm.NewObject()

	// Helper function to format arguments properly
	formatArgs := func(args []goja.Value) []interface{} {
		formatted := make([]interface{}, len(args))
		for i, arg := range args {
			if r.debugMode {
				// In debug mode, export full object details
				formatted[i] = arg.Export()
			} else {
				// In normal mode, use cleaner string representation
				formatted[i] = r.formatConsoleArg(arg)
			}
		}
		return formatted
	}

	// console.log
	r.console.Set("log", func(call goja.FunctionCall) goja.Value {
		if r.debugMode {
			args := formatArgs(call.Arguments)
			fmt.Println(args...)
		}
		// In non-debug mode, suppress console.log entirely to reduce noise
		return goja.Undefined()
	})

	// console.error
	r.console.Set("error", func(call goja.FunctionCall) goja.Value {
		args := formatArgs(call.Arguments)
		allArgs := append([]interface{}{"ERROR:"}, args...)
		fmt.Println(allArgs...)
		return goja.Undefined()
	})

	// console.warn
	r.console.Set("warn", func(call goja.FunctionCall) goja.Value {
		args := formatArgs(call.Arguments)
		allArgs := append([]interface{}{"WARN:"}, args...)
		fmt.Println(allArgs...)
		return goja.Undefined()
	})

	// console.info
	r.console.Set("info", func(call goja.FunctionCall) goja.Value {
		if r.debugMode {
			args := formatArgs(call.Arguments)
			allArgs := append([]interface{}{"INFO:"}, args...)
			fmt.Println(allArgs...)
		}
		return goja.Undefined()
	})

	r.global.Set("console", r.console)
	r.window.Set("console", r.console)
}

// formatConsoleArg formats a JavaScript value for clean console output
func (r *Runtime) formatConsoleArg(arg goja.Value) interface{} {
	if arg == nil || goja.IsUndefined(arg) {
		return "undefined"
	}
	if goja.IsNull(arg) {
		return "null"
	}

	// For objects, provide a cleaner representation
	if obj := arg.ToObject(r.vm); obj != nil {
		// Check if it's a DOM element
		if domNode := obj.Get("__domNode"); domNode != nil && !goja.IsUndefined(domNode) {
			if tagName := obj.Get("tagName"); tagName != nil && !goja.IsUndefined(tagName) {
				return fmt.Sprintf("<%s>", tagName.String())
			}
			if nodeName := obj.Get("nodeName"); nodeName != nil && !goja.IsUndefined(nodeName) {
				return fmt.Sprintf("[%s]", nodeName.String())
			}
			return "[DOM Node]"
		}

		// For plain objects, try to get a reasonable string representation
		if typeProp := obj.Get("type"); typeProp != nil && !goja.IsUndefined(typeProp) {
			return fmt.Sprintf("[Event: %s]", typeProp.String())
		}

		return "[object Object]"
	}

	// For primitives, use string representation
	return arg.String()
}

// SetDebugMode enables or disables verbose console output
func (r *Runtime) SetDebugMode(debug bool) {
	r.debugMode = debug
}

// IsDebugMode returns whether debug mode is enabled
func (r *Runtime) IsDebugMode() bool {
	return r.debugMode
}

// setupTimers initializes setTimeout and setInterval
func (r *Runtime) setupTimers() {
	// setTimeout
	r.global.Set("setTimeout", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 2 {
			panic(r.vm.NewTypeError("setTimeout requires at least 2 arguments"))
		}

		callback, ok := goja.AssertFunction(call.Arguments[0])
		if !ok {
			panic(r.vm.NewTypeError("First argument must be a function"))
		}

		delay := call.Arguments[1].ToInteger()
		if delay < 0 {
			delay = 0
		}

		timer := &Timer{
			ID:       r.nextTimerID,
			Callback: callback,
			Timeout:  time.Duration(delay) * time.Millisecond,
			Interval: false,
			Active:   true,
		}

		r.timers[timer.ID] = timer
		r.nextTimerID++

		// Start the timer
		timer.timer = time.AfterFunc(timer.Timeout, func() {
			if timer.Active {
				_, _ = timer.Callback(goja.Undefined())
				delete(r.timers, timer.ID)
			}
		})

		return r.vm.ToValue(timer.ID)
	})

	// setInterval
	r.global.Set("setInterval", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 2 {
			panic(r.vm.NewTypeError("setInterval requires at least 2 arguments"))
		}

		callback, ok := goja.AssertFunction(call.Arguments[0])
		if !ok {
			panic(r.vm.NewTypeError("First argument must be a function"))
		}

		delay := call.Arguments[1].ToInteger()
		if delay < 0 {
			delay = 0
		}

		timer := &Timer{
			ID:       r.nextTimerID,
			Callback: callback,
			Timeout:  time.Duration(delay) * time.Millisecond,
			Interval: true,
			Active:   true,
		}

		r.timers[timer.ID] = timer
		r.nextTimerID++

		// Start the interval timer
		timer.timer = time.AfterFunc(timer.Timeout, r.intervalCallback(timer))

		return r.vm.ToValue(timer.ID)
	})

	// clearTimeout / clearInterval
	clearTimer := func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return goja.Undefined()
		}

		timerID := int(call.Arguments[0].ToInteger())
		if timer, exists := r.timers[timerID]; exists {
			timer.Active = false
			if timer.timer != nil {
				timer.timer.Stop()
			}
			delete(r.timers, timerID)
		}

		return goja.Undefined()
	}

	r.global.Set("clearTimeout", clearTimer)
	r.global.Set("clearInterval", clearTimer)

	// Also set on window
	r.window.Set("setTimeout", r.global.Get("setTimeout"))
	r.window.Set("setInterval", r.global.Get("setInterval"))
	r.window.Set("clearTimeout", r.global.Get("clearTimeout"))
	r.window.Set("clearInterval", r.global.Get("clearInterval"))
}

// intervalCallback creates a recursive callback for setInterval
func (r *Runtime) intervalCallback(timer *Timer) func() {
	return func() {
		if timer.Active {
			_, _ = timer.Callback(goja.Undefined())
			// Schedule next interval
			timer.timer = time.AfterFunc(timer.Timeout, r.intervalCallback(timer))
		}
	}
}

// setupDOM initializes DOM objects and bindings
func (r *Runtime) setupDOM() {
	// Create DOM bindings wrapper
	domBindings := NewDOMBindings(r.vm, r.document)
	domBindings.jsRuntime = r // Set reference to runtime for event loop access

	// Set document as global
	r.global.Set("document", domBindings.WrapDocument())
	r.window.Set("document", r.global.Get("document"))

	// Set Node constants
	nodeConstants := r.vm.NewObject()
	nodeConstants.Set("ELEMENT_NODE", int(dom.ElementNode))
	nodeConstants.Set("ATTRIBUTE_NODE", int(dom.AttributeNode))
	nodeConstants.Set("TEXT_NODE", int(dom.TextNode))
	nodeConstants.Set("CDATA_SECTION_NODE", int(dom.CDataSectionNode))
	nodeConstants.Set("ENTITY_REFERENCE_NODE", int(dom.EntityReferenceNode))
	nodeConstants.Set("ENTITY_NODE", int(dom.EntityNode))
	nodeConstants.Set("PROCESSING_INSTRUCTION_NODE", int(dom.ProcessingInstructionNode))
	nodeConstants.Set("COMMENT_NODE", int(dom.CommentNode))
	nodeConstants.Set("DOCUMENT_NODE", int(dom.DocumentNode))
	nodeConstants.Set("DOCUMENT_TYPE_NODE", int(dom.DocumentTypeNode))
	nodeConstants.Set("DOCUMENT_FRAGMENT_NODE", int(dom.DocumentFragmentNode))
	nodeConstants.Set("NOTATION_NODE", int(dom.NotationNode))

	r.global.Set("Node", nodeConstants)
	r.window.Set("Node", nodeConstants)

	// Setup browser APIs and global APIs
	domBindings.SetupBrowserAPIs()
	domBindings.SetupGlobalAPIs()
}

// setupStorage initializes localStorage and sessionStorage
func (r *Runtime) setupStorage() {
	storage.SetupStorageAPIs(r.vm, r.storageManager)

	// Also set on window
	r.window.Set("localStorage", r.global.Get("localStorage"))
	r.window.Set("sessionStorage", r.global.Get("sessionStorage"))
}

// SetupFetch adds the fetch API to the runtime with optional network mocks
func (r *Runtime) SetupFetch(networkMocks interface{}) {
	// This will be set externally to avoid import cycles
	// For now, just set a placeholder that will be replaced by the test harness
	r.global.Set("fetch", func(call goja.FunctionCall) goja.Value {
		panic(r.vm.NewTypeError("fetch API not configured - use runtime.SetupFetch() with network mocks"))
	})
	r.window.Set("fetch", r.global.Get("fetch"))
}

// setupEventLoopAPIs initializes modern async JavaScript APIs
func (r *Runtime) setupEventLoopAPIs() {
	if r.eventLoop == nil {
		return // Event loop not initialized
	}

	// queueMicrotask() - W3C Microtask API
	r.global.Set("queueMicrotask", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(r.vm.NewTypeError("queueMicrotask requires a callback function"))
		}

		callback, ok := goja.AssertFunction(call.Arguments[0])
		if !ok {
			panic(r.vm.NewTypeError("First argument must be a function"))
		}

		// Convert to Go function and queue as microtask
		r.eventLoop.QueueMicrotask(func() {
			_, _ = callback(goja.Undefined())
		}, loop.QueueMicrotaskSource)

		return goja.Undefined()
	})

	// requestAnimationFrame() - Animation timing API
	r.global.Set("requestAnimationFrame", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(r.vm.NewTypeError("requestAnimationFrame requires a callback function"))
		}

		callback, ok := goja.AssertFunction(call.Arguments[0])
		if !ok {
			panic(r.vm.NewTypeError("First argument must be a function"))
		}

		id := r.eventLoop.RequestAnimationFrame(callback)
		return r.vm.ToValue(id)
	})

	// cancelAnimationFrame() - Cancel animation frame
	r.global.Set("cancelAnimationFrame", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return goja.Undefined()
		}

		id := call.Arguments[0].ToInteger()
		r.eventLoop.CancelAnimationFrame(id)
		return goja.Undefined()
	})

	// requestIdleCallback() - Idle callback API
	r.global.Set("requestIdleCallback", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(r.vm.NewTypeError("requestIdleCallback requires a callback function"))
		}

		callback, ok := goja.AssertFunction(call.Arguments[0])
		if !ok {
			panic(r.vm.NewTypeError("First argument must be a function"))
		}

		// Optional timeout parameter
		timeout := time.Duration(0)
		if len(call.Arguments) > 1 {
			if options := call.Arguments[1].ToObject(r.vm); options != nil {
				if timeoutValue := options.Get("timeout"); timeoutValue != nil && !goja.IsUndefined(timeoutValue) {
					timeout = time.Duration(timeoutValue.ToInteger()) * time.Millisecond
				}
			}
		}

		id := r.eventLoop.RequestIdleCallback(callback, timeout)
		return r.vm.ToValue(id)
	})

	// cancelIdleCallback() - Cancel idle callback
	r.global.Set("cancelIdleCallback", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return goja.Undefined()
		}

		id := call.Arguments[0].ToInteger()
		r.eventLoop.CancelIdleCallback(id)
		return goja.Undefined()
	})

	// Also set these APIs on window object
	r.window.Set("queueMicrotask", r.global.Get("queueMicrotask"))
	r.window.Set("requestAnimationFrame", r.global.Get("requestAnimationFrame"))
	r.window.Set("cancelAnimationFrame", r.global.Get("cancelAnimationFrame"))
	r.window.Set("requestIdleCallback", r.global.Get("requestIdleCallback"))
	r.window.Set("cancelIdleCallback", r.global.Get("cancelIdleCallback"))
}

// setupEventLoopTimers replaces legacy timers with event loop-based versions
func (r *Runtime) setupEventLoopTimers() {
	if r.eventLoop == nil {
		return // Fall back to legacy timers
	}

	// Enhanced setTimeout with event loop integration
	r.global.Set("setTimeout", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 2 {
			panic(r.vm.NewTypeError("setTimeout requires at least 2 arguments"))
		}

		callback, ok := goja.AssertFunction(call.Arguments[0])
		if !ok {
			panic(r.vm.NewTypeError("First argument must be a function"))
		}

		delay := call.Arguments[1].ToInteger()
		if delay < 0 {
			delay = 0
		}

		// Additional arguments passed to callback
		var args []goja.Value
		if len(call.Arguments) > 2 {
			args = call.Arguments[2:]
		}

		id := r.eventLoop.ScheduleTask(
			loop.TimerTask,
			loop.TimerTaskSource,
			callback,
			args,
			time.Duration(delay)*time.Millisecond,
		)

		return r.vm.ToValue(id)
	})

	// Enhanced setInterval with event loop integration
	r.global.Set("setInterval", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 2 {
			panic(r.vm.NewTypeError("setInterval requires at least 2 arguments"))
		}

		callback, ok := goja.AssertFunction(call.Arguments[0])
		if !ok {
			panic(r.vm.NewTypeError("First argument must be a function"))
		}

		delay := call.Arguments[1].ToInteger()
		if delay < 0 {
			delay = 0
		}

		// Additional arguments passed to callback
		var args []goja.Value
		if len(call.Arguments) > 2 {
			args = call.Arguments[2:]
		}

		// For intervals, we need to create a wrapper that reschedules itself
		var intervalCallback goja.Callable
		intervalCallback = func(this goja.Value, arguments ...goja.Value) (goja.Value, error) {
			// Execute the original callback
			result, err := callback(this, args...)

			// Schedule the next interval
			r.eventLoop.ScheduleTask(
				loop.TimerTask,
				loop.TimerTaskSource,
				intervalCallback,
				nil,
				time.Duration(delay)*time.Millisecond,
			)

			return result, err
		}

		id := r.eventLoop.ScheduleTask(
			loop.TimerTask,
			loop.TimerTaskSource,
			intervalCallback,
			nil,
			time.Duration(delay)*time.Millisecond,
		)

		return r.vm.ToValue(id)
	})

	// clearTimeout/clearInterval remain the same as they work with task IDs
	// The event loop handles task cancellation internally

	// Update window object
	r.window.Set("setTimeout", r.global.Get("setTimeout"))
	r.window.Set("setInterval", r.global.Get("setInterval"))
}

// Shutdown cleans up the runtime and stops all timers
func (r *Runtime) Shutdown() {
	// Stop event loop if it exists and is running
	if r.eventLoop != nil && r.eventLoop.IsRunning() {
		r.eventLoop.Stop()
		r.eventLoop.Wait()
	}

	// Clean up legacy timers
	for _, timer := range r.timers {
		timer.Active = false
		if timer.timer != nil {
			timer.timer.Stop()
		}
	}
	r.timers = make(map[int]*Timer)
}
