## [3. ]{.secno}[Aborting ongoing activities]{.content}[](#aborting-ongoing-activities){.self-link} {#aborting-ongoing-activities .heading .settled level="3"}

Though promises do not have a built-in aborting mechanism, many APIs
using them require abort semantics.
[`AbortController`{.idl}](#abortcontroller){#ref-for-abortcontroller
link-type="idl"} is meant to support these requirements by providing an
[`abort()`{.idl}](#dom-abortcontroller-abort){#ref-for-dom-abortcontroller-abort①
link-type="idl"} method that toggles the state of a corresponding
[`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal④
link-type="idl"} object. The API which wishes to support aborting can
accept an [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal⑤
link-type="idl"} object, and use its state to determine how to proceed.

APIs that rely upon
[`AbortController`{.idl}](#abortcontroller){#ref-for-abortcontroller①
link-type="idl"} are encouraged to respond to
[`abort()`{.idl}](#dom-abortcontroller-abort){#ref-for-dom-abortcontroller-abort②
link-type="idl"} by rejecting any unsettled promise with the
[`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal⑥
link-type="idl"}'s [abort
reason](#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason
link-type="dfn"}.

::: {#aborting-ongoing-activities-example .example}
[](#aborting-ongoing-activities-example){.self-link}

A hypothetical `doAmazingness({ ... })` method could accept an
[`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal⑦
link-type="idl"} object to support aborting as follows:

``` {.lang-javascript .highlight}
const controller = new AbortController();
const signal = controller.signal;

startSpinner();

doAmazingness({ ..., signal })
  .then(result => ...)
  .catch(err => {
    if (err.name == 'AbortError') return;
    showUserErrorMessage();
  })
  .then(() => stopSpinner());

// …

controller.abort();
```

`doAmazingness` could be implemented as follows:

``` {.lang-javascript .highlight}
function doAmazingness({signal}) {
  return new Promise((resolve, reject) => {
    signal.throwIfAborted();

    // Begin doing amazingness, and call resolve(result) when done.
    // But also, watch for signals:
    signal.addEventListener('abort', () => {
      // Stop doing amazingness, and:
      reject(signal.reason);
    });
  });
}
```
:::

APIs that do not return promises can either react in an equivalent
manner or opt to not surface the
[`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal⑧
link-type="idl"}'s [abort
reason](#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason①
link-type="dfn"} at all.
[`addEventListener()`{.idl}](#dom-eventtarget-addeventlistener){#ref-for-dom-eventtarget-addeventlistener④
link-type="idl"} is an example of an API where the latter made sense.

APIs that require more granular control could extend both
[`AbortController`{.idl}](#abortcontroller){#ref-for-abortcontroller②
link-type="idl"} and
[`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal⑨
link-type="idl"} objects according to their needs.

### [3.1. ]{.secno}[Interface [`AbortController`{.idl}](#abortcontroller){#ref-for-abortcontroller③ link-type="idl"}]{.content}[](#interface-abortcontroller){.self-link} {#interface-abortcontroller .heading .settled level="3.1"}

``` {.idl .highlight .def}
[Exposed=*]
interface AbortController {
  constructor();

  [SameObject] readonly attribute AbortSignal signal;

  undefined abort(optional any reason);
};
```

`controller`{.variable}` = new `[`AbortController`](#dom-abortcontroller-abortcontroller){#ref-for-dom-abortcontroller-abortcontroller① .idl-code link-type="constructor"}`()`
:   Returns a new `controller`{.variable} whose
    [`signal`{.idl}](#dom-abortcontroller-signal){#ref-for-dom-abortcontroller-signal①
    link-type="idl"} is set to a newly created
    [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal①①
    link-type="idl"} object.

`controller`{.variable}` . `[`signal`](#dom-abortcontroller-signal){#ref-for-dom-abortcontroller-signal② .idl-code link-type="attribute"}
:   Returns the
    [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal①②
    link-type="idl"} object associated with this object.

`controller`{.variable}` . `[`abort`](#dom-abortcontroller-abort){#ref-for-dom-abortcontroller-abort④ .idl-code link-type="method"}`(``reason`{.variable}`)`
:   Invoking this method will store `reason`{.variable} in this object's
    [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal①③
    link-type="idl"}'s [abort
    reason](#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason②
    link-type="dfn"}, and signal to any observers that the associated
    activity is to be aborted. If `reason`{.variable} is undefined, then
    an
    \"[`AbortError`{.idl}](https://webidl.spec.whatwg.org/#aborterror){#ref-for-aborterror
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑦
    link-type="idl"} will be stored.

An [`AbortController`{.idl}](#abortcontroller){#ref-for-abortcontroller④
link-type="idl"} object has an associated
[signal]{#abortcontroller-signal .dfn .dfn-paneled
dfn-for="AbortController" dfn-type="dfn" export=""} (an
[`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal①④
link-type="idl"} object).

