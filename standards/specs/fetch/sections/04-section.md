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

