## [3]{.secno} Semantics, structure, and APIs of HTML documents[](#dom){.self-link} {#dom}

### [3.1]{.secno} Documents[](#documents){.self-link}

Every XML and HTML document in an HTML UA is represented by a
[`Document`{#documents:document}](#document) object.
[\[DOM\]](references.html#refsDOM)

The [`Document`{#documents:document-2}](#document) object\'s
[[URL](https://dom.spec.whatwg.org/#concept-document-url)]{#the-document's-address
.dfn} is defined in DOM. It is initially set when the
[`Document`{#documents:document-3}](#document) object is created, but
can change during the lifetime of the
[`Document`{#documents:document-4}](#document) object; for example, it
changes when the user
[navigates](browsing-the-web.html#navigate){#documents:navigate} to a
[fragment](browsing-the-web.html#navigate-fragid){#documents:navigate-fragid}
on the page and when the
[`pushState()`{#documents:dom-history-pushstate}](nav-history-apis.html#dom-history-pushstate)
method is called with a new
[URL](https://url.spec.whatwg.org/#concept-url){#documents:url
x-internal="url"}. [\[DOM\]](references.html#refsDOM)

Interactive user agents typically expose the
[`Document`{#documents:document-5}](#document) object\'s
[URL](https://dom.spec.whatwg.org/#concept-document-url){#documents:the-document's-address
x-internal="the-document's-address"} in their user interface. This is
the primary mechanism by which a user can tell if a site is attempting
to impersonate another.

The [`Document`{#documents:document-6}](#document) object\'s
[[origin](https://dom.spec.whatwg.org/#concept-document-origin)]{#concept-document-origin
.dfn} is defined in DOM. It is initially set when the
[`Document`{#documents:document-7}](#document) object is created, and
can change during the lifetime of the
[`Document`{#documents:document-8}](#document) only upon setting
[`document.domain`{#documents:dom-document-domain}](browsers.html#dom-document-domain).
A [`Document`{#documents:document-9}](#document)\'s
[origin](https://dom.spec.whatwg.org/#concept-document-origin){#documents:concept-document-origin
x-internal="concept-document-origin"} can differ from the
[origin](https://url.spec.whatwg.org/#concept-url-origin){#documents:concept-url-origin
x-internal="concept-url-origin"} of its
[URL](https://dom.spec.whatwg.org/#concept-document-url){#documents:the-document's-address-2
x-internal="the-document's-address"}; for example when a [child
navigable](document-sequences.html#child-navigable){#documents:child-navigable}
is
[created](document-sequences.html#create-a-new-child-navigable){#documents:create-a-new-child-navigable},
its [active
document](document-sequences.html#nav-document){#documents:nav-document}\'s
[origin](https://dom.spec.whatwg.org/#concept-document-origin){#documents:concept-document-origin-2
x-internal="concept-document-origin"} is inherited from its
[parent](document-sequences.html#nav-parent){#documents:nav-parent}\'s
[active
document](document-sequences.html#nav-document){#documents:nav-document-2}\'s
[origin](https://dom.spec.whatwg.org/#concept-document-origin){#documents:concept-document-origin-3
x-internal="concept-document-origin"}, even though its [active
document](document-sequences.html#nav-document){#documents:nav-document-3}\'s
[URL](https://dom.spec.whatwg.org/#concept-document-url){#documents:the-document's-address-3
x-internal="the-document's-address"} is
[`about:blank`{#documents:about:blank}](infrastructure.html#about:blank).
[\[DOM\]](references.html#refsDOM)

When a [`Document`{#documents:document-10}](#document) is created by a
[script](webappapis.html#concept-script){#documents:concept-script}
using the
[`createDocument()`{#documents:dom-domimplementation-createdocument}](https://dom.spec.whatwg.org/#dom-domimplementation-createdocument){x-internal="dom-domimplementation-createdocument"}
or
[`createHTMLDocument()`{#documents:dom-domimplementation-createhtmldocument}](https://dom.spec.whatwg.org/#dom-domimplementation-createhtmldocument){x-internal="dom-domimplementation-createhtmldocument"}
methods, the [`Document`{#documents:document-11}](#document) is [ready
for post-load
tasks](parsing.html#ready-for-post-load-tasks){#documents:ready-for-post-load-tasks}
immediately.

[The document\'s referrer]{#the-document's-referrer .dfn export=""} is a
string (representing a
[URL](https://url.spec.whatwg.org/#concept-url){#documents:url-2
x-internal="url"}) that can be set when the
[`Document`{#documents:document-12}](#document) is created. If it is not
explicitly set, then its value is the empty string.

#### [3.1.1]{.secno} The [`Document`{#the-document-object:document}](#document) object[](#the-document-object){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Document](https://developer.mozilla.org/en-US/docs/Web/API/Document "The Document interface represents any web page loaded in the browser and serves as an entry point into the web page's content, which is the DOM tree.")

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
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

DOM defines a
[`Document`{#the-document-object:dom-document}](https://dom.spec.whatwg.org/#interface-document){x-internal="dom-document"}
interface, which this specification extends significantly.

``` idl
enum DocumentReadyState { "loading", "interactive", "complete" };
enum DocumentVisibilityState { "visible", "hidden" };
typedef (HTMLScriptElement or SVGScriptElement) HTMLOrSVGScriptElement;

[LegacyOverrideBuiltIns]
partial interface Document {
  static Document parseHTMLUnsafe((TrustedHTML or DOMString) html);

  // resource metadata management
  [PutForwards=href, LegacyUnforgeable] readonly attribute Location? location;
  attribute USVString domain;
  readonly attribute USVString referrer;
  attribute USVString cookie;
  readonly attribute DOMString lastModified;
  readonly attribute DocumentReadyState readyState;

  // DOM tree accessors
  getter object (DOMString name);
  [CEReactions] attribute DOMString title;
  [CEReactions] attribute DOMString dir;
  [CEReactions] attribute HTMLElement? body;
  readonly attribute HTMLHeadElement? head;
  [SameObject] readonly attribute HTMLCollection images;
  [SameObject] readonly attribute HTMLCollection embeds;
  [SameObject] readonly attribute HTMLCollection plugins;
  [SameObject] readonly attribute HTMLCollection links;
  [SameObject] readonly attribute HTMLCollection forms;
  [SameObject] readonly attribute HTMLCollection scripts;
  NodeList getElementsByName(DOMString elementName);
  readonly attribute HTMLOrSVGScriptElement? currentScript; // classic scripts in a document tree only

  // dynamic markup insertion
  [CEReactions] Document open(optional DOMString unused1, optional DOMString unused2); // both arguments are ignored
  WindowProxy? open(USVString url, DOMString name, DOMString features);
  [CEReactions] undefined close();
  [CEReactions] undefined write((TrustedHTML or DOMString)... text);
  [CEReactions] undefined writeln((TrustedHTML or DOMString)... text);

  // user interaction
  readonly attribute WindowProxy? defaultView;
  boolean hasFocus();
  [CEReactions] attribute DOMString designMode;
  [CEReactions] boolean execCommand(DOMString commandId, optional boolean showUI = false, optional DOMString value = "");
  boolean queryCommandEnabled(DOMString commandId);
  boolean queryCommandIndeterm(DOMString commandId);
  boolean queryCommandState(DOMString commandId);
  boolean queryCommandSupported(DOMString commandId);
  DOMString queryCommandValue(DOMString commandId);
  readonly attribute boolean hidden;
  readonly attribute DocumentVisibilityState visibilityState;

  // special event handler IDL attributes that only apply to Document objects
  [LegacyLenientThis] attribute EventHandler onreadystatechange;
  attribute EventHandler onvisibilitychange;

  // also has obsolete members
};
Document includes GlobalEventHandlers;
```

Each [`Document`{#the-document-object:document-5}](#document) has a
[policy container]{#concept-document-policy-container .dfn
dfn-for="Document" export=""} (a [policy
container](browsers.html#policy-container){#the-document-object:policy-container}),
initially a new policy container, which contains policies which apply to
the [`Document`{#the-document-object:document-6}](#document).

Each [`Document`{#the-document-object:document-7}](#document) has a
[permissions policy]{#concept-document-permissions-policy .dfn
dfn-for="Document" export=""}, which is a [permissions
policy](https://w3c.github.io/webappsec-feature-policy/#permissions-policy){#the-document-object:concept-permissions-policy
x-internal="concept-permissions-policy"}, which is initially empty.

Each [`Document`{#the-document-object:document-8}](#document) has a
[module map]{#concept-document-module-map .dfn}, which is a [module
map](webappapis.html#module-map){#the-document-object:module-map},
initially empty.

Each [`Document`{#the-document-object:document-9}](#document) has an
[opener policy]{#concept-document-coop .dfn}, which is an [opener
policy](browsers.html#cross-origin-opener-policy){#the-document-object:cross-origin-opener-policy},
initially a new opener policy.

Each [`Document`{#the-document-object:document-10}](#document) has an
[is initial `about:blank`]{#is-initial-about:blank .dfn}, which is a
boolean, initially false.

Each [`Document`{#the-document-object:document-11}](#document) has a
[during-loading navigation ID for WebDriver
BiDi]{#concept-document-navigation-id .dfn}, which is a [navigation
ID](browsing-the-web.html#navigation-id){#the-document-object:navigation-id}
or null, initially null.

As the name indicates, this is used for interfacing with the WebDriver
BiDi specification, which needs to be informed about certain occurrences
during the early parts of the
[`Document`{#the-document-object:document-12}](#document)\'s lifecycle,
in a way that ties them to the original [navigation
ID](browsing-the-web.html#navigation-id){#the-document-object:navigation-id-2}
used when the navigation that created this
[`Document`{#the-document-object:document-13}](#document) was the
[ongoing
navigation](browsing-the-web.html#ongoing-navigation){#the-document-object:ongoing-navigation}.
This eventually gets set back to null, after WebDriver BiDi considers
the loading process to be finished. [\[BIDI\]](references.html#refsBIDI)

Each [`Document`{#the-document-object:document-14}](#document) has an
[about base URL]{#concept-document-about-base-url .dfn}, which is a
[URL](https://url.spec.whatwg.org/#concept-url){#the-document-object:url
x-internal="url"} or null, initially null.

This is only populated for \"`about:`\"-schemed
[`Document`{#the-document-object:document-15}](#document)s.

Each [`Document`{#the-document-object:document-16}](#document) has a
[bfcache blocking details]{#concept-document-bfcache-blocking-details
.dfn}, which is a
[set](https://infra.spec.whatwg.org/#ordered-set){#the-document-object:set
x-internal="set"} of [not restored reason
details](nav-history-apis.html#nrr-details-struct){#the-document-object:nrr-details-struct},
initially empty.

Each [`Document`{#the-document-object:document-17}](#document) has an
[open dialogs list]{#open-dialogs-list .dfn}, which is a
[list](https://infra.spec.whatwg.org/#list){#the-document-object:list
x-internal="list"} of
[`dialog`{#the-document-object:the-dialog-element}](interactive-elements.html#the-dialog-element)
elements, initially empty.

#### [3.1.2]{.secno} The [`DocumentOrShadowRoot`{#the-documentorshadowroot-interface:documentorshadowroot}](#documentorshadowroot) interface[](#the-documentorshadowroot-interface){.self-link}

DOM defines the
[`DocumentOrShadowRoot`{#the-documentorshadowroot-interface:dom-documentorshadowroot}](https://dom.spec.whatwg.org/#documentorshadowroot){x-internal="dom-documentorshadowroot"}
mixin, which this specification extends.

``` idl
partial interface mixin DocumentOrShadowRoot {
  readonly attribute Element? activeElement;
};
```

#### [3.1.3]{.secno} [Resource metadata management]{.dfn}[](#resource-metadata-management){.self-link}

`document`{.variable}`.`[`referrer`](#dom-document-referrer){#dom-document-referrer-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Document/referrer](https://developer.mozilla.org/en-US/docs/Web/API/Document/referrer "The Document.referrer property returns the URI of the page that linked to this page.")

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

Returns the
[URL](https://dom.spec.whatwg.org/#concept-document-url){#resource-metadata-management:the-document's-address
x-internal="the-document's-address"} of the
[`Document`{#resource-metadata-management:document}](#document) from
which the user navigated to this one, unless it was blocked or there was
no such document, in which case it returns the empty string.

The
[`noreferrer`{#resource-metadata-management:link-type-noreferrer}](links.html#link-type-noreferrer)
link type can be used to block the referrer.

The [`referrer`]{#dom-document-referrer .dfn dfn-for="Document"
dfn-type="attribute"} attribute must return [the document\'s
referrer](#the-document's-referrer){#resource-metadata-management:the-document's-referrer}.

------------------------------------------------------------------------

`document`{.variable}`.`[`cookie`](#dom-document-cookie){#dom-document-cookie-dev}` [ = ``value`{.variable}` ]`

:   Returns the HTTP cookies that apply to the
    [`Document`{#resource-metadata-management:document-2}](#document).
    If there are no cookies or cookies can\'t be applied to this
    resource, the empty string will be returned.

    Can be set, to add a new cookie to the element\'s set of HTTP
    cookies.

    If the contents are [sandboxed into an opaque
    origin](browsers.html#sandboxed-origin-browsing-context-flag){#resource-metadata-management:sandboxed-origin-browsing-context-flag}
    (e.g., in an
    [`iframe`{#resource-metadata-management:the-iframe-element}](iframe-embed-object.html#the-iframe-element)
    with the
    [`sandbox`{#resource-metadata-management:attr-iframe-sandbox}](iframe-embed-object.html#attr-iframe-sandbox)
    attribute), a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#resource-metadata-management:securityerror
    x-internal="securityerror"}
    [`DOMException`{#resource-metadata-management:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
    will be thrown on getting and setting.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Document/cookie](https://developer.mozilla.org/en-US/docs/Web/API/Document/cookie "The Document property cookie lets you read and write cookies associated with the document. It serves as a getter and setter for the actual values of the cookies.")

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

The [`cookie`]{#dom-document-cookie .dfn dfn-for="Document"
dfn-type="attribute"} attribute represents the cookies of the resource
identified by the document\'s
[URL](https://dom.spec.whatwg.org/#concept-document-url){#resource-metadata-management:the-document's-address-2
x-internal="the-document's-address"}.

A [`Document`{#resource-metadata-management:document-3}](#document)
object that falls into one of the following conditions is a
[cookie-averse `Document` object]{#cookie-averse-document-object .dfn}:

- A [`Document`{#resource-metadata-management:document-4}](#document)
  object whose [browsing
  context](document-sequences.html#concept-document-bc){#resource-metadata-management:concept-document-bc}
  is null.
- A [`Document`{#resource-metadata-management:document-5}](#document)
  whose
  [URL](https://dom.spec.whatwg.org/#concept-document-url){#resource-metadata-management:the-document's-address-3
  x-internal="the-document's-address"}\'s
  [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#resource-metadata-management:concept-url-scheme
  x-internal="concept-url-scheme"} is not an [HTTP(S)
  scheme](https://fetch.spec.whatwg.org/#http-scheme){#resource-metadata-management:http(s)-scheme
  x-internal="http(s)-scheme"}.

[![(This is a tracking
vector.)](https://resources.whatwg.org/tracking-vector.svg){width="46"
crossorigin=""
height="64"}](https://infra.spec.whatwg.org/#tracking-vector "There is a tracking vector here."){#resource-metadata-management:tracking-vector
.tracking-vector x-internal="tracking-vector"} On getting, if the
document is a [cookie-averse `Document`
object](#cookie-averse-document-object){#resource-metadata-management:cookie-averse-document-object},
then the user agent must return the empty string. Otherwise, if the
[`Document`{#resource-metadata-management:document-6}](#document)\'s
[origin](https://dom.spec.whatwg.org/#concept-document-origin){#resource-metadata-management:concept-document-origin
x-internal="concept-document-origin"} is an [opaque
origin](browsers.html#concept-origin-opaque){#resource-metadata-management:concept-origin-opaque},
the user agent must throw a
[\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#resource-metadata-management:securityerror-2
x-internal="securityerror"}
[`DOMException`{#resource-metadata-management:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.
Otherwise, the user agent must return the
[cookie-string](https://httpwg.org/specs/rfc6265.html#sane-cookie-syntax){#resource-metadata-management:cookie-string
x-internal="cookie-string"} for the document\'s
[URL](https://dom.spec.whatwg.org/#concept-document-url){#resource-metadata-management:the-document's-address-4
x-internal="the-document's-address"} for a \"non-HTTP\" API, decoded
using [UTF-8 decode without
BOM](https://encoding.spec.whatwg.org/#utf-8-decode-without-bom){#resource-metadata-management:utf-8-decode-without-bom
x-internal="utf-8-decode-without-bom"}.
[\[COOKIES\]](references.html#refsCOOKIES)

On setting, if the document is a [cookie-averse `Document`
object](#cookie-averse-document-object){#resource-metadata-management:cookie-averse-document-object-2},
then the user agent must do nothing. Otherwise, if the
[`Document`{#resource-metadata-management:document-7}](#document)\'s
[origin](https://dom.spec.whatwg.org/#concept-document-origin){#resource-metadata-management:concept-document-origin-2
x-internal="concept-document-origin"} is an [opaque
origin](browsers.html#concept-origin-opaque){#resource-metadata-management:concept-origin-opaque-2},
the user agent must throw a
[\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#resource-metadata-management:securityerror-3
x-internal="securityerror"}
[`DOMException`{#resource-metadata-management:domexception-3}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.
Otherwise, the user agent must act as it would when [receiving a
set-cookie-string](https://httpwg.org/specs/rfc6265.html#storage-model){#resource-metadata-management:receives-a-set-cookie-string
x-internal="receives-a-set-cookie-string"} for the document\'s
[URL](https://dom.spec.whatwg.org/#concept-document-url){#resource-metadata-management:the-document's-address-5
x-internal="the-document's-address"} via a \"non-HTTP\" API, consisting
of the new value [encoded as
UTF-8](https://encoding.spec.whatwg.org/#utf-8-encode){#resource-metadata-management:utf-8-encode
x-internal="utf-8-encode"}. [\[COOKIES\]](references.html#refsCOOKIES)
[\[ENCODING\]](references.html#refsENCODING)

Since the
[`cookie`{#resource-metadata-management:dom-document-cookie}](#dom-document-cookie)
attribute is accessible across frames, the path restrictions on cookies
are only a tool to help manage which cookies are sent to which parts of
the site, and are not in any way a security feature.

The
[`cookie`{#resource-metadata-management:dom-document-cookie-2}](#dom-document-cookie)
attribute\'s getter and setter synchronously access shared state. Since
there is no locking mechanism, other browsing contexts in a multiprocess
user agent can modify cookies while scripts are running. A site could,
for instance, try to read a cookie, increment its value, then write it
back out, using the new value of the cookie as a unique identifier for
the session; if the site does this twice in two different browser
windows at the same time, it might end up using the same \"unique\"
identifier for both sessions, with potentially disastrous effects.

------------------------------------------------------------------------

`document`{.variable}`.`[`lastModified`](#dom-document-lastmodified){#dom-document-lastmodified-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Document/lastModified](https://developer.mozilla.org/en-US/docs/Web/API/Document/lastModified "The lastModified property of the Document interface returns a string containing the date and time on which the current document was last modified.")

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

Returns the date of the last modification to the document, as reported
by the server, in the form \"`MM/DD/YYYY hh:mm:ss`\", in the user\'s
local time zone.

If the last modification date is not known, the current time is returned
instead.

The [`lastModified`]{#dom-document-lastmodified .dfn dfn-for="Document"
dfn-type="attribute"} attribute, on getting, must return the date and
time of the
[`Document`{#resource-metadata-management:document-8}](#document)\'s
source file\'s last modification, in the user\'s local time zone, in the
following format:

1.  The month component of the date.

2.  A U+002F SOLIDUS character (/).

3.  The day component of the date.

4.  A U+002F SOLIDUS character (/).

5.  The year component of the date.

6.  A U+0020 SPACE character.

7.  The hours component of the time.

8.  A U+003A COLON character (:).

9.  The minutes component of the time.

10. A U+003A COLON character (:).

11. The seconds component of the time.

All the numeric components above, other than the year, must be given as
two [ASCII
digits](https://infra.spec.whatwg.org/#ascii-digit){#resource-metadata-management:ascii-digits
x-internal="ascii-digits"} representing the number in base ten,
zero-padded if necessary. The year must be given as the shortest
possible string of four or more [ASCII
digits](https://infra.spec.whatwg.org/#ascii-digit){#resource-metadata-management:ascii-digits-2
x-internal="ascii-digits"} representing the number in base ten,
zero-padded if necessary.

The [`Document`{#resource-metadata-management:document-9}](#document)\'s
source file\'s last modification date and time must be derived from
relevant features of the networking protocols used, e.g. from the value
of the HTTP
\`[`Last-Modified`{#resource-metadata-management:http-last-modified}](https://httpwg.org/specs/rfc7232.html#header.last-modified){x-internal="http-last-modified"}\`
header of the document, or from metadata in the file system for local
files. If the last modification date and time are not known, the
attribute must return the current date and time in the above format.

#### [3.1.4]{.secno} Reporting document loading status[](#reporting-document-loading-status){.self-link}

`document`{.variable}`.`[`readyState`](#dom-document-readystate){#dom-document-readystate-dev}

:   Returns \"`loading`\" while the
    [`Document`{#reporting-document-loading-status:document}](#document)
    is loading, \"`interactive`\" once it is finished parsing but still
    loading subresources, and \"`complete`\" once it has loaded.

    The
    [`readystatechange`{#reporting-document-loading-status:event-readystatechange}](indices.html#event-readystatechange)
    event fires on the
    [`Document`{#reporting-document-loading-status:document-2}](#document)
    object when this value changes.

    The
    [`DOMContentLoaded`{#reporting-document-loading-status:event-domcontentloaded}](indices.html#event-domcontentloaded)
    event fires after the transition to \"`interactive`\" but before the
    transition to \"`complete`\", at the point where all subresources
    apart from
    [`async`{#reporting-document-loading-status:attr-script-async}](scripting.html#attr-script-async)
    [`script`{#reporting-document-loading-status:the-script-element}](scripting.html#the-script-element)
    elements have loaded.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Document/readyState](https://developer.mozilla.org/en-US/docs/Web/API/Document/readyState "The Document.readyState property describes the loading state of the document. When the value of this property changes, a readystatechange event fires on the document object.")

Support in all current engines.

::: support
[Firefox3.6+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer11]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

Each
[`Document`{#reporting-document-loading-status:document-3}](#document)
has a [current document readiness]{#current-document-readiness .dfn}, a
string, initially \"`complete`\".

For
[`Document`{#reporting-document-loading-status:document-4}](#document)
objects created via the [create and initialize a `Document`
object](document-lifecycle.html#initialise-the-document-object){#reporting-document-loading-status:initialise-the-document-object}
algorithm, this will be immediately reset to \"`loading`\" before any
script can observe the value of
[`document.readyState`{#reporting-document-loading-status:dom-document-readystate}](#dom-document-readystate).
This default applies to other cases such as [initial
`about:blank`](#is-initial-about:blank){#reporting-document-loading-status:is-initial-about:blank}
[`Document`{#reporting-document-loading-status:document-5}](#document)s
or
[`Document`{#reporting-document-loading-status:document-6}](#document)s
without a [browsing
context](document-sequences.html#concept-document-bc){#reporting-document-loading-status:concept-document-bc}.

The [`readyState`]{#dom-document-readystate .dfn dfn-for="Document"
dfn-type="attribute"} getter steps are to return
[this](https://webidl.spec.whatwg.org/#this){#reporting-document-loading-status:this
x-internal="this"}\'s [current document
readiness](#current-document-readiness){#reporting-document-loading-status:current-document-readiness}.

To [update the current document
readiness]{#update-the-current-document-readiness .dfn} for
[`Document`{#reporting-document-loading-status:document-7}](#document)
`document`{.variable} to `readinessValue`{.variable}:

1.  If `document`{.variable}\'s [current document
    readiness](#current-document-readiness){#reporting-document-loading-status:current-document-readiness-2}
    equals `readinessValue`{.variable}, then return.

2.  Set `document`{.variable}\'s [current document
    readiness](#current-document-readiness){#reporting-document-loading-status:current-document-readiness-3}
    to `readinessValue`{.variable}.

3.  If `document`{.variable} is associated with an [HTML
    parser](parsing.html#html-parser){#reporting-document-loading-status:html-parser},
    then:

    1.  Let `now`{.variable} be the [current high resolution
        time](https://w3c.github.io/hr-time/#dfn-current-high-resolution-time){#reporting-document-loading-status:current-high-resolution-time
        x-internal="current-high-resolution-time"} given
        `document`{.variable}\'s [relevant global
        object](webappapis.html#concept-relevant-global){#reporting-document-loading-status:concept-relevant-global}.

    2.  If `readinessValue`{.variable} is \"`complete`\", and
        `document`{.variable}\'s [load timing
        info](#load-timing-info){#reporting-document-loading-status:load-timing-info}\'s
        [DOM complete
        time](#dom-complete-time){#reporting-document-loading-status:dom-complete-time}
        is 0, then set `document`{.variable}\'s [load timing
        info](#load-timing-info){#reporting-document-loading-status:load-timing-info-2}\'s
        [DOM complete
        time](#dom-complete-time){#reporting-document-loading-status:dom-complete-time-2}
        to `now`{.variable}.

    3.  Otherwise, if `readinessValue`{.variable} is \"`interactive`\",
        and `document`{.variable}\'s [load timing
        info](#load-timing-info){#reporting-document-loading-status:load-timing-info-3}\'s
        [DOM interactive
        time](#dom-interactive-time){#reporting-document-loading-status:dom-interactive-time}
        is 0, then set `document`{.variable}\'s [load timing
        info](#load-timing-info){#reporting-document-loading-status:load-timing-info-4}\'s
        [DOM interactive
        time](#dom-interactive-time){#reporting-document-loading-status:dom-interactive-time-2}
        to `now`{.variable}.

4.  [Fire an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#reporting-document-loading-status:concept-event-fire
    x-internal="concept-event-fire"} named
    [`readystatechange`{#reporting-document-loading-status:event-readystatechange-2}](indices.html#event-readystatechange)
    at `document`{.variable}.

------------------------------------------------------------------------

A [`Document`{#reporting-document-loading-status:document-8}](#document)
is said to have an [active parser]{#active-parser .dfn} if it is
associated with an [HTML
parser](parsing.html#html-parser){#reporting-document-loading-status:html-parser-2}
or an [XML
parser](xhtml.html#xml-parser){#reporting-document-loading-status:xml-parser}
that has not yet been
[stopped](parsing.html#stop-parsing){#reporting-document-loading-status:stop-parsing}
or
[aborted](parsing.html#abort-a-parser){#reporting-document-loading-status:abort-a-parser}.

------------------------------------------------------------------------

A [`Document`{#reporting-document-loading-status:document-9}](#document)
has a [document load timing
info](#document-load-timing-info){#reporting-document-loading-status:document-load-timing-info}
[load timing info]{#load-timing-info .dfn dfn-for="Document" export=""}.

A
[`Document`{#reporting-document-loading-status:document-10}](#document)
has a [document unload timing
info](#document-unload-timing-info){#reporting-document-loading-status:document-unload-timing-info}
[previous document unload timing]{#previous-document-unload-timing .dfn
dfn-for="Document" export=""}.

A
[`Document`{#reporting-document-loading-status:document-11}](#document)
has a boolean [was created via cross-origin
redirects]{#was-created-via-cross-origin-redirects .dfn export=""},
initially false.

The [document load timing info]{#document-load-timing-info .dfn
export=""}
[struct](https://infra.spec.whatwg.org/#struct){#reporting-document-loading-status:struct
x-internal="struct"} has the following
[items](https://infra.spec.whatwg.org/#struct-item){#reporting-document-loading-status:struct-item
x-internal="struct-item"}:

[navigation start time]{#navigation-start-time .dfn dfn-for="document load timing info" export=""} (default 0)
:   A number

[DOM interactive time]{#dom-interactive-time .dfn dfn-for="document load timing info" export=""} (default 0)\
[DOM content loaded event start time]{#dom-content-loaded-event-start-time .dfn dfn-for="document load timing info" export=""} (default 0)\
[DOM content loaded event end time]{#dom-content-loaded-event-end-time .dfn dfn-for="document load timing info" export=""} (default 0)\
[DOM complete time]{#dom-complete-time .dfn dfn-for="document load timing info" export=""} (default 0)\
[load event start time]{#load-event-start-time .dfn dfn-for="document load timing info" export=""} (default 0)\
[load event end time]{#load-event-end-time .dfn dfn-for="document load timing info" export=""} (default 0)
:   [`DOMHighResTimeStamp`{#reporting-document-loading-status:domhighrestimestamp}](https://w3c.github.io/hr-time/#dom-domhighrestimestamp){x-internal="domhighrestimestamp"}
    values

The [document unload timing info]{#document-unload-timing-info .dfn
export=""}
[struct](https://infra.spec.whatwg.org/#struct){#reporting-document-loading-status:struct-2
x-internal="struct"} has the following
[items](https://infra.spec.whatwg.org/#struct-item){#reporting-document-loading-status:struct-item-2
x-internal="struct-item"}:

[unload event start time]{#unload-event-start-time .dfn dfn-for="document unload timing info" export=""} (default 0)\
[unload event end time]{#unload-event-end-time .dfn dfn-for="document unload timing info" export=""} (default 0)
:   [`DOMHighResTimeStamp`{#reporting-document-loading-status:domhighrestimestamp-2}](https://w3c.github.io/hr-time/#dom-domhighrestimestamp){x-internal="domhighrestimestamp"}
    values

#### [3.1.5]{.secno} [Render-blocking mechanism]{.dfn}[](#render-blocking-mechanism){.self-link}

Each [`Document`{#render-blocking-mechanism:document}](#document) has a
[render-blocking element set]{#render-blocking-element-set .dfn}, a
[set](https://infra.spec.whatwg.org/#ordered-set){#render-blocking-mechanism:set
x-internal="set"} of elements, initially the empty set.

A [`Document`{#render-blocking-mechanism:document-2}](#document)
`document`{.variable} [allows adding render-blocking
elements]{#allows-adding-render-blocking-elements .dfn} if
`document`{.variable}\'s [content
type](https://dom.spec.whatwg.org/#concept-document-content-type){#render-blocking-mechanism:concept-document-content-type
x-internal="concept-document-content-type"} is
\"[`text/html`{#render-blocking-mechanism:text/html}](iana.html#text/html)\"
and [the body
element](#the-body-element-2){#render-blocking-mechanism:the-body-element-2}
of `document`{.variable} is null.

A [`Document`{#render-blocking-mechanism:document-3}](#document)
`document`{.variable} is [render-blocked]{#render-blocked .dfn} if both
of the following are true:

- `document`{.variable}\'s [render-blocking element
  set](#render-blocking-element-set){#render-blocking-mechanism:render-blocking-element-set}
  is non-empty, or `document`{.variable} [allows adding render-blocking
  elements](#allows-adding-render-blocking-elements){#render-blocking-mechanism:allows-adding-render-blocking-elements}.

- The [current high resolution
  time](https://w3c.github.io/hr-time/#dfn-current-high-resolution-time){#render-blocking-mechanism:current-high-resolution-time
  x-internal="current-high-resolution-time"} given
  `document`{.variable}\'s [relevant global
  object](webappapis.html#concept-relevant-global){#render-blocking-mechanism:concept-relevant-global}
  has not exceeded an
  [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#render-blocking-mechanism:implementation-defined
  x-internal="implementation-defined"} timeout value.

An element `el`{.variable} is [render-blocking]{#render-blocking .dfn}
if `el`{.variable}\'s [node
document](https://dom.spec.whatwg.org/#concept-node-document){#render-blocking-mechanism:node-document
x-internal="node-document"} `document`{.variable} is
[render-blocked](#render-blocked){#render-blocking-mechanism:render-blocked},
and `el`{.variable} is in `document`{.variable}\'s [render-blocking
element
set](#render-blocking-element-set){#render-blocking-mechanism:render-blocking-element-set-2}.

To [block rendering]{#block-rendering .dfn} on an element
`el`{.variable}:

1.  Let `document`{.variable} be `el`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#render-blocking-mechanism:node-document-2
    x-internal="node-document"}.

2.  If `document`{.variable} [allows adding render-blocking
    elements](#allows-adding-render-blocking-elements){#render-blocking-mechanism:allows-adding-render-blocking-elements-2},
    then
    [append](https://infra.spec.whatwg.org/#set-append){#render-blocking-mechanism:set-append
    x-internal="set-append"} `el`{.variable} to `document`{.variable}\'s
    [render-blocking element
    set](#render-blocking-element-set){#render-blocking-mechanism:render-blocking-element-set-3}.

To [unblock rendering]{#unblock-rendering .dfn} on an element
`el`{.variable}:

1.  Let `document`{.variable} be `el`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#render-blocking-mechanism:node-document-3
    x-internal="node-document"}.

2.  [Remove](https://infra.spec.whatwg.org/#list-remove){#render-blocking-mechanism:list-remove
    x-internal="list-remove"} `el`{.variable} from
    `document`{.variable}\'s [render-blocking element
    set](#render-blocking-element-set){#render-blocking-mechanism:render-blocking-element-set-4}.

Whenever a
[render-blocking](#render-blocking){#render-blocking-mechanism:render-blocking}
element `el`{.variable} [becomes browsing-context
disconnected](infrastructure.html#becomes-browsing-context-disconnected){#render-blocking-mechanism:becomes-browsing-context-disconnected},
or `el`{.variable}\'s [blocking
attribute](urls-and-fetching.html#blocking-attribute){#render-blocking-mechanism:blocking-attribute}\'s
value is changed so that `el`{.variable} is no longer [potentially
render-blocking](urls-and-fetching.html#potentially-render-blocking){#render-blocking-mechanism:potentially-render-blocking},
then [unblock
rendering](#unblock-rendering){#render-blocking-mechanism:unblock-rendering}
on `el`{.variable}.

#### [3.1.6]{.secno} [DOM tree accessors]{.dfn}[](#dom-tree-accessors){.self-link}

[The `html` element]{#the-html-element-2 .dfn} of a document is its
[document
element](https://dom.spec.whatwg.org/#document-element){#dom-tree-accessors:document-element
x-internal="document-element"}, if it\'s an
[`html`{#dom-tree-accessors:the-html-element}](semantics.html#the-html-element)
element, and null otherwise.

------------------------------------------------------------------------

`document`{.variable}`.`[`head`](#dom-document-head){#dom-document-head-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Document/head](https://developer.mozilla.org/en-US/docs/Web/API/Document/head "The head read-only property of the Document interface returns the <head> element of the current document.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS4+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

Returns [the `head`
element](#the-head-element-2){#dom-tree-accessors:the-head-element-2}.

[The `head` element]{#the-head-element-2 .dfn} of a document is the
first
[`head`{#dom-tree-accessors:the-head-element}](semantics.html#the-head-element)
element that is a child of [the `html`
element](#the-html-element-2){#dom-tree-accessors:the-html-element-2},
if there is one, or null otherwise.

The [`head`]{#dom-document-head .dfn dfn-for="Document"
dfn-type="attribute"} attribute, on getting, must return [the `head`
element](#the-head-element-2){#dom-tree-accessors:the-head-element-2-2}
of the document (a
[`head`{#dom-tree-accessors:the-head-element-3}](semantics.html#the-head-element)
element or null).

------------------------------------------------------------------------

`document`{.variable}`.`[`title`](#document.title){#dom-document-title-dev}` [ = ``value`{.variable}` ]`

:   Returns the document\'s title, as given by [the `title`
    element](#the-title-element-2){#dom-tree-accessors:the-title-element-2}
    for HTML and as given by the [SVG
    `title`](https://svgwg.org/svg2-draft/struct.html#TitleElement){#dom-tree-accessors:svg-title
    x-internal="svg-title"} element for SVG.

    Can be set, to update the document\'s title. If there is no
    appropriate element to update, the new value is ignored.

[The `title` element]{#the-title-element-2 .dfn} of a document is the
first
[`title`{#dom-tree-accessors:the-title-element}](semantics.html#the-title-element)
element in the document (in [tree
order](https://dom.spec.whatwg.org/#concept-tree-order){#dom-tree-accessors:tree-order
x-internal="tree-order"}), if there is one, or null otherwise.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Document/title](https://developer.mozilla.org/en-US/docs/Web/API/Document/title "The document.title property gets or sets the current title of the document. When present, it defaults to the value of the <title>.")

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

The [`title`]{#document.title .dfn dfn-for="Document"
dfn-type="attribute"} attribute must, on getting, run the following
algorithm:

1.  If the [document
    element](https://dom.spec.whatwg.org/#document-element){#dom-tree-accessors:document-element-2
    x-internal="document-element"} is an [SVG
    `svg`](https://svgwg.org/svg2-draft/struct.html#SVGElement){#dom-tree-accessors:svg-svg
    x-internal="svg-svg"} element, then let `value`{.variable} be the
    [child text
    content](https://dom.spec.whatwg.org/#concept-child-text-content){#dom-tree-accessors:child-text-content
    x-internal="child-text-content"} of the first [SVG
    `title`](https://svgwg.org/svg2-draft/struct.html#TitleElement){#dom-tree-accessors:svg-title-2
    x-internal="svg-title"} element that is a child of the [document
    element](https://dom.spec.whatwg.org/#document-element){#dom-tree-accessors:document-element-3
    x-internal="document-element"}.

2.  Otherwise, let `value`{.variable} be the [child text
    content](https://dom.spec.whatwg.org/#concept-child-text-content){#dom-tree-accessors:child-text-content-2
    x-internal="child-text-content"} of [the `title`
    element](#the-title-element-2){#dom-tree-accessors:the-title-element-2-2},
    or the empty string if [the `title`
    element](#the-title-element-2){#dom-tree-accessors:the-title-element-2-3}
    is null.

3.  [Strip and collapse ASCII
    whitespace](https://infra.spec.whatwg.org/#strip-and-collapse-ascii-whitespace){#dom-tree-accessors:strip-and-collapse-ascii-whitespace
    x-internal="strip-and-collapse-ascii-whitespace"} in
    `value`{.variable}.

4.  Return `value`{.variable}.

On setting, the steps corresponding to the first matching condition in
the following list must be run:

If the [document element](https://dom.spec.whatwg.org/#document-element){#dom-tree-accessors:document-element-4 x-internal="document-element"} is an [SVG `svg`](https://svgwg.org/svg2-draft/struct.html#SVGElement){#dom-tree-accessors:svg-svg-2 x-internal="svg-svg"} element

:   1.  If there is an [SVG
        `title`](https://svgwg.org/svg2-draft/struct.html#TitleElement){#dom-tree-accessors:svg-title-3
        x-internal="svg-title"} element that is a child of the [document
        element](https://dom.spec.whatwg.org/#document-element){#dom-tree-accessors:document-element-5
        x-internal="document-element"}, let `element`{.variable} be the
        first such element.

    2.  Otherwise:

        1.  Let `element`{.variable} be the result of [creating an
            element](https://dom.spec.whatwg.org/#concept-create-element){#dom-tree-accessors:create-an-element
            x-internal="create-an-element"} given the [document
            element](https://dom.spec.whatwg.org/#document-element){#dom-tree-accessors:document-element-6
            x-internal="document-element"}\'s [node
            document](https://dom.spec.whatwg.org/#concept-node-document){#dom-tree-accessors:node-document
            x-internal="node-document"}, \"`title`\", and the [SVG
            namespace](https://infra.spec.whatwg.org/#svg-namespace){#dom-tree-accessors:svg-namespace
            x-internal="svg-namespace"}.

        2.  Insert `element`{.variable} as the [first
            child](https://dom.spec.whatwg.org/#concept-tree-first-child){#dom-tree-accessors:first-child
            x-internal="first-child"} of the [document
            element](https://dom.spec.whatwg.org/#document-element){#dom-tree-accessors:document-element-7
            x-internal="document-element"}.

    3.  [String replace
        all](https://dom.spec.whatwg.org/#string-replace-all){#dom-tree-accessors:string-replace-all
        x-internal="string-replace-all"} with the given value within
        `element`{.variable}.

If the [document element](https://dom.spec.whatwg.org/#document-element){#dom-tree-accessors:document-element-8 x-internal="document-element"} is in the [HTML namespace](https://infra.spec.whatwg.org/#html-namespace){#dom-tree-accessors:html-namespace-2 x-internal="html-namespace-2"}

:   1.  If [the `title`
        element](#the-title-element-2){#dom-tree-accessors:the-title-element-2-4}
        is null and [the `head`
        element](#the-head-element-2){#dom-tree-accessors:the-head-element-2-3}
        is null, then return.

    2.  If [the `title`
        element](#the-title-element-2){#dom-tree-accessors:the-title-element-2-5}
        is non-null, let `element`{.variable} be [the `title`
        element](#the-title-element-2){#dom-tree-accessors:the-title-element-2-6}.

    3.  Otherwise:

        1.  Let `element`{.variable} be the result of [creating an
            element](https://dom.spec.whatwg.org/#concept-create-element){#dom-tree-accessors:create-an-element-2
            x-internal="create-an-element"} given the [document
            element](https://dom.spec.whatwg.org/#document-element){#dom-tree-accessors:document-element-9
            x-internal="document-element"}\'s [node
            document](https://dom.spec.whatwg.org/#concept-node-document){#dom-tree-accessors:node-document-2
            x-internal="node-document"}, \"`title`\", and the [HTML
            namespace](https://infra.spec.whatwg.org/#html-namespace){#dom-tree-accessors:html-namespace-2-2
            x-internal="html-namespace-2"}.

        2.  [Append](https://dom.spec.whatwg.org/#concept-node-append){#dom-tree-accessors:concept-node-append
            x-internal="concept-node-append"} `element`{.variable} to
            [the `head`
            element](#the-head-element-2){#dom-tree-accessors:the-head-element-2-4}.

    4.  [String replace
        all](https://dom.spec.whatwg.org/#string-replace-all){#dom-tree-accessors:string-replace-all-2
        x-internal="string-replace-all"} with the given value within
        `element`{.variable}.

Otherwise

:   Do nothing.

------------------------------------------------------------------------

`document`{.variable}`.`[`body`](#dom-document-body){#dom-document-body-dev}` [ = ``value`{.variable}` ]`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Document/body](https://developer.mozilla.org/en-US/docs/Web/API/Document/body "The Document.body property represents the <body> or <frameset> node of the current document, or null if no such element exists.")

Support in all current engines.

::: support
[Firefox60+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera9.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

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

Returns [the body
element](#the-body-element-2){#dom-tree-accessors:the-body-element-2}.

Can be set, to replace [the body
element](#the-body-element-2){#dom-tree-accessors:the-body-element-2-2}.

If the new value is not a
[`body`{#dom-tree-accessors:the-body-element}](sections.html#the-body-element)
or [`frameset`{#dom-tree-accessors:frameset}](obsolete.html#frameset)
element, this will throw a
[\"`HierarchyRequestError`\"](https://webidl.spec.whatwg.org/#hierarchyrequesterror){#dom-tree-accessors:hierarchyrequesterror
x-internal="hierarchyrequesterror"}
[`DOMException`{#dom-tree-accessors:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

[The body element]{#the-body-element-2 .dfn export=""} of a document is
the first of [the `html`
element](#the-html-element-2){#dom-tree-accessors:the-html-element-2-2}\'s
children that is either a
[`body`{#dom-tree-accessors:the-body-element-3}](sections.html#the-body-element)
element or a
[`frameset`{#dom-tree-accessors:frameset-2}](obsolete.html#frameset)
element, or null if there is no such element.

The [`body`]{#dom-document-body .dfn dfn-for="Document"
dfn-type="attribute"} attribute, on getting, must return [the body
element](#the-body-element-2){#dom-tree-accessors:the-body-element-2-3}
of the document (either a
[`body`{#dom-tree-accessors:the-body-element-4}](sections.html#the-body-element)
element, a
[`frameset`{#dom-tree-accessors:frameset-3}](obsolete.html#frameset)
element, or null). On setting, the following algorithm must be run:

1.  If the new value is not a
    [`body`{#dom-tree-accessors:the-body-element-5}](sections.html#the-body-element)
    or
    [`frameset`{#dom-tree-accessors:frameset-4}](obsolete.html#frameset)
    element, then throw a
    [\"`HierarchyRequestError`\"](https://webidl.spec.whatwg.org/#hierarchyrequesterror){#dom-tree-accessors:hierarchyrequesterror-2
    x-internal="hierarchyrequesterror"}
    [`DOMException`{#dom-tree-accessors:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.
2.  Otherwise, if the new value is the same as [the body
    element](#the-body-element-2){#dom-tree-accessors:the-body-element-2-4},
    return.
3.  Otherwise, if [the body
    element](#the-body-element-2){#dom-tree-accessors:the-body-element-2-5}
    is not null, then
    [replace](https://dom.spec.whatwg.org/#concept-node-replace){#dom-tree-accessors:concept-node-replace
    x-internal="concept-node-replace"} [the body
    element](#the-body-element-2){#dom-tree-accessors:the-body-element-2-6}
    with the new value within [the body
    element](#the-body-element-2){#dom-tree-accessors:the-body-element-2-7}\'s
    parent and return.
4.  Otherwise, if there is no [document
    element](https://dom.spec.whatwg.org/#document-element){#dom-tree-accessors:document-element-10
    x-internal="document-element"}, throw a
    [\"`HierarchyRequestError`\"](https://webidl.spec.whatwg.org/#hierarchyrequesterror){#dom-tree-accessors:hierarchyrequesterror-3
    x-internal="hierarchyrequesterror"}
    [`DOMException`{#dom-tree-accessors:domexception-3}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.
5.  Otherwise, [the body
    element](#the-body-element-2){#dom-tree-accessors:the-body-element-2-8}
    is null, but there\'s a [document
    element](https://dom.spec.whatwg.org/#document-element){#dom-tree-accessors:document-element-11
    x-internal="document-element"}.
    [Append](https://dom.spec.whatwg.org/#concept-node-append){#dom-tree-accessors:concept-node-append-2
    x-internal="concept-node-append"} the new value to the [document
    element](https://dom.spec.whatwg.org/#document-element){#dom-tree-accessors:document-element-12
    x-internal="document-element"}.

The value returned by the
[`body`{#dom-tree-accessors:dom-document-body}](#dom-document-body)
getter is not always the one passed to the setter.

::: example
In this example, the setter successfully inserts a
[`body`{#dom-tree-accessors:the-body-element-6}](sections.html#the-body-element)
element (though this is non-conforming since SVG does not allow a
[`body`{#dom-tree-accessors:the-body-element-7}](sections.html#the-body-element)
as child of [SVG
`svg`](https://svgwg.org/svg2-draft/struct.html#SVGElement){#dom-tree-accessors:svg-svg-3
x-internal="svg-svg"}). However the getter will return null because the
document element is not
[`html`{#dom-tree-accessors:the-html-element-3}](semantics.html#the-html-element).

``` html
<svg xmlns="http://www.w3.org/2000/svg">
 <script>
  document.body = document.createElementNS("http://www.w3.org/1999/xhtml", "body");
  console.assert(document.body === null);
 </script>
</svg>
```
:::

------------------------------------------------------------------------

`document`{.variable}`.`[`images`](#dom-document-images){#dom-document-images-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Document/images](https://developer.mozilla.org/en-US/docs/Web/API/Document/images "The images read-only property of the Document interface returns a collection of the images in the current HTML document.")

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

Returns an
[`HTMLCollection`{#dom-tree-accessors:htmlcollection}](https://dom.spec.whatwg.org/#interface-htmlcollection){x-internal="htmlcollection"}
of the
[`img`{#dom-tree-accessors:the-img-element}](embedded-content.html#the-img-element)
elements in the [`Document`{#dom-tree-accessors:document}](#document).

`document`{.variable}`.`[`embeds`](#dom-document-embeds){#dom-document-embeds-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Document/embeds](https://developer.mozilla.org/en-US/docs/Web/API/Document/embeds "The embeds read-only property of the Document interface returns a list of the embedded <embed> elements within the current document.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari10.1+]{.safari .yes}[Chrome64+]{.chrome
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

`document`{.variable}`.`[`plugins`](#dom-document-plugins){#dom-document-plugins-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Document/plugins](https://developer.mozilla.org/en-US/docs/Web/API/Document/plugins "The plugins read-only property of the Document interface returns an HTMLCollection object containing one or more HTMLEmbedElements representing the <embed> elements in the current document.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari10.1+]{.safari .yes}[Chrome64+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera51+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer4+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari
iOS10.3+]{.safari_ios .yes}[Chrome Android?]{.chrome_android
.unknown}[WebView Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android47+]{.opera_android .yes}
:::
::::
:::::

Returns an
[`HTMLCollection`{#dom-tree-accessors:htmlcollection-2}](https://dom.spec.whatwg.org/#interface-htmlcollection){x-internal="htmlcollection"}
of the
[`embed`{#dom-tree-accessors:the-embed-element}](iframe-embed-object.html#the-embed-element)
elements in the [`Document`{#dom-tree-accessors:document-2}](#document).

`document`{.variable}`.`[`links`](#dom-document-links){#dom-document-links-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Document/links](https://developer.mozilla.org/en-US/docs/Web/API/Document/links "The links read-only property of the Document interface returns a collection of all <area> elements and <a> elements in a document with a value for the href attribute.")

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

Returns an
[`HTMLCollection`{#dom-tree-accessors:htmlcollection-3}](https://dom.spec.whatwg.org/#interface-htmlcollection){x-internal="htmlcollection"}
of the
[`a`{#dom-tree-accessors:the-a-element}](text-level-semantics.html#the-a-element)
and
[`area`{#dom-tree-accessors:the-area-element}](image-maps.html#the-area-element)
elements in the [`Document`{#dom-tree-accessors:document-3}](#document)
that have
[`href`{#dom-tree-accessors:attr-hyperlink-href}](links.html#attr-hyperlink-href)
attributes.

`document`{.variable}`.`[`forms`](#dom-document-forms){#dom-document-forms-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Document/forms](https://developer.mozilla.org/en-US/docs/Web/API/Document/forms "The forms read-only property of the Document interface returns an HTMLCollection listing all the <form> elements contained in the document.")

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

Returns an
[`HTMLCollection`{#dom-tree-accessors:htmlcollection-4}](https://dom.spec.whatwg.org/#interface-htmlcollection){x-internal="htmlcollection"}
of the
[`form`{#dom-tree-accessors:the-form-element}](forms.html#the-form-element)
elements in the [`Document`{#dom-tree-accessors:document-4}](#document).

`document`{.variable}`.`[`scripts`](#dom-document-scripts){#dom-document-scripts-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Document/scripts](https://developer.mozilla.org/en-US/docs/Web/API/Document/scripts "The scripts property of the Document interface returns a list of the <script> elements in the document. The returned object is an HTMLCollection.")

Support in all current engines.

::: support
[Firefox9+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer4+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns an
[`HTMLCollection`{#dom-tree-accessors:htmlcollection-5}](https://dom.spec.whatwg.org/#interface-htmlcollection){x-internal="htmlcollection"}
of the
[`script`{#dom-tree-accessors:the-script-element}](scripting.html#the-script-element)
elements in the [`Document`{#dom-tree-accessors:document-5}](#document).

The [`images`]{#dom-document-images .dfn dfn-for="Document"
dfn-type="attribute"} attribute must return an
[`HTMLCollection`{#dom-tree-accessors:htmlcollection-6}](https://dom.spec.whatwg.org/#interface-htmlcollection){x-internal="htmlcollection"}
rooted at the [`Document`{#dom-tree-accessors:document-6}](#document)
node, whose filter matches only
[`img`{#dom-tree-accessors:the-img-element-2}](embedded-content.html#the-img-element)
elements.

The [`embeds`]{#dom-document-embeds .dfn dfn-for="Document"
dfn-type="attribute"} attribute must return an
[`HTMLCollection`{#dom-tree-accessors:htmlcollection-7}](https://dom.spec.whatwg.org/#interface-htmlcollection){x-internal="htmlcollection"}
rooted at the [`Document`{#dom-tree-accessors:document-7}](#document)
node, whose filter matches only
[`embed`{#dom-tree-accessors:the-embed-element-2}](iframe-embed-object.html#the-embed-element)
elements.

The [`plugins`]{#dom-document-plugins .dfn dfn-for="Document"
dfn-type="attribute"} attribute must return the same object as that
returned by the
[`embeds`{#dom-tree-accessors:dom-document-embeds}](#dom-document-embeds)
attribute.

The [`links`]{#dom-document-links .dfn dfn-for="Document"
dfn-type="attribute"} attribute must return an
[`HTMLCollection`{#dom-tree-accessors:htmlcollection-8}](https://dom.spec.whatwg.org/#interface-htmlcollection){x-internal="htmlcollection"}
rooted at the [`Document`{#dom-tree-accessors:document-8}](#document)
node, whose filter matches only
[`a`{#dom-tree-accessors:the-a-element-2}](text-level-semantics.html#the-a-element)
elements with
[`href`{#dom-tree-accessors:attr-hyperlink-href-2}](links.html#attr-hyperlink-href)
attributes and
[`area`{#dom-tree-accessors:the-area-element-2}](image-maps.html#the-area-element)
elements with
[`href`{#dom-tree-accessors:attr-hyperlink-href-3}](links.html#attr-hyperlink-href)
attributes.

The [`forms`]{#dom-document-forms .dfn dfn-for="Document"
dfn-type="attribute"} attribute must return an
[`HTMLCollection`{#dom-tree-accessors:htmlcollection-9}](https://dom.spec.whatwg.org/#interface-htmlcollection){x-internal="htmlcollection"}
rooted at the [`Document`{#dom-tree-accessors:document-9}](#document)
node, whose filter matches only
[`form`{#dom-tree-accessors:the-form-element-2}](forms.html#the-form-element)
elements.

The [`scripts`]{#dom-document-scripts .dfn dfn-for="Document"
dfn-type="attribute"} attribute must return an
[`HTMLCollection`{#dom-tree-accessors:htmlcollection-10}](https://dom.spec.whatwg.org/#interface-htmlcollection){x-internal="htmlcollection"}
rooted at the [`Document`{#dom-tree-accessors:document-10}](#document)
node, whose filter matches only
[`script`{#dom-tree-accessors:the-script-element-2}](scripting.html#the-script-element)
elements.

------------------------------------------------------------------------

`collection`{.variable}` = ``document`{.variable}`.`[`getElementsByName`](#dom-document-getelementsbyname){#dom-document-getelementsbyname-dev}`(``name`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Document/getElementsByName](https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementsByName "The getElementsByName() method of the Document object returns a NodeList Collection of elements with a given name attribute in the document.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

Returns a
[`NodeList`{#dom-tree-accessors:nodelist}](https://dom.spec.whatwg.org/#interface-nodelist){x-internal="nodelist"}
of elements in the
[`Document`{#dom-tree-accessors:document-11}](#document) that have a
`name` attribute with the value `name`{.variable}.

The
[`getElementsByName(``elementName`{.variable}`)`]{#dom-document-getelementsbyname
.dfn dfn-for="Document" dfn-type="method"} method steps are to return a
[live](infrastructure.html#live){#dom-tree-accessors:live}
[`NodeList`{#dom-tree-accessors:nodelist-2}](https://dom.spec.whatwg.org/#interface-nodelist){x-internal="nodelist"}
containing all the [HTML
elements](infrastructure.html#html-elements){#dom-tree-accessors:html-elements}
in that document that have a `name` attribute whose value is [identical
to](https://infra.spec.whatwg.org/#string-is){#dom-tree-accessors:identical-to
x-internal="identical-to"} the `elementName`{.variable} argument, in
[tree
order](https://dom.spec.whatwg.org/#concept-tree-order){#dom-tree-accessors:tree-order-2
x-internal="tree-order"}. When the method is invoked on a
[`Document`{#dom-tree-accessors:document-12}](#document) object again
with the same argument, the user agent may return the same as the object
returned by the earlier call. In other cases, a new
[`NodeList`{#dom-tree-accessors:nodelist-3}](https://dom.spec.whatwg.org/#interface-nodelist){x-internal="nodelist"}
object must be returned.

------------------------------------------------------------------------

`document`{.variable}`.`[`currentScript`](#dom-document-currentscript){#dom-document-currentscript-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Document/currentScript](https://developer.mozilla.org/en-US/docs/Web/API/Document/currentScript "The Document.currentScript property returns the <script> element whose script is currently being processed and isn't a JavaScript module. (For modules use import.meta instead.)")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari8+]{.safari .yes}[Chrome29+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

Returns the
[`script`{#dom-tree-accessors:the-script-element-3}](scripting.html#the-script-element)
element, or the [SVG
`script`](https://svgwg.org/svg2-draft/interact.html#ScriptElement){#dom-tree-accessors:svg-script
x-internal="svg-script"} element, that is currently executing, as long
as the element represents a [classic
script](webappapis.html#classic-script){#dom-tree-accessors:classic-script}.
In the case of reentrant script execution, returns the one that most
recently started executing amongst those that have not yet finished
executing.

Returns null if the
[`Document`{#dom-tree-accessors:document-13}](#document) is not
currently executing a
[`script`{#dom-tree-accessors:the-script-element-4}](scripting.html#the-script-element)
or [SVG
`script`](https://svgwg.org/svg2-draft/interact.html#ScriptElement){#dom-tree-accessors:svg-script-2
x-internal="svg-script"} element (e.g., because the running script is an
event handler, or a timeout), or if the currently executing
[`script`{#dom-tree-accessors:the-script-element-5}](scripting.html#the-script-element)
or [SVG
`script`](https://svgwg.org/svg2-draft/interact.html#ScriptElement){#dom-tree-accessors:svg-script-3
x-internal="svg-script"} element represents a [module
script](webappapis.html#module-script){#dom-tree-accessors:module-script}.

The [`currentScript`]{#dom-document-currentscript .dfn
dfn-for="Document" dfn-type="attribute"} attribute, on getting, must
return the value to which it was most recently set. When the
[`Document`{#dom-tree-accessors:document-14}](#document) is created, the
[`currentScript`{#dom-tree-accessors:dom-document-currentscript}](#dom-document-currentscript)
must be initialized to null.

This API has fallen out of favor in the implementer and standards
community, as it globally exposes
[`script`{#dom-tree-accessors:the-script-element-6}](scripting.html#the-script-element)
or [SVG
`script`](https://svgwg.org/svg2-draft/interact.html#ScriptElement){#dom-tree-accessors:svg-script-4
x-internal="svg-script"} elements. As such, it is not available in newer
contexts, such as when running [module
scripts](webappapis.html#module-script){#dom-tree-accessors:module-script-2}
or when running scripts in a [shadow
tree](https://dom.spec.whatwg.org/#concept-shadow-tree){#dom-tree-accessors:shadow-tree
x-internal="shadow-tree"}. We are looking into creating a new solution
for identifying the running script in such contexts, which does not make
it globally available: see [issue
#1013](https://github.com/whatwg/html/issues/1013).

------------------------------------------------------------------------

The [`Document`{#dom-tree-accessors:document-15}](#document) interface
[supports named
properties](https://webidl.spec.whatwg.org/#dfn-support-named-properties){#dom-tree-accessors:support-named-properties
x-internal="support-named-properties"}. The [supported property
names](https://webidl.spec.whatwg.org/#dfn-supported-property-names){#dom-tree-accessors:supported-property-names
x-internal="supported-property-names"} of a
[`Document`{#dom-tree-accessors:document-16}](#document) object
`document`{.variable} at any moment consist of the following, in [tree
order](https://dom.spec.whatwg.org/#concept-tree-order){#dom-tree-accessors:tree-order-3
x-internal="tree-order"} according to the element that contributed them,
ignoring later duplicates, and with values from
[`id`{#dom-tree-accessors:the-id-attribute}](#the-id-attribute)
attributes coming before values from `name` attributes when the same
element contributes both:

- the value of the `name` content attribute for all
  [exposed](#exposed){#dom-tree-accessors:exposed}
  [`embed`{#dom-tree-accessors:the-embed-element-3}](iframe-embed-object.html#the-embed-element),
  [`form`{#dom-tree-accessors:the-form-element-3}](forms.html#the-form-element),
  [`iframe`{#dom-tree-accessors:the-iframe-element}](iframe-embed-object.html#the-iframe-element),
  [`img`{#dom-tree-accessors:the-img-element-3}](embedded-content.html#the-img-element),
  and [exposed](#exposed){#dom-tree-accessors:exposed-2}
  [`object`{#dom-tree-accessors:the-object-element}](iframe-embed-object.html#the-object-element)
  elements that have a non-empty `name` content attribute and are [in a
  document
  tree](https://dom.spec.whatwg.org/#in-a-document-tree){#dom-tree-accessors:in-a-document-tree
  x-internal="in-a-document-tree"} with `document`{.variable} as their
  [root](https://dom.spec.whatwg.org/#concept-tree-root){#dom-tree-accessors:root
  x-internal="root"};

- the value of the
  [`id`{#dom-tree-accessors:the-id-attribute-2}](#the-id-attribute)
  content attribute for all
  [exposed](#exposed){#dom-tree-accessors:exposed-3}
  [`object`{#dom-tree-accessors:the-object-element-2}](iframe-embed-object.html#the-object-element)
  elements that have a non-empty
  [`id`{#dom-tree-accessors:the-id-attribute-3}](#the-id-attribute)
  content attribute and are [in a document
  tree](https://dom.spec.whatwg.org/#in-a-document-tree){#dom-tree-accessors:in-a-document-tree-2
  x-internal="in-a-document-tree"} with `document`{.variable} as their
  [root](https://dom.spec.whatwg.org/#concept-tree-root){#dom-tree-accessors:root-2
  x-internal="root"}; and

- the value of the
  [`id`{#dom-tree-accessors:the-id-attribute-4}](#the-id-attribute)
  content attribute for all
  [`img`{#dom-tree-accessors:the-img-element-4}](embedded-content.html#the-img-element)
  elements that have both a non-empty
  [`id`{#dom-tree-accessors:the-id-attribute-5}](#the-id-attribute)
  content attribute and a non-empty `name` content attribute, and are
  [in a document
  tree](https://dom.spec.whatwg.org/#in-a-document-tree){#dom-tree-accessors:in-a-document-tree-3
  x-internal="in-a-document-tree"} with `document`{.variable} as their
  [root](https://dom.spec.whatwg.org/#concept-tree-root){#dom-tree-accessors:root-3
  x-internal="root"}.

To [determine the value of a named
property](https://webidl.spec.whatwg.org/#dfn-determine-the-value-of-a-named-property){#dom-tree-accessors:determine-the-value-of-a-named-property
x-internal="determine-the-value-of-a-named-property"} `name`{.variable}
for a [`Document`{#dom-tree-accessors:document-17}](#document), the user
agent must return the value obtained using the following steps:

1.  Let `elements`{.variable} be the list of [named
    elements](#dom-document-nameditem-filter){#dom-tree-accessors:dom-document-nameditem-filter}
    with the name `name`{.variable} that are [in a document
    tree](https://dom.spec.whatwg.org/#in-a-document-tree){#dom-tree-accessors:in-a-document-tree-4
    x-internal="in-a-document-tree"} with the
    [`Document`{#dom-tree-accessors:document-18}](#document) as their
    [root](https://dom.spec.whatwg.org/#concept-tree-root){#dom-tree-accessors:root-4
    x-internal="root"}.

    There will be at least one such element, since the algorithm would
    otherwise not have been [invoked by Web
    IDL](https://webidl.spec.whatwg.org/#LegacyPlatformObjectGetOwnProperty){#dom-tree-accessors:legacyplatformobjectgetownproperty
    x-internal="legacyplatformobjectgetownproperty"}.

2.  If `elements`{.variable} has only one element, and that element is
    an
    [`iframe`{#dom-tree-accessors:the-iframe-element-2}](iframe-embed-object.html#the-iframe-element)
    element, and that
    [`iframe`{#dom-tree-accessors:the-iframe-element-3}](iframe-embed-object.html#the-iframe-element)
    element\'s [content
    navigable](document-sequences.html#content-navigable){#dom-tree-accessors:content-navigable}
    is not null, then return the [active
    `WindowProxy`](document-sequences.html#nav-wp){#dom-tree-accessors:nav-wp}
    of the element\'s [content
    navigable](document-sequences.html#content-navigable){#dom-tree-accessors:content-navigable-2}.

3.  Otherwise, if `elements`{.variable} has only one element, return
    that element.

4.  Otherwise, return an
    [`HTMLCollection`{#dom-tree-accessors:htmlcollection-11}](https://dom.spec.whatwg.org/#interface-htmlcollection){x-internal="htmlcollection"}
    rooted at the
    [`Document`{#dom-tree-accessors:document-19}](#document) node, whose
    filter matches only [named
    elements](#dom-document-nameditem-filter){#dom-tree-accessors:dom-document-nameditem-filter-2}
    with the name `name`{.variable}.

[Named elements]{#dom-document-nameditem-filter .dfn} with the name
`name`{.variable}, for the purposes of the above algorithm, are those
that are either:

- [Exposed](#exposed){#dom-tree-accessors:exposed-4}
  [`embed`{#dom-tree-accessors:the-embed-element-4}](iframe-embed-object.html#the-embed-element),
  [`form`{#dom-tree-accessors:the-form-element-4}](forms.html#the-form-element),
  [`iframe`{#dom-tree-accessors:the-iframe-element-4}](iframe-embed-object.html#the-iframe-element),
  [`img`{#dom-tree-accessors:the-img-element-5}](embedded-content.html#the-img-element),
  or [exposed](#exposed){#dom-tree-accessors:exposed-5}
  [`object`{#dom-tree-accessors:the-object-element-3}](iframe-embed-object.html#the-object-element)
  elements that have a `name` content attribute whose value is
  `name`{.variable}, or
- [Exposed](#exposed){#dom-tree-accessors:exposed-6}
  [`object`{#dom-tree-accessors:the-object-element-4}](iframe-embed-object.html#the-object-element)
  elements that have an
  [`id`{#dom-tree-accessors:the-id-attribute-6}](#the-id-attribute)
  content attribute whose value is `name`{.variable}, or
- [`img`{#dom-tree-accessors:the-img-element-6}](embedded-content.html#the-img-element)
  elements that have an
  [`id`{#dom-tree-accessors:the-id-attribute-7}](#the-id-attribute)
  content attribute whose value is `name`{.variable}, and that have a
  non-empty `name` content attribute present also.

An
[`embed`{#dom-tree-accessors:the-embed-element-5}](iframe-embed-object.html#the-embed-element)
or
[`object`{#dom-tree-accessors:the-object-element-5}](iframe-embed-object.html#the-object-element)
element is said to be [exposed]{#exposed .dfn} if it has no
[exposed](#exposed){#dom-tree-accessors:exposed-7}
[`object`{#dom-tree-accessors:the-object-element-6}](iframe-embed-object.html#the-object-element)
ancestor, and, for
[`object`{#dom-tree-accessors:the-object-element-7}](iframe-embed-object.html#the-object-element)
elements, is additionally either not showing its [fallback
content](#fallback-content){#dom-tree-accessors:fallback-content} or has
no
[`object`{#dom-tree-accessors:the-object-element-8}](iframe-embed-object.html#the-object-element)
or
[`embed`{#dom-tree-accessors:the-embed-element-6}](iframe-embed-object.html#the-embed-element)
descendants.

------------------------------------------------------------------------

The [`dir`{#dom-tree-accessors:dom-document-dir}](#dom-document-dir)
attribute on the
[`Document`{#dom-tree-accessors:document-20}](#document) interface is
defined along with the [`dir`{#dom-tree-accessors:attr-dir}](#attr-dir)
content attribute.

### [3.2]{.secno} Elements[](#elements){.self-link}

#### [3.2.1]{.secno} Semantics[](#semantics-2){.self-link} {#semantics-2}

Elements, attributes, and attribute values in HTML are defined (by this
specification) to have certain meanings (semantics). For example, the
[`ol`{#semantics-2:the-ol-element}](grouping-content.html#the-ol-element)
element represents an ordered list, and the
[`lang`{#semantics-2:attr-lang}](#attr-lang) attribute represents the
language of the content.

These definitions allow HTML processors, such as web browsers or search
engines, to present and use documents and applications in a wide variety
of contexts that the author might not have considered.

::: example
As a simple example, consider a web page written by an author who only
considered desktop computer web browsers:

``` html
<!DOCTYPE HTML>
<html lang="en">
 <head>
  <title>My Page</title>
 </head>
 <body>
  <h1>Welcome to my page</h1>
  <p>I like cars and lorries and have a big Jeep!</p>
  <h2>Where I live</h2>
  <p>I live in a small hut on a mountain!</p>
 </body>
</html>
```

Because HTML conveys *meaning*, rather than presentation, the same page
can also be used by a small browser on a mobile phone, without any
change to the page. Instead of headings being in large letters as on the
desktop, for example, the browser on the mobile phone might use the same
size text for the whole page, but with the headings in bold.

But it goes further than just differences in screen size: the same page
could equally be used by a blind user using a browser based around
speech synthesis, which instead of displaying the page on a screen,
reads the page to the user, e.g. using headphones. Instead of large text
for the headings, the speech browser might use a different volume or a
slower voice.

That\'s not all, either. Since the browsers know which parts of the page
are the headings, they can create a document outline that the user can
use to quickly navigate around the document, using keys for \"jump to
next heading\" or \"jump to previous heading\". Such features are
especially common with speech browsers, where users would otherwise find
quickly navigating a page quite difficult.

Even beyond browsers, software can make use of this information. Search
engines can use the headings to more effectively index a page, or to
provide quick links to subsections of the page from their results. Tools
can use the headings to create a table of contents (that is in fact how
this very specification\'s table of contents is generated).

This example has focused on headings, but the same principle applies to
all of the semantics in HTML.
:::

Authors must not use elements, attributes, or attribute values for
purposes other than their appropriate intended semantic purpose, as
doing so prevents software from correctly processing the page.

::: example
For example, the following snippet, intended to represent the heading of
a corporate site, is non-conforming because the second line is not
intended to be a heading of a subsection, but merely a subheading or
subtitle (a subordinate heading for the same section).

``` bad
<body>
 <h1>ACME Corporation</h1>
 <h2>The leaders in arbitrary fast delivery since 1920</h2>
 ...
```

The
[`hgroup`{#semantics-2:the-hgroup-element}](sections.html#the-hgroup-element)
element can be used for these kinds of situations:

``` html
<body>
 <hgroup>
  <h1>ACME Corporation</h1>
  <p>The leaders in arbitrary fast delivery since 1920</p>
 </hgroup>
 ...
```
:::

::: example
The document in this next example is similarly non-conforming, despite
being syntactically correct, because the data placed in the cells is
clearly not tabular data, and the
[`cite`{#semantics-2:the-cite-element}](text-level-semantics.html#the-cite-element)
element mis-used:

``` {.bad lang="en-GB"}
<!DOCTYPE HTML>
<html lang="en-GB">
 <head> <title> Demonstration </title> </head>
 <body>
  <table>
   <tr> <td> My favourite animal is the cat. </td> </tr>
   <tr>
    <td>
     —<a href="https://example.org/~ernest/"><cite>Ernest</cite></a>,
     in an essay from 1992
    </td>
   </tr>
  </table>
 </body>
</html>
```

This would make software that relies on these semantics fail: for
example, a speech browser that allowed a blind user to navigate tables
in the document would report the quote above as a table, confusing the
user; similarly, a tool that extracted titles of works from pages would
extract \"Ernest\" as the title of a work, even though it\'s actually a
person\'s name, not a title.

A corrected version of this document might be:

``` {lang="en-GB"}
<!DOCTYPE HTML>
<html lang="en-GB">
 <head> <title> Demonstration </title> </head>
 <body>
  <blockquote>
   <p> My favourite animal is the cat. </p>
  </blockquote>
  <p>
   —<a href="https://example.org/~ernest/">Ernest</a>,
   in an essay from 1992
  </p>
 </body>
</html>
```
:::

Authors must not use elements, attributes, or attribute values that are
not permitted by this specification or [other applicable
specifications](infrastructure.html#other-applicable-specifications){#semantics-2:other-applicable-specifications},
as doing so makes it significantly harder for the language to be
extended in the future.

::: example
In the next example, there is a non-conforming attribute value
(\"carpet\") and a non-conforming attribute (\"texture\"), which is not
permitted by this specification:

``` bad
<label>Carpet: <input type="carpet" name="c" texture="deep pile"></label>
```

Here would be an alternative and correct way to mark this up:

``` html
<label>Carpet: <input type="text" class="carpet" name="c" data-texture="deep pile"></label>
```
:::

DOM nodes whose [node
document](https://dom.spec.whatwg.org/#concept-node-document){#semantics-2:node-document
x-internal="node-document"}\'s [browsing
context](document-sequences.html#concept-document-bc){#semantics-2:concept-document-bc}
is null are exempt from all document conformance requirements other than
the [HTML syntax](syntax.html#writing) requirements and [XML
syntax](xhtml.html#writing-xhtml-documents) requirements.

::: example
In particular, the
[`template`{#semantics-2:the-template-element}](scripting.html#the-template-element)
element\'s [template
contents](scripting.html#template-contents){#semantics-2:template-contents}\'s
[node
document](https://dom.spec.whatwg.org/#concept-node-document){#semantics-2:node-document-2
x-internal="node-document"}\'s [browsing
context](document-sequences.html#concept-document-bc){#semantics-2:concept-document-bc-2}
is null. For example, the [content
model](#concept-element-content-model){#semantics-2:concept-element-content-model}
requirements and attribute value microsyntax requirements do not apply
to a
[`template`{#semantics-2:the-template-element-2}](scripting.html#the-template-element)
element\'s [template
contents](scripting.html#template-contents){#semantics-2:template-contents-2}.
In this example an
[`img`{#semantics-2:the-img-element}](embedded-content.html#the-img-element)
element has attribute values that are placeholders that would be invalid
outside a
[`template`{#semantics-2:the-template-element-3}](scripting.html#the-template-element)
element.

``` html
<template>
 <article>
  <img src="{{src}}" alt="{{alt}}">
  <h1></h1>
 </article>
</template>
```

However, if the above markup were to omit the `</h1>` end tag, that
would be a violation of the [HTML syntax](syntax.html#writing), and
would thus be flagged as an error by conformance checkers.
:::

Through scripting and using other mechanisms, the values of attributes,
text, and indeed the entire structure of the document may change
dynamically while a user agent is processing it. The semantics of a
document at an instant in time are those represented by the state of the
document at that instant in time, and the semantics of a document can
therefore change over time. User agents must update their presentation
of the document as this occurs.

HTML has a
[`progress`{#semantics-2:the-progress-element}](form-elements.html#the-progress-element)
element that describes a progress bar. If its \"value\" attribute is
dynamically updated by a script, the UA would update the rendering to
show the progress changing.

#### [3.2.2]{.secno} Elements in the DOM[](#elements-in-the-dom){.self-link}

The nodes representing [HTML
elements](infrastructure.html#html-elements){#elements-in-the-dom:html-elements}
in the DOM must implement, and expose to scripts, the interfaces listed
for them in the relevant sections of this specification. This includes
[HTML
elements](infrastructure.html#html-elements){#elements-in-the-dom:html-elements-2}
in [XML
documents](https://dom.spec.whatwg.org/#xml-document){#elements-in-the-dom:xml-documents
x-internal="xml-documents"}, even when those documents are in another
context (e.g. inside an XSLT transform).

Elements in the DOM [represent]{#represents .dfn} things; that is, they
have intrinsic *meaning*, also known as semantics.

For example, an
[`ol`{#elements-in-the-dom:the-ol-element}](grouping-content.html#the-ol-element)
element represents an ordered list.

Elements can be [referenced]{#referenced .dfn} (referred to) in some
way, either explicitly or implicitly. One way that an element in the DOM
can be explicitly referenced is by giving an
[`id`{#elements-in-the-dom:the-id-attribute}](#the-id-attribute)
attribute to the element, and then creating a
[hyperlink](links.html#hyperlink){#elements-in-the-dom:hyperlink} with
that [`id`{#elements-in-the-dom:the-id-attribute-2}](#the-id-attribute)
attribute\'s value as the
[fragment](browsing-the-web.html#navigate-fragid){#elements-in-the-dom:navigate-fragid}
for the
[hyperlink](links.html#hyperlink){#elements-in-the-dom:hyperlink-2}\'s
[`href`{#elements-in-the-dom:attr-hyperlink-href}](links.html#attr-hyperlink-href)
attribute value. Hyperlinks are not necessary for a reference, however;
any manner of referring to the element in question will suffice.

::: example
Consider the following
[`figure`{#elements-in-the-dom:the-figure-element}](grouping-content.html#the-figure-element)
element, which is given an
[`id`{#elements-in-the-dom:the-id-attribute-3}](#the-id-attribute)
attribute:

``` html
<figure id="module-script-graph">
  <img src="module-script-graph.svg"
       alt="Module A depends on module B, which depends
            on modules C and D.">
  <figcaption>Figure 27: a simple module graph</figcaption>
</figure>
```

A
[hyperlink](links.html#hyperlink){#elements-in-the-dom:hyperlink-3}-based
[reference](#referenced){#elements-in-the-dom:referenced} could be
created using the
[`a`{#elements-in-the-dom:the-a-element}](text-level-semantics.html#the-a-element)
element, like so:

``` html
As we can see in <a href="#module-script-graph">figure 27</a>, ...
```

However, there are many other ways of
[referencing](#referenced){#elements-in-the-dom:referenced-2} the
[`figure`{#elements-in-the-dom:the-figure-element-2}](grouping-content.html#the-figure-element)
element, such as:

- \"As depicted in the figure of modules A, B, C, and D\...\"

- \"In Figure 27\...\" (without a hyperlink)

- \"From the contents of the \'simple module graph\' figure\...\"

- \"In the figure below\...\" (but [this is
  discouraged](grouping-content.html#figure-note-about-references))
:::

The basic interface, from which all the [HTML
elements](infrastructure.html#html-elements){#elements-in-the-dom:html-elements-3}\'
interfaces inherit, and which must be used by elements that have no
additional requirements, is the
[`HTMLElement`{#elements-in-the-dom:htmlelement}](#htmlelement)
interface.

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement "The HTMLElement interface represents any HTML element. Some elements directly implement this interface, while others implement it via an interface that inherits it.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera8+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android1+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::

:::: feature
[HTMLUnknownElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLUnknownElement "The HTMLUnknownElement interface represents an invalid HTML element and derives from the HTMLElement interface, but without implementing any additional properties or methods.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome15+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::::

``` idl
[Exposed=Window]
interface HTMLElement : Element {
  [HTMLConstructor] constructor();

  // metadata attributes
  [CEReactions] attribute DOMString title;
  [CEReactions] attribute DOMString lang;
  [CEReactions] attribute boolean translate;
  [CEReactions] attribute DOMString dir;

  // user interaction
  [CEReactions] attribute (boolean or unrestricted double or DOMString)? hidden;
  [CEReactions] attribute boolean inert;
  undefined click();
  [CEReactions] attribute DOMString accessKey;
  readonly attribute DOMString accessKeyLabel;
  [CEReactions] attribute boolean draggable;
  [CEReactions] attribute boolean spellcheck;
  [CEReactions] attribute DOMString writingSuggestions;
  [CEReactions] attribute DOMString autocapitalize;
  [CEReactions] attribute boolean autocorrect;

  [CEReactions] attribute [LegacyNullToEmptyString] DOMString innerText;
  [CEReactions] attribute [LegacyNullToEmptyString] DOMString outerText;

  ElementInternals attachInternals();

  // The popover API
  undefined showPopover(optional ShowPopoverOptions options = {});
  undefined hidePopover();
  boolean togglePopover(optional (TogglePopoverOptions or boolean) options = {});
  [CEReactions] attribute DOMString? popover;
};

dictionary ShowPopoverOptions {
  HTMLElement source;
};

dictionary TogglePopoverOptions : ShowPopoverOptions {
  boolean force;
};

HTMLElement includes GlobalEventHandlers;
HTMLElement includes ElementContentEditable;
HTMLElement includes HTMLOrSVGElement;

[Exposed=Window]
interface HTMLUnknownElement : HTMLElement {
  // Note: intentionally no [HTMLConstructor]
};
```

The [`HTMLElement`{#elements-in-the-dom:htmlelement-7}](#htmlelement)
interface holds methods and attributes related to a number of disparate
features, and the members of this interface are therefore described in
various different sections of this specification.

------------------------------------------------------------------------

The [element
interface](https://dom.spec.whatwg.org/#concept-element-interface){#elements-in-the-dom:element-interface
x-internal="element-interface"} for an element with name
`name`{.variable} in the [HTML
namespace](https://infra.spec.whatwg.org/#html-namespace){#elements-in-the-dom:html-namespace-2
x-internal="html-namespace-2"} is determined as follows:

1.  If `name`{.variable} is
    [`applet`{#elements-in-the-dom:applet}](obsolete.html#applet),
    [`bgsound`{#elements-in-the-dom:bgsound}](obsolete.html#bgsound),
    [`blink`{#elements-in-the-dom:blink}](obsolete.html#blink),
    [`isindex`{#elements-in-the-dom:isindex}](obsolete.html#isindex),
    [`keygen`{#elements-in-the-dom:keygen}](obsolete.html#keygen),
    [`multicol`{#elements-in-the-dom:multicol}](obsolete.html#multicol),
    [`nextid`{#elements-in-the-dom:nextid}](obsolete.html#nextid), or
    [`spacer`{#elements-in-the-dom:spacer}](obsolete.html#spacer), then
    return
    [`HTMLUnknownElement`{#elements-in-the-dom:htmlunknownelement}](#htmlunknownelement).

2.  If `name`{.variable} is
    [`acronym`{#elements-in-the-dom:acronym}](obsolete.html#acronym),
    [`basefont`{#elements-in-the-dom:basefont}](obsolete.html#basefont),
    [`big`{#elements-in-the-dom:big}](obsolete.html#big),
    [`center`{#elements-in-the-dom:center}](obsolete.html#center),
    [`nobr`{#elements-in-the-dom:nobr}](obsolete.html#nobr),
    [`noembed`{#elements-in-the-dom:noembed}](obsolete.html#noembed),
    [`noframes`{#elements-in-the-dom:noframes}](obsolete.html#noframes),
    [`plaintext`{#elements-in-the-dom:plaintext}](obsolete.html#plaintext),
    [`rb`{#elements-in-the-dom:rb}](obsolete.html#rb),
    [`rtc`{#elements-in-the-dom:rtc}](obsolete.html#rtc),
    [`strike`{#elements-in-the-dom:strike}](obsolete.html#strike), or
    [`tt`{#elements-in-the-dom:tt}](obsolete.html#tt), then return
    [`HTMLElement`{#elements-in-the-dom:htmlelement-8}](#htmlelement).

3.  If `name`{.variable} is
    [`listing`{#elements-in-the-dom:listing}](obsolete.html#listing) or
    [`xmp`{#elements-in-the-dom:xmp}](obsolete.html#xmp), then return
    [`HTMLPreElement`{#elements-in-the-dom:htmlpreelement}](grouping-content.html#htmlpreelement).

4.  Otherwise, if this specification defines an interface appropriate
    for the [element
    type](infrastructure.html#element-type){#elements-in-the-dom:element-type}
    corresponding to the local name `name`{.variable}, then return that
    interface.

5.  If [other applicable
    specifications](infrastructure.html#other-applicable-specifications){#elements-in-the-dom:other-applicable-specifications}
    define an appropriate interface for `name`{.variable}, then return
    the interface they define.

6.  If `name`{.variable} is a [valid custom element
    name](custom-elements.html#valid-custom-element-name){#elements-in-the-dom:valid-custom-element-name},
    then return
    [`HTMLElement`{#elements-in-the-dom:htmlelement-9}](#htmlelement).

7.  Return
    [`HTMLUnknownElement`{#elements-in-the-dom:htmlunknownelement-2}](#htmlunknownelement).

The use of
[`HTMLElement`{#elements-in-the-dom:htmlelement-10}](#htmlelement)
instead of
[`HTMLUnknownElement`{#elements-in-the-dom:htmlunknownelement-3}](#htmlunknownelement)
in the case of [valid custom element
names](custom-elements.html#valid-custom-element-name){#elements-in-the-dom:valid-custom-element-name-2}
is done to ensure that any potential future
[upgrades](custom-elements.html#upgrades){#elements-in-the-dom:upgrades}
only cause a linear transition of the element\'s prototype chain, from
[`HTMLElement`{#elements-in-the-dom:htmlelement-11}](#htmlelement) to a
subclass, instead of a lateral one, from
[`HTMLUnknownElement`{#elements-in-the-dom:htmlunknownelement-4}](#htmlunknownelement)
to an unrelated subclass.

Features shared between HTML and SVG elements use the
[`HTMLOrSVGElement`{#elements-in-the-dom:htmlorsvgelement-2}](#htmlorsvgelement)
interface mixin: [\[SVG\]](references.html#refsSVG)

``` idl
interface mixin HTMLOrSVGElement {
  [SameObject] readonly attribute DOMStringMap dataset;
  attribute DOMString nonce; // intentionally no [CEReactions]

  [CEReactions] attribute boolean autofocus;
  [CEReactions] attribute long tabIndex;
  undefined focus(optional FocusOptions options = {});
  undefined blur();
};
```

::: example
An example of an element that is neither an HTML nor SVG element is one
created as follows:

``` html
const el = document.createElementNS("some namespace", "example");
console.assert(el.constructor === Element);
```
:::

#### [3.2.3]{.secno} HTML element constructors[](#html-element-constructors){.self-link}

To support the [custom elements](custom-elements.html#custom-elements)
feature, all HTML elements have special constructor behavior. This is
indicated via the [`[HTMLConstructor]`]{#htmlconstructor .dfn
dfn-type="extended-attribute" lt="HTMLConstructor"} IDL [extended
attribute](https://webidl.spec.whatwg.org/#dfn-extended-attribute){#html-element-constructors:extended-attribute
x-internal="extended-attribute"}. It indicates that the interface object
for the given interface will have a specific behavior when called, as
defined in detail below.

The
[`[HTMLConstructor]`{#html-element-constructors:htmlconstructor}](#htmlconstructor)
extended attribute must take no arguments, and must only appear on
[constructor
operations](https://webidl.spec.whatwg.org/#idl-constructors){#html-element-constructors:constructor-operation
x-internal="constructor-operation"}. It must appear only once on a
constructor operation, and the interface must contain only the single,
annotated constructor operation, and no others. The annotated
constructor operation must be declared to take no arguments.

Interfaces declared with constructor operations that are annotated with
the
[`[HTMLConstructor]`{#html-element-constructors:htmlconstructor-2}](#htmlconstructor)
extended attribute have the following [overridden constructor
steps](https://webidl.spec.whatwg.org/#overridden-constructor-steps){#html-element-constructors:overridden-constructor-steps
x-internal="overridden-constructor-steps"}:

1.  If
    [NewTarget](https://tc39.es/ecma262/#sec-built-in-function-objects){#html-element-constructors:newtarget
    x-internal="newtarget"} is equal to the [active function
    object](https://tc39.es/ecma262/#active-function-object){#html-element-constructors:active-function-object
    x-internal="active-function-object"}, then throw a
    [`TypeError`{#html-element-constructors:typeerror}](https://tc39.es/ecma262/#sec-native-error-types-used-in-this-standard-typeerror){x-internal="typeerror"}.

    ::: example
    This can occur when a custom element is defined using an [element
    interface](https://dom.spec.whatwg.org/#concept-element-interface){#html-element-constructors:element-interface
    x-internal="element-interface"} as its constructor:

    ``` js
    customElements.define("bad-1", HTMLButtonElement);
    new HTMLButtonElement();          // (1)
    document.createElement("bad-1");  // (2)
    ```

    In this case, during the execution of
    [`HTMLButtonElement`{#html-element-constructors:htmlbuttonelement}](form-elements.html#htmlbuttonelement)
    (either explicitly, as in (1), or implicitly, as in (2)), both the
    [active function
    object](https://tc39.es/ecma262/#active-function-object){#html-element-constructors:active-function-object-2
    x-internal="active-function-object"} and
    [NewTarget](https://tc39.es/ecma262/#sec-built-in-function-objects){#html-element-constructors:newtarget-2
    x-internal="newtarget"} are
    [`HTMLButtonElement`{#html-element-constructors:htmlbuttonelement-2}](form-elements.html#htmlbuttonelement).
    If this check was not present, it would be possible to create an
    instance of
    [`HTMLButtonElement`{#html-element-constructors:htmlbuttonelement-3}](form-elements.html#htmlbuttonelement)
    whose local name was `bad-1`.
    :::

2.  Let `registry`{.variable} be null.

3.  If the [surrounding
    agent](https://tc39.es/ecma262/#surrounding-agent){#html-element-constructors:surrounding-agent
    x-internal="surrounding-agent"}\'s [active custom element
    constructor
    map](custom-elements.html#active-custom-element-constructor-map){#html-element-constructors:active-custom-element-constructor-map}\[[NewTarget](https://tc39.es/ecma262/#sec-built-in-function-objects){#html-element-constructors:newtarget-3
    x-internal="newtarget"}\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#html-element-constructors:map-exists
    x-internal="map-exists"}:

    1.  Set `registry`{.variable} to the [surrounding
        agent](https://tc39.es/ecma262/#surrounding-agent){#html-element-constructors:surrounding-agent-2
        x-internal="surrounding-agent"}\'s [active custom element
        constructor
        map](custom-elements.html#active-custom-element-constructor-map){#html-element-constructors:active-custom-element-constructor-map-2}\[[NewTarget](https://tc39.es/ecma262/#sec-built-in-function-objects){#html-element-constructors:newtarget-4
        x-internal="newtarget"}\].

    2.  [Remove](https://infra.spec.whatwg.org/#map-remove){#html-element-constructors:map-remove
        x-internal="map-remove"} the [surrounding
        agent](https://tc39.es/ecma262/#surrounding-agent){#html-element-constructors:surrounding-agent-3
        x-internal="surrounding-agent"}\'s [active custom element
        constructor
        map](custom-elements.html#active-custom-element-constructor-map){#html-element-constructors:active-custom-element-constructor-map-3}\[[NewTarget](https://tc39.es/ecma262/#sec-built-in-function-objects){#html-element-constructors:newtarget-5
        x-internal="newtarget"}\].

4.  Otherwise, set `registry`{.variable} to [current global
    object](webappapis.html#current-global-object){#html-element-constructors:current-global-object}\'s
    [associated
    `Document`](nav-history-apis.html#concept-document-window){#html-element-constructors:concept-document-window}\'s
    [custom element
    registry](https://dom.spec.whatwg.org/#document-custom-element-registry){#html-element-constructors:document-custom-element-registry
    x-internal="document-custom-element-registry"}.

5.  Let `definition`{.variable} be the item in `registry`{.variable}\'s
    [custom element definition
    set](custom-elements.html#custom-element-definition-set){#html-element-constructors:custom-element-definition-set}
    with
    [constructor](custom-elements.html#concept-custom-element-definition-constructor){#html-element-constructors:concept-custom-element-definition-constructor}
    equal to
    [NewTarget](https://tc39.es/ecma262/#sec-built-in-function-objects){#html-element-constructors:newtarget-6
    x-internal="newtarget"}. If there is no such item, then throw a
    [`TypeError`{#html-element-constructors:typeerror-2}](https://tc39.es/ecma262/#sec-native-error-types-used-in-this-standard-typeerror){x-internal="typeerror"}.

    Since there can be no item in `registry`{.variable}\'s [custom
    element definition
    set](custom-elements.html#custom-element-definition-set){#html-element-constructors:custom-element-definition-set-2}
    with a
    [constructor](custom-elements.html#concept-custom-element-definition-constructor){#html-element-constructors:concept-custom-element-definition-constructor-2}
    of undefined, this step also prevents HTML element constructors from
    being called as functions (since in that case
    [NewTarget](https://tc39.es/ecma262/#sec-built-in-function-objects){#html-element-constructors:newtarget-7
    x-internal="newtarget"} will be undefined).

6.  Let `isValue`{.variable} be null.

7.  If `definition`{.variable}\'s [local
    name](custom-elements.html#concept-custom-element-definition-local-name){#html-element-constructors:concept-custom-element-definition-local-name}
    is equal to `definition`{.variable}\'s
    [name](custom-elements.html#concept-custom-element-definition-name){#html-element-constructors:concept-custom-element-definition-name}
    (i.e., `definition`{.variable} is for an [autonomous custom
    element](custom-elements.html#autonomous-custom-element){#html-element-constructors:autonomous-custom-element}):

    1.  If the [active function
        object](https://tc39.es/ecma262/#active-function-object){#html-element-constructors:active-function-object-3
        x-internal="active-function-object"} is not
        [`HTMLElement`{#html-element-constructors:htmlelement}](#htmlelement),
        then throw a
        [`TypeError`{#html-element-constructors:typeerror-3}](https://tc39.es/ecma262/#sec-native-error-types-used-in-this-standard-typeerror){x-internal="typeerror"}.

        ::: example
        This can occur when a custom element is defined to not extend
        any local names, but inherits from a
        non-[`HTMLElement`{#html-element-constructors:htmlelement-2}](#htmlelement)
        class:

        ``` js
        customElements.define("bad-2", class Bad2 extends HTMLParagraphElement {});
        ```

        In this case, during the (implicit) `super()` call that occurs
        when constructing an instance of `Bad2`, the [active function
        object](https://tc39.es/ecma262/#active-function-object){#html-element-constructors:active-function-object-4
        x-internal="active-function-object"} is
        [`HTMLParagraphElement`{#html-element-constructors:htmlparagraphelement}](grouping-content.html#htmlparagraphelement),
        not
        [`HTMLElement`{#html-element-constructors:htmlelement-3}](#htmlelement).
        :::

8.  Otherwise (i.e., if `definition`{.variable} is for a [customized
    built-in
    element](custom-elements.html#customized-built-in-element){#html-element-constructors:customized-built-in-element}):

    1.  Let `valid local names`{.variable} be the list of local names
        for elements defined in this specification or in [other
        applicable
        specifications](infrastructure.html#other-applicable-specifications){#html-element-constructors:other-applicable-specifications}
        that use the [active function
        object](https://tc39.es/ecma262/#active-function-object){#html-element-constructors:active-function-object-5
        x-internal="active-function-object"} as their [element
        interface](https://dom.spec.whatwg.org/#concept-element-interface){#html-element-constructors:element-interface-2
        x-internal="element-interface"}.

    2.  If `valid local names`{.variable} does not contain
        `definition`{.variable}\'s [local
        name](custom-elements.html#concept-custom-element-definition-local-name){#html-element-constructors:concept-custom-element-definition-local-name-2},
        then throw a
        [`TypeError`{#html-element-constructors:typeerror-4}](https://tc39.es/ecma262/#sec-native-error-types-used-in-this-standard-typeerror){x-internal="typeerror"}.

        ::: example
        This can occur when a custom element is defined to extend a
        given local name but inherits from the wrong class:

        ``` js
        customElements.define("bad-3", class Bad3 extends HTMLQuoteElement {}, { extends: "p" });
        ```

        In this case, during the (implicit) `super()` call that occurs
        when constructing an instance of `Bad3`,
        `valid local names`{.variable} is the list containing
        [`q`{#html-element-constructors:the-q-element}](text-level-semantics.html#the-q-element)
        and
        [`blockquote`{#html-element-constructors:the-blockquote-element}](grouping-content.html#the-blockquote-element),
        but `definition`{.variable}\'s [local
        name](custom-elements.html#concept-custom-element-definition-local-name){#html-element-constructors:concept-custom-element-definition-local-name-3}
        is
        [`p`{#html-element-constructors:the-p-element}](grouping-content.html#the-p-element),
        which is not in that list.
        :::

    3.  Set `isValue`{.variable} to `definition`{.variable}\'s
        [name](custom-elements.html#concept-custom-element-definition-name){#html-element-constructors:concept-custom-element-definition-name-2}.

9.  If `definition`{.variable}\'s [construction
    stack](custom-elements.html#concept-custom-element-definition-construction-stack){#html-element-constructors:concept-custom-element-definition-construction-stack}
    is empty:

    1.  Let `element`{.variable} be the result of [internally creating a
        new object implementing the
        interface](https://webidl.spec.whatwg.org/#internally-create-a-new-object-implementing-the-interface){#html-element-constructors:internally-create-a-new-object-implementing-the-interface
        x-internal="internally-create-a-new-object-implementing-the-interface"}
        to which the [active function
        object](https://tc39.es/ecma262/#active-function-object){#html-element-constructors:active-function-object-6
        x-internal="active-function-object"} corresponds, given the
        [current
        realm](https://tc39.es/ecma262/#current-realm){#html-element-constructors:current-realm
        x-internal="current-realm"} and
        [NewTarget](https://tc39.es/ecma262/#sec-built-in-function-objects){#html-element-constructors:newtarget-8
        x-internal="newtarget"}.

    2.  Set `element`{.variable}\'s [node
        document](https://dom.spec.whatwg.org/#concept-node-document){#html-element-constructors:node-document
        x-internal="node-document"} to the [current global
        object](webappapis.html#current-global-object){#html-element-constructors:current-global-object-2}\'s
        [associated
        `Document`](nav-history-apis.html#concept-document-window){#html-element-constructors:concept-document-window-2}.

    3.  Set `element`{.variable}\'s
        [namespace](https://dom.spec.whatwg.org/#concept-element-namespace){#html-element-constructors:concept-element-namespace
        x-internal="concept-element-namespace"} to the [HTML
        namespace](https://infra.spec.whatwg.org/#html-namespace){#html-element-constructors:html-namespace-2
        x-internal="html-namespace-2"}.

    4.  Set `element`{.variable}\'s [namespace
        prefix](https://dom.spec.whatwg.org/#concept-element-namespace-prefix){#html-element-constructors:concept-element-namespace-prefix
        x-internal="concept-element-namespace-prefix"} to null.

    5.  Set `element`{.variable}\'s [local
        name](https://dom.spec.whatwg.org/#concept-element-local-name){#html-element-constructors:concept-element-local-name
        x-internal="concept-element-local-name"} to
        `definition`{.variable}\'s [local
        name](custom-elements.html#concept-custom-element-definition-local-name){#html-element-constructors:concept-custom-element-definition-local-name-4}.

    6.  Set `element`{.variable}\'s [custom element
        state](https://dom.spec.whatwg.org/#concept-element-custom-element-state){#html-element-constructors:custom-element-state
        x-internal="custom-element-state"} to \"`custom`\".

    7.  Set `element`{.variable}\'s [custom element
        definition](https://dom.spec.whatwg.org/#concept-element-custom-element-definition){#html-element-constructors:concept-element-custom-element-definition
        x-internal="concept-element-custom-element-definition"} to
        `definition`{.variable}.

    8.  Set `element`{.variable}\'s [`is`
        value](https://dom.spec.whatwg.org/#concept-element-is-value){#html-element-constructors:concept-element-is-value
        x-internal="concept-element-is-value"} to `isValue`{.variable}.

    9.  Return `element`{.variable}.

    This occurs when author script constructs a new custom element
    directly, e.g. via `new MyCustomElement()`.

10. Let `prototype`{.variable} be ?
    [Get](https://tc39.es/ecma262/#sec-get-o-p){#html-element-constructors:js-get
    x-internal="js-get"}([NewTarget](https://tc39.es/ecma262/#sec-built-in-function-objects){#html-element-constructors:newtarget-9
    x-internal="newtarget"}, \"prototype\").

11. If `prototype`{.variable} [is not an
    Object](https://tc39.es/ecma262/#sec-object-type){#html-element-constructors:js-object
    x-internal="js-object"}, then:

    1.  Let `realm`{.variable} be ?
        [GetFunctionRealm](https://tc39.es/ecma262/#sec-getfunctionrealm){#html-element-constructors:getfunctionrealm
        x-internal="getfunctionrealm"}([NewTarget](https://tc39.es/ecma262/#sec-built-in-function-objects){#html-element-constructors:newtarget-10
        x-internal="newtarget"}).

    2.  Set `prototype`{.variable} to the [interface prototype
        object](https://webidl.spec.whatwg.org/#dfn-interface-prototype-object){#html-element-constructors:interface-prototype-object
        x-internal="interface-prototype-object"} of `realm`{.variable}
        whose interface is the same as the interface of the [active
        function
        object](https://tc39.es/ecma262/#active-function-object){#html-element-constructors:active-function-object-7
        x-internal="active-function-object"}.

    The realm of the [active function
    object](https://tc39.es/ecma262/#active-function-object){#html-element-constructors:active-function-object-8
    x-internal="active-function-object"} might not be
    `realm`{.variable}, so we are using the more general concept of
    \"the same interface\" across realms; we are not looking for
    equality of [interface
    objects](https://webidl.spec.whatwg.org/#dfn-interface-object){#html-element-constructors:interface-object
    x-internal="interface-object"}. This fallback behavior, including
    using the realm of
    [NewTarget](https://tc39.es/ecma262/#sec-built-in-function-objects){#html-element-constructors:newtarget-11
    x-internal="newtarget"} and looking up the appropriate prototype
    there, is designed to match analogous behavior for the JavaScript
    built-ins and Web IDL\'s [internally create a new object
    implementing the
    interface](https://webidl.spec.whatwg.org/#internally-create-a-new-object-implementing-the-interface){#html-element-constructors:internally-create-a-new-object-implementing-the-interface-2
    x-internal="internally-create-a-new-object-implementing-the-interface"}
    algorithm.

12. Let `element`{.variable} be the last entry in
    `definition`{.variable}\'s [construction
    stack](custom-elements.html#concept-custom-element-definition-construction-stack){#html-element-constructors:concept-custom-element-definition-construction-stack-2}.

13. If `element`{.variable} is an [*already constructed*
    marker](custom-elements.html#concept-already-constructed-marker){#html-element-constructors:concept-already-constructed-marker},
    then throw a
    [`TypeError`{#html-element-constructors:typeerror-5}](https://tc39.es/ecma262/#sec-native-error-types-used-in-this-standard-typeerror){x-internal="typeerror"}.

    ::: example
    This can occur when the author code inside the [custom element
    constructor](custom-elements.html#custom-element-constructor){#html-element-constructors:custom-element-constructor}
    [non-conformantly](custom-elements.html#custom-element-conformance)
    creates another instance of the class being constructed, before
    calling `super()`:

    ``` js
    let doSillyThing = true;

    class DontDoThis extends HTMLElement {
      constructor() {
        if (doSillyThing) {
          doSillyThing = false;
          new DontDoThis();
          // Now the construction stack will contain an already constructed marker.
        }

        // This will then fail with a TypeError:
        super();
      }
    }
    ```
    :::

    ::: example
    This can also occur when author code inside the [custom element
    constructor](custom-elements.html#custom-element-constructor){#html-element-constructors:custom-element-constructor-2}
    [non-conformantly](custom-elements.html#custom-element-conformance)
    calls `super()` twice, since per the JavaScript specification, this
    actually executes the superclass constructor (i.e. this algorithm)
    twice, before throwing an error:

    ``` js
    class DontDoThisEither extends HTMLElement {
      constructor() {
        super();

        // This will throw, but not until it has already called into the HTMLElement constructor
        super();
      }
    }
    ```
    :::

14. Perform ?
    `element`{.variable}.\[\[SetPrototypeOf\]\](`prototype`{.variable}).

15. Replace the last entry in `definition`{.variable}\'s [construction
    stack](custom-elements.html#concept-custom-element-definition-construction-stack){#html-element-constructors:concept-custom-element-definition-construction-stack-3}
    with an [*already constructed*
    marker](custom-elements.html#concept-already-constructed-marker){#html-element-constructors:concept-already-constructed-marker-2}.

16. Return `element`{.variable}.

    This step is normally reached when
    [upgrading](custom-elements.html#upgrades){#html-element-constructors:upgrades}
    a custom element; the existing element is returned, so that the
    `super()` call inside the [custom element
    constructor](custom-elements.html#custom-element-constructor){#html-element-constructors:custom-element-constructor-3}
    assigns that existing element to **this**.

------------------------------------------------------------------------

In addition to the constructor behavior implied by
[`[HTMLConstructor]`{#html-element-constructors:htmlconstructor-3}](#htmlconstructor),
some elements also have [named
constructors](https://webidl.spec.whatwg.org/#dfn-named-constructor){#html-element-constructors:named-constructor
x-internal="named-constructor"} (which are really factory functions with
a modified `prototype` property).

::: example
Named constructors for HTML elements can also be used in an
[`extends`{#html-element-constructors:dom-elementdefinitionoptions-extends}](custom-elements.html#dom-elementdefinitionoptions-extends)
clause when defining a [custom element
constructor](custom-elements.html#custom-element-constructor){#html-element-constructors:custom-element-constructor-4}:

``` js
class AutoEmbiggenedImage extends Image {
  constructor(width, height) {
    super(width * 10, height * 10);
  }
}

customElements.define("auto-embiggened", AutoEmbiggenedImage, { extends: "img" });

const image = new AutoEmbiggenedImage(15, 20);
console.assert(image.width === 150);
console.assert(image.height === 200);
```
:::

#### [3.2.4]{.secno} Element definitions[](#element-definitions){.self-link}

Each element in this specification has a definition that includes the
following information:

[Categories]{#concept-element-categories .dfn}
:   A list of
    [categories](#content-categories){#element-definitions:content-categories}
    to which the element belongs. These are used when defining the
    [content
    models](#content-models){#element-definitions:content-models} for
    each element.

[Contexts in which this element can be used]{#concept-element-contexts .dfn}

:   A *non-normative* description of where the element can be used. This
    information is redundant with the content models of elements that
    allow this one as a child, and is provided only as a convenience.

    ::: note
    For simplicity, only the most specific expectations are listed.

    For example, all [phrasing
    content](#phrasing-content-2){#element-definitions:phrasing-content-2}
    is [flow
    content](#flow-content-2){#element-definitions:flow-content-2}.
    Thus, elements that are [phrasing
    content](#phrasing-content-2){#element-definitions:phrasing-content-2-2}
    will only be listed as \"where [phrasing
    content](#phrasing-content-2){#element-definitions:phrasing-content-2-3}
    is expected\", since this is the more-specific expectation. Anywhere
    that expects [flow
    content](#flow-content-2){#element-definitions:flow-content-2-2}
    also expects [phrasing
    content](#phrasing-content-2){#element-definitions:phrasing-content-2-4},
    and thus also meets this expectation.
    :::

[Content model]{#concept-element-content-model .dfn}
:   A normative description of what content must be included as children
    and descendants of the element.

[Tag omission in text/html]{#concept-element-tag-omission .dfn}
:   A *non-normative* description of whether, in the
    [`text/html`{#element-definitions:text/html}](iana.html#text/html)
    syntax, the
    [start](syntax.html#syntax-start-tag){#element-definitions:syntax-start-tag}
    and
    [end](syntax.html#syntax-end-tag){#element-definitions:syntax-end-tag}
    tags can be omitted. This information is redundant with the
    normative requirements given in the [optional
    tags](syntax.html#syntax-tag-omission){#element-definitions:syntax-tag-omission}
    section, and is provided in the element definitions only as a
    convenience.

[Content attributes]{#concept-element-attributes .dfn}
:   A normative list of attributes that may be specified on the element
    (except where otherwise disallowed), along with non-normative
    descriptions of those attributes. (The content to the left of the
    dash is normative, the content to the right of the dash is not.)

[Accessibility considerations]{#concept-element-accessibility-considerations .dfn}

:   For authors: Conformance requirements for use of ARIA
    [`role`{#element-definitions:attr-aria-role}](infrastructure.html#attr-aria-role)
    and
    [`aria-*`{#element-definitions:attr-aria-*}](infrastructure.html#attr-aria-*)
    attributes are defined in ARIA in HTML.
    [\[ARIA\]](references.html#refsARIA)
    [\[ARIAHTML\]](references.html#refsARIAHTML)

    For implementers: User agent requirements for implementing
    accessibility API semantics are defined in HTML Accessibility API
    Mappings. [\[HTMLAAM\]](references.html#refsHTMLAAM)

[DOM interface]{#concept-element-dom .dfn}

:   A normative definition of a DOM interface that such elements must
    implement.

This is then followed by a description of what the element
[represents](#represents){#element-definitions:represents}, along with
any additional normative conformance criteria that may apply to authors
and implementations. Examples are sometimes also included.

##### [3.2.4.1]{.secno} Attributes[](#attributes){.self-link}

An attribute value is a string. Except where otherwise specified,
attribute values on [HTML
elements](infrastructure.html#html-elements){#attributes:html-elements}
may be any string value, including the empty string, and there is no
restriction on what text can be specified in such attribute values.

#### [3.2.5]{.secno} [Content models]{.dfn}[](#content-models){.self-link}

Each element defined in this specification has a content model: a
description of the element\'s expected
[contents](#concept-html-contents){#content-models:concept-html-contents}.
An [HTML
element](infrastructure.html#html-elements){#content-models:html-elements}
must have contents that match the requirements described in the
element\'s content model. The [contents]{#concept-html-contents .dfn} of
an element are its children in the DOM.

[ASCII
whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#content-models:space-characters
x-internal="space-characters"} is always allowed between elements. User
agents represent these characters between elements in the source markup
as
[`Text`{#content-models:text}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
nodes in the DOM. Empty
[`Text`{#content-models:text-2}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
nodes and
[`Text`{#content-models:text-3}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
nodes consisting of just sequences of those characters are considered
[inter-element whitespace]{#inter-element-whitespace .dfn}.

[Inter-element
whitespace](#inter-element-whitespace){#content-models:inter-element-whitespace},
comment nodes, and processing instruction nodes must be ignored when
establishing whether an element\'s contents match the element\'s content
model or not, and must be ignored when following algorithms that define
document and element semantics.

Thus, an element `A`{.variable} is said to be *preceded or followed* by
a second element `B`{.variable} if `A`{.variable} and `B`{.variable}
have the same parent node and there are no other element nodes or
[`Text`{#content-models:text-4}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
nodes (other than [inter-element
whitespace](#inter-element-whitespace){#content-models:inter-element-whitespace-2})
between them. Similarly, a node is the *only child* of an element if
that element contains no other nodes other than [inter-element
whitespace](#inter-element-whitespace){#content-models:inter-element-whitespace-3},
comment nodes, and processing instruction nodes.

Authors must not use [HTML
elements](infrastructure.html#html-elements){#content-models:html-elements-2}
anywhere except where they are explicitly allowed, as defined for each
element, or as explicitly required by other specifications. For XML
compound documents, these contexts could be inside elements from other
namespaces, if those elements are defined as providing the relevant
contexts.

The Atom Syndication Format defines a `content` element. When its `type`
attribute has the value `xhtml`, The Atom Syndication Format requires
that it contain a single HTML
[`div`{#content-models:the-div-element}](grouping-content.html#the-div-element)
element. Thus, a
[`div`{#content-models:the-div-element-2}](grouping-content.html#the-div-element)
element is allowed in that context, even though this is not explicitly
normatively stated by this specification.
[\[ATOM\]](references.html#refsATOM)

In addition, [HTML
elements](infrastructure.html#html-elements){#content-models:html-elements-3}
may be orphan nodes (i.e. without a parent node).

::: example
For example, creating a
[`td`{#content-models:the-td-element}](tables.html#the-td-element)
element and storing it in a global variable in a script is conforming,
even though
[`td`{#content-models:the-td-element-2}](tables.html#the-td-element)
elements are otherwise only supposed to be used inside
[`tr`{#content-models:the-tr-element}](tables.html#the-tr-element)
elements.

``` js
var data = {
  name: "Banana",
  cell: document.createElement('td'),
};
```
:::

##### [3.2.5.1]{.secno} The \"nothing\" content model[](#the-nothing-content-model){.self-link}

When an element\'s content model is [nothing]{#concept-content-nothing
.dfn}, the element must contain no
[`Text`{#the-nothing-content-model:text}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
nodes (other than [inter-element
whitespace](#inter-element-whitespace){#the-nothing-content-model:inter-element-whitespace})
and no element nodes.

Most HTML elements whose content model is \"nothing\" are also, for
convenience, [void
elements](syntax.html#void-elements){#the-nothing-content-model:void-elements}
(elements that have no [end
tag](syntax.html#syntax-end-tag){#the-nothing-content-model:syntax-end-tag}
in the [HTML syntax](syntax.html#syntax)). However, these are entirely
separate concepts.

##### [3.2.5.2]{.secno} Kinds of content[](#kinds-of-content){.self-link}

Each element in HTML falls into zero or more
[categories]{#content-categories .dfn} that group elements with similar
characteristics together. The following broad categories are used in
this specification:

- [Metadata
  content](#metadata-content-2){#kinds-of-content:metadata-content-2}
- [Flow content](#flow-content-2){#kinds-of-content:flow-content-2}
- [Sectioning
  content](#sectioning-content-2){#kinds-of-content:sectioning-content-2}
- [Heading
  content](#heading-content-2){#kinds-of-content:heading-content-2}
- [Phrasing
  content](#phrasing-content-2){#kinds-of-content:phrasing-content-2}
- [Embedded
  content](#embedded-content-category){#kinds-of-content:embedded-content-category}
- [Interactive
  content](#interactive-content-2){#kinds-of-content:interactive-content-2}

Some elements also fall into other categories, which are defined in
other parts of this specification.

These categories are related as follows:

Sectioning content, heading content, phrasing content, embedded content,
and interactive content are all types of flow content. Metadata is
sometimes flow content. Metadata and interactive content are sometimes
phrasing content. Embedded content is also a type of phrasing content,
and sometimes is interactive content.

Other categories are also used for specific purposes, e.g. form controls
are specified using a number of categories to define common
requirements. Some elements have unique requirements and do not fit into
any particular category.

###### [3.2.5.2.1]{.secno} Metadata content[](#metadata-content){.self-link}

[Metadata content]{#metadata-content-2 .dfn export=""} is content that
sets up the presentation or behavior of the rest of the content, or that
sets up the relationship of the document with other documents, or that
conveys other \"out of band\" information.

- [`base`{#metadata-content:the-base-element}](semantics.html#the-base-element)
- [`link`{#metadata-content:the-link-element}](semantics.html#the-link-element)
- [`meta`{#metadata-content:the-meta-element}](semantics.html#the-meta-element)
- [`noscript`{#metadata-content:the-noscript-element}](scripting.html#the-noscript-element)
- [`script`{#metadata-content:the-script-element}](scripting.html#the-script-element)
- [`style`{#metadata-content:the-style-element}](semantics.html#the-style-element)
- [`template`{#metadata-content:the-template-element}](scripting.html#the-template-element)
- [`title`{#metadata-content:the-title-element}](semantics.html#the-title-element)

Elements from other namespaces whose semantics are primarily
metadata-related (e.g. RDF) are also [metadata
content](#metadata-content-2){#metadata-content:metadata-content-2}.

::: example
Thus, in the XML serialization, one can use RDF, like this:

``` html
<html xmlns="http://www.w3.org/1999/xhtml"
      xmlns:r="http://www.w3.org/1999/02/22-rdf-syntax-ns#" xml:lang="en">
 <head>
  <title>Hedral's Home Page</title>
  <r:RDF>
   <Person xmlns="http://www.w3.org/2000/10/swap/pim/contact#"
           r:about="https://hedral.example.com/#">
    <fullName>Cat Hedral</fullName>
    <mailbox r:resource="mailto:hedral@damowmow.com"/>
    <personalTitle>Sir</personalTitle>
   </Person>
  </r:RDF>
 </head>
 <body>
  <h1>My home page</h1>
  <p>I like playing with string, I guess. Sister says squirrels are fun
  too so sometimes I follow her to play with them.</p>
 </body>
</html>
```

This isn\'t possible in the HTML serialization, however.
:::

###### [3.2.5.2.2]{.secno} Flow content[](#flow-content){.self-link}

Most elements that are used in the body of documents and applications
are categorized as [flow content]{#flow-content-2 .dfn export=""}.

- [`a`{#flow-content:the-a-element}](text-level-semantics.html#the-a-element)
- [`abbr`{#flow-content:the-abbr-element}](text-level-semantics.html#the-abbr-element)
- [`address`{#flow-content:the-address-element}](sections.html#the-address-element)
- [`area`{#flow-content:the-area-element}](image-maps.html#the-area-element)
  (if it is a descendant of a
  [`map`{#flow-content:the-map-element}](image-maps.html#the-map-element)
  element)
- [`article`{#flow-content:the-article-element}](sections.html#the-article-element)
- [`aside`{#flow-content:the-aside-element}](sections.html#the-aside-element)
- [`audio`{#flow-content:the-audio-element}](media.html#the-audio-element)
- [`b`{#flow-content:the-b-element}](text-level-semantics.html#the-b-element)
- [`bdi`{#flow-content:the-bdi-element}](text-level-semantics.html#the-bdi-element)
- [`bdo`{#flow-content:the-bdo-element}](text-level-semantics.html#the-bdo-element)
- [`blockquote`{#flow-content:the-blockquote-element}](grouping-content.html#the-blockquote-element)
- [`br`{#flow-content:the-br-element}](text-level-semantics.html#the-br-element)
- [`button`{#flow-content:the-button-element}](form-elements.html#the-button-element)
- [`canvas`{#flow-content:the-canvas-element}](canvas.html#the-canvas-element)
- [`cite`{#flow-content:the-cite-element}](text-level-semantics.html#the-cite-element)
- [`code`{#flow-content:the-code-element}](text-level-semantics.html#the-code-element)
- [`data`{#flow-content:the-data-element}](text-level-semantics.html#the-data-element)
- [`datalist`{#flow-content:the-datalist-element}](form-elements.html#the-datalist-element)
- [`del`{#flow-content:the-del-element}](edits.html#the-del-element)
- [`details`{#flow-content:the-details-element}](interactive-elements.html#the-details-element)
- [`dfn`{#flow-content:the-dfn-element}](text-level-semantics.html#the-dfn-element)
- [`dialog`{#flow-content:the-dialog-element}](interactive-elements.html#the-dialog-element)
- [`div`{#flow-content:the-div-element}](grouping-content.html#the-div-element)
- [`dl`{#flow-content:the-dl-element}](grouping-content.html#the-dl-element)
- [`em`{#flow-content:the-em-element}](text-level-semantics.html#the-em-element)
- [`embed`{#flow-content:the-embed-element}](iframe-embed-object.html#the-embed-element)
- [`fieldset`{#flow-content:the-fieldset-element}](form-elements.html#the-fieldset-element)
- [`figure`{#flow-content:the-figure-element}](grouping-content.html#the-figure-element)
- [`footer`{#flow-content:the-footer-element}](sections.html#the-footer-element)
- [`form`{#flow-content:the-form-element}](forms.html#the-form-element)
- [`h1`{#flow-content:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements)
- [`h2`{#flow-content:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-2}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements)
- [`h3`{#flow-content:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-3}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements)
- [`h4`{#flow-content:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-4}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements)
- [`h5`{#flow-content:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-5}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements)
- [`h6`{#flow-content:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-6}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements)
- [`header`{#flow-content:the-header-element}](sections.html#the-header-element)
- [`hgroup`{#flow-content:the-hgroup-element}](sections.html#the-hgroup-element)
- [`hr`{#flow-content:the-hr-element}](grouping-content.html#the-hr-element)
- [`i`{#flow-content:the-i-element}](text-level-semantics.html#the-i-element)
- [`iframe`{#flow-content:the-iframe-element}](iframe-embed-object.html#the-iframe-element)
- [`img`{#flow-content:the-img-element}](embedded-content.html#the-img-element)
- [`input`{#flow-content:the-input-element}](input.html#the-input-element)
- [`ins`{#flow-content:the-ins-element}](edits.html#the-ins-element)
- [`kbd`{#flow-content:the-kbd-element}](text-level-semantics.html#the-kbd-element)
- [`label`{#flow-content:the-label-element}](forms.html#the-label-element)
- [`link`{#flow-content:the-link-element}](semantics.html#the-link-element)
  (if it is [allowed in the
  body](semantics.html#allowed-in-the-body){#flow-content:allowed-in-the-body})
- [`main`{#flow-content:the-main-element}](grouping-content.html#the-main-element)
  (if it is a [hierarchically correct `main`
  element](grouping-content.html#hierarchically-correct-main-element){#flow-content:hierarchically-correct-main-element})
- [`map`{#flow-content:the-map-element-2}](image-maps.html#the-map-element)
- [`mark`{#flow-content:the-mark-element}](text-level-semantics.html#the-mark-element)
- [MathML
  `math`](https://w3c.github.io/mathml-core/#the-top-level-math-element){#flow-content:mathml-math
  x-internal="mathml-math"}
- [`menu`{#flow-content:the-menu-element}](grouping-content.html#the-menu-element)
- [`meta`{#flow-content:the-meta-element}](semantics.html#the-meta-element)
  (if the
  [`itemprop`{#flow-content:names:-the-itemprop-attribute}](microdata.html#names:-the-itemprop-attribute)
  attribute is present)
- [`meter`{#flow-content:the-meter-element}](form-elements.html#the-meter-element)
- [`nav`{#flow-content:the-nav-element}](sections.html#the-nav-element)
- [`noscript`{#flow-content:the-noscript-element}](scripting.html#the-noscript-element)
- [`object`{#flow-content:the-object-element}](iframe-embed-object.html#the-object-element)
- [`ol`{#flow-content:the-ol-element}](grouping-content.html#the-ol-element)
- [`output`{#flow-content:the-output-element}](form-elements.html#the-output-element)
- [`p`{#flow-content:the-p-element}](grouping-content.html#the-p-element)
- [`picture`{#flow-content:the-picture-element}](embedded-content.html#the-picture-element)
- [`pre`{#flow-content:the-pre-element}](grouping-content.html#the-pre-element)
- [`progress`{#flow-content:the-progress-element}](form-elements.html#the-progress-element)
- [`q`{#flow-content:the-q-element}](text-level-semantics.html#the-q-element)
- [`ruby`{#flow-content:the-ruby-element}](text-level-semantics.html#the-ruby-element)
- [`s`{#flow-content:the-s-element}](text-level-semantics.html#the-s-element)
- [`samp`{#flow-content:the-samp-element}](text-level-semantics.html#the-samp-element)
- [`script`{#flow-content:the-script-element}](scripting.html#the-script-element)
- [`search`{#flow-content:the-search-element}](grouping-content.html#the-search-element)
- [`section`{#flow-content:the-section-element}](sections.html#the-section-element)
- [`select`{#flow-content:the-select-element}](form-elements.html#the-select-element)
- [`slot`{#flow-content:the-slot-element}](scripting.html#the-slot-element)
- [`small`{#flow-content:the-small-element}](text-level-semantics.html#the-small-element)
- [`span`{#flow-content:the-span-element}](text-level-semantics.html#the-span-element)
- [`strong`{#flow-content:the-strong-element}](text-level-semantics.html#the-strong-element)
- [`sub`{#flow-content:the-sub-and-sup-elements}](text-level-semantics.html#the-sub-and-sup-elements)
- [`sup`{#flow-content:the-sub-and-sup-elements-2}](text-level-semantics.html#the-sub-and-sup-elements)
- [SVG
  `svg`](https://svgwg.org/svg2-draft/struct.html#SVGElement){#flow-content:svg-svg
  x-internal="svg-svg"}
- [`table`{#flow-content:the-table-element}](tables.html#the-table-element)
- [`template`{#flow-content:the-template-element}](scripting.html#the-template-element)
- [`textarea`{#flow-content:the-textarea-element}](form-elements.html#the-textarea-element)
- [`time`{#flow-content:the-time-element}](text-level-semantics.html#the-time-element)
- [`u`{#flow-content:the-u-element}](text-level-semantics.html#the-u-element)
- [`ul`{#flow-content:the-ul-element}](grouping-content.html#the-ul-element)
- [`var`{#flow-content:the-var-element}](text-level-semantics.html#the-var-element)
- [`video`{#flow-content:the-video-element}](media.html#the-video-element)
- [`wbr`{#flow-content:the-wbr-element}](text-level-semantics.html#the-wbr-element)
- [autonomous custom
  elements](custom-elements.html#autonomous-custom-element){#flow-content:autonomous-custom-element}
- [text](#text-content){#flow-content:text-content}

###### [3.2.5.2.3]{.secno} Sectioning content[](#sectioning-content){.self-link}

[Sectioning content]{#sectioning-content-2 .dfn export=""} is content
that defines the scope of
[`header`{#sectioning-content:the-header-element}](sections.html#the-header-element)
and
[`footer`{#sectioning-content:the-footer-element}](sections.html#the-footer-element)
elements.

- [`article`{#sectioning-content:the-article-element}](sections.html#the-article-element)
- [`aside`{#sectioning-content:the-aside-element}](sections.html#the-aside-element)
- [`nav`{#sectioning-content:the-nav-element}](sections.html#the-nav-element)
- [`section`{#sectioning-content:the-section-element}](sections.html#the-section-element)

###### [3.2.5.2.4]{.secno} Heading content[](#heading-content){.self-link}

[Heading content]{#heading-content-2 .dfn export=""} defines the heading
of a section (whether explicitly marked up using [sectioning
content](#sectioning-content-2){#heading-content:sectioning-content-2}
elements, or implied by the heading content itself).

- [`h1`{#heading-content:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements)
- [`h2`{#heading-content:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-2}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements)
- [`h3`{#heading-content:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-3}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements)
- [`h4`{#heading-content:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-4}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements)
- [`h5`{#heading-content:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-5}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements)
- [`h6`{#heading-content:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-6}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements)
- [`hgroup`{#heading-content:the-hgroup-element}](sections.html#the-hgroup-element)
  (if it has a descendant
  [`h1`{#heading-content:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-7}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements)
  to
  [`h6`{#heading-content:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-8}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements)
  element)

###### [3.2.5.2.5]{.secno} Phrasing content[](#phrasing-content){.self-link}

[Phrasing content]{#phrasing-content-2 .dfn export=""} is the text of
the document, as well as elements that mark up that text at the
intra-paragraph level. Runs of [phrasing
content](#phrasing-content-2){#phrasing-content:phrasing-content-2} form
[paragraphs](#paragraph){#phrasing-content:paragraph}.

- [`a`{#phrasing-content:the-a-element}](text-level-semantics.html#the-a-element)
- [`abbr`{#phrasing-content:the-abbr-element}](text-level-semantics.html#the-abbr-element)
- [`area`{#phrasing-content:the-area-element}](image-maps.html#the-area-element)
  (if it is a descendant of a
  [`map`{#phrasing-content:the-map-element}](image-maps.html#the-map-element)
  element)
- [`audio`{#phrasing-content:the-audio-element}](media.html#the-audio-element)
- [`b`{#phrasing-content:the-b-element}](text-level-semantics.html#the-b-element)
- [`bdi`{#phrasing-content:the-bdi-element}](text-level-semantics.html#the-bdi-element)
- [`bdo`{#phrasing-content:the-bdo-element}](text-level-semantics.html#the-bdo-element)
- [`br`{#phrasing-content:the-br-element}](text-level-semantics.html#the-br-element)
- [`button`{#phrasing-content:the-button-element}](form-elements.html#the-button-element)
- [`canvas`{#phrasing-content:the-canvas-element}](canvas.html#the-canvas-element)
- [`cite`{#phrasing-content:the-cite-element}](text-level-semantics.html#the-cite-element)
- [`code`{#phrasing-content:the-code-element}](text-level-semantics.html#the-code-element)
- [`data`{#phrasing-content:the-data-element}](text-level-semantics.html#the-data-element)
- [`datalist`{#phrasing-content:the-datalist-element}](form-elements.html#the-datalist-element)
- [`del`{#phrasing-content:the-del-element}](edits.html#the-del-element)
- [`dfn`{#phrasing-content:the-dfn-element}](text-level-semantics.html#the-dfn-element)
- [`em`{#phrasing-content:the-em-element}](text-level-semantics.html#the-em-element)
- [`embed`{#phrasing-content:the-embed-element}](iframe-embed-object.html#the-embed-element)
- [`i`{#phrasing-content:the-i-element}](text-level-semantics.html#the-i-element)
- [`iframe`{#phrasing-content:the-iframe-element}](iframe-embed-object.html#the-iframe-element)
- [`img`{#phrasing-content:the-img-element}](embedded-content.html#the-img-element)
- [`input`{#phrasing-content:the-input-element}](input.html#the-input-element)
- [`ins`{#phrasing-content:the-ins-element}](edits.html#the-ins-element)
- [`kbd`{#phrasing-content:the-kbd-element}](text-level-semantics.html#the-kbd-element)
- [`label`{#phrasing-content:the-label-element}](forms.html#the-label-element)
- [`link`{#phrasing-content:the-link-element}](semantics.html#the-link-element)
  (if it is [allowed in the
  body](semantics.html#allowed-in-the-body){#phrasing-content:allowed-in-the-body})
- [`map`{#phrasing-content:the-map-element-2}](image-maps.html#the-map-element)
- [`mark`{#phrasing-content:the-mark-element}](text-level-semantics.html#the-mark-element)
- [MathML
  `math`](https://w3c.github.io/mathml-core/#the-top-level-math-element){#phrasing-content:mathml-math
  x-internal="mathml-math"}
- [`meta`{#phrasing-content:the-meta-element}](semantics.html#the-meta-element)
  (if the
  [`itemprop`{#phrasing-content:names:-the-itemprop-attribute}](microdata.html#names:-the-itemprop-attribute)
  attribute is present)
- [`meter`{#phrasing-content:the-meter-element}](form-elements.html#the-meter-element)
- [`noscript`{#phrasing-content:the-noscript-element}](scripting.html#the-noscript-element)
- [`object`{#phrasing-content:the-object-element}](iframe-embed-object.html#the-object-element)
- [`output`{#phrasing-content:the-output-element}](form-elements.html#the-output-element)
- [`picture`{#phrasing-content:the-picture-element}](embedded-content.html#the-picture-element)
- [`progress`{#phrasing-content:the-progress-element}](form-elements.html#the-progress-element)
- [`q`{#phrasing-content:the-q-element}](text-level-semantics.html#the-q-element)
- [`ruby`{#phrasing-content:the-ruby-element}](text-level-semantics.html#the-ruby-element)
- [`s`{#phrasing-content:the-s-element}](text-level-semantics.html#the-s-element)
- [`samp`{#phrasing-content:the-samp-element}](text-level-semantics.html#the-samp-element)
- [`script`{#phrasing-content:the-script-element}](scripting.html#the-script-element)
- [`select`{#phrasing-content:the-select-element}](form-elements.html#the-select-element)
- [`slot`{#phrasing-content:the-slot-element}](scripting.html#the-slot-element)
- [`small`{#phrasing-content:the-small-element}](text-level-semantics.html#the-small-element)
- [`span`{#phrasing-content:the-span-element}](text-level-semantics.html#the-span-element)
- [`strong`{#phrasing-content:the-strong-element}](text-level-semantics.html#the-strong-element)
- [`sub`{#phrasing-content:the-sub-and-sup-elements}](text-level-semantics.html#the-sub-and-sup-elements)
- [`sup`{#phrasing-content:the-sub-and-sup-elements-2}](text-level-semantics.html#the-sub-and-sup-elements)
- [SVG
  `svg`](https://svgwg.org/svg2-draft/struct.html#SVGElement){#phrasing-content:svg-svg
  x-internal="svg-svg"}
- [`template`{#phrasing-content:the-template-element}](scripting.html#the-template-element)
- [`textarea`{#phrasing-content:the-textarea-element}](form-elements.html#the-textarea-element)
- [`time`{#phrasing-content:the-time-element}](text-level-semantics.html#the-time-element)
- [`u`{#phrasing-content:the-u-element}](text-level-semantics.html#the-u-element)
- [`var`{#phrasing-content:the-var-element}](text-level-semantics.html#the-var-element)
- [`video`{#phrasing-content:the-video-element}](media.html#the-video-element)
- [`wbr`{#phrasing-content:the-wbr-element}](text-level-semantics.html#the-wbr-element)
- [autonomous custom
  elements](custom-elements.html#autonomous-custom-element){#phrasing-content:autonomous-custom-element}
- [text](#text-content){#phrasing-content:text-content}

Most elements that are categorized as phrasing content can only contain
elements that are themselves categorized as phrasing content, not any
flow content.

[Text]{#text-content .dfn}, in the context of content models, means
either nothing, or
[`Text`{#phrasing-content:text}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
nodes. [Text](#text-content){#phrasing-content:text-content-2} is
sometimes used as a content model on its own, but is also [phrasing
content](#phrasing-content-2){#phrasing-content:phrasing-content-2-2},
and can be [inter-element
whitespace](#inter-element-whitespace){#phrasing-content:inter-element-whitespace}
(if the
[`Text`{#phrasing-content:text-2}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
nodes are empty or contain just [ASCII
whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#phrasing-content:space-characters
x-internal="space-characters"}).

[`Text`{#phrasing-content:text-3}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
nodes and attribute values must consist of [scalar
values](https://infra.spec.whatwg.org/#scalar-value){#phrasing-content:scalar-value
x-internal="scalar-value"}, excluding
[noncharacters](https://infra.spec.whatwg.org/#noncharacter){#phrasing-content:noncharacter
x-internal="noncharacter"}, and
[controls](https://infra.spec.whatwg.org/#control){#phrasing-content:control
x-internal="control"} other than [ASCII
whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#phrasing-content:space-characters-2
x-internal="space-characters"}. This specification includes extra
constraints on the exact value of
[`Text`{#phrasing-content:text-4}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
nodes and attribute values depending on their precise context.

###### [3.2.5.2.6]{.secno} Embedded content[](#embedded-content-2){.self-link} {#embedded-content-2}

[Embedded content]{#embedded-content-category .dfn export=""} is content
that imports another resource into the document, or content from another
vocabulary that is inserted into the document.

- [`audio`{#embedded-content-2:the-audio-element}](media.html#the-audio-element)
- [`canvas`{#embedded-content-2:the-canvas-element}](canvas.html#the-canvas-element)
- [`embed`{#embedded-content-2:the-embed-element}](iframe-embed-object.html#the-embed-element)
- [`iframe`{#embedded-content-2:the-iframe-element}](iframe-embed-object.html#the-iframe-element)
- [`img`{#embedded-content-2:the-img-element}](embedded-content.html#the-img-element)
- [MathML
  `math`](https://w3c.github.io/mathml-core/#the-top-level-math-element){#embedded-content-2:mathml-math
  x-internal="mathml-math"}
- [`object`{#embedded-content-2:the-object-element}](iframe-embed-object.html#the-object-element)
- [`picture`{#embedded-content-2:the-picture-element}](embedded-content.html#the-picture-element)
- [SVG
  `svg`](https://svgwg.org/svg2-draft/struct.html#SVGElement){#embedded-content-2:svg-svg
  x-internal="svg-svg"}
- [`video`{#embedded-content-2:the-video-element}](media.html#the-video-element)

Elements that are from namespaces other than the [HTML
namespace](https://infra.spec.whatwg.org/#html-namespace){#embedded-content-2:html-namespace-2
x-internal="html-namespace-2"} and that convey content but not metadata,
are [embedded
content](#embedded-content-category){#embedded-content-2:embedded-content-category}
for the purposes of the content models defined in this specification.
(For example, MathML or SVG.)

Some embedded content elements can have [fallback
content]{#fallback-content .dfn}: content that is to be used when the
external resource cannot be used (e.g. because it is of an unsupported
format). The element definitions state what the fallback is, if any.

###### [3.2.5.2.7]{.secno} Interactive content[](#interactive-content){.self-link}

[Interactive content]{#interactive-content-2 .dfn export=""} is content
that is specifically intended for user interaction.

- [`a`{#interactive-content:the-a-element}](text-level-semantics.html#the-a-element)
  (if the
  [`href`{#interactive-content:attr-hyperlink-href}](links.html#attr-hyperlink-href)
  attribute is present)
- [`audio`{#interactive-content:the-audio-element}](media.html#the-audio-element)
  (if the
  [`controls`{#interactive-content:attr-media-controls}](media.html#attr-media-controls)
  attribute is present)
- [`button`{#interactive-content:the-button-element}](form-elements.html#the-button-element)
- [`details`{#interactive-content:the-details-element}](interactive-elements.html#the-details-element)
- [`embed`{#interactive-content:the-embed-element}](iframe-embed-object.html#the-embed-element)
- [`iframe`{#interactive-content:the-iframe-element}](iframe-embed-object.html#the-iframe-element)
- [`img`{#interactive-content:the-img-element}](embedded-content.html#the-img-element)
  (if the
  [`usemap`{#interactive-content:attr-hyperlink-usemap}](image-maps.html#attr-hyperlink-usemap)
  attribute is present)
- [`input`{#interactive-content:the-input-element}](input.html#the-input-element)
  (if the
  [`type`{#interactive-content:attr-input-type}](input.html#attr-input-type)
  attribute is *not* in the
  [Hidden](input.html#hidden-state-(type=hidden)){#interactive-content:hidden-state-(type=hidden)}
  state)
- [`label`{#interactive-content:the-label-element}](forms.html#the-label-element)
- [`select`{#interactive-content:the-select-element}](form-elements.html#the-select-element)
- [`textarea`{#interactive-content:the-textarea-element}](form-elements.html#the-textarea-element)
- [`video`{#interactive-content:the-video-element}](media.html#the-video-element)
  (if the
  [`controls`{#interactive-content:attr-media-controls-2}](media.html#attr-media-controls)
  attribute is present)

###### [3.2.5.2.8]{.secno} Palpable content[](#palpable-content){.self-link}

As a general rule, elements whose content model allows any [flow
content](#flow-content-2){#palpable-content:flow-content-2} or [phrasing
content](#phrasing-content-2){#palpable-content:phrasing-content-2}
should have at least one node in its
[contents](#concept-html-contents){#palpable-content:concept-html-contents}
that is [palpable content]{#palpable-content-2 .dfn export=""} and that
does not have the
[`hidden`{#palpable-content:attr-hidden}](interaction.html#attr-hidden)
attribute specified.

[Palpable
content](#palpable-content-2){#palpable-content:palpable-content-2}
makes an element non-empty by providing either some descendant non-empty
[text](#text-content){#palpable-content:text-content}, or else something
users can hear
([`audio`{#palpable-content:the-audio-element}](media.html#the-audio-element)
elements) or view
([`video`{#palpable-content:the-video-element}](media.html#the-video-element),
[`img`{#palpable-content:the-img-element}](embedded-content.html#the-img-element),
or
[`canvas`{#palpable-content:the-canvas-element}](canvas.html#the-canvas-element)
elements) or otherwise interact with (for example, interactive form
controls).

This requirement is not a hard requirement, however, as there are many
cases where an element can be empty legitimately, for example when it is
used as a placeholder which will later be filled in by a script, or when
the element is part of a template and would on most pages be filled in
but on some pages is not relevant.

Conformance checkers are encouraged to provide a mechanism for authors
to find elements that fail to fulfill this requirement, as an authoring
aid.

The following elements are palpable content:

- [`a`{#palpable-content:the-a-element}](text-level-semantics.html#the-a-element)
- [`abbr`{#palpable-content:the-abbr-element}](text-level-semantics.html#the-abbr-element)
- [`address`{#palpable-content:the-address-element}](sections.html#the-address-element)
- [`article`{#palpable-content:the-article-element}](sections.html#the-article-element)
- [`aside`{#palpable-content:the-aside-element}](sections.html#the-aside-element)
- [`audio`{#palpable-content:the-audio-element-2}](media.html#the-audio-element)
  (if the
  [`controls`{#palpable-content:attr-media-controls}](media.html#attr-media-controls)
  attribute is present)
- [`b`{#palpable-content:the-b-element}](text-level-semantics.html#the-b-element)
- [`bdi`{#palpable-content:the-bdi-element}](text-level-semantics.html#the-bdi-element)
- [`bdo`{#palpable-content:the-bdo-element}](text-level-semantics.html#the-bdo-element)
- [`blockquote`{#palpable-content:the-blockquote-element}](grouping-content.html#the-blockquote-element)
- [`button`{#palpable-content:the-button-element}](form-elements.html#the-button-element)
- [`canvas`{#palpable-content:the-canvas-element-2}](canvas.html#the-canvas-element)
- [`cite`{#palpable-content:the-cite-element}](text-level-semantics.html#the-cite-element)
- [`code`{#palpable-content:the-code-element}](text-level-semantics.html#the-code-element)
- [`data`{#palpable-content:the-data-element}](text-level-semantics.html#the-data-element)
- [`del`{#palpable-content:the-del-element}](edits.html#the-del-element)
- [`details`{#palpable-content:the-details-element}](interactive-elements.html#the-details-element)
- [`dfn`{#palpable-content:the-dfn-element}](text-level-semantics.html#the-dfn-element)
- [`div`{#palpable-content:the-div-element}](grouping-content.html#the-div-element)
- [`dl`{#palpable-content:the-dl-element}](grouping-content.html#the-dl-element)
  (if the element\'s children include at least one name-value group)
- [`em`{#palpable-content:the-em-element}](text-level-semantics.html#the-em-element)
- [`embed`{#palpable-content:the-embed-element}](iframe-embed-object.html#the-embed-element)
- [`fieldset`{#palpable-content:the-fieldset-element}](form-elements.html#the-fieldset-element)
- [`figure`{#palpable-content:the-figure-element}](grouping-content.html#the-figure-element)
- [`footer`{#palpable-content:the-footer-element}](sections.html#the-footer-element)
- [`form`{#palpable-content:the-form-element}](forms.html#the-form-element)
- [`h1`{#palpable-content:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements)
- [`h2`{#palpable-content:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-2}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements)
- [`h3`{#palpable-content:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-3}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements)
- [`h4`{#palpable-content:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-4}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements)
- [`h5`{#palpable-content:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-5}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements)
- [`h6`{#palpable-content:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-6}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements)
- [`header`{#palpable-content:the-header-element}](sections.html#the-header-element)
- [`hgroup`{#palpable-content:the-hgroup-element}](sections.html#the-hgroup-element)
- [`i`{#palpable-content:the-i-element}](text-level-semantics.html#the-i-element)
- [`iframe`{#palpable-content:the-iframe-element}](iframe-embed-object.html#the-iframe-element)
- [`img`{#palpable-content:the-img-element-2}](embedded-content.html#the-img-element)
- [`input`{#palpable-content:the-input-element}](input.html#the-input-element)
  (if the
  [`type`{#palpable-content:attr-input-type}](input.html#attr-input-type)
  attribute is *not* in the
  [Hidden](input.html#hidden-state-(type=hidden)){#palpable-content:hidden-state-(type=hidden)}
  state)
- [`ins`{#palpable-content:the-ins-element}](edits.html#the-ins-element)
- [`kbd`{#palpable-content:the-kbd-element}](text-level-semantics.html#the-kbd-element)
- [`label`{#palpable-content:the-label-element}](forms.html#the-label-element)
- [`main`{#palpable-content:the-main-element}](grouping-content.html#the-main-element)
- [`map`{#palpable-content:the-map-element}](image-maps.html#the-map-element)
- [`mark`{#palpable-content:the-mark-element}](text-level-semantics.html#the-mark-element)
- [MathML
  `math`](https://w3c.github.io/mathml-core/#the-top-level-math-element){#palpable-content:mathml-math
  x-internal="mathml-math"}
- [`menu`{#palpable-content:the-menu-element}](grouping-content.html#the-menu-element)
  (if the element\'s children include at least one
  [`li`{#palpable-content:the-li-element}](grouping-content.html#the-li-element)
  element)
- [`meter`{#palpable-content:the-meter-element}](form-elements.html#the-meter-element)
- [`nav`{#palpable-content:the-nav-element}](sections.html#the-nav-element)
- [`object`{#palpable-content:the-object-element}](iframe-embed-object.html#the-object-element)
- [`ol`{#palpable-content:the-ol-element}](grouping-content.html#the-ol-element)
  (if the element\'s children include at least one
  [`li`{#palpable-content:the-li-element-2}](grouping-content.html#the-li-element)
  element)
- [`output`{#palpable-content:the-output-element}](form-elements.html#the-output-element)
- [`p`{#palpable-content:the-p-element}](grouping-content.html#the-p-element)
- [`picture`{#palpable-content:the-picture-element}](embedded-content.html#the-picture-element)
- [`pre`{#palpable-content:the-pre-element}](grouping-content.html#the-pre-element)
- [`progress`{#palpable-content:the-progress-element}](form-elements.html#the-progress-element)
- [`q`{#palpable-content:the-q-element}](text-level-semantics.html#the-q-element)
- [`ruby`{#palpable-content:the-ruby-element}](text-level-semantics.html#the-ruby-element)
- [`s`{#palpable-content:the-s-element}](text-level-semantics.html#the-s-element)
- [`samp`{#palpable-content:the-samp-element}](text-level-semantics.html#the-samp-element)
- [`search`{#palpable-content:the-search-element}](grouping-content.html#the-search-element)
- [`section`{#palpable-content:the-section-element}](sections.html#the-section-element)
- [`select`{#palpable-content:the-select-element}](form-elements.html#the-select-element)
- [`small`{#palpable-content:the-small-element}](text-level-semantics.html#the-small-element)
- [`span`{#palpable-content:the-span-element}](text-level-semantics.html#the-span-element)
- [`strong`{#palpable-content:the-strong-element}](text-level-semantics.html#the-strong-element)
- [`sub`{#palpable-content:the-sub-and-sup-elements}](text-level-semantics.html#the-sub-and-sup-elements)
- [`sup`{#palpable-content:the-sub-and-sup-elements-2}](text-level-semantics.html#the-sub-and-sup-elements)
- [SVG
  `svg`](https://svgwg.org/svg2-draft/struct.html#SVGElement){#palpable-content:svg-svg
  x-internal="svg-svg"}
- [`table`{#palpable-content:the-table-element}](tables.html#the-table-element)
- [`textarea`{#palpable-content:the-textarea-element}](form-elements.html#the-textarea-element)
- [`time`{#palpable-content:the-time-element}](text-level-semantics.html#the-time-element)
- [`u`{#palpable-content:the-u-element}](text-level-semantics.html#the-u-element)
- [`ul`{#palpable-content:the-ul-element}](grouping-content.html#the-ul-element)
  (if the element\'s children include at least one
  [`li`{#palpable-content:the-li-element-3}](grouping-content.html#the-li-element)
  element)
- [`var`{#palpable-content:the-var-element}](text-level-semantics.html#the-var-element)
- [`video`{#palpable-content:the-video-element-2}](media.html#the-video-element)
- [autonomous custom
  elements](custom-elements.html#autonomous-custom-element){#palpable-content:autonomous-custom-element}
- [text](#text-content){#palpable-content:text-content-2} that is not
  [inter-element
  whitespace](#inter-element-whitespace){#palpable-content:inter-element-whitespace}

###### [3.2.5.2.9]{.secno} Script-supporting elements[](#script-supporting-elements){.self-link}

[Script-supporting elements]{#script-supporting-elements-2 .dfn} are
those that do not
[represent](#represents){#script-supporting-elements:represents}
anything themselves (i.e. they are not rendered), but are used to
support scripts, e.g. to provide functionality for the user.

The following elements are script-supporting elements:

- [`script`{#script-supporting-elements:the-script-element}](scripting.html#the-script-element)
- [`template`{#script-supporting-elements:the-template-element}](scripting.html#the-template-element)

##### [3.2.5.3]{.secno} Transparent content models[](#transparent-content-models){.self-link}

Some elements are described as [transparent]{#transparent .dfn}; they
have \"transparent\" in the description of their content model. The
content model of a
[transparent](#transparent){#transparent-content-models:transparent}
element is derived from the content model of its parent element: the
elements required in the part of the content model that is
\"transparent\" are the same elements as required in the part of the
content model of the parent of the transparent element in which the
transparent element finds itself.

::: example
For instance, an
[`ins`{#transparent-content-models:the-ins-element}](edits.html#the-ins-element)
element inside a
[`ruby`{#transparent-content-models:the-ruby-element}](text-level-semantics.html#the-ruby-element)
element cannot contain an
[`rt`{#transparent-content-models:the-rt-element}](text-level-semantics.html#the-rt-element)
element, because the part of the
[`ruby`{#transparent-content-models:the-ruby-element-2}](text-level-semantics.html#the-ruby-element)
element\'s content model that allows
[`ins`{#transparent-content-models:the-ins-element-2}](edits.html#the-ins-element)
elements is the part that allows [phrasing
content](#phrasing-content-2){#transparent-content-models:phrasing-content-2},
and the
[`rt`{#transparent-content-models:the-rt-element-2}](text-level-semantics.html#the-rt-element)
element is not [phrasing
content](#phrasing-content-2){#transparent-content-models:phrasing-content-2-2}.
:::

In some cases, where transparent elements are nested in each other, the
process has to be applied iteratively.

::: example
Consider the following markup fragment:

``` html
<p><object><ins><map><a href="/">Apples</a></map></ins></object></p>
```

To check whether \"Apples\" is allowed inside the
[`a`{#transparent-content-models:the-a-element}](text-level-semantics.html#the-a-element)
element, the content models are examined. The
[`a`{#transparent-content-models:the-a-element-2}](text-level-semantics.html#the-a-element)
element\'s content model is transparent, as is the
[`map`{#transparent-content-models:the-map-element}](image-maps.html#the-map-element)
element\'s, as is the
[`ins`{#transparent-content-models:the-ins-element-3}](edits.html#the-ins-element)
element\'s, as is the
[`object`{#transparent-content-models:the-object-element}](iframe-embed-object.html#the-object-element)
element\'s. The
[`object`{#transparent-content-models:the-object-element-2}](iframe-embed-object.html#the-object-element)
element is found in the
[`p`{#transparent-content-models:the-p-element}](grouping-content.html#the-p-element)
element, whose content model is [phrasing
content](#phrasing-content-2){#transparent-content-models:phrasing-content-2-3}.
Thus, \"Apples\" is allowed, as text is phrasing content.
:::

When a transparent element has no parent, then the part of its content
model that is \"transparent\" must instead be treated as accepting any
[flow
content](#flow-content-2){#transparent-content-models:flow-content-2}.

##### [3.2.5.4]{.secno} Paragraphs[](#paragraphs){.self-link}

The term [paragraph](#paragraph){#paragraphs:paragraph} as defined in
this section is used for more than just the definition of the
[`p`{#paragraphs:the-p-element}](grouping-content.html#the-p-element)
element. The [paragraph](#paragraph){#paragraphs:paragraph-2} concept
defined here is used to describe how to interpret documents. The
[`p`{#paragraphs:the-p-element-2}](grouping-content.html#the-p-element)
element is merely one of several ways of marking up a
[paragraph](#paragraph){#paragraphs:paragraph-3}.

A [paragraph]{#paragraph .dfn} is typically a run of [phrasing
content](#phrasing-content-2){#paragraphs:phrasing-content-2} that forms
a block of text with one or more sentences that discuss a particular
topic, as in typography, but can also be used for more general thematic
grouping. For instance, an address is also a paragraph, as is a part of
a form, a byline, or a stanza in a poem.

::: example
In the following example, there are two paragraphs in a section. There
is also a heading, which contains phrasing content that is not a
paragraph. Note how the comments and [inter-element
whitespace](#inter-element-whitespace){#paragraphs:inter-element-whitespace}
do not form paragraphs.

``` html
<section>
  <h2>Example of paragraphs</h2>
  This is the <em>first</em> paragraph in this example.
  <p>This is the second.</p>
  <!-- This is not a paragraph. -->
</section>
```
:::

Paragraphs in [flow
content](#flow-content-2){#paragraphs:flow-content-2} are defined
relative to what the document looks like without the
[`a`{#paragraphs:the-a-element}](text-level-semantics.html#the-a-element),
[`ins`{#paragraphs:the-ins-element}](edits.html#the-ins-element),
[`del`{#paragraphs:the-del-element}](edits.html#the-del-element), and
[`map`{#paragraphs:the-map-element}](image-maps.html#the-map-element)
elements complicating matters, since those elements, with their hybrid
content models, can straddle paragraph boundaries, as shown in the first
two examples below.

Generally, having elements straddle paragraph boundaries is best
avoided. Maintaining such markup can be difficult.

::: example
The following example takes the markup from the earlier example and puts
[`ins`{#paragraphs:the-ins-element-2}](edits.html#the-ins-element) and
[`del`{#paragraphs:the-del-element-2}](edits.html#the-del-element)
elements around some of the markup to show that the text was changed
(though in this case, the changes admittedly don\'t make much sense).
Notice how this example has exactly the same paragraphs as the previous
one, despite the
[`ins`{#paragraphs:the-ins-element-3}](edits.html#the-ins-element) and
[`del`{#paragraphs:the-del-element-3}](edits.html#the-del-element)
elements --- the
[`ins`{#paragraphs:the-ins-element-4}](edits.html#the-ins-element)
element straddles the heading and the first paragraph, and the
[`del`{#paragraphs:the-del-element-4}](edits.html#the-del-element)
element straddles the boundary between the two paragraphs.

``` html
<section>
  <ins><h2>Example of paragraphs</h2>
  This is the <em>first</em> paragraph in</ins> this example<del>.
  <p>This is the second.</p></del>
  <!-- This is not a paragraph. -->
</section>
```
:::

Let `view`{.variable} be a view of the DOM that replaces all
[`a`{#paragraphs:the-a-element-2}](text-level-semantics.html#the-a-element),
[`ins`{#paragraphs:the-ins-element-5}](edits.html#the-ins-element),
[`del`{#paragraphs:the-del-element-5}](edits.html#the-del-element), and
[`map`{#paragraphs:the-map-element-2}](image-maps.html#the-map-element)
elements in the document with their
[contents](#concept-html-contents){#paragraphs:concept-html-contents}.
Then, in `view`{.variable}, for each run of sibling [phrasing
content](#phrasing-content-2){#paragraphs:phrasing-content-2-2} nodes
uninterrupted by other types of content, in an element that accepts
content other than [phrasing
content](#phrasing-content-2){#paragraphs:phrasing-content-2-3} as well
as [phrasing
content](#phrasing-content-2){#paragraphs:phrasing-content-2-4}, let
`first`{.variable} be the first node of the run, and let
`last`{.variable} be the last node of the run. For each such run that
consists of at least one node that is neither [embedded
content](#embedded-content-category){#paragraphs:embedded-content-category}
nor [inter-element
whitespace](#inter-element-whitespace){#paragraphs:inter-element-whitespace-2},
a paragraph exists in the original DOM from immediately before
`first`{.variable} to immediately after `last`{.variable}. (Paragraphs
can thus span across
[`a`{#paragraphs:the-a-element-3}](text-level-semantics.html#the-a-element),
[`ins`{#paragraphs:the-ins-element-6}](edits.html#the-ins-element),
[`del`{#paragraphs:the-del-element-6}](edits.html#the-del-element), and
[`map`{#paragraphs:the-map-element-3}](image-maps.html#the-map-element)
elements.)

Conformance checkers may warn authors of cases where they have
paragraphs that overlap each other (this can happen with
[`object`{#paragraphs:the-object-element}](iframe-embed-object.html#the-object-element),
[`video`{#paragraphs:the-video-element}](media.html#the-video-element),
[`audio`{#paragraphs:the-audio-element}](media.html#the-audio-element),
and
[`canvas`{#paragraphs:the-canvas-element}](canvas.html#the-canvas-element)
elements, and indirectly through elements in other namespaces that allow
HTML to be further embedded therein, like [SVG
`svg`](https://svgwg.org/svg2-draft/struct.html#SVGElement){#paragraphs:svg-svg
x-internal="svg-svg"} or [MathML
`math`](https://w3c.github.io/mathml-core/#the-top-level-math-element){#paragraphs:mathml-math
x-internal="mathml-math"}).

A [paragraph](#paragraph){#paragraphs:paragraph-4} is also formed
explicitly by
[`p`{#paragraphs:the-p-element-3}](grouping-content.html#the-p-element)
elements.

The
[`p`{#paragraphs:the-p-element-4}](grouping-content.html#the-p-element)
element can be used to wrap individual paragraphs when there would
otherwise not be any content other than phrasing content to separate the
paragraphs from each other.

::: example
In the following example, the link spans half of the first paragraph,
all of the heading separating the two paragraphs, and half of the second
paragraph. It straddles the paragraphs and the heading.

``` html
<header>
 Welcome!
 <a href="about.html">
  This is home of...
  <h1>The Falcons!</h1>
  The Lockheed Martin multirole jet fighter aircraft!
 </a>
 This page discusses the F-16 Fighting Falcon's innermost secrets.
</header>
```

Here is another way of marking this up, this time showing the paragraphs
explicitly, and splitting the one link element into three:

``` html
<header>
 <p>Welcome! <a href="about.html">This is home of...</a></p>
 <h1><a href="about.html">The Falcons!</a></h1>
 <p><a href="about.html">The Lockheed Martin multirole jet
 fighter aircraft!</a> This page discusses the F-16 Fighting
 Falcon's innermost secrets.</p>
</header>
```
:::

::: example
It is possible for paragraphs to overlap when using certain elements
that define fallback content. For example, in the following section:

``` html
<section>
 <h2>My Cats</h2>
 You can play with my cat simulator.
 <object data="cats.sim">
  To see the cat simulator, use one of the following links:
  <ul>
   <li><a href="cats.sim">Download simulator file</a>
   <li><a href="https://sims.example.com/watch?v=LYds5xY4INU">Use online simulator</a>
  </ul>
  Alternatively, upgrade to the Mellblom Browser.
 </object>
 I'm quite proud of it.
</section>
```

There are five paragraphs:

1.  The paragraph that says \"You can play with my cat simulator.
    *object* I\'m quite proud of it.\", where *object* is the
    [`object`{#paragraphs:the-object-element-2}](iframe-embed-object.html#the-object-element)
    element.
2.  The paragraph that says \"To see the cat simulator, use one of the
    following links:\".
3.  The paragraph that says \"Download simulator file\".
4.  The paragraph that says \"Use online simulator\".
5.  The paragraph that says \"Alternatively, upgrade to the Mellblom
    Browser.\".

The first paragraph is overlapped by the other four. A user agent that
supports the \"cats.sim\" resource will only show the first one, but a
user agent that shows the fallback will confusingly show the first
sentence of the first paragraph as if it was in the same paragraph as
the second one, and will show the last paragraph as if it was at the
start of the second sentence of the first paragraph.

To avoid this confusion, explicit
[`p`{#paragraphs:the-p-element-5}](grouping-content.html#the-p-element)
elements can be used. For example:

``` html
<section>
 <h2>My Cats</h2>
 <p>You can play with my cat simulator.</p>
 <object data="cats.sim">
  <p>To see the cat simulator, use one of the following links:</p>
  <ul>
   <li><a href="cats.sim">Download simulator file</a>
   <li><a href="https://sims.example.com/watch?v=LYds5xY4INU">Use online simulator</a>
  </ul>
  <p>Alternatively, upgrade to the Mellblom Browser.</p>
 </object>
 <p>I'm quite proud of it.</p>
</section>
```
:::

#### [3.2.6]{.secno} [Global attributes]{.dfn}[](#global-attributes){.self-link}

:::: {.mdn-anno .wrapped}
MDN

::: feature
[Global_attributes](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes "Global attributes are attributes common to all HTML elements; they can be used on all elements, though they may have no effect on some elements.")
:::
::::

The following attributes are common to and may be specified on all [HTML
elements](infrastructure.html#html-elements){#global-attributes:html-elements}
(even those not defined in this specification):

- [`accesskey`{#global-attributes:the-accesskey-attribute}](interaction.html#the-accesskey-attribute)
- [`autocapitalize`{#global-attributes:attr-autocapitalize}](interaction.html#attr-autocapitalize)
- [`autocorrect`{#global-attributes:attr-autocorrect}](interaction.html#attr-autocorrect)
- [`autofocus`{#global-attributes:attr-fe-autofocus}](interaction.html#attr-fe-autofocus)
- [`contenteditable`{#global-attributes:attr-contenteditable}](interaction.html#attr-contenteditable)
- [`dir`{#global-attributes:attr-dir}](#attr-dir)
- [`draggable`{#global-attributes:attr-draggable}](dnd.html#attr-draggable)
- [`enterkeyhint`{#global-attributes:attr-enterkeyhint}](interaction.html#attr-enterkeyhint)
- [`hidden`{#global-attributes:attr-hidden}](interaction.html#attr-hidden)
- [`inert`{#global-attributes:the-inert-attribute}](interaction.html#the-inert-attribute)
- [`inputmode`{#global-attributes:attr-inputmode}](interaction.html#attr-inputmode)
- [`is`{#global-attributes:attr-is}](custom-elements.html#attr-is)
- [`itemid`{#global-attributes:attr-itemid}](microdata.html#attr-itemid)
- [`itemprop`{#global-attributes:names:-the-itemprop-attribute}](microdata.html#names:-the-itemprop-attribute)
- [`itemref`{#global-attributes:attr-itemref}](microdata.html#attr-itemref)
- [`itemscope`{#global-attributes:attr-itemscope}](microdata.html#attr-itemscope)
- [`itemtype`{#global-attributes:attr-itemtype}](microdata.html#attr-itemtype)
- [`lang`{#global-attributes:attr-lang}](#attr-lang)
- [`nonce`{#global-attributes:attr-nonce}](urls-and-fetching.html#attr-nonce)
- [`popover`{#global-attributes:attr-popover}](popover.html#attr-popover)
- [`spellcheck`{#global-attributes:attr-spellcheck}](interaction.html#attr-spellcheck)
- [`style`{#global-attributes:attr-style}](#attr-style)
- [`tabindex`{#global-attributes:attr-tabindex}](interaction.html#attr-tabindex)
- [`title`{#global-attributes:attr-title}](#attr-title)
- [`translate`{#global-attributes:attr-translate}](#attr-translate)
- [`writingsuggestions`{#global-attributes:attr-writingsuggestions}](interaction.html#attr-writingsuggestions)

These attributes are only defined by this specification as attributes
for [HTML
elements](infrastructure.html#html-elements){#global-attributes:html-elements-2}.
When this specification refers to elements having these attributes,
elements from namespaces that are not defined as having these attributes
must not be considered as being elements with these attributes.

::: example
For example, in the following XML fragment, the \"`bogus`\" element does
not have a [`dir`{#global-attributes:attr-dir-2}](#attr-dir) attribute
as defined in this specification, despite having an attribute with the
literal name \"`dir`\". Thus, [the
directionality](#the-directionality){#global-attributes:the-directionality}
of the inner-most
[`span`{#global-attributes:the-span-element}](text-level-semantics.html#the-span-element)
element is \'[rtl](#concept-rtl){#global-attributes:concept-rtl}\',
inherited from the
[`div`{#global-attributes:the-div-element}](grouping-content.html#the-div-element)
element indirectly through the \"`bogus`\" element.

``` bad
<div xmlns="http://www.w3.org/1999/xhtml" dir="rtl">
 <bogus xmlns="https://example.net/ns" dir="ltr">
  <span xmlns="http://www.w3.org/1999/xhtml">
  </span>
 </bogus>
</div>
```
:::

------------------------------------------------------------------------

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Global_attributes/slot](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/slot "The slot global attribute assigns a slot in a shadow DOM shadow tree to an element: An element with a slot attribute is assigned to the slot created by the <slot> element whose name attribute's value matches that slot attribute's value.")

Support in all current engines.

::: support
[Firefox63+]{.firefox .yes}[Safari10+]{.safari .yes}[Chrome53+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)No]{.edge .no}[Internet Explorer?]{.ie .unknown}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

DOM defines the user agent requirements for the [`class`]{#classes .dfn
dfn-for="global" dfn-type="element-attr"}, [`id`]{#the-id-attribute .dfn
dfn-for="global" dfn-type="element-attr"}, and [`slot`]{#attr-slot .dfn
dfn-for="global" dfn-type="element-attr"} attributes for any element in
any namespace. [\[DOM\]](references.html#refsDOM)

The [`class`{#global-attributes:classes}](#classes),
[`id`{#global-attributes:the-id-attribute}](#the-id-attribute), and
[`slot`{#global-attributes:attr-slot}](#attr-slot) attributes may be
specified on all [HTML
elements](infrastructure.html#html-elements){#global-attributes:html-elements-3}.

When specified on [HTML
elements](infrastructure.html#html-elements){#global-attributes:html-elements-4},
the [`class`{#global-attributes:classes-2}](#classes) attribute must
have a value that is a [set of space-separated
tokens](common-microsyntaxes.html#set-of-space-separated-tokens){#global-attributes:set-of-space-separated-tokens}
representing the various classes that the element belongs to.

::: note
Assigning classes to an element affects class matching in selectors in
CSS, the
[`getElementsByClassName()`{#global-attributes:dom-document-getelementsbyclassname}](https://dom.spec.whatwg.org/#dom-document-getelementsbyclassname){x-internal="dom-document-getelementsbyclassname"}
method in the DOM, and other such features.

There are no additional restrictions on the tokens authors can use in
the [`class`{#global-attributes:classes-3}](#classes) attribute, but
authors are encouraged to use values that describe the nature of the
content, rather than values that describe the desired presentation of
the content.
:::

When specified on [HTML
elements](infrastructure.html#html-elements){#global-attributes:html-elements-5},
the [`id`{#global-attributes:the-id-attribute-2}](#the-id-attribute)
attribute value must be unique amongst all the
[IDs](https://dom.spec.whatwg.org/#concept-id){#global-attributes:concept-id
x-internal="concept-id"} in the element\'s
[tree](https://dom.spec.whatwg.org/#concept-tree){#global-attributes:tree
x-internal="tree"} and must contain at least one character. The value
must not contain any [ASCII
whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#global-attributes:space-characters
x-internal="space-characters"}.

::: note
The [`id`{#global-attributes:the-id-attribute-3}](#the-id-attribute)
attribute specifies its element\'s [unique identifier
(ID)](https://dom.spec.whatwg.org/#concept-id){#global-attributes:concept-id-2
x-internal="concept-id"}.

There are no other restrictions on what form an ID can take; in
particular, IDs can consist of just digits, start with a digit, start
with an underscore, consist of just punctuation, etc.

An element\'s [unique
identifier](https://dom.spec.whatwg.org/#concept-id){#global-attributes:concept-id-3
x-internal="concept-id"} can be used for a variety of purposes, most
notably as a way to link to specific parts of a document using
[fragments](https://url.spec.whatwg.org/#concept-url-fragment){#global-attributes:concept-url-fragment
x-internal="concept-url-fragment"}, as a way to target an element when
scripting, and as a way to style a specific element from CSS.
:::

Identifiers are opaque strings. Particular meanings should not be
derived from the value of the
[`id`{#global-attributes:the-id-attribute-4}](#the-id-attribute)
attribute.

There are no conformance requirements for the
[`slot`{#global-attributes:attr-slot-2}](#attr-slot) attribute specific
to [HTML
elements](infrastructure.html#html-elements){#global-attributes:html-elements-6}.

The [`slot`{#global-attributes:attr-slot-3}](#attr-slot) attribute is
used to [assign a
slot](https://dom.spec.whatwg.org/#assign-a-slot){#global-attributes:assign-a-slot
x-internal="assign-a-slot"} to an element: an element with a
[`slot`{#global-attributes:attr-slot-4}](#attr-slot) attribute is
[assigned](https://dom.spec.whatwg.org/#assign-a-slot){#global-attributes:assign-a-slot-2
x-internal="assign-a-slot"} to the
[slot](https://dom.spec.whatwg.org/#concept-slot){#global-attributes:concept-slot
x-internal="concept-slot"} created by the
[`slot`{#global-attributes:the-slot-element}](scripting.html#the-slot-element)
element whose
[name](scripting.html#attr-slot-name){#global-attributes:attr-slot-name}
attribute\'s value matches that
[`slot`{#global-attributes:attr-slot-5}](#attr-slot) attribute\'s value
--- but only if that
[`slot`{#global-attributes:the-slot-element-2}](scripting.html#the-slot-element)
element finds itself in the [shadow
tree](https://dom.spec.whatwg.org/#concept-shadow-tree){#global-attributes:shadow-tree
x-internal="shadow-tree"} whose
[root](https://dom.spec.whatwg.org/#concept-tree-root){#global-attributes:root
x-internal="root"}\'s
[host](https://dom.spec.whatwg.org/#concept-documentfragment-host){#global-attributes:concept-documentfragment-host
x-internal="concept-documentfragment-host"} has the corresponding
[`slot`{#global-attributes:attr-slot-6}](#attr-slot) attribute value.

------------------------------------------------------------------------

To enable assistive technology products to expose a more fine-grained
interface than is otherwise possible with HTML elements and attributes,
a set of [annotations for assistive technology products](#wai-aria) can
be specified (the ARIA
[`role`{#global-attributes:attr-aria-role}](infrastructure.html#attr-aria-role)
and
[`aria-*`{#global-attributes:attr-aria-*}](infrastructure.html#attr-aria-*)
attributes). [\[ARIA\]](references.html#refsARIA)

------------------------------------------------------------------------

The following [event handler content
attributes](webappapis.html#event-handler-content-attributes){#global-attributes:event-handler-content-attributes}
may be specified on any [HTML
element](infrastructure.html#html-elements){#global-attributes:html-elements-7}:

- [`onauxclick`{#global-attributes:handler-onauxclick}](webappapis.html#handler-onauxclick)
- [`onbeforeinput`{#global-attributes:handler-onbeforeinput}](webappapis.html#handler-onbeforeinput)
- [`onbeforematch`{#global-attributes:handler-onbeforematch}](webappapis.html#handler-onbeforematch)
- [`onbeforetoggle`{#global-attributes:handler-onbeforetoggle}](webappapis.html#handler-onbeforetoggle)
- [`onblur`{#global-attributes:handler-onblur}](webappapis.html#handler-onblur)\*
- [`oncancel`{#global-attributes:handler-oncancel}](webappapis.html#handler-oncancel)
- [`oncanplay`{#global-attributes:handler-oncanplay}](webappapis.html#handler-oncanplay)
- [`oncanplaythrough`{#global-attributes:handler-oncanplaythrough}](webappapis.html#handler-oncanplaythrough)
- [`onchange`{#global-attributes:handler-onchange}](webappapis.html#handler-onchange)
- [`onclick`{#global-attributes:handler-onclick}](webappapis.html#handler-onclick)
- [`onclose`{#global-attributes:handler-onclose}](webappapis.html#handler-onclose)
- [`oncommand`{#global-attributes:handler-oncommand}](webappapis.html#handler-oncommand)
- [`oncontextlost`{#global-attributes:handler-oncontextlost}](webappapis.html#handler-oncontextlost)
- [`oncontextmenu`{#global-attributes:handler-oncontextmenu}](webappapis.html#handler-oncontextmenu)
- [`oncontextrestored`{#global-attributes:handler-oncontextrestored}](webappapis.html#handler-oncontextrestored)
- [`oncopy`{#global-attributes:handler-oncopy}](webappapis.html#handler-oncopy)
- [`oncuechange`{#global-attributes:handler-oncuechange}](webappapis.html#handler-oncuechange)
- [`oncut`{#global-attributes:handler-oncut}](webappapis.html#handler-oncut)
- [`ondblclick`{#global-attributes:handler-ondblclick}](webappapis.html#handler-ondblclick)
- [`ondrag`{#global-attributes:handler-ondrag}](webappapis.html#handler-ondrag)
- [`ondragend`{#global-attributes:handler-ondragend}](webappapis.html#handler-ondragend)
- [`ondragenter`{#global-attributes:handler-ondragenter}](webappapis.html#handler-ondragenter)
- [`ondragleave`{#global-attributes:handler-ondragleave}](webappapis.html#handler-ondragleave)
- [`ondragover`{#global-attributes:handler-ondragover}](webappapis.html#handler-ondragover)
- [`ondragstart`{#global-attributes:handler-ondragstart}](webappapis.html#handler-ondragstart)
- [`ondrop`{#global-attributes:handler-ondrop}](webappapis.html#handler-ondrop)
- [`ondurationchange`{#global-attributes:handler-ondurationchange}](webappapis.html#handler-ondurationchange)
- [`onemptied`{#global-attributes:handler-onemptied}](webappapis.html#handler-onemptied)
- [`onended`{#global-attributes:handler-onended}](webappapis.html#handler-onended)
- [`onerror`{#global-attributes:handler-onerror}](webappapis.html#handler-onerror)\*
- [`onfocus`{#global-attributes:handler-onfocus}](webappapis.html#handler-onfocus)\*
- [`onformdata`{#global-attributes:handler-onformdata}](webappapis.html#handler-onformdata)
- [`oninput`{#global-attributes:handler-oninput}](webappapis.html#handler-oninput)
- [`oninvalid`{#global-attributes:handler-oninvalid}](webappapis.html#handler-oninvalid)
- [`onkeydown`{#global-attributes:handler-onkeydown}](webappapis.html#handler-onkeydown)
- [`onkeypress`{#global-attributes:handler-onkeypress}](webappapis.html#handler-onkeypress)
- [`onkeyup`{#global-attributes:handler-onkeyup}](webappapis.html#handler-onkeyup)
- [`onload`{#global-attributes:handler-onload}](webappapis.html#handler-onload)\*
- [`onloadeddata`{#global-attributes:handler-onloadeddata}](webappapis.html#handler-onloadeddata)
- [`onloadedmetadata`{#global-attributes:handler-onloadedmetadata}](webappapis.html#handler-onloadedmetadata)
- [`onloadstart`{#global-attributes:handler-onloadstart}](webappapis.html#handler-onloadstart)
- [`onmousedown`{#global-attributes:handler-onmousedown}](webappapis.html#handler-onmousedown)
- [`onmouseenter`{#global-attributes:handler-onmouseenter}](webappapis.html#handler-onmouseenter)
- [`onmouseleave`{#global-attributes:handler-onmouseleave}](webappapis.html#handler-onmouseleave)
- [`onmousemove`{#global-attributes:handler-onmousemove}](webappapis.html#handler-onmousemove)
- [`onmouseout`{#global-attributes:handler-onmouseout}](webappapis.html#handler-onmouseout)
- [`onmouseover`{#global-attributes:handler-onmouseover}](webappapis.html#handler-onmouseover)
- [`onmouseup`{#global-attributes:handler-onmouseup}](webappapis.html#handler-onmouseup)
- [`onpaste`{#global-attributes:handler-onpaste}](webappapis.html#handler-onpaste)
- [`onpause`{#global-attributes:handler-onpause}](webappapis.html#handler-onpause)
- [`onplay`{#global-attributes:handler-onplay}](webappapis.html#handler-onplay)
- [`onplaying`{#global-attributes:handler-onplaying}](webappapis.html#handler-onplaying)
- [`onprogress`{#global-attributes:handler-onprogress}](webappapis.html#handler-onprogress)
- [`onratechange`{#global-attributes:handler-onratechange}](webappapis.html#handler-onratechange)
- [`onreset`{#global-attributes:handler-onreset}](webappapis.html#handler-onreset)
- [`onresize`{#global-attributes:handler-onresize}](webappapis.html#handler-onresize)\*
- [`onscroll`{#global-attributes:handler-onscroll}](webappapis.html#handler-onscroll)\*
- [`onscrollend`{#global-attributes:handler-onscrollend}](webappapis.html#handler-onscrollend)\*
- [`onsecuritypolicyviolation`{#global-attributes:handler-onsecuritypolicyviolation}](webappapis.html#handler-onsecuritypolicyviolation)
- [`onseeked`{#global-attributes:handler-onseeked}](webappapis.html#handler-onseeked)
- [`onseeking`{#global-attributes:handler-onseeking}](webappapis.html#handler-onseeking)
- [`onselect`{#global-attributes:handler-onselect}](webappapis.html#handler-onselect)
- [`onslotchange`{#global-attributes:handler-onslotchange}](webappapis.html#handler-onslotchange)
- [`onstalled`{#global-attributes:handler-onstalled}](webappapis.html#handler-onstalled)
- [`onsubmit`{#global-attributes:handler-onsubmit}](webappapis.html#handler-onsubmit)
- [`onsuspend`{#global-attributes:handler-onsuspend}](webappapis.html#handler-onsuspend)
- [`ontimeupdate`{#global-attributes:handler-ontimeupdate}](webappapis.html#handler-ontimeupdate)
- [`ontoggle`{#global-attributes:handler-ontoggle}](webappapis.html#handler-ontoggle)
- [`onvolumechange`{#global-attributes:handler-onvolumechange}](webappapis.html#handler-onvolumechange)
- [`onwaiting`{#global-attributes:handler-onwaiting}](webappapis.html#handler-onwaiting)
- [`onwheel`{#global-attributes:handler-onwheel}](webappapis.html#handler-onwheel)

The attributes marked with an asterisk have a different meaning when
specified on
[`body`{#global-attributes:the-body-element}](sections.html#the-body-element)
elements as those elements expose [event
handlers](webappapis.html#event-handlers){#global-attributes:event-handlers}
of the
[`Window`{#global-attributes:window}](nav-history-apis.html#window)
object with the same names.

While these attributes apply to all elements, they are not useful on all
elements. For example, only [media
elements](media.html#media-element){#global-attributes:media-element}
will ever receive a
[`volumechange`{#global-attributes:event-media-volumechange}](media.html#event-media-volumechange)
event fired by the user agent.

------------------------------------------------------------------------

[Custom data
attributes](#custom-data-attribute){#global-attributes:custom-data-attribute}
(e.g. `data-foldername` or `data-msgid`) can be specified on any [HTML
element](infrastructure.html#html-elements){#global-attributes:html-elements-8},
to store custom data, state, annotations, and similar, specific to the
page.

------------------------------------------------------------------------

In [HTML
documents](https://dom.spec.whatwg.org/#html-document){#global-attributes:html-documents
x-internal="html-documents"}, elements in the [HTML
namespace](https://infra.spec.whatwg.org/#html-namespace){#global-attributes:html-namespace-2
x-internal="html-namespace-2"} may have an `xmlns` attribute specified,
if, and only if, it has the exact value
\"`http://www.w3.org/1999/xhtml`\". This does not apply to [XML
documents](https://dom.spec.whatwg.org/#xml-document){#global-attributes:xml-documents
x-internal="xml-documents"}.

In HTML, the `xmlns` attribute has absolutely no effect. It is basically
a talisman. It is allowed merely to make migration to and from XML
mildly easier. When parsed by an [HTML
parser](parsing.html#html-parser){#global-attributes:html-parser}, the
attribute ends up in no namespace, not the
\"`http://www.w3.org/2000/xmlns/`\" namespace like namespace declaration
attributes in XML do.

In XML, an `xmlns` attribute is part of the namespace declaration
mechanism, and an element cannot actually have an `xmlns` attribute in
no namespace specified.

------------------------------------------------------------------------

XML also allows the use of the
[`xml:space`{#global-attributes:attr-xml-space}](https://www.w3.org/TR/xml/#sec-white-space){x-internal="attr-xml-space"}
attribute in the [XML
namespace](https://infra.spec.whatwg.org/#xml-namespace){#global-attributes:xml-namespace
x-internal="xml-namespace"} on any element in an [XML
document](https://dom.spec.whatwg.org/#xml-document){#global-attributes:xml-documents-2
x-internal="xml-documents"}. This attribute has no effect on [HTML
elements](infrastructure.html#html-elements){#global-attributes:html-elements-9},
as the default behavior in HTML is to preserve whitespace.
[\[XML\]](references.html#refsXML)

There is no way to serialize the
[`xml:space`{#global-attributes:attr-xml-space-2}](https://www.w3.org/TR/xml/#sec-white-space){x-internal="attr-xml-space"}
attribute on [HTML
elements](infrastructure.html#html-elements){#global-attributes:html-elements-10}
in the [`text/html`{#global-attributes:text/html}](iana.html#text/html)
syntax.

##### [3.2.6.1]{.secno} The [`title`{#the-title-attribute:attr-title}](#attr-title) attribute[](#the-title-attribute){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Global_attributes/title](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/title "The title global attribute contains text representing advisory information related to the element it belongs to.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerYes]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

The [`title`]{#attr-title .dfn dfn-for="html-global"
dfn-type="element-attr"} attribute
[represents](#represents){#the-title-attribute:represents} advisory
information for the element, such as would be appropriate for a tooltip.
On a link, this could be the title or a description of the target
resource; on an image, it could be the image credit or a description of
the image; on a paragraph, it could be a footnote or commentary on the
text; on a citation, it could be further information about the source;
on [interactive
content](#interactive-content-2){#the-title-attribute:interactive-content-2},
it could be a label for, or instructions for, use of the element; and so
forth. The value is text.

Relying on the [`title`{#the-title-attribute:attr-title-2}](#attr-title)
attribute is currently discouraged as many user agents do not expose the
attribute in an accessible manner as required by this specification
(e.g., requiring a pointing device such as a mouse to cause a tooltip to
appear, which excludes keyboard-only users and touch-only users, such as
anyone with a modern phone or tablet).

If this attribute is omitted from an element, then it implies that the
[`title`{#the-title-attribute:attr-title-3}](#attr-title) attribute of
the nearest ancestor [HTML
element](infrastructure.html#html-elements){#the-title-attribute:html-elements}
with a [`title`{#the-title-attribute:attr-title-4}](#attr-title)
attribute set is also relevant to this element. Setting the attribute
overrides this, explicitly stating that the advisory information of any
ancestors is not relevant to this element. Setting the attribute to the
empty string indicates that the element has no advisory information.

If the [`title`{#the-title-attribute:attr-title-5}](#attr-title)
attribute\'s value contains U+000A LINE FEED (LF) characters, the
content is split into multiple lines. Each U+000A LINE FEED (LF)
character represents a line break.

::: example
Caution is advised with respect to the use of newlines in
[`title`{#the-title-attribute:attr-title-6}](#attr-title) attributes.

For instance, the following snippet actually defines an abbreviation\'s
expansion *with a line break in it*:

``` bad
<p>My logs show that there was some interest in <abbr title="Hypertext
Transport Protocol">HTTP</abbr> today.</p>
```
:::

Some elements, such as
[`link`{#the-title-attribute:the-link-element}](semantics.html#the-link-element),
[`abbr`{#the-title-attribute:the-abbr-element}](text-level-semantics.html#the-abbr-element),
and
[`input`{#the-title-attribute:the-input-element}](input.html#the-input-element),
define additional semantics for the
[`title`{#the-title-attribute:attr-title-7}](#attr-title) attribute
beyond the semantics described above.

The [advisory information]{#advisory-information .dfn} of an element is
the value that the following algorithm returns, with the algorithm being
aborted once a value is returned. When the algorithm returns the empty
string, then there is no advisory information.

1.  If the element has a
    [`title`{#the-title-attribute:attr-title-8}](#attr-title) attribute,
    then return the result of running [normalize
    newlines](https://infra.spec.whatwg.org/#normalize-newlines){#the-title-attribute:normalize-newlines
    x-internal="normalize-newlines"} on its value.

2.  If the element has a parent element, then return the parent
    element\'s [advisory
    information](#advisory-information){#the-title-attribute:advisory-information}.

3.  Return the empty string.

User agents should inform the user when elements have [advisory
information](#advisory-information){#the-title-attribute:advisory-information-2},
otherwise the information would not be discoverable.

------------------------------------------------------------------------

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLElement/title](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/title "The HTMLElement.title property represents the title of the element: the text usually displayed in a 'tooltip' popup when the mouse is over the node.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The [`title`]{#dom-title .dfn dfn-for="HTMLElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-title-attribute:reflect}
the [`title`{#the-title-attribute:attr-title-9}](#attr-title) content
attribute.

##### [3.2.6.2]{.secno} The [`lang`{#the-lang-and-xml:lang-attributes:attr-lang}](#attr-lang) and [`xml:lang`{#the-lang-and-xml:lang-attributes:attr-xml-lang}](https://www.w3.org/TR/xml/#sec-lang-tag){x-internal="attr-xml-lang"} attributes[](#the-lang-and-xml:lang-attributes){.self-link} {#the-lang-and-xml:lang-attributes}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Global_attributes/lang](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/lang "The lang global attribute helps define the language of an element: the language that non-editable elements are written in, or the language that the editable elements should be written in by the user. The attribute contains a single "language tag" in the format defined in RFC 5646: Tags for Identifying Languages (also known as BCP 47).")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerYes]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

The [`lang`]{#attr-lang .dfn dfn-for="html-global"
dfn-type="element-attr"} attribute (in no namespace) specifies the
primary language for the element\'s contents and for any of the
element\'s attributes that contain text. Its value must be a valid BCP
47 language tag, or the empty string. Setting the attribute to the empty
string indicates that the primary language is unknown.
[\[BCP47\]](references.html#refsBCP47)

The [[`lang`](https://www.w3.org/TR/xml/#sec-lang-tag)]{#attr-xml-lang
.dfn} attribute in the [XML
namespace](https://infra.spec.whatwg.org/#xml-namespace){#the-lang-and-xml:lang-attributes:xml-namespace
x-internal="xml-namespace"} is defined in XML.
[\[XML\]](references.html#refsXML)

If these attributes are omitted from an element, then the language of
this element is the same as the language of its parent element, if any
(except for
[`slot`{#the-lang-and-xml:lang-attributes:the-slot-element}](scripting.html#the-slot-element)
elements in a [shadow
tree](https://dom.spec.whatwg.org/#concept-shadow-tree){#the-lang-and-xml:lang-attributes:shadow-tree
x-internal="shadow-tree"}).

The [`lang`{#the-lang-and-xml:lang-attributes:attr-lang-2}](#attr-lang)
attribute in no namespace may be used on any [HTML
element](infrastructure.html#html-elements){#the-lang-and-xml:lang-attributes:html-elements}.

The [`lang` attribute in the XML
namespace](https://www.w3.org/TR/xml/#sec-lang-tag){#the-lang-and-xml:lang-attributes:attr-xml-lang-2
x-internal="attr-xml-lang"} may be used on [HTML
elements](infrastructure.html#html-elements){#the-lang-and-xml:lang-attributes:html-elements-2}
in [XML
documents](https://dom.spec.whatwg.org/#xml-document){#the-lang-and-xml:lang-attributes:xml-documents
x-internal="xml-documents"}, as well as elements in other namespaces if
the relevant specifications allow it (in particular, MathML and SVG
allow [`lang` attributes in the XML
namespace](https://www.w3.org/TR/xml/#sec-lang-tag){#the-lang-and-xml:lang-attributes:attr-xml-lang-3
x-internal="attr-xml-lang"} to be specified on their elements). If both
the [`lang`{#the-lang-and-xml:lang-attributes:attr-lang-3}](#attr-lang)
attribute in no namespace and the [`lang` attribute in the XML
namespace](https://www.w3.org/TR/xml/#sec-lang-tag){#the-lang-and-xml:lang-attributes:attr-xml-lang-4
x-internal="attr-xml-lang"} are specified on the same element, they must
have exactly the same value when compared in an [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#the-lang-and-xml:lang-attributes:ascii-case-insensitive
x-internal="ascii-case-insensitive"} manner.

Authors must not use the [`lang` attribute in the XML
namespace](https://www.w3.org/TR/xml/#sec-lang-tag){#the-lang-and-xml:lang-attributes:attr-xml-lang-5
x-internal="attr-xml-lang"} on [HTML
elements](infrastructure.html#html-elements){#the-lang-and-xml:lang-attributes:html-elements-3}
in [HTML
documents](https://dom.spec.whatwg.org/#html-document){#the-lang-and-xml:lang-attributes:html-documents
x-internal="html-documents"}. To ease migration to and from XML, authors
may specify an attribute in no namespace with no prefix and with the
literal localname \"`xml:lang`\" on [HTML
elements](infrastructure.html#html-elements){#the-lang-and-xml:lang-attributes:html-elements-4}
in [HTML
documents](https://dom.spec.whatwg.org/#html-document){#the-lang-and-xml:lang-attributes:html-documents-2
x-internal="html-documents"}, but such attributes must only be specified
if a [`lang`{#the-lang-and-xml:lang-attributes:attr-lang-4}](#attr-lang)
attribute in no namespace is also specified, and both attributes must
have the same value when compared in an [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#the-lang-and-xml:lang-attributes:ascii-case-insensitive-2
x-internal="ascii-case-insensitive"} manner.

The attribute in no namespace with no prefix and with the literal
localname \"`xml:lang`\" has no effect on language processing.

------------------------------------------------------------------------

To determine the [language]{#language .dfn dfn-for="Node" export=""} of
a node, user agents must use the first appropriate step in the following
list:

If the node is an element that has a [`lang` attribute in the XML namespace](https://www.w3.org/TR/xml/#sec-lang-tag){#the-lang-and-xml:lang-attributes:attr-xml-lang-6 x-internal="attr-xml-lang"} set
:   Use the value of that attribute.

If the node is an [HTML element](infrastructure.html#html-elements){#the-lang-and-xml:lang-attributes:html-elements-5} or an element in the [SVG namespace](https://infra.spec.whatwg.org/#svg-namespace){#the-lang-and-xml:lang-attributes:svg-namespace x-internal="svg-namespace"}, and it has a [`lang`{#the-lang-and-xml:lang-attributes:attr-lang-5}](#attr-lang) in no namespace attribute set
:   Use the value of that attribute.

If the node\'s parent is a [shadow root](https://dom.spec.whatwg.org/#concept-shadow-root){#the-lang-and-xml:lang-attributes:shadow-root x-internal="shadow-root"}
:   Use the
    [language](#language){#the-lang-and-xml:lang-attributes:language} of
    that [shadow
    root](https://dom.spec.whatwg.org/#concept-shadow-root){#the-lang-and-xml:lang-attributes:shadow-root-2
    x-internal="shadow-root"}\'s
    [host](https://dom.spec.whatwg.org/#concept-documentfragment-host){#the-lang-and-xml:lang-attributes:concept-documentfragment-host
    x-internal="concept-documentfragment-host"}.

If the node\'s [parent element](https://dom.spec.whatwg.org/#parent-element){#the-lang-and-xml:lang-attributes:parent-element x-internal="parent-element"} is not null
:   Use the
    [language](#language){#the-lang-and-xml:lang-attributes:language-2}
    of that [parent
    element](https://dom.spec.whatwg.org/#parent-element){#the-lang-and-xml:lang-attributes:parent-element-2
    x-internal="parent-element"}.

Otherwise

:   If there is a [pragma-set default
    language](semantics.html#pragma-set-default-language){#the-lang-and-xml:lang-attributes:pragma-set-default-language}
    set, then that is the language of the node. If there is no
    [pragma-set default
    language](semantics.html#pragma-set-default-language){#the-lang-and-xml:lang-attributes:pragma-set-default-language-2}
    set, then language information from a higher-level protocol (such as
    HTTP), if any, must be used as the final fallback language instead.
    In the absence of any such language information, and in cases where
    the higher-level protocol reports multiple languages, the language
    of the node is unknown, and the corresponding language tag is the
    empty string.

If the resulting value is not a recognized language tag, then it must be
treated as an unknown language having the given language tag, distinct
from all other languages. For the purposes of round-tripping or
communicating with other services that expect language tags, user agents
should pass unknown language tags through unmodified, and tagged as
being BCP 47 language tags, so that subsequent services do not interpret
the data as another type of language description.
[\[BCP47\]](references.html#refsBCP47)

Thus, for instance, an element with `lang="xyzzy"` would be matched by
the selector `:lang(xyzzy)` (e.g. in CSS), but it would not be matched
by `:lang(abcde)`, even though both are equally invalid. Similarly, if a
web browser and screen reader working in unison communicated about the
language of the element, the browser would tell the screen reader that
the language was \"xyzzy\", even if it knew it was invalid, just in case
the screen reader actually supported a language with that tag after all.
Even if the screen reader supported both BCP 47 and another syntax for
encoding language names, and in that other syntax the string \"xyzzy\"
was a way to denote the Belarusian language, it would be *incorrect* for
the screen reader to then start treating text as Belarusian, because
\"xyzzy\" is not how Belarusian is described in BCP 47 codes (BCP 47
uses the code \"be\" for Belarusian).

If the resulting value is the empty string, then it must be interpreted
as meaning that the language of the node is [explicitly
unknown]{#concept-explicitly-unknown .dfn}.

------------------------------------------------------------------------

User agents may use the element\'s language to determine proper
processing or rendering (e.g. in the selection of appropriate fonts or
pronunciations, for dictionary selection, or for the user interfaces of
form controls such as date pickers).

------------------------------------------------------------------------

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLElement/lang](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/lang "The HTMLElement.lang property gets or sets the base language of an element's attribute values and text content.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The [`lang`]{#dom-lang .dfn dfn-for="HTMLElement" dfn-type="attribute"}
IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-lang-and-xml:lang-attributes:reflect}
the [`lang`{#the-lang-and-xml:lang-attributes:attr-lang-6}](#attr-lang)
content attribute in no namespace.

##### [3.2.6.3]{.secno} The [`translate`{#the-translate-attribute:attr-translate}](#attr-translate) attribute[](#the-translate-attribute){.self-link}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Global_attributes/translate](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/translate "The translate global attribute is an enumerated attribute that is used to specify whether an element's translatable attribute values and its Text node children should be translated when the page is localized, or whether to leave them unchanged.")

Support in all current engines.

::: support
[Firefox111+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome19+]{.chrome
.yes}

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

The [`translate`]{#attr-translate .dfn dfn-for="html-global"
dfn-type="element-attr"} attribute is used to specify whether an
element\'s attribute values and the values of its
[`Text`{#the-translate-attribute:text}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
node children are to be translated when the page is localized, or
whether to leave them unchanged. It is an attribute is an [enumerated
attribute](common-microsyntaxes.html#enumerated-attribute){#the-translate-attribute:enumerated-attribute}
with the following keywords and states:

Keyword

State

Brief description

[`yes`]{#attr-translate-yes-keyword .dfn dfn-for="html-global/translate"
dfn-type="attr-value"}

[Yes]{#attr-translate-yes-state .dfn}

Sets [translation
mode](#translation-mode){#the-translate-attribute:translation-mode} to
[translate-enabled](#translate-enabled){#the-translate-attribute:translate-enabled}.

(the empty string)

[`no`]{#attr-translate-no-keyword .dfn dfn-for="html-global/translate"
dfn-type="attr-value"}

[No]{#attr-translate-no-state .dfn}

Sets [translation
mode](#translation-mode){#the-translate-attribute:translation-mode-2} to
[no-translate](#no-translate){#the-translate-attribute:no-translate}.

The attribute\'s *[missing value
default](common-microsyntaxes.html#missing-value-default)* and *[invalid
value default](common-microsyntaxes.html#invalid-value-default)* are
both the [Inherit]{#attr-translate-inherit-state .dfn} state.

Each element (even non-HTML elements) has a [translation
mode]{#translation-mode .dfn}, which is in either the
[translate-enabled](#translate-enabled){#the-translate-attribute:translate-enabled-2}
state or the
[no-translate](#no-translate){#the-translate-attribute:no-translate-2}
state. If an [HTML
element](infrastructure.html#html-elements){#the-translate-attribute:html-elements}\'s
[`translate`{#the-translate-attribute:attr-translate-2}](#attr-translate)
attribute is in the
[Yes](#attr-translate-yes-state){#the-translate-attribute:attr-translate-yes-state}
state, then the element\'s [translation
mode](#translation-mode){#the-translate-attribute:translation-mode-3} is
in the
[translate-enabled](#translate-enabled){#the-translate-attribute:translate-enabled-3}
state; otherwise, if the element\'s
[`translate`{#the-translate-attribute:attr-translate-3}](#attr-translate)
attribute is in the
[No](#attr-translate-no-state){#the-translate-attribute:attr-translate-no-state}
state, then the element\'s [translation
mode](#translation-mode){#the-translate-attribute:translation-mode-4} is
in the
[no-translate](#no-translate){#the-translate-attribute:no-translate-3}
state. Otherwise, either the element\'s
[`translate`{#the-translate-attribute:attr-translate-4}](#attr-translate)
attribute is in the
[Inherit](#attr-translate-inherit-state){#the-translate-attribute:attr-translate-inherit-state}
state, or the element is not an [HTML
element](infrastructure.html#html-elements){#the-translate-attribute:html-elements-2}
and thus does not have a
[`translate`{#the-translate-attribute:attr-translate-5}](#attr-translate)
attribute; in either case, the element\'s [translation
mode](#translation-mode){#the-translate-attribute:translation-mode-5} is
in the same state as its [parent
element](https://dom.spec.whatwg.org/#parent-element){#the-translate-attribute:parent-element
x-internal="parent-element"}\'s, if any, or in the
[translate-enabled](#translate-enabled){#the-translate-attribute:translate-enabled-4}
state, if the element\'s [parent
element](https://dom.spec.whatwg.org/#parent-element){#the-translate-attribute:parent-element-2
x-internal="parent-element"} is null.

When an element is in the [translate-enabled]{#translate-enabled .dfn}
state, the element\'s [translatable
attributes](#translatable-attributes){#the-translate-attribute:translatable-attributes}
and the values of its
[`Text`{#the-translate-attribute:text-2}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
node children are to be translated when the page is localized.

When an element is in the [no-translate]{#no-translate .dfn} state, the
element\'s attribute values and the values of its
[`Text`{#the-translate-attribute:text-3}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
node children are to be left as-is when the page is localized, e.g.
because the element contains a person\'s name or a name of a computer
program.

The following attributes are [translatable
attributes]{#translatable-attributes .dfn}:

- [`abbr`{#the-translate-attribute:attr-th-abbr}](tables.html#attr-th-abbr)
  on
  [`th`{#the-translate-attribute:the-th-element}](tables.html#the-th-element)
  elements
- `alt` on
  [`area`{#the-translate-attribute:attr-area-alt}](image-maps.html#attr-area-alt),
  [`img`{#the-translate-attribute:attr-img-alt}](embedded-content.html#attr-img-alt),
  and
  [`input`{#the-translate-attribute:attr-input-alt}](input.html#attr-input-alt)
  elements
- [`content`{#the-translate-attribute:attr-meta-content}](semantics.html#attr-meta-content)
  on
  [`meta`{#the-translate-attribute:the-meta-element}](semantics.html#the-meta-element)
  elements, if the
  [`name`{#the-translate-attribute:attr-meta-name}](semantics.html#attr-meta-name)
  attribute specifies a metadata name whose value is known to be
  translatable
- [`download`{#the-translate-attribute:attr-hyperlink-download}](links.html#attr-hyperlink-download)
  on
  [`a`{#the-translate-attribute:the-a-element}](text-level-semantics.html#the-a-element)
  and
  [`area`{#the-translate-attribute:the-area-element}](image-maps.html#the-area-element)
  elements
- `label` on
  [`optgroup`{#the-translate-attribute:attr-optgroup-label}](form-elements.html#attr-optgroup-label),
  [`option`{#the-translate-attribute:attr-option-label}](form-elements.html#attr-option-label),
  and
  [`track`{#the-translate-attribute:attr-track-label}](media.html#attr-track-label)
  elements
- [`lang`{#the-translate-attribute:attr-lang}](#attr-lang) on [HTML
  elements](infrastructure.html#html-elements){#the-translate-attribute:html-elements-3};
  must be \"translated\" to match the language used in the translation
- `placeholder` on
  [`input`{#the-translate-attribute:attr-input-placeholder}](input.html#attr-input-placeholder)
  and
  [`textarea`{#the-translate-attribute:attr-textarea-placeholder}](form-elements.html#attr-textarea-placeholder)
  elements
- [`srcdoc`{#the-translate-attribute:attr-iframe-srcdoc}](iframe-embed-object.html#attr-iframe-srcdoc)
  on
  [`iframe`{#the-translate-attribute:the-iframe-element}](iframe-embed-object.html#the-iframe-element)
  elements; must be parsed and recursively processed
- [`style`{#the-translate-attribute:attr-style}](#attr-style) on [HTML
  elements](infrastructure.html#html-elements){#the-translate-attribute:html-elements-4};
  must be parsed and recursively processed (e.g. for the values of
  [\'content\'](https://drafts.csswg.org/css2/#content%E2%91%A0){#the-translate-attribute:'content'
  x-internal="'content'"} properties)
- [`title`{#the-translate-attribute:attr-title}](#attr-title) on all
  [HTML
  elements](infrastructure.html#html-elements){#the-translate-attribute:html-elements-5}
- [`value`{#the-translate-attribute:attr-input-value}](input.html#attr-input-value)
  on
  [`input`{#the-translate-attribute:the-input-element}](input.html#the-input-element)
  elements with a
  [`type`{#the-translate-attribute:attr-input-type}](input.html#attr-input-type)
  attribute in the
  [Button](input.html#button-state-(type=button)){#the-translate-attribute:button-state-(type=button)}
  state or the [Reset
  Button](input.html#reset-button-state-(type=reset)){#the-translate-attribute:reset-button-state-(type=reset)}
  state

Other specifications may define other attributes that are also
[translatable
attributes](#translatable-attributes){#the-translate-attribute:translatable-attributes-2}.
For example, ARIA would define the
[`aria-label`{#the-translate-attribute:attr-aria-label}](https://w3c.github.io/aria/#aria-label){x-internal="attr-aria-label"}
attribute as translatable.

------------------------------------------------------------------------

The [`translate`]{#dom-translate .dfn dfn-for="HTMLElement"
dfn-type="attribute"} IDL attribute must, on getting, return true if the
element\'s [translation
mode](#translation-mode){#the-translate-attribute:translation-mode-6} is
[translate-enabled](#translate-enabled){#the-translate-attribute:translate-enabled-5},
and false otherwise. On setting, it must set the content attribute\'s
value to \"`yes`\" if the new value is true, and set the content
attribute\'s value to \"`no`\" otherwise.

::: example
In this example, everything in the document is to be translated when the
page is localized, except the sample keyboard input and sample program
output:

``` html
<!DOCTYPE HTML>
<html lang=en> <!-- default on the document element is translate=yes -->
 <head>
  <title>The Bee Game</title> <!-- implied translate=yes inherited from ancestors -->
 </head>
 <body>
  <p>The Bee Game is a text adventure game in English.</p>
  <p>When the game launches, the first thing you should do is type
  <kbd translate=no>eat honey</kbd>. The game will respond with:</p>
  <pre><samp translate=no>Yum yum! That was some good honey!</samp></pre>
 </body>
</html>
```
:::

##### [3.2.6.4]{.secno} The [`dir`{#the-dir-attribute:attr-dir}](#attr-dir) attribute[](#the-dir-attribute){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Global_attributes/dir](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/dir "The dir global attribute is an enumerated attribute that indicates the directionality of the element's text.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome1+]{.chrome
.yes}

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

The [`dir`]{#attr-dir .dfn dfn-for="html-global"
dfn-type="element-attr"} attribute is an [enumerated
attribute](common-microsyntaxes.html#enumerated-attribute){#the-dir-attribute:enumerated-attribute}
with the following keywords and states:

Keyword

State

Brief description

[`ltr`]{#attr-dir-ltr .dfn dfn-for="html-global/dir"
dfn-type="attr-value"}

[LTR]{#attr-dir-ltr-state .dfn}

The contents of the element are explicitly directionally isolated
left-to-right text.

[`rtl`]{#attr-dir-rtl .dfn dfn-for="html-global/dir"
dfn-type="attr-value"}

[RTL]{#attr-dir-rtl-state .dfn}

The contents of the element are explicitly directionally isolated
right-to-left text.

[`auto`]{#attr-dir-auto .dfn dfn-for="html-global/dir"
dfn-type="attr-value"}

[Auto]{#attr-dir-auto-state .dfn}

The contents of the element are explicitly directionally isolated text,
but the direction is to be determined programmatically using the
contents of the element (as described below).

::: note
The heuristic used by the
[Auto](#attr-dir-auto-state){#the-dir-attribute:attr-dir-auto-state}
state is very crude (it just looks at the first character with a strong
directionality, in a manner analogous to the Paragraph Level
determination in the bidirectional algorithm). Authors are urged to only
use this value as a last resort when the direction of the text is truly
unknown and no better server-side heuristic can be applied.
[\[BIDI\]](references.html#refsBIDI)

For
[`textarea`{#the-dir-attribute:the-textarea-element}](form-elements.html#the-textarea-element)
and
[`pre`{#the-dir-attribute:the-pre-element}](grouping-content.html#the-pre-element)
elements, the heuristic is applied on a per-paragraph level.
:::

The attribute\'s *[missing value
default](common-microsyntaxes.html#missing-value-default)* and *[invalid
value default](common-microsyntaxes.html#invalid-value-default)* are
both the [Undefined]{#attr-dir-undefined-state .dfn} state.

------------------------------------------------------------------------

The [directionality]{#the-directionality .dfn} of an element (any
element, not just an [HTML
element](infrastructure.html#html-elements){#the-dir-attribute:html-elements})
is either \'[ltr]{#concept-ltr .dfn}\' or \'[rtl]{#concept-rtl .dfn}\'.
To compute the
[directionality](#the-directionality){#the-dir-attribute:the-directionality}
given an element `element`{.variable}, switch on `element`{.variable}\'s
[`dir`{#the-dir-attribute:attr-dir-2}](#attr-dir) attribute state:

[LTR](#attr-dir-ltr-state){#the-dir-attribute:attr-dir-ltr-state}
:   Return \'[ltr](#concept-ltr){#the-dir-attribute:concept-ltr}\'.

[RTL](#attr-dir-rtl-state){#the-dir-attribute:attr-dir-rtl-state}
:   Return \'[rtl](#concept-rtl){#the-dir-attribute:concept-rtl}\'.

[Auto](#attr-dir-auto-state){#the-dir-attribute:attr-dir-auto-state-2}

:   1.  Let `result`{.variable} be the [auto
        directionality](#auto-directionality){#the-dir-attribute:auto-directionality}
        of `element`{.variable}.

    2.  If `result`{.variable} is null, then return
        \'[ltr](#concept-ltr){#the-dir-attribute:concept-ltr-2}\'.

    3.  Return `result`{.variable}.

[Undefined](#attr-dir-undefined-state){#the-dir-attribute:attr-dir-undefined-state}

:   

    If `element`{.variable} is a [`bdi`{#the-dir-attribute:the-bdi-element}](text-level-semantics.html#the-bdi-element) element

    :   1.  Let `result`{.variable} be the [auto
            directionality](#auto-directionality){#the-dir-attribute:auto-directionality-2}
            of `element`{.variable}.

        2.  If `result`{.variable} is null, then return
            \'[ltr](#concept-ltr){#the-dir-attribute:concept-ltr-3}\'.

        3.  Return `result`{.variable}.

    If `element`{.variable} is an [`input`{#the-dir-attribute:the-input-element}](input.html#the-input-element) element whose [`type`{#the-dir-attribute:attr-input-type}](input.html#attr-input-type) attribute is in the [Telephone](input.html#telephone-state-(type=tel)){#the-dir-attribute:telephone-state-(type=tel)} state
    :   Return
        \'[ltr](#concept-ltr){#the-dir-attribute:concept-ltr-4}\'.

    Otherwise

    :   Return the [parent
        directionality](#parent-directionality){#the-dir-attribute:parent-directionality}
        of `element`{.variable}.

Since the [`dir`{#the-dir-attribute:attr-dir-3}](#attr-dir) attribute is
only defined for [HTML
elements](infrastructure.html#html-elements){#the-dir-attribute:html-elements-2},
it cannot be present on elements from other namespaces. Thus, elements
from other namespaces always end up using the [parent
directionality](#parent-directionality){#the-dir-attribute:parent-directionality-2}.

The [auto-directionality form-associated
elements]{#auto-directionality-form-associated-elements .dfn} are:

- [`input`{#the-dir-attribute:the-input-element-2}](input.html#the-input-element)
  elements whose
  [`type`{#the-dir-attribute:attr-input-type-2}](input.html#attr-input-type)
  attribute is in the
  [Hidden](input.html#hidden-state-(type=hidden)){#the-dir-attribute:hidden-state-(type=hidden)},
  [Text](input.html#text-(type=text)-state-and-search-state-(type=search)){#the-dir-attribute:text-(type=text)-state-and-search-state-(type=search)},
  [Search](input.html#text-(type=text)-state-and-search-state-(type=search)){#the-dir-attribute:text-(type=text)-state-and-search-state-(type=search)-2},
  [Telephone](input.html#telephone-state-(type=tel)){#the-dir-attribute:telephone-state-(type=tel)-2},
  [URL](input.html#url-state-(type=url)){#the-dir-attribute:url-state-(type=url)},
  [Email](input.html#email-state-(type=email)){#the-dir-attribute:email-state-(type=email)},
  [Password](input.html#password-state-(type=password)){#the-dir-attribute:password-state-(type=password)},
  [Submit
  Button](input.html#submit-button-state-(type=submit)){#the-dir-attribute:submit-button-state-(type=submit)},
  [Reset
  Button](input.html#reset-button-state-(type=reset)){#the-dir-attribute:reset-button-state-(type=reset)},
  or
  [Button](input.html#button-state-(type=button)){#the-dir-attribute:button-state-(type=button)}
  state, and

- [`textarea`{#the-dir-attribute:the-textarea-element-2}](form-elements.html#the-textarea-element)
  elements.

To compute the [auto directionality]{#auto-directionality .dfn} given an
element `element`{.variable}:

1.  If `element`{.variable} is an [auto-directionality form-associated
    element](#auto-directionality-form-associated-elements){#the-dir-attribute:auto-directionality-form-associated-elements}:

    1.  If `element`{.variable}\'s
        [value](form-control-infrastructure.html#concept-fe-value){#the-dir-attribute:concept-fe-value}
        contains a character of bidirectional character type AL or R,
        and there is no character of bidirectional character type L
        anywhere before it in the element\'s
        [value](form-control-infrastructure.html#concept-fe-value){#the-dir-attribute:concept-fe-value-2},
        then return
        \'[rtl](#concept-rtl){#the-dir-attribute:concept-rtl-2}\'.
        [\[BIDI\]](references.html#refsBIDI)

    2.  If `element`{.variable}\'s
        [value](form-control-infrastructure.html#concept-fe-value){#the-dir-attribute:concept-fe-value-3}
        is not the empty string, then return
        \'[ltr](#concept-ltr){#the-dir-attribute:concept-ltr-5}\'.

    3.  Return null.

2.  If `element`{.variable} is a
    [`slot`{#the-dir-attribute:the-slot-element}](scripting.html#the-slot-element)
    element whose
    [root](https://dom.spec.whatwg.org/#concept-tree-root){#the-dir-attribute:root
    x-internal="root"} is a [shadow
    root](https://dom.spec.whatwg.org/#concept-shadow-root){#the-dir-attribute:shadow-root
    x-internal="shadow-root"} and `element`{.variable}\'s [assigned
    nodes](https://dom.spec.whatwg.org/#slot-assigned-nodes){#the-dir-attribute:assigned-nodes
    x-internal="assigned-nodes"} are not empty:

    1.  [For
        each](https://infra.spec.whatwg.org/#list-iterate){#the-dir-attribute:list-iterate
        x-internal="list-iterate"} node `child`{.variable} of
        `element`{.variable}\'s [assigned
        nodes](https://dom.spec.whatwg.org/#slot-assigned-nodes){#the-dir-attribute:assigned-nodes-2
        x-internal="assigned-nodes"}:

        1.  Let `childDirection`{.variable} be null.

        2.  If `child`{.variable} is a
            [`Text`{#the-dir-attribute:text}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
            node, then set `childDirection`{.variable} to the [text node
            directionality](#text-node-directionality){#the-dir-attribute:text-node-directionality}
            of `child`{.variable}.

        3.  Otherwise:

            1.  [Assert](https://infra.spec.whatwg.org/#assert){#the-dir-attribute:assert
                x-internal="assert"}: `child`{.variable} is an
                [`Element`{#the-dir-attribute:element}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"}
                node.

            2.  Set `childDirection`{.variable} to the [contained text
                auto
                directionality](#contained-text-auto-directionality){#the-dir-attribute:contained-text-auto-directionality}
                of `child`{.variable} with
                *[canExcludeRoot](#auto-directionality-can-exclude-root){#the-dir-attribute:auto-directionality-can-exclude-root}*
                set to true.

        4.  If `childDirection`{.variable} is not null, then return
            `childDirection`{.variable}.

    2.  Return null.

3.  Return the [contained text auto
    directionality](#contained-text-auto-directionality){#the-dir-attribute:contained-text-auto-directionality-2}
    of `element`{.variable} with
    *[canExcludeRoot](#auto-directionality-can-exclude-root){#the-dir-attribute:auto-directionality-can-exclude-root-2}*
    set to false.

To compute the [contained text auto
directionality]{#contained-text-auto-directionality .dfn} of an element
`element`{.variable} with a boolean
[`canExcludeRoot`{.variable}]{#auto-directionality-can-exclude-root
.dfn}:

1.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#the-dir-attribute:list-iterate-2
    x-internal="list-iterate"} node `descendant`{.variable} of
    `element`{.variable}\'s
    [descendants](https://dom.spec.whatwg.org/#concept-tree-descendant){#the-dir-attribute:descendant
    x-internal="descendant"}, in [tree
    order](https://dom.spec.whatwg.org/#concept-tree-order){#the-dir-attribute:tree-order
    x-internal="tree-order"}:

    1.  If any of

        - `descendant`{.variable}
        - any ancestor element of `descendant`{.variable} that is a
          descendant of `element`{.variable}
        - if `canExcludeRoot`{.variable} is true, `element`{.variable}

        is one of

        - a
          [`bdi`{#the-dir-attribute:the-bdi-element-2}](text-level-semantics.html#the-bdi-element)
          element
        - a
          [`script`{#the-dir-attribute:the-script-element}](scripting.html#the-script-element)
          element
        - a
          [`style`{#the-dir-attribute:the-style-element}](semantics.html#the-style-element)
          element
        - a
          [`textarea`{#the-dir-attribute:the-textarea-element-3}](form-elements.html#the-textarea-element)
          element
        - an element whose
          [`dir`{#the-dir-attribute:attr-dir-4}](#attr-dir) attribute is
          not in the
          [Undefined](#attr-dir-undefined-state){#the-dir-attribute:attr-dir-undefined-state-2}
          state

        then
        [continue](https://infra.spec.whatwg.org/#iteration-continue){#the-dir-attribute:continue
        x-internal="continue"}.

    2.  If `descendant`{.variable} is a
        [`slot`{#the-dir-attribute:the-slot-element-2}](scripting.html#the-slot-element)
        element whose
        [root](https://dom.spec.whatwg.org/#concept-tree-root){#the-dir-attribute:root-2
        x-internal="root"} is a [shadow
        root](https://dom.spec.whatwg.org/#concept-shadow-root){#the-dir-attribute:shadow-root-2
        x-internal="shadow-root"}, then return the
        [directionality](#the-directionality){#the-dir-attribute:the-directionality-2}
        of that [shadow
        root](https://dom.spec.whatwg.org/#concept-shadow-root){#the-dir-attribute:shadow-root-3
        x-internal="shadow-root"}\'s
        [host](https://dom.spec.whatwg.org/#concept-documentfragment-host){#the-dir-attribute:concept-documentfragment-host
        x-internal="concept-documentfragment-host"}.

    3.  If `descendant`{.variable} is not a
        [`Text`{#the-dir-attribute:text-2}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
        node, then
        [continue](https://infra.spec.whatwg.org/#iteration-continue){#the-dir-attribute:continue-2
        x-internal="continue"}.

    4.  Let `result`{.variable} be the [text node
        directionality](#text-node-directionality){#the-dir-attribute:text-node-directionality-2}
        of `descendant`{.variable}.

    5.  If `result`{.variable} is not null, then return
        `result`{.variable}.

2.  Return null.

To compute the [text node directionality]{#text-node-directionality
.dfn} given a
[`Text`{#the-dir-attribute:text-3}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
node `text`{.variable}:

1.  If `text`{.variable}\'s
    [data](https://dom.spec.whatwg.org/#concept-cd-data){#the-dir-attribute:concept-cd-data
    x-internal="concept-cd-data"} does not contain a code point whose
    bidirectional character type is L, AL, or R, then return null.
    [\[BIDI\]](references.html#refsBIDI)

2.  Let `codePoint`{.variable} be the first code point in
    `text`{.variable}\'s
    [data](https://dom.spec.whatwg.org/#concept-cd-data){#the-dir-attribute:concept-cd-data-2
    x-internal="concept-cd-data"} whose bidirectional character type is
    L, AL, or R.

3.  If `codePoint`{.variable} is of bidirectional character type AL or
    R, then return
    \'[rtl](#concept-rtl){#the-dir-attribute:concept-rtl-3}\'.

4.  If `codePoint`{.variable} is of bidirectional character type L, then
    return \'[ltr](#concept-ltr){#the-dir-attribute:concept-ltr-6}\'.

To compute the [parent directionality]{#parent-directionality .dfn}
given an element `element`{.variable}:

1.  Let `parentNode`{.variable} be `element`{.variable}\'s parent node.

2.  If `parentNode`{.variable} is a [shadow
    root](https://dom.spec.whatwg.org/#concept-shadow-root){#the-dir-attribute:shadow-root-4
    x-internal="shadow-root"}, then return the
    [directionality](#the-directionality){#the-dir-attribute:the-directionality-3}
    of `parentNode`{.variable}\'s
    [host](https://dom.spec.whatwg.org/#concept-documentfragment-host){#the-dir-attribute:concept-documentfragment-host-2
    x-internal="concept-documentfragment-host"}.

3.  If `parentNode`{.variable} is an element, then return the
    [directionality](#the-directionality){#the-dir-attribute:the-directionality-4}
    of `parentNode`{.variable}.

4.  Return \'[ltr](#concept-ltr){#the-dir-attribute:concept-ltr-7}\'.

This attribute [has rendering requirements involving the bidirectional
algorithm](#bidireq).

------------------------------------------------------------------------

The [directionality of an attribute]{#directionality-of-the-attribute
.dfn} of an [HTML
element](infrastructure.html#html-elements){#the-dir-attribute:html-elements-3},
which is used when the text of that attribute is to be included in the
rendering in some manner, is determined as per the first appropriate set
of steps from the following list:

If the attribute is a [directionality-capable attribute](#directionality-capable-attribute){#the-dir-attribute:directionality-capable-attribute} and the element\'s [`dir`{#the-dir-attribute:attr-dir-5}](#attr-dir) attribute is in the [Auto](#attr-dir-auto-state){#the-dir-attribute:attr-dir-auto-state-3} state

:   Find the first character (in logical order) of the attribute\'s
    value that is of bidirectional character type L, AL, or R.
    [\[BIDI\]](references.html#refsBIDI)

    If such a character is found and it is of bidirectional character
    type AL or R, the [directionality of the
    attribute](#directionality-of-the-attribute){#the-dir-attribute:directionality-of-the-attribute}
    is \'[rtl](#concept-rtl){#the-dir-attribute:concept-rtl-4}\'.

    Otherwise, the [directionality of the
    attribute](#directionality-of-the-attribute){#the-dir-attribute:directionality-of-the-attribute-2}
    is \'[ltr](#concept-ltr){#the-dir-attribute:concept-ltr-8}\'.

Otherwise
:   The [directionality of the
    attribute](#directionality-of-the-attribute){#the-dir-attribute:directionality-of-the-attribute-3}
    is the same as [the element\'s
    directionality](#the-directionality){#the-dir-attribute:the-directionality-5}.

The following attributes are [directionality-capable
attributes]{#directionality-capable-attribute .dfn}:

- [`abbr`{#the-dir-attribute:attr-th-abbr}](tables.html#attr-th-abbr) on
  [`th`{#the-dir-attribute:the-th-element}](tables.html#the-th-element)
  elements
- `alt` on
  [`area`{#the-dir-attribute:attr-area-alt}](image-maps.html#attr-area-alt),
  [`img`{#the-dir-attribute:attr-img-alt}](embedded-content.html#attr-img-alt),
  and
  [`input`{#the-dir-attribute:attr-input-alt}](input.html#attr-input-alt)
  elements
- [`content`{#the-dir-attribute:attr-meta-content}](semantics.html#attr-meta-content)
  on
  [`meta`{#the-dir-attribute:the-meta-element}](semantics.html#the-meta-element)
  elements, if the
  [`name`{#the-dir-attribute:attr-meta-name}](semantics.html#attr-meta-name)
  attribute specifies a metadata name whose value is primarily intended
  to be human-readable rather than machine-readable
- `label` on
  [`optgroup`{#the-dir-attribute:attr-optgroup-label}](form-elements.html#attr-optgroup-label),
  [`option`{#the-dir-attribute:attr-option-label}](form-elements.html#attr-option-label),
  and
  [`track`{#the-dir-attribute:attr-track-label}](media.html#attr-track-label)
  elements
- `placeholder` on
  [`input`{#the-dir-attribute:attr-input-placeholder}](input.html#attr-input-placeholder)
  and
  [`textarea`{#the-dir-attribute:attr-textarea-placeholder}](form-elements.html#attr-textarea-placeholder)
  elements
- [`title`{#the-dir-attribute:attr-title}](#attr-title) on all [HTML
  elements](infrastructure.html#html-elements){#the-dir-attribute:html-elements-4}

------------------------------------------------------------------------

`document`{.variable}`.`[`dir`](#dom-dir){#dom-dir-dev}` [ = ``value`{.variable}` ]`

:   Returns [the `html`
    element](#the-html-element-2){#the-dir-attribute:the-html-element-2}\'s
    [`dir`{#the-dir-attribute:attr-dir-6}](#attr-dir) attribute\'s
    value, if any.

    Can be set, to either \"`ltr`\", \"`rtl`\", or \"`auto`\" to replace
    [the `html`
    element](#the-html-element-2){#the-dir-attribute:the-html-element-2-2}\'s
    [`dir`{#the-dir-attribute:attr-dir-7}](#attr-dir) attribute\'s
    value.

    If there is no [`html`
    element](#the-html-element-2){#the-dir-attribute:the-html-element-2-3},
    returns the empty string and ignores new values.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLElement/dir](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/dir "The HTMLElement.dir property gets or sets the text writing directionality of the content of the current element.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The [`dir`]{#dom-dir .dfn dfn-for="HTMLElement" dfn-type="attribute"}
IDL attribute on an element must
[reflect](common-dom-interfaces.html#reflect){#the-dir-attribute:reflect}
the [`dir`{#the-dir-attribute:attr-dir-8}](#attr-dir) content attribute
of that element, [limited to only known
values](common-dom-interfaces.html#limited-to-only-known-values){#the-dir-attribute:limited-to-only-known-values}.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Document/dir](https://developer.mozilla.org/en-US/docs/Web/API/Document/dir "The Document.dir property is a string representing the directionality of the text of the document, whether left to right (default) or right to left. Possible values are 'rtl', right to left, and 'ltr', left to right.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari10.1+]{.safari .yes}[Chrome64+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera51+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android47+]{.opera_android .yes}
:::
::::
:::::

The [`dir`]{#dom-document-dir .dfn dfn-for="Document"
dfn-type="attribute"} IDL attribute on
[`Document`{#the-dir-attribute:document}](#document) objects must
[reflect](common-dom-interfaces.html#reflect){#the-dir-attribute:reflect-2}
the [`dir`{#the-dir-attribute:attr-dir-9}](#attr-dir) content attribute
of [the `html`
element](#the-html-element-2){#the-dir-attribute:the-html-element-2-4},
if any, [limited to only known
values](common-dom-interfaces.html#limited-to-only-known-values){#the-dir-attribute:limited-to-only-known-values-2}.
If there is no such element, then the attribute must return the empty
string and do nothing on setting.

Authors are strongly encouraged to use the
[`dir`{#the-dir-attribute:attr-dir-10}](#attr-dir) attribute to indicate
text direction rather than using CSS, since that way their documents
will continue to render correctly even in the absence of CSS (e.g. as
interpreted by search engines).

::: example
This markup fragment is of an IM conversation.

``` html
<p dir=auto class="u1"><b><bdi>Student</bdi>:</b> How do you write "What's your name?" in Arabic?</p>
<p dir=auto class="u2"><b><bdi>Teacher</bdi>:</b> ما اسمك؟</p>
<p dir=auto class="u1"><b><bdi>Student</bdi>:</b> Thanks.</p>
<p dir=auto class="u2"><b><bdi>Teacher</bdi>:</b> That's written "شكرًا".</p>
<p dir=auto class="u2"><b><bdi>Teacher</bdi>:</b> Do you know how to write "Please"?</p>
<p dir=auto class="u1"><b><bdi>Student</bdi>:</b> "من فضلك", right?</p>
```

Given a suitable style sheet and the default alignment styles for the
[`p`{#the-dir-attribute:the-p-element}](grouping-content.html#the-p-element)
element, namely to align the text to the *start edge* of the paragraph,
the resulting rendering could be as follows:

![Each paragraph rendered as a separate block, with the paragraphs
left-aligned except the second paragraph and the last one, which would
be right aligned, with the usernames (\'Student\' and \'Teacher\' in
this example) flush right, with a colon to their left, and the text
first to the left of that.](/images/im.png){width="366" height="157"}

As noted earlier, the
[`auto`{#the-dir-attribute:attr-dir-auto}](#attr-dir-auto) value is not
a panacea. The final paragraph in this example is misinterpreted as
being right-to-left text, since it begins with an Arabic character,
which causes the \"right?\" to be to the left of the Arabic text.
:::

##### [3.2.6.5]{.secno} The [`style`{#the-style-attribute:attr-style}](#attr-style) attribute[](#the-style-attribute){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Global_attributes/style](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/style "The style global attribute contains CSS styling declarations to be applied to the element. Note that it is recommended for styles to be defined in a separate file or files. This attribute and the <style> element have mainly the purpose of allowing for quick styling, for example for testing purposes.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerYes]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

All [HTML
elements](infrastructure.html#html-elements){#the-style-attribute:html-elements}
may have the [`style`]{#attr-style .dfn dfn-for="html-global"
dfn-type="element-attr"} content attribute set. This is a [style
attribute](https://drafts.csswg.org/css-style-attr/#style-attribute){#the-style-attribute:css-styling-attribute
x-internal="css-styling-attribute"} as defined by CSS Style Attributes.
[\[CSSATTR\]](references.html#refsCSSATTR)

In user agents that support CSS, the attribute\'s value must be parsed
when the attribute is added or has its value changed, according to the
rules given for [style
attributes](https://drafts.csswg.org/css-style-attr/#style-attribute){#the-style-attribute:css-styling-attribute-2
x-internal="css-styling-attribute"}.
[\[CSSATTR\]](references.html#refsCSSATTR)

However, if the [Should element\'s inline behavior be blocked by Content
Security
Policy?](https://w3c.github.io/webappsec-csp/#should-block-inline){#the-style-attribute:should-element's-inline-behavior-be-blocked-by-content-security-policy
x-internal="should-element's-inline-behavior-be-blocked-by-content-security-policy"}
algorithm returns \"`Blocked`\" when executed upon the attribute\'s
[element](https://dom.spec.whatwg.org/#interface-element){#the-style-attribute:element
x-internal="element"}, \"`style attribute`\", and the attribute\'s
value, then the style rules defined in the attribute\'s value must not
be applied to the
[element](https://dom.spec.whatwg.org/#interface-element){#the-style-attribute:element-2
x-internal="element"}. [\[CSP\]](references.html#refsCSP)

Documents that use
[`style`{#the-style-attribute:attr-style-2}](#attr-style) attributes on
any of their elements must still be comprehensible and usable if those
attributes were removed.

In particular, using the
[`style`{#the-style-attribute:attr-style-3}](#attr-style) attribute to
hide and show content, or to convey meaning that is otherwise not
included in the document, is non-conforming. (To hide and show content,
use the
[`hidden`{#the-style-attribute:attr-hidden}](interaction.html#attr-hidden)
attribute.)

------------------------------------------------------------------------

`element`{.variable}`.`[`style`](https://drafts.csswg.org/cssom/#dom-elementcssinlinestyle-style){#dom-style-dev x-internal="dom-style"}

:   Returns a
    [`CSSStyleDeclaration`{#the-style-attribute:cssstyledeclaration}](https://drafts.csswg.org/cssom/#the-cssstyledeclaration-interface){x-internal="cssstyledeclaration"}
    object for the element\'s
    [`style`{#the-style-attribute:attr-style-4}](#attr-style) attribute.

The
[`style`{#the-style-attribute:dom-style}](https://drafts.csswg.org/cssom/#dom-elementcssinlinestyle-style){x-internal="dom-style"}
IDL attribute is defined in CSS Object Model.
[\[CSSOM\]](references.html#refsCSSOM)

::: example
In the following example, the words that refer to colors are marked up
using the
[`span`{#the-style-attribute:the-span-element}](text-level-semantics.html#the-span-element)
element and the
[`style`{#the-style-attribute:attr-style-5}](#attr-style) attribute to
make those words show up in the relevant colors in visual media.

``` html
<p>My sweat suit is <span style="color: green; background:
transparent">green</span> and my eyes are <span style="color: blue;
background: transparent">blue</span>.</p>
```
:::

##### [3.2.6.6]{.secno} [Embedding custom non-visible data]{.dfn} with the [`data-*`{#embedding-custom-non-visible-data-with-the-data-*-attributes:attr-data-*}](#attr-data-*) attributes[](#embedding-custom-non-visible-data-with-the-data-*-attributes){.self-link} {#embedding-custom-non-visible-data-with-the-data-*-attributes}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Global_attributes/data-\*](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/data-* "The data-* global attributes form a class of attributes called custom data attributes, that allow proprietary information to be exchanged between the HTML and its DOM representation by scripts.")

Support in all current engines.

::: support
[Firefox6+]{.firefox .yes}[Safari5.1+]{.safari .yes}[Chrome7+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerYes]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

A [custom data attribute]{#custom-data-attribute .dfn} is an attribute
in no namespace whose name starts with the string
\"[`data-`]{#attr-data-* .dfn}\", has at least one character after the
hyphen, is
[XML-compatible](infrastructure.html#xml-compatible){#embedding-custom-non-visible-data-with-the-data-*-attributes:xml-compatible},
and contains no [ASCII upper
alphas](https://infra.spec.whatwg.org/#ascii-upper-alpha){#embedding-custom-non-visible-data-with-the-data-*-attributes:uppercase-ascii-letters
x-internal="uppercase-ascii-letters"}.

All attribute names on [HTML
elements](infrastructure.html#html-elements){#embedding-custom-non-visible-data-with-the-data-*-attributes:html-elements}
in [HTML
documents](https://dom.spec.whatwg.org/#html-document){#embedding-custom-non-visible-data-with-the-data-*-attributes:html-documents
x-internal="html-documents"} get ASCII-lowercased automatically, so the
restriction on ASCII uppercase letters doesn\'t affect such documents.

[Custom data
attributes](#custom-data-attribute){#embedding-custom-non-visible-data-with-the-data-*-attributes:custom-data-attribute}
are intended to store custom data, state, annotations, and similar,
private to the page or application, for which there are no more
appropriate attributes or elements.

These attributes are not intended for use by software that is not known
to the administrators of the site that uses the attributes. For generic
extensions that are to be used by multiple independent tools, either
this specification should be extended to provide the feature explicitly,
or a technology like
[microdata](microdata.html#microdata){#embedding-custom-non-visible-data-with-the-data-*-attributes:microdata}
should be used (with a standardized vocabulary).

::: example
For instance, a site about music could annotate list items representing
tracks in an album with custom data attributes containing the length of
each track. This information could then be used by the site itself to
allow the user to sort the list by track length, or to filter the list
for tracks of certain lengths.

``` html
<ol>
 <li data-length="2m11s">Beyond The Sea</li>
 ...
</ol>
```

It would be inappropriate, however, for the user to use generic software
not associated with that music site to search for tracks of a certain
length by looking at this data.

This is because these attributes are intended for use by the site\'s own
scripts, and are not a generic extension mechanism for publicly-usable
metadata.
:::

::: example
Similarly, a page author could write markup that provides information
for a translation tool that they are intending to use:

``` html
<p>The third <span data-mytrans-de="Anspruch">claim</span> covers the case of <span
translate="no">HTML</span> markup.</p>
```

In this example, the \"`data-mytrans-de`\" attribute gives specific text
for the MyTrans product to use when translating the phrase \"claim\" to
German. However, the standard
[`translate`{#embedding-custom-non-visible-data-with-the-data-*-attributes:attr-translate}](#attr-translate)
attribute is used to tell it that in all languages, \"HTML\" is to
remain unchanged. When a standard attribute is available, there is no
need for a [custom data
attribute](#custom-data-attribute){#embedding-custom-non-visible-data-with-the-data-*-attributes:custom-data-attribute-2}
to be used.
:::

::: example
In this example, custom data attributes are used to store the result of
a feature detection for
[`PaymentRequest`{#embedding-custom-non-visible-data-with-the-data-*-attributes:paymentrequest}](https://w3c.github.io/payment-request/#dom-paymentrequest){x-internal="paymentrequest"},
which could be used in CSS to style a checkout page differently.

``` html
<script>
 if ('PaymentRequest' in window) {
   document.documentElement.dataset.hasPaymentRequest = '';
 }
</script>
```

Here, the `data-has-payment-request` attribute is effectively being used
as a [boolean
attribute](common-microsyntaxes.html#boolean-attribute){#embedding-custom-non-visible-data-with-the-data-*-attributes:boolean-attribute};
it is enough to check the presence of the attribute. However, if the
author so wishes, it could later be populated with some value, maybe to
indicate limited functionality of the feature.
:::

Every [HTML
element](infrastructure.html#html-elements){#embedding-custom-non-visible-data-with-the-data-*-attributes:html-elements-2}
may have any number of [custom data
attributes](#custom-data-attribute){#embedding-custom-non-visible-data-with-the-data-*-attributes:custom-data-attribute-3}
specified, with any value.

Authors should carefully design such extensions so that when the
attributes are ignored and any associated CSS dropped, the page is still
usable.

User agents must not derive any implementation behavior from these
attributes or values. Specifications intended for user agents must not
define these attributes to have any meaningful values.

JavaScript libraries may use the [custom data
attributes](#custom-data-attribute){#embedding-custom-non-visible-data-with-the-data-*-attributes:custom-data-attribute-4},
as they are considered to be part of the page on which they are used.
Authors of libraries that are reused by many authors are encouraged to
include their name in the attribute names, to reduce the risk of
clashes. Where it makes sense, library authors are also encouraged to
make the exact name used in the attribute names customizable, so that
libraries whose authors unknowingly picked the same name can be used on
the same page, and so that multiple versions of a particular library can
be used on the same page even when those versions are not mutually
compatible.

::: example
For example, a library called \"DoQuery\" could use attribute names like
`data-doquery-range`, and a library called \"jJo\" could use attributes
names like `data-jjo-range`. The jJo library could also provide an API
to set which prefix to use (e.g. `J.setDataPrefix('j2')`, making the
attributes have names like `data-j2-range`).
:::

------------------------------------------------------------------------

`element`{.variable}`.`[`dataset`](#dom-dataset){#dom-dataset-dev}

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLElement/dataset](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/dataset "The dataset read-only property of the HTMLElement interface provides read/write access to custom data attributes (data-*) on elements. It exposes a map of strings (DOMStringMap) with an entry for each data-* attribute.")

Support in all current engines.

::: support
[Firefox6+]{.firefox .yes}[Safari5.1+]{.safari .yes}[Chrome7+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer11]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android3+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::

:::: feature
[SVGElement/dataset](https://developer.mozilla.org/en-US/docs/Web/API/SVGElement/dataset "The dataset read-only property of the HTMLElement interface provides read/write access to custom data attributes (data-*) on elements. It exposes a map of strings (DOMStringMap) with an entry for each data-* attribute.")

Support in all current engines.

::: support
[Firefox51+]{.firefox .yes}[Safari5.1+]{.safari .yes}[Chrome55+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera41+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)17+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android41+]{.opera_android .yes}
:::
::::
:::::::

Returns a
[`DOMStringMap`{#embedding-custom-non-visible-data-with-the-data-*-attributes:domstringmap}](#domstringmap)
object for the element\'s
[`data-*`{#embedding-custom-non-visible-data-with-the-data-*-attributes:attr-data-*-2}](#attr-data-*)
attributes.

Hyphenated names become camel-cased. For example, `data-foo-bar=""`
becomes `element.dataset.fooBar`.

The [`dataset`]{#dom-dataset .dfn dfn-for="HTMLOrSVGElement"
dfn-type="attribute"} IDL attribute provides convenient accessors for
all the
[`data-*`{#embedding-custom-non-visible-data-with-the-data-*-attributes:attr-data-*-3}](#attr-data-*)
attributes on an element. On getting, the
[`dataset`{#embedding-custom-non-visible-data-with-the-data-*-attributes:dom-dataset}](#dom-dataset)
IDL attribute must return a
[`DOMStringMap`{#embedding-custom-non-visible-data-with-the-data-*-attributes:domstringmap-2}](#domstringmap)
whose associated element is this element.

The
[`DOMStringMap`{#embedding-custom-non-visible-data-with-the-data-*-attributes:domstringmap-3}](#domstringmap)
interface is used for the
[`dataset`{#embedding-custom-non-visible-data-with-the-data-*-attributes:dom-dataset-2}](#dom-dataset)
attribute. Each
[`DOMStringMap`{#embedding-custom-non-visible-data-with-the-data-*-attributes:domstringmap-4}](#domstringmap)
has an [associated element]{#concept-domstringmap-element .dfn}.

``` idl
[Exposed=Window,
 LegacyOverrideBuiltIns]
interface DOMStringMap {
  getter DOMString (DOMString name);
  [CEReactions] setter undefined (DOMString name, DOMString value);
  [CEReactions] deleter undefined (DOMString name);
};
```

To [get a `DOMStringMap`\'s name-value
pairs]{#concept-domstringmap-pairs .dfn}, run the following algorithm:

1.  Let `list`{.variable} be an empty list of name-value pairs.

2.  For each content attribute on the
    [`DOMStringMap`{#embedding-custom-non-visible-data-with-the-data-*-attributes:domstringmap-5}](#domstringmap)\'s
    [associated
    element](#concept-domstringmap-element){#embedding-custom-non-visible-data-with-the-data-*-attributes:concept-domstringmap-element}
    whose first five characters are the string \"`data-`\" and whose
    remaining characters (if any) do not include any [ASCII upper
    alphas](https://infra.spec.whatwg.org/#ascii-upper-alpha){#embedding-custom-non-visible-data-with-the-data-*-attributes:uppercase-ascii-letters-2
    x-internal="uppercase-ascii-letters"}, in the order that those
    attributes are listed in the element\'s [attribute
    list](https://dom.spec.whatwg.org/#concept-element-attribute){#embedding-custom-non-visible-data-with-the-data-*-attributes:attribute-list
    x-internal="attribute-list"}, add a name-value pair to
    `list`{.variable} whose name is the attribute\'s name with the first
    five characters removed and whose value is the attribute\'s value.

3.  For each name in `list`{.variable}, for each U+002D HYPHEN-MINUS
    character (-) in the name that is followed by an [ASCII lower
    alpha](https://infra.spec.whatwg.org/#ascii-lower-alpha){#embedding-custom-non-visible-data-with-the-data-*-attributes:lowercase-ascii-letters
    x-internal="lowercase-ascii-letters"}, remove the U+002D
    HYPHEN-MINUS character (-) and replace the character that followed
    it by the same character [converted to ASCII
    uppercase](https://infra.spec.whatwg.org/#ascii-uppercase){#embedding-custom-non-visible-data-with-the-data-*-attributes:converted-to-ascii-uppercase
    x-internal="converted-to-ascii-uppercase"}.

4.  Return `list`{.variable}.

The [supported property
names](https://webidl.spec.whatwg.org/#dfn-supported-property-names){#embedding-custom-non-visible-data-with-the-data-*-attributes:supported-property-names
x-internal="supported-property-names"} on a
[`DOMStringMap`{#embedding-custom-non-visible-data-with-the-data-*-attributes:domstringmap-6}](#domstringmap)
object at any instant are the names of each pair returned from [getting
the `DOMStringMap`\'s name-value
pairs](#concept-domstringmap-pairs){#embedding-custom-non-visible-data-with-the-data-*-attributes:concept-domstringmap-pairs}
at that instant, in the order returned.

To [determine the value of a named
property](https://webidl.spec.whatwg.org/#dfn-determine-the-value-of-a-named-property){#embedding-custom-non-visible-data-with-the-data-*-attributes:determine-the-value-of-a-named-property
x-internal="determine-the-value-of-a-named-property"} `name`{.variable}
for a
[`DOMStringMap`{#embedding-custom-non-visible-data-with-the-data-*-attributes:domstringmap-7}](#domstringmap),
return the value component of the name-value pair whose name component
is `name`{.variable} in the list returned from [getting the
`DOMStringMap`\'s name-value
pairs](#concept-domstringmap-pairs){#embedding-custom-non-visible-data-with-the-data-*-attributes:concept-domstringmap-pairs-2}.

To [set the value of a new named
property](https://webidl.spec.whatwg.org/#dfn-set-the-value-of-a-new-named-property){#embedding-custom-non-visible-data-with-the-data-*-attributes:set-the-value-of-a-new-named-property
x-internal="set-the-value-of-a-new-named-property"} or [set the value of
an existing named
property](https://webidl.spec.whatwg.org/#dfn-set-the-value-of-an-existing-named-property){#embedding-custom-non-visible-data-with-the-data-*-attributes:set-the-value-of-an-existing-named-property
x-internal="set-the-value-of-an-existing-named-property"} for a
[`DOMStringMap`{#embedding-custom-non-visible-data-with-the-data-*-attributes:domstringmap-8}](#domstringmap),
given a property name `name`{.variable} and a new value
`value`{.variable}, run the following steps:

1.  If `name`{.variable} contains a U+002D HYPHEN-MINUS character (-)
    followed by an [ASCII lower
    alpha](https://infra.spec.whatwg.org/#ascii-lower-alpha){#embedding-custom-non-visible-data-with-the-data-*-attributes:lowercase-ascii-letters-2
    x-internal="lowercase-ascii-letters"}, then throw a
    [\"`SyntaxError`\"](https://webidl.spec.whatwg.org/#syntaxerror){#embedding-custom-non-visible-data-with-the-data-*-attributes:syntaxerror
    x-internal="syntaxerror"}
    [`DOMException`{#embedding-custom-non-visible-data-with-the-data-*-attributes:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  For each [ASCII upper
    alpha](https://infra.spec.whatwg.org/#ascii-upper-alpha){#embedding-custom-non-visible-data-with-the-data-*-attributes:uppercase-ascii-letters-3
    x-internal="uppercase-ascii-letters"} in `name`{.variable}, insert a
    U+002D HYPHEN-MINUS character (-) before the character and replace
    the character with the same character [converted to ASCII
    lowercase](https://infra.spec.whatwg.org/#ascii-lowercase){#embedding-custom-non-visible-data-with-the-data-*-attributes:converted-to-ascii-lowercase
    x-internal="converted-to-ascii-lowercase"}.

3.  Insert the string `data-` at the front of `name`{.variable}.

4.  If `name`{.variable} does not match the XML
    [`Name`{#embedding-custom-non-visible-data-with-the-data-*-attributes:xml-name}](https://www.w3.org/TR/xml/#NT-Name){x-internal="xml-name"}
    production, throw an
    [\"`InvalidCharacterError`\"](https://webidl.spec.whatwg.org/#invalidcharactererror){#embedding-custom-non-visible-data-with-the-data-*-attributes:invalidcharactererror
    x-internal="invalidcharactererror"}
    [`DOMException`{#embedding-custom-non-visible-data-with-the-data-*-attributes:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

5.  [Set an attribute
    value](https://dom.spec.whatwg.org/#concept-element-attributes-set-value){#embedding-custom-non-visible-data-with-the-data-*-attributes:concept-element-attributes-set-value
    x-internal="concept-element-attributes-set-value"} for the
    [`DOMStringMap`{#embedding-custom-non-visible-data-with-the-data-*-attributes:domstringmap-9}](#domstringmap)\'s
    [associated
    element](#concept-domstringmap-element){#embedding-custom-non-visible-data-with-the-data-*-attributes:concept-domstringmap-element-2}
    using `name`{.variable} and `value`{.variable}.

To [delete an existing named
property](https://webidl.spec.whatwg.org/#dfn-delete-an-existing-named-property){#embedding-custom-non-visible-data-with-the-data-*-attributes:delete-an-existing-named-property
x-internal="delete-an-existing-named-property"} `name`{.variable} for a
[`DOMStringMap`{#embedding-custom-non-visible-data-with-the-data-*-attributes:domstringmap-10}](#domstringmap),
run the following steps:

1.  For each [ASCII upper
    alpha](https://infra.spec.whatwg.org/#ascii-upper-alpha){#embedding-custom-non-visible-data-with-the-data-*-attributes:uppercase-ascii-letters-4
    x-internal="uppercase-ascii-letters"} in `name`{.variable}, insert a
    U+002D HYPHEN-MINUS character (-) before the character and replace
    the character with the same character [converted to ASCII
    lowercase](https://infra.spec.whatwg.org/#ascii-lowercase){#embedding-custom-non-visible-data-with-the-data-*-attributes:converted-to-ascii-lowercase-2
    x-internal="converted-to-ascii-lowercase"}.

2.  Insert the string `data-` at the front of `name`{.variable}.

3.  [Remove an attribute by
    name](https://dom.spec.whatwg.org/#concept-element-attributes-remove){#embedding-custom-non-visible-data-with-the-data-*-attributes:concept-element-attributes-remove
    x-internal="concept-element-attributes-remove"} given
    `name`{.variable} and the
    [`DOMStringMap`{#embedding-custom-non-visible-data-with-the-data-*-attributes:domstringmap-11}](#domstringmap)\'s
    [associated
    element](#concept-domstringmap-element){#embedding-custom-non-visible-data-with-the-data-*-attributes:concept-domstringmap-element-3}.

This algorithm will only get invoked by Web IDL for names that are given
by the earlier algorithm for [getting the `DOMStringMap`\'s name-value
pairs](#concept-domstringmap-pairs){#embedding-custom-non-visible-data-with-the-data-*-attributes:concept-domstringmap-pairs-3}.
[\[WEBIDL\]](references.html#refsWEBIDL)

::: example
If a web page wanted an element to represent a space ship, e.g. as part
of a game, it would have to use the
[`class`{#embedding-custom-non-visible-data-with-the-data-*-attributes:classes}](#classes)
attribute along with
[`data-*`{#embedding-custom-non-visible-data-with-the-data-*-attributes:attr-data-*-4}](#attr-data-*)
attributes:

``` html
<div class="spaceship" data-ship-id="92432"
     data-weapons="laser 2" data-shields="50%"
     data-x="30" data-y="10" data-z="90">
 <button class="fire"
         onclick="spaceships[this.parentNode.dataset.shipId].fire()">
  Fire
 </button>
</div>
```

Notice how the hyphenated attribute name becomes camel-cased in the API.
:::

::: example
Given the following fragment and elements with similar constructions:

``` html
<img class="tower" id="tower5" data-x="12" data-y="5"
     data-ai="robotarget" data-hp="46" data-ability="flames"
     src="towers/rocket.png" alt="Rocket Tower">
```

\...one could imagine a function `splashDamage()` that takes some
arguments, the first of which is the element to process:

``` js
function splashDamage(node, x, y, damage) {
  if (node.classList.contains('tower') && // checking the 'class' attribute
      node.dataset.x == x && // reading the 'data-x' attribute
      node.dataset.y == y) { // reading the 'data-y' attribute
    var hp = parseInt(node.dataset.hp); // reading the 'data-hp' attribute
    hp = hp - damage;
    if (hp < 0) {
      hp = 0;
      node.dataset.ai = 'dead'; // setting the 'data-ai' attribute
      delete node.dataset.ability; // removing the 'data-ability' attribute
    }
    node.dataset.hp = hp; // setting the 'data-hp' attribute
  }
}
```
:::

#### [3.2.7]{.secno} The [`innerText`{#the-innertext-idl-attribute:dom-innertext}](#dom-innertext) and [`outerText`{#the-innertext-idl-attribute:dom-outertext}](#dom-outertext) properties[](#the-innertext-idl-attribute){.self-link} {#the-innertext-idl-attribute}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[HTMLElement/innerText](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/innerText "The innerText property of the HTMLElement interface represents the rendered text content of a node and its descendants.")

Support in all current engines.

::: support
[Firefox45+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera9.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android1+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

`element`{.variable}`.`[`innerText`](#dom-innertext){#dom-innertext-dev}` [ = ``value`{.variable}` ]`

:   Returns the element\'s text content \"as rendered\".

    Can be set, to replace the element\'s children with the given value,
    but with line breaks converted to
    [`br`{#the-innertext-idl-attribute:the-br-element}](text-level-semantics.html#the-br-element)
    elements.

`element`{.variable}`.`[`outerText`](#dom-outertext){#dom-outertext-dev}` [ = ``value`{.variable}` ]`

:   Returns the element\'s text content \"as rendered\".

    Can be set, to replace the element with the given value, but with
    line breaks converted to
    [`br`{#the-innertext-idl-attribute:the-br-element-2}](text-level-semantics.html#the-br-element)
    elements.

The [get the text steps]{#get-the-text-steps .dfn export=""}, given an
[HTMLElement](#htmlelement){#the-innertext-idl-attribute:htmlelement}
`element`{.variable}, are:

1.  If `element`{.variable} is not [being
    rendered](rendering.html#being-rendered){#the-innertext-idl-attribute:being-rendered}
    or if the user agent is a non-CSS user agent, then return
    `element`{.variable}\'s [descendant text
    content](https://dom.spec.whatwg.org/#concept-descendant-text-content){#the-innertext-idl-attribute:descendant-text-content
    x-internal="descendant-text-content"}.

    This step can produce surprising results, as when the
    [`innerText`{#the-innertext-idl-attribute:dom-innertext-2}](#dom-innertext)
    getter is invoked on an element not [being
    rendered](rendering.html#being-rendered){#the-innertext-idl-attribute:being-rendered-2},
    its text contents are returned, but when accessed on an element that
    is [being
    rendered](rendering.html#being-rendered){#the-innertext-idl-attribute:being-rendered-3},
    all of its children that are not [being
    rendered](rendering.html#being-rendered){#the-innertext-idl-attribute:being-rendered-4}
    have their text contents ignored.

2.  Let `results`{.variable} be a new empty
    [list](https://infra.spec.whatwg.org/#list){#the-innertext-idl-attribute:list
    x-internal="list"}.

3.  For each child node `node`{.variable} of `element`{.variable}:

    1.  Let `current`{.variable} be the
        [list](https://infra.spec.whatwg.org/#list){#the-innertext-idl-attribute:list-2
        x-internal="list"} resulting in running the [rendered text
        collection
        steps](#rendered-text-collection-steps){#the-innertext-idl-attribute:rendered-text-collection-steps}
        with `node`{.variable}. Each item in `results`{.variable} will
        either be a
        [string](https://infra.spec.whatwg.org/#string){#the-innertext-idl-attribute:string
        x-internal="string"} or a positive integer (a *required line
        break count*).

        Intuitively, a *required line break count* item means that a
        certain number of line breaks appear at that point, but they can
        be collapsed with the line breaks induced by adjacent *required
        line break count* items, reminiscent to CSS margin-collapsing.

    2.  For each item `item`{.variable} in `current`{.variable}, append
        `item`{.variable} to `results`{.variable}.

4.  [Remove](https://infra.spec.whatwg.org/#list-remove){#the-innertext-idl-attribute:list-remove
    x-internal="list-remove"} any items from `results`{.variable} that
    are the empty string.

5.  [Remove](https://infra.spec.whatwg.org/#list-remove){#the-innertext-idl-attribute:list-remove-2
    x-internal="list-remove"} any runs of consecutive *required line
    break count* items at the start or end of `results`{.variable}.

6.  [Replace](https://infra.spec.whatwg.org/#list-replace){#the-innertext-idl-attribute:list-replace
    x-internal="list-replace"} each remaining run of consecutive
    *required line break count* items with a string consisting of as
    many U+000A LF code points as the maximum of the values in the
    *required line break count* items.

7.  Return the concatenation of the string items in
    `results`{.variable}.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLElement/outerText](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/outerText "The outerText property of the HTMLElement interface returns the same value as HTMLElement.innerText. When used as a setter it replaces the whole current node with the given text (this differs from innerText, which replaces the content inside the current node).")

Support in all current engines.

::: support
[Firefox98+]{.firefox .yes}[Safari1.3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android1+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The [`innerText`]{#dom-innertext .dfn dfn-for="HTMLElement"
dfn-type="attribute"} and [`outerText`]{#dom-outertext .dfn
dfn-for="HTMLElement" dfn-type="attribute"} getter steps are to return
the result of running [get the text
steps](#get-the-text-steps){#the-innertext-idl-attribute:get-the-text-steps}
with
[this](https://webidl.spec.whatwg.org/#this){#the-innertext-idl-attribute:this
x-internal="this"}.

The [rendered text collection steps]{#rendered-text-collection-steps
.dfn}, given a
[node](https://dom.spec.whatwg.org/#interface-node){#the-innertext-idl-attribute:node
x-internal="node"} `node`{.variable}, are as follows:

1.  Let `items`{.variable} be the result of running the [rendered text
    collection
    steps](#rendered-text-collection-steps){#the-innertext-idl-attribute:rendered-text-collection-steps-2}
    with each child node of `node`{.variable} in [tree
    order](https://dom.spec.whatwg.org/#concept-tree-order){#the-innertext-idl-attribute:tree-order
    x-internal="tree-order"}, and then concatenating the results to a
    single
    [list](https://infra.spec.whatwg.org/#list){#the-innertext-idl-attribute:list-3
    x-internal="list"}.

2.  If `node`{.variable}\'s [computed
    value](https://drafts.csswg.org/css-cascade/#computed-value){#the-innertext-idl-attribute:computed-value
    x-internal="computed-value"} of
    [\'visibility\'](https://drafts.csswg.org/css2/#propdef-visibility){#the-innertext-idl-attribute:'visibility'
    x-internal="'visibility'"} is not \'visible\', then return
    `items`{.variable}.

3.  If `node`{.variable} is not [being
    rendered](rendering.html#being-rendered){#the-innertext-idl-attribute:being-rendered-5},
    then return `items`{.variable}. For the purpose of this step, the
    following elements must act as described if the [computed
    value](https://drafts.csswg.org/css-cascade/#computed-value){#the-innertext-idl-attribute:computed-value-2
    x-internal="computed-value"} of the
    [\'display\'](https://drafts.csswg.org/css2/#display-prop){#the-innertext-idl-attribute:'display'
    x-internal="'display'"} property is not \'none\':

    - [`select`{#the-innertext-idl-attribute:the-select-element}](form-elements.html#the-select-element)
      elements have an associated non-replaced inline [CSS
      box](https://drafts.csswg.org/css-display/#css-box){#the-innertext-idl-attribute:css-box
      x-internal="css-box"} whose child boxes include only those of
      [`optgroup`{#the-innertext-idl-attribute:the-optgroup-element}](form-elements.html#the-optgroup-element)
      and
      [`option`{#the-innertext-idl-attribute:the-option-element}](form-elements.html#the-option-element)
      element child nodes;

    - [`optgroup`{#the-innertext-idl-attribute:the-optgroup-element-2}](form-elements.html#the-optgroup-element)
      elements have an associated non-replaced block-level [CSS
      box](https://drafts.csswg.org/css-display/#css-box){#the-innertext-idl-attribute:css-box-2
      x-internal="css-box"} whose child boxes include only those of
      [`option`{#the-innertext-idl-attribute:the-option-element-2}](form-elements.html#the-option-element)
      element child nodes; and

    - [`option`{#the-innertext-idl-attribute:the-option-element-3}](form-elements.html#the-option-element)
      elements have an associated non-replaced block-level [CSS
      box](https://drafts.csswg.org/css-display/#css-box){#the-innertext-idl-attribute:css-box-3
      x-internal="css-box"} whose child boxes are as normal for
      non-replaced block-level [CSS
      boxes](https://drafts.csswg.org/css-display/#css-box){#the-innertext-idl-attribute:css-box-4
      x-internal="css-box"}.

    `items`{.variable} can be non-empty due to \'display:contents\'.

4.  If `node`{.variable} is a
    [`Text`{#the-innertext-idl-attribute:text}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
    node, then for each CSS text box produced by `node`{.variable}, in
    content order, compute the text of the box after application of the
    CSS
    [\'white-space\'](https://drafts.csswg.org/css-text/#white-space-property){#the-innertext-idl-attribute:'white-space'
    x-internal="'white-space'"} processing rules and
    [\'text-transform\'](https://drafts.csswg.org/css-text/#text-transform-property){#the-innertext-idl-attribute:'text-transform'
    x-internal="'text-transform'"} rules, set `items`{.variable} to the
    [list](https://infra.spec.whatwg.org/#list){#the-innertext-idl-attribute:list-4
    x-internal="list"} of the resulting strings, and return
    `items`{.variable}. The CSS
    [\'white-space\'](https://drafts.csswg.org/css-text/#white-space-property){#the-innertext-idl-attribute:'white-space'-2
    x-internal="'white-space'"} processing rules are slightly modified:
    collapsible spaces at the end of lines are always collapsed, but
    they are only removed if the line is the last line of the block, or
    it ends with a
    [`br`{#the-innertext-idl-attribute:the-br-element-3}](text-level-semantics.html#the-br-element)
    element. Soft hyphens should be preserved.
    [\[CSSTEXT\]](references.html#refsCSSTEXT)

5.  If `node`{.variable} is a
    [`br`{#the-innertext-idl-attribute:the-br-element-4}](text-level-semantics.html#the-br-element)
    element, then
    [append](https://infra.spec.whatwg.org/#list-append){#the-innertext-idl-attribute:list-append
    x-internal="list-append"} a string containing a single U+000A LF
    code point to `items`{.variable}.

6.  If `node`{.variable}\'s [computed
    value](https://drafts.csswg.org/css-cascade/#computed-value){#the-innertext-idl-attribute:computed-value-3
    x-internal="computed-value"} of
    [\'display\'](https://drafts.csswg.org/css2/#display-prop){#the-innertext-idl-attribute:'display'-2
    x-internal="'display'"} is
    [\'table-cell\'](https://drafts.csswg.org/css-tables/#table-cell){#the-innertext-idl-attribute:'table-cell'
    x-internal="'table-cell'"}, and `node`{.variable}\'s [CSS
    box](https://drafts.csswg.org/css-display/#css-box){#the-innertext-idl-attribute:css-box-5
    x-internal="css-box"} is not the last
    [\'table-cell\'](https://drafts.csswg.org/css-tables/#table-cell){#the-innertext-idl-attribute:'table-cell'-2
    x-internal="'table-cell'"} box of its enclosing
    [\'table-row\'](https://drafts.csswg.org/css-tables/#table-row){#the-innertext-idl-attribute:'table-row'
    x-internal="'table-row'"} box, then
    [append](https://infra.spec.whatwg.org/#list-append){#the-innertext-idl-attribute:list-append-2
    x-internal="list-append"} a string containing a single U+0009 TAB
    code point to `items`{.variable}.

7.  If `node`{.variable}\'s [computed
    value](https://drafts.csswg.org/css-cascade/#computed-value){#the-innertext-idl-attribute:computed-value-4
    x-internal="computed-value"} of
    [\'display\'](https://drafts.csswg.org/css2/#display-prop){#the-innertext-idl-attribute:'display'-3
    x-internal="'display'"} is
    [\'table-row\'](https://drafts.csswg.org/css-tables/#table-row){#the-innertext-idl-attribute:'table-row'-2
    x-internal="'table-row'"}, and `node`{.variable}\'s [CSS
    box](https://drafts.csswg.org/css-display/#css-box){#the-innertext-idl-attribute:css-box-6
    x-internal="css-box"} is not the last
    [\'table-row\'](https://drafts.csswg.org/css-tables/#table-row){#the-innertext-idl-attribute:'table-row'-3
    x-internal="'table-row'"} box of the nearest ancestor
    [\'table\'](https://drafts.csswg.org/css-tables/#table){#the-innertext-idl-attribute:'table'
    x-internal="'table'"} box, then
    [append](https://infra.spec.whatwg.org/#list-append){#the-innertext-idl-attribute:list-append-3
    x-internal="list-append"} a string containing a single U+000A LF
    code point to `items`{.variable}.

8.  If `node`{.variable} is a
    [`p`{#the-innertext-idl-attribute:the-p-element}](grouping-content.html#the-p-element)
    element, then
    [append](https://infra.spec.whatwg.org/#list-append){#the-innertext-idl-attribute:list-append-4
    x-internal="list-append"} 2 (a *required line break count*) at the
    beginning and end of `items`{.variable}.

9.  If `node`{.variable}\'s [used
    value](https://drafts.csswg.org/css-cascade/#used-value){#the-innertext-idl-attribute:used-value
    x-internal="used-value"} of
    [\'display\'](https://drafts.csswg.org/css2/#display-prop){#the-innertext-idl-attribute:'display'-4
    x-internal="'display'"} is
    [block-level](https://drafts.csswg.org/css-display/#block-level){#the-innertext-idl-attribute:block-level
    x-internal="block-level"} or
    [\'table-caption\'](https://drafts.csswg.org/css-tables/#table-caption){#the-innertext-idl-attribute:'table-caption'
    x-internal="'table-caption'"}, then
    [append](https://infra.spec.whatwg.org/#list-append){#the-innertext-idl-attribute:list-append-5
    x-internal="list-append"} 1 (a *required line break count*) at the
    beginning and end of `items`{.variable}.
    [\[CSSDISPLAY\]](references.html#refsCSSDISPLAY)

    Floats and absolutely-positioned elements fall into this category.

10. Return `items`{.variable}.

Note that descendant nodes of most replaced elements (e.g.,
[`textarea`{#the-innertext-idl-attribute:the-textarea-element}](form-elements.html#the-textarea-element),
[`input`{#the-innertext-idl-attribute:the-input-element}](input.html#the-input-element),
and
[`video`{#the-innertext-idl-attribute:the-video-element}](media.html#the-video-element)
--- but not
[`button`{#the-innertext-idl-attribute:the-button-element}](form-elements.html#the-button-element))
are not rendered by CSS, strictly speaking, and therefore have no [CSS
boxes](https://drafts.csswg.org/css-display/#css-box){#the-innertext-idl-attribute:css-box-7
x-internal="css-box"} for the purposes of this algorithm.

This algorithm is amenable to being generalized to work on
[ranges](https://dom.spec.whatwg.org/#concept-range){#the-innertext-idl-attribute:concept-range
x-internal="concept-range"}. Then we can use it as the basis for
[`Selection`{#the-innertext-idl-attribute:selection}](https://w3c.github.io/selection-api/#selection-interface){x-internal="selection"}\'s
stringifier and maybe expose it directly on
[ranges](https://dom.spec.whatwg.org/#concept-range){#the-innertext-idl-attribute:concept-range-2
x-internal="concept-range"}. See [Bugzilla bug
10583](https://www.w3.org/Bugs/Public/show_bug.cgi?id=10583).

------------------------------------------------------------------------

The [set the inner text steps]{#set-the-inner-text-steps .dfn
export=""}, given an
[HTMLElement](#htmlelement){#the-innertext-idl-attribute:htmlelement-2}
`element`{.variable}, and a string `value`{.variable} are:

1.  Let `fragment`{.variable} be the [rendered text
    fragment](#rendered-text-fragment){#the-innertext-idl-attribute:rendered-text-fragment}
    for `value`{.variable} given `element`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-innertext-idl-attribute:node-document
    x-internal="node-document"}.

2.  [Replace
    all](https://dom.spec.whatwg.org/#concept-node-replace-all){#the-innertext-idl-attribute:concept-node-replace-all
    x-internal="concept-node-replace-all"} with `fragment`{.variable}
    within `element`{.variable}.

The
[`innerText`{#the-innertext-idl-attribute:dom-innertext-3}](#dom-innertext)
setter steps are to run [set the inner text
steps](#set-the-inner-text-steps){#the-innertext-idl-attribute:set-the-inner-text-steps}.

The
[`outerText`{#the-innertext-idl-attribute:dom-outertext-2}](#dom-outertext)
setter steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-innertext-idl-attribute:this-2
    x-internal="this"}\'s parent is null, then throw a
    [\"`NoModificationAllowedError`\"](https://webidl.spec.whatwg.org/#nomodificationallowederror){#the-innertext-idl-attribute:nomodificationallowederror
    x-internal="nomodificationallowederror"}
    [`DOMException`{#the-innertext-idl-attribute:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  Let `next`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#the-innertext-idl-attribute:this-3
    x-internal="this"}\'s [next
    sibling](https://dom.spec.whatwg.org/#concept-tree-next-sibling){#the-innertext-idl-attribute:next-sibling
    x-internal="next-sibling"}.

3.  Let `previous`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#the-innertext-idl-attribute:this-4
    x-internal="this"}\'s [previous
    sibling](https://dom.spec.whatwg.org/#concept-tree-previous-sibling){#the-innertext-idl-attribute:previous-sibling
    x-internal="previous-sibling"}.

4.  Let `fragment`{.variable} be the [rendered text
    fragment](#rendered-text-fragment){#the-innertext-idl-attribute:rendered-text-fragment-2}
    for the given value given
    [this](https://webidl.spec.whatwg.org/#this){#the-innertext-idl-attribute:this-5
    x-internal="this"}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-innertext-idl-attribute:node-document-2
    x-internal="node-document"}.

5.  If `fragment`{.variable} has no
    [children](https://dom.spec.whatwg.org/#concept-tree-child){#the-innertext-idl-attribute:concept-tree-child
    x-internal="concept-tree-child"}, then
    [append](https://dom.spec.whatwg.org/#concept-node-append){#the-innertext-idl-attribute:concept-node-append
    x-internal="concept-node-append"} a new
    [`Text`{#the-innertext-idl-attribute:text-2}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
    node whose
    [data](https://dom.spec.whatwg.org/#concept-cd-data){#the-innertext-idl-attribute:concept-cd-data
    x-internal="concept-cd-data"} is the empty string and [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-innertext-idl-attribute:node-document-3
    x-internal="node-document"} is
    [this](https://webidl.spec.whatwg.org/#this){#the-innertext-idl-attribute:this-6
    x-internal="this"}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-innertext-idl-attribute:node-document-4
    x-internal="node-document"} to `fragment`{.variable}.

6.  [Replace](https://dom.spec.whatwg.org/#concept-node-replace){#the-innertext-idl-attribute:concept-node-replace
    x-internal="concept-node-replace"}
    [this](https://webidl.spec.whatwg.org/#this){#the-innertext-idl-attribute:this-7
    x-internal="this"} with `fragment`{.variable} within
    [this](https://webidl.spec.whatwg.org/#this){#the-innertext-idl-attribute:this-8
    x-internal="this"}\'s parent.

7.  If `next`{.variable} is non-null and `next`{.variable}\'s [previous
    sibling](https://dom.spec.whatwg.org/#concept-tree-previous-sibling){#the-innertext-idl-attribute:previous-sibling-2
    x-internal="previous-sibling"} is a
    [`Text`{#the-innertext-idl-attribute:text-3}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
    node, then [merge with the next text
    node](#merge-with-the-next-text-node){#the-innertext-idl-attribute:merge-with-the-next-text-node}
    given `next`{.variable}\'s [previous
    sibling](https://dom.spec.whatwg.org/#concept-tree-previous-sibling){#the-innertext-idl-attribute:previous-sibling-3
    x-internal="previous-sibling"}.

8.  If `previous`{.variable} is a
    [`Text`{#the-innertext-idl-attribute:text-4}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
    node, then [merge with the next text
    node](#merge-with-the-next-text-node){#the-innertext-idl-attribute:merge-with-the-next-text-node-2}
    given `previous`{.variable}.

The [rendered text fragment]{#rendered-text-fragment .dfn} for a string
`input`{.variable} given a
[`Document`{#the-innertext-idl-attribute:document}](#document)
`document`{.variable} is the result of running the following steps:

1.  Let `fragment`{.variable} be a new
    [`DocumentFragment`{#the-innertext-idl-attribute:documentfragment}](https://dom.spec.whatwg.org/#interface-documentfragment){x-internal="documentfragment"}
    whose [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-innertext-idl-attribute:node-document-5
    x-internal="node-document"} is `document`{.variable}.

2.  Let `position`{.variable} be a [position
    variable](https://infra.spec.whatwg.org/#string-position-variable){#the-innertext-idl-attribute:position-variable
    x-internal="position-variable"} for `input`{.variable}, initially
    pointing at the start of `input`{.variable}.

3.  Let `text`{.variable} be the empty string.

4.  While `position`{.variable} is not past the end of
    `input`{.variable}:

    1.  [Collect a sequence of code
        points](https://infra.spec.whatwg.org/#collect-a-sequence-of-code-points){#the-innertext-idl-attribute:collect-a-sequence-of-code-points
        x-internal="collect-a-sequence-of-code-points"} that are not
        U+000A LF or U+000D CR from `input`{.variable} given
        `position`{.variable}, and set `text`{.variable} to the result.

    2.  If `text`{.variable} is not the empty string, then
        [append](https://dom.spec.whatwg.org/#concept-node-append){#the-innertext-idl-attribute:concept-node-append-2
        x-internal="concept-node-append"} a new
        [`Text`{#the-innertext-idl-attribute:text-5}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
        node whose
        [data](https://dom.spec.whatwg.org/#concept-cd-data){#the-innertext-idl-attribute:concept-cd-data-2
        x-internal="concept-cd-data"} is `text`{.variable} and [node
        document](https://dom.spec.whatwg.org/#concept-node-document){#the-innertext-idl-attribute:node-document-6
        x-internal="node-document"} is `document`{.variable} to
        `fragment`{.variable}.

    3.  While `position`{.variable} is not past the end of
        `input`{.variable}, and the code point at `position`{.variable}
        is either U+000A LF or U+000D CR:

        1.  If the code point at `position`{.variable} is U+000D CR and
            the next code point is U+000A LF, then advance
            `position`{.variable} to the next code point in
            `input`{.variable}.

        2.  Advance `position`{.variable} to the next code point in
            `input`{.variable}.

        3.  [Append](https://dom.spec.whatwg.org/#concept-node-append){#the-innertext-idl-attribute:concept-node-append-3
            x-internal="concept-node-append"} the result of [creating an
            element](https://dom.spec.whatwg.org/#concept-create-element){#the-innertext-idl-attribute:create-an-element
            x-internal="create-an-element"} given `document`{.variable},
            \"`br`\", and the [HTML
            namespace](https://infra.spec.whatwg.org/#html-namespace){#the-innertext-idl-attribute:html-namespace-2
            x-internal="html-namespace-2"} to `fragment`{.variable}.

5.  Return `fragment`{.variable}.

To [merge with the next text node]{#merge-with-the-next-text-node .dfn}
given a
[`Text`{#the-innertext-idl-attribute:text-6}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
node `node`{.variable}:

1.  Let `next`{.variable} be `node`{.variable}\'s [next
    sibling](https://dom.spec.whatwg.org/#concept-tree-next-sibling){#the-innertext-idl-attribute:next-sibling-2
    x-internal="next-sibling"}.

2.  If `next`{.variable} is not a
    [`Text`{#the-innertext-idl-attribute:text-7}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
    node, then return.

3.  [Replace
    data](https://dom.spec.whatwg.org/#concept-cd-replace){#the-innertext-idl-attribute:replace-data
    x-internal="replace-data"} with `node`{.variable},
    `node`{.variable}\'s
    [data](https://dom.spec.whatwg.org/#concept-cd-data){#the-innertext-idl-attribute:concept-cd-data-3
    x-internal="concept-cd-data"}\'s
    [length](https://infra.spec.whatwg.org/#string-length){#the-innertext-idl-attribute:length
    x-internal="length"}, 0, and `next`{.variable}\'s
    [data](https://dom.spec.whatwg.org/#concept-cd-data){#the-innertext-idl-attribute:concept-cd-data-4
    x-internal="concept-cd-data"}.

4.  [Remove](https://dom.spec.whatwg.org/#concept-node-remove){#the-innertext-idl-attribute:concept-node-remove
    x-internal="concept-node-remove"} `next`{.variable}.

#### [3.2.8]{.secno} Requirements relating to the bidirectional algorithm[](#requirements-relating-to-the-bidirectional-algorithm){.self-link}

##### [3.2.8.1]{.secno} Authoring conformance criteria for bidirectional-algorithm formatting characters[](#authoring-conformance-criteria-for-bidirectional-algorithm-formatting-characters){.self-link}

[Text
content](#text-content){#authoring-conformance-criteria-for-bidirectional-algorithm-formatting-characters:text-content}
in [HTML
elements](infrastructure.html#html-elements){#authoring-conformance-criteria-for-bidirectional-algorithm-formatting-characters:html-elements}
with
[`Text`{#authoring-conformance-criteria-for-bidirectional-algorithm-formatting-characters:text}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
nodes in their
[contents](#concept-html-contents){#authoring-conformance-criteria-for-bidirectional-algorithm-formatting-characters:concept-html-contents},
and text in attributes of [HTML
elements](infrastructure.html#html-elements){#authoring-conformance-criteria-for-bidirectional-algorithm-formatting-characters:html-elements-2}
that allow free-form text, may contain characters in the ranges U+202A
to U+202E and U+2066 to U+2069 (the bidirectional-algorithm formatting
characters). [\[BIDI\]](references.html#refsBIDI)

Authors are encouraged to use the
[`dir`{#authoring-conformance-criteria-for-bidirectional-algorithm-formatting-characters:attr-dir}](#attr-dir)
attribute, the
[`bdo`{#authoring-conformance-criteria-for-bidirectional-algorithm-formatting-characters:the-bdo-element}](text-level-semantics.html#the-bdo-element)
element, and the
[`bdi`{#authoring-conformance-criteria-for-bidirectional-algorithm-formatting-characters:the-bdi-element}](text-level-semantics.html#the-bdi-element)
element, rather than maintaining the bidirectional-algorithm formatting
characters manually. The bidirectional-algorithm formatting characters
interact poorly with CSS.

##### [3.2.8.2]{.secno} User agent conformance criteria[](#user-agent-conformance-criteria){.self-link}

User agents must implement the Unicode bidirectional algorithm to
determine the proper ordering of characters when rendering documents and
parts of documents. [\[BIDI\]](references.html#refsBIDI)

The mapping of HTML to the Unicode bidirectional algorithm must be done
in one of three ways. Either the user agent must implement CSS,
including in particular the CSS
[\'unicode-bidi\'](https://drafts.csswg.org/css-writing-modes/#unicode-bidi){#user-agent-conformance-criteria:'unicode-bidi'
x-internal="'unicode-bidi'"},
[\'direction\'](https://drafts.csswg.org/css-writing-modes/#direction){#user-agent-conformance-criteria:'direction'
x-internal="'direction'"}, and
[\'content\'](https://drafts.csswg.org/css2/#content%E2%91%A0){#user-agent-conformance-criteria:'content'
x-internal="'content'"} properties, and must have, in its user agent
style sheet, the rules using those properties given in this
specification\'s [rendering](rendering.html#rendering) section, or,
alternatively, the user agent must act as if it implemented just the
aforementioned properties and had a user agent style sheet that included
all the aforementioned rules, but without letting style sheets specified
in documents override them, or, alternatively, the user agent must
implement another styling language with equivalent semantics.
[\[CSSGC\]](references.html#refsCSSGC)

The following elements and attributes have requirements defined by the
[rendering](rendering.html#rendering) section that, due to the
requirements in this section, are requirements on all user agents (not
just those that [support the suggested default
rendering](infrastructure.html#renderingUA)):

- [`dir`{#user-agent-conformance-criteria:attr-dir}](#attr-dir)
  attribute
- [`bdi`{#user-agent-conformance-criteria:the-bdi-element}](text-level-semantics.html#the-bdi-element)
  element
- [`bdo`{#user-agent-conformance-criteria:the-bdo-element}](text-level-semantics.html#the-bdo-element)
  element
- [`br`{#user-agent-conformance-criteria:the-br-element}](text-level-semantics.html#the-br-element)
  element
- [`pre`{#user-agent-conformance-criteria:the-pre-element}](grouping-content.html#the-pre-element)
  element
- [`textarea`{#user-agent-conformance-criteria:the-textarea-element}](form-elements.html#the-textarea-element)
  element
- [`wbr`{#user-agent-conformance-criteria:the-wbr-element}](text-level-semantics.html#the-wbr-element)
  element

#### [3.2.9]{.secno} Requirements related to ARIA and to platform accessibility APIs[](#wai-aria){.self-link} {#wai-aria}

User agent requirements for implementing Accessibility API semantics on
[HTML
elements](infrastructure.html#html-elements){#wai-aria:html-elements}
are defined in HTML Accessibility API Mappings. In addition to the rules
there, for a [custom
element](custom-elements.html#custom-element){#wai-aria:custom-element}
`element`{.variable}, the default ARIA role semantics are determined as
follows: [\[HTMLAAM\]](references.html#refsHTMLAAM)

1.  Let `map`{.variable} be `element`{.variable}\'s [internal content
    attribute
    map](custom-elements.html#internal-content-attribute-map){#wai-aria:internal-content-attribute-map}.

2.  If `map`{.variable}\[\"`role`\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#wai-aria:map-exists
    x-internal="map-exists"}, then return it.

3.  Return no role.

Similarly, for a [custom
element](custom-elements.html#custom-element){#wai-aria:custom-element-2}
`element`{.variable}, the default ARIA state and property semantics, for
a state or property named `stateOrProperty`{.variable}, are determined
as follows:

1.  If `element`{.variable}\'s [attached
    internals](custom-elements.html#attached-internals){#wai-aria:attached-internals}
    is non-null:

    1.  If `element`{.variable}\'s [attached
        internals](custom-elements.html#attached-internals){#wai-aria:attached-internals-2}\'s
        [get the `stateOrProperty`{.variable}-associated
        element](common-dom-interfaces.html#attr-associated-element){#wai-aria:attr-associated-element}
        exists, then return the result of running it.

    2.  If `element`{.variable}\'s [attached
        internals](custom-elements.html#attached-internals){#wai-aria:attached-internals-3}\'s
        [get the `stateOrProperty`{.variable}-associated
        elements](common-dom-interfaces.html#attr-associated-elements){#wai-aria:attr-associated-elements}
        exists, then return the result of running it.

2.  If `element`{.variable}\'s [internal content attribute
    map](custom-elements.html#internal-content-attribute-map){#wai-aria:internal-content-attribute-map-2}\[`stateOrProperty`{.variable}\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#wai-aria:map-exists-2
    x-internal="map-exists"}, then return it.

3.  Return the default value for `stateOrProperty`{.variable}.

The \"default semantics\" referred to here are sometimes also called
\"native\", \"implicit\", or \"host language\" semantics in ARIA.
[\[ARIA\]](references.html#refsARIA)

One implication of these definitions is that the default semantics can
change over time. This allows custom elements the same expressivity as
built-in elements; e.g., compare to how the default ARIA role semantics
of an
[`a`{#wai-aria:the-a-element}](text-level-semantics.html#the-a-element)
element change as the
[`href`{#wai-aria:attr-hyperlink-href}](links.html#attr-hyperlink-href)
attribute is added or removed.

For an example of this in action, see [the custom elements
section](custom-elements.html#custom-elements-accessibility-example).

------------------------------------------------------------------------

Conformance checker requirements for checking use of ARIA
[`role`{#wai-aria:attr-aria-role}](infrastructure.html#attr-aria-role)
and [`aria-*`{#wai-aria:attr-aria-*}](infrastructure.html#attr-aria-*)
attributes on [HTML
elements](infrastructure.html#html-elements){#wai-aria:html-elements-2}
are defined in ARIA in HTML.
[\[ARIAHTML\]](references.html#refsARIAHTML)

[← 2.7 Safe passing of structured data](structured-data.html) --- [Table
of Contents](./) --- [4 The elements of HTML →](semantics.html)