::: {.algorithm algorithm="AbortController()" algorithm-for="AbortController"}
The [`new AbortController()`]{#dom-abortcontroller-abortcontroller .dfn
.dfn-paneled .idl-code dfn-for="AbortController" dfn-type="constructor"
export="" lt="AbortController()|constructor()"} constructor steps are:

1.  Let `signal`{.variable} be a new
    [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal①⑤
    link-type="idl"} object.

2.  Set [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②④
    link-type="dfn"}'s
    [signal](#abortcontroller-signal){#ref-for-abortcontroller-signal
    link-type="dfn"} to `signal`{.variable}.
:::

The [`signal`]{#dom-abortcontroller-signal .dfn .dfn-paneled .idl-code
dfn-for="AbortController" dfn-type="attribute" export=""} getter steps
are to return
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑤
link-type="dfn"}'s
[signal](#abortcontroller-signal){#ref-for-abortcontroller-signal①
link-type="dfn"}.

::: {.algorithm algorithm="abort(reason)" algorithm-for="AbortController"}
The [`abort(``reason`{.variable}`)`]{#dom-abortcontroller-abort .dfn
.dfn-paneled .idl-code dfn-for="AbortController" dfn-type="method"
export="" lt="abort(reason)|abort()"} method steps are to [signal
abort](#abortcontroller-signal-abort){#ref-for-abortcontroller-signal-abort
link-type="dfn"} on
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑥
link-type="dfn"} with `reason`{.variable} if it is given.
:::

::: {.algorithm algorithm="signal abort" algorithm-for="AbortController"}
To [signal abort]{#abortcontroller-signal-abort .dfn .dfn-paneled
dfn-for="AbortController" dfn-type="dfn" export=""} on an
[`AbortController`{.idl}](#abortcontroller){#ref-for-abortcontroller⑤
link-type="idl"} `controller`{.variable} with an optional
`reason`{.variable}, [signal
abort](#abortsignal-signal-abort){#ref-for-abortsignal-signal-abort
link-type="dfn"} on `controller`{.variable}'s
[signal](#abortcontroller-signal){#ref-for-abortcontroller-signal②
link-type="dfn"} with `reason`{.variable} if it is given.
:::

### [3.2. ]{.secno}[Interface [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal①⑥ link-type="idl"}]{.content}[](#interface-AbortSignal){.self-link} {#interface-AbortSignal .heading .settled level="3.2"}

``` {.idl .highlight .def}
[Exposed=*]
interface AbortSignal : EventTarget {
  [NewObject] static AbortSignal abort(optional any reason);
  [Exposed=(Window,Worker), NewObject] static AbortSignal timeout([EnforceRange] unsigned long long milliseconds);
  [NewObject] static AbortSignal _any(sequence<AbortSignal> signals);

  readonly attribute boolean aborted;
  readonly attribute any reason;
  undefined throwIfAborted();

  attribute EventHandler onabort;
};
```

`AbortSignal . `[`abort`](#dom-abortsignal-abort){#ref-for-dom-abortsignal-abort① .idl-code link-type="method"}`(``reason`{.variable}`)`
:   Returns an
    [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal②①
    link-type="idl"} instance whose [abort
    reason](#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason③
    link-type="dfn"} is set to `reason`{.variable} if not undefined;
    otherwise to an
    \"[`AbortError`{.idl}](https://webidl.spec.whatwg.org/#aborterror){#ref-for-aborterror①
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑧
    link-type="idl"}.

