HTML Standard

::: {#head .head .with-buttons}
[![WHATWG](https://resources.whatwg.org/logo.svg){.darkmode-aware
width="100" crossorigin="" height="100"}](https://whatwg.org/){.logo}

# HTML {#html .allcaps}

Living Standard --- Last Updated [2 June 2025]{.pubdate}
:::

[← 8.6 Timers](timers-and-user-prompts.html) --- [Table of Contents](./)
--- [8.10 Images →](imagebitmap-and-animations.html)

1.  ::: {#toc-webappapis}
    1.  [[8.9]{.secno} System state and
        capabilities](system-state.html#system-state-and-capabilities)
        1.  [[8.9.1]{.secno} The `Navigator`
            object](system-state.html#the-navigator-object)
            1.  [[8.9.1.1]{.secno} Client
                identification](system-state.html#client-identification)
            2.  [[8.9.1.2]{.secno} Language
                preferences](system-state.html#language-preferences)
            3.  [[8.9.1.3]{.secno} Browser
                state](system-state.html#navigator.online)
            4.  [[8.9.1.4]{.secno} Custom scheme handlers: the
                `registerProtocolHandler()`
                method](system-state.html#custom-handlers)
                1.  [[8.9.1.4.1]{.secno} Security and
                    privacy](system-state.html#security-and-privacy)
                2.  [[8.9.1.4.2]{.secno} User agent
                    automation](system-state.html#user-agent-automation)
            5.  [[8.9.1.5]{.secno} Cookies](system-state.html#cookies)
            6.  [[8.9.1.6]{.secno} PDF viewing
                support](system-state.html#pdf-viewing-support)
    :::

### [8.9]{.secno} System state and capabilities[](#system-state-and-capabilities){.self-link}

#### [8.9.1]{.secno} The [`Navigator`{#the-navigator-object:navigator}](#navigator) object[](#the-navigator-object){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Navigator](https://developer.mozilla.org/en-US/docs/Web/API/Navigator "The Navigator interface represents the state and the identity of the user agent. It allows scripts to query it and to register themselves to carry on some activities.")

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

Instances of
[`Navigator`{#the-navigator-object:navigator-2}](#navigator) represent
the identity and state of the user agent (the client). It has also been
used as a generic global under which various APIs are located, but this
is not precedent to build upon. Instead use the
[`WindowOrWorkerGlobalScope`{#the-navigator-object:windoworworkerglobalscope}](webappapis.html#windoworworkerglobalscope)
mixin or equivalent.

``` idl
[Exposed=Window]
interface Navigator {
  // objects implementing this interface also implement the interfaces given below
};
Navigator includes NavigatorID;
Navigator includes NavigatorLanguage;
Navigator includes NavigatorOnLine;
Navigator includes NavigatorContentUtils;
Navigator includes NavigatorCookies;
Navigator includes NavigatorPlugins;
Navigator includes NavigatorConcurrentHardware;
```

These interface mixins are defined separately so that
[`WorkerNavigator`{#the-navigator-object:workernavigator}](workers.html#workernavigator)
can reuse parts of the
[`Navigator`{#the-navigator-object:navigator-10}](#navigator) interface.

Each
[`Window`{#the-navigator-object:window}](nav-history-apis.html#window)
has an [associated `Navigator`]{#associated-navigator .dfn}, which is a
[`Navigator`{#the-navigator-object:navigator-11}](#navigator) object.
Upon creation of the
[`Window`{#the-navigator-object:window-2}](nav-history-apis.html#window)
object, its [associated
`Navigator`](#associated-navigator){#the-navigator-object:associated-navigator}
must be set to a
[new](https://webidl.spec.whatwg.org/#new){#the-navigator-object:new
x-internal="new"}
[`Navigator`{#the-navigator-object:navigator-12}](#navigator) object
created in the
[`Window`{#the-navigator-object:window-3}](nav-history-apis.html#window)
object\'s [relevant
realm](webappapis.html#concept-relevant-realm){#the-navigator-object:concept-relevant-realm}.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Window/navigator](https://developer.mozilla.org/en-US/docs/Web/API/Window/navigator "The Window.navigator read-only property returns a reference to the Navigator object, which has methods and properties about the application running the script.")

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

The [`navigator`]{#dom-navigator .dfn dfn-for="Window"
dfn-type="attribute"} and [`clientInformation`]{#dom-clientinformation
.dfn dfn-for="Window" dfn-type="attribute"} getter steps are to return
[this](https://webidl.spec.whatwg.org/#this){#the-navigator-object:this
x-internal="this"}\'s [associated
`Navigator`](#associated-navigator){#the-navigator-object:associated-navigator-2}.

##### [8.9.1.1]{.secno} Client identification[](#client-identification){.self-link}

``` idl
interface mixin NavigatorID {
  readonly attribute DOMString appCodeName; // constant "Mozilla"
  readonly attribute DOMString appName; // constant "Netscape"
  readonly attribute DOMString appVersion;
  readonly attribute DOMString platform;
  readonly attribute DOMString product; // constant "Gecko"
  [Exposed=Window] readonly attribute DOMString productSub;
  readonly attribute DOMString userAgent;
  [Exposed=Window] readonly attribute DOMString vendor;
  [Exposed=Window] readonly attribute DOMString vendorSub; // constant ""
};
```

In certain cases, despite the best efforts of the entire industry, web
browsers have bugs and limitations that web authors are forced to work
around.

This section defines a collection of attributes that can be used to
determine, from script, the kind of user agent in use, in order to work
around these issues.

The user agent has a [navigator compatibility
mode]{#concept-navigator-compatibility-mode .dfn}, which is either
*Chrome*, *Gecko*, or *WebKit*.

The [navigator compatibility
mode](#concept-navigator-compatibility-mode){#client-identification:concept-navigator-compatibility-mode}
constrains the
[`NavigatorID`{#client-identification:navigatorid}](#navigatorid) mixin
to the combinations of attribute values and presence of
[`taintEnabled()`{#client-identification:dom-navigator-taintenabled}](#dom-navigator-taintenabled)
and
[`oscpu`{#client-identification:dom-navigator-oscpu}](#dom-navigator-oscpu)
that are known to be compatible with existing web content.

Client detection should always be limited to detecting known current
versions; future versions and unknown versions should always be assumed
to be fully compliant.

`self`{.variable}`.`[`navigator`](#dom-navigator){#client-identification:dom-navigator}`.`[`appCodeName`](#dom-navigator-appcodename){#dom-navigator-appcodename-dev}

Returns the string \"`Mozilla`\".

`self`{.variable}`.`[`navigator`](#dom-navigator){#client-identification:dom-navigator-2}`.`[`appName`](#dom-navigator-appname){#dom-navigator-appname-dev}

Returns the string \"`Netscape`\".

`self`{.variable}`.`[`navigator`](#dom-navigator){#client-identification:dom-navigator-3}`.`[`appVersion`](#dom-navigator-appversion){#dom-navigator-appversion-dev}

Returns the version of the browser.

`self`{.variable}`.`[`navigator`](#dom-navigator){#client-identification:dom-navigator-4}`.`[`platform`](#dom-navigator-platform){#dom-navigator-platform-dev}

Returns the name of the platform.

`self`{.variable}`.`[`navigator`](#dom-navigator){#client-identification:dom-navigator-5}`.`[`product`](#dom-navigator-product){#dom-navigator-product-dev}

Returns the string \"`Gecko`\".

`window`{.variable}`.`[`navigator`](#dom-navigator){#client-identification:dom-navigator-6}`.`[`productSub`](#dom-navigator-productsub){#dom-navigator-productsub-dev}

Returns either the string \"`20030107`\", or the string \"`20100101`\".

`self`{.variable}`.`[`navigator`](#dom-navigator){#client-identification:dom-navigator-7}`.`[`userAgent`](#dom-navigator-useragent){#dom-navigator-useragent-dev}

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Navigator/userAgent](https://developer.mozilla.org/en-US/docs/Web/API/Navigator/userAgent "The Navigator.userAgent read-only property returns the user agent string for the current browser.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer4+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::

:::: feature
[WorkerNavigator/userAgent](https://developer.mozilla.org/en-US/docs/Web/API/WorkerNavigator/userAgent "The WorkerNavigator.userAgent read-only property returns the user agent string for the current browser.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::::

Returns the complete \``User-Agent`\` header.

`window`{.variable}`.`[`navigator`](#dom-navigator){#client-identification:dom-navigator-8}`.`[`vendor`](#dom-navigator-vendor){#dom-navigator-vendor-dev}

Returns either the empty string, the string \"`Apple Computer, Inc.`\",
or the string \"`Google Inc.`\".

`window`{.variable}`.`[`navigator`](#dom-navigator){#client-identification:dom-navigator-9}`.`[`vendorSub`](#dom-navigator-vendorsub){#dom-navigator-vendorsub-dev}

Returns the empty string.

[`appCodeName`]{#dom-navigator-appcodename .dfn dfn-for="NavigatorID" dfn-type="attribute"}
:   Must return the string \"`Mozilla`\".

[`appName`]{#dom-navigator-appname .dfn dfn-for="NavigatorID" dfn-type="attribute"}
:   Must return the string \"`Netscape`\".

[`appVersion`]{#dom-navigator-appversion .dfn dfn-for="NavigatorID" dfn-type="attribute"}

:   Must return the appropriate string that starts with \"`5.0 (`\", as
    follows:

    Let `trail`{.variable} be the substring of [default \``User-Agent`\`
    value](https://fetch.spec.whatwg.org/#default-user-agent-value){#client-identification:default-user-agent-value
    x-internal="default-user-agent-value"} that follows the
    \"`Mozilla/`\" prefix.

    If the [navigator compatibility mode](#concept-navigator-compatibility-mode){#client-identification:concept-navigator-compatibility-mode-2} is *Chrome* or *WebKit*
    :   Return `trail`{.variable}.

    If the [navigator compatibility mode](#concept-navigator-compatibility-mode){#client-identification:concept-navigator-compatibility-mode-3} is *Gecko*

    :   If `trail`{.variable} starts with \"`5.0 (Windows`\", then
        return \"`5.0 (Windows)`\".

        Otherwise, return the prefix of `trail`{.variable} up to but not
        including the first U+003B (;), concatenated with the character
        U+0029 RIGHT PARENTHESIS. For example, \"`5.0 (Macintosh)`\",
        \"`5.0 (Android 10)`\", or \"`5.0 (X11)`\".

[`platform`]{#dom-navigator-platform .dfn dfn-for="NavigatorID" dfn-type="attribute"}
:   Must return a string representing the platform on which the browser
    is executing (e.g. \"`MacIntel`\", \"`Win32`\", \"`Linux x86_64`\",
    \"`Linux armv81`\") or, for privacy and compatibility, a string that
    is commonly returned on another platform.

[`product`]{#dom-navigator-product .dfn dfn-for="NavigatorID" dfn-type="attribute"}
:   Must return the string \"`Gecko`\".

[`productSub`]{#dom-navigator-productsub .dfn dfn-for="NavigatorID" dfn-type="attribute"}

:   Must return the appropriate string from the following list:

    If the [navigator compatibility mode](#concept-navigator-compatibility-mode){#client-identification:concept-navigator-compatibility-mode-4} is *Chrome* or *WebKit*
    :   The string \"`20030107`\".

    If the [navigator compatibility mode](#concept-navigator-compatibility-mode){#client-identification:concept-navigator-compatibility-mode-5} is *Gecko*

    :   The string \"`20100101`\".

[`userAgent`]{#dom-navigator-useragent .dfn dfn-for="NavigatorID" dfn-type="attribute"}
:   Must return the [default \``User-Agent`\`
    value](https://fetch.spec.whatwg.org/#default-user-agent-value){#client-identification:default-user-agent-value-2
    x-internal="default-user-agent-value"}.

[`vendor`]{#dom-navigator-vendor .dfn dfn-for="NavigatorID" dfn-type="attribute"}

:   Must return the appropriate string from the following list:

    If the [navigator compatibility mode](#concept-navigator-compatibility-mode){#client-identification:concept-navigator-compatibility-mode-6} is *Chrome*
    :   The string \"`Google Inc.`\".

    If the [navigator compatibility mode](#concept-navigator-compatibility-mode){#client-identification:concept-navigator-compatibility-mode-7} is *Gecko*
    :   The empty string.

    If the [navigator compatibility mode](#concept-navigator-compatibility-mode){#client-identification:concept-navigator-compatibility-mode-8} is *WebKit*

    :   The string \"`Apple Computer, Inc.`\".

[`vendorSub`]{#dom-navigator-vendorsub .dfn dfn-for="NavigatorID" dfn-type="attribute"}

:   Must return the empty string.

If the [navigator compatibility
mode](#concept-navigator-compatibility-mode){#client-identification:concept-navigator-compatibility-mode-9}
is *Gecko*, then the user agent must also support the following partial
interface:

``` idl
partial interface mixin NavigatorID {
  [Exposed=Window] boolean taintEnabled(); // constant false
  [Exposed=Window] readonly attribute DOMString oscpu;
};
```

The [`taintEnabled()`]{#dom-navigator-taintenabled .dfn
dfn-for="NavigatorID" dfn-type="method"} method must return false.

The [`oscpu`]{#dom-navigator-oscpu .dfn dfn-for="NavigatorID"
dfn-type="attribute"} attribute\'s getter must return either the empty
string or a string representing the platform on which the browser is
executing, e.g. \"`Windows NT 10.0; Win64; x64`\", \"`Linux x86_64`\".

[![(This is a tracking
vector.)](https://resources.whatwg.org/tracking-vector.svg){width="46"
crossorigin=""
height="64"}](https://infra.spec.whatwg.org/#tracking-vector "There is a tracking vector here."){#client-identification:tracking-vector
.tracking-vector x-internal="tracking-vector"} Any information in this
API that varies from user to user can be used to profile the user. In
fact, if enough such information is available, a user can actually be
uniquely identified. For this reason, user agent implementers are
strongly urged to include as little information in this API as possible.

##### [8.9.1.2]{.secno} Language preferences[](#language-preferences){.self-link}

``` idl
interface mixin NavigatorLanguage {
  readonly attribute DOMString language;
  readonly attribute FrozenArray<DOMString> languages;
};
```

`self`{.variable}`.`[`navigator`](#dom-navigator){#language-preferences:dom-navigator}`.`[`language`](#dom-navigator-language){#dom-navigator-language-dev}

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Navigator/language](https://developer.mozilla.org/en-US/docs/Web/API/Navigator/language "The Navigator.language read-only property returns a string representing the preferred language of the user, usually the language of the browser UI.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera4+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer11]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android4+]{.firefox_android .yes}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::

:::: feature
[WorkerNavigator/language](https://developer.mozilla.org/en-US/docs/Web/API/WorkerNavigator/language "The WorkerNavigator.language read-only property returns a string representing the preferred language of the user, usually the language of the browser UI.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari10+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera4+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer11]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android4+]{.firefox_android .yes}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::::

Returns a language tag representing the user\'s preferred language.

`self`{.variable}`.`[`navigator`](#dom-navigator){#language-preferences:dom-navigator-2}`.`[`languages`](#dom-navigator-languages){#dom-navigator-languages-dev}

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Navigator/languages](https://developer.mozilla.org/en-US/docs/Web/API/Navigator/languages "The Navigator.languages read-only property returns an array of strings representing the user's preferred languages. The language is described using language tags according to RFC 5646: Tags for Identifying Languages (also known as BCP 47). In the returned array they are ordered by preference with the most preferred language first.")

Support in all current engines.

::: support
[Firefox32+]{.firefox .yes}[Safari10.1+]{.safari
.yes}[Chrome37+]{.chrome .yes}

------------------------------------------------------------------------

[Opera24+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)16+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet3.0+]{.samsunginternet_android .yes}[Opera
Android24+]{.opera_android .yes}
:::
::::

:::: feature
[WorkerNavigator/languages](https://developer.mozilla.org/en-US/docs/Web/API/WorkerNavigator/languages "The WorkerNavigator.languages read-only property returns an array of strings representing the user's preferred languages. The language is described using language tags according to RFC 5646: Tags for Identifying Languages (also known as BCP 47). In the returned array they are ordered by preference with the most preferred language first.")

Support in all current engines.

::: support
[Firefox32+]{.firefox .yes}[Safari10.1+]{.safari
.yes}[Chrome37+]{.chrome .yes}

------------------------------------------------------------------------

[Opera24+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)16+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet3.0+]{.samsunginternet_android .yes}[Opera
Android24+]{.opera_android .yes}
:::
::::
:::::::

Returns an array of language tags representing the user\'s preferred
languages, with the most preferred language first.

The most preferred language is the one returned by
[`navigator.language`{#language-preferences:dom-navigator-language-2}](#dom-navigator-language).

A
[`languagechange`{#language-preferences:event-languagechange}](indices.html#event-languagechange)
event is fired at the
[`Window`{#language-preferences:window}](nav-history-apis.html#window)
or
[`WorkerGlobalScope`{#language-preferences:workerglobalscope}](workers.html#workerglobalscope)
object when the user agent\'s understanding of what the user\'s
preferred languages are changes.

[`language`]{#dom-navigator-language .dfn dfn-for="NavigatorLanguage" dfn-type="attribute"}
:   Must return a valid BCP 47 language tag representing either [a
    plausible
    language](#a-plausible-language){#language-preferences:a-plausible-language}
    or the user\'s most preferred language.
    [\[BCP47\]](references.html#refsBCP47)

[`languages`]{#dom-navigator-languages .dfn dfn-for="NavigatorLanguage" dfn-type="attribute"}

:   Must return a [frozen
    array](https://webidl.spec.whatwg.org/#dfn-frozen-array-type){#language-preferences:frozen-array
    x-internal="frozen-array"} of valid BCP 47 language tags
    representing either one or more [plausible
    languages](#a-plausible-language){#language-preferences:a-plausible-language-2},
    or the user\'s preferred languages, ordered by preference with the
    most preferred language first. The same object must be returned
    until the user agent needs to return different values, or values in
    a different order. [\[BCP47\]](references.html#refsBCP47)

    Whenever the user agent needs to make the
    [`navigator.languages`{#language-preferences:dom-navigator-languages-2}](#dom-navigator-languages)
    attribute of a
    [`Window`{#language-preferences:window-2}](nav-history-apis.html#window)
    or
    [`WorkerGlobalScope`{#language-preferences:workerglobalscope-2}](workers.html#workerglobalscope)
    object `global`{.variable} return a new set of language tags, the
    user agent must [queue a global
    task](webappapis.html#queue-a-global-task){#language-preferences:queue-a-global-task}
    on the [DOM manipulation task
    source](webappapis.html#dom-manipulation-task-source){#language-preferences:dom-manipulation-task-source}
    given `global`{.variable} to [fire an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#language-preferences:concept-event-fire
    x-internal="concept-event-fire"} named
    [`languagechange`{#language-preferences:event-languagechange-2}](indices.html#event-languagechange)
    at `global`{.variable}, and wait until that task begins to be
    executed before actually returning a new value.

To determine [a plausible language]{#a-plausible-language .dfn}, the
user agent should bear in mind the following:

- [![(This is a tracking
  vector.)](https://resources.whatwg.org/tracking-vector.svg){width="46"
  crossorigin=""
  height="64"}](https://infra.spec.whatwg.org/#tracking-vector "There is a tracking vector here."){#language-preferences:tracking-vector
  .tracking-vector x-internal="tracking-vector"} Any information in this
  API that varies from user to user can be used to profile or identify
  the user.
- If the user is not using a service that obfuscates the user\'s point
  of origin (e.g. the Tor anonymity network), then the value that is
  least likely to distinguish the user from other users with similar
  origins (e.g. from the same IP address block) is the language used by
  the majority of such users. [\[TOR\]](references.html#refsTOR)
- If the user is using an anonymizing service, then the value
  \"`en-US`\" is suggested; if all users of the service use that same
  value, that reduces the possibility of distinguishing the users from
  each other.

[![(This is a tracking
vector.)](https://resources.whatwg.org/tracking-vector.svg){width="46"
crossorigin=""
height="64"}](https://infra.spec.whatwg.org/#tracking-vector "There is a tracking vector here."){#language-preferences:tracking-vector-2
.tracking-vector x-internal="tracking-vector"} To avoid introducing any
more fingerprinting vectors, user agents should use the same list for
the APIs defined in this function as for the HTTP
\`[`Accept-Language`{#language-preferences:http-accept-language}](https://httpwg.org/specs/rfc7231.html#header.accept-language){x-internal="http-accept-language"}\`
header.

##### [8.9.1.3]{.secno} []{#browser-state}Browser state[](#navigator.online){.self-link} {#navigator.online}

``` idl
interface mixin NavigatorOnLine {
  readonly attribute boolean onLine;
};
```

`self`{.variable}`.`[`navigator`](#dom-navigator){#navigator.online:dom-navigator}`.`[`onLine`](#dom-navigator-online){#dom-navigator-online-dev}

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Navigator/onLine](https://developer.mozilla.org/en-US/docs/Web/API/Navigator/onLine "Returns the online status of the browser. The property returns a boolean value, with true meaning online and false meaning offline. The property sends updates whenever the browser's ability to connect to the network changes. The update occurs when the user follows links or when a script requests a remote page. For example, the property should return false when users click links soon after they lose internet connection.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome2+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera3+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer4+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android4+]{.firefox_android .yes}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android18+]{.chrome_android .yes}[WebView Android[🔰
37+]{title="Partial implementation."}]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::

:::: feature
[WorkerNavigator/onLine](https://developer.mozilla.org/en-US/docs/Web/API/WorkerNavigator/onLine "Returns the online status of the browser. The property returns a boolean value, with true meaning online and false meaning offline. The property sends updates whenever the browser's ability to connect to the network changes. The update occurs when the user follows links or when a script requests a remote page.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android4+]{.firefox_android .yes}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView Android[🔰
4.4+]{title="Partial implementation."}]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::::

Returns false if the user agent is definitely offline (disconnected from
the network). Returns true if the user agent might be online.

The events
[`online`{#navigator.online:event-online}](indices.html#event-online)
and
[`offline`{#navigator.online:event-offline}](indices.html#event-offline)
are fired when the value of this attribute changes.

The [`onLine`]{#dom-navigator-online .dfn dfn-for="NavigatorOnLine"
dfn-type="attribute"} attribute must return false if the user agent will
not contact the network when the user follows links or when a script
requests a remote page (or knows that such an attempt would fail), and
must return true otherwise.

When the value that would be returned by the
[`navigator.onLine`{#navigator.online:dom-navigator-online-2}](#dom-navigator-online)
attribute of a
[`Window`{#navigator.online:window}](nav-history-apis.html#window) or
[`WorkerGlobalScope`{#navigator.online:workerglobalscope}](workers.html#workerglobalscope)
`global`{.variable} changes from true to false, the user agent must
[queue a global
task](webappapis.html#queue-a-global-task){#navigator.online:queue-a-global-task}
on the [networking task
source](webappapis.html#networking-task-source){#navigator.online:networking-task-source}
given `global`{.variable} to [fire an
event](https://dom.spec.whatwg.org/#concept-event-fire){#navigator.online:concept-event-fire
x-internal="concept-event-fire"} named
[`offline`{#navigator.online:event-offline-2}](indices.html#event-offline)
at `global`{.variable}.

On the other hand, when the value that would be returned by the
[`navigator.onLine`{#navigator.online:dom-navigator-online-3}](#dom-navigator-online)
attribute of a
[`Window`{#navigator.online:window-2}](nav-history-apis.html#window) or
[`WorkerGlobalScope`{#navigator.online:workerglobalscope-2}](workers.html#workerglobalscope)
`global`{.variable} changes from false to true, the user agent must
[queue a global
task](webappapis.html#queue-a-global-task){#navigator.online:queue-a-global-task-2}
on the [networking task
source](webappapis.html#networking-task-source){#navigator.online:networking-task-source-2}
given `global`{.variable} to [fire an
event](https://dom.spec.whatwg.org/#concept-event-fire){#navigator.online:concept-event-fire-2
x-internal="concept-event-fire"} named
[`online`{#navigator.online:event-online-2}](indices.html#event-online)
at the
[`Window`{#navigator.online:window-3}](nav-history-apis.html#window) or
[`WorkerGlobalScope`{#navigator.online:workerglobalscope-3}](workers.html#workerglobalscope)
object.

This attribute is inherently unreliable. A computer can be connected to
a network without having Internet access.

::: example
In this example, an indicator is updated as the browser goes online and
offline.

``` html
<!DOCTYPE HTML>
<html lang="en">
 <head>
  <title>Online status</title>
  <script>
   function updateIndicator() {
     document.getElementById('indicator').textContent = navigator.onLine ? 'online' : 'offline';
   }
  </script>
 </head>
 <body onload="updateIndicator()" ononline="updateIndicator()" onoffline="updateIndicator()">
  <p>The network is: <span id="indicator">(state unknown)</span>
 </body>
</html>
```
:::

##### [8.9.1.4]{.secno} Custom scheme handlers: the [`registerProtocolHandler()`{#custom-handlers:dom-navigator-registerprotocolhandler}](#dom-navigator-registerprotocolhandler) method[](#custom-handlers){.self-link} {#custom-handlers}

::::: {.mdn-anno .wrapped}
MDN

:::: feature
[Navigator/registerProtocolHandler](https://developer.mozilla.org/en-US/docs/Web/API/Navigator/registerProtocolHandler "The Navigator method registerProtocolHandler() lets websites register their ability to open or handle particular URL schemes (aka protocols).")

::: support
[Firefox2+]{.firefox .yes}[SafariNo]{.safari .no}[Chrome13+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)No]{.edge .no}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome AndroidNo]{.chrome_android .no}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

``` idl
interface mixin NavigatorContentUtils {
  [SecureContext] undefined registerProtocolHandler(DOMString scheme, USVString url);
  [SecureContext] undefined unregisterProtocolHandler(DOMString scheme, USVString url);
};
```

`window`{.variable}`.`[`navigator`](#dom-navigator){#custom-handlers:dom-navigator}`.`[`registerProtocolHandler`](#dom-navigator-registerprotocolhandler){#dom-navigator-registerprotocolhandler-dev}`(``scheme`{.variable}`, ``url`{.variable}`)`

Registers a handler for `scheme`{.variable} at `url`{.variable}. For
example, an online telephone messaging service could register itself as
a handler of the
[`sms:`{#custom-handlers:sms-protocol}](https://www.rfc-editor.org/rfc/rfc5724#section-2){x-internal="sms-protocol"}
scheme, so that if the user clicks on such a link, they are given the
opportunity to use that web site. [\[SMS\]](references.html#refsSMS)

The string \"`%s`\" in `url`{.variable} is used as a placeholder for
where to put the URL of the content to be handled.

Throws a
[\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#custom-handlers:securityerror
x-internal="securityerror"}
[`DOMException`{#custom-handlers:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if the user agent blocks the registration (this might happen if trying
to register as a handler for \"`http`\", for instance).

Throws a
[\"`SyntaxError`\"](https://webidl.spec.whatwg.org/#syntaxerror){#custom-handlers:syntaxerror
x-internal="syntaxerror"}
[`DOMException`{#custom-handlers:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if the \"`%s`\" string is missing in `url`{.variable}.

`window`{.variable}`.`[`navigator`](#dom-navigator){#custom-handlers:dom-navigator-2}`.`[`unregisterProtocolHandler`](#dom-navigator-unregisterprotocolhandler){#dom-navigator-unregisterprotocolhandler-dev}`(``scheme`{.variable}`, ``url`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**⚠**MDN

:::: feature
[Navigator/unregisterProtocolHandler](https://developer.mozilla.org/en-US/docs/Web/API/Navigator/unregisterProtocolHandler "The Navigator method unregisterProtocolHandler() removes a protocol handler for a given URL scheme.")

Support in one engine only.

::: support
[FirefoxNo]{.firefox .no}[SafariNo]{.safari .no}[Chrome38+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera25+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome AndroidNo]{.chrome_android .no}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

Unregisters the handler given by the arguments.

Throws a
[\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#custom-handlers:securityerror-2
x-internal="securityerror"}
[`DOMException`{#custom-handlers:domexception-3}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if the user agent blocks the deregistration (this might happen if with
invalid schemes, for instance).

Throws a
[\"`SyntaxError`\"](https://webidl.spec.whatwg.org/#syntaxerror){#custom-handlers:syntaxerror-2
x-internal="syntaxerror"}
[`DOMException`{#custom-handlers:domexception-4}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if the \"`%s`\" string is missing in `url`{.variable}.

The
[`registerProtocolHandler(``scheme`{.variable}`, ``url`{.variable}`)`]{#dom-navigator-registerprotocolhandler
.dfn dfn-for="NavigatorContentUtils" dfn-type="method"} method steps
are:

1.  Let (`normalizedScheme`{.variable},
    `normalizedURLString`{.variable}) be the result of running
    [normalize protocol handler
    parameters](#normalize-protocol-handler-parameters){#custom-handlers:normalize-protocol-handler-parameters}
    with `scheme`{.variable}, `url`{.variable}, and
    [this](https://webidl.spec.whatwg.org/#this){#custom-handlers:this
    x-internal="this"}\'s [relevant settings
    object](webappapis.html#relevant-settings-object){#custom-handlers:relevant-settings-object}.

2.  [In
    parallel](infrastructure.html#in-parallel){#custom-handlers:in-parallel}:
    [register a protocol handler]{#protocol-handler-registration .dfn
    export=""} for `normalizedScheme`{.variable} and
    `normalizedURLString`{.variable}. User agents may, within the
    constraints described, do whatever they like. A user agent could,
    for instance, prompt the user and offer the user the opportunity to
    add the site to a shortlist of handlers, or make the handlers their
    default, or cancel the request. User agents could also silently
    collect the information, providing it only when relevant to the
    user.

    User agents should keep track of which sites have registered
    handlers (even if the user has declined such registrations) so that
    the user is not repeatedly prompted with the same request.

    If the [`registerProtocolHandler()` automation
    mode](#registerprotocolhandler()-automation-mode){#custom-handlers:registerprotocolhandler()-automation-mode}
    of
    [this](https://webidl.spec.whatwg.org/#this){#custom-handlers:this-2
    x-internal="this"}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#custom-handlers:concept-relevant-global}\'s
    [associated
    `Document`](nav-history-apis.html#concept-document-window){#custom-handlers:concept-document-window}
    is not \"[`none`]{#rph-automation-mode-none .dfn}\", the user agent
    should first verify that it is in an automation context (see
    [WebDriver\'s security
    considerations](https://w3c.github.io/webdriver/#security){#custom-handlers:webdriver's-security-considerations
    x-internal="webdriver's-security-considerations"}). The user agent
    should then bypass the above communication of information and
    gathering of user consent, and instead do the following based on the
    value of the [`registerProtocolHandler()` automation
    mode](#registerprotocolhandler()-automation-mode){#custom-handlers:registerprotocolhandler()-automation-mode-2}:

    \"[`autoAccept`]{#rph-automation-mode-auto-accept .dfn}\"
    :   Act as if the user has seen the registration details and
        accepted the request.

    \"[`autoReject`]{#rph-automation-mode-auto-reject .dfn}\"

    :   Act as if the user has seen the registration details and
        rejected the request.

    When the [user agent uses this handler]{#protocol-handler-invocation
    .dfn lt="invoke a protocol
        handler" export=""} for a
    [URL](https://url.spec.whatwg.org/#concept-url){#custom-handlers:url
    x-internal="url"} `inputURL`{.variable}:

    1.  [Assert](https://infra.spec.whatwg.org/#assert){#custom-handlers:assert
        x-internal="assert"}: `inputURL`{.variable}\'s
        [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#custom-handlers:concept-url-scheme
        x-internal="concept-url-scheme"} is
        `normalizedScheme`{.variable}.

    2.  [Set the
        username](https://url.spec.whatwg.org/#set-the-username){#custom-handlers:set-the-username
        x-internal="set-the-username"} given `inputURL`{.variable} and
        the empty string.

    3.  [Set the
        password](https://url.spec.whatwg.org/#set-the-password){#custom-handlers:set-the-password
        x-internal="set-the-password"} given `inputURL`{.variable} and
        the empty string.

    4.  Let `inputURLString`{.variable} be the
        [serialization](https://url.spec.whatwg.org/#concept-url-serializer){#custom-handlers:concept-url-serializer
        x-internal="concept-url-serializer"} of `inputURL`{.variable}.

    5.  Let `encodedURL`{.variable} be the result of running [UTF-8
        percent-encode](https://url.spec.whatwg.org/#string-utf-8-percent-encode){#custom-handlers:utf-8-percent-encode
        x-internal="utf-8-percent-encode"} on
        `inputURLString`{.variable} using the [component percent-encode
        set](https://url.spec.whatwg.org/#component-percent-encode-set){#custom-handlers:component-percent-encode-set
        x-internal="component-percent-encode-set"}.

    6.  Let `handlerURLString`{.variable} be
        `normalizedURLString`{.variable}.

    7.  Replace the first instance of \"`%s`\" in
        `handlerURLString`{.variable} with `encodedURL`{.variable}.

    8.  Let `resultURL`{.variable} be the result of
        [parsing](https://url.spec.whatwg.org/#concept-url-parser){#custom-handlers:url-parser
        x-internal="url-parser"} `handlerURLString`{.variable}.

    9.  [Navigate](browsing-the-web.html#navigate){#custom-handlers:navigate}
        an appropriate
        [navigable](document-sequences.html#navigable){#custom-handlers:navigable}
        to `resultURL`{.variable}.

    ::: example
    If the user had visited a site at `https://example.com/` that made
    the following call:

    ``` js
    navigator.registerProtocolHandler('web+soup', 'soup?url=%s')
    ```

    \...and then, much later, while visiting `https://www.example.net/`,
    clicked on a link such as:

    ``` html
    <a href="web+soup:chicken-kïwi">Download our Chicken Kïwi soup!</a>
    ```

    \...then the UA might navigate to the following URL:

        https://example.com/soup?url=web+soup:chicken-k%C3%AFwi

    This site could then do whatever it is that it does with soup
    (synthesize it and ship it to the user, or whatever).
    :::

    This does not define when the handler is used. To some extent, the
    [processing model for navigating across
    documents](browsing-the-web.html#navigate){#custom-handlers:navigate-2}
    defines some cases where it is relevant, but in general user agents
    may use this information wherever they would otherwise consider
    handing schemes to native plugins or helper applications.

The
[`unregisterProtocolHandler(``scheme`{.variable}`, ``url`{.variable}`)`]{#dom-navigator-unregisterprotocolhandler
.dfn dfn-for="NavigatorContentUtils" dfn-type="method"} method steps
are:

1.  Let (`normalizedScheme`{.variable},
    `normalizedURLString`{.variable}) be the result of running
    [normalize protocol handler
    parameters](#normalize-protocol-handler-parameters){#custom-handlers:normalize-protocol-handler-parameters-2}
    with `scheme`{.variable}, `url`{.variable}, and
    [this](https://webidl.spec.whatwg.org/#this){#custom-handlers:this-3
    x-internal="this"}\'s [relevant settings
    object](webappapis.html#relevant-settings-object){#custom-handlers:relevant-settings-object-2}.

2.  [In
    parallel](infrastructure.html#in-parallel){#custom-handlers:in-parallel-2}:
    unregister the handler described by `normalizedScheme`{.variable}
    and `normalizedURLString`{.variable}.

------------------------------------------------------------------------

To [normalize protocol handler
parameters]{#normalize-protocol-handler-parameters .dfn export=""},
given a string `scheme`{.variable}, a string `url`{.variable}, and an
[environment settings
object](webappapis.html#environment-settings-object){#custom-handlers:environment-settings-object}
`environment`{.variable}, run these steps:

1.  Set `scheme`{.variable} to `scheme`{.variable}, [converted to ASCII
    lowercase](https://infra.spec.whatwg.org/#ascii-lowercase){#custom-handlers:converted-to-ascii-lowercase
    x-internal="converted-to-ascii-lowercase"}.

2.  If `scheme`{.variable} is neither a [safelisted
    scheme](#safelisted-scheme){#custom-handlers:safelisted-scheme} nor
    a string starting with \"`web+`\" followed by one or more [ASCII
    lower
    alphas](https://infra.spec.whatwg.org/#ascii-lower-alpha){#custom-handlers:lowercase-ascii-letters
    x-internal="lowercase-ascii-letters"}, then throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#custom-handlers:securityerror-3
    x-internal="securityerror"}
    [`DOMException`{#custom-handlers:domexception-5}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    This means that including a colon in `scheme`{.variable} (as in
    \"`mailto:`\") will throw.

    The following schemes are the [safelisted
    schemes]{#safelisted-scheme .dfn export=""}:

    - `bitcoin`
    - `ftp`
    - `ftps`
    - `geo`
    - `im`
    - `irc`
    - `ircs`
    - `magnet`
    - `mailto`
    - `matrix`
    - `mms`
    - `news`
    - `nntp`
    - `openpgp4fpr`
    - `sftp`
    - `sip`
    - `sms`
    - `smsto`
    - `ssh`
    - `tel`
    - `urn`
    - `webcal`
    - `wtai`
    - `xmpp`

    This list can be changed. If there are schemes that ought to be
    added, please send feedback.

3.  If `url`{.variable} does not contain \"`%s`\", then throw a
    [\"`SyntaxError`\"](https://webidl.spec.whatwg.org/#syntaxerror){#custom-handlers:syntaxerror-3
    x-internal="syntaxerror"}
    [`DOMException`{#custom-handlers:domexception-6}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

4.  Let `urlRecord`{.variable} be the result of [encoding-parsing a
    URL](urls-and-fetching.html#encoding-parsing-a-url){#custom-handlers:encoding-parsing-a-url}
    given `url`{.variable}, relative to `environment`{.variable}.

5.  If `urlRecord`{.variable} is failure, then throw a
    [\"`SyntaxError`\"](https://webidl.spec.whatwg.org/#syntaxerror){#custom-handlers:syntaxerror-4
    x-internal="syntaxerror"}
    [`DOMException`{#custom-handlers:domexception-7}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    This is forcibly the case if the `%s` placeholder is in the host or
    port of the URL.

6.  If `urlRecord`{.variable}\'s
    [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#custom-handlers:concept-url-scheme-2
    x-internal="concept-url-scheme"} is not an [HTTP(S)
    scheme](https://fetch.spec.whatwg.org/#http-scheme){#custom-handlers:http(s)-scheme
    x-internal="http(s)-scheme"} or `urlRecord`{.variable}\'s
    [origin](https://url.spec.whatwg.org/#concept-url-origin){#custom-handlers:concept-url-origin
    x-internal="concept-url-origin"} is not [same
    origin](browsers.html#same-origin){#custom-handlers:same-origin}
    with `environment`{.variable}\'s
    [origin](webappapis.html#concept-settings-object-origin){#custom-handlers:concept-settings-object-origin},
    then throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#custom-handlers:securityerror-4
    x-internal="securityerror"}
    [`DOMException`{#custom-handlers:domexception-8}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

7.  [Assert](https://infra.spec.whatwg.org/#assert){#custom-handlers:assert-2
    x-internal="assert"}: the result of [Is url potentially
    trustworthy?](https://w3c.github.io/webappsec-secure-contexts/#potentially-trustworthy-url){#custom-handlers:is-url-potentially-trustworthy
    x-internal="is-url-potentially-trustworthy"} given
    `urlRecord`{.variable} is \"`Potentially Trustworthy`\".

    Because [normalize protocol handler
    parameters](#normalize-protocol-handler-parameters){#custom-handlers:normalize-protocol-handler-parameters-3}
    is run within a [secure
    context](webappapis.html#secure-context){#custom-handlers:secure-context},
    this is implied by the [same
    origin](browsers.html#same-origin){#custom-handlers:same-origin-2}
    condition.

8.  Return (`scheme`{.variable}, `urlRecord`{.variable}).

    The
    [serialization](https://url.spec.whatwg.org/#concept-url-serializer){#custom-handlers:concept-url-serializer-2
    x-internal="concept-url-serializer"} of `urlRecord`{.variable} will
    by definition not be a [valid URL
    string](https://url.spec.whatwg.org/#valid-url-string){#custom-handlers:valid-url-string
    x-internal="valid-url-string"} as it includes the string \"`%s`\"
    which is not a valid component in a URL.

###### [8.9.1.4.1]{.secno} Security and privacy[](#security-and-privacy){.self-link}

Custom scheme handlers can introduce a number of concerns, in particular
privacy concerns.

**Hijacking all web usage.** User agents should not allow schemes that
are key to its normal operation, such as an [HTTP(S)
scheme](https://fetch.spec.whatwg.org/#http-scheme){#security-and-privacy:http(s)-scheme
x-internal="http(s)-scheme"}, to be rerouted through third-party sites.
This would allow a user\'s activities to be trivially tracked, and would
allow user information, even in secure connections, to be collected.

**Hijacking defaults.** User agents are strongly urged to not
automatically change any defaults, as this could lead the user to send
data to remote hosts that the user is not expecting. New handlers
registering themselves should never automatically cause those sites to
be used.

**Registration spamming.** User agents should consider the possibility
that a site will attempt to register a large number of handlers,
possibly from multiple domains (e.g., by redirecting through a series of
pages each on a different domain, and each registering a handler for
`web+spam:` --- analogous practices abusing other web browser features
have been used by pornography web sites for many years). User agents
should gracefully handle such hostile attempts, protecting the user.

**Hostile handler metadata.** User agents should protect against typical
attacks against strings embedded in their interface, for example
ensuring that markup or escape characters in such strings are not
executed, that null bytes are properly handled, that over-long strings
do not cause crashes or buffer overruns, and so forth.

**Leaking private data.** Web page authors may reference a custom scheme
handler using URL data considered private. They might do so with the
expectation that the user\'s choice of handler points to a page inside
the organization, ensuring that sensitive data will not be exposed to
third parties. However, a user may have registered a handler pointing to
an external site, resulting in a data leak to that third party.
Implementers might wish to consider allowing administrators to disable
custom handlers on certain subdomains, content types, or schemes.

**Interface interference.** User agents should be prepared to handle
intentionally long arguments to the methods. For example, if the user
interface exposed consists of an \"accept\" button and a \"deny\"
button, with the \"accept\" binding containing the name of the handler,
it\'s important that a long name not cause the \"deny\" button to be
pushed off the screen.

###### [8.9.1.4.2]{.secno} [User agent automation]{#rph-user-agent-automation .dfn}[](#user-agent-automation){.self-link}

Each [`Document`{#user-agent-automation:document}](dom.html#document)
has a [`registerProtocolHandler()` automation
mode]{#registerprotocolhandler()-automation-mode .dfn}. It defaults to
\"[`none`{#user-agent-automation:rph-automation-mode-none}](#rph-automation-mode-none)\",
but it also can be either
\"[`autoAccept`{#user-agent-automation:rph-automation-mode-auto-accept}](#rph-automation-mode-auto-accept)\"
or
\"[`autoReject`{#user-agent-automation:rph-automation-mode-auto-reject}](#rph-automation-mode-auto-reject)\".

For the purposes of user agent automation and website testing, this
standard defines [Set RPH Registration Mode]{#set-rph-registration-mode
.dfn} WebDriver [extension
command](https://w3c.github.io/webdriver/#dfn-extension-commands){#user-agent-automation:extension-command
x-internal="extension-command"}. It instructs the user agent to place a
[`Document`{#user-agent-automation:document-2}](dom.html#document) into
a mode where it will automatically simulate a user either accepting or
rejecting and registration confirmation prompt dialog.

HTTP Method

URI Template

\``POST`\`

`/session/{session id}/custom-handlers/set-mode`

The [remote end
steps](https://w3c.github.io/webdriver/#dfn-remote-end-steps){#user-agent-automation:remote-end-steps
x-internal="remote-end-steps"} are:

1.  If `parameters`{.variable} is not a JSON Object, return a [WebDriver
    error](https://w3c.github.io/webdriver/#dfn-errors){#user-agent-automation:webdriver-error
    x-internal="webdriver-error"} with [WebDriver error
    code](https://w3c.github.io/webdriver/#dfn-error-code){#user-agent-automation:webdriver-error-code
    x-internal="webdriver-error-code"} [invalid
    argument](https://w3c.github.io/webdriver/#dfn-invalid-argument){#user-agent-automation:invalid-argument
    x-internal="invalid-argument"}.

2.  Let `mode`{.variable} be the result of [getting a
    property](https://w3c.github.io/webdriver/#dfn-getting-properties){#user-agent-automation:getting-a-property
    x-internal="getting-a-property"} named \"`mode`\" from
    `parameters`{.variable}.

3.  If `mode`{.variable} is not
    \"[`autoAccept`{#user-agent-automation:rph-automation-mode-auto-accept-2}](#rph-automation-mode-auto-accept)\",
    \"[`autoReject`{#user-agent-automation:rph-automation-mode-auto-reject-2}](#rph-automation-mode-auto-reject)\",
    or
    \"[`none`{#user-agent-automation:rph-automation-mode-none-2}](#rph-automation-mode-none)\",
    return a [WebDriver
    error](https://w3c.github.io/webdriver/#dfn-errors){#user-agent-automation:webdriver-error-2
    x-internal="webdriver-error"} with [WebDriver error
    code](https://w3c.github.io/webdriver/#dfn-error-code){#user-agent-automation:webdriver-error-code-2
    x-internal="webdriver-error-code"} [invalid
    argument](https://w3c.github.io/webdriver/#dfn-invalid-argument){#user-agent-automation:invalid-argument-2
    x-internal="invalid-argument"}.

4.  Let `document`{.variable} be the [current browsing
    context](https://w3c.github.io/webdriver/#dfn-current-browsing-context){#user-agent-automation:webdriver-current-browsing-context
    x-internal="webdriver-current-browsing-context"}\'s [active
    document](document-sequences.html#active-document){#user-agent-automation:active-document}.

5.  Set `document`{.variable}\'s [`registerProtocolHandler()` automation
    mode](#registerprotocolhandler()-automation-mode){#user-agent-automation:registerprotocolhandler()-automation-mode}
    to `mode`{.variable}.

6.  Return
    [success](https://w3c.github.io/webdriver/#dfn-success){#user-agent-automation:success-value
    x-internal="success-value"} with data null.

##### [8.9.1.5]{.secno} Cookies[](#cookies){.self-link}

``` idl
interface mixin NavigatorCookies {
  readonly attribute boolean cookieEnabled;
};
```

`window`{.variable}`.`[`navigator`](#dom-navigator){#cookies:dom-navigator}`.`[`cookieEnabled`](#dom-navigator-cookieenabled){#dom-navigator-cookieenabled-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Navigator/cookieEnabled](https://developer.mozilla.org/en-US/docs/Web/API/Navigator/cookieEnabled "navigator.cookieEnabled returns a Boolean value that indicates whether cookies are enabled or not.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer4+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns false if setting a cookie will be ignored, and true otherwise.

The [`cookieEnabled`]{#dom-navigator-cookieenabled .dfn
dfn-for="NavigatorCookies" dfn-type="attribute"} attribute must return
true if the user agent attempts to handle cookies according to HTTP
State Management Mechanism, and false if it ignores cookie change
requests. [\[COOKIES\]](references.html#refsCOOKIES)

##### [8.9.1.6]{.secno} PDF viewing support[](#pdf-viewing-support){.self-link}

`window`{.variable}`.`[`navigator`](#dom-navigator){#pdf-viewing-support:dom-navigator}`.`[`pdfViewerEnabled`](#dom-navigator-pdfviewerenabled){#dom-navigator-pdfviewerenabled-dev}

:   Returns true if the user agent supports inline viewing of PDF files
    when
    [navigating](browsing-the-web.html#navigate){#pdf-viewing-support:navigate}
    to them, or false otherwise. In the latter case, PDF files will be
    handled by [external
    software](browsing-the-web.html#hand-off-to-external-software){#pdf-viewing-support:hand-off-to-external-software}.

``` idl
interface mixin NavigatorPlugins {
  [SameObject] readonly attribute PluginArray plugins;
  [SameObject] readonly attribute MimeTypeArray mimeTypes;
  boolean javaEnabled();
  readonly attribute boolean pdfViewerEnabled;
};

[Exposed=Window,
 LegacyUnenumerableNamedProperties]
interface PluginArray {
  undefined refresh();
  readonly attribute unsigned long length;
  getter Plugin? item(unsigned long index);
  getter Plugin? namedItem(DOMString name);
};

[Exposed=Window,
 LegacyUnenumerableNamedProperties]
interface MimeTypeArray {
  readonly attribute unsigned long length;
  getter MimeType? item(unsigned long index);
  getter MimeType? namedItem(DOMString name);
};

[Exposed=Window,
 LegacyUnenumerableNamedProperties]
interface Plugin {
  readonly attribute DOMString name;
  readonly attribute DOMString description;
  readonly attribute DOMString filename;
  readonly attribute unsigned long length;
  getter MimeType? item(unsigned long index);
  getter MimeType? namedItem(DOMString name);
};

[Exposed=Window]
interface MimeType {
  readonly attribute DOMString type;
  readonly attribute DOMString description;
  readonly attribute DOMString suffixes;
  readonly attribute Plugin enabledPlugin;
};
```

Although these days detecting PDF viewer support can be done via
[`navigator.pdfViewerEnabled`{#pdf-viewing-support:dom-navigator-pdfviewerenabled-2}](#dom-navigator-pdfviewerenabled),
for historical reasons, there are a number of complex and intertwined
interfaces that provide the same capability, which legacy code relies
on. This section specifies both the simple modern variant and the
complicated historical one.

Each user agent has a [PDF viewer supported]{#pdf-viewer-supported .dfn}
boolean, whose value is
[implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#pdf-viewing-support:implementation-defined
x-internal="implementation-defined"} (and might vary according to user
preferences).

This value also impacts the
[navigation](browsing-the-web.html#navigate){#pdf-viewing-support:navigate-2}
processing model.

------------------------------------------------------------------------

Each
[`Window`{#pdf-viewing-support:window}](nav-history-apis.html#window)
object has a [PDF viewer plugin objects]{#pdf-viewer-plugin-objects
.dfn} list. If the user agent\'s [PDF viewer
supported](#pdf-viewer-supported){#pdf-viewing-support:pdf-viewer-supported}
is false, then it is the empty list. Otherwise, it is a list containing
five [`Plugin`{#pdf-viewing-support:dom-plugin}](#dom-plugin) objects,
whose
[names](#concept-plugin-name){#pdf-viewing-support:concept-plugin-name}
are, respectively:

0.  \"`PDF Viewer`\"
1.  \"`Chrome PDF Viewer`\"
2.  \"`Chromium PDF Viewer`\"
3.  \"`Microsoft Edge PDF Viewer`\"
4.  \"`WebKit built-in PDF`\"

The values of the above list form the [PDF viewer plugin
names]{#pdf-viewer-plugin-names .dfn} list.

These names were chosen based on evidence of what websites historically
search for, and thus what is necessary for user agents to expose in
order to maintain compatibility with existing content. They are ordered
alphabetically. The \"`PDF Viewer`\" name was then inserted in the 0th
position so that the
[`enabledPlugin`{#pdf-viewing-support:dom-mimetype-enabledplugin-2}](#dom-mimetype-enabledplugin)
getter could point to a generic plugin name.

Each
[`Window`{#pdf-viewing-support:window-2}](nav-history-apis.html#window)
object has a [PDF viewer mime type
objects]{#pdf-viewer-mime-type-objects .dfn} list. If the user agent\'s
[PDF viewer
supported](#pdf-viewer-supported){#pdf-viewing-support:pdf-viewer-supported-2}
is false, then it is the empty list. Otherwise, it is a list containing
two [`MimeType`{#pdf-viewing-support:mimetype-5}](#mimetype) objects,
whose
[types](#concept-mimetype-type){#pdf-viewing-support:concept-mimetype-type}
are, respectively:

0.  \"`application/pdf`\"
1.  \"`text/pdf`\"

The values of the above list form the [PDF viewer mime
types]{#pdf-viewer-mime-types .dfn} list.

------------------------------------------------------------------------

Each
[`NavigatorPlugins`{#pdf-viewing-support:navigatorplugins}](#navigatorplugins)
object has a [plugins array]{#plugins-array .dfn}, which is a new
[`PluginArray`{#pdf-viewing-support:pluginarray-2}](#pluginarray), and a
[mime types array]{#mime-types-array .dfn}, which is a new
[`MimeTypeArray`{#pdf-viewing-support:mimetypearray-2}](#mimetypearray).

The
[`NavigatorPlugins`{#pdf-viewing-support:navigatorplugins-2}](#navigatorplugins)
mixin\'s [`plugins`]{#dom-navigator-plugins .dfn
dfn-for="NavigatorPlugins" dfn-type="attribute"} getter steps are to
return
[this](https://webidl.spec.whatwg.org/#this){#pdf-viewing-support:this
x-internal="this"}\'s [plugins
array](#plugins-array){#pdf-viewing-support:plugins-array}.

The
[`NavigatorPlugins`{#pdf-viewing-support:navigatorplugins-3}](#navigatorplugins)
mixin\'s [`mimeTypes`]{#dom-navigator-mimetypes .dfn
dfn-for="NavigatorPlugins" dfn-type="attribute"} getter steps are to
return
[this](https://webidl.spec.whatwg.org/#this){#pdf-viewing-support:this-2
x-internal="this"}\'s [mime types
array](#mime-types-array){#pdf-viewing-support:mime-types-array}.

The
[`NavigatorPlugins`{#pdf-viewing-support:navigatorplugins-4}](#navigatorplugins)
mixin\'s [`javaEnabled()`]{#dom-navigator-javaenabled .dfn
dfn-for="NavigatorPlugins" dfn-type="method"} method steps are to return
false.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Navigator/pdfViewerEnabled](https://developer.mozilla.org/en-US/docs/Web/API/Navigator/pdfViewerEnabled "The pdfViewerEnabled property of the Navigator interface indicates whether the browser supports inline display of PDF files when navigating to them.")

Support in all current engines.

::: support
[Firefox99+]{.firefox .yes}[Safari16.4+]{.safari
.yes}[Chrome94+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge94+]{.edge_blink .yes}

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

The
[`NavigatorPlugins`{#pdf-viewing-support:navigatorplugins-5}](#navigatorplugins)
mixin\'s [`pdfViewerEnabled`]{#dom-navigator-pdfviewerenabled .dfn
dfn-for="NavigatorPlugins" dfn-type="attribute"} getter steps are to
return the user agent\'s [PDF viewer
supported](#pdf-viewer-supported){#pdf-viewing-support:pdf-viewer-supported-3}.

------------------------------------------------------------------------

The [`PluginArray`{#pdf-viewing-support:pluginarray-3}](#pluginarray)
interface [supports named
properties](https://webidl.spec.whatwg.org/#dfn-support-named-properties){#pdf-viewing-support:support-named-properties
x-internal="support-named-properties"}. If the user agent\'s [PDF viewer
supported](#pdf-viewer-supported){#pdf-viewing-support:pdf-viewer-supported-4}
is true, then they are the [PDF viewer plugin
names](#pdf-viewer-plugin-names){#pdf-viewing-support:pdf-viewer-plugin-names}.
Otherwise, they are the empty list.

The [`PluginArray`{#pdf-viewing-support:pluginarray-4}](#pluginarray)
interface\'s
[`namedItem(``name`{.variable}`)`]{#dom-pluginarray-nameditem .dfn
dfn-for="PluginArray" dfn-type="method"} method steps are:

1.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#pdf-viewing-support:list-iterate
    x-internal="list-iterate"}
    [`Plugin`{#pdf-viewing-support:dom-plugin-2}](#dom-plugin)
    `plugin`{.variable} of
    [this](https://webidl.spec.whatwg.org/#this){#pdf-viewing-support:this-3
    x-internal="this"}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#pdf-viewing-support:concept-relevant-global}\'s
    [PDF viewer plugin
    objects](#pdf-viewer-plugin-objects){#pdf-viewing-support:pdf-viewer-plugin-objects}:
    if `plugin`{.variable}\'s
    [name](#concept-plugin-name){#pdf-viewing-support:concept-plugin-name-2}
    is `name`{.variable}, then return `plugin`{.variable}.

2.  Return null.

The [`PluginArray`{#pdf-viewing-support:pluginarray-5}](#pluginarray)
interface [supports indexed
properties](https://webidl.spec.whatwg.org/#dfn-support-indexed-properties){#pdf-viewing-support:supports-indexed-properties
x-internal="supports-indexed-properties"}. The [supported property
indices](https://webidl.spec.whatwg.org/#dfn-supported-property-indices){#pdf-viewing-support:supported-property-indices
x-internal="supported-property-indices"} are the
[indices](https://infra.spec.whatwg.org/#list-get-the-indices){#pdf-viewing-support:indices
x-internal="indices"} of
[this](https://webidl.spec.whatwg.org/#this){#pdf-viewing-support:this-4
x-internal="this"}\'s [relevant global
object](webappapis.html#concept-relevant-global){#pdf-viewing-support:concept-relevant-global-2}\'s
[PDF viewer plugin
objects](#pdf-viewer-plugin-objects){#pdf-viewing-support:pdf-viewer-plugin-objects-2}.

The [`PluginArray`{#pdf-viewing-support:pluginarray-6}](#pluginarray)
interface\'s [`item(``index`{.variable}`)`]{#dom-pluginarray-item .dfn
dfn-for="PluginArray" dfn-type="method"} method steps are:

1.  Let `plugins`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#pdf-viewing-support:this-5
    x-internal="this"}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#pdf-viewing-support:concept-relevant-global-3}\'s
    [PDF viewer plugin
    objects](#pdf-viewer-plugin-objects){#pdf-viewing-support:pdf-viewer-plugin-objects-3}.

2.  If `index`{.variable} \< `plugins`{.variable}\'s
    [size](https://infra.spec.whatwg.org/#list-size){#pdf-viewing-support:list-size
    x-internal="list-size"}, then return
    `plugins`{.variable}\[`index`{.variable}\].

3.  Return null.

The [`PluginArray`{#pdf-viewing-support:pluginarray-7}](#pluginarray)
interface\'s [`length`]{#dom-pluginarray-length .dfn
dfn-for="PluginArray" dfn-type="attribute"} getter steps are to return
[this](https://webidl.spec.whatwg.org/#this){#pdf-viewing-support:this-6
x-internal="this"}\'s [relevant global
object](webappapis.html#concept-relevant-global){#pdf-viewing-support:concept-relevant-global-4}\'s
[PDF viewer plugin
objects](#pdf-viewer-plugin-objects){#pdf-viewing-support:pdf-viewer-plugin-objects-4}\'s
[size](https://infra.spec.whatwg.org/#list-size){#pdf-viewing-support:list-size-2
x-internal="list-size"}.

The [`PluginArray`{#pdf-viewing-support:pluginarray-8}](#pluginarray)
interface\'s [`refresh()`]{#dom-pluginarray-refresh .dfn
dfn-for="PluginArray" dfn-type="method"} method steps are to do nothing.

------------------------------------------------------------------------

The
[`MimeTypeArray`{#pdf-viewing-support:mimetypearray-3}](#mimetypearray)
interface [supports named
properties](https://webidl.spec.whatwg.org/#dfn-support-named-properties){#pdf-viewing-support:support-named-properties-2
x-internal="support-named-properties"}. If the user agent\'s [PDF viewer
supported](#pdf-viewer-supported){#pdf-viewing-support:pdf-viewer-supported-5}
is true, then they are the [PDF viewer mime
types](#pdf-viewer-mime-types){#pdf-viewing-support:pdf-viewer-mime-types}.
Otherwise, they are the empty list.

The
[`MimeTypeArray`{#pdf-viewing-support:mimetypearray-4}](#mimetypearray)
interface\'s
[`namedItem(``name`{.variable}`)`]{#dom-mimetypearray-nameditem .dfn
dfn-for="MimeTypeArray" dfn-type="method"} method steps are:

1.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#pdf-viewing-support:list-iterate-2
    x-internal="list-iterate"}
    [`MimeType`{#pdf-viewing-support:mimetype-6}](#mimetype)
    `mimeType`{.variable} of
    [this](https://webidl.spec.whatwg.org/#this){#pdf-viewing-support:this-7
    x-internal="this"}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#pdf-viewing-support:concept-relevant-global-5}\'s
    [PDF viewer mime type
    objects](#pdf-viewer-mime-type-objects){#pdf-viewing-support:pdf-viewer-mime-type-objects}:
    if `mimeType`{.variable}\'s
    [type](#concept-mimetype-type){#pdf-viewing-support:concept-mimetype-type-2}
    is `name`{.variable}, then return `mimeType`{.variable}.

2.  Return null.

The
[`MimeTypeArray`{#pdf-viewing-support:mimetypearray-5}](#mimetypearray)
interface [supports indexed
properties](https://webidl.spec.whatwg.org/#dfn-support-indexed-properties){#pdf-viewing-support:supports-indexed-properties-2
x-internal="supports-indexed-properties"}. The [supported property
indices](https://webidl.spec.whatwg.org/#dfn-supported-property-indices){#pdf-viewing-support:supported-property-indices-2
x-internal="supported-property-indices"} are the
[indices](https://infra.spec.whatwg.org/#list-get-the-indices){#pdf-viewing-support:indices-2
x-internal="indices"} of
[this](https://webidl.spec.whatwg.org/#this){#pdf-viewing-support:this-8
x-internal="this"}\'s [relevant global
object](webappapis.html#concept-relevant-global){#pdf-viewing-support:concept-relevant-global-6}\'s
[PDF viewer mime type
objects](#pdf-viewer-mime-type-objects){#pdf-viewing-support:pdf-viewer-mime-type-objects-2}.

The
[`MimeTypeArray`{#pdf-viewing-support:mimetypearray-6}](#mimetypearray)
interface\'s [`item(``index`{.variable}`)`]{#dom-mimetypearray-item .dfn
dfn-for="MimeTypeArray" dfn-type="method"} method steps are:

1.  Let `mimeTypes`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#pdf-viewing-support:this-9
    x-internal="this"}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#pdf-viewing-support:concept-relevant-global-7}\'s
    [PDF viewer mime type
    objects](#pdf-viewer-mime-type-objects){#pdf-viewing-support:pdf-viewer-mime-type-objects-3}.

2.  If `index`{.variable} \< `mimeTypes`{.variable}\'s
    [size](https://infra.spec.whatwg.org/#list-size){#pdf-viewing-support:list-size-3
    x-internal="list-size"}, then return
    `mimeTypes`{.variable}\[`index`{.variable}\].

3.  Return null.

The
[`MimeTypeArray`{#pdf-viewing-support:mimetypearray-7}](#mimetypearray)
interface\'s [`length`]{#dom-mimetypearray-length .dfn
dfn-for="MimeTypeArray" dfn-type="attribute"} getter steps are to return
[this](https://webidl.spec.whatwg.org/#this){#pdf-viewing-support:this-10
x-internal="this"}\'s [relevant global
object](webappapis.html#concept-relevant-global){#pdf-viewing-support:concept-relevant-global-8}\'s
[PDF viewer mime type
objects](#pdf-viewer-mime-type-objects){#pdf-viewing-support:pdf-viewer-mime-type-objects-4}\'s
[size](https://infra.spec.whatwg.org/#list-size){#pdf-viewing-support:list-size-4
x-internal="list-size"}.

------------------------------------------------------------------------

Each [`Plugin`{#pdf-viewing-support:dom-plugin-3}](#dom-plugin) object
has a [name]{#concept-plugin-name .dfn}, which is set when the object is
created.

The [`Plugin`{#pdf-viewing-support:dom-plugin-4}](#dom-plugin)
interface\'s [`name`]{#dom-plugin-name .dfn dfn-for="Plugin"
dfn-type="attribute"} getter steps are to return
[this](https://webidl.spec.whatwg.org/#this){#pdf-viewing-support:this-11
x-internal="this"}\'s
[name](#concept-plugin-name){#pdf-viewing-support:concept-plugin-name-3}.

The [`Plugin`{#pdf-viewing-support:dom-plugin-5}](#dom-plugin)
interface\'s [`description`]{#dom-plugin-description .dfn
dfn-for="Plugin" dfn-type="attribute"} getter steps are to return
\"`Portable Document Format`\".

The [`Plugin`{#pdf-viewing-support:dom-plugin-6}](#dom-plugin)
interface\'s [`filename`]{#dom-plugin-filename .dfn dfn-for="Plugin"
dfn-type="attribute"} getter steps are to return
\"`internal-pdf-viewer`\".

The [`Plugin`{#pdf-viewing-support:dom-plugin-7}](#dom-plugin) interface
[supports named
properties](https://webidl.spec.whatwg.org/#dfn-support-named-properties){#pdf-viewing-support:support-named-properties-3
x-internal="support-named-properties"}. If the user agent\'s [PDF viewer
supported](#pdf-viewer-supported){#pdf-viewing-support:pdf-viewer-supported-6}
is true, then they are the [PDF viewer mime
types](#pdf-viewer-mime-types){#pdf-viewing-support:pdf-viewer-mime-types-2}.
Otherwise, they are the empty list.

The [`Plugin`{#pdf-viewing-support:dom-plugin-8}](#dom-plugin)
interface\'s [`namedItem(``name`{.variable}`)`]{#dom-plugin-nameditem
.dfn dfn-for="Plugin" dfn-type="method"} method steps are:

1.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#pdf-viewing-support:list-iterate-3
    x-internal="list-iterate"}
    [`MimeType`{#pdf-viewing-support:mimetype-7}](#mimetype)
    `mimeType`{.variable} of
    [this](https://webidl.spec.whatwg.org/#this){#pdf-viewing-support:this-12
    x-internal="this"}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#pdf-viewing-support:concept-relevant-global-9}\'s
    [PDF viewer mime type
    objects](#pdf-viewer-mime-type-objects){#pdf-viewing-support:pdf-viewer-mime-type-objects-5}:
    if `mimeType`{.variable}\'s
    [type](#concept-mimetype-type){#pdf-viewing-support:concept-mimetype-type-3}
    is `name`{.variable}, then return `mimeType`{.variable}.

2.  Return null.

The [`Plugin`{#pdf-viewing-support:dom-plugin-9}](#dom-plugin) interface
[supports indexed
properties](https://webidl.spec.whatwg.org/#dfn-support-indexed-properties){#pdf-viewing-support:supports-indexed-properties-3
x-internal="supports-indexed-properties"}. The [supported property
indices](https://webidl.spec.whatwg.org/#dfn-supported-property-indices){#pdf-viewing-support:supported-property-indices-3
x-internal="supported-property-indices"} are the
[indices](https://infra.spec.whatwg.org/#list-get-the-indices){#pdf-viewing-support:indices-3
x-internal="indices"} of
[this](https://webidl.spec.whatwg.org/#this){#pdf-viewing-support:this-13
x-internal="this"}\'s [relevant global
object](webappapis.html#concept-relevant-global){#pdf-viewing-support:concept-relevant-global-10}\'s
[PDF viewer mime type
objects](#pdf-viewer-mime-type-objects){#pdf-viewing-support:pdf-viewer-mime-type-objects-6}.

The [`Plugin`{#pdf-viewing-support:dom-plugin-10}](#dom-plugin)
interface\'s [`item(``index`{.variable}`)`]{#dom-plugin-item .dfn
dfn-for="Plugin" dfn-type="method"} method steps are:

1.  Let `mimeTypes`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#pdf-viewing-support:this-14
    x-internal="this"}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#pdf-viewing-support:concept-relevant-global-11}\'s
    [PDF viewer mime type
    objects](#pdf-viewer-mime-type-objects){#pdf-viewing-support:pdf-viewer-mime-type-objects-7}.

2.  If `index`{.variable} \< `mimeTypes`{.variable}\'s
    [size](https://infra.spec.whatwg.org/#list-size){#pdf-viewing-support:list-size-5
    x-internal="list-size"}, then return
    `mimeTypes`{.variable}\[`index`{.variable}\].

3.  Return null.

The [`Plugin`{#pdf-viewing-support:dom-plugin-11}](#dom-plugin)
interface\'s [`length`]{#dom-plugin-length .dfn dfn-for="Plugin"
dfn-type="attribute"} getter steps are to return
[this](https://webidl.spec.whatwg.org/#this){#pdf-viewing-support:this-15
x-internal="this"}\'s [relevant global
object](webappapis.html#concept-relevant-global){#pdf-viewing-support:concept-relevant-global-12}\'s
[PDF viewer mime type
objects](#pdf-viewer-mime-type-objects){#pdf-viewing-support:pdf-viewer-mime-type-objects-8}\'s
[size](https://infra.spec.whatwg.org/#list-size){#pdf-viewing-support:list-size-6
x-internal="list-size"}.

------------------------------------------------------------------------

Each [`MimeType`{#pdf-viewing-support:mimetype-8}](#mimetype) object has
a [type]{#concept-mimetype-type .dfn}, which is set when the object is
created.

The [`MimeType`{#pdf-viewing-support:mimetype-9}](#mimetype)
interface\'s [`type`]{#dom-mimetype-type .dfn dfn-for="MimeType"
dfn-type="attribute"} getter steps are to return
[this](https://webidl.spec.whatwg.org/#this){#pdf-viewing-support:this-16
x-internal="this"}\'s
[type](#concept-mimetype-type){#pdf-viewing-support:concept-mimetype-type-4}.

The [`MimeType`{#pdf-viewing-support:mimetype-10}](#mimetype)
interface\'s [`description`]{#dom-mimetype-description .dfn
dfn-for="MimeType" dfn-type="attribute"} getter steps are to return
\"`Portable Document Format`\".

The [`MimeType`{#pdf-viewing-support:mimetype-11}](#mimetype)
interface\'s [`suffixes`]{#dom-mimetype-suffixes .dfn dfn-for="MimeType"
dfn-type="attribute"} getter steps are to return \"`pdf`\".

The [`MimeType`{#pdf-viewing-support:mimetype-12}](#mimetype)
interface\'s [`enabledPlugin`]{#dom-mimetype-enabledplugin .dfn
dfn-for="MimeType" dfn-type="attribute"} getter steps are to return
[this](https://webidl.spec.whatwg.org/#this){#pdf-viewing-support:this-17
x-internal="this"}\'s [relevant global
object](webappapis.html#concept-relevant-global){#pdf-viewing-support:concept-relevant-global-13}\'s
[PDF viewer plugin
objects](#pdf-viewer-plugin-objects){#pdf-viewing-support:pdf-viewer-plugin-objects-5}\[0\]
(i.e., the generic \"`PDF Viewer`\" one).

[← 8.6 Timers](timers-and-user-prompts.html) --- [Table of Contents](./)
--- [8.10 Images →](imagebitmap-and-animations.html)
