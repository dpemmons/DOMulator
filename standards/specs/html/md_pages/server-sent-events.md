HTML Standard

::: {#head .head .with-buttons}
[![WHATWG](https://resources.whatwg.org/logo.svg){.darkmode-aware
width="100" crossorigin="" height="100"}](https://whatwg.org/){.logo}

# HTML {#html .allcaps}

Living Standard --- Last Updated [2 June 2025]{.pubdate}
:::

[← 9 Communication](comms.html) --- [Table of Contents](./) --- [9.3
Cross-document messaging →](web-messaging.html)

1.  ::: {#toc-comms}
    1.  [[9.2]{.secno} Server-sent
        events](server-sent-events.html#server-sent-events)
        1.  [[9.2.1]{.secno}
            Introduction](server-sent-events.html#server-sent-events-intro)
        2.  [[9.2.2]{.secno} The `EventSource`
            interface](server-sent-events.html#the-eventsource-interface)
        3.  [[9.2.3]{.secno} Processing
            model](server-sent-events.html#sse-processing-model)
        4.  [[9.2.4]{.secno} The \``Last-Event-ID`\`
            header](server-sent-events.html#the-last-event-id-header)
        5.  [[9.2.5]{.secno} Parsing an event
            stream](server-sent-events.html#parsing-an-event-stream)
        6.  [[9.2.6]{.secno} Interpreting an event
            stream](server-sent-events.html#event-stream-interpretation)
        7.  [[9.2.7]{.secno} Authoring
            notes](server-sent-events.html#authoring-notes)
        8.  [[9.2.8]{.secno} Connectionless push and other
            features](server-sent-events.html#eventsource-push)
        9.  [[9.2.9]{.secno} Garbage
            collection](server-sent-events.html#garbage-collection)
        10. [[9.2.10]{.secno} Implementation
            advice](server-sent-events.html#implementation-advice)
    :::

### [9.2]{.secno} [Server-sent events]{.dfn}[](#server-sent-events){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Server-sent_events](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events "Traditionally, a web page has to send a request to the server to receive new data; that is, the page requests data from the server. With server-sent events, it's possible for a server to send new data to a web page at any time, by pushing messages to the web page. These incoming messages can be treated as Events + data inside the web page.")

Support in all current engines.

::: support
[Firefox6+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome6+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android45+]{.firefox_android .yes}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

#### [9.2.1]{.secno} Introduction[](#server-sent-events-intro){.self-link} {#server-sent-events-intro}

*This section is non-normative.*

To enable servers to push data to web pages over HTTP or using dedicated
server-push protocols, this specification introduces the
[`EventSource`{#server-sent-events-intro:eventsource}](#eventsource)
interface.

Using this API consists of creating an
[`EventSource`{#server-sent-events-intro:eventsource-2}](#eventsource)
object and registering an event listener.

``` js
var source = new EventSource('updates.cgi');
source.onmessage = function (event) {
  alert(event.data);
};
```

On the server-side, the script (\"`updates.cgi`\" in this case) sends
messages in the following form, with the
[`text/event-stream`{#server-sent-events-intro:text/event-stream}](iana.html#text/event-stream)
MIME type:

    data: This is the first message.

    data: This is the second message, it
    data: has two lines.

    data: This is the third message.

------------------------------------------------------------------------

Authors can separate events by using different event types. Here is a
stream that has two event types, \"add\" and \"remove\":

    event: add
    data: 73857293

    event: remove
    data: 2153

    event: add
    data: 113411

The script to handle such a stream would look like this (where
`addHandler` and `removeHandler` are functions that take one argument,
the event):

``` js
var source = new EventSource('updates.cgi');
source.addEventListener('add', addHandler, false);
source.addEventListener('remove', removeHandler, false);
```

The default event type is \"message\".

Event streams are always decoded as UTF-8. There is no way to specify
another character encoding.

------------------------------------------------------------------------

Event stream requests can be redirected using HTTP 301 and 307 redirects
as with normal HTTP requests. Clients will reconnect if the connection
is closed; a client can be told to stop reconnecting using the HTTP 204
No Content response code.

Using this API rather than emulating it using
[`XMLHttpRequest`{#server-sent-events-intro:xmlhttprequest}](https://xhr.spec.whatwg.org/#xmlhttprequest){x-internal="xmlhttprequest"}
or an
[`iframe`{#server-sent-events-intro:the-iframe-element}](iframe-embed-object.html#the-iframe-element)
allows the user agent to make better use of network resources in cases
where the user agent implementer and the network operator are able to
coordinate in advance. Amongst other benefits, this can result in
significant savings in battery life on portable devices. This is
discussed further in the section below on [connectionless
push](#eventsource-push).

#### [9.2.2]{.secno} The [`EventSource`{#the-eventsource-interface:eventsource}](#eventsource) interface[](#the-eventsource-interface){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[EventSource](https://developer.mozilla.org/en-US/docs/Web/API/EventSource "The EventSource interface is web content's interface to server-sent events.")

Support in all current engines.

::: support
[Firefox6+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome6+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android45+]{.firefox_android .yes}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

``` idl
[Exposed=(Window,Worker)]
interface EventSource : EventTarget {
  constructor(USVString url, optional EventSourceInit eventSourceInitDict = {});

  readonly attribute USVString url;
  readonly attribute boolean withCredentials;

  // ready state
  const unsigned short CONNECTING = 0;
  const unsigned short OPEN = 1;
  const unsigned short CLOSED = 2;
  readonly attribute unsigned short readyState;

  // networking
  attribute EventHandler onopen;
  attribute EventHandler onmessage;
  attribute EventHandler onerror;
  undefined close();
};

dictionary EventSourceInit {
  boolean withCredentials = false;
};
```

Each
[`EventSource`{#the-eventsource-interface:eventsource-2}](#eventsource)
object has the following associated with it:

- A [url]{#concept-eventsource-url .dfn dfn-for="EventSource"} (a [URL
  record](https://url.spec.whatwg.org/#concept-url){#the-eventsource-interface:url-record
  x-internal="url-record"}). Set during construction.

- A [request]{#concept-event-stream-request .dfn}. This must initially
  be null.

- A [reconnection time]{#concept-event-stream-reconnection-time .dfn},
  in milliseconds. This must initially be an
  [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#the-eventsource-interface:implementation-defined
  x-internal="implementation-defined"} value, probably in the region of
  a few seconds.

- A [last event ID string]{#concept-event-stream-last-event-id .dfn}.
  This must initially be the empty string.

Apart from
[url](#concept-eventsource-url){#the-eventsource-interface:concept-eventsource-url}
these are not currently exposed on the
[`EventSource`{#the-eventsource-interface:eventsource-3}](#eventsource)
object.

`source`{.variable}` = new `[`EventSource`](#dom-eventsource){#dom-eventsource-dev}`( ``url`{.variable}` [, { `[`withCredentials`](#dom-eventsourceinit-withcredentials){#the-eventsource-interface:dom-eventsourceinit-withcredentials}`: true } ])`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[EventSource/EventSource](https://developer.mozilla.org/en-US/docs/Web/API/EventSource/EventSource "The EventSource() constructor returns a newly-created EventSource, which represents a remote resource.")

Support in all current engines.

::: support
[Firefox6+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome6+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android45+]{.firefox_android .yes}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12+]{.opera_android .yes}
:::
::::
:::::

Creates a new
[`EventSource`{#the-eventsource-interface:eventsource-4}](#eventsource)
object.

`url`{.variable} is a string giving the
[URL](https://url.spec.whatwg.org/#concept-url){#the-eventsource-interface:url
x-internal="url"} that will provide the event stream.

Setting
[`withCredentials`{#the-eventsource-interface:dom-eventsourceinit-withcredentials-2}](#dom-eventsourceinit-withcredentials)
to true will set the [credentials
mode](https://fetch.spec.whatwg.org/#concept-request-credentials-mode){#the-eventsource-interface:concept-request-credentials-mode
x-internal="concept-request-credentials-mode"} for connection requests
to `url`{.variable} to \"`include`\".

`source`{.variable}`.`[`close`](#dom-eventsource-close){#dom-eventsource-close-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[EventSource/close](https://developer.mozilla.org/en-US/docs/Web/API/EventSource/close "The close() method of the EventSource interface closes the connection, if one is made, and sets the EventSource.readyState attribute to 2 (closed).")

Support in all current engines.

::: support
[Firefox6+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome6+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android45+]{.firefox_android .yes}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12+]{.opera_android .yes}
:::
::::
:::::

Aborts any instances of the
[fetch](https://fetch.spec.whatwg.org/#concept-fetch){#the-eventsource-interface:concept-fetch
x-internal="concept-fetch"} algorithm started for this
[`EventSource`{#the-eventsource-interface:eventsource-5}](#eventsource)
object, and sets the
[`readyState`{#the-eventsource-interface:dom-eventsource-readystate-2}](#dom-eventsource-readystate)
attribute to
[`CLOSED`{#the-eventsource-interface:dom-eventsource-closed-2}](#dom-eventsource-closed).

`source`{.variable}`.`[`url`](#dom-eventsource-url){#dom-eventsource-url-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[EventSource/url](https://developer.mozilla.org/en-US/docs/Web/API/EventSource/url "The url read-only property of the EventSource interface returns a string representing the URL of the source.")

Support in all current engines.

::: support
[Firefox6+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome18+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android45+]{.firefox_android .yes}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12+]{.opera_android .yes}
:::
::::
:::::

Returns the [URL providing the event
stream](#concept-eventsource-url){#the-eventsource-interface:concept-eventsource-url-2}.

`source`{.variable}`.`[`withCredentials`](#dom-eventsource-withcredentials){#dom-eventsource-withcredentials-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[EventSource/withCredentials](https://developer.mozilla.org/en-US/docs/Web/API/EventSource/withCredentials "The withCredentials read-only property of the EventSource interface returns a boolean value indicating whether the EventSource object was instantiated with CORS credentials set.")

Support in all current engines.

::: support
[Firefox6+]{.firefox .yes}[Safari7+]{.safari .yes}[Chrome26+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android45+]{.firefox_android .yes}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12+]{.opera_android .yes}
:::
::::
:::::

Returns true if the [credentials
mode](https://fetch.spec.whatwg.org/#concept-request-credentials-mode){#the-eventsource-interface:concept-request-credentials-mode-2
x-internal="concept-request-credentials-mode"} for connection requests
to the [URL providing the event
stream](#concept-eventsource-url){#the-eventsource-interface:concept-eventsource-url-3}
is set to \"`include`\", and false otherwise.

`source`{.variable}`.`[`readyState`](#dom-eventsource-readystate){#dom-eventsource-readystate-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[EventSource/readyState](https://developer.mozilla.org/en-US/docs/Web/API/EventSource/readyState "The readyState read-only property of the EventSource interface returns a number representing the state of the connection.")

Support in all current engines.

::: support
[Firefox6+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome6+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android45+]{.firefox_android .yes}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12+]{.opera_android .yes}
:::
::::
:::::

Returns the state of this
[`EventSource`{#the-eventsource-interface:eventsource-6}](#eventsource)
object\'s connection. It can have the values described below.

The
[`EventSource(``url`{.variable}`, ``eventSourceInitDict`{.variable}`)`]{#dom-eventsource
.dfn dfn-for="EventSource" dfn-type="constructor"} constructor, when
invoked, must run these steps:

1.  Let `ev`{.variable} be a new
    [`EventSource`{#the-eventsource-interface:eventsource-7}](#eventsource)
    object.

2.  Let `settings`{.variable} be `ev`{.variable}\'s [relevant settings
    object](webappapis.html#relevant-settings-object){#the-eventsource-interface:relevant-settings-object}.

3.  Let `urlRecord`{.variable} be the result of [encoding-parsing a
    URL](urls-and-fetching.html#encoding-parsing-a-url){#the-eventsource-interface:encoding-parsing-a-url}
    given `url`{.variable}, relative to `settings`{.variable}.

4.  If `urlRecord`{.variable} is failure, then throw a
    [\"`SyntaxError`\"](https://webidl.spec.whatwg.org/#syntaxerror){#the-eventsource-interface:syntaxerror
    x-internal="syntaxerror"}
    [`DOMException`{#the-eventsource-interface:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

5.  Set `ev`{.variable}\'s
    [url](#concept-eventsource-url){#the-eventsource-interface:concept-eventsource-url-4}
    to `urlRecord`{.variable}.

6.  Let `corsAttributeState`{.variable} be
    [Anonymous](urls-and-fetching.html#attr-crossorigin-anonymous){#the-eventsource-interface:attr-crossorigin-anonymous}.

7.  If the value of `eventSourceInitDict`{.variable}\'s
    [`withCredentials`{#the-eventsource-interface:dom-eventsourceinit-withcredentials-3}](#dom-eventsourceinit-withcredentials)
    member is true, then set `corsAttributeState`{.variable} to [Use
    Credentials](urls-and-fetching.html#attr-crossorigin-use-credentials){#the-eventsource-interface:attr-crossorigin-use-credentials}
    and set `ev`{.variable}\'s
    [`withCredentials`{#the-eventsource-interface:dom-eventsource-withcredentials-2}](#dom-eventsource-withcredentials)
    attribute to true.

8.  Let `request`{.variable} be the result of [creating a potential-CORS
    request](urls-and-fetching.html#create-a-potential-cors-request){#the-eventsource-interface:create-a-potential-cors-request}
    given `urlRecord`{.variable}, the empty string, and
    `corsAttributeState`{.variable}.

9.  Set `request`{.variable}\'s
    [client](https://fetch.spec.whatwg.org/#concept-request-client){#the-eventsource-interface:concept-request-client
    x-internal="concept-request-client"} to `settings`{.variable}.

10. User agents may
    [set](https://fetch.spec.whatwg.org/#concept-header-list-set){#the-eventsource-interface:concept-header-list-set
    x-internal="concept-header-list-set"}
    (\`[`Accept`{#the-eventsource-interface:http-accept}](https://httpwg.org/specs/rfc7231.html#header.accept){x-internal="http-accept"}\`,
    \`[`text/event-stream`{#the-eventsource-interface:text/event-stream}](iana.html#text/event-stream)\`)
    in `request`{.variable}\'s [header
    list](https://fetch.spec.whatwg.org/#concept-request-header-list){#the-eventsource-interface:concept-request-header-list
    x-internal="concept-request-header-list"}.

11. Set `request`{.variable}\'s [cache
    mode](https://fetch.spec.whatwg.org/#concept-request-cache-mode){#the-eventsource-interface:concept-request-cache-mode
    x-internal="concept-request-cache-mode"} to \"`no-store`\".

12. Set `request`{.variable}\'s [initiator
    type](https://fetch.spec.whatwg.org/#request-initiator-type){#the-eventsource-interface:concept-request-initiator-type
    x-internal="concept-request-initiator-type"} to \"`other`\".

13. Set `ev`{.variable}\'s
    [request](#concept-event-stream-request){#the-eventsource-interface:concept-event-stream-request}
    to `request`{.variable}.

14. Let `processEventSourceEndOfBody`{.variable} given
    [response](https://fetch.spec.whatwg.org/#concept-response){#the-eventsource-interface:concept-response
    x-internal="concept-response"} `res`{.variable} be the following
    step: if `res`{.variable} is not a [network
    error](https://fetch.spec.whatwg.org/#concept-network-error){#the-eventsource-interface:network-error
    x-internal="network-error"}, then [reestablish the
    connection](#reestablish-the-connection){#the-eventsource-interface:reestablish-the-connection}.

15. [Fetch](https://fetch.spec.whatwg.org/#concept-fetch){#the-eventsource-interface:concept-fetch-2
    x-internal="concept-fetch"} `request`{.variable}, with
    *[processResponseEndOfBody](https://fetch.spec.whatwg.org/#fetch-processresponseendofbody){x-internal="processresponseendofbody"}*
    set to `processEventSourceEndOfBody`{.variable} and
    *[processResponse](https://fetch.spec.whatwg.org/#process-response){x-internal="processresponse"}*
    set to the following steps given
    [response](https://fetch.spec.whatwg.org/#concept-response){#the-eventsource-interface:concept-response-2
    x-internal="concept-response"} `res`{.variable}:

    1.  If `res`{.variable} is an [aborted network
        error](https://fetch.spec.whatwg.org/#concept-aborted-network-error){#the-eventsource-interface:aborted-network-error
        x-internal="aborted-network-error"}, then [fail the
        connection](#fail-the-connection){#the-eventsource-interface:fail-the-connection}.

    2.  Otherwise, if `res`{.variable} is a [network
        error](https://fetch.spec.whatwg.org/#concept-network-error){#the-eventsource-interface:network-error-2
        x-internal="network-error"}, then [reestablish the
        connection](#reestablish-the-connection){#the-eventsource-interface:reestablish-the-connection-2},
        unless the user agent knows that to be futile, in which case the
        user agent may [fail the
        connection](#fail-the-connection){#the-eventsource-interface:fail-the-connection-2}.

    3.  Otherwise, if `res`{.variable}\'s
        [status](https://fetch.spec.whatwg.org/#concept-response-status){#the-eventsource-interface:concept-response-status
        x-internal="concept-response-status"} is not 200, or if
        `res`{.variable}\'s
        \`[`Content-Type`{#the-eventsource-interface:content-type}](urls-and-fetching.html#content-type)\`
        is not
        \`[`text/event-stream`{#the-eventsource-interface:text/event-stream-2}](iana.html#text/event-stream)\`,
        then [fail the
        connection](#fail-the-connection){#the-eventsource-interface:fail-the-connection-3}.

    4.  Otherwise, [announce the
        connection](#announce-the-connection){#the-eventsource-interface:announce-the-connection}
        and [interpret](#event-stream-interpretation)
        `res`{.variable}\'s
        [body](https://fetch.spec.whatwg.org/#concept-response-body){#the-eventsource-interface:concept-response-body
        x-internal="concept-response-body"} line by line.

16. Return `ev`{.variable}.

------------------------------------------------------------------------

The [`url`]{#dom-eventsource-url .dfn dfn-for="EventSource"
dfn-type="attribute"} attribute\'s getter must return the
[serialization](https://url.spec.whatwg.org/#concept-url-serializer){#the-eventsource-interface:concept-url-serializer
x-internal="concept-url-serializer"} of this
[`EventSource`{#the-eventsource-interface:eventsource-8}](#eventsource)
object\'s
[url](#concept-eventsource-url){#the-eventsource-interface:concept-eventsource-url-5}.

The [`withCredentials`]{#dom-eventsource-withcredentials .dfn
dfn-for="EventSource" dfn-type="attribute"} attribute must return the
value to which it was last initialized. When the object is created, it
must be initialized to false.

The [`readyState`]{#dom-eventsource-readystate .dfn
dfn-for="EventSource" dfn-type="attribute"} attribute represents the
state of the connection. It can have the following values:

[`CONNECTING`]{#dom-eventsource-connecting .dfn dfn-for="EventSource" dfn-type="const"} (numeric value 0)
:   The connection has not yet been established, or it was closed and
    the user agent is reconnecting.

[`OPEN`]{#dom-eventsource-open .dfn dfn-for="EventSource" dfn-type="const"} (numeric value 1)
:   The user agent has an open connection and is dispatching events as
    it receives them.

[`CLOSED`]{#dom-eventsource-closed .dfn dfn-for="EventSource" dfn-type="const"} (numeric value 2)
:   The connection is not open, and the user agent is not trying to
    reconnect. Either there was a fatal error or the
    [`close()`{#the-eventsource-interface:dom-eventsource-close-2}](#dom-eventsource-close)
    method was invoked.

When the object is created its
[`readyState`{#the-eventsource-interface:dom-eventsource-readystate-3}](#dom-eventsource-readystate)
must be set to
[`CONNECTING`{#the-eventsource-interface:dom-eventsource-connecting-2}](#dom-eventsource-connecting)
(0). The rules given below for handling the connection define when the
value changes.

The [`close()`]{#dom-eventsource-close .dfn dfn-for="EventSource"
dfn-type="method"} method must abort any instances of the
[fetch](https://fetch.spec.whatwg.org/#concept-fetch){#the-eventsource-interface:concept-fetch-3
x-internal="concept-fetch"} algorithm started for this
[`EventSource`{#the-eventsource-interface:eventsource-9}](#eventsource)
object, and must set the
[`readyState`{#the-eventsource-interface:dom-eventsource-readystate-4}](#dom-eventsource-readystate)
attribute to
[`CLOSED`{#the-eventsource-interface:dom-eventsource-closed-3}](#dom-eventsource-closed).

The following are the [event
handlers](webappapis.html#event-handlers){#the-eventsource-interface:event-handlers}
(and their corresponding [event handler event
types](webappapis.html#event-handler-event-type){#the-eventsource-interface:event-handler-event-type})
that must be supported, as [event handler IDL
attributes](webappapis.html#event-handler-idl-attributes){#the-eventsource-interface:event-handler-idl-attributes},
by all objects implementing the
[`EventSource`{#the-eventsource-interface:eventsource-10}](#eventsource)
interface:

[Event
handler](webappapis.html#event-handlers){#the-eventsource-interface:event-handlers-2}

[Event handler event
type](webappapis.html#event-handler-event-type){#the-eventsource-interface:event-handler-event-type-2}

[`onopen`]{#handler-eventsource-onopen .dfn dfn-for="EventSource"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[EventSource/open_event](https://developer.mozilla.org/en-US/docs/Web/API/EventSource/open_event "The open event of the EventSource API is fired when a connection with an event source is opened.")

Support in all current engines.

::: support
[Firefox6+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome6+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android45+]{.firefox_android .yes}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12+]{.opera_android .yes}
:::
::::
:::::

[`open`{#the-eventsource-interface:event-open}](indices.html#event-open)

[`onmessage`]{#handler-eventsource-onmessage .dfn dfn-for="EventSource"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[EventSource/message_event](https://developer.mozilla.org/en-US/docs/Web/API/EventSource/message_event "The message event of the EventSource API is fired when data is received through an event source.")

Support in all current engines.

::: support
[Firefox6+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome6+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android45+]{.firefox_android .yes}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12+]{.opera_android .yes}
:::
::::
:::::

[`message`{#the-eventsource-interface:event-message}](indices.html#event-message)

[`onerror`]{#handler-eventsource-onerror .dfn dfn-for="EventSource"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[EventSource/error_event](https://developer.mozilla.org/en-US/docs/Web/API/EventSource/error_event "The error event of the EventSource API is fired when a connection with an event source fails to be opened.")

Support in all current engines.

::: support
[Firefox6+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome6+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android45+]{.firefox_android .yes}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12+]{.opera_android .yes}
:::
::::
:::::

[`error`{#the-eventsource-interface:event-error}](indices.html#event-error)

------------------------------------------------------------------------

#### [9.2.3]{.secno} []{#processing-model-9}Processing model[](#sse-processing-model){.self-link} {#sse-processing-model}

When a user agent is to [announce the
connection]{#announce-the-connection .dfn}, the user agent must [queue a
task](webappapis.html#queue-a-task){#sse-processing-model:queue-a-task}
which, if the
[`readyState`{#sse-processing-model:dom-eventsource-readystate}](#dom-eventsource-readystate)
attribute is set to a value other than
[`CLOSED`{#sse-processing-model:dom-eventsource-closed}](#dom-eventsource-closed),
sets the
[`readyState`{#sse-processing-model:dom-eventsource-readystate-2}](#dom-eventsource-readystate)
attribute to
[`OPEN`{#sse-processing-model:dom-eventsource-open}](#dom-eventsource-open)
and [fires an
event](https://dom.spec.whatwg.org/#concept-event-fire){#sse-processing-model:concept-event-fire
x-internal="concept-event-fire"} named
[`open`{#sse-processing-model:event-open}](indices.html#event-open) at
the [`EventSource`{#sse-processing-model:eventsource}](#eventsource)
object.

When a user agent is to [reestablish the
connection]{#reestablish-the-connection .dfn}, the user agent must run
the following steps. These steps are run [in
parallel](infrastructure.html#in-parallel){#sse-processing-model:in-parallel},
not as part of a
[task](webappapis.html#concept-task){#sse-processing-model:concept-task}.
(The tasks that it queues, of course, are run like normal tasks and not
themselves [in
parallel](infrastructure.html#in-parallel){#sse-processing-model:in-parallel-2}.)

1.  [Queue a
    task](webappapis.html#queue-a-task){#sse-processing-model:queue-a-task-2}
    to run the following steps:

    1.  If the
        [`readyState`{#sse-processing-model:dom-eventsource-readystate-3}](#dom-eventsource-readystate)
        attribute is set to
        [`CLOSED`{#sse-processing-model:dom-eventsource-closed-2}](#dom-eventsource-closed),
        abort the task.

    2.  Set the
        [`readyState`{#sse-processing-model:dom-eventsource-readystate-4}](#dom-eventsource-readystate)
        attribute to
        [`CONNECTING`{#sse-processing-model:dom-eventsource-connecting}](#dom-eventsource-connecting).

    3.  [Fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#sse-processing-model:concept-event-fire-2
        x-internal="concept-event-fire"} named
        [`error`{#sse-processing-model:event-error}](indices.html#event-error)
        at the
        [`EventSource`{#sse-processing-model:eventsource-2}](#eventsource)
        object.

2.  Wait a delay equal to the reconnection time of the event source.

3.  Optionally, wait some more. In particular, if the previous attempt
    failed, then user agents might introduce an exponential backoff
    delay to avoid overloading a potentially already overloaded server.
    Alternatively, if the operating system has reported that there is no
    network connectivity, user agents might wait for the operating
    system to announce that the network connection has returned before
    retrying.

4.  Wait until the aforementioned task has run, if it has not yet run.

5.  [Queue a
    task](webappapis.html#queue-a-task){#sse-processing-model:queue-a-task-3}
    to run the following steps:

    1.  If the
        [`EventSource`{#sse-processing-model:eventsource-3}](#eventsource)
        object\'s
        [`readyState`{#sse-processing-model:dom-eventsource-readystate-5}](#dom-eventsource-readystate)
        attribute is not set to
        [`CONNECTING`{#sse-processing-model:dom-eventsource-connecting-2}](#dom-eventsource-connecting),
        then return.

    2.  Let `request`{.variable} be the
        [`EventSource`{#sse-processing-model:eventsource-4}](#eventsource)
        object\'s
        [request](#concept-event-stream-request){#sse-processing-model:concept-event-stream-request}.

    3.  If the
        [`EventSource`{#sse-processing-model:eventsource-5}](#eventsource)
        object\'s [last event ID
        string](#concept-event-stream-last-event-id){#sse-processing-model:concept-event-stream-last-event-id}
        is not the empty string, then:

        1.  Let `lastEventIDValue`{.variable} be the
            [`EventSource`{#sse-processing-model:eventsource-6}](#eventsource)
            object\'s [last event ID
            string](#concept-event-stream-last-event-id){#sse-processing-model:concept-event-stream-last-event-id-2},
            [encoded as
            UTF-8](https://encoding.spec.whatwg.org/#utf-8-encode){#sse-processing-model:utf-8-encode
            x-internal="utf-8-encode"}.

        2.  [Set](https://fetch.spec.whatwg.org/#concept-header-list-set){#sse-processing-model:concept-header-list-set
            x-internal="concept-header-list-set"}
            (\`[`Last-Event-ID`{#sse-processing-model:last-event-id}](#last-event-id)\`,
            `lastEventIDValue`{.variable}) in `request`{.variable}\'s
            [header
            list](https://fetch.spec.whatwg.org/#concept-request-header-list){#sse-processing-model:concept-request-header-list
            x-internal="concept-request-header-list"}.

    4.  [Fetch](https://fetch.spec.whatwg.org/#concept-fetch){#sse-processing-model:concept-fetch
        x-internal="concept-fetch"} `request`{.variable} and process the
        response obtained in this fashion, if any, as described earlier
        in this section.

When a user agent is to [fail the connection]{#fail-the-connection
.dfn}, the user agent must [queue a
task](webappapis.html#queue-a-task){#sse-processing-model:queue-a-task-4}
which, if the
[`readyState`{#sse-processing-model:dom-eventsource-readystate-6}](#dom-eventsource-readystate)
attribute is set to a value other than
[`CLOSED`{#sse-processing-model:dom-eventsource-closed-3}](#dom-eventsource-closed),
sets the
[`readyState`{#sse-processing-model:dom-eventsource-readystate-7}](#dom-eventsource-readystate)
attribute to
[`CLOSED`{#sse-processing-model:dom-eventsource-closed-4}](#dom-eventsource-closed)
and [fires an
event](https://dom.spec.whatwg.org/#concept-event-fire){#sse-processing-model:concept-event-fire-3
x-internal="concept-event-fire"} named
[`error`{#sse-processing-model:event-error-2}](indices.html#event-error)
at the
[`EventSource`{#sse-processing-model:eventsource-7}](#eventsource)
object. **Once the user agent has [failed the
connection](#fail-the-connection){#sse-processing-model:fail-the-connection},
it does *not* attempt to reconnect.**

------------------------------------------------------------------------

The [task
source](webappapis.html#task-source){#sse-processing-model:task-source}
for any
[tasks](webappapis.html#concept-task){#sse-processing-model:concept-task-2}
that are
[queued](webappapis.html#queue-a-task){#sse-processing-model:queue-a-task-5}
by [`EventSource`{#sse-processing-model:eventsource-8}](#eventsource)
objects is the [remote event task source]{#remote-event-task-source
.dfn}.

#### [9.2.4]{.secno} The \`[`Last-Event-ID`{#the-last-event-id-header:last-event-id}](#last-event-id)\` header[](#the-last-event-id-header){.self-link}

The [`Last-Event-ID`]{#last-event-id .dfn dfn-type="http-header"}\` HTTP
request header reports an
[`EventSource`{#the-last-event-id-header:eventsource}](#eventsource)
object\'s [last event ID
string](#concept-event-stream-last-event-id){#the-last-event-id-header:concept-event-stream-last-event-id}
to the server when the user agent is to [reestablish the
connection](#reestablish-the-connection){#the-last-event-id-header:reestablish-the-connection}.

See [whatwg/html issue
#7363](https://github.com/whatwg/html/issues/7363) to define the value
space better. It is essentially any UTF-8 encoded string, that does not
contain U+0000 NULL, U+000A LF, or U+000D CR.

#### [9.2.5]{.secno} Parsing an event stream[](#parsing-an-event-stream){.self-link}

This event stream format\'s [MIME
type](https://mimesniff.spec.whatwg.org/#mime-type){#parsing-an-event-stream:mime-type
x-internal="mime-type"} is
[`text/event-stream`{#parsing-an-event-stream:text/event-stream}](iana.html#text/event-stream).

The event stream format is as described by the `stream` production of
the following ABNF, the character set for which is Unicode.
[\[ABNF\]](references.html#refsABNF)

``` abnf
stream        = [ bom ] *event
event         = *( comment / field ) end-of-line
comment       = colon *any-char end-of-line
field         = 1*name-char [ colon [ space ] *any-char ] end-of-line
end-of-line   = ( cr lf / cr / lf )

; characters
lf            = %x000A ; U+000A LINE FEED (LF)
cr            = %x000D ; U+000D CARRIAGE RETURN (CR)
space         = %x0020 ; U+0020 SPACE
colon         = %x003A ; U+003A COLON (:)
bom           = %xFEFF ; U+FEFF BYTE ORDER MARK
name-char     = %x0000-0009 / %x000B-000C / %x000E-0039 / %x003B-10FFFF
                ; a scalar value other than U+000A LINE FEED (LF), U+000D CARRIAGE RETURN (CR), or U+003A COLON (:)
any-char      = %x0000-0009 / %x000B-000C / %x000E-10FFFF
                ; a scalar value other than U+000A LINE FEED (LF) or U+000D CARRIAGE RETURN (CR)
```

Event streams in this format must always be encoded as UTF-8.
[\[ENCODING\]](references.html#refsENCODING)

Lines must be separated by either a U+000D CARRIAGE RETURN U+000A LINE
FEED (CRLF) character pair, a single U+000A LINE FEED (LF) character, or
a single U+000D CARRIAGE RETURN (CR) character.

Since connections established to remote servers for such resources are
expected to be long-lived, UAs should ensure that appropriate buffering
is used. In particular, while line buffering with lines are defined to
end with a single U+000A LINE FEED (LF) character is safe, block
buffering or line buffering with different expected line endings can
cause delays in event dispatch.

#### [9.2.6]{.secno} Interpreting an event stream[](#event-stream-interpretation){.self-link} {#event-stream-interpretation}

Streams must be decoded using the [UTF-8
decode](https://encoding.spec.whatwg.org/#utf-8-decode){#event-stream-interpretation:utf-8-decode
x-internal="utf-8-decode"} algorithm.

The [UTF-8
decode](https://encoding.spec.whatwg.org/#utf-8-decode){#event-stream-interpretation:utf-8-decode-2
x-internal="utf-8-decode"} algorithm strips one leading UTF-8 Byte Order
Mark (BOM), if any.

The stream must then be parsed by reading everything line by line, with
a U+000D CARRIAGE RETURN U+000A LINE FEED (CRLF) character pair, a
single U+000A LINE FEED (LF) character not preceded by a U+000D CARRIAGE
RETURN (CR) character, and a single U+000D CARRIAGE RETURN (CR)
character not followed by a U+000A LINE FEED (LF) character being the
ways in which a line can end.

When a stream is parsed, a `data`{.variable} buffer, an
`event type`{.variable} buffer, and a `last event ID`{.variable} buffer
must be associated with it. They must be initialized to the empty
string.

Lines must be processed, in the order they are received, as follows:

If the line is empty (a blank line)
:   [Dispatch the
    event](#dispatchMessage){#event-stream-interpretation:dispatchMessage},
    as defined below.

If the line starts with a U+003A COLON character (:)
:   Ignore the line.

If the line contains a U+003A COLON character (:)

:   Collect the characters on the line before the first U+003A COLON
    character (:), and let `field`{.variable} be that string.

    Collect the characters on the line after the first U+003A COLON
    character (:), and let `value`{.variable} be that string. If
    `value`{.variable} starts with a U+0020 SPACE character, remove it
    from `value`{.variable}.

    [Process the field](#processField) using the steps described below,
    using `field`{.variable} as the field name and `value`{.variable} as
    the field value.

Otherwise, the string is not empty but does not contain a U+003A COLON character (:)

:   [Process the field](#processField) using the steps described below,
    using the whole line as the field name, and the empty string as the
    field value.

Once the end of the file is reached, any pending data must be discarded.
(If the file ends in the middle of an event, before the final empty
line, the incomplete event is not dispatched.)

------------------------------------------------------------------------

The steps to [process the field]{.dfn} given a field name and a field
value depend on the field name, as given in the following list. Field
names must be compared literally, with no case folding performed.

If the field name is \"event\"
:   Set the `event type`{.variable} buffer to the field value.

If the field name is \"data\"
:   Append the field value to the `data`{.variable} buffer, then append
    a single U+000A LINE FEED (LF) character to the `data`{.variable}
    buffer.

If the field name is \"id\"
:   If the field value does not contain U+0000 NULL, then set the
    `last event ID`{.variable} buffer to the field value. Otherwise,
    ignore the field.

If the field name is \"retry\"
:   If the field value consists of only [ASCII
    digits](https://infra.spec.whatwg.org/#ascii-digit){#event-stream-interpretation:ascii-digits
    x-internal="ascii-digits"}, then interpret the field value as an
    integer in base ten, and set the event stream\'s [reconnection
    time](#concept-event-stream-reconnection-time){#event-stream-interpretation:concept-event-stream-reconnection-time}
    to that integer. Otherwise, ignore the field.

Otherwise

:   The field is ignored.

When the user agent is required to [dispatch the event]{#dispatchMessage
.dfn}, the user agent must process the `data`{.variable} buffer, the
`event type`{.variable} buffer, and the `last event ID`{.variable}
buffer using steps appropriate for the user agent.

For web browsers, the appropriate steps to [dispatch the
event](#dispatchMessage){#event-stream-interpretation:dispatchMessage-2}
are as follows:

1.  Set the [last event ID
    string](#concept-event-stream-last-event-id){#event-stream-interpretation:concept-event-stream-last-event-id}
    of the event source to the value of the `last event ID`{.variable}
    buffer. The buffer does not get reset, so the [last event ID
    string](#concept-event-stream-last-event-id){#event-stream-interpretation:concept-event-stream-last-event-id-2}
    of the event source remains set to this value until the next time it
    is set by the server.

2.  If the `data`{.variable} buffer is an empty string, set the
    `data`{.variable} buffer and the `event type`{.variable} buffer to
    the empty string and return.

3.  If the `data`{.variable} buffer\'s last character is a U+000A LINE
    FEED (LF) character, then remove the last character from the
    `data`{.variable} buffer.

4.  Let `event`{.variable} be the result of [creating an
    event](https://dom.spec.whatwg.org/#concept-event-create){#event-stream-interpretation:creating-an-event
    x-internal="creating-an-event"} using
    [`MessageEvent`{#event-stream-interpretation:messageevent}](comms.html#messageevent),
    in the [relevant
    realm](webappapis.html#concept-relevant-realm){#event-stream-interpretation:concept-relevant-realm}
    of the
    [`EventSource`{#event-stream-interpretation:eventsource}](#eventsource)
    object.

5.  Initialize `event`{.variable}\'s
    [`type`{#event-stream-interpretation:dom-event-type}](https://dom.spec.whatwg.org/#dom-event-type){x-internal="dom-event-type"}
    attribute to
    \"[`message`{#event-stream-interpretation:event-message}](indices.html#event-message)\",
    its
    [`data`{#event-stream-interpretation:dom-messageevent-data}](comms.html#dom-messageevent-data)
    attribute to `data`{.variable}, its
    [`origin`{#event-stream-interpretation:dom-messageevent-origin}](comms.html#dom-messageevent-origin)
    attribute to the
    [serialization](browsers.html#ascii-serialisation-of-an-origin){#event-stream-interpretation:ascii-serialisation-of-an-origin}
    of the
    [origin](https://url.spec.whatwg.org/#concept-url-origin){#event-stream-interpretation:concept-url-origin
    x-internal="concept-url-origin"} of the event stream\'s final URL
    (i.e., the URL after redirects), and its
    [`lastEventId`{#event-stream-interpretation:dom-messageevent-lasteventid}](comms.html#dom-messageevent-lasteventid)
    attribute to the [last event ID
    string](#concept-event-stream-last-event-id){#event-stream-interpretation:concept-event-stream-last-event-id-3}
    of the event source.

6.  If the `event type`{.variable} buffer has a value other than the
    empty string, change the
    [type](https://dom.spec.whatwg.org/#dom-event-type){#event-stream-interpretation:concept-event-type
    x-internal="concept-event-type"} of the newly created event to equal
    the value of the `event type`{.variable} buffer.

7.  Set the `data`{.variable} buffer and the `event type`{.variable}
    buffer to the empty string.

8.  [Queue a
    task](webappapis.html#queue-a-task){#event-stream-interpretation:queue-a-task}
    which, if the
    [`readyState`{#event-stream-interpretation:dom-eventsource-readystate}](#dom-eventsource-readystate)
    attribute is set to a value other than
    [`CLOSED`{#event-stream-interpretation:dom-eventsource-closed}](#dom-eventsource-closed),
    [dispatches](https://dom.spec.whatwg.org/#concept-event-dispatch){#event-stream-interpretation:concept-event-dispatch
    x-internal="concept-event-dispatch"} the newly created event at the
    [`EventSource`{#event-stream-interpretation:eventsource-2}](#eventsource)
    object.

If an event doesn\'t have an \"id\" field, but an earlier event did set
the event source\'s [last event ID
string](#concept-event-stream-last-event-id){#event-stream-interpretation:concept-event-stream-last-event-id-4},
then the event\'s
[`lastEventId`{#event-stream-interpretation:dom-messageevent-lasteventid-2}](comms.html#dom-messageevent-lasteventid)
field will be set to the value of whatever the last seen \"id\" field
was.

For other user agents, the appropriate steps to [dispatch the
event](#dispatchMessage){#event-stream-interpretation:dispatchMessage-3}
are implementation dependent, but at a minimum they must set the
`data`{.variable} and `event type`{.variable} buffers to the empty
string before returning.

::: example
The following event stream, once followed by a blank line:

    data: YHOO
    data: +2
    data: 10

\...would cause an event
[`message`{#event-stream-interpretation:event-message-2}](indices.html#event-message)
with the interface
[`MessageEvent`{#event-stream-interpretation:messageevent-2}](comms.html#messageevent)
to be dispatched on the
[`EventSource`{#event-stream-interpretation:eventsource-3}](#eventsource)
object. The event\'s
[`data`{#event-stream-interpretation:dom-messageevent-data-2}](comms.html#dom-messageevent-data)
attribute would contain the string \"`YHOO\n+2\n10`\" (where \"`\n`\"
represents a newline).

This could be used as follows:

``` js
var stocks = new EventSource("https://stocks.example.com/ticker.php");
stocks.onmessage = function (event) {
  var data = event.data.split('\n');
  updateStocks(data[0], data[1], data[2]);
};
```

\...where `updateStocks()` is a function defined as:

``` js
function updateStocks(symbol, delta, value) { ... }
```

\...or some such.
:::

::: example
The following stream contains four blocks. The first block has just a
comment, and will fire nothing. The second block has two fields with
names \"data\" and \"id\" respectively; an event will be fired for this
block, with the data \"first event\", and will then set the last event
ID to \"1\" so that if the connection died between this block and the
next, the server would be sent a
\`[`Last-Event-ID`{#event-stream-interpretation:last-event-id}](#last-event-id)\`
header with the value \``1`\`. The third block fires an event with data
\"second event\", and also has an \"id\" field, this time with no value,
which resets the last event ID to the empty string (meaning no
\`[`Last-Event-ID`{#event-stream-interpretation:last-event-id-2}](#last-event-id)\`
header will now be sent in the event of a reconnection being attempted).
Finally, the last block just fires an event with the data
\" third event\" (with a single leading space character). Note that the
last still has to end with a blank line, the end of the stream is not
enough to trigger the dispatch of the last event.

    : test stream

    data: first event
    id: 1

    data:second event
    id

    data:  third event
:::

::: example
The following stream fires two events:

    data

    data
    data

    data:

The first block fires events with the data set to the empty string, as
would the last block if it was followed by a blank line. The middle
block fires an event with the data set to a single newline character.
The last block is discarded because it is not followed by a blank line.
:::

::: example
The following stream fires two identical events:

    data:test

    data: test

This is because the space after the colon is ignored if present.
:::

#### [9.2.7]{.secno} Authoring notes[](#authoring-notes){.self-link}

Legacy proxy servers are known to, in certain cases, drop HTTP
connections after a short timeout. To protect against such proxy
servers, authors can include a comment line (one starting with a \':\'
character) every 15 seconds or so.

Authors wishing to relate event source connections to each other or to
specific documents previously served might find that relying on IP
addresses doesn\'t work, as individual clients can have multiple IP
addresses (due to having multiple proxy servers) and individual IP
addresses can have multiple clients (due to sharing a proxy server). It
is better to include a unique identifier in the document when it is
served and then pass that identifier as part of the URL when the
connection is established.

Authors are also cautioned that HTTP chunking can have unexpected
negative effects on the reliability of this protocol, in particular if
the chunking is done by a different layer unaware of the timing
requirements. If this is a problem, chunking can be disabled for serving
event streams.

Clients that support HTTP\'s per-server connection limitation might run
into trouble when opening multiple pages from a site if each page has an
[`EventSource`{#authoring-notes:eventsource}](#eventsource) to the same
domain. Authors can avoid this using the relatively complex mechanism of
using unique domain names per connection, or by allowing the user to
enable or disable the
[`EventSource`{#authoring-notes:eventsource-2}](#eventsource)
functionality on a per-page basis, or by sharing a single
[`EventSource`{#authoring-notes:eventsource-3}](#eventsource) object
using a [shared
worker](workers.html#sharedworkerglobalscope){#authoring-notes:sharedworkerglobalscope}.

#### [9.2.8]{.secno} Connectionless push and other features[](#eventsource-push){.self-link} {#eventsource-push}

User agents running in controlled environments, e.g. browsers on mobile
handsets tied to specific carriers, may offload the management of the
connection to a proxy on the network. In such a situation, the user
agent for the purposes of conformance is considered to include both the
handset software and the network proxy.

::: example
For example, a browser on a mobile device, after having established a
connection, might detect that it is on a supporting network and request
that a proxy server on the network take over the management of the
connection. The timeline for such a situation might be as follows:

1.  Browser connects to a remote HTTP server and requests the resource
    specified by the author in the
    [`EventSource`{#eventsource-push:dom-eventsource}](#dom-eventsource)
    constructor.
2.  The server sends occasional messages.
3.  In between two messages, the browser detects that it is idle except
    for the network activity involved in keeping the TCP connection
    alive, and decides to switch to sleep mode to save power.
4.  The browser disconnects from the server.
5.  The browser contacts a service on the network, and requests that the
    service, a \"push proxy\", maintain the connection instead.
6.  The \"push proxy\" service contacts the remote HTTP server and
    requests the resource specified by the author in the
    [`EventSource`{#eventsource-push:dom-eventsource-2}](#dom-eventsource)
    constructor (possibly including a
    \`[`Last-Event-ID`{#eventsource-push:last-event-id}](#last-event-id)\`
    HTTP header, etc.).
7.  The browser allows the mobile device to go to sleep.
8.  The server sends another message.
9.  The \"push proxy\" service uses a technology such as OMA push to
    convey the event to the mobile device, which wakes only enough to
    process the event and then returns to sleep.
:::

This can reduce the total data usage, and can therefore result in
considerable power savings.

As well as implementing the existing API and
[`text/event-stream`{#eventsource-push:text/event-stream}](iana.html#text/event-stream)
wire format as defined by this specification and in more distributed
ways as described above, formats of event framing defined by [other
applicable
specifications](infrastructure.html#other-applicable-specifications){#eventsource-push:other-applicable-specifications}
may be supported. This specification does not define how they are to be
parsed or processed.

#### [9.2.9]{.secno} Garbage collection[](#garbage-collection){.self-link}

While an [`EventSource`{#garbage-collection:eventsource}](#eventsource)
object\'s
[`readyState`{#garbage-collection:dom-eventsource-readystate}](#dom-eventsource-readystate)
is
[`CONNECTING`{#garbage-collection:dom-eventsource-connecting}](#dom-eventsource-connecting),
and the object has one or more event listeners registered for
[`open`{#garbage-collection:event-open}](indices.html#event-open),
[`message`{#garbage-collection:event-message}](indices.html#event-message),
or [`error`{#garbage-collection:event-error}](indices.html#event-error)
events, there must be a strong reference from the
[`Window`{#garbage-collection:window}](nav-history-apis.html#window) or
[`WorkerGlobalScope`{#garbage-collection:workerglobalscope}](workers.html#workerglobalscope)
object that the
[`EventSource`{#garbage-collection:eventsource-2}](#eventsource)
object\'s constructor was invoked from to the
[`EventSource`{#garbage-collection:eventsource-3}](#eventsource) object
itself.

While an
[`EventSource`{#garbage-collection:eventsource-4}](#eventsource)
object\'s
[`readyState`{#garbage-collection:dom-eventsource-readystate-2}](#dom-eventsource-readystate)
is
[`OPEN`{#garbage-collection:dom-eventsource-open}](#dom-eventsource-open),
and the object has one or more event listeners registered for
[`message`{#garbage-collection:event-message-2}](indices.html#event-message)
or
[`error`{#garbage-collection:event-error-2}](indices.html#event-error)
events, there must be a strong reference from the
[`Window`{#garbage-collection:window-2}](nav-history-apis.html#window)
or
[`WorkerGlobalScope`{#garbage-collection:workerglobalscope-2}](workers.html#workerglobalscope)
object that the
[`EventSource`{#garbage-collection:eventsource-5}](#eventsource)
object\'s constructor was invoked from to the
[`EventSource`{#garbage-collection:eventsource-6}](#eventsource) object
itself.

While there is a task queued by an
[`EventSource`{#garbage-collection:eventsource-7}](#eventsource) object
on the [remote event task
source](#remote-event-task-source){#garbage-collection:remote-event-task-source},
there must be a strong reference from the
[`Window`{#garbage-collection:window-3}](nav-history-apis.html#window)
or
[`WorkerGlobalScope`{#garbage-collection:workerglobalscope-3}](workers.html#workerglobalscope)
object that the
[`EventSource`{#garbage-collection:eventsource-8}](#eventsource)
object\'s constructor was invoked from to that
[`EventSource`{#garbage-collection:eventsource-9}](#eventsource) object.

If a user agent is to [forcibly
close]{#concept-eventsource-forcibly-close .dfn} an
[`EventSource`{#garbage-collection:eventsource-10}](#eventsource) object
(this happens when a
[`Document`{#garbage-collection:document}](dom.html#document) object
goes away permanently), the user agent must abort any instances of the
[fetch](https://fetch.spec.whatwg.org/#concept-fetch){#garbage-collection:concept-fetch
x-internal="concept-fetch"} algorithm started for this
[`EventSource`{#garbage-collection:eventsource-11}](#eventsource)
object, and must set the
[`readyState`{#garbage-collection:dom-eventsource-readystate-3}](#dom-eventsource-readystate)
attribute to
[`CLOSED`{#garbage-collection:dom-eventsource-closed}](#dom-eventsource-closed).

If an [`EventSource`{#garbage-collection:eventsource-12}](#eventsource)
object is garbage collected while its connection is still open, the user
agent must abort any instance of the
[fetch](https://fetch.spec.whatwg.org/#concept-fetch){#garbage-collection:concept-fetch-2
x-internal="concept-fetch"} algorithm opened by this
[`EventSource`{#garbage-collection:eventsource-13}](#eventsource).

#### [9.2.10]{.secno} Implementation advice[](#implementation-advice){.self-link}

*This section is non-normative.*

User agents are strongly urged to provide detailed diagnostic
information about
[`EventSource`{#implementation-advice:eventsource}](#eventsource)
objects and their related network connections in their development
consoles, to aid authors in debugging code using this API.

For example, a user agent could have a panel displaying all the
[`EventSource`{#implementation-advice:eventsource-2}](#eventsource)
objects a page has created, each listing the constructor\'s arguments,
whether there was a network error, what the CORS status of the
connection is and what headers were sent by the client and received from
the server to lead to that status, the messages that were received and
how they were parsed, and so forth.

Implementations are especially encouraged to report detailed information
to their development consoles whenever an
[`error`{#implementation-advice:event-error}](indices.html#event-error)
event is fired, since little to no information can be made available in
the events themselves.

[← 9 Communication](comms.html) --- [Table of Contents](./) --- [9.3
Cross-document messaging →](web-messaging.html)
