HTML Standard

::: {#head .head .with-buttons}
[![WHATWG](https://resources.whatwg.org/logo.svg){.darkmode-aware
width="100" crossorigin="" height="100"}](https://whatwg.org/){.logo}

# HTML {#html .allcaps}

Living Standard --- Last Updated [2 June 2025]{.pubdate}
:::

[← 4.5 Text-level semantics](text-level-semantics.html) --- [Table of
Contents](./) --- [4.7 Edits →](edits.html)

1.  ::: {#toc-semantics}
    1.  [[4.6]{.secno} Links](links.html#links)
        1.  [[4.6.1]{.secno} Introduction](links.html#introduction-2)
        2.  [[4.6.2]{.secno} Links created by `a` and `area`
            elements](links.html#links-created-by-a-and-area-elements)
        3.  [[4.6.3]{.secno} API for `a` and `area`
            elements](links.html#api-for-a-and-area-elements)
        4.  [[4.6.4]{.secno} Following
            hyperlinks](links.html#following-hyperlinks)
        5.  [[4.6.5]{.secno} Downloading
            resources](links.html#downloading-resources)
        6.  [[4.6.6]{.secno} Hyperlink
            auditing](links.html#hyperlink-auditing)
            1.  [[4.6.6.1]{.secno} The \``Ping-From`\` and \``Ping-To`\`
                headers](links.html#the-ping-headers)
        7.  [[4.6.7]{.secno} Link types](links.html#linkTypes)
            1.  [[4.6.7.1]{.secno} Link type
                \"`alternate`\"](links.html#rel-alternate)
            2.  [[4.6.7.2]{.secno} Link type
                \"`author`\"](links.html#link-type-author)
            3.  [[4.6.7.3]{.secno} Link type
                \"`bookmark`\"](links.html#link-type-bookmark)
            4.  [[4.6.7.4]{.secno} Link type
                \"`canonical`\"](links.html#link-type-canonical)
            5.  [[4.6.7.5]{.secno} Link type
                \"`dns-prefetch`\"](links.html#link-type-dns-prefetch)
            6.  [[4.6.7.6]{.secno} Link type
                \"`expect`\"](links.html#link-type-expect)
            7.  [[4.6.7.7]{.secno} Link type
                \"`external`\"](links.html#link-type-external)
            8.  [[4.6.7.8]{.secno} Link type
                \"`help`\"](links.html#link-type-help)
            9.  [[4.6.7.9]{.secno} Link type
                \"`icon`\"](links.html#rel-icon)
            10. [[4.6.7.10]{.secno} Link type
                \"`license`\"](links.html#link-type-license)
            11. [[4.6.7.11]{.secno} Link type
                \"`manifest`\"](links.html#link-type-manifest)
            12. [[4.6.7.12]{.secno} Link type
                \"`modulepreload`\"](links.html#link-type-modulepreload)
            13. [[4.6.7.13]{.secno} Link type
                \"`nofollow`\"](links.html#link-type-nofollow)
            14. [[4.6.7.14]{.secno} Link type
                \"`noopener`\"](links.html#link-type-noopener)
            15. [[4.6.7.15]{.secno} Link type
                \"`noreferrer`\"](links.html#link-type-noreferrer)
            16. [[4.6.7.16]{.secno} Link type
                \"`opener`\"](links.html#link-type-opener)
            17. [[4.6.7.17]{.secno} Link type
                \"`pingback`\"](links.html#link-type-pingback)
            18. [[4.6.7.18]{.secno} Link type
                \"`preconnect`\"](links.html#link-type-preconnect)
            19. [[4.6.7.19]{.secno} Link type
                \"`prefetch`\"](links.html#link-type-prefetch)
            20. [[4.6.7.20]{.secno} Link type
                \"`preload`\"](links.html#link-type-preload)
            21. [[4.6.7.21]{.secno} Link type
                \"`privacy-policy`\"](links.html#link-type-privacy-policy)
            22. [[4.6.7.22]{.secno} Link type
                \"`search`\"](links.html#link-type-search)
            23. [[4.6.7.23]{.secno} Link type
                \"`stylesheet`\"](links.html#link-type-stylesheet)
            24. [[4.6.7.24]{.secno} Link type
                \"`tag`\"](links.html#link-type-tag)
            25. [[4.6.7.25]{.secno} Link Type
                \"`terms-of-service`\"](links.html#link-type-terms-of-service)
            26. [[4.6.7.26]{.secno} Sequential link
                types](links.html#sequential-link-types)
                1.  [[4.6.7.26.1]{.secno} Link type
                    \"`next`\"](links.html#link-type-next)
                2.  [[4.6.7.26.2]{.secno} Link type
                    \"`prev`\"](links.html#link-type-prev)
            27. [[4.6.7.27]{.secno} Other link
                types](links.html#other-link-types)
    :::

### [4.6]{.secno} Links[](#links){.self-link}

#### [4.6.1]{.secno} Introduction[](#introduction-2){.self-link} {#introduction-2}

Links are a conceptual construct, created by
[`a`{#introduction-2:the-a-element}](text-level-semantics.html#the-a-element),
[`area`{#introduction-2:the-area-element}](image-maps.html#the-area-element),
[`form`{#introduction-2:the-form-element}](forms.html#the-form-element),
and
[`link`{#introduction-2:the-link-element}](semantics.html#the-link-element)
elements, that
[represent](dom.html#represents){#introduction-2:represents} a
connection between two resources, one of which is the current
[`Document`{#introduction-2:document}](dom.html#document). There are
three kinds of links in HTML:

[Links to external resources]{#external-resource-link .dfn lt="external resource link" export=""}
:   These are links to resources that are to be used to augment the
    current document, generally automatically processed by the user
    agent. All [external resource
    links](#external-resource-link){#introduction-2:external-resource-link}
    have a [fetch and process the linked
    resource](semantics.html#fetch-and-process-the-linked-resource){#introduction-2:fetch-and-process-the-linked-resource}
    algorithm which describes how the resource is obtained.

[Hyperlinks]{#hyperlink .dfn lt="hyperlink" export=""}
:   These are links to other resources that are generally exposed to the
    user by the user agent so that the user can cause the user agent to
    [navigate](browsing-the-web.html#navigate){#introduction-2:navigate}
    to those resources, e.g. to visit them in a browser or download
    them.

[Internal resource links]{#internal-resource-link .dfn export=""}

:   These are links to resources within the current document, used to
    give those resources special meaning or behavior.

For
[`link`{#introduction-2:the-link-element-2}](semantics.html#the-link-element)
elements with an
[`href`{#introduction-2:attr-link-href}](semantics.html#attr-link-href)
attribute and a
[`rel`{#introduction-2:attr-link-rel}](semantics.html#attr-link-rel)
attribute, links must be created for the keywords of the
[`rel`{#introduction-2:attr-link-rel-2}](semantics.html#attr-link-rel)
attribute, as defined for those keywords in the [link types](#linkTypes)
section.

Similarly, for
[`a`{#introduction-2:the-a-element-2}](text-level-semantics.html#the-a-element)
and
[`area`{#introduction-2:the-area-element-2}](image-maps.html#the-area-element)
elements with an
[`href`{#introduction-2:attr-hyperlink-href}](#attr-hyperlink-href)
attribute and a
[`rel`{#introduction-2:attr-hyperlink-rel}](#attr-hyperlink-rel)
attribute, links must be created for the keywords of the
[`rel`{#introduction-2:attr-hyperlink-rel-2}](#attr-hyperlink-rel)
attribute as defined for those keywords in the [link types](#linkTypes)
section. Unlike
[`link`{#introduction-2:the-link-element-3}](semantics.html#the-link-element)
elements, however,
[`a`{#introduction-2:the-a-element-3}](text-level-semantics.html#the-a-element)
and
[`area`{#introduction-2:the-area-element-3}](image-maps.html#the-area-element)
elements with an
[`href`{#introduction-2:attr-hyperlink-href-2}](#attr-hyperlink-href)
attribute that either do not have a
[`rel`{#introduction-2:attr-hyperlink-rel-3}](#attr-hyperlink-rel)
attribute, or whose
[`rel`{#introduction-2:attr-hyperlink-rel-4}](#attr-hyperlink-rel)
attribute has no keywords that are defined as specifying
[hyperlinks](#hyperlink){#introduction-2:hyperlink}, must also create a
[hyperlink](#hyperlink){#introduction-2:hyperlink-2}. This implied
hyperlink has no special meaning (it has no [link type](#linkTypes))
beyond linking the element\'s [node
document](https://dom.spec.whatwg.org/#concept-node-document){#introduction-2:node-document
x-internal="node-document"} to the resource given by the element\'s
[`href`{#introduction-2:attr-hyperlink-href-3}](#attr-hyperlink-href)
attribute.

Similarly, for
[`form`{#introduction-2:the-form-element-2}](forms.html#the-form-element)
elements with a
[`rel`{#introduction-2:attr-form-rel}](forms.html#attr-form-rel)
attribute, links must be created for the keywords of the
[`rel`{#introduction-2:attr-form-rel-2}](forms.html#attr-form-rel)
attribute as defined for those keywords in the [link types](#linkTypes)
section.
[`form`{#introduction-2:the-form-element-3}](forms.html#the-form-element)
elements that do not have a
[`rel`{#introduction-2:attr-form-rel-3}](forms.html#attr-form-rel)
attribute, or whose
[`rel`{#introduction-2:attr-form-rel-4}](forms.html#attr-form-rel)
attribute has no keywords that are defined as specifying
[hyperlinks](#hyperlink){#introduction-2:hyperlink-3}, must also create
a [hyperlink](#hyperlink){#introduction-2:hyperlink-4}.

A [hyperlink](#hyperlink){#introduction-2:hyperlink-5} can have one or
more [hyperlink annotations]{#hyperlink-annotation .dfn} that modify the
processing semantics of that hyperlink.

#### [4.6.2]{.secno} Links created by [`a`{#links-created-by-a-and-area-elements:the-a-element}](text-level-semantics.html#the-a-element) and [`area`{#links-created-by-a-and-area-elements:the-area-element}](image-maps.html#the-area-element) elements[](#links-created-by-a-and-area-elements){.self-link}

The [`href`]{#attr-hyperlink-href .dfn dfn-for="a,area"
dfn-type="element-attr"} attribute on
[`a`{#links-created-by-a-and-area-elements:the-a-element-2}](text-level-semantics.html#the-a-element)
and
[`area`{#links-created-by-a-and-area-elements:the-area-element-2}](image-maps.html#the-area-element)
elements must have a value that is a [valid URL potentially surrounded
by
spaces](urls-and-fetching.html#valid-url-potentially-surrounded-by-spaces){#links-created-by-a-and-area-elements:valid-url-potentially-surrounded-by-spaces}.

The
[`href`{#links-created-by-a-and-area-elements:attr-hyperlink-href}](#attr-hyperlink-href)
attribute on
[`a`{#links-created-by-a-and-area-elements:the-a-element-3}](text-level-semantics.html#the-a-element)
and
[`area`{#links-created-by-a-and-area-elements:the-area-element-3}](image-maps.html#the-area-element)
elements is not required; when those elements do not have
[`href`{#links-created-by-a-and-area-elements:attr-hyperlink-href-2}](#attr-hyperlink-href)
attributes they do not create hyperlinks.

The [`target`]{#attr-hyperlink-target .dfn dfn-for="a,area"
dfn-type="element-attr"} attribute, if present, must be a [valid
navigable target name or
keyword](document-sequences.html#valid-navigable-target-name-or-keyword){#links-created-by-a-and-area-elements:valid-navigable-target-name-or-keyword}.
It gives the name of the
[navigable](document-sequences.html#navigable){#links-created-by-a-and-area-elements:navigable}
that will be used. User agents use this name when [following
hyperlinks](#following-hyperlinks-2){#links-created-by-a-and-area-elements:following-hyperlinks-2}.

The [`download`]{#attr-hyperlink-download .dfn dfn-for="a,area"
dfn-type="element-attr"} attribute, if present, indicates that the
author intends the hyperlink to be used for [downloading a
resource](#downloading-hyperlinks){#links-created-by-a-and-area-elements:downloading-hyperlinks}.
The attribute may have a value; the value, if any, specifies the default
filename that the author recommends for use in labeling the resource in
a local file system. There are no restrictions on allowed values, but
authors are cautioned that most file systems have limitations with
regard to what punctuation is supported in filenames, and user agents
are likely to adjust filenames accordingly.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Element/a#attr-ping](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/a#attr-ping "The <a> HTML element (or anchor element), with its href attribute, creates a hyperlink to web pages, files, email addresses, locations in the same page, or anything else a URL can address.")

Support in all current engines.

::: support
[Firefox[🔰
1+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari6+]{.safari .yes}[Chrome12+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)17+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android≤37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

The [`ping`]{#ping .dfn dfn-for="a,area" dfn-type="element-attr"}
attribute, if present, gives the URLs of the resources that are
interested in being notified if the user follows the hyperlink. The
value must be a [set of space-separated
tokens](common-microsyntaxes.html#set-of-space-separated-tokens){#links-created-by-a-and-area-elements:set-of-space-separated-tokens},
each of which must be a [valid non-empty
URL](urls-and-fetching.html#valid-non-empty-url){#links-created-by-a-and-area-elements:valid-non-empty-url}
whose
[scheme](https://url.spec.whatwg.org/#concept-url-scheme){#links-created-by-a-and-area-elements:concept-url-scheme
x-internal="concept-url-scheme"} is an [HTTP(S)
scheme](https://fetch.spec.whatwg.org/#http-scheme){#links-created-by-a-and-area-elements:http(s)-scheme
x-internal="http(s)-scheme"}. The value is used by the user agent for
[hyperlink
auditing](#hyperlink-auditing){#links-created-by-a-and-area-elements:hyperlink-auditing}.

The [`rel`]{#attr-hyperlink-rel .dfn dfn-for="a,area"
dfn-type="element-attr"} attribute on
[`a`{#links-created-by-a-and-area-elements:the-a-element-4}](text-level-semantics.html#the-a-element)
and
[`area`{#links-created-by-a-and-area-elements:the-area-element-4}](image-maps.html#the-area-element)
elements controls what kinds of links the elements create. The
attribute\'s value must be an [unordered set of unique space-separated
tokens](common-microsyntaxes.html#unordered-set-of-unique-space-separated-tokens){#links-created-by-a-and-area-elements:unordered-set-of-unique-space-separated-tokens}.
The [allowed keywords and their meanings](#linkTypes) are defined below.

[`rel`{#links-created-by-a-and-area-elements:attr-hyperlink-rel}](#attr-hyperlink-rel)\'s
[supported
tokens](https://dom.spec.whatwg.org/#concept-supported-tokens){#links-created-by-a-and-area-elements:concept-supported-tokens
x-internal="concept-supported-tokens"} are the keywords defined in [HTML
link types](#linkTypes) which are allowed on
[`a`{#links-created-by-a-and-area-elements:the-a-element-5}](text-level-semantics.html#the-a-element)
and
[`area`{#links-created-by-a-and-area-elements:the-area-element-5}](image-maps.html#the-area-element)
elements, impact the processing model, and are supported by the user
agent. The possible [supported
tokens](https://dom.spec.whatwg.org/#concept-supported-tokens){#links-created-by-a-and-area-elements:concept-supported-tokens-2
x-internal="concept-supported-tokens"} are
[`noreferrer`{#links-created-by-a-and-area-elements:link-type-noreferrer}](#link-type-noreferrer),
[`noopener`{#links-created-by-a-and-area-elements:link-type-noopener}](#link-type-noopener),
and
[`opener`{#links-created-by-a-and-area-elements:link-type-opener}](#link-type-opener).
[`rel`{#links-created-by-a-and-area-elements:attr-hyperlink-rel-2}](#attr-hyperlink-rel)\'s
[supported
tokens](https://dom.spec.whatwg.org/#concept-supported-tokens){#links-created-by-a-and-area-elements:concept-supported-tokens-3
x-internal="concept-supported-tokens"} must only include the tokens from
this list that the user agent implements the processing model for.

The
[`rel`{#links-created-by-a-and-area-elements:attr-hyperlink-rel-3}](#attr-hyperlink-rel)
attribute has no default value. If the attribute is omitted or if none
of the values in the attribute are recognized by the user agent, then
the document has no particular relationship with the destination
resource other than there being a hyperlink between the two.

The [`hreflang`]{#attr-hyperlink-hreflang .dfn dfn-for="a,area"
dfn-type="element-attr"} attribute on
[`a`{#links-created-by-a-and-area-elements:the-a-element-6}](text-level-semantics.html#the-a-element)
elements that create
[hyperlinks](#hyperlink){#links-created-by-a-and-area-elements:hyperlink},
if present, gives the language of the linked resource. It is purely
advisory. The value must be a valid BCP 47 language tag.
[\[BCP47\]](references.html#refsBCP47) User agents must not consider
this attribute authoritative --- upon fetching the resource, user agents
must use only language information associated with the resource to
determine its language, not metadata included in the link to the
resource.

The [`type`]{#attr-hyperlink-type .dfn dfn-for="a,area"
dfn-type="element-attr"} attribute, if present, gives the [MIME
type](https://mimesniff.spec.whatwg.org/#mime-type){#links-created-by-a-and-area-elements:mime-type
x-internal="mime-type"} of the linked resource. It is purely advisory.
The value must be a [valid MIME type
string](https://mimesniff.spec.whatwg.org/#valid-mime-type){#links-created-by-a-and-area-elements:valid-mime-type-string
x-internal="valid-mime-type-string"}. User agents must not consider the
[`type`{#links-created-by-a-and-area-elements:attr-hyperlink-type}](#attr-hyperlink-type)
attribute authoritative --- upon fetching the resource, user agents must
not use metadata included in the link to the resource to determine its
type.

The [`referrerpolicy`]{#attr-hyperlink-referrerpolicy .dfn
dfn-for="a,area" dfn-type="element-attr"} attribute is a [referrer
policy
attribute](urls-and-fetching.html#referrer-policy-attribute){#links-created-by-a-and-area-elements:referrer-policy-attribute}.
Its purpose is to set the [referrer
policy](https://w3c.github.io/webappsec-referrer-policy/#referrer-policy){#links-created-by-a-and-area-elements:referrer-policy
x-internal="referrer-policy"} used when [following
hyperlinks](#following-hyperlinks-2){#links-created-by-a-and-area-elements:following-hyperlinks-2-2}.
[\[REFERRERPOLICY\]](references.html#refsREFERRERPOLICY)

------------------------------------------------------------------------

When an
[`a`{#links-created-by-a-and-area-elements:the-a-element-7}](text-level-semantics.html#the-a-element)
or
[`area`{#links-created-by-a-and-area-elements:the-area-element-6}](image-maps.html#the-area-element)
element\'s [activation
behavior](https://dom.spec.whatwg.org/#eventtarget-activation-behavior){#links-created-by-a-and-area-elements:activation-behaviour
x-internal="activation-behaviour"} is invoked, the user agent may allow
the user to indicate a preference regarding whether the hyperlink is to
be used for
[navigation](browsing-the-web.html#navigate){#links-created-by-a-and-area-elements:navigate}
or whether the resource it specifies is to be downloaded.

In the absence of a user preference, the default should be navigation if
the element has no
[`download`{#links-created-by-a-and-area-elements:attr-hyperlink-download}](#attr-hyperlink-download)
attribute, and should be to download the specified resource if it does.

The [activation
behavior](https://dom.spec.whatwg.org/#eventtarget-activation-behavior){#links-created-by-a-and-area-elements:activation-behaviour-2
x-internal="activation-behaviour"} of an
[`a`{#links-created-by-a-and-area-elements:the-a-element-8}](text-level-semantics.html#the-a-element)
or
[`area`{#links-created-by-a-and-area-elements:the-area-element-7}](image-maps.html#the-area-element)
element `element`{.variable} given an event `event`{.variable} is:

1.  If `element`{.variable} has no
    [`href`{#links-created-by-a-and-area-elements:attr-hyperlink-href-3}](#attr-hyperlink-href)
    attribute, then return.

2.  Let `hyperlinkSuffix`{.variable} be null.

3.  If `element`{.variable} is an
    [`a`{#links-created-by-a-and-area-elements:the-a-element-9}](text-level-semantics.html#the-a-element)
    element, and `event`{.variable}\'s
    [target](https://dom.spec.whatwg.org/#concept-event-target){#links-created-by-a-and-area-elements:concept-event-target
    x-internal="concept-event-target"} is an
    [`img`{#links-created-by-a-and-area-elements:the-img-element}](embedded-content.html#the-img-element)
    with an
    [`ismap`{#links-created-by-a-and-area-elements:attr-img-ismap}](embedded-content.html#attr-img-ismap)
    attribute specified, then:

    1.  Let `x`{.variable} and `y`{.variable} be 0.

    2.  If `event`{.variable}\'s
        [`isTrusted`{#links-created-by-a-and-area-elements:dom-event-istrusted}](https://dom.spec.whatwg.org/#dom-event-istrusted){x-internal="dom-event-istrusted"}
        attribute is initialized to true, then set `x`{.variable} to the
        distance in [CSS
        pixels](https://drafts.csswg.org/css-values/#px){#links-created-by-a-and-area-elements:'px'
        x-internal="'px'"} from the left edge of the image to the
        location of the click, and set `y`{.variable} to the distance in
        [CSS
        pixels](https://drafts.csswg.org/css-values/#px){#links-created-by-a-and-area-elements:'px'-2
        x-internal="'px'"} from the top edge of the image to the
        location of the click.

    3.  If `x`{.variable} is negative, set `x`{.variable} to 0.

    4.  If `y`{.variable} is negative, set `y`{.variable} to 0.

    5.  Set `hyperlinkSuffix`{.variable} to the concatenation of U+003F
        (?), the value of `x`{.variable} expressed as a base-ten integer
        using [ASCII
        digits](https://infra.spec.whatwg.org/#ascii-digit){#links-created-by-a-and-area-elements:ascii-digits
        x-internal="ascii-digits"}, U+002C (,), and the value of
        `y`{.variable} expressed as a base-ten integer using [ASCII
        digits](https://infra.spec.whatwg.org/#ascii-digit){#links-created-by-a-and-area-elements:ascii-digits-2
        x-internal="ascii-digits"}.

4.  Let `userInvolvement`{.variable} be `event`{.variable}\'s [user
    navigation
    involvement](browsing-the-web.html#event-uni){#links-created-by-a-and-area-elements:event-uni}.

5.  If the user has expressed a preference to download the hyperlink,
    then set `userInvolvement`{.variable} to
    \"[`browser UI`{#links-created-by-a-and-area-elements:uni-browser-ui}](browsing-the-web.html#uni-browser-ui)\".

    That is, if the user has expressed a specific preference for
    downloading, this no longer counts as merely
    \"[`activation`{#links-created-by-a-and-area-elements:uni-activation}](browsing-the-web.html#uni-activation)\".

6.  If `element`{.variable} has a
    [`download`{#links-created-by-a-and-area-elements:attr-hyperlink-download-2}](#attr-hyperlink-download)
    attribute, or if the user has expressed a preference to download the
    hyperlink, then [download the
    hyperlink](#downloading-hyperlinks){#links-created-by-a-and-area-elements:downloading-hyperlinks-2}
    created by `element`{.variable} with
    *[hyperlinkSuffix](#downloading-hyperlinksuffix)* set to
    `hyperlinkSuffix`{.variable} and
    *[userInvolvement](#downloading-userinvolvement)* set to
    `userInvolvement`{.variable}.

7.  Otherwise, [follow the
    hyperlink](#following-hyperlinks-2){#links-created-by-a-and-area-elements:following-hyperlinks-2-3}
    created by `element`{.variable} with
    *[hyperlinkSuffix](#following-hyperlinksuffix)* set to
    `hyperlinkSuffix`{.variable} and
    *[userInvolvement](#following-userinvolvement)* set to
    `userInvolvement`{.variable}.

#### [4.6.3]{.secno} API for [`a`{#api-for-a-and-area-elements:the-a-element}](text-level-semantics.html#the-a-element) and [`area`{#api-for-a-and-area-elements:the-area-element}](image-maps.html#the-area-element) elements[](#api-for-a-and-area-elements){.self-link}

``` idl
interface mixin HTMLHyperlinkElementUtils {
  [CEReactions] stringifier attribute USVString href;
  readonly attribute USVString origin;
  [CEReactions] attribute USVString protocol;
  [CEReactions] attribute USVString username;
  [CEReactions] attribute USVString password;
  [CEReactions] attribute USVString host;
  [CEReactions] attribute USVString hostname;
  [CEReactions] attribute USVString port;
  [CEReactions] attribute USVString pathname;
  [CEReactions] attribute USVString search;
  [CEReactions] attribute USVString hash;
};
```

`hyperlink`{.variable}`.``toString()`

`hyperlink`{.variable}`.`[`href`](#dom-hyperlink-href){#dom-hyperlink-href-dev}

::::::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLAnchorElement/href](https://developer.mozilla.org/en-US/docs/Web/API/HTMLAnchorElement/href "The HTMLAnchorElement.href property is a stringifier that returns a string containing the whole URL, and allows the href to be updated.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::

:::: feature
[HTMLAnchorElement/toString](https://developer.mozilla.org/en-US/docs/Web/API/HTMLAnchorElement/toString "The HTMLAnchorElement.toString() stringifier method returns a string containing the whole URL. It is a read-only version of HTMLAnchorElement.href.")

Support in all current engines.

::: support
[Firefox22+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome52+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::

:::: feature
[HTMLAreaElement/href](https://developer.mozilla.org/en-US/docs/Web/API/HTMLAreaElement/href "The HTMLAreaElement.href property is a stringifier that returns a string containing the whole URL, and allows the href to be updated.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::

:::: feature
[HTMLAreaElement/toString](https://developer.mozilla.org/en-US/docs/Web/API/HTMLAreaElement/toString "The HTMLAreaElement.toString() stringifier method returns a string containing the whole URL. It is a read-only version of HTMLAreaElement.href.")

Support in all current engines.

::: support
[Firefox22+]{.firefox .yes}[Safari10.1+]{.safari
.yes}[Chrome32+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::::::::

Returns the hyperlink\'s URL.

Can be set, to change the URL.

`hyperlink`{.variable}`.`[`origin`](#dom-hyperlink-origin){#dom-hyperlink-origin-dev}

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLAnchorElement/origin](https://developer.mozilla.org/en-US/docs/Web/API/HTMLAnchorElement/origin "The HTMLAnchorElement.origin read-only property is a string containing the Unicode serialization of the origin of the represented URL.")

Support in all current engines.

::: support
[Firefox26+]{.firefox .yes}[Safari5.1+]{.safari .yes}[Chrome8+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)17+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android3+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::

:::: feature
[HTMLAreaElement/origin](https://developer.mozilla.org/en-US/docs/Web/API/HTMLAreaElement/origin "The HTMLAreaElement.origin read-only property is a string containing the Unicode serialization of the origin of the represented URL.")

Support in all current engines.

::: support
[Firefox26+]{.firefox .yes}[Safari10+]{.safari .yes}[Chrome32+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)17+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::::

Returns the hyperlink\'s URL\'s origin.

`hyperlink`{.variable}`.`[`protocol`](#dom-hyperlink-protocol){#dom-hyperlink-protocol-dev}

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLAnchorElement/protocol](https://developer.mozilla.org/en-US/docs/Web/API/HTMLAnchorElement/protocol "The HTMLAnchorElement.protocol property is a string representing the protocol scheme of the URL, including the final ':'.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::

:::: feature
[HTMLAreaElement/protocol](https://developer.mozilla.org/en-US/docs/Web/API/HTMLAreaElement/protocol "The HTMLAreaElement.protocol property is a string representing the protocol scheme of the URL, including the final ':'.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::::

Returns the hyperlink\'s URL\'s scheme.

Can be set, to change the URL\'s scheme.

`hyperlink`{.variable}`.`[`username`](#dom-hyperlink-username){#dom-hyperlink-username-dev}

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLAnchorElement/username](https://developer.mozilla.org/en-US/docs/Web/API/HTMLAnchorElement/username "The HTMLAnchorElement.username property is a string containing the username specified before the domain name.")

Support in all current engines.

::: support
[Firefox26+]{.firefox .yes}[Safari10+]{.safari .yes}[Chrome32+]{.chrome
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

:::: feature
[HTMLAreaElement/username](https://developer.mozilla.org/en-US/docs/Web/API/HTMLAreaElement/username "The HTMLAreaElement.username property is a string containing the username specified before the domain name.")

Support in all current engines.

::: support
[Firefox26+]{.firefox .yes}[Safari10+]{.safari .yes}[Chrome32+]{.chrome
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
:::::::

Returns the hyperlink\'s URL\'s username.

Can be set, to change the URL\'s username.

`hyperlink`{.variable}`.`[`password`](#dom-hyperlink-password){#dom-hyperlink-password-dev}

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLAnchorElement/password](https://developer.mozilla.org/en-US/docs/Web/API/HTMLAnchorElement/password "The HTMLAnchorElement.password property is a string containing the password specified before the domain name.")

Support in all current engines.

::: support
[Firefox26+]{.firefox .yes}[Safari10+]{.safari .yes}[Chrome32+]{.chrome
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

:::: feature
[HTMLAreaElement/password](https://developer.mozilla.org/en-US/docs/Web/API/HTMLAreaElement/password "The HTMLAreaElement.password property is a string containing the password specified before the domain name.")

Support in all current engines.

::: support
[Firefox26+]{.firefox .yes}[Safari10+]{.safari .yes}[Chrome32+]{.chrome
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
:::::::

Returns the hyperlink\'s URL\'s password.

Can be set, to change the URL\'s password.

`hyperlink`{.variable}`.`[`host`](#dom-hyperlink-host){#dom-hyperlink-host-dev}

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLAnchorElement/host](https://developer.mozilla.org/en-US/docs/Web/API/HTMLAnchorElement/host "The HTMLAnchorElement.host property is a string containing the host, that is the hostname, and then, if the port of the URL is nonempty, a ':', and the port of the URL.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::

:::: feature
[HTMLAreaElement/host](https://developer.mozilla.org/en-US/docs/Web/API/HTMLAreaElement/host "The HTMLAreaElement.host property is a string containing the host, that is the hostname, and then, if the port of the URL is nonempty, a ':', and the port of the URL.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::::

Returns the hyperlink\'s URL\'s host and port (if different from the
default port for the scheme).

Can be set, to change the URL\'s host and port.

`hyperlink`{.variable}`.`[`hostname`](#dom-hyperlink-hostname){#dom-hyperlink-hostname-dev}

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLAnchorElement/hostname](https://developer.mozilla.org/en-US/docs/Web/API/HTMLAnchorElement/hostname "The HTMLAnchorElement.hostname property is a string containing the domain of the URL.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::

:::: feature
[HTMLAreaElement/hostname](https://developer.mozilla.org/en-US/docs/Web/API/HTMLAreaElement/hostname "The HTMLAreaElement.hostname property is a string containing the domain of the URL.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::::

Returns the hyperlink\'s URL\'s host.

Can be set, to change the URL\'s host.

`hyperlink`{.variable}`.`[`port`](#dom-hyperlink-port){#dom-hyperlink-port-dev}

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLAnchorElement/port](https://developer.mozilla.org/en-US/docs/Web/API/HTMLAnchorElement/port "The HTMLAnchorElement.port property is a string containing the port number of the URL. If the URL does not contain an explicit port number, it will be set to ''.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::

:::: feature
[HTMLAreaElement/port](https://developer.mozilla.org/en-US/docs/Web/API/HTMLAreaElement/port "The HTMLAreaElement.port property is a string containing the port number of the URL. If the URL does not contain an explicit port number, it will be set to ''.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::::

Returns the hyperlink\'s URL\'s port.

Can be set, to change the URL\'s port.

`hyperlink`{.variable}`.`[`pathname`](#dom-hyperlink-pathname){#dom-hyperlink-pathname-dev}

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLAnchorElement/pathname](https://developer.mozilla.org/en-US/docs/Web/API/HTMLAnchorElement/pathname "The HTMLAnchorElement.pathname property is a string containing an initial '/' followed by the path of the URL not including the query string or fragment (or the empty string if there is no path).")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::

:::: feature
[HTMLAreaElement/pathname](https://developer.mozilla.org/en-US/docs/Web/API/HTMLAreaElement/pathname "The HTMLAreaElement.pathname property is a string containing an initial '/' followed by the path of the URL not including the query string or fragment (or the empty string if there is no path).")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::::

Returns the hyperlink\'s URL\'s path.

Can be set, to change the URL\'s path.

`hyperlink`{.variable}`.`[`search`](#dom-hyperlink-search){#dom-hyperlink-search-dev}

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLAnchorElement/search](https://developer.mozilla.org/en-US/docs/Web/API/HTMLAnchorElement/search "The HTMLAnchorElement.search property is a search string, also called a query string, that is a string containing a '?' followed by the parameters of the URL.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::

:::: feature
[HTMLAreaElement/search](https://developer.mozilla.org/en-US/docs/Web/API/HTMLAreaElement/search "The HTMLAreaElement.search property is a search string, also called a query string, that is a string containing a '?' followed by the parameters of the URL.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::::

Returns the hyperlink\'s URL\'s query (includes leading \"`?`\" if
non-empty).

Can be set, to change the URL\'s query (ignores leading \"`?`\").

`hyperlink`{.variable}`.`[`hash`](#dom-hyperlink-hash){#dom-hyperlink-hash-dev}

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLAnchorElement/hash](https://developer.mozilla.org/en-US/docs/Web/API/HTMLAnchorElement/hash "The HTMLAnchorElement.hash property returns a string containing a '#' followed by the fragment identifier of the URL.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::

:::: feature
[HTMLAreaElement/hash](https://developer.mozilla.org/en-US/docs/Web/API/HTMLAreaElement/hash "The HTMLAreaElement.hash property returns a string containing a '#' followed by the fragment identifier of the URL.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::::

Returns the hyperlink\'s URL\'s fragment (includes leading \"`#`\" if
non-empty).

Can be set, to change the URL\'s fragment (ignores leading \"`#`\").

An element implementing the
[`HTMLHyperlinkElementUtils`{#api-for-a-and-area-elements:htmlhyperlinkelementutils}](#htmlhyperlinkelementutils)
mixin has an associated [url]{#concept-hyperlink-url .dfn} (null or a
[URL](https://url.spec.whatwg.org/#concept-url){#api-for-a-and-area-elements:url
x-internal="url"}). It is initially null.

An element implementing the
[`HTMLHyperlinkElementUtils`{#api-for-a-and-area-elements:htmlhyperlinkelementutils-2}](#htmlhyperlinkelementutils)
mixin has an associated [set the url]{#concept-hyperlink-url-set .dfn}
algorithm, which runs these steps:

1.  Set this element\'s
    [url](#concept-hyperlink-url){#api-for-a-and-area-elements:concept-hyperlink-url}
    to null.

2.  If this element\'s
    [`href`{#api-for-a-and-area-elements:attr-hyperlink-href}](#attr-hyperlink-href)
    content attribute is absent, then return.

3.  Let `url`{.variable} be the result of [encoding-parsing a
    URL](urls-and-fetching.html#encoding-parsing-a-url){#api-for-a-and-area-elements:encoding-parsing-a-url}
    given this element\'s
    [`href`{#api-for-a-and-area-elements:attr-hyperlink-href-2}](#attr-hyperlink-href)
    content attribute\'s value, relative to this element\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#api-for-a-and-area-elements:node-document
    x-internal="node-document"}.

4.  If `url`{.variable} is not failure, then set this element\'s
    [url](#concept-hyperlink-url){#api-for-a-and-area-elements:concept-hyperlink-url-2}
    to `url`{.variable}.

When elements implementing the
[`HTMLHyperlinkElementUtils`{#api-for-a-and-area-elements:htmlhyperlinkelementutils-3}](#htmlhyperlinkelementutils)
mixin are created, and whenever those elements have their
[`href`{#api-for-a-and-area-elements:attr-hyperlink-href-3}](#attr-hyperlink-href)
content attribute set, changed, or removed, the user agent must [set the
url](#concept-hyperlink-url-set){#api-for-a-and-area-elements:concept-hyperlink-url-set}.

This is only observable for
[`blob:`{#api-for-a-and-area-elements:blob-protocol}](https://w3c.github.io/FileAPI/#DefinitionOfScheme){x-internal="blob-protocol"}
URLs as
[parsing](https://url.spec.whatwg.org/#concept-url-parser){#api-for-a-and-area-elements:url-parser
x-internal="url-parser"} them involves a [Blob URL
Store](https://w3c.github.io/FileAPI/#BlobURLStore){#api-for-a-and-area-elements:blob-url-store
x-internal="blob-url-store"} lookup.

An element implementing the
[`HTMLHyperlinkElementUtils`{#api-for-a-and-area-elements:htmlhyperlinkelementutils-4}](#htmlhyperlinkelementutils)
mixin has an associated [reinitialize url]{#reinitialise-url .dfn}
algorithm, which runs these steps:

1.  If the element\'s
    [url](#concept-hyperlink-url){#api-for-a-and-area-elements:concept-hyperlink-url-3}
    is non-null, its
    [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#api-for-a-and-area-elements:concept-url-scheme
    x-internal="concept-url-scheme"} is \"`blob`\", and it has an
    [opaque
    path](https://url.spec.whatwg.org/#url-opaque-path){#api-for-a-and-area-elements:opaque-path
    x-internal="opaque-path"}, then terminate these steps.

2.  [Set the
    url](#concept-hyperlink-url-set){#api-for-a-and-area-elements:concept-hyperlink-url-set-2}.

To [update `href`]{#update-href .dfn}, set the element\'s
[`href`{#api-for-a-and-area-elements:attr-hyperlink-href-4}](#attr-hyperlink-href)
content attribute\'s value to the element\'s
[url](#concept-hyperlink-url){#api-for-a-and-area-elements:concept-hyperlink-url-4},
[serialized](https://url.spec.whatwg.org/#concept-url-serializer){#api-for-a-and-area-elements:concept-url-serializer
x-internal="concept-url-serializer"}.

------------------------------------------------------------------------

The [`href`]{#dom-hyperlink-href .dfn
dfn-for="HTMLHyperlinkElementUtils" dfn-type="attribute"} getter steps
are:

1.  [Reinitialize
    url](#reinitialise-url){#api-for-a-and-area-elements:reinitialise-url}.

2.  Let `url`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#api-for-a-and-area-elements:this
    x-internal="this"}\'s
    [url](#concept-hyperlink-url){#api-for-a-and-area-elements:concept-hyperlink-url-5}.

3.  If `url`{.variable} is null and
    [this](https://webidl.spec.whatwg.org/#this){#api-for-a-and-area-elements:this-2
    x-internal="this"} has no
    [`href`{#api-for-a-and-area-elements:attr-hyperlink-href-5}](#attr-hyperlink-href)
    content attribute, return the empty string.

4.  Otherwise, if `url`{.variable} is null, return
    [this](https://webidl.spec.whatwg.org/#this){#api-for-a-and-area-elements:this-3
    x-internal="this"}\'s
    [`href`{#api-for-a-and-area-elements:attr-hyperlink-href-6}](#attr-hyperlink-href)
    content attribute\'s value.

5.  Return `url`{.variable},
    [serialized](https://url.spec.whatwg.org/#concept-url-serializer){#api-for-a-and-area-elements:concept-url-serializer-2
    x-internal="concept-url-serializer"}.

The
[`href`{#api-for-a-and-area-elements:dom-hyperlink-href-2}](#dom-hyperlink-href)
setter steps are to set
[this](https://webidl.spec.whatwg.org/#this){#api-for-a-and-area-elements:this-4
x-internal="this"}\'s
[`href`{#api-for-a-and-area-elements:attr-hyperlink-href-7}](#attr-hyperlink-href)
content attribute\'s value to the given value.

The [`origin`]{#dom-hyperlink-origin .dfn
dfn-for="HTMLHyperlinkElementUtils" dfn-type="attribute"} getter steps
are:

1.  [Reinitialize
    url](#reinitialise-url){#api-for-a-and-area-elements:reinitialise-url-2}.

2.  If
    [this](https://webidl.spec.whatwg.org/#this){#api-for-a-and-area-elements:this-5
    x-internal="this"}\'s
    [url](#concept-hyperlink-url){#api-for-a-and-area-elements:concept-hyperlink-url-6}
    is null, return the empty string.

3.  Return the
    [serialization](browsers.html#ascii-serialisation-of-an-origin){#api-for-a-and-area-elements:ascii-serialisation-of-an-origin}
    of
    [this](https://webidl.spec.whatwg.org/#this){#api-for-a-and-area-elements:this-6
    x-internal="this"}\'s
    [url](#concept-hyperlink-url){#api-for-a-and-area-elements:concept-hyperlink-url-7}\'s
    [origin](https://url.spec.whatwg.org/#concept-url-origin){#api-for-a-and-area-elements:concept-url-origin
    x-internal="concept-url-origin"}.

The [`protocol`]{#dom-hyperlink-protocol .dfn
dfn-for="HTMLHyperlinkElementUtils" dfn-type="attribute"} getter steps
are:

1.  [Reinitialize
    url](#reinitialise-url){#api-for-a-and-area-elements:reinitialise-url-3}.

2.  If
    [this](https://webidl.spec.whatwg.org/#this){#api-for-a-and-area-elements:this-7
    x-internal="this"}\'s
    [url](#concept-hyperlink-url){#api-for-a-and-area-elements:concept-hyperlink-url-8}
    is null, return \"`:`\".

3.  Return
    [this](https://webidl.spec.whatwg.org/#this){#api-for-a-and-area-elements:this-8
    x-internal="this"}\'s
    [url](#concept-hyperlink-url){#api-for-a-and-area-elements:concept-hyperlink-url-9}\'s
    [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#api-for-a-and-area-elements:concept-url-scheme-2
    x-internal="concept-url-scheme"}, followed by \"`:`\".

The
[`protocol`{#api-for-a-and-area-elements:dom-hyperlink-protocol-2}](#dom-hyperlink-protocol)
setter steps are:

1.  [Reinitialize
    url](#reinitialise-url){#api-for-a-and-area-elements:reinitialise-url-4}.

2.  If
    [this](https://webidl.spec.whatwg.org/#this){#api-for-a-and-area-elements:this-9
    x-internal="this"}\'s
    [url](#concept-hyperlink-url){#api-for-a-and-area-elements:concept-hyperlink-url-10}
    is null, then return.

3.  [Basic URL
    parse](https://url.spec.whatwg.org/#concept-basic-url-parser){#api-for-a-and-area-elements:basic-url-parser
    x-internal="basic-url-parser"} the given value, followed by \"`:`\",
    with
    [this](https://webidl.spec.whatwg.org/#this){#api-for-a-and-area-elements:this-10
    x-internal="this"}\'s
    [url](#concept-hyperlink-url){#api-for-a-and-area-elements:concept-hyperlink-url-11}
    as
    [*url*](https://url.spec.whatwg.org/#basic-url-parser-url){#api-for-a-and-area-elements:basic-url-parser-url
    x-internal="basic-url-parser-url"} and [scheme start
    state](https://url.spec.whatwg.org/#scheme-start-state){#api-for-a-and-area-elements:scheme-start-state
    x-internal="scheme-start-state"} as [*state
    override*](https://url.spec.whatwg.org/#basic-url-parser-state-override){#api-for-a-and-area-elements:basic-url-parser-state-override
    x-internal="basic-url-parser-state-override"}.

    Because the URL parser ignores multiple consecutive colons,
    providing a value of \"`https:`\" (or even \"`https::::`\") is the
    same as providing a value of \"`https`\".

4.  [Update
    `href`](#update-href){#api-for-a-and-area-elements:update-href}.

The [`username`]{#dom-hyperlink-username .dfn
dfn-for="HTMLHyperlinkElementUtils" dfn-type="attribute"} getter steps
are:

1.  [Reinitialize
    url](#reinitialise-url){#api-for-a-and-area-elements:reinitialise-url-5}.

2.  If
    [this](https://webidl.spec.whatwg.org/#this){#api-for-a-and-area-elements:this-11
    x-internal="this"}\'s
    [url](#concept-hyperlink-url){#api-for-a-and-area-elements:concept-hyperlink-url-12}
    is null, return the empty string.

3.  Return
    [this](https://webidl.spec.whatwg.org/#this){#api-for-a-and-area-elements:this-12
    x-internal="this"}\'s
    [url](#concept-hyperlink-url){#api-for-a-and-area-elements:concept-hyperlink-url-13}\'s
    [username](https://url.spec.whatwg.org/#concept-url-username){#api-for-a-and-area-elements:concept-url-username
    x-internal="concept-url-username"}.

The
[`username`{#api-for-a-and-area-elements:dom-hyperlink-username-2}](#dom-hyperlink-username)
setter steps are:

1.  [Reinitialize
    url](#reinitialise-url){#api-for-a-and-area-elements:reinitialise-url-6}.

2.  Let `url`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#api-for-a-and-area-elements:this-13
    x-internal="this"}\'s
    [url](#concept-hyperlink-url){#api-for-a-and-area-elements:concept-hyperlink-url-14}.

3.  If `url`{.variable} is null or `url`{.variable} [cannot have a
    username/password/port](https://url.spec.whatwg.org/#cannot-have-a-username-password-port){#api-for-a-and-area-elements:cannot-have-a-username/password/port
    x-internal="cannot-have-a-username/password/port"}, then return.

4.  [Set the
    username](https://url.spec.whatwg.org/#set-the-username){#api-for-a-and-area-elements:set-the-username
    x-internal="set-the-username"}, given `url`{.variable} and the given
    value.

5.  [Update
    `href`](#update-href){#api-for-a-and-area-elements:update-href-2}.

The [`password`]{#dom-hyperlink-password .dfn
dfn-for="HTMLHyperlinkElementUtils" dfn-type="attribute"} getter steps
are:

1.  [Reinitialize
    url](#reinitialise-url){#api-for-a-and-area-elements:reinitialise-url-7}.

2.  Let `url`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#api-for-a-and-area-elements:this-14
    x-internal="this"}\'s
    [url](#concept-hyperlink-url){#api-for-a-and-area-elements:concept-hyperlink-url-15}.

3.  If `url`{.variable} is null, then return the empty string.

4.  Return `url`{.variable}\'s
    [password](https://url.spec.whatwg.org/#concept-url-password){#api-for-a-and-area-elements:concept-url-password
    x-internal="concept-url-password"}.

The
[`password`{#api-for-a-and-area-elements:dom-hyperlink-password-2}](#dom-hyperlink-password)
setter steps are:

1.  [Reinitialize
    url](#reinitialise-url){#api-for-a-and-area-elements:reinitialise-url-8}.

2.  Let `url`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#api-for-a-and-area-elements:this-15
    x-internal="this"}\'s
    [url](#concept-hyperlink-url){#api-for-a-and-area-elements:concept-hyperlink-url-16}.

3.  If `url`{.variable} is null or `url`{.variable} [cannot have a
    username/password/port](https://url.spec.whatwg.org/#cannot-have-a-username-password-port){#api-for-a-and-area-elements:cannot-have-a-username/password/port-2
    x-internal="cannot-have-a-username/password/port"}, then return.

4.  [Set the
    password](https://url.spec.whatwg.org/#set-the-password){#api-for-a-and-area-elements:set-the-password
    x-internal="set-the-password"}, given `url`{.variable} and the given
    value.

5.  [Update
    `href`](#update-href){#api-for-a-and-area-elements:update-href-3}.

The [`host`]{#dom-hyperlink-host .dfn
dfn-for="HTMLHyperlinkElementUtils" dfn-type="attribute"} getter steps
are:

1.  [Reinitialize
    url](#reinitialise-url){#api-for-a-and-area-elements:reinitialise-url-9}.

2.  Let `url`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#api-for-a-and-area-elements:this-16
    x-internal="this"}\'s
    [url](#concept-hyperlink-url){#api-for-a-and-area-elements:concept-hyperlink-url-17}.

3.  If `url`{.variable} or `url`{.variable}\'s
    [host](https://url.spec.whatwg.org/#concept-url-host){#api-for-a-and-area-elements:concept-url-host
    x-internal="concept-url-host"} is null, return the empty string.

4.  If `url`{.variable}\'s
    [port](https://url.spec.whatwg.org/#concept-url-port){#api-for-a-and-area-elements:concept-url-port
    x-internal="concept-url-port"} is null, return `url`{.variable}\'s
    [host](https://url.spec.whatwg.org/#concept-url-host){#api-for-a-and-area-elements:concept-url-host-2
    x-internal="concept-url-host"},
    [serialized](https://url.spec.whatwg.org/#concept-host-serializer){#api-for-a-and-area-elements:host-serializer
    x-internal="host-serializer"}.

5.  Return `url`{.variable}\'s
    [host](https://url.spec.whatwg.org/#concept-url-host){#api-for-a-and-area-elements:concept-url-host-3
    x-internal="concept-url-host"},
    [serialized](https://url.spec.whatwg.org/#concept-host-serializer){#api-for-a-and-area-elements:host-serializer-2
    x-internal="host-serializer"}, followed by \"`:`\" and
    `url`{.variable}\'s
    [port](https://url.spec.whatwg.org/#concept-url-port){#api-for-a-and-area-elements:concept-url-port-2
    x-internal="concept-url-port"},
    [serialized](https://url.spec.whatwg.org/#serialize-an-integer){#api-for-a-and-area-elements:serialize-an-integer
    x-internal="serialize-an-integer"}.

The
[`host`{#api-for-a-and-area-elements:dom-hyperlink-host-2}](#dom-hyperlink-host)
setter steps are:

1.  [Reinitialize
    url](#reinitialise-url){#api-for-a-and-area-elements:reinitialise-url-10}.

2.  Let `url`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#api-for-a-and-area-elements:this-17
    x-internal="this"}\'s
    [url](#concept-hyperlink-url){#api-for-a-and-area-elements:concept-hyperlink-url-18}.

3.  If `url`{.variable} is null or `url`{.variable} has an [opaque
    path](https://url.spec.whatwg.org/#url-opaque-path){#api-for-a-and-area-elements:opaque-path-2
    x-internal="opaque-path"}, then return.

4.  [Basic URL
    parse](https://url.spec.whatwg.org/#concept-basic-url-parser){#api-for-a-and-area-elements:basic-url-parser-2
    x-internal="basic-url-parser"} the given value, with
    `url`{.variable} as
    [*url*](https://url.spec.whatwg.org/#basic-url-parser-url){#api-for-a-and-area-elements:basic-url-parser-url-2
    x-internal="basic-url-parser-url"} and [host
    state](https://url.spec.whatwg.org/#host-state){#api-for-a-and-area-elements:host-state
    x-internal="host-state"} as [*state
    override*](https://url.spec.whatwg.org/#basic-url-parser-state-override){#api-for-a-and-area-elements:basic-url-parser-state-override-2
    x-internal="basic-url-parser-state-override"}.

5.  [Update
    `href`](#update-href){#api-for-a-and-area-elements:update-href-4}.

The [`hostname`]{#dom-hyperlink-hostname .dfn
dfn-for="HTMLHyperlinkElementUtils" dfn-type="attribute"} getter steps
are:

1.  [Reinitialize
    url](#reinitialise-url){#api-for-a-and-area-elements:reinitialise-url-11}.

2.  Let `url`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#api-for-a-and-area-elements:this-18
    x-internal="this"}\'s
    [url](#concept-hyperlink-url){#api-for-a-and-area-elements:concept-hyperlink-url-19}.

3.  If `url`{.variable} or `url`{.variable}\'s
    [host](https://url.spec.whatwg.org/#concept-url-host){#api-for-a-and-area-elements:concept-url-host-4
    x-internal="concept-url-host"} is null, return the empty string.

4.  Return `url`{.variable}\'s
    [host](https://url.spec.whatwg.org/#concept-url-host){#api-for-a-and-area-elements:concept-url-host-5
    x-internal="concept-url-host"},
    [serialized](https://url.spec.whatwg.org/#concept-host-serializer){#api-for-a-and-area-elements:host-serializer-3
    x-internal="host-serializer"}.

The
[`hostname`{#api-for-a-and-area-elements:dom-hyperlink-hostname-2}](#dom-hyperlink-hostname)
setter steps are:

1.  [Reinitialize
    url](#reinitialise-url){#api-for-a-and-area-elements:reinitialise-url-12}.

2.  Let `url`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#api-for-a-and-area-elements:this-19
    x-internal="this"}\'s
    [url](#concept-hyperlink-url){#api-for-a-and-area-elements:concept-hyperlink-url-20}.

3.  If `url`{.variable} is null or `url`{.variable} has an [opaque
    path](https://url.spec.whatwg.org/#url-opaque-path){#api-for-a-and-area-elements:opaque-path-3
    x-internal="opaque-path"}, then return.

4.  [Basic URL
    parse](https://url.spec.whatwg.org/#concept-basic-url-parser){#api-for-a-and-area-elements:basic-url-parser-3
    x-internal="basic-url-parser"} the given value, with
    `url`{.variable} as
    [*url*](https://url.spec.whatwg.org/#basic-url-parser-url){#api-for-a-and-area-elements:basic-url-parser-url-3
    x-internal="basic-url-parser-url"} and [hostname
    state](https://url.spec.whatwg.org/#hostname-state){#api-for-a-and-area-elements:hostname-state
    x-internal="hostname-state"} as [*state
    override*](https://url.spec.whatwg.org/#basic-url-parser-state-override){#api-for-a-and-area-elements:basic-url-parser-state-override-3
    x-internal="basic-url-parser-state-override"}.

5.  [Update
    `href`](#update-href){#api-for-a-and-area-elements:update-href-5}.

The [`port`]{#dom-hyperlink-port .dfn
dfn-for="HTMLHyperlinkElementUtils" dfn-type="attribute"} getter steps
are:

1.  [Reinitialize
    url](#reinitialise-url){#api-for-a-and-area-elements:reinitialise-url-13}.

2.  Let `url`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#api-for-a-and-area-elements:this-20
    x-internal="this"}\'s
    [url](#concept-hyperlink-url){#api-for-a-and-area-elements:concept-hyperlink-url-21}.

3.  If `url`{.variable} or `url`{.variable}\'s
    [port](https://url.spec.whatwg.org/#concept-url-port){#api-for-a-and-area-elements:concept-url-port-3
    x-internal="concept-url-port"} is null, return the empty string.

4.  Return `url`{.variable}\'s
    [port](https://url.spec.whatwg.org/#concept-url-port){#api-for-a-and-area-elements:concept-url-port-4
    x-internal="concept-url-port"},
    [serialized](https://url.spec.whatwg.org/#serialize-an-integer){#api-for-a-and-area-elements:serialize-an-integer-2
    x-internal="serialize-an-integer"}.

The
[`port`{#api-for-a-and-area-elements:dom-hyperlink-port-2}](#dom-hyperlink-port)
setter steps are:

1.  [Reinitialize
    url](#reinitialise-url){#api-for-a-and-area-elements:reinitialise-url-14}.

2.  Let `url`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#api-for-a-and-area-elements:this-21
    x-internal="this"}\'s
    [url](#concept-hyperlink-url){#api-for-a-and-area-elements:concept-hyperlink-url-22}.

3.  If `url`{.variable} is null or `url`{.variable} [cannot have a
    username/password/port](https://url.spec.whatwg.org/#cannot-have-a-username-password-port){#api-for-a-and-area-elements:cannot-have-a-username/password/port-3
    x-internal="cannot-have-a-username/password/port"}, then return.

4.  If the given value is the empty string, then set `url`{.variable}\'s
    [port](https://url.spec.whatwg.org/#concept-url-port){#api-for-a-and-area-elements:concept-url-port-5
    x-internal="concept-url-port"} to null.

5.  Otherwise, [basic URL
    parse](https://url.spec.whatwg.org/#concept-basic-url-parser){#api-for-a-and-area-elements:basic-url-parser-4
    x-internal="basic-url-parser"} the given value, with
    `url`{.variable} as
    [*url*](https://url.spec.whatwg.org/#basic-url-parser-url){#api-for-a-and-area-elements:basic-url-parser-url-4
    x-internal="basic-url-parser-url"} and [port
    state](https://url.spec.whatwg.org/#port-state){#api-for-a-and-area-elements:port-state
    x-internal="port-state"} as [*state
    override*](https://url.spec.whatwg.org/#basic-url-parser-state-override){#api-for-a-and-area-elements:basic-url-parser-state-override-4
    x-internal="basic-url-parser-state-override"}.

6.  [Update
    `href`](#update-href){#api-for-a-and-area-elements:update-href-6}.

The [`pathname`]{#dom-hyperlink-pathname .dfn
dfn-for="HTMLHyperlinkElementUtils" dfn-type="attribute"} getter steps
are:

1.  [Reinitialize
    url](#reinitialise-url){#api-for-a-and-area-elements:reinitialise-url-15}.

2.  Let `url`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#api-for-a-and-area-elements:this-22
    x-internal="this"}\'s
    [url](#concept-hyperlink-url){#api-for-a-and-area-elements:concept-hyperlink-url-23}.

3.  If `url`{.variable} is null, then return the empty string.

4.  Return the result of [URL path
    serializing](https://url.spec.whatwg.org/#url-path-serializer){#api-for-a-and-area-elements:url-path-serializer
    x-internal="url-path-serializer"} `url`{.variable}.

The
[`pathname`{#api-for-a-and-area-elements:dom-hyperlink-pathname-2}](#dom-hyperlink-pathname)
setter steps are:

1.  [Reinitialize
    url](#reinitialise-url){#api-for-a-and-area-elements:reinitialise-url-16}.

2.  Let `url`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#api-for-a-and-area-elements:this-23
    x-internal="this"}\'s
    [url](#concept-hyperlink-url){#api-for-a-and-area-elements:concept-hyperlink-url-24}.

3.  If `url`{.variable} is null or `url`{.variable} has an [opaque
    path](https://url.spec.whatwg.org/#url-opaque-path){#api-for-a-and-area-elements:opaque-path-4
    x-internal="opaque-path"}, then return.

4.  Set `url`{.variable}\'s
    [path](https://url.spec.whatwg.org/#concept-url-path){#api-for-a-and-area-elements:concept-url-path
    x-internal="concept-url-path"} to the empty list.

5.  [Basic URL
    parse](https://url.spec.whatwg.org/#concept-basic-url-parser){#api-for-a-and-area-elements:basic-url-parser-5
    x-internal="basic-url-parser"} the given value, with
    `url`{.variable} as
    [*url*](https://url.spec.whatwg.org/#basic-url-parser-url){#api-for-a-and-area-elements:basic-url-parser-url-5
    x-internal="basic-url-parser-url"} and [path start
    state](https://url.spec.whatwg.org/#path-start-state){#api-for-a-and-area-elements:path-start-state
    x-internal="path-start-state"} as [*state
    override*](https://url.spec.whatwg.org/#basic-url-parser-state-override){#api-for-a-and-area-elements:basic-url-parser-state-override-5
    x-internal="basic-url-parser-state-override"}.

6.  [Update
    `href`](#update-href){#api-for-a-and-area-elements:update-href-7}.

The [`search`]{#dom-hyperlink-search .dfn
dfn-for="HTMLHyperlinkElementUtils" dfn-type="attribute"} getter steps
are:

1.  [Reinitialize
    url](#reinitialise-url){#api-for-a-and-area-elements:reinitialise-url-17}.

2.  Let `url`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#api-for-a-and-area-elements:this-24
    x-internal="this"}\'s
    [url](#concept-hyperlink-url){#api-for-a-and-area-elements:concept-hyperlink-url-25}.

3.  If `url`{.variable} is null, or `url`{.variable}\'s
    [query](https://url.spec.whatwg.org/#concept-url-query){#api-for-a-and-area-elements:concept-url-query
    x-internal="concept-url-query"} is either null or the empty string,
    return the empty string.

4.  Return \"`?`\", followed by `url`{.variable}\'s
    [query](https://url.spec.whatwg.org/#concept-url-query){#api-for-a-and-area-elements:concept-url-query-2
    x-internal="concept-url-query"}.

The
[`search`{#api-for-a-and-area-elements:dom-hyperlink-search-2}](#dom-hyperlink-search)
setter steps are:

1.  [Reinitialize
    url](#reinitialise-url){#api-for-a-and-area-elements:reinitialise-url-18}.

2.  Let `url`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#api-for-a-and-area-elements:this-25
    x-internal="this"}\'s
    [url](#concept-hyperlink-url){#api-for-a-and-area-elements:concept-hyperlink-url-26}.

3.  If `url`{.variable} is null, terminate these steps.

4.  If the given value is the empty string, set `url`{.variable}\'s
    [query](https://url.spec.whatwg.org/#concept-url-query){#api-for-a-and-area-elements:concept-url-query-3
    x-internal="concept-url-query"} to null.

5.  Otherwise:

    1.  Let `input`{.variable} be the given value with a single leading
        \"`?`\" removed, if any.

    2.  Set `url`{.variable}\'s
        [query](https://url.spec.whatwg.org/#concept-url-query){#api-for-a-and-area-elements:concept-url-query-4
        x-internal="concept-url-query"} to the empty string.

    3.  [Basic URL
        parse](https://url.spec.whatwg.org/#concept-basic-url-parser){#api-for-a-and-area-elements:basic-url-parser-6
        x-internal="basic-url-parser"} `input`{.variable}, with
        `url`{.variable} as
        [*url*](https://url.spec.whatwg.org/#basic-url-parser-url){#api-for-a-and-area-elements:basic-url-parser-url-6
        x-internal="basic-url-parser-url"} and [query
        state](https://url.spec.whatwg.org/#query-state){#api-for-a-and-area-elements:query-state
        x-internal="query-state"} as [*state
        override*](https://url.spec.whatwg.org/#basic-url-parser-state-override){#api-for-a-and-area-elements:basic-url-parser-state-override-6
        x-internal="basic-url-parser-state-override"}.

6.  [Update
    `href`](#update-href){#api-for-a-and-area-elements:update-href-8}.

The [`hash`]{#dom-hyperlink-hash .dfn
dfn-for="HTMLHyperlinkElementUtils" dfn-type="attribute"} getter steps
are:

1.  [Reinitialize
    url](#reinitialise-url){#api-for-a-and-area-elements:reinitialise-url-19}.

2.  Let `url`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#api-for-a-and-area-elements:this-26
    x-internal="this"}\'s
    [url](#concept-hyperlink-url){#api-for-a-and-area-elements:concept-hyperlink-url-27}.

3.  If `url`{.variable} is null, or `url`{.variable}\'s
    [fragment](https://url.spec.whatwg.org/#concept-url-fragment){#api-for-a-and-area-elements:concept-url-fragment
    x-internal="concept-url-fragment"} is either null or the empty
    string, return the empty string.

4.  Return \"`#`\", followed by `url`{.variable}\'s
    [fragment](https://url.spec.whatwg.org/#concept-url-fragment){#api-for-a-and-area-elements:concept-url-fragment-2
    x-internal="concept-url-fragment"}.

The
[`hash`{#api-for-a-and-area-elements:dom-hyperlink-hash-2}](#dom-hyperlink-hash)
setter steps are:

1.  [Reinitialize
    url](#reinitialise-url){#api-for-a-and-area-elements:reinitialise-url-20}.

2.  Let `url`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#api-for-a-and-area-elements:this-27
    x-internal="this"}\'s
    [url](#concept-hyperlink-url){#api-for-a-and-area-elements:concept-hyperlink-url-28}.

3.  If `url`{.variable} is null, then return.

4.  If the given value is the empty string, set `url`{.variable}\'s
    [fragment](https://url.spec.whatwg.org/#concept-url-fragment){#api-for-a-and-area-elements:concept-url-fragment-3
    x-internal="concept-url-fragment"} to null.

5.  Otherwise:

    1.  Let `input`{.variable} be the given value with a single leading
        \"`#`\" removed, if any.

    2.  Set `url`{.variable}\'s
        [fragment](https://url.spec.whatwg.org/#concept-url-fragment){#api-for-a-and-area-elements:concept-url-fragment-4
        x-internal="concept-url-fragment"} to the empty string.

    3.  [Basic URL
        parse](https://url.spec.whatwg.org/#concept-basic-url-parser){#api-for-a-and-area-elements:basic-url-parser-7
        x-internal="basic-url-parser"} `input`{.variable}, with
        `url`{.variable} as
        [*url*](https://url.spec.whatwg.org/#basic-url-parser-url){#api-for-a-and-area-elements:basic-url-parser-url-7
        x-internal="basic-url-parser-url"} and [fragment
        state](https://url.spec.whatwg.org/#fragment-state){#api-for-a-and-area-elements:fragment-state
        x-internal="fragment-state"} as [*state
        override*](https://url.spec.whatwg.org/#basic-url-parser-state-override){#api-for-a-and-area-elements:basic-url-parser-state-override-7
        x-internal="basic-url-parser-state-override"}.

6.  [Update
    `href`](#update-href){#api-for-a-and-area-elements:update-href-9}.

#### [4.6.4]{.secno} Following hyperlinks[](#following-hyperlinks){.self-link}

An element `element`{.variable} [cannot navigate]{#cannot-navigate .dfn}
if any of the following are true:

- `element`{.variable}\'s [node
  document](https://dom.spec.whatwg.org/#concept-node-document){#following-hyperlinks:node-document
  x-internal="node-document"} is not [fully
  active](document-sequences.html#fully-active){#following-hyperlinks:fully-active};
  or

- `element`{.variable} is not an
  [`a`{#following-hyperlinks:the-a-element}](text-level-semantics.html#the-a-element)
  element and is not
  [connected](https://dom.spec.whatwg.org/#connected){#following-hyperlinks:connected
  x-internal="connected"}.

This is also used by [form
submission](form-control-infrastructure.html#concept-form-submit){#following-hyperlinks:concept-form-submit}
for the
[`form`{#following-hyperlinks:the-form-element}](forms.html#the-form-element)
element. The exception for
[`a`{#following-hyperlinks:the-a-element-2}](text-level-semantics.html#the-a-element)
elements is for compatibility with web content.

To [get an element\'s noopener]{#get-an-element's-noopener .dfn}, given
an
[`a`{#following-hyperlinks:the-a-element-3}](text-level-semantics.html#the-a-element),
[`area`{#following-hyperlinks:the-area-element}](image-maps.html#the-area-element),
or
[`form`{#following-hyperlinks:the-form-element-2}](forms.html#the-form-element)
element `element`{.variable}, a [URL
record](https://url.spec.whatwg.org/#concept-url){#following-hyperlinks:url-record
x-internal="url-record"} `url`{.variable}, and a string
`target`{.variable}, perform the following steps. They return a boolean.

1.  If `element`{.variable}\'s [link types](#linkTypes) include the
    [`noopener`{#following-hyperlinks:link-type-noopener}](#link-type-noopener)
    or
    [`noreferrer`{#following-hyperlinks:link-type-noreferrer}](#link-type-noreferrer)
    keyword, then return true.

2.  [If `element`{.variable}\'s [link types](#linkTypes) do not include
    the
    [`opener`{#following-hyperlinks:link-type-opener}](#link-type-opener)
    keyword and `target`{.variable} is an [ASCII
    case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#following-hyperlinks:ascii-case-insensitive
    x-internal="ascii-case-insensitive"} match for \"`_blank`\", then
    return true.]{#opener-processing-model}

3.  If `url`{.variable}\'s [blob URL
    entry](https://url.spec.whatwg.org/#concept-url-blob-entry){#following-hyperlinks:concept-url-blob-entry
    x-internal="concept-url-blob-entry"} is not null:

    1.  Let `blobOrigin`{.variable} be `url`{.variable}\'s [blob URL
        entry](https://url.spec.whatwg.org/#concept-url-blob-entry){#following-hyperlinks:concept-url-blob-entry-2
        x-internal="concept-url-blob-entry"}\'s
        [environment](https://w3c.github.io/FileAPI/#blob-url-entry-environment){#following-hyperlinks:blob-url-entry-environment
        x-internal="blob-url-entry-environment"}\'s
        [origin](webappapis.html#concept-settings-object-origin){#following-hyperlinks:concept-settings-object-origin}.

    2.  Let `topLevelOrigin`{.variable} be `element`{.variable}\'s
        [relevant settings
        object](webappapis.html#relevant-settings-object){#following-hyperlinks:relevant-settings-object}\'s
        [top-level
        origin](webappapis.html#concept-environment-top-level-origin){#following-hyperlinks:concept-environment-top-level-origin}.

    3.  If `blobOrigin`{.variable} is not [same
        site](browsers.html#same-site){#following-hyperlinks:same-site}
        with `topLevelOrigin`{.variable}, then return true.

4.  Return false.

To [follow the hyperlink]{#following-hyperlinks-2 .dfn} created by an
element `subject`{.variable}, given an optional
[`hyperlinkSuffix`{.variable}]{#following-hyperlinksuffix .dfn} (default
null) and an optional
[`userInvolvement`{.variable}]{#following-userinvolvement .dfn} (default
\"[`none`{#following-hyperlinks:uni-none}](browsing-the-web.html#uni-none)\"):

1.  If `subject`{.variable} [cannot
    navigate](#cannot-navigate){#following-hyperlinks:cannot-navigate},
    then return.

2.  Let `targetAttributeValue`{.variable} be the empty string.

3.  If `subject`{.variable} is an
    [`a`{#following-hyperlinks:the-a-element-4}](text-level-semantics.html#the-a-element)
    or
    [`area`{#following-hyperlinks:the-area-element-2}](image-maps.html#the-area-element)
    element, then set `targetAttributeValue`{.variable} to the result of
    [getting an element\'s
    target](semantics.html#get-an-element's-target){#following-hyperlinks:get-an-element's-target}
    given `subject`{.variable}.

4.  Let `urlRecord`{.variable} be the result of [encoding-parsing a
    URL](urls-and-fetching.html#encoding-parsing-a-url){#following-hyperlinks:encoding-parsing-a-url}
    given `subject`{.variable}\'s
    [`href`{#following-hyperlinks:attr-hyperlink-href}](#attr-hyperlink-href)
    attribute value, relative to `subject`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#following-hyperlinks:node-document-2
    x-internal="node-document"}.

5.  If `urlRecord`{.variable} is failure, then return.

6.  Let `noopener`{.variable} be the result of [getting an element\'s
    noopener](#get-an-element's-noopener){#following-hyperlinks:get-an-element's-noopener}
    with `subject`{.variable}, `urlRecord`{.variable}, and
    `targetAttributeValue`{.variable}.

7.  Let `targetNavigable`{.variable} be the first return value of
    applying [the rules for choosing a
    navigable](document-sequences.html#the-rules-for-choosing-a-navigable){#following-hyperlinks:the-rules-for-choosing-a-navigable}
    given `targetAttributeValue`{.variable}, `subject`{.variable}\'s
    [node
    navigable](document-sequences.html#node-navigable){#following-hyperlinks:node-navigable},
    and `noopener`{.variable}.

8.  If `targetNavigable`{.variable} is null, then return.

9.  Let `urlString`{.variable} be the result of applying the [URL
    serializer](https://url.spec.whatwg.org/#concept-url-serializer){#following-hyperlinks:concept-url-serializer
    x-internal="concept-url-serializer"} to `urlRecord`{.variable}.

10. If `hyperlinkSuffix`{.variable} is non-null, then append it to
    `urlString`{.variable}.

11. Let `referrerPolicy`{.variable} be the current state of
    `subject`{.variable}\'s `referrerpolicy` content attribute.

12. [If `subject`{.variable}\'s [link types](#linkTypes) includes the
    [`noreferrer`{#following-hyperlinks:link-type-noreferrer-2}](#link-type-noreferrer)
    keyword, then set `referrerPolicy`{.variable} to
    \"`no-referrer`\".]{#noreferrer-a-area-processing-model}

13. [Navigate](browsing-the-web.html#navigate){#following-hyperlinks:navigate}
    `targetNavigable`{.variable} to `urlString`{.variable} using
    `subject`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#following-hyperlinks:node-document-3
    x-internal="node-document"}, with
    *[referrerPolicy](browsing-the-web.html#navigation-referrer-policy)*
    set to `referrerPolicy`{.variable},
    *[userInvolvement](browsing-the-web.html#navigation-user-involvement)*
    set to `userInvolvement`{.variable}, and
    *[sourceElement](browsing-the-web.html#navigation-source-element)*
    set to `subject`{.variable}.

    Unlike many other types of navigations, following hyperlinks does
    not have special
    \"[`replace`{#following-hyperlinks:navigationhistorybehavior-replace}](browsing-the-web.html#navigationhistorybehavior-replace)\"
    behavior for when documents are not [completely
    loaded](document-lifecycle.html#completely-loaded){#following-hyperlinks:completely-loaded}.
    This is true for both user-initiated instances of following
    hyperlinks, as well as script-triggered ones via, e.g.,
    `aElement.click()`.

#### [4.6.5]{.secno} Downloading resources[](#downloading-resources){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[HTMLAnchorElement/download](https://developer.mozilla.org/en-US/docs/Web/API/HTMLAnchorElement/download "The HTMLAnchorElement.download property is a string indicating that the linked resource is intended to be downloaded rather than displayed in the browser. The value, if any, specifies the default file name for use in labeling the resource in a local file system. If the name is not a valid file name in the underlying OS, the browser will adjust it.")

Support in all current engines.

::: support
[Firefox20+]{.firefox .yes}[Safari10.1+]{.safari
.yes}[Chrome15+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)13+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

In some cases, resources are intended for later use rather than
immediate viewing. To indicate that a resource is intended to be
downloaded for use later, rather than immediately used, the
[`download`{#downloading-resources:attr-hyperlink-download}](#attr-hyperlink-download)
attribute can be specified on the
[`a`{#downloading-resources:the-a-element}](text-level-semantics.html#the-a-element)
or
[`area`{#downloading-resources:the-area-element}](image-maps.html#the-area-element)
element that creates the
[hyperlink](#hyperlink){#downloading-resources:hyperlink} to that
resource.

The attribute can furthermore be given a value, to specify the filename
that user agents are to use when storing the resource in a file system.
This value can be overridden by the
\`[`Content-Disposition`{#downloading-resources:http-content-disposition}](https://httpwg.org/specs/rfc6266.html){x-internal="http-content-disposition"}\`
HTTP header\'s filename parameters.
[\[RFC6266\]](references.html#refsRFC6266)

In cross-origin situations, the
[`download`{#downloading-resources:attr-hyperlink-download-2}](#attr-hyperlink-download)
attribute has to be combined with the
\`[`Content-Disposition`{#downloading-resources:http-content-disposition-2}](https://httpwg.org/specs/rfc6266.html){x-internal="http-content-disposition"}\`
HTTP header, specifically with the `attachment` disposition type, to
avoid the user being warned of possibly nefarious activity. (This is to
protect users from being made to download sensitive personal or
confidential information without their full understanding.)

------------------------------------------------------------------------

To [download the hyperlink]{#downloading-hyperlinks .dfn export=""}
created by an element `subject`{.variable}, given an optional
[`hyperlinkSuffix`{.variable}]{#downloading-hyperlinksuffix .dfn}
(default null) and an optional
[`userInvolvement`{.variable}]{#downloading-userinvolvement .dfn}
(default
\"[`none`{#downloading-resources:uni-none}](browsing-the-web.html#uni-none)\"):

1.  If `subject`{.variable} [cannot
    navigate](#cannot-navigate){#downloading-resources:cannot-navigate},
    then return.

2.  [If `subject`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#downloading-resources:node-document
    x-internal="node-document"}\'s [active sandboxing flag
    set](browsers.html#active-sandboxing-flag-set){#downloading-resources:active-sandboxing-flag-set}
    has the [sandboxed downloads browsing context
    flag](browsers.html#sandboxed-downloads-browsing-context-flag){#downloading-resources:sandboxed-downloads-browsing-context-flag}
    set, then return.]{#allowed-to-download}

3.  Let `urlString`{.variable} be the result of
    [encoding-parsing-and-serializing a
    URL](urls-and-fetching.html#encoding-parsing-and-serializing-a-url){#downloading-resources:encoding-parsing-and-serializing-a-url}
    given `subject`{.variable}\'s
    [`href`{#downloading-resources:attr-hyperlink-href}](#attr-hyperlink-href)
    attribute value, relative to `subject`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#downloading-resources:node-document-2
    x-internal="node-document"}.

4.  If `urlString`{.variable} is failure, then return.

5.  If `hyperlinkSuffix`{.variable} is non-null, then append it to
    `urlString`{.variable}.

6.  If `userInvolvement`{.variable} is not
    \"[`browser UI`{#downloading-resources:uni-browser-ui}](browsing-the-web.html#uni-browser-ui)\",
    then:

    1.  [Assert](https://infra.spec.whatwg.org/#assert){#downloading-resources:assert
        x-internal="assert"}: `subject`{.variable} has a
        [`download`{#downloading-resources:attr-hyperlink-download-3}](#attr-hyperlink-download)
        attribute.

    2.  Let `navigation`{.variable} be `subject`{.variable}\'s [relevant
        global
        object](webappapis.html#concept-relevant-global){#downloading-resources:concept-relevant-global}\'s
        [navigation
        API](nav-history-apis.html#window-navigation-api){#downloading-resources:window-navigation-api}.

    3.  Let `filename`{.variable} be the value of
        `subject`{.variable}\'s
        [`download`{#downloading-resources:attr-hyperlink-download-4}](#attr-hyperlink-download)
        attribute.

    4.  Let `continue`{.variable} be the result of [firing a download
        request `navigate`
        event](nav-history-apis.html#fire-a-download-request-navigate-event){#downloading-resources:fire-a-download-request-navigate-event}
        at `navigation`{.variable} with
        *[destinationURL](nav-history-apis.html#fire-navigate-download-destinationurl)*
        set to `urlString`{.variable},
        *[userInvolvement](nav-history-apis.html#fire-navigate-download-userinvolvement)*
        set to `userInvolvement`{.variable},
        *[sourceElement](nav-history-apis.html#fire-navigate-download-sourceelement)*
        set to `subject`{.variable}, and
        *[filename](nav-history-apis.html#fire-navigate-download-filename)*
        set to `filename`{.variable}.

    5.  If `continue`{.variable} is false, then return.

7.  Run these steps [in
    parallel](infrastructure.html#in-parallel){#downloading-resources:in-parallel}:

    1.  Optionally, the user agent may abort these steps, if it believes
        doing so would safeguard the user from a potentially hostile
        download.

    2.  Let `request`{.variable} be a new
        [request](https://fetch.spec.whatwg.org/#concept-request){#downloading-resources:concept-request
        x-internal="concept-request"} whose
        [URL](https://fetch.spec.whatwg.org/#concept-request-url){#downloading-resources:concept-request-url
        x-internal="concept-request-url"} is `urlString`{.variable},
        [client](https://fetch.spec.whatwg.org/#concept-request-client){#downloading-resources:concept-request-client
        x-internal="concept-request-client"} is [entry settings
        object](webappapis.html#entry-settings-object){#downloading-resources:entry-settings-object},
        [initiator](https://fetch.spec.whatwg.org/#concept-request-initiator){#downloading-resources:concept-request-initiator
        x-internal="concept-request-initiator"} is \"`download`\",
        [destination](https://fetch.spec.whatwg.org/#concept-request-destination){#downloading-resources:concept-request-destination
        x-internal="concept-request-destination"} is the empty string,
        and whose [synchronous
        flag](https://fetch.spec.whatwg.org/#synchronous-flag){#downloading-resources:synchronous-flag
        x-internal="synchronous-flag"} and [use-URL-credentials
        flag](https://fetch.spec.whatwg.org/#concept-request-use-url-credentials-flag){#downloading-resources:use-url-credentials-flag
        x-internal="use-url-credentials-flag"} are set.

    3.  [Handle as a
        download](#handle-as-a-download){#downloading-resources:handle-as-a-download}
        the result of
        [fetching](https://fetch.spec.whatwg.org/#concept-fetch){#downloading-resources:concept-fetch
        x-internal="concept-fetch"} `request`{.variable}.

To [handle as a download]{#handle-as-a-download .dfn} a
[response](https://fetch.spec.whatwg.org/#concept-response){#downloading-resources:concept-response
x-internal="concept-response"} `response`{.variable}:

1.  Let `suggestedFilename`{.variable} be the result of [getting the
    suggested
    filename](#getting-the-suggested-filename){#downloading-resources:getting-the-suggested-filename}
    for `response`{.variable}.

2.  Provide the user with a way to save `response`{.variable} for later
    use. If the user agent needs a filename, it should use
    `suggestedFilename`{.variable}. Report any problems downloading the
    file to the user.

3.  Return `suggestedFilename`{.variable}.

To [get the suggested filename]{#getting-the-suggested-filename .dfn}
for a
[response](https://fetch.spec.whatwg.org/#concept-response){#downloading-resources:concept-response-2
x-internal="concept-response"} `response`{.variable}:

This algorithm is intended to mitigate security dangers involved in
downloading files from untrusted sites, and user agents are strongly
urged to follow it.

1.  Let `filename`{.variable} be the undefined value.

2.  If `response`{.variable} has a
    \`[`Content-Disposition`{#downloading-resources:http-content-disposition-3}](https://httpwg.org/specs/rfc6266.html){x-internal="http-content-disposition"}\`
    header, that header specifies the `attachment` disposition type, and
    the header includes filename information, then let
    `filename`{.variable} have the value specified by the header, and
    jump to the step labeled *sanitize* below.
    [\[RFC6266\]](references.html#refsRFC6266)

3.  Let `interface origin`{.variable} be the
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#downloading-resources:concept-document-origin
    x-internal="concept-document-origin"} of the
    [`Document`{#downloading-resources:document}](dom.html#document) in
    which the
    [download](#downloading-hyperlinks){#downloading-resources:downloading-hyperlinks}
    or
    [navigate](browsing-the-web.html#navigate){#downloading-resources:navigate}
    action resulting in the download was initiated, if any.

4.  Let `response origin`{.variable} be the
    [origin](browsers.html#concept-origin){#downloading-resources:concept-origin}
    of the URL of `response`{.variable}, unless that URL\'s
    [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#downloading-resources:concept-url-scheme
    x-internal="concept-url-scheme"} component is `data`, in which case
    let `response origin`{.variable} be the same as the
    `interface origin`{.variable}, if any.

5.  If there is no `interface origin`{.variable}, then let
    `trusted operation`{.variable} be true. Otherwise, let
    `trusted operation`{.variable} be true if
    `response origin`{.variable} is the [same
    origin](browsers.html#same-origin){#downloading-resources:same-origin}
    as `interface origin`{.variable}, and false otherwise.

6.  If `trusted operation`{.variable} is true and `response`{.variable}
    has a
    \`[`Content-Disposition`{#downloading-resources:http-content-disposition-4}](https://httpwg.org/specs/rfc6266.html){x-internal="http-content-disposition"}\`
    header and that header includes filename information, then let
    `filename`{.variable} have the value specified by the header, and
    jump to the step labeled *sanitize* below.
    [\[RFC6266\]](references.html#refsRFC6266)

7.  If the download was not initiated from a
    [hyperlink](#hyperlink){#downloading-resources:hyperlink-2} created
    by an
    [`a`{#downloading-resources:the-a-element-2}](text-level-semantics.html#the-a-element)
    or
    [`area`{#downloading-resources:the-area-element-2}](image-maps.html#the-area-element)
    element, or if the element of the
    [hyperlink](#hyperlink){#downloading-resources:hyperlink-3} from
    which it was initiated did not have a
    [`download`{#downloading-resources:attr-hyperlink-download-5}](#attr-hyperlink-download)
    attribute when the download was initiated, or if there was such an
    attribute but its value when the download was initiated was the
    empty string, then jump to the step labeled *no proposed filename*.

8.  Let `proposed filename`{.variable} have the value of the
    [`download`{#downloading-resources:attr-hyperlink-download-6}](#attr-hyperlink-download)
    attribute of the element of the
    [hyperlink](#hyperlink){#downloading-resources:hyperlink-4} that
    initiated the download at the time the download was initiated.

9.  If `trusted operation`{.variable} is true, let `filename`{.variable}
    have the value of `proposed filename`{.variable}, and jump to the
    step labeled *sanitize* below.

10. If `response`{.variable} has a
    \`[`Content-Disposition`{#downloading-resources:http-content-disposition-5}](https://httpwg.org/specs/rfc6266.html){x-internal="http-content-disposition"}\`
    header and that header specifies the `attachment` disposition type,
    let `filename`{.variable} have the value of
    `proposed filename`{.variable}, and jump to the step labeled
    *sanitize* below. [\[RFC6266\]](references.html#refsRFC6266)

11. *No proposed filename*: If `trusted operation`{.variable} is true,
    or if the user indicated a preference for having the response in
    question downloaded, let `filename`{.variable} have a value derived
    from the
    [URL](https://url.spec.whatwg.org/#concept-url){#downloading-resources:url
    x-internal="url"} of `response`{.variable} in an
    [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#downloading-resources:implementation-defined
    x-internal="implementation-defined"} manner, and jump to the step
    labeled *sanitize* below.

12. Let `filename`{.variable} be set to the user\'s preferred filename
    or to a filename selected by the user agent, and jump to the step
    labeled *sanitize* below.

    ::: warning
    If the algorithm reaches this step, then a download was begun from a
    different origin than `response`{.variable}, and the origin did not
    mark the file as suitable for downloading, and the download was not
    initiated by the user. This could be because a
    [`download`{#downloading-resources:attr-hyperlink-download-7}](#attr-hyperlink-download)
    attribute was used to trigger the download, or because
    `response`{.variable} is not of a type that the user agent supports.

    This could be dangerous, because, for instance, a hostile server
    could be trying to get a user to unknowingly download private
    information and then re-upload it to the hostile server, by tricking
    the user into thinking the data is from the hostile server.

    Thus, it is in the user\'s interests that the user be somehow
    notified that `response`{.variable} comes from quite a different
    source, and to prevent confusion, any suggested filename from the
    potentially hostile `interface origin`{.variable} should be ignored.
    :::

13. *Sanitize*: Optionally, allow the user to influence
    `filename`{.variable}. For example, a user agent could prompt the
    user for a filename, potentially providing the value of
    `filename`{.variable} as determined above as a default value.

14. Adjust `filename`{.variable} to be suitable for the local file
    system.

    For example, this could involve removing characters that are not
    legal in filenames, or trimming leading and trailing whitespace.

15. If the platform conventions do not in any way use
    [extensions](#concept-extension){#downloading-resources:concept-extension}
    to determine the types of file on the file system, then return
    `filename`{.variable} as the filename.

16. Let `claimed type`{.variable} be the type given by
    `response`{.variable}\'s [Content-Type
    metadata](urls-and-fetching.html#content-type){#downloading-resources:content-type},
    if any is known. Let `named type`{.variable} be the type given by
    `filename`{.variable}\'s
    [extension](#concept-extension){#downloading-resources:concept-extension-2},
    if any is known. For the purposes of this step, a *type* is a
    mapping of a [MIME
    type](https://mimesniff.spec.whatwg.org/#mime-type){#downloading-resources:mime-type
    x-internal="mime-type"} to an
    [extension](#concept-extension){#downloading-resources:concept-extension-3}.

17. If `named type`{.variable} is consistent with the user\'s
    preferences (e.g., because the value of `filename`{.variable} was
    determined by prompting the user), then return `filename`{.variable}
    as the filename.

18. If `claimed type`{.variable} and `named type`{.variable} are the
    same type (i.e., the type given by `response`{.variable}\'s
    [Content-Type
    metadata](urls-and-fetching.html#content-type){#downloading-resources:content-type-2}
    is consistent with the type given by `filename`{.variable}\'s
    [extension](#concept-extension){#downloading-resources:concept-extension-4}),
    then return `filename`{.variable} as the filename.

19. If the `claimed type`{.variable} is known, then alter
    `filename`{.variable} to add an
    [extension](#concept-extension){#downloading-resources:concept-extension-5}
    corresponding to `claimed type`{.variable}.

    Otherwise, if `named type`{.variable} is known to be potentially
    dangerous (e.g. it will be treated by the platform conventions as a
    native executable, shell script, HTML application, or
    executable-macro-capable document), then optionally alter
    `filename`{.variable} to add a known-safe
    [extension](#concept-extension){#downloading-resources:concept-extension-6}
    (e.g. \"`.txt`\").

    This last step would make it impossible to download executables,
    which might not be desirable. As always, implementers are forced to
    balance security and usability in this matter.

20. Return `filename`{.variable} as the filename.

For the purposes of this algorithm, a file
[extension]{#concept-extension .dfn} consists of any part of the
filename that platform conventions dictate will be used for identifying
the type of the file. For example, many operating systems use the part
of the filename following the last dot (\"`.`\") in the filename to
determine the type of the file, and from that the manner in which the
file is to be opened or executed.

User agents should ignore any directory or path information provided by
the response itself, its
[URL](https://url.spec.whatwg.org/#concept-url){#downloading-resources:url-2
x-internal="url"}, and any
[`download`{#downloading-resources:attr-hyperlink-download-8}](#attr-hyperlink-download)
attribute, in deciding where to store the resulting file in the user\'s
file system.

#### [4.6.6]{.secno} [Hyperlink auditing]{.dfn}[](#hyperlink-auditing){.self-link}

If a [hyperlink](#hyperlink){#hyperlink-auditing:hyperlink} created by
an
[`a`{#hyperlink-auditing:the-a-element}](text-level-semantics.html#the-a-element)
or
[`area`{#hyperlink-auditing:the-area-element}](image-maps.html#the-area-element)
element has a [`ping`{#hyperlink-auditing:ping}](#ping) attribute, and
the user follows the hyperlink, and the value of the element\'s
[`href`{#hyperlink-auditing:attr-hyperlink-href}](#attr-hyperlink-href)
attribute can be
[parsed](urls-and-fetching.html#encoding-parsing-a-url){#hyperlink-auditing:encoding-parsing-a-url},
relative to the element\'s [node
document](https://dom.spec.whatwg.org/#concept-node-document){#hyperlink-auditing:node-document
x-internal="node-document"}, without failure, then the user agent must
take the [`ping`{#hyperlink-auditing:ping-2}](#ping) attribute\'s value,
[split that string on ASCII
whitespace](https://infra.spec.whatwg.org/#split-on-ascii-whitespace){#hyperlink-auditing:split-a-string-on-spaces
x-internal="split-a-string-on-spaces"},
[parse](urls-and-fetching.html#encoding-parsing-a-url){#hyperlink-auditing:encoding-parsing-a-url-2}
each resulting token, relative to the element\'s [node
document](https://dom.spec.whatwg.org/#concept-node-document){#hyperlink-auditing:node-document-2
x-internal="node-document"}, and then run these steps for each resulting
[URL](https://url.spec.whatwg.org/#concept-url){#hyperlink-auditing:url
x-internal="url"} `ping URL`{.variable}, ignoring when parsing returns
failure:

1.  If `ping URL`{.variable}\'s
    [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#hyperlink-auditing:concept-url-scheme
    x-internal="concept-url-scheme"} is not an [HTTP(S)
    scheme](https://fetch.spec.whatwg.org/#http-scheme){#hyperlink-auditing:http(s)-scheme
    x-internal="http(s)-scheme"}, then return.

2.  Optionally, return. (For example, the user agent might wish to
    ignore any or all ping URLs in accordance with the user\'s expressed
    preferences.)

3.  Let `settingsObject`{.variable} be the element\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#hyperlink-auditing:node-document-3
    x-internal="node-document"}\'s [relevant settings
    object](webappapis.html#relevant-settings-object){#hyperlink-auditing:relevant-settings-object}.

4.  Let `request`{.variable} be a new
    [request](https://fetch.spec.whatwg.org/#concept-request){#hyperlink-auditing:concept-request
    x-internal="concept-request"} whose
    [URL](https://fetch.spec.whatwg.org/#concept-request-url){#hyperlink-auditing:concept-request-url
    x-internal="concept-request-url"} is `ping URL`{.variable},
    [method](https://fetch.spec.whatwg.org/#concept-request-method){#hyperlink-auditing:concept-request-method
    x-internal="concept-request-method"} is \``POST`\`, [header
    list](https://fetch.spec.whatwg.org/#concept-request-header-list){#hyperlink-auditing:concept-request-header-list
    x-internal="concept-request-header-list"} is «
    (\`[`Content-Type`{#hyperlink-auditing:content-type}](urls-and-fetching.html#content-type)\`,
    \`[`text/ping`{#hyperlink-auditing:text/ping}](iana.html#text/ping)\`)
    »,
    [body](https://fetch.spec.whatwg.org/#concept-request-body){#hyperlink-auditing:concept-request-body
    x-internal="concept-request-body"} is \``PING`\`,
    [client](https://fetch.spec.whatwg.org/#concept-request-client){#hyperlink-auditing:concept-request-client
    x-internal="concept-request-client"} is `settingsObject`{.variable},
    [destination](https://fetch.spec.whatwg.org/#concept-request-destination){#hyperlink-auditing:concept-request-destination
    x-internal="concept-request-destination"} is the empty string,
    [credentials
    mode](https://fetch.spec.whatwg.org/#concept-request-credentials-mode){#hyperlink-auditing:concept-request-credentials-mode
    x-internal="concept-request-credentials-mode"} is \"`include`\",
    [referrer](https://fetch.spec.whatwg.org/#concept-request-referrer){#hyperlink-auditing:concept-request-referrer
    x-internal="concept-request-referrer"} is \"`no-referrer`\", and
    whose [use-URL-credentials
    flag](https://fetch.spec.whatwg.org/#concept-request-use-url-credentials-flag){#hyperlink-auditing:use-url-credentials-flag
    x-internal="use-url-credentials-flag"} is set, and whose [initiator
    type](https://fetch.spec.whatwg.org/#request-initiator-type){#hyperlink-auditing:concept-request-initiator-type
    x-internal="concept-request-initiator-type"} is \"`ping`\".

5.  Let `target URL`{.variable} be the result of
    [encoding-parsing-and-serializing a
    URL](urls-and-fetching.html#encoding-parsing-and-serializing-a-url){#hyperlink-auditing:encoding-parsing-and-serializing-a-url}
    given the element\'s
    [`href`{#hyperlink-auditing:attr-hyperlink-href-2}](#attr-hyperlink-href)
    attribute\'s value, relative to the element\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#hyperlink-auditing:node-document-4
    x-internal="node-document"}, and then:

    If the [URL](https://dom.spec.whatwg.org/#concept-document-url){#hyperlink-auditing:the-document's-address x-internal="the-document's-address"} of the [`Document`{#hyperlink-auditing:document}](dom.html#document) object containing the hyperlink being audited and `ping URL`{.variable} have the [same origin](browsers.html#same-origin){#hyperlink-auditing:same-origin}\
    If the origins are different, but the [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#hyperlink-auditing:concept-url-scheme-2 x-internal="concept-url-scheme"} of the [URL](https://dom.spec.whatwg.org/#concept-document-url){#hyperlink-auditing:the-document's-address-2 x-internal="the-document's-address"} of the [`Document`{#hyperlink-auditing:document-2}](dom.html#document) containing the hyperlink being audited is not \"`https`\"
    :   `request`{.variable} must include a
        \`[`Ping-From`{#hyperlink-auditing:ping-from}](#ping-from)\`
        header with, as its value, the
        [URL](https://dom.spec.whatwg.org/#concept-document-url){#hyperlink-auditing:the-document's-address-3
        x-internal="the-document's-address"} of the document containing
        the hyperlink, and a
        \`[`Ping-To`{#hyperlink-auditing:ping-to}](#ping-to)\` HTTP
        header with, as its value, the `target URL`{.variable}.

    Otherwise
    :   `request`{.variable} must include a
        \`[`Ping-To`{#hyperlink-auditing:ping-to-2}](#ping-to)\` HTTP
        header with, as its value, `target URL`{.variable}.
        [`request`{.variable} does not include a
        \`[`Ping-From`{#hyperlink-auditing:ping-from-2}](#ping-from)\`
        header.]{.note}

6.  [Fetch](https://fetch.spec.whatwg.org/#concept-fetch){#hyperlink-auditing:concept-fetch
    x-internal="concept-fetch"} `request`{.variable}.

This may be done [in
parallel](infrastructure.html#in-parallel){#hyperlink-auditing:in-parallel}
with the primary fetch, and is independent of the result of that fetch.

User agents should allow the user to adjust this behavior, for example
in conjunction with a setting that disables the sending of HTTP
\`[`Referer`{#hyperlink-auditing:http-referer}](https://httpwg.org/specs/rfc7231.html#header.referer){x-internal="http-referer"}\`
(sic) headers. Based on the user\'s preferences, UAs may either
[ignore](infrastructure.html#ignore){#hyperlink-auditing:ignore} the
[`ping`{#hyperlink-auditing:ping-3}](#ping) attribute altogether, or
selectively ignore URLs in the list (e.g. ignoring any third-party
URLs); this is explicitly accounted for in the steps above.

User agents must ignore any entity bodies returned in the responses.
User agents may close the connection prematurely once they start
receiving a response body.

When the [`ping`{#hyperlink-auditing:ping-4}](#ping) attribute is
present, user agents should clearly indicate to the user that following
the hyperlink will also cause secondary requests to be sent in the
background, possibly including listing the actual target URLs.

For example, a visual user agent could include the hostnames of the
target ping URLs along with the hyperlink\'s actual URL in a status bar
or tooltip.

::: note
The [`ping`{#hyperlink-auditing:ping-5}](#ping) attribute is redundant
with pre-existing technologies like HTTP redirects and JavaScript in
allowing web pages to track which off-site links are most popular or
allowing advertisers to track click-through rates.

However, the [`ping`{#hyperlink-auditing:ping-6}](#ping) attribute
provides these advantages to the user over those alternatives:

- It allows the user to see the final target URL unobscured.
- It allows the UA to inform the user about the out-of-band
  notifications.
- It allows the user to disable the notifications without losing the
  underlying link functionality.
- It allows the UA to optimize the use of available network bandwidth so
  that the target page loads faster.

Thus, while it is possible to track users without this feature, authors
are encouraged to use the [`ping`{#hyperlink-auditing:ping-7}](#ping)
attribute so that the user agent can make the user experience more
transparent.
:::

##### [4.6.6.1]{.secno} The \`[`Ping-From`{#the-ping-headers:ping-from}](#ping-from)\` and \`[`Ping-To`{#the-ping-headers:ping-to}](#ping-to)\` headers[](#the-ping-headers){.self-link} {#the-ping-headers}

The \`[`Ping-From`]{#ping-from .dfn dfn-type="http-header"}\` and
\`[`Ping-To`]{#ping-to .dfn dfn-type="http-header"}\` HTTP request
headers are included in [hyperlink
auditing](#hyperlink-auditing){#the-ping-headers:hyperlink-auditing}
requests. Their value is a
[URL](https://url.spec.whatwg.org/#concept-url){#the-ping-headers:url
x-internal="url"},
[serialized](https://url.spec.whatwg.org/#concept-url-serializer){#the-ping-headers:concept-url-serializer
x-internal="concept-url-serializer"}.

#### [4.6.7]{.secno} Link types[](#linkTypes){.self-link} {#linkTypes}

:::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Link_types](https://developer.mozilla.org/en-US/docs/Web/HTML/Link_types "The rel attribute defines the relationship between a linked resource and the current document. Valid on <link>, <a>, <area>, and <form>, the supported values depend on the element on which the attribute is found.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera9+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerYes]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::

::: feature
[Link_types](https://developer.mozilla.org/en-US/docs/Web/HTML/Link_types "The rel attribute defines the relationship between a linked resource and the current document. Valid on <link>, <a>, <area>, and <form>, the supported values depend on the element on which the attribute is found.")
:::
::::::

The following table summarizes the link types that are defined by this
specification, by their corresponding keywords. This table is
non-normative; the actual definitions for the link types are given in
the next few sections.

In this section, the term *referenced document* refers to the resource
identified by the element representing the link, and the term *current
document* refers to the resource within which the element representing
the link finds itself.

To determine which link types apply to a
[`link`{#linkTypes:the-link-element}](semantics.html#the-link-element),
[`a`{#linkTypes:the-a-element}](text-level-semantics.html#the-a-element),
[`area`{#linkTypes:the-area-element}](image-maps.html#the-area-element),
or [`form`{#linkTypes:the-form-element}](forms.html#the-form-element)
element, the element\'s `rel` attribute must be [split on ASCII
whitespace](https://infra.spec.whatwg.org/#split-on-ascii-whitespace){#linkTypes:split-a-string-on-spaces
x-internal="split-a-string-on-spaces"}. The resulting tokens are the
keywords for the link types that apply to that element.

Except where otherwise specified, a keyword must not be specified more
than once per
[`rel`{#linkTypes:attr-hyperlink-rel}](#attr-hyperlink-rel) attribute.

Some of the sections that follow the table below list synonyms for
certain keywords. The indicated synonyms are to be handled as specified
by user agents, but must not be used in documents (for example, the
keyword \"`copyright`\").

Keywords are always [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#linkTypes:ascii-case-insensitive
x-internal="ascii-case-insensitive"}, and must be compared as such.

Thus, `rel="next"` is the same as `rel="NEXT"`.

Keywords that are [body-ok]{#body-ok .dfn} affect whether
[`link`{#linkTypes:the-link-element-2}](semantics.html#the-link-element)
elements are [allowed in the
body](semantics.html#allowed-in-the-body){#linkTypes:allowed-in-the-body}.
The [body-ok](#body-ok){#linkTypes:body-ok} keywords are
[`dns-prefetch`{#linkTypes:link-type-dns-prefetch}](#link-type-dns-prefetch),
[`modulepreload`{#linkTypes:link-type-modulepreload}](#link-type-modulepreload),
[`pingback`{#linkTypes:link-type-pingback}](#link-type-pingback),
[`preconnect`{#linkTypes:link-type-preconnect}](#link-type-preconnect),
[`prefetch`{#linkTypes:link-type-prefetch}](#link-type-prefetch),
[`preload`{#linkTypes:link-type-preload}](#link-type-preload), and
[`stylesheet`{#linkTypes:link-type-stylesheet}](#link-type-stylesheet).

New link types that are to be implemented by web browsers are to be
added to this standard. The remainder can be [registered as
extensions](#concept-rel-extensions){#linkTypes:concept-rel-extensions}.

Link type

Effect on\...

[body-ok](#body-ok){#linkTypes:body-ok-2}

Has
\`[`Link`{#linkTypes:http-link}](https://httpwg.org/specs/rfc8288.html#header){x-internal="http-link"}\`
processing

Brief description

[`link`{#linkTypes:the-link-element-3}](semantics.html#the-link-element)

[`a`{#linkTypes:the-a-element-2}](text-level-semantics.html#the-a-element)
and
[`area`{#linkTypes:the-area-element-2}](image-maps.html#the-area-element)

[`form`{#linkTypes:the-form-element-2}](forms.html#the-form-element)

[`alternate`{#linkTypes:rel-alternate}](#rel-alternate)

[Hyperlink](#hyperlink){#linkTypes:hyperlink}

*not allowed*

·

·

Gives alternate representations of the current document.

[`canonical`{#linkTypes:link-type-canonical}](#link-type-canonical)

[Hyperlink](#hyperlink){#linkTypes:hyperlink-2}

*not allowed*

·

·

Gives the preferred URL for the current document.

[`author`{#linkTypes:link-type-author}](#link-type-author)

[Hyperlink](#hyperlink){#linkTypes:hyperlink-3}

*not allowed*

·

·

Gives a link to the author of the current document or article.

[`bookmark`{#linkTypes:link-type-bookmark}](#link-type-bookmark)

*not allowed*

[Hyperlink](#hyperlink){#linkTypes:hyperlink-4}

*not allowed*

·

·

Gives the permalink for the nearest ancestor section.

[`dns-prefetch`{#linkTypes:link-type-dns-prefetch-2}](#link-type-dns-prefetch)

[External
Resource](#external-resource-link){#linkTypes:external-resource-link}

*not allowed*

Yes

·

Specifies that the user agent should preemptively perform DNS resolution
for the target resource\'s
[origin](browsers.html#concept-origin){#linkTypes:concept-origin}.

[`expect`{#linkTypes:link-type-expect}](#link-type-expect)

[Internal
Resource](#internal-resource-link){#linkTypes:internal-resource-link}

*not allowed*

·

·

Expect an element with the target ID to appear in the current document.

[`external`{#linkTypes:link-type-external}](#link-type-external)

*not allowed*

[Annotation](#hyperlink-annotation){#linkTypes:hyperlink-annotation}

·

·

Indicates that the referenced document is not part of the same site as
the current document.

[`help`{#linkTypes:link-type-help}](#link-type-help)

[Hyperlink](#hyperlink){#linkTypes:hyperlink-5}

·

·

Provides a link to context-sensitive help.

[`icon`{#linkTypes:rel-icon}](#rel-icon)

[External
Resource](#external-resource-link){#linkTypes:external-resource-link-2}

*not allowed*

·

·

Imports an icon to represent the current document.

[`manifest`{#linkTypes:link-type-manifest}](#link-type-manifest)

[External
Resource](#external-resource-link){#linkTypes:external-resource-link-3}

*not allowed*

·

·

Imports or links to an [application
manifest](https://w3c.github.io/manifest/#dfn-manifest){#linkTypes:application-manifest
x-internal="application-manifest"}.
[\[MANIFEST\]](references.html#refsMANIFEST)

[`modulepreload`{#linkTypes:link-type-modulepreload-2}](#link-type-modulepreload)

[External
Resource](#external-resource-link){#linkTypes:external-resource-link-4}

*not allowed*

Yes

·

Specifies that the user agent must preemptively [fetch the module
script](webappapis.html#fetch-a-single-module-script){#linkTypes:fetch-a-single-module-script}
and store it in the document\'s [module
map](dom.html#concept-document-module-map){#linkTypes:concept-document-module-map}
for later evaluation. Optionally, the module\'s dependencies can be
fetched as well.

[`license`{#linkTypes:link-type-license}](#link-type-license)

[Hyperlink](#hyperlink){#linkTypes:hyperlink-6}

·

·

Indicates that the main content of the current document is covered by
the copyright license described by the referenced document.

[`next`{#linkTypes:link-type-next}](#link-type-next)

[Hyperlink](#hyperlink){#linkTypes:hyperlink-7}

·

·

Indicates that the current document is a part of a series, and that the
next document in the series is the referenced document.

[`nofollow`{#linkTypes:link-type-nofollow}](#link-type-nofollow)

*not allowed*

[Annotation](#hyperlink-annotation){#linkTypes:hyperlink-annotation-2}

·

·

Indicates that the current document\'s original author or publisher does
not endorse the referenced document.

[`noopener`{#linkTypes:link-type-noopener}](#link-type-noopener)

*not allowed*

[Annotation](#hyperlink-annotation){#linkTypes:hyperlink-annotation-3}

·

·

Creates a [top-level
traversable](document-sequences.html#top-level-traversable){#linkTypes:top-level-traversable}
with a non-[auxiliary browsing
context](document-sequences.html#auxiliary-browsing-context){#linkTypes:auxiliary-browsing-context}
if the hyperlink would otherwise create one that was auxiliary (i.e.,
has an appropriate
[`target`{#linkTypes:attr-hyperlink-target}](#attr-hyperlink-target)
attribute value).

[`noreferrer`{#linkTypes:link-type-noreferrer}](#link-type-noreferrer)

*not allowed*

[Annotation](#hyperlink-annotation){#linkTypes:hyperlink-annotation-4}

·

·

No
\`[`Referer`{#linkTypes:http-referer}](https://httpwg.org/specs/rfc7231.html#header.referer){x-internal="http-referer"}\`
(sic) header will be included. Additionally, has the same effect as
[`noopener`{#linkTypes:link-type-noopener-2}](#link-type-noopener).

[`opener`{#linkTypes:link-type-opener}](#link-type-opener)

*not allowed*

[Annotation](#hyperlink-annotation){#linkTypes:hyperlink-annotation-5}

·

·

Creates an [auxiliary browsing
context](document-sequences.html#auxiliary-browsing-context){#linkTypes:auxiliary-browsing-context-2}
if the hyperlink would otherwise create a [top-level
traversable](document-sequences.html#top-level-traversable){#linkTypes:top-level-traversable-2}
with a non-[auxiliary browsing
context](document-sequences.html#auxiliary-browsing-context){#linkTypes:auxiliary-browsing-context-3}
(i.e., has \"`_blank`\" as
[`target`{#linkTypes:attr-hyperlink-target-2}](#attr-hyperlink-target)
attribute value).

[`pingback`{#linkTypes:link-type-pingback-2}](#link-type-pingback)

[External
Resource](#external-resource-link){#linkTypes:external-resource-link-5}

*not allowed*

Yes

·

Gives the address of the pingback server that handles pingbacks to the
current document.

[`preconnect`{#linkTypes:link-type-preconnect-2}](#link-type-preconnect)

[External
Resource](#external-resource-link){#linkTypes:external-resource-link-6}

*not allowed*

Yes

Yes

Specifies that the user agent should preemptively connect to the target
resource\'s
[origin](browsers.html#concept-origin){#linkTypes:concept-origin-2}.

[`prefetch`{#linkTypes:link-type-prefetch-2}](#link-type-prefetch)

[External
Resource](#external-resource-link){#linkTypes:external-resource-link-7}

*not allowed*

Yes

·

Specifies that the user agent should preemptively
[fetch](https://fetch.spec.whatwg.org/#concept-fetch){#linkTypes:concept-fetch
x-internal="concept-fetch"} and cache the target resource as it is
likely to be required for a followup
[navigation](browsing-the-web.html#navigate){#linkTypes:navigate}.

[`preload`{#linkTypes:link-type-preload-2}](#link-type-preload)

[External
Resource](#external-resource-link){#linkTypes:external-resource-link-8}

*not allowed*

Yes

Yes

Specifies that the user agent must preemptively
[fetch](https://fetch.spec.whatwg.org/#concept-fetch){#linkTypes:concept-fetch-2
x-internal="concept-fetch"} and cache the target resource for current
[navigation](browsing-the-web.html#navigate){#linkTypes:navigate-2}
according to the [potential
destination](https://fetch.spec.whatwg.org/#concept-potential-destination){#linkTypes:concept-potential-destination
x-internal="concept-potential-destination"} given by the
[`as`{#linkTypes:attr-link-as}](semantics.html#attr-link-as) attribute
(and the
[priority](https://fetch.spec.whatwg.org/#request-priority){#linkTypes:concept-request-priority
x-internal="concept-request-priority"} associated with the
[corresponding](https://fetch.spec.whatwg.org/#concept-potential-destination-translate){#linkTypes:concept-potential-destination-translate
x-internal="concept-potential-destination-translate"}
[destination](https://fetch.spec.whatwg.org/#concept-request-destination){#linkTypes:concept-request-destination
x-internal="concept-request-destination"}).

[`prev`{#linkTypes:link-type-prev}](#link-type-prev)

[Hyperlink](#hyperlink){#linkTypes:hyperlink-8}

·

·

Indicates that the current document is a part of a series, and that the
previous document in the series is the referenced document.

[`privacy-policy`{#linkTypes:link-type-privacy-policy}](#link-type-privacy-policy)

[Hyperlink](#hyperlink){#linkTypes:hyperlink-9}

*not allowed*

·

·

Gives a link to information about the data collection and usage
practices that apply to the current document.

[`search`{#linkTypes:link-type-search}](#link-type-search)

[Hyperlink](#hyperlink){#linkTypes:hyperlink-10}

·

·

Gives a link to a resource that can be used to search through the
current document and its related pages.

[`stylesheet`{#linkTypes:link-type-stylesheet-2}](#link-type-stylesheet)

[External
Resource](#external-resource-link){#linkTypes:external-resource-link-9}

*not allowed*

Yes

·

Imports a style sheet.

[`tag`{#linkTypes:link-type-tag}](#link-type-tag)

*not allowed*

[Hyperlink](#hyperlink){#linkTypes:hyperlink-11}

*not allowed*

·

·

Gives a tag (identified by the given address) that applies to the
current document.

[`terms-of-service`{#linkTypes:link-type-terms-of-service}](#link-type-terms-of-service)

[Hyperlink](#hyperlink){#linkTypes:hyperlink-12}

*not allowed*

·

·

Gives a link to information about the agreements between the current
document\'s provider and users who wish to use the current document.

##### [4.6.7.1]{.secno} []{#link-type-alternate}Link type \"[`alternate`]{.dfn dfn-for="link/rel,a/rel,area/rel" dfn-type="attr-value"}\"[](#rel-alternate){.self-link} {#rel-alternate}

::::: {.mdn-anno .wrapped}
**⚠**MDN

:::: feature
[Alternative_style_sheets](https://developer.mozilla.org/en-US/docs/Web/CSS/Alternative_style_sheets "Specifying alternative style sheets in a web page provides a way for users to see multiple versions of a page, based on their needs or preferences.")

Support in one engine only.

::: support
[Firefox3+]{.firefox .yes}[Safari?]{.safari
.unknown}[Chrome1--48]{.chrome .no}

------------------------------------------------------------------------

[OperaYes]{.opera .yes}[EdgeNo]{.edge_blink .no}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet Explorer8+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

The [`alternate`{#rel-alternate:rel-alternate}](#rel-alternate) keyword
may be used with
[`link`{#rel-alternate:the-link-element}](semantics.html#the-link-element),
[`a`{#rel-alternate:the-a-element}](text-level-semantics.html#the-a-element),
and
[`area`{#rel-alternate:the-area-element}](image-maps.html#the-area-element)
elements.

The meaning of this keyword depends on the values of the other
attributes.

If the element is a [`link`{#rel-alternate:the-link-element-2}](semantics.html#the-link-element) element and the [`rel`{#rel-alternate:attr-link-rel}](semantics.html#attr-link-rel) attribute also contains the keyword [`stylesheet`{#rel-alternate:link-type-stylesheet}](#link-type-stylesheet)

:   The [`alternate`{#rel-alternate:rel-alternate-2}](#rel-alternate)
    keyword modifies the meaning of the
    [`stylesheet`{#rel-alternate:link-type-stylesheet-2}](#link-type-stylesheet)
    keyword in the way described for that keyword. The
    [`alternate`{#rel-alternate:rel-alternate-3}](#rel-alternate)
    keyword does not create a link of its own.

    ::: example
    Here, a set of
    [`link`{#rel-alternate:the-link-element-3}](semantics.html#the-link-element)
    elements provide some style sheets:

    ``` html
    <!-- a persistent style sheet -->
    <link rel="stylesheet" href="default.css">

    <!-- the preferred alternate style sheet -->
    <link rel="stylesheet" href="green.css" title="Green styles">

    <!-- some alternate style sheets -->
    <link rel="alternate stylesheet" href="contrast.css" title="High contrast">
    <link rel="alternate stylesheet" href="big.css" title="Big fonts">
    <link rel="alternate stylesheet" href="wide.css" title="Wide screen">
    ```
    :::

If the [`alternate`{#rel-alternate:rel-alternate-4}](#rel-alternate) keyword is used with the [`type`{#rel-alternate:attr-hyperlink-type}](#attr-hyperlink-type) attribute set to the value `application/rss+xml` or the value `application/atom+xml`

:   The keyword creates a
    [hyperlink](#hyperlink){#rel-alternate:hyperlink} referencing a
    syndication feed (though not necessarily syndicating exactly the
    same content as the current page).

    For the purposes of feed autodiscovery, user agents should consider
    all
    [`link`{#rel-alternate:the-link-element-4}](semantics.html#the-link-element)
    elements in the document with the
    [`alternate`{#rel-alternate:rel-alternate-5}](#rel-alternate)
    keyword used and with their
    [`type`{#rel-alternate:attr-hyperlink-type-2}](#attr-hyperlink-type)
    attribute set to the value `application/rss+xml` or the value
    `application/atom+xml`. If the user agent has the concept of a
    default syndication feed, the first such element (in [tree
    order](https://dom.spec.whatwg.org/#concept-tree-order){#rel-alternate:tree-order
    x-internal="tree-order"}) should be used as the default.

    ::: example
    The following
    [`link`{#rel-alternate:the-link-element-5}](semantics.html#the-link-element)
    elements give syndication feeds for a blog:

    ``` html
    <link rel="alternate" type="application/atom+xml" href="posts.xml" title="Cool Stuff Blog">
    <link rel="alternate" type="application/atom+xml" href="posts.xml?category=robots" title="Cool Stuff Blog: robots category">
    <link rel="alternate" type="application/atom+xml" href="comments.xml" title="Cool Stuff Blog: Comments">
    ```

    Such
    [`link`{#rel-alternate:the-link-element-6}](semantics.html#the-link-element)
    elements would be used by user agents engaged in feed autodiscovery,
    with the first being the default (where applicable).

    The following example offers various different syndication feeds to
    the user, using
    [`a`{#rel-alternate:the-a-element-2}](text-level-semantics.html#the-a-element)
    elements:

    ``` html
    <p>You can access the planets database using Atom feeds:</p>
    <ul>
     <li><a href="recently-visited-planets.xml" rel="alternate" type="application/atom+xml">Recently Visited Planets</a></li>
     <li><a href="known-bad-planets.xml" rel="alternate" type="application/atom+xml">Known Bad Planets</a></li>
     <li><a href="unexplored-planets.xml" rel="alternate" type="application/atom+xml">Unexplored Planets</a></li>
    </ul>
    ```

    These links would not be used in feed autodiscovery.
    :::

Otherwise

:   The keyword creates a
    [hyperlink](#hyperlink){#rel-alternate:hyperlink-2} referencing an
    alternate representation of the current document.

    The nature of the referenced document is given by the
    [`hreflang`{#rel-alternate:attr-hyperlink-hreflang}](#attr-hyperlink-hreflang),
    and
    [`type`{#rel-alternate:attr-hyperlink-type-3}](#attr-hyperlink-type)
    attributes.

    If the [`alternate`{#rel-alternate:rel-alternate-6}](#rel-alternate)
    keyword is used with the
    [`hreflang`{#rel-alternate:attr-hyperlink-hreflang-2}](#attr-hyperlink-hreflang)
    attribute, and that attribute\'s value differs from the [document
    element](https://dom.spec.whatwg.org/#document-element){#rel-alternate:document-element
    x-internal="document-element"}\'s
    [language](dom.html#language){#rel-alternate:language}, it indicates
    that the referenced document is a translation.

    If the [`alternate`{#rel-alternate:rel-alternate-7}](#rel-alternate)
    keyword is used with the
    [`type`{#rel-alternate:attr-hyperlink-type-4}](#attr-hyperlink-type)
    attribute, it indicates that the referenced document is a
    reformulation of the current document in the specified format.

    The
    [`hreflang`{#rel-alternate:attr-hyperlink-hreflang-3}](#attr-hyperlink-hreflang)
    and
    [`type`{#rel-alternate:attr-hyperlink-type-5}](#attr-hyperlink-type)
    attributes can be combined when specified with the
    [`alternate`{#rel-alternate:rel-alternate-8}](#rel-alternate)
    keyword.

    ::: example
    The following example shows how you can specify versions of the page
    that use alternative formats, are aimed at other languages, and that
    are intended for other media:

    ``` html
    <link rel=alternate href="/en/html" hreflang=en type=text/html title="English HTML">
    <link rel=alternate href="/fr/html" hreflang=fr type=text/html title="French HTML">
    <link rel=alternate href="/en/html/print" hreflang=en type=text/html media=print title="English HTML (for printing)">
    <link rel=alternate href="/fr/html/print" hreflang=fr type=text/html media=print title="French HTML (for printing)">
    <link rel=alternate href="/en/pdf" hreflang=en type=application/pdf title="English PDF">
    <link rel=alternate href="/fr/pdf" hreflang=fr type=application/pdf title="French PDF">
    ```
    :::

    This relationship is transitive --- that is, if a document links to
    two other documents with the link type
    \"[`alternate`{#rel-alternate:rel-alternate-9}](#rel-alternate)\",
    then, in addition to implying that those documents are alternative
    representations of the first document, it is also implying that
    those two documents are alternative representations of each other.

##### [4.6.7.2]{.secno} Link type \"[`author`]{.dfn dfn-for="link/rel,a/rel,area/rel" dfn-type="attr-value"}\"[](#link-type-author){.self-link}

The [`author`{#link-type-author:link-type-author}](#link-type-author)
keyword may be used with
[`link`{#link-type-author:the-link-element}](semantics.html#the-link-element),
[`a`{#link-type-author:the-a-element}](text-level-semantics.html#the-a-element),
and
[`area`{#link-type-author:the-area-element}](image-maps.html#the-area-element)
elements. This keyword creates a
[hyperlink](#hyperlink){#link-type-author:hyperlink}.

For
[`a`{#link-type-author:the-a-element-2}](text-level-semantics.html#the-a-element)
and
[`area`{#link-type-author:the-area-element-2}](image-maps.html#the-area-element)
elements, the
[`author`{#link-type-author:link-type-author-2}](#link-type-author)
keyword indicates that the referenced document provides further
information about the author of the nearest
[`article`{#link-type-author:the-article-element}](sections.html#the-article-element)
element ancestor of the element defining the hyperlink, if there is one,
or of the page as a whole, otherwise.

For
[`link`{#link-type-author:the-link-element-2}](semantics.html#the-link-element)
elements, the
[`author`{#link-type-author:link-type-author-3}](#link-type-author)
keyword indicates that the referenced document provides further
information about the author for the page as a whole.

The \"referenced document\" can be, and often is, a
[`mailto:`{#link-type-author:mailto-protocol}](https://www.rfc-editor.org/rfc/rfc6068#section-2){x-internal="mailto-protocol"}
URL giving the email address of the author.
[\[MAILTO\]](references.html#refsMAILTO)

**Synonyms**: For historical reasons, user agents must also treat
[`link`{#link-type-author:the-link-element-3}](semantics.html#the-link-element),
[`a`{#link-type-author:the-a-element-3}](text-level-semantics.html#the-a-element),
and
[`area`{#link-type-author:the-area-element-3}](image-maps.html#the-area-element)
elements that have a `rev` attribute with the value \"`made`\" as having
the [`author`{#link-type-author:link-type-author-4}](#link-type-author)
keyword specified as a link relationship.

##### [4.6.7.3]{.secno} Link type \"[`bookmark`]{.dfn dfn-for="a/rel,area/rel" dfn-type="attr-value"}\"[](#link-type-bookmark){.self-link}

The
[`bookmark`{#link-type-bookmark:link-type-bookmark}](#link-type-bookmark)
keyword may be used with
[`a`{#link-type-bookmark:the-a-element}](text-level-semantics.html#the-a-element)
and
[`area`{#link-type-bookmark:the-area-element}](image-maps.html#the-area-element)
elements. This keyword creates a
[hyperlink](#hyperlink){#link-type-bookmark:hyperlink}.

The
[`bookmark`{#link-type-bookmark:link-type-bookmark-2}](#link-type-bookmark)
keyword gives a permalink for the nearest ancestor
[`article`{#link-type-bookmark:the-article-element}](sections.html#the-article-element)
element of the linking element in question, or of the section the
linking element is most closely associated with, if there are no
ancestor
[`article`{#link-type-bookmark:the-article-element-2}](sections.html#the-article-element)
elements.

::: example
The following snippet has three permalinks. A user agent could determine
which permalink applies to which part of the spec by looking at where
the permalinks are given.

``` html
 ...
 <body>
  <h1>Example of permalinks</h1>
  <div id="a">
   <h2>First example</h2>
   <p><a href="a.html" rel="bookmark">This permalink applies to
   only the content from the first H2 to the second H2</a>. The DIV isn't
   exactly that section, but it roughly corresponds to it.</p>
  </div>
  <h2>Second example</h2>
  <article id="b">
   <p><a href="b.html" rel="bookmark">This permalink applies to
   the outer ARTICLE element</a> (which could be, e.g., a blog post).</p>
   <article id="c">
    <p><a href="c.html" rel="bookmark">This permalink applies to
    the inner ARTICLE element</a> (which could be, e.g., a blog comment).</p>
   </article>
  </article>
 </body>
 ...
```
:::

##### [4.6.7.4]{.secno} Link type \"[`canonical`]{.dfn dfn-for="link/rel" dfn-type="attr-value"}\"[](#link-type-canonical){.self-link}

The
[`canonical`{#link-type-canonical:link-type-canonical}](#link-type-canonical)
keyword may be used with
[`link`{#link-type-canonical:the-link-element}](semantics.html#the-link-element)
element. This keyword creates a
[hyperlink](#hyperlink){#link-type-canonical:hyperlink}.

The
[`canonical`{#link-type-canonical:link-type-canonical-2}](#link-type-canonical)
keyword indicates that URL given by the
[`href`{#link-type-canonical:attr-link-href}](semantics.html#attr-link-href)
attribute is the preferred URL for the current document. That helps
search engines reduce duplicate content, as described in more detail in
The Canonical Link Relation. [\[RFC6596\]](references.html#refsRFC6596)

##### [4.6.7.5]{.secno} Link type \"[`dns-prefetch`]{.dfn dfn-for="link/rel" dfn-type="attr-value"}\"[](#link-type-dns-prefetch){.self-link}

::::: {.mdn-anno .wrapped}
MDN

:::: feature
[Link_types/dns-prefetch](https://developer.mozilla.org/en-US/docs/Web/HTML/Link_types/dns-prefetch "The dns-prefetch keyword for the rel attribute of the <link> element is a hint to browsers that the user is likely to need resources from the target resource's origin, and therefore the browser can likely improve the user experience by preemptively performing DNS resolution for that origin.")

::: support
[Firefox3+]{.firefox .yes}[Safari?]{.safari .unknown}[Chrome46+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)No]{.edge .no}[Internet Explorer?]{.ie .unknown}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome AndroidYes]{.chrome_android .yes}[WebView
Android46+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

The
[`dns-prefetch`{#link-type-dns-prefetch:link-type-dns-prefetch}](#link-type-dns-prefetch)
keyword may be used with
[`link`{#link-type-dns-prefetch:the-link-element}](semantics.html#the-link-element)
elements. This keyword creates an [external resource
link](#external-resource-link){#link-type-dns-prefetch:external-resource-link}.
This keyword is [body-ok](#body-ok){#link-type-dns-prefetch:body-ok}.

The
[`dns-prefetch`{#link-type-dns-prefetch:link-type-dns-prefetch-2}](#link-type-dns-prefetch)
keyword indicates that preemptively performing DNS resolution for the
[origin](browsers.html#concept-origin){#link-type-dns-prefetch:concept-origin}
of the specified resource is likely to be beneficial, as it is highly
likely that the user will require resources located at that
[origin](browsers.html#concept-origin){#link-type-dns-prefetch:concept-origin-2},
and the user experience would be improved by preempting the latency
costs associated with DNS resolution.

There is no default type for resources given by the
[`dns-prefetch`{#link-type-dns-prefetch:link-type-dns-prefetch-3}](#link-type-dns-prefetch)
keyword.

The appropriate times to [fetch and
process](semantics.html#fetch-and-process-the-linked-resource){#link-type-dns-prefetch:fetch-and-process-the-linked-resource}
this type of link are:

- When the [external resource
  link](#external-resource-link){#link-type-dns-prefetch:external-resource-link-2}
  is created on a
  [`link`{#link-type-dns-prefetch:the-link-element-2}](semantics.html#the-link-element)
  element that is already [browsing-context
  connected](infrastructure.html#browsing-context-connected){#link-type-dns-prefetch:browsing-context-connected}.

- When the [external resource
  link](#external-resource-link){#link-type-dns-prefetch:external-resource-link-3}\'s
  [`link`{#link-type-dns-prefetch:the-link-element-3}](semantics.html#the-link-element)
  element [becomes browsing-context
  connected](infrastructure.html#becomes-browsing-context-connected){#link-type-dns-prefetch:becomes-browsing-context-connected}.

- When the
  [`href`{#link-type-dns-prefetch:attr-link-href}](semantics.html#attr-link-href)
  attribute of the
  [`link`{#link-type-dns-prefetch:the-link-element-4}](semantics.html#the-link-element)
  element of an [external resource
  link](#external-resource-link){#link-type-dns-prefetch:external-resource-link-4}
  that is already [browsing-context
  connected](infrastructure.html#browsing-context-connected){#link-type-dns-prefetch:browsing-context-connected-2}
  is changed.

The [fetch and process the linked
resource](semantics.html#fetch-and-process-the-linked-resource){#link-type-dns-prefetch:fetch-and-process-the-linked-resource-2}
steps for this type of linked resource, given a
[`link`{#link-type-dns-prefetch:the-link-element-5}](semantics.html#the-link-element)
element `el`{.variable}, are:

1.  Let `url`{.variable} be the result of [encoding-parsing a
    URL](urls-and-fetching.html#encoding-parsing-a-url){#link-type-dns-prefetch:encoding-parsing-a-url}
    given `el`{.variable}\'s
    [`href`{#link-type-dns-prefetch:attr-link-href-2}](semantics.html#attr-link-href)
    attribute\'s value, relative to `el`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#link-type-dns-prefetch:node-document
    x-internal="node-document"}.

2.  If `url`{.variable} is failure, then return.

3.  Let `partitionKey`{.variable} be the result of [determining the
    network partition
    key](https://fetch.spec.whatwg.org/#determine-the-network-partition-key){#link-type-dns-prefetch:determine-the-network-partition-key
    x-internal="determine-the-network-partition-key"} given
    `el`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#link-type-dns-prefetch:node-document-2
    x-internal="node-document"}\'s [relevant settings
    object](webappapis.html#relevant-settings-object){#link-type-dns-prefetch:relevant-settings-object}.

4.  The user agent should [resolve an
    origin](https://fetch.spec.whatwg.org/#resolve-an-origin){#link-type-dns-prefetch:resolve-an-origin
    x-internal="resolve-an-origin"} given `partitionKey`{.variable} and
    `url`{.variable}\'s
    [origin](https://url.spec.whatwg.org/#concept-url-origin){#link-type-dns-prefetch:concept-url-origin
    x-internal="concept-url-origin"}.

    As the results of this algorithm can be cached, future fetches could
    be faster.

##### [4.6.7.6]{.secno} Link type \"[`expect`]{.dfn dfn-for="link/rel" dfn-type="attr-value"}\"[](#link-type-expect){.self-link}

The [`expect`{#link-type-expect:link-type-expect}](#link-type-expect)
keyword may be used with
[`link`{#link-type-expect:the-link-element}](semantics.html#the-link-element)
elements. This keyword creates an [internal resource
link](#internal-resource-link){#link-type-expect:internal-resource-link}.

An [internal resource
link](#internal-resource-link){#link-type-expect:internal-resource-link-2}
created by the
[`expect`{#link-type-expect:link-type-expect-2}](#link-type-expect)
keyword can be used to [block
rendering](dom.html#render-blocked){#link-type-expect:render-blocked}
until the element that it
[indicates](browsing-the-web.html#the-indicated-part-of-the-document){#link-type-expect:the-indicated-part-of-the-document}
is connected to the document and fully parsed.

There is no default type for resources given by the
[`expect`{#link-type-expect:link-type-expect-3}](#link-type-expect)
keyword.

Whenever any of the following conditions occur for a
[`link`{#link-type-expect:the-link-element-2}](semantics.html#the-link-element)
element `el`{.variable}:

- the
  [`expect`{#link-type-expect:link-type-expect-4}](#link-type-expect)
  [internal resource
  link](#internal-resource-link){#link-type-expect:internal-resource-link-3}
  is created on `el`{.variable} that is already [browsing-context
  connected](infrastructure.html#browsing-context-connected){#link-type-expect:browsing-context-connected};

- an [`expect`{#link-type-expect:link-type-expect-5}](#link-type-expect)
  [internal resource
  link](#internal-resource-link){#link-type-expect:internal-resource-link-4}
  has been created on `el`{.variable} and `el`{.variable} becomes
  [browsing-context
  connected](infrastructure.html#browsing-context-connected){#link-type-expect:browsing-context-connected-2};

- an
  [`expect`{#link-type-expect:link-type-expect-6}](#link-type-expect)[internal
  resource
  link](#internal-resource-link){#link-type-expect:internal-resource-link-5}
  has been created on `el`{.variable}, `el`{.variable} is already
  [browsing-context
  connected](infrastructure.html#browsing-context-connected){#link-type-expect:browsing-context-connected-3},
  and `el`{.variable}\'s
  [`href`{#link-type-expect:attr-link-href}](semantics.html#attr-link-href)
  attribute is set, changed, or removed; or

- an [`expect`{#link-type-expect:link-type-expect-7}](#link-type-expect)
  [internal resource
  link](#internal-resource-link){#link-type-expect:internal-resource-link-6}
  has been created on `el`{.variable}, `el`{.variable} is already
  [browsing-context
  connected](infrastructure.html#browsing-context-connected){#link-type-expect:browsing-context-connected-4},
  and `el`{.variable}\'s
  [`media`{#link-type-expect:attr-link-media}](semantics.html#attr-link-media)
  attribute is set, changed, or removed,

then
[process](#process-internal-resource-link){#link-type-expect:process-internal-resource-link}
`el`{.variable}.

To [process internal resource link]{#process-internal-resource-link
.dfn} given a
[`link`{#link-type-expect:the-link-element-3}](semantics.html#the-link-element)
element `el`{.variable}, run these steps:

1.  Let `doc`{.variable} be `el`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#link-type-expect:node-document
    x-internal="node-document"}.

2.  Let `url`{.variable} be the result of [encoding-parsing a
    URL](urls-and-fetching.html#encoding-parsing-a-url){#link-type-expect:encoding-parsing-a-url}
    given `el`{.variable}\'s
    [`href`{#link-type-expect:attr-link-href-2}](semantics.html#attr-link-href)
    attribute\'s value, relative to `doc`{.variable}.

3.  If this fails, or if `url`{.variable} does not
    [equal](https://url.spec.whatwg.org/#concept-url-equals){#link-type-expect:concept-url-equals
    x-internal="concept-url-equals"} `doc`{.variable}\'s
    [URL](https://dom.spec.whatwg.org/#concept-document-url){#link-type-expect:the-document's-address
    x-internal="the-document's-address"} with *[exclude
    fragments](https://url.spec.whatwg.org/#url-equals-exclude-fragments){x-internal="url-equals-exclude-fragments"}*
    set to false, then [unblock
    rendering](dom.html#unblock-rendering){#link-type-expect:unblock-rendering}
    on `el`{.variable} and return.

4.  Let `indicatedElement`{.variable} be the result of [selecting the
    indicated
    part](browsing-the-web.html#select-the-indicated-part){#link-type-expect:select-the-indicated-part}
    given `doc`{.variable} and `url`{.variable}.

5.  If all of the following are true:

    - `doc`{.variable}\'s [current document
      readiness](dom.html#current-document-readiness){#link-type-expect:current-document-readiness}
      is \"`loading`\";

    - `el`{.variable} creates an [internal resource
      link](#internal-resource-link){#link-type-expect:internal-resource-link-7};

    - `el`{.variable} is [browsing-context
      connected](infrastructure.html#browsing-context-connected){#link-type-expect:browsing-context-connected-5};

    - `el`{.variable}\'s
      [`rel`{#link-type-expect:attr-link-rel}](semantics.html#attr-link-rel)
      attribute contains
      [`expect`{#link-type-expect:link-type-expect-8}](#link-type-expect);

    - `el`{.variable} is [potentially
      render-blocking](urls-and-fetching.html#potentially-render-blocking){#link-type-expect:potentially-render-blocking};

    - `el`{.variable}\'s
      [`media`{#link-type-expect:attr-link-media-2}](semantics.html#attr-link-media)
      attribute [matches the
      environment](common-microsyntaxes.html#matches-the-environment){#link-type-expect:matches-the-environment};
      and

    - `indicatedElement`{.variable} is not an element, or is on a [stack
      of open
      elements](parsing.html#stack-of-open-elements){#link-type-expect:stack-of-open-elements}
      of an [HTML
      parser](parsing.html#html-parser){#link-type-expect:html-parser}
      whose associated
      [`Document`{#link-type-expect:document}](dom.html#document) is
      `doc`{.variable},

    then [block
    rendering](dom.html#block-rendering){#link-type-expect:block-rendering}
    on `el`{.variable}.

6.  Otherwise, [unblock
    rendering](dom.html#unblock-rendering){#link-type-expect:unblock-rendering-2}
    on `el`{.variable}.

To [process internal resource links]{#process-internal-resource-links
.dfn} given a
[`Document`{#link-type-expect:document-2}](dom.html#document)
`doc`{.variable}:

1.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#link-type-expect:list-iterate
    x-internal="list-iterate"}
    [`expect`{#link-type-expect:link-type-expect-9}](#link-type-expect)
    [`link`{#link-type-expect:the-link-element-4}](semantics.html#the-link-element)
    element `link`{.variable} in `doc`{.variable}\'s [render-blocking
    element
    set](dom.html#render-blocking-element-set){#link-type-expect:render-blocking-element-set},
    [process](#process-internal-resource-link){#link-type-expect:process-internal-resource-link-2}
    `link`{.variable}.

The following [attribute change
steps](https://dom.spec.whatwg.org/#concept-element-attributes-change-ext){#link-type-expect:concept-element-attributes-change-ext
x-internal="concept-element-attributes-change-ext"}, given
`element`{.variable}, `localName`{.variable}, `value`{.variable}, and
`namespace`{.variable}, are used to ensure
[`expect`{#link-type-expect:link-type-expect-10}](#link-type-expect)
[`link`{#link-type-expect:the-link-element-5}](semantics.html#the-link-element)
elements respond to dynamic
[`id`{#link-type-expect:the-id-attribute}](dom.html#the-id-attribute)
and [`name`{#link-type-expect:attr-a-name}](obsolete.html#attr-a-name)
changes:

1.  If `namespace`{.variable} is not null, then return.

2.  If `element`{.variable} is in a [stack of open
    elements](parsing.html#stack-of-open-elements){#link-type-expect:stack-of-open-elements-2}
    of an [HTML
    parser](parsing.html#html-parser){#link-type-expect:html-parser-2},
    then return.

3.  If any of the following is true:

    - `localName`{.variable} is
      [`id`{#link-type-expect:the-id-attribute-2}](dom.html#the-id-attribute);
      or

    - `localName`{.variable} is
      [`name`{#link-type-expect:attr-a-name-2}](obsolete.html#attr-a-name)
      and `element`{.variable} is an
      [`a`{#link-type-expect:the-a-element}](text-level-semantics.html#the-a-element)
      element,

    then [process internal resource
    links](#process-internal-resource-links){#link-type-expect:process-internal-resource-links}
    given `element`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#link-type-expect:node-document-2
    x-internal="node-document"}.

##### [4.6.7.7]{.secno} Link type \"[`external`]{.dfn dfn-for="a/rel,area/rel,form/rel" dfn-type="attr-value"}\"[](#link-type-external){.self-link}

The
[`external`{#link-type-external:link-type-external}](#link-type-external)
keyword may be used with
[`a`{#link-type-external:the-a-element}](text-level-semantics.html#the-a-element),
[`area`{#link-type-external:the-area-element}](image-maps.html#the-area-element),
and
[`form`{#link-type-external:the-form-element}](forms.html#the-form-element)
elements. This keyword does not create a
[hyperlink](#hyperlink){#link-type-external:hyperlink}, but
[annotates](#hyperlink-annotation){#link-type-external:hyperlink-annotation}
any other hyperlinks created by the element (the implied hyperlink, if
no other keywords create one).

The
[`external`{#link-type-external:link-type-external-2}](#link-type-external)
keyword indicates that the link is leading to a document that is not
part of the site that the current document forms a part of.

##### [4.6.7.8]{.secno} Link type \"[`help`]{.dfn dfn-for="link/rel,a/rel,area/rel,form/rel" dfn-type="attr-value"}\"[](#link-type-help){.self-link}

The [`help`{#link-type-help:link-type-help}](#link-type-help) keyword
may be used with
[`link`{#link-type-help:the-link-element}](semantics.html#the-link-element),
[`a`{#link-type-help:the-a-element}](text-level-semantics.html#the-a-element),
[`area`{#link-type-help:the-area-element}](image-maps.html#the-area-element),
and
[`form`{#link-type-help:the-form-element}](forms.html#the-form-element)
elements. This keyword creates a
[hyperlink](#hyperlink){#link-type-help:hyperlink}.

For
[`a`{#link-type-help:the-a-element-2}](text-level-semantics.html#the-a-element),
[`area`{#link-type-help:the-area-element-2}](image-maps.html#the-area-element),
and
[`form`{#link-type-help:the-form-element-2}](forms.html#the-form-element)
elements, the
[`help`{#link-type-help:link-type-help-2}](#link-type-help) keyword
indicates that the referenced document provides further help information
for the parent of the element defining the hyperlink, and its children.

::: example
In the following example, the form control has associated
context-sensitive help. The user agent could use this information, for
example, displaying the referenced document if the user presses the
\"Help\" or \"F1\" key.

``` html
 <p><label> Topic: <input name=topic> <a href="help/topic.html" rel="help">(Help)</a></label></p>
```
:::

For
[`link`{#link-type-help:the-link-element-2}](semantics.html#the-link-element)
elements, the
[`help`{#link-type-help:link-type-help-3}](#link-type-help) keyword
indicates that the referenced document provides help for the page as a
whole.

For
[`a`{#link-type-help:the-a-element-3}](text-level-semantics.html#the-a-element)
and
[`area`{#link-type-help:the-area-element-3}](image-maps.html#the-area-element)
elements, on some browsers, the
[`help`{#link-type-help:link-type-help-4}](#link-type-help) keyword
causes the link to use a different cursor.

##### [4.6.7.9]{.secno} Link type \"[`icon`]{.dfn dfn-for="link/rel" dfn-type="attr-value"}\"[](#rel-icon){.self-link} {#rel-icon}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Link_types#icon](https://developer.mozilla.org/en-US/docs/Web/HTML/Link_types#icon "The rel attribute defines the relationship between a linked resource and the current document. Valid on <link>, <a>, <area>, and <form>, the supported values depend on the element on which the attribute is found.")

Support in all current engines.

::: support
[Firefox2+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera9+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer11]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android4+]{.firefox_android .yes}[Safari iOSNo]{.safari_ios
.no}[Chrome Android18+]{.chrome_android .yes}[WebView
Android38+]{.webview_android .yes}[Samsung
Internet4.0+]{.samsunginternet_android .yes}[Opera
AndroidNo]{.opera_android .no}

------------------------------------------------------------------------

[[caniuse.com
table](https://caniuse.com/#feat=link-icon-png "Favicons")]{.caniuse}
:::
::::
:::::

The [`icon`{#rel-icon:rel-icon}](#rel-icon) keyword may be used with
[`link`{#rel-icon:the-link-element}](semantics.html#the-link-element)
elements. This keyword creates an [external resource
link](#external-resource-link){#rel-icon:external-resource-link}.

The specified resource is an icon representing the page or site, and
should be used by the user agent when representing the page in the user
interface.

Icons could be auditory icons, visual icons, or other kinds of icons. If
multiple icons are provided, the user agent must select the most
appropriate icon according to the
[`type`{#rel-icon:attr-link-type}](semantics.html#attr-link-type),
[`media`{#rel-icon:attr-link-media}](semantics.html#attr-link-media),
and [`sizes`{#rel-icon:attr-link-sizes}](semantics.html#attr-link-sizes)
attributes. If there are multiple equally appropriate icons, user agents
must use the last one declared in [tree
order](https://dom.spec.whatwg.org/#concept-tree-order){#rel-icon:tree-order
x-internal="tree-order"} at the time that the user agent collected the
list of icons. If the user agent tries to use an icon but that icon is
determined, upon closer examination, to in fact be inappropriate (e.g.
because it uses an unsupported format), then the user agent must try the
next-most-appropriate icon as determined by the attributes.

User agents are not required to update icons when the list of icons
changes, but are encouraged to do so.

There is no default type for resources given by the
[`icon`{#rel-icon:rel-icon-2}](#rel-icon) keyword. However, for the
purposes of [determining the type of the
resource](semantics.html#concept-link-type-sniffing), user agents must
expect the resource to be an image.

The
[`sizes`{#rel-icon:attr-link-sizes-2}](semantics.html#attr-link-sizes)
keywords represent icon sizes in raw pixels (as opposed to [CSS
pixels](https://drafts.csswg.org/css-values/#px){#rel-icon:'px'
x-internal="'px'"}).

An icon that is 50 [CSS
pixels](https://drafts.csswg.org/css-values/#px){#rel-icon:'px'-2
x-internal="'px'"} wide intended for displays with a device pixel
density of two device pixels per [CSS
pixel](https://drafts.csswg.org/css-values/#px){#rel-icon:'px'-3
x-internal="'px'"} (2x, 192dpi) would have a width of 100 raw pixels.
This feature does not support indicating that a different resource is to
be used for small high-resolution icons vs large low-resolution icons
(e.g. 50×50 2x vs 100×100 1x).

To parse and process the attribute\'s value, the user agent must first
[split the attribute\'s value on ASCII
whitespace](https://infra.spec.whatwg.org/#split-on-ascii-whitespace){#rel-icon:split-a-string-on-spaces
x-internal="split-a-string-on-spaces"}, and must then parse each
resulting keyword to determine what it represents.

The [`any`]{#attr-link-sizes-any .dfn} keyword represents that the
resource contains a scalable icon, e.g. as provided by an SVG image.

Other keywords must be further parsed as follows to determine what they
represent:

1.  If the keyword doesn\'t contain exactly one U+0078 LATIN SMALL
    LETTER X or U+0058 LATIN CAPITAL LETTER X character, then this
    keyword doesn\'t represent anything. Return for that keyword.

2.  Let `width string`{.variable} be the string before the \"`x`\" or
    \"`X`\".

3.  Let `height string`{.variable} be the string after the \"`x`\" or
    \"`X`\".

4.  If either `width string`{.variable} or `height string`{.variable}
    start with a U+0030 DIGIT ZERO (0) character or contain any
    characters other than [ASCII
    digits](https://infra.spec.whatwg.org/#ascii-digit){#rel-icon:ascii-digits
    x-internal="ascii-digits"}, then this keyword doesn\'t represent
    anything. Return for that keyword.

5.  Apply the [rules for parsing non-negative
    integers](common-microsyntaxes.html#rules-for-parsing-non-negative-integers){#rel-icon:rules-for-parsing-non-negative-integers}
    to `width string`{.variable} to obtain `width`{.variable}.

6.  Apply the [rules for parsing non-negative
    integers](common-microsyntaxes.html#rules-for-parsing-non-negative-integers){#rel-icon:rules-for-parsing-non-negative-integers-2}
    to `height string`{.variable} to obtain `height`{.variable}.

7.  The keyword represents that the resource contains a bitmap icon with
    a width of `width`{.variable} device pixels and a height of
    `height`{.variable} device pixels.

The keywords specified on the
[`sizes`{#rel-icon:attr-link-sizes-3}](semantics.html#attr-link-sizes)
attribute must not represent icon sizes that are not actually available
in the linked resource.

The [linked resource fetch setup
steps](semantics.html#linked-resource-fetch-setup-steps){#rel-icon:linked-resource-fetch-setup-steps}
for this type of linked resource, given a
[`link`{#rel-icon:the-link-element-2}](semantics.html#the-link-element)
element `el`{.variable} and
[request](https://fetch.spec.whatwg.org/#concept-request){#rel-icon:concept-request
x-internal="concept-request"} `request`{.variable}, are:

1.  Set `request`{.variable}\'s
    [destination](https://fetch.spec.whatwg.org/#concept-request-destination){#rel-icon:concept-request-destination
    x-internal="concept-request-destination"} to \"`image`\".

2.  Return true.

The [process a link
header](semantics.html#process-a-link-header){#rel-icon:process-a-link-header}
steps for this type of linked resource are to do nothing.

In the absence of a
[`link`{#rel-icon:the-link-element-3}](semantics.html#the-link-element)
with the [`icon`{#rel-icon:rel-icon-3}](#rel-icon) keyword, for
[`Document`{#rel-icon:document}](dom.html#document) objects whose
[URL](https://dom.spec.whatwg.org/#concept-document-url){#rel-icon:the-document's-address
x-internal="the-document's-address"}\'s
[scheme](https://url.spec.whatwg.org/#concept-url-scheme){#rel-icon:concept-url-scheme
x-internal="concept-url-scheme"} is an [HTTP(S)
scheme](https://fetch.spec.whatwg.org/#http-scheme){#rel-icon:http(s)-scheme
x-internal="http(s)-scheme"}, user agents may instead run these steps
[in parallel](infrastructure.html#in-parallel){#rel-icon:in-parallel}:

1.  Let `request`{.variable} be a new
    [request](https://fetch.spec.whatwg.org/#concept-request){#rel-icon:concept-request-2
    x-internal="concept-request"} whose
    [URL](https://fetch.spec.whatwg.org/#concept-request-url){#rel-icon:concept-request-url
    x-internal="concept-request-url"} is the [URL
    record](https://url.spec.whatwg.org/#concept-url){#rel-icon:url-record
    x-internal="url-record"} obtained by resolving the
    [URL](https://url.spec.whatwg.org/#concept-url){#rel-icon:url
    x-internal="url"} \"`/favicon.ico`\" against the
    [`Document`{#rel-icon:document-2}](dom.html#document) object\'s
    [URL](https://dom.spec.whatwg.org/#concept-document-url){#rel-icon:the-document's-address-2
    x-internal="the-document's-address"},
    [client](https://fetch.spec.whatwg.org/#concept-request-client){#rel-icon:concept-request-client
    x-internal="concept-request-client"} is the
    [`Document`{#rel-icon:document-3}](dom.html#document) object\'s
    [relevant settings
    object](webappapis.html#relevant-settings-object){#rel-icon:relevant-settings-object},
    [destination](https://fetch.spec.whatwg.org/#concept-request-destination){#rel-icon:concept-request-destination-2
    x-internal="concept-request-destination"} is \"`image`\",
    [synchronous
    flag](https://fetch.spec.whatwg.org/#synchronous-flag){#rel-icon:synchronous-flag
    x-internal="synchronous-flag"} is set, [credentials
    mode](https://fetch.spec.whatwg.org/#concept-request-credentials-mode){#rel-icon:concept-request-credentials-mode
    x-internal="concept-request-credentials-mode"} is \"`include`\", and
    whose [use-URL-credentials
    flag](https://fetch.spec.whatwg.org/#concept-request-use-url-credentials-flag){#rel-icon:use-url-credentials-flag
    x-internal="use-url-credentials-flag"} is set.

2.  Let `response`{.variable} be the result of
    [fetching](https://fetch.spec.whatwg.org/#concept-fetch){#rel-icon:concept-fetch
    x-internal="concept-fetch"} `request`{.variable}.

3.  Use `response`{.variable}\'s [unsafe
    response](urls-and-fetching.html#unsafe-response){#rel-icon:unsafe-response}
    as an icon as if it had been declared using the
    [`icon`{#rel-icon:rel-icon-4}](#rel-icon) keyword.

::: example
The following snippet shows the top part of an application with several
icons.

``` html
<!DOCTYPE HTML>
<html lang="en">
 <head>
  <title>lsForums — Inbox</title>
  <link rel=icon href=favicon.png sizes="16x16" type="image/png">
  <link rel=icon href=windows.ico sizes="32x32 48x48" type="image/vnd.microsoft.icon">
  <link rel=icon href=mac.icns sizes="128x128 512x512 8192x8192 32768x32768">
  <link rel=icon href=iphone.png sizes="57x57" type="image/png">
  <link rel=icon href=gnome.svg sizes="any" type="image/svg+xml">
  <link rel=stylesheet href=lsforums.css>
  <script src=lsforums.js></script>
  <meta name=application-name content="lsForums">
 </head>
 <body>
  ...
```
:::

For historical reasons, the [`icon`{#rel-icon:rel-icon-5}](#rel-icon)
keyword may be preceded by the keyword \"`shortcut`\". If the
\"`shortcut`\" keyword is present, the
[`rel`{#rel-icon:attr-hyperlink-rel}](#attr-hyperlink-rel) attribute\'s
entire value must be an [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#rel-icon:ascii-case-insensitive
x-internal="ascii-case-insensitive"} match for the string
\"`shortcut icon`\" (with a single U+0020 SPACE character between the
tokens and no other [ASCII
whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#rel-icon:space-characters
x-internal="space-characters"}).

##### [4.6.7.10]{.secno} Link type \"[`license`]{.dfn dfn-for="link/rel,a/rel,area/rel,form/rel" dfn-type="attr-value"}\"[](#link-type-license){.self-link}

The
[`license`{#link-type-license:link-type-license}](#link-type-license)
keyword may be used with
[`link`{#link-type-license:the-link-element}](semantics.html#the-link-element),
[`a`{#link-type-license:the-a-element}](text-level-semantics.html#the-a-element),
[`area`{#link-type-license:the-area-element}](image-maps.html#the-area-element),
and
[`form`{#link-type-license:the-form-element}](forms.html#the-form-element)
elements. This keyword creates a
[hyperlink](#hyperlink){#link-type-license:hyperlink}.

The
[`license`{#link-type-license:link-type-license-2}](#link-type-license)
keyword indicates that the referenced document provides the copyright
license terms under which the main content of the current document is
provided.

This specification does not specify how to distinguish between the main
content of a document and content that is not deemed to be part of that
main content. The distinction should be made clear to the user.

::: example
Consider a photo sharing site. A page on that site might describe and
show a photograph, and the page might be marked up as follows:

``` html
<!DOCTYPE HTML>
<html lang="en">
 <head>
  <title>Exampl Pictures: Kissat</title>
  <link rel="stylesheet" href="/style/default">
 </head>
 <body>
  <h1>Kissat</h1>
  <nav>
   <a href="../">Return to photo index</a>
  </nav>
  <figure>
   <img src="/pix/39627052_fd8dcd98b5.jpg">
   <figcaption>Kissat</figcaption>
  </figure>
  <p>One of them has six toes!</p>
  <p><small><a rel="license" href="http://www.opensource.org/licenses/mit-license.php">MIT Licensed</a></small></p>
  <footer>
   <a href="/">Home</a> | <a href="../">Photo index</a>
   <p><small>© copyright 2009 Exampl Pictures. All Rights Reserved.</small></p>
  </footer>
 </body>
</html>
```

In this case the
[`license`{#link-type-license:link-type-license-3}](#link-type-license)
applies to just the photo (the main content of the document), not the
whole document. In particular not the design of the page itself, which
is covered by the copyright given at the bottom of the document. This
could be made clearer in the styling (e.g. making the license link
prominently positioned near the photograph, while having the page
copyright in light small text at the foot of the page).
:::

**Synonyms**: For historical reasons, user agents must also treat the
keyword \"`copyright`\" like the
[`license`{#link-type-license:link-type-license-4}](#link-type-license)
keyword.

##### [4.6.7.11]{.secno} Link type \"[`manifest`]{.dfn dfn-for="link/rel" dfn-type="attr-value"}\"[](#link-type-manifest){.self-link}

::::: {.mdn-anno .wrapped}
**⚠**MDN

:::: feature
[Link_types/manifest](https://developer.mozilla.org/en-US/docs/Web/HTML/Link_types/manifest "The manifest keyword for the rel attribute of the <link> element indicates that the target resource is a Web app manifest.")

Support in one engine only.

::: support
[Firefox?]{.firefox .unknown}[Safari?]{.safari
.unknown}[ChromeNo]{.chrome .no}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[EdgeNo]{.edge_blink .no}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet Explorer?]{.ie .unknown}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android39+]{.chrome_android .yes}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

The
[`manifest`{#link-type-manifest:link-type-manifest}](#link-type-manifest)
keyword may be used with
[`link`{#link-type-manifest:the-link-element}](semantics.html#the-link-element)
elements. This keyword creates an [external resource
link](#external-resource-link){#link-type-manifest:external-resource-link}.

The
[`manifest`{#link-type-manifest:link-type-manifest-2}](#link-type-manifest)
keyword indicates the manifest file that provides metadata associated
with the current document.

There is no default type for resources given by the
[`manifest`{#link-type-manifest:link-type-manifest-3}](#link-type-manifest)
keyword.

When a web application is not
[installed](https://w3c.github.io/manifest/#dfn-installed-web-application){#link-type-manifest:installed-web-application
x-internal="installed-web-application"}, the appropriate time to [fetch
and process the linked
resource](semantics.html#fetch-and-process-the-linked-resource){#link-type-manifest:fetch-and-process-the-linked-resource}
for this link type is when the user agent deems it necessary. For
example, when the user chooses to [install the web
application](https://w3c.github.io/manifest/#dfn-installed-web-application){#link-type-manifest:installed-web-application-2
x-internal="installed-web-application"}.

For an [installed web
application](https://w3c.github.io/manifest/#dfn-installed-web-application){#link-type-manifest:installed-web-application-3
x-internal="installed-web-application"}, the appropriate times to [fetch
and process the linked
resource](semantics.html#fetch-and-process-the-linked-resource){#link-type-manifest:fetch-and-process-the-linked-resource-2}
for this link type are:

- When the [external resource
  link](#external-resource-link){#link-type-manifest:external-resource-link-2}
  is created on a
  [`link`{#link-type-manifest:the-link-element-2}](semantics.html#the-link-element)
  element that is already [browsing-context
  connected](infrastructure.html#browsing-context-connected){#link-type-manifest:browsing-context-connected}.

- When the [external resource
  link](#external-resource-link){#link-type-manifest:external-resource-link-3}\'s
  [`link`{#link-type-manifest:the-link-element-3}](semantics.html#the-link-element)
  element [becomes browsing-context
  connected](infrastructure.html#becomes-browsing-context-connected){#link-type-manifest:becomes-browsing-context-connected}.

- When the
  [`href`{#link-type-manifest:attr-link-href}](semantics.html#attr-link-href)
  attribute of the
  [`link`{#link-type-manifest:the-link-element-4}](semantics.html#the-link-element)
  element of an [external resource
  link](#external-resource-link){#link-type-manifest:external-resource-link-4}
  that is already [browsing-context
  connected](infrastructure.html#browsing-context-connected){#link-type-manifest:browsing-context-connected-2}
  is changed.

In any case, only the first
[`link`{#link-type-manifest:the-link-element-5}](semantics.html#the-link-element)
element in [tree
order](https://dom.spec.whatwg.org/#concept-tree-order){#link-type-manifest:tree-order
x-internal="tree-order"} whose
[`rel`{#link-type-manifest:attr-link-rel}](semantics.html#attr-link-rel)
attribute contains the token
[`manifest`{#link-type-manifest:link-type-manifest-4}](#link-type-manifest)
may be used.

A user agent must not [delay the load
event](parsing.html#delay-the-load-event){#link-type-manifest:delay-the-load-event}
for this link type.

The [linked resource fetch setup
steps](semantics.html#linked-resource-fetch-setup-steps){#link-type-manifest:linked-resource-fetch-setup-steps}
for this type of linked resource, given a
[`link`{#link-type-manifest:the-link-element-6}](semantics.html#the-link-element)
element `el`{.variable} and
[request](https://fetch.spec.whatwg.org/#concept-request){#link-type-manifest:concept-request
x-internal="concept-request"} `request`{.variable}, are:

1.  Let `navigable`{.variable} be `el`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#link-type-manifest:node-document
    x-internal="node-document"}\'s [node
    navigable](document-sequences.html#node-navigable){#link-type-manifest:node-navigable}.

2.  If `navigable`{.variable} is null, then return false.

3.  If `navigable`{.variable} is not a [top-level
    traversable](document-sequences.html#top-level-traversable){#link-type-manifest:top-level-traversable},
    then return false.

4.  Set `request`{.variable}\'s
    [initiator](https://fetch.spec.whatwg.org/#concept-request-initiator){#link-type-manifest:concept-request-initiator
    x-internal="concept-request-initiator"} to \"`manifest`\".

5.  Set `request`{.variable}\'s
    [destination](https://fetch.spec.whatwg.org/#concept-request-destination){#link-type-manifest:concept-request-destination
    x-internal="concept-request-destination"} to \"`manifest`\".

6.  Set `request`{.variable}\'s
    [mode](https://fetch.spec.whatwg.org/#concept-request-mode){#link-type-manifest:concept-request-mode
    x-internal="concept-request-mode"} to \"`cors`\".

7.  Set `request`{.variable}\'s [credentials
    mode](https://fetch.spec.whatwg.org/#concept-request-credentials-mode){#link-type-manifest:concept-request-credentials-mode
    x-internal="concept-request-credentials-mode"} to the [CORS settings
    attribute credentials
    mode](urls-and-fetching.html#cors-settings-attribute-credentials-mode){#link-type-manifest:cors-settings-attribute-credentials-mode}
    for `el`{.variable}\'s
    [`crossorigin`{#link-type-manifest:attr-link-crossorigin}](semantics.html#attr-link-crossorigin)
    content attribute.

8.  Return true.

To [process this type of linked
resource](semantics.html#process-the-linked-resource){#link-type-manifest:process-the-linked-resource}
given a
[`link`{#link-type-manifest:the-link-element-7}](semantics.html#the-link-element)
element `el`{.variable}, boolean `success`{.variable},
[response](https://fetch.spec.whatwg.org/#concept-response){#link-type-manifest:concept-response
x-internal="concept-response"} `response`{.variable}, and [byte
sequence](https://infra.spec.whatwg.org/#byte-sequence){#link-type-manifest:byte-sequence
x-internal="byte-sequence"} `bodyBytes`{.variable}:

1.  If `response`{.variable}\'s [Content-Type
    metadata](urls-and-fetching.html#content-type){#link-type-manifest:content-type}
    is not a [JSON MIME
    type](https://mimesniff.spec.whatwg.org/#json-mime-type){#link-type-manifest:json-mime-type
    x-internal="json-mime-type"}, then set `success`{.variable} to
    false.

2.  If `success`{.variable} is true:

    1.  Let `document URL`{.variable} be `el`{.variable}\'s [node
        document](https://dom.spec.whatwg.org/#concept-node-document){#link-type-manifest:node-document-2
        x-internal="node-document"}\'s
        [URL](https://dom.spec.whatwg.org/#concept-document-url){#link-type-manifest:the-document's-address
        x-internal="the-document's-address"}.

    2.  Let `manifest URL`{.variable} be `response`{.variable}\'s
        [URL](https://fetch.spec.whatwg.org/#concept-response-url){#link-type-manifest:concept-response-url
        x-internal="concept-response-url"}.

    3.  [Process the
        manifest](https://w3c.github.io/manifest/#dfn-processing-a-manifest){#link-type-manifest:process-the-manifest
        x-internal="process-the-manifest"} given
        `document URL`{.variable}, `manifest URL`{.variable}, and
        `bodyBytes`{.variable}.
        [\[MANIFEST\]](references.html#refsMANIFEST)

The [process a link
header](semantics.html#process-a-link-header){#link-type-manifest:process-a-link-header}
steps for this type of linked resource are to do nothing.

##### [4.6.7.12]{.secno} Link type \"[`modulepreload`]{.dfn dfn-for="link/rel" dfn-type="attr-value"}\"[](#link-type-modulepreload){.self-link}

::::: {.mdn-anno .wrapped}
MDN

:::: feature
[Link_types/modulepreload](https://developer.mozilla.org/en-US/docs/Web/HTML/Link_types/modulepreload "The modulepreload keyword, for the rel attribute of the <link> element, provides a declarative way to preemptively fetch a module script, parse and compile it, and store it in the document's module map for later execution.")

::: support
[Firefox115+]{.firefox .yes}[Safari?]{.safari
.unknown}[Chrome66+]{.chrome .yes}

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

The
[`modulepreload`{#link-type-modulepreload:link-type-modulepreload}](#link-type-modulepreload)
keyword may be used with
[`link`{#link-type-modulepreload:the-link-element}](semantics.html#the-link-element)
elements. This keyword creates an [external resource
link](#external-resource-link){#link-type-modulepreload:external-resource-link}.
This keyword is [body-ok](#body-ok){#link-type-modulepreload:body-ok}.

The
[`modulepreload`{#link-type-modulepreload:link-type-modulepreload-2}](#link-type-modulepreload)
keyword is a specialized alternative to the
[`preload`{#link-type-modulepreload:link-type-preload}](#link-type-preload)
keyword, with a processing model geared toward preloading [module
scripts](webappapis.html#module-script){#link-type-modulepreload:module-script}.
In particular, it uses the specific fetch behavior for module scripts
(including, e.g., a different interpretation of the
[`crossorigin`{#link-type-modulepreload:attr-link-crossorigin}](semantics.html#attr-link-crossorigin)
attribute), and places the result into the appropriate [module
map](dom.html#concept-document-module-map){#link-type-modulepreload:concept-document-module-map}
for later evaluation. In contrast, a similar [external resource
link](#external-resource-link){#link-type-modulepreload:external-resource-link-2}
using the
[`preload`{#link-type-modulepreload:link-type-preload-2}](#link-type-preload)
keyword would place the result in the preload cache, without affecting
the document\'s [module
map](dom.html#concept-document-module-map){#link-type-modulepreload:concept-document-module-map-2}.

Additionally, implementations can take advantage of the fact that
[module
scripts](webappapis.html#module-script){#link-type-modulepreload:module-script-2}
declare their dependencies in order to fetch the specified module\'s
dependency as well. This is intended as an optimization opportunity,
since the user agent knows that, in all likelihood, those dependencies
will also be needed later. It will not generally be observable without
using technology such as service workers, or monitoring on the server
side. Notably, the appropriate
[`load`{#link-type-modulepreload:event-load}](indices.html#event-load)
or
[`error`{#link-type-modulepreload:event-error}](indices.html#event-error)
events will occur after the specified module is fetched, and will not
wait for any dependencies.

A user agent must not [delay the load
event](parsing.html#delay-the-load-event){#link-type-modulepreload:delay-the-load-event}
for this link type.

The appropriate times to [fetch and process the linked
resource](semantics.html#fetch-and-process-the-linked-resource){#link-type-modulepreload:fetch-and-process-the-linked-resource}
for such a link are:

- When the [external resource
  link](#external-resource-link){#link-type-modulepreload:external-resource-link-3}
  is created on a
  [`link`{#link-type-modulepreload:the-link-element-2}](semantics.html#the-link-element)
  element that is already [browsing-context
  connected](infrastructure.html#browsing-context-connected){#link-type-modulepreload:browsing-context-connected}.

- When the [external resource
  link](#external-resource-link){#link-type-modulepreload:external-resource-link-4}\'s
  [`link`{#link-type-modulepreload:the-link-element-3}](semantics.html#the-link-element)
  element [becomes browsing-context
  connected](infrastructure.html#becomes-browsing-context-connected){#link-type-modulepreload:becomes-browsing-context-connected}.

- When the
  [`href`{#link-type-modulepreload:attr-link-href}](semantics.html#attr-link-href)
  attribute of the
  [`link`{#link-type-modulepreload:the-link-element-4}](semantics.html#the-link-element)
  element of an [external resource
  link](#external-resource-link){#link-type-modulepreload:external-resource-link-5}
  that is already [browsing-context
  connected](infrastructure.html#browsing-context-connected){#link-type-modulepreload:browsing-context-connected-2}
  is changed.

Unlike some other link relations, changing the relevant attributes (such
as
[`as`{#link-type-modulepreload:attr-link-as}](semantics.html#attr-link-as),
[`crossorigin`{#link-type-modulepreload:attr-link-crossorigin-2}](semantics.html#attr-link-crossorigin),
and
[`referrerpolicy`{#link-type-modulepreload:attr-link-referrerpolicy}](semantics.html#attr-link-referrerpolicy))
of such a
[`link`{#link-type-modulepreload:the-link-element-5}](semantics.html#the-link-element)
does not trigger a new fetch. This is because the document\'s [module
map](dom.html#concept-document-module-map){#link-type-modulepreload:concept-document-module-map-3}
has already been populated by a previous fetch, and so re-fetching would
be pointless.

The [fetch and process the linked
resource](semantics.html#fetch-and-process-the-linked-resource){#link-type-modulepreload:fetch-and-process-the-linked-resource-2}
algorithm for
[`modulepreload`{#link-type-modulepreload:link-type-modulepreload-3}](#link-type-modulepreload)
links, given a
[`link`{#link-type-modulepreload:the-link-element-6}](semantics.html#the-link-element)
element `el`{.variable}, is as follows:

1.  If `el`{.variable}\'s
    [`href`{#link-type-modulepreload:attr-link-href-2}](semantics.html#attr-link-href)
    attribute\'s value is the empty string, then return.

2.  Let `destination`{.variable} be the current state of
    `el`{.variable}\'s
    [`as`{#link-type-modulepreload:attr-link-as-2}](semantics.html#attr-link-as)
    attribute (a
    [destination](https://fetch.spec.whatwg.org/#concept-request-destination){#link-type-modulepreload:concept-request-destination
    x-internal="concept-request-destination"}), or \"`script`\" if it is
    in no state.

3.  If `destination`{.variable} is not
    [script-like](https://fetch.spec.whatwg.org/#request-destination-script-like){#link-type-modulepreload:concept-script-like-destination
    x-internal="concept-script-like-destination"}, then [queue an
    element
    task](webappapis.html#queue-an-element-task){#link-type-modulepreload:queue-an-element-task}
    on the [networking task
    source](webappapis.html#networking-task-source){#link-type-modulepreload:networking-task-source}
    given `el`{.variable} to [fire an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#link-type-modulepreload:concept-event-fire
    x-internal="concept-event-fire"} named
    [`error`{#link-type-modulepreload:event-error-2}](indices.html#event-error)
    at `el`{.variable}, and return.

4.  Let `url`{.variable} be the result of [encoding-parsing a
    URL](urls-and-fetching.html#encoding-parsing-a-url){#link-type-modulepreload:encoding-parsing-a-url}
    given `el`{.variable}\'s
    [`href`{#link-type-modulepreload:attr-link-href-3}](semantics.html#attr-link-href)
    attribute\'s value, relative to `el`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#link-type-modulepreload:node-document
    x-internal="node-document"}.

5.  If `url`{.variable} is failure, then return.

6.  Let `settings object`{.variable} be `el`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#link-type-modulepreload:node-document-2
    x-internal="node-document"}\'s [relevant settings
    object](webappapis.html#relevant-settings-object){#link-type-modulepreload:relevant-settings-object}.

7.  Let `credentials mode`{.variable} be the [CORS settings attribute
    credentials
    mode](urls-and-fetching.html#cors-settings-attribute-credentials-mode){#link-type-modulepreload:cors-settings-attribute-credentials-mode}
    for `el`{.variable}\'s
    [`crossorigin`{#link-type-modulepreload:attr-link-crossorigin-3}](semantics.html#attr-link-crossorigin)
    attribute.

8.  Let `cryptographic nonce`{.variable} be
    `el`{.variable}.[\[\[CryptographicNonce\]\]](urls-and-fetching.html#cryptographicnonce){#link-type-modulepreload:cryptographicnonce}.

9.  Let `integrity metadata`{.variable} be the value of
    `el`{.variable}\'s
    [`integrity`{#link-type-modulepreload:attr-link-integrity}](semantics.html#attr-link-integrity)
    attribute, if it is specified, or the empty string otherwise.

10. If `el`{.variable} does not have an
    [`integrity`{#link-type-modulepreload:attr-link-integrity-2}](semantics.html#attr-link-integrity)
    attribute, then set `integrity metadata`{.variable} to the result of
    [resolving a module integrity
    metadata](webappapis.html#resolving-a-module-integrity-metadata){#link-type-modulepreload:resolving-a-module-integrity-metadata}
    with `url`{.variable} and `settings object`{.variable}.

11. Let `referrer policy`{.variable} be the current state of
    `el`{.variable}\'s
    [`referrerpolicy`{#link-type-modulepreload:attr-link-referrerpolicy-2}](semantics.html#attr-link-referrerpolicy)
    attribute.

12. Let `fetch priority`{.variable} be the current state of
    `el`{.variable}\'s
    [`fetchpriority`{#link-type-modulepreload:attr-link-fetchpriority}](semantics.html#attr-link-fetchpriority)
    attribute.

13. Let `options`{.variable} be a [script fetch
    options](webappapis.html#script-fetch-options){#link-type-modulepreload:script-fetch-options}
    whose [cryptographic
    nonce](webappapis.html#concept-script-fetch-options-nonce){#link-type-modulepreload:concept-script-fetch-options-nonce}
    is `cryptographic nonce`{.variable}, [integrity
    metadata](webappapis.html#concept-script-fetch-options-integrity){#link-type-modulepreload:concept-script-fetch-options-integrity}
    is `integrity metadata`{.variable}, [parser
    metadata](webappapis.html#concept-script-fetch-options-parser){#link-type-modulepreload:concept-script-fetch-options-parser}
    is \"`not-parser-inserted`\", [credentials
    mode](webappapis.html#concept-script-fetch-options-credentials){#link-type-modulepreload:concept-script-fetch-options-credentials}
    is `credentials mode`{.variable}, [referrer
    policy](webappapis.html#concept-script-fetch-options-referrer-policy){#link-type-modulepreload:concept-script-fetch-options-referrer-policy}
    is `referrer policy`{.variable}, and [fetch
    priority](webappapis.html#concept-script-fetch-options-fetch-priority){#link-type-modulepreload:concept-script-fetch-options-fetch-priority}
    is `fetch priority`{.variable}.

14. [Fetch a modulepreload module script
    graph](webappapis.html#fetch-a-modulepreload-module-script-graph){#link-type-modulepreload:fetch-a-modulepreload-module-script-graph}
    given `url`{.variable}, `destination`{.variable},
    `settings object`{.variable}, `options`{.variable}, and with the
    following steps given `result`{.variable}:

    1.  If `result`{.variable} is null, then [fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#link-type-modulepreload:concept-event-fire-2
        x-internal="concept-event-fire"} named
        [`error`{#link-type-modulepreload:event-error-3}](indices.html#event-error)
        at `el`{.variable}, and return.

    2.  [Fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#link-type-modulepreload:concept-event-fire-3
        x-internal="concept-event-fire"} named
        [`load`{#link-type-modulepreload:event-load-2}](indices.html#event-load)
        at `el`{.variable}.

The [process a link
header](semantics.html#process-a-link-header){#link-type-modulepreload:process-a-link-header}
steps for this type of linked resource are to do nothing.

::: {#example-modulepreload-manifest .example}
The following snippet shows the top part of an application with several
modules preloaded:

``` html
<!DOCTYPE html>
<html lang="en">
<title>IRCFog</title>

<link rel="modulepreload" href="app.mjs">
<link rel="modulepreload" href="helpers.mjs">
<link rel="modulepreload" href="irc.mjs">
<link rel="modulepreload" href="fog-machine.mjs">

<script type="module" src="app.mjs">
...
```

Assume that the module graph for the application is as follows:

![The module graph is rooted at app.mjs, which depends on irc.mjs and
fog-machine.mjs. In turn, irc.mjs depends on
helpers.mjs.](/images/ircfog-modules.svg){.darkmode-aware width="301"
height="151"}

Here we see the application developer has used
[`modulepreload`{#link-type-modulepreload:link-type-modulepreload-4}](#link-type-modulepreload)
to declare all of the modules in their module graph, ensuring that the
user agent initiates fetches for them all. Without such preloading, the
user agent might need to go through multiple network roundtrips before
discovering `helpers.mjs`, if technologies such as HTTP/2 Server Push
are not in play. In this way,
[`modulepreload`{#link-type-modulepreload:link-type-modulepreload-5}](#link-type-modulepreload)
[`link`{#link-type-modulepreload:the-link-element-7}](semantics.html#the-link-element)
elements can be used as a sort of \"manifest\" of the application\'s
modules.
:::

::: {#example-modulepreload-dynamic-import .example}
The following code shows how
[`modulepreload`{#link-type-modulepreload:link-type-modulepreload-6}](#link-type-modulepreload)
links can be used in conjunction with
[`import()`{#link-type-modulepreload:import()}](https://tc39.es/ecma262/#sec-import-calls){x-internal="import()"}
to ensure network fetching is done ahead of time, so that when
[`import()`{#link-type-modulepreload:import()-2}](https://tc39.es/ecma262/#sec-import-calls){x-internal="import()"}
is called, the module is already ready (but not evaluated) in the
[module
map](webappapis.html#module-map){#link-type-modulepreload:module-map}:

``` html
<link rel="modulepreload" href="awesome-viewer.mjs">

<button onclick="import('./awesome-viewer.mjs').then(m => m.view())">
  View awesome thing
</button>
```
:::

##### [4.6.7.13]{.secno} Link type \"[`nofollow`]{.dfn dfn-for="a/rel,area/rel,form/rel" dfn-type="attr-value"}\"[](#link-type-nofollow){.self-link}

The
[`nofollow`{#link-type-nofollow:link-type-nofollow}](#link-type-nofollow)
keyword may be used with
[`a`{#link-type-nofollow:the-a-element}](text-level-semantics.html#the-a-element),
[`area`{#link-type-nofollow:the-area-element}](image-maps.html#the-area-element),
and
[`form`{#link-type-nofollow:the-form-element}](forms.html#the-form-element)
elements. This keyword does not create a
[hyperlink](#hyperlink){#link-type-nofollow:hyperlink}, but
[annotates](#hyperlink-annotation){#link-type-nofollow:hyperlink-annotation}
any other hyperlinks created by the element (the implied hyperlink, if
no other keywords create one).

The
[`nofollow`{#link-type-nofollow:link-type-nofollow-2}](#link-type-nofollow)
keyword indicates that the link is not endorsed by the original author
or publisher of the page, or that the link to the referenced document
was included primarily because of a commercial relationship between
people affiliated with the two pages.

##### [4.6.7.14]{.secno} Link type \"[`noopener`]{.dfn dfn-for="a/rel,area/rel,form/rel" dfn-type="attr-value"}\"[](#link-type-noopener){.self-link}

::::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Link_types/noopener](https://developer.mozilla.org/en-US/docs/Web/HTML/Link_types/noopener "The noopener keyword for the rel attribute of the <a>, <area>, and <form> elements instructs the browser to navigate to the target resource without granting the new browsing context access to the document that opened it — by not setting the Window.opener property on the opened window (it returns null).")

Support in all current engines.

::: support
[Firefox52+]{.firefox .yes}[Safari10.1+]{.safari
.yes}[Chrome49+]{.chrome .yes}

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
[Link_types/noopener](https://developer.mozilla.org/en-US/docs/Web/HTML/Link_types/noopener "The noopener keyword for the rel attribute of the <a>, <area>, and <form> elements instructs the browser to navigate to the target resource without granting the new browsing context access to the document that opened it — by not setting the Window.opener property on the opened window (it returns null).")

Support in all current engines.

::: support
[Firefox52+]{.firefox .yes}[Safari10.1+]{.safari
.yes}[Chrome49+]{.chrome .yes}

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

The
[`noopener`{#link-type-noopener:link-type-noopener}](#link-type-noopener)
keyword may be used with
[`a`{#link-type-noopener:the-a-element}](text-level-semantics.html#the-a-element),
[`area`{#link-type-noopener:the-area-element}](image-maps.html#the-area-element),
and
[`form`{#link-type-noopener:the-form-element}](forms.html#the-form-element)
elements. This keyword does not create a
[hyperlink](#hyperlink){#link-type-noopener:hyperlink}, but
[annotates](#hyperlink-annotation){#link-type-noopener:hyperlink-annotation}
any other hyperlinks created by the element (the implied hyperlink, if
no other keywords create one).

The keyword indicates that any newly created [top-level
traversable](document-sequences.html#top-level-traversable){#link-type-noopener:top-level-traversable}
which results from following the
[hyperlink](#hyperlink){#link-type-noopener:hyperlink-2} will not
contain an [auxiliary browsing
context](document-sequences.html#auxiliary-browsing-context){#link-type-noopener:auxiliary-browsing-context}.
E.g., the resulting
[`Window`{#link-type-noopener:window}](nav-history-apis.html#window)\'s
[`opener`{#link-type-noopener:dom-opener}](nav-history-apis.html#dom-opener)
getter will return null.

See also the [processing model](document-sequences.html#noopener).

::: example
This typically creates a [top-level
traversable](document-sequences.html#top-level-traversable){#link-type-noopener:top-level-traversable-2}
with an [auxiliary browsing
context](document-sequences.html#auxiliary-browsing-context){#link-type-noopener:auxiliary-browsing-context-2}
(assuming there is no existing
[navigable](document-sequences.html#navigable){#link-type-noopener:navigable}
whose [target
name](document-sequences.html#nav-target){#link-type-noopener:nav-target}
is \"`example`\"):

``` html
<a href=help.html target=example>Help!</a>
```

This creates a [top-level
traversable](document-sequences.html#top-level-traversable){#link-type-noopener:top-level-traversable-3}
with a non-[auxiliary browsing
context](document-sequences.html#auxiliary-browsing-context){#link-type-noopener:auxiliary-browsing-context-3}
(assuming the same thing):

``` html
<a href=help.html target=example rel=noopener>Help!</a>
```

These are equivalent and only navigate the [parent
navigable](document-sequences.html#nav-parent){#link-type-noopener:nav-parent}:

``` html
<a href=index.html target=_parent>Home</a>
```

``` html
<a href=index.html target=_parent rel=noopener>Home</a>
```
:::

##### [4.6.7.15]{.secno} Link type \"[`noreferrer`]{.dfn dfn-for="a/rel,area/rel,form/rel" dfn-type="attr-value"}\"[](#link-type-noreferrer){.self-link}

::::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Link_types/noreferrer](https://developer.mozilla.org/en-US/docs/Web/HTML/Link_types/noreferrer "The noreferrer keyword for the rel attribute of the <a>, <area>, and <form> elements instructs the browser, when navigating to the target resource, to omit the Referer header and otherwise leak no referrer information — and additionally to behave as if the noopener keyword were also specified.")

Support in all current engines.

::: support
[Firefox33+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome16+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)13+]{.edge .yes}[Internet Explorer[🔰
11]{title="Partial implementation."}]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet1.5+]{.samsunginternet_android .yes}[Opera
Android?]{.opera_android .unknown}
:::
::::

:::: feature
[Link_types/noreferrer](https://developer.mozilla.org/en-US/docs/Web/HTML/Link_types/noreferrer "The noreferrer keyword for the rel attribute of the <a>, <area>, and <form> elements instructs the browser, when navigating to the target resource, to omit the Referer header and otherwise leak no referrer information — and additionally to behave as if the noopener keyword were also specified.")

Support in all current engines.

::: support
[Firefox33+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome16+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)13+]{.edge .yes}[Internet Explorer[🔰
11]{title="Partial implementation."}]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet1.5+]{.samsunginternet_android .yes}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::::

The
[`noreferrer`{#link-type-noreferrer:link-type-noreferrer}](#link-type-noreferrer)
keyword may be used with
[`a`{#link-type-noreferrer:the-a-element}](text-level-semantics.html#the-a-element),
[`area`{#link-type-noreferrer:the-area-element}](image-maps.html#the-area-element),
and
[`form`{#link-type-noreferrer:the-form-element}](forms.html#the-form-element)
elements. This keyword does not create a
[hyperlink](#hyperlink){#link-type-noreferrer:hyperlink}, but
[annotates](#hyperlink-annotation){#link-type-noreferrer:hyperlink-annotation}
any other hyperlinks created by the element (the implied hyperlink, if
no other keywords create one).

It indicates that no referrer information is to be leaked when following
the link and also implies the
[`noopener`{#link-type-noreferrer:link-type-noopener}](#link-type-noopener)
keyword behavior under the same conditions.

See also the [processing model](#noreferrer-a-area-processing-model)
where referrer is directly manipulated.

`<a href="..." rel="noreferrer" target="_blank">` has the same behavior
as `<a href="..." rel="noreferrer noopener" target="_blank">`.

##### [4.6.7.16]{.secno} Link type \"[`opener`]{.dfn dfn-for="a/rel,area/rel,form/rel" dfn-type="attr-value"}\"[](#link-type-opener){.self-link}

The [`opener`{#link-type-opener:link-type-opener}](#link-type-opener)
keyword may be used with
[`a`{#link-type-opener:the-a-element}](text-level-semantics.html#the-a-element),
[`area`{#link-type-opener:the-area-element}](image-maps.html#the-area-element),
and
[`form`{#link-type-opener:the-form-element}](forms.html#the-form-element)
elements. This keyword does not create a
[hyperlink](#hyperlink){#link-type-opener:hyperlink}, but
[annotates](#hyperlink-annotation){#link-type-opener:hyperlink-annotation}
any other hyperlinks created by the element (the implied hyperlink, if
no other keywords create one).

The keyword indicates that any newly created [top-level
traversable](document-sequences.html#top-level-traversable){#link-type-opener:top-level-traversable}
which results from following the
[hyperlink](#hyperlink){#link-type-opener:hyperlink-2} will contain an
[auxiliary browsing
context](document-sequences.html#auxiliary-browsing-context){#link-type-opener:auxiliary-browsing-context}.

See also the [processing model](#opener-processing-model).

::: example
In the following example the
[`opener`{#link-type-opener:link-type-opener-2}](#link-type-opener) is
used to allow the help page popup to navigate its opener, e.g., in case
what the user is looking for can be found elsewhere. An alternative
might be to use a named target, rather than `_blank`, but this has the
potential to clash with existing names.

``` html
<a href="..." rel=opener target=_blank>Help!</a>
```
:::

##### [4.6.7.17]{.secno} Link type \"[`pingback`]{.dfn dfn-for="link/rel" dfn-type="attr-value"}\"[](#link-type-pingback){.self-link}

The
[`pingback`{#link-type-pingback:link-type-pingback}](#link-type-pingback)
keyword may be used with
[`link`{#link-type-pingback:the-link-element}](semantics.html#the-link-element)
elements. This keyword creates an [external resource
link](#external-resource-link){#link-type-pingback:external-resource-link}.
This keyword is [body-ok](#body-ok){#link-type-pingback:body-ok}.

For the semantics of the
[`pingback`{#link-type-pingback:link-type-pingback-2}](#link-type-pingback)
keyword, see Pingback 1.0. [\[PINGBACK\]](references.html#refsPINGBACK)

##### [4.6.7.18]{.secno} Link type \"[`preconnect`]{.dfn dfn-for="link/rel" dfn-type="attr-value"}\"[](#link-type-preconnect){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Link_types/preconnect](https://developer.mozilla.org/en-US/docs/Web/HTML/Link_types/preconnect "The preconnect keyword for the rel attribute of the <link> element is a hint to browsers that the user is likely to need resources from the target resource's origin, and therefore the browser can likely improve the user experience by preemptively initiating a connection to that origin. Preconnecting speeds up future loads from a given origin by preemptively performing part or all of the handshake (DNS+TCP for HTTP, and DNS+TCP+TLS for HTTPS origins).")

Support in all current engines.

::: support
[Firefox39+]{.firefox .yes}[Safari11.1+]{.safari
.yes}[Chrome46+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet4.0+]{.samsunginternet_android .yes}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

The
[`preconnect`{#link-type-preconnect:link-type-preconnect}](#link-type-preconnect)
keyword may be used with
[`link`{#link-type-preconnect:the-link-element}](semantics.html#the-link-element)
elements. This keyword creates an [external resource
link](#external-resource-link){#link-type-preconnect:external-resource-link}.
This keyword is [body-ok](#body-ok){#link-type-preconnect:body-ok}.

The
[`preconnect`{#link-type-preconnect:link-type-preconnect-2}](#link-type-preconnect)
keyword indicates that preemptively initiating a connection to the
[origin](browsers.html#concept-origin){#link-type-preconnect:concept-origin}
of the specified resource is likely to be beneficial, as it is highly
likely that the user will require resources located at that
[origin](browsers.html#concept-origin){#link-type-preconnect:concept-origin-2},
and the user experience would be improved by preempting the latency
costs associated with establishing the connection.

There is no default type for resources given by the
[`preconnect`{#link-type-preconnect:link-type-preconnect-3}](#link-type-preconnect)
keyword.

A user agent must not [delay the load
event](parsing.html#delay-the-load-event){#link-type-preconnect:delay-the-load-event}
for this link type.

The appropriate times to [fetch and
process](semantics.html#fetch-and-process-the-linked-resource){#link-type-preconnect:fetch-and-process-the-linked-resource}
this type of link are:

- When the [external resource
  link](#external-resource-link){#link-type-preconnect:external-resource-link-2}
  is created on a
  [`link`{#link-type-preconnect:the-link-element-2}](semantics.html#the-link-element)
  element that is already [browsing-context
  connected](infrastructure.html#browsing-context-connected){#link-type-preconnect:browsing-context-connected}.

- When the [external resource
  link](#external-resource-link){#link-type-preconnect:external-resource-link-3}\'s
  [`link`{#link-type-preconnect:the-link-element-3}](semantics.html#the-link-element)
  element [becomes browsing-context
  connected](infrastructure.html#becomes-browsing-context-connected){#link-type-preconnect:becomes-browsing-context-connected}.

- When the
  [`href`{#link-type-preconnect:attr-link-href}](semantics.html#attr-link-href)
  attribute of the
  [`link`{#link-type-preconnect:the-link-element-4}](semantics.html#the-link-element)
  element of an [external resource
  link](#external-resource-link){#link-type-preconnect:external-resource-link-4}
  that is already [browsing-context
  connected](infrastructure.html#browsing-context-connected){#link-type-preconnect:browsing-context-connected-2}
  is changed.

- When the
  [`crossorigin`{#link-type-preconnect:attr-link-crossorigin}](semantics.html#attr-link-crossorigin)
  attribute of the
  [`link`{#link-type-preconnect:the-link-element-5}](semantics.html#the-link-element)
  element of an [external resource
  link](#external-resource-link){#link-type-preconnect:external-resource-link-5}
  that is already [browsing-context
  connected](infrastructure.html#browsing-context-connected){#link-type-preconnect:browsing-context-connected-3}
  is set, changed, or removed.

The [fetch and process the linked
resource](semantics.html#fetch-and-process-the-linked-resource){#link-type-preconnect:fetch-and-process-the-linked-resource-2}
steps for this type of linked resource, given a
[`link`{#link-type-preconnect:the-link-element-6}](semantics.html#the-link-element)
element `el`{.variable}, are to [create link
options](semantics.html#create-link-options-from-element){#link-type-preconnect:create-link-options-from-element}
from `el`{.variable} and to
[preconnect](#preconnect){#link-type-preconnect:preconnect} given the
result.

The [process a link
header](semantics.html#process-a-link-header){#link-type-preconnect:process-a-link-header}
step for this type of linked resource given a [link processing
options](semantics.html#link-processing-options){#link-type-preconnect:link-processing-options}
`options`{.variable} are to
[preconnect](#preconnect){#link-type-preconnect:preconnect-2} given
`options`{.variable}.

To [preconnect]{#preconnect .dfn} given a [link processing
options](semantics.html#link-processing-options){#link-type-preconnect:link-processing-options-2}
`options`{.variable}:

1.  If `options`{.variable}\'s
    [href](semantics.html#link-options-href){#link-type-preconnect:link-options-href}
    is an empty string, return.

2.  Let `url`{.variable} be the result of [encoding-parsing a
    URL](urls-and-fetching.html#encoding-parsing-a-url){#link-type-preconnect:encoding-parsing-a-url}
    given `options`{.variable}\'s
    [href](semantics.html#link-options-href){#link-type-preconnect:link-options-href-2},
    relative to `options`{.variable}\'s [base
    URL](semantics.html#link-options-base-url){#link-type-preconnect:link-options-base-url}.

    Passing the base URL instead of a document or environment is tracked
    by [issue #9715](https://github.com/whatwg/html/issues/9715).

3.  If `url`{.variable} is failure, then return.

4.  If `url`{.variable}\'s
    [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#link-type-preconnect:concept-url-scheme
    x-internal="concept-url-scheme"} is not an [HTTP(S)
    scheme](https://fetch.spec.whatwg.org/#http-scheme){#link-type-preconnect:http(s)-scheme
    x-internal="http(s)-scheme"}, then return.

5.  Let `partitionKey`{.variable} be the result of [determining the
    network partition
    key](https://fetch.spec.whatwg.org/#determine-the-network-partition-key){#link-type-preconnect:determine-the-network-partition-key
    x-internal="determine-the-network-partition-key"} given
    `options`{.variable}\'s
    [environment](semantics.html#link-options-environment){#link-type-preconnect:link-options-environment}.

6.  Let `useCredentials`{.variable} be true.

7.  If `options`{.variable}\'s
    [crossorigin](semantics.html#link-options-crossorigin){#link-type-preconnect:link-options-crossorigin}
    is
    [Anonymous](urls-and-fetching.html#attr-crossorigin-anonymous){#link-type-preconnect:attr-crossorigin-anonymous}
    and `options`{.variable}\'s
    [origin](semantics.html#link-options-origin){#link-type-preconnect:link-options-origin}
    does not have the [same
    origin](browsers.html#same-origin){#link-type-preconnect:same-origin}
    as `url`{.variable}\'s
    [origin](https://url.spec.whatwg.org/#concept-url-origin){#link-type-preconnect:concept-url-origin
    x-internal="concept-url-origin"}, then set
    `useCredentials`{.variable} to false.

8.  The user agent should [obtain a
    connection](https://fetch.spec.whatwg.org/#concept-connection-obtain){#link-type-preconnect:obtain-a-connection
    x-internal="obtain-a-connection"} given `partitionKey`{.variable},
    `url`{.variable}\'s
    [origin](https://url.spec.whatwg.org/#concept-url-origin){#link-type-preconnect:concept-url-origin-2
    x-internal="concept-url-origin"}, and `useCredentials`{.variable}.

    This connection is obtained but not used directly. It will remain in
    the [connection
    pool](https://fetch.spec.whatwg.org/#concept-connection-pool){#link-type-preconnect:connection-pool
    x-internal="connection-pool"} for subsequent use.

    The user agent should attempt to initiate a preconnect and perform
    the full connection handshake (DNS+TCP for HTTP, and DNS+TCP+TLS for
    HTTPS origins) whenever possible, but is allowed to elect to perform
    a partial handshake (DNS only for HTTP, and DNS or DNS+TCP for HTTPS
    origins), or skip it entirely, due to resource constraints or other
    reasons.

    The optimal number of connections per origin is dependent on the
    negotiated protocol, users current connectivity profile, available
    device resources, global connection limits, and other context
    specific variables. As a result, the decision for how many
    connections should be opened is deferred to the user agent.

##### [4.6.7.19]{.secno} Link type \"[`prefetch`]{.dfn dfn-for="link/rel" dfn-type="attr-value"}\"[](#link-type-prefetch){.self-link}

::::: {.mdn-anno .wrapped}
MDN

:::: feature
[Link_types/prefetch](https://developer.mozilla.org/en-US/docs/Web/HTML/Link_types/prefetch "The prefetch keyword for the rel attribute of the <link> element is a hint to browsers that the user is likely to need the target resource for future navigations, and therefore the browser can likely improve the user experience by preemptively fetching and caching the resource.")

::: support
[Firefox2+]{.firefox .yes}[SafariNo]{.safari .no}[Chrome8+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer11]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet1.5+]{.samsunginternet_android .yes}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

The
[`prefetch`{#link-type-prefetch:link-type-prefetch}](#link-type-prefetch)
keyword may be used with
[`link`{#link-type-prefetch:the-link-element}](semantics.html#the-link-element)
elements. This keyword creates an [external resource
link](#external-resource-link){#link-type-prefetch:external-resource-link}.
This keyword is [body-ok](#body-ok){#link-type-prefetch:body-ok}.

The
[`prefetch`{#link-type-prefetch:link-type-prefetch-2}](#link-type-prefetch)
keyword indicates that preemptively
[fetching](https://fetch.spec.whatwg.org/#concept-fetch){#link-type-prefetch:concept-fetch
x-internal="concept-fetch"} and caching the specified resource or
same-site document is likely to be beneficial, as it is highly likely
that the user will require this resource for future navigations.

There is no default type for resources given by the
[`prefetch`{#link-type-prefetch:link-type-prefetch-3}](#link-type-prefetch)
keyword.

The appropriate times to [fetch and
process](semantics.html#fetch-and-process-the-linked-resource){#link-type-prefetch:fetch-and-process-the-linked-resource}
this type of link are:

- When the [external resource
  link](#external-resource-link){#link-type-prefetch:external-resource-link-2}
  is created on a
  [`link`{#link-type-prefetch:the-link-element-2}](semantics.html#the-link-element)
  element that is already [browsing-context
  connected](infrastructure.html#browsing-context-connected){#link-type-prefetch:browsing-context-connected}.

- When the [external resource
  link](#external-resource-link){#link-type-prefetch:external-resource-link-3}\'s
  [`link`{#link-type-prefetch:the-link-element-3}](semantics.html#the-link-element)
  element [becomes browsing-context
  connected](infrastructure.html#becomes-browsing-context-connected){#link-type-prefetch:becomes-browsing-context-connected}.

- When the
  [`href`{#link-type-prefetch:attr-link-href}](semantics.html#attr-link-href)
  attribute of the
  [`link`{#link-type-prefetch:the-link-element-4}](semantics.html#the-link-element)
  element of an [external resource
  link](#external-resource-link){#link-type-prefetch:external-resource-link-4}
  that is already [browsing-context
  connected](infrastructure.html#browsing-context-connected){#link-type-prefetch:browsing-context-connected-2}
  is changed.

- When the
  [`crossorigin`{#link-type-prefetch:attr-link-crossorigin}](semantics.html#attr-link-crossorigin)
  attribute of the
  [`link`{#link-type-prefetch:the-link-element-5}](semantics.html#the-link-element)
  element of an [external resource
  link](#external-resource-link){#link-type-prefetch:external-resource-link-5}
  that is already [browsing-context
  connected](infrastructure.html#browsing-context-connected){#link-type-prefetch:browsing-context-connected-3}
  is set, changed, or removed.

The [fetch and process the linked
resource](semantics.html#fetch-and-process-the-linked-resource){#link-type-prefetch:fetch-and-process-the-linked-resource-2}
algorithm for
[`prefetch`{#link-type-prefetch:link-type-prefetch-4}](#link-type-prefetch)
links, given a
[`link`{#link-type-prefetch:the-link-element-6}](semantics.html#the-link-element)
element `el`{.variable}, is as follows:

1.  If `el`{.variable}\'s
    [`href`{#link-type-prefetch:attr-link-href-2}](semantics.html#attr-link-href)
    attribute\'s value is the empty string, then return.

2.  Let `options`{.variable} be the result of [creating link
    options](semantics.html#create-link-options-from-element){#link-type-prefetch:create-link-options-from-element}
    from `el`{.variable}.

3.  Set `options`{.variable}\'s
    [destination](semantics.html#link-options-destination){#link-type-prefetch:link-options-destination}
    to the empty string.

4.  Let `request`{.variable} be the result of [creating a link
    request](semantics.html#create-a-link-request){#link-type-prefetch:create-a-link-request}
    given `options`{.variable}.

5.  If `request`{.variable} is null, then return.

6.  Set `request`{.variable}\'s
    [initiator](https://fetch.spec.whatwg.org/#concept-request-initiator){#link-type-prefetch:concept-request-initiator
    x-internal="concept-request-initiator"} to \"`prefetch`\".

7.  Let `processPrefetchResponse`{.variable} be the following steps
    given a
    [response](https://fetch.spec.whatwg.org/#concept-response){#link-type-prefetch:concept-response
    x-internal="concept-response"} `response`{.variable} and null,
    failure, or a [byte
    sequence](https://infra.spec.whatwg.org/#byte-sequence){#link-type-prefetch:byte-sequence
    x-internal="byte-sequence"} `bytesOrNull`{.variable}:

    1.  If `response`{.variable} is a [network
        error](https://fetch.spec.whatwg.org/#concept-network-error){#link-type-prefetch:network-error
        x-internal="network-error"}, [fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#link-type-prefetch:concept-event-fire
        x-internal="concept-event-fire"} named
        [`error`{#link-type-prefetch:event-error}](indices.html#event-error)
        at `el`{.variable}.

    2.  Otherwise, [fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#link-type-prefetch:concept-event-fire-2
        x-internal="concept-event-fire"} named
        [`load`{#link-type-prefetch:event-load}](indices.html#event-load)
        at `el`{.variable}.

8.  The user agent should
    [fetch](https://fetch.spec.whatwg.org/#concept-fetch){#link-type-prefetch:concept-fetch-2
    x-internal="concept-fetch"} `request`{.variable}, with
    *[processResponseConsumeBody](https://fetch.spec.whatwg.org/#process-response-end-of-body){x-internal="processresponseconsumebody"}*
    set to `processPrefetchResponse`{.variable}. User agents may delay
    the fetching of `request`{.variable} to prioritize other requests
    that are necessary for the current document.

The [process a link
header](semantics.html#process-a-link-header){#link-type-prefetch:process-a-link-header}
steps for this type of linked resource are to do nothing.

##### [4.6.7.20]{.secno} Link type \"[`preload`]{.dfn dfn-for="link/rel" dfn-type="attr-value"}\"[](#link-type-preload){.self-link}

::::: {.mdn-anno .wrapped}
**⚠**MDN

:::: feature
[Link_types/preload](https://developer.mozilla.org/en-US/docs/Web/HTML/Link_types/preload "The preload value of the <link> element's rel attribute lets you declare fetch requests in the HTML's <head>, specifying resources that your page will need very soon, which you want to start loading early in the page lifecycle, before browsers' main rendering machinery kicks in. This ensures they are available earlier and are less likely to block the page's render, improving performance. Even though the name contains the term load, it doesn't load and execute the script but only schedules it to be downloaded and cached with a higher priority.")

Support in one engine only.

::: support
[Firefox85+]{.firefox .yes}[Safari?]{.safari .unknown}[Chrome[🔰
50+]{title="Partial implementation."}]{.chrome .yes}

------------------------------------------------------------------------

[Opera37+]{.opera .yes}[Edge[🔰
79+]{title="Partial implementation."}]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)No]{.edge .no}[Internet Explorer?]{.ie .unknown}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android50+]{.webview_android .yes}[Samsung
Internet5.0+]{.samsunginternet_android .yes}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

The
[`preload`{#link-type-preload:link-type-preload}](#link-type-preload)
keyword may be used with
[`link`{#link-type-preload:the-link-element}](semantics.html#the-link-element)
elements. This keyword creates an [external resource
link](#external-resource-link){#link-type-preload:external-resource-link}.
This keyword is [body-ok](#body-ok){#link-type-preload:body-ok}.

The
[`preload`{#link-type-preload:link-type-preload-2}](#link-type-preload)
keyword indicates that the user agent will preemptively
[fetch](https://fetch.spec.whatwg.org/#concept-fetch){#link-type-preload:concept-fetch
x-internal="concept-fetch"} and cache the specified resource according
to the [potential
destination](https://fetch.spec.whatwg.org/#concept-potential-destination){#link-type-preload:concept-potential-destination
x-internal="concept-potential-destination"} given by the
[`as`{#link-type-preload:attr-link-as}](semantics.html#attr-link-as)
attribute, and the
[priority](https://fetch.spec.whatwg.org/#request-priority){#link-type-preload:concept-request-priority
x-internal="concept-request-priority"} given by the
[`fetchpriority`{#link-type-preload:attr-link-fetchpriority}](semantics.html#attr-link-fetchpriority)
attribute, as it is highly likely that the user will require this
resource for the current navigation.

User-agents might perform additional operations when a resource is
loaded, such as preemptively [decoding
images](embedded-content.html#dom-img-decode){#link-type-preload:dom-img-decode}
or [creating
stylesheets](https://drafts.csswg.org/cssom/#create-a-css-style-sheet){#link-type-preload:create-a-css-style-sheet
x-internal="create-a-css-style-sheet"}. However, these additional
operations cannot have observable effects.

There is no default type for resources given by the
[`preload`{#link-type-preload:link-type-preload-3}](#link-type-preload)
keyword.

A user agent must not [delay the load
event](parsing.html#delay-the-load-event){#link-type-preload:delay-the-load-event}
for this link type.

The appropriate times to [fetch and process the linked
resource](semantics.html#fetch-and-process-the-linked-resource){#link-type-preload:fetch-and-process-the-linked-resource}
for such a link are:

- When the [external resource
  link](#external-resource-link){#link-type-preload:external-resource-link-2}
  is created on a
  [`link`{#link-type-preload:the-link-element-2}](semantics.html#the-link-element)
  element that is already [browsing-context
  connected](infrastructure.html#browsing-context-connected){#link-type-preload:browsing-context-connected}.

- When the [external resource
  link](#external-resource-link){#link-type-preload:external-resource-link-3}\'s
  [`link`{#link-type-preload:the-link-element-3}](semantics.html#the-link-element)
  element [becomes browsing-context
  connected](infrastructure.html#becomes-browsing-context-connected){#link-type-preload:becomes-browsing-context-connected}.

- When the
  [`href`{#link-type-preload:attr-link-href}](semantics.html#attr-link-href)
  attribute of the
  [`link`{#link-type-preload:the-link-element-4}](semantics.html#the-link-element)
  element of an [external resource
  link](#external-resource-link){#link-type-preload:external-resource-link-4}
  that is already [browsing-context
  connected](infrastructure.html#browsing-context-connected){#link-type-preload:browsing-context-connected-2}
  is changed.

- When the
  [`as`{#link-type-preload:attr-link-as-2}](semantics.html#attr-link-as)
  attribute of the
  [`link`{#link-type-preload:the-link-element-5}](semantics.html#the-link-element)
  element of an [external resource
  link](#external-resource-link){#link-type-preload:external-resource-link-5}
  that is already [browsing-context
  connected](infrastructure.html#browsing-context-connected){#link-type-preload:browsing-context-connected-3}
  is changed.

- When the
  [`type`{#link-type-preload:attr-link-type}](semantics.html#attr-link-type)
  attribute of the
  [`link`{#link-type-preload:the-link-element-6}](semantics.html#the-link-element)
  element of an [external resource
  link](#external-resource-link){#link-type-preload:external-resource-link-6}
  that is already [browsing-context
  connected](infrastructure.html#browsing-context-connected){#link-type-preload:browsing-context-connected-4},
  but was previously not obtained due to the
  [`type`{#link-type-preload:attr-link-type-2}](semantics.html#attr-link-type)
  attribute specifying an unsupported type for the request
  [destination](https://fetch.spec.whatwg.org/#concept-request-destination){#link-type-preload:concept-request-destination
  x-internal="concept-request-destination"}, is set, removed, or
  changed.

- When the
  [`media`{#link-type-preload:attr-link-media}](semantics.html#attr-link-media)
  attribute of the
  [`link`{#link-type-preload:the-link-element-7}](semantics.html#the-link-element)
  element of an [external resource
  link](#external-resource-link){#link-type-preload:external-resource-link-7}
  that is already [browsing-context
  connected](infrastructure.html#browsing-context-connected){#link-type-preload:browsing-context-connected-5},
  but was previously not obtained due to the
  [`media`{#link-type-preload:attr-link-media-2}](semantics.html#attr-link-media)
  attribute not [matching the
  environment](common-microsyntaxes.html#matches-the-environment){#link-type-preload:matches-the-environment},
  is changed or removed.

A [`Document`{#link-type-preload:document}](dom.html#document) has a
[map of preloaded resources]{#map-of-preloaded-resources .dfn}, which is
an [ordered
map](https://infra.spec.whatwg.org/#ordered-map){#link-type-preload:ordered-map
x-internal="ordered-map"}, initially empty.

A [preload key]{#preload-key .dfn} is a
[struct](https://infra.spec.whatwg.org/#struct){#link-type-preload:struct
x-internal="struct"}. It has the following
[items](https://infra.spec.whatwg.org/#struct-item){#link-type-preload:struct-item
x-internal="struct-item"}:

[URL]{#preload-url .dfn}
:   A
    [URL](https://url.spec.whatwg.org/#concept-url){#link-type-preload:url
    x-internal="url"}

[destination]{#preload-destination .dfn}
:   A string

[mode]{#preload-mode .dfn}
:   A [request
    mode](https://fetch.spec.whatwg.org/#concept-request-mode){#link-type-preload:concept-request-mode
    x-internal="concept-request-mode"}, either \"`same-origin`\",
    \"`cors`\", or \"`no-cors`\"

[credentials mode]{#preload-credentials-mode .dfn}
:   A [credentials
    mode](https://fetch.spec.whatwg.org/#concept-request-credentials-mode){#link-type-preload:concept-request-credentials-mode
    x-internal="concept-request-credentials-mode"}

A [preload entry]{#preload-entry .dfn} is a
[struct](https://infra.spec.whatwg.org/#struct){#link-type-preload:struct-2
x-internal="struct"}. It has the following
[items](https://infra.spec.whatwg.org/#struct-item){#link-type-preload:struct-item-2
x-internal="struct-item"}:

[integrity metadata]{#preload-integrity-metadata .dfn}
:   A string

[response]{#preload-response .dfn}
:   Null or a
    [response](https://fetch.spec.whatwg.org/#concept-response){#link-type-preload:concept-response
    x-internal="concept-response"}

[on response available]{#preload-on-response-available .dfn}
:   Null, or an algorithm accepting a
    [response](https://fetch.spec.whatwg.org/#concept-response){#link-type-preload:concept-response-2
    x-internal="concept-response"} or null

To [consume a preloaded resource]{#consume-a-preloaded-resource .dfn
export=""} for
[`Window`{#link-type-preload:window}](nav-history-apis.html#window)
`window`{.variable}, given a
[URL](https://url.spec.whatwg.org/#concept-url){#link-type-preload:url-2
x-internal="url"} `url`{.variable}, a string `destination`{.variable}, a
string `mode`{.variable}, a string `credentialsMode`{.variable}, a
string `integrityMetadata`{.variable}, and
`onResponseAvailable`{.variable}, which is an algorithm accepting a
[response](https://fetch.spec.whatwg.org/#concept-response){#link-type-preload:concept-response-3
x-internal="concept-response"}:

1.  Let `key`{.variable} be a [preload
    key](#preload-key){#link-type-preload:preload-key} whose
    [URL](#preload-url){#link-type-preload:preload-url} is
    `url`{.variable},
    [destination](#preload-destination){#link-type-preload:preload-destination}
    is `destination`{.variable},
    [mode](#preload-mode){#link-type-preload:preload-mode} is
    `mode`{.variable}, and [credentials
    mode](#preload-credentials-mode){#link-type-preload:preload-credentials-mode}
    is `credentialsMode`{.variable}.

2.  Let `preloads`{.variable} be `window`{.variable}\'s [associated
    `Document`](nav-history-apis.html#concept-document-window){#link-type-preload:concept-document-window}\'s
    [map of preloaded
    resources](#map-of-preloaded-resources){#link-type-preload:map-of-preloaded-resources}.

3.  If `key`{.variable} does not
    [exist](https://infra.spec.whatwg.org/#map-exists){#link-type-preload:map-exists
    x-internal="map-exists"} in `preloads`{.variable}, then return
    false.

4.  Let `entry`{.variable} be `preloads`{.variable}\[`key`{.variable}\].

5.  Let `consumerIntegrityMetadata`{.variable} be the result of
    [parsing](https://w3c.github.io/webappsec-subresource-integrity/#parse-metadata){#link-type-preload:parse-integrity-metadata
    x-internal="parse-integrity-metadata"}
    `integrityMetadata`{.variable}.

6.  Let `preloadIntegrityMetadata`{.variable} be the result of
    [parsing](https://w3c.github.io/webappsec-subresource-integrity/#parse-metadata){#link-type-preload:parse-integrity-metadata-2
    x-internal="parse-integrity-metadata"} `entry`{.variable}\'s
    [integrity
    metadata](#preload-integrity-metadata){#link-type-preload:preload-integrity-metadata}.

7.  If none of the following conditions apply:

    - `consumerIntegrityMetadata`{.variable} is `no metadata`;

    - `consumerIntegrityMetadata`{.variable} is equal to
      `preloadIntegrityMetadata`{.variable}; or

      This comparison would ignore unknown integrity options. See [issue
      #116.](https://github.com/w3c/webappsec-subresource-integrity/issues/116)

    then return false.

    A mismatch in integrity metadata between the preload and the
    consumer, even if both match the data, would lead to an additional
    fetch from the network.

    It is important that [network
    errors](https://fetch.spec.whatwg.org/#concept-network-error){#link-type-preload:network-error
    x-internal="network-error"} are added to the preload cache so that
    if a preload request results in an error, the erroneous response
    isn\'t re-requested from the network later. This also has security
    implications; consider the case where a developer specifies
    subresource integrity metadata on a preload request, but not the
    following resource request. If the preload request fails subresource
    integrity verification and is discarded, the resource request will
    fetch and consume a potentially-malicious response from the network
    without verifying its integrity. [\[SRI\]](references.html#refsSRI)

8.  [Remove](https://infra.spec.whatwg.org/#map-remove){#link-type-preload:map-remove
    x-internal="map-remove"} `preloads`{.variable}\[`key`{.variable}\].

9.  If `entry`{.variable}\'s
    [response](#preload-response){#link-type-preload:preload-response}
    is null, then set `entry`{.variable}\'s [on response
    available](#preload-on-response-available){#link-type-preload:preload-on-response-available}
    to `onResponseAvailable`{.variable}.

10. Otherwise, call `onResponseAvailable`{.variable} with
    `entry`{.variable}\'s
    [response](#preload-response){#link-type-preload:preload-response-2}.

11. Return true.

For the purposes of this section, a string `type`{.variable}
[matches]{#match-preload-type .dfn} a string `destination`{.variable} if
the following algorithm returns true:

1.  If `type`{.variable} is an empty string, then return true.

2.  If `destination`{.variable} is \"`fetch`\", then return true.

3.  Let `mimeTypeRecord`{.variable} be the result of
    [parsing](https://mimesniff.spec.whatwg.org/#parse-a-mime-type){#link-type-preload:parse-a-mime-type
    x-internal="parse-a-mime-type"} `type`{.variable}.

4.  If `mimeTypeRecord`{.variable} is failure, then return false.

5.  If `mimeTypeRecord`{.variable} is not [supported by the user
    agent](https://mimesniff.spec.whatwg.org/#supported-by-the-user-agent){#link-type-preload:is-mime-type-supported-by-the-user-agent
    x-internal="is-mime-type-supported-by-the-user-agent"}, then return
    false.

6.  If any of the following are true:

    - `destination`{.variable} is \"`audio`\" or \"`video`\", and
      `mimeTypeRecord`{.variable} is an [audio or video MIME
      type](https://mimesniff.spec.whatwg.org/#audio-or-video-mime-type){#link-type-preload:audio-or-video-mime-type
      x-internal="audio-or-video-mime-type"};

    - `destination`{.variable} is a [script-like
      destination](https://fetch.spec.whatwg.org/#request-destination-script-like){#link-type-preload:concept-script-like-destination
      x-internal="concept-script-like-destination"} and
      `mimeTypeRecord`{.variable} is a [JavaScript MIME
      type](https://mimesniff.spec.whatwg.org/#javascript-mime-type){#link-type-preload:javascript-mime-type
      x-internal="javascript-mime-type"};

    - `destination`{.variable} is \"`image`\" and
      `mimeTypeRecord`{.variable} is an [image MIME
      type](https://mimesniff.spec.whatwg.org/#image-mime-type){#link-type-preload:image-mime-type
      x-internal="image-mime-type"};

    - `destination`{.variable} is \"`font`\" and
      `mimeTypeRecord`{.variable} is a [font MIME
      type](https://mimesniff.spec.whatwg.org/#font-mime-type){#link-type-preload:font-mime-type
      x-internal="font-mime-type"};

    - `destination`{.variable} is \"`json`\" and
      `mimeTypeRecord`{.variable} is a [JSON MIME
      type](https://mimesniff.spec.whatwg.org/#json-mime-type){#link-type-preload:json-mime-type
      x-internal="json-mime-type"};

    - `destination`{.variable} is \"`style`\" and
      `mimeTypeRecord`{.variable}\'s
      [essence](https://mimesniff.spec.whatwg.org/#mime-type-essence){#link-type-preload:mime-type-essence
      x-internal="mime-type-essence"} is
      [`text/css`{#link-type-preload:text/css}](indices.html#text/css);
      or

    - `destination`{.variable} is \"`track`\" and
      `mimeTypeRecord`{.variable}\'s
      [essence](https://mimesniff.spec.whatwg.org/#mime-type-essence){#link-type-preload:mime-type-essence-2
      x-internal="mime-type-essence"} is
      [`text/vtt`{#link-type-preload:text/vtt}](indices.html#text/vtt),

    then return true.

7.  Return false.

To [create a preload key]{#create-a-preload-key .dfn} for a
[request](https://fetch.spec.whatwg.org/#concept-request){#link-type-preload:concept-request
x-internal="concept-request"} `request`{.variable}, return a new
[preload key](#preload-key){#link-type-preload:preload-key-2} whose
[URL](#preload-url){#link-type-preload:preload-url-2} is
`request`{.variable}\'s
[URL](https://fetch.spec.whatwg.org/#concept-request-url){#link-type-preload:concept-request-url
x-internal="concept-request-url"},
[destination](#preload-destination){#link-type-preload:preload-destination-2}
is `request`{.variable}\'s
[destination](https://fetch.spec.whatwg.org/#concept-request-destination){#link-type-preload:concept-request-destination-2
x-internal="concept-request-destination"},
[mode](#preload-mode){#link-type-preload:preload-mode-2} is
`request`{.variable}\'s
[mode](https://fetch.spec.whatwg.org/#concept-request-mode){#link-type-preload:concept-request-mode-2
x-internal="concept-request-mode"}, and [credentials
mode](#preload-credentials-mode){#link-type-preload:preload-credentials-mode-2}
is `request`{.variable}\'s [credentials
mode](https://fetch.spec.whatwg.org/#concept-request-credentials-mode){#link-type-preload:concept-request-credentials-mode-2
x-internal="concept-request-credentials-mode"}.

To [translate a preload destination]{#translate-a-preload-destination
.dfn} given a string `destination`{.variable}:

1.  If `destination`{.variable} is not \"`fetch`\", \"`font`\",
    \"`image`\", \"`script`\", \"`style`\", or \"`track`\", then return
    null.

2.  Return the result of
    [translating](https://fetch.spec.whatwg.org/#concept-potential-destination-translate){#link-type-preload:concept-potential-destination-translate
    x-internal="concept-potential-destination-translate"}
    `destination`{.variable}.

To [preload]{#preload .dfn} given a [link processing
options](semantics.html#link-processing-options){#link-type-preload:link-processing-options}
`options`{.variable} and an optional `processResponse`{.variable}, which
is an algorithm accepting a
[response](https://fetch.spec.whatwg.org/#concept-response){#link-type-preload:concept-response-4
x-internal="concept-response"}:

1.  If `options`{.variable}\'s
    [type](semantics.html#link-options-type){#link-type-preload:link-options-type}
    doesn\'t
    [match](#match-preload-type){#link-type-preload:match-preload-type}
    `options`{.variable}\'s
    [destination](semantics.html#link-options-destination){#link-type-preload:link-options-destination},
    then return.

2.  If `options`{.variable}\'s
    [destination](semantics.html#link-options-destination){#link-type-preload:link-options-destination-2}
    is \"`image`\" and `options`{.variable}\'s [source
    set](semantics.html#link-options-source-set){#link-type-preload:link-options-source-set}
    is not null, then set `options`{.variable}\'s
    [href](semantics.html#link-options-href){#link-type-preload:link-options-href}
    to the result of [selecting an image
    source](images.html#select-an-image-source-from-a-source-set){#link-type-preload:select-an-image-source-from-a-source-set}
    from `options`{.variable}\'s [source
    set](semantics.html#link-options-source-set){#link-type-preload:link-options-source-set-2}.

3.  Let `request`{.variable} be the result of [creating a link
    request](semantics.html#create-a-link-request){#link-type-preload:create-a-link-request}
    given `options`{.variable}.

4.  If `request`{.variable} is null, then return.

5.  Let `unsafeEndTime`{.variable} be 0.

6.  Let `entry`{.variable} be a new [preload
    entry](#preload-entry){#link-type-preload:preload-entry} whose
    [integrity
    metadata](#preload-integrity-metadata){#link-type-preload:preload-integrity-metadata-2}
    is `options`{.variable}\'s
    [integrity](semantics.html#link-options-integrity){#link-type-preload:link-options-integrity}.

7.  Let `key`{.variable} be the result of [creating a preload
    key](#create-a-preload-key){#link-type-preload:create-a-preload-key}
    given `request`{.variable}.

8.  If `options`{.variable}\'s
    [document](semantics.html#link-options-document){#link-type-preload:link-options-document}
    is \"`pending`\", then set `request`{.variable}\'s [initiator
    type](https://fetch.spec.whatwg.org/#request-initiator-type){#link-type-preload:concept-request-initiator-type
    x-internal="concept-request-initiator-type"} to \"`early hint`\".

9.  Let `controller`{.variable} be null.

10. Let `reportTiming`{.variable} given a
    [`Document`{#link-type-preload:document-2}](dom.html#document)
    `document`{.variable} be to [report
    timing](https://fetch.spec.whatwg.org/#finalize-and-report-timing){#link-type-preload:report-timing
    x-internal="report-timing"} for `controller`{.variable} given
    `document`{.variable}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#link-type-preload:concept-relevant-global}.

11. Set `controller`{.variable} to the result of
    [fetching](https://fetch.spec.whatwg.org/#concept-fetch){#link-type-preload:concept-fetch-2
    x-internal="concept-fetch"} `request`{.variable}, with
    *[processResponseConsumeBody](https://fetch.spec.whatwg.org/#process-response-end-of-body){x-internal="processresponseconsumebody"}*
    set to the following steps given a
    [response](https://fetch.spec.whatwg.org/#concept-response){#link-type-preload:concept-response-5
    x-internal="concept-response"} `response`{.variable} and null,
    failure, or a [byte
    sequence](https://infra.spec.whatwg.org/#byte-sequence){#link-type-preload:byte-sequence
    x-internal="byte-sequence"} `bodyBytes`{.variable}:

    1.  If `bodyBytes`{.variable} is a [byte
        sequence](https://infra.spec.whatwg.org/#byte-sequence){#link-type-preload:byte-sequence-2
        x-internal="byte-sequence"}, then set `response`{.variable}\'s
        [body](https://fetch.spec.whatwg.org/#concept-response-body){#link-type-preload:concept-response-body
        x-internal="concept-response-body"} to `bodyBytes`{.variable}
        [as a
        body](https://fetch.spec.whatwg.org/#byte-sequence-as-a-body){#link-type-preload:as-a-body
        x-internal="as-a-body"}.

        By using
        *[processResponseConsumeBody](https://fetch.spec.whatwg.org/#process-response-end-of-body){x-internal="processresponseconsumebody"}*,
        we have
        [extracted](https://fetch.spec.whatwg.org/#bodyinit-safely-extract){#link-type-preload:body-safely-extract
        x-internal="body-safely-extract"} the entire
        [body](https://fetch.spec.whatwg.org/#concept-response-body){#link-type-preload:concept-response-body-2
        x-internal="concept-response-body"}. This is necessary to ensure
        the preloader loads the entire body from the network, regardless
        of whether the preload will be consumed (which is uncertain at
        this point). This step then resets the request\'s body to a new
        body containing the same bytes, so that other specifications can
        read from it at the time of actual consumption, despite us
        having already done so once.

    2.  Otherwise, set `response`{.variable} to a [network
        error](https://fetch.spec.whatwg.org/#concept-network-error){#link-type-preload:network-error-2
        x-internal="network-error"}.

    3.  Set `unsafeEndTime`{.variable} to the [unsafe shared current
        time](https://w3c.github.io/hr-time/#dfn-unsafe-shared-current-time){#link-type-preload:unsafe-shared-current-time
        x-internal="unsafe-shared-current-time"}.

    4.  If `options`{.variable}\'s
        [document](semantics.html#link-options-document){#link-type-preload:link-options-document-2}
        is not null, then call `reportTiming`{.variable} given
        `options`{.variable}\'s
        [document](semantics.html#link-options-document){#link-type-preload:link-options-document-3}.

    5.  If `entry`{.variable}\'s [on response
        available](#preload-on-response-available){#link-type-preload:preload-on-response-available-2}
        is null, then set `entry`{.variable}\'s
        [response](#preload-response){#link-type-preload:preload-response-3}
        to `response`{.variable}; otherwise call `entry`{.variable}\'s
        [on response
        available](#preload-on-response-available){#link-type-preload:preload-on-response-available-3}
        given `response`{.variable}.

    6.  If `processResponse`{.variable} is given, then call
        `processResponse`{.variable} with `response`{.variable}.

12. Let `commit`{.variable} be the following steps given a
    [`Document`{#link-type-preload:document-3}](dom.html#document)
    `document`{.variable}:

    1.  If `entry`{.variable}\'s
        [response](#preload-response){#link-type-preload:preload-response-4}
        is not null, then call `reportTiming`{.variable} given
        `document`{.variable}.

    2.  Set `document`{.variable}\'s [map of preloaded
        resources](#map-of-preloaded-resources){#link-type-preload:map-of-preloaded-resources-2}\[`key`{.variable}\]
        to `entry`{.variable}.

13. If `options`{.variable}\'s
    [document](semantics.html#link-options-document){#link-type-preload:link-options-document-4}
    is null, then set `options`{.variable}\'s [on document
    ready](semantics.html#link-options-on-document-ready){#link-type-preload:link-options-on-document-ready}
    to `commit`{.variable}. Otherwise, call `commit`{.variable} with
    `options`{.variable}\'s
    [document](semantics.html#link-options-document){#link-type-preload:link-options-document-5}.

The [fetch and process the linked
resource](semantics.html#fetch-and-process-the-linked-resource){#link-type-preload:fetch-and-process-the-linked-resource-2}
steps for this type of linked resource, given a
[`link`{#link-type-preload:the-link-element-8}](semantics.html#the-link-element)
element `el`{.variable}, are:

1.  [Update the source
    set](images.html#update-the-source-set){#link-type-preload:update-the-source-set}
    for `el`{.variable}.

2.  Let `options`{.variable} be the result of [creating link
    options](semantics.html#create-link-options-from-element){#link-type-preload:create-link-options-from-element}
    from `el`{.variable}.

3.  [Preload](#preload){#link-type-preload:preload}
    `options`{.variable}, with the following steps given a
    [response](https://fetch.spec.whatwg.org/#concept-response){#link-type-preload:concept-response-6
    x-internal="concept-response"} `response`{.variable}:

    1.  If `response`{.variable} is a [network
        error](https://fetch.spec.whatwg.org/#concept-network-error){#link-type-preload:network-error-3
        x-internal="network-error"}, [fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#link-type-preload:concept-event-fire
        x-internal="concept-event-fire"} named
        [`error`{#link-type-preload:event-error}](indices.html#event-error)
        at `el`{.variable}. Otherwise, [fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#link-type-preload:concept-event-fire-2
        x-internal="concept-event-fire"} named
        [`load`{#link-type-preload:event-load}](indices.html#event-load)
        at `el`{.variable}.

        The actual browsers\' behavior is different from the spec here,
        and the feasibility of changing the behavior has not yet been
        investigated. See [issue
        #1142](https://github.com/whatwg/html/issues/1142).

The [process a link
header](semantics.html#process-a-link-header){#link-type-preload:process-a-link-header}
step for this type of link given a [link processing
options](semantics.html#link-processing-options){#link-type-preload:link-processing-options-2}
`options`{.variable} is to
[preload](#preload){#link-type-preload:preload-2} `options`{.variable}.

##### [4.6.7.21]{.secno} Link type \"[`privacy-policy`]{.dfn dfn-for="link/rel,a/rel,area/rel" dfn-type="attr-value"}\"[](#link-type-privacy-policy){.self-link}

The
[`privacy-policy`{#link-type-privacy-policy:link-type-privacy-policy}](#link-type-privacy-policy)
keyword may be used with
[`link`{#link-type-privacy-policy:the-link-element}](semantics.html#the-link-element),
[`a`{#link-type-privacy-policy:the-a-element}](text-level-semantics.html#the-a-element),
and
[`area`{#link-type-privacy-policy:the-area-element}](image-maps.html#the-area-element)
elements. This keyword creates a
[hyperlink](#hyperlink){#link-type-privacy-policy:hyperlink}.

The
[`privacy-policy`{#link-type-privacy-policy:link-type-privacy-policy-2}](#link-type-privacy-policy)
keyword indicates that the referenced document contains information
about the data collection and usage practices that apply to the current
document, as described in more detail in Additional Link Relation Types.
The referenced document may be a standalone privacy policy, or a
specific section of some more general document.
[\[RFC6903\]](references.html#refsRFC6903)

##### [4.6.7.22]{.secno} Link type \"[`search`]{.dfn dfn-for="link/rel,a/rel,area/rel,form/rel" dfn-type="attr-value"}\"[](#link-type-search){.self-link}

The [`search`{#link-type-search:link-type-search}](#link-type-search)
keyword may be used with
[`link`{#link-type-search:the-link-element}](semantics.html#the-link-element),
[`a`{#link-type-search:the-a-element}](text-level-semantics.html#the-a-element),
[`area`{#link-type-search:the-area-element}](image-maps.html#the-area-element),
and
[`form`{#link-type-search:the-form-element}](forms.html#the-form-element)
elements. This keyword creates a
[hyperlink](#hyperlink){#link-type-search:hyperlink}.

The [`search`{#link-type-search:link-type-search-2}](#link-type-search)
keyword indicates that the referenced document provides an interface
specifically for searching the document and its related resources.

OpenSearch description documents can be used with
[`link`{#link-type-search:the-link-element-2}](semantics.html#the-link-element)
elements and the
[`search`{#link-type-search:link-type-search-3}](#link-type-search) link
type to enable user agents to autodiscover search interfaces.
[\[OPENSEARCH\]](references.html#refsOPENSEARCH)

##### [4.6.7.23]{.secno} Link type \"[`stylesheet`]{.dfn dfn-for="link/rel" dfn-type="attr-value"}\"[](#link-type-stylesheet){.self-link}

The
[`stylesheet`{#link-type-stylesheet:link-type-stylesheet}](#link-type-stylesheet)
keyword may be used with
[`link`{#link-type-stylesheet:the-link-element}](semantics.html#the-link-element)
elements. This keyword creates an [external resource
link](#external-resource-link){#link-type-stylesheet:external-resource-link}
that contributes to the styling processing model. This keyword is
[body-ok](#body-ok){#link-type-stylesheet:body-ok}.

The specified resource is a [CSS style
sheet](https://drafts.csswg.org/cssom/#css-style-sheet){#link-type-stylesheet:css-style-sheet
x-internal="css-style-sheet"} that describes how to present the
document.

::::: {.mdn-anno .wrapped .before}
**⚠**MDN

:::: feature
[Alternative_style_sheets](https://developer.mozilla.org/en-US/docs/Web/CSS/Alternative_style_sheets "Specifying alternative style sheets in a web page provides a way for users to see multiple versions of a page, based on their needs or preferences.")

Support in one engine only.

::: support
[Firefox3+]{.firefox .yes}[Safari?]{.safari
.unknown}[Chrome1--48]{.chrome .no}

------------------------------------------------------------------------

[OperaYes]{.opera .yes}[EdgeNo]{.edge_blink .no}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet Explorer8+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

If the
[`alternate`{#link-type-stylesheet:rel-alternate}](#rel-alternate)
keyword is also specified on the
[`link`{#link-type-stylesheet:the-link-element-2}](semantics.html#the-link-element)
element, then [the link is an alternative style
sheet]{#the-link-is-an-alternative-stylesheet .dfn}; in this case, the
[`title`{#link-type-stylesheet:attr-title}](dom.html#attr-title)
attribute must be specified on the
[`link`{#link-type-stylesheet:the-link-element-3}](semantics.html#the-link-element)
element, with a non-empty value.

The default type for resources given by the
[`stylesheet`{#link-type-stylesheet:link-type-stylesheet-2}](#link-type-stylesheet)
keyword is
[`text/css`{#link-type-stylesheet:text/css}](indices.html#text/css).

A
[`link`{#link-type-stylesheet:the-link-element-4}](semantics.html#the-link-element)
element of this type is [implicitly potentially
render-blocking](urls-and-fetching.html#implicitly-potentially-render-blocking){#link-type-stylesheet:implicitly-potentially-render-blocking}
if the element was created by its [node
document](https://dom.spec.whatwg.org/#concept-node-document){#link-type-stylesheet:node-document
x-internal="node-document"}\'s parser.

When the
[`disabled`{#link-type-stylesheet:attr-link-disabled}](semantics.html#attr-link-disabled)
attribute of a
[`link`{#link-type-stylesheet:the-link-element-5}](semantics.html#the-link-element)
element with a
[`stylesheet`{#link-type-stylesheet:link-type-stylesheet-3}](#link-type-stylesheet)
keyword is set,
[disable](https://drafts.csswg.org/cssom/#disable-a-css-style-sheet){#link-type-stylesheet:disable-a-css-style-sheet
x-internal="disable-a-css-style-sheet"} the [associated CSS style
sheet](https://drafts.csswg.org/cssom/#associated-css-style-sheet){#link-type-stylesheet:associated-css-style-sheet
x-internal="associated-css-style-sheet"}.

The appropriate times to [fetch and
process](semantics.html#fetch-and-process-the-linked-resource){#link-type-stylesheet:fetch-and-process-the-linked-resource}
this type of link are:

- When the [external resource
  link](#external-resource-link){#link-type-stylesheet:external-resource-link-2}
  is created on a
  [`link`{#link-type-stylesheet:the-link-element-6}](semantics.html#the-link-element)
  element that is already [browsing-context
  connected](infrastructure.html#browsing-context-connected){#link-type-stylesheet:browsing-context-connected}.

- When the [external resource
  link](#external-resource-link){#link-type-stylesheet:external-resource-link-3}\'s
  [`link`{#link-type-stylesheet:the-link-element-7}](semantics.html#the-link-element)
  element [becomes browsing-context
  connected](infrastructure.html#becomes-browsing-context-connected){#link-type-stylesheet:becomes-browsing-context-connected}.

- When the
  [`href`{#link-type-stylesheet:attr-link-href}](semantics.html#attr-link-href)
  attribute of the
  [`link`{#link-type-stylesheet:the-link-element-8}](semantics.html#the-link-element)
  element of an [external resource
  link](#external-resource-link){#link-type-stylesheet:external-resource-link-4}
  that is already [browsing-context
  connected](infrastructure.html#browsing-context-connected){#link-type-stylesheet:browsing-context-connected-2}
  is changed.

- When the
  [`disabled`{#link-type-stylesheet:attr-link-disabled-2}](semantics.html#attr-link-disabled)
  attribute of the
  [`link`{#link-type-stylesheet:the-link-element-9}](semantics.html#the-link-element)
  element of an [external resource
  link](#external-resource-link){#link-type-stylesheet:external-resource-link-5}
  that is already [browsing-context
  connected](infrastructure.html#browsing-context-connected){#link-type-stylesheet:browsing-context-connected-3}
  is set, changed, or removed.

- When the
  [`crossorigin`{#link-type-stylesheet:attr-link-crossorigin}](semantics.html#attr-link-crossorigin)
  attribute of the
  [`link`{#link-type-stylesheet:the-link-element-10}](semantics.html#the-link-element)
  element of an [external resource
  link](#external-resource-link){#link-type-stylesheet:external-resource-link-6}
  that is already [browsing-context
  connected](infrastructure.html#browsing-context-connected){#link-type-stylesheet:browsing-context-connected-4}
  is set, changed, or removed.

- When the
  [`type`{#link-type-stylesheet:attr-link-type}](semantics.html#attr-link-type)
  attribute of the
  [`link`{#link-type-stylesheet:the-link-element-11}](semantics.html#the-link-element)
  element of an [external resource
  link](#external-resource-link){#link-type-stylesheet:external-resource-link-7}
  that is already [browsing-context
  connected](infrastructure.html#browsing-context-connected){#link-type-stylesheet:browsing-context-connected-5}
  is set or changed to a value that does not or no longer matches the
  [Content-Type
  metadata](urls-and-fetching.html#content-type){#link-type-stylesheet:content-type}
  of the previous obtained external resource, if any.

- When the
  [`type`{#link-type-stylesheet:attr-link-type-2}](semantics.html#attr-link-type)
  attribute of the
  [`link`{#link-type-stylesheet:the-link-element-12}](semantics.html#the-link-element)
  element of an [external resource
  link](#external-resource-link){#link-type-stylesheet:external-resource-link-8}
  that is already [browsing-context
  connected](infrastructure.html#browsing-context-connected){#link-type-stylesheet:browsing-context-connected-6},
  but was previously not obtained due to the
  [`type`{#link-type-stylesheet:attr-link-type-3}](semantics.html#attr-link-type)
  attribute specifying an unsupported type, is removed or changed.

- When the [external resource
  link](#external-resource-link){#link-type-stylesheet:external-resource-link-9}
  that is already [browsing-context
  connected](infrastructure.html#browsing-context-connected){#link-type-stylesheet:browsing-context-connected-7}
  changes from being [an alternative style
  sheet](#the-link-is-an-alternative-stylesheet){#link-type-stylesheet:the-link-is-an-alternative-stylesheet}
  to not being one, or vice versa.

**Quirk**: If the document has been set to [quirks
mode](https://dom.spec.whatwg.org/#concept-document-quirks){#link-type-stylesheet:quirks-mode
x-internal="quirks-mode"}, has the [same
origin](browsers.html#same-origin){#link-type-stylesheet:same-origin} as
the
[URL](https://url.spec.whatwg.org/#concept-url){#link-type-stylesheet:url
x-internal="url"} of the external resource, and the [Content-Type
metadata](urls-and-fetching.html#content-type){#link-type-stylesheet:content-type-2}
of the external resource is not a supported style sheet type, the user
agent must instead assume it to be
[`text/css`{#link-type-stylesheet:text/css-2}](indices.html#text/css).

The [linked resource fetch setup
steps](semantics.html#linked-resource-fetch-setup-steps){#link-type-stylesheet:linked-resource-fetch-setup-steps}
for this type of linked resource, given a
[`link`{#link-type-stylesheet:the-link-element-13}](semantics.html#the-link-element)
element `el`{.variable} and
[request](https://fetch.spec.whatwg.org/#concept-request){#link-type-stylesheet:concept-request
x-internal="concept-request"} `request`{.variable}, are:

1.  If `el`{.variable}\'s
    [`disabled`{#link-type-stylesheet:attr-link-disabled-3}](semantics.html#attr-link-disabled)
    attribute is set, then return false.

2.  If `el`{.variable} [contributes a script-blocking style
    sheet](semantics.html#contributes-a-script-blocking-style-sheet){#link-type-stylesheet:contributes-a-script-blocking-style-sheet},
    [append](https://infra.spec.whatwg.org/#set-append){#link-type-stylesheet:set-append
    x-internal="set-append"} `el`{.variable} to its [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#link-type-stylesheet:node-document-2
    x-internal="node-document"}\'s [script-blocking style sheet
    set](semantics.html#script-blocking-style-sheet-set){#link-type-stylesheet:script-blocking-style-sheet-set}.

3.  If `el`{.variable}\'s
    [`media`{#link-type-stylesheet:attr-link-media}](semantics.html#attr-link-media)
    attribute\'s value [matches the
    environment](common-microsyntaxes.html#matches-the-environment){#link-type-stylesheet:matches-the-environment}
    and `el`{.variable} is [potentially
    render-blocking](urls-and-fetching.html#potentially-render-blocking){#link-type-stylesheet:potentially-render-blocking},
    then [block
    rendering](dom.html#block-rendering){#link-type-stylesheet:block-rendering}
    on `el`{.variable}.

4.  If `el`{.variable} is currently
    [render-blocking](dom.html#render-blocking){#link-type-stylesheet:render-blocking},
    then set `request`{.variable}\'s
    [render-blocking](https://fetch.spec.whatwg.org/#request-render-blocking){#link-type-stylesheet:concept-request-render-blocking
    x-internal="concept-request-render-blocking"} to true.

5.  Return true.

See [issue #968](https://github.com/whatwg/html/issues/968) for plans to
use the CSSOM [fetch a CSS style
sheet](https://drafts.csswg.org/cssom/#fetching-css-style-sheets)
algorithm instead of the [default fetch and process the linked
resource](semantics.html#default-fetch-and-process-the-linked-resource){#link-type-stylesheet:default-fetch-and-process-the-linked-resource}
algorithm. In the meantime, any [critical
subresource](infrastructure.html#critical-subresources){#link-type-stylesheet:critical-subresources}
[request](https://fetch.spec.whatwg.org/#concept-request){#link-type-stylesheet:concept-request-2
x-internal="concept-request"} should have its
[render-blocking](https://fetch.spec.whatwg.org/#request-render-blocking){#link-type-stylesheet:concept-request-render-blocking-2
x-internal="concept-request-render-blocking"} set to whether or not the
[`link`{#link-type-stylesheet:the-link-element-14}](semantics.html#the-link-element)
element is currently
[render-blocking](dom.html#render-blocking){#link-type-stylesheet:render-blocking-2}.

To [process this type of linked
resource](semantics.html#process-the-linked-resource){#link-type-stylesheet:process-the-linked-resource}
given a
[`link`{#link-type-stylesheet:the-link-element-15}](semantics.html#the-link-element)
element `el`{.variable}, boolean `success`{.variable},
[response](https://fetch.spec.whatwg.org/#concept-response){#link-type-stylesheet:concept-response
x-internal="concept-response"} `response`{.variable}, and [byte
sequence](https://infra.spec.whatwg.org/#byte-sequence){#link-type-stylesheet:byte-sequence
x-internal="byte-sequence"} `bodyBytes`{.variable}:

1.  If the resource\'s [Content-Type
    metadata](urls-and-fetching.html#content-type){#link-type-stylesheet:content-type-3}
    is not
    [`text/css`{#link-type-stylesheet:text/css-3}](indices.html#text/css),
    then set `success`{.variable} to false.

2.  If `el`{.variable} no longer creates an [external resource
    link](#external-resource-link){#link-type-stylesheet:external-resource-link-10}
    that contributes to the styling processing model, or if, since the
    resource in question was
    [fetched](semantics.html#fetch-and-process-the-linked-resource){#link-type-stylesheet:fetch-and-process-the-linked-resource-2},
    it has become appropriate to
    [fetch](semantics.html#fetch-and-process-the-linked-resource){#link-type-stylesheet:fetch-and-process-the-linked-resource-3}
    it again, then:

    1.  [Remove](https://infra.spec.whatwg.org/#list-remove){#link-type-stylesheet:list-remove
        x-internal="list-remove"} `el`{.variable} from
        `el`{.variable}\'s [node
        document](https://dom.spec.whatwg.org/#concept-node-document){#link-type-stylesheet:node-document-3
        x-internal="node-document"}\'s [script-blocking style sheet
        set](semantics.html#script-blocking-style-sheet-set){#link-type-stylesheet:script-blocking-style-sheet-set-2}.

    2.  Return.

3.  If `el`{.variable} has an [associated CSS style
    sheet](https://drafts.csswg.org/cssom/#associated-css-style-sheet){#link-type-stylesheet:associated-css-style-sheet-2
    x-internal="associated-css-style-sheet"}, [remove the CSS style
    sheet](https://drafts.csswg.org/cssom/#remove-a-css-style-sheet){#link-type-stylesheet:remove-a-css-style-sheet
    x-internal="remove-a-css-style-sheet"}.

4.  If `success`{.variable} is true, then:

    1.  [Create a CSS style
        sheet](https://drafts.csswg.org/cssom/#create-a-css-style-sheet){#link-type-stylesheet:create-a-css-style-sheet
        x-internal="create-a-css-style-sheet"} with the following
        properties:

        [type](https://drafts.csswg.org/cssom/#concept-css-style-sheet-type){#link-type-stylesheet:concept-css-style-sheet-type x-internal="concept-css-style-sheet-type"}
        :   [`text/css`{#link-type-stylesheet:text/css-4}](indices.html#text/css)

        [location](https://drafts.csswg.org/cssom/#concept-css-style-sheet-location){#link-type-stylesheet:concept-css-style-sheet-location x-internal="concept-css-style-sheet-location"}

        :   `response`{.variable}\'s [URL
            list](https://fetch.spec.whatwg.org/#concept-response-url-list){#link-type-stylesheet:concept-response-url-list
            x-internal="concept-response-url-list"}\[0\]

            We provide a URL here on the assumption that
            [w3c/csswg-drafts issue
            #9316](https://github.com/w3c/csswg-drafts/issues/9316) will
            be fixed.

        [owner node](https://drafts.csswg.org/cssom/#concept-css-style-sheet-owner-node){#link-type-stylesheet:concept-css-style-sheet-owner-node x-internal="concept-css-style-sheet-owner-node"}
        :   `el`{.variable}

        [media](https://drafts.csswg.org/cssom/#concept-css-style-sheet-media){#link-type-stylesheet:concept-css-style-sheet-media x-internal="concept-css-style-sheet-media"}

        :   The
            [`media`{#link-type-stylesheet:attr-link-media-2}](semantics.html#attr-link-media)
            attribute of `el`{.variable}.

            This is a reference to the (possibly absent at this time)
            attribute, rather than a copy of the attribute\'s current
            value. CSSOM defines what happens when the attribute is
            dynamically set, changed, or removed.

        [title](https://drafts.csswg.org/cssom/#concept-css-style-sheet-title){#link-type-stylesheet:concept-css-style-sheet-title x-internal="concept-css-style-sheet-title"}

        :   The
            [`title`{#link-type-stylesheet:attr-link-title}](semantics.html#attr-link-title)
            attribute of `el`{.variable}, if `el`{.variable} is [in a
            document
            tree](https://dom.spec.whatwg.org/#in-a-document-tree){#link-type-stylesheet:in-a-document-tree
            x-internal="in-a-document-tree"}, or the empty string
            otherwise.

            This is similarly a reference to the attribute, rather than
            a copy of the attribute\'s current value.

        [alternate flag](https://drafts.csswg.org/cssom/#concept-css-style-sheet-alternate-flag){#link-type-stylesheet:concept-css-style-sheet-alternate-flag x-internal="concept-css-style-sheet-alternate-flag"}
        :   Set if [the link is an alternative style
            sheet](#the-link-is-an-alternative-stylesheet){#link-type-stylesheet:the-link-is-an-alternative-stylesheet-2}
            and `el`{.variable}\'s [explicitly
            enabled](semantics.html#explicitly-enabled){#link-type-stylesheet:explicitly-enabled}
            is false; unset otherwise.

        [origin-clean flag](https://drafts.csswg.org/cssom/#concept-css-style-sheet-origin-clean-flag){#link-type-stylesheet:concept-css-style-sheet-origin-clean-flag x-internal="concept-css-style-sheet-origin-clean-flag"}
        :   Set if the resource is
            [CORS-same-origin](urls-and-fetching.html#cors-same-origin){#link-type-stylesheet:cors-same-origin};
            unset otherwise.

        [parent CSS style sheet](https://drafts.csswg.org/cssom/#concept-css-style-sheet-parent-css-style-sheet){#link-type-stylesheet:concept-css-style-sheet-parent-css-style-sheet x-internal="concept-css-style-sheet-parent-css-style-sheet"}\
        [owner CSS rule](https://drafts.csswg.org/cssom/#concept-css-style-sheet-owner-css-rule){#link-type-stylesheet:concept-css-style-sheet-owner-css-rule x-internal="concept-css-style-sheet-owner-css-rule"}
        :   null

        [disabled flag](https://drafts.csswg.org/cssom/#concept-css-style-sheet-disabled-flag){#link-type-stylesheet:concept-css-style-sheet-disabled-flag x-internal="concept-css-style-sheet-disabled-flag"}
        :   Left at its default value.

        [CSS rules](https://drafts.csswg.org/cssom/#concept-css-style-sheet-css-rules){#link-type-stylesheet:concept-css-style-sheet-css-rules x-internal="concept-css-style-sheet-css-rules"}

        :   Left uninitialized.

            This doesn\'t seem right. Presumably we should be using
            `bodyBytes`{.variable}? Tracked as [issue
            #2997](https://github.com/whatwg/html/issues/2997).

        The CSS [environment
        encoding](https://drafts.csswg.org/css-syntax/#environment-encoding){#link-type-stylesheet:environment-encoding
        x-internal="environment-encoding"} is the result of running the
        following steps: [\[CSSSYNTAX\]](references.html#refsCSSSYNTAX)

        1.  If `el`{.variable} has a
            [`charset`{#link-type-stylesheet:attr-link-charset}](obsolete.html#attr-link-charset)
            attribute, [get an
            encoding](https://encoding.spec.whatwg.org/#concept-encoding-get){#link-type-stylesheet:getting-an-encoding
            x-internal="getting-an-encoding"} from that attribute\'s
            value. If that succeeds, return the resulting encoding.
            [\[ENCODING\]](references.html#refsENCODING)

        2.  Otherwise, return the [document\'s character
            encoding](https://dom.spec.whatwg.org/#concept-document-encoding){#link-type-stylesheet:document's-character-encoding
            x-internal="document's-character-encoding"}.
            [\[DOM\]](references.html#refsDOM)

    2.  [Fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#link-type-stylesheet:concept-event-fire
        x-internal="concept-event-fire"} named
        [`load`{#link-type-stylesheet:event-load}](indices.html#event-load)
        at `el`{.variable}.

5.  Otherwise, [fire an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#link-type-stylesheet:concept-event-fire-2
    x-internal="concept-event-fire"} named
    [`error`{#link-type-stylesheet:event-error}](indices.html#event-error)
    at `el`{.variable}.

6.  If `el`{.variable} [contributes a script-blocking style
    sheet](semantics.html#contributes-a-script-blocking-style-sheet){#link-type-stylesheet:contributes-a-script-blocking-style-sheet-2},
    then:

    1.  [Assert](https://infra.spec.whatwg.org/#assert){#link-type-stylesheet:assert
        x-internal="assert"}: `el`{.variable}\'s [node
        document](https://dom.spec.whatwg.org/#concept-node-document){#link-type-stylesheet:node-document-4
        x-internal="node-document"}\'s [script-blocking style sheet
        set](semantics.html#script-blocking-style-sheet-set){#link-type-stylesheet:script-blocking-style-sheet-set-3}
        [contains](https://infra.spec.whatwg.org/#list-contain){#link-type-stylesheet:list-contains
        x-internal="list-contains"} `el`{.variable}.

    2.  [Remove](https://infra.spec.whatwg.org/#list-remove){#link-type-stylesheet:list-remove-2
        x-internal="list-remove"} `el`{.variable} from its [node
        document](https://dom.spec.whatwg.org/#concept-node-document){#link-type-stylesheet:node-document-5
        x-internal="node-document"}\'s [script-blocking style sheet
        set](semantics.html#script-blocking-style-sheet-set){#link-type-stylesheet:script-blocking-style-sheet-set-4}.

7.  [Unblock
    rendering](dom.html#unblock-rendering){#link-type-stylesheet:unblock-rendering}
    on `el`{.variable}.

The [process a link
header](semantics.html#process-a-link-header){#link-type-stylesheet:process-a-link-header}
steps for this type of linked resource are to do nothing.

##### [4.6.7.24]{.secno} Link type \"[`tag`]{.dfn dfn-for="a/rel,area/rel" dfn-type="attr-value"}\"[](#link-type-tag){.self-link}

The [`tag`{#link-type-tag:link-type-tag}](#link-type-tag) keyword may be
used with
[`a`{#link-type-tag:the-a-element}](text-level-semantics.html#the-a-element)
and
[`area`{#link-type-tag:the-area-element}](image-maps.html#the-area-element)
elements. This keyword creates a
[hyperlink](#hyperlink){#link-type-tag:hyperlink}.

The [`tag`{#link-type-tag:link-type-tag-2}](#link-type-tag) keyword
indicates that the *tag* that the referenced document represents applies
to the current document.

Since it indicates that the tag *applies to the current document*, it
would be inappropriate to use this keyword in the markup of a [tag
cloud](semantics-other.html#tag-cloud), which lists the popular tags
across a set of pages.

::: example
This document is about some gems, and so it is *tagged* with
\"`https://en.wikipedia.org/wiki/Gemstone`\" to unambiguously categorize
it as applying to the \"jewel\" kind of gems, and not to, say, the towns
in the US, the Ruby package format, or the Swiss locomotive class:

``` html
<!DOCTYPE HTML>
<html lang="en">
 <head>
  <title>My Precious</title>
 </head>
 <body>
  <header><h1>My precious</h1> <p>Summer 2012</p></header>
  <p>Recently I managed to dispose of a red gem that had been
  bothering me. I now have a much nicer blue sapphire.</p>
  <p>The red gem had been found in a bauxite stone while I was digging
  out the office level, but nobody was willing to haul it away. The
  same red gem stayed there for literally years.</p>
  <footer>
   Tags: <a rel=tag href="https://en.wikipedia.org/wiki/Gemstone">Gemstone</a>
  </footer>
 </body>
</html>
```
:::

::: example
In *this* document, there are two articles. The
\"[`tag`{#link-type-tag:link-type-tag-3}](#link-type-tag)\" link,
however, applies to the whole page (and would do so wherever it was
placed, including if it was within the
[`article`{#link-type-tag:the-article-element}](sections.html#the-article-element)
elements).

``` html
<!DOCTYPE HTML>
<html lang="en">
 <head>
  <title>Gem 4/4</title>
 </head>
 <body>
  <article>
   <h1>801: Steinbock</h1>
   <p>The number 801 Gem 4/4 electro-diesel has an ibex and was rebuilt in 2002.</p>
  </article>
  <article>
   <h1>802: Murmeltier</h1>
   <figure>
    <img src="https://upload.wikimedia.org/wikipedia/commons/b/b0/Trains_de_la_Bernina_en_hiver_2.jpg"
         alt="The 802 was red with pantographs and tall vents on the side.">
    <figcaption>The 802 in the 1980s, above Lago Bianco.</figcaption>
   </figure>
   <p>The number 802 Gem 4/4 electro-diesel has a marmot and was rebuilt in 2003.</p>
  </article>
  <p class="topic"><a rel=tag href="https://en.wikipedia.org/wiki/Rhaetian_Railway_Gem_4/4">Gem 4/4</a></p>
 </body>
</html>
```
:::

##### [4.6.7.25]{.secno} Link Type \"[`terms-of-service`]{.dfn dfn-for="link/rel,a/rel,area/rel" dfn-type="attr-value"}\"[](#link-type-terms-of-service){.self-link}

The
[`terms-of-service`{#link-type-terms-of-service:link-type-terms-of-service}](#link-type-terms-of-service)
keyword may be used with
[`link`{#link-type-terms-of-service:the-link-element}](semantics.html#the-link-element),
[`a`{#link-type-terms-of-service:the-a-element}](text-level-semantics.html#the-a-element),
and
[`area`{#link-type-terms-of-service:the-area-element}](image-maps.html#the-area-element)
elements. This keyword creates a
[hyperlink](#hyperlink){#link-type-terms-of-service:hyperlink}.

The
[`terms-of-service`{#link-type-terms-of-service:link-type-terms-of-service-2}](#link-type-terms-of-service)
keyword indicates that the referenced document contains information
about the agreements between the current document\'s provider and users
who wish to use the current document, as described in more detail in
Additional Link Relation Types.
[\[RFC6903\]](references.html#refsRFC6903)

##### [4.6.7.26]{.secno} Sequential link types[](#sequential-link-types){.self-link}

Some documents form part of a sequence of documents.

A sequence of documents is one where each document can have a *previous
sibling* and a *next sibling*. A document with no previous sibling is
the start of its sequence, a document with no next sibling is the end of
its sequence.

A document may be part of multiple sequences.

###### [4.6.7.26.1]{.secno} Link type \"[`next`]{.dfn dfn-for="link/rel,a/rel,area/rel,form/rel" dfn-type="attr-value"}\"[](#link-type-next){.self-link}

The [`next`{#link-type-next:link-type-next}](#link-type-next) keyword
may be used with
[`link`{#link-type-next:the-link-element}](semantics.html#the-link-element),
[`a`{#link-type-next:the-a-element}](text-level-semantics.html#the-a-element),
[`area`{#link-type-next:the-area-element}](image-maps.html#the-area-element),
and
[`form`{#link-type-next:the-form-element}](forms.html#the-form-element)
elements. This keyword creates a
[hyperlink](#hyperlink){#link-type-next:hyperlink}.

The [`next`{#link-type-next:link-type-next-2}](#link-type-next) keyword
indicates that the document is part of a sequence, and that the link is
leading to the document that is the next logical document in the
sequence.

When the [`next`{#link-type-next:link-type-next-3}](#link-type-next)
keyword is used with a
[`link`{#link-type-next:the-link-element-2}](semantics.html#the-link-element)
element, user agents should process such links as if they were using one
of the
[`dns-prefetch`{#link-type-next:link-type-dns-prefetch}](#link-type-dns-prefetch),
[`preconnect`{#link-type-next:link-type-preconnect}](#link-type-preconnect),
or [`prefetch`{#link-type-next:link-type-prefetch}](#link-type-prefetch)
keywords. Which keyword the user agent wishes to use is
implementation-dependent; for example, a user agent may wish to use the
less-costly
[`preconnect`{#link-type-next:link-type-preconnect-2}](#link-type-preconnect)
processing model when trying to conserve data, battery power, or
processing power, or may wish to pick a keyword depending on heuristic
analysis of past user behavior in similar scenarios.

###### [4.6.7.26.2]{.secno} Link type \"[`prev`]{.dfn dfn-for="link/rel,a/rel,area/rel,form/rel" dfn-type="attr-value"}\"[](#link-type-prev){.self-link}

The [`prev`{#link-type-prev:link-type-prev}](#link-type-prev) keyword
may be used with
[`link`{#link-type-prev:the-link-element}](semantics.html#the-link-element),
[`a`{#link-type-prev:the-a-element}](text-level-semantics.html#the-a-element),
[`area`{#link-type-prev:the-area-element}](image-maps.html#the-area-element),
and
[`form`{#link-type-prev:the-form-element}](forms.html#the-form-element)
elements. This keyword creates a
[hyperlink](#hyperlink){#link-type-prev:hyperlink}.

The [`prev`{#link-type-prev:link-type-prev-2}](#link-type-prev) keyword
indicates that the document is part of a sequence, and that the link is
leading to the document that is the previous logical document in the
sequence.

**Synonyms**: For historical reasons, user agents must also treat the
keyword \"`previous`\" like the
[`prev`{#link-type-prev:link-type-prev-3}](#link-type-prev) keyword.

##### [4.6.7.27]{.secno} Other link types[](#other-link-types){.self-link}

[Extensions to the predefined set of link types]{#concept-rel-extensions
.dfn} may be registered on the [microformats page for existing rel
values](https://microformats.org/wiki/existing-rel-values#HTML5_link_type_extensions).
[\[MFREL\]](references.html#refsMFREL)

Anyone is free to edit the microformats page for existing rel values at
any time to add a type. Extension types must be specified with the
following information:

Keyword

:   The actual value being defined. The value should not be confusingly
    similar to any other defined value (e.g. differing only in case).

    If the value contains a U+003A COLON character (:), it must also be
    an [absolute
    URL](https://url.spec.whatwg.org/#syntax-url-absolute){#other-link-types:absolute-url
    x-internal="absolute-url"}.

Effect on\... [`link`{#other-link-types:the-link-element}](semantics.html#the-link-element)

:   One of the following:

    Not allowed
    :   The keyword must not be specified on
        [`link`{#other-link-types:the-link-element-2}](semantics.html#the-link-element)
        elements.

    Hyperlink
    :   The keyword may be specified on a
        [`link`{#other-link-types:the-link-element-3}](semantics.html#the-link-element)
        element; it creates a
        [hyperlink](#hyperlink){#other-link-types:hyperlink}.

    External Resource
    :   The keyword may be specified on a
        [`link`{#other-link-types:the-link-element-4}](semantics.html#the-link-element)
        element; it creates an [external resource
        link](#external-resource-link){#other-link-types:external-resource-link}.

Effect on\... [`a`{#other-link-types:the-a-element}](text-level-semantics.html#the-a-element) and [`area`{#other-link-types:the-area-element}](image-maps.html#the-area-element)

:   One of the following:

    Not allowed
    :   The keyword must not be specified on
        [`a`{#other-link-types:the-a-element-2}](text-level-semantics.html#the-a-element)
        and
        [`area`{#other-link-types:the-area-element-2}](image-maps.html#the-area-element)
        elements.

    Hyperlink
    :   The keyword may be specified on
        [`a`{#other-link-types:the-a-element-3}](text-level-semantics.html#the-a-element)
        and
        [`area`{#other-link-types:the-area-element-3}](image-maps.html#the-area-element)
        elements; it creates a
        [hyperlink](#hyperlink){#other-link-types:hyperlink-2}.

    External Resource
    :   The keyword may be specified on
        [`a`{#other-link-types:the-a-element-4}](text-level-semantics.html#the-a-element)
        and
        [`area`{#other-link-types:the-area-element-4}](image-maps.html#the-area-element)
        elements; it creates an [external resource
        link](#external-resource-link){#other-link-types:external-resource-link-2}.

    Hyperlink Annotation
    :   The keyword may be specified on
        [`a`{#other-link-types:the-a-element-5}](text-level-semantics.html#the-a-element)
        and
        [`area`{#other-link-types:the-area-element-5}](image-maps.html#the-area-element)
        elements; it
        [annotates](#hyperlink-annotation){#other-link-types:hyperlink-annotation}
        other [hyperlinks](#hyperlink){#other-link-types:hyperlink-3}
        created by the element.

Effect on\... [`form`{#other-link-types:the-form-element}](forms.html#the-form-element)

:   One of the following:

    Not allowed
    :   The keyword must not be specified on
        [`form`{#other-link-types:the-form-element-2}](forms.html#the-form-element)
        elements.

    Hyperlink
    :   The keyword may be specified on
        [`form`{#other-link-types:the-form-element-3}](forms.html#the-form-element)
        elements; it creates a
        [hyperlink](#hyperlink){#other-link-types:hyperlink-4}.

    External Resource
    :   The keyword may be specified on
        [`form`{#other-link-types:the-form-element-4}](forms.html#the-form-element)
        elements; it creates an [external resource
        link](#external-resource-link){#other-link-types:external-resource-link-3}.

    Hyperlink Annotation
    :   The keyword may be specified on
        [`form`{#other-link-types:the-form-element-5}](forms.html#the-form-element)
        elements; it
        [annotates](#hyperlink-annotation){#other-link-types:hyperlink-annotation-2}
        other [hyperlinks](#hyperlink){#other-link-types:hyperlink-5}
        created by the element.

Brief description
:   A short non-normative description of what the keyword\'s meaning is.

Specification
:   A link to a more detailed description of the keyword\'s semantics
    and requirements. It could be another page on the wiki, or a link to
    an external page.

Synonyms
:   A list of other keyword values that have exactly the same processing
    requirements. Authors should not use the values defined to be
    synonyms, they are only intended to allow user agents to support
    legacy content. Anyone may remove synonyms that are not used in
    practice; only names that need to be processed as synonyms for
    compatibility with legacy content are to be registered in this way.

Status

:   One of the following:

    Proposed
    :   The keyword has not received wide peer review and approval.
        Someone has proposed it and is, or soon will be, using it.

    Ratified
    :   The keyword has received wide peer review and approval. It has a
        specification that unambiguously defines how to handle pages
        that use the keyword, including when they use it in incorrect
        ways.

    Discontinued
    :   The keyword has received wide peer review and it has been found
        wanting. Existing pages are using this keyword, but new pages
        should avoid it. The \"brief description\" and \"specification\"
        entries will give details of what authors should use instead, if
        anything.

    If a keyword is found to be redundant with existing values, it
    should be removed and listed as a synonym for the existing value.

    If a keyword is registered in the \"proposed\" state for a period of
    a month or more without being used or specified, then it may be
    removed from the registry.

    If a keyword is added with the \"proposed\" status and found to be
    redundant with existing values, it should be removed and listed as a
    synonym for the existing value. If a keyword is added with the
    \"proposed\" status and found to be harmful, then it should be
    changed to \"discontinued\" status.

    Anyone can change the status at any time, but should only do so in
    accordance with the definitions above.

Conformance checkers must use the information given on the microformats
page for existing rel values to establish if a value is allowed or not:
values defined in this specification or marked as \"proposed\" or
\"ratified\" must be accepted when used on the elements for which they
apply as described in the \"Effect on\...\" field, whereas values marked
as \"discontinued\" or not listed in either this specification or on the
aforementioned page must be rejected as invalid. Conformance checkers
may cache this information (e.g. for performance reasons or to avoid the
use of unreliable network connectivity).

When an author uses a new type not defined by either this specification
or the wiki page, conformance checkers should offer to add the value to
the wiki, with the details described above, with the \"proposed\"
status.

Types defined as extensions in the [microformats page for existing rel
values](https://microformats.org/wiki/existing-rel-values#HTML5_link_type_extensions)
with the status \"proposed\" or \"ratified\" may be used with the `rel`
attribute on
[`link`{#other-link-types:the-link-element-5}](semantics.html#the-link-element),
[`a`{#other-link-types:the-a-element-6}](text-level-semantics.html#the-a-element),
and
[`area`{#other-link-types:the-area-element-6}](image-maps.html#the-area-element)
elements in accordance to the \"Effect on\...\" field.
[\[MFREL\]](references.html#refsMFREL)

[← 4.5 Text-level semantics](text-level-semantics.html) --- [Table of
Contents](./) --- [4.7 Edits →](edits.html)
