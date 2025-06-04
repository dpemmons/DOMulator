HTML Standard

::: {#head .head .with-buttons}
[![WHATWG](https://resources.whatwg.org/logo.svg){.darkmode-aware
width="100" crossorigin="" height="100"}](https://whatwg.org/){.logo}

# HTML {#html .allcaps}

Living Standard --- Last Updated [2 June 2025]{.pubdate}
:::

[← 4.13 Custom elements](custom-elements.html) --- [Table of
Contents](./) --- [5 Microdata →](microdata.html)

1.  ::: {#toc-semantics}
    1.  [[4.14]{.secno} Common idioms without dedicated
        elements](semantics-other.html#common-idioms)
        1.  [[4.14.1]{.secno} Breadcrumb
            navigation](semantics-other.html#rel-up)
        2.  [[4.14.2]{.secno} Tag
            clouds](semantics-other.html#tag-clouds)
        3.  [[4.14.3]{.secno}
            Conversations](semantics-other.html#conversations)
        4.  [[4.14.4]{.secno} Footnotes](semantics-other.html#footnotes)
    2.  [[4.15]{.secno} Disabled
        elements](semantics-other.html#disabled-elements)
    3.  [[4.16]{.secno} Matching HTML elements using selectors and
        CSS](semantics-other.html#selectors)
        1.  [[4.16.1]{.secno} Case-sensitivity of the CSS \'attr()\'
            function](semantics-other.html#case-sensitivity-of-the-css-'attr()'-function)
        2.  [[4.16.2]{.secno} Case-sensitivity of
            selectors](semantics-other.html#case-sensitivity-of-selectors)
        3.  [[4.16.3]{.secno}
            Pseudo-classes](semantics-other.html#pseudo-classes)
    :::

### [4.14]{.secno} Common idioms without dedicated elements[](#common-idioms){.self-link} {#common-idioms}

#### [4.14.1]{.secno} Breadcrumb navigation[](#rel-up){.self-link} {#rel-up}

This specification does not provide a machine-readable way of describing
breadcrumb navigation menus. Authors are encouraged to just use a series
of links in a paragraph. The
[`nav`{#rel-up:the-nav-element}](sections.html#the-nav-element) element
can be used to mark the section containing these paragraphs as being
navigation blocks.

::: example
In the following example, the current page can be reached via two paths.

``` html
<nav>
 <p>
  <a href="/">Main</a> ▸
  <a href="/products/">Products</a> ▸
  <a href="/products/dishwashers/">Dishwashers</a> ▸
  <a>Second hand</a>
 </p>
 <p>
  <a href="/">Main</a> ▸
  <a href="/second-hand/">Second hand</a> ▸
  <a>Dishwashers</a>
 </p>
</nav>
```
:::

#### [4.14.2]{.secno} Tag clouds[](#tag-clouds){.self-link}

This specification does not define any markup specifically for marking
up lists of keywords that apply to a group of pages (also known as *tag
clouds*). In general, authors are encouraged to either mark up such
lists using
[`ul`{#tag-clouds:the-ul-element}](grouping-content.html#the-ul-element)
elements with explicit inline counts that are then hidden and turned
into a presentational effect using a style sheet, or to use SVG.

::: example
Here, three tags are included in a short tag cloud:

``` html
<style>
.tag-cloud > li > span { display: none; }
.tag-cloud > li { display: inline; }
.tag-cloud-1 { font-size: 0.7em; }
.tag-cloud-2 { font-size: 0.9em; }
.tag-cloud-3 { font-size: 1.1em; }
.tag-cloud-4 { font-size: 1.3em; }
.tag-cloud-5 { font-size: 1.5em; }

@media speech {
  .tag-cloud > li > span { display:inline }
}
</style>
...
<ul class="tag-cloud">
 <li class="tag-cloud-4"><a title="28 instances" href="/t/apple">apple</a> <span>(popular)</span>
 <li class="tag-cloud-2"><a title="6 instances"  href="/t/kiwi">kiwi</a> <span>(rare)</span>
 <li class="tag-cloud-5"><a title="41 instances" href="/t/pear">pear</a> <span>(very popular)</span>
</ul>
```

The actual frequency of each tag is given using the
[`title`{#tag-clouds:attr-title}](dom.html#attr-title) attribute. A CSS
style sheet is provided to convert the markup into a cloud of
differently-sized words, but for user agents that do not support CSS or
are not visual, the markup contains annotations like \"(popular)\" or
\"(rare)\" to categorize the various tags by frequency, thus enabling
all users to benefit from the information.

The
[`ul`{#tag-clouds:the-ul-element-2}](grouping-content.html#the-ul-element)
element is used (rather than
[`ol`{#tag-clouds:the-ol-element}](grouping-content.html#the-ol-element))
because the order is not particularly important: while the list is in
fact ordered alphabetically, it would convey the same information if
ordered by, say, the length of the tag.

The [`tag`{#tag-clouds:link-type-tag}](links.html#link-type-tag)
[`rel`{#tag-clouds:attr-hyperlink-rel}](links.html#attr-hyperlink-rel)-keyword
is *not* used on these
[`a`{#tag-clouds:the-a-element}](text-level-semantics.html#the-a-element)
elements because they do not represent tags that apply to the page
itself; they are just part of an index listing the tags themselves.
:::

#### [4.14.3]{.secno} Conversations[](#conversations){.self-link}

This specification does not define a specific element for marking up
conversations, meeting minutes, chat transcripts, dialogues in
screenplays, instant message logs, and other situations where different
players take turns in discourse.

Instead, authors are encouraged to mark up conversations using
[`p`{#conversations:the-p-element}](grouping-content.html#the-p-element)
elements and punctuation. Authors who need to mark the speaker for
styling purposes are encouraged to use
[`span`{#conversations:the-span-element}](text-level-semantics.html#the-span-element)
or
[`b`{#conversations:the-b-element}](text-level-semantics.html#the-b-element).
Paragraphs with their text wrapped in the
[`i`{#conversations:the-i-element}](text-level-semantics.html#the-i-element)
element can be used for marking up stage directions.

::: example
This example demonstrates this using an extract from Abbot and
Costello\'s famous sketch, Who\'s on first:

``` html
<p> Costello: Look, you gotta first baseman?
<p> Abbott: Certainly.
<p> Costello: Who's playing first?
<p> Abbott: That's right.
<p> Costello becomes exasperated.
<p> Costello: When you pay off the first baseman every month, who gets the money?
<p> Abbott: Every dollar of it.
```
:::

::: example
The following extract shows how an IM conversation log could be marked
up, using the
[`data`{#conversations:the-data-element}](text-level-semantics.html#the-data-element)
element to provide Unix timestamps for each line. Note that the
timestamps are provided in a format that the
[`time`{#conversations:the-time-element}](text-level-semantics.html#the-time-element)
element does not support, so the
[`data`{#conversations:the-data-element-2}](text-level-semantics.html#the-data-element)
element is used instead (namely, Unix `time_t` timestamps). Had the
author wished to mark up the data using one of the date and time formats
supported by the
[`time`{#conversations:the-time-element-2}](text-level-semantics.html#the-time-element)
element, that element could have been used instead of
[`data`{#conversations:the-data-element-3}](text-level-semantics.html#the-data-element).
This could be advantageous as it would allow data analysis tools to
detect the timestamps unambiguously, without coordination with the page
author.

``` html
<p> <data value="1319898155">14:22</data> <b>egof</b> I'm not that nerdy, I've only seen 30% of the star trek episodes
<p> <data value="1319898192">14:23</data> <b>kaj</b> if you know what percentage of the star trek episodes you have seen, you are inarguably nerdy
<p> <data value="1319898200">14:23</data> <b>egof</b> it's unarguably
<p> <data value="1319898228">14:23</data> <i>* kaj blinks</i>
<p> <data value="1319898260">14:24</data> <b>kaj</b> you are not helping your case
```
:::

::: example
HTML does not have a good way to mark up graphs, so descriptions of
interactive conversations from games are more difficult to mark up. This
example shows one possible convention using
[`dl`{#conversations:the-dl-element}](grouping-content.html#the-dl-element)
elements to list the possible responses at each point in the
conversation. Another option to consider is describing the conversation
in the form of a DOT file, and outputting the result as an SVG image to
place in the document. [\[DOT\]](references.html#refsDOT)

``` html
<p> Next, you meet a fisher. You can say one of several greetings:
<dl>
 <dt> "Hello there!"
 <dd>
  <p> She responds with "Hello, how may I help you?"; you can respond with:
  <dl>
   <dt> "I would like to buy a fish."
   <dd> <p> She sells you a fish and the conversation finishes.
   <dt> "Can I borrow your boat?"
   <dd>
    <p> She is surprised and asks "What are you offering in return?".
    <dl>
     <dt> "Five gold." (if you have enough)
     <dt> "Ten gold." (if you have enough)
     <dt> "Fifteen gold." (if you have enough)
     <dd> <p> She lends you her boat. The conversation ends.
     <dt> "A fish." (if you have one)
     <dt> "A newspaper." (if you have one)
     <dt> "A pebble." (if you have one)
     <dd> <p> "No thanks", she replies. Your conversation options
     at this point are the same as they were after asking to borrow
     her boat, minus any options you've suggested before.
    </dl>
   </dd>
  </dl>
 </dd>
 <dt> "Vote for me in the next election!"
 <dd> <p> She turns away. The conversation finishes.
 <dt> "Madam, are you aware that your fish are running away?"
 <dd>
  <p> She looks at you skeptically and says "Fish cannot run, miss".
  <dl>
   <dt> "You got me!"
   <dd> <p> The fisher sighs and the conversation ends.
   <dt> "Only kidding."
   <dd> <p> "Good one!" she retorts. Your conversation options at this
   point are the same as those following "Hello there!" above.
   <dt> "Oh, then what are they doing?"
   <dd> <p> She looks at her fish, giving you an opportunity to steal
   her boat, which you do. The conversation ends.
  </dl>
 </dd>
</dl>
```
:::

::: example
In some games, conversations are simpler: each character merely has a
fixed set of lines that they say. In this example, a game
FAQ/walkthrough lists some of the known possible responses for each
character:

``` html
<section>
 <h1>Dialogue</h1>
 <p><small>Some characters repeat their lines in order each time you interact
 with them, others randomly pick from amongst their lines. Those who respond in
 order have numbered entries in the lists below.</small>
 <h2>The Shopkeeper</h2>
 <ul>
  <li>How may I help you?
  <li>Fresh apples!
  <li>A loaf of bread for madam?
 </ul>
 <h2>The pilot</h2>
 <p>Before the accident:
 <ul>
  <li>I'm about to fly out, sorry!
  <li>Sorry, I'm just waiting for flight clearance and then I'll be off!
 </ul>
 <p>After the accident:
 <ol>
  <li>I'm about to fly out, sorry!
  <li>Ok, I'm not leaving right now, my plane is being cleaned.
  <li>Ok, it's not being cleaned, it needs a minor repair first.
  <li>Ok, ok, stop bothering me! Truth is, I had a crash.
 </ol>
 <h2>Clan Leader</h2>
 <p>During the first clan meeting:
 <ul>
  <li>Hey, have you seen my daughter? I bet she's up to something nefarious again...
  <li>Nice weather we're having today, eh?
  <li>The name is Bailey, Jeff Bailey. How can I help you today?
  <li>A glass of water? Fresh from the well!
 </ul>
 <p>After the earthquake:
 <ol>
  <li>Everyone is safe in the shelter, we just have to put out the fire!
  <li>I'll go and tell the fire brigade, you keep hosing it down!
 </ol>
</section>
```
:::

#### [4.14.4]{.secno} Footnotes[](#footnotes){.self-link}

HTML does not have a dedicated mechanism for marking up footnotes. Here
are the suggested alternatives.

------------------------------------------------------------------------

For short inline annotations, the
[`title`{#footnotes:attr-title}](dom.html#attr-title) attribute could be
used.

::: example
In this example, two parts of a dialogue are annotated with
footnote-like content using the
[`title`{#footnotes:attr-title-2}](dom.html#attr-title) attribute.

``` html
<p> <b>Customer</b>: Hello! I wish to register a complaint. Hello. Miss?
<p> <b>Shopkeeper</b>: <span title="Colloquial pronunciation of 'What do you'"
>Watcha</span> mean, miss?
<p> <b>Customer</b>: Uh, I'm sorry, I have a cold. I wish to make a complaint.
<p> <b>Shopkeeper</b>: Sorry, <span title="This is, of course, a lie.">we're
closing for lunch</span>.
```
:::

Unfortunately, relying on the
[`title`{#footnotes:attr-title-3}](dom.html#attr-title) attribute is
currently discouraged as many user agents do not expose the attribute in
an accessible manner as required by this specification (e.g. requiring a
pointing device such as a mouse to cause a tooltip to appear, which
excludes keyboard-only users and touch-only users, such as anyone with a
modern phone or tablet).

If the [`title`{#footnotes:attr-title-4}](dom.html#attr-title) attribute
is used, CSS can be used to draw the reader\'s attention to the elements
with the attribute.

::: example
For example, the following CSS places a dashed line below elements that
have a [`title`{#footnotes:attr-title-5}](dom.html#attr-title)
attribute.

``` css
[title] { border-bottom: thin dashed; }
```
:::

------------------------------------------------------------------------

For longer annotations, the
[`a`{#footnotes:the-a-element}](text-level-semantics.html#the-a-element)
element should be used, pointing to an element later in the document.
The convention is that the contents of the link be a number in square
brackets.

::: example
In this example, a footnote in the dialogue links to a paragraph below
the dialogue. The paragraph then reciprocally links back to the
dialogue, allowing the user to return to the location of the footnote.

``` html
<p> Announcer: Number 16: The <i>hand</i>.
<p> Interviewer: Good evening. I have with me in the studio tonight
Mr Norman St John Polevaulter, who for the past few years has been
contradicting people. Mr Polevaulter, why <em>do</em> you
contradict people?
<p> Norman: I don't. <sup><a href="#fn1" id="r1">[1]</a></sup>
<p> Interviewer: You told me you did!
...
<section>
 <p id="fn1"><a href="#r1">[1]</a> This is, naturally, a lie,
 but paradoxically if it were true he could not say so without
 contradicting the interviewer and thus making it false.</p>
</section>
```
:::

------------------------------------------------------------------------

For side notes, longer annotations that apply to entire sections of the
text rather than just specific words or sentences, the
[`aside`{#footnotes:the-aside-element}](sections.html#the-aside-element)
element should be used.

::: example
In this example, a sidebar is given after a dialogue, giving it some
context.

``` html
<p> <span class="speaker">Customer</span>: I will not buy this record, it is scratched.
<p> <span class="speaker">Shopkeeper</span>: I'm sorry?
<p> <span class="speaker">Customer</span>: I will not buy this record, it is scratched.
<p> <span class="speaker">Shopkeeper</span>: No no no, this's'a tobacconist's.
<aside>
 <p>In 1970, the British Empire lay in ruins, and foreign
 nationalists frequented the streets — many of them Hungarians
 (not the streets — the foreign nationals). Sadly, Alexander
 Yalt has been publishing incompetently-written phrase books.
</aside>
```
:::

------------------------------------------------------------------------

For figures or tables, footnotes can be included in the relevant
[`figcaption`{#footnotes:the-figcaption-element}](grouping-content.html#the-figcaption-element)
or
[`caption`{#footnotes:the-caption-element}](tables.html#the-caption-element)
element, or in surrounding prose.

::: example
In this example, a table has cells with footnotes that are given in
prose. A
[`figure`{#footnotes:the-figure-element}](grouping-content.html#the-figure-element)
element is used to give a single legend to the combination of the table
and its footnotes.

``` html
<figure>
 <figcaption>Table 1. Alternative activities for knights.</figcaption>
 <table>
  <tr>
   <th> Activity
   <th> Location
   <th> Cost
  <tr>
   <td> Dance
   <td> Wherever possible
   <td> £0<sup><a href="#fn1">1</a></sup>
  <tr>
   <td> Routines, chorus scenes<sup><a href="#fn2">2</a></sup>
   <td> Undisclosed
   <td> Undisclosed
  <tr>
   <td> Dining<sup><a href="#fn3">3</a></sup>
   <td> Camelot
   <td> Cost of ham, jam, and spam<sup><a href="#fn4">4</a></sup>
 </table>
 <p id="fn1">1. Assumed.</p>
 <p id="fn2">2. Footwork impeccable.</p>
 <p id="fn3">3. Quality described as "well".</p>
 <p id="fn4">4. A lot.</p>
</figure>
```
:::

### [4.15]{.secno} Disabled elements[](#disabled-elements){.self-link}

An element is said to be [actually disabled]{#concept-element-disabled
.dfn} if it is one of the following:

- a
  [`button`{#disabled-elements:the-button-element}](form-elements.html#the-button-element)
  element that is
  [disabled](form-control-infrastructure.html#concept-fe-disabled){#disabled-elements:concept-fe-disabled}
- an
  [`input`{#disabled-elements:the-input-element}](input.html#the-input-element)
  element that is
  [disabled](form-control-infrastructure.html#concept-fe-disabled){#disabled-elements:concept-fe-disabled-2}
- a
  [`select`{#disabled-elements:the-select-element}](form-elements.html#the-select-element)
  element that is
  [disabled](form-control-infrastructure.html#concept-fe-disabled){#disabled-elements:concept-fe-disabled-3}
- a
  [`textarea`{#disabled-elements:the-textarea-element}](form-elements.html#the-textarea-element)
  element that is
  [disabled](form-control-infrastructure.html#concept-fe-disabled){#disabled-elements:concept-fe-disabled-4}
- an
  [`optgroup`{#disabled-elements:the-optgroup-element}](form-elements.html#the-optgroup-element)
  element that has a
  [`disabled`{#disabled-elements:attr-optgroup-disabled}](form-elements.html#attr-optgroup-disabled)
  attribute
- an
  [`option`{#disabled-elements:the-option-element}](form-elements.html#the-option-element)
  element that is
  [disabled](form-elements.html#concept-option-disabled){#disabled-elements:concept-option-disabled}
- a
  [`fieldset`{#disabled-elements:the-fieldset-element}](form-elements.html#the-fieldset-element)
  element that is a [disabled
  fieldset](form-elements.html#concept-fieldset-disabled){#disabled-elements:concept-fieldset-disabled}
- a [form-associated custom
  element](custom-elements.html#form-associated-custom-element){#disabled-elements:form-associated-custom-element}
  that is
  [disabled](form-control-infrastructure.html#concept-fe-disabled){#disabled-elements:concept-fe-disabled-5}

This definition is used to determine what elements are
[focusable](interaction.html#focusable){#disabled-elements:focusable}
and which elements match the
[`:enabled`{#disabled-elements:selector-enabled}](#selector-enabled) and
[`:disabled`{#disabled-elements:selector-disabled}](#selector-disabled)
[pseudo
classes](https://drafts.csswg.org/selectors/#pseudo-class){#disabled-elements:pseudo-class
x-internal="pseudo-class"}.

### [4.16]{.secno} Matching HTML elements using selectors and CSS[](#selectors){.self-link} {#selectors}

#### [4.16.1]{.secno} Case-sensitivity of the CSS [\'attr()\'](https://drafts.csswg.org/css-values/#funcdef-attr){#case-sensitivity-of-the-css-'attr()'-function:'attr()' x-internal="'attr()'"} function[](#case-sensitivity-of-the-css-'attr()'-function){.self-link} {#case-sensitivity-of-the-css-'attr()'-function}

CSS Values and Units leaves the case-sensitivity of attribute names for
the purpose of the
[\'attr()\'](https://drafts.csswg.org/css-values/#funcdef-attr){#case-sensitivity-of-the-css-'attr()'-function:'attr()'-2
x-internal="'attr()'"} function to be defined by the host language.
[\[CSSVALUES\]](references.html#refsCSSVALUES)

When comparing the attribute name part of a CSS
[\'attr()\'](https://drafts.csswg.org/css-values/#funcdef-attr){#case-sensitivity-of-the-css-'attr()'-function:'attr()'-3
x-internal="'attr()'"} function to the names of namespace-less
attributes on [HTML
elements](infrastructure.html#html-elements){#case-sensitivity-of-the-css-'attr()'-function:html-elements}
in [HTML
documents](https://dom.spec.whatwg.org/#html-document){#case-sensitivity-of-the-css-'attr()'-function:html-documents
x-internal="html-documents"}, the name part of the CSS
[\'attr()\'](https://drafts.csswg.org/css-values/#funcdef-attr){#case-sensitivity-of-the-css-'attr()'-function:'attr()'-4
x-internal="'attr()'"} function must first be [converted to ASCII
lowercase](https://infra.spec.whatwg.org/#ascii-lowercase){#case-sensitivity-of-the-css-'attr()'-function:converted-to-ascii-lowercase
x-internal="converted-to-ascii-lowercase"}. The same function when
compared to other attributes must be compared according to its original
case. In both cases, to match the values must be [identical
to](https://infra.spec.whatwg.org/#string-is){#case-sensitivity-of-the-css-'attr()'-function:identical-to
x-internal="identical-to"} each other (and therefore the comparison is
case sensitive).

This is the same as comparing the name part of a CSS [attribute
selector](https://drafts.csswg.org/selectors/#attribute-selector){#case-sensitivity-of-the-css-'attr()'-function:attribute-selector
x-internal="attribute-selector"}, specified in the next section.

#### [4.16.2]{.secno} Case-sensitivity of selectors[](#case-sensitivity-of-selectors){.self-link}

Selectors leaves the case-sensitivity of element names, attribute names,
and attribute values to be defined by the host language.
[\[SELECTORS\]](references.html#refsSELECTORS)

When comparing a CSS element [type
selector](https://drafts.csswg.org/selectors/#type-selector){#case-sensitivity-of-selectors:type-selector
x-internal="type-selector"} to the names of [HTML
elements](infrastructure.html#html-elements){#case-sensitivity-of-selectors:html-elements}
in [HTML
documents](https://dom.spec.whatwg.org/#html-document){#case-sensitivity-of-selectors:html-documents
x-internal="html-documents"}, the CSS element [type
selector](https://drafts.csswg.org/selectors/#type-selector){#case-sensitivity-of-selectors:type-selector-2
x-internal="type-selector"} must first be [converted to ASCII
lowercase](https://infra.spec.whatwg.org/#ascii-lowercase){#case-sensitivity-of-selectors:converted-to-ascii-lowercase
x-internal="converted-to-ascii-lowercase"}. The same selector when
compared to other elements must be compared according to its original
case. In both cases, to match the values must be [identical
to](https://infra.spec.whatwg.org/#string-is){#case-sensitivity-of-selectors:identical-to
x-internal="identical-to"} each other (and therefore the comparison is
case sensitive).

When comparing the name part of a CSS [attribute
selector](https://drafts.csswg.org/selectors/#attribute-selector){#case-sensitivity-of-selectors:attribute-selector
x-internal="attribute-selector"} to the names of attributes on [HTML
elements](infrastructure.html#html-elements){#case-sensitivity-of-selectors:html-elements-2}
in [HTML
documents](https://dom.spec.whatwg.org/#html-document){#case-sensitivity-of-selectors:html-documents-2
x-internal="html-documents"}, the name part of the CSS [attribute
selector](https://drafts.csswg.org/selectors/#attribute-selector){#case-sensitivity-of-selectors:attribute-selector-2
x-internal="attribute-selector"} must first be [converted to ASCII
lowercase](https://infra.spec.whatwg.org/#ascii-lowercase){#case-sensitivity-of-selectors:converted-to-ascii-lowercase-2
x-internal="converted-to-ascii-lowercase"}. The same selector when
compared to other attributes must be compared according to its original
case. In both cases, the comparison is case-sensitive.

[Attribute
selectors](https://drafts.csswg.org/selectors/#attribute-selector){#case-sensitivity-of-selectors:attribute-selector-3
x-internal="attribute-selector"} on an [HTML
element](infrastructure.html#html-elements){#case-sensitivity-of-selectors:html-elements-3}
in an [HTML
document](https://dom.spec.whatwg.org/#html-document){#case-sensitivity-of-selectors:html-documents-3
x-internal="html-documents"} must treat the *values* of attributes with
the following names as [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#case-sensitivity-of-selectors:ascii-case-insensitive
x-internal="ascii-case-insensitive"}:

- `accept`
- `accept-charset`
- `align`
- `alink`
- `axis`
- `bgcolor`
- `charset`
- `checked`
- `clear`
- `codetype`
- `color`
- `compact`
- `declare`
- `defer`
- `dir`
- `direction`
- `disabled`
- `enctype`
- `face`
- `frame`
- `hreflang`
- `http-equiv`
- `lang`
- `language`
- `link`
- `media`
- `method`
- `multiple`
- `nohref`
- `noresize`
- `noshade`
- `nowrap`
- `readonly`
- `rel`
- `rev`
- `rules`
- `scope`
- `scrolling`
- `selected`
- `shape`
- `target`
- `text`
- `type`
- `valign`
- `valuetype`
- `vlink`

::: example
For example, the selector `[bgcolor="#ffffff"]` will match any HTML
element with a `bgcolor` attribute with values including `#ffffff`,
`#FFFFFF` and `#fffFFF`. This happens even if `bgcolor` has no effect
for a given element (e.g.,
[`div`{#case-sensitivity-of-selectors:the-div-element}](grouping-content.html#the-div-element)).

The selector `[type=a s]` will match any HTML element with a `type`
attribute whose value is `a`, but not whose value is `A`, due to the `s`
flag.
:::

All other attribute values and everything else must be treated as
entirely [identical
to](https://infra.spec.whatwg.org/#string-is){#case-sensitivity-of-selectors:identical-to-2
x-internal="identical-to"} each other for the purposes of selector
matching. This includes:

- [IDs](https://dom.spec.whatwg.org/#concept-id){#case-sensitivity-of-selectors:concept-id
  x-internal="concept-id"} and
  [classes](https://dom.spec.whatwg.org/#concept-class){#case-sensitivity-of-selectors:concept-class
  x-internal="concept-class"} in [no-quirks
  mode](https://dom.spec.whatwg.org/#concept-document-no-quirks){#case-sensitivity-of-selectors:no-quirks-mode
  x-internal="no-quirks-mode"} and [limited-quirks
  mode](https://dom.spec.whatwg.org/#concept-document-limited-quirks){#case-sensitivity-of-selectors:limited-quirks-mode
  x-internal="limited-quirks-mode"}

- the names of elements not in the [HTML
  namespace](https://infra.spec.whatwg.org/#html-namespace){#case-sensitivity-of-selectors:html-namespace-2
  x-internal="html-namespace-2"}

- the names of [HTML
  elements](infrastructure.html#html-elements){#case-sensitivity-of-selectors:html-elements-4}
  in [XML
  documents](https://dom.spec.whatwg.org/#xml-document){#case-sensitivity-of-selectors:xml-documents
  x-internal="xml-documents"}

- the names of attributes of elements not in the [HTML
  namespace](https://infra.spec.whatwg.org/#html-namespace){#case-sensitivity-of-selectors:html-namespace-2-2
  x-internal="html-namespace-2"}

- the names of attributes of [HTML
  elements](infrastructure.html#html-elements){#case-sensitivity-of-selectors:html-elements-5}
  in [XML
  documents](https://dom.spec.whatwg.org/#xml-document){#case-sensitivity-of-selectors:xml-documents-2
  x-internal="xml-documents"}

- the names of attributes that themselves have namespaces

Selectors defines that ID and class selectors (such as `#foo` and
`.bar`), when matched against elements in documents that are in [quirks
mode](https://dom.spec.whatwg.org/#concept-document-quirks){#case-sensitivity-of-selectors:quirks-mode
x-internal="quirks-mode"}, will be matched in an [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#case-sensitivity-of-selectors:ascii-case-insensitive-2
x-internal="ascii-case-insensitive"} manner. However, this does not
apply for attribute selectors with \"`id`\" or \"`class`\" as the name
part. The selector `[class="foobar"]` will treat its value as
case-sensitive even in [quirks
mode](https://dom.spec.whatwg.org/#concept-document-quirks){#case-sensitivity-of-selectors:quirks-mode-2
x-internal="quirks-mode"}.

#### [4.16.3]{.secno} Pseudo-classes[](#pseudo-classes){.self-link}

:::: {.mdn-anno .wrapped}
MDN

::: feature
[Pseudo-classes](https://developer.mozilla.org/en-US/docs/Web/CSS/Pseudo-classes "A CSS pseudo-class is a keyword added to a selector that specifies a special state of the selected element(s). For example, the pseudo-class :hover can be used to select a button when a user's pointer hovers over the button and this selected button can then be styled.")
:::
::::

There are a number of dynamic selectors that can be used with HTML. This
section defines when these selectors match HTML elements.
[\[SELECTORS\]](references.html#refsSELECTORS)
[\[CSSUI\]](references.html#refsCSSUI)

[`:defined`]{#selector-defined .dfn dfn-type="selector" noexport=""}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[:defined](https://developer.mozilla.org/en-US/docs/Web/CSS/:defined "The :defined CSS pseudo-class represents any element that has been defined. This includes any standard element built in to the browser, and custom elements that have been successfully defined (i.e. with the CustomElementRegistry.define() method).")

Support in all current engines.

::: support
[Firefox63+]{.firefox .yes}[Safari10+]{.safari .yes}[Chrome54+]{.chrome
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

The [`:defined`{#pseudo-classes:selector-defined}](#selector-defined)
[pseudo-class](https://drafts.csswg.org/selectors/#pseudo-class){#pseudo-classes:pseudo-class
x-internal="pseudo-class"} must match any element that is
[defined](https://dom.spec.whatwg.org/#concept-element-defined){#pseudo-classes:concept-element-defined
x-internal="concept-element-defined"}.

[`:link`]{#selector-link .dfn dfn-type="selector" noexport=""}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[:link](https://developer.mozilla.org/en-US/docs/Web/CSS/:link "The :link CSS pseudo-class represents an element that has not yet been visited. It matches every unvisited <a> or <area> element that has an href attribute.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera3.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer3+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari
iOS3.2+]{.safari_ios .yes}[Chrome Android?]{.chrome_android
.unknown}[WebView Android1.5+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

[`:visited`]{#selector-visited .dfn dfn-type="selector" noexport=""}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[:visited](https://developer.mozilla.org/en-US/docs/Web/CSS/:visited "The :visited CSS pseudo-class applies once the link has been visited by the user. For privacy reasons, the styles that can be modified using this selector are very limited. The :visited pseudo-class applies only <a> and <area> elements that have an href attribute.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera3.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

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

All
[`a`{#pseudo-classes:the-a-element}](text-level-semantics.html#the-a-element)
elements that have an
[`href`{#pseudo-classes:attr-hyperlink-href}](links.html#attr-hyperlink-href)
attribute, and all
[`area`{#pseudo-classes:the-area-element}](image-maps.html#the-area-element)
elements that have an
[`href`{#pseudo-classes:attr-hyperlink-href-2}](links.html#attr-hyperlink-href)
attribute, must match one of
[`:link`{#pseudo-classes:selector-link}](#selector-link) and
[`:visited`{#pseudo-classes:selector-visited}](#selector-visited).

Other specifications might apply more specific rules regarding how these
elements are to match these
[pseudo-classes](https://drafts.csswg.org/selectors/#pseudo-class){#pseudo-classes:pseudo-class-2
x-internal="pseudo-class"}, to mitigate some privacy concerns that apply
with straightforward implementations of this requirement.

[`:active`]{#selector-active .dfn dfn-type="selector" noexport=""}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[:active](https://developer.mozilla.org/en-US/docs/Web/CSS/:active "The :active CSS pseudo-class represents an element (such as a button) that is being activated by the user. When using a mouse, "activation" typically starts when the user presses down the primary mouse button.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

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

The [`:active`{#pseudo-classes:selector-active}](#selector-active)
[pseudo-class](https://drafts.csswg.org/selectors/#pseudo-class){#pseudo-classes:pseudo-class-3
x-internal="pseudo-class"} is defined to match an element "[while an
element is [*being activated*]{#concept-selector-active .dfn} by the
user]{cite="https://drafts.csswg.org/selectors/#the-active-pseudo"}".

To determine whether a particular element is *[being
activated](#concept-selector-active)* for the purposes of defining the
[`:active`{#pseudo-classes:selector-active-2}](#selector-active)
[pseudo-class](https://drafts.csswg.org/selectors/#pseudo-class){#pseudo-classes:pseudo-class-4
x-internal="pseudo-class"} only, an HTML user agent must use the first
relevant entry in the following list.

If the element is a [`button`{#pseudo-classes:the-button-element}](form-elements.html#the-button-element) element\
If the element is an [`input`{#pseudo-classes:the-input-element}](input.html#the-input-element) element whose [`type`{#pseudo-classes:attr-input-type}](input.html#attr-input-type) attribute is in the [Submit Button](input.html#submit-button-state-(type=submit)){#pseudo-classes:submit-button-state-(type=submit)}, [Image Button](input.html#image-button-state-(type=image)){#pseudo-classes:image-button-state-(type=image)}, [Reset Button](input.html#reset-button-state-(type=reset)){#pseudo-classes:reset-button-state-(type=reset)}, or [Button](input.html#button-state-(type=button)){#pseudo-classes:button-state-(type=button)} state\
If the element is an [`a`{#pseudo-classes:the-a-element-2}](text-level-semantics.html#the-a-element) element that has an [`href`{#pseudo-classes:attr-hyperlink-href-3}](links.html#attr-hyperlink-href) attribute\
If the element is an [`area`{#pseudo-classes:the-area-element-2}](image-maps.html#the-area-element) element that has an [`href`{#pseudo-classes:attr-hyperlink-href-4}](links.html#attr-hyperlink-href) attribute\
If the element is [focusable](interaction.html#focusable){#pseudo-classes:focusable}

:   The element is *[being activated](#concept-selector-active)* if it
    is [in a formal activation
    state](#in-a-formal-activation-state){#pseudo-classes:in-a-formal-activation-state}.

    For example, if the user is using a keyboard to push a
    [`button`{#pseudo-classes:the-button-element-2}](form-elements.html#the-button-element)
    element by pressing the space bar, the element would match this
    [pseudo-class](https://drafts.csswg.org/selectors/#pseudo-class){#pseudo-classes:pseudo-class-5
    x-internal="pseudo-class"} in between the time that the element
    received the
    [`keydown`{#pseudo-classes:event-keydown}](https://w3c.github.io/uievents/#event-type-keydown){x-internal="event-keydown"}
    event and the time the element received the
    [`keyup`{#pseudo-classes:event-keyup}](https://w3c.github.io/uievents/#event-type-keyup){x-internal="event-keyup"}
    event.

If the element is [being actively pointed at](#being-actively-pointed-at){#pseudo-classes:being-actively-pointed-at}

:   The element is *[being activated](#concept-selector-active)*.

An element is said to be [in a formal activation
state]{#in-a-formal-activation-state .dfn} between the time the user
begins to indicate an intent to trigger the element\'s [activation
behavior](https://dom.spec.whatwg.org/#eventtarget-activation-behavior){#pseudo-classes:activation-behaviour
x-internal="activation-behaviour"} and either the time the user stops
indicating an intent to trigger the element\'s [activation
behavior](https://dom.spec.whatwg.org/#eventtarget-activation-behavior){#pseudo-classes:activation-behaviour-2
x-internal="activation-behaviour"}, or the time the element\'s
[activation
behavior](https://dom.spec.whatwg.org/#eventtarget-activation-behavior){#pseudo-classes:activation-behaviour-3
x-internal="activation-behaviour"} has finished running, which ever
comes first.

An element is said to be [being actively pointed
at]{#being-actively-pointed-at .dfn} while the user indicates the
element using a pointing device while that pointing device is in the
\"down\" state (e.g. for a mouse, between the time the mouse button is
pressed and the time it is depressed; for a finger in a multitouch
environment, while the finger is touching the display surface).

Per the definition in Selectors,
[`:active`{#pseudo-classes:selector-active-3}](#selector-active) also
matches [flat
tree](https://drafts.csswg.org/css-scoping/#flat-tree){#pseudo-classes:flat-tree
x-internal="flat-tree"} ancestors of elements that are *[being
activated](#concept-selector-active)*.
[\[SELECTORS\]](references.html#refsSELECTORS)

Additionally, any element that is the [labeled
control](forms.html#labeled-control){#pseudo-classes:labeled-control} of
a
[`label`{#pseudo-classes:the-label-element}](forms.html#the-label-element)
element that is currently matching
[`:active`{#pseudo-classes:selector-active-4}](#selector-active), also
matches
[`:active`{#pseudo-classes:selector-active-5}](#selector-active). (But,
it does not count as being *[being
activated](#concept-selector-active)*.)

[`:hover`]{#selector-hover .dfn dfn-type="selector" noexport=""}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[:hover](https://developer.mozilla.org/en-US/docs/Web/CSS/:hover "The :hover CSS pseudo-class matches when the user interacts with an element with a pointing device, but does not necessarily activate it. It is generally triggered when the user hovers over an element with the cursor (mouse pointer).")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari2+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera4+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer4+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

The [`:hover`{#pseudo-classes:selector-hover}](#selector-hover)
[pseudo-class](https://drafts.csswg.org/selectors/#pseudo-class){#pseudo-classes:pseudo-class-6
x-internal="pseudo-class"} is defined to match an element "[while the
user [*designates*]{#concept-selector-hover .dfn} an element with a
pointing
device]{cite="https://drafts.csswg.org/selectors/#the-hover-pseudo"}".
For the purposes of defining the
[`:hover`{#pseudo-classes:selector-hover-2}](#selector-hover)
[pseudo-class](https://drafts.csswg.org/selectors/#pseudo-class){#pseudo-classes:pseudo-class-7
x-internal="pseudo-class"} only, an HTML user agent must consider an
element as being one that the user
*[designates](#concept-selector-hover)* if it is an element that the
user indicates using a pointing device.

Per the definition in Selectors,
[`:hover`{#pseudo-classes:selector-hover-3}](#selector-hover) also
matches [flat
tree](https://drafts.csswg.org/css-scoping/#flat-tree){#pseudo-classes:flat-tree-2
x-internal="flat-tree"} ancestors of elements that are
*[designated](#concept-selector-hover)*.
[\[SELECTORS\]](references.html#refsSELECTORS)

Additionally, any element that is the [labeled
control](forms.html#labeled-control){#pseudo-classes:labeled-control-2}
of a
[`label`{#pseudo-classes:the-label-element-2}](forms.html#the-label-element)
element that is currently matching
[`:hover`{#pseudo-classes:selector-hover-4}](#selector-hover), also
matches [`:hover`{#pseudo-classes:selector-hover-5}](#selector-hover).
(But, it does not count as being
*[designated](#concept-selector-hover)*.)

::: example
Consider in particular a fragment such as:

``` html
<p> <label for=c> <input id=a> </label> <span id=b> <input id=c> </span> </p>
```

If the user designates the element with ID \"`a`\" with their pointing
device, then the
[`p`{#pseudo-classes:the-p-element}](grouping-content.html#the-p-element)
element (and all its ancestors not shown in the snippet above), the
[`label`{#pseudo-classes:the-label-element-3}](forms.html#the-label-element)
element, the element with ID \"`a`\", and the element with ID \"`c`\"
will match the
[`:hover`{#pseudo-classes:selector-hover-6}](#selector-hover)
[pseudo-class](https://drafts.csswg.org/selectors/#pseudo-class){#pseudo-classes:pseudo-class-8
x-internal="pseudo-class"}. The element with ID \"`a`\" matches it by
being *[designated](#concept-selector-hover)*; the
[`label`{#pseudo-classes:the-label-element-4}](forms.html#the-label-element)
and
[`p`{#pseudo-classes:the-p-element-2}](grouping-content.html#the-p-element)
elements match it because of the condition in Selectors about flat tree
ancestors; and the element with ID \"`c`\" matches it through the
additional condition above on [labeled
controls](forms.html#labeled-control){#pseudo-classes:labeled-control-3}
(i.e., its
[`label`{#pseudo-classes:the-label-element-5}](forms.html#the-label-element)
element matches
[`:hover`{#pseudo-classes:selector-hover-7}](#selector-hover)). However,
the element with ID \"`b`\" does *not* match
[`:hover`{#pseudo-classes:selector-hover-8}](#selector-hover): its flat
tree descendant is not designated, even though that flat tree descendant
matches [`:hover`{#pseudo-classes:selector-hover-9}](#selector-hover).
:::

[`:focus`]{#selector-focus .dfn dfn-type="selector" noexport=""}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[:focus](https://developer.mozilla.org/en-US/docs/Web/CSS/:focus "The :focus CSS pseudo-class represents an element (such as a form input) that has received focus. It is generally triggered when the user clicks or taps on an element or selects it with the keyboard's Tab key.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera7+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer8+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

For the purposes of the CSS
[`:focus`{#pseudo-classes:selector-focus}](#selector-focus)
[pseudo-class](https://drafts.csswg.org/selectors/#pseudo-class){#pseudo-classes:pseudo-class-9
x-internal="pseudo-class"}, an [element has the
focus]{#element-has-the-focus .dfn} when:

- it is not itself a [navigable
  container](document-sequences.html#navigable-container){#pseudo-classes:navigable-container};
  and

- any of the following are true:

  - it is one of the elements listed in the [current focus chain of the
    top-level
    traversable](interaction.html#current-focus-chain-of-a-top-level-traversable){#pseudo-classes:current-focus-chain-of-a-top-level-traversable};
    or

  - its [shadow
    root](https://dom.spec.whatwg.org/#concept-element-shadow-root){#pseudo-classes:concept-element-shadow-root
    x-internal="concept-element-shadow-root"} `shadowRoot`{.variable} is
    not null and `shadowRoot`{.variable} is the
    [root](https://dom.spec.whatwg.org/#concept-tree-root){#pseudo-classes:root
    x-internal="root"} of at least one element that [has the
    focus](#element-has-the-focus){#pseudo-classes:element-has-the-focus}.

[`:target`]{#selector-target .dfn dfn-type="selector" noexport=""}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[:target](https://developer.mozilla.org/en-US/docs/Web/CSS/:target "The :target CSS pseudo-class represents a unique element (the target element) with an id matching the URL's fragment.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1.3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera9.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS2+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android2+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

For the purposes of the CSS
[`:target`{#pseudo-classes:selector-target}](#selector-target)
[pseudo-class](https://drafts.csswg.org/selectors/#pseudo-class){#pseudo-classes:pseudo-class-10
x-internal="pseudo-class"}, the
[`Document`{#pseudo-classes:document}](dom.html#document)\'s *target
elements* are a list containing the
[`Document`{#pseudo-classes:document-2}](dom.html#document)\'s [target
element](browsing-the-web.html#target-element){#pseudo-classes:target-element},
if it is not null, or containing no elements, if it is.
[\[SELECTORS\]](references.html#refsSELECTORS)

[`:popover-open`]{#selector-popover-open .dfn dfn-type="selector"
noexport=""}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[:popover-open](https://developer.mozilla.org/en-US/docs/Web/CSS/:popover-open "The :popover-open CSS pseudo-class represents a popover element (i.e. one with a popover attribute) that is in the showing state. You can use this to apply style to popover elements only when they are shown.")

Support in all current engines.

::: support
[Firefox[🔰
114+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safaripreview+]{.safari .yes}[Chrome114+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge114+]{.edge_blink .yes}

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
[`:popover-open`{#pseudo-classes:selector-popover-open}](#selector-popover-open)
[pseudo-class](https://drafts.csswg.org/selectors/#pseudo-class){#pseudo-classes:pseudo-class-11
x-internal="pseudo-class"} is defined to match any [HTML
element](infrastructure.html#html-elements){#pseudo-classes:html-elements}
whose
[`popover`{#pseudo-classes:attr-popover}](popover.html#attr-popover)
attribute is not in the [No Popover
State](popover.html#attr-popover-none-state){#pseudo-classes:attr-popover-none-state}
and whose [popover visibility
state](popover.html#popover-visibility-state){#pseudo-classes:popover-visibility-state}
is
[showing](popover.html#popover-showing-state){#pseudo-classes:popover-showing-state}.

[`:enabled`]{#selector-enabled .dfn dfn-type="selector" noexport=""}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[:enabled](https://developer.mozilla.org/en-US/docs/Web/CSS/:enabled "The :enabled CSS pseudo-class represents any enabled element. An element is enabled if it can be activated (selected, clicked on, typed into, etc.) or accept focus. The element also has a disabled state, in which it can't be activated or accept focus.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera9+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android2+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

The [`:enabled`{#pseudo-classes:selector-enabled}](#selector-enabled)
[pseudo-class](https://drafts.csswg.org/selectors/#pseudo-class){#pseudo-classes:pseudo-class-12
x-internal="pseudo-class"} must match any
[`button`{#pseudo-classes:the-button-element-3}](form-elements.html#the-button-element),
[`input`{#pseudo-classes:the-input-element-2}](input.html#the-input-element),
[`select`{#pseudo-classes:the-select-element}](form-elements.html#the-select-element),
[`textarea`{#pseudo-classes:the-textarea-element}](form-elements.html#the-textarea-element),
[`optgroup`{#pseudo-classes:the-optgroup-element}](form-elements.html#the-optgroup-element),
[`option`{#pseudo-classes:the-option-element}](form-elements.html#the-option-element),
[`fieldset`{#pseudo-classes:the-fieldset-element}](form-elements.html#the-fieldset-element)
element, or [form-associated custom
element](custom-elements.html#form-associated-custom-element){#pseudo-classes:form-associated-custom-element}
that is not [actually
disabled](#concept-element-disabled){#pseudo-classes:concept-element-disabled}.

[`:disabled`]{#selector-disabled .dfn dfn-type="selector" noexport=""}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[:disabled](https://developer.mozilla.org/en-US/docs/Web/CSS/:disabled "The :disabled CSS pseudo-class represents any disabled element. An element is disabled if it can't be activated (selected, clicked on, typed into, etc.) or accept focus. The element also has an enabled state, in which it can be activated or accept focus.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera9+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android2+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

The [`:disabled`{#pseudo-classes:selector-disabled}](#selector-disabled)
[pseudo-class](https://drafts.csswg.org/selectors/#pseudo-class){#pseudo-classes:pseudo-class-13
x-internal="pseudo-class"} must match any element that is [actually
disabled](#concept-element-disabled){#pseudo-classes:concept-element-disabled-2}.

[`:checked`]{#selector-checked .dfn dfn-type="selector" noexport=""}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[:checked](https://developer.mozilla.org/en-US/docs/Web/CSS/:checked "The :checked CSS pseudo-class selector represents any radio (<input type="radio">), checkbox (<input type="checkbox">), or option (<option> in a <select>) element that is checked or toggled to an on state.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera9+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android4+]{.firefox_android .yes}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android18+]{.chrome_android .yes}[WebView
Android2+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

The [`:checked`{#pseudo-classes:selector-checked}](#selector-checked)
[pseudo-class](https://drafts.csswg.org/selectors/#pseudo-class){#pseudo-classes:pseudo-class-14
x-internal="pseudo-class"} must match any element falling into one of
the following categories:

- [`input`{#pseudo-classes:the-input-element-3}](input.html#the-input-element)
  elements whose
  [`type`{#pseudo-classes:attr-input-type-2}](input.html#attr-input-type)
  attribute is in the
  [Checkbox](input.html#checkbox-state-(type=checkbox)){#pseudo-classes:checkbox-state-(type=checkbox)}
  state and whose
  [checkedness](form-control-infrastructure.html#concept-fe-checked){#pseudo-classes:concept-fe-checked}
  state is true
- [`input`{#pseudo-classes:the-input-element-4}](input.html#the-input-element)
  elements whose
  [`type`{#pseudo-classes:attr-input-type-3}](input.html#attr-input-type)
  attribute is in the [Radio
  Button](input.html#radio-button-state-(type=radio)){#pseudo-classes:radio-button-state-(type=radio)}
  state and whose
  [checkedness](form-control-infrastructure.html#concept-fe-checked){#pseudo-classes:concept-fe-checked-2}
  state is true
- [`option`{#pseudo-classes:the-option-element-2}](form-elements.html#the-option-element)
  elements whose
  [selectedness](form-elements.html#concept-option-selectedness){#pseudo-classes:concept-option-selectedness}
  is true

[`:indeterminate`]{#selector-indeterminate .dfn dfn-type="selector"
noexport=""}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[:indeterminate](https://developer.mozilla.org/en-US/docs/Web/CSS/:indeterminate "The :indeterminate CSS pseudo-class represents any form element whose state is indeterminate, such as checkboxes which have their HTML indeterminate attribute set to true, radio buttons which are members of a group in which all radio buttons are unchecked, and indeterminate <progress> elements.")

Support in all current engines.

::: support
[Firefox2+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera9+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

The
[`:indeterminate`{#pseudo-classes:selector-indeterminate}](#selector-indeterminate)
[pseudo-class](https://drafts.csswg.org/selectors/#pseudo-class){#pseudo-classes:pseudo-class-15
x-internal="pseudo-class"} must match any element falling into one of
the following categories:

- [`input`{#pseudo-classes:the-input-element-5}](input.html#the-input-element)
  elements whose
  [`type`{#pseudo-classes:attr-input-type-4}](input.html#attr-input-type)
  attribute is in the
  [Checkbox](input.html#checkbox-state-(type=checkbox)){#pseudo-classes:checkbox-state-(type=checkbox)-2}
  state and whose
  [`indeterminate`{#pseudo-classes:dom-input-indeterminate}](input.html#dom-input-indeterminate)
  IDL attribute is set to true
- [`input`{#pseudo-classes:the-input-element-6}](input.html#the-input-element)
  elements whose
  [`type`{#pseudo-classes:attr-input-type-5}](input.html#attr-input-type)
  attribute is in the [Radio
  Button](input.html#radio-button-state-(type=radio)){#pseudo-classes:radio-button-state-(type=radio)-2}
  state and whose [radio button
  group](input.html#radio-button-group){#pseudo-classes:radio-button-group}
  contains no
  [`input`{#pseudo-classes:the-input-element-7}](input.html#the-input-element)
  elements whose
  [checkedness](form-control-infrastructure.html#concept-fe-checked){#pseudo-classes:concept-fe-checked-3}
  state is true.
- [`progress`{#pseudo-classes:the-progress-element}](form-elements.html#the-progress-element)
  elements with no
  [`value`{#pseudo-classes:attr-progress-value}](form-elements.html#attr-progress-value)
  content attribute

[`:default`]{#selector-default .dfn dfn-type="selector" noexport=""}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[:default](https://developer.mozilla.org/en-US/docs/Web/CSS/:default "The :default CSS pseudo-class selects form elements that are the default in a group of related elements.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome10+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

The [`:default`{#pseudo-classes:selector-default}](#selector-default)
[pseudo-class](https://drafts.csswg.org/selectors/#pseudo-class){#pseudo-classes:pseudo-class-16
x-internal="pseudo-class"} must match any element falling into one of
the following categories:

- [Submit
  buttons](forms.html#concept-submit-button){#pseudo-classes:concept-submit-button}
  that are [default
  buttons](form-control-infrastructure.html#default-button){#pseudo-classes:default-button}
  of their [form
  owner](form-control-infrastructure.html#form-owner){#pseudo-classes:form-owner}.
- [`input`{#pseudo-classes:the-input-element-8}](input.html#the-input-element)
  elements to which the
  [`checked`{#pseudo-classes:attr-input-checked}](input.html#attr-input-checked)
  attribute applies and that have a
  [`checked`{#pseudo-classes:attr-input-checked-2}](input.html#attr-input-checked)
  attribute
- [`option`{#pseudo-classes:the-option-element-3}](form-elements.html#the-option-element)
  elements that have a
  [`selected`{#pseudo-classes:attr-option-selected}](form-elements.html#attr-option-selected)
  attribute

[`:placeholder-shown`]{#selector-placeholder-shown .dfn
dfn-type="selector" noexport=""}

The
[`:placeholder-shown`{#pseudo-classes:selector-placeholder-shown}](#selector-placeholder-shown)
[pseudo-class](https://drafts.csswg.org/selectors/#pseudo-class){#pseudo-classes:pseudo-class-17
x-internal="pseudo-class"} must match any element falling into one of
the following categories:

- [`input`{#pseudo-classes:the-input-element-9}](input.html#the-input-element)
  elements that have a
  [`placeholder`{#pseudo-classes:attr-input-placeholder}](input.html#attr-input-placeholder)
  attribute whose value is currently being presented to the user
- [`textarea`{#pseudo-classes:the-textarea-element-2}](form-elements.html#the-textarea-element)
  elements that have a
  [`placeholder`{#pseudo-classes:attr-textarea-placeholder}](form-elements.html#attr-textarea-placeholder)
  attribute whose value is currently being presented to the user

[`:valid`]{#selector-valid .dfn dfn-type="selector" noexport=""}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[:valid](https://developer.mozilla.org/en-US/docs/Web/CSS/:valid "The :valid CSS pseudo-class represents any <input> or other <form> element whose contents validate successfully. This allows to easily make valid fields adopt an appearance that helps the user confirm that their data is formatted properly.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome10+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

The [`:valid`{#pseudo-classes:selector-valid}](#selector-valid)
[pseudo-class](https://drafts.csswg.org/selectors/#pseudo-class){#pseudo-classes:pseudo-class-18
x-internal="pseudo-class"} must match any element falling into one of
the following categories:

- elements that are [candidates for constraint
  validation](form-control-infrastructure.html#candidate-for-constraint-validation){#pseudo-classes:candidate-for-constraint-validation}
  and that [satisfy their
  constraints](form-control-infrastructure.html#concept-fv-valid){#pseudo-classes:concept-fv-valid}
- [`form`{#pseudo-classes:the-form-element}](forms.html#the-form-element)
  elements that are not the [form
  owner](form-control-infrastructure.html#form-owner){#pseudo-classes:form-owner-2}
  of any elements that themselves are [candidates for constraint
  validation](form-control-infrastructure.html#candidate-for-constraint-validation){#pseudo-classes:candidate-for-constraint-validation-2}
  but do not [satisfy their
  constraints](form-control-infrastructure.html#concept-fv-valid){#pseudo-classes:concept-fv-valid-2}
- [`fieldset`{#pseudo-classes:the-fieldset-element-2}](form-elements.html#the-fieldset-element)
  elements that have no descendant elements that themselves are
  [candidates for constraint
  validation](form-control-infrastructure.html#candidate-for-constraint-validation){#pseudo-classes:candidate-for-constraint-validation-3}
  but do not [satisfy their
  constraints](form-control-infrastructure.html#concept-fv-valid){#pseudo-classes:concept-fv-valid-3}

[`:invalid`]{#selector-invalid .dfn dfn-type="selector" noexport=""}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[:invalid](https://developer.mozilla.org/en-US/docs/Web/CSS/:invalid "The :invalid CSS pseudo-class represents any <form>, <fieldset>, <input> or other <form> element whose contents fail to validate.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome10+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

The [`:invalid`{#pseudo-classes:selector-invalid}](#selector-invalid)
[pseudo-class](https://drafts.csswg.org/selectors/#pseudo-class){#pseudo-classes:pseudo-class-19
x-internal="pseudo-class"} must match any element falling into one of
the following categories:

- elements that are [candidates for constraint
  validation](form-control-infrastructure.html#candidate-for-constraint-validation){#pseudo-classes:candidate-for-constraint-validation-4}
  but that do not [satisfy their
  constraints](form-control-infrastructure.html#concept-fv-valid){#pseudo-classes:concept-fv-valid-4}
- [`form`{#pseudo-classes:the-form-element-2}](forms.html#the-form-element)
  elements that are the [form
  owner](form-control-infrastructure.html#form-owner){#pseudo-classes:form-owner-3}
  of one or more elements that themselves are [candidates for constraint
  validation](form-control-infrastructure.html#candidate-for-constraint-validation){#pseudo-classes:candidate-for-constraint-validation-5}
  but do not [satisfy their
  constraints](form-control-infrastructure.html#concept-fv-valid){#pseudo-classes:concept-fv-valid-5}
- [`fieldset`{#pseudo-classes:the-fieldset-element-3}](form-elements.html#the-fieldset-element)
  elements that have of one or more descendant elements that themselves
  are [candidates for constraint
  validation](form-control-infrastructure.html#candidate-for-constraint-validation){#pseudo-classes:candidate-for-constraint-validation-6}
  but do not [satisfy their
  constraints](form-control-infrastructure.html#concept-fv-valid){#pseudo-classes:concept-fv-valid-6}

[`:user-valid`]{#selector-user-valid .dfn dfn-type="selector"
noexport=""}

The
[`:user-valid`{#pseudo-classes:selector-user-valid}](#selector-user-valid)
[pseudo-class](https://drafts.csswg.org/selectors/#pseudo-class){#pseudo-classes:pseudo-class-20
x-internal="pseudo-class"} must match
[`input`{#pseudo-classes:the-input-element-10}](input.html#the-input-element),
[`textarea`{#pseudo-classes:the-textarea-element-3}](form-elements.html#the-textarea-element),
and
[`select`{#pseudo-classes:the-select-element-2}](form-elements.html#the-select-element)
elements whose [user
validity](form-control-infrastructure.html#user-validity){#pseudo-classes:user-validity}
is true, are [candidates for constraint
validation](form-control-infrastructure.html#candidate-for-constraint-validation){#pseudo-classes:candidate-for-constraint-validation-7},
and that [satisfy their
constraints](form-control-infrastructure.html#concept-fv-valid){#pseudo-classes:concept-fv-valid-7}.

[`:user-invalid`]{#selector-user-invalid .dfn dfn-type="selector"
noexport=""}

The
[`:user-invalid`{#pseudo-classes:selector-user-invalid}](#selector-user-invalid)
[pseudo-class](https://drafts.csswg.org/selectors/#pseudo-class){#pseudo-classes:pseudo-class-21
x-internal="pseudo-class"} must match
[`input`{#pseudo-classes:the-input-element-11}](input.html#the-input-element),
[`textarea`{#pseudo-classes:the-textarea-element-4}](form-elements.html#the-textarea-element),
and
[`select`{#pseudo-classes:the-select-element-3}](form-elements.html#the-select-element)
elements whose [user
validity](form-control-infrastructure.html#user-validity){#pseudo-classes:user-validity-2}
is true, are [candidates for constraint
validation](form-control-infrastructure.html#candidate-for-constraint-validation){#pseudo-classes:candidate-for-constraint-validation-8}
but do not [satisfy their
constraints](form-control-infrastructure.html#concept-fv-valid){#pseudo-classes:concept-fv-valid-8}.

[`:in-range`]{#selector-in-range .dfn dfn-type="selector" noexport=""}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[:in-range](https://developer.mozilla.org/en-US/docs/Web/CSS/:in-range "The :in-range CSS pseudo-class represents an <input> element whose current value is within the range limits specified by the min and max attributes.")

Support in all current engines.

::: support
[Firefox29+]{.firefox .yes}[Safari5.1+]{.safari .yes}[Chrome10+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)13+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android16+]{.firefox_android .yes}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android2.2+]{.webview_android .yes}[Samsung
Internet1.0+]{.samsunginternet_android .yes}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

The [`:in-range`{#pseudo-classes:selector-in-range}](#selector-in-range)
[pseudo-class](https://drafts.csswg.org/selectors/#pseudo-class){#pseudo-classes:pseudo-class-22
x-internal="pseudo-class"} must match all elements that are [candidates
for constraint
validation](form-control-infrastructure.html#candidate-for-constraint-validation){#pseudo-classes:candidate-for-constraint-validation-9},
[have range
limitations](input.html#have-range-limitations){#pseudo-classes:have-range-limitations},
and that are neither [suffering from an
underflow](form-control-infrastructure.html#suffering-from-an-underflow){#pseudo-classes:suffering-from-an-underflow}
nor [suffering from an
overflow](form-control-infrastructure.html#suffering-from-an-overflow){#pseudo-classes:suffering-from-an-overflow}.

[`:out-of-range`]{#selector-out-of-range .dfn dfn-type="selector"
noexport=""}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[:out-of-range](https://developer.mozilla.org/en-US/docs/Web/CSS/:out-of-range "The :out-of-range CSS pseudo-class represents an <input> element whose current value is outside the range limits specified by the min and max attributes.")

Support in all current engines.

::: support
[Firefox29+]{.firefox .yes}[Safari5.1+]{.safari .yes}[Chrome10+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)13+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android16+]{.firefox_android .yes}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android2.2+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

The
[`:out-of-range`{#pseudo-classes:selector-out-of-range}](#selector-out-of-range)
[pseudo-class](https://drafts.csswg.org/selectors/#pseudo-class){#pseudo-classes:pseudo-class-23
x-internal="pseudo-class"} must match all elements that are [candidates
for constraint
validation](form-control-infrastructure.html#candidate-for-constraint-validation){#pseudo-classes:candidate-for-constraint-validation-10},
[have range
limitations](input.html#have-range-limitations){#pseudo-classes:have-range-limitations-2},
and that are either [suffering from an
underflow](form-control-infrastructure.html#suffering-from-an-underflow){#pseudo-classes:suffering-from-an-underflow-2}
or [suffering from an
overflow](form-control-infrastructure.html#suffering-from-an-overflow){#pseudo-classes:suffering-from-an-overflow-2}.

[`:required`]{#selector-required .dfn dfn-type="selector" noexport=""}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[:required](https://developer.mozilla.org/en-US/docs/Web/CSS/:required "The :required CSS pseudo-class represents any <input>, <select>, or <textarea> element that has the required attribute set on it.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome10+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android4.4.3+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

The [`:required`{#pseudo-classes:selector-required}](#selector-required)
[pseudo-class](https://drafts.csswg.org/selectors/#pseudo-class){#pseudo-classes:pseudo-class-24
x-internal="pseudo-class"} must match any element falling into one of
the following categories:

- [`input`{#pseudo-classes:the-input-element-12}](input.html#the-input-element)
  elements that are *[required](input.html#concept-input-required)*
- [`select`{#pseudo-classes:the-select-element-4}](form-elements.html#the-select-element)
  elements that have a
  [`required`{#pseudo-classes:attr-select-required}](form-elements.html#attr-select-required)
  attribute
- [`textarea`{#pseudo-classes:the-textarea-element-5}](form-elements.html#the-textarea-element)
  elements that have a
  [`required`{#pseudo-classes:attr-textarea-required}](form-elements.html#attr-textarea-required)
  attribute

[`:optional`]{#selector-optional .dfn dfn-type="selector" noexport=""}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[:optional](https://developer.mozilla.org/en-US/docs/Web/CSS/:optional "The :optional CSS pseudo-class represents any <input>, <select>, or <textarea> element that does not have the required attribute set on it.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome10+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

The [`:optional`{#pseudo-classes:selector-optional}](#selector-optional)
[pseudo-class](https://drafts.csswg.org/selectors/#pseudo-class){#pseudo-classes:pseudo-class-25
x-internal="pseudo-class"} must match any element falling into one of
the following categories:

- [`input`{#pseudo-classes:the-input-element-13}](input.html#the-input-element)
  elements to which the
  [`required`{#pseudo-classes:attr-input-required}](input.html#attr-input-required)
  attribute applies that are not
  *[required](input.html#concept-input-required)*
- [`select`{#pseudo-classes:the-select-element-5}](form-elements.html#the-select-element)
  elements that do not have a
  [`required`{#pseudo-classes:attr-select-required-2}](form-elements.html#attr-select-required)
  attribute
- [`textarea`{#pseudo-classes:the-textarea-element-6}](form-elements.html#the-textarea-element)
  elements that do not have a
  [`required`{#pseudo-classes:attr-textarea-required-2}](form-elements.html#attr-textarea-required)
  attribute

[`:autofill`]{#selector-autofill .dfn dfn-type="selector" noexport=""}

::::: {.mdn-anno .wrapped .before}
MDN

:::: feature
[:autofill](https://developer.mozilla.org/en-US/docs/Web/CSS/:autofill "The :autofill CSS pseudo-class matches when an <input> element has its value autofilled by the browser. The class stops matching if the user edits the field.")

::: support
[Firefox86+]{.firefox .yes}[Safari15+]{.safari .yes}[Chrome[🔰
1+]{title="Requires a prefix or alternative name."}]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge[🔰
79+]{title="Requires a prefix or alternative name."}]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS15+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

[`:-webkit-autofill`]{#selector-webkit-autofill .dfn dfn-type="selector"
noexport=""}

The [`:autofill`{#pseudo-classes:selector-autofill}](#selector-autofill)
and
[`:-webkit-autofill`{#pseudo-classes:selector-webkit-autofill}](#selector-webkit-autofill)
[pseudo-classes](https://drafts.csswg.org/selectors/#pseudo-class){#pseudo-classes:pseudo-class-26
x-internal="pseudo-class"} must match
[`input`{#pseudo-classes:the-input-element-14}](input.html#the-input-element)
elements which have been autofilled by user agent. These pseudo-classes
must stop matching if the user edits the autofilled field.

One way such autofilling might happen is via the
[`autocomplete`{#pseudo-classes:attr-fe-autocomplete}](form-control-infrastructure.html#attr-fe-autocomplete)
attribute, but user agents could autofill even without that attribute
being involved.

[`:read-only`]{#selector-read-only .dfn dfn-type="selector" noexport=""}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[:read-only](https://developer.mozilla.org/en-US/docs/Web/CSS/:read-only "The :read-only CSS pseudo-class represents an element (such as input or textarea) that is not editable by the user.")

Support in all current engines.

::: support
[Firefox78+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera9+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)13+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android[🔰
4+]{title="Requires a prefix or alternative name."}]{.firefox_android
.yes}[Safari iOS?]{.safari_ios .unknown}[Chrome
Android?]{.chrome_android .unknown}[WebView Android37+]{.webview_android
.yes}[Samsung Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

[`:read-write`]{#selector-read-write .dfn dfn-type="selector"
noexport=""}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[:read-write](https://developer.mozilla.org/en-US/docs/Web/CSS/:read-write "The :read-write CSS pseudo-class represents an element (such as input or textarea) that is editable by the user.")

Support in all current engines.

::: support
[Firefox78+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera9+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)13+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android[🔰
4+]{title="Requires a prefix or alternative name."}]{.firefox_android
.yes}[Safari iOS?]{.safari_ios .unknown}[Chrome
Android?]{.chrome_android .unknown}[WebView Android37+]{.webview_android
.yes}[Samsung Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

The
[`:read-write`{#pseudo-classes:selector-read-write}](#selector-read-write)
[pseudo-class](https://drafts.csswg.org/selectors/#pseudo-class){#pseudo-classes:pseudo-class-27
x-internal="pseudo-class"} must match any element falling into one of
the following categories, which for the purposes of Selectors are thus
considered *user-alterable*:
[\[SELECTORS\]](references.html#refsSELECTORS)

- [`input`{#pseudo-classes:the-input-element-15}](input.html#the-input-element)
  elements to which the
  [`readonly`{#pseudo-classes:attr-input-readonly}](input.html#attr-input-readonly)
  attribute applies, and that are
  *[mutable](form-control-infrastructure.html#concept-fe-mutable)* (i.e.
  that do not have the
  [`readonly`{#pseudo-classes:attr-input-readonly-2}](input.html#attr-input-readonly)
  attribute specified and that are not
  [disabled](form-control-infrastructure.html#concept-fe-disabled){#pseudo-classes:concept-fe-disabled})
- [`textarea`{#pseudo-classes:the-textarea-element-7}](form-elements.html#the-textarea-element)
  elements that do not have a
  [`readonly`{#pseudo-classes:attr-textarea-readonly}](form-elements.html#attr-textarea-readonly)
  attribute, and that are not
  [disabled](form-control-infrastructure.html#concept-fe-disabled){#pseudo-classes:concept-fe-disabled-2}
- elements that are [editing
  hosts](interaction.html#editing-host){#pseudo-classes:editing-host} or
  [editable](https://w3c.github.io/editing/docs/execCommand/#editable){#pseudo-classes:editable
  x-internal="editable"} and are neither
  [`input`{#pseudo-classes:the-input-element-16}](input.html#the-input-element)
  elements nor
  [`textarea`{#pseudo-classes:the-textarea-element-8}](form-elements.html#the-textarea-element)
  elements

The
[`:read-only`{#pseudo-classes:selector-read-only}](#selector-read-only)
[pseudo-class](https://drafts.csswg.org/selectors/#pseudo-class){#pseudo-classes:pseudo-class-28
x-internal="pseudo-class"} must match all other [HTML
elements](infrastructure.html#html-elements){#pseudo-classes:html-elements-2}.

[`:modal`]{#selector-modal .dfn dfn-type="selector" noexport=""}

The [`:modal`{#pseudo-classes:selector-modal}](#selector-modal)
[pseudo-class](https://drafts.csswg.org/selectors/#pseudo-class){#pseudo-classes:pseudo-class-29
x-internal="pseudo-class"} must match any element falling into one of
the following categories:

- [`dialog`{#pseudo-classes:the-dialog-element}](interactive-elements.html#the-dialog-element)
  elements whose [is
  modal](interactive-elements.html#is-modal){#pseudo-classes:is-modal}
  is true
- elements whose [fullscreen
  flag](https://fullscreen.spec.whatwg.org/#fullscreen-flag){#pseudo-classes:fullscreen-flag
  x-internal="fullscreen-flag"} is true

[`:dir(ltr)`]{#selector-ltr .dfn dfn-type="selector" noexport=""}

::::: {.mdn-anno .wrapped .before}
MDN

:::: feature
[:dir](https://developer.mozilla.org/en-US/docs/Web/CSS/:dir "The :dir() CSS pseudo-class matches elements based on the directionality of the text contained in them.")

::: support
[Firefox49+]{.firefox .yes}[Safari16.4+]{.safari .yes}[ChromeNo]{.chrome
.no}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[EdgeNo]{.edge_blink .no}

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

The [`:dir(ltr)`{#pseudo-classes:selector-ltr}](#selector-ltr)
[pseudo-class](https://drafts.csswg.org/selectors/#pseudo-class){#pseudo-classes:pseudo-class-30
x-internal="pseudo-class"} must match all elements whose
[directionality](dom.html#the-directionality){#pseudo-classes:the-directionality}
is \'[ltr](dom.html#concept-ltr){#pseudo-classes:concept-ltr}\'.

[`:dir(rtl)`]{#selector-rtl .dfn dfn-type="selector" noexport=""}

The [`:dir(rtl)`{#pseudo-classes:selector-rtl}](#selector-rtl)
[pseudo-class](https://drafts.csswg.org/selectors/#pseudo-class){#pseudo-classes:pseudo-class-31
x-internal="pseudo-class"} must match all elements whose
[directionality](dom.html#the-directionality){#pseudo-classes:the-directionality-2}
is \'[rtl](dom.html#concept-rtl){#pseudo-classes:concept-rtl}\'.

[Custom state pseudo-class]{#selector-custom .dfn dfn-type="selector"
noexport=""}

The
[`:state(`{#pseudo-classes:selector-custom}`identifier`{.variable}`)`{#pseudo-classes:selector-custom}](#selector-custom)
pseudo-class must match all [custom
element](custom-elements.html#custom-element){#pseudo-classes:custom-element}s
whose [states
set](custom-elements.html#states-set){#pseudo-classes:states-set}\'s
[set
entries](https://webidl.spec.whatwg.org/#dfn-set-entries){#pseudo-classes:set-entries
x-internal="set-entries"} contains `identifier`{.variable}.

[`:playing`]{#selector-playing .dfn dfn-type="selector" noexport=""}

The [`:playing`{#pseudo-classes:selector-playing}](#selector-playing)
[pseudo-class](https://drafts.csswg.org/selectors/#pseudo-class){#pseudo-classes:pseudo-class-32
x-internal="pseudo-class"} must match all [media
elements](media.html#media-element){#pseudo-classes:media-element} whose
[`paused`{#pseudo-classes:dom-media-paused}](media.html#dom-media-paused)
attribute is false.

[`:paused`]{#selector-paused .dfn dfn-type="selector" noexport=""}

The [`:paused`{#pseudo-classes:selector-paused}](#selector-paused)
[pseudo-class](https://drafts.csswg.org/selectors/#pseudo-class){#pseudo-classes:pseudo-class-33
x-internal="pseudo-class"} must match all [media
elements](media.html#media-element){#pseudo-classes:media-element-2}
whose
[`paused`{#pseudo-classes:dom-media-paused-2}](media.html#dom-media-paused)
attribute is true.

[`:seeking`]{#selector-seeking .dfn dfn-type="selector" noexport=""}

The [`:seeking`{#pseudo-classes:selector-seeking}](#selector-seeking)
[pseudo-class](https://drafts.csswg.org/selectors/#pseudo-class){#pseudo-classes:pseudo-class-34
x-internal="pseudo-class"} must match all [media
elements](media.html#media-element){#pseudo-classes:media-element-3}
whose
[`seeking`{#pseudo-classes:dom-media-seeking}](media.html#dom-media-seeking)
attribute is true.

[`:buffering`]{#selector-buffering .dfn dfn-type="selector" noexport=""}

The
[`:buffering`{#pseudo-classes:selector-buffering}](#selector-buffering)
[pseudo-class](https://drafts.csswg.org/selectors/#pseudo-class){#pseudo-classes:pseudo-class-35
x-internal="pseudo-class"} must match all [media
elements](media.html#media-element){#pseudo-classes:media-element-4}
whose
[`paused`{#pseudo-classes:dom-media-paused-3}](media.html#dom-media-paused)
attribute is false,
[`networkState`{#pseudo-classes:dom-media-networkstate}](media.html#dom-media-networkstate)
attribute is
[`NETWORK_LOADING`{#pseudo-classes:dom-media-network_loading}](media.html#dom-media-network_loading),
and ready state is
[`HAVE_CURRENT_DATA`{#pseudo-classes:dom-media-have_current_data}](media.html#dom-media-have_current_data)
or less.

[`:stalled`]{#selector-stalled .dfn dfn-type="selector" noexport=""}

The [`:stalled`{#pseudo-classes:selector-stalled}](#selector-stalled)
[pseudo-class](https://drafts.csswg.org/selectors/#pseudo-class){#pseudo-classes:pseudo-class-36
x-internal="pseudo-class"} must match all [media
elements](media.html#media-element){#pseudo-classes:media-element-5}
that match the
[`:buffering`{#pseudo-classes:selector-buffering-2}](#selector-buffering)
[pseudo-class](https://drafts.csswg.org/selectors/#pseudo-class){#pseudo-classes:pseudo-class-37
x-internal="pseudo-class"} and whose [is currently
stalled](media.html#is-currently-stalled){#pseudo-classes:is-currently-stalled}
is true.

[`:muted`]{#selector-muted .dfn dfn-type="selector" noexport=""}

The [`:muted`{#pseudo-classes:selector-muted}](#selector-muted)
[pseudo-class](https://drafts.csswg.org/selectors/#pseudo-class){#pseudo-classes:pseudo-class-38
x-internal="pseudo-class"} must match all [media
elements](media.html#media-element){#pseudo-classes:media-element-6}
that are
[muted](media.html#concept-media-muted){#pseudo-classes:concept-media-muted}.

[`:volume-locked`]{#selector-volume-locked .dfn dfn-type="selector"
noexport=""}

The
[`:volume-locked`{#pseudo-classes:selector-volume-locked}](#selector-volume-locked)
[pseudo-class](https://drafts.csswg.org/selectors/#pseudo-class){#pseudo-classes:pseudo-class-39
x-internal="pseudo-class"} must match all [media
elements](media.html#media-element){#pseudo-classes:media-element-7}
when the user agent\'s [volume
locked](media.html#volume-locked){#pseudo-classes:volume-locked} is
true.

[`:open`]{#selector-open .dfn dfn-type="selector"}

The [`:open`{#pseudo-classes:selector-open}](#selector-open)
[pseudo-class](https://drafts.csswg.org/selectors/#pseudo-class){#pseudo-classes:pseudo-class-40
x-internal="pseudo-class"} must match any element falling into one of
the following categories:

- [`details`{#pseudo-classes:the-details-element}](interactive-elements.html#the-details-element)
  elements that have an
  [`open`{#pseudo-classes:attr-details-open}](interactive-elements.html#attr-details-open)
  attribute

- [`dialog`{#pseudo-classes:the-dialog-element-2}](interactive-elements.html#the-dialog-element)
  elements that have an
  [`open`{#pseudo-classes:attr-dialog-open}](interactive-elements.html#attr-dialog-open)
  attribute

- [`select`{#pseudo-classes:the-select-element-6}](form-elements.html#the-select-element)
  elements that are a [drop-down
  box](rendering.html#drop-down-box){#pseudo-classes:drop-down-box} and
  whose drop-down boxes are open

- [`input`{#pseudo-classes:the-input-element-17}](input.html#the-input-element)
  elements that [support a
  picker](input.html#input-support-picker){#pseudo-classes:input-support-picker}
  and whose pickers are open

This specification does not define when an element matches the `:lang()`
dynamic
[pseudo-class](https://drafts.csswg.org/selectors/#pseudo-class){#pseudo-classes:pseudo-class-41
x-internal="pseudo-class"}, as it is defined in sufficient detail in a
language-agnostic fashion in Selectors.
[\[SELECTORS\]](references.html#refsSELECTORS)

[← 4.13 Custom elements](custom-elements.html) --- [Table of
Contents](./) --- [5 Microdata →](microdata.html)
