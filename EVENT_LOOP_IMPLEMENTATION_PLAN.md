# HTML5 Event Loop Implementation Plan
**DOMulator Phase 3: Advanced Framework Compatibility**

## 🎯 Strategic Objective

Implement a complete HTML5-compliant event loop to achieve **99% compatibility with React, Vue, Angular** and other modern JavaScript frameworks that rely heavily on Promise-based async patterns and precise timing control.

## 📊 Current State Analysis

### ✅ Foundation Complete
- **DOM Implementation**: Complete W3C-compliant DOM with all node types
- **JavaScript Runtime**: Full Goja integration with DOM bindings  
- **CSS Selectors**: Complete querySelector/querySelectorAll support
- **HTML Parser**: Robust golang.org/x/net/html integration
- **Event System**: Complete event propagation and handling
- **Browser APIs**: HTTP/Fetch, FormData, CustomEvent, Storage, URL, History, Performance
- **Testing Framework**: Comprehensive testing harness with network mocking

### ⚠️ Current Gap: Event Loop
- **Current Approach**: Simplified Go timer system for setTimeout/setInterval
- **Missing Components**: Task/microtask queues, animation frames, render steps
- **Framework Impact**: React/Vue/Angular require precise async timing for proper operation

### 🎯 Target Architecture
- **HTML5 Compliance**: Full event loop specification implementation
- **Task Queues**: Macrotasks, microtasks, animation frames, idle callbacks
- **Render Steps**: Style recalculation, layout, paint simulation
- **Testing Integration**: Deterministic async testing with time control

## 🏗️ Implementation Plan: 4-Phase Approach - Main Thread Architecture

### **🎯 Core Architectural Decision: Main Thread Event Loop**

**Rationale**: Single-threaded execution for optimal Goja VM compatibility and deterministic testing
- **Thread Safety**: No race conditions, no complex synchronization needed
- **Goja Compatibility**: Perfect fit with Goja's single-threaded design  
- **Testing Benefits**: Completely deterministic execution for reliable test scenarios
- **Simplicity**: Eliminates goroutine coordination complexity
- **HTML5 Compliance**: Maintains spec compliance with cooperative multitasking

**Implementation Interdependencies** (Critical Path):
```
1. Core EventLoop struct ← Foundation for everything
   ↓
2. Task & Microtask Queues ← HTML5 algorithm core  
   ↓
3. ProcessEventLoopIteration() ← Main algorithm implementation
   ↓
4. Timing system ← Foundation for all scheduling
   ↓
5. Animation & Idle queues ← Can be parallel to #3-4
   ↓
6. Core tests ← Validate foundation before Phase 3B
```

### **Phase 3A: Core Event Loop Foundation** (Week 1-2, 40 hours)

#### **Package Structure: `internal/loop/`**
```
internal/loop/
├── eventloop.go          # Main EventLoop struct and Run() method
├── task.go              # Task definitions and TaskQueue
├── microtask.go         # Microtask definitions and MicrotaskQueue  
├── animation.go         # Animation frame queue and timing
├── idle.go              # Idle callback support
├── timing.go            # Performance timing integration
└── eventloop_test.go    # Comprehensive test suite
```

#### **Core Components Implementation**

##### **1. EventLoop Core Structure**
```go
package loop

import (
    "sync"
    "sync/atomic"
    "time"
    "github.com/dop251/goja"
    "github.com/dpemmons/DOMulator/internal/dom"
)

type EventLoop struct {
    // Core queues (priority order)
    microtaskQueue    *MicrotaskQueue    // Highest priority
    taskQueue         *TaskQueue         // Normal priority  
    animationQueue    *AnimationQueue    // Frame-locked
    idleQueue         *IdleQueue         // Lowest priority
    
    // Runtime integration
    vm                *goja.Runtime
    document          *dom.Document
    
    // State management
    running           atomic.Bool
    blocked           atomic.Bool
    renderingEnabled  atomic.Bool
    
    // Timing control
    frameRate         time.Duration      // 16.67ms for 60fps
    lastFrameTime     time.Time
    performanceStart  time.Time
    
    // Concurrency control
    mu                sync.RWMutex
    shutdown          chan struct{}
    done              chan struct{}
    
    // Metrics
    taskCount         atomic.Int64
    microtaskCount    atomic.Int64
    frameCount        atomic.Int64
}

func New(vm *goja.Runtime, document *dom.Document) *EventLoop {
    return &EventLoop{
        microtaskQueue: NewMicrotaskQueue(),
        taskQueue:      NewTaskQueue(),
        animationQueue: NewAnimationQueue(),
        idleQueue:      NewIdleQueue(),
        vm:            vm,
        document:      document,
        frameRate:     16667 * time.Microsecond, // 60fps
        performanceStart: time.Now(),
        shutdown:      make(chan struct{}),
        done:          make(chan struct{}),
    }
}
```

