## [9]{.secno} Communication[](#comms){.self-link} {#comms}

[]{#network-intro}[]{#the-websocket-interface}[]{#feedback-from-the-protocol}[]{#ping-and-pong-frames}[]{#the-closeevent-interface}[]{#garbage-collection-2}[]{#websocket}[]{#binarytype}[]{#closeevent}[]{#closeeventinit}The
[`WebSocket`{#comms:websocket-2}](https://websockets.spec.whatwg.org/#websocket){x-internal="websocket-2"}
interface used to be defined here. It is now defined in WebSockets.
[\[WEBSOCKETS\]](references.html#refsWEBSOCKETS)

### [9.1]{.secno} The [`MessageEvent`{#the-messageevent-interface:messageevent}](#messageevent) interface[](#the-messageevent-interface){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[MessageEvent](https://developer.mozilla.org/en-US/docs/Web/API/MessageEvent "The MessageEvent interface represents a message received by a target object.")

Support in all current engines.

::: support
[Firefox3+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome2+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

Messages in [server-sent
events](server-sent-events.html#server-sent-events){#the-messageevent-interface:server-sent-events},
[cross-document
messaging](web-messaging.html#web-messaging){#the-messageevent-interface:web-messaging},
[channel
messaging](web-messaging.html#channel-messaging){#the-messageevent-interface:channel-messaging},
[broadcast
channels](web-messaging.html#broadcasting-to-other-browsing-contexts){#the-messageevent-interface:broadcasting-to-other-browsing-contexts},
and WebSockets use the
[`MessageEvent`{#the-messageevent-interface:messageevent-2}](#messageevent)
interface for their
[`message`{#the-messageevent-interface:event-message}](indices.html#event-message)
events: [\[WEBSOCKETS\]](references.html#refsWEBSOCKETS)

``` idl
[Exposed=(Window,Worker,AudioWorklet)]
interface MessageEvent : Event {
  constructor(DOMString type, optional MessageEventInit eventInitDict = {});

  readonly attribute any data;
  readonly attribute USVString origin;
  readonly attribute DOMString lastEventId;
  readonly attribute MessageEventSource? source;
  readonly attribute FrozenArray<MessagePort> ports;

  undefined initMessageEvent(DOMString type, optional boolean bubbles = false, optional boolean cancelable = false, optional any data = null, optional USVString origin = "", optional DOMString lastEventId = "", optional MessageEventSource? source = null, optional sequence<MessagePort> ports = []);
};

dictionary MessageEventInit : EventInit {
  any data = null;
  USVString origin = "";
  DOMString lastEventId = "";
  MessageEventSource? source = null;
  sequence<MessagePort> ports = [];
};

typedef (WindowProxy or MessagePort or ServiceWorker) MessageEventSource;
```

`event`{.variable}`.`[`data`](#dom-messageevent-data){#dom-messageevent-data-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[MessageEvent/data](https://developer.mozilla.org/en-US/docs/Web/API/MessageEvent/data "The data read-only property of the MessageEvent interface represents the data sent by the message emitter.")

Support in all current engines.

::: support
[Firefox3+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome2+]{.chrome
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

Returns the data of the message.

`event`{.variable}`.`[`origin`](#dom-messageevent-origin){#dom-messageevent-origin-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[MessageEvent/origin](https://developer.mozilla.org/en-US/docs/Web/API/MessageEvent/origin "The origin read-only property of the MessageEvent interface is a string representing the origin of the message emitter.")

Support in all current engines.

::: support
[Firefox3+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome2+]{.chrome
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

Returns the origin of the message, for [server-sent
events](server-sent-events.html#server-sent-events){#the-messageevent-interface:server-sent-events-2}
and [cross-document
messaging](web-messaging.html#web-messaging){#the-messageevent-interface:web-messaging-2}.

`event`{.variable}`.`[`lastEventId`](#dom-messageevent-lasteventid){#dom-messageevent-lasteventid-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[MessageEvent/lastEventId](https://developer.mozilla.org/en-US/docs/Web/API/MessageEvent/lastEventId "The lastEventId read-only property of the MessageEvent interface is a string representing a unique ID for the event.")

Support in all current engines.

::: support
[Firefox3+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome2+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)17+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns the [last event ID
string](server-sent-events.html#concept-event-stream-last-event-id){#the-messageevent-interface:concept-event-stream-last-event-id},
for [server-sent
events](server-sent-events.html#server-sent-events){#the-messageevent-interface:server-sent-events-3}.

`event`{.variable}`.`[`source`](#dom-messageevent-source){#dom-messageevent-source-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[MessageEvent/source](https://developer.mozilla.org/en-US/docs/Web/API/MessageEvent/source "The source read-only property of the MessageEvent interface is a MessageEventSource (which can be a WindowProxy, MessagePort, or ServiceWorker object) representing the message emitter.")

Support in all current engines.

::: support
[Firefox3+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome2+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns the
[`WindowProxy`{#the-messageevent-interface:windowproxy-2}](nav-history-apis.html#windowproxy)
of the source window, for [cross-document
messaging](web-messaging.html#web-messaging){#the-messageevent-interface:web-messaging-3},
and the
[`MessagePort`{#the-messageevent-interface:messageport-5}](web-messaging.html#messageport)
being attached, in the
[`connect`{#the-messageevent-interface:event-workerglobalscope-connect}](indices.html#event-workerglobalscope-connect)
event fired at
[`SharedWorkerGlobalScope`{#the-messageevent-interface:sharedworkerglobalscope}](workers.html#sharedworkerglobalscope)
objects.

`event`{.variable}`.`[`ports`](#dom-messageevent-ports){#dom-messageevent-ports-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[MessageEvent/ports](https://developer.mozilla.org/en-US/docs/Web/API/MessageEvent/ports "The ports read-only property of the MessageEvent interface is an array of MessagePort objects representing the ports associated with the channel the message is being sent through (where appropriate, e.g. in channel messaging or when sending a message to a shared worker).")

Support in all current engines.

::: support
[Firefox3+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome4+]{.chrome
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

Returns the
[`MessagePort`{#the-messageevent-interface:messageport-6}](web-messaging.html#messageport)
array sent with the message, for [cross-document
messaging](web-messaging.html#web-messaging){#the-messageevent-interface:web-messaging-4}
and [channel
messaging](web-messaging.html#channel-messaging){#the-messageevent-interface:channel-messaging-2}.

The [`data`]{#dom-messageevent-data .dfn dfn-for="MessageEvent"
dfn-type="attribute"} attribute must return the value it was initialized
to. It represents the message being sent.

The [`origin`]{#dom-messageevent-origin .dfn dfn-for="MessageEvent"
dfn-type="attribute"} attribute must return the value it was initialized
to. It represents, in [server-sent
events](server-sent-events.html#server-sent-events){#the-messageevent-interface:server-sent-events-4}
and [cross-document
messaging](web-messaging.html#web-messaging){#the-messageevent-interface:web-messaging-5},
the
[origin](https://dom.spec.whatwg.org/#concept-document-origin){#the-messageevent-interface:concept-document-origin
x-internal="concept-document-origin"} of the document that sent the
message (typically the scheme, hostname, and port of the document, but
not its path or
[fragment](https://url.spec.whatwg.org/#concept-url-fragment){#the-messageevent-interface:concept-url-fragment
x-internal="concept-url-fragment"}).

The [`lastEventId`]{#dom-messageevent-lasteventid .dfn
dfn-for="MessageEvent" dfn-type="attribute"} attribute must return the
value it was initialized to. It represents, in [server-sent
events](server-sent-events.html#server-sent-events){#the-messageevent-interface:server-sent-events-5},
the [last event ID
string](server-sent-events.html#concept-event-stream-last-event-id){#the-messageevent-interface:concept-event-stream-last-event-id-2}
of the event source.

The [`source`]{#dom-messageevent-source .dfn dfn-for="MessageEvent"
dfn-type="attribute"} attribute must return the value it was initialized
to. It represents, in [cross-document
messaging](web-messaging.html#web-messaging){#the-messageevent-interface:web-messaging-6},
the
[`WindowProxy`{#the-messageevent-interface:windowproxy-3}](nav-history-apis.html#windowproxy)
of the [browsing
context](document-sequences.html#browsing-context){#the-messageevent-interface:browsing-context}
of the
[`Window`{#the-messageevent-interface:window}](nav-history-apis.html#window)
object from which the message came; and in the
[`connect`{#the-messageevent-interface:event-workerglobalscope-connect-2}](indices.html#event-workerglobalscope-connect)
events used by [shared
workers](workers.html#sharedworkerglobalscope){#the-messageevent-interface:sharedworkerglobalscope-2},
the newly connecting
[`MessagePort`{#the-messageevent-interface:messageport-7}](web-messaging.html#messageport).

The [`ports`]{#dom-messageevent-ports .dfn dfn-for="MessageEvent"
dfn-type="attribute"} attribute must return the value it was initialized
to. It represents, in [cross-document
messaging](web-messaging.html#web-messaging){#the-messageevent-interface:web-messaging-7}
and [channel
messaging](web-messaging.html#channel-messaging){#the-messageevent-interface:channel-messaging-3},
the
[`MessagePort`{#the-messageevent-interface:messageport-8}](web-messaging.html#messageport)
array being sent.

The
[`initMessageEvent(``type`{.variable}`, ``bubbles`{.variable}`, ``cancelable`{.variable}`, ``data`{.variable}`, ``origin`{.variable}`, ``lastEventId`{.variable}`, ``source`{.variable}`, ``ports`{.variable}`)`]{#dom-messageevent-initmessageevent
.dfn dfn-for="MessageEvent" dfn-type="method"} method must initialize
the event in a manner analogous to the similarly-named
[`initEvent()`{#the-messageevent-interface:dom-event-initevent}](https://dom.spec.whatwg.org/#dom-event-initevent){x-internal="dom-event-initevent"}
method. [\[DOM\]](references.html#refsDOM)

Various APIs (e.g.,
[`WebSocket`{#the-messageevent-interface:websocket-2}](https://websockets.spec.whatwg.org/#websocket){x-internal="websocket-2"},
[`EventSource`{#the-messageevent-interface:eventsource}](server-sent-events.html#eventsource))
use the
[`MessageEvent`{#the-messageevent-interface:messageevent-3}](#messageevent)
interface for their
[`message`{#the-messageevent-interface:event-message-2}](indices.html#event-message)
event without using the
[`MessagePort`{#the-messageevent-interface:messageport-9}](web-messaging.html#messageport)
API.

[← 8.10 Images](imagebitmap-and-animations.html) --- [Table of
Contents](./) --- [9.2 Server-sent events →](server-sent-events.html)