`AbortSignal . `[`any`](#dom-abortsignal-any){#ref-for-dom-abortsignal-any① .idl-code link-type="method"}`(``signals`{.variable}`)`
:   Returns an
    [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal②②
    link-type="idl"} instance which will be aborted once any of
    `signals`{.variable} is aborted. Its [abort
    reason](#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason④
    link-type="dfn"} will be set to whichever one of
    `signals`{.variable} caused it to be aborted.

`AbortSignal . `[`timeout`](#dom-abortsignal-timeout){#ref-for-dom-abortsignal-timeout① .idl-code link-type="method"}`(``milliseconds`{.variable}`)`
:   Returns an
    [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal②③
    link-type="idl"} instance which will be aborted in
    `milliseconds`{.variable} milliseconds. Its [abort
    reason](#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason⑤
    link-type="dfn"} will be set to a
    \"[`TimeoutError`{.idl}](https://webidl.spec.whatwg.org/#timeouterror){#ref-for-timeouterror
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑨
    link-type="idl"}.

`signal`{.variable}` . `[`aborted`](#dom-abortsignal-aborted){#ref-for-dom-abortsignal-aborted① .idl-code link-type="attribute"}
:   Returns true if `signal`{.variable}'s
    [`AbortController`{.idl}](#abortcontroller){#ref-for-abortcontroller⑥
    link-type="idl"} has signaled to abort; otherwise false.

`signal`{.variable}` . `[`reason`](#dom-abortsignal-reason){#ref-for-dom-abortsignal-reason① .idl-code link-type="attribute"}
:   Returns `signal`{.variable}'s [abort
    reason](#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason⑥
    link-type="dfn"}.

`signal`{.variable}` . `[`throwIfAborted`](#dom-abortsignal-throwifaborted){#ref-for-dom-abortsignal-throwifaborted① .idl-code link-type="method"}`()`
:   Throws `signal`{.variable}'s [abort
    reason](#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason⑦
    link-type="dfn"}, if `signal`{.variable}'s
    [`AbortController`{.idl}](#abortcontroller){#ref-for-abortcontroller⑦
    link-type="idl"} has signaled to abort; otherwise, does nothing.

An [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal②④
link-type="idl"} object has an associated [abort
reason]{#abortsignal-abort-reason .dfn .dfn-paneled
dfn-for="AbortSignal" dfn-type="dfn" export=""} (a JavaScript value),
which is initially undefined.

An [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal②⑤
link-type="idl"} object has associated [abort
algorithms]{#abortsignal-abort-algorithms .dfn .dfn-paneled
dfn-for="AbortSignal" dfn-type="dfn" noexport=""}, (a
[set](https://infra.spec.whatwg.org/#ordered-set){#ref-for-ordered-set②
link-type="dfn"} of algorithms which are to be executed when it is
[aborted](#abortsignal-aborted){#ref-for-abortsignal-aborted①
link-type="dfn"}), which is initially empty.

The [abort
algorithms](#abortsignal-abort-algorithms){#ref-for-abortsignal-abort-algorithms
link-type="dfn"} enable APIs with complex requirements to react in a
reasonable way to
[`abort()`{.idl}](#dom-abortcontroller-abort){#ref-for-dom-abortcontroller-abort⑤
link-type="idl"}. For example, a given API's [abort
reason](#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason⑧
link-type="dfn"} might need to be propagated to a cross-thread
environment, such as a service worker.

An [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal②⑥
link-type="idl"} object has a [dependent]{#abortsignal-dependent .dfn
.dfn-paneled dfn-for="AbortSignal" dfn-type="dfn" noexport=""} (a
boolean), which is initially false.

An [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal②⑦
link-type="idl"} object has associated [source
signals]{#abortsignal-source-signals .dfn .dfn-paneled
dfn-for="AbortSignal" dfn-type="dfn" noexport=""} (a weak
[set](https://infra.spec.whatwg.org/#ordered-set){#ref-for-ordered-set③
link-type="dfn"} of
[`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal②⑧
link-type="idl"} objects that the object is dependent on for its
[aborted](#abortsignal-aborted){#ref-for-abortsignal-aborted②
link-type="dfn"} state), which is initially empty.

An [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal②⑨
link-type="idl"} object has associated [dependent
signals]{#abortsignal-dependent-signals .dfn .dfn-paneled
dfn-for="AbortSignal" dfn-type="dfn" noexport=""} (a weak
[set](https://infra.spec.whatwg.org/#ordered-set){#ref-for-ordered-set④
link-type="dfn"} of
[`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal③⓪
link-type="idl"} objects that are dependent on the object for their
[aborted](#abortsignal-aborted){#ref-for-abortsignal-aborted③
link-type="dfn"} state), which is initially empty.

------------------------------------------------------------------------

::: {.algorithm algorithm="abort(reason)" algorithm-for="AbortSignal"}
The static [`abort(``reason`{.variable}`)`]{#dom-abortsignal-abort .dfn
.dfn-paneled .idl-code dfn-for="AbortSignal" dfn-type="method" export=""
lt="abort(reason)|abort()"} method steps are:

1.  Let `signal`{.variable} be a new
    [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal③①
    link-type="idl"} object.

2.  Set `signal`{.variable}'s [abort
    reason](#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason⑨
    link-type="dfn"} to `reason`{.variable} if it is given; otherwise to
    a new
    \"[`AbortError`{.idl}](https://webidl.spec.whatwg.org/#aborterror){#ref-for-aborterror②
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException①⓪
    link-type="idl"}.

