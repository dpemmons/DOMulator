HTML Standard

::: {#head .head .with-buttons}
[![WHATWG](https://resources.whatwg.org/logo.svg){.darkmode-aware
width="100" crossorigin="" height="100"}](https://whatwg.org/){.logo}

# HTML {#html .allcaps}

Living Standard --- Last Updated [2 June 2025]{.pubdate}
:::

[← 4 The elements of HTML](semantics.html) --- [Table of Contents](./)
--- [4.4 Grouping content →](grouping-content.html)

1.  ::: {#toc-semantics}
    1.  [[4.3]{.secno} Sections](sections.html#sections)
        1.  [[4.3.1]{.secno} The `body`
            element](sections.html#the-body-element)
        2.  [[4.3.2]{.secno} The `article`
            element](sections.html#the-article-element)
        3.  [[4.3.3]{.secno} The `section`
            element](sections.html#the-section-element)
        4.  [[4.3.4]{.secno} The `nav`
            element](sections.html#the-nav-element)
        5.  [[4.3.5]{.secno} The `aside`
            element](sections.html#the-aside-element)
        6.  [[4.3.6]{.secno} The `h1`, `h2`, `h3`, `h4`, `h5`, and `h6`
            elements](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements)
        7.  [[4.3.7]{.secno} The `hgroup`
            element](sections.html#the-hgroup-element)
        8.  [[4.3.8]{.secno} The `header`
            element](sections.html#the-header-element)
        9.  [[4.3.9]{.secno} The `footer`
            element](sections.html#the-footer-element)
        10. [[4.3.10]{.secno} The `address`
            element](sections.html#the-address-element)
        11. [[4.3.11]{.secno} Headings and
            outlines](sections.html#headings-and-outlines-2)
            1.  [[4.3.11.1]{.secno} Sample
                outlines](sections.html#sample-outlines)
            2.  [[4.3.11.2]{.secno} Exposing outlines to
                users](sections.html#exposing-outlines-to-users)
        12. [[4.3.12]{.secno} Usage
            summary](sections.html#usage-summary-2)
            1.  [[4.3.12.1]{.secno} Article or
                section?](sections.html#article-or-section)
    :::

### [4.3]{.secno} Sections[](#sections){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Introduction_to_HTML/Document_and_website_structure#HTML_for_structuring_content](https://developer.mozilla.org/en-US/docs/Web/HTML/Introduction_to_HTML/Document_and_website_structure#HTML_for_structuring_content "At this point, you should have a better idea about how to structure a web page/site. In the last article of this module, we'll learn how to debug HTML.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome5+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11.1+]{.opera_android .yes}
:::
::::
:::::

#### [4.3.1]{.secno} The [`body`]{.dfn dfn-type="element"} element[](#the-body-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/body](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/body "The <body> HTML element represents the content of an HTML document. There can be only one <body> element in a document.")

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

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLBodyElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLBodyElement "The HTMLBodyElement interface provides special properties (beyond those inherited from the regular HTMLElement interface) for manipulating <body> elements.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
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

[Categories](dom.html#concept-element-categories){#the-body-element:concept-element-categories}:
:   None.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-body-element:concept-element-contexts}:
:   As the second element in an
    [`html`{#the-body-element:the-html-element}](semantics.html#the-html-element)
    element.

[Content model](dom.html#concept-element-content-model){#the-body-element:concept-element-content-model}:
:   [Flow
    content](dom.html#flow-content-2){#the-body-element:flow-content-2}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-body-element:concept-element-tag-omission}:
:   A [`body`{#the-body-element:the-body-element}](#the-body-element)
    element\'s [start
    tag](syntax.html#syntax-start-tag){#the-body-element:syntax-start-tag}
    can be omitted if the element is empty, or if the first thing inside
    the
    [`body`{#the-body-element:the-body-element-2}](#the-body-element)
    element is not [ASCII
    whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#the-body-element:space-characters
    x-internal="space-characters"} or a
    [comment](syntax.html#syntax-comments){#the-body-element:syntax-comments},
    except if the first thing inside the
    [`body`{#the-body-element:the-body-element-3}](#the-body-element)
    element is a
    [`meta`{#the-body-element:the-meta-element}](semantics.html#the-meta-element),
    [`noscript`{#the-body-element:the-noscript-element}](scripting.html#the-noscript-element),
    [`link`{#the-body-element:the-link-element}](semantics.html#the-link-element),
    [`script`{#the-body-element:the-script-element}](scripting.html#the-script-element),
    [`style`{#the-body-element:the-style-element}](semantics.html#the-style-element),
    or
    [`template`{#the-body-element:the-template-element}](scripting.html#the-template-element)
    element.
:   A [`body`{#the-body-element:the-body-element-4}](#the-body-element)
    element\'s [end
    tag](syntax.html#syntax-end-tag){#the-body-element:syntax-end-tag}
    can be omitted if the
    [`body`{#the-body-element:the-body-element-5}](#the-body-element)
    element is not immediately followed by a
    [comment](syntax.html#syntax-comments){#the-body-element:syntax-comments-2}.

[Content attributes](dom.html#concept-element-attributes){#the-body-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-body-element:global-attributes}
:   [`onafterprint`{#the-body-element:handler-window-onafterprint}](webappapis.html#handler-window-onafterprint)
:   [`onbeforeprint`{#the-body-element:handler-window-onbeforeprint}](webappapis.html#handler-window-onbeforeprint)
:   [`onbeforeunload`{#the-body-element:handler-window-onbeforeunload}](webappapis.html#handler-window-onbeforeunload)
:   [`onhashchange`{#the-body-element:handler-window-onhashchange}](webappapis.html#handler-window-onhashchange)
:   [`onlanguagechange`{#the-body-element:handler-window-onlanguagechange}](webappapis.html#handler-window-onlanguagechange)
:   [`onmessage`{#the-body-element:handler-window-onmessage}](webappapis.html#handler-window-onmessage)
:   [`onmessageerror`{#the-body-element:handler-window-onmessageerror}](webappapis.html#handler-window-onmessageerror)
:   [`onoffline`{#the-body-element:handler-window-onoffline}](webappapis.html#handler-window-onoffline)
:   [`ononline`{#the-body-element:handler-window-ononline}](webappapis.html#handler-window-ononline)
:   [`onpageswap`{#the-body-element:handler-window-onpageswap}](webappapis.html#handler-window-onpageswap)
:   [`onpagehide`{#the-body-element:handler-window-onpagehide}](webappapis.html#handler-window-onpagehide)
:   [`onpagereveal`{#the-body-element:handler-window-onpagereveal}](webappapis.html#handler-window-onpagereveal)
:   [`onpageshow`{#the-body-element:handler-window-onpageshow}](webappapis.html#handler-window-onpageshow)
:   [`onpopstate`{#the-body-element:handler-window-onpopstate}](webappapis.html#handler-window-onpopstate)
:   [`onrejectionhandled`{#the-body-element:handler-window-onrejectionhandled}](webappapis.html#handler-window-onrejectionhandled)
:   [`onstorage`{#the-body-element:handler-window-onstorage}](webappapis.html#handler-window-onstorage)
:   [`onunhandledrejection`{#the-body-element:handler-window-onunhandledrejection}](webappapis.html#handler-window-onunhandledrejection)
:   [`onunload`{#the-body-element:handler-window-onunload}](webappapis.html#handler-window-onunload)

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-body-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-body).
:   [For implementers](https://w3c.github.io/html-aam/#el-body).

[DOM interface](dom.html#concept-element-dom){#the-body-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLBodyElement : HTMLElement {
      [HTMLConstructor] constructor();

      // also has obsolete members
    };

    HTMLBodyElement includes WindowEventHandlers;
    ```

The [`body`{#the-body-element:the-body-element-6}](#the-body-element)
element [represents](dom.html#represents){#the-body-element:represents}
the contents of the document.

In conforming documents, there is only one
[`body`{#the-body-element:the-body-element-7}](#the-body-element)
element. The
[`document.body`{#the-body-element:dom-document-body}](dom.html#dom-document-body)
IDL attribute provides scripts with easy access to a document\'s
[`body`{#the-body-element:the-body-element-8}](#the-body-element)
element.

Some DOM operations (for example, parts of the [drag and
drop](dnd.html#dnd){#the-body-element:dnd} model) are defined in terms
of \"[the body
element](dom.html#the-body-element-2){#the-body-element:the-body-element-2-2}\".
This refers to a particular element in the DOM, as per the definition of
the term, and not any arbitrary
[`body`{#the-body-element:the-body-element-9}](#the-body-element)
element.

The [`body`{#the-body-element:the-body-element-10}](#the-body-element)
element exposes as [event handler content
attributes](webappapis.html#event-handler-content-attributes){#the-body-element:event-handler-content-attributes}
a number of the [event
handlers](webappapis.html#event-handlers){#the-body-element:event-handlers}
of the
[`Window`{#the-body-element:window}](nav-history-apis.html#window)
object. It also mirrors their [event handler IDL
attributes](webappapis.html#event-handler-idl-attributes){#the-body-element:event-handler-idl-attributes}.

The [event
handlers](webappapis.html#event-handlers){#the-body-element:event-handlers-2}
of the
[`Window`{#the-body-element:window-2}](nav-history-apis.html#window)
object named by the [`Window`-reflecting body element event handler
set](webappapis.html#window-reflecting-body-element-event-handler-set){#the-body-element:window-reflecting-body-element-event-handler-set},
exposed on the
[`body`{#the-body-element:the-body-element-11}](#the-body-element)
element, replace the generic [event
handlers](webappapis.html#event-handlers){#the-body-element:event-handlers-3}
with the same names normally supported by [HTML
elements](infrastructure.html#html-elements){#the-body-element:html-elements}.

Thus, for example, a bubbling
[`error`{#the-body-element:event-error}](indices.html#event-error) event
dispatched on a child of [the body
element](dom.html#the-body-element-2){#the-body-element:the-body-element-2-3}
of a [`Document`{#the-body-element:document}](dom.html#document) would
first trigger the
[`onerror`{#the-body-element:handler-onerror}](webappapis.html#handler-onerror)
[event handler content
attributes](webappapis.html#event-handler-content-attributes){#the-body-element:event-handler-content-attributes-2}
of that element, then that of the root
[`html`{#the-body-element:the-html-element-2}](semantics.html#the-html-element)
element, and only *then* would it trigger the
[`onerror`{#the-body-element:handler-onerror-2}](webappapis.html#handler-onerror)
[event handler content
attribute](webappapis.html#event-handler-content-attributes){#the-body-element:event-handler-content-attributes-3}
on the
[`body`{#the-body-element:the-body-element-12}](#the-body-element)
element. This is because the event would bubble from the target, to the
[`body`{#the-body-element:the-body-element-13}](#the-body-element), to
the
[`html`{#the-body-element:the-html-element-3}](semantics.html#the-html-element),
to the [`Document`{#the-body-element:document-2}](dom.html#document), to
the
[`Window`{#the-body-element:window-3}](nav-history-apis.html#window),
and the [event
handler](webappapis.html#event-handlers){#the-body-element:event-handlers-4}
on the
[`body`{#the-body-element:the-body-element-14}](#the-body-element) is
watching the
[`Window`{#the-body-element:window-4}](nav-history-apis.html#window) not
the [`body`{#the-body-element:the-body-element-15}](#the-body-element).
A regular event listener attached to the
[`body`{#the-body-element:the-body-element-16}](#the-body-element) using
`addEventListener()`, however, would be run when the event bubbled
through the
[`body`{#the-body-element:the-body-element-17}](#the-body-element) and
not when it reaches the
[`Window`{#the-body-element:window-5}](nav-history-apis.html#window)
object.

::: example
This page updates an indicator to show whether or not the user is
online:

``` html
<!DOCTYPE HTML>
<html lang="en">
 <head>
  <title>Online or offline?</title>
  <script>
   function update(online) {
     document.getElementById('status').textContent =
       online ? 'Online' : 'Offline';
   }
  </script>
 </head>
 <body ononline="update(true)"
       onoffline="update(false)"
       onload="update(navigator.onLine)">
  <p>You are: <span id="status">(Unknown)</span></p>
 </body>
</html>
```
:::

#### [4.3.2]{.secno} The [`article`]{.dfn dfn-type="element"} element[](#the-article-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/article](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/article "The <article> HTML element represents a self-contained composition in a document, page, application, or site, which is intended to be independently distributable or reusable (e.g., in syndication). Examples include: a forum post, a magazine or newspaper article, or a blog entry, a product card, a user-submitted comment, an interactive widget or gadget, or any other independent item of content.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome5+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11.1+]{.opera_android .yes}
:::
::::
:::::

[Categories](dom.html#concept-element-categories){#the-article-element:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-article-element:flow-content-2}.
:   [Sectioning
    content](dom.html#sectioning-content-2){#the-article-element:sectioning-content-2}.
:   [Palpable
    content](dom.html#palpable-content-2){#the-article-element:palpable-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-article-element:concept-element-contexts}:
:   Where [sectioning
    content](dom.html#sectioning-content-2){#the-article-element:sectioning-content-2-2}
    is expected.

[Content model](dom.html#concept-element-content-model){#the-article-element:concept-element-content-model}:
:   [Flow
    content](dom.html#flow-content-2){#the-article-element:flow-content-2-2}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-article-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-article-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-article-element:global-attributes}

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-article-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-article).
:   [For implementers](https://w3c.github.io/html-aam/#el-article).

[DOM interface](dom.html#concept-element-dom){#the-article-element:concept-element-dom}:
:   Uses
    [`HTMLElement`{#the-article-element:htmlelement}](dom.html#htmlelement).

The
[`article`{#the-article-element:the-article-element}](#the-article-element)
element
[represents](dom.html#represents){#the-article-element:represents} a
complete, or self-contained, composition in a document, page,
application, or site and that is, in principle, independently
distributable or reusable, e.g. in syndication. This could be a forum
post, a magazine or newspaper article, a blog entry, a user-submitted
comment, an interactive widget or gadget, or any other independent item
of content.

When
[`article`{#the-article-element:the-article-element-2}](#the-article-element)
elements are nested, the inner
[`article`{#the-article-element:the-article-element-3}](#the-article-element)
elements represent articles that are in principle related to the
contents of the outer article. For instance, a blog entry on a site that
accepts user-submitted comments could represent the comments as
[`article`{#the-article-element:the-article-element-4}](#the-article-element)
elements nested within the
[`article`{#the-article-element:the-article-element-5}](#the-article-element)
element for the blog entry.

Author information associated with an
[`article`{#the-article-element:the-article-element-6}](#the-article-element)
element (q.v. the
[`address`{#the-article-element:the-address-element}](#the-address-element)
element) does not apply to nested
[`article`{#the-article-element:the-article-element-7}](#the-article-element)
elements.

When used specifically with content to be redistributed in syndication,
the
[`article`{#the-article-element:the-article-element-8}](#the-article-element)
element is similar in purpose to the `entry` element in Atom.
[\[ATOM\]](references.html#refsATOM)

The schema.org microdata vocabulary can be used to provide the
publication date for an
[`article`{#the-article-element:the-article-element-9}](#the-article-element)
element, using one of the CreativeWork subtypes.

When the main content of the page (i.e. excluding footers, headers,
navigation blocks, and sidebars) is all one single self-contained
composition, that content may be marked with an
[`article`{#the-article-element:the-article-element-10}](#the-article-element),
but it is technically redundant in that case (since it\'s self-evident
that the page is a single composition, as it is a single document).

::: {#article-example .example}
This example shows a blog post using the
[`article`{#the-article-element:the-article-element-11}](#the-article-element)
element, with some schema.org annotations:

``` html
<article itemscope itemtype="http://schema.org/BlogPosting">
 <header>
  <h2 itemprop="headline">The Very First Rule of Life</h2>
  <p><time itemprop="datePublished" datetime="2009-10-09">3 days ago</time></p>
  <link itemprop="url" href="?comments=0">
 </header>
 <p>If there's a microphone anywhere near you, assume it's hot and
 sending whatever you're saying to the world. Seriously.</p>
 <p>...</p>
 <footer>
  <a itemprop="discussionUrl" href="?comments=1">Show comments...</a>
 </footer>
</article>
```

Here is that same blog post, but showing some of the comments:

``` html
<article itemscope itemtype="http://schema.org/BlogPosting">
 <header>
  <h2 itemprop="headline">The Very First Rule of Life</h2>
  <p><time itemprop="datePublished" datetime="2009-10-09">3 days ago</time></p>
  <link itemprop="url" href="?comments=0">
 </header>
 <p>If there's a microphone anywhere near you, assume it's hot and
 sending whatever you're saying to the world. Seriously.</p>
 <p>...</p>
 <section>
  <h1>Comments</h1>
  <article itemprop="comment" itemscope itemtype="http://schema.org/Comment" id="c1">
   <link itemprop="url" href="#c1">
   <footer>
    <p>Posted by: <span itemprop="creator" itemscope itemtype="http://schema.org/Person">
     <span itemprop="name">George Washington</span>
    </span></p>
    <p><time itemprop="dateCreated" datetime="2009-10-10">15 minutes ago</time></p>
   </footer>
   <p>Yeah! Especially when talking about your lobbyist friends!</p>
  </article>
  <article itemprop="comment" itemscope itemtype="http://schema.org/Comment" id="c2">
   <link itemprop="url" href="#c2">
   <footer>
    <p>Posted by: <span itemprop="creator" itemscope itemtype="http://schema.org/Person">
     <span itemprop="name">George Hammond</span>
    </span></p>
    <p><time itemprop="dateCreated" datetime="2009-10-10">5 minutes ago</time></p>
   </footer>
   <p>Hey, you have the same first name as me.</p>
  </article>
 </section>
</article>
```

Notice the use of
[`footer`{#the-article-element:the-footer-element}](#the-footer-element)
to give the information for each comment (such as who wrote it and
when): the
[`footer`{#the-article-element:the-footer-element-2}](#the-footer-element)
element *can* appear at the start of its section when appropriate, such
as in this case. (Using
[`header`{#the-article-element:the-header-element}](#the-header-element)
in this case wouldn\'t be wrong either; it\'s mostly a matter of
authoring preference.)
:::

::: example
In this example,
[`article`{#the-article-element:the-article-element-12}](#the-article-element)
elements are used to host widgets on a portal page. The widgets are
implemented as [customized built-in
elements](custom-elements.html#customized-built-in-element){#the-article-element:customized-built-in-element}
in order to get specific styling and scripted behavior.

``` html
<!DOCTYPE HTML>
<html lang=en>
<title>eHome Portal</title>
<script src="/scripts/widgets.js"></script>
<link rel=stylesheet href="/styles/main.css">
<article is="stock-widget">
 <h2>Stocks</h2>
 <table>
  <thead> <tr> <th> Stock <th> Value <th> Delta
  <tbody> <template> <tr> <td> <td> <td> </template>
 </table>
 <p> <input type=button value="Refresh" onclick="this.parentElement.refresh()">
</article>
<article is="news-widget">
 <h2>News</h2>
 <ul>
  <template>
   <li>
    <p><img> <strong></strong>
    <p>
  </template>
 </ul>
 <p> <input type=button value="Refresh" onclick="this.parentElement.refresh()">
</article>
```
:::

#### [4.3.3]{.secno} The [`section`]{.dfn dfn-type="element"} element[](#the-section-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/section](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/section "The <section> HTML element represents a generic standalone section of a document, which doesn't have a more specific semantic element to represent it. Sections should always have a heading, with very few exceptions.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome5+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11.1+]{.opera_android .yes}
:::
::::
:::::

[Categories](dom.html#concept-element-categories){#the-section-element:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-section-element:flow-content-2}.
:   [Sectioning
    content](dom.html#sectioning-content-2){#the-section-element:sectioning-content-2}.
:   [Palpable
    content](dom.html#palpable-content-2){#the-section-element:palpable-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-section-element:concept-element-contexts}:
:   Where [sectioning
    content](dom.html#sectioning-content-2){#the-section-element:sectioning-content-2-2}
    is expected.

[Content model](dom.html#concept-element-content-model){#the-section-element:concept-element-content-model}:
:   [Flow
    content](dom.html#flow-content-2){#the-section-element:flow-content-2-2}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-section-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-section-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-section-element:global-attributes}

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-section-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-section).
:   [For implementers](https://w3c.github.io/html-aam/#el-section).

[DOM interface](dom.html#concept-element-dom){#the-section-element:concept-element-dom}:
:   Uses
    [`HTMLElement`{#the-section-element:htmlelement}](dom.html#htmlelement).

The
[`section`{#the-section-element:the-section-element}](#the-section-element)
element
[represents](dom.html#represents){#the-section-element:represents} a
generic section of a document or application. A section, in this
context, is a thematic grouping of content, typically with a heading.

Examples of sections would be chapters, the various tabbed pages in a
tabbed dialog box, or the numbered sections of a thesis. A web site\'s
home page could be split into sections for an introduction, news items,
and contact information.

Authors are encouraged to use the
[`article`{#the-section-element:the-article-element}](#the-article-element)
element instead of the
[`section`{#the-section-element:the-section-element-2}](#the-section-element)
element when it would make sense to syndicate the contents of the
element.

The
[`section`{#the-section-element:the-section-element-3}](#the-section-element)
element is not a generic container element. When an element is needed
only for styling purposes or as a convenience for scripting, authors are
encouraged to use the
[`div`{#the-section-element:the-div-element}](grouping-content.html#the-div-element)
element instead. A general rule is that the
[`section`{#the-section-element:the-section-element-4}](#the-section-element)
element is appropriate only if the element\'s contents would be listed
explicitly in the document\'s
[outline](#outline){#the-section-element:outline}.

::: example
In the following example, we see an article (part of a larger web page)
about apples, containing two short sections.

``` html
<article>
 <hgroup>
  <h2>Apples</h2>
  <p>Tasty, delicious fruit!</p>
 </hgroup>
 <p>The apple is the pomaceous fruit of the apple tree.</p>
 <section>
  <h3>Red Delicious</h3>
  <p>These bright red apples are the most common found in many
  supermarkets.</p>
 </section>
 <section>
  <h3>Granny Smith</h3>
  <p>These juicy, green apples make a great filling for
  apple pies.</p>
 </section>
</article>
```
:::

::: example
Here is a graduation programme with two sections, one for the list of
people graduating, and one for the description of the ceremony. (The
markup in this example features an uncommon style sometimes used to
minimize the amount of [inter-element
whitespace](dom.html#inter-element-whitespace){#the-section-element:inter-element-whitespace}.)

``` html
<!DOCTYPE Html>
<Html Lang=En
 ><Head
   ><Title
     >Graduation Ceremony Summer 2022</Title
   ></Head
 ><Body
   ><H1
     >Graduation</H1
   ><Section
     ><H2
       >Ceremony</H2
     ><P
       >Opening Procession</P
     ><P
       >Speech by Valedictorian</P
     ><P
       >Speech by Class President</P
     ><P
       >Presentation of Diplomas</P
     ><P
       >Closing Speech by Headmaster</P
   ></Section
   ><Section
     ><H2
       >Graduates</H2
     ><Ul
       ><Li
         >Molly Carpenter</Li
       ><Li
         >Anastasia Luccio</Li
       ><Li
         >Ebenezar McCoy</Li
       ><Li
         >Karrin Murphy</Li
       ><Li
         >Thomas Raith</Li
       ><Li
         >Susan Rodriguez</Li
     ></Ul
   ></Section
 ></Body
></Html>
```
:::

::: example
In this example, a book author has marked up some sections as chapters
and some as appendices, and uses CSS to style the headers in these two
classes of section differently.

``` html
<style>
 section { border: double medium; margin: 2em; }
 section.chapter h2 { font: 2em Roboto, Helvetica Neue, sans-serif; }
 section.appendix h2 { font: small-caps 2em Roboto, Helvetica Neue, sans-serif; }
</style>
<header>
 <hgroup>
  <h1>My Book</h1>
  <p>A sample with not much content</p>
 </hgroup>
 <p><small>Published by Dummy Publicorp Ltd.</small></p>
</header>
<section class="chapter">
 <h2>My First Chapter</h2>
 <p>This is the first of my chapters. It doesn't say much.</p>
 <p>But it has two paragraphs!</p>
</section>
<section class="chapter">
 <h2>It Continues: The Second Chapter</h2>
 <p>Bla dee bla, dee bla dee bla. Boom.</p>
</section>
<section class="chapter">
 <h2>Chapter Three: A Further Example</h2>
 <p>It's not like a battle between brightness and earthtones would go
 unnoticed.</p>
 <p>But it might ruin my story.</p>
</section>
<section class="appendix">
 <h2>Appendix A: Overview of Examples</h2>
 <p>These are demonstrations.</p>
</section>
<section class="appendix">
 <h2>Appendix B: Some Closing Remarks</h2>
 <p>Hopefully this long example shows that you <em>can</em> style
 sections, so long as they are used to indicate actual sections.</p>
</section>
```
:::

#### [4.3.4]{.secno} The [`nav`]{.dfn dfn-type="element"} element[](#the-nav-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/nav](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/nav "The <nav> HTML element represents a section of a page whose purpose is to provide navigation links, either within the current document or to other documents. Common examples of navigation sections are menus, tables of contents, and indexes.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome5+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11.1+]{.opera_android .yes}
:::
::::
:::::

[Categories](dom.html#concept-element-categories){#the-nav-element:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-nav-element:flow-content-2}.
:   [Sectioning
    content](dom.html#sectioning-content-2){#the-nav-element:sectioning-content-2}.
:   [Palpable
    content](dom.html#palpable-content-2){#the-nav-element:palpable-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-nav-element:concept-element-contexts}:
:   Where [sectioning
    content](dom.html#sectioning-content-2){#the-nav-element:sectioning-content-2-2}
    is expected.

[Content model](dom.html#concept-element-content-model){#the-nav-element:concept-element-content-model}:
:   [Flow
    content](dom.html#flow-content-2){#the-nav-element:flow-content-2-2}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-nav-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-nav-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-nav-element:global-attributes}

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-nav-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-nav).
:   [For implementers](https://w3c.github.io/html-aam/#el-nav).

[DOM interface](dom.html#concept-element-dom){#the-nav-element:concept-element-dom}:
:   Uses
    [`HTMLElement`{#the-nav-element:htmlelement}](dom.html#htmlelement).

The [`nav`{#the-nav-element:the-nav-element}](#the-nav-element) element
[represents](dom.html#represents){#the-nav-element:represents} a section
of a page that links to other pages or to parts within the page: a
section with navigation links.

Not all groups of links on a page need to be in a
[`nav`{#the-nav-element:the-nav-element-2}](#the-nav-element) element
--- the element is primarily intended for sections that consist of major
navigation blocks. In particular, it is common for footers to have a
short list of links to various pages of a site, such as the terms of
service, the home page, and a copyright page. The
[`footer`{#the-nav-element:the-footer-element}](#the-footer-element)
element alone is sufficient for such cases; while a
[`nav`{#the-nav-element:the-nav-element-3}](#the-nav-element) element
can be used in such cases, it is usually unnecessary.

User agents (such as screen readers) that are targeted at users who can
benefit from navigation information being omitted in the initial
rendering, or who can benefit from navigation information being
immediately available, can use this element as a way to determine what
content on the page to initially skip or provide on request (or both).

::: example
In the following example, there are two
[`nav`{#the-nav-element:the-nav-element-4}](#the-nav-element) elements,
one for primary navigation around the site, and one for secondary
navigation around the page itself.

``` html
<body>
 <h1>The Wiki Center Of Exampland</h1>
 <nav>
  <ul>
   <li><a href="/">Home</a></li>
   <li><a href="/events">Current Events</a></li>
   ...more...
  </ul>
 </nav>
 <article>
  <header>
   <h2>Demos in Exampland</h2>
   <p>Written by A. N. Other.</p>
  </header>
  <nav>
   <ul>
    <li><a href="#public">Public demonstrations</a></li>
    <li><a href="#destroy">Demolitions</a></li>
    ...more...
   </ul>
  </nav>
  <div>
   <section id="public">
    <h2>Public demonstrations</h2>
    <p>...more...</p>
   </section>
   <section id="destroy">
    <h2>Demolitions</h2>
    <p>...more...</p>
   </section>
   ...more...
  </div>
  <footer>
   <p><a href="?edit">Edit</a> | <a href="?delete">Delete</a> | <a href="?Rename">Rename</a></p>
  </footer>
 </article>
 <footer>
  <p><small>© copyright 1998 Exampland Emperor</small></p>
 </footer>
</body>
```
:::

::: example
In the following example, the page has several places where links are
present, but only one of those places is considered a navigation
section.

``` html
<body itemscope itemtype="http://schema.org/Blog">
 <header>
  <h1>Wake up sheeple!</h1>
  <p><a href="news.html">News</a> -
     <a href="blog.html">Blog</a> -
     <a href="forums.html">Forums</a></p>
  <p>Last Modified: <span itemprop="dateModified">2009-04-01</span></p>
  <nav>
   <h2>Navigation</h2>
   <ul>
    <li><a href="articles.html">Index of all articles</a></li>
    <li><a href="today.html">Things sheeple need to wake up for today</a></li>
    <li><a href="successes.html">Sheeple we have managed to wake</a></li>
   </ul>
  </nav>
 </header>
 <main>
  <article itemprop="blogPosts" itemscope itemtype="http://schema.org/BlogPosting">
   <header>
    <h2 itemprop="headline">My Day at the Beach</h2>
   </header>
   <div itemprop="articleBody">
    <p>Today I went to the beach and had a lot of fun.</p>
    ...more content...
   </div>
   <footer>
    <p>Posted <time itemprop="datePublished" datetime="2009-10-10">Thursday</time>.</p>
   </footer>
  </article>
  ...more blog posts...
 </main>
 <footer>
  <p>Copyright ©
   <span itemprop="copyrightYear">2010</span>
   <span itemprop="copyrightHolder">The Example Company</span>
  </p>
  <p><a href="about.html">About</a> -
     <a href="policy.html">Privacy Policy</a> -
     <a href="contact.html">Contact Us</a></p>
 </footer>
</body>
```

You can also see microdata annotations in the above example that use the
schema.org vocabulary to provide the publication date and other metadata
about the blog post.
:::

::: example
A [`nav`{#the-nav-element:the-nav-element-5}](#the-nav-element) element
doesn\'t have to contain a list, it can contain other kinds of content
as well. In this navigation block, links are provided in prose:

``` html
<nav>
 <h1>Navigation</h1>
 <p>You are on my home page. To the north lies <a href="/blog">my
 blog</a>, from whence the sounds of battle can be heard. To the east
 you can see a large mountain, upon which many <a
 href="/school">school papers</a> are littered. Far up thus mountain
 you can spy a little figure who appears to be me, desperately
 scribbling a <a href="/school/thesis">thesis</a>.</p>
 <p>To the west are several exits. One fun-looking exit is labeled <a
 href="https://games.example.com/">"games"</a>. Another more
 boring-looking exit is labeled <a
 href="https://isp.example.net/">ISP™</a>.</p>
 <p>To the south lies a dark and dank <a href="/about">contacts
 page</a>. Cobwebs cover its disused entrance, and at one point you
 see a rat run quickly out of the page.</p>
</nav>
```
:::

::: example
In this example,
[`nav`{#the-nav-element:the-nav-element-6}](#the-nav-element) is used in
an email application, to let the user switch folders:

``` html
<p><input type=button value="Compose" onclick="compose()"></p>
<nav>
 <h1>Folders</h1>
 <ul>
  <li> <a href="/inbox" onclick="return openFolder(this.href)">Inbox</a> <span class=count></span>
  <li> <a href="/sent" onclick="return openFolder(this.href)">Sent</a>
  <li> <a href="/drafts" onclick="return openFolder(this.href)">Drafts</a>
  <li> <a href="/trash" onclick="return openFolder(this.href)">Trash</a>
  <li> <a href="/customers" onclick="return openFolder(this.href)">Customers</a>
 </ul>
</nav>
```
:::

#### [4.3.5]{.secno} The [`aside`]{.dfn dfn-type="element"} element[](#the-aside-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/aside](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/aside "The <aside> HTML element represents a portion of a document whose content is only indirectly related to the document's main content. Asides are frequently presented as sidebars or call-out boxes.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome5+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11.1+]{.opera_android .yes}
:::
::::
:::::

[Categories](dom.html#concept-element-categories){#the-aside-element:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-aside-element:flow-content-2}.
:   [Sectioning
    content](dom.html#sectioning-content-2){#the-aside-element:sectioning-content-2}.
:   [Palpable
    content](dom.html#palpable-content-2){#the-aside-element:palpable-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-aside-element:concept-element-contexts}:
:   Where [sectioning
    content](dom.html#sectioning-content-2){#the-aside-element:sectioning-content-2-2}
    is expected.

[Content model](dom.html#concept-element-content-model){#the-aside-element:concept-element-content-model}:
:   [Flow
    content](dom.html#flow-content-2){#the-aside-element:flow-content-2-2}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-aside-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-aside-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-aside-element:global-attributes}

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-aside-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-aside).
:   [For implementers](https://w3c.github.io/html-aam/#el-aside).

[DOM interface](dom.html#concept-element-dom){#the-aside-element:concept-element-dom}:
:   Uses
    [`HTMLElement`{#the-aside-element:htmlelement}](dom.html#htmlelement).

The [`aside`{#the-aside-element:the-aside-element}](#the-aside-element)
element [represents](dom.html#represents){#the-aside-element:represents}
a section of a page that consists of content that is tangentially
related to the content around the
[`aside`{#the-aside-element:the-aside-element-2}](#the-aside-element)
element, and which could be considered separate from that content. Such
sections are often represented as sidebars in printed typography.

The element can be used for typographical effects like pull quotes or
sidebars, for advertising, for groups of
[`nav`{#the-aside-element:the-nav-element}](#the-nav-element) elements,
and for other content that is considered separate from the main content
of the page.

It\'s not appropriate to use the
[`aside`{#the-aside-element:the-aside-element-3}](#the-aside-element)
element just for parentheticals, since those are part of the main flow
of the document.

::: example
The following example shows how an aside is used to mark up background
material on Switzerland in a much longer news story on Europe.

``` html
<aside>
 <h2>Switzerland</h2>
 <p>Switzerland, a land-locked country in the middle of geographic
 Europe, has not joined the geopolitical European Union, though it is
 a signatory to a number of European treaties.</p>
</aside>
```
:::

::: example
The following example shows how an aside is used to mark up a pull quote
in a longer article.

``` html
...

<p>He later joined a large company, continuing on the same work.
<q>I love my job. People ask me what I do for fun when I'm not at
work. But I'm paid to do my hobby, so I never know what to
answer. Some people wonder what they would do if they didn't have to
work... but I know what I would do, because I was unemployed for a
year, and I filled that time doing exactly what I do now.</q></p>

<aside>
 <q>People ask me what I do for fun when I'm not at work. But I'm
 paid to do my hobby, so I never know what to answer.</q>
</aside>

<p>Of course his work — or should that be hobby? —
isn't his only passion. He also enjoys other pleasures.</p>

...
```
:::

::: example
The following extract shows how
[`aside`{#the-aside-element:the-aside-element-4}](#the-aside-element)
can be used for blogrolls and other side content on a blog:

``` html
<body>
 <header>
  <h1>My wonderful blog</h1>
  <p>My tagline</p>
 </header>
 <aside>
  <!-- this aside contains two sections that are tangentially related
  to the page, namely, links to other blogs, and links to blog posts
  from this blog -->
  <nav>
   <h2>My blogroll</h2>
   <ul>
    <li><a href="https://blog.example.com/">Example Blog</a>
   </ul>
  </nav>
  <nav>
   <h2>Archives</h2>
   <ol reversed>
    <li><a href="/last-post">My last post</a>
    <li><a href="/first-post">My first post</a>
   </ol>
  </nav>
 </aside>
 <aside>
  <!-- this aside is tangentially related to the page also, it
  contains twitter messages from the blog author -->
  <h1>Twitter Feed</h1>
  <blockquote cite="https://twitter.example.net/t31351234">
   I'm on vacation, writing my blog.
  </blockquote>
  <blockquote cite="https://twitter.example.net/t31219752">
   I'm going to go on vacation soon.
  </blockquote>
 </aside>
 <article>
  <!-- this is a blog post -->
  <h2>My last post</h2>
  <p>This is my last post.</p>
  <footer>
   <p><a href="/last-post" rel=bookmark>Permalink</a>
  </footer>
 </article>
 <article>
  <!-- this is also a blog post -->
  <h2>My first post</h2>
  <p>This is my first post.</p>
  <aside>
   <!-- this aside is about the blog post, since it's inside the
   <article> element; it would be wrong, for instance, to put the
   blogroll here, since the blogroll isn't really related to this post
   specifically, only to the page as a whole -->
   <h2>Posting</h2>
   <p>While I'm thinking about it, I wanted to say something about
   posting. Posting is fun!</p>
  </aside>
  <footer>
   <p><a href="/first-post" rel=bookmark>Permalink</a>
  </footer>
 </article>
 <footer>
  <p><a href="/archives">Archives</a> -
   <a href="/about">About me</a> -
   <a href="/copyright">Copyright</a></p>
 </footer>
</body>
```
:::

#### [4.3.6]{.secno} The [`h1`]{#the-h1-element .dfn dfn-type="element"}, [`h2`]{#the-h2-element .dfn dfn-type="element"}, [`h3`]{#the-h3-element .dfn dfn-type="element"}, [`h4`]{#the-h4-element .dfn dfn-type="element"}, [`h5`]{#the-h5-element .dfn dfn-type="element"}, and [`h6`]{#the-h6-element .dfn dfn-type="element"} elements[](#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements){.self-link} {#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements}

::::::::::::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/Heading_Elements](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements "The <h1> to <h6> HTML elements represent six levels of section headings. <h1> is the highest section level and <h6> is the lowest.")

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

:::: feature
[Element/Heading_Elements](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements "The <h1> to <h6> HTML elements represent six levels of section headings. <h1> is the highest section level and <h6> is the lowest.")

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

:::: feature
[Element/Heading_Elements](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements "The <h1> to <h6> HTML elements represent six levels of section headings. <h1> is the highest section level and <h6> is the lowest.")

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

:::: feature
[Element/Heading_Elements](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements "The <h1> to <h6> HTML elements represent six levels of section headings. <h1> is the highest section level and <h6> is the lowest.")

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

:::: feature
[Element/Heading_Elements](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements "The <h1> to <h6> HTML elements represent six levels of section headings. <h1> is the highest section level and <h6> is the lowest.")

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

:::: feature
[Element/Heading_Elements](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements "The <h1> to <h6> HTML elements represent six levels of section headings. <h1> is the highest section level and <h6> is the lowest.")

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
:::::::::::::::

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLHeadingElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLHeadingElement "The HTMLHeadingElement interface represents the different heading elements, <h1> through <h6>. It inherits methods and properties from the HTMLElement interface.")

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

[Categories](dom.html#concept-element-categories){#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements:flow-content-2}.
:   [Heading
    content](dom.html#heading-content-2){#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements:heading-content-2}.
:   [Palpable
    content](dom.html#palpable-content-2){#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements:palpable-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements:concept-element-contexts}:
:   As a child of an
    [`hgroup`{#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements:the-hgroup-element}](#the-hgroup-element)
    element.
:   Where [heading
    content](dom.html#heading-content-2){#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements:heading-content-2-2}
    is expected.

[Content model](dom.html#concept-element-content-model){#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements:concept-element-content-model}:
:   [Phrasing
    content](dom.html#phrasing-content-2){#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements:phrasing-content-2}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements:global-attributes}

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-h1-h6).
:   [For implementers](https://w3c.github.io/html-aam/#el-h1-h6).

[DOM interface](dom.html#concept-element-dom){#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLHeadingElement : HTMLElement {
      [HTMLConstructor] constructor();

      // also has obsolete members
    };
    ```

These elements
[represent](dom.html#represents){#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements:represents}
headings for their sections.

The semantics and meaning of these elements are defined in the section
on [headings and outlines](#headings-and-outlines).

These elements have a [heading
level](#heading-level){#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements:heading-level}
given by the number in their name. The [heading
level](#heading-level){#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements:heading-level-2}
corresponds to the levels of nested sections. The
[`h1`{#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements}](#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements)
element is for a top-level section,
[`h2`{#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-2}](#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements)
for a subsection,
[`h3`{#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-3}](#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements)
for a sub-subsection, and so on.

::: example
As far as their respective document outlines (their heading and section
structures) are concerned, these two snippets are semantically
equivalent:

``` html
<body>
<h1>Let's call it a draw(ing surface)</h1>
<h2>Diving in</h2>
<h2>Simple shapes</h2>
<h2>Canvas coordinates</h2>
<h3>Canvas coordinates diagram</h3>
<h2>Paths</h2>
</body>
```

``` html
<body>
 <h1>Let's call it a draw(ing surface)</h1>
 <section>
  <h2>Diving in</h2>
 </section>
 <section>
  <h2>Simple shapes</h2>
 </section>
 <section>
  <h2>Canvas coordinates</h2>
  <section>
   <h3>Canvas coordinates diagram</h3>
  </section>
 </section>
 <section>
  <h2>Paths</h2>
 </section>
</body>
```

Authors might prefer the former style for its terseness, or the latter
style for its additional styling hooks. Which is best is purely an issue
of preferred authoring style.
:::

#### [4.3.7]{.secno} The [`hgroup`]{.dfn dfn-type="element"} element[](#the-hgroup-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/hgroup](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/hgroup "The <hgroup> HTML element represents a heading and related content. It groups a single <h1>–<h6> element with one or more <p>.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome5+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android2.2+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11.1+]{.opera_android .yes}
:::
::::
:::::

[Categories](dom.html#concept-element-categories){#the-hgroup-element:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-hgroup-element:flow-content-2}.
:   [Heading
    content](dom.html#heading-content-2){#the-hgroup-element:heading-content-2}.
:   [Palpable
    content](dom.html#palpable-content-2){#the-hgroup-element:palpable-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-hgroup-element:concept-element-contexts}:
:   Where [heading
    content](dom.html#heading-content-2){#the-hgroup-element:heading-content-2-2}
    is expected.

[Content model](dom.html#concept-element-content-model){#the-hgroup-element:concept-element-content-model}:
:   Zero or more
    [`p`{#the-hgroup-element:the-p-element}](grouping-content.html#the-p-element)
    elements, followed by one
    [`h1`{#the-hgroup-element:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements}](#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements),
    [`h2`{#the-hgroup-element:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-2}](#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements),
    [`h3`{#the-hgroup-element:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-3}](#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements),
    [`h4`{#the-hgroup-element:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-4}](#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements),
    [`h5`{#the-hgroup-element:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-5}](#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements),
    or
    [`h6`{#the-hgroup-element:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-6}](#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements)
    element, followed by zero or more
    [`p`{#the-hgroup-element:the-p-element-2}](grouping-content.html#the-p-element)
    elements, optionally intermixed with [script-supporting
    elements](dom.html#script-supporting-elements-2){#the-hgroup-element:script-supporting-elements-2}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-hgroup-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-hgroup-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-hgroup-element:global-attributes}

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-hgroup-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-hgroup).
:   [For implementers](https://w3c.github.io/html-aam/#el-hgroup).

[DOM interface](dom.html#concept-element-dom){#the-hgroup-element:concept-element-dom}:
:   Uses
    [`HTMLElement`{#the-hgroup-element:htmlelement}](dom.html#htmlelement).

The
[`hgroup`{#the-hgroup-element:the-hgroup-element}](#the-hgroup-element)
element
[represents](dom.html#represents){#the-hgroup-element:represents} a
heading and related content. The element may be used to group an
[`h1`{#the-hgroup-element:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-7}](#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements)--[`h6`{#the-hgroup-element:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-8}](#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements)
element with one or more
[`p`{#the-hgroup-element:the-p-element-3}](grouping-content.html#the-p-element)
elements containing content representing a subheading, alternative
title, or tagline.

::: example
Here are some examples of valid headings contained within an
[`hgroup`{#the-hgroup-element:the-hgroup-element-2}](#the-hgroup-element)
element.

``` html
<hgroup>
 <h1>The reality dysfunction</h1>
 <p>Space is not the only void</p>
</hgroup>
```

``` html
<hgroup>
 <h1>Dr. Strangelove</h1>
 <p>Or: How I Learned to Stop Worrying and Love the Bomb</p>
</hgroup>
```
:::

#### [4.3.8]{.secno} The [`header`]{.dfn dfn-type="element"} element[](#the-header-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/header](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/header "The <header> HTML element represents introductory content, typically a group of introductory or navigational aids. It may contain some heading elements but also a logo, a search form, an author name, and other elements.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome5+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11.1+]{.opera_android .yes}
:::
::::
:::::

[Categories](dom.html#concept-element-categories){#the-header-element:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-header-element:flow-content-2}.
:   [Palpable
    content](dom.html#palpable-content-2){#the-header-element:palpable-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-header-element:concept-element-contexts}:
:   Where [flow
    content](dom.html#flow-content-2){#the-header-element:flow-content-2-2}
    is expected.

[Content model](dom.html#concept-element-content-model){#the-header-element:concept-element-content-model}:
:   [Flow
    content](dom.html#flow-content-2){#the-header-element:flow-content-2-3},
    but with no
    [`header`{#the-header-element:the-header-element}](#the-header-element)
    or
    [`footer`{#the-header-element:the-footer-element}](#the-footer-element)
    element descendants.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-header-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-header-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-header-element:global-attributes}

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-header-element:concept-element-accessibility-considerations}:
:   If there is an ancestor [sectioning
    content](dom.html#sectioning-content-2){#the-header-element:sectioning-content-2}
    element: [for authors](https://w3c.github.io/html-aria/#el-header);
    [for implementers](https://w3c.github.io/html-aam/#el-header).
:   Otherwise: [for
    authors](https://w3c.github.io/html-aria/#el-header); [for
    implementers](https://w3c.github.io/html-aam/#el-header-ancestorbody).

[DOM interface](dom.html#concept-element-dom){#the-header-element:concept-element-dom}:
:   Uses
    [`HTMLElement`{#the-header-element:htmlelement}](dom.html#htmlelement).

The
[`header`{#the-header-element:the-header-element-2}](#the-header-element)
element
[represents](dom.html#represents){#the-header-element:represents} a
group of introductory or navigational aids.

A
[`header`{#the-header-element:the-header-element-3}](#the-header-element)
element is intended to usually contain a heading (an
[`h1`{#the-header-element:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements}](#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements)--[`h6`{#the-header-element:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-2}](#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements)
element or an
[`hgroup`{#the-header-element:the-hgroup-element}](#the-hgroup-element)
element), but this is not required. The
[`header`{#the-header-element:the-header-element-4}](#the-header-element)
element can also be used to wrap a section\'s table of contents, a
search form, or any relevant logos.

::: example
Here are some sample headers. This first one is for a game:

``` html
<header>
 <p>Welcome to...</p>
 <h1>Voidwars!</h1>
</header>
```

The following snippet shows how the element can be used to mark up a
specification\'s header:

``` html
<header>
 <hgroup>
  <h1>Fullscreen API</h1>
  <p>Living Standard — Last Updated 19 October 2015<p>
 </hgroup>
 <dl>
  <dt>Participate:</dt>
  <dd><a href="https://github.com/whatwg/fullscreen">GitHub whatwg/fullscreen</a></dd>
  <dt>Commits:</dt>
  <dd><a href="https://github.com/whatwg/fullscreen/commits">GitHub whatwg/fullscreen/commits</a></dd>
 </dl>
</header>
```
:::

The
[`header`{#the-header-element:the-header-element-5}](#the-header-element)
element is not [sectioning
content](dom.html#sectioning-content-2){#the-header-element:sectioning-content-2-2};
it doesn\'t introduce a new section.

::: example
In this example, the page has a page heading given by the
[`h1`{#the-header-element:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-3}](#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements)
element, and two subsections whose headings are given by
[`h2`{#the-header-element:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-4}](#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements)
elements. The content after the
[`header`{#the-header-element:the-header-element-6}](#the-header-element)
element is still part of the last subsection started in the
[`header`{#the-header-element:the-header-element-7}](#the-header-element)
element, because the
[`header`{#the-header-element:the-header-element-8}](#the-header-element)
element doesn\'t take part in the
[outline](#outline){#the-header-element:outline} algorithm.

``` html
<body>
 <header>
  <h1>Little Green Guys With Guns</h1>
  <nav>
   <ul>
    <li><a href="/games">Games</a>
    <li><a href="/forum">Forum</a>
    <li><a href="/download">Download</a>
   </ul>
  </nav>
  <h2>Important News</h2> <!-- this starts a second subsection -->
  <!-- this is part of the subsection entitled "Important News" -->
  <p>To play today's games you will need to update your client.</p>
  <h2>Games</h2> <!-- this starts a third subsection -->
 </header>
 <p>You have three active games:</p>
 <!-- this is still part of the subsection entitled "Games" -->
 ...
```
:::

#### [4.3.9]{.secno} The [`footer`]{.dfn dfn-type="element"} element[](#the-footer-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/footer](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/footer "The <footer> HTML element represents a footer for its nearest ancestor sectioning content or sectioning root element. A <footer> typically contains information about the author of the section, copyright data or links to related documents.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome5+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11.1+]{.opera_android .yes}
:::
::::
:::::

[Categories](dom.html#concept-element-categories){#the-footer-element:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-footer-element:flow-content-2}.
:   [Palpable
    content](dom.html#palpable-content-2){#the-footer-element:palpable-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-footer-element:concept-element-contexts}:
:   Where [flow
    content](dom.html#flow-content-2){#the-footer-element:flow-content-2-2}
    is expected.

[Content model](dom.html#concept-element-content-model){#the-footer-element:concept-element-content-model}:
:   [Flow
    content](dom.html#flow-content-2){#the-footer-element:flow-content-2-3},
    but with no
    [`header`{#the-footer-element:the-header-element}](#the-header-element)
    or
    [`footer`{#the-footer-element:the-footer-element}](#the-footer-element)
    element descendants.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-footer-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-footer-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-footer-element:global-attributes}

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-footer-element:concept-element-accessibility-considerations}:
:   If there is an ancestor [sectioning
    content](dom.html#sectioning-content-2){#the-footer-element:sectioning-content-2}
    element: [for authors](https://w3c.github.io/html-aria/#el-footer);
    [for implementers](https://w3c.github.io/html-aam/#el-footer).
:   Otherwise: [for
    authors](https://w3c.github.io/html-aria/#el-footer); [for
    implementers](https://w3c.github.io/html-aam/#el-footer-ancestorbody).

[DOM interface](dom.html#concept-element-dom){#the-footer-element:concept-element-dom}:
:   Uses
    [`HTMLElement`{#the-footer-element:htmlelement}](dom.html#htmlelement).

The
[`footer`{#the-footer-element:the-footer-element-2}](#the-footer-element)
element
[represents](dom.html#represents){#the-footer-element:represents} a
footer for its nearest ancestor [sectioning
content](dom.html#sectioning-content-2){#the-footer-element:sectioning-content-2-2}
element, or for [the body
element](dom.html#the-body-element-2){#the-footer-element:the-body-element-2}
if there is no such ancestor. A footer typically contains information
about its section such as who wrote it, links to related documents,
copyright data, and the like.

When the
[`footer`{#the-footer-element:the-footer-element-3}](#the-footer-element)
element contains entire sections, they
[represent](dom.html#represents){#the-footer-element:represents-2}
appendices, indices, long colophons, verbose license agreements, and
other such content.

Contact information for the author or editor of a section belongs in an
[`address`{#the-footer-element:the-address-element}](#the-address-element)
element, possibly itself inside a
[`footer`{#the-footer-element:the-footer-element-4}](#the-footer-element).
Bylines and other information that could be suitable for both a
[`header`{#the-footer-element:the-header-element-2}](#the-header-element)
or a
[`footer`{#the-footer-element:the-footer-element-5}](#the-footer-element)
can be placed in either (or neither). The primary purpose of these
elements is merely to help the author write self-explanatory markup that
is easy to maintain and style; they are not intended to impose specific
structures on authors.

Footers don\'t necessarily have to appear at the *end* of a section,
though they usually do.

When there is no ancestor [sectioning
content](dom.html#sectioning-content-2){#the-footer-element:sectioning-content-2-3}
element, then it applies to the whole page.

The
[`footer`{#the-footer-element:the-footer-element-6}](#the-footer-element)
element is not itself [sectioning
content](dom.html#sectioning-content-2){#the-footer-element:sectioning-content-2-4};
it doesn\'t introduce a new section.

::: example
Here is a page with two footers, one at the top and one at the bottom,
with the same content:

``` html
<body>
 <footer><a href="../">Back to index...</a></footer>
 <hgroup>
  <h1>Lorem ipsum</h1>
  <p>The ipsum of all lorems</p>
 </hgroup>
 <p>A dolor sit amet, consectetur adipisicing elit, sed do eiusmod
 tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim
 veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex
 ea commodo consequat. Duis aute irure dolor in reprehenderit in
 voluptate velit esse cillum dolore eu fugiat nulla
 pariatur. Excepteur sint occaecat cupidatat non proident, sunt in
 culpa qui officia deserunt mollit anim id est laborum.</p>
 <footer><a href="../">Back to index...</a></footer>
</body>
```
:::

::: example
Here is an example which shows the
[`footer`{#the-footer-element:the-footer-element-7}](#the-footer-element)
element being used both for a site-wide footer and for a section footer.

``` html
<!DOCTYPE HTML>
<HTML LANG="en"><HEAD>
<TITLE>The Ramblings of a Scientist</TITLE>
<BODY>
<H1>The Ramblings of a Scientist</H1>
<ARTICLE>
 <H1>Episode 15</H1>
 <VIDEO SRC="/fm/015.ogv" CONTROLS PRELOAD>
  <P><A HREF="/fm/015.ogv">Download video</A>.</P>
 </VIDEO>
 <FOOTER> <!-- footer for article -->
  <P>Published <TIME DATETIME="2009-10-21T18:26-07:00">on 2009/10/21 at 6:26pm</TIME></P>
 </FOOTER>
</ARTICLE>
<ARTICLE>
 <H1>My Favorite Trains</H1>
 <P>I love my trains. My favorite train of all time is a Köf.</P>
 <P>It is fun to see them pull some coal cars because they look so
 dwarfed in comparison.</P>
 <FOOTER> <!-- footer for article -->
  <P>Published <TIME DATETIME="2009-09-15T14:54-07:00">on 2009/09/15 at 2:54pm</TIME></P>
 </FOOTER>
</ARTICLE>
<FOOTER> <!-- site wide footer -->
 <NAV>
  <P><A HREF="/credits.html">Credits</A> —
     <A HREF="/tos.html">Terms of Service</A> —
     <A HREF="/index.html">Blog Index</A></P>
 </NAV>
 <P>Copyright © 2009 Gordon Freeman</P>
</FOOTER>
</BODY>
</HTML>
```
:::

::: example
Some site designs have what is sometimes referred to as \"fat footers\"
--- footers that contain a lot of material, including images, links to
other articles, links to pages for sending feedback, special offers\...
in some ways, a whole \"front page\" in the footer.

This fragment shows the bottom of a page on a site with a \"fat
footer\":

``` html
...
 <footer>
  <nav>
   <section>
    <h1>Articles</h1>
    <p><img src="images/somersaults.jpeg" alt=""> Go to the gym with
    our somersaults class! Our teacher Jim takes you through the paces
    in this two-part article. <a href="articles/somersaults/1">Part
    1</a> · <a href="articles/somersaults/2">Part 2</a></p>
    <p><img src="images/kindplus.jpeg"> Tired of walking on the edge of
    a clif<!-- sic -->? Our guest writer Lara shows you how to bumble
    your way through the bars. <a href="articles/kindplus/1">Read
    more...</a></p>
    <p><img src="images/crisps.jpeg"> The chips are down, now all
    that's left is a potato. What can you do with it? <a
    href="articles/crisps/1">Read more...</a></p>
   </section>
   <ul>
    <li> <a href="/about">About us...</a>
    <li> <a href="/feedback">Send feedback!</a>
    <li> <a href="/sitemap">Sitemap</a>
   </ul>
  </nav>
  <p><small>Copyright © 2015 The Snacker —
  <a href="/tos">Terms of Service</a></small></p>
 </footer>
</body>
```
:::

#### [4.3.10]{.secno} The [`address`]{.dfn dfn-type="element"} element[](#the-address-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/address](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/address "The <address> HTML element indicates that the enclosed HTML provides contact information for a person or people, or for an organization.")

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

[Categories](dom.html#concept-element-categories){#the-address-element:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-address-element:flow-content-2}.
:   [Palpable
    content](dom.html#palpable-content-2){#the-address-element:palpable-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-address-element:concept-element-contexts}:
:   Where [flow
    content](dom.html#flow-content-2){#the-address-element:flow-content-2-2}
    is expected.

[Content model](dom.html#concept-element-content-model){#the-address-element:concept-element-content-model}:
:   [Flow
    content](dom.html#flow-content-2){#the-address-element:flow-content-2-3},
    but with no [heading
    content](dom.html#heading-content-2){#the-address-element:heading-content-2}
    descendants, no [sectioning
    content](dom.html#sectioning-content-2){#the-address-element:sectioning-content-2}
    descendants, and no
    [`header`{#the-address-element:the-header-element}](#the-header-element),
    [`footer`{#the-address-element:the-footer-element}](#the-footer-element),
    or
    [`address`{#the-address-element:the-address-element}](#the-address-element)
    element descendants.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-address-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-address-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-address-element:global-attributes}

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-address-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-address).
:   [For implementers](https://w3c.github.io/html-aam/#el-address).

[DOM interface](dom.html#concept-element-dom){#the-address-element:concept-element-dom}:
:   Uses
    [`HTMLElement`{#the-address-element:htmlelement}](dom.html#htmlelement).

The
[`address`{#the-address-element:the-address-element-2}](#the-address-element)
element
[represents](dom.html#represents){#the-address-element:represents} the
contact information for its nearest
[`article`{#the-address-element:the-article-element}](#the-article-element)
or [`body`{#the-address-element:the-body-element}](#the-body-element)
element ancestor. If that is [the body
element](dom.html#the-body-element-2){#the-address-element:the-body-element-2},
then the contact information applies to the document as a whole.

::: example
For example, a page at the W3C web site related to HTML might include
the following contact information:

``` html
<ADDRESS>
 <A href="../People/Raggett/">Dave Raggett</A>,
 <A href="../People/Arnaud/">Arnaud Le Hors</A>,
 contact persons for the <A href="Activity">W3C HTML Activity</A>
</ADDRESS>
```
:::

The
[`address`{#the-address-element:the-address-element-3}](#the-address-element)
element must not be used to represent arbitrary addresses (e.g. postal
addresses), unless those addresses are in fact the relevant contact
information. (The
[`p`{#the-address-element:the-p-element}](grouping-content.html#the-p-element)
element is the appropriate element for marking up postal addresses in
general.)

The
[`address`{#the-address-element:the-address-element-4}](#the-address-element)
element must not contain information other than contact information.

::: example
For example, the following is non-conforming use of the
[`address`{#the-address-element:the-address-element-5}](#the-address-element)
element:

``` bad
<ADDRESS>Last Modified: 1999/12/24 23:37:50</ADDRESS>
```
:::

Typically, the
[`address`{#the-address-element:the-address-element-6}](#the-address-element)
element would be included along with other information in a
[`footer`{#the-address-element:the-footer-element-2}](#the-footer-element)
element.

The contact information for a node `node`{.variable} is a collection of
[`address`{#the-address-element:the-address-element-7}](#the-address-element)
elements defined by the first applicable entry from the following list:

If `node`{.variable} is an [`article`{#the-address-element:the-article-element-2}](#the-article-element) element\
If `node`{.variable} is a [`body`{#the-address-element:the-body-element-3}](#the-body-element) element

:   The contact information consists of all the
    [`address`{#the-address-element:the-address-element-8}](#the-address-element)
    elements that have `node`{.variable} as an ancestor and do not have
    another
    [`body`{#the-address-element:the-body-element-4}](#the-body-element)
    or
    [`article`{#the-address-element:the-article-element-3}](#the-article-element)
    element ancestor that is a descendant of `node`{.variable}.

If `node`{.variable} has an ancestor element that is an [`article`{#the-address-element:the-article-element-4}](#the-article-element) element\
If `node`{.variable} has an ancestor element that is a [`body`{#the-address-element:the-body-element-5}](#the-body-element) element

:   The contact information of `node`{.variable} is the same as the
    contact information of the nearest
    [`article`{#the-address-element:the-article-element-5}](#the-article-element)
    or
    [`body`{#the-address-element:the-body-element-6}](#the-body-element)
    element ancestor, whichever is nearest.

If `node`{.variable}\'s [node document](https://dom.spec.whatwg.org/#concept-node-document){#the-address-element:node-document x-internal="node-document"} has [a body element](dom.html#the-body-element-2){#the-address-element:the-body-element-2-2}

:   The contact information of `node`{.variable} is the same as the
    contact information of [the body
    element](dom.html#the-body-element-2){#the-address-element:the-body-element-2-3}
    of the
    [`Document`{#the-address-element:document}](dom.html#document).

Otherwise

:   There is no contact information for `node`{.variable}.

User agents may expose the contact information of a node to the user, or
use it for other purposes, such as indexing sections based on the
sections\' contact information.

::: example
In this example the footer contains contact information and a copyright
notice.

``` html
<footer>
 <address>
  For more details, contact
  <a href="mailto:js@example.com">John Smith</a>.
 </address>
 <p><small>© copyright 2038 Example Corp.</small></p>
</footer>
```
:::

#### [4.3.11]{.secno} []{#headings-and-outlines}[]{#outlines}Headings and outlines[](#headings-and-outlines-2){.self-link} {#headings-and-outlines-2}

[`h1`{#headings-and-outlines-2:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements}](#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements)--[`h6`{#headings-and-outlines-2:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-2}](#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements)
elements have a [heading level]{#heading-level .dfn}, which is given by
the number in the element\'s name.

These elements
[represent](dom.html#represents){#headings-and-outlines-2:represents}
[headings]{#concept-heading .dfn}. The lower a
[heading](#concept-heading){#headings-and-outlines-2:concept-heading}\'s
[heading level](#heading-level){#headings-and-outlines-2:heading-level}
is, the fewer ancestor sections the
[heading](#concept-heading){#headings-and-outlines-2:concept-heading-2}
has.

The [outline]{#outline .dfn} is all
[headings](#concept-heading){#headings-and-outlines-2:concept-heading-3}
in a document, in [tree
order](https://dom.spec.whatwg.org/#concept-tree-order){#headings-and-outlines-2:tree-order
x-internal="tree-order"}.

The [outline](#outline){#headings-and-outlines-2:outline} should be used
for generating document outlines, for example when generating tables of
contents. When creating an interactive table of contents, entries should
jump the user to the relevant
[heading](#concept-heading){#headings-and-outlines-2:concept-heading-4}.

If a document has one or more
[headings](#concept-heading){#headings-and-outlines-2:concept-heading-5},
at least a single
[heading](#concept-heading){#headings-and-outlines-2:concept-heading-6}
within the [outline](#outline){#headings-and-outlines-2:outline-2}
should have a [heading
level](#heading-level){#headings-and-outlines-2:heading-level-2} of 1.

Each
[heading](#concept-heading){#headings-and-outlines-2:concept-heading-7}
following another
[heading](#concept-heading){#headings-and-outlines-2:concept-heading-8}
`lead`{.variable} in the
[outline](#outline){#headings-and-outlines-2:outline-3} must have a
[heading
level](#heading-level){#headings-and-outlines-2:heading-level-3} that is
less than, equal to, or 1 greater than `lead`{.variable}\'s [heading
level](#heading-level){#headings-and-outlines-2:heading-level-4}.

::: example
The following example is non-conforming:

``` bad
<body>
 <h1>Apples</h1>
 <p>Apples are fruit.</p>
 <section>
  <h3>Taste</h3>
  <p>They taste lovely.</p>
 </section>
</body>
```

It could be written as follows and then it would be conforming:

``` html
<body>
 <h1>Apples</h1>
 <p>Apples are fruit.</p>
 <section>
  <h2>Taste</h2>
  <p>They taste lovely.</p>
 </section>
</body>
```
:::

##### [4.3.11.1]{.secno} Sample outlines[](#sample-outlines){.self-link}

::: example
The following markup fragment:

``` html
<body>
  <hgroup id="document-title">
    <h1>HTML: Living Standard</h1>
    <p>Last Updated 12 August 2016</p>
  </hgroup>
  <p>Some intro to the document.</p>
  <h2>Table of contents</h2>
  <ol id=toc>...</ol>
  <h2>First section</h2>
  <p>Some intro to the first section.</p>
</body>
```

\...results in 3 document headings:

1.  `<h1>HTML: Living Standard</h1>`

2.  `<h2>Table of contents</h2>`.

3.  `<h2>First section</h2>`.

A rendered view of the [outline](#outline){#sample-outlines:outline}
might look like:

![Top-level section with the heading \"HTML: Living Standard\" and two
subsections; \"Table of contents\" and \"First
section\".](/images/outline.svg){.darkmode-aware width="465"
height="120"}
:::

::: example
First, here is a document, which is a book with very short chapters and
subsections:

``` html
<!DOCTYPE HTML>
<html lang=en>
<title>The Tax Book (all in one page)</title>
<h1>The Tax Book</h1>
<h2>Earning money</h2>
<p>Earning money is good.</p>
<h3>Getting a job</h3>
<p>To earn money you typically need a job.</p>
<h2>Spending money</h2>
<p>Spending is what money is mainly used for.</p>
<h3>Cheap things</h3>
<p>Buying cheap things often not cost-effective.</p>
<h3>Expensive things</h3>
<p>The most expensive thing is often not the most cost-effective either.</p>
<h2>Investing money</h2>
<p>You can lend your money to other people.</p>
<h2>Losing money</h2>
<p>If you spend money or invest money, sooner or later you will lose money.
<h3>Poor judgement</h3>
<p>Usually if you lose money it's because you made a mistake.</p>
```

Its [outline](#outline){#sample-outlines:outline-2} could be presented
as follows:

1.  The Tax Book
    1.  Earning money
        1.  Getting a job
    2.  Spending money
        1.  Cheap things
        2.  Expensive things
    3.  Investing money
    4.  Losing money
        1.  Poor judgement

Notice that the
[`title`{#sample-outlines:the-title-element}](semantics.html#the-title-element)
element is not a
[heading](#concept-heading){#sample-outlines:concept-heading}.
:::

::: example
A document can contain multiple top-level headings:

``` html
<!DOCTYPE HTML>
<html lang=en>
<title>Alphabetic Fruit</title>
<h1>Apples</h1>
<p>Pomaceous.</p>
<h1>Bananas</h1>
<p>Edible.</p>
<h1>Carambola</h1>
<p>Star.</p>
```

The document\'s [outline](#outline){#sample-outlines:outline-3} could be
presented as follows:

1.  Apples
2.  Bananas
3.  Carambola
:::

::: example
[`header`{#sample-outlines:the-header-element}](#the-header-element)
elements do not influence the
[outline](#outline){#sample-outlines:outline-4} of a document:

``` html
<!DOCTYPE HTML>
<html lang="en">
<title>We're adopting a child! — Ray's blog</title>
<h1>Ray's blog</h1>
<article>
 <header>
  <nav>
   <a href="?t=-1d">Yesterday</a>;
   <a href="?t=-7d">Last week</a>;
   <a href="?t=-1m">Last month</a>
  </nav>
  <h2>We're adopting a child!</h2>
 </header>
 <p>As of today, Janine and I have signed the papers to become
 the proud parents of baby Diane! We've been looking forward to
 this day for weeks.</p>
</article>
</html>
```

The document\'s [outline](#outline){#sample-outlines:outline-5} could be
presented as follows:

1.  Ray\'s blog
    1.  We\'re adopting a child!
:::

------------------------------------------------------------------------

::: example
The following example is conforming, but not encouraged as it has no
[heading](#concept-heading){#sample-outlines:concept-heading-2} whose
[heading level](#heading-level){#sample-outlines:heading-level} is 1:

``` html
<!DOCTYPE HTML>
<html lang=en>
<title>Alphabetic Fruit</title>
<section>
 <h2>Apples</h2>
 <p>Pomaceous.</p>
</section>
<section>
 <h2>Bananas</h2>
 <p>Edible.</p>
</section>
<section>
 <h2>Carambola</h2>
 <p>Star.</p>
</section>
```

The document\'s [outline](#outline){#sample-outlines:outline-6} could be
presented as follows:

1.  1.  Apples
    2.  Bananas
    3.  Carambola
:::

::: example
The following example is conforming, but not encouraged as the first
[heading](#concept-heading){#sample-outlines:concept-heading-3}\'s
[heading level](#heading-level){#sample-outlines:heading-level-2} is not
1:

``` html
<!DOCTYPE HTML>
<html lang=en>
<title>Feathers on The Site of Encyclopedic Knowledge</title>
 <h2>A plea from our caretakers</h2>
 <p>Please, we beg of you, send help! We're stuck in the server room!</p>
<h1>Feathers</h1>
<p>Epidermal growths.</p>
```

The document\'s [outline](#outline){#sample-outlines:outline-7} could be
presented as follows:

1.  1.  A plea from our caretakers
2.  Feathers
:::

##### [4.3.11.2]{.secno} Exposing outlines to users[](#exposing-outlines-to-users){.self-link}

User agents are encouraged to expose page
[outlines](#outline){#exposing-outlines-to-users:outline} to users to
aid in navigation. This is especially true for non-visual media, e.g.
screen readers.

::: example
For instance, a user agent could map the arrow keys as follows:

[[Shift]{.kbd} + [← Left]{.kbd}]{.kbd}
:   Go to previous heading

[[Shift]{.kbd} + [→ Right]{.kbd}]{.kbd}
:   Go to next heading

[[Shift]{.kbd} + [↑ Up]{.kbd}]{.kbd}
:   Go to next heading whose
    [level](#heading-level){#exposing-outlines-to-users:heading-level}
    is one less than the current heading\'s level

[[Shift]{.kbd} + [↓ Down]{.kbd}]{.kbd}
:   Go to next heading whose
    [level](#heading-level){#exposing-outlines-to-users:heading-level-2}
    is the same as the current heading\'s level
:::

#### [4.3.12]{.secno} Usage summary[](#usage-summary-2){.self-link} {#usage-summary-2}

*This section is non-normative.*

Element

Purpose

Example

[`body`{#usage-summary-2:the-body-element}](#the-body-element)

The contents of the document.

``` example
<!DOCTYPE HTML>
<html lang="en">
 <head> <title>Steve Hill's Home Page</title> </head>
 <body> <p>Hard Trance is My Life.</p> </body>
</html>
```

[`article`{#usage-summary-2:the-article-element}](#the-article-element)

A complete, or self-contained, composition in a document, page,
application, or site and that is, in principle, independently
distributable or reusable, e.g. in syndication. This could be a forum
post, a magazine or newspaper article, a blog entry, a user-submitted
comment, an interactive widget or gadget, or any other independent item
of content.

``` example
<article>
 <img src="/tumblr_masqy2s5yn1rzfqbpo1_500.jpg" alt="Yellow smiley face with the caption 'masif'">
 <p>My fave Masif tee so far!</p>
 <footer>Posted 2 days ago</footer>
</article>
<article>
 <img src="/tumblr_m9tf6wSr6W1rzfqbpo1_500.jpg" alt="">
 <p>Happy 2nd birthday Masif Saturdays!!!</p>
 <footer>Posted 3 weeks ago</footer>
</article>
```

[`section`{#usage-summary-2:the-section-element}](#the-section-element)

A generic section of a document or application. A section, in this
context, is a thematic grouping of content, typically with a heading.

``` example
<h1>Biography</h1>
<section>
 <h1>The facts</h1>
 <p>1500+ shows, 14+ countries</p>
</section>
<section>
 <h1>2010/2011 figures per year</h1>
 <p>100+ shows, 8+ countries</p>
</section>
```

[`nav`{#usage-summary-2:the-nav-element}](#the-nav-element)

A section of a page that links to other pages or to parts within the
page: a section with navigation links.

``` example
<nav>
 <p><a href="/">Home</a>
 <p><a href="/biog.html">Bio</a>
 <p><a href="/discog.html">Discog</a>
</nav>
```

[`aside`{#usage-summary-2:the-aside-element}](#the-aside-element)

A section of a page that consists of content that is tangentially
related to the content around the
[`aside`{#usage-summary-2:the-aside-element-2}](#the-aside-element)
element, and which could be considered separate from that content. Such
sections are often represented as sidebars in printed typography.

``` example
<h1>Music</h1>
<p>As any burner can tell you, the event has a lot of trance.</p>
<aside>You can buy the music we played at our <a href="buy.html">playlist page</a>.</aside>
<p>This year we played a kind of trance that originated in Belgium, Germany, and the Netherlands in the mid-90s.</p>
```

[`h1`{#usage-summary-2:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements}](#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements)--[`h6`{#usage-summary-2:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-2}](#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements)

A heading

``` example
<h1>The Guide To Music On The Playa</h1>
<h2>The Main Stage</h2>
<p>If you want to play on a stage, you should bring one.</p>
<h2>Amplified Music</h2>
<p>Amplifiers up to 300W or 90dB are welcome.</p>
```

[`hgroup`{#usage-summary-2:the-hgroup-element}](#the-hgroup-element)

A heading and related content. The element may be used to group an
[`h1`{#usage-summary-2:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-3}](#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements)--[`h6`{#usage-summary-2:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-4}](#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements)
element with one or more
[`p`{#usage-summary-2:the-p-element}](grouping-content.html#the-p-element)
elements containing content representing a subheading, alternative
title, or tagline.

``` example
<hgroup>
 <h1>Burning Music</h1>
 <p>The Guide To Music On The Playa</p>
</hgroup>
<section>
 <hgroup>
  <h1>Main Stage</h1>
  <p>The Fiction Of A Music Festival</p>
 </hgroup>
 <p>If you want to play on a stage, you should bring one.</p>
</section>
<section>
 <hgroup>
  <h1>Loudness!</h1>
  <p>Questions About Amplified Music</p>
 </hgroup>
 <p>Amplifiers up to 300W or 90dB are welcome.</p>
</section>
```

[`header`{#usage-summary-2:the-header-element}](#the-header-element)

A group of introductory or navigational aids.

``` example
<article>
 <header>
  <h1>Hard Trance is My Life</h1>
  <p>By DJ Steve Hill and Technikal</p>
 </header>
 <p>The album with the amusing punctuation has red artwork.</p>
</article>
```

[`footer`{#usage-summary-2:the-footer-element}](#the-footer-element)

A footer for its nearest ancestor [sectioning
content](dom.html#sectioning-content-2){#usage-summary-2:sectioning-content-2}
element, or for [the body
element](dom.html#the-body-element-2){#usage-summary-2:the-body-element-2}
if there is no such ancestor. A footer typically contains information
about its section such as who wrote it, links to related documents,
copyright data, and the like.

``` example
<article>
 <h1>Hard Trance is My Life</h1>
 <p>The album with the amusing punctuation has red artwork.</p>
 <footer>
  <p>Artists: DJ Steve Hill and Technikal</p>
 </footer>
</article>
```

##### [4.3.12.1]{.secno} Article or section?[](#article-or-section){.self-link}

*This section is non-normative.*

A
[`section`{#article-or-section:the-section-element}](#the-section-element)
forms part of something else. An
[`article`{#article-or-section:the-article-element}](#the-article-element)
is its own thing. But how does one know which is which? Mostly the real
answer is \"it depends on author intent\".

For example, one could imagine a book with a \"Granny Smith\" chapter
that just said \"These juicy, green apples make a great filling for
apple pies.\"; that would be a
[`section`{#article-or-section:the-section-element-2}](#the-section-element)
because there\'d be lots of other chapters on (maybe) other kinds of
apples.

On the other hand, one could imagine a tweet or reddit comment or tumblr
post or newspaper classified ad that just said \"Granny Smith. These
juicy, green apples make a great filling for apple pies.\"; it would
then be
[`article`{#article-or-section:the-article-element-2}](#the-article-element)s
because that was the whole thing.

A comment on an article is not part of the
[`article`{#article-or-section:the-article-element-3}](#the-article-element)
on which it is commenting, therefore it is its own
[`article`{#article-or-section:the-article-element-4}](#the-article-element).

[← 4 The elements of HTML](semantics.html) --- [Table of Contents](./)
--- [4.4 Grouping content →](grouping-content.html)
