# Event Loop Compliance Matrix

**Standard**: WHATWG HTML Living Standard (Section 8.1.6 Event loops)
**Version**: [Specify Version/Date when analysis begins]
**Last Updated**: [Specify Date]

## Feature Compliance

| Feature/Section | Spec Section | Implemented | Compliance Level | Test Coverage | Notes |
|---|---|---|---|---|---|
| **8.1.6 Event loops** | [html.spec.whatwg.org/multipage/webappapis.html#event-loops](https://html.spec.whatwg.org/multipage/webappapis.html#event-loops) |  |  |  |  |
| 8.1.6.1 Definitions | [html.spec.whatwg.org/multipage/webappapis.html#definitions-2](https://html.spec.whatwg.org/multipage/webappapis.html#definitions-2) |  |  |  |  |
| 8.1.6.2 Generic task sources | [html.spec.whatwg.org/multipage/webappapis.html#generic-task-sources](https://html.spec.whatwg.org/multipage/webappapis.html#generic-task-sources) |  |  |  |  |
| 8.1.6.3 Processing model | [html.spec.whatwg.org/multipage/webappapis.html#processing-model](https://html.spec.whatwg.org/multipage/webappapis.html#processing-model) |  |  |  |  |
| - Perform a microtask checkpoint | [html.spec.whatwg.org/multipage/webappapis.html#perform-a-microtask-checkpoint](https://html.spec.whatwg.org/multipage/webappapis.html#perform-a-microtask-checkpoint) |  |  |  |  |
| - Update the rendering | [html.spec.whatwg.org/multipage/webappapis.html#update-the-rendering](https://html.spec.whatwg.org/multipage/webappapis.html#update-the-rendering) |  |  |  |  |
| - Run idle period tasks | [html.spec.whatwg.org/multipage/webappapis.html#run-idle-period-tasks](https://html.spec.whatwg.org/multipage/webappapis.html#run-idle-period-tasks) |  |  |  |  |
| 8.1.6.4 Queuing tasks | [html.spec.whatwg.org/multipage/webappapis.html#queueing-tasks](https://html.spec.whatwg.org/multipage/webappapis.html#queueing-tasks) |  |  |  |  |
| 8.1.6.5 Microtasks | [html.spec.whatwg.org/multipage/webappapis.html#microtasks](https://html.spec.whatwg.org/multipage/webappapis.html#microtasks) |  |  |  |  |
| - `queueMicrotask()` | [html.spec.whatwg.org/multipage/webappapis.html#dom-queuemicrotask](https://html.spec.whatwg.org/multipage/webappapis.html#dom-queuemicrotask) |  |  |  |  |

## Compliance Levels
- **FullCompliance**: Implemented and passes all relevant tests.
- **PartialCompliance**: Partially implemented or some tests failing.
- **IntentionalDeviation**: Implemented differently for specific reasons (e.g., performance, simplification).
- **NotImplemented**: Feature not implemented.
- **NotApplicable**: Feature not applicable to DOMulator's use case.
