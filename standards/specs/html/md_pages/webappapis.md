HTML Standard

::: {#head .head .with-buttons}
[![WHATWG](https://resources.whatwg.org/logo.svg){.darkmode-aware
width="100" crossorigin="" height="100"}](https://whatwg.org/){.logo}

# HTML {#html .allcaps}

Living Standard --- Last Updated [2 June 2025]{.pubdate}
:::

[← 7.5 Document lifecycle](document-lifecycle.html) --- [Table of
Contents](./) --- [8.4 Dynamic markup insertion
→](dynamic-markup-insertion.html)

1.  [[[8]{.secno} Web application
    APIs](webappapis.html#webappapis)]{#toc-webappapis}
    1.  [[8.1]{.secno} Scripting](webappapis.html#scripting)
        1.  [[8.1.1]{.secno}
            Introduction](webappapis.html#introduction-11)
        2.  [[8.1.2]{.secno} Agents and agent
            clusters](webappapis.html#agents-and-agent-clusters)
            1.  [[8.1.2.1]{.secno} Integration with the JavaScript agent
                formalism](webappapis.html#integration-with-the-javascript-agent-formalism)
            2.  [[8.1.2.2]{.secno} Integration with the JavaScript agent
                cluster
                formalism](webappapis.html#integration-with-the-javascript-agent-cluster-formalism)
        3.  [[8.1.3]{.secno} Realms and their
            counterparts](webappapis.html#realms-and-their-counterparts)
            1.  [[8.1.3.1]{.secno}
                Environments](webappapis.html#environments)
            2.  [[8.1.3.2]{.secno} Environment settings
                objects](webappapis.html#environment-settings-objects)
            3.  [[8.1.3.3]{.secno} Realms, settings objects, and global
                objects](webappapis.html#realms-settings-objects-global-objects)
                1.  [[8.1.3.3.1]{.secno} Entry](webappapis.html#entry)
                2.  [[8.1.3.3.2]{.secno}
                    Incumbent](webappapis.html#incumbent)
                3.  [[8.1.3.3.3]{.secno}
                    Current](webappapis.html#current)
                4.  [[8.1.3.3.4]{.secno}
                    Relevant](webappapis.html#relevant)
            4.  [[8.1.3.4]{.secno} Enabling and disabling
                scripting](webappapis.html#enabling-and-disabling-scripting)
            5.  [[8.1.3.5]{.secno} Secure
                contexts](webappapis.html#secure-contexts)
        4.  [[8.1.4]{.secno} Script processing
            model](webappapis.html#scripting-processing-model)
            1.  [[8.1.4.1]{.secno}
                Scripts](webappapis.html#script-structs)
            2.  [[8.1.4.2]{.secno} Fetching
                scripts](webappapis.html#fetching-scripts)
            3.  [[8.1.4.3]{.secno} Creating
                scripts](webappapis.html#creating-scripts)
            4.  [[8.1.4.4]{.secno} Calling
                scripts](webappapis.html#calling-scripts)
            5.  [[8.1.4.5]{.secno} Killing
                scripts](webappapis.html#killing-scripts)
            6.  [[8.1.4.6]{.secno} Runtime script
                errors](webappapis.html#runtime-script-errors)
            7.  [[8.1.4.7]{.secno} Unhandled promise
                rejections](webappapis.html#unhandled-promise-rejections)
            8.  [[8.1.4.8]{.secno} Import map parse
                results](webappapis.html#import-map-parse-results)
        5.  [[8.1.5]{.secno} Module specifier
            resolution](webappapis.html#module-specifier-resolution)
            1.  [[8.1.5.1]{.secno} The resolution
                algorithm](webappapis.html#the-resolution-algorithm)
            2.  [[8.1.5.2]{.secno} Import
                maps](webappapis.html#import-maps)
            3.  [[8.1.5.3]{.secno} Import map processing
                model](webappapis.html#import-map-processing-model)
        6.  [[8.1.6]{.secno} JavaScript specification host
            hooks](webappapis.html#javascript-specification-host-hooks)
            1.  [[8.1.6.1]{.secno}
                HostEnsureCanAddPrivateElement(`O`{.variable})](webappapis.html#the-hostensurecanaddprivateelement-implementation)
            2.  [[8.1.6.2]{.secno}
                HostEnsureCanCompileStrings(`realm`{.variable},
                `parameterStrings`{.variable}, `bodyString`{.variable},
                `codeString`{.variable}, `compilationType`{.variable},
                `parameterArgs`{.variable},
                `bodyArg`{.variable})](webappapis.html#hostensurecancompilestrings(realm,-parameterstrings,-bodystring,-codestring,-compilationtype,-parameterargs,-bodyarg))
            3.  [[8.1.6.3]{.secno}
                HostGetCodeForEval(`argument`{.variable})](webappapis.html#hostgetcodeforeval(argument))
            4.  [[8.1.6.4]{.secno}
                HostPromiseRejectionTracker(`promise`{.variable},
                `operation`{.variable})](webappapis.html#the-hostpromiserejectiontracker-implementation)
            5.  [[8.1.6.5]{.secno}
                HostSystemUTCEpochNanoseconds(`global`{.variable})](webappapis.html#hostsystemutcepochnanoseconds)
            6.  [[8.1.6.6]{.secno} Job-related host
                hooks](webappapis.html#integration-with-javascript-jobs)
                1.  [[8.1.6.6.1]{.secno}
                    HostCallJobCallback(`callback`{.variable},
                    `V`{.variable},
                    `argumentsList`{.variable})](webappapis.html#hostcalljobcallback)
                2.  [[8.1.6.6.2]{.secno}
                    HostEnqueueFinalizationRegistryCleanupJob(`finalizationRegistry`{.variable})](webappapis.html#hostenqueuefinalizationregistrycleanupjob)
                3.  [[8.1.6.6.3]{.secno}
                    HostEnqueueGenericJob(`job`{.variable},
                    `realm`{.variable})](webappapis.html#hostenqueuegenericjob)
                4.  [[8.1.6.6.4]{.secno}
                    HostEnqueuePromiseJob(`job`{.variable},
                    `realm`{.variable})](webappapis.html#hostenqueuepromisejob)
                5.  [[8.1.6.6.5]{.secno}
                    HostEnqueueTimeoutJob(`job`{.variable},
                    `realm`{.variable},
                    `milliseconds`{.variable})](webappapis.html#hostenqueuetimeoutjob)
                6.  [[8.1.6.6.6]{.secno}
                    HostMakeJobCallback(`callable`{.variable})](webappapis.html#hostmakejobcallback)
            7.  [[8.1.6.7]{.secno} Module-related host
                hooks](webappapis.html#integration-with-the-javascript-module-system)
                1.  [[8.1.6.7.1]{.secno}
                    HostGetImportMetaProperties(`moduleRecord`{.variable})](webappapis.html#hostgetimportmetaproperties)
                2.  [[8.1.6.7.2]{.secno}
                    HostGetSupportedImportAttributes()](webappapis.html#hostgetsupportedimportattributes)
                3.  [[8.1.6.7.3]{.secno}
                    HostLoadImportedModule(`referrer`{.variable},
                    `moduleRequest`{.variable}, `loadState`{.variable},
                    `payload`{.variable})](webappapis.html#hostloadimportedmodule)
        7.  [[8.1.7]{.secno} Event loops](webappapis.html#event-loops)
            1.  [[8.1.7.1]{.secno}
                Definitions](webappapis.html#definitions-3)
            2.  [[8.1.7.2]{.secno} Queuing
                tasks](webappapis.html#queuing-tasks)
            3.  [[8.1.7.3]{.secno} Processing
                model](webappapis.html#event-loop-processing-model)
            4.  [[8.1.7.4]{.secno} Generic task
                sources](webappapis.html#generic-task-sources)
            5.  [[8.1.7.5]{.secno} Dealing with the event loop from
                other
                specifications](webappapis.html#event-loop-for-spec-authors)
        8.  [[8.1.8]{.secno} Events](webappapis.html#events)
            1.  [[8.1.8.1]{.secno} Event
                handlers](webappapis.html#event-handler-attributes)
            2.  [[8.1.8.2]{.secno} Event handlers on elements,
                `Document` objects, and `Window`
                objects](webappapis.html#event-handlers-on-elements,-document-objects,-and-window-objects)
                1.  [[8.1.8.2.1]{.secno} IDL
                    definitions](webappapis.html#idl-definitions)
            3.  [[8.1.8.3]{.secno} Event
                firing](webappapis.html#event-firing)
    2.  [[8.2]{.secno} The `WindowOrWorkerGlobalScope`
        mixin](webappapis.html#windoworworkerglobalscope-mixin)
    3.  [[8.3]{.secno} Base64 utility methods](webappapis.html#atob)

## [8]{.secno} Web application APIs[](#webappapis){.self-link} {#webappapis}

### [8.1]{.secno} Scripting[](#scripting){.self-link}

#### [8.1.1]{.secno} Introduction[](#introduction-11){.self-link} {#introduction-11}

Various mechanisms can cause author-provided executable code to run in
the context of a document. These mechanisms include, but are probably
not limited to:

- Processing of
  [`script`{#introduction-11:the-script-element}](scripting.html#the-script-element)
  elements.
- Navigating to [`javascript:`
  URLs](browsing-the-web.html#the-javascript:-url-special-case){#introduction-11:the-javascript:-url-special-case}.
- Event handlers, whether registered through the DOM using
  `addEventListener()`, by explicit [event handler content
  attributes](#event-handler-content-attributes){#introduction-11:event-handler-content-attributes},
  by [event handler IDL
  attributes](#event-handler-idl-attributes){#introduction-11:event-handler-idl-attributes},
  or otherwise.
- Processing of technologies like SVG that have their own scripting
  features.

#### [8.1.2]{.secno} Agents and agent clusters[](#agents-and-agent-clusters){.self-link}

##### [8.1.2.1]{.secno} Integration with the JavaScript agent formalism[](#integration-with-the-javascript-agent-formalism){.self-link}

JavaScript defines the concept of an
[agent](https://tc39.es/ecma262/#sec-agents){#integration-with-the-javascript-agent-formalism:agent
x-internal="agent"}. This section gives the mapping of that
language-level concept on to the web platform.

::: note
Conceptually, the
[agent](https://tc39.es/ecma262/#sec-agents){#integration-with-the-javascript-agent-formalism:agent-2
x-internal="agent"} concept is an architecture-independent, idealized
\"thread\" in which JavaScript code runs. Such code can involve multiple
globals/[realms](#concept-global-object-realm){#integration-with-the-javascript-agent-formalism:concept-global-object-realm}
that can synchronously access each other, and thus needs to run in a
single execution thread.

Two
[`Window`{#integration-with-the-javascript-agent-formalism:window}](nav-history-apis.html#window)
objects having the same
[agent](https://tc39.es/ecma262/#sec-agents){#integration-with-the-javascript-agent-formalism:agent-3
x-internal="agent"} does not indicate they can directly access all
objects created in each other\'s realms. They would have to be [same
origin-domain](browsers.html#same-origin-domain){#integration-with-the-javascript-agent-formalism:same-origin-domain};
see
[IsPlatformObjectSameOrigin](nav-history-apis.html#isplatformobjectsameorigin-(-o-)){#integration-with-the-javascript-agent-formalism:isplatformobjectsameorigin-(-o-)}.
:::

The following types of agents exist on the web platform:

[Similar-origin window agent]{#similar-origin-window-agent .dfn export=""}

:   Contains various
    [`Window`{#integration-with-the-javascript-agent-formalism:window-2}](nav-history-apis.html#window)
    objects which can potentially reach each other, either directly or
    by using
    [`document.domain`{#integration-with-the-javascript-agent-formalism:dom-document-domain}](browsers.html#dom-document-domain).

    If the encompassing [agent
    cluster](https://tc39.es/ecma262/#sec-agent-clusters){#integration-with-the-javascript-agent-formalism:agent-cluster
    x-internal="agent-cluster"}\'s [is
    origin-keyed](#is-origin-keyed){#integration-with-the-javascript-agent-formalism:is-origin-keyed}
    is true, then all the
    [`Window`{#integration-with-the-javascript-agent-formalism:window-3}](nav-history-apis.html#window)
    objects will be [same
    origin](browsers.html#same-origin){#integration-with-the-javascript-agent-formalism:same-origin},
    can reach each other directly, and
    [`document.domain`{#integration-with-the-javascript-agent-formalism:dom-document-domain-2}](browsers.html#dom-document-domain)
    will no-op.

    Two
    [`Window`{#integration-with-the-javascript-agent-formalism:window-4}](nav-history-apis.html#window)
    objects that are [same
    origin](browsers.html#same-origin){#integration-with-the-javascript-agent-formalism:same-origin-2}
    can be in different [similar-origin window
    agents](#similar-origin-window-agent){#integration-with-the-javascript-agent-formalism:similar-origin-window-agent},
    for instance if they are each in their own [browsing context
    group](document-sequences.html#browsing-context-group){#integration-with-the-javascript-agent-formalism:browsing-context-group}.

[Dedicated worker agent]{#dedicated-worker-agent .dfn export=""}
:   Contains a single
    [`DedicatedWorkerGlobalScope`{#integration-with-the-javascript-agent-formalism:dedicatedworkerglobalscope}](workers.html#dedicatedworkerglobalscope).

[Shared worker agent]{#shared-worker-agent .dfn export=""}
:   Contains a single
    [`SharedWorkerGlobalScope`{#integration-with-the-javascript-agent-formalism:sharedworkerglobalscope}](workers.html#sharedworkerglobalscope).

[Service worker agent]{#service-worker-agent .dfn export=""}
:   Contains a single
    [`ServiceWorkerGlobalScope`{#integration-with-the-javascript-agent-formalism:serviceworkerglobalscope}](https://w3c.github.io/ServiceWorker/#serviceworkerglobalscope){x-internal="serviceworkerglobalscope"}.

[Worklet agent]{#worklet-agent .dfn export=""}

:   Contains a single
    [`WorkletGlobalScope`{#integration-with-the-javascript-agent-formalism:workletglobalscope}](worklets.html#workletglobalscope)
    object.

    Although a given worklet can have multiple realms, each such realm
    needs its own agent, as each realm can be executing code
    independently and at the same time as the others.

Only
[shared](#shared-worker-agent){#integration-with-the-javascript-agent-formalism:shared-worker-agent}
and [dedicated worker
agents](#dedicated-worker-agent){#integration-with-the-javascript-agent-formalism:dedicated-worker-agent}
allow the use of JavaScript
[`Atomics`{#integration-with-the-javascript-agent-formalism:atomics}](https://tc39.es/ecma262/#sec-atomics-object){x-internal="atomics"}
APIs to potentially
[block](https://tc39.es/ecma262/#sec-forward-progress){#integration-with-the-javascript-agent-formalism:forward-progress
x-internal="forward-progress"}.

To [create an agent]{#create-an-agent .dfn}, given a boolean
`canBlock`{.variable}:

1.  Let `signifier`{.variable} be a new unique internal value.

2.  Let `candidateExecution`{.variable} be a new [candidate
    execution](https://tc39.es/ecma262/#sec-candidate-executions){#integration-with-the-javascript-agent-formalism:candidate-execution
    x-internal="candidate-execution"}.

3.  Let `agent`{.variable} be a new
    [agent](https://tc39.es/ecma262/#sec-agents){#integration-with-the-javascript-agent-formalism:agent-4
    x-internal="agent"} whose \[\[CanBlock\]\] is `canBlock`{.variable},
    \[\[Signifier\]\] is `signifier`{.variable},
    \[\[CandidateExecution\]\] is `candidateExecution`{.variable}, and
    \[\[IsLockFree1\]\], \[\[IsLockFree2\]\], and \[\[LittleEndian\]\]
    are set at the implementation\'s discretion.

4.  Set `agent`{.variable}\'s [event
    loop](#concept-agent-event-loop){#integration-with-the-javascript-agent-formalism:concept-agent-event-loop}
    to a new [event
    loop](#event-loop){#integration-with-the-javascript-agent-formalism:event-loop}.

5.  Return `agent`{.variable}.

For a
[realm](https://tc39.es/ecma262/#sec-code-realms){#integration-with-the-javascript-agent-formalism:realm
x-internal="realm"} `realm`{.variable}, the
[agent](https://tc39.es/ecma262/#sec-agents){#integration-with-the-javascript-agent-formalism:agent-5
x-internal="agent"} whose \[\[Signifier\]\] is
`realm`{.variable}.\[\[AgentSignifier\]\] is [the realm\'s
agent]{#realm's-agent .dfn dfn-for="realm" lt="agent" export=""}.

The [relevant agent]{#relevant-agent .dfn export=""} for a [platform
object](https://webidl.spec.whatwg.org/#dfn-platform-object){#integration-with-the-javascript-agent-formalism:platform-object
x-internal="platform-object"} `platformObject`{.variable} is
`platformObject`{.variable}\'s [relevant
realm](#concept-relevant-realm){#integration-with-the-javascript-agent-formalism:concept-relevant-realm}\'s
[agent](#realm's-agent){#integration-with-the-javascript-agent-formalism:realm's-agent}.

The agent equivalent of the [current
realm](https://tc39.es/ecma262/#current-realm){#integration-with-the-javascript-agent-formalism:current-realm
x-internal="current-realm"} is the [surrounding
agent](https://tc39.es/ecma262/#surrounding-agent){#integration-with-the-javascript-agent-formalism:surrounding-agent
x-internal="surrounding-agent"}.

##### [8.1.2.2]{.secno} Integration with the JavaScript agent cluster formalism[](#integration-with-the-javascript-agent-cluster-formalism){.self-link}

JavaScript also defines the concept of an [agent
cluster](https://tc39.es/ecma262/#sec-agent-clusters){#integration-with-the-javascript-agent-cluster-formalism:agent-cluster
x-internal="agent-cluster"}, which this standard maps to the web
platform by placing agents appropriately when they are created using the
[obtain a similar-origin window
agent](#obtain-similar-origin-window-agent){#integration-with-the-javascript-agent-cluster-formalism:obtain-similar-origin-window-agent}
or [obtain a worker/worklet
agent](#obtaining-a-worker/worklet-agent){#integration-with-the-javascript-agent-cluster-formalism:obtaining-a-worker/worklet-agent}
algorithms.

The [agent
cluster](https://tc39.es/ecma262/#sec-agent-clusters){#integration-with-the-javascript-agent-cluster-formalism:agent-cluster-2
x-internal="agent-cluster"} concept is crucial for defining the
JavaScript memory model, and in particular among which
[agents](https://tc39.es/ecma262/#sec-agents){#integration-with-the-javascript-agent-cluster-formalism:agent
x-internal="agent"} the backing data of
[`SharedArrayBuffer`{#integration-with-the-javascript-agent-cluster-formalism:sharedarraybuffer}](https://tc39.es/ecma262/#sec-sharedarraybuffer-objects){x-internal="sharedarraybuffer"}
objects can be shared.

Conceptually, the [agent
cluster](https://tc39.es/ecma262/#sec-agent-clusters){#integration-with-the-javascript-agent-cluster-formalism:agent-cluster-3
x-internal="agent-cluster"} concept is an architecture-independent,
idealized \"process boundary\" that groups together multiple \"threads\"
([agents](https://tc39.es/ecma262/#sec-agents){#integration-with-the-javascript-agent-cluster-formalism:agent-2
x-internal="agent"}). The [agent
clusters](https://tc39.es/ecma262/#sec-agent-clusters){#integration-with-the-javascript-agent-cluster-formalism:agent-cluster-4
x-internal="agent-cluster"} defined by the specification are generally
more restrictive than the actual process boundaries implemented in user
agents. By enforcing these idealized divisions at the specification
level, we ensure that web developers see interoperable behavior with
regard to shared memory, even in the face of varying and changing user
agent process models.

An [agent
cluster](https://tc39.es/ecma262/#sec-agent-clusters){#integration-with-the-javascript-agent-cluster-formalism:agent-cluster-5
x-internal="agent-cluster"} has an associated [cross-origin isolation
mode]{#agent-cluster-cross-origin-isolation .dfn}, which is a
[cross-origin isolation
mode](document-sequences.html#cross-origin-isolation-mode){#integration-with-the-javascript-agent-cluster-formalism:cross-origin-isolation-mode}.
It is initially
\"[`none`{#integration-with-the-javascript-agent-cluster-formalism:cross-origin-isolation-none}](document-sequences.html#cross-origin-isolation-none)\".

An [agent
cluster](https://tc39.es/ecma262/#sec-agent-clusters){#integration-with-the-javascript-agent-cluster-formalism:agent-cluster-6
x-internal="agent-cluster"} has an associated [is
origin-keyed]{#is-origin-keyed .dfn} (a boolean), which is initially
false.

------------------------------------------------------------------------

The following defines the allocation of the [agent
clusters](https://tc39.es/ecma262/#sec-agent-clusters){#integration-with-the-javascript-agent-cluster-formalism:agent-cluster-7
x-internal="agent-cluster"} of [similar-origin window
agents](#similar-origin-window-agent){#integration-with-the-javascript-agent-cluster-formalism:similar-origin-window-agent}.

An [agent cluster key]{#agent-cluster-key .dfn} is a
[site](browsers.html#site){#integration-with-the-javascript-agent-cluster-formalism:site}
or [tuple
origin](browsers.html#concept-origin-tuple){#integration-with-the-javascript-agent-cluster-formalism:concept-origin-tuple}.
Without web developer action to achieve [origin-keyed agent
clusters](browsers.html#origin-keyed-agent-clusters), it will be a
[site](browsers.html#site){#integration-with-the-javascript-agent-cluster-formalism:site-2}.

An equivalent formulation is that an [agent cluster
key](#agent-cluster-key){#integration-with-the-javascript-agent-cluster-formalism:agent-cluster-key}
can be a
[scheme-and-host](browsers.html#scheme-and-host){#integration-with-the-javascript-agent-cluster-formalism:scheme-and-host}
or an
[origin](browsers.html#concept-origin){#integration-with-the-javascript-agent-cluster-formalism:concept-origin}.

To [obtain a similar-origin window
agent]{#obtain-similar-origin-window-agent .dfn}, given an
[origin](browsers.html#concept-origin){#integration-with-the-javascript-agent-cluster-formalism:concept-origin-2}
`origin`{.variable}, a [browsing context
group](document-sequences.html#browsing-context-group){#integration-with-the-javascript-agent-cluster-formalism:browsing-context-group}
`group`{.variable}, and a boolean `requestsOAC`{.variable}, run these
steps:

1.  Let `site`{.variable} be the result of [obtaining a
    site](browsers.html#obtain-a-site){#integration-with-the-javascript-agent-cluster-formalism:obtain-a-site}
    with `origin`{.variable}.

2.  Let `key`{.variable} be `site`{.variable}.

3.  If `group`{.variable}\'s [cross-origin isolation
    mode](document-sequences.html#bcg-cross-origin-isolation){#integration-with-the-javascript-agent-cluster-formalism:bcg-cross-origin-isolation}
    is not
    \"[`none`{#integration-with-the-javascript-agent-cluster-formalism:cross-origin-isolation-none-2}](document-sequences.html#cross-origin-isolation-none)\",
    then set `key`{.variable} to `origin`{.variable}.

4.  Otherwise, if `group`{.variable}\'s [historical agent cluster key
    map](document-sequences.html#historical-agent-cluster-key-map){#integration-with-the-javascript-agent-cluster-formalism:historical-agent-cluster-key-map}\[`origin`{.variable}\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#integration-with-the-javascript-agent-cluster-formalism:map-exists
    x-internal="map-exists"}, then set `key`{.variable} to
    `group`{.variable}\'s [historical agent cluster key
    map](document-sequences.html#historical-agent-cluster-key-map){#integration-with-the-javascript-agent-cluster-formalism:historical-agent-cluster-key-map-2}\[`origin`{.variable}\].

5.  Otherwise:

    1.  If `requestsOAC`{.variable} is true, then set `key`{.variable}
        to `origin`{.variable}.

    2.  Set `group`{.variable}\'s [historical agent cluster key
        map](document-sequences.html#historical-agent-cluster-key-map){#integration-with-the-javascript-agent-cluster-formalism:historical-agent-cluster-key-map-3}\[`origin`{.variable}\]
        to `key`{.variable}.

6.  If `group`{.variable}\'s [agent cluster
    map](document-sequences.html#agent-cluster-map){#integration-with-the-javascript-agent-cluster-formalism:agent-cluster-map}\[`key`{.variable}\]
    [does not
    exist](https://infra.spec.whatwg.org/#map-exists){#integration-with-the-javascript-agent-cluster-formalism:map-exists-2
    x-internal="map-exists"}, then:

    1.  Let `agentCluster`{.variable} be a new [agent
        cluster](https://tc39.es/ecma262/#sec-agent-clusters){#integration-with-the-javascript-agent-cluster-formalism:agent-cluster-8
        x-internal="agent-cluster"}.

    2.  Set `agentCluster`{.variable}\'s [cross-origin isolation
        mode](#agent-cluster-cross-origin-isolation){#integration-with-the-javascript-agent-cluster-formalism:agent-cluster-cross-origin-isolation}
        to `group`{.variable}\'s [cross-origin isolation
        mode](document-sequences.html#bcg-cross-origin-isolation){#integration-with-the-javascript-agent-cluster-formalism:bcg-cross-origin-isolation-2}.

    3.  If `key`{.variable} is an
        [origin](browsers.html#concept-origin){#integration-with-the-javascript-agent-cluster-formalism:concept-origin-3}:

        1.  [Assert](https://infra.spec.whatwg.org/#assert){#integration-with-the-javascript-agent-cluster-formalism:assert
            x-internal="assert"}: `key`{.variable} is
            `origin`{.variable}.

        2.  Set `agentCluster`{.variable}\'s [is
            origin-keyed](#is-origin-keyed){#integration-with-the-javascript-agent-cluster-formalism:is-origin-keyed}
            to true.

    4.  Add the result of [creating an
        agent](#create-an-agent){#integration-with-the-javascript-agent-cluster-formalism:create-an-agent},
        given false, to `agentCluster`{.variable}.

    5.  Set `group`{.variable}\'s [agent cluster
        map](document-sequences.html#agent-cluster-map){#integration-with-the-javascript-agent-cluster-formalism:agent-cluster-map-2}\[`key`{.variable}\]
        to `agentCluster`{.variable}.

7.  Return the single [similar-origin window
    agent](#similar-origin-window-agent){#integration-with-the-javascript-agent-cluster-formalism:similar-origin-window-agent-2}
    contained in `group`{.variable}\'s [agent cluster
    map](document-sequences.html#agent-cluster-map){#integration-with-the-javascript-agent-cluster-formalism:agent-cluster-map-3}\[`key`{.variable}\].

This means that there is only one [similar-origin window
agent](#similar-origin-window-agent){#integration-with-the-javascript-agent-cluster-formalism:similar-origin-window-agent-3}
per browsing context agent cluster. (However, [dedicated
worker](#dedicated-worker-agent){#integration-with-the-javascript-agent-cluster-formalism:dedicated-worker-agent}
and [worklet
agents](#worklet-agent){#integration-with-the-javascript-agent-cluster-formalism:worklet-agent}
might be in the same cluster.)

------------------------------------------------------------------------

The following defines the allocation of the [agent
clusters](https://tc39.es/ecma262/#sec-agent-clusters){#integration-with-the-javascript-agent-cluster-formalism:agent-cluster-9
x-internal="agent-cluster"} of all other types of agents.

To [obtain a worker/worklet agent]{#obtaining-a-worker/worklet-agent
.dfn}, given an [environment settings
object](#environment-settings-object){#integration-with-the-javascript-agent-cluster-formalism:environment-settings-object}
or null `outside settings`{.variable}, a boolean
`isTopLevel`{.variable}, and a boolean `canBlock`{.variable}, run these
steps:

1.  Let `agentCluster`{.variable} be null.

2.  If `isTopLevel`{.variable} is true, then:

    1.  Set `agentCluster`{.variable} to a new [agent
        cluster](https://tc39.es/ecma262/#sec-agent-clusters){#integration-with-the-javascript-agent-cluster-formalism:agent-cluster-10
        x-internal="agent-cluster"}.

    2.  Set `agentCluster`{.variable}\'s [is
        origin-keyed](#is-origin-keyed){#integration-with-the-javascript-agent-cluster-formalism:is-origin-keyed-2}
        to true.

        These workers can be considered to be origin-keyed. However,
        this is not exposed through any APIs (in the way that
        [`originAgentCluster`{#integration-with-the-javascript-agent-cluster-formalism:dom-originagentcluster}](browsers.html#dom-originagentcluster)
        exposes the origin-keyedness for windows).

3.  Otherwise:

    1.  [Assert](https://infra.spec.whatwg.org/#assert){#integration-with-the-javascript-agent-cluster-formalism:assert-2
        x-internal="assert"}: `outside settings`{.variable} is not null.

    2.  Let `ownerAgent`{.variable} be `outside settings`{.variable}\'s
        [realm](#environment-settings-object's-realm){#integration-with-the-javascript-agent-cluster-formalism:environment-settings-object's-realm}\'s
        [agent](#realm's-agent){#integration-with-the-javascript-agent-cluster-formalism:realm's-agent}.

    3.  Set `agentCluster`{.variable} to the agent cluster which
        contains `ownerAgent`{.variable}.

4.  Let `agent`{.variable} be the result of [creating an
    agent](#create-an-agent){#integration-with-the-javascript-agent-cluster-formalism:create-an-agent-2}
    given `canBlock`{.variable}.

5.  Add `agent`{.variable} to `agentCluster`{.variable}.

6.  Return `agent`{.variable}.

To [obtain a dedicated/shared worker
agent]{#obtain-a-dedicated/shared-worker-agent .dfn}, given an
[environment settings
object](#environment-settings-object){#integration-with-the-javascript-agent-cluster-formalism:environment-settings-object-2}
`outside settings`{.variable} and a boolean `isShared`{.variable},
return the result of [obtaining a worker/worklet
agent](#obtaining-a-worker/worklet-agent){#integration-with-the-javascript-agent-cluster-formalism:obtaining-a-worker/worklet-agent-2}
given `outside settings`{.variable}, `isShared`{.variable}, and true.

To [obtain a worklet agent]{#obtain-a-worklet-agent .dfn}, given an
[environment settings
object](#environment-settings-object){#integration-with-the-javascript-agent-cluster-formalism:environment-settings-object-3}
`outside settings`{.variable}, return the result of [obtaining a
worker/worklet
agent](#obtaining-a-worker/worklet-agent){#integration-with-the-javascript-agent-cluster-formalism:obtaining-a-worker/worklet-agent-3}
given `outside settings`{.variable}, false, and false.

To [obtain a service worker agent]{#obtain-a-service-worker-agent .dfn
export=""}, return the result of [obtaining a worker/worklet
agent](#obtaining-a-worker/worklet-agent){#integration-with-the-javascript-agent-cluster-formalism:obtaining-a-worker/worklet-agent-4}
given null, true, and false.

------------------------------------------------------------------------

::: {#can-share-memory-with .example}
The following pairs of global objects are each within the same [agent
cluster](https://tc39.es/ecma262/#sec-agent-clusters){#integration-with-the-javascript-agent-cluster-formalism:agent-cluster-11
x-internal="agent-cluster"}, and thus can use
[`SharedArrayBuffer`{#integration-with-the-javascript-agent-cluster-formalism:sharedarraybuffer-2}](https://tc39.es/ecma262/#sec-sharedarraybuffer-objects){x-internal="sharedarraybuffer"}
instances to share memory with each other:

- A
  [`Window`{#integration-with-the-javascript-agent-cluster-formalism:window}](nav-history-apis.html#window)
  object and a dedicated worker that it created.

- A worker (of any type) and a dedicated worker it created.

- A
  [`Window`{#integration-with-the-javascript-agent-cluster-formalism:window-2}](nav-history-apis.html#window)
  object `A`{.variable} and the
  [`Window`{#integration-with-the-javascript-agent-cluster-formalism:window-3}](nav-history-apis.html#window)
  object of an
  [`iframe`{#integration-with-the-javascript-agent-cluster-formalism:the-iframe-element}](iframe-embed-object.html#the-iframe-element)
  element that `A`{.variable} created that could be [same
  origin-domain](browsers.html#same-origin-domain){#integration-with-the-javascript-agent-cluster-formalism:same-origin-domain}
  with `A`{.variable}.

- A
  [`Window`{#integration-with-the-javascript-agent-cluster-formalism:window-4}](nav-history-apis.html#window)
  object and a [same
  origin-domain](browsers.html#same-origin-domain){#integration-with-the-javascript-agent-cluster-formalism:same-origin-domain-2}
  [`Window`{#integration-with-the-javascript-agent-cluster-formalism:window-5}](nav-history-apis.html#window)
  object that opened it.

- A
  [`Window`{#integration-with-the-javascript-agent-cluster-formalism:window-6}](nav-history-apis.html#window)
  object and a worklet that it created.

The following pairs of global objects are *not* within the same [agent
cluster](https://tc39.es/ecma262/#sec-agent-clusters){#integration-with-the-javascript-agent-cluster-formalism:agent-cluster-12
x-internal="agent-cluster"}, and thus cannot share memory:

- A
  [`Window`{#integration-with-the-javascript-agent-cluster-formalism:window-7}](nav-history-apis.html#window)
  object and a shared worker it created.

- A worker (of any type) and a shared worker it created.

- A
  [`Window`{#integration-with-the-javascript-agent-cluster-formalism:window-8}](nav-history-apis.html#window)
  object and a service worker it created.

- A
  [`Window`{#integration-with-the-javascript-agent-cluster-formalism:window-9}](nav-history-apis.html#window)
  object `A`{.variable} and the
  [`Window`{#integration-with-the-javascript-agent-cluster-formalism:window-10}](nav-history-apis.html#window)
  object of an
  [`iframe`{#integration-with-the-javascript-agent-cluster-formalism:the-iframe-element-2}](iframe-embed-object.html#the-iframe-element)
  element that `A`{.variable} created that cannot be [same
  origin-domain](browsers.html#same-origin-domain){#integration-with-the-javascript-agent-cluster-formalism:same-origin-domain-3}
  with `A`{.variable}.

- Any two
  [`Window`{#integration-with-the-javascript-agent-cluster-formalism:window-11}](nav-history-apis.html#window)
  objects with no opener or ancestor relationship. This holds even if
  the two
  [`Window`{#integration-with-the-javascript-agent-cluster-formalism:window-12}](nav-history-apis.html#window)
  objects are [same
  origin](browsers.html#same-origin){#integration-with-the-javascript-agent-cluster-formalism:same-origin}.
:::

#### [8.1.3]{.secno} Realms and their counterparts[](#realms-and-their-counterparts){.self-link}

The JavaScript specification introduces the
[realm](https://tc39.es/ecma262/#sec-code-realms){#realms-and-their-counterparts:realm
x-internal="realm"} concept, representing a global environment in which
script is run. Each realm comes with an
[implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#realms-and-their-counterparts:implementation-defined
x-internal="implementation-defined"} [global
object](#global-object){#realms-and-their-counterparts:global-object};
much of this specification is devoted to defining that global object and
its properties.

For web specifications, it is often useful to associate values or
algorithms with a realm/global object pair. When the values are specific
to a particular type of realm, they are associated directly with the
global object in question, e.g., in the definition of the
[`Window`{#realms-and-their-counterparts:window}](nav-history-apis.html#window)
or
[`WorkerGlobalScope`{#realms-and-their-counterparts:workerglobalscope}](workers.html#workerglobalscope)
interfaces. When the values have utility across multiple realms, we use
the [environment settings
object](#environment-settings-object){#realms-and-their-counterparts:environment-settings-object}
concept.

Finally, in some cases it is necessary to track associated values before
a realm/global object/environment settings object even comes into
existence (for example, during
[navigation](browsing-the-web.html#navigate){#realms-and-their-counterparts:navigate}).
These values are tracked in the
[environment](#environment){#realms-and-their-counterparts:environment}
concept.

##### [8.1.3.1]{.secno} Environments[](#environments){.self-link}

An [environment]{#environment .dfn export=""} is an object that
identifies the settings of a current or potential execution environment
(i.e., a [navigation
params](browsing-the-web.html#navigation-params){#environments:navigation-params}\'s
[reserved
environment](browsing-the-web.html#navigation-params-reserved-environment){#environments:navigation-params-reserved-environment}
or a
[request](https://fetch.spec.whatwg.org/#concept-request){#environments:concept-request
x-internal="concept-request"}\'s [reserved
client](https://fetch.spec.whatwg.org/#concept-request-reserved-client){#environments:concept-request-reserved-client
x-internal="concept-request-reserved-client"}). An
[environment](#environment){#environments:environment} has the following
fields:

An [id]{#concept-environment-id .dfn dfn-for="environment" export=""}
:   An opaque string that uniquely identifies this
    [environment](#environment){#environments:environment-2}.

A [creation URL]{#concept-environment-creation-url .dfn dfn-for="environment" export=""}

:   A [URL](https://url.spec.whatwg.org/#concept-url){#environments:url
    x-internal="url"} that represents the location of the resource with
    which this [environment](#environment){#environments:environment-3}
    is associated.

    In the case of a
    [`Window`{#environments:window}](nav-history-apis.html#window)
    [environment settings
    object](#environment-settings-object){#environments:environment-settings-object},
    this URL might be distinct from its [global
    object](#concept-settings-object-global){#environments:concept-settings-object-global}\'s
    [associated
    `Document`](nav-history-apis.html#concept-document-window){#environments:concept-document-window}\'s
    [URL](https://dom.spec.whatwg.org/#concept-document-url){#environments:the-document's-address
    x-internal="the-document's-address"}, due to mechanisms such as
    [`history.pushState()`{#environments:dom-history-pushstate}](nav-history-apis.html#dom-history-pushstate)
    which modify the latter.

A [top-level creation URL]{#concept-environment-top-level-creation-url .dfn dfn-for="environment" export=""}
:   Null or a
    [URL](https://url.spec.whatwg.org/#concept-url){#environments:url-2
    x-internal="url"} that represents the [creation
    URL](#concept-environment-creation-url){#environments:concept-environment-creation-url}
    of the \"top-level\"
    [environment](#environment){#environments:environment-4}. It is null
    for workers and worklets.

A [top-level origin]{#concept-environment-top-level-origin .dfn dfn-for="environment" export=""}

:   A [for now]{.XXX}
    [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#environments:implementation-defined
    x-internal="implementation-defined"} value, null, or an
    [origin](browsers.html#concept-origin){#environments:concept-origin}.
    For a \"top-level\" potential execution environment it is null
    (i.e., when there is no response yet); otherwise it is the
    \"top-level\"
    [environment](#environment){#environments:environment-5}\'s
    [origin](#concept-settings-object-origin){#environments:concept-settings-object-origin}.
    For a dedicated worker or worklet it is the [top-level
    origin](#concept-environment-top-level-origin){#environments:concept-environment-top-level-origin}
    of its creator. For a shared or service worker it is an
    [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#environments:implementation-defined-2
    x-internal="implementation-defined"} value.

    This is distinct from the [top-level creation
    URL](#concept-environment-top-level-creation-url){#environments:concept-environment-top-level-creation-url}\'s
    [origin](https://url.spec.whatwg.org/#concept-url-origin){#environments:concept-url-origin
    x-internal="concept-url-origin"} when sandboxing, workers, and
    worklets are involved.

A [target browsing context]{#concept-environment-target-browsing-context .dfn dfn-for="environment" export=""}
:   Null or a target [browsing
    context](document-sequences.html#browsing-context){#environments:browsing-context}
    for a [navigation
    request](https://fetch.spec.whatwg.org/#navigation-request){#environments:navigation-request
    x-internal="navigation-request"}.

An [active service worker]{#concept-environment-active-service-worker .dfn dfn-for="environment" export=""}
:   Null or a [service
    worker](https://w3c.github.io/ServiceWorker/#dfn-service-worker){#environments:dfn-service-worker
    x-internal="dfn-service-worker"} that
    [controls](https://w3c.github.io/ServiceWorker/#dfn-control){#environments:dfn-control
    x-internal="dfn-control"} the
    [environment](#environment){#environments:environment-6}.

An [execution ready flag]{#concept-environment-execution-ready-flag .dfn dfn-for="environment" export=""}

:   A flag that indicates whether the environment setup is done. It is
    initially unset.

Specifications may define [environment discarding
steps]{#environment-discarding-steps .dfn export=""} for environments.
The steps take an
[environment](#environment){#environments:environment-7} as input.

The [environment discarding
steps](#environment-discarding-steps){#environments:environment-discarding-steps}
are run for only a select few environments: the ones that will never
become execution ready because, for example, they failed to load.

##### [8.1.3.2]{.secno} Environment settings objects[](#environment-settings-objects){.self-link}

An [environment settings object]{#environment-settings-object .dfn
export=""} is an
[environment](#environment){#environment-settings-objects:environment}
that additionally specifies algorithms for:

A [realm execution context]{#realm-execution-context .dfn dfn-for="environment settings object" export=""}

:   A [JavaScript execution
    context](https://tc39.es/ecma262/#sec-execution-contexts){#environment-settings-objects:javascript-execution-context
    x-internal="javascript-execution-context"} shared by all
    [scripts](scripting.html#the-script-element){#environment-settings-objects:the-script-element}
    that use this settings object, i.e., all scripts in a given
    [realm](https://tc39.es/ecma262/#sec-code-realms){#environment-settings-objects:realm
    x-internal="realm"}. When we [run a classic
    script](#run-a-classic-script){#environment-settings-objects:run-a-classic-script}
    or [run a module
    script](#run-a-module-script){#environment-settings-objects:run-a-module-script},
    this execution context becomes the top of the [JavaScript execution
    context
    stack](https://tc39.es/ecma262/#execution-context-stack){#environment-settings-objects:javascript-execution-context-stack
    x-internal="javascript-execution-context-stack"}, on top of which
    another execution context specific to the script in question is
    pushed. (This setup ensures [Source Text Module
    Record](https://tc39.es/ecma262/#sec-source-text-module-records){#environment-settings-objects:source-text-module-record
    x-internal="source-text-module-record"}\'s
    [Evaluate](https://tc39.es/ecma262/#sec-moduleevaluation){#environment-settings-objects:js-evaluate
    x-internal="js-evaluate"} knows which realm to use.)

A [module map]{#concept-settings-object-module-map .dfn dfn-for="environment
   settings object" export=""}

:   A [module
    map](#module-map){#environment-settings-objects:module-map} that is
    used when importing JavaScript modules.

An [API base URL]{#api-base-url .dfn dfn-for="environment settings object" export=""}

:   A
    [URL](https://url.spec.whatwg.org/#concept-url){#environment-settings-objects:url
    x-internal="url"} used by APIs called by scripts that use this
    [environment settings
    object](#environment-settings-object){#environment-settings-objects:environment-settings-object}
    to [parse
    URLs](urls-and-fetching.html#parse-a-url){#environment-settings-objects:parse-a-url}.

An [origin]{#concept-settings-object-origin .dfn dfn-for="environment
   settings object" export=""}

:   An
    [origin](browsers.html#concept-origin){#environment-settings-objects:concept-origin}
    used in security checks.

A [policy container]{#concept-settings-object-policy-container .dfn dfn-for="environment settings object" export=""}

:   A [policy
    container](browsers.html#policy-container){#environment-settings-objects:policy-container}
    containing policies used for security checks.

A [cross-origin isolated capability]{#concept-settings-object-cross-origin-isolated-capability .dfn dfn-for="environment settings object" export=""}
:   A boolean representing whether scripts that use this [environment
    settings
    object](#environment-settings-object){#environment-settings-objects:environment-settings-object-2}
    are allowed to use APIs that require cross-origin isolation.

A [time origin]{#concept-settings-object-time-origin .dfn dfn-for="environment settings object" export=""}
:   A number used as the baseline for performance-related timestamps.
    [\[HRT\]](references.html#refsHRT)

An [environment settings
object](#environment-settings-object){#environment-settings-objects:environment-settings-object-3}\'s
[responsible event loop]{#responsible-event-loop .dfn
dfn-for="environment settings
  object" export=""} is its [global
object](#concept-settings-object-global){#environment-settings-objects:concept-settings-object-global}\'s
[relevant
agent](#relevant-agent){#environment-settings-objects:relevant-agent}\'s
[event
loop](#concept-agent-event-loop){#environment-settings-objects:concept-agent-event-loop}.

##### [8.1.3.3]{.secno} Realms, settings objects, and global objects[](#realms-settings-objects-global-objects){.self-link} {#realms-settings-objects-global-objects}

A [global object]{#global-object .dfn export=""} is a JavaScript object
that is the \[\[GlobalObject\]\] field of a
[realm](https://tc39.es/ecma262/#sec-code-realms){#realms-settings-objects-global-objects:realm
x-internal="realm"}.

In this specification, all
[realms](https://tc39.es/ecma262/#sec-code-realms){#realms-settings-objects-global-objects:realm-2
x-internal="realm"} are
[created](#creating-a-new-javascript-realm){#realms-settings-objects-global-objects:creating-a-new-javascript-realm}
with [global
objects](#global-object){#realms-settings-objects-global-objects:global-object}
that are either
[`Window`{#realms-settings-objects-global-objects:window}](nav-history-apis.html#window),
[`WorkerGlobalScope`{#realms-settings-objects-global-objects:workerglobalscope}](workers.html#workerglobalscope),
or
[`WorkletGlobalScope`{#realms-settings-objects-global-objects:workletglobalscope}](worklets.html#workletglobalscope)
objects.

A [global
object](#global-object){#realms-settings-objects-global-objects:global-object-2}
has an [in error reporting mode]{#in-error-reporting-mode .dfn
dfn-for="global object"} boolean, which is initially false.

A [global
object](#global-object){#realms-settings-objects-global-objects:global-object-3}
has an [outstanding rejected promises weak
set]{#outstanding-rejected-promises-weak-set .dfn}, a
[set](https://infra.spec.whatwg.org/#ordered-set){#realms-settings-objects-global-objects:set
x-internal="set"} of
[`Promise`{#realms-settings-objects-global-objects:idl-promise}](https://webidl.spec.whatwg.org/#idl-promise){x-internal="idl-promise"}
objects, initially empty. This set must not create strong references to
any of its members, and implementations are free to limit its size in an
[implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#realms-settings-objects-global-objects:implementation-defined
x-internal="implementation-defined"} manner, e.g., by removing old
entries from it when new ones are added.

A [global
object](#global-object){#realms-settings-objects-global-objects:global-object-4}
has an [about-to-be-notified rejected promises
list]{#about-to-be-notified-rejected-promises-list .dfn}, a
[list](https://infra.spec.whatwg.org/#list){#realms-settings-objects-global-objects:list
x-internal="list"} of
[`Promise`{#realms-settings-objects-global-objects:idl-promise-2}](https://webidl.spec.whatwg.org/#idl-promise){x-internal="idl-promise"}
objects, initially empty.

A [global
object](#global-object){#realms-settings-objects-global-objects:global-object-5}
has an [import map]{#concept-global-import-map .dfn}, initially an
[empty import
map](#empty-import-map){#realms-settings-objects-global-objects:empty-import-map}.

For now, only
[`Window`{#realms-settings-objects-global-objects:window-2}](nav-history-apis.html#window)
[global
objects](#global-object){#realms-settings-objects-global-objects:global-object-6}
have their [import
map](#concept-global-import-map){#realms-settings-objects-global-objects:concept-global-import-map}
modified from the initial empty one. The [import
map](#concept-global-import-map){#realms-settings-objects-global-objects:concept-global-import-map-2}
is only accessed for the resolution of a root [module
script](#module-script){#realms-settings-objects-global-objects:module-script}.

A [global
object](#global-object){#realms-settings-objects-global-objects:global-object-7}
has a [resolved module set]{#resolved-module-set .dfn}, a
[set](https://infra.spec.whatwg.org/#ordered-set){#realms-settings-objects-global-objects:set-2
x-internal="set"} of [specifier resolution
records](#specifier-resolution-record){#realms-settings-objects-global-objects:specifier-resolution-record},
initially empty.

The [resolved module
set](#resolved-module-set){#realms-settings-objects-global-objects:resolved-module-set}
ensures that module specifier resolution returns the same result when
called multiple times with the same (referrer, specifier) pair. It does
that by ensuring that [import
map](#import-map){#realms-settings-objects-global-objects:import-map}
rules that impact the specifier in its referrer\'s scope cannot be
defined after its initial resolution. For now, only
[`Window`{#realms-settings-objects-global-objects:window-3}](nav-history-apis.html#window)
[global
objects](#global-object){#realms-settings-objects-global-objects:global-object-8}
have their module set data structures modified from the initial empty
one.

------------------------------------------------------------------------

There is always a 1-to-1-to-1 mapping between
[realms](https://tc39.es/ecma262/#sec-code-realms){#realms-settings-objects-global-objects:realm-3
x-internal="realm"}, [global
objects](#global-object){#realms-settings-objects-global-objects:global-object-9},
and [environment settings
objects](#environment-settings-object){#realms-settings-objects-global-objects:environment-settings-object}:

- A
  [realm](https://tc39.es/ecma262/#sec-code-realms){#realms-settings-objects-global-objects:realm-4
  x-internal="realm"} has a \[\[HostDefined\]\] field, which contains
  [the realm\'s settings object]{#concept-realm-settings-object .dfn
  dfn-for="realm" lt="settings object" export=""}.

- A
  [realm](https://tc39.es/ecma262/#sec-code-realms){#realms-settings-objects-global-objects:realm-5
  x-internal="realm"} has a \[\[GlobalObject\]\] field, which contains
  [the realm\'s global object]{#concept-realm-global .dfn
  dfn-for="realm" lt="global object" export=""}.

- Each [global
  object](#global-object){#realms-settings-objects-global-objects:global-object-10}
  in this specification is created during the
  [creation](#creating-a-new-javascript-realm){#realms-settings-objects-global-objects:creating-a-new-javascript-realm-2}
  of a corresponding
  [realm](https://tc39.es/ecma262/#sec-code-realms){#realms-settings-objects-global-objects:realm-6
  x-internal="realm"}, known as [the global object\'s
  realm]{#concept-global-object-realm .dfn dfn-for="global object"
  lt="realm" export=""}.

- ::: {#relevant-settings-object-for-a-global-object}
  Each [global
  object](#global-object){#realms-settings-objects-global-objects:global-object-11}
  in this specification is created alongside a corresponding
  [environment settings
  object](#environment-settings-object){#realms-settings-objects-global-objects:environment-settings-object-2},
  known as its [relevant settings
  object](#relevant-settings-object){#realms-settings-objects-global-objects:relevant-settings-object}.
  :::

- An [environment settings
  object](#environment-settings-object){#realms-settings-objects-global-objects:environment-settings-object-3}\'s
  [realm execution
  context](#realm-execution-context){#realms-settings-objects-global-objects:realm-execution-context}\'s
  Realm component is [the environment settings object\'s
  realm]{#environment-settings-object's-realm .dfn
  dfn-for="environment settings object" lt="realm" export=""}.

- An [environment settings
  object](#environment-settings-object){#realms-settings-objects-global-objects:environment-settings-object-4}\'s
  [realm](#environment-settings-object's-realm){#realms-settings-objects-global-objects:environment-settings-object's-realm}
  then has a \[\[GlobalObject\]\] field, which contains [the environment
  settings object\'s global object]{#concept-settings-object-global .dfn
  dfn-for="environment settings
     object" lt="global object" export=""}.

To [create a new realm]{#creating-a-new-javascript-realm .dfn export=""}
in an
[agent](https://tc39.es/ecma262/#sec-agents){#realms-settings-objects-global-objects:agent
x-internal="agent"} `agent`{.variable}, optionally with instructions to
create a global object or a global **this** binding (or both), the
following steps are taken:

1.  Perform
    [InitializeHostDefinedRealm](https://tc39.es/ecma262/#sec-initializehostdefinedrealm){#realms-settings-objects-global-objects:js-initializehostdefinedrealm
    x-internal="js-initializehostdefinedrealm"}() with the provided
    customizations for creating the global object and the global
    **this** binding.

2.  Let `realm execution context`{.variable} be the [running JavaScript
    execution
    context](https://tc39.es/ecma262/#running-execution-context){#realms-settings-objects-global-objects:running-javascript-execution-context
    x-internal="running-javascript-execution-context"}.

    This is the [JavaScript execution
    context](https://tc39.es/ecma262/#sec-execution-contexts){#realms-settings-objects-global-objects:javascript-execution-context
    x-internal="javascript-execution-context"} created in the previous
    step.

3.  Remove `realm execution context`{.variable} from the [JavaScript
    execution context
    stack](https://tc39.es/ecma262/#execution-context-stack){#realms-settings-objects-global-objects:javascript-execution-context-stack
    x-internal="javascript-execution-context-stack"}.

4.  Let `realm`{.variable} be `realm execution context`{.variable}\'s
    Realm component.

5.  If `agent`{.variable}\'s [agent
    cluster](https://tc39.es/ecma262/#sec-agent-clusters){#realms-settings-objects-global-objects:agent-cluster
    x-internal="agent-cluster"}\'s [cross-origin isolation
    mode](#agent-cluster-cross-origin-isolation){#realms-settings-objects-global-objects:agent-cluster-cross-origin-isolation}
    is
    \"[`none`{#realms-settings-objects-global-objects:cross-origin-isolation-none}](document-sequences.html#cross-origin-isolation-none)\",
    then:

    1.  Let `global`{.variable} be `realm`{.variable}\'s [global
        object](#concept-realm-global){#realms-settings-objects-global-objects:concept-realm-global}.

    2.  Let `status`{.variable} be !
        `global`{.variable}.\[\[Delete\]\](\"`SharedArrayBuffer`\").

    3.  [Assert](https://infra.spec.whatwg.org/#assert){#realms-settings-objects-global-objects:assert
        x-internal="assert"}: `status`{.variable} is true.

    This is done for compatibility with web content and there is some
    hope that this can be removed in the future. Web developers can
    still get at the constructor through
    `new WebAssembly.Memory({ shared:true, initial:0, maximum:0 }).buffer.constructor`{.js}.

6.  Return `realm execution context`{.variable}.

------------------------------------------------------------------------

When defining algorithm steps throughout this specification, it is often
important to indicate what
[realm](https://tc39.es/ecma262/#sec-code-realms){#realms-settings-objects-global-objects:realm-7
x-internal="realm"} is to be used---or, equivalently, what [global
object](#global-object){#realms-settings-objects-global-objects:global-object-12}
or [environment settings
object](#environment-settings-object){#realms-settings-objects-global-objects:environment-settings-object-5}
is to be used. In general, there are at least four possibilities:

[Entry]{#concept-entry-everything .dfn}
:   This corresponds to the script that initiated the currently running
    script action: i.e., the function or script that the user agent
    called into when it called into author code.

[Incumbent]{#concept-incumbent-everything .dfn}
:   This corresponds to the most-recently-entered author function or
    script on the stack, or the author function or script that
    originally scheduled the currently-running callback.

[Current]{#concept-current-everything .dfn}
:   This corresponds to the currently-running function object, including
    built-in user-agent functions which might not be implemented as
    JavaScript. (It is derived from the [current
    realm](https://tc39.es/ecma262/#current-realm){#realms-settings-objects-global-objects:current-realm
    x-internal="current-realm"}.)

[Relevant]{#concept-relevant-everything .dfn}
:   Every [platform
    object](https://webidl.spec.whatwg.org/#dfn-platform-object){#realms-settings-objects-global-objects:platform-object
    x-internal="platform-object"} has a [relevant
    realm](#concept-relevant-realm){#realms-settings-objects-global-objects:concept-relevant-realm},
    which is roughly the
    [realm](https://tc39.es/ecma262/#sec-code-realms){#realms-settings-objects-global-objects:realm-8
    x-internal="realm"} in which it was created. When writing
    algorithms, the most prominent [platform
    object](https://webidl.spec.whatwg.org/#dfn-platform-object){#realms-settings-objects-global-objects:platform-object-2
    x-internal="platform-object"} whose [relevant
    realm](#concept-relevant-realm){#realms-settings-objects-global-objects:concept-relevant-realm-2}
    might be important is the **this** value of the currently-running
    function object. In some cases, there can be other important
    [relevant
    realms](#concept-relevant-realm){#realms-settings-objects-global-objects:concept-relevant-realm-3},
    such as those of any arguments.

Note how the
[entry](#concept-entry-everything){#realms-settings-objects-global-objects:concept-entry-everything},
[incumbent](#concept-incumbent-everything){#realms-settings-objects-global-objects:concept-incumbent-everything},
and
[current](#concept-current-everything){#realms-settings-objects-global-objects:concept-current-everything}
concepts are usable without qualification, whereas the
[relevant](#concept-relevant-everything){#realms-settings-objects-global-objects:concept-relevant-everything}
concept must be applied to a particular [platform
object](https://webidl.spec.whatwg.org/#dfn-platform-object){#realms-settings-objects-global-objects:platform-object-3
x-internal="platform-object"}.

The
[incumbent](#concept-incumbent-everything){#realms-settings-objects-global-objects:concept-incumbent-everything-2}
and
[entry](#concept-entry-everything){#realms-settings-objects-global-objects:concept-entry-everything-2}
concepts should not be used by new specifications, as they are
excessively complicated and unintuitive to work with. We are working to
remove almost all existing uses from the platform: see [issue
#1430](https://github.com/whatwg/html/issues/1430) for
[incumbent](#concept-incumbent-everything){#realms-settings-objects-global-objects:concept-incumbent-everything-3},
and [issue #1431](https://github.com/whatwg/html/issues/1431) for
[entry](#concept-entry-everything){#realms-settings-objects-global-objects:concept-entry-everything-3}.

In general, web platform specifications should use the
[relevant](#concept-relevant-everything){#realms-settings-objects-global-objects:concept-relevant-everything-2}
concept, applied to the object being operated on (usually the **this**
value of the current method). This mismatches the JavaScript
specification, where
[current](#concept-current-everything){#realms-settings-objects-global-objects:concept-current-everything-2}
is generally used as the default (e.g., in determining the
[realm](https://tc39.es/ecma262/#sec-code-realms){#realms-settings-objects-global-objects:realm-9
x-internal="realm"} whose `Array` constructor should be used to
construct the result in `Array.prototype.map`). But this inconsistency
is so embedded in the platform that we have to accept it going forward.

::: example
Consider the following pages, with `a.html` being loaded in a browser
window, `b.html` being loaded in an
[`iframe`{#realms-settings-objects-global-objects:the-iframe-element}](iframe-embed-object.html#the-iframe-element)
as shown, and `c.html` and `d.html` omitted (they can simply be empty
documents):

``` html
<!-- a.html -->
<!DOCTYPE html>
<html lang="en">
<title>Entry page</title>

<iframe src="b.html"></iframe>
<button onclick="frames[0].hello()">Hello</button>

<!--b.html -->
<!DOCTYPE html>
<html lang="en">
<title>Incumbent page</title>

<iframe src="c.html" id="c"></iframe>
<iframe src="d.html" id="d"></iframe>

<script>
  const c = document.querySelector("#c").contentWindow;
  const d = document.querySelector("#d").contentWindow;

  window.hello = () => {
    c.print.call(d);
  };
</script>
```

Each page has its own [browsing
context](document-sequences.html#browsing-context){#realms-settings-objects-global-objects:browsing-context},
and thus its own
[realm](https://tc39.es/ecma262/#sec-code-realms){#realms-settings-objects-global-objects:realm-10
x-internal="realm"}, [global
object](#global-object){#realms-settings-objects-global-objects:global-object-13},
and [environment settings
object](#environment-settings-object){#realms-settings-objects-global-objects:environment-settings-object-6}.

When the
[`print()`{#realms-settings-objects-global-objects:dom-print}](timers-and-user-prompts.html#dom-print)
method is called in response to pressing the button in `a.html`, then:

- The [entry
  realm](#concept-entry-realm){#realms-settings-objects-global-objects:concept-entry-realm}
  is that of `a.html`.

- The [incumbent
  realm](#concept-incumbent-realm){#realms-settings-objects-global-objects:concept-incumbent-realm}
  is that of `b.html`.

- The [current
  realm](https://tc39.es/ecma262/#current-realm){#realms-settings-objects-global-objects:current-realm-2
  x-internal="current-realm"} is that of `c.html` (since it is the
  [`print()`{#realms-settings-objects-global-objects:dom-print-2}](timers-and-user-prompts.html#dom-print)
  method from `c.html` whose code is running).

- The [relevant
  realm](#concept-relevant-realm){#realms-settings-objects-global-objects:concept-relevant-realm-4}
  of the object on which the
  [`print()`{#realms-settings-objects-global-objects:dom-print-3}](timers-and-user-prompts.html#dom-print)
  method is being called is that of `d.html`.
:::

::: example
One reason why the
[relevant](#concept-relevant-everything){#realms-settings-objects-global-objects:concept-relevant-everything-3}
concept is generally a better default choice than the
[current](#concept-current-everything){#realms-settings-objects-global-objects:concept-current-everything-3}
concept is that it is more suitable for creating an object that is to be
persisted and returned multiple times. For example, the
[`navigator.getBattery()`{#realms-settings-objects-global-objects:dom-navigator-getbattery}](https://w3c.github.io/battery/#widl-Navigator-getBattery-Promise-BatteryManager){x-internal="dom-navigator-getbattery"}
method creates promises in the [relevant
realm](#concept-relevant-realm){#realms-settings-objects-global-objects:concept-relevant-realm-5}
for the
[`Navigator`{#realms-settings-objects-global-objects:navigator}](system-state.html#navigator)
object on which it is invoked. This has the following impact:
[\[BATTERY\]](references.html#refsBATTERY)

``` html
<!-- outer.html -->
<!DOCTYPE html>
<html lang="en">
<title>Relevant realm demo: outer page</title>
<script>
  function doTest() {
    const promise = navigator.getBattery.call(frames[0].navigator);

    console.log(promise instanceof Promise);           // logs false
    console.log(promise instanceof frames[0].Promise); // logs true

    frames[0].hello();
  }
</script>
<iframe src="inner.html" onload="doTest()"></iframe>

<!-- inner.html -->
<!DOCTYPE html>
<html lang="en">
<title>Relevant realm demo: inner page</title>
<script>
  function hello() {
    const promise = navigator.getBattery();

    console.log(promise instanceof Promise);        // logs true
    console.log(promise instanceof parent.Promise); // logs false
  }
</script>
```

If the algorithm for the
[`getBattery()`{#realms-settings-objects-global-objects:dom-navigator-getbattery-2}](https://w3c.github.io/battery/#widl-Navigator-getBattery-Promise-BatteryManager){x-internal="dom-navigator-getbattery"}
method had instead used the [current
realm](https://tc39.es/ecma262/#current-realm){#realms-settings-objects-global-objects:current-realm-3
x-internal="current-realm"}, all the results would be reversed. That is,
after the first call to
[`getBattery()`{#realms-settings-objects-global-objects:dom-navigator-getbattery-3}](https://w3c.github.io/battery/#widl-Navigator-getBattery-Promise-BatteryManager){x-internal="dom-navigator-getbattery"}
in `outer.html`, the
[`Navigator`{#realms-settings-objects-global-objects:navigator-2}](system-state.html#navigator)
object in `inner.html` would be permanently storing a `Promise` object
created in `outer.html`\'s
[realm](https://tc39.es/ecma262/#sec-code-realms){#realms-settings-objects-global-objects:realm-11
x-internal="realm"}, and calls like that inside the `hello()` function
would thus return a promise from the \"wrong\" realm. Since this is
undesirable, the algorithm instead uses the [relevant
realm](#concept-relevant-realm){#realms-settings-objects-global-objects:concept-relevant-realm-6},
giving the sensible results indicated in the comments above.
:::

------------------------------------------------------------------------

The rest of this section deals with formally defining the
[entry](#concept-entry-everything){#realms-settings-objects-global-objects:concept-entry-everything-4},
[incumbent](#concept-incumbent-everything){#realms-settings-objects-global-objects:concept-incumbent-everything-4},
[current](#concept-current-everything){#realms-settings-objects-global-objects:concept-current-everything-4},
and
[relevant](#concept-relevant-everything){#realms-settings-objects-global-objects:concept-relevant-everything-4}
concepts.

###### [8.1.3.3.1]{.secno} Entry[](#entry){.self-link}

The process of [calling scripts](#calling-scripts) will push or pop
[realm execution
contexts](#realm-execution-context){#entry:realm-execution-context} onto
the [JavaScript execution context
stack](https://tc39.es/ecma262/#execution-context-stack){#entry:javascript-execution-context-stack
x-internal="javascript-execution-context-stack"}, interspersed with
other [execution
contexts](https://tc39.es/ecma262/#sec-execution-contexts){#entry:javascript-execution-context
x-internal="javascript-execution-context"}.

With this in hand, we define the [entry execution
context]{#entry-execution-context .dfn} to be the most recently pushed
item in the [JavaScript execution context
stack](https://tc39.es/ecma262/#execution-context-stack){#entry:javascript-execution-context-stack-2
x-internal="javascript-execution-context-stack"} that is a [realm
execution
context](#realm-execution-context){#entry:realm-execution-context-2}.
The [entry realm]{#concept-entry-realm .dfn} is the [entry execution
context](#entry-execution-context){#entry:entry-execution-context}\'s
Realm component.

Then, the [entry settings object]{#entry-settings-object .dfn} is the
[environment settings
object](#concept-realm-settings-object){#entry:concept-realm-settings-object}
of the [entry realm](#concept-entry-realm){#entry:concept-entry-realm}.

Similarly, the [entry global object]{#entry-global-object .dfn} is the
[global object](#concept-realm-global){#entry:concept-realm-global} of
the [entry realm](#concept-entry-realm){#entry:concept-entry-realm-2}.

###### [8.1.3.3.2]{.secno} Incumbent[](#incumbent){.self-link}

All [JavaScript execution
contexts](https://tc39.es/ecma262/#sec-execution-contexts){#incumbent:javascript-execution-context
x-internal="javascript-execution-context"} must contain, as part of
their code evaluation state, a [skip-when-determining-incumbent
counter]{#skip-when-determining-incumbent-counter .dfn} value, which is
initially zero. In the process of [preparing to run a
callback](#prepare-to-run-a-callback){#incumbent:prepare-to-run-a-callback}
and [cleaning up after running a
callback](#clean-up-after-running-a-callback){#incumbent:clean-up-after-running-a-callback},
this value will be incremented and decremented.

Every [event loop](#event-loop){#incumbent:event-loop} has an associated
[backup incumbent settings object
stack]{#backup-incumbent-settings-object-stack .dfn}, initially empty.
Roughly speaking, it is used to determine the [incumbent settings
object](#incumbent-settings-object){#incumbent:incumbent-settings-object}
when no author code is on the stack, but author code is responsible for
the current algorithm having been run in some way. The process of
[preparing to run a
callback](#prepare-to-run-a-callback){#incumbent:prepare-to-run-a-callback-2}
and [cleaning up after running a
callback](#clean-up-after-running-a-callback){#incumbent:clean-up-after-running-a-callback-2}
manipulate this stack. [\[WEBIDL\]](references.html#refsWEBIDL)

When Web IDL is used to
[invoke](https://webidl.spec.whatwg.org/#invoke-a-callback-function){#incumbent:es-invoking-callback-functions
x-internal="es-invoking-callback-functions"} author code, or when
[HostEnqueuePromiseJob](#hostenqueuepromisejob){#incumbent:hostenqueuepromisejob}
invokes a promise job, they use the following algorithms to track
relevant data for determining the [incumbent settings
object](#incumbent-settings-object){#incumbent:incumbent-settings-object-2}:

To [prepare to run a callback]{#prepare-to-run-a-callback .dfn
export=""} with an [environment settings
object](#environment-settings-object){#incumbent:environment-settings-object}
`settings`{.variable}:

1.  Push `settings`{.variable} onto the [backup incumbent settings
    object
    stack](#backup-incumbent-settings-object-stack){#incumbent:backup-incumbent-settings-object-stack}.

2.  Let `context`{.variable} be the [topmost script-having execution
    context](#topmost-script-having-execution-context){#incumbent:topmost-script-having-execution-context}.

3.  If `context`{.variable} is not null, increment
    `context`{.variable}\'s [skip-when-determining-incumbent
    counter](#skip-when-determining-incumbent-counter){#incumbent:skip-when-determining-incumbent-counter}.

To [clean up after running a
callback]{#clean-up-after-running-a-callback .dfn export=""} with an
[environment settings
object](#environment-settings-object){#incumbent:environment-settings-object-2}
`settings`{.variable}:

1.  Let `context`{.variable} be the [topmost script-having execution
    context](#topmost-script-having-execution-context){#incumbent:topmost-script-having-execution-context-2}.

    This will be the same as the [topmost script-having execution
    context](#topmost-script-having-execution-context){#incumbent:topmost-script-having-execution-context-3}
    inside the corresponding invocation of [prepare to run a
    callback](#prepare-to-run-a-callback){#incumbent:prepare-to-run-a-callback-3}.

2.  If `context`{.variable} is not null, decrement
    `context`{.variable}\'s [skip-when-determining-incumbent
    counter](#skip-when-determining-incumbent-counter){#incumbent:skip-when-determining-incumbent-counter-2}.

3.  [Assert](https://infra.spec.whatwg.org/#assert){#incumbent:assert
    x-internal="assert"}: the topmost entry of the [backup incumbent
    settings object
    stack](#backup-incumbent-settings-object-stack){#incumbent:backup-incumbent-settings-object-stack-2}
    is `settings`{.variable}.

4.  Remove `settings`{.variable} from the [backup incumbent settings
    object
    stack](#backup-incumbent-settings-object-stack){#incumbent:backup-incumbent-settings-object-stack-3}.

Here, the [topmost script-having execution
context]{#topmost-script-having-execution-context .dfn} is the topmost
entry of the [JavaScript execution context
stack](https://tc39.es/ecma262/#execution-context-stack){#incumbent:javascript-execution-context-stack
x-internal="javascript-execution-context-stack"} that has a non-null
ScriptOrModule component, or null if there is no such entry in the
[JavaScript execution context
stack](https://tc39.es/ecma262/#execution-context-stack){#incumbent:javascript-execution-context-stack-2
x-internal="javascript-execution-context-stack"}.

With all this in place, the [incumbent settings
object]{#incumbent-settings-object .dfn export=""} is determined as
follows:

1.  Let `context`{.variable} be the [topmost script-having execution
    context](#topmost-script-having-execution-context){#incumbent:topmost-script-having-execution-context-4}.

2.  If `context`{.variable} is null, or if `context`{.variable}\'s
    [skip-when-determining-incumbent
    counter](#skip-when-determining-incumbent-counter){#incumbent:skip-when-determining-incumbent-counter-3}
    is greater than zero, then:

    1.  [Assert](https://infra.spec.whatwg.org/#assert){#incumbent:assert-2
        x-internal="assert"}: the [backup incumbent settings object
        stack](#backup-incumbent-settings-object-stack){#incumbent:backup-incumbent-settings-object-stack-4}
        is not empty.

        This assert would fail if you try to obtain the [incumbent
        settings
        object](#incumbent-settings-object){#incumbent:incumbent-settings-object-3}
        from inside an algorithm that was triggered neither by [calling
        scripts](#calling-scripts) nor by Web IDL
        [invoking](https://webidl.spec.whatwg.org/#invoke-a-callback-function){#incumbent:es-invoking-callback-functions-2
        x-internal="es-invoking-callback-functions"} a callback. For
        example, it would trigger if you tried to obtain the [incumbent
        settings
        object](#incumbent-settings-object){#incumbent:incumbent-settings-object-4}
        inside an algorithm that ran periodically as part of the [event
        loop](#event-loop){#incumbent:event-loop-2}, with no involvement
        of author code. In such cases the
        [incumbent](#concept-incumbent-everything){#incumbent:concept-incumbent-everything}
        concept cannot be used.

    2.  Return the topmost entry of the [backup incumbent settings
        object
        stack](#backup-incumbent-settings-object-stack){#incumbent:backup-incumbent-settings-object-stack-5}.

3.  Return `context`{.variable}\'s Realm component\'s [settings
    object](#concept-realm-settings-object){#incumbent:concept-realm-settings-object}.

Then, the [incumbent realm]{#concept-incumbent-realm .dfn export=""} is
the
[realm](#environment-settings-object's-realm){#incumbent:environment-settings-object's-realm}
of the [incumbent settings
object](#incumbent-settings-object){#incumbent:incumbent-settings-object-5}.

Similarly, the [incumbent global object]{#concept-incumbent-global .dfn
export=""} is the [global
object](#concept-settings-object-global){#incumbent:concept-settings-object-global}
of the [incumbent settings
object](#incumbent-settings-object){#incumbent:incumbent-settings-object-6}.

------------------------------------------------------------------------

The following series of examples is intended to make it clear how all of
the different mechanisms contribute to the definition of the
[incumbent](#incumbent) concept:

::: {#example-incumbent-1 .example}
Consider the following starter example:

``` html
<!DOCTYPE html>
<iframe></iframe>
<script>
  frames[0].postMessage("some data", "*");
</script>
```

There are two interesting [environment settings
objects](#environment-settings-object){#incumbent:environment-settings-object-3}
here: that of `window`, and that of `frames[0]`. Our concern is: what is
the [incumbent settings
object](#incumbent-settings-object){#incumbent:incumbent-settings-object-7}
at the time that the algorithm for
[`postMessage()`{#incumbent:dom-window-postmessage}](web-messaging.html#dom-window-postmessage)
executes?

It should be that of `window`, to capture the intuitive notion that the
author script responsible for causing the algorithm to happen is
executing in `window`, not `frames[0]`. This makes sense: the [window
post message
steps](web-messaging.html#window-post-message-steps){#incumbent:window-post-message-steps}
use the [incumbent settings
object](#incumbent-settings-object){#incumbent:incumbent-settings-object-8}
to determine the
[`source`{#incumbent:dom-messageevent-source}](comms.html#dom-messageevent-source)
property of the resulting
[`MessageEvent`{#incumbent:messageevent}](comms.html#messageevent), and
in this case `window` is definitely the source of the message.

Let us now explain how the steps given above give us our
intuitively-desired result of `window`\'s [relevant settings
object](#relevant-settings-object){#incumbent:relevant-settings-object}.

When the [window post message
steps](web-messaging.html#window-post-message-steps){#incumbent:window-post-message-steps-2}
look up the [incumbent settings
object](#incumbent-settings-object){#incumbent:incumbent-settings-object-9},
the [topmost script-having execution
context](#topmost-script-having-execution-context){#incumbent:topmost-script-having-execution-context-5}
will be that corresponding to the
[`script`{#incumbent:the-script-element}](scripting.html#the-script-element)
element: it was pushed onto the [JavaScript execution context
stack](https://tc39.es/ecma262/#execution-context-stack){#incumbent:javascript-execution-context-stack-3
x-internal="javascript-execution-context-stack"} as part of
[ScriptEvaluation](https://tc39.es/ecma262/#sec-runtime-semantics-scriptevaluation){#incumbent:js-scriptevaluation
x-internal="js-scriptevaluation"} during the [run a classic
script](#run-a-classic-script){#incumbent:run-a-classic-script}
algorithm. Since there are no Web IDL callback invocations involved, the
context\'s [skip-when-determining-incumbent
counter](#skip-when-determining-incumbent-counter){#incumbent:skip-when-determining-incumbent-counter-4}
is zero, so it is used to determine the [incumbent settings
object](#incumbent-settings-object){#incumbent:incumbent-settings-object-10};
the result is the [environment settings
object](#environment-settings-object){#incumbent:environment-settings-object-4}
of `window`.

(Note how the [environment settings
object](#environment-settings-object){#incumbent:environment-settings-object-5}
of `frames[0]` is the [relevant settings
object](#relevant-settings-object){#incumbent:relevant-settings-object-2}
of [this](https://webidl.spec.whatwg.org/#this){#incumbent:this
x-internal="this"} at the time the
[`postMessage()`{#incumbent:dom-window-postmessage-2}](web-messaging.html#dom-window-postmessage)
method is called, and thus is involved in determining the *target* of
the message. Whereas the incumbent is used to determine the *source*.)
:::

::: {#example-incumbent-2 .example}
Consider the following more complicated example:

``` html
<!DOCTYPE html>
<iframe></iframe>
<script>
  const bound = frames[0].postMessage.bind(frames[0], "some data", "*");
  window.setTimeout(bound);
</script>
```

This example is very similar to the previous one, but with an extra
indirection through `Function.prototype.bind` as well as
[`setTimeout()`{#incumbent:dom-settimeout}](timers-and-user-prompts.html#dom-settimeout).
But, the answer should be the same: invoking algorithms asynchronously
should not change the
[incumbent](#concept-incumbent-everything){#incumbent:concept-incumbent-everything-2}
concept.

This time, the result involves more complicated mechanisms:

When `bound` is
[converted](https://webidl.spec.whatwg.org/#es-type-mapping){#incumbent:concept-idl-convert
x-internal="concept-idl-convert"} to a Web IDL callback type, the
[incumbent settings
object](#incumbent-settings-object){#incumbent:incumbent-settings-object-11}
is that corresponding to `window` (in the same manner as in our starter
example above). Web IDL stores this as the resulting callback value\'s
[callback
context](https://webidl.spec.whatwg.org/#dfn-callback-context){#incumbent:callback-context
x-internal="callback-context"}.

When the [task](#concept-task){#incumbent:concept-task} posted by
[`setTimeout()`{#incumbent:dom-settimeout-2}](timers-and-user-prompts.html#dom-settimeout)
executes, the algorithm for that task uses Web IDL to
[invoke](https://webidl.spec.whatwg.org/#invoke-a-callback-function){#incumbent:es-invoking-callback-functions-3
x-internal="es-invoking-callback-functions"} the stored callback value.
Web IDL in turn calls the above [prepare to run a
callback](#prepare-to-run-a-callback){#incumbent:prepare-to-run-a-callback-4}
algorithm. This pushes the stored [callback
context](https://webidl.spec.whatwg.org/#dfn-callback-context){#incumbent:callback-context-2
x-internal="callback-context"} onto the [backup incumbent settings
object
stack](#backup-incumbent-settings-object-stack){#incumbent:backup-incumbent-settings-object-stack-6}.
At this time (inside the timer task) there is no author code on the
stack, so the [topmost script-having execution
context](#topmost-script-having-execution-context){#incumbent:topmost-script-having-execution-context-6}
is null, and nothing gets its [skip-when-determining-incumbent
counter](#skip-when-determining-incumbent-counter){#incumbent:skip-when-determining-incumbent-counter-5}
incremented.

Invoking the callback then calls `bound`, which in turn calls the
[`postMessage()`{#incumbent:dom-window-postmessage-3}](web-messaging.html#dom-window-postmessage)
method of `frames[0]`. When the
[`postMessage()`{#incumbent:dom-window-postmessage-4}](web-messaging.html#dom-window-postmessage)
algorithm looks up the [incumbent settings
object](#incumbent-settings-object){#incumbent:incumbent-settings-object-12},
there is still no author code on the stack, since the bound function
just directly calls the built-in method. So the [topmost script-having
execution
context](#topmost-script-having-execution-context){#incumbent:topmost-script-having-execution-context-7}
will be null: the [JavaScript execution
context](https://tc39.es/ecma262/#sec-execution-contexts){#incumbent:javascript-execution-context-2
x-internal="javascript-execution-context"} stack only contains an
execution context corresponding to
[`postMessage()`{#incumbent:dom-window-postmessage-5}](web-messaging.html#dom-window-postmessage),
with no
[ScriptEvaluation](https://tc39.es/ecma262/#sec-runtime-semantics-scriptevaluation){#incumbent:js-scriptevaluation-2
x-internal="js-scriptevaluation"} context or similar below it.

This is where we fall back to the [backup incumbent settings object
stack](#backup-incumbent-settings-object-stack){#incumbent:backup-incumbent-settings-object-stack-7}.
As noted above, it will contain as its topmost entry the [relevant
settings
object](#relevant-settings-object){#incumbent:relevant-settings-object-3}
of `window`. So that is what is used as the [incumbent settings
object](#incumbent-settings-object){#incumbent:incumbent-settings-object-13}
while executing the
[`postMessage()`{#incumbent:dom-window-postmessage-6}](web-messaging.html#dom-window-postmessage)
algorithm.
:::

::: {#example-incumbent-3 .example}
Consider this final, even more convoluted example:

``` html
<!-- a.html -->
<!DOCTYPE html>
<button>click me</button>
<iframe></iframe>
<script>
const bound = frames[0].location.assign.bind(frames[0].location, "https://example.com/");
document.querySelector("button").addEventListener("click", bound);
</script>
```

``` html
<!-- b.html -->
<!DOCTYPE html>
<iframe src="a.html"></iframe>
<script>
  const iframe = document.querySelector("iframe");
  iframe.onload = function onLoad() {
    iframe.contentWindow.document.querySelector("button").click();
  };
</script>
```

Again there are two interesting [environment settings
objects](#environment-settings-object){#incumbent:environment-settings-object-6}
in play: that of `a.html`, and that of `b.html`. When the
[`location.assign()`{#incumbent:dom-location-assign}](nav-history-apis.html#dom-location-assign)
method triggers the [`Location`-object
navigate](nav-history-apis.html#location-object-navigate){#incumbent:location-object-navigate}
algorithm, what will be the [incumbent settings
object](#incumbent-settings-object){#incumbent:incumbent-settings-object-14}?
As before, it should intuitively be that of `a.html`: the
[`click`{#incumbent:event-click}](https://w3c.github.io/uievents/#event-type-click){x-internal="event-click"}
listener was originally scheduled by `a.html`, so even if something
involving `b.html` causes the listener to fire, the
[incumbent](#concept-incumbent-everything){#incumbent:concept-incumbent-everything-3}
responsible is that of `a.html`.

The callback setup is similar to the previous example: when `bound` is
[converted](https://webidl.spec.whatwg.org/#es-type-mapping){#incumbent:concept-idl-convert-2
x-internal="concept-idl-convert"} to a Web IDL callback type, the
[incumbent settings
object](#incumbent-settings-object){#incumbent:incumbent-settings-object-15}
is that corresponding to `a.html`, which is stored as the callback\'s
[callback
context](https://webidl.spec.whatwg.org/#dfn-callback-context){#incumbent:callback-context-3
x-internal="callback-context"}.

When the [`click()`{#incumbent:dom-click}](interaction.html#dom-click)
method is called inside `b.html`, it
[dispatches](https://dom.spec.whatwg.org/#concept-event-dispatch){#incumbent:concept-event-dispatch
x-internal="concept-event-dispatch"} a
[`click`{#incumbent:event-click-2}](https://w3c.github.io/uievents/#event-type-click){x-internal="event-click"}
event on the button that is inside `a.html`. This time, when the
[prepare to run a
callback](#prepare-to-run-a-callback){#incumbent:prepare-to-run-a-callback-5}
algorithm executes as part of event dispatch, there *is* author code on
the stack; the [topmost script-having execution
context](#topmost-script-having-execution-context){#incumbent:topmost-script-having-execution-context-8}
is that of the `onLoad` function, whose [skip-when-determining-incumbent
counter](#skip-when-determining-incumbent-counter){#incumbent:skip-when-determining-incumbent-counter-6}
gets incremented. Additionally, `a.html`\'s [environment settings
object](#environment-settings-object){#incumbent:environment-settings-object-7}
(stored as the
[`EventHandler`{#incumbent:eventhandler}](#eventhandler)\'s [callback
context](https://webidl.spec.whatwg.org/#dfn-callback-context){#incumbent:callback-context-4
x-internal="callback-context"}) is pushed onto the [backup incumbent
settings object
stack](#backup-incumbent-settings-object-stack){#incumbent:backup-incumbent-settings-object-stack-8}.

Now, when the [`Location`-object
navigate](nav-history-apis.html#location-object-navigate){#incumbent:location-object-navigate-2}
algorithm looks up the [incumbent settings
object](#incumbent-settings-object){#incumbent:incumbent-settings-object-16},
the [topmost script-having execution
context](#topmost-script-having-execution-context){#incumbent:topmost-script-having-execution-context-9}
is still that of the `onLoad` function (due to the fact we are using a
bound function as the callback). Its [skip-when-determining-incumbent
counter](#skip-when-determining-incumbent-counter){#incumbent:skip-when-determining-incumbent-counter-7}
value is one, however, so we fall back to the [backup incumbent settings
object
stack](#backup-incumbent-settings-object-stack){#incumbent:backup-incumbent-settings-object-stack-9}.
This gives us the [environment settings
object](#environment-settings-object){#incumbent:environment-settings-object-8}
of `a.html`, as expected.

Note that this means that even though it is the
[`iframe`{#incumbent:the-iframe-element}](iframe-embed-object.html#the-iframe-element)
inside `a.html` that navigates, it is `a.html` itself that is used as
the source [`Document`{#incumbent:document}](dom.html#document), which
determines among other things the [request
client](https://fetch.spec.whatwg.org/#concept-request-client){#incumbent:concept-request-client
x-internal="concept-request-client"}. This is [perhaps the only
justifiable use of the incumbent concept on the web
platform](https://www.w3.org/Bugs/Public/show_bug.cgi?id=26603#c36); in
all other cases the consequences of using it are simply confusing and we
hope to one day switch them to use
[current](#concept-current-everything){#incumbent:concept-current-everything}
or
[relevant](#concept-relevant-everything){#incumbent:concept-relevant-everything}
as appropriate.
:::

###### [8.1.3.3.3]{.secno} Current[](#current){.self-link}

The JavaScript specification defines the [current
realm](https://tc39.es/ecma262/#current-realm){#current:current-realm
x-internal="current-realm"}, also known as the \"current Realm Record\".
[\[JAVASCRIPT\]](references.html#refsJAVASCRIPT)

Then, the [current settings object]{#current-settings-object .dfn
export=""} is the [environment settings
object](#concept-realm-settings-object){#current:concept-realm-settings-object}
of the [current
realm](https://tc39.es/ecma262/#current-realm){#current:current-realm-2
x-internal="current-realm"}.

Similarly, the [current global object]{#current-global-object .dfn
export=""} is the [global
object](#concept-realm-global){#current:concept-realm-global} of the
[current
realm](https://tc39.es/ecma262/#current-realm){#current:current-realm-3
x-internal="current-realm"}.

###### [8.1.3.3.4]{.secno} Relevant[](#relevant){.self-link}

The [relevant realm]{#concept-relevant-realm .dfn export=""} for a
[platform
object](https://webidl.spec.whatwg.org/#dfn-platform-object){#relevant:platform-object
x-internal="platform-object"} is the value of [its \[\[Realm\]\]
field](https://webidl.spec.whatwg.org/#es-platform-objects){#relevant:realm-field-of-a-platform-object
x-internal="realm-field-of-a-platform-object"}.

Then, the [relevant settings object]{#relevant-settings-object .dfn
export=""} for a [platform
object](https://webidl.spec.whatwg.org/#dfn-platform-object){#relevant:platform-object-2
x-internal="platform-object"} `o`{.variable} is the [environment
settings
object](#concept-realm-settings-object){#relevant:concept-realm-settings-object}
of the [relevant
realm](#concept-relevant-realm){#relevant:concept-relevant-realm} for
`o`{.variable}.

Similarly, the [relevant global object]{#concept-relevant-global .dfn
export=""} for a [platform
object](https://webidl.spec.whatwg.org/#dfn-platform-object){#relevant:platform-object-3
x-internal="platform-object"} `o`{.variable} is the [global
object](#concept-realm-global){#relevant:concept-realm-global} of the
[relevant
realm](#concept-relevant-realm){#relevant:concept-relevant-realm-2} for
`o`{.variable}.

##### [8.1.3.4]{.secno} Enabling and disabling scripting[](#enabling-and-disabling-scripting){.self-link}

[Scripting is enabled]{#concept-environment-script .dfn} for an
[environment settings
object](#environment-settings-object){#enabling-and-disabling-scripting:environment-settings-object}
`settings`{.variable} when all of the following conditions are true:

- The user agent supports scripting.
- [![(This is a tracking
  vector.)](https://resources.whatwg.org/tracking-vector.svg){width="46"
  crossorigin=""
  height="64"}](https://infra.spec.whatwg.org/#tracking-vector "There is a tracking vector here."){#enabling-and-disabling-scripting:tracking-vector
  .tracking-vector x-internal="tracking-vector"} The user has not
  disabled scripting for `settings`{.variable} at this time. (User
  agents may provide users with the option to disable scripting
  globally, or in a finer-grained manner, e.g., on a per-origin basis,
  down to the level of individual [environment settings
  objects](#environment-settings-object){#enabling-and-disabling-scripting:environment-settings-object-2}.)
- [Either `settings`{.variable}\'s [global
  object](#concept-settings-object-global){#enabling-and-disabling-scripting:concept-settings-object-global}
  is not a
  [`Window`{#enabling-and-disabling-scripting:window}](nav-history-apis.html#window)
  object, or `settings`{.variable}\'s [global
  object](#concept-settings-object-global){#enabling-and-disabling-scripting:concept-settings-object-global-2}\'s
  [associated
  `Document`](nav-history-apis.html#concept-document-window){#enabling-and-disabling-scripting:concept-document-window}\'s
  [active sandboxing flag
  set](browsers.html#active-sandboxing-flag-set){#enabling-and-disabling-scripting:active-sandboxing-flag-set}
  does not have its [sandboxed scripts browsing context
  flag](browsers.html#sandboxed-scripts-browsing-context-flag){#enabling-and-disabling-scripting:sandboxed-scripts-browsing-context-flag}
  set.]{#sandboxScriptBlocked}

[Scripting is disabled]{#concept-environment-noscript .dfn} for an
[environment settings
object](#environment-settings-object){#enabling-and-disabling-scripting:environment-settings-object-3}
when scripting is not
[enabled](#concept-environment-script){#enabling-and-disabling-scripting:concept-environment-script}
for it, i.e., when any of the above conditions are false.

------------------------------------------------------------------------

[Scripting is enabled]{#concept-n-script .dfn} for a node
`node`{.variable} if `node`{.variable}\'s [node
document](https://dom.spec.whatwg.org/#concept-node-document){#enabling-and-disabling-scripting:node-document
x-internal="node-document"}\'s [browsing
context](document-sequences.html#concept-document-bc){#enabling-and-disabling-scripting:concept-document-bc}
is non-null, and [scripting is
enabled](#concept-environment-script){#enabling-and-disabling-scripting:concept-environment-script-2}
for `node`{.variable}\'s [relevant settings
object](#relevant-settings-object){#enabling-and-disabling-scripting:relevant-settings-object}.

[Scripting is disabled]{#concept-n-noscript .dfn} for a node when
scripting is not
[enabled](#concept-n-script){#enabling-and-disabling-scripting:concept-n-script},
i.e., when its [node
document](https://dom.spec.whatwg.org/#concept-node-document){#enabling-and-disabling-scripting:node-document-2
x-internal="node-document"}\'s [browsing
context](document-sequences.html#concept-document-bc){#enabling-and-disabling-scripting:concept-document-bc-2}
is null or when [scripting is
disabled](#concept-environment-noscript){#enabling-and-disabling-scripting:concept-environment-noscript}
for its [relevant settings
object](#relevant-settings-object){#enabling-and-disabling-scripting:relevant-settings-object-2}.

##### [8.1.3.5]{.secno} Secure contexts[](#secure-contexts){.self-link}

An [environment](#environment){#secure-contexts:environment}
`environment`{.variable} is a [secure context]{#secure-context .dfn
lt="secure context|Is
  an environment settings object contextually secure?" export=""} if the
following algorithm returns true:

1.  If `environment`{.variable} is an [environment settings
    object](#environment-settings-object){#secure-contexts:environment-settings-object},
    then:

    1.  Let `global`{.variable} be `environment`{.variable}\'s [global
        object](#concept-settings-object-global){#secure-contexts:concept-settings-object-global}.

    2.  If `global`{.variable} is a
        [`WorkerGlobalScope`{#secure-contexts:workerglobalscope}](workers.html#workerglobalscope),
        then:

        1.  If `global`{.variable}\'s [owner
            set](workers.html#concept-WorkerGlobalScope-owner-set){#secure-contexts:concept-WorkerGlobalScope-owner-set}\[0\]\'s
            [relevant settings
            object](#relevant-settings-object){#secure-contexts:relevant-settings-object}
            is a [secure
            context](#secure-context){#secure-contexts:secure-context},
            then return true.

            We only need to check the 0th item since they will
            necessarily all be consistent.

        2.  Return false.

    3.  If `global`{.variable} is a
        [`WorkletGlobalScope`{#secure-contexts:workletglobalscope}](worklets.html#workletglobalscope),
        then return true.

        Worklets can only be created in secure contexts.

2.  If the result of [Is url potentially
    trustworthy?](https://w3c.github.io/webappsec-secure-contexts/#potentially-trustworthy-url){#secure-contexts:is-url-potentially-trustworthy
    x-internal="is-url-potentially-trustworthy"} given
    `environment`{.variable}\'s [top-level creation
    URL](#concept-environment-top-level-creation-url){#secure-contexts:concept-environment-top-level-creation-url}
    is \"`Potentially Trustworthy`\", then return true.

3.  Return false.

An [environment](#environment){#secure-contexts:environment-2} is a
[non-secure context]{#non-secure-context .dfn export=""} if it is not a
[secure context](#secure-context){#secure-contexts:secure-context-2}.

#### [8.1.4]{.secno} []{#processing-model-7}Script processing model[](#scripting-processing-model){.self-link} {#scripting-processing-model}

##### [8.1.4.1]{.secno} []{#definitions-2}Scripts[](#script-structs){.self-link} {#script-structs}

A [script]{#concept-script .dfn export=""} is one of two possible
[structs](https://infra.spec.whatwg.org/#struct){#script-structs:struct
x-internal="struct"} (namely, a [classic
script](#classic-script){#script-structs:classic-script} or a [module
script](#module-script){#script-structs:module-script}). All scripts
have:

A [settings object]{#settings-object .dfn dfn-for="script" export=""}
:   An [environment settings
    object](#environment-settings-object){#script-structs:environment-settings-object},
    containing various settings that are shared with other
    [scripts](#concept-script){#script-structs:concept-script} in the
    same context.

A [record]{#concept-script-record .dfn dfn-for="script" export=""}

:   One of the following:

    - a [script
      record](https://tc39.es/ecma262/#sec-script-records){#script-structs:script-record
      x-internal="script-record"}, for [classic
      scripts](#classic-script){#script-structs:classic-script-2};

    - a [Source Text Module
      Record](https://tc39.es/ecma262/#sec-source-text-module-records){#script-structs:source-text-module-record
      x-internal="source-text-module-record"}, for [JavaScript module
      scripts](#javascript-module-script){#script-structs:javascript-module-script};

    - a [Synthetic Module
      Record](https://tc39.es/ecma262/#sec-synthetic-module-records){#script-structs:synthetic-module-record
      x-internal="synthetic-module-record"}, for [CSS module
      scripts](#css-module-script){#script-structs:css-module-script}
      and [JSON module
      scripts](#json-module-script){#script-structs:json-module-script};

    - a [WebAssembly Module
      Record](https://webassembly.github.io/esm-integration/js-api/index.html#webassembly-module-record){#script-structs:webassembly-module-record
      x-internal="webassembly-module-record"}, for [WebAssembly module
      scripts](#webassembly-module-script){#script-structs:webassembly-module-script};
      or

    - null, representing a parsing failure.

A [parse error]{#concept-script-parse-error .dfn dfn-for="script"}

:   A JavaScript value, which has meaning only if the
    [record](#concept-script-record){#script-structs:concept-script-record}
    is null, indicating that the corresponding script source text could
    not be parsed.

    This value is used for internal tracking of immediate parse errors
    when [creating scripts](#creating-scripts), and is not to be used
    directly. Instead, consult the [error to
    rethrow](#concept-script-error-to-rethrow){#script-structs:concept-script-error-to-rethrow}
    when determining \"what went wrong\" for this script.

An [error to rethrow]{#concept-script-error-to-rethrow .dfn dfn-for="script" export=""}

:   A JavaScript value representing an error that will prevent
    evaluation from succeeding. It will be re-thrown by any attempts to
    [run](#calling-scripts) the script.

    This could be the script\'s [parse
    error](#concept-script-parse-error){#script-structs:concept-script-parse-error},
    but in the case of a [module
    script](#module-script){#script-structs:module-script-2} it could
    instead be the [parse
    error](#concept-script-parse-error){#script-structs:concept-script-parse-error-2}
    from one of its dependencies, or an error from [resolve a module
    specifier](#resolve-a-module-specifier){#script-structs:resolve-a-module-specifier}.

    Since this exception value is provided by the JavaScript
    specification, we know that it is never null, so we use null to
    signal that no error has occurred.

[Fetch options]{#concept-script-script-fetch-options .dfn dfn-for="script" export=""}
:   Null or a [script fetch
    options](#script-fetch-options){#script-structs:script-fetch-options},
    containing various options related to fetching this script or
    [module scripts](#module-script){#script-structs:module-script-3}
    that it imports.

A [base URL]{#concept-script-base-url .dfn dfn-for="script" export=""}

:   Null or a base
    [URL](https://url.spec.whatwg.org/#concept-url){#script-structs:url
    x-internal="url"} used for [resolving module
    specifiers](#resolve-a-module-specifier){#script-structs:resolve-a-module-specifier-2}.
    When non-null, this will either be the URL from which the script was
    obtained, for external scripts, or the [document base
    URL](urls-and-fetching.html#document-base-url){#script-structs:document-base-url}
    of the containing document, for inline scripts.

A [classic script]{#classic-script .dfn export=""} is a type of
[script](#concept-script){#script-structs:concept-script-2} that has the
following additional
[item](https://infra.spec.whatwg.org/#struct-item){#script-structs:struct-item
x-internal="struct-item"}:

A [muted errors]{#muted-errors .dfn} boolean

:   A boolean which, if true, means that error information will not be
    provided for errors in this script. This is used to mute errors for
    cross-origin scripts, since that can leak private information.

A [module script]{#module-script .dfn export=""} is another type of
[script](#concept-script){#script-structs:concept-script-3}. It has no
additional
[items](https://infra.spec.whatwg.org/#struct-item){#script-structs:struct-item-2
x-internal="struct-item"}.

[Module scripts](#module-script){#script-structs:module-script-4} can be
classified into four types:

- A [module script](#module-script){#script-structs:module-script-5} is
  a [JavaScript module script]{#javascript-module-script .dfn export=""}
  if its
  [record](#concept-script-record){#script-structs:concept-script-record-2}
  is a [Source Text Module
  Record](https://tc39.es/ecma262/#sec-source-text-module-records){#script-structs:source-text-module-record-2
  x-internal="source-text-module-record"}.

- A [module script](#module-script){#script-structs:module-script-6} is
  a [CSS module script]{#css-module-script .dfn export=""} if its
  [record](#concept-script-record){#script-structs:concept-script-record-3}
  is a [Synthetic Module
  Record](https://tc39.es/ecma262/#sec-synthetic-module-records){#script-structs:synthetic-module-record-2
  x-internal="synthetic-module-record"}, and it was created via the
  [create a CSS module
  script](#creating-a-css-module-script){#script-structs:creating-a-css-module-script}
  algorithm. CSS module scripts represent a parsed CSS stylesheet.

- A [module script](#module-script){#script-structs:module-script-7} is
  a [JSON module script]{#json-module-script .dfn export=""} if its
  [record](#concept-script-record){#script-structs:concept-script-record-4}
  is a [Synthetic Module
  Record](https://tc39.es/ecma262/#sec-synthetic-module-records){#script-structs:synthetic-module-record-3
  x-internal="synthetic-module-record"}, and it was created via the
  [create a JSON module
  script](#creating-a-json-module-script){#script-structs:creating-a-json-module-script}
  algorithm. JSON module scripts represent a parsed JSON document.

- A [module script](#module-script){#script-structs:module-script-8} is
  a [WebAssembly module script]{#webassembly-module-script .dfn
  export=""} if its
  [record](#concept-script-record){#script-structs:concept-script-record-5}
  is a [WebAssembly Module
  Record](https://webassembly.github.io/esm-integration/js-api/index.html#webassembly-module-record){#script-structs:webassembly-module-record-2
  x-internal="webassembly-module-record"}.

As CSS stylesheets and JSON documents do not import dependent modules,
and do not throw exceptions on evaluation, the [fetch
options](#concept-script-script-fetch-options){#script-structs:concept-script-script-fetch-options}
and [base
URL](#concept-script-base-url){#script-structs:concept-script-base-url}
of [CSS module
scripts](#css-module-script){#script-structs:css-module-script-2} and
[JSON module
scripts](#json-module-script){#script-structs:json-module-script-2} and
are always null.

The [active script]{#active-script .dfn} is determined by the following
algorithm:

1.  Let `record`{.variable} be
    [GetActiveScriptOrModule](https://tc39.es/ecma262/#sec-getactivescriptormodule){#script-structs:getactivescriptormodule
    x-internal="getactivescriptormodule"}().

2.  If `record`{.variable} is null, return null.

3.  Return `record`{.variable}.\[\[HostDefined\]\].

The [active script](#active-script){#script-structs:active-script}
concept is so far only used by the
[`import()`{#script-structs:import()}](https://tc39.es/ecma262/#sec-import-calls){x-internal="import()"}
feature, to determine the [base
URL](#concept-script-base-url){#script-structs:concept-script-base-url-2}
to use for resolving relative module specifiers.

##### [8.1.4.2]{.secno} Fetching scripts[](#fetching-scripts){.self-link}

This section introduces a number of algorithms for fetching scripts,
taking various necessary inputs and resulting in
[classic](#classic-script){#fetching-scripts:classic-script} or [module
scripts](#module-script){#fetching-scripts:module-script}.

------------------------------------------------------------------------

[Script fetch options]{#script-fetch-options .dfn} is a
[struct](https://infra.spec.whatwg.org/#struct){#fetching-scripts:struct
x-internal="struct"} with the following
[items](https://infra.spec.whatwg.org/#struct-item){#fetching-scripts:struct-item
x-internal="struct-item"}:

[cryptographic nonce]{#concept-script-fetch-options-nonce .dfn}
:   The [cryptographic nonce
    metadata](https://fetch.spec.whatwg.org/#concept-request-nonce-metadata){#fetching-scripts:concept-request-nonce-metadata
    x-internal="concept-request-nonce-metadata"} used for the initial
    fetch and for fetching any imported modules

[integrity metadata]{#concept-script-fetch-options-integrity .dfn}
:   The [integrity
    metadata](https://fetch.spec.whatwg.org/#concept-request-integrity-metadata){#fetching-scripts:concept-request-integrity-metadata
    x-internal="concept-request-integrity-metadata"} used for the
    initial fetch

[parser metadata]{#concept-script-fetch-options-parser .dfn}
:   The [parser
    metadata](https://fetch.spec.whatwg.org/#concept-request-parser-metadata){#fetching-scripts:concept-request-parser-metadata
    x-internal="concept-request-parser-metadata"} used for the initial
    fetch and for fetching any imported modules

[credentials mode]{#concept-script-fetch-options-credentials .dfn}
:   The [credentials
    mode](https://fetch.spec.whatwg.org/#concept-request-credentials-mode){#fetching-scripts:concept-request-credentials-mode
    x-internal="concept-request-credentials-mode"} used for the initial
    fetch (for [module
    scripts](#module-script){#fetching-scripts:module-script-2}) and for
    fetching any imported modules (for both [module
    scripts](#module-script){#fetching-scripts:module-script-3} and
    [classic
    scripts](#classic-script){#fetching-scripts:classic-script-2})

[referrer policy]{#concept-script-fetch-options-referrer-policy .dfn}

:   The [referrer
    policy](https://fetch.spec.whatwg.org/#concept-request-referrer-policy){#fetching-scripts:concept-request-referrer-policy
    x-internal="concept-request-referrer-policy"} used for the initial
    fetch and for fetching any imported modules

    This policy can mutate after a [module
    script](#module-script){#fetching-scripts:module-script-4}\'s
    [response](https://fetch.spec.whatwg.org/#concept-response){#fetching-scripts:concept-response
    x-internal="concept-response"} is received, to be the [referrer
    policy](https://w3c.github.io/webappsec-referrer-policy/#referrer-policy){#fetching-scripts:referrer-policy
    x-internal="referrer-policy"}
    [parsed](https://w3c.github.io/webappsec-referrer-policy/#parse-referrer-policy-from-header){#fetching-scripts:parse-referrer-policy-header
    x-internal="parse-referrer-policy-header"} from the
    [response](https://fetch.spec.whatwg.org/#concept-response){#fetching-scripts:concept-response-2
    x-internal="concept-response"}, and used when fetching any module
    dependencies.

[render-blocking]{#concept-script-fetch-options-render-blocking .dfn}
:   The boolean value of
    [render-blocking](https://fetch.spec.whatwg.org/#request-render-blocking){#fetching-scripts:concept-request-render-blocking
    x-internal="concept-request-render-blocking"} used for the initial
    fetch and for fetching any imported modules. Unless otherwise
    stated, its value is false.

[fetch priority]{#concept-script-fetch-options-fetch-priority .dfn}

:   The
    [priority](https://fetch.spec.whatwg.org/#request-priority){#fetching-scripts:concept-request-priority
    x-internal="concept-request-priority"} used for the initial fetch

Recall that via the
[`import()`{#fetching-scripts:import()}](https://tc39.es/ecma262/#sec-import-calls){x-internal="import()"}
feature, [classic
scripts](#classic-script){#fetching-scripts:classic-script-3} can import
[module scripts](#module-script){#fetching-scripts:module-script-5}.

The [default script fetch options]{#default-script-fetch-options .dfn}
are a [script fetch
options](#script-fetch-options){#fetching-scripts:script-fetch-options}
whose [cryptographic
nonce](#concept-script-fetch-options-nonce){#fetching-scripts:concept-script-fetch-options-nonce}
is the empty string, [integrity
metadata](#concept-script-fetch-options-integrity){#fetching-scripts:concept-script-fetch-options-integrity}
is the empty string, [parser
metadata](#concept-script-fetch-options-parser){#fetching-scripts:concept-script-fetch-options-parser}
is \"`not-parser-inserted`\", [credentials
mode](#concept-script-fetch-options-credentials){#fetching-scripts:concept-script-fetch-options-credentials}
is \"`same-origin`\", [referrer
policy](#concept-script-fetch-options-referrer-policy){#fetching-scripts:concept-script-fetch-options-referrer-policy}
is the empty string, and [fetch
priority](#concept-script-fetch-options-fetch-priority){#fetching-scripts:concept-script-fetch-options-fetch-priority}
is \"`auto`\".

Given a
[request](https://fetch.spec.whatwg.org/#concept-request){#fetching-scripts:concept-request
x-internal="concept-request"} `request`{.variable} and a [script fetch
options](#script-fetch-options){#fetching-scripts:script-fetch-options-2}
`options`{.variable}, we define:

[set up the classic script request]{#set-up-the-classic-script-request .dfn}
:   Set `request`{.variable}\'s [cryptographic nonce
    metadata](https://fetch.spec.whatwg.org/#concept-request-nonce-metadata){#fetching-scripts:concept-request-nonce-metadata-2
    x-internal="concept-request-nonce-metadata"} to
    `options`{.variable}\'s [cryptographic
    nonce](#concept-script-fetch-options-nonce){#fetching-scripts:concept-script-fetch-options-nonce-2},
    its [integrity
    metadata](https://fetch.spec.whatwg.org/#concept-request-integrity-metadata){#fetching-scripts:concept-request-integrity-metadata-2
    x-internal="concept-request-integrity-metadata"} to
    `options`{.variable}\'s [integrity
    metadata](#concept-script-fetch-options-integrity){#fetching-scripts:concept-script-fetch-options-integrity-2},
    its [parser
    metadata](https://fetch.spec.whatwg.org/#concept-request-parser-metadata){#fetching-scripts:concept-request-parser-metadata-2
    x-internal="concept-request-parser-metadata"} to
    `options`{.variable}\'s [parser
    metadata](#concept-script-fetch-options-parser){#fetching-scripts:concept-script-fetch-options-parser-2},
    its [referrer
    policy](https://fetch.spec.whatwg.org/#concept-request-referrer-policy){#fetching-scripts:concept-request-referrer-policy-2
    x-internal="concept-request-referrer-policy"} to
    `options`{.variable}\'s [referrer
    policy](#concept-script-fetch-options-referrer-policy){#fetching-scripts:concept-script-fetch-options-referrer-policy-2},
    its
    [render-blocking](https://fetch.spec.whatwg.org/#request-render-blocking){#fetching-scripts:concept-request-render-blocking-2
    x-internal="concept-request-render-blocking"} to
    `options`{.variable}\'s
    [render-blocking](#concept-script-fetch-options-render-blocking){#fetching-scripts:concept-script-fetch-options-render-blocking},
    and its
    [priority](https://fetch.spec.whatwg.org/#request-priority){#fetching-scripts:concept-request-priority-2
    x-internal="concept-request-priority"} to `options`{.variable}\'s
    [fetch
    priority](#concept-script-fetch-options-fetch-priority){#fetching-scripts:concept-script-fetch-options-fetch-priority-2}.

[set up the module script request]{#set-up-the-module-script-request .dfn}

:   Set `request`{.variable}\'s [cryptographic nonce
    metadata](https://fetch.spec.whatwg.org/#concept-request-nonce-metadata){#fetching-scripts:concept-request-nonce-metadata-3
    x-internal="concept-request-nonce-metadata"} to
    `options`{.variable}\'s [cryptographic
    nonce](#concept-script-fetch-options-nonce){#fetching-scripts:concept-script-fetch-options-nonce-3},
    its [integrity
    metadata](https://fetch.spec.whatwg.org/#concept-request-integrity-metadata){#fetching-scripts:concept-request-integrity-metadata-3
    x-internal="concept-request-integrity-metadata"} to
    `options`{.variable}\'s [integrity
    metadata](#concept-script-fetch-options-integrity){#fetching-scripts:concept-script-fetch-options-integrity-3},
    its [parser
    metadata](https://fetch.spec.whatwg.org/#concept-request-parser-metadata){#fetching-scripts:concept-request-parser-metadata-3
    x-internal="concept-request-parser-metadata"} to
    `options`{.variable}\'s [parser
    metadata](#concept-script-fetch-options-parser){#fetching-scripts:concept-script-fetch-options-parser-3},
    its [credentials
    mode](https://fetch.spec.whatwg.org/#concept-request-credentials-mode){#fetching-scripts:concept-request-credentials-mode-2
    x-internal="concept-request-credentials-mode"} to
    `options`{.variable}\'s [credentials
    mode](#concept-script-fetch-options-credentials){#fetching-scripts:concept-script-fetch-options-credentials-2},
    its [referrer
    policy](https://fetch.spec.whatwg.org/#concept-request-referrer-policy){#fetching-scripts:concept-request-referrer-policy-3
    x-internal="concept-request-referrer-policy"} to
    `options`{.variable}\'s [referrer
    policy](#concept-script-fetch-options-referrer-policy){#fetching-scripts:concept-script-fetch-options-referrer-policy-3},
    its
    [render-blocking](https://fetch.spec.whatwg.org/#request-render-blocking){#fetching-scripts:concept-request-render-blocking-3
    x-internal="concept-request-render-blocking"} to
    `options`{.variable}\'s
    [render-blocking](#concept-script-fetch-options-render-blocking){#fetching-scripts:concept-script-fetch-options-render-blocking-2},
    and its
    [priority](https://fetch.spec.whatwg.org/#request-priority){#fetching-scripts:concept-request-priority-3
    x-internal="concept-request-priority"} to `options`{.variable}\'s
    [fetch
    priority](#concept-script-fetch-options-fetch-priority){#fetching-scripts:concept-script-fetch-options-fetch-priority-3}.

To [get the descendant script fetch
options]{#get-the-descendant-script-fetch-options .dfn} given a [script
fetch
options](#script-fetch-options){#fetching-scripts:script-fetch-options-3}
`originalOptions`{.variable}, a
[URL](https://url.spec.whatwg.org/#concept-url){#fetching-scripts:url
x-internal="url"} `url`{.variable}, and an [environment settings
object](#environment-settings-object){#fetching-scripts:environment-settings-object}
`settingsObject`{.variable}:

1.  Let `newOptions`{.variable} be a copy of
    `originalOptions`{.variable}.

2.  Let `integrity`{.variable} be the result of [resolving a module
    integrity
    metadata](#resolving-a-module-integrity-metadata){#fetching-scripts:resolving-a-module-integrity-metadata}
    with `url`{.variable} and `settingsObject`{.variable}.

3.  Set `newOptions`{.variable}\'s [integrity
    metadata](#concept-script-fetch-options-integrity){#fetching-scripts:concept-script-fetch-options-integrity-4}
    to `integrity`{.variable}.

4.  Set `newOptions`{.variable}\'s [fetch
    priority](#concept-script-fetch-options-fetch-priority){#fetching-scripts:concept-script-fetch-options-fetch-priority-4}
    to \"`auto`\".

5.  Return `newOptions`{.variable}.

To [resolve a module integrity
metadata]{#resolving-a-module-integrity-metadata .dfn}, given a
[URL](https://url.spec.whatwg.org/#concept-url){#fetching-scripts:url-2
x-internal="url"} `url`{.variable} and an [environment settings
object](#environment-settings-object){#fetching-scripts:environment-settings-object-2}
`settingsObject`{.variable}:

1.  Let `map`{.variable} be `settingsObject`{.variable}\'s [global
    object](#concept-settings-object-global){#fetching-scripts:concept-settings-object-global}\'s
    [import
    map](#concept-global-import-map){#fetching-scripts:concept-global-import-map}.

2.  If `map`{.variable}\'s
    [integrity](#concept-import-map-integrity){#fetching-scripts:concept-import-map-integrity}\[`url`{.variable}\]
    does not
    [exist](https://infra.spec.whatwg.org/#map-exists){#fetching-scripts:map-exists
    x-internal="map-exists"}, then return the empty string.

3.  Return `map`{.variable}\'s
    [integrity](#concept-import-map-integrity){#fetching-scripts:concept-import-map-integrity-2}\[`url`{.variable}\].

------------------------------------------------------------------------

Several of the below algorithms can be customized with a [perform the
fetch hook]{#fetching-scripts-perform-fetch .dfn
dfn-for="fetching scripts" export=""} algorithm, which takes a
[request](https://fetch.spec.whatwg.org/#concept-request){#fetching-scripts:concept-request-2
x-internal="concept-request"}, a boolean
[`isTopLevel`{#fetching-scripts:fetching-scripts-is-top-level
.variable}](#fetching-scripts-is-top-level), and a
[`processCustomFetchResponse`{.variable}]{#fetching-scripts-processcustomfetchresponse
.dfn dfn-for="fetching
  scripts" export=""} algorithm. It runs
[`processCustomFetchResponse`{#fetching-scripts:fetching-scripts-processcustomfetchresponse
.variable}](#fetching-scripts-processcustomfetchresponse) with a
[response](https://fetch.spec.whatwg.org/#concept-response){#fetching-scripts:concept-response-3
x-internal="concept-response"} and either null (on failure) or a [byte
sequence](https://infra.spec.whatwg.org/#byte-sequence){#fetching-scripts:byte-sequence
x-internal="byte-sequence"} containing the response body.
[`isTopLevel`{.variable}]{#fetching-scripts-is-top-level .dfn
dfn-for="fetching scripts" export=""} will be true for all [classic
script](#classic-script){#fetching-scripts:classic-script-4} fetches,
and for the initial fetch when [fetching an external module script
graph](#fetch-a-module-script-tree){#fetching-scripts:fetch-a-module-script-tree}
or [fetching a module worker script
graph](#fetch-a-module-worker-script-tree){#fetching-scripts:fetch-a-module-worker-script-tree},
but false for the fetches resulting from `import` statements encountered
throughout the graph or from `import()` expressions.

By default, not supplying a [perform the fetch
hook](#fetching-scripts-perform-fetch){#fetching-scripts:fetching-scripts-perform-fetch}
will cause the below algorithms to simply
[fetch](https://fetch.spec.whatwg.org/#concept-fetch){#fetching-scripts:concept-fetch
x-internal="concept-fetch"} the given
[request](https://fetch.spec.whatwg.org/#concept-request){#fetching-scripts:concept-request-3
x-internal="concept-request"}, with algorithm-specific customizations to
the
[request](https://fetch.spec.whatwg.org/#concept-request){#fetching-scripts:concept-request-4
x-internal="concept-request"} and validations of the resulting
[response](https://fetch.spec.whatwg.org/#concept-response){#fetching-scripts:concept-response-4
x-internal="concept-response"}.

To layer your own customizations on top of these algorithm-specific
ones, supply a [perform the fetch
hook](#fetching-scripts-perform-fetch){#fetching-scripts:fetching-scripts-perform-fetch-2}
that modifies the given
[request](https://fetch.spec.whatwg.org/#concept-request){#fetching-scripts:concept-request-5
x-internal="concept-request"},
[fetches](https://fetch.spec.whatwg.org/#concept-fetch){#fetching-scripts:concept-fetch-2
x-internal="concept-fetch"} it, and then performs specific validations
of the resulting
[response](https://fetch.spec.whatwg.org/#concept-response){#fetching-scripts:concept-response-5
x-internal="concept-response"} (completing with a [network
error](https://fetch.spec.whatwg.org/#concept-network-error){#fetching-scripts:network-error
x-internal="network-error"} if the validations fail).

The hook can also be used to perform more subtle customizations, such as
keeping a cache of
[responses](https://fetch.spec.whatwg.org/#concept-response){#fetching-scripts:concept-response-6
x-internal="concept-response"} and avoiding performing a
[fetch](https://fetch.spec.whatwg.org/#concept-fetch){#fetching-scripts:concept-fetch-3
x-internal="concept-fetch"} at all.

Service Workers is an example of a specification that runs these
algorithms with its own options for the hook.
[\[SW\]](references.html#refsSW)

------------------------------------------------------------------------

Now for the algorithms themselves.

To [fetch a classic script]{#fetch-a-classic-script .dfn export=""}
given a
[URL](https://url.spec.whatwg.org/#concept-url){#fetching-scripts:url-3
x-internal="url"} `url`{.variable}, an [environment settings
object](#environment-settings-object){#fetching-scripts:environment-settings-object-3}
`settingsObject`{.variable}, a [script fetch
options](#script-fetch-options){#fetching-scripts:script-fetch-options-4}
`options`{.variable}, a [CORS settings attribute
state](urls-and-fetching.html#cors-settings-attribute){#fetching-scripts:cors-settings-attribute}
`corsSetting`{.variable}, an
[encoding](https://encoding.spec.whatwg.org/#encoding){#fetching-scripts:encoding
x-internal="encoding"} `encoding`{.variable}, and an algorithm
`onComplete`{.variable}, run these steps. `onComplete`{.variable} must
be an algorithm accepting null (on failure) or a [classic
script](#classic-script){#fetching-scripts:classic-script-5} (on
success).

1.  Let `request`{.variable} be the result of [creating a potential-CORS
    request](urls-and-fetching.html#create-a-potential-cors-request){#fetching-scripts:create-a-potential-cors-request}
    given `url`{.variable}, \"`script`\", and `corsSetting`{.variable}.

2.  Set `request`{.variable}\'s
    [client](https://fetch.spec.whatwg.org/#concept-request-client){#fetching-scripts:concept-request-client
    x-internal="concept-request-client"} to `settingsObject`{.variable}.

3.  Set `request`{.variable}\'s [initiator
    type](https://fetch.spec.whatwg.org/#request-initiator-type){#fetching-scripts:concept-request-initiator-type
    x-internal="concept-request-initiator-type"} to \"`script`\".

4.  [Set up the classic script
    request](#set-up-the-classic-script-request){#fetching-scripts:set-up-the-classic-script-request}
    given `request`{.variable} and `options`{.variable}.

5.  [Fetch](https://fetch.spec.whatwg.org/#concept-fetch){#fetching-scripts:concept-fetch-4
    x-internal="concept-fetch"} `request`{.variable} with the following
    *[processResponseConsumeBody](https://fetch.spec.whatwg.org/#process-response-end-of-body){x-internal="processresponseconsumebody"}*
    steps given
    [response](https://fetch.spec.whatwg.org/#concept-response){#fetching-scripts:concept-response-7
    x-internal="concept-response"} `response`{.variable} and null,
    failure, or a [byte
    sequence](https://infra.spec.whatwg.org/#byte-sequence){#fetching-scripts:byte-sequence-2
    x-internal="byte-sequence"} `bodyBytes`{.variable}:

    `response`{.variable} can be either
    [CORS-same-origin](urls-and-fetching.html#cors-same-origin){#fetching-scripts:cors-same-origin}
    or
    [CORS-cross-origin](urls-and-fetching.html#cors-cross-origin){#fetching-scripts:cors-cross-origin}.
    This only affects how error reporting happens.

    1.  Set `response`{.variable} to `response`{.variable}\'s [unsafe
        response](urls-and-fetching.html#unsafe-response){#fetching-scripts:unsafe-response}.

    2.  If any of the following are true:

        - `bodyBytes`{.variable} is null or failure; or

        - `response`{.variable}\'s
          [status](https://fetch.spec.whatwg.org/#concept-response-status){#fetching-scripts:concept-response-status
          x-internal="concept-response-status"} is not an [ok
          status](https://fetch.spec.whatwg.org/#ok-status){#fetching-scripts:ok-status
          x-internal="ok-status"},

        then run `onComplete`{.variable} given null, and abort these
        steps.

        For historical reasons, this algorithm does not include MIME
        type checking, unlike the other script-fetching algorithms in
        this section.

    3.  Let `potentialMIMETypeForEncoding`{.variable} be the result of
        [extracting a MIME
        type](https://fetch.spec.whatwg.org/#concept-header-extract-mime-type){#fetching-scripts:extract-a-mime-type
        x-internal="extract-a-mime-type"} given `response`{.variable}\'s
        [header
        list](https://fetch.spec.whatwg.org/#concept-response-header-list){#fetching-scripts:concept-response-header-list
        x-internal="concept-response-header-list"}.

    4.  Set `encoding`{.variable} to the result of [legacy extracting an
        encoding](https://fetch.spec.whatwg.org/#legacy-extract-an-encoding){#fetching-scripts:legacy-extract-an-encoding
        x-internal="legacy-extract-an-encoding"} given
        `potentialMIMETypeForEncoding`{.variable} and
        `encoding`{.variable}.

        This intentionally ignores the [MIME type
        essence](https://mimesniff.spec.whatwg.org/#mime-type-essence){#fetching-scripts:mime-type-essence
        x-internal="mime-type-essence"}.

    5.  Let `sourceText`{.variable} be the result of
        [decoding](https://encoding.spec.whatwg.org/#decode){#fetching-scripts:decode
        x-internal="decode"} `bodyBytes`{.variable} to Unicode, using
        `encoding`{.variable} as the fallback encoding.

        The
        [decode](https://encoding.spec.whatwg.org/#decode){#fetching-scripts:decode-2
        x-internal="decode"} algorithm overrides `encoding`{.variable}
        if the file contains a BOM.

    6.  Let `mutedErrors`{.variable} be true if `response`{.variable}
        was
        [CORS-cross-origin](urls-and-fetching.html#cors-cross-origin){#fetching-scripts:cors-cross-origin-2},
        and false otherwise.

    7.  Let `script`{.variable} be the result of [creating a classic
        script](#creating-a-classic-script){#fetching-scripts:creating-a-classic-script}
        given `sourceText`{.variable}, `settingsObject`{.variable},
        `response`{.variable}\'s
        [URL](https://fetch.spec.whatwg.org/#concept-response-url){#fetching-scripts:concept-response-url
        x-internal="concept-response-url"}, `options`{.variable},
        `mutedErrors`{.variable}, and `url`{.variable}.

    8.  Run `onComplete`{.variable} given `script`{.variable}.

To [fetch a classic worker script]{#fetch-a-classic-worker-script .dfn
export=""} given a
[URL](https://url.spec.whatwg.org/#concept-url){#fetching-scripts:url-4
x-internal="url"} `url`{.variable}, an [environment settings
object](#environment-settings-object){#fetching-scripts:environment-settings-object-4}
`fetchClient`{.variable}, a
[destination](https://fetch.spec.whatwg.org/#concept-request-destination){#fetching-scripts:concept-request-destination
x-internal="concept-request-destination"} `destination`{.variable}, an
[environment settings
object](#environment-settings-object){#fetching-scripts:environment-settings-object-5}
`settingsObject`{.variable}, an algorithm `onComplete`{.variable}, and
an optional [perform the fetch
hook](#fetching-scripts-perform-fetch){#fetching-scripts:fetching-scripts-perform-fetch-3}
`performFetch`{.variable}, run these steps. `onComplete`{.variable} must
be an algorithm accepting null (on failure) or a [classic
script](#classic-script){#fetching-scripts:classic-script-6} (on
success).

1.  Let `request`{.variable} be a new
    [request](https://fetch.spec.whatwg.org/#concept-request){#fetching-scripts:concept-request-6
    x-internal="concept-request"} whose
    [URL](https://fetch.spec.whatwg.org/#concept-request-url){#fetching-scripts:concept-request-url
    x-internal="concept-request-url"} is `url`{.variable},
    [client](https://fetch.spec.whatwg.org/#concept-request-client){#fetching-scripts:concept-request-client-2
    x-internal="concept-request-client"} is `fetchClient`{.variable},
    [destination](https://fetch.spec.whatwg.org/#concept-request-destination){#fetching-scripts:concept-request-destination-2
    x-internal="concept-request-destination"} is
    `destination`{.variable}, [initiator
    type](https://fetch.spec.whatwg.org/#request-initiator-type){#fetching-scripts:concept-request-initiator-type-2
    x-internal="concept-request-initiator-type"} is \"`other`\",
    [mode](https://fetch.spec.whatwg.org/#concept-request-mode){#fetching-scripts:concept-request-mode
    x-internal="concept-request-mode"} is \"`same-origin`\",
    [credentials
    mode](https://fetch.spec.whatwg.org/#concept-request-credentials-mode){#fetching-scripts:concept-request-credentials-mode-3
    x-internal="concept-request-credentials-mode"} is \"`same-origin`\",
    [parser
    metadata](https://fetch.spec.whatwg.org/#concept-request-parser-metadata){#fetching-scripts:concept-request-parser-metadata-4
    x-internal="concept-request-parser-metadata"} is
    \"`not parser-inserted`\", and whose [use-URL-credentials
    flag](https://fetch.spec.whatwg.org/#concept-request-use-url-credentials-flag){#fetching-scripts:use-url-credentials-flag
    x-internal="use-url-credentials-flag"} is set.

2.  If `performFetch`{.variable} was given, run
    `performFetch`{.variable} with `request`{.variable}, true, and with
    `processResponseConsumeBody`{.variable} as defined below.

    Otherwise,
    [fetch](https://fetch.spec.whatwg.org/#concept-fetch){#fetching-scripts:concept-fetch-5
    x-internal="concept-fetch"} `request`{.variable} with
    *[processResponseConsumeBody](https://fetch.spec.whatwg.org/#process-response-end-of-body){x-internal="processresponseconsumebody"}*
    set to `processResponseConsumeBody`{.variable} as defined below.

    In both cases, let `processResponseConsumeBody`{.variable} given
    [response](https://fetch.spec.whatwg.org/#concept-response){#fetching-scripts:concept-response-8
    x-internal="concept-response"} `response`{.variable} and null,
    failure, or a [byte
    sequence](https://infra.spec.whatwg.org/#byte-sequence){#fetching-scripts:byte-sequence-3
    x-internal="byte-sequence"} `bodyBytes`{.variable} be the following
    algorithm:

    1.  Set `response`{.variable} to `response`{.variable}\'s [unsafe
        response](urls-and-fetching.html#unsafe-response){#fetching-scripts:unsafe-response-2}.

    2.  If any of the following are true:

        - `bodyBytes`{.variable} is null or failure; or

        - `response`{.variable}\'s
          [status](https://fetch.spec.whatwg.org/#concept-response-status){#fetching-scripts:concept-response-status-2
          x-internal="concept-response-status"} is not an [ok
          status](https://fetch.spec.whatwg.org/#ok-status){#fetching-scripts:ok-status-2
          x-internal="ok-status"},

        then run `onComplete`{.variable} given null, and abort these
        steps.

    3.  If all of the following are true:

        - `response`{.variable}\'s
          [URL](https://fetch.spec.whatwg.org/#concept-response-url){#fetching-scripts:concept-response-url-2
          x-internal="concept-response-url"}\'s
          [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#fetching-scripts:concept-url-scheme
          x-internal="concept-url-scheme"} is an [HTTP(S)
          scheme](https://fetch.spec.whatwg.org/#http-scheme){#fetching-scripts:http(s)-scheme
          x-internal="http(s)-scheme"}; and

        - the result of [extracting a MIME
          type](https://fetch.spec.whatwg.org/#concept-header-extract-mime-type){#fetching-scripts:extract-a-mime-type-2
          x-internal="extract-a-mime-type"} from
          `response`{.variable}\'s [header
          list](https://fetch.spec.whatwg.org/#concept-response-header-list){#fetching-scripts:concept-response-header-list-2
          x-internal="concept-response-header-list"} is not a
          [JavaScript MIME
          type](https://mimesniff.spec.whatwg.org/#javascript-mime-type){#fetching-scripts:javascript-mime-type
          x-internal="javascript-mime-type"},

        then run `onComplete`{.variable} given null, and abort these
        steps.

        Other [fetch
        schemes](https://fetch.spec.whatwg.org/#fetch-scheme){#fetching-scripts:fetch-scheme
        x-internal="fetch-scheme"} are exempted from MIME type checking
        for historical web-compatibility reasons. We might be able to
        tighten this in the future; see [issue
        #3255](https://github.com/whatwg/html/issues/3255).

    4.  Let `sourceText`{.variable} be the result of [UTF-8
        decoding](https://encoding.spec.whatwg.org/#utf-8-decode){#fetching-scripts:utf-8-decode
        x-internal="utf-8-decode"} `bodyBytes`{.variable}.

    5.  Let `script`{.variable} be the result of [creating a classic
        script](#creating-a-classic-script){#fetching-scripts:creating-a-classic-script-2}
        using `sourceText`{.variable}, `settingsObject`{.variable},
        `response`{.variable}\'s
        [URL](https://fetch.spec.whatwg.org/#concept-response-url){#fetching-scripts:concept-response-url-3
        x-internal="concept-response-url"}, and the [default script
        fetch
        options](#default-script-fetch-options){#fetching-scripts:default-script-fetch-options}.

    6.  Run `onComplete`{.variable} given `script`{.variable}.

To [fetch a classic worker-imported
script]{#fetch-a-classic-worker-imported-script .dfn export=""} given a
[URL](https://url.spec.whatwg.org/#concept-url){#fetching-scripts:url-5
x-internal="url"} `url`{.variable}, an [environment settings
object](#environment-settings-object){#fetching-scripts:environment-settings-object-6}
`settingsObject`{.variable}, and an optional [perform the fetch
hook](#fetching-scripts-perform-fetch){#fetching-scripts:fetching-scripts-perform-fetch-4}
`performFetch`{.variable}, run these steps. The algorithm will return a
[classic script](#classic-script){#fetching-scripts:classic-script-7} on
success, or throw an exception on failure.

1.  Let `response`{.variable} be null.

2.  Let `bodyBytes`{.variable} be null.

3.  Let `request`{.variable} be a new
    [request](https://fetch.spec.whatwg.org/#concept-request){#fetching-scripts:concept-request-7
    x-internal="concept-request"} whose
    [URL](https://fetch.spec.whatwg.org/#concept-request-url){#fetching-scripts:concept-request-url-2
    x-internal="concept-request-url"} is `url`{.variable},
    [client](https://fetch.spec.whatwg.org/#concept-request-client){#fetching-scripts:concept-request-client-3
    x-internal="concept-request-client"} is `settingsObject`{.variable},
    [destination](https://fetch.spec.whatwg.org/#concept-request-destination){#fetching-scripts:concept-request-destination-3
    x-internal="concept-request-destination"} is \"`script`\",
    [initiator
    type](https://fetch.spec.whatwg.org/#request-initiator-type){#fetching-scripts:concept-request-initiator-type-3
    x-internal="concept-request-initiator-type"} is \"`other`\", [parser
    metadata](https://fetch.spec.whatwg.org/#concept-request-parser-metadata){#fetching-scripts:concept-request-parser-metadata-5
    x-internal="concept-request-parser-metadata"} is
    \"`not parser-inserted`\", and whose [use-URL-credentials
    flag](https://fetch.spec.whatwg.org/#concept-request-use-url-credentials-flag){#fetching-scripts:use-url-credentials-flag-2
    x-internal="use-url-credentials-flag"} is set.

4.  If `performFetch`{.variable} was given, run
    `performFetch`{.variable} with `request`{.variable},
    `isTopLevel`{.variable}, and with
    `processResponseConsumeBody`{.variable} as defined below.

    Otherwise,
    [fetch](https://fetch.spec.whatwg.org/#concept-fetch){#fetching-scripts:concept-fetch-6
    x-internal="concept-fetch"} `request`{.variable} with
    *[processResponseConsumeBody](https://fetch.spec.whatwg.org/#process-response-end-of-body){x-internal="processresponseconsumebody"}*
    set to `processResponseConsumeBody`{.variable} as defined below.

    In both cases, let `processResponseConsumeBody`{.variable} given
    [response](https://fetch.spec.whatwg.org/#concept-response){#fetching-scripts:concept-response-9
    x-internal="concept-response"} `res`{.variable} and null, failure,
    or a [byte
    sequence](https://infra.spec.whatwg.org/#byte-sequence){#fetching-scripts:byte-sequence-4
    x-internal="byte-sequence"} `bb`{.variable} be the following
    algorithm:

    1.  Set `bodyBytes`{.variable} to `bb`{.variable}.

    2.  Set `response`{.variable} to `res`{.variable}.

5.  [Pause](#pause){#fetching-scripts:pause} until `response`{.variable}
    is not null.

    Unlike other algorithms in this section, the fetching process is
    synchronous here.

6.  Set `response`{.variable} to `response`{.variable}\'s [unsafe
    response](urls-and-fetching.html#unsafe-response){#fetching-scripts:unsafe-response-3}.

7.  If any of the following are true:

    - `bodyBytes`{.variable} is null or failure;

    - `response`{.variable}\'s
      [status](https://fetch.spec.whatwg.org/#concept-response-status){#fetching-scripts:concept-response-status-3
      x-internal="concept-response-status"} is not an [ok
      status](https://fetch.spec.whatwg.org/#ok-status){#fetching-scripts:ok-status-3
      x-internal="ok-status"}; or

    - the result of [extracting a MIME
      type](https://fetch.spec.whatwg.org/#concept-header-extract-mime-type){#fetching-scripts:extract-a-mime-type-3
      x-internal="extract-a-mime-type"} from `response`{.variable}\'s
      [header
      list](https://fetch.spec.whatwg.org/#concept-response-header-list){#fetching-scripts:concept-response-header-list-3
      x-internal="concept-response-header-list"} is not a [JavaScript
      MIME
      type](https://mimesniff.spec.whatwg.org/#javascript-mime-type){#fetching-scripts:javascript-mime-type-2
      x-internal="javascript-mime-type"},

    then throw a
    [\"`NetworkError`\"](https://webidl.spec.whatwg.org/#networkerror){#fetching-scripts:networkerror
    x-internal="networkerror"}
    [`DOMException`{#fetching-scripts:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

8.  Let `sourceText`{.variable} be the result of [UTF-8
    decoding](https://encoding.spec.whatwg.org/#utf-8-decode){#fetching-scripts:utf-8-decode-2
    x-internal="utf-8-decode"} `bodyBytes`{.variable}.

9.  Let `mutedErrors`{.variable} be true if `response`{.variable} was
    [CORS-cross-origin](urls-and-fetching.html#cors-cross-origin){#fetching-scripts:cors-cross-origin-3},
    and false otherwise.

10. Let `script`{.variable} be the result of [creating a classic
    script](#creating-a-classic-script){#fetching-scripts:creating-a-classic-script-3}
    given `sourceText`{.variable}, `settingsObject`{.variable},
    `response`{.variable}\'s
    [URL](https://fetch.spec.whatwg.org/#concept-response-url){#fetching-scripts:concept-response-url-4
    x-internal="concept-response-url"}, the [default script fetch
    options](#default-script-fetch-options){#fetching-scripts:default-script-fetch-options-2},
    and `mutedErrors`{.variable}.

11. Return `script`{.variable}.

To [fetch an external module script graph]{#fetch-a-module-script-tree
.dfn export=""} given a
[URL](https://url.spec.whatwg.org/#concept-url){#fetching-scripts:url-6
x-internal="url"} `url`{.variable}, an [environment settings
object](#environment-settings-object){#fetching-scripts:environment-settings-object-7}
`settingsObject`{.variable}, a [script fetch
options](#script-fetch-options){#fetching-scripts:script-fetch-options-5}
`options`{.variable}, and an algorithm `onComplete`{.variable}, run
these steps. `onComplete`{.variable} must be an algorithm accepting null
(on failure) or a [module
script](#module-script){#fetching-scripts:module-script-6} (on success).

1.  [Fetch a single module
    script](#fetch-a-single-module-script){#fetching-scripts:fetch-a-single-module-script}
    given `url`{.variable}, `settingsObject`{.variable}, \"`script`\",
    `options`{.variable}, `settingsObject`{.variable}, \"`client`\",
    true, and with the following steps given `result`{.variable}:

    1.  If `result`{.variable} is null, run `onComplete`{.variable}
        given null, and abort these steps.

    2.  [Fetch the descendants of and
        link](#fetch-the-descendants-of-and-link-a-module-script){#fetching-scripts:fetch-the-descendants-of-and-link-a-module-script}
        `result`{.variable} given `settingsObject`{.variable},
        \"`script`\", and `onComplete`{.variable}.

To [fetch a modulepreload module script
graph]{#fetch-a-modulepreload-module-script-graph .dfn} given a
[URL](https://url.spec.whatwg.org/#concept-url){#fetching-scripts:url-7
x-internal="url"} `url`{.variable}, a
[destination](https://fetch.spec.whatwg.org/#concept-request-destination){#fetching-scripts:concept-request-destination-4
x-internal="concept-request-destination"} `destination`{.variable}, an
[environment settings
object](#environment-settings-object){#fetching-scripts:environment-settings-object-8}
`settingsObject`{.variable}, a [script fetch
options](#script-fetch-options){#fetching-scripts:script-fetch-options-6}
`options`{.variable}, and an algorithm `onComplete`{.variable}, run
these steps. `onComplete`{.variable} must be an algorithm accepting null
(on failure) or a [module
script](#module-script){#fetching-scripts:module-script-7} (on success).

1.  [Fetch a single module
    script](#fetch-a-single-module-script){#fetching-scripts:fetch-a-single-module-script-2}
    given `url`{.variable}, `settingsObject`{.variable},
    `destination`{.variable}, `options`{.variable},
    `settingsObject`{.variable}, \"`client`\", true, and with the
    following steps given `result`{.variable}:

    1.  Run `onComplete`{.variable} given `result`{.variable}.

    2.  [Assert](https://infra.spec.whatwg.org/#assert){#fetching-scripts:assert
        x-internal="assert"}: `settingsObject`{.variable}\'s [global
        object](#concept-settings-object-global){#fetching-scripts:concept-settings-object-global-2}
        implements
        [`Window`{#fetching-scripts:window}](nav-history-apis.html#window).

    3.  If `result`{.variable} is not null, optionally [fetch the
        descendants of and
        link](#fetch-the-descendants-of-and-link-a-module-script){#fetching-scripts:fetch-the-descendants-of-and-link-a-module-script-2}
        `result`{.variable} given `settingsObject`{.variable},
        `destination`{.variable}, and an empty algorithm.

        Generally, performing this step will be beneficial for
        performance, as it allows pre-loading the modules that will
        invariably be requested later, via algorithms such as [fetch an
        external module script
        graph](#fetch-a-module-script-tree){#fetching-scripts:fetch-a-module-script-tree-2}
        that fetch the entire graph. However, user agents might wish to
        skip them in bandwidth-constrained situations, or situations
        where the relevant fetches are already in flight.

To [fetch an inline module script
graph]{#fetch-an-inline-module-script-graph .dfn} given a
[string](https://infra.spec.whatwg.org/#string){#fetching-scripts:string
x-internal="string"} `sourceText`{.variable}, a
[URL](https://url.spec.whatwg.org/#concept-url){#fetching-scripts:url-8
x-internal="url"} `baseURL`{.variable}, an [environment settings
object](#environment-settings-object){#fetching-scripts:environment-settings-object-9}
`settingsObject`{.variable}, a [script fetch
options](#script-fetch-options){#fetching-scripts:script-fetch-options-7}
`options`{.variable}, and an algorithm `onComplete`{.variable}, run
these steps. `onComplete`{.variable} must be an algorithm accepting null
(on failure) or a [module
script](#module-script){#fetching-scripts:module-script-8} (on success).

1.  Let `script`{.variable} be the result of [creating a JavaScript
    module
    script](#creating-a-javascript-module-script){#fetching-scripts:creating-a-javascript-module-script}
    using `sourceText`{.variable}, `settingsObject`{.variable},
    `baseURL`{.variable}, and `options`{.variable}.

2.  [Fetch the descendants of and
    link](#fetch-the-descendants-of-and-link-a-module-script){#fetching-scripts:fetch-the-descendants-of-and-link-a-module-script-3}
    `script`{.variable}, given `settingsObject`{.variable},
    \"`script`\", and `onComplete`{.variable}.

To [fetch a module worker script
graph]{#fetch-a-module-worker-script-tree .dfn export=""} given a
[URL](https://url.spec.whatwg.org/#concept-url){#fetching-scripts:url-9
x-internal="url"} `url`{.variable}, an [environment settings
object](#environment-settings-object){#fetching-scripts:environment-settings-object-10}
`fetchClient`{.variable}, a
[destination](https://fetch.spec.whatwg.org/#concept-request-destination){#fetching-scripts:concept-request-destination-5
x-internal="concept-request-destination"} `destination`{.variable}, a
[credentials
mode](#concept-script-fetch-options-credentials){#fetching-scripts:concept-script-fetch-options-credentials-3}
`credentialsMode`{.variable}, an [environment settings
object](#environment-settings-object){#fetching-scripts:environment-settings-object-11}
`settingsObject`{.variable}, and an algorithm `onComplete`{.variable},
[fetch a worklet/module worker script
graph](#fetch-a-worklet/module-worker-script-graph){#fetching-scripts:fetch-a-worklet/module-worker-script-graph}
given `url`{.variable}, `fetchClient`{.variable},
`destination`{.variable}, `credentialsMode`{.variable},
`settingsObject`{.variable}, and `onComplete`{.variable}.

To [fetch a worklet script graph]{#fetch-a-worklet-script-graph .dfn}
given a
[URL](https://url.spec.whatwg.org/#concept-url){#fetching-scripts:url-10
x-internal="url"} `url`{.variable}, an [environment settings
object](#environment-settings-object){#fetching-scripts:environment-settings-object-12}
`fetchClient`{.variable}, a
[destination](https://fetch.spec.whatwg.org/#concept-request-destination){#fetching-scripts:concept-request-destination-6
x-internal="concept-request-destination"} `destination`{.variable}, a
[credentials
mode](#concept-script-fetch-options-credentials){#fetching-scripts:concept-script-fetch-options-credentials-4}
`credentialsMode`{.variable}, an [environment settings
object](#environment-settings-object){#fetching-scripts:environment-settings-object-13}
`settingsObject`{.variable}, a [module responses
map](worklets.html#concept-worklet-module-responses-map){#fetching-scripts:concept-worklet-module-responses-map}
`moduleResponsesMap`{.variable}, and an algorithm
`onComplete`{.variable}, [fetch a worklet/module worker script
graph](#fetch-a-worklet/module-worker-script-graph){#fetching-scripts:fetch-a-worklet/module-worker-script-graph-2}
given `url`{.variable}, `fetchClient`{.variable},
`destination`{.variable}, `credentialsMode`{.variable},
`settingsObject`{.variable}, `onComplete`{.variable}, and the following
[perform the fetch
hook](#fetching-scripts-perform-fetch){#fetching-scripts:fetching-scripts-perform-fetch-5}
given `request`{.variable} and
*[processCustomFetchResponse](#fetching-scripts-processcustomfetchresponse)*:

1.  Let `requestURL`{.variable} be `request`{.variable}\'s
    [URL](https://fetch.spec.whatwg.org/#concept-request-url){#fetching-scripts:concept-request-url-3
    x-internal="concept-request-url"}.

2.  If `moduleResponsesMap`{.variable}\[`requestURL`{.variable}\] is
    \"`fetching`\", wait [in
    parallel](infrastructure.html#in-parallel){#fetching-scripts:in-parallel}
    until that entry\'s value changes, then [queue a
    task](#queue-a-task){#fetching-scripts:queue-a-task} on the
    [networking task
    source](#networking-task-source){#fetching-scripts:networking-task-source}
    to proceed with running the following steps.

3.  If `moduleResponsesMap`{.variable}\[`requestURL`{.variable}\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#fetching-scripts:map-exists-2
    x-internal="map-exists"}, then:

    1.  Let `cached`{.variable} be
        `moduleResponsesMap`{.variable}\[`requestURL`{.variable}\].

    2.  Run `processCustomFetchResponse`{.variable} with
        `cached`{.variable}\[0\] and `cached`{.variable}\[1\].

    3.  Return.

4.  [Set](https://infra.spec.whatwg.org/#map-set){#fetching-scripts:map-set
    x-internal="map-set"}
    `moduleResponsesMap`{.variable}\[`requestURL`{.variable}\] to
    \"`fetching`\".

5.  [Fetch](https://fetch.spec.whatwg.org/#concept-fetch){#fetching-scripts:concept-fetch-7
    x-internal="concept-fetch"} `request`{.variable}, with
    *[processResponseConsumeBody](https://fetch.spec.whatwg.org/#process-response-end-of-body){x-internal="processresponseconsumebody"}*
    set to the following steps given
    [response](https://fetch.spec.whatwg.org/#concept-response){#fetching-scripts:concept-response-10
    x-internal="concept-response"} `response`{.variable} and null,
    failure, or a [byte
    sequence](https://infra.spec.whatwg.org/#byte-sequence){#fetching-scripts:byte-sequence-5
    x-internal="byte-sequence"} `bodyBytes`{.variable}:

    1.  Set `moduleResponsesMap`{.variable}\[`requestURL`{.variable}\]
        to (`response`{.variable}, `bodyBytes`{.variable}).

    2.  Run `processCustomFetchResponse`{.variable} with
        `response`{.variable} and `bodyBytes`{.variable}.

------------------------------------------------------------------------

The following algorithms are meant for internal use by this
specification only as part of [fetching an external module script
graph](#fetch-a-module-script-tree){#fetching-scripts:fetch-a-module-script-tree-3}
or other similar concepts above, and should not be used directly by
other specifications.

This diagram illustrates how these algorithms relate to the ones above,
as well as to each other:

![](data:image/svg+xml;base64,PHN2ZyB2aWV3Ym94PSIwIDAgOTQxIDE2NiIgaWQ9Im1vZHVsZS1zY3JpcHQtZmV0Y2hpbmctZGlhZ3JhbSIgYXJpYS1sYWJlbD0iRmV0Y2ggYW4gZXh0ZXJuYWwgbW9kdWxlIHNjcmlwdCwgZmV0Y2ggYSBtb2R1bGVwcmVsb2FkIG1vZHVsZSBzY3JpcHQgZ3JhcGgsIGZldGNoIGFuIGlubGluZSBtb2R1bGUgc2NyaXB0IGdyYXBoLCBhbmQgZmV0Y2ggYSBtb2R1bGUgd29ya2VyIHNjcmlwdCBncmFwaCBhbGwgY2FsbCBmZXRjaCB0aGUgZGVzY2VuZGFudHMgb2YgYW5kIGxpbmsgYSBtb2R1bGUgc2NyaXB0LiIgcm9sZT0iaW1nIj4KICAgPGRlZnM+CiAgICA8bWFya2VyIG1hcmtlcmhlaWdodD0iNiIgaWQ9Im1vZHVsZS1zY3JpcHQtZmV0Y2hpbmctZGlhZ3JhbS1hcnJvdyIgb3JpZW50PSJhdXRvIiBtYXJrZXJ3aWR0aD0iNiIgcmVmeD0iNSIgcmVmeT0iMyI+CiAgICAgPHBhdGggZD0iTTAsMCBWNiBMNSwzIFoiIHN0eWxlPSJmaWxsOiB2YXIoLS10ZXh0LCBDYW52YXNUZXh0KSIgLz4KICAgIDwvbWFya2VyPgogICA8L2RlZnM+CgogICA8ZyB0cmFuc2Zvcm09InRyYW5zbGF0ZSgwLjUsMC41KSIgY2xhc3M9ImNhbGxlciI+CiAgICA8cmVjdCB3aWR0aD0iMTgwIiBoZWlnaHQ9IjUwIiAvPgogICAgPGZvcmVpZ25vYmplY3Qgd2lkdGg9IjE4MCIgaGVpZ2h0PSI1MCI+CiAgICAgPGEgaHJlZj0iI2ZldGNoLWEtbW9kdWxlLXNjcmlwdC10cmVlIiBpZD0iZmV0Y2hpbmctc2NyaXB0czpmZXRjaC1hLW1vZHVsZS1zY3JpcHQtdHJlZS00IiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMTk5OS94aHRtbCI+ZmV0Y2ggYW4gZXh0ZXJuYWwgbW9kdWxlIHNjcmlwdCBncmFwaDwvYT4KICAgIDwvZm9yZWlnbm9iamVjdD4KICAgPC9nPgogICA8cGF0aCBkPSJNOTAuNSw1MC41IEw0NzAuNSwxNDAuNSIgbWFya2VyLWVuZD0idXJsKCNtb2R1bGUtc2NyaXB0LWZldGNoaW5nLWRpYWdyYW0tYXJyb3cpIiAvPgoKICAgPGcgdHJhbnNmb3JtPSJ0cmFuc2xhdGUoMTkwLjUsMC41KSIgY2xhc3M9ImNhbGxlciI+CiAgICA8cmVjdCB3aWR0aD0iMTgwIiBoZWlnaHQ9IjUwIiAvPgogICAgPGZvcmVpZ25vYmplY3Qgd2lkdGg9IjE4MCIgaGVpZ2h0PSI1MCI+CiAgICAgPGEgaHJlZj0iI2ZldGNoLWEtbW9kdWxlcHJlbG9hZC1tb2R1bGUtc2NyaXB0LWdyYXBoIiBpZD0iZmV0Y2hpbmctc2NyaXB0czpmZXRjaC1hLW1vZHVsZXByZWxvYWQtbW9kdWxlLXNjcmlwdC1ncmFwaCIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzE5OTkveGh0bWwiPmZldGNoIGEgbW9kdWxlcHJlbG9hZCBtb2R1bGUgc2NyaXB0IGdyYXBoPC9hPgogICAgPC9mb3JlaWdub2JqZWN0PgogICA8L2c+CiAgIDxwYXRoIGQ9Ik0yODAuNSw1MC41IEw0NzAuNSwxNDAuNSIgbWFya2VyLWVuZD0idXJsKCNtb2R1bGUtc2NyaXB0LWZldGNoaW5nLWRpYWdyYW0tYXJyb3cpIiAvPgoKICAgPGcgdHJhbnNmb3JtPSJ0cmFuc2xhdGUoMzgwLjUsMC41KSIgY2xhc3M9ImNhbGxlciI+CiAgICA8cmVjdCB3aWR0aD0iMTgwIiBoZWlnaHQ9IjUwIiAvPgogICAgPGZvcmVpZ25vYmplY3Qgd2lkdGg9IjE4MCIgaGVpZ2h0PSI1MCI+CiAgICAgPGEgaHJlZj0iI2ZldGNoLWFuLWlubGluZS1tb2R1bGUtc2NyaXB0LWdyYXBoIiBpZD0iZmV0Y2hpbmctc2NyaXB0czpmZXRjaC1hbi1pbmxpbmUtbW9kdWxlLXNjcmlwdC1ncmFwaCIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzE5OTkveGh0bWwiPmZldGNoIGFuIGlubGluZSBtb2R1bGUgc2NyaXB0IGdyYXBoPC9hPgogICAgPC9mb3JlaWdub2JqZWN0PgogICA8L2c+CiAgIDxwYXRoIGQ9Ik00NzAuNSw1MC41IEw0NzAuNSwxNDAuNSIgbWFya2VyLWVuZD0idXJsKCNtb2R1bGUtc2NyaXB0LWZldGNoaW5nLWRpYWdyYW0tYXJyb3cpIiAvPgoKICAgPGcgdHJhbnNmb3JtPSJ0cmFuc2xhdGUoNTcwLjUsMC41KSIgY2xhc3M9ImNhbGxlciI+CiAgICA8cmVjdCB3aWR0aD0iMTgwIiBoZWlnaHQ9IjUwIiAvPgogICAgPGZvcmVpZ25vYmplY3Qgd2lkdGg9IjE4MCIgaGVpZ2h0PSI1MCI+CiAgICAgPGEgaHJlZj0iI2ZldGNoLWEtbW9kdWxlLXdvcmtlci1zY3JpcHQtdHJlZSIgaWQ9ImZldGNoaW5nLXNjcmlwdHM6ZmV0Y2gtYS1tb2R1bGUtd29ya2VyLXNjcmlwdC10cmVlLTIiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5L3hodG1sIj5mZXRjaCBhIG1vZHVsZSB3b3JrZXIgc2NyaXB0IGdyYXBoPC9hPgogICAgPC9mb3JlaWdub2JqZWN0PgogICA8L2c+CiAgIDxwYXRoIGQ9Ik02NjAuNSw1MC41IEw3NTUuNSw3MC41IiBtYXJrZXItZW5kPSJ1cmwoI21vZHVsZS1zY3JpcHQtZmV0Y2hpbmctZGlhZ3JhbS1hcnJvdykiIC8+CgogICA8ZyB0cmFuc2Zvcm09InRyYW5zbGF0ZSg3NjAuNSwwLjUpIiBjbGFzcz0iY2FsbGVyIj4KICAgIDxyZWN0IHdpZHRoPSIxODAiIGhlaWdodD0iNTAiIC8+CiAgICA8Zm9yZWlnbm9iamVjdCB3aWR0aD0iMTgwIiBoZWlnaHQ9IjUwIj4KICAgICA8YSBocmVmPSIjZmV0Y2gtYS13b3JrbGV0LXNjcmlwdC1ncmFwaCIgaWQ9ImZldGNoaW5nLXNjcmlwdHM6ZmV0Y2gtYS13b3JrbGV0LXNjcmlwdC1ncmFwaCIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzE5OTkveGh0bWwiPmZldGNoIGEgd29ya2xldCBzY3JpcHQgZ3JhcGg8L2E+CiAgICA8L2ZvcmVpZ25vYmplY3Q+CiAgIDwvZz4KICAgPHBhdGggZD0iTTg1MC41LDUwLjUgTDc1NS41LDcwLjUiIG1hcmtlci1lbmQ9InVybCgjbW9kdWxlLXNjcmlwdC1mZXRjaGluZy1kaWFncmFtLWFycm93KSIgLz4KCiAgIDxnIHRyYW5zZm9ybT0idHJhbnNsYXRlKDY2NS41LDcwLjUpIiBjbGFzcz0iY2FsbGVyIj4KICAgIDxyZWN0IHdpZHRoPSIxODAiIGhlaWdodD0iNTAiIC8+CiAgICA8Zm9yZWlnbm9iamVjdCB3aWR0aD0iMTgwIiBoZWlnaHQ9IjUwIj4KICAgICA8YSBocmVmPSIjZmV0Y2gtYS13b3JrbGV0L21vZHVsZS13b3JrZXItc2NyaXB0LWdyYXBoIiBpZD0iZmV0Y2hpbmctc2NyaXB0czpmZXRjaC1hLXdvcmtsZXQvbW9kdWxlLXdvcmtlci1zY3JpcHQtZ3JhcGgtMyIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzE5OTkveGh0bWwiPmZldGNoIGEgd29ya2xldC9tb2R1bGUgd29ya2VyIHNjcmlwdCBncmFwaDwvYT4KICAgIDwvZm9yZWlnbm9iamVjdD4KICAgPC9nPgogICA8cGF0aCBkPSJNNzU1LjUsMTIwLjUgTDQ3MC41LDE0MC41IiBtYXJrZXItZW5kPSJ1cmwoI21vZHVsZS1zY3JpcHQtZmV0Y2hpbmctZGlhZ3JhbS1hcnJvdykiIC8+CgogICA8ZyB0cmFuc2Zvcm09InRyYW5zbGF0ZSgyNjAuNSwxNDAuNSkiIGNsYXNzPSJzdWJhbGdvcml0aG0iPgogICAgPHJlY3Qgd2lkdGg9IjQyMCIgaGVpZ2h0PSIyNSIgLz4KICAgIDxmb3JlaWdub2JqZWN0IHdpZHRoPSI0MjAiIGhlaWdodD0iMjUiPgogICAgIDxhIGhyZWY9IiNmZXRjaC10aGUtZGVzY2VuZGFudHMtb2YtYW5kLWxpbmstYS1tb2R1bGUtc2NyaXB0IiBpZD0iZmV0Y2hpbmctc2NyaXB0czpmZXRjaC10aGUtZGVzY2VuZGFudHMtb2YtYW5kLWxpbmstYS1tb2R1bGUtc2NyaXB0LTQiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5L3hodG1sIj5mZXRjaCB0aGUgZGVzY2VuZGFudHMgb2YgYW5kIGxpbmsgYSBtb2R1bGUgc2NyaXB0PC9hPgogICAgPC9mb3JlaWdub2JqZWN0PgogICA8L2c+CiAgPC9zdmc+){#module-script-fetching-diagram}

To [fetch a worklet/module worker script
graph]{#fetch-a-worklet/module-worker-script-graph .dfn} given a
[URL](https://url.spec.whatwg.org/#concept-url){#fetching-scripts:url-11
x-internal="url"} `url`{.variable}, an [environment settings
object](#environment-settings-object){#fetching-scripts:environment-settings-object-14}
`fetchClient`{.variable}, a
[destination](https://fetch.spec.whatwg.org/#concept-request-destination){#fetching-scripts:concept-request-destination-7
x-internal="concept-request-destination"} `destination`{.variable}, a
[credentials
mode](#concept-script-fetch-options-credentials){#fetching-scripts:concept-script-fetch-options-credentials-5}
`credentialsMode`{.variable}, an [environment settings
object](#environment-settings-object){#fetching-scripts:environment-settings-object-15}
`settingsObject`{.variable}, an algorithm `onComplete`{.variable}, and
an optional [perform the fetch
hook](#fetching-scripts-perform-fetch){#fetching-scripts:fetching-scripts-perform-fetch-6}
`performFetch`{.variable}, run these steps. `onComplete`{.variable} must
be an algorithm accepting null (on failure) or a [module
script](#module-script){#fetching-scripts:module-script-9} (on success).

1.  Let `options`{.variable} be a [script fetch
    options](#script-fetch-options){#fetching-scripts:script-fetch-options-8}
    whose [cryptographic
    nonce](#concept-script-fetch-options-nonce){#fetching-scripts:concept-script-fetch-options-nonce-4}
    is the empty string, [integrity
    metadata](#concept-script-fetch-options-integrity){#fetching-scripts:concept-script-fetch-options-integrity-5}
    is the empty string, [parser
    metadata](#concept-script-fetch-options-parser){#fetching-scripts:concept-script-fetch-options-parser-4}
    is \"`not-parser-inserted`\", [credentials
    mode](#concept-script-fetch-options-credentials){#fetching-scripts:concept-script-fetch-options-credentials-6}
    is `credentialsMode`{.variable}, [referrer
    policy](#concept-script-fetch-options-referrer-policy){#fetching-scripts:concept-script-fetch-options-referrer-policy-4}
    is the empty string, and [fetch
    priority](#concept-script-fetch-options-fetch-priority){#fetching-scripts:concept-script-fetch-options-fetch-priority-5}
    is \"`auto`\".

2.  [Fetch a single module
    script](#fetch-a-single-module-script){#fetching-scripts:fetch-a-single-module-script-3}
    given `url`{.variable}, `fetchClient`{.variable},
    `destination`{.variable}, `options`{.variable},
    `settingsObject`{.variable}, \"`client`\", true, and
    `onSingleFetchComplete`{.variable} as defined below. If
    `performFetch`{.variable} was given, pass it along as well.

    `onSingleFetchComplete`{.variable} given `result`{.variable} is the
    following algorithm:

    1.  If `result`{.variable} is null, run `onComplete`{.variable}
        given null, and abort these steps.

    2.  [Fetch the descendants of and
        link](#fetch-the-descendants-of-and-link-a-module-script){#fetching-scripts:fetch-the-descendants-of-and-link-a-module-script-5}
        `result`{.variable} given `fetchClient`{.variable},
        `destination`{.variable}, and `onComplete`{.variable}. If
        `performFetch`{.variable} was given, pass it along as well.

To [fetch the descendants of and
link]{#fetch-the-descendants-of-and-link-a-module-script .dfn} a [module
script](#module-script){#fetching-scripts:module-script-10}
`moduleScript`{.variable}, given an [environment settings
object](#environment-settings-object){#fetching-scripts:environment-settings-object-16}
`fetchClient`{.variable}, a
[destination](https://fetch.spec.whatwg.org/#concept-request-destination){#fetching-scripts:concept-request-destination-8
x-internal="concept-request-destination"} `destination`{.variable}, an
algorithm `onComplete`{.variable}, and an optional [perform the fetch
hook](#fetching-scripts-perform-fetch){#fetching-scripts:fetching-scripts-perform-fetch-7}
`performFetch`{.variable}, run these steps. `onComplete`{.variable} must
be an algorithm accepting null (on failure) or a [module
script](#module-script){#fetching-scripts:module-script-11} (on
success).

1.  Let `record`{.variable} be `moduleScript`{.variable}\'s
    [record](#concept-script-record){#fetching-scripts:concept-script-record}.

2.  If `record`{.variable} is null, then:

    1.  Set `moduleScript`{.variable}\'s [error to
        rethrow](#concept-script-error-to-rethrow){#fetching-scripts:concept-script-error-to-rethrow}
        to `moduleScript`{.variable}\'s [parse
        error](parsing.html#parse-errors){#fetching-scripts:parse-errors}.

    2.  Run `onComplete`{.variable} given `moduleScript`{.variable}.

    3.  Return.

3.  Let `state`{.variable} be
    [Record](https://tc39.es/ecma262/#sec-list-and-record-specification-type){#fetching-scripts:record
    x-internal="record"} { \[\[ParseError\]\]: null,
    \[\[Destination\]\]: `destination`{.variable}, \[\[PerformFetch\]\]:
    null, \[\[FetchClient\]\]: `fetchClient`{.variable} }.

4.  If `performFetch`{.variable} was given, set
    `state`{.variable}.\[\[PerformFetch\]\] to
    `performFetch`{.variable}.

5.  Let `loadingPromise`{.variable} be
    `record`{.variable}.[LoadRequestedModules](https://tc39.es/ecma262/#sec-LoadRequestedModules){#fetching-scripts:js-loadrequestedmodules
    x-internal="js-loadrequestedmodules"}(`state`{.variable}).

    This step will recursively load all the module transitive
    dependencies.

6.  [Upon
    fulfillment](https://webidl.spec.whatwg.org/#upon-fulfillment){#fetching-scripts:upon-fulfillment
    x-internal="upon-fulfillment"} of `loadingPromise`{.variable}, run
    the following steps:

    1.  Perform
        `record`{.variable}.[Link](https://tc39.es/ecma262/#sec-moduledeclarationlinking){#fetching-scripts:js-link
        x-internal="js-link"}().

        This step will recursively call
        [Link](https://tc39.es/ecma262/#sec-moduledeclarationlinking){#fetching-scripts:js-link-2
        x-internal="js-link"} on all of the module\'s unlinked
        dependencies.

        If this throws an exception, catch it, and set
        `moduleScript`{.variable}\'s [error to
        rethrow](#concept-script-error-to-rethrow){#fetching-scripts:concept-script-error-to-rethrow-2}
        to that exception.

    2.  Run `onComplete`{.variable} given `moduleScript`{.variable}.

7.  [Upon
    rejection](https://webidl.spec.whatwg.org/#upon-rejection){#fetching-scripts:upon-rejection
    x-internal="upon-rejection"} of `loadingPromise`{.variable}, run the
    following steps:

    1.  If `state`{.variable}.\[\[ParseError\]\] is not null, set
        `moduleScript`{.variable}\'s [error to
        rethrow](#concept-script-error-to-rethrow){#fetching-scripts:concept-script-error-to-rethrow-3}
        to `state`{.variable}.\[\[ParseError\]\] and run
        `onComplete`{.variable} given `moduleScript`{.variable}.

    2.  Otherwise, run `onComplete`{.variable} given null.

        `state`{.variable}.\[\[ParseError\]\] is null when
        `loadingPromise`{.variable} is rejected due to a loading error.

To [fetch a single module script]{#fetch-a-single-module-script .dfn},
given a
[URL](https://url.spec.whatwg.org/#concept-url){#fetching-scripts:url-12
x-internal="url"} `url`{.variable}, an [environment settings
object](#environment-settings-object){#fetching-scripts:environment-settings-object-17}
`fetchClient`{.variable}, a
[destination](https://fetch.spec.whatwg.org/#concept-request-destination){#fetching-scripts:concept-request-destination-9
x-internal="concept-request-destination"} `destination`{.variable}, a
[script fetch
options](#script-fetch-options){#fetching-scripts:script-fetch-options-9}
`options`{.variable}, an [environment settings
object](#environment-settings-object){#fetching-scripts:environment-settings-object-18}
`settingsObject`{.variable}, a
[referrer](https://fetch.spec.whatwg.org/#concept-request-referrer){#fetching-scripts:concept-request-referrer
x-internal="concept-request-referrer"} `referrer`{.variable}, an
optional [ModuleRequest
Record](https://tc39.es/ecma262/#modulerequest-record){#fetching-scripts:modulerequest-record
x-internal="modulerequest-record"} `moduleRequest`{.variable}, a boolean
[`isTopLevel`{#fetching-scripts:fetching-scripts-is-top-level-2
.variable}](#fetching-scripts-is-top-level), an algorithm
`onComplete`{.variable}, and an optional [perform the fetch
hook](#fetching-scripts-perform-fetch){#fetching-scripts:fetching-scripts-perform-fetch-8}
`performFetch`{.variable}, run these steps. `onComplete`{.variable} must
be an algorithm accepting null (on failure) or a [module
script](#module-script){#fetching-scripts:module-script-12} (on
success).

1.  Let `moduleType`{.variable} be \"`javascript-or-wasm`\".

2.  If `moduleRequest`{.variable} was given, then set
    `moduleType`{.variable} to the result of running the [module type
    from module
    request](#module-type-from-module-request){#fetching-scripts:module-type-from-module-request}
    steps given `moduleRequest`{.variable}.

3.  [Assert](https://infra.spec.whatwg.org/#assert){#fetching-scripts:assert-2
    x-internal="assert"}: the result of running the [module type
    allowed](#module-type-allowed){#fetching-scripts:module-type-allowed}
    steps given `moduleType`{.variable} and `settingsObject`{.variable}
    is true. Otherwise, we would not have reached this point because a
    failure would have been raised when inspecting
    `moduleRequest`{.variable}.\[\[Attributes\]\] in
    [HostLoadImportedModule](#validate-requested-module-specifiers) or
    [fetch a single imported module
    script](#fetch-a-single-imported-module-script){#fetching-scripts:fetch-a-single-imported-module-script}.

4.  Let `moduleMap`{.variable} be `settingsObject`{.variable}\'s [module
    map](#concept-settings-object-module-map){#fetching-scripts:concept-settings-object-module-map}.

5.  If `moduleMap`{.variable}\[(`url`{.variable},
    `moduleType`{.variable})\] is \"`fetching`\", wait [in
    parallel](infrastructure.html#in-parallel){#fetching-scripts:in-parallel-2}
    until that entry\'s value changes, then [queue a
    task](#queue-a-task){#fetching-scripts:queue-a-task-2} on the
    [networking task
    source](#networking-task-source){#fetching-scripts:networking-task-source-2}
    to proceed with running the following steps.

6.  If `moduleMap`{.variable}\[(`url`{.variable},
    `moduleType`{.variable})\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#fetching-scripts:map-exists-3
    x-internal="map-exists"}, run `onComplete`{.variable} given
    `moduleMap`{.variable}\[(`url`{.variable},
    `moduleType`{.variable})\], and return.

7.  [Set](https://infra.spec.whatwg.org/#map-set){#fetching-scripts:map-set-2
    x-internal="map-set"} `moduleMap`{.variable}\[(`url`{.variable},
    `moduleType`{.variable})\] to \"`fetching`\".

8.  Let `request`{.variable} be a new
    [request](https://fetch.spec.whatwg.org/#concept-request){#fetching-scripts:concept-request-8
    x-internal="concept-request"} whose
    [URL](https://fetch.spec.whatwg.org/#concept-request-url){#fetching-scripts:concept-request-url-4
    x-internal="concept-request-url"} is `url`{.variable},
    [mode](https://fetch.spec.whatwg.org/#concept-request-mode){#fetching-scripts:concept-request-mode-2
    x-internal="concept-request-mode"} is \"`cors`\",
    [referrer](https://fetch.spec.whatwg.org/#concept-request-referrer){#fetching-scripts:concept-request-referrer-2
    x-internal="concept-request-referrer"} is `referrer`{.variable}, and
    [client](https://fetch.spec.whatwg.org/#concept-request-client){#fetching-scripts:concept-request-client-4
    x-internal="concept-request-client"} is `fetchClient`{.variable}.

9.  Set `request`{.variable}\'s
    [destination](https://fetch.spec.whatwg.org/#concept-request-destination){#fetching-scripts:concept-request-destination-10
    x-internal="concept-request-destination"} to the result of running
    the [fetch destination from module
    type](#fetch-destination-from-module-type){#fetching-scripts:fetch-destination-from-module-type}
    steps given `destination`{.variable} and `moduleType`{.variable}.

10. If `destination`{.variable} is \"`worker`\", \"`sharedworker`\", or
    \"`serviceworker`\", and `isTopLevel`{.variable} is true, then set
    `request`{.variable}\'s
    [mode](https://fetch.spec.whatwg.org/#concept-request-mode){#fetching-scripts:concept-request-mode-3
    x-internal="concept-request-mode"} to \"`same-origin`\".

11. Set `request`{.variable}\'s [initiator
    type](https://fetch.spec.whatwg.org/#request-initiator-type){#fetching-scripts:concept-request-initiator-type-4
    x-internal="concept-request-initiator-type"} to \"`script`\".

12. [Set up the module script
    request](#set-up-the-module-script-request){#fetching-scripts:set-up-the-module-script-request}
    given `request`{.variable} and `options`{.variable}.

13. If `performFetch`{.variable} was given, run
    `performFetch`{.variable} with `request`{.variable},
    `isTopLevel`{.variable}, and with
    `processResponseConsumeBody`{.variable} as defined below.

    Otherwise,
    [fetch](https://fetch.spec.whatwg.org/#concept-fetch){#fetching-scripts:concept-fetch-8
    x-internal="concept-fetch"} `request`{.variable} with
    *[processResponseConsumeBody](https://fetch.spec.whatwg.org/#process-response-end-of-body){x-internal="processresponseconsumebody"}*
    set to `processResponseConsumeBody`{.variable} as defined below.

    In both cases, let `processResponseConsumeBody`{.variable} given
    [response](https://fetch.spec.whatwg.org/#concept-response){#fetching-scripts:concept-response-11
    x-internal="concept-response"} `response`{.variable} and null,
    failure, or a [byte
    sequence](https://infra.spec.whatwg.org/#byte-sequence){#fetching-scripts:byte-sequence-6
    x-internal="byte-sequence"} `bodyBytes`{.variable} be the following
    algorithm:

    `response`{.variable} is always
    [CORS-same-origin](urls-and-fetching.html#cors-same-origin){#fetching-scripts:cors-same-origin-2}.

    1.  If any of the following are true:

        - `bodyBytes`{.variable} is null or failure; or

        - `response`{.variable}\'s
          [status](https://fetch.spec.whatwg.org/#concept-response-status){#fetching-scripts:concept-response-status-4
          x-internal="concept-response-status"} is not an [ok
          status](https://fetch.spec.whatwg.org/#ok-status){#fetching-scripts:ok-status-4
          x-internal="ok-status"},

        then
        [set](https://infra.spec.whatwg.org/#map-set){#fetching-scripts:map-set-3
        x-internal="map-set"} `moduleMap`{.variable}\[(`url`{.variable},
        `moduleType`{.variable})\] to null, run `onComplete`{.variable}
        given null, and abort these steps.

    2.  Let `mimeType`{.variable} be the result of [extracting a MIME
        type](https://fetch.spec.whatwg.org/#concept-header-extract-mime-type){#fetching-scripts:extract-a-mime-type-4
        x-internal="extract-a-mime-type"} from `response`{.variable}\'s
        [header
        list](https://fetch.spec.whatwg.org/#concept-response-header-list){#fetching-scripts:concept-response-header-list-4
        x-internal="concept-response-header-list"}.

    3.  Let `moduleScript`{.variable} be null.

    4.  Let `referrerPolicy`{.variable} be the result of [parsing the
        \``Referrer-Policy`\`
        header](https://w3c.github.io/webappsec-referrer-policy/#parse-referrer-policy-from-header){#fetching-scripts:parse-referrer-policy-header-2
        x-internal="parse-referrer-policy-header"} given
        `response`{.variable}.
        [\[REFERRERPOLICY\]](references.html#refsREFERRERPOLICY)

    5.  If `referrerPolicy`{.variable} is not the empty string, set
        `options`{.variable}\'s [referrer
        policy](#concept-script-fetch-options-referrer-policy){#fetching-scripts:concept-script-fetch-options-referrer-policy-5}
        to `referrerPolicy`{.variable}.

    6.  If `mimeType`{.variable}\'s
        [essence](https://mimesniff.spec.whatwg.org/#mime-type-essence){#fetching-scripts:mime-type-essence-2
        x-internal="mime-type-essence"} is
        \"[`application/wasm`{#fetching-scripts:application/wasm}](indices.html#application/wasm)\"
        and `moduleType`{.variable} is \"`javascript-or-wasm`\", then
        set `moduleScript`{.variable} to the result of [creating a
        WebAssembly module
        script](#creating-a-webassembly-module-script){#fetching-scripts:creating-a-webassembly-module-script}
        given `bodyBytes`{.variable}, `settingsObject`{.variable},
        `response`{.variable}\'s
        [URL](https://fetch.spec.whatwg.org/#concept-response-url){#fetching-scripts:concept-response-url-5
        x-internal="concept-response-url"}, and `options`{.variable}.

    7.  Otherwise:

        1.  Let `sourceText`{.variable} be the result of [UTF-8
            decoding](https://encoding.spec.whatwg.org/#utf-8-decode){#fetching-scripts:utf-8-decode-3
            x-internal="utf-8-decode"} `bodyBytes`{.variable}.

        2.  If `mimeType`{.variable} is a [JavaScript MIME
            type](https://mimesniff.spec.whatwg.org/#javascript-mime-type){#fetching-scripts:javascript-mime-type-3
            x-internal="javascript-mime-type"} and
            `moduleType`{.variable} is \"`javascript-or-wasm`\", then
            set `moduleScript`{.variable} to the result of [creating a
            JavaScript module
            script](#creating-a-javascript-module-script){#fetching-scripts:creating-a-javascript-module-script-2}
            given `sourceText`{.variable}, `settingsObject`{.variable},
            `response`{.variable}\'s
            [URL](https://fetch.spec.whatwg.org/#concept-response-url){#fetching-scripts:concept-response-url-6
            x-internal="concept-response-url"}, and
            `options`{.variable}.

        3.  If the [MIME type
            essence](https://mimesniff.spec.whatwg.org/#mime-type-essence){#fetching-scripts:mime-type-essence-3
            x-internal="mime-type-essence"} of `mimeType`{.variable} is
            \"[`text/css`{#fetching-scripts:text/css}](indices.html#text/css)\"
            and `moduleType`{.variable} is \"`css`\", then set
            `moduleScript`{.variable} to the result of [creating a CSS
            module
            script](#creating-a-css-module-script){#fetching-scripts:creating-a-css-module-script}
            given `sourceText`{.variable} and
            `settingsObject`{.variable}.

        4.  If `mimeType`{.variable} is a [JSON MIME
            type](https://mimesniff.spec.whatwg.org/#json-mime-type){#fetching-scripts:json-mime-type
            x-internal="json-mime-type"} and `moduleType`{.variable} is
            \"`json`\", then set `moduleScript`{.variable} to the result
            of [creating a JSON module
            script](#creating-a-json-module-script){#fetching-scripts:creating-a-json-module-script}
            given `sourceText`{.variable} and
            `settingsObject`{.variable}.

    8.  [Set](https://infra.spec.whatwg.org/#map-set){#fetching-scripts:map-set-4
        x-internal="map-set"} `moduleMap`{.variable}\[(`url`{.variable},
        `moduleType`{.variable})\] to `moduleScript`{.variable}, and run
        `onComplete`{.variable} given `moduleScript`{.variable}.

        It is intentional that the [module
        map](#module-map){#fetching-scripts:module-map} is keyed by the
        [request
        URL](https://fetch.spec.whatwg.org/#concept-request-url){#fetching-scripts:concept-request-url-5
        x-internal="concept-request-url"}, whereas the [base
        URL](#concept-script-base-url){#fetching-scripts:concept-script-base-url}
        for the [module
        script](#module-script){#fetching-scripts:module-script-13} is
        set to the [response
        URL](https://fetch.spec.whatwg.org/#concept-response-url){#fetching-scripts:concept-response-url-7
        x-internal="concept-response-url"}. The former is used to
        deduplicate fetches, while the latter is used for URL
        resolution.

To [fetch a single imported module
script]{#fetch-a-single-imported-module-script .dfn}, given a
[URL](https://url.spec.whatwg.org/#concept-url){#fetching-scripts:url-13
x-internal="url"} `url`{.variable}, an [environment settings
object](#environment-settings-object){#fetching-scripts:environment-settings-object-19}
`fetchClient`{.variable}, a
[destination](https://fetch.spec.whatwg.org/#concept-request-destination){#fetching-scripts:concept-request-destination-11
x-internal="concept-request-destination"} `destination`{.variable}, a
[script fetch
options](#script-fetch-options){#fetching-scripts:script-fetch-options-10}
`options`{.variable}, [environment settings
object](#environment-settings-object){#fetching-scripts:environment-settings-object-20}
`settingsObject`{.variable}, a
[referrer](https://fetch.spec.whatwg.org/#concept-request-referrer){#fetching-scripts:concept-request-referrer-3
x-internal="concept-request-referrer"} `referrer`{.variable}, a
[ModuleRequest
Record](https://tc39.es/ecma262/#modulerequest-record){#fetching-scripts:modulerequest-record-2
x-internal="modulerequest-record"} `moduleRequest`{.variable}, an
algorithm `onComplete`{.variable}, and an optional [perform the fetch
hook](#fetching-scripts-perform-fetch){#fetching-scripts:fetching-scripts-perform-fetch-9}
`performFetch`{.variable}, run these steps. `onComplete`{.variable} must
be an algorithm accepting null (on failure) or a [module
script](#module-script){#fetching-scripts:module-script-14} (on
success).

1.  [Assert](https://infra.spec.whatwg.org/#assert){#fetching-scripts:assert-3
    x-internal="assert"}: `moduleRequest`{.variable}.\[\[Attributes\]\]
    does not contain any
    [Record](https://tc39.es/ecma262/#sec-list-and-record-specification-type){#fetching-scripts:record-2
    x-internal="record"} `entry`{.variable} such that
    `entry`{.variable}.\[\[Key\]\] is not \"`type`\", because we only
    asked for \"`type`\" attributes in
    [HostGetSupportedImportAttributes](#hostgetsupportedimportattributes){#fetching-scripts:hostgetsupportedimportattributes}.

2.  Let `moduleType`{.variable} be the result of running the [module
    type from module
    request](#module-type-from-module-request){#fetching-scripts:module-type-from-module-request-2}
    steps given `moduleRequest`{.variable}.

3.  If the result of running the [module type
    allowed](#module-type-allowed){#fetching-scripts:module-type-allowed-2}
    steps given `moduleType`{.variable} and `settingsObject`{.variable}
    is false, then run `onComplete`{.variable} given null, and return.

4.  [Fetch a single module
    script](#fetch-a-single-module-script){#fetching-scripts:fetch-a-single-module-script-4}
    given `url`{.variable}, `fetchClient`{.variable},
    `destination`{.variable}, `options`{.variable},
    `settingsObject`{.variable}, `referrer`{.variable},
    `moduleRequest`{.variable}, false, and `onComplete`{.variable}. If
    `performFetch`{.variable} was given, pass it along as well.

##### [8.1.4.3]{.secno} Creating scripts[](#creating-scripts){.self-link}

To [create a classic script]{#creating-a-classic-script .dfn}, given a
[string](https://infra.spec.whatwg.org/#string){#creating-scripts:string
x-internal="string"} `source`{.variable}, an [environment settings
object](#environment-settings-object){#creating-scripts:environment-settings-object}
`settings`{.variable}, a
[URL](https://url.spec.whatwg.org/#concept-url){#creating-scripts:url
x-internal="url"} `baseURL`{.variable}, a [script fetch
options](#script-fetch-options){#creating-scripts:script-fetch-options}
`options`{.variable}, an optional boolean `mutedErrors`{.variable}
(default false), and an optional
[URL](https://url.spec.whatwg.org/#concept-url){#creating-scripts:url-2
x-internal="url"}-or-null `sourceURLForWindowScripts`{.variable}
(default null):

1.  If `mutedErrors`{.variable} is true, then set `baseURL`{.variable}
    to
    [`about:blank`{#creating-scripts:about:blank}](infrastructure.html#about:blank).

    When `mutedErrors`{.variable} is true, `baseURL`{.variable} is the
    script\'s
    [CORS-cross-origin](urls-and-fetching.html#cors-cross-origin){#creating-scripts:cors-cross-origin}
    [response](https://fetch.spec.whatwg.org/#concept-response){#creating-scripts:concept-response
    x-internal="concept-response"}\'s
    [url](https://fetch.spec.whatwg.org/#concept-response-url){#creating-scripts:concept-response-url
    x-internal="concept-response-url"}, which shouldn\'t be exposed to
    JavaScript. Therefore, `baseURL`{.variable} is sanitized here.

2.  If [scripting is
    disabled](#concept-environment-noscript){#creating-scripts:concept-environment-noscript}
    for `settings`{.variable}, then set `source`{.variable} to the empty
    string.

3.  Let `script`{.variable} be a new [classic
    script](#classic-script){#creating-scripts:classic-script} that this
    algorithm will subsequently initialize.

4.  Set `script`{.variable}\'s [settings
    object](#settings-object){#creating-scripts:settings-object} to
    `settings`{.variable}.

5.  Set `script`{.variable}\'s [base
    URL](#concept-script-base-url){#creating-scripts:concept-script-base-url}
    to `baseURL`{.variable}.

6.  Set `script`{.variable}\'s [fetch
    options](#concept-script-script-fetch-options){#creating-scripts:concept-script-script-fetch-options}
    to `options`{.variable}.

7.  Set `script`{.variable}\'s [muted
    errors](#muted-errors){#creating-scripts:muted-errors} to
    `mutedErrors`{.variable}.

8.  Set `script`{.variable}\'s [parse
    error](#concept-script-parse-error){#creating-scripts:concept-script-parse-error}
    and [error to
    rethrow](#concept-script-error-to-rethrow){#creating-scripts:concept-script-error-to-rethrow}
    to null.

9.  [Record classic script creation
    time](https://w3c.github.io/long-animation-frames/#record-classic-script-creation-time){#creating-scripts:record-classic-script-creation-time
    x-internal="record-classic-script-creation-time"} given
    `script`{.variable} and `sourceURLForWindowScripts`{.variable}.

10. Let `result`{.variable} be
    [ParseScript](https://tc39.es/ecma262/#sec-parse-script){#creating-scripts:js-parsescript
    x-internal="js-parsescript"}(`source`{.variable},
    `settings`{.variable}\'s
    [realm](#environment-settings-object's-realm){#creating-scripts:environment-settings-object's-realm},
    `script`{.variable}).

    Passing `script`{.variable} as the last parameter here ensures
    `result`{.variable}.\[\[HostDefined\]\] will be `script`{.variable}.

11. If `result`{.variable} is a
    [list](https://infra.spec.whatwg.org/#list){#creating-scripts:list
    x-internal="list"} of errors, then:

    1.  Set `script`{.variable}\'s [parse
        error](#concept-script-parse-error){#creating-scripts:concept-script-parse-error-2}
        and its [error to
        rethrow](#concept-script-error-to-rethrow){#creating-scripts:concept-script-error-to-rethrow-2}
        to `result`{.variable}\[0\].

    2.  Return `script`{.variable}.

12. Set `script`{.variable}\'s
    [record](#concept-script-record){#creating-scripts:concept-script-record}
    to `result`{.variable}.

13. Return `script`{.variable}.

To [create a JavaScript module
script]{#creating-a-javascript-module-script .dfn}, given a
[string](https://infra.spec.whatwg.org/#string){#creating-scripts:string-2
x-internal="string"} `source`{.variable}, an [environment settings
object](#environment-settings-object){#creating-scripts:environment-settings-object-2}
`settings`{.variable}, a
[URL](https://url.spec.whatwg.org/#concept-url){#creating-scripts:url-3
x-internal="url"} `baseURL`{.variable}, and a [script fetch
options](#script-fetch-options){#creating-scripts:script-fetch-options-2}
`options`{.variable}:

1.  If [scripting is
    disabled](#concept-environment-noscript){#creating-scripts:concept-environment-noscript-2}
    for `settings`{.variable}, then set `source`{.variable} to the empty
    string.

2.  Let `script`{.variable} be a new [module
    script](#module-script){#creating-scripts:module-script} that this
    algorithm will subsequently initialize.

3.  Set `script`{.variable}\'s [settings
    object](#settings-object){#creating-scripts:settings-object-2} to
    `settings`{.variable}.

4.  Set `script`{.variable}\'s [base
    URL](#concept-script-base-url){#creating-scripts:concept-script-base-url-2}
    to `baseURL`{.variable}.

5.  Set `script`{.variable}\'s [fetch
    options](#concept-script-script-fetch-options){#creating-scripts:concept-script-script-fetch-options-2}
    to `options`{.variable}.

6.  Set `script`{.variable}\'s [parse
    error](#concept-script-parse-error){#creating-scripts:concept-script-parse-error-3}
    and [error to
    rethrow](#concept-script-error-to-rethrow){#creating-scripts:concept-script-error-to-rethrow-3}
    to null.

7.  Let `result`{.variable} be
    [ParseModule](https://tc39.es/ecma262/#sec-parsemodule){#creating-scripts:js-parsemodule
    x-internal="js-parsemodule"}(`source`{.variable},
    `settings`{.variable}\'s
    [realm](#environment-settings-object's-realm){#creating-scripts:environment-settings-object's-realm-2},
    `script`{.variable}).

    Passing `script`{.variable} as the last parameter here ensures
    `result`{.variable}.\[\[HostDefined\]\] will be `script`{.variable}.

8.  If `result`{.variable} is a
    [list](https://infra.spec.whatwg.org/#list){#creating-scripts:list-2
    x-internal="list"} of errors, then:

    1.  Set `script`{.variable}\'s [parse
        error](#concept-script-parse-error){#creating-scripts:concept-script-parse-error-4}
        to `result`{.variable}\[0\].

    2.  Return `script`{.variable}.

9.  Set `script`{.variable}\'s
    [record](#concept-script-record){#creating-scripts:concept-script-record-2}
    to `result`{.variable}.

10. Return `script`{.variable}.

To [create a WebAssembly module
script]{#creating-a-webassembly-module-script .dfn}, given a [byte
sequence](https://infra.spec.whatwg.org/#byte-sequence){#creating-scripts:byte-sequence
x-internal="byte-sequence"} `bodyBytes`{.variable}, an [environment
settings
object](#environment-settings-object){#creating-scripts:environment-settings-object-3}
`settings`{.variable}, a
[URL](https://url.spec.whatwg.org/#concept-url){#creating-scripts:url-4
x-internal="url"} `baseURL`{.variable}, and a [script fetch
options](#script-fetch-options){#creating-scripts:script-fetch-options-3}
`options`{.variable}:

1.  If [scripting is
    disabled](#concept-environment-noscript){#creating-scripts:concept-environment-noscript-3}
    for `settings`{.variable}, then set `bodyBytes`{.variable} to the
    byte sequence 0x00 0x61 0x73 0x6D 0x01 0x00 0x00 0x00.

    This byte sequence corresponds to an empty WebAssembly module with
    only the magic bytes and version number provided.

2.  Let `script`{.variable} be a new [module
    script](#module-script){#creating-scripts:module-script-2} that this
    algorithm will subsequently initialize.

3.  Set `script`{.variable}\'s [settings
    object](#settings-object){#creating-scripts:settings-object-3} to
    `settings`{.variable}.

4.  Set `script`{.variable}\'s [base
    URL](#concept-script-base-url){#creating-scripts:concept-script-base-url-3}
    to `baseURL`{.variable}.

5.  Set `script`{.variable}\'s [fetch
    options](#concept-script-script-fetch-options){#creating-scripts:concept-script-script-fetch-options-3}
    to `options`{.variable}.

6.  Set `script`{.variable}\'s [parse
    error](#concept-script-parse-error){#creating-scripts:concept-script-parse-error-5}
    and [error to
    rethrow](#concept-script-error-to-rethrow){#creating-scripts:concept-script-error-to-rethrow-4}
    to null.

7.  Let `result`{.variable} be the result of [parsing a WebAssembly
    module](https://webassembly.github.io/esm-integration/js-api/index.html#parse-a-webassembly-module){#creating-scripts:parse-a-webassembly-module
    x-internal="parse-a-webassembly-module"} given
    `bodyBytes`{.variable}, `settings`{.variable}\'s
    [realm](#environment-settings-object's-realm){#creating-scripts:environment-settings-object's-realm-3},
    and `script`{.variable}.

    Passing `script`{.variable} as the last parameter here ensures
    `result`{.variable}.\[\[HostDefined\]\] will be `script`{.variable}.

8.  If the previous step threw an error `error`{.variable}, then:

    1.  Set `script`{.variable}\'s [parse
        error](#concept-script-parse-error){#creating-scripts:concept-script-parse-error-6}
        to `error`{.variable}.

    2.  Return `script`{.variable}.

9.  Set `script`{.variable}\'s
    [record](#concept-script-record){#creating-scripts:concept-script-record-3}
    to `result`{.variable}.

10. Return `script`{.variable}.

WebAssembly JavaScript Interface: ESM Integration specifies the hooks
for the WebAssembly integration with ECMA-262 module loading. This
includes support both for direct dependency imports, as well as for
source phase imports, which support virtualization and
multi-instantiation. [\[WASMESM\]](references.html#refsWASMESM)

To [create a CSS module script]{#creating-a-css-module-script .dfn},
given a string `source`{.variable} and an [environment settings
object](#environment-settings-object){#creating-scripts:environment-settings-object-4}
`settings`{.variable}:

1.  Let `script`{.variable} be a new [module
    script](#module-script){#creating-scripts:module-script-3} that this
    algorithm will subsequently initialize.

2.  Set `script`{.variable}\'s [settings
    object](#settings-object){#creating-scripts:settings-object-4} to
    `settings`{.variable}.

3.  Set `script`{.variable}\'s [base
    URL](#concept-script-base-url){#creating-scripts:concept-script-base-url-4}
    and [fetch
    options](#concept-script-script-fetch-options){#creating-scripts:concept-script-script-fetch-options-4}
    to null.

4.  Set `script`{.variable}\'s [parse
    error](#concept-script-parse-error){#creating-scripts:concept-script-parse-error-7}
    and [error to
    rethrow](#concept-script-error-to-rethrow){#creating-scripts:concept-script-error-to-rethrow-5}
    to null.

5.  Let `sheet`{.variable} be the result of running the steps to [create
    a constructed
    `CSSStyleSheet`](https://drafts.csswg.org/cssom/#create-a-constructed-cssstylesheet){#creating-scripts:create-a-constructed-cssstylesheet
    x-internal="create-a-constructed-cssstylesheet"} with an empty
    dictionary as the argument.

6.  Run the steps to [synchronously replace the rules of a
    `CSSStyleSheet`](https://drafts.csswg.org/cssom/#synchronously-replace-the-rules-of-a-cssstylesheet){#creating-scripts:synchronously-replace-the-rules-of-a-cssstylesheet
    x-internal="synchronously-replace-the-rules-of-a-cssstylesheet"} on
    `sheet`{.variable} given `source`{.variable}.

    If this throws an exception, catch it, and set
    `script`{.variable}\'s [parse
    error](#concept-script-parse-error){#creating-scripts:concept-script-parse-error-8}
    to that exception, and return `script`{.variable}.

    The steps to [synchronously replace the rules of a
    `CSSStyleSheet`](https://drafts.csswg.org/cssom/#synchronously-replace-the-rules-of-a-cssstylesheet){#creating-scripts:synchronously-replace-the-rules-of-a-cssstylesheet-2
    x-internal="synchronously-replace-the-rules-of-a-cssstylesheet"}
    will throw if `source`{.variable} contains any `@import` rules. This
    is by-design for now because there is not yet an agreement on how to
    handle these for CSS module scripts; therefore they are blocked
    altogether until a consensus is reached.

7.  Set `script`{.variable}\'s
    [record](#concept-script-record){#creating-scripts:concept-script-record-4}
    to the result of
    [CreateDefaultExportSyntheticModule](https://tc39.es/ecma262/#sec-create-default-export-synthetic-module){#creating-scripts:createdefaultexportsyntheticmodule
    x-internal="createdefaultexportsyntheticmodule"}(`sheet`{.variable}).

8.  Return `script`{.variable}.

To [create a JSON module script]{#creating-a-json-module-script .dfn},
given a string `source`{.variable} and an [environment settings
object](#environment-settings-object){#creating-scripts:environment-settings-object-5}
`settings`{.variable}:

1.  Let `script`{.variable} be a new [module
    script](#module-script){#creating-scripts:module-script-4} that this
    algorithm will subsequently initialize.

2.  Set `script`{.variable}\'s [settings
    object](#settings-object){#creating-scripts:settings-object-5} to
    `settings`{.variable}.

3.  Set `script`{.variable}\'s [base
    URL](#concept-script-base-url){#creating-scripts:concept-script-base-url-5}
    and [fetch
    options](#concept-script-script-fetch-options){#creating-scripts:concept-script-script-fetch-options-5}
    to null.

4.  Set `script`{.variable}\'s [parse
    error](#concept-script-parse-error){#creating-scripts:concept-script-parse-error-9}
    and [error to
    rethrow](#concept-script-error-to-rethrow){#creating-scripts:concept-script-error-to-rethrow-6}
    to null.

5.  Let `result`{.variable} be
    [ParseJSONModule](https://tc39.es/ecma262/#sec-parse-json-module){#creating-scripts:parsejsonmodule
    x-internal="parsejsonmodule"}(`source`{.variable}).

    If this throws an exception, catch it, and set
    `script`{.variable}\'s [parse
    error](#concept-script-parse-error){#creating-scripts:concept-script-parse-error-10}
    to that exception, and return `script`{.variable}.

6.  Set `script`{.variable}\'s
    [record](#concept-script-record){#creating-scripts:concept-script-record-5}
    to `result`{.variable}.

7.  Return `script`{.variable}.

The [module type from module request]{#module-type-from-module-request
.dfn} steps, given a [ModuleRequest
Record](https://tc39.es/ecma262/#modulerequest-record){#creating-scripts:modulerequest-record
x-internal="modulerequest-record"} `moduleRequest`{.variable}, are as
follows:

1.  Let `moduleType`{.variable} be \"`javascript-or-wasm`\".

2.  If `moduleRequest`{.variable}.\[\[Attributes\]\] has a
    [Record](https://tc39.es/ecma262/#sec-list-and-record-specification-type){#creating-scripts:record
    x-internal="record"} `entry`{.variable} such that
    `entry`{.variable}.\[\[Key\]\] is \"`type`\", then:

    1.  If `entry`{.variable}.\[\[Value\]\] is \"`javascript-or-wasm`\",
        then set `moduleType`{.variable} to null.

        This specification uses the \"`javascript-or-wasm`\" module type
        internally for [JavaScript module
        scripts](#javascript-module-script){#creating-scripts:javascript-module-script}
        or [WebAssembly module
        scripts](#webassembly-module-script){#creating-scripts:webassembly-module-script},
        so this step is needed to prevent modules from being imported
        using a \"`javascript-or-wasm`\" type attribute (a null
        `moduleType`{.variable} will cause the [module type
        allowed](#module-type-allowed){#creating-scripts:module-type-allowed}
        check to fail).

    2.  Otherwise, set `moduleType`{.variable} to
        `entry`{.variable}.\[\[Value\]\].

3.  Return `moduleType`{.variable}.

The [module type allowed]{#module-type-allowed .dfn} steps, given a
[string](https://infra.spec.whatwg.org/#string){#creating-scripts:string-3
x-internal="string"} `moduleType`{.variable} and an [environment
settings
object](#environment-settings-object){#creating-scripts:environment-settings-object-6}
`settings`{.variable}, are as follows:

1.  If `moduleType`{.variable} is not \"`javascript-or-wasm`\",
    \"`css`\", or \"`json`\", then return false.

2.  If `moduleType`{.variable} is \"`css`\" and the
    [`CSSStyleSheet`{#creating-scripts:cssstylesheet}](https://drafts.csswg.org/cssom/#the-cssstylesheet-interface){x-internal="cssstylesheet"}
    interface is not
    [exposed](https://webidl.spec.whatwg.org/#dfn-exposed){#creating-scripts:idl-exposed
    x-internal="idl-exposed"} in `settings`{.variable}\'s
    [realm](#environment-settings-object's-realm){#creating-scripts:environment-settings-object's-realm-4},
    then return false.

3.  Return true.

The [fetch destination from module
type]{#fetch-destination-from-module-type .dfn} steps, given a
[destination](https://fetch.spec.whatwg.org/#concept-request-destination){#creating-scripts:concept-request-destination
x-internal="concept-request-destination"}
`defaultDestination`{.variable} and a
[string](https://infra.spec.whatwg.org/#string){#creating-scripts:string-4
x-internal="string"} `moduleType`{.variable}, are as follows:

1.  If `moduleType`{.variable} is \"`json`\", then return \"`json`\".
2.  If `moduleType`{.variable} is \"`css`\", then return \"`style`\".
3.  Return `defaultDestination`{.variable}.

##### [8.1.4.4]{.secno} Calling scripts[](#calling-scripts){.self-link}

To [run a classic script]{#run-a-classic-script .dfn export=""} given a
[classic script](#classic-script){#calling-scripts:classic-script}
`script`{.variable} and an optional boolean `rethrow errors`{.variable}
(default false):

1.  Let `settings`{.variable} be the [settings
    object](#settings-object){#calling-scripts:settings-object} of
    `script`{.variable}.

2.  [Check if we can run
    script](#check-if-we-can-run-script){#calling-scripts:check-if-we-can-run-script}
    with `settings`{.variable}. If this returns \"do not run\", then
    return
    [NormalCompletion](https://tc39.es/ecma262/#sec-normalcompletion){#calling-scripts:normalcompletion
    x-internal="normalcompletion"}(empty).

3.  [Record classic script execution start
    time](https://w3c.github.io/long-animation-frames/#record-classic-script-execution-start-time){#calling-scripts:record-classic-script-execution-start-time
    x-internal="record-classic-script-execution-start-time"} given
    `script`{.variable}.

4.  [Prepare to run
    script](#prepare-to-run-script){#calling-scripts:prepare-to-run-script}
    given `settings`{.variable}.

5.  Let `evaluationStatus`{.variable} be null.

6.  If `script`{.variable}\'s [error to
    rethrow](#concept-script-error-to-rethrow){#calling-scripts:concept-script-error-to-rethrow}
    is not null, then set `evaluationStatus`{.variable} to Completion {
    \[\[Type\]\]: throw, \[\[Value\]\]: `script`{.variable}\'s [error to
    rethrow](#concept-script-error-to-rethrow){#calling-scripts:concept-script-error-to-rethrow-2},
    \[\[Target\]\]: empty }.

7.  Otherwise, set `evaluationStatus`{.variable} to
    [ScriptEvaluation](https://tc39.es/ecma262/#sec-runtime-semantics-scriptevaluation){#calling-scripts:js-scriptevaluation
    x-internal="js-scriptevaluation"}(`script`{.variable}\'s
    [record](#concept-script-record){#calling-scripts:concept-script-record}).

    If
    [ScriptEvaluation](https://tc39.es/ecma262/#sec-runtime-semantics-scriptevaluation){#calling-scripts:js-scriptevaluation-2
    x-internal="js-scriptevaluation"} does not complete because the user
    agent has [aborted the running
    script](#abort-a-running-script){#calling-scripts:abort-a-running-script},
    leave `evaluationStatus`{.variable} as null.

8.  If `evaluationStatus`{.variable} is an [abrupt
    completion](https://tc39.es/ecma262/#sec-completion-record-specification-type){#calling-scripts:completion-record
    x-internal="completion-record"}, then:

    1.  If `rethrow errors`{.variable} is true and
        `script`{.variable}\'s [muted
        errors](#muted-errors){#calling-scripts:muted-errors} is false,
        then:

        1.  [Clean up after running
            script](#clean-up-after-running-script){#calling-scripts:clean-up-after-running-script}
            with `settings`{.variable}.

        2.  Rethrow `evaluationStatus`{.variable}.\[\[Value\]\].

    2.  If `rethrow errors`{.variable} is true and
        `script`{.variable}\'s [muted
        errors](#muted-errors){#calling-scripts:muted-errors-2} is true,
        then:

        1.  [Clean up after running
            script](#clean-up-after-running-script){#calling-scripts:clean-up-after-running-script-2}
            with `settings`{.variable}.

        2.  Throw a
            [\"`NetworkError`\"](https://webidl.spec.whatwg.org/#networkerror){#calling-scripts:networkerror
            x-internal="networkerror"}
            [`DOMException`{#calling-scripts:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    3.  Otherwise, `rethrow errors`{.variable} is false. Perform the
        following steps:

        1.  [Report an
            exception](#report-an-exception){#calling-scripts:report-an-exception}
            given by `evaluationStatus`{.variable}.\[\[Value\]\] for
            `script`{.variable}\'s [settings
            object](#settings-object){#calling-scripts:settings-object-2}\'s
            [global
            object](#concept-settings-object-global){#calling-scripts:concept-settings-object-global}.

        2.  [Clean up after running
            script](#clean-up-after-running-script){#calling-scripts:clean-up-after-running-script-3}
            with `settings`{.variable}.

        3.  Return `evaluationStatus`{.variable}.

9.  [Clean up after running
    script](#clean-up-after-running-script){#calling-scripts:clean-up-after-running-script-4}
    with `settings`{.variable}.

10. If `evaluationStatus`{.variable} is a normal completion, then return
    `evaluationStatus`{.variable}.

11. If we\'ve reached this point, `evaluationStatus`{.variable} was left
    as null because the script was [aborted
    prematurely](#abort-a-running-script){#calling-scripts:abort-a-running-script-2}
    during evaluation. Return Completion { \[\[Type\]\]: throw,
    \[\[Value\]\]: a new
    [\"`QuotaExceededError`\"](https://webidl.spec.whatwg.org/#quotaexceedederror){#calling-scripts:quotaexceedederror
    x-internal="quotaexceedederror"}
    [`DOMException`{#calling-scripts:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"},
    \[\[Target\]\]: empty }.

To [run a module script]{#run-a-module-script .dfn export=""} given a
[module script](#module-script){#calling-scripts:module-script}
`script`{.variable} and an optional boolean
`preventErrorReporting`{.variable} (default false):

1.  Let `settings`{.variable} be the [settings
    object](#settings-object){#calling-scripts:settings-object-3} of
    `script`{.variable}.

2.  [Check if we can run
    script](#check-if-we-can-run-script){#calling-scripts:check-if-we-can-run-script-2}
    with `settings`{.variable}. If this returns \"do not run\", then
    return [a promise resolved
    with](https://webidl.spec.whatwg.org/#a-promise-resolved-with){#calling-scripts:a-promise-resolved-with
    x-internal="a-promise-resolved-with"} undefined.

3.  [Record module script execution start
    time](https://w3c.github.io/long-animation-frames/#record-module-script-execution-start-time){#calling-scripts:record-module-script-execution-start-time
    x-internal="record-module-script-execution-start-time"} given
    `script`{.variable}.

4.  [Prepare to run
    script](#prepare-to-run-script){#calling-scripts:prepare-to-run-script-2}
    given `settings`{.variable}.

5.  Let `evaluationPromise`{.variable} be null.

6.  If `script`{.variable}\'s [error to
    rethrow](#concept-script-error-to-rethrow){#calling-scripts:concept-script-error-to-rethrow-3}
    is not null, then set `evaluationPromise`{.variable} to [a promise
    rejected
    with](https://webidl.spec.whatwg.org/#a-promise-rejected-with){#calling-scripts:a-promise-rejected-with
    x-internal="a-promise-rejected-with"} `script`{.variable}\'s [error
    to
    rethrow](#concept-script-error-to-rethrow){#calling-scripts:concept-script-error-to-rethrow-4}.

7.  Otherwise:

    1.  Let `record`{.variable} be `script`{.variable}\'s
        [record](#concept-script-record){#calling-scripts:concept-script-record-2}.

    2.  Set `evaluationPromise`{.variable} to
        `record`{.variable}.[Evaluate](https://tc39.es/ecma262/#sec-moduleevaluation){#calling-scripts:js-evaluate
        x-internal="js-evaluate"}().

        This step will recursively evaluate all of the module\'s
        dependencies.

        If
        [Evaluate](https://tc39.es/ecma262/#sec-moduleevaluation){#calling-scripts:js-evaluate-2
        x-internal="js-evaluate"} fails to complete as a result of the
        user agent [aborting the running
        script](#abort-a-running-script){#calling-scripts:abort-a-running-script-3},
        then set `evaluationPromise`{.variable} to [a promise rejected
        with](https://webidl.spec.whatwg.org/#a-promise-rejected-with){#calling-scripts:a-promise-rejected-with-2
        x-internal="a-promise-rejected-with"} a new
        [\"`QuotaExceededError`\"](https://webidl.spec.whatwg.org/#quotaexceedederror){#calling-scripts:quotaexceedederror-2
        x-internal="quotaexceedederror"}
        [`DOMException`{#calling-scripts:domexception-3}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

8.  If `preventErrorReporting`{.variable} is false, then [upon
    rejection](https://webidl.spec.whatwg.org/#upon-rejection){#calling-scripts:upon-rejection
    x-internal="upon-rejection"} of `evaluationPromise`{.variable} with
    `reason`{.variable}, [report an
    exception](#report-an-exception){#calling-scripts:report-an-exception-2}
    given by `reason`{.variable} for `script`{.variable}\'s [settings
    object](#settings-object){#calling-scripts:settings-object-4}\'s
    [global
    object](#concept-settings-object-global){#calling-scripts:concept-settings-object-global-2}.

9.  [Clean up after running
    script](#clean-up-after-running-script){#calling-scripts:clean-up-after-running-script-5}
    with `settings`{.variable}.

10. Return `evaluationPromise`{.variable}.

The steps to [check if we can run script]{#check-if-we-can-run-script
.dfn} with an [environment settings
object](#environment-settings-object){#calling-scripts:environment-settings-object}
`settings`{.variable} are as follows. They return either \"run\" or \"do
not run\".

1.  If the [global
    object](#concept-settings-object-global){#calling-scripts:concept-settings-object-global-3}
    specified by `settings`{.variable} is a
    [`Window`{#calling-scripts:window}](nav-history-apis.html#window)
    object whose
    [`Document`{#calling-scripts:document}](dom.html#document) object is
    not [fully
    active](document-sequences.html#fully-active){#calling-scripts:fully-active},
    then return \"do not run\".

2.  If [scripting is
    disabled](#concept-environment-noscript){#calling-scripts:concept-environment-noscript}
    for `settings`{.variable}, then return \"do not run\".

3.  Return \"run\".

The steps to [prepare to run script]{#prepare-to-run-script .dfn
export=""} with an [environment settings
object](#environment-settings-object){#calling-scripts:environment-settings-object-2}
`settings`{.variable} are as follows:

1.  Push `settings`{.variable}\'s [realm execution
    context](#realm-execution-context){#calling-scripts:realm-execution-context}
    onto the [JavaScript execution context
    stack](https://tc39.es/ecma262/#execution-context-stack){#calling-scripts:javascript-execution-context-stack
    x-internal="javascript-execution-context-stack"}; it is now the
    [running JavaScript execution
    context](https://tc39.es/ecma262/#running-execution-context){#calling-scripts:running-javascript-execution-context
    x-internal="running-javascript-execution-context"}.

2.  Add `settings`{.variable} to the [surrounding
    agent](https://tc39.es/ecma262/#surrounding-agent){#calling-scripts:surrounding-agent
    x-internal="surrounding-agent"}\'s [event
    loop](#concept-agent-event-loop){#calling-scripts:concept-agent-event-loop}\'s
    [currently running
    task](#currently-running-task){#calling-scripts:currently-running-task}\'s
    [script evaluation environment settings object
    set](#script-evaluation-environment-settings-object-set){#calling-scripts:script-evaluation-environment-settings-object-set}.

The steps to [clean up after running
script]{#clean-up-after-running-script .dfn export=""} with an
[environment settings
object](#environment-settings-object){#calling-scripts:environment-settings-object-3}
`settings`{.variable} are as follows:

1.  [Assert](https://infra.spec.whatwg.org/#assert){#calling-scripts:assert
    x-internal="assert"}: `settings`{.variable}\'s [realm execution
    context](#realm-execution-context){#calling-scripts:realm-execution-context-2}
    is the [running JavaScript execution
    context](https://tc39.es/ecma262/#running-execution-context){#calling-scripts:running-javascript-execution-context-2
    x-internal="running-javascript-execution-context"}.

2.  Remove `settings`{.variable}\'s [realm execution
    context](#realm-execution-context){#calling-scripts:realm-execution-context-3}
    from the [JavaScript execution context
    stack](https://tc39.es/ecma262/#execution-context-stack){#calling-scripts:javascript-execution-context-stack-2
    x-internal="javascript-execution-context-stack"}.

3.  If the [JavaScript execution context
    stack](https://tc39.es/ecma262/#execution-context-stack){#calling-scripts:javascript-execution-context-stack-3
    x-internal="javascript-execution-context-stack"} is now empty,
    [perform a microtask
    checkpoint](#perform-a-microtask-checkpoint){#calling-scripts:perform-a-microtask-checkpoint}.
    (If this runs scripts, these algorithms will be invoked
    reentrantly.)

These algorithms are not invoked by one script directly calling another,
but they can be invoked reentrantly in an indirect manner, e.g. if a
script dispatches an event which has event listeners registered.

The [running script]{#running-script .dfn} is the
[script](#concept-script){#calling-scripts:concept-script} in the
\[\[HostDefined\]\] field in the ScriptOrModule component of the
[running JavaScript execution
context](https://tc39.es/ecma262/#running-execution-context){#calling-scripts:running-javascript-execution-context-3
x-internal="running-javascript-execution-context"}.

##### [8.1.4.5]{.secno} Killing scripts[](#killing-scripts){.self-link}

Although the JavaScript specification does not account for this
possibility, it\'s sometimes necessary to [abort a running
script]{#abort-a-running-script .dfn
lt="abort a running script|abort the running script" export=""}. This
causes any
[ScriptEvaluation](https://tc39.es/ecma262/#sec-runtime-semantics-scriptevaluation){#killing-scripts:js-scriptevaluation
x-internal="js-scriptevaluation"} or [Source Text Module
Record](https://tc39.es/ecma262/#sec-source-text-module-records){#killing-scripts:source-text-module-record
x-internal="source-text-module-record"}
[Evaluate](https://tc39.es/ecma262/#sec-moduleevaluation){#killing-scripts:js-evaluate
x-internal="js-evaluate"} invocations to cease immediately, emptying the
[JavaScript execution context
stack](https://tc39.es/ecma262/#execution-context-stack){#killing-scripts:javascript-execution-context-stack
x-internal="javascript-execution-context-stack"} without triggering any
of the normal mechanisms like `finally` blocks.
[\[JAVASCRIPT\]](references.html#refsJAVASCRIPT)

User agents may impose resource limitations on scripts, for example CPU
quotas, memory limits, total execution time limits, or bandwidth
limitations. When a script exceeds a limit, the user agent may either
throw a
[\"`QuotaExceededError`\"](https://webidl.spec.whatwg.org/#quotaexceedederror){#killing-scripts:quotaexceedederror
x-internal="quotaexceedederror"}
[`DOMException`{#killing-scripts:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"},
[abort the
script](#abort-a-running-script){#killing-scripts:abort-a-running-script}
without an exception, prompt the user, or throttle script execution.

::: example
For example, the following script never terminates. A user agent could,
after waiting for a few seconds, prompt the user to either terminate the
script or let it continue.

``` html
<script>
 while (true) { /* loop */ }
</script>
```
:::

User agents are encouraged to allow users to disable scripting whenever
the user is prompted either by a script (e.g. using the
[`window.alert()`{#killing-scripts:dom-alert}](timers-and-user-prompts.html#dom-alert)
API) or because of a script\'s actions (e.g. because it has exceeded a
time limit).

If scripting is disabled while a script is executing, the script should
be terminated immediately.

User agents may allow users to specifically disable scripts just for the
purposes of closing a [browsing
context](document-sequences.html#browsing-context){#killing-scripts:browsing-context}.

For example, the prompt mentioned in the example above could also offer
the user with a mechanism to just close the page entirely, without
running any
[`unload`{#killing-scripts:event-unload}](indices.html#event-unload)
event handlers.

##### [8.1.4.6]{.secno} Runtime script errors[](#runtime-script-errors){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[reportError](https://developer.mozilla.org/en-US/docs/Web/API/reportError "The reportError() global method may be used to report errors to the console or global event handlers, emulating an uncaught JavaScript exception.")

Support in all current engines.

::: support
[Firefox93+]{.firefox .yes}[Safari15.4+]{.safari
.yes}[Chrome95+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge95+]{.edge_blink .yes}

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

`self.`[`reportError`](#dom-reporterror){#dom-reporterror-dev}`(``e`{.variable}`)`

:   Dispatches an
    [`error`{#runtime-script-errors:event-error}](indices.html#event-error)
    event at the global object for the given value `e`{.variable}, in
    the same fashion as an unhandled exception.

To [extract error information]{#extract-error .dfn} from a JavaScript
value `exception`{.variable}:

1.  Let `attributes`{.variable} be an empty
    [map](https://infra.spec.whatwg.org/#ordered-map){#runtime-script-errors:ordered-map
    x-internal="ordered-map"} keyed by IDL attributes.

2.  Set
    `attributes`{.variable}\[[`error`{#runtime-script-errors:dom-errorevent-error}](#dom-errorevent-error)\]
    to `exception`{.variable}.

3.  Set
    `attributes`{.variable}\[[`message`{#runtime-script-errors:dom-errorevent-message}](#dom-errorevent-message)\],
    `attributes`{.variable}\[[`filename`{#runtime-script-errors:dom-errorevent-filename}](#dom-errorevent-filename)\],
    `attributes`{.variable}\[[`lineno`{#runtime-script-errors:dom-errorevent-lineno}](#dom-errorevent-lineno)\],
    and
    `attributes`{.variable}\[[`colno`{#runtime-script-errors:dom-errorevent-colno}](#dom-errorevent-colno)\]
    to
    [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#runtime-script-errors:implementation-defined
    x-internal="implementation-defined"} values derived from
    `exception`{.variable}.

    Browsers implement behavior not specified here or in the JavaScript
    specification to gather values which are helpful, including in
    unusual cases (e.g., `eval`). In the future, this might be specified
    in greater detail.

4.  Return `attributes`{.variable}.

[]{#report-the-error}To [report an exception]{#report-an-exception .dfn
export=""} `exception`{.variable} which is a JavaScript value, for a
particular [global
object](#global-object){#runtime-script-errors:global-object}
`global`{.variable} and optional boolean
[`omitError`{.variable}]{#report-exception-omiterror .dfn} (default
false):

1.  Let `notHandled`{.variable} be true.

2.  Let `errorInfo`{.variable} be the result of [extracting error
    information](#extract-error){#runtime-script-errors:extract-error}
    from `exception`{.variable}.

3.  Let `script`{.variable} be a
    [script](#concept-script){#runtime-script-errors:concept-script}
    found in an
    [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#runtime-script-errors:implementation-defined-2
    x-internal="implementation-defined"} way, or null. This should
    usually be the [running
    script](#running-script){#runtime-script-errors:running-script}
    (most notably during [run a classic
    script](#run-a-classic-script){#runtime-script-errors:run-a-classic-script}).

    Implementations have not yet settled on interoperable behavior for
    which script is used to determine whether errors are muted in less
    common cases.

4.  If `script`{.variable} is a [classic
    script](#classic-script){#runtime-script-errors:classic-script} and
    `script`{.variable}\'s [muted
    errors](#muted-errors){#runtime-script-errors:muted-errors} is true,
    then set
    `errorInfo`{.variable}\[[`error`{#runtime-script-errors:dom-errorevent-error-2}](#dom-errorevent-error)\]
    to null,
    `errorInfo`{.variable}\[[`message`{#runtime-script-errors:dom-errorevent-message-2}](#dom-errorevent-message)\]
    to \"`Script error.`\",
    `errorInfo`{.variable}\[[`filename`{#runtime-script-errors:dom-errorevent-filename-2}](#dom-errorevent-filename)\]
    to the empty string,
    `errorInfo`{.variable}\[[`lineno`{#runtime-script-errors:dom-errorevent-lineno-2}](#dom-errorevent-lineno)\]
    to 0, and
    `errorInfo`{.variable}\[[`colno`{#runtime-script-errors:dom-errorevent-colno-2}](#dom-errorevent-colno)\]
    to 0.

5.  If `omitError`{.variable} is true, then set
    `errorInfo`{.variable}\[[`error`{#runtime-script-errors:dom-errorevent-error-3}](#dom-errorevent-error)\]
    to null.

6.  If `global`{.variable} is not [in error reporting
    mode](#in-error-reporting-mode){#runtime-script-errors:in-error-reporting-mode},
    then:

    1.  Set `global`{.variable}\'s [in error reporting
        mode](#in-error-reporting-mode){#runtime-script-errors:in-error-reporting-mode-2}
        to true.

    2.  If `global`{.variable} implements
        [`EventTarget`{#runtime-script-errors:eventtarget}](https://dom.spec.whatwg.org/#interface-eventtarget){x-internal="eventtarget"},
        then set `notHandled`{.variable} to the result of [firing an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#runtime-script-errors:concept-event-fire
        x-internal="concept-event-fire"} named
        [`error`{#runtime-script-errors:event-error-2}](indices.html#event-error)
        at `global`{.variable}, using
        [`ErrorEvent`{#runtime-script-errors:errorevent}](#errorevent),
        with the
        [`cancelable`{#runtime-script-errors:dom-event-cancelable}](https://dom.spec.whatwg.org/#dom-event-cancelable){x-internal="dom-event-cancelable"}
        attribute initialized to true, and additional attributes
        initialized according to `errorInfo`{.variable}.

        Returning true in an event handler cancels the event per [the
        event handler processing
        algorithm](#the-event-handler-processing-algorithm){#runtime-script-errors:the-event-handler-processing-algorithm}.

    3.  Set `global`{.variable}\'s [in error reporting
        mode](#in-error-reporting-mode){#runtime-script-errors:in-error-reporting-mode-3}
        to false.

7.  If `notHandled`{.variable} is true, then:

    1.  Set
        `errorInfo`{.variable}\[[`error`{#runtime-script-errors:dom-errorevent-error-4}](#dom-errorevent-error)\]
        to null.

    2.  If `global`{.variable} implements
        [`DedicatedWorkerGlobalScope`{#runtime-script-errors:dedicatedworkerglobalscope}](workers.html#dedicatedworkerglobalscope),
        [queue a global
        task](#queue-a-global-task){#runtime-script-errors:queue-a-global-task}
        on the [DOM manipulation task
        source](#dom-manipulation-task-source){#runtime-script-errors:dom-manipulation-task-source}
        with the `global`{.variable}\'s associated
        [`Worker`{#runtime-script-errors:worker}](workers.html#worker)\'s
        [relevant global
        object](#concept-relevant-global){#runtime-script-errors:concept-relevant-global}
        to run these steps:

        1.  Let `workerObject`{.variable} be the
            [`Worker`{#runtime-script-errors:worker-2}](workers.html#worker)
            object associated with `global`{.variable}.

        2.  Set `notHandled`{.variable} to the result of [firing an
            event](https://dom.spec.whatwg.org/#concept-event-fire){#runtime-script-errors:concept-event-fire-2
            x-internal="concept-event-fire"} named
            [`error`{#runtime-script-errors:event-error-3}](indices.html#event-error)
            at `workerObject`{.variable}, using
            [`ErrorEvent`{#runtime-script-errors:errorevent-2}](#errorevent),
            with the
            [`cancelable`{#runtime-script-errors:dom-event-cancelable-2}](https://dom.spec.whatwg.org/#dom-event-cancelable){x-internal="dom-event-cancelable"}
            attribute initialized to true, and additional attributes
            initialized according to `errorInfo`{.variable}.

        3.  If `notHandled`{.variable} is true, then
            [report](#report-an-exception){#runtime-script-errors:report-an-exception}
            `exception`{.variable} for `workerObject`{.variable}\'s
            [relevant global
            object](#concept-relevant-global){#runtime-script-errors:concept-relevant-global-2}
            with
            [`omitError`{.variable}](#report-exception-omiterror){#runtime-script-errors:report-exception-omiterror}
            set to true.

            The actual `exception`{.variable} value will not be
            available in the owner realm, but the user agent still
            carries through enough information to set the message,
            filename, and other attributes, as well as potentially
            report to a developer console.

    3.  Otherwise, the user agent may report `exception`{.variable} to a
        developer console.

If the implicit port connecting a worker to its
[`Worker`{#runtime-script-errors:worker-3}](workers.html#worker) object
has been disentangled (i.e. if the parent worker has been terminated),
then the user agent must act as if the
[`Worker`{#runtime-script-errors:worker-4}](workers.html#worker) object
had no
[`error`{#runtime-script-errors:event-error-4}](indices.html#event-error)
event handler and as if that worker\'s
[`onerror`{#runtime-script-errors:handler-workerglobalscope-onerror}](workers.html#handler-workerglobalscope-onerror)
attribute was null, but must otherwise act as described above.

Thus, error reports propagate up to the chain of dedicated workers up to
the original
[`Document`{#runtime-script-errors:document}](dom.html#document), even
if some of the workers along this chain have been terminated and garbage
collected.

Previous revisions of this standard defined an algorithm to [report the
exception]{#report-the-exception .dfn lt="report the exception"
export=""}. As part of [issue #958](https://github.com/whatwg/html),
this has been superseded by [report an
exception](#report-an-exception){#runtime-script-errors:report-an-exception-2}
which behaves differently and takes different inputs. [Issue
#10516](https://github.com/whatwg/html/issues/10516) tracks updating the
specification ecosystem.

------------------------------------------------------------------------

The [`reportError(``e`{.variable}`)`]{#dom-reporterror .dfn
dfn-for="WindowOrWorkerGlobalScope" dfn-type="method"} method steps are
to [report an
exception](#report-an-exception){#runtime-script-errors:report-an-exception-3}
`e`{.variable} for
[this](https://webidl.spec.whatwg.org/#this){#runtime-script-errors:this
x-internal="this"}.

It is unclear whether
[muting](#muted-errors){#runtime-script-errors:muted-errors-2} is
applicable here. In Chrome and Safari it is muted, but in Firefox it is
not. See also [issue #958](https://github.com/whatwg/html/issues/958).

------------------------------------------------------------------------

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[ErrorEvent](https://developer.mozilla.org/en-US/docs/Web/API/ErrorEvent "The ErrorEvent interface represents events providing information related to errors in scripts or in files.")

Support in all current engines.

::: support
[Firefox27+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome10+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

The [`ErrorEvent`{#runtime-script-errors:errorevent-3}](#errorevent)
interface is defined as follows:

``` idl
[Exposed=*]
interface ErrorEvent : Event {
  constructor(DOMString type, optional ErrorEventInit eventInitDict = {});

  readonly attribute DOMString message;
  readonly attribute USVString filename;
  readonly attribute unsigned long lineno;
  readonly attribute unsigned long colno;
  readonly attribute any error;
};

dictionary ErrorEventInit : EventInit {
  DOMString message = "";
  USVString filename = "";
  unsigned long lineno = 0;
  unsigned long colno = 0;
  any error;
};
```

The [`message`]{#dom-errorevent-message .dfn dfn-for="ErrorEvent"
dfn-type="attribute"} attribute must return the value it was initialized
to. It represents the error message.

The [`filename`]{#dom-errorevent-filename .dfn dfn-for="ErrorEvent"
dfn-type="attribute"} attribute must return the value it was initialized
to. It represents the
[URL](https://url.spec.whatwg.org/#concept-url){#runtime-script-errors:url
x-internal="url"} of the script in which the error originally occurred.

The [`lineno`]{#dom-errorevent-lineno .dfn dfn-for="ErrorEvent"
dfn-type="attribute"} attribute must return the value it was initialized
to. It represents the line number where the error occurred in the
script.

The [`colno`]{#dom-errorevent-colno .dfn dfn-for="ErrorEvent"
dfn-type="attribute"} attribute must return the value it was initialized
to. It represents the column number where the error occurred in the
script.

The [`error`]{#dom-errorevent-error .dfn dfn-for="ErrorEvent"
dfn-type="attribute"} attribute must return the value it was initialized
to. It must initially be initialized to undefined. Where appropriate, it
is set to the object representing the error (e.g., the exception object
in the case of an uncaught exception).

##### [8.1.4.7]{.secno} Unhandled promise rejections[](#unhandled-promise-rejections){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Window/rejectionhandled_event](https://developer.mozilla.org/en-US/docs/Web/API/Window/rejectionhandled_event "The rejectionhandled event is sent to the script's global scope (usually window but also Worker) whenever a rejected JavaScript Promise is handled late, i.e. when a handler is attached to the promise after its rejection had caused an unhandledrejection event.")

Support in all current engines.

::: support
[Firefox69+]{.firefox .yes}[Safari11+]{.safari .yes}[Chrome49+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari
iOS11.3+]{.safari_ios .yes}[Chrome Android?]{.chrome_android
.unknown}[WebView Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

In addition to synchronous [runtime script
errors](#runtime-script-errors), scripts may experience asynchronous
promise rejections, tracked via the
[`unhandledrejection`{#unhandled-promise-rejections:event-unhandledrejection}](indices.html#event-unhandledrejection)
and
[`rejectionhandled`{#unhandled-promise-rejections:event-rejectionhandled}](indices.html#event-rejectionhandled)
events. Tracking these rejections is done via the
[HostPromiseRejectionTracker](#the-hostpromiserejectiontracker-implementation){#unhandled-promise-rejections:the-hostpromiserejectiontracker-implementation}
abstract operation, but reporting them is defined here.

To [notify about rejected promises]{#notify-about-rejected-promises
.dfn} given a [global
object](#global-object){#unhandled-promise-rejections:global-object}
`global`{.variable}:

1.  Let `list`{.variable} be a
    [clone](https://infra.spec.whatwg.org/#list-clone){#unhandled-promise-rejections:list-clone
    x-internal="list-clone"} of `global`{.variable}\'s
    [about-to-be-notified rejected promises
    list](#about-to-be-notified-rejected-promises-list){#unhandled-promise-rejections:about-to-be-notified-rejected-promises-list}.

2.  If `list`{.variable} [is
    empty](https://infra.spec.whatwg.org/#list-is-empty){#unhandled-promise-rejections:list-is-empty
    x-internal="list-is-empty"}, then return.

3.  [Empty](https://infra.spec.whatwg.org/#list-empty){#unhandled-promise-rejections:list-empty
    x-internal="list-empty"} `global`{.variable}\'s
    [about-to-be-notified rejected promises
    list](#about-to-be-notified-rejected-promises-list){#unhandled-promise-rejections:about-to-be-notified-rejected-promises-list-2}.

4.  [Queue a global
    task](#queue-a-global-task){#unhandled-promise-rejections:queue-a-global-task}
    on the [DOM manipulation task
    source](#dom-manipulation-task-source){#unhandled-promise-rejections:dom-manipulation-task-source}
    given `global`{.variable} to run the following step:

    1.  [For
        each](https://infra.spec.whatwg.org/#list-iterate){#unhandled-promise-rejections:list-iterate
        x-internal="list-iterate"} promise `p`{.variable} of
        `list`{.variable}:

        1.  If `p`{.variable}.\[\[PromiseIsHandled\]\] is true, then
            [continue](https://infra.spec.whatwg.org/#iteration-continue){#unhandled-promise-rejections:continue
            x-internal="continue"}.

        2.  Let `notCanceled`{.variable} be the result of [firing an
            event](https://dom.spec.whatwg.org/#concept-event-fire){#unhandled-promise-rejections:concept-event-fire
            x-internal="concept-event-fire"} named
            [`unhandledrejection`{#unhandled-promise-rejections:event-unhandledrejection-2}](indices.html#event-unhandledrejection)
            at `global`{.variable}, using
            [`PromiseRejectionEvent`{#unhandled-promise-rejections:promiserejectionevent}](#promiserejectionevent),
            with the
            [`cancelable`{#unhandled-promise-rejections:dom-event-cancelable}](https://dom.spec.whatwg.org/#dom-event-cancelable){x-internal="dom-event-cancelable"}
            attribute initialized to true, the
            [`promise`{#unhandled-promise-rejections:dom-promiserejectionevent-promise}](#dom-promiserejectionevent-promise)
            attribute initialized to `p`{.variable}, and the
            [`reason`{#unhandled-promise-rejections:dom-promiserejectionevent-reason}](#dom-promiserejectionevent-reason)
            attribute initialized to
            `p`{.variable}.\[\[PromiseResult\]\].

        3.  [If `notCanceled`{.variable} is true, then the user agent
            may report `p`{.variable}.\[\[PromiseResult\]\] to a
            developer console.]{#concept-promise-rejection-handled}

        4.  If `p`{.variable}.\[\[PromiseIsHandled\]\] is false, then
            [append](https://infra.spec.whatwg.org/#set-append){#unhandled-promise-rejections:set-append
            x-internal="set-append"} `p`{.variable} to
            `global`{.variable}\'s [outstanding rejected promises weak
            set](#outstanding-rejected-promises-weak-set){#unhandled-promise-rejections:outstanding-rejected-promises-weak-set}.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[PromiseRejectionEvent](https://developer.mozilla.org/en-US/docs/Web/API/PromiseRejectionEvent "The PromiseRejectionEvent interface represents events which are sent to the global script context when JavaScript Promises are rejected. These events are particularly useful for telemetry and debugging purposes.")

Support in all current engines.

::: support
[Firefox69+]{.firefox .yes}[Safari11+]{.safari .yes}[Chrome49+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari
iOS11.3+]{.safari_ios .yes}[Chrome Android?]{.chrome_android
.unknown}[WebView Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

The
[`PromiseRejectionEvent`{#unhandled-promise-rejections:promiserejectionevent-2}](#promiserejectionevent)
interface is defined as follows:

``` idl
[Exposed=*]
interface PromiseRejectionEvent : Event {
  constructor(DOMString type, PromiseRejectionEventInit eventInitDict);

  readonly attribute object promise;
  readonly attribute any reason;
};

dictionary PromiseRejectionEventInit : EventInit {
  required object promise;
  any reason;
};
```

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[PromiseRejectionEvent/promise](https://developer.mozilla.org/en-US/docs/Web/API/PromiseRejectionEvent/promise "The PromiseRejectionEvent interface's promise read-only property indicates the JavaScript Promise which was rejected. You can examine the event's PromiseRejectionEvent.reason property to learn why the promise was rejected.")

Support in all current engines.

::: support
[Firefox69+]{.firefox .yes}[Safari11+]{.safari .yes}[Chrome49+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari
iOS11.3+]{.safari_ios .yes}[Chrome Android?]{.chrome_android
.unknown}[WebView Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

The [`promise`]{#dom-promiserejectionevent-promise .dfn
dfn-for="PromiseRejectionEvent" dfn-type="attribute"} attribute must
return the value it was initialized to. It represents the promise which
this notification is about.

Because of how Web IDL conversion rules for
[`Promise`](https://webidl.spec.whatwg.org/#idl-promise){#unhandled-promise-rejections:idl-promise
x-internal="idl-promise"}`<``T`{.variable}`>` types always wrap the
input into a new promise, the
[`promise`{#unhandled-promise-rejections:dom-promiserejectionevent-promise-3}](#dom-promiserejectionevent-promise)
attribute is of type
[`object`{#unhandled-promise-rejections:idl-object-3}](https://webidl.spec.whatwg.org/#idl-object){x-internal="idl-object"}
instead, which is more appropriate for representing an opaque handle to
the original promise object.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[PromiseRejectionEvent/reason](https://developer.mozilla.org/en-US/docs/Web/API/PromiseRejectionEvent/reason "The PromiseRejectionEvent reason read-only property is any JavaScript value or Object which provides the reason passed into Promise.reject(). This in theory provides information about why the promise was rejected.")

Support in all current engines.

::: support
[Firefox69+]{.firefox .yes}[Safari11+]{.safari .yes}[Chrome49+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari
iOS11.3+]{.safari_ios .yes}[Chrome Android?]{.chrome_android
.unknown}[WebView Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

The [`reason`]{#dom-promiserejectionevent-reason .dfn
dfn-for="PromiseRejectionEvent" dfn-type="attribute"} attribute must
return the value it was initialized to. It represents the rejection
reason for the promise.

##### [8.1.4.8]{.secno} Import map parse results[](#import-map-parse-results){.self-link}

An [import map parse result]{#import-map-parse-result .dfn} is a
[struct](https://infra.spec.whatwg.org/#struct){#import-map-parse-results:struct
x-internal="struct"} that is similar to a
[script](#concept-script){#import-map-parse-results:concept-script}, and
also can be stored in a
[`script`{#import-map-parse-results:the-script-element}](scripting.html#the-script-element)
element\'s
[result](scripting.html#concept-script-result){#import-map-parse-results:concept-script-result},
but is not counted as a
[script](#concept-script){#import-map-parse-results:concept-script-2}
for other purposes. It has the following
[items](https://infra.spec.whatwg.org/#struct-item){#import-map-parse-results:struct-item
x-internal="struct-item"}:

An [import map]{#impr-import-map .dfn}
:   An [import map](#import-map){#import-map-parse-results:import-map}
    or null.

An [error to rethrow]{#impr-error-to-rethrow .dfn}
:   A JavaScript value representing an error that will prevent using
    this import map, when non-null.

To [create an import map parse
result]{#create-an-import-map-parse-result .dfn} given a
[string](https://infra.spec.whatwg.org/#string){#import-map-parse-results:string
x-internal="string"} `input`{.variable} and a
[URL](https://url.spec.whatwg.org/#concept-url){#import-map-parse-results:url
x-internal="url"} `baseURL`{.variable}:

1.  Let `result`{.variable} be an [import map parse
    result](#import-map-parse-result){#import-map-parse-results:import-map-parse-result}
    whose [import
    map](#impr-import-map){#import-map-parse-results:impr-import-map} is
    null and whose [error to
    rethrow](#impr-error-to-rethrow){#import-map-parse-results:impr-error-to-rethrow}
    is null.

2.  [Parse an import map
    string](#parse-an-import-map-string){#import-map-parse-results:parse-an-import-map-string}
    given `input`{.variable} and `baseURL`{.variable}, catching any
    exceptions. If this threw an exception, then set
    `result`{.variable}\'s [error to
    rethrow](#impr-error-to-rethrow){#import-map-parse-results:impr-error-to-rethrow-2}
    to that exception. Otherwise, set `result`{.variable}\'s [import
    map](#impr-import-map){#import-map-parse-results:impr-import-map-2}
    to the return value.

3.  Return `result`{.variable}.

To [register an import map]{#register-an-import-map .dfn} given a
[`Window`{#import-map-parse-results:window}](nav-history-apis.html#window)
`global`{.variable} and an [import map parse
result](#import-map-parse-result){#import-map-parse-results:import-map-parse-result-2}
`result`{.variable}:

1.  If `result`{.variable}\'s [error to
    rethrow](#impr-error-to-rethrow){#import-map-parse-results:impr-error-to-rethrow-3}
    is not null, then [report an
    exception](#report-an-exception){#import-map-parse-results:report-an-exception}
    given by `result`{.variable}\'s [error to
    rethrow](#impr-error-to-rethrow){#import-map-parse-results:impr-error-to-rethrow-4}
    for `global`{.variable} and return.

2.  [Merge existing and new import
    maps](#merge-existing-and-new-import-maps){#import-map-parse-results:merge-existing-and-new-import-maps},
    given `global`{.variable} and `result`{.variable}\'s [import
    map](#impr-import-map){#import-map-parse-results:impr-import-map-3}.

#### [8.1.5]{.secno} Module specifier resolution[](#module-specifier-resolution){.self-link}

##### [8.1.5.1]{.secno} The resolution algorithm[](#the-resolution-algorithm){.self-link}

The [resolve a module
specifier](#resolve-a-module-specifier){#the-resolution-algorithm:resolve-a-module-specifier}
algorithm is the primary entry point for converting module specifier
strings into
[URLs](https://url.spec.whatwg.org/#concept-url){#the-resolution-algorithm:url
x-internal="url"}. When no [import
maps](#import-map){#the-resolution-algorithm:import-map} are involved,
it is relatively straightforward, and reduces to [resolving a URL-like
module
specifier](#resolving-a-url-like-module-specifier){#the-resolution-algorithm:resolving-a-url-like-module-specifier}.

When there is a non-empty [import
map](#import-map){#the-resolution-algorithm:import-map-2} present, the
behavior is more complex. It checks candidate entries from all
applicable [module specifier
maps](#module-specifier-map){#the-resolution-algorithm:module-specifier-map},
from most-specific to least-specific
[scopes](#concept-import-map-scopes){#the-resolution-algorithm:concept-import-map-scopes}
(falling back to the top-level unscoped
[imports](#concept-import-map-imports){#the-resolution-algorithm:concept-import-map-imports}),
and from most-specific to least-specific prefixes. For each candidate,
the [resolve an imports
match](#resolving-an-imports-match){#the-resolution-algorithm:resolving-an-imports-match}
algorithm will give on the following results:

- Successful resolution of the specifier to a
  [URL](https://url.spec.whatwg.org/#concept-url){#the-resolution-algorithm:url-2
  x-internal="url"}. Then the [resolve a module
  specifier](#resolve-a-module-specifier){#the-resolution-algorithm:resolve-a-module-specifier-2}
  algorithm will return that URL.

- Throwing an exception. Then the [resolve a module
  specifier](#resolve-a-module-specifier){#the-resolution-algorithm:resolve-a-module-specifier-3}
  algorithm will rethrow that exception, without any further fallbacks.

- Failing to resolve, without an error. In this case the outer [resolve
  a module
  specifier](#resolve-a-module-specifier){#the-resolution-algorithm:resolve-a-module-specifier-4}
  algorithm will move on to the next candidate.

In the end, if no successful resolution is found via any of the
candidate [module specifier
maps](#module-specifier-map){#the-resolution-algorithm:module-specifier-map-2},
[resolve a module
specifier](#resolve-a-module-specifier){#the-resolution-algorithm:resolve-a-module-specifier-5}
will throw an exception. Thus the result is always either a
[URL](https://url.spec.whatwg.org/#concept-url){#the-resolution-algorithm:url-3
x-internal="url"} or a thrown exception.

To [resolve a module specifier]{#resolve-a-module-specifier .dfn} given
a
[script](scripting.html#the-script-element){#the-resolution-algorithm:the-script-element}-or-null
`referringScript`{.variable} and a
[string](https://infra.spec.whatwg.org/#string){#the-resolution-algorithm:string
x-internal="string"} `specifier`{.variable}:

1.  Let `settingsObject`{.variable} and `baseURL`{.variable} be null.

2.  If `referringScript`{.variable} is not null, then:

    1.  Set `settingsObject`{.variable} to
        `referringScript`{.variable}\'s [settings
        object](#settings-object){#the-resolution-algorithm:settings-object}.

    2.  Set `baseURL`{.variable} to `referringScript`{.variable}\'s
        [base
        URL](#concept-script-base-url){#the-resolution-algorithm:concept-script-base-url}.

3.  Otherwise:

    1.  [Assert](https://infra.spec.whatwg.org/#assert){#the-resolution-algorithm:assert
        x-internal="assert"}: there is a [current settings
        object](#current-settings-object){#the-resolution-algorithm:current-settings-object}.

    2.  Set `settingsObject`{.variable} to the [current settings
        object](#current-settings-object){#the-resolution-algorithm:current-settings-object-2}.

    3.  Set `baseURL`{.variable} to `settingsObject`{.variable}\'s [API
        base
        URL](#api-base-url){#the-resolution-algorithm:api-base-url}.

4.  Let `importMap`{.variable} be an [empty import
    map](#empty-import-map){#the-resolution-algorithm:empty-import-map}.

5.  If `settingsObject`{.variable}\'s [global
    object](#concept-settings-object-global){#the-resolution-algorithm:concept-settings-object-global}
    implements
    [`Window`{#the-resolution-algorithm:window}](nav-history-apis.html#window),
    then set `importMap`{.variable} to `settingsObject`{.variable}\'s
    [global
    object](#concept-settings-object-global){#the-resolution-algorithm:concept-settings-object-global-2}\'s
    [import
    map](#concept-global-import-map){#the-resolution-algorithm:concept-global-import-map}.

6.  Let `serializedBaseURL`{.variable} be `baseURL`{.variable},
    [serialized](https://url.spec.whatwg.org/#concept-url-serializer){#the-resolution-algorithm:concept-url-serializer
    x-internal="concept-url-serializer"}.

7.  Let `asURL`{.variable} be the result of [resolving a URL-like module
    specifier](#resolving-a-url-like-module-specifier){#the-resolution-algorithm:resolving-a-url-like-module-specifier-2}
    given `specifier`{.variable} and `baseURL`{.variable}.

8.  Let `normalizedSpecifier`{.variable} be the
    [serialization](https://url.spec.whatwg.org/#concept-url-serializer){#the-resolution-algorithm:concept-url-serializer-2
    x-internal="concept-url-serializer"} of `asURL`{.variable}, if
    `asURL`{.variable} is non-null; otherwise, `specifier`{.variable}.

9.  Let `result`{.variable} be a
    [URL](https://url.spec.whatwg.org/#concept-url){#the-resolution-algorithm:url-4
    x-internal="url"}-or-null, initially null.

10. [For
    each](https://infra.spec.whatwg.org/#list-iterate){#the-resolution-algorithm:list-iterate
    x-internal="list-iterate"} `scopePrefix`{.variable} →
    `scopeImports`{.variable} of `importMap`{.variable}\'s
    [scopes](#concept-import-map-scopes){#the-resolution-algorithm:concept-import-map-scopes-2}:

    1.  If `scopePrefix`{.variable} is `serializedBaseURL`{.variable},
        or if `scopePrefix`{.variable} ends with U+002F (/) and
        `scopePrefix`{.variable} is a [code unit
        prefix](https://infra.spec.whatwg.org/#code-unit-prefix){#the-resolution-algorithm:code-unit-prefix
        x-internal="code-unit-prefix"} of
        `serializedBaseURL`{.variable}, then:

        1.  Let `scopeImportsMatch`{.variable} be the result of
            [resolving an imports
            match](#resolving-an-imports-match){#the-resolution-algorithm:resolving-an-imports-match-2}
            given `normalizedSpecifier`{.variable}, `asURL`{.variable},
            and `scopeImports`{.variable}.

        2.  If `scopeImportsMatch`{.variable} is not null, then set
            `result`{.variable} to `scopeImportsMatch`{.variable}, and
            [break](https://infra.spec.whatwg.org/#iteration-break){#the-resolution-algorithm:break
            x-internal="break"}.

11. If `result`{.variable} is null, set `result`{.variable} to the
    result of [resolving an imports
    match](#resolving-an-imports-match){#the-resolution-algorithm:resolving-an-imports-match-3}
    given `normalizedSpecifier`{.variable}, `asURL`{.variable}, and
    `importMap`{.variable}\'s
    [imports](#concept-import-map-imports){#the-resolution-algorithm:concept-import-map-imports-2}.

12. If `result`{.variable} is null, set it to `asURL`{.variable}.

    By this point, if `result`{.variable} was null,
    `specifier`{.variable} wasn\'t remapped to anything by
    `importMap`{.variable}, but it might have been able to be turned
    into a URL.

13. If `result`{.variable} is not null, then:

    1.  [Add module to resolved module
        set](#add-module-to-resolved-module-set){#the-resolution-algorithm:add-module-to-resolved-module-set}
        given `settingsObject`{.variable},
        `serializedBaseURL`{.variable},
        `normalizedSpecifier`{.variable}, and `asURL`{.variable}.

    2.  Return `result`{.variable}.

14. Throw a
    [`TypeError`{#the-resolution-algorithm:typeerror}](https://tc39.es/ecma262/#sec-native-error-types-used-in-this-standard-typeerror){x-internal="typeerror"}
    indicating that `specifier`{.variable} was a bare specifier, but was
    not remapped to anything by `importMap`{.variable}.

To [resolve an imports match]{#resolving-an-imports-match .dfn}, given a
[string](https://infra.spec.whatwg.org/#string){#the-resolution-algorithm:string-2
x-internal="string"} `normalizedSpecifier`{.variable}, a
[URL](https://url.spec.whatwg.org/#concept-url){#the-resolution-algorithm:url-5
x-internal="url"}-or-null `asURL`{.variable}, and a [module specifier
map](#module-specifier-map){#the-resolution-algorithm:module-specifier-map-3}
`specifierMap`{.variable}:

1.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#the-resolution-algorithm:list-iterate-2
    x-internal="list-iterate"} `specifierKey`{.variable} →
    `resolutionResult`{.variable} of `specifierMap`{.variable}:

    1.  If `specifierKey`{.variable} is
        `normalizedSpecifier`{.variable}, then:

        1.  If `resolutionResult`{.variable} is null, then throw a
            [`TypeError`{#the-resolution-algorithm:typeerror-2}](https://tc39.es/ecma262/#sec-native-error-types-used-in-this-standard-typeerror){x-internal="typeerror"}
            indicating that resolution of `specifierKey`{.variable} was
            blocked by a null entry.

            This will terminate the entire [resolve a module
            specifier](#resolve-a-module-specifier){#the-resolution-algorithm:resolve-a-module-specifier-6}
            algorithm, without any further fallbacks.

        2.  [Assert](https://infra.spec.whatwg.org/#assert){#the-resolution-algorithm:assert-2
            x-internal="assert"}: `resolutionResult`{.variable} is a
            [URL](https://url.spec.whatwg.org/#concept-url){#the-resolution-algorithm:url-6
            x-internal="url"}.

        3.  Return `resolutionResult`{.variable}.

    2.  If all of the following are true:

        - `specifierKey`{.variable} ends with U+002F (/);

        - `specifierKey`{.variable} is a [code unit
          prefix](https://infra.spec.whatwg.org/#code-unit-prefix){#the-resolution-algorithm:code-unit-prefix-2
          x-internal="code-unit-prefix"} of
          `normalizedSpecifier`{.variable}; and

        - either `asURL`{.variable} is null, or `asURL`{.variable} [is
          special](https://url.spec.whatwg.org/#is-special){#the-resolution-algorithm:is-special
          x-internal="is-special"},

        then:

        1.  If `resolutionResult`{.variable} is null, then throw a
            [`TypeError`{#the-resolution-algorithm:typeerror-3}](https://tc39.es/ecma262/#sec-native-error-types-used-in-this-standard-typeerror){x-internal="typeerror"}
            indicating that the resolution of `specifierKey`{.variable}
            was blocked by a null entry.

            This will terminate the entire [resolve a module
            specifier](#resolve-a-module-specifier){#the-resolution-algorithm:resolve-a-module-specifier-7}
            algorithm, without any further fallbacks.

        2.  [Assert](https://infra.spec.whatwg.org/#assert){#the-resolution-algorithm:assert-3
            x-internal="assert"}: `resolutionResult`{.variable} is a
            [URL](https://url.spec.whatwg.org/#concept-url){#the-resolution-algorithm:url-7
            x-internal="url"}.

        3.  Let `afterPrefix`{.variable} be the portion of
            `normalizedSpecifier`{.variable} after the initial
            `specifierKey`{.variable} prefix.

        4.  [Assert](https://infra.spec.whatwg.org/#assert){#the-resolution-algorithm:assert-4
            x-internal="assert"}: `resolutionResult`{.variable},
            [serialized](https://url.spec.whatwg.org/#concept-url-serializer){#the-resolution-algorithm:concept-url-serializer-3
            x-internal="concept-url-serializer"}, ends with U+002F (/),
            as enforced during
            [parsing](#parse-an-import-map-string){#the-resolution-algorithm:parse-an-import-map-string}.

        5.  Let `url`{.variable} be the result of [URL
            parsing](https://url.spec.whatwg.org/#concept-url-parser){#the-resolution-algorithm:url-parser
            x-internal="url-parser"} `afterPrefix`{.variable} with
            `resolutionResult`{.variable}.

        6.  If `url`{.variable} is failure, then throw a
            [`TypeError`{#the-resolution-algorithm:typeerror-4}](https://tc39.es/ecma262/#sec-native-error-types-used-in-this-standard-typeerror){x-internal="typeerror"}
            indicating that resolution of
            `normalizedSpecifier`{.variable} was blocked since the
            `afterPrefix`{.variable} portion could not be URL-parsed
            relative to the `resolutionResult`{.variable} mapped to by
            the `specifierKey`{.variable} prefix.

            This will terminate the entire [resolve a module
            specifier](#resolve-a-module-specifier){#the-resolution-algorithm:resolve-a-module-specifier-8}
            algorithm, without any further fallbacks.

        7.  [Assert](https://infra.spec.whatwg.org/#assert){#the-resolution-algorithm:assert-5
            x-internal="assert"}: `url`{.variable} is a
            [URL](https://url.spec.whatwg.org/#concept-url){#the-resolution-algorithm:url-8
            x-internal="url"}.

        8.  If the
            [serialization](https://url.spec.whatwg.org/#concept-url-serializer){#the-resolution-algorithm:concept-url-serializer-4
            x-internal="concept-url-serializer"} of
            `resolutionResult`{.variable} is not a [code unit
            prefix](https://infra.spec.whatwg.org/#code-unit-prefix){#the-resolution-algorithm:code-unit-prefix-3
            x-internal="code-unit-prefix"} of the
            [serialization](https://url.spec.whatwg.org/#concept-url-serializer){#the-resolution-algorithm:concept-url-serializer-5
            x-internal="concept-url-serializer"} of `url`{.variable},
            then throw a
            [`TypeError`{#the-resolution-algorithm:typeerror-5}](https://tc39.es/ecma262/#sec-native-error-types-used-in-this-standard-typeerror){x-internal="typeerror"}
            indicating that the resolution of
            `normalizedSpecifier`{.variable} was blocked due to it
            backtracking above its prefix `specifierKey`{.variable}.

            This will terminate the entire [resolve a module
            specifier](#resolve-a-module-specifier){#the-resolution-algorithm:resolve-a-module-specifier-9}
            algorithm, without any further fallbacks.

        9.  Return `url`{.variable}.

2.  Return null.

    The [resolve a module
    specifier](#resolve-a-module-specifier){#the-resolution-algorithm:resolve-a-module-specifier-10}
    algorithm will fall back to a less-specific scope, or to
    \"`imports`\", if possible.

To [resolve a URL-like module
specifier]{#resolving-a-url-like-module-specifier .dfn}, given a
[string](https://infra.spec.whatwg.org/#string){#the-resolution-algorithm:string-3
x-internal="string"} `specifier`{.variable} and a
[URL](https://url.spec.whatwg.org/#concept-url){#the-resolution-algorithm:url-9
x-internal="url"} `baseURL`{.variable}:

1.  If `specifier`{.variable} [starts
    with](https://infra.spec.whatwg.org/#string-starts-with){#the-resolution-algorithm:starts-with
    x-internal="starts-with"} \"`/`\", \"`./`\", or \"`../`\", then:

    1.  Let `url`{.variable} be the result of [URL
        parsing](https://url.spec.whatwg.org/#concept-url-parser){#the-resolution-algorithm:url-parser-2
        x-internal="url-parser"} `specifier`{.variable} with
        `baseURL`{.variable}.

    2.  If `url`{.variable} is failure, then return null.

        One way this could happen is if `specifier`{.variable} is
        \"`../foo`\" and `baseURL`{.variable} is a
        [`data:`{#the-resolution-algorithm:data-protocol}](https://www.rfc-editor.org/rfc/rfc2397#section-2){x-internal="data-protocol"}
        URL.

    3.  Return `url`{.variable}.

    This includes cases where `specifier`{.variable} [starts
    with](https://infra.spec.whatwg.org/#string-starts-with){#the-resolution-algorithm:starts-with-2
    x-internal="starts-with"} \"`//`\", i.e., scheme-relative URLs.
    Thus, `url`{.variable} might end up with a different
    [host](https://url.spec.whatwg.org/#concept-url-host){#the-resolution-algorithm:concept-url-host
    x-internal="concept-url-host"} than `baseURL`{.variable}.

2.  Let `url`{.variable} be the result of [URL
    parsing](https://url.spec.whatwg.org/#concept-url-parser){#the-resolution-algorithm:url-parser-3
    x-internal="url-parser"} `specifier`{.variable} (with no base URL).

3.  If `url`{.variable} is failure, then return null.

4.  Return `url`{.variable}.

##### [8.1.5.2]{.secno} Import maps[](#import-maps){.self-link}

An [import map](#import-map){#import-map-dev} allows control over module
specifier resolution. Import maps are delivered via inline
[`script`{#import-maps:the-script-element}](scripting.html#the-script-element)
elements with their
[`type`{#import-maps:attr-script-type}](scripting.html#attr-script-type)
attribute set to \"`importmap`\", and with their [child text
content](https://dom.spec.whatwg.org/#concept-child-text-content){#import-maps:child-text-content
x-internal="child-text-content"} containing a JSON representation of the
import map.

A [`Document`{#import-maps:document}](dom.html#document) can have
multiple import maps processed, which can happen either before or after
any modules have been imported, e.g., via
[`import()`{#import-maps:import()}](https://tc39.es/ecma262/#sec-import-calls){x-internal="import()"}
expressions or
[`script`{#import-maps:the-script-element-2}](scripting.html#the-script-element)
elements with their
[`type`{#import-maps:attr-script-type-2}](scripting.html#attr-script-type)
attribute set to \"`module`\". The [merge existing and new import
maps](#merge-existing-and-new-import-maps){#import-maps:merge-existing-and-new-import-maps}
algorithm ensures that new import maps cannot define the module
resolution for modules that were already defined by past import maps, or
for ones that were already resolved.

::: {#example-import-map-bare-specifier .example}
The simplest use of import maps is to globally remap a bare module
specifier:

``` json
{
  "imports": {
    "moment": "/node_modules/moment/src/moment.js"
  }
}
```

This enables statements like `import moment from "moment";`{.js} to
work, fetching and evaluating the JavaScript module at the
`/node_modules/moment/src/moment.js` URL.
:::

::: {#example-import-map-trailing-slashes .example}
An import map can remap a class of module specifiers into a class of
URLs by using trailing slashes, like so:

``` json
{
  "imports": {
    "moment/": "/node_modules/moment/src/"
  }
}
```

This enables statements like
`import localeData from "moment/locale/zh-cn.js";`{.js} to work,
fetching and evaluating the JavaScript module at the
`/node_modules/moment/src/locale/zh-cn.js` URL. Such trailing-slash
mappings are often combined with bare-specifier mappings, e.g.

``` json
{
  "imports": {
    "moment": "/node_modules/moment/src/moment.js",
    "moment/": "/node_modules/moment/src/"
  }
}
```

so that both the \"main module\" specified by \"`moment`\" and the
\"submodules\" specified by paths such as \"`moment/locale/zh-cn.js`\"
are available.
:::

::: {#example-import-map-url-like-specifier .example}
Bare specifiers are not the only type of module specifiers which import
maps can remap. \"URL-like\" specifiers, i.e., those that are either
parseable as absolute URLs or start with \"`/`\", \"`./`\", or
\"`../`\", can be remapped as well:

``` json
{
  "imports": {
    "https://cdn.example.com/vue/dist/vue.runtime.esm.js": "/node_modules/vue/dist/vue.runtime.esm.js",
    "/js/app.mjs": "/js/app-8e0d62a03.mjs",
    "../helpers/": "https://cdn.example/helpers/"
  }
}
```

Note how the URL to be remapped, as well as the URL being mapped to, can
be specified either as absolute URLs, or as relative URLs starting with
\"`/`\", \"`./`\", or \"`../`\". (They cannot be specified as relative
URLs without those starting sigils, as those help distinguish from bare
module specifiers.) Also note how the [trailing slash
mapping](#example-import-map-trailing-slashes) works in this context as
well.

Such remappings operate on the post-canonicalization URL, and do not
require a match between the literal strings supplied in the import map
key and the imported module specifier. So for example, if this import
map was included on `https://example.com/app.html`, then not only would
`import "/js/app.mjs"`{.js} be remapped, but so would
`import "./js/app.mjs"`{.js} and `import "./foo/../js/app.mjs"`{.js}.
:::

::: {#example-import-map-scopes .example}
All previous examples have globally remapped module specifiers, by using
the top-level \"`imports`\" key in the import map. The top-level
\"`scopes`\" key can be used to provide localized remappings, which only
apply when the referring module matches a specific URL prefix. For
example:

``` json
{
  "scopes": {
    "/a/" : {
      "moment": "/node_modules/moment/src/moment.js"
    },
    "/b/" : {
      "moment": "https://cdn.example.com/moment/src/moment.js"
    }
  }
}
```

With this import map, the statement `import "moment"` will have
different meanings depending on which referrer script contains the
statement:

- Inside scripts located under `/a/`, this will import
  `/node_modules/moment/src/moment.js`.

- Inside scripts located under `/b/`, this will import
  `https://cdn.example.com/moment/src/moment.js`.

- Inside scripts located under `/c/`, this will fail to resolve and thus
  throw an exception.

A typical usage of scopes is to allow multiple versions of the \"same\"
module to exist in a web application, with some parts of the module
graph importing one version, and other parts importing another version.
:::

::: {#example-import-map-scopes-overlapping .example}
Scopes can overlap each other, and overlap the global \"`imports`\"
specifier map. At resolution time, scopes are consulted in order of
most- to least-specific, where specificity is measured by sorting the
scopes using the [code unit less
than](https://infra.spec.whatwg.org/#code-unit-less-than){#import-maps:code-unit-less-than
x-internal="code-unit-less-than"} operation. So, for example,
\"`/scope2/scope3/`\" is treated as more specific than \"`/scope2/`\",
which is treated as more specific than the top-level (unscoped)
mappings.

The following import map illustrates this:

``` json
{
  "imports": {
    "a": "/a-1.mjs",
    "b": "/b-1.mjs",
    "c": "/c-1.mjs"
  },
  "scopes": {
    "/scope2/": {
      "a": "/a-2.mjs"
    },
    "/scope2/scope3/": {
      "b": "/b-3.mjs"
    }
  }
}
```

This results in the following resolutions (using relative URLs for
brevity):

Specifier

\"`a`\"

\"`b`\"

\"`c`\"

Referrer

`/scope1/r.mjs`

`/a-1.mjs`

`/b-1.mjs`

`/c-1.mjs`

`/scope2/r.mjs`

`/a-2.mjs`

`/b-1.mjs`

`/c-1.mjs`

`/scope2/scope3/r.mjs`

`/a-2.mjs`

`/b-3.mjs`

`/c-1.mjs`
:::

::: {#example-import-map-integrity .example}
Import maps can also be used to provide modules with integrity metadata
to be used in Subresource Integrity checks.
[\[SRI\]](references.html#refsSRI)

The following import map illustrates this:

``` json
{
  "imports": {
    "a": "/a-1.mjs",
    "b": "/b-1.mjs",
    "c": "/c-1.mjs"
  },
  "integrity": {
    "/a-1.mjs": "sha384-Li9vy3DqF8tnTXuiaAJuML3ky+er10rcgNR/VqsVpcw+ThHmYcwiB1pbOxEbzJr7",
    "/d-1.mjs": "sha384-MBO5IDfYaE6c6Aao94oZrIOiC6CGiSN2n4QUbHNPhzk5Xhm0djZLQqTpL0HzTUxk"
  }
}
```

The above example provides integrity metadata to be enforced on the
modules `/a-1.mjs` and `/d-1.mjs`, even if the latter is not defined as
an import in the map.
:::

------------------------------------------------------------------------

The [child text
content](https://dom.spec.whatwg.org/#concept-child-text-content){#import-maps:child-text-content-2
x-internal="child-text-content"} of a
[`script`{#import-maps:the-script-element-3}](scripting.html#the-script-element)
element representing an [import
map](#import-map){#import-maps:import-map} must match the following
[import map authoring requirements]{#import-map-authoring-requirements
.dfn}:

- It must be valid JSON. [\[JSON\]](references.html#refsJSON)

- The JSON must represent a JSON object, with at most the three keys
  \"`imports`\", \"`scopes`\", and \"`integrity`\".

- The values corresponding to the \"`imports`\", \"`scopes`\", and
  \"`integrity`\" keys, if present, must themselves be JSON objects.

- The value corresponding to the \"`imports`\" key, if present, must be
  a [valid module specifier
  map](#valid-module-specifier-map){#import-maps:valid-module-specifier-map}.

- The value corresponding to the \"`scopes`\" key, if present, must be a
  JSON object, whose keys are [valid URL
  strings](https://url.spec.whatwg.org/#valid-url-string){#import-maps:valid-url-string
  x-internal="valid-url-string"} and whose values are [valid module
  specifier
  maps](#valid-module-specifier-map){#import-maps:valid-module-specifier-map-2}.

- The value corresponding to the \"`integrity`\" key, if present, must
  be a JSON object, whose keys are [valid URL
  strings](https://url.spec.whatwg.org/#valid-url-string){#import-maps:valid-url-string-2
  x-internal="valid-url-string"} and whose values fit [the requirements
  of the integrity
  attribute](https://w3c.github.io/webappsec-subresource-integrity/#the-integrity-attribute){#import-maps:the-requirements-of-the-integrity-attribute
  x-internal="the-requirements-of-the-integrity-attribute"}.

A [valid module specifier map]{#valid-module-specifier-map .dfn} is a
JSON object that meets the following requirements:

- All of its keys must be nonempty.

- All of its values must be strings.

- Each value must be either a [valid absolute
  URL](https://url.spec.whatwg.org/#syntax-url-absolute){#import-maps:absolute-url
  x-internal="absolute-url"} or a [valid URL
  string](https://url.spec.whatwg.org/#valid-url-string){#import-maps:valid-url-string-3
  x-internal="valid-url-string"} that [starts
  with](https://infra.spec.whatwg.org/#string-starts-with){#import-maps:starts-with
  x-internal="starts-with"} \"`/`\", \"`./`\", or \"`../`\".

- If a given key [ends
  with](https://infra.spec.whatwg.org/#string-ends-with){#import-maps:ends-with
  x-internal="ends-with"} \"`/`\", then the corresponding value must
  also.

##### [8.1.5.3]{.secno} Import map processing model[](#import-map-processing-model){.self-link}

Formally, an [import map]{#import-map .dfn} is a
[struct](https://infra.spec.whatwg.org/#struct){#import-map-processing-model:struct
x-internal="struct"} with three
[items](https://infra.spec.whatwg.org/#struct-item){#import-map-processing-model:struct-item
x-internal="struct-item"}:

- [imports]{#concept-import-map-imports .dfn}, a [module specifier
  map](#module-specifier-map){#import-map-processing-model:module-specifier-map};

- [scopes]{#concept-import-map-scopes .dfn}, an [ordered
  map](https://infra.spec.whatwg.org/#ordered-map){#import-map-processing-model:ordered-map
  x-internal="ordered-map"} of
  [URLs](https://url.spec.whatwg.org/#concept-url){#import-map-processing-model:url
  x-internal="url"} to [module specifier
  maps](#module-specifier-map){#import-map-processing-model:module-specifier-map-2};
  and

- [integrity]{#concept-import-map-integrity .dfn}, a [module integrity
  map](#module-integrity-map){#import-map-processing-model:module-integrity-map}.

A [module specifier map]{#module-specifier-map .dfn} is an [ordered
map](https://infra.spec.whatwg.org/#ordered-map){#import-map-processing-model:ordered-map-2
x-internal="ordered-map"} whose
[keys](https://infra.spec.whatwg.org/#map-key){#import-map-processing-model:map-key
x-internal="map-key"} are
[strings](https://infra.spec.whatwg.org/#string){#import-map-processing-model:string
x-internal="string"} and whose
[values](https://infra.spec.whatwg.org/#map-value){#import-map-processing-model:map-value
x-internal="map-value"} are either
[URLs](https://url.spec.whatwg.org/#concept-url){#import-map-processing-model:url-2
x-internal="url"} or nulls.

A [module integrity map]{#module-integrity-map .dfn} is an [ordered
map](https://infra.spec.whatwg.org/#ordered-map){#import-map-processing-model:ordered-map-3
x-internal="ordered-map"} whose
[keys](https://infra.spec.whatwg.org/#map-key){#import-map-processing-model:map-key-2
x-internal="map-key"} are
[URLs](https://url.spec.whatwg.org/#concept-url){#import-map-processing-model:url-3
x-internal="url"} and whose
[values](https://infra.spec.whatwg.org/#map-value){#import-map-processing-model:map-value-2
x-internal="map-value"} are
[strings](https://infra.spec.whatwg.org/#string){#import-map-processing-model:string-2
x-internal="string"} that will be used as [integrity
metadata](https://fetch.spec.whatwg.org/#concept-request-integrity-metadata){#import-map-processing-model:concept-request-integrity-metadata
x-internal="concept-request-integrity-metadata"}.

An [empty import map]{#empty-import-map .dfn} is an [import
map](#import-map){#import-map-processing-model:import-map} with its
[imports](#concept-import-map-imports){#import-map-processing-model:concept-import-map-imports}
and
[scopes](#concept-import-map-scopes){#import-map-processing-model:concept-import-map-scopes}
both being empty maps.

------------------------------------------------------------------------

A [specifier resolution record]{#specifier-resolution-record .dfn} is a
[struct](https://infra.spec.whatwg.org/#struct){#import-map-processing-model:struct-2
x-internal="struct"}. It has the following
[items](https://infra.spec.whatwg.org/#struct-item){#import-map-processing-model:struct-item-2
x-internal="struct-item"}:

A [serialized base URL]{#specifier-resolution-record-serialized-base-url .dfn}
:   A
    [string](https://infra.spec.whatwg.org/#string){#import-map-processing-model:string-3
    x-internal="string"}-or-null that represents the base URL of the
    specifier, when one exists.

A [specifier]{#specifier-resolution-record-specifier .dfn}
:   A
    [string](https://infra.spec.whatwg.org/#string){#import-map-processing-model:string-4
    x-internal="string"} representing the specifier.

A [specifier as a URL]{#specifier-resolution-record-as-url .dfn}
:   A
    [URL](https://url.spec.whatwg.org/#concept-url){#import-map-processing-model:url-4
    x-internal="url"}-or-null that represents the URL in case of a
    URL-like module specifier.

Implementations can replace [specifier as a
URL](#specifier-resolution-record-as-url){#import-map-processing-model:specifier-resolution-record-as-url}
with a boolean that indicates that the specifier is either bare or
URL-like that [is
special](https://url.spec.whatwg.org/#is-special){#import-map-processing-model:is-special
x-internal="is-special"}.

To [add module to resolved module
set]{#add-module-to-resolved-module-set .dfn} given an [environment
settings
object](#environment-settings-object){#import-map-processing-model:environment-settings-object}
`settingsObject`{.variable}, a
[string](https://infra.spec.whatwg.org/#string){#import-map-processing-model:string-5
x-internal="string"} `serializedBaseURL`{.variable}, a
[string](https://infra.spec.whatwg.org/#string){#import-map-processing-model:string-6
x-internal="string"} `normalizedSpecifier`{.variable}, and a
[URL](https://url.spec.whatwg.org/#concept-url){#import-map-processing-model:url-5
x-internal="url"}-or-null `asURL`{.variable}:

1.  Let `global`{.variable} be `settingsObject`{.variable}\'s [global
    object](#concept-settings-object-global){#import-map-processing-model:concept-settings-object-global}.

2.  If `global`{.variable} does not implement
    [`Window`{#import-map-processing-model:window}](nav-history-apis.html#window),
    then return.

3.  Let `record`{.variable} be a new [specifier resolution
    record](#specifier-resolution-record){#import-map-processing-model:specifier-resolution-record},
    with [serialized base
    URL](#specifier-resolution-record-serialized-base-url){#import-map-processing-model:specifier-resolution-record-serialized-base-url}
    set to `serializedBaseURL`{.variable},
    [specifier](#specifier-resolution-record-specifier){#import-map-processing-model:specifier-resolution-record-specifier}
    set to `normalizedSpecifier`{.variable}, and [specifier as a
    URL](#specifier-resolution-record-as-url){#import-map-processing-model:specifier-resolution-record-as-url-2}
    set to `asURL`{.variable}.

4.  [Append](https://infra.spec.whatwg.org/#set-append){#import-map-processing-model:set-append
    x-internal="set-append"} `record`{.variable} to
    `global`{.variable}\'s [resolved module
    set](#resolved-module-set){#import-map-processing-model:resolved-module-set}.

------------------------------------------------------------------------

To [parse an import map string]{#parse-an-import-map-string .dfn}, given
a
[string](https://infra.spec.whatwg.org/#string){#import-map-processing-model:string-7
x-internal="string"} `input`{.variable} and a
[URL](https://url.spec.whatwg.org/#concept-url){#import-map-processing-model:url-6
x-internal="url"} `baseURL`{.variable}:

1.  Let `parsed`{.variable} be the result of [parsing a JSON string to
    an Infra
    value](https://infra.spec.whatwg.org/#parse-a-json-string-to-an-infra-value){#import-map-processing-model:parse-a-json-string-to-an-infra-value
    x-internal="parse-a-json-string-to-an-infra-value"} given
    `input`{.variable}.

2.  If `parsed`{.variable} is not an [ordered
    map](https://infra.spec.whatwg.org/#ordered-map){#import-map-processing-model:ordered-map-4
    x-internal="ordered-map"}, then throw a
    [`TypeError`{#import-map-processing-model:typeerror}](https://tc39.es/ecma262/#sec-native-error-types-used-in-this-standard-typeerror){x-internal="typeerror"}
    indicating that the top-level value needs to be a JSON object.

3.  Let `sortedAndNormalizedImports`{.variable} be an empty [ordered
    map](https://infra.spec.whatwg.org/#ordered-map){#import-map-processing-model:ordered-map-5
    x-internal="ordered-map"}.

4.  If `parsed`{.variable}\[\"`imports`\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#import-map-processing-model:map-exists
    x-internal="map-exists"}, then:

    1.  If `parsed`{.variable}\[\"`imports`\"\] is not an [ordered
        map](https://infra.spec.whatwg.org/#ordered-map){#import-map-processing-model:ordered-map-6
        x-internal="ordered-map"}, then throw a
        [`TypeError`{#import-map-processing-model:typeerror-2}](https://tc39.es/ecma262/#sec-native-error-types-used-in-this-standard-typeerror){x-internal="typeerror"}
        indicating that the value for the \"`imports`\" top-level key
        needs to be a JSON object.

    2.  Set `sortedAndNormalizedImports`{.variable} to the result of
        [sorting and normalizing a module specifier
        map](#sorting-and-normalizing-a-module-specifier-map){#import-map-processing-model:sorting-and-normalizing-a-module-specifier-map}
        given `parsed`{.variable}\[\"`imports`\"\] and
        `baseURL`{.variable}.

5.  Let `sortedAndNormalizedScopes`{.variable} be an empty [ordered
    map](https://infra.spec.whatwg.org/#ordered-map){#import-map-processing-model:ordered-map-7
    x-internal="ordered-map"}.

6.  If `parsed`{.variable}\[\"`scopes`\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#import-map-processing-model:map-exists-2
    x-internal="map-exists"}, then:

    1.  If `parsed`{.variable}\[\"`scopes`\"\] is not an [ordered
        map](https://infra.spec.whatwg.org/#ordered-map){#import-map-processing-model:ordered-map-8
        x-internal="ordered-map"}, then throw a
        [`TypeError`{#import-map-processing-model:typeerror-3}](https://tc39.es/ecma262/#sec-native-error-types-used-in-this-standard-typeerror){x-internal="typeerror"}
        indicating that the value for the \"`scopes`\" top-level key
        needs to be a JSON object.

    2.  Set `sortedAndNormalizedScopes`{.variable} to the result of
        [sorting and normalizing
        scopes](#sorting-and-normalizing-scopes){#import-map-processing-model:sorting-and-normalizing-scopes}
        given `parsed`{.variable}\[\"`scopes`\"\] and
        `baseURL`{.variable}.

7.  Let `normalizedIntegrity`{.variable} be an empty [ordered
    map](https://infra.spec.whatwg.org/#ordered-map){#import-map-processing-model:ordered-map-9
    x-internal="ordered-map"}.

8.  If `parsed`{.variable}\[\"`integrity`\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#import-map-processing-model:map-exists-3
    x-internal="map-exists"}, then:

    1.  If `parsed`{.variable}\[\"`integrity`\"\] is not an [ordered
        map](https://infra.spec.whatwg.org/#ordered-map){#import-map-processing-model:ordered-map-10
        x-internal="ordered-map"}, then throw a
        [`TypeError`{#import-map-processing-model:typeerror-4}](https://tc39.es/ecma262/#sec-native-error-types-used-in-this-standard-typeerror){x-internal="typeerror"}
        indicating that the value for the \"`integrity`\" top-level key
        needs to be a JSON object.

    2.  Set `normalizedIntegrity`{.variable} to the result of
        [normalizing a module integrity
        map](#normalizing-a-module-integrity-map){#import-map-processing-model:normalizing-a-module-integrity-map}
        given `parsed`{.variable}\[\"`integrity`\"\] and
        `baseURL`{.variable}.

9.  If `parsed`{.variable}\'s
    [keys](https://infra.spec.whatwg.org/#map-key){#import-map-processing-model:map-key-3
    x-internal="map-key"}
    [contains](https://infra.spec.whatwg.org/#list-contain){#import-map-processing-model:list-contains
    x-internal="list-contains"} any items besides \"`imports`\",
    \"`scopes`\", or \"`integrity`\", then the user agent should [report
    a warning to the
    console](https://console.spec.whatwg.org/#report-a-warning-to-the-console){#import-map-processing-model:report-a-warning-to-the-console
    x-internal="report-a-warning-to-the-console"} indicating that an
    invalid top-level key was present in the import map.

    This can help detect typos. It is not an error, because that would
    prevent any future extensions from being added backward-compatibly.

10. Return an [import
    map](#import-map){#import-map-processing-model:import-map-2} whose
    [imports](#concept-import-map-imports){#import-map-processing-model:concept-import-map-imports-2}
    are `sortedAndNormalizedImports`{.variable}, whose
    [scopes](#concept-import-map-scopes){#import-map-processing-model:concept-import-map-scopes-2}
    are `sortedAndNormalizedScopes`{.variable}, and whose
    [integrity](#concept-import-map-integrity){#import-map-processing-model:concept-import-map-integrity}
    are `normalizedIntegrity`{.variable}.

::: {#example-import-map-normalization .example}
The [import map](#import-map){#import-map-processing-model:import-map-3}
that results from this parsing algorithm is highly normalized. For
example, given a base URL of `https://example.com/base/page.html`, the
input

``` json
{
  "imports": {
    "/app/helper": "node_modules/helper/index.mjs",
    "lodash": "/node_modules/lodash-es/lodash.js"
  }
}
```

will generate an [import
map](#import-map){#import-map-processing-model:import-map-4} with
[imports](#concept-import-map-imports){#import-map-processing-model:concept-import-map-imports-3}
of

    «[
      "https://example.com/app/helper" → https://example.com/base/node_modules/helper/index.mjs
      "lodash" → https://example.com/node_modules/lodash-es/lodash.js
    ]»

and (despite nothing being present in the input string) an empty
[ordered
map](https://infra.spec.whatwg.org/#ordered-map){#import-map-processing-model:ordered-map-11
x-internal="ordered-map"} for its
[scopes](#concept-import-map-scopes){#import-map-processing-model:concept-import-map-scopes-3}.
:::

------------------------------------------------------------------------

To [merge module specifier maps]{#merge-module-specifier-maps .dfn},
given a [module specifier
map](#module-specifier-map){#import-map-processing-model:module-specifier-map-3}
`newMap`{.variable} and a [module specifier
map](#module-specifier-map){#import-map-processing-model:module-specifier-map-4}
`oldMap`{.variable}:

1.  Let `mergedMap`{.variable} be a deep copy of `oldMap`{.variable}.

2.  [For
    each](https://infra.spec.whatwg.org/#map-iterate){#import-map-processing-model:map-iterate
    x-internal="map-iterate"} `specifier`{.variable} → `url`{.variable}
    of `newMap`{.variable}:

    1.  If `specifier`{.variable}
        [exists](https://infra.spec.whatwg.org/#map-exists){#import-map-processing-model:map-exists-4
        x-internal="map-exists"} in `oldMap`{.variable}, then:

        1.  The user agent may [report a warning to the
            console](https://console.spec.whatwg.org/#report-a-warning-to-the-console){#import-map-processing-model:report-a-warning-to-the-console-2
            x-internal="report-a-warning-to-the-console"} indicating the
            ignored rule. They may choose to avoid reporting if the rule
            is identical to an existing one.

        2.  [Continue](https://infra.spec.whatwg.org/#iteration-continue){#import-map-processing-model:continue
            x-internal="continue"}.

    2.  Set `mergedMap`{.variable}\[`specifier`{.variable}\] to
        `url`{.variable}.

3.  Return `mergedMap`{.variable}.

To [merge existing and new import
maps]{#merge-existing-and-new-import-maps .dfn}, given a [global
object](#concept-settings-object-global){#import-map-processing-model:concept-settings-object-global-2}
`global`{.variable} and an [import
map](#import-map){#import-map-processing-model:import-map-5}
`newImportMap`{.variable}:

1.  Let `newImportMapScopes`{.variable} be a deep copy of
    `newImportMap`{.variable}\'s
    [scopes](#concept-import-map-scopes){#import-map-processing-model:concept-import-map-scopes-4}.

    We\'re mutating these copies and removing items from them when they
    are used to ignore scope-specific rules. This is true for
    `newImportMapScopes`{.variable}, as well as to
    `newImportMapImports`{.variable} below.

2.  Let `oldImportMap`{.variable} be `global`{.variable}\'s [import
    map](#concept-global-import-map){#import-map-processing-model:concept-global-import-map}.

3.  Let `newImportMapImports`{.variable} be a deep copy of
    `newImportMap`{.variable}\'s
    [imports](#concept-import-map-imports){#import-map-processing-model:concept-import-map-imports-4}.

4.  [For
    each](https://infra.spec.whatwg.org/#map-iterate){#import-map-processing-model:map-iterate-2
    x-internal="map-iterate"} `scopePrefix`{.variable} →
    `scopeImports`{.variable} of `newImportMapScopes`{.variable}:

    1.  [For
        each](https://infra.spec.whatwg.org/#list-iterate){#import-map-processing-model:list-iterate
        x-internal="list-iterate"} `record`{.variable} of
        `global`{.variable}\'s [resolved module
        set](#resolved-module-set){#import-map-processing-model:resolved-module-set-2}:

        1.  If `scopePrefix`{.variable} is `record`{.variable}\'s
            [serialized base
            URL](#specifier-resolution-record-serialized-base-url){#import-map-processing-model:specifier-resolution-record-serialized-base-url-2},
            or if `scopePrefix`{.variable} ends with U+002F (/) and
            `scopePrefix`{.variable} is a [code unit
            prefix](https://infra.spec.whatwg.org/#code-unit-prefix){#import-map-processing-model:code-unit-prefix
            x-internal="code-unit-prefix"} of `record`{.variable}\'s
            [serialized base
            URL](#specifier-resolution-record-serialized-base-url){#import-map-processing-model:specifier-resolution-record-serialized-base-url-3},
            then:

            1.  [For
                each](https://infra.spec.whatwg.org/#list-iterate){#import-map-processing-model:list-iterate-2
                x-internal="list-iterate"} `specifierKey`{.variable} →
                `resolutionResult`{.variable} of
                `scopeImports`{.variable}:

                1.  If `specifierKey`{.variable} is
                    `record`{.variable}\'s
                    [specifier](#specifier-resolution-record-specifier){#import-map-processing-model:specifier-resolution-record-specifier-2},
                    or if all of the following conditions are true:

                    - `specifierKey`{.variable} ends with U+002F (/);

                    - `specifierKey`{.variable} is a [code unit
                      prefix](https://infra.spec.whatwg.org/#code-unit-prefix){#import-map-processing-model:code-unit-prefix-2
                      x-internal="code-unit-prefix"} of
                      `record`{.variable}\'s
                      [specifier](#specifier-resolution-record-specifier){#import-map-processing-model:specifier-resolution-record-specifier-3};

                    - either `record`{.variable}\'s [specifier as a
                      URL](#specifier-resolution-record-as-url){#import-map-processing-model:specifier-resolution-record-as-url-3}
                      is null or [is
                      special](https://url.spec.whatwg.org/#is-special){#import-map-processing-model:is-special-2
                      x-internal="is-special"},

                    then:

                    1.  The user agent may [report a warning to the
                        console](https://console.spec.whatwg.org/#report-a-warning-to-the-console){#import-map-processing-model:report-a-warning-to-the-console-3
                        x-internal="report-a-warning-to-the-console"}
                        indicating the ignored rule. They may choose to
                        avoid reporting if the rule is identical to an
                        existing one.

                    2.  Remove
                        `scopeImports`{.variable}\[`specifierKey`{.variable}\].

        Implementers are encouraged to implement a more efficient
        matching algorithm when working with the [resolved module
        set](#resolved-module-set){#import-map-processing-model:resolved-module-set-3}.
        As guidance, the number of resolved/mapped modules in a large
        application can be on the order of thousands.

    2.  If `scopePrefix`{.variable}
        [exists](https://infra.spec.whatwg.org/#map-exists){#import-map-processing-model:map-exists-5
        x-internal="map-exists"} in `oldImportMap`{.variable}\'s
        [scopes](#concept-import-map-scopes){#import-map-processing-model:concept-import-map-scopes-5},
        then set `oldImportMap`{.variable}\'s
        [scopes](#concept-import-map-scopes){#import-map-processing-model:concept-import-map-scopes-6}\[`scopePrefix`{.variable}\]
        to the result of [merging module specifier
        maps](#merge-module-specifier-maps){#import-map-processing-model:merge-module-specifier-maps},
        given `scopeImports`{.variable} and `oldImportMap`{.variable}\'s
        [scopes](#concept-import-map-scopes){#import-map-processing-model:concept-import-map-scopes-7}\[`scopePrefix`{.variable}\].

    3.  Otherwise, set `oldImportMap`{.variable}\'s
        [scopes](#concept-import-map-scopes){#import-map-processing-model:concept-import-map-scopes-8}\[`scopePrefix`{.variable}\]
        to `scopeImports`{.variable}.

5.  [For
    each](https://infra.spec.whatwg.org/#map-iterate){#import-map-processing-model:map-iterate-3
    x-internal="map-iterate"} `url`{.variable} → `integrity`{.variable}
    of `newImportMap`{.variable}\'s
    [integrity](#concept-import-map-integrity){#import-map-processing-model:concept-import-map-integrity-2}:

    1.  If `url`{.variable}
        [exists](https://infra.spec.whatwg.org/#map-exists){#import-map-processing-model:map-exists-6
        x-internal="map-exists"} in `oldImportMap`{.variable}\'s
        [integrity](#concept-import-map-integrity){#import-map-processing-model:concept-import-map-integrity-3},
        then:

        1.  The user agent may [report a warning to the
            console](https://console.spec.whatwg.org/#report-a-warning-to-the-console){#import-map-processing-model:report-a-warning-to-the-console-4
            x-internal="report-a-warning-to-the-console"} indicating the
            ignored rule. They may choose to avoid reporting if the rule
            is identical to an existing one.

        2.  [Continue](https://infra.spec.whatwg.org/#iteration-continue){#import-map-processing-model:continue-2
            x-internal="continue"}.

    2.  Set `oldImportMap`{.variable}\'s
        [integrity](#concept-import-map-integrity){#import-map-processing-model:concept-import-map-integrity-4}\[`url`{.variable}\]
        to `integrity`{.variable}.

6.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#import-map-processing-model:list-iterate-3
    x-internal="list-iterate"} `record`{.variable} of
    `global`{.variable}\'s [resolved module
    set](#resolved-module-set){#import-map-processing-model:resolved-module-set-4}:

    1.  [For
        each](https://infra.spec.whatwg.org/#list-iterate){#import-map-processing-model:list-iterate-4
        x-internal="list-iterate"} `specifier`{.variable} →
        `url`{.variable} of `newImportMapImports`{.variable}:

        1.  If `specifier`{.variable} [starts
            with](https://infra.spec.whatwg.org/#string-starts-with){#import-map-processing-model:starts-with
            x-internal="starts-with"} `record`{.variable}\'s
            [specifier](#specifier-resolution-record-specifier){#import-map-processing-model:specifier-resolution-record-specifier-4},
            then:

            1.  The user agent may [report a warning to the
                console](https://console.spec.whatwg.org/#report-a-warning-to-the-console){#import-map-processing-model:report-a-warning-to-the-console-5
                x-internal="report-a-warning-to-the-console"} indicating
                the ignored rule. They may choose to avoid reporting if
                the rule is identical to an existing one.

            2.  Remove
                `newImportMapImports`{.variable}\[`specifier`{.variable}\].

7.  Set `oldImportMap`{.variable}\'s
    [imports](#concept-import-map-imports){#import-map-processing-model:concept-import-map-imports-5}
    to the result of [merge module specifier
    maps](#merge-module-specifier-maps){#import-map-processing-model:merge-module-specifier-maps-2},
    given `newImportMapImports`{.variable} and
    `oldImportMap`{.variable}\'s
    [imports](#concept-import-map-imports){#import-map-processing-model:concept-import-map-imports-6}.

The above algorithm merges a new import map into the given [environment
settings
object](#environment-settings-object){#import-map-processing-model:environment-settings-object-2}\'s
[global
object](#global-object){#import-map-processing-model:global-object}\'s
[import map](#import-map){#import-map-processing-model:import-map-6}.
Let\'s examine a few examples:

::: {#example-import-map-merge-unrelated .example}
There are two cases when rules of the new import map don\'t get merged
into the existing one.

1.  The new import map rule has the exact same scope and specifier as a
    rule in the existing import map. We\'ll call that \"conflicting
    rule\".

2.  The new import map rule may impact the resolution of an already
    resolved module. We\'ll call that \"impacted already resolved
    module\".

When the new import map has no conflicting rules, and there are no
impacted resolved modules, the resulting map would be a combination of
the new and existing maps. Rules that would have individually impacted
similar modules (e.g. \"/app/\" and \"/app/helper\") but are not an
exact match are not conflicting, and all make it to the merged map.

So, the following existing and new import maps:

``` json
{
   "imports": {
    "/app/": "./original-app/",
  }
}
{
  "imports": {
    "/app/helper": "./helper/index.mjs"
  },
  "scopes": {
    "/js": {
      "/app/": "./js-app/"
    }
  }
}
```

Would be equivalent to the following single import map:

``` json
{
  "imports": {
    "/app/": "./original-app/",
    "/app/helper": "./helper/index.mjs"
  },
  "scopes": {
    "/js": {
      "/app/": "./js-app/"
    }
  }
}
```
:::

::: {#example-import-map-merge-conflict-imports .example}
When the new import map impacts an already resolved module, that rule
gets dropped from the import map.

So, if the [resolved module
set](#resolved-module-set){#import-map-processing-model:resolved-module-set-5}
already contains the \"`/app/helper`\", the following new import map:

``` json
{
   "imports": {
    "/app/helper": "./helper/index.mjs",
    "lodash": "/node_modules/lodash-es/lodash.js"
  }
}
```

Would be equivalent to the following one:

``` json
{
  "imports": {
    "lodash": "/node_modules/lodash-es/lodash.js"
  }
}
```
:::

::: {#example-import-map-merge-conflict-scopes .example}
The same is true for rules that impact already resolved modules defined
in specific scopes. If we already resolved \"`/app/helper`\" from
\"`/app/main.mjs`\" the following new import map:

``` json
{
  "scopes": {
    "/app/": {
      "/app/helper": "./helper/index.mjs"
    }
  },
   "imports": {
    "lodash": "/node_modules/lodash-es/lodash.js"
  }
}
```

Would similarly be equivalent to:

``` json
{
  "imports": {
    "lodash": "/node_modules/lodash-es/lodash.js"
  }
}
```
:::

::: {#example-import-map-merge-conflict-imports-and-scopes .example}
We could also have cases where a single already-resolved module
specifier has multiple rules for its resolution, depending on the
referring script. In such cases, only the relevant rules would not be
added to the map.

For example, if we already resolved \"`/app/helper`\" from
\"`/app/vendor/main.mjs`\", the following new import map:

``` json
{
  "scopes": {
    "/app/": {
      "/app/helper": "./helper/index.mjs"
    },
    "/app/vendor/": {
      "/app/": "./vendor_helper/"
    },
    "/vendor/": {
      "/app/helper": "./helper/vendor_index.mjs"
    }
  },
   "imports": {
    "lodash": "/node_modules/lodash-es/lodash.js"
    "/app/": "./general_app_path/"
    "/app/helper": "./other_path/helper/index.mjs"
  }
}
```

Would be equivalent to:

``` json
{
  "scopes": {
    "/vendor/": {
      "/app/helper": "./helper/vendor_index.mjs"
    }
  },
  "imports": {
    "lodash": "/node_modules/lodash-es/lodash.js"
  }
}
```

This is achieved by the fact that the merge algorithm tracks already
resolved modules and removes rules affecting them from new import maps
before they are merged into the existing one.
:::

::: {#example-import-map-merge-two-map-conflict .example}
When the new import map has conflicting rules to the existing import
map, with no impacted already resolved modules, the existing import map
rules persist.

For example, the following existing and new import maps:

``` json
{
   "imports": {
    "/app/helper": "./helper/index.mjs",
    "lodash": "/node_modules/lodash-es/lodash.js"
  }
}
{
  "imports": {
    "/app/helper": "./main/helper/index.mjs"
  }
}
```

Would be equivalent to the following single import map:

``` json
{
  "imports": {
    "/app/helper": "./helper/index.mjs",
    "lodash": "/node_modules/lodash-es/lodash.js",
  }
}
```
:::

To [sort and normalize a module specifier
map]{#sorting-and-normalizing-a-module-specifier-map .dfn}, given an
[ordered
map](https://infra.spec.whatwg.org/#ordered-map){#import-map-processing-model:ordered-map-12
x-internal="ordered-map"} `originalMap`{.variable} and a
[URL](https://url.spec.whatwg.org/#concept-url){#import-map-processing-model:url-7
x-internal="url"} `baseURL`{.variable}:

1.  Let `normalized`{.variable} be an empty [ordered
    map](https://infra.spec.whatwg.org/#ordered-map){#import-map-processing-model:ordered-map-13
    x-internal="ordered-map"}.

2.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#import-map-processing-model:list-iterate-5
    x-internal="list-iterate"} `specifierKey`{.variable} →
    `value`{.variable} of `originalMap`{.variable}:

    1.  Let `normalizedSpecifierKey`{.variable} be the result of
        [normalizing a specifier
        key](#normalizing-a-specifier-key){#import-map-processing-model:normalizing-a-specifier-key}
        given `specifierKey`{.variable} and `baseURL`{.variable}.

    2.  If `normalizedSpecifierKey`{.variable} is null, then
        [continue](https://infra.spec.whatwg.org/#iteration-continue){#import-map-processing-model:continue-3
        x-internal="continue"}.

    3.  If `value`{.variable} is not a
        [string](https://infra.spec.whatwg.org/#string){#import-map-processing-model:string-8
        x-internal="string"}, then:

        1.  The user agent may [report a warning to the
            console](https://console.spec.whatwg.org/#report-a-warning-to-the-console){#import-map-processing-model:report-a-warning-to-the-console-6
            x-internal="report-a-warning-to-the-console"} indicating
            that addresses need to be strings.

        2.  Set
            `normalized`{.variable}\[`normalizedSpecifierKey`{.variable}\]
            to null.

        3.  [Continue](https://infra.spec.whatwg.org/#iteration-continue){#import-map-processing-model:continue-4
            x-internal="continue"}.

    4.  Let `addressURL`{.variable} be the result of [resolving a
        URL-like module
        specifier](#resolving-a-url-like-module-specifier){#import-map-processing-model:resolving-a-url-like-module-specifier}
        given `value`{.variable} and `baseURL`{.variable}.

    5.  If `addressURL`{.variable} is null, then:

        1.  The user agent may [report a warning to the
            console](https://console.spec.whatwg.org/#report-a-warning-to-the-console){#import-map-processing-model:report-a-warning-to-the-console-7
            x-internal="report-a-warning-to-the-console"} indicating
            that the address was invalid.

        2.  Set
            `normalized`{.variable}\[`normalizedSpecifierKey`{.variable}\]
            to null.

        3.  [Continue](https://infra.spec.whatwg.org/#iteration-continue){#import-map-processing-model:continue-5
            x-internal="continue"}.

    6.  If `specifierKey`{.variable} ends with U+002F (/), and the
        [serialization](https://url.spec.whatwg.org/#concept-url-serializer){#import-map-processing-model:concept-url-serializer
        x-internal="concept-url-serializer"} of `addressURL`{.variable}
        does not end with U+002F (/), then:

        1.  The user agent may [report a warning to the
            console](https://console.spec.whatwg.org/#report-a-warning-to-the-console){#import-map-processing-model:report-a-warning-to-the-console-8
            x-internal="report-a-warning-to-the-console"} indicating
            that an invalid address was given for the specifier key
            `specifierKey`{.variable}; since `specifierKey`{.variable}
            ends with a slash, the address needs to as well.

        2.  Set
            `normalized`{.variable}\[`normalizedSpecifierKey`{.variable}\]
            to null.

        3.  [Continue](https://infra.spec.whatwg.org/#iteration-continue){#import-map-processing-model:continue-6
            x-internal="continue"}.

    7.  Set
        `normalized`{.variable}\[`normalizedSpecifierKey`{.variable}\]
        to `addressURL`{.variable}.

3.  Return the result of [sorting in descending
    order](https://infra.spec.whatwg.org/#map-sort-in-descending-order){#import-map-processing-model:map-sort-descending
    x-internal="map-sort-descending"} `normalized`{.variable}, with an
    entry `a`{.variable} being less than an entry `b`{.variable} if
    `a`{.variable}\'s
    [key](https://infra.spec.whatwg.org/#map-key){#import-map-processing-model:map-key-4
    x-internal="map-key"} is [code unit less
    than](https://infra.spec.whatwg.org/#code-unit-less-than){#import-map-processing-model:code-unit-less-than
    x-internal="code-unit-less-than"} `b`{.variable}\'s
    [key](https://infra.spec.whatwg.org/#map-key){#import-map-processing-model:map-key-5
    x-internal="map-key"}.

To [sort and normalize scopes]{#sorting-and-normalizing-scopes .dfn},
given an [ordered
map](https://infra.spec.whatwg.org/#ordered-map){#import-map-processing-model:ordered-map-14
x-internal="ordered-map"} `originalMap`{.variable} and a
[URL](https://url.spec.whatwg.org/#concept-url){#import-map-processing-model:url-8
x-internal="url"} `baseURL`{.variable}:

1.  Let `normalized`{.variable} be an empty [ordered
    map](https://infra.spec.whatwg.org/#ordered-map){#import-map-processing-model:ordered-map-15
    x-internal="ordered-map"}.

2.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#import-map-processing-model:list-iterate-6
    x-internal="list-iterate"} `scopePrefix`{.variable} →
    `potentialSpecifierMap`{.variable} of `originalMap`{.variable}:

    1.  If `potentialSpecifierMap`{.variable} is not an [ordered
        map](https://infra.spec.whatwg.org/#ordered-map){#import-map-processing-model:ordered-map-16
        x-internal="ordered-map"}, then throw a
        [`TypeError`{#import-map-processing-model:typeerror-5}](https://tc39.es/ecma262/#sec-native-error-types-used-in-this-standard-typeerror){x-internal="typeerror"}
        indicating that the value of the scope with prefix
        `scopePrefix`{.variable} needs to be a JSON object.

    2.  Let `scopePrefixURL`{.variable} be the result of [URL
        parsing](https://url.spec.whatwg.org/#concept-url-parser){#import-map-processing-model:url-parser
        x-internal="url-parser"} `scopePrefix`{.variable} with
        `baseURL`{.variable}.

    3.  If `scopePrefixURL`{.variable} is failure, then:

        1.  The user agent may [report a warning to the
            console](https://console.spec.whatwg.org/#report-a-warning-to-the-console){#import-map-processing-model:report-a-warning-to-the-console-9
            x-internal="report-a-warning-to-the-console"} that the scope
            prefix URL was not parseable.

        2.  [Continue](https://infra.spec.whatwg.org/#iteration-continue){#import-map-processing-model:continue-7
            x-internal="continue"}.

    4.  Let `normalizedScopePrefix`{.variable} be the
        [serialization](https://url.spec.whatwg.org/#concept-url-serializer){#import-map-processing-model:concept-url-serializer-2
        x-internal="concept-url-serializer"} of
        `scopePrefixURL`{.variable}.

    5.  Set
        `normalized`{.variable}\[`normalizedScopePrefix`{.variable}\] to
        the result of [sorting and normalizing a module specifier
        map](#sorting-and-normalizing-a-module-specifier-map){#import-map-processing-model:sorting-and-normalizing-a-module-specifier-map-2}
        given `potentialSpecifierMap`{.variable} and
        `baseURL`{.variable}.

3.  Return the result of [sorting in descending
    order](https://infra.spec.whatwg.org/#map-sort-in-descending-order){#import-map-processing-model:map-sort-descending-2
    x-internal="map-sort-descending"} `normalized`{.variable}, with an
    entry `a`{.variable} being less than an entry `b`{.variable} if
    `a`{.variable}\'s
    [key](https://infra.spec.whatwg.org/#map-key){#import-map-processing-model:map-key-6
    x-internal="map-key"} is [code unit less
    than](https://infra.spec.whatwg.org/#code-unit-less-than){#import-map-processing-model:code-unit-less-than-2
    x-internal="code-unit-less-than"} `b`{.variable}\'s
    [key](https://infra.spec.whatwg.org/#map-key){#import-map-processing-model:map-key-7
    x-internal="map-key"}.

In the above two algorithms, sorting keys and scopes in descending order
has the effect of putting \"`foo/bar/`\" before \"`foo/`\". This in turn
gives \"`foo/bar/`\" a higher priority than \"`foo/`\" during [module
specifier
resolution](#resolve-a-module-specifier){#import-map-processing-model:resolve-a-module-specifier}.

To [normalize a module integrity
map]{#normalizing-a-module-integrity-map .dfn}, given an [ordered
map](https://infra.spec.whatwg.org/#ordered-map){#import-map-processing-model:ordered-map-17
x-internal="ordered-map"} `originalMap`{.variable}:

1.  Let `normalized`{.variable} be an empty [ordered
    map](https://infra.spec.whatwg.org/#ordered-map){#import-map-processing-model:ordered-map-18
    x-internal="ordered-map"}.

2.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#import-map-processing-model:list-iterate-7
    x-internal="list-iterate"} `key`{.variable} → `value`{.variable} of
    `originalMap`{.variable}:

    1.  Let `resolvedURL`{.variable} be the result of [resolving a
        URL-like module
        specifier](#resolving-a-url-like-module-specifier){#import-map-processing-model:resolving-a-url-like-module-specifier-2}
        given `key`{.variable} and `baseURL`{.variable}.

        Unlike \"`imports`\", keys of the integrity map are treated as
        URLs, not module specifiers. However, we use the [resolve a
        URL-like module
        specifier](#resolving-a-url-like-module-specifier){#import-map-processing-model:resolving-a-url-like-module-specifier-3}
        algorithm to prohibit \"bare\" relative URLs like `foo`, which
        could be mistaken for module specifiers.

    2.  If `resolvedURL`{.variable} is null, then:

        1.  The user agent may [report a warning to the
            console](https://console.spec.whatwg.org/#report-a-warning-to-the-console){#import-map-processing-model:report-a-warning-to-the-console-10
            x-internal="report-a-warning-to-the-console"} indicating
            that the key failed to resolve.

        2.  [Continue](https://infra.spec.whatwg.org/#iteration-continue){#import-map-processing-model:continue-8
            x-internal="continue"}.

    3.  If `value`{.variable} is not a
        [string](https://infra.spec.whatwg.org/#string){#import-map-processing-model:string-9
        x-internal="string"}, then:

        1.  The user agent may [report a warning to the
            console](https://console.spec.whatwg.org/#report-a-warning-to-the-console){#import-map-processing-model:report-a-warning-to-the-console-11
            x-internal="report-a-warning-to-the-console"} indicating
            that [integrity
            metadata](https://fetch.spec.whatwg.org/#concept-request-integrity-metadata){#import-map-processing-model:concept-request-integrity-metadata-2
            x-internal="concept-request-integrity-metadata"} values need
            to be
            [strings](https://infra.spec.whatwg.org/#string){#import-map-processing-model:string-10
            x-internal="string"}.

        2.  [Continue](https://infra.spec.whatwg.org/#iteration-continue){#import-map-processing-model:continue-9
            x-internal="continue"}.

    4.  Set `normalized`{.variable}\[`resolvedURL`{.variable}\] to
        `value`{.variable}.

3.  Return `normalized`{.variable}.

To [normalize a specifier key]{#normalizing-a-specifier-key .dfn}, given
a
[string](https://infra.spec.whatwg.org/#string){#import-map-processing-model:string-11
x-internal="string"} `specifierKey`{.variable} and a
[URL](https://url.spec.whatwg.org/#concept-url){#import-map-processing-model:url-9
x-internal="url"} `baseURL`{.variable}:

1.  If `specifierKey`{.variable} is the empty string, then:

    1.  The user agent may [report a warning to the
        console](https://console.spec.whatwg.org/#report-a-warning-to-the-console){#import-map-processing-model:report-a-warning-to-the-console-12
        x-internal="report-a-warning-to-the-console"} indicating that
        specifier keys may not be the empty string.

    2.  Return null.

2.  Let `url`{.variable} be the result of [resolving a URL-like module
    specifier](#resolving-a-url-like-module-specifier){#import-map-processing-model:resolving-a-url-like-module-specifier-4},
    given `specifierKey`{.variable} and `baseURL`{.variable}.

3.  If `url`{.variable} is not null, then return the
    [serialization](https://url.spec.whatwg.org/#concept-url-serializer){#import-map-processing-model:concept-url-serializer-3
    x-internal="concept-url-serializer"} of `url`{.variable}.

4.  Return `specifierKey`{.variable}.

#### [8.1.6]{.secno} JavaScript specification host hooks[](#javascript-specification-host-hooks){.self-link}

The JavaScript specification contains a number of
[implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#javascript-specification-host-hooks:implementation-defined
x-internal="implementation-defined"} abstract operations, that vary
depending on the host environment. This section defines them for user
agent hosts.

##### [8.1.6.1]{.secno} [HostEnsureCanAddPrivateElement]{.dfn}(`O`{.variable})[](#the-hostensurecanaddprivateelement-implementation){.self-link} {#the-hostensurecanaddprivateelement-implementation}

JavaScript contains an
[implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#the-hostensurecanaddprivateelement-implementation:implementation-defined
x-internal="implementation-defined"}
[HostEnsureCanAddPrivateElement](https://tc39.es/ecma262/#sec-hostensurecanaddprivateelement){#the-hostensurecanaddprivateelement-implementation:js-hostensurecanaddprivateelement
x-internal="js-hostensurecanaddprivateelement"}(`O`{.variable}) abstract
operation. User agents must use the following implementation:
[\[JAVASCRIPT\]](references.html#refsJAVASCRIPT)

1.  If `O`{.variable} is a
    [`WindowProxy`{#the-hostensurecanaddprivateelement-implementation:windowproxy}](nav-history-apis.html#windowproxy)
    object, or
    [implements](https://webidl.spec.whatwg.org/#implements){#the-hostensurecanaddprivateelement-implementation:implements
    x-internal="implements"}
    [`Location`{#the-hostensurecanaddprivateelement-implementation:location}](nav-history-apis.html#location),
    then return Completion { \[\[Type\]\]: throw, \[\[Value\]\]: a new
    [`TypeError`{#the-hostensurecanaddprivateelement-implementation:typeerror}](https://tc39.es/ecma262/#sec-native-error-types-used-in-this-standard-typeerror){x-internal="typeerror"}
    }.

2.  Return
    [NormalCompletion](https://tc39.es/ecma262/#sec-normalcompletion){#the-hostensurecanaddprivateelement-implementation:normalcompletion
    x-internal="normalcompletion"}(unused).

JavaScript private fields can be applied to arbitrary objects. Since
this can dramatically complicate implementation for particularly-exotic
host objects, the JavaScript language specification provides this hook
to allow hosts to reject private fields on objects meeting a
host-defined criteria. In the case of HTML,
[`WindowProxy`{#the-hostensurecanaddprivateelement-implementation:windowproxy-2}](nav-history-apis.html#windowproxy)
and
[`Location`{#the-hostensurecanaddprivateelement-implementation:location-2}](nav-history-apis.html#location)
have complicated semantics --- particularly around navigation and
security --- that make implementation of private field semantics
challenging, so our implementation simply rejects those objects.

##### [8.1.6.2]{.secno} [HostEnsureCanCompileStrings]{.dfn}(`realm`{.variable}, `parameterStrings`{.variable}, `bodyString`{.variable}, `codeString`{.variable}, `compilationType`{.variable}, `parameterArgs`{.variable}, `bodyArg`{.variable})[](#hostensurecancompilestrings(realm,-parameterstrings,-bodystring,-codestring,-compilationtype,-parameterargs,-bodyarg)){.self-link} {#hostensurecancompilestrings(realm,-parameterstrings,-bodystring,-codestring,-compilationtype,-parameterargs,-bodyarg)}

JavaScript contains an
[implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#hostensurecancompilestrings(realm,-parameterstrings,-bodystring,-codestring,-compilationtype,-parameterargs,-bodyarg):implementation-defined
x-internal="implementation-defined"}
[HostEnsureCanCompileStrings](https://tc39.es/proposal-dynamic-code-brand-checks/#sec-hostensurecancompilestrings){#hostensurecancompilestrings(realm,-parameterstrings,-bodystring,-codestring,-compilationtype,-parameterargs,-bodyarg):js-hostensurecancompilestrings
x-internal="js-hostensurecancompilestrings"} abstract operation,
redefined by the Dynamic Code Brand Checks proposal. User agents must
use the following implementation:
[\[JAVASCRIPT\]](references.html#refsJAVASCRIPT)
[\[JSDYNAMICCODEBRANDCHECKS\]](references.html#refsJSDYNAMICCODEBRANDCHECKS)

1.  Perform ?
    [EnsureCSPDoesNotBlockStringCompilation](https://w3c.github.io/webappsec-csp/#can-compile-strings){#hostensurecancompilestrings(realm,-parameterstrings,-bodystring,-codestring,-compilationtype,-parameterargs,-bodyarg):csp-ensurecspdoesnotblockstringcompilation
    x-internal="csp-ensurecspdoesnotblockstringcompilation"}(`realm`{.variable},
    `parameterStrings`{.variable}, `bodyString`{.variable},
    `codeString`{.variable}, `compilationType`{.variable},
    `parameterArgs`{.variable}, `bodyArg`{.variable}).
    [\[CSP\]](references.html#refsCSP)

##### [8.1.6.3]{.secno} [HostGetCodeForEval]{.dfn}(`argument`{.variable})[](#hostgetcodeforeval(argument)){.self-link} {#hostgetcodeforeval(argument)}

The Dynamic Code Brand Checks proposal contains an
[implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#hostgetcodeforeval(argument):implementation-defined
x-internal="implementation-defined"}
[HostGetCodeForEval](https://tc39.es/proposal-dynamic-code-brand-checks/#sec-hostgetcodeforeval){#hostgetcodeforeval(argument):js-hostgetcodeforeval
x-internal="js-hostgetcodeforeval"}(`argument`{.variable}) abstract
operation. User agents must use the following implementation:
[\[JSDYNAMICCODEBRANDCHECKS\]](references.html#refsJSDYNAMICCODEBRANDCHECKS)

1.  If `argument`{.variable} is a
    [`TrustedScript`{#hostgetcodeforeval(argument):tt-trustedscript}](https://w3c.github.io/trusted-types/dist/spec/#trusted-script){x-internal="tt-trustedscript"}
    object, then return `argument`{.variable}\'s
    [data](https://w3c.github.io/trusted-types/dist/spec/#trustedscript-data){#hostgetcodeforeval(argument):tt-trustedscript-data
    x-internal="tt-trustedscript-data"}.

2.  Otherwise, return no-code.

##### [8.1.6.4]{.secno} [HostPromiseRejectionTracker]{.dfn}(`promise`{.variable}, `operation`{.variable})[](#the-hostpromiserejectiontracker-implementation){.self-link} {#the-hostpromiserejectiontracker-implementation}

JavaScript contains an
[implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#the-hostpromiserejectiontracker-implementation:implementation-defined
x-internal="implementation-defined"}
[HostPromiseRejectionTracker](https://tc39.es/ecma262/#sec-host-promise-rejection-tracker){#the-hostpromiserejectiontracker-implementation:js-hostpromiserejectiontracker
x-internal="js-hostpromiserejectiontracker"}(`promise`{.variable},
`operation`{.variable}) abstract operation. User agents must use the
following implementation:
[\[JAVASCRIPT\]](references.html#refsJAVASCRIPT)

1.  Let `script`{.variable} be the [running
    script](#running-script){#the-hostpromiserejectiontracker-implementation:running-script}.

2.  If `script`{.variable} is a [classic
    script](#classic-script){#the-hostpromiserejectiontracker-implementation:classic-script}
    and `script`{.variable}\'s [muted
    errors](#muted-errors){#the-hostpromiserejectiontracker-implementation:muted-errors}
    is true, then return.

3.  Let `settingsObject`{.variable} be the [current settings
    object](#current-settings-object){#the-hostpromiserejectiontracker-implementation:current-settings-object}.

4.  If `script`{.variable} is not null, then set
    `settingsObject`{.variable} to `script`{.variable}\'s [settings
    object](#settings-object){#the-hostpromiserejectiontracker-implementation:settings-object}.

5.  Let `global`{.variable} be `settingsObject`{.variable}\'s [global
    object](#concept-settings-object-global){#the-hostpromiserejectiontracker-implementation:concept-settings-object-global}.

6.  If `operation`{.variable} is \"`reject`\", then:

    1.  [Append](https://infra.spec.whatwg.org/#list-append){#the-hostpromiserejectiontracker-implementation:list-append
        x-internal="list-append"} `promise`{.variable} to
        `global`{.variable}\'s [about-to-be-notified rejected promises
        list](#about-to-be-notified-rejected-promises-list){#the-hostpromiserejectiontracker-implementation:about-to-be-notified-rejected-promises-list}.

7.  If `operation`{.variable} is \"`handle`\", then:

    1.  If `global`{.variable}\'s [about-to-be-notified rejected
        promises
        list](#about-to-be-notified-rejected-promises-list){#the-hostpromiserejectiontracker-implementation:about-to-be-notified-rejected-promises-list-2}
        [contains](https://infra.spec.whatwg.org/#list-contain){#the-hostpromiserejectiontracker-implementation:list-contains
        x-internal="list-contains"} `promise`{.variable}, then
        [remove](https://infra.spec.whatwg.org/#list-remove){#the-hostpromiserejectiontracker-implementation:list-remove
        x-internal="list-remove"} `promise`{.variable} from that list
        and return.

    2.  If `global`{.variable}\'s [outstanding rejected promises weak
        set](#outstanding-rejected-promises-weak-set){#the-hostpromiserejectiontracker-implementation:outstanding-rejected-promises-weak-set}
        does not
        [contain](https://infra.spec.whatwg.org/#list-contain){#the-hostpromiserejectiontracker-implementation:list-contains-2
        x-internal="list-contains"} `promise`{.variable}, then return.

    3.  [Remove](https://infra.spec.whatwg.org/#list-remove){#the-hostpromiserejectiontracker-implementation:list-remove-2
        x-internal="list-remove"} `promise`{.variable} from
        `global`{.variable}\'s [outstanding rejected promises weak
        set](#outstanding-rejected-promises-weak-set){#the-hostpromiserejectiontracker-implementation:outstanding-rejected-promises-weak-set-2}.

    4.  [Queue a global
        task](#queue-a-global-task){#the-hostpromiserejectiontracker-implementation:queue-a-global-task}
        on the [DOM manipulation task
        source](#dom-manipulation-task-source){#the-hostpromiserejectiontracker-implementation:dom-manipulation-task-source}
        given `global`{.variable} to [fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#the-hostpromiserejectiontracker-implementation:concept-event-fire
        x-internal="concept-event-fire"} named
        [`rejectionhandled`{#the-hostpromiserejectiontracker-implementation:event-rejectionhandled}](indices.html#event-rejectionhandled)
        at `global`{.variable}, using
        [`PromiseRejectionEvent`{#the-hostpromiserejectiontracker-implementation:promiserejectionevent}](#promiserejectionevent),
        with the
        [`promise`{#the-hostpromiserejectiontracker-implementation:dom-promiserejectionevent-promise}](#dom-promiserejectionevent-promise)
        attribute initialized to `promise`{.variable}, and the
        [`reason`{#the-hostpromiserejectiontracker-implementation:dom-promiserejectionevent-reason}](#dom-promiserejectionevent-reason)
        attribute initialized to
        `promise`{.variable}.\[\[PromiseResult\]\].

##### [8.1.6.5]{.secno} [HostSystemUTCEpochNanoseconds]{.dfn}(`global`{.variable})[](#hostsystemutcepochnanoseconds){.self-link} {#hostsystemutcepochnanoseconds}

The Temporal proposal contains an
[implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#hostsystemutcepochnanoseconds:implementation-defined
x-internal="implementation-defined"}
[HostSystemUTCEpochNanoseconds](https://tc39.es/proposal-temporal/#sec-hostsystemutcepochnanoseconds){#hostsystemutcepochnanoseconds:js-hostsystemutcepochnanoseconds
x-internal="js-hostsystemutcepochnanoseconds"} abstract operation. User
agents must use the following implementation:
[\[JSTEMPORAL\]](references.html#refsJSTEMPORAL)

1.  Let `settingsObject`{.variable} be `global`{.variable}\'s [relevant
    settings
    object](#relevant-settings-object){#hostsystemutcepochnanoseconds:relevant-settings-object}.

2.  Let `time`{.variable} be `settingsObject`{.variable}\'s [current
    wall
    time](https://w3c.github.io/hr-time/#dfn-current-wall-time){#hostsystemutcepochnanoseconds:current-wall-time
    x-internal="current-wall-time"}.

3.  Let `ns`{.variable} be the number of nanoseconds from the [Unix
    epoch](https://w3c.github.io/hr-time/#dfn-unix-epoch){#hostsystemutcepochnanoseconds:unix-epoch
    x-internal="unix-epoch"} to `time`{.variable}, rounded to the
    nearest integer.

4.  Return the result of
    [clamping](https://tc39.es/ecma262/#clamping){#hostsystemutcepochnanoseconds:clamping
    x-internal="clamping"} `ns`{.variable} between
    [nsMinInstant](https://tc39.es/proposal-temporal/#eqn-nsMinInstant){#hostsystemutcepochnanoseconds:nsmininstant
    x-internal="nsmininstant"} and
    [nsMaxInstant](https://tc39.es/proposal-temporal/#eqn-nsMaxInstant){#hostsystemutcepochnanoseconds:nsmaxinstant
    x-internal="nsmaxinstant"}.

##### [8.1.6.6]{.secno} Job-related host hooks[](#integration-with-javascript-jobs){.self-link} {#integration-with-javascript-jobs}

::::: {.mdn-anno .wrapped .before}
**⚠**MDN

:::: feature
[Reference/Global_Objects/Promise#Incumbent_settings_object_tracking](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise#Incumbent_settings_object_tracking "The Promise object represents the eventual completion (or failure) of an asynchronous operation and its resulting value.")

Support in one engine only.

::: support
[Firefox50+]{.firefox .yes}[SafariNo]{.safari .no}[ChromeNo]{.chrome
.no}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[EdgeNo]{.edge_blink .no}

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

The JavaScript specification defines Jobs to be scheduled and run later
by the host, as well as [JobCallback
Records](https://tc39.es/ecma262/#sec-jobcallback-records){#integration-with-javascript-jobs:jobcallback-record
x-internal="jobcallback-record"} which encapsulate JavaScript functions
that are called as part of jobs. The JavaScript specification contains a
number of
[implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#integration-with-javascript-jobs:implementation-defined
x-internal="implementation-defined"} abstract operations that lets the
host define how jobs are scheduled and how JobCallbacks are handled.
HTML uses these abstract operations to [track the [incumbent settings
object](#incumbent-settings-object){#integration-with-javascript-jobs:incumbent-settings-object}
in promises and
[`FinalizationRegistry`{#integration-with-javascript-jobs:finalizationregistry}](https://tc39.es/ecma262/#sec-finalization-registry-objects){x-internal="finalizationregistry"}
callbacks]{#incumbent-settings-object-tracking-in-promises} by saving
and restoring the [incumbent settings
object](#incumbent-settings-object){#integration-with-javascript-jobs:incumbent-settings-object-2}
and a [JavaScript execution
context](https://tc39.es/ecma262/#sec-execution-contexts){#integration-with-javascript-jobs:javascript-execution-context
x-internal="javascript-execution-context"} for the [active
script](#active-script){#integration-with-javascript-jobs:active-script}
in JobCallbacks. This section defines them for user agent hosts.

###### [8.1.6.6.1]{.secno} [HostCallJobCallback]{.dfn}(`callback`{.variable}, `V`{.variable}, `argumentsList`{.variable})[](#hostcalljobcallback){.self-link} {#hostcalljobcallback}

JavaScript contains an
[implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#hostcalljobcallback:implementation-defined
x-internal="implementation-defined"}
[HostCallJobCallback](https://tc39.es/ecma262/#sec-hostcalljobcallback){#hostcalljobcallback:js-hostcalljobcallback
x-internal="js-hostcalljobcallback"}(`callback`{.variable},
`V`{.variable}, `argumentsList`{.variable}) abstract operation to let
hosts restore state when invoking JavaScript callbacks from inside
tasks. User agents must use the following implementation:
[\[JAVASCRIPT\]](references.html#refsJAVASCRIPT)

1.  Let `incumbent settings`{.variable} be
    `callback`{.variable}.\[\[HostDefined\]\].\[\[IncumbentSettings\]\].

2.  Let `script execution context`{.variable} be
    `callback`{.variable}.\[\[HostDefined\]\].\[\[ActiveScriptContext\]\].

3.  [Prepare to run a
    callback](#prepare-to-run-a-callback){#hostcalljobcallback:prepare-to-run-a-callback}
    with `incumbent settings`{.variable}.

    This affects the
    [incumbent](#concept-incumbent-everything){#hostcalljobcallback:concept-incumbent-everything}
    concept while the callback runs.

4.  If `script execution context`{.variable} is not null, then
    [push](https://infra.spec.whatwg.org/#stack-push){#hostcalljobcallback:stack-push
    x-internal="stack-push"} `script execution context`{.variable} onto
    the [JavaScript execution context
    stack](https://tc39.es/ecma262/#execution-context-stack){#hostcalljobcallback:javascript-execution-context-stack
    x-internal="javascript-execution-context-stack"}.

    This affects the [active
    script](#active-script){#hostcalljobcallback:active-script} while
    the callback runs.

5.  Let `result`{.variable} be
    [Call](https://tc39.es/ecma262/#sec-call){#hostcalljobcallback:call
    x-internal="call"}(`callback`{.variable}.\[\[Callback\]\],
    `V`{.variable}, `argumentsList`{.variable}).

6.  If `script execution context`{.variable} is not null, then
    [pop](https://infra.spec.whatwg.org/#stack-pop){#hostcalljobcallback:stack-pop
    x-internal="stack-pop"} `script execution context`{.variable} from
    the [JavaScript execution context
    stack](https://tc39.es/ecma262/#execution-context-stack){#hostcalljobcallback:javascript-execution-context-stack-2
    x-internal="javascript-execution-context-stack"}.

7.  [Clean up after running a
    callback](#clean-up-after-running-a-callback){#hostcalljobcallback:clean-up-after-running-a-callback}
    with `incumbent settings`{.variable}.

8.  Return `result`{.variable}.

###### [8.1.6.6.2]{.secno} [HostEnqueueFinalizationRegistryCleanupJob]{.dfn}(`finalizationRegistry`{.variable})[](#hostenqueuefinalizationregistrycleanupjob){.self-link} {#hostenqueuefinalizationregistrycleanupjob}

JavaScript has the ability to register objects with
[`FinalizationRegistry`{#hostenqueuefinalizationregistrycleanupjob:finalizationregistry}](https://tc39.es/ecma262/#sec-finalization-registry-objects){x-internal="finalizationregistry"}
objects, in order to schedule a cleanup action if they are found to be
garbage collected. The JavaScript specification contains an
[implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#hostenqueuefinalizationregistrycleanupjob:implementation-defined
x-internal="implementation-defined"}
[HostEnqueueFinalizationRegistryCleanupJob](https://tc39.es/ecma262/#sec-host-cleanup-finalization-registry){#hostenqueuefinalizationregistrycleanupjob:js-hostenqueuefinalizationregistrycleanupjob
x-internal="js-hostenqueuefinalizationregistrycleanupjob"}(`finalizationRegistry`{.variable})
abstract operation to schedule the cleanup action.

The timing and occurrence of cleanup work is
[implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#hostenqueuefinalizationregistrycleanupjob:implementation-defined-2
x-internal="implementation-defined"} in the JavaScript specification.
User agents might differ in when and whether an object is garbage
collected, affecting both whether the return value of the
[`WeakRef.prototype.deref()`{#hostenqueuefinalizationregistrycleanupjob:weakref.prototype.deref()}](https://tc39.es/ecma262/#sec-weak-ref.prototype.deref){x-internal="weakref.prototype.deref()"}
method is undefined, and whether
[`FinalizationRegistry`{#hostenqueuefinalizationregistrycleanupjob:finalizationregistry-2}](https://tc39.es/ecma262/#sec-finalization-registry-objects){x-internal="finalizationregistry"}
cleanup callbacks occur. There are well-known cases in popular web
browsers where objects are not accessible to JavaScript, but they remain
retained by the garbage collector indefinitely. HTML clears kept-alive
[`WeakRef`{#hostenqueuefinalizationregistrycleanupjob:weakref}](https://tc39.es/ecma262/#sec-weak-ref-objects){x-internal="weakref"}
objects in the [perform a microtask
checkpoint](#perform-a-microtask-checkpoint){#hostenqueuefinalizationregistrycleanupjob:perform-a-microtask-checkpoint}
algorithm. Authors would be best off not depending on the timing details
of garbage collection implementations.

Cleanup actions do not take place interspersed with synchronous
JavaScript execution, but rather happen in queued
[tasks](#concept-task){#hostenqueuefinalizationregistrycleanupjob:concept-task}.
User agents must use the following implementation:
[\[JAVASCRIPT\]](references.html#refsJAVASCRIPT)

1.  Let `global`{.variable} be
    `finalizationRegistry`{.variable}.\[\[Realm\]\]\'s [global
    object](#global-object){#hostenqueuefinalizationregistrycleanupjob:global-object}.

2.  [Queue a global
    task](#queue-a-global-task){#hostenqueuefinalizationregistrycleanupjob:queue-a-global-task}
    on the [JavaScript engine task
    source]{#javascript-engine-task-source .dfn} given
    `global`{.variable} to perform the following steps:

    1.  Let `entry`{.variable} be
        `finalizationRegistry`{.variable}.\[\[CleanupCallback\]\].\[\[Callback\]\].\[\[Realm\]\]\'s
        [environment settings
        object](#concept-realm-settings-object){#hostenqueuefinalizationregistrycleanupjob:concept-realm-settings-object}.

    2.  [Check if we can run
        script](#check-if-we-can-run-script){#hostenqueuefinalizationregistrycleanupjob:check-if-we-can-run-script}
        with `entry`{.variable}. If this returns \"do not run\", then
        return.

    3.  [Prepare to run
        script](#prepare-to-run-script){#hostenqueuefinalizationregistrycleanupjob:prepare-to-run-script}
        with `entry`{.variable}.

        This affects the
        [entry](#concept-entry-everything){#hostenqueuefinalizationregistrycleanupjob:concept-entry-everything}
        concept while the cleanup callback runs.

    4.  Let `result`{.variable} be the result of performing
        [CleanupFinalizationRegistry](https://tc39.es/ecma262/#sec-cleanup-finalization-registry){#hostenqueuefinalizationregistrycleanupjob:cleanupfinalizationregistry
        x-internal="cleanupfinalizationregistry"}(`finalizationRegistry`{.variable}).

    5.  [Clean up after running
        script](#clean-up-after-running-script){#hostenqueuefinalizationregistrycleanupjob:clean-up-after-running-script}
        with `entry`{.variable}.

    6.  If `result`{.variable} is an [abrupt
        completion](https://tc39.es/ecma262/#sec-completion-record-specification-type){#hostenqueuefinalizationregistrycleanupjob:completion-record
        x-internal="completion-record"}, then [report an
        exception](#report-an-exception){#hostenqueuefinalizationregistrycleanupjob:report-an-exception}
        given by `result`{.variable}.\[\[Value\]\] for
        `global`{.variable}.

###### [8.1.6.6.3]{.secno} [HostEnqueueGenericJob]{.dfn}(`job`{.variable}, `realm`{.variable})[](#hostenqueuegenericjob){.self-link} {#hostenqueuegenericjob}

JavaScript contains an
[implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#hostenqueuegenericjob:implementation-defined
x-internal="implementation-defined"}
[HostEnqueueGenericJob](https://tc39.es/ecma262/#sec-hostenqueuegenericjob){#hostenqueuegenericjob:js-hostenqueuegenericjob
x-internal="js-hostenqueuegenericjob"}(`job`{.variable},
`realm`{.variable}) abstract operation to perform generic jobs in a
particular realm (e.g., resolve promises resulting from
[`Atomics.waitAsync`{#hostenqueuegenericjob:atomics.waitasync}](https://tc39.es/ecma262/#sec-atomics.waitasync){x-internal="atomics.waitasync"}).
User agents must use the following implementation:
[\[JAVASCRIPT\]](references.html#refsJAVASCRIPT)

1.  Let `global`{.variable} be `realm`{.variable}\'s [global
    object](#concept-realm-global){#hostenqueuegenericjob:concept-realm-global}.

2.  [Queue a global
    task](#queue-a-global-task){#hostenqueuegenericjob:queue-a-global-task}
    on the [JavaScript engine task
    source](#javascript-engine-task-source){#hostenqueuegenericjob:javascript-engine-task-source}
    given `global`{.variable} to perform `job`{.variable}().

[]{#integration-with-the-javascript-job-queue}
[]{#enqueuejob(queuename,-job,-arguments)}

###### [8.1.6.6.4]{.secno} [HostEnqueuePromiseJob]{.dfn}(`job`{.variable}, `realm`{.variable})[](#hostenqueuepromisejob){.self-link} {#hostenqueuepromisejob}

JavaScript contains an
[implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#hostenqueuepromisejob:implementation-defined
x-internal="implementation-defined"}
[HostEnqueuePromiseJob](https://tc39.es/ecma262/#sec-hostenqueuepromisejob){#hostenqueuepromisejob:js-hostenqueuepromisejob
x-internal="js-hostenqueuepromisejob"}(`job`{.variable},
`realm`{.variable}) abstract operation to schedule Promise-related
operations. HTML schedules these operations in the microtask queue. User
agents must use the following implementation:
[\[JAVASCRIPT\]](references.html#refsJAVASCRIPT)

1.  If `realm`{.variable} is not null, then let
    `job settings`{.variable} be the [settings
    object](#concept-realm-settings-object){#hostenqueuepromisejob:concept-realm-settings-object}
    for `realm`{.variable}. Otherwise, let `job settings`{.variable} be
    null.

    ::: note
    If `realm`{.variable} is not null, it is the
    [realm](https://tc39.es/ecma262/#sec-code-realms){#hostenqueuepromisejob:realm
    x-internal="realm"} of the author code that will run. When
    `job`{.variable} is returned by
    [NewPromiseReactionJob](https://tc39.es/ecma262/#sec-newpromisereactionjob){#hostenqueuepromisejob:newpromisereactionjob
    x-internal="newpromisereactionjob"}, it is the realm of the
    promise\'s handler function. When `job`{.variable} is returned by
    [NewPromiseResolveThenableJob](https://tc39.es/ecma262/#sec-newpromiseresolvethenablejob){#hostenqueuepromisejob:newpromiseresolvethenablejob
    x-internal="newpromiseresolvethenablejob"}, it is the realm of the
    `then` function.

    If `realm`{.variable} is null, either no author code will run or
    author code is guaranteed to throw. For the former, the author may
    not have passed in code to run, such as in
    `promise.then(null, null)`. For the latter, it is because a revoked
    Proxy was passed. In both cases, all the steps below that would
    otherwise use `job settings`{.variable} get skipped.

    [NewPromiseResolveThenableJob](https://tc39.es/ecma262/#sec-newpromiseresolvethenablejob)
    and
    [NewPromiseReactionJob](https://tc39.es/ecma262/#sec-newpromisereactionjob)
    both seem to provide non-null realms (the current Realm Record) in
    the case of a revoked proxy. The previous text could be updated to
    reflect that.
    :::

2.  [Queue a
    microtask](#queue-a-microtask){#hostenqueuepromisejob:queue-a-microtask}
    to perform the following steps:

    1.  If `job settings`{.variable} is not null, then [check if we can
        run
        script](#check-if-we-can-run-script){#hostenqueuepromisejob:check-if-we-can-run-script}
        with `job settings`{.variable}. If this returns \"do not run\"
        then return.

    2.  If `job settings`{.variable} is not null, then [prepare to run
        script](#prepare-to-run-script){#hostenqueuepromisejob:prepare-to-run-script}
        with `job settings`{.variable}.

        This affects the
        [entry](#concept-entry-everything){#hostenqueuepromisejob:concept-entry-everything}
        concept while the job runs.

    3.  Let `result`{.variable} be `job`{.variable}().

        `job`{.variable} is an [abstract
        closure](https://tc39.es/ecma262/#sec-abstract-closure){#hostenqueuepromisejob:abstract-closure
        x-internal="abstract-closure"} returned by
        [NewPromiseReactionJob](https://tc39.es/ecma262/#sec-newpromisereactionjob){#hostenqueuepromisejob:newpromisereactionjob-2
        x-internal="newpromisereactionjob"} or
        [NewPromiseResolveThenableJob](https://tc39.es/ecma262/#sec-newpromiseresolvethenablejob){#hostenqueuepromisejob:newpromiseresolvethenablejob-2
        x-internal="newpromiseresolvethenablejob"}. The promise\'s
        handler function when `job`{.variable} is returned by
        [NewPromiseReactionJob](https://tc39.es/ecma262/#sec-newpromisereactionjob){#hostenqueuepromisejob:newpromisereactionjob-3
        x-internal="newpromisereactionjob"}, and the `then` function
        when `job`{.variable} is returned by
        [NewPromiseResolveThenableJob](https://tc39.es/ecma262/#sec-newpromiseresolvethenablejob){#hostenqueuepromisejob:newpromiseresolvethenablejob-3
        x-internal="newpromiseresolvethenablejob"}, are wrapped in
        [JobCallback
        Records](https://tc39.es/ecma262/#sec-jobcallback-records){#hostenqueuepromisejob:jobcallback-record
        x-internal="jobcallback-record"}. HTML saves the [incumbent
        settings
        object](#incumbent-settings-object){#hostenqueuepromisejob:incumbent-settings-object}
        and a [JavaScript execution
        context](https://tc39.es/ecma262/#sec-execution-contexts){#hostenqueuepromisejob:javascript-execution-context
        x-internal="javascript-execution-context"} for to the [active
        script](#active-script){#hostenqueuepromisejob:active-script} in
        [HostMakeJobCallback](#hostmakejobcallback){#hostenqueuepromisejob:hostmakejobcallback}
        and restores them in
        [HostCallJobCallback](#hostcalljobcallback){#hostenqueuepromisejob:hostcalljobcallback}.

    4.  If `job settings`{.variable} is not null, then [clean up after
        running
        script](#clean-up-after-running-script){#hostenqueuepromisejob:clean-up-after-running-script}
        with `job settings`{.variable}.

    5.  If `result`{.variable} is an [abrupt
        completion](https://tc39.es/ecma262/#sec-completion-record-specification-type){#hostenqueuepromisejob:completion-record
        x-internal="completion-record"}, then [report an
        exception](#report-an-exception){#hostenqueuepromisejob:report-an-exception}
        given by `result`{.variable}.\[\[Value\]\] for
        `realm`{.variable}\'s [global
        object](#concept-realm-global){#hostenqueuepromisejob:concept-realm-global}.

        There is a very gnarly case where HostEnqueuePromiseJob is
        called with a null realm (e.g., because Promise.prototype.then
        was called with null handlers) but also the job returns abruptly
        (because the promise capability\'s resolve or reject handler
        threw, possibly because this is a subclass of Promise that takes
        the supplied functions and wraps them in throwing functions
        before passing them on to the function passed to the Promise
        superclass constructor. Which global is to be used then,
        considering that the current realm could be different at each of
        those steps, by using a Promise constructor or
        Promise.prototype.then from another realm? See [issue
        #10526](https://github.com/whatwg/html/issues/10526).

###### [8.1.6.6.5]{.secno} [HostEnqueueTimeoutJob]{.dfn}(`job`{.variable}, `realm`{.variable}, `milliseconds`{.variable})[](#hostenqueuetimeoutjob){.self-link} {#hostenqueuetimeoutjob}

JavaScript contains an
[implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#hostenqueuetimeoutjob:implementation-defined
x-internal="implementation-defined"}
[HostEnqueueTimeoutJob](https://tc39.es/ecma262/#sec-hostenqueuetimeoutjob){#hostenqueuetimeoutjob:js-hostenqueuetimeoutjob
x-internal="js-hostenqueuetimeoutjob"}(`job`{.variable},
`milliseconds`{.variable}) abstract operation to schedule an operation
to be performed after a timeout. HTML schedules these operations using
[run steps after a
timeout](timers-and-user-prompts.html#run-steps-after-a-timeout){#hostenqueuetimeoutjob:run-steps-after-a-timeout}.
User agents must use the following implementation:
[\[JAVASCRIPT\]](references.html#refsJAVASCRIPT)

1.  Let `global`{.variable} be `realm`{.variable}\'s [global
    object](#concept-realm-global){#hostenqueuetimeoutjob:concept-realm-global}.

2.  Let `timeoutStep`{.variable} be an algorithm step which [queues a
    global
    task](#queue-a-global-task){#hostenqueuetimeoutjob:queue-a-global-task}
    on the [JavaScript engine task
    source](#javascript-engine-task-source){#hostenqueuetimeoutjob:javascript-engine-task-source}
    given `global`{.variable} to perform `job`{.variable}().

3.  [Run steps after a
    timeout](timers-and-user-prompts.html#run-steps-after-a-timeout){#hostenqueuetimeoutjob:run-steps-after-a-timeout-2}
    given `global`{.variable}, \"`JavaScript`\",
    `milliseconds`{.variable}, and `timeoutStep`{.variable}.

###### [8.1.6.6.6]{.secno} [HostMakeJobCallback]{.dfn}(`callable`{.variable})[](#hostmakejobcallback){.self-link} {#hostmakejobcallback}

JavaScript contains an
[implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#hostmakejobcallback:implementation-defined
x-internal="implementation-defined"}
[HostMakeJobCallback](https://tc39.es/ecma262/#sec-hostmakejobcallback){#hostmakejobcallback:js-hostmakejobcallback
x-internal="js-hostmakejobcallback"}(`callable`{.variable}) abstract
operation to let hosts attach state to JavaScript callbacks that are
called from inside
[task](#concept-task){#hostmakejobcallback:concept-task}s. User agents
must use the following implementation:
[\[JAVASCRIPT\]](references.html#refsJAVASCRIPT)

1.  Let `incumbent settings`{.variable} be the [incumbent settings
    object](#incumbent-settings-object){#hostmakejobcallback:incumbent-settings-object}.

2.  Let `active script`{.variable} be the [active
    script](#active-script){#hostmakejobcallback:active-script}.

3.  Let `script execution context`{.variable} be null.

4.  If `active script`{.variable} is not null, set
    `script execution context`{.variable} to a new [JavaScript execution
    context](https://tc39.es/ecma262/#sec-execution-contexts){#hostmakejobcallback:javascript-execution-context
    x-internal="javascript-execution-context"}, with its Function field
    set to null, its Realm field set to `active script`{.variable}\'s
    [settings
    object](#settings-object){#hostmakejobcallback:settings-object}\'s
    [realm](#environment-settings-object's-realm){#hostmakejobcallback:environment-settings-object's-realm},
    and its ScriptOrModule set to `active script`{.variable}\'s
    [record](#concept-script-record){#hostmakejobcallback:concept-script-record}.

    As seen below, this is used in order to propagate the current
    [active
    script](#active-script){#hostmakejobcallback:active-script-2}
    forward to the time when the job callback is invoked.

    ::: example
    A case where `active script`{.variable} is non-null, and saving it
    in this way is useful, is the following:

    ``` js
    Promise.resolve('import(`./example.mjs`)').then(eval);
    ```

    Without this step (and the steps that use it in
    [HostCallJobCallback](#hostcalljobcallback){#hostmakejobcallback:hostcalljobcallback}),
    there would be no [active
    script](#active-script){#hostmakejobcallback:active-script-3} when
    the
    [`import()`{#hostmakejobcallback:import()}](https://tc39.es/ecma262/#sec-import-calls){x-internal="import()"}
    expression is evaluated, since
    [`eval()`{#hostmakejobcallback:eval()}](https://tc39.es/ecma262/#sec-eval-x){x-internal="eval()"}
    is a built-in function that does not originate from any particular
    [script](#concept-script){#hostmakejobcallback:concept-script}.

    With this step in place, the active script is propagated from the
    above code into the job, allowing
    [`import()`{#hostmakejobcallback:import()-2}](https://tc39.es/ecma262/#sec-import-calls){x-internal="import()"}
    to use the original script\'s [base
    URL](#concept-script-base-url){#hostmakejobcallback:concept-script-base-url}
    appropriately.
    :::

    ::: example
    `active script`{.variable} can be null if the user clicks on the
    following button:

    ``` html
    <button onclick="Promise.resolve('import(`./example.mjs`)').then(eval)">Click me</button>
    ```

    In this case, the JavaScript function for the [event
    handler](#event-handlers){#hostmakejobcallback:event-handlers} will
    be created by the [get the current value of the event
    handler](#getting-the-current-value-of-the-event-handler){#hostmakejobcallback:getting-the-current-value-of-the-event-handler}
    algorithm, which creates a function with null \[\[ScriptOrModule\]\]
    value. Thus, when the promise machinery calls
    [HostMakeJobCallback](#hostmakejobcallback){#hostmakejobcallback:hostmakejobcallback},
    there will be no [active
    script](#active-script){#hostmakejobcallback:active-script-4} to
    pass along.

    As a consequence, this means that when the
    [`import()`{#hostmakejobcallback:import()-3}](https://tc39.es/ecma262/#sec-import-calls){x-internal="import()"}
    expression is evaluated, there will still be no [active
    script](#active-script){#hostmakejobcallback:active-script-5}.
    Fortunately that is handled by our implementation of
    [HostLoadImportedModule](#hostloadimportedmodule){#hostmakejobcallback:hostloadimportedmodule}
    by falling back to using the [current settings
    object](#current-settings-object){#hostmakejobcallback:current-settings-object}\'s
    [API base URL](#api-base-url){#hostmakejobcallback:api-base-url}.
    :::

5.  Return the [JobCallback
    Record](https://tc39.es/ecma262/#sec-jobcallback-records){#hostmakejobcallback:jobcallback-record
    x-internal="jobcallback-record"} { \[\[Callback\]\]:
    `callable`{.variable}, \[\[HostDefined\]\]: {
    \[\[IncumbentSettings\]\]: `incumbent settings`{.variable},
    \[\[ActiveScriptContext\]\]: `script execution context`{.variable} }
    }.

##### [8.1.6.7]{.secno} Module-related host hooks[](#integration-with-the-javascript-module-system){.self-link} {#integration-with-the-javascript-module-system}

The JavaScript specification defines a syntax for modules, as well as
some host-agnostic parts of their processing model. This specification
defines the rest of their processing model: how the module system is
bootstrapped, via the
[`script`{#integration-with-the-javascript-module-system:the-script-element}](scripting.html#the-script-element)
element with
[`type`{#integration-with-the-javascript-module-system:attr-script-type}](scripting.html#attr-script-type)
attribute set to \"`module`\", and how modules are fetched, resolved,
and executed. [\[JAVASCRIPT\]](references.html#refsJAVASCRIPT)

Although the JavaScript specification speaks in terms of \"scripts\"
versus \"modules\", in general this specification speaks in terms of
[classic
scripts](#classic-script){#integration-with-the-javascript-module-system:classic-script}
versus [module
scripts](#module-script){#integration-with-the-javascript-module-system:module-script},
since both of them use the
[`script`{#integration-with-the-javascript-module-system:the-script-element-2}](scripting.html#the-script-element)
element.

`modulePromise`{.variable}` = `[`import(``specifier`{.variable}`)`](https://tc39.es/ecma262/#sec-import-calls){#integration-with-the-javascript-module-system:import() x-internal="import()"}

:   Returns a promise for the module namespace object for the [module
    script](#module-script){#integration-with-the-javascript-module-system:module-script-2}
    identified by `specifier`{.variable}. This allows dynamic importing
    of module scripts at runtime, instead of statically using the
    `import` statement form. The specifier will be
    [resolved](#resolve-a-module-specifier){#integration-with-the-javascript-module-system:resolve-a-module-specifier}
    relative to the [active
    script](#active-script){#integration-with-the-javascript-module-system:active-script}.

    The returned promise will be rejected if an invalid specifier is
    given, or if a failure is encountered while
    [fetching](#hostloadimportedmodule){#integration-with-the-javascript-module-system:hostloadimportedmodule}
    or evaluating the resulting module graph.

    This syntax can be used inside both
    [classic](#classic-script){#integration-with-the-javascript-module-system:classic-script-2}
    and [module
    scripts](#module-script){#integration-with-the-javascript-module-system:module-script-3}.
    It thus provides a bridge into the module-script world, from the
    classic-script world.

`url`{.variable}` = `[`import.meta`](https://tc39.es/ecma262/#sec-meta-properties){#integration-with-the-javascript-module-system:import.meta x-internal="import.meta"}`.`[`url`](#import-meta-url){#integration-with-the-javascript-module-system:import-meta-url}

:   Returns the [active module
    script](#active-script){#integration-with-the-javascript-module-system:active-script-2}\'s
    [base
    URL](#concept-script-base-url){#integration-with-the-javascript-module-system:concept-script-base-url}.

    This syntax can only be used inside [module
    scripts](#module-script){#integration-with-the-javascript-module-system:module-script-4}.

`url`{.variable}` = `[`import.meta`](https://tc39.es/ecma262/#sec-meta-properties){#integration-with-the-javascript-module-system:import.meta-2 x-internal="import.meta"}`.`[`resolve`](#import-meta-resolve){#integration-with-the-javascript-module-system:import-meta-resolve}`(``specifier`{.variable}`)`

:   Returns `specifier`{.variable},
    [resolved](#resolve-a-module-specifier){#integration-with-the-javascript-module-system:resolve-a-module-specifier-2}
    relative to the [active
    script](#active-script){#integration-with-the-javascript-module-system:active-script-3}.
    That is, this returns the URL that would be imported by using
    [`import(`{#integration-with-the-javascript-module-system:import()-2}`specifier`{.variable}`)`{#integration-with-the-javascript-module-system:import()-2}](https://tc39.es/ecma262/#sec-import-calls){x-internal="import()"}.

    Throws a
    [`TypeError`{#integration-with-the-javascript-module-system:typeerror}](https://tc39.es/ecma262/#sec-native-error-types-used-in-this-standard-typeerror){x-internal="typeerror"}
    exception if an invalid specifier is given.

    This syntax can only be used inside [module
    scripts](#module-script){#integration-with-the-javascript-module-system:module-script-5}.

A [module map]{#module-map .dfn} is a
[map](https://infra.spec.whatwg.org/#ordered-map){#integration-with-the-javascript-module-system:ordered-map
x-internal="ordered-map"} keyed by
[tuples](https://infra.spec.whatwg.org/#tuple){#integration-with-the-javascript-module-system:tuple
x-internal="tuple"} consisting of a [URL
record](https://url.spec.whatwg.org/#concept-url){#integration-with-the-javascript-module-system:url-record
x-internal="url-record"} and a
[string](https://infra.spec.whatwg.org/#string){#integration-with-the-javascript-module-system:string
x-internal="string"}. The [URL
record](https://url.spec.whatwg.org/#concept-url){#integration-with-the-javascript-module-system:url-record-2
x-internal="url-record"} is the [request
URL](https://fetch.spec.whatwg.org/#concept-request-url){#integration-with-the-javascript-module-system:concept-request-url
x-internal="concept-request-url"} at which the module was fetched, and
the
[string](https://infra.spec.whatwg.org/#string){#integration-with-the-javascript-module-system:string-2
x-internal="string"} indicates the type of the module (e.g.
\"`javascript-or-wasm`\"). The [module
map](#module-map){#integration-with-the-javascript-module-system:module-map}\'s
values are either a [module
script](#module-script){#integration-with-the-javascript-module-system:module-script-6},
null (used to represent failed fetches), or a placeholder value
\"`fetching`\". [Module
maps](#module-map){#integration-with-the-javascript-module-system:module-map-2}
are used to ensure that imported module scripts are only fetched,
parsed, and evaluated once per
[`Document`{#integration-with-the-javascript-module-system:document}](dom.html#document)
or [worker](workers.html#workers).

::: example
Since [module
maps](#module-map){#integration-with-the-javascript-module-system:module-map-3}
are keyed by (URL, module type), the following code will create three
separate entries in the [module
map](#module-map){#integration-with-the-javascript-module-system:module-map-4},
since it results in three different (URL, module type)
[tuples](https://infra.spec.whatwg.org/#tuple){#integration-with-the-javascript-module-system:tuple-2
x-internal="tuple"} (all with \"`javascript-or-wasm`\" type):

``` js
import "https://example.com/module.mjs";
import "https://example.com/module.mjs#map-buster";
import "https://example.com/module.mjs?debug=true";
```

That is, URL
[queries](https://url.spec.whatwg.org/#concept-url-query){#integration-with-the-javascript-module-system:concept-url-query
x-internal="concept-url-query"} and
[fragments](https://url.spec.whatwg.org/#concept-url-fragment){#integration-with-the-javascript-module-system:concept-url-fragment
x-internal="concept-url-fragment"} can be varied to create distinct
entries in the [module
map](#module-map){#integration-with-the-javascript-module-system:module-map-5};
they are not ignored. Thus, three separate fetches and three separate
module evaluations will be performed.

In contrast, the following code would only create a single entry in the
[module
map](#module-map){#integration-with-the-javascript-module-system:module-map-6},
since after applying the [URL
parser](https://url.spec.whatwg.org/#concept-url-parser){#integration-with-the-javascript-module-system:url-parser
x-internal="url-parser"} to these inputs, the resulting [URL
records](https://url.spec.whatwg.org/#concept-url){#integration-with-the-javascript-module-system:url-record-3
x-internal="url-record"} are equal:

``` js
import "https://example.com/module2.mjs";
import "https:example.com/module2.mjs";
import "https://///example.com\\module2.mjs";
import "https://example.com/foo/../module2.mjs";
```

So in this second example, only one fetch and one module evaluation will
occur.

Note that this behavior is the same as how [shared
workers](workers.html#sharedworker){#integration-with-the-javascript-module-system:sharedworker}
are keyed by their parsed [constructor
url](workers.html#concept-sharedworkerglobalscope-constructor-url){#integration-with-the-javascript-module-system:concept-sharedworkerglobalscope-constructor-url}.
:::

::: example
Since module type is also part of the [module
map](#module-map){#integration-with-the-javascript-module-system:module-map-7}
key, the following code will create two separate entries in the [module
map](#module-map){#integration-with-the-javascript-module-system:module-map-8}
(the type is \"`javascript-or-wasm`\" for the first, and \"`css`\" for
the second):

``` html
<script type=module>
  import "https://example.com/module";
</script>
<script type=module>
  import "https://example.com/module" with { type: "css" };
</script>
```

This can result in two separate fetches and two separate module
evaluations being performed.

In practice, due to the as-yet-unspecified memory cache (see issue
[#6110](https://github.com/whatwg/html/issues/6110)) the resource may
only be fetched once in WebKit and Blink-based browsers. Additionally,
as long as all module types are mutually exclusive, the module type
check in [fetch a single module
script](#fetch-a-single-module-script){#integration-with-the-javascript-module-system:fetch-a-single-module-script}
will fail for at least one of the imports, so at most one module
evaluation will occur.

The purpose of including the type in the [module
map](#module-map){#integration-with-the-javascript-module-system:module-map-9}
key is so that an import with the wrong type attribute does not prevent
a different import of the same specifier but with the correct type from
succeeding.
:::

::: example
JavaScript module scripts are the default import type when importing
from another JavaScript module; that is, when an `import` statement
lacks a `type` import attribute the imported module script\'s type will
be JavaScript. Attempting to import a JavaScript resource using an
`import` statement with a `type` import attribute will fail:

``` html
<script type="module">
    // All of the following will fail, assuming that the imported .mjs files are served with a
    // JavaScript MIME type. JavaScript module scripts are the default and cannot be imported with
    // any import type attribute.
    import foo from "./foo.mjs" with { type: "javascript" };
    import foo2 from "./foo2.mjs" with { type: "js" };
    import foo3 from "./foo3.mjs" with { type: "" };
    await import("./foo4.mjs", { with: { type: null } });
    await import("./foo5.mjs", { with: { type: undefined } });
</script>
```
:::

###### [8.1.6.7.1]{.secno} [HostGetImportMetaProperties]{.dfn}(`moduleRecord`{.variable})[](#hostgetimportmetaproperties){.self-link} {#hostgetimportmetaproperties}

::::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Reference/Operators/import.meta/resolve](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/import.meta/resolve "import.meta.resolve() is a built-in function defined on the import.meta object of a JavaScript module that resolves a module specifier to a URL using the current module's URL as base.")

Support in all current engines.

::: support
[Firefox106+]{.firefox .yes}[Safari16.4+]{.safari
.yes}[Chrome105+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge105+]{.edge_blink .yes}

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

:::: feature
[Reference/Operators/import.meta](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/import.meta "The import.meta meta-property exposes context-specific metadata to a JavaScript module. It contains information about the module, such as the module's URL.")

Support in all current engines.

::: support
[Firefox62+]{.firefox .yes}[Safari11.1+]{.safari
.yes}[Chrome64+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS12+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::::

JavaScript contains an
[implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#hostgetimportmetaproperties:implementation-defined
x-internal="implementation-defined"}
[HostGetImportMetaProperties](https://tc39.es/ecma262/#sec-hostgetimportmetaproperties){#hostgetimportmetaproperties:js-hostgetimportmetaproperties
x-internal="js-hostgetimportmetaproperties"} abstract operation. User
agents must use the following implementation:
[\[JAVASCRIPT\]](references.html#refsJAVASCRIPT)

1.  Let `moduleScript`{.variable} be
    `moduleRecord`{.variable}.\[\[HostDefined\]\].

2.  [Assert](https://infra.spec.whatwg.org/#assert){#hostgetimportmetaproperties:assert
    x-internal="assert"}: `moduleScript`{.variable}\'s [base
    URL](#concept-script-base-url){#hostgetimportmetaproperties:concept-script-base-url}
    is not null, as `moduleScript`{.variable} is a [JavaScript module
    script](#javascript-module-script){#hostgetimportmetaproperties:javascript-module-script}.

3.  Let `urlString`{.variable} be `moduleScript`{.variable}\'s [base
    URL](#concept-script-base-url){#hostgetimportmetaproperties:concept-script-base-url-2},
    [serialized](https://url.spec.whatwg.org/#concept-url-serializer){#hostgetimportmetaproperties:concept-url-serializer
    x-internal="concept-url-serializer"}.

4.  Let `steps`{.variable} be the following steps, given the argument
    `specifier`{.variable}:

    1.  Set `specifier`{.variable} to ?
        [ToString](https://tc39.es/ecma262/#sec-tostring){#hostgetimportmetaproperties:tostring
        x-internal="tostring"}(`specifier`{.variable}).

    2.  Let `url`{.variable} be the result of [resolving a module
        specifier](#resolve-a-module-specifier){#hostgetimportmetaproperties:resolve-a-module-specifier}
        given `moduleScript`{.variable} and `specifier`{.variable}.

    3.  Return the
        [serialization](https://url.spec.whatwg.org/#concept-url-serializer){#hostgetimportmetaproperties:concept-url-serializer-2
        x-internal="concept-url-serializer"} of `url`{.variable}.

5.  Let `resolveFunction`{.variable} be !
    [CreateBuiltinFunction](https://tc39.es/ecma262/#sec-createbuiltinfunction){#hostgetimportmetaproperties:createbuiltinfunction
    x-internal="createbuiltinfunction"}(`steps`{.variable}, 1,
    \"`resolve`\", « »).

6.  Return «
    [Record](https://tc39.es/ecma262/#sec-list-and-record-specification-type){#hostgetimportmetaproperties:record
    x-internal="record"} { \[\[Key\]\]: \"[`url`]{#import-meta-url
    .dfn}\", \[\[Value\]\]: `urlString`{.variable} },
    [Record](https://tc39.es/ecma262/#sec-list-and-record-specification-type){#hostgetimportmetaproperties:record-2
    x-internal="record"} { \[\[Key\]\]:
    \"[`resolve`]{#import-meta-resolve .dfn}\", \[\[Value\]\]:
    `resolveFunction`{.variable} } ».

###### [8.1.6.7.2]{.secno} []{#hostgetsupportedimportassertions}[HostGetSupportedImportAttributes]{.dfn}()[](#hostgetsupportedimportattributes){.self-link}

JavaScript contains an an
[implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#hostgetsupportedimportattributes:implementation-defined
x-internal="implementation-defined"}
[HostGetSupportedImportAttributes](https://tc39.es/ecma262/#sec-hostgetsupportedimportattributes){#hostgetsupportedimportattributes:js-hostgetsupportedimportattributes
x-internal="js-hostgetsupportedimportattributes"} abstract operation.
User agents must use the following implementation:
[\[JAVASCRIPT\]](references.html#refsJAVASCRIPT)

1.  Return « \"`type`\" ».

###### [8.1.6.7.3]{.secno} [HostLoadImportedModule]{.dfn}(`referrer`{.variable}, `moduleRequest`{.variable}, `loadState`{.variable}, `payload`{.variable})[](#hostloadimportedmodule){.self-link} {#hostloadimportedmodule}

JavaScript contains an
[implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#hostloadimportedmodule:implementation-defined
x-internal="implementation-defined"}
[HostLoadImportedModule](https://tc39.es/ecma262/#sec-HostLoadImportedModule){#hostloadimportedmodule:js-hostloadimportedmodule
x-internal="js-hostloadimportedmodule"} abstract operation. User agents
must use the following implementation:
[\[JAVASCRIPT\]](references.html#refsJAVASCRIPT)

1.  Let `settingsObject`{.variable} be the [current settings
    object](#current-settings-object){#hostloadimportedmodule:current-settings-object}.

2.  If `settingsObject`{.variable}\'s [global
    object](#concept-settings-object-global){#hostloadimportedmodule:concept-settings-object-global}
    implements
    [`WorkletGlobalScope`{#hostloadimportedmodule:workletglobalscope}](worklets.html#workletglobalscope)
    or
    [`ServiceWorkerGlobalScope`{#hostloadimportedmodule:serviceworkerglobalscope}](https://w3c.github.io/ServiceWorker/#serviceworkerglobalscope){x-internal="serviceworkerglobalscope"}
    and `loadState`{.variable} is undefined, then:

    `loadState`{.variable} is undefined when the current fetching
    process has been initiated by a dynamic
    [`import()`{#hostloadimportedmodule:import()}](https://tc39.es/ecma262/#sec-import-calls){x-internal="import()"}
    call, either directly or when loading the transitive dependencies of
    the dynamically imported module.

    1.  Let `completion`{.variable} be [Completion
        Record](https://tc39.es/ecma262/#sec-completion-record-specification-type){#hostloadimportedmodule:completion-record
        x-internal="completion-record"} { \[\[Type\]\]: throw,
        \[\[Value\]\]: a new
        [`TypeError`{#hostloadimportedmodule:typeerror}](https://tc39.es/ecma262/#sec-native-error-types-used-in-this-standard-typeerror){x-internal="typeerror"},
        \[\[Target\]\]: empty }.

    2.  Perform
        [FinishLoadingImportedModule](https://tc39.es/ecma262/#sec-FinishLoadingImportedModule){#hostloadimportedmodule:finishloadingimportedmodule
        x-internal="finishloadingimportedmodule"}(`referrer`{.variable},
        `moduleRequest`{.variable}, `payload`{.variable},
        `completion`{.variable}).

    3.  Return.

3.  Let `referencingScript`{.variable} be null.

4.  Let `originalFetchOptions`{.variable} be the [default script fetch
    options](#default-script-fetch-options){#hostloadimportedmodule:default-script-fetch-options}.

5.  Let `fetchReferrer`{.variable} be \"`client`\".

6.  If `referrer`{.variable} is a [Script
    Record](https://tc39.es/ecma262/#sec-script-records){#hostloadimportedmodule:script-record
    x-internal="script-record"} or a [Cyclic Module
    Record](https://tc39.es/ecma262/#sec-cyclic-module-records){#hostloadimportedmodule:cyclic-module-record
    x-internal="cyclic-module-record"}, then:

    1.  Set `referencingScript`{.variable} to
        `referrer`{.variable}.\[\[HostDefined\]\].

    2.  Set `settingsObject`{.variable} to
        `referencingScript`{.variable}\'s [settings
        object](#settings-object){#hostloadimportedmodule:settings-object}.

    3.  Set `fetchReferrer`{.variable} to
        `referencingScript`{.variable}\'s [base
        URL](#concept-script-base-url){#hostloadimportedmodule:concept-script-base-url}.

    4.  Set `originalFetchOptions`{.variable} to
        `referencingScript`{.variable}\'s [fetch
        options](#concept-script-script-fetch-options){#hostloadimportedmodule:concept-script-script-fetch-options}.

    ::: example
    `referrer`{.variable} is usually a [Script
    Record](https://tc39.es/ecma262/#sec-script-records){#hostloadimportedmodule:script-record-2
    x-internal="script-record"} or a [Cyclic Module
    Record](https://tc39.es/ecma262/#sec-cyclic-module-records){#hostloadimportedmodule:cyclic-module-record-2
    x-internal="cyclic-module-record"}, but it will not be so for event
    handlers per the [get the current value of the event
    handler](#getting-the-current-value-of-the-event-handler){#hostloadimportedmodule:getting-the-current-value-of-the-event-handler}
    algorithm. For example, given:

    ``` html
    <button onclick="import('./foo.mjs')">Click me</button>
    ```

    If a
    [`click`{#hostloadimportedmodule:event-click}](https://w3c.github.io/uievents/#event-type-click){x-internal="event-click"}
    event occurs, then at the time the
    [`import()`{#hostloadimportedmodule:import()-2}](https://tc39.es/ecma262/#sec-import-calls){x-internal="import()"}
    expression runs,
    [GetActiveScriptOrModule](https://tc39.es/ecma262/#sec-getactivescriptormodule){#hostloadimportedmodule:getactivescriptormodule
    x-internal="getactivescriptormodule"} will return null, and this
    operation will receive the [current
    realm](https://tc39.es/ecma262/#current-realm){#hostloadimportedmodule:current-realm
    x-internal="current-realm"} as a fallback `referrer`{.variable}.
    :::

7.  If `referrer`{.variable} is a [Cyclic Module
    Record](https://tc39.es/ecma262/#sec-cyclic-module-records){#hostloadimportedmodule:cyclic-module-record-3
    x-internal="cyclic-module-record"} and `moduleRequest`{.variable} is
    equal to the first element of
    `referrer`{.variable}.\[\[RequestedModules\]\], then:

    1.  ::: {#validate-requested-module-specifiers}
        [For
        each](https://infra.spec.whatwg.org/#list-iterate){#hostloadimportedmodule:list-iterate
        x-internal="list-iterate"} [ModuleRequest
        record](https://tc39.es/ecma262/#modulerequest-record){#hostloadimportedmodule:modulerequest-record
        x-internal="modulerequest-record"} `requested`{.variable} of
        `referrer`{.variable}.\[\[RequestedModules\]\]:

        1.  If `moduleRequest`{.variable}.\[\[Attributes\]\] contains a
            [Record](https://tc39.es/ecma262/#sec-list-and-record-specification-type){#hostloadimportedmodule:record
            x-internal="record"} `entry`{.variable} such that
            `entry`{.variable}.\[\[Key\]\] is not \"`type`\", then:

            1.  Let `completion`{.variable} be [Completion
                Record](https://tc39.es/ecma262/#sec-completion-record-specification-type){#hostloadimportedmodule:completion-record-2
                x-internal="completion-record"} { \[\[Type\]\]: throw,
                \[\[Value\]\]: a new
                [`SyntaxError`{#hostloadimportedmodule:syntaxerror-2}](https://tc39.es/ecma262/#sec-native-error-types-used-in-this-standard-syntaxerror){x-internal="syntaxerror-2"}
                exception, \[\[Target\]\]: empty }.

            2.  Perform
                [FinishLoadingImportedModule](https://tc39.es/ecma262/#sec-FinishLoadingImportedModule){#hostloadimportedmodule:finishloadingimportedmodule-2
                x-internal="finishloadingimportedmodule"}(`referrer`{.variable},
                `moduleRequest`{.variable}, `payload`{.variable},
                `completion`{.variable}).

            3.  Return.

            The JavaScript specification re-performs this validation but
            it is duplicated here to avoid unnecessarily loading any of
            the dependencies on validation failure.

        2.  [Resolve a module
            specifier](#resolve-a-module-specifier){#hostloadimportedmodule:resolve-a-module-specifier}
            given `referencingScript`{.variable} and
            `moduleRequest`{.variable}.\[\[Specifier\]\], catching any
            exceptions. If they throw an exception, let
            `resolutionError`{.variable} be the thrown exception.

        3.  If the previous step threw an exception, then:

            1.  Let `completion`{.variable} be [Completion
                Record](https://tc39.es/ecma262/#sec-completion-record-specification-type){#hostloadimportedmodule:completion-record-3
                x-internal="completion-record"} { \[\[Type\]\]: throw,
                \[\[Value\]\]: `resolutionError`{.variable},
                \[\[Target\]\]: empty }.

            2.  Perform
                [FinishLoadingImportedModule](https://tc39.es/ecma262/#sec-FinishLoadingImportedModule){#hostloadimportedmodule:finishloadingimportedmodule-3
                x-internal="finishloadingimportedmodule"}(`referrer`{.variable},
                `moduleRequest`{.variable}, `payload`{.variable},
                `completion`{.variable}).

            3.  Return.

        4.  Let `moduleType`{.variable} be the result of running the
            [module type from module
            request](#module-type-from-module-request){#hostloadimportedmodule:module-type-from-module-request}
            steps given `moduleRequest`{.variable}.

        5.  If the result of running the [module type
            allowed](#module-type-allowed){#hostloadimportedmodule:module-type-allowed}
            steps given `moduleType`{.variable} and
            `settingsObject`{.variable} is false, then:

            1.  Let `completion`{.variable} be [Completion
                Record](https://tc39.es/ecma262/#sec-completion-record-specification-type){#hostloadimportedmodule:completion-record-4
                x-internal="completion-record"} { \[\[Type\]\]: throw,
                \[\[Value\]\]: a new
                [`TypeError`{#hostloadimportedmodule:typeerror-2}](https://tc39.es/ecma262/#sec-native-error-types-used-in-this-standard-typeerror){x-internal="typeerror"}
                exception, \[\[Target\]\]: empty }.

            2.  Perform
                [FinishLoadingImportedModule](https://tc39.es/ecma262/#sec-FinishLoadingImportedModule){#hostloadimportedmodule:finishloadingimportedmodule-4
                x-internal="finishloadingimportedmodule"}(`referrer`{.variable},
                `moduleRequest`{.variable}, `payload`{.variable},
                `completion`{.variable}).

            3.  Return.

        This step is essentially validating all of the requested module
        specifiers and type attributes when the first call to
        [HostLoadImportedModule](#hostloadimportedmodule){#hostloadimportedmodule:hostloadimportedmodule}
        for a static module dependency list is made, to avoid further
        loading operations in the case any one of the dependencies has a
        static error. We treat a module with unresolvable module
        specifiers or unsupported type attributes the same as one that
        cannot be parsed; in both cases, a syntactic issue makes it
        impossible to ever contemplate linking the module later.
        :::

8.  Let `url`{.variable} be the result of [resolving a module
    specifier](#resolve-a-module-specifier){#hostloadimportedmodule:resolve-a-module-specifier-2}
    given `referencingScript`{.variable} and
    `moduleRequest`{.variable}.\[\[Specifier\]\], catching any
    exceptions. If they throw an exception, let
    `resolutionError`{.variable} be the thrown exception.

9.  If the previous step threw an exception, then:

    1.  Let `completion`{.variable} be [Completion
        Record](https://tc39.es/ecma262/#sec-completion-record-specification-type){#hostloadimportedmodule:completion-record-5
        x-internal="completion-record"} { \[\[Type\]\]: throw,
        \[\[Value\]\]: `resolutionError`{.variable}, \[\[Target\]\]:
        empty }.

    2.  Perform
        [FinishLoadingImportedModule](https://tc39.es/ecma262/#sec-FinishLoadingImportedModule){#hostloadimportedmodule:finishloadingimportedmodule-5
        x-internal="finishloadingimportedmodule"}(`referrer`{.variable},
        `moduleRequest`{.variable}, `payload`{.variable},
        `completion`{.variable}).

    3.  Return.

10. Let `fetchOptions`{.variable} be the result of [getting the
    descendant script fetch
    options](#get-the-descendant-script-fetch-options){#hostloadimportedmodule:get-the-descendant-script-fetch-options}
    given `originalFetchOptions`{.variable}, `url`{.variable}, and
    `settingsObject`{.variable}.

11. Let `destination`{.variable} be `"script"`.

12. Let `fetchClient`{.variable} be `settingsObject`{.variable}.

13. If `loadState`{.variable} is not undefined, then:

    1.  Set `destination`{.variable} to
        `loadState`{.variable}.\[\[Destination\]\].

    2.  Set `fetchClient`{.variable} to
        `loadState`{.variable}.\[\[FetchClient\]\].

14. [Fetch a single imported module
    script](#fetch-a-single-imported-module-script){#hostloadimportedmodule:fetch-a-single-imported-module-script}
    given `url`{.variable}, `fetchClient`{.variable},
    `destination`{.variable}, `fetchOptions`{.variable},
    `settingsObject`{.variable}, `fetchReferrer`{.variable},
    `moduleRequest`{.variable}, and `onSingleFetchComplete`{.variable}
    as defined below. If `loadState`{.variable} is not undefined and
    `loadState`{.variable}.\[\[PerformFetch\]\] is not null, pass
    `loadState`{.variable}.\[\[PerformFetch\]\] along as well.

    `onSingleFetchComplete`{.variable} given `moduleScript`{.variable}
    is the following algorithm:

    1.  Let `completion`{.variable} be null.

    2.  If `moduleScript`{.variable} is null, then set
        `completion`{.variable} to [Completion
        Record](https://tc39.es/ecma262/#sec-completion-record-specification-type){#hostloadimportedmodule:completion-record-6
        x-internal="completion-record"} { \[\[Type\]\]: throw,
        \[\[Value\]\]: a new
        [`TypeError`{#hostloadimportedmodule:typeerror-3}](https://tc39.es/ecma262/#sec-native-error-types-used-in-this-standard-typeerror){x-internal="typeerror"},
        \[\[Target\]\]: empty }.

    3.  Otherwise, if `moduleScript`{.variable}\'s [parse
        error](#concept-script-parse-error){#hostloadimportedmodule:concept-script-parse-error}
        is not null, then:

        1.  Let `parseError`{.variable} be `moduleScript`{.variable}\'s
            [parse
            error](#concept-script-parse-error){#hostloadimportedmodule:concept-script-parse-error-2}.

        2.  Set `completion`{.variable} to [Completion
            Record](https://tc39.es/ecma262/#sec-completion-record-specification-type){#hostloadimportedmodule:completion-record-7
            x-internal="completion-record"} { \[\[Type\]\]: throw,
            \[\[Value\]\]: `parseError`{.variable}, \[\[Target\]\]:
            empty }.

        3.  If `loadState`{.variable} is not undefined and
            `loadState`{.variable}.\[\[ParseError\]\] is null, set
            `loadState`{.variable}.\[\[ParseError\]\] to
            `parseError`{.variable}.

    4.  Otherwise, set `completion`{.variable} to [Completion
        Record](https://tc39.es/ecma262/#sec-completion-record-specification-type){#hostloadimportedmodule:completion-record-8
        x-internal="completion-record"} { \[\[Type\]\]: normal,
        \[\[Value\]\]: `moduleScript`{.variable}\'s
        [record](#concept-script-record){#hostloadimportedmodule:concept-script-record},
        \[\[Target\]\]: empty }.

    5.  Perform
        [FinishLoadingImportedModule](https://tc39.es/ecma262/#sec-FinishLoadingImportedModule){#hostloadimportedmodule:finishloadingimportedmodule-6
        x-internal="finishloadingimportedmodule"}(`referrer`{.variable},
        `moduleRequest`{.variable}, `payload`{.variable},
        `completion`{.variable}).

#### [8.1.7]{.secno} Event loops[](#event-loops){.self-link}

##### [8.1.7.1]{.secno} Definitions[](#definitions-3){.self-link} {#definitions-3}

To coordinate events, user interaction, scripts, rendering, networking,
and so forth, user agents must use [event loops]{#event-loop .dfn
lt="event loop" export=""} as described in this section. Each
[agent](https://tc39.es/ecma262/#sec-agents){#definitions-3:agent
x-internal="agent"} has an associated [event
loop]{#concept-agent-event-loop .dfn dfn-for="agent" export=""}, which
is unique to that agent.

The [event
loop](#concept-agent-event-loop){#definitions-3:concept-agent-event-loop}
of a [similar-origin window
agent](#similar-origin-window-agent){#definitions-3:similar-origin-window-agent}
is known as a [window event loop]{#window-event-loop .dfn}. The [event
loop](#concept-agent-event-loop){#definitions-3:concept-agent-event-loop-2}
of a [dedicated worker
agent](#dedicated-worker-agent){#definitions-3:dedicated-worker-agent},
[shared worker
agent](#shared-worker-agent){#definitions-3:shared-worker-agent}, or
[service worker
agent](#service-worker-agent){#definitions-3:service-worker-agent} is
known as a [worker event loop]{#worker-event-loop-2 .dfn}. And the
[event
loop](#concept-agent-event-loop){#definitions-3:concept-agent-event-loop-3}
of a [worklet agent](#worklet-agent){#definitions-3:worklet-agent} is
known as a [worklet event loop]{#worklet-event-loop .dfn}.

::: note
[Event loops](#event-loop){#definitions-3:event-loop} do not necessarily
correspond to implementation threads. For example, multiple [window
event loops](#window-event-loop){#definitions-3:window-event-loop} could
be cooperatively scheduled in a single thread.

However, for the various worker
[agents](https://tc39.es/ecma262/#sec-agents){#definitions-3:agent-2
x-internal="agent"} that are allocated with \[\[CanBlock\]\] set to
true, the JavaScript specification does place requirements on them
regarding [forward
progress](https://tc39.es/ecma262/#sec-forward-progress){#definitions-3:forward-progress
x-internal="forward-progress"}, which effectively amount to requiring
dedicated per-agent threads in those cases.
:::

------------------------------------------------------------------------

An [event loop](#event-loop){#definitions-3:event-loop-2} has one or
more [task queues]{#task-queue .dfn}. A [task
queue](#task-queue){#definitions-3:task-queue} is a
[set](https://infra.spec.whatwg.org/#ordered-set){#definitions-3:set
x-internal="set"} of
[tasks](#concept-task){#definitions-3:concept-task}.

[Task queues](#task-queue){#definitions-3:task-queue-2} are
[sets](https://infra.spec.whatwg.org/#ordered-set){#definitions-3:set-2
x-internal="set"}, not
[queues](https://infra.spec.whatwg.org/#queue){#definitions-3:queue
x-internal="queue"}, because the [event loop processing
model](#event-loop-processing-model) grabs the first
[*runnable*](#concept-task-runnable){#definitions-3:concept-task-runnable}
[task](#concept-task){#definitions-3:concept-task-2} from the chosen
queue, instead of
[dequeuing](https://infra.spec.whatwg.org/#queue-dequeue){#definitions-3:dequeue
x-internal="dequeue"} the first task.

The [microtask queue](#microtask-queue){#definitions-3:microtask-queue}
is not a [task queue](#task-queue){#definitions-3:task-queue-3}.

Tasks encapsulate algorithms that are responsible for such work as:

Events

:   Dispatching an
    [`Event`{#definitions-3:event}](https://dom.spec.whatwg.org/#interface-event){x-internal="event"}
    object at a particular
    [`EventTarget`{#definitions-3:eventtarget}](https://dom.spec.whatwg.org/#interface-eventtarget){x-internal="eventtarget"}
    object is often done by a dedicated task.

    Not all events are dispatched using the [task
    queue](#task-queue){#definitions-3:task-queue-4}; many are
    dispatched during other tasks.

Parsing
:   The [HTML
    parser](parsing.html#html-parser){#definitions-3:html-parser}
    tokenizing one or more bytes, and then processing any resulting
    tokens, is typically a task.

Callbacks
:   Calling a callback is often done by a dedicated task.

Using a resource
:   When an algorithm
    [fetches](https://fetch.spec.whatwg.org/#concept-fetch){#definitions-3:concept-fetch
    x-internal="concept-fetch"} a resource, if the fetching occurs in a
    non-blocking fashion then the processing of the resource once some
    or all of the resource is available is performed by a task.

Reacting to DOM manipulation

:   Some elements have tasks that trigger in response to DOM
    manipulation, e.g. when that element is [inserted into the
    document](infrastructure.html#insert-an-element-into-a-document){#definitions-3:insert-an-element-into-a-document}.

Formally, a [task]{#concept-task .dfn lt="task" export=""} is a
[struct](https://infra.spec.whatwg.org/#struct){#definitions-3:struct
x-internal="struct"} which has:

[Steps]{#concept-task-steps .dfn}
:   A series of steps specifying the work to be done by the task.

A [source]{#concept-task-source .dfn}
:   One of the [task sources](#task-source){#definitions-3:task-source},
    used to group and serialize related tasks.

A [document]{#concept-task-document .dfn}
:   A [`Document`{#definitions-3:document}](dom.html#document)
    associated with the task, or null for tasks that are not in a
    [window event
    loop](#window-event-loop){#definitions-3:window-event-loop-2}.

A [script evaluation environment settings object set]{#script-evaluation-environment-settings-object-set .dfn export=""}
:   A
    [set](https://infra.spec.whatwg.org/#ordered-set){#definitions-3:set-3
    x-internal="set"} of [environment settings
    objects](#environment-settings-object){#definitions-3:environment-settings-object}
    used for tracking script evaluation during the task.

A [task](#concept-task){#definitions-3:concept-task-3} is
[runnable]{#concept-task-runnable .dfn} if its
[document](#concept-task-document){#definitions-3:concept-task-document}
is either null or [fully
active](document-sequences.html#fully-active){#definitions-3:fully-active}.

Per its
[source](#concept-task-source){#definitions-3:concept-task-source}
field, each [task](#concept-task){#definitions-3:concept-task-4} is
defined as coming from a specific [task source]{#task-source .dfn
export=""}. For each [event
loop](#event-loop){#definitions-3:event-loop-3}, every [task
source](#task-source){#definitions-3:task-source-2} must be associated
with a specific [task queue](#task-queue){#definitions-3:task-queue-5}.

Essentially, [task sources](#task-source){#definitions-3:task-source-3}
are used within standards to separate logically-different types of
tasks, which a user agent might wish to distinguish between. [Task
queues](#task-queue){#definitions-3:task-queue-6} are used by user
agents to coalesce task sources within a given [event
loop](#event-loop){#definitions-3:event-loop-4}.

For example, a user agent could have one [task
queue](#task-queue){#definitions-3:task-queue-7} for mouse and key
events (to which the [user interaction task
source](#user-interaction-task-source){#definitions-3:user-interaction-task-source}
is associated), and another to which all other [task
sources](#task-source){#definitions-3:task-source-4} are associated.
Then, using the freedom granted in the initial step of the [event loop
processing model](#event-loop-processing-model), it could give keyboard
and mouse events preference over other tasks three-quarters of the time,
keeping the interface responsive but not starving other task queues.
Note that in this setup, the processing model still enforces that the
user agent would never process events from any one [task
source](#task-source){#definitions-3:task-source-5} out of order.

------------------------------------------------------------------------

Each [event loop](#event-loop){#definitions-3:event-loop-5} has a
[currently running task]{#currently-running-task .dfn}, which is either
a [task](#concept-task){#definitions-3:concept-task-5} or null.
Initially, this is null. It is used to handle reentrancy.

Each [event loop](#event-loop){#definitions-3:event-loop-6} has a
[microtask queue]{#microtask-queue .dfn}, which is a
[queue](https://infra.spec.whatwg.org/#queue){#definitions-3:queue-2
x-internal="queue"} of
[microtasks](#microtask){#definitions-3:microtask}, initially empty. A
[microtask]{#microtask .dfn export=""} is a colloquial way of referring
to a [task](#concept-task){#definitions-3:concept-task-6} that was
created via the [queue a
microtask](#queue-a-microtask){#definitions-3:queue-a-microtask}
algorithm.

Each [event loop](#event-loop){#definitions-3:event-loop-7} has a
[performing a microtask checkpoint]{#performing-a-microtask-checkpoint
.dfn} boolean, which is initially false. It is used to prevent reentrant
invocation of the [perform a microtask
checkpoint](#perform-a-microtask-checkpoint){#definitions-3:perform-a-microtask-checkpoint}
algorithm.

Each [window event
loop](#window-event-loop){#definitions-3:window-event-loop-3} has a
[`DOMHighResTimeStamp`{#definitions-3:domhighrestimestamp}](https://w3c.github.io/hr-time/#dom-domhighrestimestamp){x-internal="domhighrestimestamp"}
[last render opportunity time]{#last-render-opportunity-time .dfn},
initially set to zero.

Each [window event
loop](#window-event-loop){#definitions-3:window-event-loop-4} has a
[`DOMHighResTimeStamp`{#definitions-3:domhighrestimestamp-2}](https://w3c.github.io/hr-time/#dom-domhighrestimestamp){x-internal="domhighrestimestamp"}
[last idle period start time]{#last-idle-period-start-time .dfn},
initially set to zero.

To get the [same-loop windows]{#same-loop-windows .dfn} for a [window
event loop](#window-event-loop){#definitions-3:window-event-loop-5}
`loop`{.variable}, return all
[`Window`{#definitions-3:window}](nav-history-apis.html#window) objects
whose [relevant
agent](#relevant-agent){#definitions-3:relevant-agent}\'s [event
loop](#concept-agent-event-loop){#definitions-3:concept-agent-event-loop-4}
is `loop`{.variable}.

##### [8.1.7.2]{.secno} Queuing tasks[](#queuing-tasks){.self-link}

To [queue a task]{#queue-a-task .dfn export=""} on a [task
source](#task-source){#queuing-tasks:task-source} `source`{.variable},
which performs a series of steps `steps`{.variable}, optionally given an
event loop `event loop`{.variable} and a document `document`{.variable}:

1.  If `event loop`{.variable} was not given, set
    `event loop`{.variable} to the [implied event
    loop](#implied-event-loop){#queuing-tasks:implied-event-loop}.

2.  If `document`{.variable} was not given, set `document`{.variable} to
    the [implied
    document](#implied-document){#queuing-tasks:implied-document}.

3.  Let `task`{.variable} be a new
    [task](#concept-task){#queuing-tasks:concept-task}.

4.  Set `task`{.variable}\'s
    [steps](#concept-task-steps){#queuing-tasks:concept-task-steps} to
    `steps`{.variable}.

5.  Set `task`{.variable}\'s
    [source](#concept-task-source){#queuing-tasks:concept-task-source}
    to `source`{.variable}.

6.  Set `task`{.variable}\'s
    [document](#concept-task-document){#queuing-tasks:concept-task-document}
    to the `document`{.variable}.

7.  Set `task`{.variable}\'s [script evaluation environment settings
    object
    set](#script-evaluation-environment-settings-object-set){#queuing-tasks:script-evaluation-environment-settings-object-set}
    to an empty
    [set](https://infra.spec.whatwg.org/#ordered-set){#queuing-tasks:set
    x-internal="set"}.

8.  Let `queue`{.variable} be the [task
    queue](#task-queue){#queuing-tasks:task-queue} to which
    `source`{.variable} is associated on `event loop`{.variable}.

9.  [Append](https://infra.spec.whatwg.org/#list-append){#queuing-tasks:list-append
    x-internal="list-append"} `task`{.variable} to `queue`{.variable}.

Failing to pass an event loop and document to [queue a
task](#queue-a-task){#queuing-tasks:queue-a-task} means relying on the
ambiguous and poorly-specified [implied event
loop](#implied-event-loop){#queuing-tasks:implied-event-loop-2} and
[implied document](#implied-document){#queuing-tasks:implied-document-2}
concepts. Specification authors should either always pass these values,
or use the wrapper algorithms [queue a global
task](#queue-a-global-task){#queuing-tasks:queue-a-global-task} or
[queue an element
task](#queue-an-element-task){#queuing-tasks:queue-an-element-task}
instead. Using the wrapper algorithms is recommended.

To [queue a global task]{#queue-a-global-task .dfn export=""} on a [task
source](#task-source){#queuing-tasks:task-source-2} `source`{.variable},
with a [global object](#global-object){#queuing-tasks:global-object}
`global`{.variable} and a series of steps `steps`{.variable}:

1.  Let `event loop`{.variable} be `global`{.variable}\'s [relevant
    agent](#relevant-agent){#queuing-tasks:relevant-agent}\'s [event
    loop](#concept-agent-event-loop){#queuing-tasks:concept-agent-event-loop}.

2.  Let `document`{.variable} be `global`{.variable}\'s [associated
    `Document`](nav-history-apis.html#concept-document-window){#queuing-tasks:concept-document-window},
    if `global`{.variable} is a
    [`Window`{#queuing-tasks:window}](nav-history-apis.html#window)
    object; otherwise null.

3.  [Queue a task](#queue-a-task){#queuing-tasks:queue-a-task-2} given
    `source`{.variable}, `event loop`{.variable}, `document`{.variable},
    and `steps`{.variable}.

To [queue an element task]{#queue-an-element-task .dfn export=""} on a
[task source](#task-source){#queuing-tasks:task-source-3}
`source`{.variable}, with an element `element`{.variable} and a series
of steps `steps`{.variable}:

1.  Let `global`{.variable} be `element`{.variable}\'s [relevant global
    object](#concept-relevant-global){#queuing-tasks:concept-relevant-global}.

2.  [Queue a global
    task](#queue-a-global-task){#queuing-tasks:queue-a-global-task-2}
    given `source`{.variable}, `global`{.variable}, and
    `steps`{.variable}.

To [queue a microtask]{#queue-a-microtask .dfn export=""} which performs
a series of steps `steps`{.variable}, optionally given a document
`document`{.variable}:

1.  [Assert](https://infra.spec.whatwg.org/#assert){#queuing-tasks:assert
    x-internal="assert"}: there is a [surrounding
    agent](https://tc39.es/ecma262/#surrounding-agent){#queuing-tasks:surrounding-agent
    x-internal="surrounding-agent"}. I.e., this algorithm is not called
    while [in
    parallel](infrastructure.html#in-parallel){#queuing-tasks:in-parallel}.

2.  Let `eventLoop`{.variable} be the [surrounding
    agent](https://tc39.es/ecma262/#surrounding-agent){#queuing-tasks:surrounding-agent-2
    x-internal="surrounding-agent"}\'s [event
    loop](#concept-agent-event-loop){#queuing-tasks:concept-agent-event-loop-2}.

3.  If `document`{.variable} was not given, set `document`{.variable} to
    the [implied
    document](#implied-document){#queuing-tasks:implied-document-3}.

4.  Let `microtask`{.variable} be a new
    [task](#concept-task){#queuing-tasks:concept-task-2}.

5.  Set `microtask`{.variable}\'s
    [steps](#concept-task-steps){#queuing-tasks:concept-task-steps-2} to
    `steps`{.variable}.

6.  Set `microtask`{.variable}\'s
    [source](#concept-task-source){#queuing-tasks:concept-task-source-2}
    to the [microtask task source]{#microtask-task-source .dfn}.

7.  Set `microtask`{.variable}\'s
    [document](#concept-task-document){#queuing-tasks:concept-task-document-2}
    to `document`{.variable}.

8.  Set `microtask`{.variable}\'s [script evaluation environment
    settings object
    set](#script-evaluation-environment-settings-object-set){#queuing-tasks:script-evaluation-environment-settings-object-set-2}
    to an empty
    [set](https://infra.spec.whatwg.org/#ordered-set){#queuing-tasks:set-2
    x-internal="set"}.

9.  [Enqueue](https://infra.spec.whatwg.org/#queue-enqueue){#queuing-tasks:enqueue
    x-internal="enqueue"} `microtask`{.variable} on
    `eventLoop`{.variable}\'s [microtask
    queue](#microtask-queue){#queuing-tasks:microtask-queue}.

It is possible for a [microtask](#microtask){#queuing-tasks:microtask}
to be moved to a regular [task
queue](#task-queue){#queuing-tasks:task-queue-2}, if, during its initial
execution, it [spins the event
loop](#spin-the-event-loop){#queuing-tasks:spin-the-event-loop}. This is
the only case in which the
[source](#concept-task-source){#queuing-tasks:concept-task-source-3},
[document](#concept-task-document){#queuing-tasks:concept-task-document-3},
and [script evaluation environment settings object
set](#script-evaluation-environment-settings-object-set){#queuing-tasks:script-evaluation-environment-settings-object-set-3}
of the microtask are consulted; they are ignored by the [perform a
microtask
checkpoint](#perform-a-microtask-checkpoint){#queuing-tasks:perform-a-microtask-checkpoint}
algorithm.

The [implied event loop]{#implied-event-loop .dfn} when queuing a task
is the one that can deduced from the context of the calling algorithm.
This is generally unambiguous, as most specification algorithms only
ever involve a single
[agent](https://tc39.es/ecma262/#sec-agents){#queuing-tasks:agent
x-internal="agent"} (and thus a single [event
loop](#event-loop){#queuing-tasks:event-loop}). The exception is
algorithms involving or specifying cross-agent communication (e.g.,
between a window and a worker); for those cases, the [implied event
loop](#implied-event-loop){#queuing-tasks:implied-event-loop-3} concept
must not be relied upon and specifications must explicitly provide an
[event loop](#event-loop){#queuing-tasks:event-loop-2} when [queuing a
task](#queue-a-task){#queuing-tasks:queue-a-task-3}.

The [implied document]{#implied-document .dfn} when queuing a task on an
[event loop](#event-loop){#queuing-tasks:event-loop-3}
`event loop`{.variable} is determined as follows:

1.  If `event loop`{.variable} is not a [window event
    loop](#window-event-loop){#queuing-tasks:window-event-loop}, then
    return null.

2.  If the task is being queued in the context of an element, then
    return the element\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#queuing-tasks:node-document
    x-internal="node-document"}.

3.  If the task is being queued in the context of a [browsing
    context](document-sequences.html#browsing-context){#queuing-tasks:browsing-context},
    then return the browsing context\'s [active
    document](document-sequences.html#active-document){#queuing-tasks:active-document}.

4.  If the task is being queued by or for a
    [script](#concept-script){#queuing-tasks:concept-script}, then
    return the script\'s [settings
    object](#settings-object){#queuing-tasks:settings-object}\'s [global
    object](#concept-settings-object-global){#queuing-tasks:concept-settings-object-global}\'s
    [associated
    `Document`](nav-history-apis.html#concept-document-window){#queuing-tasks:concept-document-window-2}.

5.  [Assert](https://infra.spec.whatwg.org/#assert){#queuing-tasks:assert-2
    x-internal="assert"}: this step is never reached, because one of the
    previous conditions is true. [Really?]{.XXX}

Both [implied event
loop](#implied-event-loop){#queuing-tasks:implied-event-loop-4} and
[implied document](#implied-document){#queuing-tasks:implied-document-4}
are vaguely-defined and have a lot of action-at-a-distance. The hope is
to remove these, especially [implied
document](#implied-document){#queuing-tasks:implied-document-5}. See
[issue #4980](https://github.com/whatwg/html/issues/4980).

##### [8.1.7.3]{.secno} []{#processing-model-8}Processing model[](#event-loop-processing-model){.self-link} {#event-loop-processing-model dfn-type="dfn" lt="event loop processing model"}

An [event loop](#event-loop){#event-loop-processing-model:event-loop}
must continually run through the following steps for as long as it
exists:

1.  Let `oldestTask`{.variable} and `taskStartTime`{.variable} be null.

2.  If the [event
    loop](#event-loop){#event-loop-processing-model:event-loop-2} has a
    [task queue](#task-queue){#event-loop-processing-model:task-queue}
    with at least one
    [runnable](#concept-task-runnable){#event-loop-processing-model:concept-task-runnable}
    [task](#concept-task){#event-loop-processing-model:concept-task},
    then:

    1.  Let `taskQueue`{.variable} be one such [task
        queue](#task-queue){#event-loop-processing-model:task-queue-2},
        chosen in an
        [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#event-loop-processing-model:implementation-defined
        x-internal="implementation-defined"} manner.

        Remember that the [microtask
        queue](#microtask-queue){#event-loop-processing-model:microtask-queue}
        is not a [task
        queue](#task-queue){#event-loop-processing-model:task-queue-3},
        so it will not be chosen in this step. However, a [task
        queue](#task-queue){#event-loop-processing-model:task-queue-4}
        to which the [microtask task
        source](#microtask-task-source){#event-loop-processing-model:microtask-task-source}
        is associated might be chosen in this step. In that case, the
        [task](#concept-task){#event-loop-processing-model:concept-task-2}
        chosen in the next step was originally a
        [microtask](#microtask){#event-loop-processing-model:microtask},
        but it got moved as part of [spinning the event
        loop](#spin-the-event-loop){#event-loop-processing-model:spin-the-event-loop}.

    2.  Set `taskStartTime`{.variable} to the [unsafe shared current
        time](https://w3c.github.io/hr-time/#dfn-unsafe-shared-current-time){#event-loop-processing-model:unsafe-shared-current-time
        x-internal="unsafe-shared-current-time"}.

    3.  Set `oldestTask`{.variable} to the first
        [runnable](#concept-task-runnable){#event-loop-processing-model:concept-task-runnable-2}
        [task](#concept-task){#event-loop-processing-model:concept-task-3}
        in `taskQueue`{.variable}, and
        [remove](https://infra.spec.whatwg.org/#list-remove){#event-loop-processing-model:list-remove
        x-internal="list-remove"} it from `taskQueue`{.variable}.

    4.  If `oldestTask`{.variable}\'s
        [document](#concept-task-document){#event-loop-processing-model:concept-task-document}
        is not null, then [record task start
        time](https://w3c.github.io/long-animation-frames/#record-task-start-time){#event-loop-processing-model:record-task-start-time
        x-internal="record-task-start-time"} given
        `taskStartTime`{.variable} and `oldestTask`{.variable}\'s
        [document](#concept-task-document){#event-loop-processing-model:concept-task-document-2}.

    5.  Set the [event
        loop](#event-loop){#event-loop-processing-model:event-loop-3}\'s
        [currently running
        task](#currently-running-task){#event-loop-processing-model:currently-running-task}
        to `oldestTask`{.variable}.

    6.  Perform `oldestTask`{.variable}\'s
        [steps](#concept-task-steps){#event-loop-processing-model:concept-task-steps}.

    7.  Set the [event
        loop](#event-loop){#event-loop-processing-model:event-loop-4}\'s
        [currently running
        task](#currently-running-task){#event-loop-processing-model:currently-running-task-2}
        back to null.

    8.  [Perform a microtask
        checkpoint](#perform-a-microtask-checkpoint){#event-loop-processing-model:perform-a-microtask-checkpoint}.

3.  Let `taskEndTime`{.variable} be the [unsafe shared current
    time](https://w3c.github.io/hr-time/#dfn-unsafe-shared-current-time){#event-loop-processing-model:unsafe-shared-current-time-2
    x-internal="unsafe-shared-current-time"}.
    [\[HRT\]](references.html#refsHRT)

4.  If `oldestTask`{.variable} is not null, then:

    1.  Let `top-level browsing contexts`{.variable} be an empty
        [set](https://infra.spec.whatwg.org/#ordered-set){#event-loop-processing-model:set
        x-internal="set"}.

    2.  For each [environment settings
        object](#environment-settings-object){#event-loop-processing-model:environment-settings-object}
        `settings`{.variable} of `oldestTask`{.variable}\'s [script
        evaluation environment settings object
        set](#script-evaluation-environment-settings-object-set){#event-loop-processing-model:script-evaluation-environment-settings-object-set}:

        1.  Let `global`{.variable} be `settings`{.variable}\'s [global
            object](#concept-settings-object-global){#event-loop-processing-model:concept-settings-object-global}.

        2.  If `global`{.variable} is not a
            [`Window`{#event-loop-processing-model:window}](nav-history-apis.html#window)
            object, then
            [continue](https://infra.spec.whatwg.org/#iteration-continue){#event-loop-processing-model:continue
            x-internal="continue"}.

        3.  If `global`{.variable}\'s [browsing
            context](nav-history-apis.html#window-bc){#event-loop-processing-model:window-bc}
            is null, then
            [continue](https://infra.spec.whatwg.org/#iteration-continue){#event-loop-processing-model:continue-2
            x-internal="continue"}.

        4.  Let `tlbc`{.variable} be `global`{.variable}\'s [browsing
            context](nav-history-apis.html#window-bc){#event-loop-processing-model:window-bc-2}\'s
            [top-level browsing
            context](document-sequences.html#bc-tlbc){#event-loop-processing-model:bc-tlbc}.

        5.  If `tlbc`{.variable} is not null, then
            [append](https://infra.spec.whatwg.org/#set-append){#event-loop-processing-model:set-append
            x-internal="set-append"} it to
            `top-level browsing contexts`{.variable}.

    3.  [Report long
        tasks](https://w3c.github.io/longtasks/#report-long-tasks){#event-loop-processing-model:report-long-tasks
        x-internal="report-long-tasks"}, passing in
        `taskStartTime`{.variable}, `taskEndTime`{.variable},
        `top-level browsing contexts`{.variable}, and
        `oldestTask`{.variable}.

    4.  If `oldestTask`{.variable}\'s
        [document](#concept-task-document){#event-loop-processing-model:concept-task-document-3}
        is not null, then [record task end
        time](https://w3c.github.io/long-animation-frames/#record-task-end-time){#event-loop-processing-model:record-task-end-time
        x-internal="record-task-end-time"} given
        `taskEndTime`{.variable} and `oldestTask`{.variable}\'s
        [document](#concept-task-document){#event-loop-processing-model:concept-task-document-4}.

5.  ::: {#idle-deadline-computation}
    If this is a [window event
    loop](#window-event-loop){#event-loop-processing-model:window-event-loop}
    that has no
    [*runnable*](#concept-task-runnable){#event-loop-processing-model:concept-task-runnable-3}
    [task](#concept-task){#event-loop-processing-model:concept-task-4}
    in this [event
    loop](#event-loop){#event-loop-processing-model:event-loop-5}\'s
    [task
    queues](#task-queue){#event-loop-processing-model:task-queue-5},
    then:

    1.  Set this [event
        loop](#event-loop){#event-loop-processing-model:event-loop-6}\'s
        [last idle period start
        time](#last-idle-period-start-time){#event-loop-processing-model:last-idle-period-start-time}
        to the [unsafe shared current
        time](https://w3c.github.io/hr-time/#dfn-unsafe-shared-current-time){#event-loop-processing-model:unsafe-shared-current-time-3
        x-internal="unsafe-shared-current-time"}.

    2.  Let `computeDeadline`{.variable} be the following steps:

        1.  Let `deadline`{.variable} be this [event
            loop](#event-loop){#event-loop-processing-model:event-loop-7}\'s
            [last idle period start
            time](#last-idle-period-start-time){#event-loop-processing-model:last-idle-period-start-time-2}
            plus 50.

            The cap of 50ms in the future is to ensure responsiveness to
            new user input within the threshold of human perception.

        2.  Let `hasPendingRenders`{.variable} be false.

        3.  For each `windowInSameLoop`{.variable} of the [same-loop
            windows](#same-loop-windows){#event-loop-processing-model:same-loop-windows}
            for this [event
            loop](#event-loop){#event-loop-processing-model:event-loop-8}:

            1.  If `windowInSameLoop`{.variable}\'s [map of animation
                frame
                callbacks](imagebitmap-and-animations.html#list-of-animation-frame-callbacks){#event-loop-processing-model:list-of-animation-frame-callbacks}
                is not
                [empty](https://infra.spec.whatwg.org/#map-is-empty){#event-loop-processing-model:map-empty
                x-internal="map-empty"}, or if the user agent believes
                that the `windowInSameLoop`{.variable} might have
                pending rendering updates, set
                `hasPendingRenders`{.variable} to true.

            2.  Let `timerCallbackEstimates`{.variable} be the result of
                [getting the
                values](https://infra.spec.whatwg.org/#map-getting-the-values){#event-loop-processing-model:map-get-the-values
                x-internal="map-get-the-values"} of
                `windowInSameLoop`{.variable}\'s [map of active
                timers](timers-and-user-prompts.html#map-of-active-timers){#event-loop-processing-model:map-of-active-timers}.

            3.  For each `timeoutDeadline`{.variable} of
                `timerCallbackEstimates`{.variable}, if
                `timeoutDeadline`{.variable} is less than
                `deadline`{.variable}, set `deadline`{.variable} to
                `timeoutDeadline`{.variable}.

        4.  If `hasPendingRenders`{.variable} is true, then:

            1.  Let `nextRenderDeadline`{.variable} be this [event
                loop](#event-loop){#event-loop-processing-model:event-loop-9}\'s
                [last render opportunity
                time](#last-render-opportunity-time){#event-loop-processing-model:last-render-opportunity-time}
                plus (1000 divided by the current refresh rate).

                The refresh rate can be hardware- or
                implementation-specific. For a refresh rate of 60Hz, the
                `nextRenderDeadline`{.variable} would be about 16.67ms
                after the [last render opportunity
                time](#last-render-opportunity-time){#event-loop-processing-model:last-render-opportunity-time-2}.

            2.  If `nextRenderDeadline`{.variable} is less than
                `deadline`{.variable}, then return
                `nextRenderDeadline`{.variable}.

        5.  Return `deadline`{.variable}.

    3.  For each `win`{.variable} of the [same-loop
        windows](#same-loop-windows){#event-loop-processing-model:same-loop-windows-2}
        for this [event
        loop](#event-loop){#event-loop-processing-model:event-loop-10},
        perform the [start an idle period
        algorithm](https://w3c.github.io/requestidlecallback/#start-an-idle-period-algorithm){#event-loop-processing-model:start-an-idle-period-algorithm
        x-internal="start-an-idle-period-algorithm"} for
        `win`{.variable} with the following step: return the result of
        calling `computeDeadline`{.variable},
        [coarsened](https://w3c.github.io/hr-time/#dfn-coarsen-time){#event-loop-processing-model:coarsen-time
        x-internal="coarsen-time"} given `win`{.variable}\'s [relevant
        settings
        object](#relevant-settings-object){#event-loop-processing-model:relevant-settings-object}\'s
        [cross-origin isolated
        capability](#concept-settings-object-cross-origin-isolated-capability){#event-loop-processing-model:concept-settings-object-cross-origin-isolated-capability}.
        [\[REQUESTIDLECALLBACK\]](references.html#refsREQUESTIDLECALLBACK)
    :::

6.  If this is a [worker event
    loop](#worker-event-loop-2){#event-loop-processing-model:worker-event-loop-2},
    then:

    1.  If this [event
        loop](#event-loop){#event-loop-processing-model:event-loop-11}\'s
        [agent](https://tc39.es/ecma262/#sec-agents){#event-loop-processing-model:agent
        x-internal="agent"}\'s single
        [realm](https://tc39.es/ecma262/#sec-code-realms){#event-loop-processing-model:realm
        x-internal="realm"}\'s [global
        object](#concept-realm-global){#event-loop-processing-model:concept-realm-global}
        is a
        [supported](imagebitmap-and-animations.html#concept-animationframeprovider-supported){#event-loop-processing-model:concept-animationframeprovider-supported}
        [`DedicatedWorkerGlobalScope`{#event-loop-processing-model:dedicatedworkerglobalscope}](workers.html#dedicatedworkerglobalscope)
        and the user agent believes that it would benefit from having
        its rendering updated at this time, then:

        1.  Let `now`{.variable} be the [current high resolution
            time](https://w3c.github.io/hr-time/#dfn-current-high-resolution-time){#event-loop-processing-model:current-high-resolution-time
            x-internal="current-high-resolution-time"} given the
            [`DedicatedWorkerGlobalScope`{#event-loop-processing-model:dedicatedworkerglobalscope-2}](workers.html#dedicatedworkerglobalscope).
            [\[HRT\]](references.html#refsHRT)

        2.  [Run the animation frame
            callbacks](imagebitmap-and-animations.html#run-the-animation-frame-callbacks){#event-loop-processing-model:run-the-animation-frame-callbacks}
            for that
            [`DedicatedWorkerGlobalScope`{#event-loop-processing-model:dedicatedworkerglobalscope-3}](workers.html#dedicatedworkerglobalscope),
            passing in `now`{.variable} as the timestamp.

        3.  Update the rendering of that dedicated worker to reflect the
            current state.

        Similar to the notes for [updating the
        rendering](#update-the-rendering){#event-loop-processing-model:update-the-rendering}
        in a [window event
        loop](#window-event-loop){#event-loop-processing-model:window-event-loop-2},
        a user agent can determine the rate of rendering in the
        dedicated worker.

    2.  If there are no
        [tasks](#concept-task){#event-loop-processing-model:concept-task-5}
        in the [event
        loop](#event-loop){#event-loop-processing-model:event-loop-12}\'s
        [task
        queues](#task-queue){#event-loop-processing-model:task-queue-6}
        and the
        [`WorkerGlobalScope`{#event-loop-processing-model:workerglobalscope}](workers.html#workerglobalscope)
        object\'s
        [closing](workers.html#dom-workerglobalscope-closing){#event-loop-processing-model:dom-workerglobalscope-closing}
        flag is true, then destroy the [event
        loop](#event-loop){#event-loop-processing-model:event-loop-13},
        aborting these steps, resuming the [run a
        worker](workers.html#run-a-worker){#event-loop-processing-model:run-a-worker}
        steps described in the [Web workers](workers.html#workers)
        section below.

A [window event
loop](#window-event-loop){#event-loop-processing-model:window-event-loop-3}
`eventLoop`{.variable} must also run the following [in
parallel](infrastructure.html#in-parallel){#event-loop-processing-model:in-parallel},
as long as it exists:

1.  Wait until at least one
    [navigable](document-sequences.html#navigable){#event-loop-processing-model:navigable}
    whose [active
    document](document-sequences.html#nav-document){#event-loop-processing-model:nav-document}\'s
    [relevant
    agent](#relevant-agent){#event-loop-processing-model:relevant-agent}\'s
    [event
    loop](#concept-agent-event-loop){#event-loop-processing-model:concept-agent-event-loop}
    is `eventLoop`{.variable} might have a [rendering
    opportunity](#rendering-opportunity){#event-loop-processing-model:rendering-opportunity}.

2.  Set `eventLoop`{.variable}\'s [last render opportunity
    time](#last-render-opportunity-time){#event-loop-processing-model:last-render-opportunity-time-3}
    to the [unsafe shared current
    time](https://w3c.github.io/hr-time/#dfn-unsafe-shared-current-time){#event-loop-processing-model:unsafe-shared-current-time-4
    x-internal="unsafe-shared-current-time"}.

3.  For each `navigable`{.variable} that has a [rendering
    opportunity](#rendering-opportunity){#event-loop-processing-model:rendering-opportunity-2},
    [queue a global
    task](#queue-a-global-task){#event-loop-processing-model:queue-a-global-task}
    on the [rendering task
    source](#rendering-task-source){#event-loop-processing-model:rendering-task-source}
    given `navigable`{.variable}\'s [active
    window](document-sequences.html#nav-window){#event-loop-processing-model:nav-window}
    to [update the rendering]{#update-the-rendering .dfn export=""}:

    This might cause redundant calls to [update the
    rendering](#update-the-rendering){#event-loop-processing-model:update-the-rendering-2}.
    However, these calls would have no observable effect because there
    will be no rendering necessary, as per the *Unnecessary rendering*
    step. Implementations can introduce further optimizations such as
    only queuing this task when it is not already queued. However, note
    that the document associated with the task might become inactive
    before the task is processed.

    1.  Let `frameTimestamp`{.variable} be `eventLoop`{.variable}\'s
        [last render opportunity
        time](#last-render-opportunity-time){#event-loop-processing-model:last-render-opportunity-time-4}.

    2.  Let `docs`{.variable} be all [fully
        active](document-sequences.html#fully-active){#event-loop-processing-model:fully-active}
        [`Document`{#event-loop-processing-model:document}](dom.html#document)
        objects whose [relevant
        agent](#relevant-agent){#event-loop-processing-model:relevant-agent-2}\'s
        [event
        loop](#concept-agent-event-loop){#event-loop-processing-model:concept-agent-event-loop-2}
        is `eventLoop`{.variable}, sorted arbitrarily except that the
        following conditions must be met:

        - Any
          [`Document`{#event-loop-processing-model:document-2}](dom.html#document)
          `B`{.variable} whose [container
          document](document-sequences.html#doc-container-document){#event-loop-processing-model:doc-container-document}
          is `A`{.variable} must be listed after `A`{.variable} in the
          list.

        - If there are two documents `A`{.variable} and `B`{.variable}
          that both have the same non-null [container
          document](document-sequences.html#doc-container-document){#event-loop-processing-model:doc-container-document-2}
          `C`{.variable}, then the order of `A`{.variable} and
          `B`{.variable} in the list must match the [shadow-including
          tree
          order](https://dom.spec.whatwg.org/#concept-shadow-including-tree-order){#event-loop-processing-model:shadow-including-tree-order
          x-internal="shadow-including-tree-order"} of their respective
          [navigable
          containers](document-sequences.html#navigable-container){#event-loop-processing-model:navigable-container}
          in `C`{.variable}\'s [node
          tree](https://dom.spec.whatwg.org/#concept-node-tree){#event-loop-processing-model:node-tree
          x-internal="node-tree"}.

        In the steps below that iterate over `docs`{.variable}, each
        [`Document`{#event-loop-processing-model:document-3}](dom.html#document)
        must be processed in the order it is found in the list.

    3.  *Filter non-renderable documents*: Remove from `docs`{.variable}
        any
        [`Document`{#event-loop-processing-model:document-4}](dom.html#document)
        object `doc`{.variable} for which any of the following are true:

        - `doc`{.variable} is
          [render-blocked](dom.html#render-blocked){#event-loop-processing-model:render-blocked};

        - `doc`{.variable}\'s [visibility
          state](interaction.html#visibility-state){#event-loop-processing-model:visibility-state}
          is \"`hidden`\";

        - `doc`{.variable}\'s rendering is [suppressed for view
          transitions](https://drafts.csswg.org/css-view-transitions/#document-rendering-suppression-for-view-transitions){#event-loop-processing-model:rendering-suppression-for-view-transitions
          x-internal="rendering-suppression-for-view-transitions"}; or

        - `doc`{.variable}\'s [node
          navigable](document-sequences.html#node-navigable){#event-loop-processing-model:node-navigable}
          doesn\'t currently have a [rendering
          opportunity](#rendering-opportunity){#event-loop-processing-model:rendering-opportunity-3}.

        We have to check for rendering opportunities here, in addition
        to checking that in the [in
        parallel](infrastructure.html#in-parallel){#event-loop-processing-model:in-parallel-2}
        steps, as some documents that share the same [event
        loop](#event-loop){#event-loop-processing-model:event-loop-14}
        might not have a [rendering
        opportunity](#rendering-opportunity){#event-loop-processing-model:rendering-opportunity-4}
        at the same time.

    4.  *Unnecessary rendering*: Remove from `docs`{.variable} any
        [`Document`{#event-loop-processing-model:document-5}](dom.html#document)
        object `doc`{.variable} for which all of the following are true:

        - the user agent believes that updating the rendering of
          `doc`{.variable}\'s [node
          navigable](document-sequences.html#node-navigable){#event-loop-processing-model:node-navigable-2}
          would have no visible effect; and

        - `doc`{.variable}\'s [map of animation frame
          callbacks](imagebitmap-and-animations.html#list-of-animation-frame-callbacks){#event-loop-processing-model:list-of-animation-frame-callbacks-2}
          is empty.

    5.  Remove from `docs`{.variable} all
        [`Document`{#event-loop-processing-model:document-6}](dom.html#document)
        objects for which the user agent believes that it\'s preferable
        to skip updating the rendering for other reasons.

        ::: note
        The step labeled *Filter non-renderable documents* prevents the
        user agent from updating the rendering when it is unable to
        present new content to the user.

        The step labeled *Unnecessary rendering* prevents the user agent
        from updating the rendering when there\'s no new content to
        draw.

        This step enables the user agent to prevent the steps below from
        running for other reasons, for example, to ensure certain
        [tasks](#concept-task){#event-loop-processing-model:concept-task-6}
        are executed immediately after each other, with only [microtask
        checkpoints](#perform-a-microtask-checkpoint){#event-loop-processing-model:perform-a-microtask-checkpoint-2}
        interleaved (and without, e.g., [animation frame
        callbacks](imagebitmap-and-animations.html#run-the-animation-frame-callbacks){#event-loop-processing-model:run-the-animation-frame-callbacks-2}
        interleaved). Concretely, a user agent might wish to coalesce
        timer callbacks together, with no intermediate rendering
        updates.
        :::

    6.  For each `doc`{.variable} of `docs`{.variable},
        [reveal](browsing-the-web.html#reveal){#event-loop-processing-model:reveal}
        `doc`{.variable}.

    7.  For each `doc`{.variable} of `docs`{.variable}, [flush autofocus
        candidates](interaction.html#flush-autofocus-candidates){#event-loop-processing-model:flush-autofocus-candidates}
        for `doc`{.variable} if its [node
        navigable](document-sequences.html#node-navigable){#event-loop-processing-model:node-navigable-3}
        is a [top-level
        traversable](document-sequences.html#top-level-traversable){#event-loop-processing-model:top-level-traversable}.

    8.  For each `doc`{.variable} of `docs`{.variable}, [run the resize
        steps](https://drafts.csswg.org/cssom-view/#document-run-the-resize-steps){#event-loop-processing-model:run-the-resize-steps
        x-internal="run-the-resize-steps"} for `doc`{.variable}.
        [\[CSSOMVIEW\]](references.html#refsCSSOMVIEW)

    9.  For each `doc`{.variable} of `docs`{.variable}, [run the scroll
        steps](https://drafts.csswg.org/cssom-view/#document-run-the-scroll-steps){#event-loop-processing-model:run-the-scroll-steps
        x-internal="run-the-scroll-steps"} for `doc`{.variable}.
        [\[CSSOMVIEW\]](references.html#refsCSSOMVIEW)

    10. For each `doc`{.variable} of `docs`{.variable}, [evaluate media
        queries and report
        changes](https://drafts.csswg.org/cssom-view/#evaluate-media-queries-and-report-changes){#event-loop-processing-model:evaluate-media-queries-and-report-changes
        x-internal="evaluate-media-queries-and-report-changes"} for
        `doc`{.variable}. [\[CSSOMVIEW\]](references.html#refsCSSOMVIEW)

    11. For each `doc`{.variable} of `docs`{.variable}, [update
        animations and send
        events](https://drafts.csswg.org/web-animations-1/#update-animations-and-send-events){#event-loop-processing-model:update-animations-and-send-events
        x-internal="update-animations-and-send-events"} for
        `doc`{.variable}, passing in [relative high resolution
        time](https://w3c.github.io/hr-time/#dfn-relative-high-resolution-time){#event-loop-processing-model:relative-high-resolution-time
        x-internal="relative-high-resolution-time"} given
        `frameTimestamp`{.variable} and `doc`{.variable}\'s [relevant
        global
        object](#concept-relevant-global){#event-loop-processing-model:concept-relevant-global}
        as the timestamp
        [\[WEBANIMATIONS\]](references.html#refsWEBANIMATIONS)

    12. For each `doc`{.variable} of `docs`{.variable}, [run the
        fullscreen
        steps](https://fullscreen.spec.whatwg.org/#run-the-fullscreen-steps){#event-loop-processing-model:run-the-fullscreen-steps
        x-internal="run-the-fullscreen-steps"} for `doc`{.variable}.
        [\[FULLSCREEN\]](references.html#refsFULLSCREEN)

    13. For each `doc`{.variable} of `docs`{.variable}, if the user
        agent detects that the backing storage associated with a
        [`CanvasRenderingContext2D`{#event-loop-processing-model:canvasrenderingcontext2d}](canvas.html#canvasrenderingcontext2d)
        or an
        [`OffscreenCanvasRenderingContext2D`{#event-loop-processing-model:offscreencanvasrenderingcontext2d}](canvas.html#offscreencanvasrenderingcontext2d),
        `context`{.variable}, has been lost, then it must run the
        [context lost steps]{#context-lost-steps .dfn} for each such
        `context`{.variable}:

        1.  Let `canvas`{.variable} be the value of
            `context`{.variable}\'s
            [`canvas`{#event-loop-processing-model:dom-context-2d-canvas}](canvas.html#dom-context-2d-canvas)
            attribute, if `context`{.variable} is a
            [`CanvasRenderingContext2D`{#event-loop-processing-model:canvasrenderingcontext2d-2}](canvas.html#canvasrenderingcontext2d),
            or the [associated `OffscreenCanvas`
            object](canvas.html#associated-offscreencanvas-object){#event-loop-processing-model:associated-offscreencanvas-object}
            for `context`{.variable} otherwise.

        2.  Set `context`{.variable}\'s [context
            lost](canvas.html#concept-canvas-context-lost){#event-loop-processing-model:concept-canvas-context-lost}
            to true.

        3.  [Reset the rendering context to its default
            state](canvas.html#reset-the-rendering-context-to-its-default-state){#event-loop-processing-model:reset-the-rendering-context-to-its-default-state}
            given `context`{.variable}.

        4.  Let `shouldRestore`{.variable} be the result of [firing an
            event](https://dom.spec.whatwg.org/#concept-event-fire){#event-loop-processing-model:concept-event-fire
            x-internal="concept-event-fire"} named
            [`contextlost`{#event-loop-processing-model:event-contextlost}](indices.html#event-contextlost)
            at `canvas`{.variable}, with the
            [`cancelable`{#event-loop-processing-model:dom-event-cancelable}](https://dom.spec.whatwg.org/#dom-event-cancelable){x-internal="dom-event-cancelable"}
            attribute initialized to true.

        5.  If `shouldRestore`{.variable} is false, then abort these
            steps.

        6.  Attempt to restore `context`{.variable} by creating a
            backing storage using `context`{.variable}\'s attributes and
            associating them with `context`{.variable}. If this fails,
            then abort these steps.

        7.  Set `context`{.variable}\'s [context
            lost](canvas.html#concept-canvas-context-lost){#event-loop-processing-model:concept-canvas-context-lost-2}
            to false.

        8.  [Fire an
            event](https://dom.spec.whatwg.org/#concept-event-fire){#event-loop-processing-model:concept-event-fire-2
            x-internal="concept-event-fire"} named
            [`contextrestored`{#event-loop-processing-model:event-contextrestored}](indices.html#event-contextrestored)
            at `canvas`{.variable}.

    14. For each `doc`{.variable} of `docs`{.variable}, [run the
        animation frame
        callbacks](imagebitmap-and-animations.html#run-the-animation-frame-callbacks){#event-loop-processing-model:run-the-animation-frame-callbacks-3}
        for `doc`{.variable}, passing in the [relative high resolution
        time](https://w3c.github.io/hr-time/#dfn-relative-high-resolution-time){#event-loop-processing-model:relative-high-resolution-time-2
        x-internal="relative-high-resolution-time"} given
        `frameTimestamp`{.variable} and `doc`{.variable}\'s [relevant
        global
        object](#concept-relevant-global){#event-loop-processing-model:concept-relevant-global-2}
        as the timestamp.

    15. Let `unsafeStyleAndLayoutStartTime`{.variable} be the [unsafe
        shared current
        time](https://w3c.github.io/hr-time/#dfn-unsafe-shared-current-time){#event-loop-processing-model:unsafe-shared-current-time-5
        x-internal="unsafe-shared-current-time"}.

    16. For each `doc`{.variable} of `docs`{.variable}:

        1.  Let `resizeObserverDepth`{.variable} be 0.

        2.  While true:

            1.  Recalculate styles and update layout for
                `doc`{.variable}.

            2.  Let
                `hadInitialVisibleContentVisibilityDetermination`{.variable}
                be false.

            3.  For each element `element`{.variable} with
                [\'auto\'](https://drafts.csswg.org/css-contain/#propdef-content-visibility){#event-loop-processing-model:content-visibility-auto
                x-internal="content-visibility-auto"} used value of
                [\'content-visibility\'](https://drafts.csswg.org/css-contain/#content-visibility){#event-loop-processing-model:'content-visibility'
                x-internal="'content-visibility'"}:

                1.  Let `checkForInitialDetermination`{.variable} be
                    true if `element`{.variable}\'s [proximity to the
                    viewport](https://drafts.csswg.org/css-contain/#proximity-to-the-viewport){#event-loop-processing-model:proximity-to-the-viewport
                    x-internal="proximity-to-the-viewport"} is not
                    determined and it is not [relevant to the
                    user](https://drafts.csswg.org/css-contain/#relevant-to-the-user){#event-loop-processing-model:relevant-to-the-user
                    x-internal="relevant-to-the-user"}. Otherwise, let
                    `checkForInitialDetermination`{.variable} be false.

                2.  Determine [proximity to the
                    viewport](https://drafts.csswg.org/css-contain/#proximity-to-the-viewport){#event-loop-processing-model:proximity-to-the-viewport-2
                    x-internal="proximity-to-the-viewport"} for
                    `element`{.variable}.

                3.  If `checkForInitialDetermination`{.variable} is true
                    and `element`{.variable} is now [relevant to the
                    user](https://drafts.csswg.org/css-contain/#relevant-to-the-user){#event-loop-processing-model:relevant-to-the-user-2
                    x-internal="relevant-to-the-user"}, then set
                    `hadInitialVisibleContentVisibilityDetermination`{.variable}
                    to true.

            4.  If
                `hadInitialVisibleContentVisibilityDetermination`{.variable}
                is true, then
                [continue](https://infra.spec.whatwg.org/#iteration-continue){#event-loop-processing-model:continue-3
                x-internal="continue"}.

                The intent of this step is for the initial viewport
                proximity determination, which takes effect immediately,
                to be reflected in the style and layout calculation
                which is carried out in a previous step of this loop.
                Proximity determinations other than the initial one take
                effect at the next [rendering
                opportunity](#rendering-opportunity){#event-loop-processing-model:rendering-opportunity-5}.
                [\[CSSCONTAIN\]](references.html#refsCSSCONTAIN)

            5.  [Gather active resize observations at
                depth](https://drafts.csswg.org/resize-observer-1/#gather-active-observations-h){#event-loop-processing-model:gather-active-resize-observations-at-depth
                x-internal="gather-active-resize-observations-at-depth"}
                `resizeObserverDepth`{.variable} for `doc`{.variable}.

            6.  If `doc`{.variable} [has active resize
                observations](https://drafts.csswg.org/resize-observer-1/#has-active-observations-h){#event-loop-processing-model:has-active-resize-observations
                x-internal="has-active-resize-observations"}:

                1.  Set `resizeObserverDepth`{.variable} to the result
                    of [broadcasting active resize
                    observations](https://drafts.csswg.org/resize-observer-1/#broadcast-resize-notifications-h){#event-loop-processing-model:broadcast-active-resize-observations
                    x-internal="broadcast-active-resize-observations"}
                    given `doc`{.variable}.
                2.  [Continue](https://infra.spec.whatwg.org/#iteration-continue){#event-loop-processing-model:continue-4
                    x-internal="continue"}.

            7.  Otherwise,
                [break](https://infra.spec.whatwg.org/#iteration-break){#event-loop-processing-model:break
                x-internal="break"}.

        3.  If `doc`{.variable} [has skipped resize
            observations](https://drafts.csswg.org/resize-observer-1/#has-skipped-observations-h){#event-loop-processing-model:has-skipped-resize-observations
            x-internal="has-skipped-resize-observations"}, then [deliver
            resize loop
            error](https://drafts.csswg.org/resize-observer-1/#deliver-resize-error){#event-loop-processing-model:deliver-resize-loop-error
            x-internal="deliver-resize-loop-error"} given
            `doc`{.variable}.

    17. ::: {#focus-fixup-rule}
        For each `doc`{.variable} of `docs`{.variable}, if the [focused
        area](interaction.html#focused-area-of-the-document){#event-loop-processing-model:focused-area-of-the-document}
        of `doc`{.variable} is not a [focusable
        area](interaction.html#focusable-area){#event-loop-processing-model:focusable-area},
        then run the [focusing
        steps](interaction.html#focusing-steps){#event-loop-processing-model:focusing-steps}
        for `doc`{.variable}\'s
        [viewport](https://drafts.csswg.org/css2/#viewport){#event-loop-processing-model:viewport
        x-internal="viewport"}, and set `doc`{.variable}\'s [relevant
        global
        object](#concept-relevant-global){#event-loop-processing-model:concept-relevant-global-3}\'s
        [navigation
        API](nav-history-apis.html#window-navigation-api){#event-loop-processing-model:window-navigation-api}\'s
        [focus changed during ongoing
        navigation](nav-history-apis.html#focus-changed-during-ongoing-navigation){#event-loop-processing-model:focus-changed-during-ongoing-navigation}
        to false.

        For example, this might happen because an element has the
        [`hidden`{#event-loop-processing-model:attr-hidden}](interaction.html#attr-hidden)
        attribute added, causing it to stop [being
        rendered](rendering.html#being-rendered){#event-loop-processing-model:being-rendered}.
        It might also happen to an
        [`input`{#event-loop-processing-model:the-input-element}](input.html#the-input-element)
        element when the element gets
        [disabled](form-control-infrastructure.html#concept-fe-disabled){#event-loop-processing-model:concept-fe-disabled}.

        This will
        [usually](interaction.html#note-sometimes-no-blur-event) fire
        [`blur`{#event-loop-processing-model:event-blur}](indices.html#event-blur)
        events, and possibly
        [`change`{#event-loop-processing-model:event-change}](indices.html#event-change)
        events.

        In addition to this asynchronous fixup, if the [focused area of
        the
        document](interaction.html#focused-area-of-the-document){#event-loop-processing-model:focused-area-of-the-document-2}
        is removed, there is a [synchronous
        fixup](infrastructure.html#node-remove-focus-fixup). That one
        will *not* fire
        [`blur`{#event-loop-processing-model:event-blur-2}](indices.html#event-blur)
        or
        [`change`{#event-loop-processing-model:event-change-2}](indices.html#event-change)
        events.
        :::

    18. For each `doc`{.variable} of `docs`{.variable}, [perform pending
        transition
        operations](https://drafts.csswg.org/css-view-transitions/#perform-pending-transition-operations){#event-loop-processing-model:perform-pending-transition-operations
        x-internal="perform-pending-transition-operations"} for
        `doc`{.variable}.
        [\[CSSVIEWTRANSITIONS\]](references.html#refsCSSVIEWTRANSITIONS)

    19. For each `doc`{.variable} of `docs`{.variable}, [run the update
        intersection observations
        steps](https://w3c.github.io/IntersectionObserver/#run-the-update-intersection-observations-steps){#event-loop-processing-model:run-the-update-intersection-observations-steps
        x-internal="run-the-update-intersection-observations-steps"} for
        `doc`{.variable}, passing in the [relative high resolution
        time](https://w3c.github.io/hr-time/#dfn-relative-high-resolution-time){#event-loop-processing-model:relative-high-resolution-time-3
        x-internal="relative-high-resolution-time"} given
        `now`{.variable} and `doc`{.variable}\'s [relevant global
        object](#concept-relevant-global){#event-loop-processing-model:concept-relevant-global-4}
        as the timestamp.
        [\[INTERSECTIONOBSERVER\]](references.html#refsINTERSECTIONOBSERVER)

    20. For each `doc`{.variable} of `docs`{.variable}, [record
        rendering
        time](https://w3c.github.io/long-animation-frames/#record-rendering-time){#event-loop-processing-model:record-rendering-time
        x-internal="record-rendering-time"} for `doc`{.variable} given
        `unsafeStyleAndLayoutStartTime`{.variable}.

    21. For each `doc`{.variable} of `docs`{.variable}, [mark paint
        timing](https://w3c.github.io/paint-timing/#mark-paint-timing){#event-loop-processing-model:mark-paint-timing
        x-internal="mark-paint-timing"} for `doc`{.variable}.

    22. For each `doc`{.variable} of `docs`{.variable}, update the
        rendering or user interface of `doc`{.variable} and its [node
        navigable](document-sequences.html#node-navigable){#event-loop-processing-model:node-navigable-4}
        to reflect the current state.

    23. For each `doc`{.variable} of `docs`{.variable}, [process top
        layer
        removals](https://drafts.csswg.org/css-position-4/#process-top-layer-removals){#event-loop-processing-model:process-top-layer-removals
        x-internal="process-top-layer-removals"} given `doc`{.variable}.

A
[navigable](document-sequences.html#navigable){#event-loop-processing-model:navigable-2}
has a [rendering opportunity]{#rendering-opportunity .dfn export=""} if
the user agent is currently able to present the contents of the
[navigable](document-sequences.html#navigable){#event-loop-processing-model:navigable-3}
to the user, accounting for hardware refresh rate constraints and user
agent throttling for performance reasons, but considering content
presentable even if it\'s outside the viewport.

A
[navigable](document-sequences.html#navigable){#event-loop-processing-model:navigable-4}\'s
[rendering
opportunities](#rendering-opportunity){#event-loop-processing-model:rendering-opportunity-6}
are determined based on hardware constraints such as display refresh
rates and other factors such as page performance or whether its [active
document](document-sequences.html#nav-document){#event-loop-processing-model:nav-document-2}\'s
[visibility
state](interaction.html#visibility-state){#event-loop-processing-model:visibility-state-2}
is \"`visible`\". Rendering opportunities typically occur at regular
intervals.

This specification does not mandate any particular model for selecting
rendering opportunities. But for example, if the browser is attempting
to achieve a 60Hz refresh rate, then rendering opportunities occur at a
maximum of every 60th of a second (about 16.7ms). If the browser finds
that a
[navigable](document-sequences.html#navigable){#event-loop-processing-model:navigable-5}
is not able to sustain this rate, it might drop to a more sustainable 30
rendering opportunities per second for that
[navigable](document-sequences.html#navigable){#event-loop-processing-model:navigable-6},
rather than occasionally dropping frames. Similarly, if a
[navigable](document-sequences.html#navigable){#event-loop-processing-model:navigable-7}
is not visible, the user agent might decide to drop that page to a much
slower 4 rendering opportunities per second, or even less.

------------------------------------------------------------------------

When a user agent is to [perform a microtask
checkpoint]{#perform-a-microtask-checkpoint .dfn export=""}:

1.  If the [event
    loop](#event-loop){#event-loop-processing-model:event-loop-15}\'s
    [performing a microtask
    checkpoint](#performing-a-microtask-checkpoint){#event-loop-processing-model:performing-a-microtask-checkpoint}
    is true, then return.

2.  Set the [event
    loop](#event-loop){#event-loop-processing-model:event-loop-16}\'s
    [performing a microtask
    checkpoint](#performing-a-microtask-checkpoint){#event-loop-processing-model:performing-a-microtask-checkpoint-2}
    to true.

3.  While the [event
    loop](#event-loop){#event-loop-processing-model:event-loop-17}\'s
    [microtask
    queue](#microtask-queue){#event-loop-processing-model:microtask-queue-2}
    is not
    [empty](https://infra.spec.whatwg.org/#list-is-empty){#event-loop-processing-model:list-is-empty
    x-internal="list-is-empty"}:

    1.  Let `oldestMicrotask`{.variable} be the result of
        [dequeuing](https://infra.spec.whatwg.org/#queue-dequeue){#event-loop-processing-model:dequeue
        x-internal="dequeue"} from the [event
        loop](#event-loop){#event-loop-processing-model:event-loop-18}\'s
        [microtask
        queue](#microtask-queue){#event-loop-processing-model:microtask-queue-3}.

    2.  Set the [event
        loop](#event-loop){#event-loop-processing-model:event-loop-19}\'s
        [currently running
        task](#currently-running-task){#event-loop-processing-model:currently-running-task-3}
        to `oldestMicrotask`{.variable}.

    3.  Run `oldestMicrotask`{.variable}.

        This might involve invoking scripted callbacks, which eventually
        calls the [clean up after running
        script](#clean-up-after-running-script){#event-loop-processing-model:clean-up-after-running-script}
        steps, which call this [perform a microtask
        checkpoint](#perform-a-microtask-checkpoint){#event-loop-processing-model:perform-a-microtask-checkpoint-3}
        algorithm again, which is why we use the [performing a microtask
        checkpoint](#performing-a-microtask-checkpoint){#event-loop-processing-model:performing-a-microtask-checkpoint-3}
        flag to avoid reentrancy.

    4.  Set the [event
        loop](#event-loop){#event-loop-processing-model:event-loop-20}\'s
        [currently running
        task](#currently-running-task){#event-loop-processing-model:currently-running-task-4}
        back to null.

4.  For each [environment settings
    object](#environment-settings-object){#event-loop-processing-model:environment-settings-object-2}
    `settingsObject`{.variable} whose [responsible event
    loop](#responsible-event-loop){#event-loop-processing-model:responsible-event-loop}
    is this [event
    loop](#event-loop){#event-loop-processing-model:event-loop-21},
    [notify about rejected
    promises](#notify-about-rejected-promises){#event-loop-processing-model:notify-about-rejected-promises}
    given `settingsObject`{.variable}\'s [global
    object](#concept-settings-object-global){#event-loop-processing-model:concept-settings-object-global-2}.

5.  [Cleanup Indexed Database
    transactions](https://w3c.github.io/IndexedDB/#cleanup-indexed-database-transactions){#event-loop-processing-model:cleanup-indexed-database-transactions
    x-internal="cleanup-indexed-database-transactions"}.

6.  Perform
    [ClearKeptObjects](https://tc39.es/ecma262/#sec-clear-kept-objects){#event-loop-processing-model:clearkeptobjects
    x-internal="clearkeptobjects"}().

    When
    [`WeakRef.prototype.deref()`{#event-loop-processing-model:weakref.prototype.deref()}](https://tc39.es/ecma262/#sec-weak-ref.prototype.deref){x-internal="weakref.prototype.deref()"}
    returns an object, that object is kept alive until the next
    invocation of
    [ClearKeptObjects](https://tc39.es/ecma262/#sec-clear-kept-objects){#event-loop-processing-model:clearkeptobjects-2
    x-internal="clearkeptobjects"}(), after which it is again subject to
    garbage collection.

7.  Set the [event
    loop](#event-loop){#event-loop-processing-model:event-loop-22}\'s
    [performing a microtask
    checkpoint](#performing-a-microtask-checkpoint){#event-loop-processing-model:performing-a-microtask-checkpoint-4}
    to false.

8.  [Record timing info for microtask
    checkpoint](https://w3c.github.io/long-animation-frames/#record-timing-info-for-microtask-checkpoint){#event-loop-processing-model:record-timing-info-for-microtask-checkpoint
    x-internal="record-timing-info-for-microtask-checkpoint"}.

------------------------------------------------------------------------

When an algorithm running [in
parallel](infrastructure.html#in-parallel){#event-loop-processing-model:in-parallel-3}
is to [await a stable state]{#await-a-stable-state .dfn}, the user agent
must [queue a
microtask](#queue-a-microtask){#event-loop-processing-model:queue-a-microtask}
that runs the following steps, and must then stop executing (execution
of the algorithm resumes when the microtask is run, as described in the
following steps):

1.  Run the algorithm\'s [synchronous section]{#synchronous-section
    .dfn}.

2.  Resume execution of the algorithm [in
    parallel](infrastructure.html#in-parallel){#event-loop-processing-model:in-parallel-4},
    if appropriate, as described in the algorithm\'s steps.

Steps in [synchronous
sections](#synchronous-section){#event-loop-processing-model:synchronous-section}
are marked with ⌛.

------------------------------------------------------------------------

Algorithm steps that say to [spin the event loop]{#spin-the-event-loop
.dfn} until a condition `goal`{.variable} is met are equivalent to
substituting in the following algorithm steps:

1.  Let `task`{.variable} be the [event
    loop](#event-loop){#event-loop-processing-model:event-loop-23}\'s
    [currently running
    task](#currently-running-task){#event-loop-processing-model:currently-running-task-5}.

    `task`{.variable} could be a
    [microtask](#microtask){#event-loop-processing-model:microtask-2}.

2.  Let `task source`{.variable} be `task`{.variable}\'s
    [source](#concept-task-source){#event-loop-processing-model:concept-task-source}.

3.  Let `old stack`{.variable} be a copy of the [JavaScript execution
    context
    stack](https://tc39.es/ecma262/#execution-context-stack){#event-loop-processing-model:javascript-execution-context-stack
    x-internal="javascript-execution-context-stack"}.

4.  Empty the [JavaScript execution context
    stack](https://tc39.es/ecma262/#execution-context-stack){#event-loop-processing-model:javascript-execution-context-stack-2
    x-internal="javascript-execution-context-stack"}.

5.  [Perform a microtask
    checkpoint](#perform-a-microtask-checkpoint){#event-loop-processing-model:perform-a-microtask-checkpoint-4}.

    If `task`{.variable} is a
    [microtask](#microtask){#event-loop-processing-model:microtask-3}
    this step will be a no-op due to [performing a microtask
    checkpoint](#performing-a-microtask-checkpoint){#event-loop-processing-model:performing-a-microtask-checkpoint-5}
    being true.

6.  [In
    parallel](infrastructure.html#in-parallel){#event-loop-processing-model:in-parallel-5}:

    1.  Wait until the condition `goal`{.variable} is met.

    2.  [Queue a
        task](#queue-a-task){#event-loop-processing-model:queue-a-task}
        on `task source`{.variable} to:

        1.  Replace the [JavaScript execution context
            stack](https://tc39.es/ecma262/#execution-context-stack){#event-loop-processing-model:javascript-execution-context-stack-3
            x-internal="javascript-execution-context-stack"} with
            `old stack`{.variable}.

        2.  Perform any steps that appear after this [spin the event
            loop](#spin-the-event-loop){#event-loop-processing-model:spin-the-event-loop-2}
            instance in the original algorithm.

            This resumes `task`{.variable}.

7.  Stop `task`{.variable}, allowing whatever algorithm that invoked it
    to resume.

    This causes the [event
    loop](#event-loop){#event-loop-processing-model:event-loop-24}\'s
    main set of steps or the [perform a microtask
    checkpoint](#perform-a-microtask-checkpoint){#event-loop-processing-model:perform-a-microtask-checkpoint-5}
    algorithm to continue.

Unlike other algorithms in this and other specifications, which behave
similar to programming-language function calls, [spin the event
loop](#spin-the-event-loop){#event-loop-processing-model:spin-the-event-loop-3}
is more like a macro, which saves typing and indentation at the usage
site by expanding into a series of steps and operations.

::: example
An algorithm whose steps are:

1.  Do something.

2.  [Spin the event
    loop](#spin-the-event-loop){#event-loop-processing-model:spin-the-event-loop-4}
    until awesomeness happens.

3.  Do something else.

is a shorthand which, after \"macro expansion\", becomes

1.  Do something.

2.  Let `old stack`{.variable} be a copy of the [JavaScript execution
    context
    stack](https://tc39.es/ecma262/#execution-context-stack){#event-loop-processing-model:javascript-execution-context-stack-4
    x-internal="javascript-execution-context-stack"}.

3.  Empty the [JavaScript execution context
    stack](https://tc39.es/ecma262/#execution-context-stack){#event-loop-processing-model:javascript-execution-context-stack-5
    x-internal="javascript-execution-context-stack"}.

4.  [Perform a microtask
    checkpoint](#perform-a-microtask-checkpoint){#event-loop-processing-model:perform-a-microtask-checkpoint-6}.

5.  [In
    parallel](infrastructure.html#in-parallel){#event-loop-processing-model:in-parallel-6}:

    1.  Wait until awesomeness happens.

    2.  [Queue a
        task](#queue-a-task){#event-loop-processing-model:queue-a-task-2}
        on the task source in which \"do something\" was done to:

        1.  Replace the [JavaScript execution context
            stack](https://tc39.es/ecma262/#execution-context-stack){#event-loop-processing-model:javascript-execution-context-stack-6
            x-internal="javascript-execution-context-stack"} with
            `old stack`{.variable}.

        2.  Do something else.
:::

::: example
Here is a more full example of the substitution, where the event loop is
spun from inside a task that is queued from work in parallel. The
version using [spin the event
loop](#spin-the-event-loop){#event-loop-processing-model:spin-the-event-loop-5}:

1.  [In
    parallel](infrastructure.html#in-parallel){#event-loop-processing-model:in-parallel-7}:

    1.  Do parallel thing 1.

    2.  [Queue a
        task](#queue-a-task){#event-loop-processing-model:queue-a-task-3}
        on the [DOM manipulation task
        source](#dom-manipulation-task-source){#event-loop-processing-model:dom-manipulation-task-source}
        to:

        1.  Do task thing 1.

        2.  [Spin the event
            loop](#spin-the-event-loop){#event-loop-processing-model:spin-the-event-loop-6}
            until awesomeness happens.

        3.  Do task thing 2.

    3.  Do parallel thing 2.

The fully expanded version:

1.  [In
    parallel](infrastructure.html#in-parallel){#event-loop-processing-model:in-parallel-8}:

    1.  Do parallel thing 1.

    2.  Let `old stack`{.variable} be null.

    3.  [Queue a
        task](#queue-a-task){#event-loop-processing-model:queue-a-task-4}
        on the [DOM manipulation task
        source](#dom-manipulation-task-source){#event-loop-processing-model:dom-manipulation-task-source-2}
        to:

        1.  Do task thing 1.

        2.  Set `old stack`{.variable} to a copy of the [JavaScript
            execution context
            stack](https://tc39.es/ecma262/#execution-context-stack){#event-loop-processing-model:javascript-execution-context-stack-7
            x-internal="javascript-execution-context-stack"}.

        3.  Empty the [JavaScript execution context
            stack](https://tc39.es/ecma262/#execution-context-stack){#event-loop-processing-model:javascript-execution-context-stack-8
            x-internal="javascript-execution-context-stack"}.

        4.  [Perform a microtask
            checkpoint](#perform-a-microtask-checkpoint){#event-loop-processing-model:perform-a-microtask-checkpoint-7}.

    4.  Wait until awesomeness happens.

    5.  [Queue a
        task](#queue-a-task){#event-loop-processing-model:queue-a-task-5}
        on the [DOM manipulation task
        source](#dom-manipulation-task-source){#event-loop-processing-model:dom-manipulation-task-source-3}
        to:

        1.  Replace the [JavaScript execution context
            stack](https://tc39.es/ecma262/#execution-context-stack){#event-loop-processing-model:javascript-execution-context-stack-9
            x-internal="javascript-execution-context-stack"} with
            `old stack`{.variable}.

        2.  Do task thing 2.

    6.  Do parallel thing 2.
:::

------------------------------------------------------------------------

Some of the algorithms in this specification, for historical reasons,
require the user agent to [pause]{#pause .dfn export=""} while running a
[task](#concept-task){#event-loop-processing-model:concept-task-7} until
a condition `goal`{.variable} is met. This means running the following
steps:

1.  Let `global`{.variable} be the [current global
    object](#current-global-object){#event-loop-processing-model:current-global-object}.

2.  Let `timeBeforePause`{.variable} be the [current high resolution
    time](https://w3c.github.io/hr-time/#dfn-current-high-resolution-time){#event-loop-processing-model:current-high-resolution-time-2
    x-internal="current-high-resolution-time"} given
    `global`{.variable}.

3.  If necessary, update the rendering or user interface of any
    [`Document`{#event-loop-processing-model:document-7}](dom.html#document)
    or
    [navigable](document-sequences.html#navigable){#event-loop-processing-model:navigable-8}
    to reflect the current state.

4.  Wait until the condition `goal`{.variable} is met. While a user
    agent has a paused
    [task](#concept-task){#event-loop-processing-model:concept-task-8},
    the corresponding [event
    loop](#event-loop){#event-loop-processing-model:event-loop-25} must
    not run further
    [tasks](#concept-task){#event-loop-processing-model:concept-task-9},
    and any script in the currently running
    [task](#concept-task){#event-loop-processing-model:concept-task-10}
    must block. User agents should remain responsive to user input while
    paused, however, albeit in a reduced capacity since the [event
    loop](#event-loop){#event-loop-processing-model:event-loop-26} will
    not be doing anything.

5.  [Record pause
    duration](https://w3c.github.io/long-animation-frames/#record-pause-duration){#event-loop-processing-model:record-pause-duration
    x-internal="record-pause-duration"} given the [duration
    from](https://w3c.github.io/hr-time/#dfn-duration-from){#event-loop-processing-model:duration-from
    x-internal="duration-from"} `timeBeforePause`{.variable} to the
    [current high resolution
    time](https://w3c.github.io/hr-time/#dfn-current-high-resolution-time){#event-loop-processing-model:current-high-resolution-time-3
    x-internal="current-high-resolution-time"} given
    `global`{.variable}.

::: warning
[Pausing](#pause){#event-loop-processing-model:pause} is highly
detrimental to the user experience, especially in scenarios where a
single [event
loop](#event-loop){#event-loop-processing-model:event-loop-27} is shared
among multiple documents. User agents are encouraged to experiment with
alternatives to [pausing](#pause){#event-loop-processing-model:pause-2},
such as [spinning the event
loop](#spin-the-event-loop){#event-loop-processing-model:spin-the-event-loop-7}
or even simply proceeding without any kind of suspended execution at
all, insofar as it is possible to do so while preserving compatibility
with existing content. This specification will happily change if a
less-drastic alternative is discovered to be web-compatible.

In the interim, implementers should be aware that the variety of
alternatives that user agents might experiment with can change subtle
aspects of [event
loop](#event-loop){#event-loop-processing-model:event-loop-28} behavior,
including
[task](#concept-task){#event-loop-processing-model:concept-task-11} and
[microtask](#microtask){#event-loop-processing-model:microtask-4}
timing. Implementations should continue experimenting even if doing so
causes them to violate the exact semantics implied by the
[pause](#pause){#event-loop-processing-model:pause-3} operation.
:::

##### [8.1.7.4]{.secno} Generic task sources[](#generic-task-sources){.self-link}

The following [task
sources](#task-source){#generic-task-sources:task-source} are used by a
number of mostly unrelated features in this and other specifications.

The [DOM manipulation task source]{#dom-manipulation-task-source .dfn export=""}

:   This [task
    source](#task-source){#generic-task-sources:task-source-2} is used
    for features that react to DOM manipulations, such as things that
    happen in a non-blocking fashion when an element is [inserted into
    the
    document](infrastructure.html#insert-an-element-into-a-document){#generic-task-sources:insert-an-element-into-a-document}.

The [user interaction task source]{#user-interaction-task-source .dfn export=""}

:   This [task
    source](#task-source){#generic-task-sources:task-source-3} is used
    for features that react to user interaction, for example keyboard or
    mouse input.

    Events sent in response to user input (e.g.,
    [`click`{#generic-task-sources:event-click}](https://w3c.github.io/uievents/#event-type-click){x-internal="event-click"}
    events) must be fired using
    [tasks](#concept-task){#generic-task-sources:concept-task}
    [queued](#queue-a-task){#generic-task-sources:queue-a-task} with the
    [user interaction task
    source](#user-interaction-task-source){#generic-task-sources:user-interaction-task-source}.
    [\[UIEVENTS\]](references.html#refsUIEVENTS)

The [networking task source]{#networking-task-source .dfn export=""}

:   This [task
    source](#task-source){#generic-task-sources:task-source-4} is used
    for features that trigger in response to network activity.

The [navigation and traversal task source]{#navigation-and-traversal-task-source .dfn}

:   This [task
    source](#task-source){#generic-task-sources:task-source-5} is used
    to queue tasks involved in
    [navigation](browsing-the-web.html#navigate){#generic-task-sources:navigate}
    and [history
    traversal](browsing-the-web.html#apply-the-traverse-history-step){#generic-task-sources:apply-the-traverse-history-step}.

The [rendering task source]{#rendering-task-source .dfn}

:   This [task
    source](#task-source){#generic-task-sources:task-source-6} is used
    solely to [update the
    rendering](#update-the-rendering){#generic-task-sources:update-the-rendering}.

##### [8.1.7.5]{.secno} Dealing with the event loop from other specifications[](#event-loop-for-spec-authors){.self-link} {#event-loop-for-spec-authors}

Writing specifications that correctly interact with the [event
loop](#event-loop){#event-loop-for-spec-authors:event-loop} can be
tricky. This is compounded by how this specification uses
concurrency-model-independent terminology, so we say things like
\"[event loop](#event-loop){#event-loop-for-spec-authors:event-loop-2}\"
and \"[in
parallel](infrastructure.html#in-parallel){#event-loop-for-spec-authors:in-parallel}\"
instead of using more familiar model-specific terms like \"main thread\"
or \"on a background thread\".

By default, specification text generally runs on the [event
loop](#event-loop){#event-loop-for-spec-authors:event-loop-3}. This
falls out from the formal [event loop processing
model](#event-loop-processing-model), in that you can eventually trace
most algorithms back to a
[task](#concept-task){#event-loop-for-spec-authors:concept-task}
[queued](#queue-a-task){#event-loop-for-spec-authors:queue-a-task}
there.

The algorithm steps for any JavaScript method will be invoked by author
code calling that method. And author code can only be run via queued
tasks, usually originating somewhere in the [`script` processing
model](scripting.html#script-processing-model).

From this starting point, the overriding guideline is that any work a
specification needs to perform that would otherwise block the [event
loop](#event-loop){#event-loop-for-spec-authors:event-loop-4} must
instead be performed [in
parallel](infrastructure.html#in-parallel){#event-loop-for-spec-authors:in-parallel-2}
with it. This includes (but is not limited to):

- performing heavy computation;

- displaying a user-facing prompt;

- performing operations which could require involving outside systems
  (i.e. \"going out of process\").

The next complication is that, in algorithm sections that are [in
parallel](infrastructure.html#in-parallel){#event-loop-for-spec-authors:in-parallel-3},
you must not create or manipulate objects associated to a specific
[realm](https://tc39.es/ecma262/#sec-code-realms){#event-loop-for-spec-authors:realm
x-internal="realm"},
[global](#global-object){#event-loop-for-spec-authors:global-object}, or
[environment settings
object](#environment-settings-object){#event-loop-for-spec-authors:environment-settings-object}.
(Stated in more familiar terms, you must not directly access main-thread
artifacts from a background thread.) Doing so would create data races
observable to JavaScript code, since after all, your algorithm steps are
running *[in
parallel](infrastructure.html#in-parallel){#event-loop-for-spec-authors:in-parallel-4}*
to the JavaScript code.

By extension, you cannot access Web IDL\'s
[this](https://webidl.spec.whatwg.org/#this){#event-loop-for-spec-authors:this
x-internal="this"} value from steps running [in
parallel](infrastructure.html#in-parallel){#event-loop-for-spec-authors:in-parallel-5},
even if those steps were activated by an algorithm that *does* have
access to the
[this](https://webidl.spec.whatwg.org/#this){#event-loop-for-spec-authors:this-2
x-internal="this"} value.

You can, however, manipulate specification-level data structures and
values from Infra, as those are realm-agnostic. They are never directly
exposed to JavaScript without a specific conversion taking place (often
[via Web
IDL](https://webidl.spec.whatwg.org/#es-type-mapping){#event-loop-for-spec-authors:concept-idl-convert
x-internal="concept-idl-convert"}).
[\[INFRA\]](references.html#refsINFRA)
[\[WEBIDL\]](references.html#refsWEBIDL)

To affect the world of observable JavaScript objects, then, you must
[queue a global
task](#queue-a-global-task){#event-loop-for-spec-authors:queue-a-global-task}
to perform any such manipulations. This ensures your steps are properly
interleaved with respect to other things happening on the [event
loop](#event-loop){#event-loop-for-spec-authors:event-loop-5}.
Furthermore, you must choose a [task
source](#task-source){#event-loop-for-spec-authors:task-source} when
[queuing a global
task](#queue-a-global-task){#event-loop-for-spec-authors:queue-a-global-task-2};
this governs the relative order of your steps versus others. If you are
unsure which [task
source](#task-source){#event-loop-for-spec-authors:task-source-2} to
use, pick one of the [generic task sources](#generic-task-sources) that
sounds most applicable. Finally, you must indicate which [global
object](#global-object){#event-loop-for-spec-authors:global-object-2}
your queued task is associated with; this ensures that if that global
object is inactive, the task does not run.

The base primitive, on which [queue a global
task](#queue-a-global-task){#event-loop-for-spec-authors:queue-a-global-task-3}
builds, is the [queue a
task](#queue-a-task){#event-loop-for-spec-authors:queue-a-task-2}
algorithm. In general, [queue a global
task](#queue-a-global-task){#event-loop-for-spec-authors:queue-a-global-task-4}
is better because it automatically picks the right [event
loop](#event-loop){#event-loop-for-spec-authors:event-loop-6} and, where
appropriate,
[document](#concept-task-document){#event-loop-for-spec-authors:concept-task-document}.
Older specifications often use [queue a
task](#queue-a-task){#event-loop-for-spec-authors:queue-a-task-3}
combined with the [implied event
loop](#implied-event-loop){#event-loop-for-spec-authors:implied-event-loop}
and [implied
document](#implied-document){#event-loop-for-spec-authors:implied-document}
concepts, but this is discouraged.

Putting this all together, we can provide a template for a typical
algorithm that needs to do work asynchronously:

1.  Do any synchronous setup work, while still on the [event
    loop](#event-loop){#event-loop-for-spec-authors:event-loop-7}. This
    may include converting
    [realm](https://tc39.es/ecma262/#sec-code-realms){#event-loop-for-spec-authors:realm-2
    x-internal="realm"}-specific JavaScript values into realm-agnostic
    specification-level values.

2.  Perform a set of potentially-expensive steps [in
    parallel](infrastructure.html#in-parallel){#event-loop-for-spec-authors:in-parallel-6},
    operating entirely on realm-agnostic values, and producing a
    realm-agnostic result.

3.  [Queue a global
    task](#queue-a-global-task){#event-loop-for-spec-authors:queue-a-global-task-5},
    on a specified [task
    source](#task-source){#event-loop-for-spec-authors:task-source-3}
    and given an appropriate [global
    object](#global-object){#event-loop-for-spec-authors:global-object-3},
    to convert the realm-agnostic result back into observable effects on
    the observable world of JavaScript objects on the [event
    loop](#event-loop){#event-loop-for-spec-authors:event-loop-8}.

::: {#example-event-loop-using-algorithm .example}
The following is an algorithm that \"encrypts\" a passed-in
[list](https://infra.spec.whatwg.org/#list){#event-loop-for-spec-authors:list
x-internal="list"} of [scalar value
strings](https://infra.spec.whatwg.org/#scalar-value-string){#event-loop-for-spec-authors:scalar-value-string
x-internal="scalar-value-string"} `input`{.variable}, after parsing them
as URLs:

1.  Let `urls`{.variable} be an empty
    [list](https://infra.spec.whatwg.org/#list){#event-loop-for-spec-authors:list-2
    x-internal="list"}.

2.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#event-loop-for-spec-authors:list-iterate
    x-internal="list-iterate"} `string`{.variable} of
    `input`{.variable}:

    1.  Let `parsed`{.variable} be the result of [encoding-parsing a
        URL](urls-and-fetching.html#encoding-parsing-a-url){#event-loop-for-spec-authors:encoding-parsing-a-url}
        given `string`{.variable}, relative to the [current settings
        object](#current-settings-object){#event-loop-for-spec-authors:current-settings-object}.

    2.  If `parsed`{.variable} is failure, then return [a promise
        rejected
        with](https://webidl.spec.whatwg.org/#a-promise-rejected-with){#event-loop-for-spec-authors:a-promise-rejected-with
        x-internal="a-promise-rejected-with"} a
        [\"`SyntaxError`\"](https://webidl.spec.whatwg.org/#syntaxerror){#event-loop-for-spec-authors:syntaxerror
        x-internal="syntaxerror"}
        [`DOMException`{#event-loop-for-spec-authors:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    3.  Let `serialized`{.variable} be the result of applying the [URL
        serializer](https://url.spec.whatwg.org/#concept-url-serializer){#event-loop-for-spec-authors:concept-url-serializer
        x-internal="concept-url-serializer"} to `parsed`{.variable}.

    4.  [Append](https://infra.spec.whatwg.org/#list-append){#event-loop-for-spec-authors:list-append
        x-internal="list-append"} `serialized`{.variable} to
        `urls`{.variable}.

3.  Let `realm`{.variable} be the [current
    realm](https://tc39.es/ecma262/#current-realm){#event-loop-for-spec-authors:current-realm
    x-internal="current-realm"}.

4.  Let `p`{.variable} be a new promise.

5.  Run the following steps [in
    parallel](infrastructure.html#in-parallel){#event-loop-for-spec-authors:in-parallel-7}:

    1.  Let `encryptedURLs`{.variable} be an empty
        [list](https://infra.spec.whatwg.org/#list){#event-loop-for-spec-authors:list-3
        x-internal="list"}.

    2.  [For
        each](https://infra.spec.whatwg.org/#list-iterate){#event-loop-for-spec-authors:list-iterate-2
        x-internal="list-iterate"} `url`{.variable} of
        `urls`{.variable}:

        1.  Wait 100 milliseconds, so that people think we\'re doing
            heavy-duty encryption.

        2.  Let `encrypted`{.variable} be a new
            [string](https://infra.spec.whatwg.org/#string){#event-loop-for-spec-authors:string
            x-internal="string"} derived from `url`{.variable}, whose
            `n`{.variable}th [code
            unit](https://infra.spec.whatwg.org/#code-unit){#event-loop-for-spec-authors:code-unit
            x-internal="code-unit"} is equal to `url`{.variable}\'s
            `n`{.variable}th [code
            unit](https://infra.spec.whatwg.org/#code-unit){#event-loop-for-spec-authors:code-unit-2
            x-internal="code-unit"} plus 13.

        3.  [Append](https://infra.spec.whatwg.org/#list-append){#event-loop-for-spec-authors:list-append-2
            x-internal="list-append"} `encrypted`{.variable} to
            `encryptedURLs`{.variable}.

    3.  [Queue a global
        task](#queue-a-global-task){#event-loop-for-spec-authors:queue-a-global-task-6}
        on the [networking task
        source](#networking-task-source){#event-loop-for-spec-authors:networking-task-source},
        given `realm`{.variable}\'s [global
        object](#concept-realm-global){#event-loop-for-spec-authors:concept-realm-global},
        to perform the following steps:

        1.  Let `array`{.variable} be the result of
            [converting](https://webidl.spec.whatwg.org/#es-type-mapping){#event-loop-for-spec-authors:concept-idl-convert-2
            x-internal="concept-idl-convert"} `encryptedURLs`{.variable}
            to a JavaScript array, in `realm`{.variable}.

        2.  Resolve `p`{.variable} with `array`{.variable}.

6.  Return `p`{.variable}.

Here are several things to notice about this algorithm:

- It does its URL parsing up front, on the [event
  loop](#event-loop){#event-loop-for-spec-authors:event-loop-9}, before
  going to the [in
  parallel](infrastructure.html#in-parallel){#event-loop-for-spec-authors:in-parallel-8}
  steps. This is necessary, since parsing depends on the [current
  settings
  object](#current-settings-object){#event-loop-for-spec-authors:current-settings-object-2},
  which would no longer be current after going [in
  parallel](infrastructure.html#in-parallel){#event-loop-for-spec-authors:in-parallel-9}.

- Alternately, it could have saved a reference to the [current settings
  object](#current-settings-object){#event-loop-for-spec-authors:current-settings-object-3}\'s
  [API base
  URL](#api-base-url){#event-loop-for-spec-authors:api-base-url} and
  used it during the [in
  parallel](infrastructure.html#in-parallel){#event-loop-for-spec-authors:in-parallel-10}
  steps; that would have been equivalent. However, we recommend instead
  doing as much work as possible up front, as this example does.
  Attempting to save the correct values can be error prone; for example,
  if we\'d saved just the [current settings
  object](#current-settings-object){#event-loop-for-spec-authors:current-settings-object-4},
  instead of its [API base
  URL](#api-base-url){#event-loop-for-spec-authors:api-base-url-2},
  there would have been a potential race.

- It implicitly passes a
  [list](https://infra.spec.whatwg.org/#list){#event-loop-for-spec-authors:list-4
  x-internal="list"} of
  [strings](https://infra.spec.whatwg.org/#string){#event-loop-for-spec-authors:string-2
  x-internal="string"} from the initial steps to the [in
  parallel](infrastructure.html#in-parallel){#event-loop-for-spec-authors:in-parallel-11}
  steps. This is OK, as both
  [lists](https://infra.spec.whatwg.org/#list){#event-loop-for-spec-authors:list-5
  x-internal="list"} and
  [strings](https://infra.spec.whatwg.org/#string){#event-loop-for-spec-authors:string-3
  x-internal="string"} are
  [realm](https://tc39.es/ecma262/#sec-code-realms){#event-loop-for-spec-authors:realm-3
  x-internal="realm"}-agnostic.

- It performs \"expensive computation\" (waiting for 100 milliseconds
  per input URL) during the [in
  parallel](infrastructure.html#in-parallel){#event-loop-for-spec-authors:in-parallel-12}
  steps, thus not blocking the main [event
  loop](#event-loop){#event-loop-for-spec-authors:event-loop-10}.

- Promises, as observable JavaScript objects, are never created and
  manipulated during the [in
  parallel](infrastructure.html#in-parallel){#event-loop-for-spec-authors:in-parallel-13}
  steps. `p`{.variable} is created before entering those steps, and then
  is manipulated during a
  [task](#concept-task){#event-loop-for-spec-authors:concept-task-2}
  that is
  [queued](#queue-a-global-task){#event-loop-for-spec-authors:queue-a-global-task-7}
  specifically for that purpose.

- The creation of a JavaScript array object also happens during the
  queued task, and is careful to specify which realm it creates the
  array in since that is no longer obvious from context.

(On these last two points, see also [whatwg/webidl issue
#135](https://github.com/whatwg/webidl/issues/135) and [whatwg/webidl
issue #371](https://github.com/whatwg/webidl/issues/371), where we are
still mulling over the subtleties of the above promise-resolution
pattern.)

Another thing to note is that, in the event this algorithm was called
from a Web IDL-specified operation taking a
`sequence`\<[`USVString`{#event-loop-for-spec-authors:idl-usvstring}](https://webidl.spec.whatwg.org/#idl-USVString){x-internal="idl-usvstring"}\>,
there was an automatic conversion from
[realm](https://tc39.es/ecma262/#sec-code-realms){#event-loop-for-spec-authors:realm-4
x-internal="realm"}-specific JavaScript objects provided by the author
as input, into the realm-agnostic
`sequence`\<[`USVString`{#event-loop-for-spec-authors:idl-usvstring-2}](https://webidl.spec.whatwg.org/#idl-USVString){x-internal="idl-usvstring"}\>
Web IDL type, which we then treat as a
[list](https://infra.spec.whatwg.org/#list){#event-loop-for-spec-authors:list-6
x-internal="list"} of [scalar value
strings](https://infra.spec.whatwg.org/#scalar-value-string){#event-loop-for-spec-authors:scalar-value-string-2
x-internal="scalar-value-string"}. So depending on how your
specification is structured, there may be other implicit steps happening
on the main [event
loop](#event-loop){#event-loop-for-spec-authors:event-loop-11} that play
a part in this whole process of getting you ready to go [in
parallel](infrastructure.html#in-parallel){#event-loop-for-spec-authors:in-parallel-14}.
:::

#### [8.1.8]{.secno} Events[](#events){.self-link}

##### [8.1.8.1]{.secno} Event handlers[](#event-handler-attributes){.self-link} {#event-handler-attributes}

:::: {.mdn-anno .wrapped}
MDN

::: feature
[Events/Event_handlers](https://developer.mozilla.org/en-US/docs/Web/Guide/Events/Event_handlers "Events are signals fired inside the browser window that notify of changes in the browser or operating system environment. Programmers can create event handler code that will run when an event fires, allowing web pages to respond appropriately to change.")
:::
::::

Many objects can have [event handlers]{#event-handlers .dfn
lt="event handler" export=""} specified. These act as non-capture [event
listeners](https://dom.spec.whatwg.org/#concept-event-listener){#event-handler-attributes:event-listener
x-internal="event-listener"} for the object on which they are specified.
[\[DOM\]](references.html#refsDOM)

An [event
handler](#event-handlers){#event-handler-attributes:event-handlers} is a
[struct](https://infra.spec.whatwg.org/#struct){#event-handler-attributes:struct
x-internal="struct"} with two
[items](https://infra.spec.whatwg.org/#struct-item){#event-handler-attributes:struct-item
x-internal="struct-item"}:

- a [value]{#event-handler-value .dfn dfn-for="event handler"
  export=""}, which is either null, a callback object, or an [internal
  raw uncompiled
  handler](#internal-raw-uncompiled-handler){#event-handler-attributes:internal-raw-uncompiled-handler}.
  The
  [`EventHandler`{#event-handler-attributes:eventhandler}](#eventhandler)
  callback function type describes how this is exposed to scripts.
  Initially, an [event
  handler](#event-handlers){#event-handler-attributes:event-handlers-2}\'s
  [value](#event-handler-value){#event-handler-attributes:event-handler-value}
  must be set to null.

- a [listener]{#event-handler-listener .dfn dfn-for="event handler"
  export=""}, which is either null or an [event
  listener](https://dom.spec.whatwg.org/#concept-event-listener){#event-handler-attributes:event-listener-2
  x-internal="event-listener"} responsible for running [the event
  handler processing
  algorithm](#the-event-handler-processing-algorithm){#event-handler-attributes:the-event-handler-processing-algorithm}.
  Initially, an [event
  handler](#event-handlers){#event-handler-attributes:event-handlers-3}\'s
  [listener](#event-handler-listener){#event-handler-attributes:event-handler-listener}
  must be set to null.

Event handlers are exposed in two ways.

The first way, common to all event handlers, is as an [event handler IDL
attribute](#event-handler-idl-attributes){#event-handler-attributes:event-handler-idl-attributes}.

The second way is as an [event handler content
attribute](#event-handler-content-attributes){#event-handler-attributes:event-handler-content-attributes}.
Event handlers on [HTML
elements](infrastructure.html#html-elements){#event-handler-attributes:html-elements}
and some of the event handlers on
[`Window`{#event-handler-attributes:window}](nav-history-apis.html#window)
objects are exposed in this way.

For both of these two ways, the [event
handler](#event-handlers){#event-handler-attributes:event-handlers-4} is
exposed through a [name]{#event-handler-name .dfn}, which is a string
that always starts with \"`on`\" and is followed by the name of the
event for which the handler is intended.

------------------------------------------------------------------------

Most of the time, the object that exposes an [event
handler](#event-handlers){#event-handler-attributes:event-handlers-5} is
the same as the object on which the corresponding [event
listener](https://dom.spec.whatwg.org/#concept-event-listener){#event-handler-attributes:event-listener-3
x-internal="event-listener"} is added. However, the
[`body`{#event-handler-attributes:the-body-element}](sections.html#the-body-element)
and
[`frameset`{#event-handler-attributes:frameset}](obsolete.html#frameset)
elements expose several [event
handlers](#event-handlers){#event-handler-attributes:event-handlers-6}
that act upon the element\'s
[`Window`{#event-handler-attributes:window-2}](nav-history-apis.html#window)
object, if one exists. In either case, we call the object an [event
handler](#event-handlers){#event-handler-attributes:event-handlers-7}
acts upon the [target]{#event-handler-target .dfn} of that [event
handler](#event-handlers){#event-handler-attributes:event-handlers-8}.

To [determine the target of an event
handler]{#determining-the-target-of-an-event-handler .dfn}, given an
[`EventTarget`{#event-handler-attributes:eventtarget}](https://dom.spec.whatwg.org/#interface-eventtarget){x-internal="eventtarget"}
object `eventTarget`{.variable} on which the [event
handler](#event-handlers){#event-handler-attributes:event-handlers-9} is
exposed, and an [event handler
name](#event-handler-name){#event-handler-attributes:event-handler-name}
`name`{.variable}, the following steps are taken:

1.  If `eventTarget`{.variable} is not a
    [`body`{#event-handler-attributes:the-body-element-2}](sections.html#the-body-element)
    element or a
    [`frameset`{#event-handler-attributes:frameset-2}](obsolete.html#frameset)
    element, then return `eventTarget`{.variable}.

2.  If `name`{.variable} is not the name of an attribute member of the
    [`WindowEventHandlers`{#event-handler-attributes:windoweventhandlers}](#windoweventhandlers)
    interface mixin and the [`Window`-reflecting body element event
    handler
    set](#window-reflecting-body-element-event-handler-set){#event-handler-attributes:window-reflecting-body-element-event-handler-set}
    does not
    [contain](https://infra.spec.whatwg.org/#list-contain){#event-handler-attributes:list-contains
    x-internal="list-contains"} `name`{.variable}, then return
    `eventTarget`{.variable}.

3.  If `eventTarget`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#event-handler-attributes:node-document
    x-internal="node-document"} is not an [active
    document](document-sequences.html#active-document){#event-handler-attributes:active-document},
    then return null.

    This could happen if this object is a
    [`body`{#event-handler-attributes:the-body-element-3}](sections.html#the-body-element)
    element without a corresponding
    [`Window`{#event-handler-attributes:window-3}](nav-history-apis.html#window)
    object, for example.

    This check does not necessarily prevent
    [`body`{#event-handler-attributes:the-body-element-4}](sections.html#the-body-element)
    and
    [`frameset`{#event-handler-attributes:frameset-3}](obsolete.html#frameset)
    elements that are not [the body
    element](dom.html#the-body-element-2){#event-handler-attributes:the-body-element-2-2}
    of their [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#event-handler-attributes:node-document-2
    x-internal="node-document"} from reaching the next step. In
    particular, a
    [`body`{#event-handler-attributes:the-body-element-5}](sections.html#the-body-element)
    element created in an [active
    document](document-sequences.html#active-document){#event-handler-attributes:active-document-2}
    (perhaps with
    [`document.createElement()`{#event-handler-attributes:dom-document-createelement}](https://dom.spec.whatwg.org/#dom-document-createelement){x-internal="dom-document-createelement"})
    but not
    [connected](https://dom.spec.whatwg.org/#connected){#event-handler-attributes:connected
    x-internal="connected"} will also have its corresponding
    [`Window`{#event-handler-attributes:window-4}](nav-history-apis.html#window)
    object as the
    [target](#event-handler-target){#event-handler-attributes:event-handler-target}
    of several [event
    handlers](#event-handlers){#event-handler-attributes:event-handlers-10}
    exposed through it.

4.  Return `eventTarget`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#event-handler-attributes:node-document-3
    x-internal="node-document"}\'s [relevant global
    object](#concept-relevant-global){#event-handler-attributes:concept-relevant-global}.

------------------------------------------------------------------------

Each
[`EventTarget`{#event-handler-attributes:eventtarget-2}](https://dom.spec.whatwg.org/#interface-eventtarget){x-internal="eventtarget"}
object that has one or more [event
handlers](#event-handlers){#event-handler-attributes:event-handlers-11}
specified has an associated [event handler map]{#event-handler-map .dfn
dfn-for="EventTarget" export=""}, which is a
[map](https://infra.spec.whatwg.org/#ordered-map){#event-handler-attributes:ordered-map
x-internal="ordered-map"} of strings representing
[names](#event-handler-name){#event-handler-attributes:event-handler-name-2}
of [event
handlers](#event-handlers){#event-handler-attributes:event-handlers-12}
to [event
handlers](#event-handlers){#event-handler-attributes:event-handlers-13}.

When an
[`EventTarget`{#event-handler-attributes:eventtarget-3}](https://dom.spec.whatwg.org/#interface-eventtarget){x-internal="eventtarget"}
object that has one or more [event
handlers](#event-handlers){#event-handler-attributes:event-handlers-14}
specified is created, its [event handler
map](#event-handler-map){#event-handler-attributes:event-handler-map}
must be initialized such that it contains an
[entry](https://infra.spec.whatwg.org/#map-entry){#event-handler-attributes:map-entry
x-internal="map-entry"} for each [event
handler](#event-handlers){#event-handler-attributes:event-handlers-15}
that has that object as
[target](#event-handler-target){#event-handler-attributes:event-handler-target-2},
with
[items](https://infra.spec.whatwg.org/#struct-item){#event-handler-attributes:struct-item-2
x-internal="struct-item"} in those [event
handlers](#event-handlers){#event-handler-attributes:event-handlers-16}
set to their initial values.

The order of the
[entries](https://infra.spec.whatwg.org/#map-entry){#event-handler-attributes:map-entry-2
x-internal="map-entry"} of [event handler
map](#event-handler-map){#event-handler-attributes:event-handler-map-2}
could be arbitrary. It is not observable through any algorithms that
operate on the map.

[Entries](https://infra.spec.whatwg.org/#map-entry){#event-handler-attributes:map-entry-3
x-internal="map-entry"} are not created in the [event handler
map](#event-handler-map){#event-handler-attributes:event-handler-map-3}
of an object for [event
handlers](#event-handlers){#event-handler-attributes:event-handlers-17}
that are merely exposed on that object, but have some other object as
their
[targets](#event-handler-target){#event-handler-attributes:event-handler-target-3}.

------------------------------------------------------------------------

An [event handler IDL attribute]{#event-handler-idl-attributes .dfn
export=""} is an IDL attribute for a specific [event
handler](#event-handlers){#event-handler-attributes:event-handlers-18}.
The name of the IDL attribute is the same as the
[name](#event-handler-name){#event-handler-attributes:event-handler-name-3}
of the [event
handler](#event-handlers){#event-handler-attributes:event-handlers-19}.

The getter of an [event handler IDL
attribute](#event-handler-idl-attributes){#event-handler-attributes:event-handler-idl-attributes-2}
with name `name`{.variable}, when called, must run these steps:

1.  Let `eventTarget`{.variable} be the result of [determining the
    target of an event
    handler](#determining-the-target-of-an-event-handler){#event-handler-attributes:determining-the-target-of-an-event-handler}
    given this object and `name`{.variable}.

2.  If `eventTarget`{.variable} is null, then return null.

3.  Return the result of [getting the current value of the event
    handler](#getting-the-current-value-of-the-event-handler){#event-handler-attributes:getting-the-current-value-of-the-event-handler}
    given `eventTarget`{.variable} and `name`{.variable}.

The setter of an [event handler IDL
attribute](#event-handler-idl-attributes){#event-handler-attributes:event-handler-idl-attributes-3}
with name `name`{.variable}, when called, must run these steps:

1.  Let `eventTarget`{.variable} be the result of [determining the
    target of an event
    handler](#determining-the-target-of-an-event-handler){#event-handler-attributes:determining-the-target-of-an-event-handler-2}
    given this object and `name`{.variable}.

2.  If `eventTarget`{.variable} is null, then return.

3.  If the given value is null, then [deactivate an event
    handler](#deactivate-an-event-handler){#event-handler-attributes:deactivate-an-event-handler}
    given `eventTarget`{.variable} and `name`{.variable}.

4.  Otherwise:

    1.  Let `handlerMap`{.variable} be `eventTarget`{.variable}\'s
        [event handler
        map](#event-handler-map){#event-handler-attributes:event-handler-map-4}.

    2.  Let `eventHandler`{.variable} be
        `handlerMap`{.variable}\[`name`{.variable}\].

    3.  Set `eventHandler`{.variable}\'s
        [value](#event-handler-value){#event-handler-attributes:event-handler-value-2}
        to the given value.

    4.  [Activate an event
        handler](#activate-an-event-handler){#event-handler-attributes:activate-an-event-handler}
        given `eventTarget`{.variable} and `name`{.variable}.

Certain [event handler IDL
attributes](#event-handler-idl-attributes){#event-handler-attributes:event-handler-idl-attributes-4}
have additional requirements, in particular the
[`onmessage`{#event-handler-attributes:handler-messageeventtarget-onmessage}](web-messaging.html#handler-messageeventtarget-onmessage)
attribute of
[`MessagePort`{#event-handler-attributes:messageport}](web-messaging.html#messageport)
objects.

------------------------------------------------------------------------

An [event handler content attribute]{#event-handler-content-attributes
.dfn export=""} is a content attribute for a specific [event
handler](#event-handlers){#event-handler-attributes:event-handlers-20}.
The name of the content attribute is the same as the
[name](#event-handler-name){#event-handler-attributes:event-handler-name-4}
of the [event
handler](#event-handlers){#event-handler-attributes:event-handlers-21}.

[Event handler content
attributes](#event-handler-content-attributes){#event-handler-attributes:event-handler-content-attributes-2},
when specified, must contain valid JavaScript code which, when parsed,
would match the
*[FunctionBody](https://tc39.es/ecma262/#prod-FunctionBody){x-internal="js-prod-functionbody"}*
production after [automatic semicolon
insertion](https://tc39.es/ecma262/#sec-automatic-semicolon-insertion){#event-handler-attributes:automatic-semicolon-insertion
x-internal="automatic-semicolon-insertion"}.

The following [attribute change
steps](https://dom.spec.whatwg.org/#concept-element-attributes-change-ext){#event-handler-attributes:concept-element-attributes-change-ext
x-internal="concept-element-attributes-change-ext"} are used to
synchronize between [event handler content
attributes](#event-handler-content-attributes){#event-handler-attributes:event-handler-content-attributes-3}
and [event
handlers](#event-handlers){#event-handler-attributes:event-handlers-22}:
[\[DOM\]](references.html#refsDOM)

1.  If `namespace`{.variable} is not null, or `localName`{.variable} is
    not the name of an [event handler content
    attribute](#event-handler-content-attributes){#event-handler-attributes:event-handler-content-attributes-4}
    on `element`{.variable}, then return.

2.  Let `eventTarget`{.variable} be the result of [determining the
    target of an event
    handler](#determining-the-target-of-an-event-handler){#event-handler-attributes:determining-the-target-of-an-event-handler-3}
    given `element`{.variable} and `localName`{.variable}.

3.  If `eventTarget`{.variable} is null, then return.

4.  If `value`{.variable} is null, then [deactivate an event
    handler](#deactivate-an-event-handler){#event-handler-attributes:deactivate-an-event-handler-2}
    given `eventTarget`{.variable} and `localName`{.variable}.

5.  Otherwise:

    1.  If the [Should element\'s inline behavior be blocked by Content
        Security
        Policy?](https://w3c.github.io/webappsec-csp/#should-block-inline){#event-handler-attributes:should-element's-inline-behavior-be-blocked-by-content-security-policy
        x-internal="should-element's-inline-behavior-be-blocked-by-content-security-policy"}
        algorithm returns \"`Blocked`\" when executed upon
        `element`{.variable}, \"`script attribute`\", and
        `value`{.variable}, then return.
        [\[CSP\]](references.html#refsCSP)

    2.  Let `handlerMap`{.variable} be `eventTarget`{.variable}\'s
        [event handler
        map](#event-handler-map){#event-handler-attributes:event-handler-map-5}.

    3.  Let `eventHandler`{.variable} be
        `handlerMap`{.variable}\[`localName`{.variable}\].

    4.  Let `location`{.variable} be the script location that triggered
        the execution of these steps.

    5.  Set `eventHandler`{.variable}\'s
        [value](#event-handler-value){#event-handler-attributes:event-handler-value-3}
        to the [internal raw uncompiled
        handler](#internal-raw-uncompiled-handler){#event-handler-attributes:internal-raw-uncompiled-handler-2}
        `value`{.variable}/`location`{.variable}.

    6.  [Activate an event
        handler](#activate-an-event-handler){#event-handler-attributes:activate-an-event-handler-2}
        given `eventTarget`{.variable} and `localName`{.variable}.

Per the DOM Standard, these steps are run even if `oldValue`{.variable}
and `value`{.variable} are identical (setting an attribute to its
current value), but *not* if `oldValue`{.variable} and
`value`{.variable} are both null (removing an attribute that doesn\'t
currently exist). [\[DOM\]](references.html#refsDOM)

------------------------------------------------------------------------

To [deactivate an event handler]{#deactivate-an-event-handler .dfn}
given an
[`EventTarget`{#event-handler-attributes:eventtarget-4}](https://dom.spec.whatwg.org/#interface-eventtarget){x-internal="eventtarget"}
object `eventTarget`{.variable} and a string `name`{.variable} that is
the
[name](#event-handler-name){#event-handler-attributes:event-handler-name-5}
of an [event
handler](#event-handlers){#event-handler-attributes:event-handlers-23},
run these steps:

1.  Let `handlerMap`{.variable} be `eventTarget`{.variable}\'s [event
    handler
    map](#event-handler-map){#event-handler-attributes:event-handler-map-6}.

2.  Let `eventHandler`{.variable} be
    `handlerMap`{.variable}\[`name`{.variable}\].

3.  Set `eventHandler`{.variable}\'s
    [value](#event-handler-value){#event-handler-attributes:event-handler-value-4}
    to null.

4.  Let `listener`{.variable} be `eventHandler`{.variable}\'s
    [listener](#event-handler-listener){#event-handler-attributes:event-handler-listener-2}.

5.  If `listener`{.variable} is not null, then [remove an event
    listener](https://dom.spec.whatwg.org/#remove-an-event-listener){#event-handler-attributes:remove-an-event-listener
    x-internal="remove-an-event-listener"} with `eventTarget`{.variable}
    and `listener`{.variable}.

6.  Set `eventHandler`{.variable}\'s
    [listener](#event-handler-listener){#event-handler-attributes:event-handler-listener-3}
    to null.

To [erase all event listeners and
handlers]{#erase-all-event-listeners-and-handlers .dfn} given an
[`EventTarget`{#event-handler-attributes:eventtarget-5}](https://dom.spec.whatwg.org/#interface-eventtarget){x-internal="eventtarget"}
object `eventTarget`{.variable}, run these steps:

1.  If `eventTarget`{.variable} has an associated [event handler
    map](#event-handler-map){#event-handler-attributes:event-handler-map-7},
    then for each `name`{.variable} → `eventHandler`{.variable} of
    `eventTarget`{.variable}\'s associated [event handler
    map](#event-handler-map){#event-handler-attributes:event-handler-map-8},
    [deactivate an event
    handler](#deactivate-an-event-handler){#event-handler-attributes:deactivate-an-event-handler-3}
    given `eventTarget`{.variable} and `name`{.variable}.

2.  [Remove all event
    listeners](https://dom.spec.whatwg.org/#remove-all-event-listeners){#event-handler-attributes:remove-all-event-listeners
    x-internal="remove-all-event-listeners"} given
    `eventTarget`{.variable}.

This algorithm is used to define
[`document.open()`{#event-handler-attributes:dom-document-open}](dynamic-markup-insertion.html#dom-document-open).

To [activate an event handler]{#activate-an-event-handler .dfn} given an
[`EventTarget`{#event-handler-attributes:eventtarget-6}](https://dom.spec.whatwg.org/#interface-eventtarget){x-internal="eventtarget"}
object `eventTarget`{.variable} and a string `name`{.variable} that is
the
[name](#event-handler-name){#event-handler-attributes:event-handler-name-6}
of an [event
handler](#event-handlers){#event-handler-attributes:event-handlers-24},
run these steps:

1.  Let `handlerMap`{.variable} be `eventTarget`{.variable}\'s [event
    handler
    map](#event-handler-map){#event-handler-attributes:event-handler-map-9}.

2.  Let `eventHandler`{.variable} be
    `handlerMap`{.variable}\[`name`{.variable}\].

3.  If `eventHandler`{.variable}\'s
    [listener](#event-handler-listener){#event-handler-attributes:event-handler-listener-4}
    is not null, then return.

4.  Let `callback`{.variable} be the result of creating a Web IDL
    [`EventListener`{#event-handler-attributes:dom-eventlistener}](https://dom.spec.whatwg.org/#callbackdef-eventlistener){x-internal="dom-eventlistener"}
    instance representing a reference to a function of one argument that
    executes the steps of [the event handler processing
    algorithm](#the-event-handler-processing-algorithm){#event-handler-attributes:the-event-handler-processing-algorithm-2},
    given `eventTarget`{.variable}, `name`{.variable}, and its argument.

    The
    [`EventListener`{#event-handler-attributes:dom-eventlistener-2}](https://dom.spec.whatwg.org/#callbackdef-eventlistener){x-internal="dom-eventlistener"}\'s
    [callback
    context](https://webidl.spec.whatwg.org/#dfn-callback-context){#event-handler-attributes:callback-context
    x-internal="callback-context"} can be arbitrary; it does not impact
    the steps of [the event handler processing
    algorithm](#the-event-handler-processing-algorithm){#event-handler-attributes:the-event-handler-processing-algorithm-3}.
    [\[DOM\]](references.html#refsDOM)

    The callback is emphatically *not* the [event
    handler](#event-handlers){#event-handler-attributes:event-handlers-25}
    itself. Every event handler ends up registering the same callback,
    the algorithm defined below, which takes care of invoking the right
    code, and processing the code\'s return value.

5.  Let `listener`{.variable} be a new [event
    listener](https://dom.spec.whatwg.org/#concept-event-listener){#event-handler-attributes:event-listener-4
    x-internal="event-listener"} whose
    [type](https://dom.spec.whatwg.org/#event-listener-type){#event-handler-attributes:event-listener-type
    x-internal="event-listener-type"} is the [event handler event
    type]{#event-handler-event-type .dfn export=""} corresponding to
    `eventHandler`{.variable} and
    [callback](https://dom.spec.whatwg.org/#event-listener-callback){#event-handler-attributes:event-listener-callback
    x-internal="event-listener-callback"} is `callback`{.variable}.

    To be clear, an [event
    listener](https://dom.spec.whatwg.org/#concept-event-listener){#event-handler-attributes:event-listener-5
    x-internal="event-listener"} is different from an
    [`EventListener`{#event-handler-attributes:dom-eventlistener-3}](https://dom.spec.whatwg.org/#callbackdef-eventlistener){x-internal="dom-eventlistener"}.

6.  [Add an event
    listener](https://dom.spec.whatwg.org/#add-an-event-listener){#event-handler-attributes:add-an-event-listener
    x-internal="add-an-event-listener"} with `eventTarget`{.variable}
    and `listener`{.variable}.

7.  Set `eventHandler`{.variable}\'s
    [listener](#event-handler-listener){#event-handler-attributes:event-handler-listener-5}
    to `listener`{.variable}.

::: note
The event listener registration happens only if the [event
handler](#event-handlers){#event-handler-attributes:event-handlers-26}\'s
[value](#event-handler-value){#event-handler-attributes:event-handler-value-5}
is being set to non-null, and the [event
handler](#event-handlers){#event-handler-attributes:event-handlers-27}
is not already activated. Since listeners are called in the order they
were registered, assuming no
[deactivation](#deactivate-an-event-handler){#event-handler-attributes:deactivate-an-event-handler-4}
occurred, the order of event listeners for a particular event type will
always be:

1.  the event listeners registered with
    [`addEventListener()`{#event-handler-attributes:dom-eventtarget-addeventlistener}](https://dom.spec.whatwg.org/#dom-eventtarget-addeventlistener){x-internal="dom-eventtarget-addeventlistener"}
    before the first time the [event
    handler](#event-handlers){#event-handler-attributes:event-handlers-28}\'s
    [value](#event-handler-value){#event-handler-attributes:event-handler-value-6}
    was set to non-null

2.  then the callback to which it is currently set, if any

3.  and finally the event listeners registered with
    [`addEventListener()`{#event-handler-attributes:dom-eventtarget-addeventlistener-2}](https://dom.spec.whatwg.org/#dom-eventtarget-addeventlistener){x-internal="dom-eventtarget-addeventlistener"}
    *after* the first time the [event
    handler](#event-handlers){#event-handler-attributes:event-handlers-29}\'s
    [value](#event-handler-value){#event-handler-attributes:event-handler-value-7}
    was set to non-null.
:::

::: example
This example demonstrates the order in which event listeners are
invoked. If the button in this example is clicked by the user, the page
will show four alerts, with the text \"ONE\", \"TWO\", \"THREE\", and
\"FOUR\" respectively.

``` html
<button id="test">Start Demo</button>
<script>
 var button = document.getElementById('test');
 button.addEventListener('click', function () { alert('ONE') }, false);
 button.setAttribute('onclick', "alert('NOT CALLED')"); // event handler listener is registered here
 button.addEventListener('click', function () { alert('THREE') }, false);
 button.onclick = function () { alert('TWO'); };
 button.addEventListener('click', function () { alert('FOUR') }, false);
</script>
```

However, in the following example, the event handler is
[deactivated](#deactivate-an-event-handler){#event-handler-attributes:deactivate-an-event-handler-5}
after its initial activation (and its event listener is removed), before
being reactivated at a later time. The page will show five alerts with
\"ONE\", \"TWO\", \"THREE\", \"FOUR\", and \"FIVE\" respectively, in
order.

``` html
<button id="test">Start Demo</button>
<script>
 var button = document.getElementById('test');
 button.addEventListener('click', function () { alert('ONE') }, false);
 button.setAttribute('onclick', "alert('NOT CALLED')"); // event handler is activated here
 button.addEventListener('click', function () { alert('TWO') }, false);
 button.onclick = null;                                 // but deactivated here
 button.addEventListener('click', function () { alert('THREE') }, false);
 button.onclick = function () { alert('FOUR'); };       // and re-activated here
 button.addEventListener('click', function () { alert('FIVE') }, false);
</script>
```
:::

The interfaces implemented by the event object do not influence whether
an [event
handler](#event-handlers){#event-handler-attributes:event-handlers-30}
is triggered or not.

[The event handler processing
algorithm]{#the-event-handler-processing-algorithm .dfn} for an
[`EventTarget`{#event-handler-attributes:eventtarget-7}](https://dom.spec.whatwg.org/#interface-eventtarget){x-internal="eventtarget"}
object `eventTarget`{.variable}, a string `name`{.variable} representing
the
[name](#event-handler-name){#event-handler-attributes:event-handler-name-7}
of an [event
handler](#event-handlers){#event-handler-attributes:event-handlers-31},
and an
[`Event`{#event-handler-attributes:event}](https://dom.spec.whatwg.org/#interface-event){x-internal="event"}
object `event`{.variable} is as follows:

1.  Let `callback`{.variable} be the result of [getting the current
    value of the event
    handler](#getting-the-current-value-of-the-event-handler){#event-handler-attributes:getting-the-current-value-of-the-event-handler-2}
    given `eventTarget`{.variable} and `name`{.variable}.

2.  If `callback`{.variable} is null, then return.

3.  Let `special error event handling`{.variable} be true if
    `event`{.variable} is an
    [`ErrorEvent`{#event-handler-attributes:errorevent}](#errorevent)
    object, `event`{.variable}\'s
    [`type`{#event-handler-attributes:dom-event-type}](https://dom.spec.whatwg.org/#dom-event-type){x-internal="dom-event-type"}
    is
    \"[`error`{#event-handler-attributes:event-error}](indices.html#event-error)\",
    and `event`{.variable}\'s
    [`currentTarget`{#event-handler-attributes:dom-event-currenttarget}](https://dom.spec.whatwg.org/#dom-event-currenttarget){x-internal="dom-event-currenttarget"}
    implements the
    [`WindowOrWorkerGlobalScope`{#event-handler-attributes:windoworworkerglobalscope}](#windoworworkerglobalscope)
    mixin. Otherwise, let `special error event handling`{.variable} be
    false.

4.  Process the
    [`Event`{#event-handler-attributes:event-2}](https://dom.spec.whatwg.org/#interface-event){x-internal="event"}
    object `event`{.variable} as follows:

    If `special error event handling`{.variable} is true

    :   Let `return value`{.variable} be the result of
        [invoking](https://webidl.spec.whatwg.org/#invoke-a-callback-function){#event-handler-attributes:es-invoking-callback-functions
        x-internal="es-invoking-callback-functions"}
        `callback`{.variable} with « `event`{.variable}\'s
        [`message`{#event-handler-attributes:dom-errorevent-message}](#dom-errorevent-message),
        `event`{.variable}\'s
        [`filename`{#event-handler-attributes:dom-errorevent-filename}](#dom-errorevent-filename),
        `event`{.variable}\'s
        [`lineno`{#event-handler-attributes:dom-errorevent-lineno}](#dom-errorevent-lineno),
        `event`{.variable}\'s
        [`colno`{#event-handler-attributes:dom-errorevent-colno}](#dom-errorevent-colno),
        `event`{.variable}\'s
        [`error`{#event-handler-attributes:dom-errorevent-error}](#dom-errorevent-error)
        », \"`rethrow`\", and with *[callback this
        value](https://webidl.spec.whatwg.org/#dfn-callback-this-value){x-internal="dfn-callback-this-value"}*
        set to `event`{.variable}\'s
        [`currentTarget`{#event-handler-attributes:dom-event-currenttarget-2}](https://dom.spec.whatwg.org/#dom-event-currenttarget){x-internal="dom-event-currenttarget"}.

    Otherwise

    :   Let `return value`{.variable} be the result of
        [invoking](https://webidl.spec.whatwg.org/#invoke-a-callback-function){#event-handler-attributes:es-invoking-callback-functions-2
        x-internal="es-invoking-callback-functions"}
        `callback`{.variable} with « `event`{.variable} »,
        \"`rethrow`\", and with *[callback this
        value](https://webidl.spec.whatwg.org/#dfn-callback-this-value){x-internal="dfn-callback-this-value"}*
        set to `event`{.variable}\'s
        [`currentTarget`{#event-handler-attributes:dom-event-currenttarget-3}](https://dom.spec.whatwg.org/#dom-event-currenttarget){x-internal="dom-event-currenttarget"}.

    If an exception gets thrown by the callback, it will be rethrown,
    ending these steps. The exception will propagate to the [DOM event
    dispatch
    logic](https://dom.spec.whatwg.org/#concept-event-dispatch){#event-handler-attributes:concept-event-dispatch
    x-internal="concept-event-dispatch"}, which will then
    [report](#report-an-exception){#event-handler-attributes:report-an-exception}
    it.

5.  Process `return value`{.variable} as follows:

    If `event`{.variable} is a [`BeforeUnloadEvent`{#event-handler-attributes:beforeunloadevent}](nav-history-apis.html#beforeunloadevent) object and `event`{.variable}\'s [`type`{#event-handler-attributes:dom-event-type-2}](https://dom.spec.whatwg.org/#dom-event-type){x-internal="dom-event-type"} is \"[`beforeunload`{#event-handler-attributes:event-beforeunload}](indices.html#event-beforeunload)\"

    :   In this case, the [event handler IDL
        attribute](#event-handler-idl-attributes){#event-handler-attributes:event-handler-idl-attributes-5}\'s
        type will be
        [`OnBeforeUnloadEventHandler`{#event-handler-attributes:onbeforeunloadeventhandler}](#onbeforeunloadeventhandler),
        so `return value`{.variable} will have been coerced into either
        null or a
        [`DOMString`{#event-handler-attributes:idl-domstring}](https://webidl.spec.whatwg.org/#idl-DOMString){x-internal="idl-domstring"}.

        If `return value`{.variable} is not null, then:

        1.  Set `event`{.variable}\'s [canceled
            flag](https://dom.spec.whatwg.org/#canceled-flag){#event-handler-attributes:canceled-flag
            x-internal="canceled-flag"}.

        2.  If `event`{.variable}\'s
            [`returnValue`{#event-handler-attributes:dom-beforeunloadevent-returnvalue}](nav-history-apis.html#dom-beforeunloadevent-returnvalue)
            attribute\'s value is the empty string, then set
            `event`{.variable}\'s
            [`returnValue`{#event-handler-attributes:dom-beforeunloadevent-returnvalue-2}](nav-history-apis.html#dom-beforeunloadevent-returnvalue)
            attribute\'s value to `return value`{.variable}.

    If `special error event handling`{.variable} is true
    :   If `return value`{.variable} is true, then set
        `event`{.variable}\'s [canceled
        flag](https://dom.spec.whatwg.org/#canceled-flag){#event-handler-attributes:canceled-flag-2
        x-internal="canceled-flag"}.

    Otherwise

    :   If `return value`{.variable} is false, then set
        `event`{.variable}\'s [canceled
        flag](https://dom.spec.whatwg.org/#canceled-flag){#event-handler-attributes:canceled-flag-3
        x-internal="canceled-flag"}.

        If we\'ve gotten to this \"Otherwise\" clause because
        `event`{.variable}\'s
        [`type`{#event-handler-attributes:dom-event-type-3}](https://dom.spec.whatwg.org/#dom-event-type){x-internal="dom-event-type"}
        is
        \"[`beforeunload`{#event-handler-attributes:event-beforeunload-2}](indices.html#event-beforeunload)\"
        but `event`{.variable} is *not* a
        [`BeforeUnloadEvent`{#event-handler-attributes:beforeunloadevent-2}](nav-history-apis.html#beforeunloadevent)
        object, then `return value`{.variable} will never be false,
        since in such cases `return value`{.variable} will have been
        coerced into either null or a
        [`DOMString`{#event-handler-attributes:idl-domstring-2}](https://webidl.spec.whatwg.org/#idl-DOMString){x-internal="idl-domstring"}.

------------------------------------------------------------------------

The
[`EventHandler`{#event-handler-attributes:eventhandler-2}](#eventhandler)
callback function type represents a callback used for event handlers. It
is represented in Web IDL as follows:

``` idl
[LegacyTreatNonObjectAsNull]
callback EventHandlerNonNull = any (Event event);
typedef EventHandlerNonNull? EventHandler;
```

In JavaScript, any
[`Function`{#event-handler-attributes:idl-function}](https://webidl.spec.whatwg.org/#common-Function){x-internal="idl-function"}
object implements this interface.

::: example
For example, the following document fragment:

``` html
<body onload="alert(this)" onclick="alert(this)">
```

\...leads to an alert saying \"`[object Window]`\" when the document is
loaded, and an alert saying \"`[object HTMLBodyElement]`\" whenever the
user clicks something in the page.
:::

::: note
The return value of the function affects whether the event is canceled
or not: as described above, if the return value is false, the event is
canceled.

There are two exceptions in the platform, for historical reasons:

- The
  [`onerror`{#event-handler-attributes:handler-onerror}](#handler-onerror)
  handlers on global objects, where returning *true* cancels the event.

- The
  [`onbeforeunload`{#event-handler-attributes:handler-window-onbeforeunload}](#handler-window-onbeforeunload)
  handler, where returning any non-null and non-undefined value will
  cancel the event.
:::

For historical reasons, the
[`onerror`{#event-handler-attributes:handler-onerror-2}](#handler-onerror)
handler has different arguments:

``` idl
[LegacyTreatNonObjectAsNull]
callback OnErrorEventHandlerNonNull = any ((Event or DOMString) event, optional DOMString source, optional unsigned long lineno, optional unsigned long colno, optional any error);
typedef OnErrorEventHandlerNonNull? OnErrorEventHandler;
```

::: example
``` js
window.onerror = (message, source, lineno, colno, error) => { … };
```
:::

Similarly, the
[`onbeforeunload`{#event-handler-attributes:handler-window-onbeforeunload-2}](#handler-window-onbeforeunload)
handler has a different return value:

``` idl
[LegacyTreatNonObjectAsNull]
callback OnBeforeUnloadEventHandlerNonNull = DOMString? (Event event);
typedef OnBeforeUnloadEventHandlerNonNull? OnBeforeUnloadEventHandler;
```

------------------------------------------------------------------------

An [internal raw uncompiled handler]{#internal-raw-uncompiled-handler
.dfn} is a tuple with the following information:

- An uncompiled script body

- A location where the script body originated, in case an error needs to
  be reported

When the user agent is to [get the current value of the event
handler]{#getting-the-current-value-of-the-event-handler .dfn} given an
[`EventTarget`{#event-handler-attributes:eventtarget-8}](https://dom.spec.whatwg.org/#interface-eventtarget){x-internal="eventtarget"}
object `eventTarget`{.variable} and a string `name`{.variable} that is
the
[name](#event-handler-name){#event-handler-attributes:event-handler-name-8}
of an [event
handler](#event-handlers){#event-handler-attributes:event-handlers-32},
it must run these steps:

1.  Let `handlerMap`{.variable} be `eventTarget`{.variable}\'s [event
    handler
    map](#event-handler-map){#event-handler-attributes:event-handler-map-10}.

2.  Let `eventHandler`{.variable} be
    `handlerMap`{.variable}\[`name`{.variable}\].

3.  If `eventHandler`{.variable}\'s
    [value](#event-handler-value){#event-handler-attributes:event-handler-value-8}
    is an [internal raw uncompiled
    handler](#internal-raw-uncompiled-handler){#event-handler-attributes:internal-raw-uncompiled-handler-3},
    then:

    1.  If `eventTarget`{.variable} is an element, then let
        `element`{.variable} be `eventTarget`{.variable}, and
        `document`{.variable} be `element`{.variable}\'s [node
        document](https://dom.spec.whatwg.org/#concept-node-document){#event-handler-attributes:node-document-4
        x-internal="node-document"}. Otherwise, `eventTarget`{.variable}
        is a
        [`Window`{#event-handler-attributes:window-5}](nav-history-apis.html#window)
        object, let `element`{.variable} be null, and
        `document`{.variable} be `eventTarget`{.variable}\'s [associated
        `Document`](nav-history-apis.html#concept-document-window){#event-handler-attributes:concept-document-window}.

    2.  If [scripting is
        disabled](#concept-n-noscript){#event-handler-attributes:concept-n-noscript}
        for `document`{.variable}, then return null.

    3.  Let `body`{.variable} be the uncompiled script body in
        `eventHandler`{.variable}\'s
        [value](#event-handler-value){#event-handler-attributes:event-handler-value-9}.

    4.  Let `location`{.variable} be the location where the script body
        originated, as given by `eventHandler`{.variable}\'s
        [value](#event-handler-value){#event-handler-attributes:event-handler-value-10}.

    5.  If `element`{.variable} is not null and `element`{.variable} has
        a [form
        owner](form-control-infrastructure.html#form-owner){#event-handler-attributes:form-owner},
        let `form owner`{.variable} be that [form
        owner](form-control-infrastructure.html#form-owner){#event-handler-attributes:form-owner-2}.
        Otherwise, let `form owner`{.variable} be null.

    6.  Let `settings object`{.variable} be the [relevant settings
        object](#relevant-settings-object){#event-handler-attributes:relevant-settings-object}
        of `document`{.variable}.

    7.  If `body`{.variable} is not parsable as
        *[FunctionBody](https://tc39.es/ecma262/#prod-FunctionBody){x-internal="js-prod-functionbody"}*
        or if parsing detects an [early
        error](https://tc39.es/ecma262/#early-error-rule){#event-handler-attributes:early-error
        x-internal="early-error"}, then follow these substeps:

        1.  Set `eventHandler`{.variable}\'s
            [value](#event-handler-value){#event-handler-attributes:event-handler-value-11}
            to null.

            This does not
            [deactivate](#deactivate-an-event-handler){#event-handler-attributes:deactivate-an-event-handler-6}
            the event handler, which additionally
            [removes](https://dom.spec.whatwg.org/#remove-an-event-listener){#event-handler-attributes:remove-an-event-listener-2
            x-internal="remove-an-event-listener"} the event handler\'s
            [listener](#event-handler-listener){#event-handler-attributes:event-handler-listener-6}
            (if present).

        2.  Let `syntaxError`{.variable} be a new
            [`SyntaxError`{#event-handler-attributes:syntaxerror-2}](https://tc39.es/ecma262/#sec-native-error-types-used-in-this-standard-syntaxerror){x-internal="syntaxerror-2"}
            exception associated with `settings object`{.variable}\'s
            [realm](#environment-settings-object's-realm){#event-handler-attributes:environment-settings-object's-realm}
            which describes the error while parsing. It should be based
            on `location`{.variable}, where the script body originated.

        3.  [Report an
            exception](#report-an-exception){#event-handler-attributes:report-an-exception-2}
            with `syntaxError`{.variable} for
            `settings object`{.variable}\'s [global
            object](#concept-settings-object-global){#event-handler-attributes:concept-settings-object-global}.

        4.  Return null.

    8.  Push `settings object`{.variable}\'s [realm execution
        context](#realm-execution-context){#event-handler-attributes:realm-execution-context}
        onto the [JavaScript execution context
        stack](https://tc39.es/ecma262/#execution-context-stack){#event-handler-attributes:javascript-execution-context-stack
        x-internal="javascript-execution-context-stack"}; it is now the
        [running JavaScript execution
        context](https://tc39.es/ecma262/#running-execution-context){#event-handler-attributes:running-javascript-execution-context
        x-internal="running-javascript-execution-context"}.

        This is necessary so the subsequent invocation of
        [OrdinaryFunctionCreate](https://tc39.es/ecma262/#sec-ordinaryfunctioncreate){#event-handler-attributes:js-ordinaryfunctioncreate
        x-internal="js-ordinaryfunctioncreate"} takes place in the
        correct
        [realm](https://tc39.es/ecma262/#sec-code-realms){#event-handler-attributes:realm
        x-internal="realm"}.

    9.  Let `function`{.variable} be the result of calling
        [OrdinaryFunctionCreate](https://tc39.es/ecma262/#sec-ordinaryfunctioncreate){#event-handler-attributes:js-ordinaryfunctioncreate-2
        x-internal="js-ordinaryfunctioncreate"}, with arguments:

        `functionPrototype`{.variable}
        :   [%Function.prototype%](https://tc39.es/ecma262/#sec-properties-of-the-function-prototype-object){#event-handler-attributes:function.prototype
            x-internal="function.prototype"}

        `sourceText`{.variable}

        :   

            If `name`{.variable} is [`onerror`{#event-handler-attributes:handler-onerror-3}](#handler-onerror) and `eventTarget`{.variable} is a [`Window`{#event-handler-attributes:window-6}](nav-history-apis.html#window) object
            :   The string formed by concatenating \"`function `\",
                `name`{.variable},
                \"`(event, source, lineno, colno, error) {`\", U+000A
                LF, `body`{.variable}, U+000A LF, and \"`}`\".

            Otherwise
            :   The string formed by concatenating \"`function `\",
                `name`{.variable}, \"`(event) {`\", U+000A LF,
                `body`{.variable}, U+000A LF, and \"`}`\".

        `ParameterList`{.variable}

        :   

            If `name`{.variable} is [`onerror`{#event-handler-attributes:handler-onerror-4}](#handler-onerror) and `eventTarget`{.variable} is a [`Window`{#event-handler-attributes:window-7}](nav-history-apis.html#window) object
            :   Let the function have five arguments, named `event`,
                `source`, `lineno`, `colno`, and `error`.

            Otherwise
            :   Let the function have a single argument called `event`.

        `body`{.variable}
        :   The result of parsing `body`{.variable} above.

        `thisMode`{.variable}
        :   non-lexical-this

        `scope`{.variable}

        :   1.  Let `realm`{.variable} be
                `settings object`{.variable}\'s
                [realm](#environment-settings-object's-realm){#event-handler-attributes:environment-settings-object's-realm-2}.

            2.  Let `scope`{.variable} be
                `realm`{.variable}.\[\[GlobalEnv\]\].

            3.  If `eventHandler`{.variable} is an element\'s [event
                handler](#event-handlers){#event-handler-attributes:event-handlers-33},
                then set `scope`{.variable} to
                [NewObjectEnvironment](https://tc39.es/ecma262/#sec-newobjectenvironment){#event-handler-attributes:js-newobjectenvironment
                x-internal="js-newobjectenvironment"}(`document`{.variable},
                true, `scope`{.variable}).

                (Otherwise, `eventHandler`{.variable} is a
                [`Window`{#event-handler-attributes:window-8}](nav-history-apis.html#window)
                object\'s [event
                handler](#event-handlers){#event-handler-attributes:event-handlers-34}.)

            4.  If `form owner`{.variable} is not null, then set
                `scope`{.variable} to
                [NewObjectEnvironment](https://tc39.es/ecma262/#sec-newobjectenvironment){#event-handler-attributes:js-newobjectenvironment-2
                x-internal="js-newobjectenvironment"}(`form owner`{.variable},
                true, `scope`{.variable}).

            5.  If `element`{.variable} is not null, then set
                `scope`{.variable} to
                [NewObjectEnvironment](https://tc39.es/ecma262/#sec-newobjectenvironment){#event-handler-attributes:js-newobjectenvironment-3
                x-internal="js-newobjectenvironment"}(`element`{.variable},
                true, `scope`{.variable}).

            6.  Return `scope`{.variable}.

    10. Remove `settings object`{.variable}\'s [realm execution
        context](#realm-execution-context){#event-handler-attributes:realm-execution-context-2}
        from the [JavaScript execution context
        stack](https://tc39.es/ecma262/#execution-context-stack){#event-handler-attributes:javascript-execution-context-stack-2
        x-internal="javascript-execution-context-stack"}.

    11. Set `function`{.variable}.\[\[ScriptOrModule\]\] to null.

        ::: note
        This is done because the default behavior, of associating the
        created function with the nearest
        [script](#concept-script){#event-handler-attributes:concept-script}
        on the stack, can lead to path-dependent results. For example,
        an event handler which is first invoked by user interaction
        would end up with null \[\[ScriptOrModule\]\] (since then this
        algorithm would be first invoked when the [active
        script](#active-script){#event-handler-attributes:active-script}
        is null), whereas one that is first invoked by dispatching an
        event from script would have its \[\[ScriptOrModule\]\] set to
        that script.

        Instead, we just always set \[\[ScriptOrModule\]\] to null. This
        is more intuitive anyway; the idea that the first script which
        dispatches an event is somehow responsible for the event handler
        code is dubious.

        In practice, this only affects the resolution of relative URLs
        via
        [`import()`{#event-handler-attributes:import()}](https://tc39.es/ecma262/#sec-import-calls){x-internal="import()"},
        which consult the [base
        URL](#concept-script-base-url){#event-handler-attributes:concept-script-base-url}
        of the associated script. Nulling out \[\[ScriptOrModule\]\]
        means that
        [HostLoadImportedModule](#hostloadimportedmodule){#event-handler-attributes:hostloadimportedmodule}
        will fall back to the [current settings
        object](#current-settings-object){#event-handler-attributes:current-settings-object}\'s
        [API base
        URL](#api-base-url){#event-handler-attributes:api-base-url}.
        :::

    12. Set `eventHandler`{.variable}\'s
        [value](#event-handler-value){#event-handler-attributes:event-handler-value-12}
        to the result of creating a Web IDL
        [`EventHandler`{#event-handler-attributes:eventhandler-3}](#eventhandler)
        callback function object whose object reference is
        `function`{.variable} and whose [callback
        context](https://webidl.spec.whatwg.org/#dfn-callback-context){#event-handler-attributes:callback-context-2
        x-internal="callback-context"} is `settings object`{.variable}.

4.  Return `eventHandler`{.variable}\'s
    [value](#event-handler-value){#event-handler-attributes:event-handler-value-13}.

##### [8.1.8.2]{.secno} Event handlers on elements, [`Document`{#event-handlers-on-elements,-document-objects,-and-window-objects:document}](dom.html#document) objects, and [`Window`{#event-handlers-on-elements,-document-objects,-and-window-objects:window}](nav-history-apis.html#window) objects[](#event-handlers-on-elements,-document-objects,-and-window-objects){.self-link} {#event-handlers-on-elements,-document-objects,-and-window-objects}

The following are the [event
handlers](#event-handlers){#event-handlers-on-elements,-document-objects,-and-window-objects:event-handlers}
(and their corresponding [event handler event
types](#event-handler-event-type){#event-handlers-on-elements,-document-objects,-and-window-objects:event-handler-event-type})
that must be supported by all [HTML
elements](infrastructure.html#html-elements){#event-handlers-on-elements,-document-objects,-and-window-objects:html-elements},
as both [event handler content
attributes](#event-handler-content-attributes){#event-handlers-on-elements,-document-objects,-and-window-objects:event-handler-content-attributes}
and [event handler IDL
attributes](#event-handler-idl-attributes){#event-handlers-on-elements,-document-objects,-and-window-objects:event-handler-idl-attributes};
and that must be supported by all
[`Document`{#event-handlers-on-elements,-document-objects,-and-window-objects:document-2}](dom.html#document)
and
[`Window`{#event-handlers-on-elements,-document-objects,-and-window-objects:window-2}](nav-history-apis.html#window)
objects, as [event handler IDL
attributes](#event-handler-idl-attributes){#event-handlers-on-elements,-document-objects,-and-window-objects:event-handler-idl-attributes-2}:

[Event
handler](#event-handlers){#event-handlers-on-elements,-document-objects,-and-window-objects:event-handlers-2}

[Event handler event
type](#event-handler-event-type){#event-handlers-on-elements,-document-objects,-and-window-objects:event-handler-event-type-2}

[`onabort`]{#handler-onabort .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/abort_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/abort_event "The abort event is fired when the resource was not fully loaded, but not as the result of an error.")

Support in all current engines.

::: support
[Firefox9+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

`abort`

[`onauxclick`]{#handler-onauxclick .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
MDN

:::: feature
[Element/auxclick_event](https://developer.mozilla.org/en-US/docs/Web/API/Element/auxclick_event "The auxclick event is fired at an Element when a non-primary pointing device button (any mouse button other than the primary—usually leftmost—button) has been pressed and released both within the same element.")

::: support
[Firefox53+]{.firefox .yes}[SafariNo]{.safari .no}[Chrome55+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android53+]{.firefox_android .yes}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

[`auxclick`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-auxclick}](https://w3c.github.io/uievents/#event-type-auxclick){x-internal="event-auxclick"}

[`onbeforeinput`]{#handler-onbeforeinput .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

[`beforeinput`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-beforeinput}](https://w3c.github.io/uievents/#event-type-beforeinput){x-internal="event-beforeinput"}

[`onbeforematch`]{#handler-onbeforematch .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

[`beforematch`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-beforematch}](indices.html#event-beforematch)

[`onbeforetoggle`]{#handler-onbeforetoggle .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

[`beforetoggle`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-beforetoggle}](indices.html#event-beforetoggle)

[`oncancel`]{#handler-oncancel .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLDialogElement/cancel_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLDialogElement/cancel_event "The cancel event fires on a <dialog> when the user instructs the browser that they wish to dismiss the current open dialog. The browser fires this event when the user presses the Esc key.")

Support in all current engines.

::: support
[Firefox98+]{.firefox .yes}[Safari15.4+]{.safari
.yes}[Chrome37+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome AndroidNo]{.chrome_android .no}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

[`cancel`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-cancel}](indices.html#event-cancel)

[`oncanplay`]{#handler-oncanplay .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/canplay_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/canplay_event "The canplay event is fired when the user agent can play the media, but estimates that not enough data has been loaded to play the media up to its end without having to stop for further buffering of content.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

[`canplay`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-media-canplay}](media.html#event-media-canplay)

[`oncanplaythrough`]{#handler-oncanplaythrough .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/canplaythrough_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/canplaythrough_event "The canplaythrough event is fired when the user agent can play the media, and estimates that enough data has been loaded to play the media up to its end without having to stop for further buffering of content.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

[`canplaythrough`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-media-canplaythrough}](media.html#event-media-canplaythrough)

[`onchange`]{#handler-onchange .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLElement/change_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/change_event "The change event is fired for <input>, <select>, and <textarea> elements when the user modifies the element's value. Unlike the input event, the change event is not necessarily fired for each alteration to an element's value.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera9+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

[`change`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-change}](indices.html#event-change)

[`onclick`]{#handler-onclick .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Element/click_event](https://developer.mozilla.org/en-US/docs/Web/API/Element/click_event "An element receives a click event when a pointing device button (such as a mouse's primary mouse button) is both pressed and released while the pointer is located inside the element.")

Support in all current engines.

::: support
[Firefox6+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android6+]{.firefox_android .yes}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

[`click`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-click}](https://w3c.github.io/uievents/#event-type-click){x-internal="event-click"}

[`onclose`]{#handler-onclose .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

[`close`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-close}](indices.html#event-close)

[`oncommand`]{#handler-oncommand .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

[`command`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-command}](indices.html#event-command)

[`oncontextlost`]{#handler-oncontextlost .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

[`contextlost`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-contextlost}](indices.html#event-contextlost)

[`oncontextmenu`]{#handler-oncontextmenu .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

[`contextmenu`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-contextmenu}](https://w3c.github.io/uievents/#event-type-contextmenu){x-internal="event-contextmenu"}

[`oncontextrestored`]{#handler-oncontextrestored .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

[`contextrestored`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-contextrestored}](indices.html#event-contextrestored)

[`oncopy`]{#handler-oncopy .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Element/copy_event](https://developer.mozilla.org/en-US/docs/Web/API/Element/copy_event "The copy event fires when the user initiates a copy action through the browser's user interface.")

Support in all current engines.

::: support
[Firefox22+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

[`copy`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-copy}](https://w3c.github.io/clipboard-apis/#clipboard-event-copy){x-internal="event-copy"}

[`oncuechange`]{#handler-oncuechange .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLTrackElement/cuechange_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLTrackElement/cuechange_event "The cuechange event fires when a TextTrack has changed the currently displaying cues. The event is fired on both the TextTrack and the HTMLTrackElement in which it's being presented, if any.")

Support in all current engines.

::: support
[Firefox68+]{.firefox .yes}[Safari10+]{.safari .yes}[Chrome32+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera19+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)14+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android4.4.3+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android19+]{.opera_android .yes}
:::
::::
:::::

[`cuechange`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-media-cuechange}](media.html#event-media-cuechange)

[`oncut`]{#handler-oncut .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Element/cut_event](https://developer.mozilla.org/en-US/docs/Web/API/Element/cut_event "The cut event is fired when the user has initiated a "cut" action through the browser's user interface.")

Support in all current engines.

::: support
[Firefox22+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

[`cut`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-cut}](https://w3c.github.io/clipboard-apis/#clipboard-event-cut){x-internal="event-cut"}

[`ondblclick`]{#handler-ondblclick .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Element/dblclick_event](https://developer.mozilla.org/en-US/docs/Web/API/Element/dblclick_event "The dblclick event fires when a pointing device button (such as a mouse's primary button) is double-clicked; that is, when it's rapidly clicked twice on a single element within a very short span of time.")

Support in all current engines.

::: support
[Firefox6+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer8+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android6+]{.firefox_android .yes}[Safari iOS1+]{.safari_ios
.yes}[Chrome AndroidNo]{.chrome_android .no}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

[`dblclick`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-dblclick}](https://w3c.github.io/uievents/#event-type-dblclick){x-internal="event-dblclick"}

[`ondrag`]{#handler-ondrag .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

[`drag`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-dnd-drag}](dnd.html#event-dnd-drag)

[`ondragend`]{#handler-ondragend .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

[`dragend`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-dnd-dragend}](dnd.html#event-dnd-dragend)

[`ondragenter`]{#handler-ondragenter .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

[`dragenter`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-dnd-dragenter}](dnd.html#event-dnd-dragenter)

[`ondragleave`]{#handler-ondragleave .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

[`dragleave`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-dnd-dragleave}](dnd.html#event-dnd-dragleave)

[`ondragover`]{#handler-ondragover .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

[`dragover`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-dnd-dragover}](dnd.html#event-dnd-dragover)

[`ondragstart`]{#handler-ondragstart .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

[`dragstart`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-dnd-dragstart}](dnd.html#event-dnd-dragstart)

[`ondrop`]{#handler-ondrop .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

[`drop`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-dnd-drop}](dnd.html#event-dnd-drop)

[`ondurationchange`]{#handler-ondurationchange .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/durationchange_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/durationchange_event "The durationchange event is fired when the duration attribute has been updated.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

[`durationchange`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-media-durationchange}](media.html#event-media-durationchange)

[`onemptied`]{#handler-onemptied .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/emptied_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/emptied_event "The emptied event is fired when the media has become empty; for example, this event is sent if the media has already been loaded (or partially loaded), and the load() method is called to reload it.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

[`emptied`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-media-emptied}](media.html#event-media-emptied)

[`onended`]{#handler-onended .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/ended_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/ended_event "The ended event is fired when playback or streaming has stopped because the end of the media was reached or because no further data is available.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

[`ended`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-media-ended}](media.html#event-media-ended)

[`onformdata`]{#handler-onformdata .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

[`formdata`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-formdata}](indices.html#event-formdata)

[`oninput`]{#handler-oninput .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLElement/input_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/input_event "The input event fires when the value of an <input>, <select>, or <textarea> element has been changed.")

Support in all current engines.

::: support
[Firefox6+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)No]{.edge .no}[Internet Explorer[🔰
9+]{title="Partial implementation."}]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12+]{.opera_android .yes}
:::
::::
:::::

[`input`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-input}](https://w3c.github.io/uievents/#event-type-input){x-internal="event-input"}

[`oninvalid`]{#handler-oninvalid .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

[`invalid`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-invalid}](indices.html#event-invalid)

[`onkeydown`]{#handler-onkeydown .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Element/keydown_event](https://developer.mozilla.org/en-US/docs/Web/API/Element/keydown_event "The keydown event is fired when a key is pressed.")

Support in all current engines.

::: support
[Firefox6+]{.firefox .yes}[Safari1.2+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

[`keydown`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-keydown}](https://w3c.github.io/uievents/#event-type-keydown){x-internal="event-keydown"}

[`onkeypress`]{#handler-onkeypress .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

[`keypress`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-keypress}](https://w3c.github.io/uievents/#event-type-keypress){x-internal="event-keypress"}

[`onkeyup`]{#handler-onkeyup .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Element/keyup_event](https://developer.mozilla.org/en-US/docs/Web/API/Element/keyup_event "The keyup event is fired when a key is released.")

Support in all current engines.

::: support
[Firefox6+]{.firefox .yes}[Safari1.2+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

[`keyup`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-keyup}](https://w3c.github.io/uievents/#event-type-keyup){x-internal="event-keyup"}

[`onloadeddata`]{#handler-onloadeddata .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/loadeddata_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/loadeddata_event "The loadeddata event is fired when the frame at the current playback position of the media has finished loading; often the first frame.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

[`loadeddata`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-media-loadeddata}](media.html#event-media-loadeddata)

[`onloadedmetadata`]{#handler-onloadedmetadata .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/loadedmetadata_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/loadedmetadata_event "The loadedmetadata event is fired when the metadata has been loaded.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

[`loadedmetadata`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-media-loadedmetadata}](media.html#event-media-loadedmetadata)

[`onloadstart`]{#handler-onloadstart .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/loadstart_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/loadstart_event "The loadstart event is fired when the browser has started to load a resource.")

Support in all current engines.

::: support
[Firefox6+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

[`loadstart`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-media-loadstart}](media.html#event-media-loadstart)

[`onmousedown`]{#handler-onmousedown .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Element/mousedown_event](https://developer.mozilla.org/en-US/docs/Web/API/Element/mousedown_event "The mousedown event is fired at an Element when a pointing device button is pressed while the pointer is inside the element.")

Support in all current engines.

::: support
[Firefox6+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome2+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

[`mousedown`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-mousedown}](https://w3c.github.io/uievents/#event-type-mousedown){x-internal="event-mousedown"}

[`onmouseenter`]{#handler-onmouseenter .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Element/mouseenter_event](https://developer.mozilla.org/en-US/docs/Web/API/Element/mouseenter_event "The mouseenter event is fired at an Element when a pointing device (usually a mouse) is initially moved so that its hotspot is within the element at which the event was fired.")

Support in all current engines.

::: support
[Firefox10+]{.firefox .yes}[Safari7+]{.safari .yes}[Chrome30+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

[`mouseenter`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-mouseenter}](https://w3c.github.io/uievents/#event-type-mouseenter){x-internal="event-mouseenter"}

[`onmouseleave`]{#handler-onmouseleave .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Element/mouseleave_event](https://developer.mozilla.org/en-US/docs/Web/API/Element/mouseleave_event "The mouseleave event is fired at an Element when the cursor of a pointing device (usually a mouse) is moved out of it.")

Support in all current engines.

::: support
[Firefox10+]{.firefox .yes}[Safari7+]{.safari .yes}[Chrome30+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

[`mouseleave`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-mouseleave}](https://w3c.github.io/uievents/#event-type-mouseleave){x-internal="event-mouseleave"}

[`onmousemove`]{#handler-onmousemove .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Element/mousemove_event](https://developer.mozilla.org/en-US/docs/Web/API/Element/mousemove_event "The mousemove event is fired at an element when a pointing device (usually a mouse) is moved while the cursor's hotspot is inside it.")

Support in all current engines.

::: support
[Firefox6+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome2+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

[`mousemove`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-mousemove}](https://w3c.github.io/uievents/#event-type-mousemove){x-internal="event-mousemove"}

[`onmouseout`]{#handler-onmouseout .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Element/mouseout_event](https://developer.mozilla.org/en-US/docs/Web/API/Element/mouseout_event "The mouseout event is fired at an Element when a pointing device (usually a mouse) is used to move the cursor so that it is no longer contained within the element or one of its children.")

Support in all current engines.

::: support
[Firefox6+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

[`mouseout`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-mouseout}](https://w3c.github.io/uievents/#event-type-mouseout){x-internal="event-mouseout"}

[`onmouseover`]{#handler-onmouseover .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Element/mouseover_event](https://developer.mozilla.org/en-US/docs/Web/API/Element/mouseover_event "The mouseover event is fired at an Element when a pointing device (such as a mouse or trackpad) is used to move the cursor onto the element or one of its child elements.")

Support in all current engines.

::: support
[Firefox6+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome2+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera9.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

[`mouseover`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-mouseover}](https://w3c.github.io/uievents/#event-type-mouseover){x-internal="event-mouseover"}

[`onmouseup`]{#handler-onmouseup .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Element/mouseup_event](https://developer.mozilla.org/en-US/docs/Web/API/Element/mouseup_event "The mouseup event is fired at an Element when a button on a pointing device (such as a mouse or trackpad) is released while the pointer is located inside it.")

Support in all current engines.

::: support
[Firefox6+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome2+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

[`mouseup`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-mouseup}](https://w3c.github.io/uievents/#event-type-mouseup){x-internal="event-mouseup"}

[`onpaste`]{#handler-onpaste .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Element/paste_event](https://developer.mozilla.org/en-US/docs/Web/API/Element/paste_event "The paste event is fired when the user has initiated a "paste" action through the browser's user interface.")

Support in all current engines.

::: support
[Firefox22+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

[`paste`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-paste}](https://w3c.github.io/clipboard-apis/#clipboard-event-paste){x-internal="event-paste"}

[`onpause`]{#handler-onpause .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/pause_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/pause_event "The pause event is sent when a request to pause an activity is handled and the activity has entered its paused state, most commonly after the media has been paused through a call to the element's pause() method.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

[`pause`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-media-pause}](media.html#event-media-pause)

[`onplay`]{#handler-onplay .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/play_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/play_event "The play event is fired when the paused property is changed from true to false, as a result of the play method, or the autoplay attribute.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

[`play`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-media-play}](media.html#event-media-play)

[`onplaying`]{#handler-onplaying .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/playing_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/playing_event "The playing event is fired after playback is first started, and whenever it is restarted. For example it is fired when playback resumes after having been paused or delayed due to lack of data.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

[`playing`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-media-playing}](media.html#event-media-playing)

[`onprogress`]{#handler-onprogress .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/progress_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/progress_event "The progress event is fired periodically as the browser loads a resource.")

Support in all current engines.

::: support
[Firefox6+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

[`progress`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-media-progress}](media.html#event-media-progress)

[`onratechange`]{#handler-onratechange .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/ratechange_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/ratechange_event "The ratechange event is fired when the playback rate has changed.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

[`ratechange`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-media-ratechange}](media.html#event-media-ratechange)

[`onreset`]{#handler-onreset .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

[`reset`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-reset}](indices.html#event-reset)

[`onscrollend`]{#handler-onscrollend .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

::::::: {.mdn-anno .wrapped .before}
MDN

:::: feature
[Document/scrollend_event](https://developer.mozilla.org/en-US/docs/Web/API/Document/scrollend_event "The scrollend event fires when the document view has completed scrolling. Scrolling is considered completed when the scroll position has no more pending updates and the user has completed their gesture.")

::: support
[Firefox109+]{.firefox .yes}[SafariNo]{.safari .no}[Chrome114+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge114+]{.edge_blink .yes}

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

:::: feature
[Element/scrollend_event](https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollend_event "The scrollend event fires when element scrolling has completed. Scrolling is considered completed when the scroll position has no more pending updates and the user has completed their gesture.")

::: support
[Firefox109+]{.firefox .yes}[SafariNo]{.safari .no}[Chrome114+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge114+]{.edge_blink .yes}

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
:::::::

[`scrollend`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-scrollend}](https://drafts.csswg.org/cssom-view/#eventdef-document-scrollend){x-internal="event-scrollend"}

[`onsecuritypolicyviolation`]{#handler-onsecuritypolicyviolation .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Element/securitypolicyviolation_event](https://developer.mozilla.org/en-US/docs/Web/API/Element/securitypolicyviolation_event "The securitypolicyviolation event is fired when a Content Security Policy is violated.")

Support in all current engines.

::: support
[Firefox63+]{.firefox .yes}[Safari10+]{.safari .yes}[Chrome41+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)15+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

[`securitypolicyviolation`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-securitypolicyviolation}](https://w3c.github.io/webappsec-csp/#eventdef-globaleventhandlers-securitypolicyviolation){x-internal="event-securitypolicyviolation"}

[`onseeked`]{#handler-onseeked .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/seeked_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/seeked_event "The seeked event is fired when a seek operation completed, the current playback position has changed, and the Boolean seeking attribute is changed to false.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

[`seeked`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-media-seeked}](media.html#event-media-seeked)

[`onseeking`]{#handler-onseeking .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/seeking_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/seeking_event "The seeking event is fired when a seek operation starts, meaning the Boolean seeking attribute has changed to true and the media is seeking a new position.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

[`seeking`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-media-seeking}](media.html#event-media-seeking)

[`onselect`]{#handler-onselect .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLInputElement/select_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLInputElement/select_event "The select event fires when some text has been selected.")

Support in all current engines.

::: support
[Firefox6+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::

:::: feature
[HTMLTextAreaElement/select_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLTextAreaElement/select_event "The select event fires when some text has been selected.")

Support in all current engines.

::: support
[Firefox6+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::::

[`select`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-select}](indices.html#event-select)

[`onslotchange`]{#handler-onslotchange .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLSlotElement/slotchange_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLSlotElement/slotchange_event "The slotchange event is fired on an HTMLSlotElement instance (<slot> element) when the node(s) contained in that slot change.")

Support in all current engines.

::: support
[Firefox63+]{.firefox .yes}[Safari10.1+]{.safari
.yes}[Chrome53+]{.chrome .yes}

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

[`slotchange`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-slotchange}](https://dom.spec.whatwg.org/#eventdef-htmlslotelement-slotchange){x-internal="event-slotchange"}

[`onstalled`]{#handler-onstalled .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/stalled_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/stalled_event "The stalled event is fired when the user agent is trying to fetch media data, but data is unexpectedly not forthcoming.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

[`stalled`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-media-stalled}](media.html#event-media-stalled)

[`onsubmit`]{#handler-onsubmit .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLFormElement/submit_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLFormElement/submit_event "The submit event fires when a <form> is submitted.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera8+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

[`submit`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-submit}](indices.html#event-submit)

[`onsuspend`]{#handler-onsuspend .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/suspend_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/suspend_event "The suspend event is fired when media data loading has been suspended.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

[`suspend`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-media-suspend}](media.html#event-media-suspend)

[`ontimeupdate`]{#handler-ontimeupdate .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/timeupdate_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/timeupdate_event "The timeupdate event is fired when the time indicated by the currentTime attribute has been updated.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

[`timeupdate`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-media-timeupdate}](media.html#event-media-timeupdate)

[`ontoggle`]{#handler-ontoggle .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

[`toggle`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-toggle}](indices.html#event-toggle)

[`onvolumechange`]{#handler-onvolumechange .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/volumechange_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/volumechange_event "The volumechange event is fired when either the volume attribute or the muted attribute has changed.")

Support in all current engines.

::: support
[Firefox6+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

[`volumechange`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-media-volumechange}](media.html#event-media-volumechange)

[`onwaiting`]{#handler-onwaiting .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/waiting_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/waiting_event "The waiting event is fired when playback has stopped because of a temporary lack of data.")

Support in all current engines.

::: support
[Firefox6+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

[`waiting`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-media-waiting}](media.html#event-media-waiting)

[`onwebkitanimationend`]{#handler-onwebkitanimationend .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

`webkitAnimationEnd`

[`onwebkitanimationiteration`]{#handler-onwebkitanimationiteration .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

`webkitAnimationIteration`

[`onwebkitanimationstart`]{#handler-onwebkitanimationstart .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

`webkitAnimationStart`

[`onwebkittransitionend`]{#handler-onwebkittransitionend .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

`webkitTransitionEnd`

[`onwheel`]{#handler-onwheel .dfn
dfn-for="HTMLElement,Document,Window,GlobalEventHandlers"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Element/wheel_event](https://developer.mozilla.org/en-US/docs/Web/API/Element/wheel_event "The wheel event fires when the user rotates a wheel button on a pointing device (typically a mouse).")

Support in all current engines.

::: support
[Firefox17+]{.firefox .yes}[Safari7+]{.safari .yes}[Chrome31+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOSNo]{.safari_ios
.no}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

[`wheel`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-wheel}](https://w3c.github.io/uievents/#event-type-wheel){x-internal="event-wheel"}

------------------------------------------------------------------------

The following are the [event
handlers](#event-handlers){#event-handlers-on-elements,-document-objects,-and-window-objects:event-handlers-3}
(and their corresponding [event handler event
types](#event-handler-event-type){#event-handlers-on-elements,-document-objects,-and-window-objects:event-handler-event-type-3})
that must be supported by all [HTML
elements](infrastructure.html#html-elements){#event-handlers-on-elements,-document-objects,-and-window-objects:html-elements-2}
other than
[`body`{#event-handlers-on-elements,-document-objects,-and-window-objects:the-body-element}](sections.html#the-body-element)
and
[`frameset`{#event-handlers-on-elements,-document-objects,-and-window-objects:frameset}](obsolete.html#frameset)
elements, as both [event handler content
attributes](#event-handler-content-attributes){#event-handlers-on-elements,-document-objects,-and-window-objects:event-handler-content-attributes-2}
and [event handler IDL
attributes](#event-handler-idl-attributes){#event-handlers-on-elements,-document-objects,-and-window-objects:event-handler-idl-attributes-3};
that must be supported by all
[`Document`{#event-handlers-on-elements,-document-objects,-and-window-objects:document-3}](dom.html#document)
objects, as [event handler IDL
attributes](#event-handler-idl-attributes){#event-handlers-on-elements,-document-objects,-and-window-objects:event-handler-idl-attributes-4};
and that must be supported by all
[`Window`{#event-handlers-on-elements,-document-objects,-and-window-objects:window-3}](nav-history-apis.html#window)
objects, as [event handler IDL
attributes](#event-handler-idl-attributes){#event-handlers-on-elements,-document-objects,-and-window-objects:event-handler-idl-attributes-5}
on the
[`Window`{#event-handlers-on-elements,-document-objects,-and-window-objects:window-4}](nav-history-apis.html#window)
objects themselves, and with corresponding [event handler content
attributes](#event-handler-content-attributes){#event-handlers-on-elements,-document-objects,-and-window-objects:event-handler-content-attributes-3}
and [event handler IDL
attributes](#event-handler-idl-attributes){#event-handlers-on-elements,-document-objects,-and-window-objects:event-handler-idl-attributes-6}
exposed on all
[`body`{#event-handlers-on-elements,-document-objects,-and-window-objects:the-body-element-2}](sections.html#the-body-element)
and
[`frameset`{#event-handlers-on-elements,-document-objects,-and-window-objects:frameset-2}](obsolete.html#frameset)
elements that are owned by that
[`Window`{#event-handlers-on-elements,-document-objects,-and-window-objects:window-5}](nav-history-apis.html#window)
object\'s [associated
`Document`](nav-history-apis.html#concept-document-window){#event-handlers-on-elements,-document-objects,-and-window-objects:concept-document-window}:

[Event
handler](#event-handlers){#event-handlers-on-elements,-document-objects,-and-window-objects:event-handlers-4}

[Event handler event
type](#event-handler-event-type){#event-handlers-on-elements,-document-objects,-and-window-objects:event-handler-event-type-4}

[`onblur`]{#handler-onblur .dfn dfn-for="GlobalEventHandlers"
dfn-type="attribute"}

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Element/blur_event](https://developer.mozilla.org/en-US/docs/Web/API/Element/blur_event "The blur event fires when an element has lost focus. The event does not bubble, but the related focusout event that follows does bubble.")

Support in all current engines.

::: support
[Firefox24+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::

:::: feature
[Window/blur_event](https://developer.mozilla.org/en-US/docs/Web/API/Window/blur_event "The blur event fires when an element has lost focus.")

Support in all current engines.

::: support
[Firefox6+]{.firefox .yes}[Safari5.1+]{.safari .yes}[Chrome5+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer11]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::::

[`blur`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-blur}](indices.html#event-blur)

[`onerror`]{#handler-onerror .dfn dfn-for="GlobalEventHandlers"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Window/error_event](https://developer.mozilla.org/en-US/docs/Web/API/Window/error_event "The error event is fired on a Window object when a resource failed to load or couldn't be used — for example if a script has an execution error.")

Support in all current engines.

::: support
[Firefox6+]{.firefox .yes}[Safari5.1+]{.safari .yes}[Chrome10+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

[`error`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-error}](indices.html#event-error)

[`onfocus`]{#handler-onfocus .dfn dfn-for="GlobalEventHandlers"
dfn-type="attribute"}

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Element/focus_event](https://developer.mozilla.org/en-US/docs/Web/API/Element/focus_event "The focus event fires when an element has received focus. The event does not bubble, but the related focusin event that follows does bubble.")

Support in all current engines.

::: support
[Firefox24+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::

:::: feature
[Window/focus_event](https://developer.mozilla.org/en-US/docs/Web/API/Window/focus_event "The focus event fires when an element has received focus.")

Support in all current engines.

::: support
[Firefox6+]{.firefox .yes}[Safari5.1+]{.safari .yes}[Chrome5+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer11]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::::

[`focus`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-focus}](indices.html#event-focus)

[`onload`]{#handler-onload .dfn dfn-for="GlobalEventHandlers"
dfn-type="attribute"}

[`load`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-load}](indices.html#event-load)

[`onresize`]{#handler-onresize .dfn dfn-for="GlobalEventHandlers"
dfn-type="attribute"}

[`resize`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-resize}](https://drafts.csswg.org/cssom-view/#eventdef-window-resize){x-internal="event-resize"}

[`onscroll`]{#handler-onscroll .dfn dfn-for="GlobalEventHandlers"
dfn-type="attribute"}

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Document/scroll_event](https://developer.mozilla.org/en-US/docs/Web/API/Document/scroll_event "The scroll event fires when the document view has been scrolled. To detect when scrolling has completed, see the Document: scrollend event. For element scrolling, see Element: scroll event.")

Support in all current engines.

::: support
[Firefox6+]{.firefox .yes}[Safari2+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12+]{.opera_android .yes}
:::
::::

:::: feature
[Element/scroll_event](https://developer.mozilla.org/en-US/docs/Web/API/Element/scroll_event "The scroll event fires when an element has been scrolled. To detect when scrolling has completed, see the Element: scrollend event.")

Support in all current engines.

::: support
[Firefox6+]{.firefox .yes}[Safari1.3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::::

[`scroll`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-scroll}](https://drafts.csswg.org/cssom-view/#eventdef-document-scroll){x-internal="event-scroll"}

We call the
[set](https://infra.spec.whatwg.org/#ordered-set){#event-handlers-on-elements,-document-objects,-and-window-objects:set
x-internal="set"} of the
[names](#event-handler-name){#event-handlers-on-elements,-document-objects,-and-window-objects:event-handler-name}
of the [event
handlers](#event-handlers){#event-handlers-on-elements,-document-objects,-and-window-objects:event-handlers-5}
listed in the first column of this table the [`Window`-reflecting body
element event handler
set]{#window-reflecting-body-element-event-handler-set .dfn}.

------------------------------------------------------------------------

The following are the [event
handlers](#event-handlers){#event-handlers-on-elements,-document-objects,-and-window-objects:event-handlers-6}
(and their corresponding [event handler event
types](#event-handler-event-type){#event-handlers-on-elements,-document-objects,-and-window-objects:event-handler-event-type-5})
that must be supported by
[`Window`{#event-handlers-on-elements,-document-objects,-and-window-objects:window-6}](nav-history-apis.html#window)
objects, as [event handler IDL
attributes](#event-handler-idl-attributes){#event-handlers-on-elements,-document-objects,-and-window-objects:event-handler-idl-attributes-7}
on the
[`Window`{#event-handlers-on-elements,-document-objects,-and-window-objects:window-7}](nav-history-apis.html#window)
objects themselves, and with corresponding [event handler content
attributes](#event-handler-content-attributes){#event-handlers-on-elements,-document-objects,-and-window-objects:event-handler-content-attributes-4}
and [event handler IDL
attributes](#event-handler-idl-attributes){#event-handlers-on-elements,-document-objects,-and-window-objects:event-handler-idl-attributes-8}
exposed on all
[`body`{#event-handlers-on-elements,-document-objects,-and-window-objects:the-body-element-3}](sections.html#the-body-element)
and
[`frameset`{#event-handlers-on-elements,-document-objects,-and-window-objects:frameset-3}](obsolete.html#frameset)
elements that are owned by that
[`Window`{#event-handlers-on-elements,-document-objects,-and-window-objects:window-8}](nav-history-apis.html#window)
object\'s [associated
`Document`](nav-history-apis.html#concept-document-window){#event-handlers-on-elements,-document-objects,-and-window-objects:concept-document-window-2}:

[Event
handler](#event-handlers){#event-handlers-on-elements,-document-objects,-and-window-objects:event-handlers-7}

[Event handler event
type](#event-handler-event-type){#event-handlers-on-elements,-document-objects,-and-window-objects:event-handler-event-type-6}

[`onafterprint`]{#handler-window-onafterprint .dfn
dfn-for="WindowEventHandlers" dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Window/afterprint_event](https://developer.mozilla.org/en-US/docs/Web/API/Window/afterprint_event "The afterprint event is fired after the associated document has started printing or the print preview has been closed.")

Support in all current engines.

::: support
[Firefox6+]{.firefox .yes}[Safari13+]{.safari .yes}[Chrome63+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

[`afterprint`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-afterprint}](indices.html#event-afterprint)

[`onbeforeprint`]{#handler-window-onbeforeprint .dfn
dfn-for="WindowEventHandlers" dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Window/beforeprint_event](https://developer.mozilla.org/en-US/docs/Web/API/Window/beforeprint_event "The beforeprint event is fired when the associated document is about to be printed or previewed for printing.")

Support in all current engines.

::: support
[Firefox6+]{.firefox .yes}[Safari13+]{.safari .yes}[Chrome63+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

[`beforeprint`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-beforeprint}](indices.html#event-beforeprint)

[`onbeforeunload`]{#handler-window-onbeforeunload .dfn
dfn-for="WindowEventHandlers" dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Window/beforeunload_event](https://developer.mozilla.org/en-US/docs/Web/API/Window/beforeunload_event "The beforeunload event is fired when the window, the document and its resources are about to be unloaded. The document is still visible and the event is still cancelable at this point.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer4+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12+]{.opera_android .yes}
:::
::::
:::::

[`beforeunload`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-beforeunload}](indices.html#event-beforeunload)

[`onhashchange`]{#handler-window-onhashchange .dfn
dfn-for="WindowEventHandlers" dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Window/hashchange_event](https://developer.mozilla.org/en-US/docs/Web/API/Window/hashchange_event "The hashchange event is fired when the fragment identifier of the URL has changed (the part of the URL beginning with and following the # symbol).")

Support in all current engines.

::: support
[Firefox3.6+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome8+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer8+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

[`hashchange`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-hashchange}](indices.html#event-hashchange)

[`onlanguagechange`]{#handler-window-onlanguagechange .dfn
dfn-for="WindowEventHandlers" dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Window/languagechange_event](https://developer.mozilla.org/en-US/docs/Web/API/Window/languagechange_event "The languagechange event is fired at the global scope object when the user's preferred language changes.")

Support in all current engines.

::: support
[Firefox32+]{.firefox .yes}[Safari10.1+]{.safari
.yes}[Chrome37+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android4+]{.firefox_android .yes}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet4.0+]{.samsunginternet_android .yes}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

[`languagechange`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-languagechange}](indices.html#event-languagechange)

[`onmessage`]{#handler-window-onmessage .dfn
dfn-for="WindowEventHandlers" dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Window/message_event](https://developer.mozilla.org/en-US/docs/Web/API/Window/message_event "The message event is fired on a Window object when the window receives a message, for example from a call to Window.postMessage() from another browsing context.")

Support in all current engines.

::: support
[Firefox9+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome60+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer8+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS4+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android47+]{.opera_android .yes}
:::
::::
:::::

[`message`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-message}](indices.html#event-message)

[`onmessageerror`]{#handler-window-onmessageerror .dfn
dfn-for="WindowEventHandlers" dfn-type="attribute"}

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/error_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/error_event "The error event is fired when the resource could not be loaded due to an error (for example, a network connectivity problem).")

Support in all current engines.

::: support
[Firefox6+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12+]{.opera_android .yes}
:::
::::

:::: feature
[Window/messageerror_event](https://developer.mozilla.org/en-US/docs/Web/API/Window/messageerror_event "The messageerror event is fired on a Window object when it receives a message that can't be deserialized.")

Support in all current engines.

::: support
[Firefox57+]{.firefox .yes}[Safari16.4+]{.safari
.yes}[Chrome60+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)18]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android47+]{.opera_android .yes}
:::
::::
:::::::

[`messageerror`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-messageerror}](indices.html#event-messageerror)

[`onoffline`]{#handler-window-onoffline .dfn
dfn-for="WindowEventHandlers" dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Window/offline_event](https://developer.mozilla.org/en-US/docs/Web/API/Window/offline_event "The offline event of the Window interface is fired when the browser has lost access to the network and the value of Navigator.onLine switches to false.")

Support in all current engines.

::: support
[Firefox9+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

[`offline`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-offline}](indices.html#event-offline)

[`ononline`]{#handler-window-ononline .dfn dfn-for="WindowEventHandlers"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Window/online_event](https://developer.mozilla.org/en-US/docs/Web/API/Window/online_event "The online event of the Window interface is fired when the browser has gained access to the network and the value of Navigator.onLine switches to true.")

Support in all current engines.

::: support
[Firefox9+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

[`online`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-online}](indices.html#event-online)

[`onpageswap`]{#handler-window-onpageswap .dfn
dfn-for="WindowEventHandlers" dfn-type="attribute"}

[`pageswap`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-pageswap}](indices.html#event-pageswap)

[`onpagehide`]{#handler-window-onpagehide .dfn
dfn-for="WindowEventHandlers" dfn-type="attribute"}

[`pagehide`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-pagehide}](indices.html#event-pagehide)

[`onpagereveal`]{#handler-window-onpagereveal .dfn
dfn-for="WindowEventHandlers" dfn-type="attribute"}

[`pagereveal`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-pagereveal}](indices.html#event-pagereveal)

[`onpageshow`]{#handler-window-onpageshow .dfn
dfn-for="WindowEventHandlers" dfn-type="attribute"}

[`pageshow`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-pageshow}](indices.html#event-pageshow)

[`onpopstate`]{#handler-window-onpopstate .dfn
dfn-for="WindowEventHandlers" dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Window/popstate_event](https://developer.mozilla.org/en-US/docs/Web/API/Window/popstate_event "The popstate event of the Window interface is fired when the active history entry changes while the user navigates the session history. It changes the current history entry to that of the last page the user visited or, if history.pushState() has been used to add a history entry to the history stack, that history entry is used instead.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome5+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11.5+]{.opera_android .yes}
:::
::::
:::::

[`popstate`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-popstate}](indices.html#event-popstate)

[`onrejectionhandled`]{#handler-window-onrejectionhandled .dfn
dfn-for="WindowEventHandlers" dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Window/rejectionhandled_event](https://developer.mozilla.org/en-US/docs/Web/API/Window/rejectionhandled_event "The rejectionhandled event is sent to the script's global scope (usually window but also Worker) whenever a rejected JavaScript Promise is handled late, i.e. when a handler is attached to the promise after its rejection had caused an unhandledrejection event.")

Support in all current engines.

::: support
[Firefox69+]{.firefox .yes}[Safari11+]{.safari .yes}[Chrome49+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari
iOS11.3+]{.safari_ios .yes}[Chrome Android?]{.chrome_android
.unknown}[WebView Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

[`rejectionhandled`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-rejectionhandled}](indices.html#event-rejectionhandled)

[`onstorage`]{#handler-window-onstorage .dfn
dfn-for="WindowEventHandlers" dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Window/storage_event](https://developer.mozilla.org/en-US/docs/Web/API/Window/storage_event "The storage event of the Window interface fires when a storage area (localStorage) has been modified in the context of another document.")

Support in all current engines.

::: support
[Firefox45+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)15+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS4+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

[`storage`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-storage}](indices.html#event-storage)

[`onunhandledrejection`]{#handler-window-onunhandledrejection .dfn
dfn-for="WindowEventHandlers" dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Window/unhandledrejection_event](https://developer.mozilla.org/en-US/docs/Web/API/Window/unhandledrejection_event "The unhandledrejection event is sent to the global scope of a script when a JavaScript Promise that has no rejection handler is rejected; typically, this is the window, but may also be a Worker.")

Support in all current engines.

::: support
[Firefox69+]{.firefox .yes}[Safari11+]{.safari .yes}[Chrome49+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari
iOS11.3+]{.safari_ios .yes}[Chrome Android?]{.chrome_android
.unknown}[WebView Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

[`unhandledrejection`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-unhandledrejection}](indices.html#event-unhandledrejection)

[`onunload`]{#handler-window-onunload .dfn dfn-for="WindowEventHandlers"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Window/unload_event](https://developer.mozilla.org/en-US/docs/Web/API/Window/unload_event "The unload event is fired when the document or a child resource is being unloaded.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera4+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer4+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

[`unload`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-unload}](indices.html#event-unload)

This list of [event
handlers](#event-handlers){#event-handlers-on-elements,-document-objects,-and-window-objects:event-handlers-8}
is reified as [event handler IDL
attributes](#event-handler-idl-attributes){#event-handlers-on-elements,-document-objects,-and-window-objects:event-handler-idl-attributes-9}
through the
[`WindowEventHandlers`{#event-handlers-on-elements,-document-objects,-and-window-objects:windoweventhandlers}](#windoweventhandlers)
interface mixin.

------------------------------------------------------------------------

The following are the [event
handlers](#event-handlers){#event-handlers-on-elements,-document-objects,-and-window-objects:event-handlers-9}
(and their corresponding [event handler event
types](#event-handler-event-type){#event-handlers-on-elements,-document-objects,-and-window-objects:event-handler-event-type-7})
that must be supported on
[`Document`{#event-handlers-on-elements,-document-objects,-and-window-objects:document-4}](dom.html#document)
objects as [event handler IDL
attributes](#event-handler-idl-attributes){#event-handlers-on-elements,-document-objects,-and-window-objects:event-handler-idl-attributes-10}:

[Event
handler](#event-handlers){#event-handlers-on-elements,-document-objects,-and-window-objects:event-handlers-10}

[Event handler event
type](#event-handler-event-type){#event-handlers-on-elements,-document-objects,-and-window-objects:event-handler-event-type-8}

[`onreadystatechange`]{#handler-onreadystatechange .dfn
dfn-for="Document" dfn-type="attribute"}

[`readystatechange`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-readystatechange}](indices.html#event-readystatechange)

[`onvisibilitychange`]{#handler-onvisibilitychange .dfn
dfn-for="Document" dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Document/visibilitychange_event](https://developer.mozilla.org/en-US/docs/Web/API/Document/visibilitychange_event "The visibilitychange event is fired at the document when the contents of its tab have become visible or have been hidden.")

Support in all current engines.

::: support
[Firefox56+]{.firefox .yes}[Safari14.1+]{.safari
.yes}[Chrome62+]{.chrome .yes}

------------------------------------------------------------------------

[Opera49+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)18]{.edge .yes}[Internet Explorer[🔰
10+]{title="Partial implementation."}]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android62+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android46+]{.opera_android .yes}
:::
::::
:::::

[`visibilitychange`{#event-handlers-on-elements,-document-objects,-and-window-objects:event-visibilitychange}](indices.html#event-visibilitychange)

###### [8.1.8.2.1]{.secno} IDL definitions[](#idl-definitions){.self-link}

``` idl
interface mixin GlobalEventHandlers {
  attribute EventHandler onabort;
  attribute EventHandler onauxclick;
  attribute EventHandler onbeforeinput;
  attribute EventHandler onbeforematch;
  attribute EventHandler onbeforetoggle;
  attribute EventHandler onblur;
  attribute EventHandler oncancel;
  attribute EventHandler oncanplay;
  attribute EventHandler oncanplaythrough;
  attribute EventHandler onchange;
  attribute EventHandler onclick;
  attribute EventHandler onclose;
  attribute EventHandler oncommand;
  attribute EventHandler oncontextlost;
  attribute EventHandler oncontextmenu;
  attribute EventHandler oncontextrestored;
  attribute EventHandler oncopy;
  attribute EventHandler oncuechange;
  attribute EventHandler oncut;
  attribute EventHandler ondblclick;
  attribute EventHandler ondrag;
  attribute EventHandler ondragend;
  attribute EventHandler ondragenter;
  attribute EventHandler ondragleave;
  attribute EventHandler ondragover;
  attribute EventHandler ondragstart;
  attribute EventHandler ondrop;
  attribute EventHandler ondurationchange;
  attribute EventHandler onemptied;
  attribute EventHandler onended;
  attribute OnErrorEventHandler onerror;
  attribute EventHandler onfocus;
  attribute EventHandler onformdata;
  attribute EventHandler oninput;
  attribute EventHandler oninvalid;
  attribute EventHandler onkeydown;
  attribute EventHandler onkeypress;
  attribute EventHandler onkeyup;
  attribute EventHandler onload;
  attribute EventHandler onloadeddata;
  attribute EventHandler onloadedmetadata;
  attribute EventHandler onloadstart;
  attribute EventHandler onmousedown;
  [LegacyLenientThis] attribute EventHandler onmouseenter;
  [LegacyLenientThis] attribute EventHandler onmouseleave;
  attribute EventHandler onmousemove;
  attribute EventHandler onmouseout;
  attribute EventHandler onmouseover;
  attribute EventHandler onmouseup;
  attribute EventHandler onpaste;
  attribute EventHandler onpause;
  attribute EventHandler onplay;
  attribute EventHandler onplaying;
  attribute EventHandler onprogress;
  attribute EventHandler onratechange;
  attribute EventHandler onreset;
  attribute EventHandler onresize;
  attribute EventHandler onscroll;
  attribute EventHandler onscrollend;
  attribute EventHandler onsecuritypolicyviolation;
  attribute EventHandler onseeked;
  attribute EventHandler onseeking;
  attribute EventHandler onselect;
  attribute EventHandler onslotchange;
  attribute EventHandler onstalled;
  attribute EventHandler onsubmit;
  attribute EventHandler onsuspend;
  attribute EventHandler ontimeupdate;
  attribute EventHandler ontoggle;
  attribute EventHandler onvolumechange;
  attribute EventHandler onwaiting;
  attribute EventHandler onwebkitanimationend;
  attribute EventHandler onwebkitanimationiteration;
  attribute EventHandler onwebkitanimationstart;
  attribute EventHandler onwebkittransitionend;
  attribute EventHandler onwheel;
};

interface mixin WindowEventHandlers {
  attribute EventHandler onafterprint;
  attribute EventHandler onbeforeprint;
  attribute OnBeforeUnloadEventHandler onbeforeunload;
  attribute EventHandler onhashchange;
  attribute EventHandler onlanguagechange;
  attribute EventHandler onmessage;
  attribute EventHandler onmessageerror;
  attribute EventHandler onoffline;
  attribute EventHandler ononline;
  attribute EventHandler onpagehide;
  attribute EventHandler onpagereveal;
  attribute EventHandler onpageshow;
  attribute EventHandler onpageswap;
  attribute EventHandler onpopstate;
  attribute EventHandler onrejectionhandled;
  attribute EventHandler onstorage;
  attribute EventHandler onunhandledrejection;
  attribute EventHandler onunload;
};
```

##### [8.1.8.3]{.secno} Event firing[](#event-firing){.self-link}

Certain operations and methods are defined as firing events on elements.
For example, the
[`click()`{#event-firing:dom-click}](interaction.html#dom-click) method
on the [`HTMLElement`{#event-firing:htmlelement}](dom.html#htmlelement)
interface is defined as firing a
[`click`{#event-firing:event-click}](https://w3c.github.io/uievents/#event-type-click){x-internal="event-click"}
event on the element. [\[UIEVENTS\]](references.html#refsUIEVENTS)

[Firing a synthetic pointer event named
`e`{.variable}]{#fire-a-synthetic-pointer-event .dfn} at
`target`{.variable}, with an optional `not trusted flag`{.variable},
means running these steps:

1.  Let `event`{.variable} be the result of [creating an
    event](https://dom.spec.whatwg.org/#concept-event-create){#event-firing:creating-an-event
    x-internal="creating-an-event"} using
    [`PointerEvent`{#event-firing:pointerevent}](https://w3c.github.io/pointerevents/#pointerevent-interface){x-internal="pointerevent"}.

2.  Initialize `event`{.variable}\'s
    [`type`{#event-firing:dom-event-type}](https://dom.spec.whatwg.org/#dom-event-type){x-internal="dom-event-type"}
    attribute to `e`{.variable}.

3.  Initialize `event`{.variable}\'s
    [`bubbles`{#event-firing:dom-event-bubbles}](https://dom.spec.whatwg.org/#dom-event-bubbles){x-internal="dom-event-bubbles"}
    and
    [`cancelable`{#event-firing:dom-event-cancelable}](https://dom.spec.whatwg.org/#dom-event-cancelable){x-internal="dom-event-cancelable"}
    attributes to true.

4.  Set `event`{.variable}\'s [composed
    flag](https://dom.spec.whatwg.org/#composed-flag){#event-firing:composed-flag
    x-internal="composed-flag"}.

5.  If the `not trusted flag`{.variable} is set, initialize
    `event`{.variable}\'s
    [`isTrusted`{#event-firing:dom-event-istrusted}](https://dom.spec.whatwg.org/#dom-event-istrusted){x-internal="dom-event-istrusted"}
    attribute to false.

6.  Initialize `event`{.variable}\'s `ctrlKey`, `shiftKey`, `altKey`,
    and `metaKey` attributes according to the current state of the key
    input device, if any (false for any keys that are not available).

7.  Initialize `event`{.variable}\'s
    [`view`{#event-firing:dom-uievent-view}](https://w3c.github.io/uievents/#dom-uievent-view){x-internal="dom-uievent-view"}
    attribute to `target`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#event-firing:node-document
    x-internal="node-document"}\'s
    [`Window`{#event-firing:window}](nav-history-apis.html#window)
    object, if any, and null otherwise.

8.  `event`{.variable}\'s `getModifierState()` method is to return
    values appropriately describing the current state of the key input
    device.

9.  Return the result of
    [dispatching](https://dom.spec.whatwg.org/#concept-event-dispatch){#event-firing:concept-event-dispatch
    x-internal="concept-event-dispatch"} `event`{.variable} at
    `target`{.variable}.

[Firing a `click` event]{#fire-a-click-event .dfn} at
`target`{.variable} means [firing a synthetic pointer event named
`click`](#fire-a-synthetic-pointer-event){#event-firing:fire-a-synthetic-pointer-event}
at `target`{.variable}.

### [8.2]{.secno} The [`WindowOrWorkerGlobalScope`{#windoworworkerglobalscope-mixin:windoworworkerglobalscope}](#windoworworkerglobalscope) mixin[](#windoworworkerglobalscope-mixin){.self-link} {#windoworworkerglobalscope-mixin}

The
[`WindowOrWorkerGlobalScope`{#windoworworkerglobalscope-mixin:windoworworkerglobalscope-2}](#windoworworkerglobalscope)
mixin is for use of APIs that are to be exposed on
[`Window`{#windoworworkerglobalscope-mixin:window}](nav-history-apis.html#window)
and
[`WorkerGlobalScope`{#windoworworkerglobalscope-mixin:workerglobalscope}](workers.html#workerglobalscope)
objects.

Other standards are encouraged to further extend it using
`partial interface mixin `[`WindowOrWorkerGlobalScope`](#windoworworkerglobalscope){#windoworworkerglobalscope-mixin:windoworworkerglobalscope-3}` { … };`
along with an appropriate reference.

``` idl
typedef (DOMString or Function or TrustedScript) TimerHandler;

interface mixin WindowOrWorkerGlobalScope {
  [Replaceable] readonly attribute USVString origin;
  readonly attribute boolean isSecureContext;
  readonly attribute boolean crossOriginIsolated;

  undefined reportError(any e);

  // base64 utility methods
  DOMString btoa(DOMString data);
  ByteString atob(DOMString data);

  // timers
  long setTimeout(TimerHandler handler, optional long timeout = 0, any... arguments);
  undefined clearTimeout(optional long id = 0);
  long setInterval(TimerHandler handler, optional long timeout = 0, any... arguments);
  undefined clearInterval(optional long id = 0);

  // microtask queuing
  undefined queueMicrotask(VoidFunction callback);

  // ImageBitmap
  Promise<ImageBitmap> createImageBitmap(ImageBitmapSource image, optional ImageBitmapOptions options = {});
  Promise<ImageBitmap> createImageBitmap(ImageBitmapSource image, long sx, long sy, long sw, long sh, optional ImageBitmapOptions options = {});

  // structured cloning
  any structuredClone(any value, optional StructuredSerializeOptions options = {});
};
Window includes WindowOrWorkerGlobalScope;
WorkerGlobalScope includes WindowOrWorkerGlobalScope;
```

`self.`[`isSecureContext`](#dom-issecurecontext){#dom-issecurecontext-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[isSecureContext](https://developer.mozilla.org/en-US/docs/Web/API/isSecureContext "The global isSecureContext read-only property returns a boolean indicating whether the current context is secure (true) or not (false).")

Support in all current engines.

::: support
[Firefox49+]{.firefox .yes}[Safari11.1+]{.safari
.yes}[Chrome47+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)15+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

Returns whether or not this global object represents a [secure
context](#secure-context){#windoworworkerglobalscope-mixin:secure-context}.
[\[SECURE-CONTEXTS\]](references.html#refsSECURE-CONTEXTS)

`self.`[`origin`](#dom-origin){#dom-origin-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[origin](https://developer.mozilla.org/en-US/docs/Web/API/origin "The global origin read-only property returns the origin of the global scope, serialized as a string.")

Support in all current engines.

::: support
[Firefox54+]{.firefox .yes}[Safari11+]{.safari .yes}[Chrome59+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)18]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

Returns the global object\'s
[origin](browsers.html#concept-origin){#windoworworkerglobalscope-mixin:concept-origin},
serialized as string.

`self.`[`crossOriginIsolated`](#dom-crossoriginisolated){#dom-crossoriginisolated-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[crossOriginIsolated](https://developer.mozilla.org/en-US/docs/Web/API/crossOriginIsolated "The global crossOriginIsolated read-only property returns a boolean value that indicates whether the website is in a cross-origin isolation state. That state mitigates the risk of side-channel attacks and unlocks a few capabilities:")

Support in all current engines.

::: support
[Firefox72+]{.firefox .yes}[Safari15.2+]{.safari
.yes}[Chrome87+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge87+]{.edge_blink .yes}

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

Returns whether scripts running in this global are allowed to use APIs
that require cross-origin isolation. This depends on the
\`[`Cross-Origin-Opener-Policy`{#windoworworkerglobalscope-mixin:cross-origin-opener-policy-2}](browsers.html#cross-origin-opener-policy-2)\`
and
\`[`Cross-Origin-Embedder-Policy`{#windoworworkerglobalscope-mixin:cross-origin-embedder-policy}](browsers.html#cross-origin-embedder-policy)\`
HTTP response headers and the
\"[`cross-origin-isolated`{#windoworworkerglobalscope-mixin:cross-origin-isolated-feature}](infrastructure.html#cross-origin-isolated-feature)\"
feature.

::: example
Developers are strongly encouraged to use `self.origin` over
`location.origin`. The former returns the
[origin](browsers.html#concept-origin){#windoworworkerglobalscope-mixin:concept-origin-2}
of the environment, the latter of the URL of the environment. Imagine
the following script executing in a document on
`https://stargate.example/`:

``` js
var frame = document.createElement("iframe")
frame.onload = function() {
  var frameWin = frame.contentWindow
  console.log(frameWin.location.origin) // "null"
  console.log(frameWin.origin) // "https://stargate.example"
}
document.body.appendChild(frame)
```

`self.origin` is a more reliable security indicator.
:::

The [`isSecureContext`]{#dom-issecurecontext .dfn
dfn-for="WindowOrWorkerGlobalScope" dfn-type="attribute"} getter steps
are to return true if
[this](https://webidl.spec.whatwg.org/#this){#windoworworkerglobalscope-mixin:this
x-internal="this"}\'s [relevant settings
object](#relevant-settings-object){#windoworworkerglobalscope-mixin:relevant-settings-object}
is a [secure
context](#secure-context){#windoworworkerglobalscope-mixin:secure-context-2},
or false otherwise.

The [`origin`]{#dom-origin .dfn dfn-for="WindowOrWorkerGlobalScope"
dfn-type="attribute"} getter steps are to return
[this](https://webidl.spec.whatwg.org/#this){#windoworworkerglobalscope-mixin:this-2
x-internal="this"}\'s [relevant settings
object](#relevant-settings-object){#windoworworkerglobalscope-mixin:relevant-settings-object-2}\'s
[origin](#concept-settings-object-origin){#windoworworkerglobalscope-mixin:concept-settings-object-origin},
[serialized](browsers.html#ascii-serialisation-of-an-origin){#windoworworkerglobalscope-mixin:ascii-serialisation-of-an-origin}.

The [`crossOriginIsolated`]{#dom-crossoriginisolated .dfn
dfn-for="WindowOrWorkerGlobalScope" dfn-type="attribute"} getter steps
are to return
[this](https://webidl.spec.whatwg.org/#this){#windoworworkerglobalscope-mixin:this-3
x-internal="this"}\'s [relevant settings
object](#relevant-settings-object){#windoworworkerglobalscope-mixin:relevant-settings-object-3}\'s
[cross-origin isolated
capability](#concept-settings-object-cross-origin-isolated-capability){#windoworworkerglobalscope-mixin:concept-settings-object-cross-origin-isolated-capability}.

### [8.3]{.secno} Base64 utility methods[](#atob){.self-link} {#atob}

The [`atob()`{#atob:dom-atob}](#dom-atob) and
[`btoa()`{#atob:dom-btoa}](#dom-btoa) methods allow developers to
transform content to and from the base64 encoding.

In these APIs, for mnemonic purposes, the \"b\" can be considered to
stand for \"binary\", and the \"a\" for \"ASCII\". In practice, though,
for primarily historical reasons, both the input and output of these
functions are Unicode strings.

`result`{.variable}` = self.`[`btoa`](#dom-btoa){#dom-btoa-dev}`(``data`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[btoa](https://developer.mozilla.org/en-US/docs/Web/API/btoa "The btoa() method creates a Base64-encoded ASCII string from a binary string (i.e., a string in which each character in the string is treated as a byte of binary data).")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

Takes the input data, in the form of a Unicode string containing only
characters in the range U+0000 to U+00FF, each representing a binary
byte with values 0x00 to 0xFF respectively, and converts it to its
base64 representation, which it returns.

Throws an
[\"`InvalidCharacterError`\"](https://webidl.spec.whatwg.org/#invalidcharactererror){#atob:invalidcharactererror
x-internal="invalidcharactererror"}
[`DOMException`{#atob:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if the input string contains any out-of-range characters.

`result`{.variable}` = self.`[`atob`](#dom-atob){#dom-atob-dev}`(``data`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[atob](https://developer.mozilla.org/en-US/docs/Web/API/atob "The atob() function decodes a string of data which has been encoded using Base64 encoding. You can use the btoa() method to encode and transmit data which may otherwise cause communication problems, then transmit it and use the atob() method to decode the data again. For example, you can encode, transmit, and decode control characters such as ASCII values 0 through 31.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

Takes the input data, in the form of a Unicode string containing
base64-encoded binary data, decodes it, and returns a string consisting
of characters in the range U+0000 to U+00FF, each representing a binary
byte with values 0x00 to 0xFF respectively, corresponding to that binary
data.

Throws an
[\"`InvalidCharacterError`\"](https://webidl.spec.whatwg.org/#invalidcharactererror){#atob:invalidcharactererror-2
x-internal="invalidcharactererror"}
[`DOMException`{#atob:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if the input string is not valid base64 data.

The
[`btoa(`{#dom-windowbase64-btoa}`data`{.variable}`)`{#dom-windowbase64-btoa}]{#dom-btoa
.dfn dfn-for="WindowOrWorkerGlobalScope" dfn-type="method"} method must
throw an
[\"`InvalidCharacterError`\"](https://webidl.spec.whatwg.org/#invalidcharactererror){#atob:invalidcharactererror-3
x-internal="invalidcharactererror"}
[`DOMException`{#atob:domexception-3}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if `data`{.variable} contains any character whose code point is greater
than U+00FF. Otherwise, the user agent must convert `data`{.variable} to
a byte sequence whose `n`{.variable}th byte is the eight-bit
representation of the `n`{.variable}th code point of `data`{.variable},
and then must apply [forgiving-base64
encode](https://infra.spec.whatwg.org/#forgiving-base64-encode){#atob:forgiving-base64-encode
x-internal="forgiving-base64-encode"} to that byte sequence and return
the result.

The
[`atob(`{#dom-windowbase64-atob}`data`{.variable}`)`{#dom-windowbase64-atob}]{#dom-atob
.dfn dfn-for="WindowOrWorkerGlobalScope" dfn-type="method"} method steps
are:

1.  Let `decodedData`{.variable} be the result of running
    [forgiving-base64
    decode](https://infra.spec.whatwg.org/#forgiving-base64-decode){#atob:forgiving-base64-decode
    x-internal="forgiving-base64-decode"} on `data`{.variable}.

2.  If `decodedData`{.variable} is failure, then throw an
    [\"`InvalidCharacterError`\"](https://webidl.spec.whatwg.org/#invalidcharactererror){#atob:invalidcharactererror-4
    x-internal="invalidcharactererror"}
    [`DOMException`{#atob:domexception-4}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

3.  Return `decodedData`{.variable}.

[← 7.5 Document lifecycle](document-lifecycle.html) --- [Table of
Contents](./) --- [8.4 Dynamic markup insertion
→](dynamic-markup-insertion.html)
