HTML Standard

::: {#head .head .with-buttons}
[![WHATWG](https://resources.whatwg.org/logo.svg){.darkmode-aware
width="100" crossorigin="" height="100"}](https://whatwg.org/){.logo}

# HTML {#html .allcaps}

Living Standard --- Last Updated [2 June 2025]{.pubdate}
:::

[← 4.8.12 The map element](image-maps.html) --- [Table of Contents](./)
--- [4.9 Tabular data →](tables.html)

1.  ::: {#toc-semantics}
    1.  1.  [[4.8.15]{.secno}
            MathML](embedded-content-other.html#mathml)
        2.  [[4.8.16]{.secno} SVG](embedded-content-other.html#svg-0)
        3.  [[4.8.17]{.secno} Dimension
            attributes](embedded-content-other.html#dimension-attributes)
    :::

#### [4.8.15]{.secno} MathML[](#mathml){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[HTML/HTML5/HTML5_Parser#Inline_SVG_and_MathML_support](https://developer.mozilla.org/en-US/docs/Web/Guide/HTML/HTML5/HTML5_Parser#Inline_SVG_and_MathML_support "To build websites, you should know about HTML — the fundamental technology used to define the structure of a webpage. HTML is used to specify whether your web content should be recognized as a paragraph, list, heading, link, image, multimedia player, form, or one of many other available elements or even a new element that you define.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5.1+]{.safari .yes}[Chrome7+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android5+]{.firefox_android .yes}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android18+]{.chrome_android .yes}[WebView
Android3+]{.webview_android .yes}[Samsung
Internet1.0+]{.samsunginternet_android .yes}[Opera
Android12+]{.opera_android .yes}
:::
::::
:::::

The [MathML
`math`](https://w3c.github.io/mathml-core/#the-top-level-math-element){#mathml:mathml-math
x-internal="mathml-math"} element falls into the [embedded
content](dom.html#embedded-content-category){#mathml:embedded-content-category},
[phrasing
content](dom.html#phrasing-content-2){#mathml:phrasing-content-2}, [flow
content](dom.html#flow-content-2){#mathml:flow-content-2}, and [palpable
content](dom.html#palpable-content-2){#mathml:palpable-content-2}
categories for the purposes of the content models in this specification.

When the [MathML
`annotation-xml`](https://w3c.github.io/mathml-core/#dfn-annotation-xml){#mathml:mathml-annotation-xml
x-internal="mathml-annotation-xml"} element contains elements from the
[HTML
namespace](https://infra.spec.whatwg.org/#html-namespace){#mathml:html-namespace-2
x-internal="html-namespace-2"}, such elements must all be [flow
content](dom.html#flow-content-2){#mathml:flow-content-2-2}.

When the MathML token elements
([`mi`{#mathml:mathml-mi}](https://w3c.github.io/mathml-core/#the-mi-element){x-internal="mathml-mi"},
[`mo`{#mathml:mathml-mo}](https://w3c.github.io/mathml-core/#operator-fence-separator-or-accent-mo){x-internal="mathml-mo"},
[`mn`{#mathml:mathml-mn}](https://w3c.github.io/mathml-core/#number-mn){x-internal="mathml-mn"},
[`ms`{#mathml:mathml-ms}](https://w3c.github.io/mathml-core/#string-literal-ms){x-internal="mathml-ms"},
and
[`mtext`{#mathml:mathml-mtext}](https://w3c.github.io/mathml-core/#text-mtext){x-internal="mathml-mtext"})
are descendants of HTML elements, they may contain [phrasing
content](dom.html#phrasing-content-2){#mathml:phrasing-content-2-2}
elements from the [HTML
namespace](https://infra.spec.whatwg.org/#html-namespace){#mathml:html-namespace-2-2
x-internal="html-namespace-2"}.

User agents must handle text other than [inter-element
whitespace](dom.html#inter-element-whitespace){#mathml:inter-element-whitespace}
found in MathML elements whose content models do not allow straight text
by pretending for the purposes of MathML content models, layout, and
rendering that the text is actually wrapped in a [MathML
`mtext`](https://w3c.github.io/mathml-core/#text-mtext){#mathml:mathml-mtext-2
x-internal="mathml-mtext"} element. (Such text is not, however,
conforming.)

User agents must act as if any MathML element whose contents does not
match the element\'s content model was replaced, for the purposes of
MathML layout and rendering, by a [MathML
`merror`](https://w3c.github.io/mathml-core/#error-message-merror){#mathml:mathml-merror
x-internal="mathml-merror"} element containing some appropriate error
message.

The semantics of MathML elements are defined by MathML and [other
applicable
specifications](infrastructure.html#other-applicable-specifications){#mathml:other-applicable-specifications}.
[\[MATHML\]](references.html#refsMATHML)

::: example
Here is an example of the use of MathML in an HTML document:

``` html
<!DOCTYPE html>
<html lang="en">
 <head>
  <title>The quadratic formula</title>
 </head>
 <body>
  <h1>The quadratic formula</h1>
  <p>
   <math>
    <mi>x</mi>
    <mo>=</mo>
    <mfrac>
     <mrow>
      <mo form="prefix">−</mo> <mi>b</mi>
      <mo>±</mo>
      <msqrt>
       <msup> <mi>b</mi> <mn>2</mn> </msup>
       <mo>−</mo>
       <mn>4</mn> <mo>⁢</mo> <mi>a</mi> <mo>⁢</mo> <mi>c</mi>
      </msqrt>
     </mrow>
     <mrow>
      <mn>2</mn> <mo>⁢</mo> <mi>a</mi>
     </mrow>
    </mfrac>
   </math>
  </p>
 </body>
</html>
```
:::

#### [4.8.16]{.secno} SVG[](#svg-0){.self-link} {#svg-0}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[HTML/HTML5/HTML5_Parser#Inline_SVG_and_MathML_support](https://developer.mozilla.org/en-US/docs/Web/Guide/HTML/HTML5/HTML5_Parser#Inline_SVG_and_MathML_support "To build websites, you should know about HTML — the fundamental technology used to define the structure of a webpage. HTML is used to specify whether your web content should be recognized as a paragraph, list, heading, link, image, multimedia player, form, or one of many other available elements or even a new element that you define.")

Support in all current engines.

::: support
[Firefox37+]{.firefox .yes}[Safari11.1+]{.safari .yes}[Chrome7+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera15+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android37+]{.firefox_android .yes}[Safari iOS11.3+]{.safari_ios
.yes}[Chrome Android18+]{.chrome_android .yes}[WebView
Android4.4+]{.webview_android .yes}[Samsung
Internet4+]{.samsunginternet_android .yes}[Opera
Android15+]{.opera_android .yes}
:::
::::
:::::

The [SVG
`svg`](https://svgwg.org/svg2-draft/struct.html#SVGElement){#svg-0:svg-svg
x-internal="svg-svg"} element falls into the [embedded
content](dom.html#embedded-content-category){#svg-0:embedded-content-category},
[phrasing
content](dom.html#phrasing-content-2){#svg-0:phrasing-content-2}, [flow
content](dom.html#flow-content-2){#svg-0:flow-content-2}, and [palpable
content](dom.html#palpable-content-2){#svg-0:palpable-content-2}
categories for the purposes of the content models in this specification.

When the [SVG
`foreignObject`](https://svgwg.org/svg2-draft/embedded.html#ForeignObjectElement){#svg-0:svg-foreignobject
x-internal="svg-foreignobject"} element contains elements from the [HTML
namespace](https://infra.spec.whatwg.org/#html-namespace){#svg-0:html-namespace-2
x-internal="html-namespace-2"}, such elements must all be [flow
content](dom.html#flow-content-2){#svg-0:flow-content-2-2}.

The content model for the [SVG
`title`](https://svgwg.org/svg2-draft/struct.html#TitleElement){#svg-0:svg-title
x-internal="svg-title"} element inside [HTML
documents](https://dom.spec.whatwg.org/#html-document){#svg-0:html-documents
x-internal="html-documents"} is [phrasing
content](dom.html#phrasing-content-2){#svg-0:phrasing-content-2-2}.
(This further constrains the requirements given in SVG 2.)

The semantics of SVG elements are defined by SVG 2 and [other applicable
specifications](infrastructure.html#other-applicable-specifications){#svg-0:other-applicable-specifications}.
[\[SVG\]](references.html#refsSVG)

------------------------------------------------------------------------

`doc`{.variable}` = ``iframe`{.variable}`.`[`getSVGDocument`](#dom-media-getsvgdocument){#dom-media-getsvgdocument-dev}`()`\
`doc`{.variable}` = ``embed`{.variable}`.`[`getSVGDocument`](#dom-media-getsvgdocument){#svg-0:dom-media-getsvgdocument}`()`\
`doc`{.variable}` = ``object`{.variable}`.`[`getSVGDocument`](#dom-media-getsvgdocument){#svg-0:dom-media-getsvgdocument-2}`()`

:   Returns the [`Document`{#svg-0:document}](dom.html#document) object,
    in the case of
    [`iframe`{#svg-0:the-iframe-element}](iframe-embed-object.html#the-iframe-element),
    [`embed`{#svg-0:the-embed-element}](iframe-embed-object.html#the-embed-element),
    or
    [`object`{#svg-0:the-object-element}](iframe-embed-object.html#the-object-element)
    elements being used to embed SVG.

The [`getSVGDocument()`]{#dom-media-getsvgdocument .dfn
dfn-for="HTMLIFrameElement,HTMLEmbedElement,HTMLObjectElement"
dfn-type="method"} method steps are:

1.  Let `document`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#svg-0:this
    x-internal="this"}\'s [content
    document](document-sequences.html#concept-bcc-content-document){#svg-0:concept-bcc-content-document}.

2.  If `document`{.variable} is non-null and was created by the [page
    load processing model for XML
    files](document-lifecycle.html#read-xml){#svg-0:read-xml} section
    because the [computed type of the
    resource](https://mimesniff.spec.whatwg.org/#computed-mime-type){#svg-0:content-type-sniffing-2
    x-internal="content-type-sniffing-2"} in the
    [navigate](browsing-the-web.html#navigate){#svg-0:navigate}
    algorithm was
    [`image/svg+xml`{#svg-0:image/svg+xml}](indices.html#image/svg+xml),
    then return `document`{.variable}.

3.  Return null.

#### [4.8.17]{.secno} [Dimension attributes]{.dfn}[](#dimension-attributes){.self-link}

**Author requirements**: The [`width`]{#attr-dim-width .dfn
dfn-for="img,iframe,embed,object,source,video" dfn-type="element-attr"}
and [`height`]{#attr-dim-height .dfn
dfn-for="img,iframe,embed,object,source,video" dfn-type="element-attr"}
attributes on
[`img`{#dimension-attributes:the-img-element}](embedded-content.html#the-img-element),
[`iframe`{#dimension-attributes:the-iframe-element}](iframe-embed-object.html#the-iframe-element),
[`embed`{#dimension-attributes:the-embed-element}](iframe-embed-object.html#the-embed-element),
[`object`{#dimension-attributes:the-object-element}](iframe-embed-object.html#the-object-element),
[`video`{#dimension-attributes:the-video-element}](media.html#the-video-element),
[`source`{#dimension-attributes:the-source-element}](embedded-content.html#the-source-element)
when the parent is a
[`picture`{#dimension-attributes:the-picture-element}](embedded-content.html#the-picture-element)
element and, when their
[`type`{#dimension-attributes:attr-input-type}](input.html#attr-input-type)
attribute is in the [Image
Button](input.html#image-button-state-(type=image)){#dimension-attributes:image-button-state-(type=image)}
state,
[`input`{#dimension-attributes:the-input-element}](input.html#the-input-element)
elements may be specified to give the dimensions of the visual content
of the element (the width and height respectively, relative to the
nominal direction of the output medium), in [CSS
pixels](https://drafts.csswg.org/css-values/#px){#dimension-attributes:'px'
x-internal="'px'"}. The attributes, if specified, must have values that
are [valid non-negative
integers](common-microsyntaxes.html#valid-non-negative-integer){#dimension-attributes:valid-non-negative-integer}.

The specified dimensions given may differ from the dimensions specified
in the resource itself, since the resource may have a resolution that
differs from the CSS pixel resolution. (On screens, [CSS
pixels](https://drafts.csswg.org/css-values/#px){#dimension-attributes:'px'-2
x-internal="'px'"} have a resolution of 96ppi, but in general the CSS
pixel resolution depends on the reading distance.) If both attributes
are specified, then one of the following statements must be true:

- `specified width`{.variable} - 0.5 ≤ `specified height`{.variable} \*
  `target ratio`{.variable} ≤ `specified width`{.variable} + 0.5
- `specified height`{.variable} - 0.5 ≤ `specified width`{.variable} /
  `target ratio`{.variable} ≤ `specified height`{.variable} + 0.5
- `specified height`{.variable} = `specified width`{.variable} = 0

The `target ratio`{.variable} is the ratio of the [natural
width](https://drafts.csswg.org/css-images/#natural-width){#dimension-attributes:natural-width
x-internal="natural-width"} to the [natural
height](https://drafts.csswg.org/css-images/#natural-height){#dimension-attributes:natural-height
x-internal="natural-height"} in the resource. The
`specified width`{.variable} and `specified height`{.variable} are the
values of the
[`width`{#dimension-attributes:attr-dim-width}](#attr-dim-width) and
[`height`{#dimension-attributes:attr-dim-height}](#attr-dim-height)
attributes respectively.

The two attributes must be omitted if the resource in question does not
have both a [natural
width](https://drafts.csswg.org/css-images/#natural-width){#dimension-attributes:natural-width-2
x-internal="natural-width"} and a [natural
height](https://drafts.csswg.org/css-images/#natural-height){#dimension-attributes:natural-height-2
x-internal="natural-height"}.

If the two attributes are both 0, it indicates that the element is not
intended for the user (e.g. it might be a part of a service to count
page views).

The dimension attributes are not intended to be used to stretch the
image.

**User agent requirements**: User agents are expected to use these
attributes [as hints for the rendering](rendering.html#dimRendering).

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLObjectElement/width](https://developer.mozilla.org/en-US/docs/Web/API/HTMLObjectElement/width "The width property of the HTMLObjectElement interface returns a string that reflects the width HTML attribute, specifying the displayed width of the resource in CSS pixels.")

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
[HTMLObjectElement/height](https://developer.mozilla.org/en-US/docs/Web/API/HTMLObjectElement/height "The height property of the HTMLObjectElement interface Returns a string that reflects the height HTML attribute, specifying the displayed height of the resource in CSS pixels.")

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
:::::::

The [`width`]{#dom-dim-width .dfn
dfn-for="HTMLIFrameElement,HTMLEmbedElement,HTMLObjectElement,HTMLSourceElement,HTMLVideoElement"
dfn-type="attribute"} and [`height`]{#dom-dim-height .dfn
dfn-for="HTMLIFrameElement,HTMLEmbedElement,HTMLObjectElement,HTMLSourceElement,HTMLVideoElement"
dfn-type="attribute"} IDL attributes on the
[`iframe`{#dimension-attributes:the-iframe-element-2}](iframe-embed-object.html#the-iframe-element),
[`embed`{#dimension-attributes:the-embed-element-2}](iframe-embed-object.html#the-embed-element),
[`object`{#dimension-attributes:the-object-element-2}](iframe-embed-object.html#the-object-element),
[`source`{#dimension-attributes:the-source-element-2}](embedded-content.html#the-source-element),
and
[`video`{#dimension-attributes:the-video-element-2}](media.html#the-video-element)
elements must
[reflect](common-dom-interfaces.html#reflect){#dimension-attributes:reflect}
the respective content attributes of the same name.

For
[`iframe`{#dimension-attributes:the-iframe-element-3}](iframe-embed-object.html#the-iframe-element),
[`embed`{#dimension-attributes:the-embed-element-3}](iframe-embed-object.html#the-embed-element)
and
[`object`{#dimension-attributes:the-object-element-3}](iframe-embed-object.html#the-object-element)
the IDL attributes are
[`DOMString`{#dimension-attributes:idl-domstring}](https://webidl.spec.whatwg.org/#idl-DOMString){x-internal="idl-domstring"};
for
[`video`{#dimension-attributes:the-video-element-3}](media.html#the-video-element)
and
[`source`{#dimension-attributes:the-source-element-3}](embedded-content.html#the-source-element)
the IDL attributes are
[`unsigned long`{#dimension-attributes:idl-unsigned-long}](https://webidl.spec.whatwg.org/#idl-unsigned-long){x-internal="idl-unsigned-long"}.

The corresponding IDL attributes for
[`img`{#dimension-attributes:dom-img-height}](embedded-content.html#dom-img-height)
and
[`input`{#dimension-attributes:dom-input-height}](input.html#dom-input-height)
elements are defined in those respective elements\' sections, as they
are slightly more specific to those elements\' other behaviors.

[← 4.8.12 The map element](image-maps.html) --- [Table of Contents](./)
--- [4.9 Tabular data →](tables.html)
