HTML Standard

::: {#head .head .with-buttons}
[![WHATWG](https://resources.whatwg.org/logo.svg){.darkmode-aware
width="100" crossorigin="" height="100"}](https://whatwg.org/){.logo}

# HTML {#html .allcaps}

Living Standard --- Last Updated [2 June 2025]{.pubdate}
:::

[← 4.3 Sections](sections.html) --- [Table of Contents](./) --- [4.5
Text-level semantics →](text-level-semantics.html)

1.  ::: {#toc-semantics}
    1.  [[4.4]{.secno} Grouping
        content](grouping-content.html#grouping-content)
        1.  [[4.4.1]{.secno} The `p`
            element](grouping-content.html#the-p-element)
        2.  [[4.4.2]{.secno} The `hr`
            element](grouping-content.html#the-hr-element)
        3.  [[4.4.3]{.secno} The `pre`
            element](grouping-content.html#the-pre-element)
        4.  [[4.4.4]{.secno} The `blockquote`
            element](grouping-content.html#the-blockquote-element)
        5.  [[4.4.5]{.secno} The `ol`
            element](grouping-content.html#the-ol-element)
        6.  [[4.4.6]{.secno} The `ul`
            element](grouping-content.html#the-ul-element)
        7.  [[4.4.7]{.secno} The `menu`
            element](grouping-content.html#the-menu-element)
        8.  [[4.4.8]{.secno} The `li`
            element](grouping-content.html#the-li-element)
        9.  [[4.4.9]{.secno} The `dl`
            element](grouping-content.html#the-dl-element)
        10. [[4.4.10]{.secno} The `dt`
            element](grouping-content.html#the-dt-element)
        11. [[4.4.11]{.secno} The `dd`
            element](grouping-content.html#the-dd-element)
        12. [[4.4.12]{.secno} The `figure`
            element](grouping-content.html#the-figure-element)
        13. [[4.4.13]{.secno} The `figcaption`
            element](grouping-content.html#the-figcaption-element)
        14. [[4.4.14]{.secno} The `main`
            element](grouping-content.html#the-main-element)
        15. [[4.4.15]{.secno} The `search`
            element](grouping-content.html#the-search-element)
        16. [[4.4.16]{.secno} The `div`
            element](grouping-content.html#the-div-element)
    :::

### [4.4]{.secno} Grouping content[](#grouping-content){.self-link}

#### [4.4.1]{.secno} The [`p`]{.dfn dfn-type="element"} element[](#the-p-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/p](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/p "The <p> HTML element represents a paragraph. Paragraphs are usually represented in visual media as blocks of text separated from adjacent blocks by blank lines and/or first-line indentation, but HTML paragraphs can be any structural grouping of related content, such as images or form fields.")

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
[HTMLParagraphElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLParagraphElement "The HTMLParagraphElement interface provides special properties (beyond those of the regular HTMLElement object interface it inherits) for manipulating <p> elements.")

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

[Categories](dom.html#concept-element-categories){#the-p-element:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-p-element:flow-content-2}.
:   [Palpable
    content](dom.html#palpable-content-2){#the-p-element:palpable-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-p-element:concept-element-contexts}:
:   Where [flow
    content](dom.html#flow-content-2){#the-p-element:flow-content-2-2}
    is expected.

[Content model](dom.html#concept-element-content-model){#the-p-element:concept-element-content-model}:
:   [Phrasing
    content](dom.html#phrasing-content-2){#the-p-element:phrasing-content-2}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-p-element:concept-element-tag-omission}:
:   A [`p`{#the-p-element:the-p-element}](#the-p-element) element\'s
    [end tag](syntax.html#syntax-end-tag){#the-p-element:syntax-end-tag}
    can be omitted if the
    [`p`{#the-p-element:the-p-element-2}](#the-p-element) element is
    immediately followed by an
    [`address`{#the-p-element:the-address-element}](sections.html#the-address-element),
    [`article`{#the-p-element:the-article-element}](sections.html#the-article-element),
    [`aside`{#the-p-element:the-aside-element}](sections.html#the-aside-element),
    [`blockquote`{#the-p-element:the-blockquote-element}](#the-blockquote-element),
    [`details`{#the-p-element:the-details-element}](interactive-elements.html#the-details-element),
    [`dialog`{#the-p-element:the-dialog-element}](interactive-elements.html#the-dialog-element),
    [`div`{#the-p-element:the-div-element}](#the-div-element),
    [`dl`{#the-p-element:the-dl-element}](#the-dl-element),
    [`fieldset`{#the-p-element:the-fieldset-element}](form-elements.html#the-fieldset-element),
    [`figcaption`{#the-p-element:the-figcaption-element}](#the-figcaption-element),
    [`figure`{#the-p-element:the-figure-element}](#the-figure-element),
    [`footer`{#the-p-element:the-footer-element}](sections.html#the-footer-element),
    [`form`{#the-p-element:the-form-element}](forms.html#the-form-element),
    [`h1`{#the-p-element:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements),
    [`h2`{#the-p-element:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-2}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements),
    [`h3`{#the-p-element:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-3}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements),
    [`h4`{#the-p-element:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-4}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements),
    [`h5`{#the-p-element:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-5}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements),
    [`h6`{#the-p-element:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-6}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements),
    [`header`{#the-p-element:the-header-element}](sections.html#the-header-element),
    [`hgroup`{#the-p-element:the-hgroup-element}](sections.html#the-hgroup-element),
    [`hr`{#the-p-element:the-hr-element}](#the-hr-element),
    [`main`{#the-p-element:the-main-element}](#the-main-element),
    [`menu`{#the-p-element:the-menu-element}](#the-menu-element),
    [`nav`{#the-p-element:the-nav-element}](sections.html#the-nav-element),
    [`ol`{#the-p-element:the-ol-element}](#the-ol-element),
    [`p`{#the-p-element:the-p-element-3}](#the-p-element),
    [`pre`{#the-p-element:the-pre-element}](#the-pre-element),
    [`search`{#the-p-element:the-search-element}](#the-search-element),
    [`section`{#the-p-element:the-section-element}](sections.html#the-section-element),
    [`table`{#the-p-element:the-table-element}](tables.html#the-table-element),
    or [`ul`{#the-p-element:the-ul-element}](#the-ul-element) element,
    or if there is no more content in the parent element and the parent
    element is an [HTML
    element](infrastructure.html#html-elements){#the-p-element:html-elements}
    that is not an
    [`a`{#the-p-element:the-a-element}](text-level-semantics.html#the-a-element),
    [`audio`{#the-p-element:the-audio-element}](media.html#the-audio-element),
    [`del`{#the-p-element:the-del-element}](edits.html#the-del-element),
    [`ins`{#the-p-element:the-ins-element}](edits.html#the-ins-element),
    [`map`{#the-p-element:the-map-element}](image-maps.html#the-map-element),
    [`noscript`{#the-p-element:the-noscript-element}](scripting.html#the-noscript-element),
    or
    [`video`{#the-p-element:the-video-element}](media.html#the-video-element)
    element, or an [autonomous custom
    element](custom-elements.html#autonomous-custom-element){#the-p-element:autonomous-custom-element}.

[Content attributes](dom.html#concept-element-attributes){#the-p-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-p-element:global-attributes}

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-p-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-p).
:   [For implementers](https://w3c.github.io/html-aam/#el-p).

[DOM interface](dom.html#concept-element-dom){#the-p-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLParagraphElement : HTMLElement {
      [HTMLConstructor] constructor();

      // also has obsolete members
    };
    ```

The [`p`{#the-p-element:the-p-element-4}](#the-p-element) element
[represents](dom.html#represents){#the-p-element:represents} a
[paragraph](dom.html#paragraph){#the-p-element:paragraph}.

While paragraphs are usually represented in visual media by blocks of
text that are physically separated from adjacent blocks through blank
lines, a style sheet or user agent would be equally justified in
presenting paragraph breaks in a different manner, for instance using
inline pilcrows (¶).

::: example
The following examples are conforming HTML fragments:

``` html
<p>The little kitten gently seated herself on a piece of
carpet. Later in her life, this would be referred to as the time the
cat sat on the mat.</p>
```

``` html
<fieldset>
 <legend>Personal information</legend>
 <p>
   <label>Name: <input name="n"></label>
   <label><input name="anon" type="checkbox"> Hide from other users</label>
 </p>
 <p><label>Address: <textarea name="a"></textarea></label></p>
</fieldset>
```

``` html
<p>There was once an example from Femley,<br>
Whose markup was of dubious quality.<br>
The validator complained,<br>
So the author was pained,<br>
To move the error from the markup to the rhyming.</p>
```
:::

The [`p`{#the-p-element:the-p-element-5}](#the-p-element) element should
not be used when a more specific element is more appropriate.

::: example
The following example is technically correct:

``` html
<section>
 <!-- ... -->
 <p>Last modified: 2001-04-23</p>
 <p>Author: fred@example.com</p>
</section>
```

However, it would be better marked-up as:

``` html
<section>
 <!-- ... -->
 <footer>Last modified: 2001-04-23</footer>
 <address>Author: fred@example.com</address>
</section>
```

Or:

``` html
<section>
 <!-- ... -->
 <footer>
  <p>Last modified: 2001-04-23</p>
  <address>Author: fred@example.com</address>
 </footer>
</section>
```
:::

:::::: note
List elements (in particular,
[`ol`{#the-p-element:the-ol-element-2}](#the-ol-element) and
[`ul`{#the-p-element:the-ul-element-2}](#the-ul-element) elements)
cannot be children of
[`p`{#the-p-element:the-p-element-6}](#the-p-element) elements. When a
sentence contains a bulleted list, therefore, one might wonder how it
should be marked up.

::: example
For instance, this fantastic sentence has bullets relating to

- wizards,
- faster-than-light travel, and
- telepathy,

and is further discussed below.
:::

The solution is to realize that a *[paragraph](dom.html#paragraph)*, in
HTML terms, is not a logical concept, but a structural one. In the
fantastic example above, there are actually *five*
[paragraphs](dom.html#paragraph){#the-p-element:paragraph-3} as defined
by this specification: one before the list, one for each bullet, and one
after the list.

::: example
The markup for the above example could therefore be:

``` html
<p>For instance, this fantastic sentence has bullets relating to</p>
<ul>
 <li>wizards,
 <li>faster-than-light travel, and
 <li>telepathy,
</ul>
<p>and is further discussed below.</p>
```
:::

Authors wishing to conveniently style such \"logical\" paragraphs
consisting of multiple \"structural\" paragraphs can use the
[`div`{#the-p-element:the-div-element-2}](#the-div-element) element
instead of the [`p`{#the-p-element:the-p-element-7}](#the-p-element)
element.

::: example
Thus for instance the above example could become the following:

``` html
<div>For instance, this fantastic sentence has bullets relating to
<ul>
 <li>wizards,
 <li>faster-than-light travel, and
 <li>telepathy,
</ul>
and is further discussed below.</div>
```

This example still has five structural paragraphs, but now the author
can style just the
[`div`{#the-p-element:the-div-element-3}](#the-div-element) instead of
having to consider each part of the example separately.
:::
::::::

#### [4.4.2]{.secno} The [`hr`]{.dfn dfn-type="element"} element[](#the-hr-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/hr](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/hr "The <hr> HTML element represents a thematic break between paragraph-level elements: for example, a change of scene in a story, or a shift of topic within a section.")

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

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLHRElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLHRElement "The HTMLHRElement interface provides special properties (beyond those of the HTMLElement interface it also has available to it by inheritance) for manipulating <hr> elements.")

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

[Categories](dom.html#concept-element-categories){#the-hr-element:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-hr-element:flow-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-hr-element:concept-element-contexts}:
:   Where [flow
    content](dom.html#flow-content-2){#the-hr-element:flow-content-2-2}
    is expected.
:   As a child of a
    [`select`{#the-hr-element:the-select-element}](form-elements.html#the-select-element)
    element.

[Content model](dom.html#concept-element-content-model){#the-hr-element:concept-element-content-model}:
:   [Nothing](dom.html#concept-content-nothing){#the-hr-element:concept-content-nothing}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-hr-element:concept-element-tag-omission}:
:   No [end
    tag](syntax.html#syntax-end-tag){#the-hr-element:syntax-end-tag}.

[Content attributes](dom.html#concept-element-attributes){#the-hr-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-hr-element:global-attributes}

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-hr-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-hr).
:   [For implementers](https://w3c.github.io/html-aam/#el-hr).

[DOM interface](dom.html#concept-element-dom){#the-hr-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLHRElement : HTMLElement {
      [HTMLConstructor] constructor();

      // also has obsolete members
    };
    ```

The [`hr`{#the-hr-element:the-hr-element}](#the-hr-element) element
[represents](dom.html#represents){#the-hr-element:represents} a
[paragraph](dom.html#paragraph){#the-hr-element:paragraph}-level
thematic break, e.g., a scene change in a story, or a transition to
another topic within a section of a reference book; alternatively, it
represents a separator between a set of options of a
[`select`{#the-hr-element:the-select-element-2}](form-elements.html#the-select-element)
element.

::: example
The following fictional extract from a project manual shows two sections
that use the [`hr`{#the-hr-element:the-hr-element-2}](#the-hr-element)
element to separate topics within the section.

``` html
<section>
 <h1>Communication</h1>
 <p>There are various methods of communication. This section
 covers a few of the important ones used by the project.</p>
 <hr>
 <p>Communication stones seem to come in pairs and have mysterious
 properties:</p>
 <ul>
  <li>They can transfer thoughts in two directions once activated
  if used alone.</li>
  <li>If used with another device, they can transfer one's
  consciousness to another body.</li>
  <li>If both stones are used with another device, the
  consciousnesses switch bodies.</li>
 </ul>
 <hr>
 <p>Radios use the electromagnetic spectrum in the meter range and
 longer.</p>
 <hr>
 <p>Signal flares use the electromagnetic spectrum in the
 nanometer range.</p>
</section>
<section>
 <h1>Food</h1>
 <p>All food at the project is rationed:</p>
 <dl>
  <dt>Potatoes</dt>
  <dd>Two per day</dd>
  <dt>Soup</dt>
  <dd>One bowl per day</dd>
 </dl>
 <hr>
 <p>Cooking is done by the chefs on a set rotation.</p>
</section>
```

There is no need for an
[`hr`{#the-hr-element:the-hr-element-3}](#the-hr-element) element
between the sections themselves, since the
[`section`{#the-hr-element:the-section-element}](sections.html#the-section-element)
elements and the
[`h1`{#the-hr-element:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements)
elements imply thematic changes themselves.
:::

::: example
The following extract from Pandora\'s Star by Peter F. Hamilton shows
two paragraphs that precede a scene change and the paragraph that
follows it. The scene change, represented in the printed book by a gap
containing a solitary centered star between the second and third
paragraphs, is here represented using the
[`hr`{#the-hr-element:the-hr-element-4}](#the-hr-element) element.

``` {lang="en-GB"}
<p>Dudley was ninety-two, in his second life, and fast approaching
time for another rejuvenation. Despite his body having the physical
age of a standard fifty-year-old, the prospect of a long degrading
campaign within academia was one he regarded with dread. For a
supposedly advanced civilization, the Intersolar Commonwealth could be
appallingly backward at times, not to mention cruel.</p>
<p><i>Maybe it won't be that bad</i>, he told himself. The lie was
comforting enough to get him through the rest of the night's
shift.</p>
<hr>
<p>The Carlton AllLander drove Dudley home just after dawn. Like the
astronomer, the vehicle was old and worn, but perfectly capable of
doing its job. It had a cheap diesel engine, common enough on a
semi-frontier world like Gralmond, although its drive array was a
thoroughly modern photoneural processor. With its high suspension and
deep-tread tyres it could plough along the dirt track to the
observatory in all weather and seasons, including the metre-deep snow
of Gralmond's winters.</p>
```
:::

The [`hr`{#the-hr-element:the-hr-element-5}](#the-hr-element) element
does not affect the document\'s
[outline](sections.html#outline){#the-hr-element:outline}.

#### [4.4.3]{.secno} The [`pre`]{.dfn dfn-type="element"} element[](#the-pre-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/pre](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/pre "The <pre> HTML element represents preformatted text which is to be presented exactly as written in the HTML file. The text is typically rendered using a non-proportional, or monospaced, font. Whitespace inside this element is displayed as written.")

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

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLPreElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLPreElement "The HTMLPreElement interface exposes specific properties and methods (beyond those of the HTMLElement interface it also has available to it by inheritance) for manipulating a block of preformatted text (<pre>).")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer6+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

[Categories](dom.html#concept-element-categories){#the-pre-element:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-pre-element:flow-content-2}.
:   [Palpable
    content](dom.html#palpable-content-2){#the-pre-element:palpable-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-pre-element:concept-element-contexts}:
:   Where [flow
    content](dom.html#flow-content-2){#the-pre-element:flow-content-2-2}
    is expected.

[Content model](dom.html#concept-element-content-model){#the-pre-element:concept-element-content-model}:
:   [Phrasing
    content](dom.html#phrasing-content-2){#the-pre-element:phrasing-content-2}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-pre-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-pre-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-pre-element:global-attributes}

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-pre-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-pre).
:   [For implementers](https://w3c.github.io/html-aam/#el-pre).

[DOM interface](dom.html#concept-element-dom){#the-pre-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLPreElement : HTMLElement {
      [HTMLConstructor] constructor();

      // also has obsolete members
    };
    ```

The [`pre`{#the-pre-element:the-pre-element}](#the-pre-element) element
[represents](dom.html#represents){#the-pre-element:represents} a block
of preformatted text, in which structure is represented by typographic
conventions rather than by elements.

In [the HTML syntax](syntax.html#syntax){#the-pre-element:syntax}, a
leading newline character immediately following the
[`pre`{#the-pre-element:the-pre-element-2}](#the-pre-element) element
start tag is stripped.

Some examples of cases where the
[`pre`{#the-pre-element:the-pre-element-3}](#the-pre-element) element
could be used:

- Including an email, with paragraphs indicated by blank lines, lists
  indicated by lines prefixed with a bullet, and so on.
- Including fragments of computer code, with structure indicated
  according to the conventions of that language.
- Displaying ASCII art.

Authors are encouraged to consider how preformatted text will be
experienced when the formatting is lost, as will be the case for users
of speech synthesizers, braille displays, and the like. For cases like
ASCII art, it is likely that an alternative presentation, such as a
textual description, would be more universally accessible to the readers
of the document.

To represent a block of computer code, the
[`pre`{#the-pre-element:the-pre-element-4}](#the-pre-element) element
can be used with a
[`code`{#the-pre-element:the-code-element}](text-level-semantics.html#the-code-element)
element; to represent a block of computer output the
[`pre`{#the-pre-element:the-pre-element-5}](#the-pre-element) element
can be used with a
[`samp`{#the-pre-element:the-samp-element}](text-level-semantics.html#the-samp-element)
element. Similarly, the
[`kbd`{#the-pre-element:the-kbd-element}](text-level-semantics.html#the-kbd-element)
element can be used within a
[`pre`{#the-pre-element:the-pre-element-6}](#the-pre-element) element to
indicate text that the user is to enter.

This element [has rendering requirements involving the bidirectional
algorithm](dom.html#bidireq).

::: example
In the following snippet, a sample of computer code is presented.

``` html
<p>This is the <code>Panel</code> constructor:</p>
<pre><code>function Panel(element, canClose, closeHandler) {
  this.element = element;
  this.canClose = canClose;
  this.closeHandler = function () { if (closeHandler) closeHandler() };
}</code></pre>
```
:::

::: example
In the following snippet,
[`samp`{#the-pre-element:the-samp-element-2}](text-level-semantics.html#the-samp-element)
and
[`kbd`{#the-pre-element:the-kbd-element-2}](text-level-semantics.html#the-kbd-element)
elements are mixed in the contents of a
[`pre`{#the-pre-element:the-pre-element-7}](#the-pre-element) element to
show a session of Zork I.

``` html
<pre><samp>You are in an open field west of a big white house with a boarded
front door.
There is a small mailbox here.

></samp> <kbd>open mailbox</kbd>

<samp>Opening the mailbox reveals:
A leaflet.

></samp></pre>
```
:::

::: example
The following shows a contemporary poem that uses the
[`pre`{#the-pre-element:the-pre-element-8}](#the-pre-element) element to
preserve its unusual formatting, which forms an intrinsic part of the
poem itself.

``` html
<pre>                maxling

it is with a          heart
               heavy

that i admit loss of a feline
        so           loved

a friend lost to the
        unknown
                                (night)

~cdr 11dec07</pre>
```
:::

#### [4.4.4]{.secno} The [`blockquote`]{.dfn dfn-type="element"} element[](#the-blockquote-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/blockquote](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/blockquote "The <blockquote> HTML element indicates that the enclosed text is an extended quotation. Usually, this is rendered visually by indentation (see Notes for how to change it). A URL for the source of the quotation may be given using the cite attribute, while a text representation of the source can be given using the <cite> element.")

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

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLQuoteElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLQuoteElement "The HTMLQuoteElement interface provides special properties and methods (beyond the regular HTMLElement interface it also has available to it by inheritance) for manipulating quoting elements, like <blockquote> and <q>, but not the <cite> element.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer6+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

[Categories](dom.html#concept-element-categories){#the-blockquote-element:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-blockquote-element:flow-content-2}.
:   [Palpable
    content](dom.html#palpable-content-2){#the-blockquote-element:palpable-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-blockquote-element:concept-element-contexts}:
:   Where [flow
    content](dom.html#flow-content-2){#the-blockquote-element:flow-content-2-2}
    is expected.

[Content model](dom.html#concept-element-content-model){#the-blockquote-element:concept-element-content-model}:
:   [Flow
    content](dom.html#flow-content-2){#the-blockquote-element:flow-content-2-3}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-blockquote-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-blockquote-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-blockquote-element:global-attributes}
:   [`cite`{#the-blockquote-element:attr-blockquote-cite}](#attr-blockquote-cite)
    --- Link to the source of the quotation or more information about
    the edit

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-blockquote-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-blockquote).
:   [For implementers](https://w3c.github.io/html-aam/#el-blockquote).

[DOM interface](dom.html#concept-element-dom){#the-blockquote-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLQuoteElement : HTMLElement {
      [HTMLConstructor] constructor();

      [CEReactions] attribute USVString cite;
    };
    ```

    The
    [`HTMLQuoteElement`{#the-blockquote-element:htmlquoteelement}](#htmlquoteelement)
    interface is also used by the
    [`q`{#the-blockquote-element:the-q-element}](text-level-semantics.html#the-q-element)
    element.

The
[`blockquote`{#the-blockquote-element:the-blockquote-element}](#the-blockquote-element)
element
[represents](dom.html#represents){#the-blockquote-element:represents} a
section that is quoted from another source.

Content inside a
[`blockquote`{#the-blockquote-element:the-blockquote-element-2}](#the-blockquote-element)
must be quoted from another source, whose address, if it has one, may be
cited in the [`cite`]{#attr-blockquote-cite .dfn dfn-for="blockquote"
dfn-type="element-attr"} attribute.

If the
[`cite`{#the-blockquote-element:attr-blockquote-cite-2}](#attr-blockquote-cite)
attribute is present, it must be a [valid URL potentially surrounded by
spaces](urls-and-fetching.html#valid-url-potentially-surrounded-by-spaces){#the-blockquote-element:valid-url-potentially-surrounded-by-spaces}.
To obtain the corresponding citation link, the value of the attribute
must be
[parsed](urls-and-fetching.html#encoding-parsing-a-url){#the-blockquote-element:encoding-parsing-a-url}
relative to the element\'s [node
document](https://dom.spec.whatwg.org/#concept-node-document){#the-blockquote-element:node-document
x-internal="node-document"}. User agents may allow users to follow such
citation links, but they are primarily intended for private use (e.g.,
by server-side scripts collecting statistics about a site\'s use of
quotations), not for readers.

The content of a
[`blockquote`{#the-blockquote-element:the-blockquote-element-3}](#the-blockquote-element)
may be abbreviated or may have context added in the conventional manner
for the text\'s language.

::: example
For example, in English this is traditionally done using square
brackets. Consider a page with the sentence \"Jane ate the cracker. She
then said she liked apples and fish.\"; it could be quoted as follows:

``` html
<blockquote>
 <p>[Jane] then said she liked [...] fish.</p>
</blockquote>
```
:::

Attribution for the quotation, if any, must be placed outside the
[`blockquote`{#the-blockquote-element:the-blockquote-element-4}](#the-blockquote-element)
element.

::: example
For example, here the attribution is given in a paragraph after the
quote:

``` html
<blockquote>
 <p>I contend that we are both atheists. I just believe in one fewer
 god than you do. When you understand why you dismiss all the other
 possible gods, you will understand why I dismiss yours.</p>
</blockquote>
<p>— Stephen Roberts</p>
```

The other examples below show other ways of showing attribution.
:::

The [`cite`]{#dom-quote-cite .dfn dfn-for="HTMLQuoteElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-blockquote-element:reflect}
the element\'s `cite` content attribute.

::: example
Here a
[`blockquote`{#the-blockquote-element:the-blockquote-element-5}](#the-blockquote-element)
element is used in conjunction with a
[`figure`{#the-blockquote-element:the-figure-element}](#the-figure-element)
element and its
[`figcaption`{#the-blockquote-element:the-figcaption-element}](#the-figcaption-element)
to clearly relate a quote to its attribution (which is not part of the
quote and therefore doesn\'t belong inside the
[`blockquote`{#the-blockquote-element:the-blockquote-element-6}](#the-blockquote-element)
itself):

``` html
<figure>
 <blockquote>
  <p>The truth may be puzzling. It may take some work to grapple with.
  It may be counterintuitive. It may contradict deeply held
  prejudices. It may not be consonant with what we desperately want to
  be true. But our preferences do not determine what's true. We have a
  method, and that method helps us to reach not absolute truth, only
  asymptotic approaches to the truth — never there, just closer
  and closer, always finding vast new oceans of undiscovered
  possibilities. Cleverly designed experiments are the key.</p>
 </blockquote>
 <figcaption>Carl Sagan, in "<cite>Wonder and Skepticism</cite>", from
 the <cite>Skeptical Inquirer</cite> Volume 19, Issue 1 (January-February
 1995)</figcaption>
</figure>
```
:::

::: example
This next example shows the use of
[`cite`{#the-blockquote-element:the-cite-element}](text-level-semantics.html#the-cite-element)
alongside
[`blockquote`{#the-blockquote-element:the-blockquote-element-7}](#the-blockquote-element):

``` html
<p>His next piece was the aptly named <cite>Sonnet 130</cite>:</p>
<blockquote cite="https://quotes.example.org/s/sonnet130.html">
  <p>My mistress' eyes are nothing like the sun,<br>
  Coral is far more red, than her lips red,<br>
  ...
```
:::

::: example
This example shows how a forum post could use
[`blockquote`{#the-blockquote-element:the-blockquote-element-8}](#the-blockquote-element)
to show what post a user is replying to. The
[`article`{#the-blockquote-element:the-article-element}](sections.html#the-article-element)
element is used for each post, to mark up the threading.

``` html
<article>
 <h1><a href="https://bacon.example.com/?blog=109431">Bacon on a crowbar</a></h1>
 <article>
  <header><strong>t3yw</strong> 12 points 1 hour ago</header>
  <p>I bet a narwhal would love that.</p>
  <footer><a href="?pid=29578">permalink</a></footer>
  <article>
   <header><strong>greg</strong> 8 points 1 hour ago</header>
   <blockquote><p>I bet a narwhal would love that.</p></blockquote>
   <p>Dude narwhals don't eat bacon.</p>
   <footer><a href="?pid=29579">permalink</a></footer>
   <article>
    <header><strong>t3yw</strong> 15 points 1 hour ago</header>
    <blockquote>
     <blockquote><p>I bet a narwhal would love that.</p></blockquote>
     <p>Dude narwhals don't eat bacon.</p>
    </blockquote>
    <p>Next thing you'll be saying they don't get capes and wizard
    hats either!</p>
    <footer><a href="?pid=29580">permalink</a></footer>
    <article>
     <article>
      <header><strong>boing</strong> -5 points 1 hour ago</header>
      <p>narwhals are worse than ceiling cat</p>
      <footer><a href="?pid=29581">permalink</a></footer>
     </article>
    </article>
   </article>
  </article>
  <article>
   <header><strong>fred</strong> 1 points 23 minutes ago</header>
   <blockquote><p>I bet a narwhal would love that.</p></blockquote>
   <p>I bet they'd love to peel a banana too.</p>
   <footer><a href="?pid=29582">permalink</a></footer>
  </article>
 </article>
</article>
```
:::

::: example
This example shows the use of a
[`blockquote`{#the-blockquote-element:the-blockquote-element-9}](#the-blockquote-element)
for short snippets, demonstrating that one does not have to use
[`p`{#the-blockquote-element:the-p-element}](#the-p-element) elements
inside
[`blockquote`{#the-blockquote-element:the-blockquote-element-10}](#the-blockquote-element)
elements:

``` html
<p>He began his list of "lessons" with the following:</p>
<blockquote>One should never assume that his side of
the issue will be recognized, let alone that it will
be conceded to have merits.</blockquote>
<p>He continued with a number of similar points, ending with:</p>
<blockquote>Finally, one should be prepared for the threat
of breakdown in negotiations at any given moment and not
be cowed by the possibility.</blockquote>
<p>We shall now discuss these points...
```
:::

[Examples of how to represent a
conversation](semantics-other.html#conversations) are shown in a later
section; it is not appropriate to use the
[`cite`{#the-blockquote-element:the-cite-element-2}](text-level-semantics.html#the-cite-element)
and
[`blockquote`{#the-blockquote-element:the-blockquote-element-11}](#the-blockquote-element)
elements for this purpose.

#### [4.4.5]{.secno} The [`ol`]{.dfn dfn-type="element"} element[](#the-ol-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/ol](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ol "The <ol> HTML element represents an ordered list of items — typically rendered as a numbered list.")

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

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLOListElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLOListElement "The HTMLOListElement interface provides special properties (beyond those defined on the regular HTMLElement interface it also has available to it by inheritance) for manipulating ordered list elements.")

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

[Categories](dom.html#concept-element-categories){#the-ol-element:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-ol-element:flow-content-2}.
:   If the element\'s children include at least one
    [`li`{#the-ol-element:the-li-element}](#the-li-element) element:
    [Palpable
    content](dom.html#palpable-content-2){#the-ol-element:palpable-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-ol-element:concept-element-contexts}:
:   Where [flow
    content](dom.html#flow-content-2){#the-ol-element:flow-content-2-2}
    is expected.

[Content model](dom.html#concept-element-content-model){#the-ol-element:concept-element-content-model}:
:   Zero or more
    [`li`{#the-ol-element:the-li-element-2}](#the-li-element) and
    [script-supporting](dom.html#script-supporting-elements-2){#the-ol-element:script-supporting-elements-2}
    elements.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-ol-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-ol-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-ol-element:global-attributes}
:   [`reversed`{#the-ol-element:attr-ol-reversed}](#attr-ol-reversed)
    --- Number the list backwards
:   [`start`{#the-ol-element:attr-ol-start}](#attr-ol-start) ---
    [Starting
    value](#concept-ol-start){#the-ol-element:concept-ol-start} of the
    list
:   [`type`{#the-ol-element:attr-ol-type}](#attr-ol-type) --- Kind of
    list marker

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-ol-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-ol).
:   [For implementers](https://w3c.github.io/html-aam/#el-ol).

[DOM interface](dom.html#concept-element-dom){#the-ol-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLOListElement : HTMLElement {
      [HTMLConstructor] constructor();

      [CEReactions] attribute boolean reversed;
      [CEReactions] attribute long start;
      [CEReactions] attribute DOMString type;

      // also has obsolete members
    };
    ```

The [`ol`{#the-ol-element:the-ol-element}](#the-ol-element) element
[represents](dom.html#represents){#the-ol-element:represents} a list of
items, where the items have been intentionally ordered, such that
changing the order would change the meaning of the document.

The items of the list are the
[`li`{#the-ol-element:the-li-element-3}](#the-li-element) element child
nodes of the [`ol`{#the-ol-element:the-ol-element-2}](#the-ol-element)
element, in [tree
order](https://dom.spec.whatwg.org/#concept-tree-order){#the-ol-element:tree-order
x-internal="tree-order"}.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Element/ol#attr-reversed](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ol#attr-reversed "The <ol> HTML element represents an ordered list of items — typically rendered as a numbered list.")

Support in all current engines.

::: support
[Firefox18+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome18+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)≤79+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

The [`reversed`]{#attr-ol-reversed .dfn dfn-for="ol"
dfn-type="element-attr"} attribute is a [boolean
attribute](common-microsyntaxes.html#boolean-attribute){#the-ol-element:boolean-attribute}.
If present, it indicates that the list is a descending list (\..., 3, 2,
1). If the attribute is omitted, the list is an ascending list (1, 2, 3,
\...).

The [`start`]{#attr-ol-start .dfn dfn-for="ol" dfn-type="element-attr"}
attribute, if present, must be a [valid
integer](common-microsyntaxes.html#valid-integer){#the-ol-element:valid-integer}.
It is used to determine the [starting
value](#concept-ol-start){#the-ol-element:concept-ol-start-2} of the
list.

An [`ol`{#the-ol-element:the-ol-element-3}](#the-ol-element) element has
a [starting value]{#concept-ol-start .dfn}, which is an integer
determined as follows:

1.  If the [`ol`{#the-ol-element:the-ol-element-4}](#the-ol-element)
    element has a
    [`start`{#the-ol-element:attr-ol-start-2}](#attr-ol-start)
    attribute, then:

    1.  Let `parsed`{.variable} be the result of [parsing the value of
        the attribute as an
        integer](common-microsyntaxes.html#rules-for-parsing-integers){#the-ol-element:rules-for-parsing-integers}.

    2.  If `parsed`{.variable} is not an error, then return
        `parsed`{.variable}.

2.  If the [`ol`{#the-ol-element:the-ol-element-5}](#the-ol-element)
    element has a
    [`reversed`{#the-ol-element:attr-ol-reversed-2}](#attr-ol-reversed)
    attribute, then return the number of [owned `li`
    elements](#list-owner){#the-ol-element:list-owner}.

3.  Return 1.

The [`type`]{#attr-ol-type .dfn dfn-for="ol" dfn-type="element-attr"}
attribute can be used to specify the kind of marker to use in the list,
in the cases where that matters (e.g. because items are to be
[referenced](dom.html#referenced){#the-ol-element:referenced} by their
number/letter). The attribute, if specified, must have a value that is
[identical
to](https://infra.spec.whatwg.org/#string-is){#the-ol-element:identical-to
x-internal="identical-to"} one of the characters given in the first cell
of one of the rows of the following table. The
[`type`{#the-ol-element:attr-ol-type-2}](#attr-ol-type) attribute
represents the state given in the cell in the second column of the row
whose first cell matches the attribute\'s value; if none of the cells
match, or if the attribute is omitted, then the attribute represents the
[decimal](#attr-ol-type-state-decimal){#the-ol-element:attr-ol-type-state-decimal}
state.

Keyword

State

Description

Examples for values 1-3 and 3999-4001

[`1`]{#attr-ol-type-keyword-decimal .dfn dfn-for="ol/type"
dfn-type="attr-value"} (U+0031)

[decimal]{#attr-ol-type-state-decimal .dfn}

Decimal numbers

`1.`{.sample}

`2.`{.sample}

`3.`{.sample}

\...

`3999.`{.sample}

`4000.`{.sample}

`4001.`{.sample}

\...

[`a`]{#attr-ol-type-keyword-lower-alpha .dfn dfn-for="ol/type"
dfn-type="attr-value"} (U+0061)

[lower-alpha]{#attr-ol-type-state-lower-alpha .dfn}

Lowercase latin alphabet

`a.`{.sample}

`b.`{.sample}

`c.`{.sample}

\...

`ewu.`{.sample}

`ewv.`{.sample}

`eww.`{.sample}

\...

[`A`]{#attr-ol-type-keyword-upper-alpha .dfn dfn-for="ol/type"
dfn-type="attr-value"} (U+0041)

[upper-alpha]{#attr-ol-type-state-upper-alpha .dfn}

Uppercase latin alphabet

`A.`{.sample}

`B.`{.sample}

`C.`{.sample}

\...

`EWU.`{.sample}

`EWV.`{.sample}

`EWW.`{.sample}

\...

[`i`]{#attr-ol-type-keyword-lower-roman .dfn dfn-for="ol/type"
dfn-type="attr-value"} (U+0069)

[lower-roman]{#attr-ol-type-state-lower-roman .dfn}

Lowercase roman numerals

`i.`{.sample}

`ii.`{.sample}

`iii.`{.sample}

\...

`mmmcmxcix.`{.sample}

`i̅v̅.`{.sample}

`i̅v̅i.`{.sample}

\...

[`I`]{#attr-ol-type-keyword-upper-roman .dfn dfn-for="ol/type"
dfn-type="attr-value"} (U+0049)

[upper-roman]{#attr-ol-type-state-upper-roman .dfn}

Uppercase roman numerals

`I.`{.sample}

`II.`{.sample}

`III.`{.sample}

\...

`MMMCMXCIX.`{.sample}

`I̅V̅.`{.sample}

`I̅V̅I.`{.sample}

\...

User agents should render the items of the list in a manner consistent
with the state of the
[`type`{#the-ol-element:attr-ol-type-3}](#attr-ol-type) attribute of the
[`ol`{#the-ol-element:the-ol-element-6}](#the-ol-element) element.
Numbers less than or equal to zero should always use the decimal system
regardless of the
[`type`{#the-ol-element:attr-ol-type-4}](#attr-ol-type) attribute.

For CSS user agents, a mapping for this attribute to the
[\'list-style-type\'](https://drafts.csswg.org/css-lists/#propdef-list-style-type){#the-ol-element:'list-style-type'
x-internal="'list-style-type'"} CSS property is given [in the Rendering
section](rendering.html#decohints) (the mapping is straightforward: the
states above have the same names as their corresponding CSS values).

It is possible to redefine the default CSS list styles used to implement
this attribute in CSS user agents; doing so will affect how list items
are rendered.

The [`reversed`]{#dom-ol-reversed .dfn dfn-for="HTMLOListElement"
dfn-type="attribute"} and [`type`]{#dom-ol-type .dfn
dfn-for="HTMLOListElement" dfn-type="attribute"} IDL attributes must
[reflect](common-dom-interfaces.html#reflect){#the-ol-element:reflect}
the respective content attributes of the same name.

The [`start`]{#dom-ol-start .dfn dfn-for="HTMLOListElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-ol-element:reflect-2}
the content attribute of the same name, with a [default
value](common-dom-interfaces.html#default-value){#the-ol-element:default-value}
of 1.

This means that the
[`start`{#the-ol-element:dom-ol-start-2}](#dom-ol-start) IDL attribute
does not necessarily match the list\'s [starting
value](#concept-ol-start){#the-ol-element:concept-ol-start-3}, in cases
where the [`start`{#the-ol-element:attr-ol-start-3}](#attr-ol-start)
content attribute is omitted and the
[`reversed`{#the-ol-element:attr-ol-reversed-3}](#attr-ol-reversed)
content attribute is specified.

::: example
The following markup shows a list where the order matters, and where the
[`ol`{#the-ol-element:the-ol-element-7}](#the-ol-element) element is
therefore appropriate. Compare this list to the equivalent list in the
[`ul`{#the-ol-element:the-ul-element}](#the-ul-element) section to see
an example of the same items using the
[`ul`{#the-ol-element:the-ul-element-2}](#the-ul-element) element.

``` html
<p>I have lived in the following countries (given in the order of when
I first lived there):</p>
<ol>
 <li>Switzerland
 <li>United Kingdom
 <li>United States
 <li>Norway
</ol>
```

Note how changing the order of the list changes the meaning of the
document. In the following example, changing the relative order of the
first two items has changed the birthplace of the author:

``` html
<p>I have lived in the following countries (given in the order of when
I first lived there):</p>
<ol>
 <li>United Kingdom
 <li>Switzerland
 <li>United States
 <li>Norway
</ol>
```
:::

#### [4.4.6]{.secno} The [`ul`]{.dfn dfn-type="element"} element[](#the-ul-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/ul](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ul "The <ul> HTML element represents an unordered list of items, typically rendered as a bulleted list.")

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

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLUListElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLUListElement "The HTMLUListElement interface provides special properties (beyond those defined on the regular HTMLElement interface it also has available to it by inheritance) for manipulating unordered list (<ul>) elements.")

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

[Categories](dom.html#concept-element-categories){#the-ul-element:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-ul-element:flow-content-2}.
:   If the element\'s children include at least one
    [`li`{#the-ul-element:the-li-element}](#the-li-element) element:
    [Palpable
    content](dom.html#palpable-content-2){#the-ul-element:palpable-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-ul-element:concept-element-contexts}:
:   Where [flow
    content](dom.html#flow-content-2){#the-ul-element:flow-content-2-2}
    is expected.

[Content model](dom.html#concept-element-content-model){#the-ul-element:concept-element-content-model}:
:   Zero or more
    [`li`{#the-ul-element:the-li-element-2}](#the-li-element) and
    [script-supporting](dom.html#script-supporting-elements-2){#the-ul-element:script-supporting-elements-2}
    elements.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-ul-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-ul-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-ul-element:global-attributes}

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-ul-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-ul).
:   [For implementers](https://w3c.github.io/html-aam/#el-ul).

[DOM interface](dom.html#concept-element-dom){#the-ul-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLUListElement : HTMLElement {
      [HTMLConstructor] constructor();

      // also has obsolete members
    };
    ```

The [`ul`{#the-ul-element:the-ul-element}](#the-ul-element) element
[represents](dom.html#represents){#the-ul-element:represents} a list of
items, where the order of the items is not important --- that is, where
changing the order would not materially change the meaning of the
document.

The items of the list are the
[`li`{#the-ul-element:the-li-element-3}](#the-li-element) element child
nodes of the [`ul`{#the-ul-element:the-ul-element-2}](#the-ul-element)
element.

::: example
The following markup shows a list where the order does not matter, and
where the [`ul`{#the-ul-element:the-ul-element-3}](#the-ul-element)
element is therefore appropriate. Compare this list to the equivalent
list in the [`ol`{#the-ul-element:the-ol-element}](#the-ol-element)
section to see an example of the same items using the
[`ol`{#the-ul-element:the-ol-element-2}](#the-ol-element) element.

``` html
<p>I have lived in the following countries:</p>
<ul>
 <li>Norway
 <li>Switzerland
 <li>United Kingdom
 <li>United States
</ul>
```

Note that changing the order of the list does not change the meaning of
the document. The items in the snippet above are given in alphabetical
order, but in the snippet below they are given in order of the size of
their current account balance in 2007, without changing the meaning of
the document whatsoever:

``` html
<p>I have lived in the following countries:</p>
<ul>
 <li>Switzerland
 <li>Norway
 <li>United Kingdom
 <li>United States
</ul>
```
:::

#### [4.4.7]{.secno} The [`menu`]{#menus .dfn dfn-type="element"} element[](#the-menu-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/menu](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/menu "The <menu> HTML element is described in the HTML specification as a semantic alternative to <ul>, but treated by browsers (and exposed through the accessibility tree) as no different than <ul>. It represents an unordered list of items (which are represented by <li> elements).")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer6+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMenuElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMenuElement "The HTMLMenuElement interface provides special properties (beyond those defined on the regular HTMLElement interface it also has available to it by inheritance) for manipulating <menu> elements.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer6+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

[Categories](dom.html#concept-element-categories){#the-menu-element:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-menu-element:flow-content-2}.
:   If the element\'s children include at least one
    [`li`{#the-menu-element:the-li-element}](#the-li-element) element:
    [Palpable
    content](dom.html#palpable-content-2){#the-menu-element:palpable-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-menu-element:concept-element-contexts}:
:   Where [flow
    content](dom.html#flow-content-2){#the-menu-element:flow-content-2-2}
    is expected.

[Content model](dom.html#concept-element-content-model){#the-menu-element:concept-element-content-model}:
:   Zero or more
    [`li`{#the-menu-element:the-li-element-2}](#the-li-element) and
    [script-supporting](dom.html#script-supporting-elements-2){#the-menu-element:script-supporting-elements-2}
    elements.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-menu-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-menu-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-menu-element:global-attributes}

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-menu-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-menu).
:   [For implementers](https://w3c.github.io/html-aam/#el-menu).

[DOM interface](dom.html#concept-element-dom){#the-menu-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLMenuElement : HTMLElement {
      [HTMLConstructor] constructor();

      // also has obsolete members
    };
    ```

The [`menu`{#the-menu-element:the-menu-element}](#the-menu-element)
element [represents](dom.html#represents){#the-menu-element:represents}
a toolbar consisting of its contents, in the form of an unordered list
of items (represented by
[`li`{#the-menu-element:the-li-element-3}](#the-li-element) elements),
each of which represents a command that the user can perform or
activate.

The [`menu`{#the-menu-element:the-menu-element-2}](#the-menu-element)
element is simply a semantic alternative to
[`ul`{#the-menu-element:the-ul-element}](#the-ul-element) to express an
unordered list of commands (a \"toolbar\").

::: example
In this example, a text-editing application uses a
[`menu`{#the-menu-element:the-menu-element-3}](#the-menu-element)
element to provide a series of editing commands:

``` html
<menu>
 <li><button onclick="copy()"><img src="copy.svg" alt="Copy"></button></li>
 <li><button onclick="cut()"><img src="cut.svg" alt="Cut"></button></li>
 <li><button onclick="paste()"><img src="paste.svg" alt="Paste"></button></li>
</menu>
```

Note that the styling to make this look like a conventional toolbar menu
is up to the application.
:::

#### [4.4.8]{.secno} The [`li`]{.dfn dfn-type="element"} element[](#the-li-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/li](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/li "The <li> HTML element is used to represent an item in a list. It must be contained in a parent element: an ordered list (<ol>), an unordered list (<ul>), or a menu (<menu>). In menus and unordered lists, list items are usually displayed using bullet points. In ordered lists, they are usually displayed with an ascending counter on the left, such as a number or letter.")

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

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLLIElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLLIElement "The HTMLLIElement interface exposes specific properties and methods (beyond those defined by regular HTMLElement interface it also has available to it by inheritance) for manipulating list elements.")

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

[Categories](dom.html#concept-element-categories){#the-li-element:concept-element-categories}:
:   None.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-li-element:concept-element-contexts}:
:   Inside [`ol`{#the-li-element:the-ol-element}](#the-ol-element)
    elements.
:   Inside [`ul`{#the-li-element:the-ul-element}](#the-ul-element)
    elements.
:   Inside [`menu`{#the-li-element:the-menu-element}](#the-menu-element)
    elements.

[Content model](dom.html#concept-element-content-model){#the-li-element:concept-element-content-model}:
:   [Flow
    content](dom.html#flow-content-2){#the-li-element:flow-content-2}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-li-element:concept-element-tag-omission}:
:   An [`li`{#the-li-element:the-li-element}](#the-li-element)
    element\'s [end
    tag](syntax.html#syntax-end-tag){#the-li-element:syntax-end-tag} can
    be omitted if the
    [`li`{#the-li-element:the-li-element-2}](#the-li-element) element is
    immediately followed by another
    [`li`{#the-li-element:the-li-element-3}](#the-li-element) element or
    if there is no more content in the parent element.

[Content attributes](dom.html#concept-element-attributes){#the-li-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-li-element:global-attributes}
:   If the element is not a child of an
    [`ul`{#the-li-element:the-ul-element-2}](#the-ul-element) or
    [`menu`{#the-li-element:the-menu-element-2}](#the-menu-element)
    element: [`value`{#the-li-element:attr-li-value}](#attr-li-value)
    --- [Ordinal value](#ordinal-value){#the-li-element:ordinal-value}
    of the list item

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-li-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-li).
:   [For implementers](https://w3c.github.io/html-aam/#el-li).

[DOM interface](dom.html#concept-element-dom){#the-li-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLLIElement : HTMLElement {
      [HTMLConstructor] constructor();

      [CEReactions] attribute long value;

      // also has obsolete members
    };
    ```

The [`li`{#the-li-element:the-li-element-4}](#the-li-element) element
[represents](dom.html#represents){#the-li-element:represents} a list
item. If its parent element is an
[`ol`{#the-li-element:the-ol-element-2}](#the-ol-element),
[`ul`{#the-li-element:the-ul-element-3}](#the-ul-element), or
[`menu`{#the-li-element:the-menu-element-3}](#the-menu-element) element,
then the element is an item of the parent element\'s list, as defined
for those elements. Otherwise, the list item has no defined list-related
relationship to any other
[`li`{#the-li-element:the-li-element-5}](#the-li-element) element.

The [`value`]{#attr-li-value .dfn dfn-for="li" dfn-type="element-attr"}
attribute, if present, must be a [valid
integer](common-microsyntaxes.html#valid-integer){#the-li-element:valid-integer}.
It is used to determine the [ordinal
value](#ordinal-value){#the-li-element:ordinal-value-2} of the list
item, when the
[`li`{#the-li-element:the-li-element-6}](#the-li-element)\'s [list
owner](#list-owner){#the-li-element:list-owner} is an
[`ol`{#the-li-element:the-ol-element-3}](#the-ol-element) element.

------------------------------------------------------------------------

Any element whose [computed
value](https://drafts.csswg.org/css-cascade/#computed-value){#the-li-element:computed-value
x-internal="computed-value"} of
[\'display\'](https://drafts.csswg.org/css2/#display-prop){#the-li-element:'display'
x-internal="'display'"} is \'list-item\' has a [list owner]{#list-owner
.dfn}, which is determined as follows:

1.  If the element is not [being
    rendered](rendering.html#being-rendered){#the-li-element:being-rendered},
    return null; the element has no [list
    owner](#list-owner){#the-li-element:list-owner-2}.

2.  Let `ancestor`{.variable} be the element\'s parent.

3.  If the element has an
    [`ol`{#the-li-element:the-ol-element-4}](#the-ol-element),
    [`ul`{#the-li-element:the-ul-element-4}](#the-ul-element), or
    [`menu`{#the-li-element:the-menu-element-4}](#the-menu-element)
    ancestor, set `ancestor`{.variable} to the closest such ancestor
    element.

4.  Return the closest inclusive ancestor of `ancestor`{.variable} that
    produces a [CSS
    box](https://drafts.csswg.org/css-display/#css-box){#the-li-element:css-box
    x-internal="css-box"}.

    Such an element will always exist, as at the very least the
    [document
    element](https://dom.spec.whatwg.org/#document-element){#the-li-element:document-element
    x-internal="document-element"} will always produce a [CSS
    box](https://drafts.csswg.org/css-display/#css-box){#the-li-element:css-box-2
    x-internal="css-box"}.

To determine the [ordinal value]{#ordinal-value .dfn} of each element
owned by a given [list owner](#list-owner){#the-li-element:list-owner-3}
`owner`{.variable}, perform the following steps:

1.  Let `i`{.variable} be 1.

2.  If `owner`{.variable} is an
    [`ol`{#the-li-element:the-ol-element-5}](#the-ol-element) element,
    let `numbering`{.variable} be `owner`{.variable}\'s [starting
    value](#concept-ol-start){#the-li-element:concept-ol-start}.
    Otherwise, let `numbering`{.variable} be 1.

3.  *Loop*: If `i`{.variable} is greater than the number of [list items
    that `owner`{.variable}
    owns](#list-owner){#the-li-element:list-owner-4}, then return; all
    of `owner`{.variable}\'s [owned list
    items](#list-owner){#the-li-element:list-owner-5} have been assigned
    [ordinal values](#ordinal-value){#the-li-element:ordinal-value-3}.

4.  Let `item`{.variable} be the `i`{.variable}th of
    `owner`{.variable}\'s [owned list
    items](#list-owner){#the-li-element:list-owner-6}, in [tree
    order](https://dom.spec.whatwg.org/#concept-tree-order){#the-li-element:tree-order
    x-internal="tree-order"}.

5.  If `item`{.variable} is an
    [`li`{#the-li-element:the-li-element-7}](#the-li-element) element
    that has a
    [`value`{#the-li-element:attr-li-value-2}](#attr-li-value)
    attribute, then:

    1.  Let `parsed`{.variable} be the result of [parsing the value of
        the attribute as an
        integer](common-microsyntaxes.html#rules-for-parsing-integers){#the-li-element:rules-for-parsing-integers}.

    2.  If `parsed`{.variable} is not an error, then set
        `numbering`{.variable} to `parsed`{.variable}.

6.  The [ordinal value](#ordinal-value){#the-li-element:ordinal-value-4}
    of `item`{.variable} is `numbering`{.variable}.

7.  If `owner`{.variable} is an
    [`ol`{#the-li-element:the-ol-element-6}](#the-ol-element) element,
    and `owner`{.variable} has a
    [`reversed`{#the-li-element:attr-ol-reversed}](#attr-ol-reversed)
    attribute, decrement `numbering`{.variable} by 1; otherwise,
    increment `numbering`{.variable} by 1.

8.  Increment `i`{.variable} by 1.

9.  Go to the step labeled *loop*.

------------------------------------------------------------------------

The [`value`]{#dom-li-value .dfn dfn-for="HTMLLIElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-li-element:reflect}
the value of the
[`value`{#the-li-element:attr-li-value-3}](#attr-li-value) content
attribute.

::: example
The element\'s [`value`{#the-li-element:dom-li-value-2}](#dom-li-value)
IDL attribute does not directly correspond to its [ordinal
value](#ordinal-value){#the-li-element:ordinal-value-5}; it simply
[reflects](common-dom-interfaces.html#reflect){#the-li-element:reflect-2}
the content attribute. For example, given this list:

``` html
<ol>
 <li>Item 1
 <li value="3">Item 3
 <li>Item 4
</ol>
```

The [ordinal values](#ordinal-value){#the-li-element:ordinal-value-6}
are 1, 3, and 4, whereas the
[`value`{#the-li-element:dom-li-value-3}](#dom-li-value) IDL attributes
return 0, 3, 0 on getting.
:::

::: example
The following example, the top ten movies are listed (in reverse order).
Note the way the list is given a title by using a
[`figure`{#the-li-element:the-figure-element}](#the-figure-element)
element and its
[`figcaption`{#the-li-element:the-figcaption-element}](#the-figcaption-element)
element.

``` html
<figure>
 <figcaption>The top 10 movies of all time</figcaption>
 <ol>
  <li value="10"><cite>Josie and the Pussycats</cite>, 2001</li>
  <li value="9"><cite lang="sh">Црна мачка, бели мачор</cite>, 1998</li>
  <li value="8"><cite>A Bug's Life</cite>, 1998</li>
  <li value="7"><cite>Toy Story</cite>, 1995</li>
  <li value="6"><cite>Monsters, Inc</cite>, 2001</li>
  <li value="5"><cite>Cars</cite>, 2006</li>
  <li value="4"><cite>Toy Story 2</cite>, 1999</li>
  <li value="3"><cite>Finding Nemo</cite>, 2003</li>
  <li value="2"><cite>The Incredibles</cite>, 2004</li>
  <li value="1"><cite>Ratatouille</cite>, 2007</li>
 </ol>
</figure>
```

The markup could also be written as follows, using the
[`reversed`{#the-li-element:attr-ol-reversed-2}](#attr-ol-reversed)
attribute on the
[`ol`{#the-li-element:the-ol-element-7}](#the-ol-element) element:

``` html
<figure>
 <figcaption>The top 10 movies of all time</figcaption>
 <ol reversed>
  <li><cite>Josie and the Pussycats</cite>, 2001</li>
  <li><cite lang="sh">Црна мачка, бели мачор</cite>, 1998</li>
  <li><cite>A Bug's Life</cite>, 1998</li>
  <li><cite>Toy Story</cite>, 1995</li>
  <li><cite>Monsters, Inc</cite>, 2001</li>
  <li><cite>Cars</cite>, 2006</li>
  <li><cite>Toy Story 2</cite>, 1999</li>
  <li><cite>Finding Nemo</cite>, 2003</li>
  <li><cite>The Incredibles</cite>, 2004</li>
  <li><cite>Ratatouille</cite>, 2007</li>
 </ol>
</figure>
```
:::

While it is conforming to include heading elements (e.g.
[`h1`{#the-li-element:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements))
inside [`li`{#the-li-element:the-li-element-8}](#the-li-element)
elements, it likely does not convey the semantics that the author
intended. A heading starts a new section, so a heading in a list
implicitly splits the list into spanning multiple sections.

#### [4.4.9]{.secno} The [`dl`]{.dfn dfn-type="element"} element[](#the-dl-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/dl](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dl "The <dl> HTML element represents a description list. The element encloses a list of groups of terms (specified using the <dt> element) and descriptions (provided by <dd> elements). Common uses for this element are to implement a glossary or to display metadata (a list of key-value pairs).")

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

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLDListElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLDListElement "The HTMLDListElement interface provides special properties (beyond those of the regular HTMLElement interface it also has available to it by inheritance) for manipulating definition list (<dl>) elements.")

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

[Categories](dom.html#concept-element-categories){#the-dl-element:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-dl-element:flow-content-2}.
:   If the element\'s children include at least one name-value group:
    [Palpable
    content](dom.html#palpable-content-2){#the-dl-element:palpable-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-dl-element:concept-element-contexts}:
:   Where [flow
    content](dom.html#flow-content-2){#the-dl-element:flow-content-2-2}
    is expected.

[Content model](dom.html#concept-element-content-model){#the-dl-element:concept-element-content-model}:
:   Either: Zero or more groups each consisting of one or more
    [`dt`{#the-dl-element:the-dt-element}](#the-dt-element) elements
    followed by one or more
    [`dd`{#the-dl-element:the-dd-element}](#the-dd-element) elements,
    optionally intermixed with [script-supporting
    elements](dom.html#script-supporting-elements-2){#the-dl-element:script-supporting-elements-2}.
:   Or: One or more
    [`div`{#the-dl-element:the-div-element}](#the-div-element) elements,
    optionally intermixed with [script-supporting
    elements](dom.html#script-supporting-elements-2){#the-dl-element:script-supporting-elements-2-2}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-dl-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-dl-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-dl-element:global-attributes}

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-dl-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-dl).
:   [For implementers](https://w3c.github.io/html-aam/#el-dl).

[DOM interface](dom.html#concept-element-dom){#the-dl-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLDListElement : HTMLElement {
      [HTMLConstructor] constructor();

      // also has obsolete members
    };
    ```

The [`dl`{#the-dl-element:the-dl-element}](#the-dl-element) element
[represents](dom.html#represents){#the-dl-element:represents} an
association list consisting of zero or more name-value groups (a
description list). A name-value group consists of one or more names
([`dt`{#the-dl-element:the-dt-element-2}](#the-dt-element) elements,
possibly as children of a
[`div`{#the-dl-element:the-div-element-2}](#the-div-element) element
child) followed by one or more values
([`dd`{#the-dl-element:the-dd-element-2}](#the-dd-element) elements,
possibly as children of a
[`div`{#the-dl-element:the-div-element-3}](#the-div-element) element
child), ignoring any nodes other than
[`dt`{#the-dl-element:the-dt-element-3}](#the-dt-element) and
[`dd`{#the-dl-element:the-dd-element-3}](#the-dd-element) element
children, and [`dt`{#the-dl-element:the-dt-element-4}](#the-dt-element)
and [`dd`{#the-dl-element:the-dd-element-4}](#the-dd-element) elements
that are children of
[`div`{#the-dl-element:the-div-element-4}](#the-div-element) element
children. Within a single
[`dl`{#the-dl-element:the-dl-element-2}](#the-dl-element) element, there
should not be more than one
[`dt`{#the-dl-element:the-dt-element-5}](#the-dt-element) element for
each name.

Name-value groups may be terms and definitions, metadata topics and
values, questions and answers, or any other groups of name-value data.

The values within a group are alternatives; multiple paragraphs forming
part of the same value must all be given within the same
[`dd`{#the-dl-element:the-dd-element-5}](#the-dd-element) element.

The order of the list of groups, and of the names and values within each
group, may be significant.

In order to annotate groups with
[microdata](microdata.html#microdata){#the-dl-element:microdata}
attributes, or other [global
attributes](dom.html#global-attributes){#the-dl-element:global-attributes-2}
that apply to whole groups, or just for styling purposes, each group in
a [`dl`{#the-dl-element:the-dl-element-3}](#the-dl-element) element can
be wrapped in a
[`div`{#the-dl-element:the-div-element-5}](#the-div-element) element.
This does not change the semantics of the
[`dl`{#the-dl-element:the-dl-element-4}](#the-dl-element) element.

The name-value groups of a
[`dl`{#the-dl-element:the-dl-element-5}](#the-dl-element) element
`dl`{.variable} are determined using the following algorithm. A
name-value group has a name (a list of
[`dt`{#the-dl-element:the-dt-element-6}](#the-dt-element) elements,
initially empty) and a value (a list of
[`dd`{#the-dl-element:the-dd-element-6}](#the-dd-element) elements,
initially empty).

1.  Let `groups`{.variable} be an empty list of name-value groups.

2.  Let `current`{.variable} be a new name-value group.

3.  Let `seenDd`{.variable} be false.

4.  Let `child`{.variable} be `dl`{.variable}\'s [first
    child](https://dom.spec.whatwg.org/#concept-tree-first-child){#the-dl-element:first-child
    x-internal="first-child"}.

5.  Let `grandchild`{.variable} be null.

6.  While `child`{.variable} is not null:

    1.  If `child`{.variable} is a
        [`div`{#the-dl-element:the-div-element-6}](#the-div-element)
        element, then:

        1.  Let `grandchild`{.variable} be `child`{.variable}\'s [first
            child](https://dom.spec.whatwg.org/#concept-tree-first-child){#the-dl-element:first-child-2
            x-internal="first-child"}.

        2.  While `grandchild`{.variable} is not null:

            1.  [Process `dt` or
                `dd`](#process-dt-or-dd){#the-dl-element:process-dt-or-dd}
                for `grandchild`{.variable}.

            2.  Set `grandchild`{.variable} to
                `grandchild`{.variable}\'s [next
                sibling](https://dom.spec.whatwg.org/#concept-tree-next-sibling){#the-dl-element:next-sibling
                x-internal="next-sibling"}.

    2.  Otherwise, [process `dt` or
        `dd`](#process-dt-or-dd){#the-dl-element:process-dt-or-dd-2} for
        `child`{.variable}.

    3.  Set `child`{.variable} to `child`{.variable}\'s [next
        sibling](https://dom.spec.whatwg.org/#concept-tree-next-sibling){#the-dl-element:next-sibling-2
        x-internal="next-sibling"}.

7.  If `current`{.variable} is not empty, then append
    `current`{.variable} to `groups`{.variable}.

8.  Return `groups`{.variable}.

To [process `dt` or `dd`]{#process-dt-or-dd .dfn} for a node
`node`{.variable} means to follow these steps:

1.  Let `groups`{.variable}, `current`{.variable}, and
    `seenDd`{.variable} be the same variables as those of the same name
    in the algorithm that invoked these steps.

2.  If `node`{.variable} is a
    [`dt`{#the-dl-element:the-dt-element-7}](#the-dt-element) element,
    then:

    1.  If `seenDd`{.variable} is true, then append `current`{.variable}
        to `groups`{.variable}, set `current`{.variable} to a new
        name-value group, and set `seenDd`{.variable} to false.

    2.  Append `node`{.variable} to `current`{.variable}\'s name.

3.  Otherwise, if `node`{.variable} is a
    [`dd`{#the-dl-element:the-dd-element-7}](#the-dd-element) element,
    then append `node`{.variable} to `current`{.variable}\'s value and
    set `seenDd`{.variable} to true.

When a name-value group has an empty list as name or value, it is often
due to accidentally using
[`dd`{#the-dl-element:the-dd-element-8}](#the-dd-element) elements in
the place of [`dt`{#the-dl-element:the-dt-element-8}](#the-dt-element)
elements and vice versa. Conformance checkers can spot such mistakes and
might be able to advise authors how to correctly use the markup.

::: example
In the following example, one entry (\"Authors\") is linked to two
values (\"John\" and \"Luke\").

``` html
<dl>
 <dt> Authors
 <dd> John
 <dd> Luke
 <dt> Editor
 <dd> Frank
</dl>
```
:::

::: example
In the following example, one definition is linked to two terms.

``` html
<dl>
 <dt lang="en-US"> <dfn>color</dfn> </dt>
 <dt lang="en-GB"> <dfn>colour</dfn> </dt>
 <dd> A sensation which (in humans) derives from the ability of
 the fine structure of the eye to distinguish three differently
 filtered analyses of a view. </dd>
</dl>
```
:::

::: example
The following example illustrates the use of the
[`dl`{#the-dl-element:the-dl-element-6}](#the-dl-element) element to
mark up metadata of sorts. At the end of the example, one group has two
metadata labels (\"Authors\" and \"Editors\") and two values (\"Robert
Rothman\" and \"Daniel Jackson\"). This example also uses the
[`div`{#the-dl-element:the-div-element-7}](#the-div-element) element
around the groups of
[`dt`{#the-dl-element:the-dt-element-9}](#the-dt-element) and
[`dd`{#the-dl-element:the-dd-element-9}](#the-dd-element) element, to
aid with styling.

``` html
<dl>
 <div>
  <dt> Last modified time </dt>
  <dd> 2004-12-23T23:33Z </dd>
 </div>
 <div>
  <dt> Recommended update interval </dt>
  <dd> 60s </dd>
 </div>
 <div>
  <dt> Authors </dt>
  <dt> Editors </dt>
  <dd> Robert Rothman </dd>
  <dd> Daniel Jackson </dd>
 </div>
</dl>
```
:::

::: example
The following example shows the
[`dl`{#the-dl-element:the-dl-element-7}](#the-dl-element) element used
to give a set of instructions. The order of the instructions here is
important (in the other examples, the order of the blocks was not
important).

``` html
<p>Determine the victory points as follows (use the
first matching case):</p>
<dl>
 <dt> If you have exactly five gold coins </dt>
 <dd> You get five victory points </dd>
 <dt> If you have one or more gold coins, and you have one or more silver coins </dt>
 <dd> You get two victory points </dd>
 <dt> If you have one or more silver coins </dt>
 <dd> You get one victory point </dd>
 <dt> Otherwise </dt>
 <dd> You get no victory points </dd>
</dl>
```
:::

::: example
The following snippet shows a
[`dl`{#the-dl-element:the-dl-element-8}](#the-dl-element) element being
used as a glossary. Note the use of
[`dfn`{#the-dl-element:the-dfn-element}](text-level-semantics.html#the-dfn-element)
to indicate the word being defined.

``` html
<dl>
 <dt><dfn>Apartment</dfn>, n.</dt>
 <dd>An execution context grouping one or more threads with one or
 more COM objects.</dd>
 <dt><dfn>Flat</dfn>, n.</dt>
 <dd>A deflated tire.</dd>
 <dt><dfn>Home</dfn>, n.</dt>
 <dd>The user's login directory.</dd>
</dl>
```
:::

::: example
This example uses
[microdata](microdata.html#microdata){#the-dl-element:microdata-2}
attributes in a
[`dl`{#the-dl-element:the-dl-element-9}](#the-dl-element) element,
together with the
[`div`{#the-dl-element:the-div-element-8}](#the-div-element) element, to
annotate the ice cream desserts at a French restaurant.

``` {lang="fr"}
<dl>
 <div itemscope itemtype="http://schema.org/Product">
  <dt itemprop="name">Café ou Chocolat Liégeois
  <dd itemprop="offers" itemscope itemtype="http://schema.org/Offer">
   <span itemprop="price">3.50</span>
   <data itemprop="priceCurrency" value="EUR">€</data>
  <dd itemprop="description">
   2 boules Café ou Chocolat, 1 boule Vanille, sauce café ou chocolat, chantilly
 </div>

 <div itemscope itemtype="http://schema.org/Product">
  <dt itemprop="name">Américaine
  <dd itemprop="offers" itemscope itemtype="http://schema.org/Offer">
   <span itemprop="price">3.50</span>
   <data itemprop="priceCurrency" value="EUR">€</data>
  <dd itemprop="description">
   1 boule Crème brûlée, 1 boule Vanille, 1 boule Caramel, chantilly
 </div>
</dl>
```

Without the [`div`{#the-dl-element:the-div-element-9}](#the-div-element)
element the markup would need to use the
[`itemref`{#the-dl-element:attr-itemref}](microdata.html#attr-itemref)
attribute to link the data in the
[`dd`{#the-dl-element:the-dd-element-10}](#the-dd-element) elements with
the item, as follows.

``` {lang="fr"}
<dl>
 <dt itemscope itemtype="http://schema.org/Product" itemref="1-offer 1-description">
  <span itemprop="name">Café ou Chocolat Liégeois</span>
 <dd id="1-offer" itemprop="offers" itemscope itemtype="http://schema.org/Offer">
  <span itemprop="price">3.50</span>
  <data itemprop="priceCurrency" value="EUR">€</data>
 <dd id="1-description" itemprop="description">
  2 boules Café ou Chocolat, 1 boule Vanille, sauce café ou chocolat, chantilly

 <dt itemscope itemtype="http://schema.org/Product" itemref="2-offer 2-description">
  <span itemprop="name">Américaine</span>
 <dd id="2-offer" itemprop="offers" itemscope itemtype="http://schema.org/Offer">
  <span itemprop="price">3.50</span>
  <data itemprop="priceCurrency" value="EUR">€</data>
 <dd id="2-description" itemprop="description">
  1 boule Crème brûlée, 1 boule Vanille, 1 boule Caramel, chantilly
</dl>
```
:::

The [`dl`{#the-dl-element:the-dl-element-10}](#the-dl-element) element
is inappropriate for marking up dialogue. See some [examples of how to
mark up dialogue](semantics-other.html#conversations).

#### [4.4.10]{.secno} The [`dt`]{.dfn dfn-type="element"} element[](#the-dt-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/dt](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dt "The <dt> HTML element specifies a term in a description or definition list, and as such must be used inside a <dl> element. It is usually followed by a <dd> element; however, multiple <dt> elements in a row indicate several terms that are all defined by the immediate next <dd> element.")

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

[Categories](dom.html#concept-element-categories){#the-dt-element:concept-element-categories}:
:   None.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-dt-element:concept-element-contexts}:
:   Before [`dd`{#the-dt-element:the-dd-element}](#the-dd-element) or
    [`dt`{#the-dt-element:the-dt-element}](#the-dt-element) elements
    inside [`dl`{#the-dt-element:the-dl-element}](#the-dl-element)
    elements.
:   Before [`dd`{#the-dt-element:the-dd-element-2}](#the-dd-element) or
    [`dt`{#the-dt-element:the-dt-element-2}](#the-dt-element) elements
    inside [`div`{#the-dt-element:the-div-element}](#the-div-element)
    elements that are children of a
    [`dl`{#the-dt-element:the-dl-element-2}](#the-dl-element) element.

[Content model](dom.html#concept-element-content-model){#the-dt-element:concept-element-content-model}:
:   [Flow
    content](dom.html#flow-content-2){#the-dt-element:flow-content-2},
    but with no
    [`header`{#the-dt-element:the-header-element}](sections.html#the-header-element),
    [`footer`{#the-dt-element:the-footer-element}](sections.html#the-footer-element),
    [sectioning
    content](dom.html#sectioning-content-2){#the-dt-element:sectioning-content-2},
    or [heading
    content](dom.html#heading-content-2){#the-dt-element:heading-content-2}
    descendants.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-dt-element:concept-element-tag-omission}:
:   A [`dt`{#the-dt-element:the-dt-element-3}](#the-dt-element)
    element\'s [end
    tag](syntax.html#syntax-end-tag){#the-dt-element:syntax-end-tag} can
    be omitted if the
    [`dt`{#the-dt-element:the-dt-element-4}](#the-dt-element) element is
    immediately followed by another
    [`dt`{#the-dt-element:the-dt-element-5}](#the-dt-element) element or
    a [`dd`{#the-dt-element:the-dd-element-3}](#the-dd-element) element.

[Content attributes](dom.html#concept-element-attributes){#the-dt-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-dt-element:global-attributes}

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-dt-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-dt).
:   [For implementers](https://w3c.github.io/html-aam/#el-dt).

[DOM interface](dom.html#concept-element-dom){#the-dt-element:concept-element-dom}:
:   Uses
    [`HTMLElement`{#the-dt-element:htmlelement}](dom.html#htmlelement).

The [`dt`{#the-dt-element:the-dt-element-6}](#the-dt-element) element
[represents](dom.html#represents){#the-dt-element:represents} the term,
or name, part of a term-description group in a description list
([`dl`{#the-dt-element:the-dl-element-3}](#the-dl-element) element).

The [`dt`{#the-dt-element:the-dt-element-7}](#the-dt-element) element
itself, when used in a
[`dl`{#the-dt-element:the-dl-element-4}](#the-dl-element) element, does
not indicate that its contents are a term being defined, but this can be
indicated using the
[`dfn`{#the-dt-element:the-dfn-element}](text-level-semantics.html#the-dfn-element)
element.

::: example
This example shows a list of frequently asked questions (a FAQ) marked
up using the [`dt`{#the-dt-element:the-dt-element-8}](#the-dt-element)
element for questions and the
[`dd`{#the-dt-element:the-dd-element-4}](#the-dd-element) element for
answers.

``` html
<article>
 <h1>FAQ</h1>
 <dl>
  <dt>What do we want?</dt>
  <dd>Our data.</dd>
  <dt>When do we want it?</dt>
  <dd>Now.</dd>
  <dt>Where is it?</dt>
  <dd>We are not sure.</dd>
 </dl>
</article>
```
:::

#### [4.4.11]{.secno} The [`dd`]{.dfn dfn-type="element"} element[](#the-dd-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/dd](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dd "The <dd> HTML element provides the description, definition, or value for the preceding term (<dt>) in a description list (<dl>).")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerYes]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android4+]{.firefox_android .yes}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

[Categories](dom.html#concept-element-categories){#the-dd-element:concept-element-categories}:
:   None.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-dd-element:concept-element-contexts}:
:   After [`dt`{#the-dd-element:the-dt-element}](#the-dt-element) or
    [`dd`{#the-dd-element:the-dd-element}](#the-dd-element) elements
    inside [`dl`{#the-dd-element:the-dl-element}](#the-dl-element)
    elements.
:   After [`dt`{#the-dd-element:the-dt-element-2}](#the-dt-element) or
    [`dd`{#the-dd-element:the-dd-element-2}](#the-dd-element) elements
    inside [`div`{#the-dd-element:the-div-element}](#the-div-element)
    elements that are children of a
    [`dl`{#the-dd-element:the-dl-element-2}](#the-dl-element) element.

[Content model](dom.html#concept-element-content-model){#the-dd-element:concept-element-content-model}:
:   [Flow
    content](dom.html#flow-content-2){#the-dd-element:flow-content-2}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-dd-element:concept-element-tag-omission}:
:   A [`dd`{#the-dd-element:the-dd-element-3}](#the-dd-element)
    element\'s [end
    tag](syntax.html#syntax-end-tag){#the-dd-element:syntax-end-tag} can
    be omitted if the
    [`dd`{#the-dd-element:the-dd-element-4}](#the-dd-element) element is
    immediately followed by another
    [`dd`{#the-dd-element:the-dd-element-5}](#the-dd-element) element or
    a [`dt`{#the-dd-element:the-dt-element-3}](#the-dt-element) element,
    or if there is no more content in the parent element.

[Content attributes](dom.html#concept-element-attributes){#the-dd-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-dd-element:global-attributes}

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-dd-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-dd).
:   [For implementers](https://w3c.github.io/html-aam/#el-dd).

[DOM interface](dom.html#concept-element-dom){#the-dd-element:concept-element-dom}:
:   Uses
    [`HTMLElement`{#the-dd-element:htmlelement}](dom.html#htmlelement).

The [`dd`{#the-dd-element:the-dd-element-6}](#the-dd-element) element
[represents](dom.html#represents){#the-dd-element:represents} the
description, definition, or value, part of a term-description group in a
description list
([`dl`{#the-dd-element:the-dl-element-3}](#the-dl-element) element).

::: example
A [`dl`{#the-dd-element:the-dl-element-4}](#the-dl-element) can be used
to define a vocabulary list, like in a dictionary. In the following
example, each entry, given by a
[`dt`{#the-dd-element:the-dt-element-4}](#the-dt-element) with a
[`dfn`{#the-dd-element:the-dfn-element}](text-level-semantics.html#the-dfn-element),
has several [`dd`{#the-dd-element:the-dd-element-7}](#the-dd-element)s,
showing the various parts of the definition.

``` html
<dl>
 <dt><dfn>happiness</dfn></dt>
 <dd class="pronunciation">/ˈhæpinəs/</dd>
 <dd class="part-of-speech"><i><abbr>n.</abbr></i></dd>
 <dd>The state of being happy.</dd>
 <dd>Good fortune; success. <q>Oh <b>happiness</b>! It worked!</q></dd>
 <dt><dfn>rejoice</dfn></dt>
 <dd class="pronunciation">/rɪˈdʒɔɪs/</dd>
 <dd><i class="part-of-speech"><abbr>v.intr.</abbr></i> To be delighted oneself.</dd>
 <dd><i class="part-of-speech"><abbr>v.tr.</abbr></i> To cause one to be delighted.</dd>
</dl>
```
:::

#### [4.4.12]{.secno} The [`figure`]{.dfn dfn-type="element"} element[](#the-figure-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/figure](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/figure "The <figure> HTML element represents self-contained content, potentially with an optional caption, which is specified using the <figcaption> element. The figure, its caption, and its contents are referenced as a single unit.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5.1+]{.safari .yes}[Chrome8+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

[Categories](dom.html#concept-element-categories){#the-figure-element:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-figure-element:flow-content-2}.
:   [Palpable
    content](dom.html#palpable-content-2){#the-figure-element:palpable-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-figure-element:concept-element-contexts}:
:   Where [flow
    content](dom.html#flow-content-2){#the-figure-element:flow-content-2-2}
    is expected.

[Content model](dom.html#concept-element-content-model){#the-figure-element:concept-element-content-model}:
:   Either: one
    [`figcaption`{#the-figure-element:the-figcaption-element}](#the-figcaption-element)
    element followed by [flow
    content](dom.html#flow-content-2){#the-figure-element:flow-content-2-3}.
:   Or: [flow
    content](dom.html#flow-content-2){#the-figure-element:flow-content-2-4}
    followed by one
    [`figcaption`{#the-figure-element:the-figcaption-element-2}](#the-figcaption-element)
    element.
:   Or: [flow
    content](dom.html#flow-content-2){#the-figure-element:flow-content-2-5}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-figure-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-figure-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-figure-element:global-attributes}

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-figure-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-figure).
:   [For implementers](https://w3c.github.io/html-aam/#el-figure).

[DOM interface](dom.html#concept-element-dom){#the-figure-element:concept-element-dom}:
:   Uses
    [`HTMLElement`{#the-figure-element:htmlelement}](dom.html#htmlelement).

The
[`figure`{#the-figure-element:the-figure-element}](#the-figure-element)
element
[represents](dom.html#represents){#the-figure-element:represents} some
[flow
content](dom.html#flow-content-2){#the-figure-element:flow-content-2-6},
optionally with a caption, that is self-contained (like a complete
sentence) and is typically
[referenced](dom.html#referenced){#the-figure-element:referenced} as a
single unit from the main flow of the document.

\"Self-contained\" in this context does not necessarily mean
independent. For example, each sentence in a paragraph is
self-contained; an image that is part of a sentence would be
inappropriate for
[`figure`{#the-figure-element:the-figure-element-2}](#the-figure-element),
but an entire sentence made of images would be fitting.

The element can thus be used to annotate illustrations, diagrams,
photos, code listings, etc.

::: {#figure-note-about-references .note}
When a
[`figure`{#the-figure-element:the-figure-element-3}](#the-figure-element)
is referred to from the main content of the document by identifying it
by its caption (e.g., by figure number), it enables such content to be
easily moved away from that primary content, e.g., to the side of the
page, to dedicated pages, or to an appendix, without affecting the flow
of the document.

If a
[`figure`{#the-figure-element:the-figure-element-4}](#the-figure-element)
element is
[referenced](dom.html#referenced){#the-figure-element:referenced-2} by
its relative position, e.g., \"in the photograph above\" or \"as the
next figure shows\", then moving the figure would disrupt the page\'s
meaning. Authors are encouraged to consider using labels to refer to
figures, rather than using such relative references, so that the page
can easily be restyled without affecting the page\'s meaning.
:::

The first
[`figcaption`{#the-figure-element:the-figcaption-element-3}](#the-figcaption-element)
element child of the element, if any, represents the caption of the
[`figure`{#the-figure-element:the-figure-element-5}](#the-figure-element)
element\'s contents. If there is no child
[`figcaption`{#the-figure-element:the-figcaption-element-4}](#the-figcaption-element)
element, then there is no caption.

A
[`figure`{#the-figure-element:the-figure-element-6}](#the-figure-element)
element\'s contents are part of the surrounding flow. If the purpose of
the page is to display the figure, for example a photograph on an image
sharing site, the
[`figure`{#the-figure-element:the-figure-element-7}](#the-figure-element)
and
[`figcaption`{#the-figure-element:the-figcaption-element-5}](#the-figcaption-element)
elements can be used to explicitly provide a caption for that figure.
For content that is only tangentially related, or that serves a separate
purpose than the surrounding flow, the
[`aside`{#the-figure-element:the-aside-element}](sections.html#the-aside-element)
element should be used (and can itself wrap a
[`figure`{#the-figure-element:the-figure-element-8}](#the-figure-element)).
For example, a pull quote that repeats content from an
[`article`{#the-figure-element:the-article-element}](sections.html#the-article-element)
would be more appropriate in an
[`aside`{#the-figure-element:the-aside-element-2}](sections.html#the-aside-element)
than in a
[`figure`{#the-figure-element:the-figure-element-9}](#the-figure-element),
because it isn\'t part of the content, it\'s a repetition of the content
for the purposes of enticing readers or highlighting key topics.

::: example
This example shows the
[`figure`{#the-figure-element:the-figure-element-10}](#the-figure-element)
element to mark up a code listing.

``` html
<p>In <a href="#l4">listing 4</a> we see the primary core interface
API declaration.</p>
<figure id="l4">
 <figcaption>Listing 4. The primary core interface API declaration.</figcaption>
 <pre><code>interface PrimaryCore {
 boolean verifyDataLine();
 undefined sendData(sequence&lt;byte> data);
 undefined initSelfDestruct();
}</code></pre>
</figure>
<p>The API is designed to use UTF-8.</p>
```
:::

::: example
Here we see a
[`figure`{#the-figure-element:the-figure-element-11}](#the-figure-element)
element to mark up a photo that is the main content of the page (as in a
gallery).

``` html
<!DOCTYPE HTML>
<html lang="en">
<title>Bubbles at work — My Gallery™</title>
<figure>
 <img src="bubbles-work.jpeg"
      alt="Bubbles, sitting in his office chair, works on his
           latest project intently.">
 <figcaption>Bubbles at work</figcaption>
</figure>
<nav><a href="19414.html">Prev</a> — <a href="19416.html">Next</a></nav>
```
:::

::: example
In this example, we see an image that is *not* a figure, as well as an
image and a video that are. The first image is literally part of the
example\'s second sentence, so it\'s not a self-contained unit, and thus
[`figure`{#the-figure-element:the-figure-element-12}](#the-figure-element)
would be inappropriate.

``` html
<h2>Malinko's comics</h2>

<p>This case centered on some sort of "intellectual property"
infringement related to a comic (see Exhibit A). The suit started
after a trailer ending with these words:

<blockquote>
 <img src="promblem-packed-action.png" alt="ROUGH COPY! Promblem-Packed Action!">
</blockquote>

<p>...was aired. A lawyer, armed with a Bigger Notebook, launched a
preemptive strike using snowballs. A complete copy of the trailer is
included with Exhibit B.

<figure>
 <img src="ex-a.png" alt="Two squiggles on a dirty piece of paper.">
 <figcaption>Exhibit A. The alleged <cite>rough copy</cite> comic.</figcaption>
</figure>

<figure>
 <video src="ex-b.mov"></video>
 <figcaption>Exhibit B. The <cite>Rough Copy</cite> trailer.</figcaption>
</figure>

<p>The case was resolved out of court.
```
:::

::: example
Here, a part of a poem is marked up using
[`figure`{#the-figure-element:the-figure-element-13}](#the-figure-element).

``` html
<figure>
 <p>'Twas brillig, and the slithy toves<br>
 Did gyre and gimble in the wabe;<br>
 All mimsy were the borogoves,<br>
 And the mome raths outgrabe.</p>
 <figcaption><cite>Jabberwocky</cite> (first verse). Lewis Carroll, 1832-98</figcaption>
</figure>
```
:::

::: example
In this example, which could be part of a much larger work discussing a
castle, nested
[`figure`{#the-figure-element:the-figure-element-14}](#the-figure-element)
elements are used to provide both a group caption and individual
captions for each figure in the group:

``` html
<figure>
 <figcaption>The castle through the ages: 1423, 1858, and 1999 respectively.</figcaption>
 <figure>
  <figcaption>Etching. Anonymous, ca. 1423.</figcaption>
  <img src="castle1423.jpeg" alt="The castle has one tower, and a tall wall around it.">
 </figure>
 <figure>
  <figcaption>Oil-based paint on canvas. Maria Towle, 1858.</figcaption>
  <img src="castle1858.jpeg" alt="The castle now has two towers and two walls.">
 </figure>
 <figure>
  <figcaption>Film photograph. Peter Jankle, 1999.</figcaption>
  <img src="castle1999.jpeg" alt="The castle lies in ruins, the original tower all that remains in one piece.">
 </figure>
</figure>
```
:::

::: example
The previous example could also be more succinctly written as follows
(using [`title`{#the-figure-element:attr-title}](dom.html#attr-title)
attributes in place of the nested
[`figure`{#the-figure-element:the-figure-element-15}](#the-figure-element)/[`figcaption`{#the-figure-element:the-figcaption-element-6}](#the-figcaption-element)
pairs):

``` html
<figure>
 <img src="castle1423.jpeg" title="Etching. Anonymous, ca. 1423."
      alt="The castle has one tower, and a tall wall around it.">
 <img src="castle1858.jpeg" title="Oil-based paint on canvas. Maria Towle, 1858."
      alt="The castle now has two towers and two walls.">
 <img src="castle1999.jpeg" title="Film photograph. Peter Jankle, 1999."
      alt="The castle lies in ruins, the original tower all that remains in one piece.">
 <figcaption>The castle through the ages: 1423, 1858, and 1999 respectively.</figcaption>
</figure>
```
:::

::: example
The figure is sometimes
[referenced](dom.html#referenced){#the-figure-element:referenced-3} only
implicitly from the content:

``` html
<article>
 <h1>Fiscal negotiations stumble in Congress as deadline nears</h1>
 <figure>
  <img src="obama-reid.jpeg" alt="Obama and Reid sit together smiling in the Oval Office.">
  <figcaption>Barack Obama and Harry Reid. White House press photograph.</figcaption>
 </figure>
 <p>Negotiations in Congress to end the fiscal impasse sputtered on Tuesday, leaving both chambers
 grasping for a way to reopen the government and raise the country's borrowing authority with a
 Thursday deadline drawing near.</p>
 ...
</article>
```
:::

#### [4.4.13]{.secno} The [`figcaption`]{.dfn dfn-type="element"} element[](#the-figcaption-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/figcaption](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/figcaption "The <figcaption> HTML element represents a caption or legend describing the rest of the contents of its parent <figure> element.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5.1+]{.safari .yes}[Chrome8+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

[Categories](dom.html#concept-element-categories){#the-figcaption-element:concept-element-categories}:
:   None.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-figcaption-element:concept-element-contexts}:
:   As the first or last child of a
    [`figure`{#the-figcaption-element:the-figure-element}](#the-figure-element)
    element.

[Content model](dom.html#concept-element-content-model){#the-figcaption-element:concept-element-content-model}:
:   [Flow
    content](dom.html#flow-content-2){#the-figcaption-element:flow-content-2}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-figcaption-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-figcaption-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-figcaption-element:global-attributes}

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-figcaption-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-figcaption).
:   [For implementers](https://w3c.github.io/html-aam/#el-figcaption).

[DOM interface](dom.html#concept-element-dom){#the-figcaption-element:concept-element-dom}:
:   Uses
    [`HTMLElement`{#the-figcaption-element:htmlelement}](dom.html#htmlelement).

The
[`figcaption`{#the-figcaption-element:the-figcaption-element}](#the-figcaption-element)
element
[represents](dom.html#represents){#the-figcaption-element:represents} a
caption or legend for the rest of the contents of the
[`figcaption`{#the-figcaption-element:the-figcaption-element-2}](#the-figcaption-element)
element\'s parent
[`figure`{#the-figcaption-element:the-figure-element-2}](#the-figure-element)
element, if any.

::: example
The element can contain additional information about the source:

``` html
<figcaption>
 <p>A duck.</p>
 <p><small>Photograph courtesy of 🌟 News.</small></p>
</figcaption>
```

``` html
<figcaption>
 <p>Average rent for 3-room apartments, excluding non-profit apartments</p>
 <p>Zürich’s Statistics Office — <time datetime=2017-11-14>14 November 2017</time></p>
</figcaption>
```
:::

#### [4.4.14]{.secno} The [`main`]{.dfn dfn-type="element"} element[](#the-main-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/main](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/main "The <main> HTML element represents the dominant content of the <body> of a document. The main content area consists of content that is directly related to or expands upon the central topic of a document, or the central functionality of an application.")

Support in all current engines.

::: support
[Firefox21+]{.firefox .yes}[Safari7+]{.safari .yes}[Chrome26+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera16+]{.opera .yes}[Edge79+]{.edge_blink .yes}

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

[Categories](dom.html#concept-element-categories){#the-main-element:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-main-element:flow-content-2}.
:   [Palpable
    content](dom.html#palpable-content-2){#the-main-element:palpable-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-main-element:concept-element-contexts}:
:   Where [flow
    content](dom.html#flow-content-2){#the-main-element:flow-content-2-2}
    is expected, but only if it is a [hierarchically correct `main`
    element](#hierarchically-correct-main-element){#the-main-element:hierarchically-correct-main-element}.

[Content model](dom.html#concept-element-content-model){#the-main-element:concept-element-content-model}:
:   [Flow
    content](dom.html#flow-content-2){#the-main-element:flow-content-2-3}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-main-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-main-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-main-element:global-attributes}

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-main-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-main).
:   [For implementers](https://w3c.github.io/html-aam/#el-main).

[DOM interface](dom.html#concept-element-dom){#the-main-element:concept-element-dom}:
:   Uses
    [`HTMLElement`{#the-main-element:htmlelement}](dom.html#htmlelement).

The [`main`{#the-main-element:the-main-element}](#the-main-element)
element [represents](dom.html#represents){#the-main-element:represents}
the dominant contents of the document.

A document must not have more than one
[`main`{#the-main-element:the-main-element-2}](#the-main-element)
element that does not have the
[`hidden`{#the-main-element:attr-hidden}](interaction.html#attr-hidden)
attribute specified.

A [hierarchically correct `main`
element]{#hierarchically-correct-main-element .dfn} is one whose
ancestor elements are limited to
[`html`{#the-main-element:the-html-element}](semantics.html#the-html-element),
[`body`{#the-main-element:the-body-element}](sections.html#the-body-element),
[`div`{#the-main-element:the-div-element}](#the-div-element),
[`form`{#the-main-element:the-form-element}](forms.html#the-form-element)
without an [accessible
name](https://w3c.github.io/aria/#dfn-accessible-name){#the-main-element:concept-accessible-name
x-internal="concept-accessible-name"}, and [autonomous custom
elements](custom-elements.html#autonomous-custom-element){#the-main-element:autonomous-custom-element}.
Each [`main`{#the-main-element:the-main-element-3}](#the-main-element)
element must be a [hierarchically correct `main`
element](#hierarchically-correct-main-element){#the-main-element:hierarchically-correct-main-element-2}.

::: example
In this example, the author has used a presentation where each component
of the page is rendered in a box. To wrap the main content of the page
(as opposed to the header, the footer, the navigation bar, and a
sidebar), the
[`main`{#the-main-element:the-main-element-4}](#the-main-element)
element is used.

``` html
<!DOCTYPE html>
<html lang="en">
<title>RPG System 17</title>
<style>
 header, nav, aside, main, footer {
   margin: 0.5em; border: thin solid; padding: 0.5em;
   background: #EFF; color: black; box-shadow: 0 0 0.25em #033;
 }
 h1, h2, p { margin: 0; }
 nav, main { float: left; }
 aside { float: right; }
 footer { clear: both; }
</style>
<header>
 <h1>System Eighteen</h1>
</header>
<nav>
 <a href="../16/">← System 17</a>
 <a href="../18/">RPXIX →</a>
</nav>
<aside>
 <p>This system has no HP mechanic, so there's no healing.
</aside>
<main>
 <h2>Character creation</h2>
 <p>Attributes (magic, strength, agility) are purchased at the cost of one point per level.</p>
 <h2>Rolls</h2>
 <p>Each encounter, roll the dice for all your skills. If you roll more than the opponent, you win.</p>
</main>
<footer>
 <p>Copyright © 2013
</footer>
</html>
```

In the following example, multiple
[`main`{#the-main-element:the-main-element-5}](#the-main-element)
elements are used and script is used to make navigation work without a
server roundtrip and to set the
[`hidden`{#the-main-element:attr-hidden-2}](interaction.html#attr-hidden)
attribute on those that are not current:

``` html
<!doctype html>
<html lang=en-CA>
<meta charset=utf-8>
<title> … </title>
<link rel=stylesheet href=spa.css>
<script src=spa.js async></script>
<nav>
 <a href=/>Home</a>
 <a href=/about>About</a>
 <a href=/contact>Contact</a>
</nav>
<main>
 <h1>Home</h1>
 …
</main>
<main hidden>
 <h1>About</h1>
 …
</main>
<main hidden>
 <h1>Contact</h1>
 …
</main>
<footer>Made with ❤️ by <a href=https://example.com/>Example 👻</a>.</footer>
```
:::

#### [4.4.15]{.secno} The [`search`]{.dfn dfn-type="element"} element[](#the-search-element){.self-link}

::::: {.mdn-anno .wrapped}
**⚠**MDN

:::: feature
[Element/search](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/search "The <search> HTML element is a container representing the parts of the document or application with form controls or other content related to performing a search or filtering operation. The <search> element semantically identifies the purpose of the element's contents as having search or filtering capabilities. The search or filtering functionality can be for the website or application, the current web page or document, or the entire Internet or subsection thereof.")

No support in current engines.

::: support
[FirefoxNo]{.firefox .no}[SafariNo]{.safari .no}[ChromeNo]{.chrome .no}

------------------------------------------------------------------------

[OperaNo]{.opera .no}[EdgeNo]{.edge_blink .no}

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

[Categories](dom.html#concept-element-categories){#the-search-element:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-search-element:flow-content-2}.
:   [Palpable
    content](dom.html#palpable-content-2){#the-search-element:palpable-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-search-element:concept-element-contexts}:
:   Where [flow
    content](dom.html#flow-content-2){#the-search-element:flow-content-2-2}
    is expected.

[Content model](dom.html#concept-element-content-model){#the-search-element:concept-element-content-model}:
:   [Flow
    content](dom.html#flow-content-2){#the-search-element:flow-content-2-3}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-search-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-search-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-search-element:global-attributes}

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-search-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-search).
:   [For implementers](https://w3c.github.io/html-aam/#el-search).

[DOM interface](dom.html#concept-element-dom){#the-search-element:concept-element-dom}:
:   Uses
    [`HTMLElement`{#the-search-element:htmlelement}](dom.html#htmlelement).

The
[`search`{#the-search-element:the-search-element}](#the-search-element)
element
[represents](dom.html#represents){#the-search-element:represents} a part
of a document or application that contains a set of form controls or
other content related to performing a search or filtering operation.
This could be a search of the web site or application; a way of
searching or filtering search results on the current web page; or a
global or Internet-wide search function.

It\'s not appropriate to use the
[`search`{#the-search-element:the-search-element-2}](#the-search-element)
element just for presenting search results, though suggestions and links
as part of \"quick search\" results can be included as part of a search
feature. Rather, a returned web page of search results would instead be
expected to be presented as part of the main content of that web page.

::: example
In the following example, the author is including a search form within
the
[`header`{#the-search-element:the-header-element}](sections.html#the-header-element)
of the web page:

``` html
<header>
  <h1><a href="/">My fancy blog</a></h1>
  ...
  <search>
    <form action="search.php">
      <label for="query">Find an article</label>
      <input id="query" name="q" type="search">
      <button type="submit">Go!</button>
    </form>
  </search>
</header>
```
:::

::: example
In this example, the author has implemented their web application\'s
search functionality entirely with JavaScript. There is no use of the
[`form`{#the-search-element:the-form-element}](forms.html#the-form-element)
element to perform server-side submission, but the containing
[`search`{#the-search-element:the-search-element-3}](#the-search-element)
element semantically identifies the purpose of the descendant content as
representing search capabilities.

``` html
<search>
  <label>
    Find and filter your query
    <input type="search" id="query">
  </label>
  <label>
    <input type="checkbox" id="exact-only">
    Exact matches only
  </label>

  <section>
    <h3>Results found:</h3>
    <ul id="results">
      <li>
        <p><a href="services/consulting">Consulting services</a></p>
        <p>
          Find out how can we help you improve your business with our integrated consultants, Bob and Bob.
        </p>
      </li>
      ...
    </ul>
    <!--
      when a query returns or filters out all results
      render the no results message here
    -->
    <output id="no-results"></output>
  </section>
</search>
```
:::

::: example
In the following example, the page has two search features. The first is
located in the web page\'s
[`header`{#the-search-element:the-header-element-2}](sections.html#the-header-element)
and serves as a global mechanism to search the web site\'s content. Its
purpose is indicated by its specified
[`title`{#the-search-element:the-title-element}](semantics.html#the-title-element)
attribute. The second is included as part of the main content of the
page, as it represents a mechanism to search and filter the content of
the current page. It contains a heading to indicate its purpose.

``` html
<body>
  <header>
    ...
    <search title="Website">
      ...
    </search>
  </header>
  <main>
    <h1>Hotels near your location</h1>
     <search>
       <h2>Filter results</h2>
       ...
     </search>
     <article>
      <!-- search result content -->
    </article>
  </main>
</body>
```
:::

#### [4.4.16]{.secno} The [`div`]{.dfn dfn-type="element"} element[](#the-div-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/div](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/div "The <div> HTML element is the generic container for flow content. It has no effect on the content or layout until styled in some way using CSS (e.g. styling is directly applied to it, or some kind of layout model like Flexbox is applied to its parent element).")

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
[HTMLDivElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLDivElement "The HTMLDivElement interface provides special properties (beyond the regular HTMLElement interface it also has available to it by inheritance) for manipulating <div> elements.")

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

[Categories](dom.html#concept-element-categories){#the-div-element:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-div-element:flow-content-2}.
:   [Palpable
    content](dom.html#palpable-content-2){#the-div-element:palpable-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-div-element:concept-element-contexts}:
:   Where [flow
    content](dom.html#flow-content-2){#the-div-element:flow-content-2-2}
    is expected.
:   As a child of a
    [`dl`{#the-div-element:the-dl-element}](#the-dl-element) element.

[Content model](dom.html#concept-element-content-model){#the-div-element:concept-element-content-model}:
:   If the element is a child of a
    [`dl`{#the-div-element:the-dl-element-2}](#the-dl-element) element:
    one or more [`dt`{#the-div-element:the-dt-element}](#the-dt-element)
    elements followed by one or more
    [`dd`{#the-div-element:the-dd-element}](#the-dd-element) elements,
    optionally intermixed with [script-supporting
    elements](dom.html#script-supporting-elements-2){#the-div-element:script-supporting-elements-2}.
:   If the element is not a child of a
    [`dl`{#the-div-element:the-dl-element-3}](#the-dl-element) element:
    [flow
    content](dom.html#flow-content-2){#the-div-element:flow-content-2-3}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-div-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-div-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-div-element:global-attributes}

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-div-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-div).
:   [For implementers](https://w3c.github.io/html-aam/#el-div).

[DOM interface](dom.html#concept-element-dom){#the-div-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLDivElement : HTMLElement {
      [HTMLConstructor] constructor();

      // also has obsolete members
    };
    ```

The [`div`{#the-div-element:the-div-element}](#the-div-element) element
has no special meaning at all. It
[represents](dom.html#represents){#the-div-element:represents} its
children. It can be used with the
[`class`{#the-div-element:classes}](dom.html#classes),
[`lang`{#the-div-element:attr-lang}](dom.html#attr-lang), and
[`title`{#the-div-element:attr-title}](dom.html#attr-title) attributes
to mark up semantics common to a group of consecutive elements. It can
also be used in a
[`dl`{#the-div-element:the-dl-element-4}](#the-dl-element) element,
wrapping groups of
[`dt`{#the-div-element:the-dt-element-2}](#the-dt-element) and
[`dd`{#the-div-element:the-dd-element-2}](#the-dd-element) elements.

Authors are strongly encouraged to view the
[`div`{#the-div-element:the-div-element-2}](#the-div-element) element as
an element of last resort, for when no other element is suitable. Use of
more appropriate elements instead of the
[`div`{#the-div-element:the-div-element-3}](#the-div-element) element
leads to better accessibility for readers and easier maintainability for
authors.

::: example
For example, a blog post would be marked up using
[`article`{#the-div-element:the-article-element}](sections.html#the-article-element),
a chapter using
[`section`{#the-div-element:the-section-element}](sections.html#the-section-element),
a page\'s navigation aids using
[`nav`{#the-div-element:the-nav-element}](sections.html#the-nav-element),
and a group of form controls using
[`fieldset`{#the-div-element:the-fieldset-element}](form-elements.html#the-fieldset-element).

On the other hand,
[`div`{#the-div-element:the-div-element-4}](#the-div-element) elements
can be useful for stylistic purposes or to wrap multiple paragraphs
within a section that are all to be annotated in a similar way. In the
following example, we see
[`div`{#the-div-element:the-div-element-5}](#the-div-element) elements
used as a way to set the language of two paragraphs at once, instead of
setting the language on the two paragraph elements separately:

``` html
<article lang="en-US">
 <h1>My use of language and my cats</h1>
 <p>My cat's behavior hasn't changed much since her absence, except
 that she plays her new physique to the neighbors regularly, in an
 attempt to get pets.</p>
 <div lang="en-GB">
  <p>My other cat, coloured black and white, is a sweetie. He followed
  us to the pool today, walking down the pavement with us. Yesterday
  he apparently visited our neighbours. I wonder if he recognises that
  their flat is a mirror image of ours.</p>
  <p>Hm, I just noticed that in the last paragraph I used British
  English. But I'm supposed to write in American English. So I
  shouldn't say "pavement" or "flat" or "colour"...</p>
 </div>
 <p>I should say "sidewalk" and "apartment" and "color"!</p>
</article>
```
:::

[← 4.3 Sections](sections.html) --- [Table of Contents](./) --- [4.5
Text-level semantics →](text-level-semantics.html)
