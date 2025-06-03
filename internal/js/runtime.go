package js

import (
	"fmt"
	"time"

	"github.com/dop251/goja"
	"github.com/dpemmons/DOMulator/internal/browser/storage"
	"github.com/dpemmons/DOMulator/internal/dom"
)

// Runtime represents a JavaScript runtime environment with DOM integration
type Runtime struct {
	vm             *goja.Runtime
	global         *goja.Object
	document       *dom.Document
	window         *goja.Object
	console        *goja.Object
	timers         map[int]*Timer
	nextTimerID    int
	storageManager *storage.StorageManager
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
}

// setupConsole initializes the console object
func (r *Runtime) setupConsole() {
	r.console = r.vm.NewObject()

	// console.log
	r.console.Set("log", func(call goja.FunctionCall) goja.Value {
		args := make([]interface{}, len(call.Arguments))
		for i, arg := range call.Arguments {
			args[i] = arg.Export()
		}
		fmt.Println(args...)
		return goja.Undefined()
	})

	// console.error
	r.console.Set("error", func(call goja.FunctionCall) goja.Value {
		args := make([]interface{}, len(call.Arguments))
		for i, arg := range call.Arguments {
			args[i] = arg.Export()
		}
		allArgs := append([]interface{}{"ERROR:"}, args...)
		fmt.Println(allArgs...)
		return goja.Undefined()
	})

	// console.warn
	r.console.Set("warn", func(call goja.FunctionCall) goja.Value {
		args := make([]interface{}, len(call.Arguments))
		for i, arg := range call.Arguments {
			args[i] = arg.Export()
		}
		allArgs := append([]interface{}{"WARN:"}, args...)
		fmt.Println(allArgs...)
		return goja.Undefined()
	})

	// console.info
	r.console.Set("info", func(call goja.FunctionCall) goja.Value {
		args := make([]interface{}, len(call.Arguments))
		for i, arg := range call.Arguments {
			args[i] = arg.Export()
		}
		allArgs := append([]interface{}{"INFO:"}, args...)
		fmt.Println(allArgs...)
		return goja.Undefined()
	})

	r.global.Set("console", r.console)
	r.window.Set("console", r.console)
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

// Shutdown cleans up the runtime and stops all timers
func (r *Runtime) Shutdown() {
	for _, timer := range r.timers {
		timer.Active = false
		if timer.timer != nil {
			timer.timer.Stop()
		}
	}
	r.timers = make(map[int]*Timer)
}