3.  Return `signal`{.variable}.
:::

::: {.algorithm algorithm="timeout(milliseconds)" algorithm-for="AbortSignal"}
The static
[`timeout(``milliseconds`{.variable}`)`]{#dom-abortsignal-timeout .dfn
.dfn-paneled .idl-code dfn-for="AbortSignal" dfn-type="method"
export=""} method steps are:

1.  Let `signal`{.variable} be a new
    [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal③②
    link-type="idl"} object.

2.  Let `global`{.variable} be `signal`{.variable}'s [relevant global
    object](https://html.spec.whatwg.org/multipage/webappapis.html#concept-relevant-global){#ref-for-concept-relevant-global①
    link-type="dfn"}.

3.  [Run steps after a
    timeout](https://html.spec.whatwg.org/multipage/timers-and-user-prompts.html#run-steps-after-a-timeout){#ref-for-run-steps-after-a-timeout
    link-type="dfn"} given `global`{.variable},
    \"`AbortSignal-timeout`\", `milliseconds`{.variable}, and the
    following step:

    1.  [Queue a global
        task](https://html.spec.whatwg.org/multipage/webappapis.html#queue-a-global-task){#ref-for-queue-a-global-task
        link-type="dfn"} on the [timer task
        source](https://html.spec.whatwg.org/multipage/timers-and-user-prompts.html#timer-task-source){#ref-for-timer-task-source
        link-type="dfn"} given `global`{.variable} to [signal
        abort](#abortsignal-signal-abort){#ref-for-abortsignal-signal-abort①
        link-type="dfn"} given `signal`{.variable} and a new
        \"[`TimeoutError`{.idl}](https://webidl.spec.whatwg.org/#timeouterror){#ref-for-timeouterror①
        .idl-code link-type="exception"}\"
        [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException①①
        link-type="idl"}.

    For the duration of this timeout, if `signal`{.variable} has any
    event listeners registered for its
    [`abort`{.idl}](#eventdef-abortsignal-abort){#ref-for-eventdef-abortsignal-abort
    link-type="idl"} event, there must be a strong reference from
    `global`{.variable} to `signal`{.variable}.

4.  Return `signal`{.variable}.
:::