##### **2. HTML5 Event Loop Algorithm**
```go
func (el *EventLoop) Run() {
    defer close(el.done)
    el.running.Store(true)
    
    for el.running.Load() {
        select {
        case <-el.shutdown:
            return
        default:
            el.processEventLoopIteration()
        }
    }
}

func (el *EventLoop) processEventLoopIteration() {
    // 1. Select a task from the task queue
    task := el.taskQueue.SelectTask()
    if task != nil {
        el.executeTask(task)
        el.taskCount.Add(1)
    }
    
    // 2. Process all microtasks
    el.processMicrotasks()
    
    // 3. Update rendering if needed
    if el.shouldRender() {
        el.performRenderSteps()
        el.frameCount.Add(1)
    }
    
    // 4. Process idle callbacks if time permits
    if el.hasIdleTime() {
        el.processIdleCallbacks()
    }
}

func (el *EventLoop) processMicrotasks() {
    // Process ALL microtasks until queue is empty
    // Prevent infinite loops with a reasonable limit
    const maxMicrotasks = 1000
    processed := 0
    
    for processed < maxMicrotasks {
        microtask := el.microtaskQueue.Dequeue()
        if microtask == nil {
            break
        }
        el.executeMicrotask(microtask)
        el.microtaskCount.Add(1)
        processed++
        
        // New microtasks may have been queued during execution
    }
}
```

##### **3. Task and Microtask Definitions**
```go
// task.go
type TaskType int
const (
    TimerTask TaskType = iota
    DOMEventTask
    NetworkTask
    UserInteractionTask
    HistoryTraversalTask
)

type TaskSource string
const (
    TimerTaskSource           TaskSource = "timer"
    DOMManipulationTaskSource TaskSource = "DOM manipulation"
    UserInteractionTaskSource TaskSource = "user interaction"
    NetworkingTaskSource      TaskSource = "networking"
)

type Task struct {
    ID          int64
    Type        TaskType
    Callback    goja.Callable
    Args        []goja.Value
    Delay       time.Duration
    ScheduledAt time.Time
    Source      TaskSource
}

type TaskQueue struct {
    tasks       []*Task
    mu          sync.Mutex
    nextID      atomic.Int64
}

// microtask.go  
type MicrotaskSource string
const (
    PromiseTaskSource           MicrotaskSource = "promise"
    QueueMicrotaskSource        MicrotaskSource = "queueMicrotask"
    MutationObserverSource      MicrotaskSource = "MutationObserver"
)

type Microtask struct {
    ID       int64
    Callback goja.Callable
    Args     []goja.Value
    Source   MicrotaskSource
}

type MicrotaskQueue struct {
    microtasks []Microtask
    mu         sync.Mutex
    nextID     atomic.Int64
}
```

##### **4. Animation Frame System**
```go
// animation.go
type AnimationFrameCallback struct {
    ID       int64
    Callback goja.Callable
    Timestamp time.Time
}

type AnimationQueue struct {
    callbacks []AnimationFrameCallback
    mu        sync.RWMutex
    nextID    atomic.Int64
}

func (aq *AnimationQueue) Schedule(callback goja.Callable) int64 {
    aq.mu.Lock()
    defer aq.mu.Unlock()
    
    id := aq.nextID.Add(1)
    aq.callbacks = append(aq.callbacks, AnimationFrameCallback{
        ID:        id,
        Callback:  callback,
        Timestamp: time.Now(),
    })
    
    return id
}

func (aq *AnimationQueue) Cancel(id int64) {
    aq.mu.Lock()
    defer aq.mu.Unlock()
    
    for i, callback := range aq.callbacks {
        if callback.ID == id {
            aq.callbacks = append(aq.callbacks[:i], aq.callbacks[i+1:]...)
            break
        }
    }
}
```

#### **Testing Strategy - Phase 3A**
```go
// eventloop_test.go
func TestEventLoopBasicOperation(t *testing.T) {
    // Test basic event loop startup and shutdown
}

func TestTaskQueueOrdering(t *testing.T) {
    // Verify tasks execute in correct order
}

func TestMicrotaskCheckpoints(t *testing.T) {
    // Verify microtasks process at correct checkpoints
}

func TestAnimationFrameTiming(t *testing.T) {
    // Verify animation frame callbacks execute at 60fps
}

func TestEventLoopStarvation(t *testing.T) {
    // Verify no queue starves others
}
```

