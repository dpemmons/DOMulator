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

