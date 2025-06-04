## [4]{.secno} The elements of HTML[](#semantics){.self-link} {#semantics}

### [4.1]{.secno} The document element[](#the-root-element){.self-link} {#the-root-element}

#### [4.1.1]{.secno} The [`html`]{.dfn dfn-type="element"} element[](#the-html-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/html](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/html "The <html> HTML element represents the root (top-level element) of an HTML document, so it is also referred to as the root element. All other elements must be descendants of this element.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
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
[HTMLHtmlElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLHtmlElement "The HTMLHtmlElement interface serves as the root node for a given HTML document. This object inherits the properties and methods described in the HTMLElement interface.")

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
:::::

[Categories](dom.html#concept-element-categories){#the-html-element:concept-element-categories}:
:   None.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-html-element:concept-element-contexts}:
:   As document\'s [document
    element](https://dom.spec.whatwg.org/#document-element){#the-html-element:document-element
    x-internal="document-element"}.
:   Wherever a subdocument fragment is allowed in a compound document.

[Content model](dom.html#concept-element-content-model){#the-html-element:concept-element-content-model}:
:   A [`head`{#the-html-element:the-head-element}](#the-head-element)
    element followed by a
    [`body`{#the-html-element:the-body-element}](sections.html#the-body-element)
    element.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-html-element:concept-element-tag-omission}:
:   An [`html`{#the-html-element:the-html-element}](#the-html-element)
    element\'s [start
    tag](syntax.html#syntax-start-tag){#the-html-element:syntax-start-tag}
    can be omitted if the first thing inside the
    [`html`{#the-html-element:the-html-element-2}](#the-html-element)
    element is not a
    [comment](syntax.html#syntax-comments){#the-html-element:syntax-comments}.
:   An [`html`{#the-html-element:the-html-element-3}](#the-html-element)
    element\'s [end
    tag](syntax.html#syntax-end-tag){#the-html-element:syntax-end-tag}
    can be omitted if the
    [`html`{#the-html-element:the-html-element-4}](#the-html-element)
    element is not immediately followed by a
    [comment](syntax.html#syntax-comments){#the-html-element:syntax-comments-2}.

[Content attributes](dom.html#concept-element-attributes){#the-html-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-html-element:global-attributes}

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-html-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-html).
:   [For implementers](https://w3c.github.io/html-aam/#el-html).

[DOM interface](dom.html#concept-element-dom){#the-html-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLHtmlElement : HTMLElement {
      [HTMLConstructor] constructor();

      // also has obsolete members
    };
    ```

The [`html`{#the-html-element:the-html-element-5}](#the-html-element)
element [represents](dom.html#represents){#the-html-element:represents}
the root of an HTML document.

Authors are encouraged to specify a
[`lang`{#the-html-element:attr-lang}](dom.html#attr-lang) attribute on
the root
[`html`{#the-html-element:the-html-element-6}](#the-html-element)
element, giving the document\'s language. This aids speech synthesis
tools to determine what pronunciations to use, translation tools to
determine what rules to use, and so forth.

::: example
The [`html`{#the-html-element:the-html-element-7}](#the-html-element)
element in the following example declares that the document\'s language
is English.

``` html
<!DOCTYPE html>
<html lang="en">
<head>
<title>Swapping Songs</title>
</head>
<body>
<h1>Swapping Songs</h1>
<p>Tonight I swapped some of the songs I wrote with some friends, who
gave me some of the songs they wrote. I love sharing my music.</p>
</body>
</html>
```
:::

### [4.2]{.secno} Document metadata[](#document-metadata){.self-link}

#### [4.2.1]{.secno} The [`head`]{.dfn dfn-type="element"} element[](#the-head-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/head](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/head "The <head> HTML element contains machine-readable information (metadata) about the document, like its title, scripts, and style sheets.")

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
[HTMLHeadElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLHeadElement "The HTMLHeadElement interface contains the descriptive information, or metadata, for a document. This object inherits all of the properties and methods described in the HTMLElement interface.")

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
:::::

[Categories](dom.html#concept-element-categories){#the-head-element:concept-element-categories}:
:   None.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-head-element:concept-element-contexts}:
:   As the first element in an
    [`html`{#the-head-element:the-html-element}](#the-html-element)
    element.

[Content model](dom.html#concept-element-content-model){#the-head-element:concept-element-content-model}:
:   If the document is [an `iframe` `srcdoc`
    document](iframe-embed-object.html#an-iframe-srcdoc-document){#the-head-element:an-iframe-srcdoc-document}
    or if title information is available from a higher-level protocol:
    Zero or more elements of [metadata
    content](dom.html#metadata-content-2){#the-head-element:metadata-content-2},
    of which no more than one is a
    [`title`{#the-head-element:the-title-element}](#the-title-element)
    element and no more than one is a
    [`base`{#the-head-element:the-base-element}](#the-base-element)
    element.
:   Otherwise: One or more elements of [metadata
    content](dom.html#metadata-content-2){#the-head-element:metadata-content-2-2},
    of which exactly one is a
    [`title`{#the-head-element:the-title-element-2}](#the-title-element)
    element and no more than one is a
    [`base`{#the-head-element:the-base-element-2}](#the-base-element)
    element.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-head-element:concept-element-tag-omission}:
:   A [`head`{#the-head-element:the-head-element}](#the-head-element)
    element\'s [start
    tag](syntax.html#syntax-start-tag){#the-head-element:syntax-start-tag}
    can be omitted if the element is empty, or if the first thing inside
    the
    [`head`{#the-head-element:the-head-element-2}](#the-head-element)
    element is an element.
:   A [`head`{#the-head-element:the-head-element-3}](#the-head-element)
    element\'s [end
    tag](syntax.html#syntax-end-tag){#the-head-element:syntax-end-tag}
    can be omitted if the
    [`head`{#the-head-element:the-head-element-4}](#the-head-element)
    element is not immediately followed by [ASCII
    whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#the-head-element:space-characters
    x-internal="space-characters"} or a
    [comment](syntax.html#syntax-comments){#the-head-element:syntax-comments}.

[Content attributes](dom.html#concept-element-attributes){#the-head-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-head-element:global-attributes}

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-head-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-head).
:   [For implementers](https://w3c.github.io/html-aam/#el-head).

[DOM interface](dom.html#concept-element-dom){#the-head-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLHeadElement : HTMLElement {
      [HTMLConstructor] constructor();
    };
    ```

The [`head`{#the-head-element:the-head-element-5}](#the-head-element)
element [represents](dom.html#represents){#the-head-element:represents}
a collection of metadata for the
[`Document`{#the-head-element:document}](dom.html#document).

::: example
The collection of metadata in a
[`head`{#the-head-element:the-head-element-6}](#the-head-element)
element can be large or small. Here is an example of a very short one:

``` html
<!doctype html>
<html lang=en>
 <head>
  <title>A document with a short head</title>
 </head>
 <body>
 ...
```

Here is an example of a longer one:

``` html
<!DOCTYPE HTML>
<HTML LANG="EN">
 <HEAD>
  <META CHARSET="UTF-8">
  <BASE HREF="https://www.example.com/">
  <TITLE>An application with a long head</TITLE>
  <LINK REL="STYLESHEET" HREF="default.css">
  <LINK REL="STYLESHEET ALTERNATE" HREF="big.css" TITLE="Big Text">
  <SCRIPT SRC="support.js"></SCRIPT>
  <META NAME="APPLICATION-NAME" CONTENT="Long headed application">
 </HEAD>
 <BODY>
 ...
```
:::

The [`title`{#the-head-element:the-title-element-3}](#the-title-element)
element is a required child in most situations, but when a higher-level
protocol provides title information, e.g., in the subject line of an
email when HTML is used as an email authoring format, the
[`title`{#the-head-element:the-title-element-4}](#the-title-element)
element can be omitted.

#### [4.2.2]{.secno} The [`title`]{.dfn dfn-type="element"} element[](#the-title-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/title](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/title "The <title> HTML element defines the document's title that is shown in a browser's title bar or a page's tab. It only contains text; tags within the element are ignored.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer1+]{.ie .yes}

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
[HTMLTitleElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLTitleElement "The HTMLTitleElement interface is implemented by a document's <title>. This element inherits all of the properties and methods of the HTMLElement interface.")

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

[Categories](dom.html#concept-element-categories){#the-title-element:concept-element-categories}:
:   [Metadata
    content](dom.html#metadata-content-2){#the-title-element:metadata-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-title-element:concept-element-contexts}:
:   In a
    [`head`{#the-title-element:the-head-element}](#the-head-element)
    element containing no other
    [`title`{#the-title-element:the-title-element}](#the-title-element)
    elements.

[Content model](dom.html#concept-element-content-model){#the-title-element:concept-element-content-model}:
:   [Text](dom.html#text-content){#the-title-element:text-content} that
    is not [inter-element
    whitespace](dom.html#inter-element-whitespace){#the-title-element:inter-element-whitespace}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-title-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-title-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-title-element:global-attributes}

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-title-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-title).
:   [For implementers](https://w3c.github.io/html-aam/#el-title).

[DOM interface](dom.html#concept-element-dom){#the-title-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLTitleElement : HTMLElement {
      [HTMLConstructor] constructor();

      [CEReactions] attribute DOMString text;
    };
    ```

The
[`title`{#the-title-element:the-title-element-2}](#the-title-element)
element [represents](dom.html#represents){#the-title-element:represents}
the document\'s title or name. Authors should use titles that identify
their documents even when they are used out of context, for example in a
user\'s history or bookmarks, or in search results. The document\'s
title is often different from its first heading, since the first heading
does not have to stand alone when taken out of context.

There must be no more than one
[`title`{#the-title-element:the-title-element-3}](#the-title-element)
element per document.

If it\'s reasonable for the
[`Document`{#the-title-element:document}](dom.html#document) to have no
title, then the
[`title`{#the-title-element:the-title-element-4}](#the-title-element)
element is probably not required. See the
[`head`{#the-title-element:the-head-element-2}](#the-head-element)
element\'s content model for a description of when the element is
required.

`title`{.variable}`.`[`text`](#dom-title-text){#dom-title-text-dev}` [ = ``value`{.variable}` ]`

:   Returns the [child text
    content](https://dom.spec.whatwg.org/#concept-child-text-content){#the-title-element:child-text-content
    x-internal="child-text-content"} of the element.

    Can be set, to replace the element\'s children with the given value.

The [`text`]{#dom-title-text .dfn dfn-for="HTMLTitleElement"
dfn-type="attribute"} attribute\'s getter must return this
[`title`{#the-title-element:the-title-element-5}](#the-title-element)
element\'s [child text
content](https://dom.spec.whatwg.org/#concept-child-text-content){#the-title-element:child-text-content-2
x-internal="child-text-content"}.

The [`text`{#the-title-element:dom-title-text-2}](#dom-title-text)
attribute\'s setter must [string replace
all](https://dom.spec.whatwg.org/#string-replace-all){#the-title-element:string-replace-all
x-internal="string-replace-all"} with the given value within this
[`title`{#the-title-element:the-title-element-6}](#the-title-element)
element.

::: example
Here are some examples of appropriate titles, contrasted with the
top-level headings that might be used on those same pages.

``` html
  <title>Introduction to The Mating Rituals of Bees</title>
    ...
  <h1>Introduction</h1>
  <p>This companion guide to the highly successful
  <cite>Introduction to Medieval Bee-Keeping</cite> book is...
```

The next page might be a part of the same site. Note how the title
describes the subject matter unambiguously, while the first heading
assumes the reader knows what the context is and therefore won\'t wonder
if the dances are Salsa or Waltz:

``` html
  <title>Dances used during bee mating rituals</title>
    ...
  <h1>The Dances</h1>
```
:::

The string to use as the document\'s title is given by the
[`document.title`{#the-title-element:document.title}](dom.html#document.title)
IDL attribute.

User agents should use the document\'s title when referring to the
document in their user interface. When the contents of a
[`title`{#the-title-element:the-title-element-7}](#the-title-element)
element are used in this way, [the
directionality](dom.html#the-directionality){#the-title-element:the-directionality}
of that
[`title`{#the-title-element:the-title-element-8}](#the-title-element)
element should be used to set the directionality of the document\'s
title in the user interface.

#### [4.2.3]{.secno} The [`base`]{.dfn dfn-type="element"} element[](#the-base-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/base](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/base "The <base> HTML element specifies the base URL to use for all relative URLs in a document. There can be only one <base> element in a document.")

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
[HTMLBaseElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLBaseElement "The HTMLBaseElement interface contains the base URI for a document. This object inherits all of the properties and methods as described in the HTMLElement interface.")

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

[Categories](dom.html#concept-element-categories){#the-base-element:concept-element-categories}:
:   [Metadata
    content](dom.html#metadata-content-2){#the-base-element:metadata-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-base-element:concept-element-contexts}:
:   In a [`head`{#the-base-element:the-head-element}](#the-head-element)
    element containing no other
    [`base`{#the-base-element:the-base-element}](#the-base-element)
    elements.

[Content model](dom.html#concept-element-content-model){#the-base-element:concept-element-content-model}:
:   [Nothing](dom.html#concept-content-nothing){#the-base-element:concept-content-nothing}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-base-element:concept-element-tag-omission}:
:   No [end
    tag](syntax.html#syntax-end-tag){#the-base-element:syntax-end-tag}.

[Content attributes](dom.html#concept-element-attributes){#the-base-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-base-element:global-attributes}
:   [`href`{#the-base-element:attr-base-href}](#attr-base-href) ---
    [Document base
    URL](urls-and-fetching.html#document-base-url){#the-base-element:document-base-url}
:   [`target`{#the-base-element:attr-base-target}](#attr-base-target)
    --- Default
    [navigable](document-sequences.html#navigable){#the-base-element:navigable}
    for [hyperlink](links.html#hyperlink){#the-base-element:hyperlink}
    [navigation](browsing-the-web.html#navigate){#the-base-element:navigate}
    and [form
    submission](form-control-infrastructure.html#form-submission-2){#the-base-element:form-submission-2}

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-base-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-base).
:   [For implementers](https://w3c.github.io/html-aam/#el-base).

[DOM interface](dom.html#concept-element-dom){#the-base-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLBaseElement : HTMLElement {
      [HTMLConstructor] constructor();

      [CEReactions] attribute USVString href;
      [CEReactions] attribute DOMString target;
    };
    ```

The [`base`{#the-base-element:the-base-element-2}](#the-base-element)
element allows authors to specify the [document base
URL](urls-and-fetching.html#document-base-url){#the-base-element:document-base-url-2}
for the purposes of parsing
[URLs](https://url.spec.whatwg.org/#concept-url){#the-base-element:url
x-internal="url"}, and the name of the default
[navigable](document-sequences.html#navigable){#the-base-element:navigable-2}
for the purposes of [following
hyperlinks](links.html#following-hyperlinks-2){#the-base-element:following-hyperlinks-2}.
The element does not
[represent](dom.html#represents){#the-base-element:represents} any
content beyond this information.

There must be no more than one
[`base`{#the-base-element:the-base-element-3}](#the-base-element)
element per document.

A [`base`{#the-base-element:the-base-element-4}](#the-base-element)
element must have either an
[`href`{#the-base-element:attr-base-href-2}](#attr-base-href) attribute,
a [`target`{#the-base-element:attr-base-target-2}](#attr-base-target)
attribute, or both.

The [`href`]{#attr-base-href .dfn dfn-for="base"
dfn-type="element-attr"} content attribute, if specified, must contain a
[valid URL potentially surrounded by
spaces](urls-and-fetching.html#valid-url-potentially-surrounded-by-spaces){#the-base-element:valid-url-potentially-surrounded-by-spaces}.

A [`base`{#the-base-element:the-base-element-5}](#the-base-element)
element, if it has an
[`href`{#the-base-element:attr-base-href-3}](#attr-base-href) attribute,
must come before any other elements in the tree that have attributes
defined as taking
[URLs](https://url.spec.whatwg.org/#concept-url){#the-base-element:url-2
x-internal="url"}.

If there are multiple
[`base`{#the-base-element:the-base-element-6}](#the-base-element)
elements with
[`href`{#the-base-element:attr-base-href-4}](#attr-base-href)
attributes, all but the first are ignored.

The [`target`]{#attr-base-target .dfn dfn-for="base"
dfn-type="element-attr"} attribute, if specified, must contain a [valid
navigable target name or
keyword](document-sequences.html#valid-navigable-target-name-or-keyword){#the-base-element:valid-navigable-target-name-or-keyword},
which specifies which
[navigable](document-sequences.html#navigable){#the-base-element:navigable-3}
is to be used as the default when
[hyperlinks](links.html#hyperlink){#the-base-element:hyperlink-2} and
[forms](forms.html#the-form-element){#the-base-element:the-form-element}
in the [`Document`{#the-base-element:document}](dom.html#document) cause
[navigation](browsing-the-web.html#navigate){#the-base-element:navigate-2}.

A [`base`{#the-base-element:the-base-element-7}](#the-base-element)
element, if it has a
[`target`{#the-base-element:attr-base-target-3}](#attr-base-target)
attribute, must come before any elements in the tree that represent
[hyperlinks](links.html#hyperlink){#the-base-element:hyperlink-3}.

If there are multiple
[`base`{#the-base-element:the-base-element-8}](#the-base-element)
elements with
[`target`{#the-base-element:attr-base-target-4}](#attr-base-target)
attributes, all but the first are ignored.

To [get an element\'s target]{#get-an-element's-target .dfn}, given an
[`a`{#the-base-element:the-a-element}](text-level-semantics.html#the-a-element),
[`area`{#the-base-element:the-area-element}](image-maps.html#the-area-element),
or
[`form`{#the-base-element:the-form-element-2}](forms.html#the-form-element)
element `element`{.variable}, and an optional string-or-null
`target`{.variable} (default null), run these steps:

1.  If `target`{.variable} is null, then:

    1.  If `element`{.variable} has a `target` attribute, then set
        `target`{.variable} to that attribute\'s value.

    2.  Otherwise, if `element`{.variable}\'s [node
        document](https://dom.spec.whatwg.org/#concept-node-document){#the-base-element:node-document
        x-internal="node-document"} contains a
        [`base`{#the-base-element:the-base-element-9}](#the-base-element)
        element with a
        [`target`{#the-base-element:attr-base-target-5}](#attr-base-target)
        attribute, set `target`{.variable} to the value of the
        [`target`{#the-base-element:attr-base-target-6}](#attr-base-target)
        attribute of the first such
        [`base`{#the-base-element:the-base-element-10}](#the-base-element)
        element.

2.  If `target`{.variable} is not null, and contains an [ASCII tab or
    newline](https://infra.spec.whatwg.org/#ascii-tab-or-newline){#the-base-element:ascii-tab-or-newline
    x-internal="ascii-tab-or-newline"} and a U+003C (\<), then set
    `target`{.variable} to \"`_blank`\".

3.  Return `target`{.variable}.

------------------------------------------------------------------------

A [`base`{#the-base-element:the-base-element-11}](#the-base-element)
element that is the first
[`base`{#the-base-element:the-base-element-12}](#the-base-element)
element with an
[`href`{#the-base-element:attr-base-href-5}](#attr-base-href) content
attribute [in a document
tree](https://dom.spec.whatwg.org/#in-a-document-tree){#the-base-element:in-a-document-tree
x-internal="in-a-document-tree"} has a [frozen base
URL]{#frozen-base-url .dfn}. The [frozen base
URL](#frozen-base-url){#the-base-element:frozen-base-url} must be
[immediately](infrastructure.html#immediately){#the-base-element:immediately}
[set](#set-the-frozen-base-url){#the-base-element:set-the-frozen-base-url}
for an element whenever any of the following situations occur:

- The [`base`{#the-base-element:the-base-element-13}](#the-base-element)
  element becomes the first
  [`base`{#the-base-element:the-base-element-14}](#the-base-element)
  element in [tree
  order](https://dom.spec.whatwg.org/#concept-tree-order){#the-base-element:tree-order
  x-internal="tree-order"} with an
  [`href`{#the-base-element:attr-base-href-6}](#attr-base-href) content
  attribute in its
  [`Document`{#the-base-element:document-2}](dom.html#document).

- The [`base`{#the-base-element:the-base-element-15}](#the-base-element)
  element is the first
  [`base`{#the-base-element:the-base-element-16}](#the-base-element)
  element in [tree
  order](https://dom.spec.whatwg.org/#concept-tree-order){#the-base-element:tree-order-2
  x-internal="tree-order"} with an
  [`href`{#the-base-element:attr-base-href-7}](#attr-base-href) content
  attribute in its
  [`Document`{#the-base-element:document-3}](dom.html#document), and its
  [`href`{#the-base-element:attr-base-href-8}](#attr-base-href) content
  attribute is changed.

To [set the frozen base URL]{#set-the-frozen-base-url .dfn export=""}
for an element `element`{.variable}:

1.  Let `document`{.variable} be `element`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-base-element:node-document-2
    x-internal="node-document"}.

2.  Let `urlRecord`{.variable} be the result of
    [parsing](https://url.spec.whatwg.org/#concept-url-parser){#the-base-element:url-parser
    x-internal="url-parser"} the value of `element`{.variable}\'s
    [`href`{#the-base-element:attr-base-href-9}](#attr-base-href)
    content attribute with `document`{.variable}\'s [fallback base
    URL](urls-and-fetching.html#fallback-base-url){#the-base-element:fallback-base-url},
    and `document`{.variable}\'s [character
    encoding](https://dom.spec.whatwg.org/#concept-document-encoding){#the-base-element:document's-character-encoding
    x-internal="document's-character-encoding"}. (Thus, the
    [`base`{#the-base-element:the-base-element-17}](#the-base-element)
    element isn\'t affected by itself.)

3.  If any of the following are true:

    - `urlRecord`{.variable} is failure;

    - `urlRecord`{.variable}\'s
      [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#the-base-element:concept-url-scheme
      x-internal="concept-url-scheme"} is \"`data`\" or
      \"`javascript`\"; or

    - running [Is base allowed for
      Document?](https://w3c.github.io/webappsec-csp/#allow-base-for-document){#the-base-element:is-base-allowed-for-document
      x-internal="is-base-allowed-for-document"} on
      `urlRecord`{.variable} and `document`{.variable} returns
      \"`Blocked`\",

    then set `element`{.variable}\'s [frozen base
    URL](#frozen-base-url){#the-base-element:frozen-base-url-2} to
    `document`{.variable}\'s [fallback base
    URL](urls-and-fetching.html#fallback-base-url){#the-base-element:fallback-base-url-2}
    and return.

4.  Set `element`{.variable}\'s [frozen base
    URL](#frozen-base-url){#the-base-element:frozen-base-url-3} to
    `urlRecord`{.variable}.

The [`href`]{#dom-base-href .dfn dfn-for="HTMLBaseElement"
dfn-type="attribute"} IDL attribute, on getting, must return the result
of running the following algorithm:

1.  Let `document`{.variable} be `element`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-base-element:node-document-3
    x-internal="node-document"}.

2.  Let `url`{.variable} be the value of the
    [`href`{#the-base-element:attr-base-href-10}](#attr-base-href)
    attribute of this element, if it has one, and the empty string
    otherwise.

3.  Let `urlRecord`{.variable} be the result of
    [parsing](https://url.spec.whatwg.org/#concept-url-parser){#the-base-element:url-parser-2
    x-internal="url-parser"} `url`{.variable} with
    `document`{.variable}\'s [fallback base
    URL](urls-and-fetching.html#fallback-base-url){#the-base-element:fallback-base-url-3},
    and `document`{.variable}\'s [character
    encoding](https://dom.spec.whatwg.org/#concept-document-encoding){#the-base-element:document's-character-encoding-2
    x-internal="document's-character-encoding"}. (Thus, the
    [`base`{#the-base-element:the-base-element-18}](#the-base-element)
    element isn\'t affected by other
    [`base`{#the-base-element:the-base-element-19}](#the-base-element)
    elements or itself.)

4.  If `urlRecord`{.variable} is failure, return `url`{.variable}.

5.  Return the
    [serialization](https://url.spec.whatwg.org/#concept-url-serializer){#the-base-element:concept-url-serializer
    x-internal="concept-url-serializer"} of `urlRecord`{.variable}.

The [`href`{#the-base-element:dom-base-href-2}](#dom-base-href) IDL
attribute, on setting, must set the
[`href`{#the-base-element:attr-base-href-11}](#attr-base-href) content
attribute to the given new value.

The [`target`]{#dom-base-target .dfn dfn-for="HTMLBaseElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-base-element:reflect}
the content attribute of the same name.

::: example
In this example, a
[`base`{#the-base-element:the-base-element-20}](#the-base-element)
element is used to set the [document base
URL](urls-and-fetching.html#document-base-url){#the-base-element:document-base-url-3}:

``` html
<!DOCTYPE html>
<html lang="en">
    <head>
        <title>This is an example for the &lt;base&gt; element</title>
        <base href="https://www.example.com/news/index.html">
    </head>
    <body>
        <p>Visit the <a href="archives.html">archives</a>.</p>
    </body>
</html>
```

The link in the above example would be a link to
\"`https://www.example.com/news/archives.html`\".
:::

#### [4.2.4]{.secno} The [`link`]{.dfn dfn-type="element"} element[](#the-link-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/link](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/link "The <link> HTML element specifies relationships between the current document and an external resource. This element is most commonly used to link to stylesheets, but is also used to establish site icons (both "favicon" style icons and icons for the home screen and apps on mobile devices) among other things.")

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
[HTMLLinkElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLLinkElement "The HTMLLinkElement interface represents reference information for external resources and the relationship of those resources to a document and vice versa (corresponds to <link> element; not to be confused with <a>, which is represented by HTMLAnchorElement). This object inherits all of the properties and methods of the HTMLElement interface.")

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
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

[Categories](dom.html#concept-element-categories){#the-link-element:concept-element-categories}:
:   [Metadata
    content](dom.html#metadata-content-2){#the-link-element:metadata-content-2}.
:   If the element is [allowed in the
    body](#allowed-in-the-body){#the-link-element:allowed-in-the-body}:
    [flow
    content](dom.html#flow-content-2){#the-link-element:flow-content-2}.
:   If the element is [allowed in the
    body](#allowed-in-the-body){#the-link-element:allowed-in-the-body-2}:
    [phrasing
    content](dom.html#phrasing-content-2){#the-link-element:phrasing-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-link-element:concept-element-contexts}:
:   Where [metadata
    content](dom.html#metadata-content-2){#the-link-element:metadata-content-2-2}
    is expected.
:   In a
    [`noscript`{#the-link-element:the-noscript-element}](scripting.html#the-noscript-element)
    element that is a child of a
    [`head`{#the-link-element:the-head-element}](#the-head-element)
    element.
:   If the element is [allowed in the
    body](#allowed-in-the-body){#the-link-element:allowed-in-the-body-3}:
    where [phrasing
    content](dom.html#phrasing-content-2){#the-link-element:phrasing-content-2-2}
    is expected.

[Content model](dom.html#concept-element-content-model){#the-link-element:concept-element-content-model}:
:   [Nothing](dom.html#concept-content-nothing){#the-link-element:concept-content-nothing}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-link-element:concept-element-tag-omission}:
:   No [end
    tag](syntax.html#syntax-end-tag){#the-link-element:syntax-end-tag}.

[Content attributes](dom.html#concept-element-attributes){#the-link-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-link-element:global-attributes}
:   [`href`{#the-link-element:attr-link-href}](#attr-link-href) ---
    Address of the
    [hyperlink](links.html#hyperlink){#the-link-element:hyperlink}
:   [`crossorigin`{#the-link-element:attr-link-crossorigin}](#attr-link-crossorigin)
    --- How the element handles crossorigin requests
:   [`rel`{#the-link-element:attr-link-rel}](#attr-link-rel) ---
    Relationship between the document containing the
    [hyperlink](links.html#hyperlink){#the-link-element:hyperlink-2} and
    the destination resource
:   [`media`{#the-link-element:attr-link-media}](#attr-link-media) ---
    Applicable media
:   [`integrity`{#the-link-element:attr-link-integrity}](#attr-link-integrity)
    --- Integrity metadata used in Subresource Integrity checks
    [\[SRI\]](references.html#refsSRI)
:   [`hreflang`{#the-link-element:attr-link-hreflang}](#attr-link-hreflang)
    --- Language of the linked resource
:   [`type`{#the-link-element:attr-link-type}](#attr-link-type) --- Hint
    for the type of the referenced resource
:   [`referrerpolicy`{#the-link-element:attr-link-referrerpolicy}](#attr-link-referrerpolicy)
    --- [Referrer
    policy](https://w3c.github.io/webappsec-referrer-policy/#referrer-policy){#the-link-element:referrer-policy
    x-internal="referrer-policy"} for
    [fetches](https://fetch.spec.whatwg.org/#concept-fetch){#the-link-element:concept-fetch
    x-internal="concept-fetch"} initiated by the element
:   [`sizes`{#the-link-element:attr-link-sizes}](#attr-link-sizes) ---
    Sizes of the icons (for
    [`rel`{#the-link-element:attr-link-rel-2}](#attr-link-rel)=\"[`icon`{#the-link-element:rel-icon}](links.html#rel-icon)\")
:   [`imagesrcset`{#the-link-element:attr-link-imagesrcset}](#attr-link-imagesrcset)
    --- Images to use in different situations, e.g., high-resolution
    displays, small monitors, etc. (for
    [`rel`{#the-link-element:attr-link-rel-3}](#attr-link-rel)=\"[`preload`{#the-link-element:link-type-preload}](links.html#link-type-preload)\")
:   [`imagesizes`{#the-link-element:attr-link-imagesizes}](#attr-link-imagesizes)
    --- Image sizes for different page layouts (for
    [`rel`{#the-link-element:attr-link-rel-4}](#attr-link-rel)=\"[`preload`{#the-link-element:link-type-preload-2}](links.html#link-type-preload)\")
:   [`as`{#the-link-element:attr-link-as}](#attr-link-as) --- [Potential
    destination](https://fetch.spec.whatwg.org/#concept-potential-destination){#the-link-element:concept-potential-destination
    x-internal="concept-potential-destination"} for a preload request
    (for
    [`rel`{#the-link-element:attr-link-rel-5}](#attr-link-rel)=\"[`preload`{#the-link-element:link-type-preload-3}](links.html#link-type-preload)\"
    and
    [`rel`{#the-link-element:attr-link-rel-6}](#attr-link-rel)=\"[`modulepreload`{#the-link-element:link-type-modulepreload}](links.html#link-type-modulepreload)\")
:   [`blocking`{#the-link-element:attr-link-blocking}](#attr-link-blocking)
    --- Whether the element is [potentially
    render-blocking](urls-and-fetching.html#potentially-render-blocking){#the-link-element:potentially-render-blocking}
:   [`color`{#the-link-element:attr-link-color}](#attr-link-color) ---
    Color to use when customizing a site\'s icon (for
    [`rel`{#the-link-element:attr-link-rel-7}](#attr-link-rel)=\"`mask-icon`\")
:   [`disabled`{#the-link-element:attr-link-disabled}](#attr-link-disabled)
    --- Whether the link is disabled
:   [`fetchpriority`{#the-link-element:attr-link-fetchpriority}](#attr-link-fetchpriority)
    --- Sets the
    [priority](https://fetch.spec.whatwg.org/#request-priority){#the-link-element:concept-request-priority
    x-internal="concept-request-priority"} for
    [fetches](https://fetch.spec.whatwg.org/#concept-fetch){#the-link-element:concept-fetch-2
    x-internal="concept-fetch"} initiated by the element
:   Also, the
    [`title`{#the-link-element:attr-link-title}](#attr-link-title)
    attribute [has special
    semantics](#attr-link-title){#the-link-element:attr-link-title-2} on
    this element: Title of the link; [CSS style sheet set
    name](https://drafts.csswg.org/cssom/#css-style-sheet-set-name){#the-link-element:css-style-sheet-set-name
    x-internal="css-style-sheet-set-name"}

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-link-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-link).
:   [For implementers](https://w3c.github.io/html-aam/#el-link).

[DOM interface](dom.html#concept-element-dom){#the-link-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLLinkElement : HTMLElement {
      [HTMLConstructor] constructor();

      [CEReactions] attribute USVString href;
      [CEReactions] attribute DOMString? crossOrigin;
      [CEReactions] attribute DOMString rel;
      [CEReactions] attribute DOMString as;
      [SameObject, PutForwards=value] readonly attribute DOMTokenList relList;
      [CEReactions] attribute DOMString media;
      [CEReactions] attribute DOMString integrity;
      [CEReactions] attribute DOMString hreflang;
      [CEReactions] attribute DOMString type;
      [SameObject, PutForwards=value] readonly attribute DOMTokenList sizes;
      [CEReactions] attribute USVString imageSrcset;
      [CEReactions] attribute DOMString imageSizes;
      [CEReactions] attribute DOMString referrerPolicy;
      [SameObject, PutForwards=value] readonly attribute DOMTokenList blocking;
      [CEReactions] attribute boolean disabled;
      [CEReactions] attribute DOMString fetchPriority;

      // also has obsolete members
    };
    HTMLLinkElement includes LinkStyle;
    ```

The [`link`{#the-link-element:the-link-element}](#the-link-element)
element allows authors to link their document to other resources.

The address of the link(s) is given by the [`href`]{#attr-link-href .dfn
dfn-for="link" dfn-type="element-attr"} attribute. If the
[`href`{#the-link-element:attr-link-href-2}](#attr-link-href) attribute
is present, then its value must be a [valid non-empty URL potentially
surrounded by
spaces](urls-and-fetching.html#valid-non-empty-url-potentially-surrounded-by-spaces){#the-link-element:valid-non-empty-url-potentially-surrounded-by-spaces}.
One or both of the
[`href`{#the-link-element:attr-link-href-3}](#attr-link-href) or
[`imagesrcset`{#the-link-element:attr-link-imagesrcset-2}](#attr-link-imagesrcset)
attributes must be present.

If both the
[`href`{#the-link-element:attr-link-href-4}](#attr-link-href) and
[`imagesrcset`{#the-link-element:attr-link-imagesrcset-3}](#attr-link-imagesrcset)
attributes are absent, then the element does not define a link.

The types of link indicated (the relationships) are given by the value
of the [`rel`]{#attr-link-rel .dfn dfn-for="link"
dfn-type="element-attr"} attribute, which, if present, must have a value
that is a [unordered set of unique space-separated
tokens](common-microsyntaxes.html#unordered-set-of-unique-space-separated-tokens){#the-link-element:unordered-set-of-unique-space-separated-tokens}.
The [allowed keywords and their meanings](links.html#linkTypes) are
defined in a later section. If the
[`rel`{#the-link-element:attr-link-rel-8}](#attr-link-rel) attribute is
absent, has no keywords, or if none of the keywords used are allowed
according to the definitions in this specification, then the element
does not create any links.

[`rel`{#the-link-element:attr-link-rel-9}](#attr-link-rel)\'s [supported
tokens](https://dom.spec.whatwg.org/#concept-supported-tokens){#the-link-element:concept-supported-tokens
x-internal="concept-supported-tokens"} are the keywords defined in [HTML
link types](links.html#linkTypes) which are allowed on
[`link`{#the-link-element:the-link-element-2}](#the-link-element)
elements, impact the processing model, and are supported by the user
agent. The possible [supported
tokens](https://dom.spec.whatwg.org/#concept-supported-tokens){#the-link-element:concept-supported-tokens-2
x-internal="concept-supported-tokens"} are
[`alternate`{#the-link-element:rel-alternate}](links.html#rel-alternate),
[`dns-prefetch`{#the-link-element:link-type-dns-prefetch}](links.html#link-type-dns-prefetch),
[`expect`{#the-link-element:link-type-expect}](links.html#link-type-expect),
[`icon`{#the-link-element:rel-icon-2}](links.html#rel-icon),
[`manifest`{#the-link-element:link-type-manifest}](links.html#link-type-manifest),
[`modulepreload`{#the-link-element:link-type-modulepreload-2}](links.html#link-type-modulepreload),
[`next`{#the-link-element:link-type-next}](links.html#link-type-next),
[`pingback`{#the-link-element:link-type-pingback}](links.html#link-type-pingback),
[`preconnect`{#the-link-element:link-type-preconnect}](links.html#link-type-preconnect),
[`prefetch`{#the-link-element:link-type-prefetch}](links.html#link-type-prefetch),
[`preload`{#the-link-element:link-type-preload-4}](links.html#link-type-preload),
[`search`{#the-link-element:link-type-search}](links.html#link-type-search),
and
[`stylesheet`{#the-link-element:link-type-stylesheet}](links.html#link-type-stylesheet).
[`rel`{#the-link-element:attr-link-rel-10}](#attr-link-rel)\'s
[supported
tokens](https://dom.spec.whatwg.org/#concept-supported-tokens){#the-link-element:concept-supported-tokens-3
x-internal="concept-supported-tokens"} must only include the tokens from
this list that the user agent implements the processing model for.

Theoretically a user agent could support the processing model for the
[`canonical`{#the-link-element:link-type-canonical}](links.html#link-type-canonical)
keyword --- if it were a search engine that executed JavaScript. But in
practice that\'s quite unlikely. So in most cases,
[`canonical`{#the-link-element:link-type-canonical-2}](links.html#link-type-canonical)
ought not be included in
[`rel`{#the-link-element:attr-link-rel-11}](#attr-link-rel)\'s
[supported
tokens](https://dom.spec.whatwg.org/#concept-supported-tokens){#the-link-element:concept-supported-tokens-4
x-internal="concept-supported-tokens"}.

A [`link`{#the-link-element:the-link-element-3}](#the-link-element)
element must have either a
[`rel`{#the-link-element:attr-link-rel-12}](#attr-link-rel) attribute or
an
[`itemprop`{#the-link-element:names:-the-itemprop-attribute}](microdata.html#names:-the-itemprop-attribute)
attribute, but not both.

If a [`link`{#the-link-element:the-link-element-4}](#the-link-element)
element has an
[`itemprop`{#the-link-element:names:-the-itemprop-attribute-2}](microdata.html#names:-the-itemprop-attribute)
attribute, or has a
[`rel`{#the-link-element:attr-link-rel-13}](#attr-link-rel) attribute
that contains only keywords that are
[body-ok](links.html#body-ok){#the-link-element:body-ok}, then the
element is said to be [allowed in the body]{#allowed-in-the-body .dfn}.
This means that the element can be used where [phrasing
content](dom.html#phrasing-content-2){#the-link-element:phrasing-content-2-3}
is expected.

If the [`rel`{#the-link-element:attr-link-rel-14}](#attr-link-rel)
attribute is used, the element can only sometimes be used in the
[`body`{#the-link-element:the-body-element}](sections.html#the-body-element)
of the page. When used with the
[`itemprop`{#the-link-element:names:-the-itemprop-attribute-3}](microdata.html#names:-the-itemprop-attribute)
attribute, the element can be used both in the
[`head`{#the-link-element:the-head-element-2}](#the-head-element)
element and in the
[`body`{#the-link-element:the-body-element-2}](sections.html#the-body-element)
of the page, subject to the constraints of the microdata model.

------------------------------------------------------------------------

Two categories of links can be created using the
[`link`{#the-link-element:the-link-element-5}](#the-link-element)
element: [links to external
resources](links.html#external-resource-link){#the-link-element:external-resource-link}
and [hyperlinks](links.html#hyperlink){#the-link-element:hyperlink-3}.
The [link types section](links.html#linkTypes) defines whether a
particular link type is an external resource or a hyperlink. One
[`link`{#the-link-element:the-link-element-6}](#the-link-element)
element can create multiple links (of which some might be [external
resource
links](links.html#external-resource-link){#the-link-element:external-resource-link-2}
and some might be
[hyperlinks](links.html#hyperlink){#the-link-element:hyperlink-4});
exactly which and how many links are created depends on the keywords
given in the [`rel`{#the-link-element:attr-link-rel-15}](#attr-link-rel)
attribute. User agents must process the links on a per-link basis, not a
per-element basis.

Each link created for a
[`link`{#the-link-element:the-link-element-7}](#the-link-element)
element is handled separately. For instance, if there are two
[`link`{#the-link-element:the-link-element-8}](#the-link-element)
elements with `rel="stylesheet"`, they each count as a separate external
resource, and each is affected by its own attributes independently.
Similarly, if a single
[`link`{#the-link-element:the-link-element-9}](#the-link-element)
element has a
[`rel`{#the-link-element:attr-link-rel-16}](#attr-link-rel) attribute
with the value `next stylesheet`, it creates both a
[hyperlink](links.html#hyperlink){#the-link-element:hyperlink-5} (for
the
[`next`{#the-link-element:link-type-next-2}](links.html#link-type-next)
keyword) and an [external resource
link](links.html#external-resource-link){#the-link-element:external-resource-link-3}
(for the
[`stylesheet`{#the-link-element:link-type-stylesheet-2}](links.html#link-type-stylesheet)
keyword), and they are affected by other attributes (such as
[`media`{#the-link-element:attr-link-media-2}](#attr-link-media) or
[`title`{#the-link-element:attr-link-title-3}](#attr-link-title))
differently.

::: example
For example, the following
[`link`{#the-link-element:the-link-element-10}](#the-link-element)
element creates two
[hyperlinks](links.html#hyperlink){#the-link-element:hyperlink-6} (to
the same page):

``` html
<link rel="author license" href="/about">
```

The two links created by this element are one whose semantic is that the
target page has information about the current page\'s author, and one
whose semantic is that the target page has information regarding the
license under which the current page is provided.
:::

[Hyperlinks](links.html#hyperlink){#the-link-element:hyperlink-7}
created with the
[`link`{#the-link-element:the-link-element-11}](#the-link-element)
element and its
[`rel`{#the-link-element:attr-link-rel-17}](#attr-link-rel) attribute
apply to the whole document. This contrasts with the
[`rel`{#the-link-element:attr-hyperlink-rel}](links.html#attr-hyperlink-rel)
attribute of
[`a`{#the-link-element:the-a-element}](text-level-semantics.html#the-a-element)
and
[`area`{#the-link-element:the-area-element}](image-maps.html#the-area-element)
elements, which indicates the type of a link whose context is given by
the link\'s location within the document.

Unlike those created by
[`a`{#the-link-element:the-a-element-2}](text-level-semantics.html#the-a-element)
and
[`area`{#the-link-element:the-area-element-2}](image-maps.html#the-area-element)
elements,
[hyperlinks](links.html#hyperlink){#the-link-element:hyperlink-8}
created by
[`link`{#the-link-element:the-link-element-12}](#the-link-element)
elements are not displayed as part of the document by default, in user
agents that [support the suggested default
rendering](infrastructure.html#renderingUA). And even if they are
force-displayed using CSS, they have no [activation
behavior](https://dom.spec.whatwg.org/#eventtarget-activation-behavior){#the-link-element:activation-behaviour
x-internal="activation-behaviour"}. Instead, they primarily provide
semantic information which might be used by the page or by other
software that consumes the page\'s contents. Additionally, the user
agent can [provide its own UI for following such
hyperlinks](#providing-users-with-a-means-to-follow-hyperlinks-created-using-the-link-element).

The exact behavior for [links to external
resources](links.html#external-resource-link){#the-link-element:external-resource-link-4}
depends on the exact relationship, as defined for the relevant [link
type](links.html#linkTypes).

------------------------------------------------------------------------

The [`crossorigin`]{#attr-link-crossorigin .dfn dfn-for="link"
dfn-type="element-attr"} attribute is a [CORS settings
attribute](urls-and-fetching.html#cors-settings-attribute){#the-link-element:cors-settings-attribute}.
It is intended for use with [external resource
links](links.html#external-resource-link){#the-link-element:external-resource-link-5}.

The [`media`]{#attr-link-media .dfn dfn-for="link"
dfn-type="element-attr"} attribute says which media the resource applies
to. The value must be a [valid media query
list](common-microsyntaxes.html#valid-media-query-list){#the-link-element:valid-media-query-list}.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Subresource_Integrity](https://developer.mozilla.org/en-US/docs/Web/Security/Subresource_Integrity "Subresource Integrity (SRI) is a security feature that enables browsers to verify that resources they fetch (for example, from a CDN) are delivered without unexpected manipulation. It works by allowing you to provide a cryptographic hash that a fetched resource must match.")

Support in all current engines.

::: support
[Firefox43+]{.firefox .yes}[Safari11.1+]{.safari
.yes}[Chrome45+]{.chrome .yes}

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

The [`integrity`]{#attr-link-integrity .dfn dfn-for="link"
dfn-type="element-attr"} attribute represents the [integrity
metadata](https://fetch.spec.whatwg.org/#concept-request-integrity-metadata){#the-link-element:concept-request-integrity-metadata
x-internal="concept-request-integrity-metadata"} for requests which this
element is responsible for. The value is text. The attribute must only
be specified on
[`link`{#the-link-element:the-link-element-13}](#the-link-element)
elements that have a
[`rel`{#the-link-element:attr-link-rel-18}](#attr-link-rel) attribute
that contains the
[`stylesheet`{#the-link-element:link-type-stylesheet-3}](links.html#link-type-stylesheet),
[`preload`{#the-link-element:link-type-preload-5}](links.html#link-type-preload),
or
[`modulepreload`{#the-link-element:link-type-modulepreload-3}](links.html#link-type-modulepreload)
keyword. [\[SRI\]](references.html#refsSRI)

The [`hreflang`]{#attr-link-hreflang .dfn dfn-for="link"
dfn-type="element-attr"} attribute on the
[`link`{#the-link-element:the-link-element-14}](#the-link-element)
element has the same semantics as the [`hreflang` attribute on the `a`
element](links.html#attr-hyperlink-hreflang){#the-link-element:attr-hyperlink-hreflang}.

The [`type`]{#attr-link-type .dfn dfn-for="link"
dfn-type="element-attr"} attribute gives the [MIME
type](https://mimesniff.spec.whatwg.org/#mime-type){#the-link-element:mime-type
x-internal="mime-type"} of the linked resource. It is purely advisory.
The value must be a [valid MIME type
string](https://mimesniff.spec.whatwg.org/#valid-mime-type){#the-link-element:valid-mime-type-string
x-internal="valid-mime-type-string"}.

For [external resource
links](links.html#external-resource-link){#the-link-element:external-resource-link-6},
the [`type`{#the-link-element:attr-link-type-2}](#attr-link-type)
attribute is used as a hint to user agents so that they can avoid
fetching resources they do not support.

The [`referrerpolicy`]{#attr-link-referrerpolicy .dfn dfn-for="link"
dfn-type="element-attr"} attribute is a [referrer policy
attribute](urls-and-fetching.html#referrer-policy-attribute){#the-link-element:referrer-policy-attribute}.
It is intended for use with [external resource
links](links.html#external-resource-link){#the-link-element:external-resource-link-7},
where it helps set the [referrer
policy](https://w3c.github.io/webappsec-referrer-policy/#referrer-policy){#the-link-element:referrer-policy-2
x-internal="referrer-policy"} used when [fetching and processing the
linked
resource](#fetch-and-process-the-linked-resource){#the-link-element:fetch-and-process-the-linked-resource}.
[\[REFERRERPOLICY\]](references.html#refsREFERRERPOLICY)

The [`title`]{#attr-link-title .dfn dfn-for="link"
dfn-type="element-attr"} attribute gives the title of the link. With one
exception, it is purely advisory. The value is text. The exception is
for style sheet links that are [in a document
tree](https://dom.spec.whatwg.org/#in-a-document-tree){#the-link-element:in-a-document-tree
x-internal="in-a-document-tree"}, for which the
[`title`{#the-link-element:attr-link-title-4}](#attr-link-title)
attribute defines [CSS style sheet
sets](https://drafts.csswg.org/cssom/#css-style-sheet-set){#the-link-element:css-style-sheet-set
x-internal="css-style-sheet-set"}.

The [`title`{#the-link-element:attr-link-title-5}](#attr-link-title)
attribute on
[`link`{#the-link-element:the-link-element-15}](#the-link-element)
elements differs from the global
[`title`{#the-link-element:attr-title}](dom.html#attr-title) attribute
of most other elements in that a link without a title does not inherit
the title of the parent element: it merely has no title.

------------------------------------------------------------------------

The [`imagesrcset`]{#attr-link-imagesrcset .dfn dfn-for="link"
dfn-type="element-attr"} attribute may be present, and is a [srcset
attribute](images.html#srcset-attribute){#the-link-element:srcset-attribute}.

The
[`imagesrcset`{#the-link-element:attr-link-imagesrcset-4}](#attr-link-imagesrcset)
and [`href`{#the-link-element:attr-link-href-5}](#attr-link-href)
attributes (if [width
descriptors](images.html#width-descriptor){#the-link-element:width-descriptor}
are not used) together contribute the [image
sources](images.html#image-source){#the-link-element:image-source} to
the [source set](images.html#source-set){#the-link-element:source-set}.

If the
[`imagesrcset`{#the-link-element:attr-link-imagesrcset-5}](#attr-link-imagesrcset)
attribute is present and has any [image candidate
strings](images.html#image-candidate-string){#the-link-element:image-candidate-string}
using a [width
descriptor](images.html#width-descriptor){#the-link-element:width-descriptor-2},
the [`imagesizes`]{#attr-link-imagesizes .dfn dfn-for="link"
dfn-type="element-attr"} attribute must also be present, and is a [sizes
attribute](images.html#sizes-attribute){#the-link-element:sizes-attribute}.
The
[`imagesizes`{#the-link-element:attr-link-imagesizes-2}](#attr-link-imagesizes)
attribute contributes the [source
size](images.html#source-size-2){#the-link-element:source-size-2} to the
[source set](images.html#source-set){#the-link-element:source-set-2}.

The
[`imagesrcset`{#the-link-element:attr-link-imagesrcset-6}](#attr-link-imagesrcset)
and
[`imagesizes`{#the-link-element:attr-link-imagesizes-3}](#attr-link-imagesizes)
attributes must only be specified on
[`link`{#the-link-element:the-link-element-16}](#the-link-element)
elements that have both a
[`rel`{#the-link-element:attr-link-rel-19}](#attr-link-rel) attribute
that specifies the
[`preload`{#the-link-element:link-type-preload-6}](links.html#link-type-preload)
keyword, as well as an
[`as`{#the-link-element:attr-link-as-2}](#attr-link-as) attribute in the
\"`image`\" state.

::: example
These attributes allow preloading the appropriate resource that is later
used by an
[`img`{#the-link-element:the-img-element}](embedded-content.html#the-img-element)
element that has the corresponding values for its
[`srcset`{#the-link-element:attr-img-srcset}](embedded-content.html#attr-img-srcset)
and
[`sizes`{#the-link-element:attr-img-sizes}](embedded-content.html#attr-img-sizes)
attributes:

``` html
<link rel="preload" as="image"
      imagesrcset="wolf_400px.jpg 400w, wolf_800px.jpg 800w, wolf_1600px.jpg 1600w"
      imagesizes="50vw">

<!-- ... later, or perhaps inserted dynamically ... -->
<img src="wolf.jpg" alt="A rad wolf"
     srcset="wolf_400px.jpg 400w, wolf_800px.jpg 800w, wolf_1600px.jpg 1600w"
     sizes="50vw">
```

Note how we omit the
[`href`{#the-link-element:attr-link-href-6}](#attr-link-href) attribute,
as it would only be relevant for browsers that do not support
[`imagesrcset`{#the-link-element:attr-link-imagesrcset-7}](#attr-link-imagesrcset),
and in those cases it would likely cause the incorrect image to be
preloaded.
:::

::: example
The
[`imagesrcset`{#the-link-element:attr-link-imagesrcset-8}](#attr-link-imagesrcset)
attribute can be combined with the
[`media`{#the-link-element:attr-link-media-3}](#attr-link-media)
attribute to preload the appropriate resource selected from a
[`picture`{#the-link-element:the-picture-element}](embedded-content.html#the-picture-element)
element\'s sources, for [art
direction](images.html#art-direction){#the-link-element:art-direction}:

``` html
<link rel="preload" as="image"
      imagesrcset="dog-cropped-1x.jpg, dog-cropped-2x.jpg 2x"
      media="(max-width: 800px)">
<link rel="preload" as="image"
      imagesrcset="dog-wide-1x.jpg, dog-wide-2x.jpg 2x"
      media="(min-width: 801px)">

<!-- ... later, or perhaps inserted dynamically ... -->
<picture>
  <source srcset="dog-cropped-1x.jpg, dog-cropped-2x.jpg 2x"
          media="(max-width: 800px)">
  <img src="dog-wide-1x.jpg" srcset="dog-wide-2x.jpg 2x"
       alt="An awesome dog">
</picture>
```
:::

------------------------------------------------------------------------

The [`sizes`]{#attr-link-sizes .dfn dfn-for="link"
dfn-type="element-attr"} attribute gives the sizes of icons for visual
media. Its value, if present, is merely advisory. User agents may use
the value to decide which icon(s) to use if multiple icons are
available. If specified, the attribute must have a value that is an
[unordered set of unique space-separated
tokens](common-microsyntaxes.html#unordered-set-of-unique-space-separated-tokens){#the-link-element:unordered-set-of-unique-space-separated-tokens-2}
which are [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#the-link-element:ascii-case-insensitive
x-internal="ascii-case-insensitive"}. Each value must be either an
[ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#the-link-element:ascii-case-insensitive-2
x-internal="ascii-case-insensitive"} match for the string
\"[`any`{#the-link-element:attr-link-sizes-any}](links.html#attr-link-sizes-any)\",
or a value that consists of two [valid non-negative
integers](common-microsyntaxes.html#valid-non-negative-integer){#the-link-element:valid-non-negative-integer}
that do not have a leading U+0030 DIGIT ZERO (0) character and that are
separated by a single U+0078 LATIN SMALL LETTER X or U+0058 LATIN
CAPITAL LETTER X character. The attribute must only be specified on
[`link`{#the-link-element:the-link-element-17}](#the-link-element)
elements that have a
[`rel`{#the-link-element:attr-link-rel-20}](#attr-link-rel) attribute
that specifies the
[`icon`{#the-link-element:rel-icon-3}](links.html#rel-icon) keyword or
the `apple-touch-icon` keyword.

The `apple-touch-icon` keyword is a registered [extension to the
predefined set of link
types](links.html#concept-rel-extensions){#the-link-element:concept-rel-extensions},
but user agents are not required to support it in any way.

------------------------------------------------------------------------

The [`as`]{#attr-link-as .dfn dfn-for="link" dfn-type="element-attr"}
attribute specifies the [potential
destination](https://fetch.spec.whatwg.org/#concept-potential-destination){#the-link-element:concept-potential-destination-2
x-internal="concept-potential-destination"} for a preload request for
the resource given by the
[`href`{#the-link-element:attr-link-href-7}](#attr-link-href) attribute.
It is an [enumerated
attribute](common-microsyntaxes.html#enumerated-attribute){#the-link-element:enumerated-attribute}.
Each [potential
destination](https://fetch.spec.whatwg.org/#concept-potential-destination){#the-link-element:concept-potential-destination-3
x-internal="concept-potential-destination"} is a keyword for this
attribute, mapping to a state of the same name. The attribute must be
specified on
[`link`{#the-link-element:the-link-element-18}](#the-link-element)
elements that have a
[`rel`{#the-link-element:attr-link-rel-21}](#attr-link-rel) attribute
that contains the
[`preload`{#the-link-element:link-type-preload-7}](links.html#link-type-preload)
keyword. It may be specified on
[`link`{#the-link-element:the-link-element-19}](#the-link-element)
elements that have a
[`rel`{#the-link-element:attr-link-rel-22}](#attr-link-rel) attribute
that contains the
[`modulepreload`{#the-link-element:link-type-modulepreload-4}](links.html#link-type-modulepreload)
keyword; in such cases it must have a value which is a [script-like
destination](https://fetch.spec.whatwg.org/#request-destination-script-like){#the-link-element:concept-script-like-destination
x-internal="concept-script-like-destination"}. For other
[`link`{#the-link-element:the-link-element-20}](#the-link-element)
elements, it must not be specified.

The processing model for how the
[`as`{#the-link-element:attr-link-as-3}](#attr-link-as) attribute is
used is given in an individual link type\'s [fetch and process the
linked
resource](#fetch-and-process-the-linked-resource){#the-link-element:fetch-and-process-the-linked-resource-2}
algorithm.

The attribute does not have a *[missing value
default](common-microsyntaxes.html#missing-value-default)* or *[invalid
value default](common-microsyntaxes.html#invalid-value-default)*,
meaning that invalid or missing values for the attribute map to no
state. This is accounted for in the processing model. For
[`preload`{#the-link-element:link-type-preload-8}](links.html#link-type-preload)
links, both conditions are an error; for
[`modulepreload`{#the-link-element:link-type-modulepreload-5}](links.html#link-type-modulepreload)
links, a missing value will be treated as \"`script`\".

------------------------------------------------------------------------

The [`blocking`]{#attr-link-blocking .dfn dfn-for="link"
dfn-type="element-attr"} attribute is a [blocking
attribute](urls-and-fetching.html#blocking-attribute){#the-link-element:blocking-attribute}.
It is used by link types
[`stylesheet`{#the-link-element:link-type-stylesheet-4}](links.html#link-type-stylesheet)
and
[`expect`{#the-link-element:link-type-expect-2}](links.html#link-type-expect),
and it must only be specified on link elements that have a
[`rel`{#the-link-element:attr-link-rel-23}](#attr-link-rel) attribute
containing those keywords.

------------------------------------------------------------------------

The [`color`]{#attr-link-color .dfn dfn-for="link"
dfn-type="element-attr"} attribute is used with the `mask-icon` link
type. The attribute must only be specified on
[`link`{#the-link-element:the-link-element-21}](#the-link-element)
elements that have a
[`rel`{#the-link-element:attr-link-rel-24}](#attr-link-rel) attribute
that contains the `mask-icon` keyword. The value must be a string that
matches the CSS
[\<color\>](https://drafts.csswg.org/css-color/#typedef-color){#the-link-element:color
x-internal="color"} production, defining a suggested color that user
agents can use to customize the display of the icon that the user sees
when they pin your site.

This specification does not have any user agent requirements for the
[`color`{#the-link-element:attr-link-color-2}](#attr-link-color)
attribute.

The `mask-icon` keyword is a registered [extension to the predefined set
of link
types](links.html#concept-rel-extensions){#the-link-element:concept-rel-extensions-2},
but user agents are not required to support it in any way.

------------------------------------------------------------------------

[`link`{#the-link-element:the-link-element-22}](#the-link-element)
elements have an associated [explicitly enabled]{#explicitly-enabled
.dfn} boolean. It is initially false.

The [`disabled`]{#attr-link-disabled .dfn dfn-for="link"
dfn-type="element-attr"} attribute is a [boolean
attribute](common-microsyntaxes.html#boolean-attribute){#the-link-element:boolean-attribute}
that is used with the
[`stylesheet`{#the-link-element:link-type-stylesheet-5}](links.html#link-type-stylesheet)
link type. The attribute must only be specified on
[`link`{#the-link-element:the-link-element-23}](#the-link-element)
elements that have a
[`rel`{#the-link-element:attr-link-rel-25}](#attr-link-rel) attribute
that contains the
[`stylesheet`{#the-link-element:link-type-stylesheet-6}](links.html#link-type-stylesheet)
keyword.

Whenever the
[`disabled`{#the-link-element:attr-link-disabled-2}](#attr-link-disabled)
attribute is removed, set the
[`link`{#the-link-element:the-link-element-24}](#the-link-element)
element\'s [explicitly
enabled](#explicitly-enabled){#the-link-element:explicitly-enabled}
attribute to true.

::: example
Removing the
[`disabled`{#the-link-element:attr-link-disabled-3}](#attr-link-disabled)
attribute dynamically, e.g., using
`document.querySelector("link").removeAttribute("disabled")`, will fetch
and apply the style sheet:

``` html
<link disabled rel="alternate stylesheet" href="css/pooh">
```
:::

::::: {.mdn-anno .wrapped .before}
MDN

:::: feature
[HTMLLinkElement/fetchPriority](https://developer.mozilla.org/en-US/docs/Web/API/HTMLLinkElement/fetchPriority "The fetchPriority property of the HTMLLinkElement interface represents a hint given to the browser on how it should prioritize the preload of the given resource relative to other resources of the same type.")

::: support
[FirefoxNo]{.firefox .no}[Safari[🔰
preview+]{title="Requires setting a user preference or runtime flag."}]{.safari
.yes}[Chrome102+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge102+]{.edge_blink .yes}

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

The [`fetchpriority`]{#attr-link-fetchpriority .dfn dfn-for="link"
dfn-type="element-attr"} attribute is a [fetch priority
attribute](urls-and-fetching.html#fetch-priority-attribute){#the-link-element:fetch-priority-attribute}
that is intended for use with [external resource
links](links.html#external-resource-link){#the-link-element:external-resource-link-8},
where it is used to set the
[priority](https://fetch.spec.whatwg.org/#request-priority){#the-link-element:concept-request-priority-2
x-internal="concept-request-priority"} used when [fetching and
processing the linked
resource](#fetch-and-process-the-linked-resource){#the-link-element:fetch-and-process-the-linked-resource-3}.

------------------------------------------------------------------------

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLLinkElement/rel](https://developer.mozilla.org/en-US/docs/Web/API/HTMLLinkElement/rel "The HTMLLinkElement.rel property reflects the rel attribute. It is a string containing a space-separated list of link types indicating the relationship between the resource represented by the <link> element and the current document.")

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
:::::

The IDL attributes [`href`]{#dom-link-href .dfn
dfn-for="HTMLLinkElement" dfn-type="attribute"},
[`hreflang`]{#dom-link-hreflang .dfn dfn-for="HTMLLinkElement"
dfn-type="attribute"}, [`integrity`]{#dom-link-integrity .dfn
dfn-for="HTMLLinkElement" dfn-type="attribute"},
[`media`]{#dom-link-media .dfn dfn-for="HTMLLinkElement"
dfn-type="attribute"}, [`rel`]{#dom-link-rel .dfn
dfn-for="HTMLLinkElement" dfn-type="attribute"},
[`sizes`]{#dom-link-sizes .dfn dfn-for="HTMLLinkElement"
dfn-type="attribute"}, [`type`]{#dom-link-type .dfn
dfn-for="HTMLLinkElement" dfn-type="attribute"},
[`blocking`]{#dom-link-blocking .dfn dfn-for="HTMLLinkElement"
dfn-type="attribute"}, and [`disabled`]{#dom-link-disabled .dfn
dfn-for="HTMLLinkElement" dfn-type="attribute"} each must
[reflect](common-dom-interfaces.html#reflect){#the-link-element:reflect}
the respective content attributes of the same name.

There is no reflecting IDL attribute for the
[`color`{#the-link-element:attr-link-color-3}](#attr-link-color)
attribute, but this might be added later.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLLinkElement/as](https://developer.mozilla.org/en-US/docs/Web/API/HTMLLinkElement/as "The as property of the HTMLLinkElement interface returns a string representing the type of content to be preloaded by a link element.")

Support in all current engines.

::: support
[Firefox56+]{.firefox .yes}[Safari10+]{.safari .yes}[Chrome50+]{.chrome
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

The [`as`]{#dom-link-as .dfn dfn-for="HTMLLinkElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-link-element:reflect-2}
the [`as`{#the-link-element:attr-link-as-4}](#attr-link-as) content
attribute, [limited to only known
values](common-dom-interfaces.html#limited-to-only-known-values){#the-link-element:limited-to-only-known-values}.

The [`crossOrigin`]{#dom-link-crossorigin .dfn dfn-for="HTMLLinkElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-link-element:reflect-3}
the
[`crossorigin`{#the-link-element:attr-link-crossorigin-2}](#attr-link-crossorigin)
content attribute, [limited to only known
values](common-dom-interfaces.html#limited-to-only-known-values){#the-link-element:limited-to-only-known-values-2}.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLLinkElement/referrerPolicy](https://developer.mozilla.org/en-US/docs/Web/API/HTMLLinkElement/referrerPolicy "The HTMLLinkElement.referrerPolicy property reflects the HTML referrerpolicy attribute of the <link> element defining which referrer is sent when fetching the resource.")

Support in all current engines.

::: support
[Firefox50+]{.firefox .yes}[Safari14.1+]{.safari
.yes}[Chrome58+]{.chrome .yes}

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

The [`referrerPolicy`]{#dom-link-referrerpolicy .dfn
dfn-for="HTMLLinkElement" dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-link-element:reflect-4}
the
[`referrerpolicy`{#the-link-element:attr-link-referrerpolicy-2}](#attr-link-referrerpolicy)
content attribute, [limited to only known
values](common-dom-interfaces.html#limited-to-only-known-values){#the-link-element:limited-to-only-known-values-3}.

The [`fetchPriority`]{#dom-link-fetchpriority .dfn
dfn-for="HTMLLinkElement" dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-link-element:reflect-5}
the
[`fetchpriority`{#the-link-element:attr-link-fetchpriority-2}](#attr-link-fetchpriority)
content attribute, [limited to only known
values](common-dom-interfaces.html#limited-to-only-known-values){#the-link-element:limited-to-only-known-values-4}.

The [`imageSrcset`]{#dom-link-imagesrcset .dfn dfn-for="HTMLLinkElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-link-element:reflect-6}
the
[`imagesrcset`{#the-link-element:attr-link-imagesrcset-9}](#attr-link-imagesrcset)
content attribute.

The [`imageSizes`]{#dom-link-imagesizes .dfn dfn-for="HTMLLinkElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-link-element:reflect-7}
the
[`imagesizes`{#the-link-element:attr-link-imagesizes-4}](#attr-link-imagesizes)
content attribute.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLLinkElement/relList](https://developer.mozilla.org/en-US/docs/Web/API/HTMLLinkElement/relList "The HTMLLinkElement.relList read-only property reflects the rel attribute. It is a live DOMTokenList containing the set of link types indicating the relationship between the resource represented by the <link> element and the current document.")

Support in all current engines.

::: support
[Firefox30+]{.firefox .yes}[Safari9+]{.safari .yes}[Chrome50+]{.chrome
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

The [`relList`]{#dom-link-rellist .dfn dfn-for="HTMLLinkElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-link-element:reflect-8}
the [`rel`{#the-link-element:attr-link-rel-26}](#attr-link-rel) content
attribute.

The [`relList`{#the-link-element:dom-link-rellist-2}](#dom-link-rellist)
attribute can be used for feature detection, by calling its
[`supports()`{#the-link-element:dom-domtokenlist-supports}](https://dom.spec.whatwg.org/#dom-domtokenlist-supports){x-internal="dom-domtokenlist-supports"}
method to check which [types of links](links.html#linkTypes) are
supported.

##### [4.2.4.1]{.secno} Processing the [`media`{#processing-the-media-attribute:attr-link-media}](#attr-link-media) attribute[](#processing-the-media-attribute){.self-link}

If the link is a
[hyperlink](links.html#hyperlink){#processing-the-media-attribute:hyperlink}
then the
[`media`{#processing-the-media-attribute:attr-link-media-2}](#attr-link-media)
attribute is purely advisory, and describes for which media the document
in question was designed.

However, if the link is an [external resource
link](links.html#external-resource-link){#processing-the-media-attribute:external-resource-link},
then the
[`media`{#processing-the-media-attribute:attr-link-media-3}](#attr-link-media)
attribute is prescriptive. The user agent must apply the external
resource when the
[`media`{#processing-the-media-attribute:attr-link-media-4}](#attr-link-media)
attribute\'s value [matches the
environment](common-microsyntaxes.html#matches-the-environment){#processing-the-media-attribute:matches-the-environment}
and the other relevant conditions apply, and must not apply it
otherwise.

The default, if the
[`media`{#processing-the-media-attribute:attr-link-media-5}](#attr-link-media)
attribute is omitted, is \"`all`\", meaning that by default links apply
to all media.

The external resource might have further restrictions defined within
that limit its applicability. For example, a CSS style sheet might have
some `@media` blocks. This specification does not override such further
restrictions or requirements.

##### [4.2.4.2]{.secno} Processing the [`type`{#processing-the-type-attribute:attr-link-type}](#attr-link-type) attribute[](#processing-the-type-attribute){.self-link}

If the
[`type`{#processing-the-type-attribute:attr-link-type-2}](#attr-link-type)
attribute is present, then the user agent must assume that the resource
is of the given type (even if that is not a [valid MIME type
string](https://mimesniff.spec.whatwg.org/#valid-mime-type){#processing-the-type-attribute:valid-mime-type-string
x-internal="valid-mime-type-string"}, e.g. the empty string). If the
attribute is omitted, but the [external resource
link](links.html#external-resource-link){#processing-the-type-attribute:external-resource-link}
type has a default type defined, then the user agent must assume that
the resource is of that type. If the UA does not support the given [MIME
type](https://mimesniff.spec.whatwg.org/#mime-type){#processing-the-type-attribute:mime-type
x-internal="mime-type"} for the given link relationship, then the UA
should not [fetch and process the linked
resource](#fetch-and-process-the-linked-resource){#processing-the-type-attribute:fetch-and-process-the-linked-resource};
if the UA does support the given [MIME
type](https://mimesniff.spec.whatwg.org/#mime-type){#processing-the-type-attribute:mime-type-2
x-internal="mime-type"} for the given link relationship, then the UA
should [fetch and process the linked
resource](#fetch-and-process-the-linked-resource){#processing-the-type-attribute:fetch-and-process-the-linked-resource-2}
at the appropriate time as specified for the [external resource
link](links.html#external-resource-link){#processing-the-type-attribute:external-resource-link-2}\'s
particular type. If the attribute is omitted, and the [external resource
link](links.html#external-resource-link){#processing-the-type-attribute:external-resource-link-3}
type does not have a default type defined, but the user agent would
[fetch and process the linked
resource](#fetch-and-process-the-linked-resource){#processing-the-type-attribute:fetch-and-process-the-linked-resource-3}
if the type was known and supported, then the user agent should [fetch
and process the linked
resource](#fetch-and-process-the-linked-resource){#processing-the-type-attribute:fetch-and-process-the-linked-resource-4}
under the assumption that it will be supported.

User agents must not consider the
[`type`{#processing-the-type-attribute:attr-link-type-3}](#attr-link-type)
attribute authoritative --- upon fetching the resource, user agents must
not use the
[`type`{#processing-the-type-attribute:attr-link-type-4}](#attr-link-type)
attribute to determine its actual type. Only the actual type (as defined
in the next paragraph) is used to determine whether to *apply* the
resource, not the aforementioned assumed type.

If the [external resource
link](links.html#external-resource-link){#processing-the-type-attribute:external-resource-link-4}
type defines rules for processing the resource\'s [Content-Type
metadata](urls-and-fetching.html#content-type){#processing-the-type-attribute:content-type},
then those rules apply. Otherwise, if the resource is expected to be an
image, user agents may apply the [image sniffing
rules](https://mimesniff.spec.whatwg.org/#rules-for-sniffing-images-specifically){#processing-the-type-attribute:content-type-sniffing:-image
x-internal="content-type-sniffing:-image"}, with the
`official type`{.variable} being the type determined from the
resource\'s [Content-Type
metadata](urls-and-fetching.html#content-type){#processing-the-type-attribute:content-type-2},
and use the resulting [computed type of the
resource](https://mimesniff.spec.whatwg.org/#computed-mime-type){#processing-the-type-attribute:content-type-sniffing-2
x-internal="content-type-sniffing-2"} as if it was the actual type.
Otherwise, if neither of these conditions apply or if the user agent
opts not to apply the image sniffing rules, then the user agent must use
the resource\'s [Content-Type
metadata](urls-and-fetching.html#content-type){#processing-the-type-attribute:content-type-3}
to determine the type of the resource. If there is no type metadata, but
the [external resource
link](links.html#external-resource-link){#processing-the-type-attribute:external-resource-link-5}
type has a default type defined, then the user agent must assume that
the resource is of that type.

The
[`stylesheet`{#processing-the-type-attribute:link-type-stylesheet}](links.html#link-type-stylesheet)
link type defines rules for processing the resource\'s [Content-Type
metadata](urls-and-fetching.html#content-type){#processing-the-type-attribute:content-type-4}.

Once the user agent has established the type of the resource, the user
agent must apply the resource if it is of a supported type and the other
relevant conditions apply, and must ignore the resource otherwise.

::: example
If a document contains style sheet links labeled as follows:

``` html
<link rel="stylesheet" href="A" type="text/plain">
<link rel="stylesheet" href="B" type="text/css">
<link rel="stylesheet" href="C">
```

\...then a compliant UA that supported only CSS style sheets would fetch
the B and C files, and skip the A file (since
[`text/plain`{#processing-the-type-attribute:text/plain}](https://www.rfc-editor.org/rfc/rfc2046#section-4.1.3){x-internal="text/plain"}
is not the [MIME
type](https://mimesniff.spec.whatwg.org/#mime-type){#processing-the-type-attribute:mime-type-3
x-internal="mime-type"} for CSS style sheets).

For files B and C, it would then check the actual types returned by the
server. For those that are sent as
[`text/css`{#processing-the-type-attribute:text/css}](indices.html#text/css),
it would apply the styles, but for those labeled as
[`text/plain`{#processing-the-type-attribute:text/plain-2}](https://www.rfc-editor.org/rfc/rfc2046#section-4.1.3){x-internal="text/plain"},
or any other type, it would not.

If one of the two files was returned without a
[Content-Type](urls-and-fetching.html#content-type){#processing-the-type-attribute:content-type-5}
metadata, or with a syntactically incorrect type like
`Content-Type: "null"`, then the default type for
[`stylesheet`{#processing-the-type-attribute:link-type-stylesheet-2}](links.html#link-type-stylesheet)
links would kick in. Since that default type is
[`text/css`{#processing-the-type-attribute:text/css-2}](indices.html#text/css),
the style sheet *would* nonetheless be applied.
:::

##### [4.2.4.3]{.secno} []{#obtaining-a-resource-from-a-link-element}Fetching and processing a resource from a [`link`{#fetching-and-processing-a-resource-from-a-link-element:the-link-element}](#the-link-element) element[](#fetching-and-processing-a-resource-from-a-link-element){.self-link}

All [external resource
links](links.html#external-resource-link){#fetching-and-processing-a-resource-from-a-link-element:external-resource-link}
have a [fetch and process the linked
resource]{#fetch-and-process-the-linked-resource .dfn} algorithm, which
takes a
[`link`{#fetching-and-processing-a-resource-from-a-link-element:the-link-element-2}](#the-link-element)
element `el`{.variable}. They also have [linked resource fetch setup
steps]{#linked-resource-fetch-setup-steps .dfn} which take a
[`link`{#fetching-and-processing-a-resource-from-a-link-element:the-link-element-3}](#the-link-element)
element `el`{.variable} and
[request](https://fetch.spec.whatwg.org/#concept-request){#fetching-and-processing-a-resource-from-a-link-element:concept-request
x-internal="concept-request"} `request`{.variable}. Individual link
types may provide their own [fetch and process the linked
resource](#fetch-and-process-the-linked-resource){#fetching-and-processing-a-resource-from-a-link-element:fetch-and-process-the-linked-resource}
algorithm, but unless explicitly stated, they use the [default fetch and
process the linked
resource](#default-fetch-and-process-the-linked-resource){#fetching-and-processing-a-resource-from-a-link-element:default-fetch-and-process-the-linked-resource}
algorithm. Similarly, individual link types may provide their own
[linked resource fetch setup
steps](#linked-resource-fetch-setup-steps){#fetching-and-processing-a-resource-from-a-link-element:linked-resource-fetch-setup-steps},
but unless explicitly stated, these steps just return true.

The [default fetch and process the linked
resource]{#default-fetch-and-process-the-linked-resource .dfn}, given a
[`link`{#fetching-and-processing-a-resource-from-a-link-element:the-link-element-4}](#the-link-element)
element `el`{.variable}, is as follows:

1.  Let `options`{.variable} be the result of [creating link
    options](#create-link-options-from-element){#fetching-and-processing-a-resource-from-a-link-element:create-link-options-from-element}
    from `el`{.variable}.

2.  Let `request`{.variable} be the result of [creating a link
    request](#create-a-link-request){#fetching-and-processing-a-resource-from-a-link-element:create-a-link-request}
    given `options`{.variable}.

3.  If `request`{.variable} is null, then return.

4.  Set `request`{.variable}\'s [synchronous
    flag](https://fetch.spec.whatwg.org/#synchronous-flag){#fetching-and-processing-a-resource-from-a-link-element:synchronous-flag
    x-internal="synchronous-flag"}.

5.  Run the [linked resource fetch setup
    steps](#linked-resource-fetch-setup-steps){#fetching-and-processing-a-resource-from-a-link-element:linked-resource-fetch-setup-steps-2},
    given `el`{.variable} and `request`{.variable}. If the result is
    false, then return.

6.  Set `request`{.variable}\'s [initiator
    type](https://fetch.spec.whatwg.org/#request-initiator-type){#fetching-and-processing-a-resource-from-a-link-element:concept-request-initiator-type
    x-internal="concept-request-initiator-type"} to \"`css`\" if
    `el`{.variable}\'s
    [`rel`{#fetching-and-processing-a-resource-from-a-link-element:attr-link-rel}](#attr-link-rel)
    attribute contains the keyword
    [`stylesheet`{#fetching-and-processing-a-resource-from-a-link-element:link-type-stylesheet}](links.html#link-type-stylesheet);
    \"`link`\" otherwise.

7.  [Fetch](https://fetch.spec.whatwg.org/#concept-fetch){#fetching-and-processing-a-resource-from-a-link-element:concept-fetch
    x-internal="concept-fetch"} `request`{.variable} with
    *[processResponseConsumeBody](https://fetch.spec.whatwg.org/#process-response-end-of-body){x-internal="processresponseconsumebody"}*
    set to the following steps given
    [response](https://fetch.spec.whatwg.org/#concept-response){#fetching-and-processing-a-resource-from-a-link-element:concept-response
    x-internal="concept-response"} `response`{.variable} and null,
    failure, or a [byte
    sequence](https://infra.spec.whatwg.org/#byte-sequence){#fetching-and-processing-a-resource-from-a-link-element:byte-sequence
    x-internal="byte-sequence"} `bodyBytes`{.variable}:

    1.  Let `success`{.variable} be true.

    2.  If any of the following are true:

        - `bodyBytes`{.variable} is null or failure; or

        - `response`{.variable}\'s
          [status](https://fetch.spec.whatwg.org/#concept-response-status){#fetching-and-processing-a-resource-from-a-link-element:concept-response-status
          x-internal="concept-response-status"} is not an [ok
          status](https://fetch.spec.whatwg.org/#ok-status){#fetching-and-processing-a-resource-from-a-link-element:ok-status
          x-internal="ok-status"},

        then set `success`{.variable} to false.

        Note that content-specific errors, e.g., CSS parse errors or PNG
        decoding errors, do not affect `success`{.variable}.

    3.  Otherwise, wait for the [link
        resource](links.html#external-resource-link){#fetching-and-processing-a-resource-from-a-link-element:external-resource-link-2}\'s
        [critical
        subresources](infrastructure.html#critical-subresources){#fetching-and-processing-a-resource-from-a-link-element:critical-subresources}
        to finish loading.

        The specification that defines a link type\'s [critical
        subresources](infrastructure.html#critical-subresources){#fetching-and-processing-a-resource-from-a-link-element:critical-subresources-2}
        (e.g., CSS) is expected to describe how these subresources are
        fetched and processed. However, since this is not currently
        explicit, this specification describes waiting for a [link
        resource](links.html#external-resource-link){#fetching-and-processing-a-resource-from-a-link-element:external-resource-link-3}\'s
        [critical
        subresources](infrastructure.html#critical-subresources){#fetching-and-processing-a-resource-from-a-link-element:critical-subresources-3}
        to be fetched and processed, with the expectation that this will
        be done correctly.

    4.  [Process the linked
        resource](#process-the-linked-resource){#fetching-and-processing-a-resource-from-a-link-element:process-the-linked-resource}
        given `el`{.variable}, `success`{.variable},
        `response`{.variable}, and `bodyBytes`{.variable}.

To [create a link request]{#create-a-link-request .dfn} given a [link
processing
options](#link-processing-options){#fetching-and-processing-a-resource-from-a-link-element:link-processing-options}
`options`{.variable}:

1.  [Assert](https://infra.spec.whatwg.org/#assert){#fetching-and-processing-a-resource-from-a-link-element:assert
    x-internal="assert"}: `options`{.variable}\'s
    [href](#link-options-href){#fetching-and-processing-a-resource-from-a-link-element:link-options-href}
    is not the empty string.

2.  If `options`{.variable}\'s
    [destination](#link-options-destination){#fetching-and-processing-a-resource-from-a-link-element:link-options-destination}
    is null, then return null.

3.  Let `url`{.variable} be the result of [encoding-parsing a
    URL](urls-and-fetching.html#encoding-parsing-a-url){#fetching-and-processing-a-resource-from-a-link-element:encoding-parsing-a-url}
    given `options`{.variable}\'s
    [href](#link-options-href){#fetching-and-processing-a-resource-from-a-link-element:link-options-href-2},
    relative to `options`{.variable}\'s [base
    URL](#link-options-base-url){#fetching-and-processing-a-resource-from-a-link-element:link-options-base-url}.

    Passing the base URL instead of a document or environment is tracked
    by [issue #9715](https://github.com/whatwg/html/issues/9715).

4.  If `url`{.variable} is failure, then return null.

5.  Let `request`{.variable} be the result of [creating a potential-CORS
    request](urls-and-fetching.html#create-a-potential-cors-request){#fetching-and-processing-a-resource-from-a-link-element:create-a-potential-cors-request}
    given `url`{.variable}, `options`{.variable}\'s
    [destination](#link-options-destination){#fetching-and-processing-a-resource-from-a-link-element:link-options-destination-2},
    and `options`{.variable}\'s
    [crossorigin](#link-options-crossorigin){#fetching-and-processing-a-resource-from-a-link-element:link-options-crossorigin}.

6.  Set `request`{.variable}\'s [policy
    container](https://fetch.spec.whatwg.org/#concept-request-policy-container){#fetching-and-processing-a-resource-from-a-link-element:concept-request-policy-container
    x-internal="concept-request-policy-container"} to
    `options`{.variable}\'s [policy
    container](#link-options-policy-container){#fetching-and-processing-a-resource-from-a-link-element:link-options-policy-container}.

7.  Set `request`{.variable}\'s [integrity
    metadata](https://fetch.spec.whatwg.org/#concept-request-integrity-metadata){#fetching-and-processing-a-resource-from-a-link-element:concept-request-integrity-metadata
    x-internal="concept-request-integrity-metadata"} to
    `options`{.variable}\'s
    [integrity](#link-options-integrity){#fetching-and-processing-a-resource-from-a-link-element:link-options-integrity}.

8.  Set `request`{.variable}\'s [cryptographic nonce
    metadata](https://fetch.spec.whatwg.org/#concept-request-nonce-metadata){#fetching-and-processing-a-resource-from-a-link-element:concept-request-nonce-metadata
    x-internal="concept-request-nonce-metadata"} to
    `options`{.variable}\'s [cryptographic nonce
    metadata](#link-options-nonce){#fetching-and-processing-a-resource-from-a-link-element:link-options-nonce}.

9.  Set `request`{.variable}\'s [referrer
    policy](https://fetch.spec.whatwg.org/#concept-request-referrer-policy){#fetching-and-processing-a-resource-from-a-link-element:concept-request-referrer-policy
    x-internal="concept-request-referrer-policy"} to
    `options`{.variable}\'s [referrer
    policy](#link-options-referrer-policy){#fetching-and-processing-a-resource-from-a-link-element:link-options-referrer-policy}.

10. Set `request`{.variable}\'s
    [client](https://fetch.spec.whatwg.org/#concept-request-client){#fetching-and-processing-a-resource-from-a-link-element:concept-request-client
    x-internal="concept-request-client"} to `options`{.variable}\'s
    [environment](#link-options-environment){#fetching-and-processing-a-resource-from-a-link-element:link-options-environment}.

11. Set `request`{.variable}\'s
    [priority](https://fetch.spec.whatwg.org/#request-priority){#fetching-and-processing-a-resource-from-a-link-element:concept-request-priority
    x-internal="concept-request-priority"} to `options`{.variable}\'s
    [fetch
    priority](#link-options-fetch-priority){#fetching-and-processing-a-resource-from-a-link-element:link-options-fetch-priority}.

12. Return `request`{.variable}.

User agents may opt to only try to [fetch and
process](#fetch-and-process-the-linked-resource){#fetching-and-processing-a-resource-from-a-link-element:fetch-and-process-the-linked-resource-2}
such resources when they are needed, instead of pro-actively fetching
all the [external
resources](links.html#external-resource-link){#fetching-and-processing-a-resource-from-a-link-element:external-resource-link-4}
that are not applied.

Similar to the [fetch and process the linked
resource](#fetch-and-process-the-linked-resource){#fetching-and-processing-a-resource-from-a-link-element:fetch-and-process-the-linked-resource-3}
algorithm, all [external resource
links](links.html#external-resource-link){#fetching-and-processing-a-resource-from-a-link-element:external-resource-link-5}
have a [process the linked resource]{#process-the-linked-resource .dfn}
algorithm which takes a
[`link`{#fetching-and-processing-a-resource-from-a-link-element:the-link-element-5}](#the-link-element)
element `el`{.variable}, boolean `success`{.variable}, a
[response](https://fetch.spec.whatwg.org/#concept-response){#fetching-and-processing-a-resource-from-a-link-element:concept-response-2
x-internal="concept-response"} `response`{.variable}, and a [byte
sequence](https://infra.spec.whatwg.org/#byte-sequence){#fetching-and-processing-a-resource-from-a-link-element:byte-sequence-2
x-internal="byte-sequence"} `bodyBytes`{.variable}. Individual link
types may provide their own [process the linked
resource](#process-the-linked-resource){#fetching-and-processing-a-resource-from-a-link-element:process-the-linked-resource-2}
algorithm, but unless explicitly stated, that algorithm does nothing.

Unless otherwise specified for a given
[`rel`{#fetching-and-processing-a-resource-from-a-link-element:attr-link-rel-2}](#attr-link-rel)
keyword, the element must [delay the load
event](parsing.html#delay-the-load-event){#fetching-and-processing-a-resource-from-a-link-element:delay-the-load-event}
of the element\'s [node
document](https://dom.spec.whatwg.org/#concept-node-document){#fetching-and-processing-a-resource-from-a-link-element:node-document
x-internal="node-document"} until all the attempts to [fetch and process
the linked
resource](#fetch-and-process-the-linked-resource){#fetching-and-processing-a-resource-from-a-link-element:fetch-and-process-the-linked-resource-4}
and its [critical
subresources](infrastructure.html#critical-subresources){#fetching-and-processing-a-resource-from-a-link-element:critical-subresources-4}
are complete. (Resources that the user agent has not yet attempted to
fetch and process, e.g., because it is waiting for the resource to be
needed, do not [delay the load
event](parsing.html#delay-the-load-event){#fetching-and-processing-a-resource-from-a-link-element:delay-the-load-event-2}.)

##### [4.2.4.4]{.secno} Processing \`[`Link`{#processing-link-headers:http-link}](https://httpwg.org/specs/rfc8288.html#header){x-internal="http-link"}\` headers[](#processing-link-headers){.self-link}

All link types that can be [external resource
links](links.html#external-resource-link){#processing-link-headers:external-resource-link}
define a [process a link header]{#process-a-link-header .dfn} algorithm,
which takes a [link processing
options](#link-processing-options){#processing-link-headers:link-processing-options}.
This algorithm defines whether and how they react to appearing in an
HTTP
\`[`Link`{#processing-link-headers:http-link-2}](https://httpwg.org/specs/rfc8288.html#header){x-internal="http-link"}\`
response header.

For most link types, this algorithm does nothing. The [summary
table](links.html#table-link-relations) is a good reference to quickly
know whether a link type has defined [process a link
header](#process-a-link-header){#processing-link-headers:process-a-link-header}
steps.

A [link processing options]{#link-processing-options .dfn} is a
[struct](https://infra.spec.whatwg.org/#struct){#processing-link-headers:struct
x-internal="struct"}. It has the following
[items](https://infra.spec.whatwg.org/#struct-item){#processing-link-headers:struct-item
x-internal="struct-item"}:

[href]{#link-options-href .dfn} (default the empty string)\
[destination]{#link-options-destination .dfn} (default the empty string)\
[initiator]{#link-options-initiator .dfn} (default \"`link`\")\
[integrity]{#link-options-integrity .dfn} (default the empty string)\
[type]{#link-options-type .dfn} (default the empty string)\
[cryptographic nonce metadata]{#link-options-nonce .dfn} (default the empty string)
:   A string

[crossorigin]{#link-options-crossorigin .dfn} (default [No CORS](urls-and-fetching.html#attr-crossorigin-none){#processing-link-headers:attr-crossorigin-none})
:   A [CORS settings
    attribute](urls-and-fetching.html#cors-settings-attribute){#processing-link-headers:cors-settings-attribute}
    state

[referrer policy]{#link-options-referrer-policy .dfn} (default the empty string)
:   A [referrer
    policy](https://w3c.github.io/webappsec-referrer-policy/#referrer-policy){#processing-link-headers:referrer-policy
    x-internal="referrer-policy"}

[source set]{#link-options-source-set .dfn} (default null)
:   Null or a [source
    set](images.html#source-set){#processing-link-headers:source-set}

[base URL]{#link-options-base-url .dfn}
:   A
    [URL](https://url.spec.whatwg.org/#concept-url){#processing-link-headers:url
    x-internal="url"}

[origin]{#link-options-origin .dfn}
:   An
    [origin](browsers.html#concept-origin){#processing-link-headers:concept-origin}

[environment]{#link-options-environment .dfn}
:   An
    [environment](webappapis.html#environment){#processing-link-headers:environment}

[policy container]{#link-options-policy-container .dfn}
:   A [policy
    container](browsers.html#policy-container){#processing-link-headers:policy-container}

[document]{#link-options-document .dfn} (default null)
:   Null or a
    [`Document`{#processing-link-headers:document}](dom.html#document)

[on document ready]{#link-options-on-document-ready .dfn} (default null)
:   Null or an algorithm accepting a
    [`Document`{#processing-link-headers:document-2}](dom.html#document)

[fetch priority]{#link-options-fetch-priority .dfn} (default [`Auto`{#processing-link-headers:attr-fetchpriority-auto-state}](urls-and-fetching.html#attr-fetchpriority-auto-state))
:   A [fetch priority
    attribute](urls-and-fetching.html#fetch-priority-attribute){#processing-link-headers:fetch-priority-attribute}
    state

A [link processing
options](#link-processing-options){#processing-link-headers:link-processing-options-2}
has a [base
URL](#link-options-base-url){#processing-link-headers:link-options-base-url}
and an
[href](#link-options-href){#processing-link-headers:link-options-href}
rather than a parsed URL because the URL could be a result of the
options\'s [source
set](#link-options-source-set){#processing-link-headers:link-options-source-set}.

To [create link options from element]{#create-link-options-from-element
.dfn} given a
[`link`{#processing-link-headers:the-link-element}](#the-link-element)
element `el`{.variable}:

1.  Let `document`{.variable} be `el`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#processing-link-headers:node-document
    x-internal="node-document"}.

2.  Let `options`{.variable} be a new [link processing
    options](#link-processing-options){#processing-link-headers:link-processing-options-3}
    with

    [destination](#link-options-destination){#processing-link-headers:link-options-destination}
    :   the result of
        [translating](links.html#translate-a-preload-destination){#processing-link-headers:translate-a-preload-destination}
        the state of `el`{.variable}\'s
        [`as`{#processing-link-headers:attr-link-as}](#attr-link-as)
        attribute

    [crossorigin](#link-options-crossorigin){#processing-link-headers:link-options-crossorigin}
    :   the state of `el`{.variable}\'s
        [`crossorigin`{#processing-link-headers:attr-link-crossorigin}](#attr-link-crossorigin)
        content attribute

    [referrer policy](#link-options-referrer-policy){#processing-link-headers:link-options-referrer-policy}
    :   the state of `el`{.variable}\'s
        [`referrerpolicy`{#processing-link-headers:attr-link-referrerpolicy}](#attr-link-referrerpolicy)
        content attribute

    [source set](#link-options-source-set){#processing-link-headers:link-options-source-set-2}
    :   `el`{.variable}\'s [source
        set](images.html#source-set){#processing-link-headers:source-set-2}

    [base URL](#link-options-base-url){#processing-link-headers:link-options-base-url-2}
    :   `document`{.variable}\'s [document base
        URL](urls-and-fetching.html#document-base-url){#processing-link-headers:document-base-url}

    [origin](#link-options-origin){#processing-link-headers:link-options-origin}
    :   `document`{.variable}\'s
        [origin](https://dom.spec.whatwg.org/#concept-document-origin){#processing-link-headers:concept-document-origin
        x-internal="concept-document-origin"}

    [environment](#link-options-environment){#processing-link-headers:link-options-environment}
    :   `document`{.variable}\'s [relevant settings
        object](webappapis.html#relevant-settings-object){#processing-link-headers:relevant-settings-object}

    [policy container](#link-options-policy-container){#processing-link-headers:link-options-policy-container}
    :   `document`{.variable}\'s [policy
        container](dom.html#concept-document-policy-container){#processing-link-headers:concept-document-policy-container}

    [document](#link-options-document){#processing-link-headers:link-options-document}
    :   `document`{.variable}

    [cryptographic nonce metadata](#link-options-nonce){#processing-link-headers:link-options-nonce}
    :   the current value of `el`{.variable}\'s
        [\[\[CryptographicNonce\]\]](urls-and-fetching.html#cryptographicnonce){#processing-link-headers:cryptographicnonce}
        internal slot

    [fetch priority](#link-options-fetch-priority){#processing-link-headers:link-options-fetch-priority}
    :   the state of `el`{.variable}\'s
        [`fetchpriority`{#processing-link-headers:attr-link-fetchpriority}](#attr-link-fetchpriority)
        content attribute

3.  If `el`{.variable} has an
    [`href`{#processing-link-headers:attr-link-href}](#attr-link-href)
    attribute, then set `options`{.variable}\'s
    [href](#link-options-href){#processing-link-headers:link-options-href-2}
    to the value of `el`{.variable}\'s
    [`href`{#processing-link-headers:attr-link-href-2}](#attr-link-href)
    attribute.

4.  If `el`{.variable} has an
    [`integrity`{#processing-link-headers:attr-link-integrity}](#attr-link-integrity)
    attribute, then set `options`{.variable}\'s
    [integrity](#link-options-integrity){#processing-link-headers:link-options-integrity}
    to the value of `el`{.variable}\'s
    [`integrity`{#processing-link-headers:attr-link-integrity-2}](#attr-link-integrity)
    content attribute.

5.  If `el`{.variable} has a
    [`type`{#processing-link-headers:attr-link-type}](#attr-link-type)
    attribute, then set `options`{.variable}\'s
    [type](#link-options-type){#processing-link-headers:link-options-type}
    to the value of `el`{.variable}\'s
    [`type`{#processing-link-headers:attr-link-type-2}](#attr-link-type)
    attribute.

6.  [Assert](https://infra.spec.whatwg.org/#assert){#processing-link-headers:assert
    x-internal="assert"}: `options`{.variable}\'s
    [href](#link-options-href){#processing-link-headers:link-options-href-3}
    is not the empty string, or `options`{.variable}\'s [source
    set](#link-options-source-set){#processing-link-headers:link-options-source-set-3}
    is not null.

    A
    [`link`{#processing-link-headers:the-link-element-2}](#the-link-element)
    element with neither an
    [`href`{#processing-link-headers:attr-link-href-3}](#attr-link-href)
    or an
    [`imagesrcset`{#processing-link-headers:attr-link-imagesrcset}](#attr-link-imagesrcset)
    does not represent a link.

7.  Return `options`{.variable}.

To [extract links from headers]{#extract-links-from-headers .dfn} given
a [header
list](https://fetch.spec.whatwg.org/#concept-header-list){#processing-link-headers:concept-header-list
x-internal="concept-header-list"} `headers`{.variable}:

1.  Let `links`{.variable} be a new
    [list](https://infra.spec.whatwg.org/#list){#processing-link-headers:list
    x-internal="list"}.

2.  Let `rawLinkHeaders`{.variable} be the result of [getting, decoding,
    and
    splitting](https://fetch.spec.whatwg.org/#concept-header-list-get-decode-split){#processing-link-headers:concept-header-list-get-decode-split
    x-internal="concept-header-list-get-decode-split"}
    \`[`Link`{#processing-link-headers:the-link-element-3}](#the-link-element)\`
    from `headers`{.variable}.

3.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#processing-link-headers:list-iterate
    x-internal="list-iterate"} `linkHeader`{.variable} of
    `rawLinkHeaders`{.variable}:

    1.  Let `linkObject`{.variable} be the result of
        [parsing](https://httpwg.org/specs/rfc8288.html#parse-fv){#processing-link-headers:parsing-a-link-field-value
        x-internal="parsing-a-link-field-value"}
        `linkHeader`{.variable}.
        [\[WEBLINK\]](references.html#refsWEBLINK)

    2.  If `linkObject`{.variable}\[\"`target_uri`\"\] does not
        [exist](https://infra.spec.whatwg.org/#map-exists){#processing-link-headers:map-exists
        x-internal="map-exists"}, then
        [continue](https://infra.spec.whatwg.org/#iteration-continue){#processing-link-headers:continue
        x-internal="continue"}.

    3.  [Append](https://infra.spec.whatwg.org/#list-append){#processing-link-headers:list-append
        x-internal="list-append"} `linkObject`{.variable} to
        `links`{.variable}.

4.  Return `links`{.variable}.

To [process link headers]{#process-link-headers .dfn} given a
[`Document`{#processing-link-headers:document-3}](dom.html#document)
`doc`{.variable}, a
[response](https://fetch.spec.whatwg.org/#concept-response){#processing-link-headers:concept-response
x-internal="concept-response"} `response`{.variable}, and a
\"`pre-media`\" or \"`media`\" `phase`{.variable}:

1.  Let `links`{.variable} be the result of [extracting
    links](#extract-links-from-headers){#processing-link-headers:extract-links-from-headers}
    from `response`{.variable}\'s [header
    list](https://fetch.spec.whatwg.org/#concept-response-header-list){#processing-link-headers:concept-response-header-list
    x-internal="concept-response-header-list"}.

2.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#processing-link-headers:list-iterate-2
    x-internal="list-iterate"} `linkObject`{.variable} in
    `links`{.variable}:

    1.  Let `rel`{.variable} be
        `linkObject`{.variable}\[\"`relation_type`\"\].

    2.  Let `attribs`{.variable} be
        `linkObject`{.variable}\[\"`target_attributes`\"\].

    3.  Let `expectedPhase`{.variable} be \"`media`\" if either
        \"[`srcset`{#processing-link-headers:attr-img-srcset}](embedded-content.html#attr-img-srcset)\",
        \"[`imagesrcset`{#processing-link-headers:attr-link-imagesrcset-2}](#attr-link-imagesrcset)\",
        or
        \"[`media`{#processing-link-headers:attr-link-media}](#attr-link-media)\"
        [exist](https://infra.spec.whatwg.org/#map-exists){#processing-link-headers:map-exists-2
        x-internal="map-exists"} in `attribs`{.variable}; otherwise
        \"`pre-media`\".

    4.  If `expectedPhase`{.variable} is not `phase`{.variable}, then
        [continue](https://infra.spec.whatwg.org/#iteration-continue){#processing-link-headers:continue-2
        x-internal="continue"}.

    5.  If
        `attribs`{.variable}\[\"[`media`{#processing-link-headers:attr-link-media-2}](#attr-link-media)\"\]
        [exists](https://infra.spec.whatwg.org/#map-exists){#processing-link-headers:map-exists-3
        x-internal="map-exists"} and
        `attribs`{.variable}\[\"[`media`{#processing-link-headers:attr-link-media-3}](#attr-link-media)\"\]
        does not [match the
        environment](common-microsyntaxes.html#matches-the-environment){#processing-link-headers:matches-the-environment},
        then
        [continue](https://infra.spec.whatwg.org/#iteration-continue){#processing-link-headers:continue-3
        x-internal="continue"}.

    6.  Let `options`{.variable} be a new [link processing
        options](#link-processing-options){#processing-link-headers:link-processing-options-4}
        with

        [href](#link-options-href){#processing-link-headers:link-options-href-4}
        :   `linkObject`{.variable}\[\"`target_uri`\"\]

        [base URL](#link-options-base-url){#processing-link-headers:link-options-base-url-3}
        :   `doc`{.variable}\'s [document base
            URL](urls-and-fetching.html#document-base-url){#processing-link-headers:document-base-url-2}

        [origin](#link-options-origin){#processing-link-headers:link-options-origin-2}
        :   `doc`{.variable}\'s
            [origin](https://dom.spec.whatwg.org/#concept-document-origin){#processing-link-headers:concept-document-origin-2
            x-internal="concept-document-origin"}

        [environment](#link-options-environment){#processing-link-headers:link-options-environment-2}
        :   `doc`{.variable}\'s [relevant settings
            object](webappapis.html#relevant-settings-object){#processing-link-headers:relevant-settings-object-2}

        [policy container](#link-options-policy-container){#processing-link-headers:link-options-policy-container-2}
        :   `doc`{.variable}\'s [policy
            container](dom.html#concept-document-policy-container){#processing-link-headers:concept-document-policy-container-2}

        [document](#link-options-document){#processing-link-headers:link-options-document-2}
        :   `doc`{.variable}

    7.  [Apply link options from parsed header
        attributes](#apply-link-options-from-parsed-header-attributes){#processing-link-headers:apply-link-options-from-parsed-header-attributes}
        to `options`{.variable} given `attribs`{.variable}.

    8.  If
        `attribs`{.variable}\[\"[`imagesrcset`{#processing-link-headers:attr-link-imagesrcset-3}](#attr-link-imagesrcset)\"\]
        [exists](https://infra.spec.whatwg.org/#map-exists){#processing-link-headers:map-exists-4
        x-internal="map-exists"} and
        `attribs`{.variable}\[\"[`imagesizes`{#processing-link-headers:attr-link-imagesizes}](#attr-link-imagesizes)\"\]
        [exists](https://infra.spec.whatwg.org/#map-exists){#processing-link-headers:map-exists-5
        x-internal="map-exists"}, then set `options`{.variable}\'s
        [source
        set](#link-options-source-set){#processing-link-headers:link-options-source-set-4}
        to the result of [creating a source
        set](images.html#create-a-source-set){#processing-link-headers:create-a-source-set}
        given `linkObject`{.variable}\[\"`target_uri`\"\],
        `attribs`{.variable}\[\"[`imagesrcset`{#processing-link-headers:attr-link-imagesrcset-4}](#attr-link-imagesrcset)\"\],
        `attribs`{.variable}\[\"[`imagesizes`{#processing-link-headers:attr-link-imagesizes-2}](#attr-link-imagesizes)\"\],
        and null.

    9.  Run the [process a link
        header](#process-a-link-header){#processing-link-headers:process-a-link-header-2}
        steps for `rel`{.variable} given `options`{.variable}.

To [apply link options from parsed header
attributes]{#apply-link-options-from-parsed-header-attributes .dfn} to a
[link processing
options](#link-processing-options){#processing-link-headers:link-processing-options-5}
`options`{.variable} given `attribs`{.variable}:

1.  If
    `attribs`{.variable}\[\"[`as`{#processing-link-headers:attr-link-as-2}](#attr-link-as)\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#processing-link-headers:map-exists-6
    x-internal="map-exists"}, then set `options`{.variable}\'s
    [destination](#link-options-destination){#processing-link-headers:link-options-destination-2}
    to the result of
    [translating](links.html#translate-a-preload-destination){#processing-link-headers:translate-a-preload-destination-2}
    `attribs`{.variable}\[\"[`as`{#processing-link-headers:attr-link-as-3}](#attr-link-as)\"\].

2.  If
    `attribs`{.variable}\[\"[`crossorigin`{#processing-link-headers:attr-link-crossorigin-2}](#attr-link-crossorigin)\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#processing-link-headers:map-exists-7
    x-internal="map-exists"} and is an [ASCII
    case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#processing-link-headers:ascii-case-insensitive
    x-internal="ascii-case-insensitive"} match for one of the [CORS
    settings
    attribute](urls-and-fetching.html#cors-settings-attribute){#processing-link-headers:cors-settings-attribute-2}
    [keywords](common-microsyntaxes.html#enumerated-attribute){#processing-link-headers:enumerated-attribute},
    then set `options`{.variable}\'s
    [crossorigin](#link-options-crossorigin){#processing-link-headers:link-options-crossorigin-2}
    to the [CORS settings
    attribute](urls-and-fetching.html#cors-settings-attribute){#processing-link-headers:cors-settings-attribute-3}
    state corresponding to that keyword.

3.  If
    `attribs`{.variable}\[\"[`integrity`{#processing-link-headers:attr-link-integrity-3}](#attr-link-integrity)\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#processing-link-headers:map-exists-8
    x-internal="map-exists"}, then set `options`{.variable}\'s
    [integrity](#link-options-integrity){#processing-link-headers:link-options-integrity-2}
    to
    `attribs`{.variable}\[\"[`integrity`{#processing-link-headers:attr-link-integrity-4}](#attr-link-integrity)\"\].

4.  If
    `attribs`{.variable}\[\"[`referrerpolicy`{#processing-link-headers:attr-link-referrerpolicy-2}](#attr-link-referrerpolicy)\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#processing-link-headers:map-exists-9
    x-internal="map-exists"} and is an [ASCII
    case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#processing-link-headers:ascii-case-insensitive-2
    x-internal="ascii-case-insensitive"} match for some [referrer
    policy](https://w3c.github.io/webappsec-referrer-policy/#referrer-policy){#processing-link-headers:referrer-policy-2
    x-internal="referrer-policy"}, then set `options`{.variable}\'s
    [referrer
    policy](#link-options-referrer-policy){#processing-link-headers:link-options-referrer-policy-2}
    to that [referrer
    policy](https://w3c.github.io/webappsec-referrer-policy/#referrer-policy){#processing-link-headers:referrer-policy-3
    x-internal="referrer-policy"}.

5.  If
    `attribs`{.variable}\[\"[`nonce`{#processing-link-headers:attr-nonce}](urls-and-fetching.html#attr-nonce)\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#processing-link-headers:map-exists-10
    x-internal="map-exists"}, then set `options`{.variable}\'s
    [nonce](#link-options-nonce){#processing-link-headers:link-options-nonce-2}
    to
    `attribs`{.variable}\[\"[`nonce`{#processing-link-headers:attr-nonce-2}](urls-and-fetching.html#attr-nonce)\"\].

6.  If
    `attribs`{.variable}\[\"[`type`{#processing-link-headers:attr-link-type-3}](#attr-link-type)\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#processing-link-headers:map-exists-11
    x-internal="map-exists"}, then set `options`{.variable}\'s
    [type](#link-options-type){#processing-link-headers:link-options-type-2}
    to
    `attribs`{.variable}\[\"[`type`{#processing-link-headers:attr-link-type-4}](#attr-link-type)\"\].

7.  If
    `attribs`{.variable}\[\"[`fetchpriority`{#processing-link-headers:attr-link-fetchpriority-2}](#attr-link-fetchpriority)\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#processing-link-headers:map-exists-12
    x-internal="map-exists"} and is an [ASCII
    case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#processing-link-headers:ascii-case-insensitive-3
    x-internal="ascii-case-insensitive"} match for a [fetch priority
    attribute](urls-and-fetching.html#fetch-priority-attribute){#processing-link-headers:fetch-priority-attribute-2}
    keyword, then set `options`{.variable}\'s [fetch
    priority](#link-options-fetch-priority){#processing-link-headers:link-options-fetch-priority-2}
    to that [fetch priority
    attribute](urls-and-fetching.html#fetch-priority-attribute){#processing-link-headers:fetch-priority-attribute-3}
    keyword.

##### [4.2.4.5]{.secno} Early hints[](#early-hints){.self-link}

::::: {.mdn-anno .wrapped}
MDN

:::: feature
[Status/103](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/103 "The HTTP 103 Early Hints information response may be sent by a server while it is still preparing a response, with hints about the resources that the server is expecting the final response will link. This allows a browser to start preloading resources even before the server has prepared and sent that final response.")

::: support
[Firefoxpreview+]{.firefox .yes}[SafariNo]{.safari
.no}[Chrome103+]{.chrome .yes}

------------------------------------------------------------------------

[OperaNo]{.opera .no}[Edge103+]{.edge_blink .yes}

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

[Early hints]{#early-hints-2 .dfn} allow user-agents to perform some
operations, such as to speculatively load resources that are likely to
be used by the document, before the navigation request is fully handled
by the server and a response code is served. Servers can indicate early
hints by serving a
[response](https://fetch.spec.whatwg.org/#concept-response){#early-hints:concept-response
x-internal="concept-response"} with a 103 status code before serving the
final
[response](https://fetch.spec.whatwg.org/#concept-response){#early-hints:concept-response-2
x-internal="concept-response"}.[\[RFC8297\]](references.html#refsRFC8297)

For compatibility reasons [early hints are typically delivered over
HTTP/2 or
above](https://httpwg.org/specs/rfc8297.html#security-considerations),
but for readability we use HTTP/1.1-style notation below.

::: example
For example, given the following sequence of responses:

    103 Early Hint
    Link: </image.png>; rel=preload; as=image

    200 OK
    Content-Type: text/html

    <!DOCTYPE html>
    ...
    <img src="/image.png">

the image will start loading before the HTML content arrives.
:::

Only the first early hint response served during the navigation is
handled, and it is discarded if it is succeeded by a cross-origin
redirect.

In addition to the
\`[`Link`{#early-hints:http-link}](https://httpwg.org/specs/rfc8288.html#header){x-internal="http-link"}\`
headers, it is possible that the 103 response contains a [Content
Security
Policy](https://w3c.github.io/webappsec-csp/#content-security-policy-object){#early-hints:content-security-policy
x-internal="content-security-policy"} header, which is enforced when
processing the early hint.

::: example
For example, given the following sequence of responses:

    103 Early Hint
    Content-Security-Policy: style-src: self;
    Link: </style.css>; rel=preload; as=style

    103 Early Hint
    Link: </image.png>; rel=preload; as=image

    302 Redirect
    Location: /alternate.html

    200 OK
    Content-Security-Policy: style-src: none;
    Link: </font.ttf>; rel=preload; as=font

The font and style would be loaded, and the image will be discarded, as
only the first early hint response in the final redirect chain is
respected. The late [Content Security
Policy](https://w3c.github.io/webappsec-csp/#content-security-policy-object){#early-hints:content-security-policy-2
x-internal="content-security-policy"} header comes after the request to
fetch the style has already been performed, but the style will not be
accessible to the document.
:::

To [process early hint headers]{#process-early-hint-headers .dfn} given
a
[response](https://fetch.spec.whatwg.org/#concept-response){#early-hints:concept-response-3
x-internal="concept-response"} `response`{.variable} and an
[environment](webappapis.html#environment){#early-hints:environment}
`reservedEnvironment`{.variable}:

Early-hint
\`[`Link`{#early-hints:http-link-2}](https://httpwg.org/specs/rfc8288.html#header){x-internal="http-link"}\`
headers are always processed before
\`[`Link`{#early-hints:http-link-3}](https://httpwg.org/specs/rfc8288.html#header){x-internal="http-link"}\`
headers from the final
[response](https://fetch.spec.whatwg.org/#concept-response){#early-hints:concept-response-4
x-internal="concept-response"}, followed by
[`link`{#early-hints:the-link-element}](#the-link-element) elements.
This is equivalent to prepending the contents of the early and final
\`[`Link`{#early-hints:http-link-4}](https://httpwg.org/specs/rfc8288.html#header){x-internal="http-link"}\`
headers to the [`Document`{#early-hints:document}](dom.html#document)\'s
[`head`{#early-hints:the-head-element}](#the-head-element) element, in
respective order.

1.  Let `earlyPolicyContainer`{.variable} be the result of [creating a
    policy container from a fetch
    response](browsers.html#creating-a-policy-container-from-a-fetch-response){#early-hints:creating-a-policy-container-from-a-fetch-response}
    given `response`{.variable} and `reservedEnvironment`{.variable}.

    This allows the early hint
    [response](https://fetch.spec.whatwg.org/#concept-response){#early-hints:concept-response-5
    x-internal="concept-response"} to include a [Content Security
    Policy](https://w3c.github.io/webappsec-csp/#content-security-policy-object){#early-hints:content-security-policy-3
    x-internal="content-security-policy"} which would be
    [enforced](https://w3c.github.io/webappsec-csp/#enforced){#early-hints:enforce-the-policy
    x-internal="enforce-the-policy"} when fetching the early hint
    [request](https://fetch.spec.whatwg.org/#concept-request){#early-hints:concept-request
    x-internal="concept-request"}.

2.  Let `links`{.variable} be the result of [extracting
    links](#extract-links-from-headers){#early-hints:extract-links-from-headers}
    from `response`{.variable}\'s [header
    list](https://fetch.spec.whatwg.org/#concept-response-header-list){#early-hints:concept-response-header-list
    x-internal="concept-response-header-list"}.

3.  Let `earlyHints`{.variable} be an empty
    [list](https://infra.spec.whatwg.org/#list){#early-hints:list
    x-internal="list"}.

4.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#early-hints:list-iterate
    x-internal="list-iterate"} `linkObject`{.variable} in
    `links`{.variable}:

    The moment we receive the early hint link header, we begin
    [fetching](https://fetch.spec.whatwg.org/#concept-fetch){#early-hints:concept-fetch
    x-internal="concept-fetch"} `earlyRequest`{.variable}. If it comes
    back before the
    [`Document`{#early-hints:document-2}](dom.html#document) is created,
    we set `earlyResponse`{.variable} to the
    [response](https://fetch.spec.whatwg.org/#concept-response){#early-hints:concept-response-6
    x-internal="concept-response"} of that
    [fetch](https://fetch.spec.whatwg.org/#concept-fetch){#early-hints:concept-fetch-2
    x-internal="concept-fetch"} and once the
    [`Document`{#early-hints:document-3}](dom.html#document) is created
    we commit it (by making it available in the [map of preloaded
    resources](links.html#map-of-preloaded-resources){#early-hints:map-of-preloaded-resources}
    as if it was a
    [`link`{#early-hints:the-link-element-2}](#the-link-element)
    element). If the
    [`Document`{#early-hints:document-4}](dom.html#document) is created
    first, the
    [response](https://fetch.spec.whatwg.org/#concept-response){#early-hints:concept-response-7
    x-internal="concept-response"} is committed as soon as it becomes
    available.

    1.  Let `rel`{.variable} be
        `linkObject`{.variable}\[\"`relation_type`\"\].

    2.  Let `options`{.variable} be a new [link processing
        options](#link-processing-options){#early-hints:link-processing-options}
        with

        [href](#link-options-href){#early-hints:link-options-href}
        :   `linkObject`{.variable}\[\"`target_uri`\"\]

        [initiator](#link-options-initiator){#early-hints:link-options-initiator}
        :   \"`early-hint`\"

        [base URL](#link-options-base-url){#early-hints:link-options-base-url}
        :   `response`{.variable}\'s
            [URL](https://fetch.spec.whatwg.org/#concept-response-url){#early-hints:concept-response-url
            x-internal="concept-response-url"}

        [origin](#link-options-origin){#early-hints:link-options-origin}
        :   `response`{.variable}\'s
            [URL](https://fetch.spec.whatwg.org/#concept-response-url){#early-hints:concept-response-url-2
            x-internal="concept-response-url"}\'s
            [origin](https://url.spec.whatwg.org/#concept-url-origin){#early-hints:concept-url-origin
            x-internal="concept-url-origin"}

        [environment](#link-options-environment){#early-hints:link-options-environment}
        :   `reservedEnvironment`{.variable}

        [policy container](#link-options-policy-container){#early-hints:link-options-policy-container}
        :   `earlyPolicyContainer`{.variable}

    3.  Let `attribs`{.variable} be
        `linkObject`{.variable}\[\"`target_attributes`\"\].

        Only the [`as`{#early-hints:attr-link-as-5}](#attr-link-as),
        [`crossorigin`{#early-hints:attr-link-crossorigin}](#attr-link-crossorigin),
        [`integrity`{#early-hints:attr-link-integrity}](#attr-link-integrity),
        and [`type`{#early-hints:attr-link-type}](#attr-link-type)
        attributes are handled as part of early hint processing. The
        other ones, in particular
        [`blocking`{#early-hints:attr-link-blocking}](#attr-link-blocking),
        [`imagesrcset`{#early-hints:attr-link-imagesrcset}](#attr-link-imagesrcset),
        [`imagesizes`{#early-hints:attr-link-imagesizes}](#attr-link-imagesizes),
        and [`media`{#early-hints:attr-link-media}](#attr-link-media)
        are only applicable once a
        [`Document`{#early-hints:document-5}](dom.html#document) is
        created.

    4.  [Apply link options from parsed header
        attributes](#apply-link-options-from-parsed-header-attributes){#early-hints:apply-link-options-from-parsed-header-attributes}
        to `options`{.variable} given `attribs`{.variable}.

    5.  Run the [process a link
        header](#process-a-link-header){#early-hints:process-a-link-header}
        steps for `rel`{.variable} given `options`{.variable}.

    6.  [Append](https://infra.spec.whatwg.org/#list-append){#early-hints:list-append
        x-internal="list-append"} `options`{.variable} to
        `earlyHints`{.variable}.

5.  Return the following substeps given
    [`Document`{#early-hints:document-6}](dom.html#document)
    `doc`{.variable}: [for
    each](https://infra.spec.whatwg.org/#list-iterate){#early-hints:list-iterate-2
    x-internal="list-iterate"} `options`{.variable} in
    `earlyHints`{.variable}:

    1.  If `options`{.variable}\'s [on document
        ready](#link-options-on-document-ready){#early-hints:link-options-on-document-ready}
        is null, then set `options`{.variable}\'s
        [document](#link-options-document){#early-hints:link-options-document}
        to `doc`{.variable}.

    2.  Otherwise, call `options`{.variable}\'s [on document
        ready](#link-options-on-document-ready){#early-hints:link-options-on-document-ready-2}
        with `doc`{.variable}.

##### [4.2.4.6]{.secno} Providing users with a means to follow hyperlinks created using the [`link`{#providing-users-with-a-means-to-follow-hyperlinks-created-using-the-link-element:the-link-element}](#the-link-element) element[](#providing-users-with-a-means-to-follow-hyperlinks-created-using-the-link-element){.self-link}

Interactive user agents may provide users with a means to [follow the
hyperlinks](links.html#following-hyperlinks-2){#providing-users-with-a-means-to-follow-hyperlinks-created-using-the-link-element:following-hyperlinks-2}
created using the
[`link`{#providing-users-with-a-means-to-follow-hyperlinks-created-using-the-link-element:the-link-element-2}](#the-link-element)
element, somewhere within their user interface. Such invocations of the
[follow the
hyperlink](links.html#following-hyperlinks-2){#providing-users-with-a-means-to-follow-hyperlinks-created-using-the-link-element:following-hyperlinks-2-2}
algorithm must set the
*[userInvolvement](links.html#following-userinvolvement)* argument to
\"[`browser UI`{#providing-users-with-a-means-to-follow-hyperlinks-created-using-the-link-element:uni-browser-ui}](browsing-the-web.html#uni-browser-ui)\".
The exact interface is not defined by this specification, but it could
include the following information (obtained from the element\'s
attributes, again as defined below), in some form or another (possibly
simplified), for each
[hyperlink](links.html#hyperlink){#providing-users-with-a-means-to-follow-hyperlinks-created-using-the-link-element:hyperlink}
created with each
[`link`{#providing-users-with-a-means-to-follow-hyperlinks-created-using-the-link-element:the-link-element-3}](#the-link-element)
element in the document:

- The relationship between this document and the resource (given by the
  [`rel`{#providing-users-with-a-means-to-follow-hyperlinks-created-using-the-link-element:attr-link-rel}](#attr-link-rel)
  attribute)
- The title of the resource (given by the
  [`title`{#providing-users-with-a-means-to-follow-hyperlinks-created-using-the-link-element:attr-link-title}](#attr-link-title)
  attribute).
- The address of the resource (given by the
  [`href`{#providing-users-with-a-means-to-follow-hyperlinks-created-using-the-link-element:attr-link-href}](#attr-link-href)
  attribute).
- The language of the resource (given by the
  [`hreflang`{#providing-users-with-a-means-to-follow-hyperlinks-created-using-the-link-element:attr-link-hreflang}](#attr-link-hreflang)
  attribute).
- The optimum media for the resource (given by the
  [`media`{#providing-users-with-a-means-to-follow-hyperlinks-created-using-the-link-element:attr-link-media}](#attr-link-media)
  attribute).

User agents could also include other information, such as the type of
the resource (as given by the
[`type`{#providing-users-with-a-means-to-follow-hyperlinks-created-using-the-link-element:attr-link-type}](#attr-link-type)
attribute).

#### [4.2.5]{.secno} The [`meta`]{#meta .dfn dfn-type="element"} element[](#the-meta-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/meta](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/meta "The <meta> HTML element represents metadata that cannot be represented by other HTML meta-related elements, like <base>, <link>, <script>, <style> or <title>.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
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
[HTMLMetaElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMetaElement "The HTMLMetaElement interface contains descriptive metadata about a document provided in HTML as <meta> elements. This interface inherits all of the properties and methods described in the HTMLElement interface.")

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
:::::

[Categories](dom.html#concept-element-categories){#the-meta-element:concept-element-categories}:
:   [Metadata
    content](dom.html#metadata-content-2){#the-meta-element:metadata-content-2}.
:   If the
    [`itemprop`{#the-meta-element:names:-the-itemprop-attribute}](microdata.html#names:-the-itemprop-attribute)
    attribute is present: [flow
    content](dom.html#flow-content-2){#the-meta-element:flow-content-2}.
:   If the
    [`itemprop`{#the-meta-element:names:-the-itemprop-attribute-2}](microdata.html#names:-the-itemprop-attribute)
    attribute is present: [phrasing
    content](dom.html#phrasing-content-2){#the-meta-element:phrasing-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-meta-element:concept-element-contexts}:
:   If the
    [`charset`{#the-meta-element:attr-meta-charset}](#attr-meta-charset)
    attribute is present, or if the element\'s
    [`http-equiv`{#the-meta-element:attr-meta-http-equiv}](#attr-meta-http-equiv)
    attribute is in the [Encoding declaration
    state](#attr-meta-http-equiv-content-type){#the-meta-element:attr-meta-http-equiv-content-type}:
    in a [`head`{#the-meta-element:the-head-element}](#the-head-element)
    element.
:   If the
    [`http-equiv`{#the-meta-element:attr-meta-http-equiv-2}](#attr-meta-http-equiv)
    attribute is present but not in the [Encoding declaration
    state](#attr-meta-http-equiv-content-type){#the-meta-element:attr-meta-http-equiv-content-type-2}:
    in a
    [`head`{#the-meta-element:the-head-element-2}](#the-head-element)
    element.
:   If the
    [`http-equiv`{#the-meta-element:attr-meta-http-equiv-3}](#attr-meta-http-equiv)
    attribute is present but not in the [Encoding declaration
    state](#attr-meta-http-equiv-content-type){#the-meta-element:attr-meta-http-equiv-content-type-3}:
    in a
    [`noscript`{#the-meta-element:the-noscript-element}](scripting.html#the-noscript-element)
    element that is a child of a
    [`head`{#the-meta-element:the-head-element-3}](#the-head-element)
    element.
:   If the [`name`{#the-meta-element:attr-meta-name}](#attr-meta-name)
    attribute is present: where [metadata
    content](dom.html#metadata-content-2){#the-meta-element:metadata-content-2-2}
    is expected.
:   If the
    [`itemprop`{#the-meta-element:names:-the-itemprop-attribute-3}](microdata.html#names:-the-itemprop-attribute)
    attribute is present: where [metadata
    content](dom.html#metadata-content-2){#the-meta-element:metadata-content-2-3}
    is expected.
:   If the
    [`itemprop`{#the-meta-element:names:-the-itemprop-attribute-4}](microdata.html#names:-the-itemprop-attribute)
    attribute is present: where [phrasing
    content](dom.html#phrasing-content-2){#the-meta-element:phrasing-content-2-2}
    is expected.

[Content model](dom.html#concept-element-content-model){#the-meta-element:concept-element-content-model}:
:   [Nothing](dom.html#concept-content-nothing){#the-meta-element:concept-content-nothing}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-meta-element:concept-element-tag-omission}:
:   No [end
    tag](syntax.html#syntax-end-tag){#the-meta-element:syntax-end-tag}.

[Content attributes](dom.html#concept-element-attributes){#the-meta-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-meta-element:global-attributes}
:   [`name`{#the-meta-element:attr-meta-name-2}](#attr-meta-name) ---
    Metadata name
:   [`http-equiv`{#the-meta-element:attr-meta-http-equiv-4}](#attr-meta-http-equiv)
    --- Pragma directive
:   [`content`{#the-meta-element:attr-meta-content}](#attr-meta-content)
    --- Value of the element
:   [`charset`{#the-meta-element:attr-meta-charset-2}](#attr-meta-charset)
    --- [Character encoding
    declaration](#character-encoding-declaration){#the-meta-element:character-encoding-declaration}
:   [`media`{#the-meta-element:attr-meta-media}](#attr-meta-media) ---
    Applicable media

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-meta-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-meta).
:   [For implementers](https://w3c.github.io/html-aam/#el-meta).

[DOM interface](dom.html#concept-element-dom){#the-meta-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLMetaElement : HTMLElement {
      [HTMLConstructor] constructor();

      [CEReactions] attribute DOMString name;
      [CEReactions] attribute DOMString httpEquiv;
      [CEReactions] attribute DOMString content;
      [CEReactions] attribute DOMString media;

      // also has obsolete members
    };
    ```

The [`meta`{#the-meta-element:the-meta-element}](#the-meta-element)
element [represents](dom.html#represents){#the-meta-element:represents}
various kinds of metadata that cannot be expressed using the
[`title`{#the-meta-element:the-title-element}](#the-title-element),
[`base`{#the-meta-element:the-base-element}](#the-base-element),
[`link`{#the-meta-element:the-link-element}](#the-link-element),
[`style`{#the-meta-element:the-style-element}](#the-style-element), and
[`script`{#the-meta-element:the-script-element}](scripting.html#the-script-element)
elements.

The [`meta`{#the-meta-element:the-meta-element-2}](#the-meta-element)
element can represent document-level metadata with the
[`name`{#the-meta-element:attr-meta-name-3}](#attr-meta-name) attribute,
pragma directives with the
[`http-equiv`{#the-meta-element:attr-meta-http-equiv-5}](#attr-meta-http-equiv)
attribute, and the file\'s [character encoding
declaration](#character-encoding-declaration){#the-meta-element:character-encoding-declaration-2}
when an HTML document is serialized to string form (e.g. for
transmission over the network or for disk storage) with the
[`charset`{#the-meta-element:attr-meta-charset-3}](#attr-meta-charset)
attribute.

Exactly one of the
[`name`{#the-meta-element:attr-meta-name-4}](#attr-meta-name),
[`http-equiv`{#the-meta-element:attr-meta-http-equiv-6}](#attr-meta-http-equiv),
[`charset`{#the-meta-element:attr-meta-charset-4}](#attr-meta-charset),
and
[`itemprop`{#the-meta-element:names:-the-itemprop-attribute-5}](microdata.html#names:-the-itemprop-attribute)
attributes must be specified.

If either [`name`{#the-meta-element:attr-meta-name-5}](#attr-meta-name),
[`http-equiv`{#the-meta-element:attr-meta-http-equiv-7}](#attr-meta-http-equiv),
or
[`itemprop`{#the-meta-element:names:-the-itemprop-attribute-6}](microdata.html#names:-the-itemprop-attribute)
is specified, then the
[`content`{#the-meta-element:attr-meta-content-2}](#attr-meta-content)
attribute must also be specified. Otherwise, it must be omitted.

The [`charset`]{#attr-meta-charset .dfn dfn-for="meta"
dfn-type="element-attr"} attribute specifies the [character
encoding](https://encoding.spec.whatwg.org/#encoding){#the-meta-element:encoding
x-internal="encoding"} used by the document. This is a [character
encoding
declaration](#character-encoding-declaration){#the-meta-element:character-encoding-declaration-3}.
If the attribute is present, its value must be an [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#the-meta-element:ascii-case-insensitive
x-internal="ascii-case-insensitive"} match for the string \"`utf-8`\".

The
[`charset`{#the-meta-element:attr-meta-charset-5}](#attr-meta-charset)
attribute on the
[`meta`{#the-meta-element:the-meta-element-3}](#the-meta-element)
element has no effect in XML documents, but is allowed in XML documents
in order to facilitate migration to and from XML.

There must not be more than one
[`meta`{#the-meta-element:the-meta-element-4}](#the-meta-element)
element with a
[`charset`{#the-meta-element:attr-meta-charset-6}](#attr-meta-charset)
attribute per document.

The [`content`]{#attr-meta-content .dfn dfn-for="meta"
dfn-type="element-attr"} attribute gives the value of the document
metadata or pragma directive when the element is used for those
purposes. The allowed values depend on the exact context, as described
in subsequent sections of this specification.

If a [`meta`{#the-meta-element:the-meta-element-5}](#the-meta-element)
element has a [`name`]{#attr-meta-name .dfn dfn-for="meta"
dfn-type="element-attr"} attribute, it sets document metadata. Document
metadata is expressed in terms of name-value pairs, the
[`name`{#the-meta-element:attr-meta-name-6}](#attr-meta-name) attribute
on the [`meta`{#the-meta-element:the-meta-element-6}](#the-meta-element)
element giving the name, and the
[`content`{#the-meta-element:attr-meta-content-3}](#attr-meta-content)
attribute on the same element giving the value. The name specifies what
aspect of metadata is being set; valid names and the meaning of their
values are described in the following sections. If a
[`meta`{#the-meta-element:the-meta-element-7}](#the-meta-element)
element has no
[`content`{#the-meta-element:attr-meta-content-4}](#attr-meta-content)
attribute, then the value part of the metadata name-value pair is the
empty string.

The [`media`]{#attr-meta-media .dfn dfn-for="meta"
dfn-type="element-attr"} attribute says which media the metadata applies
to. The value must be a [valid media query
list](common-microsyntaxes.html#valid-media-query-list){#the-meta-element:valid-media-query-list}.
Unless the [`name`{#the-meta-element:attr-meta-name-7}](#attr-meta-name)
is
[`theme-color`{#the-meta-element:meta-theme-color}](#meta-theme-color),
the [`media`{#the-meta-element:attr-meta-media-2}](#attr-meta-media)
attribute has no effect on the processing model and must not be used by
authors.

The [`name`]{#dom-meta-name .dfn dfn-for="HTMLMetaElement"
dfn-type="attribute"}, [`content`]{#dom-meta-content .dfn
dfn-for="HTMLMetaElement" dfn-type="attribute"}, and
[`media`]{#dom-meta-media .dfn dfn-for="HTMLMetaElement"
dfn-type="attribute"} IDL attributes must
[reflect](common-dom-interfaces.html#reflect){#the-meta-element:reflect}
the respective content attributes of the same name. The IDL attribute
[`httpEquiv`]{#dom-meta-httpequiv .dfn dfn-for="HTMLMetaElement"
dfn-type="attribute"} must
[reflect](common-dom-interfaces.html#reflect){#the-meta-element:reflect-2}
the content attribute
[`http-equiv`{#the-meta-element:attr-meta-http-equiv-8}](#attr-meta-http-equiv).

##### [4.2.5.1]{.secno} Standard metadata names[](#standard-metadata-names){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/meta/name](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/meta/name "The <meta> element can be used to provide document metadata in terms of name-value pairs, with the name attribute giving the metadata name, and the content attribute giving the value.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer6+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

This specification defines a few names for the
[`name`{#standard-metadata-names:attr-meta-name}](#attr-meta-name)
attribute of the
[`meta`{#standard-metadata-names:the-meta-element}](#the-meta-element)
element.

Names are case-insensitive[, and must be compared in an [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#standard-metadata-names:ascii-case-insensitive
x-internal="ascii-case-insensitive"} manner]{.impl}.

[`application-name`]{#meta-application-name .dfn dfn-for="meta/name"
dfn-type="attr-value"}

The value must be a short free-form string giving the name of the web
application that the page represents. If the page is not a web
application, the
[`application-name`{#standard-metadata-names:meta-application-name}](#meta-application-name)
metadata name must not be used. Translations of the web application\'s
name may be given, using the
[`lang`{#standard-metadata-names:attr-lang}](dom.html#attr-lang)
attribute to specify the language of each name.

There must not be more than one
[`meta`{#standard-metadata-names:the-meta-element-2}](#the-meta-element)
element with a given
[language](dom.html#language){#standard-metadata-names:language} and
where the
[`name`{#standard-metadata-names:attr-meta-name-2}](#attr-meta-name)
attribute value is an [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#standard-metadata-names:ascii-case-insensitive-2
x-internal="ascii-case-insensitive"} match for
[`application-name`{#standard-metadata-names:meta-application-name-2}](#meta-application-name)
per document.

User agents may use the application name in UI in preference to the
page\'s
[`title`{#standard-metadata-names:the-title-element}](#the-title-element),
since the title might include status messages and the like relevant to
the status of the page at a particular moment in time instead of just
being the name of the application.

To find the application name to use given an ordered list of languages
(e.g. British English, American English, and English), user agents must
run the following steps:

1.  Let `languages`{.variable} be the list of languages.

2.  Let `default language`{.variable} be the
    [language](dom.html#language){#standard-metadata-names:language-2}
    of the
    [`Document`{#standard-metadata-names:document}](dom.html#document)\'s
    [document
    element](https://dom.spec.whatwg.org/#document-element){#standard-metadata-names:document-element
    x-internal="document-element"}, if any, and if that language is not
    unknown.

3.  If there is a `default language`{.variable}, and if it is not the
    same language as any of the languages in `languages`{.variable},
    append it to `languages`{.variable}.

4.  Let `winning language`{.variable} be the first language in
    `languages`{.variable} for which there is a
    [`meta`{#standard-metadata-names:the-meta-element-3}](#the-meta-element)
    element in the
    [`Document`{#standard-metadata-names:document-2}](dom.html#document)
    where the
    [`name`{#standard-metadata-names:attr-meta-name-3}](#attr-meta-name)
    attribute value is an [ASCII
    case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#standard-metadata-names:ascii-case-insensitive-3
    x-internal="ascii-case-insensitive"} match for
    [`application-name`{#standard-metadata-names:meta-application-name-3}](#meta-application-name)
    and whose
    [language](dom.html#language){#standard-metadata-names:language-3}
    is the language in question.

    If none of the languages have such a
    [`meta`{#standard-metadata-names:the-meta-element-4}](#the-meta-element)
    element, then return; there\'s no given application name.

5.  Return the value of the
    [`content`{#standard-metadata-names:attr-meta-content}](#attr-meta-content)
    attribute of the first
    [`meta`{#standard-metadata-names:the-meta-element-5}](#the-meta-element)
    element in the
    [`Document`{#standard-metadata-names:document-3}](dom.html#document)
    in [tree
    order](https://dom.spec.whatwg.org/#concept-tree-order){#standard-metadata-names:tree-order
    x-internal="tree-order"} where the
    [`name`{#standard-metadata-names:attr-meta-name-4}](#attr-meta-name)
    attribute value is an [ASCII
    case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#standard-metadata-names:ascii-case-insensitive-4
    x-internal="ascii-case-insensitive"} match for
    [`application-name`{#standard-metadata-names:meta-application-name-4}](#meta-application-name)
    and whose
    [language](dom.html#language){#standard-metadata-names:language-4}
    is `winning language`{.variable}.

This algorithm would be used by a browser when it needs a name for the
page, for instance, to label a bookmark. The languages it would provide
to the algorithm would be the user\'s preferred languages.

[`author`]{#meta-author .dfn dfn-for="meta/name" dfn-type="attr-value"}

The value must be a free-form string giving the name of one of the
page\'s authors.

[`description`]{#meta-description .dfn dfn-for="meta/name"
dfn-type="attr-value"}

The value must be a free-form string that describes the page. The value
must be appropriate for use in a directory of pages, e.g. in a search
engine. There must not be more than one
[`meta`{#standard-metadata-names:the-meta-element-6}](#the-meta-element)
element where the
[`name`{#standard-metadata-names:attr-meta-name-5}](#attr-meta-name)
attribute value is an [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#standard-metadata-names:ascii-case-insensitive-5
x-internal="ascii-case-insensitive"} match for
[`description`{#standard-metadata-names:meta-description}](#meta-description)
per document.

[`generator`]{#meta-generator .dfn dfn-for="meta/name"
dfn-type="attr-value"}

The value must be a free-form string that identifies one of the software
packages used to generate the document. This value must not be used on
pages whose markup is not generated by software, e.g. pages whose markup
was written by a user in a text editor.

::: example
Here is what a tool called \"Frontweaver\" could include in its output,
in the page\'s
[`head`{#standard-metadata-names:the-head-element}](#the-head-element)
element, to identify itself as the tool used to generate the page:

``` html
<meta name=generator content="Frontweaver 8.2">
```
:::

[`keywords`]{#meta-keywords .dfn dfn-for="meta/name"
dfn-type="attr-value"}

The value must be a [set of comma-separated
tokens](common-microsyntaxes.html#set-of-comma-separated-tokens){#standard-metadata-names:set-of-comma-separated-tokens},
each of which is a keyword relevant to the page.

::: example
This page about typefaces on British motorways uses a
[`meta`{#standard-metadata-names:the-meta-element-7}](#the-meta-element)
element to specify some keywords that users might use to look for the
page:

``` html
<!DOCTYPE HTML>
<html lang="en-GB">
 <head>
  <title>Typefaces on UK motorways</title>
  <meta name="keywords" content="british,type face,font,fonts,highway,highways">
 </head>
 <body>
  ...
```
:::

Many search engines do not consider such keywords, because this feature
has historically been used unreliably and even misleadingly as a way to
spam search engine results in a way that is not helpful for users.

To obtain the list of keywords that the author has specified as
applicable to the page, the user agent must run the following steps:

1.  Let `keywords`{.variable} be an empty list.

2.  For each
    [`meta`{#standard-metadata-names:the-meta-element-8}](#the-meta-element)
    element with a
    [`name`{#standard-metadata-names:attr-meta-name-6}](#attr-meta-name)
    attribute and a
    [`content`{#standard-metadata-names:attr-meta-content-2}](#attr-meta-content)
    attribute and where the
    [`name`{#standard-metadata-names:attr-meta-name-7}](#attr-meta-name)
    attribute value is an [ASCII
    case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#standard-metadata-names:ascii-case-insensitive-6
    x-internal="ascii-case-insensitive"} match for
    [`keywords`{#standard-metadata-names:meta-keywords}](#meta-keywords):

    1.  [Split the value of the element\'s `content` attribute on
        commas](https://infra.spec.whatwg.org/#split-on-commas){#standard-metadata-names:split-a-string-on-commas
        x-internal="split-a-string-on-commas"}.

    2.  Add the resulting tokens, if any, to `keywords`{.variable}.

3.  Remove any duplicates from `keywords`{.variable}.

4.  Return `keywords`{.variable}. This is the list of keywords that the
    author has specified as applicable to the page.

User agents should not use this information when there is insufficient
confidence in the reliability of the value.

For instance, it would be reasonable for a content management system to
use the keyword information of pages within the system to populate the
index of a site-specific search engine, but a large-scale content
aggregator that used this information would likely find that certain
users would try to game its ranking mechanism through the use of
inappropriate keywords.

[`referrer`]{#meta-referrer .dfn dfn-for="meta/name"
dfn-type="attr-value"}

The value must be a [referrer
policy](https://w3c.github.io/webappsec-referrer-policy/#referrer-policy){#standard-metadata-names:referrer-policy
x-internal="referrer-policy"}, which defines the default [referrer
policy](https://w3c.github.io/webappsec-referrer-policy/#referrer-policy){#standard-metadata-names:referrer-policy-2
x-internal="referrer-policy"} for the
[`Document`{#standard-metadata-names:document-4}](dom.html#document).
[\[REFERRERPOLICY\]](references.html#refsREFERRERPOLICY)

If any
[`meta`{#standard-metadata-names:the-meta-element-9}](#the-meta-element)
element `element`{.variable} is [inserted into the
document](infrastructure.html#insert-an-element-into-a-document){#standard-metadata-names:insert-an-element-into-a-document},
or has its
[`name`{#standard-metadata-names:attr-meta-name-8}](#attr-meta-name) or
[`content`{#standard-metadata-names:attr-meta-content-3}](#attr-meta-content)
attributes changed, user agents must run the following algorithm:

1.  If `element`{.variable} is not [in a document
    tree](https://dom.spec.whatwg.org/#in-a-document-tree){#standard-metadata-names:in-a-document-tree
    x-internal="in-a-document-tree"}, then return.

2.  If `element`{.variable} does not have a
    [`name`{#standard-metadata-names:attr-meta-name-9}](#attr-meta-name)
    attribute whose value is an [ASCII
    case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#standard-metadata-names:ascii-case-insensitive-7
    x-internal="ascii-case-insensitive"} match for
    \"[`referrer`{#standard-metadata-names:meta-referrer}](#meta-referrer)\",
    then return.

3.  If `element`{.variable} does not have a
    [`content`{#standard-metadata-names:attr-meta-content-4}](#attr-meta-content)
    attribute, or that attribute\'s value is the empty string, then
    return.

4.  Let `value`{.variable} be the value of `element`{.variable}\'s
    [`content`{#standard-metadata-names:attr-meta-content-5}](#attr-meta-content)
    attribute, [converted to ASCII
    lowercase](https://infra.spec.whatwg.org/#ascii-lowercase){#standard-metadata-names:converted-to-ascii-lowercase
    x-internal="converted-to-ascii-lowercase"}.

5.  If `value`{.variable} is one of the values given in the first column
    of the following table, then set `value`{.variable} to the value
    given in the second column:

    Legacy value

    Referrer policy

    `never`

    [`no-referrer`{#standard-metadata-names:referrer-policy-no-referrer}](https://w3c.github.io/webappsec-referrer-policy/#referrer-policy-no-referrer){x-internal="referrer-policy-no-referrer"}

    `default`

    the [default referrer
    policy](https://w3c.github.io/webappsec-referrer-policy/#default-referrer-policy){#standard-metadata-names:default-referrer-policy
    x-internal="default-referrer-policy"}

    `always`

    [`unsafe-url`{#standard-metadata-names:referrer-policy-unsafe-url}](https://w3c.github.io/webappsec-referrer-policy/#referrer-policy-unsafe-url){x-internal="referrer-policy-unsafe-url"}

    `origin-when-crossorigin`

    [`origin-when-cross-origin`{#standard-metadata-names:referrer-policy-origin-when-cross-origin}](https://w3c.github.io/webappsec-referrer-policy/#referrer-policy-origin-when-cross-origin){x-internal="referrer-policy-origin-when-cross-origin"}

6.  If `value`{.variable} is a [referrer
    policy](https://w3c.github.io/webappsec-referrer-policy/#referrer-policy){#standard-metadata-names:referrer-policy-3
    x-internal="referrer-policy"}, then set `element`{.variable}\'s
    [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#standard-metadata-names:node-document
    x-internal="node-document"}\'s [policy
    container](dom.html#concept-document-policy-container){#standard-metadata-names:concept-document-policy-container}\'s
    [referrer
    policy](browsers.html#policy-container-referrer-policy){#standard-metadata-names:policy-container-referrer-policy}
    to `policy`{.variable}.

For historical reasons, unlike other standard metadata names, the
processing model for
[`referrer`{#standard-metadata-names:meta-referrer-2}](#meta-referrer)
is not responsive to element removals, and does not use [tree
order](https://dom.spec.whatwg.org/#concept-tree-order){#standard-metadata-names:tree-order-2
x-internal="tree-order"}. Only the most-recently-inserted or
most-recently-modified
[`meta`{#standard-metadata-names:the-meta-element-10}](#the-meta-element)
element in this state has an effect.

[`theme-color`]{#meta-theme-color .dfn dfn-for="meta/name"
dfn-type="attr-value"}

::::: {.mdn-anno .wrapped .before}
MDN

:::: feature
[Element/meta/name/theme-color](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/meta/name/theme-color "The theme-color value for the name attribute of the <meta> element indicates a suggested color that user agents should use to customize the display of the page or of the surrounding user interface. If specified, the content attribute must contain a valid CSS <color>.")

::: support
[FirefoxNo]{.firefox .no}[Safari15+]{.safari .yes}[Chrome[🔰
73+]{title="Partial implementation."}]{.chrome .yes}

------------------------------------------------------------------------

[OperaNo]{.opera .no}[Edge[🔰
79+]{title="Partial implementation."}]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android80+]{.chrome_android .yes}[WebView
AndroidNo]{.webview_android .no}[Samsung
Internet6.2+]{.samsunginternet_android .yes}[Opera
AndroidNo]{.opera_android .no}
:::
::::
:::::

The value must be a string that matches the CSS
[\<color\>](https://drafts.csswg.org/css-color/#typedef-color){#standard-metadata-names:color
x-internal="color"} production, defining a suggested color that user
agents should use to customize the display of the page or of the
surrounding user interface. For example, a browser might color the
page\'s title bar with the specified value, or use it as a color
highlight in a tab bar or task switcher.

Within an HTML document, the
[`media`{#standard-metadata-names:attr-meta-media}](#attr-meta-media)
attribute value must be unique amongst all the
[`meta`{#standard-metadata-names:the-meta-element-11}](#the-meta-element)
elements with their
[`name`{#standard-metadata-names:attr-meta-name-10}](#attr-meta-name)
attribute value set to an [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#standard-metadata-names:ascii-case-insensitive-8
x-internal="ascii-case-insensitive"} match for
[`theme-color`{#standard-metadata-names:meta-theme-color}](#meta-theme-color).

::: example
This standard itself uses \"WHATWG green\" as its theme color:

``` html
<!DOCTYPE HTML>
<title>HTML Standard</title>
<meta name="theme-color" content="#3c790a">
...
```
:::

The
[`media`{#standard-metadata-names:attr-meta-media-2}](#attr-meta-media)
attribute may be used to describe the context in which the provided
color should be used.

::: example
If we only wanted to use \"WHATWG green\" as this standard\'s theme
color in dark mode, we could use the `prefers-color-scheme` media
feature:

``` html
<!DOCTYPE HTML>
<title>HTML Standard</title>
<meta name="theme-color" content="#3c790a" media="(prefers-color-scheme: dark)">
...
```
:::

To obtain a page\'s theme color, user agents must run the following
steps:

1.  Let `candidate elements`{.variable} be the list of all
    [`meta`{#standard-metadata-names:the-meta-element-12}](#the-meta-element)
    elements that meet the following criteria, in [tree
    order](https://dom.spec.whatwg.org/#concept-tree-order){#standard-metadata-names:tree-order-3
    x-internal="tree-order"}:

    - the element is [in a document
      tree](https://dom.spec.whatwg.org/#in-a-document-tree){#standard-metadata-names:in-a-document-tree-2
      x-internal="in-a-document-tree"};

    - the element has a
      [`name`{#standard-metadata-names:attr-meta-name-11}](#attr-meta-name)
      attribute, whose value is an [ASCII
      case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#standard-metadata-names:ascii-case-insensitive-9
      x-internal="ascii-case-insensitive"} match for
      [`theme-color`{#standard-metadata-names:meta-theme-color-2}](#meta-theme-color);
      and

    - the element has a
      [`content`{#standard-metadata-names:attr-meta-content-6}](#attr-meta-content)
      attribute.

2.  For each `element`{.variable} in `candidate elements`{.variable}:

    1.  If `element`{.variable} has a
        [`media`{#standard-metadata-names:attr-link-media}](#attr-link-media)
        attribute and the value of `element`{.variable}\'s
        [`media`{#standard-metadata-names:attr-meta-media-3}](#attr-meta-media)
        attribute does not [match the
        environment](common-microsyntaxes.html#matches-the-environment){#standard-metadata-names:matches-the-environment},
        then
        [continue](https://infra.spec.whatwg.org/#iteration-continue){#standard-metadata-names:continue
        x-internal="continue"}.

    2.  Let `value`{.variable} be the result of [stripping leading and
        trailing ASCII
        whitespace](https://infra.spec.whatwg.org/#strip-leading-and-trailing-ascii-whitespace){#standard-metadata-names:strip-leading-and-trailing-ascii-whitespace
        x-internal="strip-leading-and-trailing-ascii-whitespace"} from
        the value of `element`{.variable}\'s
        [`content`{#standard-metadata-names:attr-meta-content-7}](#attr-meta-content)
        attribute.

    3.  Let `color`{.variable} be the result of
        [parsing](https://drafts.csswg.org/css-color/#parse-a-css-color-value){#standard-metadata-names:parsed-as-a-css-color-value
        x-internal="parsed-as-a-css-color-value"} `value`{.variable}.

    4.  If `color`{.variable} is not failure, then return
        `color`{.variable}.

3.  Return nothing (the page has no theme color).

If any
[`meta`{#standard-metadata-names:the-meta-element-13}](#the-meta-element)
elements are [inserted into the
document](infrastructure.html#insert-an-element-into-a-document){#standard-metadata-names:insert-an-element-into-a-document-2}
or [removed from the
document](infrastructure.html#remove-an-element-from-a-document){#standard-metadata-names:remove-an-element-from-a-document},
or existing
[`meta`{#standard-metadata-names:the-meta-element-14}](#the-meta-element)
elements have their
[`name`{#standard-metadata-names:attr-meta-name-12}](#attr-meta-name),
[`content`{#standard-metadata-names:attr-meta-content-8}](#attr-meta-content),
or
[`media`{#standard-metadata-names:attr-link-media-2}](#attr-link-media)
attributes changed, or if the environment changes such that any
[`meta`{#standard-metadata-names:the-meta-element-15}](#the-meta-element)
element\'s
[`media`{#standard-metadata-names:attr-link-media-3}](#attr-link-media)
attribute\'s value may now or may no longer [match the
environment](common-microsyntaxes.html#matches-the-environment){#standard-metadata-names:matches-the-environment-2},
user agents must re-run the above algorithm and apply the result to any
affected UI.

When using the theme color in UI, user agents may adjust it in
implementation-specific ways to make it more suitable for the UI in
question. For example, if a user agent intends to use the theme color as
a background and display white text over it, it might use a darker
variant of the theme color in that part of the UI, to ensure adequate
contrast.

[`color-scheme`]{#meta-color-scheme .dfn dfn-for="meta/name"
dfn-type="attr-value"}

To aid user agents in rendering the page background with the desired
color scheme immediately (rather than waiting for all CSS in the page to
load), a
[\'color-scheme\'](https://drafts.csswg.org/css-color-adjust/#color-scheme-prop){#standard-metadata-names:'color-scheme'
x-internal="'color-scheme'"} value can be provided in a
[`meta`{#standard-metadata-names:the-meta-element-16}](#the-meta-element)
element.

The value must be a string that matches the syntax for the CSS
[\'color-scheme\'](https://drafts.csswg.org/css-color-adjust/#color-scheme-prop){#standard-metadata-names:'color-scheme'-2
x-internal="'color-scheme'"} property value. It determines the [page\'s
supported
color-schemes](https://drafts.csswg.org/css-color-adjust/#pages-supported-color-schemes){#standard-metadata-names:page's-supported-color-schemes
x-internal="page's-supported-color-schemes"}.

There must not be more than one
[`meta`{#standard-metadata-names:the-meta-element-17}](#the-meta-element)
element with its
[`name`{#standard-metadata-names:attr-meta-name-13}](#attr-meta-name)
attribute value set to an [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#standard-metadata-names:ascii-case-insensitive-10
x-internal="ascii-case-insensitive"} match for
[`color-scheme`{#standard-metadata-names:meta-color-scheme}](#meta-color-scheme)
per document.

::: example
The following declaration indicates that the page is aware of and can
handle a color scheme with dark background colors and light foreground
colors:

``` html
<meta name="color-scheme" content="dark">
```
:::

To obtain a [page\'s supported
color-schemes](https://drafts.csswg.org/css-color-adjust/#pages-supported-color-schemes){#standard-metadata-names:page's-supported-color-schemes-2
x-internal="page's-supported-color-schemes"}, user agents must run the
following steps:

1.  Let `candidate elements`{.variable} be the list of all
    [`meta`{#standard-metadata-names:the-meta-element-18}](#the-meta-element)
    elements that meet the following criteria, in [tree
    order](https://dom.spec.whatwg.org/#concept-tree-order){#standard-metadata-names:tree-order-4
    x-internal="tree-order"}:

    - the element is [in a document
      tree](https://dom.spec.whatwg.org/#in-a-document-tree){#standard-metadata-names:in-a-document-tree-3
      x-internal="in-a-document-tree"};

    - the element has a
      [`name`{#standard-metadata-names:attr-meta-name-14}](#attr-meta-name)
      attribute, whose value is an [ASCII
      case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#standard-metadata-names:ascii-case-insensitive-11
      x-internal="ascii-case-insensitive"} match for
      [`color-scheme`{#standard-metadata-names:meta-color-scheme-2}](#meta-color-scheme);
      and

    - the element has a
      [`content`{#standard-metadata-names:attr-meta-content-9}](#attr-meta-content)
      attribute.

2.  For each `element`{.variable} in `candidate elements`{.variable}:

    1.  Let `parsed`{.variable} be the result of [parsing a list of
        component
        values](https://drafts.csswg.org/css-syntax/#parse-a-list-of-component-values){#standard-metadata-names:parse-a-list-of-component-values
        x-internal="parse-a-list-of-component-values"} given the value
        of `element`{.variable}\'s
        [`content`{#standard-metadata-names:attr-meta-content-10}](#attr-meta-content)
        attribute.
    2.  If `parsed`{.variable} is a valid CSS
        [\'color-scheme\'](https://drafts.csswg.org/css-color-adjust/#color-scheme-prop){#standard-metadata-names:'color-scheme'-3
        x-internal="'color-scheme'"} property value, then return
        `parsed`{.variable}.

3.  Return null.

If any
[`meta`{#standard-metadata-names:the-meta-element-19}](#the-meta-element)
elements are [inserted into the
document](infrastructure.html#insert-an-element-into-a-document){#standard-metadata-names:insert-an-element-into-a-document-3}
or [removed from the
document](infrastructure.html#remove-an-element-from-a-document){#standard-metadata-names:remove-an-element-from-a-document-2},
or existing
[`meta`{#standard-metadata-names:the-meta-element-20}](#the-meta-element)
elements have their
[`name`{#standard-metadata-names:attr-meta-name-15}](#attr-meta-name) or
[`content`{#standard-metadata-names:attr-meta-content-11}](#attr-meta-content)
attributes changed, user agents must re-run the above algorithm.

Because these rules check successive elements until they find a match,
an author can provide multiple such values to handle fallback for legacy
user agents. Opposite to how CSS fallback works for properties, the
multiple meta elements needs to be arranged with the legacy values after
the newer values.

##### [4.2.5.2]{.secno} Other metadata names[](#other-metadata-names){.self-link}

Anyone can create and use their own [extensions to the predefined set of
metadata names]{#concept-meta-extensions .dfn}. There is no requirement
to register such extensions.

However, a new metadata name should not be created in any of the
following cases:

- If either the name is a
  [URL](https://url.spec.whatwg.org/#concept-url){#other-metadata-names:url
  x-internal="url"}, or the value of its accompanying
  [`content`{#other-metadata-names:attr-meta-content}](#attr-meta-content)
  attribute is a
  [URL](https://url.spec.whatwg.org/#concept-url){#other-metadata-names:url-2
  x-internal="url"}; in those cases, registering it as an [extension to
  the predefined set of link
  types](links.html#concept-rel-extensions){#other-metadata-names:concept-rel-extensions}
  is encouraged (rather than creating a new metadata name).

- If the name is for something expected to have processing requirements
  in user agents; in that case it ought to be standardized.

Also, before creating and using a new metadata name, consulting the
[WHATWG Wiki MetaExtensions
page](https://wiki.whatwg.org/wiki/MetaExtensions) is encouraged --- to
avoid choosing a metadata name that\'s already in use, and to avoid
duplicating the purpose of any metadata names that are already in use,
and to avoid new standardized names clashing with your chosen name.
[\[WHATWGWIKI\]](references.html#refsWHATWGWIKI)

Anyone is free to edit the WHATWG Wiki MetaExtensions page at any time
to add a metadata name. New metadata names can be specified with the
following information:

Keyword
:   The actual name being defined. The name should not be confusingly
    similar to any other defined name (e.g. differing only in case).

Brief description
:   A short non-normative description of what the metadata name\'s
    meaning is, including the format the value is required to be in.

Specification
:   A link to a more detailed description of the metadata name\'s
    semantics and requirements. It could be another page on the wiki, or
    a link to an external page.

Synonyms
:   A list of other names that have exactly the same processing
    requirements. Authors should not use the names defined to be
    synonyms (they are only intended to allow user agents to support
    legacy content). Anyone may remove synonyms that are not used in
    practice; only names that need to be processed as synonyms for
    compatibility with legacy content are to be registered in this way.

Status

:   One of the following:

    Proposed
    :   The name has not received wide peer review and approval. Someone
        has proposed it and is, or soon will be, using it.

    Ratified
    :   The name has received wide peer review and approval. It has a
        specification that unambiguously defines how to handle pages
        that use the name, including when they use it in incorrect ways.

    Discontinued
    :   The metadata name has received wide peer review and it has been
        found wanting. Existing pages are using this metadata name, but
        new pages should avoid it. The \"brief description\" and
        \"specification\" entries will give details of what authors
        should use instead, if anything.

    If a metadata name is found to be redundant with existing values, it
    should be removed and listed as a synonym for the existing value.

    If a metadata name is added in the \"proposed\" state for a period
    of a month or more without being used or specified, then it may be
    removed from the WHATWG Wiki MetaExtensions page.

    If a metadata name is added with the \"proposed\" status and found
    to be redundant with existing values, it should be removed and
    listed as a synonym for the existing value. If a metadata name is
    added with the \"proposed\" status and found to be harmful, then it
    should be changed to \"discontinued\" status.

    Anyone can change the status at any time, but should only do so in
    accordance with the definitions above.

##### [4.2.5.3]{.secno} Pragma directives[](#pragma-directives){.self-link}

When the [`http-equiv`]{#attr-meta-http-equiv .dfn dfn-for="meta"
dfn-type="element-attr"} attribute is specified on a
[`meta`{#pragma-directives:the-meta-element}](#the-meta-element)
element, the element is a pragma directive.

The
[`http-equiv`{#pragma-directives:attr-meta-http-equiv}](#attr-meta-http-equiv)
attribute is an [enumerated
attribute](common-microsyntaxes.html#enumerated-attribute){#pragma-directives:enumerated-attribute}
with the following keywords and states:

Keyword

Conforming

State

Brief description

[`content-language`]{#attr-meta-http-equiv-keyword-content-language .dfn
dfn-for="meta/http-equiv" dfn-type="attr-value"}

No

[Content
language](#attr-meta-http-equiv-content-language){#pragma-directives:attr-meta-http-equiv-content-language}

Sets the [pragma-set default
language](#pragma-set-default-language){#pragma-directives:pragma-set-default-language}.

[`content-type`]{#attr-meta-http-equiv-keyword-content-type .dfn
dfn-for="meta/http-equiv" dfn-type="attr-value"}

[Encoding
declaration](#attr-meta-http-equiv-content-type){#pragma-directives:attr-meta-http-equiv-content-type}

An alternative form of setting the
[`charset`{#pragma-directives:attr-meta-charset}](#attr-meta-charset).

[`default-style`]{#attr-meta-http-equiv-keyword-default-style .dfn
dfn-for="meta/http-equiv" dfn-type="attr-value"}

[Default
style](#attr-meta-http-equiv-default-style){#pragma-directives:attr-meta-http-equiv-default-style}

Sets the
[name](https://drafts.csswg.org/cssom/#css-style-sheet-set-name){#pragma-directives:css-style-sheet-set-name
x-internal="css-style-sheet-set-name"} of the default [CSS style sheet
set](https://drafts.csswg.org/cssom/#css-style-sheet-set){#pragma-directives:css-style-sheet-set
x-internal="css-style-sheet-set"}.

[`refresh`]{#attr-meta-http-equiv-keyword-refresh .dfn
dfn-for="meta/http-equiv" dfn-type="attr-value"}

[Refresh](#attr-meta-http-equiv-refresh){#pragma-directives:attr-meta-http-equiv-refresh}

Acts as a timed redirect.

[`set-cookie`]{#attr-meta-http-equiv-keyword-set-cookie .dfn
dfn-for="meta/http-equiv" dfn-type="attr-value"}

No

[Set-Cookie](#attr-meta-http-equiv-set-cookie){#pragma-directives:attr-meta-http-equiv-set-cookie}

Has no effect.

[`x-ua-compatible`]{#attr-meta-http-equiv-keyword-x-ua-compatible .dfn
dfn-for="meta/http-equiv" dfn-type="attr-value"}

[X-UA-Compatible](#attr-meta-http-equiv-x-ua-compatible){#pragma-directives:attr-meta-http-equiv-x-ua-compatible}

In practice, encourages Internet Explorer to more closely follow the
specifications.

[`content-security-policy`]{#attr-meta-http-equiv-keyword-content-security-policy
.dfn dfn-for="meta/http-equiv" dfn-type="attr-value"}

[Content security
policy](#attr-meta-http-equiv-content-security-policy){#pragma-directives:attr-meta-http-equiv-content-security-policy}

[Enforces](https://w3c.github.io/webappsec-csp/#enforced){#pragma-directives:enforce-the-policy
x-internal="enforce-the-policy"} a [Content Security
Policy](https://w3c.github.io/webappsec-csp/#content-security-policy-object){#pragma-directives:content-security-policy
x-internal="content-security-policy"} on a
[`Document`{#pragma-directives:document}](dom.html#document).

When a
[`meta`{#pragma-directives:the-meta-element-2}](#the-meta-element)
element is [inserted into the
document](infrastructure.html#insert-an-element-into-a-document){#pragma-directives:insert-an-element-into-a-document},
if its
[`http-equiv`{#pragma-directives:attr-meta-http-equiv-2}](#attr-meta-http-equiv)
attribute is present and represents one of the above states, then the
user agent must run the algorithm appropriate for that state, as
described in the following list:

[Content language state]{#attr-meta-http-equiv-content-language .dfn}
(`http-equiv="`[`content-language`{#pragma-directives:attr-meta-http-equiv-keyword-content-language}](#attr-meta-http-equiv-keyword-content-language)`"`)

This feature is non-conforming. Authors are encouraged to use the
[`lang`{#pragma-directives:attr-lang}](dom.html#attr-lang) attribute
instead.

This pragma sets the [pragma-set default
language]{#pragma-set-default-language .dfn}. Until such a pragma is
successfully processed, there is no [pragma-set default
language](#pragma-set-default-language){#pragma-directives:pragma-set-default-language-2}.

1.  If the
    [`meta`{#pragma-directives:the-meta-element-3}](#the-meta-element)
    element has no
    [`content`{#pragma-directives:attr-meta-content}](#attr-meta-content)
    attribute, then return.

2.  If the element\'s
    [`content`{#pragma-directives:attr-meta-content-2}](#attr-meta-content)
    attribute contains a U+002C COMMA character (,) then return.

3.  Let `input`{.variable} be the value of the element\'s
    [`content`{#pragma-directives:attr-meta-content-3}](#attr-meta-content)
    attribute.

4.  Let `position`{.variable} point at the first character of
    `input`{.variable}.

5.  [Skip ASCII
    whitespace](https://infra.spec.whatwg.org/#skip-ascii-whitespace){#pragma-directives:skip-ascii-whitespace
    x-internal="skip-ascii-whitespace"} within `input`{.variable} given
    `position`{.variable}.

6.  [Collect a sequence of code
    points](https://infra.spec.whatwg.org/#collect-a-sequence-of-code-points){#pragma-directives:collect-a-sequence-of-code-points
    x-internal="collect-a-sequence-of-code-points"} that are not [ASCII
    whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#pragma-directives:space-characters
    x-internal="space-characters"} from `input`{.variable} given
    `position`{.variable}.

7.  Let `candidate`{.variable} be the string that resulted from the
    previous step.

8.  If `candidate`{.variable} is the empty string, return.

9.  Set the [pragma-set default
    language](#pragma-set-default-language){#pragma-directives:pragma-set-default-language-3}
    to `candidate`{.variable}.

    If the value consists of multiple space-separated tokens, tokens
    after the first are ignored.

This pragma is almost, but not quite, entirely unlike the HTTP
\`[`Content-Language`{#pragma-directives:http-content-language}](https://httpwg.org/specs/rfc7231.html#header.content-language){x-internal="http-content-language"}\`
header of the same name. [\[HTTP\]](references.html#refsHTTP)

[Encoding declaration state]{#attr-meta-http-equiv-content-type .dfn}
(`http-equiv="`[`content-type`{#pragma-directives:attr-meta-http-equiv-keyword-content-type}](#attr-meta-http-equiv-keyword-content-type)`"`)

The [Encoding declaration
state](#attr-meta-http-equiv-content-type){#pragma-directives:attr-meta-http-equiv-content-type-2}
is just an alternative form of setting the
[`charset`{#pragma-directives:attr-meta-charset-2}](#attr-meta-charset)
attribute: it is a [character encoding
declaration](#character-encoding-declaration){#pragma-directives:character-encoding-declaration}.
This state\'s user agent requirements are all handled by the parsing
section of the specification.

For [`meta`{#pragma-directives:the-meta-element-4}](#the-meta-element)
elements with an
[`http-equiv`{#pragma-directives:attr-meta-http-equiv-3}](#attr-meta-http-equiv)
attribute in the [Encoding declaration
state](#attr-meta-http-equiv-content-type){#pragma-directives:attr-meta-http-equiv-content-type-3},
the
[`content`{#pragma-directives:attr-meta-content-4}](#attr-meta-content)
attribute must have a value that is an [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#pragma-directives:ascii-case-insensitive
x-internal="ascii-case-insensitive"} match for a string that consists
of: \"`text/html;`\", optionally followed by any number of [ASCII
whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#pragma-directives:space-characters-2
x-internal="space-characters"}, followed by \"`charset=utf-8`\".

A document must not contain both a
[`meta`{#pragma-directives:the-meta-element-5}](#the-meta-element)
element with an
[`http-equiv`{#pragma-directives:attr-meta-http-equiv-4}](#attr-meta-http-equiv)
attribute in the [Encoding declaration
state](#attr-meta-http-equiv-content-type){#pragma-directives:attr-meta-http-equiv-content-type-4}
and a [`meta`{#pragma-directives:the-meta-element-6}](#the-meta-element)
element with the
[`charset`{#pragma-directives:attr-meta-charset-3}](#attr-meta-charset)
attribute present.

The [Encoding declaration
state](#attr-meta-http-equiv-content-type){#pragma-directives:attr-meta-http-equiv-content-type-5}
may be used in [HTML
documents](https://dom.spec.whatwg.org/#html-document){#pragma-directives:html-documents
x-internal="html-documents"}, but elements with an
[`http-equiv`{#pragma-directives:attr-meta-http-equiv-5}](#attr-meta-http-equiv)
attribute in that state must not be used in [XML
documents](https://dom.spec.whatwg.org/#xml-document){#pragma-directives:xml-documents
x-internal="xml-documents"}.

[Default style state]{#attr-meta-http-equiv-default-style .dfn}
(`http-equiv="`[`default-style`{#pragma-directives:attr-meta-http-equiv-keyword-default-style}](#attr-meta-http-equiv-keyword-default-style)`"`)

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

This pragma sets the
[name](https://drafts.csswg.org/cssom/#css-style-sheet-set-name){#pragma-directives:css-style-sheet-set-name-2
x-internal="css-style-sheet-set-name"} of the default [CSS style sheet
set](https://drafts.csswg.org/cssom/#css-style-sheet-set){#pragma-directives:css-style-sheet-set-2
x-internal="css-style-sheet-set"}.

1.  If the
    [`meta`{#pragma-directives:the-meta-element-7}](#the-meta-element)
    element has no
    [`content`{#pragma-directives:attr-meta-content-5}](#attr-meta-content)
    attribute, or if that attribute\'s value is the empty string, then
    return.

2.  [Change the preferred CSS style sheet set
    name](https://drafts.csswg.org/cssom/#change-the-preferred-css-style-sheet-set-name){#pragma-directives:change-the-preferred-css-style-sheet-set-name
    x-internal="change-the-preferred-css-style-sheet-set-name"} with the
    name being the value of the element\'s
    [`content`{#pragma-directives:attr-meta-content-6}](#attr-meta-content)
    attribute. [\[CSSOM\]](references.html#refsCSSOM)

[Refresh state]{#attr-meta-http-equiv-refresh .dfn}
(`http-equiv="`[`refresh`{#pragma-directives:attr-meta-http-equiv-keyword-refresh}](#attr-meta-http-equiv-keyword-refresh)`"`)

This pragma acts as a timed redirect.

A [`Document`{#pragma-directives:document-2}](dom.html#document) object
has an associated [will declaratively
refresh]{#will-declaratively-refresh .dfn dfn-for="Document"} (a
boolean). It is initially false.

1.  If the
    [`meta`{#pragma-directives:the-meta-element-8}](#the-meta-element)
    element has no
    [`content`{#pragma-directives:attr-meta-content-7}](#attr-meta-content)
    attribute, or if that attribute\'s value is the empty string, then
    return.

2.  Let `input`{.variable} be the value of the element\'s
    [`content`{#pragma-directives:attr-meta-content-8}](#attr-meta-content)
    attribute.

3.  Run the [shared declarative refresh
    steps](#shared-declarative-refresh-steps){#pragma-directives:shared-declarative-refresh-steps}
    with the
    [`meta`{#pragma-directives:the-meta-element-9}](#the-meta-element)
    element\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#pragma-directives:node-document
    x-internal="node-document"}, `input`{.variable}, and the
    [`meta`{#pragma-directives:the-meta-element-10}](#the-meta-element)
    element.

The [shared declarative refresh steps]{#shared-declarative-refresh-steps
.dfn}, given a
[`Document`{#pragma-directives:document-3}](dom.html#document) object
`document`{.variable}, string `input`{.variable}, and optionally a
[`meta`{#pragma-directives:the-meta-element-11}](#the-meta-element)
element `meta`{.variable}, are as follows:

1.  If `document`{.variable}\'s [will declaratively
    refresh](#will-declaratively-refresh){#pragma-directives:will-declaratively-refresh}
    is true, then return.

2.  Let `position`{.variable} point at the first [code
    point](https://infra.spec.whatwg.org/#code-point){#pragma-directives:code-point
    x-internal="code-point"} of `input`{.variable}.

3.  [Skip ASCII
    whitespace](https://infra.spec.whatwg.org/#skip-ascii-whitespace){#pragma-directives:skip-ascii-whitespace-2
    x-internal="skip-ascii-whitespace"} within `input`{.variable} given
    `position`{.variable}.

4.  Let `time`{.variable} be 0.

5.  [Collect a sequence of code
    points](https://infra.spec.whatwg.org/#collect-a-sequence-of-code-points){#pragma-directives:collect-a-sequence-of-code-points-2
    x-internal="collect-a-sequence-of-code-points"} that are [ASCII
    digits](https://infra.spec.whatwg.org/#ascii-digit){#pragma-directives:ascii-digits
    x-internal="ascii-digits"} from `input`{.variable} given
    `position`{.variable}, and let the result be
    `timeString`{.variable}.

6.  If `timeString`{.variable} is the empty string, then:

    1.  If the [code
        point](https://infra.spec.whatwg.org/#code-point){#pragma-directives:code-point-2
        x-internal="code-point"} in `input`{.variable} pointed to by
        `position`{.variable} is not U+002E (.), then return.

7.  Otherwise, set `time`{.variable} to the result of parsing
    `timeString`{.variable} using the [rules for parsing non-negative
    integers](common-microsyntaxes.html#rules-for-parsing-non-negative-integers){#pragma-directives:rules-for-parsing-non-negative-integers}.

8.  [Collect a sequence of code
    points](https://infra.spec.whatwg.org/#collect-a-sequence-of-code-points){#pragma-directives:collect-a-sequence-of-code-points-3
    x-internal="collect-a-sequence-of-code-points"} that are [ASCII
    digits](https://infra.spec.whatwg.org/#ascii-digit){#pragma-directives:ascii-digits-2
    x-internal="ascii-digits"} and U+002E FULL STOP characters (.) from
    `input`{.variable} given `position`{.variable}. Ignore any collected
    characters.

9.  Let `urlRecord`{.variable} be `document`{.variable}\'s
    [URL](https://dom.spec.whatwg.org/#concept-document-url){#pragma-directives:the-document's-address
    x-internal="the-document's-address"}.

10. If `position`{.variable} is not past the end of `input`{.variable},
    then:

    1.  If the [code
        point](https://infra.spec.whatwg.org/#code-point){#pragma-directives:code-point-3
        x-internal="code-point"} in `input`{.variable} pointed to by
        `position`{.variable} is not U+003B (;), U+002C (,), or [ASCII
        whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#pragma-directives:space-characters-3
        x-internal="space-characters"}, then return.

    2.  [Skip ASCII
        whitespace](https://infra.spec.whatwg.org/#skip-ascii-whitespace){#pragma-directives:skip-ascii-whitespace-3
        x-internal="skip-ascii-whitespace"} within `input`{.variable}
        given `position`{.variable}.

    3.  If the [code
        point](https://infra.spec.whatwg.org/#code-point){#pragma-directives:code-point-4
        x-internal="code-point"} in `input`{.variable} pointed to by
        `position`{.variable} is U+003B (;) or U+002C (,), then advance
        `position`{.variable} to the next [code
        point](https://infra.spec.whatwg.org/#code-point){#pragma-directives:code-point-5
        x-internal="code-point"}.

    4.  [Skip ASCII
        whitespace](https://infra.spec.whatwg.org/#skip-ascii-whitespace){#pragma-directives:skip-ascii-whitespace-4
        x-internal="skip-ascii-whitespace"} within `input`{.variable}
        given `position`{.variable}.

11. If `position`{.variable} is not past the end of `input`{.variable},
    then:

    1.  Let `urlString`{.variable} be the substring of
        `input`{.variable} from the [code
        point](https://infra.spec.whatwg.org/#code-point){#pragma-directives:code-point-6
        x-internal="code-point"} at `position`{.variable} to the end of
        the string.

    2.  If the [code
        point](https://infra.spec.whatwg.org/#code-point){#pragma-directives:code-point-7
        x-internal="code-point"} in `input`{.variable} pointed to by
        `position`{.variable} is U+0055 (U) or U+0075 (u), then advance
        `position`{.variable} to the next [code
        point](https://infra.spec.whatwg.org/#code-point){#pragma-directives:code-point-8
        x-internal="code-point"}. Otherwise, jump to the step labeled
        *skip quotes*.

    3.  If the [code
        point](https://infra.spec.whatwg.org/#code-point){#pragma-directives:code-point-9
        x-internal="code-point"} in `input`{.variable} pointed to by
        `position`{.variable} is U+0052 (R) or U+0072 (r), then advance
        `position`{.variable} to the next [code
        point](https://infra.spec.whatwg.org/#code-point){#pragma-directives:code-point-10
        x-internal="code-point"}. Otherwise, jump to the step labeled
        *parse*.

    4.  If the [code
        point](https://infra.spec.whatwg.org/#code-point){#pragma-directives:code-point-11
        x-internal="code-point"} in `input`{.variable} pointed to by
        `position`{.variable} is U+004C (L) or U+006C (l), then advance
        `position`{.variable} to the next [code
        point](https://infra.spec.whatwg.org/#code-point){#pragma-directives:code-point-12
        x-internal="code-point"}. Otherwise, jump to the step labeled
        *parse*.

    5.  [Skip ASCII
        whitespace](https://infra.spec.whatwg.org/#skip-ascii-whitespace){#pragma-directives:skip-ascii-whitespace-5
        x-internal="skip-ascii-whitespace"} within `input`{.variable}
        given `position`{.variable}.

    6.  If the [code
        point](https://infra.spec.whatwg.org/#code-point){#pragma-directives:code-point-13
        x-internal="code-point"} in `input`{.variable} pointed to by
        `position`{.variable} is U+003D (=), then advance
        `position`{.variable} to the next [code
        point](https://infra.spec.whatwg.org/#code-point){#pragma-directives:code-point-14
        x-internal="code-point"}. Otherwise, jump to the step labeled
        *parse*.

    7.  [Skip ASCII
        whitespace](https://infra.spec.whatwg.org/#skip-ascii-whitespace){#pragma-directives:skip-ascii-whitespace-6
        x-internal="skip-ascii-whitespace"} within `input`{.variable}
        given `position`{.variable}.

    8.  *Skip quotes*: If the [code
        point](https://infra.spec.whatwg.org/#code-point){#pragma-directives:code-point-15
        x-internal="code-point"} in `input`{.variable} pointed to by
        `position`{.variable} is U+0027 (\') or U+0022 (\"), then let
        `quote`{.variable} be that [code
        point](https://infra.spec.whatwg.org/#code-point){#pragma-directives:code-point-16
        x-internal="code-point"}, and advance `position`{.variable} to
        the next [code
        point](https://infra.spec.whatwg.org/#code-point){#pragma-directives:code-point-17
        x-internal="code-point"}. Otherwise, let `quote`{.variable} be
        the empty string.

    9.  Set `urlString`{.variable} to the substring of
        `input`{.variable} from the [code
        point](https://infra.spec.whatwg.org/#code-point){#pragma-directives:code-point-18
        x-internal="code-point"} at `position`{.variable} to the end of
        the string.

    10. If `quote`{.variable} is not the empty string, and there is a
        [code
        point](https://infra.spec.whatwg.org/#code-point){#pragma-directives:code-point-19
        x-internal="code-point"} in `urlString`{.variable} equal to
        `quote`{.variable}, then truncate `urlString`{.variable} at that
        [code
        point](https://infra.spec.whatwg.org/#code-point){#pragma-directives:code-point-20
        x-internal="code-point"}, so that it and all subsequent [code
        points](https://infra.spec.whatwg.org/#code-point){#pragma-directives:code-point-21
        x-internal="code-point"} are removed.

    11. *Parse*: Set `urlRecord`{.variable} to the result of
        [encoding-parsing a
        URL](urls-and-fetching.html#encoding-parsing-a-url){#pragma-directives:encoding-parsing-a-url}
        given `urlString`{.variable}, relative to `document`{.variable}.

    12. If `urlRecord`{.variable} is failure, then return.

12. Set `document`{.variable}\'s [will declaratively
    refresh](#will-declaratively-refresh){#pragma-directives:will-declaratively-refresh-2}
    to true.

13. Perform one or more of the following steps:

    - After the refresh has come due (as defined below), if the user has
      not canceled the redirect and, if `meta`{.variable} is given,
      `document`{.variable}\'s [active sandboxing flag
      set](browsers.html#active-sandboxing-flag-set){#pragma-directives:active-sandboxing-flag-set}
      does not have the [sandboxed automatic features browsing context
      flag](browsers.html#sandboxed-automatic-features-browsing-context-flag){#pragma-directives:sandboxed-automatic-features-browsing-context-flag}
      set, then
      [navigate](browsing-the-web.html#navigate){#pragma-directives:navigate}
      `document`{.variable}\'s [node
      navigable](document-sequences.html#node-navigable){#pragma-directives:node-navigable}
      to `urlRecord`{.variable} using `document`{.variable}, with
      *[historyHandling](browsing-the-web.html#navigation-hh)* set to
      \"[`replace`{#pragma-directives:navigationhistorybehavior-replace}](browsing-the-web.html#navigationhistorybehavior-replace)\".

      For the purposes of the previous paragraph, a refresh is said to
      have come due as soon as the *later* of the following two
      conditions occurs:

      - At least `time`{.variable} seconds have elapsed since
        `document`{.variable}\'s [completely loaded
        time](document-lifecycle.html#completely-loaded-time){#pragma-directives:completely-loaded-time},
        adjusted to take into account user or user agent preferences.
      - If `meta`{.variable} is given, at least `time`{.variable}
        seconds have elapsed since `meta`{.variable} was [inserted into
        the
        document](infrastructure.html#insert-an-element-into-a-document){#pragma-directives:insert-an-element-into-a-document-2}
        `document`{.variable}, adjusted to take into account user or
        user agent preferences.

      It is important to use `document`{.variable} here, and not
      `meta`{.variable}\'s [node
      document](https://dom.spec.whatwg.org/#concept-node-document){#pragma-directives:node-document-2
      x-internal="node-document"}, as that might have changed between
      the initial set of steps and the refresh coming due and
      `meta`{.variable} is not always given (in case of the HTTP
      \`[`Refresh`{#pragma-directives:refresh}](document-lifecycle.html#refresh)\`
      header).

    - Provide the user with an interface that, when selected,
      [navigates](browsing-the-web.html#navigate){#pragma-directives:navigate-2}
      `document`{.variable}\'s [node
      navigable](document-sequences.html#node-navigable){#pragma-directives:node-navigable-2}
      to `urlRecord`{.variable} using `document`{.variable}.

    - Do nothing.

    In addition, the user agent may, as with anything, inform the user
    of any and all aspects of its operation, including the state of any
    timers, the destinations of any timed redirects, and so forth.

For [`meta`{#pragma-directives:the-meta-element-12}](#the-meta-element)
elements with an
[`http-equiv`{#pragma-directives:attr-meta-http-equiv-6}](#attr-meta-http-equiv)
attribute in the [Refresh
state](#attr-meta-http-equiv-refresh){#pragma-directives:attr-meta-http-equiv-refresh-2},
the
[`content`{#pragma-directives:attr-meta-content-9}](#attr-meta-content)
attribute must have a value consisting either of:

- just a [valid non-negative
  integer](common-microsyntaxes.html#valid-non-negative-integer){#pragma-directives:valid-non-negative-integer},
  or
- a [valid non-negative
  integer](common-microsyntaxes.html#valid-non-negative-integer){#pragma-directives:valid-non-negative-integer-2},
  followed by a U+003B SEMICOLON character (;), followed by one or more
  [ASCII
  whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#pragma-directives:space-characters-4
  x-internal="space-characters"}, followed by a substring that is an
  [ASCII
  case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#pragma-directives:ascii-case-insensitive-2
  x-internal="ascii-case-insensitive"} match for the string \"`URL`\",
  followed by a U+003D EQUALS SIGN character (=), followed by a [valid
  URL
  string](https://url.spec.whatwg.org/#valid-url-string){#pragma-directives:valid-url-string
  x-internal="valid-url-string"} that does not start with a literal
  U+0027 APOSTROPHE (\') or U+0022 QUOTATION MARK (\") character.

In the former case, the integer represents a number of seconds before
the page is to be reloaded; in the latter case the integer represents a
number of seconds before the page is to be replaced by the page at the
given
[URL](https://url.spec.whatwg.org/#concept-url){#pragma-directives:url
x-internal="url"}.

::: example
A news organization\'s front page could include the following markup in
the page\'s
[`head`{#pragma-directives:the-head-element}](#the-head-element)
element, to ensure that the page automatically reloads from the server
every five minutes:

``` html
<meta http-equiv="Refresh" content="300">
```
:::

::: example
A sequence of pages could be used as an automated slide show by making
each page refresh to the next page in the sequence, using markup such as
the following:

``` html
<meta http-equiv="Refresh" content="20; URL=page4.html">
```
:::

[Set-Cookie state]{#attr-meta-http-equiv-set-cookie .dfn}
(`http-equiv="`[`set-cookie`{#pragma-directives:attr-meta-http-equiv-keyword-set-cookie}](#attr-meta-http-equiv-keyword-set-cookie)`"`)

This pragma is non-conforming and has no effect.

User agents are required to ignore this pragma.

[X-UA-Compatible state]{#attr-meta-http-equiv-x-ua-compatible .dfn}
(`http-equiv="`[`x-ua-compatible`{#pragma-directives:attr-meta-http-equiv-keyword-x-ua-compatible}](#attr-meta-http-equiv-keyword-x-ua-compatible)`"`)

In practice, this pragma encourages Internet Explorer to more closely
follow the specifications.

For [`meta`{#pragma-directives:the-meta-element-13}](#the-meta-element)
elements with an
[`http-equiv`{#pragma-directives:attr-meta-http-equiv-7}](#attr-meta-http-equiv)
attribute in the [X-UA-Compatible
state](#attr-meta-http-equiv-x-ua-compatible){#pragma-directives:attr-meta-http-equiv-x-ua-compatible-2},
the
[`content`{#pragma-directives:attr-meta-content-10}](#attr-meta-content)
attribute must have a value that is an [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#pragma-directives:ascii-case-insensitive-3
x-internal="ascii-case-insensitive"} match for the string \"`IE=edge`\".

User agents are required to ignore this pragma.

[Content security policy
state]{#attr-meta-http-equiv-content-security-policy .dfn export=""}
(`http-equiv="`[`content-security-policy`{#pragma-directives:attr-meta-http-equiv-keyword-content-security-policy}](#attr-meta-http-equiv-keyword-content-security-policy)`"`)

This pragma
[enforces](https://w3c.github.io/webappsec-csp/#enforced){#pragma-directives:enforce-the-policy-2
x-internal="enforce-the-policy"} a [Content Security
Policy](https://w3c.github.io/webappsec-csp/#content-security-policy-object){#pragma-directives:content-security-policy-2
x-internal="content-security-policy"} on a
[`Document`{#pragma-directives:document-4}](dom.html#document).
[\[CSP\]](references.html#refsCSP)

1.  If the
    [`meta`{#pragma-directives:the-meta-element-14}](#the-meta-element)
    element is not a child of a
    [`head`{#pragma-directives:the-head-element-2}](#the-head-element)
    element, return.

2.  If the
    [`meta`{#pragma-directives:the-meta-element-15}](#the-meta-element)
    element has no
    [`content`{#pragma-directives:attr-meta-content-11}](#attr-meta-content)
    attribute, or if that attribute\'s value is the empty string, then
    return.

3.  Let `policy`{.variable} be the result of executing Content Security
    Policy\'s [parse a serialized Content Security
    Policy](https://w3c.github.io/webappsec-csp/#parse-serialized-policy){#pragma-directives:parse-a-serialized-content-security-policy
    x-internal="parse-a-serialized-content-security-policy"} algorithm
    on the
    [`meta`{#pragma-directives:the-meta-element-16}](#the-meta-element)
    element\'s
    [`content`{#pragma-directives:attr-meta-content-12}](#attr-meta-content)
    attribute\'s value, with a source of \"meta\", and a disposition of
    \"enforce\".

4.  Remove all occurrences of the
    [`report-uri`{#pragma-directives:report-uri-directive}](https://w3c.github.io/webappsec-csp/#report-uri){x-internal="report-uri-directive"},
    [`frame-ancestors`{#pragma-directives:frame-ancestors-directive}](https://w3c.github.io/webappsec-csp/#frame-ancestors){x-internal="frame-ancestors-directive"},
    and
    [`sandbox`{#pragma-directives:sandbox-directive}](https://w3c.github.io/webappsec-csp/#sandbox){x-internal="sandbox-directive"}
    [directives](https://w3c.github.io/webappsec-csp/#directives){#pragma-directives:content-security-policy-directive
    x-internal="content-security-policy-directive"} from
    `policy`{.variable}.

5.  [Enforce the
    policy](https://w3c.github.io/webappsec-csp/#enforced){#pragma-directives:enforce-the-policy-3
    x-internal="enforce-the-policy"} `policy`{.variable}.

For [`meta`{#pragma-directives:the-meta-element-17}](#the-meta-element)
elements with an
[`http-equiv`{#pragma-directives:attr-meta-http-equiv-8}](#attr-meta-http-equiv)
attribute in the [Content security policy
state](#attr-meta-http-equiv-content-security-policy){#pragma-directives:attr-meta-http-equiv-content-security-policy-2},
the
[`content`{#pragma-directives:attr-meta-content-13}](#attr-meta-content)
attribute must have a value consisting of a [valid Content Security
Policy](https://w3c.github.io/webappsec-csp/#grammardef-serialized-policy){#pragma-directives:content-security-policy-syntax
x-internal="content-security-policy-syntax"}, but must not contain any
[`report-uri`{#pragma-directives:report-uri-directive-2}](https://w3c.github.io/webappsec-csp/#report-uri){x-internal="report-uri-directive"},
[`frame-ancestors`{#pragma-directives:frame-ancestors-directive-2}](https://w3c.github.io/webappsec-csp/#frame-ancestors){x-internal="frame-ancestors-directive"},
or
[`sandbox`{#pragma-directives:sandbox-directive-2}](https://w3c.github.io/webappsec-csp/#sandbox){x-internal="sandbox-directive"}
[directives](https://w3c.github.io/webappsec-csp/#directives){#pragma-directives:content-security-policy-directive-2
x-internal="content-security-policy-directive"}. The [Content Security
Policy](https://w3c.github.io/webappsec-csp/#content-security-policy-object){#pragma-directives:content-security-policy-3
x-internal="content-security-policy"} given in the
[`content`{#pragma-directives:attr-meta-content-14}](#attr-meta-content)
attribute will be
[enforced](https://w3c.github.io/webappsec-csp/#enforced){#pragma-directives:enforce-the-policy-4
x-internal="enforce-the-policy"} upon the current document.
[\[CSP\]](references.html#refsCSP)

At the time of inserting the
[`meta`{#pragma-directives:the-meta-element-18}](#the-meta-element)
element to the document, it is possible that some resources have already
been fetched. For example, images might be stored in the [list of
available
images](images.html#list-of-available-images){#pragma-directives:list-of-available-images}
prior to dynamically inserting a
[`meta`{#pragma-directives:the-meta-element-19}](#the-meta-element)
element with an
[`http-equiv`{#pragma-directives:attr-meta-http-equiv-9}](#attr-meta-http-equiv)
attribute in the [Content security policy
state](#attr-meta-http-equiv-content-security-policy){#pragma-directives:attr-meta-http-equiv-content-security-policy-3}.
Resources that have already been fetched are not guaranteed to be
blocked by a [Content Security
Policy](https://w3c.github.io/webappsec-csp/#content-security-policy-object){#pragma-directives:content-security-policy-4
x-internal="content-security-policy"} that\'s
[enforced](https://w3c.github.io/webappsec-csp/#enforced){#pragma-directives:enforce-the-policy-5
x-internal="enforce-the-policy"} late.

::: example
A page might choose to mitigate the risk of cross-site scripting attacks
by preventing the execution of inline JavaScript, as well as blocking
all plugin content, using a policy such as the following:

``` html
<meta http-equiv="Content-Security-Policy" content="script-src 'self'; object-src 'none'">
```
:::

There must not be more than one
[`meta`{#pragma-directives:the-meta-element-20}](#the-meta-element)
element with any particular state in the document at a time.

##### [4.2.5.4]{.secno} Specifying the document\'s character encoding[](#charset){.self-link} {#charset}

A [character encoding declaration]{#character-encoding-declaration .dfn}
is a mechanism by which the [character
encoding](https://encoding.spec.whatwg.org/#encoding){#charset:encoding
x-internal="encoding"} used to store or transmit a document is
specified.

The Encoding standard requires use of the
[UTF-8](https://encoding.spec.whatwg.org/#utf-8){#charset:utf-8
x-internal="utf-8"} [character
encoding](https://encoding.spec.whatwg.org/#encoding){#charset:encoding-2
x-internal="encoding"} and requires use of the \"`utf-8`\" [encoding
label](https://encoding.spec.whatwg.org/#label){#charset:encoding-label
x-internal="encoding-label"} to identify it. Those requirements
necessitate that the document\'s [character encoding
declaration](#character-encoding-declaration){#charset:character-encoding-declaration},
if it exists, specifies an [encoding
label](https://encoding.spec.whatwg.org/#label){#charset:encoding-label-2
x-internal="encoding-label"} using an [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#charset:ascii-case-insensitive
x-internal="ascii-case-insensitive"} match for \"`utf-8`\". Regardless
of whether a [character encoding
declaration](#character-encoding-declaration){#charset:character-encoding-declaration-2}
is present or not, the actual [character
encoding](https://dom.spec.whatwg.org/#concept-document-encoding){#charset:document's-character-encoding
x-internal="document's-character-encoding"} used to encode the document
must be
[UTF-8](https://encoding.spec.whatwg.org/#utf-8){#charset:utf-8-2
x-internal="utf-8"}. [\[ENCODING\]](references.html#refsENCODING)

To enforce the above rules, authoring tools must default to using
[UTF-8](https://encoding.spec.whatwg.org/#utf-8){#charset:utf-8-3
x-internal="utf-8"} for newly-created documents.

The following restrictions also apply:

- The character encoding declaration must be serialized without the use
  of [character
  references](syntax.html#syntax-charref){#charset:syntax-charref} or
  character escapes of any kind.
- [[The element containing the character encoding declaration must be
  serialized completely within the first 1024 bytes of the
  document.]{#charset512}]{#charset1024}

In addition, due to a number of restrictions on
[`meta`{#charset:the-meta-element}](#the-meta-element) elements, there
can only be one
[`meta`{#charset:the-meta-element-2}](#the-meta-element)-based character
encoding declaration per document.

If an [HTML
document](https://dom.spec.whatwg.org/#html-document){#charset:html-documents
x-internal="html-documents"} does not start with a BOM, and its
[encoding](https://encoding.spec.whatwg.org/#encoding){#charset:encoding-3
x-internal="encoding"} is not explicitly given by [Content-Type
metadata](urls-and-fetching.html#content-type){#charset:content-type},
and the document is not [an `iframe` `srcdoc`
document](iframe-embed-object.html#an-iframe-srcdoc-document){#charset:an-iframe-srcdoc-document},
then the encoding must be specified using a
[`meta`{#charset:the-meta-element-3}](#the-meta-element) element with a
[`charset`{#charset:attr-meta-charset}](#attr-meta-charset) attribute or
a [`meta`{#charset:the-meta-element-4}](#the-meta-element) element with
an [`http-equiv`{#charset:attr-meta-http-equiv}](#attr-meta-http-equiv)
attribute in the [Encoding declaration
state](#attr-meta-http-equiv-content-type){#charset:attr-meta-http-equiv-content-type}.

::: note
A character encoding declaration is required (either in the
[Content-Type
metadata](urls-and-fetching.html#content-type){#charset:content-type-2}
or explicitly in the file) even when all characters are in the ASCII
range, because a character encoding is needed to process non-ASCII
characters entered by the user in forms, in URLs generated by scripts,
and so forth.

Using non-UTF-8 encodings can have unexpected results on form submission
and URL encodings, which use the [document\'s character
encoding](https://dom.spec.whatwg.org/#concept-document-encoding){#charset:document's-character-encoding-2
x-internal="document's-character-encoding"} by default.
:::

If the document is [an `iframe` `srcdoc`
document](iframe-embed-object.html#an-iframe-srcdoc-document){#charset:an-iframe-srcdoc-document-2},
the document must not have a [character encoding
declaration](#character-encoding-declaration){#charset:character-encoding-declaration-3}.
(In this case, the source is already decoded, since it is part of the
document that contained the
[`iframe`{#charset:the-iframe-element}](iframe-embed-object.html#the-iframe-element).)

In XML, the XML declaration should be used for inline character encoding
information, if necessary.

::: example
In HTML, to declare that the character encoding is
[UTF-8](https://encoding.spec.whatwg.org/#utf-8){#charset:utf-8-4
x-internal="utf-8"}, the author could include the following markup near
the top of the document (in the
[`head`{#charset:the-head-element}](#the-head-element) element):

``` html
<meta charset="utf-8">
```

In XML, the XML declaration would be used instead, at the very top of
the markup:

``` html
<?xml version="1.0" encoding="utf-8"?>
```
:::

#### [4.2.6]{.secno} The [`style`]{.dfn dfn-type="element"} element[](#the-style-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/style](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/style "The <style> HTML element contains style information for a document, or part of a document. It contains CSS, which is applied to the contents of the document containing the <style> element.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera3.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer3+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLStyleElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLStyleElement "The HTMLStyleElement interface represents a <style> element. It inherits properties and methods from its parent, HTMLElement.")

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
:::::

[Categories](dom.html#concept-element-categories){#the-style-element:concept-element-categories}:
:   [Metadata
    content](dom.html#metadata-content-2){#the-style-element:metadata-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-style-element:concept-element-contexts}:
:   Where [metadata
    content](dom.html#metadata-content-2){#the-style-element:metadata-content-2-2}
    is expected.
:   In a
    [`noscript`{#the-style-element:the-noscript-element}](scripting.html#the-noscript-element)
    element that is a child of a
    [`head`{#the-style-element:the-head-element}](#the-head-element)
    element.

[Content model](dom.html#concept-element-content-model){#the-style-element:concept-element-content-model}:
:   [Text](dom.html#text-content){#the-style-element:text-content} that
    gives a [conformant style
    sheet](https://drafts.csswg.org/css-syntax/#conform-classes){#the-style-element:conformant-style-sheet
    x-internal="conformant-style-sheet"}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-style-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-style-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-style-element:global-attributes}
:   [`media`{#the-style-element:attr-style-media}](#attr-style-media)
    --- Applicable media
:   [`blocking`{#the-style-element:attr-style-blocking}](#attr-style-blocking)
    --- Whether the element is [potentially
    render-blocking](urls-and-fetching.html#potentially-render-blocking){#the-style-element:potentially-render-blocking}
:   Also, the
    [`title`{#the-style-element:attr-style-title}](#attr-style-title)
    attribute [has special
    semantics](#attr-style-title){#the-style-element:attr-style-title-2}
    on this element: [CSS style sheet set
    name](https://drafts.csswg.org/cssom/#css-style-sheet-set-name){#the-style-element:css-style-sheet-set-name
    x-internal="css-style-sheet-set-name"}

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-style-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-style).
:   [For implementers](https://w3c.github.io/html-aam/#el-style).

[DOM interface](dom.html#concept-element-dom){#the-style-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLStyleElement : HTMLElement {
      [HTMLConstructor] constructor();

      attribute boolean disabled;
      [CEReactions] attribute DOMString media;
      [SameObject, PutForwards=value] readonly attribute DOMTokenList blocking;

      // also has obsolete members
    };
    HTMLStyleElement includes LinkStyle;
    ```

The [`style`{#the-style-element:the-style-element}](#the-style-element)
element allows authors to embed CSS style sheets in their documents. The
[`style`{#the-style-element:the-style-element-2}](#the-style-element)
element is one of several inputs to the styling processing model. The
element does not
[represent](dom.html#represents){#the-style-element:represents} content
for the user.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLStyleElement/disabled](https://developer.mozilla.org/en-US/docs/Web/API/HTMLStyleElement/disabled "The HTMLStyleElement.disabled property can be used to get and set whether the stylesheet is disabled (true) or not (false).")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)13+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The [`disabled`]{#dom-style-disabled .dfn dfn-for="HTMLStyleElement"
dfn-type="attribute"} getter steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-style-element:this
    x-internal="this"} does not have an [associated CSS style
    sheet](https://drafts.csswg.org/cssom/#associated-css-style-sheet){#the-style-element:associated-css-style-sheet
    x-internal="associated-css-style-sheet"}, return false.

2.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-style-element:this-2
    x-internal="this"}\'s [associated CSS style
    sheet](https://drafts.csswg.org/cssom/#associated-css-style-sheet){#the-style-element:associated-css-style-sheet-2
    x-internal="associated-css-style-sheet"}\'s [disabled
    flag](https://drafts.csswg.org/cssom/#concept-css-style-sheet-disabled-flag){#the-style-element:concept-css-style-sheet-disabled-flag
    x-internal="concept-css-style-sheet-disabled-flag"} is set, return
    true.

3.  Return false.

The
[`disabled`{#the-style-element:dom-style-disabled-2}](#dom-style-disabled)
setter steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-style-element:this-3
    x-internal="this"} does not have an [associated CSS style
    sheet](https://drafts.csswg.org/cssom/#associated-css-style-sheet){#the-style-element:associated-css-style-sheet-3
    x-internal="associated-css-style-sheet"}, return.

2.  If the given value is true, set
    [this](https://webidl.spec.whatwg.org/#this){#the-style-element:this-4
    x-internal="this"}\'s [associated CSS style
    sheet](https://drafts.csswg.org/cssom/#associated-css-style-sheet){#the-style-element:associated-css-style-sheet-4
    x-internal="associated-css-style-sheet"}\'s [disabled
    flag](https://drafts.csswg.org/cssom/#concept-css-style-sheet-disabled-flag){#the-style-element:concept-css-style-sheet-disabled-flag-2
    x-internal="concept-css-style-sheet-disabled-flag"}. Otherwise,
    unset
    [this](https://webidl.spec.whatwg.org/#this){#the-style-element:this-5
    x-internal="this"}\'s [associated CSS style
    sheet](https://drafts.csswg.org/cssom/#associated-css-style-sheet){#the-style-element:associated-css-style-sheet-5
    x-internal="associated-css-style-sheet"}\'s [disabled
    flag](https://drafts.csswg.org/cssom/#concept-css-style-sheet-disabled-flag){#the-style-element:concept-css-style-sheet-disabled-flag-3
    x-internal="concept-css-style-sheet-disabled-flag"}.

::: example
Importantly,
[`disabled`{#the-style-element:dom-style-disabled-3}](#dom-style-disabled)
attribute assignments only take effect when the
[`style`{#the-style-element:the-style-element-3}](#the-style-element)
element has an [associated CSS style
sheet](https://drafts.csswg.org/cssom/#associated-css-style-sheet){#the-style-element:associated-css-style-sheet-6
x-internal="associated-css-style-sheet"}:

``` js
const style = document.createElement('style');
style.disabled = true;
style.textContent = 'body { background-color: red; }';
document.body.append(style);
console.log(style.disabled); // false
```
:::

The [`media`]{#attr-style-media .dfn dfn-for="style"
dfn-type="element-attr"} attribute says which media the styles apply to.
The value must be a [valid media query
list](common-microsyntaxes.html#valid-media-query-list){#the-style-element:valid-media-query-list}.
The user agent must apply the styles when the
[`media`{#the-style-element:attr-style-media-2}](#attr-style-media)
attribute\'s value [matches the
environment](common-microsyntaxes.html#matches-the-environment){#the-style-element:matches-the-environment}
and the other relevant conditions apply, and must not apply them
otherwise.

The styles might be further limited in scope, e.g. in CSS with the use
of `@media` blocks. This specification does not override such further
restrictions or requirements.

The default, if the
[`media`{#the-style-element:attr-style-media-3}](#attr-style-media)
attribute is omitted, is \"`all`\", meaning that by default styles apply
to all media.

The [`blocking`]{#attr-style-blocking .dfn dfn-for="style"
dfn-type="element-attr"} attribute is a [blocking
attribute](urls-and-fetching.html#blocking-attribute){#the-style-element:blocking-attribute}.

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

The [`title`]{#attr-style-title .dfn dfn-for="style"
dfn-type="element-attr"} attribute on
[`style`{#the-style-element:the-style-element-4}](#the-style-element)
elements defines [CSS style sheet
sets](https://drafts.csswg.org/cssom/#css-style-sheet-set){#the-style-element:css-style-sheet-set
x-internal="css-style-sheet-set"}. If the
[`style`{#the-style-element:the-style-element-5}](#the-style-element)
element has no
[`title`{#the-style-element:attr-style-title-3}](#attr-style-title)
attribute, then it has no title; the
[`title`{#the-style-element:attr-title}](dom.html#attr-title) attribute
of ancestors does not apply to the
[`style`{#the-style-element:the-style-element-6}](#the-style-element)
element. If the
[`style`{#the-style-element:the-style-element-7}](#the-style-element)
element is not [in a document
tree](https://dom.spec.whatwg.org/#in-a-document-tree){#the-style-element:in-a-document-tree
x-internal="in-a-document-tree"}, then the
[`title`{#the-style-element:attr-style-title-4}](#attr-style-title)
attribute is ignored. [\[CSSOM\]](references.html#refsCSSOM)

The [`title`{#the-style-element:attr-style-title-5}](#attr-style-title)
attribute on
[`style`{#the-style-element:the-style-element-8}](#the-style-element)
elements, like the
[`title`{#the-style-element:attr-link-title}](#attr-link-title)
attribute on
[`link`{#the-style-element:the-link-element}](#the-link-element)
elements, differs from the global
[`title`{#the-style-element:attr-title-2}](dom.html#attr-title)
attribute in that a
[`style`{#the-style-element:the-style-element-9}](#the-style-element)
block without a title does not inherit the title of the parent element:
it merely has no title.

The [child text
content](https://dom.spec.whatwg.org/#concept-child-text-content){#the-style-element:child-text-content
x-internal="child-text-content"} of a
[`style`{#the-style-element:the-style-element-10}](#the-style-element)
element must be that of a [conformant style
sheet](https://drafts.csswg.org/css-syntax/#conform-classes){#the-style-element:conformant-style-sheet-2
x-internal="conformant-style-sheet"}.

A [`style`{#the-style-element:the-style-element-11}](#the-style-element)
element is [implicitly potentially
render-blocking](urls-and-fetching.html#implicitly-potentially-render-blocking){#the-style-element:implicitly-potentially-render-blocking}
if the element was created by its [node
document](https://dom.spec.whatwg.org/#concept-node-document){#the-style-element:node-document
x-internal="node-document"}\'s parser.

------------------------------------------------------------------------

The user agent must run the [update a `style`
block](#update-a-style-block){#the-style-element:update-a-style-block}
algorithm whenever any of the following conditions occur:

- The element is popped off the [stack of open
  elements](parsing.html#stack-of-open-elements){#the-style-element:stack-of-open-elements}
  of an [HTML
  parser](parsing.html#html-parser){#the-style-element:html-parser} or
  [XML parser](xhtml.html#xml-parser){#the-style-element:xml-parser}.

- The element is not on the [stack of open
  elements](parsing.html#stack-of-open-elements){#the-style-element:stack-of-open-elements-2}
  of an [HTML
  parser](parsing.html#html-parser){#the-style-element:html-parser-2} or
  [XML parser](xhtml.html#xml-parser){#the-style-element:xml-parser-2},
  and it [becomes
  connected](infrastructure.html#becomes-connected){#the-style-element:becomes-connected}
  or
  [disconnected](infrastructure.html#becomes-disconnected){#the-style-element:becomes-disconnected}.

- The element\'s [children changed
  steps](https://dom.spec.whatwg.org/#concept-node-children-changed-ext){#the-style-element:children-changed-steps
  x-internal="children-changed-steps"} run.

The [update a `style` block]{#update-a-style-block .dfn export=""}
algorithm is as follows:

1.  Let `element`{.variable} be the
    [`style`{#the-style-element:the-style-element-12}](#the-style-element)
    element.

2.  If `element`{.variable} has an [associated CSS style
    sheet](https://drafts.csswg.org/cssom/#associated-css-style-sheet){#the-style-element:associated-css-style-sheet-7
    x-internal="associated-css-style-sheet"}, [remove the CSS style
    sheet](https://drafts.csswg.org/cssom/#remove-a-css-style-sheet){#the-style-element:remove-a-css-style-sheet
    x-internal="remove-a-css-style-sheet"} in question.

3.  If `element`{.variable} is not
    [connected](https://dom.spec.whatwg.org/#connected){#the-style-element:connected
    x-internal="connected"}, then return.

4.  If `element`{.variable}\'s
    [`type`{#the-style-element:attr-style-type}](obsolete.html#attr-style-type)
    attribute is present and its value is neither the empty string nor
    an [ASCII
    case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#the-style-element:ascii-case-insensitive
    x-internal="ascii-case-insensitive"} match for
    \"[`text/css`{#the-style-element:text/css}](indices.html#text/css)\",
    then return.

    In particular, a
    [`type`{#the-style-element:attr-style-type-2}](obsolete.html#attr-style-type)
    value with parameters, such as \"`text/css; charset=utf-8`\", will
    cause this algorithm to return early.

5.  If the [Should element\'s inline behavior be blocked by Content
    Security
    Policy?](https://w3c.github.io/webappsec-csp/#should-block-inline){#the-style-element:should-element's-inline-behavior-be-blocked-by-content-security-policy
    x-internal="should-element's-inline-behavior-be-blocked-by-content-security-policy"}
    algorithm returns \"`Blocked`\" when executed upon the
    [`style`{#the-style-element:the-style-element-13}](#the-style-element)
    element, \"`style`\", and the
    [`style`{#the-style-element:the-style-element-14}](#the-style-element)
    element\'s [child text
    content](https://dom.spec.whatwg.org/#concept-child-text-content){#the-style-element:child-text-content-2
    x-internal="child-text-content"}, then return.
    [\[CSP\]](references.html#refsCSP)

6.  [Create a CSS style
    sheet](https://drafts.csswg.org/cssom/#create-a-css-style-sheet){#the-style-element:create-a-css-style-sheet
    x-internal="create-a-css-style-sheet"} with the following
    properties:

    [type](https://drafts.csswg.org/cssom/#concept-css-style-sheet-type){#the-style-element:concept-css-style-sheet-type x-internal="concept-css-style-sheet-type"}
    :   [`text/css`{#the-style-element:text/css-2}](indices.html#text/css)

    [owner node](https://drafts.csswg.org/cssom/#concept-css-style-sheet-owner-node){#the-style-element:concept-css-style-sheet-owner-node x-internal="concept-css-style-sheet-owner-node"}
    :   `element`{.variable}

    [media](https://drafts.csswg.org/cssom/#concept-css-style-sheet-media){#the-style-element:concept-css-style-sheet-media x-internal="concept-css-style-sheet-media"}

    :   The
        [`media`{#the-style-element:attr-style-media-4}](#attr-style-media)
        attribute of `element`{.variable}.

        This is a reference to the (possibly absent at this time)
        attribute, rather than a copy of the attribute\'s current value.
        CSSOM defines what happens when the attribute is dynamically
        set, changed, or removed.

    [title](https://drafts.csswg.org/cssom/#concept-css-style-sheet-title){#the-style-element:concept-css-style-sheet-title x-internal="concept-css-style-sheet-title"}

    :   The
        [`title`{#the-style-element:attr-style-title-6}](#attr-style-title)
        attribute of `element`{.variable}, if `element`{.variable} is
        [in a document
        tree](https://dom.spec.whatwg.org/#in-a-document-tree){#the-style-element:in-a-document-tree-2
        x-internal="in-a-document-tree"}, or the empty string otherwise.

        Again, this is a *reference* to the attribute.

    [alternate flag](https://drafts.csswg.org/cssom/#concept-css-style-sheet-alternate-flag){#the-style-element:concept-css-style-sheet-alternate-flag x-internal="concept-css-style-sheet-alternate-flag"}
    :   Unset.

    [origin-clean flag](https://drafts.csswg.org/cssom/#concept-css-style-sheet-origin-clean-flag){#the-style-element:concept-css-style-sheet-origin-clean-flag x-internal="concept-css-style-sheet-origin-clean-flag"}
    :   Set.

    [location](https://drafts.csswg.org/cssom/#concept-css-style-sheet-location){#the-style-element:concept-css-style-sheet-location x-internal="concept-css-style-sheet-location"}\
    [parent CSS style sheet](https://drafts.csswg.org/cssom/#concept-css-style-sheet-parent-css-style-sheet){#the-style-element:concept-css-style-sheet-parent-css-style-sheet x-internal="concept-css-style-sheet-parent-css-style-sheet"}\
    [owner CSS rule](https://drafts.csswg.org/cssom/#concept-css-style-sheet-owner-css-rule){#the-style-element:concept-css-style-sheet-owner-css-rule x-internal="concept-css-style-sheet-owner-css-rule"}
    :   null

    [disabled flag](https://drafts.csswg.org/cssom/#concept-css-style-sheet-disabled-flag){#the-style-element:concept-css-style-sheet-disabled-flag-4 x-internal="concept-css-style-sheet-disabled-flag"}
    :   Left at its default value.

    [CSS rules](https://drafts.csswg.org/cssom/#concept-css-style-sheet-css-rules){#the-style-element:concept-css-style-sheet-css-rules x-internal="concept-css-style-sheet-css-rules"}

    :   Left uninitialized.

        This doesn\'t seem right. Presumably we should be using the
        element\'s [child text
        content](https://dom.spec.whatwg.org/#concept-child-text-content){#the-style-element:child-text-content-3
        x-internal="child-text-content"}? Tracked as [issue
        #2997](https://github.com/whatwg/html/issues/2997).

7.  If `element`{.variable} [contributes a script-blocking style
    sheet](#contributes-a-script-blocking-style-sheet){#the-style-element:contributes-a-script-blocking-style-sheet},
    [append](https://infra.spec.whatwg.org/#set-append){#the-style-element:set-append
    x-internal="set-append"} `element`{.variable} to its [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-style-element:node-document-2
    x-internal="node-document"}\'s [script-blocking style sheet
    set](#script-blocking-style-sheet-set){#the-style-element:script-blocking-style-sheet-set}.

8.  If `element`{.variable}\'s
    [`media`{#the-style-element:attr-style-media-5}](#attr-style-media)
    attribute\'s value [matches the
    environment](common-microsyntaxes.html#matches-the-environment){#the-style-element:matches-the-environment-2}
    and `element`{.variable} is [potentially
    render-blocking](urls-and-fetching.html#potentially-render-blocking){#the-style-element:potentially-render-blocking-2},
    then [block
    rendering](dom.html#block-rendering){#the-style-element:block-rendering}
    on `element`{.variable}.

Once the attempts to obtain the style sheet\'s [critical
subresources](infrastructure.html#critical-subresources){#the-style-element:critical-subresources},
if any, are complete, or, if the style sheet has no [critical
subresources](infrastructure.html#critical-subresources){#the-style-element:critical-subresources-2},
once the style sheet has been parsed and processed, the user agent must
run these steps:

Fetching the [critical
subresources](infrastructure.html#critical-subresources){#the-style-element:critical-subresources-3}
is not well-defined; probably [issue
#968](https://github.com/whatwg/html/issues/968) is the best resolution
for that. In the meantime, any [critical
subresource](infrastructure.html#critical-subresources){#the-style-element:critical-subresources-4}
[request](https://fetch.spec.whatwg.org/#concept-request){#the-style-element:concept-request
x-internal="concept-request"} should have its
[render-blocking](https://fetch.spec.whatwg.org/#request-render-blocking){#the-style-element:concept-request-render-blocking
x-internal="concept-request-render-blocking"} set to whether or not the
[`style`{#the-style-element:the-style-element-15}](#the-style-element)
element is currently
[render-blocking](dom.html#render-blocking){#the-style-element:render-blocking}.

1.  Let `element`{.variable} be the
    [`style`{#the-style-element:the-style-element-16}](#the-style-element)
    element associated with the style sheet in question.

2.  Let `success`{.variable} be true.

3.  If the attempts to obtain any of the style sheet\'s [critical
    subresources](infrastructure.html#critical-subresources){#the-style-element:critical-subresources-5}
    failed for any reason (e.g., DNS error, HTTP 404 response, a
    connection being prematurely closed, unsupported Content-Type), set
    `success`{.variable} to false.

    Note that content-specific errors, e.g., CSS parse errors or PNG
    decoding errors, do not affect `success`{.variable}.

4.  [Queue an element
    task](webappapis.html#queue-an-element-task){#the-style-element:queue-an-element-task}
    on the [networking task
    source](webappapis.html#networking-task-source){#the-style-element:networking-task-source}
    given `element`{.variable} and the following steps:

    1.  If `success`{.variable} is true, [fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#the-style-element:concept-event-fire
        x-internal="concept-event-fire"} named
        [`load`{#the-style-element:event-load}](indices.html#event-load)
        at `element`{.variable}.

    2.  Otherwise, [fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#the-style-element:concept-event-fire-2
        x-internal="concept-event-fire"} named
        [`error`{#the-style-element:event-error}](indices.html#event-error)
        at `element`{.variable}.

    3.  If `element`{.variable} [contributes a script-blocking style
        sheet](#contributes-a-script-blocking-style-sheet){#the-style-element:contributes-a-script-blocking-style-sheet-2}:

        1.  [Assert](https://infra.spec.whatwg.org/#assert){#the-style-element:assert
            x-internal="assert"}: `element`{.variable}\'s [node
            document](https://dom.spec.whatwg.org/#concept-node-document){#the-style-element:node-document-3
            x-internal="node-document"}\'s [script-blocking style sheet
            set](#script-blocking-style-sheet-set){#the-style-element:script-blocking-style-sheet-set-2}
            [contains](https://infra.spec.whatwg.org/#list-contain){#the-style-element:list-contains
            x-internal="list-contains"} `element`{.variable}.

        2.  [Remove](https://infra.spec.whatwg.org/#list-remove){#the-style-element:list-remove
            x-internal="list-remove"} `element`{.variable} from its
            [node
            document](https://dom.spec.whatwg.org/#concept-node-document){#the-style-element:node-document-4
            x-internal="node-document"}\'s [script-blocking style sheet
            set](#script-blocking-style-sheet-set){#the-style-element:script-blocking-style-sheet-set-3}.

    4.  [Unblock
        rendering](dom.html#unblock-rendering){#the-style-element:unblock-rendering}
        on `element`{.variable}.

The element must [delay the load
event](parsing.html#delay-the-load-event){#the-style-element:delay-the-load-event}
of the element\'s [node
document](https://dom.spec.whatwg.org/#concept-node-document){#the-style-element:node-document-5
x-internal="node-document"} until all the attempts to obtain the style
sheet\'s [critical
subresources](infrastructure.html#critical-subresources){#the-style-element:critical-subresources-6},
if any, are complete.

This specification does not specify a style system, but CSS is expected
to be supported by most web browsers. [\[CSS\]](references.html#refsCSS)

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLStyleElement/media](https://developer.mozilla.org/en-US/docs/Web/API/HTMLStyleElement/media "The HTMLStyleElement.media property specifies the intended destination medium for style information.")

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
:::::

The [`media`]{#dom-style-media .dfn dfn-for="HTMLStyleElement"
dfn-type="attribute"} and [`blocking`]{#dom-style-blocking .dfn
dfn-for="HTMLStyleElement" dfn-type="attribute"} IDL attributes must
each
[reflect](common-dom-interfaces.html#reflect){#the-style-element:reflect}
the respective content attributes of the same name.

The
[`LinkStyle`{#the-style-element:linkstyle-2}](https://drafts.csswg.org/cssom/#the-linkstyle-interface){x-internal="linkstyle"}
interface is also implemented by this element.
[\[CSSOM\]](references.html#refsCSSOM)

::: example
The following document has its stress emphasis styled as bright red text
rather than italics text, while leaving titles of works and Latin words
in their default italics. It shows how using appropriate elements
enables easier restyling of documents.

``` html
<!DOCTYPE html>
<html lang="en-US">
 <head>
  <title>My favorite book</title>
  <style>
   body { color: black; background: white; }
   em { font-style: normal; color: red; }
  </style>
 </head>
 <body>
  <p>My <em>favorite</em> book of all time has <em>got</em> to be
  <cite>A Cat's Life</cite>. It is a book by P. Rahmel that talks
  about the <i lang="la">Felis catus</i> in modern human society.</p>
 </body>
</html>
```
:::

#### [4.2.7]{.secno} Interactions of styling and scripting[](#interactions-of-styling-and-scripting){.self-link}

If the style sheet referenced no other resources (e.g., it was an
internal style sheet given by a
[`style`{#interactions-of-styling-and-scripting:the-style-element}](#the-style-element)
element with no `@import` rules), then the style rules must be
[immediately](infrastructure.html#immediately){#interactions-of-styling-and-scripting:immediately}
made available to script; otherwise, the style rules must only be made
available to script once the [event
loop](webappapis.html#event-loop){#interactions-of-styling-and-scripting:event-loop}
reaches its [update the
rendering](webappapis.html#update-the-rendering){#interactions-of-styling-and-scripting:update-the-rendering}
step.

An element `el`{.variable} in the context of a
[`Document`{#interactions-of-styling-and-scripting:document}](dom.html#document)
of an [HTML
parser](parsing.html#html-parser){#interactions-of-styling-and-scripting:html-parser}
or [XML
parser](xhtml.html#xml-parser){#interactions-of-styling-and-scripting:xml-parser}
[contributes a script-blocking style
sheet]{#contributes-a-script-blocking-style-sheet .dfn} if all of the
following are true:

- `el`{.variable} was created by that
  [`Document`{#interactions-of-styling-and-scripting:document-2}](dom.html#document)\'s
  parser.

- `el`{.variable} is either a
  [`style`{#interactions-of-styling-and-scripting:the-style-element-2}](#the-style-element)
  element or a
  [`link`{#interactions-of-styling-and-scripting:the-link-element}](#the-link-element)
  element that was an [external resource link that contributes to the
  styling processing
  model](links.html#link-type-stylesheet){#interactions-of-styling-and-scripting:link-type-stylesheet}
  when the `el`{.variable} was created by the parser.

- `el`{.variable}\'s `media` attribute\'s value [matches the
  environment](common-microsyntaxes.html#matches-the-environment){#interactions-of-styling-and-scripting:matches-the-environment}.

- `el`{.variable}\'s style sheet was enabled when the element was
  created by the parser.

- The last time the [event
  loop](webappapis.html#event-loop){#interactions-of-styling-and-scripting:event-loop-2}
  reached [step 1](webappapis.html#step1), `el`{.variable}\'s
  [root](https://dom.spec.whatwg.org/#concept-tree-root){#interactions-of-styling-and-scripting:root
  x-internal="root"} was that
  [`Document`{#interactions-of-styling-and-scripting:document-3}](dom.html#document).

- The user agent hasn\'t given up on loading that particular style sheet
  yet. A user agent may give up on loading a style sheet at any time.

  Giving up on a style sheet before the style sheet loads, if the style
  sheet eventually does still load, means that the script might end up
  operating with incorrect information. For example, if a style sheet
  sets the color of an element to green, but a script that inspects the
  resulting style is executed before the sheet is loaded, the script
  will find that the element is black (or whatever the default color
  is), and might thus make poor choices (e.g., deciding to use black as
  the color elsewhere on the page, instead of green). Implementers have
  to balance the likelihood of a script using incorrect information with
  the performance impact of doing nothing while waiting for a slow
  network request to finish.

It is expected that counterparts to the above rules also apply to
[`<?xml-stylesheet?>`{#interactions-of-styling-and-scripting:xml-stylesheet}](https://www.w3.org/TR/xml-stylesheet/#the-xml-stylesheet-processing-instruction){x-internal="xml-stylesheet"}
PIs. However, this has not yet been thoroughly investigated.

A
[`Document`{#interactions-of-styling-and-scripting:document-4}](dom.html#document)
has a [script-blocking style sheet set]{#script-blocking-style-sheet-set
.dfn}, which is an [ordered
set](https://infra.spec.whatwg.org/#ordered-set){#interactions-of-styling-and-scripting:set
x-internal="set"}, initially empty.

A
[`Document`{#interactions-of-styling-and-scripting:document-5}](dom.html#document)
`document`{.variable} [has a style sheet that is blocking
scripts]{#has-a-style-sheet-that-is-blocking-scripts .dfn} if the
following steps return true:

1.  If `document`{.variable}\'s [script-blocking style sheet
    set](#script-blocking-style-sheet-set){#interactions-of-styling-and-scripting:script-blocking-style-sheet-set}
    is not
    [empty](https://infra.spec.whatwg.org/#list-is-empty){#interactions-of-styling-and-scripting:list-is-empty
    x-internal="list-is-empty"}, then return true.

2.  If `document`{.variable}\'s [node
    navigable](document-sequences.html#node-navigable){#interactions-of-styling-and-scripting:node-navigable}
    is null, then return false.

3.  Let `containerDocument`{.variable} be `document`{.variable}\'s [node
    navigable](document-sequences.html#node-navigable){#interactions-of-styling-and-scripting:node-navigable-2}\'s
    [container
    document](document-sequences.html#nav-container){#interactions-of-styling-and-scripting:nav-container}.

4.  If `containerDocument`{.variable} is non-null and
    `containerDocument`{.variable}\'s [script-blocking style sheet
    set](#script-blocking-style-sheet-set){#interactions-of-styling-and-scripting:script-blocking-style-sheet-set-2}
    is not
    [empty](https://infra.spec.whatwg.org/#list-is-empty){#interactions-of-styling-and-scripting:list-is-empty-2
    x-internal="list-is-empty"}, then return true.

5.  Return false.

A
[`Document`{#interactions-of-styling-and-scripting:document-6}](dom.html#document)
[has no style sheet that is blocking
scripts]{#has-no-style-sheet-that-is-blocking-scripts .dfn} if it does
not [have a style sheet that is blocking
scripts](#has-a-style-sheet-that-is-blocking-scripts){#interactions-of-styling-and-scripting:has-a-style-sheet-that-is-blocking-scripts}.

[← 3 Semantics, structure, and APIs of HTML documents](dom.html) ---
[Table of Contents](./) --- [4.3 Sections →](sections.html)
