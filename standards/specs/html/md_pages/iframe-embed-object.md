HTML Standard

::: {#head .head .with-buttons}
[![WHATWG](https://resources.whatwg.org/logo.svg){.darkmode-aware
width="100" crossorigin="" height="100"}](https://whatwg.org/){.logo}

# HTML {#html .allcaps}

Living Standard --- Last Updated [2 June 2025]{.pubdate}
:::

[← 4.8.4 Images](images.html) --- [Table of Contents](./) --- [4.8.8 The
video element →](media.html)

1.  ::: {#toc-semantics}
    1.  1.  [[4.8.5]{.secno} The `iframe`
            element](iframe-embed-object.html#the-iframe-element)
        2.  [[4.8.6]{.secno} The `embed`
            element](iframe-embed-object.html#the-embed-element)
        3.  [[4.8.7]{.secno} The `object`
            element](iframe-embed-object.html#the-object-element)
    :::

#### [4.8.5]{.secno} The [`iframe`]{.dfn dfn-type="element"} element[](#the-iframe-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/iframe](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/iframe "The <iframe> HTML element represents a nested browsing context, embedding another HTML page into the current one.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera15+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerYes]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android14+]{.opera_android .yes}
:::
::::
:::::

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLIFrameElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLIFrameElement "The HTMLIFrameElement interface provides special properties and methods (beyond those of the HTMLElement interface it also has available to it by inheritance) for manipulating the layout and presentation of inline frame elements.")

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
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

[Categories](dom.html#concept-element-categories){#the-iframe-element:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-iframe-element:flow-content-2}.
:   [Phrasing
    content](dom.html#phrasing-content-2){#the-iframe-element:phrasing-content-2}.
:   [Embedded
    content](dom.html#embedded-content-category){#the-iframe-element:embedded-content-category}.
:   [Interactive
    content](dom.html#interactive-content-2){#the-iframe-element:interactive-content-2}.
:   [Palpable
    content](dom.html#palpable-content-2){#the-iframe-element:palpable-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-iframe-element:concept-element-contexts}:
:   Where [embedded
    content](dom.html#embedded-content-category){#the-iframe-element:embedded-content-category-2}
    is expected.

[Content model](dom.html#concept-element-content-model){#the-iframe-element:concept-element-content-model}:
:   [Nothing](dom.html#concept-content-nothing){#the-iframe-element:concept-content-nothing}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-iframe-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-iframe-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-iframe-element:global-attributes}
:   [`src`{#the-iframe-element:attr-iframe-src}](#attr-iframe-src) ---
    Address of the resource
:   [`srcdoc`{#the-iframe-element:attr-iframe-srcdoc}](#attr-iframe-srcdoc)
    --- A document to render in the
    [`iframe`{#the-iframe-element:the-iframe-element}](#the-iframe-element)
:   [`name`{#the-iframe-element:attr-iframe-name}](#attr-iframe-name)
    --- Name of [content
    navigable](document-sequences.html#content-navigable){#the-iframe-element:content-navigable}
:   [`sandbox`{#the-iframe-element:attr-iframe-sandbox}](#attr-iframe-sandbox)
    --- Security rules for nested content
:   [`allow`{#the-iframe-element:attr-iframe-allow}](#attr-iframe-allow)
    --- [Permissions
    policy](https://w3c.github.io/webappsec-feature-policy/#permissions-policy){#the-iframe-element:concept-permissions-policy
    x-internal="concept-permissions-policy"} to be applied to the
    [`iframe`{#the-iframe-element:the-iframe-element-2}](#the-iframe-element)\'s
    contents
:   [`allowfullscreen`{#the-iframe-element:attr-iframe-allowfullscreen}](#attr-iframe-allowfullscreen)
    --- Whether to allow the
    [`iframe`{#the-iframe-element:the-iframe-element-3}](#the-iframe-element)\'s
    contents to use
    [`requestFullscreen()`{#the-iframe-element:dom-element-requestfullscreen}](https://fullscreen.spec.whatwg.org/#dom-element-requestfullscreen){x-internal="dom-element-requestfullscreen"}
:   [`width`{#the-iframe-element:attr-dim-width}](embedded-content-other.html#attr-dim-width)
    --- Horizontal dimension
:   [`height`{#the-iframe-element:attr-dim-height}](embedded-content-other.html#attr-dim-height)
    --- Vertical dimension
:   [`referrerpolicy`{#the-iframe-element:attr-iframe-referrerpolicy}](#attr-iframe-referrerpolicy)
    --- [Referrer
    policy](https://w3c.github.io/webappsec-referrer-policy/#referrer-policy){#the-iframe-element:referrer-policy
    x-internal="referrer-policy"} for
    [fetches](https://fetch.spec.whatwg.org/#concept-fetch){#the-iframe-element:concept-fetch
    x-internal="concept-fetch"} initiated by the element
:   [`loading`{#the-iframe-element:attr-iframe-loading}](#attr-iframe-loading)
    --- Used when determining loading deferral

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-iframe-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-iframe).
:   [For implementers](https://w3c.github.io/html-aam/#el-iframe).

[DOM interface](dom.html#concept-element-dom){#the-iframe-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLIFrameElement : HTMLElement {
      [HTMLConstructor] constructor();

      [CEReactions] attribute USVString src;
      [CEReactions] attribute (TrustedHTML or DOMString) srcdoc;
      [CEReactions] attribute DOMString name;
      [SameObject, PutForwards=value] readonly attribute DOMTokenList sandbox;
      [CEReactions] attribute DOMString allow;
      [CEReactions] attribute boolean allowFullscreen;
      [CEReactions] attribute DOMString width;
      [CEReactions] attribute DOMString height;
      [CEReactions] attribute DOMString referrerPolicy;
      [CEReactions] attribute DOMString loading;
      readonly attribute Document? contentDocument;
      readonly attribute WindowProxy? contentWindow;
      Document? getSVGDocument();

      // also has obsolete members
    };
    ```

The
[`iframe`{#the-iframe-element:the-iframe-element-4}](#the-iframe-element)
element
[represents](dom.html#represents){#the-iframe-element:represents} its
[content
navigable](document-sequences.html#content-navigable){#the-iframe-element:content-navigable-2}.

The [`src`]{#attr-iframe-src .dfn dfn-for="iframe"
dfn-type="element-attr"} attribute gives the
[URL](https://url.spec.whatwg.org/#concept-url){#the-iframe-element:url
x-internal="url"} of a page that the element\'s [content
navigable](document-sequences.html#content-navigable){#the-iframe-element:content-navigable-3}
is to contain. The attribute, if present, must be a [valid non-empty URL
potentially surrounded by
spaces](urls-and-fetching.html#valid-non-empty-url-potentially-surrounded-by-spaces){#the-iframe-element:valid-non-empty-url-potentially-surrounded-by-spaces}.
If the
[`itemprop`{#the-iframe-element:names:-the-itemprop-attribute}](microdata.html#names:-the-itemprop-attribute)
attribute is specified on an
[`iframe`{#the-iframe-element:the-iframe-element-5}](#the-iframe-element)
element, then the
[`src`{#the-iframe-element:attr-iframe-src-2}](#attr-iframe-src)
attribute must also be specified.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Element/iframe#attr-srcdoc](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/iframe#attr-srcdoc "The <iframe> HTML element represents a nested browsing context, embedding another HTML page into the current one.")

Support in all current engines.

::: support
[Firefox25+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome20+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

The [`srcdoc`]{#attr-iframe-srcdoc .dfn dfn-for="iframe"
dfn-type="element-attr"} attribute gives the content of the page that
the element\'s [content
navigable](document-sequences.html#content-navigable){#the-iframe-element:content-navigable-4}
is to contain. The value of the attribute is used to
[construct](browsing-the-web.html#create-navigation-params-from-a-srcdoc-resource){#the-iframe-element:create-navigation-params-from-a-srcdoc-resource}
[an `iframe` `srcdoc` document]{#an-iframe-srcdoc-document .dfn
export=""}, which is a
[`Document`{#the-iframe-element:document-3}](dom.html#document) whose
[URL](https://dom.spec.whatwg.org/#concept-document-url){#the-iframe-element:the-document's-address
x-internal="the-document's-address"} [matches
`about:srcdoc`](urls-and-fetching.html#matches-about:srcdoc){#the-iframe-element:matches-about:srcdoc}.

The
[`srcdoc`{#the-iframe-element:attr-iframe-srcdoc-2}](#attr-iframe-srcdoc)
attribute, if present, must have a value using [the HTML
syntax](syntax.html#syntax){#the-iframe-element:syntax} that consists of
the following syntactic components, in the given order:

1.  Any number of
    [comments](syntax.html#syntax-comments){#the-iframe-element:syntax-comments}
    and [ASCII
    whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#the-iframe-element:space-characters
    x-internal="space-characters"}.
2.  Optionally, a
    [DOCTYPE](syntax.html#syntax-doctype){#the-iframe-element:syntax-doctype}.
3.  Any number of
    [comments](syntax.html#syntax-comments){#the-iframe-element:syntax-comments-2}
    and [ASCII
    whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#the-iframe-element:space-characters-2
    x-internal="space-characters"}.
4.  The [document
    element](https://dom.spec.whatwg.org/#document-element){#the-iframe-element:document-element
    x-internal="document-element"}, in the form of an
    [`html`{#the-iframe-element:the-html-element}](semantics.html#the-html-element)
    [element](syntax.html#syntax-elements){#the-iframe-element:syntax-elements}.
5.  Any number of
    [comments](syntax.html#syntax-comments){#the-iframe-element:syntax-comments-3}
    and [ASCII
    whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#the-iframe-element:space-characters-3
    x-internal="space-characters"}.

The above requirements apply in [XML
documents](https://dom.spec.whatwg.org/#xml-document){#the-iframe-element:xml-documents
x-internal="xml-documents"} as well.

::: example
Here a blog uses the
[`srcdoc`{#the-iframe-element:attr-iframe-srcdoc-3}](#attr-iframe-srcdoc)
attribute in conjunction with the
[`sandbox`{#the-iframe-element:attr-iframe-sandbox-2}](#attr-iframe-sandbox)
attribute described below to provide users of user agents that support
this feature with an extra layer of protection from script injection in
the blog post comments:

``` html
<article>
 <h1>I got my own magazine!</h1>
 <p>After much effort, I've finally found a publisher, and so now I
 have my own magazine! Isn't that awesome?! The first issue will come
 out in September, and we have articles about getting food, and about
 getting in boxes, it's going to be great!</p>
 <footer>
  <p>Written by <a href="/users/cap">cap</a>, 1 hour ago.
 </footer>
 <article>
  <footer> Thirteen minutes ago, <a href="/users/ch">ch</a> wrote: </footer>
  <iframe sandbox srcdoc="<p>did you get a cover picture yet?"></iframe>
 </article>
 <article>
  <footer> Nine minutes ago, <a href="/users/cap">cap</a> wrote: </footer>
  <iframe sandbox srcdoc="<p>Yeah, you can see it <a href=&quot;/gallery?mode=cover&amp;amp;page=1&quot;>in my gallery</a>."></iframe>
 </article>
 <article>
  <footer> Five minutes ago, <a href="/users/ch">ch</a> wrote: </footer>
  <iframe sandbox srcdoc="<p>hey that's earl's table.
<p>you should get earl&amp;amp;me on the next cover."></iframe>
 </article>
```

Notice the way that quotes have to be escaped (otherwise the
[`srcdoc`{#the-iframe-element:attr-iframe-srcdoc-4}](#attr-iframe-srcdoc)
attribute would end prematurely), and the way raw ampersands (e.g. in
URLs or in prose) mentioned in the sandboxed content have to be *doubly*
escaped --- once so that the ampersand is preserved when originally
parsing the
[`srcdoc`{#the-iframe-element:attr-iframe-srcdoc-5}](#attr-iframe-srcdoc)
attribute, and once more to prevent the ampersand from being
misinterpreted when parsing the sandboxed content.

Furthermore, notice that since the
[DOCTYPE](syntax.html#syntax-doctype){#the-iframe-element:syntax-doctype-2}
is optional in [`iframe` `srcdoc`
documents](#an-iframe-srcdoc-document){#the-iframe-element:an-iframe-srcdoc-document},
and the
[`html`{#the-iframe-element:the-html-element-2}](semantics.html#the-html-element),
[`head`{#the-iframe-element:the-head-element}](semantics.html#the-head-element),
and
[`body`{#the-iframe-element:the-body-element}](sections.html#the-body-element)
elements have [optional start and end
tags](syntax.html#syntax-tag-omission), and the
[`title`{#the-iframe-element:the-title-element}](semantics.html#the-title-element)
element is also optional in [`iframe` `srcdoc`
documents](#an-iframe-srcdoc-document){#the-iframe-element:an-iframe-srcdoc-document-2},
the markup in a
[`srcdoc`{#the-iframe-element:attr-iframe-srcdoc-6}](#attr-iframe-srcdoc)
attribute can be relatively succinct despite representing an entire
document, since only the contents of the
[`body`{#the-iframe-element:the-body-element-2}](sections.html#the-body-element)
element need appear literally in the syntax. The other elements are
still present, but only by implication.
:::

In [the HTML syntax](syntax.html#syntax){#the-iframe-element:syntax-2},
authors need only remember to use U+0022 QUOTATION MARK characters (\")
to wrap the attribute contents and then to escape all U+0026 AMPERSAND
(&) and U+0022 QUOTATION MARK (\") characters, and to specify the
[`sandbox`{#the-iframe-element:attr-iframe-sandbox-3}](#attr-iframe-sandbox)
attribute, to ensure safe embedding of content. (And remember to escape
ampersands before quotation marks, to ensure quotation marks become
&quot; and not &amp;quot;.)

In XML the U+003C LESS-THAN SIGN character (\<) needs to be escaped as
well. In order to prevent [attribute-value
normalization](https://www.w3.org/TR/xml/#AVNormalize), some of XML\'s
whitespace characters --- specifically U+0009 CHARACTER TABULATION
(tab), U+000A LINE FEED (LF), and U+000D CARRIAGE RETURN (CR) --- also
need to be escaped. [\[XML\]](references.html#refsXML)

If the [`src`{#the-iframe-element:attr-iframe-src-3}](#attr-iframe-src)
attribute and the
[`srcdoc`{#the-iframe-element:attr-iframe-srcdoc-7}](#attr-iframe-srcdoc)
attribute are both specified together, the
[`srcdoc`{#the-iframe-element:attr-iframe-srcdoc-8}](#attr-iframe-srcdoc)
attribute takes priority. This allows authors to provide a fallback
[URL](https://url.spec.whatwg.org/#concept-url){#the-iframe-element:url-2
x-internal="url"} for legacy user agents that do not support the
[`srcdoc`{#the-iframe-element:attr-iframe-srcdoc-9}](#attr-iframe-srcdoc)
attribute.

------------------------------------------------------------------------

The
[`iframe`{#the-iframe-element:the-iframe-element-6}](#the-iframe-element)
[HTML element post-connection
steps](infrastructure.html#html-element-post-connection-steps){#the-iframe-element:html-element-post-connection-steps},
given `insertedNode`{.variable}, are:

1.  [Create a new child
    navigable](document-sequences.html#create-a-new-child-navigable){#the-iframe-element:create-a-new-child-navigable}
    for `insertedNode`{.variable}.

2.  If `insertedNode`{.variable} has a
    [`sandbox`{#the-iframe-element:attr-iframe-sandbox-4}](#attr-iframe-sandbox)
    attribute, then [parse the sandboxing
    directive](browsers.html#parse-a-sandboxing-directive){#the-iframe-element:parse-a-sandboxing-directive}
    given the attribute\'s value and `insertedNode`{.variable}\'s
    [`iframe` sandboxing flag
    set](browsers.html#iframe-sandboxing-flag-set){#the-iframe-element:iframe-sandboxing-flag-set}.

3.  [Process the `iframe`
    attributes](#process-the-iframe-attributes){#the-iframe-element:process-the-iframe-attributes}
    for `insertedNode`{.variable}, with
    *[initialInsertion](#process-iframe-initial-insertion)* set to true.

The
[`iframe`{#the-iframe-element:the-iframe-element-7}](#the-iframe-element)
[HTML element removing
steps](infrastructure.html#html-element-removing-steps){#the-iframe-element:html-element-removing-steps},
given `removedNode`{.variable}, are to [destroy a child
navigable](document-sequences.html#destroy-a-child-navigable){#the-iframe-element:destroy-a-child-navigable}
given `removedNode`{.variable}.

This happens without any
[`unload`{#the-iframe-element:event-unload}](indices.html#event-unload)
events firing (the element\'s [content
document](document-sequences.html#concept-bcc-content-document){#the-iframe-element:concept-bcc-content-document}
is
*[destroyed](document-sequences.html#destroy-a-child-navigable){#the-iframe-element:destroy-a-child-navigable-2}*,
not
*[unloaded](document-lifecycle.html#unload-a-document){#the-iframe-element:unload-a-document}*).

Although
[`iframe`{#the-iframe-element:the-iframe-element-8}](#the-iframe-element)s
are processed while in a [shadow
tree](https://dom.spec.whatwg.org/#concept-shadow-tree){#the-iframe-element:shadow-tree
x-internal="shadow-tree"}, per the above, several other aspects of their
behavior are not well-defined with regards to shadow trees. See [issue
#763](https://github.com/whatwg/html/issues/763) for more detail.

Whenever an
[`iframe`{#the-iframe-element:the-iframe-element-9}](#the-iframe-element)
element with a non-null [content
navigable](document-sequences.html#content-navigable){#the-iframe-element:content-navigable-5}
has its
[`srcdoc`{#the-iframe-element:attr-iframe-srcdoc-10}](#attr-iframe-srcdoc)
attribute set, changed, or removed, the user agent must [process the
`iframe`
attributes](#process-the-iframe-attributes){#the-iframe-element:process-the-iframe-attributes-2}.

Similarly, whenever an
[`iframe`{#the-iframe-element:the-iframe-element-10}](#the-iframe-element)
element with a non-null [content
navigable](document-sequences.html#content-navigable){#the-iframe-element:content-navigable-6}
but with no
[`srcdoc`{#the-iframe-element:attr-iframe-srcdoc-11}](#attr-iframe-srcdoc)
attribute specified has its
[`src`{#the-iframe-element:attr-iframe-src-4}](#attr-iframe-src)
attribute set, changed, or removed, the user agent must [process the
`iframe`
attributes](#process-the-iframe-attributes){#the-iframe-element:process-the-iframe-attributes-3}.

To [process the `iframe` attributes]{#process-the-iframe-attributes
.dfn} for an element `element`{.variable}, with an optional boolean
[`initialInsertion`{.variable}]{#process-iframe-initial-insertion .dfn}
(default false):

1.  If `element`{.variable}\'s
    [`srcdoc`{#the-iframe-element:attr-iframe-srcdoc-12}](#attr-iframe-srcdoc)
    attribute is specified, then:

    1.  Set `element`{.variable}\'s [current navigation was lazy
        loaded](#current-navigation-was-lazy-loaded){#the-iframe-element:current-navigation-was-lazy-loaded}
        boolean to false.

    2.  If the [will lazy load element
        steps](urls-and-fetching.html#will-lazy-load-element-steps){#the-iframe-element:will-lazy-load-element-steps}
        given `element`{.variable} return true, then:

        1.  Set `element`{.variable}\'s [lazy load resumption
            steps](urls-and-fetching.html#lazy-load-resumption-steps){#the-iframe-element:lazy-load-resumption-steps}
            to the rest of this algorithm starting with the step labeled
            *navigate to the srcdoc resource*.

        2.  Set `element`{.variable}\'s [current navigation was lazy
            loaded](#current-navigation-was-lazy-loaded){#the-iframe-element:current-navigation-was-lazy-loaded-2}
            boolean to true.

        3.  [Start intersection-observing a lazy loading
            element](urls-and-fetching.html#start-intersection-observing-a-lazy-loading-element){#the-iframe-element:start-intersection-observing-a-lazy-loading-element}
            for `element`{.variable}.

        4.  Return.

    3.  *Navigate to the srcdoc resource*: [Navigate an `iframe` or
        `frame`](#navigate-an-iframe-or-frame){#the-iframe-element:navigate-an-iframe-or-frame}
        given `element`{.variable},
        [`about:srcdoc`{#the-iframe-element:about:srcdoc}](urls-and-fetching.html#about:srcdoc),
        the empty string, and the value of `element`{.variable}\'s
        [`srcdoc`{#the-iframe-element:attr-iframe-srcdoc-13}](#attr-iframe-srcdoc)
        attribute.

        The resulting
        [`Document`{#the-iframe-element:document-4}](dom.html#document)
        must be considered [an `iframe` `srcdoc`
        document](#an-iframe-srcdoc-document){#the-iframe-element:an-iframe-srcdoc-document-3}.

2.  Otherwise:

    1.  Let `url`{.variable} be the result of running the [shared
        attribute processing steps for `iframe` and `frame`
        elements](#shared-attribute-processing-steps-for-iframe-and-frame-elements){#the-iframe-element:shared-attribute-processing-steps-for-iframe-and-frame-elements}
        given `element`{.variable} and `initialInsertion`{.variable}.

    2.  If `url`{.variable} is null, then return.

    3.  If `url`{.variable} [matches
        `about:blank`](urls-and-fetching.html#matches-about:blank){#the-iframe-element:matches-about:blank}
        and `initialInsertion`{.variable} is true, then:

        1.  Run the [iframe load event
            steps](#iframe-load-event-steps){#the-iframe-element:iframe-load-event-steps}
            given `element`{.variable}.

        2.  Return.

    4.  Let `referrerPolicy`{.variable} be the current state of
        `element`{.variable}\'s
        [`referrerpolicy`{#the-iframe-element:attr-iframe-referrerpolicy-2}](#attr-iframe-referrerpolicy)
        content attribute.

    5.  Set `element`{.variable}\'s [current navigation was lazy
        loaded](#current-navigation-was-lazy-loaded){#the-iframe-element:current-navigation-was-lazy-loaded-3}
        boolean to false.

    6.  If the [will lazy load element
        steps](urls-and-fetching.html#will-lazy-load-element-steps){#the-iframe-element:will-lazy-load-element-steps-2}
        given `element`{.variable} return true, then:

        1.  Set `element`{.variable}\'s [lazy load resumption
            steps](urls-and-fetching.html#lazy-load-resumption-steps){#the-iframe-element:lazy-load-resumption-steps-2}
            to the rest of this algorithm starting with the step labeled
            *navigate*.

        2.  Set `element`{.variable}\'s [current navigation was lazy
            loaded](#current-navigation-was-lazy-loaded){#the-iframe-element:current-navigation-was-lazy-loaded-4}
            boolean to true.

        3.  [Start intersection-observing a lazy loading
            element](urls-and-fetching.html#start-intersection-observing-a-lazy-loading-element){#the-iframe-element:start-intersection-observing-a-lazy-loading-element-2}
            for `element`{.variable}.

        4.  Return.

    7.  *Navigate*: [Navigate an `iframe` or
        `frame`](#navigate-an-iframe-or-frame){#the-iframe-element:navigate-an-iframe-or-frame-2}
        given `element`{.variable}, `url`{.variable}, and
        `referrerPolicy`{.variable}.

The [shared attribute processing steps for `iframe` and `frame`
elements]{#shared-attribute-processing-steps-for-iframe-and-frame-elements
.dfn}, given an element `element`{.variable} and a boolean
`initialInsertion`{.variable}, are:

1.  Let `url`{.variable} be the [URL
    record](https://url.spec.whatwg.org/#concept-url){#the-iframe-element:url-record
    x-internal="url-record"}
    [`about:blank`{#the-iframe-element:about:blank}](infrastructure.html#about:blank).

2.  If `element`{.variable} has a
    [`src`{#the-iframe-element:attr-iframe-src-5}](#attr-iframe-src)
    attribute specified, and its value is not the empty string, then:

    1.  Let `maybeURL`{.variable} be the result of [encoding-parsing a
        URL](urls-and-fetching.html#encoding-parsing-a-url){#the-iframe-element:encoding-parsing-a-url}
        given that attribute\'s value, relative to
        `element`{.variable}\'s [node
        document](https://dom.spec.whatwg.org/#concept-node-document){#the-iframe-element:node-document
        x-internal="node-document"}.

    2.  If `maybeURL`{.variable} is not failure, then set
        `url`{.variable} to `maybeURL`{.variable}.

3.  If the [inclusive ancestor
    navigables](document-sequences.html#inclusive-ancestor-navigables){#the-iframe-element:inclusive-ancestor-navigables}
    of `element`{.variable}\'s [node
    navigable](document-sequences.html#node-navigable){#the-iframe-element:node-navigable}
    contains a
    [navigable](document-sequences.html#navigable){#the-iframe-element:navigable}
    whose [active
    document](document-sequences.html#nav-document){#the-iframe-element:nav-document}\'s
    [URL](https://dom.spec.whatwg.org/#concept-document-url){#the-iframe-element:the-document's-address-2
    x-internal="the-document's-address"}
    [equals](https://url.spec.whatwg.org/#concept-url-equals){#the-iframe-element:concept-url-equals
    x-internal="concept-url-equals"} `url`{.variable} with *[exclude
    fragments](https://url.spec.whatwg.org/#url-equals-exclude-fragments){x-internal="url-equals-exclude-fragments"}*
    set to true, then return null.

4.  If `url`{.variable} [matches
    `about:blank`](urls-and-fetching.html#matches-about:blank){#the-iframe-element:matches-about:blank-2}
    and `initialInsertion`{.variable} is true, then perform the [URL and
    history update
    steps](browsing-the-web.html#url-and-history-update-steps){#the-iframe-element:url-and-history-update-steps}
    given `element`{.variable}\'s [content
    navigable](document-sequences.html#content-navigable){#the-iframe-element:content-navigable-7}\'s
    [active
    document](document-sequences.html#nav-document){#the-iframe-element:nav-document-2}
    and `url`{.variable}.

    This is necessary in case `url`{.variable} is something like
    `about:blank?foo`. If `url`{.variable} is just plain `about:blank`,
    this will do nothing.

5.  Return `url`{.variable}.

To [navigate an `iframe` or `frame`]{#navigate-an-iframe-or-frame .dfn}
given an element `element`{.variable}, a
[URL](https://url.spec.whatwg.org/#concept-url){#the-iframe-element:url-3
x-internal="url"} `url`{.variable}, a [referrer
policy](https://w3c.github.io/webappsec-referrer-policy/#referrer-policy){#the-iframe-element:referrer-policy-2
x-internal="referrer-policy"} `referrerPolicy`{.variable}, an optional
string-or-null `srcdocString`{.variable} (default null), and an optional
boolean `initialInsertion`{.variable} (default false):

1.  Let `historyHandling`{.variable} be
    \"[`auto`{#the-iframe-element:navigationhistorybehavior-auto}](browsing-the-web.html#navigationhistorybehavior-auto)\".

2.  If `element`{.variable}\'s [content
    navigable](document-sequences.html#content-navigable){#the-iframe-element:content-navigable-8}\'s
    [active
    document](document-sequences.html#nav-document){#the-iframe-element:nav-document-3}
    is not [completely
    loaded](document-lifecycle.html#completely-loaded){#the-iframe-element:completely-loaded},
    then set `historyHandling`{.variable} to
    \"[`replace`{#the-iframe-element:navigationhistorybehavior-replace}](browsing-the-web.html#navigationhistorybehavior-replace)\".

3.  If `element`{.variable} is an
    [`iframe`{#the-iframe-element:the-iframe-element-11}](#the-iframe-element),
    then set `element`{.variable}\'s [pending resource-timing start
    time](#iframe-pending-resource-timing-start-time){#the-iframe-element:iframe-pending-resource-timing-start-time}
    to the [current high resolution
    time](https://w3c.github.io/hr-time/#dfn-current-high-resolution-time){#the-iframe-element:current-high-resolution-time
    x-internal="current-high-resolution-time"} given
    `element`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-iframe-element:node-document-2
    x-internal="node-document"}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#the-iframe-element:concept-relevant-global}.

4.  [Navigate](browsing-the-web.html#navigate){#the-iframe-element:navigate}
    `element`{.variable}\'s [content
    navigable](document-sequences.html#content-navigable){#the-iframe-element:content-navigable-9}
    to `url`{.variable} using `element`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-iframe-element:node-document-3
    x-internal="node-document"}, with
    *[historyHandling](browsing-the-web.html#navigation-hh)* set to
    `historyHandling`{.variable},
    *[referrerPolicy](browsing-the-web.html#navigation-referrer-policy)*
    set to `referrerPolicy`{.variable},
    *[documentResource](browsing-the-web.html#navigation-resource)* set
    to `srcdocString`{.variable}, and
    *[initialInsertion](browsing-the-web.html#navigation-initial-insertion)*
    set to `initialInsertion`{.variable}.

Each [`Document`{#the-iframe-element:document-5}](dom.html#document) has
an [iframe load in progress]{#iframe-load-in-progress .dfn} flag and a
[mute iframe load]{#mute-iframe-load .dfn} flag. When a
[`Document`{#the-iframe-element:document-6}](dom.html#document) is
created, these flags must be unset for that
[`Document`{#the-iframe-element:document-7}](dom.html#document).

To run the [iframe load event steps]{#iframe-load-event-steps .dfn},
given an
[`iframe`{#the-iframe-element:the-iframe-element-12}](#the-iframe-element)
element `element`{.variable}:

1.  [Assert](https://infra.spec.whatwg.org/#assert){#the-iframe-element:assert
    x-internal="assert"}: `element`{.variable}\'s [content
    navigable](document-sequences.html#content-navigable){#the-iframe-element:content-navigable-10}
    is not null.

2.  Let `childDocument`{.variable} be `element`{.variable}\'s [content
    navigable](document-sequences.html#content-navigable){#the-iframe-element:content-navigable-11}\'s
    [active
    document](document-sequences.html#nav-document){#the-iframe-element:nav-document-4}.

3.  If `childDocument`{.variable} has its [mute iframe
    load](#mute-iframe-load){#the-iframe-element:mute-iframe-load} flag
    set, then return.

4.  If `element`{.variable}\'s [pending resource-timing start
    time](#iframe-pending-resource-timing-start-time){#the-iframe-element:iframe-pending-resource-timing-start-time-2}
    is not null, then:

    1.  Let `global`{.variable} be `element`{.variable}\'s [node
        document](https://dom.spec.whatwg.org/#concept-node-document){#the-iframe-element:node-document-4
        x-internal="node-document"}\'s [relevant global
        object](webappapis.html#concept-relevant-global){#the-iframe-element:concept-relevant-global-2}.

    2.  Let `fallbackTimingInfo`{.variable} be a new [fetch timing
        info](https://fetch.spec.whatwg.org/#fetch-timing-info){#the-iframe-element:fetch-timing-info
        x-internal="fetch-timing-info"} whose [start
        time](https://fetch.spec.whatwg.org/#fetch-timing-info-start-time){#the-iframe-element:fetch-timing-info-start-time
        x-internal="fetch-timing-info-start-time"} is
        `element`{.variable}\'s [pending resource-timing start
        time](#iframe-pending-resource-timing-start-time){#the-iframe-element:iframe-pending-resource-timing-start-time-3}
        and whose [response end
        time](https://fetch.spec.whatwg.org/#fetch-timing-info-end-time){#the-iframe-element:fetch-timing-info-end-time
        x-internal="fetch-timing-info-end-time"} is the [current high
        resolution
        time](https://w3c.github.io/hr-time/#dfn-current-high-resolution-time){#the-iframe-element:current-high-resolution-time-2
        x-internal="current-high-resolution-time"} given
        `global`{.variable}.

    3.  [Mark resource
        timing](https://w3c.github.io/resource-timing/#dfn-mark-resource-timing){#the-iframe-element:mark-resource-timing
        x-internal="mark-resource-timing"} given
        `fallbackTimingInfo`{.variable}, `url`{.variable},
        \"[`iframe`{#the-iframe-element:the-iframe-element-13}](#the-iframe-element)\",
        `global`{.variable}, the empty string, a new [response body
        info](https://fetch.spec.whatwg.org/#response-body-info){#the-iframe-element:response-body-info
        x-internal="response-body-info"}, and 0.

    4.  Set `element`{.variable}\'s [pending resource-timing start
        time](#iframe-pending-resource-timing-start-time){#the-iframe-element:iframe-pending-resource-timing-start-time-4}
        to null.

5.  Set `childDocument`{.variable}\'s [iframe load in
    progress](#iframe-load-in-progress){#the-iframe-element:iframe-load-in-progress}
    flag.

6.  [Fire an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#the-iframe-element:concept-event-fire
    x-internal="concept-event-fire"} named
    [`load`{#the-iframe-element:event-load}](indices.html#event-load) at
    `element`{.variable}.

7.  Unset `childDocument`{.variable}\'s [iframe load in
    progress](#iframe-load-in-progress){#the-iframe-element:iframe-load-in-progress-2}
    flag.

This, in conjunction with scripting, can be used to probe the URL space
of the local network\'s HTTP servers. User agents may implement
[cross-origin](browsers.html#concept-origin){#the-iframe-element:concept-origin}
access control policies that are stricter than those described above to
mitigate this attack, but unfortunately such policies are typically not
compatible with existing web content.

If an element type [potentially delays the load
event]{#potentially-delays-the-load-event .dfn}, then for each element
`element`{.variable} of that type, the user agent must [delay the load
event](parsing.html#delay-the-load-event){#the-iframe-element:delay-the-load-event}
of `element`{.variable}\'s [node
document](https://dom.spec.whatwg.org/#concept-node-document){#the-iframe-element:node-document-5
x-internal="node-document"} if `element`{.variable}\'s [content
navigable](document-sequences.html#content-navigable){#the-iframe-element:content-navigable-12}
is non-null and any of the following are true:

- `element`{.variable}\'s [content
  navigable](document-sequences.html#content-navigable){#the-iframe-element:content-navigable-13}\'s
  [active
  document](document-sequences.html#nav-document){#the-iframe-element:nav-document-5}
  is not [ready for post-load
  tasks](parsing.html#ready-for-post-load-tasks){#the-iframe-element:ready-for-post-load-tasks};

- `element`{.variable}\'s [content
  navigable](document-sequences.html#content-navigable){#the-iframe-element:content-navigable-14}\'s
  [is delaying `load`
  events](document-sequences.html#delaying-load-events-mode){#the-iframe-element:delaying-load-events-mode}
  is true; or

- anything is [delaying the load
  event](parsing.html#delay-the-load-event){#the-iframe-element:delay-the-load-event-2}
  of `element`{.variable}\'s [content
  navigable](document-sequences.html#content-navigable){#the-iframe-element:content-navigable-15}\'s
  [active
  document](document-sequences.html#nav-document){#the-iframe-element:nav-document-6}.

If, during the handling of the
[`load`{#the-iframe-element:event-load-2}](indices.html#event-load)
event, `element`{.variable}\'s [content
navigable](document-sequences.html#content-navigable){#the-iframe-element:content-navigable-16}
is again
[navigated](browsing-the-web.html#navigate){#the-iframe-element:navigate-2},
that will further [delay the load
event](parsing.html#delay-the-load-event){#the-iframe-element:delay-the-load-event-3}.

Each
[`iframe`{#the-iframe-element:the-iframe-element-14}](#the-iframe-element)
element has an associated [current navigation was lazy
loaded]{#current-navigation-was-lazy-loaded .dfn} boolean, initially
false. It is set and unset in the [process the `iframe`
attributes](#process-the-iframe-attributes){#the-iframe-element:process-the-iframe-attributes-4}
algorithm.

An
[`iframe`{#the-iframe-element:the-iframe-element-15}](#the-iframe-element)
element whose [current navigation was lazy
loaded](#current-navigation-was-lazy-loaded){#the-iframe-element:current-navigation-was-lazy-loaded-5}
boolean is false [potentially delays the load
event](#potentially-delays-the-load-event){#the-iframe-element:potentially-delays-the-load-event}.

Each
[`iframe`{#the-iframe-element:the-iframe-element-16}](#the-iframe-element)
element has an associated null or
[`DOMHighResTimeStamp`](https://w3c.github.io/hr-time/#dom-domhighrestimestamp){#the-iframe-element:domhighrestimestamp
x-internal="domhighrestimestamp"} [pending resource-timing start
time]{#iframe-pending-resource-timing-start-time .dfn}, initially set to
null.

If, when the element is created, the
[`srcdoc`{#the-iframe-element:attr-iframe-srcdoc-14}](#attr-iframe-srcdoc)
attribute is not set, and the
[`src`{#the-iframe-element:attr-iframe-src-6}](#attr-iframe-src)
attribute is either also not set or set but its value cannot be
[parsed](urls-and-fetching.html#encoding-parsing-a-url){#the-iframe-element:encoding-parsing-a-url-2},
the element\'s [content
navigable](document-sequences.html#content-navigable){#the-iframe-element:content-navigable-17}
will remain at the [initial
`about:blank`](dom.html#is-initial-about:blank){#the-iframe-element:is-initial-about:blank}
[`Document`{#the-iframe-element:document-8}](dom.html#document).

If the user
[navigates](browsing-the-web.html#navigate){#the-iframe-element:navigate-3}
away from this page, the
[`iframe`{#the-iframe-element:the-iframe-element-17}](#the-iframe-element)\'s
[content
navigable](document-sequences.html#content-navigable){#the-iframe-element:content-navigable-18}\'s
[active
`WindowProxy`](document-sequences.html#nav-wp){#the-iframe-element:nav-wp}
object will proxy new
[`Window`{#the-iframe-element:window}](nav-history-apis.html#window)
objects for new
[`Document`{#the-iframe-element:document-9}](dom.html#document) objects,
but the [`src`{#the-iframe-element:attr-iframe-src-7}](#attr-iframe-src)
attribute will not change.

------------------------------------------------------------------------

The [`name`]{#attr-iframe-name .dfn dfn-for="iframe"
dfn-type="element-attr"} attribute, if present, must be a [valid
navigable target
name](document-sequences.html#valid-navigable-target-name){#the-iframe-element:valid-navigable-target-name}.
The given value is used to name the element\'s [content
navigable](document-sequences.html#content-navigable){#the-iframe-element:content-navigable-19}
if present when that is
[created](document-sequences.html#create-a-new-child-navigable){#the-iframe-element:create-a-new-child-navigable-2}.

------------------------------------------------------------------------

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Element/iframe#attr-sandbox](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/iframe#attr-sandbox "The <iframe> HTML element represents a nested browsing context, embedding another HTML page into the current one.")

Support in all current engines.

::: support
[Firefox17+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

The [`sandbox`]{#attr-iframe-sandbox .dfn dfn-for="iframe"
dfn-type="element-attr"} attribute, when specified, enables a set of
extra restrictions on any content hosted by the
[`iframe`{#the-iframe-element:the-iframe-element-18}](#the-iframe-element).
Its value must be an [unordered set of unique space-separated
tokens](common-microsyntaxes.html#unordered-set-of-unique-space-separated-tokens){#the-iframe-element:unordered-set-of-unique-space-separated-tokens}
that are [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#the-iframe-element:ascii-case-insensitive
x-internal="ascii-case-insensitive"}. The allowed values are:

- [`allow-downloads`{#the-iframe-element:attr-iframe-sandbox-allow-downloads}](browsers.html#attr-iframe-sandbox-allow-downloads)
- [`allow-forms`{#the-iframe-element:attr-iframe-sandbox-allow-forms}](browsers.html#attr-iframe-sandbox-allow-forms)
- [`allow-modals`{#the-iframe-element:attr-iframe-sandbox-allow-modals}](browsers.html#attr-iframe-sandbox-allow-modals)
- [`allow-orientation-lock`{#the-iframe-element:attr-iframe-sandbox-allow-orientation-lock}](browsers.html#attr-iframe-sandbox-allow-orientation-lock)
- [`allow-pointer-lock`{#the-iframe-element:attr-iframe-sandbox-allow-pointer-lock}](browsers.html#attr-iframe-sandbox-allow-pointer-lock)
- [`allow-popups`{#the-iframe-element:attr-iframe-sandbox-allow-popups}](browsers.html#attr-iframe-sandbox-allow-popups)
- [`allow-popups-to-escape-sandbox`{#the-iframe-element:attr-iframe-sandbox-allow-popups-to-escape-sandbox}](browsers.html#attr-iframe-sandbox-allow-popups-to-escape-sandbox)
- [`allow-presentation`{#the-iframe-element:attr-iframe-sandbox-allow-presentation}](browsers.html#attr-iframe-sandbox-allow-presentation)
- [`allow-same-origin`{#the-iframe-element:attr-iframe-sandbox-allow-same-origin}](browsers.html#attr-iframe-sandbox-allow-same-origin)
- [`allow-scripts`{#the-iframe-element:attr-iframe-sandbox-allow-scripts}](browsers.html#attr-iframe-sandbox-allow-scripts)
- [`allow-top-navigation`{#the-iframe-element:attr-iframe-sandbox-allow-top-navigation}](browsers.html#attr-iframe-sandbox-allow-top-navigation)
- [`allow-top-navigation-by-user-activation`{#the-iframe-element:attr-iframe-sandbox-allow-top-navigation-by-user-activation}](browsers.html#attr-iframe-sandbox-allow-top-navigation-by-user-activation)
- [`allow-top-navigation-to-custom-protocols`{#the-iframe-element:attr-iframe-sandbox-allow-top-navigation-to-custom-protocols}](browsers.html#attr-iframe-sandbox-allow-top-navigation-to-custom-protocols)

When the attribute is set, the content is treated as being from a unique
[opaque
origin](browsers.html#concept-origin-opaque){#the-iframe-element:concept-origin-opaque},
forms, scripts, and various potentially annoying APIs are disabled, and
links are prevented from targeting other
[navigables](document-sequences.html#navigable){#the-iframe-element:navigable-2}.
The
[`allow-same-origin`{#the-iframe-element:attr-iframe-sandbox-allow-same-origin-2}](browsers.html#attr-iframe-sandbox-allow-same-origin)
keyword causes the content to be treated as being from its real origin
instead of forcing it into an [opaque
origin](browsers.html#concept-origin-opaque){#the-iframe-element:concept-origin-opaque-2};
the
[`allow-top-navigation`{#the-iframe-element:attr-iframe-sandbox-allow-top-navigation-2}](browsers.html#attr-iframe-sandbox-allow-top-navigation)
keyword allows the content to
[navigate](browsing-the-web.html#navigate){#the-iframe-element:navigate-4}
its [traversable
navigable](document-sequences.html#nav-traversable){#the-iframe-element:nav-traversable};
the
[`allow-top-navigation-by-user-activation`{#the-iframe-element:attr-iframe-sandbox-allow-top-navigation-by-user-activation-2}](browsers.html#attr-iframe-sandbox-allow-top-navigation-by-user-activation)
keyword behaves similarly but allows such
[navigation](browsing-the-web.html#navigate){#the-iframe-element:navigate-5}
only when the browsing context\'s [active
window](document-sequences.html#nav-window){#the-iframe-element:nav-window}
has [transient
activation](interaction.html#transient-activation){#the-iframe-element:transient-activation};
the
[`allow-top-navigation-to-custom-protocols`{#the-iframe-element:attr-iframe-sandbox-allow-top-navigation-to-custom-protocols-2}](browsers.html#attr-iframe-sandbox-allow-top-navigation-to-custom-protocols)
reenables navigations toward non [fetch
scheme](https://fetch.spec.whatwg.org/#fetch-scheme){#the-iframe-element:fetch-scheme
x-internal="fetch-scheme"} to be [handed off to external
software](browsing-the-web.html#hand-off-to-external-software){#the-iframe-element:hand-off-to-external-software};
and the
[`allow-forms`{#the-iframe-element:attr-iframe-sandbox-allow-forms-2}](browsers.html#attr-iframe-sandbox-allow-forms),
[`allow-modals`{#the-iframe-element:attr-iframe-sandbox-allow-modals-2}](browsers.html#attr-iframe-sandbox-allow-modals),
[`allow-orientation-lock`{#the-iframe-element:attr-iframe-sandbox-allow-orientation-lock-2}](browsers.html#attr-iframe-sandbox-allow-orientation-lock),
[`allow-pointer-lock`{#the-iframe-element:attr-iframe-sandbox-allow-pointer-lock-2}](browsers.html#attr-iframe-sandbox-allow-pointer-lock),
[`allow-popups`{#the-iframe-element:attr-iframe-sandbox-allow-popups-2}](browsers.html#attr-iframe-sandbox-allow-popups),
[`allow-presentation`{#the-iframe-element:attr-iframe-sandbox-allow-presentation-2}](browsers.html#attr-iframe-sandbox-allow-presentation),
[`allow-scripts`{#the-iframe-element:attr-iframe-sandbox-allow-scripts-2}](browsers.html#attr-iframe-sandbox-allow-scripts),
and
[`allow-popups-to-escape-sandbox`{#the-iframe-element:attr-iframe-sandbox-allow-popups-to-escape-sandbox-2}](browsers.html#attr-iframe-sandbox-allow-popups-to-escape-sandbox)
keywords re-enable forms, modal dialogs, screen orientation lock, the
pointer lock API, popups, the presentation API, scripts, and the
creation of unsandboxed [auxiliary browsing
contexts](document-sequences.html#auxiliary-browsing-context){#the-iframe-element:auxiliary-browsing-context}
respectively. The
[`allow-downloads`{#the-iframe-element:attr-iframe-sandbox-allow-downloads-2}](browsers.html#attr-iframe-sandbox-allow-downloads)
keyword allows content to perform downloads.
[\[POINTERLOCK\]](references.html#refsPOINTERLOCK)
[\[SCREENORIENTATION\]](references.html#refsSCREENORIENTATION)
[\[PRESENTATION\]](references.html#refsPRESENTATION)

The
[`allow-top-navigation`{#the-iframe-element:attr-iframe-sandbox-allow-top-navigation-3}](browsers.html#attr-iframe-sandbox-allow-top-navigation)
and
[`allow-top-navigation-by-user-activation`{#the-iframe-element:attr-iframe-sandbox-allow-top-navigation-by-user-activation-3}](browsers.html#attr-iframe-sandbox-allow-top-navigation-by-user-activation)
keywords must not both be specified, as doing so is redundant; only
[`allow-top-navigation`{#the-iframe-element:attr-iframe-sandbox-allow-top-navigation-4}](browsers.html#attr-iframe-sandbox-allow-top-navigation)
will have an effect in such non-conformant markup.

Similarly, the
[`allow-top-navigation-to-custom-protocols`{#the-iframe-element:attr-iframe-sandbox-allow-top-navigation-to-custom-protocols-3}](browsers.html#attr-iframe-sandbox-allow-top-navigation-to-custom-protocols)
keyword must not be specified if either
[`allow-top-navigation`{#the-iframe-element:attr-iframe-sandbox-allow-top-navigation-5}](browsers.html#attr-iframe-sandbox-allow-top-navigation)
or
[`allow-popups`{#the-iframe-element:attr-iframe-sandbox-allow-popups-3}](browsers.html#attr-iframe-sandbox-allow-popups)
are specified, as doing so is redundant.

To allow
[`alert()`{#the-iframe-element:dom-alert}](timers-and-user-prompts.html#dom-alert),
[`confirm()`{#the-iframe-element:dom-confirm}](timers-and-user-prompts.html#dom-confirm),
and
[`prompt()`{#the-iframe-element:dom-prompt}](timers-and-user-prompts.html#dom-prompt)
inside sandboxed content, both the
[`allow-modals`{#the-iframe-element:attr-iframe-sandbox-allow-modals-3}](browsers.html#attr-iframe-sandbox-allow-modals)
and
[`allow-same-origin`{#the-iframe-element:attr-iframe-sandbox-allow-same-origin-3}](browsers.html#attr-iframe-sandbox-allow-same-origin)
keywords need to be specified, and the loaded URL needs to be [same
origin](browsers.html#same-origin){#the-iframe-element:same-origin} with
the [top-level
origin](webappapis.html#concept-environment-top-level-origin){#the-iframe-element:concept-environment-top-level-origin}.
Without the
[`allow-same-origin`{#the-iframe-element:attr-iframe-sandbox-allow-same-origin-4}](browsers.html#attr-iframe-sandbox-allow-same-origin)
keyword, the content is always treated as cross-origin, and cross-origin
content [cannot show simple
dialogs](timers-and-user-prompts.html#cannot-show-simple-dialogs){#the-iframe-element:cannot-show-simple-dialogs}.

Setting both the
[`allow-scripts`{#the-iframe-element:attr-iframe-sandbox-allow-scripts-3}](browsers.html#attr-iframe-sandbox-allow-scripts)
and
[`allow-same-origin`{#the-iframe-element:attr-iframe-sandbox-allow-same-origin-5}](browsers.html#attr-iframe-sandbox-allow-same-origin)
keywords together when the embedded page has the [same
origin](browsers.html#same-origin){#the-iframe-element:same-origin-2} as
the page containing the
[`iframe`{#the-iframe-element:the-iframe-element-19}](#the-iframe-element)
allows the embedded page to simply remove the
[`sandbox`{#the-iframe-element:attr-iframe-sandbox-5}](#attr-iframe-sandbox)
attribute and then reload itself, effectively breaking out of the
sandbox altogether.

These flags only take effect when the [content
navigable](document-sequences.html#content-navigable){#the-iframe-element:content-navigable-20}
of the
[`iframe`{#the-iframe-element:the-iframe-element-20}](#the-iframe-element)
element is
[navigated](browsing-the-web.html#navigate){#the-iframe-element:navigate-6}.
Removing them, or removing the entire
[`sandbox`{#the-iframe-element:attr-iframe-sandbox-6}](#attr-iframe-sandbox)
attribute, has no effect on an already-loaded page.

Potentially hostile files should not be served from the same server as
the file containing the
[`iframe`{#the-iframe-element:the-iframe-element-21}](#the-iframe-element)
element. Sandboxing hostile content is of minimal help if an attacker
can convince the user to just visit the hostile content directly, rather
than in the
[`iframe`{#the-iframe-element:the-iframe-element-22}](#the-iframe-element).
To limit the damage that can be caused by hostile HTML content, it
should be served from a separate dedicated domain. Using a different
domain ensures that scripts in the files are unable to attack the site,
even if the user is tricked into visiting those pages directly, without
the protection of the
[`sandbox`{#the-iframe-element:attr-iframe-sandbox-7}](#attr-iframe-sandbox)
attribute.

When an
[`iframe`{#the-iframe-element:the-iframe-element-23}](#the-iframe-element)
element\'s
[`sandbox`{#the-iframe-element:attr-iframe-sandbox-8}](#attr-iframe-sandbox)
attribute is set or changed while it has a non-null [content
navigable](document-sequences.html#content-navigable){#the-iframe-element:content-navigable-21},
the user agent must [parse the sandboxing
directive](browsers.html#parse-a-sandboxing-directive){#the-iframe-element:parse-a-sandboxing-directive-2}
given the attribute\'s value and the
[`iframe`{#the-iframe-element:the-iframe-element-24}](#the-iframe-element)
element\'s [`iframe` sandboxing flag
set](browsers.html#iframe-sandboxing-flag-set){#the-iframe-element:iframe-sandboxing-flag-set-2}.

When an
[`iframe`{#the-iframe-element:the-iframe-element-25}](#the-iframe-element)
element\'s
[`sandbox`{#the-iframe-element:attr-iframe-sandbox-9}](#attr-iframe-sandbox)
attribute is removed while it has a non-null [content
navigable](document-sequences.html#content-navigable){#the-iframe-element:content-navigable-22},
the user agent must empty the
[`iframe`{#the-iframe-element:the-iframe-element-26}](#the-iframe-element)
element\'s [`iframe` sandboxing flag
set](browsers.html#iframe-sandboxing-flag-set){#the-iframe-element:iframe-sandboxing-flag-set-3}.

::: example
In this example, some completely-unknown, potentially hostile,
user-provided HTML content is embedded in a page. Because it is served
from a separate domain, it is affected by all the normal cross-site
restrictions. In addition, the embedded page has scripting disabled,
plugins disabled, forms disabled, and it cannot navigate any frames or
windows other than itself (or any frames or windows it itself embeds).

``` html
<p>We're not scared of you! Here is your content, unedited:</p>
<iframe sandbox src="https://usercontent.example.net/getusercontent.cgi?id=12193"></iframe>
```

It is important to use a separate domain so that if the attacker
convinces the user to visit that page directly, the page doesn\'t run in
the context of the site\'s origin, which would make the user vulnerable
to any attack found in the page.
:::

::: example
In this example, a gadget from another site is embedded. The gadget has
scripting and forms enabled, and the origin sandbox restrictions are
lifted, allowing the gadget to communicate with its originating server.
The sandbox is still useful, however, as it disables plugins and popups,
thus reducing the risk of the user being exposed to malware and other
annoyances.

``` html
<iframe sandbox="allow-same-origin allow-forms allow-scripts"
        src="https://maps.example.com/embedded.html"></iframe>
```
:::

::: example
Suppose a file A contained the following fragment:

``` html
<iframe sandbox="allow-same-origin allow-forms" src=B></iframe>
```

Suppose that file B contained an iframe also:

``` html
<iframe sandbox="allow-scripts" src=C></iframe>
```

Further, suppose that file C contained a link:

``` html
<a href=D>Link</a>
```

For this example, suppose all the files were served as
[`text/html`{#the-iframe-element:text/html}](iana.html#text/html).

Page C in this scenario has all the sandboxing flags set. Scripts are
disabled, because the
[`iframe`{#the-iframe-element:the-iframe-element-27}](#the-iframe-element)
in A has scripts disabled, and this overrides the
[`allow-scripts`{#the-iframe-element:attr-iframe-sandbox-allow-scripts-4}](browsers.html#attr-iframe-sandbox-allow-scripts)
keyword set on the
[`iframe`{#the-iframe-element:the-iframe-element-28}](#the-iframe-element)
in B. Forms are also disabled, because the inner
[`iframe`{#the-iframe-element:the-iframe-element-29}](#the-iframe-element)
(in B) does not have the
[`allow-forms`{#the-iframe-element:attr-iframe-sandbox-allow-forms-3}](browsers.html#attr-iframe-sandbox-allow-forms)
keyword set.

Suppose now that a script in A removes all the
[`sandbox`{#the-iframe-element:attr-iframe-sandbox-10}](#attr-iframe-sandbox)
attributes in A and B. This would change nothing immediately. If the
user clicked the link in C, loading page D into the
[`iframe`{#the-iframe-element:the-iframe-element-30}](#the-iframe-element)
in B, page D would now act as if the
[`iframe`{#the-iframe-element:the-iframe-element-31}](#the-iframe-element)
in B had the
[`allow-same-origin`{#the-iframe-element:attr-iframe-sandbox-allow-same-origin-6}](browsers.html#attr-iframe-sandbox-allow-same-origin)
and
[`allow-forms`{#the-iframe-element:attr-iframe-sandbox-allow-forms-4}](browsers.html#attr-iframe-sandbox-allow-forms)
keywords set, because that was the state of the [content
navigable](document-sequences.html#content-navigable){#the-iframe-element:content-navigable-23}
in the
[`iframe`{#the-iframe-element:the-iframe-element-32}](#the-iframe-element)
in A when page B was loaded.

Generally speaking, dynamically removing or changing the
[`sandbox`{#the-iframe-element:attr-iframe-sandbox-11}](#attr-iframe-sandbox)
attribute is ill-advised, because it can make it quite hard to reason
about what will be allowed and what will not.
:::

------------------------------------------------------------------------

The [`allow`]{#attr-iframe-allow .dfn dfn-for="iframe"
dfn-type="element-attr"} attribute, when specified, determines the
[container
policy](https://w3c.github.io/webappsec-feature-policy/#container-policy){#the-iframe-element:concept-container-policy
x-internal="concept-container-policy"} that will be used when the
[permissions
policy](dom.html#concept-document-permissions-policy){#the-iframe-element:concept-document-permissions-policy}
for a [`Document`{#the-iframe-element:document-10}](dom.html#document)
in the
[`iframe`{#the-iframe-element:the-iframe-element-33}](#the-iframe-element)\'s
[content
navigable](document-sequences.html#content-navigable){#the-iframe-element:content-navigable-24}
is initialized. Its value must be a [serialized permissions
policy](https://w3c.github.io/webappsec-feature-policy/#serialized-permissions-policy){#the-iframe-element:concept-serialized-permissions-policy
x-internal="concept-serialized-permissions-policy"}.
[\[PERMISSIONSPOLICY\]](references.html#refsPERMISSIONSPOLICY)

::: example
In this example, an
[`iframe`{#the-iframe-element:the-iframe-element-34}](#the-iframe-element)
is used to embed a map from an online navigation service. The
[`allow`{#the-iframe-element:attr-iframe-allow-2}](#attr-iframe-allow)
attribute is used to enable the Geolocation API within the nested
context.

``` html
<iframe src="https://maps.example.com/" allow="geolocation"></iframe>
```
:::

The [`allowfullscreen`]{#attr-iframe-allowfullscreen .dfn
dfn-for="iframe" dfn-type="element-attr"} attribute is a [boolean
attribute](common-microsyntaxes.html#boolean-attribute){#the-iframe-element:boolean-attribute}.
When specified, it indicates that
[`Document`{#the-iframe-element:document-11}](dom.html#document) objects
in the
[`iframe`{#the-iframe-element:the-iframe-element-35}](#the-iframe-element)
element\'s [content
navigable](document-sequences.html#content-navigable){#the-iframe-element:content-navigable-25}
will be initialized with a [permissions
policy](dom.html#concept-document-permissions-policy){#the-iframe-element:concept-document-permissions-policy-2}
which allows the \"`fullscreen`\" feature to be used from any
[origin](browsers.html#concept-origin){#the-iframe-element:concept-origin-2}.
This is enforced by the [process permissions policy
attributes](https://w3c.github.io/webappsec-feature-policy/#process-permissions-policy-attributes){#the-iframe-element:process-permissions-policy-attributes
x-internal="process-permissions-policy-attributes"} algorithm.
[\[PERMISSIONSPOLICY\]](references.html#refsPERMISSIONSPOLICY)

::: example
Here, an
[`iframe`{#the-iframe-element:the-iframe-element-36}](#the-iframe-element)
is used to embed a player from a video site. The
[`allowfullscreen`{#the-iframe-element:attr-iframe-allowfullscreen-2}](#attr-iframe-allowfullscreen)
attribute is needed to enable the player to show its video fullscreen.

``` html
<article>
 <header>
  <p><img src="/usericons/1627591962735"> <b>Fred Flintstone</b></p>
  <p><a href="/posts/3095182851" rel=bookmark>12:44</a> — <a href="#acl-3095182851">Private Post</a></p>
 </header>
 <p>Check out my new ride!</p>
 <iframe src="https://video.example.com/embed?id=92469812" allowfullscreen></iframe>
</article>
```
:::

Neither
[`allow`{#the-iframe-element:attr-iframe-allow-3}](#attr-iframe-allow)
nor
[`allowfullscreen`{#the-iframe-element:attr-iframe-allowfullscreen-3}](#attr-iframe-allowfullscreen)
can grant access to a feature in an
[`iframe`{#the-iframe-element:the-iframe-element-37}](#the-iframe-element)
element\'s [content
navigable](document-sequences.html#content-navigable){#the-iframe-element:content-navigable-26}
if the element\'s [node
document](https://dom.spec.whatwg.org/#concept-node-document){#the-iframe-element:node-document-6
x-internal="node-document"} is not already allowed to use that feature.

To determine whether a
[`Document`{#the-iframe-element:document-12}](dom.html#document) object
`document`{.variable} is [allowed to use]{#allowed-to-use .dfn
export=""} the policy-controlled-feature `feature`{.variable}, run these
steps:

1.  If `document`{.variable}\'s [browsing
    context](document-sequences.html#concept-document-bc){#the-iframe-element:concept-document-bc}
    is null, then return false.

2.  If `document`{.variable} is not [fully
    active](document-sequences.html#fully-active){#the-iframe-element:fully-active},
    then return false.

3.  If the result of running [is feature enabled in document for
    origin](https://w3c.github.io/webappsec-feature-policy/#is-feature-enabled){#the-iframe-element:is-feature-enabled
    x-internal="is-feature-enabled"} on `feature`{.variable},
    `document`{.variable}, and `document`{.variable}\'s
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#the-iframe-element:concept-document-origin
    x-internal="concept-document-origin"} is \"`Enabled`\", then return
    true.

4.  Return false.

Because they only influence the [permissions
policy](dom.html#concept-document-permissions-policy){#the-iframe-element:concept-document-permissions-policy-3}
of the [content
navigable](document-sequences.html#content-navigable){#the-iframe-element:content-navigable-27}\'s
[active
document](document-sequences.html#nav-document){#the-iframe-element:nav-document-7},
the
[`allow`{#the-iframe-element:attr-iframe-allow-4}](#attr-iframe-allow)
and
[`allowfullscreen`{#the-iframe-element:attr-iframe-allowfullscreen-4}](#attr-iframe-allowfullscreen)
attributes only take effect when the [content
navigable](document-sequences.html#content-navigable){#the-iframe-element:content-navigable-28}
of the
[`iframe`{#the-iframe-element:the-iframe-element-38}](#the-iframe-element)
is
[navigated](browsing-the-web.html#navigate){#the-iframe-element:navigate-7}.
Adding or removing them has no effect on an already-loaded document.

------------------------------------------------------------------------

The
[`iframe`{#the-iframe-element:the-iframe-element-39}](#the-iframe-element)
element supports [dimension
attributes](embedded-content-other.html#dimension-attributes){#the-iframe-element:dimension-attributes}
for cases where the embedded content has specific dimensions (e.g. ad
units have well-defined dimensions).

An
[`iframe`{#the-iframe-element:the-iframe-element-40}](#the-iframe-element)
element never has [fallback
content](dom.html#fallback-content){#the-iframe-element:fallback-content},
as it will always [create a new child
navigable](document-sequences.html#create-a-new-child-navigable){#the-iframe-element:create-a-new-child-navigable-3},
regardless of whether the specified initial contents are successfully
used.

------------------------------------------------------------------------

The [`referrerpolicy`]{#attr-iframe-referrerpolicy .dfn dfn-for="iframe"
dfn-type="element-attr"} attribute is a [referrer policy
attribute](urls-and-fetching.html#referrer-policy-attribute){#the-iframe-element:referrer-policy-attribute}.
Its purpose is to set the [referrer
policy](https://w3c.github.io/webappsec-referrer-policy/#referrer-policy){#the-iframe-element:referrer-policy-3
x-internal="referrer-policy"} used when [processing the `iframe`
attributes](#process-the-iframe-attributes){#the-iframe-element:process-the-iframe-attributes-5}.
[\[REFERRERPOLICY\]](references.html#refsREFERRERPOLICY)

The [`loading`]{#attr-iframe-loading .dfn dfn-for="iframe"
dfn-type="element-attr"} attribute is a [lazy loading
attribute](urls-and-fetching.html#lazy-loading-attribute){#the-iframe-element:lazy-loading-attribute}.
Its purpose is to indicate the policy for loading
[`iframe`{#the-iframe-element:the-iframe-element-41}](#the-iframe-element)
elements that are outside the viewport.

When the
[`loading`{#the-iframe-element:attr-iframe-loading-2}](#attr-iframe-loading)
attribute\'s state is changed to the
[Eager](urls-and-fetching.html#attr-loading-eager-state){#the-iframe-element:attr-loading-eager-state}
state, the user agent must run these steps:

1.  Let `resumptionSteps`{.variable} be the
    [`iframe`{#the-iframe-element:the-iframe-element-42}](#the-iframe-element)
    element\'s [lazy load resumption
    steps](urls-and-fetching.html#lazy-load-resumption-steps){#the-iframe-element:lazy-load-resumption-steps-3}.

2.  If `resumptionSteps`{.variable} is null, then return.

3.  Set the
    [`iframe`{#the-iframe-element:the-iframe-element-43}](#the-iframe-element)\'s
    [lazy load resumption
    steps](urls-and-fetching.html#lazy-load-resumption-steps){#the-iframe-element:lazy-load-resumption-steps-4}
    to null.

4.  Invoke `resumptionSteps`{.variable}.

------------------------------------------------------------------------

Descendants of
[`iframe`{#the-iframe-element:the-iframe-element-44}](#the-iframe-element)
elements represent nothing. (In legacy user agents that do not support
[`iframe`{#the-iframe-element:the-iframe-element-45}](#the-iframe-element)
elements, the contents would be parsed as markup that could act as
fallback content.)

The [HTML
parser](parsing.html#html-parser){#the-iframe-element:html-parser}
treats markup inside
[`iframe`{#the-iframe-element:the-iframe-element-46}](#the-iframe-element)
elements as text.

------------------------------------------------------------------------

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLIFrameElement/src](https://developer.mozilla.org/en-US/docs/Web/API/HTMLIFrameElement/src "The HTMLIFrameElement.src property reflects the HTML referrerpolicy attribute of the <iframe> element defining which referrer is sent when fetching the resource.")

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

The IDL attributes [`src`]{#dom-iframe-src .dfn
dfn-for="HTMLIFrameElement" dfn-type="attribute"},
[`name`]{#dom-iframe-name .dfn dfn-for="HTMLIFrameElement"
dfn-type="attribute"}, [`sandbox`]{#dom-iframe-sandbox .dfn
dfn-for="HTMLIFrameElement" dfn-type="attribute"}, and
[`allow`]{#dom-iframe-allow .dfn dfn-for="HTMLIFrameElement"
dfn-type="attribute"} must
[reflect](common-dom-interfaces.html#reflect){#the-iframe-element:reflect}
the respective content attributes of the same name.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLIFrameElement/srcdoc](https://developer.mozilla.org/en-US/docs/Web/API/HTMLIFrameElement/srcdoc "The srcdoc property of the HTMLIFrameElement specifies the content of the page.")

Support in all current engines.

::: support
[Firefox25+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome20+]{.chrome
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

The [`srcdoc`]{#dom-iframe-srcdoc .dfn dfn-for="HTMLIFrameElement"
dfn-type="attribute"} getter steps are:

1.  Let `attribute`{.variable} be the result of running [get an
    attribute by namespace and local
    name](https://dom.spec.whatwg.org/#concept-element-attributes-get-by-namespace){#the-iframe-element:concept-element-attributes-get-by-namespace
    x-internal="concept-element-attributes-get-by-namespace"} given
    null,
    [`srcdoc`{#the-iframe-element:attr-iframe-srcdoc-15}](#attr-iframe-srcdoc)\'s
    [local
    name](https://dom.spec.whatwg.org/#concept-attribute-local-name){#the-iframe-element:concept-attribute-local-name
    x-internal="concept-attribute-local-name"}, and
    [this](https://webidl.spec.whatwg.org/#this){#the-iframe-element:this
    x-internal="this"}.

2.  If `attribute`{.variable} is null, then return the empty string.

3.  Return `attribute`{.variable}\'s
    [value](https://dom.spec.whatwg.org/#concept-attribute-value){#the-iframe-element:concept-attribute-value
    x-internal="concept-attribute-value"}.

The
[`srcdoc`{#the-iframe-element:dom-iframe-srcdoc-2}](#dom-iframe-srcdoc)
setter steps are:

1.  Let `compliantString`{.variable} be the result of invoking the [Get
    Trusted Type compliant
    string](https://w3c.github.io/trusted-types/dist/spec/#get-trusted-type-compliant-string-algorithm){#the-iframe-element:tt-getcompliantstring
    x-internal="tt-getcompliantstring"} algorithm with
    [`TrustedHTML`{#the-iframe-element:tt-trustedhtml-2}](https://w3c.github.io/trusted-types/dist/spec/#trustedhtml){x-internal="tt-trustedhtml"},
    [this](https://webidl.spec.whatwg.org/#this){#the-iframe-element:this-2
    x-internal="this"}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#the-iframe-element:concept-relevant-global-3},
    the given value, \"`HTMLIFrameElement srcdoc`\", and \"`script`\".

2.  [Set an attribute
    value](https://dom.spec.whatwg.org/#concept-element-attributes-set-value){#the-iframe-element:concept-element-attributes-set-value
    x-internal="concept-element-attributes-set-value"} given
    [this](https://webidl.spec.whatwg.org/#this){#the-iframe-element:this-3
    x-internal="this"},
    [`srcdoc`{#the-iframe-element:attr-iframe-srcdoc-16}](#attr-iframe-srcdoc)\'s
    [local
    name](https://dom.spec.whatwg.org/#concept-attribute-local-name){#the-iframe-element:concept-attribute-local-name-2
    x-internal="concept-attribute-local-name"}, and
    `compliantString`{.variable}.

The [supported
tokens](https://dom.spec.whatwg.org/#concept-supported-tokens){#the-iframe-element:concept-supported-tokens
x-internal="concept-supported-tokens"} for
[`sandbox`{#the-iframe-element:dom-iframe-sandbox-2}](#dom-iframe-sandbox)\'s
[`DOMTokenList`{#the-iframe-element:domtokenlist-2}](https://dom.spec.whatwg.org/#interface-domtokenlist){x-internal="domtokenlist"}
are the allowed values defined in the
[`sandbox`{#the-iframe-element:attr-iframe-sandbox-12}](#attr-iframe-sandbox)
attribute and supported by the user agent.

The [`allowFullscreen`]{#dom-iframe-allowfullscreen .dfn
dfn-for="HTMLIFrameElement" dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-iframe-element:reflect-2}
the
[`allowfullscreen`{#the-iframe-element:attr-iframe-allowfullscreen-5}](#attr-iframe-allowfullscreen)
content attribute.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLIFrameElement/referrerPolicy](https://developer.mozilla.org/en-US/docs/Web/API/HTMLIFrameElement/referrerPolicy "The HTMLIFrameElement.referrerPolicy property reflects the HTML referrerpolicy attribute of the <iframe> element defining which referrer is sent when fetching the resource.")

Support in all current engines.

::: support
[Firefox50+]{.firefox .yes}[Safari14+]{.safari .yes}[Chrome52+]{.chrome
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

The [`referrerPolicy`]{#dom-iframe-referrerpolicy .dfn
dfn-for="HTMLIFrameElement" dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-iframe-element:reflect-3}
the
[`referrerpolicy`{#the-iframe-element:attr-iframe-referrerpolicy-3}](#attr-iframe-referrerpolicy)
content attribute, [limited to only known
values](common-dom-interfaces.html#limited-to-only-known-values){#the-iframe-element:limited-to-only-known-values}.

The [`loading`]{#dom-iframe-loading .dfn dfn-for="HTMLIFrameElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-iframe-element:reflect-4}
the
[`loading`{#the-iframe-element:attr-iframe-loading-3}](#attr-iframe-loading)
content attribute, [limited to only known
values](common-dom-interfaces.html#limited-to-only-known-values){#the-iframe-element:limited-to-only-known-values-2}.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLIFrameElement/contentDocument](https://developer.mozilla.org/en-US/docs/Web/API/HTMLIFrameElement/contentDocument "If the iframe and the iframe's parent document are Same Origin, returns a Document (that is, the active document in the inline frame's nested browsing context), else returns null.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer8+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The [`contentDocument`]{#dom-iframe-contentdocument .dfn
dfn-for="HTMLIFrameElement" dfn-type="attribute"} getter steps are to
return
[this](https://webidl.spec.whatwg.org/#this){#the-iframe-element:this-4
x-internal="this"}\'s [content
document](document-sequences.html#concept-bcc-content-document){#the-iframe-element:concept-bcc-content-document-2}.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLIFrameElement/contentWindow](https://developer.mozilla.org/en-US/docs/Web/API/HTMLIFrameElement/contentWindow "The contentWindow property returns the Window object of an HTMLIFrameElement. You can use this Window object to access the iframe's document and its internal DOM. This attribute is read-only, but its properties can be manipulated like the global Window object.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera8+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

The [`contentWindow`]{#dom-iframe-contentwindow .dfn
dfn-for="HTMLIFrameElement" dfn-type="attribute"} getter steps are to
return
[this](https://webidl.spec.whatwg.org/#this){#the-iframe-element:this-5
x-internal="this"}\'s [content
window](document-sequences.html#content-window){#the-iframe-element:content-window}.

::: example
Here is an example of a page using an
[`iframe`{#the-iframe-element:the-iframe-element-47}](#the-iframe-element)
to include advertising from an advertising broker:

``` html
<iframe src="https://ads.example.com/?customerid=923513721&amp;format=banner"
        width="468" height="60"></iframe>
```
:::

#### [4.8.6]{.secno} The [`embed`]{.dfn dfn-type="element"} element[](#the-embed-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/embed](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/embed "The <embed> HTML element embeds external content at the specified point in the document. This content is provided by an external application or other source of interactive content such as a browser plug-in.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerYes]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLEmbedElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLEmbedElement "The HTMLEmbedElement interface provides special properties (beyond the regular HTMLElement interface it also has available to it by inheritance) for manipulating <embed> elements.")

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

[Categories](dom.html#concept-element-categories){#the-embed-element:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-embed-element:flow-content-2}.
:   [Phrasing
    content](dom.html#phrasing-content-2){#the-embed-element:phrasing-content-2}.
:   [Embedded
    content](dom.html#embedded-content-category){#the-embed-element:embedded-content-category}.
:   [Interactive
    content](dom.html#interactive-content-2){#the-embed-element:interactive-content-2}.
:   [Palpable
    content](dom.html#palpable-content-2){#the-embed-element:palpable-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-embed-element:concept-element-contexts}:
:   Where [embedded
    content](dom.html#embedded-content-category){#the-embed-element:embedded-content-category-2}
    is expected.

[Content model](dom.html#concept-element-content-model){#the-embed-element:concept-element-content-model}:
:   [Nothing](dom.html#concept-content-nothing){#the-embed-element:concept-content-nothing}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-embed-element:concept-element-tag-omission}:
:   No [end
    tag](syntax.html#syntax-end-tag){#the-embed-element:syntax-end-tag}.

[Content attributes](dom.html#concept-element-attributes){#the-embed-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-embed-element:global-attributes}
:   [`src`{#the-embed-element:attr-embed-src}](#attr-embed-src) ---
    Address of the resource
:   [`type`{#the-embed-element:attr-embed-type}](#attr-embed-type) ---
    Type of embedded resource
:   [`width`{#the-embed-element:attr-dim-width}](embedded-content-other.html#attr-dim-width)
    --- Horizontal dimension
:   [`height`{#the-embed-element:attr-dim-height}](embedded-content-other.html#attr-dim-height)
    --- Vertical dimension
:   Any other attribute that has no namespace (see prose).

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-embed-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-embed).
:   [For implementers](https://w3c.github.io/html-aam/#el-embed).

[DOM interface](dom.html#concept-element-dom){#the-embed-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLEmbedElement : HTMLElement {
      [HTMLConstructor] constructor();

      [CEReactions] attribute USVString src;
      [CEReactions] attribute DOMString type;
      [CEReactions] attribute DOMString width;
      [CEReactions] attribute DOMString height;
      Document? getSVGDocument();

      // also has obsolete members
    };
    ```

The [`embed`{#the-embed-element:the-embed-element}](#the-embed-element)
element provides an integration point for an external application or
interactive content.

The [`src`]{#attr-embed-src .dfn dfn-for="embed"
dfn-type="element-attr"} attribute gives the
[URL](https://url.spec.whatwg.org/#concept-url){#the-embed-element:url
x-internal="url"} of the resource being embedded. The attribute, if
present, must contain a [valid non-empty URL potentially surrounded by
spaces](urls-and-fetching.html#valid-non-empty-url-potentially-surrounded-by-spaces){#the-embed-element:valid-non-empty-url-potentially-surrounded-by-spaces}.

If the
[`itemprop`{#the-embed-element:names:-the-itemprop-attribute}](microdata.html#names:-the-itemprop-attribute)
attribute is specified on an
[`embed`{#the-embed-element:the-embed-element-2}](#the-embed-element)
element, then the
[`src`{#the-embed-element:attr-embed-src-2}](#attr-embed-src) attribute
must also be specified.

The [`type`]{#attr-embed-type .dfn dfn-for="embed"
dfn-type="element-attr"} attribute, if present, gives the [MIME
type](https://mimesniff.spec.whatwg.org/#mime-type){#the-embed-element:mime-type
x-internal="mime-type"} by which the plugin to instantiate is selected.
The value must be a [valid MIME type
string](https://mimesniff.spec.whatwg.org/#valid-mime-type){#the-embed-element:valid-mime-type-string
x-internal="valid-mime-type-string"}. If both the
[`type`{#the-embed-element:attr-embed-type-2}](#attr-embed-type)
attribute and the
[`src`{#the-embed-element:attr-embed-src-3}](#attr-embed-src) attribute
are present, then the
[`type`{#the-embed-element:attr-embed-type-3}](#attr-embed-type)
attribute must specify the same type as the [explicit Content-Type
metadata](urls-and-fetching.html#content-type){#the-embed-element:content-type}
of the resource given by the
[`src`{#the-embed-element:attr-embed-src-4}](#attr-embed-src) attribute.

While any of the following conditions are occurring, any
[plugin](infrastructure.html#plugin){#the-embed-element:plugin}
instantiated for the element must be removed, and the
[`embed`{#the-embed-element:the-embed-element-3}](#the-embed-element)
element [represents](dom.html#represents){#the-embed-element:represents}
nothing:

- The element has neither a
  [`src`{#the-embed-element:attr-embed-src-5}](#attr-embed-src)
  attribute nor a
  [`type`{#the-embed-element:attr-embed-type-4}](#attr-embed-type)
  attribute.

- The element has a [media
  element](media.html#media-element){#the-embed-element:media-element}
  ancestor.

- The element has an ancestor
  [`object`{#the-embed-element:the-object-element}](#the-object-element)
  element that is *not* showing its [fallback
  content](dom.html#fallback-content){#the-embed-element:fallback-content}.

An [`embed`{#the-embed-element:the-embed-element-4}](#the-embed-element)
element is said to be [potentially active]{#concept-embed-active .dfn}
when the following conditions are all met simultaneously:

- The element is [in a
  document](https://dom.spec.whatwg.org/#in-a-document){#the-embed-element:in-a-document
  x-internal="in-a-document"} or was [in a
  document](https://dom.spec.whatwg.org/#in-a-document){#the-embed-element:in-a-document-2
  x-internal="in-a-document"} the last time the [event
  loop](webappapis.html#event-loop){#the-embed-element:event-loop}
  reached [step 1](webappapis.html#step1).

- The element\'s [node
  document](https://dom.spec.whatwg.org/#concept-node-document){#the-embed-element:node-document
  x-internal="node-document"} is [fully
  active](document-sequences.html#fully-active){#the-embed-element:fully-active}.

- The element has either a
  [`src`{#the-embed-element:attr-embed-src-6}](#attr-embed-src)
  attribute set or a
  [`type`{#the-embed-element:attr-embed-type-5}](#attr-embed-type)
  attribute set (or both).

- The element\'s
  [`src`{#the-embed-element:attr-embed-src-7}](#attr-embed-src)
  attribute is either absent or its value is not the empty string.

- The element is not a descendant of a [media
  element](media.html#media-element){#the-embed-element:media-element-2}.

- The element is not a descendant of an
  [`object`{#the-embed-element:the-object-element-2}](#the-object-element)
  element that is not showing its [fallback
  content](dom.html#fallback-content){#the-embed-element:fallback-content-2}.

- The element is [being
  rendered](rendering.html#being-rendered){#the-embed-element:being-rendered},
  or was [being
  rendered](rendering.html#being-rendered){#the-embed-element:being-rendered-2}
  the last time the [event
  loop](webappapis.html#event-loop){#the-embed-element:event-loop-2}
  reached [step 1](webappapis.html#step1).

Whenever an
[`embed`{#the-embed-element:the-embed-element-5}](#the-embed-element)
element that was not [potentially
active](#concept-embed-active){#the-embed-element:concept-embed-active}
becomes [potentially
active](#concept-embed-active){#the-embed-element:concept-embed-active-2},
and whenever a [potentially
active](#concept-embed-active){#the-embed-element:concept-embed-active-3}
[`embed`{#the-embed-element:the-embed-element-6}](#the-embed-element)
element that is remaining [potentially
active](#concept-embed-active){#the-embed-element:concept-embed-active-4}
and has its
[`src`{#the-embed-element:attr-embed-type-6}](#attr-embed-type)
attribute set, changed, or removed or its
[`type`{#the-embed-element:attr-embed-type-7}](#attr-embed-type)
attribute set, changed, or removed, the user agent must [queue an
element
task](webappapis.html#queue-an-element-task){#the-embed-element:queue-an-element-task}
on the [embed task source]{#embed-task-source .dfn} given the element to
run [the `embed` element setup
steps](#the-embed-element-setup-steps){#the-embed-element:the-embed-element-setup-steps}
for that element.

[The `embed` element setup steps]{#the-embed-element-setup-steps .dfn}
for a given
[`embed`{#the-embed-element:the-embed-element-7}](#the-embed-element)
element `element`{.variable} are as follows:

1.  If another
    [task](webappapis.html#concept-task){#the-embed-element:concept-task}
    has since been queued to run [the `embed` element setup
    steps](#the-embed-element-setup-steps){#the-embed-element:the-embed-element-setup-steps-2}
    for `element`{.variable}, then return.

2.  If `element`{.variable} has a
    [`src`{#the-embed-element:attr-embed-src-8}](#attr-embed-src)
    attribute set, then:

    1.  Let `url`{.variable} be the result of [encoding-parsing a
        URL](urls-and-fetching.html#encoding-parsing-a-url){#the-embed-element:encoding-parsing-a-url}
        given `element`{.variable}\'s
        [`src`{#the-embed-element:attr-embed-src-9}](#attr-embed-src)
        attribute\'s value, relative to `element`{.variable}\'s [node
        document](https://dom.spec.whatwg.org/#concept-node-document){#the-embed-element:node-document-2
        x-internal="node-document"}.

    2.  If `url`{.variable} is failure, then return.

    3.  Let `request`{.variable} be a new
        [request](https://fetch.spec.whatwg.org/#concept-request){#the-embed-element:concept-request
        x-internal="concept-request"} whose
        [URL](https://fetch.spec.whatwg.org/#concept-request-url){#the-embed-element:concept-request-url
        x-internal="concept-request-url"} is `url`{.variable},
        [client](https://fetch.spec.whatwg.org/#concept-request-client){#the-embed-element:concept-request-client
        x-internal="concept-request-client"} is `element`{.variable}\'s
        [node
        document](https://dom.spec.whatwg.org/#concept-node-document){#the-embed-element:node-document-3
        x-internal="node-document"}\'s [relevant settings
        object](webappapis.html#relevant-settings-object){#the-embed-element:relevant-settings-object},
        [destination](https://fetch.spec.whatwg.org/#concept-request-destination){#the-embed-element:concept-request-destination
        x-internal="concept-request-destination"} is \"`embed`\",
        [credentials
        mode](https://fetch.spec.whatwg.org/#concept-request-credentials-mode){#the-embed-element:concept-request-credentials-mode
        x-internal="concept-request-credentials-mode"} is \"`include`\",
        [mode](https://fetch.spec.whatwg.org/#concept-request-mode){#the-embed-element:concept-request-mode
        x-internal="concept-request-mode"} is \"`navigate`\", [initiator
        type](https://fetch.spec.whatwg.org/#request-initiator-type){#the-embed-element:concept-request-initiator-type
        x-internal="concept-request-initiator-type"} is \"`embed`\", and
        whose [use-URL-credentials
        flag](https://fetch.spec.whatwg.org/#concept-request-use-url-credentials-flag){#the-embed-element:use-url-credentials-flag
        x-internal="use-url-credentials-flag"} is set.

    4.  [Fetch](https://fetch.spec.whatwg.org/#concept-fetch){#the-embed-element:concept-fetch
        x-internal="concept-fetch"} `request`{.variable}, with
        *[processResponse](https://fetch.spec.whatwg.org/#process-response){x-internal="processresponse"}*
        set to the following steps given
        [response](https://fetch.spec.whatwg.org/#concept-response){#the-embed-element:concept-response
        x-internal="concept-response"} `response`{.variable}:

        1.  If another
            [task](webappapis.html#concept-task){#the-embed-element:concept-task-2}
            has since been queued to run [the `embed` element setup
            steps](#the-embed-element-setup-steps){#the-embed-element:the-embed-element-setup-steps-3}
            for `element`{.variable}, then return.

        2.  If `response`{.variable} is a [network
            error](https://fetch.spec.whatwg.org/#concept-network-error){#the-embed-element:network-error
            x-internal="network-error"}, then [fire an
            event](https://dom.spec.whatwg.org/#concept-event-fire){#the-embed-element:concept-event-fire
            x-internal="concept-event-fire"} named
            [`load`{#the-embed-element:event-load}](indices.html#event-load)
            at `element`{.variable}, and return.

        3.  Let `type`{.variable} be the result of determining the [type
            of
            content](#concept-embed-type){#the-embed-element:concept-embed-type}
            given `element`{.variable} and `response`{.variable}.

        4.  Switch on `type`{.variable}:

            null

            :   1.  [Display no
                    plugin](#display-no-plugin){#the-embed-element:display-no-plugin}
                    for `element`{.variable}.

            Otherwise

            :   1.  If `element`{.variable}\'s [content
                    navigable](document-sequences.html#content-navigable){#the-embed-element:content-navigable}
                    is null, then [create a new child
                    navigable](document-sequences.html#create-a-new-child-navigable){#the-embed-element:create-a-new-child-navigable}
                    for `element`{.variable}.

                2.  [Navigate](browsing-the-web.html#navigate){#the-embed-element:navigate}
                    `element`{.variable}\'s [content
                    navigable](document-sequences.html#content-navigable){#the-embed-element:content-navigable-2}
                    to `response`{.variable}\'s
                    [URL](https://fetch.spec.whatwg.org/#concept-response-url){#the-embed-element:concept-response-url
                    x-internal="concept-response-url"} using
                    `element`{.variable}\'s [node
                    document](https://dom.spec.whatwg.org/#concept-node-document){#the-embed-element:node-document-4
                    x-internal="node-document"}, with
                    *[response](browsing-the-web.html#navigation-response)*
                    set to `response`{.variable}, and
                    *[historyHandling](browsing-the-web.html#navigation-hh)*
                    set to
                    \"[`replace`{#the-embed-element:navigationhistorybehavior-replace}](browsing-the-web.html#navigationhistorybehavior-replace)\".

                    `element`{.variable}\'s
                    [`src`{#the-embed-element:attr-embed-src-10}](#attr-embed-src)
                    attribute does not get updated if the [content
                    navigable](document-sequences.html#content-navigable){#the-embed-element:content-navigable-3}
                    gets further navigated to other locations.

                3.  `element`{.variable} now
                    [represents](dom.html#represents){#the-embed-element:represents-2}
                    its [content
                    navigable](document-sequences.html#content-navigable){#the-embed-element:content-navigable-4}.

        Fetching the resource must [delay the load
        event](parsing.html#delay-the-load-event){#the-embed-element:delay-the-load-event}
        of `element`{.variable}\'s [node
        document](https://dom.spec.whatwg.org/#concept-node-document){#the-embed-element:node-document-5
        x-internal="node-document"}.

3.  Otherwise, [display no
    plugin](#display-no-plugin){#the-embed-element:display-no-plugin-2}
    for `element`{.variable}.

To determine the [type of the content]{#concept-embed-type .dfn} given
an [`embed`{#the-embed-element:the-embed-element-8}](#the-embed-element)
element `element`{.variable} and a
[response](https://fetch.spec.whatwg.org/#concept-response){#the-embed-element:concept-response-2
x-internal="concept-response"} `response`{.variable}, run the following
steps:

1.  If `element`{.variable} has a
    [`type`{#the-embed-element:attr-embed-type-8}](#attr-embed-type)
    attribute, and that attribute\'s value is a type that a
    [plugin](infrastructure.html#plugin){#the-embed-element:plugin-2}
    supports, then return the value of the
    [`type`{#the-embed-element:attr-embed-type-9}](#attr-embed-type)
    attribute.

2.  If the
    [path](https://url.spec.whatwg.org/#concept-url-path){#the-embed-element:concept-url-path
    x-internal="concept-url-path"} component of `response`{.variable}\'s
    [url](https://fetch.spec.whatwg.org/#concept-response-url){#the-embed-element:concept-response-url-2
    x-internal="concept-response-url"} matches a pattern that a
    [plugin](infrastructure.html#plugin){#the-embed-element:plugin-3}
    supports, then return the type that that plugin can handle.

    For example, a plugin might say that it can handle URLs with
    [path](https://url.spec.whatwg.org/#concept-url-path){#the-embed-element:concept-url-path-2
    x-internal="concept-url-path"} components that end with the four
    character string \"`.swf`\".

3.  If `response`{.variable} has [explicit Content-Type
    metadata](urls-and-fetching.html#content-type){#the-embed-element:content-type-2},
    and that value is a type that a
    [plugin](infrastructure.html#plugin){#the-embed-element:plugin-4}
    supports, then return that value.

4.  Return null.

It is intentional that the above algorithm allows `response`{.variable}
to have a non-[ok
status](https://fetch.spec.whatwg.org/#ok-status){#the-embed-element:ok-status
x-internal="ok-status"}. This allows servers to return data for plugins
even with error responses (e.g., HTTP 500 Internal Server Error codes
can still contain plugin data).

To [display no plugin]{#display-no-plugin .dfn} for an
[`embed`{#the-embed-element:the-embed-element-9}](#the-embed-element)
element `element`{.variable}:

1.  [Destroy a child
    navigable](document-sequences.html#destroy-a-child-navigable){#the-embed-element:destroy-a-child-navigable}
    given `element`{.variable}.

2.  Display an indication that no
    [plugin](infrastructure.html#plugin){#the-embed-element:plugin-5}
    could be found for `element`{.variable}, as the contents of
    `element`{.variable}.

3.  `element`{.variable} now
    [represents](dom.html#represents){#the-embed-element:represents-3}
    nothing.

The
[`embed`{#the-embed-element:the-embed-element-10}](#the-embed-element)
element has no [fallback
content](dom.html#fallback-content){#the-embed-element:fallback-content-3};
its descendants are ignored.

Whenever an
[`embed`{#the-embed-element:the-embed-element-11}](#the-embed-element)
element that was [potentially
active](#concept-embed-active){#the-embed-element:concept-embed-active-5}
stops being [potentially
active](#concept-embed-active){#the-embed-element:concept-embed-active-6},
any [plugin](infrastructure.html#plugin){#the-embed-element:plugin-6}
that had been instantiated for that element must be unloaded.

The
[`embed`{#the-embed-element:the-embed-element-12}](#the-embed-element)
element [potentially delays the load
event](#potentially-delays-the-load-event){#the-embed-element:potentially-delays-the-load-event}.

The
[`embed`{#the-embed-element:the-embed-element-13}](#the-embed-element)
element supports [dimension
attributes](embedded-content-other.html#dimension-attributes){#the-embed-element:dimension-attributes}.

The IDL attributes [`src`]{#dom-embed-src .dfn
dfn-for="HTMLEmbedElement" dfn-type="attribute"} and
[`type`]{#dom-embed-type .dfn dfn-for="HTMLEmbedElement"
dfn-type="attribute"} each must
[reflect](common-dom-interfaces.html#reflect){#the-embed-element:reflect}
the respective content attributes of the same name.

#### [4.8.7]{.secno} The [`object`]{.dfn dfn-type="element"} element[](#the-object-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/object](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/object "The <object> HTML element represents an external resource, which can be treated as an image, a nested browsing context, or a resource to be handled by a plugin.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
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

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLObjectElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLObjectElement "The HTMLObjectElement interface provides special properties and methods (beyond those on the HTMLElement interface it also has available to it by inheritance) for manipulating the layout and presentation of <object> element, representing external resources.")

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

[Categories](dom.html#concept-element-categories){#the-object-element:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-object-element:flow-content-2}.
:   [Phrasing
    content](dom.html#phrasing-content-2){#the-object-element:phrasing-content-2}.
:   [Embedded
    content](dom.html#embedded-content-category){#the-object-element:embedded-content-category}.
:   [Listed](forms.html#category-listed){#the-object-element:category-listed}
    [form-associated
    element](forms.html#form-associated-element){#the-object-element:form-associated-element}.
:   [Palpable
    content](dom.html#palpable-content-2){#the-object-element:palpable-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-object-element:concept-element-contexts}:
:   Where [embedded
    content](dom.html#embedded-content-category){#the-object-element:embedded-content-category-2}
    is expected.

[Content model](dom.html#concept-element-content-model){#the-object-element:concept-element-content-model}:
:   [Transparent](dom.html#transparent){#the-object-element:transparent}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-object-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-object-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-object-element:global-attributes}
:   [`data`{#the-object-element:attr-object-data}](#attr-object-data)
    --- Address of the resource
:   [`type`{#the-object-element:attr-object-type}](#attr-object-type)
    --- Type of embedded resource
:   [`name`{#the-object-element:attr-object-name}](#attr-object-name)
    --- Name of [content
    navigable](document-sequences.html#content-navigable){#the-object-element:content-navigable}
:   [`form`{#the-object-element:attr-fae-form}](form-control-infrastructure.html#attr-fae-form)
    --- Associates the element with a
    [`form`{#the-object-element:the-form-element}](forms.html#the-form-element)
    element
:   [`width`{#the-object-element:attr-dim-width}](embedded-content-other.html#attr-dim-width)
    --- Horizontal dimension
:   [`height`{#the-object-element:attr-dim-height}](embedded-content-other.html#attr-dim-height)
    --- Vertical dimension

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-object-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-object).
:   [For implementers](https://w3c.github.io/html-aam/#el-object).

[DOM interface](dom.html#concept-element-dom){#the-object-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLObjectElement : HTMLElement {
      [HTMLConstructor] constructor();

      [CEReactions] attribute USVString data;
      [CEReactions] attribute DOMString type;
      [CEReactions] attribute DOMString name;
      readonly attribute HTMLFormElement? form;
      [CEReactions] attribute DOMString width;
      [CEReactions] attribute DOMString height;
      readonly attribute Document? contentDocument;
      readonly attribute WindowProxy? contentWindow;
      Document? getSVGDocument();

      readonly attribute boolean willValidate;
      readonly attribute ValidityState validity;
      readonly attribute DOMString validationMessage;
      boolean checkValidity();
      boolean reportValidity();
      undefined setCustomValidity(DOMString error);

      // also has obsolete members
    };
    ```

    Depending on the type of content instantiated by the
    [`object`{#the-object-element:the-object-element}](#the-object-element)
    element, the node also supports other interfaces.

The
[`object`{#the-object-element:the-object-element-2}](#the-object-element)
element can represent an external resource, which, depending on the type
of the resource, will either be treated as an image or as a [child
navigable](document-sequences.html#child-navigable){#the-object-element:child-navigable}.

The [`data`]{#attr-object-data .dfn dfn-for="object"
dfn-type="element-attr"} attribute specifies the
[URL](https://url.spec.whatwg.org/#concept-url){#the-object-element:url
x-internal="url"} of the resource. It must be present, and must contain
a [valid non-empty URL potentially surrounded by
spaces](urls-and-fetching.html#valid-non-empty-url-potentially-surrounded-by-spaces){#the-object-element:valid-non-empty-url-potentially-surrounded-by-spaces}.

The [`type`]{#attr-object-type .dfn dfn-for="object"
dfn-type="element-attr"} attribute, if present, specifies the type of
the resource. If present, the attribute must be a [valid MIME type
string](https://mimesniff.spec.whatwg.org/#valid-mime-type){#the-object-element:valid-mime-type-string
x-internal="valid-mime-type-string"}.

The [`name`]{#attr-object-name .dfn dfn-for="object"
dfn-type="element-attr"} attribute, if present, must be a [valid
navigable target
name](document-sequences.html#valid-navigable-target-name){#the-object-element:valid-navigable-target-name}.
The given value is used to name the element\'s [content
navigable](document-sequences.html#content-navigable){#the-object-element:content-navigable-2},
if applicable, and if present when the element\'s [content
navigable](document-sequences.html#content-navigable){#the-object-element:content-navigable-3}
is
[created](document-sequences.html#create-a-new-child-navigable){#the-object-element:create-a-new-child-navigable}.

Whenever one of the following conditions occur:

- the element is created,
- the element is popped off the [stack of open
  elements](parsing.html#stack-of-open-elements){#the-object-element:stack-of-open-elements}
  of an [HTML
  parser](parsing.html#html-parser){#the-object-element:html-parser} or
  [XML parser](xhtml.html#xml-parser){#the-object-element:xml-parser},
- the element is not on the [stack of open
  elements](parsing.html#stack-of-open-elements){#the-object-element:stack-of-open-elements-2}
  of an [HTML
  parser](parsing.html#html-parser){#the-object-element:html-parser-2}
  or [XML
  parser](xhtml.html#xml-parser){#the-object-element:xml-parser-2}, and
  it is either [inserted into a
  document](infrastructure.html#insert-an-element-into-a-document){#the-object-element:insert-an-element-into-a-document}
  or [removed from a
  document](infrastructure.html#remove-an-element-from-a-document){#the-object-element:remove-an-element-from-a-document},
- the element\'s [node
  document](https://dom.spec.whatwg.org/#concept-node-document){#the-object-element:node-document
  x-internal="node-document"} changes whether it is [fully
  active](document-sequences.html#fully-active){#the-object-element:fully-active},
- one of the element\'s ancestor
  [`object`{#the-object-element:the-object-element-3}](#the-object-element)
  elements changes to or from showing its [fallback
  content](dom.html#fallback-content){#the-object-element:fallback-content},
- the element\'s
  [`classid`{#the-object-element:attr-object-classid}](obsolete.html#attr-object-classid)
  attribute is set, changed, or removed,
- the element\'s
  [`classid`{#the-object-element:attr-object-classid-2}](obsolete.html#attr-object-classid)
  attribute is not present, and its
  [`data`{#the-object-element:attr-object-data-2}](#attr-object-data)
  attribute is set, changed, or removed,
- neither the element\'s
  [`classid`{#the-object-element:attr-object-classid-3}](obsolete.html#attr-object-classid)
  attribute nor its
  [`data`{#the-object-element:attr-object-data-3}](#attr-object-data)
  attribute are present, and its
  [`type`{#the-object-element:attr-object-type-2}](#attr-object-type)
  attribute is set, changed, or removed,
- the element changes from [being
  rendered](rendering.html#being-rendered){#the-object-element:being-rendered}
  to not being rendered, or vice versa,

\...the user agent must [queue an element
task](webappapis.html#queue-an-element-task){#the-object-element:queue-an-element-task}
on the [DOM manipulation task
source](webappapis.html#dom-manipulation-task-source){#the-object-element:dom-manipulation-task-source}
given the
[`object`{#the-object-element:the-object-element-4}](#the-object-element)
element to run the following steps to (re)determine what the
[`object`{#the-object-element:the-object-element-5}](#the-object-element)
element represents. This
[task](webappapis.html#concept-task){#the-object-element:concept-task}
being
[queued](webappapis.html#queue-a-task){#the-object-element:queue-a-task}
or actively running must [delay the load
event](parsing.html#delay-the-load-event){#the-object-element:delay-the-load-event}
of the element\'s [node
document](https://dom.spec.whatwg.org/#concept-node-document){#the-object-element:node-document-2
x-internal="node-document"}.

1.  If the user has indicated a preference that this
    [`object`{#the-object-element:the-object-element-6}](#the-object-element)
    element\'s [fallback
    content](dom.html#fallback-content){#the-object-element:fallback-content-2}
    be shown instead of the element\'s usual behavior, then jump to the
    step below labeled *fallback*.

    For example, a user could ask for the element\'s [fallback
    content](dom.html#fallback-content){#the-object-element:fallback-content-3}
    to be shown because that content uses a format that the user finds
    more accessible.

2.  If the element has an ancestor [media
    element](media.html#media-element){#the-object-element:media-element},
    or has an ancestor
    [`object`{#the-object-element:the-object-element-7}](#the-object-element)
    element that is *not* showing its [fallback
    content](dom.html#fallback-content){#the-object-element:fallback-content-4},
    or if the element is not [in a
    document](https://dom.spec.whatwg.org/#in-a-document){#the-object-element:in-a-document
    x-internal="in-a-document"} whose [browsing
    context](document-sequences.html#concept-document-bc){#the-object-element:concept-document-bc}
    is non-null, or if the element\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-object-element:node-document-3
    x-internal="node-document"} is not [fully
    active](document-sequences.html#fully-active){#the-object-element:fully-active-2},
    or if the element is still in the [stack of open
    elements](parsing.html#stack-of-open-elements){#the-object-element:stack-of-open-elements-3}
    of an [HTML
    parser](parsing.html#html-parser){#the-object-element:html-parser-3}
    or [XML
    parser](xhtml.html#xml-parser){#the-object-element:xml-parser-3}, or
    if the element is not [being
    rendered](rendering.html#being-rendered){#the-object-element:being-rendered-2},
    then jump to the step below labeled *fallback*.

3.  If the
    [`data`{#the-object-element:attr-object-data-4}](#attr-object-data)
    attribute is present and its value is not the empty string, then:

    1.  If the
        [`type`{#the-object-element:attr-object-type-3}](#attr-object-type)
        attribute is present and its value is not a type that the user
        agent supports, then the user agent may jump to the step below
        labeled *fallback* without fetching the content to examine its
        real type.

    2.  Let `url`{.variable} be the result of [encoding-parsing a
        URL](urls-and-fetching.html#encoding-parsing-a-url){#the-object-element:encoding-parsing-a-url}
        given the
        [`data`{#the-object-element:attr-object-data-5}](#attr-object-data)
        attribute\'s value, relative to the element\'s [node
        document](https://dom.spec.whatwg.org/#concept-node-document){#the-object-element:node-document-4
        x-internal="node-document"}.

    3.  If `url`{.variable} is failure, then [fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#the-object-element:concept-event-fire
        x-internal="concept-event-fire"} named
        [`error`{#the-object-element:event-error}](indices.html#event-error)
        at the element and jump to the step below labeled *fallback*.

    4.  Let `request`{.variable} be a new
        [request](https://fetch.spec.whatwg.org/#concept-request){#the-object-element:concept-request
        x-internal="concept-request"} whose
        [URL](https://fetch.spec.whatwg.org/#concept-request-url){#the-object-element:concept-request-url
        x-internal="concept-request-url"} is `url`{.variable},
        [client](https://fetch.spec.whatwg.org/#concept-request-client){#the-object-element:concept-request-client
        x-internal="concept-request-client"} is the element\'s [node
        document](https://dom.spec.whatwg.org/#concept-node-document){#the-object-element:node-document-5
        x-internal="node-document"}\'s [relevant settings
        object](webappapis.html#relevant-settings-object){#the-object-element:relevant-settings-object},
        [destination](https://fetch.spec.whatwg.org/#concept-request-destination){#the-object-element:concept-request-destination
        x-internal="concept-request-destination"} is \"`object`\",
        [credentials
        mode](https://fetch.spec.whatwg.org/#concept-request-credentials-mode){#the-object-element:concept-request-credentials-mode
        x-internal="concept-request-credentials-mode"} is \"`include`\",
        [mode](https://fetch.spec.whatwg.org/#concept-request-mode){#the-object-element:concept-request-mode
        x-internal="concept-request-mode"} is \"`navigate`\", [initiator
        type](https://fetch.spec.whatwg.org/#request-initiator-type){#the-object-element:concept-request-initiator-type
        x-internal="concept-request-initiator-type"} is \"`object`\",
        and whose [use-URL-credentials
        flag](https://fetch.spec.whatwg.org/#concept-request-use-url-credentials-flag){#the-object-element:use-url-credentials-flag
        x-internal="use-url-credentials-flag"} is set.

    5.  [Fetch](https://fetch.spec.whatwg.org/#concept-fetch){#the-object-element:concept-fetch
        x-internal="concept-fetch"} `request`{.variable}.

        Fetching the resource must [delay the load
        event](parsing.html#delay-the-load-event){#the-object-element:delay-the-load-event-2}
        of the element\'s [node
        document](https://dom.spec.whatwg.org/#concept-node-document){#the-object-element:node-document-6
        x-internal="node-document"} until the
        [task](webappapis.html#concept-task){#the-object-element:concept-task-2}
        that is
        [queued](webappapis.html#queue-a-task){#the-object-element:queue-a-task-2}
        by the [networking task
        source](webappapis.html#networking-task-source){#the-object-element:networking-task-source}
        once the resource has been fetched (defined next) has been run.

    6.  If the resource is not yet available (e.g. because the resource
        was not available in the cache, so that loading the resource
        required making a request over the network), then jump to the
        step below labeled *fallback*. The
        [task](webappapis.html#concept-task){#the-object-element:concept-task-3}
        that is
        [queued](webappapis.html#queue-a-task){#the-object-element:queue-a-task-3}
        by the [networking task
        source](webappapis.html#networking-task-source){#the-object-element:networking-task-source-2}
        once the resource is available must restart this algorithm from
        this step. Resources can load incrementally; user agents may opt
        to consider a resource \"available\" whenever enough data has
        been obtained to begin processing the resource.

    7.  If the load failed (e.g. there was an HTTP 404 error, there was
        a DNS error), [fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#the-object-element:concept-event-fire-2
        x-internal="concept-event-fire"} named
        [`error`{#the-object-element:event-error-2}](indices.html#event-error)
        at the element, then jump to the step below labeled *fallback*.

    8.  ::: {#object-type-detection}
        Determine the `resource type`{.variable}, as follows:

        1.  Let the `resource type`{.variable} be unknown.

        2.  If the user agent is configured to strictly obey
            Content-Type headers for this resource, and the resource has
            [associated Content-Type
            metadata](urls-and-fetching.html#content-type){#the-object-element:content-type},
            then let the `resource type`{.variable} be the type
            specified in [the resource\'s Content-Type
            metadata](urls-and-fetching.html#content-type){#the-object-element:content-type-2},
            and jump to the step below labeled *handler*.

            This can introduce a vulnerability, wherein a site is trying
            to embed a resource that uses a particular type, but the
            remote site overrides that and instead furnishes the user
            agent with a resource that triggers a different type of
            content with different security characteristics.

        3.  Run the appropriate set of steps from the following list:

            If the resource has [associated Content-Type metadata](urls-and-fetching.html#content-type){#the-object-element:content-type-3}

            :   1.  Let `binary`{.variable} be false.

                2.  If the type specified in [the resource\'s
                    Content-Type
                    metadata](urls-and-fetching.html#content-type){#the-object-element:content-type-4}
                    is
                    \"[`text/plain`{#the-object-element:text/plain}](https://www.rfc-editor.org/rfc/rfc2046#section-4.1.3){x-internal="text/plain"}\",
                    and the result of applying the [rules for
                    distinguishing if a resource is text or
                    binary](https://mimesniff.spec.whatwg.org/#rules-for-text-or-binary){#the-object-element:content-type-sniffing:-text-or-binary
                    x-internal="content-type-sniffing:-text-or-binary"}
                    to the resource is that the resource is not
                    [`text/plain`{#the-object-element:text/plain-2}](https://www.rfc-editor.org/rfc/rfc2046#section-4.1.3){x-internal="text/plain"},
                    then set `binary`{.variable} to true.

                3.  If the type specified in [the resource\'s
                    Content-Type
                    metadata](urls-and-fetching.html#content-type){#the-object-element:content-type-5}
                    is
                    \"[`application/octet-stream`{#the-object-element:application/octet-stream}](https://www.rfc-editor.org/rfc/rfc2046#section-4.5.1){x-internal="application/octet-stream"}\",
                    then set `binary`{.variable} to true.

                4.  If `binary`{.variable} is false, then let the
                    `resource type`{.variable} be the type specified in
                    [the resource\'s Content-Type
                    metadata](urls-and-fetching.html#content-type){#the-object-element:content-type-6},
                    and jump to the step below labeled *handler*.

                5.  If there is a
                    [`type`{#the-object-element:attr-object-type-4}](#attr-object-type)
                    attribute present on the
                    [`object`{#the-object-element:the-object-element-8}](#the-object-element)
                    element, and its value is not
                    [`application/octet-stream`{#the-object-element:application/octet-stream-2}](https://www.rfc-editor.org/rfc/rfc2046#section-4.5.1){x-internal="application/octet-stream"},
                    then run the following steps:

                    1.  If the attribute\'s value is a type that starts
                        with \"`image/`\" that is not also an [XML MIME
                        type](https://mimesniff.spec.whatwg.org/#xml-mime-type){#the-object-element:xml-mime-type
                        x-internal="xml-mime-type"}, then let the
                        `resource type`{.variable} be the type specified
                        in that
                        [`type`{#the-object-element:attr-object-type-5}](#attr-object-type)
                        attribute.

                    2.  Jump to the step below labeled *handler*.

            Otherwise, if the resource does not have [associated Content-Type metadata](urls-and-fetching.html#content-type){#the-object-element:content-type-7}

            :   1.  If there is a
                    [`type`{#the-object-element:attr-object-type-6}](#attr-object-type)
                    attribute present on the
                    [`object`{#the-object-element:the-object-element-9}](#the-object-element)
                    element, then let the `tentative type`{.variable} be
                    the type specified in that
                    [`type`{#the-object-element:attr-object-type-7}](#attr-object-type)
                    attribute.

                    Otherwise, let `tentative type`{.variable} be the
                    [computed type of the
                    resource](https://mimesniff.spec.whatwg.org/#computed-mime-type){#the-object-element:content-type-sniffing-2
                    x-internal="content-type-sniffing-2"}.

                2.  If `tentative type`{.variable} is *not*
                    [`application/octet-stream`{#the-object-element:application/octet-stream-3}](https://www.rfc-editor.org/rfc/rfc2046#section-4.5.1){x-internal="application/octet-stream"},
                    then let `resource type`{.variable} be
                    `tentative type`{.variable} and jump to the step
                    below labeled *handler*.

        4.  If applying the [URL
            parser](https://url.spec.whatwg.org/#concept-url-parser){#the-object-element:url-parser
            x-internal="url-parser"} algorithm to the
            [URL](https://url.spec.whatwg.org/#concept-url){#the-object-element:url-2
            x-internal="url"} of the specified resource (after any
            redirects) results in a [URL
            record](https://url.spec.whatwg.org/#concept-url){#the-object-element:url-record
            x-internal="url-record"} whose
            [path](https://url.spec.whatwg.org/#concept-url-path){#the-object-element:concept-url-path
            x-internal="concept-url-path"} component matches a pattern
            that a
            [plugin](infrastructure.html#plugin){#the-object-element:plugin}
            supports, then let `resource type`{.variable} be the type
            that that plugin can handle.

            For example, a plugin might say that it can handle resources
            with
            [path](https://url.spec.whatwg.org/#concept-url-path){#the-object-element:concept-url-path-2
            x-internal="concept-url-path"} components that end with the
            four character string \"`.swf`\".

        It is possible for this step to finish, or for one of the
        substeps above to jump straight to the next step, with
        `resource type`{.variable} still being unknown. In both cases,
        the next step will trigger fallback.
        :::

    9.  *Handler*: Handle the content as given by the first of the
        following cases that matches:

        If the `resource type`{.variable} is an [XML MIME type](https://mimesniff.spec.whatwg.org/#xml-mime-type){#the-object-element:xml-mime-type-2 x-internal="xml-mime-type"}, or if the `resource type`{.variable} does not start with \"`image/`\"

        :   If the
            [`object`{#the-object-element:the-object-element-10}](#the-object-element)
            element\'s [content
            navigable](document-sequences.html#content-navigable){#the-object-element:content-navigable-4}
            is null, then [create a new child
            navigable](document-sequences.html#create-a-new-child-navigable){#the-object-element:create-a-new-child-navigable-2}
            for the element.

            Let `response`{.variable} be the
            [response](https://fetch.spec.whatwg.org/#concept-response){#the-object-element:concept-response
            x-internal="concept-response"} from
            [fetch](https://fetch.spec.whatwg.org/#concept-fetch){#the-object-element:concept-fetch-2
            x-internal="concept-fetch"}.

            If `response`{.variable}\'s
            [URL](https://fetch.spec.whatwg.org/#concept-response-url){#the-object-element:concept-response-url
            x-internal="concept-response-url"} does not [match
            `about:blank`](urls-and-fetching.html#matches-about:blank){#the-object-element:matches-about:blank},
            then
            [navigate](browsing-the-web.html#navigate){#the-object-element:navigate}
            the element\'s [content
            navigable](document-sequences.html#content-navigable){#the-object-element:content-navigable-5}
            to `response`{.variable}\'s
            [URL](https://fetch.spec.whatwg.org/#concept-response-url){#the-object-element:concept-response-url-2
            x-internal="concept-response-url"} using the element\'s
            [node
            document](https://dom.spec.whatwg.org/#concept-node-document){#the-object-element:node-document-7
            x-internal="node-document"}, with
            *[historyHandling](browsing-the-web.html#navigation-hh)* set
            to
            \"[`replace`{#the-object-element:navigationhistorybehavior-replace}](browsing-the-web.html#navigationhistorybehavior-replace)\".

            The
            [`data`{#the-object-element:attr-object-data-6}](#attr-object-data)
            attribute of the
            [`object`{#the-object-element:the-object-element-11}](#the-object-element)
            element doesn\'t get updated if the [content
            navigable](document-sequences.html#content-navigable){#the-object-element:content-navigable-6}
            gets further
            [navigated](browsing-the-web.html#navigate){#the-object-element:navigate-2}
            to other locations.

            The
            [`object`{#the-object-element:the-object-element-12}](#the-object-element)
            element
            [represents](dom.html#represents){#the-object-element:represents}
            its [content
            navigable](document-sequences.html#content-navigable){#the-object-element:content-navigable-7}.

        If the `resource type`{.variable} starts with \"`image/`\", and support for images has not been disabled

        :   [Destroy a child
            navigable](document-sequences.html#destroy-a-child-navigable){#the-object-element:destroy-a-child-navigable}
            given the
            [`object`{#the-object-element:the-object-element-13}](#the-object-element)
            element.

            Apply the [image
            sniffing](https://mimesniff.spec.whatwg.org/#rules-for-sniffing-images-specifically){#the-object-element:content-type-sniffing:-image
            x-internal="content-type-sniffing:-image"} rules to
            determine the type of the image.

            The
            [`object`{#the-object-element:the-object-element-14}](#the-object-element)
            element
            [represents](dom.html#represents){#the-object-element:represents-2}
            the specified image.

            If the image cannot be rendered, e.g. because it is
            malformed or in an unsupported format, jump to the step
            below labeled *fallback*.

        Otherwise

        :   The given `resource type`{.variable} is not supported. Jump
            to the step below labeled *fallback*.

            If the previous step ended with the
            `resource type`{.variable} being unknown, this is the case
            that is triggered.

    10. The element\'s contents are not part of what the
        [`object`{#the-object-element:the-object-element-15}](#the-object-element)
        element represents.

    11. If the
        [`object`{#the-object-element:the-object-element-16}](#the-object-element)
        element does not represent its [content
        navigable](document-sequences.html#content-navigable){#the-object-element:content-navigable-8},
        then once the resource is completely loaded, [queue an element
        task](webappapis.html#queue-an-element-task){#the-object-element:queue-an-element-task-2}
        on the [DOM manipulation task
        source](webappapis.html#dom-manipulation-task-source){#the-object-element:dom-manipulation-task-source-2}
        given the
        [`object`{#the-object-element:the-object-element-17}](#the-object-element)
        element to [fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#the-object-element:concept-event-fire-3
        x-internal="concept-event-fire"} named
        [`load`{#the-object-element:event-load}](indices.html#event-load)
        at the element.

        If the element *does* represent its [content
        navigable](document-sequences.html#content-navigable){#the-object-element:content-navigable-9},
        then an analogous task will be queued when the created
        [`Document`{#the-object-element:document-3}](dom.html#document)
        is [completely finished
        loading](document-lifecycle.html#completely-finish-loading){#the-object-element:completely-finish-loading}.

    12. Return.

4.  *Fallback*: The
    [`object`{#the-object-element:the-object-element-18}](#the-object-element)
    element
    [represents](dom.html#represents){#the-object-element:represents-3}
    the element\'s children. This is the element\'s [fallback
    content](dom.html#fallback-content){#the-object-element:fallback-content-5}.
    [Destroy a child
    navigable](document-sequences.html#destroy-a-child-navigable){#the-object-element:destroy-a-child-navigable-2}
    given the element.

Due to the algorithm above, the contents of
[`object`{#the-object-element:the-object-element-19}](#the-object-element)
elements act as [fallback
content](dom.html#fallback-content){#the-object-element:fallback-content-6},
used only when referenced resources can\'t be shown (e.g. because it
returned a 404 error). This allows multiple
[`object`{#the-object-element:the-object-element-20}](#the-object-element)
elements to be nested inside each other, targeting multiple user agents
with different capabilities, with the user agent picking the first one
it supports.

The
[`object`{#the-object-element:the-object-element-21}](#the-object-element)
element [potentially delays the load
event](#potentially-delays-the-load-event){#the-object-element:potentially-delays-the-load-event}.

The
[`form`{#the-object-element:attr-fae-form-2}](form-control-infrastructure.html#attr-fae-form)
attribute is used to explicitly associate the
[`object`{#the-object-element:the-object-element-22}](#the-object-element)
element with its [form
owner](form-control-infrastructure.html#form-owner){#the-object-element:form-owner}.

The
[`object`{#the-object-element:the-object-element-23}](#the-object-element)
element supports [dimension
attributes](embedded-content-other.html#dimension-attributes){#the-object-element:dimension-attributes}.

::::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLObjectElement/data](https://developer.mozilla.org/en-US/docs/Web/API/HTMLObjectElement/data "The data property of the HTMLObjectElement interface returns a string that reflects the data HTML attribute, specifying the address of a resource's data.")

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

:::: feature
[HTMLObjectElement/type](https://developer.mozilla.org/en-US/docs/Web/API/HTMLObjectElement/type "The type property of the HTMLObjectElement interface returns a string that reflects the type HTML attribute, specifying the MIME type of the resource.")

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

:::: feature
[HTMLObjectElement/name](https://developer.mozilla.org/en-US/docs/Web/API/HTMLObjectElement/name "The name property of the HTMLObjectElement interface returns a string that reflects the name HTML attribute, specifying the name of the browsing context.")

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
:::::::::

The IDL attributes [`data`]{#dom-object-data .dfn
dfn-for="HTMLObjectElement" dfn-type="attribute"},
[`type`]{#dom-object-type .dfn dfn-for="HTMLObjectElement"
dfn-type="attribute"}, and [`name`]{#dom-object-name .dfn
dfn-for="HTMLObjectElement" dfn-type="attribute"} each must
[reflect](common-dom-interfaces.html#reflect){#the-object-element:reflect}
the respective content attributes of the same name.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLObjectElement/contentDocument](https://developer.mozilla.org/en-US/docs/Web/API/HTMLObjectElement/contentDocument "The contentDocument read-only property of the HTMLObjectElement interface Returns a Document representing the active document of the object element's nested browsing context, if any; otherwise null.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer8+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The [`contentDocument`]{#dom-object-contentdocument .dfn
dfn-for="HTMLObjectElement" dfn-type="attribute"} getter steps are to
return
[this](https://webidl.spec.whatwg.org/#this){#the-object-element:this
x-internal="this"}\'s [content
document](document-sequences.html#concept-bcc-content-document){#the-object-element:concept-bcc-content-document}.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLObjectElement/contentWindow](https://developer.mozilla.org/en-US/docs/Web/API/HTMLObjectElement/contentWindow "The contentWindow read-only property of the HTMLObjectElement interface returns a WindowProxy representing the window proxy of the object element's nested browsing context, if any; otherwise null.")

Support in all current engines.

::: support
[Firefox22+]{.firefox .yes}[Safari13+]{.safari .yes}[Chrome53+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)17+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

The [`contentWindow`]{#dom-object-contentwindow .dfn
dfn-for="HTMLObjectElement" dfn-type="attribute"} getter steps are to
return
[this](https://webidl.spec.whatwg.org/#this){#the-object-element:this-2
x-internal="this"}\'s [content
window](document-sequences.html#content-window){#the-object-element:content-window}.

The
[`willValidate`{#the-object-element:dom-cva-willvalidate-2}](form-control-infrastructure.html#dom-cva-willvalidate),
[`validity`{#the-object-element:dom-cva-validity-2}](form-control-infrastructure.html#dom-cva-validity),
and
[`validationMessage`{#the-object-element:dom-cva-validationmessage-2}](form-control-infrastructure.html#dom-cva-validationmessage)
attributes, and the
[`checkValidity()`{#the-object-element:dom-cva-checkvalidity-2}](form-control-infrastructure.html#dom-cva-checkvalidity),
[`reportValidity()`{#the-object-element:dom-cva-reportvalidity-2}](form-control-infrastructure.html#dom-cva-reportvalidity),
and
[`setCustomValidity()`{#the-object-element:dom-cva-setcustomvalidity-2}](form-control-infrastructure.html#dom-cva-setcustomvalidity)
methods, are part of the [constraint validation
API](form-control-infrastructure.html#the-constraint-validation-api){#the-object-element:the-constraint-validation-api}.
The
[`form`{#the-object-element:dom-fae-form-2}](form-control-infrastructure.html#dom-fae-form)
IDL attribute is part of the element\'s forms API.

::: example
In this example, an HTML page is embedded in another using the
[`object`{#the-object-element:the-object-element-24}](#the-object-element)
element.

``` html
<figure>
 <object data="clock.html"></object>
 <figcaption>My HTML Clock</figcaption>
</figure>
```
:::

[← 4.8.4 Images](images.html) --- [Table of Contents](./) --- [4.8.8 The
video element →](media.html)
