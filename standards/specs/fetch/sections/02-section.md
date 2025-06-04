## [2. ]{.secno}[Infrastructure]{.content}[](#infrastructure){.self-link} {#infrastructure .heading .settled level="2"}

This specification depends on the Infra Standard.
[\[INFRA\]](#biblio-infra "Infra Standard"){link-type="biblio"}

This specification uses terminology from ABNF, Encoding, HTML, HTTP,
MIME Sniffing, Streams, URL, Web IDL, and WebSockets.
[\[ABNF\]](#biblio-abnf "Augmented BNF for Syntax Specifications: ABNF"){link-type="biblio"}
[\[ENCODING\]](#biblio-encoding "Encoding Standard"){link-type="biblio"}
[\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}
[\[HTTP\]](#biblio-http "HTTP Semantics"){link-type="biblio"}
[\[MIMESNIFF\]](#biblio-mimesniff "MIME Sniffing Standard"){link-type="biblio"}
[\[STREAMS\]](#biblio-streams "Streams Standard"){link-type="biblio"}
[\[URL\]](#biblio-url "URL Standard"){link-type="biblio"}
[\[WEBIDL\]](#biblio-webidl "Web IDL Standard"){link-type="biblio"}
[\[WEBSOCKETS\]](#biblio-websockets "WebSockets Standard"){link-type="biblio"}

[ABNF]{#abnf .dfn .dfn-paneled dfn-type="dfn" noexport=""} means ABNF as
augmented by HTTP (in particular the addition of `#`) and RFC 7405.
[\[RFC7405\]](#biblio-rfc7405 "Case-Sensitive String Support in ABNF"){link-type="biblio"}

------------------------------------------------------------------------

[Credentials]{#credentials .dfn .dfn-paneled dfn-type="dfn" export=""}
are HTTP cookies, TLS client certificates, and [authentication
entries](#authentication-entry){#ref-for-authentication-entry
link-type="dfn"} (for HTTP authentication).
[\[COOKIES\]](#biblio-cookies "HTTP State Management Mechanism"){link-type="biblio"}
[\[TLS\]](#biblio-tls "The Transport Layer Security (TLS) Protocol Version 1.3"){link-type="biblio"}
[\[HTTP\]](#biblio-http "HTTP Semantics"){link-type="biblio"}

------------------------------------------------------------------------

A [fetch params]{#fetch-params .dfn .dfn-paneled dfn-type="dfn"
noexport=""} is a
[struct](https://infra.spec.whatwg.org/#struct){#ref-for-struct
link-type="dfn"} used as a bookkeeping detail by the
[fetch](#concept-fetch){#ref-for-concept-fetch link-type="dfn"}
algorithm. It has the following
[items](https://infra.spec.whatwg.org/#struct-item){#ref-for-struct-item
link-type="dfn"}:

[request]{#fetch-params-request .dfn .dfn-paneled dfn-for="fetch params" dfn-type="dfn" noexport=""}
:   A [request](#concept-request){#ref-for-concept-request
    link-type="dfn"}.

[process request body chunk length]{#fetch-params-process-request-body .dfn .dfn-paneled dfn-for="fetch params" dfn-type="dfn" noexport=""} (default null)\
[process request end-of-body]{#fetch-params-process-request-end-of-body .dfn .dfn-paneled dfn-for="fetch params" dfn-type="dfn" noexport=""} (default null)\
[process early hints response]{#fetch-params-process-early-hints-response .dfn .dfn-paneled dfn-for="fetch params" dfn-type="dfn" noexport=""} (default null)\
[process response]{#fetch-params-process-response .dfn .dfn-paneled dfn-for="fetch params" dfn-type="dfn" noexport=""} (default null)\
[process response end-of-body]{#fetch-params-process-response-end-of-body .dfn .dfn-paneled dfn-for="fetch params" dfn-type="dfn" noexport=""} (default null)\
[process response consume body]{#fetch-params-process-response-consume-body .dfn .dfn-paneled dfn-for="fetch params" dfn-type="dfn" noexport=""} (default null)
:   Null or an algorithm.

[task destination]{#fetch-params-task-destination .dfn .dfn-paneled dfn-for="fetch params" dfn-type="dfn" noexport=""} (default null)
:   Null, a [global
    object](https://html.spec.whatwg.org/multipage/webappapis.html#global-object){#ref-for-global-object
    link-type="dfn"}, or a [parallel
    queue](https://html.spec.whatwg.org/multipage/infrastructure.html#parallel-queue){#ref-for-parallel-queue
    link-type="dfn"}.

[cross-origin isolated capability]{#fetch-params-cross-origin-isolated-capability .dfn .dfn-paneled dfn-for="fetch params" dfn-type="dfn" noexport=""} (default false)
:   A boolean.

[controller]{#fetch-params-controller .dfn .dfn-paneled dfn-for="fetch params" dfn-type="dfn" noexport=""} (default a new [fetch controller](#fetch-controller){#ref-for-fetch-controller link-type="dfn"})
:   A [fetch controller](#fetch-controller){#ref-for-fetch-controller①
    link-type="dfn"}.

[timing info]{#fetch-params-timing-info .dfn .dfn-paneled dfn-for="fetch params" dfn-type="dfn" noexport=""}
:   A [fetch timing info](#fetch-timing-info){#ref-for-fetch-timing-info
    link-type="dfn"}.

[preloaded response candidate]{#fetch-params-preloaded-response-candidate .dfn .dfn-paneled dfn-for="fetch params" dfn-type="dfn" export=""} (default null)
:   Null, \"`pending`\", or a
    [response](#concept-response){#ref-for-concept-response
    link-type="dfn"}.

A [fetch controller]{#fetch-controller .dfn .dfn-paneled dfn-type="dfn"
export=""} is a
[struct](https://infra.spec.whatwg.org/#struct){#ref-for-struct①
link-type="dfn"} used to enable callers of
[fetch](#concept-fetch){#ref-for-concept-fetch① link-type="dfn"} to
perform certain operations on it after it has started. It has the
following
[items](https://infra.spec.whatwg.org/#struct-item){#ref-for-struct-item①
link-type="dfn"}:

[state]{#fetch-controller-state .dfn .dfn-paneled dfn-for="fetch controller" dfn-type="dfn" export=""} (default \"`ongoing`\")
:   \"`ongoing`\", \"`terminated`\", or \"`aborted`\"

[full timing info]{#fetch-controller-full-timing-info .dfn .dfn-paneled dfn-for="fetch controller" dfn-type="dfn" noexport=""} (default null)
:   Null or a [fetch timing
    info](#fetch-timing-info){#ref-for-fetch-timing-info①
    link-type="dfn"}.

[report timing steps]{#fetch-controller-report-timing-steps .dfn .dfn-paneled dfn-for="fetch controller" dfn-type="dfn" noexport=""} (default null)
:   Null or an algorithm accepting a [global
    object](https://html.spec.whatwg.org/multipage/webappapis.html#global-object){#ref-for-global-object①
    link-type="dfn"}.

[serialized abort reason]{#fetch-controller-serialized-abort-reason .dfn .dfn-paneled dfn-for="fetch controller" dfn-type="dfn" export=""} (default null)
:   Null or a
    [Record](https://tc39.es/ecma262/#sec-list-and-record-specification-type){#ref-for-sec-list-and-record-specification-type
    link-type="dfn"} (result of
    [StructuredSerialize](https://html.spec.whatwg.org/multipage/structured-data.html#structuredserialize){#ref-for-structuredserialize
    link-type="abstract-op"}).

[next manual redirect steps]{#fetch-controller-next-manual-redirect-steps .dfn .dfn-paneled dfn-for="fetch controller" dfn-type="dfn" noexport=""} (default null)
:   Null or an algorithm accepting nothing.

::: {.algorithm algorithm="report timing" algorithm-for="fetch controller"}
To [report timing]{#finalize-and-report-timing .dfn .dfn-paneled
dfn-for="fetch controller" dfn-type="dfn" export=""} for a [fetch
controller](#fetch-controller){#ref-for-fetch-controller②
link-type="dfn"} `controller`{.variable} given a [global
object](https://html.spec.whatwg.org/multipage/webappapis.html#global-object){#ref-for-global-object②
link-type="dfn"} `global`{.variable}:

1.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert
    link-type="dfn"}: `controller`{.variable}'s [report timing
    steps](#fetch-controller-report-timing-steps){#ref-for-fetch-controller-report-timing-steps
    link-type="dfn"} is non-null.

2.  Call `controller`{.variable}'s [report timing
    steps](#fetch-controller-report-timing-steps){#ref-for-fetch-controller-report-timing-steps①
    link-type="dfn"} with `global`{.variable}.
:::

::: {.algorithm algorithm="process the next manual redirect" algorithm-for="fetch controller"}
To [process the next manual
redirect]{#fetch-controller-process-the-next-manual-redirect .dfn
.dfn-paneled dfn-for="fetch controller" dfn-type="dfn" export=""} for a
[fetch controller](#fetch-controller){#ref-for-fetch-controller③
link-type="dfn"} `controller`{.variable}:

1.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert①
    link-type="dfn"}: `controller`{.variable}'s [next manual redirect
    steps](#fetch-controller-next-manual-redirect-steps){#ref-for-fetch-controller-next-manual-redirect-steps
    link-type="dfn"} is non-null.

2.  Call `controller`{.variable}'s [next manual redirect
    steps](#fetch-controller-next-manual-redirect-steps){#ref-for-fetch-controller-next-manual-redirect-steps①
    link-type="dfn"}.
:::

::: {.algorithm algorithm="extract full timing info" algorithm-for="fetch controller"}
To [extract full timing info]{#extract-full-timing-info .dfn
.dfn-paneled dfn-for="fetch controller" dfn-type="dfn" export=""} given
a [fetch controller](#fetch-controller){#ref-for-fetch-controller④
link-type="dfn"} `controller`{.variable}:

1.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert②
    link-type="dfn"}: `controller`{.variable}'s [full timing
    info](#fetch-controller-full-timing-info){#ref-for-fetch-controller-full-timing-info
    link-type="dfn"} is non-null.

2.  Return `controller`{.variable}'s [full timing
    info](#fetch-controller-full-timing-info){#ref-for-fetch-controller-full-timing-info①
    link-type="dfn"}.
:::

::: {.algorithm algorithm="abort" algorithm-for="fetch controller"}
To [abort]{#fetch-controller-abort .dfn .dfn-paneled
dfn-for="fetch controller" dfn-type="dfn" export=""} a [fetch
controller](#fetch-controller){#ref-for-fetch-controller⑤
link-type="dfn"} `controller`{.variable} with an optional
`error`{.variable}:

1.  Set `controller`{.variable}'s
    [state](#fetch-controller-state){#ref-for-fetch-controller-state
    link-type="dfn"} to \"`aborted`\".

2.  Let `fallbackError`{.variable} be an
    \"[`AbortError`{.idl}](https://webidl.spec.whatwg.org/#aborterror){#ref-for-aborterror
    link-type="idl"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException
    link-type="idl"}.

3.  Set `error`{.variable} to `fallbackError`{.variable} if it is not
    given.

4.  Let `serializedError`{.variable} be
    [StructuredSerialize](https://html.spec.whatwg.org/multipage/structured-data.html#structuredserialize){#ref-for-structuredserialize①
    link-type="abstract-op"}(`error`{.variable}). If that threw an
    exception, catch it, and let `serializedError`{.variable} be
    [StructuredSerialize](https://html.spec.whatwg.org/multipage/structured-data.html#structuredserialize){#ref-for-structuredserialize②
    link-type="abstract-op"}(`fallbackError`{.variable}).

5.  Set `controller`{.variable}'s [serialized abort
    reason](#fetch-controller-serialized-abort-reason){#ref-for-fetch-controller-serialized-abort-reason
    link-type="dfn"} to `serializedError`{.variable}.
:::

::: {.algorithm algorithm="deserialize a serialized abort reason"}
To [deserialize a serialized abort
reason]{#deserialize-a-serialized-abort-reason .dfn .dfn-paneled
dfn-type="dfn" export=""}, given null or a
[Record](https://tc39.es/ecma262/#sec-list-and-record-specification-type){#ref-for-sec-list-and-record-specification-type①
link-type="dfn"} `abortReason`{.variable} and a
[realm](https://tc39.es/ecma262/#realm){#ref-for-realm link-type="dfn"}
`realm`{.variable}:

1.  Let `fallbackError`{.variable} be an
    \"[`AbortError`{.idl}](https://webidl.spec.whatwg.org/#aborterror){#ref-for-aborterror①
    link-type="idl"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException①
    link-type="idl"}.

2.  Let `deserializedError`{.variable} be `fallbackError`{.variable}.

3.  If `abortReason`{.variable} is non-null, then set
    `deserializedError`{.variable} to
    [StructuredDeserialize](https://html.spec.whatwg.org/multipage/structured-data.html#structureddeserialize){#ref-for-structureddeserialize
    link-type="abstract-op"}(`abortReason`{.variable},
    `realm`{.variable}). If that threw an exception or returned
    undefined, then set `deserializedError`{.variable} to
    `fallbackError`{.variable}.

4.  Return `deserializedError`{.variable}.
:::

::: {.algorithm algorithm="terminate" algorithm-for="fetch controller"}
To [terminate]{#fetch-controller-terminate .dfn .dfn-paneled
dfn-for="fetch controller" dfn-type="dfn" export=""} a [fetch
controller](#fetch-controller){#ref-for-fetch-controller⑥
link-type="dfn"} `controller`{.variable}, set `controller`{.variable}'s
[state](#fetch-controller-state){#ref-for-fetch-controller-state①
link-type="dfn"} to \"`terminated`\".
:::

A [fetch params](#fetch-params){#ref-for-fetch-params link-type="dfn"}
`fetchParams`{.variable} is [aborted]{#fetch-params-aborted .dfn
.dfn-paneled dfn-for="fetch params" dfn-type="dfn" noexport=""} if its
[controller](#fetch-params-controller){#ref-for-fetch-params-controller
link-type="dfn"}'s
[state](#fetch-controller-state){#ref-for-fetch-controller-state②
link-type="dfn"} is \"`aborted`\".

A [fetch params](#fetch-params){#ref-for-fetch-params① link-type="dfn"}
`fetchParams`{.variable} is [canceled]{#fetch-params-canceled .dfn
.dfn-paneled dfn-for="fetch params" dfn-type="dfn" noexport=""} if its
[controller](#fetch-params-controller){#ref-for-fetch-params-controller①
link-type="dfn"}'s
[state](#fetch-controller-state){#ref-for-fetch-controller-state③
link-type="dfn"} is \"`aborted`\" or \"`terminated`\".

A [fetch timing info]{#fetch-timing-info .dfn .dfn-paneled
dfn-type="dfn" export=""} is a
[struct](https://infra.spec.whatwg.org/#struct){#ref-for-struct②
link-type="dfn"} used to maintain timing information needed by Resource
Timing and Navigation Timing. It has the following
[items](https://infra.spec.whatwg.org/#struct-item){#ref-for-struct-item②
link-type="dfn"}:
[\[RESOURCE-TIMING\]](#biblio-resource-timing "Resource Timing"){link-type="biblio"}
[\[NAVIGATION-TIMING\]](#biblio-navigation-timing "Navigation Timing"){link-type="biblio"}

[start time]{#fetch-timing-info-start-time .dfn .dfn-paneled dfn-for="fetch timing info" dfn-type="dfn" export=""} (default 0)\
[redirect start time]{#fetch-timing-info-redirect-start-time .dfn .dfn-paneled dfn-for="fetch timing info" dfn-type="dfn" export=""} (default 0)\
[redirect end time]{#fetch-timing-info-redirect-end-time .dfn .dfn-paneled dfn-for="fetch timing info" dfn-type="dfn" export=""} (default 0)\
[post-redirect start time]{#fetch-timing-info-post-redirect-start-time .dfn .dfn-paneled dfn-for="fetch timing info" dfn-type="dfn" export=""} (default 0)\
[final service worker start time]{#fetch-timing-info-final-service-worker-start-time .dfn .dfn-paneled dfn-for="fetch timing info" dfn-type="dfn" export=""} (default 0)\
[final network-request start time]{#fetch-timing-info-final-network-request-start-time .dfn .dfn-paneled dfn-for="fetch timing info" dfn-type="dfn" export=""} (default 0)\
[first interim network-response start time]{#fetch-timing-info-first-interim-network-response-start-time .dfn .dfn-paneled dfn-for="fetch timing info" dfn-type="dfn" export=""} (default 0)\
[final network-response start time]{#fetch-timing-info-final-network-response-start-time .dfn .dfn-paneled dfn-for="fetch timing info" dfn-type="dfn" export=""} (default 0)\
[end time]{#fetch-timing-info-end-time .dfn .dfn-paneled dfn-for="fetch timing info" dfn-type="dfn" export=""} (default 0)
:   A
    [`DOMHighResTimeStamp`{.idl}](https://w3c.github.io/hr-time/#dom-domhighrestimestamp){#ref-for-dom-domhighrestimestamp
    link-type="idl"}.

[final connection timing info]{#fetch-timing-info-final-connection-timing-info .dfn .dfn-paneled dfn-for="fetch timing info" dfn-type="dfn" export=""} (default null)
:   Null or a [connection timing
    info](#connection-timing-info){#ref-for-connection-timing-info
    link-type="dfn"}.

[server-timing headers]{#fetch-timing-info-server-timing-headers .dfn .dfn-paneled dfn-for="fetch timing info" dfn-type="dfn" export=""} (default « »)
:   A [list](https://infra.spec.whatwg.org/#list){#ref-for-list
    link-type="dfn"} of strings.

[render-blocking]{#fetch-timing-info-render-blocking .dfn .dfn-paneled dfn-for="fetch timing info" dfn-type="dfn" export=""} (default false)
:   A boolean.

A [response body info]{#response-body-info .dfn .dfn-paneled
dfn-type="dfn" export=""} is a
[struct](https://infra.spec.whatwg.org/#struct){#ref-for-struct③
link-type="dfn"} used to maintain information needed by Resource Timing
and Navigation Timing. It has the following
[items](https://infra.spec.whatwg.org/#struct-item){#ref-for-struct-item③
link-type="dfn"}:
[\[RESOURCE-TIMING\]](#biblio-resource-timing "Resource Timing"){link-type="biblio"}
[\[NAVIGATION-TIMING\]](#biblio-navigation-timing "Navigation Timing"){link-type="biblio"}

[encoded size]{#fetch-timing-info-encoded-body-size .dfn .dfn-paneled dfn-for="response body info" dfn-type="dfn" export=""} (default 0)\
[decoded size]{#fetch-timing-info-decoded-body-size .dfn .dfn-paneled dfn-for="response body info" dfn-type="dfn" export=""} (default 0)
:   A number.

[content type]{#response-body-info-content-type .dfn .dfn-paneled dfn-for="response body info" dfn-type="dfn" export=""} (default the empty string)
:   An [ASCII
    string](https://infra.spec.whatwg.org/#ascii-string){#ref-for-ascii-string
    link-type="dfn"}.

[content encoding]{#response-body-info-content-encoding .dfn .dfn-paneled dfn-for="response body info" dfn-type="dfn" export=""} (default the empty string)
:   An [ASCII
    string](https://infra.spec.whatwg.org/#ascii-string){#ref-for-ascii-string①
    link-type="dfn"}.

::: {.algorithm algorithm="create an opaque timing info"}
To [create an opaque timing info]{#create-an-opaque-timing-info .dfn
.dfn-paneled dfn-type="dfn" export=""
lt="create an opaque timing info|creating an opaque timing info"}, given
a [fetch timing info](#fetch-timing-info){#ref-for-fetch-timing-info②
link-type="dfn"} `timingInfo`{.variable}, return a new [fetch timing
info](#fetch-timing-info){#ref-for-fetch-timing-info③ link-type="dfn"}
whose [start
time](#fetch-timing-info-start-time){#ref-for-fetch-timing-info-start-time
link-type="dfn"} and [post-redirect start
time](#fetch-timing-info-post-redirect-start-time){#ref-for-fetch-timing-info-post-redirect-start-time
link-type="dfn"} are `timingInfo`{.variable}'s [start
time](#fetch-timing-info-start-time){#ref-for-fetch-timing-info-start-time①
link-type="dfn"}.
:::

::: {.algorithm algorithm="queue a fetch task"}
To [queue a fetch task]{#queue-a-fetch-task .dfn .dfn-paneled
dfn-type="dfn" noexport=""}, given an algorithm `algorithm`{.variable},
a [global
object](https://html.spec.whatwg.org/multipage/webappapis.html#global-object){#ref-for-global-object③
link-type="dfn"} or a [parallel
queue](https://html.spec.whatwg.org/multipage/infrastructure.html#parallel-queue){#ref-for-parallel-queue①
link-type="dfn"} `taskDestination`{.variable}, run these steps:

1.  If `taskDestination`{.variable} is a [parallel
    queue](https://html.spec.whatwg.org/multipage/infrastructure.html#parallel-queue){#ref-for-parallel-queue②
    link-type="dfn"}, then
    [enqueue](https://html.spec.whatwg.org/multipage/infrastructure.html#enqueue-the-following-steps){#ref-for-enqueue-the-following-steps
    link-type="dfn"} `algorithm`{.variable} to
    `taskDestination`{.variable}.

2.  Otherwise, [queue a global
    task](https://html.spec.whatwg.org/multipage/webappapis.html#queue-a-global-task){#ref-for-queue-a-global-task
    link-type="dfn"} on the [networking task
    source](https://html.spec.whatwg.org/multipage/webappapis.html#networking-task-source){#ref-for-networking-task-source
    link-type="dfn"} with `taskDestination`{.variable} and
    `algorithm`{.variable}.
:::

------------------------------------------------------------------------

To [serialize an integer]{#serialize-an-integer .dfn .dfn-paneled
dfn-type="dfn" noexport=""}, represent it as a string of the shortest
possible decimal number.

This will be replaced by a more descriptive algorithm in Infra. See
[infra/201](https://github.com/whatwg/infra/issues/201).

### [2.1. ]{.secno}[URL]{.content}[](#url){.self-link} {#url .heading .settled level="2.1"}

A [local scheme]{#local-scheme .dfn .dfn-paneled dfn-type="dfn"
export=""} is \"`about`\", \"`blob`\", or \"`data`\".

A [URL](https://url.spec.whatwg.org/#concept-url){#ref-for-concept-url
link-type="dfn"} [is local]{#is-local .dfn .dfn-paneled dfn-type="dfn"
export=""} if its
[scheme](https://url.spec.whatwg.org/#concept-url-scheme){#ref-for-concept-url-scheme
link-type="dfn"} is a [local
scheme](#local-scheme){#ref-for-local-scheme link-type="dfn"}.

This definition is also used by Referrer Policy.
[\[REFERRER\]](#biblio-referrer "Referrer Policy"){link-type="biblio"}

An [HTTP(S) scheme]{#http-scheme .dfn .dfn-paneled dfn-type="dfn"
export=""} is \"`http`\" or \"`https`\".

A [fetch scheme]{#fetch-scheme .dfn .dfn-paneled dfn-type="dfn"
export=""} is \"`about`\", \"`blob`\", \"`data`\", \"`file`\", or an
[HTTP(S) scheme](#http-scheme){#ref-for-http-scheme link-type="dfn"}.

[HTTP(S) scheme](#http-scheme){#ref-for-http-scheme① link-type="dfn"}
and [fetch scheme](#fetch-scheme){#ref-for-fetch-scheme link-type="dfn"}
are also used by HTML.
[\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}

### [2.2. ]{.secno}[HTTP]{.content}[](#http){.self-link} {#http .heading .settled level="2.2"}

While [fetching](#concept-fetch){#ref-for-concept-fetch②
link-type="dfn"} encompasses more than just HTTP, it borrows a number of
concepts from HTTP and applies these to resources obtained via other
means (e.g., `data` URLs).

An [HTTP tab or space]{#http-tab-or-space .dfn .dfn-paneled
dfn-type="dfn" export=""} is U+0009 TAB or U+0020 SPACE.

[HTTP whitespace]{#http-whitespace .dfn .dfn-paneled dfn-type="dfn"
export=""} is U+000A LF, U+000D CR, or an [HTTP tab or
space](#http-tab-or-space){#ref-for-http-tab-or-space link-type="dfn"}.

[HTTP whitespace](#http-whitespace){#ref-for-http-whitespace
link-type="dfn"} is only useful for specific constructs that are reused
outside the context of HTTP headers (e.g., [MIME
types](https://mimesniff.spec.whatwg.org/#mime-type){#ref-for-mime-type
link-type="dfn"}). For HTTP header values, using [HTTP tab or
space](#http-tab-or-space){#ref-for-http-tab-or-space① link-type="dfn"}
is preferred, and outside that context [ASCII
whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#ref-for-ascii-whitespace
link-type="dfn"} is preferred. Unlike [ASCII
whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#ref-for-ascii-whitespace①
link-type="dfn"} this excludes U+000C FF.

An [HTTP newline byte]{#http-newline-byte .dfn .dfn-paneled
dfn-type="dfn" export=""} is 0x0A (LF) or 0x0D (CR).

An [HTTP tab or space byte]{#http-tab-or-space-byte .dfn .dfn-paneled
dfn-type="dfn" export=""} is 0x09 (HT) or 0x20 (SP).

An [HTTP whitespace byte]{#http-whitespace-byte .dfn .dfn-paneled
dfn-type="dfn" export=""} is an [HTTP newline
byte](#http-newline-byte){#ref-for-http-newline-byte link-type="dfn"} or
[HTTP tab or space
byte](#http-tab-or-space-byte){#ref-for-http-tab-or-space-byte
link-type="dfn"}.

:::: {.algorithm algorithm="collect an HTTP quoted string"}
To [collect an HTTP quoted string]{#collect-an-http-quoted-string .dfn
.dfn-paneled dfn-type="dfn" export=""
lt="collect an HTTP quoted string|collecting an HTTP quoted string"}
from a [string](https://infra.spec.whatwg.org/#string){#ref-for-string
link-type="dfn"} `input`{.variable}, given a [position
variable](https://infra.spec.whatwg.org/#string-position-variable){#ref-for-string-position-variable
link-type="dfn"} `position`{.variable} and an optional boolean
`extract-value`{.variable} (default false):

1.  Let `positionStart`{.variable} be `position`{.variable}.

2.  Let `value`{.variable} be the empty string.

3.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert③
    link-type="dfn"}: the [code
    point](https://infra.spec.whatwg.org/#code-point){#ref-for-code-point
    link-type="dfn"} at `position`{.variable} within `input`{.variable}
    is U+0022 (\").

4.  Advance `position`{.variable} by 1.

5.  While true:

    1.  Append the result of [collecting a sequence of code
        points](https://infra.spec.whatwg.org/#collect-a-sequence-of-code-points){#ref-for-collect-a-sequence-of-code-points
        link-type="dfn"} that are not U+0022 (\") or U+005C (\\) from
        `input`{.variable}, given `position`{.variable}, to
        `value`{.variable}.

    2.  If `position`{.variable} is past the end of `input`{.variable},
        then
        [break](https://infra.spec.whatwg.org/#iteration-break){#ref-for-iteration-break
        link-type="dfn"}.

    3.  Let `quoteOrBackslash`{.variable} be the [code
        point](https://infra.spec.whatwg.org/#code-point){#ref-for-code-point①
        link-type="dfn"} at `position`{.variable} within
        `input`{.variable}.

    4.  Advance `position`{.variable} by 1.

    5.  If `quoteOrBackslash`{.variable} is U+005C (\\), then:

        1.  If `position`{.variable} is past the end of
            `input`{.variable}, then append U+005C (\\) to
            `value`{.variable} and
            [break](https://infra.spec.whatwg.org/#iteration-break){#ref-for-iteration-break①
            link-type="dfn"}.

        2.  Append the [code
            point](https://infra.spec.whatwg.org/#code-point){#ref-for-code-point②
            link-type="dfn"} at `position`{.variable} within
            `input`{.variable} to `value`{.variable}.

        3.  Advance `position`{.variable} by 1.

    6.  Otherwise:

        1.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert④
            link-type="dfn"}: `quoteOrBackslash`{.variable} is U+0022
            (\").

        2.  [Break](https://infra.spec.whatwg.org/#iteration-break){#ref-for-iteration-break②
            link-type="dfn"}.

6.  If `extract-value`{.variable} is true, then return
    `value`{.variable}.

7.  Return the [code
    points](https://infra.spec.whatwg.org/#code-point){#ref-for-code-point③
    link-type="dfn"} from `positionStart`{.variable} to
    `position`{.variable}, inclusive, within `input`{.variable}.

::: {#example-http-quoted-string .example}
[](#example-http-quoted-string){.self-link}

Input

Output

Output with `extract-value`{.variable} set to true

Final [position
variable](https://infra.spec.whatwg.org/#string-position-variable){#ref-for-string-position-variable①
link-type="dfn"} value

\"[`"\`]{.mark}\"

\"`"\`\"

\"`\`\"

2

\"[`"Hello"`]{.mark}` World`\"

\"`"Hello"`\"

\"`Hello`\"

7

\"[`"Hello \\ World\""`]{.mark}\"

\"`"Hello \\ World\""`\"

\"`Hello \ World"`\"

18

[The [position
variable](https://infra.spec.whatwg.org/#string-position-variable){#ref-for-string-position-variable②
link-type="dfn"} always starts at 0 in these examples.]{.small}
:::
::::

#### [2.2.1. ]{.secno}[Methods]{.content}[](#methods){.self-link} {#methods .heading .settled level="2.2.1"}

A [method]{#concept-method .dfn .dfn-paneled dfn-type="dfn" export=""}
is a [byte
sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence
link-type="dfn"} that matches the
[method](https://httpwg.org/specs/rfc9110.html#method.overview){#ref-for-method.overview
link-type="dfn"} token production.

A [CORS-safelisted method]{#cors-safelisted-method .dfn .dfn-paneled
dfn-type="dfn" export=""} is a
[method](#concept-method){#ref-for-concept-method link-type="dfn"} that
is \``GET`\`, \``HEAD`\`, or \``POST`\`.

A [forbidden method]{#forbidden-method .dfn .dfn-paneled dfn-type="dfn"
export=""} is a [method](#concept-method){#ref-for-concept-method①
link-type="dfn"} that is a
[byte-case-insensitive](https://infra.spec.whatwg.org/#byte-case-insensitive){#ref-for-byte-case-insensitive
link-type="dfn"} match for \``CONNECT`\`, \``TRACE`\`, or \``TRACK`\`.
[\[HTTPVERBSEC1\]](#biblio-httpverbsec1 "Multiple vendors' web servers enable HTTP TRACE method by default."){link-type="biblio"},
[\[HTTPVERBSEC2\]](#biblio-httpverbsec2 "Microsoft Internet Information Server (IIS) vulnerable to cross-site scripting via HTTP TRACK method."){link-type="biblio"},
[\[HTTPVERBSEC3\]](#biblio-httpverbsec3 "HTTP proxy default configurations allow arbitrary TCP connections."){link-type="biblio"}

To [normalize]{#concept-method-normalize .dfn .dfn-paneled
dfn-for="method" dfn-type="dfn" export=""} a
[method](#concept-method){#ref-for-concept-method② link-type="dfn"}, if
it is a
[byte-case-insensitive](https://infra.spec.whatwg.org/#byte-case-insensitive){#ref-for-byte-case-insensitive①
link-type="dfn"} match for \``DELETE`\`, \``GET`\`, \``HEAD`\`,
\``OPTIONS`\`, \``POST`\`, or \``PUT`\`,
[byte-uppercase](https://infra.spec.whatwg.org/#byte-uppercase){#ref-for-byte-uppercase
link-type="dfn"} it.

[Normalization](#concept-method-normalize){#ref-for-concept-method-normalize
link-type="dfn"} is done for backwards compatibility and consistency
across APIs as [methods](#concept-method){#ref-for-concept-method③
link-type="dfn"} are actually \"case-sensitive\".

[](#example-normalization){.self-link}Using \``patch`\` is highly likely
to result in a \``405 Method Not Allowed`\`. \``PATCH`\` is much more
likely to succeed.

There are no restrictions on
[methods](#concept-method){#ref-for-concept-method④ link-type="dfn"}.
\``CHICKEN`\` is perfectly acceptable (and not a misspelling of
\``CHECKIN`\`). Other than those that are
[normalized](#concept-method-normalize){#ref-for-concept-method-normalize①
link-type="dfn"} there are no casing restrictions either. \``Egg`\` or
\``eGg`\` would be fine, though uppercase is encouraged for consistency.

#### [2.2.2. ]{.secno}[Headers]{.content}[](#terminology-headers){.self-link} {#terminology-headers .heading .settled level="2.2.2"}

HTTP generally refers to a header as a \"field\" or \"header field\".
The web platform uses the more colloquial term \"header\".
[\[HTTP\]](#biblio-http "HTTP Semantics"){link-type="biblio"}

A [header list]{#concept-header-list .dfn .dfn-paneled dfn-type="dfn"
export=""} is a
[list](https://infra.spec.whatwg.org/#list){#ref-for-list①
link-type="dfn"} of zero or more
[headers](#concept-header){#ref-for-concept-header link-type="dfn"}. It
is initially « ».

A [header list](#concept-header-list){#ref-for-concept-header-list
link-type="dfn"} is essentially a specialized multimap: an ordered list
of key-value pairs with potentially duplicate keys. Since headers other
than \``Set-Cookie`\` are always combined when exposed to client-side
JavaScript, implementations could choose a more efficient
representation, as long as they also support an associated data
structure for \``Set-Cookie`\` headers.

::: {.algorithm algorithm="get a structured field value" algorithm-for="header list"}
To [get a structured field
value]{#concept-header-list-get-structured-header .dfn .dfn-paneled
dfn-for="header list" dfn-type="dfn" export=""} given a [header
name](#header-name){#ref-for-header-name link-type="dfn"}
`name`{.variable} and a string `type`{.variable} from a [header
list](#concept-header-list){#ref-for-concept-header-list①
link-type="dfn"} `list`{.variable}, run these steps. They return null or
a [structured field
value](https://httpwg.org/specs/rfc9651.html#rfc.section.2){#ref-for-rfc.section.2
link-type="dfn"}.

1.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert⑤
    link-type="dfn"}: `type`{.variable} is one of \"`dictionary`\",
    \"`list`\", or \"`item`\".

2.  Let `value`{.variable} be the result of
    [getting](#concept-header-list-get){#ref-for-concept-header-list-get
    link-type="dfn"} `name`{.variable} from `list`{.variable}.

3.  If `value`{.variable} is null, then return null.

4.  Let `result`{.variable} be the result of [parsing structured
    fields](https://httpwg.org/specs/rfc9651.html#text-parse){#ref-for-text-parse
    link-type="dfn"} with `input_string`{.variable} set to
    `value`{.variable} and `header_type`{.variable} set to
    `type`{.variable}.

5.  If parsing failed, then return null.

6.  Return `result`{.variable}.

[Get a structured field
value](#concept-header-list-get-structured-header){#ref-for-concept-header-list-get-structured-header
link-type="dfn"} intentionally does not distinguish between a
[header](#concept-header){#ref-for-concept-header① link-type="dfn"} not
being present and its
[value](#concept-header-value){#ref-for-concept-header-value
link-type="dfn"} failing to parse as a [structured field
value](https://httpwg.org/specs/rfc9651.html#rfc.section.2){#ref-for-rfc.section.2①
link-type="dfn"}. This ensures uniform processing across the web
platform.
:::

::: {.algorithm algorithm="set a structured field value" algorithm-for="header list"}
To [set a structured field
value]{#concept-header-list-set-structured-header .dfn .dfn-paneled
dfn-for="header list" dfn-type="dfn" export=""} given a
[tuple](https://infra.spec.whatwg.org/#tuple){#ref-for-tuple
link-type="dfn"} ([header name](#header-name){#ref-for-header-name①
link-type="dfn"} `name`{.variable}, [structured field
value](https://httpwg.org/specs/rfc9651.html#rfc.section.2){#ref-for-rfc.section.2②
link-type="dfn"} `structuredValue`{.variable}), in a [header
list](#concept-header-list){#ref-for-concept-header-list②
link-type="dfn"} `list`{.variable}:

1.  Let `serializedValue`{.variable} be the result of executing the
    [serializing structured
    fields](https://httpwg.org/specs/rfc9651.html#text-serialize){#ref-for-text-serialize
    link-type="dfn"} algorithm on `structuredValue`{.variable}.

2.  [Set](#concept-header-list-set){#ref-for-concept-header-list-set
    link-type="dfn"} (`name`{.variable}, `serializedValue`{.variable})
    in `list`{.variable}.

[Structured field
values](https://httpwg.org/specs/rfc9651.html#rfc.section.2){#ref-for-rfc.section.2③
link-type="dfn"} are defined as objects which HTTP can (eventually)
serialize in interesting and efficient ways. For the moment, Fetch only
supports [header values](#header-value){#ref-for-header-value
link-type="dfn"} as [byte
sequences](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence①
link-type="dfn"}, which means that these objects can be set in [header
lists](#concept-header-list){#ref-for-concept-header-list③
link-type="dfn"} only via serialization, and they can be obtained from
[header lists](#concept-header-list){#ref-for-concept-header-list④
link-type="dfn"} only by parsing. In the future the fact that they are
objects might be preserved end-to-end.
[\[RFC9651\]](#biblio-rfc9651 "Structured Field Values for HTTP"){link-type="biblio"}
:::

------------------------------------------------------------------------

::: {.algorithm algorithm="contains" algorithm-for="header list"}
A [header list](#concept-header-list){#ref-for-concept-header-list⑤
link-type="dfn"} `list`{.variable} [contains]{#header-list-contains .dfn
.dfn-paneled dfn-for="header list" dfn-type="dfn" export=""
lt="contains|does not contain"} a [header
name](#header-name){#ref-for-header-name② link-type="dfn"}
`name`{.variable} if `list`{.variable}
[contains](https://infra.spec.whatwg.org/#list-contain){#ref-for-list-contain
link-type="dfn"} a [header](#concept-header){#ref-for-concept-header②
link-type="dfn"} whose
[name](#concept-header-name){#ref-for-concept-header-name
link-type="dfn"} is a
[byte-case-insensitive](https://infra.spec.whatwg.org/#byte-case-insensitive){#ref-for-byte-case-insensitive②
link-type="dfn"} match for `name`{.variable}.
:::

::: {.algorithm algorithm="get" algorithm-for="header list"}
To [get]{#concept-header-list-get .dfn .dfn-paneled
dfn-for="header list" dfn-type="dfn" export=""} a [header
name](#header-name){#ref-for-header-name③ link-type="dfn"}
`name`{.variable} from a [header
list](#concept-header-list){#ref-for-concept-header-list⑥
link-type="dfn"} `list`{.variable}, run these steps. They return null or
a [header value](#header-value){#ref-for-header-value① link-type="dfn"}.

1.  If `list`{.variable} [does not
    contain](#header-list-contains){#ref-for-header-list-contains
    link-type="dfn"} `name`{.variable}, then return null.

2.  Return the
    [values](#concept-header-value){#ref-for-concept-header-value①
    link-type="dfn"} of all
    [headers](#concept-header){#ref-for-concept-header③ link-type="dfn"}
    in `list`{.variable} whose
    [name](#concept-header-name){#ref-for-concept-header-name①
    link-type="dfn"} is a
    [byte-case-insensitive](https://infra.spec.whatwg.org/#byte-case-insensitive){#ref-for-byte-case-insensitive③
    link-type="dfn"} match for `name`{.variable}, separated from each
    other by 0x2C 0x20, in order.
:::

::: {.algorithm algorithm="get, decode, and split" algorithm-for="header list"}
To [get, decode, and split]{#concept-header-list-get-decode-split .dfn
.dfn-paneled dfn-for="header list" dfn-type="dfn" export=""
lt="get, decode, and split|getting, decoding, and splitting"} a [header
name](#header-name){#ref-for-header-name④ link-type="dfn"}
`name`{.variable} from [header
list](#concept-header-list){#ref-for-concept-header-list⑦
link-type="dfn"} `list`{.variable}, run these steps. They return null or
a [list](https://infra.spec.whatwg.org/#list){#ref-for-list②
link-type="dfn"} of
[strings](https://infra.spec.whatwg.org/#string){#ref-for-string①
link-type="dfn"}.

1.  Let `value`{.variable} be the result of
    [getting](#concept-header-list-get){#ref-for-concept-header-list-get①
    link-type="dfn"} `name`{.variable} from `list`{.variable}.

2.  If `value`{.variable} is null, then return null.

3.  Return the result of [getting, decoding, and
    splitting](#header-value-get-decode-and-split){#ref-for-header-value-get-decode-and-split
    link-type="dfn"} `value`{.variable}.
:::

::: {#example-header-list-get-decode-split .example}
[](#example-header-list-get-decode-split){.self-link}

This is how [get, decode, and
split](#concept-header-list-get-decode-split){#ref-for-concept-header-list-get-decode-split
link-type="dfn"} functions in practice with \``A`\` as the
`name`{.variable} argument:

Headers (as on the network)

Output

``` {.lang-http .highlight}
A: nosniff,
```

« \"`nosniff`\", \"\" »

``` {.lang-http .highlight}
A: nosniff
B: sniff
A:
```

``` {.lang-http .highlight}
A:
B: sniff
```

« \"\" »

``` {.lang-http .highlight}
B: sniff
```

null

``` {.lang-http .highlight}
A: text/html;", x/x
```

« \"`text/html;", x/x`\" »

``` {.lang-http .highlight}
A: text/html;"
A: x/x
```

``` {.lang-http .highlight}
A: x/x;test="hi",y/y
```

« \"`x/x;test="hi"`\", \"`y/y`\" »

``` {.lang-http .highlight}
A: x/x;test="hi"
C: **bingo**
A: y/y
```

``` {.lang-http .highlight}
A: x / x,,,1
```

« \"`x / x`\", \"\", \"\", \"`1`\" »

``` {.lang-http .highlight}
A: x / x
A: ,
A: 1
```

``` {.lang-http .highlight}
A: "1,2", 3
```

« \"`"1,2"`\", \"`3`\" »

``` {.lang-http .highlight}
A: "1,2"
D: 4
A: 3
```
:::

::: {.algorithm algorithm="get, decode, and split" algorithm-for="header value"}
To [get, decode, and split]{#header-value-get-decode-and-split .dfn
.dfn-paneled dfn-for="header value" dfn-type="dfn"
lt="get, decode, and split|getting, decoding, and splitting"
noexport=""} a [header value](#header-value){#ref-for-header-value②
link-type="dfn"} `value`{.variable}, run these steps. They return a
[list](https://infra.spec.whatwg.org/#list){#ref-for-list③
link-type="dfn"} of
[strings](https://infra.spec.whatwg.org/#string){#ref-for-string②
link-type="dfn"}.

1.  Let `input`{.variable} be the result of [isomorphic
    decoding](https://infra.spec.whatwg.org/#isomorphic-decode){#ref-for-isomorphic-decode
    link-type="dfn"} `value`{.variable}.

2.  Let `position`{.variable} be a [position
    variable](https://infra.spec.whatwg.org/#string-position-variable){#ref-for-string-position-variable③
    link-type="dfn"} for `input`{.variable}, initially pointing at the
    start of `input`{.variable}.

3.  Let `values`{.variable} be a
    [list](https://infra.spec.whatwg.org/#list){#ref-for-list④
    link-type="dfn"} of
    [strings](https://infra.spec.whatwg.org/#string){#ref-for-string③
    link-type="dfn"}, initially « ».

4.  Let `temporaryValue`{.variable} be the empty string.

5.  While true:

    1.  Append the result of [collecting a sequence of code
        points](https://infra.spec.whatwg.org/#collect-a-sequence-of-code-points){#ref-for-collect-a-sequence-of-code-points①
        link-type="dfn"} that are not U+0022 (\") or U+002C (,) from
        `input`{.variable}, given `position`{.variable}, to
        `temporaryValue`{.variable}.

        The result might be the empty string.

    2.  If `position`{.variable} is not past the end of
        `input`{.variable} and the [code
        point](https://infra.spec.whatwg.org/#code-point){#ref-for-code-point④
        link-type="dfn"} at `position`{.variable} within
        `input`{.variable} is U+0022 (\"):

        1.  Append the result of [collecting an HTTP quoted
            string](#collect-an-http-quoted-string){#ref-for-collect-an-http-quoted-string
            link-type="dfn"} from `input`{.variable}, given
            `position`{.variable}, to `temporaryValue`{.variable}.

        2.  If `position`{.variable} is not past the end of
            `input`{.variable}, then
            [continue](https://infra.spec.whatwg.org/#iteration-continue){#ref-for-iteration-continue
            link-type="dfn"}.

    3.  Remove all [HTTP tab or
        space](#http-tab-or-space){#ref-for-http-tab-or-space②
        link-type="dfn"} from the start and end of
        `temporaryValue`{.variable}.

    4.  [Append](https://infra.spec.whatwg.org/#list-append){#ref-for-list-append
        link-type="dfn"} `temporaryValue`{.variable} to
        `values`{.variable}.

    5.  Set `temporaryValue`{.variable} to the empty string.

    6.  If `position`{.variable} is past the end of `input`{.variable},
        then return `values`{.variable}.

    7.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert⑥
        link-type="dfn"}: the [code
        point](https://infra.spec.whatwg.org/#code-point){#ref-for-code-point⑤
        link-type="dfn"} at `position`{.variable} within
        `input`{.variable} is U+002C (,).

    8.  Advance `position`{.variable} by 1.

Except for blessed call sites, the algorithm directly above is not to be
invoked directly. Use [get, decode, and
split](#concept-header-list-get-decode-split){#ref-for-concept-header-list-get-decode-split①
link-type="dfn"} instead.
:::

::: {.algorithm algorithm="append" algorithm-for="header list"}
To [append]{#concept-header-list-append .dfn .dfn-paneled
dfn-for="header list" dfn-type="dfn" export=""} a
[header](#concept-header){#ref-for-concept-header④ link-type="dfn"}
(`name`{.variable}, `value`{.variable}) to a [header
list](#concept-header-list){#ref-for-concept-header-list⑧
link-type="dfn"} `list`{.variable}:

1.  If `list`{.variable}
    [contains](#header-list-contains){#ref-for-header-list-contains①
    link-type="dfn"} `name`{.variable}, then set `name`{.variable} to
    the first such [header](#concept-header){#ref-for-concept-header⑤
    link-type="dfn"}'s
    [name](#concept-header-name){#ref-for-concept-header-name②
    link-type="dfn"}.

    This reuses the casing of the
    [name](#concept-header-name){#ref-for-concept-header-name③
    link-type="dfn"} of the
    [header](#concept-header){#ref-for-concept-header⑥ link-type="dfn"}
    already in `list`{.variable}, if any. If there are multiple matched
    [headers](#concept-header){#ref-for-concept-header⑦ link-type="dfn"}
    their [names](#concept-header-name){#ref-for-concept-header-name④
    link-type="dfn"} will all be identical.

2.  [Append](https://infra.spec.whatwg.org/#list-append){#ref-for-list-append①
    link-type="dfn"} (`name`{.variable}, `value`{.variable}) to
    `list`{.variable}.
:::

::: {.algorithm algorithm="delete" algorithm-for="header list"}
To [delete]{#concept-header-list-delete .dfn .dfn-paneled
dfn-for="header list" dfn-type="dfn" export=""} a [header
name](#header-name){#ref-for-header-name⑤ link-type="dfn"}
`name`{.variable} from a [header
list](#concept-header-list){#ref-for-concept-header-list⑨
link-type="dfn"} `list`{.variable},
[remove](https://infra.spec.whatwg.org/#list-remove){#ref-for-list-remove
link-type="dfn"} all [headers](#concept-header){#ref-for-concept-header⑧
link-type="dfn"} whose
[name](#concept-header-name){#ref-for-concept-header-name⑤
link-type="dfn"} is a
[byte-case-insensitive](https://infra.spec.whatwg.org/#byte-case-insensitive){#ref-for-byte-case-insensitive④
link-type="dfn"} match for `name`{.variable} from `list`{.variable}.
:::

::: {.algorithm algorithm="set" algorithm-for="header list"}
To [set]{#concept-header-list-set .dfn .dfn-paneled
dfn-for="header list" dfn-type="dfn" export=""} a
[header](#concept-header){#ref-for-concept-header⑨ link-type="dfn"}
(`name`{.variable}, `value`{.variable}) in a [header
list](#concept-header-list){#ref-for-concept-header-list①⓪
link-type="dfn"} `list`{.variable}:

1.  If `list`{.variable}
    [contains](#header-list-contains){#ref-for-header-list-contains②
    link-type="dfn"} `name`{.variable}, then set the
    [value](#concept-header-value){#ref-for-concept-header-value②
    link-type="dfn"} of the first such
    [header](#concept-header){#ref-for-concept-header①⓪ link-type="dfn"}
    to `value`{.variable} and
    [remove](https://infra.spec.whatwg.org/#list-remove){#ref-for-list-remove①
    link-type="dfn"} the others.

2.  Otherwise,
    [append](https://infra.spec.whatwg.org/#list-append){#ref-for-list-append②
    link-type="dfn"} (`name`{.variable}, `value`{.variable}) to
    `list`{.variable}.
:::

::: {.algorithm algorithm="combine" algorithm-for="header list"}
To [combine]{#concept-header-list-combine .dfn .dfn-paneled
dfn-for="header list" dfn-type="dfn" export=""} a
[header](#concept-header){#ref-for-concept-header①① link-type="dfn"}
(`name`{.variable}, `value`{.variable}) in a [header
list](#concept-header-list){#ref-for-concept-header-list①①
link-type="dfn"} `list`{.variable}:

1.  If `list`{.variable}
    [contains](#header-list-contains){#ref-for-header-list-contains③
    link-type="dfn"} `name`{.variable}, then set the
    [value](#concept-header-value){#ref-for-concept-header-value③
    link-type="dfn"} of the first such
    [header](#concept-header){#ref-for-concept-header①② link-type="dfn"}
    to its [value](#concept-header-value){#ref-for-concept-header-value④
    link-type="dfn"}, followed by 0x2C 0x20, followed by
    `value`{.variable}.

2.  Otherwise,
    [append](https://infra.spec.whatwg.org/#list-append){#ref-for-list-append③
    link-type="dfn"} (`name`{.variable}, `value`{.variable}) to
    `list`{.variable}.

[Combine](#concept-header-list-combine){#ref-for-concept-header-list-combine
link-type="dfn"} is used by
[`XMLHttpRequest`{.idl}](https://xhr.spec.whatwg.org/#xmlhttprequest){#ref-for-xmlhttprequest
link-type="idl"} and the [WebSocket protocol
handshake](https://websockets.spec.whatwg.org/#concept-websocket-establish){#ref-for-concept-websocket-establish
link-type="dfn"}.
:::

::: {.algorithm algorithm="convert header names to a sorted-lowercase set"}
To [convert header names to a sorted-lowercase
set]{#convert-header-names-to-a-sorted-lowercase-set .dfn .dfn-paneled
dfn-type="dfn" noexport=""}, given a
[list](https://infra.spec.whatwg.org/#list){#ref-for-list⑤
link-type="dfn"} of
[names](#concept-header-name){#ref-for-concept-header-name⑥
link-type="dfn"} `headerNames`{.variable}, run these steps. They return
an [ordered
set](https://infra.spec.whatwg.org/#ordered-set){#ref-for-ordered-set
link-type="dfn"} of [header names](#header-name){#ref-for-header-name⑥
link-type="dfn"}.

1.  Let `headerNamesSet`{.variable} be a new [ordered
    set](https://infra.spec.whatwg.org/#ordered-set){#ref-for-ordered-set①
    link-type="dfn"}.

2.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate
    link-type="dfn"} `name`{.variable} of `headerNames`{.variable},
    [append](https://infra.spec.whatwg.org/#set-append){#ref-for-set-append
    link-type="dfn"} the result of
    [byte-lowercasing](https://infra.spec.whatwg.org/#byte-lowercase){#ref-for-byte-lowercase
    link-type="dfn"} `name`{.variable} to `headerNamesSet`{.variable}.

3.  Return the result of
    [sorting](https://infra.spec.whatwg.org/#list-sort-in-ascending-order){#ref-for-list-sort-in-ascending-order
    link-type="dfn"} `headerNamesSet`{.variable} in ascending order with
    [byte less
    than](https://infra.spec.whatwg.org/#byte-less-than){#ref-for-byte-less-than
    link-type="dfn"}.
:::

::: {.algorithm algorithm="sort and combine" algorithm-for="header list"}
To [sort and combine]{#concept-header-list-sort-and-combine .dfn
.dfn-paneled dfn-for="header list" dfn-type="dfn" export=""} a [header
list](#concept-header-list){#ref-for-concept-header-list①②
link-type="dfn"} `list`{.variable}, run these steps. They return a
[header list](#concept-header-list){#ref-for-concept-header-list①③
link-type="dfn"}.

1.  Let `headers`{.variable} be a [header
    list](#concept-header-list){#ref-for-concept-header-list①④
    link-type="dfn"}.

2.  Let `names`{.variable} be the result of [convert header names to a
    sorted-lowercase
    set](#convert-header-names-to-a-sorted-lowercase-set){#ref-for-convert-header-names-to-a-sorted-lowercase-set
    link-type="dfn"} with all the
    [names](#concept-header-name){#ref-for-concept-header-name⑦
    link-type="dfn"} of the
    [headers](#concept-header){#ref-for-concept-header①③
    link-type="dfn"} in `list`{.variable}.

3.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate①
    link-type="dfn"} `name`{.variable} of `names`{.variable}:

    1.  If `name`{.variable} is \``set-cookie`\`, then:

        1.  Let `values`{.variable} be a list of all
            [values](#concept-header-value){#ref-for-concept-header-value⑤
            link-type="dfn"} of
            [headers](#concept-header){#ref-for-concept-header①④
            link-type="dfn"} in `list`{.variable} whose
            [name](#concept-header-name){#ref-for-concept-header-name⑧
            link-type="dfn"} is a
            [byte-case-insensitive](https://infra.spec.whatwg.org/#byte-case-insensitive){#ref-for-byte-case-insensitive⑤
            link-type="dfn"} match for `name`{.variable}, in order.

        2.  [For
            each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate②
            link-type="dfn"} `value`{.variable} of `values`{.variable}:

            1.  [Append](https://infra.spec.whatwg.org/#list-append){#ref-for-list-append④
                link-type="dfn"} (`name`{.variable}, `value`{.variable})
                to `headers`{.variable}.

    2.  Otherwise:

        1.  Let `value`{.variable} be the result of
            [getting](#concept-header-list-get){#ref-for-concept-header-list-get②
            link-type="dfn"} `name`{.variable} from `list`{.variable}.

        2.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert⑦
            link-type="dfn"}: `value`{.variable} is non-null.

        3.  [Append](https://infra.spec.whatwg.org/#list-append){#ref-for-list-append⑤
            link-type="dfn"} (`name`{.variable}, `value`{.variable}) to
            `headers`{.variable}.

4.  Return `headers`{.variable}.
:::

------------------------------------------------------------------------

A [header]{#concept-header .dfn .dfn-paneled dfn-type="dfn" export=""}
is a [tuple](https://infra.spec.whatwg.org/#tuple){#ref-for-tuple①
link-type="dfn"} that consists of a [name]{#concept-header-name .dfn
.dfn-paneled dfn-for="header" dfn-type="dfn" export=""} (a [header
name](#header-name){#ref-for-header-name⑦ link-type="dfn"}) and
[value]{#concept-header-value .dfn .dfn-paneled dfn-for="header"
dfn-type="dfn" export=""} (a [header
value](#header-value){#ref-for-header-value③ link-type="dfn"}).

A [header name]{#header-name .dfn .dfn-paneled dfn-type="dfn" export=""}
is a [byte
sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence②
link-type="dfn"} that matches the
[field-name](https://httpwg.org/specs/rfc9110.html#fields.names){#ref-for-fields.names
link-type="dfn"} token production.

A [header value]{#header-value .dfn .dfn-paneled dfn-type="dfn"
export=""} is a [byte
sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence③
link-type="dfn"} that matches the following conditions:

- Has no leading or trailing [HTTP tab or space
  bytes](#http-tab-or-space-byte){#ref-for-http-tab-or-space-byte①
  link-type="dfn"}.

- Contains no 0x00 (NUL) or [HTTP newline
  bytes](#http-newline-byte){#ref-for-http-newline-byte①
  link-type="dfn"}.

The definition of [header value](#header-value){#ref-for-header-value④
link-type="dfn"} is not defined in terms of the
[field-value](https://httpwg.org/specs/rfc9110.html#fields.values){#ref-for-fields.values
link-type="dfn"} token production as it is [not compatible with deployed
content](https://github.com/httpwg/http-core/issues/215 "field-value value space").

::: {.algorithm algorithm="normalize" algorithm-for="header value"}
To [normalize]{#concept-header-value-normalize .dfn .dfn-paneled
dfn-for="header value" dfn-type="dfn" export=""} a [byte
sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence④
link-type="dfn"} `potentialValue`{.variable}, remove any leading and
trailing [HTTP whitespace
bytes](#http-whitespace-byte){#ref-for-http-whitespace-byte
link-type="dfn"} from `potentialValue`{.variable}.
:::

------------------------------------------------------------------------

::: {.algorithm algorithm="CORS-safelisted request-header"}
To determine whether a
[header](#concept-header){#ref-for-concept-header①⑤ link-type="dfn"}
(`name`{.variable}, `value`{.variable}) is a [CORS-safelisted
request-header]{#cors-safelisted-request-header .dfn .dfn-paneled
dfn-type="dfn" export=""}, run these steps:

1.  If `value`{.variable}'s
    [length](https://infra.spec.whatwg.org/#byte-sequence-length){#ref-for-byte-sequence-length
    link-type="dfn"} is greater than 128, then return false.

2.  [Byte-lowercase](https://infra.spec.whatwg.org/#byte-lowercase){#ref-for-byte-lowercase①
    link-type="dfn"} `name`{.variable} and switch on the result:

    \``accept`\`

    :   If `value`{.variable} contains a [CORS-unsafe request-header
        byte](#cors-unsafe-request-header-byte){#ref-for-cors-unsafe-request-header-byte
        link-type="dfn"}, then return false.

    \``accept-language`\`\
    \``content-language`\`

    :   If `value`{.variable} contains a byte that is not in the range
        0x30 (0) to 0x39 (9), inclusive, is not in the range 0x41 (A) to
        0x5A (Z), inclusive, is not in the range 0x61 (a) to 0x7A (z),
        inclusive, and is not 0x20 (SP), 0x2A (\*), 0x2C (,), 0x2D (-),
        0x2E (.), 0x3B (;), or 0x3D (=), then return false.

    \``content-type`\`

    :   1.  If `value`{.variable} contains a [CORS-unsafe request-header
            byte](#cors-unsafe-request-header-byte){#ref-for-cors-unsafe-request-header-byte①
            link-type="dfn"}, then return false.

        2.  Let `mimeType`{.variable} be the result of
            [parsing](https://mimesniff.spec.whatwg.org/#parse-a-mime-type){#ref-for-parse-a-mime-type
            link-type="dfn"} the result of [isomorphic
            decoding](https://infra.spec.whatwg.org/#isomorphic-decode){#ref-for-isomorphic-decode①
            link-type="dfn"} `value`{.variable}.

        3.  If `mimeType`{.variable} is failure, then return false.

        4.  If `mimeType`{.variable}'s
            [essence](https://mimesniff.spec.whatwg.org/#mime-type-essence){#ref-for-mime-type-essence
            link-type="dfn"} is not
            \"`application/x-www-form-urlencoded`\",
            \"`multipart/form-data`\", or \"`text/plain`\", then return
            false.

        This intentionally does not use [extract a MIME
        type](#concept-header-extract-mime-type){#ref-for-concept-header-extract-mime-type
        link-type="dfn"} as that algorithm is rather forgiving and
        servers are not expected to implement it.

        ::: {#example-cors-safelisted-request-header-content-type .example}
        [](#example-cors-safelisted-request-header-content-type){.self-link}
        If [extract a MIME
        type](#concept-header-extract-mime-type){#ref-for-concept-header-extract-mime-type①
        link-type="dfn"} were used the following request would not
        result in a CORS preflight and a naïve parser on the server
        might treat the request body as JSON:

        ``` {.lang-javascript .highlight}
        fetch("https://victim.example/naïve-endpoint", {
          method: "POST",
          headers: [
            ["Content-Type", "application/json"],
            ["Content-Type", "text/plain"]
          ],
          credentials: "include",
          body: JSON.stringify(exerciseForTheReader)
        });
        ```
        :::

    \``range`\`

    :   1.  Let `rangeValue`{.variable} be the result of [parsing a
            single range header
            value](#simple-range-header-value){#ref-for-simple-range-header-value
            link-type="dfn"} given `value`{.variable} and false.

        2.  If `rangeValue`{.variable} is failure, then return false.

        3.  If `rangeValue`{.variable}\[0\] is null, then return false.

            As web browsers have historically not emitted ranges such as
            \``bytes=-500`\` this algorithm does not safelist them.

    Otherwise

    :   Return false.

3.  Return true.

There are limited exceptions to the \``Content-Type`\` header safelist,
as documented in [CORS protocol exceptions](#cors-protocol-exceptions).
:::

::: {.algorithm algorithm="CORS-unsafe request-header byte"}
A [CORS-unsafe request-header byte]{#cors-unsafe-request-header-byte
.dfn .dfn-paneled dfn-type="dfn" noexport=""} is a byte
`byte`{.variable} for which one of the following is true:

- `byte`{.variable} is less than 0x20 and is not 0x09 HT

- `byte`{.variable} is 0x22 (\"), 0x28 (left parenthesis), 0x29 (right
  parenthesis), 0x3A (:), 0x3C (\<), 0x3E (\>), 0x3F (?), 0x40 (@), 0x5B
  (\[), 0x5C (\\), 0x5D (\]), 0x7B ({), 0x7D (}), or 0x7F DEL.
:::

::: {.algorithm algorithm="CORS-unsafe request-header names"}
The [CORS-unsafe request-header names]{#cors-unsafe-request-header-names
.dfn .dfn-paneled dfn-type="dfn" noexport=""}, given a [header
list](#concept-header-list){#ref-for-concept-header-list①⑤
link-type="dfn"} `headers`{.variable}, are determined as follows:

1.  Let `unsafeNames`{.variable} be a new
    [list](https://infra.spec.whatwg.org/#list){#ref-for-list⑥
    link-type="dfn"}.

2.  Let `potentiallyUnsafeNames`{.variable} be a new
    [list](https://infra.spec.whatwg.org/#list){#ref-for-list⑦
    link-type="dfn"}.

3.  Let `safelistValueSize`{.variable} be 0.

4.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate③
    link-type="dfn"} `header`{.variable} of `headers`{.variable}:

    1.  If `header`{.variable} is not a [CORS-safelisted
        request-header](#cors-safelisted-request-header){#ref-for-cors-safelisted-request-header
        link-type="dfn"}, then
        [append](https://infra.spec.whatwg.org/#list-append){#ref-for-list-append⑥
        link-type="dfn"} `header`{.variable}'s
        [name](#concept-header-name){#ref-for-concept-header-name⑨
        link-type="dfn"} to `unsafeNames`{.variable}.

    2.  Otherwise,
        [append](https://infra.spec.whatwg.org/#list-append){#ref-for-list-append⑦
        link-type="dfn"} `header`{.variable}'s
        [name](#concept-header-name){#ref-for-concept-header-name①⓪
        link-type="dfn"} to `potentiallyUnsafeNames`{.variable} and
        increase `safelistValueSize`{.variable} by `header`{.variable}'s
        [value](#concept-header-value){#ref-for-concept-header-value⑥
        link-type="dfn"}'s
        [length](https://infra.spec.whatwg.org/#byte-sequence-length){#ref-for-byte-sequence-length①
        link-type="dfn"}.

5.  If `safelistValueSize`{.variable} is greater than 1024, then [for
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate④
    link-type="dfn"} `name`{.variable} of
    `potentiallyUnsafeNames`{.variable},
    [append](https://infra.spec.whatwg.org/#list-append){#ref-for-list-append⑧
    link-type="dfn"} `name`{.variable} to `unsafeNames`{.variable}.

6.  Return the result of [convert header names to a sorted-lowercase
    set](#convert-header-names-to-a-sorted-lowercase-set){#ref-for-convert-header-names-to-a-sorted-lowercase-set①
    link-type="dfn"} with `unsafeNames`{.variable}.
:::

A [CORS non-wildcard request-header
name]{#cors-non-wildcard-request-header-name .dfn .dfn-paneled
dfn-type="dfn" export=""} is a [header
name](#header-name){#ref-for-header-name⑧ link-type="dfn"} that is a
[byte-case-insensitive](https://infra.spec.whatwg.org/#byte-case-insensitive){#ref-for-byte-case-insensitive⑥
link-type="dfn"} match for \``Authorization`\`.

A [privileged no-CORS request-header
name]{#privileged-no-cors-request-header-name .dfn .dfn-paneled
dfn-type="dfn" export=""} is a [header
name](#header-name){#ref-for-header-name⑨ link-type="dfn"} that is a
[byte-case-insensitive](https://infra.spec.whatwg.org/#byte-case-insensitive){#ref-for-byte-case-insensitive⑦
link-type="dfn"} match for one of

- \``Range`\`.

::: {.note role="note"}
These are headers that can be set by privileged APIs, and will be
preserved if their associated request object is copied, but will be
removed if the request is modified by unprivileged APIs.

\``Range`\` headers are commonly used by
[downloads](https://html.spec.whatwg.org/multipage/links.html#downloading-hyperlinks){#ref-for-downloading-hyperlinks
link-type="dfn"} and [media
fetches](https://html.spec.whatwg.org/multipage/media.html#concept-media-load-resource){#ref-for-concept-media-load-resource
link-type="dfn"}.

A helper is provided to [add a range
header](#concept-request-add-range-header){#ref-for-concept-request-add-range-header
link-type="dfn"} to a particular request.
:::

A [CORS-safelisted response-header
name]{#cors-safelisted-response-header-name .dfn .dfn-paneled
dfn-type="dfn" export=""}, given a
[list](https://infra.spec.whatwg.org/#list){#ref-for-list⑧
link-type="dfn"} of [header names](#header-name){#ref-for-header-name①⓪
link-type="dfn"} `list`{.variable}, is a [header
name](#header-name){#ref-for-header-name①① link-type="dfn"} that is a
[byte-case-insensitive](https://infra.spec.whatwg.org/#byte-case-insensitive){#ref-for-byte-case-insensitive⑧
link-type="dfn"} match for one of

- \``Cache-Control`\`
- \``Content-Language`\`
- \``Content-Length`\`
- \``Content-Type`\`
- \``Expires`\`
- \``Last-Modified`\`
- \``Pragma`\`
- Any
  [item](https://infra.spec.whatwg.org/#list-item){#ref-for-list-item
  link-type="dfn"} in `list`{.variable} that is not a [forbidden
  response-header
  name](#forbidden-response-header-name){#ref-for-forbidden-response-header-name
  link-type="dfn"}.

A [no-CORS-safelisted request-header
name]{#no-cors-safelisted-request-header-name .dfn .dfn-paneled
dfn-type="dfn" noexport=""} is a [header
name](#header-name){#ref-for-header-name①② link-type="dfn"} that is a
[byte-case-insensitive](https://infra.spec.whatwg.org/#byte-case-insensitive){#ref-for-byte-case-insensitive⑨
link-type="dfn"} match for one of

- \``Accept`\`
- \``Accept-Language`\`
- \``Content-Language`\`
- \``Content-Type`\`

::: {.algorithm algorithm="no-CORS-safelisted request-header"}
To determine whether a
[header](#concept-header){#ref-for-concept-header①⑥ link-type="dfn"}
(`name`{.variable}, `value`{.variable}) is a [no-CORS-safelisted
request-header]{#no-cors-safelisted-request-header .dfn .dfn-paneled
dfn-type="dfn" noexport=""}, run these steps:

1.  If `name`{.variable} is not a [no-CORS-safelisted request-header
    name](#no-cors-safelisted-request-header-name){#ref-for-no-cors-safelisted-request-header-name
    link-type="dfn"}, then return false.

2.  Return whether (`name`{.variable}, `value`{.variable}) is a
    [CORS-safelisted
    request-header](#cors-safelisted-request-header){#ref-for-cors-safelisted-request-header①
    link-type="dfn"}.
:::

:::: {.algorithm algorithm="forbidden request-header"}
A [header](#concept-header){#ref-for-concept-header①⑦ link-type="dfn"}
(`name`{.variable}, `value`{.variable}) is [forbidden
request-header]{#forbidden-request-header .dfn .dfn-paneled
dfn-type="dfn" export=""} if these steps return true:

1.  If `name`{.variable} is a
    [byte-case-insensitive](https://infra.spec.whatwg.org/#byte-case-insensitive){#ref-for-byte-case-insensitive①⓪
    link-type="dfn"} match for one of:

    - \``Accept-Charset`\`
    - \``Accept-Encoding`\`
    - \`[`Access-Control-Request-Headers`](#http-access-control-request-headers){#ref-for-http-access-control-request-headers
      link-type="http-header"}\`
    - \`[`Access-Control-Request-Method`](#http-access-control-request-method){#ref-for-http-access-control-request-method
      link-type="http-header"}\`
    - \``Connection`\`
    - \``Content-Length`\`
    - \``Cookie`\`
    - \``Cookie2`\`
    - \``Date`\`
    - \``DNT`\`
    - \``Expect`\`
    - \``Host`\`
    - \``Keep-Alive`\`
    - \`[`Origin`](#http-origin){#ref-for-http-origin①
      link-type="http-header"}\`
    - \``Referer`\`
    - \``Set-Cookie`\`
    - \``TE`\`
    - \``Trailer`\`
    - \``Transfer-Encoding`\`
    - \``Upgrade`\`
    - \``Via`\`

    then return true.

2.  If `name`{.variable} when
    [byte-lowercased](https://infra.spec.whatwg.org/#byte-lowercase){#ref-for-byte-lowercase②
    link-type="dfn"} [starts
    with](https://infra.spec.whatwg.org/#byte-sequence-starts-with){#ref-for-byte-sequence-starts-with
    link-type="dfn"} \``proxy-`\` or \``sec-`\`, then return true.

3.  If `name`{.variable} is a
    [byte-case-insensitive](https://infra.spec.whatwg.org/#byte-case-insensitive){#ref-for-byte-case-insensitive①①
    link-type="dfn"} match for one of:

    - \``X-HTTP-Method`\`
    - \``X-HTTP-Method-Override`\`
    - \``X-Method-Override`\`

    then:

    1.  Let `parsedValues`{.variable} be the result of [getting,
        decoding, and
        splitting](#header-value-get-decode-and-split){#ref-for-header-value-get-decode-and-split①
        link-type="dfn"} `value`{.variable}.

    2.  [For
        each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate⑤
        link-type="dfn"} `method`{.variable} of
        `parsedValues`{.variable}: if the [isomorphic
        encoding](https://infra.spec.whatwg.org/#isomorphic-encode){#ref-for-isomorphic-encode
        link-type="dfn"} of `method`{.variable} is a [forbidden
        method](#forbidden-method){#ref-for-forbidden-method
        link-type="dfn"}, then return true.

4.  Return false.

::: {.note role="note"}
These are forbidden so the user agent remains in full control over them.

[Header names](#header-name){#ref-for-header-name①③ link-type="dfn"}
starting with \``Sec-`\` are reserved to allow new
[headers](#concept-header){#ref-for-concept-header①⑧ link-type="dfn"} to
be minted that are safe from APIs using
[fetch](#concept-fetch){#ref-for-concept-fetch③ link-type="dfn"} that
allow control over [headers](#concept-header){#ref-for-concept-header①⑨
link-type="dfn"} by developers, such as
[`XMLHttpRequest`{.idl}](https://xhr.spec.whatwg.org/#xmlhttprequest){#ref-for-xmlhttprequest①
link-type="idl"}.
[\[XHR\]](#biblio-xhr "XMLHttpRequest Standard"){link-type="biblio"}

The \``Set-Cookie`\` header is semantically a response header, so it is
not useful on requests. Because \``Set-Cookie`\` headers cannot be
combined, they require more complex handling in the
[`Headers`{.idl}](#headers){#ref-for-headers link-type="idl"} object. It
is forbidden here to avoid leaking this complexity into requests.
:::
::::

A [forbidden response-header name]{#forbidden-response-header-name .dfn
.dfn-paneled dfn-type="dfn" export=""} is a [header
name](#header-name){#ref-for-header-name①④ link-type="dfn"} that is a
[byte-case-insensitive](https://infra.spec.whatwg.org/#byte-case-insensitive){#ref-for-byte-case-insensitive①②
link-type="dfn"} match for one of:

- \``Set-Cookie`\`
- \``Set-Cookie2`\`

A [request-body-header name]{#request-body-header-name .dfn .dfn-paneled
dfn-type="dfn" export=""} is a [header
name](#header-name){#ref-for-header-name①⑤ link-type="dfn"} that is a
[byte-case-insensitive](https://infra.spec.whatwg.org/#byte-case-insensitive){#ref-for-byte-case-insensitive①③
link-type="dfn"} match for one of:

- \``Content-Encoding`\`
- \``Content-Language`\`
- \``Content-Location`\`
- \``Content-Type`\`

------------------------------------------------------------------------

::: {.algorithm algorithm="extract header values"}
To [extract header values]{#extract-header-values .dfn .dfn-paneled
dfn-type="dfn" export=""
lt="extract header values|extracting header values"} given a
[header](#concept-header){#ref-for-concept-header②⓪ link-type="dfn"}
`header`{.variable}, run these steps:

1.  If parsing `header`{.variable}'s
    [value](#concept-header-value){#ref-for-concept-header-value⑦
    link-type="dfn"}, per the [ABNF](#abnf){#ref-for-abnf
    link-type="dfn"} for `header`{.variable}'s
    [name](#concept-header-name){#ref-for-concept-header-name①①
    link-type="dfn"}, fails, then return failure.

2.  Return one or more
    [values](#concept-header-value){#ref-for-concept-header-value⑧
    link-type="dfn"} resulting from parsing `header`{.variable}'s
    [value](#concept-header-value){#ref-for-concept-header-value⑨
    link-type="dfn"}, per the [ABNF](#abnf){#ref-for-abnf①
    link-type="dfn"} for `header`{.variable}'s
    [name](#concept-header-name){#ref-for-concept-header-name①②
    link-type="dfn"}.
:::

::: {.algorithm algorithm="extract header list values"}
To [extract header list values]{#extract-header-list-values .dfn
.dfn-paneled dfn-type="dfn" export=""
lt="extract header list values|extracting header list values"} given a
[header name](#header-name){#ref-for-header-name①⑥ link-type="dfn"}
`name`{.variable} and a [header
list](#concept-header-list){#ref-for-concept-header-list①⑥
link-type="dfn"} `list`{.variable}, run these steps:

1.  If `list`{.variable} [does not
    contain](#header-list-contains){#ref-for-header-list-contains④
    link-type="dfn"} `name`{.variable}, then return null.

2.  If the [ABNF](#abnf){#ref-for-abnf② link-type="dfn"} for
    `name`{.variable} allows a single
    [header](#concept-header){#ref-for-concept-header②① link-type="dfn"}
    and `list`{.variable}
    [contains](#header-list-contains){#ref-for-header-list-contains⑤
    link-type="dfn"} more than one, then return failure.

    If different error handling is needed, extract the desired
    [header](#concept-header){#ref-for-concept-header②② link-type="dfn"}
    first.

3.  Let `values`{.variable} be an empty
    [list](https://infra.spec.whatwg.org/#list){#ref-for-list⑨
    link-type="dfn"}.

4.  For each [header](#concept-header){#ref-for-concept-header②③
    link-type="dfn"} `header`{.variable} `list`{.variable}
    [contains](#header-list-contains){#ref-for-header-list-contains⑥
    link-type="dfn"} whose
    [name](#concept-header-name){#ref-for-concept-header-name①③
    link-type="dfn"} is `name`{.variable}:

    1.  Let `extract`{.variable} be the result of [extracting header
        values](#extract-header-values){#ref-for-extract-header-values
        link-type="dfn"} from `header`{.variable}.

    2.  If `extract`{.variable} is failure, then return failure.

    3.  Append each
        [value](#concept-header-value){#ref-for-concept-header-value①⓪
        link-type="dfn"} in `extract`{.variable}, in order, to
        `values`{.variable}.

5.  Return `values`{.variable}.
:::

::: {.algorithm algorithm="build a content range"}
To [build a content range]{#build-a-content-range .dfn .dfn-paneled
dfn-type="dfn" noexport=""} given an integer `rangeStart`{.variable}, an
integer `rangeEnd`{.variable}, and an integer `fullLength`{.variable},
run these steps:

1.  Let `contentRange`{.variable} be \``bytes `\`.

2.  Append `rangeStart`{.variable},
    [serialized](#serialize-an-integer){#ref-for-serialize-an-integer
    link-type="dfn"} and [isomorphic
    encoded](https://infra.spec.whatwg.org/#isomorphic-encode){#ref-for-isomorphic-encode①
    link-type="dfn"}, to `contentRange`{.variable}.

3.  Append 0x2D (-) to `contentRange`{.variable}.

4.  Append `rangeEnd`{.variable},
    [serialized](#serialize-an-integer){#ref-for-serialize-an-integer①
    link-type="dfn"} and [isomorphic
    encoded](https://infra.spec.whatwg.org/#isomorphic-encode){#ref-for-isomorphic-encode②
    link-type="dfn"} to `contentRange`{.variable}.

5.  Append 0x2F (/) to `contentRange`{.variable}.

6.  Append `fullLength`{.variable},
    [serialized](#serialize-an-integer){#ref-for-serialize-an-integer②
    link-type="dfn"} and [isomorphic
    encoded](https://infra.spec.whatwg.org/#isomorphic-encode){#ref-for-isomorphic-encode③
    link-type="dfn"} to `contentRange`{.variable}.

7.  Return `contentRange`{.variable}.
:::

::: {.algorithm algorithm="parse a single range header value"}
To [parse a single range header value]{#simple-range-header-value .dfn
.dfn-paneled dfn-type="dfn" noexport=""} from a [byte
sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence⑤
link-type="dfn"} `value`{.variable} and a boolean
`allowWhitespace`{.variable}, run these steps:

1.  Let `data`{.variable} be the [isomorphic
    decoding](https://infra.spec.whatwg.org/#isomorphic-decode){#ref-for-isomorphic-decode②
    link-type="dfn"} of `value`{.variable}.

2.  If `data`{.variable} does not [start
    with](https://infra.spec.whatwg.org/#string-starts-with){#ref-for-string-starts-with
    link-type="dfn"} \"`bytes`\", then return failure.

3.  Let `position`{.variable} be a [position
    variable](https://infra.spec.whatwg.org/#string-position-variable){#ref-for-string-position-variable④
    link-type="dfn"} for `data`{.variable}, initially pointing at the
    5th [code
    point](https://infra.spec.whatwg.org/#code-point){#ref-for-code-point⑥
    link-type="dfn"} of `data`{.variable}.

4.  If `allowWhitespace`{.variable} is true, [collect a sequence of code
    points](https://infra.spec.whatwg.org/#collect-a-sequence-of-code-points){#ref-for-collect-a-sequence-of-code-points②
    link-type="dfn"} that are [HTTP tab or
    space](#http-tab-or-space){#ref-for-http-tab-or-space③
    link-type="dfn"}, from `data`{.variable} given
    `position`{.variable}.

5.  If the [code
    point](https://infra.spec.whatwg.org/#code-point){#ref-for-code-point⑦
    link-type="dfn"} at `position`{.variable} within `data`{.variable}
    is not U+003D (=), then return failure.

6.  Advance `position`{.variable} by 1.

7.  If `allowWhitespace`{.variable} is true, [collect a sequence of code
    points](https://infra.spec.whatwg.org/#collect-a-sequence-of-code-points){#ref-for-collect-a-sequence-of-code-points③
    link-type="dfn"} that are [HTTP tab or
    space](#http-tab-or-space){#ref-for-http-tab-or-space④
    link-type="dfn"}, from `data`{.variable} given
    `position`{.variable}.

8.  Let `rangeStart`{.variable} be the result of [collecting a sequence
    of code
    points](https://infra.spec.whatwg.org/#collect-a-sequence-of-code-points){#ref-for-collect-a-sequence-of-code-points④
    link-type="dfn"} that are [ASCII
    digits](https://infra.spec.whatwg.org/#ascii-digit){#ref-for-ascii-digit
    link-type="dfn"}, from `data`{.variable} given
    `position`{.variable}.

9.  Let `rangeStartValue`{.variable} be `rangeStart`{.variable},
    interpreted as decimal number, if `rangeStart`{.variable} is not the
    empty string; otherwise null.

10. If `allowWhitespace`{.variable} is true, [collect a sequence of code
    points](https://infra.spec.whatwg.org/#collect-a-sequence-of-code-points){#ref-for-collect-a-sequence-of-code-points⑤
    link-type="dfn"} that are [HTTP tab or
    space](#http-tab-or-space){#ref-for-http-tab-or-space⑤
    link-type="dfn"}, from `data`{.variable} given
    `position`{.variable}.

11. If the [code
    point](https://infra.spec.whatwg.org/#code-point){#ref-for-code-point⑧
    link-type="dfn"} at `position`{.variable} within `data`{.variable}
    is not U+002D (-), then return failure.

12. Advance `position`{.variable} by 1.

13. If `allowWhitespace`{.variable} is true, [collect a sequence of code
    points](https://infra.spec.whatwg.org/#collect-a-sequence-of-code-points){#ref-for-collect-a-sequence-of-code-points⑥
    link-type="dfn"} that are [HTTP tab or
    space](#http-tab-or-space){#ref-for-http-tab-or-space⑥
    link-type="dfn"}, from `data`{.variable} given
    `position`{.variable}.

14. Let `rangeEnd`{.variable} be the result of [collecting a sequence of
    code
    points](https://infra.spec.whatwg.org/#collect-a-sequence-of-code-points){#ref-for-collect-a-sequence-of-code-points⑦
    link-type="dfn"} that are [ASCII
    digits](https://infra.spec.whatwg.org/#ascii-digit){#ref-for-ascii-digit①
    link-type="dfn"}, from `data`{.variable} given
    `position`{.variable}.

15. Let `rangeEndValue`{.variable} be `rangeEnd`{.variable}, interpreted
    as decimal number, if `rangeEnd`{.variable} is not the empty string;
    otherwise null.

16. If `position`{.variable} is not past the end of `data`{.variable},
    then return failure.

17. If `rangeEndValue`{.variable} and `rangeStartValue`{.variable} are
    null, then return failure.

18. If `rangeStartValue`{.variable} and `rangeEndValue`{.variable} are
    numbers, and `rangeStartValue`{.variable} is greater than
    `rangeEndValue`{.variable}, then return failure.

19. Return (`rangeStartValue`{.variable}, `rangeEndValue`{.variable}).

    The range end or start can be omitted, e.g., \``bytes=0-`\` or
    \``bytes=-500`\` are valid ranges.

[Parse a single range header
value](#simple-range-header-value){#ref-for-simple-range-header-value①
link-type="dfn"} succeeds for a subset of allowed range header values,
but it is the most common form used by user agents when requesting media
or resuming downloads. This format of range header value can be set
using [add a range
header](#concept-request-add-range-header){#ref-for-concept-request-add-range-header①
link-type="dfn"}.
:::

------------------------------------------------------------------------

A [default \``User-Agent`\` value]{#default-user-agent-value .dfn
.dfn-paneled dfn-type="dfn" export=""} is an
[implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#ref-for-implementation-defined
link-type="dfn"} [header value](#header-value){#ref-for-header-value⑤
link-type="dfn"} for the \``User-Agent`\`
[header](#concept-header){#ref-for-concept-header②④ link-type="dfn"}.

The [document \``Accept`\` header value]{#document-accept-header-value
.dfn .dfn-paneled dfn-type="dfn" noexport=""} is
\``text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8`\`.

#### [2.2.3. ]{.secno}[Statuses]{.content}[](#statuses){.self-link} {#statuses .heading .settled level="2.2.3"}

A [status]{#concept-status .dfn .dfn-paneled dfn-type="dfn" export=""}
is an integer in the range 0 to 999, inclusive.

Various edge cases in mapping HTTP/1's `status-code` to this concept are
worked on in [issue #1156](https://github.com/whatwg/fetch/issues/1156).

A [null body status]{#null-body-status .dfn .dfn-paneled dfn-type="dfn"
export=""} is a [status](#concept-status){#ref-for-concept-status
link-type="dfn"} that is 101, 103, 204, 205, or 304.

An [ok status]{#ok-status .dfn .dfn-paneled dfn-type="dfn" export=""} is
a [status](#concept-status){#ref-for-concept-status① link-type="dfn"} in
the range 200 to 299, inclusive.

A [redirect status]{#redirect-status .dfn .dfn-paneled dfn-type="dfn"
export=""} is a [status](#concept-status){#ref-for-concept-status②
link-type="dfn"} that is 301, 302, 303, 307, or 308.

#### [2.2.4. ]{.secno}[Bodies]{.content}[](#bodies){.self-link} {#bodies .heading .settled level="2.2.4"}

A [body]{#concept-body .dfn .dfn-paneled dfn-type="dfn" export=""}
consists of:

- A [stream]{#concept-body-stream .dfn .dfn-paneled dfn-for="body"
  dfn-type="dfn" export=""} (a
  [`ReadableStream`{.idl}](https://streams.spec.whatwg.org/#readablestream){#ref-for-readablestream
  link-type="idl"} object).

- A [source]{#concept-body-source .dfn .dfn-paneled dfn-for="body"
  dfn-type="dfn" export=""} (null, a [byte
  sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence⑥
  link-type="dfn"}, a
  [`Blob`{.idl}](https://w3c.github.io/FileAPI/#dfn-Blob){#ref-for-dfn-Blob
  link-type="idl"} object, or a
  [`FormData`{.idl}](https://xhr.spec.whatwg.org/#formdata){#ref-for-formdata
  link-type="idl"} object), initially null.

- A [length]{#concept-body-total-bytes .dfn .dfn-paneled dfn-for="body"
  dfn-type="dfn" export=""} (null or an integer), initially null.

::: {.algorithm algorithm="clone" algorithm-for="body"}
To [clone]{#concept-body-clone .dfn .dfn-paneled dfn-for="body"
dfn-type="dfn" export=""} a [body](#concept-body){#ref-for-concept-body
link-type="dfn"} `body`{.variable}, run these steps:

1.  Let « `out1`{.variable}, `out2`{.variable} » be the result of
    [teeing](https://streams.spec.whatwg.org/#readablestream-tee){#ref-for-readablestream-tee
    link-type="dfn"} `body`{.variable}'s
    [stream](#concept-body-stream){#ref-for-concept-body-stream
    link-type="dfn"}.

2.  Set `body`{.variable}'s
    [stream](#concept-body-stream){#ref-for-concept-body-stream①
    link-type="dfn"} to `out1`{.variable}.

3.  Return a [body](#concept-body){#ref-for-concept-body①
    link-type="dfn"} whose
    [stream](#concept-body-stream){#ref-for-concept-body-stream②
    link-type="dfn"} is `out2`{.variable} and other members are copied
    from `body`{.variable}.
:::

::: {.algorithm algorithm="as a body" algorithm-for="byte sequence"}
To get a [byte
sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence⑦
link-type="dfn"} `bytes`{.variable} [as a body]{#byte-sequence-as-a-body
.dfn .dfn-paneled dfn-for="byte sequence" dfn-type="dfn" export=""},
return the [body](#body-with-type-body){#ref-for-body-with-type-body
link-type="dfn"} of the result of [safely
extracting](#bodyinit-safely-extract){#ref-for-bodyinit-safely-extract
link-type="dfn"} `bytes`{.variable}.
:::

------------------------------------------------------------------------

::: {.algorithm algorithm="incrementally read" algorithm-for="body"}
To [incrementally read]{#body-incrementally-read .dfn .dfn-paneled
dfn-for="body" dfn-type="dfn" export=""} a
[body](#concept-body){#ref-for-concept-body② link-type="dfn"}
`body`{.variable}, given an algorithm `processBodyChunk`{.variable}, an
algorithm `processEndOfBody`{.variable}, an algorithm
`processBodyError`{.variable}, and an optional null, [parallel
queue](https://html.spec.whatwg.org/multipage/infrastructure.html#parallel-queue){#ref-for-parallel-queue③
link-type="dfn"}, or [global
object](https://html.spec.whatwg.org/multipage/webappapis.html#global-object){#ref-for-global-object④
link-type="dfn"} `taskDestination`{.variable} (default null), run these
steps. `processBodyChunk`{.variable} must be an algorithm accepting a
[byte
sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence⑧
link-type="dfn"}. `processEndOfBody`{.variable} must be an algorithm
accepting no arguments. `processBodyError`{.variable} must be an
algorithm accepting an exception.

1.  If `taskDestination`{.variable} is null, then set
    `taskDestination`{.variable} to the result of [starting a new
    parallel
    queue](https://html.spec.whatwg.org/multipage/infrastructure.html#starting-a-new-parallel-queue){#ref-for-starting-a-new-parallel-queue
    link-type="dfn"}.

2.  Let `reader`{.variable} be the result of [getting a
    reader](https://streams.spec.whatwg.org/#readablestream-get-a-reader){#ref-for-readablestream-get-a-reader
    link-type="dfn"} for `body`{.variable}'s
    [stream](#concept-body-stream){#ref-for-concept-body-stream③
    link-type="dfn"}.

    This operation will not throw an exception.

3.  Perform the [incrementally-read
    loop](#incrementally-read-loop){#ref-for-incrementally-read-loop
    link-type="dfn"} given `reader`{.variable},
    `taskDestination`{.variable}, `processBodyChunk`{.variable},
    `processEndOfBody`{.variable}, and `processBodyError`{.variable}.
:::

::: {.algorithm algorithm="incrementally-read loop"}
To perform the [incrementally-read loop]{#incrementally-read-loop .dfn
.dfn-paneled dfn-type="dfn" noexport=""}, given a
[`ReadableStreamDefaultReader`{.idl}](https://streams.spec.whatwg.org/#readablestreamdefaultreader){#ref-for-readablestreamdefaultreader
link-type="idl"} object `reader`{.variable}, [parallel
queue](https://html.spec.whatwg.org/multipage/infrastructure.html#parallel-queue){#ref-for-parallel-queue④
link-type="dfn"} or [global
object](https://html.spec.whatwg.org/multipage/webappapis.html#global-object){#ref-for-global-object⑤
link-type="dfn"} `taskDestination`{.variable}, algorithm
`processBodyChunk`{.variable}, algorithm `processEndOfBody`{.variable},
and algorithm `processBodyError`{.variable}:

1.  Let `readRequest`{.variable} be the following [read
    request](https://streams.spec.whatwg.org/#read-request){#ref-for-read-request
    link-type="dfn"}:

    [chunk steps](https://streams.spec.whatwg.org/#read-request-chunk-steps){#ref-for-read-request-chunk-steps link-type="dfn"}, given `chunk`{.variable}

    :   1.  Let `continueAlgorithm`{.variable} be null.

        2.  If `chunk`{.variable} is not a
            [`Uint8Array`{.idl}](https://webidl.spec.whatwg.org/#idl-Uint8Array){#ref-for-idl-Uint8Array
            link-type="idl"} object, then set
            `continueAlgorithm`{.variable} to this step: run
            `processBodyError`{.variable} given a
            [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror
            link-type="idl"}.

        3.  Otherwise:

            1.  Let `bytes`{.variable} be a [copy
                of](https://webidl.spec.whatwg.org/#dfn-get-buffer-source-copy){#ref-for-dfn-get-buffer-source-copy
                link-type="dfn"} `chunk`{.variable}.

                Implementations are strongly encouraged to use an
                implementation strategy that avoids this copy where
                possible.

            2.  Set `continueAlgorithm`{.variable} to these steps:

                1.  Run `processBodyChunk`{.variable} given
                    `bytes`{.variable}.

                2.  Perform the [incrementally-read
                    loop](#incrementally-read-loop){#ref-for-incrementally-read-loop①
                    link-type="dfn"} given `reader`{.variable},
                    `taskDestination`{.variable},
                    `processBodyChunk`{.variable},
                    `processEndOfBody`{.variable}, and
                    `processBodyError`{.variable}.

        4.  [Queue a fetch
            task](#queue-a-fetch-task){#ref-for-queue-a-fetch-task
            link-type="dfn"} given `continueAlgorithm`{.variable} and
            `taskDestination`{.variable}.

    [close steps](https://streams.spec.whatwg.org/#read-request-close-steps){#ref-for-read-request-close-steps link-type="dfn"}

    :   1.  [Queue a fetch
            task](#queue-a-fetch-task){#ref-for-queue-a-fetch-task①
            link-type="dfn"} given `processEndOfBody`{.variable} and
            `taskDestination`{.variable}.

    [error steps](https://streams.spec.whatwg.org/#read-request-error-steps){#ref-for-read-request-error-steps link-type="dfn"}, given `e`{.variable}

    :   1.  [Queue a fetch
            task](#queue-a-fetch-task){#ref-for-queue-a-fetch-task②
            link-type="dfn"} to run `processBodyError`{.variable} given
            `e`{.variable}, with `taskDestination`{.variable}.

2.  [Read a
    chunk](https://streams.spec.whatwg.org/#readablestreamdefaultreader-read-a-chunk){#ref-for-readablestreamdefaultreader-read-a-chunk
    link-type="dfn"} from `reader`{.variable} given
    `readRequest`{.variable}.
:::

::: {.algorithm algorithm="fully read" algorithm-for="body"}
To [fully read]{#body-fully-read .dfn .dfn-paneled dfn-for="body"
dfn-type="dfn" export=""} a [body](#concept-body){#ref-for-concept-body③
link-type="dfn"} `body`{.variable}, given an algorithm
`processBody`{.variable}, an algorithm `processBodyError`{.variable},
and an optional null, [parallel
queue](https://html.spec.whatwg.org/multipage/infrastructure.html#parallel-queue){#ref-for-parallel-queue⑤
link-type="dfn"}, or [global
object](https://html.spec.whatwg.org/multipage/webappapis.html#global-object){#ref-for-global-object⑥
link-type="dfn"} `taskDestination`{.variable} (default null), run these
steps. `processBody`{.variable} must be an algorithm accepting a [byte
sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence⑨
link-type="dfn"}. `processBodyError`{.variable} must be an algorithm
optionally accepting an
[exception](https://webidl.spec.whatwg.org/#dfn-exception){#ref-for-dfn-exception
link-type="dfn"}.

1.  If `taskDestination`{.variable} is null, then set
    `taskDestination`{.variable} to the result of [starting a new
    parallel
    queue](https://html.spec.whatwg.org/multipage/infrastructure.html#starting-a-new-parallel-queue){#ref-for-starting-a-new-parallel-queue①
    link-type="dfn"}.

2.  Let `successSteps`{.variable} given a [byte
    sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence①⓪
    link-type="dfn"} `bytes`{.variable} be to [queue a fetch
    task](#queue-a-fetch-task){#ref-for-queue-a-fetch-task③
    link-type="dfn"} to run `processBody`{.variable} given
    `bytes`{.variable}, with `taskDestination`{.variable}.

3.  Let `errorSteps`{.variable} optionally given an
    [exception](https://webidl.spec.whatwg.org/#dfn-exception){#ref-for-dfn-exception①
    link-type="dfn"} `exception`{.variable} be to [queue a fetch
    task](#queue-a-fetch-task){#ref-for-queue-a-fetch-task④
    link-type="dfn"} to run `processBodyError`{.variable} given
    `exception`{.variable}, with `taskDestination`{.variable}.

4.  Let `reader`{.variable} be the result of [getting a
    reader](https://streams.spec.whatwg.org/#readablestream-get-a-reader){#ref-for-readablestream-get-a-reader①
    link-type="dfn"} for `body`{.variable}'s
    [stream](#concept-body-stream){#ref-for-concept-body-stream④
    link-type="dfn"}. If that threw an exception, then run
    `errorSteps`{.variable} with that exception and return.

5.  [Read all
    bytes](https://streams.spec.whatwg.org/#readablestreamdefaultreader-read-all-bytes){#ref-for-readablestreamdefaultreader-read-all-bytes
    link-type="dfn"} from `reader`{.variable}, given
    `successSteps`{.variable} and `errorSteps`{.variable}.
:::

------------------------------------------------------------------------

A [body with type]{#body-with-type .dfn .dfn-paneled dfn-type="dfn"
export=""} is a
[tuple](https://infra.spec.whatwg.org/#tuple){#ref-for-tuple②
link-type="dfn"} that consists of a [body]{#body-with-type-body .dfn
.dfn-paneled dfn-for="body with type" dfn-type="dfn" export=""} (a
[body](#concept-body){#ref-for-concept-body④ link-type="dfn"}) and a
[type]{#body-with-type-type .dfn .dfn-paneled dfn-for="body with type"
dfn-type="dfn" export=""} (a [header
value](#header-value){#ref-for-header-value⑥ link-type="dfn"} or null).

------------------------------------------------------------------------

::: {.algorithm algorithm="handle content codings"}
To [handle content codings]{#handle-content-codings .dfn .dfn-paneled
dfn-type="dfn" export=""} given `codings`{.variable} and
`bytes`{.variable}, run these steps:

1.  If `codings`{.variable} are not supported, then return
    `bytes`{.variable}.

2.  Return the result of decoding `bytes`{.variable} with
    `codings`{.variable} as explained in HTTP, if decoding does not
    result in an error, and failure otherwise.
    [\[HTTP\]](#biblio-http "HTTP Semantics"){link-type="biblio"}
:::

#### [2.2.5. ]{.secno}[Requests]{.content}[](#requests){.self-link} {#requests .heading .settled level="2.2.5"}

This section documents how requests work in detail. To get started, see
[Setting up a request](#fetch-elsewhere-request).

The input to [fetch](#concept-fetch){#ref-for-concept-fetch④
link-type="dfn"} is a [request]{#concept-request .dfn .dfn-paneled
dfn-type="dfn" export=""}.

A [request](#concept-request){#ref-for-concept-request① link-type="dfn"}
has an associated [method]{#concept-request-method .dfn .dfn-paneled
dfn-for="request" dfn-type="dfn" export=""} (a
[method](#concept-method){#ref-for-concept-method⑤ link-type="dfn"}).
Unless stated otherwise it is \``GET`\`.

This can be updated during redirects to \``GET`\` as described in [HTTP
fetch](#concept-http-fetch){#ref-for-concept-http-fetch
link-type="dfn"}.

A [request](#concept-request){#ref-for-concept-request② link-type="dfn"}
has an associated [URL]{#concept-request-url .dfn .dfn-paneled
dfn-for="request" dfn-type="dfn" export=""} (a
[URL](https://url.spec.whatwg.org/#concept-url){#ref-for-concept-url①
link-type="dfn"}).

Implementations are encouraged to make this a pointer to the first
[URL](https://url.spec.whatwg.org/#concept-url){#ref-for-concept-url②
link-type="dfn"} in
[request](#concept-request){#ref-for-concept-request③ link-type="dfn"}'s
[URL list](#concept-request-url-list){#ref-for-concept-request-url-list
link-type="dfn"}. It is provided as a distinct field solely for the
convenience of other standards hooking into Fetch.

A [request](#concept-request){#ref-for-concept-request④ link-type="dfn"}
has an associated [local-URLs-only flag]{#local-urls-only-flag .dfn
.dfn-paneled dfn-type="dfn" export=""}. Unless stated otherwise it is
unset.

A [request](#concept-request){#ref-for-concept-request⑤ link-type="dfn"}
has an associated [header list]{#concept-request-header-list .dfn
.dfn-paneled dfn-for="request" dfn-type="dfn" export=""} (a [header
list](#concept-header-list){#ref-for-concept-header-list①⑦
link-type="dfn"}). Unless stated otherwise it is « ».

A [request](#concept-request){#ref-for-concept-request⑥ link-type="dfn"}
has an associated [unsafe-request flag]{#unsafe-request-flag .dfn
.dfn-paneled dfn-for="request" dfn-type="dfn" export=""}. Unless stated
otherwise it is unset.

The [unsafe-request
flag](#unsafe-request-flag){#ref-for-unsafe-request-flag
link-type="dfn"} is set by APIs such as
[`fetch()`](#dom-global-fetch){#ref-for-dom-global-fetch① .idl-code
link-type="method"} and
[`XMLHttpRequest`{.idl}](https://xhr.spec.whatwg.org/#xmlhttprequest){#ref-for-xmlhttprequest②
link-type="idl"} to ensure a [CORS-preflight
fetch](#cors-preflight-fetch-0){#ref-for-cors-preflight-fetch-0
link-type="dfn"} is done based on the supplied
[method](#concept-request-method){#ref-for-concept-request-method
link-type="dfn"} and [header
list](#concept-request-header-list){#ref-for-concept-request-header-list
link-type="dfn"}. It does not free an API from outlawing [forbidden
methods](#forbidden-method){#ref-for-forbidden-method① link-type="dfn"}
and [forbidden
request-headers](#forbidden-request-header){#ref-for-forbidden-request-header
link-type="dfn"}.

A [request](#concept-request){#ref-for-concept-request⑦ link-type="dfn"}
has an associated [body]{#concept-request-body .dfn .dfn-paneled
dfn-for="request" dfn-type="dfn" export=""} (null, a [byte
sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence①①
link-type="dfn"}, or a [body](#concept-body){#ref-for-concept-body⑤
link-type="dfn"}). Unless stated otherwise it is null.

A [byte
sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence①②
link-type="dfn"} will be [safely
extracted](#bodyinit-safely-extract){#ref-for-bodyinit-safely-extract①
link-type="dfn"} into a [body](#concept-body){#ref-for-concept-body⑥
link-type="dfn"} early on in
[fetch](#concept-fetch){#ref-for-concept-fetch⑤ link-type="dfn"}. As
part of [HTTP fetch](#concept-http-fetch){#ref-for-concept-http-fetch①
link-type="dfn"} it is possible for this field to be set to null due to
certain redirects.

------------------------------------------------------------------------

A [request](#concept-request){#ref-for-concept-request⑧ link-type="dfn"}
has an associated [client]{#concept-request-client .dfn .dfn-paneled
dfn-for="request" dfn-type="dfn" export=""} (null or an [environment
settings
object](https://html.spec.whatwg.org/multipage/webappapis.html#environment-settings-object){#ref-for-environment-settings-object
link-type="dfn"}).

A [request](#concept-request){#ref-for-concept-request⑨ link-type="dfn"}
has an associated [reserved client]{#concept-request-reserved-client
.dfn .dfn-paneled dfn-for="request" dfn-type="dfn" export=""} (null, an
[environment](https://html.spec.whatwg.org/multipage/webappapis.html#environment){#ref-for-environment
link-type="dfn"}, or an [environment settings
object](https://html.spec.whatwg.org/multipage/webappapis.html#environment-settings-object){#ref-for-environment-settings-object①
link-type="dfn"}). Unless stated otherwise it is null.

This is only used by [navigation
requests](#navigation-request){#ref-for-navigation-request
link-type="dfn"} and worker requests, but not service worker requests.
It references an
[environment](https://html.spec.whatwg.org/multipage/webappapis.html#environment){#ref-for-environment①
link-type="dfn"} for a [navigation
request](#navigation-request){#ref-for-navigation-request①
link-type="dfn"} and an [environment settings
object](https://html.spec.whatwg.org/multipage/webappapis.html#environment-settings-object){#ref-for-environment-settings-object②
link-type="dfn"} for a worker request.

A [request](#concept-request){#ref-for-concept-request①⓪
link-type="dfn"} has an associated [replaces client
id]{#concept-request-replaces-client-id .dfn .dfn-paneled
dfn-for="request" dfn-type="dfn" export=""} (a string). Unless stated
otherwise it is the empty string.

This is only used by [navigation
requests](#navigation-request){#ref-for-navigation-request②
link-type="dfn"}. It is the
[id](https://html.spec.whatwg.org/multipage/webappapis.html#concept-environment-id){#ref-for-concept-environment-id
link-type="dfn"} of the [target browsing
context](https://html.spec.whatwg.org/multipage/webappapis.html#concept-environment-target-browsing-context){#ref-for-concept-environment-target-browsing-context
link-type="dfn"}'s [active
document](https://html.spec.whatwg.org/multipage/document-sequences.html#nav-document){#ref-for-nav-document
link-type="dfn"}'s [environment settings
object](https://html.spec.whatwg.org/multipage/webappapis.html#environment-settings-object){#ref-for-environment-settings-object③
link-type="dfn"}.

A [request](#concept-request){#ref-for-concept-request①①
link-type="dfn"} has an associated [traversable for user
prompts]{#concept-request-window .dfn .dfn-paneled dfn-for="request"
dfn-type="dfn" export=""}, that is \"`no-traversable`\", \"`client`\",
or a [traversable
navigable](https://html.spec.whatwg.org/multipage/document-sequences.html#traversable-navigable){#ref-for-traversable-navigable
link-type="dfn"}. Unless stated otherwise it is \"`client`\".

::: {.note role="note"}
This is used to determine whether and where to show necessary UI for the
request, such as authentication prompts or client certificate dialogs.

\"`no-traversable`\"
:   No UI is shown; usually the request fails with a [network
    error](#concept-network-error){#ref-for-concept-network-error
    link-type="dfn"}.

\"`client`\"
:   This value will automatically be changed to either
    \"`no-traversable`\" or to a [traversable
    navigable](https://html.spec.whatwg.org/multipage/document-sequences.html#traversable-navigable){#ref-for-traversable-navigable①
    link-type="dfn"} derived from the request's
    [client](#concept-request-client){#ref-for-concept-request-client
    link-type="dfn"} during
    [fetching](#concept-fetch){#ref-for-concept-fetch⑥ link-type="dfn"}.
    This provides a convenient way for standards to not have to
    explicitly set a request's [traversable for user
    prompts](#concept-request-window){#ref-for-concept-request-window
    link-type="dfn"}.

a [traversable navigable](https://html.spec.whatwg.org/multipage/document-sequences.html#traversable-navigable){#ref-for-traversable-navigable② link-type="dfn"}
:   The UI shown will be associated with the browser interface elements
    that are displaying that [traversable
    navigable](https://html.spec.whatwg.org/multipage/document-sequences.html#traversable-navigable){#ref-for-traversable-navigable③
    link-type="dfn"}.
:::

When displaying a user interface associated with a request in that
request's [traversable for user
prompts](#concept-request-window){#ref-for-concept-request-window①
link-type="dfn"}, the user agent should update the address bar to
display something derived from the request's [current
URL](#concept-request-current-url){#ref-for-concept-request-current-url
link-type="dfn"} (and not, e.g., leave it at its previous value, derived
from the URL of the request's initiator). Additionally, the user agent
should avoid displaying content from the request's initiator in the
[traversable for user
prompts](#concept-request-window){#ref-for-concept-request-window②
link-type="dfn"}, especially in the case of cross-origin requests.
Displaying a blank page behind such prompts is a good way to fulfill
these requirements. Failing to follow these guidelines can confuse users
as to which origin is responsible for the prompt.

A [request](#concept-request){#ref-for-concept-request①②
link-type="dfn"} has an associated boolean
[keepalive]{#request-keepalive-flag .dfn .dfn-paneled dfn-for="request"
dfn-type="dfn" export=""}. Unless stated otherwise it is false.

This can be used to allow the request to outlive the [environment
settings
object](https://html.spec.whatwg.org/multipage/webappapis.html#environment-settings-object){#ref-for-environment-settings-object④
link-type="dfn"}, e.g., `navigator.sendBeacon()` and the HTML `img`
element use this. Requests with this set to true are subject to
additional processing requirements.

A [request](#concept-request){#ref-for-concept-request①③
link-type="dfn"} has an associated [initiator
type]{#request-initiator-type .dfn .dfn-paneled dfn-for="request"
dfn-type="dfn" export=""}, which is null, \"`audio`\", \"`beacon`\",
\"`body`\", \"`css`\", \"`early-hints`\", \"`embed`\", \"`fetch`\",
\"`font`\", \"`frame`\", \"`iframe`\", \"`image`\", \"`img`\",
\"`input`\", \"`link`\", \"`object`\", \"`ping`\", \"`script`\",
\"`track`\", \"`video`\", \"`xmlhttprequest`\", or \"`other`\". Unless
stated otherwise it is null.
[\[RESOURCE-TIMING\]](#biblio-resource-timing "Resource Timing"){link-type="biblio"}

A [request](#concept-request){#ref-for-concept-request①④
link-type="dfn"} has an associated [service-workers
mode]{#request-service-workers-mode .dfn .dfn-paneled dfn-for="request"
dfn-type="dfn" export=""}, that is \"`all`\" or \"`none`\". Unless
stated otherwise it is \"`all`\".

::: {.note role="note"}
This determines which service workers will receive a
[`fetch`{.idl}](https://w3c.github.io/ServiceWorker/#service-worker-global-scope-fetch-event){#ref-for-service-worker-global-scope-fetch-event
.idl-code link-type="event"} event for this fetch.

\"`all`\"
:   Relevant service workers will get a
    [`fetch`{.idl}](https://w3c.github.io/ServiceWorker/#service-worker-global-scope-fetch-event){#ref-for-service-worker-global-scope-fetch-event①
    .idl-code link-type="event"} event for this fetch.

\"`none`\"
:   No service workers will get events for this fetch.
:::

A [request](#concept-request){#ref-for-concept-request①⑤
link-type="dfn"} has an associated
[initiator]{#concept-request-initiator .dfn .dfn-paneled
dfn-for="request" dfn-type="dfn" export=""}, which is the empty string,
\"`download`\", \"`imageset`\", \"`manifest`\", \"`prefetch`\",
\"`prerender`\", or \"`xslt`\". Unless stated otherwise it is the empty
string.

A [request](#concept-request){#ref-for-concept-request①⑥
link-type="dfn"}'s
[initiator](#concept-request-initiator){#ref-for-concept-request-initiator
link-type="dfn"} is not particularly granular for the time being as
other specifications do not require it to be. It is primarily a
specification device to assist defining CSP and Mixed Content. It is not
exposed to JavaScript.
[\[CSP\]](#biblio-csp "Content Security Policy Level 3"){link-type="biblio"}
[\[MIX\]](#biblio-mix "Mixed Content"){link-type="biblio"}

A [request](#concept-request){#ref-for-concept-request①⑦
link-type="dfn"} has an associated
[destination]{#concept-request-destination .dfn .dfn-paneled
dfn-for="request" dfn-type="dfn" export=""}, which is the empty string,
\"`audio`\", \"`audioworklet`\", \"`document`\", \"`embed`\",
\"`font`\", \"`frame`\", \"`iframe`\", \"`image`\", \"`json`\",
\"`manifest`\", \"`object`\", \"`paintworklet`\", \"`report`\",
\"`script`\", \"`serviceworker`\", \"`sharedworker`\", \"`style`\",
\"`track`\", \"`video`\", \"`webidentity`\", \"`worker`\", or
\"`xslt`\". Unless stated otherwise it is the empty string.

These are reflected on
[`RequestDestination`{.idl}](#requestdestination){#ref-for-requestdestination
link-type="idl"} except for \"`serviceworker`\" and \"`webidentity`\" as
fetches with those destinations skip service workers.

A [request](#concept-request){#ref-for-concept-request①⑧
link-type="dfn"}'s
[destination](#concept-request-destination){#ref-for-concept-request-destination
link-type="dfn"} is [script-like]{#request-destination-script-like .dfn
.dfn-paneled dfn-for="request/destination" dfn-type="dfn" export=""} if
it is \"`audioworklet`\", \"`paintworklet`\", \"`script`\",
\"`serviceworker`\", \"`sharedworker`\", or \"`worker`\".

Algorithms that use
[script-like](#request-destination-script-like){#ref-for-request-destination-script-like
link-type="dfn"} should also consider \"`xslt`\" as that too can cause
script execution. It is not included in the list as it is not always
relevant and might require different behavior.

::: {#destination-table .note role="note"}
[](#destination-table){.self-link}

The following table illustrates the relationship between a
[request](#concept-request){#ref-for-concept-request①⑨
link-type="dfn"}'s
[initiator](#concept-request-initiator){#ref-for-concept-request-initiator①
link-type="dfn"},
[destination](#concept-request-destination){#ref-for-concept-request-destination①
link-type="dfn"}, CSP directives, and features. It is not exhaustive
with respect to features. Features need to have the relevant values
defined in their respective standards.

[Initiator](#concept-request-initiator){#ref-for-concept-request-initiator②
link-type="dfn"}

[Destination](#concept-request-destination){#ref-for-concept-request-destination②
link-type="dfn"}

CSP directive

Features

\"\"

\"`report`\"

---

CSP, NEL reports.

\"`document`\"

HTML's navigate algorithm (top-level only).

\"`frame`\"

`child-src`

HTML's `<frame>`

\"`iframe`\"

`child-src`

HTML's `<iframe>`

\"\"

`connect-src`

`navigator.sendBeacon()`,
[`EventSource`{.idl}](https://html.spec.whatwg.org/multipage/server-sent-events.html#eventsource){#ref-for-eventsource
link-type="idl"}, HTML's `<a ping="">` and `<area ping="">`,
[`fetch()`](#dom-global-fetch){#ref-for-dom-global-fetch② .idl-code
link-type="method"},
[`XMLHttpRequest`{.idl}](https://xhr.spec.whatwg.org/#xmlhttprequest){#ref-for-xmlhttprequest③
link-type="idl"},
[`WebSocket`{.idl}](https://websockets.spec.whatwg.org/#websocket){#ref-for-websocket
link-type="idl"}, Cache API

\"`object`\"

`object-src`

HTML's `<object>`

\"`embed`\"

`object-src`

HTML's `<embed>`

\"`audio`\"

`media-src`

HTML's `<audio>`

\"`font`\"

`font-src`

CSS\' `@font-face`

\"`image`\"

`img-src`

HTML's `<img src>`, `/favicon.ico` resource, SVG's `<image>`, CSS\'
`background-image`, CSS\' `cursor`, CSS\' `list-style-image`, ...

\"`audioworklet`\"

`script-src`

`audioWorklet.addModule()`

\"`paintworklet`\"

`script-src`

`CSS.paintWorklet.addModule()`

\"`script`\"

`script-src`

HTML's `<script>`, `importScripts()`

\"`serviceworker`\"

`child-src`, `script-src`, `worker-src`

`navigator.serviceWorker.register()`

\"`sharedworker`\"

`child-src`, `script-src`, `worker-src`

`SharedWorker`

\"`webidentity`\"

`connect-src`

`Federated Credential Management requests`

\"`worker`\"

`child-src`, `script-src`, `worker-src`

`Worker`

\"`json`\"

`connect-src`

`import "..." with { type: "json" }`

\"`style`\"

`style-src`

HTML's `<link rel=stylesheet>`, CSS\' `@import`,
`import "..." with { type: "css" }`

\"`track`\"

`media-src`

HTML's `<track>`

\"`video`\"

`media-src`

HTML's `<video>` element

\"`download`\"

\"\"

---

HTML's `download=""`, \"Save Link As...\" UI

\"`imageset`\"

\"`image`\"

`img-src`

HTML's `<img srcset>` and `<picture>`

\"`manifest`\"

\"`manifest`\"

`manifest-src`

HTML's `<link rel=manifest>`

\"`prefetch`\"

\"\"

`default-src` (no specific directive)

HTML's `<link rel=prefetch>`

\"`prerender`\"

HTML's `<link rel=prerender>`

\"`xslt`\"

\"`xslt`\"

`script-src`

`<?xml-stylesheet>`

CSP's `form-action` needs to be a hook directly in HTML's navigate or
form submission algorithm.

CSP will also need to check
[request](#concept-request){#ref-for-concept-request②⓪
link-type="dfn"}'s
[client](#concept-request-client){#ref-for-concept-request-client①
link-type="dfn"}'s [global
object](https://html.spec.whatwg.org/multipage/webappapis.html#concept-settings-object-global){#ref-for-concept-settings-object-global
link-type="dfn"}'s [associated
`Document`](https://html.spec.whatwg.org/multipage/nav-history-apis.html#concept-document-window){#ref-for-concept-document-window
link-type="dfn"}'s [ancestor
navigables](https://html.spec.whatwg.org/multipage/document-sequences.html#ancestor-navigables){#ref-for-ancestor-navigables
link-type="dfn"} for various CSP directives.
:::

------------------------------------------------------------------------

A [request](#concept-request){#ref-for-concept-request②①
link-type="dfn"} has an associated [priority]{#request-priority .dfn
.dfn-paneled dfn-for="request" dfn-type="dfn" export=""}, which is
\"`high`\", \"`low`\", or \"`auto`\". Unless stated otherwise it is
\"`auto`\".

A [request](#concept-request){#ref-for-concept-request②②
link-type="dfn"} has an associated [[]{#concept-request-priority
.bs-old-id}internal priority]{#request-internal-priority .dfn
.dfn-paneled dfn-for="request" dfn-type="dfn" export=""} (null or an
[implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#ref-for-implementation-defined①
link-type="dfn"} object). Unless otherwise stated it is null.

A [request](#concept-request){#ref-for-concept-request②③
link-type="dfn"} has an associated [origin]{#concept-request-origin .dfn
.dfn-paneled dfn-for="request" dfn-type="dfn" export=""}, which is
\"`client`\" or an
[origin](https://html.spec.whatwg.org/multipage/browsers.html#concept-origin){#ref-for-concept-origin
link-type="dfn"}. Unless stated otherwise it is \"`client`\".

\"`client`\" is changed to an
[origin](https://html.spec.whatwg.org/multipage/browsers.html#concept-origin){#ref-for-concept-origin①
link-type="dfn"} during
[fetching](#concept-fetch){#ref-for-concept-fetch⑦ link-type="dfn"}. It
provides a convenient way for standards to not have to set
[request](#concept-request){#ref-for-concept-request②④
link-type="dfn"}'s
[origin](#concept-request-origin){#ref-for-concept-request-origin
link-type="dfn"}.

A [request](#concept-request){#ref-for-concept-request②⑤
link-type="dfn"} has an associated [policy
container]{#concept-request-policy-container .dfn .dfn-paneled
dfn-for="request" dfn-type="dfn" export=""}, which is \"`client`\" or a
[policy
container](https://html.spec.whatwg.org/multipage/browsers.html#policy-container){#ref-for-policy-container
link-type="dfn"}. Unless stated otherwise it is \"`client`\".

\"`client`\" is changed to a [policy
container](https://html.spec.whatwg.org/multipage/browsers.html#policy-container){#ref-for-policy-container①
link-type="dfn"} during
[fetching](#concept-fetch){#ref-for-concept-fetch⑧ link-type="dfn"}. It
provides a convenient way for standards to not have to set
[request](#concept-request){#ref-for-concept-request②⑥
link-type="dfn"}'s [policy
container](#concept-request-policy-container){#ref-for-concept-request-policy-container
link-type="dfn"}.

A [request](#concept-request){#ref-for-concept-request②⑦
link-type="dfn"} has an associated [referrer]{#concept-request-referrer
.dfn .dfn-paneled dfn-for="request" dfn-type="dfn" export=""}, which is
\"`no-referrer`\", \"`client`\", or a
[URL](https://url.spec.whatwg.org/#concept-url){#ref-for-concept-url③
link-type="dfn"}. Unless stated otherwise it is \"`client`\".

\"`client`\" is changed to \"`no-referrer`\" or a
[URL](https://url.spec.whatwg.org/#concept-url){#ref-for-concept-url④
link-type="dfn"} during
[fetching](#concept-fetch){#ref-for-concept-fetch⑨ link-type="dfn"}. It
provides a convenient way for standards to not have to set
[request](#concept-request){#ref-for-concept-request②⑧
link-type="dfn"}'s
[referrer](#concept-request-referrer){#ref-for-concept-request-referrer
link-type="dfn"}.

A [request](#concept-request){#ref-for-concept-request②⑨
link-type="dfn"} has an associated [referrer
policy]{#concept-request-referrer-policy .dfn .dfn-paneled
dfn-for="request" dfn-type="dfn" export=""}, which is a [referrer
policy](https://w3c.github.io/webappsec-referrer-policy/#referrer-policy){#ref-for-referrer-policy
link-type="dfn"}. Unless stated otherwise it is the empty string.
[\[REFERRER\]](#biblio-referrer "Referrer Policy"){link-type="biblio"}

This can be used to override the referrer policy to be used for this
[request](#concept-request){#ref-for-concept-request③⓪ link-type="dfn"}.

A [request](#concept-request){#ref-for-concept-request③①
link-type="dfn"} has an associated [mode]{#concept-request-mode .dfn
.dfn-paneled dfn-for="request" dfn-type="dfn" export=""}, which is
\"`same-origin`\", \"`cors`\", \"`no-cors`\", \"`navigate`\", or
\"`websocket`\". Unless stated otherwise, it is \"`no-cors`\".

::: {.note role="note"}

\"`same-origin`\"
:   Used to ensure requests are made to same-origin URLs.
    [Fetch](#concept-fetch){#ref-for-concept-fetch①⓪ link-type="dfn"}
    will return a [network
    error](#concept-network-error){#ref-for-concept-network-error①
    link-type="dfn"} if the request is not made to a same-origin URL.

\"`cors`\"
:   For requests whose [response
    tainting](#concept-request-response-tainting){#ref-for-concept-request-response-tainting
    link-type="dfn"} gets set to \"`cors`\", makes the request a [CORS
    request](#cors-request){#ref-for-cors-request link-type="dfn"} ---
    in which case, fetch will return a [network
    error](#concept-network-error){#ref-for-concept-network-error②
    link-type="dfn"} if the requested resource does not understand the
    [CORS protocol](#cors-protocol){#ref-for-cors-protocol
    link-type="dfn"}, or if the requested resource is one that
    intentionally does not participate in the [CORS
    protocol](#cors-protocol){#ref-for-cors-protocol① link-type="dfn"}.

\"`no-cors`\"
:   Restricts requests to using [CORS-safelisted
    methods](#cors-safelisted-method){#ref-for-cors-safelisted-method
    link-type="dfn"} and [CORS-safelisted
    request-headers](#cors-safelisted-request-header){#ref-for-cors-safelisted-request-header②
    link-type="dfn"}. Upon success, fetch will return an [opaque
    filtered
    response](#concept-filtered-response-opaque){#ref-for-concept-filtered-response-opaque
    link-type="dfn"}.

\"`navigate`\"
:   This is a special mode used only when
    [navigating](https://html.spec.whatwg.org/multipage/nav-history-apis.html#blocking-navigating){#ref-for-blocking-navigating
    link-type="dfn"} between documents.

\"`websocket`\"
:   This is a special mode used only when [establishing a WebSocket
    connection](https://websockets.spec.whatwg.org/#concept-websocket-establish){#ref-for-concept-websocket-establish①
    link-type="dfn"}.

Even though the default
[request](#concept-request){#ref-for-concept-request③② link-type="dfn"}
[mode](#concept-request-mode){#ref-for-concept-request-mode
link-type="dfn"} is \"`no-cors`\", standards are highly discouraged from
using it for new features. It is rather unsafe.
:::

A [request](#concept-request){#ref-for-concept-request③③
link-type="dfn"} has an associated [use-CORS-preflight
flag]{#use-cors-preflight-flag .dfn .dfn-paneled dfn-for="request"
dfn-type="dfn" export=""}. Unless stated otherwise, it is unset.

The [use-CORS-preflight
flag](#use-cors-preflight-flag){#ref-for-use-cors-preflight-flag
link-type="dfn"} being set is one of several conditions that results in
a [CORS-preflight
request](#cors-preflight-request){#ref-for-cors-preflight-request
link-type="dfn"}. The [use-CORS-preflight
flag](#use-cors-preflight-flag){#ref-for-use-cors-preflight-flag①
link-type="dfn"} is set if either one or more event listeners are
registered on an
[`XMLHttpRequestUpload`{.idl}](https://xhr.spec.whatwg.org/#xmlhttprequestupload){#ref-for-xmlhttprequestupload
link-type="idl"} object or if a
[`ReadableStream`{.idl}](https://streams.spec.whatwg.org/#readablestream){#ref-for-readablestream①
link-type="idl"} object is used in a request.

A [request](#concept-request){#ref-for-concept-request③④
link-type="dfn"} has an associated [credentials
mode]{#concept-request-credentials-mode .dfn .dfn-paneled
dfn-for="request" dfn-type="dfn" export=""}, which is \"`omit`\",
\"`same-origin`\", or \"`include`\". Unless stated otherwise, it is
\"`same-origin`\".

::: {.note role="note"}

\"`omit`\"
:   Excludes credentials from this request, and causes any credentials
    sent back in the response to be ignored.

\"`same-origin`\"
:   Include credentials with requests made to same-origin URLs, and use
    any credentials sent back in responses from same-origin URLs.

\"`include`\"
:   Always includes credentials with this request, and always use any
    credentials sent back in the response.

[Request](#concept-request){#ref-for-concept-request③⑤
link-type="dfn"}'s [credentials
mode](#concept-request-credentials-mode){#ref-for-concept-request-credentials-mode
link-type="dfn"} controls the flow of
[credentials](#credentials){#ref-for-credentials link-type="dfn"} during
a [fetch](#concept-fetch){#ref-for-concept-fetch①① link-type="dfn"}.
When [request](#concept-request){#ref-for-concept-request③⑥
link-type="dfn"}'s
[mode](#concept-request-mode){#ref-for-concept-request-mode①
link-type="dfn"} is \"`navigate`\", its [credentials
mode](#concept-request-credentials-mode){#ref-for-concept-request-credentials-mode①
link-type="dfn"} is assumed to be \"`include`\" and
[fetch](#concept-fetch){#ref-for-concept-fetch①② link-type="dfn"} does
not currently account for other values. If HTML changes here, this
standard will need corresponding changes.
:::

A [request](#concept-request){#ref-for-concept-request③⑦
link-type="dfn"} has an associated [use-URL-credentials
flag]{#concept-request-use-url-credentials-flag .dfn .dfn-paneled
dfn-for="request" dfn-type="dfn" export=""}. Unless stated otherwise, it
is unset.

When this flag is set, when a
[request](#concept-request){#ref-for-concept-request③⑧
link-type="dfn"}'s
[URL](#concept-request-url){#ref-for-concept-request-url
link-type="dfn"} has a
[username](https://url.spec.whatwg.org/#concept-url-username){#ref-for-concept-url-username
link-type="dfn"} and
[password](https://url.spec.whatwg.org/#concept-url-password){#ref-for-concept-url-password
link-type="dfn"}, and there is an available [authentication
entry](#authentication-entry){#ref-for-authentication-entry①
link-type="dfn"} for the
[request](#concept-request){#ref-for-concept-request③⑨ link-type="dfn"},
then the
[URL](https://url.spec.whatwg.org/#concept-url){#ref-for-concept-url⑤
link-type="dfn"}'s credentials are preferred over that of the
[authentication
entry](#authentication-entry){#ref-for-authentication-entry②
link-type="dfn"}. Modern specifications avoid setting this flag, since
putting credentials in
[URLs](https://url.spec.whatwg.org/#concept-url){#ref-for-concept-url⑥
link-type="dfn"} is discouraged, but some older features set it for
compatibility reasons.

A [request](#concept-request){#ref-for-concept-request④⓪
link-type="dfn"} has an associated [cache
mode]{#concept-request-cache-mode .dfn .dfn-paneled dfn-for="request"
dfn-type="dfn" export=""}, which is \"`default`\", \"`no-store`\",
\"`reload`\", \"`no-cache`\", \"`force-cache`\", or
\"`only-if-cached`\". Unless stated otherwise, it is \"`default`\".

::: {.note role="note"}

\"`default`\"
:   [Fetch](#concept-fetch){#ref-for-concept-fetch①③ link-type="dfn"}
    will inspect the HTTP cache on the way to the network. If the HTTP
    cache contains a matching [fresh
    response](#concept-fresh-response){#ref-for-concept-fresh-response
    link-type="dfn"} it will be returned. If the HTTP cache contains a
    matching [stale-while-revalidate
    response](#concept-stale-while-revalidate-response){#ref-for-concept-stale-while-revalidate-response
    link-type="dfn"} it will be returned, and a conditional network
    fetch will be made to update the entry in the HTTP cache. If the
    HTTP cache contains a matching [stale
    response](#concept-stale-response){#ref-for-concept-stale-response
    link-type="dfn"}, a conditional network fetch will be returned to
    update the entry in the HTTP cache. Otherwise, a non-conditional
    network fetch will be returned to update the entry in the HTTP
    cache. [\[HTTP\]](#biblio-http "HTTP Semantics"){link-type="biblio"}
    [\[HTTP-CACHING\]](#biblio-http-caching "HTTP Caching"){link-type="biblio"}
    [\[STALE-WHILE-REVALIDATE\]](#biblio-stale-while-revalidate "HTTP Cache-Control Extensions for Stale Content"){link-type="biblio"}

\"`no-store`\"
:   Fetch behaves as if there is no HTTP cache at all.

\"`reload`\"
:   Fetch behaves as if there is no HTTP cache on the way to the
    network. Ergo, it creates a normal request and updates the HTTP
    cache with the response.

\"`no-cache`\"
:   Fetch creates a conditional request if there is a response in the
    HTTP cache and a normal request otherwise. It then updates the HTTP
    cache with the response.

\"`force-cache`\"
:   Fetch uses any response in the HTTP cache matching the request, not
    paying attention to staleness. If there was no response, it creates
    a normal request and updates the HTTP cache with the response.

\"`only-if-cached`\"
:   Fetch uses any response in the HTTP cache matching the request, not
    paying attention to staleness. If there was no response, it returns
    a network error. (Can only be used when
    [request](#concept-request){#ref-for-concept-request④①
    link-type="dfn"}'s
    [mode](#concept-request-mode){#ref-for-concept-request-mode②
    link-type="dfn"} is \"`same-origin`\". Any cached redirects will be
    followed assuming
    [request](#concept-request){#ref-for-concept-request④②
    link-type="dfn"}'s [redirect
    mode](#concept-request-redirect-mode){#ref-for-concept-request-redirect-mode
    link-type="dfn"} is \"`follow`\" and the redirects do not violate
    [request](#concept-request){#ref-for-concept-request④③
    link-type="dfn"}'s
    [mode](#concept-request-mode){#ref-for-concept-request-mode③
    link-type="dfn"}.)

If [header
list](#concept-request-header-list){#ref-for-concept-request-header-list①
link-type="dfn"}
[contains](#header-list-contains){#ref-for-header-list-contains⑦
link-type="dfn"} \``If-Modified-Since`\`, \``If-None-Match`\`,
\``If-Unmodified-Since`\`, \``If-Match`\`, or \``If-Range`\`,
[fetch](#concept-fetch){#ref-for-concept-fetch①④ link-type="dfn"} will
set [cache
mode](#concept-request-cache-mode){#ref-for-concept-request-cache-mode
link-type="dfn"} to \"`no-store`\" if it is \"`default`\".
:::

A [request](#concept-request){#ref-for-concept-request④④
link-type="dfn"} has an associated [redirect
mode]{#concept-request-redirect-mode .dfn .dfn-paneled dfn-for="request"
dfn-type="dfn" export=""}, which is \"`follow`\", \"`error`\", or
\"`manual`\". Unless stated otherwise, it is \"`follow`\".

::: {.note role="note"}

\"`follow`\"
:   Follow all redirects incurred when fetching a resource.

\"`error`\"
:   Return a [network
    error](#concept-network-error){#ref-for-concept-network-error③
    link-type="dfn"} when a request is met with a redirect.

\"`manual`\"
:   Retrieves an [opaque-redirect filtered
    response](#concept-filtered-response-opaque-redirect){#ref-for-concept-filtered-response-opaque-redirect
    link-type="dfn"} when a request is met with a redirect, to allow a
    service worker to replay the redirect offline. The response is
    otherwise indistinguishable from a [network
    error](#concept-network-error){#ref-for-concept-network-error④
    link-type="dfn"}, to not violate [atomic HTTP redirect
    handling](#atomic-http-redirect-handling){#ref-for-atomic-http-redirect-handling
    link-type="dfn"}.
:::

A [request](#concept-request){#ref-for-concept-request④⑤
link-type="dfn"} has associated [integrity
metadata]{#concept-request-integrity-metadata .dfn .dfn-paneled
dfn-for="request" dfn-type="dfn" export=""} (a string). Unless stated
otherwise, it is the empty string.

A [request](#concept-request){#ref-for-concept-request④⑥
link-type="dfn"} has associated [cryptographic nonce
metadata]{#concept-request-nonce-metadata .dfn .dfn-paneled
dfn-for="request" dfn-type="dfn" export=""} (a string). Unless stated
otherwise, it is the empty string.

A [request](#concept-request){#ref-for-concept-request④⑦
link-type="dfn"} has associated [parser
metadata]{#concept-request-parser-metadata .dfn .dfn-paneled
dfn-for="request" dfn-type="dfn" export=""} which is the empty string,
\"`parser-inserted`\", or \"`not-parser-inserted`\". Unless otherwise
stated, it is the empty string.

A [request](#concept-request){#ref-for-concept-request④⑧
link-type="dfn"}'s [cryptographic nonce
metadata](#concept-request-nonce-metadata){#ref-for-concept-request-nonce-metadata
link-type="dfn"} and [parser
metadata](#concept-request-parser-metadata){#ref-for-concept-request-parser-metadata
link-type="dfn"} are generally populated from attributes and flags on
the HTML element responsible for creating a
[request](#concept-request){#ref-for-concept-request④⑨ link-type="dfn"}.
They are used by various algorithms in Content Security Policy to
determine whether requests or responses are to be blocked in a given
context.
[\[CSP\]](#biblio-csp "Content Security Policy Level 3"){link-type="biblio"}

A [request](#concept-request){#ref-for-concept-request⑤⓪
link-type="dfn"} has an associated [reload-navigation
flag]{#concept-request-reload-navigation-flag .dfn .dfn-paneled
dfn-for="request" dfn-type="dfn" export=""}. Unless stated otherwise, it
is unset.

This flag is for exclusive use by HTML's navigate algorithm.
[\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}

A [request](#concept-request){#ref-for-concept-request⑤①
link-type="dfn"} has an associated [history-navigation
flag]{#concept-request-history-navigation-flag .dfn .dfn-paneled
dfn-for="request" dfn-type="dfn" export=""}. Unless stated otherwise, it
is unset.

This flag is for exclusive use by HTML's navigate algorithm.
[\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}

A [request](#concept-request){#ref-for-concept-request⑤②
link-type="dfn"} has an associated boolean
[user-activation]{#request-user-activation .dfn .dfn-paneled
dfn-for="request" dfn-type="dfn" export=""}. Unless stated otherwise, it
is false.

This is for exclusive use by HTML's navigate algorithm.
[\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}

A [request](#concept-request){#ref-for-concept-request⑤③
link-type="dfn"} has an associated boolean
[render-blocking]{#request-render-blocking .dfn .dfn-paneled
dfn-for="request" dfn-type="dfn" export=""}. Unless stated otherwise, it
is false.

This flag is for exclusive use by HTML's render-blocking mechanism.
[\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}

------------------------------------------------------------------------

A [request](#concept-request){#ref-for-concept-request⑤④
link-type="dfn"} has an associated [URL list]{#concept-request-url-list
.dfn .dfn-paneled dfn-for="request" dfn-type="dfn" export=""} (a
[list](https://infra.spec.whatwg.org/#list){#ref-for-list①⓪
link-type="dfn"} of one or more
[URLs](https://url.spec.whatwg.org/#concept-url){#ref-for-concept-url⑦
link-type="dfn"}). Unless stated otherwise, it is a list containing a
copy of [request](#concept-request){#ref-for-concept-request⑤⑤
link-type="dfn"}'s
[URL](#concept-request-url){#ref-for-concept-request-url①
link-type="dfn"}.

A [request](#concept-request){#ref-for-concept-request⑤⑥
link-type="dfn"} has an associated [current
URL]{#concept-request-current-url .dfn .dfn-paneled dfn-for="request"
dfn-type="dfn" export=""}. It is a pointer to the last
[URL](https://url.spec.whatwg.org/#concept-url){#ref-for-concept-url⑧
link-type="dfn"} in
[request](#concept-request){#ref-for-concept-request⑤⑦
link-type="dfn"}'s [URL
list](#concept-request-url-list){#ref-for-concept-request-url-list①
link-type="dfn"}.

A [request](#concept-request){#ref-for-concept-request⑤⑧
link-type="dfn"} has an associated [redirect
count]{#concept-request-redirect-count .dfn .dfn-paneled
dfn-for="request" dfn-type="dfn" export=""}. Unless stated otherwise, it
is zero.

A [request](#concept-request){#ref-for-concept-request⑤⑨
link-type="dfn"} has an associated [response
tainting]{#concept-request-response-tainting .dfn .dfn-paneled
dfn-for="request" dfn-type="dfn" export=""}, which is \"`basic`\",
\"`cors`\", or \"`opaque`\". Unless stated otherwise, it is \"`basic`\".

A [request](#concept-request){#ref-for-concept-request⑥⓪
link-type="dfn"} has an associated [prevent no-cache cache-control
header modification flag]{#no-cache-prevent-cache-control .dfn
.dfn-paneled dfn-for="request" dfn-type="dfn" export=""}. Unless stated
otherwise, it is unset.

A [request](#concept-request){#ref-for-concept-request⑥①
link-type="dfn"} has an associated [done flag]{#done-flag .dfn
.dfn-paneled dfn-for="request" dfn-type="dfn" export=""}. Unless stated
otherwise, it is unset.

A [request](#concept-request){#ref-for-concept-request⑥②
link-type="dfn"} has an associated [timing allow failed
flag]{#timing-allow-failed .dfn .dfn-paneled dfn-for="request"
dfn-type="dfn" export=""}. Unless stated otherwise, it is unset.

A [request](#concept-request){#ref-for-concept-request⑥③
link-type="dfn"}'s [URL
list](#concept-request-url-list){#ref-for-concept-request-url-list②
link-type="dfn"}, [current
URL](#concept-request-current-url){#ref-for-concept-request-current-url①
link-type="dfn"}, [redirect
count](#concept-request-redirect-count){#ref-for-concept-request-redirect-count
link-type="dfn"}, [response
tainting](#concept-request-response-tainting){#ref-for-concept-request-response-tainting①
link-type="dfn"}, [done flag](#done-flag){#ref-for-done-flag
link-type="dfn"}, and [timing allow failed
flag](#timing-allow-failed){#ref-for-timing-allow-failed
link-type="dfn"} are used as bookkeeping details by the
[fetch](#concept-fetch){#ref-for-concept-fetch①⑤ link-type="dfn"}
algorithm.

------------------------------------------------------------------------

A [subresource request]{#subresource-request .dfn .dfn-paneled
dfn-type="dfn" export=""} is a
[request](#concept-request){#ref-for-concept-request⑥④ link-type="dfn"}
whose
[destination](#concept-request-destination){#ref-for-concept-request-destination③
link-type="dfn"} is \"`audio`\", \"`audioworklet`\", \"`font`\",
\"`image`\", \"`json`\", \"`manifest`\", \"`paintworklet`\",
\"`script`\", \"`style`\", \"`track`\", \"`video`\", \"`xslt`\", or the
empty string.

A [non-subresource request]{#non-subresource-request .dfn .dfn-paneled
dfn-type="dfn" export=""} is a
[request](#concept-request){#ref-for-concept-request⑥⑤ link-type="dfn"}
whose
[destination](#concept-request-destination){#ref-for-concept-request-destination④
link-type="dfn"} is \"`document`\", \"`embed`\", \"`frame`\",
\"`iframe`\", \"`object`\", \"`report`\", \"`serviceworker`\",
\"`sharedworker`\", or \"`worker`\".

A [navigation request]{#navigation-request .dfn .dfn-paneled
dfn-type="dfn" export=""} is a
[request](#concept-request){#ref-for-concept-request⑥⑥ link-type="dfn"}
whose
[destination](#concept-request-destination){#ref-for-concept-request-destination⑤
link-type="dfn"} is \"`document`\", \"`embed`\", \"`frame`\",
\"`iframe`\", or \"`object`\".

See [handle
fetch](https://w3c.github.io/ServiceWorker/#handle-fetch){#ref-for-handle-fetch
link-type="dfn"} for usage of these terms.
[\[SW\]](#biblio-sw "Service Workers"){link-type="biblio"}

------------------------------------------------------------------------

::: {.algorithm algorithm="redirect-tainted origin" algorithm-for="request"}
A [request](#concept-request){#ref-for-concept-request⑥⑦
link-type="dfn"} `request`{.variable} has a [redirect-tainted
origin]{#concept-request-tainted-origin .dfn .dfn-paneled
dfn-for="request" dfn-type="dfn" noexport=""} if these steps return
true:

1.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert⑧
    link-type="dfn"}: `request`{.variable}'s
    [origin](#concept-request-origin){#ref-for-concept-request-origin①
    link-type="dfn"} is not \"`client`\".

2.  Let `lastURL`{.variable} be null.

3.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate⑥
    link-type="dfn"} `url`{.variable} of `request`{.variable}'s [URL
    list](#concept-request-url-list){#ref-for-concept-request-url-list③
    link-type="dfn"}:

    1.  If `lastURL`{.variable} is null, then set `lastURL`{.variable}
        to `url`{.variable} and
        [continue](https://infra.spec.whatwg.org/#iteration-continue){#ref-for-iteration-continue①
        link-type="dfn"}.

    2.  If `url`{.variable}'s
        [origin](https://url.spec.whatwg.org/#concept-url-origin){#ref-for-concept-url-origin
        link-type="dfn"} is not [same
        origin](https://html.spec.whatwg.org/multipage/browsers.html#same-origin){#ref-for-same-origin
        link-type="dfn"} with `lastURL`{.variable}'s
        [origin](https://url.spec.whatwg.org/#concept-url-origin){#ref-for-concept-url-origin①
        link-type="dfn"} and `request`{.variable}'s
        [origin](#concept-request-origin){#ref-for-concept-request-origin②
        link-type="dfn"} is not [same
        origin](https://html.spec.whatwg.org/multipage/browsers.html#same-origin){#ref-for-same-origin①
        link-type="dfn"} with `lastURL`{.variable}'s
        [origin](https://url.spec.whatwg.org/#concept-url-origin){#ref-for-concept-url-origin②
        link-type="dfn"}, then return true.

    3.  Set `lastURL`{.variable} to `url`{.variable}.

4.  Return false.
:::

::: {.algorithm algorithm="Serializing a request origin"}
[Serializing a request origin]{#serializing-a-request-origin .dfn
.dfn-paneled dfn-type="dfn" noexport=""}, given a
[request](#concept-request){#ref-for-concept-request⑥⑧ link-type="dfn"}
`request`{.variable}, is to run these steps:

1.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert⑨
    link-type="dfn"}: `request`{.variable}'s
    [origin](#concept-request-origin){#ref-for-concept-request-origin③
    link-type="dfn"} is not \"`client`\".

2.  If `request`{.variable} has a [redirect-tainted
    origin](#concept-request-tainted-origin){#ref-for-concept-request-tainted-origin
    link-type="dfn"}, then return \"`null`\".

3.  Return `request`{.variable}'s
    [origin](#concept-request-origin){#ref-for-concept-request-origin④
    link-type="dfn"},
    [serialized](https://html.spec.whatwg.org/multipage/browsers.html#ascii-serialisation-of-an-origin){#ref-for-ascii-serialisation-of-an-origin
    link-type="dfn"}.
:::

::: {.algorithm algorithm="Byte-serializing a request origin"}
[Byte-serializing a request origin]{#byte-serializing-a-request-origin
.dfn .dfn-paneled dfn-type="dfn" noexport=""}, given a
[request](#concept-request){#ref-for-concept-request⑥⑨ link-type="dfn"}
`request`{.variable}, is to return the result of [serializing a request
origin](#serializing-a-request-origin){#ref-for-serializing-a-request-origin
link-type="dfn"} with `request`{.variable}, [isomorphic
encoded](https://infra.spec.whatwg.org/#isomorphic-encode){#ref-for-isomorphic-encode④
link-type="dfn"}.
:::

------------------------------------------------------------------------

::: {.algorithm algorithm="clone" algorithm-for="request"}
To [clone]{#concept-request-clone .dfn .dfn-paneled dfn-for="request"
dfn-type="dfn" export=""} a
[request](#concept-request){#ref-for-concept-request⑦⓪ link-type="dfn"}
`request`{.variable}, run these steps:

1.  Let `newRequest`{.variable} be a copy of `request`{.variable},
    except for its
    [body](#concept-request-body){#ref-for-concept-request-body
    link-type="dfn"}.

2.  If `request`{.variable}'s
    [body](#concept-request-body){#ref-for-concept-request-body①
    link-type="dfn"} is non-null, set `newRequest`{.variable}'s
    [body](#concept-request-body){#ref-for-concept-request-body②
    link-type="dfn"} to the result of
    [cloning](#concept-body-clone){#ref-for-concept-body-clone
    link-type="dfn"} `request`{.variable}'s
    [body](#concept-request-body){#ref-for-concept-request-body③
    link-type="dfn"}.

3.  Return `newRequest`{.variable}.
:::

------------------------------------------------------------------------

::: {.algorithm algorithm="add a range header" algorithm-for="request"}
To [add a range header]{#concept-request-add-range-header .dfn
.dfn-paneled dfn-for="request" dfn-type="dfn" export=""} to a
[request](#concept-request){#ref-for-concept-request⑦① link-type="dfn"}
`request`{.variable}, with an integer `first`{.variable}, and an
optional integer `last`{.variable}, run these steps:

1.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert①⓪
    link-type="dfn"}: `last`{.variable} is not given, or
    `first`{.variable} is less than or equal to `last`{.variable}.

2.  Let `rangeValue`{.variable} be \``bytes=`\`.

3.  [Serialize](#serialize-an-integer){#ref-for-serialize-an-integer③
    link-type="dfn"} and [isomorphic
    encode](https://infra.spec.whatwg.org/#isomorphic-encode){#ref-for-isomorphic-encode⑤
    link-type="dfn"} `first`{.variable}, and append the result to
    `rangeValue`{.variable}.

4.  Append 0x2D (-) to `rangeValue`{.variable}.

5.  If `last`{.variable} is given, then
    [serialize](#serialize-an-integer){#ref-for-serialize-an-integer④
    link-type="dfn"} and [isomorphic
    encode](https://infra.spec.whatwg.org/#isomorphic-encode){#ref-for-isomorphic-encode⑥
    link-type="dfn"} it, and append the result to
    `rangeValue`{.variable}.

6.  [Append](#concept-header-list-append){#ref-for-concept-header-list-append
    link-type="dfn"} (\``Range`\`, `rangeValue`{.variable}) to
    `request`{.variable}'s [header
    list](#concept-request-header-list){#ref-for-concept-request-header-list②
    link-type="dfn"}.

A range header denotes an inclusive byte range. There a range header
where `first`{.variable} is 0 and `last`{.variable} is 500, is a range
of 501 bytes.

Features that combine multiple responses into one logical resource are
historically a source of security bugs. Please seek security review for
features that deal with partial responses.
:::

------------------------------------------------------------------------

::: {.algorithm algorithm="serialize a response URL for reporting"}
To [serialize a response URL for
reporting]{#serialize-a-response-url-for-reporting .dfn .dfn-paneled
dfn-type="dfn" export=""}, given a
[response](#concept-response){#ref-for-concept-response①
link-type="dfn"} `response`{.variable}, run these steps:

1.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert①①
    link-type="dfn"}: `response`{.variable}'s [URL
    list](#concept-response-url-list){#ref-for-concept-response-url-list
    link-type="dfn"} [is not
    empty](https://infra.spec.whatwg.org/#list-is-empty){#ref-for-list-is-empty
    link-type="dfn"}.

2.  Let `url`{.variable} be a copy of `response`{.variable}'s [URL
    list](#concept-response-url-list){#ref-for-concept-response-url-list①
    link-type="dfn"}\[0\].

    This is not `response`{.variable}'s
    [URL](#concept-response-url){#ref-for-concept-response-url
    link-type="dfn"} in order to avoid leaking information about
    redirect targets (see [similar considerations for CSP
    reporting](https://w3c.github.io/webappsec-csp/#security-violation-reports)
    too).
    [\[CSP\]](#biblio-csp "Content Security Policy Level 3"){link-type="biblio"}

3.  [Set the
    username](https://url.spec.whatwg.org/#set-the-username){#ref-for-set-the-username
    link-type="dfn"} given `url`{.variable} and the empty string.

4.  [Set the
    password](https://url.spec.whatwg.org/#set-the-password){#ref-for-set-the-password
    link-type="dfn"} given `url`{.variable} and the empty string.

5.  Return the
    [serialization](https://url.spec.whatwg.org/#concept-url-serializer){#ref-for-concept-url-serializer
    link-type="dfn"} of `url`{.variable} with [*exclude
    fragment*](https://url.spec.whatwg.org/#url-serializer-exclude-fragment){#ref-for-url-serializer-exclude-fragment
    link-type="dfn"} set to true.
:::

::: {.algorithm algorithm="Cross-Origin-Embedder-Policy allows credentials"}
To check if [Cross-Origin-Embedder-Policy allows
credentials]{#cross-origin-embedder-policy-allows-credentials .dfn
.dfn-paneled dfn-type="dfn" export=""}, given a
[request](#concept-request){#ref-for-concept-request⑦② link-type="dfn"}
`request`{.variable}, run these steps:

1.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert①②
    link-type="dfn"}: `request`{.variable}'s
    [origin](#concept-request-origin){#ref-for-concept-request-origin⑤
    link-type="dfn"} is not \"`client`\".

2.  If `request`{.variable}'s
    [mode](#concept-request-mode){#ref-for-concept-request-mode④
    link-type="dfn"} is not \"`no-cors`\", then return true.

3.  If `request`{.variable}'s
    [client](#concept-request-client){#ref-for-concept-request-client②
    link-type="dfn"} is null, then return true.

4.  If `request`{.variable}'s
    [client](#concept-request-client){#ref-for-concept-request-client③
    link-type="dfn"}'s [policy
    container](https://html.spec.whatwg.org/multipage/webappapis.html#concept-settings-object-policy-container){#ref-for-concept-settings-object-policy-container
    link-type="dfn"}'s [embedder
    policy](https://html.spec.whatwg.org/multipage/browsers.html#policy-container-embedder-policy){#ref-for-policy-container-embedder-policy
    link-type="dfn"}'s
    [value](https://html.spec.whatwg.org/multipage/browsers.html#embedder-policy-value-2){#ref-for-embedder-policy-value-2
    link-type="dfn"} is not
    \"[`credentialless`](https://html.spec.whatwg.org/multipage/browsers.html#coep-credentialless){#ref-for-coep-credentialless
    link-type="dfn"}\", then return true.

5.  If `request`{.variable}'s
    [origin](#concept-request-origin){#ref-for-concept-request-origin⑥
    link-type="dfn"} is [same
    origin](https://html.spec.whatwg.org/multipage/browsers.html#same-origin){#ref-for-same-origin②
    link-type="dfn"} with `request`{.variable}'s [current
    URL](#concept-request-current-url){#ref-for-concept-request-current-url②
    link-type="dfn"}'s
    [origin](https://url.spec.whatwg.org/#concept-url-origin){#ref-for-concept-url-origin③
    link-type="dfn"} and `request`{.variable} does not have a
    [redirect-tainted
    origin](#concept-request-tainted-origin){#ref-for-concept-request-tainted-origin①
    link-type="dfn"}, then return true.

6.  Return false.
:::

#### [2.2.6. ]{.secno}[Responses]{.content}[](#responses){.self-link} {#responses .heading .settled level="2.2.6"}

The result of [fetch](#concept-fetch){#ref-for-concept-fetch①⑥
link-type="dfn"} is a [response]{#concept-response .dfn .dfn-paneled
dfn-type="dfn" export=""}. A
[response](#concept-response){#ref-for-concept-response②
link-type="dfn"} evolves over time. That is, not all its fields are
available straight away.

A [response](#concept-response){#ref-for-concept-response③
link-type="dfn"} has an associated [type]{#concept-response-type .dfn
.dfn-paneled dfn-for="response" dfn-type="dfn" export=""} which is
\"`basic`\", \"`cors`\", \"`default`\", \"`error`\", \"`opaque`\", or
\"`opaqueredirect`\". Unless stated otherwise, it is \"`default`\".

A [response](#concept-response){#ref-for-concept-response④
link-type="dfn"} can have an associated [aborted
flag]{#concept-response-aborted .dfn .dfn-paneled dfn-for="response"
dfn-type="dfn" export=""}, which is initially unset.

This indicates that the request was intentionally aborted by the
developer or end-user.

A [response](#concept-response){#ref-for-concept-response⑤
link-type="dfn"} has an associated [URL]{#concept-response-url .dfn
.dfn-paneled dfn-for="response" dfn-type="dfn" export=""}. It is a
pointer to the last
[URL](https://url.spec.whatwg.org/#concept-url){#ref-for-concept-url⑨
link-type="dfn"} in
[response](#concept-response){#ref-for-concept-response⑥
link-type="dfn"}'s [URL
list](#concept-response-url-list){#ref-for-concept-response-url-list②
link-type="dfn"} and null if
[response](#concept-response){#ref-for-concept-response⑦
link-type="dfn"}'s [URL
list](#concept-response-url-list){#ref-for-concept-response-url-list③
link-type="dfn"} [is
empty](https://infra.spec.whatwg.org/#list-is-empty){#ref-for-list-is-empty①
link-type="dfn"}.

A [response](#concept-response){#ref-for-concept-response⑧
link-type="dfn"} has an associated [URL list]{#concept-response-url-list
.dfn .dfn-paneled dfn-for="response" dfn-type="dfn" export=""} (a
[list](https://infra.spec.whatwg.org/#list){#ref-for-list①①
link-type="dfn"} of zero or more
[URLs](https://url.spec.whatwg.org/#concept-url){#ref-for-concept-url①⓪
link-type="dfn"}). Unless stated otherwise, it is « ».

Except for the first and last
[URL](https://url.spec.whatwg.org/#concept-url){#ref-for-concept-url①①
link-type="dfn"}, if any, a
[response](#concept-response){#ref-for-concept-response⑨
link-type="dfn"}'s [URL
list](#concept-response-url-list){#ref-for-concept-response-url-list④
link-type="dfn"} is not directly exposed to script as that would violate
[atomic HTTP redirect
handling](#atomic-http-redirect-handling){#ref-for-atomic-http-redirect-handling①
link-type="dfn"}.

A [response](#concept-response){#ref-for-concept-response①⓪
link-type="dfn"} has an associated [status]{#concept-response-status
.dfn .dfn-paneled dfn-for="response" dfn-type="dfn" export=""}, which is
a [status](#concept-status){#ref-for-concept-status③ link-type="dfn"}.
Unless stated otherwise it is 200.

A [response](#concept-response){#ref-for-concept-response①①
link-type="dfn"} has an associated [status
message]{#concept-response-status-message .dfn .dfn-paneled
dfn-for="response" dfn-type="dfn" export=""}. Unless stated otherwise it
is the empty byte sequence.

Responses over an HTTP/2 connection will always have the empty byte
sequence as status message as HTTP/2 does not support them.

A [response](#concept-response){#ref-for-concept-response①②
link-type="dfn"} has an associated [header
list]{#concept-response-header-list .dfn .dfn-paneled dfn-for="response"
dfn-type="dfn" export=""} (a [header
list](#concept-header-list){#ref-for-concept-header-list①⑧
link-type="dfn"}). Unless stated otherwise it is « ».

A [response](#concept-response){#ref-for-concept-response①③
link-type="dfn"} has an associated [body]{#concept-response-body .dfn
.dfn-paneled dfn-for="response" dfn-type="dfn" export=""} (null or a
[body](#concept-body){#ref-for-concept-body⑦ link-type="dfn"}). Unless
stated otherwise it is null.

The [source](#concept-body-source){#ref-for-concept-body-source
link-type="dfn"} and
[length](#concept-body-total-bytes){#ref-for-concept-body-total-bytes
link-type="dfn"} concepts of a network's
[response](#concept-response){#ref-for-concept-response①④
link-type="dfn"}'s
[body](#concept-response-body){#ref-for-concept-response-body
link-type="dfn"} are always null.

A [response](#concept-response){#ref-for-concept-response①⑤
link-type="dfn"} has an associated [cache
state]{#concept-response-cache-state .dfn .dfn-paneled
dfn-for="response" dfn-type="dfn" export=""} (the empty string,
\"`local`\", or \"`validated`\"). Unless stated otherwise, it is the
empty string.

This is intended for usage by Service Workers and Resource Timing.
[\[SW\]](#biblio-sw "Service Workers"){link-type="biblio"}
[\[RESOURCE-TIMING\]](#biblio-resource-timing "Resource Timing"){link-type="biblio"}

A [response](#concept-response){#ref-for-concept-response①⑥
link-type="dfn"} has an associated [CORS-exposed header-name
list]{#concept-response-cors-exposed-header-name-list .dfn .dfn-paneled
dfn-for="response" dfn-type="dfn" export=""} (a list of zero or more
[header](#concept-header){#ref-for-concept-header②⑤ link-type="dfn"}
[names](#concept-header-name){#ref-for-concept-header-name①④
link-type="dfn"}). The list is empty unless otherwise specified.

A [response](#concept-response){#ref-for-concept-response①⑦
link-type="dfn"} will typically get its [CORS-exposed header-name
list](#concept-response-cors-exposed-header-name-list){#ref-for-concept-response-cors-exposed-header-name-list
link-type="dfn"} set by [extracting header
values](#extract-header-values){#ref-for-extract-header-values①
link-type="dfn"} from the
\`[`Access-Control-Expose-Headers`](#http-access-control-expose-headers){#ref-for-http-access-control-expose-headers
link-type="http-header"}\` header. This list is used by a [CORS filtered
response](#concept-filtered-response-cors){#ref-for-concept-filtered-response-cors
link-type="dfn"} to determine which headers to expose.

A [response](#concept-response){#ref-for-concept-response①⑧
link-type="dfn"} has an associated [range-requested
flag]{#concept-response-range-requested-flag .dfn .dfn-paneled
dfn-for="response" dfn-type="dfn" noexport=""}, which is initially
unset.

This is used to prevent a partial response from an earlier ranged
request being provided to an API that didn't make a range request. See
the flag's usage for a detailed description of the attack.

A [response](#concept-response){#ref-for-concept-response①⑨
link-type="dfn"} has an associated
[request-includes-credentials]{#response-request-includes-credentials
.dfn .dfn-paneled dfn-for="response" dfn-type="dfn" noexport=""} (a
boolean), which is initially true.

A [response](#concept-response){#ref-for-concept-response②⓪
link-type="dfn"} has an associated [timing allow passed
flag]{#concept-response-timing-allow-passed .dfn .dfn-paneled
dfn-for="response" dfn-type="dfn" noexport=""}, which is initially
unset.

This is used so that the caller to a fetch can determine if sensitive
timing data is allowed on the resource fetched by looking at the flag of
the response returned. Because the flag on the response of a redirect
has to be set if it was set for previous responses in the redirect
chain, this is also tracked internally using the request's [timing allow
failed flag](#timing-allow-failed){#ref-for-timing-allow-failed①
link-type="dfn"}.

A [response](#concept-response){#ref-for-concept-response②①
link-type="dfn"} has an associated [body
info]{#concept-response-body-info .dfn .dfn-paneled dfn-for="response"
dfn-type="dfn" export=""} (a [response body
info](#response-body-info){#ref-for-response-body-info
link-type="dfn"}). Unless stated otherwise, it is a new [response body
info](#response-body-info){#ref-for-response-body-info①
link-type="dfn"}.

A [response](#concept-response){#ref-for-concept-response②②
link-type="dfn"} has an associated [service worker timing
info]{#response-service-worker-timing-info .dfn .dfn-paneled
dfn-for="response" dfn-type="dfn" export=""} (null or a [service worker
timing
info](https://w3c.github.io/ServiceWorker/#service-worker-timing-info){#ref-for-service-worker-timing-info
link-type="dfn"}), which is initially null.

A [response](#concept-response){#ref-for-concept-response②③
link-type="dfn"} has an associated
[has-cross-origin-redirects]{#response-has-cross-origin-redirects .dfn
.dfn-paneled dfn-for="response" dfn-type="dfn" noexport=""} (a boolean),
which is initially false.

------------------------------------------------------------------------

A [network error]{#concept-network-error .dfn .dfn-paneled
dfn-type="dfn" export=""} is a
[response](#concept-response){#ref-for-concept-response②④
link-type="dfn"} whose
[type](#concept-response-type){#ref-for-concept-response-type
link-type="dfn"} is \"`error`\",
[status](#concept-response-status){#ref-for-concept-response-status
link-type="dfn"} is 0, [status
message](#concept-response-status-message){#ref-for-concept-response-status-message
link-type="dfn"} is the empty byte sequence, [header
list](#concept-response-header-list){#ref-for-concept-response-header-list
link-type="dfn"} is « »,
[body](#concept-response-body){#ref-for-concept-response-body①
link-type="dfn"} is null, and [body
info](#concept-response-body-info){#ref-for-concept-response-body-info
link-type="dfn"} is a new [response body
info](#response-body-info){#ref-for-response-body-info②
link-type="dfn"}.

An [aborted network error]{#concept-aborted-network-error .dfn
.dfn-paneled dfn-type="dfn" export=""} is a [network
error](#concept-network-error){#ref-for-concept-network-error⑤
link-type="dfn"} whose [aborted
flag](#concept-response-aborted){#ref-for-concept-response-aborted
link-type="dfn"} is set.

::: {.algorithm algorithm="appropriate network error"}
To create the [appropriate network error]{#appropriate-network-error
.dfn .dfn-paneled dfn-type="dfn" noexport=""} given [fetch
params](#fetch-params){#ref-for-fetch-params② link-type="dfn"}
`fetchParams`{.variable}:

1.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert①③
    link-type="dfn"}: `fetchParams`{.variable} is
    [canceled](#fetch-params-canceled){#ref-for-fetch-params-canceled
    link-type="dfn"}.

2.  Return an [aborted network
    error](#concept-aborted-network-error){#ref-for-concept-aborted-network-error
    link-type="dfn"} if `fetchParams`{.variable} is
    [aborted](#fetch-params-aborted){#ref-for-fetch-params-aborted
    link-type="dfn"}; otherwise return a [network
    error](#concept-network-error){#ref-for-concept-network-error⑥
    link-type="dfn"}.
:::

------------------------------------------------------------------------

A [filtered response]{#concept-filtered-response .dfn .dfn-paneled
dfn-type="dfn" export=""} is a
[response](#concept-response){#ref-for-concept-response②⑤
link-type="dfn"} that offers a limited view on an associated
[response](#concept-response){#ref-for-concept-response②⑥
link-type="dfn"}. This associated
[response](#concept-response){#ref-for-concept-response②⑦
link-type="dfn"} can be accessed through [filtered
response](#concept-filtered-response){#ref-for-concept-filtered-response
link-type="dfn"}'s [internal response]{#concept-internal-response .dfn
.dfn-paneled dfn-for="filtered response" dfn-type="dfn" export=""} (a
[response](#concept-response){#ref-for-concept-response②⑧
link-type="dfn"} that is neither a [network
error](#concept-network-error){#ref-for-concept-network-error⑦
link-type="dfn"} nor a [filtered
response](#concept-filtered-response){#ref-for-concept-filtered-response①
link-type="dfn"}).

Unless stated otherwise a [filtered
response](#concept-filtered-response){#ref-for-concept-filtered-response②
link-type="dfn"}'s associated concepts (such as its
[body](#concept-response-body){#ref-for-concept-response-body②
link-type="dfn"}) refer to the associated concepts of its [internal
response](#concept-internal-response){#ref-for-concept-internal-response
link-type="dfn"}. (The exceptions to this are listed below as part of
defining the concrete types of [filtered
responses](#concept-filtered-response){#ref-for-concept-filtered-response③
link-type="dfn"}.)

::: {.note role="note"}
The [fetch](#concept-fetch){#ref-for-concept-fetch①⑦ link-type="dfn"}
algorithm by way of
[*processResponse*](#process-response){#ref-for-process-response
link-type="dfn"} and equivalent parameters exposes [filtered
responses](#concept-filtered-response){#ref-for-concept-filtered-response④
link-type="dfn"} to callers to ensure they do not accidentally leak
information. If the information needs to be revealed for legacy reasons,
e.g., to feed image data to a decoder, the associated [internal
response](#concept-internal-response){#ref-for-concept-internal-response①
link-type="dfn"} can be used by specification algorithms.

New specifications ought not to build further on [opaque filtered
responses](#concept-filtered-response-opaque){#ref-for-concept-filtered-response-opaque①
link-type="dfn"} or [opaque-redirect filtered
responses](#concept-filtered-response-opaque-redirect){#ref-for-concept-filtered-response-opaque-redirect①
link-type="dfn"}. Those are legacy constructs and cannot always be
adequately protected given contemporary computer architecture.
:::

A [basic filtered response]{#concept-filtered-response-basic .dfn
.dfn-paneled dfn-type="dfn" export=""} is a [filtered
response](#concept-filtered-response){#ref-for-concept-filtered-response⑤
link-type="dfn"} whose
[type](#concept-response-type){#ref-for-concept-response-type①
link-type="dfn"} is \"`basic`\" and [header
list](#concept-response-header-list){#ref-for-concept-response-header-list①
link-type="dfn"} excludes any
[headers](#concept-header){#ref-for-concept-header②⑥ link-type="dfn"} in
[internal
response](#concept-internal-response){#ref-for-concept-internal-response②
link-type="dfn"}'s [header
list](#concept-response-header-list){#ref-for-concept-response-header-list②
link-type="dfn"} whose
[name](#concept-header-name){#ref-for-concept-header-name①⑤
link-type="dfn"} is a [forbidden response-header
name](#forbidden-response-header-name){#ref-for-forbidden-response-header-name①
link-type="dfn"}.

A [CORS filtered response]{#concept-filtered-response-cors .dfn
.dfn-paneled dfn-type="dfn" export=""} is a [filtered
response](#concept-filtered-response){#ref-for-concept-filtered-response⑥
link-type="dfn"} whose
[type](#concept-response-type){#ref-for-concept-response-type②
link-type="dfn"} is \"`cors`\" and [header
list](#concept-response-header-list){#ref-for-concept-response-header-list③
link-type="dfn"} excludes any
[headers](#concept-header){#ref-for-concept-header②⑦ link-type="dfn"} in
[internal
response](#concept-internal-response){#ref-for-concept-internal-response③
link-type="dfn"}'s [header
list](#concept-response-header-list){#ref-for-concept-response-header-list④
link-type="dfn"} whose
[name](#concept-header-name){#ref-for-concept-header-name①⑥
link-type="dfn"} is *not* a [CORS-safelisted response-header
name](#cors-safelisted-response-header-name){#ref-for-cors-safelisted-response-header-name
link-type="dfn"}, given [internal
response](#concept-internal-response){#ref-for-concept-internal-response④
link-type="dfn"}'s [CORS-exposed header-name
list](#concept-response-cors-exposed-header-name-list){#ref-for-concept-response-cors-exposed-header-name-list①
link-type="dfn"}.

An [opaque filtered response]{#concept-filtered-response-opaque .dfn
.dfn-paneled dfn-type="dfn" export=""} is a [filtered
response](#concept-filtered-response){#ref-for-concept-filtered-response⑦
link-type="dfn"} whose
[type](#concept-response-type){#ref-for-concept-response-type③
link-type="dfn"} is \"`opaque`\", [URL
list](#concept-response-url-list){#ref-for-concept-response-url-list⑤
link-type="dfn"} is « »,
[status](#concept-response-status){#ref-for-concept-response-status①
link-type="dfn"} is 0, [status
message](#concept-response-status-message){#ref-for-concept-response-status-message①
link-type="dfn"} is the empty byte sequence, [header
list](#concept-response-header-list){#ref-for-concept-response-header-list⑤
link-type="dfn"} is « »,
[body](#concept-response-body){#ref-for-concept-response-body③
link-type="dfn"} is null, and [body
info](#concept-response-body-info){#ref-for-concept-response-body-info①
link-type="dfn"} is a new [response body
info](#response-body-info){#ref-for-response-body-info③
link-type="dfn"}.

An [opaque-redirect filtered
response]{#concept-filtered-response-opaque-redirect .dfn .dfn-paneled
dfn-type="dfn" export=""} is a [filtered
response](#concept-filtered-response){#ref-for-concept-filtered-response⑧
link-type="dfn"} whose
[type](#concept-response-type){#ref-for-concept-response-type④
link-type="dfn"} is \"`opaqueredirect`\",
[status](#concept-response-status){#ref-for-concept-response-status②
link-type="dfn"} is 0, [status
message](#concept-response-status-message){#ref-for-concept-response-status-message②
link-type="dfn"} is the empty byte sequence, [header
list](#concept-response-header-list){#ref-for-concept-response-header-list⑥
link-type="dfn"} is « »,
[body](#concept-response-body){#ref-for-concept-response-body④
link-type="dfn"} is null, and [body
info](#concept-response-body-info){#ref-for-concept-response-body-info②
link-type="dfn"} is a new [response body
info](#response-body-info){#ref-for-response-body-info④
link-type="dfn"}.

::: {.note role="note"}
Exposing the [URL
list](#concept-response-url-list){#ref-for-concept-response-url-list⑥
link-type="dfn"} for [opaque-redirect filtered
responses](#concept-filtered-response-opaque-redirect){#ref-for-concept-filtered-response-opaque-redirect②
link-type="dfn"} is harmless since no redirects are followed.

In other words, an [opaque filtered
response](#concept-filtered-response-opaque){#ref-for-concept-filtered-response-opaque②
link-type="dfn"} and an [opaque-redirect filtered
response](#concept-filtered-response-opaque-redirect){#ref-for-concept-filtered-response-opaque-redirect③
link-type="dfn"} are nearly indistinguishable from a [network
error](#concept-network-error){#ref-for-concept-network-error⑧
link-type="dfn"}. When introducing new APIs, do not use the [internal
response](#concept-internal-response){#ref-for-concept-internal-response⑤
link-type="dfn"} for internal specification algorithms as that will leak
information.

This also means that JavaScript APIs, such as
[`response.ok`](#dom-response-ok){#ref-for-dom-response-ok .idl-code
link-type="attribute"}, will return rather useless results.
:::

::: {#example-filtered-responses .example}
[](#example-filtered-responses){.self-link}

The [type](#concept-response-type){#ref-for-concept-response-type⑤
link-type="dfn"} of a
[response](#concept-response){#ref-for-concept-response②⑨
link-type="dfn"} is exposed to script through the
[`type`{.idl}](#dom-response-type){#ref-for-dom-response-type
link-type="idl"} getter:

``` {.lang-javascript .highlight}
console.log(new Response().type); // "default"

console.log((await fetch("/")).type); // "basic"

console.log((await fetch("https://api.example/status")).type); // "cors"

console.log((await fetch("https://crossorigin.example/image", { mode: "no-cors" })).type); // "opaque"

console.log((await fetch("/surprise-me", { redirect: "manual" })).type); // "opaqueredirect"
```

(This assumes that the various resources exist,
`https://api.example/status` has the appropriate CORS headers, and
`/surprise-me` uses a [redirect
status](#redirect-status){#ref-for-redirect-status link-type="dfn"}.)
:::

------------------------------------------------------------------------

::: {.algorithm algorithm="clone" algorithm-for="response"}
To [clone]{#concept-response-clone .dfn .dfn-paneled dfn-for="response"
dfn-type="dfn" export=""} a
[response](#concept-response){#ref-for-concept-response③⓪
link-type="dfn"} `response`{.variable}, run these steps:

1.  If `response`{.variable} is a [filtered
    response](#concept-filtered-response){#ref-for-concept-filtered-response⑨
    link-type="dfn"}, then return a new identical [filtered
    response](#concept-filtered-response){#ref-for-concept-filtered-response①⓪
    link-type="dfn"} whose [internal
    response](#concept-internal-response){#ref-for-concept-internal-response⑥
    link-type="dfn"} is a
    [clone](#concept-response-clone){#ref-for-concept-response-clone
    link-type="dfn"} of `response`{.variable}'s [internal
    response](#concept-internal-response){#ref-for-concept-internal-response⑦
    link-type="dfn"}.

2.  Let `newResponse`{.variable} be a copy of `response`{.variable},
    except for its
    [body](#concept-response-body){#ref-for-concept-response-body⑤
    link-type="dfn"}.

3.  If `response`{.variable}'s
    [body](#concept-response-body){#ref-for-concept-response-body⑥
    link-type="dfn"} is non-null, then set `newResponse`{.variable}'s
    [body](#concept-response-body){#ref-for-concept-response-body⑦
    link-type="dfn"} to the result of
    [cloning](#concept-body-clone){#ref-for-concept-body-clone①
    link-type="dfn"} `response`{.variable}'s
    [body](#concept-response-body){#ref-for-concept-response-body⑧
    link-type="dfn"}.

4.  Return `newResponse`{.variable}.
:::

------------------------------------------------------------------------

A [fresh response]{#concept-fresh-response .dfn .dfn-paneled
dfn-type="dfn" noexport=""} is a
[response](#concept-response){#ref-for-concept-response③①
link-type="dfn"} whose [current
age](https://httpwg.org/specs/rfc9111.html#age.calculations){#ref-for-age.calculations
link-type="dfn"} is within its [freshness
lifetime](https://httpwg.org/specs/rfc9111.html#calculating.freshness.lifetime){#ref-for-calculating.freshness.lifetime
link-type="dfn"}.

A [stale-while-revalidate
response]{#concept-stale-while-revalidate-response .dfn .dfn-paneled
dfn-type="dfn" noexport=""} is a
[response](#concept-response){#ref-for-concept-response③②
link-type="dfn"} that is not a [fresh
response](#concept-fresh-response){#ref-for-concept-fresh-response①
link-type="dfn"} and whose [current
age](https://httpwg.org/specs/rfc9111.html#age.calculations){#ref-for-age.calculations①
link-type="dfn"} is within the [stale-while-revalidate
lifetime](https://httpwg.org/specs/rfc5861.html#n-the-stale-while-revalidate-cache-control-extension){#ref-for-n-the-stale-while-revalidate-cache-control-extension
link-type="dfn"}.
[\[HTTP-CACHING\]](#biblio-http-caching "HTTP Caching"){link-type="biblio"}
[\[STALE-WHILE-REVALIDATE\]](#biblio-stale-while-revalidate "HTTP Cache-Control Extensions for Stale Content"){link-type="biblio"}

A [stale response]{#concept-stale-response .dfn .dfn-paneled
dfn-type="dfn" export=""} is a
[response](#concept-response){#ref-for-concept-response③③
link-type="dfn"} that is not a [fresh
response](#concept-fresh-response){#ref-for-concept-fresh-response②
link-type="dfn"} or a [stale-while-revalidate
response](#concept-stale-while-revalidate-response){#ref-for-concept-stale-while-revalidate-response①
link-type="dfn"}.

------------------------------------------------------------------------

::: {.algorithm algorithm="location URL" algorithm-for="response"}
The [location URL]{#concept-response-location-url .dfn .dfn-paneled
dfn-for="response" dfn-type="dfn" export=""} of a
[response](#concept-response){#ref-for-concept-response③④
link-type="dfn"} `response`{.variable}, given null or an [ASCII
string](https://infra.spec.whatwg.org/#ascii-string){#ref-for-ascii-string②
link-type="dfn"} `requestFragment`{.variable}, is the value returned by
the following steps. They return null, failure, or a
[URL](https://url.spec.whatwg.org/#concept-url){#ref-for-concept-url①②
link-type="dfn"}.

1.  If `response`{.variable}'s
    [status](#concept-response-status){#ref-for-concept-response-status③
    link-type="dfn"} is not a [redirect
    status](#redirect-status){#ref-for-redirect-status①
    link-type="dfn"}, then return null.

2.  Let `location`{.variable} be the result of [extracting header list
    values](#extract-header-list-values){#ref-for-extract-header-list-values
    link-type="dfn"} given \``Location`\` and `response`{.variable}'s
    [header
    list](#concept-response-header-list){#ref-for-concept-response-header-list⑦
    link-type="dfn"}.

3.  If `location`{.variable} is a [header
    value](#header-value){#ref-for-header-value⑦ link-type="dfn"}, then
    set `location`{.variable} to the result of
    [parsing](https://url.spec.whatwg.org/#concept-url-parser){#ref-for-concept-url-parser
    link-type="dfn"} `location`{.variable} with `response`{.variable}'s
    [URL](#concept-response-url){#ref-for-concept-response-url①
    link-type="dfn"}.

    If `response`{.variable} was constructed through the
    [`Response`{.idl}](#response){#ref-for-response link-type="idl"}
    constructor, `response`{.variable}'s
    [URL](#concept-response-url){#ref-for-concept-response-url②
    link-type="dfn"} will be null, meaning that `location`{.variable}
    will only parse successfully if it is an [absolute-URL-with-fragment
    string](https://url.spec.whatwg.org/#absolute-url-with-fragment-string){#ref-for-absolute-url-with-fragment-string
    link-type="dfn"}.

4.  If `location`{.variable} is a
    [URL](https://url.spec.whatwg.org/#concept-url){#ref-for-concept-url①③
    link-type="dfn"} whose
    [fragment](https://url.spec.whatwg.org/#concept-url-fragment){#ref-for-concept-url-fragment
    link-type="dfn"} is null, then set `location`{.variable}'s
    [fragment](https://url.spec.whatwg.org/#concept-url-fragment){#ref-for-concept-url-fragment①
    link-type="dfn"} to `requestFragment`{.variable}.

    This ensures that synthetic (indeed, all) responses follow the
    processing model for redirects defined by HTTP.
    [\[HTTP\]](#biblio-http "HTTP Semantics"){link-type="biblio"}

5.  Return `location`{.variable}.

The [location
URL](#concept-response-location-url){#ref-for-concept-response-location-url
link-type="dfn"} algorithm is exclusively used for redirect handling in
this standard and in HTML's navigate algorithm which handles redirects
manually. [\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}
:::

#### [2.2.7. ]{.secno}[Miscellaneous]{.content}[](#miscellaneous){.self-link} {#miscellaneous .heading .settled level="2.2.7"}

A [potential destination]{#concept-potential-destination .dfn
.dfn-paneled dfn-type="dfn" export=""} is \"`fetch`\" or a
[destination](#concept-request-destination){#ref-for-concept-request-destination⑥
link-type="dfn"} which is not the empty string.

::: {.algorithm algorithm="translate" algorithm-for="destination"}
To [translate]{#concept-potential-destination-translate .dfn
.dfn-paneled dfn-for="destination" dfn-type="dfn" export=""} a
[potential
destination](#concept-potential-destination){#ref-for-concept-potential-destination
link-type="dfn"} `potentialDestination`{.variable}, run these steps:

1.  If `potentialDestination`{.variable} is \"`fetch`\", then return the
    empty string.

2.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert①④
    link-type="dfn"}: `potentialDestination`{.variable} is a
    [destination](#concept-request-destination){#ref-for-concept-request-destination⑦
    link-type="dfn"}.

3.  Return `potentialDestination`{.variable}.
:::

### [2.3. ]{.secno}[Authentication entries]{.content}[](#authentication-entries){.self-link} {#authentication-entries .heading .settled level="2.3"}

An [authentication entry]{#authentication-entry .dfn .dfn-paneled
dfn-type="dfn" export=""} and a [proxy-authentication
entry]{#proxy-authentication-entry .dfn .dfn-paneled dfn-type="dfn"
export=""} are tuples of username, password, and realm, used for HTTP
authentication and HTTP proxy authentication, and associated with one or
more [requests](#concept-request){#ref-for-concept-request⑦③
link-type="dfn"}.

User agents should allow both to be cleared together with HTTP cookies
and similar tracking functionality.

Further details are defined by HTTP.
[\[HTTP\]](#biblio-http "HTTP Semantics"){link-type="biblio"}
[\[HTTP-CACHING\]](#biblio-http-caching "HTTP Caching"){link-type="biblio"}

### [2.4. ]{.secno}[Fetch groups]{.content}[](#fetch-groups){.self-link} {#fetch-groups .heading .settled level="2.4"}

Each [environment settings
object](https://html.spec.whatwg.org/multipage/webappapis.html#environment-settings-object){#ref-for-environment-settings-object⑤
link-type="dfn"} has an associated [fetch group]{#concept-fetch-group
.dfn .dfn-paneled dfn-for="fetch" dfn-type="dfn" export=""}.

A [fetch group](#concept-fetch-group){#ref-for-concept-fetch-group
link-type="dfn"} holds an ordered list of [fetch
records]{#concept-fetch-record .dfn .dfn-paneled dfn-for="fetch group"
dfn-type="dfn" export="" lt="fetch record"}.

A [fetch record](#concept-fetch-record){#ref-for-concept-fetch-record
link-type="dfn"} has an associated
[request]{#concept-fetch-record-request .dfn .dfn-paneled
dfn-for="fetch record" dfn-type="dfn" export=""} (a
[request](#concept-request){#ref-for-concept-request⑦④
link-type="dfn"}).

A [fetch record](#concept-fetch-record){#ref-for-concept-fetch-record①
link-type="dfn"} has an associated
[controller]{#concept-fetch-record-fetch .dfn .dfn-paneled
dfn-for="fetch record" dfn-type="dfn" export=""} (a [fetch
controller](#fetch-controller){#ref-for-fetch-controller⑦
link-type="dfn"} or null).

------------------------------------------------------------------------

When a [fetch group](#concept-fetch-group){#ref-for-concept-fetch-group①
link-type="dfn"} is [terminated]{#concept-fetch-group-terminate .dfn
.dfn-paneled dfn-for="fetch group" dfn-type="dfn" export=""}, for each
associated [fetch
record](#concept-fetch-record){#ref-for-concept-fetch-record②
link-type="dfn"} whose [fetch
record](#concept-fetch-record){#ref-for-concept-fetch-record③
link-type="dfn"}'s
[controller](#concept-fetch-record-fetch){#ref-for-concept-fetch-record-fetch
link-type="dfn"} is non-null, and whose
[request](#concept-fetch-record-request){#ref-for-concept-fetch-record-request
link-type="dfn"}'s [done flag](#done-flag){#ref-for-done-flag①
link-type="dfn"} is unset or
[keepalive](#request-keepalive-flag){#ref-for-request-keepalive-flag
link-type="dfn"} is false,
[terminate](#fetch-controller-terminate){#ref-for-fetch-controller-terminate
link-type="dfn"} the [fetch
record](#concept-fetch-record){#ref-for-concept-fetch-record④
link-type="dfn"}'s
[controller](#concept-fetch-record-fetch){#ref-for-concept-fetch-record-fetch①
link-type="dfn"}.

### [2.5. ]{.secno}[Resolving domains]{.content}[](#resolving-domains){.self-link} {#resolving-domains .heading .settled level="2.5"}

:::: {.algorithm algorithm="resolve an origin"}
[![(This is a tracking
vector.)](https://resources.whatwg.org/tracking-vector.svg "There is a tracking vector here."){.darkmode-aware
crossorigin="" height="64"
width="46"}](https://infra.spec.whatwg.org/#tracking-vector){.tracking-vector
style="color: currentcolor"} To [[]{#resolve-a-domain .bs-old-id}resolve
an origin]{#resolve-an-origin .dfn .dfn-paneled dfn-type="dfn"
export=""}, given a [network partition
key](#network-partition-key){#ref-for-network-partition-key
link-type="dfn"} `key`{.variable} and an
[origin](https://html.spec.whatwg.org/multipage/browsers.html#concept-origin){#ref-for-concept-origin②
link-type="dfn"} `origin`{.variable}:

1.  If `origin`{.variable}'s
    [host](https://html.spec.whatwg.org/multipage/browsers.html#concept-origin-host){#ref-for-concept-origin-host
    link-type="dfn"} is an [IP
    address](https://url.spec.whatwg.org/#ip-address){#ref-for-ip-address
    link-type="dfn"}, then return « `origin`{.variable}'s
    [host](https://html.spec.whatwg.org/multipage/browsers.html#concept-origin-host){#ref-for-concept-origin-host①
    link-type="dfn"} ».

2.  If `origin`{.variable}'s
    [host](https://html.spec.whatwg.org/multipage/browsers.html#concept-origin-host){#ref-for-concept-origin-host②
    link-type="dfn"}'s [public
    suffix](https://url.spec.whatwg.org/#host-public-suffix){#ref-for-host-public-suffix
    link-type="dfn"} is \"`localhost`\" or \"`localhost.`\", then return
    « `::1`, `127.0.0.1` ».

3.  Perform an
    [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#ref-for-implementation-defined②
    link-type="dfn"} operation to turn `origin`{.variable} into a
    [set](https://infra.spec.whatwg.org/#ordered-set){#ref-for-ordered-set②
    link-type="dfn"} of one or more [IP
    addresses](https://url.spec.whatwg.org/#ip-address){#ref-for-ip-address①
    link-type="dfn"}.

    It is also
    [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#ref-for-implementation-defined③
    link-type="dfn"} whether other operations might be performed to get
    connection information beyond just [IP
    addresses](https://url.spec.whatwg.org/#ip-address){#ref-for-ip-address②
    link-type="dfn"}. For example, if `origin`{.variable}'s
    [scheme](https://html.spec.whatwg.org/multipage/browsers.html#concept-origin-scheme){#ref-for-concept-origin-scheme
    link-type="dfn"} is an [HTTP(S)
    scheme](#http-scheme){#ref-for-http-scheme② link-type="dfn"}, the
    implementation might perform a DNS query for HTTPS RRs.
    [\[SVCB\]](#biblio-svcb "Service binding and parameter specification via the DNS (DNS SVCB and HTTPS RRs)"){link-type="biblio"}

    If this operation succeeds, return the
    [set](https://infra.spec.whatwg.org/#ordered-set){#ref-for-ordered-set③
    link-type="dfn"} of [IP
    addresses](https://url.spec.whatwg.org/#ip-address){#ref-for-ip-address③
    link-type="dfn"} and any additional
    [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#ref-for-implementation-defined④
    link-type="dfn"} information.

4.  Return failure.

The results of [resolve an
origin](#resolve-an-origin){#ref-for-resolve-an-origin link-type="dfn"}
may be cached. If they are cached, `key`{.variable} should be used as
part of the cache key.

::: {.note role="note"}
Typically this operation would involve DNS and as such caching can
happen on DNS servers without `key`{.variable} being taken into account.
Depending on the implementation it might also not be possible to take
`key`{.variable} into account locally.
[\[RFC1035\]](#biblio-rfc1035 "Domain names - implementation and specification"){link-type="biblio"}

The order of the [IP
addresses](https://url.spec.whatwg.org/#ip-address){#ref-for-ip-address④
link-type="dfn"} that the [resolve an
origin](#resolve-an-origin){#ref-for-resolve-an-origin① link-type="dfn"}
algorithm can return can differ between invocations.

The particulars (apart from the cache key) are not tied down as they are
not pertinent to the system the Fetch Standard establishes. Other
documents ought not to build on this primitive without having a
considered discussion with the Fetch Standard community first.
:::
::::

### [2.6. ]{.secno}[Connections]{.content}[](#connections){.self-link} {#connections .heading .settled level="2.6"}

A user agent has an associated [connection
pool]{#concept-connection-pool .dfn .dfn-paneled dfn-type="dfn"
export=""}. A [connection
pool](#concept-connection-pool){#ref-for-concept-connection-pool
link-type="dfn"} is an [ordered
set](https://infra.spec.whatwg.org/#ordered-set){#ref-for-ordered-set④
link-type="dfn"} of zero or more [connections]{#concept-connection .dfn
.dfn-paneled dfn-type="dfn" export="" lt="connection"}. Each
[connection](#concept-connection){#ref-for-concept-connection
link-type="dfn"} is identified by an associated [key]{#connection-key
.dfn .dfn-paneled dfn-for="connection" dfn-type="dfn" noexport=""} (a
[network partition
key](#network-partition-key){#ref-for-network-partition-key①
link-type="dfn"}), [origin]{#connection-origin .dfn .dfn-paneled
dfn-for="connection" dfn-type="dfn" noexport=""} (an
[origin](https://html.spec.whatwg.org/multipage/browsers.html#concept-origin){#ref-for-concept-origin③
link-type="dfn"}), and [credentials]{#connection-credentials .dfn
.dfn-paneled dfn-for="connection" dfn-type="dfn" noexport=""} (a
boolean).

Each [connection](#concept-connection){#ref-for-concept-connection①
link-type="dfn"} has an associated [timing
info]{#concept-connection-timing-info .dfn .dfn-paneled
dfn-for="connection" dfn-type="dfn" noexport=""} (a [connection timing
info](#connection-timing-info){#ref-for-connection-timing-info①
link-type="dfn"}).

A [connection timing info]{#connection-timing-info .dfn .dfn-paneled
dfn-type="dfn" export=""} is a
[struct](https://infra.spec.whatwg.org/#struct){#ref-for-struct④
link-type="dfn"} used to maintain timing information pertaining to the
process of obtaining a connection. It has the following
[items](https://infra.spec.whatwg.org/#struct-item){#ref-for-struct-item④
link-type="dfn"}:

[domain lookup start time]{#connection-timing-info-domain-lookup-start-time .dfn .dfn-paneled dfn-for="connection timing info" dfn-type="dfn" export=""} (default 0)\
[domain lookup end time]{#connection-timing-info-domain-lookup-end-time .dfn .dfn-paneled dfn-for="connection timing info" dfn-type="dfn" export=""} (default 0)\
[connection start time]{#connection-timing-info-connection-start-time .dfn .dfn-paneled dfn-for="connection timing info" dfn-type="dfn" export=""} (default 0)\
[connection end time]{#connection-timing-info-connection-end-time .dfn .dfn-paneled dfn-for="connection timing info" dfn-type="dfn" export=""} (default 0)\
[secure connection start time]{#connection-timing-info-secure-connection-start-time .dfn .dfn-paneled dfn-for="connection timing info" dfn-type="dfn" export=""} (default 0)
:   A
    [`DOMHighResTimeStamp`{.idl}](https://w3c.github.io/hr-time/#dom-domhighrestimestamp){#ref-for-dom-domhighrestimestamp①
    link-type="idl"}.

[ALPN negotiated protocol]{#connection-timing-info-alpn-negotiated-protocol .dfn .dfn-paneled dfn-for="connection timing info" dfn-type="dfn" export=""} (default the empty [byte sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence①③ link-type="dfn"})
:   A [byte
    sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence①④
    link-type="dfn"}.

::: {.algorithm algorithm="clamp and coarsen connection timing info"}
To [clamp and coarsen connection timing
info]{#clamp-and-coarsen-connection-timing-info .dfn .dfn-paneled
dfn-type="dfn" noexport=""}, given a [connection timing
info](#connection-timing-info){#ref-for-connection-timing-info②
link-type="dfn"} `timingInfo`{.variable}, a
[`DOMHighResTimeStamp`{.idl}](https://w3c.github.io/hr-time/#dom-domhighrestimestamp){#ref-for-dom-domhighrestimestamp②
link-type="idl"} `defaultStartTime`{.variable}, and a boolean
`crossOriginIsolatedCapability`{.variable}, run these steps:

1.  If `timingInfo`{.variable}'s [connection start
    time](#connection-timing-info-connection-start-time){#ref-for-connection-timing-info-connection-start-time
    link-type="dfn"} is less than `defaultStartTime`{.variable}, then
    return a new [connection timing
    info](#connection-timing-info){#ref-for-connection-timing-info③
    link-type="dfn"} whose [domain lookup start
    time](#connection-timing-info-domain-lookup-start-time){#ref-for-connection-timing-info-domain-lookup-start-time
    link-type="dfn"} is `defaultStartTime`{.variable}, [domain lookup
    end
    time](#connection-timing-info-domain-lookup-end-time){#ref-for-connection-timing-info-domain-lookup-end-time
    link-type="dfn"} is `defaultStartTime`{.variable}, [connection start
    time](#connection-timing-info-connection-start-time){#ref-for-connection-timing-info-connection-start-time①
    link-type="dfn"} is `defaultStartTime`{.variable}, [connection end
    time](#connection-timing-info-connection-end-time){#ref-for-connection-timing-info-connection-end-time
    link-type="dfn"} is `defaultStartTime`{.variable}, [secure
    connection start
    time](#connection-timing-info-secure-connection-start-time){#ref-for-connection-timing-info-secure-connection-start-time
    link-type="dfn"} is `defaultStartTime`{.variable}, and [ALPN
    negotiated
    protocol](#connection-timing-info-alpn-negotiated-protocol){#ref-for-connection-timing-info-alpn-negotiated-protocol
    link-type="dfn"} is `timingInfo`{.variable}'s [ALPN negotiated
    protocol](#connection-timing-info-alpn-negotiated-protocol){#ref-for-connection-timing-info-alpn-negotiated-protocol①
    link-type="dfn"}.

2.  Return a new [connection timing
    info](#connection-timing-info){#ref-for-connection-timing-info④
    link-type="dfn"} whose [domain lookup start
    time](#connection-timing-info-domain-lookup-start-time){#ref-for-connection-timing-info-domain-lookup-start-time①
    link-type="dfn"} is the result of [coarsen
    time](https://w3c.github.io/hr-time/#dfn-coarsen-time){#ref-for-dfn-coarsen-time
    link-type="dfn"} given `timingInfo`{.variable}'s [domain lookup
    start
    time](#connection-timing-info-domain-lookup-start-time){#ref-for-connection-timing-info-domain-lookup-start-time②
    link-type="dfn"} and `crossOriginIsolatedCapability`{.variable},
    [domain lookup end
    time](#connection-timing-info-domain-lookup-end-time){#ref-for-connection-timing-info-domain-lookup-end-time①
    link-type="dfn"} is the result of [coarsen
    time](https://w3c.github.io/hr-time/#dfn-coarsen-time){#ref-for-dfn-coarsen-time①
    link-type="dfn"} given `timingInfo`{.variable}'s [domain lookup end
    time](#connection-timing-info-domain-lookup-end-time){#ref-for-connection-timing-info-domain-lookup-end-time②
    link-type="dfn"} and `crossOriginIsolatedCapability`{.variable},
    [connection start
    time](#connection-timing-info-connection-start-time){#ref-for-connection-timing-info-connection-start-time②
    link-type="dfn"} is the result of [coarsen
    time](https://w3c.github.io/hr-time/#dfn-coarsen-time){#ref-for-dfn-coarsen-time②
    link-type="dfn"} given `timingInfo`{.variable}'s [connection start
    time](#connection-timing-info-connection-start-time){#ref-for-connection-timing-info-connection-start-time③
    link-type="dfn"} and `crossOriginIsolatedCapability`{.variable},
    [connection end
    time](#connection-timing-info-connection-end-time){#ref-for-connection-timing-info-connection-end-time①
    link-type="dfn"} is the result of [coarsen
    time](https://w3c.github.io/hr-time/#dfn-coarsen-time){#ref-for-dfn-coarsen-time③
    link-type="dfn"} given `timingInfo`{.variable}'s [connection end
    time](#connection-timing-info-connection-end-time){#ref-for-connection-timing-info-connection-end-time②
    link-type="dfn"} and `crossOriginIsolatedCapability`{.variable},
    [secure connection start
    time](#connection-timing-info-secure-connection-start-time){#ref-for-connection-timing-info-secure-connection-start-time①
    link-type="dfn"} is the result of [coarsen
    time](https://w3c.github.io/hr-time/#dfn-coarsen-time){#ref-for-dfn-coarsen-time④
    link-type="dfn"} given `timingInfo`{.variable}'s [connection end
    time](#connection-timing-info-connection-end-time){#ref-for-connection-timing-info-connection-end-time③
    link-type="dfn"} and `crossOriginIsolatedCapability`{.variable}, and
    [ALPN negotiated
    protocol](#connection-timing-info-alpn-negotiated-protocol){#ref-for-connection-timing-info-alpn-negotiated-protocol②
    link-type="dfn"} is `timingInfo`{.variable}'s [ALPN negotiated
    protocol](#connection-timing-info-alpn-negotiated-protocol){#ref-for-connection-timing-info-alpn-negotiated-protocol③
    link-type="dfn"}.
:::

------------------------------------------------------------------------

A [new connection setting]{#new-connection-setting .dfn .dfn-paneled
dfn-type="dfn" noexport=""} is \"`no`\", \"`yes`\", or
\"`yes-and-dedicated`\".

::: {.algorithm algorithm="obtain a connection"}
To [obtain a connection]{#concept-connection-obtain .dfn .dfn-paneled
dfn-type="dfn" export=""}, given a [network partition
key](#network-partition-key){#ref-for-network-partition-key②
link-type="dfn"} `key`{.variable},
[URL](https://url.spec.whatwg.org/#concept-url){#ref-for-concept-url①④
link-type="dfn"} `url`{.variable}, boolean `credentials`{.variable}, an
optional [new connection
setting](#new-connection-setting){#ref-for-new-connection-setting
link-type="dfn"} `new`{.variable} (default \"`no`\"), and an optional
boolean
[`requireUnreliable`{.variable}]{#obtain-a-connection-requireunreliable
.dfn .dfn-paneled dfn-for="obtain a connection" dfn-type="dfn"
export=""} (default false), run these steps:

1.  If `new`{.variable} is \"`no`\", then:

    1.  Let `connections`{.variable} be a set of
        [connections](#concept-connection){#ref-for-concept-connection②
        link-type="dfn"} in the user agent's [connection
        pool](#concept-connection-pool){#ref-for-concept-connection-pool①
        link-type="dfn"} whose
        [key](#connection-key){#ref-for-connection-key link-type="dfn"}
        is `key`{.variable},
        [origin](#connection-origin){#ref-for-connection-origin
        link-type="dfn"} is `url`{.variable}'s
        [origin](https://url.spec.whatwg.org/#concept-url-origin){#ref-for-concept-url-origin④
        link-type="dfn"}, and
        [credentials](#connection-credentials){#ref-for-connection-credentials
        link-type="dfn"} is `credentials`{.variable}.

    2.  If `connections`{.variable} is not empty and
        `requireUnreliable`{.variable} is false, then return one of
        `connections`{.variable}.

    3.  If there is a
        [connection](#concept-connection){#ref-for-concept-connection③
        link-type="dfn"} capable of supporting unreliable transport in
        `connections`{.variable}, e.g., HTTP/3, then return that
        [connection](#concept-connection){#ref-for-concept-connection④
        link-type="dfn"}.

2.  Let `proxies`{.variable} be the result of finding proxies for
    `url`{.variable} in an
    [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#ref-for-implementation-defined⑤
    link-type="dfn"} manner. If there are no proxies, let
    `proxies`{.variable} be « \"`DIRECT`\" ».

    This is where non-standard technology such as [Web Proxy
    Auto-Discovery Protocol
    (WPAD)](https://en.wikipedia.org/wiki/Web_Proxy_Auto-Discovery_Protocol)
    and [proxy auto-config
    (PAC)](https://en.wikipedia.org/wiki/Proxy_auto-config) come into
    play. The \"`DIRECT`\" value means to not use a proxy for this
    particular `url`{.variable}.

3.  Let `timingInfo`{.variable} be a new [connection timing
    info](#connection-timing-info){#ref-for-connection-timing-info⑤
    link-type="dfn"}.

4.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate⑦
    link-type="dfn"} `proxy`{.variable} of `proxies`{.variable}:

    1.  Set `timingInfo`{.variable}'s [domain lookup start
        time](#connection-timing-info-domain-lookup-start-time){#ref-for-connection-timing-info-domain-lookup-start-time③
        link-type="dfn"} to the [unsafe shared current
        time](https://w3c.github.io/hr-time/#dfn-unsafe-shared-current-time){#ref-for-dfn-unsafe-shared-current-time
        link-type="dfn"}.

    2.  Let `hosts`{.variable} be « `url`{.variable}'s
        [origin](https://url.spec.whatwg.org/#concept-url-origin){#ref-for-concept-url-origin⑤
        link-type="dfn"}'s
        [host](https://html.spec.whatwg.org/multipage/browsers.html#concept-origin-host){#ref-for-concept-origin-host③
        link-type="dfn"} ».

    3.  If `proxy`{.variable} is \"`DIRECT`\", then set
        `hosts`{.variable} to the result of running [resolve an
        origin](#resolve-an-origin){#ref-for-resolve-an-origin②
        link-type="dfn"} given `key`{.variable} and `url`{.variable}'s
        [origin](https://url.spec.whatwg.org/#concept-url-origin){#ref-for-concept-url-origin⑥
        link-type="dfn"}.

    4.  If `hosts`{.variable} is failure, then
        [continue](https://infra.spec.whatwg.org/#iteration-continue){#ref-for-iteration-continue②
        link-type="dfn"}.

    5.  Set `timingInfo`{.variable}'s [domain lookup end
        time](#connection-timing-info-domain-lookup-end-time){#ref-for-connection-timing-info-domain-lookup-end-time③
        link-type="dfn"} to the [unsafe shared current
        time](https://w3c.github.io/hr-time/#dfn-unsafe-shared-current-time){#ref-for-dfn-unsafe-shared-current-time①
        link-type="dfn"}.

    6.  Let `connection`{.variable} be the result of running this step:
        run [create a
        connection](#create-a-connection){#ref-for-create-a-connection
        link-type="dfn"} given `key`{.variable}, `url`{.variable}'s
        [origin](https://url.spec.whatwg.org/#concept-url-origin){#ref-for-concept-url-origin⑦
        link-type="dfn"}, `credentials`{.variable}, `proxy`{.variable},
        an
        [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#ref-for-implementation-defined⑥
        link-type="dfn"}
        [host](https://url.spec.whatwg.org/#concept-host){#ref-for-concept-host
        link-type="dfn"} from `hosts`{.variable},
        `timingInfo`{.variable}, and `requireUnreliable`{.variable} an
        [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#ref-for-implementation-defined⑦
        link-type="dfn"} number of times, [in
        parallel](https://html.spec.whatwg.org/multipage/infrastructure.html#in-parallel){#ref-for-in-parallel
        link-type="dfn"} from each other, and wait for at least 1 to
        return a value. In an
        [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#ref-for-implementation-defined⑧
        link-type="dfn"} manner, select a value to return from the
        returned values and return it. Any other returned values that
        are
        [connections](#concept-connection){#ref-for-concept-connection⑤
        link-type="dfn"} may be closed.

        Essentially this allows an implementation to pick one or more
        [IP
        addresses](https://url.spec.whatwg.org/#ip-address){#ref-for-ip-address⑤
        link-type="dfn"} from the return value of [resolve an
        origin](#resolve-an-origin){#ref-for-resolve-an-origin③
        link-type="dfn"} (assuming `proxy`{.variable} is \"`DIRECT`\")
        and race them against each other, favor [IPv6
        addresses](https://url.spec.whatwg.org/#concept-ipv6){#ref-for-concept-ipv6
        link-type="dfn"}, retry in case of a timeout, etc.

    7.  If `connection`{.variable} is failure, then
        [continue](https://infra.spec.whatwg.org/#iteration-continue){#ref-for-iteration-continue③
        link-type="dfn"}.

    8.  If `new`{.variable} is not \"`yes-and-dedicated`\", then
        [append](https://infra.spec.whatwg.org/#set-append){#ref-for-set-append①
        link-type="dfn"} `connection`{.variable} to the user agent's
        [connection
        pool](#concept-connection-pool){#ref-for-concept-connection-pool②
        link-type="dfn"}.

    9.  Return `connection`{.variable}.

5.  Return failure.

This is intentionally a little vague as there are a lot of nuances to
connection management that are best left to the discretion of
implementers. Describing this helps explain the `<link rel=preconnect>`
feature and clearly stipulates that
[connections](#concept-connection){#ref-for-concept-connection⑥
link-type="dfn"} are keyed on
[credentials](#credentials){#ref-for-credentials① link-type="dfn"}. The
latter clarifies that, e.g., TLS session identifiers are not reused
across [connections](#concept-connection){#ref-for-concept-connection⑦
link-type="dfn"} whose
[credentials](#connection-credentials){#ref-for-connection-credentials①
link-type="dfn"} are false with
[connections](#concept-connection){#ref-for-concept-connection⑧
link-type="dfn"} whose
[credentials](#connection-credentials){#ref-for-connection-credentials②
link-type="dfn"} are true.
:::

------------------------------------------------------------------------

::: {.algorithm algorithm="create a connection"}
To [create a connection]{#create-a-connection .dfn .dfn-paneled
dfn-type="dfn" noexport=""}, given a [network partition
key](#network-partition-key){#ref-for-network-partition-key③
link-type="dfn"} `key`{.variable},
[origin](https://html.spec.whatwg.org/multipage/browsers.html#concept-origin){#ref-for-concept-origin④
link-type="dfn"} `origin`{.variable}, boolean `credentials`{.variable},
string `proxy`{.variable},
[host](https://url.spec.whatwg.org/#concept-host){#ref-for-concept-host①
link-type="dfn"} `host`{.variable}, [connection timing
info](#connection-timing-info){#ref-for-connection-timing-info⑥
link-type="dfn"} `timingInfo`{.variable}, and boolean
`requireUnreliable`{.variable}, run these steps:

1.  Set `timingInfo`{.variable}'s [connection start
    time](#connection-timing-info-connection-start-time){#ref-for-connection-timing-info-connection-start-time④
    link-type="dfn"} to the [unsafe shared current
    time](https://w3c.github.io/hr-time/#dfn-unsafe-shared-current-time){#ref-for-dfn-unsafe-shared-current-time②
    link-type="dfn"}.

2.  Let `connection`{.variable} be a new
    [connection](#concept-connection){#ref-for-concept-connection⑨
    link-type="dfn"} whose
    [key](#connection-key){#ref-for-connection-key① link-type="dfn"} is
    `key`{.variable},
    [origin](#connection-origin){#ref-for-connection-origin①
    link-type="dfn"} is `origin`{.variable},
    [credentials](#connection-credentials){#ref-for-connection-credentials③
    link-type="dfn"} is `credentials`{.variable}, and [timing
    info](#concept-connection-timing-info){#ref-for-concept-connection-timing-info
    link-type="dfn"} is `timingInfo`{.variable}. [Record connection
    timing
    info](#record-connection-timing-info){#ref-for-record-connection-timing-info
    link-type="dfn"} given `connection`{.variable} and use
    `connection`{.variable} to establish an HTTP connection to
    `host`{.variable}, taking `proxy`{.variable} and `origin`{.variable}
    into account, with the following caveats:
    [\[HTTP\]](#biblio-http "HTTP Semantics"){link-type="biblio"}
    [\[HTTP1\]](#biblio-http1 "HTTP/1.1"){link-type="biblio"}
    [\[TLS\]](#biblio-tls "The Transport Layer Security (TLS) Protocol Version 1.3"){link-type="biblio"}

    - If `requireUnreliable`{.variable} is true, then establish a
      connection capable of unreliable transport, e.g., an HTTP/3
      connection.
      [\[HTTP3\]](#biblio-http3 "HTTP/3"){link-type="biblio"}

    - When establishing a connection capable of unreliable transport,
      enable options that are necessary for WebTransport. For HTTP/3,
      this means including `SETTINGS_ENABLE_WEBTRANSPORT` with a value
      of `1` and `H3_DATAGRAM` with a value of `1` in the initial
      `SETTINGS` frame.
      [\[WEBTRANSPORT-HTTP3\]](#biblio-webtransport-http3 "WebTransport over HTTP/3"){link-type="biblio"}
      [\[HTTP3-DATAGRAM\]](#biblio-http3-datagram "HTTP Datagrams and the Capsule Protocol"){link-type="biblio"}

    - If `credentials`{.variable} is false, then do not send a TLS
      client certificate.

    - If establishing a connection does not succeed (e.g., a UDP, TCP,
      or TLS error), then return failure.

3.  Set `timingInfo`{.variable}'s [ALPN negotiated
    protocol](#connection-timing-info-alpn-negotiated-protocol){#ref-for-connection-timing-info-alpn-negotiated-protocol④
    link-type="dfn"} to `connection`{.variable}'s ALPN Protocol ID, with
    the following caveats:
    [\[RFC7301\]](#biblio-rfc7301 "Transport Layer Security (TLS) Application-Layer Protocol Negotiation Extension"){link-type="biblio"}

    - When a proxy is configured, if a tunnel connection is established
      then this must be the ALPN Protocol ID of the tunneled protocol,
      otherwise it must be the ALPN Protocol ID of the first hop to the
      proxy.

    - In case the user agent is using an experimental, non-registered
      protocol, the user agent must use the used ALPN Protocol ID, if
      any. If ALPN was not used for protocol negotiations, the user
      agent may use another descriptive string.

      `timingInfo`{.variable}'s [ALPN negotiated
      protocol](#connection-timing-info-alpn-negotiated-protocol){#ref-for-connection-timing-info-alpn-negotiated-protocol⑤
      link-type="dfn"} is intended to identify the network protocol in
      use regardless of how it was actually negotiated; that is, even if
      ALPN is not used to negotiate the network protocol, this is the
      ALPN Protocol IDs that indicates the protocol in use.

    IANA maintains a [list of ALPN Protocol
    IDs](https://www.iana.org/assignments/tls-extensiontype-values/tls-extensiontype-values.xhtml#alpn-protocol-ids).

4.  Return `connection`{.variable}.
:::

------------------------------------------------------------------------

::: {.algorithm algorithm="record connection timing info"}
To [record connection timing info]{#record-connection-timing-info .dfn
.dfn-paneled dfn-type="dfn" noexport=""} given a
[connection](#concept-connection){#ref-for-concept-connection①⓪
link-type="dfn"} `connection`{.variable}, let `timingInfo`{.variable} be
`connection`{.variable}'s [timing
info](#concept-connection-timing-info){#ref-for-concept-connection-timing-info①
link-type="dfn"} and observe these requirements:

- `timingInfo`{.variable}'s [connection end
  time](#connection-timing-info-connection-end-time){#ref-for-connection-timing-info-connection-end-time④
  link-type="dfn"} should be the [unsafe shared current
  time](https://w3c.github.io/hr-time/#dfn-unsafe-shared-current-time){#ref-for-dfn-unsafe-shared-current-time③
  link-type="dfn"} immediately after establishing the connection to the
  server or proxy, as follows:

  - The returned time must include the time interval to establish the
    transport connection, as well as other time intervals such as SOCKS
    authentication. It must include the time interval to complete enough
    of the TLS handshake to request the resource.

  - If the user agent used TLS False Start for this connection, this
    interval must not include the time needed to receive the server's
    Finished message.
    [\[RFC7918\]](#biblio-rfc7918 "Transport Layer Security (TLS) False Start"){link-type="biblio"}

  - If the user agent sends the request with early data without waiting
    for the full handshake to complete, this interval must not include
    the time needed to receive the server's ServerHello message.
    [\[RFC8470\]](#biblio-rfc8470 "Using Early Data in HTTP"){link-type="biblio"}

  - If the user agent waits for full handshake completion to send the
    request, this interval includes the full TLS handshake even if other
    requests were sent using early data on `connection`{.variable}.

  [](#example-connection-end-time){.self-link}Suppose the user agent
  establishes an HTTP/2 connection over TLS 1.3 to send a `GET` request
  and a `POST` request. It sends the ClientHello at time `t1`{.variable}
  and then sends the `GET` request with early data. The `POST` request
  is not safe
  ([\[HTTP\]](#biblio-http "HTTP Semantics"){link-type="biblio"},
  section 9.2.1), so the user agent waits to complete the handshake at
  time `t2`{.variable} before sending it. Although both requests used
  the same connection, the `GET` request reports a connection end time
  of `t1`{.variable}, while the `POST` request reports `t2`{.variable}.

- If a secure transport is used, `timingInfo`{.variable}'s [secure
  connection start
  time](#connection-timing-info-secure-connection-start-time){#ref-for-connection-timing-info-secure-connection-start-time②
  link-type="dfn"} should be the result of calling [unsafe shared
  current
  time](https://w3c.github.io/hr-time/#dfn-unsafe-shared-current-time){#ref-for-dfn-unsafe-shared-current-time④
  link-type="dfn"} immediately before starting the handshake process to
  secure `connection`{.variable}.
  [\[TLS\]](#biblio-tls "The Transport Layer Security (TLS) Protocol Version 1.3"){link-type="biblio"}

- If `connection`{.variable} is an HTTP/3 connection,
  `timingInfo`{.variable}'s [connection start
  time](#connection-timing-info-connection-start-time){#ref-for-connection-timing-info-connection-start-time⑤
  link-type="dfn"} and `timingInfo`{.variable}'s [secure connection
  start
  time](#connection-timing-info-secure-connection-start-time){#ref-for-connection-timing-info-secure-connection-start-time③
  link-type="dfn"} must be equal. (In HTTP/3 the secure transport
  handshake process is performed as part of the initial connection
  setup.) [\[HTTP3\]](#biblio-http3 "HTTP/3"){link-type="biblio"}

The [clamp and coarsen connection timing
info](#clamp-and-coarsen-connection-timing-info){#ref-for-clamp-and-coarsen-connection-timing-info
link-type="dfn"} algorithm ensures that details of reused connections
are not exposed and time values are coarsened.
:::

### [2.7. ]{.secno}[Network partition keys]{.content}[](#network-partition-keys){.self-link} {#network-partition-keys .heading .settled level="2.7"}

A [network partition key]{#network-partition-key .dfn .dfn-paneled
dfn-type="dfn" export=""} is a tuple consisting of a
[site](https://html.spec.whatwg.org/multipage/browsers.html#site){#ref-for-site
link-type="dfn"} and null or an
[implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#ref-for-implementation-defined⑨
link-type="dfn"} value.

::: {.algorithm algorithm="determine the network partition key"}
To [determine the network partition
key]{#determine-the-network-partition-key .dfn .dfn-paneled
dfn-type="dfn" export=""}, given an
[environment](https://html.spec.whatwg.org/multipage/webappapis.html#environment){#ref-for-environment②
link-type="dfn"} `environment`{.variable}:

1.  Let `topLevelOrigin`{.variable} be `environment`{.variable}'s
    [top-level
    origin](https://html.spec.whatwg.org/multipage/webappapis.html#concept-environment-top-level-origin){#ref-for-concept-environment-top-level-origin
    link-type="dfn"}.

2.  If `topLevelOrigin`{.variable} is null, then set
    `topLevelOrigin`{.variable} to `environment`{.variable}'s [top-level
    creation
    URL](https://html.spec.whatwg.org/multipage/webappapis.html#concept-environment-top-level-creation-url){#ref-for-concept-environment-top-level-creation-url
    link-type="dfn"}'s
    [origin](https://url.spec.whatwg.org/#concept-url-origin){#ref-for-concept-url-origin⑧
    link-type="dfn"}.

3.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert①⑤
    link-type="dfn"}: `topLevelOrigin`{.variable} is an
    [origin](https://html.spec.whatwg.org/multipage/browsers.html#concept-origin){#ref-for-concept-origin⑤
    link-type="dfn"}.

4.  Let `topLevelSite`{.variable} be the result of [obtaining a
    site](https://html.spec.whatwg.org/multipage/browsers.html#obtain-a-site){#ref-for-obtain-a-site
    link-type="dfn"}, given `topLevelOrigin`{.variable}.

5.  Let `secondKey`{.variable} be null or an
    [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#ref-for-implementation-defined①⓪
    link-type="dfn"} value.

    The second key is intentionally a little vague as the finer points
    are still evolving. See [issue
    #1035](https://github.com/whatwg/fetch/issues/1035).

6.  Return (`topLevelSite`{.variable}, `secondKey`{.variable}).
:::

::: {.algorithm algorithm="determine the network partition key" algorithm-for="request"}
To [determine the network partition
key]{#request-determine-the-network-partition-key .dfn .dfn-paneled
dfn-for="request" dfn-type="dfn" noexport=""}, given a
[request](#concept-request){#ref-for-concept-request⑦⑤ link-type="dfn"}
`request`{.variable}:

1.  If `request`{.variable}'s [reserved
    client](#concept-request-reserved-client){#ref-for-concept-request-reserved-client
    link-type="dfn"} is non-null, then return the result of [determining
    the network partition
    key](#determine-the-network-partition-key){#ref-for-determine-the-network-partition-key
    link-type="dfn"} given `request`{.variable}'s [reserved
    client](#concept-request-reserved-client){#ref-for-concept-request-reserved-client①
    link-type="dfn"}.

2.  If `request`{.variable}'s
    [client](#concept-request-client){#ref-for-concept-request-client④
    link-type="dfn"} is non-null, then return the result of [determining
    the network partition
    key](#determine-the-network-partition-key){#ref-for-determine-the-network-partition-key①
    link-type="dfn"} given `request`{.variable}'s
    [client](#concept-request-client){#ref-for-concept-request-client⑤
    link-type="dfn"}.

3.  Return null.
:::

### [2.8. ]{.secno}[HTTP cache partitions]{.content}[](#http-cache-partitions){.self-link} {#http-cache-partitions .heading .settled level="2.8"}

::: {.algorithm algorithm="determine the HTTP cache partition"}
To [determine the HTTP cache
partition]{#determine-the-http-cache-partition .dfn .dfn-paneled
dfn-type="dfn" noexport=""}, given a
[request](#concept-request){#ref-for-concept-request⑦⑥ link-type="dfn"}
`request`{.variable}:

1.  Let `key`{.variable} be the result of [determining the network
    partition
    key](#request-determine-the-network-partition-key){#ref-for-request-determine-the-network-partition-key
    link-type="dfn"} given `request`{.variable}.

2.  If `key`{.variable} is null, then return null.

3.  Return the unique HTTP cache associated with `key`{.variable}.
    [\[HTTP-CACHING\]](#biblio-http-caching "HTTP Caching"){link-type="biblio"}
:::

### [2.9. ]{.secno}[Port blocking]{.content}[](#port-blocking){.self-link} {#port-blocking .heading .settled level="2.9"}

New protocols can avoid the need for blocking ports by negotiating the
protocol through TLS using ALPN. The protocol cannot be spoofed through
HTTP requests in that case.
[\[RFC7301\]](#biblio-rfc7301 "Transport Layer Security (TLS) Application-Layer Protocol Negotiation Extension"){link-type="biblio"}

::: {.algorithm algorithm="block bad port"}
To determine whether fetching a
[request](#concept-request){#ref-for-concept-request⑦⑦ link-type="dfn"}
`request`{.variable} [should be blocked due to a bad
port]{#block-bad-port .dfn .dfn-paneled dfn-type="dfn" export=""
lt="block bad port"}:

1.  Let `url`{.variable} be `request`{.variable}'s [current
    URL](#concept-request-current-url){#ref-for-concept-request-current-url③
    link-type="dfn"}.

2.  If `url`{.variable}'s
    [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#ref-for-concept-url-scheme①
    link-type="dfn"} is an [HTTP(S)
    scheme](#http-scheme){#ref-for-http-scheme③ link-type="dfn"} and
    `url`{.variable}'s
    [port](https://url.spec.whatwg.org/#concept-url-port){#ref-for-concept-url-port
    link-type="dfn"} is a [bad port](#bad-port){#ref-for-bad-port
    link-type="dfn"}, then return **blocked**.

3.  Return **allowed**.
:::

A
[port](https://url.spec.whatwg.org/#concept-url-port){#ref-for-concept-url-port①
link-type="dfn"} is a [bad port]{#bad-port .dfn .dfn-paneled
dfn-type="dfn" export=""} if it is listed in the first column of the
following table.

Port

Typical service

1

tcpmux

7

echo

9

discard

11

systat

13

daytime

15

netstat

17

qotd

19

chargen

20

ftp-data

21

ftp

22

ssh

23

telnet

25

smtp

37

time

42

name

43

nicname

53

domain

69

tftp

77

---​

79

finger

87

---​

95

supdup

101

hostname

102

iso-tsap

103

gppitnp

104

acr-nema

109

pop2

110

pop3

111

sunrpc

113

auth

115

sftp

117

uucp-path

119

nntp

123

ntp

135

epmap

137

netbios-ns

139

netbios-ssn

143

imap

161

snmp

179

bgp

389

ldap

427

svrloc

465

submissions

512

exec

513

login

514

shell

515

printer

526

tempo

530

courier

531

chat

532

netnews

540

uucp

548

afp

554

rtsp

556

remotefs

563

nntps

587

submission

601

syslog-conn

636

ldaps

989

ftps-data

990

ftps

993

imaps

995

pop3s

1719

h323gatestat

1720

h323hostcall

1723

pptp

2049

nfs

3659

apple-sasl

4045

npp

4190

sieve

5060

sip

5061

sips

6000

x11

6566

sane-port

6665

ircu

6666

ircu

6667

ircu

6668

ircu

6669

ircu

6679

osaut

6697

ircs-u

10080

amanda

::: {.algorithm algorithm="should response to request be blocked due to mime type"}
### [2.10. ]{.secno}[Should `response`{.variable} to `request`{.variable} be blocked due to its MIME type?]{.content}[](#should-response-to-request-be-blocked-due-to-mime-type?){#ref-for-should-response-to-request-be-blocked-due-to-mime-type? .self-link} {#should-response-to-request-be-blocked-due-to-mime-type? .heading .settled .dfn-paneled dfn-type="dfn" export="" level="2.10" lt="should response to request be blocked due to mime type"}

Run these steps:

1.  Let `mimeType`{.variable} be the result of [extracting a MIME
    type](#concept-header-extract-mime-type){#ref-for-concept-header-extract-mime-type②
    link-type="dfn"} from `response`{.variable}'s [header
    list](#concept-response-header-list){#ref-for-concept-response-header-list⑧
    link-type="dfn"}.

2.  If `mimeType`{.variable} is failure, then return **allowed**.

3.  Let `destination`{.variable} be `request`{.variable}'s
    [destination](#concept-request-destination){#ref-for-concept-request-destination⑧
    link-type="dfn"}.

4.  If `destination`{.variable} is
    [script-like](#request-destination-script-like){#ref-for-request-destination-script-like①
    link-type="dfn"} and one of the following is true, then return
    **blocked**:

    - `mimeType`{.variable}'s
      [essence](https://mimesniff.spec.whatwg.org/#mime-type-essence){#ref-for-mime-type-essence①
      link-type="dfn"} [starts
      with](https://infra.spec.whatwg.org/#string-starts-with){#ref-for-string-starts-with①
      link-type="dfn"} \"`audio/`\", \"`image/`\", or \"`video/`\".
    - `mimeType`{.variable}'s
      [essence](https://mimesniff.spec.whatwg.org/#mime-type-essence){#ref-for-mime-type-essence②
      link-type="dfn"} is \"`text/csv`\".

5.  Return **allowed**.
:::