::: {.algorithm algorithm="any(signals)" algorithm-for="AbortSignal"}
The static [`any(``signals`{.variable}`)`]{#dom-abortsignal-any .dfn
.dfn-paneled .idl-code dfn-for="AbortSignal" dfn-type="method"
export=""} method steps are to return the result of [creating a
dependent abort
signal](#create-a-dependent-abort-signal){#ref-for-create-a-dependent-abort-signal
link-type="dfn"} from `signals`{.variable} using
[`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal③③
link-type="idl"} and the [current
realm](https://tc39.es/ecma262/#current-realm){#ref-for-current-realm
link-type="dfn"}.
:::

The [`aborted`]{#dom-abortsignal-aborted .dfn .dfn-paneled .idl-code
dfn-for="AbortSignal" dfn-type="attribute" export=""} getter steps are
to return true if
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑦
link-type="dfn"} is
[aborted](#abortsignal-aborted){#ref-for-abortsignal-aborted④
link-type="dfn"}; otherwise false.

The [`reason`]{#dom-abortsignal-reason .dfn .dfn-paneled .idl-code
dfn-for="AbortSignal" dfn-type="attribute" export=""} getter steps are
to return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑧
link-type="dfn"}'s [abort
reason](#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason①⓪
link-type="dfn"}.

The [`throwIfAborted()`]{#dom-abortsignal-throwifaborted .dfn
.dfn-paneled .idl-code dfn-for="AbortSignal" dfn-type="method"
export=""} method steps are to throw
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑨
link-type="dfn"}'s [abort
reason](#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason①①
link-type="dfn"}, if
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this③⓪
link-type="dfn"} is
[aborted](#abortsignal-aborted){#ref-for-abortsignal-aborted⑤
link-type="dfn"}.

::: {#example-throwifaborted .example}
[](#example-throwifaborted){.self-link}

This method is primarily useful for when functions accepting
[`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal③④
link-type="idl"}s want to throw (or return a rejected promise) at
specific checkpoints, instead of passing along the
[`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal③⑤
link-type="idl"} to other methods. For example, the following function
allows aborting in between each attempt to poll for a condition. This
gives opportunities to abort the polling process, even though the actual
asynchronous operation (i.e., `await func()`{.lang-javascript
.highlight}) does not accept an
[`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal③⑥
link-type="idl"}.

``` {.lang-javascript .highlight}
async function waitForCondition(func, targetValue, { signal } = {}) {
  while (true) {
    signal?.throwIfAborted();

    const result = await func();
    if (result === targetValue) {
      return;
    }
  }
}
```
:::

The [`onabort`]{#dom-abortsignal-onabort .dfn .dfn-paneled .idl-code
dfn-for="AbortSignal" dfn-type="attribute" export=""} attribute is an
[event handler IDL
attribute](https://html.spec.whatwg.org/multipage/webappapis.html#event-handler-idl-attributes){#ref-for-event-handler-idl-attributes
link-type="dfn"} for the [`onabort`]{#abortsignal-onabort .dfn
.dfn-paneled dfn-for="AbortSignal" dfn-type="dfn" export=""} [event
handler](https://html.spec.whatwg.org/multipage/webappapis.html#event-handlers){#ref-for-event-handlers①
link-type="dfn"}, whose [event handler event
type](https://html.spec.whatwg.org/multipage/webappapis.html#event-handler-event-type){#ref-for-event-handler-event-type
link-type="dfn"} is [`abort`]{#eventdef-abortsignal-abort .dfn
.dfn-paneled .idl-code dfn-for="AbortSignal" dfn-type="event"
export=""}.

Changes to an [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal③⑦
link-type="idl"} object represent the wishes of the corresponding
[`AbortController`{.idl}](#abortcontroller){#ref-for-abortcontroller⑧
link-type="idl"} object, but an API observing the
[`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal③⑧
link-type="idl"} object can choose to ignore them. For instance, if the
operation has already completed.

------------------------------------------------------------------------

An [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal③⑨
link-type="idl"} object is [aborted]{#abortsignal-aborted .dfn
.dfn-paneled dfn-for="AbortSignal" dfn-type="dfn" export=""} when its
[abort
reason](#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason①②
link-type="dfn"} is not undefined.

::: {.algorithm algorithm="add" algorithm-for="AbortSignal"}
To [add]{#abortsignal-add .dfn .dfn-paneled dfn-for="AbortSignal"
dfn-type="dfn" export=""} an algorithm `algorithm`{.variable} to an
[`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal④⓪
link-type="idl"} object `signal`{.variable}:

1.  If `signal`{.variable} is
    [aborted](#abortsignal-aborted){#ref-for-abortsignal-aborted⑥
    link-type="dfn"}, then return.

