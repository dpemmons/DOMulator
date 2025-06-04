HTML Standard

::: {#head .head .with-buttons}
[![WHATWG](https://resources.whatwg.org/logo.svg){.darkmode-aware
width="100" crossorigin="" height="100"}](https://whatwg.org/){.logo}

# HTML {#html .allcaps}

Living Standard --- Last Updated [2 June 2025]{.pubdate}
:::

[← 9.2 Server-sent events](server-sent-events.html) --- [Table of
Contents](./) --- [10 Web workers →](workers.html)

1.  ::: {#toc-comms}
    1.  [[9.3]{.secno} Cross-document
        messaging](web-messaging.html#web-messaging)
        1.  [[9.3.1]{.secno}
            Introduction](web-messaging.html#introduction-12)
        2.  [[9.3.2]{.secno}
            Security](web-messaging.html#security-postmsg)
            1.  [[9.3.2.1]{.secno} Authors](web-messaging.html#authors)
            2.  [[9.3.2.2]{.secno} User
                agents](web-messaging.html#user-agents)
        3.  [[9.3.3]{.secno} Posting
            messages](web-messaging.html#posting-messages)
    2.  [[9.4]{.secno} Channel
        messaging](web-messaging.html#channel-messaging)
        1.  [[9.4.1]{.secno}
            Introduction](web-messaging.html#introduction-13)
            1.  [[9.4.1.1]{.secno}
                Examples](web-messaging.html#examples-5)
            2.  [[9.4.1.2]{.secno} Ports as the basis of an
                object-capability model on the
                web](web-messaging.html#ports-as-the-basis-of-an-object-capability-model-on-the-web)
            3.  [[9.4.1.3]{.secno} Ports as the basis of abstracting out
                service
                implementations](web-messaging.html#ports-as-the-basis-of-abstracting-out-service-implementations)
        2.  [[9.4.2]{.secno} Message
            channels](web-messaging.html#message-channels)
        3.  [[9.4.3]{.secno} The `MessageEventTarget`
            mixin](web-messaging.html#the-messageeventtarget-mixin)
        4.  [[9.4.4]{.secno} Message
            ports](web-messaging.html#message-ports)
        5.  [[9.4.5]{.secno} Ports and garbage
            collection](web-messaging.html#ports-and-garbage-collection)
    3.  [[9.5]{.secno} Broadcasting to other browsing
        contexts](web-messaging.html#broadcasting-to-other-browsing-contexts)
    :::

### [9.3]{.secno} [Cross-document messaging]{#crossDocumentMessages .dfn}[](#web-messaging){.self-link} {#web-messaging}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Window/postMessage](https://developer.mozilla.org/en-US/docs/Web/API/Window/postMessage "The window.postMessage() method safely enables cross-origin communication between Window objects; e.g., between a page and a pop-up that it spawned, or between a page and an iframe embedded within it.")

Support in all current engines.

::: support
[Firefox3+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome2+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera9.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android≤37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

Web browsers, for security and privacy reasons, prevent documents in
different domains from affecting each other; that is, cross-site
scripting is disallowed.

While this is an important security feature, it prevents pages from
different domains from communicating even when those pages are not
hostile. This section introduces a messaging system that allows
documents to communicate with each other regardless of their source
domain, in a way designed to not enable cross-site scripting attacks.

The
[`postMessage()`{#web-messaging:dom-window-postmessage}](#dom-window-postmessage)
API can be used as a [tracking
vector](https://infra.spec.whatwg.org/#tracking-vector){#web-messaging:tracking-vector
x-internal="tracking-vector"}.

#### [9.3.1]{.secno} Introduction[](#introduction-12){.self-link} {#introduction-12}

*This section is non-normative.*

::: example
For example, if document A contains an
[`iframe`{#introduction-12:the-iframe-element}](iframe-embed-object.html#the-iframe-element)
element that contains document B, and script in document A calls
[`postMessage()`{#introduction-12:dom-window-postmessage}](#dom-window-postmessage)
on the [`Window`{#introduction-12:window}](nav-history-apis.html#window)
object of document B, then a message event will be fired on that object,
marked as originating from the
[`Window`{#introduction-12:window-2}](nav-history-apis.html#window) of
document A. The script in document A might look like:

``` js
var o = document.getElementsByTagName('iframe')[0];
o.contentWindow.postMessage('Hello world', 'https://b.example.org/');
```

To register an event handler for incoming events, the script would use
`addEventListener()` (or similar mechanisms). For example, the script in
document B might look like:

``` js
window.addEventListener('message', receiver, false);
function receiver(e) {
  if (e.origin == 'https://example.com') {
    if (e.data == 'Hello world') {
      e.source.postMessage('Hello', e.origin);
    } else {
      alert(e.data);
    }
  }
}
```

This script first checks the domain is the expected domain, and then
looks at the message, which it either displays to the user, or responds
to by sending a message back to the document which sent the message in
the first place.
:::

#### [9.3.2]{.secno} Security[](#security-postmsg){.self-link} {#security-postmsg}

##### [9.3.2.1]{.secno} Authors[](#authors){.self-link}

Use of this API requires extra care to protect users from hostile
entities abusing a site for their own purposes.

Authors should check the
[`origin`{#authors:dom-messageevent-origin}](comms.html#dom-messageevent-origin)
attribute to ensure that messages are only accepted from domains that
they expect to receive messages from. Otherwise, bugs in the author\'s
message handling code could be exploited by hostile sites.

Furthermore, even after checking the
[`origin`{#authors:dom-messageevent-origin-2}](comms.html#dom-messageevent-origin)
attribute, authors should also check that the data in question is of the
expected format. Otherwise, if the source of the event has been attacked
using a cross-site scripting flaw, further unchecked processing of
information sent using the
[`postMessage()`{#authors:dom-window-postmessage}](#dom-window-postmessage)
method could result in the attack being propagated into the receiver.

Authors should not use the wildcard keyword (\*) in the
`targetOrigin`{.variable} argument in messages that contain any
confidential information, as otherwise there is no way to guarantee that
the message is only delivered to the recipient to which it was intended.

------------------------------------------------------------------------

Authors who accept messages from any origin are encouraged to consider
the risks of a denial-of-service attack. An attacker could send a high
volume of messages; if the receiving page performs expensive computation
or causes network traffic to be sent for each such message, the
attacker\'s message could be multiplied into a denial-of-service attack.
Authors are encouraged to employ rate limiting (only accepting a certain
number of messages per minute) to make such attacks impractical.

##### [9.3.2.2]{.secno} User agents[](#user-agents){.self-link}

The integrity of this API is based on the inability for scripts of one
[origin](browsers.html#concept-origin){#user-agents:concept-origin} to
post arbitrary events (using `dispatchEvent()` or otherwise) to objects
in other origins (those that are not the
[same](browsers.html#same-origin){#user-agents:same-origin}).

Implementers are urged to take extra care in the implementation of this
feature. It allows authors to transmit information from one domain to
another domain, which is normally disallowed for security reasons. It
also requires that UAs be careful to allow access to certain properties
but not others.

------------------------------------------------------------------------

User agents are also encouraged to consider rate-limiting message
traffic between different
[origins](browsers.html#concept-origin){#user-agents:concept-origin-2},
to protect naïve sites from denial-of-service attacks.

#### [9.3.3]{.secno} Posting messages[](#posting-messages){.self-link}

`window`{.variable}`.`[`postMessage`](#dom-window-postmessage-options){#dom-window-postmessage-options-dev}`(``message`{.variable}` [, ``options`{.variable}` ])`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Window/postMessage](https://developer.mozilla.org/en-US/docs/Web/API/Window/postMessage "The window.postMessage() method safely enables cross-origin communication between Window objects; e.g., between a page and a pop-up that it spawned, or between a page and an iframe embedded within it.")

Support in all current engines.

::: support
[Firefox3+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome2+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera9.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

Posts a message to the given window. Messages can be structured objects,
e.g. nested objects and arrays, can contain JavaScript values (strings,
numbers,
[`Date`{#posting-messages:date}](https://tc39.es/ecma262/#sec-date-objects){x-internal="date"}
objects, etc.), and can contain certain data objects such as
[`File`{#posting-messages:file}](https://w3c.github.io/FileAPI/#dfn-file){x-internal="file"}
[`Blob`{#posting-messages:blob}](https://w3c.github.io/FileAPI/#dfn-Blob){x-internal="blob"},
[`FileList`{#posting-messages:filelist}](https://w3c.github.io/FileAPI/#filelist-section){x-internal="filelist"},
and
[`ArrayBuffer`{#posting-messages:idl-arraybuffer}](https://webidl.spec.whatwg.org/#idl-ArrayBuffer){x-internal="idl-arraybuffer"}
objects.

Objects listed in the
[`transfer`{#posting-messages:dom-structuredserializeoptions-transfer}](#dom-structuredserializeoptions-transfer)
member of `options`{.variable} are transferred, not just cloned, meaning
that they are no longer usable on the sending side.

A target origin can be specified using the
[`targetOrigin`{#posting-messages:dom-windowpostmessageoptions-targetorigin}](nav-history-apis.html#dom-windowpostmessageoptions-targetorigin)
member of `options`{.variable}. If not provided, it defaults to \"`/`\".
This default restricts the message to same-origin targets only.

If the origin of the target window doesn\'t match the given target
origin, the message is discarded, to avoid information leakage. To send
the message to the target regardless of origin, set the target origin to
\"`*`\".

Throws a
[\"`DataCloneError`\"](https://webidl.spec.whatwg.org/#datacloneerror){#posting-messages:datacloneerror
x-internal="datacloneerror"}
[`DOMException`{#posting-messages:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if `transfer`{.variable} array contains duplicate objects or if
`message`{.variable} could not be cloned.

`window`{.variable}`.`[`postMessage`](#dom-window-postmessage){#posting-messages:dom-window-postmessage}`(``message`{.variable}`, ``targetOrigin`{.variable}` [, ``transfer`{.variable}` ])`

This is an alternate version of
[`postMessage()`{#posting-messages:dom-window-postmessage-options}](#dom-window-postmessage-options)
where the target origin is specified as a parameter. Calling
`window.postMessage(message, target, transfer)` is equivalent to
`window.postMessage(message, {targetOrigin, transfer})`.

When posting a message to a
[`Window`{#posting-messages:window}](nav-history-apis.html#window) of a
[browsing
context](document-sequences.html#browsing-context){#posting-messages:browsing-context}
that has just been navigated to a new
[`Document`{#posting-messages:document}](dom.html#document) is likely to
result in the message not receiving its intended recipient: the scripts
in the target [browsing
context](document-sequences.html#browsing-context){#posting-messages:browsing-context-2}
have to have had time to set up listeners for the messages. Thus, for
instance, in situations where a message is to be sent to the
[`Window`{#posting-messages:window-2}](nav-history-apis.html#window) of
newly created child
[`iframe`{#posting-messages:the-iframe-element}](iframe-embed-object.html#the-iframe-element),
authors are advised to have the child
[`Document`{#posting-messages:document-2}](dom.html#document) post a
message to their parent announcing their readiness to receive messages,
and for the parent to wait for this message before beginning posting
messages.

The [window post message steps]{#window-post-message-steps .dfn}, given
a `targetWindow`{.variable}, `message`{.variable}, and
`options`{.variable}, are as follows:

1.  Let `targetRealm`{.variable} be `targetWindow`{.variable}\'s
    [realm](webappapis.html#concept-global-object-realm){#posting-messages:concept-global-object-realm}.

2.  Let `incumbentSettings`{.variable} be the [incumbent settings
    object](webappapis.html#incumbent-settings-object){#posting-messages:incumbent-settings-object}.

3.  Let `targetOrigin`{.variable} be
    `options`{.variable}\[\"[`targetOrigin`{#posting-messages:dom-windowpostmessageoptions-targetorigin-2}](nav-history-apis.html#dom-windowpostmessageoptions-targetorigin)\"\].

4.  If `targetOrigin`{.variable} is a single U+002F SOLIDUS character
    (/), then set `targetOrigin`{.variable} to
    `incumbentSettings`{.variable}\'s
    [origin](webappapis.html#concept-settings-object-origin){#posting-messages:concept-settings-object-origin}.

5.  Otherwise, if `targetOrigin`{.variable} is not a single U+002A
    ASTERISK character (\*), then:

    1.  Let `parsedURL`{.variable} be the result of running the [URL
        parser](https://url.spec.whatwg.org/#concept-url-parser){#posting-messages:url-parser
        x-internal="url-parser"} on `targetOrigin`{.variable}.

    2.  If `parsedURL`{.variable} is failure, then throw a
        [\"`SyntaxError`\"](https://webidl.spec.whatwg.org/#syntaxerror){#posting-messages:syntaxerror
        x-internal="syntaxerror"}
        [`DOMException`{#posting-messages:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    3.  Set `targetOrigin`{.variable} to `parsedURL`{.variable}\'s
        [origin](https://url.spec.whatwg.org/#concept-url-origin){#posting-messages:concept-url-origin
        x-internal="concept-url-origin"}.

6.  Let `transfer`{.variable} be
    `options`{.variable}\[\"[`transfer`{#posting-messages:dom-structuredserializeoptions-transfer-2}](#dom-structuredserializeoptions-transfer)\"\].

7.  Let `serializeWithTransferResult`{.variable} be
    [StructuredSerializeWithTransfer](structured-data.html#structuredserializewithtransfer){#posting-messages:structuredserializewithtransfer}(`message`{.variable},
    `transfer`{.variable}). Rethrow any exceptions.

8.  [Queue a global
    task](webappapis.html#queue-a-global-task){#posting-messages:queue-a-global-task}
    on the [posted message task source]{#posted-message-task-source
    .dfn} given `targetWindow`{.variable} to run the following steps:

    1.  If the `targetOrigin`{.variable} argument is not a single
        literal U+002A ASTERISK character (\*) and
        `targetWindow`{.variable}\'s [associated
        `Document`](nav-history-apis.html#concept-document-window){#posting-messages:concept-document-window}\'s
        [origin](https://dom.spec.whatwg.org/#concept-document-origin){#posting-messages:concept-document-origin
        x-internal="concept-document-origin"} is not [same
        origin](browsers.html#same-origin){#posting-messages:same-origin}
        with `targetOrigin`{.variable}, then return.

    2.  Let `origin`{.variable} be the
        [serialization](browsers.html#ascii-serialisation-of-an-origin){#posting-messages:ascii-serialisation-of-an-origin}
        of `incumbentSettings`{.variable}\'s
        [origin](webappapis.html#concept-settings-object-origin){#posting-messages:concept-settings-object-origin-2}.

    3.  Let `source`{.variable} be the
        [`WindowProxy`{#posting-messages:windowproxy}](nav-history-apis.html#windowproxy)
        object corresponding to `incumbentSettings`{.variable}\'s
        [global
        object](webappapis.html#concept-settings-object-global){#posting-messages:concept-settings-object-global}
        (a
        [`Window`{#posting-messages:window-3}](nav-history-apis.html#window)
        object).

    4.  Let `deserializeRecord`{.variable} be
        [StructuredDeserializeWithTransfer](structured-data.html#structureddeserializewithtransfer){#posting-messages:structureddeserializewithtransfer}(`serializeWithTransferResult`{.variable},
        `targetRealm`{.variable}).

        If this throws an exception, catch it, [fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#posting-messages:concept-event-fire
        x-internal="concept-event-fire"} named
        [`messageerror`{#posting-messages:event-messageerror}](indices.html#event-messageerror)
        at `targetWindow`{.variable}, using
        [`MessageEvent`{#posting-messages:messageevent}](comms.html#messageevent),
        with the
        [`origin`{#posting-messages:dom-messageevent-origin}](comms.html#dom-messageevent-origin)
        attribute initialized to `origin`{.variable} and the
        [`source`{#posting-messages:dom-messageevent-source}](comms.html#dom-messageevent-source)
        attribute initialized to `source`{.variable}, and then return.

    5.  Let `messageClone`{.variable} be
        `deserializeRecord`{.variable}.\[\[Deserialized\]\].

    6.  Let `newPorts`{.variable} be a new [frozen
        array](https://webidl.spec.whatwg.org/#dfn-frozen-array-type){#posting-messages:frozen-array
        x-internal="frozen-array"} consisting of all
        [`MessagePort`{#posting-messages:messageport}](#messageport)
        objects in
        `deserializeRecord`{.variable}.\[\[TransferredValues\]\], if
        any, maintaining their relative order.

    7.  [Fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#posting-messages:concept-event-fire-2
        x-internal="concept-event-fire"} named
        [`message`{#posting-messages:event-message}](indices.html#event-message)
        at `targetWindow`{.variable}, using
        [`MessageEvent`{#posting-messages:messageevent-2}](comms.html#messageevent),
        with the
        [`origin`{#posting-messages:dom-messageevent-origin-2}](comms.html#dom-messageevent-origin)
        attribute initialized to `origin`{.variable}, the
        [`source`{#posting-messages:dom-messageevent-source-2}](comms.html#dom-messageevent-source)
        attribute initialized to `source`{.variable}, the
        [`data`{#posting-messages:dom-messageevent-data}](comms.html#dom-messageevent-data)
        attribute initialized to `messageClone`{.variable}, and the
        [`ports`{#posting-messages:dom-messageevent-ports}](comms.html#dom-messageevent-ports)
        attribute initialized to `newPorts`{.variable}.

The [`Window`{#posting-messages:window-4}](nav-history-apis.html#window)
interface\'s
[`postMessage(``message`{.variable}`, ``options`{.variable}`)`]{#dom-window-postmessage-options
.dfn dfn-for="Window" dfn-type="method"} method steps are to run the
[window post message
steps](#window-post-message-steps){#posting-messages:window-post-message-steps}
given
[this](https://webidl.spec.whatwg.org/#this){#posting-messages:this
x-internal="this"}, `message`{.variable}, and `options`{.variable}.

The [`Window`{#posting-messages:window-5}](nav-history-apis.html#window)
interface\'s
[`postMessage(``message`{.variable}`, ``targetOrigin`{.variable}`, ``transfer`{.variable}`)`]{#dom-window-postmessage
.dfn dfn-for="Window" dfn-type="method"} method steps are to run the
[window post message
steps](#window-post-message-steps){#posting-messages:window-post-message-steps-2}
given
[this](https://webidl.spec.whatwg.org/#this){#posting-messages:this-2
x-internal="this"}, `message`{.variable}, and «\[
\"[`targetOrigin`{#posting-messages:dom-windowpostmessageoptions-targetorigin-3}](nav-history-apis.html#dom-windowpostmessageoptions-targetorigin)\"
→ `targetOrigin`{.variable},
\"[`transfer`{#posting-messages:dom-structuredserializeoptions-transfer-3}](#dom-structuredserializeoptions-transfer)\"
→ `transfer`{.variable} \]».

### [9.4]{.secno} [Channel messaging]{.dfn}[](#channel-messaging){.self-link}

::::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Channel_Messaging_API](https://developer.mozilla.org/en-US/docs/Web/API/Channel_Messaging_API "The Channel Messaging API allows two separate scripts running in different browsing contexts attached to the same document (e.g., two IFrames, or the main document and an IFrame, two documents via a SharedWorker, or two workers) to communicate directly, passing messages between one another through two-way channels (or pipes) with a port at each end.")

Support in all current engines.

::: support
[Firefox41+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome2+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

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

:::: feature
[Channel_Messaging_API/Using_channel_messaging](https://developer.mozilla.org/en-US/docs/Web/API/Channel_Messaging_API/Using_channel_messaging "The Channel Messaging API allows two separate scripts running in different browsing contexts attached to the same document (e.g., two <iframe> elements, the main document and a single <iframe>, or two documents via a SharedWorker) to communicate directly, passing messages between each other through two-way channels (or pipes) with a port at each end.")

Support in all current engines.

::: support
[Firefox41+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome2+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

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
:::::::

#### [9.4.1]{.secno} Introduction[](#introduction-13){.self-link} {#introduction-13}

*This section is non-normative.*

To enable independent pieces of code (e.g. running in different
[browsing
contexts](document-sequences.html#browsing-context){#introduction-13:browsing-context})
to communicate directly, authors can use [channel
messaging](#channel-messaging){#introduction-13:channel-messaging}.

Communication channels in this mechanism are implemented as two-ways
pipes, with a port at each end. Messages sent in one port are delivered
at the other port, and vice-versa. Messages are delivered as DOM events,
without interrupting or blocking running
[tasks](webappapis.html#concept-task){#introduction-13:concept-task}.

To create a connection (two \"entangled\" ports), the `MessageChannel()`
constructor is called:

``` js
var channel = new MessageChannel();
```

One of the ports is kept as the local port, and the other port is sent
to the remote code, e.g. using
[`postMessage()`{#introduction-13:dom-window-postmessage}](#dom-window-postmessage):

``` js
otherWindow.postMessage('hello', 'https://example.com', [channel.port2]);
```

To send messages, the
[`postMessage()`{#introduction-13:dom-messageport-postmessage}](#dom-messageport-postmessage)
method on the port is used:

``` js
channel.port1.postMessage('hello');
```

To receive messages, one listens to
[`message`{#introduction-13:event-message}](indices.html#event-message)
events:

``` js
channel.port1.onmessage = handleMessage;
function handleMessage(event) {
  // message is in event.data
  // ...
}
```

Data sent on a port can be structured data; for example here an array of
strings is passed on a
[`MessagePort`{#introduction-13:messageport}](#messageport):

``` js
port1.postMessage(['hello', 'world']);
```

##### [9.4.1.1]{.secno} Examples[](#examples-5){.self-link} {#examples-5}

*This section is non-normative.*

::: example
In this example, two JavaScript libraries are connected to each other
using [`MessagePort`{#examples-5:messageport}](#messageport)s. This
allows the libraries to later be hosted in different frames, or in
[`Worker`{#examples-5:worker}](workers.html#worker) objects, without any
change to the APIs.

``` html
<script src="contacts.js"></script> <!-- exposes a contacts object -->
<script src="compose-mail.js"></script> <!-- exposes a composer object -->
<script>
 var channel = new MessageChannel();
 composer.addContactsProvider(channel.port1);
 contacts.registerConsumer(channel.port2);
</script>
```

Here\'s what the \"addContactsProvider()\" function\'s implementation
could look like:

``` js
function addContactsProvider(port) {
  port.onmessage = function (event) {
    switch (event.data.messageType) {
      case 'search-result': handleSearchResult(event.data.results); break;
      case 'search-done': handleSearchDone(); break;
      case 'search-error': handleSearchError(event.data.message); break;
      // ...
    }
  };
};
```

Alternatively, it could be implemented as follows:

``` js
function addContactsProvider(port) {
  port.addEventListener('message', function (event) {
    if (event.data.messageType == 'search-result')
      handleSearchResult(event.data.results);
  });
  port.addEventListener('message', function (event) {
    if (event.data.messageType == 'search-done')
      handleSearchDone();
  });
  port.addEventListener('message', function (event) {
    if (event.data.messageType == 'search-error')
      handleSearchError(event.data.message);
  });
  // ...
  port.start();
};
```

The key difference is that when using
[`addEventListener()`{#examples-5:dom-eventtarget-addeventlistener}](https://dom.spec.whatwg.org/#dom-eventtarget-addeventlistener){x-internal="dom-eventtarget-addeventlistener"},
the
[`start()`{#examples-5:dom-messageport-start}](#dom-messageport-start)
method must also be invoked. When using
[`onmessage`{#examples-5:handler-messageeventtarget-onmessage}](#handler-messageeventtarget-onmessage),
the call to
[`start()`{#examples-5:dom-messageport-start-2}](#dom-messageport-start)
is implied.

The
[`start()`{#examples-5:dom-messageport-start-3}](#dom-messageport-start)
method, whether called explicitly or implicitly (by setting
[`onmessage`{#examples-5:handler-messageeventtarget-onmessage-2}](#handler-messageeventtarget-onmessage)),
starts the flow of messages: messages posted on message ports are
initially paused, so that they don\'t get dropped on the floor before
the script has had a chance to set up its handlers.
:::

##### [9.4.1.2]{.secno} Ports as the basis of an object-capability model on the web[](#ports-as-the-basis-of-an-object-capability-model-on-the-web){.self-link}

*This section is non-normative.*

Ports can be viewed as a way to expose limited capabilities (in the
object-capability model sense) to other actors in the system. This can
either be a weak capability system, where the ports are merely used as a
convenient model within a particular origin, or as a strong capability
model, where they are provided by one origin `provider`{.variable} as
the only mechanism by which another origin `consumer`{.variable} can
effect change in or obtain information from `provider`{.variable}.

For example, consider a situation in which a social web site embeds in
one
[`iframe`{#ports-as-the-basis-of-an-object-capability-model-on-the-web:the-iframe-element}](iframe-embed-object.html#the-iframe-element)
the user\'s email contacts provider (an address book site, from a second
origin), and in a second
[`iframe`{#ports-as-the-basis-of-an-object-capability-model-on-the-web:the-iframe-element-2}](iframe-embed-object.html#the-iframe-element)
a game (from a third origin). The outer social site and the game in the
second
[`iframe`{#ports-as-the-basis-of-an-object-capability-model-on-the-web:the-iframe-element-3}](iframe-embed-object.html#the-iframe-element)
cannot access anything inside the first
[`iframe`{#ports-as-the-basis-of-an-object-capability-model-on-the-web:the-iframe-element-4}](iframe-embed-object.html#the-iframe-element);
together they can only:

- [Navigate](browsing-the-web.html#navigate){#ports-as-the-basis-of-an-object-capability-model-on-the-web:navigate}
  the
  [`iframe`{#ports-as-the-basis-of-an-object-capability-model-on-the-web:the-iframe-element-5}](iframe-embed-object.html#the-iframe-element)
  to a new
  [URL](https://url.spec.whatwg.org/#concept-url){#ports-as-the-basis-of-an-object-capability-model-on-the-web:url
  x-internal="url"}, such as the same
  [URL](https://url.spec.whatwg.org/#concept-url){#ports-as-the-basis-of-an-object-capability-model-on-the-web:url-2
  x-internal="url"} but with a different
  [fragment](https://url.spec.whatwg.org/#concept-url-fragment){#ports-as-the-basis-of-an-object-capability-model-on-the-web:concept-url-fragment
  x-internal="concept-url-fragment"}, causing the
  [`Window`{#ports-as-the-basis-of-an-object-capability-model-on-the-web:window}](nav-history-apis.html#window)
  in the
  [`iframe`{#ports-as-the-basis-of-an-object-capability-model-on-the-web:the-iframe-element-6}](iframe-embed-object.html#the-iframe-element)
  to receive a
  [`hashchange`{#ports-as-the-basis-of-an-object-capability-model-on-the-web:event-hashchange}](indices.html#event-hashchange)
  event.

- Resize the
  [`iframe`{#ports-as-the-basis-of-an-object-capability-model-on-the-web:the-iframe-element-7}](iframe-embed-object.html#the-iframe-element),
  causing the
  [`Window`{#ports-as-the-basis-of-an-object-capability-model-on-the-web:window-2}](nav-history-apis.html#window)
  in the
  [`iframe`{#ports-as-the-basis-of-an-object-capability-model-on-the-web:the-iframe-element-8}](iframe-embed-object.html#the-iframe-element)
  to receive a
  [`resize`{#ports-as-the-basis-of-an-object-capability-model-on-the-web:event-resize}](https://drafts.csswg.org/cssom-view/#eventdef-window-resize){x-internal="event-resize"}
  event.

- Send a
  [`message`{#ports-as-the-basis-of-an-object-capability-model-on-the-web:event-message}](indices.html#event-message)
  event to the
  [`Window`{#ports-as-the-basis-of-an-object-capability-model-on-the-web:window-3}](nav-history-apis.html#window)
  in the
  [`iframe`{#ports-as-the-basis-of-an-object-capability-model-on-the-web:the-iframe-element-9}](iframe-embed-object.html#the-iframe-element)
  using the
  [`window.postMessage()`{#ports-as-the-basis-of-an-object-capability-model-on-the-web:dom-window-postmessage}](#dom-window-postmessage)
  API.

The contacts provider can use these methods, most particularly the third
one, to provide an API that can be accessed by other origins to
manipulate the user\'s address book. For example, it could respond to a
message \"`add-contact Guillaume Tell <tell@pomme.example.net>`\" by
adding the given person and email address to the user\'s address book.

To avoid any site on the web being able to manipulate the user\'s
contacts, the contacts provider might only allow certain trusted sites,
such as the social site, to do this.

Now suppose the game wanted to add a contact to the user\'s address
book, and that the social site was willing to allow it to do so on its
behalf, essentially \"sharing\" the trust that the contacts provider had
with the social site. There are several ways it could do this; most
simply, it could just proxy messages between the game site and the
contacts site. However, this solution has a number of difficulties: it
requires the social site to either completely trust the game site not to
abuse the privilege, or it requires that the social site verify each
request to make sure it\'s not a request that it doesn\'t want to allow
(such as adding multiple contacts, reading the contacts, or deleting
them); it also requires some additional complexity if there\'s ever the
possibility of multiple games simultaneously trying to interact with the
contacts provider.

Using message channels and
[`MessagePort`{#ports-as-the-basis-of-an-object-capability-model-on-the-web:messageport}](#messageport)
objects, however, all of these problems can go away. When the game tells
the social site that it wants to add a contact, the social site can ask
the contacts provider not for it to add a contact, but for the
*capability* to add a single contact. The contacts provider then creates
a pair of
[`MessagePort`{#ports-as-the-basis-of-an-object-capability-model-on-the-web:messageport-2}](#messageport)
objects, and sends one of them back to the social site, who forwards it
on to the game. The game and the contacts provider then have a direct
connection, and the contacts provider knows to only honor a single \"add
contact\" request, nothing else. In other words, the game has been
granted the capability to add a single contact.

##### [9.4.1.3]{.secno} Ports as the basis of abstracting out service implementations[](#ports-as-the-basis-of-abstracting-out-service-implementations){.self-link}

*This section is non-normative.*

Continuing the example from the previous section, consider the contacts
provider in particular. While an initial implementation might have
simply used
[`XMLHttpRequest`{#ports-as-the-basis-of-abstracting-out-service-implementations:xmlhttprequest}](https://xhr.spec.whatwg.org/#xmlhttprequest){x-internal="xmlhttprequest"}
objects in the service\'s
[`iframe`{#ports-as-the-basis-of-abstracting-out-service-implementations:the-iframe-element}](iframe-embed-object.html#the-iframe-element),
an evolution of the service might instead want to use a [shared
worker](workers.html#sharedworker){#ports-as-the-basis-of-abstracting-out-service-implementations:sharedworker}
with a single
[`WebSocket`{#ports-as-the-basis-of-abstracting-out-service-implementations:websocket-2}](https://websockets.spec.whatwg.org/#websocket){x-internal="websocket-2"}
connection.

If the initial design used
[`MessagePort`{#ports-as-the-basis-of-abstracting-out-service-implementations:messageport}](#messageport)
objects to grant capabilities, or even just to allow multiple
simultaneous independent sessions, the service implementation can switch
from the
[`XMLHttpRequest`{#ports-as-the-basis-of-abstracting-out-service-implementations:xmlhttprequest-2}](https://xhr.spec.whatwg.org/#xmlhttprequest){x-internal="xmlhttprequest"}s-in-each-[`iframe`{#ports-as-the-basis-of-abstracting-out-service-implementations:the-iframe-element-2}](iframe-embed-object.html#the-iframe-element)
model to the
shared-[`WebSocket`{#ports-as-the-basis-of-abstracting-out-service-implementations:websocket-2-2}](https://websockets.spec.whatwg.org/#websocket){x-internal="websocket-2"}
model without changing the API at all: the ports on the service provider
side can all be forwarded to the shared worker without it affecting the
users of the API in the slightest.

#### [9.4.2]{.secno} Message channels[](#message-channels){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[MessageChannel](https://developer.mozilla.org/en-US/docs/Web/API/MessageChannel "The MessageChannel interface of the Channel Messaging API allows us to create a new message channel and send data through it via its two MessagePort properties.")

Support in all current engines.

::: support
[Firefox41+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome2+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

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

``` idl
[Exposed=(Window,Worker)]
interface MessageChannel {
  constructor();

  readonly attribute MessagePort port1;
  readonly attribute MessagePort port2;
};
```

`channel`{.variable}` = new `[`MessageChannel`](#dom-messagechannel){#dom-messagechannel-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[MessageChannel/MessageChannel](https://developer.mozilla.org/en-US/docs/Web/API/MessageChannel/MessageChannel "The MessageChannel() constructor of the MessageChannel interface returns a new MessageChannel object with two new MessagePort objects.")

Support in all current engines.

::: support
[Firefox41+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome2+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

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

Returns a new
[`MessageChannel`{#message-channels:messagechannel}](#messagechannel)
object with two new
[`MessagePort`{#message-channels:messageport-3}](#messageport) objects.

`channel`{.variable}`.`[`port1`](#dom-messagechannel-port1){#dom-messagechannel-port1-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[MessageChannel/port1](https://developer.mozilla.org/en-US/docs/Web/API/MessageChannel/port1 "The port1 read-only property of the MessageChannel interface returns the first port of the message channel — the port attached to the context that originated the channel.")

Support in all current engines.

::: support
[Firefox41+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome2+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

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

Returns the first
[`MessagePort`{#message-channels:messageport-4}](#messageport) object.

`channel`{.variable}`.`[`port2`](#dom-messagechannel-port2){#dom-messagechannel-port2-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[MessageChannel/port2](https://developer.mozilla.org/en-US/docs/Web/API/MessageChannel/port2 "The port2 read-only property of the MessageChannel interface returns the second port of the message channel — the port attached to the context at the other end of the channel, which the message is initially sent to.")

Support in all current engines.

::: support
[Firefox41+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome2+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

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

Returns the second
[`MessagePort`{#message-channels:messageport-5}](#messageport) object.

A
[`MessageChannel`{#message-channels:messagechannel-2}](#messagechannel)
object has an associated [port 1]{#port-1 .dfn} and an associated [port
2]{#port-2 .dfn}, both
[`MessagePort`{#message-channels:messageport-6}](#messageport) objects.

The [`new MessageChannel()`]{#dom-messagechannel .dfn} constructor steps
are:

1.  Set
    [this](https://webidl.spec.whatwg.org/#this){#message-channels:this
    x-internal="this"}\'s [port 1](#port-1){#message-channels:port-1} to
    a [new](https://webidl.spec.whatwg.org/#new){#message-channels:new
    x-internal="new"}
    [`MessagePort`{#message-channels:messageport-7}](#messageport) in
    [this](https://webidl.spec.whatwg.org/#this){#message-channels:this-2
    x-internal="this"}\'s [relevant
    realm](webappapis.html#concept-relevant-realm){#message-channels:concept-relevant-realm}.

2.  Set
    [this](https://webidl.spec.whatwg.org/#this){#message-channels:this-3
    x-internal="this"}\'s [port 2](#port-2){#message-channels:port-2} to
    a [new](https://webidl.spec.whatwg.org/#new){#message-channels:new-2
    x-internal="new"}
    [`MessagePort`{#message-channels:messageport-8}](#messageport) in
    [this](https://webidl.spec.whatwg.org/#this){#message-channels:this-4
    x-internal="this"}\'s [relevant
    realm](webappapis.html#concept-relevant-realm){#message-channels:concept-relevant-realm-2}.

3.  [Entangle](#entangle){#message-channels:entangle}
    [this](https://webidl.spec.whatwg.org/#this){#message-channels:this-5
    x-internal="this"}\'s [port 1](#port-1){#message-channels:port-1-2}
    and
    [this](https://webidl.spec.whatwg.org/#this){#message-channels:this-6
    x-internal="this"}\'s [port 2](#port-2){#message-channels:port-2-2}.

The [`port1`]{#dom-messagechannel-port1 .dfn dfn-for="MessageChannel"
dfn-type="attribute"} getter steps are to return
[this](https://webidl.spec.whatwg.org/#this){#message-channels:this-7
x-internal="this"}\'s [port 1](#port-1){#message-channels:port-1-3}.

The [`port2`]{#dom-messagechannel-port2 .dfn dfn-for="MessageChannel"
dfn-type="attribute"} getter steps are to return
[this](https://webidl.spec.whatwg.org/#this){#message-channels:this-8
x-internal="this"}\'s [port 2](#port-2){#message-channels:port-2-3}.

#### [9.4.3]{.secno} []{#the-messageeventtarget-abstract-interface}The [`MessageEventTarget`{#the-messageeventtarget-mixin:messageeventtarget}](#messageeventtarget) mixin[](#the-messageeventtarget-mixin){.self-link}

``` idl
interface mixin MessageEventTarget {
  attribute EventHandler onmessage;
  attribute EventHandler onmessageerror;
};
```

The following are the [event
handlers](webappapis.html#event-handlers){#the-messageeventtarget-mixin:event-handlers}
(and their corresponding [event handler event
types](webappapis.html#event-handler-event-type){#the-messageeventtarget-mixin:event-handler-event-type})
that must be supported, as [event handler IDL
attributes](webappapis.html#event-handler-idl-attributes){#the-messageeventtarget-mixin:event-handler-idl-attributes},
by objects implementing the
[`MessageEventTarget`{#the-messageeventtarget-mixin:messageeventtarget-2}](#messageeventtarget)
interface:

[Event
handler](webappapis.html#event-handlers){#the-messageeventtarget-mixin:event-handlers-2}

[Event handler event
type](webappapis.html#event-handler-event-type){#the-messageeventtarget-mixin:event-handler-event-type-2}

[]{#handler-messageport-onmessage}[]{#handler-worker-onmessage}[]{#handler-dedicatedworkerglobalscope-onmessage}[`onmessage`]{#handler-messageeventtarget-onmessage
.dfn dfn-for="MessageEventTarget" dfn-type="attribute"}

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[MessagePort/message_event](https://developer.mozilla.org/en-US/docs/Web/API/MessagePort/message_event "The message event is fired on a MessagePort object when a message arrives on that channel.")

Support in all current engines.

::: support
[Firefox41+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome2+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

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

:::: feature
[DedicatedWorkerGlobalScope/message_event](https://developer.mozilla.org/en-US/docs/Web/API/DedicatedWorkerGlobalScope/message_event "The message event is fired on a DedicatedWorkerGlobalScope object when the worker receives a message from its parent (i.e. when the parent sends a message using Worker.postMessage()).")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11.5+]{.opera_android .yes}
:::
::::
:::::::

[`message`{#the-messageeventtarget-mixin:event-message}](indices.html#event-message)

[]{#handler-messageport-onmessageerror}[]{#handler-worker-onmessageerror}[]{#handler-dedicatedworkerglobalscope-onmessageerror}[`onmessageerror`]{#handler-messageeventtarget-onmessageerror
.dfn dfn-for="MessageEventTarget" dfn-type="attribute"}

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[MessagePort/messageerror_event](https://developer.mozilla.org/en-US/docs/Web/API/MessagePort/messageerror_event "The messageerror event is fired on a MessagePort object when it receives a message that can't be deserialized.")

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

:::: feature
[DedicatedWorkerGlobalScope/messageerror_event](https://developer.mozilla.org/en-US/docs/Web/API/DedicatedWorkerGlobalScope/messageerror_event "The messageerror event is fired on a DedicatedWorkerGlobalScope object when it receives a message that can't be deserialized.")

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

[`messageerror`{#the-messageeventtarget-mixin:event-messageerror}](indices.html#event-messageerror)

#### [9.4.4]{.secno} Message ports[](#message-ports){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[MessagePort](https://developer.mozilla.org/en-US/docs/Web/API/MessagePort "The MessagePort interface of the Channel Messaging API represents one of the two ports of a MessageChannel, allowing messages to be sent from one port and listening out for them arriving at the other.")

Support in all current engines.

::: support
[Firefox41+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome2+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

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

Each channel has two message ports. Data sent through one port is
received by the other port, and vice versa.

``` idl
[Exposed=(Window,Worker,AudioWorklet), Transferable]
interface MessagePort : EventTarget {
  undefined postMessage(any message, sequence<object> transfer);
  undefined postMessage(any message, optional StructuredSerializeOptions options = {});
  undefined start();
  undefined close();

  // event handlers
  attribute EventHandler onclose;
};

MessagePort includes MessageEventTarget;

dictionary StructuredSerializeOptions {
  sequence<object> transfer = [];
};
```

`port`{.variable}`.`[`postMessage`](#dom-messageport-postmessage){#dom-messageport-postmessage-dev}`(``message`{.variable}` [, ``transfer`{.variable}`])`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[MessagePort/postMessage](https://developer.mozilla.org/en-US/docs/Web/API/MessagePort/postMessage "The postMessage() method of the MessagePort interface sends a message from the port, and optionally, transfers ownership of objects to other browsing contexts.")

Support in all current engines.

::: support
[Firefox41+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome2+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

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

`port`{.variable}`.`[`postMessage`](#dom-messageport-postmessage-options){#dom-messageport-postmessage-options-dev}`(``message`{.variable}` [, { ``transfer`` }])`

Posts a message through the channel. Objects listed in
`transfer`{.variable} are transferred, not just cloned, meaning that
they are no longer usable on the sending side.

Throws a
[\"`DataCloneError`\"](https://webidl.spec.whatwg.org/#datacloneerror){#message-ports:datacloneerror
x-internal="datacloneerror"}
[`DOMException`{#message-ports:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if `transfer`{.variable} contains duplicate objects or
`port`{.variable}, or if `message`{.variable} could not be cloned.

`port`{.variable}`.`[`start`](#dom-messageport-start){#dom-messageport-start-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[MessagePort/start](https://developer.mozilla.org/en-US/docs/Web/API/MessagePort/start "The start() method of the MessagePort interface starts the sending of messages queued on the port. This method is only needed when using EventTarget.addEventListener; it is implied when using onmessage.")

Support in all current engines.

::: support
[Firefox41+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome2+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

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

Begins dispatching messages received on the port.

`port`{.variable}`.`[`close`](#dom-messageport-close){#dom-messageport-close-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[MessagePort/close](https://developer.mozilla.org/en-US/docs/Web/API/MessagePort/close "The close() method of the MessagePort interface disconnects the port, so it is no longer active. This stops the flow of messages to that port.")

Support in all current engines.

::: support
[Firefox41+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome2+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

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

Disconnects the port, so that it is no longer active.

Each [`MessagePort`{#message-ports:messageport-2}](#messageport) object
has a [message event target]{#message-event-target .dfn} (a
[`MessageEventTarget`{#message-ports:messageeventtarget-2}](#messageeventtarget)),
to which the
[`message`{#message-ports:event-message}](indices.html#event-message)
and
[`messageerror`{#message-ports:event-messageerror}](indices.html#event-messageerror)
events are dispatched. Unless otherwise specified, it defaults to the
[`MessagePort`{#message-ports:messageport-3}](#messageport) object
itself.

Each [`MessagePort`{#message-ports:messageport-4}](#messageport) object
can be entangled with another (a symmetric relationship). Each
[`MessagePort`{#message-ports:messageport-5}](#messageport) object also
has a [task
source](webappapis.html#task-source){#message-ports:task-source} called
the [port message queue]{#port-message-queue .dfn}, initially empty. A
[port message
queue](#port-message-queue){#message-ports:port-message-queue} can be
enabled or disabled, and is initially disabled. Once enabled, a port can
never be disabled again (though messages in the queue can get moved to
another queue or removed altogether, which has much the same effect). A
[`MessagePort`{#message-ports:messageport-6}](#messageport) also has a
[has been shipped]{#has-been-shipped .dfn} flag, which must initially be
false.

When a port\'s [port message
queue](#port-message-queue){#message-ports:port-message-queue-2} is
enabled, the [event
loop](webappapis.html#event-loop){#message-ports:event-loop} must use it
as one of its [task
sources](webappapis.html#task-source){#message-ports:task-source-2}.
When a port\'s [relevant global
object](webappapis.html#concept-relevant-global){#message-ports:concept-relevant-global}
is a [`Window`{#message-ports:window}](nav-history-apis.html#window),
all [tasks](webappapis.html#concept-task){#message-ports:concept-task}
[queued](webappapis.html#queue-a-task){#message-ports:queue-a-task} on
its [port message
queue](#port-message-queue){#message-ports:port-message-queue-3} must be
associated with the port\'s [relevant global
object](webappapis.html#concept-relevant-global){#message-ports:concept-relevant-global-2}\'s
[associated
`Document`](nav-history-apis.html#concept-document-window){#message-ports:concept-document-window}.

If the document is [fully
active](document-sequences.html#fully-active){#message-ports:fully-active},
but the event listeners were all created in the context of documents
that are *not* [fully
active](document-sequences.html#fully-active){#message-ports:fully-active-2},
then the messages will not be received unless and until the documents
become [fully
active](document-sequences.html#fully-active){#message-ports:fully-active-3}
again.

Each [event
loop](webappapis.html#event-loop){#message-ports:event-loop-2} has a
[task source](webappapis.html#task-source){#message-ports:task-source-3}
called the [unshipped port message queue]{#unshipped-port-message-queue
.dfn}. This is a virtual [task
source](webappapis.html#task-source){#message-ports:task-source-4}: it
must act as if it contained the
[tasks](webappapis.html#concept-task){#message-ports:concept-task-2} of
each [port message
queue](#port-message-queue){#message-ports:port-message-queue-4} of each
[`MessagePort`{#message-ports:messageport-7}](#messageport) whose [has
been shipped](#has-been-shipped){#message-ports:has-been-shipped} flag
is false, whose [port message
queue](#port-message-queue){#message-ports:port-message-queue-5} is
enabled, and whose [relevant
agent](webappapis.html#relevant-agent){#message-ports:relevant-agent}\'s
[event
loop](webappapis.html#concept-agent-event-loop){#message-ports:concept-agent-event-loop}
is that [event
loop](webappapis.html#event-loop){#message-ports:event-loop-3}, in the
order in which they were added to their respective [task
source](webappapis.html#task-source){#message-ports:task-source-5}. When
a [task](webappapis.html#concept-task){#message-ports:concept-task-3}
would be removed from the [unshipped port message
queue](#unshipped-port-message-queue){#message-ports:unshipped-port-message-queue},
it must instead be removed from its [port message
queue](#port-message-queue){#message-ports:port-message-queue-6}.

When a [`MessagePort`{#message-ports:messageport-8}](#messageport)\'s
[has been shipped](#has-been-shipped){#message-ports:has-been-shipped-2}
flag is false, its [port message
queue](#port-message-queue){#message-ports:port-message-queue-7} must be
ignored for the purposes of the [event
loop](webappapis.html#event-loop){#message-ports:event-loop-4}. (The
[unshipped port message
queue](#unshipped-port-message-queue){#message-ports:unshipped-port-message-queue-2}
is used instead.)

The [has been
shipped](#has-been-shipped){#message-ports:has-been-shipped-3} flag is
set to true when a port, its twin, or the object it was cloned from, is
or has been transferred. When a
[`MessagePort`{#message-ports:messageport-9}](#messageport)\'s [has been
shipped](#has-been-shipped){#message-ports:has-been-shipped-4} flag is
true, its [port message
queue](#port-message-queue){#message-ports:port-message-queue-8} acts as
a first-class [task
source](webappapis.html#task-source){#message-ports:task-source-6},
unaffected to any [unshipped port message
queue](#unshipped-port-message-queue){#message-ports:unshipped-port-message-queue-3}.

When the user agent is to [entangle]{#entangle .dfn} two
[`MessagePort`{#message-ports:messageport-10}](#messageport) objects, it
must run the following steps:

1.  If one of the ports is already entangled, then disentangle it and
    the port that it was entangled with.

    If those two previously entangled ports were the two ports of a
    [`MessageChannel`{#message-ports:messagechannel}](#messagechannel)
    object, then that
    [`MessageChannel`{#message-ports:messagechannel-2}](#messagechannel)
    object no longer represents an actual channel: the two ports in that
    object are no longer entangled.

2.  Associate the two ports to be entangled, so that they form the two
    parts of a new channel. (There is no
    [`MessageChannel`{#message-ports:messagechannel-3}](#messagechannel)
    object that represents this channel.)

    Two ports `A`{.variable} and `B`{.variable} that have gone through
    this step are now said to be entangled; one is entangled to the
    other, and vice versa.

    While this specification describes this process as instantaneous,
    implementations are more likely to implement it via message passing.
    As with all algorithms, the key is \"merely\" that the end result be
    indistinguishable, in a black-box sense, from the specification.

The [disentangle]{#disentangle .dfn} steps, given a
[`MessagePort`{#message-ports:messageport-11}](#messageport)
`initiatorPort`{.variable} which initiates disentangling, are as
follows:

1.  Let `otherPort`{.variable} be the
    [`MessagePort`{#message-ports:messageport-12}](#messageport) which
    `initiatorPort`{.variable} was entangled with.

2.  [Assert](https://infra.spec.whatwg.org/#assert){#message-ports:assert
    x-internal="assert"}: `otherPort`{.variable} exists.

3.  Disentangle `initiatorPort`{.variable} and `otherPort`{.variable},
    so that they are no longer entangled or associated with each other.

4.  [Fire an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#message-ports:concept-event-fire
    x-internal="concept-event-fire"} named
    [`close`{#message-ports:event-close}](indices.html#event-close) at
    `otherPort`{.variable}.

::: note
The [`close`{#message-ports:event-close-2}](indices.html#event-close)
event will be fired even if the port is not explicitly closed. The cases
where this event is dispatched are:

- the
  [`close()`{#message-ports:dom-messageport-close-2}](#dom-messageport-close)
  method was called;
- the [`Document`{#message-ports:document}](dom.html#document) was
  [destroyed](document-lifecycle.html#destroy-a-document){#message-ports:destroy-a-document};
  or
- the [`MessagePort`{#message-ports:messageport-13}](#messageport) was
  [garbage collected](#ports-and-garbage-collection).

We only dispatch the event on `otherPort`{.variable} because
`initiatorPort`{.variable} explicitly triggered the close, its
[`Document`{#message-ports:document-2}](dom.html#document) no longer
exists, or it was already garbage collected, as described above.
:::

------------------------------------------------------------------------

[`MessagePort`{#message-ports:messageport-14}](#messageport) objects are
[transferable
objects](structured-data.html#transferable-objects){#message-ports:transferable-objects}.
Their [transfer
steps](structured-data.html#transfer-steps){#message-ports:transfer-steps},
given `value`{.variable} and `dataHolder`{.variable}, are:

1.  Set `value`{.variable}\'s [has been
    shipped](#has-been-shipped){#message-ports:has-been-shipped-5} flag
    to true.

2.  Set `dataHolder`{.variable}.\[\[PortMessageQueue\]\] to
    `value`{.variable}\'s [port message
    queue](#port-message-queue){#message-ports:port-message-queue-9}.

3.  If `value`{.variable} is entangled with another port
    `remotePort`{.variable}, then:

    1.  Set `remotePort`{.variable}\'s [has been
        shipped](#has-been-shipped){#message-ports:has-been-shipped-6}
        flag to true.

    2.  Set `dataHolder`{.variable}.\[\[RemotePort\]\] to
        `remotePort`{.variable}.

4.  Otherwise, set `dataHolder`{.variable}.\[\[RemotePort\]\] to null.

Their [transfer-receiving
steps](structured-data.html#transfer-receiving-steps){#message-ports:transfer-receiving-steps},
given `dataHolder`{.variable} and `value`{.variable}, are:

1.  Set `value`{.variable}\'s [has been
    shipped](#has-been-shipped){#message-ports:has-been-shipped-7} flag
    to true.

2.  Move all the
    [tasks](webappapis.html#concept-task){#message-ports:concept-task-4}
    that are to fire
    [`message`{#message-ports:event-message-2}](indices.html#event-message)
    events in `dataHolder`{.variable}.\[\[PortMessageQueue\]\] to the
    [port message
    queue](#port-message-queue){#message-ports:port-message-queue-10} of
    `value`{.variable}, if any, leaving `value`{.variable}\'s [port
    message
    queue](#port-message-queue){#message-ports:port-message-queue-11} in
    its initial disabled state, and, if `value`{.variable}\'s [relevant
    global
    object](webappapis.html#concept-relevant-global){#message-ports:concept-relevant-global-3}
    is a
    [`Window`{#message-ports:window-2}](nav-history-apis.html#window),
    associating the moved
    [tasks](webappapis.html#concept-task){#message-ports:concept-task-5}
    with `value`{.variable}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#message-ports:concept-relevant-global-4}\'s
    [associated
    `Document`](nav-history-apis.html#concept-document-window){#message-ports:concept-document-window-2}.

3.  If `dataHolder`{.variable}.\[\[RemotePort\]\] is not null, then
    [entangle](#entangle){#message-ports:entangle}
    `dataHolder`{.variable}.\[\[RemotePort\]\] and `value`{.variable}.
    (This will disentangle `dataHolder`{.variable}.\[\[RemotePort\]\]
    from the original port that was transferred.)

------------------------------------------------------------------------

The [message port post message steps]{#message-port-post-message-steps
.dfn}, given `sourcePort`{.variable}, `targetPort`{.variable},
`message`{.variable}, and `options`{.variable} are as follows:

1.  Let `transfer`{.variable} be
    `options`{.variable}\[\"[`transfer`{#message-ports:dom-structuredserializeoptions-transfer}](#dom-structuredserializeoptions-transfer)\"\].

2.  If `transfer`{.variable}
    [contains](https://infra.spec.whatwg.org/#list-contain){#message-ports:list-contains
    x-internal="list-contains"} `sourcePort`{.variable}, then throw a
    [\"`DataCloneError`\"](https://webidl.spec.whatwg.org/#datacloneerror){#message-ports:datacloneerror-2
    x-internal="datacloneerror"}
    [`DOMException`{#message-ports:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

3.  Let `doomed`{.variable} be false.

4.  If `targetPort`{.variable} is not null and `transfer`{.variable}
    [contains](https://infra.spec.whatwg.org/#list-contain){#message-ports:list-contains-2
    x-internal="list-contains"} `targetPort`{.variable}, then set
    `doomed`{.variable} to true and optionally report to a developer
    console that the target port was posted to itself, causing the
    communication channel to be lost.

5.  Let `serializeWithTransferResult`{.variable} be
    [StructuredSerializeWithTransfer](structured-data.html#structuredserializewithtransfer){#message-ports:structuredserializewithtransfer}(`message`{.variable},
    `transfer`{.variable}). Rethrow any exceptions.

6.  If `targetPort`{.variable} is null, or if `doomed`{.variable} is
    true, then return.

7.  Add a
    [task](webappapis.html#concept-task){#message-ports:concept-task-6}
    that runs the following steps to the [port message
    queue](#port-message-queue){#message-ports:port-message-queue-12} of
    `targetPort`{.variable}:

    1.  Let `finalTargetPort`{.variable} be the
        [`MessagePort`{#message-ports:messageport-15}](#messageport) in
        whose [port message
        queue](#port-message-queue){#message-ports:port-message-queue-13}
        the task now finds itself.

        This can be different from `targetPort`{.variable}, if
        `targetPort`{.variable} itself was transferred and thus all its
        tasks moved along with it.

    2.  Let `messageEventTarget`{.variable} be
        `finalTargetPort`{.variable}\'s [message event
        target](#message-event-target){#message-ports:message-event-target}.

    3.  Let `targetRealm`{.variable} be `finalTargetPort`{.variable}\'s
        [relevant
        realm](webappapis.html#concept-relevant-realm){#message-ports:concept-relevant-realm}.

    4.  Let `deserializeRecord`{.variable} be
        [StructuredDeserializeWithTransfer](structured-data.html#structureddeserializewithtransfer){#message-ports:structureddeserializewithtransfer}(`serializeWithTransferResult`{.variable},
        `targetRealm`{.variable}).

        If this throws an exception, catch it, [fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#message-ports:concept-event-fire-2
        x-internal="concept-event-fire"} named
        [`messageerror`{#message-ports:event-messageerror-2}](indices.html#event-messageerror)
        at `messageEventTarget`{.variable}, using
        [`MessageEvent`{#message-ports:messageevent}](comms.html#messageevent),
        and then return.

    5.  Let `messageClone`{.variable} be
        `deserializeRecord`{.variable}.\[\[Deserialized\]\].

    6.  Let `newPorts`{.variable} be a new [frozen
        array](https://webidl.spec.whatwg.org/#dfn-frozen-array-type){#message-ports:frozen-array
        x-internal="frozen-array"} consisting of all
        [`MessagePort`{#message-ports:messageport-16}](#messageport)
        objects in
        `deserializeRecord`{.variable}.\[\[TransferredValues\]\], if
        any, maintaining their relative order.

    7.  [Fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#message-ports:concept-event-fire-3
        x-internal="concept-event-fire"} named
        [`message`{#message-ports:event-message-3}](indices.html#event-message)
        at `messageEventTarget`{.variable}, using
        [`MessageEvent`{#message-ports:messageevent-2}](comms.html#messageevent),
        with the
        [`data`{#message-ports:dom-messageevent-data}](comms.html#dom-messageevent-data)
        attribute initialized to `messageClone`{.variable} and the
        [`ports`{#message-ports:dom-messageevent-ports}](comms.html#dom-messageevent-ports)
        attribute initialized to `newPorts`{.variable}.

The
[`postMessage(``message`{.variable}`, ``options`{.variable}`)`]{#dom-messageport-postmessage-options
.dfn dfn-for="MessagePort" dfn-type="method"} method steps are:

1.  Let `targetPort`{.variable} be the port with which
    [this](https://webidl.spec.whatwg.org/#this){#message-ports:this
    x-internal="this"} is entangled, if any; otherwise let it be null.

2.  Run the [message port post message
    steps](#message-port-post-message-steps){#message-ports:message-port-post-message-steps}
    providing
    [this](https://webidl.spec.whatwg.org/#this){#message-ports:this-2
    x-internal="this"}, `targetPort`{.variable}, `message`{.variable}
    and `options`{.variable}.

The
[`postMessage(``message`{.variable}`, ``transfer`{.variable}`)`]{#dom-messageport-postmessage
.dfn dfn-for="MessagePort" dfn-type="method"} method steps are:

1.  Let `targetPort`{.variable} be the port with which
    [this](https://webidl.spec.whatwg.org/#this){#message-ports:this-3
    x-internal="this"} is entangled, if any; otherwise let it be null.

2.  Let `options`{.variable} be «\[
    \"[`transfer`{#message-ports:dom-structuredserializeoptions-transfer-2}](#dom-structuredserializeoptions-transfer)\"
    → `transfer`{.variable} \]».

3.  Run the [message port post message
    steps](#message-port-post-message-steps){#message-ports:message-port-post-message-steps-2}
    providing
    [this](https://webidl.spec.whatwg.org/#this){#message-ports:this-4
    x-internal="this"}, `targetPort`{.variable}, `message`{.variable}
    and `options`{.variable}.

------------------------------------------------------------------------

The [`start()`]{#dom-messageport-start .dfn dfn-for="MessagePort"
dfn-type="method"} method steps are to enable
[this](https://webidl.spec.whatwg.org/#this){#message-ports:this-5
x-internal="this"}\'s [port message
queue](#port-message-queue){#message-ports:port-message-queue-14}, if it
is not already enabled.

------------------------------------------------------------------------

The [`close()`]{#dom-messageport-close .dfn dfn-for="MessagePort"
dfn-type="method"} method steps are:

1.  Set
    [this](https://webidl.spec.whatwg.org/#this){#message-ports:this-6
    x-internal="this"}\'s
    [\[\[Detached\]\]](structured-data.html#detached){#message-ports:detached}
    internal slot value to true.

2.  If
    [this](https://webidl.spec.whatwg.org/#this){#message-ports:this-7
    x-internal="this"} is entangled,
    [disentangle](#disentangle){#message-ports:disentangle} it.

------------------------------------------------------------------------

The following are the [event
handlers](webappapis.html#event-handlers){#message-ports:event-handlers}
(and their corresponding [event handler event
types](webappapis.html#event-handler-event-type){#message-ports:event-handler-event-type})
that must be supported, as [event handler IDL
attributes](webappapis.html#event-handler-idl-attributes){#message-ports:event-handler-idl-attributes},
by all objects implementing the
[`MessagePort`{#message-ports:messageport-17}](#messageport) interface:

[Event
handler](webappapis.html#event-handlers){#message-ports:event-handlers-2}

[Event handler event
type](webappapis.html#event-handler-event-type){#message-ports:event-handler-event-type-2}

[`onclose`]{#handler-messageport-onclose .dfn dfn-for="MessagePort"
dfn-type="attribute"}

[`close`{#message-ports:event-close-3}](indices.html#event-close)

The first time a
[`MessagePort`{#message-ports:messageport-18}](#messageport) object\'s
[`onmessage`{#message-ports:handler-messageeventtarget-onmessage}](#handler-messageeventtarget-onmessage)
IDL attribute is set, the port\'s [port message
queue](#port-message-queue){#message-ports:port-message-queue-15} must
be enabled, as if the
[`start()`{#message-ports:dom-messageport-start-2}](#dom-messageport-start)
method had been called.

#### [9.4.5]{.secno} Ports and garbage collection[](#ports-and-garbage-collection){.self-link}

When a
[`MessagePort`{#ports-and-garbage-collection:messageport}](#messageport)
object `o`{.variable} is garbage collected, if `o`{.variable} is
entangled, then the user agent must
[disentangle](#disentangle){#ports-and-garbage-collection:disentangle}
`o`{.variable}.

When a
[`MessagePort`{#ports-and-garbage-collection:messageport-2}](#messageport)
object `o`{.variable} is entangled and
[`message`{#ports-and-garbage-collection:event-message}](indices.html#event-message)
or
[`messageerror`{#ports-and-garbage-collection:event-messageerror}](indices.html#event-messageerror)
event listener is registered, user agents must act as if
`o`{.variable}\'s entangled
[`MessagePort`{#ports-and-garbage-collection:messageport-3}](#messageport)
object has a strong reference to `o`{.variable}.

Furthermore, a
[`MessagePort`{#ports-and-garbage-collection:messageport-4}](#messageport)
object must not be garbage collected while there exists an event
referenced by a
[task](webappapis.html#concept-task){#ports-and-garbage-collection:concept-task}
in a [task
queue](webappapis.html#task-queue){#ports-and-garbage-collection:task-queue}
that is to be dispatched on that
[`MessagePort`{#ports-and-garbage-collection:messageport-5}](#messageport)
object, or while the
[`MessagePort`{#ports-and-garbage-collection:messageport-6}](#messageport)
object\'s [port message
queue](#port-message-queue){#ports-and-garbage-collection:port-message-queue}
is enabled and not empty.

::: note
Thus, a message port can be received, given an event listener, and then
forgotten, and so long as that event listener could receive a message,
the channel will be maintained.

Of course, if this was to occur on both sides of the channel, then both
ports could be garbage collected, since they would not be reachable from
live code, despite having a strong reference to each other. However, if
a message port has a pending message, it is not garbage collected.
:::

Authors are strongly encouraged to explicitly close
[`MessagePort`{#ports-and-garbage-collection:messageport-7}](#messageport)
objects to disentangle them, so that their resources can be recollected.
Creating many
[`MessagePort`{#ports-and-garbage-collection:messageport-8}](#messageport)
objects and discarding them without closing them can lead to high
transient memory usage since garbage collection is not necessarily
performed promptly, especially for
[`MessagePort`{#ports-and-garbage-collection:messageport-9}](#messageport)s
where garbage collection can involve cross-process coordination.

### [9.5]{.secno} [Broadcasting to other browsing contexts]{.dfn}[](#broadcasting-to-other-browsing-contexts){.self-link}

::::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[BroadcastChannel](https://developer.mozilla.org/en-US/docs/Web/API/BroadcastChannel "The BroadcastChannel interface represents a named channel that any browsing context of a given origin can subscribe to. It allows communication between different documents (in different windows, tabs, frames or iframes) of the same origin. Messages are broadcasted via a message event fired at all BroadcastChannel objects listening to the channel, except the object that sent the message.")

Support in all current engines.

::: support
[Firefox38+]{.firefox .yes}[Safari15.4+]{.safari
.yes}[Chrome54+]{.chrome .yes}

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

:::: feature
[Broadcast_Channel_API](https://developer.mozilla.org/en-US/docs/Web/API/Broadcast_Channel_API "The Broadcast Channel API allows basic communication between browsing contexts (that is, windows, tabs, frames, or iframes) and workers on the same origin.")

Support in all current engines.

::: support
[Firefox38+]{.firefox .yes}[Safari15.4+]{.safari
.yes}[Chrome54+]{.chrome .yes}

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
:::::::

Pages on a single
[origin](browsers.html#concept-origin){#broadcasting-to-other-browsing-contexts:concept-origin}
opened by the same user in the same user agent but in different
unrelated [browsing
contexts](document-sequences.html#browsing-context){#broadcasting-to-other-browsing-contexts:browsing-context}
sometimes need to send notifications to each other, for example \"hey,
the user logged in over here, check your credentials again\".

For elaborate cases, e.g. to manage locking of shared state, to manage
synchronization of resources between a server and multiple local
clients, to share a
[`WebSocket`{#broadcasting-to-other-browsing-contexts:websocket-2}](https://websockets.spec.whatwg.org/#websocket){x-internal="websocket-2"}
connection with a remote host, and so forth, [shared
workers](workers.html#sharedworker){#broadcasting-to-other-browsing-contexts:sharedworker}
are the most appropriate solution.

For simple cases, though, where a shared worker would be an unreasonable
overhead, authors can use the simple channel-based broadcast mechanism
described in this section.

``` idl
[Exposed=(Window,Worker)]
interface BroadcastChannel : EventTarget {
  constructor(DOMString name);

  readonly attribute DOMString name;
  undefined postMessage(any message);
  undefined close();
  attribute EventHandler onmessage;
  attribute EventHandler onmessageerror;
};
```

`broadcastChannel`{.variable}` = new `[`BroadcastChannel`](#dom-broadcastchannel){#dom-broadcastchannel-dev}`(``name`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[BroadcastChannel/BroadcastChannel](https://developer.mozilla.org/en-US/docs/Web/API/BroadcastChannel/BroadcastChannel "The BroadcastChannel() constructor creates a new BroadcastChannel and connects it to the underlying channel.")

Support in all current engines.

::: support
[Firefox38+]{.firefox .yes}[Safari15.4+]{.safari
.yes}[Chrome54+]{.chrome .yes}

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

Returns a new
[`BroadcastChannel`{#broadcasting-to-other-browsing-contexts:broadcastchannel}](#broadcastchannel)
object via which messages for the given channel name can be sent and
received.

`broadcastChannel`{.variable}`.`[`name`](#dom-broadcastchannel-name){#dom-broadcastchannel-name-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[BroadcastChannel/name](https://developer.mozilla.org/en-US/docs/Web/API/BroadcastChannel/name "The read-only BroadcastChannel.name property returns a string, which uniquely identifies the given channel with its name. This name is passed to the BroadcastChannel() constructor at creation time and is therefore read-only.")

Support in all current engines.

::: support
[Firefox38+]{.firefox .yes}[Safari15.4+]{.safari
.yes}[Chrome54+]{.chrome .yes}

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

Returns the channel name (as passed to the constructor).

`broadcastChannel`{.variable}`.`[`postMessage`](#dom-broadcastchannel-postmessage){#dom-broadcastchannel-postmessage-dev}`(``message`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[BroadcastChannel/postMessage](https://developer.mozilla.org/en-US/docs/Web/API/BroadcastChannel/postMessage "The BroadcastChannel.postMessage() sends a message, which can be of any kind of Object, to each listener in any browsing context with the same origin. The message is transmitted as a 'message' event targeted at each BroadcastChannel bound to the channel.")

Support in all current engines.

::: support
[Firefox38+]{.firefox .yes}[Safari15.4+]{.safari
.yes}[Chrome54+]{.chrome .yes}

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

Sends the given message to other
[`BroadcastChannel`{#broadcasting-to-other-browsing-contexts:broadcastchannel-2}](#broadcastchannel)
objects set up for this channel. Messages can be structured objects,
e.g. nested objects and arrays.

`broadcastChannel`{.variable}`.`[`close`](#dom-broadcastchannel-close){#dom-broadcastchannel-close-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[BroadcastChannel/close](https://developer.mozilla.org/en-US/docs/Web/API/BroadcastChannel/close "The BroadcastChannel.close() terminates the connection to the underlying channel, allowing the object to be garbage collected. This is a necessary step to perform as there is no other way for a browser to know that this channel is not needed anymore.")

Support in all current engines.

::: support
[Firefox38+]{.firefox .yes}[Safari15.4+]{.safari
.yes}[Chrome54+]{.chrome .yes}

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

Closes the
[`BroadcastChannel`{#broadcasting-to-other-browsing-contexts:broadcastchannel-3}](#broadcastchannel)
object, opening it up to garbage collection.

A
[`BroadcastChannel`{#broadcasting-to-other-browsing-contexts:broadcastchannel-4}](#broadcastchannel)
object has a [channel name]{#channel-name .dfn} and a [closed
flag]{#concept-broadcastchannel-closed .dfn}.

The [`new BroadcastChannel(``name`{.variable}`)`]{#dom-broadcastchannel
.dfn} constructor steps are:

1.  Set
    [this](https://webidl.spec.whatwg.org/#this){#broadcasting-to-other-browsing-contexts:this
    x-internal="this"}\'s [channel
    name](#channel-name){#broadcasting-to-other-browsing-contexts:channel-name}
    to `name`{.variable}.

2.  Set
    [this](https://webidl.spec.whatwg.org/#this){#broadcasting-to-other-browsing-contexts:this-2
    x-internal="this"}\'s [closed
    flag](#concept-broadcastchannel-closed){#broadcasting-to-other-browsing-contexts:concept-broadcastchannel-closed}
    to false.

The [`name`]{#dom-broadcastchannel-name .dfn dfn-for="BroadcastChannel"
dfn-type="attribute"} getter steps are to return
[this](https://webidl.spec.whatwg.org/#this){#broadcasting-to-other-browsing-contexts:this-3
x-internal="this"}\'s [channel
name](#channel-name){#broadcasting-to-other-browsing-contexts:channel-name-2}.

A
[`BroadcastChannel`{#broadcasting-to-other-browsing-contexts:broadcastchannel-5}](#broadcastchannel)
object is said to be [eligible for messaging]{#eligible-for-messaging
.dfn} when its [relevant global
object](webappapis.html#concept-relevant-global){#broadcasting-to-other-browsing-contexts:concept-relevant-global}
is either:

- a
  [`Window`{#broadcasting-to-other-browsing-contexts:window}](nav-history-apis.html#window)
  object whose [associated
  `Document`](nav-history-apis.html#concept-document-window){#broadcasting-to-other-browsing-contexts:concept-document-window}
  is [fully
  active](document-sequences.html#fully-active){#broadcasting-to-other-browsing-contexts:fully-active},
  or

- a
  [`WorkerGlobalScope`{#broadcasting-to-other-browsing-contexts:workerglobalscope}](workers.html#workerglobalscope)
  object whose
  [closing](workers.html#dom-workerglobalscope-closing){#broadcasting-to-other-browsing-contexts:dom-workerglobalscope-closing}
  flag is false and whose
  [worker](workers.html#worker){#broadcasting-to-other-browsing-contexts:worker}
  is not a [suspendable
  worker](workers.html#suspendable-worker){#broadcasting-to-other-browsing-contexts:suspendable-worker}.

The
[`postMessage(``message`{.variable}`)`]{#dom-broadcastchannel-postmessage
.dfn dfn-for="BroadcastChannel" dfn-type="method"} method steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#broadcasting-to-other-browsing-contexts:this-4
    x-internal="this"} is not [eligible for
    messaging](#eligible-for-messaging){#broadcasting-to-other-browsing-contexts:eligible-for-messaging},
    then return.

2.  If
    [this](https://webidl.spec.whatwg.org/#this){#broadcasting-to-other-browsing-contexts:this-5
    x-internal="this"}\'s [closed
    flag](#concept-broadcastchannel-closed){#broadcasting-to-other-browsing-contexts:concept-broadcastchannel-closed-2}
    is true, then throw an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#broadcasting-to-other-browsing-contexts:invalidstateerror
    x-internal="invalidstateerror"}
    [`DOMException`{#broadcasting-to-other-browsing-contexts:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

3.  Let `serialized`{.variable} be
    [StructuredSerialize](structured-data.html#structuredserialize){#broadcasting-to-other-browsing-contexts:structuredserialize}(`message`{.variable}).
    Rethrow any exceptions.

4.  Let `sourceOrigin`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#broadcasting-to-other-browsing-contexts:this-6
    x-internal="this"}\'s [relevant settings
    object](webappapis.html#relevant-settings-object){#broadcasting-to-other-browsing-contexts:relevant-settings-object}\'s
    [origin](webappapis.html#concept-settings-object-origin){#broadcasting-to-other-browsing-contexts:concept-settings-object-origin}.

5.  Let `sourceStorageKey`{.variable} be the result of running [obtain a
    storage key for non-storage
    purposes](https://storage.spec.whatwg.org/#obtain-a-storage-key-for-non-storage-purposes){#broadcasting-to-other-browsing-contexts:obtain-a-storage-key-for-non-storage-purposes
    x-internal="obtain-a-storage-key-for-non-storage-purposes"} with
    [this](https://webidl.spec.whatwg.org/#this){#broadcasting-to-other-browsing-contexts:this-7
    x-internal="this"}\'s [relevant settings
    object](webappapis.html#relevant-settings-object){#broadcasting-to-other-browsing-contexts:relevant-settings-object-2}.

6.  Let `destinations`{.variable} be a list of
    [`BroadcastChannel`{#broadcasting-to-other-browsing-contexts:broadcastchannel-6}](#broadcastchannel)
    objects that match the following criteria:

    - They are [eligible for
      messaging](#eligible-for-messaging){#broadcasting-to-other-browsing-contexts:eligible-for-messaging-2}.

    - The result of running [obtain a storage key for non-storage
      purposes](https://storage.spec.whatwg.org/#obtain-a-storage-key-for-non-storage-purposes){#broadcasting-to-other-browsing-contexts:obtain-a-storage-key-for-non-storage-purposes-2
      x-internal="obtain-a-storage-key-for-non-storage-purposes"} with
      their [relevant settings
      object](webappapis.html#relevant-settings-object){#broadcasting-to-other-browsing-contexts:relevant-settings-object-3}
      [equals](https://storage.spec.whatwg.org/#storage-key-equal){#broadcasting-to-other-browsing-contexts:storage-key-equal
      x-internal="storage-key-equal"} `sourceStorageKey`{.variable}.

    - Their [channel
      name](#channel-name){#broadcasting-to-other-browsing-contexts:channel-name-3}
      [is](https://infra.spec.whatwg.org/#string-is){#broadcasting-to-other-browsing-contexts:is
      x-internal="is"}
      [this](https://webidl.spec.whatwg.org/#this){#broadcasting-to-other-browsing-contexts:this-8
      x-internal="this"}\'s [channel
      name](#channel-name){#broadcasting-to-other-browsing-contexts:channel-name-4}.

7.  Remove `source`{.variable} from `destinations`{.variable}.

8.  Sort `destinations`{.variable} such that all
    [`BroadcastChannel`{#broadcasting-to-other-browsing-contexts:broadcastchannel-7}](#broadcastchannel)
    objects whose [relevant
    agents](webappapis.html#relevant-agent){#broadcasting-to-other-browsing-contexts:relevant-agent}
    are the same are sorted in creation order, oldest first. (This does
    not define a complete ordering. Within this constraint, user agents
    may sort the list in any
    [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#broadcasting-to-other-browsing-contexts:implementation-defined
    x-internal="implementation-defined"} manner.)

9.  For each `destination`{.variable} in `destinations`{.variable},
    [queue a global
    task](webappapis.html#queue-a-global-task){#broadcasting-to-other-browsing-contexts:queue-a-global-task}
    on the [DOM manipulation task
    source](webappapis.html#dom-manipulation-task-source){#broadcasting-to-other-browsing-contexts:dom-manipulation-task-source}
    given `destination`{.variable}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#broadcasting-to-other-browsing-contexts:concept-relevant-global-2}
    to perform the following steps:

    1.  If `destination`{.variable}\'s [closed
        flag](#concept-broadcastchannel-closed){#broadcasting-to-other-browsing-contexts:concept-broadcastchannel-closed-3}
        is true, then abort these steps.

    2.  Let `targetRealm`{.variable} be `destination`{.variable}\'s
        [relevant
        realm](webappapis.html#concept-relevant-realm){#broadcasting-to-other-browsing-contexts:concept-relevant-realm}.

    3.  Let `data`{.variable} be
        [StructuredDeserialize](structured-data.html#structureddeserialize){#broadcasting-to-other-browsing-contexts:structureddeserialize}(`serialized`{.variable},
        `targetRealm`{.variable}).

        If this throws an exception, catch it, [fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#broadcasting-to-other-browsing-contexts:concept-event-fire
        x-internal="concept-event-fire"} named
        [`messageerror`{#broadcasting-to-other-browsing-contexts:event-messageerror}](indices.html#event-messageerror)
        at `destination`{.variable}, using
        [`MessageEvent`{#broadcasting-to-other-browsing-contexts:messageevent}](comms.html#messageevent),
        with the
        [`origin`{#broadcasting-to-other-browsing-contexts:dom-messageevent-origin}](comms.html#dom-messageevent-origin)
        attribute initialized to the
        [serialization](browsers.html#ascii-serialisation-of-an-origin){#broadcasting-to-other-browsing-contexts:ascii-serialisation-of-an-origin}
        of `sourceOrigin`{.variable}, and then abort these steps.

    4.  [Fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#broadcasting-to-other-browsing-contexts:concept-event-fire-2
        x-internal="concept-event-fire"} named
        [`message`{#broadcasting-to-other-browsing-contexts:event-message}](indices.html#event-message)
        at `destination`{.variable}, using
        [`MessageEvent`{#broadcasting-to-other-browsing-contexts:messageevent-2}](comms.html#messageevent),
        with the
        [`data`{#broadcasting-to-other-browsing-contexts:dom-messageevent-data}](comms.html#dom-messageevent-data)
        attribute initialized to `data`{.variable} and the
        [`origin`{#broadcasting-to-other-browsing-contexts:dom-messageevent-origin-2}](comms.html#dom-messageevent-origin)
        attribute initialized to the
        [serialization](browsers.html#ascii-serialisation-of-an-origin){#broadcasting-to-other-browsing-contexts:ascii-serialisation-of-an-origin-2}
        of `sourceOrigin`{.variable}.

While a
[`BroadcastChannel`{#broadcasting-to-other-browsing-contexts:broadcastchannel-8}](#broadcastchannel)
object whose [closed
flag](#concept-broadcastchannel-closed){#broadcasting-to-other-browsing-contexts:concept-broadcastchannel-closed-4}
is false has an event listener registered for
[`message`{#broadcasting-to-other-browsing-contexts:event-message-2}](indices.html#event-message)
or
[`messageerror`{#broadcasting-to-other-browsing-contexts:event-messageerror-2}](indices.html#event-messageerror)
events, there must be a strong reference from the
[`BroadcastChannel`{#broadcasting-to-other-browsing-contexts:broadcastchannel-9}](#broadcastchannel)
object\'s [relevant global
object](webappapis.html#concept-relevant-global){#broadcasting-to-other-browsing-contexts:concept-relevant-global-3}
to the
[`BroadcastChannel`{#broadcasting-to-other-browsing-contexts:broadcastchannel-10}](#broadcastchannel)
object itself.

The [`close()`]{#dom-broadcastchannel-close .dfn
dfn-for="BroadcastChannel" dfn-type="method"} method steps are to set
[this](https://webidl.spec.whatwg.org/#this){#broadcasting-to-other-browsing-contexts:this-9
x-internal="this"}\'s [closed
flag](#concept-broadcastchannel-closed){#broadcasting-to-other-browsing-contexts:concept-broadcastchannel-closed-5}
to true.

Authors are strongly encouraged to explicitly close
[`BroadcastChannel`{#broadcasting-to-other-browsing-contexts:broadcastchannel-11}](#broadcastchannel)
objects when they are no longer needed, so that they can be garbage
collected. Creating many
[`BroadcastChannel`{#broadcasting-to-other-browsing-contexts:broadcastchannel-12}](#broadcastchannel)
objects and discarding them while leaving them with an event listener
and without closing them can lead to an apparent memory leak, since the
objects will continue to live for as long as they have an event listener
(or until their page or worker is closed).

------------------------------------------------------------------------

The following are the [event
handlers](webappapis.html#event-handlers){#broadcasting-to-other-browsing-contexts:event-handlers}
(and their corresponding [event handler event
types](webappapis.html#event-handler-event-type){#broadcasting-to-other-browsing-contexts:event-handler-event-type})
that must be supported, as [event handler IDL
attributes](webappapis.html#event-handler-idl-attributes){#broadcasting-to-other-browsing-contexts:event-handler-idl-attributes},
by all objects implementing the
[`BroadcastChannel`{#broadcasting-to-other-browsing-contexts:broadcastchannel-13}](#broadcastchannel)
interface:

[Event
handler](webappapis.html#event-handlers){#broadcasting-to-other-browsing-contexts:event-handlers-2}

[Event handler event
type](webappapis.html#event-handler-event-type){#broadcasting-to-other-browsing-contexts:event-handler-event-type-2}

[`onmessage`]{#handler-broadcastchannel-onmessage .dfn
dfn-for="BroadcastChannel" dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[BroadcastChannel/message_event](https://developer.mozilla.org/en-US/docs/Web/API/BroadcastChannel/message_event "The message event is fired on a BroadcastChannel object when a message arrives on that channel.")

Support in all current engines.

::: support
[Firefox38+]{.firefox .yes}[Safari15.4+]{.safari
.yes}[Chrome54+]{.chrome .yes}

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

[`message`{#broadcasting-to-other-browsing-contexts:event-message-3}](indices.html#event-message)

[`onmessageerror`]{#handler-broadcastchannel-onmessageerror .dfn
dfn-for="BroadcastChannel" dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[BroadcastChannel/messageerror_event](https://developer.mozilla.org/en-US/docs/Web/API/BroadcastChannel/messageerror_event "The messageerror event is fired on a BroadcastChannel object when a message that can't be deserialized arrives on the channel.")

Support in all current engines.

::: support
[Firefox57+]{.firefox .yes}[Safari15.4+]{.safari
.yes}[Chrome60+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android47+]{.opera_android .yes}
:::
::::
:::::

[`messageerror`{#broadcasting-to-other-browsing-contexts:event-messageerror-3}](indices.html#event-messageerror)

::: example
Suppose a page wants to know when the user logs out, even when the user
does so from another tab at the same site:

``` js
var authChannel = new BroadcastChannel('auth');
authChannel.onmessage = function (event) {
  if (event.data == 'logout')
    showLogout();
}

function logoutRequested() {
  // called when the user asks us to log them out
  doLogout();
  showLogout();
  authChannel.postMessage('logout');
}

function doLogout() {
  // actually log the user out (e.g. clearing cookies)
  // ...
}

function showLogout() {
  // update the UI to indicate we're logged out
  // ...
}
```
:::

[← 9.2 Server-sent events](server-sent-events.html) --- [Table of
Contents](./) --- [10 Web workers →](workers.html)
