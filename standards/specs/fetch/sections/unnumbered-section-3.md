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