2.  [Append](https://infra.spec.whatwg.org/#set-append){#ref-for-set-append①
    link-type="dfn"} `algorithm`{.variable} to `signal`{.variable}'s
    [abort
    algorithms](#abortsignal-abort-algorithms){#ref-for-abortsignal-abort-algorithms①
    link-type="dfn"}.
:::

::: {.algorithm algorithm="remove" algorithm-for="AbortSignal"}
To [remove]{#abortsignal-remove .dfn .dfn-paneled dfn-for="AbortSignal"
dfn-type="dfn" export=""} an algorithm `algorithm`{.variable} from an
[`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal④①
link-type="idl"} `signal`{.variable},
[remove](https://infra.spec.whatwg.org/#list-remove){#ref-for-list-remove①
link-type="dfn"} `algorithm`{.variable} from `signal`{.variable}'s
[abort
algorithms](#abortsignal-abort-algorithms){#ref-for-abortsignal-abort-algorithms②
link-type="dfn"}.
:::

::: {.algorithm algorithm="signal abort" algorithm-for="AbortSignal"}
To [signal abort]{#abortsignal-signal-abort .dfn .dfn-paneled
dfn-for="AbortSignal" dfn-type="dfn" noexport=""}, given an
[`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal④②
link-type="idl"} object `signal`{.variable} and an optional
`reason`{.variable}:

1.  If `signal`{.variable} is
    [aborted](#abortsignal-aborted){#ref-for-abortsignal-aborted⑦
    link-type="dfn"}, then return.

2.  Set `signal`{.variable}'s [abort
    reason](#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason①③
    link-type="dfn"} to `reason`{.variable} if it is given; otherwise to
    a new
    \"[`AbortError`{.idl}](https://webidl.spec.whatwg.org/#aborterror){#ref-for-aborterror③
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException①②
    link-type="idl"}.

3.  Let `dependentSignalsToAbort`{.variable} be a new
    [list](https://infra.spec.whatwg.org/#list){#ref-for-list⑧
    link-type="dfn"}.

4.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate⑧
    link-type="dfn"} `dependentSignal`{.variable} of
    `signal`{.variable}'s [dependent
    signals](#abortsignal-dependent-signals){#ref-for-abortsignal-dependent-signals
    link-type="dfn"}:

    1.  If `dependentSignal`{.variable} is not
        [aborted](#abortsignal-aborted){#ref-for-abortsignal-aborted⑧
        link-type="dfn"}:

        1.  Set `dependentSignal`{.variable}'s [abort
            reason](#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason①④
            link-type="dfn"} to `signal`{.variable}'s [abort
            reason](#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason①⑤
            link-type="dfn"}.

        2.  [Append](https://infra.spec.whatwg.org/#list-append){#ref-for-list-append⑦
            link-type="dfn"} `dependentSignal`{.variable} to
            `dependentSignalsToAbort`{.variable}.

5.  [Run the abort
    steps](#run-the-abort-steps){#ref-for-run-the-abort-steps
    link-type="dfn"} for `signal`{.variable}.

6.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate⑨
    link-type="dfn"} `dependentSignal`{.variable} of
    `dependentSignalsToAbort`{.variable}, [run the abort
    steps](#run-the-abort-steps){#ref-for-run-the-abort-steps①
    link-type="dfn"} for `dependentSignal`{.variable}.
:::

