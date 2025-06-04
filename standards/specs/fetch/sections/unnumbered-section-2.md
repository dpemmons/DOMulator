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