---

### **Phase 3B: JavaScript API Integration** (Week 3, 20 hours)

#### **Enhanced Promise Implementation**
```go
// js/promises.go
type Promise struct {
    state      PromiseState
    value      goja.Value
    reason     goja.Value
    handlers   []PromiseHandler
    eventLoop  *loop.EventLoop
    mu         sync.Mutex
}

type PromiseState int
const (
    PromisePending PromiseState = iota
    PromiseFulfilled
    PromiseRejected
)

func (el *EventLoop) CreatePromiseConstructor() goja.Value {
    return el.vm.ToValue(func(call goja.ConstructorCall) *goja.Object {
        executor := call.Arguments[0]
        
        promise := &Promise{
            state:     PromisePending,
            eventLoop: el,
            handlers:  make([]PromiseHandler, 0),
        }
        
        // Create resolve/reject functions
        resolve := el.createResolveFunction(promise)
        reject := el.createRejectFunction(promise)
        
        // Execute executor immediately
        _, err := executor.Call(goja.Undefined(), resolve, reject)
        if err != nil {
            // Reject promise if executor throws
            promise.reject(el.vm.ToValue(err))
        }
        
        return el.wrapPromise(promise)
    })
}

func (el *EventLoop) createResolveFunction(promise *Promise) goja.Value {
    return el.vm.ToValue(func(call goja.FunctionCall) goja.Value {
        value := goja.Undefined()
        if len(call.Arguments) > 0 {
            value = call.Arguments[0]
        }
        
        // Queue microtask for promise resolution
        el.QueueMicrotask(func() {
            promise.fulfill(value)
        }, loop.PromiseTaskSource)
        
        return goja.Undefined()
    })
}
```

#### **Core Async APIs**
```go
// Enhanced runtime.go integration
func (r *Runtime) SetupEventLoop() {
    r.eventLoop = loop.New(r.vm, r.document)
    
    // Start event loop in background goroutine
    go r.eventLoop.Run()
    
    // Setup async APIs
    r.setupPromiseAPI()
    r.setupMicrotaskAPI()
    r.setupAnimationFrameAPI()
    r.setupIdleCallbackAPI()
    r.enhanceTimerAPIs()
}

func (r *Runtime) setupMicrotaskAPI() {
    r.vm.Set("queueMicrotask", func(call goja.FunctionCall) goja.Value {
        if len(call.Arguments) < 1 {
            panic(r.vm.NewTypeError("queueMicrotask requires a callback"))
        }
        
        callback, ok := goja.AssertFunction(call.Arguments[0])
        if !ok {
            panic(r.vm.NewTypeError("Callback must be a function"))
        }
        
        r.eventLoop.QueueMicrotask(func() {
            callback.Call(goja.Undefined())
        }, loop.QueueMicrotaskSource)
        
        return goja.Undefined()
    })
}

func (r *Runtime) setupAnimationFrameAPI() {
    r.vm.Set("requestAnimationFrame", func(call goja.FunctionCall) goja.Value {
        if len(call.Arguments) < 1 {
            panic(r.vm.NewTypeError("requestAnimationFrame requires a callback"))
        }
        
        callback, ok := goja.AssertFunction(call.Arguments[0])
        if !ok {
            panic(r.vm.NewTypeError("Callback must be a function"))
        }
        
        id := r.eventLoop.RequestAnimationFrame(callback)
        return r.vm.ToValue(id)
    })
    
    r.vm.Set("cancelAnimationFrame", func(call goja.FunctionCall) goja.Value {
        if len(call.Arguments) < 1 {
            return goja.Undefined()
        }
        
        id := call.Arguments[0].ToInteger()
        r.eventLoop.CancelAnimationFrame(id)
        return goja.Undefined()
    })
}
```

#### **Enhanced Timer APIs**
```go
// Refactor existing setTimeout/setInterval to use event loop
func (r *Runtime) enhanceTimerAPIs() {
    r.vm.Set("setTimeout", func(call goja.FunctionCall) goja.Value {
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
        // HTML5 spec minimum 4ms delay
        if delay < 4 {
            delay = 4
        }

        id := r.eventLoop.SetTimeout(callback, time.Duration(delay)*time.Millisecond)
        return r.vm.ToValue(id)
    })
}
```

---

### **Phase 3C: Advanced Features & Optimization** (Week 4, 20 hours)