::: {.algorithm algorithm="run the abort steps"}
To [run the abort steps]{#run-the-abort-steps .dfn .dfn-paneled
dfn-type="dfn" noexport=""} for an
[`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal④③
link-type="idl"} `signal`{.variable}:

1.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate①⓪
    link-type="dfn"} `algorithm`{.variable} of `signal`{.variable}'s
    [abort
    algorithms](#abortsignal-abort-algorithms){#ref-for-abortsignal-abort-algorithms③
    link-type="dfn"}: run `algorithm`{.variable}.

2.  [Empty](https://infra.spec.whatwg.org/#list-empty){#ref-for-list-empty
    link-type="dfn"} `signal`{.variable}'s [abort
    algorithms](#abortsignal-abort-algorithms){#ref-for-abortsignal-abort-algorithms④
    link-type="dfn"}.

3.  [Fire an event](#concept-event-fire){#ref-for-concept-event-fire⑤
    link-type="dfn"} named
    [`abort`{.idl}](#eventdef-abortsignal-abort){#ref-for-eventdef-abortsignal-abort①
    link-type="idl"} at `signal`{.variable}.
:::

::: {.algorithm algorithm="create a dependent abort signal"}
To [create a dependent abort signal]{#create-a-dependent-abort-signal
.dfn .dfn-paneled dfn-type="dfn" export=""} from a list of
[`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal④④
link-type="idl"} objects `signals`{.variable}, using
`signalInterface`{.variable}, which must be either
[`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal④⑤
link-type="idl"} or an interface that inherits from it, and a
`realm`{.variable}:

1.  Let `resultSignal`{.variable} be a
    [new](https://webidl.spec.whatwg.org/#new){#ref-for-new
    link-type="dfn"} object implementing `signalInterface`{.variable}
    using `realm`{.variable}.

2.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate①①
    link-type="dfn"} `signal`{.variable} of `signals`{.variable}: if
    `signal`{.variable} is
    [aborted](#abortsignal-aborted){#ref-for-abortsignal-aborted⑨
    link-type="dfn"}, then set `resultSignal`{.variable}'s [abort
    reason](#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason①⑥
    link-type="dfn"} to `signal`{.variable}'s [abort
    reason](#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason①⑦
    link-type="dfn"} and return `resultSignal`{.variable}.

3.  Set `resultSignal`{.variable}'s
    [dependent](#abortsignal-dependent){#ref-for-abortsignal-dependent
    link-type="dfn"} to true.

4.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate①②
    link-type="dfn"} `signal`{.variable} of `signals`{.variable}:

    1.  If `signal`{.variable}'s
        [dependent](#abortsignal-dependent){#ref-for-abortsignal-dependent①
        link-type="dfn"} is false:

        1.  [Append](https://infra.spec.whatwg.org/#set-append){#ref-for-set-append②
            link-type="dfn"} `signal`{.variable} to
            `resultSignal`{.variable}'s [source
            signals](#abortsignal-source-signals){#ref-for-abortsignal-source-signals
            link-type="dfn"}.

        2.  [Append](https://infra.spec.whatwg.org/#set-append){#ref-for-set-append③
            link-type="dfn"} `resultSignal`{.variable} to
            `signal`{.variable}'s [dependent
            signals](#abortsignal-dependent-signals){#ref-for-abortsignal-dependent-signals①
            link-type="dfn"}.

    2.  Otherwise, [for
        each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate①③
        link-type="dfn"} `sourceSignal`{.variable} of
        `signal`{.variable}'s [source
        signals](#abortsignal-source-signals){#ref-for-abortsignal-source-signals①
        link-type="dfn"}:

        1.  Assert: `sourceSignal`{.variable} is not
            [aborted](#abortsignal-aborted){#ref-for-abortsignal-aborted①⓪
            link-type="dfn"} and not
            [dependent](#abortsignal-dependent){#ref-for-abortsignal-dependent②
            link-type="dfn"}.

        2.  [Append](https://infra.spec.whatwg.org/#set-append){#ref-for-set-append④
            link-type="dfn"} `sourceSignal`{.variable} to
            `resultSignal`{.variable}'s [source
            signals](#abortsignal-source-signals){#ref-for-abortsignal-source-signals②
            link-type="dfn"}.

        3.  [Append](https://infra.spec.whatwg.org/#set-append){#ref-for-set-append⑤
            link-type="dfn"} `resultSignal`{.variable} to
            `sourceSignal`{.variable}'s [dependent
            signals](#abortsignal-dependent-signals){#ref-for-abortsignal-dependent-signals②
            link-type="dfn"}.

5.  Return `resultSignal`{.variable}.
:::

#### [3.2.1. ]{.secno}[Garbage collection]{.content}[](#abort-signal-garbage-collection){.self-link} {#abort-signal-garbage-collection .heading .settled level="3.2.1"}

A non-[aborted](#abortsignal-aborted){#ref-for-abortsignal-aborted①①
link-type="dfn"}
[dependent](#abortsignal-dependent){#ref-for-abortsignal-dependent③
link-type="dfn"}
[`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal④⑥
link-type="idl"} object must not be garbage collected while its [source
signals](#abortsignal-source-signals){#ref-for-abortsignal-source-signals③
link-type="dfn"} is non-empty and it has registered event listeners for
its
[`abort`{.idl}](#eventdef-abortsignal-abort){#ref-for-eventdef-abortsignal-abort②
link-type="idl"} event or its [abort
algorithms](#abortsignal-abort-algorithms){#ref-for-abortsignal-abort-algorithms⑤
link-type="dfn"} is non-empty.

