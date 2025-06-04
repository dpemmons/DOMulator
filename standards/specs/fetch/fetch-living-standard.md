## [Goals]{.content} {#goals .no-num .short .heading .settled}

The goal is to unify fetching across the web platform and provide
consistent handling of everything that involves, including:

- URL schemes
- Redirects
- Cross-origin semantics
- CSP
  [\[CSP\]](#biblio-csp "Content Security Policy Level 3"){link-type="biblio"}
- Fetch Metadata
  [\[FETCH-METADATA\]](#biblio-fetch-metadata "Fetch Metadata Request Headers"){link-type="biblio"}
- Service workers
  [\[SW\]](#biblio-sw "Service Workers"){link-type="biblio"}
- Mixed Content
  [\[MIX\]](#biblio-mix "Mixed Content"){link-type="biblio"}
- Upgrade Insecure Requests
  [\[UPGRADE-INSECURE-REQUESTS\]](#biblio-upgrade-insecure-requests "Upgrade Insecure Requests"){link-type="biblio"}
- \``Referer`\`
  [\[REFERRER\]](#biblio-referrer "Referrer Policy"){link-type="biblio"}

To do so it also supersedes the HTTP
\`[`Origin`](#http-origin){#ref-for-http-origin
link-type="http-header"}\` header semantics originally defined in The
Web Origin Concept.
[\[ORIGIN\]](#biblio-origin "The Web Origin Concept"){link-type="biblio"}

## [1. ]{.secno}[Preface]{.content}[](#preface){.self-link} {#preface .short .heading .settled level="1"}

At a high level, fetching a resource is a fairly simple operation. A
request goes in, a response comes out. The details of that operation are
however quite involved and used to not be written down carefully and
differ from one API to the next.

Numerous APIs provide the ability to fetch a resource, e.g. HTML's `img`
and `script` element, CSS\' `cursor` and `list-style-image`, the
`navigator.sendBeacon()` and `self.importScripts()` JavaScript APIs. The
Fetch Standard provides a unified architecture for these features so
they are all consistent when it comes to various aspects of fetching,
such as redirects and the CORS protocol.

The Fetch Standard also defines the
[`fetch()`](#dom-global-fetch){#ref-for-dom-global-fetch .idl-code
link-type="method"} JavaScript API, which exposes most of the networking
functionality at a fairly low level of abstraction.

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

## [3. ]{.secno}[HTTP extensions]{.content}[](#http-extensions){.self-link} {#http-extensions .heading .settled level="3"}

### [3.1. ]{.secno}[\``Origin`\` header]{.content}[](#origin-header){.self-link} {#origin-header .heading .settled level="3.1"}

The \`[`Origin`]{#http-origin .dfn .dfn-paneled dfn-type="http-header"
export=""}\` request [header](#concept-header){#ref-for-concept-header②⑧
link-type="dfn"} indicates where a
[fetch](#concept-fetch){#ref-for-concept-fetch①⑧ link-type="dfn"}
originates from.

The \`[`Origin`](#http-origin){#ref-for-http-origin②
link-type="http-header"}\` header is a version of the \``Referer`\`
\[sic\] header that does not reveal a
[path](https://url.spec.whatwg.org/#concept-url-path){#ref-for-concept-url-path
link-type="dfn"}. It is used for all [HTTP
fetches](#concept-http-fetch){#ref-for-concept-http-fetch②
link-type="dfn"} whose
[request](#concept-request){#ref-for-concept-request⑦⑧
link-type="dfn"}'s [response
tainting](#concept-request-response-tainting){#ref-for-concept-request-response-tainting②
link-type="dfn"} is \"`cors`\", as well as those where
[request](#concept-request){#ref-for-concept-request⑦⑨
link-type="dfn"}'s
[method](#concept-request-method){#ref-for-concept-request-method①
link-type="dfn"} is neither \``GET`\` nor \``HEAD`\`. Due to
compatibility constraints it is not included in all
[fetches](#concept-fetch){#ref-for-concept-fetch①⑨ link-type="dfn"}.

Its possible
[values](#concept-header-value){#ref-for-concept-header-value①①
link-type="dfn"} are all the return values of [byte-serializing a
request
origin](#byte-serializing-a-request-origin){#ref-for-byte-serializing-a-request-origin
link-type="dfn"}, given a
[request](#concept-request){#ref-for-concept-request⑧⓪ link-type="dfn"}.

This supplants the definition in The Web Origin Concept.
[\[ORIGIN\]](#biblio-origin "The Web Origin Concept"){link-type="biblio"}

------------------------------------------------------------------------

::: {.algorithm algorithm="append a request `Origin` header"}
To [append a request \``Origin`\`
header]{#append-a-request-origin-header .dfn .dfn-paneled dfn-type="dfn"
noexport=""}, given a
[request](#concept-request){#ref-for-concept-request⑧① link-type="dfn"}
`request`{.variable}, run these steps:

1.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert①⑥
    link-type="dfn"}: `request`{.variable}'s
    [origin](#concept-request-origin){#ref-for-concept-request-origin⑦
    link-type="dfn"} is not \"`client`\".

2.  Let `serializedOrigin`{.variable} be the result of [byte-serializing
    a request
    origin](#byte-serializing-a-request-origin){#ref-for-byte-serializing-a-request-origin①
    link-type="dfn"} with `request`{.variable}.

3.  If `request`{.variable}'s [response
    tainting](#concept-request-response-tainting){#ref-for-concept-request-response-tainting③
    link-type="dfn"} is \"`cors`\" or `request`{.variable}'s
    [mode](#concept-request-mode){#ref-for-concept-request-mode⑤
    link-type="dfn"} is \"`websocket`\", then
    [append](#concept-header-list-append){#ref-for-concept-header-list-append①
    link-type="dfn"} (\``Origin`\`, `serializedOrigin`{.variable}) to
    `request`{.variable}'s [header
    list](#concept-request-header-list){#ref-for-concept-request-header-list③
    link-type="dfn"}.

4.  Otherwise, if `request`{.variable}'s
    [method](#concept-request-method){#ref-for-concept-request-method②
    link-type="dfn"} is neither \``GET`\` nor \``HEAD`\`, then:

    1.  If `request`{.variable}'s
        [mode](#concept-request-mode){#ref-for-concept-request-mode⑥
        link-type="dfn"} is not \"`cors`\", then switch on
        `request`{.variable}'s [referrer
        policy](#concept-request-referrer-policy){#ref-for-concept-request-referrer-policy
        link-type="dfn"}:

        \"`no-referrer`\"

        :   Set `serializedOrigin`{.variable} to \``null`\`.

        \"`no-referrer-when-downgrade`\"\
        \"`strict-origin`\"\
        \"`strict-origin-when-cross-origin`\"

        :   If `request`{.variable}'s
            [origin](#concept-request-origin){#ref-for-concept-request-origin⑧
            link-type="dfn"} is a [tuple
            origin](https://html.spec.whatwg.org/multipage/browsers.html#concept-origin-tuple){#ref-for-concept-origin-tuple
            link-type="dfn"}, its `scheme`{.variable} is \"`https`\",
            and `request`{.variable}'s [current
            URL](#concept-request-current-url){#ref-for-concept-request-current-url④
            link-type="dfn"}'s `scheme`{.variable} is not \"`https`\",
            then set `serializedOrigin`{.variable} to \``null`\`.

        \"`same-origin`\"

        :   If `request`{.variable}'s
            [origin](#concept-request-origin){#ref-for-concept-request-origin⑨
            link-type="dfn"} is not [same
            origin](https://html.spec.whatwg.org/multipage/browsers.html#same-origin){#ref-for-same-origin③
            link-type="dfn"} with `request`{.variable}'s [current
            URL](#concept-request-current-url){#ref-for-concept-request-current-url⑤
            link-type="dfn"}'s
            [origin](https://url.spec.whatwg.org/#concept-url-origin){#ref-for-concept-url-origin⑨
            link-type="dfn"}, then set `serializedOrigin`{.variable} to
            \``null`\`.

        Otherwise
        :   Do nothing.

    2.  [Append](#concept-header-list-append){#ref-for-concept-header-list-append②
        link-type="dfn"} (\``Origin`\`, `serializedOrigin`{.variable})
        to `request`{.variable}'s [header
        list](#concept-request-header-list){#ref-for-concept-request-header-list④
        link-type="dfn"}.

A [request](#concept-request){#ref-for-concept-request⑧②
link-type="dfn"}'s [referrer
policy](#concept-request-referrer-policy){#ref-for-concept-request-referrer-policy①
link-type="dfn"} is taken into account for all fetches where the fetcher
did not explicitly opt into sharing their
[origin](https://html.spec.whatwg.org/multipage/browsers.html#concept-origin){#ref-for-concept-origin⑥
link-type="dfn"} with the server, e.g., via using the [CORS
protocol](#cors-protocol){#ref-for-cors-protocol② link-type="dfn"}.
:::

### [3.2. ]{.secno}[CORS protocol]{.content}[](#http-cors-protocol){.self-link} {#http-cors-protocol .heading .settled level="3.2"}

To allow sharing responses cross-origin and allow for more versatile
[fetches](#concept-fetch){#ref-for-concept-fetch②⓪ link-type="dfn"} than
possible with HTML's
[`form`](https://html.spec.whatwg.org/multipage/forms.html#the-form-element){#ref-for-the-form-element
link-type="element"} element, the [CORS protocol]{#cors-protocol .dfn
.dfn-paneled dfn-type="dfn" export=""} exists. It is layered on top of
HTTP and allows responses to declare they can be shared with other
[origins](https://html.spec.whatwg.org/multipage/browsers.html#concept-origin){#ref-for-concept-origin⑦
link-type="dfn"}.

It needs to be an opt-in mechanism to prevent leaking data from
responses behind a firewall (intranets). Additionally, for
[requests](#concept-request){#ref-for-concept-request⑧③ link-type="dfn"}
including [credentials](#credentials){#ref-for-credentials②
link-type="dfn"} it needs to be opt-in to prevent leaking
potentially-sensitive data.

This section explains the [CORS
protocol](#cors-protocol){#ref-for-cors-protocol③ link-type="dfn"} as it
pertains to server developers. Requirements for user agents are part of
the [fetch](#concept-fetch){#ref-for-concept-fetch②① link-type="dfn"}
algorithm, except for the [new HTTP header
syntax](#http-new-header-syntax).

#### [3.2.1. ]{.secno}[General]{.content}[](#general){.self-link} {#general .heading .settled level="3.2.1"}

The [CORS protocol](#cors-protocol){#ref-for-cors-protocol④
link-type="dfn"} consists of a set of headers that indicates whether a
response can be shared cross-origin.

For [requests](#concept-request){#ref-for-concept-request⑧④
link-type="dfn"} that are more involved than what is possible with
HTML's
[`form`](https://html.spec.whatwg.org/multipage/forms.html#the-form-element){#ref-for-the-form-element①
link-type="element"} element, a [CORS-preflight
request](#cors-preflight-request){#ref-for-cors-preflight-request①
link-type="dfn"} is performed, to ensure
[request](#concept-request){#ref-for-concept-request⑧⑤
link-type="dfn"}'s [current
URL](#concept-request-current-url){#ref-for-concept-request-current-url⑥
link-type="dfn"} supports the [CORS
protocol](#cors-protocol){#ref-for-cors-protocol⑤ link-type="dfn"}.

#### [3.2.2. ]{.secno}[HTTP requests]{.content}[](#http-requests){.self-link} {#http-requests .heading .settled level="3.2.2"}

A [CORS request]{#cors-request .dfn .dfn-paneled dfn-type="dfn"
export=""} is an HTTP request that includes an
\`[`Origin`](#http-origin){#ref-for-http-origin③
link-type="http-header"}\` header. It cannot be reliably identified as
participating in the [CORS
protocol](#cors-protocol){#ref-for-cors-protocol⑥ link-type="dfn"} as
the \`[`Origin`](#http-origin){#ref-for-http-origin④
link-type="http-header"}\` header is also included for all
[requests](#concept-request){#ref-for-concept-request⑧⑥ link-type="dfn"}
whose [method](#concept-request-method){#ref-for-concept-request-method③
link-type="dfn"} is neither \``GET`\` nor \``HEAD`\`.

A [CORS-preflight request]{#cors-preflight-request .dfn .dfn-paneled
dfn-type="dfn" export=""} is a [CORS
request](#cors-request){#ref-for-cors-request① link-type="dfn"} that
checks to see if the [CORS
protocol](#cors-protocol){#ref-for-cors-protocol⑦ link-type="dfn"} is
understood. It uses \``OPTIONS`\` as
[method](#concept-method){#ref-for-concept-method⑥ link-type="dfn"} and
includes the following
[header](#concept-header){#ref-for-concept-header②⑨ link-type="dfn"}:

\`[`Access-Control-Request-Method`]{#http-access-control-request-method .dfn .dfn-paneled dfn-type="http-header" export=""}\`

:   Indicates which [method](#concept-method){#ref-for-concept-method⑦
    link-type="dfn"} a future [CORS
    request](#cors-request){#ref-for-cors-request② link-type="dfn"} to
    the same resource might use.

A [CORS-preflight
request](#cors-preflight-request){#ref-for-cors-preflight-request②
link-type="dfn"} can also include the following
[header](#concept-header){#ref-for-concept-header③⓪ link-type="dfn"}:

\`[`Access-Control-Request-Headers`]{#http-access-control-request-headers .dfn .dfn-paneled dfn-type="http-header" export=""}\`

:   Indicates which [headers](#concept-header){#ref-for-concept-header③①
    link-type="dfn"} a future [CORS
    request](#cors-request){#ref-for-cors-request③ link-type="dfn"} to
    the same resource might use.

#### [3.2.3. ]{.secno}[HTTP responses]{.content}[](#http-responses){.self-link} {#http-responses .heading .settled level="3.2.3"}

An HTTP response to a [CORS
request](#cors-request){#ref-for-cors-request④ link-type="dfn"} can
include the following
[headers](#concept-header){#ref-for-concept-header③② link-type="dfn"}:

\`[`Access-Control-Allow-Origin`]{#http-access-control-allow-origin .dfn .dfn-paneled dfn-type="http-header" export=""}\`

:   Indicates whether the response can be shared, via returning the
    literal
    [value](#concept-header-value){#ref-for-concept-header-value①②
    link-type="dfn"} of the
    \`[`Origin`](#http-origin){#ref-for-http-origin⑤
    link-type="http-header"}\` request
    [header](#concept-header){#ref-for-concept-header③③ link-type="dfn"}
    (which can be \``null`\`) or \``*`\` in a response.

\`[`Access-Control-Allow-Credentials`]{#http-access-control-allow-credentials .dfn .dfn-paneled dfn-type="http-header" export=""}\`

:   Indicates whether the response can be shared when
    [request](#concept-request){#ref-for-concept-request⑧⑦
    link-type="dfn"}'s [credentials
    mode](#concept-request-credentials-mode){#ref-for-concept-request-credentials-mode②
    link-type="dfn"} is \"`include`\".

    For a [CORS-preflight
    request](#cors-preflight-request){#ref-for-cors-preflight-request③
    link-type="dfn"},
    [request](#concept-request){#ref-for-concept-request⑧⑧
    link-type="dfn"}'s [credentials
    mode](#concept-request-credentials-mode){#ref-for-concept-request-credentials-mode③
    link-type="dfn"} is always \"`same-origin`\", i.e., it excludes
    credentials, but for any subsequent [CORS
    requests](#cors-request){#ref-for-cors-request⑤ link-type="dfn"} it
    might not be. Support therefore needs to be indicated as part of the
    HTTP response to the [CORS-preflight
    request](#cors-preflight-request){#ref-for-cors-preflight-request④
    link-type="dfn"} as well.

An HTTP response to a [CORS-preflight
request](#cors-preflight-request){#ref-for-cors-preflight-request⑤
link-type="dfn"} can include the following
[headers](#concept-header){#ref-for-concept-header③④ link-type="dfn"}:

\`[`Access-Control-Allow-Methods`]{#http-access-control-allow-methods .dfn .dfn-paneled dfn-type="http-header" export=""}\`

:   Indicates which [methods](#concept-method){#ref-for-concept-method⑧
    link-type="dfn"} are supported by the
    [response](#concept-response){#ref-for-concept-response③⑤
    link-type="dfn"}'s
    [URL](#concept-response-url){#ref-for-concept-response-url③
    link-type="dfn"} for the purposes of the [CORS
    protocol](#cors-protocol){#ref-for-cors-protocol⑧ link-type="dfn"}.

    The \``Allow`\` [header](#concept-header){#ref-for-concept-header③⑤
    link-type="dfn"} is not relevant for the purposes of the [CORS
    protocol](#cors-protocol){#ref-for-cors-protocol⑨ link-type="dfn"}.

\`[`Access-Control-Allow-Headers`]{#http-access-control-allow-headers .dfn .dfn-paneled dfn-type="http-header" export=""}\`

:   Indicates which [headers](#concept-header){#ref-for-concept-header③⑥
    link-type="dfn"} are supported by the
    [response](#concept-response){#ref-for-concept-response③⑥
    link-type="dfn"}'s
    [URL](#concept-response-url){#ref-for-concept-response-url④
    link-type="dfn"} for the purposes of the [CORS
    protocol](#cors-protocol){#ref-for-cors-protocol①⓪ link-type="dfn"}.

\`[`Access-Control-Max-Age`]{#http-access-control-max-age .dfn .dfn-paneled dfn-type="http-header" export=""}\`

:   Indicates the number of seconds (5 by default) the information
    provided by the
    \`[`Access-Control-Allow-Methods`](#http-access-control-allow-methods){#ref-for-http-access-control-allow-methods
    link-type="http-header"}\` and
    \`[`Access-Control-Allow-Headers`](#http-access-control-allow-headers){#ref-for-http-access-control-allow-headers
    link-type="http-header"}\`
    [headers](#concept-header){#ref-for-concept-header③⑦
    link-type="dfn"} can be cached.

An HTTP response to a [CORS
request](#cors-request){#ref-for-cors-request⑥ link-type="dfn"} that is
not a [CORS-preflight
request](#cors-preflight-request){#ref-for-cors-preflight-request⑥
link-type="dfn"} can also include the following
[header](#concept-header){#ref-for-concept-header③⑧ link-type="dfn"}:

\`[`Access-Control-Expose-Headers`]{#http-access-control-expose-headers .dfn .dfn-paneled dfn-type="http-header" export=""}\`

:   Indicates which [headers](#concept-header){#ref-for-concept-header③⑨
    link-type="dfn"} can be exposed as part of the response by listing
    their [names](#concept-header-name){#ref-for-concept-header-name①⑦
    link-type="dfn"}.

------------------------------------------------------------------------

A successful HTTP response, i.e., one where the server developer intends
to share it, to a [CORS request](#cors-request){#ref-for-cors-request⑦
link-type="dfn"} can use any
[status](#concept-status){#ref-for-concept-status④ link-type="dfn"}, as
long as it includes the
[headers](#concept-header){#ref-for-concept-header④⓪ link-type="dfn"}
stated above with
[values](#concept-header-value){#ref-for-concept-header-value①③
link-type="dfn"} matching up with the request.

A successful HTTP response to a [CORS-preflight
request](#cors-preflight-request){#ref-for-cors-preflight-request⑦
link-type="dfn"} is similar, except it is restricted to an [ok
status](#ok-status){#ref-for-ok-status link-type="dfn"}, e.g., 200 or
204.

Any other kind of HTTP response is not successful and will either end up
not being shared or fail the [CORS-preflight
request](#cors-preflight-request){#ref-for-cors-preflight-request⑧
link-type="dfn"}. Be aware that any work the server performs might
nonetheless leak through side channels, such as timing. If server
developers wish to denote this explicitly, the 403
[status](#concept-status){#ref-for-concept-status⑤ link-type="dfn"} can
be used, coupled with omitting the relevant
[headers](#concept-header){#ref-for-concept-header④① link-type="dfn"}.

If desired, "failure" could also be shared, but that would make it a
successful HTTP response. That is why for a successful HTTP response to
a [CORS request](#cors-request){#ref-for-cors-request⑧ link-type="dfn"}
that is not a [CORS-preflight
request](#cors-preflight-request){#ref-for-cors-preflight-request⑨
link-type="dfn"} the [status](#concept-status){#ref-for-concept-status⑥
link-type="dfn"} can be anything, including 403.

Ultimately server developers have a lot of freedom in how they handle
HTTP responses and these tactics can differ between the response to the
[CORS-preflight
request](#cors-preflight-request){#ref-for-cors-preflight-request①⓪
link-type="dfn"} and the [CORS
request](#cors-request){#ref-for-cors-request⑨ link-type="dfn"} that
follows it:

- They can provide a static response. This can be helpful when working
  with caching intermediaries. A static response can both be successful
  and not successful depending on the [CORS
  request](#cors-request){#ref-for-cors-request①⓪ link-type="dfn"}. This
  is okay.

- They can provide a dynamic response, tuned to [CORS
  request](#cors-request){#ref-for-cors-request①① link-type="dfn"}. This
  can be helpful when the response body is to be tailored to a specific
  origin or a response needs to have credentials and be successful for a
  set of origins.

#### [3.2.4. ]{.secno}[HTTP new-header syntax]{.content}[](#http-new-header-syntax){.self-link} {#http-new-header-syntax .heading .settled level="3.2.4"}

[ABNF](#abnf){#ref-for-abnf③ link-type="dfn"} for the
[values](#concept-header-value){#ref-for-concept-header-value①④
link-type="dfn"} of the
[headers](#concept-header){#ref-for-concept-header④② link-type="dfn"}
used by the [CORS protocol](#cors-protocol){#ref-for-cors-protocol①①
link-type="dfn"}:

``` {.lang-abnf .highlight}
Access-Control-Request-Method    = method
Access-Control-Request-Headers   = 1#field-name

wildcard                         = "*"
Access-Control-Allow-Origin      = origin-or-null / wildcard
Access-Control-Allow-Credentials = %s"true" ; case-sensitive
Access-Control-Expose-Headers    = #field-name
Access-Control-Max-Age           = delta-seconds
Access-Control-Allow-Methods     = #method
Access-Control-Allow-Headers     = #field-name
```

For \``Access-Control-Expose-Headers`\`,
\``Access-Control-Allow-Methods`\`, and
\``Access-Control-Allow-Headers`\` response
[headers](#concept-header){#ref-for-concept-header④③ link-type="dfn"},
the [value](#concept-header-value){#ref-for-concept-header-value①⑤
link-type="dfn"} \``*`\` counts as a wildcard for
[requests](#concept-request){#ref-for-concept-request⑧⑨ link-type="dfn"}
without [credentials](#credentials){#ref-for-credentials③
link-type="dfn"}. For such
[requests](#concept-request){#ref-for-concept-request⑨⓪ link-type="dfn"}
there is no way to solely match a [header
name](#header-name){#ref-for-header-name①⑦ link-type="dfn"} or
[method](#concept-method){#ref-for-concept-method⑨ link-type="dfn"} that
is \``*`\`.

#### [3.2.5. ]{.secno}[CORS protocol and credentials]{.content}[](#cors-protocol-and-credentials){.self-link} {#cors-protocol-and-credentials .heading .settled level="3.2.5"}

When [request](#concept-request){#ref-for-concept-request⑨①
link-type="dfn"}'s [credentials
mode](#concept-request-credentials-mode){#ref-for-concept-request-credentials-mode④
link-type="dfn"} is \"`include`\" it has an impact on the functioning of
the [CORS protocol](#cors-protocol){#ref-for-cors-protocol①②
link-type="dfn"} other than including
[credentials](#credentials){#ref-for-credentials④ link-type="dfn"} in
the [fetch](#concept-fetch){#ref-for-concept-fetch②② link-type="dfn"}.

::: {#example-xhr-credentials .example}
[](#example-xhr-credentials){.self-link}

In the old days,
[`XMLHttpRequest`{.idl}](https://xhr.spec.whatwg.org/#xmlhttprequest){#ref-for-xmlhttprequest④
link-type="idl"} could be used to set
[request](#concept-request){#ref-for-concept-request⑨②
link-type="dfn"}'s [credentials
mode](#concept-request-credentials-mode){#ref-for-concept-request-credentials-mode⑤
link-type="dfn"} to \"`include`\":

``` {.lang-javascript .highlight}
var client = new XMLHttpRequest()
client.open("GET", "./")
client.withCredentials = true
/* … */
```

Nowadays, `fetch("./", { credentials:"include" }).then(/* … */)`
suffices.
:::

A [request](#concept-request){#ref-for-concept-request⑨③
link-type="dfn"}'s [credentials
mode](#concept-request-credentials-mode){#ref-for-concept-request-credentials-mode⑥
link-type="dfn"} is not necessarily observable on the server; only when
[credentials](#credentials){#ref-for-credentials⑤ link-type="dfn"} exist
for a [request](#concept-request){#ref-for-concept-request⑨④
link-type="dfn"} can it be observed by virtue of the
[credentials](#credentials){#ref-for-credentials⑥ link-type="dfn"} being
included. Note that even so, a [CORS-preflight
request](#cors-preflight-request){#ref-for-cors-preflight-request①①
link-type="dfn"} never includes
[credentials](#credentials){#ref-for-credentials⑦ link-type="dfn"}.

The server developer therefore needs to decide whether or not responses
\"tainted\" with [credentials](#credentials){#ref-for-credentials⑧
link-type="dfn"} can be shared. And also needs to decide if
[requests](#concept-request){#ref-for-concept-request⑨⑤ link-type="dfn"}
necessitating a [CORS-preflight
request](#cors-preflight-request){#ref-for-cors-preflight-request①②
link-type="dfn"} can include
[credentials](#credentials){#ref-for-credentials⑨ link-type="dfn"}.
Generally speaking, both sharing responses and allowing requests with
[credentials](#credentials){#ref-for-credentials①⓪ link-type="dfn"} is
rather unsafe, and extreme care has to be taken to avoid the [confused
deputy problem](https://en.wikipedia.org/wiki/Confused_deputy_problem).

To share responses with
[credentials](#credentials){#ref-for-credentials①① link-type="dfn"}, the
\`[`Access-Control-Allow-Origin`](#http-access-control-allow-origin){#ref-for-http-access-control-allow-origin
link-type="http-header"}\` and
\`[`Access-Control-Allow-Credentials`](#http-access-control-allow-credentials){#ref-for-http-access-control-allow-credentials
link-type="http-header"}\`
[headers](#concept-header){#ref-for-concept-header④④ link-type="dfn"}
are important. The following table serves to illustrate the various
legal and illegal combinations for a request to
`https://rabbit.invalid/`:

Request's credentials mode

\`[`Access-Control-Allow-Origin`](#http-access-control-allow-origin){#ref-for-http-access-control-allow-origin①
link-type="http-header"}\`

\`[`Access-Control-Allow-Credentials`](#http-access-control-allow-credentials){#ref-for-http-access-control-allow-credentials①
link-type="http-header"}\`

Shared?

Notes

\"`omit`\"

\``*`\`

Omitted

✅

---​

\"`omit`\"

\``*`\`

\``true`\`

✅

If credentials mode is not \"`include`\", then
\`[`Access-Control-Allow-Credentials`](#http-access-control-allow-credentials){#ref-for-http-access-control-allow-credentials②
link-type="http-header"}\` is ignored.

\"`omit`\"

\``https://rabbit.invalid/`\`

Omitted

❌

A
[serialized](https://html.spec.whatwg.org/multipage/browsers.html#ascii-serialisation-of-an-origin){#ref-for-ascii-serialisation-of-an-origin①
link-type="dfn"} origin has no trailing slash.

\"`omit`\"

\``https://rabbit.invalid`\`

Omitted

✅

---​

\"`include`\"

\``*`\`

\``true`\`

❌

If credentials mode is \"`include`\", then
\`[`Access-Control-Allow-Origin`](#http-access-control-allow-origin){#ref-for-http-access-control-allow-origin②
link-type="http-header"}\` cannot be \``*`\`.

\"`include`\"

\``https://rabbit.invalid`\`

\``true`\`

✅

---​

\"`include`\"

\``https://rabbit.invalid`\`

\``True`\`

❌

\``true`\` is (byte) case-sensitive.

Similarly,
\`[`Access-Control-Expose-Headers`](#http-access-control-expose-headers){#ref-for-http-access-control-expose-headers①
link-type="http-header"}\`,
\`[`Access-Control-Allow-Methods`](#http-access-control-allow-methods){#ref-for-http-access-control-allow-methods①
link-type="http-header"}\`, and
\`[`Access-Control-Allow-Headers`](#http-access-control-allow-headers){#ref-for-http-access-control-allow-headers①
link-type="http-header"}\` response headers can only use \``*`\` as
value when [request](#concept-request){#ref-for-concept-request⑨⑥
link-type="dfn"}'s [credentials
mode](#concept-request-credentials-mode){#ref-for-concept-request-credentials-mode⑦
link-type="dfn"} is not \"`include`\".

#### [3.2.6. ]{.secno}[Examples]{.content}[](#cors-protocol-examples){.self-link} {#cors-protocol-examples .heading .settled level="3.2.6"}

::: {#example-simple-cors .example}
[](#example-simple-cors){.self-link}

A script at `https://foo.invalid/` wants to fetch some data from
`https://bar.invalid/`. (Neither
[credentials](#credentials){#ref-for-credentials①② link-type="dfn"} nor
response header access is important.)

``` {#unicorn}
var url = "https://bar.invalid/api?key=730d67a37d7f3d802e96396d00280768773813fbe726d116944d814422fc1a45&data=about:unicorn";
fetch(url).then(success, failure)
```

This will use the [CORS
protocol](#cors-protocol){#ref-for-cors-protocol①③ link-type="dfn"},
though this is entirely transparent to the developer from `foo.invalid`.
As part of the [CORS protocol](#cors-protocol){#ref-for-cors-protocol①④
link-type="dfn"}, the user agent will include the
\`[`Origin`](#http-origin){#ref-for-http-origin⑥
link-type="http-header"}\` header in the request:

``` {.lang-http .highlight}
Origin: https://foo.invalid
```

Upon receiving a response from `bar.invalid`, the user agent will verify
the
\`[`Access-Control-Allow-Origin`](#http-access-control-allow-origin){#ref-for-http-access-control-allow-origin③
link-type="http-header"}\` response header. If its value is either
\``https://foo.invalid`\` or \``*`\`, the user agent will invoke the
`success` callback. If it has any other value, or is missing, the user
agent will invoke the `failure` callback.
:::

::: {#example-cors-with-response-header .example}
[](#example-cors-with-response-header){.self-link}

The developer of `foo.invalid` is back, and now wants to fetch some data
from `bar.invalid` while also accessing a response header.

``` {.lang-javascript .highlight}
fetch(url).then(response => {
  var hsts = response.headers.get("strict-transport-security"),
      csp = response.headers.get("content-security-policy")
  log(hsts, csp)
})
```

`bar.invalid` provides a correct
\`[`Access-Control-Allow-Origin`](#http-access-control-allow-origin){#ref-for-http-access-control-allow-origin④
link-type="http-header"}\` response header per the earlier example. The
values of `hsts` and `csp` will depend on the
\`[`Access-Control-Expose-Headers`](#http-access-control-expose-headers){#ref-for-http-access-control-expose-headers②
link-type="http-header"}\` response header. For example, if the response
included the following headers

``` {.lang-http .highlight}
Content-Security-Policy: default-src 'self'
Strict-Transport-Security: max-age=31536000; includeSubdomains; preload
Access-Control-Expose-Headers: Content-Security-Policy
```

then `hsts` would be null and `csp` would be \"`default-src 'self'`\",
even though the response did include both headers. This is because
`bar.invalid` needs to explicitly share each header by listing their
names in the
\`[`Access-Control-Expose-Headers`](#http-access-control-expose-headers){#ref-for-http-access-control-expose-headers③
link-type="http-header"}\` response header.

Alternatively, if `bar.invalid` wanted to share all its response
headers, for requests that do not include
[credentials](#credentials){#ref-for-credentials①③ link-type="dfn"}, it
could use \``*`\` as value for the
\`[`Access-Control-Expose-Headers`](#http-access-control-expose-headers){#ref-for-http-access-control-expose-headers④
link-type="http-header"}\` response header. If the request would have
included [credentials](#credentials){#ref-for-credentials①④
link-type="dfn"}, the response header names would have to be listed
explicitly and \``*`\` could not be used.
:::

::: {#example-cors-with-credentials .example}
[](#example-cors-with-credentials){.self-link}

The developer of `foo.invalid` returns, now fetching some data from
`bar.invalid` while including
[credentials](#credentials){#ref-for-credentials①⑤ link-type="dfn"}.
This time around the [CORS
protocol](#cors-protocol){#ref-for-cors-protocol①⑤ link-type="dfn"} is
no longer transparent to the developer as
[credentials](#credentials){#ref-for-credentials①⑥ link-type="dfn"}
require an explicit opt-in:

``` {.lang-javascript .highlight}
fetch(url, { credentials:"include" }).then(success, failure)
```

This also makes any \``Set-Cookie`\` response headers `bar.invalid`
includes fully functional (they are ignored otherwise).

The user agent will make sure to include any relevant
[credentials](#credentials){#ref-for-credentials①⑦ link-type="dfn"} in
the request. It will also put stricter requirements on the response. Not
only will `bar.invalid` need to list \``https://foo.invalid`\` as value
for the
\`[`Access-Control-Allow-Origin`](#http-access-control-allow-origin){#ref-for-http-access-control-allow-origin⑤
link-type="http-header"}\` header (\``*`\` is not allowed when
[credentials](#credentials){#ref-for-credentials①⑧ link-type="dfn"} are
involved), the
\`[`Access-Control-Allow-Credentials`](#http-access-control-allow-credentials){#ref-for-http-access-control-allow-credentials③
link-type="http-header"}\` header has to be present too:

``` {.lang-http .highlight}
Access-Control-Allow-Origin: https://foo.invalid
Access-Control-Allow-Credentials: true
```

If the response does not include those two headers with those values,
the `failure` callback will be invoked. However, any \``Set-Cookie`\`
response headers will be respected.
:::

#### [3.2.7. ]{.secno}[CORS protocol exceptions]{.content}[](#cors-protocol-exceptions){.self-link} {#cors-protocol-exceptions .heading .settled level="3.2.7"}

Specifications have allowed limited exceptions to the CORS safelist for
non-safelisted \``Content-Type`\` header values. These exceptions are
made for requests that can be triggered by web content but whose headers
and bodies can be only minimally controlled by the web content.
Therefore, servers should expect cross-origin web content to be allowed
to trigger non-preflighted requests with the following non-safelisted
\``Content-Type`\` header values:

- \``application/csp-report`\`
  [\[CSP\]](#biblio-csp "Content Security Policy Level 3"){link-type="biblio"}
- \``application/expect-ct-report+json`\`
  [\[RFC9163\]](#biblio-rfc9163 "Expect-CT Extension for HTTP"){link-type="biblio"}
- \``application/xss-auditor-report`\`
- \``application/ocsp-request`\`
  [\[RFC6960\]](#biblio-rfc6960 "X.509 Internet Public Key Infrastructure Online Certificate Status Protocol - OCSP"){link-type="biblio"}

Specifications should avoid introducing new exceptions and should only
do so with careful consideration for the security consequences. New
exceptions can be proposed by [filing an
issue](https://github.com/whatwg/fetch/issues/new).

### [3.3. ]{.secno}[\``Content-Length`\` header]{.content}[](#content-length-header){.self-link} {#content-length-header .heading .settled level="3.3"}

The \``Content-Length`\` header is largely defined in HTTP. Its
processing model is defined here as the model defined in HTTP is not
compatible with web content.
[\[HTTP\]](#biblio-http "HTTP Semantics"){link-type="biblio"}

::: {.algorithm algorithm="extract a length" algorithm-for="header list"}
To [extract a length]{#header-list-extract-a-length .dfn .dfn-paneled
dfn-for="header list" dfn-type="dfn" export=""
lt="extract a length|extracting a length"} from a [header
list](#concept-header-list){#ref-for-concept-header-list①⑨
link-type="dfn"} `headers`{.variable}, run these steps:

1.  Let `values`{.variable} be the result of [getting, decoding, and
    splitting](#concept-header-list-get-decode-split){#ref-for-concept-header-list-get-decode-split②
    link-type="dfn"} \``Content-Length`\` from `headers`{.variable}.

2.  If `values`{.variable} is null, then return null.

3.  Let `candidateValue`{.variable} be null.

4.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate⑧
    link-type="dfn"} `value`{.variable} of `values`{.variable}:

    1.  If `candidateValue`{.variable} is null, then set
        `candidateValue`{.variable} to `value`{.variable}.

    2.  Otherwise, if `value`{.variable} is not
        `candidateValue`{.variable}, return failure.

5.  If `candidateValue`{.variable} is the empty string or has a [code
    point](https://infra.spec.whatwg.org/#code-point){#ref-for-code-point⑨
    link-type="dfn"} that is not an [ASCII
    digit](https://infra.spec.whatwg.org/#ascii-digit){#ref-for-ascii-digit②
    link-type="dfn"}, then return null.

6.  Return `candidateValue`{.variable}, interpreted as decimal number.
:::

### [3.4. ]{.secno}[\``Content-Type`\` header]{.content}[](#content-type-header){.self-link} {#content-type-header .heading .settled level="3.4"}

The \``Content-Type`\` header is largely defined in HTTP. Its processing
model is defined here as the model defined in HTTP is not compatible
with web content.
[\[HTTP\]](#biblio-http "HTTP Semantics"){link-type="biblio"}

::: {.algorithm algorithm="extract a MIME type" algorithm-for="header list"}
To [extract a MIME type]{#concept-header-extract-mime-type .dfn
.dfn-paneled dfn-for="header list" dfn-type="dfn" export=""
lt="extract a MIME type|extracting a MIME type"} from a [header
list](#concept-header-list){#ref-for-concept-header-list②⓪
link-type="dfn"} `headers`{.variable}, run these steps. They return
failure or a [MIME
type](https://mimesniff.spec.whatwg.org/#mime-type){#ref-for-mime-type①
link-type="dfn"}.

1.  Let `charset`{.variable} be null.

2.  Let `essence`{.variable} be null.

3.  Let `mimeType`{.variable} be null.

4.  Let `values`{.variable} be the result of [getting, decoding, and
    splitting](#concept-header-list-get-decode-split){#ref-for-concept-header-list-get-decode-split③
    link-type="dfn"} \``Content-Type`\` from `headers`{.variable}.

5.  If `values`{.variable} is null, then return failure.

6.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate⑨
    link-type="dfn"} `value`{.variable} of `values`{.variable}:

    1.  Let `temporaryMimeType`{.variable} be the result of
        [parsing](https://mimesniff.spec.whatwg.org/#parse-a-mime-type){#ref-for-parse-a-mime-type①
        link-type="dfn"} `value`{.variable}.

    2.  If `temporaryMimeType`{.variable} is failure or its
        [essence](https://mimesniff.spec.whatwg.org/#mime-type-essence){#ref-for-mime-type-essence③
        link-type="dfn"} is \"`*/*`\", then
        [continue](https://infra.spec.whatwg.org/#iteration-continue){#ref-for-iteration-continue④
        link-type="dfn"}.

    3.  Set `mimeType`{.variable} to `temporaryMimeType`{.variable}.

    4.  If `mimeType`{.variable}'s
        [essence](https://mimesniff.spec.whatwg.org/#mime-type-essence){#ref-for-mime-type-essence④
        link-type="dfn"} is not `essence`{.variable}, then:

        1.  Set `charset`{.variable} to null.

        2.  If `mimeType`{.variable}'s
            [parameters](https://mimesniff.spec.whatwg.org/#parameters){#ref-for-parameters
            link-type="dfn"}\[\"`charset`\"\]
            [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists
            link-type="dfn"}, then set `charset`{.variable} to
            `mimeType`{.variable}'s
            [parameters](https://mimesniff.spec.whatwg.org/#parameters){#ref-for-parameters①
            link-type="dfn"}\[\"`charset`\"\].

        3.  Set `essence`{.variable} to `mimeType`{.variable}'s
            [essence](https://mimesniff.spec.whatwg.org/#mime-type-essence){#ref-for-mime-type-essence⑤
            link-type="dfn"}.

    5.  Otherwise, if `mimeType`{.variable}'s
        [parameters](https://mimesniff.spec.whatwg.org/#parameters){#ref-for-parameters②
        link-type="dfn"}\[\"`charset`\"\] does not
        [exist](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists①
        link-type="dfn"}, and `charset`{.variable} is non-null, set
        `mimeType`{.variable}'s
        [parameters](https://mimesniff.spec.whatwg.org/#parameters){#ref-for-parameters③
        link-type="dfn"}\[\"`charset`\"\] to `charset`{.variable}.

7.  If `mimeType`{.variable} is null, then return failure.

8.  Return `mimeType`{.variable}.
:::

When [extract a MIME
type](#concept-header-extract-mime-type){#ref-for-concept-header-extract-mime-type③
link-type="dfn"} returns failure or a [MIME
type](https://mimesniff.spec.whatwg.org/#mime-type){#ref-for-mime-type②
link-type="dfn"} whose
[essence](https://mimesniff.spec.whatwg.org/#mime-type-essence){#ref-for-mime-type-essence⑥
link-type="dfn"} is incorrect for a given format, treat this as a fatal
error. Existing web platform features have not always followed this
pattern, which has been a major source of security vulnerabilities in
those features over the years. In contrast, a [MIME
type](https://mimesniff.spec.whatwg.org/#mime-type){#ref-for-mime-type③
link-type="dfn"}'s
[parameters](https://mimesniff.spec.whatwg.org/#parameters){#ref-for-parameters④
link-type="dfn"} can typically be safely ignored.

::: {#example-extract-a-mime-type .example}
[](#example-extract-a-mime-type){.self-link}

This is how [extract a MIME
type](#concept-header-extract-mime-type){#ref-for-concept-header-extract-mime-type④
link-type="dfn"} functions in practice:

Headers (as on the network)

Output
([serialized](https://mimesniff.spec.whatwg.org/#serialize-a-mime-type){#ref-for-serialize-a-mime-type
link-type="dfn"})

``` {.lang-http .highlight}
Content-Type: text/plain;charset=gbk, text/html
```

`text/html`

``` {.lang-http .highlight}
Content-Type: text/html;charset=gbk;a=b, text/html;x=y
```

`text/html;x=y;charset=gbk`

``` {.lang-http .highlight}
Content-Type: text/html;charset=gbk;a=b
Content-Type: text/html;x=y
```

``` {.lang-http .highlight}
Content-Type: text/html;charset=gbk
Content-Type: x/x
Content-Type: text/html;x=y
```

`text/html;x=y`

``` {.lang-http .highlight}
Content-Type: text/html
Content-Type: cannot-parse
```

`text/html`

``` {.lang-http .highlight}
Content-Type: text/html
Content-Type: */*
```

``` {.lang-http .highlight}
Content-Type: text/html
Content-Type:
```
:::

:::: {.algorithm algorithm="legacy extract an encoding"}
To [legacy extract an encoding]{#legacy-extract-an-encoding .dfn
.dfn-paneled dfn-type="dfn" export=""} given failure or a [MIME
type](https://mimesniff.spec.whatwg.org/#mime-type){#ref-for-mime-type④
link-type="dfn"} `mimeType`{.variable} and an
[encoding](https://encoding.spec.whatwg.org/#encoding){#ref-for-encoding
link-type="dfn"} `fallbackEncoding`{.variable}, run these steps:

1.  If `mimeType`{.variable} is failure, then return
    `fallbackEncoding`{.variable}.

2.  If `mimeType`{.variable}\[\"`charset`\"\] does not
    [exist](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists②
    link-type="dfn"}, then return `fallbackEncoding`{.variable}.

3.  Let `tentativeEncoding`{.variable} be the result of [getting an
    encoding](https://encoding.spec.whatwg.org/#concept-encoding-get){#ref-for-concept-encoding-get
    link-type="dfn"} from `mimeType`{.variable}\[\"`charset`\"\].

4.  If `tentativeEncoding`{.variable} is failure, then return
    `fallbackEncoding`{.variable}.

5.  Return `tentativeEncoding`{.variable}.

::: {.note role="note"}
This algorithm allows `mimeType`{.variable} to be failure so it can be
more easily combined with [extract a MIME
type](#concept-header-extract-mime-type){#ref-for-concept-header-extract-mime-type⑤
link-type="dfn"}.

It is denoted as legacy as modern formats are to exclusively use
[UTF-8](https://encoding.spec.whatwg.org/#utf-8){#ref-for-utf-8
link-type="dfn"}.
:::
::::

### [3.5. ]{.secno}[\``X-Content-Type-Options`\` header]{.content}[](#x-content-type-options-header){.self-link} {#x-content-type-options-header .heading .settled level="3.5"}

The \`[`X-Content-Type-Options`]{#http-x-content-type-options .dfn
.dfn-paneled dfn-type="http-header" export=""}\` response
[header](#concept-header){#ref-for-concept-header④⑤ link-type="dfn"} can
be used to require checking of a
[response](#concept-response){#ref-for-concept-response③⑦
link-type="dfn"}'s \``Content-Type`\`
[header](#concept-header){#ref-for-concept-header④⑥ link-type="dfn"}
against the
[destination](#concept-request-destination){#ref-for-concept-request-destination⑨
link-type="dfn"} of a
[request](#concept-request){#ref-for-concept-request⑨⑦ link-type="dfn"}.

::: {.algorithm algorithm="determine nosniff"}
To [determine nosniff]{#determine-nosniff .dfn .dfn-paneled
dfn-type="dfn" export=""}, given a [header
list](#concept-header-list){#ref-for-concept-header-list②①
link-type="dfn"} `list`{.variable}, run these steps:

1.  Let `values`{.variable} be the result of [getting, decoding, and
    splitting](#concept-header-list-get-decode-split){#ref-for-concept-header-list-get-decode-split④
    link-type="dfn"}
    \`[`X-Content-Type-Options`](#http-x-content-type-options){#ref-for-http-x-content-type-options
    link-type="http-header"}\` from `list`{.variable}.

2.  If `values`{.variable} is null, then return false.

3.  If `values`{.variable}\[0\] is an [ASCII
    case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#ref-for-ascii-case-insensitive
    link-type="dfn"} match for \"`nosniff`\", then return true.

4.  Return false.
:::

Web developers and conformance checkers must use the following
[value](#concept-header-value){#ref-for-concept-header-value①⑥
link-type="dfn"} [ABNF](#abnf){#ref-for-abnf④ link-type="dfn"} for
\`[`X-Content-Type-Options`](#http-x-content-type-options){#ref-for-http-x-content-type-options①
link-type="http-header"}\`:

``` {.lang-abnf .highlight}
X-Content-Type-Options           = "nosniff" ; case-insensitive
```

::: {.algorithm algorithm="should response to request be blocked due to nosniff"}
#### [3.5.1. ]{.secno}[Should `response`{.variable} to `request`{.variable} be blocked due to nosniff?]{.content}[](#should-response-to-request-be-blocked-due-to-nosniff?){#ref-for-should-response-to-request-be-blocked-due-to-nosniff? .self-link} {#should-response-to-request-be-blocked-due-to-nosniff? .heading .settled .dfn-paneled dfn-type="dfn" level="3.5.1" lt="should response to request be blocked due to nosniff" noexport=""}

Run these steps:

1.  If [determine
    nosniff](#determine-nosniff){#ref-for-determine-nosniff
    link-type="dfn"} with `response`{.variable}'s [header
    list](#concept-response-header-list){#ref-for-concept-response-header-list⑨
    link-type="dfn"} is false, then return **allowed**.

2.  Let `mimeType`{.variable} be the result of [extracting a MIME
    type](#concept-header-extract-mime-type){#ref-for-concept-header-extract-mime-type⑥
    link-type="dfn"} from `response`{.variable}'s [header
    list](#concept-response-header-list){#ref-for-concept-response-header-list①⓪
    link-type="dfn"}.

3.  Let `destination`{.variable} be `request`{.variable}'s
    [destination](#concept-request-destination){#ref-for-concept-request-destination①⓪
    link-type="dfn"}.

4.  If `destination`{.variable} is
    [script-like](#request-destination-script-like){#ref-for-request-destination-script-like②
    link-type="dfn"} and `mimeType`{.variable} is failure or is not a
    [JavaScript MIME
    type](https://mimesniff.spec.whatwg.org/#javascript-mime-type){#ref-for-javascript-mime-type
    link-type="dfn"}, then return **blocked**.

5.  If `destination`{.variable} is \"`style`\" and `mimeType`{.variable}
    is failure or its
    [essence](https://mimesniff.spec.whatwg.org/#mime-type-essence){#ref-for-mime-type-essence⑦
    link-type="dfn"} is not \"`text/css`\", then return **blocked**.

6.  Return **allowed**.

Only [request](#concept-request){#ref-for-concept-request⑨⑧
link-type="dfn"}
[destinations](#concept-request-destination){#ref-for-concept-request-destination①①
link-type="dfn"} that are
[script-like](#request-destination-script-like){#ref-for-request-destination-script-like③
link-type="dfn"} or \"`style`\" are considered as any exploits pertain
to them. Also, considering \"`image`\" was not compatible with deployed
content.
:::

### [3.6. ]{.secno}[\``Cross-Origin-Resource-Policy`\` header]{.content}[](#cross-origin-resource-policy-header){.self-link} {#cross-origin-resource-policy-header .heading .settled level="3.6"}

The
\`[`Cross-Origin-Resource-Policy`]{#http-cross-origin-resource-policy
.dfn .dfn-paneled dfn-type="http-header" export=""}\` response
[header](#concept-header){#ref-for-concept-header④⑦ link-type="dfn"} can
be used to require checking a
[request](#concept-request){#ref-for-concept-request⑨⑨
link-type="dfn"}'s [current
URL](#concept-request-current-url){#ref-for-concept-request-current-url⑦
link-type="dfn"}'s
[origin](https://url.spec.whatwg.org/#concept-url-origin){#ref-for-concept-url-origin①⓪
link-type="dfn"} against a
[request](#concept-request){#ref-for-concept-request①⓪⓪
link-type="dfn"}'s
[origin](#concept-request-origin){#ref-for-concept-request-origin①⓪
link-type="dfn"} when
[request](#concept-request){#ref-for-concept-request①⓪①
link-type="dfn"}'s
[mode](#concept-request-mode){#ref-for-concept-request-mode⑦
link-type="dfn"} is \"`no-cors`\".

Its [value](#concept-header-value){#ref-for-concept-header-value①⑦
link-type="dfn"} [ABNF](#abnf){#ref-for-abnf⑤ link-type="dfn"}:

``` {.lang-abnf .highlight}
Cross-Origin-Resource-Policy     = %s"same-origin" / %s"same-site" / %s"cross-origin" ; case-sensitive
```

::: {.algorithm algorithm="cross-origin resource policy check"}
To perform a [cross-origin resource policy
check]{#cross-origin-resource-policy-check .dfn .dfn-paneled
dfn-type="dfn" export=""}, given an
[origin](https://url.spec.whatwg.org/#concept-url-origin){#ref-for-concept-url-origin①①
link-type="dfn"} `origin`{.variable}, an [environment settings
object](https://html.spec.whatwg.org/multipage/webappapis.html#environment-settings-object){#ref-for-environment-settings-object⑥
link-type="dfn"} `settingsObject`{.variable}, a string
`destination`{.variable}, a
[response](#concept-response){#ref-for-concept-response③⑧
link-type="dfn"} `response`{.variable}, and an optional boolean
`forNavigation`{.variable}, run these steps:

1.  Set `forNavigation`{.variable} to false if it is not given.

2.  Let `embedderPolicy`{.variable} be `settingsObject`{.variable}'s
    [policy
    container](https://html.spec.whatwg.org/multipage/webappapis.html#concept-settings-object-policy-container){#ref-for-concept-settings-object-policy-container①
    link-type="dfn"}'s [embedder
    policy](https://html.spec.whatwg.org/multipage/browsers.html#policy-container-embedder-policy){#ref-for-policy-container-embedder-policy①
    link-type="dfn"}.

3.  If the [cross-origin resource policy internal
    check](#cross-origin-resource-policy-internal-check){#ref-for-cross-origin-resource-policy-internal-check
    link-type="dfn"} with `origin`{.variable},
    \"[`unsafe-none`](https://html.spec.whatwg.org/multipage/browsers.html#coep-unsafe-none){#ref-for-coep-unsafe-none
    link-type="dfn"}\", `response`{.variable}, and
    `forNavigation`{.variable} returns **blocked**, then return
    **blocked**.

    This step is needed because we don't want to report violations not
    related to Cross-Origin Embedder Policy below.

4.  If the [cross-origin resource policy internal
    check](#cross-origin-resource-policy-internal-check){#ref-for-cross-origin-resource-policy-internal-check①
    link-type="dfn"} with `origin`{.variable},
    `embedderPolicy`{.variable}'s [report only
    value](https://html.spec.whatwg.org/multipage/browsers.html#embedder-policy-report-only-value){#ref-for-embedder-policy-report-only-value
    link-type="dfn"}, `response`{.variable}, and
    `forNavigation`{.variable} returns **blocked**, then [queue a
    cross-origin embedder policy CORP violation
    report](#queue-a-cross-origin-embedder-policy-corp-violation-report){#ref-for-queue-a-cross-origin-embedder-policy-corp-violation-report
    link-type="dfn"} with `response`{.variable},
    `settingsObject`{.variable}, `destination`{.variable}, and true.

5.  If the [cross-origin resource policy internal
    check](#cross-origin-resource-policy-internal-check){#ref-for-cross-origin-resource-policy-internal-check②
    link-type="dfn"} with `origin`{.variable},
    `embedderPolicy`{.variable}'s
    [value](https://html.spec.whatwg.org/multipage/browsers.html#embedder-policy-value-2){#ref-for-embedder-policy-value-2①
    link-type="dfn"}, `response`{.variable}, and
    `forNavigation`{.variable} returns **allowed**, then return
    **allowed**.

6.  [Queue a cross-origin embedder policy CORP violation
    report](#queue-a-cross-origin-embedder-policy-corp-violation-report){#ref-for-queue-a-cross-origin-embedder-policy-corp-violation-report①
    link-type="dfn"} with `response`{.variable},
    `settingsObject`{.variable}, `destination`{.variable}, and false.

7.  Return **blocked**.

Only HTML's navigate algorithm uses this check with
`forNavigation`{.variable} set to true, and it's always for nested
navigations. Otherwise, `response`{.variable} is either the [internal
response](#concept-internal-response){#ref-for-concept-internal-response⑧
link-type="dfn"} of an [opaque filtered
response](#concept-filtered-response-opaque){#ref-for-concept-filtered-response-opaque③
link-type="dfn"} or a
[response](#concept-response){#ref-for-concept-response③⑨
link-type="dfn"} which will be the [internal
response](#concept-internal-response){#ref-for-concept-internal-response⑨
link-type="dfn"} of an [opaque filtered
response](#concept-filtered-response-opaque){#ref-for-concept-filtered-response-opaque④
link-type="dfn"}.
[\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}
:::

::: {.algorithm algorithm="cross-origin resource policy internal check"}
To perform a [cross-origin resource policy internal
check]{#cross-origin-resource-policy-internal-check .dfn .dfn-paneled
dfn-type="dfn" noexport=""}, given an
[origin](https://url.spec.whatwg.org/#concept-url-origin){#ref-for-concept-url-origin①②
link-type="dfn"} `origin`{.variable}, an [embedder policy
value](https://html.spec.whatwg.org/multipage/browsers.html#embedder-policy-value){#ref-for-embedder-policy-value
link-type="dfn"} `embedderPolicyValue`{.variable}, a
[response](#concept-response){#ref-for-concept-response④⓪
link-type="dfn"} `response`{.variable}, and a boolean
`forNavigation`{.variable}, run these steps:

1.  If `forNavigation`{.variable} is true and
    `embedderPolicyValue`{.variable} is
    \"[`unsafe-none`](https://html.spec.whatwg.org/multipage/browsers.html#coep-unsafe-none){#ref-for-coep-unsafe-none①
    link-type="dfn"}\", then return **allowed**.

2.  Let `policy`{.variable} be the result of
    [getting](#concept-header-list-get){#ref-for-concept-header-list-get③
    link-type="dfn"}
    \`[`Cross-Origin-Resource-Policy`](#http-cross-origin-resource-policy){#ref-for-http-cross-origin-resource-policy
    link-type="http-header"}\` from `response`{.variable}'s [header
    list](#concept-response-header-list){#ref-for-concept-response-header-list①①
    link-type="dfn"}.

    This means that
    \``Cross-Origin-Resource-Policy: same-site, same-origin`\` ends up
    as **allowed** below as it will never match anything, as long as
    `embedderPolicyValue`{.variable} is
    \"[`unsafe-none`](https://html.spec.whatwg.org/multipage/browsers.html#coep-unsafe-none){#ref-for-coep-unsafe-none②
    link-type="dfn"}\". Two or more
    \`[`Cross-Origin-Resource-Policy`](#http-cross-origin-resource-policy){#ref-for-http-cross-origin-resource-policy①
    link-type="http-header"}\` headers will have the same effect.

3.  If `policy`{.variable} is neither \``same-origin`\`,
    \``same-site`\`, nor \``cross-origin`\`, then set
    `policy`{.variable} to null.

4.  If `policy`{.variable} is null, then switch on
    `embedderPolicyValue`{.variable}:

    \"[`unsafe-none`](https://html.spec.whatwg.org/multipage/browsers.html#coep-unsafe-none){#ref-for-coep-unsafe-none③ link-type="dfn"}\"

    :   Do nothing.

    \"[`credentialless`](https://html.spec.whatwg.org/multipage/browsers.html#coep-credentialless){#ref-for-coep-credentialless① link-type="dfn"}\"

    :   Set `policy`{.variable} to \``same-origin`\` if:

        - `response`{.variable}'s
          [request-includes-credentials](#response-request-includes-credentials){#ref-for-response-request-includes-credentials
          link-type="dfn"} is true, or
        - `forNavigation`{.variable} is true.

    \"[`require-corp`](https://html.spec.whatwg.org/multipage/browsers.html#coep-require-corp){#ref-for-coep-require-corp link-type="dfn"}\"

    :   Set `policy`{.variable} to \``same-origin`\`.

5.  Switch on `policy`{.variable}:

    null\
    \``cross-origin`\`

    :   Return **allowed**.

    \``same-origin`\`

    :   If `origin`{.variable} is [same
        origin](https://html.spec.whatwg.org/multipage/browsers.html#same-origin){#ref-for-same-origin④
        link-type="dfn"} with `response`{.variable}'s
        [URL](#concept-response-url){#ref-for-concept-response-url⑤
        link-type="dfn"}'s
        [origin](https://url.spec.whatwg.org/#concept-url-origin){#ref-for-concept-url-origin①③
        link-type="dfn"}, then return **allowed**.

        Otherwise, return **blocked**.

    \``same-site`\`

    :   If all of the following are true

        - `origin`{.variable} is [schemelessly same
          site](https://html.spec.whatwg.org/multipage/browsers.html#schemelessly-same-site){#ref-for-schemelessly-same-site
          link-type="dfn"} with `response`{.variable}'s
          [URL](#concept-response-url){#ref-for-concept-response-url⑥
          link-type="dfn"}'s
          [origin](https://url.spec.whatwg.org/#concept-url-origin){#ref-for-concept-url-origin①④
          link-type="dfn"}

        - `origin`{.variable}'s
          [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#ref-for-concept-url-scheme②
          link-type="dfn"} is \"`https`\" or `response`{.variable}'s
          [URL](#concept-response-url){#ref-for-concept-response-url⑦
          link-type="dfn"}'s
          [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#ref-for-concept-url-scheme③
          link-type="dfn"} is not \"`https`\"

        then return **allowed**.

        Otherwise, return **blocked**.

        \``Cross-Origin-Resource-Policy: same-site`\` does not consider
        a response delivered via a secure transport to match a
        non-secure requesting origin, even if their hosts are otherwise
        same site. Securely-transported responses will only match a
        securely-transported initiator.
:::

::: {.algorithm algorithm="queue a cross-origin embedder policy CORP violation report"}
To [queue a cross-origin embedder policy CORP violation
report]{#queue-a-cross-origin-embedder-policy-corp-violation-report .dfn
.dfn-paneled dfn-type="dfn" noexport=""}, given a
[response](#concept-response){#ref-for-concept-response④①
link-type="dfn"} `response`{.variable}, an [environment settings
object](https://html.spec.whatwg.org/multipage/webappapis.html#environment-settings-object){#ref-for-environment-settings-object⑦
link-type="dfn"} `settingsObject`{.variable}, a string
`destination`{.variable}, and a boolean `reportOnly`{.variable}, run
these steps:

1.  Let `endpoint`{.variable} be `settingsObject`{.variable}'s [policy
    container](https://html.spec.whatwg.org/multipage/webappapis.html#concept-settings-object-policy-container){#ref-for-concept-settings-object-policy-container②
    link-type="dfn"}'s [embedder
    policy](https://html.spec.whatwg.org/multipage/browsers.html#policy-container-embedder-policy){#ref-for-policy-container-embedder-policy②
    link-type="dfn"}'s [report only reporting
    endpoint](https://html.spec.whatwg.org/multipage/browsers.html#embedder-policy-report-only-reporting-endpoint){#ref-for-embedder-policy-report-only-reporting-endpoint
    link-type="dfn"} if `reportOnly`{.variable} is true and
    `settingsObject`{.variable}'s [policy
    container](https://html.spec.whatwg.org/multipage/webappapis.html#concept-settings-object-policy-container){#ref-for-concept-settings-object-policy-container③
    link-type="dfn"}'s [embedder
    policy](https://html.spec.whatwg.org/multipage/browsers.html#policy-container-embedder-policy){#ref-for-policy-container-embedder-policy③
    link-type="dfn"}'s [reporting
    endpoint](https://html.spec.whatwg.org/multipage/browsers.html#embedder-policy-reporting-endpoint){#ref-for-embedder-policy-reporting-endpoint
    link-type="dfn"} otherwise.

2.  Let `serializedURL`{.variable} be the result of [serializing a
    response URL for
    reporting](#serialize-a-response-url-for-reporting){#ref-for-serialize-a-response-url-for-reporting
    link-type="dfn"} with `response`{.variable}.

3.  Let `disposition`{.variable} be \"`reporting`\" if
    `reportOnly`{.variable} is true; otherwise \"`enforce`\".

4.  Let `body`{.variable} be a new object containing the following
    properties:

    key

    value

    \"`type`\"

    \"`corp`\"

    \"`blockedURL`\"

    `serializedURL`{.variable}

    \"`destination`\"

    `destination`{.variable}

    \"`disposition`\"

    `disposition`{.variable}

5.  [Generate and queue a
    report](https://w3c.github.io/reporting/#generate-and-queue-a-report){#ref-for-generate-and-queue-a-report
    link-type="dfn"} for `settingsObject`{.variable}'s [global
    object](https://html.spec.whatwg.org/multipage/webappapis.html#concept-settings-object-global){#ref-for-concept-settings-object-global①
    link-type="dfn"} given the [\"`coep`\" report
    type](https://html.spec.whatwg.org/multipage/browsers.html#coep-report-type){#ref-for-coep-report-type
    link-type="dfn"}, `endpoint`{.variable}, and `body`{.variable}.
    [\[REPORTING\]](#biblio-reporting "Reporting API"){link-type="biblio"}
:::

### [3.7. ]{.secno}[\``Sec-Purpose`\` header]{.content}[](#sec-purpose-header){.self-link} {#sec-purpose-header .heading .settled level="3.7"}

The \`[`Sec-Purpose`]{#http-sec-purpose .dfn .dfn-paneled
dfn-type="http-header" export=""}\` HTTP request header specifies that
the request serves one or more purposes other than requesting the
resource for immediate use by the user.

The \`[`Sec-Purpose`](#http-sec-purpose){#ref-for-http-sec-purpose
link-type="http-header"}\` header field is a [structured
header](https://httpwg.org/specs/rfc9651.html#){#ref-for-something
link-type="dfn"} whose value must be a
[token](https://httpwg.org/specs/rfc9651.html#token){#ref-for-token
link-type="dfn"}.

The sole
[token](https://httpwg.org/specs/rfc9651.html#token){#ref-for-token①
link-type="dfn"} defined is `prefetch`. It indicates the request's
purpose is to fetch a resource that is anticipated to be needed shortly.

The server can use this to adjust the caching expiry for prefetches, to
disallow the prefetch, or to treat it differently when counting page
visits.

## [4. ]{.secno}[Fetching]{.content}[](#fetching){.self-link} {#fetching .heading .settled level="4"}

The algorithm below defines
[fetching](#concept-fetch){#ref-for-concept-fetch②③ link-type="dfn"}. In
broad strokes, it takes a
[request](#concept-request){#ref-for-concept-request①⓪② link-type="dfn"}
and one or more algorithms to run at various points during the
operation. A [response](#concept-response){#ref-for-concept-response④②
link-type="dfn"} is passed to the last two algorithms listed below. The
first two algorithms can be used to capture uploads.

::: {.algorithm algorithm="fetch"}
To [fetch]{#concept-fetch .dfn .dfn-paneled dfn-type="dfn" export=""},
given a [request](#concept-request){#ref-for-concept-request①⓪③
link-type="dfn"} `request`{.variable}, an optional algorithm
[`processRequestBodyChunkLength`{.variable}]{#process-request-body .dfn
.dfn-paneled dfn-for="fetch" dfn-type="dfn" export=""}, an optional
algorithm [[]{#process-request-end-of-file
.bs-old-id}`processRequestEndOfBody`{.variable}]{#process-request-end-of-body
.dfn .dfn-paneled dfn-for="fetch" dfn-type="dfn" export=""}, an optional
algorithm
[`processEarlyHintsResponse`{.variable}]{#fetch-processearlyhintsresponse
.dfn .dfn-paneled dfn-for="fetch" dfn-type="dfn" export=""}, an optional
algorithm [`processResponse`{.variable}]{#process-response .dfn
.dfn-paneled dfn-for="fetch" dfn-type="dfn" export=""}, an optional
algorithm
[`processResponseEndOfBody`{.variable}]{#fetch-processresponseendofbody
.dfn .dfn-paneled dfn-for="fetch" dfn-type="dfn" export=""}, an optional
algorithm [[]{#process-response-end-of-file
.bs-old-id}`processResponseConsumeBody`{.variable}]{#process-response-end-of-body
.dfn .dfn-paneled dfn-for="fetch" dfn-type="dfn" export=""}, and an
optional boolean [`useParallelQueue`{.variable}]{#fetch-useparallelqueue
.dfn .dfn-paneled dfn-for="fetch" dfn-type="dfn" export=""} (default
false), run the steps below. If given,
`processRequestBodyChunkLength`{.variable} must be an algorithm
accepting an integer representing the number of bytes transmitted. If
given, `processRequestEndOfBody`{.variable} must be an algorithm
accepting no arguments. If given, `processEarlyHintsResponse`{.variable}
must be an algorithm accepting a
[response](#concept-response){#ref-for-concept-response④③
link-type="dfn"}. If given, `processResponse`{.variable} must be an
algorithm accepting a
[response](#concept-response){#ref-for-concept-response④④
link-type="dfn"}. If given, `processResponseEndOfBody`{.variable} must
be an algorithm accepting a
[response](#concept-response){#ref-for-concept-response④⑤
link-type="dfn"}. If given, `processResponseConsumeBody`{.variable} must
be an algorithm accepting a
[response](#concept-response){#ref-for-concept-response④⑥
link-type="dfn"} and null, failure, or a [byte
sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence①⑤
link-type="dfn"}.

The user agent may be asked to [suspend]{#concept-fetch-suspend .dfn
.dfn-paneled dfn-for="fetch" dfn-type="dfn" export=""} the ongoing
fetch. The user agent may either accept or ignore the suspension
request. The suspended fetch can be [resumed]{#concept-fetch-resume .dfn
.dfn-paneled dfn-for="fetch" dfn-type="dfn" export=""}. The user agent
should ignore the suspension request if the ongoing fetch is updating
the response in the HTTP cache for the request.

The user agent does not update the entry in the HTTP cache for a
[request](#concept-request){#ref-for-concept-request①⓪④ link-type="dfn"}
if request's cache mode is \"no-store\" or a
\``Cache-Control: no-store`\` header appears in the response.
[\[HTTP-CACHING\]](#biblio-http-caching "HTTP Caching"){link-type="biblio"}

1.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert①⑦
    link-type="dfn"}: `request`{.variable}'s
    [mode](#concept-request-mode){#ref-for-concept-request-mode⑧
    link-type="dfn"} is \"`navigate`\" or
    `processEarlyHintsResponse`{.variable} is null.

    Processing of early hints
    ([responses](#concept-response){#ref-for-concept-response④⑦
    link-type="dfn"} whose
    [status](#concept-response-status){#ref-for-concept-response-status④
    link-type="dfn"} is 103) is only vetted for navigations.

2.  Let `taskDestination`{.variable} be null.

3.  Let `crossOriginIsolatedCapability`{.variable} be false.

4.  [Populate request from
    client](#populate-request-from-client){#ref-for-populate-request-from-client
    link-type="dfn"} given `request`{.variable}.

5.  If `request`{.variable}'s
    [client](#concept-request-client){#ref-for-concept-request-client⑥
    link-type="dfn"} is non-null, then:

    1.  Set `taskDestination`{.variable} to `request`{.variable}'s
        [client](#concept-request-client){#ref-for-concept-request-client⑦
        link-type="dfn"}'s [global
        object](https://html.spec.whatwg.org/multipage/webappapis.html#concept-settings-object-global){#ref-for-concept-settings-object-global②
        link-type="dfn"}.

    2.  Set `crossOriginIsolatedCapability`{.variable} to
        `request`{.variable}'s
        [client](#concept-request-client){#ref-for-concept-request-client⑧
        link-type="dfn"}'s [cross-origin isolated
        capability](https://html.spec.whatwg.org/multipage/webappapis.html#concept-settings-object-cross-origin-isolated-capability){#ref-for-concept-settings-object-cross-origin-isolated-capability
        link-type="dfn"}.

6.  If `useParallelQueue`{.variable} is true, then set
    `taskDestination`{.variable} to the result of [starting a new
    parallel
    queue](https://html.spec.whatwg.org/multipage/infrastructure.html#starting-a-new-parallel-queue){#ref-for-starting-a-new-parallel-queue②
    link-type="dfn"}.

7.  Let `timingInfo`{.variable} be a new [fetch timing
    info](#fetch-timing-info){#ref-for-fetch-timing-info④
    link-type="dfn"} whose [start
    time](#fetch-timing-info-start-time){#ref-for-fetch-timing-info-start-time②
    link-type="dfn"} and [post-redirect start
    time](#fetch-timing-info-post-redirect-start-time){#ref-for-fetch-timing-info-post-redirect-start-time①
    link-type="dfn"} are the [coarsened shared current
    time](https://w3c.github.io/hr-time/#dfn-coarsened-shared-current-time){#ref-for-dfn-coarsened-shared-current-time
    link-type="dfn"} given `crossOriginIsolatedCapability`{.variable},
    and
    [render-blocking](#fetch-timing-info-render-blocking){#ref-for-fetch-timing-info-render-blocking
    link-type="dfn"} is set to `request`{.variable}'s
    [render-blocking](#request-render-blocking){#ref-for-request-render-blocking
    link-type="dfn"}.

8.  Let `fetchParams`{.variable} be a new [fetch
    params](#fetch-params){#ref-for-fetch-params③ link-type="dfn"} whose
    [request](#fetch-params-request){#ref-for-fetch-params-request
    link-type="dfn"} is `request`{.variable}, [timing
    info](#fetch-params-timing-info){#ref-for-fetch-params-timing-info
    link-type="dfn"} is `timingInfo`{.variable}, [process request body
    chunk
    length](#fetch-params-process-request-body){#ref-for-fetch-params-process-request-body
    link-type="dfn"} is `processRequestBodyChunkLength`{.variable},
    [process request
    end-of-body](#fetch-params-process-request-end-of-body){#ref-for-fetch-params-process-request-end-of-body
    link-type="dfn"} is `processRequestEndOfBody`{.variable}, [process
    early hints
    response](#fetch-params-process-early-hints-response){#ref-for-fetch-params-process-early-hints-response
    link-type="dfn"} is `processEarlyHintsResponse`{.variable}, [process
    response](#fetch-params-process-response){#ref-for-fetch-params-process-response
    link-type="dfn"} is `processResponse`{.variable}, [process response
    consume
    body](#fetch-params-process-response-consume-body){#ref-for-fetch-params-process-response-consume-body
    link-type="dfn"} is `processResponseConsumeBody`{.variable},
    [process response
    end-of-body](#fetch-params-process-response-end-of-body){#ref-for-fetch-params-process-response-end-of-body
    link-type="dfn"} is `processResponseEndOfBody`{.variable}, [task
    destination](#fetch-params-task-destination){#ref-for-fetch-params-task-destination
    link-type="dfn"} is `taskDestination`{.variable}, and [cross-origin
    isolated
    capability](#fetch-params-cross-origin-isolated-capability){#ref-for-fetch-params-cross-origin-isolated-capability
    link-type="dfn"} is `crossOriginIsolatedCapability`{.variable}.

9.  If `request`{.variable}'s
    [body](#concept-request-body){#ref-for-concept-request-body④
    link-type="dfn"} is a [byte
    sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence①⑥
    link-type="dfn"}, then set `request`{.variable}'s
    [body](#concept-request-body){#ref-for-concept-request-body⑤
    link-type="dfn"} to `request`{.variable}'s
    [body](#concept-request-body){#ref-for-concept-request-body⑥
    link-type="dfn"} [as a
    body](#byte-sequence-as-a-body){#ref-for-byte-sequence-as-a-body
    link-type="dfn"}.

10. If all of the following conditions are true:

    - `request`{.variable}'s
      [URL](#concept-request-url){#ref-for-concept-request-url②
      link-type="dfn"}'s
      [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#ref-for-concept-url-scheme④
      link-type="dfn"} is an [HTTP(S)
      scheme](#http-scheme){#ref-for-http-scheme④ link-type="dfn"}

    - `request`{.variable}'s
      [mode](#concept-request-mode){#ref-for-concept-request-mode⑨
      link-type="dfn"} is \"`same-origin`\", \"`cors`\", or
      \"`no-cors`\"

    - `request`{.variable}'s
      [client](#concept-request-client){#ref-for-concept-request-client⑨
      link-type="dfn"} is not null, and `request`{.variable}'s
      [client](#concept-request-client){#ref-for-concept-request-client①⓪
      link-type="dfn"}'s [global
      object](https://html.spec.whatwg.org/multipage/webappapis.html#concept-settings-object-global){#ref-for-concept-settings-object-global③
      link-type="dfn"} is a
      [`Window`{.idl}](https://html.spec.whatwg.org/multipage/nav-history-apis.html#window){#ref-for-window
      link-type="idl"} object

    - `request`{.variable}'s
      [method](#concept-request-method){#ref-for-concept-request-method④
      link-type="dfn"} is \``GET`\`

    - `request`{.variable}'s [unsafe-request
      flag](#unsafe-request-flag){#ref-for-unsafe-request-flag①
      link-type="dfn"} is not set or `request`{.variable}'s [header
      list](#concept-request-header-list){#ref-for-concept-request-header-list⑤
      link-type="dfn"} [is
      empty](https://infra.spec.whatwg.org/#list-is-empty){#ref-for-list-is-empty②
      link-type="dfn"}

    then:

    1.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert①⑧
        link-type="dfn"}: `request`{.variable}'s
        [origin](#concept-request-origin){#ref-for-concept-request-origin①①
        link-type="dfn"} is [same
        origin](https://html.spec.whatwg.org/multipage/browsers.html#same-origin){#ref-for-same-origin⑤
        link-type="dfn"} with `request`{.variable}'s
        [client](#concept-request-client){#ref-for-concept-request-client①①
        link-type="dfn"}'s
        [origin](https://html.spec.whatwg.org/multipage/webappapis.html#concept-settings-object-origin){#ref-for-concept-settings-object-origin
        link-type="dfn"}.

    2.  Let `onPreloadedResponseAvailable`{.variable} be an algorithm
        that runs the following step given a
        [response](#concept-response){#ref-for-concept-response④⑧
        link-type="dfn"} `response`{.variable}: set
        `fetchParams`{.variable}'s [preloaded response
        candidate](#fetch-params-preloaded-response-candidate){#ref-for-fetch-params-preloaded-response-candidate
        link-type="dfn"} to `response`{.variable}.

    3.  Let `foundPreloadedResource`{.variable} be the result of
        invoking [consume a preloaded
        resource](https://html.spec.whatwg.org/multipage/links.html#consume-a-preloaded-resource){#ref-for-consume-a-preloaded-resource
        link-type="dfn"} for `request`{.variable}'s
        [client](#concept-request-client){#ref-for-concept-request-client①②
        link-type="dfn"}, given `request`{.variable}'s
        [URL](#concept-request-url){#ref-for-concept-request-url③
        link-type="dfn"}, `request`{.variable}'s
        [destination](#concept-request-destination){#ref-for-concept-request-destination①②
        link-type="dfn"}, `request`{.variable}'s
        [mode](#concept-request-mode){#ref-for-concept-request-mode①⓪
        link-type="dfn"}, `request`{.variable}'s [credentials
        mode](#concept-request-credentials-mode){#ref-for-concept-request-credentials-mode⑧
        link-type="dfn"}, `request`{.variable}'s [integrity
        metadata](#concept-request-integrity-metadata){#ref-for-concept-request-integrity-metadata
        link-type="dfn"}, and `onPreloadedResponseAvailable`{.variable}.

    4.  If `foundPreloadedResource`{.variable} is true and
        `fetchParams`{.variable}'s [preloaded response
        candidate](#fetch-params-preloaded-response-candidate){#ref-for-fetch-params-preloaded-response-candidate①
        link-type="dfn"} is null, then set `fetchParams`{.variable}'s
        [preloaded response
        candidate](#fetch-params-preloaded-response-candidate){#ref-for-fetch-params-preloaded-response-candidate②
        link-type="dfn"} to \"`pending`\".

11. If `request`{.variable}'s [header
    list](#concept-request-header-list){#ref-for-concept-request-header-list⑥
    link-type="dfn"} [does not
    contain](#header-list-contains){#ref-for-header-list-contains⑧
    link-type="dfn"} \``Accept`\`, then:

    1.  Let `value`{.variable} be \``*/*`\`.

    2.  If `request`{.variable}'s
        [initiator](#concept-request-initiator){#ref-for-concept-request-initiator③
        link-type="dfn"} is \"`prefetch`\", then set `value`{.variable}
        to the [document \``Accept`\` header
        value](#document-accept-header-value){#ref-for-document-accept-header-value
        link-type="dfn"}.

    3.  Otherwise, the user agent should set `value`{.variable} to the
        first matching statement, if any, switching on
        `request`{.variable}'s
        [destination](#concept-request-destination){#ref-for-concept-request-destination①③
        link-type="dfn"}:

        \"`document`\"\
        \"`frame`\"\
        \"`iframe`\"
        :   the [document \``Accept`\` header
            value](#document-accept-header-value){#ref-for-document-accept-header-value①
            link-type="dfn"}

        \"`image`\"
        :   \``image/png,image/svg+xml,image/*;q=0.8,*/*;q=0.5`\`

        \"`json`\"
        :   \``application/json,*/*;q=0.5`\`

        \"`style`\"
        :   \``text/css,*/*;q=0.1`\`

    4.  [Append](#concept-header-list-append){#ref-for-concept-header-list-append③
        link-type="dfn"} (\``Accept`\`, `value`{.variable}) to
        `request`{.variable}'s [header
        list](#concept-request-header-list){#ref-for-concept-request-header-list⑦
        link-type="dfn"}.

12. If `request`{.variable}'s [header
    list](#concept-request-header-list){#ref-for-concept-request-header-list⑧
    link-type="dfn"} [does not
    contain](#header-list-contains){#ref-for-header-list-contains⑨
    link-type="dfn"} \``Accept-Language`\`, then user agents should
    [append](#concept-header-list-append){#ref-for-concept-header-list-append④
    link-type="dfn"} (\``Accept-Language`, an appropriate [header
    value](#header-value){#ref-for-header-value⑧ link-type="dfn"}) to
    `request`{.variable}'s [header
    list](#concept-request-header-list){#ref-for-concept-request-header-list⑨
    link-type="dfn"}.

13. If `request`{.variable}'s [internal
    priority](#request-internal-priority){#ref-for-request-internal-priority
    link-type="dfn"} is null, then use `request`{.variable}'s
    [priority](#request-priority){#ref-for-request-priority
    link-type="dfn"},
    [initiator](#concept-request-initiator){#ref-for-concept-request-initiator④
    link-type="dfn"},
    [destination](#concept-request-destination){#ref-for-concept-request-destination①④
    link-type="dfn"}, and
    [render-blocking](#request-render-blocking){#ref-for-request-render-blocking①
    link-type="dfn"} in an
    [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#ref-for-implementation-defined①①
    link-type="dfn"} manner to set `request`{.variable}'s [internal
    priority](#request-internal-priority){#ref-for-request-internal-priority①
    link-type="dfn"} to an
    [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#ref-for-implementation-defined①②
    link-type="dfn"} object.

    The
    [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#ref-for-implementation-defined①③
    link-type="dfn"} object could encompass stream weight and dependency
    for HTTP/2, priorities used in Extensible Prioritization Scheme for
    HTTP for transports where it applies (including HTTP/3), and
    equivalent information used to prioritize dispatch and processing of
    HTTP/1 fetches.
    [\[RFC9218\]](#biblio-rfc9218 "Extensible Prioritization Scheme for HTTP"){link-type="biblio"}

14. If `request`{.variable} is a [subresource
    request](#subresource-request){#ref-for-subresource-request
    link-type="dfn"}, then:

    1.  Let `record`{.variable} be a new [fetch
        record](#concept-fetch-record){#ref-for-concept-fetch-record⑤
        link-type="dfn"} whose
        [request](#concept-fetch-record-request){#ref-for-concept-fetch-record-request①
        link-type="dfn"} is `request`{.variable} and
        [controller](#concept-fetch-record-fetch){#ref-for-concept-fetch-record-fetch②
        link-type="dfn"} is `fetchParams`{.variable}'s
        [controller](#fetch-params-controller){#ref-for-fetch-params-controller②
        link-type="dfn"}.

    2.  Append `record`{.variable} to `request`{.variable}'s
        [client](#concept-request-client){#ref-for-concept-request-client①③
        link-type="dfn"}'s [fetch
        group](#concept-fetch-group){#ref-for-concept-fetch-group②
        link-type="dfn"} list of [fetch
        records](#concept-fetch-record){#ref-for-concept-fetch-record⑥
        link-type="dfn"}.

15. Run [main fetch](#concept-main-fetch){#ref-for-concept-main-fetch
    link-type="dfn"} given `fetchParams`{.variable}.

16. Return `fetchParams`{.variable}'s
    [controller](#fetch-params-controller){#ref-for-fetch-params-controller③
    link-type="dfn"}.
:::

::: {.algorithm algorithm="populate request from client"}
To [populate request from client]{#populate-request-from-client .dfn
.dfn-paneled dfn-type="dfn" noexport=""} given a
[request](#concept-request){#ref-for-concept-request①⓪⑤ link-type="dfn"}
`request`{.variable}:

1.  If `request`{.variable}'s [traversable for user
    prompts](#concept-request-window){#ref-for-concept-request-window③
    link-type="dfn"} is \"`client`\":

    1.  Set `request`{.variable}'s [traversable for user
        prompts](#concept-request-window){#ref-for-concept-request-window④
        link-type="dfn"} to \"`no-traversable`\".

    2.  If `request`{.variable}'s
        [client](#concept-request-client){#ref-for-concept-request-client①④
        link-type="dfn"} is non-null:

        1.  Let `global`{.variable} be `request`{.variable}'s
            [client](#concept-request-client){#ref-for-concept-request-client①⑤
            link-type="dfn"}'s [global
            object](https://html.spec.whatwg.org/multipage/webappapis.html#concept-settings-object-global){#ref-for-concept-settings-object-global④
            link-type="dfn"}.

        2.  If `global`{.variable} is a
            [`Window`{.idl}](https://html.spec.whatwg.org/multipage/nav-history-apis.html#window){#ref-for-window①
            link-type="idl"} object and `global`{.variable}'s
            [navigable](https://html.spec.whatwg.org/multipage/nav-history-apis.html#window-navigable){#ref-for-window-navigable
            link-type="dfn"} is not null, then set
            `request`{.variable}'s [traversable for user
            prompts](#concept-request-window){#ref-for-concept-request-window⑤
            link-type="dfn"} to `global`{.variable}'s
            [navigable](https://html.spec.whatwg.org/multipage/nav-history-apis.html#window-navigable){#ref-for-window-navigable①
            link-type="dfn"}'s [traversable
            navigable](https://html.spec.whatwg.org/multipage/document-sequences.html#nav-traversable){#ref-for-nav-traversable
            link-type="dfn"}.

2.  If `request`{.variable}'s
    [origin](#concept-request-origin){#ref-for-concept-request-origin①②
    link-type="dfn"} is \"`client`\":

    1.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert①⑨
        link-type="dfn"}: `request`{.variable}'s
        [client](#concept-request-client){#ref-for-concept-request-client①⑥
        link-type="dfn"} is non-null.

    2.  Set `request`{.variable}'s
        [origin](#concept-request-origin){#ref-for-concept-request-origin①③
        link-type="dfn"} to `request`{.variable}'s
        [client](#concept-request-client){#ref-for-concept-request-client①⑦
        link-type="dfn"}'s
        [origin](https://html.spec.whatwg.org/multipage/webappapis.html#concept-settings-object-origin){#ref-for-concept-settings-object-origin①
        link-type="dfn"}.

3.  If `request`{.variable}'s [policy
    container](#concept-request-policy-container){#ref-for-concept-request-policy-container①
    link-type="dfn"} is \"`client`\":

    1.  If `request`{.variable}'s
        [client](#concept-request-client){#ref-for-concept-request-client①⑧
        link-type="dfn"} is non-null, then set `request`{.variable}'s
        [policy
        container](#concept-request-policy-container){#ref-for-concept-request-policy-container②
        link-type="dfn"} to a
        [clone](https://html.spec.whatwg.org/multipage/browsers.html#clone-a-policy-container){#ref-for-clone-a-policy-container
        link-type="dfn"} of `request`{.variable}'s
        [client](#concept-request-client){#ref-for-concept-request-client①⑨
        link-type="dfn"}'s [policy
        container](https://html.spec.whatwg.org/multipage/webappapis.html#concept-settings-object-policy-container){#ref-for-concept-settings-object-policy-container④
        link-type="dfn"}.
        [\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}

    2.  Otherwise, set `request`{.variable}'s [policy
        container](#concept-request-policy-container){#ref-for-concept-request-policy-container③
        link-type="dfn"} to a new [policy
        container](https://html.spec.whatwg.org/multipage/browsers.html#policy-container){#ref-for-policy-container②
        link-type="dfn"}.
:::

### [4.1. ]{.secno}[Main fetch]{.content}[](#main-fetch){.self-link} {#main-fetch .heading .settled level="4.1"}

::: {.algorithm algorithm="main fetch"}
To [main fetch]{#concept-main-fetch .dfn .dfn-paneled dfn-type="dfn"
noexport=""}, given a [fetch
params](#fetch-params){#ref-for-fetch-params④ link-type="dfn"}
`fetchParams`{.variable} and an optional boolean `recursive`{.variable}
(default false), run these steps:

1.  Let `request`{.variable} be `fetchParams`{.variable}'s
    [request](#fetch-params-request){#ref-for-fetch-params-request①
    link-type="dfn"}.

2.  Let `response`{.variable} be null.

3.  If `request`{.variable}'s [local-URLs-only
    flag](#local-urls-only-flag){#ref-for-local-urls-only-flag
    link-type="dfn"} is set and `request`{.variable}'s [current
    URL](#concept-request-current-url){#ref-for-concept-request-current-url⑧
    link-type="dfn"} is not [local](#is-local){#ref-for-is-local
    link-type="dfn"}, then set `response`{.variable} to a [network
    error](#concept-network-error){#ref-for-concept-network-error⑨
    link-type="dfn"}.

4.  Run [report Content Security Policy violations for
    `request`{.variable}](https://w3c.github.io/webappsec-csp/#report-for-request){#ref-for-report-for-request
    link-type="dfn"}.

5.  [Upgrade `request`{.variable} to a potentially trustworthy URL, if
    appropriate](https://w3c.github.io/webappsec-upgrade-insecure-requests/#upgrade-request){#ref-for-upgrade-request
    link-type="dfn"}.

6.  [Upgrade a mixed content `request`{.variable} to a potentially
    trustworthy URL, if
    appropriate](https://w3c.github.io/webappsec-mixed-content/#upgrade-algorithm){#ref-for-upgrade-algorithm
    link-type="dfn"}.

7.  If [should `request`{.variable} be blocked due to a bad
    port](#block-bad-port){#ref-for-block-bad-port link-type="dfn"},
    [should fetching `request`{.variable} be blocked as mixed
    content](https://w3c.github.io/webappsec-mixed-content/#should-block-fetch){#ref-for-should-block-fetch
    link-type="dfn"}, or [should `request`{.variable} be blocked by
    Content Security
    Policy](https://w3c.github.io/webappsec-csp/#should-block-request){#ref-for-should-block-request
    link-type="dfn"} returns **blocked**, then set `response`{.variable}
    to a [network
    error](#concept-network-error){#ref-for-concept-network-error①⓪
    link-type="dfn"}.

8.  If `request`{.variable}'s [referrer
    policy](#concept-request-referrer-policy){#ref-for-concept-request-referrer-policy②
    link-type="dfn"} is the empty string, then set
    `request`{.variable}'s [referrer
    policy](#concept-request-referrer-policy){#ref-for-concept-request-referrer-policy③
    link-type="dfn"} to `request`{.variable}'s [policy
    container](#concept-request-policy-container){#ref-for-concept-request-policy-container④
    link-type="dfn"}'s [referrer
    policy](https://html.spec.whatwg.org/multipage/browsers.html#policy-container-referrer-policy){#ref-for-policy-container-referrer-policy
    link-type="dfn"}.

9.  If `request`{.variable}'s
    [referrer](#concept-request-referrer){#ref-for-concept-request-referrer①
    link-type="dfn"} is not \"`no-referrer`\", then set
    `request`{.variable}'s
    [referrer](#concept-request-referrer){#ref-for-concept-request-referrer②
    link-type="dfn"} to the result of invoking [determine
    `request`{.variable}'s
    referrer](https://w3c.github.io/webappsec-referrer-policy/#determine-requests-referrer){#ref-for-determine-requests-referrer
    link-type="dfn"}.
    [\[REFERRER\]](#biblio-referrer "Referrer Policy"){link-type="biblio"}

    As stated in Referrer Policy, user agents can provide the end user
    with options to override `request`{.variable}'s
    [referrer](#concept-request-referrer){#ref-for-concept-request-referrer③
    link-type="dfn"} to \"`no-referrer`\" or have it expose less
    sensitive information.

10. Set `request`{.variable}'s [current
    URL](#concept-request-current-url){#ref-for-concept-request-current-url⑨
    link-type="dfn"}'s
    [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#ref-for-concept-url-scheme⑤
    link-type="dfn"} to \"`https`\" if all of the following conditions
    are true:

    - `request`{.variable}'s [current
      URL](#concept-request-current-url){#ref-for-concept-request-current-url①⓪
      link-type="dfn"}'s
      [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#ref-for-concept-url-scheme⑥
      link-type="dfn"} is \"`http`\"
    - `request`{.variable}'s [current
      URL](#concept-request-current-url){#ref-for-concept-request-current-url①①
      link-type="dfn"}'s
      [host](https://url.spec.whatwg.org/#concept-url-host){#ref-for-concept-url-host
      link-type="dfn"} is a
      [domain](https://url.spec.whatwg.org/#concept-domain){#ref-for-concept-domain
      link-type="dfn"}
    - `request`{.variable}'s [current
      URL](#concept-request-current-url){#ref-for-concept-request-current-url①②
      link-type="dfn"}'s
      [host](https://url.spec.whatwg.org/#concept-url-host){#ref-for-concept-url-host①
      link-type="dfn"}'s [public
      suffix](https://url.spec.whatwg.org/#host-public-suffix){#ref-for-host-public-suffix①
      link-type="dfn"} is not \"`localhost`\" or \"`localhost.`\"
    - Matching `request`{.variable}'s [current
      URL](#concept-request-current-url){#ref-for-concept-request-current-url①③
      link-type="dfn"}'s
      [host](https://url.spec.whatwg.org/#concept-url-host){#ref-for-concept-url-host②
      link-type="dfn"} per [Known HSTS Host Domain Name
      Matching](https://www.rfc-editor.org/rfc/rfc6797.html#section-8.2)
      results in either a superdomain match with an asserted
      `includeSubDomains` directive or a congruent match (with or
      without an asserted `includeSubDomains` directive)
      [\[HSTS\]](#biblio-hsts "HTTP Strict Transport Security (HSTS)"){link-type="biblio"};
      or DNS resolution for the request finds a matching HTTPS RR per
      [section
      9.5](https://datatracker.ietf.org/doc/html/draft-ietf-dnsop-svcb-https#section-9.5)
      of
      [\[SVCB\]](#biblio-svcb "Service binding and parameter specification via the DNS (DNS SVCB and HTTPS RRs)"){link-type="biblio"}.
      [\[HSTS\]](#biblio-hsts "HTTP Strict Transport Security (HSTS)"){link-type="biblio"}
      [\[SVCB\]](#biblio-svcb "Service binding and parameter specification via the DNS (DNS SVCB and HTTPS RRs)"){link-type="biblio"}

    As all DNS operations are generally
    [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#ref-for-implementation-defined①④
    link-type="dfn"}, how it is determined that DNS resolution contains
    an HTTPS RR is also
    [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#ref-for-implementation-defined①⑤
    link-type="dfn"}. As DNS operations are not traditionally performed
    until attempting to [obtain a
    connection](#concept-connection-obtain){#ref-for-concept-connection-obtain
    link-type="dfn"}, user agents might need to perform DNS operations
    earlier, consult local DNS caches, or wait until later in the fetch
    algorithm and potentially unwind logic on discovering the need to
    change `request`{.variable}'s [current
    URL](#concept-request-current-url){#ref-for-concept-request-current-url①④
    link-type="dfn"}'s
    [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#ref-for-concept-url-scheme⑦
    link-type="dfn"}.

11. If `recursive`{.variable} is false, then run the remaining steps [in
    parallel](https://html.spec.whatwg.org/multipage/infrastructure.html#in-parallel){#ref-for-in-parallel①
    link-type="dfn"}.

12. If `response`{.variable} is null, then set `response`{.variable} to
    the result of running the steps corresponding to the first matching
    statement:

    `fetchParams`{.variable}'s [preloaded response candidate](#fetch-params-preloaded-response-candidate){#ref-for-fetch-params-preloaded-response-candidate③ link-type="dfn"} is non-null

    :   1.  Wait until `fetchParams`{.variable}'s [preloaded response
            candidate](#fetch-params-preloaded-response-candidate){#ref-for-fetch-params-preloaded-response-candidate④
            link-type="dfn"} is not \"`pending`\".

        2.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert②⓪
            link-type="dfn"}: `fetchParams`{.variable}'s [preloaded
            response
            candidate](#fetch-params-preloaded-response-candidate){#ref-for-fetch-params-preloaded-response-candidate⑤
            link-type="dfn"} is a
            [response](#concept-response){#ref-for-concept-response④⑨
            link-type="dfn"}.

        3.  Return `fetchParams`{.variable}'s [preloaded response
            candidate](#fetch-params-preloaded-response-candidate){#ref-for-fetch-params-preloaded-response-candidate⑥
            link-type="dfn"}.

    `request`{.variable}'s [current URL](#concept-request-current-url){#ref-for-concept-request-current-url①⑤ link-type="dfn"}'s [origin](https://url.spec.whatwg.org/#concept-url-origin){#ref-for-concept-url-origin①⑤ link-type="dfn"} is [same origin](https://html.spec.whatwg.org/multipage/browsers.html#same-origin){#ref-for-same-origin⑥ link-type="dfn"} with `request`{.variable}'s [origin](#concept-request-origin){#ref-for-concept-request-origin①④ link-type="dfn"}, and `request`{.variable}'s [response tainting](#concept-request-response-tainting){#ref-for-concept-request-response-tainting④ link-type="dfn"} is \"`basic`\"\
    `request`{.variable}'s [current URL](#concept-request-current-url){#ref-for-concept-request-current-url①⑥ link-type="dfn"}'s [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#ref-for-concept-url-scheme⑧ link-type="dfn"} is \"`data`\"\
    `request`{.variable}'s [mode](#concept-request-mode){#ref-for-concept-request-mode①① link-type="dfn"} is \"`navigate`\" or \"`websocket`\"

    :   1.  Set `request`{.variable}'s [response
            tainting](#concept-request-response-tainting){#ref-for-concept-request-response-tainting⑤
            link-type="dfn"} to \"`basic`\".

        2.  Return the result of running [scheme
            fetch](#concept-scheme-fetch){#ref-for-concept-scheme-fetch
            link-type="dfn"} given `fetchParams`{.variable}.

        HTML assigns any documents and workers created from
        [URLs](https://url.spec.whatwg.org/#concept-url){#ref-for-concept-url①⑤
        link-type="dfn"} whose
        [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#ref-for-concept-url-scheme⑨
        link-type="dfn"} is \"`data`\" a unique [opaque
        origin](https://html.spec.whatwg.org/multipage/browsers.html#concept-origin-opaque){#ref-for-concept-origin-opaque
        link-type="dfn"}. Service workers can only be created from
        [URLs](https://url.spec.whatwg.org/#concept-url){#ref-for-concept-url①⑥
        link-type="dfn"} whose
        [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#ref-for-concept-url-scheme①⓪
        link-type="dfn"} is an [HTTP(S)
        scheme](#http-scheme){#ref-for-http-scheme⑤ link-type="dfn"}.
        [\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}
        [\[SW\]](#biblio-sw "Service Workers"){link-type="biblio"}

    `request`{.variable}'s [mode](#concept-request-mode){#ref-for-concept-request-mode①② link-type="dfn"} is \"`same-origin`\"

    :   Return a [network
        error](#concept-network-error){#ref-for-concept-network-error①①
        link-type="dfn"}.

    `request`{.variable}'s [mode](#concept-request-mode){#ref-for-concept-request-mode①③ link-type="dfn"} is \"`no-cors`\"

    :   1.  If `request`{.variable}'s [redirect
            mode](#concept-request-redirect-mode){#ref-for-concept-request-redirect-mode①
            link-type="dfn"} is not \"`follow`\", then return a [network
            error](#concept-network-error){#ref-for-concept-network-error①②
            link-type="dfn"}.

        2.  Set `request`{.variable}'s [response
            tainting](#concept-request-response-tainting){#ref-for-concept-request-response-tainting⑥
            link-type="dfn"} to \"`opaque`\".

        3.  Return the result of running [scheme
            fetch](#concept-scheme-fetch){#ref-for-concept-scheme-fetch①
            link-type="dfn"} given `fetchParams`{.variable}.

    `request`{.variable}'s [current URL](#concept-request-current-url){#ref-for-concept-request-current-url①⑦ link-type="dfn"}'s [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#ref-for-concept-url-scheme①① link-type="dfn"} is not an [HTTP(S) scheme](#http-scheme){#ref-for-http-scheme⑥ link-type="dfn"}

    :   Return a [network
        error](#concept-network-error){#ref-for-concept-network-error①③
        link-type="dfn"}.

    `request`{.variable}'s [use-CORS-preflight flag](#use-cors-preflight-flag){#ref-for-use-cors-preflight-flag② link-type="dfn"} is set\
    `request`{.variable}'s [unsafe-request flag](#unsafe-request-flag){#ref-for-unsafe-request-flag② link-type="dfn"} is set and either `request`{.variable}'s [method](#concept-request-method){#ref-for-concept-request-method⑤ link-type="dfn"} is not a [CORS-safelisted method](#cors-safelisted-method){#ref-for-cors-safelisted-method① link-type="dfn"} or [CORS-unsafe request-header names](#cors-unsafe-request-header-names){#ref-for-cors-unsafe-request-header-names link-type="dfn"} with `request`{.variable}'s [header list](#concept-request-header-list){#ref-for-concept-request-header-list①⓪ link-type="dfn"} [is not empty](https://infra.spec.whatwg.org/#list-is-empty){#ref-for-list-is-empty③ link-type="dfn"}

    :   1.  Set `request`{.variable}'s [response
            tainting](#concept-request-response-tainting){#ref-for-concept-request-response-tainting⑦
            link-type="dfn"} to \"`cors`\".

        2.  Let `corsWithPreflightResponse`{.variable} be the result of
            running [HTTP
            fetch](#concept-http-fetch){#ref-for-concept-http-fetch③
            link-type="dfn"} given `fetchParams`{.variable} and true.

        3.  If `corsWithPreflightResponse`{.variable} is a [network
            error](#concept-network-error){#ref-for-concept-network-error①④
            link-type="dfn"}, then [clear cache
            entries](#concept-cache-clear){#ref-for-concept-cache-clear
            link-type="dfn"} using `request`{.variable}.

        4.  Return `corsWithPreflightResponse`{.variable}.

    Otherwise

    :   1.  Set `request`{.variable}'s [response
            tainting](#concept-request-response-tainting){#ref-for-concept-request-response-tainting⑧
            link-type="dfn"} to \"`cors`\".

        2.  Return the result of running [HTTP
            fetch](#concept-http-fetch){#ref-for-concept-http-fetch④
            link-type="dfn"} given `fetchParams`{.variable}.

13. If `recursive`{.variable} is true, then return
    `response`{.variable}.

14. If `response`{.variable} is not a [network
    error](#concept-network-error){#ref-for-concept-network-error①⑤
    link-type="dfn"} and `response`{.variable} is not a [filtered
    response](#concept-filtered-response){#ref-for-concept-filtered-response①①
    link-type="dfn"}, then:

    1.  If `request`{.variable}'s [response
        tainting](#concept-request-response-tainting){#ref-for-concept-request-response-tainting⑨
        link-type="dfn"} is \"`cors`\", then:

        1.  Let `headerNames`{.variable} be the result of [extracting
            header list
            values](#extract-header-list-values){#ref-for-extract-header-list-values①
            link-type="dfn"} given
            \`[`Access-Control-Expose-Headers`](#http-access-control-expose-headers){#ref-for-http-access-control-expose-headers⑤
            link-type="http-header"}\` and `response`{.variable}'s
            [header
            list](#concept-response-header-list){#ref-for-concept-response-header-list①②
            link-type="dfn"}.

        2.  If `request`{.variable}'s [credentials
            mode](#concept-request-credentials-mode){#ref-for-concept-request-credentials-mode⑨
            link-type="dfn"} is not \"`include`\" and
            `headerNames`{.variable} contains \``*`\`, then set
            `response`{.variable}'s [CORS-exposed header-name
            list](#concept-response-cors-exposed-header-name-list){#ref-for-concept-response-cors-exposed-header-name-list②
            link-type="dfn"} to all unique
            [header](#concept-header){#ref-for-concept-header④⑧
            link-type="dfn"}
            [names](#concept-header-name){#ref-for-concept-header-name①⑧
            link-type="dfn"} in `response`{.variable}'s [header
            list](#concept-response-header-list){#ref-for-concept-response-header-list①③
            link-type="dfn"}.

        3.  Otherwise, if `headerNames`{.variable} is non-null or
            failure, then set `response`{.variable}'s [CORS-exposed
            header-name
            list](#concept-response-cors-exposed-header-name-list){#ref-for-concept-response-cors-exposed-header-name-list③
            link-type="dfn"} to `headerNames`{.variable}.

            One of the `headerNames`{.variable} can still be \``*`\` at
            this point, but will only match a
            [header](#concept-header){#ref-for-concept-header④⑨
            link-type="dfn"} whose
            [name](#concept-header-name){#ref-for-concept-header-name①⑨
            link-type="dfn"} is \``*`\`.

    2.  Set `response`{.variable} to the following [filtered
        response](#concept-filtered-response){#ref-for-concept-filtered-response①②
        link-type="dfn"} with `response`{.variable} as its [internal
        response](#concept-internal-response){#ref-for-concept-internal-response①⓪
        link-type="dfn"}, depending on `request`{.variable}'s [response
        tainting](#concept-request-response-tainting){#ref-for-concept-request-response-tainting①⓪
        link-type="dfn"}:

        \"`basic`\"
        :   [basic filtered
            response](#concept-filtered-response-basic){#ref-for-concept-filtered-response-basic
            link-type="dfn"}

        \"`cors`\"
        :   [CORS filtered
            response](#concept-filtered-response-cors){#ref-for-concept-filtered-response-cors①
            link-type="dfn"}

        \"`opaque`\"
        :   [opaque filtered
            response](#concept-filtered-response-opaque){#ref-for-concept-filtered-response-opaque⑤
            link-type="dfn"}

15. Let `internalResponse`{.variable} be `response`{.variable}, if
    `response`{.variable} is a [network
    error](#concept-network-error){#ref-for-concept-network-error①⑥
    link-type="dfn"}; otherwise `response`{.variable}'s [internal
    response](#concept-internal-response){#ref-for-concept-internal-response①①
    link-type="dfn"}.

16. If `internalResponse`{.variable}'s [URL
    list](#concept-response-url-list){#ref-for-concept-response-url-list⑦
    link-type="dfn"} [is
    empty](https://infra.spec.whatwg.org/#list-is-empty){#ref-for-list-is-empty④
    link-type="dfn"}, then set it to a
    [clone](https://infra.spec.whatwg.org/#list-clone){#ref-for-list-clone
    link-type="dfn"} of `request`{.variable}'s [URL
    list](#concept-request-url-list){#ref-for-concept-request-url-list④
    link-type="dfn"}.

    A [response](#concept-response){#ref-for-concept-response⑤⓪
    link-type="dfn"}'s [URL
    list](#concept-response-url-list){#ref-for-concept-response-url-list⑧
    link-type="dfn"} can be empty, e.g., when fetching an `about:` URL.

17. If `request`{.variable} has a [redirect-tainted
    origin](#concept-request-tainted-origin){#ref-for-concept-request-tainted-origin②
    link-type="dfn"}, then set `internalResponse`{.variable}'s
    [has-cross-origin-redirects](#response-has-cross-origin-redirects){#ref-for-response-has-cross-origin-redirects
    link-type="dfn"} to true.

18. If `request`{.variable}'s [timing allow failed
    flag](#timing-allow-failed){#ref-for-timing-allow-failed②
    link-type="dfn"} is unset, then set `internalResponse`{.variable}'s
    [timing allow passed
    flag](#concept-response-timing-allow-passed){#ref-for-concept-response-timing-allow-passed
    link-type="dfn"}.

19. If `response`{.variable} is not a [network
    error](#concept-network-error){#ref-for-concept-network-error①⑦
    link-type="dfn"} and any of the following returns **blocked**

    - [should `internalResponse`{.variable} to `request`{.variable} be
      blocked as mixed
      content](https://w3c.github.io/webappsec-mixed-content/#should-block-response){#ref-for-should-block-response
      link-type="dfn"}

    - [should `internalResponse`{.variable} to `request`{.variable} be
      blocked by Content Security
      Policy](https://w3c.github.io/webappsec-csp/#should-block-response){#ref-for-should-block-response①
      link-type="dfn"}

    - [should `internalResponse`{.variable} to `request`{.variable} be
      blocked due to its MIME
      type](#should-response-to-request-be-blocked-due-to-mime-type?){#ref-for-should-response-to-request-be-blocked-due-to-mime-type?①
      link-type="dfn"}

    - [should `internalResponse`{.variable} to `request`{.variable} be
      blocked due to
      nosniff](#should-response-to-request-be-blocked-due-to-nosniff?){#ref-for-should-response-to-request-be-blocked-due-to-nosniff?①
      link-type="dfn"}

    then set `response`{.variable} and `internalResponse`{.variable} to
    a [network
    error](#concept-network-error){#ref-for-concept-network-error①⑧
    link-type="dfn"}.

20. If `response`{.variable}'s
    [type](#concept-response-type){#ref-for-concept-response-type⑥
    link-type="dfn"} is \"`opaque`\", `internalResponse`{.variable}'s
    [status](#concept-response-status){#ref-for-concept-response-status⑤
    link-type="dfn"} is 206, `internalResponse`{.variable}'s
    [range-requested
    flag](#concept-response-range-requested-flag){#ref-for-concept-response-range-requested-flag
    link-type="dfn"} is set, and `request`{.variable}'s [header
    list](#concept-request-header-list){#ref-for-concept-request-header-list①①
    link-type="dfn"} [does not
    contain](#header-list-contains){#ref-for-header-list-contains①⓪
    link-type="dfn"} \``Range`\`, then set `response`{.variable} and
    `internalResponse`{.variable} to a [network
    error](#concept-network-error){#ref-for-concept-network-error①⑨
    link-type="dfn"}.

    ::: {.note role="note"}
    Traditionally, APIs accept a ranged response even if a range was not
    requested. This prevents a partial response from an earlier ranged
    request being provided to an API that did not make a range request.

    Further details
    The above steps prevent the following attack:

    A media element is used to request a range of a cross-origin HTML
    resource. Although this is invalid media, a reference to a clone of
    the response can be retained in a service worker. This can later be
    used as the response to a script element's fetch. If the partial
    response is valid JavaScript (even though the whole resource is
    not), executing it would leak private data.
    :::

21. If `response`{.variable} is not a [network
    error](#concept-network-error){#ref-for-concept-network-error②⓪
    link-type="dfn"} and either `request`{.variable}'s
    [method](#concept-request-method){#ref-for-concept-request-method⑥
    link-type="dfn"} is \``HEAD`\` or \``CONNECT`\`, or
    `internalResponse`{.variable}'s
    [status](#concept-response-status){#ref-for-concept-response-status⑥
    link-type="dfn"} is a [null body
    status](#null-body-status){#ref-for-null-body-status
    link-type="dfn"}, set `internalResponse`{.variable}'s
    [body](#concept-response-body){#ref-for-concept-response-body⑨
    link-type="dfn"} to null and disregard any enqueuing toward it (if
    any).

    This standardizes the error handling for servers that violate HTTP.

22. If `request`{.variable}'s [integrity
    metadata](#concept-request-integrity-metadata){#ref-for-concept-request-integrity-metadata①
    link-type="dfn"} is not the empty string, then:

    1.  Let `processBodyError`{.variable} be this step: run [fetch
        response handover](#fetch-finale){#ref-for-fetch-finale
        link-type="dfn"} given `fetchParams`{.variable} and a [network
        error](#concept-network-error){#ref-for-concept-network-error②①
        link-type="dfn"}.

    2.  If `response`{.variable}'s
        [body](#concept-response-body){#ref-for-concept-response-body①⓪
        link-type="dfn"} is null, then run `processBodyError`{.variable}
        and abort these steps.

    3.  Let `processBody`{.variable} given `bytes`{.variable} be these
        steps:

        1.  If `bytes`{.variable} do not
            [match](https://w3c.github.io/webappsec-subresource-integrity/#does-response-match-metadatalist){#ref-for-does-response-match-metadatalist
            link-type="dfn"} `request`{.variable}'s [integrity
            metadata](#concept-request-integrity-metadata){#ref-for-concept-request-integrity-metadata②
            link-type="dfn"}, then run `processBodyError`{.variable} and
            abort these steps.
            [\[SRI\]](#biblio-sri "Subresource Integrity"){link-type="biblio"}

        2.  Set `response`{.variable}'s
            [body](#concept-response-body){#ref-for-concept-response-body①①
            link-type="dfn"} to `bytes`{.variable} [as a
            body](#byte-sequence-as-a-body){#ref-for-byte-sequence-as-a-body①
            link-type="dfn"}.

        3.  Run [fetch response
            handover](#fetch-finale){#ref-for-fetch-finale①
            link-type="dfn"} given `fetchParams`{.variable} and
            `response`{.variable}.

    4.  [Fully read](#body-fully-read){#ref-for-body-fully-read
        link-type="dfn"} `response`{.variable}'s
        [body](#concept-response-body){#ref-for-concept-response-body①②
        link-type="dfn"} given `processBody`{.variable} and
        `processBodyError`{.variable}.

23. Otherwise, run [fetch response
    handover](#fetch-finale){#ref-for-fetch-finale② link-type="dfn"}
    given `fetchParams`{.variable} and `response`{.variable}.
:::

------------------------------------------------------------------------

::: {.algorithm algorithm="fetch response handover"}
The [fetch response handover]{#fetch-finale .dfn .dfn-paneled
dfn-type="dfn" noexport=""}, given a [fetch
params](#fetch-params){#ref-for-fetch-params⑤ link-type="dfn"}
`fetchParams`{.variable} and a
[response](#concept-response){#ref-for-concept-response⑤①
link-type="dfn"} `response`{.variable}, run these steps:

1.  Let `timingInfo`{.variable} be `fetchParams`{.variable}'s [timing
    info](#fetch-params-timing-info){#ref-for-fetch-params-timing-info①
    link-type="dfn"}.

2.  If `response`{.variable} is not a [network
    error](#concept-network-error){#ref-for-concept-network-error②②
    link-type="dfn"} and `fetchParams`{.variable}'s
    [request](#fetch-params-request){#ref-for-fetch-params-request②
    link-type="dfn"}'s
    [client](#concept-request-client){#ref-for-concept-request-client②⓪
    link-type="dfn"} is a [secure
    context](https://html.spec.whatwg.org/multipage/webappapis.html#secure-context){#ref-for-secure-context
    link-type="dfn"}, then set `timingInfo`{.variable}'s [server-timing
    headers](#fetch-timing-info-server-timing-headers){#ref-for-fetch-timing-info-server-timing-headers
    link-type="dfn"} to the result of [getting, decoding, and
    splitting](#concept-header-list-get-decode-split){#ref-for-concept-header-list-get-decode-split⑤
    link-type="dfn"} \``Server-Timing`\` from `response`{.variable}'s
    [internal
    response](#concept-internal-response){#ref-for-concept-internal-response①②
    link-type="dfn"}'s [header
    list](#concept-response-header-list){#ref-for-concept-response-header-list①④
    link-type="dfn"}.

    Using \_response\_'s [internal
    response](#concept-internal-response){#ref-for-concept-internal-response①③
    link-type="dfn"} is safe as exposing \``Server-Timing`\` header data
    is guarded through the \``Timing-Allow-Origin`\` header.

    The user agent may decide to expose \``Server-Timing`\` headers to
    non-secure contexts requests as well.

3.  Let `processResponseEndOfBody`{.variable} be the following steps:

    1.  Let `unsafeEndTime`{.variable} be the [unsafe shared current
        time](https://w3c.github.io/hr-time/#dfn-unsafe-shared-current-time){#ref-for-dfn-unsafe-shared-current-time⑤
        link-type="dfn"}.

    2.  If `fetchParams`{.variable}'s
        [request](#fetch-params-request){#ref-for-fetch-params-request③
        link-type="dfn"}'s
        [destination](#concept-request-destination){#ref-for-concept-request-destination①⑤
        link-type="dfn"} is \"`document`\", then set
        `fetchParams`{.variable}'s
        [controller](#fetch-params-controller){#ref-for-fetch-params-controller④
        link-type="dfn"}'s [full timing
        info](#fetch-controller-full-timing-info){#ref-for-fetch-controller-full-timing-info②
        link-type="dfn"} to `fetchParams`{.variable}'s [timing
        info](#fetch-params-timing-info){#ref-for-fetch-params-timing-info②
        link-type="dfn"}.

    3.  Set `fetchParams`{.variable}'s
        [controller](#fetch-params-controller){#ref-for-fetch-params-controller⑤
        link-type="dfn"}'s [report timing
        steps](#fetch-controller-report-timing-steps){#ref-for-fetch-controller-report-timing-steps②
        link-type="dfn"} to the following steps given a [global
        object](https://html.spec.whatwg.org/multipage/webappapis.html#global-object){#ref-for-global-object⑦
        link-type="dfn"} `global`{.variable}:

        1.  If `fetchParams`{.variable}'s
            [request](#fetch-params-request){#ref-for-fetch-params-request④
            link-type="dfn"}'s
            [URL](#concept-request-url){#ref-for-concept-request-url④
            link-type="dfn"}'s
            [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#ref-for-concept-url-scheme①②
            link-type="dfn"} is not an [HTTP(S)
            scheme](#http-scheme){#ref-for-http-scheme⑦
            link-type="dfn"}, then return.

        2.  Set `timingInfo`{.variable}'s [end
            time](#fetch-timing-info-end-time){#ref-for-fetch-timing-info-end-time
            link-type="dfn"} to the [relative high resolution
            time](https://w3c.github.io/hr-time/#dfn-relative-high-resolution-time){#ref-for-dfn-relative-high-resolution-time
            link-type="dfn"} given `unsafeEndTime`{.variable} and
            `global`{.variable}.

        3.  Let `cacheState`{.variable} be `response`{.variable}'s
            [cache
            state](#concept-response-cache-state){#ref-for-concept-response-cache-state
            link-type="dfn"}.

        4.  Let `bodyInfo`{.variable} be `response`{.variable}'s [body
            info](#concept-response-body-info){#ref-for-concept-response-body-info③
            link-type="dfn"}.

        5.  If `response`{.variable}'s [timing allow passed
            flag](#concept-response-timing-allow-passed){#ref-for-concept-response-timing-allow-passed①
            link-type="dfn"} is not set, then set
            `timingInfo`{.variable} to the result of [creating an opaque
            timing
            info](#create-an-opaque-timing-info){#ref-for-create-an-opaque-timing-info
            link-type="dfn"} for `timingInfo`{.variable} and set
            `cacheState`{.variable} to the empty string.

            This covers the case of `response`{.variable} being a
            [network
            error](#concept-network-error){#ref-for-concept-network-error②③
            link-type="dfn"}.

        6.  Let `responseStatus`{.variable} be 0.

        7.  If `fetchParams`{.variable}'s
            [request](#fetch-params-request){#ref-for-fetch-params-request⑤
            link-type="dfn"}'s
            [mode](#concept-request-mode){#ref-for-concept-request-mode①④
            link-type="dfn"} is not \"`navigate`\" or
            `response`{.variable}'s
            [has-cross-origin-redirects](#response-has-cross-origin-redirects){#ref-for-response-has-cross-origin-redirects①
            link-type="dfn"} is false:

            1.  Set `responseStatus`{.variable} to
                `response`{.variable}'s
                [status](#concept-response-status){#ref-for-concept-response-status⑦
                link-type="dfn"}.

            2.  Let `mimeType`{.variable} be the result of [extracting a
                MIME
                type](#concept-header-extract-mime-type){#ref-for-concept-header-extract-mime-type⑦
                link-type="dfn"} from `response`{.variable}'s [header
                list](#concept-response-header-list){#ref-for-concept-response-header-list①⑤
                link-type="dfn"}.

            3.  If `mimeType`{.variable} is not failure, then set
                `bodyInfo`{.variable}'s [content
                type](#response-body-info-content-type){#ref-for-response-body-info-content-type
                link-type="dfn"} to the result of [minimizing a
                supported MIME
                type](https://mimesniff.spec.whatwg.org/#minimize-a-supported-mime-type){#ref-for-minimize-a-supported-mime-type
                link-type="dfn"} given `mimeType`{.variable}.

        8.  If `fetchParams`{.variable}'s
            [request](#fetch-params-request){#ref-for-fetch-params-request⑥
            link-type="dfn"}'s [initiator
            type](#request-initiator-type){#ref-for-request-initiator-type
            link-type="dfn"} is non-null, then [mark resource
            timing](https://w3c.github.io/resource-timing/#dfn-mark-resource-timing){#ref-for-dfn-mark-resource-timing
            link-type="dfn"} given `timingInfo`{.variable},
            `fetchParams`{.variable}'s
            [request](#fetch-params-request){#ref-for-fetch-params-request⑦
            link-type="dfn"}'s
            [URL](#concept-request-url){#ref-for-concept-request-url⑤
            link-type="dfn"}, `fetchParams`{.variable}'s
            [request](#fetch-params-request){#ref-for-fetch-params-request⑧
            link-type="dfn"}'s [initiator
            type](#request-initiator-type){#ref-for-request-initiator-type①
            link-type="dfn"}, `global`{.variable},
            `cacheState`{.variable}, `bodyInfo`{.variable}, and
            `responseStatus`{.variable}.

    4.  Let `processResponseEndOfBodyTask`{.variable} be the following
        steps:

        1.  Set `fetchParams`{.variable}'s
            [request](#fetch-params-request){#ref-for-fetch-params-request⑨
            link-type="dfn"}'s [done
            flag](#done-flag){#ref-for-done-flag② link-type="dfn"}.

        2.  If `fetchParams`{.variable}'s [process response
            end-of-body](#fetch-params-process-response-end-of-body){#ref-for-fetch-params-process-response-end-of-body①
            link-type="dfn"} is non-null, then run
            `fetchParams`{.variable}'s [process response
            end-of-body](#fetch-params-process-response-end-of-body){#ref-for-fetch-params-process-response-end-of-body②
            link-type="dfn"} given `response`{.variable}.

        3.  If `fetchParams`{.variable}'s
            [request](#fetch-params-request){#ref-for-fetch-params-request①⓪
            link-type="dfn"}'s [initiator
            type](#request-initiator-type){#ref-for-request-initiator-type②
            link-type="dfn"} is non-null and `fetchParams`{.variable}'s
            [request](#fetch-params-request){#ref-for-fetch-params-request①①
            link-type="dfn"}'s
            [client](#concept-request-client){#ref-for-concept-request-client②①
            link-type="dfn"}'s [global
            object](https://html.spec.whatwg.org/multipage/webappapis.html#concept-settings-object-global){#ref-for-concept-settings-object-global⑤
            link-type="dfn"} is `fetchParams`{.variable}'s [task
            destination](#fetch-params-task-destination){#ref-for-fetch-params-task-destination①
            link-type="dfn"}, then run `fetchParams`{.variable}'s
            [controller](#fetch-params-controller){#ref-for-fetch-params-controller⑥
            link-type="dfn"}'s [report timing
            steps](#fetch-controller-report-timing-steps){#ref-for-fetch-controller-report-timing-steps③
            link-type="dfn"} given `fetchParams`{.variable}'s
            [request](#fetch-params-request){#ref-for-fetch-params-request①②
            link-type="dfn"}'s
            [client](#concept-request-client){#ref-for-concept-request-client②②
            link-type="dfn"}'s [global
            object](https://html.spec.whatwg.org/multipage/webappapis.html#concept-settings-object-global){#ref-for-concept-settings-object-global⑥
            link-type="dfn"}.

    5.  [Queue a fetch
        task](#queue-a-fetch-task){#ref-for-queue-a-fetch-task⑤
        link-type="dfn"} to run
        `processResponseEndOfBodyTask`{.variable} with
        `fetchParams`{.variable}'s [task
        destination](#fetch-params-task-destination){#ref-for-fetch-params-task-destination②
        link-type="dfn"}.

4.  If `fetchParams`{.variable}'s [process
    response](#fetch-params-process-response){#ref-for-fetch-params-process-response①
    link-type="dfn"} is non-null, then [queue a fetch
    task](#queue-a-fetch-task){#ref-for-queue-a-fetch-task⑥
    link-type="dfn"} to run `fetchParams`{.variable}'s [process
    response](#fetch-params-process-response){#ref-for-fetch-params-process-response②
    link-type="dfn"} given `response`{.variable}, with
    `fetchParams`{.variable}'s [task
    destination](#fetch-params-task-destination){#ref-for-fetch-params-task-destination③
    link-type="dfn"}.

5.  Let `internalResponse`{.variable} be `response`{.variable}, if
    `response`{.variable} is a [network
    error](#concept-network-error){#ref-for-concept-network-error②④
    link-type="dfn"}; otherwise `response`{.variable}'s [internal
    response](#concept-internal-response){#ref-for-concept-internal-response①④
    link-type="dfn"}.

6.  If `internalResponse`{.variable}'s
    [body](#concept-response-body){#ref-for-concept-response-body①③
    link-type="dfn"} is null, then run
    `processResponseEndOfBody`{.variable}.

7.  Otherwise:

    1.  Let `transformStream`{.variable} be a new
        [`TransformStream`{.idl}](https://streams.spec.whatwg.org/#transformstream){#ref-for-transformstream
        link-type="idl"}.

    2.  Let `identityTransformAlgorithm`{.variable} be an algorithm
        which, given `chunk`{.variable},
        [enqueues](https://streams.spec.whatwg.org/#transformstream-enqueue){#ref-for-transformstream-enqueue
        link-type="dfn"} `chunk`{.variable} in
        `transformStream`{.variable}.

    3.  [Set
        up](https://streams.spec.whatwg.org/#transformstream-set-up){#ref-for-transformstream-set-up
        link-type="dfn"} `transformStream`{.variable} with
        [*transformAlgorithm*](https://streams.spec.whatwg.org/#transformstream-set-up-transformalgorithm){#ref-for-transformstream-set-up-transformalgorithm
        link-type="dfn"} set to `identityTransformAlgorithm`{.variable}
        and
        [*flushAlgorithm*](https://streams.spec.whatwg.org/#transformstream-set-up-flushalgorithm){#ref-for-transformstream-set-up-flushalgorithm
        link-type="dfn"} set to `processResponseEndOfBody`{.variable}.

    4.  Set `internalResponse`{.variable}'s
        [body](#concept-response-body){#ref-for-concept-response-body①④
        link-type="dfn"}'s
        [stream](#concept-body-stream){#ref-for-concept-body-stream⑤
        link-type="dfn"} to the result of
        `internalResponse`{.variable}'s
        [body](#concept-response-body){#ref-for-concept-response-body①⑤
        link-type="dfn"}'s
        [stream](#concept-body-stream){#ref-for-concept-body-stream⑥
        link-type="dfn"} [piped
        through](https://streams.spec.whatwg.org/#readablestream-pipe-through){#ref-for-readablestream-pipe-through
        link-type="dfn"} `transformStream`{.variable}.

    This
    [`TransformStream`{.idl}](https://streams.spec.whatwg.org/#transformstream){#ref-for-transformstream①
    link-type="idl"} is needed for the purpose of receiving a
    notification when the stream reaches its end, and is otherwise an
    [identity transform
    stream](https://streams.spec.whatwg.org/#identity-transform-stream){#ref-for-identity-transform-stream
    link-type="dfn"}.

8.  If `fetchParams`{.variable}'s [process response consume
    body](#fetch-params-process-response-consume-body){#ref-for-fetch-params-process-response-consume-body①
    link-type="dfn"} is non-null, then:

    1.  Let `processBody`{.variable} given `nullOrBytes`{.variable} be
        this step: run `fetchParams`{.variable}'s [process response
        consume
        body](#fetch-params-process-response-consume-body){#ref-for-fetch-params-process-response-consume-body②
        link-type="dfn"} given `response`{.variable} and
        `nullOrBytes`{.variable}.

    2.  Let `processBodyError`{.variable} be this step: run
        `fetchParams`{.variable}'s [process response consume
        body](#fetch-params-process-response-consume-body){#ref-for-fetch-params-process-response-consume-body③
        link-type="dfn"} given `response`{.variable} and failure.

    3.  If `internalResponse`{.variable}'s
        [body](#concept-response-body){#ref-for-concept-response-body①⑥
        link-type="dfn"} is null, then [queue a fetch
        task](#queue-a-fetch-task){#ref-for-queue-a-fetch-task⑦
        link-type="dfn"} to run `processBody`{.variable} given null,
        with `fetchParams`{.variable}'s [task
        destination](#fetch-params-task-destination){#ref-for-fetch-params-task-destination④
        link-type="dfn"}.

    4.  Otherwise, [fully
        read](#body-fully-read){#ref-for-body-fully-read①
        link-type="dfn"} `internalResponse`{.variable}'s
        [body](#concept-response-body){#ref-for-concept-response-body①⑦
        link-type="dfn"} given `processBody`{.variable},
        `processBodyError`{.variable}, and `fetchParams`{.variable}'s
        [task
        destination](#fetch-params-task-destination){#ref-for-fetch-params-task-destination⑤
        link-type="dfn"}.
:::

### [4.2. ]{.secno}[]{#basic-fetch .bs-old-id}[Scheme fetch]{.content}[](#scheme-fetch){.self-link} {#scheme-fetch .heading .settled level="4.2"}

::: {.algorithm algorithm="scheme fetch"}
To [[]{#concept-basic-fetch .bs-old-id}scheme
fetch]{#concept-scheme-fetch .dfn .dfn-paneled dfn-type="dfn"
noexport=""}, given a [fetch
params](#fetch-params){#ref-for-fetch-params⑥ link-type="dfn"}
`fetchParams`{.variable}:

1.  If `fetchParams`{.variable} is
    [canceled](#fetch-params-canceled){#ref-for-fetch-params-canceled①
    link-type="dfn"}, then return the [appropriate network
    error](#appropriate-network-error){#ref-for-appropriate-network-error
    link-type="dfn"} for `fetchParams`{.variable}.

2.  Let `request`{.variable} be `fetchParams`{.variable}'s
    [request](#fetch-params-request){#ref-for-fetch-params-request①③
    link-type="dfn"}.

3.  Switch on `request`{.variable}'s [current
    URL](#concept-request-current-url){#ref-for-concept-request-current-url①⑧
    link-type="dfn"}'s
    [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#ref-for-concept-url-scheme①③
    link-type="dfn"} and run the associated steps:

    \"`about`\"

    :   If `request`{.variable}'s [current
        URL](#concept-request-current-url){#ref-for-concept-request-current-url①⑨
        link-type="dfn"}'s
        [path](https://url.spec.whatwg.org/#concept-url-path){#ref-for-concept-url-path①
        link-type="dfn"} is the string \"`blank`\", then return a new
        [response](#concept-response){#ref-for-concept-response⑤②
        link-type="dfn"} whose [status
        message](#concept-response-status-message){#ref-for-concept-response-status-message③
        link-type="dfn"} is \``OK`\`, [header
        list](#concept-response-header-list){#ref-for-concept-response-header-list①⑥
        link-type="dfn"} is « (\``Content-Type`\`,
        \``text/html;charset=utf-8`\`) », and
        [body](#concept-response-body){#ref-for-concept-response-body①⑧
        link-type="dfn"} is the empty byte sequence [as a
        body](#byte-sequence-as-a-body){#ref-for-byte-sequence-as-a-body②
        link-type="dfn"}.

        [URLs](https://url.spec.whatwg.org/#concept-url){#ref-for-concept-url①⑦
        link-type="dfn"} such as \"`about:config`\" are handled during
        [navigation](https://html.spec.whatwg.org/multipage/browsing-the-web.html#navigate){#ref-for-navigate
        link-type="dfn"} and result in a [network
        error](#concept-network-error){#ref-for-concept-network-error②⑤
        link-type="dfn"} in the context of
        [fetching](#concept-fetch){#ref-for-concept-fetch②④
        link-type="dfn"}.

    \"`blob`\"

    :   1.  Let `blobURLEntry`{.variable} be `request`{.variable}'s
            [current
            URL](#concept-request-current-url){#ref-for-concept-request-current-url②⓪
            link-type="dfn"}'s [blob URL
            entry](https://url.spec.whatwg.org/#concept-url-blob-entry){#ref-for-concept-url-blob-entry
            link-type="dfn"}.

        2.  If `request`{.variable}'s
            [method](#concept-request-method){#ref-for-concept-request-method⑦
            link-type="dfn"} is not \``GET`\` or
            `blobURLEntry`{.variable} is null, then return a [network
            error](#concept-network-error){#ref-for-concept-network-error②⑥
            link-type="dfn"}.
            [\[FILEAPI\]](#biblio-fileapi "File API"){link-type="biblio"}

            The \``GET`\`
            [method](#concept-method){#ref-for-concept-method①⓪
            link-type="dfn"} restriction serves no useful purpose other
            than being interoperable.

        3.  Let `requestEnvironment`{.variable} be the result of
            [determining the
            environment](#request-determine-the-environment){#ref-for-request-determine-the-environment
            link-type="dfn"} given `request`{.variable}.

        4.  Let `isTopLevelNavigation`{.variable} be true if
            `request`{.variable}'s
            [destination](#concept-request-destination){#ref-for-concept-request-destination①⑥
            link-type="dfn"} is \"`document`\"; otherwise, false.

        5.  If `isTopLevelNavigation`{.variable} is false and
            `requestEnvironment`{.variable} is null, then return a
            [network
            error](#concept-network-error){#ref-for-concept-network-error②⑦
            link-type="dfn"}.

        6.  Let `navigationOrEnvironment`{.variable} be the string
            \"`navigation`\" if `isTopLevelNavigation`{.variable} is
            true; otherwise, `requestEnvironment`{.variable}.

        7.  Let `blob`{.variable} be the result of [obtaining a blob
            object](https://w3c.github.io/FileAPI/#blob-url-obtain-object){#ref-for-blob-url-obtain-object
            link-type="dfn"} given `blobURLEntry`{.variable} and
            `navigationOrEnvironment`{.variable}.

        8.  If `blob`{.variable} is not a
            [`Blob`{.idl}](https://w3c.github.io/FileAPI/#dfn-Blob){#ref-for-dfn-Blob①
            link-type="idl"} object, then return a [network
            error](#concept-network-error){#ref-for-concept-network-error②⑧
            link-type="dfn"}.

        9.  Let `response`{.variable} be a new
            [response](#concept-response){#ref-for-concept-response⑤③
            link-type="dfn"}.

        10. Let `fullLength`{.variable} be `blob`{.variable}'s
            [`size`{.idl}](https://w3c.github.io/FileAPI/#dfn-size){#ref-for-dfn-size
            link-type="idl"}.

        11. Let `serializedFullLength`{.variable} be
            `fullLength`{.variable},
            [serialized](#serialize-an-integer){#ref-for-serialize-an-integer⑤
            link-type="dfn"} and [isomorphic
            encoded](https://infra.spec.whatwg.org/#isomorphic-encode){#ref-for-isomorphic-encode⑦
            link-type="dfn"}.

        12. Let `type`{.variable} be `blob`{.variable}'s
            [`type`{.idl}](https://w3c.github.io/FileAPI/#dfn-type){#ref-for-dfn-type
            link-type="idl"}.

        13. If `request`{.variable}'s [header
            list](#concept-request-header-list){#ref-for-concept-request-header-list①②
            link-type="dfn"} [does not
            contain](#header-list-contains){#ref-for-header-list-contains①①
            link-type="dfn"} \``Range`\`:

            1.  Let `bodyWithType`{.variable} be the result of [safely
                extracting](#bodyinit-safely-extract){#ref-for-bodyinit-safely-extract②
                link-type="dfn"} `blob`{.variable}.

            2.  Set `response`{.variable}'s [status
                message](#concept-response-status-message){#ref-for-concept-response-status-message④
                link-type="dfn"} to \``OK`\`.

            3.  Set `response`{.variable}'s
                [body](#concept-response-body){#ref-for-concept-response-body①⑨
                link-type="dfn"} to `bodyWithType`{.variable}'s
                [body](#body-with-type-body){#ref-for-body-with-type-body①
                link-type="dfn"}.

            4.  Set `response`{.variable}'s [header
                list](#concept-response-header-list){#ref-for-concept-response-header-list①⑦
                link-type="dfn"} to « (\``Content-Length`\`,
                `serializedFullLength`{.variable}), (\``Content-Type`\`,
                `type`{.variable}) ».

        14. Otherwise:

            1.  Set `response`{.variable}'s [range-requested
                flag](#concept-response-range-requested-flag){#ref-for-concept-response-range-requested-flag①
                link-type="dfn"}.

            2.  Let `rangeHeader`{.variable} be the result of
                [getting](#concept-header-list-get){#ref-for-concept-header-list-get④
                link-type="dfn"} \``Range`\` from `request`{.variable}'s
                [header
                list](#concept-request-header-list){#ref-for-concept-request-header-list①③
                link-type="dfn"}.

            3.  Let `rangeValue`{.variable} be the result of [parsing a
                single range header
                value](#simple-range-header-value){#ref-for-simple-range-header-value②
                link-type="dfn"} given `rangeHeader`{.variable} and
                true.

            4.  If `rangeValue`{.variable} is failure, then return a
                [network
                error](#concept-network-error){#ref-for-concept-network-error②⑨
                link-type="dfn"}.

            5.  Let (`rangeStart`{.variable}, `rangeEnd`{.variable}) be
                `rangeValue`{.variable}.

            6.  If `rangeStart`{.variable} is null:

                1.  Set `rangeStart`{.variable} to
                    `fullLength`{.variable} − `rangeEnd`{.variable}.

                2.  Set `rangeEnd`{.variable} to
                    `rangeStart`{.variable} + `rangeEnd`{.variable} − 1.

            7.  Otherwise:

                1.  If `rangeStart`{.variable} is greater than or equal
                    to `fullLength`{.variable}, then return a [network
                    error](#concept-network-error){#ref-for-concept-network-error③⓪
                    link-type="dfn"}.

                2.  If `rangeEnd`{.variable} is null or
                    `rangeEnd`{.variable} is greater than or equal to
                    `fullLength`{.variable}, then set
                    `rangeEnd`{.variable} to `fullLength`{.variable} −
                    1.

            8.  Let `slicedBlob`{.variable} be the result of invoking
                [slice
                blob](https://w3c.github.io/FileAPI/#slice-blob){#ref-for-slice-blob
                link-type="dfn"} given `blob`{.variable},
                `rangeStart`{.variable}, `rangeEnd`{.variable} + 1, and
                `type`{.variable}.

                A range header denotes an inclusive byte range, while
                the [slice
                blob](https://w3c.github.io/FileAPI/#slice-blob){#ref-for-slice-blob①
                link-type="dfn"} algorithm input range does not. To use
                the [slice
                blob](https://w3c.github.io/FileAPI/#slice-blob){#ref-for-slice-blob②
                link-type="dfn"} algorithm, we have to increment
                `rangeEnd`{.variable}.

            9.  Let `slicedBodyWithType`{.variable} be the result of
                [safely
                extracting](#bodyinit-safely-extract){#ref-for-bodyinit-safely-extract③
                link-type="dfn"} `slicedBlob`{.variable}.

            10. Set `response`{.variable}'s
                [body](#concept-response-body){#ref-for-concept-response-body②⓪
                link-type="dfn"} to `slicedBodyWithType`{.variable}'s
                [body](#body-with-type-body){#ref-for-body-with-type-body②
                link-type="dfn"}.

            11. Let `serializedSlicedLength`{.variable} be
                `slicedBlob`{.variable}'s
                [`size`{.idl}](https://w3c.github.io/FileAPI/#dfn-size){#ref-for-dfn-size①
                link-type="idl"},
                [serialized](#serialize-an-integer){#ref-for-serialize-an-integer⑥
                link-type="dfn"} and [isomorphic
                encoded](https://infra.spec.whatwg.org/#isomorphic-encode){#ref-for-isomorphic-encode⑧
                link-type="dfn"}.

            12. Let `contentRange`{.variable} be the result of invoking
                [build a content
                range](#build-a-content-range){#ref-for-build-a-content-range
                link-type="dfn"} given `rangeStart`{.variable},
                `rangeEnd`{.variable}, and `fullLength`{.variable}.

            13. Set `response`{.variable}'s
                [status](#concept-response-status){#ref-for-concept-response-status⑧
                link-type="dfn"} to 206.

            14. Set `response`{.variable}'s [status
                message](#concept-response-status-message){#ref-for-concept-response-status-message⑤
                link-type="dfn"} to \``Partial Content`\`.

            15. Set `response`{.variable}'s [header
                list](#concept-response-header-list){#ref-for-concept-response-header-list①⑧
                link-type="dfn"} to « (\``Content-Length`\`,
                `serializedSlicedLength`{.variable}),
                (\``Content-Type`\`, `type`{.variable}),
                (\``Content-Range`\`, `contentRange`{.variable}) ».

        15. Return `response`{.variable}.

    \"`data`\"

    :   1.  Let `dataURLStruct`{.variable} be the result of running the
            [`data:` URL
            processor](#data-url-processor){#ref-for-data-url-processor
            link-type="dfn"} on `request`{.variable}'s [current
            URL](#concept-request-current-url){#ref-for-concept-request-current-url②①
            link-type="dfn"}.

        2.  If `dataURLStruct`{.variable} is failure, then return a
            [network
            error](#concept-network-error){#ref-for-concept-network-error③①
            link-type="dfn"}.

        3.  Let `mimeType`{.variable} be `dataURLStruct`{.variable}'s
            [MIME
            type](#data-url-struct-mime-type){#ref-for-data-url-struct-mime-type
            link-type="dfn"},
            [serialized](https://mimesniff.spec.whatwg.org/#serialize-a-mime-type-to-bytes){#ref-for-serialize-a-mime-type-to-bytes
            link-type="dfn"}.

        4.  Return a new
            [response](#concept-response){#ref-for-concept-response⑤④
            link-type="dfn"} whose [status
            message](#concept-response-status-message){#ref-for-concept-response-status-message⑥
            link-type="dfn"} is \``OK`\`, [header
            list](#concept-response-header-list){#ref-for-concept-response-header-list①⑨
            link-type="dfn"} is « (\``Content-Type`\`,
            `mimeType`{.variable}) », and
            [body](#concept-response-body){#ref-for-concept-response-body②①
            link-type="dfn"} is `dataURLStruct`{.variable}'s
            [body](#data-url-struct-body){#ref-for-data-url-struct-body
            link-type="dfn"} [as a
            body](#byte-sequence-as-a-body){#ref-for-byte-sequence-as-a-body③
            link-type="dfn"}.

    \"`file`\"

    :   For now, unfortunate as it is, `file:`
        [URLs](https://url.spec.whatwg.org/#concept-url){#ref-for-concept-url①⑧
        link-type="dfn"} are left as an exercise for the reader.

        When in doubt, return a [network
        error](#concept-network-error){#ref-for-concept-network-error③②
        link-type="dfn"}.

    [HTTP(S) scheme](#http-scheme){#ref-for-http-scheme⑧ link-type="dfn"}

    :   Return the result of running [HTTP
        fetch](#concept-http-fetch){#ref-for-concept-http-fetch⑤
        link-type="dfn"} given `fetchParams`{.variable}.

4.  Return a [network
    error](#concept-network-error){#ref-for-concept-network-error③③
    link-type="dfn"}.
:::

::: {.algorithm algorithm="determine the environment" algorithm-for="request"}
To [determine the environment]{#request-determine-the-environment .dfn
.dfn-paneled dfn-for="request" dfn-type="dfn" noexport=""}, given a
[request](#concept-request){#ref-for-concept-request①⓪⑥ link-type="dfn"}
`request`{.variable}:

1.  If `request`{.variable}'s [reserved
    client](#concept-request-reserved-client){#ref-for-concept-request-reserved-client②
    link-type="dfn"} is non-null, then return `request`{.variable}'s
    [reserved
    client](#concept-request-reserved-client){#ref-for-concept-request-reserved-client③
    link-type="dfn"}.

2.  If `request`{.variable}'s
    [client](#concept-request-client){#ref-for-concept-request-client②③
    link-type="dfn"} is non-null, then return `request`{.variable}'s
    [client](#concept-request-client){#ref-for-concept-request-client②④
    link-type="dfn"}.

3.  Return null.
:::

### [4.3. ]{.secno}[HTTP fetch]{.content}[](#http-fetch){.self-link} {#http-fetch .heading .settled level="4.3"}

::: {.algorithm algorithm="HTTP fetch"}
To [HTTP fetch]{#concept-http-fetch .dfn .dfn-paneled dfn-type="dfn"
export=""}, given a [fetch params](#fetch-params){#ref-for-fetch-params⑦
link-type="dfn"} `fetchParams`{.variable} and an optional boolean
`makeCORSPreflight`{.variable} (default false), run these steps:

1.  Let `request`{.variable} be `fetchParams`{.variable}'s
    [request](#fetch-params-request){#ref-for-fetch-params-request①④
    link-type="dfn"}.

2.  Let `response`{.variable} and `internalResponse`{.variable} be null.

3.  If `request`{.variable}'s [service-workers
    mode](#request-service-workers-mode){#ref-for-request-service-workers-mode
    link-type="dfn"} is \"`all`\", then:

    1.  Let `requestForServiceWorker`{.variable} be a
        [clone](#concept-request-clone){#ref-for-concept-request-clone
        link-type="dfn"} of `request`{.variable}.

    2.  If `requestForServiceWorker`{.variable}'s
        [body](#concept-body){#ref-for-concept-body⑧ link-type="dfn"} is
        non-null, then:

        1.  Let `transformStream`{.variable} be a new
            [`TransformStream`{.idl}](https://streams.spec.whatwg.org/#transformstream){#ref-for-transformstream②
            link-type="idl"}.

        2.  Let `transformAlgorithm`{.variable} given `chunk`{.variable}
            be these steps:

            1.  If `fetchParams`{.variable} is
                [canceled](#fetch-params-canceled){#ref-for-fetch-params-canceled②
                link-type="dfn"}, then abort these steps.

            2.  If `chunk`{.variable} is not a
                [`Uint8Array`{.idl}](https://webidl.spec.whatwg.org/#idl-Uint8Array){#ref-for-idl-Uint8Array①
                link-type="idl"} object, then
                [terminate](#fetch-controller-terminate){#ref-for-fetch-controller-terminate①
                link-type="dfn"} `fetchParams`{.variable}'s
                [controller](#fetch-params-controller){#ref-for-fetch-params-controller⑦
                link-type="dfn"}.

            3.  Otherwise,
                [enqueue](https://streams.spec.whatwg.org/#readablestream-enqueue){#ref-for-readablestream-enqueue
                link-type="dfn"} `chunk`{.variable} in
                `transformStream`{.variable}. The user agent may split
                the chunk into
                [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#ref-for-implementation-defined①⑥
                link-type="dfn"} practical sizes and
                [enqueue](https://streams.spec.whatwg.org/#readablestream-enqueue){#ref-for-readablestream-enqueue①
                link-type="dfn"} each of them. The user agent also may
                concatenate the chunks into an
                [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#ref-for-implementation-defined①⑦
                link-type="dfn"} practical size and
                [enqueue](https://streams.spec.whatwg.org/#readablestream-enqueue){#ref-for-readablestream-enqueue②
                link-type="dfn"} it.

        3.  [Set
            up](https://streams.spec.whatwg.org/#transformstream-set-up){#ref-for-transformstream-set-up①
            link-type="dfn"} `transformStream`{.variable} with
            [*transformAlgorithm*](https://streams.spec.whatwg.org/#transformstream-set-up-transformalgorithm){#ref-for-transformstream-set-up-transformalgorithm①
            link-type="dfn"} set to `transformAlgorithm`{.variable}.

        4.  Set `requestForServiceWorker`{.variable}'s
            [body](#concept-body){#ref-for-concept-body⑨
            link-type="dfn"}'s
            [stream](#concept-body-stream){#ref-for-concept-body-stream⑦
            link-type="dfn"} to the result of
            `requestForServiceWorker`{.variable}'s
            [body](#concept-body){#ref-for-concept-body①⓪
            link-type="dfn"}'s
            [stream](#concept-body-stream){#ref-for-concept-body-stream⑧
            link-type="dfn"} [piped
            through](https://streams.spec.whatwg.org/#readablestream-pipe-through){#ref-for-readablestream-pipe-through①
            link-type="dfn"} `transformStream`{.variable}.

    3.  Let `serviceWorkerStartTime`{.variable} be the [coarsened shared
        current
        time](https://w3c.github.io/hr-time/#dfn-coarsened-shared-current-time){#ref-for-dfn-coarsened-shared-current-time①
        link-type="dfn"} given `fetchParams`{.variable}'s [cross-origin
        isolated
        capability](#fetch-params-cross-origin-isolated-capability){#ref-for-fetch-params-cross-origin-isolated-capability①
        link-type="dfn"}.

    4.  Set `response`{.variable} to the result of invoking [handle
        fetch](https://w3c.github.io/ServiceWorker/#handle-fetch){#ref-for-handle-fetch①
        link-type="dfn"} for `requestForServiceWorker`{.variable}, with
        `fetchParams`{.variable}'s
        [controller](#fetch-params-controller){#ref-for-fetch-params-controller⑧
        link-type="dfn"} and `fetchParams`{.variable}'s [cross-origin
        isolated
        capability](#fetch-params-cross-origin-isolated-capability){#ref-for-fetch-params-cross-origin-isolated-capability②
        link-type="dfn"}.
        [\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}
        [\[SW\]](#biblio-sw "Service Workers"){link-type="biblio"}

    5.  If `response`{.variable} is non-null, then:

        1.  Set `fetchParams`{.variable}'s [timing
            info](#fetch-params-timing-info){#ref-for-fetch-params-timing-info③
            link-type="dfn"}'s [final service worker start
            time](#fetch-timing-info-final-service-worker-start-time){#ref-for-fetch-timing-info-final-service-worker-start-time
            link-type="dfn"} to `serviceWorkerStartTime`{.variable}.

        2.  If `request`{.variable}'s
            [body](#concept-request-body){#ref-for-concept-request-body⑦
            link-type="dfn"} is non-null, then
            [cancel](https://streams.spec.whatwg.org/#readablestream-cancel){#ref-for-readablestream-cancel
            link-type="dfn"} `request`{.variable}'s
            [body](#concept-request-body){#ref-for-concept-request-body⑧
            link-type="dfn"} with undefined.

        3.  Set `internalResponse`{.variable} to `response`{.variable},
            if `response`{.variable} is not a [filtered
            response](#concept-filtered-response){#ref-for-concept-filtered-response①③
            link-type="dfn"}; otherwise to `response`{.variable}'s
            [internal
            response](#concept-internal-response){#ref-for-concept-internal-response①⑤
            link-type="dfn"}.

        4.  If one of the following is true

            - `response`{.variable}'s
              [type](#concept-response-type){#ref-for-concept-response-type⑦
              link-type="dfn"} is \"`error`\"

            - `request`{.variable}'s
              [mode](#concept-request-mode){#ref-for-concept-request-mode①⑤
              link-type="dfn"} is \"`same-origin`\" and
              `response`{.variable}'s
              [type](#concept-response-type){#ref-for-concept-response-type⑧
              link-type="dfn"} is \"`cors`\"

            - `request`{.variable}'s
              [mode](#concept-request-mode){#ref-for-concept-request-mode①⑥
              link-type="dfn"} is not \"`no-cors`\" and
              `response`{.variable}'s
              [type](#concept-response-type){#ref-for-concept-response-type⑨
              link-type="dfn"} is \"`opaque`\"

            - `request`{.variable}'s [redirect
              mode](#concept-request-redirect-mode){#ref-for-concept-request-redirect-mode②
              link-type="dfn"} is not \"`manual`\" and
              `response`{.variable}'s
              [type](#concept-response-type){#ref-for-concept-response-type①⓪
              link-type="dfn"} is \"`opaqueredirect`\"

            - `request`{.variable}'s [redirect
              mode](#concept-request-redirect-mode){#ref-for-concept-request-redirect-mode③
              link-type="dfn"} is not \"`follow`\" and
              `response`{.variable}'s [URL
              list](#concept-response-url-list){#ref-for-concept-response-url-list⑨
              link-type="dfn"} has more than one item.

            then return a [network
            error](#concept-network-error){#ref-for-concept-network-error③④
            link-type="dfn"}.

4.  If `response`{.variable} is null, then:

    1.  If `makeCORSPreflight`{.variable} is true and one of these
        conditions is true:

        - There is no [method cache entry
          match](#concept-cache-match-method){#ref-for-concept-cache-match-method
          link-type="dfn"} for `request`{.variable}'s
          [method](#concept-request-method){#ref-for-concept-request-method⑧
          link-type="dfn"} using `request`{.variable}, and either
          `request`{.variable}'s
          [method](#concept-request-method){#ref-for-concept-request-method⑨
          link-type="dfn"} is not a [CORS-safelisted
          method](#cors-safelisted-method){#ref-for-cors-safelisted-method②
          link-type="dfn"} or `request`{.variable}'s [use-CORS-preflight
          flag](#use-cors-preflight-flag){#ref-for-use-cors-preflight-flag③
          link-type="dfn"} is set.

        - There is at least one
          [item](https://infra.spec.whatwg.org/#list-item){#ref-for-list-item①
          link-type="dfn"} in the [CORS-unsafe request-header
          names](#cors-unsafe-request-header-names){#ref-for-cors-unsafe-request-header-names①
          link-type="dfn"} with `request`{.variable}'s [header
          list](#concept-request-header-list){#ref-for-concept-request-header-list①④
          link-type="dfn"} for which there is no [header-name cache
          entry
          match](#concept-cache-match-header){#ref-for-concept-cache-match-header
          link-type="dfn"} using `request`{.variable}.

        Then:

        1.  Let `preflightResponse`{.variable} be the result of running
            [CORS-preflight
            fetch](#cors-preflight-fetch-0){#ref-for-cors-preflight-fetch-0①
            link-type="dfn"} given `request`{.variable}.

        2.  If `preflightResponse`{.variable} is a [network
            error](#concept-network-error){#ref-for-concept-network-error③⑤
            link-type="dfn"}, then return
            `preflightResponse`{.variable}.

        This step checks the [CORS-preflight
        cache](#concept-cache){#ref-for-concept-cache link-type="dfn"}
        and if there is no suitable entry it performs a [CORS-preflight
        fetch](#cors-preflight-fetch-0){#ref-for-cors-preflight-fetch-0②
        link-type="dfn"} which, if successful, populates the cache. The
        purpose of the [CORS-preflight
        fetch](#cors-preflight-fetch-0){#ref-for-cors-preflight-fetch-0③
        link-type="dfn"} is to ensure the
        [fetched](#concept-fetch){#ref-for-concept-fetch②⑤
        link-type="dfn"} resource is familiar with the [CORS
        protocol](#cors-protocol){#ref-for-cors-protocol①⑥
        link-type="dfn"}. The cache is there to minimize the number of
        [CORS-preflight
        fetches](#cors-preflight-fetch-0){#ref-for-cors-preflight-fetch-0④
        link-type="dfn"}.

    2.  If `request`{.variable}'s [redirect
        mode](#concept-request-redirect-mode){#ref-for-concept-request-redirect-mode④
        link-type="dfn"} is \"`follow`\", then set
        `request`{.variable}'s [service-workers
        mode](#request-service-workers-mode){#ref-for-request-service-workers-mode①
        link-type="dfn"} to \"`none`\".

        Redirects coming from the network (as opposed to from a service
        worker) are not to be exposed to a service worker.

    3.  Set `response`{.variable} and `internalResponse`{.variable} to
        the result of running [HTTP-network-or-cache
        fetch](#concept-http-network-or-cache-fetch){#ref-for-concept-http-network-or-cache-fetch
        link-type="dfn"} given `fetchParams`{.variable}.

    4.  If `request`{.variable}'s [response
        tainting](#concept-request-response-tainting){#ref-for-concept-request-response-tainting①①
        link-type="dfn"} is \"`cors`\" and a [CORS
        check](#concept-cors-check){#ref-for-concept-cors-check
        link-type="dfn"} for `request`{.variable} and
        `response`{.variable} returns failure, then return a [network
        error](#concept-network-error){#ref-for-concept-network-error③⑥
        link-type="dfn"}.

        As the [CORS
        check](#concept-cors-check){#ref-for-concept-cors-check①
        link-type="dfn"} is not to be applied to
        [responses](#concept-response){#ref-for-concept-response⑤⑤
        link-type="dfn"} whose
        [status](#concept-response-status){#ref-for-concept-response-status⑨
        link-type="dfn"} is 304 or 407, or
        [responses](#concept-response){#ref-for-concept-response⑤⑥
        link-type="dfn"} from a service worker for that matter, it is
        applied here.

    5.  If the [TAO
        check](#concept-tao-check){#ref-for-concept-tao-check
        link-type="dfn"} for `request`{.variable} and
        `response`{.variable} returns failure, then set
        `request`{.variable}'s [timing allow failed
        flag](#timing-allow-failed){#ref-for-timing-allow-failed③
        link-type="dfn"}.

5.  If either `request`{.variable}'s [response
    tainting](#concept-request-response-tainting){#ref-for-concept-request-response-tainting①②
    link-type="dfn"} or `response`{.variable}'s
    [type](#concept-response-type){#ref-for-concept-response-type①①
    link-type="dfn"} is \"`opaque`\", and the [cross-origin resource
    policy
    check](#cross-origin-resource-policy-check){#ref-for-cross-origin-resource-policy-check
    link-type="dfn"} with `request`{.variable}'s
    [origin](#concept-request-origin){#ref-for-concept-request-origin①⑤
    link-type="dfn"}, `request`{.variable}'s
    [client](#concept-request-client){#ref-for-concept-request-client②⑤
    link-type="dfn"}, `request`{.variable}'s
    [destination](#concept-request-destination){#ref-for-concept-request-destination①⑦
    link-type="dfn"}, and `internalResponse`{.variable} returns
    **blocked**, then return a [network
    error](#concept-network-error){#ref-for-concept-network-error③⑦
    link-type="dfn"}.

    The [cross-origin resource policy
    check](#cross-origin-resource-policy-check){#ref-for-cross-origin-resource-policy-check①
    link-type="dfn"} runs for responses coming from the network and
    responses coming from the service worker. This is different from the
    [CORS check](#concept-cors-check){#ref-for-concept-cors-check②
    link-type="dfn"}, as `request`{.variable}'s
    [client](#concept-request-client){#ref-for-concept-request-client②⑥
    link-type="dfn"} and the service worker can have different embedder
    policies.

6.  If `internalResponse`{.variable}'s
    [status](#concept-response-status){#ref-for-concept-response-status①⓪
    link-type="dfn"} is a [redirect
    status](#redirect-status){#ref-for-redirect-status②
    link-type="dfn"}:

    1.  If `internalResponse`{.variable}'s
        [status](#concept-response-status){#ref-for-concept-response-status①①
        link-type="dfn"} is not 303, `request`{.variable}'s
        [body](#concept-request-body){#ref-for-concept-request-body⑨
        link-type="dfn"} is non-null, and the
        [connection](#concept-connection){#ref-for-concept-connection①①
        link-type="dfn"} uses HTTP/2, then user agents may, and are even
        encouraged to, transmit an `RST_STREAM` frame.

        303 is excluded as certain communities ascribe special status to
        it.

    2.  Switch on `request`{.variable}'s [redirect
        mode](#concept-request-redirect-mode){#ref-for-concept-request-redirect-mode⑤
        link-type="dfn"}:

        \"`error`\"

        :   1.  Set `response`{.variable} to a [network
                error](#concept-network-error){#ref-for-concept-network-error③⑧
                link-type="dfn"}.

        \"`manual`\"

        :   1.  If `request`{.variable}'s
                [mode](#concept-request-mode){#ref-for-concept-request-mode①⑦
                link-type="dfn"} is \"`navigate`\", then set
                `fetchParams`{.variable}'s
                [controller](#fetch-params-controller){#ref-for-fetch-params-controller⑨
                link-type="dfn"}'s [next manual redirect
                steps](#fetch-controller-next-manual-redirect-steps){#ref-for-fetch-controller-next-manual-redirect-steps②
                link-type="dfn"} to run [HTTP-redirect
                fetch](#concept-http-redirect-fetch){#ref-for-concept-http-redirect-fetch
                link-type="dfn"} given `fetchParams`{.variable} and
                `response`{.variable}.

            2.  Otherwise, set `response`{.variable} to an
                [opaque-redirect filtered
                response](#concept-filtered-response-opaque-redirect){#ref-for-concept-filtered-response-opaque-redirect④
                link-type="dfn"} whose [internal
                response](#concept-internal-response){#ref-for-concept-internal-response①⑥
                link-type="dfn"} is `internalResponse`{.variable}.

        \"`follow`\"

        :   1.  Set `response`{.variable} to the result of running
                [HTTP-redirect
                fetch](#concept-http-redirect-fetch){#ref-for-concept-http-redirect-fetch①
                link-type="dfn"} given `fetchParams`{.variable} and
                `response`{.variable}.

7.  Return `response`{.variable}. [Typically
    `internalResponse`{.variable}'s
    [body](#concept-response-body){#ref-for-concept-response-body②②
    link-type="dfn"}'s
    [stream](#concept-body-stream){#ref-for-concept-body-stream⑨
    link-type="dfn"} is still being enqueued to after returning.]{.note}
:::

### [4.4. ]{.secno}[HTTP-redirect fetch]{.content}[](#http-redirect-fetch){.self-link} {#http-redirect-fetch .heading .settled level="4.4"}

::: {.algorithm algorithm="HTTP-redirect fetch"}
To [HTTP-redirect fetch]{#concept-http-redirect-fetch .dfn .dfn-paneled
dfn-type="dfn" export=""}, given a [fetch
params](#fetch-params){#ref-for-fetch-params⑧ link-type="dfn"}
`fetchParams`{.variable} and a
[response](#concept-response){#ref-for-concept-response⑤⑦
link-type="dfn"} `response`{.variable}, run these steps:

1.  Let `request`{.variable} be `fetchParams`{.variable}'s
    [request](#fetch-params-request){#ref-for-fetch-params-request①⑤
    link-type="dfn"}.

2.  Let `internalResponse`{.variable} be `response`{.variable}, if
    `response`{.variable} is not a [filtered
    response](#concept-filtered-response){#ref-for-concept-filtered-response①④
    link-type="dfn"}; otherwise `response`{.variable}'s [internal
    response](#concept-internal-response){#ref-for-concept-internal-response①⑦
    link-type="dfn"}.

3.  Let `locationURL`{.variable} be `internalResponse`{.variable}'s
    [location
    URL](#concept-response-location-url){#ref-for-concept-response-location-url①
    link-type="dfn"} given `request`{.variable}'s [current
    URL](#concept-request-current-url){#ref-for-concept-request-current-url②②
    link-type="dfn"}'s
    [fragment](https://url.spec.whatwg.org/#concept-url-fragment){#ref-for-concept-url-fragment②
    link-type="dfn"}.

4.  If `locationURL`{.variable} is null, then return
    `response`{.variable}.

5.  If `locationURL`{.variable} is failure, then return a [network
    error](#concept-network-error){#ref-for-concept-network-error③⑨
    link-type="dfn"}.

6.  If `locationURL`{.variable}'s
    [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#ref-for-concept-url-scheme①④
    link-type="dfn"} is not an [HTTP(S)
    scheme](#http-scheme){#ref-for-http-scheme⑨ link-type="dfn"}, then
    return a [network
    error](#concept-network-error){#ref-for-concept-network-error④⓪
    link-type="dfn"}.

7.  If `request`{.variable}'s [redirect
    count](#concept-request-redirect-count){#ref-for-concept-request-redirect-count①
    link-type="dfn"} is 20, then return a [network
    error](#concept-network-error){#ref-for-concept-network-error④①
    link-type="dfn"}.

8.  Increase `request`{.variable}'s [redirect
    count](#concept-request-redirect-count){#ref-for-concept-request-redirect-count②
    link-type="dfn"} by 1.

9.  If `request`{.variable}'s
    [mode](#concept-request-mode){#ref-for-concept-request-mode①⑧
    link-type="dfn"} is \"`cors`\", `locationURL`{.variable} [includes
    credentials](https://url.spec.whatwg.org/#include-credentials){#ref-for-include-credentials
    link-type="dfn"}, and `request`{.variable}'s
    [origin](#concept-request-origin){#ref-for-concept-request-origin①⑥
    link-type="dfn"} is not [same
    origin](https://html.spec.whatwg.org/multipage/browsers.html#same-origin){#ref-for-same-origin⑦
    link-type="dfn"} with `locationURL`{.variable}'s
    [origin](https://url.spec.whatwg.org/#concept-url-origin){#ref-for-concept-url-origin①⑥
    link-type="dfn"}, then return a [network
    error](#concept-network-error){#ref-for-concept-network-error④②
    link-type="dfn"}.

10. If `request`{.variable}'s [response
    tainting](#concept-request-response-tainting){#ref-for-concept-request-response-tainting①③
    link-type="dfn"} is \"`cors`\" and `locationURL`{.variable}
    [includes
    credentials](https://url.spec.whatwg.org/#include-credentials){#ref-for-include-credentials①
    link-type="dfn"}, then return a [network
    error](#concept-network-error){#ref-for-concept-network-error④③
    link-type="dfn"}.

    This catches a cross-origin resource redirecting to a same-origin
    URL.

11. If `internalResponse`{.variable}'s
    [status](#concept-response-status){#ref-for-concept-response-status①②
    link-type="dfn"} is not 303, `request`{.variable}'s
    [body](#concept-request-body){#ref-for-concept-request-body①⓪
    link-type="dfn"} is non-null, and `request`{.variable}'s
    [body](#concept-request-body){#ref-for-concept-request-body①①
    link-type="dfn"}'s
    [source](#concept-body-source){#ref-for-concept-body-source①
    link-type="dfn"} is null, then return a [network
    error](#concept-network-error){#ref-for-concept-network-error④④
    link-type="dfn"}.

12. If one of the following is true

    - `internalResponse`{.variable}'s
      [status](#concept-response-status){#ref-for-concept-response-status①③
      link-type="dfn"} is 301 or 302 and `request`{.variable}'s
      [method](#concept-request-method){#ref-for-concept-request-method①⓪
      link-type="dfn"} is \``POST`\`

    - `internalResponse`{.variable}'s
      [status](#concept-response-status){#ref-for-concept-response-status①④
      link-type="dfn"} is 303 and `request`{.variable}'s
      [method](#concept-request-method){#ref-for-concept-request-method①①
      link-type="dfn"} is not \``GET`\` or \``HEAD`\`

    then:

    1.  Set `request`{.variable}'s
        [method](#concept-request-method){#ref-for-concept-request-method①②
        link-type="dfn"} to \``GET`\` and `request`{.variable}'s
        [body](#concept-request-body){#ref-for-concept-request-body①②
        link-type="dfn"} to null.

    2.  [For
        each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate①⓪
        link-type="dfn"} `headerName`{.variable} of [request-body-header
        name](#request-body-header-name){#ref-for-request-body-header-name
        link-type="dfn"},
        [delete](#concept-header-list-delete){#ref-for-concept-header-list-delete
        link-type="dfn"} `headerName`{.variable} from
        `request`{.variable}'s [header
        list](#concept-request-header-list){#ref-for-concept-request-header-list①⑤
        link-type="dfn"}.

13. If `request`{.variable}'s [current
    URL](#concept-request-current-url){#ref-for-concept-request-current-url②③
    link-type="dfn"}'s
    [origin](https://url.spec.whatwg.org/#concept-url-origin){#ref-for-concept-url-origin①⑦
    link-type="dfn"} is not [same
    origin](https://html.spec.whatwg.org/multipage/browsers.html#same-origin){#ref-for-same-origin⑧
    link-type="dfn"} with `locationURL`{.variable}'s
    [origin](https://url.spec.whatwg.org/#concept-url-origin){#ref-for-concept-url-origin①⑧
    link-type="dfn"}, then [for
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate①①
    link-type="dfn"} `headerName`{.variable} of [CORS non-wildcard
    request-header
    name](#cors-non-wildcard-request-header-name){#ref-for-cors-non-wildcard-request-header-name
    link-type="dfn"},
    [delete](#concept-header-list-delete){#ref-for-concept-header-list-delete①
    link-type="dfn"} `headerName`{.variable} from `request`{.variable}'s
    [header
    list](#concept-request-header-list){#ref-for-concept-request-header-list①⑥
    link-type="dfn"}.

    I.e., the moment another origin is seen after the initial request,
    the \``Authorization`\` header is removed.

14. If `request`{.variable}'s
    [body](#concept-request-body){#ref-for-concept-request-body①③
    link-type="dfn"} is non-null, then set `request`{.variable}'s
    [body](#concept-request-body){#ref-for-concept-request-body①④
    link-type="dfn"} to the
    [body](#body-with-type-body){#ref-for-body-with-type-body③
    link-type="dfn"} of the result of [safely
    extracting](#bodyinit-safely-extract){#ref-for-bodyinit-safely-extract④
    link-type="dfn"} `request`{.variable}'s
    [body](#concept-request-body){#ref-for-concept-request-body①⑤
    link-type="dfn"}'s
    [source](#concept-body-source){#ref-for-concept-body-source②
    link-type="dfn"}.

    `request`{.variable}'s
    [body](#concept-request-body){#ref-for-concept-request-body①⑥
    link-type="dfn"}'s
    [source](#concept-body-source){#ref-for-concept-body-source③
    link-type="dfn"}'s nullity has already been checked.

15. Let `timingInfo`{.variable} be `fetchParams`{.variable}'s [timing
    info](#fetch-params-timing-info){#ref-for-fetch-params-timing-info④
    link-type="dfn"}.

16. Set `timingInfo`{.variable}'s [redirect end
    time](#fetch-timing-info-redirect-end-time){#ref-for-fetch-timing-info-redirect-end-time
    link-type="dfn"} and [post-redirect start
    time](#fetch-timing-info-post-redirect-start-time){#ref-for-fetch-timing-info-post-redirect-start-time②
    link-type="dfn"} to the [coarsened shared current
    time](https://w3c.github.io/hr-time/#dfn-coarsened-shared-current-time){#ref-for-dfn-coarsened-shared-current-time②
    link-type="dfn"} given `fetchParams`{.variable}'s [cross-origin
    isolated
    capability](#fetch-params-cross-origin-isolated-capability){#ref-for-fetch-params-cross-origin-isolated-capability③
    link-type="dfn"}.

17. If `timingInfo`{.variable}'s [redirect start
    time](#fetch-timing-info-redirect-start-time){#ref-for-fetch-timing-info-redirect-start-time
    link-type="dfn"} is 0, then set `timingInfo`{.variable}'s [redirect
    start
    time](#fetch-timing-info-redirect-start-time){#ref-for-fetch-timing-info-redirect-start-time①
    link-type="dfn"} to `timingInfo`{.variable}'s [start
    time](#fetch-timing-info-start-time){#ref-for-fetch-timing-info-start-time③
    link-type="dfn"}.

18. [Append](https://infra.spec.whatwg.org/#list-append){#ref-for-list-append⑨
    link-type="dfn"} `locationURL`{.variable} to `request`{.variable}'s
    [URL
    list](#concept-request-url-list){#ref-for-concept-request-url-list⑤
    link-type="dfn"}.

19. Invoke [set `request`{.variable}'s referrer policy on
    redirect](https://w3c.github.io/webappsec-referrer-policy/#set-requests-referrer-policy-on-redirect){#ref-for-set-requests-referrer-policy-on-redirect
    link-type="dfn"} on `request`{.variable} and
    `internalResponse`{.variable}.
    [\[REFERRER\]](#biblio-referrer "Referrer Policy"){link-type="biblio"}

20. Let `recursive`{.variable} be true.

21. If `request`{.variable}'s [redirect
    mode](#concept-request-redirect-mode){#ref-for-concept-request-redirect-mode⑥
    link-type="dfn"} is \"`manual`\", then:

    1.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert②①
        link-type="dfn"}: `request`{.variable}'s
        [mode](#concept-request-mode){#ref-for-concept-request-mode①⑨
        link-type="dfn"} is \"`navigate`\".

    2.  Set `recursive`{.variable} to false.

22. Return the result of running [main
    fetch](#concept-main-fetch){#ref-for-concept-main-fetch①
    link-type="dfn"} given `fetchParams`{.variable} and
    `recursive`{.variable}.

    This has to invoke [main
    fetch](#concept-main-fetch){#ref-for-concept-main-fetch②
    link-type="dfn"} to get `request`{.variable}'s [response
    tainting](#concept-request-response-tainting){#ref-for-concept-request-response-tainting①④
    link-type="dfn"} correct.
:::

### [4.5. ]{.secno}[HTTP-network-or-cache fetch]{.content}[](#http-network-or-cache-fetch){.self-link} {#http-network-or-cache-fetch .heading .settled level="4.5"}

::: {.algorithm algorithm="HTTP-network-or-cache fetch"}
To [HTTP-network-or-cache fetch]{#concept-http-network-or-cache-fetch
.dfn .dfn-paneled dfn-type="dfn" noexport=""}, given a [fetch
params](#fetch-params){#ref-for-fetch-params⑨ link-type="dfn"}
`fetchParams`{.variable}, an optional boolean
`isAuthenticationFetch`{.variable} (default false), and an optional
boolean `isNewConnectionFetch`{.variable} (default false), run these
steps:

Some implementations might support caching of partial content, as per
HTTP Caching. However, this is not widely supported by browser caches.
[\[HTTP-CACHING\]](#biblio-http-caching "HTTP Caching"){link-type="biblio"}

1.  Let `request`{.variable} be `fetchParams`{.variable}'s
    [request](#fetch-params-request){#ref-for-fetch-params-request①⑥
    link-type="dfn"}.

2.  Let `httpFetchParams`{.variable} be null.

3.  Let `httpRequest`{.variable} be null.

4.  Let `response`{.variable} be null.

5.  Let `storedResponse`{.variable} be null.

6.  Let `httpCache`{.variable} be null.

7.  Let the `revalidatingFlag`{.variable} be unset.

8.  Run these steps, but [abort
    when](https://infra.spec.whatwg.org/#abort-when){#ref-for-abort-when
    link-type="dfn"} `fetchParams`{.variable} is
    [canceled](#fetch-params-canceled){#ref-for-fetch-params-canceled③
    link-type="dfn"}:

    1.  If `request`{.variable}'s [traversable for user
        prompts](#concept-request-window){#ref-for-concept-request-window⑥
        link-type="dfn"} is \"`no-traversable`\" and
        `request`{.variable}'s [redirect
        mode](#concept-request-redirect-mode){#ref-for-concept-request-redirect-mode⑦
        link-type="dfn"} is \"`error`\", then set
        `httpFetchParams`{.variable} to `fetchParams`{.variable} and
        `httpRequest`{.variable} to `request`{.variable}.

    2.  Otherwise:

        1.  Set `httpRequest`{.variable} to a
            [clone](#concept-request-clone){#ref-for-concept-request-clone①
            link-type="dfn"} of `request`{.variable}.

            Implementations are encouraged to avoid teeing
            `request`{.variable}'s
            [body](#concept-request-body){#ref-for-concept-request-body①⑦
            link-type="dfn"}'s
            [stream](#concept-body-stream){#ref-for-concept-body-stream①⓪
            link-type="dfn"} when `request`{.variable}'s
            [body](#concept-request-body){#ref-for-concept-request-body①⑧
            link-type="dfn"}'s
            [source](#concept-body-source){#ref-for-concept-body-source④
            link-type="dfn"} is null as only a single body is needed in
            that case. E.g., when `request`{.variable}'s
            [body](#concept-request-body){#ref-for-concept-request-body①⑨
            link-type="dfn"}'s
            [source](#concept-body-source){#ref-for-concept-body-source⑤
            link-type="dfn"} is null, redirects and authentication will
            end up failing the fetch.

        2.  Set `httpFetchParams`{.variable} to a copy of
            `fetchParams`{.variable}.

        3.  Set `httpFetchParams`{.variable}'s
            [request](#fetch-params-request){#ref-for-fetch-params-request①⑦
            link-type="dfn"} to `httpRequest`{.variable}.

        If user prompts or redirects are possible, then the user agent
        might need to re-send the request with a new set of headers
        after the user answers the prompt or the redirect location is
        determined. At that time, the original request body might have
        been partially sent already, so we need to clone the request
        (including the body) beforehand so that we have a spare copy
        available.

    3.  Let `includeCredentials`{.variable} be true if one of

        - `request`{.variable}'s [credentials
          mode](#concept-request-credentials-mode){#ref-for-concept-request-credentials-mode①⓪
          link-type="dfn"} is \"`include`\"
        - `request`{.variable}'s [credentials
          mode](#concept-request-credentials-mode){#ref-for-concept-request-credentials-mode①①
          link-type="dfn"} is \"`same-origin`\" and
          `request`{.variable}'s [response
          tainting](#concept-request-response-tainting){#ref-for-concept-request-response-tainting①⑤
          link-type="dfn"} is \"`basic`\"

        is true; otherwise false.

    4.  If [Cross-Origin-Embedder-Policy allows
        credentials](#cross-origin-embedder-policy-allows-credentials){#ref-for-cross-origin-embedder-policy-allows-credentials
        link-type="dfn"} with `request`{.variable} returns false, then
        set `includeCredentials`{.variable} to false.

    5.  Let `contentLength`{.variable} be `httpRequest`{.variable}'s
        [body](#concept-request-body){#ref-for-concept-request-body②⓪
        link-type="dfn"}'s
        [length](#concept-body-total-bytes){#ref-for-concept-body-total-bytes①
        link-type="dfn"}, if `httpRequest`{.variable}'s
        [body](#concept-request-body){#ref-for-concept-request-body②①
        link-type="dfn"} is non-null; otherwise null.

    6.  Let `contentLengthHeaderValue`{.variable} be null.

    7.  If `httpRequest`{.variable}'s
        [body](#concept-request-body){#ref-for-concept-request-body②②
        link-type="dfn"} is null and `httpRequest`{.variable}'s
        [method](#concept-request-method){#ref-for-concept-request-method①③
        link-type="dfn"} is \``POST`\` or \``PUT`\`, then set
        `contentLengthHeaderValue`{.variable} to \``0`\`.

    8.  If `contentLength`{.variable} is non-null, then set
        `contentLengthHeaderValue`{.variable} to
        `contentLength`{.variable},
        [serialized](#serialize-an-integer){#ref-for-serialize-an-integer⑦
        link-type="dfn"} and [isomorphic
        encoded](https://infra.spec.whatwg.org/#isomorphic-encode){#ref-for-isomorphic-encode⑨
        link-type="dfn"}.

    9.  If `contentLengthHeaderValue`{.variable} is non-null, then
        [append](#concept-header-list-append){#ref-for-concept-header-list-append⑤
        link-type="dfn"} (\``Content-Length`\`,
        `contentLengthHeaderValue`{.variable}) to
        `httpRequest`{.variable}'s [header
        list](#concept-request-header-list){#ref-for-concept-request-header-list①⑦
        link-type="dfn"}.

    10. If `contentLength`{.variable} is non-null and
        `httpRequest`{.variable}'s
        [keepalive](#request-keepalive-flag){#ref-for-request-keepalive-flag①
        link-type="dfn"} is true, then:

        1.  Let `inflightKeepaliveBytes`{.variable} be 0.

        2.  Let `group`{.variable} be `httpRequest`{.variable}'s
            [client](#concept-request-client){#ref-for-concept-request-client②⑦
            link-type="dfn"}'s [fetch
            group](#concept-fetch-group){#ref-for-concept-fetch-group③
            link-type="dfn"}.

        3.  Let `inflightRecords`{.variable} be the set of [fetch
            records](#concept-fetch-record){#ref-for-concept-fetch-record⑦
            link-type="dfn"} in `group`{.variable} whose
            [request](#concept-fetch-record-request){#ref-for-concept-fetch-record-request②
            link-type="dfn"}'s
            [keepalive](#request-keepalive-flag){#ref-for-request-keepalive-flag②
            link-type="dfn"} is true and [done
            flag](#done-flag){#ref-for-done-flag③ link-type="dfn"} is
            unset.

        4.  [For
            each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate①②
            link-type="dfn"} `fetchRecord`{.variable} of
            `inflightRecords`{.variable}:

            1.  Let `inflightRequest`{.variable} be
                `fetchRecord`{.variable}'s
                [request](#concept-fetch-record-request){#ref-for-concept-fetch-record-request③
                link-type="dfn"}.

            2.  Increment `inflightKeepaliveBytes`{.variable} by
                `inflightRequest`{.variable}'s
                [body](#concept-request-body){#ref-for-concept-request-body②③
                link-type="dfn"}'s
                [length](#concept-body-total-bytes){#ref-for-concept-body-total-bytes②
                link-type="dfn"}.

        5.  If the sum of `contentLength`{.variable} and
            `inflightKeepaliveBytes`{.variable} is greater than 64
            kibibytes, then return a [network
            error](#concept-network-error){#ref-for-concept-network-error④⑤
            link-type="dfn"}.

        The above limit ensures that requests that are allowed to
        outlive the [environment settings
        object](https://html.spec.whatwg.org/multipage/webappapis.html#environment-settings-object){#ref-for-environment-settings-object⑧
        link-type="dfn"} and contain a body, have a bounded size and are
        not allowed to stay alive indefinitely.

    11. If `httpRequest`{.variable}'s
        [referrer](#concept-request-referrer){#ref-for-concept-request-referrer④
        link-type="dfn"} is a
        [URL](https://url.spec.whatwg.org/#concept-url){#ref-for-concept-url①⑨
        link-type="dfn"}, then:

        1.  Let `referrerValue`{.variable} be `httpRequest`{.variable}'s
            [referrer](#concept-request-referrer){#ref-for-concept-request-referrer⑤
            link-type="dfn"},
            [serialized](https://url.spec.whatwg.org/#concept-url-serializer){#ref-for-concept-url-serializer①
            link-type="dfn"} and [isomorphic
            encoded](https://infra.spec.whatwg.org/#isomorphic-encode){#ref-for-isomorphic-encode①⓪
            link-type="dfn"}.

        2.  [Append](#concept-header-list-append){#ref-for-concept-header-list-append⑥
            link-type="dfn"} (\``Referer`\`, `referrerValue`{.variable})
            to `httpRequest`{.variable}'s [header
            list](#concept-request-header-list){#ref-for-concept-request-header-list①⑧
            link-type="dfn"}.

    12. [Append a request \``Origin`\`
        header](#append-a-request-origin-header){#ref-for-append-a-request-origin-header
        link-type="dfn"} for `httpRequest`{.variable}.

    13. [Append the Fetch metadata headers for
        `httpRequest`{.variable}](https://w3c.github.io/webappsec-fetch-metadata/#abstract-opdef-append-the-fetch-metadata-headers-for-a-request){#ref-for-abstract-opdef-append-the-fetch-metadata-headers-for-a-request
        link-type="abstract-op"}.
        [\[FETCH-METADATA\]](#biblio-fetch-metadata "Fetch Metadata Request Headers"){link-type="biblio"}

    14. If `httpRequest`{.variable}'s
        [initiator](#concept-request-initiator){#ref-for-concept-request-initiator⑤
        link-type="dfn"} is \"`prefetch`\", then [set a structured field
        value](#concept-header-list-set-structured-header){#ref-for-concept-header-list-set-structured-header
        link-type="dfn"} given
        (\`[`Sec-Purpose`](#http-sec-purpose){#ref-for-http-sec-purpose①
        link-type="http-header"}\`, the
        [token](https://httpwg.org/specs/rfc9651.html#token){#ref-for-token②
        link-type="dfn"} `prefetch`) in `httpRequest`{.variable}'s
        [header
        list](#concept-request-header-list){#ref-for-concept-request-header-list①⑨
        link-type="dfn"}.

    15. If `httpRequest`{.variable}'s [header
        list](#concept-request-header-list){#ref-for-concept-request-header-list②⓪
        link-type="dfn"} [does not
        contain](#header-list-contains){#ref-for-header-list-contains①②
        link-type="dfn"} \``User-Agent`\`, then user agents should
        [append](#concept-header-list-append){#ref-for-concept-header-list-append⑦
        link-type="dfn"} (\``User-Agent`\`, [default \``User-Agent`\`
        value](#default-user-agent-value){#ref-for-default-user-agent-value
        link-type="dfn"}) to `httpRequest`{.variable}'s [header
        list](#concept-request-header-list){#ref-for-concept-request-header-list②①
        link-type="dfn"}.

    16. If `httpRequest`{.variable}'s [cache
        mode](#concept-request-cache-mode){#ref-for-concept-request-cache-mode①
        link-type="dfn"} is \"`default`\" and `httpRequest`{.variable}'s
        [header
        list](#concept-request-header-list){#ref-for-concept-request-header-list②②
        link-type="dfn"}
        [contains](#header-list-contains){#ref-for-header-list-contains①③
        link-type="dfn"} \``If-Modified-Since`\`, \``If-None-Match`\`,
        \``If-Unmodified-Since`\`, \``If-Match`\`, or \``If-Range`\`,
        then set `httpRequest`{.variable}'s [cache
        mode](#concept-request-cache-mode){#ref-for-concept-request-cache-mode②
        link-type="dfn"} to \"`no-store`\".

    17. If `httpRequest`{.variable}'s [cache
        mode](#concept-request-cache-mode){#ref-for-concept-request-cache-mode③
        link-type="dfn"} is \"`no-cache`\", `httpRequest`{.variable}'s
        [prevent no-cache cache-control header modification
        flag](#no-cache-prevent-cache-control){#ref-for-no-cache-prevent-cache-control
        link-type="dfn"} is unset, and `httpRequest`{.variable}'s
        [header
        list](#concept-request-header-list){#ref-for-concept-request-header-list②③
        link-type="dfn"} [does not
        contain](#header-list-contains){#ref-for-header-list-contains①④
        link-type="dfn"} \``Cache-Control`\`, then
        [append](#concept-header-list-append){#ref-for-concept-header-list-append⑧
        link-type="dfn"} (\``Cache-Control`\`, \``max-age=0`\`) to
        `httpRequest`{.variable}'s [header
        list](#concept-request-header-list){#ref-for-concept-request-header-list②④
        link-type="dfn"}.

    18. If `httpRequest`{.variable}'s [cache
        mode](#concept-request-cache-mode){#ref-for-concept-request-cache-mode④
        link-type="dfn"} is \"`no-store`\" or \"`reload`\", then:

        1.  If `httpRequest`{.variable}'s [header
            list](#concept-request-header-list){#ref-for-concept-request-header-list②⑤
            link-type="dfn"} [does not
            contain](#header-list-contains){#ref-for-header-list-contains①⑤
            link-type="dfn"} \``Pragma`\`, then
            [append](#concept-header-list-append){#ref-for-concept-header-list-append⑨
            link-type="dfn"} (\``Pragma`\`, \``no-cache`\`) to
            `httpRequest`{.variable}'s [header
            list](#concept-request-header-list){#ref-for-concept-request-header-list②⑥
            link-type="dfn"}.

        2.  If `httpRequest`{.variable}'s [header
            list](#concept-request-header-list){#ref-for-concept-request-header-list②⑦
            link-type="dfn"} [does not
            contain](#header-list-contains){#ref-for-header-list-contains①⑥
            link-type="dfn"} \``Cache-Control`\`, then
            [append](#concept-header-list-append){#ref-for-concept-header-list-append①⓪
            link-type="dfn"} (\``Cache-Control`\`, \``no-cache`\`) to
            `httpRequest`{.variable}'s [header
            list](#concept-request-header-list){#ref-for-concept-request-header-list②⑧
            link-type="dfn"}.

    19. If `httpRequest`{.variable}'s [header
        list](#concept-request-header-list){#ref-for-concept-request-header-list②⑨
        link-type="dfn"}
        [contains](#header-list-contains){#ref-for-header-list-contains①⑦
        link-type="dfn"} \``Range`\`, then
        [append](#concept-header-list-append){#ref-for-concept-header-list-append①①
        link-type="dfn"} (\``Accept-Encoding`\`, \``identity`\`) to
        `httpRequest`{.variable}'s [header
        list](#concept-request-header-list){#ref-for-concept-request-header-list③⓪
        link-type="dfn"}.

        ::: {.note role="note"}
        This avoids a failure when [handling content
        codings](#handle-content-codings){#ref-for-handle-content-codings
        link-type="dfn"} with a part of an encoded
        [response](#concept-response){#ref-for-concept-response⑤⑧
        link-type="dfn"}.

        Additionally, [many
        servers](https://jakearchibald.github.io/accept-encoding-range-test/)
        mistakenly ignore \``Range`\` headers if a non-identity encoding
        is accepted.
        :::

    20. Modify `httpRequest`{.variable}'s [header
        list](#concept-request-header-list){#ref-for-concept-request-header-list③①
        link-type="dfn"} per HTTP. Do not
        [append](#concept-header-list-append){#ref-for-concept-header-list-append①②
        link-type="dfn"} a given
        [header](#concept-header){#ref-for-concept-header⑤⓪
        link-type="dfn"} if `httpRequest`{.variable}'s [header
        list](#concept-request-header-list){#ref-for-concept-request-header-list③②
        link-type="dfn"}
        [contains](#header-list-contains){#ref-for-header-list-contains①⑧
        link-type="dfn"} that
        [header](#concept-header){#ref-for-concept-header⑤①
        link-type="dfn"}'s
        [name](#concept-header-name){#ref-for-concept-header-name②⓪
        link-type="dfn"}.

        It would be great if we could make this more normative somehow.
        At this point
        [headers](#concept-header){#ref-for-concept-header⑤②
        link-type="dfn"} such as \``Accept-Encoding`\`,
        \``Connection`\`, \``DNT`\`, and \``Host`\`, are to be
        [appended](#concept-header-list-append){#ref-for-concept-header-list-append①③
        link-type="dfn"} if necessary.

        \``Accept`\`, \``Accept-Charset`\`, and \``Accept-Language`\`
        must not be included at this point.

        \``Accept`\` and \``Accept-Language`\` are already included
        (unless
        [`fetch()`](#dom-global-fetch){#ref-for-dom-global-fetch③
        .idl-code link-type="method"} is used, which does not include
        the latter by default), and \``Accept-Charset`\` is a waste of
        bytes. See [HTTP header layer
        division](#http-header-layer-division){#ref-for-http-header-layer-division
        link-type="dfn"} for more details.

    21. If `includeCredentials`{.variable} is true, then:

        1.  If the user agent is not configured to block cookies for
            `httpRequest`{.variable} (see [section
            7](https://httpwg.org/specs/rfc6265.html#privacy-considerations)
            of
            [\[COOKIES\]](#biblio-cookies "HTTP State Management Mechanism"){link-type="biblio"}),
            then:

            1.  Let `cookies`{.variable} be the result of running the
                \"cookie-string\" algorithm (see [section
                5.4](https://httpwg.org/specs/rfc6265.html#cookie) of
                [\[COOKIES\]](#biblio-cookies "HTTP State Management Mechanism"){link-type="biblio"})
                with the user agent's cookie store and
                `httpRequest`{.variable}'s [current
                URL](#concept-request-current-url){#ref-for-concept-request-current-url②④
                link-type="dfn"}.

            2.  If `cookies`{.variable} is not the empty string, then
                [append](#concept-header-list-append){#ref-for-concept-header-list-append①④
                link-type="dfn"} (\``Cookie`\`, `cookies`{.variable}) to
                `httpRequest`{.variable}'s [header
                list](#concept-request-header-list){#ref-for-concept-request-header-list③③
                link-type="dfn"}.

        2.  If `httpRequest`{.variable}'s [header
            list](#concept-request-header-list){#ref-for-concept-request-header-list③④
            link-type="dfn"} [does not
            contain](#header-list-contains){#ref-for-header-list-contains①⑨
            link-type="dfn"} \``Authorization`\`, then:

            1.  Let `authorizationValue`{.variable} be null.

            2.  If there's an [authentication
                entry](#authentication-entry){#ref-for-authentication-entry③
                link-type="dfn"} for `httpRequest`{.variable} and either
                `httpRequest`{.variable}'s [use-URL-credentials
                flag](#concept-request-use-url-credentials-flag){#ref-for-concept-request-use-url-credentials-flag
                link-type="dfn"} is unset or `httpRequest`{.variable}'s
                [current
                URL](#concept-request-current-url){#ref-for-concept-request-current-url②⑤
                link-type="dfn"} does not [include
                credentials](https://url.spec.whatwg.org/#include-credentials){#ref-for-include-credentials②
                link-type="dfn"}, then set
                `authorizationValue`{.variable} to [authentication
                entry](#authentication-entry){#ref-for-authentication-entry④
                link-type="dfn"}.

            3.  Otherwise, if `httpRequest`{.variable}'s [current
                URL](#concept-request-current-url){#ref-for-concept-request-current-url②⑥
                link-type="dfn"} does [include
                credentials](https://url.spec.whatwg.org/#include-credentials){#ref-for-include-credentials③
                link-type="dfn"} and `isAuthenticationFetch`{.variable}
                is true, set `authorizationValue`{.variable} to
                `httpRequest`{.variable}'s [current
                URL](#concept-request-current-url){#ref-for-concept-request-current-url②⑦
                link-type="dfn"}, [converted to an \``Authorization`\`
                value]{.XXX}.

            4.  If `authorizationValue`{.variable} is non-null, then
                [append](#concept-header-list-append){#ref-for-concept-header-list-append①⑤
                link-type="dfn"} (\``Authorization`\`,
                `authorizationValue`{.variable}) to
                `httpRequest`{.variable}'s [header
                list](#concept-request-header-list){#ref-for-concept-request-header-list③⑤
                link-type="dfn"}.

    22. If there's a [proxy-authentication
        entry](#proxy-authentication-entry){#ref-for-proxy-authentication-entry
        link-type="dfn"}, use it as appropriate.

        This intentionally does not depend on `httpRequest`{.variable}'s
        [credentials
        mode](#concept-request-credentials-mode){#ref-for-concept-request-credentials-mode①②
        link-type="dfn"}.

    23. Set `httpCache`{.variable} to the result of [determining the
        HTTP cache
        partition](#determine-the-http-cache-partition){#ref-for-determine-the-http-cache-partition
        link-type="dfn"}, given `httpRequest`{.variable}.

    24. If `httpCache`{.variable} is null, then set
        `httpRequest`{.variable}'s [cache
        mode](#concept-request-cache-mode){#ref-for-concept-request-cache-mode⑤
        link-type="dfn"} to \"`no-store`\".

    25. If `httpRequest`{.variable}'s [cache
        mode](#concept-request-cache-mode){#ref-for-concept-request-cache-mode⑥
        link-type="dfn"} is neither \"`no-store`\" nor \"`reload`\",
        then:

        1.  Set `storedResponse`{.variable} to the result of selecting a
            response from the `httpCache`{.variable}, possibly needing
            validation, as per the \"[Constructing Responses from
            Caches](https://httpwg.org/specs/rfc9111.html#constructing.responses.from.caches){#ref-for-constructing.responses.from.caches
            link-type="dfn"}\" chapter of HTTP Caching, if any.
            [\[HTTP-CACHING\]](#biblio-http-caching "HTTP Caching"){link-type="biblio"}

            As mandated by HTTP, this still takes the \``Vary`\`
            [header](#concept-header){#ref-for-concept-header⑤③
            link-type="dfn"} into account.

        2.  If `storedResponse`{.variable} is non-null, then:

            1.  If [cache
                mode](#concept-request-cache-mode){#ref-for-concept-request-cache-mode⑦
                link-type="dfn"} is \"`default`\",
                `storedResponse`{.variable} is a [stale-while-revalidate
                response](#concept-stale-while-revalidate-response){#ref-for-concept-stale-while-revalidate-response②
                link-type="dfn"}, and `httpRequest`{.variable}'s
                [client](#concept-request-client){#ref-for-concept-request-client②⑧
                link-type="dfn"} is non-null, then:

                1.  Set `response`{.variable} to
                    `storedResponse`{.variable}.

                2.  Set `response`{.variable}'s [cache
                    state](#concept-response-cache-state){#ref-for-concept-response-cache-state①
                    link-type="dfn"} to \"`local`\".

                3.  Let `revalidateRequest`{.variable} be a
                    [clone](#concept-request-clone){#ref-for-concept-request-clone②
                    link-type="dfn"} of `request`{.variable}.

                4.  Set `revalidateRequest`{.variable}'s [cache
                    mode](#concept-request-cache-mode){#ref-for-concept-request-cache-mode⑧
                    link-type="dfn"} set to \"`no-cache`\".

                5.  Set `revalidateRequest`{.variable}'s [prevent
                    no-cache cache-control header modification
                    flag](#no-cache-prevent-cache-control){#ref-for-no-cache-prevent-cache-control①
                    link-type="dfn"}.

                6.  Set `revalidateRequest`{.variable}'s
                    [service-workers
                    mode](#request-service-workers-mode){#ref-for-request-service-workers-mode②
                    link-type="dfn"} set to \"`none`\".

                7.  [In
                    parallel](https://html.spec.whatwg.org/multipage/infrastructure.html#in-parallel){#ref-for-in-parallel②
                    link-type="dfn"}, run [main
                    fetch](#concept-main-fetch){#ref-for-concept-main-fetch③
                    link-type="dfn"} given a new [fetch
                    params](#fetch-params){#ref-for-fetch-params①⓪
                    link-type="dfn"} whose
                    [request](#fetch-params-request){#ref-for-fetch-params-request①⑧
                    link-type="dfn"} is `revalidateRequest`{.variable}.

                    This fetch is only meant to update the state of
                    `httpCache`{.variable} and the response will be
                    unused until another cache access. The stale
                    response will be used as the response to current
                    request. This fetch is issued in the context of a
                    client so if it goes away the request will be
                    terminated.

            2.  Otherwise:

                1.  If `storedResponse`{.variable} is a [stale
                    response](#concept-stale-response){#ref-for-concept-stale-response①
                    link-type="dfn"}, then set the
                    `revalidatingFlag`{.variable}.

                2.  If the `revalidatingFlag`{.variable} is set and
                    `httpRequest`{.variable}'s [cache
                    mode](#concept-request-cache-mode){#ref-for-concept-request-cache-mode⑨
                    link-type="dfn"} is neither \"`force-cache`\" nor
                    \"`only-if-cached`\", then:

                    1.  If `storedResponse`{.variable}'s [header
                        list](#concept-response-header-list){#ref-for-concept-response-header-list②⓪
                        link-type="dfn"}
                        [contains](#header-list-contains){#ref-for-header-list-contains②⓪
                        link-type="dfn"} \``ETag`\`, then
                        [append](#concept-header-list-append){#ref-for-concept-header-list-append①⑥
                        link-type="dfn"} (\``If-None-Match`\`,
                        \``ETag`\`\'s
                        [value](#concept-header-value){#ref-for-concept-header-value①⑧
                        link-type="dfn"}) to `httpRequest`{.variable}'s
                        [header
                        list](#concept-request-header-list){#ref-for-concept-request-header-list③⑥
                        link-type="dfn"}.

                    2.  If `storedResponse`{.variable}'s [header
                        list](#concept-response-header-list){#ref-for-concept-response-header-list②①
                        link-type="dfn"}
                        [contains](#header-list-contains){#ref-for-header-list-contains②①
                        link-type="dfn"} \``Last-Modified`\`, then
                        [append](#concept-header-list-append){#ref-for-concept-header-list-append①⑦
                        link-type="dfn"} (\``If-Modified-Since`\`,
                        \``Last-Modified`\`\'s
                        [value](#concept-header-value){#ref-for-concept-header-value①⑨
                        link-type="dfn"}) to `httpRequest`{.variable}'s
                        [header
                        list](#concept-request-header-list){#ref-for-concept-request-header-list③⑦
                        link-type="dfn"}.

                    See also the \"[Sending a Validation
                    Request](https://httpwg.org/specs/rfc9111.html#validation.sent){#ref-for-validation.sent
                    link-type="dfn"}\" chapter of HTTP Caching.
                    [\[HTTP-CACHING\]](#biblio-http-caching "HTTP Caching"){link-type="biblio"}

                3.  Otherwise, set `response`{.variable} to
                    `storedResponse`{.variable} and set
                    `response`{.variable}'s [cache
                    state](#concept-response-cache-state){#ref-for-concept-response-cache-state②
                    link-type="dfn"} to \"`local`\".

9.  [If
    aborted](https://infra.spec.whatwg.org/#if-aborted){#ref-for-if-aborted
    link-type="dfn"}, then return the [appropriate network
    error](#appropriate-network-error){#ref-for-appropriate-network-error①
    link-type="dfn"} for `fetchParams`{.variable}.

10. If `response`{.variable} is null, then:

    1.  If `httpRequest`{.variable}'s [cache
        mode](#concept-request-cache-mode){#ref-for-concept-request-cache-mode①⓪
        link-type="dfn"} is \"`only-if-cached`\", then return a [network
        error](#concept-network-error){#ref-for-concept-network-error④⑥
        link-type="dfn"}.

    2.  Let `forwardResponse`{.variable} be the result of running
        [HTTP-network
        fetch](#concept-http-network-fetch){#ref-for-concept-http-network-fetch
        link-type="dfn"} given `httpFetchParams`{.variable},
        `includeCredentials`{.variable}, and
        `isNewConnectionFetch`{.variable}.

    3.  If `httpRequest`{.variable}'s
        [method](#concept-request-method){#ref-for-concept-request-method①④
        link-type="dfn"} is
        [unsafe](https://httpwg.org/specs/rfc9110.html#rfc.section.9.2.1){#ref-for-rfc.section.9.2.1
        link-type="dfn"} and `forwardResponse`{.variable}'s
        [status](#concept-response-status){#ref-for-concept-response-status①⑤
        link-type="dfn"} is in the range 200 to 399, inclusive,
        invalidate appropriate stored responses in
        `httpCache`{.variable}, as per the \"[Invalidating Stored
        Responses](https://httpwg.org/specs/rfc9111.html#invalidation){#ref-for-invalidation
        link-type="dfn"}\" chapter of HTTP Caching, and set
        `storedResponse`{.variable} to null.
        [\[HTTP-CACHING\]](#biblio-http-caching "HTTP Caching"){link-type="biblio"}

    4.  If the `revalidatingFlag`{.variable} is set and
        `forwardResponse`{.variable}'s
        [status](#concept-response-status){#ref-for-concept-response-status①⑥
        link-type="dfn"} is 304, then:

        1.  Update `storedResponse`{.variable}'s [header
            list](#concept-response-header-list){#ref-for-concept-response-header-list②②
            link-type="dfn"} using `forwardResponse`{.variable}'s
            [header
            list](#concept-response-header-list){#ref-for-concept-response-header-list②③
            link-type="dfn"}, as per the \"[Freshening Stored Responses
            upon
            Validation](https://httpwg.org/specs/rfc9111.html#freshening.responses){#ref-for-freshening.responses
            link-type="dfn"}\" chapter of HTTP Caching.
            [\[HTTP-CACHING\]](#biblio-http-caching "HTTP Caching"){link-type="biblio"}

            This updates the stored response in cache as well.

        2.  Set `response`{.variable} to `storedResponse`{.variable}.

        3.  Set `response`{.variable}'s [cache
            state](#concept-response-cache-state){#ref-for-concept-response-cache-state③
            link-type="dfn"} to \"`validated`\".

    5.  If `response`{.variable} is null, then:

        1.  Set `response`{.variable} to `forwardResponse`{.variable}.

        2.  Store `httpRequest`{.variable} and
            `forwardResponse`{.variable} in `httpCache`{.variable}, as
            per the \"[Storing Responses in
            Caches](https://httpwg.org/specs/rfc9111.html#response.cacheability){#ref-for-response.cacheability
            link-type="dfn"}\" chapter of HTTP Caching.
            [\[HTTP-CACHING\]](#biblio-http-caching "HTTP Caching"){link-type="biblio"}

            If `forwardResponse`{.variable} is a [network
            error](#concept-network-error){#ref-for-concept-network-error④⑦
            link-type="dfn"}, this effectively caches the network error,
            which is sometimes known as \"negative caching\".

            The associated [body
            info](#concept-response-body-info){#ref-for-concept-response-body-info④
            link-type="dfn"} is stored in the cache alongside the
            response.

11. Set `response`{.variable}'s [URL
    list](#concept-response-url-list){#ref-for-concept-response-url-list①⓪
    link-type="dfn"} to a
    [clone](https://infra.spec.whatwg.org/#list-clone){#ref-for-list-clone①
    link-type="dfn"} of `httpRequest`{.variable}'s [URL
    list](#concept-request-url-list){#ref-for-concept-request-url-list⑥
    link-type="dfn"}.

12. If `httpRequest`{.variable}'s [header
    list](#concept-request-header-list){#ref-for-concept-request-header-list③⑧
    link-type="dfn"}
    [contains](#header-list-contains){#ref-for-header-list-contains②②
    link-type="dfn"} \``Range`\`, then set `response`{.variable}'s
    [range-requested
    flag](#concept-response-range-requested-flag){#ref-for-concept-response-range-requested-flag②
    link-type="dfn"}.

13. Set `response`{.variable}'s
    [request-includes-credentials](#response-request-includes-credentials){#ref-for-response-request-includes-credentials①
    link-type="dfn"} to `includeCredentials`{.variable}.

14. If `response`{.variable}'s
    [status](#concept-response-status){#ref-for-concept-response-status①⑦
    link-type="dfn"} is 401, `httpRequest`{.variable}'s [response
    tainting](#concept-request-response-tainting){#ref-for-concept-request-response-tainting①⑥
    link-type="dfn"} is not \"`cors`\", `includeCredentials`{.variable}
    is true, and `request`{.variable}'s [traversable for user
    prompts](#concept-request-window){#ref-for-concept-request-window⑦
    link-type="dfn"} is a [traversable
    navigable](https://html.spec.whatwg.org/multipage/document-sequences.html#traversable-navigable){#ref-for-traversable-navigable④
    link-type="dfn"}:

    1.  Needs testing: multiple \``WWW-Authenticate`\` headers, missing,
        parsing issues.

    2.  If `request`{.variable}'s
        [body](#concept-request-body){#ref-for-concept-request-body②④
        link-type="dfn"} is non-null, then:

        1.  If `request`{.variable}'s
            [body](#concept-request-body){#ref-for-concept-request-body②⑤
            link-type="dfn"}'s
            [source](#concept-body-source){#ref-for-concept-body-source⑥
            link-type="dfn"} is null, then return a [network
            error](#concept-network-error){#ref-for-concept-network-error④⑧
            link-type="dfn"}.

        2.  Set `request`{.variable}'s
            [body](#concept-request-body){#ref-for-concept-request-body②⑥
            link-type="dfn"} to the
            [body](#body-with-type-body){#ref-for-body-with-type-body④
            link-type="dfn"} of the result of [safely
            extracting](#bodyinit-safely-extract){#ref-for-bodyinit-safely-extract⑤
            link-type="dfn"} `request`{.variable}'s
            [body](#concept-request-body){#ref-for-concept-request-body②⑦
            link-type="dfn"}'s
            [source](#concept-body-source){#ref-for-concept-body-source⑦
            link-type="dfn"}.

    3.  If `request`{.variable}'s [use-URL-credentials
        flag](#concept-request-use-url-credentials-flag){#ref-for-concept-request-use-url-credentials-flag①
        link-type="dfn"} is unset or `isAuthenticationFetch`{.variable}
        is true, then:

        1.  If `fetchParams`{.variable} is
            [canceled](#fetch-params-canceled){#ref-for-fetch-params-canceled④
            link-type="dfn"}, then return the [appropriate network
            error](#appropriate-network-error){#ref-for-appropriate-network-error②
            link-type="dfn"} for `fetchParams`{.variable}.

        2.  Let `username`{.variable} and `password`{.variable} be the
            result of prompting the end user for a username and
            password, respectively, in `request`{.variable}'s
            [traversable for user
            prompts](#concept-request-window){#ref-for-concept-request-window⑧
            link-type="dfn"}.

        3.  [Set the
            username](https://url.spec.whatwg.org/#set-the-username){#ref-for-set-the-username①
            link-type="dfn"} given `request`{.variable}'s [current
            URL](#concept-request-current-url){#ref-for-concept-request-current-url②⑧
            link-type="dfn"} and `username`{.variable}.

        4.  [Set the
            password](https://url.spec.whatwg.org/#set-the-password){#ref-for-set-the-password①
            link-type="dfn"} given `request`{.variable}'s [current
            URL](#concept-request-current-url){#ref-for-concept-request-current-url②⑨
            link-type="dfn"} and `password`{.variable}.

    4.  Set `response`{.variable} to the result of running
        [HTTP-network-or-cache
        fetch](#concept-http-network-or-cache-fetch){#ref-for-concept-http-network-or-cache-fetch①
        link-type="dfn"} given `fetchParams`{.variable} and true.

15. If `response`{.variable}'s
    [status](#concept-response-status){#ref-for-concept-response-status①⑧
    link-type="dfn"} is 407, then:

    1.  If `request`{.variable}'s [traversable for user
        prompts](#concept-request-window){#ref-for-concept-request-window⑨
        link-type="dfn"} is \"`no-traversable`\", then return a [network
        error](#concept-network-error){#ref-for-concept-network-error④⑨
        link-type="dfn"}.

    2.  Needs testing: multiple \``Proxy-Authenticate`\` headers,
        missing, parsing issues.

    3.  If `fetchParams`{.variable} is
        [canceled](#fetch-params-canceled){#ref-for-fetch-params-canceled⑤
        link-type="dfn"}, then return the [appropriate network
        error](#appropriate-network-error){#ref-for-appropriate-network-error③
        link-type="dfn"} for `fetchParams`{.variable}.

    4.  Prompt the end user as appropriate in `request`{.variable}'s
        [traversable for user
        prompts](#concept-request-window){#ref-for-concept-request-window①⓪
        link-type="dfn"} and store the result as a [proxy-authentication
        entry](#proxy-authentication-entry){#ref-for-proxy-authentication-entry①
        link-type="dfn"}.
        [\[HTTP\]](#biblio-http "HTTP Semantics"){link-type="biblio"}

        Remaining details surrounding proxy authentication are defined
        by HTTP.

    5.  Set `response`{.variable} to the result of running
        [HTTP-network-or-cache
        fetch](#concept-http-network-or-cache-fetch){#ref-for-concept-http-network-or-cache-fetch②
        link-type="dfn"} given `fetchParams`{.variable}.

16. If all of the following are true

    - `response`{.variable}'s
      [status](#concept-response-status){#ref-for-concept-response-status①⑨
      link-type="dfn"} is 421

    - `isNewConnectionFetch`{.variable} is false

    - `request`{.variable}'s
      [body](#concept-request-body){#ref-for-concept-request-body②⑧
      link-type="dfn"} is null, or `request`{.variable}'s
      [body](#concept-request-body){#ref-for-concept-request-body②⑨
      link-type="dfn"} is non-null and `request`{.variable}'s
      [body](#concept-request-body){#ref-for-concept-request-body③⓪
      link-type="dfn"}'s
      [source](#concept-body-source){#ref-for-concept-body-source⑧
      link-type="dfn"} is non-null

    then:

    1.  If `fetchParams`{.variable} is
        [canceled](#fetch-params-canceled){#ref-for-fetch-params-canceled⑥
        link-type="dfn"}, then return the [appropriate network
        error](#appropriate-network-error){#ref-for-appropriate-network-error④
        link-type="dfn"} for `fetchParams`{.variable}.

    2.  Set `response`{.variable} to the result of running
        [HTTP-network-or-cache
        fetch](#concept-http-network-or-cache-fetch){#ref-for-concept-http-network-or-cache-fetch③
        link-type="dfn"} given `fetchParams`{.variable},
        `isAuthenticationFetch`{.variable}, and true.

17. If `isAuthenticationFetch`{.variable} is true, then create an
    [authentication
    entry](#authentication-entry){#ref-for-authentication-entry⑤
    link-type="dfn"} for `request`{.variable} and the given realm.

18. Return `response`{.variable}. [Typically `response`{.variable}'s
    [body](#concept-response-body){#ref-for-concept-response-body②③
    link-type="dfn"}'s
    [stream](#concept-body-stream){#ref-for-concept-body-stream①①
    link-type="dfn"} is still being enqueued to after returning.]{.note}
:::

### [4.6. ]{.secno}[HTTP-network fetch]{.content}[](#http-network-fetch){.self-link} {#http-network-fetch .heading .settled level="4.6"}

::: {.algorithm algorithm="HTTP-network fetch"}
To [HTTP-network fetch]{#concept-http-network-fetch .dfn .dfn-paneled
dfn-type="dfn" noexport=""}, given a [fetch
params](#fetch-params){#ref-for-fetch-params①① link-type="dfn"}
`fetchParams`{.variable}, an optional boolean
`includeCredentials`{.variable} (default false), and an optional boolean
`forceNewConnection`{.variable} (default false), run these steps:

1.  Let `request`{.variable} be `fetchParams`{.variable}'s
    [request](#fetch-params-request){#ref-for-fetch-params-request①⑨
    link-type="dfn"}.

2.  Let `response`{.variable} be null.

3.  Let `timingInfo`{.variable} be `fetchParams`{.variable}'s [timing
    info](#fetch-params-timing-info){#ref-for-fetch-params-timing-info⑤
    link-type="dfn"}.

4.  Let `networkPartitionKey`{.variable} be the result of [determining
    the network partition
    key](#request-determine-the-network-partition-key){#ref-for-request-determine-the-network-partition-key①
    link-type="dfn"} given `request`{.variable}.

5.  Let `newConnection`{.variable} be \"`yes`\" if
    `forceNewConnection`{.variable} is true; otherwise \"`no`\".

6.  Switch on `request`{.variable}'s
    [mode](#concept-request-mode){#ref-for-concept-request-mode②⓪
    link-type="dfn"}:

    \"`websocket`\"

    :   Let `connection`{.variable} be the result of [obtaining a
        WebSocket
        connection](https://websockets.spec.whatwg.org/#concept-websocket-connection-obtain){#ref-for-concept-websocket-connection-obtain
        link-type="dfn"}, given `request`{.variable}'s [current
        URL](#concept-request-current-url){#ref-for-concept-request-current-url③⓪
        link-type="dfn"}.

    Otherwise

    :   Let `connection`{.variable} be the result of [obtaining a
        connection](#concept-connection-obtain){#ref-for-concept-connection-obtain①
        link-type="dfn"}, given `networkPartitionKey`{.variable},
        `request`{.variable}'s [current
        URL](#concept-request-current-url){#ref-for-concept-request-current-url③①
        link-type="dfn"}, `includeCredentials`{.variable}, and
        `newConnection`{.variable}.

7.  Run these steps, but [abort
    when](https://infra.spec.whatwg.org/#abort-when){#ref-for-abort-when①
    link-type="dfn"} `fetchParams`{.variable} is
    [canceled](#fetch-params-canceled){#ref-for-fetch-params-canceled⑦
    link-type="dfn"}:

    1.  If `connection`{.variable} is failure, then return a [network
        error](#concept-network-error){#ref-for-concept-network-error⑤⓪
        link-type="dfn"}.

    2.  Set `timingInfo`{.variable}'s [final connection timing
        info](#fetch-timing-info-final-connection-timing-info){#ref-for-fetch-timing-info-final-connection-timing-info
        link-type="dfn"} to the result of calling [clamp and coarsen
        connection timing
        info](#clamp-and-coarsen-connection-timing-info){#ref-for-clamp-and-coarsen-connection-timing-info①
        link-type="dfn"} with `connection`{.variable}'s [timing
        info](#concept-connection-timing-info){#ref-for-concept-connection-timing-info②
        link-type="dfn"}, `timingInfo`{.variable}'s [post-redirect start
        time](#fetch-timing-info-post-redirect-start-time){#ref-for-fetch-timing-info-post-redirect-start-time③
        link-type="dfn"}, and `fetchParams`{.variable}'s [cross-origin
        isolated
        capability](#fetch-params-cross-origin-isolated-capability){#ref-for-fetch-params-cross-origin-isolated-capability④
        link-type="dfn"}.

    3.  If `connection`{.variable} is an HTTP/1.x connection,
        `request`{.variable}'s
        [body](#concept-request-body){#ref-for-concept-request-body③①
        link-type="dfn"} is non-null, and `request`{.variable}'s
        [body](#concept-request-body){#ref-for-concept-request-body③②
        link-type="dfn"}'s
        [source](#concept-body-source){#ref-for-concept-body-source⑨
        link-type="dfn"} is null, then return a [network
        error](#concept-network-error){#ref-for-concept-network-error⑤①
        link-type="dfn"}.

    4.  Set `timingInfo`{.variable}'s [final network-request start
        time](#fetch-timing-info-final-network-request-start-time){#ref-for-fetch-timing-info-final-network-request-start-time
        link-type="dfn"} to the [coarsened shared current
        time](https://w3c.github.io/hr-time/#dfn-coarsened-shared-current-time){#ref-for-dfn-coarsened-shared-current-time③
        link-type="dfn"} given `fetchParams`{.variable}'s [cross-origin
        isolated
        capability](#fetch-params-cross-origin-isolated-capability){#ref-for-fetch-params-cross-origin-isolated-capability⑤
        link-type="dfn"}.

    5.  Set `response`{.variable} to the result of making an HTTP
        request over `connection`{.variable} using `request`{.variable}
        with the following caveats:

        - Follow the relevant requirements from HTTP.
          [\[HTTP\]](#biblio-http "HTTP Semantics"){link-type="biblio"}
          [\[HTTP-CACHING\]](#biblio-http-caching "HTTP Caching"){link-type="biblio"}

        - If `request`{.variable}'s
          [body](#concept-request-body){#ref-for-concept-request-body③③
          link-type="dfn"} is non-null, and `request`{.variable}'s
          [body](#concept-request-body){#ref-for-concept-request-body③④
          link-type="dfn"}'s
          [source](#concept-body-source){#ref-for-concept-body-source①⓪
          link-type="dfn"} is null, then the user agent may have a
          buffer of up to 64 kibibytes and store a part of
          `request`{.variable}'s
          [body](#concept-request-body){#ref-for-concept-request-body③⑤
          link-type="dfn"} in that buffer. If the user agent reads from
          `request`{.variable}'s
          [body](#concept-request-body){#ref-for-concept-request-body③⑥
          link-type="dfn"} beyond that buffer's size and the user agent
          needs to resend `request`{.variable}, then instead return a
          [network
          error](#concept-network-error){#ref-for-concept-network-error⑤②
          link-type="dfn"}.

          ::: {.note role="note"}
          The resending is needed when the connection is timed out, for
          example.

          The buffer is not needed when `request`{.variable}'s
          [body](#concept-request-body){#ref-for-concept-request-body③⑦
          link-type="dfn"}'s
          [source](#concept-body-source){#ref-for-concept-body-source①①
          link-type="dfn"} is non-null, because `request`{.variable}'s
          [body](#concept-request-body){#ref-for-concept-request-body③⑧
          link-type="dfn"} can be recreated from it.

          When `request`{.variable}'s
          [body](#concept-request-body){#ref-for-concept-request-body③⑨
          link-type="dfn"}'s
          [source](#concept-body-source){#ref-for-concept-body-source①②
          link-type="dfn"} is null, it means
          [body](#concept-request-body){#ref-for-concept-request-body④⓪
          link-type="dfn"} is created from a
          [`ReadableStream`{.idl}](https://streams.spec.whatwg.org/#readablestream){#ref-for-readablestream②
          link-type="idl"} object, which means
          [body](#concept-request-body){#ref-for-concept-request-body④①
          link-type="dfn"} cannot be recreated and that is why the
          buffer is needed.
          :::

        - While true:

          1.  Set `timingInfo`{.variable}'s [final network-response
              start
              time](#fetch-timing-info-final-network-response-start-time){#ref-for-fetch-timing-info-final-network-response-start-time
              link-type="dfn"} to the [coarsened shared current
              time](https://w3c.github.io/hr-time/#dfn-coarsened-shared-current-time){#ref-for-dfn-coarsened-shared-current-time④
              link-type="dfn"} given `fetchParams`{.variable}'s
              [cross-origin isolated
              capability](#fetch-params-cross-origin-isolated-capability){#ref-for-fetch-params-cross-origin-isolated-capability⑥
              link-type="dfn"}, immediately after the user agent's HTTP
              parser receives the first byte of the response (e.g.,
              frame header bytes for HTTP/2 or response status line for
              HTTP/1.x).

          2.  Wait until all the HTTP response headers are transmitted.

          3.  Let `status`{.variable} be the HTTP response's status
              code.

          4.  If `status`{.variable} is in the range 100 to 199,
              inclusive:

              1.  If `timingInfo`{.variable}'s [first interim
                  network-response start
                  time](#fetch-timing-info-first-interim-network-response-start-time){#ref-for-fetch-timing-info-first-interim-network-response-start-time
                  link-type="dfn"} is 0, then set
                  `timingInfo`{.variable}'s [first interim
                  network-response start
                  time](#fetch-timing-info-first-interim-network-response-start-time){#ref-for-fetch-timing-info-first-interim-network-response-start-time①
                  link-type="dfn"} to `timingInfo`{.variable}'s [final
                  network-response start
                  time](#fetch-timing-info-final-network-response-start-time){#ref-for-fetch-timing-info-final-network-response-start-time①
                  link-type="dfn"}.

              2.  If `request`{.variable}'s
                  [mode](#concept-request-mode){#ref-for-concept-request-mode②①
                  link-type="dfn"} is \"`websocket`\" and
                  `status`{.variable} is 101, then
                  [break](https://infra.spec.whatwg.org/#iteration-break){#ref-for-iteration-break③
                  link-type="dfn"}.

              3.  If `status`{.variable} is 103 and
                  `fetchParams`{.variable}'s [process early hints
                  response](#fetch-params-process-early-hints-response){#ref-for-fetch-params-process-early-hints-response①
                  link-type="dfn"} is non-null, then [queue a fetch
                  task](#queue-a-fetch-task){#ref-for-queue-a-fetch-task⑧
                  link-type="dfn"} to run `fetchParams`{.variable}'s
                  [process early hints
                  response](#fetch-params-process-early-hints-response){#ref-for-fetch-params-process-early-hints-response②
                  link-type="dfn"}, with
                  [response](#concept-response){#ref-for-concept-response⑤⑨
                  link-type="dfn"}.

              4.  [Continue](https://infra.spec.whatwg.org/#iteration-continue){#ref-for-iteration-continue⑤
                  link-type="dfn"}.

              These kind of HTTP responses are eventually followed by a
              \"final\" HTTP response.

          5.  [Break](https://infra.spec.whatwg.org/#iteration-break){#ref-for-iteration-break④
              link-type="dfn"}.

        The exact layering between Fetch and HTTP still needs to be
        sorted through and therefore `response`{.variable} represents
        both a [response](#concept-response){#ref-for-concept-response⑥⓪
        link-type="dfn"} and an HTTP response here.

        If the HTTP request results in a TLS client certificate dialog,
        then:

        1.  If `request`{.variable}'s [traversable for user
            prompts](#concept-request-window){#ref-for-concept-request-window①①
            link-type="dfn"} is a [traversable
            navigable](https://html.spec.whatwg.org/multipage/document-sequences.html#traversable-navigable){#ref-for-traversable-navigable⑤
            link-type="dfn"}, then make the dialog available in
            `request`{.variable}'s [traversable for user
            prompts](#concept-request-window){#ref-for-concept-request-window①②
            link-type="dfn"}.

        2.  Otherwise, return a [network
            error](#concept-network-error){#ref-for-concept-network-error⑤③
            link-type="dfn"}.

        To transmit `request`{.variable}'s
        [body](#concept-request-body){#ref-for-concept-request-body④②
        link-type="dfn"} `body`{.variable}, run these steps:

        1.  If `body`{.variable} is null and `fetchParams`{.variable}'s
            [process request
            end-of-body](#fetch-params-process-request-end-of-body){#ref-for-fetch-params-process-request-end-of-body①
            link-type="dfn"} is non-null, then [queue a fetch
            task](#queue-a-fetch-task){#ref-for-queue-a-fetch-task⑨
            link-type="dfn"} given `fetchParams`{.variable}'s [process
            request
            end-of-body](#fetch-params-process-request-end-of-body){#ref-for-fetch-params-process-request-end-of-body②
            link-type="dfn"} and `fetchParams`{.variable}'s [task
            destination](#fetch-params-task-destination){#ref-for-fetch-params-task-destination⑥
            link-type="dfn"}.

        2.  Otherwise, if `body`{.variable} is non-null:

            1.  Let `processBodyChunk`{.variable} given
                `bytes`{.variable} be these steps:

                1.  If `fetchParams`{.variable} is
                    [canceled](#fetch-params-canceled){#ref-for-fetch-params-canceled⑧
                    link-type="dfn"}, then abort these steps.

                2.  Run this step [in
                    parallel](https://html.spec.whatwg.org/multipage/infrastructure.html#in-parallel){#ref-for-in-parallel③
                    link-type="dfn"}: transmit `bytes`{.variable}.

                3.  If `fetchParams`{.variable}'s [process request body
                    chunk
                    length](#fetch-params-process-request-body){#ref-for-fetch-params-process-request-body①
                    link-type="dfn"} is non-null, then run
                    `fetchParams`{.variable}'s [process request body
                    chunk
                    length](#fetch-params-process-request-body){#ref-for-fetch-params-process-request-body②
                    link-type="dfn"} given `bytes`{.variable}'s
                    [length](https://infra.spec.whatwg.org/#byte-sequence-length){#ref-for-byte-sequence-length②
                    link-type="dfn"}.

            2.  Let `processEndOfBody`{.variable} be these steps:

                1.  If `fetchParams`{.variable} is
                    [canceled](#fetch-params-canceled){#ref-for-fetch-params-canceled⑨
                    link-type="dfn"}, then abort these steps.

                2.  If `fetchParams`{.variable}'s [process request
                    end-of-body](#fetch-params-process-request-end-of-body){#ref-for-fetch-params-process-request-end-of-body③
                    link-type="dfn"} is non-null, then run
                    `fetchParams`{.variable}'s [process request
                    end-of-body](#fetch-params-process-request-end-of-body){#ref-for-fetch-params-process-request-end-of-body④
                    link-type="dfn"}.

            3.  Let `processBodyError`{.variable} given `e`{.variable}
                be these steps:

                1.  If `fetchParams`{.variable} is
                    [canceled](#fetch-params-canceled){#ref-for-fetch-params-canceled①⓪
                    link-type="dfn"}, then abort these steps.

                2.  If `e`{.variable} is an
                    \"[`AbortError`](https://webidl.spec.whatwg.org/#aborterror){#ref-for-aborterror②
                    .idl-code link-type="exception"}\"
                    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException②
                    link-type="idl"}, then
                    [abort](#fetch-controller-abort){#ref-for-fetch-controller-abort
                    link-type="dfn"} `fetchParams`{.variable}'s
                    [controller](#fetch-params-controller){#ref-for-fetch-params-controller①⓪
                    link-type="dfn"}.

                3.  Otherwise,
                    [terminate](#fetch-controller-terminate){#ref-for-fetch-controller-terminate②
                    link-type="dfn"} `fetchParams`{.variable}'s
                    [controller](#fetch-params-controller){#ref-for-fetch-params-controller①①
                    link-type="dfn"}.

            4.  [Incrementally
                read](#body-incrementally-read){#ref-for-body-incrementally-read
                link-type="dfn"} `request`{.variable}'s
                [body](#concept-request-body){#ref-for-concept-request-body④③
                link-type="dfn"} given `processBodyChunk`{.variable},
                `processEndOfBody`{.variable},
                `processBodyError`{.variable}, and
                `fetchParams`{.variable}'s [task
                destination](#fetch-params-task-destination){#ref-for-fetch-params-task-destination⑦
                link-type="dfn"}.

8.  [If
    aborted](https://infra.spec.whatwg.org/#if-aborted){#ref-for-if-aborted①
    link-type="dfn"}, then:

    1.  If `connection`{.variable} uses HTTP/2, then transmit an
        `RST_STREAM` frame.

    2.  Return the [appropriate network
        error](#appropriate-network-error){#ref-for-appropriate-network-error⑤
        link-type="dfn"} for `fetchParams`{.variable}.

9.  Let `buffer`{.variable} be an empty [byte
    sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence①⑦
    link-type="dfn"}.

    This represents an internal buffer inside the network layer of the
    user agent.

10. Let `stream`{.variable} be a
    [new](https://webidl.spec.whatwg.org/#new){#ref-for-new
    link-type="dfn"}
    [`ReadableStream`{.idl}](https://streams.spec.whatwg.org/#readablestream){#ref-for-readablestream③
    link-type="idl"}.

11. Let `pullAlgorithm`{.variable} be the following steps:

    1.  Let `promise`{.variable} be [a new
        promise](https://webidl.spec.whatwg.org/#a-new-promise){#ref-for-a-new-promise
        link-type="dfn"}.

    2.  Run the following steps [in
        parallel](https://html.spec.whatwg.org/multipage/infrastructure.html#in-parallel){#ref-for-in-parallel④
        link-type="dfn"}:

        1.  If the size of `buffer`{.variable} is smaller than a lower
            limit chosen by the user agent and the ongoing fetch is
            [suspended](#concept-fetch-suspend){#ref-for-concept-fetch-suspend
            link-type="dfn"},
            [resume](#concept-fetch-resume){#ref-for-concept-fetch-resume
            link-type="dfn"} the fetch.

        2.  Wait until `buffer`{.variable} is not empty.

        3.  [Queue a fetch
            task](#queue-a-fetch-task){#ref-for-queue-a-fetch-task①⓪
            link-type="dfn"} to run the following steps, with
            `fetchParams`{.variable}'s [task
            destination](#fetch-params-task-destination){#ref-for-fetch-params-task-destination⑧
            link-type="dfn"}.

            1.  [Pull from
                bytes](https://streams.spec.whatwg.org/#readablestream-pull-from-bytes){#ref-for-readablestream-pull-from-bytes
                link-type="dfn"} `buffer`{.variable} into
                `stream`{.variable}.

            2.  If `stream`{.variable} is
                [errored](https://streams.spec.whatwg.org/#readablestream-errored){#ref-for-readablestream-errored
                link-type="dfn"}, then
                [terminate](#fetch-controller-terminate){#ref-for-fetch-controller-terminate③
                link-type="dfn"} `fetchParams`{.variable}'s
                [controller](#fetch-params-controller){#ref-for-fetch-params-controller①②
                link-type="dfn"}.

            3.  [Resolve](https://webidl.spec.whatwg.org/#resolve){#ref-for-resolve
                link-type="dfn"} `promise`{.variable} with undefined.

    3.  Return `promise`{.variable}.

12. Let `cancelAlgorithm`{.variable} be an algorithm that
    [aborts](#fetch-controller-abort){#ref-for-fetch-controller-abort①
    link-type="dfn"} `fetchParams`{.variable}'s
    [controller](#fetch-params-controller){#ref-for-fetch-params-controller①③
    link-type="dfn"} with `reason`{.variable}, given
    `reason`{.variable}.

13. [Set
    up](https://streams.spec.whatwg.org/#readablestream-set-up-with-byte-reading-support){#ref-for-readablestream-set-up-with-byte-reading-support
    link-type="dfn"} `stream`{.variable} with byte reading support with
    [`pullAlgorithm`{.variable}](https://streams.spec.whatwg.org/#readablestream-set-up-pullalgorithm){#ref-for-readablestream-set-up-pullalgorithm
    link-type="dfn"} set to `pullAlgorithm`{.variable},
    [`cancelAlgorithm`{.variable}](https://streams.spec.whatwg.org/#readablestream-set-up-cancelalgorithm){#ref-for-readablestream-set-up-cancelalgorithm
    link-type="dfn"} set to `cancelAlgorithm`{.variable}.

14. Set `response`{.variable}'s
    [body](#concept-response-body){#ref-for-concept-response-body②④
    link-type="dfn"} to a new
    [body](#concept-body){#ref-for-concept-body①① link-type="dfn"} whose
    [stream](#concept-body-stream){#ref-for-concept-body-stream①②
    link-type="dfn"} is `stream`{.variable}.

15. [![(This is a tracking
    vector.)](https://resources.whatwg.org/tracking-vector.svg "There is a tracking vector here."){.darkmode-aware
    crossorigin="" height="64"
    width="46"}](https://infra.spec.whatwg.org/#tracking-vector){.tracking-vector
    style="color: currentcolor"} If `includeCredentials`{.variable} is
    true and the user agent is not configured to block cookies for
    `request`{.variable} (see [section
    7](https://httpwg.org/specs/rfc6265.html#privacy-considerations) of
    [\[COOKIES\]](#biblio-cookies "HTTP State Management Mechanism"){link-type="biblio"}),
    then run the \"set-cookie-string\" parsing algorithm (see [section
    5.2](https://httpwg.org/specs/rfc6265.html#set-cookie) of
    [\[COOKIES\]](#biblio-cookies "HTTP State Management Mechanism"){link-type="biblio"})
    on the
    [value](#concept-header-value){#ref-for-concept-header-value②⓪
    link-type="dfn"} of each
    [header](#concept-header){#ref-for-concept-header⑤④ link-type="dfn"}
    whose [name](#concept-header-name){#ref-for-concept-header-name②①
    link-type="dfn"} is a
    [byte-case-insensitive](https://infra.spec.whatwg.org/#byte-case-insensitive){#ref-for-byte-case-insensitive①④
    link-type="dfn"} match for \``Set-Cookie`\` in
    `response`{.variable}'s [header
    list](#concept-response-header-list){#ref-for-concept-response-header-list②④
    link-type="dfn"}, if any, and `request`{.variable}'s [current
    URL](#concept-request-current-url){#ref-for-concept-request-current-url③②
    link-type="dfn"}.

16. Run these steps [in
    parallel](https://html.spec.whatwg.org/multipage/infrastructure.html#in-parallel){#ref-for-in-parallel⑤
    link-type="dfn"}:

    1.  Run these steps, but [abort
        when](https://infra.spec.whatwg.org/#abort-when){#ref-for-abort-when②
        link-type="dfn"} `fetchParams`{.variable} is
        [canceled](#fetch-params-canceled){#ref-for-fetch-params-canceled①①
        link-type="dfn"}:

        1.  While true:

            1.  If one or more bytes have been transmitted from
                `response`{.variable}'s message body, then:

                1.  Let `bytes`{.variable} be the transmitted bytes.

                2.  Let `codings`{.variable} be the result of
                    [extracting header list
                    values](#extract-header-list-values){#ref-for-extract-header-list-values②
                    link-type="dfn"} given \``Content-Encoding`\` and
                    `response`{.variable}'s [header
                    list](#concept-response-header-list){#ref-for-concept-response-header-list②⑤
                    link-type="dfn"}.

                3.  Let `filteredCoding`{.variable} be \"`@unknown`\".

                4.  If `codings`{.variable} is null or failure, then set
                    `filteredCoding`{.variable} to the empty string.

                5.  Otherwise, if `codings`{.variable}'s
                    [size](https://infra.spec.whatwg.org/#list-size){#ref-for-list-size
                    link-type="dfn"} is greater than 1, then set
                    `filteredCoding`{.variable} to \"`multiple`\".

                6.  Otherwise, if `codings`{.variable}\[0\] is the empty
                    string, or it is supported by the user agent, and is
                    a
                    [byte-case-insensitive](https://infra.spec.whatwg.org/#byte-case-insensitive){#ref-for-byte-case-insensitive①⑤
                    link-type="dfn"} match for an entry listed in the
                    HTTP Content Coding Registry, then set
                    `filteredCoding`{.variable} to the result of
                    [byte-lowercasing](https://infra.spec.whatwg.org/#byte-lowercase){#ref-for-byte-lowercase③
                    link-type="dfn"} `codings`{.variable}\[0\].
                    [\[IANA-HTTP-PARAMS\]](#biblio-iana-http-params "Hypertext Transfer Protocol (HTTP) Parameters"){link-type="biblio"}

                7.  Set `response`{.variable}'s [body
                    info](#concept-response-body-info){#ref-for-concept-response-body-info⑤
                    link-type="dfn"}'s [content
                    encoding](#response-body-info-content-encoding){#ref-for-response-body-info-content-encoding
                    link-type="dfn"} to `filteredCoding`{.variable}.

                8.  Increase `response`{.variable}'s [body
                    info](#concept-response-body-info){#ref-for-concept-response-body-info⑥
                    link-type="dfn"}'s [encoded
                    size](#fetch-timing-info-encoded-body-size){#ref-for-fetch-timing-info-encoded-body-size
                    link-type="dfn"} by `bytes`{.variable}'s
                    [length](https://infra.spec.whatwg.org/#byte-sequence-length){#ref-for-byte-sequence-length③
                    link-type="dfn"}.

                9.  Set `bytes`{.variable} to the result of [handling
                    content
                    codings](#handle-content-codings){#ref-for-handle-content-codings①
                    link-type="dfn"} given `codings`{.variable} and
                    `bytes`{.variable}.

                    This makes the \``Content-Length`\`
                    [header](#concept-header){#ref-for-concept-header⑤⑤
                    link-type="dfn"} unreliable to the extent that it
                    was reliable to begin with.

                10. Increase `response`{.variable}'s [body
                    info](#concept-response-body-info){#ref-for-concept-response-body-info⑦
                    link-type="dfn"}'s [decoded
                    size](#fetch-timing-info-decoded-body-size){#ref-for-fetch-timing-info-decoded-body-size
                    link-type="dfn"} by `bytes`{.variable}'s
                    [length](https://infra.spec.whatwg.org/#byte-sequence-length){#ref-for-byte-sequence-length④
                    link-type="dfn"}.

                11. If `bytes`{.variable} is failure, then
                    [terminate](#fetch-controller-terminate){#ref-for-fetch-controller-terminate④
                    link-type="dfn"} `fetchParams`{.variable}'s
                    [controller](#fetch-params-controller){#ref-for-fetch-params-controller①④
                    link-type="dfn"}.

                12. Append `bytes`{.variable} to `buffer`{.variable}.

                13. If the size of `buffer`{.variable} is larger than an
                    upper limit chosen by the user agent, ask the user
                    agent to
                    [suspend](#concept-fetch-suspend){#ref-for-concept-fetch-suspend①
                    link-type="dfn"} the ongoing fetch.

            2.  Otherwise, if the bytes transmission for
                `response`{.variable}'s message body is done normally
                and `stream`{.variable} is
                [readable](https://streams.spec.whatwg.org/#readablestream-readable){#ref-for-readablestream-readable
                link-type="dfn"}, then
                [close](https://streams.spec.whatwg.org/#readablestream-close){#ref-for-readablestream-close
                link-type="dfn"} `stream`{.variable}, and abort these
                in-parallel steps.

    2.  [If
        aborted](https://infra.spec.whatwg.org/#if-aborted){#ref-for-if-aborted②
        link-type="dfn"}, then:

        1.  If `fetchParams`{.variable} is
            [aborted](#fetch-params-aborted){#ref-for-fetch-params-aborted①
            link-type="dfn"}, then:

            1.  Set `response`{.variable}'s [aborted
                flag](#concept-response-aborted){#ref-for-concept-response-aborted①
                link-type="dfn"}.

            2.  If `stream`{.variable} is
                [readable](https://streams.spec.whatwg.org/#readablestream-readable){#ref-for-readablestream-readable①
                link-type="dfn"}, then
                [error](https://streams.spec.whatwg.org/#readablestream-error){#ref-for-readablestream-error
                link-type="dfn"} `stream`{.variable} with the result of
                [deserialize a serialized abort
                reason](#deserialize-a-serialized-abort-reason){#ref-for-deserialize-a-serialized-abort-reason
                link-type="dfn"} given `fetchParams`{.variable}'s
                [controller](#fetch-params-controller){#ref-for-fetch-params-controller①⑤
                link-type="dfn"}'s [serialized abort
                reason](#fetch-controller-serialized-abort-reason){#ref-for-fetch-controller-serialized-abort-reason①
                link-type="dfn"} and an
                [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#ref-for-implementation-defined①⑧
                link-type="dfn"}
                [realm](https://tc39.es/ecma262/#realm){#ref-for-realm①
                link-type="dfn"}.

        2.  Otherwise, if `stream`{.variable} is
            [readable](https://streams.spec.whatwg.org/#readablestream-readable){#ref-for-readablestream-readable②
            link-type="dfn"},
            [error](https://streams.spec.whatwg.org/#readablestream-error){#ref-for-readablestream-error①
            link-type="dfn"} `stream`{.variable} with a
            [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror①
            link-type="idl"}.

        3.  If `connection`{.variable} uses HTTP/2, then transmit an
            `RST_STREAM` frame.

        4.  Otherwise, the user agent should close
            `connection`{.variable} unless it would be bad for
            performance to do so.

            For instance, the user agent could keep the connection open
            if it knows there's only a few bytes of transfer remaining
            on a reusable connection. In this case it could be worse to
            close the connection and go through the handshake process
            again for the next fetch.

    These are run [in
    parallel](https://html.spec.whatwg.org/multipage/infrastructure.html#in-parallel){#ref-for-in-parallel⑥
    link-type="dfn"} as at this point it is unclear whether
    `response`{.variable}'s
    [body](#concept-response-body){#ref-for-concept-response-body②⑤
    link-type="dfn"} is relevant (`response`{.variable} might be a
    redirect).

17. Return `response`{.variable}. [Typically `response`{.variable}'s
    [body](#concept-response-body){#ref-for-concept-response-body②⑥
    link-type="dfn"}'s
    [stream](#concept-body-stream){#ref-for-concept-body-stream①③
    link-type="dfn"} is still being enqueued to after returning.]{.note}
:::

### [4.7. ]{.secno}[CORS-preflight fetch]{.content}[](#cors-preflight-fetch){.self-link} {#cors-preflight-fetch .heading .settled level="4.7"}

This is effectively the user agent implementation of the check to see if
the [CORS protocol](#cors-protocol){#ref-for-cors-protocol①⑦
link-type="dfn"} is understood. The so-called [CORS-preflight
request](#cors-preflight-request){#ref-for-cors-preflight-request①③
link-type="dfn"}. If successful it populates the [CORS-preflight
cache](#concept-cache){#ref-for-concept-cache① link-type="dfn"} to
minimize the number of these
[fetches](#cors-preflight-fetch-0){#ref-for-cors-preflight-fetch-0⑤
link-type="dfn"}.

::: {.algorithm algorithm="CORS-preflight fetch"}
To [CORS-preflight fetch]{#cors-preflight-fetch-0 .dfn .dfn-paneled
dfn-type="dfn" noexport=""}, given a
[request](#concept-request){#ref-for-concept-request①⓪⑦ link-type="dfn"}
`request`{.variable}, run these steps:

1.  Let `preflight`{.variable} be a new
    [request](#concept-request){#ref-for-concept-request①⓪⑧
    link-type="dfn"} whose
    [method](#concept-request-method){#ref-for-concept-request-method①⑤
    link-type="dfn"} is \``OPTIONS`\`, [URL
    list](#concept-request-url-list){#ref-for-concept-request-url-list⑦
    link-type="dfn"} is a
    [clone](https://infra.spec.whatwg.org/#list-clone){#ref-for-list-clone②
    link-type="dfn"} of `request`{.variable}'s [URL
    list](#concept-request-url-list){#ref-for-concept-request-url-list⑧
    link-type="dfn"},
    [initiator](#concept-request-initiator){#ref-for-concept-request-initiator⑥
    link-type="dfn"} is `request`{.variable}'s
    [initiator](#concept-request-initiator){#ref-for-concept-request-initiator⑦
    link-type="dfn"},
    [destination](#concept-request-destination){#ref-for-concept-request-destination①⑧
    link-type="dfn"} is `request`{.variable}'s
    [destination](#concept-request-destination){#ref-for-concept-request-destination①⑨
    link-type="dfn"},
    [origin](#concept-request-origin){#ref-for-concept-request-origin①⑦
    link-type="dfn"} is `request`{.variable}'s
    [origin](#concept-request-origin){#ref-for-concept-request-origin①⑧
    link-type="dfn"},
    [referrer](#concept-request-referrer){#ref-for-concept-request-referrer⑥
    link-type="dfn"} is `request`{.variable}'s
    [referrer](#concept-request-referrer){#ref-for-concept-request-referrer⑦
    link-type="dfn"}, [referrer
    policy](#concept-request-referrer-policy){#ref-for-concept-request-referrer-policy④
    link-type="dfn"} is `request`{.variable}'s [referrer
    policy](#concept-request-referrer-policy){#ref-for-concept-request-referrer-policy⑤
    link-type="dfn"},
    [mode](#concept-request-mode){#ref-for-concept-request-mode②②
    link-type="dfn"} is \"`cors`\", and [response
    tainting](#concept-request-response-tainting){#ref-for-concept-request-response-tainting①⑦
    link-type="dfn"} is \"`cors`\".

    The [service-workers
    mode](#request-service-workers-mode){#ref-for-request-service-workers-mode③
    link-type="dfn"} of `preflight`{.variable} does not matter as this
    algorithm uses [HTTP-network-or-cache
    fetch](#concept-http-network-or-cache-fetch){#ref-for-concept-http-network-or-cache-fetch④
    link-type="dfn"} rather than [HTTP
    fetch](#concept-http-fetch){#ref-for-concept-http-fetch⑥
    link-type="dfn"}.

2.  [Append](#concept-header-list-append){#ref-for-concept-header-list-append①⑧
    link-type="dfn"} (\``Accept`\`, \``*/*`\`) to
    `preflight`{.variable}'s [header
    list](#concept-request-header-list){#ref-for-concept-request-header-list③⑨
    link-type="dfn"}.

3.  [Append](#concept-header-list-append){#ref-for-concept-header-list-append①⑨
    link-type="dfn"}
    (\`[`Access-Control-Request-Method`](#http-access-control-request-method){#ref-for-http-access-control-request-method①
    link-type="http-header"}\`, `request`{.variable}'s
    [method](#concept-request-method){#ref-for-concept-request-method①⑥
    link-type="dfn"}) to `preflight`{.variable}'s [header
    list](#concept-request-header-list){#ref-for-concept-request-header-list④⓪
    link-type="dfn"}.

4.  Let `headers`{.variable} be the [CORS-unsafe request-header
    names](#cors-unsafe-request-header-names){#ref-for-cors-unsafe-request-header-names②
    link-type="dfn"} with `request`{.variable}'s [header
    list](#concept-request-header-list){#ref-for-concept-request-header-list④①
    link-type="dfn"}.

5.  If `headers`{.variable} [is not
    empty](https://infra.spec.whatwg.org/#list-is-empty){#ref-for-list-is-empty⑤
    link-type="dfn"}, then:

    1.  Let `value`{.variable} be the items in `headers`{.variable}
        separated from each other by \``,`\`.

    2.  [Append](#concept-header-list-append){#ref-for-concept-header-list-append②⓪
        link-type="dfn"}
        (\`[`Access-Control-Request-Headers`](#http-access-control-request-headers){#ref-for-http-access-control-request-headers①
        link-type="http-header"}\`, `value`{.variable}) to
        `preflight`{.variable}'s [header
        list](#concept-request-header-list){#ref-for-concept-request-header-list④②
        link-type="dfn"}.

    This intentionally does not use
    [combine](#concept-header-list-combine){#ref-for-concept-header-list-combine①
    link-type="dfn"}, as 0x20 following 0x2C is not the way this was
    implemented, for better or worse.

6.  Let `response`{.variable} be the result of running
    [HTTP-network-or-cache
    fetch](#concept-http-network-or-cache-fetch){#ref-for-concept-http-network-or-cache-fetch⑤
    link-type="dfn"} given a new [fetch
    params](#fetch-params){#ref-for-fetch-params①② link-type="dfn"}
    whose
    [request](#fetch-params-request){#ref-for-fetch-params-request②⓪
    link-type="dfn"} is `preflight`{.variable}.

7.  If a [CORS check](#concept-cors-check){#ref-for-concept-cors-check③
    link-type="dfn"} for `request`{.variable} and `response`{.variable}
    returns success and `response`{.variable}'s
    [status](#concept-response-status){#ref-for-concept-response-status②⓪
    link-type="dfn"} is an [ok status](#ok-status){#ref-for-ok-status①
    link-type="dfn"}, then:

    The [CORS check](#concept-cors-check){#ref-for-concept-cors-check④
    link-type="dfn"} is done on `request`{.variable} rather than
    `preflight`{.variable} to ensure the correct [credentials
    mode](#concept-request-credentials-mode){#ref-for-concept-request-credentials-mode①③
    link-type="dfn"} is used.

    1.  Let `methods`{.variable} be the result of [extracting header
        list
        values](#extract-header-list-values){#ref-for-extract-header-list-values③
        link-type="dfn"} given
        \`[`Access-Control-Allow-Methods`](#http-access-control-allow-methods){#ref-for-http-access-control-allow-methods②
        link-type="http-header"}\` and `response`{.variable}'s [header
        list](#concept-response-header-list){#ref-for-concept-response-header-list②⑥
        link-type="dfn"}.

    2.  Let `headerNames`{.variable} be the result of [extracting header
        list
        values](#extract-header-list-values){#ref-for-extract-header-list-values④
        link-type="dfn"} given
        \`[`Access-Control-Allow-Headers`](#http-access-control-allow-headers){#ref-for-http-access-control-allow-headers②
        link-type="http-header"}\` and `response`{.variable}'s [header
        list](#concept-response-header-list){#ref-for-concept-response-header-list②⑦
        link-type="dfn"}.

    3.  If either `methods`{.variable} or `headerNames`{.variable} is
        failure, return a [network
        error](#concept-network-error){#ref-for-concept-network-error⑤④
        link-type="dfn"}.

    4.  If `methods`{.variable} is null and `request`{.variable}'s
        [use-CORS-preflight
        flag](#use-cors-preflight-flag){#ref-for-use-cors-preflight-flag④
        link-type="dfn"} is set, then set `methods`{.variable} to a new
        list containing `request`{.variable}'s
        [method](#concept-request-method){#ref-for-concept-request-method①⑦
        link-type="dfn"}.

        This ensures that a [CORS-preflight
        fetch](#cors-preflight-fetch-0){#ref-for-cors-preflight-fetch-0⑥
        link-type="dfn"} that happened due to `request`{.variable}'s
        [use-CORS-preflight
        flag](#use-cors-preflight-flag){#ref-for-use-cors-preflight-flag⑤
        link-type="dfn"} being set is
        [cached](#concept-cache){#ref-for-concept-cache②
        link-type="dfn"}.

    5.  If `request`{.variable}'s
        [method](#concept-request-method){#ref-for-concept-request-method①⑧
        link-type="dfn"} is not in `methods`{.variable},
        `request`{.variable}'s
        [method](#concept-request-method){#ref-for-concept-request-method①⑨
        link-type="dfn"} is not a [CORS-safelisted
        method](#cors-safelisted-method){#ref-for-cors-safelisted-method③
        link-type="dfn"}, and `request`{.variable}'s [credentials
        mode](#concept-request-credentials-mode){#ref-for-concept-request-credentials-mode①④
        link-type="dfn"} is \"`include`\" or `methods`{.variable} does
        not contain \``*`\`, then return a [network
        error](#concept-network-error){#ref-for-concept-network-error⑤⑤
        link-type="dfn"}.

    6.  If one of `request`{.variable}'s [header
        list](#concept-request-header-list){#ref-for-concept-request-header-list④③
        link-type="dfn"}'s
        [names](#concept-header-name){#ref-for-concept-header-name②②
        link-type="dfn"} is a [CORS non-wildcard request-header
        name](#cors-non-wildcard-request-header-name){#ref-for-cors-non-wildcard-request-header-name①
        link-type="dfn"} and is not a
        [byte-case-insensitive](https://infra.spec.whatwg.org/#byte-case-insensitive){#ref-for-byte-case-insensitive①⑥
        link-type="dfn"} match for an
        [item](https://infra.spec.whatwg.org/#list-item){#ref-for-list-item②
        link-type="dfn"} in `headerNames`{.variable}, then return a
        [network
        error](#concept-network-error){#ref-for-concept-network-error⑤⑥
        link-type="dfn"}.

    7.  [For
        each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate①③
        link-type="dfn"} `unsafeName`{.variable} of the [CORS-unsafe
        request-header
        names](#cors-unsafe-request-header-names){#ref-for-cors-unsafe-request-header-names③
        link-type="dfn"} with `request`{.variable}'s [header
        list](#concept-request-header-list){#ref-for-concept-request-header-list④④
        link-type="dfn"}, if `unsafeName`{.variable} is not a
        [byte-case-insensitive](https://infra.spec.whatwg.org/#byte-case-insensitive){#ref-for-byte-case-insensitive①⑦
        link-type="dfn"} match for an
        [item](https://infra.spec.whatwg.org/#list-item){#ref-for-list-item③
        link-type="dfn"} in `headerNames`{.variable} and
        `request`{.variable}'s [credentials
        mode](#concept-request-credentials-mode){#ref-for-concept-request-credentials-mode①⑤
        link-type="dfn"} is \"`include`\" or `headerNames`{.variable}
        does not contain \``*`\`, return a [network
        error](#concept-network-error){#ref-for-concept-network-error⑤⑦
        link-type="dfn"}.

    8.  Let `max-age`{.variable} be the result of [extracting header
        list
        values](#extract-header-list-values){#ref-for-extract-header-list-values⑤
        link-type="dfn"} given
        \`[`Access-Control-Max-Age`](#http-access-control-max-age){#ref-for-http-access-control-max-age
        link-type="http-header"}\` and `response`{.variable}'s [header
        list](#concept-response-header-list){#ref-for-concept-response-header-list②⑧
        link-type="dfn"}.

    9.  If `max-age`{.variable} is failure or null, then set
        `max-age`{.variable} to 5.

    10. If `max-age`{.variable} is greater than an imposed limit on
        [max-age](#concept-cache-max-age){#ref-for-concept-cache-max-age
        link-type="dfn"}, then set `max-age`{.variable} to the imposed
        limit.

    11. If the user agent does not provide for a
        [cache](#concept-cache){#ref-for-concept-cache③
        link-type="dfn"}, then return `response`{.variable}.

    12. For each `method`{.variable} in `methods`{.variable} for which
        there is a [method cache entry
        match](#concept-cache-match-method){#ref-for-concept-cache-match-method①
        link-type="dfn"} using `request`{.variable}, set matching
        entry's
        [max-age](#concept-cache-max-age){#ref-for-concept-cache-max-age①
        link-type="dfn"} to `max-age`{.variable}.

    13. For each `method`{.variable} in `methods`{.variable} for which
        there is no [method cache entry
        match](#concept-cache-match-method){#ref-for-concept-cache-match-method②
        link-type="dfn"} using `request`{.variable}, [create a new cache
        entry](#concept-cache-create-entry){#ref-for-concept-cache-create-entry
        link-type="dfn"} with `request`{.variable},
        `max-age`{.variable}, `method`{.variable}, and null.

    14. For each `headerName`{.variable} in `headerNames`{.variable} for
        which there is a [header-name cache entry
        match](#concept-cache-match-header){#ref-for-concept-cache-match-header①
        link-type="dfn"} using `request`{.variable}, set matching
        entry's
        [max-age](#concept-cache-max-age){#ref-for-concept-cache-max-age②
        link-type="dfn"} to `max-age`{.variable}.

    15. For each `headerName`{.variable} in `headerNames`{.variable} for
        which there is no [header-name cache entry
        match](#concept-cache-match-header){#ref-for-concept-cache-match-header②
        link-type="dfn"} using `request`{.variable}, [create a new cache
        entry](#concept-cache-create-entry){#ref-for-concept-cache-create-entry①
        link-type="dfn"} with `request`{.variable},
        `max-age`{.variable}, null, and `headerName`{.variable}.

    16. Return `response`{.variable}.

8.  Otherwise, return a [network
    error](#concept-network-error){#ref-for-concept-network-error⑤⑧
    link-type="dfn"}.
:::

### [4.8. ]{.secno}[CORS-preflight cache]{.content}[](#cors-preflight-cache){.self-link} {#cors-preflight-cache .heading .settled level="4.8"}

A user agent has an associated [CORS-preflight
cache](#concept-cache){#ref-for-concept-cache④ link-type="dfn"}. A
[CORS-preflight cache]{#concept-cache .dfn .dfn-paneled dfn-type="dfn"
noexport=""} is a
[list](https://infra.spec.whatwg.org/#list){#ref-for-list①②
link-type="dfn"} of [cache entries](#cache-entry){#ref-for-cache-entry
link-type="dfn"}.

A [cache entry]{#cache-entry .dfn .dfn-paneled dfn-type="dfn"
noexport=""} consists of:

- [key]{#concept-cache-key .dfn .dfn-paneled dfn-for="cache entry"
  dfn-type="dfn" noexport=""} (a [network partition
  key](#network-partition-key){#ref-for-network-partition-key④
  link-type="dfn"})
- [byte-serialized origin]{#concept-cache-origin .dfn .dfn-paneled
  dfn-for="cache entry" dfn-type="dfn" noexport=""} (a [byte
  sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence①⑧
  link-type="dfn"})
- [URL]{#concept-cache-url .dfn .dfn-paneled dfn-for="cache entry"
  dfn-type="dfn" noexport=""} (a
  [URL](https://url.spec.whatwg.org/#concept-url){#ref-for-concept-url②⓪
  link-type="dfn"})
- [max-age]{#concept-cache-max-age .dfn .dfn-paneled
  dfn-for="cache entry" dfn-type="dfn" noexport=""} (a number of
  seconds)
- [credentials]{#concept-cache-credentials .dfn .dfn-paneled
  dfn-for="cache entry" dfn-type="dfn" noexport=""} (a boolean)
- [method]{#concept-cache-method .dfn .dfn-paneled dfn-for="cache entry"
  dfn-type="dfn" noexport=""} (null, \``*`\`, or a
  [method](#concept-method){#ref-for-concept-method①① link-type="dfn"})
- [header name]{#concept-cache-header-name .dfn .dfn-paneled
  dfn-for="cache entry" dfn-type="dfn" noexport=""} (null, \``*`\`, or a
  [header name](#header-name){#ref-for-header-name①⑧ link-type="dfn"})

[Cache entries](#cache-entry){#ref-for-cache-entry① link-type="dfn"}
must be removed after the seconds specified in their
[max-age](#concept-cache-max-age){#ref-for-concept-cache-max-age③
link-type="dfn"} field have passed since storing the entry. [Cache
entries](#cache-entry){#ref-for-cache-entry② link-type="dfn"} may be
removed before that moment arrives.

::: {.algorithm algorithm="create a new cache entry"}
To [create a new cache entry]{#concept-cache-create-entry .dfn
.dfn-paneled dfn-type="dfn" noexport=""}, given `request`{.variable},
`max-age`{.variable}, `method`{.variable}, and `headerName`{.variable},
run these steps:

1.  Let `entry`{.variable} be a [cache
    entry](#cache-entry){#ref-for-cache-entry③ link-type="dfn"},
    initialized as follows:

    [key](#concept-cache-key){#ref-for-concept-cache-key link-type="dfn"}

    :   The result of [determining the network partition
        key](#request-determine-the-network-partition-key){#ref-for-request-determine-the-network-partition-key②
        link-type="dfn"} given `request`{.variable}

    [byte-serialized origin](#concept-cache-origin){#ref-for-concept-cache-origin link-type="dfn"}

    :   The result of [byte-serializing a request
        origin](#byte-serializing-a-request-origin){#ref-for-byte-serializing-a-request-origin②
        link-type="dfn"} with `request`{.variable}

    [URL](#concept-cache-url){#ref-for-concept-cache-url link-type="dfn"}

    :   `request`{.variable}'s [current
        URL](#concept-request-current-url){#ref-for-concept-request-current-url③③
        link-type="dfn"}

    [max-age](#concept-cache-max-age){#ref-for-concept-cache-max-age④ link-type="dfn"}

    :   `max-age`{.variable}

    [credentials](#concept-cache-credentials){#ref-for-concept-cache-credentials link-type="dfn"}

    :   True if `request`{.variable}'s [credentials
        mode](#concept-request-credentials-mode){#ref-for-concept-request-credentials-mode①⑥
        link-type="dfn"} is \"`include`\", and false otherwise

    [method](#concept-cache-method){#ref-for-concept-cache-method link-type="dfn"}

    :   `method`{.variable}

    [header name](#concept-cache-header-name){#ref-for-concept-cache-header-name link-type="dfn"}

    :   `headerName`{.variable}

2.  [Append](https://infra.spec.whatwg.org/#list-append){#ref-for-list-append①⓪
    link-type="dfn"} `entry`{.variable} to the user agent's
    [CORS-preflight cache](#concept-cache){#ref-for-concept-cache⑤
    link-type="dfn"}.
:::

To [clear cache entries]{#concept-cache-clear .dfn .dfn-paneled
dfn-type="dfn" noexport=""}, given a `request`{.variable},
[remove](https://infra.spec.whatwg.org/#list-remove){#ref-for-list-remove②
link-type="dfn"} any [cache entries](#cache-entry){#ref-for-cache-entry④
link-type="dfn"} in the user agent's [CORS-preflight
cache](#concept-cache){#ref-for-concept-cache⑥ link-type="dfn"} whose
[key](#concept-cache-key){#ref-for-concept-cache-key① link-type="dfn"}
is the result of [determining the network partition
key](#request-determine-the-network-partition-key){#ref-for-request-determine-the-network-partition-key③
link-type="dfn"} given `request`{.variable}, [byte-serialized
origin](#concept-cache-origin){#ref-for-concept-cache-origin①
link-type="dfn"} is the result of [byte-serializing a request
origin](#byte-serializing-a-request-origin){#ref-for-byte-serializing-a-request-origin③
link-type="dfn"} with `request`{.variable}, and
[URL](#concept-cache-url){#ref-for-concept-cache-url① link-type="dfn"}
is `request`{.variable}'s [current
URL](#concept-request-current-url){#ref-for-concept-request-current-url③④
link-type="dfn"}.

There is a [cache entry match]{#concept-cache-match .dfn .dfn-paneled
dfn-type="dfn" noexport=""} for a [cache
entry](#cache-entry){#ref-for-cache-entry⑤ link-type="dfn"}
`entry`{.variable} with `request`{.variable} if `entry`{.variable}'s
[key](#concept-cache-key){#ref-for-concept-cache-key② link-type="dfn"}
is the result of [determining the network partition
key](#request-determine-the-network-partition-key){#ref-for-request-determine-the-network-partition-key④
link-type="dfn"} given `request`{.variable}, `entry`{.variable}'s
[byte-serialized
origin](#concept-cache-origin){#ref-for-concept-cache-origin②
link-type="dfn"} is the result of [byte-serializing a request
origin](#byte-serializing-a-request-origin){#ref-for-byte-serializing-a-request-origin④
link-type="dfn"} with `request`{.variable}, `entry`{.variable}'s
[URL](#concept-cache-url){#ref-for-concept-cache-url② link-type="dfn"}
is `request`{.variable}'s [current
URL](#concept-request-current-url){#ref-for-concept-request-current-url③⑤
link-type="dfn"}, and one of

- `entry`{.variable}'s
  [credentials](#concept-cache-credentials){#ref-for-concept-cache-credentials①
  link-type="dfn"} is true
- `entry`{.variable}'s
  [credentials](#concept-cache-credentials){#ref-for-concept-cache-credentials②
  link-type="dfn"} is false and `request`{.variable}'s [credentials
  mode](#concept-request-credentials-mode){#ref-for-concept-request-credentials-mode①⑦
  link-type="dfn"} is not \"`include`\".

is true.

There is a [method cache entry match]{#concept-cache-match-method .dfn
.dfn-paneled dfn-type="dfn" noexport=""} for `method`{.variable} using
`request`{.variable} when there is a [cache
entry](#cache-entry){#ref-for-cache-entry⑥ link-type="dfn"} in the user
agent's [CORS-preflight cache](#concept-cache){#ref-for-concept-cache⑦
link-type="dfn"} for which there is a [cache entry
match](#concept-cache-match){#ref-for-concept-cache-match
link-type="dfn"} with `request`{.variable} and its
[method](#concept-cache-method){#ref-for-concept-cache-method①
link-type="dfn"} is `method`{.variable} or \``*`\`.

There is a [header-name cache entry match]{#concept-cache-match-header
.dfn .dfn-paneled dfn-type="dfn" noexport=""} for
`headerName`{.variable} using `request`{.variable} when there is a
[cache entry](#cache-entry){#ref-for-cache-entry⑦ link-type="dfn"} in
the user agent's [CORS-preflight
cache](#concept-cache){#ref-for-concept-cache⑧ link-type="dfn"} for
which there is a [cache entry
match](#concept-cache-match){#ref-for-concept-cache-match①
link-type="dfn"} with `request`{.variable} and one of

- its [header
  name](#concept-cache-header-name){#ref-for-concept-cache-header-name①
  link-type="dfn"} is a
  [byte-case-insensitive](https://infra.spec.whatwg.org/#byte-case-insensitive){#ref-for-byte-case-insensitive①⑧
  link-type="dfn"} match for `headerName`{.variable}
- its [header
  name](#concept-cache-header-name){#ref-for-concept-cache-header-name②
  link-type="dfn"} is \``*`\` and `headerName`{.variable} is not a [CORS
  non-wildcard request-header
  name](#cors-non-wildcard-request-header-name){#ref-for-cors-non-wildcard-request-header-name②
  link-type="dfn"}

is true.

### [4.9. ]{.secno}[CORS check]{.content}[](#cors-check){.self-link} {#cors-check .heading .settled level="4.9"}

::: {.algorithm algorithm="CORS check"}
To perform a [CORS check]{#concept-cors-check .dfn .dfn-paneled
dfn-type="dfn" noexport=""} for a `request`{.variable} and
`response`{.variable}, run these steps:

1.  Let `origin`{.variable} be the result of
    [getting](#concept-header-list-get){#ref-for-concept-header-list-get⑤
    link-type="dfn"}
    \`[`Access-Control-Allow-Origin`](#http-access-control-allow-origin){#ref-for-http-access-control-allow-origin⑥
    link-type="http-header"}\` from `response`{.variable}'s [header
    list](#concept-response-header-list){#ref-for-concept-response-header-list②⑨
    link-type="dfn"}.

2.  If `origin`{.variable} is null, then return failure.

    Null is not \``null`\`.

3.  If `request`{.variable}'s [credentials
    mode](#concept-request-credentials-mode){#ref-for-concept-request-credentials-mode①⑧
    link-type="dfn"} is not \"`include`\" and `origin`{.variable} is
    \``*`\`, then return success.

4.  If the result of [byte-serializing a request
    origin](#byte-serializing-a-request-origin){#ref-for-byte-serializing-a-request-origin⑤
    link-type="dfn"} with `request`{.variable} is not
    `origin`{.variable}, then return failure.

5.  If `request`{.variable}'s [credentials
    mode](#concept-request-credentials-mode){#ref-for-concept-request-credentials-mode①⑨
    link-type="dfn"} is not \"`include`\", then return success.

6.  Let `credentials`{.variable} be the result of
    [getting](#concept-header-list-get){#ref-for-concept-header-list-get⑥
    link-type="dfn"}
    \`[`Access-Control-Allow-Credentials`](#http-access-control-allow-credentials){#ref-for-http-access-control-allow-credentials④
    link-type="http-header"}\` from `response`{.variable}'s [header
    list](#concept-response-header-list){#ref-for-concept-response-header-list③⓪
    link-type="dfn"}.

7.  If `credentials`{.variable} is \``true`\`, then return success.

8.  Return failure.
:::

### [4.10. ]{.secno}[TAO check]{.content}[](#tao-check){.self-link} {#tao-check .heading .settled level="4.10"}

::: {.algorithm algorithm="TAO check"}
To perform a [TAO check]{#concept-tao-check .dfn .dfn-paneled
dfn-type="dfn" noexport=""} for a `request`{.variable} and
`response`{.variable}, run these steps:

1.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert②②
    link-type="dfn"}: `request`{.variable}'s
    [origin](#concept-request-origin){#ref-for-concept-request-origin①⑨
    link-type="dfn"} is not \"`client`\".

2.  If `request`{.variable}'s [timing allow failed
    flag](#timing-allow-failed){#ref-for-timing-allow-failed④
    link-type="dfn"} is set, then return failure.

3.  Let `values`{.variable} be the result of [getting, decoding, and
    splitting](#concept-header-list-get-decode-split){#ref-for-concept-header-list-get-decode-split⑥
    link-type="dfn"} \``Timing-Allow-Origin`\` from
    `response`{.variable}'s [header
    list](#concept-response-header-list){#ref-for-concept-response-header-list③①
    link-type="dfn"}.

4.  If `values`{.variable}
    [contains](https://infra.spec.whatwg.org/#list-contain){#ref-for-list-contain①
    link-type="dfn"} \"`*`\", then return success.

5.  If `values`{.variable}
    [contains](https://infra.spec.whatwg.org/#list-contain){#ref-for-list-contain②
    link-type="dfn"} the result of [serializing a request
    origin](#serializing-a-request-origin){#ref-for-serializing-a-request-origin①
    link-type="dfn"} with `request`{.variable}, then return success.

6.  If `request`{.variable}'s
    [mode](#concept-request-mode){#ref-for-concept-request-mode②③
    link-type="dfn"} is \"`navigate`\" and `request`{.variable}'s
    [current
    URL](#concept-request-current-url){#ref-for-concept-request-current-url③⑥
    link-type="dfn"}'s
    [origin](https://url.spec.whatwg.org/#concept-url-origin){#ref-for-concept-url-origin①⑨
    link-type="dfn"} is not [same
    origin](https://html.spec.whatwg.org/multipage/browsers.html#same-origin){#ref-for-same-origin⑨
    link-type="dfn"} with `request`{.variable}'s
    [origin](#concept-request-origin){#ref-for-concept-request-origin②⓪
    link-type="dfn"}, then return failure.

    This is necessary for navigations of a nested navigable. There,
    `request`{.variable}'s
    [origin](#concept-request-origin){#ref-for-concept-request-origin②①
    link-type="dfn"} would be the container document's
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#ref-for-concept-document-origin
    link-type="dfn"} and the [TAO
    check](#concept-tao-check){#ref-for-concept-tao-check①
    link-type="dfn"} would return failure. Since navigation timing never
    validates the results of the [TAO
    check](#concept-tao-check){#ref-for-concept-tao-check②
    link-type="dfn"}, the nested document would still have access to the
    full timing information, but the container document would not.

7.  If `request`{.variable}'s [response
    tainting](#concept-request-response-tainting){#ref-for-concept-request-response-tainting①⑧
    link-type="dfn"} is \"`basic`\", then return success.

8.  Return failure.
:::

## [5. ]{.secno}[Fetch API]{.content}[](#fetch-api){.self-link} {#fetch-api .heading .settled level="5"}

The [`fetch()`](#dom-global-fetch){#ref-for-dom-global-fetch④ .idl-code
link-type="method"} method is relatively low-level API for
[fetching](#concept-fetch){#ref-for-concept-fetch②⑥ link-type="dfn"}
resources. It covers slightly more ground than
[`XMLHttpRequest`{.idl}](https://xhr.spec.whatwg.org/#xmlhttprequest){#ref-for-xmlhttprequest⑤
link-type="idl"}, although it is currently lacking when it comes to
request progression (not response progression).

::: {#fetch-blob-example .example}
[](#fetch-blob-example){.self-link}

The [`fetch()`](#dom-global-fetch){#ref-for-dom-global-fetch⑤ .idl-code
link-type="method"} method makes it quite straightforward to
[fetch](#concept-fetch){#ref-for-concept-fetch②⑦ link-type="dfn"} a
resource and extract its contents as a
[`Blob`{.idl}](https://w3c.github.io/FileAPI/#dfn-Blob){#ref-for-dfn-Blob②
link-type="idl"}:

``` {.lang-javascript .highlight}
fetch("/music/pk/altes-kamuffel.flac")
  .then(res => res.blob()).then(playBlob)
```

If you just care to log a particular response header:

``` {.lang-javascript .highlight}
fetch("/", {method:"HEAD"})
  .then(res => log(res.headers.get("strict-transport-security")))
```

If you want to check a particular response header and then process the
response of a cross-origin resource:

``` {.lang-javascript .highlight}
fetch("https://pk.example/berlin-calling.json", {mode:"cors"})
  .then(res => {
    if(res.headers.get("content-type") &&
       res.headers.get("content-type").toLowerCase().indexOf("application/json") >= 0) {
      return res.json()
    } else {
      throw new TypeError()
    }
  }).then(processJSON)
```

If you want to work with URL query parameters:

``` {.lang-javascript .highlight}
var url = new URL("https://geo.example.org/api"),
    params = {lat:35.696233, long:139.570431}
Object.keys(params).forEach(key => url.searchParams.append(key, params[key]))
fetch(url).then(/* … */)
```

If you want to receive the body data progressively:

``` {.lang-javascript .highlight}
function consume(reader) {
  var total = 0
  return pump()
  function pump() {
    return reader.read().then(({done, value}) => {
      if (done) {
        return
      }
      total += value.byteLength
      log(`received ${value.byteLength} bytes (${total} bytes in total)`)
      return pump()
    })
  }
}

fetch("/music/pk/altes-kamuffel.flac")
  .then(res => consume(res.body.getReader()))
  .then(() => log("consumed the entire body without keeping the whole thing in memory!"))
  .catch(e => log("something went wrong: " + e))
```
:::

### [5.1. ]{.secno}[Headers class]{.content}[](#headers-class){.self-link} {#headers-class .heading .settled level="5.1"}

``` {.idl .highlight .def}
typedef (sequence<sequence<ByteString>> or record<ByteString, ByteString>) HeadersInit;

[Exposed=(Window,Worker)]
interface Headers {
  constructor(optional HeadersInit init);

  undefined append(ByteString name, ByteString value);
  undefined delete(ByteString name);
  ByteString? get(ByteString name);
  sequence<ByteString> getSetCookie();
  boolean has(ByteString name);
  undefined set(ByteString name, ByteString value);
  iterable<ByteString, ByteString>;
};
```

A [`Headers`{.idl}](#headers){#ref-for-headers① link-type="idl"} object
has an associated [header list]{#concept-headers-header-list .dfn
.dfn-paneled dfn-for="Headers" dfn-type="dfn" export=""} (a [header
list](#concept-header-list){#ref-for-concept-header-list②②
link-type="dfn"}), which is initially empty. [This can be a pointer to
the [header list](#concept-header-list){#ref-for-concept-header-list②③
link-type="dfn"} of something else, e.g., of a
[request](#concept-request){#ref-for-concept-request①⓪⑨ link-type="dfn"}
as demonstrated by [`Request`{.idl}](#request){#ref-for-request
link-type="idl"} objects.]{.note}

A [`Headers`{.idl}](#headers){#ref-for-headers② link-type="idl"} object
also has an associated [guard]{#concept-headers-guard .dfn .dfn-paneled
dfn-for="Headers" dfn-type="dfn" export=""}, which is a [headers
guard]{#headers-guard .dfn .dfn-paneled dfn-type="dfn" noexport=""}. A
[headers guard](#headers-guard){#ref-for-headers-guard link-type="dfn"}
is \"`immutable`\", \"`request`\", \"`request-no-cors`\", \"`response`\"
or \"`none`\".

------------------------------------------------------------------------

`headers`{.variable}` = new `[`Headers`](#dom-headers){#ref-for-dom-headers① .idl-code link-type="constructor"}`([``init`{.variable}`])`

:   Creates a new [`Headers`{.idl}](#headers){#ref-for-headers③
    link-type="idl"} object. `init`{.variable} can be used to fill its
    internal header list, as per the example below.

    ::: {#example-headers-class .example}
    [](#example-headers-class){.self-link}
    ``` {.lang-javascript .highlight}
    const meta = { "Content-Type": "text/xml", "Breaking-Bad": "<3" };
    new Headers(meta);

    // The above is equivalent to
    const meta2 = [
      [ "Content-Type", "text/xml" ],
      [ "Breaking-Bad", "<3" ]
    ];
    new Headers(meta2);
    ```
    :::

`headers`{.variable}` . `[`append`](#dom-headers-append){#ref-for-dom-headers-append① .idl-code link-type="method"}`(``name`{.variable}`, ``value`{.variable}`)`

:   Appends a header to `headers`{.variable}.

`headers`{.variable}` . `[`delete`](#dom-headers-delete){#ref-for-dom-headers-delete① .idl-code link-type="method"}`(``name`{.variable}`)`

:   Removes a header from `headers`{.variable}.

`headers`{.variable}` . `[`get`](#dom-headers-get){#ref-for-dom-headers-get① .idl-code link-type="method"}`(``name`{.variable}`)`

:   Returns as a string the values of all headers whose name is
    `name`{.variable}, separated by a comma and a space.

`headers`{.variable}` . `[`getSetCookie`](#dom-headers-getsetcookie){#ref-for-dom-headers-getsetcookie① .idl-code link-type="method"}`()`

:   Returns a list of the values for all headers whose name is
    \``Set-Cookie`\`.

`headers`{.variable}` . `[`has`](#dom-headers-has){#ref-for-dom-headers-has① .idl-code link-type="method"}`(``name`{.variable}`)`

:   Returns whether there is a header whose name is `name`{.variable}.

`headers`{.variable}` . `[`set`](#dom-headers-set){#ref-for-dom-headers-set① .idl-code link-type="method"}`(``name`{.variable}`, ``value`{.variable}`)`

:   Replaces the value of the first header whose name is
    `name`{.variable} with `value`{.variable} and removes any remaining
    headers whose name is `name`{.variable}.

`for(const [``name`{.variable}`, ``value`{.variable}`] of ``headers`{.variable}`)`

:   `headers`{.variable} can be iterated over.

------------------------------------------------------------------------

::: {.algorithm algorithm="validate" algorithm-for="Headers"}
To [validate]{#headers-validate .dfn .dfn-paneled dfn-for="Headers"
dfn-type="dfn" noexport=""} a
[header](#concept-header){#ref-for-concept-header⑤⑥ link-type="dfn"}
(`name`{.variable}, `value`{.variable}) for a
[`Headers`{.idl}](#headers){#ref-for-headers④ link-type="idl"} object
`headers`{.variable}:

1.  If `name`{.variable} is not a [header
    name](#header-name){#ref-for-header-name①⑨ link-type="dfn"} or
    `value`{.variable} is not a [header
    value](#header-value){#ref-for-header-value⑨ link-type="dfn"}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw
    link-type="dfn"} a
    [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror②
    link-type="idl"}.

2.  If `headers`{.variable}'s
    [guard](#concept-headers-guard){#ref-for-concept-headers-guard
    link-type="dfn"} is \"`immutable`\", then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw①
    link-type="dfn"} a
    [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror③
    link-type="idl"}.

3.  If `headers`{.variable}'s
    [guard](#concept-headers-guard){#ref-for-concept-headers-guard①
    link-type="dfn"} is \"`request`\" and (`name`{.variable},
    `value`{.variable}) is a [forbidden
    request-header](#forbidden-request-header){#ref-for-forbidden-request-header①
    link-type="dfn"}, then return false.

4.  If `headers`{.variable}'s
    [guard](#concept-headers-guard){#ref-for-concept-headers-guard②
    link-type="dfn"} is \"`response`\" and `name`{.variable} is a
    [forbidden response-header
    name](#forbidden-response-header-name){#ref-for-forbidden-response-header-name②
    link-type="dfn"}, then return false.

5.  Return true.
:::

Steps for \"`request-no-cors`\" are not shared as you cannot have a fake
value (for
[`delete()`{.idl}](#dom-headers-delete){#ref-for-dom-headers-delete②
link-type="idl"}) that always succeeds in [CORS-safelisted
request-header](#cors-safelisted-request-header){#ref-for-cors-safelisted-request-header③
link-type="dfn"}.

::: {.algorithm algorithm="append" algorithm-for="Headers"}
To [append]{#concept-headers-append .dfn .dfn-paneled dfn-for="Headers"
dfn-type="dfn" export=""} a
[header](#concept-header){#ref-for-concept-header⑤⑦ link-type="dfn"}
(`name`{.variable}, `value`{.variable}) to a
[`Headers`{.idl}](#headers){#ref-for-headers⑤ link-type="idl"} object
`headers`{.variable}, run these steps:

1.  [Normalize](#concept-header-value-normalize){#ref-for-concept-header-value-normalize
    link-type="dfn"} `value`{.variable}.

2.  If [validating](#headers-validate){#ref-for-headers-validate
    link-type="dfn"} (`name`{.variable}, `value`{.variable}) for
    `headers`{.variable} returns false, then return.

3.  If `headers`{.variable}'s
    [guard](#concept-headers-guard){#ref-for-concept-headers-guard③
    link-type="dfn"} is \"`request-no-cors`\":

    1.  Let `temporaryValue`{.variable} be the result of
        [getting](#concept-header-list-get){#ref-for-concept-header-list-get⑦
        link-type="dfn"} `name`{.variable} from `headers`{.variable}'s
        [header
        list](#concept-headers-header-list){#ref-for-concept-headers-header-list
        link-type="dfn"}.

    2.  If `temporaryValue`{.variable} is null, then set
        `temporaryValue`{.variable} to `value`{.variable}.

    3.  Otherwise, set `temporaryValue`{.variable} to
        `temporaryValue`{.variable}, followed by 0x2C 0x20, followed by
        `value`{.variable}.

    4.  If (`name`{.variable}, `temporaryValue`{.variable}) is not a
        [no-CORS-safelisted
        request-header](#no-cors-safelisted-request-header){#ref-for-no-cors-safelisted-request-header
        link-type="dfn"}, then return.

4.  [Append](#concept-header-list-append){#ref-for-concept-header-list-append②①
    link-type="dfn"} (`name`{.variable}, `value`{.variable}) to
    `headers`{.variable}'s [header
    list](#concept-headers-header-list){#ref-for-concept-headers-header-list①
    link-type="dfn"}.

5.  If `headers`{.variable}'s
    [guard](#concept-headers-guard){#ref-for-concept-headers-guard④
    link-type="dfn"} is \"`request-no-cors`\", then [remove privileged
    no-CORS
    request-headers](#concept-headers-remove-privileged-no-cors-request-headers){#ref-for-concept-headers-remove-privileged-no-cors-request-headers
    link-type="dfn"} from `headers`{.variable}.
:::

::: {.algorithm algorithm="fill" algorithm-for="Headers"}
To [fill]{#concept-headers-fill .dfn .dfn-paneled dfn-for="Headers"
dfn-type="dfn" export=""} a
[`Headers`{.idl}](#headers){#ref-for-headers⑥ link-type="idl"} object
`headers`{.variable} with a given object `object`{.variable}, run these
steps:

1.  If `object`{.variable} is a
    [sequence](https://webidl.spec.whatwg.org/#idl-sequence){#ref-for-idl-sequence③
    link-type="dfn"}, then [for
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate①④
    link-type="dfn"} `header`{.variable} of `object`{.variable}:

    1.  If `header`{.variable}'s
        [size](https://infra.spec.whatwg.org/#list-size){#ref-for-list-size①
        link-type="dfn"} is not 2, then
        [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw②
        link-type="dfn"} a
        [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror④
        link-type="idl"}.

    2.  [Append](#concept-headers-append){#ref-for-concept-headers-append
        link-type="dfn"} (`header`{.variable}\[0\],
        `header`{.variable}\[1\]) to `headers`{.variable}.

2.  Otherwise, `object`{.variable} is a
    [record](https://tc39.es/ecma262/#sec-list-and-record-specification-type){#ref-for-sec-list-and-record-specification-type②
    link-type="dfn"}, then [for
    each](https://infra.spec.whatwg.org/#map-iterate){#ref-for-map-iterate
    link-type="dfn"} `key`{.variable} → `value`{.variable} of
    `object`{.variable},
    [append](#concept-headers-append){#ref-for-concept-headers-append①
    link-type="dfn"} (`key`{.variable}, `value`{.variable}) to
    `headers`{.variable}.
:::

::: {.algorithm algorithm="remove privileged no-CORS request-headers" algorithm-for="Headers"}
To [remove privileged no-CORS
request-headers]{#concept-headers-remove-privileged-no-cors-request-headers
.dfn .dfn-paneled dfn-for="Headers" dfn-type="dfn" noexport=""} from a
[`Headers`{.idl}](#headers){#ref-for-headers⑦ link-type="idl"} object
(`headers`{.variable}), run these steps:

1.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate①⑤
    link-type="dfn"} `headerName`{.variable} of [privileged no-CORS
    request-header
    names](#privileged-no-cors-request-header-name){#ref-for-privileged-no-cors-request-header-name
    link-type="dfn"}:

    1.  [Delete](#concept-header-list-delete){#ref-for-concept-header-list-delete②
        link-type="dfn"} `headerName`{.variable} from
        `headers`{.variable}'s [header
        list](#concept-headers-header-list){#ref-for-concept-headers-header-list②
        link-type="dfn"}.

This is called when headers are modified by unprivileged code.
:::

::: {.algorithm algorithm="Headers(init)" algorithm-for="Headers"}
The [`new Headers(``init`{.variable}`)`]{#dom-headers .dfn .dfn-paneled
.idl-code dfn-for="Headers" dfn-type="constructor" export=""
lt="Headers(init)|constructor(init)|Headers()|constructor()"}
constructor steps are:

1.  Set [this](https://webidl.spec.whatwg.org/#this){#ref-for-this
    link-type="dfn"}'s
    [guard](#concept-headers-guard){#ref-for-concept-headers-guard⑤
    link-type="dfn"} to \"`none`\".

2.  If `init`{.variable} is given, then
    [fill](#concept-headers-fill){#ref-for-concept-headers-fill
    link-type="dfn"}
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①
    link-type="dfn"} with `init`{.variable}.
:::

::: {.algorithm algorithm="append(name, value)" algorithm-for="Headers"}
The
[`append(``name`{.variable}`, ``value`{.variable}`)`]{#dom-headers-append
.dfn .dfn-paneled .idl-code dfn-for="Headers" dfn-type="method"
export=""} method steps are to
[append](#concept-headers-append){#ref-for-concept-headers-append②
link-type="dfn"} (`name`{.variable}, `value`{.variable}) to
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②
link-type="dfn"}.
:::

::: {.algorithm algorithm="delete(name)" algorithm-for="Headers"}
The [`delete(``name`{.variable}`)`]{#dom-headers-delete .dfn
.dfn-paneled .idl-code dfn-for="Headers" dfn-type="method" export=""}
method steps are:

1.  If [validating](#headers-validate){#ref-for-headers-validate①
    link-type="dfn"} (`name`{.variable}, \`\`) for
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③
    link-type="dfn"} returns false, then return.

    Passing a dummy [header
    value](#header-value){#ref-for-header-value①⓪ link-type="dfn"} ought
    not to have any negative repercussions.

2.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this④
    link-type="dfn"}'s
    [guard](#concept-headers-guard){#ref-for-concept-headers-guard⑥
    link-type="dfn"} is \"`request-no-cors`\", `name`{.variable} is not
    a [no-CORS-safelisted request-header
    name](#no-cors-safelisted-request-header-name){#ref-for-no-cors-safelisted-request-header-name①
    link-type="dfn"}, and `name`{.variable} is not a [privileged no-CORS
    request-header
    name](#privileged-no-cors-request-header-name){#ref-for-privileged-no-cors-request-header-name①
    link-type="dfn"}, then return.

3.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑤
    link-type="dfn"}'s [header
    list](#concept-headers-header-list){#ref-for-concept-headers-header-list③
    link-type="dfn"} [does not
    contain](#header-list-contains){#ref-for-header-list-contains②③
    link-type="dfn"} `name`{.variable}, then return.

4.  [Delete](#concept-header-list-delete){#ref-for-concept-header-list-delete③
    link-type="dfn"} `name`{.variable} from
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑥
    link-type="dfn"}'s [header
    list](#concept-headers-header-list){#ref-for-concept-headers-header-list④
    link-type="dfn"}.

5.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑦
    link-type="dfn"}'s
    [guard](#concept-headers-guard){#ref-for-concept-headers-guard⑦
    link-type="dfn"} is \"`request-no-cors`\", then [remove privileged
    no-CORS
    request-headers](#concept-headers-remove-privileged-no-cors-request-headers){#ref-for-concept-headers-remove-privileged-no-cors-request-headers①
    link-type="dfn"} from
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑧
    link-type="dfn"}.
:::

::: {.algorithm algorithm="get(name)" algorithm-for="Headers"}
The [`get(``name`{.variable}`)`]{#dom-headers-get .dfn .dfn-paneled
.idl-code dfn-for="Headers" dfn-type="method" export=""} method steps
are:

1.  If `name`{.variable} is not a [header
    name](#header-name){#ref-for-header-name②⓪ link-type="dfn"}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw③
    link-type="dfn"} a
    [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror⑤
    link-type="idl"}.

2.  Return the result of
    [getting](#concept-header-list-get){#ref-for-concept-header-list-get⑧
    link-type="dfn"} `name`{.variable} from
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑨
    link-type="dfn"}'s [header
    list](#concept-headers-header-list){#ref-for-concept-headers-header-list⑤
    link-type="dfn"}.
:::

::: {.algorithm algorithm="getSetCookie()" algorithm-for="Headers"}
The [`getSetCookie()`]{#dom-headers-getsetcookie .dfn .dfn-paneled
.idl-code dfn-for="Headers" dfn-type="method" export=""} method steps
are:

1.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⓪
    link-type="dfn"}'s [header
    list](#concept-headers-header-list){#ref-for-concept-headers-header-list⑥
    link-type="dfn"} [does not
    contain](#header-list-contains){#ref-for-header-list-contains②④
    link-type="dfn"} \``Set-Cookie`\`, then return « ».

2.  Return the
    [values](#concept-header-value){#ref-for-concept-header-value②①
    link-type="dfn"} of all
    [headers](#concept-header){#ref-for-concept-header⑤⑧
    link-type="dfn"} in
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①①
    link-type="dfn"}'s [header
    list](#concept-headers-header-list){#ref-for-concept-headers-header-list⑦
    link-type="dfn"} whose
    [name](#concept-header-name){#ref-for-concept-header-name②③
    link-type="dfn"} is a
    [byte-case-insensitive](https://infra.spec.whatwg.org/#byte-case-insensitive){#ref-for-byte-case-insensitive①⑨
    link-type="dfn"} match for \``Set-Cookie`\`, in order.
:::

::: {.algorithm algorithm="has(name)" algorithm-for="Headers"}
The [`has(``name`{.variable}`)`]{#dom-headers-has .dfn .dfn-paneled
.idl-code dfn-for="Headers" dfn-type="method" export=""} method steps
are:

1.  If `name`{.variable} is not a [header
    name](#header-name){#ref-for-header-name②① link-type="dfn"}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw④
    link-type="dfn"} a
    [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror⑥
    link-type="idl"}.

2.  Return true if
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①②
    link-type="dfn"}'s [header
    list](#concept-headers-header-list){#ref-for-concept-headers-header-list⑧
    link-type="dfn"}
    [contains](#header-list-contains){#ref-for-header-list-contains②⑤
    link-type="dfn"} `name`{.variable}; otherwise false.
:::

::: {.algorithm algorithm="set(name, value)" algorithm-for="Headers"}
The [`set(``name`{.variable}`, ``value`{.variable}`)`]{#dom-headers-set
.dfn .dfn-paneled .idl-code dfn-for="Headers" dfn-type="method"
export=""} method steps are:

1.  [Normalize](#concept-header-value-normalize){#ref-for-concept-header-value-normalize①
    link-type="dfn"} `value`{.variable}.

2.  If [validating](#headers-validate){#ref-for-headers-validate②
    link-type="dfn"} (`name`{.variable}, `value`{.variable}) for
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①③
    link-type="dfn"} returns false, then return.

3.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①④
    link-type="dfn"}'s
    [guard](#concept-headers-guard){#ref-for-concept-headers-guard⑧
    link-type="dfn"} is \"`request-no-cors`\" and (`name`{.variable},
    `value`{.variable}) is not a [no-CORS-safelisted
    request-header](#no-cors-safelisted-request-header){#ref-for-no-cors-safelisted-request-header①
    link-type="dfn"}, then return.

4.  [Set](#concept-header-list-set){#ref-for-concept-header-list-set①
    link-type="dfn"} (`name`{.variable}, `value`{.variable}) in
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑤
    link-type="dfn"}'s [header
    list](#concept-headers-header-list){#ref-for-concept-headers-header-list⑨
    link-type="dfn"}.

5.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑥
    link-type="dfn"}'s
    [guard](#concept-headers-guard){#ref-for-concept-headers-guard⑨
    link-type="dfn"} is \"`request-no-cors`\", then [remove privileged
    no-CORS
    request-headers](#concept-headers-remove-privileged-no-cors-request-headers){#ref-for-concept-headers-remove-privileged-no-cors-request-headers②
    link-type="dfn"} from
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑦
    link-type="dfn"}.
:::

The [value pairs to iterate
over](https://webidl.spec.whatwg.org/#dfn-value-pairs-to-iterate-over){#ref-for-dfn-value-pairs-to-iterate-over
link-type="dfn"} are the return value of running [sort and
combine](#concept-header-list-sort-and-combine){#ref-for-concept-header-list-sort-and-combine
link-type="dfn"} with
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑧
link-type="dfn"}'s [header
list](#concept-headers-header-list){#ref-for-concept-headers-header-list①⓪
link-type="dfn"}.

### [5.2. ]{.secno}[BodyInit unions]{.content}[](#bodyinit-unions){.self-link} {#bodyinit-unions .heading .settled level="5.2"}

``` {.idl .highlight .def}
typedef (Blob or BufferSource or FormData or URLSearchParams or USVString) XMLHttpRequestBodyInit;

typedef (ReadableStream or XMLHttpRequestBodyInit) BodyInit;
```

To [safely extract]{#bodyinit-safely-extract .dfn .dfn-paneled
dfn-for="BodyInit" dfn-type="dfn" export=""} a [body with
type](#body-with-type){#ref-for-body-with-type link-type="dfn"} from a
[byte
sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence①⑨
link-type="dfn"} or [`BodyInit`{.idl}](#bodyinit){#ref-for-bodyinit
link-type="idl"} object `object`{.variable}, run these steps:

1.  If `object`{.variable} is a
    [`ReadableStream`{.idl}](https://streams.spec.whatwg.org/#readablestream){#ref-for-readablestream⑤
    link-type="idl"} object, then:

    1.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert②③
        link-type="dfn"}: `object`{.variable} is neither
        [disturbed](https://streams.spec.whatwg.org/#is-readable-stream-disturbed){#ref-for-is-readable-stream-disturbed
        link-type="dfn"} nor
        [locked](https://streams.spec.whatwg.org/#readablestream-locked){#ref-for-readablestream-locked
        link-type="dfn"}.

2.  Return the result of
    [extracting](#concept-bodyinit-extract){#ref-for-concept-bodyinit-extract
    link-type="dfn"} `object`{.variable}.

The [safely
extract](#bodyinit-safely-extract){#ref-for-bodyinit-safely-extract⑥
link-type="dfn"} operation is a subset of the
[extract](#concept-bodyinit-extract){#ref-for-concept-bodyinit-extract①
link-type="dfn"} operation that is guaranteed to not throw an exception.

To [extract]{#concept-bodyinit-extract .dfn .dfn-paneled
dfn-for="BodyInit" dfn-type="dfn" export=""} a [body with
type](#body-with-type){#ref-for-body-with-type① link-type="dfn"} from a
[byte
sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence②⓪
link-type="dfn"} or [`BodyInit`{.idl}](#bodyinit){#ref-for-bodyinit①
link-type="idl"} object `object`{.variable}, with an optional boolean
[`keepalive`{.variable}]{#keepalive .dfn .dfn-paneled
dfn-for="BodyInit/extract" dfn-type="dfn" export=""} (default false),
run these steps:

1.  Let `stream`{.variable} be null.

2.  If `object`{.variable} is a
    [`ReadableStream`{.idl}](https://streams.spec.whatwg.org/#readablestream){#ref-for-readablestream⑥
    link-type="idl"} object, then set `stream`{.variable} to
    `object`{.variable}.

3.  Otherwise, if `object`{.variable} is a
    [`Blob`{.idl}](https://w3c.github.io/FileAPI/#dfn-Blob){#ref-for-dfn-Blob④
    link-type="idl"} object, set `stream`{.variable} to the result of
    running `object`{.variable}'s [get
    stream](https://w3c.github.io/FileAPI/#blob-get-stream){#ref-for-blob-get-stream
    link-type="dfn"}.

4.  Otherwise, set `stream`{.variable} to a
    [new](https://webidl.spec.whatwg.org/#new){#ref-for-new①
    link-type="dfn"}
    [`ReadableStream`{.idl}](https://streams.spec.whatwg.org/#readablestream){#ref-for-readablestream⑦
    link-type="idl"} object, and [set
    up](https://streams.spec.whatwg.org/#readablestream-set-up-with-byte-reading-support){#ref-for-readablestream-set-up-with-byte-reading-support①
    link-type="dfn"} `stream`{.variable} with byte reading support.

5.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert②④
    link-type="dfn"}: `stream`{.variable} is a
    [`ReadableStream`{.idl}](https://streams.spec.whatwg.org/#readablestream){#ref-for-readablestream⑧
    link-type="idl"} object.

6.  Let `action`{.variable} be null.

7.  Let `source`{.variable} be null.

8.  Let `length`{.variable} be null.

9.  Let `type`{.variable} be null.

10. Switch on `object`{.variable}:

    [`Blob`{.idl}](https://w3c.github.io/FileAPI/#dfn-Blob){#ref-for-dfn-Blob⑤ link-type="idl"}

    :   Set `source`{.variable} to `object`{.variable}.

        Set `length`{.variable} to `object`{.variable}'s
        [`size`{.idl}](https://w3c.github.io/FileAPI/#dfn-size){#ref-for-dfn-size②
        link-type="idl"}.

        If `object`{.variable}'s
        [`type`{.idl}](https://w3c.github.io/FileAPI/#dfn-type){#ref-for-dfn-type①
        link-type="idl"} attribute is not the empty byte sequence, set
        `type`{.variable} to its value.

    [byte sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence②① link-type="dfn"}

    :   Set `source`{.variable} to `object`{.variable}.

    [`BufferSource`{.idl}](https://webidl.spec.whatwg.org/#BufferSource){#ref-for-BufferSource① link-type="idl"}

    :   Set `source`{.variable} to a [copy of the
        bytes](https://webidl.spec.whatwg.org/#dfn-get-buffer-source-copy){#ref-for-dfn-get-buffer-source-copy①
        link-type="dfn"} held by `object`{.variable}.

    [`FormData`{.idl}](https://xhr.spec.whatwg.org/#formdata){#ref-for-formdata② link-type="idl"}

    :   Set `action`{.variable} to this step: run the
        [`multipart/form-data` encoding
        algorithm](https://html.spec.whatwg.org/multipage/form-control-infrastructure.html#multipart%2Fform-data-encoding-algorithm){#ref-for-multipart%2Fform-data-encoding-algorithm
        link-type="dfn"}, with `object`{.variable}'s [entry
        list](https://xhr.spec.whatwg.org/#concept-formdata-entry-list){#ref-for-concept-formdata-entry-list
        link-type="dfn"} and
        [UTF-8](https://encoding.spec.whatwg.org/#utf-8){#ref-for-utf-8①
        link-type="dfn"}.

        Set `source`{.variable} to `object`{.variable}.

        Set `length`{.variable} to [unclear, see
        [html/6424](https://github.com/whatwg/html/issues/6424) for
        improving this]{.XXX}.

        Set `type`{.variable} to \``multipart/form-data; boundary=`\`,
        followed by the [`multipart/form-data` boundary
        string](https://html.spec.whatwg.org/multipage/form-control-infrastructure.html#multipart%2Fform-data-boundary-string){#ref-for-multipart%2Fform-data-boundary-string
        link-type="dfn"} generated by the [`multipart/form-data`
        encoding
        algorithm](https://html.spec.whatwg.org/multipage/form-control-infrastructure.html#multipart%2Fform-data-encoding-algorithm){#ref-for-multipart%2Fform-data-encoding-algorithm①
        link-type="dfn"}.

    [`URLSearchParams`{.idl}](https://url.spec.whatwg.org/#urlsearchparams){#ref-for-urlsearchparams① link-type="idl"}

    :   Set `source`{.variable} to the result of running the
        [`application/x-www-form-urlencoded`
        serializer](https://url.spec.whatwg.org/#concept-urlencoded-serializer){#ref-for-concept-urlencoded-serializer
        link-type="dfn"} with `object`{.variable}'s
        [list](https://url.spec.whatwg.org/#concept-urlsearchparams-list){#ref-for-concept-urlsearchparams-list
        link-type="dfn"}.

        Set `type`{.variable} to
        \``application/x-www-form-urlencoded;charset=UTF-8`\`.

    [scalar value string](https://infra.spec.whatwg.org/#scalar-value-string){#ref-for-scalar-value-string link-type="dfn"}

    :   Set `source`{.variable} to the [UTF-8
        encoding](https://encoding.spec.whatwg.org/#utf-8-encode){#ref-for-utf-8-encode
        link-type="dfn"} of `object`{.variable}.

        Set `type`{.variable} to \``text/plain;charset=UTF-8`\`.

    [`ReadableStream`{.idl}](https://streams.spec.whatwg.org/#readablestream){#ref-for-readablestream⑨ link-type="idl"}

    :   If `keepalive`{.variable} is true, then
        [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑤
        link-type="dfn"} a
        [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror⑦
        link-type="idl"}.

        If `object`{.variable} is
        [disturbed](https://streams.spec.whatwg.org/#is-readable-stream-disturbed){#ref-for-is-readable-stream-disturbed①
        link-type="dfn"} or
        [locked](https://streams.spec.whatwg.org/#readablestream-locked){#ref-for-readablestream-locked①
        link-type="dfn"}, then
        [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑥
        link-type="dfn"} a
        [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror⑧
        link-type="idl"}.

11. If `source`{.variable} is a [byte
    sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence②②
    link-type="dfn"}, then set `action`{.variable} to a step that
    returns `source`{.variable} and `length`{.variable} to
    `source`{.variable}'s
    [length](https://infra.spec.whatwg.org/#byte-sequence-length){#ref-for-byte-sequence-length⑤
    link-type="dfn"}.

12. If `action`{.variable} is non-null, then run these steps [in
    parallel](https://html.spec.whatwg.org/multipage/infrastructure.html#in-parallel){#ref-for-in-parallel⑦
    link-type="dfn"}:

    1.  Run `action`{.variable}.

        Whenever one or more bytes are available and `stream`{.variable}
        is not
        [errored](https://streams.spec.whatwg.org/#readablestream-errored){#ref-for-readablestream-errored①
        link-type="dfn"},
        [enqueue](https://streams.spec.whatwg.org/#readablestream-enqueue){#ref-for-readablestream-enqueue③
        link-type="dfn"} the result of
        [creating](https://webidl.spec.whatwg.org/#arraybufferview-create){#ref-for-arraybufferview-create
        link-type="dfn"} a
        [`Uint8Array`{.idl}](https://webidl.spec.whatwg.org/#idl-Uint8Array){#ref-for-idl-Uint8Array②
        link-type="idl"} from the available bytes into
        `stream`{.variable}.

        When running `action`{.variable} is done,
        [close](https://streams.spec.whatwg.org/#readablestream-close){#ref-for-readablestream-close①
        link-type="dfn"} `stream`{.variable}.

13. Let `body`{.variable} be a
    [body](#concept-body){#ref-for-concept-body①② link-type="dfn"} whose
    [stream](#concept-body-stream){#ref-for-concept-body-stream①④
    link-type="dfn"} is `stream`{.variable},
    [source](#concept-body-source){#ref-for-concept-body-source①③
    link-type="dfn"} is `source`{.variable}, and
    [length](#concept-body-total-bytes){#ref-for-concept-body-total-bytes③
    link-type="dfn"} is `length`{.variable}.

14. Return (`body`{.variable}, `type`{.variable}).

### [5.3. ]{.secno}[Body mixin]{.content}[](#body-mixin){.self-link} {#body-mixin .heading .settled level="5.3"}

``` {.idl .highlight .def}
interface mixin Body {
  readonly attribute ReadableStream? body;
  readonly attribute boolean bodyUsed;
  [NewObject] Promise<ArrayBuffer> arrayBuffer();
  [NewObject] Promise<Blob> blob();
  [NewObject] Promise<Uint8Array> bytes();
  [NewObject] Promise<FormData> formData();
  [NewObject] Promise<any> json();
  [NewObject] Promise<USVString> text();
};
```

Formats you would not want a network layer to be dependent upon, such as
HTML, will likely not be exposed here. Rather, an HTML parser API might
accept a stream in due course.

Objects including the [`Body`{.idl}](#body){#ref-for-body
link-type="idl"} interface mixin have an associated
[body]{#concept-body-body .dfn .dfn-paneled dfn-for="Body"
dfn-type="dfn" noexport=""} (null or a
[body](#concept-body){#ref-for-concept-body①③ link-type="dfn"}).

An object including the [`Body`{.idl}](#body){#ref-for-body①
link-type="idl"} interface mixin is said to be [unusable]{#body-unusable
.dfn .dfn-paneled dfn-for="Body" dfn-type="dfn" export=""} if its
[body](#concept-body-body){#ref-for-concept-body-body link-type="dfn"}
is non-null and its
[body](#concept-body-body){#ref-for-concept-body-body①
link-type="dfn"}'s
[stream](#concept-body-stream){#ref-for-concept-body-stream①⑤
link-type="dfn"} is
[disturbed](https://streams.spec.whatwg.org/#is-readable-stream-disturbed){#ref-for-is-readable-stream-disturbed②
link-type="dfn"} or
[locked](https://streams.spec.whatwg.org/#readablestream-locked){#ref-for-readablestream-locked②
link-type="dfn"}.

------------------------------------------------------------------------

`requestOrResponse`{.variable}` . `[`body`](#dom-body-body){#ref-for-dom-body-body① .idl-code link-type="attribute"}

:   Returns `requestOrResponse`{.variable}'s body as
    [`ReadableStream`{.idl}](https://streams.spec.whatwg.org/#readablestream){#ref-for-readablestream①①
    link-type="idl"}.

`requestOrResponse`{.variable}` . `[`bodyUsed`](#dom-body-bodyused){#ref-for-dom-body-bodyused① .idl-code link-type="attribute"}

:   Returns whether `requestOrResponse`{.variable}'s body has been read
    from.

`requestOrResponse`{.variable}` . `[`arrayBuffer`](#dom-body-arraybuffer){#ref-for-dom-body-arraybuffer① .idl-code link-type="method"}`()`

:   Returns a promise fulfilled with `requestOrResponse`{.variable}'s
    body as
    [`ArrayBuffer`{.idl}](https://webidl.spec.whatwg.org/#idl-ArrayBuffer){#ref-for-idl-ArrayBuffer①
    link-type="idl"}.

`requestOrResponse`{.variable}` . `[`blob`](#dom-body-blob){#ref-for-dom-body-blob① .idl-code link-type="method"}`()`

:   Returns a promise fulfilled with `requestOrResponse`{.variable}'s
    body as
    [`Blob`{.idl}](https://w3c.github.io/FileAPI/#dfn-Blob){#ref-for-dfn-Blob⑦
    link-type="idl"}.

`requestOrResponse`{.variable}` . `[`bytes`](#dom-body-bytes){#ref-for-dom-body-bytes① .idl-code link-type="method"}`()`

:   Returns a promise fulfilled with `requestOrResponse`{.variable}'s
    body as
    [`Uint8Array`{.idl}](https://webidl.spec.whatwg.org/#idl-Uint8Array){#ref-for-idl-Uint8Array④
    link-type="idl"}.

`requestOrResponse`{.variable}` . `[`formData`](#dom-body-formdata){#ref-for-dom-body-formdata① .idl-code link-type="method"}`()`

:   Returns a promise fulfilled with `requestOrResponse`{.variable}'s
    body as
    [`FormData`{.idl}](https://xhr.spec.whatwg.org/#formdata){#ref-for-formdata④
    link-type="idl"}.

`requestOrResponse`{.variable}` . `[`json`](#dom-body-json){#ref-for-dom-body-json① .idl-code link-type="method"}`()`

:   Returns a promise fulfilled with `requestOrResponse`{.variable}'s
    body parsed as JSON.

`requestOrResponse`{.variable}` . `[`text`](#dom-body-text){#ref-for-dom-body-text① .idl-code link-type="method"}`()`

:   Returns a promise fulfilled with `requestOrResponse`{.variable}'s
    body as string.

------------------------------------------------------------------------

::: {.algorithm algorithm="get the MIME type" algorithm-for="Body"}
To [get the MIME type]{#concept-body-mime-type .dfn .dfn-paneled
dfn-for="Body" dfn-type="dfn" noexport=""}, given a
[`Request`{.idl}](#request){#ref-for-request① link-type="idl"} or
[`Response`{.idl}](#response){#ref-for-response① link-type="idl"} object
`requestOrResponse`{.variable}:

1.  Let `headers`{.variable} be null.

2.  If `requestOrResponse`{.variable} is a
    [`Request`{.idl}](#request){#ref-for-request② link-type="idl"}
    object, then set `headers`{.variable} to
    `requestOrResponse`{.variable}'s
    [request](#concept-request-request){#ref-for-concept-request-request
    link-type="dfn"}'s [header
    list](#concept-request-header-list){#ref-for-concept-request-header-list④⑤
    link-type="dfn"}.

3.  Otherwise, set `headers`{.variable} to
    `requestOrResponse`{.variable}'s
    [response](#concept-response-response){#ref-for-concept-response-response
    link-type="dfn"}'s [header
    list](#concept-response-header-list){#ref-for-concept-response-header-list③②
    link-type="dfn"}.

4.  Let `mimeType`{.variable} be the result of [extracting a MIME
    type](#concept-header-extract-mime-type){#ref-for-concept-header-extract-mime-type⑧
    link-type="dfn"} from `headers`{.variable}.

5.  If `mimeType`{.variable} is failure, then return null.

6.  Return `mimeType`{.variable}.
:::

::: {.algorithm algorithm="body" algorithm-for="Body"}
The [`body`]{#dom-body-body .dfn .dfn-paneled .idl-code dfn-for="Body"
dfn-type="attribute" export=""} getter steps are to return null if
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑨
link-type="dfn"}'s
[body](#concept-body-body){#ref-for-concept-body-body② link-type="dfn"}
is null; otherwise
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⓪
link-type="dfn"}'s
[body](#concept-body-body){#ref-for-concept-body-body③
link-type="dfn"}'s
[stream](#concept-body-stream){#ref-for-concept-body-stream①⑥
link-type="dfn"}.
:::

::: {.algorithm algorithm="bodyUsed" algorithm-for="Body"}
The [`bodyUsed`]{#dom-body-bodyused .dfn .dfn-paneled .idl-code
dfn-for="Body" dfn-type="attribute" export=""} getter steps are to
return true if
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②①
link-type="dfn"}'s
[body](#concept-body-body){#ref-for-concept-body-body④ link-type="dfn"}
is non-null and
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②②
link-type="dfn"}'s
[body](#concept-body-body){#ref-for-concept-body-body⑤
link-type="dfn"}'s
[stream](#concept-body-stream){#ref-for-concept-body-stream①⑦
link-type="dfn"} is
[disturbed](https://streams.spec.whatwg.org/#is-readable-stream-disturbed){#ref-for-is-readable-stream-disturbed③
link-type="dfn"}; otherwise false.
:::

::: {.algorithm algorithm="consume body" algorithm-for="Body"}
The [consume body]{#concept-body-consume-body .dfn .dfn-paneled
dfn-for="Body" dfn-type="dfn" noexport=""} algorithm, given an object
that includes [`Body`{.idl}](#body){#ref-for-body② link-type="idl"}
`object`{.variable} and an algorithm that takes a [byte
sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence②③
link-type="dfn"} and returns a JavaScript value or throws an exception
`convertBytesToJSValue`{.variable}, runs these steps:

1.  If `object`{.variable} is
    [unusable](#body-unusable){#ref-for-body-unusable link-type="dfn"},
    then return [a promise rejected
    with](https://webidl.spec.whatwg.org/#a-promise-rejected-with){#ref-for-a-promise-rejected-with
    link-type="dfn"} a
    [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror⑨
    link-type="idl"}.

2.  Let `promise`{.variable} be [a new
    promise](https://webidl.spec.whatwg.org/#a-new-promise){#ref-for-a-new-promise①
    link-type="dfn"}.

3.  Let `errorSteps`{.variable} given `error`{.variable} be to
    [reject](https://webidl.spec.whatwg.org/#reject){#ref-for-reject
    link-type="dfn"} `promise`{.variable} with `error`{.variable}.

4.  Let `successSteps`{.variable} given a [byte
    sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence②④
    link-type="dfn"} `data`{.variable} be to
    [resolve](https://webidl.spec.whatwg.org/#resolve){#ref-for-resolve①
    link-type="dfn"} `promise`{.variable} with the result of running
    `convertBytesToJSValue`{.variable} with `data`{.variable}. If that
    threw an exception, then run `errorSteps`{.variable} with that
    exception.

5.  If `object`{.variable}'s
    [body](#concept-body-body){#ref-for-concept-body-body⑥
    link-type="dfn"} is null, then run `successSteps`{.variable} with an
    empty [byte
    sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence②⑤
    link-type="dfn"}.

6.  Otherwise, [fully read](#body-fully-read){#ref-for-body-fully-read②
    link-type="dfn"} `object`{.variable}'s
    [body](#concept-body-body){#ref-for-concept-body-body⑦
    link-type="dfn"} given `successSteps`{.variable},
    `errorSteps`{.variable}, and `object`{.variable}'s [relevant global
    object](https://html.spec.whatwg.org/multipage/webappapis.html#concept-relevant-global){#ref-for-concept-relevant-global
    link-type="dfn"}.

7.  Return `promise`{.variable}.
:::

::: {.algorithm algorithm="arrayBuffer()" algorithm-for="Body"}
The [`arrayBuffer()`]{#dom-body-arraybuffer .dfn .dfn-paneled .idl-code
dfn-for="Body" dfn-type="method" export=""} method steps are to return
the result of running [consume
body](#concept-body-consume-body){#ref-for-concept-body-consume-body
link-type="dfn"} with
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②③
link-type="dfn"} and the following step given a [byte
sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence②⑥
link-type="dfn"} `bytes`{.variable}: return the result of
[creating](https://webidl.spec.whatwg.org/#arraybuffer-create){#ref-for-arraybuffer-create
link-type="dfn"} an
[`ArrayBuffer`{.idl}](https://webidl.spec.whatwg.org/#idl-ArrayBuffer){#ref-for-idl-ArrayBuffer②
link-type="idl"} from `bytes`{.variable} in
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②④
link-type="dfn"}'s [relevant
realm](https://html.spec.whatwg.org/multipage/webappapis.html#concept-relevant-realm){#ref-for-concept-relevant-realm
link-type="dfn"}.

The above method can reject with a
[`RangeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-rangeerror){#ref-for-exceptiondef-rangeerror
link-type="idl"}.
:::

::: {.algorithm algorithm="blob()" algorithm-for="Body"}
The [`blob()`]{#dom-body-blob .dfn .dfn-paneled .idl-code dfn-for="Body"
dfn-type="method" export=""} method steps are to return the result of
running [consume
body](#concept-body-consume-body){#ref-for-concept-body-consume-body①
link-type="dfn"} with
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑤
link-type="dfn"} and the following step given a [byte
sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence②⑦
link-type="dfn"} `bytes`{.variable}: return a
[`Blob`{.idl}](https://w3c.github.io/FileAPI/#dfn-Blob){#ref-for-dfn-Blob⑧
link-type="idl"} whose contents are `bytes`{.variable} and whose
[`type`{.idl}](https://w3c.github.io/FileAPI/#dfn-type){#ref-for-dfn-type②
link-type="idl"} attribute is the result of [get the MIME
type](#concept-body-mime-type){#ref-for-concept-body-mime-type
link-type="dfn"} with
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑥
link-type="dfn"}.
:::

::: {.algorithm algorithm="bytes()" algorithm-for="Body"}
The [`bytes()`]{#dom-body-bytes .dfn .dfn-paneled .idl-code
dfn-for="Body" dfn-type="method" export=""} method steps are to return
the result of running [consume
body](#concept-body-consume-body){#ref-for-concept-body-consume-body②
link-type="dfn"} with
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑦
link-type="dfn"} and the following step given a [byte
sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence②⑧
link-type="dfn"} `bytes`{.variable}: return the result of
[creating](https://webidl.spec.whatwg.org/#arraybufferview-create){#ref-for-arraybufferview-create①
link-type="dfn"} a
[`Uint8Array`{.idl}](https://webidl.spec.whatwg.org/#idl-Uint8Array){#ref-for-idl-Uint8Array⑤
link-type="idl"} from `bytes`{.variable} in
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑧
link-type="dfn"}'s [relevant
realm](https://html.spec.whatwg.org/multipage/webappapis.html#concept-relevant-realm){#ref-for-concept-relevant-realm①
link-type="dfn"}.

The above method can reject with a
[`RangeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-rangeerror){#ref-for-exceptiondef-rangeerror①
link-type="idl"}.
:::

::: {.algorithm algorithm="formData()" algorithm-for="Body"}
The [`formData()`]{#dom-body-formdata .dfn .dfn-paneled .idl-code
dfn-for="Body" dfn-type="method" export=""} method steps are to return
the result of running [consume
body](#concept-body-consume-body){#ref-for-concept-body-consume-body③
link-type="dfn"} with
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑨
link-type="dfn"} and the following steps given a [byte
sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence②⑨
link-type="dfn"} `bytes`{.variable}:

1.  Let `mimeType`{.variable} be the result of [get the MIME
    type](#concept-body-mime-type){#ref-for-concept-body-mime-type①
    link-type="dfn"} with
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③⓪
    link-type="dfn"}.

2.  If `mimeType`{.variable} is non-null, then switch on
    `mimeType`{.variable}'s
    [essence](https://mimesniff.spec.whatwg.org/#mime-type-essence){#ref-for-mime-type-essence⑧
    link-type="dfn"} and run the corresponding steps:

    \"`multipart/form-data`\"

    :   1.  Parse `bytes`{.variable}, using the value of the
            \``boundary`\` parameter from `mimeType`{.variable}, per the
            rules set forth in Returning Values from Forms:
            multipart/form-data.
            [\[RFC7578\]](#biblio-rfc7578 "Returning Values from Forms: multipart/form-data"){link-type="biblio"}

            Each part whose \``Content-Disposition`\` header contains a
            \``filename`\` parameter must be parsed into an
            [entry](https://html.spec.whatwg.org/multipage/form-control-infrastructure.html#form-entry){#ref-for-form-entry
            link-type="dfn"} whose value is a
            [`File`{.idl}](https://w3c.github.io/FileAPI/#dfn-file){#ref-for-dfn-file
            link-type="idl"} object whose contents are the contents of
            the part. The
            [`name`{.idl}](https://w3c.github.io/FileAPI/#dfn-name){#ref-for-dfn-name
            link-type="idl"} attribute of the
            [`File`{.idl}](https://w3c.github.io/FileAPI/#dfn-file){#ref-for-dfn-file①
            link-type="idl"} object must have the value of the
            \``filename`\` parameter of the part. The
            [`type`{.idl}](https://w3c.github.io/FileAPI/#dfn-type){#ref-for-dfn-type③
            link-type="idl"} attribute of the
            [`File`{.idl}](https://w3c.github.io/FileAPI/#dfn-file){#ref-for-dfn-file②
            link-type="idl"} object must have the value of the
            \``Content-Type`\` header of the part if the part has such
            header, and \``text/plain`\` (the default defined by
            [\[RFC7578\]](#biblio-rfc7578 "Returning Values from Forms: multipart/form-data"){link-type="biblio"}
            section 4.4) otherwise.

            Each part whose \``Content-Disposition`\` header does not
            contain a \``filename`\` parameter must be parsed into an
            [entry](https://html.spec.whatwg.org/multipage/form-control-infrastructure.html#form-entry){#ref-for-form-entry①
            link-type="dfn"} whose value is the [UTF-8 decoded without
            BOM](https://encoding.spec.whatwg.org/#utf-8-decode-without-bom){#ref-for-utf-8-decode-without-bom
            link-type="dfn"} content of the part. [This is done
            regardless of the presence or the value of a
            \``Content-Type`\` header and regardless of the presence or
            the value of a \``charset`\` parameter.]{.note}

            A part whose \``Content-Disposition`\` header contains a
            \``name`\` parameter whose value is \``_charset_`\` is
            parsed like any other part. It does not change the encoding.

        2.  If that fails for some reason, then
            [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑦
            link-type="dfn"} a
            [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror①⓪
            link-type="idl"}.

        3.  Return a new
            [`FormData`{.idl}](https://xhr.spec.whatwg.org/#formdata){#ref-for-formdata⑤
            link-type="idl"} object, appending each
            [entry](https://html.spec.whatwg.org/multipage/form-control-infrastructure.html#form-entry){#ref-for-form-entry②
            link-type="dfn"}, resulting from the parsing operation, to
            its [entry
            list](https://xhr.spec.whatwg.org/#concept-formdata-entry-list){#ref-for-concept-formdata-entry-list①
            link-type="dfn"}.

        The above is a rough approximation of what is needed for
        \``multipart/form-data`\`, a more detailed parsing specification
        is to be written. Volunteers welcome.

    \"`application/x-www-form-urlencoded`\"

    :   1.  Let `entries`{.variable} be the result of
            [parsing](https://url.spec.whatwg.org/#concept-urlencoded-parser){#ref-for-concept-urlencoded-parser
            link-type="dfn"} `bytes`{.variable}.

        2.  Return a new
            [`FormData`{.idl}](https://xhr.spec.whatwg.org/#formdata){#ref-for-formdata⑥
            link-type="idl"} object whose [entry
            list](https://xhr.spec.whatwg.org/#concept-formdata-entry-list){#ref-for-concept-formdata-entry-list②
            link-type="dfn"} is `entries`{.variable}.

3.  [Throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑧
    link-type="dfn"} a
    [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror①①
    link-type="idl"}.
:::

::: {.algorithm algorithm="json()" algorithm-for="Body"}
The [`json()`]{#dom-body-json .dfn .dfn-paneled .idl-code dfn-for="Body"
dfn-type="method" export=""} method steps are to return the result of
running [consume
body](#concept-body-consume-body){#ref-for-concept-body-consume-body④
link-type="dfn"} with
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this③①
link-type="dfn"} and [parse JSON from
bytes](https://infra.spec.whatwg.org/#parse-json-bytes-to-a-javascript-value){#ref-for-parse-json-bytes-to-a-javascript-value
link-type="dfn"}.

The above method can reject with a
[`SyntaxError`{.idl}](https://webidl.spec.whatwg.org/#syntaxerror){#ref-for-syntaxerror
link-type="idl"}.
:::

::: {.algorithm algorithm="text()" algorithm-for="Body"}
The [`text()`]{#dom-body-text .dfn .dfn-paneled .idl-code dfn-for="Body"
dfn-type="method" export=""} method steps are to return the result of
running [consume
body](#concept-body-consume-body){#ref-for-concept-body-consume-body⑤
link-type="dfn"} with
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this③②
link-type="dfn"} and [UTF-8
decode](https://encoding.spec.whatwg.org/#utf-8-decode){#ref-for-utf-8-decode
link-type="dfn"}.
:::

### [5.4. ]{.secno}[Request class]{.content}[](#request-class){.self-link} {#request-class .heading .settled level="5.4"}

``` {.idl .highlight .def}
typedef (Request or USVString) RequestInfo;

[Exposed=(Window,Worker)]
interface Request {
  constructor(RequestInfo input, optional RequestInit init = {});

  readonly attribute ByteString method;
  readonly attribute USVString url;
  [SameObject] readonly attribute Headers headers;

  readonly attribute RequestDestination destination;
  readonly attribute USVString referrer;
  readonly attribute ReferrerPolicy referrerPolicy;
  readonly attribute RequestMode mode;
  readonly attribute RequestCredentials credentials;
  readonly attribute RequestCache cache;
  readonly attribute RequestRedirect redirect;
  readonly attribute DOMString integrity;
  readonly attribute boolean keepalive;
  readonly attribute boolean isReloadNavigation;
  readonly attribute boolean isHistoryNavigation;
  readonly attribute AbortSignal signal;
  readonly attribute RequestDuplex duplex;

  [NewObject] Request clone();
};
Request includes Body;

dictionary RequestInit {
  ByteString method;
  HeadersInit headers;
  BodyInit? body;
  USVString referrer;
  ReferrerPolicy referrerPolicy;
  RequestMode mode;
  RequestCredentials credentials;
  RequestCache cache;
  RequestRedirect redirect;
  DOMString integrity;
  boolean keepalive;
  AbortSignal? signal;
  RequestDuplex duplex;
  RequestPriority priority;
  any window; // can only be set to null
};

enum RequestDestination { "", "audio", "audioworklet", "document", "embed", "font", "frame", "iframe", "image", "json", "manifest", "object", "paintworklet", "report", "script", "sharedworker", "style",  "track", "video", "worker", "xslt" };
enum RequestMode { "navigate", "same-origin", "no-cors", "cors" };
enum RequestCredentials { "omit", "same-origin", "include" };
enum RequestCache { "default", "no-store", "reload", "no-cache", "force-cache", "only-if-cached" };
enum RequestRedirect { "follow", "error", "manual" };
enum RequestDuplex { "half" };
enum RequestPriority { "high", "low", "auto" };
```

\"`serviceworker`\" is omitted from
[`RequestDestination`](#requestdestination){#ref-for-requestdestination②
.idl-code link-type="enum"} as it cannot be observed from JavaScript.
Implementations will still need to support it as a
[destination](#concept-request-destination){#ref-for-concept-request-destination②⓪
link-type="dfn"}. \"`websocket`\" is omitted from
[`RequestMode`](#requestmode){#ref-for-requestmode② .idl-code
link-type="enum"} as it cannot be used nor observed from JavaScript.

A [`Request`{.idl}](#request){#ref-for-request⑥ link-type="idl"} object
has an associated [request]{#concept-request-request .dfn .dfn-paneled
dfn-for="Request" dfn-type="dfn" export=""} (a
[request](#concept-request){#ref-for-concept-request①①⓪
link-type="dfn"}).

A [`Request`{.idl}](#request){#ref-for-request⑦ link-type="idl"} object
also has an associated [headers]{#request-headers .dfn .dfn-paneled
dfn-for="Request" dfn-type="dfn" export=""} (null or a
[`Headers`{.idl}](#headers){#ref-for-headers⑨ link-type="idl"} object),
initially null.

A [`Request`{.idl}](#request){#ref-for-request⑧ link-type="idl"} object
has an associated [signal]{#request-signal .dfn .dfn-paneled
dfn-for="Request" dfn-type="dfn" noexport=""} (null or an
[`AbortSignal`{.idl}](https://dom.spec.whatwg.org/#abortsignal){#ref-for-abortsignal②
link-type="idl"} object), initially null.

A [`Request`{.idl}](#request){#ref-for-request⑨ link-type="idl"}
object's [body](#concept-body-body){#ref-for-concept-body-body⑧
link-type="dfn"} is its
[request](#concept-request-request){#ref-for-concept-request-request①
link-type="dfn"}'s
[body](#concept-request-body){#ref-for-concept-request-body④④
link-type="dfn"}.

------------------------------------------------------------------------

`request`{.variable}` = new `[`Request`](#dom-request){#ref-for-dom-request① .idl-code link-type="constructor"}`(``input`{.variable}` [, ``init`{.variable}`])`

:   Returns a new `request`{.variable} whose
    [`url`{.idl}](#dom-request-url){#ref-for-dom-request-url①
    link-type="idl"} property is `input`{.variable} if
    `input`{.variable} is a string, and `input`{.variable}'s
    [`url`{.idl}](#dom-request-url){#ref-for-dom-request-url②
    link-type="idl"} if `input`{.variable} is a
    [`Request`{.idl}](#request){#ref-for-request①⓪ link-type="idl"}
    object.

    The `init`{.variable} argument is an object whose properties can be
    set as follows:

    [`method`{.idl}](#dom-requestinit-method){#ref-for-dom-requestinit-method link-type="idl"}
    :   A string to set `request`{.variable}'s
        [`method`{.idl}](#dom-request-method){#ref-for-dom-request-method①
        link-type="idl"}.

    [`headers`{.idl}](#dom-requestinit-headers){#ref-for-dom-requestinit-headers link-type="idl"}
    :   A [`Headers`{.idl}](#headers){#ref-for-headers①⓪
        link-type="idl"} object, an object literal, or an array of
        two-item arrays to set `request`{.variable}'s
        [`headers`{.idl}](#dom-request-headers){#ref-for-dom-request-headers①
        link-type="idl"}.

    [`body`{.idl}](#dom-requestinit-body){#ref-for-dom-requestinit-body link-type="idl"}
    :   A [`BodyInit`{.idl}](#bodyinit){#ref-for-bodyinit③
        link-type="idl"} object or null to set `request`{.variable}'s
        [body](#concept-request-body){#ref-for-concept-request-body④⑤
        link-type="dfn"}.

    [`referrer`{.idl}](#dom-requestinit-referrer){#ref-for-dom-requestinit-referrer link-type="idl"}
    :   A string whose value is a same-origin URL, \"`about:client`\",
        or the empty string, to set `request`{.variable}'s
        [referrer](#concept-request-referrer){#ref-for-concept-request-referrer⑧
        link-type="dfn"}.

    [`referrerPolicy`{.idl}](#dom-requestinit-referrerpolicy){#ref-for-dom-requestinit-referrerpolicy link-type="idl"}
    :   A [referrer
        policy](https://w3c.github.io/webappsec-referrer-policy/#referrer-policy){#ref-for-referrer-policy①
        link-type="dfn"} to set `request`{.variable}'s
        [`referrerPolicy`{.idl}](#dom-request-referrerpolicy){#ref-for-dom-request-referrerpolicy①
        link-type="idl"}.

    [`mode`{.idl}](#dom-requestinit-mode){#ref-for-dom-requestinit-mode link-type="idl"}
    :   A string to indicate whether the request will use CORS, or will
        be restricted to same-origin URLs. Sets `request`{.variable}'s
        [`mode`{.idl}](#dom-request-mode){#ref-for-dom-request-mode①
        link-type="idl"}. If `input`{.variable} is a string, it defaults
        to \"`cors`\".

    [`credentials`{.idl}](#dom-requestinit-credentials){#ref-for-dom-requestinit-credentials link-type="idl"}
    :   A string indicating whether credentials will be sent with the
        request always, never, or only when sent to a same-origin URL
        --- as well as whether any credentials sent back in the response
        will be used always, never, or only when received from a
        same-origin URL. Sets `request`{.variable}'s
        [`credentials`{.idl}](#dom-request-credentials){#ref-for-dom-request-credentials①
        link-type="idl"}. If `input`{.variable} is a string, it defaults
        to \"`same-origin`\".

    [`cache`{.idl}](#dom-requestinit-cache){#ref-for-dom-requestinit-cache link-type="idl"}
    :   A string indicating how the request will interact with the
        browser's cache to set `request`{.variable}'s
        [`cache`{.idl}](#dom-request-cache){#ref-for-dom-request-cache①
        link-type="idl"}.

    [`redirect`{.idl}](#dom-requestinit-redirect){#ref-for-dom-requestinit-redirect link-type="idl"}
    :   A string indicating whether `request`{.variable} follows
        redirects, results in an error upon encountering a redirect, or
        returns the redirect (in an opaque fashion). Sets
        `request`{.variable}'s
        [`redirect`{.idl}](#dom-request-redirect){#ref-for-dom-request-redirect①
        link-type="idl"}.

    [`integrity`{.idl}](#dom-requestinit-integrity){#ref-for-dom-requestinit-integrity link-type="idl"}
    :   A cryptographic hash of the resource to be fetched by
        `request`{.variable}. Sets `request`{.variable}'s
        [`integrity`{.idl}](#dom-request-integrity){#ref-for-dom-request-integrity①
        link-type="idl"}.

    [`keepalive`{.idl}](#dom-requestinit-keepalive){#ref-for-dom-requestinit-keepalive link-type="idl"}
    :   A boolean to set `request`{.variable}'s
        [`keepalive`{.idl}](#dom-request-keepalive){#ref-for-dom-request-keepalive①
        link-type="idl"}.

    [`signal`{.idl}](#dom-requestinit-signal){#ref-for-dom-requestinit-signal link-type="idl"}
    :   An
        [`AbortSignal`{.idl}](https://dom.spec.whatwg.org/#abortsignal){#ref-for-abortsignal③
        link-type="idl"} to set `request`{.variable}'s
        [`signal`{.idl}](#dom-request-signal){#ref-for-dom-request-signal①
        link-type="idl"}.

    [`window`{.idl}](#dom-requestinit-window){#ref-for-dom-requestinit-window link-type="idl"}
    :   Can only be null. Used to disassociate `request`{.variable} from
        any
        [`Window`{.idl}](https://html.spec.whatwg.org/multipage/nav-history-apis.html#window){#ref-for-window②
        link-type="idl"}.

    [`duplex`{.idl}](#dom-requestinit-duplex){#ref-for-dom-requestinit-duplex link-type="idl"}
    :   \"`half`\" is the only valid value and it is for initiating a
        half-duplex fetch (i.e., the user agent sends the entire request
        before processing the response). \"`full`\" is reserved for
        future use, for initiating a full-duplex fetch (i.e., the user
        agent can process the response before sending the entire
        request). This member needs to be set when
        [`body`{.idl}](#dom-requestinit-body){#ref-for-dom-requestinit-body①
        link-type="idl"} is a
        [`ReadableStream`{.idl}](https://streams.spec.whatwg.org/#readablestream){#ref-for-readablestream①②
        link-type="idl"} object. [See [issue
        #1254](https://github.com/whatwg/fetch/issues/1254) for defining
        \"`full`\".]{.note}

    [`priority`{.idl}](#dom-requestinit-priority){#ref-for-dom-requestinit-priority link-type="idl"}
    :   A string to set `request`{.variable}'s
        [priority](#request-priority){#ref-for-request-priority①
        link-type="dfn"}.

`request`{.variable}` . `[`method`](#dom-request-method){#ref-for-dom-request-method② .idl-code link-type="attribute"}
:   Returns `request`{.variable}'s HTTP method, which is \"`GET`\" by
    default.

`request`{.variable}` . `[`url`](#dom-request-url){#ref-for-dom-request-url③ .idl-code link-type="attribute"}
:   Returns the URL of `request`{.variable} as a string.

`request`{.variable}` . `[`headers`](#dom-request-headers){#ref-for-dom-request-headers② .idl-code link-type="attribute"}
:   Returns a [`Headers`{.idl}](#headers){#ref-for-headers①①
    link-type="idl"} object consisting of the headers associated with
    `request`{.variable}. Note that headers added in the network layer
    by the user agent will not be accounted for in this object, e.g.,
    the \"`Host`\" header.

`request`{.variable}` . `[`destination`](#dom-request-destination){#ref-for-dom-request-destination① .idl-code link-type="attribute"}
:   Returns the kind of resource requested by `request`{.variable},
    e.g., \"`document`\" or \"`script`\".

`request`{.variable}` . `[`referrer`](#dom-request-referrer){#ref-for-dom-request-referrer① .idl-code link-type="attribute"}
:   Returns the referrer of `request`{.variable}. Its value can be a
    same-origin URL if explicitly set in `init`{.variable}, the empty
    string to indicate no referrer, and \"`about:client`\" when
    defaulting to the global's default. This is used during fetching to
    determine the value of the \``Referer`\` header of the request being
    made.

`request`{.variable}` . `[`referrerPolicy`](#dom-request-referrerpolicy){#ref-for-dom-request-referrerpolicy② .idl-code link-type="attribute"}
:   Returns the referrer policy associated with `request`{.variable}.
    This is used during fetching to compute the value of the
    `request`{.variable}'s referrer.

`request`{.variable}` . `[`mode`](#dom-request-mode){#ref-for-dom-request-mode② .idl-code link-type="attribute"}
:   Returns the
    [mode](#concept-request-mode){#ref-for-concept-request-mode②④
    link-type="dfn"} associated with `request`{.variable}, which is a
    string indicating whether the request will use CORS, or will be
    restricted to same-origin URLs.

`request`{.variable}` . `[`credentials`](#dom-request-credentials){#ref-for-dom-request-credentials② .idl-code link-type="attribute"}
:   Returns the [credentials
    mode](#concept-request-credentials-mode){#ref-for-concept-request-credentials-mode②⓪
    link-type="dfn"} associated with `request`{.variable}, which is a
    string indicating whether credentials will be sent with the request
    always, never, or only when sent to a same-origin URL.

`request`{.variable}` . `[`cache`](#dom-request-cache){#ref-for-dom-request-cache② .idl-code link-type="attribute"}
:   Returns the [cache
    mode](#concept-request-cache-mode){#ref-for-concept-request-cache-mode①①
    link-type="dfn"} associated with `request`{.variable}, which is a
    string indicating how the request will interact with the browser's
    cache when fetching.

`request`{.variable}` . `[`redirect`](#dom-request-redirect){#ref-for-dom-request-redirect② .idl-code link-type="attribute"}
:   Returns the [redirect
    mode](#concept-request-redirect-mode){#ref-for-concept-request-redirect-mode⑧
    link-type="dfn"} associated with `request`{.variable}, which is a
    string indicating how redirects for the request will be handled
    during fetching. A
    [request](#concept-request){#ref-for-concept-request①①①
    link-type="dfn"} will follow redirects by default.

`request`{.variable}` . `[`integrity`](#dom-request-integrity){#ref-for-dom-request-integrity② .idl-code link-type="attribute"}
:   Returns `request`{.variable}'s subresource integrity metadata, which
    is a cryptographic hash of the resource being fetched. Its value
    consists of multiple hashes separated by whitespace.
    [\[SRI\]](#biblio-sri "Subresource Integrity"){link-type="biblio"}

`request`{.variable}` . `[`keepalive`](#dom-request-keepalive){#ref-for-dom-request-keepalive② .idl-code link-type="attribute"}
:   Returns a boolean indicating whether or not `request`{.variable} can
    outlive the global in which it was created.

`request`{.variable}` . `[`isReloadNavigation`](#dom-request-isreloadnavigation){#ref-for-dom-request-isreloadnavigation① .idl-code link-type="attribute"}
:   Returns a boolean indicating whether or not `request`{.variable} is
    for a reload navigation.

`request`{.variable}` . `[`isHistoryNavigation`](#dom-request-ishistorynavigation){#ref-for-dom-request-ishistorynavigation① .idl-code link-type="attribute"}
:   Returns a boolean indicating whether or not `request`{.variable} is
    for a history navigation (a.k.a. back-foward navigation).

`request`{.variable}` . `[`signal`](#dom-request-signal){#ref-for-dom-request-signal② .idl-code link-type="attribute"}
:   Returns the signal associated with `request`{.variable}, which is an
    [`AbortSignal`{.idl}](https://dom.spec.whatwg.org/#abortsignal){#ref-for-abortsignal④
    link-type="idl"} object indicating whether or not
    `request`{.variable} has been aborted, and its abort event handler.

`request`{.variable}` . `[`duplex`](#dom-request-duplex){#ref-for-dom-request-duplex① .idl-code link-type="attribute"}
:   Returns \"`half`\", meaning the fetch will be half-duplex (i.e., the
    user agent sends the entire request before processing the response).
    In future, it could also return \"`full`\", meaning the fetch will
    be full-duplex (i.e., the user agent can process the response before
    sending the entire request) to indicate that the fetch will be
    full-duplex. [See [issue
    #1254](https://github.com/whatwg/fetch/issues/1254) for defining
    \"`full`\".]{.note}

`request`{.variable}` . `[`clone`](#dom-request-clone){#ref-for-dom-request-clone① .idl-code link-type="method"}`()`

:   Returns a clone of `request`{.variable}.

------------------------------------------------------------------------

::: {.algorithm algorithm="create" algorithm-for="Request"}
To [create]{#request-create .dfn .dfn-paneled dfn-for="Request"
dfn-type="dfn" export="" lt="create|creating"} a
[`Request`{.idl}](#request){#ref-for-request①① link-type="idl"} object,
given a [request](#concept-request){#ref-for-concept-request①①②
link-type="dfn"} `request`{.variable}, [headers
guard](#headers-guard){#ref-for-headers-guard① link-type="dfn"}
`guard`{.variable},
[`AbortSignal`{.idl}](https://dom.spec.whatwg.org/#abortsignal){#ref-for-abortsignal⑤
link-type="idl"} object `signal`{.variable}, and
[realm](https://tc39.es/ecma262/#realm){#ref-for-realm② link-type="dfn"}
`realm`{.variable}:

1.  Let `requestObject`{.variable} be a
    [new](https://webidl.spec.whatwg.org/#new){#ref-for-new②
    link-type="dfn"} [`Request`{.idl}](#request){#ref-for-request①②
    link-type="idl"} object with `realm`{.variable}.

2.  Set `requestObject`{.variable}'s
    [request](#concept-request-request){#ref-for-concept-request-request②
    link-type="dfn"} to `request`{.variable}.

3.  Set `requestObject`{.variable}'s
    [headers](#request-headers){#ref-for-request-headers
    link-type="dfn"} to a
    [new](https://webidl.spec.whatwg.org/#new){#ref-for-new③
    link-type="dfn"} [`Headers`{.idl}](#headers){#ref-for-headers①②
    link-type="idl"} object with `realm`{.variable}, whose [headers
    list](#concept-headers-header-list){#ref-for-concept-headers-header-list①①
    link-type="dfn"} is `request`{.variable}'s [headers
    list](#concept-request-header-list){#ref-for-concept-request-header-list④⑥
    link-type="dfn"} and
    [guard](#concept-headers-guard){#ref-for-concept-headers-guard①⓪
    link-type="dfn"} is `guard`{.variable}.

4.  Set `requestObject`{.variable}'s
    [signal](#request-signal){#ref-for-request-signal link-type="dfn"}
    to `signal`{.variable}.

5.  Return `requestObject`{.variable}.
:::

------------------------------------------------------------------------

::: {.algorithm algorithm="Request(input, init)" algorithm-for="Request"}
The
[`new Request(``input`{.variable}`, ``init`{.variable}`)`]{#dom-request
.dfn .dfn-paneled .idl-code dfn-for="Request" dfn-type="constructor"
export=""
lt="Request(input, init)|constructor(input, init)|Request(input)|constructor(input)"}
constructor steps are:

1.  Let `request`{.variable} be null.

2.  Let `fallbackMode`{.variable} be null.

3.  Let `baseURL`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③③
    link-type="dfn"}'s [relevant settings
    object](https://html.spec.whatwg.org/multipage/webappapis.html#relevant-settings-object){#ref-for-relevant-settings-object
    link-type="dfn"}'s [API base
    URL](https://html.spec.whatwg.org/multipage/webappapis.html#api-base-url){#ref-for-api-base-url
    link-type="dfn"}.

4.  Let `signal`{.variable} be null.

5.  If `input`{.variable} is a string, then:

    1.  Let `parsedURL`{.variable} be the result of
        [parsing](https://url.spec.whatwg.org/#concept-url-parser){#ref-for-concept-url-parser①
        link-type="dfn"} `input`{.variable} with `baseURL`{.variable}.

    2.  If `parsedURL`{.variable} is failure, then
        [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑨
        link-type="dfn"} a
        [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror①②
        link-type="idl"}.

    3.  If `parsedURL`{.variable} [includes
        credentials](https://url.spec.whatwg.org/#include-credentials){#ref-for-include-credentials④
        link-type="dfn"}, then
        [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw①⓪
        link-type="dfn"} a
        [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror①③
        link-type="idl"}.

    4.  Set `request`{.variable} to a new
        [request](#concept-request){#ref-for-concept-request①①③
        link-type="dfn"} whose
        [URL](#concept-request-url){#ref-for-concept-request-url⑥
        link-type="dfn"} is `parsedURL`{.variable}.

    5.  Set `fallbackMode`{.variable} to \"`cors`\".

6.  Otherwise:

    1.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert②⑤
        link-type="dfn"}: `input`{.variable} is a
        [`Request`{.idl}](#request){#ref-for-request①③ link-type="idl"}
        object.

    2.  Set `request`{.variable} to `input`{.variable}'s
        [request](#concept-request-request){#ref-for-concept-request-request③
        link-type="dfn"}.

    3.  Set `signal`{.variable} to `input`{.variable}'s
        [signal](#request-signal){#ref-for-request-signal①
        link-type="dfn"}.

7.  Let `origin`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③④
    link-type="dfn"}'s [relevant settings
    object](https://html.spec.whatwg.org/multipage/webappapis.html#relevant-settings-object){#ref-for-relevant-settings-object①
    link-type="dfn"}'s
    [origin](https://html.spec.whatwg.org/multipage/webappapis.html#concept-settings-object-origin){#ref-for-concept-settings-object-origin②
    link-type="dfn"}.

8.  Let `traversableForUserPrompts`{.variable} be \"`client`\".

9.  If `request`{.variable}'s [traversable for user
    prompts](#concept-request-window){#ref-for-concept-request-window①③
    link-type="dfn"} is an [environment settings
    object](https://html.spec.whatwg.org/multipage/webappapis.html#environment-settings-object){#ref-for-environment-settings-object⑨
    link-type="dfn"} and its
    [origin](https://html.spec.whatwg.org/multipage/webappapis.html#concept-settings-object-origin){#ref-for-concept-settings-object-origin③
    link-type="dfn"} is [same
    origin](https://html.spec.whatwg.org/multipage/browsers.html#same-origin){#ref-for-same-origin①⓪
    link-type="dfn"} with `origin`{.variable}, then set
    `traversableForUserPrompts`{.variable} to `request`{.variable}'s
    [traversable for user
    prompts](#concept-request-window){#ref-for-concept-request-window①④
    link-type="dfn"}.

10. If
    `init`{.variable}\[\"[`window`{.idl}](#dom-requestinit-window){#ref-for-dom-requestinit-window①
    link-type="idl"}\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists③
    link-type="dfn"} and is non-null, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw①①
    link-type="dfn"} a
    [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror①④
    link-type="idl"}.

11. If
    `init`{.variable}\[\"[`window`{.idl}](#dom-requestinit-window){#ref-for-dom-requestinit-window②
    link-type="idl"}\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists④
    link-type="dfn"}, then set `traversableForUserPrompts`{.variable} to
    \"`no-traversable`\".

12. Set `request`{.variable} to a new
    [request](#concept-request){#ref-for-concept-request①①④
    link-type="dfn"} with the following properties:

    [URL](#concept-request-url){#ref-for-concept-request-url⑦ link-type="dfn"}
    :   `request`{.variable}'s
        [URL](#concept-request-url){#ref-for-concept-request-url⑧
        link-type="dfn"}.

    [method](#concept-request-method){#ref-for-concept-request-method②⓪ link-type="dfn"}
    :   `request`{.variable}'s
        [method](#concept-request-method){#ref-for-concept-request-method②①
        link-type="dfn"}.

    [header list](#concept-request-header-list){#ref-for-concept-request-header-list④⑦ link-type="dfn"}
    :   A copy of `request`{.variable}'s [header
        list](#concept-request-header-list){#ref-for-concept-request-header-list④⑧
        link-type="dfn"}.

    [unsafe-request flag](#unsafe-request-flag){#ref-for-unsafe-request-flag③ link-type="dfn"}
    :   Set.

    [client](#concept-request-client){#ref-for-concept-request-client②⑨ link-type="dfn"}
    :   [This](https://webidl.spec.whatwg.org/#this){#ref-for-this③⑤
        link-type="dfn"}'s [relevant settings
        object](https://html.spec.whatwg.org/multipage/webappapis.html#relevant-settings-object){#ref-for-relevant-settings-object②
        link-type="dfn"}.

    [traversable for user prompts](#concept-request-window){#ref-for-concept-request-window①⑤ link-type="dfn"}
    :   `traversableForUserPrompts`{.variable}.

    [internal priority](#request-internal-priority){#ref-for-request-internal-priority② link-type="dfn"}
    :   `request`{.variable}'s [internal
        priority](#request-internal-priority){#ref-for-request-internal-priority③
        link-type="dfn"}.

    [origin](#concept-request-origin){#ref-for-concept-request-origin②② link-type="dfn"}
    :   `request`{.variable}'s
        [origin](#concept-request-origin){#ref-for-concept-request-origin②③
        link-type="dfn"}. [The propagation of the
        [origin](https://html.spec.whatwg.org/multipage/browsers.html#concept-origin){#ref-for-concept-origin⑧
        link-type="dfn"} is only significant for navigation requests
        being handled by a service worker. In this scenario a request
        can have an origin that is different from the current
        client.]{.note}

    [referrer](#concept-request-referrer){#ref-for-concept-request-referrer⑨ link-type="dfn"}
    :   `request`{.variable}'s
        [referrer](#concept-request-referrer){#ref-for-concept-request-referrer①⓪
        link-type="dfn"}.

    [referrer policy](#concept-request-referrer-policy){#ref-for-concept-request-referrer-policy⑥ link-type="dfn"}
    :   `request`{.variable}'s [referrer
        policy](#concept-request-referrer-policy){#ref-for-concept-request-referrer-policy⑦
        link-type="dfn"}.

    [mode](#concept-request-mode){#ref-for-concept-request-mode②⑤ link-type="dfn"}
    :   `request`{.variable}'s
        [mode](#concept-request-mode){#ref-for-concept-request-mode②⑥
        link-type="dfn"}.

    [credentials mode](#concept-request-credentials-mode){#ref-for-concept-request-credentials-mode②① link-type="dfn"}
    :   `request`{.variable}'s [credentials
        mode](#concept-request-credentials-mode){#ref-for-concept-request-credentials-mode②②
        link-type="dfn"}.

    [cache mode](#concept-request-cache-mode){#ref-for-concept-request-cache-mode①② link-type="dfn"}
    :   `request`{.variable}'s [cache
        mode](#concept-request-cache-mode){#ref-for-concept-request-cache-mode①③
        link-type="dfn"}.

    [redirect mode](#concept-request-redirect-mode){#ref-for-concept-request-redirect-mode⑨ link-type="dfn"}
    :   `request`{.variable}'s [redirect
        mode](#concept-request-redirect-mode){#ref-for-concept-request-redirect-mode①⓪
        link-type="dfn"}.

    [integrity metadata](#concept-request-integrity-metadata){#ref-for-concept-request-integrity-metadata③ link-type="dfn"}
    :   `request`{.variable}'s [integrity
        metadata](#concept-request-integrity-metadata){#ref-for-concept-request-integrity-metadata④
        link-type="dfn"}.

    [keepalive](#request-keepalive-flag){#ref-for-request-keepalive-flag③ link-type="dfn"}
    :   `request`{.variable}'s
        [keepalive](#request-keepalive-flag){#ref-for-request-keepalive-flag④
        link-type="dfn"}.

    [reload-navigation flag](#concept-request-reload-navigation-flag){#ref-for-concept-request-reload-navigation-flag link-type="dfn"}
    :   `request`{.variable}'s [reload-navigation
        flag](#concept-request-reload-navigation-flag){#ref-for-concept-request-reload-navigation-flag①
        link-type="dfn"}.

    [history-navigation flag](#concept-request-history-navigation-flag){#ref-for-concept-request-history-navigation-flag link-type="dfn"}
    :   `request`{.variable}'s [history-navigation
        flag](#concept-request-history-navigation-flag){#ref-for-concept-request-history-navigation-flag①
        link-type="dfn"}.

    [URL list](#concept-request-url-list){#ref-for-concept-request-url-list⑨ link-type="dfn"}
    :   A
        [clone](https://infra.spec.whatwg.org/#list-clone){#ref-for-list-clone③
        link-type="dfn"} of `request`{.variable}'s [URL
        list](#concept-request-url-list){#ref-for-concept-request-url-list①⓪
        link-type="dfn"}.

    [initiator type](#request-initiator-type){#ref-for-request-initiator-type③ link-type="dfn"}
    :   \"`fetch`\".

13. If `init`{.variable} [is not
    empty](https://infra.spec.whatwg.org/#map-is-empty){#ref-for-map-is-empty
    link-type="dfn"}, then:

    1.  If `request`{.variable}'s
        [mode](#concept-request-mode){#ref-for-concept-request-mode②⑦
        link-type="dfn"} is \"`navigate`\", then set it to
        \"`same-origin`\".

    2.  Unset `request`{.variable}'s [reload-navigation
        flag](#concept-request-reload-navigation-flag){#ref-for-concept-request-reload-navigation-flag②
        link-type="dfn"}.

    3.  Unset `request`{.variable}'s [history-navigation
        flag](#concept-request-history-navigation-flag){#ref-for-concept-request-history-navigation-flag②
        link-type="dfn"}.

    4.  Set `request`{.variable}'s
        [origin](#concept-request-origin){#ref-for-concept-request-origin②④
        link-type="dfn"} to \"`client`\".

    5.  Set `request`{.variable}'s
        [referrer](#concept-request-referrer){#ref-for-concept-request-referrer①①
        link-type="dfn"} to \"`client`\".

    6.  Set `request`{.variable}'s [referrer
        policy](#concept-request-referrer-policy){#ref-for-concept-request-referrer-policy⑧
        link-type="dfn"} to the empty string.

    7.  Set `request`{.variable}'s
        [URL](#concept-request-url){#ref-for-concept-request-url⑨
        link-type="dfn"} to `request`{.variable}'s [current
        URL](#concept-request-current-url){#ref-for-concept-request-current-url③⑦
        link-type="dfn"}.

    8.  Set `request`{.variable}'s [URL
        list](#concept-request-url-list){#ref-for-concept-request-url-list①①
        link-type="dfn"} to « `request`{.variable}'s
        [URL](#concept-request-url){#ref-for-concept-request-url①⓪
        link-type="dfn"} ».

    This is done to ensure that when a service worker \"redirects\" a
    request, e.g., from an image in a cross-origin style sheet, and
    makes modifications, it no longer appears to come from the original
    source (i.e., the cross-origin style sheet), but instead from the
    service worker that \"redirected\" the request. This is important as
    the original source might not even be able to generate the same kind
    of requests as the service worker. Services that trust the original
    source could therefore be exploited were this not done, although
    that is somewhat farfetched.

14. If
    `init`{.variable}\[\"[`referrer`{.idl}](#dom-requestinit-referrer){#ref-for-dom-requestinit-referrer①
    link-type="idl"}\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists⑤
    link-type="dfn"}, then:

    1.  Let `referrer`{.variable} be
        `init`{.variable}\[\"[`referrer`{.idl}](#dom-requestinit-referrer){#ref-for-dom-requestinit-referrer②
        link-type="idl"}\"\].

    2.  If `referrer`{.variable} is the empty string, then set
        `request`{.variable}'s
        [referrer](#concept-request-referrer){#ref-for-concept-request-referrer①②
        link-type="dfn"} to \"`no-referrer`\".

    3.  Otherwise:

        1.  Let `parsedReferrer`{.variable} be the result of
            [parsing](https://url.spec.whatwg.org/#concept-url-parser){#ref-for-concept-url-parser②
            link-type="dfn"} `referrer`{.variable} with
            `baseURL`{.variable}.

        2.  If `parsedReferrer`{.variable} is failure, then
            [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw①②
            link-type="dfn"} a
            [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror①⑤
            link-type="idl"}.

        3.  If one of the following is true

            - `parsedReferrer`{.variable}'s
              [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#ref-for-concept-url-scheme①⑤
              link-type="dfn"} is \"`about`\" and
              [path](https://url.spec.whatwg.org/#concept-url-path){#ref-for-concept-url-path②
              link-type="dfn"} is the string \"`client`\"

            - `parsedReferrer`{.variable}'s
              [origin](https://url.spec.whatwg.org/#concept-url-origin){#ref-for-concept-url-origin②⓪
              link-type="dfn"} is not [same
              origin](https://html.spec.whatwg.org/multipage/browsers.html#same-origin){#ref-for-same-origin①①
              link-type="dfn"} with `origin`{.variable}

            then set `request`{.variable}'s
            [referrer](#concept-request-referrer){#ref-for-concept-request-referrer①③
            link-type="dfn"} to \"`client`\".

        4.  Otherwise, set `request`{.variable}'s
            [referrer](#concept-request-referrer){#ref-for-concept-request-referrer①④
            link-type="dfn"} to `parsedReferrer`{.variable}.

15. If
    `init`{.variable}\[\"[`referrerPolicy`{.idl}](#dom-requestinit-referrerpolicy){#ref-for-dom-requestinit-referrerpolicy①
    link-type="idl"}\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists⑥
    link-type="dfn"}, then set `request`{.variable}'s [referrer
    policy](#concept-request-referrer-policy){#ref-for-concept-request-referrer-policy⑨
    link-type="dfn"} to it.

16. Let `mode`{.variable} be
    `init`{.variable}\[\"[`mode`{.idl}](#dom-requestinit-mode){#ref-for-dom-requestinit-mode①
    link-type="idl"}\"\] if it
    [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists⑦
    link-type="dfn"}, and `fallbackMode`{.variable} otherwise.

17. If `mode`{.variable} is \"`navigate`\", then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw①③
    link-type="dfn"} a
    [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror①⑥
    link-type="idl"}.

18. If `mode`{.variable} is non-null, set `request`{.variable}'s
    [mode](#concept-request-mode){#ref-for-concept-request-mode②⑧
    link-type="dfn"} to `mode`{.variable}.

19. If
    `init`{.variable}\[\"[`credentials`{.idl}](#dom-requestinit-credentials){#ref-for-dom-requestinit-credentials①
    link-type="idl"}\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists⑧
    link-type="dfn"}, then set `request`{.variable}'s [credentials
    mode](#concept-request-credentials-mode){#ref-for-concept-request-credentials-mode②③
    link-type="dfn"} to it.

20. If
    `init`{.variable}\[\"[`cache`{.idl}](#dom-requestinit-cache){#ref-for-dom-requestinit-cache①
    link-type="idl"}\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists⑨
    link-type="dfn"}, then set `request`{.variable}'s [cache
    mode](#concept-request-cache-mode){#ref-for-concept-request-cache-mode①④
    link-type="dfn"} to it.

21. If `request`{.variable}'s [cache
    mode](#concept-request-cache-mode){#ref-for-concept-request-cache-mode①⑤
    link-type="dfn"} is \"`only-if-cached`\" and `request`{.variable}'s
    [mode](#concept-request-mode){#ref-for-concept-request-mode②⑨
    link-type="dfn"} is *not* \"`same-origin`\", then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw①④
    link-type="dfn"} a
    [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror①⑦
    link-type="idl"}.

22. If
    `init`{.variable}\[\"[`redirect`{.idl}](#dom-requestinit-redirect){#ref-for-dom-requestinit-redirect①
    link-type="idl"}\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists①⓪
    link-type="dfn"}, then set `request`{.variable}'s [redirect
    mode](#concept-request-redirect-mode){#ref-for-concept-request-redirect-mode①①
    link-type="dfn"} to it.

23. If
    `init`{.variable}\[\"[`integrity`{.idl}](#dom-requestinit-integrity){#ref-for-dom-requestinit-integrity①
    link-type="idl"}\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists①①
    link-type="dfn"}, then set `request`{.variable}'s [integrity
    metadata](#concept-request-integrity-metadata){#ref-for-concept-request-integrity-metadata⑤
    link-type="dfn"} to it.

24. If
    `init`{.variable}\[\"[`keepalive`{.idl}](#dom-requestinit-keepalive){#ref-for-dom-requestinit-keepalive①
    link-type="idl"}\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists①②
    link-type="dfn"}, then set `request`{.variable}'s
    [keepalive](#request-keepalive-flag){#ref-for-request-keepalive-flag⑤
    link-type="dfn"} to it.

25. If
    `init`{.variable}\[\"[`method`{.idl}](#dom-requestinit-method){#ref-for-dom-requestinit-method①
    link-type="idl"}\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists①③
    link-type="dfn"}, then:

    1.  Let `method`{.variable} be
        `init`{.variable}\[\"[`method`{.idl}](#dom-requestinit-method){#ref-for-dom-requestinit-method②
        link-type="idl"}\"\].

    2.  If `method`{.variable} is not a
        [method](#concept-method){#ref-for-concept-method①②
        link-type="dfn"} or `method`{.variable} is a [forbidden
        method](#forbidden-method){#ref-for-forbidden-method②
        link-type="dfn"}, then
        [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw①⑤
        link-type="dfn"} a
        [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror①⑧
        link-type="idl"}.

    3.  [Normalize](#concept-method-normalize){#ref-for-concept-method-normalize②
        link-type="dfn"} `method`{.variable}.

    4.  Set `request`{.variable}'s
        [method](#concept-request-method){#ref-for-concept-request-method②②
        link-type="dfn"} to `method`{.variable}.

26. If
    `init`{.variable}\[\"[`signal`{.idl}](#dom-requestinit-signal){#ref-for-dom-requestinit-signal①
    link-type="idl"}\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists①④
    link-type="dfn"}, then set `signal`{.variable} to it.

27. If
    `init`{.variable}\[\"[`priority`{.idl}](#dom-requestinit-priority){#ref-for-dom-requestinit-priority①
    link-type="idl"}\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists①⑤
    link-type="dfn"}, then:

    1.  If `request`{.variable}'s [internal
        priority](#request-internal-priority){#ref-for-request-internal-priority④
        link-type="dfn"} is not null, then update `request`{.variable}'s
        [internal
        priority](#request-internal-priority){#ref-for-request-internal-priority⑤
        link-type="dfn"} in an
        [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#ref-for-implementation-defined①⑨
        link-type="dfn"} manner.

    2.  Otherwise, set `request`{.variable}'s
        [priority](#request-priority){#ref-for-request-priority②
        link-type="dfn"} to
        `init`{.variable}\[\"[`priority`{.idl}](#dom-requestinit-priority){#ref-for-dom-requestinit-priority②
        link-type="idl"}\"\].

28. Set [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③⑥
    link-type="dfn"}'s
    [request](#concept-request-request){#ref-for-concept-request-request④
    link-type="dfn"} to `request`{.variable}.

29. Let `signals`{.variable} be « `signal`{.variable} » if
    `signal`{.variable} is non-null; otherwise « ».

30. Set [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③⑦
    link-type="dfn"}'s
    [signal](#request-signal){#ref-for-request-signal② link-type="dfn"}
    to the result of [creating a dependent abort
    signal](https://dom.spec.whatwg.org/#create-a-dependent-abort-signal){#ref-for-create-a-dependent-abort-signal
    link-type="dfn"} from `signals`{.variable}, using
    [`AbortSignal`{.idl}](https://dom.spec.whatwg.org/#abortsignal){#ref-for-abortsignal⑥
    link-type="idl"} and
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③⑧
    link-type="dfn"}'s [relevant
    realm](https://html.spec.whatwg.org/multipage/webappapis.html#concept-relevant-realm){#ref-for-concept-relevant-realm②
    link-type="dfn"}.

31. Set [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③⑨
    link-type="dfn"}'s
    [headers](#request-headers){#ref-for-request-headers①
    link-type="dfn"} to a
    [new](https://webidl.spec.whatwg.org/#new){#ref-for-new④
    link-type="dfn"} [`Headers`{.idl}](#headers){#ref-for-headers①③
    link-type="idl"} object with
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this④⓪
    link-type="dfn"}'s [relevant
    realm](https://html.spec.whatwg.org/multipage/webappapis.html#concept-relevant-realm){#ref-for-concept-relevant-realm③
    link-type="dfn"}, whose [header
    list](#concept-headers-header-list){#ref-for-concept-headers-header-list①②
    link-type="dfn"} is `request`{.variable}'s [header
    list](#concept-request-header-list){#ref-for-concept-request-header-list④⑨
    link-type="dfn"} and
    [guard](#concept-headers-guard){#ref-for-concept-headers-guard①①
    link-type="dfn"} is \"`request`\".

32. If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this④①
    link-type="dfn"}'s
    [request](#concept-request-request){#ref-for-concept-request-request⑤
    link-type="dfn"}'s
    [mode](#concept-request-mode){#ref-for-concept-request-mode③⓪
    link-type="dfn"} is \"`no-cors`\", then:

    1.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this④②
        link-type="dfn"}'s
        [request](#concept-request-request){#ref-for-concept-request-request⑥
        link-type="dfn"}'s
        [method](#concept-request-method){#ref-for-concept-request-method②③
        link-type="dfn"} is not a [CORS-safelisted
        method](#cors-safelisted-method){#ref-for-cors-safelisted-method④
        link-type="dfn"}, then
        [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw①⑥
        link-type="dfn"} a
        [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror①⑨
        link-type="idl"}.

    2.  Set [this](https://webidl.spec.whatwg.org/#this){#ref-for-this④③
        link-type="dfn"}'s
        [headers](#request-headers){#ref-for-request-headers②
        link-type="dfn"}'s
        [guard](#concept-headers-guard){#ref-for-concept-headers-guard①②
        link-type="dfn"} to \"`request-no-cors`\".

33. If `init`{.variable} [is not
    empty](https://infra.spec.whatwg.org/#map-is-empty){#ref-for-map-is-empty①
    link-type="dfn"}, then:

    The headers are sanitized as they might contain headers that are not
    allowed by this mode. Otherwise, they were previously sanitized or
    are unmodified since they were set by a privileged API.

    1.  Let `headers`{.variable} be a copy of
        [this](https://webidl.spec.whatwg.org/#this){#ref-for-this④④
        link-type="dfn"}'s
        [headers](#request-headers){#ref-for-request-headers③
        link-type="dfn"} and its associated [header
        list](#concept-headers-header-list){#ref-for-concept-headers-header-list①③
        link-type="dfn"}.

    2.  If
        `init`{.variable}\[\"[`headers`{.idl}](#dom-requestinit-headers){#ref-for-dom-requestinit-headers①
        link-type="idl"}\"\]
        [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists①⑥
        link-type="dfn"}, then set `headers`{.variable} to
        `init`{.variable}\[\"[`headers`{.idl}](#dom-requestinit-headers){#ref-for-dom-requestinit-headers②
        link-type="idl"}\"\].

    3.  Empty
        [this](https://webidl.spec.whatwg.org/#this){#ref-for-this④⑤
        link-type="dfn"}'s
        [headers](#request-headers){#ref-for-request-headers④
        link-type="dfn"}'s [header
        list](#concept-headers-header-list){#ref-for-concept-headers-header-list①④
        link-type="dfn"}.

    4.  If `headers`{.variable} is a
        [`Headers`{.idl}](#headers){#ref-for-headers①④ link-type="idl"}
        object, then [for
        each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate①⑥
        link-type="dfn"} `header`{.variable} of its [header
        list](#concept-headers-header-list){#ref-for-concept-headers-header-list①⑤
        link-type="dfn"},
        [append](#concept-headers-append){#ref-for-concept-headers-append③
        link-type="dfn"} `header`{.variable} to
        [this](https://webidl.spec.whatwg.org/#this){#ref-for-this④⑥
        link-type="dfn"}'s
        [headers](#request-headers){#ref-for-request-headers⑤
        link-type="dfn"}.

    5.  Otherwise,
        [fill](#concept-headers-fill){#ref-for-concept-headers-fill①
        link-type="dfn"}
        [this](https://webidl.spec.whatwg.org/#this){#ref-for-this④⑦
        link-type="dfn"}'s
        [headers](#request-headers){#ref-for-request-headers⑥
        link-type="dfn"} with `headers`{.variable}.

34. Let `inputBody`{.variable} be `input`{.variable}'s
    [request](#concept-request-request){#ref-for-concept-request-request⑦
    link-type="dfn"}'s
    [body](#concept-request-body){#ref-for-concept-request-body④⑥
    link-type="dfn"} if `input`{.variable} is a
    [`Request`{.idl}](#request){#ref-for-request①④ link-type="idl"}
    object; otherwise null.

35. If either
    `init`{.variable}\[\"[`body`{.idl}](#dom-requestinit-body){#ref-for-dom-requestinit-body②
    link-type="idl"}\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists①⑦
    link-type="dfn"} and is non-null or `inputBody`{.variable} is
    non-null, and `request`{.variable}'s
    [method](#concept-request-method){#ref-for-concept-request-method②④
    link-type="dfn"} is \``GET`\` or \``HEAD`\`, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw①⑦
    link-type="dfn"} a
    [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror②⓪
    link-type="idl"}.

36. Let `initBody`{.variable} be null.

37. If
    `init`{.variable}\[\"[`body`{.idl}](#dom-requestinit-body){#ref-for-dom-requestinit-body③
    link-type="idl"}\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists①⑧
    link-type="dfn"} and is non-null, then:

    1.  Let `bodyWithType`{.variable} be the result of
        [extracting](#concept-bodyinit-extract){#ref-for-concept-bodyinit-extract②
        link-type="dfn"}
        `init`{.variable}\[\"[`body`{.idl}](#dom-requestinit-body){#ref-for-dom-requestinit-body④
        link-type="idl"}\"\], with
        [*keepalive*](#keepalive){#ref-for-keepalive link-type="dfn"}
        set to `request`{.variable}'s
        [keepalive](#request-keepalive-flag){#ref-for-request-keepalive-flag⑥
        link-type="dfn"}.

    2.  Set `initBody`{.variable} to `bodyWithType`{.variable}'s
        [body](#body-with-type-body){#ref-for-body-with-type-body⑤
        link-type="dfn"}.

    3.  Let `type`{.variable} be `bodyWithType`{.variable}'s
        [type](#body-with-type-type){#ref-for-body-with-type-type
        link-type="dfn"}.

    4.  If `type`{.variable} is non-null and
        [this](https://webidl.spec.whatwg.org/#this){#ref-for-this④⑧
        link-type="dfn"}'s
        [headers](#request-headers){#ref-for-request-headers⑦
        link-type="dfn"}'s [header
        list](#concept-headers-header-list){#ref-for-concept-headers-header-list①⑥
        link-type="dfn"} [does not
        contain](#header-list-contains){#ref-for-header-list-contains②⑥
        link-type="dfn"} \``Content-Type`\`, then
        [append](#concept-headers-append){#ref-for-concept-headers-append④
        link-type="dfn"} (\``Content-Type`\`, `type`{.variable}) to
        [this](https://webidl.spec.whatwg.org/#this){#ref-for-this④⑨
        link-type="dfn"}'s
        [headers](#request-headers){#ref-for-request-headers⑧
        link-type="dfn"}.

38. Let `inputOrInitBody`{.variable} be `initBody`{.variable} if it is
    non-null; otherwise `inputBody`{.variable}.

39. If `inputOrInitBody`{.variable} is non-null and
    `inputOrInitBody`{.variable}'s
    [source](#concept-body-source){#ref-for-concept-body-source①④
    link-type="dfn"} is null, then:

    1.  If `initBody`{.variable} is non-null and
        `init`{.variable}\[\"[`duplex`{.idl}](#dom-requestinit-duplex){#ref-for-dom-requestinit-duplex①
        link-type="idl"}\"\] does not
        [exist](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists①⑨
        link-type="dfn"}, then throw a
        [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror②①
        link-type="idl"}.

    2.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑤⓪
        link-type="dfn"}'s
        [request](#concept-request-request){#ref-for-concept-request-request⑧
        link-type="dfn"}'s
        [mode](#concept-request-mode){#ref-for-concept-request-mode③①
        link-type="dfn"} is neither \"`same-origin`\" nor \"`cors`\",
        then throw a
        [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror②②
        link-type="idl"}.

    3.  Set [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑤①
        link-type="dfn"}'s
        [request](#concept-request-request){#ref-for-concept-request-request⑨
        link-type="dfn"}'s [use-CORS-preflight
        flag](#use-cors-preflight-flag){#ref-for-use-cors-preflight-flag⑥
        link-type="dfn"}.

40. Let `finalBody`{.variable} be `inputOrInitBody`{.variable}.

41. If `initBody`{.variable} is null and `inputBody`{.variable} is
    non-null, then:

    1.  If `input`{.variable} is
        [unusable](#body-unusable){#ref-for-body-unusable①
        link-type="dfn"}, then
        [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw①⑧
        link-type="dfn"} a
        [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror②③
        link-type="idl"}.

    2.  Set `finalBody`{.variable} to the result of [creating a
        proxy](https://streams.spec.whatwg.org/#readablestream-create-a-proxy){#ref-for-readablestream-create-a-proxy
        link-type="dfn"} for `inputBody`{.variable}.

42. Set [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑤②
    link-type="dfn"}'s
    [request](#concept-request-request){#ref-for-concept-request-request①⓪
    link-type="dfn"}'s
    [body](#concept-request-body){#ref-for-concept-request-body④⑦
    link-type="dfn"} to `finalBody`{.variable}.
:::

The [`method`]{#dom-request-method .dfn .dfn-paneled .idl-code
dfn-for="Request" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑤③
link-type="dfn"}'s
[request](#concept-request-request){#ref-for-concept-request-request①①
link-type="dfn"}'s
[method](#concept-request-method){#ref-for-concept-request-method②⑤
link-type="dfn"}.

The [`url`]{#dom-request-url .dfn .dfn-paneled .idl-code
dfn-for="Request" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑤④
link-type="dfn"}'s
[request](#concept-request-request){#ref-for-concept-request-request①②
link-type="dfn"}'s
[URL](#concept-request-url){#ref-for-concept-request-url①①
link-type="dfn"},
[serialized](https://url.spec.whatwg.org/#concept-url-serializer){#ref-for-concept-url-serializer②
link-type="dfn"}.

The [`headers`]{#dom-request-headers .dfn .dfn-paneled .idl-code
dfn-for="Request" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑤⑤
link-type="dfn"}'s [headers](#request-headers){#ref-for-request-headers⑨
link-type="dfn"}.

The [`destination`]{#dom-request-destination .dfn .dfn-paneled .idl-code
dfn-for="Request" dfn-type="attribute" export=""} getter are to return
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑤⑥
link-type="dfn"}'s
[request](#concept-request-request){#ref-for-concept-request-request①③
link-type="dfn"}'s
[destination](#concept-request-destination){#ref-for-concept-request-destination②①
link-type="dfn"}.

::: {.algorithm algorithm="referrer" algorithm-for="Request"}
The [`referrer`]{#dom-request-referrer .dfn .dfn-paneled .idl-code
dfn-for="Request" dfn-type="attribute" export=""} getter steps are:

1.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑤⑦
    link-type="dfn"}'s
    [request](#concept-request-request){#ref-for-concept-request-request①④
    link-type="dfn"}'s
    [referrer](#concept-request-referrer){#ref-for-concept-request-referrer①⑤
    link-type="dfn"} is \"`no-referrer`\", then return the empty string.

2.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑤⑧
    link-type="dfn"}'s
    [request](#concept-request-request){#ref-for-concept-request-request①⑤
    link-type="dfn"}'s
    [referrer](#concept-request-referrer){#ref-for-concept-request-referrer①⑥
    link-type="dfn"} is \"`client`\", then return \"`about:client`\".

3.  Return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑤⑨
    link-type="dfn"}'s
    [request](#concept-request-request){#ref-for-concept-request-request①⑥
    link-type="dfn"}'s
    [referrer](#concept-request-referrer){#ref-for-concept-request-referrer①⑦
    link-type="dfn"},
    [serialized](https://url.spec.whatwg.org/#concept-url-serializer){#ref-for-concept-url-serializer③
    link-type="dfn"}.
:::

The [`referrerPolicy`]{#dom-request-referrerpolicy .dfn .dfn-paneled
.idl-code dfn-for="Request" dfn-type="attribute" export=""} getter steps
are to return
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑥⓪
link-type="dfn"}'s
[request](#concept-request-request){#ref-for-concept-request-request①⑦
link-type="dfn"}'s [referrer
policy](#concept-request-referrer-policy){#ref-for-concept-request-referrer-policy①⓪
link-type="dfn"}.

The [`mode`]{#dom-request-mode .dfn .dfn-paneled .idl-code
dfn-for="Request" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑥①
link-type="dfn"}'s
[request](#concept-request-request){#ref-for-concept-request-request①⑧
link-type="dfn"}'s
[mode](#concept-request-mode){#ref-for-concept-request-mode③②
link-type="dfn"}.

The [`credentials`]{#dom-request-credentials .dfn .dfn-paneled .idl-code
dfn-for="Request" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑥②
link-type="dfn"}'s
[request](#concept-request-request){#ref-for-concept-request-request①⑨
link-type="dfn"}'s [credentials
mode](#concept-request-credentials-mode){#ref-for-concept-request-credentials-mode②④
link-type="dfn"}.

The [`cache`]{#dom-request-cache .dfn .dfn-paneled .idl-code
dfn-for="Request" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑥③
link-type="dfn"}'s
[request](#concept-request-request){#ref-for-concept-request-request②⓪
link-type="dfn"}'s [cache
mode](#concept-request-cache-mode){#ref-for-concept-request-cache-mode①⑥
link-type="dfn"}.

The [`redirect`]{#dom-request-redirect .dfn .dfn-paneled .idl-code
dfn-for="Request" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑥④
link-type="dfn"}'s
[request](#concept-request-request){#ref-for-concept-request-request②①
link-type="dfn"}'s [redirect
mode](#concept-request-redirect-mode){#ref-for-concept-request-redirect-mode①②
link-type="dfn"}.

The [`integrity`]{#dom-request-integrity .dfn .dfn-paneled .idl-code
dfn-for="Request" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑥⑤
link-type="dfn"}'s
[request](#concept-request-request){#ref-for-concept-request-request②②
link-type="dfn"}'s [integrity
metadata](#concept-request-integrity-metadata){#ref-for-concept-request-integrity-metadata⑥
link-type="dfn"}.

The [`keepalive`]{#dom-request-keepalive .dfn .dfn-paneled .idl-code
dfn-for="Request" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑥⑥
link-type="dfn"}'s
[request](#concept-request-request){#ref-for-concept-request-request②③
link-type="dfn"}'s
[keepalive](#request-keepalive-flag){#ref-for-request-keepalive-flag⑦
link-type="dfn"}.

The [`isReloadNavigation`]{#dom-request-isreloadnavigation .dfn
.dfn-paneled .idl-code dfn-for="Request" dfn-type="attribute" export=""}
getter steps are to return true if
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑥⑦
link-type="dfn"}'s
[request](#concept-request-request){#ref-for-concept-request-request②④
link-type="dfn"}'s [reload-navigation
flag](#concept-request-reload-navigation-flag){#ref-for-concept-request-reload-navigation-flag③
link-type="dfn"} is set; otherwise false.

The [`isHistoryNavigation`]{#dom-request-ishistorynavigation .dfn
.dfn-paneled .idl-code dfn-for="Request" dfn-type="attribute" export=""}
getter steps are to return true if
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑥⑧
link-type="dfn"}'s
[request](#concept-request-request){#ref-for-concept-request-request②⑤
link-type="dfn"}'s [history-navigation
flag](#concept-request-history-navigation-flag){#ref-for-concept-request-history-navigation-flag③
link-type="dfn"} is set; otherwise false.

The [`signal`]{#dom-request-signal .dfn .dfn-paneled .idl-code
dfn-for="Request" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑥⑨
link-type="dfn"}'s [signal](#request-signal){#ref-for-request-signal③
link-type="dfn"}.

The [`duplex`]{#dom-request-duplex .dfn .dfn-paneled .idl-code
dfn-for="Request" dfn-type="attribute" export=""} getter steps are to
return \"`half`\".

------------------------------------------------------------------------

::: {.algorithm algorithm="clone()" algorithm-for="Request"}
The [`clone()`]{#dom-request-clone .dfn .dfn-paneled .idl-code
dfn-for="Request" dfn-type="method" export=""} method steps are:

1.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑦⓪
    link-type="dfn"} is
    [unusable](#body-unusable){#ref-for-body-unusable② link-type="dfn"},
    then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw①⑨
    link-type="dfn"} a
    [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror②④
    link-type="idl"}.

2.  Let `clonedRequest`{.variable} be the result of
    [cloning](#concept-request-clone){#ref-for-concept-request-clone③
    link-type="dfn"}
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑦①
    link-type="dfn"}'s
    [request](#concept-request-request){#ref-for-concept-request-request②⑥
    link-type="dfn"}.

3.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert②⑥
    link-type="dfn"}:
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑦②
    link-type="dfn"}'s
    [signal](#request-signal){#ref-for-request-signal④ link-type="dfn"}
    is non-null.

4.  Let `clonedSignal`{.variable} be the result of [creating a dependent
    abort
    signal](https://dom.spec.whatwg.org/#create-a-dependent-abort-signal){#ref-for-create-a-dependent-abort-signal①
    link-type="dfn"} from «
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑦③
    link-type="dfn"}'s
    [signal](#request-signal){#ref-for-request-signal⑤ link-type="dfn"}
    », using
    [`AbortSignal`{.idl}](https://dom.spec.whatwg.org/#abortsignal){#ref-for-abortsignal⑦
    link-type="idl"} and
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑦④
    link-type="dfn"}'s [relevant
    realm](https://html.spec.whatwg.org/multipage/webappapis.html#concept-relevant-realm){#ref-for-concept-relevant-realm④
    link-type="dfn"}.

5.  Let `clonedRequestObject`{.variable} be the result of
    [creating](#request-create){#ref-for-request-create link-type="dfn"}
    a [`Request`{.idl}](#request){#ref-for-request①⑤ link-type="idl"}
    object, given `clonedRequest`{.variable},
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑦⑤
    link-type="dfn"}'s
    [headers](#request-headers){#ref-for-request-headers①⓪
    link-type="dfn"}'s
    [guard](#concept-headers-guard){#ref-for-concept-headers-guard①③
    link-type="dfn"}, `clonedSignal`{.variable} and
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑦⑥
    link-type="dfn"}'s [relevant
    realm](https://html.spec.whatwg.org/multipage/webappapis.html#concept-relevant-realm){#ref-for-concept-relevant-realm⑤
    link-type="dfn"}.

6.  Return `clonedRequestObject`{.variable}.
:::

### [5.5. ]{.secno}[Response class]{.content}[](#response-class){.self-link} {#response-class .heading .settled level="5.5"}

``` {.idl .highlight .def}
[Exposed=(Window,Worker)]interface Response {
  constructor(optional BodyInit? body = null, optional ResponseInit init = {});

  [NewObject] static Response error();
  [NewObject] static Response redirect(USVString url, optional unsigned short status = 302);
  [NewObject] static Response json(any data, optional ResponseInit init = {});

  readonly attribute ResponseType type;

  readonly attribute USVString url;
  readonly attribute boolean redirected;
  readonly attribute unsigned short status;
  readonly attribute boolean ok;
  readonly attribute ByteString statusText;
  [SameObject] readonly attribute Headers headers;

  [NewObject] Response clone();
};
Response includes Body;

dictionary ResponseInit {
  unsigned short status = 200;
  ByteString statusText = "";
  HeadersInit headers;
};

enum ResponseType { "basic", "cors", "default", "error", "opaque", "opaqueredirect" };
```

A [`Response`{.idl}](#response){#ref-for-response⑦ link-type="idl"}
object has an associated [response]{#concept-response-response .dfn
.dfn-paneled dfn-for="Response" dfn-type="dfn" export=""} (a
[response](#concept-response){#ref-for-concept-response⑥①
link-type="dfn"}).

A [`Response`{.idl}](#response){#ref-for-response⑧ link-type="idl"}
object also has an associated [headers]{#response-headers .dfn
.dfn-paneled dfn-for="Response" dfn-type="dfn" export=""} (null or a
[`Headers`{.idl}](#headers){#ref-for-headers①⑥ link-type="idl"} object),
initially null.

A [`Response`{.idl}](#response){#ref-for-response⑨ link-type="idl"}
object's [body](#concept-body-body){#ref-for-concept-body-body⑨
link-type="dfn"} is its
[response](#concept-response-response){#ref-for-concept-response-response①
link-type="dfn"}'s
[body](#concept-response-body){#ref-for-concept-response-body②⑦
link-type="dfn"}.

------------------------------------------------------------------------

`response`{.variable}` = new `[`Response`](#dom-response){#ref-for-dom-response① .idl-code link-type="constructor"}`(``body`{.variable}` = null [, ``init`{.variable}`])`

:   Creates a [`Response`{.idl}](#response){#ref-for-response①⓪
    link-type="idl"} whose body is `body`{.variable}, and status, status
    message, and headers are provided by `init`{.variable}.

`response`{.variable}` = `[`Response`](#response){#ref-for-response①① link-type="idl"}` . `[`error`](#dom-response-error){#ref-for-dom-response-error① .idl-code link-type="method"}`()`

:   Creates network error
    [`Response`{.idl}](#response){#ref-for-response①② link-type="idl"}.

`response`{.variable}` = `[`Response`](#response){#ref-for-response①③ link-type="idl"}` . `[`redirect`](#dom-response-redirect){#ref-for-dom-response-redirect① .idl-code link-type="method"}`(``url`{.variable}`, ``status`{.variable}` = 302)`

:   Creates a redirect [`Response`{.idl}](#response){#ref-for-response①④
    link-type="idl"} that redirects to `url`{.variable} with status
    `status`{.variable}.

`response`{.variable}` = `[`Response`](#response){#ref-for-response①⑤ link-type="idl"}` . `[`json`](#dom-response-json){#ref-for-dom-response-json① .idl-code link-type="method"}`(``data`{.variable}` [, ``init`{.variable}`])`

:   Creates a [`Response`{.idl}](#response){#ref-for-response①⑥
    link-type="idl"} whose body is the JSON-encoded `data`{.variable},
    and status, status message, and headers are provided by
    `init`{.variable}.

`response`{.variable}` . `[`type`](#dom-response-type){#ref-for-dom-response-type② .idl-code link-type="attribute"}

:   Returns `response`{.variable}'s type, e.g., \"`cors`\".

`response`{.variable}` . `[`url`](#dom-response-url){#ref-for-dom-response-url① .idl-code link-type="attribute"}

:   Returns `response`{.variable}'s URL, if it has one; otherwise the
    empty string.

`response`{.variable}` . `[`redirected`](#dom-response-redirected){#ref-for-dom-response-redirected① .idl-code link-type="attribute"}

:   Returns whether `response`{.variable} was obtained through a
    redirect.

`response`{.variable}` . `[`status`](#dom-response-status){#ref-for-dom-response-status① .idl-code link-type="attribute"}

:   Returns `response`{.variable}'s status.

`response`{.variable}` . `[`ok`](#dom-response-ok){#ref-for-dom-response-ok② .idl-code link-type="attribute"}

:   Returns whether `response`{.variable}'s status is an [ok
    status](#ok-status){#ref-for-ok-status② link-type="dfn"}.

`response`{.variable}` . `[`statusText`](#dom-response-statustext){#ref-for-dom-response-statustext① .idl-code link-type="attribute"}

:   Returns `response`{.variable}'s status message.

`response`{.variable}` . `[`headers`](#dom-response-headers){#ref-for-dom-response-headers① .idl-code link-type="attribute"}

:   Returns `response`{.variable}'s headers as
    [`Headers`{.idl}](#headers){#ref-for-headers①⑦ link-type="idl"}.

`response`{.variable}` . `[`clone`](#dom-response-clone){#ref-for-dom-response-clone① .idl-code link-type="method"}`()`

:   Returns a clone of `response`{.variable}.

------------------------------------------------------------------------

::: {.algorithm algorithm="create" algorithm-for="Response"}
To [create]{#response-create .dfn .dfn-paneled dfn-for="Response"
dfn-type="dfn" export="" lt="create|creating"} a
[`Response`{.idl}](#response){#ref-for-response①⑦ link-type="idl"}
object, given a
[response](#concept-response){#ref-for-concept-response⑥②
link-type="dfn"} `response`{.variable}, [headers
guard](#headers-guard){#ref-for-headers-guard② link-type="dfn"}
`guard`{.variable}, and
[realm](https://tc39.es/ecma262/#realm){#ref-for-realm③ link-type="dfn"}
`realm`{.variable}, run these steps:

1.  Let `responseObject`{.variable} be a
    [new](https://webidl.spec.whatwg.org/#new){#ref-for-new⑤
    link-type="dfn"} [`Response`{.idl}](#response){#ref-for-response①⑧
    link-type="idl"} object with `realm`{.variable}.

2.  Set `responseObject`{.variable}'s
    [response](#concept-response-response){#ref-for-concept-response-response②
    link-type="dfn"} to `response`{.variable}.

3.  Set `responseObject`{.variable}'s
    [headers](#response-headers){#ref-for-response-headers
    link-type="dfn"} to a
    [new](https://webidl.spec.whatwg.org/#new){#ref-for-new⑥
    link-type="dfn"} [`Headers`{.idl}](#headers){#ref-for-headers①⑧
    link-type="idl"} object with `realm`{.variable}, whose [headers
    list](#concept-headers-header-list){#ref-for-concept-headers-header-list①⑦
    link-type="dfn"} is `response`{.variable}'s [headers
    list](#concept-response-header-list){#ref-for-concept-response-header-list③③
    link-type="dfn"} and
    [guard](#concept-headers-guard){#ref-for-concept-headers-guard①④
    link-type="dfn"} is `guard`{.variable}.

4.  Return `responseObject`{.variable}.
:::

::: {.algorithm algorithm="initialize a response"}
To [initialize a response]{#initialize-a-response .dfn .dfn-paneled
dfn-type="dfn" noexport=""}, given a
[`Response`{.idl}](#response){#ref-for-response①⑨ link-type="idl"}
object `response`{.variable},
[`ResponseInit`{.idl}](#responseinit){#ref-for-responseinit②
link-type="idl"} `init`{.variable}, and null or a [body with
type](#body-with-type){#ref-for-body-with-type② link-type="dfn"}
`body`{.variable}:

1.  If
    `init`{.variable}\[\"[`status`{.idl}](#dom-responseinit-status){#ref-for-dom-responseinit-status
    link-type="idl"}\"\] is not in the range 200 to 599, inclusive, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw②⓪
    link-type="dfn"} a
    [`RangeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-rangeerror){#ref-for-exceptiondef-rangeerror②
    link-type="idl"}.

2.  If
    `init`{.variable}\[\"[`statusText`{.idl}](#dom-responseinit-statustext){#ref-for-dom-responseinit-statustext
    link-type="idl"}\"\] is not the empty string and does not match the
    [reason-phrase](https://httpwg.org/specs/rfc9112.html#status.line){#ref-for-status.line
    link-type="dfn"} token production, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw②①
    link-type="dfn"} a
    [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror②⑤
    link-type="idl"}.

3.  Set `response`{.variable}'s
    [response](#concept-response-response){#ref-for-concept-response-response③
    link-type="dfn"}'s
    [status](#concept-response-status){#ref-for-concept-response-status②①
    link-type="dfn"} to
    `init`{.variable}\[\"[`status`{.idl}](#dom-responseinit-status){#ref-for-dom-responseinit-status①
    link-type="idl"}\"\].

4.  Set `response`{.variable}'s
    [response](#concept-response-response){#ref-for-concept-response-response④
    link-type="dfn"}'s [status
    message](#concept-response-status-message){#ref-for-concept-response-status-message⑦
    link-type="dfn"} to
    `init`{.variable}\[\"[`statusText`{.idl}](#dom-responseinit-statustext){#ref-for-dom-responseinit-statustext①
    link-type="idl"}\"\].

5.  If
    `init`{.variable}\[\"[`headers`{.idl}](#dom-responseinit-headers){#ref-for-dom-responseinit-headers
    link-type="idl"}\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists②⓪
    link-type="dfn"}, then
    [fill](#concept-headers-fill){#ref-for-concept-headers-fill②
    link-type="dfn"} `response`{.variable}'s
    [headers](#response-headers){#ref-for-response-headers①
    link-type="dfn"} with
    `init`{.variable}\[\"[`headers`{.idl}](#dom-responseinit-headers){#ref-for-dom-responseinit-headers①
    link-type="idl"}\"\].

6.  If `body`{.variable} is non-null, then:

    1.  If `response`{.variable}'s
        [status](#concept-response-status){#ref-for-concept-response-status②②
        link-type="dfn"} is a [null body
        status](#null-body-status){#ref-for-null-body-status①
        link-type="dfn"}, then
        [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw②②
        link-type="dfn"} a
        [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror②⑥
        link-type="idl"}.

        101 and 103 are included in [null body
        status](#null-body-status){#ref-for-null-body-status②
        link-type="dfn"} due to their use elsewhere. They do not affect
        this step.

    2.  Set `response`{.variable}'s
        [body](#concept-response-body){#ref-for-concept-response-body②⑧
        link-type="dfn"} to `body`{.variable}'s
        [body](#body-with-type-body){#ref-for-body-with-type-body⑥
        link-type="dfn"}.

    3.  If `body`{.variable}'s
        [type](#body-with-type-type){#ref-for-body-with-type-type①
        link-type="dfn"} is non-null and `response`{.variable}'s [header
        list](#concept-response-header-list){#ref-for-concept-response-header-list③④
        link-type="dfn"} [does not
        contain](#header-list-contains){#ref-for-header-list-contains②⑦
        link-type="dfn"} \``Content-Type`\`, then
        [append](#concept-header-list-append){#ref-for-concept-header-list-append②②
        link-type="dfn"} (\``Content-Type`\`, `body`{.variable}'s
        [type](#body-with-type-type){#ref-for-body-with-type-type②
        link-type="dfn"}) to `response`{.variable}'s [header
        list](#concept-response-header-list){#ref-for-concept-response-header-list③⑤
        link-type="dfn"}.
:::

------------------------------------------------------------------------

::: {.algorithm algorithm="Response(body, init)" algorithm-for="Response"}
The
[`new Response(``body`{.variable}`, ``init`{.variable}`)`]{#dom-response
.dfn .dfn-paneled .idl-code dfn-for="Response" dfn-type="constructor"
export=""
lt="Response(body, init)|constructor(body, init)|Response(body)|constructor(body)|Response()|constructor()"}
constructor steps are:

1.  Set [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑦⑦
    link-type="dfn"}'s
    [response](#concept-response-response){#ref-for-concept-response-response⑤
    link-type="dfn"} to a new
    [response](#concept-response){#ref-for-concept-response⑥③
    link-type="dfn"}.

2.  Set [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑦⑧
    link-type="dfn"}'s
    [headers](#response-headers){#ref-for-response-headers②
    link-type="dfn"} to a
    [new](https://webidl.spec.whatwg.org/#new){#ref-for-new⑦
    link-type="dfn"} [`Headers`{.idl}](#headers){#ref-for-headers①⑨
    link-type="idl"} object with
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑦⑨
    link-type="dfn"}'s [relevant
    realm](https://html.spec.whatwg.org/multipage/webappapis.html#concept-relevant-realm){#ref-for-concept-relevant-realm⑥
    link-type="dfn"}, whose [header
    list](#concept-headers-header-list){#ref-for-concept-headers-header-list①⑧
    link-type="dfn"} is
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑧⓪
    link-type="dfn"}'s
    [response](#concept-response-response){#ref-for-concept-response-response⑥
    link-type="dfn"}'s [header
    list](#concept-response-header-list){#ref-for-concept-response-header-list③⑥
    link-type="dfn"} and
    [guard](#concept-headers-guard){#ref-for-concept-headers-guard①⑤
    link-type="dfn"} is \"`response`\".

3.  Let `bodyWithType`{.variable} be null.

4.  If `body`{.variable} is non-null, then set `bodyWithType`{.variable}
    to the result of
    [extracting](#concept-bodyinit-extract){#ref-for-concept-bodyinit-extract③
    link-type="dfn"} `body`{.variable}.

5.  Perform [initialize a
    response](#initialize-a-response){#ref-for-initialize-a-response
    link-type="dfn"} given
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑧①
    link-type="dfn"}, `init`{.variable}, and `bodyWithType`{.variable}.
:::

The static [`error()`]{#dom-response-error .dfn .dfn-paneled .idl-code
dfn-for="Response" dfn-type="method" export=""} method steps are to
return the result of
[creating](#response-create){#ref-for-response-create link-type="dfn"} a
[`Response`{.idl}](#response){#ref-for-response②⓪ link-type="idl"}
object, given a new [network
error](#concept-network-error){#ref-for-concept-network-error⑤⑨
link-type="dfn"}, \"`immutable`\", and the [current
realm](https://tc39.es/ecma262/#current-realm){#ref-for-current-realm
link-type="dfn"}.

::: {.algorithm algorithm="redirect(url, status)" algorithm-for="Response"}
The static
[`redirect(``url`{.variable}`, ``status`{.variable}`)`]{#dom-response-redirect
.dfn .dfn-paneled .idl-code dfn-for="Response" dfn-type="method"
export="" lt="redirect(url, status)|redirect(url)"} method steps are:

1.  Let `parsedURL`{.variable} be the result of
    [parsing](https://url.spec.whatwg.org/#concept-url-parser){#ref-for-concept-url-parser③
    link-type="dfn"} `url`{.variable} with [current settings
    object](https://html.spec.whatwg.org/multipage/webappapis.html#current-settings-object){#ref-for-current-settings-object
    link-type="dfn"}'s [API base
    URL](https://html.spec.whatwg.org/multipage/webappapis.html#api-base-url){#ref-for-api-base-url①
    link-type="dfn"}.

2.  If `parsedURL`{.variable} is failure, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw②③
    link-type="dfn"} a
    [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror②⑦
    link-type="idl"}.

3.  If `status`{.variable} is not a [redirect
    status](#redirect-status){#ref-for-redirect-status③
    link-type="dfn"}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw②④
    link-type="dfn"} a
    [`RangeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-rangeerror){#ref-for-exceptiondef-rangeerror③
    link-type="idl"}.

4.  Let `responseObject`{.variable} be the result of
    [creating](#response-create){#ref-for-response-create①
    link-type="dfn"} a [`Response`{.idl}](#response){#ref-for-response②①
    link-type="idl"} object, given a new
    [response](#concept-response){#ref-for-concept-response⑥④
    link-type="dfn"}, \"`immutable`\", and the [current
    realm](https://tc39.es/ecma262/#current-realm){#ref-for-current-realm①
    link-type="dfn"}.

5.  Set `responseObject`{.variable}'s
    [response](#concept-response-response){#ref-for-concept-response-response⑦
    link-type="dfn"}'s
    [status](#concept-response-status){#ref-for-concept-response-status②③
    link-type="dfn"} to `status`{.variable}.

6.  Let `value`{.variable} be `parsedURL`{.variable},
    [serialized](https://url.spec.whatwg.org/#concept-url-serializer){#ref-for-concept-url-serializer④
    link-type="dfn"} and [isomorphic
    encoded](https://infra.spec.whatwg.org/#isomorphic-encode){#ref-for-isomorphic-encode①①
    link-type="dfn"}.

7.  [Append](#concept-header-list-append){#ref-for-concept-header-list-append②③
    link-type="dfn"} (\``Location`\`, `value`{.variable}) to
    `responseObject`{.variable}'s
    [response](#concept-response-response){#ref-for-concept-response-response⑧
    link-type="dfn"}'s [header
    list](#concept-response-header-list){#ref-for-concept-response-header-list③⑦
    link-type="dfn"}.

8.  Return `responseObject`{.variable}.
:::

::: {.algorithm algorithm="json(data, init)" algorithm-for="Response"}
The static
[`json(``data`{.variable}`, ``init`{.variable}`)`]{#dom-response-json
.dfn .dfn-paneled .idl-code dfn-for="Response" dfn-type="method"
export="" lt="json(data, init)|json(data)"} method steps are:

1.  Let `bytes`{.variable} the result of running [serialize a JavaScript
    value to JSON
    bytes](https://infra.spec.whatwg.org/#serialize-a-javascript-value-to-json-bytes){#ref-for-serialize-a-javascript-value-to-json-bytes
    link-type="dfn"} on `data`{.variable}.

2.  Let `body`{.variable} be the result of
    [extracting](#concept-bodyinit-extract){#ref-for-concept-bodyinit-extract④
    link-type="dfn"} `bytes`{.variable}.

3.  Let `responseObject`{.variable} be the result of
    [creating](#response-create){#ref-for-response-create②
    link-type="dfn"} a [`Response`{.idl}](#response){#ref-for-response②②
    link-type="idl"} object, given a new
    [response](#concept-response){#ref-for-concept-response⑥⑤
    link-type="dfn"}, \"`response`\", and the [current
    realm](https://tc39.es/ecma262/#current-realm){#ref-for-current-realm②
    link-type="dfn"}.

4.  Perform [initialize a
    response](#initialize-a-response){#ref-for-initialize-a-response①
    link-type="dfn"} given `responseObject`{.variable},
    `init`{.variable}, and (`body`{.variable}, \"`application/json`\").

5.  Return `responseObject`{.variable}.
:::

The [`type`]{#dom-response-type .dfn .dfn-paneled .idl-code
dfn-for="Response" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑧②
link-type="dfn"}'s
[response](#concept-response-response){#ref-for-concept-response-response⑨
link-type="dfn"}'s
[type](#concept-response-type){#ref-for-concept-response-type①②
link-type="dfn"}.

The [`url`]{#dom-response-url .dfn .dfn-paneled .idl-code
dfn-for="Response" dfn-type="attribute" export=""} getter steps are to
return the empty string if
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑧③
link-type="dfn"}'s
[response](#concept-response-response){#ref-for-concept-response-response①⓪
link-type="dfn"}'s
[URL](#concept-response-url){#ref-for-concept-response-url⑧
link-type="dfn"} is null; otherwise
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑧④
link-type="dfn"}'s
[response](#concept-response-response){#ref-for-concept-response-response①①
link-type="dfn"}'s
[URL](#concept-response-url){#ref-for-concept-response-url⑨
link-type="dfn"},
[serialized](https://url.spec.whatwg.org/#concept-url-serializer){#ref-for-concept-url-serializer⑤
link-type="dfn"} with [*exclude
fragment*](https://url.spec.whatwg.org/#url-serializer-exclude-fragment){#ref-for-url-serializer-exclude-fragment①
link-type="dfn"} set to true.

The [`redirected`]{#dom-response-redirected .dfn .dfn-paneled .idl-code
dfn-for="Response" dfn-type="attribute" export=""} getter steps are to
return true if
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑧⑤
link-type="dfn"}'s
[response](#concept-response-response){#ref-for-concept-response-response①②
link-type="dfn"}'s [URL
list](#concept-response-url-list){#ref-for-concept-response-url-list①①
link-type="dfn"}'s
[size](https://infra.spec.whatwg.org/#list-size){#ref-for-list-size②
link-type="dfn"} is greater than 1; otherwise false.

To filter out [responses](#concept-response){#ref-for-concept-response⑥⑥
link-type="dfn"} that are the result of a redirect, do this directly
through the API, e.g., `fetch(url, { redirect:"error" })`. This way a
potentially unsafe
[response](#concept-response){#ref-for-concept-response⑥⑦
link-type="dfn"} cannot accidentally leak.

The [`status`]{#dom-response-status .dfn .dfn-paneled .idl-code
dfn-for="Response" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑧⑥
link-type="dfn"}'s
[response](#concept-response-response){#ref-for-concept-response-response①③
link-type="dfn"}'s
[status](#concept-response-status){#ref-for-concept-response-status②④
link-type="dfn"}.

The [`ok`]{#dom-response-ok .dfn .dfn-paneled .idl-code
dfn-for="Response" dfn-type="attribute" export=""} getter steps are to
return true if
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑧⑦
link-type="dfn"}'s
[response](#concept-response-response){#ref-for-concept-response-response①④
link-type="dfn"}'s
[status](#concept-response-status){#ref-for-concept-response-status②⑤
link-type="dfn"} is an [ok status](#ok-status){#ref-for-ok-status③
link-type="dfn"}; otherwise false.

The [`statusText`]{#dom-response-statustext .dfn .dfn-paneled .idl-code
dfn-for="Response" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑧⑧
link-type="dfn"}'s
[response](#concept-response-response){#ref-for-concept-response-response①⑤
link-type="dfn"}'s [status
message](#concept-response-status-message){#ref-for-concept-response-status-message⑧
link-type="dfn"}.

The [`headers`]{#dom-response-headers .dfn .dfn-paneled .idl-code
dfn-for="Response" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑧⑨
link-type="dfn"}'s
[headers](#response-headers){#ref-for-response-headers③
link-type="dfn"}.

------------------------------------------------------------------------

::: {.algorithm algorithm="clone()" algorithm-for="Response"}
The [`clone()`]{#dom-response-clone .dfn .dfn-paneled .idl-code
dfn-for="Response" dfn-type="method" export=""} method steps are:

1.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑨⓪
    link-type="dfn"} is
    [unusable](#body-unusable){#ref-for-body-unusable③ link-type="dfn"},
    then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw②⑤
    link-type="dfn"} a
    [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror②⑧
    link-type="idl"}.

2.  Let `clonedResponse`{.variable} be the result of
    [cloning](#concept-response-clone){#ref-for-concept-response-clone①
    link-type="dfn"}
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑨①
    link-type="dfn"}'s
    [response](#concept-response-response){#ref-for-concept-response-response①⑥
    link-type="dfn"}.

3.  Return the result of
    [creating](#response-create){#ref-for-response-create③
    link-type="dfn"} a [`Response`{.idl}](#response){#ref-for-response②③
    link-type="idl"} object, given `clonedResponse`{.variable},
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑨②
    link-type="dfn"}'s
    [headers](#response-headers){#ref-for-response-headers④
    link-type="dfn"}'s
    [guard](#concept-headers-guard){#ref-for-concept-headers-guard①⑥
    link-type="dfn"}, and
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑨③
    link-type="dfn"}'s [relevant
    realm](https://html.spec.whatwg.org/multipage/webappapis.html#concept-relevant-realm){#ref-for-concept-relevant-realm⑦
    link-type="dfn"}.
:::

### [5.6. ]{.secno}[Fetch method]{.content}[](#fetch-method){.self-link} {#fetch-method .heading .settled level="5.6"}

``` {.idl .highlight .def}
partial interface mixin WindowOrWorkerGlobalScope {
  [NewObject] Promise<Response> fetch(RequestInfo input, optional RequestInit init = {});
};
```

::: {.algorithm algorithm="fetch(input, init)" algorithm-for="WindowOrWorkerGlobalScope"}
The
[`fetch(``input`{.variable}`, ``init`{.variable}`)`]{#dom-global-fetch
.dfn .dfn-paneled .idl-code dfn-for="WindowOrWorkerGlobalScope"
dfn-type="method" export="" lt="fetch(input, init)|fetch(input)"} method
steps are:

1.  Let `p`{.variable} be [a new
    promise](https://webidl.spec.whatwg.org/#a-new-promise){#ref-for-a-new-promise②
    link-type="dfn"}.

2.  Let `requestObject`{.variable} be the result of invoking the initial
    value of [`Request`{.idl}](#request){#ref-for-request①⑥
    link-type="idl"} as constructor with `input`{.variable} and
    `init`{.variable} as arguments. If this throws an exception,
    [reject](https://webidl.spec.whatwg.org/#reject){#ref-for-reject①
    link-type="dfn"} `p`{.variable} with it and return `p`{.variable}.

3.  Let `request`{.variable} be `requestObject`{.variable}'s
    [request](#concept-request-request){#ref-for-concept-request-request②⑦
    link-type="dfn"}.

4.  If `requestObject`{.variable}'s
    [signal](#request-signal){#ref-for-request-signal⑥ link-type="dfn"}
    is
    [aborted](https://dom.spec.whatwg.org/#abortsignal-aborted){#ref-for-abortsignal-aborted
    link-type="dfn"}, then:

    1.  [Abort the `fetch()` call](#abort-fetch){#ref-for-abort-fetch
        link-type="dfn"} with `p`{.variable}, `request`{.variable},
        null, and `requestObject`{.variable}'s
        [signal](#request-signal){#ref-for-request-signal⑦
        link-type="dfn"}'s [abort
        reason](https://dom.spec.whatwg.org/#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason
        link-type="dfn"}.

    2.  Return `p`{.variable}.

5.  Let `globalObject`{.variable} be `request`{.variable}'s
    [client](#concept-request-client){#ref-for-concept-request-client③⓪
    link-type="dfn"}'s [global
    object](https://html.spec.whatwg.org/multipage/webappapis.html#concept-settings-object-global){#ref-for-concept-settings-object-global⑦
    link-type="dfn"}.

6.  If `globalObject`{.variable} is a
    [`ServiceWorkerGlobalScope`{.idl}](https://w3c.github.io/ServiceWorker/#serviceworkerglobalscope){#ref-for-serviceworkerglobalscope
    link-type="idl"} object, then set `request`{.variable}'s
    [service-workers
    mode](#request-service-workers-mode){#ref-for-request-service-workers-mode④
    link-type="dfn"} to \"`none`\".

7.  Let `responseObject`{.variable} be null.

8.  Let `relevantRealm`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑨④
    link-type="dfn"}'s [relevant
    realm](https://html.spec.whatwg.org/multipage/webappapis.html#concept-relevant-realm){#ref-for-concept-relevant-realm⑧
    link-type="dfn"}.

9.  Let `locallyAborted`{.variable} be false.

    This lets us reject promises with predictable timing, when the
    request to abort comes from the same thread as the call to fetch.

10. Let `controller`{.variable} be null.

11. [Add the following abort
    steps](https://dom.spec.whatwg.org/#abortsignal-add){#ref-for-abortsignal-add
    link-type="dfn"} to `requestObject`{.variable}'s
    [signal](#request-signal){#ref-for-request-signal⑧ link-type="dfn"}:

    1.  Set `locallyAborted`{.variable} to true.

    2.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert②⑦
        link-type="dfn"}: `controller`{.variable} is non-null.

    3.  [Abort](#fetch-controller-abort){#ref-for-fetch-controller-abort②
        link-type="dfn"} `controller`{.variable} with
        `requestObject`{.variable}'s
        [signal](#request-signal){#ref-for-request-signal⑨
        link-type="dfn"}'s [abort
        reason](https://dom.spec.whatwg.org/#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason①
        link-type="dfn"}.

    4.  [Abort the `fetch()` call](#abort-fetch){#ref-for-abort-fetch①
        link-type="dfn"} with `p`{.variable}, `request`{.variable},
        `responseObject`{.variable}, and `requestObject`{.variable}'s
        [signal](#request-signal){#ref-for-request-signal①⓪
        link-type="dfn"}'s [abort
        reason](https://dom.spec.whatwg.org/#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason②
        link-type="dfn"}.

12. Set `controller`{.variable} to the result of calling
    [fetch](#concept-fetch){#ref-for-concept-fetch②⑧ link-type="dfn"}
    given `request`{.variable} and
    [*processResponse*](#process-response){#ref-for-process-response①
    link-type="dfn"} given `response`{.variable} being these steps:

    1.  If `locallyAborted`{.variable} is true, then abort these steps.

    2.  If `response`{.variable}'s [aborted
        flag](#concept-response-aborted){#ref-for-concept-response-aborted②
        link-type="dfn"} is set, then:

        1.  Let `deserializedError`{.variable} be the result of
            [deserialize a serialized abort
            reason](#deserialize-a-serialized-abort-reason){#ref-for-deserialize-a-serialized-abort-reason①
            link-type="dfn"} given `controller`{.variable}'s [serialized
            abort
            reason](#fetch-controller-serialized-abort-reason){#ref-for-fetch-controller-serialized-abort-reason②
            link-type="dfn"} and `relevantRealm`{.variable}.

        2.  [Abort the `fetch()`
            call](#abort-fetch){#ref-for-abort-fetch② link-type="dfn"}
            with `p`{.variable}, `request`{.variable},
            `responseObject`{.variable}, and
            `deserializedError`{.variable}.

        3.  Abort these steps.

    3.  If `response`{.variable} is a [network
        error](#concept-network-error){#ref-for-concept-network-error⑥⓪
        link-type="dfn"}, then
        [reject](https://webidl.spec.whatwg.org/#reject){#ref-for-reject②
        link-type="dfn"} `p`{.variable} with a
        [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror②⑨
        link-type="idl"} and abort these steps.

    4.  Set `responseObject`{.variable} to the result of
        [creating](#response-create){#ref-for-response-create④
        link-type="dfn"} a
        [`Response`{.idl}](#response){#ref-for-response②⑤
        link-type="idl"} object, given `response`{.variable},
        \"`immutable`\", and `relevantRealm`{.variable}.

    5.  [Resolve](https://webidl.spec.whatwg.org/#resolve){#ref-for-resolve②
        link-type="dfn"} `p`{.variable} with
        `responseObject`{.variable}.

13. Return `p`{.variable}.
:::

::: {.algorithm algorithm="Abort the fetch() call"}
To [abort a `fetch()` call]{#abort-fetch .dfn .dfn-paneled
dfn-type="dfn" export="" lt="Abort the fetch() call"} with a
`promise`{.variable}, `request`{.variable}, `responseObject`{.variable},
and an `error`{.variable}:

1.  [Reject](https://webidl.spec.whatwg.org/#reject){#ref-for-reject③
    link-type="dfn"} `promise`{.variable} with `error`{.variable}.

    This is a no-op if `promise`{.variable} has already fulfilled.

2.  If `request`{.variable}'s
    [body](#concept-request-body){#ref-for-concept-request-body④⑧
    link-type="dfn"} is non-null and is
    [readable](https://streams.spec.whatwg.org/#readablestream-readable){#ref-for-readablestream-readable③
    link-type="dfn"}, then
    [cancel](https://streams.spec.whatwg.org/#readablestream-cancel){#ref-for-readablestream-cancel①
    link-type="dfn"} `request`{.variable}'s
    [body](#concept-request-body){#ref-for-concept-request-body④⑨
    link-type="dfn"} with `error`{.variable}.

3.  If `responseObject`{.variable} is null, then return.

4.  Let `response`{.variable} be `responseObject`{.variable}'s
    [response](#concept-response-response){#ref-for-concept-response-response①⑦
    link-type="dfn"}.

5.  If `response`{.variable}'s
    [body](#concept-response-body){#ref-for-concept-response-body②⑨
    link-type="dfn"} is non-null and is
    [readable](https://streams.spec.whatwg.org/#readablestream-readable){#ref-for-readablestream-readable④
    link-type="dfn"}, then
    [error](https://streams.spec.whatwg.org/#readablestream-error){#ref-for-readablestream-error②
    link-type="dfn"} `response`{.variable}'s
    [body](#concept-response-body){#ref-for-concept-response-body③⓪
    link-type="dfn"} with `error`{.variable}.
:::

### [5.7. ]{.secno}[Garbage collection]{.content}[](#garbage-collection){.self-link} {#garbage-collection .heading .settled level="5.7"}

The user agent may
[terminate](#fetch-controller-terminate){#ref-for-fetch-controller-terminate⑤
link-type="dfn"} an ongoing fetch if that termination is not observable
through script.

\"Observable through script\" means observable through
[`fetch()`](#dom-global-fetch){#ref-for-dom-global-fetch⑦ .idl-code
link-type="method"}'s arguments and return value. Other ways, such as
communicating with the server through a side-channel are not included.

The server being able to observe garbage collection has precedent, e.g.,
with
[`WebSocket`{.idl}](https://websockets.spec.whatwg.org/#websocket){#ref-for-websocket①
link-type="idl"} and
[`XMLHttpRequest`{.idl}](https://xhr.spec.whatwg.org/#xmlhttprequest){#ref-for-xmlhttprequest⑥
link-type="idl"} objects.

::: {#terminate-examples .example}
[](#terminate-examples){.self-link}

The user agent can terminate the fetch because the termination cannot be
observed.

``` {.lang-javascript .highlight}
fetch("https://www.example.com/")
```

The user agent cannot terminate the fetch because the termination can be
observed through the promise.

``` {.lang-javascript .highlight}
window.promise = fetch("https://www.example.com/")
```

The user agent can terminate the fetch because the associated body is
not observable.

``` {.lang-javascript .highlight}
window.promise = fetch("https://www.example.com/").then(res => res.headers)
```

The user agent can terminate the fetch because the termination cannot be
observed.

``` {.lang-javascript .highlight}
fetch("https://www.example.com/").then(res => res.body.getReader().closed)
```

The user agent cannot terminate the fetch because one can observe the
termination by registering a handler for the promise object.

``` {.lang-javascript .highlight}
window.promise = fetch("https://www.example.com/")
  .then(res => res.body.getReader().closed)
```

The user agent cannot terminate the fetch as termination would be
observable via the registered handler.

``` {.lang-javascript .highlight}
fetch("https://www.example.com/")
  .then(res => {
    res.body.getReader().closed.then(() => console.log("stream closed!"))
  })
```

(The above examples of non-observability assume that built-in properties
and functions, such as
[`body.getReader()`{.idl}](https://streams.spec.whatwg.org/#rs-get-reader){#ref-for-rs-get-reader
link-type="idl"}, have not been overwritten.)
:::

## [6. ]{.secno}[`data:` URLs]{.content}[](#data-urls){.self-link} {#data-urls .heading .settled level="6"}

For an informative description of `data:` URLs, see RFC 2397. This
section replaces that RFC's normative processing requirements to be
compatible with deployed content.
[\[RFC2397\]](#biblio-rfc2397 "The "data" URL scheme"){link-type="biblio"}

A [`data:` URL struct]{#data-url-struct .dfn .dfn-paneled dfn-type="dfn"
noexport=""} is a
[struct](https://infra.spec.whatwg.org/#struct){#ref-for-struct⑤
link-type="dfn"} that consists of a [MIME
type]{#data-url-struct-mime-type .dfn .dfn-paneled
dfn-for="data: URL struct" dfn-type="dfn" noexport=""} (a [MIME
type](https://mimesniff.spec.whatwg.org/#mime-type){#ref-for-mime-type⑤
link-type="dfn"}) and a [body]{#data-url-struct-body .dfn .dfn-paneled
dfn-for="data: URL struct" dfn-type="dfn" noexport=""} (a [byte
sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence③⓪
link-type="dfn"}).

::: {.algorithm algorithm="data: URL processor"}
The [`data:` URL processor]{#data-url-processor .dfn .dfn-paneled
dfn-type="dfn" export=""} takes a
[URL](https://url.spec.whatwg.org/#concept-url){#ref-for-concept-url②①
link-type="dfn"} `dataURL`{.variable} and then runs these steps:

1.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert②⑧
    link-type="dfn"}: `dataURL`{.variable}'s
    [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#ref-for-concept-url-scheme①⑥
    link-type="dfn"} is \"`data`\".

2.  Let `input`{.variable} be the result of running the [URL
    serializer](https://url.spec.whatwg.org/#concept-url-serializer){#ref-for-concept-url-serializer⑥
    link-type="dfn"} on `dataURL`{.variable} with [*exclude
    fragment*](https://url.spec.whatwg.org/#url-serializer-exclude-fragment){#ref-for-url-serializer-exclude-fragment②
    link-type="dfn"} set to true.

3.  Remove the leading \"`data:`\" from `input`{.variable}.

4.  Let `position`{.variable} point at the start of `input`{.variable}.

5.  Let `mimeType`{.variable} be the result of [collecting a sequence of
    code
    points](https://infra.spec.whatwg.org/#collect-a-sequence-of-code-points){#ref-for-collect-a-sequence-of-code-points⑧
    link-type="dfn"} that are not equal to U+002C (,), given
    `position`{.variable}.

6.  [Strip leading and trailing ASCII
    whitespace](https://infra.spec.whatwg.org/#strip-leading-and-trailing-ascii-whitespace){#ref-for-strip-leading-and-trailing-ascii-whitespace
    link-type="dfn"} from `mimeType`{.variable}.

    This will only remove U+0020 SPACE [code
    points](https://infra.spec.whatwg.org/#code-point){#ref-for-code-point①⓪
    link-type="dfn"}, if any.

7.  If `position`{.variable} is past the end of `input`{.variable}, then
    return failure.

8.  Advance `position`{.variable} by 1.

9.  Let `encodedBody`{.variable} be the remainder of `input`{.variable}.

10. Let `body`{.variable} be the
    [percent-decoding](https://url.spec.whatwg.org/#string-percent-decode){#ref-for-string-percent-decode
    link-type="dfn"} of `encodedBody`{.variable}.

11. If `mimeType`{.variable} ends with U+003B (;), followed by zero or
    more U+0020 SPACE, followed by an [ASCII
    case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#ref-for-ascii-case-insensitive①
    link-type="dfn"} match for \"`base64`\", then:

    1.  Let `stringBody`{.variable} be the [isomorphic
        decode](https://infra.spec.whatwg.org/#isomorphic-decode){#ref-for-isomorphic-decode③
        link-type="dfn"} of `body`{.variable}.

    2.  Set `body`{.variable} to the [forgiving-base64
        decode](https://infra.spec.whatwg.org/#forgiving-base64-decode){#ref-for-forgiving-base64-decode
        link-type="dfn"} of `stringBody`{.variable}.

    3.  If `body`{.variable} is failure, then return failure.

    4.  Remove the last 6 [code
        points](https://infra.spec.whatwg.org/#code-point){#ref-for-code-point①①
        link-type="dfn"} from `mimeType`{.variable}.

    5.  Remove trailing U+0020 SPACE [code
        points](https://infra.spec.whatwg.org/#code-point){#ref-for-code-point①②
        link-type="dfn"} from `mimeType`{.variable}, if any.

    6.  Remove the last U+003B (;) from `mimeType`{.variable}.

12. If `mimeType`{.variable} [starts
    with](https://infra.spec.whatwg.org/#string-starts-with){#ref-for-string-starts-with②
    link-type="dfn"} \"`;`\", then prepend \"`text/plain`\" to
    `mimeType`{.variable}.

13. Let `mimeTypeRecord`{.variable} be the result of
    [parsing](https://mimesniff.spec.whatwg.org/#parse-a-mime-type){#ref-for-parse-a-mime-type②
    link-type="dfn"} `mimeType`{.variable}.

14. If `mimeTypeRecord`{.variable} is failure, then set
    `mimeTypeRecord`{.variable} to `text/plain;charset=US-ASCII`.

15. Return a new [`data:` URL
    struct](#data-url-struct){#ref-for-data-url-struct link-type="dfn"}
    whose [MIME
    type](#data-url-struct-mime-type){#ref-for-data-url-struct-mime-type①
    link-type="dfn"} is `mimeTypeRecord`{.variable} and
    [body](#data-url-struct-body){#ref-for-data-url-struct-body①
    link-type="dfn"} is `body`{.variable}.
:::

## [Background reading]{.content}[](#background-reading){.self-link} {#background-reading .no-num .heading .settled}

*This section and its subsections are informative only.*

### [HTTP header layer division]{.content}[](#http-header-layer-division){#ref-for-http-header-layer-division① .self-link} {#http-header-layer-division .no-num .heading .settled .dfn-paneled dfn-type="dfn" noexport=""}

For the purposes of fetching, there is an API layer (HTML's `img`, CSS's
`background-image`), early fetch layer, service worker layer, and
network & cache layer. \``Accept`\` and \``Accept-Language`\` are set in
the early fetch layer (typically by the user agent). Most other headers
controlled by the user agent, such as \``Accept-Encoding`\`, \``Host`\`,
and \``Referer`\`, are set in the network & cache layer. Developers can
set headers either at the API layer or in the service worker layer
(typically through a [`Request`{.idl}](#request){#ref-for-request①⑦
link-type="idl"} object). Developers have almost no control over
[forbidden
request-headers](#forbidden-request-header){#ref-for-forbidden-request-header②
link-type="dfn"}, but can control \``Accept`\` and have the means to
constrain and omit \``Referer`\` for instance.

### [Atomic HTTP redirect handling]{.content}[](#atomic-http-redirect-handling){#ref-for-atomic-http-redirect-handling② .self-link} {#atomic-http-redirect-handling .no-num .heading .settled .dfn-paneled dfn-type="dfn" noexport=""}

Redirects (a [response](#concept-response){#ref-for-concept-response⑥⑧
link-type="dfn"} whose
[status](#concept-response-status){#ref-for-concept-response-status②⑥
link-type="dfn"} or [internal
response](#concept-internal-response){#ref-for-concept-internal-response①⑧
link-type="dfn"}'s (if any)
[status](#concept-response-status){#ref-for-concept-response-status②⑦
link-type="dfn"} is a [redirect
status](#redirect-status){#ref-for-redirect-status④ link-type="dfn"})
are not exposed to APIs. Exposing redirects might leak information not
otherwise available through a cross-site scripting attack.

[](#example-xss-redirect){.self-link}A fetch to
`https://example.org/auth` that includes a `Cookie` marked `HttpOnly`
could result in a redirect to
`https://other-origin.invalid/4af955781ea1c84a3b11`. This new URL
contains a secret. If we expose redirects that secret would be available
through a cross-site scripting attack.

### [Basic safe CORS protocol setup]{.content}[](#basic-safe-cors-protocol-setup){.self-link} {#basic-safe-cors-protocol-setup .no-num .heading .settled}

For resources where data is protected through IP authentication or a
firewall (unfortunately relatively common still), using the [CORS
protocol](#cors-protocol){#ref-for-cors-protocol①⑧ link-type="dfn"} is
**unsafe**. (This is the reason why the [CORS
protocol](#cors-protocol){#ref-for-cors-protocol①⑨ link-type="dfn"} had
to be invented.)

However, otherwise using the following
[header](#concept-header){#ref-for-concept-header⑤⑨ link-type="dfn"} is
**safe**:

``` {.lang-http .highlight}
Access-Control-Allow-Origin: *
```

Even if a resource exposes additional information based on cookie or
HTTP authentication, using the above
[header](#concept-header){#ref-for-concept-header⑥⓪ link-type="dfn"}
will not reveal it. It will share the resource with APIs such as
[`XMLHttpRequest`{.idl}](https://xhr.spec.whatwg.org/#xmlhttprequest){#ref-for-xmlhttprequest⑦
link-type="idl"}, much like it is already shared with `curl` and `wget`.

Thus in other words, if a resource cannot be accessed from a random
device connected to the web using `curl` and `wget` the aforementioned
[header](#concept-header){#ref-for-concept-header⑥① link-type="dfn"} is
not to be included. If it can be accessed however, it is perfectly fine
to do so.

### [CORS protocol and HTTP caches]{.content}[](#cors-protocol-and-http-caches){.self-link} {#cors-protocol-and-http-caches .no-num .heading .settled}

If [CORS protocol](#cors-protocol){#ref-for-cors-protocol②⓪
link-type="dfn"} requirements are more complicated than setting
\`[`Access-Control-Allow-Origin`](#http-access-control-allow-origin){#ref-for-http-access-control-allow-origin⑦
link-type="http-header"}\` to `*` or a static
[origin](https://html.spec.whatwg.org/multipage/browsers.html#concept-origin){#ref-for-concept-origin⑨
link-type="dfn"}, \``Vary`\` is to be used.
[\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}
[\[HTTP\]](#biblio-http "HTTP Semantics"){link-type="biblio"}
[\[HTTP-CACHING\]](#biblio-http-caching "HTTP Caching"){link-type="biblio"}

``` {#example-vary-origin .example}
Vary: Origin
```

In particular, consider what happens if \``Vary`\` is *not* used and a
server is configured to send
\`[`Access-Control-Allow-Origin`](#http-access-control-allow-origin){#ref-for-http-access-control-allow-origin⑧
link-type="http-header"}\` for a certain resource only in response to a
[CORS request](#cors-request){#ref-for-cors-request①② link-type="dfn"}.
When a user agent receives a response to a non-[CORS
request](#cors-request){#ref-for-cors-request①③ link-type="dfn"} for
that resource (for example, as the result of a [navigation
request](#navigation-request){#ref-for-navigation-request③
link-type="dfn"}), the response will lack
\`[`Access-Control-Allow-Origin`](#http-access-control-allow-origin){#ref-for-http-access-control-allow-origin⑨
link-type="http-header"}\` and the user agent will cache that response.
Then, if the user agent subsequently encounters a [CORS
request](#cors-request){#ref-for-cors-request①④ link-type="dfn"} for the
resource, it will use that cached response from the previous non-[CORS
request](#cors-request){#ref-for-cors-request①⑤ link-type="dfn"},
without
\`[`Access-Control-Allow-Origin`](#http-access-control-allow-origin){#ref-for-http-access-control-allow-origin①⓪
link-type="http-header"}\`.

But if \``Vary: Origin`\` is used in the same scenario described above,
it will cause the user agent to
[fetch](#concept-fetch){#ref-for-concept-fetch②⑨ link-type="dfn"} a
response that includes
\`[`Access-Control-Allow-Origin`](#http-access-control-allow-origin){#ref-for-http-access-control-allow-origin①①
link-type="http-header"}\`, rather than using the cached response from
the previous non-[CORS request](#cors-request){#ref-for-cors-request①⑥
link-type="dfn"} that lacks
\`[`Access-Control-Allow-Origin`](#http-access-control-allow-origin){#ref-for-http-access-control-allow-origin①②
link-type="http-header"}\`.

However, if
\`[`Access-Control-Allow-Origin`](#http-access-control-allow-origin){#ref-for-http-access-control-allow-origin①③
link-type="http-header"}\` is set to `*` or a static
[origin](https://html.spec.whatwg.org/multipage/browsers.html#concept-origin){#ref-for-concept-origin①⓪
link-type="dfn"} for a particular resource, then configure the server to
always send
\`[`Access-Control-Allow-Origin`](#http-access-control-allow-origin){#ref-for-http-access-control-allow-origin①④
link-type="http-header"}\` in responses for the resource --- for
non-[CORS requests](#cors-request){#ref-for-cors-request①⑦
link-type="dfn"} as well as [CORS
requests](#cors-request){#ref-for-cors-request①⑧ link-type="dfn"} ---
and do not use \``Vary`\`.

### []{#the-websocket-connection-is-established .bs-old-id}[]{#fail-the-websocket-connection .bs-old-id}[]{#websocket-opening-handshake .bs-old-id}[]{#websocket-connections .bs-old-id}[WebSockets]{.content}[](#websocket-protocol){.self-link} {#websocket-protocol .no-num .heading .settled}

As part of establishing a connection, the
[`WebSocket`{.idl}](https://websockets.spec.whatwg.org/#websocket){#ref-for-websocket②
link-type="idl"} object initiates a special kind of
[fetch](#concept-fetch){#ref-for-concept-fetch③⓪ link-type="dfn"} (using
a [request](#concept-request){#ref-for-concept-request①①⑤
link-type="dfn"} whose
[mode](#concept-request-mode){#ref-for-concept-request-mode③③
link-type="dfn"} is \"`websocket`\") which allows it to share in many
fetch policy decisions, such HTTP Strict Transport Security (HSTS).
Ultimately this results in fetch calling into WebSockets to obtain a
dedicated connection.
[\[WEBSOCKETS\]](#biblio-websockets "WebSockets Standard"){link-type="biblio"}
[\[HSTS\]](#biblio-hsts "HTTP Strict Transport Security (HSTS)"){link-type="biblio"}

Fetch used to define [obtain a WebSocket
connection](https://websockets.spec.whatwg.org/#concept-websocket-connection-obtain){#concept-websocket-connection-obtain
link-type="dfn"} and [establish a WebSocket
connection](https://websockets.spec.whatwg.org/#concept-websocket-establish){#concept-websocket-establish
link-type="dfn"} directly, but both are now defined in WebSockets.
[\[WEBSOCKETS\]](#biblio-websockets "WebSockets Standard"){link-type="biblio"}

## [Using fetch in other standards]{.content}[](#fetch-elsewhere){.self-link} {#fetch-elsewhere .no-num .heading .settled}

In its essence [fetching](#concept-fetch){#ref-for-concept-fetch③①
link-type="dfn"} is an exchange of a
[request](#concept-request){#ref-for-concept-request①①⑥ link-type="dfn"}
for a [response](#concept-response){#ref-for-concept-response⑥⑨
link-type="dfn"}. In reality it is rather complex mechanism for
standards to adopt and use correctly. This section aims to give some
advice.

Always ask domain experts for review.

This is a work in progress.

### [Setting up a request]{.content}[](#fetch-elsewhere-request){.self-link} {#fetch-elsewhere-request .heading .settled}

The first step in [fetching](#concept-fetch){#ref-for-concept-fetch③②
link-type="dfn"} is to create a
[request](#concept-request){#ref-for-concept-request①①⑦
link-type="dfn"}, and populate its
[items](https://infra.spec.whatwg.org/#struct-item){#ref-for-struct-item⑤
link-type="dfn"}.

Start by setting the
[request](#concept-request){#ref-for-concept-request①①⑧
link-type="dfn"}'s
[URL](#concept-request-url){#ref-for-concept-request-url①②
link-type="dfn"} and
[method](#concept-request-method){#ref-for-concept-request-method②⑥
link-type="dfn"}, as defined by HTTP. If your \``POST`\` or \``PUT`\`
[request](#concept-request){#ref-for-concept-request①①⑨ link-type="dfn"}
needs a body, you set
[request](#concept-request){#ref-for-concept-request①②⓪
link-type="dfn"}'s
[body](#concept-request-body){#ref-for-concept-request-body⑤⓪
link-type="dfn"} to a [byte
sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence③①
link-type="dfn"}, or to a new
[body](#concept-body){#ref-for-concept-body①④ link-type="dfn"} whose
[stream](#concept-body-stream){#ref-for-concept-body-stream①⑧
link-type="dfn"} is a
[`ReadableStream`{.idl}](https://streams.spec.whatwg.org/#readablestream){#ref-for-readablestream①③
link-type="idl"} you created.
[\[HTTP\]](#biblio-http "HTTP Semantics"){link-type="biblio"}

Choose your [request](#concept-request){#ref-for-concept-request①②①
link-type="dfn"}'s
[destination](#concept-request-destination){#ref-for-concept-request-destination②②
link-type="dfn"} using the guidance in the [destination
table](#destination-table).
[Destinations](#concept-request-destination){#ref-for-concept-request-destination②③
link-type="dfn"} affect Content Security Policy and have other
implications such as the
\`[`Sec-Fetch-Dest`](https://w3c.github.io/webappsec-fetch-metadata/#http-headerdef-sec-fetch-dest){#ref-for-http-headerdef-sec-fetch-dest
link-type="http-header"}\` header, so they are much more than
informative metadata. If a new feature requires a
[destination](#concept-request-destination){#ref-for-concept-request-destination②④
link-type="dfn"} that's not in the [destination
table](#destination-table), please [file an
issue](https://github.com/whatwg/fetch/issues/new?title=What%20destination%20should%20my%20feature%20use)
to discuss.
[\[CSP\]](#biblio-csp "Content Security Policy Level 3"){link-type="biblio"}

Set your [request](#concept-request){#ref-for-concept-request①②②
link-type="dfn"}'s
[client](#concept-request-client){#ref-for-concept-request-client③①
link-type="dfn"} to the [environment settings
object](https://html.spec.whatwg.org/multipage/webappapis.html#environment-settings-object){#ref-for-environment-settings-object①⓪
link-type="dfn"} you're operating in. Web-exposed APIs are generally
defined with Web IDL, for which every object that implements an
[interface](https://webidl.spec.whatwg.org/#dfn-interface){#ref-for-dfn-interface
link-type="dfn"} has a [relevant settings
object](https://html.spec.whatwg.org/multipage/webappapis.html#relevant-settings-object){#ref-for-relevant-settings-object③
link-type="dfn"} you can use. For example, a
[request](#concept-request){#ref-for-concept-request①②③ link-type="dfn"}
associated with an
[element](https://dom.spec.whatwg.org/#concept-element){#ref-for-concept-element
link-type="dfn"} would set the
[request](#concept-request){#ref-for-concept-request①②④
link-type="dfn"}'s
[client](#concept-request-client){#ref-for-concept-request-client③②
link-type="dfn"} to the element's [node
document](https://dom.spec.whatwg.org/#concept-node-document){#ref-for-concept-node-document
link-type="dfn"}'s [relevant settings
object](https://html.spec.whatwg.org/multipage/webappapis.html#relevant-settings-object){#ref-for-relevant-settings-object④
link-type="dfn"}. All features that are directly web-exposed by
JavaScript, HTML, CSS, or other
[`Document`{.idl}](https://dom.spec.whatwg.org/#document){#ref-for-document
link-type="idl"} subresources should have a
[client](#concept-request-client){#ref-for-concept-request-client③③
link-type="dfn"}.

If your [fetching](#concept-fetch){#ref-for-concept-fetch③③
link-type="dfn"} is not directly web-exposed, e.g., it is sent in the
background without relying on a current
[`Window`{.idl}](https://html.spec.whatwg.org/multipage/nav-history-apis.html#window){#ref-for-window③
link-type="idl"} or
[`Worker`{.idl}](https://html.spec.whatwg.org/multipage/workers.html#worker){#ref-for-worker
link-type="idl"}, leave
[request](#concept-request){#ref-for-concept-request①②⑤
link-type="dfn"}'s
[client](#concept-request-client){#ref-for-concept-request-client③④
link-type="dfn"} as null and set the
[request](#concept-request){#ref-for-concept-request①②⑥
link-type="dfn"}'s
[origin](#concept-request-origin){#ref-for-concept-request-origin②⑤
link-type="dfn"}, [policy
container](#concept-request-policy-container){#ref-for-concept-request-policy-container⑤
link-type="dfn"}, [service-workers
mode](#request-service-workers-mode){#ref-for-request-service-workers-mode⑤
link-type="dfn"}, and
[referrer](#concept-request-referrer){#ref-for-concept-request-referrer①⑧
link-type="dfn"} to appropriate values instead, e.g., by copying them
from the [environment settings
object](https://html.spec.whatwg.org/multipage/webappapis.html#environment-settings-object){#ref-for-environment-settings-object①①
link-type="dfn"} ahead of time. In these more advanced cases, make sure
the details of how your fetch handles Content Security Policy and
[referrer
policy](https://w3c.github.io/webappsec-referrer-policy/#referrer-policy){#ref-for-referrer-policy②
link-type="dfn"} are fleshed out. Also make sure you handle concurrency,
as callbacks (see [Invoking fetch and processing
responses](#fetch-elsewhere-fetch)) would be posted on a [parallel
queue](https://html.spec.whatwg.org/multipage/infrastructure.html#parallel-queue){#ref-for-parallel-queue⑥
link-type="dfn"}.
[\[REFERRER\]](#biblio-referrer "Referrer Policy"){link-type="biblio"}
[\[CSP\]](#biblio-csp "Content Security Policy Level 3"){link-type="biblio"}

Think through the way you intend to handle cross-origin resources. Some
features may only work in the [same
origin](https://html.spec.whatwg.org/multipage/browsers.html#same-origin){#ref-for-same-origin①②
link-type="dfn"}, in which case set your
[request](#concept-request){#ref-for-concept-request①②⑦
link-type="dfn"}'s
[mode](#concept-request-mode){#ref-for-concept-request-mode③④
link-type="dfn"} to \"`same-origin`\". Otherwise, new web-exposed
features should almost always set their
[mode](#concept-request-mode){#ref-for-concept-request-mode③⑤
link-type="dfn"} to \"`cors`\". If your feature is not web-exposed, or
you think there is another reason for it to fetch cross-origin resources
without CORS, please [file an
issue](https://github.com/whatwg/fetch/issues/new?title=Does%20my%20request%20require%20CORS)
to discuss.

For cross-origin requests, also determines if
[credentials](#credentials){#ref-for-credentials①⑨ link-type="dfn"} are
to be included with the requests, in which case set your
[request](#concept-request){#ref-for-concept-request①②⑧
link-type="dfn"}'s [credentials
mode](#concept-request-credentials-mode){#ref-for-concept-request-credentials-mode②⑤
link-type="dfn"} to \"`include`\".

Figure out if your fetch needs to be reported to Resource Timing, and
with which [initiator
type](#request-initiator-type){#ref-for-request-initiator-type④
link-type="dfn"}. By passing an [initiator
type](#request-initiator-type){#ref-for-request-initiator-type⑤
link-type="dfn"} to the
[request](#concept-request){#ref-for-concept-request①②⑨
link-type="dfn"}, reporting to Resource Timing will be done
automatically once the fetch is done and the
[response](#concept-response){#ref-for-concept-response⑦⓪
link-type="dfn"} is fully downloaded.
[\[RESOURCE-TIMING\]](#biblio-resource-timing "Resource Timing"){link-type="biblio"}

If your request requires additional HTTP headers, set its [header
list](#concept-request-header-list){#ref-for-concept-request-header-list⑤⓪
link-type="dfn"} to a [header
list](#concept-header-list){#ref-for-concept-header-list②④
link-type="dfn"} that contains those headers, e.g., «
(\``My-Header-Name`\`, \``My-Header-Value`\`) ». Sending custom headers
may have implications, such as requiring a [CORS-preflight
fetch](#cors-preflight-fetch-0){#ref-for-cors-preflight-fetch-0⑦
link-type="dfn"}, so handle with care.

If you want to override the default caching mechanism, e.g., disable
caching for this [request](#concept-request){#ref-for-concept-request①③⓪
link-type="dfn"}, set the request's [cache
mode](#concept-request-cache-mode){#ref-for-concept-request-cache-mode①⑦
link-type="dfn"} to a value other than \"`default`\".

Determine whether you want your request to support redirects. If you
don't, set its [redirect
mode](#concept-request-redirect-mode){#ref-for-concept-request-redirect-mode①③
link-type="dfn"} to \"`error`\".

Browse through the rest of the parameters for
[request](#concept-request){#ref-for-concept-request①③① link-type="dfn"}
to see if something else is relevant to you. The rest of the parameters
are used less frequently, often for special purposes, and they are
documented in detail in the [§ 2.2.5 Requests](#requests) section of
this standard.

### [Invoking fetch and processing responses]{.content}[](#fetch-elsewhere-fetch){.self-link} {#fetch-elsewhere-fetch .no-num .heading .settled}

Aside from a [request](#concept-request){#ref-for-concept-request①③②
link-type="dfn"} the [fetch](#concept-fetch){#ref-for-concept-fetch③④
link-type="dfn"} operation takes several optional arguments. For those
arguments that take an algorithm: the algorithm will be called from a
task (or in a [parallel
queue](https://html.spec.whatwg.org/multipage/infrastructure.html#parallel-queue){#ref-for-parallel-queue⑦
link-type="dfn"} if
[*useParallelQueue*](#fetch-useparallelqueue){#ref-for-fetch-useparallelqueue
link-type="dfn"} is true).

Once the [request](#concept-request){#ref-for-concept-request①③③
link-type="dfn"} is set up, to determine which algorithms to pass to
[fetch](#concept-fetch){#ref-for-concept-fetch③⑤ link-type="dfn"},
determine how you would like to process the
[response](#concept-response){#ref-for-concept-response⑦①
link-type="dfn"}, and in particular at what stage you would like to
receive a callback:

Upon completion

:   This is how most callers handle a
    [response](#concept-response){#ref-for-concept-response⑦②
    link-type="dfn"}, for example
    [scripts](https://html.spec.whatwg.org/multipage/webappapis.html#fetch-a-classic-script){#ref-for-fetch-a-classic-script
    link-type="dfn"} and [style
    resources](https://drafts.csswg.org/css-values-4/#fetch-a-style-resource){#ref-for-fetch-a-style-resource
    link-type="dfn"}. The
    [response](#concept-response){#ref-for-concept-response⑦③
    link-type="dfn"}'s
    [body](#concept-response-body){#ref-for-concept-response-body③①
    link-type="dfn"} is read in its entirety into a [byte
    sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence③②
    link-type="dfn"}, and then processed by the caller.

    To process a
    [response](#concept-response){#ref-for-concept-response⑦④
    link-type="dfn"} upon completion, pass an algorithm as the
    [*processResponseConsumeBody*](#process-response-end-of-body){#ref-for-process-response-end-of-body
    link-type="dfn"} argument of
    [fetch](#concept-fetch){#ref-for-concept-fetch③⑥ link-type="dfn"}.
    The given algorithm is passed a
    [response](#concept-response){#ref-for-concept-response⑦⑤
    link-type="dfn"} and an argument representing the fully read
    [body](#concept-response-body){#ref-for-concept-response-body③②
    link-type="dfn"} (of the
    [response](#concept-response){#ref-for-concept-response⑦⑥
    link-type="dfn"}'s [internal
    response](#concept-internal-response){#ref-for-concept-internal-response①⑨
    link-type="dfn"}). The second argument's values have the following
    meaning:

    null
    :   The [response](#concept-response){#ref-for-concept-response⑦⑦
        link-type="dfn"}'s
        [body](#concept-response-body){#ref-for-concept-response-body③③
        link-type="dfn"} is null, due to the response being a [network
        error](#concept-network-error){#ref-for-concept-network-error⑥①
        link-type="dfn"} or having a [null body
        status](#null-body-status){#ref-for-null-body-status③
        link-type="dfn"}.

    failure
    :   Attempting to [fully
        read](#body-fully-read){#ref-for-body-fully-read③
        link-type="dfn"} the contents of the
        [response](#concept-response){#ref-for-concept-response⑦⑧
        link-type="dfn"}'s
        [body](#concept-response-body){#ref-for-concept-response-body③④
        link-type="dfn"} failed, e.g., due to an I/O error.

    a [byte sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence③③ link-type="dfn"}

    :   [Fully reading](#body-fully-read){#ref-for-body-fully-read④
        link-type="dfn"} the contents of the
        [response](#concept-response){#ref-for-concept-response⑦⑨
        link-type="dfn"}'s [internal
        response](#concept-internal-response){#ref-for-concept-internal-response②⓪
        link-type="dfn"}'s
        [body](#concept-response-body){#ref-for-concept-response-body③⑤
        link-type="dfn"} succeeded.

        A [byte
        sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence③④
        link-type="dfn"} containing the full contents will be passed
        also for a
        [request](#concept-request){#ref-for-concept-request①③④
        link-type="dfn"} whose
        [mode](#concept-request-mode){#ref-for-concept-request-mode③⑥
        link-type="dfn"} is \"`no-cors`\". Callers have to be careful
        when handling such content, as it should not be accessible to
        the requesting
        [origin](https://html.spec.whatwg.org/multipage/browsers.html#concept-origin){#ref-for-concept-origin①①
        link-type="dfn"}. For example, the caller may use contents of a
        \"`no-cors`\"
        [response](#concept-response){#ref-for-concept-response⑧⓪
        link-type="dfn"} to display image contents directly to the user,
        but those image contents should not be directly exposed to
        scripts in the embedding document.

    ::: {#example-callback-upon-completion .example}
    [](#example-callback-upon-completion){.self-link}
    1.  Let `request`{.variable} be a
        [request](#concept-request){#ref-for-concept-request①③⑤
        link-type="dfn"} whose
        [URL](#concept-request-url){#ref-for-concept-request-url①③
        link-type="dfn"} is `https://stuff.example.com/` and
        [client](#concept-request-client){#ref-for-concept-request-client③⑤
        link-type="dfn"} is
        [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑨⑤
        link-type="dfn"}'s [relevant settings
        object](https://html.spec.whatwg.org/multipage/webappapis.html#relevant-settings-object){#ref-for-relevant-settings-object⑤
        link-type="dfn"}.

    2.  [Fetch](#concept-fetch){#ref-for-concept-fetch③⑦
        link-type="dfn"} `request`{.variable}, with
        [*processResponseConsumeBody*](#process-response-end-of-body){#ref-for-process-response-end-of-body①
        link-type="dfn"} set to the following steps given a
        [response](#concept-response){#ref-for-concept-response⑧①
        link-type="dfn"} `response`{.variable} and null, failure, or a
        [byte
        sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence③⑤
        link-type="dfn"} `contents`{.variable}:

        1.  If `contents`{.variable} is null or failure, then present an
            error to the user.

        2.  Otherwise, parse `contents`{.variable} considering the
            metadata from `response`{.variable}, and perform your own
            operations on it.
    :::

Headers first, then chunk-by-chunk

:   In some cases, for example when playing video or progressively
    loading images, callers might want to stream the response, and
    process it one chunk at a time. The
    [response](#concept-response){#ref-for-concept-response⑧②
    link-type="dfn"} is handed over to the fetch caller once the headers
    are processed, and the caller continues from there.

    To process a
    [response](#concept-response){#ref-for-concept-response⑧③
    link-type="dfn"} chunk-by-chunk, pass an algorithm to the
    [*processResponse*](#process-response){#ref-for-process-response②
    link-type="dfn"} argument of
    [fetch](#concept-fetch){#ref-for-concept-fetch③⑧ link-type="dfn"}.
    The given algorithm is passed a
    [response](#concept-response){#ref-for-concept-response⑧④
    link-type="dfn"} when the response's headers have been received and
    is responsible for reading the
    [response](#concept-response){#ref-for-concept-response⑧⑤
    link-type="dfn"}'s
    [body](#concept-response-body){#ref-for-concept-response-body③⑥
    link-type="dfn"}'s
    [stream](#concept-body-stream){#ref-for-concept-body-stream①⑨
    link-type="dfn"} in order to download the rest of the response. For
    convenience, you may also pass an algorithm to the
    [*processResponseEndOfBody*](#fetch-processresponseendofbody){#ref-for-fetch-processresponseendofbody
    link-type="dfn"} argument, which is called once you have finished
    fully reading the response and its
    [body](#concept-response-body){#ref-for-concept-response-body③⑦
    link-type="dfn"}. Note that unlike
    [*processResponseConsumeBody*](#process-response-end-of-body){#ref-for-process-response-end-of-body②
    link-type="dfn"}, passing the
    [*processResponse*](#process-response){#ref-for-process-response③
    link-type="dfn"} or
    [*processResponseEndOfBody*](#fetch-processresponseendofbody){#ref-for-fetch-processresponseendofbody①
    link-type="dfn"} arguments does not guarantee that the response will
    be fully read, and callers are responsible to read it themselves.

    The
    [*processResponse*](#process-response){#ref-for-process-response④
    link-type="dfn"} argument is also useful for handling the
    [response](#concept-response){#ref-for-concept-response⑧⑥
    link-type="dfn"}'s [header
    list](#concept-response-header-list){#ref-for-concept-response-header-list③⑧
    link-type="dfn"} and
    [status](#concept-response-status){#ref-for-concept-response-status②⑧
    link-type="dfn"} without handling the
    [body](#concept-response-body){#ref-for-concept-response-body③⑧
    link-type="dfn"} at all. This is used, for example, when handling
    responses that do not have an [ok
    status](#ok-status){#ref-for-ok-status④ link-type="dfn"}.

    ::: {#example-callback-chunk-by-chunk .example}
    [](#example-callback-chunk-by-chunk){.self-link}
    1.  Let `request`{.variable} be a
        [request](#concept-request){#ref-for-concept-request①③⑥
        link-type="dfn"} whose
        [URL](#concept-request-url){#ref-for-concept-request-url①④
        link-type="dfn"} is `https://stream.example.com/` and
        [client](#concept-request-client){#ref-for-concept-request-client③⑥
        link-type="dfn"} is
        [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑨⑥
        link-type="dfn"}'s [relevant settings
        object](https://html.spec.whatwg.org/multipage/webappapis.html#relevant-settings-object){#ref-for-relevant-settings-object⑥
        link-type="dfn"}.

    2.  [Fetch](#concept-fetch){#ref-for-concept-fetch③⑨
        link-type="dfn"} `request`{.variable}, with
        [*processResponse*](#process-response){#ref-for-process-response⑤
        link-type="dfn"} set to the following steps given a
        [response](#concept-response){#ref-for-concept-response⑧⑦
        link-type="dfn"} `response`{.variable}:

        1.  If `response`{.variable} is a [network
            error](#concept-network-error){#ref-for-concept-network-error⑥②
            link-type="dfn"}, then present an error to the user.

        2.  Otherwise, if `response`{.variable}'s
            [status](#concept-response-status){#ref-for-concept-response-status②⑨
            link-type="dfn"} is not an [ok
            status](#ok-status){#ref-for-ok-status⑤ link-type="dfn"},
            present some fallback value to the user.

        3.  Otherwise, [get a
            reader](https://streams.spec.whatwg.org/#readablestream-get-a-reader){#ref-for-readablestream-get-a-reader②
            link-type="dfn"} for
            [response](#concept-response){#ref-for-concept-response⑧⑧
            link-type="dfn"}'s
            [body](#concept-response-body){#ref-for-concept-response-body③⑨
            link-type="dfn"}'s
            [stream](#concept-body-stream){#ref-for-concept-body-stream②⓪
            link-type="dfn"}, and process in an appropriate way for the
            MIME type identified by [extracting a MIME
            type](#concept-header-extract-mime-type){#ref-for-concept-header-extract-mime-type⑨
            link-type="dfn"} from `response`{.variable}'s [headers
            list](#concept-response-header-list){#ref-for-concept-response-header-list③⑨
            link-type="dfn"}.
    :::

Ignore the response

:   In some cases, there is no need for a
    [response](#concept-response){#ref-for-concept-response⑧⑨
    link-type="dfn"} at all, e.g., in the case of
    [`navigator.sendBeacon()`{.idl}](https://w3c.github.io/beacon/#dom-navigator-sendbeacon){#ref-for-dom-navigator-sendbeacon
    link-type="idl"}. Processing a response and passing callbacks to
    [fetch](#concept-fetch){#ref-for-concept-fetch④⓪ link-type="dfn"} is
    optional, so omitting the callback would
    [fetch](#concept-fetch){#ref-for-concept-fetch④① link-type="dfn"}
    without expecting a response. In such cases, the
    [response](#concept-response){#ref-for-concept-response⑨⓪
    link-type="dfn"}'s
    [body](#concept-response-body){#ref-for-concept-response-body④⓪
    link-type="dfn"}'s
    [stream](#concept-body-stream){#ref-for-concept-body-stream②①
    link-type="dfn"} will be discarded, and the caller does not have to
    worry about downloading the contents unnecessarily.

    [](#example-no-callback){.self-link}[Fetch](#concept-fetch){#ref-for-concept-fetch④②
    link-type="dfn"} a
    [request](#concept-request){#ref-for-concept-request①③⑦
    link-type="dfn"} whose
    [URL](#concept-request-url){#ref-for-concept-request-url①⑤
    link-type="dfn"} is `https://fire-and-forget.example.com/`,
    [method](#concept-request-method){#ref-for-concept-request-method②⑦
    link-type="dfn"} is \``POST`\`, and
    [client](#concept-request-client){#ref-for-concept-request-client③⑦
    link-type="dfn"} is
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑨⑦
    link-type="dfn"}'s [relevant settings
    object](https://html.spec.whatwg.org/multipage/webappapis.html#relevant-settings-object){#ref-for-relevant-settings-object⑦
    link-type="dfn"}.

Apart from the callbacks to handle responses,
[fetch](#concept-fetch){#ref-for-concept-fetch④③ link-type="dfn"}
accepts additional callbacks for advanced cases.
[*processEarlyHintsResponse*](#fetch-processearlyhintsresponse){#ref-for-fetch-processearlyhintsresponse
link-type="dfn"} is intended specifically for
[responses](#concept-response){#ref-for-concept-response⑨①
link-type="dfn"} whose
[status](#concept-response-status){#ref-for-concept-response-status③⓪
link-type="dfn"} is 103, and is currently handled only by navigations.
[*processRequestBodyChunkLength*](#process-request-body){#ref-for-process-request-body
link-type="dfn"} and
[*processRequestEndOfBody*](#process-request-end-of-body){#ref-for-process-request-end-of-body
link-type="dfn"} notify the caller of request body uploading progress.

Note that the [fetch](#concept-fetch){#ref-for-concept-fetch④④
link-type="dfn"} operation starts in the same thread from which it was
called, and then breaks off to run its internal operations [in
parallel](https://html.spec.whatwg.org/multipage/infrastructure.html#in-parallel){#ref-for-in-parallel⑧
link-type="dfn"}. The aforementioned callbacks are posted to a given
[event
loop](https://html.spec.whatwg.org/multipage/webappapis.html#event-loop){#ref-for-event-loop
link-type="dfn"} which is, by default, the
[client](#concept-request-client){#ref-for-concept-request-client③⑧
link-type="dfn"}'s [global
object](https://html.spec.whatwg.org/multipage/webappapis.html#concept-settings-object-global){#ref-for-concept-settings-object-global⑧
link-type="dfn"}. To process responses [in
parallel](https://html.spec.whatwg.org/multipage/infrastructure.html#in-parallel){#ref-for-in-parallel⑨
link-type="dfn"} and handle interactions with the main thread by
yourself, [fetch](#concept-fetch){#ref-for-concept-fetch④⑤
link-type="dfn"} with
[*useParallelQueue*](#fetch-useparallelqueue){#ref-for-fetch-useparallelqueue①
link-type="dfn"} set to true.

### [Manipulating an ongoing fetch]{.content}[](#fetch-elsewhere-ongoing){.self-link} {#fetch-elsewhere-ongoing .heading .settled}

To manipulate a [fetch](#concept-fetch){#ref-for-concept-fetch④⑥
link-type="dfn"} operation that has already started, use the [fetch
controller](#fetch-controller){#ref-for-fetch-controller⑧
link-type="dfn"} returned by calling
[fetch](#concept-fetch){#ref-for-concept-fetch④⑦ link-type="dfn"}. For
example, you may
[abort](#fetch-controller-abort){#ref-for-fetch-controller-abort③
link-type="dfn"} the [fetch
controller](#fetch-controller){#ref-for-fetch-controller⑨
link-type="dfn"} due the user or page logic, or
[terminate](#fetch-controller-terminate){#ref-for-fetch-controller-terminate⑥
link-type="dfn"} it due to browser-internal circumstances.

In addition to terminating and aborting, callers may [report
timing](#finalize-and-report-timing){#ref-for-finalize-and-report-timing
link-type="dfn"} if this was not done automatically by passing the
[initiator
type](#request-initiator-type){#ref-for-request-initiator-type⑥
link-type="dfn"}, or [extract full timing
info](#extract-full-timing-info){#ref-for-extract-full-timing-info
link-type="dfn"} and handle it on the caller side (this is done only by
navigations). The [fetch
controller](#fetch-controller){#ref-for-fetch-controller①⓪
link-type="dfn"} is also used to [process the next manual
redirect](#fetch-controller-process-the-next-manual-redirect){#ref-for-fetch-controller-process-the-next-manual-redirect
link-type="dfn"} for
[requests](#concept-request){#ref-for-concept-request①③⑧
link-type="dfn"} with [redirect
mode](#concept-request-redirect-mode){#ref-for-concept-request-redirect-mode①④
link-type="dfn"} set to \"`manual`\".

## [Acknowledgments]{.content}[](#acknowledgments){.self-link} {#acknowledgments .no-num .heading .settled}

Thanks to Adam Barth, Adam Lavin, Alan Jeffrey, Alexey Proskuryakov,
Andreas Kling, Andrés Gutiérrez, Andrew Sutherland, Andrew Williams,
Ángel González, Anssi Kostiainen, Arkadiusz Michalski, Arne Johannessen,
Artem Skoretskiy, Arthur Barstow, Arthur Sonzogni, Asanka Herath, Axel
Rauschmayer, Ben Kelly, Benjamin Gruenbaum, Benjamin Hawkes-Lewis, Bert
Bos, Björn Höhrmann, Boris Zbarsky, Brad Hill, Brad Porter, Bryan Smith,
Caitlin Potter, Cameron McCormack, Carlo Cannas, 白丞祐 (Cheng-You Bai),
Chirag S Kumar, Chris Needham, Chris Rebert, Clement Pellerin, Collin
Jackson, Daniel Robertson, Daniel Veditz, Dave Tapuska, David Benjamin,
David Håsäther, David Orchard, Dean Jackson, Devdatta Akhawe, Domenic
Denicola, Dominic Farolino, Dominique Hazaël-Massieux, Doug Turner,
Douglas Creager, Eero Häkkinen, Ehsan Akhgari, Emily Stark, Eric
Lawrence, Eric Orth, Feng Yu, François Marier, Frank Ellerman, Frederick
Hirsch, Frederik Braun, Gary Blackwood, Gavin Carothers, Glenn Maynard,
Graham Klyne, Gregory Terzian, Guohui Deng(邓国辉), Hal Lockhart,
Hallvord R. M. Steen, Harris Hancock, Henri Sivonen, Henry Story,
Hiroshige Hayashizaki, Honza Bambas, Ian Hickson, Ilya Grigorik,
isonmad, Jake Archibald, James Graham, Jamie Mansfield, Janusz Majnert,
Jeena Lee, Jeff Carpenter, Jeff Hodges, Jeffrey Yasskin, Jensen
Chappell, Jeremy Roman, Jesse M. Heines, Jianjun Chen, Jinho Bang,
Jochen Eisinger, John Wilander, Jonas Sicking, Jonathan Kingston,
Jonathan Watt, 최종찬 (Jongchan Choi), Jordan Stephens, Jörn Zaefferer,
Joseph Pecoraro, Josh Matthews, jub0bs, Julian Krispel-Samsel, Julian
Reschke, 송정기 (Jungkee Song), Jussi Kalliokoski, Jxck, Kagami Sascha
Rosylight, Keith Yeung, Kenji Baheux, Lachlan Hunt, Larry Masinter, Liam
Brummitt, Linus Groh, Louis Ryan, Luca Casonato, Lucas Gonze, Łukasz
Anforowicz, 呂康豪 (Kang-Hao Lu), Maciej Stachowiak, Malisa, Manfred
Stock, Manish Goregaokar, Marc Silbey, Marcos Caceres, Marijn
Kruisselbrink, Mark Nottingham, Mark S. Miller, Martin Dürst, Martin
O'Neal, Martin Thomson, Matt Andrews, Matt Falkenhagen, Matt Menke, Matt
Oshry, Matt Seddon, Matt Womer, Mhano Harkness, Michael Ficarra, Michael
Kohler, Michael™ Smith, Mike Pennisi, Mike West, Mohamed Zergaoui,
Mohammed Zubair Ahmed, Moritz Kneilmann, Ms2ger, Nico Schlömer, Nicolás
Peña Moreno, Nidhi Jaju, Nikhil Marathe, Nikki Bee, Nikunj Mehta, Noam
Rosenthal, Odin Hørthe Omdal, Olli Pettay, Ondřej Žára, O. Opsec,
Patrick Meenan, Perry Jiang, Philip Jägenstedt, R. Auburn, Raphael Kubo
da Costa, Robert Linder, Rondinelly, Rory Hewitt, Ross A. Baker, Ryan
Sleevi, Sam Atkins, Samy Kamkar, Sébastien Cevey, Sendil Kumar N,
Shao-xuan Kang, Sharath Udupa, Shivakumar Jagalur Matt, Shivani Sharma,
Sigbjørn Finne, Simon Pieters, Simon Sapin, Simon Wülker, Srirama
Chandra Sekhar Mogali, Stephan Paul, Steven Salat, Sunava Dutta, Surya
Ismail, Tab Atkins-Bittner, Takashi Toyoshima, 吉野剛史 (Takeshi
Yoshino), Thomas Roessler, Thomas Steiner, Thomas Wisniewski, Tiancheng
\"Timothy\" Gu, Tobie Langel, Tom Schuster, Tomás Aparicio,
triple-underscore, 保呂毅 (Tsuyoshi Horo), Tyler Close, Ujjwal Sharma,
Vignesh Shanmugam, Vladimir Dzhuvinov, Wayne Carr, Xabier Rodríguez,
Yehuda Katz, Yoav Weiss, Youenn Fablet, Yoichi Osato, 平野裕 (Yutaka
Hirano), and Zhenbin Xu for being awesome.

This standard is written by [Anne van
Kesteren](https://annevankesteren.nl/){lang="nl"}
([Apple](https://www.apple.com/), <annevk@annevk.nl>).

## [Intellectual property rights]{.content}[](#ipr){.self-link} {#ipr .no-num .heading .settled}

Copyright © WHATWG (Apple, Google, Mozilla, Microsoft). This work is
licensed under a [Creative Commons Attribution 4.0 International
License](https://creativecommons.org/licenses/by/4.0/){rel="license"}.
To the extent portions of it are incorporated into source code, such
portions in the source code are licensed under the [BSD 3-Clause
License](https://opensource.org/licenses/BSD-3-Clause){rel="license"}
instead.

This is the Living Standard. Those interested in the patent-review
version should view the [Living Standard Review
Draft](/review-drafts/2024-12/).
