HTML Standard

::: {#head .head .with-buttons}
[![WHATWG](https://resources.whatwg.org/logo.svg){.darkmode-aware
width="100" crossorigin="" height="100"}](https://whatwg.org/){.logo}

# HTML {#html .allcaps}

Living Standard --- Last Updated [2 June 2025]{.pubdate}
:::

[← 8.4 Dynamic markup insertion](dynamic-markup-insertion.html) ---
[Table of Contents](./) --- [8.9 System state and capabilities
→](system-state.html)

1.  ::: {#toc-webappapis}
    1.  [[8.6]{.secno} Timers](timers-and-user-prompts.html#timers)
    2.  [[8.7]{.secno} Microtask
        queuing](timers-and-user-prompts.html#microtask-queuing)
    3.  [[8.8]{.secno} User
        prompts](timers-and-user-prompts.html#user-prompts)
        1.  [[8.8.1]{.secno} Simple
            dialogs](timers-and-user-prompts.html#simple-dialogs)
        2.  [[8.8.2]{.secno}
            Printing](timers-and-user-prompts.html#printing)
    :::

### [8.6]{.secno} Timers[](#timers){.self-link}

The [`setTimeout()`{#timers:dom-settimeout}](#dom-settimeout) and
[`setInterval()`{#timers:dom-setinterval}](#dom-setinterval) methods
allow authors to schedule timer-based callbacks.

`id`{.variable}` = self.`[`setTimeout`](#dom-settimeout){#dom-settimeout-dev}`(``handler`{.variable}` [, ``timeout`{.variable}` [, ...``arguments`{.variable}` ] ])`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[setTimeout](https://developer.mozilla.org/en-US/docs/Web/API/setTimeout "The global setTimeout() method sets a timer which executes a function or specified piece of code once the timer expires.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera4+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer4+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

Schedules a timeout to run `handler`{.variable} after
`timeout`{.variable} milliseconds. Any `arguments`{.variable} are passed
straight through to the `handler`{.variable}.

`id`{.variable}` = self.`[`setTimeout`](#dom-settimeout){#timers:dom-settimeout-2}`(``code`{.variable}` [, ``timeout`{.variable}` ])`

Schedules a timeout to compile and run `code`{.variable} after
`timeout`{.variable} milliseconds.

`self.`[`clearTimeout`](#dom-cleartimeout){#dom-cleartimeout-dev}`(``id`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[clearTimeout](https://developer.mozilla.org/en-US/docs/Web/API/clearTimeout "The global clearTimeout() method cancels a timeout previously established by calling setTimeout().")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera4+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer4+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

Cancels the timeout set with
[`setTimeout()`{#timers:dom-settimeout-3}](#dom-settimeout) or
[`setInterval()`{#timers:dom-setinterval-2}](#dom-setinterval)
identified by `id`{.variable}.

`id`{.variable}` = self.`[`setInterval`](#dom-setinterval){#dom-setinterval-dev}`(``handler`{.variable}` [, ``timeout`{.variable}` [, ...``arguments`{.variable}` ] ])`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[setInterval](https://developer.mozilla.org/en-US/docs/Web/API/setInterval "The setInterval() method, offered on the Window and Worker interfaces, repeatedly calls a function or executes a code snippet, with a fixed time delay between each call.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera4+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer4+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

Schedules a timeout to run `handler`{.variable} every
`timeout`{.variable} milliseconds. Any `arguments`{.variable} are passed
straight through to the `handler`{.variable}.

`id`{.variable}` = self.`[`setInterval`](#dom-setinterval){#timers:dom-setinterval-3}`(``code`{.variable}` [, ``timeout`{.variable}` ])`

Schedules a timeout to compile and run `code`{.variable} every
`timeout`{.variable} milliseconds.

`self.`[`clearInterval`](#dom-clearinterval){#dom-clearinterval-dev}`(``id`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[clearInterval](https://developer.mozilla.org/en-US/docs/Web/API/clearInterval "The global clearInterval() method cancels a timed, repeating action which was previously established by a call to setInterval(). If the parameter provided does not identify a previously established action, this method does nothing.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera4+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer4+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

Cancels the timeout set with
[`setInterval()`{#timers:dom-setinterval-4}](#dom-setinterval) or
[`setTimeout()`{#timers:dom-settimeout-4}](#dom-settimeout) identified
by `id`{.variable}.

Timers can be nested; after five such nested timers, however, the
interval is forced to be at least four milliseconds.

This API does not guarantee that timers will run exactly on schedule.
Delays due to CPU load, other tasks, etc, are to be expected.

Objects that implement the
[`WindowOrWorkerGlobalScope`{#timers:windoworworkerglobalscope}](webappapis.html#windoworworkerglobalscope)
mixin have a [map of setTimeout and setInterval
IDs]{#map-of-settimeout-and-setinterval-ids .dfn
dfn-for="WindowOrWorkerGlobalScope" export=""}, which is an [ordered
map](https://infra.spec.whatwg.org/#ordered-map){#timers:ordered-map
x-internal="ordered-map"}, initially empty. Each
[key](https://infra.spec.whatwg.org/#map-key){#timers:map-key
x-internal="map-key"} in this map is a positive integer, corresponding
to the return value of a
[`setTimeout()`{#timers:dom-settimeout-5}](#dom-settimeout) or
[`setInterval()`{#timers:dom-setinterval-5}](#dom-setinterval) call.
Each [value](https://infra.spec.whatwg.org/#map-value){#timers:map-value
x-internal="map-value"} is a [unique internal
value](common-microsyntaxes.html#unique-internal-value){#timers:unique-internal-value},
corresponding to a key in the object\'s [map of active
timers](#map-of-active-timers){#timers:map-of-active-timers}.

------------------------------------------------------------------------

The
[`setTimeout(`{#dom-windowtimers-setTimeout}`handler`{.variable}`, `{#dom-windowtimers-setTimeout}`timeout`{.variable}`, ...`{#dom-windowtimers-setTimeout}`arguments`{.variable}`)`{#dom-windowtimers-setTimeout}]{#dom-settimeout
.dfn dfn-for="WindowOrWorkerGlobalScope" dfn-type="method"} method steps
are to return the result of running the [timer initialization
steps](#timer-initialisation-steps){#timers:timer-initialisation-steps}
given [this](https://webidl.spec.whatwg.org/#this){#timers:this
x-internal="this"}, `handler`{.variable}, `timeout`{.variable},
`arguments`{.variable}, and false.

The
[`setInterval(`{#dom-windowtimers-setInterval}`handler`{.variable}`, `{#dom-windowtimers-setInterval}`timeout`{.variable}`, ...`{#dom-windowtimers-setInterval}`arguments`{.variable}`)`{#dom-windowtimers-setInterval}]{#dom-setinterval
.dfn dfn-for="WindowOrWorkerGlobalScope" dfn-type="method"} method steps
are to return the result of running the [timer initialization
steps](#timer-initialisation-steps){#timers:timer-initialisation-steps-2}
given [this](https://webidl.spec.whatwg.org/#this){#timers:this-2
x-internal="this"}, `handler`{.variable}, `timeout`{.variable},
`arguments`{.variable}, and true.

The
[`clearTimeout(`{#dom-windowtimers-clearTimeout}`id`{.variable}`)`{#dom-windowtimers-clearTimeout}]{#dom-cleartimeout
.dfn dfn-for="WindowOrWorkerGlobalScope" dfn-type="method"} and
[`clearInterval(`{#dom-windowtimers-clearInterval}`id`{.variable}`)`{#dom-windowtimers-clearInterval}]{#dom-clearinterval
.dfn dfn-for="WindowOrWorkerGlobalScope" dfn-type="method"} method steps
are to
[remove](https://infra.spec.whatwg.org/#map-remove){#timers:map-remove
x-internal="map-remove"}
[this](https://webidl.spec.whatwg.org/#this){#timers:this-3
x-internal="this"}\'s [map of setTimeout and setInterval
IDs](#map-of-settimeout-and-setinterval-ids){#timers:map-of-settimeout-and-setinterval-ids}\[`id`{.variable}\].

Because [`clearTimeout()`{#timers:dom-cleartimeout}](#dom-cleartimeout)
and [`clearInterval()`{#timers:dom-clearinterval}](#dom-clearinterval)
clear entries from the same map, either method can be used to clear
timers created by
[`setTimeout()`{#timers:dom-settimeout-6}](#dom-settimeout) or
[`setInterval()`{#timers:dom-setinterval-6}](#dom-setinterval).

------------------------------------------------------------------------

To perform the [timer initialization steps]{#timer-initialisation-steps
.dfn}, given a
[`WindowOrWorkerGlobalScope`{#timers:windoworworkerglobalscope-2}](webappapis.html#windoworworkerglobalscope)
`global`{.variable}, a string or
[`Function`{#timers:idl-function}](https://webidl.spec.whatwg.org/#common-Function){x-internal="idl-function"}
or
[`TrustedScript`{#timers:tt-trustedscript}](https://w3c.github.io/trusted-types/dist/spec/#trusted-script){x-internal="tt-trustedscript"}
`handler`{.variable}, a number `timeout`{.variable}, a list
`arguments`{.variable}, a boolean `repeat`{.variable}, and optionally
(and only if `repeat`{.variable} is true) a number
`previousId`{.variable}, perform the following steps. They return a
number.

1.  Let `thisArg`{.variable} be `global`{.variable} if that is a
    [`WorkerGlobalScope`{#timers:workerglobalscope}](workers.html#workerglobalscope)
    object; otherwise let `thisArg`{.variable} be the
    [`WindowProxy`{#timers:windowproxy}](nav-history-apis.html#windowproxy)
    that corresponds to `global`{.variable}.

2.  If `previousId`{.variable} was given, let `id`{.variable} be
    `previousId`{.variable}; otherwise, let `id`{.variable} be an
    [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#timers:implementation-defined
    x-internal="implementation-defined"} integer that is greater than
    zero and does not already
    [exist](https://infra.spec.whatwg.org/#map-exists){#timers:map-exists
    x-internal="map-exists"} in `global`{.variable}\'s [map of
    setTimeout and setInterval
    IDs](#map-of-settimeout-and-setinterval-ids){#timers:map-of-settimeout-and-setinterval-ids-2}.

3.  If the [surrounding
    agent](https://tc39.es/ecma262/#surrounding-agent){#timers:surrounding-agent
    x-internal="surrounding-agent"}\'s [event
    loop](webappapis.html#concept-agent-event-loop){#timers:concept-agent-event-loop}\'s
    [currently running
    task](webappapis.html#currently-running-task){#timers:currently-running-task}
    is a task that was created by this algorithm, then let
    `nesting level`{.variable} be the
    [task](webappapis.html#concept-task){#timers:concept-task}\'s [timer
    nesting level](#timer-nesting-level){#timers:timer-nesting-level}.
    Otherwise, let `nesting level`{.variable} be zero.

    The task\'s [timer nesting
    level](#timer-nesting-level){#timers:timer-nesting-level-2} is used
    both for nested calls to
    [`setTimeout()`{#timers:dom-settimeout-7}](#dom-settimeout), and for
    the repeating timers created by
    [`setInterval()`{#timers:dom-setinterval-7}](#dom-setinterval). (Or,
    indeed, for any combination of the two.) In other words, it
    represents nested invocations of this algorithm, not of a particular
    method.

4.  If `timeout`{.variable} is less than 0, then set
    `timeout`{.variable} to 0.

5.  If `nesting level`{.variable} is greater than 5, and
    `timeout`{.variable} is less than 4, then set `timeout`{.variable}
    to 4.

6.  Let `realm`{.variable} be `global`{.variable}\'s [relevant
    realm](webappapis.html#concept-relevant-realm){#timers:concept-relevant-realm}.

7.  Let `initiating script`{.variable} be the [active
    script](webappapis.html#active-script){#timers:active-script}.

8.  Let `uniqueHandle`{.variable} be null.

9.  Let `task`{.variable} be a
    [task](webappapis.html#concept-task){#timers:concept-task-2} that
    runs the following substeps:

    1.  [Assert](https://infra.spec.whatwg.org/#assert){#timers:assert
        x-internal="assert"}: `uniqueHandle`{.variable} is a [unique
        internal
        value](common-microsyntaxes.html#unique-internal-value){#timers:unique-internal-value-2},
        not null.

    2.  If `id`{.variable} does not
        [exist](https://infra.spec.whatwg.org/#map-exists){#timers:map-exists-2
        x-internal="map-exists"} in `global`{.variable}\'s [map of
        setTimeout and setInterval
        IDs](#map-of-settimeout-and-setinterval-ids){#timers:map-of-settimeout-and-setinterval-ids-3},
        then abort these steps.

    3.  If `global`{.variable}\'s [map of setTimeout and setInterval
        IDs](#map-of-settimeout-and-setinterval-ids){#timers:map-of-settimeout-and-setinterval-ids-4}\[`id`{.variable}\]
        does not equal `uniqueHandle`{.variable}, then abort these
        steps.

        This accommodates for the ID having been cleared by a
        [`clearTimeout()`{#timers:dom-cleartimeout-2}](#dom-cleartimeout)
        or
        [`clearInterval()`{#timers:dom-clearinterval-2}](#dom-clearinterval)
        call, and being reused by a subsequent
        [`setTimeout()`{#timers:dom-settimeout-8}](#dom-settimeout) or
        [`setInterval()`{#timers:dom-setinterval-8}](#dom-setinterval)
        call.

    4.  [Record timing info for timer
        handler](https://w3c.github.io/long-animation-frames/#record-timing-info-for-timer-handler){#timers:record-timing-info-for-timer-handler
        x-internal="record-timing-info-for-timer-handler"} given
        `handler`{.variable}, `global`{.variable}\'s [relevant settings
        object](webappapis.html#relevant-settings-object){#timers:relevant-settings-object},
        and `repeat`{.variable}.

    5.  If `handler`{.variable} is a
        [`Function`{#timers:idl-function-2}](https://webidl.spec.whatwg.org/#common-Function){x-internal="idl-function"},
        then
        [invoke](https://webidl.spec.whatwg.org/#invoke-a-callback-function){#timers:es-invoking-callback-functions
        x-internal="es-invoking-callback-functions"}
        `handler`{.variable} given `arguments`{.variable} and
        \"`report`\", and with *[callback this
        value](https://webidl.spec.whatwg.org/#dfn-callback-this-value){x-internal="dfn-callback-this-value"}*
        set to `thisArg`{.variable}.

    6.  Otherwise:

        1.  If `previousId`{.variable} was not given:

            1.  Let `globalName`{.variable} be \"`Window`\" if
                `global`{.variable} is a
                [`Window`{#timers:window}](nav-history-apis.html#window)
                object; \"`WorkerGlobalScope`\" otherwise.

            2.  Let `methodName`{.variable} be \"`setInterval`\" if
                `repeat`{.variable} is true; \"`setTimeout`\" otherwise.

            3.  Let `sink`{.variable} be a concatenation of
                `globalName`{.variable}, U+0020 SPACE, and
                `methodName`{.variable}.

            4.  Set `handler`{.variable} to the result of invoking the
                [Get Trusted Type compliant
                string](https://w3c.github.io/trusted-types/dist/spec/#get-trusted-type-compliant-string-algorithm){#timers:tt-getcompliantstring
                x-internal="tt-getcompliantstring"} algorithm with
                [`TrustedScript`{#timers:tt-trustedscript-2}](https://w3c.github.io/trusted-types/dist/spec/#trusted-script){x-internal="tt-trustedscript"},
                `global`{.variable}, `handler`{.variable},
                `sink`{.variable}, and \"`script`\".

        2.  [Assert](https://infra.spec.whatwg.org/#assert){#timers:assert-2
            x-internal="assert"}: `handler`{.variable} is a string.

        3.  Perform
            [EnsureCSPDoesNotBlockStringCompilation](https://w3c.github.io/webappsec-csp/#can-compile-strings){#timers:csp-ensurecspdoesnotblockstringcompilation
            x-internal="csp-ensurecspdoesnotblockstringcompilation"}(`realm`{.variable},
            « », `handler`{.variable}, `handler`{.variable}, timer, « »,
            `handler`{.variable}). If this throws an exception, catch
            it,
            [report](webappapis.html#report-an-exception){#timers:report-an-exception}
            it for `global`{.variable}, and abort these steps.

        4.  Let `settings object`{.variable} be `global`{.variable}\'s
            [relevant settings
            object](webappapis.html#relevant-settings-object){#timers:relevant-settings-object-2}.

        5.  Let `fetch options`{.variable} be the [default script fetch
            options](webappapis.html#default-script-fetch-options){#timers:default-script-fetch-options}.

        6.  Let `base URL`{.variable} be `settings object`{.variable}\'s
            [API base
            URL](webappapis.html#api-base-url){#timers:api-base-url}.

        7.  If `initiating script`{.variable} is not null, then:

            1.  Set `fetch options`{.variable} to a [script fetch
                options](webappapis.html#script-fetch-options){#timers:script-fetch-options}
                whose [cryptographic
                nonce](webappapis.html#concept-script-fetch-options-nonce){#timers:concept-script-fetch-options-nonce}
                is `initiating script`{.variable}\'s [fetch
                options](webappapis.html#concept-script-script-fetch-options){#timers:concept-script-script-fetch-options}\'s
                [cryptographic
                nonce](webappapis.html#concept-script-fetch-options-nonce){#timers:concept-script-fetch-options-nonce-2},
                [integrity
                metadata](webappapis.html#concept-script-fetch-options-integrity){#timers:concept-script-fetch-options-integrity}
                is the empty string, [parser
                metadata](webappapis.html#concept-script-fetch-options-parser){#timers:concept-script-fetch-options-parser}
                is \"`not-parser-inserted`\", [credentials
                mode](webappapis.html#concept-script-fetch-options-credentials){#timers:concept-script-fetch-options-credentials}
                is `initiating script`{.variable}\'s [fetch
                options](webappapis.html#concept-script-script-fetch-options){#timers:concept-script-script-fetch-options-2}\'s
                [credentials
                mode](webappapis.html#concept-script-fetch-options-credentials){#timers:concept-script-fetch-options-credentials-2},
                [referrer
                policy](webappapis.html#concept-script-fetch-options-referrer-policy){#timers:concept-script-fetch-options-referrer-policy}
                is `initiating script`{.variable}\'s [fetch
                options](webappapis.html#concept-script-script-fetch-options){#timers:concept-script-script-fetch-options-3}\'s
                [referrer
                policy](webappapis.html#concept-script-fetch-options-referrer-policy){#timers:concept-script-fetch-options-referrer-policy-2},
                and [fetch
                priority](webappapis.html#concept-script-fetch-options-fetch-priority){#timers:concept-script-fetch-options-fetch-priority}
                is \"`auto`\".

            2.  Set `base URL`{.variable} to
                `initiating script`{.variable}\'s [base
                URL](webappapis.html#concept-script-base-url){#timers:concept-script-base-url}.

            The effect of these steps ensures that the string
            compilation done by
            [`setTimeout()`{#timers:dom-settimeout-9}](#dom-settimeout)
            and
            [`setInterval()`{#timers:dom-setinterval-9}](#dom-setinterval)
            behaves equivalently to that done by
            [`eval()`{#timers:eval()}](https://tc39.es/ecma262/#sec-eval-x){x-internal="eval()"}.
            That is, [module
            script](webappapis.html#module-script){#timers:module-script}
            fetches via
            [`import()`{#timers:import()}](https://tc39.es/ecma262/#sec-import-calls){x-internal="import()"}
            will behave the same in both contexts.

        8.  Let `script`{.variable} be the result of [creating a classic
            script](webappapis.html#creating-a-classic-script){#timers:creating-a-classic-script}
            given `handler`{.variable}, `settings object`{.variable},
            `base URL`{.variable}, and `fetch options`{.variable}.

        9.  [Run the classic
            script](webappapis.html#run-a-classic-script){#timers:run-a-classic-script}
            `script`{.variable}.

    7.  If `id`{.variable} does not
        [exist](https://infra.spec.whatwg.org/#map-exists){#timers:map-exists-3
        x-internal="map-exists"} in `global`{.variable}\'s [map of
        setTimeout and setInterval
        IDs](#map-of-settimeout-and-setinterval-ids){#timers:map-of-settimeout-and-setinterval-ids-5},
        then abort these steps.

    8.  If `global`{.variable}\'s [map of setTimeout and setInterval
        IDs](#map-of-settimeout-and-setinterval-ids){#timers:map-of-settimeout-and-setinterval-ids-6}\[`id`{.variable}\]
        does not equal `uniqueHandle`{.variable}, then abort these
        steps.

        The ID might have been removed via the author code in
        `handler`{.variable} calling
        [`clearTimeout()`{#timers:dom-cleartimeout-3}](#dom-cleartimeout)
        or
        [`clearInterval()`{#timers:dom-clearinterval-3}](#dom-clearinterval).
        Checking that uniqueHandle isn\'t different accounts for the
        possibility of the ID, after having been cleared, being reused
        by a subsequent
        [`setTimeout()`{#timers:dom-settimeout-10}](#dom-settimeout) or
        [`setInterval()`{#timers:dom-setinterval-10}](#dom-setinterval)
        call.

    9.  If `repeat`{.variable} is true, then perform the [timer
        initialization
        steps](#timer-initialisation-steps){#timers:timer-initialisation-steps-3}
        again, given `global`{.variable}, `handler`{.variable},
        `timeout`{.variable}, `arguments`{.variable}, true, and
        `id`{.variable}.

    10. Otherwise,
        [remove](https://infra.spec.whatwg.org/#map-remove){#timers:map-remove-2
        x-internal="map-remove"} `global`{.variable}\'s [map of
        setTimeout and setInterval
        IDs](#map-of-settimeout-and-setinterval-ids){#timers:map-of-settimeout-and-setinterval-ids-7}\[`id`{.variable}\].

10. Increment `nesting level`{.variable} by one.

11. Set `task`{.variable}\'s [timer nesting level]{#timer-nesting-level
    .dfn} to `nesting level`{.variable}.

12. Let `completionStep`{.variable} be an algorithm step which [queues a
    global
    task](webappapis.html#queue-a-global-task){#timers:queue-a-global-task}
    on the [timer task source]{#timer-task-source .dfn export=""} given
    `global`{.variable} to run `task`{.variable}.

13. Set `uniqueHandle`{.variable} to the result of [running steps after
    a
    timeout](#run-steps-after-a-timeout){#timers:run-steps-after-a-timeout}
    given `global`{.variable}, \"`setTimeout/setInterval`\",
    `timeout`{.variable}, and `completionStep`{.variable}.

14. [Set](https://infra.spec.whatwg.org/#map-set){#timers:map-set
    x-internal="map-set"} `global`{.variable}\'s [map of setTimeout and
    setInterval
    IDs](#map-of-settimeout-and-setinterval-ids){#timers:map-of-settimeout-and-setinterval-ids-8}\[`id`{.variable}\]
    to `uniqueHandle`{.variable}.

15. Return `id`{.variable}.

Argument conversion as defined by Web IDL (for example, invoking
`toString()` methods on objects passed as the first argument) happens in
the algorithms defined in Web IDL, before this algorithm is invoked.

::: example
So for example, the following rather silly code will result in the log
containing \"`ONE TWO `\":

``` js
var log = '';
function logger(s) { log += s + ' '; }

setTimeout({ toString: function () {
  setTimeout("logger('ONE')", 100);
  return "logger('TWO')";
} }, 100);
```
:::

::: example
To run tasks of several milliseconds back to back without any delay,
while still yielding back to the browser to avoid starving the user
interface (and to avoid the browser killing the script for hogging the
CPU), simply queue the next timer before performing work:

``` js
function doExpensiveWork() {
  var done = false;
  // ...
  // this part of the function takes up to five milliseconds
  // set done to true if we're done
  // ...
  return done;
}

function rescheduleWork() {
  var id = setTimeout(rescheduleWork, 0); // preschedule next iteration
  if (doExpensiveWork())
    clearTimeout(id); // clear the timeout if we don't need it
}

function scheduleWork() {
  setTimeout(rescheduleWork, 0);
}

scheduleWork(); // queues a task to do lots of work
```
:::

Objects that implement the
[`WindowOrWorkerGlobalScope`{#timers:windoworworkerglobalscope-3}](webappapis.html#windoworworkerglobalscope)
mixin have a [map of active timers]{#map-of-active-timers .dfn
dfn-for="WindowOrWorkerGlobalScope" export=""}, which is an [ordered
map](https://infra.spec.whatwg.org/#ordered-map){#timers:ordered-map-2
x-internal="ordered-map"}, initially empty. Each
[key](https://infra.spec.whatwg.org/#map-key){#timers:map-key-2
x-internal="map-key"} in this map is a [unique internal
value](common-microsyntaxes.html#unique-internal-value){#timers:unique-internal-value-3}
that represents a timer, and each
[value](https://infra.spec.whatwg.org/#map-value){#timers:map-value-2
x-internal="map-value"} is a
[`DOMHighResTimeStamp`{#timers:domhighrestimestamp}](https://w3c.github.io/hr-time/#dom-domhighrestimestamp){x-internal="domhighrestimestamp"},
representing the expiry time for that timer.

To [run steps after a timeout]{#run-steps-after-a-timeout .dfn
export=""}, given a
[`WindowOrWorkerGlobalScope`{#timers:windoworworkerglobalscope-4}](webappapis.html#windoworworkerglobalscope)
`global`{.variable}, a string `orderingIdentifier`{.variable}, a number
`milliseconds`{.variable}, and a set of steps
`completionSteps`{.variable}, perform the following steps. They return a
[unique internal
value](common-microsyntaxes.html#unique-internal-value){#timers:unique-internal-value-4}.

1.  Let `timerKey`{.variable} be a new [unique internal
    value](common-microsyntaxes.html#unique-internal-value){#timers:unique-internal-value-5}.

2.  Let `startTime`{.variable} be the [current high resolution
    time](https://w3c.github.io/hr-time/#dfn-current-high-resolution-time){#timers:current-high-resolution-time
    x-internal="current-high-resolution-time"} given
    `global`{.variable}.

3.  [Set](https://infra.spec.whatwg.org/#map-set){#timers:map-set-2
    x-internal="map-set"} `global`{.variable}\'s [map of active
    timers](#map-of-active-timers){#timers:map-of-active-timers-2}\[`timerKey`{.variable}\]
    to `startTime`{.variable} plus `milliseconds`{.variable}.

4.  Run the following steps [in
    parallel](infrastructure.html#in-parallel){#timers:in-parallel}:

    1.  If `global`{.variable} is a
        [`Window`{#timers:window-2}](nav-history-apis.html#window)
        object, wait until `global`{.variable}\'s [associated
        `Document`](nav-history-apis.html#concept-document-window){#timers:concept-document-window}
        has been [fully
        active](document-sequences.html#fully-active){#timers:fully-active}
        for a further `milliseconds`{.variable} milliseconds (not
        necessarily consecutively).

        Otherwise, `global`{.variable} is a
        [`WorkerGlobalScope`{#timers:workerglobalscope-2}](workers.html#workerglobalscope)
        object; wait until `milliseconds`{.variable} milliseconds have
        passed with the worker not suspended (not necessarily
        consecutively).

    2.  Wait until any invocations of this algorithm that had the same
        `global`{.variable} and `orderingIdentifier`{.variable}, that
        started before this one, and whose `milliseconds`{.variable} is
        less than or equal to this one\'s, have completed.

    3.  Optionally, wait a further
        [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#timers:implementation-defined-2
        x-internal="implementation-defined"} length of time.

        This is intended to allow user agents to pad timeouts as needed
        to optimize the power usage of the device. For example, some
        processors have a low-power mode where the granularity of timers
        is reduced; on such platforms, user agents can slow timers down
        to fit this schedule instead of requiring the processor to use
        the more accurate mode with its associated higher power usage.

    4.  Perform `completionSteps`{.variable}.

    5.  [Remove](https://infra.spec.whatwg.org/#map-remove){#timers:map-remove-3
        x-internal="map-remove"} `global`{.variable}\'s [map of active
        timers](#map-of-active-timers){#timers:map-of-active-timers-3}\[`timerKey`{.variable}\].

5.  Return `timerKey`{.variable}.

[Run steps after a
timeout](#run-steps-after-a-timeout){#timers:run-steps-after-a-timeout-2}
is meant to be used by other specifications that want to execute
developer-supplied code after a developer-supplied timeout, in a similar
manner to [`setTimeout()`{#timers:dom-settimeout-11}](#dom-settimeout).
(Note, however, it does not have the nesting and clamping behavior of
[`setTimeout()`{#timers:dom-settimeout-12}](#dom-settimeout).) Such
specifications can choose an `orderingIdentifier`{.variable} to ensure
ordering within their specification\'s timeouts, while not constraining
ordering with respect to other specification\'s timeouts.

### [8.7]{.secno} Microtask queuing[](#microtask-queuing){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[queueMicrotask](https://developer.mozilla.org/en-US/docs/Web/API/queueMicrotask "The queueMicrotask() method, which is exposed on the Window or Worker interface, queues a microtask to be executed at a safe time prior to control returning to the browser's event loop.")

Support in all current engines.

::: support
[Firefox69+]{.firefox .yes}[Safari12.1+]{.safari
.yes}[Chrome71+]{.chrome .yes}

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

`self`{.variable}`.`[`queueMicrotask`](#dom-queuemicrotask){#dom-queuemicrotask-dev}`(``callback`{.variable}`)`

:   [Queues](webappapis.html#queue-a-microtask){#microtask-queuing:queue-a-microtask}
    a
    [microtask](webappapis.html#microtask){#microtask-queuing:microtask}
    to run the given `callback`{.variable}.

The [`queueMicrotask(``callback`{.variable}`)`]{#dom-queuemicrotask .dfn
dfn-for="WindowOrWorkerGlobalScope" dfn-type="method"} method must
[queue a
microtask](webappapis.html#queue-a-microtask){#microtask-queuing:queue-a-microtask-2}
to
[invoke](https://webidl.spec.whatwg.org/#invoke-a-callback-function){#microtask-queuing:es-invoking-callback-functions
x-internal="es-invoking-callback-functions"} `callback`{.variable} with
« » and \"`report`\".

The
[`queueMicrotask()`{#microtask-queuing:dom-queuemicrotask}](#dom-queuemicrotask)
method allows authors to schedule a callback on the [microtask
queue](webappapis.html#microtask-queue){#microtask-queuing:microtask-queue}.
This allows their code to run once the [JavaScript execution context
stack](https://tc39.es/ecma262/#execution-context-stack){#microtask-queuing:javascript-execution-context-stack
x-internal="javascript-execution-context-stack"} is next empty, which
happens once all currently executing synchronous JavaScript has run to
completion. This doesn\'t yield control back to the [event
loop](webappapis.html#event-loop){#microtask-queuing:event-loop}, as
would be the case when using, for example,
[`setTimeout(`{#microtask-queuing:dom-settimeout}`f`{.variable}`, 0)`{#microtask-queuing:dom-settimeout}](#dom-settimeout).

Authors ought to be aware that scheduling a lot of microtasks has the
same performance downsides as running a lot of synchronous code. Both
will prevent the browser from doing its own work, such as rendering. In
many cases,
[`requestAnimationFrame()`{#microtask-queuing:dom-animationframeprovider-requestanimationframe}](imagebitmap-and-animations.html#dom-animationframeprovider-requestanimationframe)
or
[`requestIdleCallback()`{#microtask-queuing:requestidlecallback()}](https://w3c.github.io/requestidlecallback/#the-requestidlecallback-method){x-internal="requestidlecallback()"}
is a better choice. In particular, if the goal is to run code before the
next rendering cycle, that is the purpose of
[`requestAnimationFrame()`{#microtask-queuing:dom-animationframeprovider-requestanimationframe-2}](imagebitmap-and-animations.html#dom-animationframeprovider-requestanimationframe).

As can be seen from the following examples, the best way of thinking
about
[`queueMicrotask()`{#microtask-queuing:dom-queuemicrotask-2}](#dom-queuemicrotask)
is as a mechanism for rearranging synchronous code, effectively placing
the queued code immediately after the currently executing synchronous
JavaScript has run to completion.

::: example
The most common reason for using
[`queueMicrotask()`{#microtask-queuing:dom-queuemicrotask-3}](#dom-queuemicrotask)
is to create consistent ordering, even in the cases where information is
available synchronously, without introducing undue delay.

For example, consider a custom element firing a `load` event, that also
maintains an internal cache of previously-loaded data. A naïve
implementation might look like:

``` js
MyElement.prototype.loadData = function (url) {
  if (this._cache[url]) {
    this._setData(this._cache[url]);
    this.dispatchEvent(new Event("load"));
  } else {
    fetch(url).then(res => res.arrayBuffer()).then(data => {
      this._cache[url] = data;
      this._setData(data);
      this.dispatchEvent(new Event("load"));
    });
  }
};
```

This naïve implementation is problematic, however, in that it causes its
users to experience inconsistent behavior. For example, code such as

``` js
element.addEventListener("load", () => console.log("loaded"));
console.log("1");
element.loadData();
console.log("2");
```

will sometimes log \"1, 2, loaded\" (if the data needs to be fetched),
and sometimes log \"1, loaded, 2\" (if the data is already cached).
Similarly, after the call to `loadData()`, it will be inconsistent
whether or not the data is set on the element.

To get a consistent ordering,
[`queueMicrotask()`{#microtask-queuing:dom-queuemicrotask-4}](#dom-queuemicrotask)
can be used:

``` js
MyElement.prototype.loadData = function (url) {
  if (this._cache[url]) {
    queueMicrotask(() => {
      this._setData(this._cache[url]);
      this.dispatchEvent(new Event("load"));
    });
  } else {
    fetch(url).then(res => res.arrayBuffer()).then(data => {
      this._cache[url] = data;
      this._setData(data);
      this.dispatchEvent(new Event("load"));
    });
  }
};
```

By essentially rearranging the queued code to be after the [JavaScript
execution context
stack](https://tc39.es/ecma262/#execution-context-stack){#microtask-queuing:javascript-execution-context-stack-2
x-internal="javascript-execution-context-stack"} empties, this ensures a
consistent ordering and update of the element\'s state.
:::

::: example
Another interesting use of
[`queueMicrotask()`{#microtask-queuing:dom-queuemicrotask-5}](#dom-queuemicrotask)
is to allow uncoordinated \"batching\" of work by multiple callers. For
example, consider a library function that wants to send data somewhere
as soon as possible, but doesn\'t want to make multiple network requests
if doing so is easily avoidable. One way to balance this would be like
so:

``` js
const queuedToSend = [];

function sendData(data) {
  queuedToSend.push(data);

  if (queuedToSend.length === 1) {
    queueMicrotask(() => {
      const stringToSend = JSON.stringify(queuedToSend);
      queuedToSend.length = 0;

      fetch("/endpoint", stringToSend);
    });
  }
}
```

With this architecture, multiple subsequent calls to `sendData()` within
the currently executing synchronous JavaScript will be batched together
into one
[`fetch()`{#microtask-queuing:fetch()}](https://fetch.spec.whatwg.org/#dom-global-fetch){x-internal="fetch()"}
call, but with no intervening event loop tasks preempting the fetch (as
would have happened with similar code that instead used
[`setTimeout()`{#microtask-queuing:dom-settimeout-2}](#dom-settimeout)).
:::

### [8.8]{.secno} User prompts[](#user-prompts){.self-link}

#### [8.8.1]{.secno} Simple dialogs[](#simple-dialogs){.self-link}

`window`{.variable}`.`[`alert`](#dom-alert){#dom-alert-dev}`(``message`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Window/alert](https://developer.mozilla.org/en-US/docs/Web/API/Window/alert "window.alert() instructs the browser to display a dialog with an optional message, and to wait until the user dismisses the dialog.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera3+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer4+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

Displays a modal alert with the given message, and waits for the user to
dismiss it.

`result`{.variable}` = ``window`{.variable}`.`[`confirm`](#dom-confirm){#dom-confirm-dev}`(``message`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Window/confirm](https://developer.mozilla.org/en-US/docs/Web/API/Window/confirm "window.confirm() instructs the browser to display a dialog with an optional message, and to wait until the user either confirms or cancels the dialog.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera3+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer4+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android1+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

Displays a modal OK/Cancel prompt with the given message, waits for the
user to dismiss it, and returns true if the user clicks OK and false if
the user clicks Cancel.

`result`{.variable}` = ``window`{.variable}`.`[`prompt`](#dom-prompt){#dom-prompt-dev}`(``message`{.variable}` [, ``default`{.variable}`])`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Window/prompt](https://developer.mozilla.org/en-US/docs/Web/API/Window/prompt "window.prompt() instructs the browser to display a dialog with an optional message prompting the user to input some text, and to wait until the user either submits the text or cancels the dialog.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera3+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer4+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

Displays a modal text control prompt with the given message, waits for
the user to dismiss it, and returns the value that the user entered. If
the user cancels the prompt, then returns null instead. If the second
argument is present, then the given value is used as a default.

Logic that depends on
[tasks](webappapis.html#concept-task){#simple-dialogs:concept-task} or
[microtasks](webappapis.html#microtask){#simple-dialogs:microtask}, such
as [media
elements](media.html#media-element){#simple-dialogs:media-element}
loading their [media
data](media.html#media-data){#simple-dialogs:media-data}, are stalled
when these methods are invoked.

The [`alert()`]{#dom-alert-noargs .dfn dfn-for="Window"
dfn-type="method"} and [`alert(``message`{.variable}`)`]{#dom-alert .dfn
dfn-for="Window" dfn-type="method"} method steps are:

1.  If we [cannot show simple
    dialogs](#cannot-show-simple-dialogs){#simple-dialogs:cannot-show-simple-dialogs}
    for
    [this](https://webidl.spec.whatwg.org/#this){#simple-dialogs:this
    x-internal="this"}, then return.

2.  If the method was invoked with no arguments, then let
    `message`{.variable} be the empty string; otherwise, let
    `message`{.variable} be the method\'s first argument.

3.  Set `message`{.variable} to the result of [normalizing
    newlines](https://infra.spec.whatwg.org/#normalize-newlines){#simple-dialogs:normalize-newlines
    x-internal="normalize-newlines"} given `message`{.variable}.

4.  Set `message`{.variable} to the result of [optionally
    truncating](#optionally-truncate-a-simple-dialog-string){#simple-dialogs:optionally-truncate-a-simple-dialog-string}
    `message`{.variable}.

5.  Let `userPromptHandler`{.variable} be [WebDriver BiDi user prompt
    opened](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-user-prompt-opened){#simple-dialogs:webdriver-bidi-user-prompt-opened
    x-internal="webdriver-bidi-user-prompt-opened"} with
    [this](https://webidl.spec.whatwg.org/#this){#simple-dialogs:this-2
    x-internal="this"}, \"`alert`\", and `message`{.variable}.

6.  If `userPromptHandler`{.variable} is \"`none`\", then:

    1.  Show `message`{.variable} to the user, treating U+000A LF as a
        line break.

    2.  Optionally,
        [pause](webappapis.html#pause){#simple-dialogs:pause} while
        waiting for the user to acknowledge the message.

7.  Invoke [WebDriver BiDi user prompt
    closed](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-user-prompt-closed){#simple-dialogs:webdriver-bidi-user-prompt-closed
    x-internal="webdriver-bidi-user-prompt-closed"} with
    [this](https://webidl.spec.whatwg.org/#this){#simple-dialogs:this-3
    x-internal="this"}, \"`alert`\", and true.

This method is defined using two overloads, instead of using an optional
argument, for historical reasons. The practical impact of this is that
`alert(undefined)`{.js} is treated as `alert("undefined")`{.js}, but
`alert()`{.js} is treated as `alert("")`{.js}.

The [`confirm(``message`{.variable}`)`]{#dom-confirm .dfn
dfn-for="Window" dfn-type="method"} method steps are:

1.  If we [cannot show simple
    dialogs](#cannot-show-simple-dialogs){#simple-dialogs:cannot-show-simple-dialogs-2}
    for
    [this](https://webidl.spec.whatwg.org/#this){#simple-dialogs:this-4
    x-internal="this"}, then return false.

2.  Set `message`{.variable} to the result of [normalizing
    newlines](https://infra.spec.whatwg.org/#normalize-newlines){#simple-dialogs:normalize-newlines-2
    x-internal="normalize-newlines"} given `message`{.variable}.

3.  Set `message`{.variable} to the result of [optionally
    truncating](#optionally-truncate-a-simple-dialog-string){#simple-dialogs:optionally-truncate-a-simple-dialog-string-2}
    `message`{.variable}.

4.  Show `message`{.variable} to the user, treating U+000A LF as a line
    break, and ask the user to respond with a positive or negative
    response.

5.  Let `userPromptHandler`{.variable} be [WebDriver BiDi user prompt
    opened](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-user-prompt-opened){#simple-dialogs:webdriver-bidi-user-prompt-opened-2
    x-internal="webdriver-bidi-user-prompt-opened"} with
    [this](https://webidl.spec.whatwg.org/#this){#simple-dialogs:this-5
    x-internal="this"}, \"`confirm`\", and `message`{.variable}.

6.  Let `accepted`{.variable} be false.

7.  If `userPromptHandler`{.variable} is \"`none`\", then:

    1.  [Pause](webappapis.html#pause){#simple-dialogs:pause-2} until
        the user responds either positively or negatively.

    2.  If the user responded positively, then set `accepted`{.variable}
        to true.

8.  If `userPromptHandler`{.variable} is \"`accept`\", then set
    `accepted`{.variable} to true.

9.  Invoke [WebDriver BiDi user prompt
    closed](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-user-prompt-closed){#simple-dialogs:webdriver-bidi-user-prompt-closed-2
    x-internal="webdriver-bidi-user-prompt-closed"} with
    [this](https://webidl.spec.whatwg.org/#this){#simple-dialogs:this-6
    x-internal="this"}, \"`confirm`\", and `accepted`{.variable}.

10. Return `accepted`{.variable}.

The
[`prompt(``message`{.variable}`, ``default`{.variable}`)`]{#dom-prompt
.dfn dfn-for="Window" dfn-type="method"} method steps are:

1.  If we [cannot show simple
    dialogs](#cannot-show-simple-dialogs){#simple-dialogs:cannot-show-simple-dialogs-3}
    for
    [this](https://webidl.spec.whatwg.org/#this){#simple-dialogs:this-7
    x-internal="this"}, then return null.

2.  Set `message`{.variable} to the result of [normalizing
    newlines](https://infra.spec.whatwg.org/#normalize-newlines){#simple-dialogs:normalize-newlines-3
    x-internal="normalize-newlines"} given `message`{.variable}.

3.  Set `message`{.variable} to the result of [optionally
    truncating](#optionally-truncate-a-simple-dialog-string){#simple-dialogs:optionally-truncate-a-simple-dialog-string-3}
    `message`{.variable}.

4.  Set `default`{.variable} to the result of [optionally
    truncating](#optionally-truncate-a-simple-dialog-string){#simple-dialogs:optionally-truncate-a-simple-dialog-string-4}
    `default`{.variable}.

5.  Show `message`{.variable} to the user, treating U+000A LF as a line
    break, and ask the user to either respond with a string value or
    abort. The response must be defaulted to the value given by
    `default`{.variable}.

6.  Let `userPromptHandler`{.variable} be [WebDriver BiDi user prompt
    opened](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-user-prompt-opened){#simple-dialogs:webdriver-bidi-user-prompt-opened-3
    x-internal="webdriver-bidi-user-prompt-opened"} with
    [this](https://webidl.spec.whatwg.org/#this){#simple-dialogs:this-8
    x-internal="this"}, \"`prompt`\", and `message`{.variable}.

7.  Let `result`{.variable} be null.

8.  If `userPromptHandler`{.variable} is \"`none`\", then:

    1.  [Pause](webappapis.html#pause){#simple-dialogs:pause-3} while
        waiting for the user\'s response.

    2.  If the user did not abort, then set `result`{.variable} to the
        string that the user responded with.

9.  Otherwise, if `userPromptHandler`{.variable} is \"`accept`\", then
    set `result`{.variable} to the empty string.

10. Invoke [WebDriver BiDi user prompt
    closed](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-user-prompt-closed){#simple-dialogs:webdriver-bidi-user-prompt-closed-3
    x-internal="webdriver-bidi-user-prompt-closed"} with
    [this](https://webidl.spec.whatwg.org/#this){#simple-dialogs:this-9
    x-internal="this"}, \"`prompt`\", false if `result`{.variable} is
    null or true otherwise, and `result`{.variable}.

11. Return `result`{.variable}.

To [optionally truncate a simple dialog
string]{#optionally-truncate-a-simple-dialog-string .dfn}
`s`{.variable}, return either `s`{.variable} itself or some string
derived from `s`{.variable} that is shorter. User agents should not
provide UI for displaying the elided portion of `s`{.variable}, as this
makes it too easy for abusers to create dialogs of the form \"Important
security alert! Click \'Show More\' for full details!\".

For example, a user agent might want to only display the first 100
characters of a message. Or, a user agent might replace the middle of
the string with \"...\". These types of modifications can be useful in
limiting the abuse potential of unnaturally large, trustworthy-looking
system dialogs.

We [cannot show simple dialogs]{#cannot-show-simple-dialogs .dfn} for a
[`Window`{#simple-dialogs:window}](nav-history-apis.html#window)
`window`{.variable} when the following algorithm returns true:

1.  If the [active sandboxing flag
    set](browsers.html#active-sandboxing-flag-set){#simple-dialogs:active-sandboxing-flag-set}
    of `window`{.variable}\'s [associated
    `Document`](nav-history-apis.html#concept-document-window){#simple-dialogs:concept-document-window}
    has the [sandboxed modals
    flag](browsers.html#sandboxed-modals-flag){#simple-dialogs:sandboxed-modals-flag}
    set, then return true.

2.  If `window`{.variable}\'s [relevant settings
    object](webappapis.html#relevant-settings-object){#simple-dialogs:relevant-settings-object}\'s
    [origin](webappapis.html#concept-settings-object-origin){#simple-dialogs:concept-settings-object-origin}
    and `window`{.variable}\'s [relevant settings
    object](webappapis.html#relevant-settings-object){#simple-dialogs:relevant-settings-object-2}\'s
    [top-level
    origin](webappapis.html#concept-environment-top-level-origin){#simple-dialogs:concept-environment-top-level-origin}
    are not [same
    origin-domain](browsers.html#same-origin-domain){#simple-dialogs:same-origin-domain},
    then return true.

3.  If `window`{.variable}\'s [relevant
    agent](webappapis.html#relevant-agent){#simple-dialogs:relevant-agent}\'s
    [event
    loop](webappapis.html#concept-agent-event-loop){#simple-dialogs:concept-agent-event-loop}\'s
    [termination nesting
    level](document-lifecycle.html#termination-nesting-level){#simple-dialogs:termination-nesting-level}
    is nonzero, then optionally return true.

4.  Optionally, return true. (For example, the user agent might give the
    user the option to ignore all modal dialogs, and would thus abort at
    this step whenever the method was invoked.)

5.  Return false.

#### [8.8.2]{.secno} Printing[](#printing){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Window/print](https://developer.mozilla.org/en-US/docs/Web/API/Window/print "Opens the print dialog to print the current document.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1.1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android114+]{.firefox_android .yes}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

`window`{.variable}`.`[`print`](#dom-print){#dom-print-dev}`()`

:   Prompts the user to print the page.

The [`print()`]{#dom-print .dfn dfn-for="Window" dfn-type="method"}
method steps are:

1.  Let `document`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#printing:this
    x-internal="this"}\'s [associated
    `Document`](nav-history-apis.html#concept-document-window){#printing:concept-document-window}.

2.  If `document`{.variable} is not [fully
    active](document-sequences.html#fully-active){#printing:fully-active},
    then return.

3.  If `document`{.variable}\'s [unload
    counter](document-lifecycle.html#unload-counter){#printing:unload-counter}
    is greater than 0, then return.

4.  If `document`{.variable} is [ready for post-load
    tasks](parsing.html#ready-for-post-load-tasks){#printing:ready-for-post-load-tasks},
    then run the [printing
    steps](#printing-steps){#printing:printing-steps} for
    `document`{.variable}.

5.  Otherwise, set `document`{.variable}\'s [print when
    loaded]{#print-when-loaded .dfn} flag.

User agents should also run the [printing
steps](#printing-steps){#printing:printing-steps-2} whenever the user
asks for the opportunity to [obtain a physical
form](rendering.html#obtain-a-physical-form){#printing:obtain-a-physical-form}
(e.g. printed copy), or the representation of a physical form (e.g. PDF
copy), of a document.

The [printing steps]{#printing-steps .dfn} for a
[`Document`{#printing:document}](dom.html#document)
`document`{.variable} are:

1.  The user agent may display a message to the user or return (or
    both).

    For instance, a kiosk browser could silently ignore any invocations
    of the [`print()`{#printing:dom-print}](#dom-print) method.

    For instance, a browser on a mobile device could detect that there
    are no printers in the vicinity and display a message saying so
    before continuing to offer a \"save to PDF\" option.

2.  If the [active sandboxing flag
    set](browsers.html#active-sandboxing-flag-set){#printing:active-sandboxing-flag-set}
    of `document`{.variable} has the [sandboxed modals
    flag](browsers.html#sandboxed-modals-flag){#printing:sandboxed-modals-flag}
    set, then return.

    If the printing dialog is blocked by a
    [`Document`{#printing:document-2}](dom.html#document)\'s sandbox,
    then neither the
    [`beforeprint`{#printing:event-beforeprint}](indices.html#event-beforeprint)
    nor
    [`afterprint`{#printing:event-afterprint}](indices.html#event-afterprint)
    events will be fired.

3.  The user agent must [fire an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#printing:concept-event-fire
    x-internal="concept-event-fire"} named
    [`beforeprint`{#printing:event-beforeprint-2}](indices.html#event-beforeprint)
    at the [relevant global
    object](webappapis.html#concept-relevant-global){#printing:concept-relevant-global}
    of `document`{.variable}, as well as any [child
    navigable](document-sequences.html#child-navigable){#printing:child-navigable}
    in it.

    Firing in children only doesn\'t seem right here, and some tasks
    likely need to be queued. See [issue
    #5096](https://github.com/whatwg/html/issues/5096).

    The
    [`beforeprint`{#printing:event-beforeprint-3}](indices.html#event-beforeprint)
    event can be used to annotate the printed copy, for instance adding
    the time at which the document was printed.

4.  The user agent should offer the user the opportunity to [obtain a
    physical
    form](rendering.html#obtain-a-physical-form){#printing:obtain-a-physical-form-2}
    (or the representation of a physical form) of `document`{.variable}.
    The user agent may wait for the user to either accept or decline
    before returning; if so, the user agent must
    [pause](webappapis.html#pause){#printing:pause} while the method is
    waiting. Even if the user agent doesn\'t wait at this point, the
    user agent must use the state of the relevant documents as they are
    at this point in the algorithm if and when it eventually creates the
    alternate form.

5.  The user agent must [fire an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#printing:concept-event-fire-2
    x-internal="concept-event-fire"} named
    [`afterprint`{#printing:event-afterprint-2}](indices.html#event-afterprint)
    at the [relevant global
    object](webappapis.html#concept-relevant-global){#printing:concept-relevant-global-2}
    of `document`{.variable}, as well as any [child
    navigables](document-sequences.html#child-navigable){#printing:child-navigable-2}
    in it.

    Firing in children only doesn\'t seem right here, and some tasks
    likely need to be queued. See [issue
    #5096](https://github.com/whatwg/html/issues/5096).

    The
    [`afterprint`{#printing:event-afterprint-3}](indices.html#event-afterprint)
    event can be used to revert annotations added in the earlier event,
    as well as showing post-printing UI. For instance, if a page is
    walking the user through the steps of applying for a home loan, the
    script could automatically advance to the next step after having
    printed a form or other.

[← 8.4 Dynamic markup insertion](dynamic-markup-insertion.html) ---
[Table of Contents](./) --- [8.9 System state and capabilities
→](system-state.html)
