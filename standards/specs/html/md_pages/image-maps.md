HTML Standard

::: {#head .head .with-buttons}
[![WHATWG](https://resources.whatwg.org/logo.svg){.darkmode-aware
width="100" crossorigin="" height="100"}](https://whatwg.org/){.logo}

# HTML {#html .allcaps}

Living Standard --- Last Updated [2 June 2025]{.pubdate}
:::

[← 4.8.8 The video element](media.html) --- [Table of Contents](./) ---
[4.8.15 MathML →](embedded-content-other.html)

1.  ::: {#toc-semantics}
    1.  1.  [[4.8.12]{.secno} The `map`
            element](image-maps.html#the-map-element)
        2.  [[4.8.13]{.secno} The `area`
            element](image-maps.html#the-area-element)
        3.  [[4.8.14]{.secno} Image maps](image-maps.html#image-maps)
            1.  [[4.8.14.1]{.secno}
                Authoring](image-maps.html#authoring)
            2.  [[4.8.14.2]{.secno} Processing
                model](image-maps.html#image-map-processing-model)
    :::

#### [4.8.12]{.secno} The [`map`]{.dfn dfn-type="element"} element[](#the-map-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/map](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/map "The <map> HTML element is used with <area> elements to define an image map (a clickable link area).")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
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

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMapElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMapElement "The HTMLMapElement interface provides special properties and methods (beyond those of the regular object HTMLElement interface it also has available to it by inheritance) for manipulating the layout and presentation of map elements.")

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

[Categories](dom.html#concept-element-categories){#the-map-element:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-map-element:flow-content-2}.
:   [Phrasing
    content](dom.html#phrasing-content-2){#the-map-element:phrasing-content-2}.
:   [Palpable
    content](dom.html#palpable-content-2){#the-map-element:palpable-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-map-element:concept-element-contexts}:
:   Where [phrasing
    content](dom.html#phrasing-content-2){#the-map-element:phrasing-content-2-2}
    is expected.

[Content model](dom.html#concept-element-content-model){#the-map-element:concept-element-content-model}:
:   [Transparent](dom.html#transparent){#the-map-element:transparent}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-map-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-map-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-map-element:global-attributes}
:   [`name`{#the-map-element:attr-map-name}](#attr-map-name) --- Name of
    [image map](#image-map){#the-map-element:image-map} to
    [reference](dom.html#referenced){#the-map-element:referenced} from
    the
    [`usemap`{#the-map-element:attr-hyperlink-usemap}](#attr-hyperlink-usemap)
    attribute

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-map-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-map).
:   [For implementers](https://w3c.github.io/html-aam/#el-map).

[DOM interface](dom.html#concept-element-dom){#the-map-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLMapElement : HTMLElement {
      [HTMLConstructor] constructor();

      [CEReactions] attribute DOMString name;
      [SameObject] readonly attribute HTMLCollection areas;
    };
    ```

The [`map`{#the-map-element:the-map-element}](#the-map-element) element,
in conjunction with an
[`img`{#the-map-element:the-img-element}](embedded-content.html#the-img-element)
element and any
[`area`{#the-map-element:the-area-element}](#the-area-element) element
descendants, defines an [image
map](#image-map){#the-map-element:image-map-2}. The element
[represents](dom.html#represents){#the-map-element:represents} its
children.

The [`name`]{#attr-map-name .dfn dfn-for="map" dfn-type="element-attr"}
attribute gives the map a name so that it can be
[referenced](dom.html#referenced){#the-map-element:referenced-2}. The
attribute must be present and must have a non-empty value with no [ASCII
whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#the-map-element:space-characters
x-internal="space-characters"}. The value of the
[`name`{#the-map-element:attr-map-name-2}](#attr-map-name) attribute
must not be equal to the value of the
[`name`{#the-map-element:attr-map-name-3}](#attr-map-name) attribute of
another [`map`{#the-map-element:the-map-element-2}](#the-map-element)
element in the same
[tree](https://dom.spec.whatwg.org/#concept-tree){#the-map-element:tree
x-internal="tree"}. If the
[`id`{#the-map-element:the-id-attribute}](dom.html#the-id-attribute)
attribute is also specified, both attributes must have the same value.

`map`{.variable}`.`[`areas`](#dom-map-areas){#dom-map-areas-dev}

:   Returns an
    [`HTMLCollection`{#the-map-element:htmlcollection-2}](https://dom.spec.whatwg.org/#interface-htmlcollection){x-internal="htmlcollection"}
    of the
    [`area`{#the-map-element:the-area-element-2}](#the-area-element)
    elements in the
    [`map`{#the-map-element:the-map-element-3}](#the-map-element).

The [`areas`]{#dom-map-areas .dfn dfn-for="HTMLMapElement"
dfn-type="attribute"} attribute must return an
[`HTMLCollection`{#the-map-element:htmlcollection-3}](https://dom.spec.whatwg.org/#interface-htmlcollection){x-internal="htmlcollection"}
rooted at the
[`map`{#the-map-element:the-map-element-4}](#the-map-element) element,
whose filter matches only
[`area`{#the-map-element:the-area-element-3}](#the-area-element)
elements.

The IDL attribute [`name`]{#dom-map-name .dfn dfn-for="HTMLMapElement"
dfn-type="attribute"} must
[reflect](common-dom-interfaces.html#reflect){#the-map-element:reflect}
the content attribute of the same name.

::: example
Image maps can be defined in conjunction with other content on the page,
to ease maintenance. This example is of a page with an image map at the
top of the page and a corresponding set of text links at the bottom.

``` html
<!DOCTYPE HTML>
<HTML LANG="EN">
<TITLE>Babies™: Toys</TITLE>
<HEADER>
 <H1>Toys</H1>
 <IMG SRC="/images/menu.gif"
      ALT="Babies™ navigation menu. Select a department to go to its page."
      USEMAP="#NAV">
</HEADER>
 ...
<FOOTER>
 <MAP NAME="NAV">
  <P>
   <A HREF="/clothes/">Clothes</A>
   <AREA ALT="Clothes" COORDS="0,0,100,50" HREF="/clothes/"> |
   <A HREF="/toys/">Toys</A>
   <AREA ALT="Toys" COORDS="100,0,200,50" HREF="/toys/"> |
   <A HREF="/food/">Food</A>
   <AREA ALT="Food" COORDS="200,0,300,50" HREF="/food/"> |
   <A HREF="/books/">Books</A>
   <AREA ALT="Books" COORDS="300,0,400,50" HREF="/books/">
  </P>
 </MAP>
</FOOTER>
```
:::

#### [4.8.13]{.secno} The [`area`]{.dfn dfn-type="element"} element[](#the-area-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/area](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/area "The <area> HTML element defines an area inside an image map that has predefined clickable areas. An image map allows geometric areas on an image to be associated with hypertext links.")

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
[HTMLAreaElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLAreaElement "The HTMLAreaElement interface provides special properties and methods (beyond those of the regular object HTMLElement interface it also has available to it by inheritance) for manipulating the layout and presentation of <area> elements.")

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

[Categories](dom.html#concept-element-categories){#the-area-element:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-area-element:flow-content-2}.
:   [Phrasing
    content](dom.html#phrasing-content-2){#the-area-element:phrasing-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-area-element:concept-element-contexts}:
:   Where [phrasing
    content](dom.html#phrasing-content-2){#the-area-element:phrasing-content-2-2}
    is expected, but only if there is a
    [`map`{#the-area-element:the-map-element}](#the-map-element) element
    ancestor.

[Content model](dom.html#concept-element-content-model){#the-area-element:concept-element-content-model}:
:   [Nothing](dom.html#concept-content-nothing){#the-area-element:concept-content-nothing}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-area-element:concept-element-tag-omission}:
:   No [end
    tag](syntax.html#syntax-end-tag){#the-area-element:syntax-end-tag}.

[Content attributes](dom.html#concept-element-attributes){#the-area-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-area-element:global-attributes}
:   [`alt`{#the-area-element:attr-area-alt}](#attr-area-alt) ---
    Replacement text for use when images are not available
:   [`coords`{#the-area-element:attr-area-coords}](#attr-area-coords)
    --- Coordinates for the shape to be created in an [image
    map](#image-map){#the-area-element:image-map}
:   [`shape`{#the-area-element:attr-area-shape}](#attr-area-shape) ---
    The kind of shape to be created in an [image
    map](#image-map){#the-area-element:image-map-2}
:   [`href`{#the-area-element:attr-hyperlink-href}](links.html#attr-hyperlink-href)
    --- Address of the
    [hyperlink](links.html#hyperlink){#the-area-element:hyperlink}
:   [`target`{#the-area-element:attr-hyperlink-target}](links.html#attr-hyperlink-target)
    ---
    [Navigable](document-sequences.html#navigable){#the-area-element:navigable}
    for [hyperlink](links.html#hyperlink){#the-area-element:hyperlink-2}
    [navigation](browsing-the-web.html#navigate){#the-area-element:navigate}
:   [`download`{#the-area-element:attr-hyperlink-download}](links.html#attr-hyperlink-download)
    --- Whether to download the resource instead of navigating to it,
    and its filename if so
:   [`ping`{#the-area-element:ping}](links.html#ping) ---
    [URLs](https://url.spec.whatwg.org/#concept-url){#the-area-element:url
    x-internal="url"} to ping
:   [`rel`{#the-area-element:attr-hyperlink-rel}](links.html#attr-hyperlink-rel)
    --- Relationship between the location in the document containing the
    [hyperlink](links.html#hyperlink){#the-area-element:hyperlink-3} and
    the destination resource
:   [`referrerpolicy`{#the-area-element:attr-hyperlink-referrerpolicy}](links.html#attr-hyperlink-referrerpolicy)
    --- [Referrer
    policy](https://w3c.github.io/webappsec-referrer-policy/#referrer-policy){#the-area-element:referrer-policy
    x-internal="referrer-policy"} for
    [fetches](https://fetch.spec.whatwg.org/#concept-fetch){#the-area-element:concept-fetch
    x-internal="concept-fetch"} initiated by the element

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-area-element:concept-element-accessibility-considerations}:
:   If the element has an
    [`href`{#the-area-element:attr-hyperlink-href-2}](links.html#attr-hyperlink-href)
    attribute: [for authors](https://w3c.github.io/html-aria/#el-area);
    [for implementers](https://w3c.github.io/html-aam/#el-area).
:   Otherwise: [for
    authors](https://w3c.github.io/html-aria/#el-area-no-href); [for
    implementers](https://w3c.github.io/html-aam/#el-area-no-href).

[DOM interface](dom.html#concept-element-dom){#the-area-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLAreaElement : HTMLElement {
      [HTMLConstructor] constructor();

      [CEReactions] attribute DOMString alt;
      [CEReactions] attribute DOMString coords;
      [CEReactions] attribute DOMString shape;
      [CEReactions] attribute DOMString target;
      [CEReactions] attribute DOMString download;
      [CEReactions] attribute USVString ping;
      [CEReactions] attribute DOMString rel;
      [SameObject, PutForwards=value] readonly attribute DOMTokenList relList;
      [CEReactions] attribute DOMString referrerPolicy;

      // also has obsolete members
    };
    HTMLAreaElement includes HTMLHyperlinkElementUtils;
    ```

The [`area`{#the-area-element:the-area-element}](#the-area-element)
element [represents](dom.html#represents){#the-area-element:represents}
either a hyperlink with some text and a corresponding area on an [image
map](#image-map){#the-area-element:image-map-3}, or a dead area on an
image map.

An [`area`{#the-area-element:the-area-element-2}](#the-area-element)
element with a parent node must have a
[`map`{#the-area-element:the-map-element-2}](#the-map-element) element
ancestor.

If the [`area`{#the-area-element:the-area-element-3}](#the-area-element)
element has an
[`href`{#the-area-element:attr-hyperlink-href-3}](links.html#attr-hyperlink-href)
attribute, then the
[`area`{#the-area-element:the-area-element-4}](#the-area-element)
element represents a
[hyperlink](links.html#hyperlink){#the-area-element:hyperlink-4}. In
this case, the [`alt`]{#attr-area-alt .dfn dfn-for="area"
dfn-type="element-attr"} attribute must be present. It specifies the
text of the hyperlink. Its value must be text that, when presented with
the texts specified for the other hyperlinks of the [image
map](#image-map){#the-area-element:image-map-4}, and with the
alternative text of the image, but without the image itself, provides
the user with the same kind of choice as the hyperlink would when used
without its text but with its shape applied to the image. The
[`alt`{#the-area-element:attr-area-alt-2}](#attr-area-alt) attribute may
be left blank if there is another
[`area`{#the-area-element:the-area-element-5}](#the-area-element)
element in the same [image
map](#image-map){#the-area-element:image-map-5} that points to the same
resource and has a non-blank
[`alt`{#the-area-element:attr-area-alt-3}](#attr-area-alt) attribute.

If the [`area`{#the-area-element:the-area-element-6}](#the-area-element)
element has no
[`href`{#the-area-element:attr-hyperlink-href-4}](links.html#attr-hyperlink-href)
attribute, then the area represented by the element cannot be selected,
and the [`alt`{#the-area-element:attr-area-alt-4}](#attr-area-alt)
attribute must be omitted.

In both cases, the
[`shape`{#the-area-element:attr-area-shape-2}](#attr-area-shape) and
[`coords`{#the-area-element:attr-area-coords-2}](#attr-area-coords)
attributes specify the area.

The [`shape`]{#attr-area-shape .dfn dfn-for="area"
dfn-type="element-attr"} attribute is an [enumerated
attribute](common-microsyntaxes.html#enumerated-attribute){#the-area-element:enumerated-attribute}
with the following keywords and states:

Keyword

Conforming

State

Brief description

[`circle`]{#attr-area-shape-keyword-circle .dfn}

[Circle
state](#attr-area-shape-circle){#the-area-element:attr-area-shape-circle}

Designates a circle, using exactly three integers in the
[`coords`{#the-area-element:attr-area-coords-3}](#attr-area-coords)
attribute.

[`circ`]{#attr-area-shape-keyword-circ .dfn}

No

[`default`]{#attr-area-shape-keyword-default .dfn}

[Default
state](#attr-area-shape-default){#the-area-element:attr-area-shape-default}

This area is the whole image. (The
[`coords`{#the-area-element:attr-area-coords-4}](#attr-area-coords)
attribute is not used.)

[`poly`]{#attr-area-shape-keyword-poly .dfn}

[Polygon
state](#attr-area-shape-poly){#the-area-element:attr-area-shape-poly}

Designates a polygon, using at-least six integers in the
[`coords`{#the-area-element:attr-area-coords-5}](#attr-area-coords)
attribute.

[`polygon`]{#attr-area-shape-keyword-polygon .dfn}

No

[`rect`]{#attr-area-shape-keyword-rect .dfn}

[Rectangle
state](#attr-area-shape-rect){#the-area-element:attr-area-shape-rect}

Designates a rectangle, using exactly four integers in the
[`coords`{#the-area-element:attr-area-coords-6}](#attr-area-coords)
attribute.

[`rectangle`]{#attr-area-shape-keyword-rectangle .dfn}

No

The attribute\'s *[missing value
default](common-microsyntaxes.html#missing-value-default)* and *[invalid
value default](common-microsyntaxes.html#invalid-value-default)* are
both the
[rectangle](#attr-area-shape-rect){#the-area-element:attr-area-shape-rect-2}
state.

The [`coords`]{#attr-area-coords .dfn dfn-for="area"
dfn-type="element-attr"} attribute must, if specified, contain a [valid
list of floating-point
numbers](common-microsyntaxes.html#valid-list-of-floating-point-numbers){#the-area-element:valid-list-of-floating-point-numbers}.
This attribute gives the coordinates for the shape described by the
[`shape`{#the-area-element:attr-area-shape-3}](#attr-area-shape)
attribute. The processing for this attribute is described as part of the
[image map](#image-map){#the-area-element:image-map-6} processing model.

In the [circle state]{#attr-area-shape-circle .dfn dfn-for="area/shape"
dfn-type="attr-value"},
[`area`{#the-area-element:the-area-element-7}](#the-area-element)
elements must have a
[`coords`{#the-area-element:attr-area-coords-7}](#attr-area-coords)
attribute present, with three integers, the last of which must be
non-negative. The first integer must be the distance in [CSS
pixels](https://drafts.csswg.org/css-values/#px){#the-area-element:'px'
x-internal="'px'"} from the left edge of the image to the center of the
circle, the second integer must be the distance in [CSS
pixels](https://drafts.csswg.org/css-values/#px){#the-area-element:'px'-2
x-internal="'px'"} from the top edge of the image to the center of the
circle, and the third integer must be the radius of the circle, again in
[CSS
pixels](https://drafts.csswg.org/css-values/#px){#the-area-element:'px'-3
x-internal="'px'"}.

In the [default state]{#attr-area-shape-default .dfn
dfn-for="area/shape" dfn-type="attr-value"},
[`area`{#the-area-element:the-area-element-8}](#the-area-element)
elements must not have a
[`coords`{#the-area-element:attr-area-coords-8}](#attr-area-coords)
attribute. (The area is the whole image.)

In the [polygon state]{#attr-area-shape-poly .dfn dfn-for="area/shape"
dfn-type="attr-value"},
[`area`{#the-area-element:the-area-element-9}](#the-area-element)
elements must have a
[`coords`{#the-area-element:attr-area-coords-9}](#attr-area-coords)
attribute with at least six integers, and the number of integers must be
even. Each pair of integers must represent a coordinate given as the
distances from the left and the top of the image in [CSS
pixels](https://drafts.csswg.org/css-values/#px){#the-area-element:'px'-4
x-internal="'px'"} respectively, and all the coordinates together must
represent the points of the polygon, in order.

In the [rectangle state]{#attr-area-shape-rect .dfn dfn-for="area/shape"
dfn-type="attr-value"},
[`area`{#the-area-element:the-area-element-10}](#the-area-element)
elements must have a
[`coords`{#the-area-element:attr-area-coords-10}](#attr-area-coords)
attribute with exactly four integers, the first of which must be less
than the third, and the second of which must be less than the fourth.
The four points must represent, respectively, the distance from the left
edge of the image to the left side of the rectangle, the distance from
the top edge to the top side, the distance from the left edge to the
right side, and the distance from the top edge to the bottom side, all
in [CSS
pixels](https://drafts.csswg.org/css-values/#px){#the-area-element:'px'-5
x-internal="'px'"}.

When user agents allow users to [follow
hyperlinks](links.html#following-hyperlinks-2){#the-area-element:following-hyperlinks-2}
or [download
hyperlinks](links.html#downloading-hyperlinks){#the-area-element:downloading-hyperlinks}
created using the
[`area`{#the-area-element:the-area-element-11}](#the-area-element)
element, the
[`href`{#the-area-element:attr-hyperlink-href-5}](links.html#attr-hyperlink-href),
[`target`{#the-area-element:attr-hyperlink-target-2}](links.html#attr-hyperlink-target),
[`download`{#the-area-element:attr-hyperlink-download-2}](links.html#attr-hyperlink-download),
and [`ping`{#the-area-element:ping-2}](links.html#ping) attributes
decide how the link is followed. The
[`rel`{#the-area-element:attr-hyperlink-rel-2}](links.html#attr-hyperlink-rel)
attribute may be used to indicate to the user the likely nature of the
target resource before the user follows the link.

The
[`target`{#the-area-element:attr-hyperlink-target-3}](links.html#attr-hyperlink-target),
[`download`{#the-area-element:attr-hyperlink-download-3}](links.html#attr-hyperlink-download),
[`ping`{#the-area-element:ping-3}](links.html#ping),
[`rel`{#the-area-element:attr-hyperlink-rel-3}](links.html#attr-hyperlink-rel),
and
[`referrerpolicy`{#the-area-element:attr-hyperlink-referrerpolicy-2}](links.html#attr-hyperlink-referrerpolicy)
attributes must be omitted if the
[`href`{#the-area-element:attr-hyperlink-href-6}](links.html#attr-hyperlink-href)
attribute is not present.

If the
[`itemprop`{#the-area-element:names:-the-itemprop-attribute}](microdata.html#names:-the-itemprop-attribute)
attribute is specified on an
[`area`{#the-area-element:the-area-element-12}](#the-area-element)
element, then the
[`href`{#the-area-element:attr-hyperlink-href-7}](links.html#attr-hyperlink-href)
attribute must also be specified.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLAreaElement/rel](https://developer.mozilla.org/en-US/docs/Web/API/HTMLAreaElement/rel "The HTMLAreaElement.rel property reflects the rel attribute. It is a string containing a space-separated list of link types indicating the relationship between the resource represented by the <area> element and the current document.")

Support in all current engines.

::: support
[Firefox30+]{.firefox .yes}[Safari9+]{.safari .yes}[Chrome54+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer11]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

The IDL attributes [`alt`]{#dom-area-alt .dfn dfn-for="HTMLAreaElement"
dfn-type="attribute"}, [`coords`]{#dom-area-coords .dfn
dfn-for="HTMLAreaElement" dfn-type="attribute"},
[`shape`]{#dom-area-shape .dfn dfn-for="HTMLAreaElement"
dfn-type="attribute"}, [`target`]{#dom-area-target .dfn
dfn-for="HTMLAreaElement" dfn-type="attribute"},
[`download`]{#dom-area-download .dfn dfn-for="HTMLAreaElement"
dfn-type="attribute"}, [`ping`]{#dom-area-ping .dfn
dfn-for="HTMLAreaElement" dfn-type="attribute"}, and
[`rel`]{#dom-area-rel .dfn dfn-for="HTMLAreaElement"
dfn-type="attribute"}, each must
[reflect](common-dom-interfaces.html#reflect){#the-area-element:reflect}
the respective content attributes of the same name.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLAreaElement/relList](https://developer.mozilla.org/en-US/docs/Web/API/HTMLAreaElement/relList "The HTMLAreaElement.relList read-only property reflects the rel attribute. It is a live DOMTokenList containing the set of link types indicating the relationship between the resource represented by the <area> element and the current document.")

Support in all current engines.

::: support
[Firefox30+]{.firefox .yes}[Safari9+]{.safari .yes}[Chrome65+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera41+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)18]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android41+]{.opera_android .yes}
:::
::::
:::::

The IDL attribute [`relList`]{#dom-area-rellist .dfn
dfn-for="HTMLAreaElement" dfn-type="attribute"} must
[reflect](common-dom-interfaces.html#reflect){#the-area-element:reflect-2}
the
[`rel`{#the-area-element:attr-hyperlink-rel-4}](links.html#attr-hyperlink-rel)
content attribute.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLAreaElement/referrerPolicy](https://developer.mozilla.org/en-US/docs/Web/API/HTMLAreaElement/referrerPolicy "The HTMLAreaElement.referrerPolicy property reflect the HTML referrerpolicy attribute of the <area> element defining which referrer is sent when fetching the resource.")

Support in all current engines.

::: support
[Firefox50+]{.firefox .yes}[Safari14.1+]{.safari
.yes}[Chrome52+]{.chrome .yes}

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

The IDL attribute [`referrerPolicy`]{#dom-area-referrerpolicy .dfn
dfn-for="HTMLAreaElement" dfn-type="attribute"} must
[reflect](common-dom-interfaces.html#reflect){#the-area-element:reflect-3}
the
[`referrerpolicy`{#the-area-element:attr-hyperlink-referrerpolicy-3}](links.html#attr-hyperlink-referrerpolicy)
content attribute, [limited to only known
values](common-dom-interfaces.html#limited-to-only-known-values){#the-area-element:limited-to-only-known-values}.

#### [4.8.14]{.secno} Image maps[](#image-maps){.self-link}

##### [4.8.14.1]{.secno} Authoring[](#authoring){.self-link}

An [image map]{#image-map .dfn} allows geometric areas on an image to be
associated with
[hyperlinks](links.html#hyperlink){#authoring:hyperlink}.

An image, in the form of an
[`img`{#authoring:the-img-element}](embedded-content.html#the-img-element)
element, may be associated with an image map (in the form of a
[`map`{#authoring:the-map-element}](#the-map-element) element) by
specifying a [`usemap`]{#attr-hyperlink-usemap .dfn dfn-for="img"
dfn-type="element-attr"} attribute on the
[`img`{#authoring:the-img-element-2}](embedded-content.html#the-img-element)
element. The
[`usemap`{#authoring:attr-hyperlink-usemap}](#attr-hyperlink-usemap)
attribute, if specified, must be a [valid hash-name
reference](common-microsyntaxes.html#valid-hash-name-reference){#authoring:valid-hash-name-reference}
to a [`map`{#authoring:the-map-element-2}](#the-map-element) element.

::: example
Consider an image that looks as follows:

![A line with four shapes in it, equally spaced: a red hollow box, a
green circle, a blue triangle, and a yellow four-pointed
star.](/images/sample-usemap.png){width="600" height="150"}

If we wanted just the colored areas to be clickable, we could do it as
follows:

``` html
<p>
 Please select a shape:
 <img src="shapes.png" usemap="#shapes"
      alt="Four shapes are available: a red hollow box, a green circle, a blue triangle, and a yellow four-pointed star.">
 <map name="shapes">
  <area shape=rect coords="50,50,100,100"> <!-- the hole in the red box -->
  <area shape=rect coords="25,25,125,125" href="red.html" alt="Red box.">
  <area shape=circle coords="200,75,50" href="green.html" alt="Green circle.">
  <area shape=poly coords="325,25,262,125,388,125" href="blue.html" alt="Blue triangle.">
  <area shape=poly coords="450,25,435,60,400,75,435,90,450,125,465,90,500,75,465,60"
        href="yellow.html" alt="Yellow star.">
 </map>
</p>
```
:::

##### [4.8.14.2]{.secno} []{#processing-model}Processing model[](#image-map-processing-model){.self-link} {#image-map-processing-model}

If an
[`img`{#image-map-processing-model:the-img-element}](embedded-content.html#the-img-element)
element has a
[`usemap`{#image-map-processing-model:attr-hyperlink-usemap}](#attr-hyperlink-usemap)
attribute specified, user agents must process it as follows:

1.  Parse the attribute\'s value using the [rules for parsing a
    hash-name
    reference](common-microsyntaxes.html#rules-for-parsing-a-hash-name-reference){#image-map-processing-model:rules-for-parsing-a-hash-name-reference}
    to a
    [`map`{#image-map-processing-model:the-map-element}](#the-map-element)
    element, with the element as the context node. This will return
    either an element (the `map`{.variable}) or null.

2.  If that returned null, then return. The image is not associated with
    an image map after all.

3.  Otherwise, the user agent must collect all the
    [`area`{#image-map-processing-model:the-area-element}](#the-area-element)
    elements that are descendants of the `map`{.variable}. Let those be
    the `areas`{.variable}.

Having obtained the list of
[`area`{#image-map-processing-model:the-area-element-2}](#the-area-element)
elements that form the image map (the `areas`{.variable}), interactive
user agents must process the list in one of two ways.

If the user agent intends to show the text that the
[`img`{#image-map-processing-model:the-img-element-2}](embedded-content.html#the-img-element)
element represents, then it must use the following steps.

1.  Remove all the
    [`area`{#image-map-processing-model:the-area-element-3}](#the-area-element)
    elements in `areas`{.variable} that have no
    [`href`{#image-map-processing-model:attr-hyperlink-href}](links.html#attr-hyperlink-href)
    attribute.

2.  Remove all the
    [`area`{#image-map-processing-model:the-area-element-4}](#the-area-element)
    elements in `areas`{.variable} that have no
    [`alt`{#image-map-processing-model:attr-area-alt}](#attr-area-alt)
    attribute, or whose
    [`alt`{#image-map-processing-model:attr-area-alt-2}](#attr-area-alt)
    attribute\'s value is the empty string, *if* there is another
    [`area`{#image-map-processing-model:the-area-element-5}](#the-area-element)
    element in `areas`{.variable} with the same value in the
    [`href`{#image-map-processing-model:attr-hyperlink-href-2}](links.html#attr-hyperlink-href)
    attribute and with a non-empty
    [`alt`{#image-map-processing-model:attr-area-alt-3}](#attr-area-alt)
    attribute.

3.  Each remaining
    [`area`{#image-map-processing-model:the-area-element-6}](#the-area-element)
    element in `areas`{.variable} represents a
    [hyperlink](links.html#hyperlink){#image-map-processing-model:hyperlink}.
    Those hyperlinks should all be made available to the user in a
    manner associated with the text of the
    [`img`{#image-map-processing-model:the-img-element-3}](embedded-content.html#the-img-element).

    In this context, user agents may represent
    [`area`{#image-map-processing-model:the-area-element-7}](#the-area-element)
    and
    [`img`{#image-map-processing-model:the-img-element-4}](embedded-content.html#the-img-element)
    elements with no specified `alt` attributes, or whose `alt`
    attributes are the empty string or some other non-visible text, in
    an
    [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#image-map-processing-model:implementation-defined
    x-internal="implementation-defined"} fashion intended to indicate
    the lack of suitable author-provided text.

If the user agent intends to show the image and allow interaction with
the image to select hyperlinks, then the image must be associated with a
set of layered shapes, taken from the
[`area`{#image-map-processing-model:the-area-element-8}](#the-area-element)
elements in `areas`{.variable}, in reverse [tree
order](https://dom.spec.whatwg.org/#concept-tree-order){#image-map-processing-model:tree-order
x-internal="tree-order"} (so the last specified
[`area`{#image-map-processing-model:the-area-element-9}](#the-area-element)
element in the `map`{.variable} is the bottom-most shape, and the first
element in the `map`{.variable}, in [tree
order](https://dom.spec.whatwg.org/#concept-tree-order){#image-map-processing-model:tree-order-2
x-internal="tree-order"}, is the top-most shape).

Each
[`area`{#image-map-processing-model:the-area-element-10}](#the-area-element)
element in `areas`{.variable} must be processed as follows to obtain a
shape to layer onto the image:

1.  Find the state that the element\'s
    [`shape`{#image-map-processing-model:attr-area-shape}](#attr-area-shape)
    attribute represents.

2.  Use the [rules for parsing a list of floating-point
    numbers](common-microsyntaxes.html#rules-for-parsing-a-list-of-floating-point-numbers){#image-map-processing-model:rules-for-parsing-a-list-of-floating-point-numbers}
    to parse the element\'s
    [`coords`{#image-map-processing-model:attr-area-coords}](#attr-area-coords)
    attribute, if it is present, and let the result be the
    `coords`{.variable} list. If the attribute is absent, let the
    `coords`{.variable} list be the empty list.

3.  If the number of items in the `coords`{.variable} list is less than
    the minimum number given for the
    [`area`{#image-map-processing-model:the-area-element-11}](#the-area-element)
    element\'s current state, as per the following table, then the shape
    is empty; return.

    State

    Minimum number of items

    [Circle
    state](#attr-area-shape-circle){#image-map-processing-model:attr-area-shape-circle}

    3

    [Default
    state](#attr-area-shape-default){#image-map-processing-model:attr-area-shape-default}

    0

    [Polygon
    state](#attr-area-shape-poly){#image-map-processing-model:attr-area-shape-poly}

    6

    [Rectangle
    state](#attr-area-shape-rect){#image-map-processing-model:attr-area-shape-rect}

    4

4.  Check for excess items in the `coords`{.variable} list as per the
    entry in the following list corresponding to the
    [`shape`{#image-map-processing-model:attr-area-shape-2}](#attr-area-shape)
    attribute\'s state:

    [Circle state](#attr-area-shape-circle){#image-map-processing-model:attr-area-shape-circle-2}
    :   Drop any items in the list beyond the third.

    [Default state](#attr-area-shape-default){#image-map-processing-model:attr-area-shape-default-2}
    :   Drop all items in the list.

    [Polygon state](#attr-area-shape-poly){#image-map-processing-model:attr-area-shape-poly-2}
    :   Drop the last item if there\'s an odd number of items.

    [Rectangle state](#attr-area-shape-rect){#image-map-processing-model:attr-area-shape-rect-2}
    :   Drop any items in the list beyond the fourth.

5.  If the
    [`shape`{#image-map-processing-model:attr-area-shape-3}](#attr-area-shape)
    attribute represents the [rectangle
    state](#attr-area-shape-rect){#image-map-processing-model:attr-area-shape-rect-3},
    and the first number in the list is numerically greater than the
    third number in the list, then swap those two numbers around.

6.  If the
    [`shape`{#image-map-processing-model:attr-area-shape-4}](#attr-area-shape)
    attribute represents the [rectangle
    state](#attr-area-shape-rect){#image-map-processing-model:attr-area-shape-rect-4},
    and the second number in the list is numerically greater than the
    fourth number in the list, then swap those two numbers around.

7.  If the
    [`shape`{#image-map-processing-model:attr-area-shape-5}](#attr-area-shape)
    attribute represents the [circle
    state](#attr-area-shape-circle){#image-map-processing-model:attr-area-shape-circle-3},
    and the third number in the list is less than or equal to zero, then
    the shape is empty; return.

8.  Now, the shape represented by the element is the one described for
    the entry in the list below corresponding to the state of the
    [`shape`{#image-map-processing-model:attr-area-shape-6}](#attr-area-shape)
    attribute:

    [Circle state](#attr-area-shape-circle){#image-map-processing-model:attr-area-shape-circle-4}

    :   Let `x`{.variable} be the first number in `coords`{.variable},
        `y`{.variable} be the second number, and `r`{.variable} be the
        third number.

        The shape is a circle whose center is `x`{.variable} [CSS
        pixels](https://drafts.csswg.org/css-values/#px){#image-map-processing-model:'px'
        x-internal="'px'"} from the left edge of the image and
        `y`{.variable} [CSS
        pixels](https://drafts.csswg.org/css-values/#px){#image-map-processing-model:'px'-2
        x-internal="'px'"} from the top edge of the image, and whose
        radius is `r`{.variable} [CSS
        pixels](https://drafts.csswg.org/css-values/#px){#image-map-processing-model:'px'-3
        x-internal="'px'"}.

    [Default state](#attr-area-shape-default){#image-map-processing-model:attr-area-shape-default-3}

    :   The shape is a rectangle that exactly covers the entire image.

    [Polygon state](#attr-area-shape-poly){#image-map-processing-model:attr-area-shape-poly-3}

    :   Let `x`{.variable}~`i`{.variable}~ be the (2`i`{.variable})th
        entry in `coords`{.variable}, and `y`{.variable}~`i`{.variable}~
        be the (2`i`{.variable}+1)th entry in `coords`{.variable} (the
        first entry in `coords`{.variable} being the one with index 0).

        Let `the coordinates`{.variable} be
        (`x`{.variable}~`i`{.variable}~,
        `y`{.variable}~`i`{.variable}~), interpreted in [CSS
        pixels](https://drafts.csswg.org/css-values/#px){#image-map-processing-model:'px'-4
        x-internal="'px'"} measured from the top left of the image, for
        all integer values of `i`{.variable} from 0 to
        (`N`{.variable}/2)-1, where `N`{.variable} is the number of
        items in `coords`{.variable}.

        The shape is a polygon whose vertices are given by
        `the coordinates`{.variable}, and whose interior is established
        using the even-odd rule.
        [\[GRAPHICS\]](references.html#refsGRAPHICS)

    [Rectangle state](#attr-area-shape-rect){#image-map-processing-model:attr-area-shape-rect-5}

    :   Let `x`{.variable}~`1`{.variable}~ be the first number in
        `coords`{.variable}, `y`{.variable}~`1`{.variable}~ be the
        second number, `x`{.variable}~`2`{.variable}~ be the third
        number, and `y`{.variable}~`2`{.variable}~ be the fourth number.

        The shape is a rectangle whose top-left corner is given by the
        coordinate (`x`{.variable}~`1`{.variable}~,
        `y`{.variable}~`1`{.variable}~) and whose bottom right corner is
        given by the coordinate (`x`{.variable}~`2`{.variable}~,
        `y`{.variable}~`2`{.variable}~), those coordinates being
        interpreted as [CSS
        pixels](https://drafts.csswg.org/css-values/#px){#image-map-processing-model:'px'-5
        x-internal="'px'"} from the top left corner of the image.

    For historical reasons, the coordinates must be interpreted relative
    to the *displayed* image after any stretching caused by the CSS
    [\'width\'](https://drafts.csswg.org/css2/#the-width-property){#image-map-processing-model:'width'
    x-internal="'width'"} and
    [\'height\'](https://drafts.csswg.org/css2/#the-height-property){#image-map-processing-model:'height'
    x-internal="'height'"} properties (or, for non-CSS browsers, the
    image element\'s `width` and `height` attributes --- CSS browsers
    map those attributes to the aforementioned CSS properties).

    Browser zoom features and transforms applied using CSS or SVG do not
    affect the coordinates.

Pointing device interaction with an image associated with a set of
layered shapes per the above algorithm must result in the relevant user
interaction events being first fired to the top-most shape covering the
point that the pointing device indicated, if any, or to the image
element itself, if there is no shape covering that point. User agents
may also allow individual
[`area`{#image-map-processing-model:the-area-element-12}](#the-area-element)
elements representing
[hyperlinks](links.html#hyperlink){#image-map-processing-model:hyperlink-2}
to be selected and activated (e.g. using a keyboard).

Because a
[`map`{#image-map-processing-model:the-map-element-2}](#the-map-element)
element (and its
[`area`{#image-map-processing-model:the-area-element-13}](#the-area-element)
elements) can be associated with multiple
[`img`{#image-map-processing-model:the-img-element-5}](embedded-content.html#the-img-element)
elements, it is possible for an
[`area`{#image-map-processing-model:the-area-element-14}](#the-area-element)
element to correspond to multiple *[focusable
areas](interaction.html#focusable-area)* of the document.

Image maps are
[live](infrastructure.html#live){#image-map-processing-model:live}; if
the DOM is mutated, then the user agent must act as if it had rerun the
algorithms for image maps.

[← 4.8.8 The video element](media.html) --- [Table of Contents](./) ---
[4.8.15 MathML →](embedded-content-other.html)
