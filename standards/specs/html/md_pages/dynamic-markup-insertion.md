HTML Standard

::: {#head .head .with-buttons}
[![WHATWG](https://resources.whatwg.org/logo.svg){.darkmode-aware
width="100" crossorigin="" height="100"}](https://whatwg.org/){.logo}

# HTML {#html .allcaps}

Living Standard --- Last Updated [2 June 2025]{.pubdate}
:::

[← 8 Web application APIs](webappapis.html) --- [Table of Contents](./)
--- [8.6 Timers →](timers-and-user-prompts.html)

1.  ::: {#toc-webappapis}
    1.  [[8.4]{.secno} Dynamic markup
        insertion](dynamic-markup-insertion.html#dynamic-markup-insertion)
        1.  [[8.4.1]{.secno} Opening the input
            stream](dynamic-markup-insertion.html#opening-the-input-stream)
        2.  [[8.4.2]{.secno} Closing the input
            stream](dynamic-markup-insertion.html#closing-the-input-stream)
        3.  [[8.4.3]{.secno}
            `document.write()`](dynamic-markup-insertion.html#document.write())
        4.  [[8.4.4]{.secno}
            `document.writeln()`](dynamic-markup-insertion.html#document.writeln())
    2.  [[8.5]{.secno} DOM parsing and serialization
        APIs](dynamic-markup-insertion.html#dom-parsing-and-serialization)
        1.  [[8.5.1]{.secno} The `DOMParser`
            interface](dynamic-markup-insertion.html#the-domparser-interface)
        2.  [[8.5.2]{.secno} Unsafe HTML parsing
            methods](dynamic-markup-insertion.html#unsafe-html-parsing-methods)
        3.  [[8.5.3]{.secno} HTML serialization
            methods](dynamic-markup-insertion.html#html-serialization-methods)
        4.  [[8.5.4]{.secno} The `innerHTML`
            property](dynamic-markup-insertion.html#the-innerhtml-property)
        5.  [[8.5.5]{.secno} The `outerHTML`
            property](dynamic-markup-insertion.html#the-outerhtml-property)
        6.  [[8.5.6]{.secno} The `insertAdjacentHTML()`
            method](dynamic-markup-insertion.html#the-insertadjacenthtml()-method)
        7.  [[8.5.7]{.secno} The `createContextualFragment()`
            method](dynamic-markup-insertion.html#the-createcontextualfragment()-method)
        8.  [[8.5.8]{.secno} The `XMLSerializer`
            interface](dynamic-markup-insertion.html#the-xmlserializer-interface)
    :::

### [8.4]{.secno} [Dynamic markup insertion]{.dfn}[](#dynamic-markup-insertion){.self-link}

APIs for dynamically inserting markup into the document interact with
the parser, and thus their behavior varies depending on whether they are
used with [HTML
documents](https://dom.spec.whatwg.org/#html-document){#dynamic-markup-insertion:html-documents
x-internal="html-documents"} (and the [HTML
parser](parsing.html#html-parser){#dynamic-markup-insertion:html-parser})
or [XML
documents](https://dom.spec.whatwg.org/#xml-document){#dynamic-markup-insertion:xml-documents
x-internal="xml-documents"} (and the [XML
parser](xhtml.html#xml-parser){#dynamic-markup-insertion:xml-parser}).

[`Document`{#dynamic-markup-insertion:document}](dom.html#document)
objects have a [throw-on-dynamic-markup-insertion
counter]{#throw-on-dynamic-markup-insertion-counter .dfn}, which is used
in conjunction with the [create an element for the
token](parsing.html#create-an-element-for-the-token){#dynamic-markup-insertion:create-an-element-for-the-token}
algorithm to prevent [custom element
constructors](custom-elements.html#custom-element-constructor){#dynamic-markup-insertion:custom-element-constructor}
from being able to use
[`document.open()`{#dynamic-markup-insertion:dom-document-open}](#dom-document-open),
[`document.close()`{#dynamic-markup-insertion:dom-document-close}](#dom-document-close),
and
[`document.write()`{#dynamic-markup-insertion:dom-document-write}](#dom-document-write)
when they are invoked by the parser. Initially, the counter must be set
to zero.

#### [8.4.1]{.secno} Opening the input stream[](#opening-the-input-stream){.self-link}

`document`{.variable}` = ``document`{.variable}`.`[`open`](#dom-document-open){#dom-document-open-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Document/open](https://developer.mozilla.org/en-US/docs/Web/API/Document/open "The Document.open() method opens a document for writing.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari11+]{.safari .yes}[Chrome64+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera51+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer4+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android47+]{.opera_android .yes}
:::
::::
:::::

Causes the
[`Document`{#opening-the-input-stream:document}](dom.html#document) to
be replaced in-place, as if it was a new
[`Document`{#opening-the-input-stream:document-2}](dom.html#document)
object, but reusing the previous object, which is then returned.

The resulting
[`Document`{#opening-the-input-stream:document-3}](dom.html#document)
has an HTML parser associated with it, which can be given data to parse
using
[`document.write()`{#opening-the-input-stream:dom-document-write}](#dom-document-write).

The method has no effect if the
[`Document`{#opening-the-input-stream:document-4}](dom.html#document) is
still being parsed.

Throws an
[\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#opening-the-input-stream:invalidstateerror
x-internal="invalidstateerror"}
[`DOMException`{#opening-the-input-stream:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if the
[`Document`{#opening-the-input-stream:document-5}](dom.html#document) is
an [XML
document](https://dom.spec.whatwg.org/#xml-document){#opening-the-input-stream:xml-documents
x-internal="xml-documents"}.

Throws an
[\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#opening-the-input-stream:invalidstateerror-2
x-internal="invalidstateerror"}
[`DOMException`{#opening-the-input-stream:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if the parser is currently executing a [custom element
constructor](custom-elements.html#custom-element-constructor){#opening-the-input-stream:custom-element-constructor}.

`window`{.variable}` = ``document`{.variable}`.`[`open`](#dom-document-open-window){#opening-the-input-stream:dom-document-open-window}`(``url`{.variable}`, ``name`{.variable}`, ``features`{.variable}`)`

Works like the
[`window.open()`{#opening-the-input-stream:dom-open}](nav-history-apis.html#dom-open)
method.

[`Document`{#opening-the-input-stream:document-6}](dom.html#document)
objects have an [active parser was aborted]{#active-parser-was-aborted
.dfn} boolean, which is used to prevent scripts from invoking the
[`document.open()`{#opening-the-input-stream:dom-document-open}](#dom-document-open)
and
[`document.write()`{#opening-the-input-stream:dom-document-write-2}](#dom-document-write)
methods (directly or indirectly) after the document\'s [active
parser](dom.html#active-parser){#opening-the-input-stream:active-parser}
has been aborted. It is initially false.

The [document open steps]{#document-open-steps .dfn}, given a
`document`{.variable}, are as follows:

1.  If `document`{.variable} is an [XML
    document](https://dom.spec.whatwg.org/#xml-document){#opening-the-input-stream:xml-documents-2
    x-internal="xml-documents"}, then throw an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#opening-the-input-stream:invalidstateerror-3
    x-internal="invalidstateerror"}
    [`DOMException`{#opening-the-input-stream:domexception-3}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  If `document`{.variable}\'s [throw-on-dynamic-markup-insertion
    counter](#throw-on-dynamic-markup-insertion-counter){#opening-the-input-stream:throw-on-dynamic-markup-insertion-counter}
    is greater than 0, then throw an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#opening-the-input-stream:invalidstateerror-4
    x-internal="invalidstateerror"}
    [`DOMException`{#opening-the-input-stream:domexception-4}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

3.  Let `entryDocument`{.variable} be the [entry global
    object](webappapis.html#entry-global-object){#opening-the-input-stream:entry-global-object}\'s
    [associated
    `Document`](nav-history-apis.html#concept-document-window){#opening-the-input-stream:concept-document-window}.

4.  If `document`{.variable}\'s
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#opening-the-input-stream:concept-document-origin
    x-internal="concept-document-origin"} is not [same
    origin](browsers.html#same-origin){#opening-the-input-stream:same-origin}
    to `entryDocument`{.variable}\'s
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#opening-the-input-stream:concept-document-origin-2
    x-internal="concept-document-origin"}, then throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#opening-the-input-stream:securityerror
    x-internal="securityerror"}
    [`DOMException`{#opening-the-input-stream:domexception-5}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

5.  If `document`{.variable} has an [active
    parser](dom.html#active-parser){#opening-the-input-stream:active-parser-2}
    whose [script nesting
    level](parsing.html#script-nesting-level){#opening-the-input-stream:script-nesting-level}
    is greater than 0, then return `document`{.variable}.

    This basically causes
    [`document.open()`{#opening-the-input-stream:dom-document-open-2}](#dom-document-open)
    to be ignored when it\'s called in an inline script found during
    parsing, while still letting it have an effect when called from a
    non-parser task such as a timer callback or event handler.

6.  Similarly, if `document`{.variable}\'s [unload
    counter](document-lifecycle.html#unload-counter){#opening-the-input-stream:unload-counter}
    is greater than 0, then return `document`{.variable}.

    This basically causes
    [`document.open()`{#opening-the-input-stream:dom-document-open-3}](#dom-document-open)
    to be ignored when it\'s called from a
    [`beforeunload`{#opening-the-input-stream:event-beforeunload}](indices.html#event-beforeunload),
    [`pagehide`{#opening-the-input-stream:event-pagehide}](indices.html#event-pagehide),
    or
    [`unload`{#opening-the-input-stream:event-unload}](indices.html#event-unload)
    event handler while the
    [`Document`{#opening-the-input-stream:document-7}](dom.html#document)
    is being unloaded.

7.  If `document`{.variable}\'s [active parser was
    aborted](#active-parser-was-aborted){#opening-the-input-stream:active-parser-was-aborted}
    is true, then return `document`{.variable}.

    This notably causes
    [`document.open()`{#opening-the-input-stream:dom-document-open-4}](#dom-document-open)
    to be ignored if it is called after a
    [navigation](browsing-the-web.html#navigate){#opening-the-input-stream:navigate}
    has started, but only during the initial parse. See [issue
    #4723](https://github.com/whatwg/html/issues/4723) for more
    background.

8.  If `document`{.variable}\'s [node
    navigable](document-sequences.html#node-navigable){#opening-the-input-stream:node-navigable}
    is non-null and `document`{.variable}\'s [node
    navigable](document-sequences.html#node-navigable){#opening-the-input-stream:node-navigable-2}\'s
    [ongoing
    navigation](browsing-the-web.html#ongoing-navigation){#opening-the-input-stream:ongoing-navigation}
    is a [navigation
    ID](browsing-the-web.html#navigation-id){#opening-the-input-stream:navigation-id},
    then [stop
    loading](document-lifecycle.html#nav-stop){#opening-the-input-stream:nav-stop}
    `document`{.variable}\'s [node
    navigable](document-sequences.html#node-navigable){#opening-the-input-stream:node-navigable-3}.

9.  For each [shadow-including inclusive
    descendant](https://dom.spec.whatwg.org/#concept-shadow-including-inclusive-descendant){#opening-the-input-stream:shadow-including-inclusive-descendant
    x-internal="shadow-including-inclusive-descendant"}
    `node`{.variable} of `document`{.variable}, [erase all event
    listeners and
    handlers](webappapis.html#erase-all-event-listeners-and-handlers){#opening-the-input-stream:erase-all-event-listeners-and-handlers}
    given `node`{.variable}.

10. If `document`{.variable} is the [associated
    `Document`](nav-history-apis.html#concept-document-window){#opening-the-input-stream:concept-document-window-2}
    of `document`{.variable}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#opening-the-input-stream:concept-relevant-global},
    then [erase all event listeners and
    handlers](webappapis.html#erase-all-event-listeners-and-handlers){#opening-the-input-stream:erase-all-event-listeners-and-handlers-2}
    given `document`{.variable}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#opening-the-input-stream:concept-relevant-global-2}.

11. [Replace
    all](https://dom.spec.whatwg.org/#concept-node-replace-all){#opening-the-input-stream:concept-node-replace-all
    x-internal="concept-node-replace-all"} with null within
    `document`{.variable}.

12. If `document`{.variable} is [fully
    active](document-sequences.html#fully-active){#opening-the-input-stream:fully-active},
    then:

    1.  Let `newURL`{.variable} be a copy of
        `entryDocument`{.variable}\'s
        [URL](https://dom.spec.whatwg.org/#concept-document-url){#opening-the-input-stream:the-document's-address
        x-internal="the-document's-address"}.

    2.  If `entryDocument`{.variable} is not `document`{.variable}, then
        set `newURL`{.variable}\'s
        [fragment](https://url.spec.whatwg.org/#concept-url-fragment){#opening-the-input-stream:concept-url-fragment
        x-internal="concept-url-fragment"} to null.

    3.  Run the [URL and history update
        steps](browsing-the-web.html#url-and-history-update-steps){#opening-the-input-stream:url-and-history-update-steps}
        with `document`{.variable} and `newURL`{.variable}.

13. Set `document`{.variable}\'s [is initial
    `about:blank`](dom.html#is-initial-about:blank){#opening-the-input-stream:is-initial-about:blank}
    to false.

14. If `document`{.variable}\'s [iframe load in
    progress](iframe-embed-object.html#iframe-load-in-progress){#opening-the-input-stream:iframe-load-in-progress}
    flag is set, then set `document`{.variable}\'s [mute iframe
    load](iframe-embed-object.html#mute-iframe-load){#opening-the-input-stream:mute-iframe-load}
    flag.

15. Set `document`{.variable} to [no-quirks
    mode](https://dom.spec.whatwg.org/#concept-document-no-quirks){#opening-the-input-stream:no-quirks-mode
    x-internal="no-quirks-mode"}.

16. Create a new [HTML
    parser](parsing.html#html-parser){#opening-the-input-stream:html-parser}
    and associate it with `document`{.variable}. This is a
    [script-created parser]{#script-created-parser .dfn} (meaning that
    it can be closed by the
    [`document.open()`{#opening-the-input-stream:dom-document-open-5}](#dom-document-open)
    and
    [`document.close()`{#opening-the-input-stream:dom-document-close}](#dom-document-close)
    methods, and that the tokenizer will wait for an explicit call to
    [`document.close()`{#opening-the-input-stream:dom-document-close-2}](#dom-document-close)
    before emitting an end-of-file token). The encoding
    [confidence](parsing.html#concept-encoding-confidence){#opening-the-input-stream:concept-encoding-confidence}
    is *irrelevant*.

17. Set the [insertion
    point](parsing.html#insertion-point){#opening-the-input-stream:insertion-point}
    to point at just before the end of the [input
    stream](parsing.html#input-stream){#opening-the-input-stream:input-stream}
    (which at this point will be empty).

18. [Update the current document
    readiness](dom.html#update-the-current-document-readiness){#opening-the-input-stream:update-the-current-document-readiness}
    of `document`{.variable} to \"`loading`\".

    This causes a
    [`readystatechange`{#opening-the-input-stream:event-readystatechange}](indices.html#event-readystatechange)
    event to fire, but the event is actually unobservable to author
    code, because of the previous step which [erased all event listeners
    and
    handlers](webappapis.html#erase-all-event-listeners-and-handlers){#opening-the-input-stream:erase-all-event-listeners-and-handlers-3}
    that could observe it.

19. Return `document`{.variable}.

The [document open
steps](#document-open-steps){#opening-the-input-stream:document-open-steps}
do not affect whether a
[`Document`{#opening-the-input-stream:document-8}](dom.html#document) is
[ready for post-load
tasks](parsing.html#ready-for-post-load-tasks){#opening-the-input-stream:ready-for-post-load-tasks}
or [completely
loaded](document-lifecycle.html#completely-loaded){#opening-the-input-stream:completely-loaded}.

The
[`open(``unused1`{.variable}`, ``unused2`{.variable}`)`]{#dom-document-open
.dfn dfn-for="Document" dfn-type="method"} method must return the result
of running the [document open
steps](#document-open-steps){#opening-the-input-stream:document-open-steps-2}
with
[this](https://webidl.spec.whatwg.org/#this){#opening-the-input-stream:this
x-internal="this"}.

The `unused1`{.variable} and `unused2`{.variable} arguments are ignored,
but kept in the IDL to allow code that calls the function with one or
two arguments to continue working. They are necessary due to Web IDL
[overload resolution
algorithm](https://webidl.spec.whatwg.org/#dfn-overload-resolution-algorithm){#opening-the-input-stream:overload-resolution-algorithm
x-internal="overload-resolution-algorithm"} rules, which would throw a
[`TypeError`{#opening-the-input-stream:typeerror}](https://tc39.es/ecma262/#sec-native-error-types-used-in-this-standard-typeerror){x-internal="typeerror"}
exception for such calls had the arguments not been there.
[whatwg/webidl issue #581](https://github.com/whatwg/webidl/issues/581)
investigates changing the algorithm to allow for their removal.
[\[WEBIDL\]](references.html#refsWEBIDL)

The
[`open(``url`{.variable}`, ``name`{.variable}`, ``features`{.variable}`)`]{#dom-document-open-window
.dfn dfn-for="Document" dfn-type="method"} method must run these steps:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#opening-the-input-stream:this-2
    x-internal="this"} is not [fully
    active](document-sequences.html#fully-active){#opening-the-input-stream:fully-active-2},
    then throw an
    [\"`InvalidAccessError`\"](https://webidl.spec.whatwg.org/#invalidaccesserror){#opening-the-input-stream:invalidaccesserror
    x-internal="invalidaccesserror"}
    [`DOMException`{#opening-the-input-stream:domexception-6}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  Return the result of running the [window open
    steps](nav-history-apis.html#window-open-steps){#opening-the-input-stream:window-open-steps}
    with `url`{.variable}, `name`{.variable}, and `features`{.variable}.

#### [8.4.2]{.secno} Closing the input stream[](#closing-the-input-stream){.self-link}

`document`{.variable}`.`[`close`](#dom-document-close){#dom-document-close-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Document/close](https://developer.mozilla.org/en-US/docs/Web/API/Document/close "The Document.close() method finishes writing to a document, opened with Document.open().")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari11+]{.safari .yes}[Chrome64+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera51+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer4+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android47+]{.opera_android .yes}
:::
::::
:::::

Closes the input stream that was opened by the
[`document.open()`{#closing-the-input-stream:dom-document-open}](#dom-document-open)
method.

Throws an
[\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#closing-the-input-stream:invalidstateerror
x-internal="invalidstateerror"}
[`DOMException`{#closing-the-input-stream:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if the
[`Document`{#closing-the-input-stream:document}](dom.html#document) is
an [XML
document](https://dom.spec.whatwg.org/#xml-document){#closing-the-input-stream:xml-documents
x-internal="xml-documents"}.

Throws an
[\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#closing-the-input-stream:invalidstateerror-2
x-internal="invalidstateerror"}
[`DOMException`{#closing-the-input-stream:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if the parser is currently executing a [custom element
constructor](custom-elements.html#custom-element-constructor){#closing-the-input-stream:custom-element-constructor}.

The [`close()`]{#dom-document-close .dfn dfn-for="Document"
dfn-type="method"} method must run the following steps:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#closing-the-input-stream:this
    x-internal="this"} is an [XML
    document](https://dom.spec.whatwg.org/#xml-document){#closing-the-input-stream:xml-documents-2
    x-internal="xml-documents"}, then throw an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#closing-the-input-stream:invalidstateerror-3
    x-internal="invalidstateerror"}
    [`DOMException`{#closing-the-input-stream:domexception-3}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  If
    [this](https://webidl.spec.whatwg.org/#this){#closing-the-input-stream:this-2
    x-internal="this"}\'s [throw-on-dynamic-markup-insertion
    counter](#throw-on-dynamic-markup-insertion-counter){#closing-the-input-stream:throw-on-dynamic-markup-insertion-counter}
    is greater than zero, then throw an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#closing-the-input-stream:invalidstateerror-4
    x-internal="invalidstateerror"}
    [`DOMException`{#closing-the-input-stream:domexception-4}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

3.  If there is no [script-created
    parser](#script-created-parser){#closing-the-input-stream:script-created-parser}
    associated with
    [this](https://webidl.spec.whatwg.org/#this){#closing-the-input-stream:this-3
    x-internal="this"}, then return.

4.  Insert an [explicit \"EOF\"
    character](parsing.html#explicit-eof-character){#closing-the-input-stream:explicit-eof-character}
    at the end of the parser\'s [input
    stream](parsing.html#input-stream){#closing-the-input-stream:input-stream}.

5.  If
    [this](https://webidl.spec.whatwg.org/#this){#closing-the-input-stream:this-4
    x-internal="this"}\'s [pending parsing-blocking
    script](scripting.html#pending-parsing-blocking-script){#closing-the-input-stream:pending-parsing-blocking-script}
    is not null, then return.

6.  Run the tokenizer, processing resulting tokens as they are emitted,
    and stopping when the tokenizer reaches the [explicit \"EOF\"
    character](parsing.html#explicit-eof-character){#closing-the-input-stream:explicit-eof-character-2}
    or [spins the event
    loop](webappapis.html#spin-the-event-loop){#closing-the-input-stream:spin-the-event-loop}.

#### [8.4.3]{.secno} [`document.write()`{#document.write():dom-document-write}](#dom-document-write)[](#document.write()){.self-link} {#document.write()}

`document`{.variable}`.`[`write`](#dom-document-write){#dom-document-write-dev}`(...``text`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Document/write](https://developer.mozilla.org/en-US/docs/Web/API/Document/write "The document.write() method writes a string of text to a document stream opened by document.open().")

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

In general, adds the given string(s) to the
[`Document`{#document.write():document}](dom.html#document)\'s input
stream.

This method has very idiosyncratic behavior. In some cases, this method
can affect the state of the [HTML
parser](parsing.html#html-parser){#document.write():html-parser} while
the parser is running, resulting in a DOM that does not correspond to
the source of the document (e.g. if the string written is the string
\"`<plaintext>`\" or \"`<!--`\"). In other cases, the call can clear the
current page first, as if
[`document.open()`{#document.write():dom-document-open}](#dom-document-open)
had been called. In yet more cases, the method is simply ignored, or
throws an exception. User agents are [explicitly allowed to avoid
executing `script` elements inserted via this
method](parsing.html#document-written-scripts-intervention). And to make
matters even worse, the exact behavior of this method can in some cases
be dependent on network latency, which can lead to failures that are
very hard to debug. **For all these reasons, use of this method is
strongly discouraged.**

Throws an
[\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#document.write():invalidstateerror
x-internal="invalidstateerror"}
[`DOMException`{#document.write():domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
when invoked on [XML
documents](https://dom.spec.whatwg.org/#xml-document){#document.write():xml-documents
x-internal="xml-documents"}.

Throws an
[\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#document.write():invalidstateerror-2
x-internal="invalidstateerror"}
[`DOMException`{#document.write():domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if the parser is currently executing a [custom element
constructor](custom-elements.html#custom-element-constructor){#document.write():custom-element-constructor}.

This method performs no sanitization to remove potentially-dangerous
elements and attributes like
[`script`{#document.write():the-script-element}](scripting.html#the-script-element)
or [event handler content
attributes](webappapis.html#event-handler-content-attributes){#document.write():event-handler-content-attributes}.

[`Document`{#document.write():document-2}](dom.html#document) objects
have an [ignore-destructive-writes
counter]{#ignore-destructive-writes-counter .dfn}, which is used in
conjunction with the processing of
[`script`{#document.write():the-script-element-2}](scripting.html#the-script-element)
elements to prevent external scripts from being able to use
[`document.write()`{#document.write():dom-document-write-2}](#dom-document-write)
to blow away the document by implicitly calling
[`document.open()`{#document.write():dom-document-open-2}](#dom-document-open).
Initially, the counter must be set to zero.

The [document write steps]{#document-write-steps .dfn}, given a
[`Document`{#document.write():document-3}](dom.html#document) object
`document`{.variable}, a list `text`{.variable}, a boolean
`lineFeed`{.variable}, and a string `sink`{.variable}, are as follows:

1.  Let `string`{.variable} be the empty string.

2.  Let `isTrusted`{.variable} be false if `text`{.variable}
    [contains](https://infra.spec.whatwg.org/#list-contain){#document.write():list-contains
    x-internal="list-contains"} a string; otherwise true.

3.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#document.write():list-iterate
    x-internal="list-iterate"} `value`{.variable} of `text`{.variable}:

    1.  If `value`{.variable} is a
        [`TrustedHTML`{#document.write():tt-trustedhtml}](https://w3c.github.io/trusted-types/dist/spec/#trustedhtml){x-internal="tt-trustedhtml"}
        object, then append `value`{.variable}\'s associated
        [data](https://w3c.github.io/trusted-types/dist/spec/#trustedhtml-data){#document.write():tt-trustedhtml-data
        x-internal="tt-trustedhtml-data"} to `string`{.variable}.

    2.  Otherwise, append `value`{.variable} to `string`{.variable}.

4.  If `isTrusted`{.variable} is false, set `string`{.variable} to the
    result of invoking the [Get Trusted Type compliant
    string](https://w3c.github.io/trusted-types/dist/spec/#get-trusted-type-compliant-string-algorithm){#document.write():tt-getcompliantstring
    x-internal="tt-getcompliantstring"} algorithm with
    [`TrustedHTML`{#document.write():tt-trustedhtml-2}](https://w3c.github.io/trusted-types/dist/spec/#trustedhtml){x-internal="tt-trustedhtml"},
    [this](https://webidl.spec.whatwg.org/#this){#document.write():this
    x-internal="this"}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#document.write():concept-relevant-global},
    `string`{.variable}, `sink`{.variable}, and \"`script`\".

5.  If `lineFeed`{.variable} is true, append U+000A LINE FEED to
    `string`{.variable}.

6.  If `document`{.variable} is an [XML
    document](https://dom.spec.whatwg.org/#xml-document){#document.write():xml-documents-2
    x-internal="xml-documents"}, then throw an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#document.write():invalidstateerror-3
    x-internal="invalidstateerror"}
    [`DOMException`{#document.write():domexception-3}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

7.  If `document`{.variable}\'s [throw-on-dynamic-markup-insertion
    counter](#throw-on-dynamic-markup-insertion-counter){#document.write():throw-on-dynamic-markup-insertion-counter}
    is greater than 0, then throw an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#document.write():invalidstateerror-4
    x-internal="invalidstateerror"}
    [`DOMException`{#document.write():domexception-4}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

8.  If `document`{.variable}\'s [active parser was
    aborted](#active-parser-was-aborted){#document.write():active-parser-was-aborted}
    is true, then return.

9.  If the [insertion
    point](parsing.html#insertion-point){#document.write():insertion-point}
    is undefined, then:

    1.  If `document`{.variable}\'s [unload
        counter](document-lifecycle.html#unload-counter){#document.write():unload-counter}
        is greater than 0 or `document`{.variable}\'s
        [ignore-destructive-writes
        counter](#ignore-destructive-writes-counter){#document.write():ignore-destructive-writes-counter}
        is greater than 0, then return.

    2.  Run the [document open
        steps](#document-open-steps){#document.write():document-open-steps}
        with `document`{.variable}.

10. Insert `string`{.variable} into the [input
    stream](parsing.html#input-stream){#document.write():input-stream}
    just before the [insertion
    point](parsing.html#insertion-point){#document.write():insertion-point-2}.

11. If `document`{.variable}\'s [pending parsing-blocking
    script](scripting.html#pending-parsing-blocking-script){#document.write():pending-parsing-blocking-script}
    is null, then have the [HTML
    parser](parsing.html#html-parser){#document.write():html-parser-2}
    process `string`{.variable}, one code point at a time, processing
    resulting tokens as they are emitted, and stopping when the
    tokenizer reaches the insertion point or when the processing of the
    tokenizer is aborted by the tree construction stage (this can happen
    if a
    [`script`{#document.write():the-script-element-3}](scripting.html#the-script-element)
    end tag token is emitted by the tokenizer).

    If the
    [`document.write()`{#document.write():dom-document-write-3}](#dom-document-write)
    method was called from script executing inline (i.e. executing
    because the parser parsed a set of
    [`script`{#document.write():the-script-element-4}](scripting.html#the-script-element)
    tags), then this is a [reentrant invocation of the
    parser](parsing.html#nestedParsing). If the [parser pause
    flag](parsing.html#parser-pause-flag){#document.write():parser-pause-flag}
    is set, the tokenizer will abort immediately and no HTML will be
    parsed, per the tokenizer\'s [parser pause flag
    check](parsing.html#check-parser-pause-flag).

The [`document.write(...``text`{.variable}`)`]{#dom-document-write .dfn
dfn-for="Document" dfn-type="method"} method steps are to run the
[document write
steps](#document-write-steps){#document.write():document-write-steps}
with
[this](https://webidl.spec.whatwg.org/#this){#document.write():this-2
x-internal="this"}, `text`{.variable}, false, and \"`Document write`\".

#### [8.4.4]{.secno} [`document.writeln()`{#document.writeln():dom-document-writeln}](#dom-document-writeln)[](#document.writeln()){.self-link} {#document.writeln()}

`document`{.variable}`.`[`writeln`](#dom-document-writeln){#dom-document-writeln-dev}`(...``text`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Document/writeln](https://developer.mozilla.org/en-US/docs/Web/API/Document/writeln "Writes a string of text followed by a newline character to a document.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari11+]{.safari .yes}[Chrome64+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera51+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer4+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android47+]{.opera_android .yes}
:::
::::
:::::

Adds the given string(s) to the
[`Document`{#document.writeln():document}](dom.html#document)\'s input
stream, followed by a newline character. If necessary, calls the
[`open()`{#document.writeln():dom-document-open}](#dom-document-open)
method implicitly first.

Throws an
[\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#document.writeln():invalidstateerror
x-internal="invalidstateerror"}
[`DOMException`{#document.writeln():domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
when invoked on [XML
documents](https://dom.spec.whatwg.org/#xml-document){#document.writeln():xml-documents
x-internal="xml-documents"}.

Throws an
[\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#document.writeln():invalidstateerror-2
x-internal="invalidstateerror"}
[`DOMException`{#document.writeln():domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if the parser is currently executing a [custom element
constructor](custom-elements.html#custom-element-constructor){#document.writeln():custom-element-constructor}.

This method performs no sanitization to remove potentially-dangerous
elements and attributes like
[`script`{#document.writeln():the-script-element}](scripting.html#the-script-element)
or [event handler content
attributes](webappapis.html#event-handler-content-attributes){#document.writeln():event-handler-content-attributes}.

The [`document.writeln(...``text`{.variable}`)`]{#dom-document-writeln
.dfn dfn-for="Document" dfn-type="method"} method steps are to run the
[document write
steps](#document-write-steps){#document.writeln():document-write-steps}
with
[this](https://webidl.spec.whatwg.org/#this){#document.writeln():this
x-internal="this"}, `text`{.variable}, true, and \"`Document writeln`\".

### [8.5]{.secno} DOM parsing and serialization APIs[](#dom-parsing-and-serialization){.self-link} {#dom-parsing-and-serialization}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[DOMParser](https://developer.mozilla.org/en-US/docs/Web/API/DOMParser "The DOMParser interface provides the ability to parse XML or HTML source code from a string into a DOM Document.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1.3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera8+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

``` idl
partial interface Element {
  [CEReactions] undefined setHTMLUnsafe((TrustedHTML or DOMString) html);
  DOMString getHTML(optional GetHTMLOptions options = {});

  [CEReactions] attribute (TrustedHTML or [LegacyNullToEmptyString] DOMString) innerHTML;
  [CEReactions] attribute (TrustedHTML or [LegacyNullToEmptyString] DOMString) outerHTML;
  [CEReactions] undefined insertAdjacentHTML(DOMString position, (TrustedHTML or DOMString) string);
};

partial interface ShadowRoot {
  [CEReactions] undefined setHTMLUnsafe((TrustedHTML or DOMString) html);
  DOMString getHTML(optional GetHTMLOptions options = {});

  [CEReactions] attribute (TrustedHTML or [LegacyNullToEmptyString] DOMString) innerHTML;
};

dictionary GetHTMLOptions {
  boolean serializableShadowRoots = false;
  sequence<ShadowRoot> shadowRoots = [];
};
```

#### [8.5.1]{.secno} The [`DOMParser`{#the-domparser-interface:domparser}](#domparser) interface[](#the-domparser-interface){.self-link}

The [`DOMParser`{#the-domparser-interface:domparser-2}](#domparser)
interface allows authors to create new
[`Document`{#the-domparser-interface:document}](dom.html#document)
objects by parsing strings, as either HTML or XML.

`parser`{.variable}` = new `[`DOMParser`](#dom-domparser-constructor){#dom-domparser-constructor-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[DOMParser/DOMParser](https://developer.mozilla.org/en-US/docs/Web/API/DOMParser/DOMParser "The DOMParser() constructor creates a new DOMParser object. This object can be used to parse the text of a document using the parseFromString() method.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1.3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera8+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

Constructs a new
[`DOMParser`{#the-domparser-interface:domparser-3}](#domparser) object.

`document`{.variable}` = ``parser`{.variable}`.`[`parseFromString`](#dom-domparser-parsefromstring){#dom-domparser-parsefromstring-dev}`(``string`{.variable}`, ``type`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[DOMParser/parseFromString](https://developer.mozilla.org/en-US/docs/Web/API/DOMParser/parseFromString "The parseFromString() method of the DOMParser interface parses a string containing either HTML or XML, returning an HTMLDocument or an XMLDocument.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1.3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera8+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

Parses `string`{.variable} using either the HTML or XML parser,
according to `type`{.variable}, and returns the resulting
[`Document`{#the-domparser-interface:document-2}](dom.html#document).
`type`{.variable} can be
\"[`text/html`{#the-domparser-interface:text/html}](iana.html#text/html)\"
(which will invoke the HTML parser), or any of
\"[`text/xml`{#the-domparser-interface:text/xml}](indices.html#text/xml)\",
\"[`application/xml`{#the-domparser-interface:application/xml}](indices.html#application/xml)\",
\"[`application/xhtml+xml`{#the-domparser-interface:application/xhtml+xml}](iana.html#application/xhtml+xml)\",
or
\"[`image/svg+xml`{#the-domparser-interface:image/svg+xml}](indices.html#image/svg+xml)\"
(which will invoke the XML parser).

For the XML parser, if `string`{.variable} cannot be parsed, then the
returned
[`Document`{#the-domparser-interface:document-3}](dom.html#document)
will contain elements describing the resulting error.

Note that
[`script`{#the-domparser-interface:the-script-element}](scripting.html#the-script-element)
elements are not evaluated during parsing, and the resulting document\'s
[encoding](https://dom.spec.whatwg.org/#concept-document-encoding){#the-domparser-interface:document's-character-encoding
x-internal="document's-character-encoding"} will always be
[UTF-8](https://encoding.spec.whatwg.org/#utf-8){#the-domparser-interface:utf-8
x-internal="utf-8"}. The document\'s
[URL](https://dom.spec.whatwg.org/#concept-document-url){#the-domparser-interface:the-document's-address
x-internal="the-document's-address"} will be inherited from
`parser`{.variable}\'s [relevant global
object](webappapis.html#concept-relevant-global){#the-domparser-interface:concept-relevant-global}.

Values other than the above for `type`{.variable} will cause a
[`TypeError`{#the-domparser-interface:typeerror}](https://tc39.es/ecma262/#sec-native-error-types-used-in-this-standard-typeerror){x-internal="typeerror"}
exception to be thrown.

The design of
[`DOMParser`{#the-domparser-interface:domparser-4}](#domparser), as a
class that needs to be constructed and then have its
[`parseFromString()`{#the-domparser-interface:dom-domparser-parsefromstring}](#dom-domparser-parsefromstring)
method called, is an unfortunate historical artifact. If we were
designing this functionality today it would be a standalone function.
For parsing HTML, the modern alternative is
[`Document.parseHTMLUnsafe()`{#the-domparser-interface:dom-parsehtmlunsafe}](#dom-parsehtmlunsafe).

This method performs no sanitization to remove potentially-dangerous
elements and attributes like
[`script`{#the-domparser-interface:the-script-element-2}](scripting.html#the-script-element)
or [event handler content
attributes](webappapis.html#event-handler-content-attributes){#the-domparser-interface:event-handler-content-attributes}.

``` idl
[Exposed=Window]
interface DOMParser {
  constructor();

  [NewObject] Document parseFromString((TrustedHTML or DOMString) string, DOMParserSupportedType type);
};

enum DOMParserSupportedType {
  "text/html",
  "text/xml",
  "application/xml",
  "application/xhtml+xml",
  "image/svg+xml"
};
```

The [`new DOMParser()`]{#dom-domparser-constructor .dfn} constructor
steps are to do nothing.

The
[`parseFromString(``string`{.variable}`, ``type`{.variable}`)`]{#dom-domparser-parsefromstring
.dfn dfn-for="DOMParser" dfn-type="method"} method steps are:

1.  Let `compliantString`{.variable} be the result of invoking the [Get
    Trusted Type compliant
    string](https://w3c.github.io/trusted-types/dist/spec/#get-trusted-type-compliant-string-algorithm){#the-domparser-interface:tt-getcompliantstring
    x-internal="tt-getcompliantstring"} algorithm with
    [`TrustedHTML`{#the-domparser-interface:tt-trustedhtml-2}](https://w3c.github.io/trusted-types/dist/spec/#trustedhtml){x-internal="tt-trustedhtml"},
    [this](https://webidl.spec.whatwg.org/#this){#the-domparser-interface:this
    x-internal="this"}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#the-domparser-interface:concept-relevant-global-2},
    `string`{.variable}, \"`DOMParser parseFromString`\", and
    \"`script`\".

2.  Let `document`{.variable} be a new
    [`Document`{#the-domparser-interface:document-5}](dom.html#document),
    whose [content
    type](https://dom.spec.whatwg.org/#concept-document-content-type){#the-domparser-interface:concept-document-content-type
    x-internal="concept-document-content-type"} is `type`{.variable} and
    [URL](https://dom.spec.whatwg.org/#concept-document-url){#the-domparser-interface:the-document's-address-2
    x-internal="the-document's-address"} is
    [this](https://webidl.spec.whatwg.org/#this){#the-domparser-interface:this-2
    x-internal="this"}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#the-domparser-interface:concept-relevant-global-3}\'s
    [associated
    `Document`](nav-history-apis.html#concept-document-window){#the-domparser-interface:concept-document-window}\'s
    [URL](https://dom.spec.whatwg.org/#concept-document-url){#the-domparser-interface:the-document's-address-3
    x-internal="the-document's-address"}.

    The document\'s
    [encoding](https://dom.spec.whatwg.org/#concept-document-encoding){#the-domparser-interface:document's-character-encoding-2
    x-internal="document's-character-encoding"} will be left as its
    default, of
    [UTF-8](https://encoding.spec.whatwg.org/#utf-8){#the-domparser-interface:utf-8-2
    x-internal="utf-8"}. In particular, any XML declarations or
    [`meta`{#the-domparser-interface:the-meta-element}](semantics.html#the-meta-element)
    elements found while parsing `compliantString`{.variable} will have
    no effect.

3.  Switch on `type`{.variable}:

    \"[`text/html`]{#dom-domparsersupportedtype-texthtml .dfn dfn-for="DOMParserSupportedType" dfn-type="enum-value"}\"

    :   1.  [Parse HTML from a
            string](#parse-html-from-a-string){#the-domparser-interface:parse-html-from-a-string}
            given `document`{.variable} and
            `compliantString`{.variable}.

        Since `document`{.variable} does not have a [browsing
        context](document-sequences.html#concept-document-bc){#the-domparser-interface:concept-document-bc},
        [scripting is
        disabled](webappapis.html#concept-n-script){#the-domparser-interface:concept-n-script}.

    [Otherwise]{#dom-domparsersupportedtype-otherwise .dfn}

    :   1.  Create an [XML
            parser](xhtml.html#xml-parser){#the-domparser-interface:xml-parser}
            `parser`{.variable}, associated with `document`{.variable},
            and with [XML scripting support
            disabled](xhtml.html#xml-scripting-support-disabled){#the-domparser-interface:xml-scripting-support-disabled}.

        2.  Parse `compliantString`{.variable} using
            `parser`{.variable}.

        3.  If the previous step resulted in an XML well-formedness or
            XML namespace well-formedness error, then:

            1.  [Assert](https://infra.spec.whatwg.org/#assert){#the-domparser-interface:assert
                x-internal="assert"}: `document`{.variable} has no child
                nodes.

            2.  Let `root`{.variable} be the result of [creating an
                element](https://dom.spec.whatwg.org/#concept-create-element){#the-domparser-interface:create-an-element
                x-internal="create-an-element"} given
                `document`{.variable}, \"`parsererror`\", and
                \"`http://www.mozilla.org/newlayout/xml/parsererror.xml`\".

            3.  Optionally, add attributes or children to
                `root`{.variable} to describe the nature of the parsing
                error.

            4.  [Append](https://dom.spec.whatwg.org/#concept-node-append){#the-domparser-interface:concept-node-append
                x-internal="concept-node-append"} `root`{.variable} to
                `document`{.variable}.

4.  Return `document`{.variable}.

To [parse HTML from a string]{#parse-html-from-a-string .dfn}, given a
[`Document`{#the-domparser-interface:document-6}](dom.html#document)
`document`{.variable} and a
[string](https://infra.spec.whatwg.org/#string){#the-domparser-interface:string
x-internal="string"} `html`{.variable}:

1.  Set `document`{.variable}\'s
    [type](https://dom.spec.whatwg.org/#concept-document-type){#the-domparser-interface:concept-document-type
    x-internal="concept-document-type"} to \"`html`\".

2.  Create an [HTML
    parser](parsing.html#html-parser){#the-domparser-interface:html-parser}
    `parser`{.variable}, associated with `document`{.variable}.

3.  Place `html`{.variable} into the [input
    stream](parsing.html#input-stream){#the-domparser-interface:input-stream}
    for `parser`{.variable}. The encoding
    [confidence](parsing.html#concept-encoding-confidence){#the-domparser-interface:concept-encoding-confidence}
    is *irrelevant*.

4.  Start `parser`{.variable} and let it run until it has consumed all
    the characters just inserted into the input stream.

    This might mutate the document\'s
    [mode](https://dom.spec.whatwg.org/#concept-document-mode){#the-domparser-interface:concept-document-mode
    x-internal="concept-document-mode"}.

#### [8.5.2]{.secno} Unsafe HTML parsing methods[](#unsafe-html-parsing-methods){.self-link}

`element`{.variable}`.`[`setHTMLUnsafe`](#dom-element-sethtmlunsafe){#dom-element-sethtmlunsafe-dev}`(``html`{.variable}`)`

:   Parses `html`{.variable} using the HTML parser, and replaces the
    children of `element`{.variable} with the result.
    `element`{.variable} provides context for the HTML parser.

`shadowRoot`{.variable}`.`[`setHTMLUnsafe`](#dom-shadowroot-sethtmlunsafe){#dom-shadowroot-sethtmlunsafe-dev}`(``html`{.variable}`)`

:   Parses `html`{.variable} using the HTML parser, and replaces the
    children of `shadowRoot`{.variable} with the result.
    `shadowRoot`{.variable}\'s
    [host](https://dom.spec.whatwg.org/#concept-documentfragment-host){#unsafe-html-parsing-methods:concept-documentfragment-host
    x-internal="concept-documentfragment-host"} provides context for the
    HTML parser.

`doc`{.variable}` = Document.`[`parseHTMLUnsafe`](#dom-parsehtmlunsafe){#unsafe-html-parsing-methods:dom-parsehtmlunsafe}`(``html`{.variable}`)`

:   Parses `html`{.variable} using the HTML parser, and returns the
    resulting
    [`Document`{#unsafe-html-parsing-methods:document}](dom.html#document).

    Note that
    [`script`{#unsafe-html-parsing-methods:the-script-element}](scripting.html#the-script-element)
    elements are not evaluated during parsing, and the resulting
    document\'s
    [encoding](https://dom.spec.whatwg.org/#concept-document-encoding){#unsafe-html-parsing-methods:document's-character-encoding
    x-internal="document's-character-encoding"} will always be
    [UTF-8](https://encoding.spec.whatwg.org/#utf-8){#unsafe-html-parsing-methods:utf-8
    x-internal="utf-8"}. The document\'s
    [URL](https://dom.spec.whatwg.org/#concept-document-url){#unsafe-html-parsing-methods:the-document's-address
    x-internal="the-document's-address"} will be
    [`about:blank`{#unsafe-html-parsing-methods:about:blank}](infrastructure.html#about:blank).

These methods perform no sanitization to remove potentially-dangerous
elements and attributes like
[`script`{#unsafe-html-parsing-methods:the-script-element-2}](scripting.html#the-script-element)
or [event handler content
attributes](webappapis.html#event-handler-content-attributes){#unsafe-html-parsing-methods:event-handler-content-attributes}.

[`Element`{#unsafe-html-parsing-methods:element}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"}\'s
[`setHTMLUnsafe(``html`{.variable}`)`]{#dom-element-sethtmlunsafe .dfn
dfn-for="Element" dfn-type="method"} method steps are:

1.  Let `compliantHTML`{.variable} be the result of invoking the [Get
    Trusted Type compliant
    string](https://w3c.github.io/trusted-types/dist/spec/#get-trusted-type-compliant-string-algorithm){#unsafe-html-parsing-methods:tt-getcompliantstring
    x-internal="tt-getcompliantstring"} algorithm with
    [`TrustedHTML`{#unsafe-html-parsing-methods:tt-trustedhtml}](https://w3c.github.io/trusted-types/dist/spec/#trustedhtml){x-internal="tt-trustedhtml"},
    [this](https://webidl.spec.whatwg.org/#this){#unsafe-html-parsing-methods:this
    x-internal="this"}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#unsafe-html-parsing-methods:concept-relevant-global},
    `html`{.variable}, \"`Element setHTMLUnsafe`\", and \"`script`\".

2.  Let `target`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#unsafe-html-parsing-methods:this-2
    x-internal="this"}\'s [template
    contents](scripting.html#template-contents){#unsafe-html-parsing-methods:template-contents}
    if
    [this](https://webidl.spec.whatwg.org/#this){#unsafe-html-parsing-methods:this-3
    x-internal="this"} is a
    [`template`{#unsafe-html-parsing-methods:the-template-element}](scripting.html#the-template-element)
    element; otherwise
    [this](https://webidl.spec.whatwg.org/#this){#unsafe-html-parsing-methods:this-4
    x-internal="this"}.

3.  [Unsafely set
    HTML](#unsafely-set-html){#unsafe-html-parsing-methods:unsafely-set-html}
    given `target`{.variable},
    [this](https://webidl.spec.whatwg.org/#this){#unsafe-html-parsing-methods:this-5
    x-internal="this"}, and `compliantHTML`{.variable}.

[`ShadowRoot`{#unsafe-html-parsing-methods:shadowroot}](https://dom.spec.whatwg.org/#interface-shadowroot){x-internal="shadowroot"}\'s
[`setHTMLUnsafe(``html`{.variable}`)`]{#dom-shadowroot-sethtmlunsafe
.dfn dfn-for="ShadowRoot" dfn-type="method"} method steps are:

1.  Let `compliantHTML`{.variable} be the result of invoking the [Get
    Trusted Type compliant
    string](https://w3c.github.io/trusted-types/dist/spec/#get-trusted-type-compliant-string-algorithm){#unsafe-html-parsing-methods:tt-getcompliantstring-2
    x-internal="tt-getcompliantstring"} algorithm with
    [`TrustedHTML`{#unsafe-html-parsing-methods:tt-trustedhtml-2}](https://w3c.github.io/trusted-types/dist/spec/#trustedhtml){x-internal="tt-trustedhtml"},
    [this](https://webidl.spec.whatwg.org/#this){#unsafe-html-parsing-methods:this-6
    x-internal="this"}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#unsafe-html-parsing-methods:concept-relevant-global-2},
    `html`{.variable}, \"`ShadowRoot setHTMLUnsafe`\", and \"`script`\".

2.  [Unsafely set
    HTML](#unsafely-set-html){#unsafe-html-parsing-methods:unsafely-set-html-2}
    given
    [this](https://webidl.spec.whatwg.org/#this){#unsafe-html-parsing-methods:this-7
    x-internal="this"},
    [this](https://webidl.spec.whatwg.org/#this){#unsafe-html-parsing-methods:this-8
    x-internal="this"}\'s [shadow
    host](https://dom.spec.whatwg.org/#concept-documentfragment-host){#unsafe-html-parsing-methods:concept-documentfragment-host-2
    x-internal="concept-documentfragment-host"}, and
    `compliantHTML`{.variable}.

To [unsafely set HTML]{#unsafely-set-html .dfn}, given an
[`Element`{#unsafe-html-parsing-methods:element-2}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"}
or
[`DocumentFragment`{#unsafe-html-parsing-methods:documentfragment}](https://dom.spec.whatwg.org/#interface-documentfragment){x-internal="documentfragment"}
`target`{.variable}, an
[`Element`{#unsafe-html-parsing-methods:element-3}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"}
`contextElement`{.variable}, and a
[string](https://infra.spec.whatwg.org/#string){#unsafe-html-parsing-methods:string
x-internal="string"} `html`{.variable}:

1.  Let `newChildren`{.variable} be the result of the [HTML fragment
    parsing
    algorithm](parsing.html#html-fragment-parsing-algorithm){#unsafe-html-parsing-methods:html-fragment-parsing-algorithm}
    given `contextElement`{.variable}, `html`{.variable}, and true.

2.  Let `fragment`{.variable} be a new
    [`DocumentFragment`{#unsafe-html-parsing-methods:documentfragment-2}](https://dom.spec.whatwg.org/#interface-documentfragment){x-internal="documentfragment"}
    whose [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#unsafe-html-parsing-methods:node-document
    x-internal="node-document"} is `contextElement`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#unsafe-html-parsing-methods:node-document-2
    x-internal="node-document"}.

3.  For each `node`{.variable} in `newChildren`{.variable},
    [append](https://dom.spec.whatwg.org/#concept-node-append){#unsafe-html-parsing-methods:concept-node-append
    x-internal="concept-node-append"} `node`{.variable} to
    `fragment`{.variable}.

4.  [Replace
    all](https://dom.spec.whatwg.org/#concept-node-replace-all){#unsafe-html-parsing-methods:concept-node-replace-all
    x-internal="concept-node-replace-all"} with `fragment`{.variable}
    within `target`{.variable}.

------------------------------------------------------------------------

The static [`parseHTMLUnsafe(``html`{.variable}`)`]{#dom-parsehtmlunsafe
.dfn dfn-for="Document" dfn-type="method"} method steps are:

1.  Let `compliantHTML`{.variable} be the result of invoking the [Get
    Trusted Type compliant
    string](https://w3c.github.io/trusted-types/dist/spec/#get-trusted-type-compliant-string-algorithm){#unsafe-html-parsing-methods:tt-getcompliantstring-3
    x-internal="tt-getcompliantstring"} algorithm with
    [`TrustedHTML`{#unsafe-html-parsing-methods:tt-trustedhtml-3}](https://w3c.github.io/trusted-types/dist/spec/#trustedhtml){x-internal="tt-trustedhtml"},
    [this](https://webidl.spec.whatwg.org/#this){#unsafe-html-parsing-methods:this-9
    x-internal="this"}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#unsafe-html-parsing-methods:concept-relevant-global-3},
    `html`{.variable}, \"`Document parseHTMLUnsafe`\", and \"`script`\".

2.  Let `document`{.variable} be a new
    [`Document`{#unsafe-html-parsing-methods:document-2}](dom.html#document),
    whose [content
    type](https://dom.spec.whatwg.org/#concept-document-content-type){#unsafe-html-parsing-methods:concept-document-content-type
    x-internal="concept-document-content-type"} is \"`text/html`\".

    Since `document`{.variable} does not have a [browsing
    context](document-sequences.html#concept-document-bc){#unsafe-html-parsing-methods:concept-document-bc},
    [scripting is
    disabled](webappapis.html#concept-n-script){#unsafe-html-parsing-methods:concept-n-script}.

3.  Set `document`{.variable}\'s [allow declarative shadow
    roots](https://dom.spec.whatwg.org/#concept-document-allow-declarative-shadow-roots){#unsafe-html-parsing-methods:concept-document-allow-declarative-shadow-roots
    x-internal="concept-document-allow-declarative-shadow-roots"} to
    true.

4.  [Parse HTML from a
    string](#parse-html-from-a-string){#unsafe-html-parsing-methods:parse-html-from-a-string}
    given `document`{.variable} and `compliantHTML`{.variable}.

5.  Return `document`{.variable}.

#### [8.5.3]{.secno} HTML serialization methods[](#html-serialization-methods){.self-link}

`html`{.variable}` = ``element`{.variable}`.`[`getHTML`](#dom-element-gethtml){#dom-element-gethtml-dev}`({ `[`serializableShadowRoots`](#dom-gethtmloptions-serializableshadowroots){#html-serialization-methods:dom-gethtmloptions-serializableshadowroots}`, `[`shadowRoots`](#dom-gethtmloptions-shadowroots){#html-serialization-methods:dom-gethtmloptions-shadowroots}` })`

:   Returns the result of serializing `element`{.variable} to HTML.
    [Shadow
    roots](https://dom.spec.whatwg.org/#concept-shadow-root){#html-serialization-methods:shadow-root
    x-internal="shadow-root"} within `element`{.variable} are serialized
    according to the provided options:

    - If
      [`serializableShadowRoots`{#html-serialization-methods:dom-gethtmloptions-serializableshadowroots-2}](#dom-gethtmloptions-serializableshadowroots)
      is true, then all shadow roots marked as
      [serializable](https://dom.spec.whatwg.org/#shadowroot-serializable){#html-serialization-methods:shadow-serializable
      x-internal="shadow-serializable"} are serialized.

    - If the
      [`shadowRoots`{#html-serialization-methods:dom-gethtmloptions-shadowroots-2}](#dom-gethtmloptions-shadowroots)
      array is provided, then all shadow roots specified in the array
      are serialized, regardless of whether or not they are marked as
      serializable.

    If neither option is provided, then no shadow roots are serialized.

`html`{.variable}` = ``shadowRoot`{.variable}`.`[`getHTML`](#dom-shadowroot-gethtml){#dom-shadowroot-gethtml-dev}`({ `[`serializableShadowRoots`](#dom-gethtmloptions-serializableshadowroots){#html-serialization-methods:dom-gethtmloptions-serializableshadowroots-3}`, `[`shadowRoots`](#dom-gethtmloptions-shadowroots){#html-serialization-methods:dom-gethtmloptions-shadowroots-3}` })`

:   Returns the result of serializing `shadowRoot`{.variable} to HTML,
    using its [shadow
    host](https://dom.spec.whatwg.org/#concept-documentfragment-host){#html-serialization-methods:concept-documentfragment-host
    x-internal="concept-documentfragment-host"} as the context element.
    [Shadow
    roots](https://dom.spec.whatwg.org/#concept-shadow-root){#html-serialization-methods:shadow-root-2
    x-internal="shadow-root"} within `shadowRoot`{.variable} are
    serialized according to the provided options, as above.

[`Element`{#html-serialization-methods:element}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"}\'s
[`getHTML(``options`{.variable}`)`]{#dom-element-gethtml .dfn
dfn-for="Element" dfn-type="method"} method steps are to return the
result of [HTML fragment serialization
algorithm](parsing.html#html-fragment-serialisation-algorithm){#html-serialization-methods:html-fragment-serialisation-algorithm}
with
[this](https://webidl.spec.whatwg.org/#this){#html-serialization-methods:this
x-internal="this"},
`options`{.variable}\[\"[`serializableShadowRoots`{#html-serialization-methods:dom-gethtmloptions-serializableshadowroots-4}](#dom-gethtmloptions-serializableshadowroots)\"\],
and
`options`{.variable}\[\"[`shadowRoots`{#html-serialization-methods:dom-gethtmloptions-shadowroots-4}](#dom-gethtmloptions-shadowroots)\"\].

[`ShadowRoot`{#html-serialization-methods:shadowroot}](https://dom.spec.whatwg.org/#interface-shadowroot){x-internal="shadowroot"}\'s
[`getHTML(``options`{.variable}`)`]{#dom-shadowroot-gethtml .dfn
dfn-for="ShadowRoot" dfn-type="method"} method steps are to return the
result of [HTML fragment serialization
algorithm](parsing.html#html-fragment-serialisation-algorithm){#html-serialization-methods:html-fragment-serialisation-algorithm-2}
with
[this](https://webidl.spec.whatwg.org/#this){#html-serialization-methods:this-2
x-internal="this"},
`options`{.variable}\[\"[`serializableShadowRoots`{#html-serialization-methods:dom-gethtmloptions-serializableshadowroots-5}](#dom-gethtmloptions-serializableshadowroots)\"\],
and
`options`{.variable}\[\"[`shadowRoots`{#html-serialization-methods:dom-gethtmloptions-shadowroots-5}](#dom-gethtmloptions-shadowroots)\"\].

#### [8.5.4]{.secno} The [`innerHTML`{#the-innerhtml-property:dom-element-innerhtml}](#dom-element-innerhtml) property[](#the-innerhtml-property){.self-link}

The
[`innerHTML`{#the-innerhtml-property:dom-element-innerhtml-2}](#dom-element-innerhtml)
property has a number of outstanding issues in the DOM Parsing and
Serialization [issue
tracker](https://github.com/w3c/DOM-Parsing/issues), documenting various
problems with its specification.

`element`{.variable}`.`[`innerHTML`](#dom-element-innerhtml){#dom-element-innerhtml-dev}

:   Returns a fragment of HTML or XML that represents the element\'s
    contents.

    In the case of an XML document, throws an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-innerhtml-property:invalidstateerror
    x-internal="invalidstateerror"}
    [`DOMException`{#the-innerhtml-property:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
    if the element cannot be serialized to XML.

`element`{.variable}`.`[`innerHTML`](#dom-element-innerhtml){#the-innerhtml-property:dom-element-innerhtml-3}` = ``value`{.variable}

:   Replaces the contents of the element with nodes parsed from the
    given string.

    In the case of an XML document, throws a
    [\"`SyntaxError`\"](https://webidl.spec.whatwg.org/#syntaxerror){#the-innerhtml-property:syntaxerror
    x-internal="syntaxerror"}
    [`DOMException`{#the-innerhtml-property:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
    if the given string is not well-formed.

`shadowRoot`{.variable}`.`[`innerHTML`](#dom-shadowroot-innerhtml){#dom-shadowroot-innerhtml-dev}

:   Returns a fragment of HTML that represents the shadow roots\'s
    contents.

`shadowRoot`{.variable}`.`[`innerHTML`](#dom-shadowroot-innerhtml){#the-innerhtml-property:dom-shadowroot-innerhtml}` = ``value`{.variable}

:   Replaces the contents of the shadow root with nodes parsed from the
    given string.

These properties\' setters perform no sanitization to remove
potentially-dangerous elements and attributes like
[`script`{#the-innerhtml-property:the-script-element}](scripting.html#the-script-element)
or [event handler content
attributes](webappapis.html#event-handler-content-attributes){#the-innerhtml-property:event-handler-content-attributes}.

The [fragment serializing algorithm
steps]{#fragment-serializing-algorithm-steps .dfn export=""}, given an
[`Element`{#the-innerhtml-property:element}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"},
[`Document`{#the-innerhtml-property:document}](dom.html#document), or
[`DocumentFragment`{#the-innerhtml-property:documentfragment}](https://dom.spec.whatwg.org/#interface-documentfragment){x-internal="documentfragment"}
`node`{.variable} and a boolean `require well-formed`{.variable}, are:

1.  Let `context document`{.variable} be `node`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-innerhtml-property:node-document
    x-internal="node-document"}.

2.  If `context document`{.variable} is an [HTML
    document](https://dom.spec.whatwg.org/#html-document){#the-innerhtml-property:html-documents
    x-internal="html-documents"}, return the result of [HTML fragment
    serialization
    algorithm](parsing.html#html-fragment-serialisation-algorithm){#the-innerhtml-property:html-fragment-serialisation-algorithm}
    with `node`{.variable}, false, and « ».

3.  Return the [XML
    serialization](https://w3c.github.io/DOM-Parsing/#dfn-xml-serialization){#the-innerhtml-property:xml-serialization
    x-internal="xml-serialization"} of `node`{.variable} given
    `require well-formed`{.variable}.

The [fragment parsing algorithm steps]{#fragment-parsing-algorithm-steps
.dfn export=""}, given an
[`Element`{#the-innerhtml-property:element-2}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"}
`context`{.variable} and a string `markup`{.variable}, are:

1.  Let `algorithm`{.variable} be the [HTML fragment parsing
    algorithm](parsing.html#html-fragment-parsing-algorithm){#the-innerhtml-property:html-fragment-parsing-algorithm}.

2.  If `context`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-innerhtml-property:node-document-2
    x-internal="node-document"} is an [XML
    document](https://dom.spec.whatwg.org/#xml-document){#the-innerhtml-property:xml-documents
    x-internal="xml-documents"}, then set `algorithm`{.variable} to the
    [XML fragment parsing
    algorithm](xhtml.html#xml-fragment-parsing-algorithm){#the-innerhtml-property:xml-fragment-parsing-algorithm}.

3.  Let `newChildren`{.variable} be the result of invoking
    `algorithm`{.variable} given `context`{.variable} and
    `markup`{.variable}.

4.  Let `fragment`{.variable} be a new
    [`DocumentFragment`{#the-innerhtml-property:documentfragment-2}](https://dom.spec.whatwg.org/#interface-documentfragment){x-internal="documentfragment"}
    whose [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-innerhtml-property:node-document-3
    x-internal="node-document"} is `context`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-innerhtml-property:node-document-4
    x-internal="node-document"}.

5.  For each `node`{.variable} of `newChildren`{.variable}, in [tree
    order](https://dom.spec.whatwg.org/#concept-tree-order){#the-innerhtml-property:tree-order
    x-internal="tree-order"}:
    [append](https://dom.spec.whatwg.org/#concept-node-append){#the-innerhtml-property:concept-node-append
    x-internal="concept-node-append"} `node`{.variable} to
    `fragment`{.variable}.

    This ensures the [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-innerhtml-property:node-document-5
    x-internal="node-document"} for the new
    [nodes](https://dom.spec.whatwg.org/#interface-node){#the-innerhtml-property:node
    x-internal="node"} is correct.

6.  Return `fragment`{.variable}.

[`Element`{#the-innerhtml-property:element-3}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"}\'s
[`innerHTML`]{#dom-element-innerhtml .dfn dfn-for="Element"
dfn-type="attribute"} getter steps are to return the result of running
[fragment serializing algorithm
steps](#fragment-serializing-algorithm-steps){#the-innerhtml-property:fragment-serializing-algorithm-steps}
with
[this](https://webidl.spec.whatwg.org/#this){#the-innerhtml-property:this
x-internal="this"} and true.

[`ShadowRoot`{#the-innerhtml-property:shadowroot}](https://dom.spec.whatwg.org/#interface-shadowroot){x-internal="shadowroot"}\'s
[`innerHTML`]{#dom-shadowroot-innerhtml .dfn dfn-for="ShadowRoot"
dfn-type="attribute"} getter steps are to return the result of running
[fragment serializing algorithm
steps](#fragment-serializing-algorithm-steps){#the-innerhtml-property:fragment-serializing-algorithm-steps-2}
with
[this](https://webidl.spec.whatwg.org/#this){#the-innerhtml-property:this-2
x-internal="this"} and true.

[`Element`{#the-innerhtml-property:element-4}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"}\'s
[`innerHTML`{#the-innerhtml-property:dom-element-innerhtml-4}](#dom-element-innerhtml)
setter steps are:

1.  Let `compliantString`{.variable} be the result of invoking the [Get
    Trusted Type compliant
    string](https://w3c.github.io/trusted-types/dist/spec/#get-trusted-type-compliant-string-algorithm){#the-innerhtml-property:tt-getcompliantstring
    x-internal="tt-getcompliantstring"} algorithm with
    [`TrustedHTML`{#the-innerhtml-property:tt-trustedhtml}](https://w3c.github.io/trusted-types/dist/spec/#trustedhtml){x-internal="tt-trustedhtml"},
    [this](https://webidl.spec.whatwg.org/#this){#the-innerhtml-property:this-3
    x-internal="this"}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#the-innerhtml-property:concept-relevant-global},
    the given value, \"`Element innerHTML`\", and \"`script`\".

2.  Let `context`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#the-innerhtml-property:this-4
    x-internal="this"}.

3.  Let `fragment`{.variable} be the result of invoking the [fragment
    parsing algorithm
    steps](#fragment-parsing-algorithm-steps){#the-innerhtml-property:fragment-parsing-algorithm-steps}
    with `context`{.variable} and `compliantString`{.variable}.

4.  If `context`{.variable} is a
    [`template`{#the-innerhtml-property:the-template-element}](scripting.html#the-template-element)
    element, then set `context`{.variable} to the
    [`template`{#the-innerhtml-property:the-template-element-2}](scripting.html#the-template-element)
    element\'s [template
    contents](scripting.html#template-contents){#the-innerhtml-property:template-contents}
    (a
    [`DocumentFragment`{#the-innerhtml-property:documentfragment-3}](https://dom.spec.whatwg.org/#interface-documentfragment){x-internal="documentfragment"}).

    Setting
    [`innerHTML`{#the-innerhtml-property:dom-element-innerhtml-5}](#dom-element-innerhtml)
    on a
    [`template`{#the-innerhtml-property:the-template-element-3}](scripting.html#the-template-element)
    element will replace all the nodes in its [template
    contents](scripting.html#template-contents){#the-innerhtml-property:template-contents-2}
    rather than its
    [children](https://dom.spec.whatwg.org/#concept-tree-child){#the-innerhtml-property:concept-tree-child
    x-internal="concept-tree-child"}.

5.  [Replace
    all](https://dom.spec.whatwg.org/#concept-node-replace-all){#the-innerhtml-property:concept-node-replace-all
    x-internal="concept-node-replace-all"} with `fragment`{.variable}
    within `context`{.variable}.

[`ShadowRoot`{#the-innerhtml-property:shadowroot-2}](https://dom.spec.whatwg.org/#interface-shadowroot){x-internal="shadowroot"}\'s
[`innerHTML`{#the-innerhtml-property:dom-shadowroot-innerhtml-2}](#dom-shadowroot-innerhtml)
setter steps are:

1.  Let `compliantString`{.variable} be the result of invoking the [Get
    Trusted Type compliant
    string](https://w3c.github.io/trusted-types/dist/spec/#get-trusted-type-compliant-string-algorithm){#the-innerhtml-property:tt-getcompliantstring-2
    x-internal="tt-getcompliantstring"} algorithm with
    [`TrustedHTML`{#the-innerhtml-property:tt-trustedhtml-2}](https://w3c.github.io/trusted-types/dist/spec/#trustedhtml){x-internal="tt-trustedhtml"},
    [this](https://webidl.spec.whatwg.org/#this){#the-innerhtml-property:this-5
    x-internal="this"}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#the-innerhtml-property:concept-relevant-global-2},
    the given value, \"`ShadowRoot innerHTML`\", and \"`script`\".

2.  Let `context`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#the-innerhtml-property:this-6
    x-internal="this"}\'s
    [host](https://dom.spec.whatwg.org/#concept-documentfragment-host){#the-innerhtml-property:concept-documentfragment-host
    x-internal="concept-documentfragment-host"}.

3.  Let `fragment`{.variable} be the result of invoking the [fragment
    parsing algorithm
    steps](#fragment-parsing-algorithm-steps){#the-innerhtml-property:fragment-parsing-algorithm-steps-2}
    with `context`{.variable} and `compliantString`{.variable}.

4.  [Replace
    all](https://dom.spec.whatwg.org/#concept-node-replace-all){#the-innerhtml-property:concept-node-replace-all-2
    x-internal="concept-node-replace-all"} with `fragment`{.variable}
    within
    [this](https://webidl.spec.whatwg.org/#this){#the-innerhtml-property:this-7
    x-internal="this"}.

#### [8.5.5]{.secno} The [`outerHTML`{#the-outerhtml-property:dom-element-outerhtml}](#dom-element-outerhtml) property[](#the-outerhtml-property){.self-link}

The
[`outerHTML`{#the-outerhtml-property:dom-element-outerhtml-2}](#dom-element-outerhtml)
property has a number of outstanding issues in the DOM Parsing and
Serialization [issue
tracker](https://github.com/w3c/DOM-Parsing/issues), documenting various
problems with its specification.

`element`{.variable}`.`[`outerHTML`](#dom-element-outerhtml){#dom-element-outerhtml-dev}

:   Returns a fragment of HTML or XML that represents the element and
    its contents.

    In the case of an XML document, throws an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-outerhtml-property:invalidstateerror
    x-internal="invalidstateerror"}
    [`DOMException`{#the-outerhtml-property:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
    if the element cannot be serialized to XML.

`element`{.variable}`.`[`outerHTML`](#dom-element-outerhtml){#the-outerhtml-property:dom-element-outerhtml-3}` = ``value`{.variable}

:   Replaces the element with nodes parsed from the given string.

    In the case of an XML document, throws a
    [\"`SyntaxError`\"](https://webidl.spec.whatwg.org/#syntaxerror){#the-outerhtml-property:syntaxerror
    x-internal="syntaxerror"}
    [`DOMException`{#the-outerhtml-property:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
    if the given string is not well-formed.

    Throws a
    [\"`NoModificationAllowedError`\"](https://webidl.spec.whatwg.org/#nomodificationallowederror){#the-outerhtml-property:nomodificationallowederror
    x-internal="nomodificationallowederror"}
    [`DOMException`{#the-outerhtml-property:domexception-3}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
    if the parent of the element is a
    [`Document`](dom.html#document){#the-outerhtml-property:document}.

This property\'s setter performs no sanitization to remove
potentially-dangerous elements and attributes like
[`script`{#the-outerhtml-property:the-script-element}](scripting.html#the-script-element)
or [event handler content
attributes](webappapis.html#event-handler-content-attributes){#the-outerhtml-property:event-handler-content-attributes}.

[`Element`{#the-outerhtml-property:element}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"}\'s
[`outerHTML`]{#dom-element-outerhtml .dfn dfn-for="Element"
dfn-type="attribute"} getter steps are:

1.  Let `element`{.variable} be a fictional node whose only child is
    [this](https://webidl.spec.whatwg.org/#this){#the-outerhtml-property:this
    x-internal="this"}.

2.  Return the result of running [fragment serializing algorithm
    steps](#fragment-serializing-algorithm-steps){#the-outerhtml-property:fragment-serializing-algorithm-steps}
    with `element`{.variable} and true.

[`Element`{#the-outerhtml-property:element-2}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"}\'s
[`outerHTML`{#the-outerhtml-property:dom-element-outerhtml-4}](#dom-element-outerhtml)
setter steps are:

1.  Let `compliantString`{.variable} be the result of invoking the [Get
    Trusted Type compliant
    string](https://w3c.github.io/trusted-types/dist/spec/#get-trusted-type-compliant-string-algorithm){#the-outerhtml-property:tt-getcompliantstring
    x-internal="tt-getcompliantstring"} algorithm with
    [`TrustedHTML`{#the-outerhtml-property:tt-trustedhtml}](https://w3c.github.io/trusted-types/dist/spec/#trustedhtml){x-internal="tt-trustedhtml"},
    [this](https://webidl.spec.whatwg.org/#this){#the-outerhtml-property:this-2
    x-internal="this"}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#the-outerhtml-property:concept-relevant-global},
    the given value, \"`Element outerHTML`\", and \"`script`\".

2.  Let `parent`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#the-outerhtml-property:this-3
    x-internal="this"}\'s
    [parent](https://dom.spec.whatwg.org/#concept-tree-parent){#the-outerhtml-property:parent
    x-internal="parent"}.

3.  If `parent`{.variable} is null, return. There would be no way to
    obtain a reference to the nodes created even if the remaining steps
    were run.

4.  If `parent`{.variable} is a
    [`Document`{#the-outerhtml-property:document-2}](dom.html#document),
    throw a
    [\"`NoModificationAllowedError`\"](https://webidl.spec.whatwg.org/#nomodificationallowederror){#the-outerhtml-property:nomodificationallowederror-2
    x-internal="nomodificationallowederror"}
    [`DOMException`{#the-outerhtml-property:domexception-4}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

5.  If `parent`{.variable} is a
    [`DocumentFragment`{#the-outerhtml-property:documentfragment}](https://dom.spec.whatwg.org/#interface-documentfragment){x-internal="documentfragment"},
    set `parent`{.variable} to the result of [creating an
    element](https://dom.spec.whatwg.org/#concept-create-element){#the-outerhtml-property:create-an-element
    x-internal="create-an-element"} given
    [this](https://webidl.spec.whatwg.org/#this){#the-outerhtml-property:this-4
    x-internal="this"}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-outerhtml-property:node-document
    x-internal="node-document"}, \"`body`\", and the [HTML
    namespace](https://infra.spec.whatwg.org/#html-namespace){#the-outerhtml-property:html-namespace-2
    x-internal="html-namespace-2"}.

6.  Let `fragment`{.variable} be the result of invoking the [fragment
    parsing algorithm
    steps](#fragment-parsing-algorithm-steps){#the-outerhtml-property:fragment-parsing-algorithm-steps}
    given `parent`{.variable} and `compliantString`{.variable}.

7.  [Replace](https://dom.spec.whatwg.org/#concept-node-replace){#the-outerhtml-property:concept-node-replace
    x-internal="concept-node-replace"}
    [this](https://webidl.spec.whatwg.org/#this){#the-outerhtml-property:this-5
    x-internal="this"} with `fragment`{.variable} within
    [this](https://webidl.spec.whatwg.org/#this){#the-outerhtml-property:this-6
    x-internal="this"}\'s
    [parent](https://dom.spec.whatwg.org/#concept-tree-parent){#the-outerhtml-property:parent-2
    x-internal="parent"}.

#### [8.5.6]{.secno} The [`insertAdjacentHTML()`{#the-insertadjacenthtml()-method:dom-element-insertadjacenthtml}](#dom-element-insertadjacenthtml) method[](#the-insertadjacenthtml()-method){.self-link} {#the-insertadjacenthtml()-method}

The
[`insertAdjacentHTML()`{#the-insertadjacenthtml()-method:dom-element-insertadjacenthtml-2}](#dom-element-insertadjacenthtml)
method has a number of outstanding issues in the DOM Parsing and
Serialization [issue
tracker](https://github.com/w3c/DOM-Parsing/issues), documenting various
problems with its specification.

`element`{.variable}`.`[`insertAdjacentHTML`](#dom-element-insertadjacenthtml){#dom-element-insertadjacenthtml-dev}`(``position`{.variable}`, ``string`{.variable}`)`

:   Parses `string`{.variable} as HTML or XML and inserts the resulting
    nodes into the tree in the position given by the
    `position`{.variable} argument, as follows:

    \"`beforebegin`\"
    :   Before the element itself (i.e., after `element`{.variable}\'s
        previous sibling)

    \"`afterbegin`\"
    :   Just inside the element, before its first child.

    \"`beforeend`\"
    :   Just inside the element, after its last child.

    \"`afterend`\"
    :   After the element itself (i.e., before `element`{.variable}\'s
        next sibling)

    Throws a
    [\"`SyntaxError`\"](https://webidl.spec.whatwg.org/#syntaxerror){#the-insertadjacenthtml()-method:syntaxerror
    x-internal="syntaxerror"}
    [`DOMException`{#the-insertadjacenthtml()-method:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
    if the arguments have invalid values (e.g., in the case of an [XML
    document](https://dom.spec.whatwg.org/#xml-document){#the-insertadjacenthtml()-method:xml-documents
    x-internal="xml-documents"}, if the given string is not
    well-formed).

    Throws a
    [\"`NoModificationAllowedError`\"](https://webidl.spec.whatwg.org/#nomodificationallowederror){#the-insertadjacenthtml()-method:nomodificationallowederror
    x-internal="nomodificationallowederror"}
    [`DOMException`{#the-insertadjacenthtml()-method:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
    if the given position isn\'t possible (e.g. inserting elements after
    the root element of a
    [`Document`{#the-insertadjacenthtml()-method:document}](dom.html#document)).

This method performs no sanitization to remove potentially-dangerous
elements and attributes like
[`script`{#the-insertadjacenthtml()-method:the-script-element}](scripting.html#the-script-element)
or [event handler content
attributes](webappapis.html#event-handler-content-attributes){#the-insertadjacenthtml()-method:event-handler-content-attributes}.

[`Element`{#the-insertadjacenthtml()-method:element}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"}\'s
[`insertAdjacentHTML(``position`{.variable}`, ``string`{.variable}`)`]{#dom-element-insertadjacenthtml
.dfn dfn-for="Element" dfn-type="method"} method steps are:

1.  Let `compliantString`{.variable} be the result of invoking the [Get
    Trusted Type compliant
    string](https://w3c.github.io/trusted-types/dist/spec/#get-trusted-type-compliant-string-algorithm){#the-insertadjacenthtml()-method:tt-getcompliantstring
    x-internal="tt-getcompliantstring"} algorithm with
    [`TrustedHTML`{#the-insertadjacenthtml()-method:tt-trustedhtml}](https://w3c.github.io/trusted-types/dist/spec/#trustedhtml){x-internal="tt-trustedhtml"},
    [this](https://webidl.spec.whatwg.org/#this){#the-insertadjacenthtml()-method:this
    x-internal="this"}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#the-insertadjacenthtml()-method:concept-relevant-global},
    `string`{.variable}, \"`Element insertAdjacentHTML`\", and
    \"`script`\".

2.  Let `context`{.variable} be null.

3.  Use the first matching item from this list:

    If `position`{.variable} is an [ASCII case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#the-insertadjacenthtml()-method:ascii-case-insensitive x-internal="ascii-case-insensitive"} match for the string \"`beforebegin`\"\
    If `position`{.variable} is an [ASCII case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#the-insertadjacenthtml()-method:ascii-case-insensitive-2 x-internal="ascii-case-insensitive"} match for the string \"`afterend`\"

    :   1.  Set `context`{.variable} to
            [this](https://webidl.spec.whatwg.org/#this){#the-insertadjacenthtml()-method:this-2
            x-internal="this"}\'s
            [parent](nav-history-apis.html#dom-parent){#the-insertadjacenthtml()-method:dom-parent}.

        2.  If `context`{.variable} is null or a
            [`Document`{#the-insertadjacenthtml()-method:document-2}](dom.html#document),
            throw a
            [\"`NoModificationAllowedError`\"](https://webidl.spec.whatwg.org/#nomodificationallowederror){#the-insertadjacenthtml()-method:nomodificationallowederror-2
            x-internal="nomodificationallowederror"}
            [`DOMException`{#the-insertadjacenthtml()-method:domexception-3}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    If `position`{.variable} is an [ASCII case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#the-insertadjacenthtml()-method:ascii-case-insensitive-3 x-internal="ascii-case-insensitive"} match for the string \"`afterbegin`\"\
    If `position`{.variable} is an [ASCII case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#the-insertadjacenthtml()-method:ascii-case-insensitive-4 x-internal="ascii-case-insensitive"} match for the string \"`beforeend`\"
    :   Set `context`{.variable} to
        [this](https://webidl.spec.whatwg.org/#this){#the-insertadjacenthtml()-method:this-3
        x-internal="this"}.

    Otherwise

    :   Throw a
        [\"`SyntaxError`\"](https://webidl.spec.whatwg.org/#syntaxerror){#the-insertadjacenthtml()-method:syntaxerror-2
        x-internal="syntaxerror"}
        [`DOMException`{#the-insertadjacenthtml()-method:domexception-4}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

4.  If `context`{.variable} is not an
    [`Element`{#the-insertadjacenthtml()-method:element-2}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"}
    or all of the following are true:

    - `context`{.variable}\'s [node
      document](https://dom.spec.whatwg.org/#concept-node-document){#the-insertadjacenthtml()-method:node-document
      x-internal="node-document"} is an HTML document;

    - `context`{.variable}\'s [local
      name](https://dom.spec.whatwg.org/#concept-element-local-name){#the-insertadjacenthtml()-method:concept-element-local-name
      x-internal="concept-element-local-name"} is
      \"[`html`{#the-insertadjacenthtml()-method:the-html-element}](semantics.html#the-html-element)\";
      and

    - `context`{.variable}\'s
      [namespace](https://dom.spec.whatwg.org/#concept-element-namespace){#the-insertadjacenthtml()-method:concept-element-namespace
      x-internal="concept-element-namespace"} is the [HTML
      namespace](https://infra.spec.whatwg.org/#html-namespace){#the-insertadjacenthtml()-method:html-namespace-2
      x-internal="html-namespace-2"},

    then set `context`{.variable} to the result of [creating an
    element](https://dom.spec.whatwg.org/#concept-create-element){#the-insertadjacenthtml()-method:create-an-element
    x-internal="create-an-element"} given
    [this](https://webidl.spec.whatwg.org/#this){#the-insertadjacenthtml()-method:this-4
    x-internal="this"}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-insertadjacenthtml()-method:node-document-2
    x-internal="node-document"}, \"`body`\", and the [HTML
    namespace](https://infra.spec.whatwg.org/#html-namespace){#the-insertadjacenthtml()-method:html-namespace-2-2
    x-internal="html-namespace-2"}.

5.  Let `fragment`{.variable} be the result of invoking the [fragment
    parsing algorithm
    steps](#fragment-parsing-algorithm-steps){#the-insertadjacenthtml()-method:fragment-parsing-algorithm-steps}
    with `context`{.variable} and `compliantString`{.variable}.

6.  Use the first matching item from this list:

    If `position`{.variable} is an [ASCII case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#the-insertadjacenthtml()-method:ascii-case-insensitive-5 x-internal="ascii-case-insensitive"} match for the string \"`beforebegin`\"

    :   [Insert](https://dom.spec.whatwg.org/#concept-node-insert){#the-insertadjacenthtml()-method:concept-node-insert
        x-internal="concept-node-insert"} `fragment`{.variable} into
        [this](https://webidl.spec.whatwg.org/#this){#the-insertadjacenthtml()-method:this-5
        x-internal="this"}\'s
        [parent](nav-history-apis.html#dom-parent){#the-insertadjacenthtml()-method:dom-parent-2}
        before
        [this](https://webidl.spec.whatwg.org/#this){#the-insertadjacenthtml()-method:this-6
        x-internal="this"}.

    If `position`{.variable} is an [ASCII case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#the-insertadjacenthtml()-method:ascii-case-insensitive-6 x-internal="ascii-case-insensitive"} match for the string \"`afterbegin`\"

    :   [Insert](https://dom.spec.whatwg.org/#concept-node-insert){#the-insertadjacenthtml()-method:concept-node-insert-2
        x-internal="concept-node-insert"} `fragment`{.variable} into
        [this](https://webidl.spec.whatwg.org/#this){#the-insertadjacenthtml()-method:this-7
        x-internal="this"} before its [first
        child](https://dom.spec.whatwg.org/#concept-tree-first-child){#the-insertadjacenthtml()-method:first-child
        x-internal="first-child"}.

    If `position`{.variable} is an [ASCII case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#the-insertadjacenthtml()-method:ascii-case-insensitive-7 x-internal="ascii-case-insensitive"} match for the string \"`beforeend`\"

    :   [Append](https://dom.spec.whatwg.org/#concept-node-append){#the-insertadjacenthtml()-method:concept-node-append
        x-internal="concept-node-append"} `fragment`{.variable} to
        [this](https://webidl.spec.whatwg.org/#this){#the-insertadjacenthtml()-method:this-8
        x-internal="this"}.

    If `position`{.variable} is an [ASCII case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#the-insertadjacenthtml()-method:ascii-case-insensitive-8 x-internal="ascii-case-insensitive"} match for the string \"`afterend`\"

    :   [Insert](https://dom.spec.whatwg.org/#concept-node-insert){#the-insertadjacenthtml()-method:concept-node-insert-3
        x-internal="concept-node-insert"} `fragment`{.variable} into
        [this](https://webidl.spec.whatwg.org/#this){#the-insertadjacenthtml()-method:this-9
        x-internal="this"}\'s
        [parent](nav-history-apis.html#dom-parent){#the-insertadjacenthtml()-method:dom-parent-3}
        before
        [this](https://webidl.spec.whatwg.org/#this){#the-insertadjacenthtml()-method:this-10
        x-internal="this"}\'s [next
        sibling](https://dom.spec.whatwg.org/#concept-tree-next-sibling){#the-insertadjacenthtml()-method:next-sibling
        x-internal="next-sibling"}.

As with other direct
[`Node`{#the-insertadjacenthtml()-method:node}](https://dom.spec.whatwg.org/#interface-node){x-internal="node"}-manipulation
APIs (and unlike
[`innerHTML`{#the-insertadjacenthtml()-method:dom-element-innerhtml}](#dom-element-innerhtml)),
[`insertAdjacentHTML()`{#the-insertadjacenthtml()-method:dom-element-insertadjacenthtml-3}](#dom-element-insertadjacenthtml)
does not include any special handling for
[`template`{#the-insertadjacenthtml()-method:the-template-element}](scripting.html#the-template-element)
elements. In most cases you will want to use
`templateEl.`[`content`](scripting.html#dom-template-content){#the-insertadjacenthtml()-method:dom-template-content}`.`[`insertAdjacentHTML()`](#dom-element-insertadjacenthtml){#the-insertadjacenthtml()-method:dom-element-insertadjacenthtml-4}
instead of directly manipulating the child nodes of a
[`template`{#the-insertadjacenthtml()-method:the-template-element-2}](scripting.html#the-template-element)
element.

#### [8.5.7]{.secno} The [`createContextualFragment()`{#the-createcontextualfragment()-method:dom-range-createcontextualfragment}](#dom-range-createcontextualfragment) method[](#the-createcontextualfragment()-method){.self-link} {#the-createcontextualfragment()-method}

The
[`createContextualFragment()`{#the-createcontextualfragment()-method:dom-range-createcontextualfragment-2}](#dom-range-createcontextualfragment)
method has a number of outstanding issues in the DOM Parsing and
Serialization [issue
tracker](https://github.com/w3c/DOM-Parsing/issues), documenting various
problems with its specification.

`docFragment`{.variable}` = ``range`{.variable}`.`[`createContextualFragment`](#dom-range-createcontextualfragment){#dom-range-createcontextualfragment-dev}`(``string`{.variable}`)`

:   Returns a
    [`DocumentFragment`{#the-createcontextualfragment()-method:documentfragment}](https://dom.spec.whatwg.org/#interface-documentfragment){x-internal="documentfragment"}
    created from the markup string `string`{.variable} using
    `range`{.variable}\'s [start
    node](https://dom.spec.whatwg.org/#concept-range-start-node){#the-createcontextualfragment()-method:concept-range-start-node
    x-internal="concept-range-start-node"} as the context in which
    `fragment`{.variable} is parsed.

This method performs no sanitization to remove potentially-dangerous
elements and attributes like
[`script`{#the-createcontextualfragment()-method:the-script-element}](scripting.html#the-script-element)
or [event handler content
attributes](webappapis.html#event-handler-content-attributes){#the-createcontextualfragment()-method:event-handler-content-attributes}.

``` idl
partial interface Range {
  [CEReactions, NewObject] DocumentFragment createContextualFragment((TrustedHTML or DOMString) string);
};
```

[`Range`{#the-createcontextualfragment()-method:range}](https://dom.spec.whatwg.org/#interface-range){x-internal="range"}\'s
[`createContextualFragment(``string`{.variable}`)`]{#dom-range-createcontextualfragment
.dfn dfn-for="Range" dfn-type="method"} method steps are:

1.  Let `compliantString`{.variable} be the result of invoking the [Get
    Trusted Type compliant
    string](https://w3c.github.io/trusted-types/dist/spec/#get-trusted-type-compliant-string-algorithm){#the-createcontextualfragment()-method:tt-getcompliantstring
    x-internal="tt-getcompliantstring"} algorithm with
    [`TrustedHTML`{#the-createcontextualfragment()-method:tt-trustedhtml-2}](https://w3c.github.io/trusted-types/dist/spec/#trustedhtml){x-internal="tt-trustedhtml"},
    [this](https://webidl.spec.whatwg.org/#this){#the-createcontextualfragment()-method:this
    x-internal="this"}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#the-createcontextualfragment()-method:concept-relevant-global},
    `string`{.variable}, \"`Range createContextualFragment`\", and
    \"`script`\".

2.  Let `node`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#the-createcontextualfragment()-method:this-2
    x-internal="this"}\'s [start
    node](https://dom.spec.whatwg.org/#concept-range-start-node){#the-createcontextualfragment()-method:concept-range-start-node-2
    x-internal="concept-range-start-node"}.

3.  Let `element`{.variable} be null.

4.  If `node`{.variable}
    [implements](https://webidl.spec.whatwg.org/#implements){#the-createcontextualfragment()-method:implements
    x-internal="implements"}
    [`Element`{#the-createcontextualfragment()-method:element}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"},
    set `element`{.variable} to `node`{.variable}.

5.  Otherwise, if `node`{.variable}
    [implements](https://webidl.spec.whatwg.org/#implements){#the-createcontextualfragment()-method:implements-2
    x-internal="implements"}
    [`Text`{#the-createcontextualfragment()-method:text}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
    or
    [`Comment`{#the-createcontextualfragment()-method:comment-2}](https://dom.spec.whatwg.org/#interface-comment){x-internal="comment-2"},
    set `element`{.variable} to `node`{.variable}\'s [parent
    element](https://dom.spec.whatwg.org/#parent-element){#the-createcontextualfragment()-method:parent-element
    x-internal="parent-element"}.

6.  If `element`{.variable} is null or all of the following are true:

    - `element`{.variable}\'s [node
      document](https://dom.spec.whatwg.org/#concept-node-document){#the-createcontextualfragment()-method:node-document
      x-internal="node-document"} is an HTML document;

    - `element`{.variable}\'s [local
      name](https://dom.spec.whatwg.org/#concept-element-local-name){#the-createcontextualfragment()-method:concept-element-local-name
      x-internal="concept-element-local-name"} is
      \"[`html`{#the-createcontextualfragment()-method:the-html-element}](semantics.html#the-html-element)\";
      and

    - `element`{.variable}\'s
      [namespace](https://dom.spec.whatwg.org/#concept-element-namespace){#the-createcontextualfragment()-method:concept-element-namespace
      x-internal="concept-element-namespace"} is the [HTML
      namespace](https://infra.spec.whatwg.org/#html-namespace){#the-createcontextualfragment()-method:html-namespace-2
      x-internal="html-namespace-2"},

    then set `element`{.variable} to the result of [creating an
    element](https://dom.spec.whatwg.org/#concept-create-element){#the-createcontextualfragment()-method:create-an-element
    x-internal="create-an-element"} given
    [this](https://webidl.spec.whatwg.org/#this){#the-createcontextualfragment()-method:this-3
    x-internal="this"}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-createcontextualfragment()-method:node-document-2
    x-internal="node-document"}, \"`body`\", and the [HTML
    namespace](https://infra.spec.whatwg.org/#html-namespace){#the-createcontextualfragment()-method:html-namespace-2-2
    x-internal="html-namespace-2"}.

7.  Let `fragment node`{.variable} be the result of invoking the
    [fragment parsing algorithm
    steps](#fragment-parsing-algorithm-steps){#the-createcontextualfragment()-method:fragment-parsing-algorithm-steps}
    with `element`{.variable} and `compliantString`{.variable}.

8.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#the-createcontextualfragment()-method:list-iterate
    x-internal="list-iterate"} `script`{.variable} of
    `fragment node`{.variable}\'s
    [`script`{#the-createcontextualfragment()-method:the-script-element-2}](scripting.html#the-script-element)
    element
    [descendants](https://dom.spec.whatwg.org/#concept-tree-descendant){#the-createcontextualfragment()-method:concept-tree-descendant
    x-internal="concept-tree-descendant"}:

    1.  Set `script`{.variable}\'s [already
        started](scripting.html#already-started){#the-createcontextualfragment()-method:already-started}
        to false.

    2.  Set `script`{.variable}\'s [parser
        document](scripting.html#parser-document){#the-createcontextualfragment()-method:parser-document}
        to null.

9.  Return `fragment node`{.variable}.

#### [8.5.8]{.secno} The [`XMLSerializer`{#the-xmlserializer-interface:xmlserializer}](#xmlserializer) interface[](#the-xmlserializer-interface){.self-link}

The
[`XMLSerializer`{#the-xmlserializer-interface:xmlserializer-2}](#xmlserializer)
interface has a number of outstanding issues in the DOM Parsing and
Serialization [issue
tracker](https://github.com/w3c/DOM-Parsing/issues), documenting various
problems with its specification. The remainder of DOM Parsing and
Serialization will be gradually upstreamed to this specification.

`xmlSerializer`{.variable}` = new `[`XMLSerializer`](#dom-xmlserializer-constructor){#dom-xmlserializer-constructor-dev}`()`
:   Constructs a new
    [`XMLSerializer`{#the-xmlserializer-interface:xmlserializer-3}](#xmlserializer)
    object.

`string`{.variable}` = ``xmlSerializer`{.variable}`.`[`serializeToString`](#dom-xmlserializer-serializetostring){#dom-xmlserializer-serializetostring-dev}`(``root`{.variable}`)`

:   Returns the result of serializing `root`{.variable} to XML.

    Throws an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-xmlserializer-interface:invalidstateerror
    x-internal="invalidstateerror"}
    [`DOMException`{#the-xmlserializer-interface:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
    if `root`{.variable} cannot be serialized to XML.

The design of
[`XMLSerializer`{#the-xmlserializer-interface:xmlserializer-4}](#xmlserializer),
as a class that needs to be constructed and then have its
[`serializeToString()`{#the-xmlserializer-interface:dom-xmlserializer-serializetostring}](#dom-xmlserializer-serializetostring)
method called, is an unfortunate historical artifact. If we were
designing this functionality today it would be a standalone function.

``` idl
[Exposed=Window]
interface XMLSerializer {
  constructor();

  DOMString serializeToString(Node root);
};
```

The [`new XMLSerializer()`]{#dom-xmlserializer-constructor .dfn}
constructor steps are to do nothing.

The
[`serializeToString(``root`{.variable}`)`]{#dom-xmlserializer-serializetostring
.dfn dfn-for="XMLSerializer" dfn-type="method"} method steps are:

1.  Return the [XML
    serialization](https://w3c.github.io/DOM-Parsing/#dfn-xml-serialization){#the-xmlserializer-interface:xml-serialization
    x-internal="xml-serialization"} of `root`{.variable} given false.

[← 8 Web application APIs](webappapis.html) --- [Table of Contents](./)
--- [8.6 Timers →](timers-and-user-prompts.html)
