## [5]{.secno} [Microdata]{.dfn}[](#microdata){.self-link}

### [5.1]{.secno} Introduction[](#introduction-7){.self-link} {#introduction-7}

#### [5.1.1]{.secno} Overview[](#overview){.self-link}

*This section is non-normative.*

Sometimes, it is desirable to annotate content with specific
machine-readable labels, e.g. to allow generic scripts to provide
services that are customized to the page, or to enable content from a
variety of cooperating authors to be processed by a single script in a
consistent manner.

For this purpose, authors can use the microdata features described in
this section. Microdata allows nested groups of name-value pairs to be
added to documents, in parallel with the existing content.

#### [5.1.2]{.secno} The basic syntax[](#the-basic-syntax){.self-link}

*This section is non-normative.*

At a high level, microdata consists of a group of name-value pairs. The
groups are called
[items](#concept-item){#the-basic-syntax:concept-item}, and each
name-value pair is a property. Items and properties are represented by
regular elements.

To create an item, the
[`itemscope`{#the-basic-syntax:attr-itemscope}](#attr-itemscope)
attribute is used.

To add a property to an item, the
[`itemprop`{#the-basic-syntax:names:-the-itemprop-attribute}](#names:-the-itemprop-attribute)
attribute is used on one of the
[item\'s](#concept-item){#the-basic-syntax:concept-item-2} descendants.

::: example
Here there are two items, each of which has the property \"name\":

``` html
<div itemscope>
 <p>My name is <span itemprop="name">Elizabeth</span>.</p>
</div>

<div itemscope>
 <p>My name is <span itemprop="name">Daniel</span>.</p>
</div>
```
:::

Markup without the microdata-related attributes does not have any effect
on the microdata model.

::: example
These two examples are exactly equivalent, at a microdata level, as the
previous two examples respectively:

``` html
<div itemscope>
 <p>My <em>name</em> is <span itemprop="name">E<strong>liz</strong>abeth</span>.</p>
</div>

<section>
 <div itemscope>
  <aside>
   <p>My name is <span itemprop="name"><a href="/?user=daniel">Daniel</a></span>.</p>
  </aside>
 </div>
</section>
```
:::

Properties generally have values that are strings.

::: example
Here the item has three properties:

``` html
<div itemscope>
 <p>My name is <span itemprop="name">Neil</span>.</p>
 <p>My band is called <span itemprop="band">Four Parts Water</span>.</p>
 <p>I am <span itemprop="nationality">British</span>.</p>
</div>
```
:::

When a string value is a
[URL](https://url.spec.whatwg.org/#concept-url){#the-basic-syntax:url
x-internal="url"}, it is expressed using the
[`a`{#the-basic-syntax:the-a-element}](text-level-semantics.html#the-a-element)
element and its
[`href`{#the-basic-syntax:attr-hyperlink-href}](links.html#attr-hyperlink-href)
attribute, the
[`img`{#the-basic-syntax:the-img-element}](embedded-content.html#the-img-element)
element and its
[`src`{#the-basic-syntax:attr-img-src}](embedded-content.html#attr-img-src)
attribute, or other elements that link to or embed external resources.

::: example
In this example, the item has one property, \"image\", whose value is a
URL:

``` html
<div itemscope>
 <img itemprop="image" src="google-logo.png" alt="Google">
</div>
```
:::

When a string value is in some machine-readable format unsuitable for
human consumption, it is expressed using the
[`value`{#the-basic-syntax:attr-data-value}](text-level-semantics.html#attr-data-value)
attribute of the
[`data`{#the-basic-syntax:the-data-element}](text-level-semantics.html#the-data-element)
element, with the human-readable version given in the element\'s
contents.

::: example
Here, there is an item with a property whose value is a product ID. The
ID is not human-friendly, so the product\'s name is used the
human-visible text instead of the ID.

``` html
<h1 itemscope>
 <data itemprop="product-id" value="9678AOU879">The Instigator 2000</data>
</h1>
```
:::

For numeric data, the
[`meter`{#the-basic-syntax:the-meter-element}](form-elements.html#the-meter-element)
element and its
[`value`{#the-basic-syntax:attr-meter-value}](form-elements.html#attr-meter-value)
attribute can be used instead.

::: example
Here a rating is given using a
[`meter`{#the-basic-syntax:the-meter-element-2}](form-elements.html#the-meter-element)
element.

``` html
<div itemscope itemtype="http://schema.org/Product">
 <span itemprop="name">Panasonic White 60L Refrigerator</span>
 <img src="panasonic-fridge-60l-white.jpg" alt="">
  <div itemprop="aggregateRating"
       itemscope itemtype="http://schema.org/AggregateRating">
   <meter itemprop="ratingValue" min=0 value=3.5 max=5>Rated 3.5/5</meter>
   (based on <span itemprop="reviewCount">11</span> customer reviews)
  </div>
</div>
```
:::

Similarly, for date- and time-related data, the
[`time`{#the-basic-syntax:the-time-element}](text-level-semantics.html#the-time-element)
element and its
[`datetime`{#the-basic-syntax:attr-time-datetime}](text-level-semantics.html#attr-time-datetime)
attribute can be used instead.

::: example
In this example, the item has one property, \"birthday\", whose value is
a date:

``` html
<div itemscope>
 I was born on <time itemprop="birthday" datetime="2009-05-10">May 10th 2009</time>.
</div>
```
:::

Properties can also themselves be groups of name-value pairs, by putting
the [`itemscope`{#the-basic-syntax:attr-itemscope-2}](#attr-itemscope)
attribute on the element that declares the property.

Items that are not part of others are called [top-level microdata
items](#top-level-microdata-items){#the-basic-syntax:top-level-microdata-items}.

::: example
In this example, the outer item represents a person, and the inner one
represents a band:

``` html
<div itemscope>
 <p>Name: <span itemprop="name">Amanda</span></p>
 <p>Band: <span itemprop="band" itemscope> <span itemprop="name">Jazz Band</span> (<span itemprop="size">12</span> players)</span></p>
</div>
```

The outer item here has two properties, \"name\" and \"band\". The
\"name\" is \"Amanda\", and the \"band\" is an item in its own right,
with two properties, \"name\" and \"size\". The \"name\" of the band is
\"Jazz Band\", and the \"size\" is \"12\".

The outer item in this example is a top-level microdata item.
:::

Properties that are not descendants of the element with the
[`itemscope`{#the-basic-syntax:attr-itemscope-3}](#attr-itemscope)
attribute can be associated with the
[item](#concept-item){#the-basic-syntax:concept-item-3} using the
[`itemref`{#the-basic-syntax:attr-itemref}](#attr-itemref) attribute.
This attribute takes a list of IDs of elements to crawl in addition to
crawling the children of the element with the
[`itemscope`{#the-basic-syntax:attr-itemscope-4}](#attr-itemscope)
attribute.

::: example
This example is the same as the previous one, but all the properties are
separated from their
[items](#concept-item){#the-basic-syntax:concept-item-4}:

``` html
<div itemscope id="amanda" itemref="a b"></div>
<p id="a">Name: <span itemprop="name">Amanda</span></p>
<div id="b" itemprop="band" itemscope itemref="c"></div>
<div id="c">
 <p>Band: <span itemprop="name">Jazz Band</span></p>
 <p>Size: <span itemprop="size">12</span> players</p>
</div>
```

This gives the same result as the previous example. The first item has
two properties, \"name\", set to \"Amanda\", and \"band\", set to
another item. That second item has two further properties, \"name\", set
to \"Jazz Band\", and \"size\", set to \"12\".
:::

An [item](#concept-item){#the-basic-syntax:concept-item-5} can have
multiple properties with the same name and different values.

::: example
This example describes an ice cream, with two flavors:

``` html
<div itemscope>
 <p>Flavors in my favorite ice cream:</p>
 <ul>
  <li itemprop="flavor">Lemon sorbet</li>
  <li itemprop="flavor">Apricot sorbet</li>
 </ul>
</div>
```

This thus results in an item with two properties, both \"flavor\",
having the values \"Lemon sorbet\" and \"Apricot sorbet\".
:::

An element introducing a property can also introduce multiple properties
at once, to avoid duplication when some of the properties have the same
value.

::: example
Here we see an item with two properties, \"favorite-color\" and
\"favorite-fruit\", both set to the value \"orange\":

``` html
<div itemscope>
 <span itemprop="favorite-color favorite-fruit">orange</span>
</div>
```
:::

It\'s important to note that there is no relationship between the
microdata and the content of the document where the microdata is marked
up.

::: example
There is no semantic difference, for instance, between the following two
examples:

``` html
<figure>
 <img src="castle.jpeg">
 <figcaption><span itemscope><span itemprop="name">The Castle</span></span> (1986)</figcaption>
</figure>
```

``` html
<span itemscope><meta itemprop="name" content="The Castle"></span>
<figure>
 <img src="castle.jpeg">
 <figcaption>The Castle (1986)</figcaption>
</figure>
```

Both have a figure with a caption, and both, completely unrelated to the
figure, have an item with a name-value pair with the name \"name\" and
the value \"The Castle\". The only difference is that if the user drags
the caption out of the document, in the former case, the item will be
included in the drag-and-drop data. In neither case is the image in any
way associated with the item.
:::

#### [5.1.3]{.secno} Typed items[](#typed-items){.self-link}

*This section is non-normative.*

The examples in the previous section show how information could be
marked up on a page that doesn\'t expect its microdata to be re-used.
Microdata is most useful, though, when it is used in contexts where
other authors and readers are able to cooperate to make new uses of the
markup.

For this purpose, it is necessary to give each
[item](#concept-item){#typed-items:concept-item} a type, such as
\"https://example.com/person\", or \"https://example.org/cat\", or
\"https://band.example.net/\". Types are identified as
[URLs](https://url.spec.whatwg.org/#concept-url){#typed-items:url
x-internal="url"}.

The type for an [item](#concept-item){#typed-items:concept-item-2} is
given as the value of an
[`itemtype`{#typed-items:attr-itemtype}](#attr-itemtype) attribute on
the same element as the
[`itemscope`{#typed-items:attr-itemscope}](#attr-itemscope) attribute.

::: example
Here, the item\'s type is \"https://example.org/animals#cat\":

``` html
<section itemscope itemtype="https://example.org/animals#cat">
 <h1 itemprop="name">Hedral</h1>
 <p itemprop="desc">Hedral is a male american domestic
 shorthair, with a fluffy black fur with white paws and belly.</p>
 <img itemprop="img" src="hedral.jpeg" alt="" title="Hedral, age 18 months">
</section>
```

In this example the \"https://example.org/animals#cat\" item has three
properties, a \"name\" (\"Hedral\"), a \"desc\" (\"Hedral is\...\"), and
an \"img\" (\"hedral.jpeg\").
:::

The type gives the context for the properties, thus selecting a
vocabulary: a property named \"class\" given for an item with the type
\"https://census.example/person\" might refer to the economic class of
an individual, while a property named \"class\" given for an item with
the type \"https://example.com/school/teacher\" might refer to the
classroom a teacher has been assigned. Several types can share a
vocabulary. For example, the types
\"`https://example.org/people/teacher`\" and
\"`https://example.org/people/engineer`\" could be defined to use the
same vocabulary (though maybe some properties would not be especially
useful in both cases, e.g. maybe the
\"`https://example.org/people/engineer`\" type might not typically be
used with the \"`classroom`\" property). Multiple types defined to use
the same vocabulary can be given for a single item by listing the URLs
as a space-separated list in the attribute\' value. An item cannot be
given two types if they do not use the same vocabulary, however.

#### [5.1.4]{.secno} Global identifiers for items[](#global-identifiers-for-items){.self-link}

*This section is non-normative.*

Sometimes, an
[item](#concept-item){#global-identifiers-for-items:concept-item} gives
information about a topic that has a global identifier. For example,
books can be identified by their ISBN number.

Vocabularies (as identified by the
[`itemtype`{#global-identifiers-for-items:attr-itemtype}](#attr-itemtype)
attribute) can be designed such that
[items](#concept-item){#global-identifiers-for-items:concept-item-2} get
associated with their global identifier in an unambiguous way by
expressing the global identifiers as
[URLs](https://url.spec.whatwg.org/#concept-url){#global-identifiers-for-items:url
x-internal="url"} given in an
[`itemid`{#global-identifiers-for-items:attr-itemid}](#attr-itemid)
attribute.

The exact meaning of the
[URLs](https://url.spec.whatwg.org/#concept-url){#global-identifiers-for-items:url-2
x-internal="url"} given in
[`itemid`{#global-identifiers-for-items:attr-itemid-2}](#attr-itemid)
attributes depends on the vocabulary used.

::: example
Here, an item is talking about a particular book:

``` html
<dl itemscope
    itemtype="https://vocab.example.net/book"
    itemid="urn:isbn:0-330-34032-8">
 <dt>Title
 <dd itemprop="title">The Reality Dysfunction
 <dt>Author
 <dd itemprop="author">Peter F. Hamilton
 <dt>Publication date
 <dd><time itemprop="pubdate" datetime="1996-01-26">26 January 1996</time>
</dl>
```

The \"`https://vocab.example.net/book`\" vocabulary in this example
would define that the
[`itemid`{#global-identifiers-for-items:attr-itemid-3}](#attr-itemid)
attribute takes a
[`urn:`{#global-identifiers-for-items:urn-protocol}](https://www.rfc-editor.org/rfc/rfc2141#section-2){x-internal="urn-protocol"}
[URL](https://url.spec.whatwg.org/#concept-url){#global-identifiers-for-items:url-3
x-internal="url"} pointing to the ISBN of the book.
:::

#### [5.1.5]{.secno} Selecting names when defining vocabularies[](#selecting-names-when-defining-vocabularies){.self-link}

*This section is non-normative.*

Using microdata means using a vocabulary. For some purposes, an ad-hoc
vocabulary is adequate. For others, a vocabulary will need to be
designed. Where possible, authors are encouraged to re-use existing
vocabularies, as this makes content re-use easier.

When designing new vocabularies, identifiers can be created either using
[URLs](https://url.spec.whatwg.org/#concept-url){#selecting-names-when-defining-vocabularies:url
x-internal="url"}, or, for properties, as plain words (with no dots or
colons). For URLs, conflicts with other vocabularies can be avoided by
only using identifiers that correspond to pages that the author has
control over.

::: example
For instance, if Jon and Adam both write content at `example.com`, at
`https://example.com/~jon/...` and `https://example.com/~adam/...`
respectively, then they could select identifiers of the form
\"https://example.com/\~jon/name\" and
\"https://example.com/\~adam/name\" respectively.
:::

Properties whose names are just plain words can only be used within the
context of the types for which they are intended; properties named using
URLs can be reused in items of any type. If an item has no type, and is
not part of another item, then if its properties have names that are
just plain words, they are not intended to be globally unique, and are
instead only intended for limited use. Generally speaking, authors are
encouraged to use either properties with globally unique names (URLs) or
ensure that their items are typed.

::: example
Here, an item is an \"https://example.org/animals#cat\", and most of the
properties have names that are words defined in the context of that
type. There are also a few additional properties whose names come from
other vocabularies.

``` html
<section itemscope itemtype="https://example.org/animals#cat">
 <h1 itemprop="name https://example.com/fn">Hedral</h1>
 <p itemprop="desc">Hedral is a male American domestic
 shorthair, with a fluffy <span
 itemprop="https://example.com/color">black</span> fur with <span
 itemprop="https://example.com/color">white</span> paws and belly.</p>
 <img itemprop="img" src="hedral.jpeg" alt="" title="Hedral, age 18 months">
</section>
```

This example has one item with the type
\"https://example.org/animals#cat\" and the following properties:

Property

Value

name

Hedral

https://example.com/fn

Hedral

desc

Hedral is a male American domestic shorthair, with a fluffy black fur
with white paws and belly.

https://example.com/color

black

https://example.com/color

white

img

\.../hedral.jpeg
:::

### [5.2]{.secno} Encoding microdata[](#encoding-microdata){.self-link}

#### [5.2.1]{.secno} The microdata model[](#the-microdata-model){.self-link}

The microdata model consists of groups of name-value pairs known as
[items](#concept-item){#the-microdata-model:concept-item}.

Each group is known as an
[item](#concept-item){#the-microdata-model:concept-item-2}. Each
[item](#concept-item){#the-microdata-model:concept-item-3} can have
[item types](#item-types){#the-microdata-model:item-types}, a [global
identifier](#global-identifier){#the-microdata-model:global-identifier}
(if the vocabulary specified by the [item
types](#item-types){#the-microdata-model:item-types-2} [support global
identifiers for
items](#support-global-identifiers-for-items){#the-microdata-model:support-global-identifiers-for-items}),
and a list of name-value pairs. Each name in the name-value pair is
known as a
[property](#the-properties-of-an-item){#the-microdata-model:the-properties-of-an-item},
and each
[property](#the-properties-of-an-item){#the-microdata-model:the-properties-of-an-item-2}
has one or more
[values](#concept-property-value){#the-microdata-model:concept-property-value}.
Each
[value](#concept-property-value){#the-microdata-model:concept-property-value-2}
is either a string or itself a group of name-value pairs (an
[item](#concept-item){#the-microdata-model:concept-item-4}). The names
are unordered relative to each other, but if a particular name has
multiple values, they do have a relative order.

#### [5.2.2]{.secno} Items[](#items){.self-link}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Global_attributes/itemscope](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/itemscope "itemscope is a boolean global attribute that defines the scope of associated metadata. Specifying the itemscope attribute for an element creates a new item, which results in a number of name-value pairs that are associated with the element.")

Support in all current engines.

::: support
[FirefoxYes]{.firefox .yes}[SafariYes]{.safari .yes}[ChromeYes]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[EdgeYes]{.edge_blink .yes}

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

Every [HTML
element](infrastructure.html#html-elements){#items:html-elements} may
have an [`itemscope`]{#attr-itemscope .dfn dfn-for="html-global"
dfn-type="element-attr"} attribute specified. The
[`itemscope`{#items:attr-itemscope}](#attr-itemscope) attribute is a
[boolean
attribute](common-microsyntaxes.html#boolean-attribute){#items:boolean-attribute}.

An element with the
[`itemscope`{#items:attr-itemscope-2}](#attr-itemscope) attribute
specified creates a new [item]{#concept-item .dfn}, a group of
name-value pairs.

------------------------------------------------------------------------

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Global_attributes/itemtype](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/itemtype "The global attribute itemtype specifies the URL of the vocabulary that will be used to define itemprop's (item properties) in the data structure.")

Support in all current engines.

::: support
[FirefoxYes]{.firefox .yes}[SafariYes]{.safari .yes}[ChromeYes]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[EdgeYes]{.edge_blink .yes}

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

Elements with an [`itemscope`{#items:attr-itemscope-3}](#attr-itemscope)
attribute may have an [`itemtype`]{#attr-itemtype .dfn
dfn-for="html-global" dfn-type="element-attr"} attribute specified, to
give the [item types](#item-types){#items:item-types} of the
[item](#concept-item){#items:concept-item}.

The [`itemtype`{#items:attr-itemtype}](#attr-itemtype) attribute, if
specified, must have a value that is an [unordered set of unique
space-separated
tokens](common-microsyntaxes.html#unordered-set-of-unique-space-separated-tokens){#items:unordered-set-of-unique-space-separated-tokens},
none of which are [identical
to](https://infra.spec.whatwg.org/#string-is){#items:identical-to
x-internal="identical-to"} another token and each of which is a [valid
URL
string](https://url.spec.whatwg.org/#valid-url-string){#items:valid-url-string
x-internal="valid-url-string"} that is an [absolute
URL](https://url.spec.whatwg.org/#syntax-url-absolute){#items:absolute-url
x-internal="absolute-url"}, and all of which are defined to use the same
vocabulary. The attribute\'s value must have at least one token.

The [item types]{#item-types .dfn} of an
[item](#concept-item){#items:concept-item-2} are the tokens obtained by
[splitting the element\'s `itemtype` attribute\'s value on ASCII
whitespace](https://infra.spec.whatwg.org/#split-on-ascii-whitespace){#items:split-a-string-on-spaces
x-internal="split-a-string-on-spaces"}. If the
[`itemtype`{#items:attr-itemtype-2}](#attr-itemtype) attribute is
missing or parsing it in this way finds no tokens, the
[item](#concept-item){#items:concept-item-3} is said to have no [item
types](#item-types){#items:item-types-2}.

The [item types](#item-types){#items:item-types-3} must all be types
defined in [applicable
specifications](infrastructure.html#other-applicable-specifications){#items:other-applicable-specifications}
and must all be defined to use the same vocabulary.

Except if otherwise specified by that specification, the
[URLs](https://url.spec.whatwg.org/#concept-url){#items:url
x-internal="url"} given as the [item
types](#item-types){#items:item-types-4} should not be automatically
dereferenced.

A specification could define that its [item
type](#item-types){#items:item-types-5} can be dereferenced to provide
the user with help information, for example. In fact, vocabulary authors
are encouraged to provide useful information at the given
[URL](https://url.spec.whatwg.org/#concept-url){#items:url-2
x-internal="url"}.

[Item types](#item-types){#items:item-types-6} are opaque identifiers,
and user agents must not dereference unknown [item
types](#item-types){#items:item-types-7}, or otherwise deconstruct them,
in order to determine how to process
[items](#concept-item){#items:concept-item-4} that use them.

The [`itemtype`{#items:attr-itemtype-3}](#attr-itemtype) attribute must
not be specified on elements that do not have an
[`itemscope`{#items:attr-itemscope-4}](#attr-itemscope) attribute
specified.

------------------------------------------------------------------------

An [item](#concept-item){#items:concept-item-5} is said to be a [typed
item]{#typed-item .dfn} when either it has an [item
type](#item-types){#items:item-types-8}, or it is the
[value](#concept-property-value){#items:concept-property-value} of a
[property](#the-properties-of-an-item){#items:the-properties-of-an-item}
of a [typed item](#typed-item){#items:typed-item}. The [relevant
types]{#relevant-types .dfn} for a [typed
item](#typed-item){#items:typed-item-2} is the
[item](#concept-item){#items:concept-item-6}\'s [item
types](#item-types){#items:item-types-9}, if it has any, or else is the
[relevant types](#relevant-types){#items:relevant-types} of the
[item](#concept-item){#items:concept-item-7} for which it is a
[property](#the-properties-of-an-item){#items:the-properties-of-an-item-2}\'s
[value](#concept-property-value){#items:concept-property-value-2}.

------------------------------------------------------------------------

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Global_attributes/itemid](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/itemid "The itemid global attribute provides microdata in the form of a unique, global identifier of an item.")

Support in all current engines.

::: support
[FirefoxYes]{.firefox .yes}[SafariYes]{.safari .yes}[ChromeYes]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[EdgeYes]{.edge_blink .yes}

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

Elements with an [`itemscope`{#items:attr-itemscope-5}](#attr-itemscope)
attribute and an [`itemtype`{#items:attr-itemtype-4}](#attr-itemtype)
attribute that references a vocabulary that is defined to [support
global identifiers for items]{#support-global-identifiers-for-items
.dfn} may also have an [`itemid`]{#attr-itemid .dfn
dfn-for="html-global" dfn-type="element-attr"} attribute specified, to
give a global identifier for the
[item](#concept-item){#items:concept-item-8}, so that it can be related
to other [items](#concept-item){#items:concept-item-9} on pages
elsewhere on the web.

The [`itemid`{#items:attr-itemid}](#attr-itemid) attribute, if
specified, must have a value that is a [valid URL potentially surrounded
by
spaces](urls-and-fetching.html#valid-url-potentially-surrounded-by-spaces){#items:valid-url-potentially-surrounded-by-spaces}.

The [global identifier]{#global-identifier .dfn} of an
[item](#concept-item){#items:concept-item-10} is the value of its
element\'s [`itemid`{#items:attr-itemid-2}](#attr-itemid) attribute, if
it has one,
[parsed](urls-and-fetching.html#encoding-parsing-a-url){#items:encoding-parsing-a-url}
relative to the [node
document](https://dom.spec.whatwg.org/#concept-node-document){#items:node-document
x-internal="node-document"} of the element on which the attribute is
specified. If the [`itemid`{#items:attr-itemid-3}](#attr-itemid)
attribute is missing or if parsing it returns failure, it is said to
have no [global
identifier](#global-identifier){#items:global-identifier}.

The [`itemid`{#items:attr-itemid-4}](#attr-itemid) attribute must not be
specified on elements that do not have both an
[`itemscope`{#items:attr-itemscope-6}](#attr-itemscope) attribute and an
[`itemtype`{#items:attr-itemtype-5}](#attr-itemtype) attribute
specified, and must not be specified on elements with an
[`itemscope`{#items:attr-itemscope-7}](#attr-itemscope) attribute whose
[`itemtype`{#items:attr-itemtype-6}](#attr-itemtype) attribute specifies
a vocabulary that does not [support global identifiers for
items](#support-global-identifiers-for-items){#items:support-global-identifiers-for-items},
as defined by that vocabulary\'s specification.

The exact meaning of a [global
identifier](#global-identifier){#items:global-identifier-2} is
determined by the vocabulary\'s specification. It is up to such
specifications to define whether multiple items with the same global
identifier (whether on the same page or on different pages) are allowed
to exist, and what the processing rules for that vocabulary are with
respect to handling the case of multiple items with the same ID.

------------------------------------------------------------------------

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Global_attributes/itemref](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/itemref "Properties that are not descendants of an element with the itemscope attribute can be associated with an item using the global attribute itemref.")

Support in all current engines.

::: support
[FirefoxYes]{.firefox .yes}[SafariYes]{.safari .yes}[ChromeYes]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[EdgeYes]{.edge_blink .yes}

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

Elements with an [`itemscope`{#items:attr-itemscope-8}](#attr-itemscope)
attribute may have an [`itemref`]{#attr-itemref .dfn
dfn-for="html-global" dfn-type="element-attr"} attribute specified, to
give a list of additional elements to crawl to find the name-value pairs
of the [item](#concept-item){#items:concept-item-11}.

The [`itemref`{#items:attr-itemref}](#attr-itemref) attribute, if
specified, must have a value that is an [unordered set of unique
space-separated
tokens](common-microsyntaxes.html#unordered-set-of-unique-space-separated-tokens){#items:unordered-set-of-unique-space-separated-tokens-2}
none of which are [identical
to](https://infra.spec.whatwg.org/#string-is){#items:identical-to-2
x-internal="identical-to"} another token and consisting of
[IDs](https://dom.spec.whatwg.org/#concept-id){#items:concept-id
x-internal="concept-id"} of elements in the same
[tree](https://dom.spec.whatwg.org/#concept-tree){#items:tree
x-internal="tree"}.

The [`itemref`{#items:attr-itemref-2}](#attr-itemref) attribute must not
be specified on elements that do not have an
[`itemscope`{#items:attr-itemscope-9}](#attr-itemscope) attribute
specified.

The [`itemref`{#items:attr-itemref-3}](#attr-itemref) attribute is not
part of the microdata data model. It is merely a syntactic construct to
aid authors in adding annotations to pages where the data to be
annotated does not follow a convenient tree structure. For example, it
allows authors to mark up data in a table so that each column defines a
separate [item](#concept-item){#items:concept-item-12}, while keeping
the properties in the cells.

::: example
This example shows a simple vocabulary used to describe the products of
a model railway manufacturer. The vocabulary has just five property
names:

product-code
:   An integer that names the product in the manufacturer\'s catalog.

name
:   A brief description of the product.

scale
:   One of \"HO\", \"1\", or \"Z\" (potentially with leading or trailing
    whitespace), indicating the scale of the product.

digital
:   If present, one of \"Digital\", \"Delta\", or \"Systems\"
    (potentially with leading or trailing whitespace) indicating that
    the product has a digital decoder of the given type.

track-type
:   For track-specific products, one of \"K\", \"M\", \"C\" (potentially
    with leading or trailing whitespace) indicating the type of track
    for which the product is intended.

This vocabulary has four defined [item
types](#item-types){#items:item-types-10}:

https://md.example.com/loco
:   Rolling stock with an engine

https://md.example.com/passengers
:   Passenger rolling stock

https://md.example.com/track
:   Track pieces

https://md.example.com/lighting
:   Equipment with lighting

Each [item](#concept-item){#items:concept-item-13} that uses this
vocabulary can be given one or more of these types, depending on what
the product is.

Thus, a locomotive might be marked up as:

``` html
<dl itemscope itemtype="https://md.example.com/loco
                        https://md.example.com/lighting">
 <dt>Name:
 <dd itemprop="name">Tank Locomotive (DB 80)
 <dt>Product code:
 <dd itemprop="product-code">33041
 <dt>Scale:
 <dd itemprop="scale">HO
 <dt>Digital:
 <dd itemprop="digital">Delta
</dl>
```

A turnout lantern retrofit kit might be marked up as:

``` html
<dl itemscope itemtype="https://md.example.com/track
                        https://md.example.com/lighting">
 <dt>Name:
 <dd itemprop="name">Turnout Lantern Kit
 <dt>Product code:
 <dd itemprop="product-code">74470
 <dt>Purpose:
 <dd>For retrofitting 2 <span itemprop="track-type">C</span> Track
 turnouts. <meta itemprop="scale" content="HO">
</dl>
```

A passenger car with no lighting might be marked up as:

``` html
<dl itemscope itemtype="https://md.example.com/passengers">
 <dt>Name:
 <dd itemprop="name">Express Train Passenger Car (DB Am 203)
 <dt>Product code:
 <dd itemprop="product-code">8710
 <dt>Scale:
 <dd itemprop="scale">Z
</dl>
```

Great care is necessary when creating new vocabularies. Often, a
hierarchical approach to types can be taken that results in a vocabulary
where each item only ever has a single type, which is generally much
simpler to manage.
:::

#### [5.2.3]{.secno} Names: the [`itemprop`]{.dfn dfn-for="html-global" dfn-type="element-attr"} attribute[](#names:-the-itemprop-attribute){.self-link} {#names:-the-itemprop-attribute}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Global_attributes/itemprop](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/itemprop "The itemprop global attribute is used to add properties to an item. Every HTML element can have an itemprop attribute specified, and an itemprop consists of a name-value pair. Each name-value pair is called a property, and a group of one or more properties forms an item. Property values are either a string or a URL and can be associated with a very wide range of elements including <audio>, <embed>, <iframe>, <img>, <link>, <object>, <source>, <track>, and <video>.")

Support in all current engines.

::: support
[FirefoxYes]{.firefox .yes}[SafariYes]{.safari .yes}[ChromeYes]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[EdgeYes]{.edge_blink .yes}

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

Every [HTML
element](infrastructure.html#html-elements){#names:-the-itemprop-attribute:html-elements}
may have an
[`itemprop`{#names:-the-itemprop-attribute:names:-the-itemprop-attribute}](#names:-the-itemprop-attribute)
attribute specified, if doing so [adds one or more
properties](#the-properties-of-an-item){#names:-the-itemprop-attribute:the-properties-of-an-item}
to one or more
[items](#concept-item){#names:-the-itemprop-attribute:concept-item} (as
defined below).

The
[`itemprop`{#names:-the-itemprop-attribute:names:-the-itemprop-attribute-2}](#names:-the-itemprop-attribute)
attribute, if specified, must have a value that is an [unordered set of
unique space-separated
tokens](common-microsyntaxes.html#unordered-set-of-unique-space-separated-tokens){#names:-the-itemprop-attribute:unordered-set-of-unique-space-separated-tokens}
none of which are [identical
to](https://infra.spec.whatwg.org/#string-is){#names:-the-itemprop-attribute:identical-to
x-internal="identical-to"} another token, representing the names of the
name-value pairs that it adds. The attribute\'s value must have at least
one token.

Each token must be either:

- If the item is a [typed
  item](#typed-item){#names:-the-itemprop-attribute:typed-item}: a
  [defined property name]{#defined-property-name .dfn} allowed in this
  situation according to the specification that defines the [relevant
  types](#relevant-types){#names:-the-itemprop-attribute:relevant-types}
  for the item, or
- A [valid URL
  string](https://url.spec.whatwg.org/#valid-url-string){#names:-the-itemprop-attribute:valid-url-string
  x-internal="valid-url-string"} that is an [absolute
  URL](https://url.spec.whatwg.org/#syntax-url-absolute){#names:-the-itemprop-attribute:absolute-url
  x-internal="absolute-url"} defined as an item property name allowed in
  this situation by a vocabulary specification, or
- A [valid URL
  string](https://url.spec.whatwg.org/#valid-url-string){#names:-the-itemprop-attribute:valid-url-string-2
  x-internal="valid-url-string"} that is an [absolute
  URL](https://url.spec.whatwg.org/#syntax-url-absolute){#names:-the-itemprop-attribute:absolute-url-2
  x-internal="absolute-url"}, used as a proprietary item property name
  (i.e. one used by the author for private purposes, not defined in a
  public specification), or
- If the item is not a [typed
  item](#typed-item){#names:-the-itemprop-attribute:typed-item-2}: a
  string that contains no U+002E FULL STOP characters (.) and no U+003A
  COLON characters (:), used as a proprietary item property name (i.e.
  one used by the author for private purposes, not defined in a public
  specification).

Specifications that introduce [defined property
names](#defined-property-name){#names:-the-itemprop-attribute:defined-property-name}
must ensure all such property names contain no U+002E FULL STOP
characters (.), no U+003A COLON characters (:), and no [ASCII
whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#names:-the-itemprop-attribute:space-characters
x-internal="space-characters"}.

The rules above disallow U+003A COLON characters (:) in non-URL values
because otherwise they could not be distinguished from URLs. Values with
U+002E FULL STOP characters (.) are reserved for future extensions.
[ASCII
whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#names:-the-itemprop-attribute:space-characters-2
x-internal="space-characters"} are disallowed because otherwise the
values would be parsed as multiple tokens.

When an element with an
[`itemprop`{#names:-the-itemprop-attribute:names:-the-itemprop-attribute-3}](#names:-the-itemprop-attribute)
attribute [adds a
property](#the-properties-of-an-item){#names:-the-itemprop-attribute:the-properties-of-an-item-2}
to multiple
[items](#concept-item){#names:-the-itemprop-attribute:concept-item-2},
the requirement above regarding the tokens applies for each
[item](#concept-item){#names:-the-itemprop-attribute:concept-item-3}
individually.

The [property names]{#property-names .dfn} of an element are the tokens
that the element\'s
[`itemprop`{#names:-the-itemprop-attribute:names:-the-itemprop-attribute-4}](#names:-the-itemprop-attribute)
attribute is found to contain when its value is [split on ASCII
whitespace](https://infra.spec.whatwg.org/#split-on-ascii-whitespace){#names:-the-itemprop-attribute:split-a-string-on-spaces
x-internal="split-a-string-on-spaces"}, with the order preserved but
with duplicates removed (leaving only the first occurrence of each
name).

Within an
[item](#concept-item){#names:-the-itemprop-attribute:concept-item-4},
the properties are unordered with respect to each other, except for
properties with the same name, which are ordered in the order they are
given by the algorithm that defines [the properties of an
item](#the-properties-of-an-item){#names:-the-itemprop-attribute:the-properties-of-an-item-3}.

::: example
In the following example, the \"a\" property has the values \"1\" and
\"2\", *in that order*, but whether the \"a\" property comes before the
\"b\" property or not is not important:

``` html
<div itemscope>
 <p itemprop="a">1</p>
 <p itemprop="a">2</p>
 <p itemprop="b">test</p>
</div>
```

Thus, the following is equivalent:

``` html
<div itemscope>
 <p itemprop="b">test</p>
 <p itemprop="a">1</p>
 <p itemprop="a">2</p>
</div>
```

As is the following:

``` html
<div itemscope>
 <p itemprop="a">1</p>
 <p itemprop="b">test</p>
 <p itemprop="a">2</p>
</div>
```

And the following:

``` html
<div id="x">
 <p itemprop="a">1</p>
</div>
<div itemscope itemref="x">
 <p itemprop="b">test</p>
 <p itemprop="a">2</p>
</div>
```
:::

#### [5.2.4]{.secno} Values[](#values){.self-link}

The [property value]{#concept-property-value .dfn} of a name-value pair
added by an element with an
[`itemprop`{#values:names:-the-itemprop-attribute}](#names:-the-itemprop-attribute)
attribute is as given for the first matching case in the following list:

If the element also has an [`itemscope`{#values:attr-itemscope}](#attr-itemscope) attribute

:   The value is the [item](#concept-item){#values:concept-item} created
    by the element.

If the element is a [`meta`{#values:the-meta-element}](semantics.html#the-meta-element) element

:   The value is the value of the element\'s
    [`content`{#values:attr-meta-content}](semantics.html#attr-meta-content)
    attribute, if any, or the empty string if there is no such
    attribute.

If the element is an [`audio`{#values:the-audio-element}](media.html#the-audio-element), [`embed`{#values:the-embed-element}](iframe-embed-object.html#the-embed-element), [`iframe`{#values:the-iframe-element}](iframe-embed-object.html#the-iframe-element), [`img`{#values:the-img-element}](embedded-content.html#the-img-element), [`source`{#values:the-source-element}](embedded-content.html#the-source-element), [`track`{#values:the-track-element}](media.html#the-track-element), or [`video`{#values:the-video-element}](media.html#the-video-element) element

:   The value is the result of [encoding-parsing-and-serializing a
    URL](urls-and-fetching.html#encoding-parsing-and-serializing-a-url){#values:encoding-parsing-and-serializing-a-url}
    given the element\'s `src` attribute\'s value, relative to the
    element\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#values:node-document
    x-internal="node-document"}, at the time the attribute is set, or
    the empty string if there is no such attribute or the result is
    failure.

If the element is an [`a`{#values:the-a-element}](text-level-semantics.html#the-a-element), [`area`{#values:the-area-element}](image-maps.html#the-area-element), or [`link`{#values:the-link-element}](semantics.html#the-link-element) element

:   The value is the result of [encoding-parsing-and-serializing a
    URL](urls-and-fetching.html#encoding-parsing-and-serializing-a-url){#values:encoding-parsing-and-serializing-a-url-2}
    given the element\'s `href` attribute\'s value, relative to the
    element\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#values:node-document-2
    x-internal="node-document"}, at the time the attribute is set, or
    the empty string if there is no such attribute or the result is
    failure.

If the element is an [`object`{#values:the-object-element}](iframe-embed-object.html#the-object-element) element

:   The value is the result of [encoding-parsing-and-serializing a
    URL](urls-and-fetching.html#encoding-parsing-and-serializing-a-url){#values:encoding-parsing-and-serializing-a-url-3}
    given the element\'s `data` attribute\'s value, relative to the
    element\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#values:node-document-3
    x-internal="node-document"}, at the time the attribute is set, or
    the empty string if there is no such attribute or the result is
    failure.

If the element is a [`data`{#values:the-data-element}](text-level-semantics.html#the-data-element) element

:   The value is the value of the element\'s
    [`value`{#values:attr-data-value}](text-level-semantics.html#attr-data-value)
    attribute, if it has one, or the empty string otherwise.

If the element is a [`meter`{#values:the-meter-element}](form-elements.html#the-meter-element) element

:   The value is the value of the element\'s
    [`value`{#values:attr-meter-value}](form-elements.html#attr-meter-value)
    attribute, if it has one, or the empty string otherwise.

If the element is a [`time`{#values:the-time-element}](text-level-semantics.html#the-time-element) element

:   The value is the element\'s [datetime
    value](text-level-semantics.html#datetime-value){#values:datetime-value}.

Otherwise

:   The value is the element\'s [descendant text
    content](https://dom.spec.whatwg.org/#concept-descendant-text-content){#values:descendant-text-content
    x-internal="descendant-text-content"}.

The [URL property elements]{#url-property-elements .dfn} are the
[`a`{#values:the-a-element-2}](text-level-semantics.html#the-a-element),
[`area`{#values:the-area-element-2}](image-maps.html#the-area-element),
[`audio`{#values:the-audio-element-2}](media.html#the-audio-element),
[`embed`{#values:the-embed-element-2}](iframe-embed-object.html#the-embed-element),
[`iframe`{#values:the-iframe-element-2}](iframe-embed-object.html#the-iframe-element),
[`img`{#values:the-img-element-2}](embedded-content.html#the-img-element),
[`link`{#values:the-link-element-2}](semantics.html#the-link-element),
[`object`{#values:the-object-element-2}](iframe-embed-object.html#the-object-element),
[`source`{#values:the-source-element-2}](embedded-content.html#the-source-element),
[`track`{#values:the-track-element-2}](media.html#the-track-element),
and [`video`{#values:the-video-element-2}](media.html#the-video-element)
elements.

If a property\'s
[value](#concept-property-value){#values:concept-property-value}, as
defined by the property\'s definition, is an [absolute
URL](https://url.spec.whatwg.org/#syntax-url-absolute){#values:absolute-url
x-internal="absolute-url"}, the property must be specified using a [URL
property
element](#url-property-elements){#values:url-property-elements}.

These requirements do not apply just because a property value happens to
match the syntax for a URL. They only apply if the property is
explicitly defined as taking such a value.

For example, a book about the first moon landing could be called
\"mission:moon\". A \"title\" property from a vocabulary that defines a
title as being a string would not expect the title to be given in an
[`a`{#values:the-a-element-3}](text-level-semantics.html#the-a-element)
element, even though it looks like a
[URL](https://url.spec.whatwg.org/#concept-url){#values:url
x-internal="url"}. On the other hand, if there was a (rather narrowly
scoped!) vocabulary for \"books whose titles look like URLs\" which had
a \"title\" property defined to take a URL, then the property *would*
expect the title to be given in an
[`a`{#values:the-a-element-4}](text-level-semantics.html#the-a-element)
element (or one of the other [URL property
elements](#url-property-elements){#values:url-property-elements-2}),
because of the requirement above.

#### [5.2.5]{.secno} Associating names with items[](#associating-names-with-items){.self-link}

To find [the properties of an item]{#the-properties-of-an-item .dfn}
defined by the element `root`{.variable}, the user agent must run the
following steps. These steps are also used to flag [microdata
errors](#microdata-error){#associating-names-with-items:microdata-error}.

1.  Let `results`{.variable}, `memory`{.variable}, and
    `pending`{.variable} be empty lists of elements.

2.  Add the element `root`{.variable} to `memory`{.variable}.

3.  Add the child elements of `root`{.variable}, if any, to
    `pending`{.variable}.

4.  If `root`{.variable} has an
    [`itemref`{#associating-names-with-items:attr-itemref}](#attr-itemref)
    attribute, [split the value of that `itemref` attribute on ASCII
    whitespace](https://infra.spec.whatwg.org/#split-on-ascii-whitespace){#associating-names-with-items:split-a-string-on-spaces
    x-internal="split-a-string-on-spaces"}. For each resulting token
    `ID`{.variable}, if there is an element in the
    [tree](https://dom.spec.whatwg.org/#concept-tree){#associating-names-with-items:tree
    x-internal="tree"} of `root`{.variable} with the
    [ID](https://dom.spec.whatwg.org/#concept-id){#associating-names-with-items:concept-id
    x-internal="concept-id"} `ID`{.variable}, then add the first such
    element to `pending`{.variable}.

5.  While `pending`{.variable} is not empty:

    1.  Remove an element from `pending`{.variable} and let
        `current`{.variable} be that element.

    2.  If `current`{.variable} is already in `memory`{.variable}, there
        is a [microdata
        error](#microdata-error){#associating-names-with-items:microdata-error-2};
        [continue](https://infra.spec.whatwg.org/#iteration-continue){#associating-names-with-items:continue
        x-internal="continue"}.

    3.  Add `current`{.variable} to `memory`{.variable}.

    4.  If `current`{.variable} does not have an
        [`itemscope`{#associating-names-with-items:attr-itemscope}](#attr-itemscope)
        attribute, then: add all the child elements of
        `current`{.variable} to `pending`{.variable}.

    5.  If `current`{.variable} has an
        [`itemprop`{#associating-names-with-items:names:-the-itemprop-attribute}](#names:-the-itemprop-attribute)
        attribute specified and has one or more [property
        names](#property-names){#associating-names-with-items:property-names},
        then add `current`{.variable} to `results`{.variable}.

6.  Sort `results`{.variable} in [tree
    order](https://dom.spec.whatwg.org/#concept-tree-order){#associating-names-with-items:tree-order
    x-internal="tree-order"}.

7.  Return `results`{.variable}.

A document must not contain any
[items](#concept-item){#associating-names-with-items:concept-item} for
which the algorithm to find [the properties of an
item](#the-properties-of-an-item){#associating-names-with-items:the-properties-of-an-item}
finds any [microdata errors]{#microdata-error .dfn}.

An [item](#concept-item){#associating-names-with-items:concept-item-2}
is a [top-level microdata item]{#top-level-microdata-items .dfn} if its
element does not have an
[`itemprop`{#associating-names-with-items:names:-the-itemprop-attribute-2}](#names:-the-itemprop-attribute)
attribute.

All
[`itemref`{#associating-names-with-items:attr-itemref-2}](#attr-itemref)
attributes in a
[`Document`{#associating-names-with-items:document}](dom.html#document)
must be such that there are no cycles in the graph formed from
representing each
[item](#concept-item){#associating-names-with-items:concept-item-3} in
the
[`Document`{#associating-names-with-items:document-2}](dom.html#document)
as a node in the graph and each
[property](#the-properties-of-an-item){#associating-names-with-items:the-properties-of-an-item-2}
of an item whose
[value](#concept-property-value){#associating-names-with-items:concept-property-value}
is another item as an edge in the graph connecting those two items.

A document must not contain any elements that have an
[`itemprop`{#associating-names-with-items:names:-the-itemprop-attribute-3}](#names:-the-itemprop-attribute)
attribute that would not be found to be a property of any of the
[items](#concept-item){#associating-names-with-items:concept-item-4} in
that document were their
[properties](#the-properties-of-an-item){#associating-names-with-items:the-properties-of-an-item-3}
all to be determined.

::: example
In this example, a single license statement is applied to two works,
using
[`itemref`{#associating-names-with-items:attr-itemref-3}](#attr-itemref)
from the items representing the works:

``` html
<!DOCTYPE HTML>
<html lang="en">
 <head>
  <title>Photo gallery</title>
 </head>
 <body>
  <h1>My photos</h1>
  <figure itemscope itemtype="http://n.whatwg.org/work" itemref="licenses">
   <img itemprop="work" src="images/house.jpeg" alt="A white house, boarded up, sits in a forest.">
   <figcaption itemprop="title">The house I found.</figcaption>
  </figure>
  <figure itemscope itemtype="http://n.whatwg.org/work" itemref="licenses">
   <img itemprop="work" src="images/mailbox.jpeg" alt="Outside the house is a mailbox. It has a leaflet inside.">
   <figcaption itemprop="title">The mailbox.</figcaption>
  </figure>
  <footer>
   <p id="licenses">All images licensed under the <a itemprop="license"
   href="http://www.opensource.org/licenses/mit-license.php">MIT
   license</a>.</p>
  </footer>
 </body>
</html>
```

The above results in two items with the type
\"`http://n.whatwg.org/work`\", one with:

work
:   `images/house.jpeg`

title
:   The house I found.

license
:   `http://www.opensource.org/licenses/mit-license.php`

\...and one with:

work
:   `images/mailbox.jpeg`

title
:   The mailbox.

license
:   `http://www.opensource.org/licenses/mit-license.php`
:::

#### [5.2.6]{.secno} Microdata and other namespaces[](#microdata-and-other-namespaces){.self-link}

Currently, the
[`itemscope`{#microdata-and-other-namespaces:attr-itemscope}](#attr-itemscope),
[`itemprop`{#microdata-and-other-namespaces:names:-the-itemprop-attribute}](#names:-the-itemprop-attribute),
and other microdata attributes are only defined for [HTML
elements](infrastructure.html#html-elements){#microdata-and-other-namespaces:html-elements}.
This means that attributes with the literal names \"`itemscope`\",
\"`itemprop`\", etc, do not cause microdata processing to occur on
elements in other namespaces, such as SVG.

::: example
Thus, in the following example there is only one item, not two.

``` bad
<p itemscope></p> <!-- this is an item (with no properties and no type) -->
<svg itemscope></svg> <!-- this is not, it's just an SVG svg element with an invalid unknown attribute -->
```
:::

### [5.3]{.secno} Sample microdata vocabularies[](#mdvocabs){.self-link} {#mdvocabs}

The vocabularies in this section are primarily intended to demonstrate
how a vocabulary is specified, though they are also usable in their own
right.

#### [5.3.1]{.secno} vCard[](#vcard){.self-link}

An item with the [item type](#item-types){#vcard:item-types}
[`http://microformats.org/profile/hcard`]{#md-vcard .dfn} represents a
person\'s or organization\'s contact information.

This vocabulary does not [support global identifiers for
items](#support-global-identifiers-for-items){#vcard:support-global-identifiers-for-items}.

The following are the type\'s [defined property
names](#defined-property-name){#vcard:defined-property-name}. They are
based on the vocabulary defined in vCard Format Specification (vCard)
and its extensions, where more information on how to interpret the
values can be found. [\[RFC6350\]](references.html#refsRFC6350)

[`kind`]{#md-vcard-kind .dfn}

:   Describes what kind of contact the item represents.

    The [value](#concept-property-value){#vcard:concept-property-value}
    must be text that is [identical
    to](https://infra.spec.whatwg.org/#string-is){#vcard:identical-to
    x-internal="identical-to"} one of the [kind
    strings](#kind-strings){#vcard:kind-strings}.

    A single property with the name
    [`kind`{#vcard:md-vcard-kind}](#md-vcard-kind) may be present within
    each [item](#concept-item){#vcard:concept-item} with the type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard}](#md-vcard).

[`fn`]{#md-vcard-fn .dfn}

:   Gives the formatted text corresponding to the name of the person or
    organization.

    The
    [value](#concept-property-value){#vcard:concept-property-value-2}
    must be text.

    Exactly one property with the name
    [`fn`{#vcard:md-vcard-fn}](#md-vcard-fn) must be present within each
    [item](#concept-item){#vcard:concept-item-2} with the type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-2}](#md-vcard).

[`n`]{#md-vcard-n .dfn}

:   Gives the structured name of the person or organization.

    The
    [value](#concept-property-value){#vcard:concept-property-value-3}
    must be an [item](#concept-item){#vcard:concept-item-3} with zero or
    more of each of the
    [`family-name`{#vcard:md-vcard-n-family-name}](#md-vcard-n-family-name),
    [`given-name`{#vcard:md-vcard-n-given-name}](#md-vcard-n-given-name),
    [`additional-name`{#vcard:md-vcard-n-additional-name}](#md-vcard-n-additional-name),
    [`honorific-prefix`{#vcard:md-vcard-n-honorific-prefix}](#md-vcard-n-honorific-prefix),
    and
    [`honorific-suffix`{#vcard:md-vcard-n-honorific-suffix}](#md-vcard-n-honorific-suffix)
    properties.

    Exactly one property with the name
    [`n`{#vcard:md-vcard-n}](#md-vcard-n) must be present within each
    [item](#concept-item){#vcard:concept-item-4} with the type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-3}](#md-vcard).

[`family-name`]{#md-vcard-n-family-name .dfn} (inside [`n`{#vcard:md-vcard-n-2}](#md-vcard-n))

:   Gives the family name of the person, or the full name of the
    organization.

    The
    [value](#concept-property-value){#vcard:concept-property-value-4}
    must be text.

    Any number of properties with the name
    [`family-name`{#vcard:md-vcard-n-family-name-2}](#md-vcard-n-family-name)
    may be present within the
    [item](#concept-item){#vcard:concept-item-5} that forms the
    [value](#concept-property-value){#vcard:concept-property-value-5} of
    the [`n`{#vcard:md-vcard-n-3}](#md-vcard-n) property of an
    [item](#concept-item){#vcard:concept-item-6} with the type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-4}](#md-vcard).

[`given-name`]{#md-vcard-n-given-name .dfn} (inside [`n`{#vcard:md-vcard-n-4}](#md-vcard-n))

:   Gives the given-name of the person.

    The
    [value](#concept-property-value){#vcard:concept-property-value-6}
    must be text.

    Any number of properties with the name
    [`given-name`{#vcard:md-vcard-n-given-name-2}](#md-vcard-n-given-name)
    may be present within the
    [item](#concept-item){#vcard:concept-item-7} that forms the
    [value](#concept-property-value){#vcard:concept-property-value-7} of
    the [`n`{#vcard:md-vcard-n-5}](#md-vcard-n) property of an
    [item](#concept-item){#vcard:concept-item-8} with the type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-5}](#md-vcard).

[`additional-name`]{#md-vcard-n-additional-name .dfn} (inside [`n`{#vcard:md-vcard-n-6}](#md-vcard-n))

:   Gives the any additional names of the person.

    The
    [value](#concept-property-value){#vcard:concept-property-value-8}
    must be text.

    Any number of properties with the name
    [`additional-name`{#vcard:md-vcard-n-additional-name-2}](#md-vcard-n-additional-name)
    may be present within the
    [item](#concept-item){#vcard:concept-item-9} that forms the
    [value](#concept-property-value){#vcard:concept-property-value-9} of
    the [`n`{#vcard:md-vcard-n-7}](#md-vcard-n) property of an
    [item](#concept-item){#vcard:concept-item-10} with the type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-6}](#md-vcard).

[`honorific-prefix`]{#md-vcard-n-honorific-prefix .dfn} (inside [`n`{#vcard:md-vcard-n-8}](#md-vcard-n))

:   Gives the honorific prefix of the person.

    The
    [value](#concept-property-value){#vcard:concept-property-value-10}
    must be text.

    Any number of properties with the name
    [`honorific-prefix`{#vcard:md-vcard-n-honorific-prefix-2}](#md-vcard-n-honorific-prefix)
    may be present within the
    [item](#concept-item){#vcard:concept-item-11} that forms the
    [value](#concept-property-value){#vcard:concept-property-value-11}
    of the [`n`{#vcard:md-vcard-n-9}](#md-vcard-n) property of an
    [item](#concept-item){#vcard:concept-item-12} with the type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-7}](#md-vcard).

[`honorific-suffix`]{#md-vcard-n-honorific-suffix .dfn} (inside [`n`{#vcard:md-vcard-n-10}](#md-vcard-n))

:   Gives the honorific suffix of the person.

    The
    [value](#concept-property-value){#vcard:concept-property-value-12}
    must be text.

    Any number of properties with the name
    [`honorific-suffix`{#vcard:md-vcard-n-honorific-suffix-2}](#md-vcard-n-honorific-suffix)
    may be present within the
    [item](#concept-item){#vcard:concept-item-13} that forms the
    [value](#concept-property-value){#vcard:concept-property-value-13}
    of the [`n`{#vcard:md-vcard-n-11}](#md-vcard-n) property of an
    [item](#concept-item){#vcard:concept-item-14} with the type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-8}](#md-vcard).

[`nickname`]{#md-vcard-nickname .dfn}

:   Gives the nickname of the person or organization.

    The nickname is the descriptive name given instead of or in addition
    to the one belonging to a person, place, or thing. It can also be
    used to specify a familiar form of a proper name specified by the
    [`fn`{#vcard:md-vcard-fn-2}](#md-vcard-fn) or
    [`n`{#vcard:md-vcard-n-12}](#md-vcard-n) properties.

    The
    [value](#concept-property-value){#vcard:concept-property-value-14}
    must be text.

    Any number of properties with the name
    [`nickname`{#vcard:md-vcard-nickname}](#md-vcard-nickname) may be
    present within each [item](#concept-item){#vcard:concept-item-15}
    with the type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-9}](#md-vcard).

[`photo`]{#md-vcard-photo .dfn}

:   Gives a photograph of the person or organization.

    The
    [value](#concept-property-value){#vcard:concept-property-value-15}
    must be an [absolute
    URL](https://url.spec.whatwg.org/#syntax-url-absolute){#vcard:absolute-url
    x-internal="absolute-url"}.

    Any number of properties with the name
    [`photo`{#vcard:md-vcard-photo}](#md-vcard-photo) may be present
    within each [item](#concept-item){#vcard:concept-item-16} with the
    type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-10}](#md-vcard).

[`bday`]{#md-vcard-bday .dfn}

:   Gives the birth date of the person or organization.

    The
    [value](#concept-property-value){#vcard:concept-property-value-16}
    must be a [valid date
    string](common-microsyntaxes.html#valid-date-string){#vcard:valid-date-string}.

    A single property with the name
    [`bday`{#vcard:md-vcard-bday}](#md-vcard-bday) may be present within
    each [item](#concept-item){#vcard:concept-item-17} with the type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-11}](#md-vcard).

[`anniversary`]{#md-vcard-anniversary .dfn}

:   Gives the birth date of the person or organization.

    The
    [value](#concept-property-value){#vcard:concept-property-value-17}
    must be a [valid date
    string](common-microsyntaxes.html#valid-date-string){#vcard:valid-date-string-2}.

    A single property with the name
    [`anniversary`{#vcard:md-vcard-anniversary}](#md-vcard-anniversary)
    may be present within each
    [item](#concept-item){#vcard:concept-item-18} with the type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-12}](#md-vcard).

[`sex`]{#md-vcard-sex .dfn}

:   Gives the biological sex of the person.

    The
    [value](#concept-property-value){#vcard:concept-property-value-18}
    must be one of `F`, meaning \"female\", `M`, meaning \"male\", `N`,
    meaning \"none or not applicable\", `O`, meaning \"other\", or `U`,
    meaning \"unknown\".

    A single property with the name
    [`sex`{#vcard:md-vcard-sex}](#md-vcard-sex) may be present within
    each [item](#concept-item){#vcard:concept-item-19} with the type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-13}](#md-vcard).

[`gender-identity`]{#md-vcard-gender-identity .dfn}

:   Gives the gender identity of the person.

    The
    [value](#concept-property-value){#vcard:concept-property-value-19}
    must be text.

    A single property with the name
    [`gender-identity`{#vcard:md-vcard-gender-identity}](#md-vcard-gender-identity)
    may be present within each
    [item](#concept-item){#vcard:concept-item-20} with the type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-14}](#md-vcard).

[`adr`]{#md-vcard-adr .dfn}

:   Gives the delivery address of the person or organization.

    The
    [value](#concept-property-value){#vcard:concept-property-value-20}
    must be an [item](#concept-item){#vcard:concept-item-21} with zero
    or more [`type`{#vcard:md-vcard-adr-type}](#md-vcard-adr-type),
    [`post-office-box`{#vcard:md-vcard-adr-post-office-box}](#md-vcard-adr-post-office-box),
    [`extended-address`{#vcard:md-vcard-adr-extended-address}](#md-vcard-adr-extended-address),
    and
    [`street-address`{#vcard:md-vcard-adr-street-address}](#md-vcard-adr-street-address)
    properties, and optionally a
    [`locality`{#vcard:md-vcard-adr-locality}](#md-vcard-adr-locality)
    property, optionally a
    [`region`{#vcard:md-vcard-adr-region}](#md-vcard-adr-region)
    property, optionally a
    [`postal-code`{#vcard:md-vcard-adr-postal-code}](#md-vcard-adr-postal-code)
    property, and optionally a
    [`country-name`{#vcard:md-vcard-adr-country-name}](#md-vcard-adr-country-name)
    property.

    If no [`type`{#vcard:md-vcard-adr-type-2}](#md-vcard-adr-type)
    properties are present within an
    [item](#concept-item){#vcard:concept-item-22} that forms the
    [value](#concept-property-value){#vcard:concept-property-value-21}
    of an [`adr`{#vcard:md-vcard-adr}](#md-vcard-adr) property of an
    [item](#concept-item){#vcard:concept-item-23} with the type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-15}](#md-vcard),
    then the [address type
    string](#address-type-strings){#vcard:address-type-strings}
    [`work`{#vcard:md-vcard-type-adr-work}](#md-vcard-type-adr-work) is
    implied.

    Any number of properties with the name
    [`adr`{#vcard:md-vcard-adr-2}](#md-vcard-adr) may be present within
    each [item](#concept-item){#vcard:concept-item-24} with the type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-16}](#md-vcard).

[`type`]{#md-vcard-adr-type .dfn} (inside [`adr`{#vcard:md-vcard-adr-3}](#md-vcard-adr))

:   Gives the type of delivery address.

    The
    [value](#concept-property-value){#vcard:concept-property-value-22}
    must be text that is [identical
    to](https://infra.spec.whatwg.org/#string-is){#vcard:identical-to-2
    x-internal="identical-to"} one of the [address type
    strings](#address-type-strings){#vcard:address-type-strings-2}.

    Any number of properties with the name
    [`type`{#vcard:md-vcard-adr-type-3}](#md-vcard-adr-type) may be
    present within the [item](#concept-item){#vcard:concept-item-25}
    that forms the
    [value](#concept-property-value){#vcard:concept-property-value-23}
    of an [`adr`{#vcard:md-vcard-adr-4}](#md-vcard-adr) property of an
    [item](#concept-item){#vcard:concept-item-26} with the type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-17}](#md-vcard),
    but within each such [`adr`{#vcard:md-vcard-adr-5}](#md-vcard-adr)
    property [item](#concept-item){#vcard:concept-item-27} there must
    only be one [`type`{#vcard:md-vcard-adr-type-4}](#md-vcard-adr-type)
    property per distinct value.

[`post-office-box`]{#md-vcard-adr-post-office-box .dfn} (inside [`adr`{#vcard:md-vcard-adr-6}](#md-vcard-adr))

:   Gives the post office box component of the delivery address of the
    person or organization.

    The
    [value](#concept-property-value){#vcard:concept-property-value-24}
    must be text.

    Any number of properties with the name
    [`post-office-box`{#vcard:md-vcard-adr-post-office-box-2}](#md-vcard-adr-post-office-box)
    may be present within the
    [item](#concept-item){#vcard:concept-item-28} that forms the
    [value](#concept-property-value){#vcard:concept-property-value-25}
    of an [`adr`{#vcard:md-vcard-adr-7}](#md-vcard-adr) property of an
    [item](#concept-item){#vcard:concept-item-29} with the type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-18}](#md-vcard).

    vCard urges authors not to use this field.

[`extended-address`]{#md-vcard-adr-extended-address .dfn} (inside [`adr`{#vcard:md-vcard-adr-8}](#md-vcard-adr))

:   Gives an additional component of the delivery address of the person
    or organization.

    The
    [value](#concept-property-value){#vcard:concept-property-value-26}
    must be text.

    Any number of properties with the name
    [`extended-address`{#vcard:md-vcard-adr-extended-address-2}](#md-vcard-adr-extended-address)
    may be present within the
    [item](#concept-item){#vcard:concept-item-30} that forms the
    [value](#concept-property-value){#vcard:concept-property-value-27}
    of an [`adr`{#vcard:md-vcard-adr-9}](#md-vcard-adr) property of an
    [item](#concept-item){#vcard:concept-item-31} with the type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-19}](#md-vcard).

    vCard urges authors not to use this field.

[`street-address`]{#md-vcard-adr-street-address .dfn} (inside [`adr`{#vcard:md-vcard-adr-10}](#md-vcard-adr))

:   Gives the street address component of the delivery address of the
    person or organization.

    The
    [value](#concept-property-value){#vcard:concept-property-value-28}
    must be text.

    Any number of properties with the name
    [`street-address`{#vcard:md-vcard-adr-street-address-2}](#md-vcard-adr-street-address)
    may be present within the
    [item](#concept-item){#vcard:concept-item-32} that forms the
    [value](#concept-property-value){#vcard:concept-property-value-29}
    of an [`adr`{#vcard:md-vcard-adr-11}](#md-vcard-adr) property of an
    [item](#concept-item){#vcard:concept-item-33} with the type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-20}](#md-vcard).

[`locality`]{#md-vcard-adr-locality .dfn} (inside [`adr`{#vcard:md-vcard-adr-12}](#md-vcard-adr))

:   Gives the locality component (e.g. city) of the delivery address of
    the person or organization.

    The
    [value](#concept-property-value){#vcard:concept-property-value-30}
    must be text.

    A single property with the name
    [`locality`{#vcard:md-vcard-adr-locality-2}](#md-vcard-adr-locality)
    may be present within the
    [item](#concept-item){#vcard:concept-item-34} that forms the
    [value](#concept-property-value){#vcard:concept-property-value-31}
    of an [`adr`{#vcard:md-vcard-adr-13}](#md-vcard-adr) property of an
    [item](#concept-item){#vcard:concept-item-35} with the type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-21}](#md-vcard).

[`region`]{#md-vcard-adr-region .dfn} (inside [`adr`{#vcard:md-vcard-adr-14}](#md-vcard-adr))

:   Gives the region component (e.g. state or province) of the delivery
    address of the person or organization.

    The
    [value](#concept-property-value){#vcard:concept-property-value-32}
    must be text.

    A single property with the name
    [`region`{#vcard:md-vcard-adr-region-2}](#md-vcard-adr-region) may
    be present within the [item](#concept-item){#vcard:concept-item-36}
    that forms the
    [value](#concept-property-value){#vcard:concept-property-value-33}
    of an [`adr`{#vcard:md-vcard-adr-15}](#md-vcard-adr) property of an
    [item](#concept-item){#vcard:concept-item-37} with the type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-22}](#md-vcard).

[`postal-code`]{#md-vcard-adr-postal-code .dfn} (inside [`adr`{#vcard:md-vcard-adr-16}](#md-vcard-adr))

:   Gives the postal code component of the delivery address of the
    person or organization.

    The
    [value](#concept-property-value){#vcard:concept-property-value-34}
    must be text.

    A single property with the name
    [`postal-code`{#vcard:md-vcard-adr-postal-code-2}](#md-vcard-adr-postal-code)
    may be present within the
    [item](#concept-item){#vcard:concept-item-38} that forms the
    [value](#concept-property-value){#vcard:concept-property-value-35}
    of an [`adr`{#vcard:md-vcard-adr-17}](#md-vcard-adr) property of an
    [item](#concept-item){#vcard:concept-item-39} with the type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-23}](#md-vcard).

[`country-name`]{#md-vcard-adr-country-name .dfn} (inside [`adr`{#vcard:md-vcard-adr-18}](#md-vcard-adr))

:   Gives the country name component of the delivery address of the
    person or organization.

    The
    [value](#concept-property-value){#vcard:concept-property-value-36}
    must be text.

    A single property with the name
    [`country-name`{#vcard:md-vcard-adr-country-name-2}](#md-vcard-adr-country-name)
    may be present within the
    [item](#concept-item){#vcard:concept-item-40} that forms the
    [value](#concept-property-value){#vcard:concept-property-value-37}
    of an [`adr`{#vcard:md-vcard-adr-19}](#md-vcard-adr) property of an
    [item](#concept-item){#vcard:concept-item-41} with the type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-24}](#md-vcard).

[`tel`]{#md-vcard-tel .dfn}

:   Gives the telephone number of the person or organization.

    The
    [value](#concept-property-value){#vcard:concept-property-value-38}
    must be either text that can be interpreted as a telephone number as
    defined in the CCITT specifications E.163 and X.121, or an
    [item](#concept-item){#vcard:concept-item-42} with zero or more
    [`type`{#vcard:md-vcard-tel-type}](#md-vcard-tel-type) properties
    and exactly one
    [`value`{#vcard:md-vcard-tel-value}](#md-vcard-tel-value) property.
    [\[E163\]](references.html#refsE163)
    [\[X121\]](references.html#refsX121)

    If no [`type`{#vcard:md-vcard-tel-type-2}](#md-vcard-tel-type)
    properties are present within an
    [item](#concept-item){#vcard:concept-item-43} that forms the
    [value](#concept-property-value){#vcard:concept-property-value-39}
    of a [`tel`{#vcard:md-vcard-tel}](#md-vcard-tel) property of an
    [item](#concept-item){#vcard:concept-item-44} with the type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-25}](#md-vcard),
    or if the
    [value](#concept-property-value){#vcard:concept-property-value-40}
    of such a [`tel`{#vcard:md-vcard-tel-2}](#md-vcard-tel) property is
    text, then the [telephone type
    string](#telephone-type-strings){#vcard:telephone-type-strings}
    [`voice`{#vcard:md-vcard-type-tel-voice}](#md-vcard-type-tel-voice)
    is implied.

    Any number of properties with the name
    [`tel`{#vcard:md-vcard-tel-3}](#md-vcard-tel) may be present within
    each [item](#concept-item){#vcard:concept-item-45} with the type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-26}](#md-vcard).

[`type`]{#md-vcard-tel-type .dfn} (inside [`tel`{#vcard:md-vcard-tel-4}](#md-vcard-tel))

:   Gives the type of telephone number.

    The
    [value](#concept-property-value){#vcard:concept-property-value-41}
    must be text that is [identical
    to](https://infra.spec.whatwg.org/#string-is){#vcard:identical-to-3
    x-internal="identical-to"} one of the [telephone type
    strings](#telephone-type-strings){#vcard:telephone-type-strings-2}.

    Any number of properties with the name
    [`type`{#vcard:md-vcard-tel-type-3}](#md-vcard-tel-type) may be
    present within the [item](#concept-item){#vcard:concept-item-46}
    that forms the
    [value](#concept-property-value){#vcard:concept-property-value-42}
    of a [`tel`{#vcard:md-vcard-tel-5}](#md-vcard-tel) property of an
    [item](#concept-item){#vcard:concept-item-47} with the type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-27}](#md-vcard),
    but within each such [`tel`{#vcard:md-vcard-tel-6}](#md-vcard-tel)
    property [item](#concept-item){#vcard:concept-item-48} there must
    only be one [`type`{#vcard:md-vcard-tel-type-4}](#md-vcard-tel-type)
    property per distinct value.

[`value`]{#md-vcard-tel-value .dfn} (inside [`tel`{#vcard:md-vcard-tel-7}](#md-vcard-tel))

:   Gives the actual telephone number of the person or organization.

    The
    [value](#concept-property-value){#vcard:concept-property-value-43}
    must be text that can be interpreted as a telephone number as
    defined in the CCITT specifications E.163 and X.121.
    [\[E163\]](references.html#refsE163)
    [\[X121\]](references.html#refsX121)

    Exactly one property with the name
    [`value`{#vcard:md-vcard-tel-value-2}](#md-vcard-tel-value) must be
    present within the [item](#concept-item){#vcard:concept-item-49}
    that forms the
    [value](#concept-property-value){#vcard:concept-property-value-44}
    of a [`tel`{#vcard:md-vcard-tel-8}](#md-vcard-tel) property of an
    [item](#concept-item){#vcard:concept-item-50} with the type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-28}](#md-vcard).

[`email`]{#md-vcard-email .dfn}

:   Gives the email address of the person or organization.

    The
    [value](#concept-property-value){#vcard:concept-property-value-45}
    must be text.

    Any number of properties with the name
    [`email`{#vcard:md-vcard-email}](#md-vcard-email) may be present
    within each [item](#concept-item){#vcard:concept-item-51} with the
    type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-29}](#md-vcard).

[`impp`]{#md-vcard-impp .dfn}

:   Gives a [URL](https://url.spec.whatwg.org/#concept-url){#vcard:url
    x-internal="url"} for instant messaging and presence protocol
    communications with the person or organization.

    The
    [value](#concept-property-value){#vcard:concept-property-value-46}
    must be an [absolute
    URL](https://url.spec.whatwg.org/#syntax-url-absolute){#vcard:absolute-url-2
    x-internal="absolute-url"}.

    Any number of properties with the name
    [`impp`{#vcard:md-vcard-impp}](#md-vcard-impp) may be present within
    each [item](#concept-item){#vcard:concept-item-52} with the type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-30}](#md-vcard).

[`lang`]{#md-vcard-lang .dfn}

:   Gives a language understood by the person or organization.

    The
    [value](#concept-property-value){#vcard:concept-property-value-47}
    must be a valid BCP 47 language tag.
    [\[BCP47\]](references.html#refsBCP47)

    Any number of properties with the name
    [`lang`{#vcard:md-vcard-lang}](#md-vcard-lang) may be present within
    each [item](#concept-item){#vcard:concept-item-53} with the type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-31}](#md-vcard).

[`tz`]{#md-vcard-tz .dfn}

:   Gives the time zone of the person or organization.

    The
    [value](#concept-property-value){#vcard:concept-property-value-48}
    must be text and must match the following syntax:

    1.  Either a U+002B PLUS SIGN character (+) or a U+002D HYPHEN-MINUS
        character (-).
    2.  A [valid non-negative
        integer](common-microsyntaxes.html#valid-non-negative-integer){#vcard:valid-non-negative-integer}
        that is exactly two digits long and that represents a number in
        the range 00..23.
    3.  A U+003A COLON character (:).
    4.  A [valid non-negative
        integer](common-microsyntaxes.html#valid-non-negative-integer){#vcard:valid-non-negative-integer-2}
        that is exactly two digits long and that represents a number in
        the range 00..59.

    Any number of properties with the name
    [`tz`{#vcard:md-vcard-tz}](#md-vcard-tz) may be present within each
    [item](#concept-item){#vcard:concept-item-54} with the type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-32}](#md-vcard).

[`geo`]{#md-vcard-geo .dfn}

:   Gives the geographical position of the person or organization.

    The
    [value](#concept-property-value){#vcard:concept-property-value-49}
    must be text and must match the following syntax:

    1.  Optionally, either a U+002B PLUS SIGN character (+) or a U+002D
        HYPHEN-MINUS character (-).
    2.  One or more [ASCII
        digits](https://infra.spec.whatwg.org/#ascii-digit){#vcard:ascii-digits
        x-internal="ascii-digits"}.
    3.  Optionally\*, a U+002E FULL STOP character (.) followed by one
        or more [ASCII
        digits](https://infra.spec.whatwg.org/#ascii-digit){#vcard:ascii-digits-2
        x-internal="ascii-digits"}.
    4.  A U+003B SEMICOLON character (;).
    5.  Optionally, either a U+002B PLUS SIGN character (+) or a U+002D
        HYPHEN-MINUS character (-).
    6.  One or more [ASCII
        digits](https://infra.spec.whatwg.org/#ascii-digit){#vcard:ascii-digits-3
        x-internal="ascii-digits"}.
    7.  Optionally\*, a U+002E FULL STOP character (.) followed by one
        or more [ASCII
        digits](https://infra.spec.whatwg.org/#ascii-digit){#vcard:ascii-digits-4
        x-internal="ascii-digits"}.

    The optional components marked with an asterisk (\*) should be
    included, and should have six digits each.

    The value specifies latitude and longitude, in that order (i.e.,
    \"LAT LON\" ordering), in decimal degrees. The longitude represents
    the location east and west of the prime meridian as a positive or
    negative real number, respectively. The latitude represents the
    location north and south of the equator as a positive or negative
    real number, respectively.

    Any number of properties with the name
    [`geo`{#vcard:md-vcard-geo}](#md-vcard-geo) may be present within
    each [item](#concept-item){#vcard:concept-item-55} with the type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-33}](#md-vcard).

[`title`]{#md-vcard-title .dfn}

:   Gives the job title, functional position or function of the person
    or organization.

    The
    [value](#concept-property-value){#vcard:concept-property-value-50}
    must be text.

    Any number of properties with the name
    [`title`{#vcard:md-vcard-title}](#md-vcard-title) may be present
    within each [item](#concept-item){#vcard:concept-item-56} with the
    type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-34}](#md-vcard).

[`role`]{#md-vcard-role .dfn}

:   Gives the role, occupation, or business category of the person or
    organization.

    The
    [value](#concept-property-value){#vcard:concept-property-value-51}
    must be text.

    Any number of properties with the name
    [`role`{#vcard:md-vcard-role}](#md-vcard-role) may be present within
    each [item](#concept-item){#vcard:concept-item-57} with the type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-35}](#md-vcard).

[`logo`]{#md-vcard-logo .dfn}

:   Gives the logo of the person or organization.

    The
    [value](#concept-property-value){#vcard:concept-property-value-52}
    must be an [absolute
    URL](https://url.spec.whatwg.org/#syntax-url-absolute){#vcard:absolute-url-3
    x-internal="absolute-url"}.

    Any number of properties with the name
    [`logo`{#vcard:md-vcard-logo}](#md-vcard-logo) may be present within
    each [item](#concept-item){#vcard:concept-item-58} with the type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-36}](#md-vcard).

[`agent`]{#md-vcard-agent .dfn}

:   Gives the contact information of another person who will act on
    behalf of the person or organization.

    The
    [value](#concept-property-value){#vcard:concept-property-value-53}
    must be either an [item](#concept-item){#vcard:concept-item-59} with
    the type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-37}](#md-vcard),
    or an [absolute
    URL](https://url.spec.whatwg.org/#syntax-url-absolute){#vcard:absolute-url-4
    x-internal="absolute-url"}, or text.

    Any number of properties with the name
    [`agent`{#vcard:md-vcard-agent}](#md-vcard-agent) may be present
    within each [item](#concept-item){#vcard:concept-item-60} with the
    type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-38}](#md-vcard).

[`org`]{#md-vcard-org .dfn}

:   Gives the name and units of the organization.

    The
    [value](#concept-property-value){#vcard:concept-property-value-54}
    must be either text or an
    [item](#concept-item){#vcard:concept-item-61} with one
    [`organization-name`{#vcard:md-vcard-org-organization-name}](#md-vcard-org-organization-name)
    property and zero or more
    [`organization-unit`{#vcard:md-vcard-org-organization-unit}](#md-vcard-org-organization-unit)
    properties.

    Any number of properties with the name
    [`org`{#vcard:md-vcard-org}](#md-vcard-org) may be present within
    each [item](#concept-item){#vcard:concept-item-62} with the type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-39}](#md-vcard).

[`organization-name`]{#md-vcard-org-organization-name .dfn} (inside [`org`{#vcard:md-vcard-org-2}](#md-vcard-org))

:   Gives the name of the organization.

    The
    [value](#concept-property-value){#vcard:concept-property-value-55}
    must be text.

    Exactly one property with the name
    [`organization-name`{#vcard:md-vcard-org-organization-name-2}](#md-vcard-org-organization-name)
    must be present within the
    [item](#concept-item){#vcard:concept-item-63} that forms the
    [value](#concept-property-value){#vcard:concept-property-value-56}
    of an [`org`{#vcard:md-vcard-org-3}](#md-vcard-org) property of an
    [item](#concept-item){#vcard:concept-item-64} with the type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-40}](#md-vcard).

[`organization-unit`]{#md-vcard-org-organization-unit .dfn} (inside [`org`{#vcard:md-vcard-org-4}](#md-vcard-org))

:   Gives the name of the organization unit.

    The
    [value](#concept-property-value){#vcard:concept-property-value-57}
    must be text.

    Any number of properties with the name
    [`organization-unit`{#vcard:md-vcard-org-organization-unit-2}](#md-vcard-org-organization-unit)
    may be present within the
    [item](#concept-item){#vcard:concept-item-65} that forms the
    [value](#concept-property-value){#vcard:concept-property-value-58}
    of the [`org`{#vcard:md-vcard-org-5}](#md-vcard-org) property of an
    [item](#concept-item){#vcard:concept-item-66} with the type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-41}](#md-vcard).

[`member`]{#md-vcard-member .dfn}

:   Gives a [URL](https://url.spec.whatwg.org/#concept-url){#vcard:url-2
    x-internal="url"} that represents a member of the group.

    The
    [value](#concept-property-value){#vcard:concept-property-value-59}
    must be an [absolute
    URL](https://url.spec.whatwg.org/#syntax-url-absolute){#vcard:absolute-url-5
    x-internal="absolute-url"}.

    Any number of properties with the name
    [`member`{#vcard:md-vcard-member}](#md-vcard-member) may be present
    within each [item](#concept-item){#vcard:concept-item-67} with the
    type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-42}](#md-vcard)
    if the [item](#concept-item){#vcard:concept-item-68} also has a
    property with the name
    [`kind`{#vcard:md-vcard-kind-2}](#md-vcard-kind) whose value is
    \"[`group`{#vcard:md-vcard-kind-group}](#md-vcard-kind-group)\".

[`related`]{#md-vcard-related .dfn}

:   Gives a relationship to another entity.

    The
    [value](#concept-property-value){#vcard:concept-property-value-60}
    must be an [item](#concept-item){#vcard:concept-item-69} with one
    [`url`{#vcard:md-vcard-related-url}](#md-vcard-related-url) property
    and one [`rel`{#vcard:md-vcard-related-rel}](#md-vcard-related-rel)
    properties.

    Any number of properties with the name
    [`related`{#vcard:md-vcard-related}](#md-vcard-related) may be
    present within each [item](#concept-item){#vcard:concept-item-70}
    with the type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-43}](#md-vcard).

[`url`]{#md-vcard-related-url .dfn} (inside [`related`{#vcard:md-vcard-related-2}](#md-vcard-related))

:   Gives the
    [URL](https://url.spec.whatwg.org/#concept-url){#vcard:url-3
    x-internal="url"} for the related entity.

    The
    [value](#concept-property-value){#vcard:concept-property-value-61}
    must be an [absolute
    URL](https://url.spec.whatwg.org/#syntax-url-absolute){#vcard:absolute-url-6
    x-internal="absolute-url"}.

    Exactly one property with the name
    [`url`{#vcard:md-vcard-related-url-2}](#md-vcard-related-url) must
    be present within the [item](#concept-item){#vcard:concept-item-71}
    that forms the
    [value](#concept-property-value){#vcard:concept-property-value-62}
    of a [`related`{#vcard:md-vcard-related-3}](#md-vcard-related)
    property of an [item](#concept-item){#vcard:concept-item-72} with
    the type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-44}](#md-vcard).

[`rel`]{#md-vcard-related-rel .dfn} (inside [`related`{#vcard:md-vcard-related-4}](#md-vcard-related))

:   Gives the relationship between the entity and the related entity.

    The
    [value](#concept-property-value){#vcard:concept-property-value-63}
    must be text that is [identical
    to](https://infra.spec.whatwg.org/#string-is){#vcard:identical-to-4
    x-internal="identical-to"} one of the [relationship
    strings](#relationship-strings){#vcard:relationship-strings}.

    Exactly one property with the name
    [`rel`{#vcard:md-vcard-related-rel-2}](#md-vcard-related-rel) must
    be present within the [item](#concept-item){#vcard:concept-item-73}
    that forms the
    [value](#concept-property-value){#vcard:concept-property-value-64}
    of a [`related`{#vcard:md-vcard-related-5}](#md-vcard-related)
    property of an [item](#concept-item){#vcard:concept-item-74} with
    the type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-45}](#md-vcard).

[`categories`]{#md-vcard-categories .dfn}

:   Gives the name of a category or tag that the person or organization
    could be classified as.

    The
    [value](#concept-property-value){#vcard:concept-property-value-65}
    must be text.

    Any number of properties with the name
    [`categories`{#vcard:md-vcard-categories}](#md-vcard-categories) may
    be present within each [item](#concept-item){#vcard:concept-item-75}
    with the type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-46}](#md-vcard).

[`note`]{#md-vcard-note .dfn}

:   Gives supplemental information or a comment about the person or
    organization.

    The
    [value](#concept-property-value){#vcard:concept-property-value-66}
    must be text.

    Any number of properties with the name
    [`note`{#vcard:md-vcard-note}](#md-vcard-note) may be present within
    each [item](#concept-item){#vcard:concept-item-76} with the type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-47}](#md-vcard).

[`rev`]{#md-vcard-rev .dfn}

:   Gives the revision date and time of the contact information.

    The
    [value](#concept-property-value){#vcard:concept-property-value-67}
    must be text that is a [valid global date and time
    string](common-microsyntaxes.html#valid-global-date-and-time-string){#vcard:valid-global-date-and-time-string}.

    The value distinguishes the current revision of the information for
    other renditions of the information.

    Any number of properties with the name
    [`rev`{#vcard:md-vcard-rev}](#md-vcard-rev) may be present within
    each [item](#concept-item){#vcard:concept-item-77} with the type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-48}](#md-vcard).

[`sound`]{#md-vcard-sound .dfn}

:   Gives a sound file relating to the person or organization.

    The
    [value](#concept-property-value){#vcard:concept-property-value-68}
    must be an [absolute
    URL](https://url.spec.whatwg.org/#syntax-url-absolute){#vcard:absolute-url-7
    x-internal="absolute-url"}.

    Any number of properties with the name
    [`sound`{#vcard:md-vcard-sound}](#md-vcard-sound) may be present
    within each [item](#concept-item){#vcard:concept-item-78} with the
    type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-49}](#md-vcard).

[`uid`]{#md-vcard-uid .dfn}

:   Gives a globally unique identifier corresponding to the person or
    organization.

    The
    [value](#concept-property-value){#vcard:concept-property-value-69}
    must be text.

    A single property with the name
    [`uid`{#vcard:md-vcard-uid}](#md-vcard-uid) may be present within
    each [item](#concept-item){#vcard:concept-item-79} with the type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-50}](#md-vcard).

[`url`]{#md-vcard-url .dfn}

:   Gives a [URL](https://url.spec.whatwg.org/#concept-url){#vcard:url-4
    x-internal="url"} relating to the person or organization.

    The
    [value](#concept-property-value){#vcard:concept-property-value-70}
    must be an [absolute
    URL](https://url.spec.whatwg.org/#syntax-url-absolute){#vcard:absolute-url-8
    x-internal="absolute-url"}.

    Any number of properties with the name
    [`url`{#vcard:md-vcard-url}](#md-vcard-url) may be present within
    each [item](#concept-item){#vcard:concept-item-80} with the type
    [`http://microformats.org/profile/hcard`{#vcard:md-vcard-51}](#md-vcard).

The [kind strings]{#kind-strings .dfn} are:

[`individual`]{#md-vcard-kind-individual .dfn}

:   Indicates a single entity (e.g. a person).

[`group`]{#md-vcard-kind-group .dfn}

:   Indicates multiple entities (e.g. a mailing list).

[`org`]{#md-vcard-kind-org .dfn}

:   Indicates a single entity that is not a person (e.g. a company).

[`location`]{#md-vcard-kind-location .dfn}

:   Indicates a geographical place (e.g. an office building).

The [address type strings]{#address-type-strings .dfn} are:

[`home`]{#md-vcard-type-adr-home .dfn}

:   Indicates a delivery address for a residence.

[`work`]{#md-vcard-type-adr-work .dfn}

:   Indicates a delivery address for a place of work.

The [telephone type strings]{#telephone-type-strings .dfn} are:

[`home`]{#md-vcard-type-tel-home .dfn}

:   Indicates a residential number.

[`work`]{#md-vcard-type-tel-work .dfn}

:   Indicates a telephone number for a place of work.

[`text`]{#md-vcard-type-tel-text .dfn}

:   Indicates that the telephone number supports text messages (SMS).

[`voice`]{#md-vcard-type-tel-voice .dfn}

:   Indicates a voice telephone number.

[`fax`]{#md-vcard-type-tel-fax .dfn}

:   Indicates a facsimile telephone number.

[`cell`]{#md-vcard-type-tel-cell .dfn}

:   Indicates a cellular telephone number.

[`video`]{#md-vcard-type-tel-video .dfn}

:   Indicates a video conferencing telephone number.

[`pager`]{#md-vcard-type-tel-pager .dfn}

:   Indicates a paging device telephone number.

[`textphone`]{#md-vcard-type-tel-textphone .dfn}

:   Indicates a telecommunication device for people with hearing or
    speech difficulties.

The [relationship strings]{#relationship-strings .dfn} are:

[`emergency`]{#md-vcard-rel-emergency .dfn}

:   An emergency contact.

[`agent`]{#md-vcard-rel-agent .dfn}

:   Another entity that acts on behalf of this entity.

[contact]{#md-vcard-rel-contact .dfn}\
[acquaintance]{#md-vcard-rel-acquaintance .dfn}\
[friend]{#md-vcard-rel-friend .dfn}\
[met]{#md-vcard-rel-met .dfn}\
[worker]{#md-vcard-rel-co-worker .dfn}\
[colleague]{#md-vcard-rel-colleague .dfn}\
[resident]{#md-vcard-rel-co-resident .dfn}\
[neighbor]{#md-vcard-rel-neighbor .dfn}\
[child]{#md-vcard-rel-child .dfn}\
[parent]{#md-vcard-rel-parent .dfn}\
[sibling]{#md-vcard-rel-sibling .dfn}\
[spouse]{#md-vcard-rel-spouse .dfn}\
[kin]{#md-vcard-rel-kin .dfn}\
[muse]{#md-vcard-rel-muse .dfn}\
[crush]{#md-vcard-rel-crush .dfn}\
[date]{#md-vcard-rel-date .dfn}\
[sweetheart]{#md-vcard-rel-sweetheart .dfn}\
[me]{#md-vcard-rel-me .dfn}

:   Has the meaning defined in XFN. [\[XFN\]](references.html#refsXFN)

##### [5.3.1.1]{.secno} Conversion to vCard[](#conversion-to-vcard){.self-link}

Given a list of nodes `nodes`{.variable} in a
[`Document`{#conversion-to-vcard:document}](dom.html#document), a user
agent must run the following algorithm to [extract any vCard data
represented by those nodes]{#extracting-a-vcard .dfn} (only the first
vCard is returned):

1.  If none of the nodes in `nodes`{.variable} are
    [items](#concept-item){#conversion-to-vcard:concept-item} with the
    [item type](#item-types){#conversion-to-vcard:item-types}
    [`http://microformats.org/profile/hcard`{#conversion-to-vcard:md-vcard}](#md-vcard),
    then there is no vCard. Abort the algorithm, returning nothing.

2.  Let `node`{.variable} be the first node in `nodes`{.variable} that
    is an [item](#concept-item){#conversion-to-vcard:concept-item-2}
    with the [item type](#item-types){#conversion-to-vcard:item-types-2}
    [`http://microformats.org/profile/hcard`{#conversion-to-vcard:md-vcard-2}](#md-vcard).

3.  Let `output`{.variable} be an empty string.

4.  [Add a vCard
    line](#add-a-vcard-line){#conversion-to-vcard:add-a-vcard-line} with
    the type \"`BEGIN`\" and the value \"`VCARD`\" to
    `output`{.variable}.

5.  [Add a vCard
    line](#add-a-vcard-line){#conversion-to-vcard:add-a-vcard-line-2}
    with the type \"`PROFILE`\" and the value \"`VCARD`\" to
    `output`{.variable}.

6.  [Add a vCard
    line](#add-a-vcard-line){#conversion-to-vcard:add-a-vcard-line-3}
    with the type \"`VERSION`\" and the value \"`4.0`\" to
    `output`{.variable}.

7.  [Add a vCard
    line](#add-a-vcard-line){#conversion-to-vcard:add-a-vcard-line-4}
    with the type \"`SOURCE`\" and the result of [escaping the vCard
    text
    string](#escaping-the-vcard-text-string){#conversion-to-vcard:escaping-the-vcard-text-string}
    that is the document\'s
    [URL](https://dom.spec.whatwg.org/#concept-document-url){#conversion-to-vcard:the-document's-address
    x-internal="the-document's-address"} as the value to
    `output`{.variable}.

8.  If [the `title`
    element](dom.html#the-title-element-2){#conversion-to-vcard:the-title-element-2}
    is not null, [add a vCard
    line](#add-a-vcard-line){#conversion-to-vcard:add-a-vcard-line-5}
    with the type \"`NAME`\" and with the result of [escaping the vCard
    text
    string](#escaping-the-vcard-text-string){#conversion-to-vcard:escaping-the-vcard-text-string-2}
    obtained from [the `title`
    element](dom.html#the-title-element-2){#conversion-to-vcard:the-title-element-2-2}\'s
    [descendant text
    content](https://dom.spec.whatwg.org/#concept-descendant-text-content){#conversion-to-vcard:descendant-text-content
    x-internal="descendant-text-content"} as the value to
    `output`{.variable}.

9.  Let `sex`{.variable} be the empty string.

10. Let `gender-identity`{.variable} be the empty string.

11. For each element `element`{.variable} that is [a property of the
    item](#the-properties-of-an-item){#conversion-to-vcard:the-properties-of-an-item}
    `node`{.variable}: for each name `name`{.variable} in
    `element`{.variable}\'s [property
    names](#property-names){#conversion-to-vcard:property-names}, run
    the following substeps:

    1.  Let `parameters`{.variable} be an empty set of name-value pairs.

    2.  Run the appropriate set of substeps from the following list. The
        steps will set a variable `value`{.variable}, which is used in
        the next step.

        If the property\'s [value](#concept-property-value){#conversion-to-vcard:concept-property-value} is an [item](#concept-item){#conversion-to-vcard:concept-item-3} `subitem`{.variable} and `name`{.variable} is [`n`{#conversion-to-vcard:md-vcard-n}](#md-vcard-n)

        :   1.  Let `value`{.variable} be the empty string.

            2.  Append to `value`{.variable} the result of [collecting
                the first vCard
                subproperty](#collecting-the-first-vcard-subproperty){#conversion-to-vcard:collecting-the-first-vcard-subproperty}
                named
                [`family-name`{#conversion-to-vcard:md-vcard-n-family-name}](#md-vcard-n-family-name)
                in `subitem`{.variable}.

            3.  Append a U+003B SEMICOLON character (;) to
                `value`{.variable}.

            4.  Append to `value`{.variable} the result of [collecting
                the first vCard
                subproperty](#collecting-the-first-vcard-subproperty){#conversion-to-vcard:collecting-the-first-vcard-subproperty-2}
                named
                [`given-name`{#conversion-to-vcard:md-vcard-n-given-name}](#md-vcard-n-given-name)
                in `subitem`{.variable}.

            5.  Append a U+003B SEMICOLON character (;) to
                `value`{.variable}.

            6.  Append to `value`{.variable} the result of [collecting
                the first vCard
                subproperty](#collecting-the-first-vcard-subproperty){#conversion-to-vcard:collecting-the-first-vcard-subproperty-3}
                named
                [`additional-name`{#conversion-to-vcard:md-vcard-n-additional-name}](#md-vcard-n-additional-name)
                in `subitem`{.variable}.

            7.  Append a U+003B SEMICOLON character (;) to
                `value`{.variable}.

            8.  Append to `value`{.variable} the result of [collecting
                the first vCard
                subproperty](#collecting-the-first-vcard-subproperty){#conversion-to-vcard:collecting-the-first-vcard-subproperty-4}
                named
                [`honorific-prefix`{#conversion-to-vcard:md-vcard-n-honorific-prefix}](#md-vcard-n-honorific-prefix)
                in `subitem`{.variable}.

            9.  Append a U+003B SEMICOLON character (;) to
                `value`{.variable}.

            10. Append to `value`{.variable} the result of [collecting
                the first vCard
                subproperty](#collecting-the-first-vcard-subproperty){#conversion-to-vcard:collecting-the-first-vcard-subproperty-5}
                named
                [`honorific-suffix`{#conversion-to-vcard:md-vcard-n-honorific-suffix}](#md-vcard-n-honorific-suffix)
                in `subitem`{.variable}.

        If the property\'s [value](#concept-property-value){#conversion-to-vcard:concept-property-value-2} is an [item](#concept-item){#conversion-to-vcard:concept-item-4} `subitem`{.variable} and `name`{.variable} is [`adr`{#conversion-to-vcard:md-vcard-adr}](#md-vcard-adr)

        :   1.  Let `value`{.variable} be the empty string.

            2.  Append to `value`{.variable} the result of [collecting
                vCard
                subproperties](#collecting-vcard-subproperties){#conversion-to-vcard:collecting-vcard-subproperties}
                named
                [`post-office-box`{#conversion-to-vcard:md-vcard-adr-post-office-box}](#md-vcard-adr-post-office-box)
                in `subitem`{.variable}.

            3.  Append a U+003B SEMICOLON character (;) to
                `value`{.variable}.

            4.  Append to `value`{.variable} the result of [collecting
                vCard
                subproperties](#collecting-vcard-subproperties){#conversion-to-vcard:collecting-vcard-subproperties-2}
                named
                [`extended-address`{#conversion-to-vcard:md-vcard-adr-extended-address}](#md-vcard-adr-extended-address)
                in `subitem`{.variable}.

            5.  Append a U+003B SEMICOLON character (;) to
                `value`{.variable}.

            6.  Append to `value`{.variable} the result of [collecting
                vCard
                subproperties](#collecting-vcard-subproperties){#conversion-to-vcard:collecting-vcard-subproperties-3}
                named
                [`street-address`{#conversion-to-vcard:md-vcard-adr-street-address}](#md-vcard-adr-street-address)
                in `subitem`{.variable}.

            7.  Append a U+003B SEMICOLON character (;) to
                `value`{.variable}.

            8.  Append to `value`{.variable} the result of [collecting
                the first vCard
                subproperty](#collecting-the-first-vcard-subproperty){#conversion-to-vcard:collecting-the-first-vcard-subproperty-6}
                named
                [`locality`{#conversion-to-vcard:md-vcard-adr-locality}](#md-vcard-adr-locality)
                in `subitem`{.variable}.

            9.  Append a U+003B SEMICOLON character (;) to
                `value`{.variable}.

            10. Append to `value`{.variable} the result of [collecting
                the first vCard
                subproperty](#collecting-the-first-vcard-subproperty){#conversion-to-vcard:collecting-the-first-vcard-subproperty-7}
                named
                [`region`{#conversion-to-vcard:md-vcard-adr-region}](#md-vcard-adr-region)
                in `subitem`{.variable}.

            11. Append a U+003B SEMICOLON character (;) to
                `value`{.variable}.

            12. Append to `value`{.variable} the result of [collecting
                the first vCard
                subproperty](#collecting-the-first-vcard-subproperty){#conversion-to-vcard:collecting-the-first-vcard-subproperty-8}
                named
                [`postal-code`{#conversion-to-vcard:md-vcard-adr-postal-code}](#md-vcard-adr-postal-code)
                in `subitem`{.variable}.

            13. Append a U+003B SEMICOLON character (;) to
                `value`{.variable}.

            14. Append to `value`{.variable} the result of [collecting
                the first vCard
                subproperty](#collecting-the-first-vcard-subproperty){#conversion-to-vcard:collecting-the-first-vcard-subproperty-9}
                named
                [`country-name`{#conversion-to-vcard:md-vcard-adr-country-name}](#md-vcard-adr-country-name)
                in `subitem`{.variable}.

            15. If there is a property named
                [`type`{#conversion-to-vcard:md-vcard-adr-type}](#md-vcard-adr-type)
                in `subitem`{.variable}, and the first such property has
                a
                [value](#concept-property-value){#conversion-to-vcard:concept-property-value-3}
                that is not an
                [item](#concept-item){#conversion-to-vcard:concept-item-5}
                and whose value consists only of [ASCII
                alphanumerics](https://infra.spec.whatwg.org/#ascii-alphanumeric){#conversion-to-vcard:alphanumeric-ascii-characters
                x-internal="alphanumeric-ascii-characters"}, then add a
                parameter named \"`TYPE`\" whose value is the
                [value](#concept-property-value){#conversion-to-vcard:concept-property-value-4}
                of that property to `parameters`{.variable}.

        If the property\'s [value](#concept-property-value){#conversion-to-vcard:concept-property-value-5} is an [item](#concept-item){#conversion-to-vcard:concept-item-6} `subitem`{.variable} and `name`{.variable} is [`org`{#conversion-to-vcard:md-vcard-org}](#md-vcard-org)

        :   1.  Let `value`{.variable} be the empty string.

            2.  Append to `value`{.variable} the result of [collecting
                the first vCard
                subproperty](#collecting-the-first-vcard-subproperty){#conversion-to-vcard:collecting-the-first-vcard-subproperty-10}
                named
                [`organization-name`{#conversion-to-vcard:md-vcard-org-organization-name}](#md-vcard-org-organization-name)
                in `subitem`{.variable}.

            3.  For each property named
                [`organization-unit`{#conversion-to-vcard:md-vcard-org-organization-unit}](#md-vcard-org-organization-unit)
                in `subitem`{.variable}, run the following steps:

                1.  If the
                    [value](#concept-property-value){#conversion-to-vcard:concept-property-value-6}
                    of the property is an
                    [item](#concept-item){#conversion-to-vcard:concept-item-7},
                    then skip this property.

                2.  Append a U+003B SEMICOLON character (;) to
                    `value`{.variable}.

                3.  Append the result of [escaping the vCard text
                    string](#escaping-the-vcard-text-string){#conversion-to-vcard:escaping-the-vcard-text-string-3}
                    given by the
                    [value](#concept-property-value){#conversion-to-vcard:concept-property-value-7}
                    of the property to `value`{.variable}.

        If the property\'s [value](#concept-property-value){#conversion-to-vcard:concept-property-value-8} is an [item](#concept-item){#conversion-to-vcard:concept-item-8} `subitem`{.variable} with the [item type](#item-types){#conversion-to-vcard:item-types-3} [`http://microformats.org/profile/hcard`{#conversion-to-vcard:md-vcard-3}](#md-vcard) and `name`{.variable} is [`related`{#conversion-to-vcard:md-vcard-related}](#md-vcard-related)

        :   1.  Let `value`{.variable} be the empty string.

            2.  If there is a property named
                [`url`{#conversion-to-vcard:md-vcard-related-url}](#md-vcard-related-url)
                in `subitem`{.variable}, and its element is a [URL
                property
                element](#url-property-elements){#conversion-to-vcard:url-property-elements},
                then append the result of [escaping the vCard text
                string](#escaping-the-vcard-text-string){#conversion-to-vcard:escaping-the-vcard-text-string-4}
                given by the
                [value](#concept-property-value){#conversion-to-vcard:concept-property-value-9}
                of the first such property to `value`{.variable}, and
                add a parameter with the name \"`VALUE`\" and the value
                \"`URI`\" to `parameters`{.variable}.

            3.  If there is a property named
                [`rel`{#conversion-to-vcard:md-vcard-related-rel}](#md-vcard-related-rel)
                in `subitem`{.variable}, and the first such property has
                a
                [value](#concept-property-value){#conversion-to-vcard:concept-property-value-10}
                that is not an
                [item](#concept-item){#conversion-to-vcard:concept-item-9}
                and whose value consists only of [ASCII
                alphanumerics](https://infra.spec.whatwg.org/#ascii-alphanumeric){#conversion-to-vcard:alphanumeric-ascii-characters-2
                x-internal="alphanumeric-ascii-characters"}, then add a
                parameter named \"`RELATION`\" whose value is the
                [value](#concept-property-value){#conversion-to-vcard:concept-property-value-11}
                of that property to `parameters`{.variable}.

        If the property\'s [value](#concept-property-value){#conversion-to-vcard:concept-property-value-12} is an [item](#concept-item){#conversion-to-vcard:concept-item-10} and `name`{.variable} is none of the above

        :   1.  Let `value`{.variable} be the result of [collecting the
                first vCard
                subproperty](#collecting-the-first-vcard-subproperty){#conversion-to-vcard:collecting-the-first-vcard-subproperty-11}
                named `value` in `subitem`{.variable}.

            2.  If there is a property named `type` in
                `subitem`{.variable}, and the first such property has a
                [value](#concept-property-value){#conversion-to-vcard:concept-property-value-13}
                that is not an
                [item](#concept-item){#conversion-to-vcard:concept-item-11}
                and whose value consists only of [ASCII
                alphanumerics](https://infra.spec.whatwg.org/#ascii-alphanumeric){#conversion-to-vcard:alphanumeric-ascii-characters-3
                x-internal="alphanumeric-ascii-characters"}, then add a
                parameter named \"`TYPE`\" whose value is the
                [value](#concept-property-value){#conversion-to-vcard:concept-property-value-14}
                of that property to `parameters`{.variable}.

        If the property\'s [value](#concept-property-value){#conversion-to-vcard:concept-property-value-15} is not an [item](#concept-item){#conversion-to-vcard:concept-item-12} and its `name`{.variable} is [`sex`{#conversion-to-vcard:md-vcard-sex}](#md-vcard-sex)

        :   If this is the first such property to be found, set
            `sex`{.variable} to the property\'s
            [value](#concept-property-value){#conversion-to-vcard:concept-property-value-16}.

        If the property\'s [value](#concept-property-value){#conversion-to-vcard:concept-property-value-17} is not an [item](#concept-item){#conversion-to-vcard:concept-item-13} and its `name`{.variable} is [`gender-identity`{#conversion-to-vcard:md-vcard-gender-identity}](#md-vcard-gender-identity)

        :   If this is the first such property to be found, set
            `gender-identity`{.variable} to the property\'s
            [value](#concept-property-value){#conversion-to-vcard:concept-property-value-18}.

        Otherwise (the property\'s [value](#concept-property-value){#conversion-to-vcard:concept-property-value-19} is not an [item](#concept-item){#conversion-to-vcard:concept-item-14})

        :   1.  Let `value`{.variable} be the property\'s
                [value](#concept-property-value){#conversion-to-vcard:concept-property-value-20}.

            2.  If `element`{.variable} is one of the [URL property
                elements](#url-property-elements){#conversion-to-vcard:url-property-elements-2},
                add a parameter with the name \"`VALUE`\" and the value
                \"`URI`\" to `parameters`{.variable}.

            3.  Otherwise, if `name`{.variable} is
                [`bday`{#conversion-to-vcard:md-vcard-bday}](#md-vcard-bday)
                or
                [`anniversary`{#conversion-to-vcard:md-vcard-anniversary}](#md-vcard-anniversary)
                and the `value`{.variable} is a [valid date
                string](common-microsyntaxes.html#valid-date-string){#conversion-to-vcard:valid-date-string},
                add a parameter with the name \"`VALUE`\" and the value
                \"`DATE`\" to `parameters`{.variable}.

            4.  Otherwise, if `name`{.variable} is
                [`rev`{#conversion-to-vcard:md-vcard-rev}](#md-vcard-rev)
                and the `value`{.variable} is a [valid global date and
                time
                string](common-microsyntaxes.html#valid-global-date-and-time-string){#conversion-to-vcard:valid-global-date-and-time-string},
                add a parameter with the name \"`VALUE`\" and the value
                \"`DATE-TIME`\" to `parameters`{.variable}.

            5.  Prefix every U+005C REVERSE SOLIDUS character (\\) in
                `value`{.variable} with another U+005C REVERSE SOLIDUS
                character (\\).

            6.  Prefix every U+002C COMMA character (,) in
                `value`{.variable} with a U+005C REVERSE SOLIDUS
                character (\\).

            7.  Unless `name`{.variable} is
                [`geo`{#conversion-to-vcard:md-vcard-geo}](#md-vcard-geo),
                prefix every U+003B SEMICOLON character (;) in
                `value`{.variable} with a U+005C REVERSE SOLIDUS
                character (\\).

            8.  Replace every U+000D CARRIAGE RETURN U+000A LINE FEED
                character pair (CRLF) in `value`{.variable} with a
                U+005C REVERSE SOLIDUS character (\\) followed by a
                U+006E LATIN SMALL LETTER N character (n).

            9.  Replace every remaining U+000D CARRIAGE RETURN (CR) or
                U+000A LINE FEED (LF) character in `value`{.variable}
                with a U+005C REVERSE SOLIDUS character (\\) followed by
                a U+006E LATIN SMALL LETTER N character (n).

    3.  [Add a vCard
        line](#add-a-vcard-line){#conversion-to-vcard:add-a-vcard-line-6}
        with the type `name`{.variable}, the parameters
        `parameters`{.variable}, and the value `value`{.variable} to
        `output`{.variable}.

12. If either `sex`{.variable} or `gender-identity`{.variable} has a
    value that is not the empty string, [add a vCard
    line](#add-a-vcard-line){#conversion-to-vcard:add-a-vcard-line-7}
    with the type \"`GENDER`\" and the value consisting of the
    concatenation of `sex`{.variable}, a U+003B SEMICOLON character (;),
    and `gender-identity`{.variable} to `output`{.variable}.

13. [Add a vCard
    line](#add-a-vcard-line){#conversion-to-vcard:add-a-vcard-line-8}
    with the type \"`END`\" and the value \"`VCARD`\" to
    `output`{.variable}.

When the above algorithm says that the user agent is to [add a vCard
line]{#add-a-vcard-line .dfn} consisting of a type `type`{.variable},
optionally some parameters, and a value `value`{.variable} to a string
`output`{.variable}, it must run the following steps:

1.  Let `line`{.variable} be an empty string.

2.  Append `type`{.variable}, [converted to ASCII
    uppercase](https://infra.spec.whatwg.org/#ascii-uppercase){#conversion-to-vcard:converted-to-ascii-uppercase
    x-internal="converted-to-ascii-uppercase"}, to `line`{.variable}.

3.  If there are any parameters, then for each parameter, in the order
    that they were added, run these substeps:

    1.  Append a U+003B SEMICOLON character (;) to `line`{.variable}.

    2.  Append the parameter\'s name to `line`{.variable}.

    3.  Append a U+003D EQUALS SIGN character (=) to `line`{.variable}.

    4.  Append the parameter\'s value to `line`{.variable}.

4.  Append a U+003A COLON character (:) to `line`{.variable}.

5.  Append `value`{.variable} to `line`{.variable}.

6.  Let `maximum length`{.variable} be 75.

7.  While `line`{.variable}\'s [code point
    length](https://infra.spec.whatwg.org/#string-code-point-length){#conversion-to-vcard:code-point-length
    x-internal="code-point-length"} is greater than
    `maximum length`{.variable}:

    1.  Append the first `maximum length`{.variable} code points of
        `line`{.variable} to `output`{.variable}.

    2.  Remove the first `maximum length`{.variable} code points from
        `line`{.variable}.

    3.  Append a U+000D CARRIAGE RETURN character (CR) to
        `output`{.variable}.

    4.  Append a U+000A LINE FEED character (LF) to `output`{.variable}.

    5.  Append a U+0020 SPACE character to `output`{.variable}.

    6.  Let `maximum length`{.variable} be 74.

8.  Append (what remains of) `line`{.variable} to `output`{.variable}.

9.  Append a U+000D CARRIAGE RETURN character (CR) to
    `output`{.variable}.

10. Append a U+000A LINE FEED character (LF) to `output`{.variable}.

When the steps above require the user agent to obtain the result of
[collecting vCard subproperties]{#collecting-vcard-subproperties .dfn}
named `subname`{.variable} in `subitem`{.variable}, the user agent must
run the following steps:

1.  Let `value`{.variable} be the empty string.

2.  For each property named `subname`{.variable} in the item
    `subitem`{.variable}, run the following substeps:

    1.  If the
        [value](#concept-property-value){#conversion-to-vcard:concept-property-value-21}
        of the property is itself an
        [item](#concept-item){#conversion-to-vcard:concept-item-15},
        then skip this property.

    2.  If this is not the first property named `subname`{.variable} in
        `subitem`{.variable} (ignoring any that were skipped by the
        previous step), then append a U+002C COMMA character (,) to
        `value`{.variable}.

    3.  Append the result of [escaping the vCard text
        string](#escaping-the-vcard-text-string){#conversion-to-vcard:escaping-the-vcard-text-string-5}
        given by the
        [value](#concept-property-value){#conversion-to-vcard:concept-property-value-22}
        of the property to `value`{.variable}.

3.  Return `value`{.variable}.

When the steps above require the user agent to obtain the result of
[collecting the first vCard
subproperty]{#collecting-the-first-vcard-subproperty .dfn} named
`subname`{.variable} in `subitem`{.variable}, the user agent must run
the following steps:

1.  If there are no properties named `subname`{.variable} in
    `subitem`{.variable}, then return the empty string.

2.  If the
    [value](#concept-property-value){#conversion-to-vcard:concept-property-value-23}
    of the first property named `subname`{.variable} in
    `subitem`{.variable} is an
    [item](#concept-item){#conversion-to-vcard:concept-item-16}, then
    return the empty string.

3.  Return the result of [escaping the vCard text
    string](#escaping-the-vcard-text-string){#conversion-to-vcard:escaping-the-vcard-text-string-6}
    given by the
    [value](#concept-property-value){#conversion-to-vcard:concept-property-value-24}
    of the first property named `subname`{.variable} in
    `subitem`{.variable}.

When the above algorithms say the user agent is to [escape the vCard
text string]{#escaping-the-vcard-text-string .dfn} `value`{.variable},
the user agent must use the following steps:

1.  Prefix every U+005C REVERSE SOLIDUS character (\\) in
    `value`{.variable} with another U+005C REVERSE SOLIDUS character
    (\\).

2.  Prefix every U+002C COMMA character (,) in `value`{.variable} with a
    U+005C REVERSE SOLIDUS character (\\).

3.  Prefix every U+003B SEMICOLON character (;) in `value`{.variable}
    with a U+005C REVERSE SOLIDUS character (\\).

4.  Replace every U+000D CARRIAGE RETURN U+000A LINE FEED character pair
    (CRLF) in `value`{.variable} with a U+005C REVERSE SOLIDUS character
    (\\) followed by a U+006E LATIN SMALL LETTER N character (n).

5.  Replace every remaining U+000D CARRIAGE RETURN (CR) or U+000A LINE
    FEED (LF) character in `value`{.variable} with a U+005C REVERSE
    SOLIDUS character (\\) followed by a U+006E LATIN SMALL LETTER N
    character (n).

6.  Return the mutated `value`{.variable}.

This algorithm can generate invalid vCard output, if the input does not
conform to the rules described for the
[`http://microformats.org/profile/hcard`{#conversion-to-vcard:md-vcard-4}](#md-vcard)
[item type](#item-types){#conversion-to-vcard:item-types-4} and [defined
property
names](#defined-property-name){#conversion-to-vcard:defined-property-name}.

##### [5.3.1.2]{.secno} Examples[](#examples-2){.self-link} {#examples-2}

*This section is non-normative.*

::: example
Here is a long example vCard for a fictional character called \"Jack
Bauer\":

``` html
<section id="jack" itemscope itemtype="http://microformats.org/profile/hcard">
 <h1 itemprop="fn">
  <span itemprop="n" itemscope>
   <span itemprop="given-name">Jack</span>
   <span itemprop="family-name">Bauer</span>
  </span>
 </h1>
 <img itemprop="photo" alt="" src="jack-bauer.jpg">
 <p itemprop="org" itemscope>
  <span itemprop="organization-name">Counter-Terrorist Unit</span>
  (<span itemprop="organization-unit">Los Angeles Division</span>)
 </p>
 <p>
  <span itemprop="adr" itemscope>
   <span itemprop="street-address">10201 W. Pico Blvd.</span><br>
   <span itemprop="locality">Los Angeles</span>,
   <span itemprop="region">CA</span>
   <span itemprop="postal-code">90064</span><br>
   <span itemprop="country-name">United States</span><br>
  </span>
  <span itemprop="geo">34.052339;-118.410623</span>
 </p>
 <h2>Assorted Contact Methods</h2>
 <ul>
  <li itemprop="tel" itemscope>
   <span itemprop="value">+1 (310) 597 3781</span> <span itemprop="type">work</span>
   <meta itemprop="type" content="voice">
  </li>
  <li><a itemprop="url" href="https://en.wikipedia.org/wiki/Jack_Bauer">I'm on Wikipedia</a>
  so you can leave a message on my user talk page.</li>
  <li><a itemprop="url" href="http://www.jackbauerfacts.com/">Jack Bauer Facts</a></li>
  <li itemprop="email"><a href="mailto:j.bauer@la.ctu.gov.invalid">j.bauer@la.ctu.gov.invalid</a></li>
  <li itemprop="tel" itemscope>
   <span itemprop="value">+1 (310) 555 3781</span> <span>
   <meta itemprop="type" content="cell">mobile phone</span>
  </li>
 </ul>
 <ins datetime="2008-07-20 21:00:00+01:00">
  <meta itemprop="rev" content="2008-07-20 21:00:00+01:00">
  <p itemprop="tel" itemscope><strong>Update!</strong>
  My new <span itemprop="type">home</span> phone number is
  <span itemprop="value">01632 960 123</span>.</p>
 </ins>
</section>
```

The odd line wrapping is needed because newlines are meaningful in
microdata: newlines would be preserved in a conversion to, for example,
the vCard format.
:::

::: example
This example shows a site\'s contact details (using the
[`address`{#examples-2:the-address-element}](sections.html#the-address-element)
element) containing an address with two street components:

``` html
<address itemscope itemtype="http://microformats.org/profile/hcard">
 <strong itemprop="fn"><span itemprop="n" itemscope><span itemprop="given-name">Alfred</span>
 <span itemprop="family-name">Person</span></span></strong> <br>
 <span itemprop="adr" itemscope>
  <span itemprop="street-address">1600 Amphitheatre Parkway</span> <br>
  <span itemprop="street-address">Building 43, Second Floor</span> <br>
  <span itemprop="locality">Mountain View</span>,
   <span itemprop="region">CA</span> <span itemprop="postal-code">94043</span>
 </span>
</address>
```
:::

::: example
The vCard vocabulary can be used to just mark up people\'s names:

``` html
<span itemscope itemtype="http://microformats.org/profile/hcard"
><span itemprop=fn><span itemprop="n" itemscope><span itemprop="given-name"
>George</span> <span itemprop="family-name">Washington</span></span
></span></span>
```

This creates a single item with a two name-value pairs, one with the
name \"fn\" and the value \"George Washington\", and the other with the
name \"n\" and a second item as its value, the second item having the
two name-value pairs \"given-name\" and \"family-name\" with the values
\"George\" and \"Washington\" respectively. This is defined to map to
the following vCard:

    BEGIN:VCARD
    PROFILE:VCARD
    VERSION:4.0
    SOURCE:document's address
    FN:George Washington
    N:Washington;George;;;
    END:VCARD
:::

#### [5.3.2]{.secno} vEvent[](#vevent){.self-link}

An item with the [item type](#item-types){#vevent:item-types}
[`http://microformats.org/profile/hcalendar#vevent`]{#md-vevent .dfn}
represents an event.

This vocabulary does not [support global identifiers for
items](#support-global-identifiers-for-items){#vevent:support-global-identifiers-for-items}.

The following are the type\'s [defined property
names](#defined-property-name){#vevent:defined-property-name}. They are
based on the vocabulary defined in Internet Calendaring and Scheduling
Core Object Specification (iCalendar), where more information on how to
interpret the values can be found.
[\[RFC5545\]](references.html#refsRFC5545)

Only the parts of the iCalendar vocabulary relating to events are used
here; this vocabulary cannot express a complete iCalendar instance.

[`attach`]{#md-vevent-attach .dfn}

:   Gives the address of an associated document for the event.

    The [value](#concept-property-value){#vevent:concept-property-value}
    must be an [absolute
    URL](https://url.spec.whatwg.org/#syntax-url-absolute){#vevent:absolute-url
    x-internal="absolute-url"}.

    Any number of properties with the name
    [`attach`{#vevent:md-vevent-attach}](#md-vevent-attach) may be
    present within each [item](#concept-item){#vevent:concept-item} with
    the type
    [`http://microformats.org/profile/hcalendar#vevent`{#vevent:md-vevent}](#md-vevent).

[`categories`]{#md-vevent-categories .dfn}

:   Gives the name of a category or tag that the event could be
    classified as.

    The
    [value](#concept-property-value){#vevent:concept-property-value-2}
    must be text.

    Any number of properties with the name
    [`categories`{#vevent:md-vevent-categories}](#md-vevent-categories)
    may be present within each
    [item](#concept-item){#vevent:concept-item-2} with the type
    [`http://microformats.org/profile/hcalendar#vevent`{#vevent:md-vevent-2}](#md-vevent).

[`class`]{#md-vevent-class .dfn}

:   Gives the access classification of the information regarding the
    event.

    The
    [value](#concept-property-value){#vevent:concept-property-value-3}
    must be text with one of the following values:

    - `public`
    - `private`
    - `confidential`

    This is merely advisory and cannot be considered a confidentiality
    measure.

    A single property with the name
    [`class`{#vevent:md-vevent-class}](#md-vevent-class) may be present
    within each [item](#concept-item){#vevent:concept-item-3} with the
    type
    [`http://microformats.org/profile/hcalendar#vevent`{#vevent:md-vevent-3}](#md-vevent).

[`comment`]{#md-vevent-comment .dfn}

:   Gives a comment regarding the event.

    The
    [value](#concept-property-value){#vevent:concept-property-value-4}
    must be text.

    Any number of properties with the name
    [`comment`{#vevent:md-vevent-comment}](#md-vevent-comment) may be
    present within each [item](#concept-item){#vevent:concept-item-4}
    with the type
    [`http://microformats.org/profile/hcalendar#vevent`{#vevent:md-vevent-4}](#md-vevent).

[`description`]{#md-vevent-description .dfn}

:   Gives a detailed description of the event.

    The
    [value](#concept-property-value){#vevent:concept-property-value-5}
    must be text.

    A single property with the name
    [`description`{#vevent:md-vevent-description}](#md-vevent-description)
    may be present within each
    [item](#concept-item){#vevent:concept-item-5} with the type
    [`http://microformats.org/profile/hcalendar#vevent`{#vevent:md-vevent-5}](#md-vevent).

[`geo`]{#md-vevent-geo .dfn}

:   Gives the geographical position of the event.

    The
    [value](#concept-property-value){#vevent:concept-property-value-6}
    must be text and must match the following syntax:

    1.  Optionally, either a U+002B PLUS SIGN character (+) or a U+002D
        HYPHEN-MINUS character (-).
    2.  One or more [ASCII
        digits](https://infra.spec.whatwg.org/#ascii-digit){#vevent:ascii-digits
        x-internal="ascii-digits"}.
    3.  Optionally\*, a U+002E FULL STOP character (.) followed by one
        or more [ASCII
        digits](https://infra.spec.whatwg.org/#ascii-digit){#vevent:ascii-digits-2
        x-internal="ascii-digits"}.
    4.  A U+003B SEMICOLON character (;).
    5.  Optionally, either a U+002B PLUS SIGN character (+) or a U+002D
        HYPHEN-MINUS character (-).
    6.  One or more [ASCII
        digits](https://infra.spec.whatwg.org/#ascii-digit){#vevent:ascii-digits-3
        x-internal="ascii-digits"}.
    7.  Optionally\*, a U+002E FULL STOP character (.) followed by one
        or more [ASCII
        digits](https://infra.spec.whatwg.org/#ascii-digit){#vevent:ascii-digits-4
        x-internal="ascii-digits"}.

    The optional components marked with an asterisk (\*) should be
    included, and should have six digits each.

    The value specifies latitude and longitude, in that order (i.e.,
    \"LAT LON\" ordering), in decimal degrees. The longitude represents
    the location east and west of the prime meridian as a positive or
    negative real number, respectively. The latitude represents the
    location north and south of the equator as a positive or negative
    real number, respectively.

    A single property with the name
    [`geo`{#vevent:md-vevent-geo}](#md-vevent-geo) may be present within
    each [item](#concept-item){#vevent:concept-item-6} with the type
    [`http://microformats.org/profile/hcalendar#vevent`{#vevent:md-vevent-6}](#md-vevent).

[`location`]{#md-vevent-location .dfn}

:   Gives the location of the event.

    The
    [value](#concept-property-value){#vevent:concept-property-value-7}
    must be text.

    A single property with the name
    [`location`{#vevent:md-vevent-location}](#md-vevent-location) may be
    present within each [item](#concept-item){#vevent:concept-item-7}
    with the type
    [`http://microformats.org/profile/hcalendar#vevent`{#vevent:md-vevent-7}](#md-vevent).

[`resources`]{#md-vevent-resources .dfn}

:   Gives a resource that will be needed for the event.

    The
    [value](#concept-property-value){#vevent:concept-property-value-8}
    must be text.

    Any number of properties with the name
    [`resources`{#vevent:md-vevent-resources}](#md-vevent-resources) may
    be present within each [item](#concept-item){#vevent:concept-item-8}
    with the type
    [`http://microformats.org/profile/hcalendar#vevent`{#vevent:md-vevent-8}](#md-vevent).

[`status`]{#md-vevent-status .dfn}

:   Gives the confirmation status of the event.

    The
    [value](#concept-property-value){#vevent:concept-property-value-9}
    must be text with one of the following values:

    - `tentative`
    - `confirmed`
    - `canceled`

    A single property with the name
    [`status`{#vevent:md-vevent-status}](#md-vevent-status) may be
    present within each [item](#concept-item){#vevent:concept-item-9}
    with the type
    [`http://microformats.org/profile/hcalendar#vevent`{#vevent:md-vevent-9}](#md-vevent).

[`summary`]{#md-vevent-summary .dfn}

:   Gives a short summary of the event.

    The
    [value](#concept-property-value){#vevent:concept-property-value-10}
    must be text.

    User agents should replace U+000A LINE FEED (LF) characters in the
    [value](#concept-property-value){#vevent:concept-property-value-11}
    by U+0020 SPACE characters when using the value.

    A single property with the name
    [`summary`{#vevent:md-vevent-summary}](#md-vevent-summary) may be
    present within each [item](#concept-item){#vevent:concept-item-10}
    with the type
    [`http://microformats.org/profile/hcalendar#vevent`{#vevent:md-vevent-10}](#md-vevent).

[`dtend`]{#md-vevent-dtend .dfn}

:   Gives the date and time by which the event ends.

    If the property with the name
    [`dtend`{#vevent:md-vevent-dtend}](#md-vevent-dtend) is present
    within an [item](#concept-item){#vevent:concept-item-11} with the
    type
    [`http://microformats.org/profile/hcalendar#vevent`{#vevent:md-vevent-11}](#md-vevent)
    that has a property with the name
    [`dtstart`{#vevent:md-vevent-dtstart}](#md-vevent-dtstart) whose
    value is a [valid date
    string](common-microsyntaxes.html#valid-date-string){#vevent:valid-date-string},
    then the
    [value](#concept-property-value){#vevent:concept-property-value-12}
    of the property with the name
    [`dtend`{#vevent:md-vevent-dtend-2}](#md-vevent-dtend) must be text
    that is a [valid date
    string](common-microsyntaxes.html#valid-date-string){#vevent:valid-date-string-2}
    also. Otherwise, the
    [value](#concept-property-value){#vevent:concept-property-value-13}
    of the property must be text that is a [valid global date and time
    string](common-microsyntaxes.html#valid-global-date-and-time-string){#vevent:valid-global-date-and-time-string}.

    In either case, the
    [value](#concept-property-value){#vevent:concept-property-value-14}
    be later in time than the value of the
    [`dtstart`{#vevent:md-vevent-dtstart-2}](#md-vevent-dtstart)
    property of the same [item](#concept-item){#vevent:concept-item-12}.

    The time given by the
    [`dtend`{#vevent:md-vevent-dtend-3}](#md-vevent-dtend) property is
    not inclusive. For day-long events, therefore, the
    [`dtend`{#vevent:md-vevent-dtend-4}](#md-vevent-dtend) property\'s
    [value](#concept-property-value){#vevent:concept-property-value-15}
    will be the day *after* the end of the event.

    A single property with the name
    [`dtend`{#vevent:md-vevent-dtend-5}](#md-vevent-dtend) may be
    present within each [item](#concept-item){#vevent:concept-item-13}
    with the type
    [`http://microformats.org/profile/hcalendar#vevent`{#vevent:md-vevent-12}](#md-vevent),
    so long as that
    [`http://microformats.org/profile/hcalendar#vevent`{#vevent:md-vevent-13}](#md-vevent)
    does not have a property with the name
    [`duration`{#vevent:md-vevent-duration}](#md-vevent-duration).

[`dtstart`]{#md-vevent-dtstart .dfn}

:   Gives the date and time at which the event starts.

    The
    [value](#concept-property-value){#vevent:concept-property-value-16}
    must be text that is either a [valid date
    string](common-microsyntaxes.html#valid-date-string){#vevent:valid-date-string-3}
    or a [valid global date and time
    string](common-microsyntaxes.html#valid-global-date-and-time-string){#vevent:valid-global-date-and-time-string-2}.

    Exactly one property with the name
    [`dtstart`{#vevent:md-vevent-dtstart-3}](#md-vevent-dtstart) must be
    present within each [item](#concept-item){#vevent:concept-item-14}
    with the type
    [`http://microformats.org/profile/hcalendar#vevent`{#vevent:md-vevent-14}](#md-vevent).

[`duration`]{#md-vevent-duration .dfn}

:   Gives the duration of the event.

    The
    [value](#concept-property-value){#vevent:concept-property-value-17}
    must be text that is a [valid vevent duration
    string](#valid-vevent-duration-string){#vevent:valid-vevent-duration-string}.

    The duration represented is the sum of all the durations represented
    by integers in the value.

    A single property with the name
    [`duration`{#vevent:md-vevent-duration-2}](#md-vevent-duration) may
    be present within each
    [item](#concept-item){#vevent:concept-item-15} with the type
    [`http://microformats.org/profile/hcalendar#vevent`{#vevent:md-vevent-15}](#md-vevent),
    so long as that
    [`http://microformats.org/profile/hcalendar#vevent`{#vevent:md-vevent-16}](#md-vevent)
    does not have a property with the name
    [`dtend`{#vevent:md-vevent-dtend-6}](#md-vevent-dtend).

[`transp`]{#md-vevent-transp .dfn}

:   Gives whether the event is to be considered as consuming time on a
    calendar, for the purpose of free-busy time searches.

    The
    [value](#concept-property-value){#vevent:concept-property-value-18}
    must be text with one of the following values:

    - `opaque`
    - `transparent`

    A single property with the name
    [`transp`{#vevent:md-vevent-transp}](#md-vevent-transp) may be
    present within each [item](#concept-item){#vevent:concept-item-16}
    with the type
    [`http://microformats.org/profile/hcalendar#vevent`{#vevent:md-vevent-17}](#md-vevent).

[`contact`]{#md-vevent-contact .dfn}

:   Gives the contact information for the event.

    The
    [value](#concept-property-value){#vevent:concept-property-value-19}
    must be text.

    Any number of properties with the name
    [`contact`{#vevent:md-vevent-contact}](#md-vevent-contact) may be
    present within each [item](#concept-item){#vevent:concept-item-17}
    with the type
    [`http://microformats.org/profile/hcalendar#vevent`{#vevent:md-vevent-18}](#md-vevent).

[`url`]{#md-vevent-url .dfn}

:   Gives a [URL](https://url.spec.whatwg.org/#concept-url){#vevent:url
    x-internal="url"} for the event.

    The
    [value](#concept-property-value){#vevent:concept-property-value-20}
    must be an [absolute
    URL](https://url.spec.whatwg.org/#syntax-url-absolute){#vevent:absolute-url-2
    x-internal="absolute-url"}.

    A single property with the name
    [`url`{#vevent:md-vevent-url}](#md-vevent-url) may be present within
    each [item](#concept-item){#vevent:concept-item-18} with the type
    [`http://microformats.org/profile/hcalendar#vevent`{#vevent:md-vevent-19}](#md-vevent).

[`uid`]{#md-vevent-uid .dfn}

:   Gives a globally unique identifier corresponding to the event.

    The
    [value](#concept-property-value){#vevent:concept-property-value-21}
    must be text.

    A single property with the name
    [`uid`{#vevent:md-vevent-uid}](#md-vevent-uid) may be present within
    each [item](#concept-item){#vevent:concept-item-19} with the type
    [`http://microformats.org/profile/hcalendar#vevent`{#vevent:md-vevent-20}](#md-vevent).

[`exdate`]{#md-vevent-exdate .dfn}

:   Gives a date and time at which the event does not occur despite the
    recurrence rules.

    The
    [value](#concept-property-value){#vevent:concept-property-value-22}
    must be text that is either a [valid date
    string](common-microsyntaxes.html#valid-date-string){#vevent:valid-date-string-4}
    or a [valid global date and time
    string](common-microsyntaxes.html#valid-global-date-and-time-string){#vevent:valid-global-date-and-time-string-3}.

    Any number of properties with the name
    [`exdate`{#vevent:md-vevent-exdate}](#md-vevent-exdate) may be
    present within each [item](#concept-item){#vevent:concept-item-20}
    with the type
    [`http://microformats.org/profile/hcalendar#vevent`{#vevent:md-vevent-21}](#md-vevent).

[`rdate`]{#md-vevent-rdate .dfn}

:   Gives a date and time at which the event recurs.

    The
    [value](#concept-property-value){#vevent:concept-property-value-23}
    must be text that is one of the following:

    - A [valid date
      string](common-microsyntaxes.html#valid-date-string){#vevent:valid-date-string-5}.
    - A [valid global date and time
      string](common-microsyntaxes.html#valid-global-date-and-time-string){#vevent:valid-global-date-and-time-string-4}.
    - A [valid global date and time
      string](common-microsyntaxes.html#valid-global-date-and-time-string){#vevent:valid-global-date-and-time-string-5}
      followed by a U+002F SOLIDUS character (/) followed by a second
      [valid global date and time
      string](common-microsyntaxes.html#valid-global-date-and-time-string){#vevent:valid-global-date-and-time-string-6}
      representing a later time.
    - A [valid global date and time
      string](common-microsyntaxes.html#valid-global-date-and-time-string){#vevent:valid-global-date-and-time-string-7}
      followed by a U+002F SOLIDUS character (/) followed by a [valid
      vevent duration
      string](#valid-vevent-duration-string){#vevent:valid-vevent-duration-string-2}.

    Any number of properties with the name
    [`rdate`{#vevent:md-vevent-rdate}](#md-vevent-rdate) may be present
    within each [item](#concept-item){#vevent:concept-item-21} with the
    type
    [`http://microformats.org/profile/hcalendar#vevent`{#vevent:md-vevent-22}](#md-vevent).

[`rrule`]{#md-vevent-rrule .dfn}

:   Gives a rule for finding dates and times at which the event occurs.

    The
    [value](#concept-property-value){#vevent:concept-property-value-24}
    must be text that matches the RECUR value type defined in iCalendar.
    [\[RFC5545\]](references.html#refsRFC5545)

    A single property with the name
    [`rrule`{#vevent:md-vevent-rrule}](#md-vevent-rrule) may be present
    within each [item](#concept-item){#vevent:concept-item-22} with the
    type
    [`http://microformats.org/profile/hcalendar#vevent`{#vevent:md-vevent-23}](#md-vevent).

[`created`]{#md-vevent-created .dfn}

:   Gives the date and time at which the event information was first
    created in a calendaring system.

    The
    [value](#concept-property-value){#vevent:concept-property-value-25}
    must be text that is a [valid global date and time
    string](common-microsyntaxes.html#valid-global-date-and-time-string){#vevent:valid-global-date-and-time-string-8}.

    A single property with the name
    [`created`{#vevent:md-vevent-created}](#md-vevent-created) may be
    present within each [item](#concept-item){#vevent:concept-item-23}
    with the type
    [`http://microformats.org/profile/hcalendar#vevent`{#vevent:md-vevent-24}](#md-vevent).

[`last-modified`]{#md-vevent-last-modified .dfn}

:   Gives the date and time at which the event information was last
    modified in a calendaring system.

    The
    [value](#concept-property-value){#vevent:concept-property-value-26}
    must be text that is a [valid global date and time
    string](common-microsyntaxes.html#valid-global-date-and-time-string){#vevent:valid-global-date-and-time-string-9}.

    A single property with the name
    [`last-modified`{#vevent:md-vevent-last-modified}](#md-vevent-last-modified)
    may be present within each
    [item](#concept-item){#vevent:concept-item-24} with the type
    [`http://microformats.org/profile/hcalendar#vevent`{#vevent:md-vevent-25}](#md-vevent).

[`sequence`]{#md-vevent-sequence .dfn}

:   Gives a revision number for the event information.

    The
    [value](#concept-property-value){#vevent:concept-property-value-27}
    must be text that is a [valid non-negative
    integer](common-microsyntaxes.html#valid-non-negative-integer){#vevent:valid-non-negative-integer}.

    A single property with the name
    [`sequence`{#vevent:md-vevent-sequence}](#md-vevent-sequence) may be
    present within each [item](#concept-item){#vevent:concept-item-25}
    with the type
    [`http://microformats.org/profile/hcalendar#vevent`{#vevent:md-vevent-26}](#md-vevent).

A string is a [valid vevent duration
string]{#valid-vevent-duration-string .dfn} if it matches the following
pattern:

1.  A U+0050 LATIN CAPITAL LETTER P character (P).

2.  One of the following:

    - A [valid non-negative
      integer](common-microsyntaxes.html#valid-non-negative-integer){#vevent:valid-non-negative-integer-2}
      followed by a U+0057 LATIN CAPITAL LETTER W character (W). The
      integer represents a duration of that number of weeks.

    - At least one, and possible both in this order, of the following:

      1.  A [valid non-negative
          integer](common-microsyntaxes.html#valid-non-negative-integer){#vevent:valid-non-negative-integer-3}
          followed by a U+0044 LATIN CAPITAL LETTER D character (D). The
          integer represents a duration of that number of days.

      2.  A U+0054 LATIN CAPITAL LETTER T character (T) followed by any
          one of the following, or the first and second of the following
          in that order, or the second and third of the following in
          that order, or all three of the following in this order:

          1.  A [valid non-negative
              integer](common-microsyntaxes.html#valid-non-negative-integer){#vevent:valid-non-negative-integer-4}
              followed by a U+0048 LATIN CAPITAL LETTER H character (H).
              The integer represents a duration of that number of hours.

          2.  A [valid non-negative
              integer](common-microsyntaxes.html#valid-non-negative-integer){#vevent:valid-non-negative-integer-5}
              followed by a U+004D LATIN CAPITAL LETTER M character (M).
              The integer represents a duration of that number of
              minutes.

          3.  A [valid non-negative
              integer](common-microsyntaxes.html#valid-non-negative-integer){#vevent:valid-non-negative-integer-6}
              followed by a U+0053 LATIN CAPITAL LETTER S character (S).
              The integer represents a duration of that number of
              seconds.

##### [5.3.2.1]{.secno} Conversion to iCalendar[](#conversion-to-icalendar){.self-link}

Given a list of nodes `nodes`{.variable} in a
[`Document`{#conversion-to-icalendar:document}](dom.html#document), a
user agent must run the following algorithm to [extract any vEvent data
represented by those nodes]{#extracting-vevent-data .dfn}:

1.  If none of the nodes in `nodes`{.variable} are
    [items](#concept-item){#conversion-to-icalendar:concept-item} with
    the type
    [`http://microformats.org/profile/hcalendar#vevent`{#conversion-to-icalendar:md-vevent}](#md-vevent),
    then there is no vEvent data. Abort the algorithm, returning
    nothing.

2.  Let `output`{.variable} be an empty string.

3.  [Add an iCalendar
    line](#add-an-icalendar-line){#conversion-to-icalendar:add-an-icalendar-line}
    with the type \"`BEGIN`\" and the value \"`VCALENDAR`\" to
    `output`{.variable}.

4.  [Add an iCalendar
    line](#add-an-icalendar-line){#conversion-to-icalendar:add-an-icalendar-line-2}
    with the type \"`PRODID`\" and the value equal to a
    user-agent-specific string representing the user agent to
    `output`{.variable}.

5.  [Add an iCalendar
    line](#add-an-icalendar-line){#conversion-to-icalendar:add-an-icalendar-line-3}
    with the type \"`VERSION`\" and the value \"`2.0`\" to
    `output`{.variable}.

6.  For each node `node`{.variable} in `nodes`{.variable} that is an
    [item](#concept-item){#conversion-to-icalendar:concept-item-2} with
    the type
    [`http://microformats.org/profile/hcalendar#vevent`{#conversion-to-icalendar:md-vevent-2}](#md-vevent),
    run the following steps:

    1.  [Add an iCalendar
        line](#add-an-icalendar-line){#conversion-to-icalendar:add-an-icalendar-line-4}
        with the type \"`BEGIN`\" and the value \"`VEVENT`\" to
        `output`{.variable}.

    2.  [Add an iCalendar
        line](#add-an-icalendar-line){#conversion-to-icalendar:add-an-icalendar-line-5}
        with the type \"`DTSTAMP`\" and a value consisting of an
        iCalendar DATE-TIME string representing the current date and
        time, with the annotation \"`VALUE=DATE-TIME`\", to
        `output`{.variable}. [\[RFC5545\]](references.html#refsRFC5545)

    3.  For each element `element`{.variable} that is [a property of the
        item](#the-properties-of-an-item){#conversion-to-icalendar:the-properties-of-an-item}
        `node`{.variable}: for each name `name`{.variable} in
        `element`{.variable}\'s [property
        names](#property-names){#conversion-to-icalendar:property-names},
        run the appropriate set of substeps from the following list:

        If the property\'s [value](#concept-property-value){#conversion-to-icalendar:concept-property-value} is an [item](#concept-item){#conversion-to-icalendar:concept-item-3}

        :   Skip the property.

        If the property is [`dtend`{#conversion-to-icalendar:md-vevent-dtend}](#md-vevent-dtend)\
        If the property is [`dtstart`{#conversion-to-icalendar:md-vevent-dtstart}](#md-vevent-dtstart)\
        If the property is [`exdate`{#conversion-to-icalendar:md-vevent-exdate}](#md-vevent-exdate)\
        If the property is [`rdate`{#conversion-to-icalendar:md-vevent-rdate}](#md-vevent-rdate)\
        If the property is [`created`{#conversion-to-icalendar:md-vevent-created}](#md-vevent-created)\
        If the property is [`last-modified`{#conversion-to-icalendar:md-vevent-last-modified}](#md-vevent-last-modified)

        :   Let `value`{.variable} be the result of stripping all U+002D
            HYPHEN-MINUS (-) and U+003A COLON (:) characters from the
            property\'s
            [value](#concept-property-value){#conversion-to-icalendar:concept-property-value-2}.

            If the property\'s
            [value](#concept-property-value){#conversion-to-icalendar:concept-property-value-3}
            is a [valid date
            string](common-microsyntaxes.html#valid-date-string){#conversion-to-icalendar:valid-date-string},
            then [add an iCalendar
            line](#add-an-icalendar-line){#conversion-to-icalendar:add-an-icalendar-line-6}
            with the type `name`{.variable} and the value
            `value`{.variable} to `output`{.variable}, with the
            annotation \"`VALUE=DATE`\".

            Otherwise, if the property\'s
            [value](#concept-property-value){#conversion-to-icalendar:concept-property-value-4}
            is a [valid global date and time
            string](common-microsyntaxes.html#valid-global-date-and-time-string){#conversion-to-icalendar:valid-global-date-and-time-string},
            then [add an iCalendar
            line](#add-an-icalendar-line){#conversion-to-icalendar:add-an-icalendar-line-7}
            with the type `name`{.variable} and the value
            `value`{.variable} to `output`{.variable}, with the
            annotation \"`VALUE=DATE-TIME`\".

            Otherwise, skip the property.

        Otherwise

        :   [Add an iCalendar
            line](#add-an-icalendar-line){#conversion-to-icalendar:add-an-icalendar-line-8}
            with the type `name`{.variable} and the property\'s
            [value](#concept-property-value){#conversion-to-icalendar:concept-property-value-5}
            to `output`{.variable}.

    4.  [Add an iCalendar
        line](#add-an-icalendar-line){#conversion-to-icalendar:add-an-icalendar-line-9}
        with the type \"`END`\" and the value \"`VEVENT`\" to
        `output`{.variable}.

7.  [Add an iCalendar
    line](#add-an-icalendar-line){#conversion-to-icalendar:add-an-icalendar-line-10}
    with the type \"`END`\" and the value \"`VCALENDAR`\" to
    `output`{.variable}.

When the above algorithm says that the user agent is to [add an
iCalendar line]{#add-an-icalendar-line .dfn} consisting of a type
`type`{.variable}, a value `value`{.variable}, and optionally an
annotation, to a string `output`{.variable}, it must run the following
steps:

1.  Let `line`{.variable} be an empty string.

2.  Append `type`{.variable}, [converted to ASCII
    uppercase](https://infra.spec.whatwg.org/#ascii-uppercase){#conversion-to-icalendar:converted-to-ascii-uppercase
    x-internal="converted-to-ascii-uppercase"}, to `line`{.variable}.

3.  If there is an annotation:

    1.  Append a U+003B SEMICOLON character (;) to `line`{.variable}.

    2.  Append the annotation to `line`{.variable}.

4.  Append a U+003A COLON character (:) to `line`{.variable}.

5.  Prefix every U+005C REVERSE SOLIDUS character (\\) in
    `value`{.variable} with another U+005C REVERSE SOLIDUS character
    (\\).

6.  Prefix every U+002C COMMA character (,) in `value`{.variable} with a
    U+005C REVERSE SOLIDUS character (\\).

7.  Prefix every U+003B SEMICOLON character (;) in `value`{.variable}
    with a U+005C REVERSE SOLIDUS character (\\).

8.  Replace every U+000D CARRIAGE RETURN U+000A LINE FEED character pair
    (CRLF) in `value`{.variable} with a U+005C REVERSE SOLIDUS character
    (\\) followed by a U+006E LATIN SMALL LETTER N character (n).

9.  Replace every remaining U+000D CARRIAGE RETURN (CR) or U+000A LINE
    FEED (LF) character in `value`{.variable} with a U+005C REVERSE
    SOLIDUS character (\\) followed by a U+006E LATIN SMALL LETTER N
    character (n).

10. Append `value`{.variable} to `line`{.variable}.

11. Let `maximum length`{.variable} be 75.

12. While `line`{.variable}\'s [code point
    length](https://infra.spec.whatwg.org/#string-code-point-length){#conversion-to-icalendar:code-point-length
    x-internal="code-point-length"} is greater than
    `maximum length`{.variable}:

    1.  Append the first `maximum length`{.variable} code points of
        `line`{.variable} to `output`{.variable}.

    2.  Remove the first `maximum length`{.variable} code points from
        `line`{.variable}.

    3.  Append a U+000D CARRIAGE RETURN character (CR) to
        `output`{.variable}.

    4.  Append a U+000A LINE FEED character (LF) to `output`{.variable}.

    5.  Append a U+0020 SPACE character to `output`{.variable}.

    6.  Let `maximum length`{.variable} be 74.

13. Append (what remains of) `line`{.variable} to `output`{.variable}.

14. Append a U+000D CARRIAGE RETURN character (CR) to
    `output`{.variable}.

15. Append a U+000A LINE FEED character (LF) to `output`{.variable}.

This algorithm can generate invalid iCalendar output, if the input does
not conform to the rules described for the
[`http://microformats.org/profile/hcalendar#vevent`{#conversion-to-icalendar:md-vevent-3}](#md-vevent)
[item type](#item-types){#conversion-to-icalendar:item-types} and
[defined property
names](#defined-property-name){#conversion-to-icalendar:defined-property-name}.

##### [5.3.2.2]{.secno} Examples[](#examples-3){.self-link} {#examples-3}

*This section is non-normative.*

::: example
Here is an example of a page that uses the vEvent vocabulary to mark up
an event:

``` html
<body itemscope itemtype="http://microformats.org/profile/hcalendar#vevent">
 ...
 <h1 itemprop="summary">Bluesday Tuesday: Money Road</h1>
 ...
 <time itemprop="dtstart" datetime="2009-05-05T19:00:00Z">May 5th @ 7pm</time>
 (until <time itemprop="dtend" datetime="2009-05-05T21:00:00Z">9pm</time>)
 ...
 <a href="http://livebrum.co.uk/2009/05/05/bluesday-tuesday-money-road"
    rel="bookmark" itemprop="url">Link to this page</a>
 ...
 <p>Location: <span itemprop="location">The RoadHouse</span></p>
 ...
 <p><input type=button value="Add to Calendar"
           onclick="location = getCalendar(this)"></p>
 ...
 <meta itemprop="description" content="via livebrum.co.uk">
</body>
```

The `getCalendar()` function is left as an exercise for the reader.

The same page could offer some markup, such as the following, for
copy-and-pasting into blogs:

``` html
<div itemscope itemtype="http://microformats.org/profile/hcalendar#vevent">
 <p>I'm going to
 <strong itemprop="summary">Bluesday Tuesday: Money Road</strong>,
 <time itemprop="dtstart" datetime="2009-05-05T19:00:00Z">May 5th at 7pm</time>
 to <time itemprop="dtend" datetime="2009-05-05T21:00:00Z">9pm</time>,
 at <span itemprop="location">The RoadHouse</span>!</p>
 <p><a href="http://livebrum.co.uk/2009/05/05/bluesday-tuesday-money-road"
       itemprop="url">See this event on livebrum.co.uk</a>.</p>
 <meta itemprop="description" content="via livebrum.co.uk">
</div>
```
:::

#### [5.3.3]{.secno} Licensing works[](#licensing-works){.self-link}

An item with the [item type](#item-types){#licensing-works:item-types}
[`http://n.whatwg.org/work`]{#md-work .dfn} represents a work (e.g. an
article, an image, a video, a song, etc.). This type is primarily
intended to allow authors to include licensing information for works.

The following are the type\'s [defined property
names](#defined-property-name){#licensing-works:defined-property-name}.

[`work`]{#md-work-work .dfn}

:   Identifies the work being described.

    The
    [value](#concept-property-value){#licensing-works:concept-property-value}
    must be an [absolute
    URL](https://url.spec.whatwg.org/#syntax-url-absolute){#licensing-works:absolute-url
    x-internal="absolute-url"}.

    Exactly one property with the name
    [`work`{#licensing-works:md-work-work}](#md-work-work) must be
    present within each
    [item](#concept-item){#licensing-works:concept-item} with the type
    [`http://n.whatwg.org/work`{#licensing-works:md-work}](#md-work).

[`title`]{#md-work-title .dfn}

:   Gives the name of the work.

    A single property with the name
    [`title`{#licensing-works:md-work-title}](#md-work-title) may be
    present within each
    [item](#concept-item){#licensing-works:concept-item-2} with the type
    [`http://n.whatwg.org/work`{#licensing-works:md-work-2}](#md-work).

[`author`]{#md-work-author .dfn}

:   Gives the name or contact information of one of the authors or
    creators of the work.

    The
    [value](#concept-property-value){#licensing-works:concept-property-value-2}
    must be either an
    [item](#concept-item){#licensing-works:concept-item-3} with the type
    [`http://microformats.org/profile/hcard`{#licensing-works:md-vcard}](#md-vcard),
    or text.

    Any number of properties with the name
    [`author`{#licensing-works:md-work-author}](#md-work-author) may be
    present within each
    [item](#concept-item){#licensing-works:concept-item-4} with the type
    [`http://n.whatwg.org/work`{#licensing-works:md-work-3}](#md-work).

[`license`]{#md-work-license .dfn}

:   Identifies one of the licenses under which the work is available.

    The
    [value](#concept-property-value){#licensing-works:concept-property-value-3}
    must be an [absolute
    URL](https://url.spec.whatwg.org/#syntax-url-absolute){#licensing-works:absolute-url-2
    x-internal="absolute-url"}.

    Any number of properties with the name
    [`license`{#licensing-works:md-work-license}](#md-work-license) may
    be present within each
    [item](#concept-item){#licensing-works:concept-item-5} with the type
    [`http://n.whatwg.org/work`{#licensing-works:md-work-4}](#md-work).

##### [5.3.3.1]{.secno} Examples[](#examples-4){.self-link} {#examples-4}

*This section is non-normative.*

::: example
This example shows an embedded image entitled My Pond, licensed under
the Creative Commons Attribution-Share Alike 4.0 International License
and the MIT license simultaneously.

``` html
<figure itemscope itemtype="http://n.whatwg.org/work">
 <img itemprop="work" src="mypond.jpeg">
 <figcaption>
  <p><cite itemprop="title">My Pond</cite></p>
  <p><small>Licensed under the <a itemprop="license"
  href="https://creativecommons.org/licenses/by-sa/4.0/">Creative
  Commons Attribution-Share Alike 4.0 International License</a>
  and the <a itemprop="license"
  href="http://www.opensource.org/licenses/mit-license.php">MIT
  license</a>.</small>
 </figcaption>
</figure>
```
:::

### [5.4]{.secno} Converting HTML to other formats[](#converting-html-to-other-formats){.self-link}

#### [5.4.1]{.secno} JSON[](#json){.self-link}

Given a list of nodes `nodes`{.variable} in a
[`Document`{#json:document}](dom.html#document), a user agent must run
the following algorithm to [extract the microdata from those nodes into
a JSON form]{#extracting-json .dfn}:

1.  Let `result`{.variable} be an empty object.

2.  Let `items`{.variable} be an empty array.

3.  For each `node`{.variable} in `nodes`{.variable}, check if the
    element is a [top-level microdata
    item](#top-level-microdata-items){#json:top-level-microdata-items},
    and if it is then [get the
    object](#get-the-object){#json:get-the-object} for that element and
    add it to `items`{.variable}.

4.  Add an entry to `result`{.variable} called \"`items`\" whose value
    is the array `items`{.variable}.

5.  Return the result of serializing `result`{.variable} to JSON in the
    shortest possible way (meaning no whitespace between tokens, no
    unnecessary zero digits in numbers, and only using Unicode escapes
    in strings for characters that do not have a dedicated escape
    sequence), and with a lowercase \"`e`\" used, when appropriate, in
    the representation of any numbers.
    [\[JSON\]](references.html#refsJSON)

This algorithm returns an object with a single property that is an
array, instead of just returning an array, so that it is possible to
extend the algorithm in the future if necessary.

When the user agent is to [get the object]{#get-the-object .dfn} for an
item `item`{.variable}, optionally with a list of elements
`memory`{.variable}, it must run the following substeps:

1.  Let `result`{.variable} be an empty object.

2.  If no `memory`{.variable} was passed to the algorithm, let
    `memory`{.variable} be an empty list.

3.  Add `item`{.variable} to `memory`{.variable}.

4.  If the `item`{.variable} has any [item
    types](#item-types){#json:item-types}, add an entry to
    `result`{.variable} called \"`type`\" whose value is an array
    listing the [item types](#item-types){#json:item-types-2} of
    `item`{.variable}, in the order they were specified on the
    [`itemtype`{#json:attr-itemtype}](#attr-itemtype) attribute.

5.  If the `item`{.variable} has a [global
    identifier](#global-identifier){#json:global-identifier}, add an
    entry to `result`{.variable} called \"`id`\" whose value is the
    [global identifier](#global-identifier){#json:global-identifier-2}
    of `item`{.variable}.

6.  Let `properties`{.variable} be an empty object.

7.  For each element `element`{.variable} that has one or more [property
    names](#property-names){#json:property-names} and is one of [the
    properties of the
    item](#the-properties-of-an-item){#json:the-properties-of-an-item}
    `item`{.variable}, in the order those elements are given by the
    algorithm that returns [the properties of an
    item](#the-properties-of-an-item){#json:the-properties-of-an-item-2},
    run the following substeps:

    1.  Let `value`{.variable} be the [property
        value](#concept-property-value){#json:concept-property-value} of
        `element`{.variable}.

    2.  If `value`{.variable} is an
        [item](#concept-item){#json:concept-item}, then: If
        `value`{.variable} is in `memory`{.variable}, then let
        `value`{.variable} be the string \"`ERROR`\". Otherwise, [get
        the object](#get-the-object){#json:get-the-object-2} for
        `value`{.variable}, passing a copy of `memory`{.variable}, and
        then replace `value`{.variable} with the object returned from
        those steps.

    3.  For each name `name`{.variable} in `element`{.variable}\'s
        [property names](#property-names){#json:property-names-2}, run
        the following substeps:

        1.  If there is no entry named `name`{.variable} in
            `properties`{.variable}, then add an entry named
            `name`{.variable} to `properties`{.variable} whose value is
            an empty array.

        2.  Append `value`{.variable} to the entry named
            `name`{.variable} in `properties`{.variable}.

8.  Add an entry to `result`{.variable} called \"`properties`\" whose
    value is the object `properties`{.variable}.

9.  Return `result`{.variable}.

::: example
For example, take this markup:

``` html
<!DOCTYPE HTML>
<html lang="en">
<title>My Blog</title>
<article itemscope itemtype="http://schema.org/BlogPosting">
 <header>
  <h1 itemprop="headline">Progress report</h1>
  <p><time itemprop="datePublished" datetime="2013-08-29">today</time></p>
  <link itemprop="url" href="?comments=0">
 </header>
 <p>All in all, he's doing well with his swim lessons. The biggest thing was he had trouble
 putting his head in, but we got it down.</p>
 <section>
  <h1>Comments</h1>
  <article itemprop="comment" itemscope itemtype="http://schema.org/UserComments" id="c1">
   <link itemprop="url" href="#c1">
   <footer>
    <p>Posted by: <span itemprop="creator" itemscope itemtype="http://schema.org/Person">
     <span itemprop="name">Greg</span>
    </span></p>
    <p><time itemprop="commentTime" datetime="2013-08-29">15 minutes ago</time></p>
   </footer>
   <p>Ha!</p>
  </article>
  <article itemprop="comment" itemscope itemtype="http://schema.org/UserComments" id="c2">
   <link itemprop="url" href="#c2">
   <footer>
    <p>Posted by: <span itemprop="creator" itemscope itemtype="http://schema.org/Person">
     <span itemprop="name">Charlotte</span>
    </span></p>
    <p><time itemprop="commentTime" datetime="2013-08-29">5 minutes ago</time></p>
   </footer>
   <p>When you say "we got it down"...</p>
  </article>
 </section>
</article>
```

It would be turned into the following JSON by the algorithm above
(supposing that the page\'s URL was
`https://blog.example.com/progress-report`):

``` json
{
  "items": [
    {
      "type": [ "http://schema.org/BlogPosting" ],
      "properties": {
        "headline": [ "Progress report" ],
        "datePublished": [ "2013-08-29" ],
        "url": [ "https://blog.example.com/progress-report?comments=0" ],
        "comment": [
          {
            "type": [ "http://schema.org/UserComments" ],
            "properties": {
              "url": [ "https://blog.example.com/progress-report#c1" ],
              "creator": [
                {
                  "type": [ "http://schema.org/Person" ],
                  "properties": {
                    "name": [ "Greg" ]
                  }
                }
              ],
              "commentTime": [ "2013-08-29" ]
            }
          },
          {
            "type": [ "http://schema.org/UserComments" ],
            "properties": {
              "url": [ "https://blog.example.com/progress-report#c2" ],
              "creator": [
                {
                  "type": [ "http://schema.org/Person" ],
                  "properties": {
                    "name": [ "Charlotte" ]
                  }
                }
              ],
              "commentTime": [ "2013-08-29" ]
            }
          }
        ]
      }
    }
  ]
}
```
:::

[← 4.14 Common idioms without dedicated elements](semantics-other.html)
--- [Table of Contents](./) --- [6 User interaction →](interaction.html)
