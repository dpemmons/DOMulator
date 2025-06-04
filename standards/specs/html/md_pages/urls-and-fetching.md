HTML Standard

::: {#head .head .with-buttons}
[![WHATWG](https://resources.whatwg.org/logo.svg){.darkmode-aware
width="100" crossorigin="" height="100"}](https://whatwg.org/){.logo}

# HTML {#html .allcaps}

Living Standard --- Last Updated [2 June 2025]{.pubdate}
:::

[← 2.3 Common microsyntaxes](common-microsyntaxes.html) --- [Table of
Contents](./) --- [2.6 Common DOM interfaces
→](common-dom-interfaces.html)

1.  ::: {#toc-infrastructure}
    1.  [[2.4]{.secno} URLs](urls-and-fetching.html#urls)
        1.  [[2.4.1]{.secno}
            Terminology](urls-and-fetching.html#terminology-2)
        2.  [[2.4.2]{.secno} Parsing
            URLs](urls-and-fetching.html#resolving-urls)
        3.  [[2.4.3]{.secno} Dynamic changes to base
            URLs](urls-and-fetching.html#dynamic-changes-to-base-urls)
    2.  [[2.5]{.secno} Fetching
        resources](urls-and-fetching.html#fetching-resources)
        1.  [[2.5.1]{.secno}
            Terminology](urls-and-fetching.html#terminology-3)
        2.  [[2.5.2]{.secno} Determining the type of a
            resource](urls-and-fetching.html#content-type-sniffing)
        3.  [[2.5.3]{.secno} Extracting character encodings from `meta`
            elements](urls-and-fetching.html#extracting-character-encodings-from-meta-elements)
        4.  [[2.5.4]{.secno} CORS settings
            attributes](urls-and-fetching.html#cors-settings-attributes)
        5.  [[2.5.5]{.secno} Referrer policy
            attributes](urls-and-fetching.html#referrer-policy-attributes)
        6.  [[2.5.6]{.secno} Nonce
            attributes](urls-and-fetching.html#nonce-attributes)
        7.  [[2.5.7]{.secno} Lazy loading
            attributes](urls-and-fetching.html#lazy-loading-attributes)
        8.  [[2.5.8]{.secno} Blocking
            attributes](urls-and-fetching.html#blocking-attributes)
        9.  [[2.5.9]{.secno} Fetch priority
            attributes](urls-and-fetching.html#fetch-priority-attributes)
    :::

### [2.4]{.secno} URLs[](#urls){.self-link}

#### [2.4.1]{.secno} Terminology[](#terminology-2){.self-link} {#terminology-2}

A string is a [valid non-empty URL]{#valid-non-empty-url .dfn} if it is
a [valid URL
string](https://url.spec.whatwg.org/#valid-url-string){#terminology-2:valid-url-string
x-internal="valid-url-string"} but it is not the empty string.

A string is a [valid URL potentially surrounded by
spaces]{#valid-url-potentially-surrounded-by-spaces .dfn} if, after
[stripping leading and trailing ASCII
whitespace](https://infra.spec.whatwg.org/#strip-leading-and-trailing-ascii-whitespace){#terminology-2:strip-leading-and-trailing-ascii-whitespace
x-internal="strip-leading-and-trailing-ascii-whitespace"} from it, it is
a [valid URL
string](https://url.spec.whatwg.org/#valid-url-string){#terminology-2:valid-url-string-2
x-internal="valid-url-string"}.

A string is a [valid non-empty URL potentially surrounded by
spaces]{#valid-non-empty-url-potentially-surrounded-by-spaces .dfn} if,
after [stripping leading and trailing ASCII
whitespace](https://infra.spec.whatwg.org/#strip-leading-and-trailing-ascii-whitespace){#terminology-2:strip-leading-and-trailing-ascii-whitespace-2
x-internal="strip-leading-and-trailing-ascii-whitespace"} from it, it is
a [valid non-empty
URL](#valid-non-empty-url){#terminology-2:valid-non-empty-url}.

This specification defines the URL
[`about:legacy-compat`]{#about:legacy-compat .dfn} as a reserved, though
unresolvable,
[`about:`{#terminology-2:about-protocol}](https://www.rfc-editor.org/rfc/rfc6694#section-2){x-internal="about-protocol"}
URL, for use in
[DOCTYPE](syntax.html#syntax-doctype){#terminology-2:syntax-doctype}s in
[HTML
documents](https://dom.spec.whatwg.org/#html-document){#terminology-2:html-documents
x-internal="html-documents"} when needed for compatibility with XML
tools. [\[ABOUT\]](references.html#refsABOUT)

This specification defines the URL [`about:html-kind`]{#about:html-kind
.dfn} as a reserved, though unresolvable,
[`about:`{#terminology-2:about-protocol-2}](https://www.rfc-editor.org/rfc/rfc6694#section-2){x-internal="about-protocol"}
URL, that is used as an identifier for kinds of media tracks.
[\[ABOUT\]](references.html#refsABOUT)

This specification defines the URL [`about:srcdoc`]{#about:srcdoc .dfn}
as a reserved, though unresolvable,
[`about:`{#terminology-2:about-protocol-3}](https://www.rfc-editor.org/rfc/rfc6694#section-2){x-internal="about-protocol"}
URL, that is used as the
[URL](https://dom.spec.whatwg.org/#concept-document-url){#terminology-2:the-document's-address
x-internal="the-document's-address"} of [`iframe` `srcdoc`
documents](iframe-embed-object.html#an-iframe-srcdoc-document){#terminology-2:an-iframe-srcdoc-document}.
[\[ABOUT\]](references.html#refsABOUT)

The [fallback base URL]{#fallback-base-url .dfn} of a
[`Document`{#terminology-2:document}](dom.html#document) object
`document`{.variable} is the [URL
record](https://url.spec.whatwg.org/#concept-url){#terminology-2:url-record
x-internal="url-record"} obtained by running these steps:

1.  If `document`{.variable} is [an `iframe` `srcdoc`
    document](iframe-embed-object.html#an-iframe-srcdoc-document){#terminology-2:an-iframe-srcdoc-document-2},
    then:

    1.  [Assert](https://infra.spec.whatwg.org/#assert){#terminology-2:assert
        x-internal="assert"}: `document`{.variable}\'s [about base
        URL](dom.html#concept-document-about-base-url){#terminology-2:concept-document-about-base-url}
        is non-null.

    2.  Return `document`{.variable}\'s [about base
        URL](dom.html#concept-document-about-base-url){#terminology-2:concept-document-about-base-url-2}.

2.  If `document`{.variable}\'s
    [URL](https://dom.spec.whatwg.org/#concept-document-url){#terminology-2:the-document's-address-2
    x-internal="the-document's-address"} [matches
    `about:blank`](#matches-about:blank){#terminology-2:matches-about:blank}
    and `document`{.variable}\'s [about base
    URL](dom.html#concept-document-about-base-url){#terminology-2:concept-document-about-base-url-3}
    is non-null, then return `document`{.variable}\'s [about base
    URL](dom.html#concept-document-about-base-url){#terminology-2:concept-document-about-base-url-4}.

3.  Return `document`{.variable}\'s
    [URL](https://dom.spec.whatwg.org/#concept-document-url){#terminology-2:the-document's-address-3
    x-internal="the-document's-address"}.

The [document base URL]{#document-base-url .dfn export=""} of a
[`Document`{#terminology-2:document-2}](dom.html#document) object is the
[URL
record](https://url.spec.whatwg.org/#concept-url){#terminology-2:url-record-2
x-internal="url-record"} obtained by running these steps:

1.  If there is no
    [`base`{#terminology-2:the-base-element}](semantics.html#the-base-element)
    element that has an
    [`href`{#terminology-2:attr-base-href}](semantics.html#attr-base-href)
    attribute in the
    [`Document`{#terminology-2:document-3}](dom.html#document), then
    return the
    [`Document`{#terminology-2:document-4}](dom.html#document)\'s
    [fallback base
    URL](#fallback-base-url){#terminology-2:fallback-base-url}.

2.  Otherwise, return the [frozen base
    URL](semantics.html#frozen-base-url){#terminology-2:frozen-base-url}
    of the first
    [`base`{#terminology-2:the-base-element-2}](semantics.html#the-base-element)
    element in the
    [`Document`{#terminology-2:document-5}](dom.html#document) that has
    an
    [`href`{#terminology-2:attr-base-href-2}](semantics.html#attr-base-href)
    attribute, in [tree
    order](https://dom.spec.whatwg.org/#concept-tree-order){#terminology-2:tree-order
    x-internal="tree-order"}.

------------------------------------------------------------------------

A [URL](https://url.spec.whatwg.org/#concept-url){#terminology-2:url
x-internal="url"} [matches `about:blank`]{#matches-about:blank .dfn} if
its
[scheme](https://url.spec.whatwg.org/#concept-url-scheme){#terminology-2:concept-url-scheme
x-internal="concept-url-scheme"} is \"`about`\", its
[path](https://url.spec.whatwg.org/#concept-url-path){#terminology-2:concept-url-path
x-internal="concept-url-path"} contains a single string \"`blank`\", its
[username](https://url.spec.whatwg.org/#concept-url-username){#terminology-2:concept-url-username
x-internal="concept-url-username"} and
[password](https://url.spec.whatwg.org/#concept-url-password){#terminology-2:concept-url-password
x-internal="concept-url-password"} are the empty string, and its
[host](https://url.spec.whatwg.org/#concept-url-host){#terminology-2:concept-url-host
x-internal="concept-url-host"} is null.

Such a URL\'s
[query](https://url.spec.whatwg.org/#concept-url-query){#terminology-2:concept-url-query
x-internal="concept-url-query"} and
[fragment](https://url.spec.whatwg.org/#concept-url-fragment){#terminology-2:concept-url-fragment
x-internal="concept-url-fragment"} can be non-null. For example, the
[URL
record](https://url.spec.whatwg.org/#concept-url){#terminology-2:url-record-3
x-internal="url-record"} created by
[parsing](https://url.spec.whatwg.org/#concept-url-parser){#terminology-2:url-parser
x-internal="url-parser"} \"`about:blank?foo#bar`\" [matches
`about:blank`](#matches-about:blank){#terminology-2:matches-about:blank-2}.

A [URL](https://url.spec.whatwg.org/#concept-url){#terminology-2:url-2
x-internal="url"} [matches `about:srcdoc`]{#matches-about:srcdoc .dfn}
if its
[scheme](https://url.spec.whatwg.org/#concept-url-scheme){#terminology-2:concept-url-scheme-2
x-internal="concept-url-scheme"} is \"`about`\", its
[path](https://url.spec.whatwg.org/#concept-url-path){#terminology-2:concept-url-path-2
x-internal="concept-url-path"} contains a single string \"`srcdoc`\",
its
[query](https://url.spec.whatwg.org/#concept-url-query){#terminology-2:concept-url-query-2
x-internal="concept-url-query"} is null, its
[username](https://url.spec.whatwg.org/#concept-url-username){#terminology-2:concept-url-username-2
x-internal="concept-url-username"} and
[password](https://url.spec.whatwg.org/#concept-url-password){#terminology-2:concept-url-password-2
x-internal="concept-url-password"} are the empty string, and its
[host](https://url.spec.whatwg.org/#concept-url-host){#terminology-2:concept-url-host-2
x-internal="concept-url-host"} is null.

The reason that [matches
`about:srcdoc`](#matches-about:srcdoc){#terminology-2:matches-about:srcdoc}
ensures that the
[URL](https://url.spec.whatwg.org/#concept-url){#terminology-2:url-3
x-internal="url"}\'s
[query](https://url.spec.whatwg.org/#concept-url-query){#terminology-2:concept-url-query-3
x-internal="concept-url-query"} is null is because it is not possible to
create [an `iframe` `srcdoc`
document](iframe-embed-object.html#an-iframe-srcdoc-document){#terminology-2:an-iframe-srcdoc-document-3}
whose
[URL](https://dom.spec.whatwg.org/#concept-document-url){#terminology-2:the-document's-address-4
x-internal="the-document's-address"} has a non-null
[query](https://url.spec.whatwg.org/#concept-url-query){#terminology-2:concept-url-query-4
x-internal="concept-url-query"}, unlike
[`Document`{#terminology-2:document-6}](dom.html#document)s whose
[URL](https://dom.spec.whatwg.org/#concept-document-url){#terminology-2:the-document's-address-5
x-internal="the-document's-address"} [matches
`about:blank`](#matches-about:blank){#terminology-2:matches-about:blank-3}.
In other words, the set of all
[URL](https://url.spec.whatwg.org/#concept-url){#terminology-2:url-4
x-internal="url"}s that [match
`about:srcdoc`](#matches-about:srcdoc){#terminology-2:matches-about:srcdoc-2}
only vary in their
[fragment](https://url.spec.whatwg.org/#concept-url-fragment){#terminology-2:concept-url-fragment-2
x-internal="concept-url-fragment"}.

#### [2.4.2]{.secno} Parsing URLs[](#resolving-urls){.self-link} {#resolving-urls}

Parsing a URL is the process of taking a string and obtaining the [URL
record](https://url.spec.whatwg.org/#concept-url){#resolving-urls:url-record
x-internal="url-record"} that it represents. While this process is
defined in URL, the HTML standard defines several wrappers to abstract
base URLs and encodings. [\[URL\]](references.html#refsURL)

Most new APIs are to use [parse a
URL](#parse-a-url){#resolving-urls:parse-a-url}. Older APIs and HTML
elements might have reason to use [encoding-parse a
URL](#encoding-parsing-a-url){#resolving-urls:encoding-parsing-a-url}.
When a custom base URL is needed or no base URL is desired, the [URL
parser](https://url.spec.whatwg.org/#concept-url-parser){#resolving-urls:url-parser
x-internal="url-parser"} can of course be used directly as well.

To [parse a URL]{#parse-a-url .dfn export=""}, given a string
`url`{.variable}, relative to a
[`Document`{#resolving-urls:document}](dom.html#document) object or
[environment settings
object](webappapis.html#environment-settings-object){#resolving-urls:environment-settings-object}
`environment`{.variable}, run these steps. They return failure or a
[URL](https://url.spec.whatwg.org/#concept-url){#resolving-urls:url
x-internal="url"}.

1.  Let `baseURL`{.variable} be `environment`{.variable}\'s [base
    URL](#document-base-url){#resolving-urls:document-base-url}, if
    `environment`{.variable} is a
    [`Document`{#resolving-urls:document-2}](dom.html#document) object;
    otherwise `environment`{.variable}\'s [API base
    URL](webappapis.html#api-base-url){#resolving-urls:api-base-url}.

2.  Return the result of applying the [URL
    parser](https://url.spec.whatwg.org/#concept-url-parser){#resolving-urls:url-parser-2
    x-internal="url-parser"} to `url`{.variable}, with
    `baseURL`{.variable}.

To [encoding-parse a URL]{#encoding-parsing-a-url .dfn export=""}, given
a string `url`{.variable}, relative to a
[`Document`{#resolving-urls:document-3}](dom.html#document) object or
[environment settings
object](webappapis.html#environment-settings-object){#resolving-urls:environment-settings-object-2}
`environment`{.variable}, run these steps. They return failure or a
[URL](https://url.spec.whatwg.org/#concept-url){#resolving-urls:url-2
x-internal="url"}.

1.  Let `encoding`{.variable} be
    [UTF-8](https://encoding.spec.whatwg.org/#utf-8){#resolving-urls:utf-8
    x-internal="utf-8"}.

2.  If `environment`{.variable} is a
    [`Document`{#resolving-urls:document-4}](dom.html#document) object,
    then set `encoding`{.variable} to `environment`{.variable}\'s
    [character
    encoding](https://dom.spec.whatwg.org/#concept-document-encoding){#resolving-urls:document's-character-encoding
    x-internal="document's-character-encoding"}.

3.  Otherwise, if `environment`{.variable}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#resolving-urls:concept-relevant-global}
    is a
    [`Window`{#resolving-urls:window}](nav-history-apis.html#window)
    object, set `encoding`{.variable} to `environment`{.variable}\'s
    [relevant global
    object](webappapis.html#concept-relevant-global){#resolving-urls:concept-relevant-global-2}\'s
    [associated
    `Document`](nav-history-apis.html#concept-document-window){#resolving-urls:concept-document-window}\'s
    [character
    encoding](https://dom.spec.whatwg.org/#concept-document-encoding){#resolving-urls:document's-character-encoding-2
    x-internal="document's-character-encoding"}.

4.  Let `baseURL`{.variable} be `environment`{.variable}\'s [base
    URL](#document-base-url){#resolving-urls:document-base-url-2}, if
    `environment`{.variable} is a
    [`Document`{#resolving-urls:document-5}](dom.html#document) object;
    otherwise `environment`{.variable}\'s [API base
    URL](webappapis.html#api-base-url){#resolving-urls:api-base-url-2}.

5.  Return the result of applying the [URL
    parser](https://url.spec.whatwg.org/#concept-url-parser){#resolving-urls:url-parser-3
    x-internal="url-parser"} to `url`{.variable}, with
    `baseURL`{.variable} and `encoding`{.variable}.

To [encoding-parse-and-serialize a
URL]{#encoding-parsing-and-serializing-a-url .dfn}, given a string
`url`{.variable}, relative to a
[`Document`{#resolving-urls:document-6}](dom.html#document) object or
[environment settings
object](webappapis.html#environment-settings-object){#resolving-urls:environment-settings-object-3}
`environment`{.variable}, run these steps. They return failure or a
string.

1.  Let `url`{.variable} be the result of [encoding-parsing a
    URL](#encoding-parsing-a-url){#resolving-urls:encoding-parsing-a-url-2}
    given `url`{.variable}, relative to `environment`{.variable}.

2.  If `url`{.variable} is failure, then return failure.

3.  Return the result of applying the [URL
    serializer](https://url.spec.whatwg.org/#concept-url-serializer){#resolving-urls:concept-url-serializer
    x-internal="concept-url-serializer"} to `url`{.variable}.

#### [2.4.3]{.secno} Dynamic changes to base URLs[](#dynamic-changes-to-base-urls){.self-link}

When a document\'s [document base
URL](#document-base-url){#dynamic-changes-to-base-urls:document-base-url}
changes, all elements in that document are [affected by a base URL
change](infrastructure.html#affected-by-a-base-url-change){#dynamic-changes-to-base-urls:affected-by-a-base-url-change}.

The following are [base URL change
steps](infrastructure.html#base-url-change-steps){#dynamic-changes-to-base-urls:base-url-change-steps},
which run when an element is [affected by a base URL
change](infrastructure.html#affected-by-a-base-url-change){#dynamic-changes-to-base-urls:affected-by-a-base-url-change-2}
(as defined by DOM):

If the element creates a [hyperlink](links.html#hyperlink){#dynamic-changes-to-base-urls:hyperlink}

:   If the
    [URL](https://url.spec.whatwg.org/#concept-url){#dynamic-changes-to-base-urls:url
    x-internal="url"} identified by the hyperlink is being shown to the
    user, or if any data derived from that
    [URL](https://url.spec.whatwg.org/#concept-url){#dynamic-changes-to-base-urls:url-2
    x-internal="url"} is affecting the display, then the
    [`href`{#dynamic-changes-to-base-urls:attr-hyperlink-href}](links.html#attr-hyperlink-href)
    attribute\'s value should be
    [reparsed](#encoding-parsing-a-url){#dynamic-changes-to-base-urls:encoding-parsing-a-url},
    relative to the element\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#dynamic-changes-to-base-urls:node-document
    x-internal="node-document"} and the UI updated appropriately.

    For example, the CSS
    [`:link`{#dynamic-changes-to-base-urls:selector-link}](semantics-other.html#selector-link)/[`:visited`{#dynamic-changes-to-base-urls:selector-visited}](semantics-other.html#selector-visited)
    [pseudo-classes](https://drafts.csswg.org/selectors/#pseudo-class){#dynamic-changes-to-base-urls:pseudo-class
    x-internal="pseudo-class"} might have been affected.

    If the hyperlink has a
    [`ping`{#dynamic-changes-to-base-urls:ping}](links.html#ping)
    attribute and its
    [URL(s)](https://url.spec.whatwg.org/#concept-url){#dynamic-changes-to-base-urls:url-3
    x-internal="url"} are being shown to the user, then the
    [`ping`{#dynamic-changes-to-base-urls:ping-2}](links.html#ping)
    attribute\'s tokens should be
    [reparsed](#encoding-parsing-a-url){#dynamic-changes-to-base-urls:encoding-parsing-a-url-2},
    relative to the element\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#dynamic-changes-to-base-urls:node-document-2
    x-internal="node-document"} and the UI updated appropriately.

If the element is a [`q`{#dynamic-changes-to-base-urls:the-q-element}](text-level-semantics.html#the-q-element), [`blockquote`{#dynamic-changes-to-base-urls:the-blockquote-element}](grouping-content.html#the-blockquote-element), [`ins`{#dynamic-changes-to-base-urls:the-ins-element}](edits.html#the-ins-element), or [`del`{#dynamic-changes-to-base-urls:the-del-element}](edits.html#the-del-element) element with a `cite` attribute

:   If the
    [URL](https://url.spec.whatwg.org/#concept-url){#dynamic-changes-to-base-urls:url-4
    x-internal="url"} identified by the `cite` attribute is being shown
    to the user, or if any data derived from that
    [URL](https://url.spec.whatwg.org/#concept-url){#dynamic-changes-to-base-urls:url-5
    x-internal="url"} is affecting the display, then the `cite`
    attribute\'s value should be
    [reparsed](#encoding-parsing-a-url){#dynamic-changes-to-base-urls:encoding-parsing-a-url-3},
    relative to the element\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#dynamic-changes-to-base-urls:node-document-3
    x-internal="node-document"} and the UI updated appropriately.

Otherwise

:   The element is not directly affected.

    For instance, changing the base URL doesn\'t affect the image
    displayed by
    [`img`{#dynamic-changes-to-base-urls:the-img-element}](embedded-content.html#the-img-element)
    elements, although subsequent accesses of the
    [`src`{#dynamic-changes-to-base-urls:dom-img-src}](embedded-content.html#dom-img-src)
    IDL attribute from script will return a new [absolute
    URL](https://url.spec.whatwg.org/#syntax-url-absolute){#dynamic-changes-to-base-urls:absolute-url
    x-internal="absolute-url"} that might no longer correspond to the
    image being shown.

### [2.5]{.secno} Fetching resources[](#fetching-resources){.self-link}

#### [2.5.1]{.secno} Terminology[](#terminology-3){.self-link} {#terminology-3}

A
[response](https://fetch.spec.whatwg.org/#concept-response){#terminology-3:concept-response
x-internal="concept-response"} whose
[type](https://fetch.spec.whatwg.org/#concept-response-type){#terminology-3:concept-response-type
x-internal="concept-response-type"} is \"`basic`\", \"`cors`\", or
\"`default`\" is [CORS-same-origin]{#cors-same-origin .dfn export=""}.
[\[FETCH\]](references.html#refsFETCH)

A
[response](https://fetch.spec.whatwg.org/#concept-response){#terminology-3:concept-response-2
x-internal="concept-response"} whose
[type](https://fetch.spec.whatwg.org/#concept-response-type){#terminology-3:concept-response-type-2
x-internal="concept-response-type"} is \"`opaque`\" or
\"`opaqueredirect`\" is [CORS-cross-origin]{#cors-cross-origin .dfn}.

A
[response](https://fetch.spec.whatwg.org/#concept-response){#terminology-3:concept-response-3
x-internal="concept-response"}\'s [unsafe response]{#unsafe-response
.dfn export=""} is its [internal
response](https://fetch.spec.whatwg.org/#concept-internal-response){#terminology-3:concept-internal-response
x-internal="concept-internal-response"} if it has one, and the
[response](https://fetch.spec.whatwg.org/#concept-response){#terminology-3:concept-response-4
x-internal="concept-response"} itself otherwise.

To [create a potential-CORS request]{#create-a-potential-cors-request
.dfn}, given a `url`{.variable}, `destination`{.variable},
`corsAttributeState`{.variable}, and an optional *same-origin fallback
flag*, run these steps:

1.  Let `mode`{.variable} be \"`no-cors`\" if
    `corsAttributeState`{.variable} is [No
    CORS](#attr-crossorigin-none){#terminology-3:attr-crossorigin-none},
    and \"`cors`\" otherwise.

2.  If *same-origin fallback flag* is set and `mode`{.variable} is
    \"`no-cors`\", set `mode`{.variable} to \"`same-origin`\".

3.  Let `credentialsMode`{.variable} be \"`include`\".

4.  If `corsAttributeState`{.variable} is
    [Anonymous](#attr-crossorigin-anonymous){#terminology-3:attr-crossorigin-anonymous},
    set `credentialsMode`{.variable} to \"`same-origin`\".

5.  Return a new
    [request](https://fetch.spec.whatwg.org/#concept-request){#terminology-3:concept-request
    x-internal="concept-request"} whose
    [URL](https://fetch.spec.whatwg.org/#concept-request-url){#terminology-3:concept-request-url
    x-internal="concept-request-url"} is `url`{.variable},
    [destination](https://fetch.spec.whatwg.org/#concept-request-destination){#terminology-3:concept-request-destination
    x-internal="concept-request-destination"} is
    `destination`{.variable},
    [mode](https://fetch.spec.whatwg.org/#concept-request-mode){#terminology-3:concept-request-mode
    x-internal="concept-request-mode"} is `mode`{.variable},
    [credentials
    mode](https://fetch.spec.whatwg.org/#concept-request-credentials-mode){#terminology-3:concept-request-credentials-mode
    x-internal="concept-request-credentials-mode"} is
    `credentialsMode`{.variable}, and whose [use-URL-credentials
    flag](https://fetch.spec.whatwg.org/#concept-request-use-url-credentials-flag){#terminology-3:use-url-credentials-flag
    x-internal="use-url-credentials-flag"} is set.

#### [2.5.2]{.secno} Determining the type of a resource[](#content-type-sniffing){.self-link} {#content-type-sniffing}

The [Content-Type metadata]{#content-type .dfn} of a resource must be
obtained and interpreted in a manner consistent with the requirements of
MIME Sniffing. [\[MIMESNIFF\]](references.html#refsMIMESNIFF)

The [[computed MIME
type](https://mimesniff.spec.whatwg.org/#computed-mime-type)]{#content-type-sniffing-2
.dfn} of a resource must be found in a manner consistent with the
requirements given in MIME Sniffing.
[\[MIMESNIFF\]](references.html#refsMIMESNIFF)

The [[rules for sniffing images
specifically](https://mimesniff.spec.whatwg.org/#rules-for-sniffing-images-specifically)]{#content-type-sniffing:-image
.dfn}, the [[rules for distinguishing if a resource is text or
binary](https://mimesniff.spec.whatwg.org/#rules-for-text-or-binary)]{#content-type-sniffing:-text-or-binary
.dfn}, and the [[rules for sniffing audio and video
specifically](https://mimesniff.spec.whatwg.org/#rules-for-sniffing-audio-and-video-specifically)]{#content-type-sniffing:-video
.dfn} are also defined in MIME Sniffing. These rules return a [MIME
type](https://mimesniff.spec.whatwg.org/#mime-type){#content-type-sniffing:mime-type
x-internal="mime-type"} as their result.
[\[MIMESNIFF\]](references.html#refsMIMESNIFF)

It is imperative that the rules in MIME Sniffing be followed exactly.
When a user agent uses different heuristics for content type detection
than the server expects, security problems can occur. For more details,
see MIME Sniffing. [\[MIMESNIFF\]](references.html#refsMIMESNIFF)

#### [2.5.3]{.secno} Extracting character encodings from [`meta`{#extracting-character-encodings-from-meta-elements:the-meta-element}](semantics.html#the-meta-element) elements[](#extracting-character-encodings-from-meta-elements){.self-link}

The [algorithm for extracting a character encoding from a `meta`
element]{#algorithm-for-extracting-a-character-encoding-from-a-meta-element
.dfn}, given a string `s`{.variable}, is as follows. It returns either a
character encoding or nothing.

1.  Let `position`{.variable} be a pointer into `s`{.variable},
    initially pointing at the start of the string.

2.  *Loop*: Find the first seven characters in `s`{.variable} after
    `position`{.variable} that are an [ASCII
    case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#extracting-character-encodings-from-meta-elements:ascii-case-insensitive
    x-internal="ascii-case-insensitive"} match for the word
    \"`charset`\". If no such match is found, return nothing.

3.  Skip any [ASCII
    whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#extracting-character-encodings-from-meta-elements:space-characters
    x-internal="space-characters"} that immediately follow the word
    \"`charset`\" (there might not be any).

4.  If the next character is not a U+003D EQUALS SIGN (=), then move
    `position`{.variable} to point just before that next character, and
    jump back to the step labeled *loop*.

5.  Skip any [ASCII
    whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#extracting-character-encodings-from-meta-elements:space-characters-2
    x-internal="space-characters"} that immediately follow the equals
    sign (there might not be any).

6.  Process the next character as follows:

    If it is a U+0022 QUOTATION MARK character (\") and there is a later U+0022 QUOTATION MARK character (\") in `s`{.variable}\
    If it is a U+0027 APOSTROPHE character (\') and there is a later U+0027 APOSTROPHE character (\') in `s`{.variable}
    :   Return the result of [getting an
        encoding](https://encoding.spec.whatwg.org/#concept-encoding-get){#extracting-character-encodings-from-meta-elements:getting-an-encoding
        x-internal="getting-an-encoding"} from the substring that is
        between this character and the next earliest occurrence of this
        character.

    If it is an unmatched U+0022 QUOTATION MARK character (\")\
    If it is an unmatched U+0027 APOSTROPHE character (\')\
    If there is no next character
    :   Return nothing.

    Otherwise
    :   Return the result of [getting an
        encoding](https://encoding.spec.whatwg.org/#concept-encoding-get){#extracting-character-encodings-from-meta-elements:getting-an-encoding-2
        x-internal="getting-an-encoding"} from the substring that
        consists of this character up to but not including the first
        [ASCII
        whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#extracting-character-encodings-from-meta-elements:space-characters-3
        x-internal="space-characters"} or U+003B SEMICOLON character
        (;), or the end of `s`{.variable}, whichever comes first.

This algorithm is distinct from those in the HTTP specifications (for
example, HTTP doesn\'t allow the use of single quotes and requires
supporting a backslash-escape mechanism that is not supported by this
algorithm). While the algorithm is used in contexts that, historically,
were related to HTTP, the syntax as supported by implementations
diverged some time ago. [\[HTTP\]](references.html#refsHTTP)

#### [2.5.4]{.secno} CORS settings attributes[](#cors-settings-attributes){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Attributes/crossorigin](https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/crossorigin "The crossorigin attribute, valid on the <audio>, <img>, <link>, <script>, and <video> elements, provides support for CORS, defining how the element handles cross-origin requests, thereby enabling the configuration of the CORS requests for the element's fetched data. Depending on the element, the attribute can be a CORS settings attribute.")

Support in all current engines.

::: support
[Firefox8+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome13+]{.chrome
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

A [CORS settings attribute]{#cors-settings-attribute .dfn} is an
[enumerated
attribute](common-microsyntaxes.html#enumerated-attribute){#cors-settings-attributes:enumerated-attribute}
with the following keywords and states:

Keyword

State

Brief description

[`anonymous`]{#attr-crossorigin-anonymous-keyword .dfn
dfn-for="audio/crossorigin,video/crossorigin,img/crossorigin,link/crossorigin,script/crossorigin"
dfn-type="attr-value"}

[Anonymous]{#attr-crossorigin-anonymous .dfn}

[Requests](https://fetch.spec.whatwg.org/#concept-request){#cors-settings-attributes:concept-request
x-internal="concept-request"} for the element will have their
[mode](https://fetch.spec.whatwg.org/#concept-request-mode){#cors-settings-attributes:concept-request-mode
x-internal="concept-request-mode"} set to \"`cors`\" and their
[credentials
mode](https://fetch.spec.whatwg.org/#concept-request-credentials-mode){#cors-settings-attributes:concept-request-credentials-mode
x-internal="concept-request-credentials-mode"} set to \"`same-origin`\".

(the empty string)

[`use-credentials`]{#attr-crossorigin-use-credentials-keyword .dfn
dfn-for="audio/crossorigin,video/crossorigin,img/crossorigin,link/crossorigin,script/crossorigin"
dfn-type="attr-value"}

[Use Credentials]{#attr-crossorigin-use-credentials .dfn}

[Requests](https://fetch.spec.whatwg.org/#concept-request){#cors-settings-attributes:concept-request-2
x-internal="concept-request"} for the element will have their
[mode](https://fetch.spec.whatwg.org/#concept-request-mode){#cors-settings-attributes:concept-request-mode-2
x-internal="concept-request-mode"} set to \"`cors`\" and their
[credentials
mode](https://fetch.spec.whatwg.org/#concept-request-credentials-mode){#cors-settings-attributes:concept-request-credentials-mode-2
x-internal="concept-request-credentials-mode"} set to \"`include`\".

The attribute\'s *[missing value
default](common-microsyntaxes.html#missing-value-default)* is the [No
CORS]{#attr-crossorigin-none .dfn} state, and its *[invalid value
default](common-microsyntaxes.html#invalid-value-default)* is the
[Anonymous](#attr-crossorigin-anonymous){#cors-settings-attributes:attr-crossorigin-anonymous}
state.

The majority of fetches governed by [CORS settings
attributes](#cors-settings-attribute){#cors-settings-attributes:cors-settings-attribute}
will be done via the [create a potential-CORS
request](#create-a-potential-cors-request){#cors-settings-attributes:create-a-potential-cors-request}
algorithm.

For more modern features, where the request\'s
[mode](https://fetch.spec.whatwg.org/#concept-request-mode){#cors-settings-attributes:concept-request-mode-3
x-internal="concept-request-mode"} is always \"`cors`\", certain [CORS
settings
attributes](#cors-settings-attribute){#cors-settings-attributes:cors-settings-attribute-2}
have been repurposed to have a slightly different meaning, wherein they
only impact the
[request](https://fetch.spec.whatwg.org/#concept-request){#cors-settings-attributes:concept-request-3
x-internal="concept-request"}\'s [credentials
mode](https://fetch.spec.whatwg.org/#concept-request-credentials-mode){#cors-settings-attributes:concept-request-credentials-mode-3
x-internal="concept-request-credentials-mode"}. To perform this
translation, we define the [CORS settings attribute credentials
mode]{#cors-settings-attribute-credentials-mode .dfn export=""} for a
given [CORS settings
attribute](#cors-settings-attribute){#cors-settings-attributes:cors-settings-attribute-3}
to be determined by switching on the attribute\'s state:

[No CORS](#attr-crossorigin-none){#cors-settings-attributes:attr-crossorigin-none}\
[Anonymous](#attr-crossorigin-anonymous){#cors-settings-attributes:attr-crossorigin-anonymous-2}
:   \"`same-origin`\"

[Use Credentials](#attr-crossorigin-none){#cors-settings-attributes:attr-crossorigin-none-2}
:   \"`include`\"

#### [2.5.5]{.secno} Referrer policy attributes[](#referrer-policy-attributes){.self-link}

A [referrer policy attribute]{#referrer-policy-attribute .dfn export=""}
is an [enumerated
attribute](common-microsyntaxes.html#enumerated-attribute){#referrer-policy-attributes:enumerated-attribute}.
Each [referrer
policy](https://w3c.github.io/webappsec-referrer-policy/#referrer-policy){#referrer-policy-attributes:referrer-policy
x-internal="referrer-policy"}, including the empty string, is a keyword
for this attribute, mapping to a state of the same name.

The attribute\'s *[missing value
default](common-microsyntaxes.html#missing-value-default)* and *[invalid
value default](common-microsyntaxes.html#invalid-value-default)* are
both the empty string state.

The impact of these states on the processing model of various
[fetches](https://fetch.spec.whatwg.org/#concept-fetch){#referrer-policy-attributes:concept-fetch
x-internal="concept-fetch"} is defined in more detail throughout this
specification, in Fetch, and in Referrer Policy.
[\[FETCH\]](references.html#refsFETCH)
[\[REFERRERPOLICY\]](references.html#refsREFERRERPOLICY)

::: note
Several signals can contribute to which processing model is used for a
given
[fetch](https://fetch.spec.whatwg.org/#concept-fetch){#referrer-policy-attributes:concept-fetch-2
x-internal="concept-fetch"}; a [referrer policy
attribute](#referrer-policy-attribute){#referrer-policy-attributes:referrer-policy-attribute}
is only one of them. In general, the order in which these signals are
processed are:

1.  First, the presence of a
    [`noreferrer`{#referrer-policy-attributes:link-type-noreferrer}](links.html#link-type-noreferrer)
    link type;

2.  Then, the value of a [referrer policy
    attribute](#referrer-policy-attribute){#referrer-policy-attributes:referrer-policy-attribute-2};

3.  Then, the presence of any
    [`meta`{#referrer-policy-attributes:the-meta-element}](semantics.html#the-meta-element)
    element with
    [`name`{#referrer-policy-attributes:attr-meta-name}](semantics.html#attr-meta-name)
    attribute set to
    [`referrer`{#referrer-policy-attributes:meta-referrer}](semantics.html#meta-referrer).

4.  Finally, the
    \`[`Referrer-Policy`{#referrer-policy-attributes:http-referrer-policy}](https://w3c.github.io/webappsec-referrer-policy/#referrer-policy-header-dfn){x-internal="http-referrer-policy"}\`
    HTTP header.
:::

#### [2.5.6]{.secno} Nonce attributes[](#nonce-attributes){.self-link}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Global_attributes/nonce](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/nonce "The nonce global attribute is a content attribute defining a cryptographic nonce ("number used once") which can be used by Content Security Policy to determine whether or not a given fetch will be allowed to proceed for a given element.")

Support in all current engines.

::: support
[Firefox31+]{.firefox .yes}[SafariYes]{.safari .yes}[ChromeYes]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[EdgeYes]{.edge_blink .yes}

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

A [`nonce`]{#attr-nonce .dfn dfn-for="htmlsvg-global"
dfn-type="element-attr"} content attribute represents a cryptographic
nonce (\"number used once\") which can be used by Content Security
Policy to determine whether or not a given fetch will be allowed to
proceed. The value is text. [\[CSP\]](references.html#refsCSP)

Elements that have a
[`nonce`{#nonce-attributes:attr-nonce}](#attr-nonce) content attribute
ensure that the cryptographic nonce is only exposed to script (and not
to side-channels like CSS attribute selectors) by taking the value from
the content attribute, moving it into an internal slot named
[\[\[CryptographicNonce\]\]]{#cryptographicnonce .dfn
dfn-for="HTMLOrSVGElement" dfn-type="attribute"}, exposing it to script
via the
[`HTMLOrSVGElement`{#nonce-attributes:htmlorsvgelement}](dom.html#htmlorsvgelement)
interface mixin, and setting the content attribute to the empty string.
Unless otherwise specified, the slot\'s value is the empty string.

`element`{.variable}`.`[`nonce`](#dom-noncedelement-nonce){#dom-htmlorsvgelement-nonce-dev}
:   Returns the value set for `element`{.variable}\'s cryptographic
    nonce. If the setter was not used, this will be the value originally
    found in the [`nonce`{#nonce-attributes:attr-nonce-2}](#attr-nonce)
    content attribute.

`element`{.variable}`.`[`nonce`](#dom-noncedelement-nonce){#nonce-attributes:dom-noncedelement-nonce}` = ``value`{.variable}

:   Updates `element`{.variable}\'s cryptographic nonce value.

::::: {.mdn-anno .wrapped .before}
MDN

:::: feature
[HTMLElement/nonce](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/nonce "The nonce property of the HTMLElement interface returns the cryptographic number used once that is used by Content Security Policy to determine whether a given fetch will be allowed to proceed.")

::: support
[Firefox75+]{.firefox .yes}[Safari[🔰
10+]{title="Partial implementation."}]{.safari .yes}[Chrome61+]{.chrome
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

The [`nonce`]{#dom-noncedelement-nonce .dfn dfn-for="HTMLOrSVGElement"
dfn-type="attribute"} IDL attribute must, on getting, return the value
of this element\'s
[\[\[CryptographicNonce\]\]](#cryptographicnonce){#nonce-attributes:cryptographicnonce};
and on setting, set this element\'s
[\[\[CryptographicNonce\]\]](#cryptographicnonce){#nonce-attributes:cryptographicnonce-2}
to the given value.

Note how the setter for the
[`nonce`{#nonce-attributes:dom-noncedelement-nonce-2}](#dom-noncedelement-nonce)
IDL attribute does not update the corresponding content attribute. This,
as well as the below setting of the
[`nonce`{#nonce-attributes:attr-nonce-3}](#attr-nonce) content attribute
to the empty string when an element [becomes browsing-context
connected](infrastructure.html#becomes-browsing-context-connected){#nonce-attributes:becomes-browsing-context-connected},
is meant to prevent exfiltration of the nonce value through mechanisms
that can easily read content attributes, such as selectors. Learn more
in [issue #2369](https://github.com/whatwg/html/issues/2369), where this
behavior was introduced.

The following [attribute change
steps](https://dom.spec.whatwg.org/#concept-element-attributes-change-ext){#nonce-attributes:concept-element-attributes-change-ext
x-internal="concept-element-attributes-change-ext"} are used for the
[`nonce`{#nonce-attributes:attr-nonce-4}](#attr-nonce) content
attribute:

1.  If `element`{.variable} does not
    [include](https://webidl.spec.whatwg.org/#include){#nonce-attributes:include
    x-internal="include"}
    [`HTMLOrSVGElement`{#nonce-attributes:htmlorsvgelement-2}](dom.html#htmlorsvgelement),
    then return.

2.  If `localName`{.variable} is not
    [`nonce`{#nonce-attributes:attr-nonce-5}](#attr-nonce) or
    `namespace`{.variable} is not null, then return.

3.  If `value`{.variable} is null, then set `element`{.variable}\'s
    [\[\[CryptographicNonce\]\]](#cryptographicnonce){#nonce-attributes:cryptographicnonce-3}
    to the empty string.

4.  Otherwise, set `element`{.variable}\'s
    [\[\[CryptographicNonce\]\]](#cryptographicnonce){#nonce-attributes:cryptographicnonce-4}
    to `value`{.variable}.

Whenever an element
[including](https://webidl.spec.whatwg.org/#include){#nonce-attributes:include-2
x-internal="include"}
[`HTMLOrSVGElement`{#nonce-attributes:htmlorsvgelement-3}](dom.html#htmlorsvgelement)
[becomes browsing-context
connected](infrastructure.html#becomes-browsing-context-connected){#nonce-attributes:becomes-browsing-context-connected-2},
the user agent must execute the following steps on the
`element`{.variable}:

1.  Let `CSP list`{.variable} be `element`{.variable}\'s
    [shadow-including
    root](https://dom.spec.whatwg.org/#concept-shadow-including-root){#nonce-attributes:shadow-including-root
    x-internal="shadow-including-root"}\'s [policy
    container](dom.html#concept-document-policy-container){#nonce-attributes:concept-document-policy-container}\'s
    [CSP
    list](browsers.html#policy-container-csp-list){#nonce-attributes:policy-container-csp-list}.

2.  If `CSP list`{.variable} [contains a header-delivered Content
    Security
    Policy](https://w3c.github.io/webappsec-csp/#contains-a-header-delivered-content-security-policy){#nonce-attributes:contains-a-header-delivered-content-security-policy
    x-internal="contains-a-header-delivered-content-security-policy"},
    and `element`{.variable} has a
    [`nonce`{#nonce-attributes:attr-nonce-6}](#attr-nonce) content
    attribute whose value is not the empty string, then:

    1.  Let `nonce`{.variable} be `element`{.variable}\'s
        [\[\[CryptographicNonce\]\]](#cryptographicnonce){#nonce-attributes:cryptographicnonce-5}.

    2.  [Set an attribute
        value](https://dom.spec.whatwg.org/#concept-element-attributes-set-value){#nonce-attributes:concept-element-attributes-set-value
        x-internal="concept-element-attributes-set-value"} for
        `element`{.variable} using
        \"[`nonce`{#nonce-attributes:attr-nonce-7}](#attr-nonce)\" and
        the empty string.

    3.  Set `element`{.variable}\'s
        [\[\[CryptographicNonce\]\]](#cryptographicnonce){#nonce-attributes:cryptographicnonce-6}
        to `nonce`{.variable}.

    If `element`{.variable}\'s
    [\[\[CryptographicNonce\]\]](#cryptographicnonce){#nonce-attributes:cryptographicnonce-7}
    were not restored it would be the empty string at this point.

The [cloning
steps](https://dom.spec.whatwg.org/#concept-node-clone-ext){#nonce-attributes:concept-node-clone-ext
x-internal="concept-node-clone-ext"} for elements that
[include](https://webidl.spec.whatwg.org/#include){#nonce-attributes:include-3
x-internal="include"}
[`HTMLOrSVGElement`{#nonce-attributes:htmlorsvgelement-4}](dom.html#htmlorsvgelement)
given `node`{.variable}, `copy`{.variable}, and `subtree`{.variable} are
to set `copy`{.variable}\'s
[\[\[CryptographicNonce\]\]](#cryptographicnonce){#nonce-attributes:cryptographicnonce-8}
to `node`{.variable}\'s
[\[\[CryptographicNonce\]\]](#cryptographicnonce){#nonce-attributes:cryptographicnonce-9}.

#### [2.5.7]{.secno} Lazy loading attributes[](#lazy-loading-attributes){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Lazy_loading](https://developer.mozilla.org/en-US/docs/Web/Performance/Lazy_loading "Lazy loading is a strategy to identify resources as non-blocking (non-critical) and load these only when needed. It's a way to shorten the length of the critical rendering path, which translates into reduced page load times.")

Support in all current engines.

::: support
[Firefox75+]{.firefox .yes}[Safari15.4+]{.safari
.yes}[Chrome77+]{.chrome .yes}

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

A [lazy loading attribute]{#lazy-loading-attribute .dfn} is an
[enumerated
attribute](common-microsyntaxes.html#enumerated-attribute){#lazy-loading-attributes:enumerated-attribute}
with the following keywords and states:

Keyword

State

Brief description

[`lazy`]{#attr-loading-lazy .dfn dfn-for="img/loading,iframe/loading"
dfn-type="attr-value"}

[Lazy]{#attr-loading-lazy-state .dfn}

Used to defer fetching a resource until some conditions are met.

[`eager`]{#attr-loading-eager .dfn dfn-for="img/loading,iframe/loading"
dfn-type="attr-value"}

[Eager]{#attr-loading-eager-state .dfn}

Used to fetch a resource immediately; the default state.

The attribute directs the user agent to fetch a resource immediately or
to defer fetching until some conditions associated with the element are
met, according to the attribute\'s current state.

The attribute\'s *[missing value
default](common-microsyntaxes.html#missing-value-default)* and *[invalid
value default](common-microsyntaxes.html#invalid-value-default)* are
both the
[Eager](#attr-loading-eager-state){#lazy-loading-attributes:attr-loading-eager-state}
state.

------------------------------------------------------------------------

The [will lazy load element steps]{#will-lazy-load-element-steps .dfn},
given an element `element`{.variable}, are as follows:

1.  If [scripting is
    disabled](webappapis.html#concept-n-noscript){#lazy-loading-attributes:concept-n-noscript}
    for `element`{.variable}, then return false.

    This is an anti-tracking measure, because if a user agent supported
    lazy loading when scripting is disabled, it would still be possible
    for a site to track a user\'s approximate scroll position throughout
    a session, by strategically placing images in a page\'s markup such
    that a server can track how many images are requested and when.

2.  If `element`{.variable}\'s [lazy loading
    attribute](#lazy-loading-attribute){#lazy-loading-attributes:lazy-loading-attribute}
    is in the
    [Lazy](#attr-loading-lazy-state){#lazy-loading-attributes:attr-loading-lazy-state}
    state, then return true.

3.  Return false.

Each
[`img`{#lazy-loading-attributes:the-img-element}](embedded-content.html#the-img-element)
and
[`iframe`{#lazy-loading-attributes:the-iframe-element}](iframe-embed-object.html#the-iframe-element)
element has associated [lazy load resumption
steps]{#lazy-load-resumption-steps .dfn}, initially null.

For
[`img`{#lazy-loading-attributes:the-img-element-2}](embedded-content.html#the-img-element)
and
[`iframe`{#lazy-loading-attributes:the-iframe-element-2}](iframe-embed-object.html#the-iframe-element)
elements that [will lazy
load](#will-lazy-load-element-steps){#lazy-loading-attributes:will-lazy-load-element-steps},
these steps are run from the [lazy load intersection
observer](#lazy-load-intersection-observer){#lazy-loading-attributes:lazy-load-intersection-observer}\'s
callback or when their [lazy loading
attribute](#lazy-loading-attribute){#lazy-loading-attributes:lazy-loading-attribute-2}
is set to the
[Eager](#attr-loading-eager-state){#lazy-loading-attributes:attr-loading-eager-state-2}
state. This causes the element to continue loading.

Each [`Document`{#lazy-loading-attributes:document}](dom.html#document)
has a [lazy load intersection observer]{#lazy-load-intersection-observer
.dfn}, initially set to null but can be set to an
[`IntersectionObserver`{#lazy-loading-attributes:intersectionobserver}](https://w3c.github.io/IntersectionObserver/#intersectionobserver){x-internal="intersectionobserver"}
instance.

To [start intersection-observing a lazy loading
element]{#start-intersection-observing-a-lazy-loading-element .dfn}
`element`{.variable}, run these steps:

1.  Let `doc`{.variable} be `element`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#lazy-loading-attributes:node-document
    x-internal="node-document"}.

2.  If `doc`{.variable}\'s [lazy load intersection
    observer](#lazy-load-intersection-observer){#lazy-loading-attributes:lazy-load-intersection-observer-2}
    is null, set it to a new
    [`IntersectionObserver`{#lazy-loading-attributes:intersectionobserver-2}](https://w3c.github.io/IntersectionObserver/#intersectionobserver){x-internal="intersectionobserver"}
    instance, initialized as follows:

    The intention is to use the original value of the
    [`IntersectionObserver`{#lazy-loading-attributes:intersectionobserver-3}](https://w3c.github.io/IntersectionObserver/#intersectionobserver){x-internal="intersectionobserver"}
    constructor. However, we\'re forced to use the JavaScript-exposed
    constructor in this specification, until Intersection Observer
    exposes low-level hooks for use in specifications. See bug
    [w3c/IntersectionObserver#464](https://github.com/w3c/IntersectionObserver/issues/464)
    which tracks this.
    [\[INTERSECTIONOBSERVER\]](references.html#refsINTERSECTIONOBSERVER)

    - The `callback`{.variable} is these steps, with arguments
      `entries`{.variable} and `observer`{.variable}:

      1.  For each `entry`{.variable} in `entries`{.variable} [using a
          method of iteration which does not trigger
          developer-modifiable array accessors or iteration
          hooks]{.XXX}:

          1.  Let `resumptionSteps`{.variable} be null.

          2.  If
              `entry`{.variable}.[`isIntersecting`{#lazy-loading-attributes:dom-intersectionobserverentry-isintersecting}](https://w3c.github.io/IntersectionObserver/#dom-intersectionobserverentry-isintersecting){x-internal="dom-intersectionobserverentry-isintersecting"}
              is true, then set `resumptionSteps`{.variable} to
              `entry`{.variable}.[`target`{#lazy-loading-attributes:dom-intersectionobserverentry-target}](https://w3c.github.io/IntersectionObserver/#dom-intersectionobserverentry-target){x-internal="dom-intersectionobserverentry-target"}\'s
              [lazy load resumption
              steps](#lazy-load-resumption-steps){#lazy-loading-attributes:lazy-load-resumption-steps}.

          3.  If `resumptionSteps`{.variable} is null, then return.

          4.  [Stop intersection-observing a lazy loading
              element](#stop-intersection-observing-a-lazy-loading-element){#lazy-loading-attributes:stop-intersection-observing-a-lazy-loading-element}
              for
              `entry`{.variable}.[`target`{#lazy-loading-attributes:dom-intersectionobserverentry-target-2}](https://w3c.github.io/IntersectionObserver/#dom-intersectionobserverentry-target){x-internal="dom-intersectionobserverentry-target"}.

          5.  Set
              `entry`{.variable}.[`target`{#lazy-loading-attributes:dom-intersectionobserverentry-target-3}](https://w3c.github.io/IntersectionObserver/#dom-intersectionobserverentry-target){x-internal="dom-intersectionobserverentry-target"}\'s
              [lazy load resumption
              steps](#lazy-load-resumption-steps){#lazy-loading-attributes:lazy-load-resumption-steps-2}
              to null.

          6.  Invoke `resumptionSteps`{.variable}.

          The intention is to use the original value of the
          [`isIntersecting`{#lazy-loading-attributes:dom-intersectionobserverentry-isintersecting-2}](https://w3c.github.io/IntersectionObserver/#dom-intersectionobserverentry-isintersecting){x-internal="dom-intersectionobserverentry-isintersecting"}
          and
          [`target`{#lazy-loading-attributes:dom-intersectionobserverentry-target-4}](https://w3c.github.io/IntersectionObserver/#dom-intersectionobserverentry-target){x-internal="dom-intersectionobserverentry-target"}
          getters. See
          [w3c/IntersectionObserver#464](https://github.com/w3c/IntersectionObserver/issues/464).
          [\[INTERSECTIONOBSERVER\]](references.html#refsINTERSECTIONOBSERVER)

    - The `options`{.variable} is an
      [`IntersectionObserverInit`{#lazy-loading-attributes:intersectionobserverinit}](https://w3c.github.io/IntersectionObserver/#dictdef-intersectionobserverinit){x-internal="intersectionobserverinit"}
      dictionary with the following dictionary members: «\[
      \"`scrollMargin`\" → [lazy load scroll
      margin](#lazy-load-root-margin){#lazy-loading-attributes:lazy-load-root-margin}
      \]»

      This allows for fetching the image during scrolling, when it does
      not yet --- but is about to --- intersect the viewport.

      The [lazy load scroll
      margin](#lazy-load-root-margin){#lazy-loading-attributes:lazy-load-root-margin-2}
      suggestions imply dynamic changes to the value, but the
      [`IntersectionObserver`{#lazy-loading-attributes:intersectionobserver-4}](https://w3c.github.io/IntersectionObserver/#intersectionobserver){x-internal="intersectionobserver"}
      API does not support changing the scroll margin. See issue
      [w3c/IntersectionObserver#428](https://github.com/w3c/IntersectionObserver/issues/428).

3.  Call `doc`{.variable}\'s [lazy load intersection
    observer](#lazy-load-intersection-observer){#lazy-loading-attributes:lazy-load-intersection-observer-3}\'s
    [`observe`{#lazy-loading-attributes:dom-intersectionobserver-observe}](https://w3c.github.io/IntersectionObserver/#dom-intersectionobserver-observe){x-internal="dom-intersectionobserver-observe"}
    method with `element`{.variable} as the argument.

    The intention is to use the original value of the
    [`observe`{#lazy-loading-attributes:dom-intersectionobserver-observe-2}](https://w3c.github.io/IntersectionObserver/#dom-intersectionobserver-observe){x-internal="dom-intersectionobserver-observe"}
    method. See
    [w3c/IntersectionObserver#464](https://github.com/w3c/IntersectionObserver/issues/464).
    [\[INTERSECTIONOBSERVER\]](references.html#refsINTERSECTIONOBSERVER)

To [stop intersection-observing a lazy loading
element]{#stop-intersection-observing-a-lazy-loading-element .dfn}
`element`{.variable}, run these steps:

1.  Let `doc`{.variable} be `element`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#lazy-loading-attributes:node-document-2
    x-internal="node-document"}.

2.  [Assert](https://infra.spec.whatwg.org/#assert){#lazy-loading-attributes:assert
    x-internal="assert"}: `doc`{.variable}\'s [lazy load intersection
    observer](#lazy-load-intersection-observer){#lazy-loading-attributes:lazy-load-intersection-observer-4}
    is not null.

3.  Call `doc`{.variable}\'s [lazy load intersection
    observer](#lazy-load-intersection-observer){#lazy-loading-attributes:lazy-load-intersection-observer-5}\'s
    [`unobserve`{#lazy-loading-attributes:dom-intersectionobserver-unobserve}](https://w3c.github.io/IntersectionObserver/#dom-intersectionobserver-unobserve){x-internal="dom-intersectionobserver-unobserve"}
    method with `element`{.variable} as the argument.

    The intention is to use the original value of the
    [`unobserve`{#lazy-loading-attributes:dom-intersectionobserver-unobserve-2}](https://w3c.github.io/IntersectionObserver/#dom-intersectionobserver-unobserve){x-internal="dom-intersectionobserver-unobserve"}
    method. See
    [w3c/IntersectionObserver#464](https://github.com/w3c/IntersectionObserver/issues/464).
    [\[INTERSECTIONOBSERVER\]](references.html#refsINTERSECTIONOBSERVER)

[![(This is a tracking
vector.)](https://resources.whatwg.org/tracking-vector.svg){width="46"
crossorigin=""
height="64"}](https://infra.spec.whatwg.org/#tracking-vector "There is a tracking vector here."){#lazy-loading-attributes:tracking-vector
.tracking-vector x-internal="tracking-vector"} The [lazy load scroll
margin]{#lazy-load-root-margin .dfn} is an
[implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#lazy-loading-attributes:implementation-defined
x-internal="implementation-defined"} value, but with the following
suggestions to consider:

- Set a minimum value that most often results in the resources being
  loaded before they intersect the viewport under normal usage patterns
  for the given device.

- The typical scrolling speed: increase the value for devices with
  faster typical scrolling speeds.

- The current scrolling speed or momentum: the UA can attempt to predict
  where the scrolling will likely stop, and adjust the value
  accordingly.

- The network quality: increase the value for slow or high-latency
  connections.

- User preferences can influence the value.

It is important [for
privacy](https://infra.spec.whatwg.org/#tracking-vector){#lazy-loading-attributes:tracking-vector-2
x-internal="tracking-vector"} that the [lazy load scroll
margin](#lazy-load-root-margin){#lazy-loading-attributes:lazy-load-root-margin-3}
not leak additional information. For example, the typical scrolling
speed on the current device could be imprecise so as to not introduce a
new fingerprinting vector.

#### [2.5.8]{.secno} Blocking attributes[](#blocking-attributes){.self-link}

A [blocking attribute]{#blocking-attribute .dfn} explicitly indicates
that certain operations should be blocked on the fetching of an external
resource. The operations that can be blocked are represented by
[possible blocking tokens]{#possible-blocking-token .dfn}, which are
strings listed by the following table:

Possible blocking token

Description

\"[`render`]{#blocking-token-render .dfn}\"

The element is [potentially
render-blocking](#potentially-render-blocking){#blocking-attributes:potentially-render-blocking}.

In the future, there might be more [possible blocking
tokens](#possible-blocking-token){#blocking-attributes:possible-blocking-token}.

A [blocking
attribute](#blocking-attribute){#blocking-attributes:blocking-attribute}
must have a value that is an [unordered set of unique space-separated
tokens](common-microsyntaxes.html#unordered-set-of-unique-space-separated-tokens){#blocking-attributes:unordered-set-of-unique-space-separated-tokens},
each of which are [possible blocking
tokens](#possible-blocking-token){#blocking-attributes:possible-blocking-token-2}.
The [supported
tokens](https://dom.spec.whatwg.org/#concept-supported-tokens){#blocking-attributes:concept-supported-tokens
x-internal="concept-supported-tokens"} of a [blocking
attribute](#blocking-attribute){#blocking-attributes:blocking-attribute-2}
are the [possible blocking
tokens](#possible-blocking-token){#blocking-attributes:possible-blocking-token-3}.
Any element can have at most one [blocking
attribute](#blocking-attribute){#blocking-attributes:blocking-attribute-3}.

The [blocking tokens set]{#blocking-tokens-set .dfn} for an element
`el`{.variable} are the result of the following steps:

1.  Let `value`{.variable} be the value of `el`{.variable}\'s [blocking
    attribute](#blocking-attribute){#blocking-attributes:blocking-attribute-4},
    or the empty string if no such attribute exists.

2.  Set `value`{.variable} to `value`{.variable}, [converted to ASCII
    lowercase](https://infra.spec.whatwg.org/#ascii-lowercase){#blocking-attributes:converted-to-ascii-lowercase
    x-internal="converted-to-ascii-lowercase"}.

3.  Let `rawTokens`{.variable} be the result of [splitting
    `value`{.variable} on ASCII
    whitespace](https://infra.spec.whatwg.org/#split-on-ascii-whitespace){#blocking-attributes:split-a-string-on-spaces
    x-internal="split-a-string-on-spaces"}.

4.  Return a set containing the elements of `rawTokens`{.variable} that
    are [possible blocking
    tokens](#possible-blocking-token){#blocking-attributes:possible-blocking-token-4}.

An element is [potentially render-blocking]{#potentially-render-blocking
.dfn} if its [blocking tokens
set](#blocking-tokens-set){#blocking-attributes:blocking-tokens-set}
contains
\"[`render`{#blocking-attributes:blocking-token-render}](#blocking-token-render)\",
or if it is [implicitly potentially
render-blocking]{#implicitly-potentially-render-blocking .dfn}, which
will be defined at the individual elements. By default, an element is
not [implicitly potentially
render-blocking](#implicitly-potentially-render-blocking){#blocking-attributes:implicitly-potentially-render-blocking}.

#### [2.5.9]{.secno} Fetch priority attributes[](#fetch-priority-attributes){.self-link}

A [fetch priority attribute]{#fetch-priority-attribute .dfn} is an
[enumerated
attribute](common-microsyntaxes.html#enumerated-attribute){#fetch-priority-attributes:enumerated-attribute}
with the following keywords and states:

Keyword

State

Brief description

[`high`]{#attr-fetchpriority-high .dfn
dfn-for="img/fetchpriority,script/fetchpriority,link/fetchpriority"
dfn-type="attr-value"}

[High]{#attr-fetchpriority-high-state .dfn}

Signals a high-priority
[fetch](https://fetch.spec.whatwg.org/#concept-fetch){#fetch-priority-attributes:concept-fetch
x-internal="concept-fetch"} relative to other resources with the same
[destination](https://fetch.spec.whatwg.org/#concept-request-destination){#fetch-priority-attributes:concept-request-destination
x-internal="concept-request-destination"}.

[`low`]{#attr-fetchpriority-low .dfn
dfn-for="img/fetchpriority,script/fetchpriority,link/fetchpriority"
dfn-type="attr-value"}

[Low]{#attr-fetchpriority-low-state .dfn}

Signals a low-priority
[fetch](https://fetch.spec.whatwg.org/#concept-fetch){#fetch-priority-attributes:concept-fetch-2
x-internal="concept-fetch"} relative to other resources with the same
[destination](https://fetch.spec.whatwg.org/#concept-request-destination){#fetch-priority-attributes:concept-request-destination-2
x-internal="concept-request-destination"}.

[`auto`]{#attr-fetchpriority-auto .dfn
dfn-for="img/fetchpriority,script/fetchpriority,link/fetchpriority"
dfn-type="attr-value"}

[Auto]{#attr-fetchpriority-auto-state .dfn}

Signals automatic determination of
[fetch](https://fetch.spec.whatwg.org/#concept-fetch){#fetch-priority-attributes:concept-fetch-3
x-internal="concept-fetch"} priority relative to other resources with
the same
[destination](https://fetch.spec.whatwg.org/#concept-request-destination){#fetch-priority-attributes:concept-request-destination-3
x-internal="concept-request-destination"}.

The attribute\'s *[missing value
default](common-microsyntaxes.html#missing-value-default)* and *[invalid
value default](common-microsyntaxes.html#invalid-value-default)* are
both the
[Auto](#attr-fetchpriority-auto-state){#fetch-priority-attributes:attr-fetchpriority-auto-state}
state.

[← 2.3 Common microsyntaxes](common-microsyntaxes.html) --- [Table of
Contents](./) --- [2.6 Common DOM interfaces
→](common-dom-interfaces.html)