#### **Render Steps Simulation**
```go
// render.go
func (el *EventLoop) performRenderSteps() {
    now := time.Now()
    
    // 1. Process animation frame callbacks
    if now.Sub(el.lastFrameTime) >= el.frameRate {
        el.processAnimationFrameCallbacks(now)
        el.lastFrameTime = now
    }
    
    // 2. Trigger any DOM mutation observers
    el.processMutationObservers()
    
    // 3. Update computed styles (simulated)
    el.updateComputedStyles()
    
    // 4. Layout calculations (simulated)
    el.performLayout()
    
    // 5. Paint (simulated - for testing framework)
    el.performPaint()
}

func (el *EventLoop) processAnimationFrameCallbacks(now time.Time) {
    callbacks := el.animationQueue.GetCallbacks()
    
    // Calculate high-resolution timestamp
    timestamp := float64(now.Sub(el.performanceStart).Nanoseconds()) / 1e6
    
    for _, callback := range callbacks {
        // Execute callback with timestamp
        el.QueueMicrotask(func() {
            callback.Callback.Call(goja.Undefined(), el.vm.ToValue(timestamp))
        }, loop.AnimationFrameSource)
    }
    
    el.animationQueue.Clear()
}
```

#### **MutationObserver Integration**
```go
// MutationObserver will be integrated with event loop
type MutationObserver struct {
    callback   goja.Callable
    eventLoop  *loop.EventLoop
    observing  map[*dom.Element]*MutationObserverInit
}

func (mo *MutationObserver) queueMutation(mutation *MutationRecord) {
    mo.eventLoop.QueueMicrotask(func() {
        mutations := []*MutationRecord{mutation}
        mo.callback.Call(goja.Undefined(), mo.eventLoop.vm.ToValue(mutations))
    }, loop.MutationObserverSource)
}
```

#### **Performance Monitoring**
```go
// performance.go
type EventLoopMetrics struct {
    TasksProcessed       int64
    MicrotasksProcessed  int64
    FramesRendered       int64
    AverageTaskTime      time.Duration
    AverageMicrotaskTime time.Duration
    AverageFrameTime     time.Duration
}

func (el *EventLoop) GetMetrics() EventLoopMetrics {
    return EventLoopMetrics{
        TasksProcessed:      el.taskCount.Load(),
        MicrotasksProcessed: el.microtaskCount.Load(),
        FramesRendered:      el.frameCount.Load(),
        // ... additional metrics
    }
}
```

---

### **Phase 3D: Testing Framework Integration** (Week 5, 20 hours)

#### **Event Loop Test Harness**
```go
// testing.go
type EventLoopTestHarness struct {
    eventLoop   *loop.EventLoop
    timeControl *TimeController
    metrics     *MetricsCollector
}

type TimeController struct {
    virtualTime time.Time
    realTime    time.Time
    paused      bool
    mu          sync.Mutex
}

func NewEventLoopTestHarness(vm *goja.Runtime, document *dom.Document) *EventLoopTestHarness {
    eventLoop := loop.New(vm, document)
    return &EventLoopTestHarness{
        eventLoop:   eventLoop,
        timeControl: NewTimeController(),
        metrics:     NewMetricsCollector(),
    }
}

func (h *EventLoopTestHarness) AdvanceTime(duration time.Duration) {
    h.timeControl.AdvanceTime(duration)
    // Process any tasks that should execute in this time window
    h.processTasksUntilTime(h.timeControl.virtualTime)
}

func (h *EventLoopTestHarness) ProcessMicrotasks() {
    h.eventLoop.ProcessAllMicrotasks()
}

func (h *EventLoopTestHarness) ProcessNextTask() *Task {
    return h.eventLoop.ProcessNextTask()
}

func (h *EventLoopTestHarness) ProcessAnimationFrame() {
    h.eventLoop.ProcessAnimationFrame()
}

func (h *EventLoopTestHarness) GetQueueCounts() QueueStats {
    return QueueStats{
        Tasks:           h.eventLoop.taskQueue.Count(),
        Microtasks:      h.eventLoop.microtaskQueue.Count(),
        AnimationFrames: h.eventLoop.animationQueue.Count(),
        IdleCallbacks:   h.eventLoop.idleQueue.Count(),
    }
}
```

