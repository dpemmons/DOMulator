HTML Standard

::: {#head .head .with-buttons}
[![WHATWG](https://resources.whatwg.org/logo.svg){.darkmode-aware
width="100" crossorigin="" height="100"}](https://whatwg.org/){.logo}

# HTML {#html .allcaps}

Living Standard --- Last Updated [2 June 2025]{.pubdate}
:::

[← 4.6 Links](links.html) --- [Table of Contents](./) --- [4.8 Embedded
content →](embedded-content.html)

1.  ::: {#toc-semantics}
    1.  [[4.7]{.secno} Edits](edits.html#edits)
        1.  [[4.7.1]{.secno} The `ins`
            element](edits.html#the-ins-element)
        2.  [[4.7.2]{.secno} The `del`
            element](edits.html#the-del-element)
        3.  [[4.7.3]{.secno} Attributes common to `ins` and `del`
            elements](edits.html#attributes-common-to-ins-and-del-elements)
        4.  [[4.7.4]{.secno} Edits and
            paragraphs](edits.html#edits-and-paragraphs)
        5.  [[4.7.5]{.secno} Edits and
            lists](edits.html#edits-and-lists)
        6.  [[4.7.6]{.secno} Edits and
            tables](edits.html#edits-and-tables)
    :::

### [4.7]{.secno} Edits[](#edits){.self-link}

The [`ins`{#edits:the-ins-element}](#the-ins-element) and
[`del`{#edits:the-del-element}](#the-del-element) elements represent
edits to the document.

#### [4.7.1]{.secno} The [`ins`]{.dfn dfn-type="element"} element[](#the-ins-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/ins](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ins "The <ins> HTML element represents a range of text that has been added to a document. You can use the <del> element to similarly represent a range of text that has been deleted from the document.")

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

[Categories](dom.html#concept-element-categories){#the-ins-element:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-ins-element:flow-content-2}.
:   [Phrasing
    content](dom.html#phrasing-content-2){#the-ins-element:phrasing-content-2}.
:   [Palpable
    content](dom.html#palpable-content-2){#the-ins-element:palpable-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-ins-element:concept-element-contexts}:
:   Where [phrasing
    content](dom.html#phrasing-content-2){#the-ins-element:phrasing-content-2-2}
    is expected.

[Content model](dom.html#concept-element-content-model){#the-ins-element:concept-element-content-model}:
:   [Transparent](dom.html#transparent){#the-ins-element:transparent}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-ins-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-ins-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-ins-element:global-attributes}
:   [`cite`{#the-ins-element:attr-mod-cite}](#attr-mod-cite) --- Link to
    the source of the quotation or more information about the edit
:   [`datetime`{#the-ins-element:attr-mod-datetime}](#attr-mod-datetime)
    --- Date and (optionally) time of the change

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-ins-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-ins).
:   [For implementers](https://w3c.github.io/html-aam/#el-ins).

[DOM interface](dom.html#concept-element-dom){#the-ins-element:concept-element-dom}:
:   Uses
    [`HTMLModElement`{#the-ins-element:htmlmodelement}](#htmlmodelement).

The [`ins`{#the-ins-element:the-ins-element}](#the-ins-element) element
[represents](dom.html#represents){#the-ins-element:represents} an
addition to the document.

::: example
The following represents the addition of a single paragraph:

``` html
<aside>
 <ins>
  <p> I like fruit. </p>
 </ins>
</aside>
```

As does the following, because everything in the
[`aside`{#the-ins-element:the-aside-element}](sections.html#the-aside-element)
element here counts as [phrasing
content](dom.html#phrasing-content-2){#the-ins-element:phrasing-content-2-3}
and therefore there is just one
[paragraph](dom.html#paragraph){#the-ins-element:paragraph}:

``` html
<aside>
 <ins>
  Apples are <em>tasty</em>.
 </ins>
 <ins>
  So are pears.
 </ins>
</aside>
```
:::

[`ins`{#the-ins-element:the-ins-element-2}](#the-ins-element) elements
should not cross [implied
paragraph](dom.html#paragraph){#the-ins-element:paragraph-2} boundaries.

::: example
The following example represents the addition of two paragraphs, the
second of which was inserted in two parts. The first
[`ins`{#the-ins-element:the-ins-element-3}](#the-ins-element) element in
this example thus crosses a paragraph boundary, which is considered poor
form.

``` bad
<aside>
 <!-- don't do this -->
 <ins datetime="2005-03-16 00:00Z">
  <p> I like fruit. </p>
  Apples are <em>tasty</em>.
 </ins>
 <ins datetime="2007-12-19 00:00Z">
  So are pears.
 </ins>
</aside>
```

Here is a better way of marking this up. It uses more elements, but none
of the elements cross implied paragraph boundaries.

``` html
<aside>
 <ins datetime="2005-03-16 00:00Z">
  <p> I like fruit. </p>
 </ins>
 <ins datetime="2005-03-16 00:00Z">
  Apples are <em>tasty</em>.
 </ins>
 <ins datetime="2007-12-19 00:00Z">
  So are pears.
 </ins>
</aside>
```
:::

#### [4.7.2]{.secno} The [`del`]{.dfn dfn-type="element"} element[](#the-del-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/del](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/del "The <del> HTML element represents a range of text that has been deleted from a document. This can be used when rendering "track changes" or source code diff information, for example. The <ins> element can be used for the opposite purpose: to indicate text that has been added to the document.")

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

[Categories](dom.html#concept-element-categories){#the-del-element:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-del-element:flow-content-2}.
:   [Phrasing
    content](dom.html#phrasing-content-2){#the-del-element:phrasing-content-2}.
:   [Palpable
    content](dom.html#palpable-content-2){#the-del-element:palpable-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-del-element:concept-element-contexts}:
:   Where [phrasing
    content](dom.html#phrasing-content-2){#the-del-element:phrasing-content-2-2}
    is expected.

[Content model](dom.html#concept-element-content-model){#the-del-element:concept-element-content-model}:
:   [Transparent](dom.html#transparent){#the-del-element:transparent}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-del-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-del-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-del-element:global-attributes}
:   [`cite`{#the-del-element:attr-mod-cite}](#attr-mod-cite) --- Link to
    the source of the quotation or more information about the edit
:   [`datetime`{#the-del-element:attr-mod-datetime}](#attr-mod-datetime)
    --- Date and (optionally) time of the change

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-del-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-del).
:   [For implementers](https://w3c.github.io/html-aam/#el-del).

[DOM interface](dom.html#concept-element-dom){#the-del-element:concept-element-dom}:
:   Uses
    [`HTMLModElement`{#the-del-element:htmlmodelement}](#htmlmodelement).

The [`del`{#the-del-element:the-del-element}](#the-del-element) element
[represents](dom.html#represents){#the-del-element:represents} a removal
from the document.

[`del`{#the-del-element:the-del-element-2}](#the-del-element) elements
should not cross [implied
paragraph](dom.html#paragraph){#the-del-element:paragraph} boundaries.

::: example
The following shows a \"to do\" list where items that have been done are
crossed-off with the date and time of their completion.

``` html
<h1>To Do</h1>
<ul>
 <li>Empty the dishwasher</li>
 <li><del datetime="2009-10-11T01:25-07:00">Watch Walter Lewin's lectures</del></li>
 <li><del datetime="2009-10-10T23:38-07:00">Download more tracks</del></li>
 <li>Buy a printer</li>
</ul>
```
:::

#### [4.7.3]{.secno} Attributes common to [`ins`{#attributes-common-to-ins-and-del-elements:the-ins-element}](#the-ins-element) and [`del`{#attributes-common-to-ins-and-del-elements:the-del-element}](#the-del-element) elements[](#attributes-common-to-ins-and-del-elements){.self-link}

The [`cite`]{#attr-mod-cite .dfn dfn-for="ins,del"
dfn-type="element-attr"} attribute may be used to specify the
[URL](https://dom.spec.whatwg.org/#concept-document-url){#attributes-common-to-ins-and-del-elements:the-document's-address
x-internal="the-document's-address"} of a document that explains the
change. When that document is long, for instance the minutes of a
meeting, authors are encouraged to include a
[fragment](https://url.spec.whatwg.org/#concept-url-fragment){#attributes-common-to-ins-and-del-elements:concept-url-fragment
x-internal="concept-url-fragment"} pointing to the specific part of that
document that discusses the change.

If the
[`cite`{#attributes-common-to-ins-and-del-elements:attr-mod-cite}](#attr-mod-cite)
attribute is present, it must be a [valid URL potentially surrounded by
spaces](urls-and-fetching.html#valid-url-potentially-surrounded-by-spaces){#attributes-common-to-ins-and-del-elements:valid-url-potentially-surrounded-by-spaces}
that explains the change. To obtain the corresponding citation link, the
value of the attribute must be
[parsed](urls-and-fetching.html#encoding-parsing-a-url){#attributes-common-to-ins-and-del-elements:encoding-parsing-a-url}
relative to the element\'s [node
document](https://dom.spec.whatwg.org/#concept-node-document){#attributes-common-to-ins-and-del-elements:node-document
x-internal="node-document"}. User agents may allow users to follow such
citation links, but they are primarily intended for private use (e.g.,
by server-side scripts collecting statistics about a site\'s edits), not
for readers.

The [`datetime`]{#attr-mod-datetime .dfn dfn-for="ins,del"
dfn-type="element-attr"} attribute may be used to specify the time and
date of the change.

If present, the
[`datetime`{#attributes-common-to-ins-and-del-elements:attr-mod-datetime}](#attr-mod-datetime)
attribute\'s value must be a [valid date string with optional
time](common-microsyntaxes.html#valid-date-string-with-optional-time){#attributes-common-to-ins-and-del-elements:valid-date-string-with-optional-time}.

User agents must parse the
[`datetime`{#attributes-common-to-ins-and-del-elements:attr-mod-datetime-2}](#attr-mod-datetime)
attribute according to the [parse a date or time
string](common-microsyntaxes.html#parse-a-date-or-time-string){#attributes-common-to-ins-and-del-elements:parse-a-date-or-time-string}
algorithm. If that doesn\'t return a
[date](common-microsyntaxes.html#concept-date){#attributes-common-to-ins-and-del-elements:concept-date}
or a [global date and
time](common-microsyntaxes.html#concept-datetime){#attributes-common-to-ins-and-del-elements:concept-datetime},
then the modification has no associated timestamp (the value is
non-conforming; it is not a [valid date string with optional
time](common-microsyntaxes.html#valid-date-string-with-optional-time){#attributes-common-to-ins-and-del-elements:valid-date-string-with-optional-time-2}).
Otherwise, the modification is marked as having been made at the given
[date](common-microsyntaxes.html#concept-date){#attributes-common-to-ins-and-del-elements:concept-date-2}
or [global date and
time](common-microsyntaxes.html#concept-datetime){#attributes-common-to-ins-and-del-elements:concept-datetime-2}.
If the given value is a [global date and
time](common-microsyntaxes.html#concept-datetime){#attributes-common-to-ins-and-del-elements:concept-datetime-3}
then user agents should use the associated time-zone offset information
to determine which time zone to present the given datetime in.

This value may be shown to the user, but it is primarily intended for
private use.

The
[`ins`{#attributes-common-to-ins-and-del-elements:the-ins-element-2}](#the-ins-element)
and
[`del`{#attributes-common-to-ins-and-del-elements:the-del-element-2}](#the-del-element)
elements must implement the
[`HTMLModElement`{#attributes-common-to-ins-and-del-elements:htmlmodelement}](#htmlmodelement)
interface:

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLModElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLModElement "The HTMLModElement interface provides special properties (beyond the regular methods and properties available through the HTMLElement interface they also have available to them by inheritance) for manipulating modification elements, that is <del> and <ins>.")

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

``` idl
[Exposed=Window]
interface HTMLModElement : HTMLElement {
  [HTMLConstructor] constructor();

  [CEReactions] attribute USVString cite;
  [CEReactions] attribute DOMString dateTime;
};
```

The [`cite`]{#dom-mod-cite .dfn dfn-for="HTMLModElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#attributes-common-to-ins-and-del-elements:reflect}
the element\'s
[`cite`{#attributes-common-to-ins-and-del-elements:attr-mod-cite-2}](#attr-mod-cite)
content attribute. The [`dateTime`]{#dom-mod-datetime .dfn
dfn-for="HTMLModElement" dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#attributes-common-to-ins-and-del-elements:reflect-2}
the element\'s
[`datetime`{#attributes-common-to-ins-and-del-elements:attr-mod-datetime-3}](#attr-mod-datetime)
content attribute.

#### [4.7.4]{.secno} Edits and paragraphs[](#edits-and-paragraphs){.self-link}

*This section is non-normative.*

Since the
[`ins`{#edits-and-paragraphs:the-ins-element}](#the-ins-element) and
[`del`{#edits-and-paragraphs:the-del-element}](#the-del-element)
elements do not affect
[paragraphing](dom.html#paragraph){#edits-and-paragraphs:paragraph}, it
is possible, in some cases where paragraphs are
[implied](dom.html#paragraph){#edits-and-paragraphs:paragraph-2}
(without explicit
[`p`{#edits-and-paragraphs:the-p-element}](grouping-content.html#the-p-element)
elements), for an
[`ins`{#edits-and-paragraphs:the-ins-element-2}](#the-ins-element) or
[`del`{#edits-and-paragraphs:the-del-element-2}](#the-del-element)
element to span both an entire paragraph or other non-[phrasing
content](dom.html#phrasing-content-2){#edits-and-paragraphs:phrasing-content-2}
elements and part of another paragraph. For example:

``` html
<section>
 <ins>
  <p>
   This is a paragraph that was inserted.
  </p>
  This is another paragraph whose first sentence was inserted
  at the same time as the paragraph above.
 </ins>
 This is a second sentence, which was there all along.
</section>
```

By only wrapping some paragraphs in
[`p`{#edits-and-paragraphs:the-p-element-2}](grouping-content.html#the-p-element)
elements, one can even get the end of one paragraph, a whole second
paragraph, and the start of a third paragraph to be covered by the same
[`ins`{#edits-and-paragraphs:the-ins-element-3}](#the-ins-element) or
[`del`{#edits-and-paragraphs:the-del-element-3}](#the-del-element)
element (though this is very confusing, and not considered good
practice):

``` bad
<section>
 This is the first paragraph. <ins>This sentence was
 inserted.
 <p>This second paragraph was inserted.</p>
 This sentence was inserted too.</ins> This is the
 third paragraph in this example.
 <!-- (don't do this) -->
</section>
```

However, due to the way [implied
paragraphs](dom.html#paragraph){#edits-and-paragraphs:paragraph-3} are
defined, it is not possible to mark up the end of one paragraph and the
start of the very next one using the same
[`ins`{#edits-and-paragraphs:the-ins-element-4}](#the-ins-element) or
[`del`{#edits-and-paragraphs:the-del-element-4}](#the-del-element)
element. You instead have to use one (or two)
[`p`{#edits-and-paragraphs:the-p-element-3}](grouping-content.html#the-p-element)
element(s) and two
[`ins`{#edits-and-paragraphs:the-ins-element-5}](#the-ins-element) or
[`del`{#edits-and-paragraphs:the-del-element-5}](#the-del-element)
elements, as for example:

``` html
<section>
 <p>This is the first paragraph. <del>This sentence was
 deleted.</del></p>
 <p><del>This sentence was deleted too.</del> That
 sentence needed a separate &lt;del&gt; element.</p>
</section>
```

Partly because of the confusion described above, authors are strongly
encouraged to always mark up all paragraphs with the
[`p`{#edits-and-paragraphs:the-p-element-4}](grouping-content.html#the-p-element)
element, instead of having
[`ins`{#edits-and-paragraphs:the-ins-element-6}](#the-ins-element) or
[`del`{#edits-and-paragraphs:the-del-element-6}](#the-del-element)
elements that cross [implied
paragraphs](dom.html#paragraph){#edits-and-paragraphs:paragraph-4}
boundaries.

#### [4.7.5]{.secno} Edits and lists[](#edits-and-lists){.self-link}

*This section is non-normative.*

The content models of the
[`ol`{#edits-and-lists:the-ol-element}](grouping-content.html#the-ol-element)
and
[`ul`{#edits-and-lists:the-ul-element}](grouping-content.html#the-ul-element)
elements do not allow
[`ins`{#edits-and-lists:the-ins-element}](#the-ins-element) and
[`del`{#edits-and-lists:the-del-element}](#the-del-element) elements as
children. Lists always represent all their items, including items that
would otherwise have been marked as deleted.

To indicate that an item is inserted or deleted, an
[`ins`{#edits-and-lists:the-ins-element-2}](#the-ins-element) or
[`del`{#edits-and-lists:the-del-element-2}](#the-del-element) element
can be wrapped around the contents of the
[`li`{#edits-and-lists:the-li-element}](grouping-content.html#the-li-element)
element. To indicate that an item has been replaced by another, a single
[`li`{#edits-and-lists:the-li-element-2}](grouping-content.html#the-li-element)
element can have one or more
[`del`{#edits-and-lists:the-del-element-3}](#the-del-element) elements
followed by one or more
[`ins`{#edits-and-lists:the-ins-element-3}](#the-ins-element) elements.

::: example
In the following example, a list that started empty had items added and
removed from it over time. The bits in the example that have been
emphasized show the parts that are the \"current\" state of the list.
The list item numbers don\'t take into account the edits, though.

``` html
<h1>Stop-ship bugs</h1>
<ol>
 <li><ins datetime="2008-02-12T15:20Z">Bug 225:
 Rain detector doesn't work in snow</ins></li>
 <li><del datetime="2008-03-01T20:22Z"><ins datetime="2008-02-14T12:02Z">Bug 228:
 Water buffer overflows in April</ins></del></li>
 <li><ins datetime="2008-02-16T13:50Z">Bug 230:
 Water heater doesn't use renewable fuels</ins></li>
 <li><del datetime="2008-02-20T21:15Z"><ins datetime="2008-02-16T14:25Z">Bug 232:
 Carbon dioxide emissions detected after startup</ins></del></li>
</ol>
```
:::

::: example
In the following example, a list that started with just fruit was
replaced by a list with just colors.

``` html
<h1>List of <del>fruits</del><ins>colors</ins></h1>
<ul>
 <li><del>Lime</del><ins>Green</ins></li>
 <li><del>Apple</del></li>
 <li>Orange</li>
 <li><del>Pear</del></li>
 <li><ins>Teal</ins></li>
 <li><del>Lemon</del><ins>Yellow</ins></li>
 <li>Olive</li>
 <li><ins>Purple</ins></li>
</ul>
```
:::

#### [4.7.6]{.secno} Edits and tables[](#edits-and-tables){.self-link}

*This section is non-normative.*

The elements that form part of the table model have complicated content
model requirements that do not allow for the
[`ins`{#edits-and-tables:the-ins-element}](#the-ins-element) and
[`del`{#edits-and-tables:the-del-element}](#the-del-element) elements,
so indicating edits to a table can be difficult.

To indicate that an entire row or an entire column has been added or
removed, the entire contents of each cell in that row or column can be
wrapped in
[`ins`{#edits-and-tables:the-ins-element-2}](#the-ins-element) or
[`del`{#edits-and-tables:the-del-element-2}](#the-del-element) elements
(respectively).

::: example
Here, a table\'s row has been added:

``` html
<table>
 <thead>
  <tr> <th> Game name           <th> Game publisher   <th> Verdict
 <tbody>
  <tr> <td> Diablo 2            <td> Blizzard         <td> 8/10
  <tr> <td> Portal              <td> Valve            <td> 10/10
  <tr> <td> <ins>Portal 2</ins> <td> <ins>Valve</ins> <td> <ins>10/10</ins>
</table>
```

Here, a column has been removed (the time at which it was removed is
given also, as is a link to the page explaining why):

``` html
<table>
 <thead>
  <tr> <th> Game name           <th> Game publisher   <th> <del cite="/edits/r192" datetime="2011-05-02 14:23Z">Verdict</del>
 <tbody>
  <tr> <td> Diablo 2            <td> Blizzard         <td> <del cite="/edits/r192" datetime="2011-05-02 14:23Z">8/10</del>
  <tr> <td> Portal              <td> Valve            <td> <del cite="/edits/r192" datetime="2011-05-02 14:23Z">10/10</del>
  <tr> <td> Portal 2            <td> Valve            <td> <del cite="/edits/r192" datetime="2011-05-02 14:23Z">10/10</del>
</table>
```
:::

Generally speaking, there is no good way to indicate more complicated
edits (e.g. that a cell was removed, moving all subsequent cells up or
to the left).

[← 4.6 Links](links.html) --- [Table of Contents](./) --- [4.8 Embedded
content →](embedded-content.html)