### [3.3. ]{.secno}[Using [`AbortController`{.idl}](#abortcontroller){#ref-for-abortcontroller⑨ link-type="idl"} and [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal④⑦ link-type="idl"} objects in APIs]{.content}[](#abortcontroller-api-integration){.self-link} {#abortcontroller-api-integration .heading .settled level="3.3"}

Any web platform API using promises to represent operations that can be
aborted must adhere to the following:

- Accept [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal④⑧
  link-type="idl"} objects through a `signal` dictionary member.
- Convey that the operation got aborted by rejecting the promise with
  [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal④⑨
  link-type="idl"} object's [abort
  reason](#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason①⑧
  link-type="dfn"}.
- Reject immediately if the
  [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal⑤⓪
  link-type="idl"} is already
  [aborted](#abortsignal-aborted){#ref-for-abortsignal-aborted①②
  link-type="dfn"}, otherwise:
- Use the [abort
  algorithms](#abortsignal-abort-algorithms){#ref-for-abortsignal-abort-algorithms⑥
  link-type="dfn"} mechanism to observe changes to the
  [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal⑤①
  link-type="idl"} object and do so in a manner that does not lead to
  clashes with other observers.

::: {#aborting-ongoing-activities-spec-example .example}
[](#aborting-ongoing-activities-spec-example){.self-link}

The method steps for a promise-returning method
`doAmazingness(``options`{.variable}`)` could be as follows:

1.  Let `global`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③①
    link-type="dfn"}'s [relevant global
    object](https://html.spec.whatwg.org/multipage/webappapis.html#concept-relevant-global){#ref-for-concept-relevant-global②
    link-type="dfn"}.

2.  Let `p`{.variable} be [a new
    promise](https://webidl.spec.whatwg.org/#a-new-promise){#ref-for-a-new-promise
    link-type="dfn"}.

3.  If `options`{.variable}\[\"`signal`\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists②
    link-type="dfn"}:

    1.  Let `signal`{.variable} be `options`{.variable}\[\"`signal`\"\].

    2.  If `signal`{.variable} is
        [aborted](#abortsignal-aborted){#ref-for-abortsignal-aborted①③
        link-type="dfn"}, then
        [reject](https://webidl.spec.whatwg.org/#reject){#ref-for-reject
        link-type="dfn"} `p`{.variable} with `signal`{.variable}'s
        [abort
        reason](#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason①⑨
        link-type="dfn"} and return `p`{.variable}.

    3.  [Add the following abort
        steps](#abortsignal-add){#ref-for-abortsignal-add①
        link-type="dfn"} to `signal`{.variable}:

        1.  Stop doing amazing things.

        2.  [Reject](https://webidl.spec.whatwg.org/#reject){#ref-for-reject①
            link-type="dfn"} `p`{.variable} with `signal`{.variable}'s
            [abort
            reason](#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason②⓪
            link-type="dfn"}.

4.  Run these steps [in
    parallel](https://html.spec.whatwg.org/multipage/infrastructure.html#in-parallel){#ref-for-in-parallel①
    link-type="dfn"}:

    1.  Let `amazingResult`{.variable} be the result of doing some
        amazing things.

    2.  [Queue a global
        task](https://html.spec.whatwg.org/multipage/webappapis.html#queue-a-global-task){#ref-for-queue-a-global-task①
        link-type="dfn"} on the amazing task source given
        `global`{.variable} to
        [resolve](https://webidl.spec.whatwg.org/#resolve){#ref-for-resolve
        link-type="dfn"} `p`{.variable} with `amazingResult`{.variable}.

5.  Return `p`{.variable}.
:::

APIs not using promises should still adhere to the above as much as
possible.