#### **Framework Compatibility Tests**
```go
// Framework compatibility test suite
func TestReactCompatibility(t *testing.T) {
    harness := NewEventLoopTestHarness()
    
    // Test useState updates
    result := harness.RunScript(`
        let state = 0;
        let setState = (newState) => {
            Promise.resolve().then(() => {
                state = newState;
                console.log('State updated:', state);
            });
        };
        
        setState(1);
        setState(2);
        state; // Should still be 0
    `)
    
    assert.Equal(t, 0, result.ToInteger())
    
    // Process microtasks - should update state
    harness.ProcessMicrotasks()
    
    finalState := harness.RunScript("state")
    assert.Equal(t, 2, finalState.ToInteger())
}

func TestVueReactivity(t *testing.T) {
    harness := NewEventLoopTestHarness()
    
    // Test Vue-like reactivity timing
    result := harness.RunScript(`
        let data = { count: 0 };
        let watchers = [];
        
        function watch(fn) {
            watchers.push(fn);
        }
        
        function setData(key, value) {
            data[key] = value;
            // Queue watchers as microtasks
            Promise.resolve().then(() => {
                watchers.forEach(fn => fn());
            });
        }
        
        let result = null;
        watch(() => {
            result = data.count * 2;
        });
        
        setData('count', 5);
        result; // Should still be null
    `)
    
    assert.Nil(t, result)
    
    harness.ProcessMicrotasks()
    
    finalResult := harness.RunScript("result")
    assert.Equal(t, 10, finalResult.ToInteger())
}
```

## 📈 Success Metrics

### **Compatibility Targets**
- **React**: 75% → 99% compatibility (hooks, effects, state updates)
- **Vue**: 75% → 99% compatibility (reactivity, watchers, async components)
- **Angular**: 70% → 95% compatibility (zone.js patterns, change detection)
- **HTMX**: Maintain 95% compatibility (no regression)

### **Performance Targets**
- **Event Loop Overhead**: <5% CPU overhead vs current implementation
- **Memory Usage**: <10MB additional memory for full event loop
- **Timing Accuracy**: <1ms deviation from browser behavior
- **Throughput**: 10,000+ microtasks/second processing capability

### **Quality Targets**
- **Test Coverage**: 95%+ coverage for event loop package
- **Integration Tests**: 50+ framework compatibility tests
- **Regression Tests**: All existing tests continue to pass
- **Documentation**: Complete API documentation and usage examples

## 🚀 Implementation Timeline

### **Week 1-2: Phase 3A - Core Foundation** (40 hours)
- **Days 1-3**: EventLoop core structure and basic task/microtask queues
- **Days 4-6**: Animation frame system and idle callback support  
- **Days 7-10**: Comprehensive testing and queue optimization

### **Week 3: Phase 3B - JavaScript Integration** (20 hours)
- **Days 1-2**: Enhanced Promise implementation with microtask scheduling
- **Days 3-4**: queueMicrotask and animation frame APIs
- **Day 5**: Enhanced timer APIs and integration testing

### **Week 4: Phase 3C - Advanced Features** (20 hours)
- **Days 1-2**: Render steps simulation and performance monitoring
- **Days 3-4**: MutationObserver integration and error handling
- **Day 5**: Memory optimization and production readiness

### **Week 5: Phase 3D - Testing Integration** (20 hours)
- **Days 1-2**: EventLoopTestHarness and time control utilities
- **Days 3-4**: Framework compatibility test suite
- **Day 5**: Documentation and final integration testing

## 🔧 Integration Points

### **Modified Components**

1. **`internal/js/runtime.go`** - Enhanced with event loop
2. **`internal/js/bindings.go`** - Event loop-aware bindings
3. **`testing/harness.go`** - Event loop testing support
4. **All browser APIs** - Enhanced with proper async timing

### **New Components**

1. **`internal/loop/`** - Complete event loop package
2. **`internal/js/promises.go`** - Full Promise/A+ implementation
3. **Enhanced testing utilities** - Deterministic async testing

## 🎯 Post-Implementation Benefits

1. **Framework Compatibility**: Full support for modern SPA frameworks
2. **Testing Capabilities**: Deterministic async testing with precise timing control
3. **Performance**: Optimized async operation handling
4. **Future-Proofing**: Foundation for Web Workers, Service Workers, etc.
5. **Market Position**: Industry-leading DOM emulation with true browser compatibility

## ⚠️ Risk Mitigation

### **Technical Risks**
- **Performance Impact**: Mitigated by careful optimization and benchmarking
- **Complexity**: Mitigated by phased implementation and comprehensive testing
- **Browser Compatibility**: Mitigated by following HTML5 specification precisely

### **Project Risks**
- **Timeline**: Conservative estimates with buffer time for unexpected issues
- **Scope Creep**: Clear phase boundaries and success criteria
- **Integration Issues**: Comprehensive regression testing throughout

This implementation will position DOMulator as the premier Go-based web testing solution, capable of handling the most complex modern web applications with the same reliability and precision as real browsers.
