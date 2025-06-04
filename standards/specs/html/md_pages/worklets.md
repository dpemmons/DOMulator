HTML Standard

::: {#head .head .with-buttons}
[![WHATWG](https://resources.whatwg.org/logo.svg){.darkmode-aware
width="100" crossorigin="" height="100"}](https://whatwg.org/){.logo}

# HTML {#html .allcaps}

Living Standard --- Last Updated [2 June 2025]{.pubdate}
:::

[← 10 Web workers](workers.html) --- [Table of Contents](./) --- [12 Web
storage →](webstorage.html)

1.  [[[11]{.secno} Worklets](worklets.html#worklets)]{#toc-worklets}
    1.  [[11.1]{.secno} Introduction](worklets.html#worklets-intro)
        1.  [[11.1.1]{.secno}
            Motivations](worklets.html#worklets-motivations)
        2.  [[11.1.2]{.secno} Code
            idempotence](worklets.html#worklets-idempotent)
        3.  [[11.1.3]{.secno} Speculative
            evaluation](worklets.html#worklets-speculative)
    2.  [[11.2]{.secno} Examples](worklets.html#worklets-examples)
        1.  [[11.2.1]{.secno} Loading
            scripts](worklets.html#worklets-examples-loading)
        2.  [[11.2.2]{.secno} Registering a class and invoking its
            methods](worklets.html#worklets-example-registering)
    3.  [[11.3]{.secno}
        Infrastructure](worklets.html#worklets-infrastructure)
        1.  [[11.3.1]{.secno} The global
            scope](worklets.html#worklets-global)
            1.  [[11.3.1.1]{.secno} Agents and event
                loops](worklets.html#worklet-agents-and-event-loops)
            2.  [[11.3.1.2]{.secno} Creation and
                termination](worklets.html#worklets-creation-termination)
            3.  [[11.3.1.3]{.secno} Script settings for
                worklets](worklets.html#script-settings-for-worklets)
        2.  [[11.3.2]{.secno} The `Worklet`
            class](worklets.html#worklets-worklet)
        3.  [[11.3.3]{.secno} The worklet\'s
            lifetime](worklets.html#worklets-lifetime)

## [11]{.secno} Worklets[](#worklets){.self-link}

### [11.1]{.secno} Introduction[](#worklets-intro){.self-link} {#worklets-intro}

*This section is non-normative.*

Worklets are a piece of specification infrastructure which can be used
for running scripts independent of the main JavaScript execution
environment, while not requiring any particular implementation model.

The worklet infrastructure specified here cannot be used directly by web
developers. Instead, other specifications build upon it to create
directly-usable worklet types, specialized for running in particular
parts of the browser implementation pipeline.

#### [11.1.1]{.secno} Motivations[](#worklets-motivations){.self-link} {#worklets-motivations}

*This section is non-normative.*

Allowing extension points to rendering, or other sensitive parts of the
implementation pipeline such as audio output, is difficult. If extension
points were done with full access to the APIs available on
[`Window`{#worklets-motivations:window}](nav-history-apis.html#window),
engines would need to abandon previously-held assumptions for what could
happen in the middle of those phases. For example, during the layout
phase, rendering engines assume that no DOM will be modified.

Additionally, defining extension points in the
[`Window`{#worklets-motivations:window-2}](nav-history-apis.html#window)
environment would restrict user agents to performing work in the same
thread as the
[`Window`{#worklets-motivations:window-3}](nav-history-apis.html#window)
object. (Unless implementations added complex, high-overhead
infrastructure to allow thread-safe APIs, as well as thread-joining
guarantees.)

Worklets are designed to allow extension points, while keeping
guarantees that user agents currently rely on. This is done through new
global environments, based on subclasses of
[`WorkletGlobalScope`{#worklets-motivations:workletglobalscope}](#workletglobalscope).

Worklets are similar to web workers. However, they:

- Are thread-agnostic. That is, they are not designed to run on a
  dedicated separate thread, like each worker is. Implementations can
  run worklets wherever they choose (including on the main thread).

- Are able to have multiple duplicate instances of the global scope
  created, for the purpose of parallelism.

- Do not use an event-based API. Instead, classes are registered on the
  global scope, whose methods are invoked by the user agent.

- Have a reduced API surface on the global scope.

- Have a lifetime for their [global
  object](webappapis.html#global-object){#worklets-motivations:global-object}
  which is defined by other specifications, often in an
  [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#worklets-motivations:implementation-defined
  x-internal="implementation-defined"} manner.

As worklets have relatively high overhead, they are best used sparingly.
Due to this, a given
[`WorkletGlobalScope`{#worklets-motivations:workletglobalscope-2}](#workletglobalscope)
is expected to be shared between multiple separate scripts. (This is
similar to how a single
[`Window`{#worklets-motivations:window-4}](nav-history-apis.html#window)
is shared between multiple separate scripts.)

Worklets are a general technology that serve different use cases. Some
worklets, such as those defined in CSS Painting API, provide extension
points intended for stateless, idempotent, and short-running
computations, which have special considerations as described in the next
couple of sections. Others, such as those defined in Web Audio API, are
used for stateful, long-running operations.
[\[CSSPAINT\]](references.html#refsCSSPAINT)
[\[WEBAUDIO\]](references.html#refsWEBAUDIO)

#### [11.1.2]{.secno} Code idempotence[](#worklets-idempotent){.self-link} {#worklets-idempotent}

Some specifications which use worklets are intended to allow user agents
to parallelize work over multiple threads, or to move work between
threads as required. In these specifications, user agents might invoke
methods on a web-developer-provided class in an
[implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#worklets-idempotent:implementation-defined
x-internal="implementation-defined"} order.

As a result of this, to prevent interoperability issues, authors who
register classes on such
[`WorkletGlobalScope`{#worklets-idempotent:workletglobalscope}](#workletglobalscope)s
should make their code idempotent. That is, a method or set of methods
on the class should produce the same output given a particular input.

This specification uses the following techniques in order to encourage
authors to write code in an idempotent way:

- No reference to the global object is available (i.e., there is no
  counterpart to
  [`self`{#worklets-idempotent:dom-workerglobalscope-self}](workers.html#dom-workerglobalscope-self)
  on
  [`WorkletGlobalScope`{#worklets-idempotent:workletglobalscope-2}](#workletglobalscope)).

  Although this was the intention when worklets were first specified,
  the introduction of `globalThis` has made it no longer true. See
  [issue #6059](https://github.com/whatwg/html/issues/6059) for more
  discussion.

- Code is loaded as a [module
  script](webappapis.html#module-script){#worklets-idempotent:module-script},
  which results in the code being executed in strict mode and with no
  shared `this` referencing the global proxy.

Together, these restrictions help prevent two different scripts from
sharing state using properties of the [global
object](webappapis.html#global-object){#worklets-idempotent:global-object}.

Additionally, specifications which use worklets and intend to allow
[implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#worklets-idempotent:implementation-defined-2
x-internal="implementation-defined"} behavior must obey the following:

- They must require user agents to always have at least two
  [`WorkletGlobalScope`{#worklets-idempotent:workletglobalscope-3}](#workletglobalscope)
  instances per [`Worklet`{#worklets-idempotent:worklet}](#worklet), and
  randomly assign a method or set of methods on a class to a particular
  [`WorkletGlobalScope`{#worklets-idempotent:workletglobalscope-4}](#workletglobalscope)
  instance. These specifications may provide an opt-out under memory
  constraints.

- These specifications must allow user agents to create and destroy
  instances of their
  [`WorkletGlobalScope`{#worklets-idempotent:workletglobalscope-5}](#workletglobalscope)
  subclasses at any time.

#### [11.1.3]{.secno} Speculative evaluation[](#worklets-speculative){.self-link} {#worklets-speculative}

Some specifications which use worklets can invoke methods on a
web-developer-provided class based on the state of the user agent. To
increase concurrency between threads, a user agent may invoke a method
speculatively, based on potential future states.

In these specifications, user agents might invoke such methods at any
time, and with any arguments, not just ones corresponding to the current
state of the user agent. The results of such speculative evaluations are
not displayed immediately, but can be cached for use if the user agent
state matches the speculated state. This can increase the concurrency
between the user agent and worklet threads.

As a result of this, to prevent interoperability risks between user
agents, authors who register classes on such
[`WorkletGlobalScope`{#worklets-speculative:workletglobalscope}](#workletglobalscope)s
should make their code stateless. That is, the only effect of invoking a
method should be its result, and not any side effects such as updating
mutable state.

The same techniques which encourage [code
idempotence](#worklets-idempotent) also encourage authors to write
stateless code.

::: non-normative
### [11.2]{.secno} Examples[](#worklets-examples){.self-link} {#worklets-examples}

*This section is non-normative.*

For these examples, we\'ll use a fake worklet. The
[`Window`{#worklets-examples:window}](nav-history-apis.html#window)
object provides two [`Worklet`{#worklets-examples:worklet}](#worklet)
instances, which each run code in their own collection of
[`FakeWorkletGlobalScope`{#worklets-examples:fakeworkletglobalscope}](#fakeworkletglobalscope)s:

``` extract
partial interface Window {
  [SameObject, SecureContext] readonly attribute Worklet fakeWorklet1;
  [SameObject, SecureContext] readonly attribute Worklet fakeWorklet2;
};
```

Each
[`Window`{#worklets-examples:window-3}](nav-history-apis.html#window)
has two [`Worklet`{#worklets-examples:worklet-4}](#worklet) instances,
[fake worklet 1]{#fake-worklet-1 .dfn} and [fake worklet
2]{#fake-worklet-2 .dfn}. Both of these have their [worklet global scope
type](#worklet-global-scope-type){#worklets-examples:worklet-global-scope-type}
set to
[`FakeWorkletGlobalScope`{#worklets-examples:fakeworkletglobalscope-2}](#fakeworkletglobalscope),
and their [worklet destination
type](#worklet-destination-type){#worklets-examples:worklet-destination-type}
set to \"`fakeworklet`\". User agents should create at least two
[`FakeWorkletGlobalScope`{#worklets-examples:fakeworkletglobalscope-3}](#fakeworkletglobalscope)
instances per worklet.

\"`fakeworklet`\" is not actually a valid
[destination](https://fetch.spec.whatwg.org/#concept-request-destination){#worklets-examples:concept-request-destination
x-internal="concept-request-destination"} per Fetch. But this
illustrates how real worklets would generally have their own
worklet-type-specific destination.
[\[FETCH\]](references.html#refsFETCH)

The [`fakeWorklet1`]{#fakeworklet1 .dfn noexport=""} getter steps are to
return
[this](https://webidl.spec.whatwg.org/#this){#worklets-examples:this
x-internal="this"}\'s [fake worklet
1](#fake-worklet-1){#worklets-examples:fake-worklet-1}.

The [`fakeWorklet2`]{#fakeworklet2 .dfn noexport=""} getter steps are to
return
[this](https://webidl.spec.whatwg.org/#this){#worklets-examples:this-2
x-internal="this"}\'s [fake worklet
2](#fake-worklet-2){#worklets-examples:fake-worklet-2}.

------------------------------------------------------------------------

``` extract
[Global=(Worklet,FakeWorklet),
 Exposed=FakeWorklet,
 SecureContext]
interface FakeWorkletGlobalScope : WorkletGlobalScope {
  undefined registerFake(DOMString type, Function classConstructor);
};
```

Each
[`FakeWorkletGlobalScope`{#worklets-examples:fakeworkletglobalscope-4}](#fakeworkletglobalscope)
has a [registered class constructors
map]{#registered-class-constructors-map .dfn}, which is an [ordered
map](https://infra.spec.whatwg.org/#ordered-map){#worklets-examples:ordered-map
x-internal="ordered-map"}, initially empty.

The
[`registerFake(``type`{.variable}`, ``classConstructor`{.variable}`)`]{#dom-fakeworkletglobalscope-registerfake
.dfn noexport=""} method steps are to set
[this](https://webidl.spec.whatwg.org/#this){#worklets-examples:this-3
x-internal="this"}\'s [registered class constructors
map](#registered-class-constructors-map){#worklets-examples:registered-class-constructors-map}\[`type`{.variable}\]
to `classConstructor`{.variable}.

#### [11.2.1]{.secno} Loading scripts[](#worklets-examples-loading){.self-link} {#worklets-examples-loading}

*This section is non-normative.*

To load scripts into [fake worklet
1](#fake-worklet-1){#worklets-examples-loading:fake-worklet-1}, a web
developer would write:

``` js
window.fakeWorklet1.addModule('script1.mjs');
window.fakeWorklet1.addModule('script2.mjs');
```

Note that which script finishes fetching and runs first is dependent on
network timing: it could be either `script1.mjs` or `script2.mjs`. This
generally won\'t matter for well-written scripts intended to be loaded
in worklets, if they follow the suggestions about preparing for
[speculative evaluation](#worklets-speculative).

If a web developer wants to perform a task only after the scripts have
successfully run and loaded into some worklets, they could write:

``` js
Promise.all([
    window.fakeWorklet1.addModule('script1.mjs'),
    window.fakeWorklet2.addModule('script2.mjs')
]).then(() => {
    // Do something which relies on those scripts being loaded.
});
```

------------------------------------------------------------------------

Another important point about script-loading is that loaded scripts can
be run in multiple
[`WorkletGlobalScope`{#worklets-examples-loading:workletglobalscope}](#workletglobalscope)s
per [`Worklet`{#worklets-examples-loading:worklet}](#worklet), as
discussed in the section on [code idempotence](#worklets-idempotent). In
particular, the specification above for [fake worklet
1](#fake-worklet-1){#worklets-examples-loading:fake-worklet-1-2} and
[fake worklet
2](#fake-worklet-2){#worklets-examples-loading:fake-worklet-2} require
this. So, consider a scenario such as the following:

``` js
// script.mjs
console.log("Hello from a FakeWorkletGlobalScope!");
```

``` js
// app.mjs
window.fakeWorklet1.addModule("script.mjs");
```

This could result in output such as the following from a user agent\'s
console:

    [fakeWorklet1#1] Hello from a FakeWorkletGlobalScope!
    [fakeWorklet1#4] Hello from a FakeWorkletGlobalScope!
    [fakeWorklet1#2] Hello from a FakeWorkletGlobalScope!
    [fakeWorklet1#3] Hello from a FakeWorkletGlobalScope!

If the user agent at some point decided to kill and restart the third
instance of
[`FakeWorkletGlobalScope`{#worklets-examples-loading:fakeworkletglobalscope}](#fakeworkletglobalscope),
the console would again print
`[fakeWorklet1#3] Hello from a FakeWorkletGlobalScope!` when this
occurs.

#### [11.2.2]{.secno} Registering a class and invoking its methods[](#worklets-example-registering){.self-link} {#worklets-example-registering}

*This section is non-normative.*

Let\'s say that one of the intended usages of our fake worklet by web
developers is to allow them to customize the highly-complex process of
boolean negation. They might register their customization as follows:

``` js
// script.mjs
registerFake('negation-processor', class {
  process(arg) {
    return !arg;
  }
});
```

``` js
// app.mjs
window.fakeWorklet1.addModule("script.mjs");
```

To make use of such registered classes, the specification for fake
worklets could define a [find the opposite of
true]{#fakeworkletglobalscope-process .dfn} algorithm, given a
[`Worklet`{#worklets-example-registering:worklet}](#worklet)
`worklet`{.variable}:

1.  Optionally, [create a worklet global
    scope](#create-a-worklet-global-scope){#worklets-example-registering:create-a-worklet-global-scope}
    for `worklet`{.variable}.

2.  Let `workletGlobalScope`{.variable} be one of
    `worklet`{.variable}\'s [global
    scopes](#concept-worklet-global-scopes){#worklets-example-registering:concept-worklet-global-scopes},
    chosen in an
    [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#worklets-example-registering:implementation-defined
    x-internal="implementation-defined"} manner.

3.  Let `classConstructor`{.variable} be
    `workletGlobalScope`{.variable}\'s [registered class constructors
    map](#registered-class-constructors-map){#worklets-example-registering:registered-class-constructors-map}\[\"`negation-processor`\"\].

4.  Let `classInstance`{.variable} be the result of
    [constructing](https://webidl.spec.whatwg.org/#construct-a-callback-function){#worklets-example-registering:es-constructing-callback-functions
    x-internal="es-constructing-callback-functions"}
    `classConstructor`{.variable}, with no arguments.

5.  Let `function`{.variable} be
    [Get](https://tc39.es/ecma262/#sec-get-o-p){#worklets-example-registering:js-get
    x-internal="js-get"}(`classInstance`{.variable}, \"`process`\").
    Rethrow any exceptions.

6.  Let `callback`{.variable} be the result of
    [converting](https://webidl.spec.whatwg.org/#es-type-mapping){#worklets-example-registering:concept-idl-convert
    x-internal="concept-idl-convert"} `function`{.variable} to a Web IDL
    [`Function`{#worklets-example-registering:idl-function}](https://webidl.spec.whatwg.org/#common-Function){x-internal="idl-function"}
    instance.

7.  Return the result of
    [invoking](https://webidl.spec.whatwg.org/#invoke-a-callback-function){#worklets-example-registering:es-invoking-callback-functions
    x-internal="es-invoking-callback-functions"} `callback`{.variable}
    with « true » and \"`rethrow`\", and with *[callback this
    value](https://webidl.spec.whatwg.org/#dfn-callback-this-value){x-internal="dfn-callback-this-value"}*
    set to `classInstance`{.variable}.

Another, perhaps better, specification architecture would be to extract
the \"`process`\" property and convert it into a
[`Function`{#worklets-example-registering:idl-function-2}](https://webidl.spec.whatwg.org/#common-Function){x-internal="idl-function"}
at registration time, as part of the
[`registerFake()`{#worklets-example-registering:dom-fakeworkletglobalscope-registerfake}](#dom-fakeworkletglobalscope-registerfake)
method steps.
:::

### [11.3]{.secno} Infrastructure[](#worklets-infrastructure){.self-link} {#worklets-infrastructure}

#### [11.3.1]{.secno} The global scope[](#worklets-global){.self-link} {#worklets-global}

Subclasses of
[`WorkletGlobalScope`{#workletglobalscope-dev}](#workletglobalscope) are
used to create [global
objects](webappapis.html#global-object){#worklets-global:global-object}
wherein code loaded into a particular
[`Worklet`{#worklets-global:worklet}](#worklet) can execute.

``` idl
[Exposed=Worklet, SecureContext]
interface WorkletGlobalScope {};
```

Other specifications are intended to subclass
[`WorkletGlobalScope`{#worklets-global:workletglobalscope}](#workletglobalscope),
adding APIs to register a class, as well as other APIs specific for
their worklet type.

Each
[`WorkletGlobalScope`{#worklets-global:workletglobalscope-2}](#workletglobalscope)
has an associated [module map]{#concept-workletglobalscope-module-map
.dfn}. It is a [module
map](webappapis.html#module-map){#worklets-global:module-map}, initially
empty.

##### [11.3.1.1]{.secno} Agents and event loops[](#worklet-agents-and-event-loops){.self-link} {#worklet-agents-and-event-loops}

*This section is non-normative.*

Each
[`WorkletGlobalScope`{#worklet-agents-and-event-loops:workletglobalscope}](#workletglobalscope)
is contained in its own [worklet
agent](webappapis.html#worklet-agent){#worklet-agents-and-event-loops:worklet-agent},
which has its corresponding [event
loop](webappapis.html#concept-agent-event-loop){#worklet-agents-and-event-loops:concept-agent-event-loop}.
However, in practice, implementation of these agents and event loops is
expected to be different from most others.

A [worklet
agent](webappapis.html#worklet-agent){#worklet-agents-and-event-loops:worklet-agent-2}
exists for each
[`WorkletGlobalScope`{#worklet-agents-and-event-loops:workletglobalscope-2}](#workletglobalscope)
since, in theory, an implementation could use a separate thread for each
[`WorkletGlobalScope`{#worklet-agents-and-event-loops:workletglobalscope-3}](#workletglobalscope)
instance, and allowing this level of parallelism is best done using
agents. However, because their \[\[CanBlock\]\] value is false, there is
no requirement that agents and threads are one-to-one. This allows
implementations the freedom to execute scripts loaded into a worklet on
any thread, including one running code from other agents with
\[\[CanBlock\]\] of false, such as the thread of a [similar-origin
window
agent](webappapis.html#similar-origin-window-agent){#worklet-agents-and-event-loops:similar-origin-window-agent}
(\"the main thread\"). Contrast this with [dedicated worker
agents](webappapis.html#dedicated-worker-agent){#worklet-agents-and-event-loops:dedicated-worker-agent},
whose true value for \[\[CanBlock\]\] effectively requires them to get a
dedicated operating system thread.

Worklet [event
loops](webappapis.html#event-loop){#worklet-agents-and-event-loops:event-loop}
are also somewhat special. They are only used for
[tasks](webappapis.html#concept-task){#worklet-agents-and-event-loops:concept-task}
associated with
[`addModule()`{#worklet-agents-and-event-loops:dom-worklet-addmodule}](#dom-worklet-addmodule),
tasks wherein the user agent invokes author-defined methods, and
[microtasks](webappapis.html#microtask){#worklet-agents-and-event-loops:microtask}.
Thus, even though the [event loop processing
model](webappapis.html#event-loop-processing-model) specifies that all
event loops run continuously, implementations can achieve
observably-equivalent results using a simpler strategy, which just
[invokes](https://webidl.spec.whatwg.org/#invoke-a-callback-function){#worklet-agents-and-event-loops:es-invoking-callback-functions
x-internal="es-invoking-callback-functions"} author-provided methods and
then relies on that process to [perform a microtask
checkpoint](webappapis.html#perform-a-microtask-checkpoint){#worklet-agents-and-event-loops:perform-a-microtask-checkpoint}.

##### [11.3.1.2]{.secno} Creation and termination[](#worklets-creation-termination){.self-link} {#worklets-creation-termination}

To [create a worklet global scope]{#create-a-worklet-global-scope .dfn
export=""} for a
[`Worklet`{#worklets-creation-termination:worklet}](#worklet)
`worklet`{.variable}:

1.  Let `outsideSettings`{.variable} be `worklet`{.variable}\'s
    [relevant settings
    object](webappapis.html#relevant-settings-object){#worklets-creation-termination:relevant-settings-object}.

2.  Let `agent`{.variable} be the result of [obtaining a worklet
    agent](webappapis.html#obtain-a-worklet-agent){#worklets-creation-termination:obtain-a-worklet-agent}
    given `outsideSettings`{.variable}. Run the rest of these steps in
    that agent.

3.  Let `realmExecutionContext`{.variable} be the result of [creating a
    new
    realm](webappapis.html#creating-a-new-javascript-realm){#worklets-creation-termination:creating-a-new-javascript-realm}
    given `agent`{.variable} and the following customizations:

    - For the global object, create a new object of the type given by
      `worklet`{.variable}\'s [worklet global scope
      type](#worklet-global-scope-type){#worklets-creation-termination:worklet-global-scope-type}.

4.  Let `workletGlobalScope`{.variable} be the [global
    object](webappapis.html#concept-realm-global){#worklets-creation-termination:concept-realm-global}
    of `realmExecutionContext`{.variable}\'s Realm component.

5.  Let `insideSettings`{.variable} be the result of [setting up a
    worklet environment settings
    object](#set-up-a-worklet-environment-settings-object){#worklets-creation-termination:set-up-a-worklet-environment-settings-object}
    given `realmExecutionContext`{.variable} and
    `outsideSettings`{.variable}.

6.  Let `pendingAddedModules`{.variable} be a
    [clone](https://infra.spec.whatwg.org/#list-clone){#worklets-creation-termination:list-clone
    x-internal="list-clone"} of `worklet`{.variable}\'s [added modules
    list](#concept-worklet-added-modules-list){#worklets-creation-termination:concept-worklet-added-modules-list}.

7.  Let `runNextAddedModule`{.variable} be the following steps:

    1.  If `pendingAddedModules`{.variable} [is not
        empty](https://infra.spec.whatwg.org/#list-is-empty){#worklets-creation-termination:list-is-empty
        x-internal="list-is-empty"}, then:

        1.  Let `moduleURL`{.variable} be the result of
            [dequeuing](https://infra.spec.whatwg.org/#queue-dequeue){#worklets-creation-termination:dequeue
            x-internal="dequeue"} from `pendingAddedModules`{.variable}.

        2.  [Fetch a worklet script
            graph](webappapis.html#fetch-a-worklet-script-graph){#worklets-creation-termination:fetch-a-worklet-script-graph}
            given `moduleURL`{.variable}, `insideSettings`{.variable},
            `worklet`{.variable}\'s [worklet destination
            type](#worklet-destination-type){#worklets-creation-termination:worklet-destination-type},
            [what credentials mode?]{.XXX}, `insideSettings`{.variable},
            `worklet`{.variable}\'s [module responses
            map](#concept-worklet-module-responses-map){#worklets-creation-termination:concept-worklet-module-responses-map},
            and with the following steps given `script`{.variable}:

            This will not actually perform a network request, as it will
            just reuse
            [responses](https://fetch.spec.whatwg.org/#concept-response){#worklets-creation-termination:concept-response
            x-internal="concept-response"} from `worklet`{.variable}\'s
            [module responses
            map](#concept-worklet-module-responses-map){#worklets-creation-termination:concept-worklet-module-responses-map-2}.
            The main purpose of this step is to create a new
            `workletGlobalScope`{.variable}-specific [module
            script](webappapis.html#module-script){#worklets-creation-termination:module-script}
            from the
            [response](https://fetch.spec.whatwg.org/#concept-response){#worklets-creation-termination:concept-response-2
            x-internal="concept-response"}.

            1.  [Assert](https://infra.spec.whatwg.org/#assert){#worklets-creation-termination:assert
                x-internal="assert"}: `script`{.variable} is not null,
                since the fetch succeeded and the source text was
                successfully parsed when `worklet`{.variable}\'s [module
                responses
                map](#concept-worklet-module-responses-map){#worklets-creation-termination:concept-worklet-module-responses-map-3}
                was initially populated with `moduleURL`{.variable}.

            2.  [Run a module
                script](webappapis.html#run-a-module-script){#worklets-creation-termination:run-a-module-script}
                given `script`{.variable}.

            3.  Run `runNextAddedModule`{.variable}.

        3.  Abort these steps.

    2.  [Append](https://infra.spec.whatwg.org/#list-append){#worklets-creation-termination:list-append
        x-internal="list-append"} `workletGlobalScope`{.variable} to
        `outsideSettings`{.variable}\'s [global
        object](webappapis.html#concept-settings-object-global){#worklets-creation-termination:concept-settings-object-global}\'s
        [associated
        `Document`](nav-history-apis.html#concept-document-window){#worklets-creation-termination:concept-document-window}\'s
        [worklet global
        scopes](#concept-document-worklet-global-scopes){#worklets-creation-termination:concept-document-worklet-global-scopes}.

    3.  [Append](https://infra.spec.whatwg.org/#list-append){#worklets-creation-termination:list-append-2
        x-internal="list-append"} `workletGlobalScope`{.variable} to
        `worklet`{.variable}\'s [global
        scopes](#concept-worklet-global-scopes){#worklets-creation-termination:concept-worklet-global-scopes}.

    4.  Run the [responsible event
        loop](webappapis.html#responsible-event-loop){#worklets-creation-termination:responsible-event-loop}
        specified by `insideSettings`{.variable}.

8.  Run `runNextAddedModule`{.variable}.

To [terminate a worklet global scope]{#terminate-a-worklet-global-scope
.dfn export=""} given a
[`WorkletGlobalScope`{#worklets-creation-termination:workletglobalscope}](#workletglobalscope)
`workletGlobalScope`{.variable}:

1.  Let `eventLoop`{.variable} be `workletGlobalScope`{.variable}\'s
    [relevant
    agent](webappapis.html#relevant-agent){#worklets-creation-termination:relevant-agent}\'s
    [event
    loop](webappapis.html#concept-agent-event-loop){#worklets-creation-termination:concept-agent-event-loop}.

2.  If there are any
    [tasks](webappapis.html#concept-task){#worklets-creation-termination:concept-task}
    queued in `eventLoop`{.variable}\'s [task
    queues](webappapis.html#task-queue){#worklets-creation-termination:task-queue},
    discard them without processing them.

3.  Wait for `eventLoop`{.variable} to complete the [currently running
    task](webappapis.html#currently-running-task){#worklets-creation-termination:currently-running-task}.

4.  If the previous step doesn\'t complete within an
    [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#worklets-creation-termination:implementation-defined
    x-internal="implementation-defined"} period of time, then [abort the
    script](webappapis.html#abort-a-running-script){#worklets-creation-termination:abort-a-running-script}
    currently running in the worklet.

5.  Destroy `eventLoop`{.variable}.

6.  [Remove](https://infra.spec.whatwg.org/#list-remove){#worklets-creation-termination:list-remove
    x-internal="list-remove"} `workletGlobalScope`{.variable} from the
    [global
    scopes](#concept-worklet-global-scopes){#worklets-creation-termination:concept-worklet-global-scopes-2}
    of the
    [`Worklet`{#worklets-creation-termination:worklet-2}](#worklet)
    whose [global
    scopes](#concept-worklet-global-scopes){#worklets-creation-termination:concept-worklet-global-scopes-3}
    contains `workletGlobalScope`{.variable}.

7.  [Remove](https://infra.spec.whatwg.org/#list-remove){#worklets-creation-termination:list-remove-2
    x-internal="list-remove"} `workletGlobalScope`{.variable} from the
    [worklet global
    scopes](#concept-document-worklet-global-scopes){#worklets-creation-termination:concept-document-worklet-global-scopes-2}
    of the
    [`Document`{#worklets-creation-termination:document}](dom.html#document)
    whose [worklet global
    scopes](#concept-document-worklet-global-scopes){#worklets-creation-termination:concept-document-worklet-global-scopes-3}
    contains `workletGlobalScope`{.variable}.

##### [11.3.1.3]{.secno} Script settings for worklets[](#script-settings-for-worklets){.self-link}

To [set up a worklet environment settings
object]{#set-up-a-worklet-environment-settings-object .dfn}, given a
[JavaScript execution
context](https://tc39.es/ecma262/#sec-execution-contexts){#script-settings-for-worklets:javascript-execution-context
x-internal="javascript-execution-context"} `executionContext`{.variable}
and an [environment settings
object](webappapis.html#environment-settings-object){#script-settings-for-worklets:environment-settings-object}
`outsideSettings`{.variable}:

1.  Let `origin`{.variable} be a unique [opaque
    origin](browsers.html#concept-origin-opaque){#script-settings-for-worklets:concept-origin-opaque}.

2.  Let `inheritedAPIBaseURL`{.variable} be
    `outsideSettings`{.variable}\'s [API base
    URL](webappapis.html#api-base-url){#script-settings-for-worklets:api-base-url}.

3.  Let `inheritedPolicyContainer`{.variable} be a
    [clone](browsers.html#clone-a-policy-container){#script-settings-for-worklets:clone-a-policy-container}
    of `outsideSettings`{.variable}\'s [policy
    container](webappapis.html#concept-settings-object-policy-container){#script-settings-for-worklets:concept-settings-object-policy-container}.

4.  Let `realm`{.variable} be the value of
    `executionContext`{.variable}\'s Realm component.

5.  Let `workletGlobalScope`{.variable} be `realm`{.variable}\'s [global
    object](webappapis.html#concept-realm-global){#script-settings-for-worklets:concept-realm-global}.

6.  Let `settingsObject`{.variable} be a new [environment settings
    object](webappapis.html#environment-settings-object){#script-settings-for-worklets:environment-settings-object-2}
    whose algorithms are defined as follows:

    The [realm execution context](webappapis.html#realm-execution-context){#script-settings-for-worklets:realm-execution-context}

    :   Return `executionContext`{.variable}.

    The [module map](webappapis.html#concept-settings-object-module-map){#script-settings-for-worklets:concept-settings-object-module-map}

    :   Return `workletGlobalScope`{.variable}\'s [module
        map](#concept-workletglobalscope-module-map){#script-settings-for-worklets:concept-workletglobalscope-module-map}.

    The [API base URL](webappapis.html#api-base-url){#script-settings-for-worklets:api-base-url-2}

    :   Return `inheritedAPIBaseURL`{.variable}.

        Unlike workers or other globals derived from a single resource,
        worklets have no primary resource; instead, multiple scripts,
        each with their own URL, are loaded into the global scope via
        [`worklet.addModule()`{#script-settings-for-worklets:dom-worklet-addmodule}](#dom-worklet-addmodule).
        So this [API base
        URL](webappapis.html#api-base-url){#script-settings-for-worklets:api-base-url-3}
        is rather unlike that of other globals. However, so far this
        doesn\'t matter, as no APIs available to worklet code make use
        of the [API base
        URL](webappapis.html#api-base-url){#script-settings-for-worklets:api-base-url-4}.

    The [origin](webappapis.html#concept-settings-object-origin){#script-settings-for-worklets:concept-settings-object-origin}

    :   Return `origin`{.variable}.

    The [policy container](webappapis.html#concept-settings-object-policy-container){#script-settings-for-worklets:concept-settings-object-policy-container-2}

    :   Return `inheritedPolicyContainer`{.variable}.

    The [cross-origin isolated capability](webappapis.html#concept-settings-object-cross-origin-isolated-capability){#script-settings-for-worklets:concept-settings-object-cross-origin-isolated-capability}
    :   Return [TODO]{.XXX}.

    The [time origin](webappapis.html#concept-settings-object-time-origin){#script-settings-for-worklets:concept-settings-object-time-origin}

    :   [Assert](https://infra.spec.whatwg.org/#assert){#script-settings-for-worklets:assert
        x-internal="assert"}: this algorithm is never called, because
        the [time
        origin](webappapis.html#concept-settings-object-time-origin){#script-settings-for-worklets:concept-settings-object-time-origin-2}
        is not available in a worklet context.

7.  Set `settingsObject`{.variable}\'s
    [id](webappapis.html#concept-environment-id){#script-settings-for-worklets:concept-environment-id}
    to a new unique opaque string, [creation
    URL](webappapis.html#concept-environment-creation-url){#script-settings-for-worklets:concept-environment-creation-url}
    to `inheritedAPIBaseURL`{.variable}, [top-level creation
    URL](webappapis.html#concept-environment-top-level-creation-url){#script-settings-for-worklets:concept-environment-top-level-creation-url}
    to null, [top-level
    origin](webappapis.html#concept-environment-top-level-origin){#script-settings-for-worklets:concept-environment-top-level-origin}
    to `outsideSettings`{.variable}\'s [top-level
    origin](webappapis.html#concept-environment-top-level-origin){#script-settings-for-worklets:concept-environment-top-level-origin-2},
    [target browsing
    context](webappapis.html#concept-environment-target-browsing-context){#script-settings-for-worklets:concept-environment-target-browsing-context}
    to null, and [active service
    worker](webappapis.html#concept-environment-active-service-worker){#script-settings-for-worklets:concept-environment-active-service-worker}
    to null.

8.  Set `realm`{.variable}\'s \[\[HostDefined\]\] field to
    `settingsObject`{.variable}.

9.  Return `settingsObject`{.variable}.

#### [11.3.2]{.secno} The [`Worklet`{#worklets-worklet:worklet}](#worklet) class[](#worklets-worklet){.self-link} {#worklets-worklet}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Worklet](https://developer.mozilla.org/en-US/docs/Web/API/Worklet "The Worklet interface is a lightweight version of Web Workers and gives developers access to low-level parts of the rendering pipeline.")

Support in all current engines.

::: support
[Firefox76+]{.firefox .yes}[Safari14.1+]{.safari
.yes}[Chrome65+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

The [`Worklet`{#worklet-dev}](#worklet) class provides the capability to
add module scripts into its associated
[`WorkletGlobalScope`{#worklets-worklet:workletglobalscope}](#workletglobalscope)s.
The user agent can then create classes registered on the
[`WorkletGlobalScope`{#worklets-worklet:workletglobalscope-2}](#workletglobalscope)s
and invoke their methods.

``` idl
[Exposed=Window, SecureContext]
interface Worklet {
  [NewObject] Promise<undefined> addModule(USVString moduleURL, optional WorkletOptions options = {});
};

dictionary WorkletOptions {
  RequestCredentials credentials = "same-origin";
};
```

Specifications that create
[`Worklet`{#worklets-worklet:worklet-2}](#worklet) instances must
specify the following for a given instance:

- its [worklet global scope type]{#worklet-global-scope-type .dfn
  export=""}, which must be a Web IDL type that
  [inherits](https://webidl.spec.whatwg.org/#dfn-inherit){#worklets-worklet:inherit
  x-internal="inherit"} from
  [`WorkletGlobalScope`{#worklets-worklet:workletglobalscope-3}](#workletglobalscope);
  and

- its [worklet destination type]{#worklet-destination-type .dfn
  export=""}, which must be a
  [destination](https://fetch.spec.whatwg.org/#concept-request-destination){#worklets-worklet:concept-request-destination
  x-internal="concept-request-destination"}, and is used when fetching
  scripts.

`await ``worklet`{.variable}`.`[`addModule`](#dom-worklet-addmodule){#dom-worklet-addmodule-dev}`(``moduleURL`{.variable}`[, { `[`credentials`](#dom-workletoptions-credentials){#worklets-worklet:dom-workletoptions-credentials}` }])`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Worklet/addModule](https://developer.mozilla.org/en-US/docs/Web/API/Worklet/addModule "The addModule() method of the Worklet interface loads the module in the given JavaScript file and adds it to the current Worklet.")

Support in all current engines.

::: support
[Firefox76+]{.firefox .yes}[Safari14.1+]{.safari
.yes}[Chrome65+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

Loads and executes the [module
script](webappapis.html#module-script){#worklets-worklet:module-script}
given by `moduleURL`{.variable} into all of `worklet`{.variable}\'s
[global
scopes](#concept-worklet-global-scopes){#worklets-worklet:concept-worklet-global-scopes}.
It can also create additional global scopes as part of this process,
depending on the worklet type. The returned promise will fulfill once
the script has been successfully loaded and run in all global scopes.

The
[`credentials`{#worklets-worklet:dom-workletoptions-credentials-2}](#dom-workletoptions-credentials)
option can be set to a [credentials
mode](https://fetch.spec.whatwg.org/#concept-request-credentials-mode){#worklets-worklet:concept-request-credentials-mode
x-internal="concept-request-credentials-mode"} to modify the
script-fetching process. It defaults to \"`same-origin`\".

Any failures in
[fetching](webappapis.html#fetch-a-worklet-script-graph){#worklets-worklet:fetch-a-worklet-script-graph}
the script or its dependencies will cause the returned promise to be
rejected with an
[\"`AbortError`\"](https://webidl.spec.whatwg.org/#aborterror){#worklets-worklet:aborterror
x-internal="aborterror"}
[`DOMException`{#worklets-worklet:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.
Any errors in parsing the script or its dependencies will cause the
returned promise to be rejected with the exception generated during
parsing.

A [`Worklet`{#worklets-worklet:worklet-3}](#worklet) has a
[list](https://infra.spec.whatwg.org/#list){#worklets-worklet:list
x-internal="list"} of [global scopes]{#concept-worklet-global-scopes
.dfn dfn-for="Worklet" export=""}, which contains instances of the
[`Worklet`{#worklets-worklet:worklet-4}](#worklet)\'s [worklet global
scope
type](#worklet-global-scope-type){#worklets-worklet:worklet-global-scope-type}.
It is initially empty.

A [`Worklet`{#worklets-worklet:worklet-5}](#worklet) has an [added
modules list]{#concept-worklet-added-modules-list .dfn}, which is a
[list](https://infra.spec.whatwg.org/#list){#worklets-worklet:list-2
x-internal="list"} of
[URLs](https://url.spec.whatwg.org/#concept-url){#worklets-worklet:url
x-internal="url"}, initially empty. Access to this list should be
thread-safe.

A [`Worklet`{#worklets-worklet:worklet-6}](#worklet) has a [module
responses map]{#concept-worklet-module-responses-map .dfn}, which is an
[ordered
map](https://infra.spec.whatwg.org/#ordered-map){#worklets-worklet:ordered-map
x-internal="ordered-map"} from
[URLs](https://url.spec.whatwg.org/#concept-url){#worklets-worklet:url-2
x-internal="url"} to either \"`fetching`\" or
[tuples](https://infra.spec.whatwg.org/#tuple){#worklets-worklet:tuple
x-internal="tuple"} consisting of a
[response](https://fetch.spec.whatwg.org/#concept-response){#worklets-worklet:concept-response
x-internal="concept-response"} and either null, failure, or a [byte
sequence](https://infra.spec.whatwg.org/#byte-sequence){#worklets-worklet:byte-sequence
x-internal="byte-sequence"} representing the response body. This map is
initially empty, and access to it should be thread-safe.

::: note
The [added modules
list](#concept-worklet-added-modules-list){#worklets-worklet:concept-worklet-added-modules-list}
and [module responses
map](#concept-worklet-module-responses-map){#worklets-worklet:concept-worklet-module-responses-map}
exist to ensure that
[`WorkletGlobalScope`{#worklets-worklet:workletglobalscope-4}](#workletglobalscope)s
created at different times get equivalent [module
scripts](webappapis.html#module-script){#worklets-worklet:module-script-2}
run in them, based on the same source text. This allows the creation of
additional
[`WorkletGlobalScope`{#worklets-worklet:workletglobalscope-5}](#workletglobalscope)s
to be transparent to the author.

In practice, user agents are not expected to implement these data
structures, and the algorithms that consult them, using thread-safe
programming techniques. Instead, when
[`addModule()`{#worklets-worklet:dom-worklet-addmodule-2}](#dom-worklet-addmodule)
is called, user agents can fetch the module graph on the main thread,
and send the fetched source text (i.e., the important data contained in
the [module responses
map](#concept-worklet-module-responses-map){#worklets-worklet:concept-worklet-module-responses-map-2})
to each thread which has a
[`WorkletGlobalScope`{#worklets-worklet:workletglobalscope-6}](#workletglobalscope).

Then, when a user agent
[creates](#create-a-worklet-global-scope){#worklets-worklet:create-a-worklet-global-scope}
a new
[`WorkletGlobalScope`{#worklets-worklet:workletglobalscope-7}](#workletglobalscope)
for a given [`Worklet`{#worklets-worklet:worklet-7}](#worklet), it can
simply send the map of fetched source text and the list of entry points
from the main thread to the thread containing the new
[`WorkletGlobalScope`{#worklets-worklet:workletglobalscope-8}](#workletglobalscope).
:::

The
[`addModule(``moduleURL`{.variable}`, ``options`{.variable}`)`]{#dom-worklet-addmodule
.dfn dfn-for="Worklet" dfn-type="method"} method steps are:

1.  Let `outsideSettings`{.variable} be the [relevant settings
    object](webappapis.html#relevant-settings-object){#worklets-worklet:relevant-settings-object}
    of
    [this](https://webidl.spec.whatwg.org/#this){#worklets-worklet:this
    x-internal="this"}.

2.  Let `moduleURLRecord`{.variable} be the result of [encoding-parsing
    a
    URL](urls-and-fetching.html#encoding-parsing-a-url){#worklets-worklet:encoding-parsing-a-url}
    given `moduleURL`{.variable}, relative to
    `outsideSettings`{.variable}.

3.  If `moduleURLRecord`{.variable} is failure, then return [a promise
    rejected
    with](https://webidl.spec.whatwg.org/#a-promise-rejected-with){#worklets-worklet:a-promise-rejected-with
    x-internal="a-promise-rejected-with"} a
    [\"`SyntaxError`\"](https://webidl.spec.whatwg.org/#syntaxerror){#worklets-worklet:syntaxerror
    x-internal="syntaxerror"}
    [`DOMException`{#worklets-worklet:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

4.  Let `promise`{.variable} be a new promise.

5.  Let `workletInstance`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#worklets-worklet:this-2
    x-internal="this"}.

6.  Run the following steps [in
    parallel](infrastructure.html#in-parallel){#worklets-worklet:in-parallel}:

    1.  If `workletInstance`{.variable}\'s [global
        scopes](#concept-worklet-global-scopes){#worklets-worklet:concept-worklet-global-scopes-2}
        [is
        empty](https://infra.spec.whatwg.org/#list-is-empty){#worklets-worklet:list-is-empty
        x-internal="list-is-empty"}, then:

        1.  [Create a worklet global
            scope](#create-a-worklet-global-scope){#worklets-worklet:create-a-worklet-global-scope-2}
            given `workletInstance`{.variable}.

        2.  Optionally,
            [create](#create-a-worklet-global-scope){#worklets-worklet:create-a-worklet-global-scope-3}
            additional global scope instances given
            `workletInstance`{.variable}, depending on the specific
            worklet in question and its specification.

        3.  Wait for all steps of the
            [creation](#create-a-worklet-global-scope){#worklets-worklet:create-a-worklet-global-scope-4}
            process(es) --- including those taking place within the
            [worklet
            agents](webappapis.html#worklet-agent){#worklets-worklet:worklet-agent}
            --- to complete, before moving on.

    2.  Let `pendingTasks`{.variable} be `workletInstance`{.variable}\'s
        [global
        scopes](#concept-worklet-global-scopes){#worklets-worklet:concept-worklet-global-scopes-3}\'s
        [size](https://infra.spec.whatwg.org/#list-size){#worklets-worklet:list-size
        x-internal="list-size"}.

    3.  Let `addedSuccessfully`{.variable} be false.

    4.  [For
        each](https://infra.spec.whatwg.org/#list-iterate){#worklets-worklet:list-iterate
        x-internal="list-iterate"} `workletGlobalScope`{.variable} of
        `workletInstance`{.variable}\'s [global
        scopes](#concept-worklet-global-scopes){#worklets-worklet:concept-worklet-global-scopes-4},
        [queue a global
        task](webappapis.html#queue-a-global-task){#worklets-worklet:queue-a-global-task}
        on the [networking task
        source](webappapis.html#networking-task-source){#worklets-worklet:networking-task-source}
        given `workletGlobalScope`{.variable} to [fetch a worklet script
        graph](webappapis.html#fetch-a-worklet-script-graph){#worklets-worklet:fetch-a-worklet-script-graph-2}
        given `moduleURLRecord`{.variable},
        `outsideSettings`{.variable}, `workletInstance`{.variable}\'s
        [worklet destination
        type](#worklet-destination-type){#worklets-worklet:worklet-destination-type},
        `options`{.variable}\[\"[`credentials`{#worklets-worklet:dom-workletoptions-credentials-3}](#dom-workletoptions-credentials)\"\],
        `workletGlobalScope`{.variable}\'s [relevant settings
        object](webappapis.html#relevant-settings-object){#worklets-worklet:relevant-settings-object-2},
        `workletInstance`{.variable}\'s [module responses
        map](#concept-worklet-module-responses-map){#worklets-worklet:concept-worklet-module-responses-map-3},
        and the following steps given `script`{.variable}:

        Only the first of these fetches will actually perform a network
        request; the ones for other
        [`WorkletGlobalScope`{#worklets-worklet:workletglobalscope-9}](#workletglobalscope)s
        will reuse
        [responses](https://fetch.spec.whatwg.org/#concept-response){#worklets-worklet:concept-response-2
        x-internal="concept-response"} from
        `workletInstance`{.variable}\'s [module responses
        map](#concept-worklet-module-responses-map){#worklets-worklet:concept-worklet-module-responses-map-4}.

        1.  If `script`{.variable} is null, then:

            1.  [Queue a global
                task](webappapis.html#queue-a-global-task){#worklets-worklet:queue-a-global-task-2}
                on the [networking task
                source](webappapis.html#networking-task-source){#worklets-worklet:networking-task-source-2}
                given `workletInstance`{.variable}\'s [relevant global
                object](webappapis.html#concept-relevant-global){#worklets-worklet:concept-relevant-global}
                to perform the following steps:

                1.  If `pendingTasks`{.variable} is not −1, then:

                    1.  Set `pendingTasks`{.variable} to −1.

                    2.  Reject `promise`{.variable} with an
                        [\"`AbortError`\"](https://webidl.spec.whatwg.org/#aborterror){#worklets-worklet:aborterror-2
                        x-internal="aborterror"}
                        [`DOMException`{#worklets-worklet:domexception-3}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

            2.  Abort these steps.

        2.  If `script`{.variable}\'s [error to
            rethrow](webappapis.html#concept-script-error-to-rethrow){#worklets-worklet:concept-script-error-to-rethrow}
            is not null, then:

            1.  [Queue a global
                task](webappapis.html#queue-a-global-task){#worklets-worklet:queue-a-global-task-3}
                on the [networking task
                source](webappapis.html#networking-task-source){#worklets-worklet:networking-task-source-3}
                given `workletInstance`{.variable}\'s [relevant global
                object](webappapis.html#concept-relevant-global){#worklets-worklet:concept-relevant-global-2}
                to perform the following steps:

                1.  If `pendingTasks`{.variable} is not −1, then:

                    1.  Set `pendingTasks`{.variable} to −1.

                    2.  Reject `promise`{.variable} with
                        `script`{.variable}\'s [error to
                        rethrow](webappapis.html#concept-script-error-to-rethrow){#worklets-worklet:concept-script-error-to-rethrow-2}.

            2.  Abort these steps.

        3.  If `addedSuccessfully`{.variable} is false, then:

            1.  [Append](https://infra.spec.whatwg.org/#list-append){#worklets-worklet:list-append
                x-internal="list-append"} `moduleURLRecord`{.variable}
                to `workletInstance`{.variable}\'s [added modules
                list](#concept-worklet-added-modules-list){#worklets-worklet:concept-worklet-added-modules-list-2}.

            2.  Set `addedSuccessfully`{.variable} to true.

        4.  [Run a module
            script](webappapis.html#run-a-module-script){#worklets-worklet:run-a-module-script}
            given `script`{.variable}.

        5.  [Queue a global
            task](webappapis.html#queue-a-global-task){#worklets-worklet:queue-a-global-task-4}
            on the [networking task
            source](webappapis.html#networking-task-source){#worklets-worklet:networking-task-source-4}
            given `workletInstance`{.variable}\'s [relevant global
            object](webappapis.html#concept-relevant-global){#worklets-worklet:concept-relevant-global-3}
            to perform the following steps:

            1.  If `pendingTasks`{.variable} is not −1, then:

                1.  Set `pendingTasks`{.variable} to
                    `pendingTasks`{.variable} − 1.

                2.  If `pendingTasks`{.variable} is 0, then resolve
                    `promise`{.variable}.

7.  Return `promise`{.variable}.

#### [11.3.3]{.secno} The worklet\'s lifetime[](#worklets-lifetime){.self-link} {#worklets-lifetime}

The lifetime of a [`Worklet`{#worklets-lifetime:worklet}](#worklet) has
no special considerations; it is tied to the object it belongs to, such
as the
[`Window`{#worklets-lifetime:window}](nav-history-apis.html#window).

Each [`Document`{#worklets-lifetime:document}](dom.html#document) has a
[worklet global scopes]{#concept-document-worklet-global-scopes .dfn},
which is a
[set](https://infra.spec.whatwg.org/#ordered-set){#worklets-lifetime:set
x-internal="set"} of
[`WorkletGlobalScope`{#worklets-lifetime:workletglobalscope}](#workletglobalscope)s,
initially empty.

The lifetime of a
[`WorkletGlobalScope`{#worklets-lifetime:workletglobalscope-2}](#workletglobalscope)
is, at a minimum, tied to the
[`Document`{#worklets-lifetime:document-2}](dom.html#document) whose
[worklet global
scopes](#concept-document-worklet-global-scopes){#worklets-lifetime:concept-document-worklet-global-scopes}
contain it. In particular,
[destroying](document-lifecycle.html#destroy-a-document){#worklets-lifetime:destroy-a-document}
the [`Document`{#worklets-lifetime:document-3}](dom.html#document) will
[terminate](#terminate-a-worklet-global-scope){#worklets-lifetime:terminate-a-worklet-global-scope}
the corresponding
[`WorkletGlobalScope`{#worklets-lifetime:workletglobalscope-3}](#workletglobalscope)
and allow it to be garbage-collected.

Additionally, user agents may, at any time,
[terminate](#terminate-a-worklet-global-scope){#worklets-lifetime:terminate-a-worklet-global-scope-2}
a given
[`WorkletGlobalScope`{#worklets-lifetime:workletglobalscope-4}](#workletglobalscope),
unless the specification defining the corresponding worklet type says
otherwise. For example, they might terminate them if the [worklet
agent](webappapis.html#worklet-agent){#worklets-lifetime:worklet-agent}\'s
[event
loop](webappapis.html#concept-agent-event-loop){#worklets-lifetime:concept-agent-event-loop}
has no
[tasks](webappapis.html#concept-task){#worklets-lifetime:concept-task}
queued, or if the user agent has no pending operations planning to make
use of the worklet, or if the user agent detects abnormal operations
such as infinite loops or callbacks exceeding imposed time limits.

Finally, specifications for specific worklet types can give more
specific details on when to
[create](#create-a-worklet-global-scope){#worklets-lifetime:create-a-worklet-global-scope}
[`WorkletGlobalScope`{#worklets-lifetime:workletglobalscope-5}](#workletglobalscope)s
for a given worklet type. For example, they might create them during
specific processes that call upon worklet code, as in the
[example](#worklets-example-registering).

[← 10 Web workers](workers.html) --- [Table of Contents](./) --- [12 Web
storage →](webstorage.html)
