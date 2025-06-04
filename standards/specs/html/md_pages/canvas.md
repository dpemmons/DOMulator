HTML Standard

::: {#head .head .with-buttons}
[![WHATWG](https://resources.whatwg.org/logo.svg){.darkmode-aware
width="100" crossorigin="" height="100"}](https://whatwg.org/){.logo}

# HTML {#html .allcaps}

Living Standard --- Last Updated [2 June 2025]{.pubdate}
:::

[← 4.12 Scripting](scripting.html) --- [Table of Contents](./) --- [4.13
Custom elements →](custom-elements.html)

1.  ::: {#toc-semantics}
    1.  1.  [[4.12.5]{.secno} The `canvas`
            element](canvas.html#the-canvas-element)
            1.  [[4.12.5.1]{.secno} The 2D rendering
                context](canvas.html#2dcontext)
                1.  [[4.12.5.1.1]{.secno} Implementation
                    notes](canvas.html#implementation-notes)
                2.  [[4.12.5.1.2]{.secno} The canvas
                    settings](canvas.html#the-canvas-settings)
                3.  [[4.12.5.1.3]{.secno} The canvas
                    state](canvas.html#the-canvas-state)
                4.  [[4.12.5.1.4]{.secno} Line
                    styles](canvas.html#line-styles)
                5.  [[4.12.5.1.5]{.secno} Text
                    styles](canvas.html#text-styles)
                6.  [[4.12.5.1.6]{.secno} Building
                    paths](canvas.html#building-paths)
                7.  [[4.12.5.1.7]{.secno} `Path2D`
                    objects](canvas.html#path2d-objects)
                8.  [[4.12.5.1.8]{.secno}
                    Transformations](canvas.html#transformations)
                9.  [[4.12.5.1.9]{.secno} Image sources for 2D rendering
                    contexts](canvas.html#image-sources-for-2d-rendering-contexts)
                10. [[4.12.5.1.10]{.secno} Fill and stroke
                    styles](canvas.html#fill-and-stroke-styles)
                11. [[4.12.5.1.11]{.secno} Drawing rectangles to the
                    bitmap](canvas.html#drawing-rectangles-to-the-bitmap)
                12. [[4.12.5.1.12]{.secno} Drawing text to the
                    bitmap](canvas.html#drawing-text-to-the-bitmap)
                13. [[4.12.5.1.13]{.secno} Drawing paths to the
                    canvas](canvas.html#drawing-paths-to-the-canvas)
                14. [[4.12.5.1.14]{.secno} Drawing focus
                    rings](canvas.html#drawing-focus-rings-and-scrolling-paths-into-view)
                15. [[4.12.5.1.15]{.secno} Drawing
                    images](canvas.html#drawing-images)
                16. [[4.12.5.1.16]{.secno} Pixel
                    manipulation](canvas.html#pixel-manipulation)
                17. [[4.12.5.1.17]{.secno}
                    Compositing](canvas.html#compositing)
                18. [[4.12.5.1.18]{.secno} Image
                    smoothing](canvas.html#image-smoothing)
                19. [[4.12.5.1.19]{.secno} Shadows](canvas.html#shadows)
                20. [[4.12.5.1.20]{.secno} Filters](canvas.html#filters)
                21. [[4.12.5.1.21]{.secno} Working with
                    externally-defined SVG
                    filters](canvas.html#working-with-externally-defined-svg-filters)
                22. [[4.12.5.1.22]{.secno} Drawing
                    model](canvas.html#drawing-model)
                23. [[4.12.5.1.23]{.secno} Best
                    practices](canvas.html#best-practices)
                24. [[4.12.5.1.24]{.secno}
                    Examples](canvas.html#examples)
            2.  [[4.12.5.2]{.secno} The `ImageBitmap` rendering
                context](canvas.html#the-imagebitmap-rendering-context)
                1.  [[4.12.5.2.1]{.secno}
                    Introduction](canvas.html#introduction-6)
                2.  [[4.12.5.2.2]{.secno} The
                    `ImageBitmapRenderingContext`
                    interface](canvas.html#the-imagebitmaprenderingcontext-interface)
            3.  [[4.12.5.3]{.secno} The `OffscreenCanvas`
                interface](canvas.html#the-offscreencanvas-interface)
                1.  [[4.12.5.3.1]{.secno} The offscreen 2D rendering
                    context](canvas.html#the-offscreen-2d-rendering-context)
            4.  [[4.12.5.4]{.secno} Color spaces and color space
                conversion](canvas.html#colour-spaces-and-colour-correction)
            5.  [[4.12.5.5]{.secno} Serializing bitmaps to a
                file](canvas.html#serialising-bitmaps-to-a-file)
            6.  [[4.12.5.6]{.secno} Security with `canvas`
                elements](canvas.html#security-with-canvas-elements)
            7.  [[4.12.5.7]{.secno} Premultiplied alpha and the 2D
                rendering
                context](canvas.html#premultiplied-alpha-and-the-2d-rendering-context)
    :::

#### [4.12.5]{.secno} The [`canvas`]{#canvas .dfn dfn-type="element"} element[](#the-canvas-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/canvas](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/canvas "Use the HTML <canvas> element with either the canvas scripting API or the WebGL API to draw graphics and animations.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari2+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera9+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLCanvasElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLCanvasElement "The HTMLCanvasElement interface provides properties and methods for manipulating the layout and presentation of <canvas> elements. The HTMLCanvasElement interface also inherits the properties and methods of the HTMLElement interface.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari2+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera9+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

[Categories](dom.html#concept-element-categories){#the-canvas-element:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-canvas-element:flow-content-2}.
:   [Phrasing
    content](dom.html#phrasing-content-2){#the-canvas-element:phrasing-content-2}.
:   [Embedded
    content](dom.html#embedded-content-category){#the-canvas-element:embedded-content-category}.
:   [Palpable
    content](dom.html#palpable-content-2){#the-canvas-element:palpable-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-canvas-element:concept-element-contexts}:
:   Where [embedded
    content](dom.html#embedded-content-category){#the-canvas-element:embedded-content-category-2}
    is expected.

[Content model](dom.html#concept-element-content-model){#the-canvas-element:concept-element-content-model}:
:   [Transparent](dom.html#transparent){#the-canvas-element:transparent},
    but with no [interactive
    content](dom.html#interactive-content-2){#the-canvas-element:interactive-content-2}
    descendants except for
    [`a`{#the-canvas-element:the-a-element}](text-level-semantics.html#the-a-element)
    elements,
    [`img`{#the-canvas-element:the-img-element}](embedded-content.html#the-img-element)
    elements with
    [`usemap`{#the-canvas-element:attr-hyperlink-usemap}](image-maps.html#attr-hyperlink-usemap)
    attributes,
    [`button`{#the-canvas-element:the-button-element}](form-elements.html#the-button-element)
    elements,
    [`input`{#the-canvas-element:the-input-element}](input.html#the-input-element)
    elements whose
    [`type`{#the-canvas-element:attr-input-type}](input.html#attr-input-type)
    attribute are in the
    [Checkbox](input.html#checkbox-state-(type=checkbox)){#the-canvas-element:checkbox-state-(type=checkbox)}
    or [Radio
    Button](input.html#radio-button-state-(type=radio)){#the-canvas-element:radio-button-state-(type=radio)}
    states,
    [`input`{#the-canvas-element:the-input-element-2}](input.html#the-input-element)
    elements that are
    [buttons](forms.html#concept-button){#the-canvas-element:concept-button},
    and
    [`select`{#the-canvas-element:the-select-element}](form-elements.html#the-select-element)
    elements with a
    [`multiple`{#the-canvas-element:attr-select-multiple}](form-elements.html#attr-select-multiple)
    attribute or a [display
    size](form-elements.html#concept-select-size){#the-canvas-element:concept-select-size}
    greater than 1.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-canvas-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-canvas-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-canvas-element:global-attributes}
:   [`width`{#the-canvas-element:attr-canvas-width}](#attr-canvas-width)
    --- Horizontal dimension
:   [`height`{#the-canvas-element:attr-canvas-height}](#attr-canvas-height)
    --- Vertical dimension

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-canvas-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-canvas).
:   [For implementers](https://w3c.github.io/html-aam/#el-canvas).

[DOM interface](dom.html#concept-element-dom){#the-canvas-element:concept-element-dom}:

:   ``` idl
    typedef (CanvasRenderingContext2D or ImageBitmapRenderingContext or WebGLRenderingContext or WebGL2RenderingContext or GPUCanvasContext) RenderingContext;

    [Exposed=Window]
    interface HTMLCanvasElement : HTMLElement {
      [HTMLConstructor] constructor();

      [CEReactions] attribute unsigned long width;
      [CEReactions] attribute unsigned long height;

      RenderingContext? getContext(DOMString contextId, optional any options = null);

      USVString toDataURL(optional DOMString type = "image/png", optional any quality);
      undefined toBlob(BlobCallback _callback, optional DOMString type = "image/png", optional any quality);
      OffscreenCanvas transferControlToOffscreen();
    };

    callback BlobCallback = undefined (Blob? blob);
    ```

The
[`canvas`{#the-canvas-element:the-canvas-element}](#the-canvas-element)
element provides scripts with a resolution-dependent bitmap canvas,
which can be used for rendering graphs, game graphics, art, or other
visual images on the fly.

Authors should not use the
[`canvas`{#the-canvas-element:the-canvas-element-2}](#the-canvas-element)
element in a document when a more suitable element is available. For
example, it is inappropriate to use a
[`canvas`{#the-canvas-element:the-canvas-element-3}](#the-canvas-element)
element to render a page heading: if the desired presentation of the
heading is graphically intense, it should be marked up using appropriate
elements (typically
[`h1`{#the-canvas-element:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements))
and then styled using CSS and supporting technologies such as [shadow
trees](https://dom.spec.whatwg.org/#concept-shadow-tree){#the-canvas-element:shadow-tree
x-internal="shadow-tree"}.

When authors use the
[`canvas`{#the-canvas-element:the-canvas-element-4}](#the-canvas-element)
element, they must also provide content that, when presented to the
user, conveys essentially the same function or purpose as the
[`canvas`{#the-canvas-element:the-canvas-element-5}](#the-canvas-element)\'s
bitmap. This content may be placed as content of the
[`canvas`{#the-canvas-element:the-canvas-element-6}](#the-canvas-element)
element. The contents of the
[`canvas`{#the-canvas-element:the-canvas-element-7}](#the-canvas-element)
element, if any, are the element\'s [fallback
content](dom.html#fallback-content){#the-canvas-element:fallback-content}.

------------------------------------------------------------------------

In interactive visual media, if [scripting is
enabled](webappapis.html#concept-n-script){#the-canvas-element:concept-n-script}
for the
[`canvas`{#the-canvas-element:the-canvas-element-8}](#the-canvas-element)
element, and if support for
[`canvas`{#the-canvas-element:the-canvas-element-9}](#the-canvas-element)
elements has been enabled, then the
[`canvas`{#the-canvas-element:the-canvas-element-10}](#the-canvas-element)
element
[represents](dom.html#represents){#the-canvas-element:represents}
[embedded
content](dom.html#embedded-content-category){#the-canvas-element:embedded-content-category-3}
consisting of a dynamically created image, the element\'s bitmap.

In non-interactive, static, visual media, if the
[`canvas`{#the-canvas-element:the-canvas-element-11}](#the-canvas-element)
element has been previously associated with a rendering context (e.g. if
the page was viewed in an interactive visual medium and is now being
printed, or if some script that ran during the page layout process
painted on the element), then the
[`canvas`{#the-canvas-element:the-canvas-element-12}](#the-canvas-element)
element
[represents](dom.html#represents){#the-canvas-element:represents-2}
[embedded
content](dom.html#embedded-content-category){#the-canvas-element:embedded-content-category-4}
with the element\'s current bitmap and size. Otherwise, the element
represents its [fallback
content](dom.html#fallback-content){#the-canvas-element:fallback-content-2}
instead.

In non-visual media, and in visual media if [scripting is
disabled](webappapis.html#concept-n-noscript){#the-canvas-element:concept-n-noscript}
for the
[`canvas`{#the-canvas-element:the-canvas-element-13}](#the-canvas-element)
element or if support for
[`canvas`{#the-canvas-element:the-canvas-element-14}](#the-canvas-element)
elements has been disabled, the
[`canvas`{#the-canvas-element:the-canvas-element-15}](#the-canvas-element)
element
[represents](dom.html#represents){#the-canvas-element:represents-3} its
[fallback
content](dom.html#fallback-content){#the-canvas-element:fallback-content-3}
instead.

When a
[`canvas`{#the-canvas-element:the-canvas-element-16}](#the-canvas-element)
element
[represents](dom.html#represents){#the-canvas-element:represents-4}
[embedded
content](dom.html#embedded-content-category){#the-canvas-element:embedded-content-category-5},
the user can still focus descendants of the
[`canvas`{#the-canvas-element:the-canvas-element-17}](#the-canvas-element)
element (in the [fallback
content](dom.html#fallback-content){#the-canvas-element:fallback-content-4}).
When an element is
[focused](interaction.html#focused){#the-canvas-element:focused}, it is
the target of keyboard interaction events (even though the element
itself is not visible). This allows authors to make an interactive
canvas keyboard-accessible: authors should have a one-to-one mapping of
interactive regions to *[focusable
areas](interaction.html#focusable-area)* in the [fallback
content](dom.html#fallback-content){#the-canvas-element:fallback-content-5}.
(Focus has no effect on mouse interaction events.)
[\[UIEVENTS\]](references.html#refsUIEVENTS)

An element whose nearest
[`canvas`{#the-canvas-element:the-canvas-element-18}](#the-canvas-element)
element ancestor is [being
rendered](rendering.html#being-rendered){#the-canvas-element:being-rendered}
and [represents](dom.html#represents){#the-canvas-element:represents-5}
[embedded
content](dom.html#embedded-content-category){#the-canvas-element:embedded-content-category-6}
is an element that is [being used as relevant canvas fallback
content]{#being-used-as-relevant-canvas-fallback-content .dfn}.

------------------------------------------------------------------------

The
[`canvas`{#the-canvas-element:the-canvas-element-19}](#the-canvas-element)
element has two attributes to control the size of the element\'s bitmap:
[`width`]{#attr-canvas-width .dfn dfn-for="canvas"
dfn-type="element-attr"} and [`height`]{#attr-canvas-height .dfn
dfn-for="canvas" dfn-type="element-attr"}. These attributes, when
specified, must have values that are [valid non-negative
integers](common-microsyntaxes.html#valid-non-negative-integer){#the-canvas-element:valid-non-negative-integer}.
The [rules for parsing non-negative
integers](common-microsyntaxes.html#rules-for-parsing-non-negative-integers){#the-canvas-element:rules-for-parsing-non-negative-integers}
must be used to [obtain their numeric values]{#obtain-numeric-values
.dfn}. If an attribute is missing, or if parsing its value returns an
error, then the default value must be used instead. The
[`width`{#the-canvas-element:attr-canvas-width-2}](#attr-canvas-width)
attribute defaults to 300, and the
[`height`{#the-canvas-element:attr-canvas-height-2}](#attr-canvas-height)
attribute defaults to 150.

When setting the value of the
[`width`{#the-canvas-element:attr-canvas-width-3}](#attr-canvas-width)
or
[`height`{#the-canvas-element:attr-canvas-height-3}](#attr-canvas-height)
attribute, if the [context
mode](#concept-canvas-context-mode){#the-canvas-element:concept-canvas-context-mode}
of the
[`canvas`{#the-canvas-element:the-canvas-element-20}](#the-canvas-element)
element is set to
[placeholder](#concept-canvas-placeholder){#the-canvas-element:concept-canvas-placeholder},
the user agent must throw an
[\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-canvas-element:invalidstateerror
x-internal="invalidstateerror"}
[`DOMException`{#the-canvas-element:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
and leave the attribute\'s value unchanged.

The [natural
dimensions](https://drafts.csswg.org/css-images/#natural-dimensions){#the-canvas-element:natural-dimensions
x-internal="natural-dimensions"} of the
[`canvas`{#the-canvas-element:the-canvas-element-21}](#the-canvas-element)
element when it
[represents](dom.html#represents){#the-canvas-element:represents-6}
[embedded
content](dom.html#embedded-content-category){#the-canvas-element:embedded-content-category-7}
are equal to the dimensions of the element\'s bitmap.

The user agent must use a square pixel density consisting of one pixel
of image data per coordinate space unit for the bitmaps of a
[`canvas`{#the-canvas-element:the-canvas-element-22}](#the-canvas-element)
and its rendering contexts.

A
[`canvas`{#the-canvas-element:the-canvas-element-23}](#the-canvas-element)
element can be sized arbitrarily by a style sheet, its bitmap is then
subject to the
[\'object-fit\'](https://drafts.csswg.org/css-images/#the-object-fit){#the-canvas-element:'object-fit'
x-internal="'object-fit'"} CSS property.

------------------------------------------------------------------------

The bitmaps of
[`canvas`{#the-canvas-element:the-canvas-element-24}](#the-canvas-element)
elements, the bitmaps of
[`ImageBitmap`{#the-canvas-element:imagebitmap}](imagebitmap-and-animations.html#imagebitmap)
objects, as well as some of the bitmaps of rendering contexts, such as
those described in the sections on the
[`CanvasRenderingContext2D`{#the-canvas-element:canvasrenderingcontext2d-2}](#canvasrenderingcontext2d),
[`OffscreenCanvasRenderingContext2D`{#the-canvas-element:offscreencanvasrenderingcontext2d}](#offscreencanvasrenderingcontext2d),
and
[`ImageBitmapRenderingContext`{#the-canvas-element:imagebitmaprenderingcontext-2}](#imagebitmaprenderingcontext)
objects below, have an [origin-clean]{#concept-canvas-origin-clean .dfn}
flag, which can be set to true or false. Initially, when the
[`canvas`{#the-canvas-element:the-canvas-element-25}](#the-canvas-element)
element or
[`ImageBitmap`{#the-canvas-element:imagebitmap-2}](imagebitmap-and-animations.html#imagebitmap)
object is created, its bitmap\'s
[origin-clean](#concept-canvas-origin-clean){#the-canvas-element:concept-canvas-origin-clean}
flag must be set to true.

A
[`canvas`{#the-canvas-element:the-canvas-element-26}](#the-canvas-element)
element can have a rendering context bound to it. Initially, it does not
have a bound rendering context. To keep track of whether it has a
rendering context or not, and what kind of rendering context it is, a
[`canvas`{#the-canvas-element:the-canvas-element-27}](#the-canvas-element)
also has a [canvas context mode]{#concept-canvas-context-mode .dfn},
which is initially [none]{#concept-canvas-none .dfn} but can be changed
to either [placeholder]{#concept-canvas-placeholder .dfn},
[2d]{#concept-canvas-2d .dfn},
[bitmaprenderer]{#concept-canvas-bitmaprenderer .dfn},
[webgl]{#concept-canvas-webgl .dfn}, [webgl2]{#concept-canvas-webgl2
.dfn}, or [webgpu]{#concept-canvas-webgpu .dfn} by algorithms defined in
this specification.

When its [canvas context
mode](#concept-canvas-context-mode){#the-canvas-element:concept-canvas-context-mode-2}
is
[none](#concept-canvas-none){#the-canvas-element:concept-canvas-none}, a
[`canvas`{#the-canvas-element:the-canvas-element-28}](#the-canvas-element)
element has no rendering context, and its bitmap must be [transparent
black](https://drafts.csswg.org/css-color/#transparent-black){#the-canvas-element:transparent-black
x-internal="transparent-black"} with a [natural
width](https://drafts.csswg.org/css-images/#natural-width){#the-canvas-element:natural-width
x-internal="natural-width"} equal to [the numeric
value](#obtain-numeric-values){#the-canvas-element:obtain-numeric-values}
of the element\'s
[`width`{#the-canvas-element:attr-canvas-width-4}](#attr-canvas-width)
attribute and a [natural
height](https://drafts.csswg.org/css-images/#natural-height){#the-canvas-element:natural-height
x-internal="natural-height"} equal to [the numeric
value](#obtain-numeric-values){#the-canvas-element:obtain-numeric-values-2}
of the element\'s
[`height`{#the-canvas-element:attr-canvas-height-4}](#attr-canvas-height)
attribute, those values being interpreted in [CSS
pixels](https://drafts.csswg.org/css-values/#px){#the-canvas-element:'px'
x-internal="'px'"}, and being updated as the attributes are set,
changed, or removed.

When its [canvas context
mode](#concept-canvas-context-mode){#the-canvas-element:concept-canvas-context-mode-3}
is
[placeholder](#concept-canvas-placeholder){#the-canvas-element:concept-canvas-placeholder-2},
a
[`canvas`{#the-canvas-element:the-canvas-element-29}](#the-canvas-element)
element has no rendering context. It serves as a placeholder for an
[`OffscreenCanvas`{#the-canvas-element:offscreencanvas-2}](#offscreencanvas)
object, and the content of the
[`canvas`{#the-canvas-element:the-canvas-element-30}](#the-canvas-element)
element is updated by the
[`OffscreenCanvas`{#the-canvas-element:offscreencanvas-3}](#offscreencanvas)
object\'s rendering context.

When a
[`canvas`{#the-canvas-element:the-canvas-element-31}](#the-canvas-element)
element represents [embedded
content](dom.html#embedded-content-category){#the-canvas-element:embedded-content-category-8},
it provides a [paint
source](https://drafts.csswg.org/css-images-4/#paint-source){#the-canvas-element:paint-source
x-internal="paint-source"} whose width is the element\'s [natural
width](https://drafts.csswg.org/css-images/#natural-width){#the-canvas-element:natural-width-2
x-internal="natural-width"}, whose height is the element\'s [natural
height](https://drafts.csswg.org/css-images/#natural-height){#the-canvas-element:natural-height-2
x-internal="natural-height"}, and whose appearance is the element\'s
bitmap.

Whenever the
[`width`{#the-canvas-element:attr-canvas-width-5}](#attr-canvas-width)
and
[`height`{#the-canvas-element:attr-canvas-height-5}](#attr-canvas-height)
content attributes are set, removed, changed, or redundantly set to the
value they already have, then the user agent must perform the action
from the row of the following table that corresponds to the
[`canvas`{#the-canvas-element:the-canvas-element-32}](#the-canvas-element)
element\'s [context
mode](#concept-canvas-context-mode){#the-canvas-element:concept-canvas-context-mode-4}.

[Context
Mode](#concept-canvas-context-mode){#the-canvas-element:concept-canvas-context-mode-5}

Action

[2d](#concept-canvas-2d){#the-canvas-element:concept-canvas-2d}

Follow the steps to [set bitmap
dimensions](#concept-canvas-set-bitmap-dimensions){#the-canvas-element:concept-canvas-set-bitmap-dimensions}
to [the numeric
values](#obtain-numeric-values){#the-canvas-element:obtain-numeric-values-3}
of the
[`width`{#the-canvas-element:attr-canvas-width-6}](#attr-canvas-width)
and
[`height`{#the-canvas-element:attr-canvas-height-6}](#attr-canvas-height)
content attributes.

[webgl](#concept-canvas-webgl){#the-canvas-element:concept-canvas-webgl}
or
[webgl2](#concept-canvas-webgl2){#the-canvas-element:concept-canvas-webgl2}

Follow the behavior defined in the WebGL specifications.
[\[WEBGL\]](references.html#refsWEBGL)

[webgpu](#concept-canvas-webgpu){#the-canvas-element:concept-canvas-webgpu}

Follow the behavior defined in WebGPU.
[\[WEBGPU\]](references.html#refsWEBGPU)

[bitmaprenderer](#concept-canvas-bitmaprenderer){#the-canvas-element:concept-canvas-bitmaprenderer}

If the context\'s [bitmap
mode](#concept-imagebitmaprenderingcontext-bitmap-mode){#the-canvas-element:concept-imagebitmaprenderingcontext-bitmap-mode}
is set to
[blank](#concept-imagebitmaprenderingcontext-blank){#the-canvas-element:concept-imagebitmaprenderingcontext-blank},
run the steps to [set an `ImageBitmapRenderingContext`\'s output
bitmap](#set-an-imagebitmaprenderingcontext's-output-bitmap){#the-canvas-element:set-an-imagebitmaprenderingcontext's-output-bitmap},
passing the
[`canvas`{#the-canvas-element:the-canvas-element-33}](#the-canvas-element)
element\'s rendering context.

[placeholder](#concept-canvas-placeholder){#the-canvas-element:concept-canvas-placeholder-3}

Do nothing.

[none](#concept-canvas-none){#the-canvas-element:concept-canvas-none-2}

Do nothing.

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLCanvasElement/height](https://developer.mozilla.org/en-US/docs/Web/API/HTMLCanvasElement/height "The HTMLCanvasElement.height property is a positive integer reflecting the height HTML attribute of the <canvas> element interpreted in CSS pixels. When the attribute is not specified, or if it is set to an invalid value, like a negative, the default value of 150 is used.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera9+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::

:::: feature
[HTMLCanvasElement/width](https://developer.mozilla.org/en-US/docs/Web/API/HTMLCanvasElement/width "The HTMLCanvasElement.width property is a positive integer reflecting the width HTML attribute of the <canvas> element interpreted in CSS pixels. When the attribute is not specified, or if it is set to an invalid value, like a negative, the default value of 300 is used.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera9+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::::

The [`width`]{#dom-canvas-width .dfn dfn-for="HTMLCanvasElement"
dfn-type="attribute"} and [`height`]{#dom-canvas-height .dfn
dfn-for="HTMLCanvasElement" dfn-type="attribute"} IDL attributes must
[reflect](common-dom-interfaces.html#reflect){#the-canvas-element:reflect}
the respective content attributes of the same name, with the same
defaults.

------------------------------------------------------------------------

`context`{.variable}` = ``canvas`{.variable}`.`[`getContext`](#dom-canvas-getcontext){#dom-canvas-getcontext-dev}`(``contextId`{.variable}` [, ``options`{.variable}` ])`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLCanvasElement/getContext](https://developer.mozilla.org/en-US/docs/Web/API/HTMLCanvasElement/getContext "The HTMLCanvasElement.getContext() method returns a drawing context on the canvas, or null if the context identifier is not supported, or the canvas has already been set to a different context mode.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari2+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera9+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

Returns an object that exposes an API for drawing on the canvas.
`contextId`{.variable} specifies the desired API:
\"[`2d`{#the-canvas-element:canvas-context-2d}](#canvas-context-2d)\",
\"[`bitmaprenderer`{#the-canvas-element:canvas-context-bitmaprenderer}](#canvas-context-bitmaprenderer)\",
\"[`webgl`{#the-canvas-element:canvas-context-webgl}](#canvas-context-webgl)\",
\"[`webgl2`{#the-canvas-element:canvas-context-webgl2}](#canvas-context-webgl2)\",
or
\"[`webgpu`{#the-canvas-element:canvas-context-webgpu}](#canvas-context-webgpu)\".
`options`{.variable} is handled by that API.

This specification defines the
\"[`2d`{#the-canvas-element:canvas-context-2d-2}](#canvas-context-2d)\"
and
\"[`bitmaprenderer`{#the-canvas-element:canvas-context-bitmaprenderer-2}](#canvas-context-bitmaprenderer)\"
contexts below. The WebGL specifications define the
\"[`webgl`{#the-canvas-element:canvas-context-webgl-2}](#canvas-context-webgl)\"
and
\"[`webgl2`{#the-canvas-element:canvas-context-webgl2-2}](#canvas-context-webgl2)\"
contexts. WebGPU defines the
\"[`webgpu`{#the-canvas-element:canvas-context-webgpu-2}](#canvas-context-webgpu)\"
context. [\[WEBGL\]](references.html#refsWEBGL)
[\[WEBGPU\]](references.html#refsWEBGPU)

Returns null if `contextId`{.variable} is not supported, or if the
canvas has already been initialized with another context type (e.g.,
trying to get a
\"[`2d`{#the-canvas-element:canvas-context-2d-3}](#canvas-context-2d)\"
context after getting a
\"[`webgl`{#the-canvas-element:canvas-context-webgl-3}](#canvas-context-webgl)\"
context).

The
[`getContext(``contextId`{.variable}`, ``options`{.variable}`)`]{#dom-canvas-getcontext
.dfn dfn-for="HTMLCanvasElement" dfn-type="method"} method of the
[`canvas`{#the-canvas-element:the-canvas-element-34}](#the-canvas-element)
element, when invoked, must run these steps:

1.  If `options`{.variable} is not an
    [object](https://webidl.spec.whatwg.org/#idl-object){#the-canvas-element:idl-object
    x-internal="idl-object"}, then set `options`{.variable} to null.

2.  Set `options`{.variable} to the result of
    [converting](https://webidl.spec.whatwg.org/#es-type-mapping){#the-canvas-element:concept-idl-convert
    x-internal="concept-idl-convert"} `options`{.variable} to a
    JavaScript value.

3.  Run the steps in the cell of the following table whose column header
    matches this
    [`canvas`{#the-canvas-element:the-canvas-element-35}](#the-canvas-element)
    element\'s [canvas context
    mode](#concept-canvas-context-mode){#the-canvas-element:concept-canvas-context-mode-6}
    and whose row header matches `contextId`{.variable}:

    [none](#concept-canvas-none){#the-canvas-element:concept-canvas-none-3}

    [2d](#concept-canvas-2d){#the-canvas-element:concept-canvas-2d-2}

    [bitmaprenderer](#concept-canvas-bitmaprenderer){#the-canvas-element:concept-canvas-bitmaprenderer-2}

    [webgl](#concept-canvas-webgl){#the-canvas-element:concept-canvas-webgl-2}
    or
    [webgl2](#concept-canvas-webgl2){#the-canvas-element:concept-canvas-webgl2-2}

    [webgpu](#concept-canvas-webgpu){#the-canvas-element:concept-canvas-webgpu-2}

    [placeholder](#concept-canvas-placeholder){#the-canvas-element:concept-canvas-placeholder-4}

    \"[`2d`]{#canvas-context-2d .dfn}\"

    1.  Let `context`{.variable} be the result of running the [2D
        context creation
        algorithm](#2d-context-creation-algorithm){#the-canvas-element:2d-context-creation-algorithm}
        given
        [this](https://webidl.spec.whatwg.org/#this){#the-canvas-element:this
        x-internal="this"} and `options`{.variable}.

    2.  Set
        [this](https://webidl.spec.whatwg.org/#this){#the-canvas-element:this-2
        x-internal="this"}\'s [context
        mode](#concept-canvas-context-mode){#the-canvas-element:concept-canvas-context-mode-7}
        to
        [2d](#concept-canvas-2d){#the-canvas-element:concept-canvas-2d-3}.

    3.  Return `context`{.variable}.

    Return the same object as was returned the last time the method was
    invoked with this same first argument.

    Return null.

    Return null.

    Return null.

    Throw an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-canvas-element:invalidstateerror-2
    x-internal="invalidstateerror"}
    [`DOMException`{#the-canvas-element:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    \"[`bitmaprenderer`]{#canvas-context-bitmaprenderer .dfn}\"

    1.  Let `context`{.variable} be the result of running the
        [`ImageBitmapRenderingContext` creation
        algorithm](#imagebitmaprenderingcontext-creation-algorithm){#the-canvas-element:imagebitmaprenderingcontext-creation-algorithm}
        given
        [this](https://webidl.spec.whatwg.org/#this){#the-canvas-element:this-3
        x-internal="this"} and `options`{.variable}.

    2.  Set
        [this](https://webidl.spec.whatwg.org/#this){#the-canvas-element:this-4
        x-internal="this"}\'s [context
        mode](#concept-canvas-context-mode){#the-canvas-element:concept-canvas-context-mode-8}
        to
        [bitmaprenderer](#concept-canvas-bitmaprenderer){#the-canvas-element:concept-canvas-bitmaprenderer-3}.

    3.  Return `context`{.variable}.

    Return null.

    Return the same object as was returned the last time the method was
    invoked with this same first argument.

    Return null.

    Return null.

    Throw an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-canvas-element:invalidstateerror-3
    x-internal="invalidstateerror"}
    [`DOMException`{#the-canvas-element:domexception-3}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    \"[`webgl`]{#canvas-context-webgl .dfn}\" or
    \"[`webgl2`]{#canvas-context-webgl2 .dfn}\", if the user agent
    supports the WebGL feature in its current configuration

    1.  Let `context`{.variable} be the result of following the
        instructions given in the WebGL specifications\' *Context
        Creation* sections. [\[WEBGL\]](references.html#refsWEBGL)

    2.  If `context`{.variable} is null, then return null; otherwise set
        [this](https://webidl.spec.whatwg.org/#this){#the-canvas-element:this-5
        x-internal="this"}\'s [context
        mode](#concept-canvas-context-mode){#the-canvas-element:concept-canvas-context-mode-9}
        to
        [webgl](#concept-canvas-webgl){#the-canvas-element:concept-canvas-webgl-3}
        or
        [webgl2](#concept-canvas-webgl2){#the-canvas-element:concept-canvas-webgl2-3}.

    3.  Return `context`{.variable}.

    Return null.

    Return null.

    Return the same object as was returned the last time the method was
    invoked with this same first argument.

    Return null.

    Throw an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-canvas-element:invalidstateerror-4
    x-internal="invalidstateerror"}
    [`DOMException`{#the-canvas-element:domexception-4}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    \"[`webgpu`]{#canvas-context-webgpu .dfn}\", if the user agent
    supports the WebGPU feature in its current configuration

    1.  Let `context`{.variable} be the result of following the
        instructions given in WebGPU\'s [Canvas
        Rendering](https://gpuweb.github.io/gpuweb/#canvas-rendering)
        section. [\[WEBGPU\]](references.html#refsWEBGPU)

    2.  If `context`{.variable} is null, then return null; otherwise set
        [this](https://webidl.spec.whatwg.org/#this){#the-canvas-element:this-6
        x-internal="this"}\'s [context
        mode](#concept-canvas-context-mode){#the-canvas-element:concept-canvas-context-mode-10}
        to
        [webgpu](#concept-canvas-webgpu){#the-canvas-element:concept-canvas-webgpu-3}.

    3.  Return `context`{.variable}.

    Return null.

    Return null.

    Return null.

    Return the same object as was returned the last time the method was
    invoked with this same first argument.

    Throw an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-canvas-element:invalidstateerror-5
    x-internal="invalidstateerror"}
    [`DOMException`{#the-canvas-element:domexception-5}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    An unsupported value\*

    Return null.

    Return null.

    Return null.

    Return null.

    Return null.

    Throw an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-canvas-element:invalidstateerror-6
    x-internal="invalidstateerror"}
    [`DOMException`{#the-canvas-element:domexception-6}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    [\* For example, the
    \"[`webgl`{#the-canvas-element:canvas-context-webgl-4}](#canvas-context-webgl)\"
    or
    \"[`webgl2`{#the-canvas-element:canvas-context-webgl2-3}](#canvas-context-webgl2)\"
    value in the case of a user agent having exhausted the graphics
    hardware\'s abilities and having no software fallback
    implementation.]{.small}

------------------------------------------------------------------------

`url`{.variable}` = ``canvas`{.variable}`.`[`toDataURL`](#dom-canvas-todataurl){#dom-canvas-todataurl-dev}`([ ``type`{.variable}` [, ``quality`{.variable}` ] ])`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLCanvasElement/toDataURL](https://developer.mozilla.org/en-US/docs/Web/API/HTMLCanvasElement/toDataURL "The HTMLCanvasElement.toDataURL() method returns a data URL containing a representation of the image in the format specified by the type parameter.")

Support in all current engines.

::: support
[Firefox2+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome2+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera9+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

Returns a [`data:`
URL](https://www.rfc-editor.org/rfc/rfc2397#section-2){#the-canvas-element:data-protocol
x-internal="data-protocol"} for the image in the canvas.

The first argument, if provided, controls the type of the image to be
returned (e.g. PNG or JPEG). The default is
\"[`image/png`{#the-canvas-element:image/png}](indices.html#image/png)\";
that type is also used if the given type isn\'t supported. The second
argument applies if the type is an image format that supports variable
quality (such as
\"[`image/jpeg`{#the-canvas-element:image/jpeg}](indices.html#image/jpeg)\"),
and is a number in the range 0.0 to 1.0 inclusive indicating the desired
quality level for the resulting image.

When trying to use types other than
\"[`image/png`{#the-canvas-element:image/png-2}](indices.html#image/png)\",
authors can check if the image was really returned in the requested
format by checking to see if the returned string starts with one of the
exact strings \"`data:image/png,`\" or \"`data:image/png;`\". If it
does, the image is PNG, and thus the requested type was not supported.
(The one exception to this is if the canvas has either no height or no
width, in which case the result might simply be \"`data:,`\".)

`canvas`{.variable}`.`[`toBlob`](#dom-canvas-toblob){#dom-canvas-toblob-dev}`(``callback`{.variable}` [, ``type`{.variable}` [, quality ] ])`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLCanvasElement/toBlob](https://developer.mozilla.org/en-US/docs/Web/API/HTMLCanvasElement/toBlob "The HTMLCanvasElement.toBlob() method creates a Blob object representing the image contained in the canvas. This file may be cached on the disk or stored in memory at the discretion of the user agent.")

Support in all current engines.

::: support
[Firefox18+]{.firefox .yes}[Safari11+]{.safari .yes}[Chrome50+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)No]{.edge .no}[Internet Explorer[🔰
10+]{title="Requires a prefix or alternative name."}]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

Creates a
[`Blob`{#the-canvas-element:blob-2}](https://w3c.github.io/FileAPI/#dfn-Blob){x-internal="blob"}
object representing a file containing the image in the canvas, and
invokes a callback with a handle to that object.

The second argument, if provided, controls the type of the image to be
returned (e.g. PNG or JPEG). The default is
\"[`image/png`{#the-canvas-element:image/png-3}](indices.html#image/png)\";
that type is also used if the given type isn\'t supported. The third
argument applies if the type is an image format that supports variable
quality (such as
\"[`image/jpeg`{#the-canvas-element:image/jpeg-2}](indices.html#image/jpeg)\"),
and is a number in the range 0.0 to 1.0 inclusive indicating the desired
quality level for the resulting image.

`canvas`{.variable}`.`[`transferControlToOffscreen`](#dom-canvas-transfercontroltooffscreen){#dom-canvas-transfercontroltooffscreen-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLCanvasElement/transferControlToOffscreen](https://developer.mozilla.org/en-US/docs/Web/API/HTMLCanvasElement/transferControlToOffscreen "The HTMLCanvasElement.transferControlToOffscreen() method transfers control to an OffscreenCanvas object, either on the main thread or on a worker.")

Support in all current engines.

::: support
[Firefox105+]{.firefox .yes}[Safari16.4+]{.safari
.yes}[Chrome69+]{.chrome .yes}

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

Returns a newly created
[`OffscreenCanvas`{#the-canvas-element:offscreencanvas-4}](#offscreencanvas)
object that uses the
[`canvas`{#the-canvas-element:the-canvas-element-36}](#the-canvas-element)
element as a placeholder. Once the
[`canvas`{#the-canvas-element:the-canvas-element-37}](#the-canvas-element)
element has become a placeholder for an
[`OffscreenCanvas`{#the-canvas-element:offscreencanvas-5}](#offscreencanvas)
object, its natural size can no longer be changed, and it cannot have a
rendering context. The content of the placeholder canvas is updated on
the
[`OffscreenCanvas`{#the-canvas-element:offscreencanvas-6}](#offscreencanvas)\'s
[relevant
agent](webappapis.html#relevant-agent){#the-canvas-element:relevant-agent}\'s
[event
loop](webappapis.html#concept-agent-event-loop){#the-canvas-element:concept-agent-event-loop}\'s
[update the
rendering](webappapis.html#update-the-rendering){#the-canvas-element:update-the-rendering}
steps.

The
[`toDataURL(``type`{.variable}`, ``quality`{.variable}`)`]{#dom-canvas-todataurl
.dfn dfn-for="HTMLCanvasElement" dfn-type="method"} method, when
invoked, must run these steps:

1.  If this
    [`canvas`{#the-canvas-element:the-canvas-element-38}](#the-canvas-element)
    element\'s bitmap\'s
    [origin-clean](#concept-canvas-origin-clean){#the-canvas-element:concept-canvas-origin-clean-2}
    flag is set to false, then throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#the-canvas-element:securityerror
    x-internal="securityerror"}
    [`DOMException`{#the-canvas-element:domexception-7}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  If this
    [`canvas`{#the-canvas-element:the-canvas-element-39}](#the-canvas-element)
    element\'s bitmap has no pixels (i.e. either its horizontal
    dimension or its vertical dimension is zero), then return the string
    \"`data:,`\". (This is the shortest [`data:`
    URL](https://www.rfc-editor.org/rfc/rfc2397#section-2){#the-canvas-element:data-protocol-2
    x-internal="data-protocol"}; it represents the empty string in a
    `text/plain` resource.)

3.  Let `file`{.variable} be [a serialization of this `canvas`
    element\'s bitmap as a
    file](#a-serialisation-of-the-bitmap-as-a-file){#the-canvas-element:a-serialisation-of-the-bitmap-as-a-file},
    passing `type`{.variable} and `quality`{.variable} if given.

4.  If `file`{.variable} is null, then return \"`data:,`\".

5.  Return a [`data:`
    URL](https://www.rfc-editor.org/rfc/rfc2397#section-2){#the-canvas-element:data-protocol-3
    x-internal="data-protocol"} representing `file`{.variable}.
    [\[RFC2397\]](references.html#refsRFC2397)

The
[`toBlob(``callback`{.variable}`, ``type`{.variable}`, ``quality`{.variable}`)`]{#dom-canvas-toblob
.dfn dfn-for="HTMLCanvasElement" dfn-type="method"} method, when
invoked, must run these steps:

1.  If this
    [`canvas`{#the-canvas-element:the-canvas-element-40}](#the-canvas-element)
    element\'s bitmap\'s
    [origin-clean](#concept-canvas-origin-clean){#the-canvas-element:concept-canvas-origin-clean-3}
    flag is set to false, then throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#the-canvas-element:securityerror-2
    x-internal="securityerror"}
    [`DOMException`{#the-canvas-element:domexception-8}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  Let `result`{.variable} be null.

3.  If this
    [`canvas`{#the-canvas-element:the-canvas-element-41}](#the-canvas-element)
    element\'s bitmap has pixels (i.e., neither its horizontal dimension
    nor its vertical dimension is zero), then set `result`{.variable} to
    a copy of this
    [`canvas`{#the-canvas-element:the-canvas-element-42}](#the-canvas-element)
    element\'s bitmap.

4.  Run these steps [in
    parallel](infrastructure.html#in-parallel){#the-canvas-element:in-parallel}:

    1.  If `result`{.variable} is non-null, then set `result`{.variable}
        to [a serialization of `result`{.variable} as a
        file](#a-serialisation-of-the-bitmap-as-a-file){#the-canvas-element:a-serialisation-of-the-bitmap-as-a-file-2}
        with `type`{.variable} and `quality`{.variable} if given.

    2.  [Queue an element
        task](webappapis.html#queue-an-element-task){#the-canvas-element:queue-an-element-task}
        on the [canvas blob serialization task
        source]{#canvas-blob-serialisation-task-source .dfn} given the
        [`canvas`{#the-canvas-element:the-canvas-element-43}](#the-canvas-element)
        element to run these steps:

        1.  If `result`{.variable} is non-null, then set
            `result`{.variable} to a new
            [`Blob`{#the-canvas-element:blob-3}](https://w3c.github.io/FileAPI/#dfn-Blob){x-internal="blob"}
            object, created in the [relevant
            realm](webappapis.html#concept-relevant-realm){#the-canvas-element:concept-relevant-realm}
            of this
            [`canvas`{#the-canvas-element:the-canvas-element-44}](#the-canvas-element)
            element, representing `result`{.variable}.
            [\[FILEAPI\]](references.html#refsFILEAPI)

        2.  [Invoke](https://webidl.spec.whatwg.org/#invoke-a-callback-function){#the-canvas-element:es-invoking-callback-functions
            x-internal="es-invoking-callback-functions"}
            `callback`{.variable} with « `result`{.variable} » and
            \"`report`\".

The
[`transferControlToOffscreen()`]{#dom-canvas-transfercontroltooffscreen
.dfn dfn-for="HTMLCanvasElement" dfn-type="method"} method, when
invoked, must run these steps:

1.  If this
    [`canvas`{#the-canvas-element:the-canvas-element-45}](#the-canvas-element)
    element\'s [context
    mode](#concept-canvas-context-mode){#the-canvas-element:concept-canvas-context-mode-11}
    is not set to
    [none](#concept-canvas-none){#the-canvas-element:concept-canvas-none-4},
    throw an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-canvas-element:invalidstateerror-7
    x-internal="invalidstateerror"}
    [`DOMException`{#the-canvas-element:domexception-9}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  Let `offscreenCanvas`{.variable} be a new
    [`OffscreenCanvas`{#the-canvas-element:offscreencanvas-7}](#offscreencanvas)
    object with its width and height equal to the values of the
    [`width`{#the-canvas-element:attr-canvas-width-7}](#attr-canvas-width)
    and
    [`height`{#the-canvas-element:attr-canvas-height-7}](#attr-canvas-height)
    content attributes of this
    [`canvas`{#the-canvas-element:the-canvas-element-46}](#the-canvas-element)
    element.

3.  Set the `offscreenCanvas`{.variable}\'s [placeholder `canvas`
    element](#offscreencanvas-placeholder){#the-canvas-element:offscreencanvas-placeholder}
    to a weak reference to this
    [`canvas`{#the-canvas-element:the-canvas-element-47}](#the-canvas-element)
    element.

4.  Set this
    [`canvas`{#the-canvas-element:the-canvas-element-48}](#the-canvas-element)
    element\'s [context
    mode](#concept-canvas-context-mode){#the-canvas-element:concept-canvas-context-mode-12}
    to
    [placeholder](#concept-canvas-placeholder){#the-canvas-element:concept-canvas-placeholder-5}.

5.  Set the `offscreenCanvas`{.variable}\'s [inherited
    language](#offscreencanvas-inherited-lang){#the-canvas-element:offscreencanvas-inherited-lang}
    to the [language](dom.html#language){#the-canvas-element:language}
    of this
    [`canvas`{#the-canvas-element:the-canvas-element-49}](#the-canvas-element)
    element.

6.  Set the `offscreenCanvas`{.variable}\'s [inherited
    direction](#offscreencanvas-inherited-direction){#the-canvas-element:offscreencanvas-inherited-direction}
    to the
    [directionality](dom.html#the-directionality){#the-canvas-element:the-directionality}
    of this
    [`canvas`{#the-canvas-element:the-canvas-element-50}](#the-canvas-element)
    element.

7.  Return `offscreenCanvas`{.variable}.

##### [4.12.5.1]{.secno} The 2D rendering context[](#2dcontext){.self-link} {#2dcontext}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[CanvasRenderingContext2D](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D "The CanvasRenderingContext2D interface, part of the Canvas API, provides the 2D rendering context for the drawing surface of a <canvas> element. It is used for drawing shapes, text, images, and other objects.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari2+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera9+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

:::::::::: {.mdn-anno .wrapped .before}
MDN

::: feature
[CanvasImageSource](https://developer.mozilla.org/en-US/docs/Web/API/CanvasImageSource "The CanvasRenderingContext2D interface, part of the Canvas API, provides the 2D rendering context for the drawing surface of a <canvas> element. It is used for drawing shapes, text, images, and other objects.")
:::

:::: feature
[CanvasGradient](https://developer.mozilla.org/en-US/docs/Web/API/CanvasGradient "The CanvasGradient interface represents an opaque object describing a gradient. It is returned by the methods CanvasRenderingContext2D.createLinearGradient(), CanvasRenderingContext2D.createConicGradient() or CanvasRenderingContext2D.createRadialGradient().")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari2+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera9+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android4+]{.firefox_android .yes}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::

:::: feature
[CanvasPattern](https://developer.mozilla.org/en-US/docs/Web/API/CanvasPattern "The CanvasPattern interface represents an opaque object describing a pattern, based on an image, a canvas, or a video, created by the CanvasRenderingContext2D.createPattern() method.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari2+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera9+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::

:::: feature
[TextMetrics](https://developer.mozilla.org/en-US/docs/Web/API/TextMetrics "The TextMetrics interface represents the dimensions of a piece of text in the canvas; a TextMetrics instance can be retrieved using the CanvasRenderingContext2D.measureText() method.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome2+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera9+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android31+]{.firefox_android .yes}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
::::::::::

``` idl
typedef (HTMLImageElement or
         SVGImageElement) HTMLOrSVGImageElement;

typedef (HTMLOrSVGImageElement or
         HTMLVideoElement or
         HTMLCanvasElement or
         ImageBitmap or
         OffscreenCanvas or
         VideoFrame) CanvasImageSource;

enum PredefinedColorSpace { "srgb", "display-p3" };

enum CanvasColorType { "unorm8", "float16" };

enum CanvasFillRule { "nonzero", "evenodd" };

dictionary CanvasRenderingContext2DSettings {
  boolean alpha = true;
  boolean desynchronized = false;
  PredefinedColorSpace colorSpace = "srgb";
  CanvasColorType colorType = "unorm8";
  boolean willReadFrequently = false;
};

enum ImageSmoothingQuality { "low", "medium", "high" };

[Exposed=Window]
interface CanvasRenderingContext2D {
  // back-reference to the canvas
  readonly attribute HTMLCanvasElement canvas;
};
CanvasRenderingContext2D includes CanvasSettings;
CanvasRenderingContext2D includes CanvasState;
CanvasRenderingContext2D includes CanvasTransform;
CanvasRenderingContext2D includes CanvasCompositing;
CanvasRenderingContext2D includes CanvasImageSmoothing;
CanvasRenderingContext2D includes CanvasFillStrokeStyles;
CanvasRenderingContext2D includes CanvasShadowStyles;
CanvasRenderingContext2D includes CanvasFilters;
CanvasRenderingContext2D includes CanvasRect;
CanvasRenderingContext2D includes CanvasDrawPath;
CanvasRenderingContext2D includes CanvasUserInterface;
CanvasRenderingContext2D includes CanvasText;
CanvasRenderingContext2D includes CanvasDrawImage;
CanvasRenderingContext2D includes CanvasImageData;
CanvasRenderingContext2D includes CanvasPathDrawingStyles;
CanvasRenderingContext2D includes CanvasTextDrawingStyles;
CanvasRenderingContext2D includes CanvasPath;

interface mixin CanvasSettings {
  // settings
  CanvasRenderingContext2DSettings getContextAttributes();
};

interface mixin CanvasState {
  // state
  undefined save(); // push state on state stack
  undefined restore(); // pop state stack and restore state
  undefined reset(); // reset the rendering context to its default state
  boolean isContextLost(); // return whether context is lost
};

interface mixin CanvasTransform {
  // transformations (default transform is the identity matrix)
  undefined scale(unrestricted double x, unrestricted double y);
  undefined rotate(unrestricted double angle);
  undefined translate(unrestricted double x, unrestricted double y);
  undefined transform(unrestricted double a, unrestricted double b, unrestricted double c, unrestricted double d, unrestricted double e, unrestricted double f);

  [NewObject] DOMMatrix getTransform();
  undefined setTransform(unrestricted double a, unrestricted double b, unrestricted double c, unrestricted double d, unrestricted double e, unrestricted double f);
  undefined setTransform(optional DOMMatrix2DInit transform = {});
  undefined resetTransform();

};

interface mixin CanvasCompositing {
  // compositing
  attribute unrestricted double globalAlpha; // (default 1.0)
  attribute DOMString globalCompositeOperation; // (default "source-over")
};

interface mixin CanvasImageSmoothing {
  // image smoothing
  attribute boolean imageSmoothingEnabled; // (default true)
  attribute ImageSmoothingQuality imageSmoothingQuality; // (default low)

};

interface mixin CanvasFillStrokeStyles {
  // colors and styles (see also the CanvasPathDrawingStyles and CanvasTextDrawingStyles interfaces)
  attribute (DOMString or CanvasGradient or CanvasPattern) strokeStyle; // (default black)
  attribute (DOMString or CanvasGradient or CanvasPattern) fillStyle; // (default black)
  CanvasGradient createLinearGradient(double x0, double y0, double x1, double y1);
  CanvasGradient createRadialGradient(double x0, double y0, double r0, double x1, double y1, double r1);
  CanvasGradient createConicGradient(double startAngle, double x, double y);
  CanvasPattern? createPattern(CanvasImageSource image, [LegacyNullToEmptyString] DOMString repetition);

};

interface mixin CanvasShadowStyles {
  // shadows
  attribute unrestricted double shadowOffsetX; // (default 0)
  attribute unrestricted double shadowOffsetY; // (default 0)
  attribute unrestricted double shadowBlur; // (default 0)
  attribute DOMString shadowColor; // (default transparent black)
};

interface mixin CanvasFilters {
  // filters
  attribute DOMString filter; // (default "none")
};

interface mixin CanvasRect {
  // rects
  undefined clearRect(unrestricted double x, unrestricted double y, unrestricted double w, unrestricted double h);
  undefined fillRect(unrestricted double x, unrestricted double y, unrestricted double w, unrestricted double h);
  undefined strokeRect(unrestricted double x, unrestricted double y, unrestricted double w, unrestricted double h);
};

interface mixin CanvasDrawPath {
  // path API (see also CanvasPath)
  undefined beginPath();
  undefined fill(optional CanvasFillRule fillRule = "nonzero");
  undefined fill(Path2D path, optional CanvasFillRule fillRule = "nonzero");
  undefined stroke();
  undefined stroke(Path2D path);
  undefined clip(optional CanvasFillRule fillRule = "nonzero");
  undefined clip(Path2D path, optional CanvasFillRule fillRule = "nonzero");
  boolean isPointInPath(unrestricted double x, unrestricted double y, optional CanvasFillRule fillRule = "nonzero");
  boolean isPointInPath(Path2D path, unrestricted double x, unrestricted double y, optional CanvasFillRule fillRule = "nonzero");
  boolean isPointInStroke(unrestricted double x, unrestricted double y);
  boolean isPointInStroke(Path2D path, unrestricted double x, unrestricted double y);
};

interface mixin CanvasUserInterface {
  undefined drawFocusIfNeeded(Element element);
  undefined drawFocusIfNeeded(Path2D path, Element element);
};

interface mixin CanvasText {
  // text (see also the CanvasPathDrawingStyles and CanvasTextDrawingStyles interfaces)
  undefined fillText(DOMString text, unrestricted double x, unrestricted double y, optional unrestricted double maxWidth);
  undefined strokeText(DOMString text, unrestricted double x, unrestricted double y, optional unrestricted double maxWidth);
  TextMetrics measureText(DOMString text);
};

interface mixin CanvasDrawImage {
  // drawing images
  undefined drawImage(CanvasImageSource image, unrestricted double dx, unrestricted double dy);
  undefined drawImage(CanvasImageSource image, unrestricted double dx, unrestricted double dy, unrestricted double dw, unrestricted double dh);
  undefined drawImage(CanvasImageSource image, unrestricted double sx, unrestricted double sy, unrestricted double sw, unrestricted double sh, unrestricted double dx, unrestricted double dy, unrestricted double dw, unrestricted double dh);
};

interface mixin CanvasImageData {
  // pixel manipulation
  ImageData createImageData([EnforceRange] long sw, [EnforceRange] long sh, optional ImageDataSettings settings = {});
  ImageData createImageData(ImageData imageData);
  ImageData getImageData([EnforceRange] long sx, [EnforceRange] long sy, [EnforceRange] long sw, [EnforceRange] long sh, optional ImageDataSettings settings = {});
  undefined putImageData(ImageData imageData, [EnforceRange] long dx, [EnforceRange] long dy);
  undefined putImageData(ImageData imageData, [EnforceRange] long dx, [EnforceRange] long dy, [EnforceRange] long dirtyX, [EnforceRange] long dirtyY, [EnforceRange] long dirtyWidth, [EnforceRange] long dirtyHeight);
};

enum CanvasLineCap { "butt", "round", "square" };
enum CanvasLineJoin { "round", "bevel", "miter" };
enum CanvasTextAlign { "start", "end", "left", "right", "center" };
enum CanvasTextBaseline { "top", "hanging", "middle", "alphabetic", "ideographic", "bottom" };
enum CanvasDirection { "ltr", "rtl", "inherit" };
enum CanvasFontKerning { "auto", "normal", "none" };
enum CanvasFontStretch { "ultra-condensed", "extra-condensed", "condensed", "semi-condensed", "normal", "semi-expanded", "expanded", "extra-expanded", "ultra-expanded" };
enum CanvasFontVariantCaps { "normal", "small-caps", "all-small-caps", "petite-caps", "all-petite-caps", "unicase", "titling-caps" };
enum CanvasTextRendering { "auto", "optimizeSpeed", "optimizeLegibility", "geometricPrecision" };

interface mixin CanvasPathDrawingStyles {
  // line caps/joins
  attribute unrestricted double lineWidth; // (default 1)
  attribute CanvasLineCap lineCap; // (default "butt")
  attribute CanvasLineJoin lineJoin; // (default "miter")
  attribute unrestricted double miterLimit; // (default 10)

  // dashed lines
  undefined setLineDash(sequence<unrestricted double> segments); // default empty
  sequence<unrestricted double> getLineDash();
  attribute unrestricted double lineDashOffset;
};

interface mixin CanvasTextDrawingStyles {
  // text
  attribute DOMString lang; // (default: "inherit")
  attribute DOMString font; // (default 10px sans-serif)
  attribute CanvasTextAlign textAlign; // (default: "start")
  attribute CanvasTextBaseline textBaseline; // (default: "alphabetic")
  attribute CanvasDirection direction; // (default: "inherit")
  attribute DOMString letterSpacing; // (default: "0px")
  attribute CanvasFontKerning fontKerning; // (default: "auto")
  attribute CanvasFontStretch fontStretch; // (default: "normal")
  attribute CanvasFontVariantCaps fontVariantCaps; // (default: "normal")
  attribute CanvasTextRendering textRendering; // (default: "auto")
  attribute DOMString wordSpacing; // (default: "0px")
};

interface mixin CanvasPath {
  // shared path API methods
  undefined closePath();
  undefined moveTo(unrestricted double x, unrestricted double y);
  undefined lineTo(unrestricted double x, unrestricted double y);
  undefined quadraticCurveTo(unrestricted double cpx, unrestricted double cpy, unrestricted double x, unrestricted double y);
  undefined bezierCurveTo(unrestricted double cp1x, unrestricted double cp1y, unrestricted double cp2x, unrestricted double cp2y, unrestricted double x, unrestricted double y);
  undefined arcTo(unrestricted double x1, unrestricted double y1, unrestricted double x2, unrestricted double y2, unrestricted double radius); 
  undefined rect(unrestricted double x, unrestricted double y, unrestricted double w, unrestricted double h);
  undefined roundRect(unrestricted double x, unrestricted double y, unrestricted double w, unrestricted double h, optional (unrestricted double or DOMPointInit or sequence<(unrestricted double or DOMPointInit)>) radii = 0);
  undefined arc(unrestricted double x, unrestricted double y, unrestricted double radius, unrestricted double startAngle, unrestricted double endAngle, optional boolean counterclockwise = false); 
  undefined ellipse(unrestricted double x, unrestricted double y, unrestricted double radiusX, unrestricted double radiusY, unrestricted double rotation, unrestricted double startAngle, unrestricted double endAngle, optional boolean counterclockwise = false); 
};

[Exposed=(Window,Worker)]
interface CanvasGradient {
  // opaque object
  undefined addColorStop(double offset, DOMString color);
};

[Exposed=(Window,Worker)]
interface CanvasPattern {
  // opaque object
  undefined setTransform(optional DOMMatrix2DInit transform = {});
};

[Exposed=(Window,Worker)]
interface TextMetrics {
  // x-direction
  readonly attribute double width; // advance width
  readonly attribute double actualBoundingBoxLeft;
  readonly attribute double actualBoundingBoxRight;

  // y-direction
  readonly attribute double fontBoundingBoxAscent;
  readonly attribute double fontBoundingBoxDescent;
  readonly attribute double actualBoundingBoxAscent;
  readonly attribute double actualBoundingBoxDescent;
  readonly attribute double emHeightAscent;
  readonly attribute double emHeightDescent;
  readonly attribute double hangingBaseline;
  readonly attribute double alphabeticBaseline;
  readonly attribute double ideographicBaseline;
};

[Exposed=(Window,Worker)]
interface Path2D {
  constructor(optional (Path2D or DOMString) path);

  undefined addPath(Path2D path, optional DOMMatrix2DInit transform = {});
};
Path2D includes CanvasPath;
```

To maintain compatibility with existing web content, user agents need to
enumerate methods defined in
[`CanvasUserInterface`{#2dcontext:canvasuserinterface-2}](#canvasuserinterface)
immediately after the
[`stroke()`{#2dcontext:dom-context-2d-stroke-2}](#dom-context-2d-stroke)
method on
[`CanvasRenderingContext2D`{#2dcontext:canvasrenderingcontext2d-18}](#canvasrenderingcontext2d)
objects.

`context`{.variable}` = ``canvas`{.variable}`.`[`getContext`](#dom-canvas-getcontext){#2dcontext:dom-canvas-getcontext}`('2d' [, { [ `[`alpha`](#dom-canvasrenderingcontext2dsettings-alpha){#2dcontext:dom-canvasrenderingcontext2dsettings-alpha-2}`: true ] [, `[`desynchronized`](#dom-canvasrenderingcontext2dsettings-desynchronized){#2dcontext:dom-canvasrenderingcontext2dsettings-desynchronized-2}`: false ] [, `[`colorSpace`](#dom-canvasrenderingcontext2dsettings-colorspace){#2dcontext:dom-canvasrenderingcontext2dsettings-colorspace-2}`: 'srgb'] [, `[`willReadFrequently`](#dom-canvasrenderingcontext2dsettings-willreadfrequently){#2dcontext:dom-canvasrenderingcontext2dsettings-willreadfrequently-2}`: false ]} ])`

Returns a
[`CanvasRenderingContext2D`{#2dcontext:canvasrenderingcontext2d-19}](#canvasrenderingcontext2d)
object that is permanently bound to a particular
[`canvas`{#2dcontext:the-canvas-element}](#the-canvas-element) element.

If the
[`alpha`{#2dcontext:dom-canvasrenderingcontext2dsettings-alpha-3}](#dom-canvasrenderingcontext2dsettings-alpha)
member is false, then the context is forced to always be opaque.

If the
[`desynchronized`{#2dcontext:dom-canvasrenderingcontext2dsettings-desynchronized-3}](#dom-canvasrenderingcontext2dsettings-desynchronized)
member is true, then the context might be
[desynchronized](#concept-canvas-desynchronized){#2dcontext:concept-canvas-desynchronized}.

The
[`colorSpace`{#2dcontext:dom-canvasrenderingcontext2dsettings-colorspace-3}](#dom-canvasrenderingcontext2dsettings-colorspace)
member specifies the [color
space](#concept-canvas-color-space){#2dcontext:concept-canvas-color-space}
of the rendering context.

The
[`colorType`{#2dcontext:dom-canvasrenderingcontext2dsettings-colortype-2}](#dom-canvasrenderingcontext2dsettings-colortype)
member specifies the [color
type](#concept-canvas-color-type){#2dcontext:concept-canvas-color-type}
of the rendering context.

If the
[`willReadFrequently`{#2dcontext:dom-canvasrenderingcontext2dsettings-willreadfrequently-3}](#dom-canvasrenderingcontext2dsettings-willreadfrequently)
member is true, then the context is marked for [readback
optimization](#concept-canvas-will-read-frequently){#2dcontext:concept-canvas-will-read-frequently}.

`context`{.variable}`.`[`canvas`](#dom-context-2d-canvas){#dom-context-2d-canvas-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/canvas](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/canvas "The CanvasRenderingContext2D.canvas property, part of the Canvas API, is a read-only reference to the HTMLCanvasElement object that is associated with a given context. It might be null if there is no associated <canvas> element.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns the
[`canvas`{#2dcontext:the-canvas-element-2}](#the-canvas-element)
element.

`attributes`{.variable}` = ``context`{.variable}`.`[`getContextAttributes`](#dom-context-2d-canvas-getcontextattributes){#2dcontext:dom-context-2d-canvas-getcontextattributes-2}`()`

Returns an object whose:

- [`alpha`{#2dcontext:concept-canvas-alpha}](#concept-canvas-alpha)
  member is true if the context has an alpha component that is not 1.0;
  otherwise false.
- [`desynchronized`{#2dcontext:dom-canvasrenderingcontext2dsettings-desynchronized-4}](#dom-canvasrenderingcontext2dsettings-desynchronized)
  member is true if the context can be
  [desynchronized](#concept-canvas-desynchronized){#2dcontext:concept-canvas-desynchronized-2}.
- [`colorSpace`{#2dcontext:dom-canvasrenderingcontext2dsettings-colorspace-4}](#dom-canvasrenderingcontext2dsettings-colorspace)
  member is a string indicating the context\'s [color
  space](#concept-canvas-color-space){#2dcontext:concept-canvas-color-space-2}.
- [`colorType`{#2dcontext:dom-canvasrenderingcontext2dsettings-colortype-3}](#dom-canvasrenderingcontext2dsettings-colortype)
  member is a string indicating the context\'s [color
  type](#concept-canvas-color-type){#2dcontext:concept-canvas-color-type-2}.
- [`willReadFrequently`{#2dcontext:dom-canvasrenderingcontext2dsettings-willreadfrequently-4}](#dom-canvasrenderingcontext2dsettings-willreadfrequently)
  member is true if the context is marked for [readback
  optimization](#concept-canvas-will-read-frequently){#2dcontext:concept-canvas-will-read-frequently-2}.

------------------------------------------------------------------------

The
[`CanvasRenderingContext2D`{#2dcontext:canvasrenderingcontext2d-20}](#canvasrenderingcontext2d)
2D rendering context represents a flat linear Cartesian surface whose
origin (0,0) is at the top left corner, with the coordinate space having
`x`{.variable} values increasing when going right, and `y`{.variable}
values increasing when going down. The `x`{.variable}-coordinate of the
right-most edge is equal to the width of the rendering context\'s
[output bitmap](#output-bitmap){#2dcontext:output-bitmap} in [CSS
pixels](https://drafts.csswg.org/css-values/#px){#2dcontext:'px'
x-internal="'px'"}; similarly, the `y`{.variable}-coordinate of the
bottom-most edge is equal to the height of the rendering context\'s
[output bitmap](#output-bitmap){#2dcontext:output-bitmap-2} in [CSS
pixels](https://drafts.csswg.org/css-values/#px){#2dcontext:'px'-2
x-internal="'px'"}.

The size of the coordinate space does not necessarily represent the size
of the actual bitmaps that the user agent will use internally or during
rendering. On high-definition displays, for instance, the user agent may
internally use bitmaps with four device pixels per unit in the
coordinate space, so that the rendering remains at high quality
throughout. Anti-aliasing can similarly be implemented using
oversampling with bitmaps of a higher resolution than the final image on
the display.

::: example
Using [CSS
pixels](https://drafts.csswg.org/css-values/#px){#2dcontext:'px'-3
x-internal="'px'"} to describe the size of a rendering context\'s
[output bitmap](#output-bitmap){#2dcontext:output-bitmap-3} does not
mean that when rendered the canvas will cover an equivalent area in [CSS
pixels](https://drafts.csswg.org/css-values/#px){#2dcontext:'px'-4
x-internal="'px'"}. [CSS
pixels](https://drafts.csswg.org/css-values/#px){#2dcontext:'px'-5
x-internal="'px'"} are reused for ease of integration with CSS features,
such as text layout.

In other words, the
[`canvas`{#2dcontext:the-canvas-element-3}](#the-canvas-element) element
below\'s rendering context has a 200x200 [output
bitmap](#output-bitmap){#2dcontext:output-bitmap-4} (which internally
uses [CSS
pixels](https://drafts.csswg.org/css-values/#px){#2dcontext:'px'-6
x-internal="'px'"} as a unit for ease of integration with CSS) and is
rendered as 100x100 [CSS
pixels](https://drafts.csswg.org/css-values/#px){#2dcontext:'px'-7
x-internal="'px'"}:

``` html
<canvas width=200 height=200 style=width:100px;height:100px>
```
:::

------------------------------------------------------------------------

The [2D context creation algorithm]{#2d-context-creation-algorithm
.dfn}, which is passed a `target`{.variable} (a
[`canvas`{#2dcontext:the-canvas-element-4}](#the-canvas-element)
element) and `options`{.variable}, consists of running these steps:

1.  Let `settings`{.variable} be the result of
    [converting](https://webidl.spec.whatwg.org/#es-type-mapping){#2dcontext:concept-idl-convert
    x-internal="concept-idl-convert"} `options`{.variable} to the
    dictionary type
    [`CanvasRenderingContext2DSettings`{#2dcontext:canvasrenderingcontext2dsettings-2}](#canvasrenderingcontext2dsettings).
    (This can throw an exception.)

2.  Let `context`{.variable} be a new
    [`CanvasRenderingContext2D`{#2dcontext:canvasrenderingcontext2d-21}](#canvasrenderingcontext2d)
    object.

3.  Initialize `context`{.variable}\'s
    [`canvas`{#2dcontext:dom-context-2d-canvas-2}](#dom-context-2d-canvas)
    attribute to point to `target`{.variable}.

4.  Set `context`{.variable}\'s [output
    bitmap](#output-bitmap){#2dcontext:output-bitmap-5} to the same
    bitmap as `target`{.variable}\'s bitmap (so that they are shared).

5.  [Set bitmap
    dimensions](#concept-canvas-set-bitmap-dimensions){#2dcontext:concept-canvas-set-bitmap-dimensions}
    to [the numeric
    values](#obtain-numeric-values){#2dcontext:obtain-numeric-values} of
    `target`{.variable}\'s
    [`width`{#2dcontext:attr-canvas-width}](#attr-canvas-width) and
    [`height`{#2dcontext:attr-canvas-height}](#attr-canvas-height)
    content attributes.

6.  Run the [canvas settings output bitmap initialization
    algorithm](#canvas-setting-init-bitmap){#2dcontext:canvas-setting-init-bitmap},
    given `context`{.variable} and `settings`{.variable}.

7.  Return `context`{.variable}.

------------------------------------------------------------------------

When the user agent is to [set bitmap
dimensions]{#concept-canvas-set-bitmap-dimensions .dfn} to
`width`{.variable} and `height`{.variable}, it must run these steps:

1.  [Reset the rendering context to its default
    state](#reset-the-rendering-context-to-its-default-state){#2dcontext:reset-the-rendering-context-to-its-default-state-2}.

2.  Resize the [output
    bitmap](#output-bitmap){#2dcontext:output-bitmap-6} to the new
    `width`{.variable} and `height`{.variable}.

3.  Let `canvas`{.variable} be the
    [`canvas`{#2dcontext:the-canvas-element-5}](#the-canvas-element)
    element to which the rendering context\'s
    [`canvas`{#2dcontext:dom-context-2d-canvas-3}](#dom-context-2d-canvas)
    attribute was initialized.

4.  If [the numeric
    value](#obtain-numeric-values){#2dcontext:obtain-numeric-values-2}
    of `canvas`{.variable}\'s
    [`width`{#2dcontext:attr-canvas-width-2}](#attr-canvas-width)
    content attribute differs from `width`{.variable}, then set
    `canvas`{.variable}\'s
    [`width`{#2dcontext:attr-canvas-width-3}](#attr-canvas-width)
    content attribute to the shortest possible string representing
    `width`{.variable} as a [valid non-negative
    integer](common-microsyntaxes.html#valid-non-negative-integer){#2dcontext:valid-non-negative-integer}.

5.  If [the numeric
    value](#obtain-numeric-values){#2dcontext:obtain-numeric-values-3}
    of `canvas`{.variable}\'s
    [`height`{#2dcontext:attr-canvas-height-2}](#attr-canvas-height)
    content attribute differs from `height`{.variable}, then set
    `canvas`{.variable}\'s
    [`height`{#2dcontext:attr-canvas-height-3}](#attr-canvas-height)
    content attribute to the shortest possible string representing
    `height`{.variable} as a [valid non-negative
    integer](common-microsyntaxes.html#valid-non-negative-integer){#2dcontext:valid-non-negative-integer-2}.

::: example
Only one square appears to be drawn in the following example:

``` js
// canvas is a reference to a <canvas> element
var context = canvas.getContext('2d');
context.fillRect(0,0,50,50);
canvas.setAttribute('width', '300'); // clears the canvas
context.fillRect(0,100,50,50);
canvas.width = canvas.width; // clears the canvas
context.fillRect(100,0,50,50); // only this square remains
```
:::

------------------------------------------------------------------------

The [`canvas`]{#dom-context-2d-canvas .dfn
dfn-for="CanvasRenderingContext2D" dfn-type="attribute"} attribute must
return the value it was initialized to when the object was created.

------------------------------------------------------------------------

The
[`PredefinedColorSpace`{#2dcontext:predefinedcolorspace-2}](#predefinedcolorspace)
enumeration is used to specify the [color
space](#concept-canvas-color-space){#2dcontext:concept-canvas-color-space-3}
of the canvas\'s backing store.

The \"[`srgb`]{#dom-predefinedcolorspace-srgb .dfn
dfn-for="PredefinedColorSpace" dfn-type="enum-value"}\" value indicates
the
[\'srgb\'](https://drafts.csswg.org/css-color/#valdef-color-srgb){#2dcontext:'srgb'
x-internal="'srgb'"} color space.

The \"[`display-p3`]{#dom-predefinedcolorspace-display-p3 .dfn
dfn-for="PredefinedColorSpace" dfn-type="enum-value"}\" value indicates
the
[\'display-p3\'](https://drafts.csswg.org/css-color/#valdef-color-display-p3){#2dcontext:'display-p3'
x-internal="'display-p3'"} color space.

The algorithm for converting between color spaces can be found in the
[Converting
Colors](https://drafts.csswg.org/css-color/#color-conversion){#2dcontext:converting-colors
x-internal="converting-colors"} section of CSS Color.
[\[CSSCOLOR\]](references.html#refsCSSCOLOR)

------------------------------------------------------------------------

The [`CanvasColorType`{#2dcontext:canvascolortype-2}](#canvascolortype)
enumeration is used to specify the [color
type](#concept-canvas-color-type){#2dcontext:concept-canvas-color-type-3}
of the canvas\'s backing store.

The \"[`unorm8`]{#dom-canvascolortype-unorm8 .dfn
dfn-for="CanvasColorType" dfn-type="enum-value"}\" value indicates that
the type for all color components is 8-bit unsigned normalized.

The \"[`float16`]{#dom-canvascolortype-float16 .dfn
dfn-for="CanvasColorType" dfn-type="enum-value"}\" value indicates that
the type for all color components is 16-bit floating point.

------------------------------------------------------------------------

The [`CanvasFillRule`{#2dcontext:canvasfillrule-7}](#canvasfillrule)
enumeration is used to select the [fill rule]{#fill-rule .dfn} algorithm
by which to determine if a point is inside or outside a path.

The \"[`nonzero`]{#dom-context-2d-fillrule-nonzero .dfn
dfn-for="CanvasFillRule" dfn-type="enum-value"}\" value indicates the
nonzero winding rule, wherein a point is considered to be outside a
shape if the number of times a half-infinite straight line drawn from
that point crosses the shape\'s path going in one direction is equal to
the number of times it crosses the path going in the other direction.

The \"[`evenodd`]{#dom-context-2d-fillrule-evenodd .dfn
dfn-for="CanvasFillRule" dfn-type="enum-value"}\" value indicates the
even-odd rule, wherein a point is considered to be outside a shape if
the number of times a half-infinite straight line drawn from that point
crosses the shape\'s path is even.

If a point is not outside a shape, it is inside the shape.

------------------------------------------------------------------------

The
[`ImageSmoothingQuality`{#2dcontext:imagesmoothingquality-2}](#imagesmoothingquality)
enumeration is used to express a preference for the interpolation
quality to use when smoothing images.

The \"[`low`]{#dom-context-2d-imagesmoothingquality-low .dfn
dfn-for="ImageSmoothingQuality" dfn-type="enum-value"}\" value indicates
a preference for a low level of image interpolation quality. Low-quality
image interpolation may be more computationally efficient than higher
settings.

The \"[`medium`]{#dom-context-2d-imagesmoothingquality-medium .dfn
dfn-for="ImageSmoothingQuality" dfn-type="enum-value"}\" value indicates
a preference for a medium level of image interpolation quality.

The \"[`high`]{#dom-context-2d-imagesmoothingquality-high .dfn
dfn-for="ImageSmoothingQuality" dfn-type="enum-value"}\" value indicates
a preference for a high level of image interpolation quality.
High-quality image interpolation may be more computationally expensive
than lower settings.

Bilinear scaling is an example of a relatively fast, lower-quality
image-smoothing algorithm. Bicubic or Lanczos scaling are examples of
image-smoothing algorithms that produce higher-quality output. This
specification does not mandate that specific interpolation algorithms be
used.

###### [4.12.5.1.1]{.secno} Implementation notes[](#implementation-notes){.self-link}

*This section is non-normative.*

The [output
bitmap](#output-bitmap){#implementation-notes:output-bitmap}, when it is
not directly displayed by the user agent, implementations can, instead
of updating this bitmap, merely remember the sequence of drawing
operations that have been applied to it until such time as the bitmap\'s
actual data is needed (for example because of a call to
[`drawImage()`{#implementation-notes:dom-context-2d-drawimage}](#dom-context-2d-drawimage),
or the
[`createImageBitmap()`{#implementation-notes:dom-createimagebitmap}](imagebitmap-and-animations.html#dom-createimagebitmap)
factory method). In many cases, this will be more memory efficient.

The bitmap of a
[`canvas`{#implementation-notes:the-canvas-element}](#the-canvas-element)
element is the one bitmap that\'s pretty much always going to be needed
in practice. The [output
bitmap](#output-bitmap){#implementation-notes:output-bitmap-2} of a
rendering context, when it has one, is always just an alias to a
[`canvas`{#implementation-notes:the-canvas-element-2}](#the-canvas-element)
element\'s bitmap.

Additional bitmaps are sometimes needed, e.g. to enable fast drawing
when the canvas is being painted at a different size than its [natural
size](https://drafts.csswg.org/css-images/#natural-dimensions){#implementation-notes:natural-dimensions
x-internal="natural-dimensions"}, or to enable double buffering so that
graphics updates, like page scrolling for example, can be processed
concurrently while canvas draw commands are being executed.

###### [4.12.5.1.2]{.secno} The canvas settings[](#the-canvas-settings){.self-link}

A
[`CanvasSettings`{#the-canvas-settings:canvassettings}](#canvassettings)
object has an [output bitmap]{#output-bitmap .dfn} that is initialized
when the object is created.

The [output bitmap](#output-bitmap){#the-canvas-settings:output-bitmap}
has an
[origin-clean](#concept-canvas-origin-clean){#the-canvas-settings:concept-canvas-origin-clean}
flag, which can be set to true or false. Initially, when one of these
bitmaps is created, its
[origin-clean](#concept-canvas-origin-clean){#the-canvas-settings:concept-canvas-origin-clean-2}
flag must be set to true.

The
[`CanvasSettings`{#the-canvas-settings:canvassettings-2}](#canvassettings)
object also has an [alpha]{#concept-canvas-alpha .dfn} boolean. When a
[`CanvasSettings`{#the-canvas-settings:canvassettings-3}](#canvassettings)
object\'s
[alpha](#concept-canvas-alpha){#the-canvas-settings:concept-canvas-alpha}
is false, then its alpha component must be fixed to 1.0 (fully opaque)
for all pixels, and attempts to change the alpha component of any pixel
must be silently ignored.

Thus, the bitmap of such a context starts off as [opaque
black](https://drafts.csswg.org/css-color/#opaque-black){#the-canvas-settings:opaque-black
x-internal="opaque-black"} instead of [transparent
black](https://drafts.csswg.org/css-color/#transparent-black){#the-canvas-settings:transparent-black
x-internal="transparent-black"};
[`clearRect()`{#the-canvas-settings:dom-context-2d-clearrect}](#dom-context-2d-clearrect)
always results in [opaque
black](https://drafts.csswg.org/css-color/#opaque-black){#the-canvas-settings:opaque-black-2
x-internal="opaque-black"} pixels, every fourth byte from
[`getImageData()`{#the-canvas-settings:dom-context-2d-getimagedata}](#dom-context-2d-getimagedata)
is always 255, the
[`putImageData()`{#the-canvas-settings:dom-context-2d-putimagedata}](#dom-context-2d-putimagedata)
method effectively ignores every fourth byte in its input, and so on.
However, the alpha component of styles and images drawn onto the canvas
are still honoured up to the point where they would impact the [output
bitmap](#output-bitmap){#the-canvas-settings:output-bitmap-2}\'s alpha
component; for instance, drawing a 50% transparent white square on a
freshly created [output
bitmap](#output-bitmap){#the-canvas-settings:output-bitmap-3} with its
[alpha](#concept-canvas-alpha){#the-canvas-settings:concept-canvas-alpha-2}
set to false will result in a fully-opaque gray square.

The
[`CanvasSettings`{#the-canvas-settings:canvassettings-4}](#canvassettings)
object also has a [desynchronized]{#concept-canvas-desynchronized .dfn}
boolean. When a
[`CanvasSettings`{#the-canvas-settings:canvassettings-5}](#canvassettings)
object\'s
[desynchronized](#concept-canvas-desynchronized){#the-canvas-settings:concept-canvas-desynchronized}
is true, then the user agent may optimize the rendering of the canvas to
reduce the latency, as measured from input events to rasterization, by
desynchronizing the canvas paint cycle from the event loop, bypassing
the ordinary user agent rendering algorithm, or both. Insofar as this
mode involves bypassing the usual paint mechanisms, rasterization, or
both, it might introduce visible tearing artifacts.

The user agent usually renders on a buffer which is not being displayed,
quickly swapping it and the one being scanned out for presentation; the
former buffer is called back buffer and the latter *front buffer*. A
popular technique for reducing latency is called front buffer rendering,
also known as *single buffer* rendering, where rendering happens in
parallel and racily with the scanning out process. This technique
reduces the latency at the price of potentially introducing tearing
artifacts and can be used to implement in total or part of the
[desynchronized](#concept-canvas-desynchronized){#the-canvas-settings:concept-canvas-desynchronized-2}
boolean. [\[MULTIPLEBUFFERING\]](references.html#refsMULTIPLEBUFFERING)

The
[desynchronized](#concept-canvas-desynchronized){#the-canvas-settings:concept-canvas-desynchronized-3}
boolean can be useful when implementing certain kinds of applications,
such as drawing applications, where the latency between input and
rasterization is critical.

The
[`CanvasSettings`{#the-canvas-settings:canvassettings-6}](#canvassettings)
object also has a [will read
frequently]{#concept-canvas-will-read-frequently .dfn} boolean. When a
[`CanvasSettings`{#the-canvas-settings:canvassettings-7}](#canvassettings)
object\'s [will read
frequently](#concept-canvas-will-read-frequently){#the-canvas-settings:concept-canvas-will-read-frequently}
is true, the user agent may optimize the canvas for readback operations.

On most devices the user agent needs to decide whether to store the
canvas\'s [output
bitmap](#output-bitmap){#the-canvas-settings:output-bitmap-4} on the GPU
(this is also called \"hardware accelerated\"), or on the CPU (also
called \"software\"). Most rendering operations are more performant for
accelerated canvases, with the major exception being readback with
[`getImageData()`{#the-canvas-settings:dom-context-2d-getimagedata-2}](#dom-context-2d-getimagedata),
[`toDataURL()`{#the-canvas-settings:dom-canvas-todataurl}](#dom-canvas-todataurl),
or
[`toBlob()`{#the-canvas-settings:dom-canvas-toblob}](#dom-canvas-toblob).
[`CanvasSettings`{#the-canvas-settings:canvassettings-8}](#canvassettings)
objects with [will read
frequently](#concept-canvas-will-read-frequently){#the-canvas-settings:concept-canvas-will-read-frequently-2}
equal to true tell the user agent that the webpage is likely to perform
many readback operations and that it is advantageous to use a software
canvas.

The
[`CanvasSettings`{#the-canvas-settings:canvassettings-9}](#canvassettings)
object also has a [color space]{#concept-canvas-color-space .dfn}
setting of type
[`PredefinedColorSpace`{#the-canvas-settings:predefinedcolorspace}](#predefinedcolorspace).
The
[`CanvasSettings`{#the-canvas-settings:canvassettings-10}](#canvassettings)
object\'s [color
space](#concept-canvas-color-space){#the-canvas-settings:concept-canvas-color-space}
indicates the color space for the [output
bitmap](#output-bitmap){#the-canvas-settings:output-bitmap-5}.

The
[`CanvasSettings`{#the-canvas-settings:canvassettings-11}](#canvassettings)
object also has a [color type]{#concept-canvas-color-type .dfn} setting
of type
[`CanvasColorType`{#the-canvas-settings:canvascolortype}](#canvascolortype).
The
[`CanvasSettings`{#the-canvas-settings:canvassettings-12}](#canvassettings)
object\'s [color
type](#concept-canvas-color-type){#the-canvas-settings:concept-canvas-color-type}
indicates the data type of the color and alpha components of the pixels
of the [output
bitmap](#output-bitmap){#the-canvas-settings:output-bitmap-6}.

To [initialize a `CanvasSettings` output
bitmap]{#canvas-setting-init-bitmap .dfn}, given a
[`CanvasSettings`{#the-canvas-settings:canvassettings-13}](#canvassettings)
`context`{.variable} and a
[`CanvasRenderingContext2DSettings`{#the-canvas-settings:canvasrenderingcontext2dsettings}](#canvasrenderingcontext2dsettings)
`settings`{.variable}:

1.  Set `context`{.variable}\'s
    [alpha](#concept-canvas-alpha){#the-canvas-settings:concept-canvas-alpha-3}
    to
    `settings`{.variable}\[\"[`alpha`]{#dom-canvasrenderingcontext2dsettings-alpha
    .dfn dfn-for="CanvasRenderingContext2DSettings"
    dfn-type="dict-member"}\"\].

2.  Set `context`{.variable}\'s
    [desynchronized](#concept-canvas-desynchronized){#the-canvas-settings:concept-canvas-desynchronized-4}
    to
    `settings`{.variable}\[\"[`desynchronized`]{#dom-canvasrenderingcontext2dsettings-desynchronized
    .dfn dfn-for="CanvasRenderingContext2DSettings"
    dfn-type="dict-member"}\"\].

3.  Set `context`{.variable}\'s [color
    space](#concept-canvas-color-space){#the-canvas-settings:concept-canvas-color-space-2}
    to
    `settings`{.variable}\[\"[`colorSpace`]{#dom-canvasrenderingcontext2dsettings-colorspace
    .dfn dfn-for="CanvasRenderingContext2DSettings"
    dfn-type="dict-member"}\"\].

4.  Set `context`{.variable}\'s [color
    type](#concept-canvas-color-type){#the-canvas-settings:concept-canvas-color-type-2}
    to
    `settings`{.variable}\[\"[`colorType`]{#dom-canvasrenderingcontext2dsettings-colortype
    .dfn dfn-for="CanvasRenderingContext2DSettings"
    dfn-type="dict-member"}\"\].

5.  Set `context`{.variable}\'s [will read
    frequently](#concept-canvas-will-read-frequently){#the-canvas-settings:concept-canvas-will-read-frequently-3}
    to
    `settings`{.variable}\[\"[`willReadFrequently`]{#dom-canvasrenderingcontext2dsettings-willreadfrequently
    .dfn dfn-for="CanvasRenderingContext2DSettings"
    dfn-type="dict-member"}\"\].

The
[`getContextAttributes()`]{#dom-context-2d-canvas-getcontextattributes
.dfn dfn-for="CanvasSettings" dfn-type="method"} method steps are to
return «\[
\"[`alpha`{#the-canvas-settings:dom-canvasrenderingcontext2dsettings-alpha}](#dom-canvasrenderingcontext2dsettings-alpha)\"
→ [this](https://webidl.spec.whatwg.org/#this){#the-canvas-settings:this
x-internal="this"}\'s
[alpha](#concept-canvas-alpha){#the-canvas-settings:concept-canvas-alpha-4},
\"[`desynchronized`{#the-canvas-settings:dom-canvasrenderingcontext2dsettings-desynchronized}](#dom-canvasrenderingcontext2dsettings-desynchronized)\"
→
[this](https://webidl.spec.whatwg.org/#this){#the-canvas-settings:this-2
x-internal="this"}\'s
[desynchronized](#concept-canvas-desynchronized){#the-canvas-settings:concept-canvas-desynchronized-5},
\"[`colorSpace`{#the-canvas-settings:dom-canvasrenderingcontext2dsettings-colorspace}](#dom-canvasrenderingcontext2dsettings-colorspace)\"
→
[this](https://webidl.spec.whatwg.org/#this){#the-canvas-settings:this-3
x-internal="this"}\'s [color
space](#concept-canvas-color-space){#the-canvas-settings:concept-canvas-color-space-3},
\"[`colorType`{#the-canvas-settings:dom-canvasrenderingcontext2dsettings-colortype}](#dom-canvasrenderingcontext2dsettings-colortype)\"
→
[this](https://webidl.spec.whatwg.org/#this){#the-canvas-settings:this-4
x-internal="this"}\'s [color
type](#concept-canvas-color-type){#the-canvas-settings:concept-canvas-color-type-3},
\"[`willReadFrequently`{#the-canvas-settings:dom-canvasrenderingcontext2dsettings-willreadfrequently}](#dom-canvasrenderingcontext2dsettings-willreadfrequently)\"
→
[this](https://webidl.spec.whatwg.org/#this){#the-canvas-settings:this-5
x-internal="this"}\'s [will read
frequently](#concept-canvas-will-read-frequently){#the-canvas-settings:concept-canvas-will-read-frequently-4}
\]».

###### [4.12.5.1.3]{.secno} The canvas state[](#the-canvas-state){.self-link}

Objects that implement the
[`CanvasState`{#the-canvas-state:canvasstate}](#canvasstate) interface
maintain a stack of drawing states. [Drawing states]{#drawing-state
.dfn} consist of:

- The current [transformation
  matrix](#transformations){#the-canvas-state:transformations}.

- The current [clipping
  region](#clipping-region){#the-canvas-state:clipping-region}.

- The current [letter
  spacing](#concept-canvastextdrawingstyles-letter-spacing){#the-canvas-state:concept-canvastextdrawingstyles-letter-spacing},
  [word
  spacing](#concept-canvastextdrawingstyles-word-spacing){#the-canvas-state:concept-canvastextdrawingstyles-word-spacing},
  [fill
  style](#concept-canvasfillstrokestyles-fill-style){#the-canvas-state:concept-canvasfillstrokestyles-fill-style},
  [stroke
  style](#concept-canvasfillstrokestyles-stroke-style){#the-canvas-state:concept-canvasfillstrokestyles-stroke-style},
  [filter](#concept-canvas-current-filter){#the-canvas-state:concept-canvas-current-filter},
  [global
  alpha](#concept-canvas-global-alpha){#the-canvas-state:concept-canvas-global-alpha},
  [compositing and blending
  operator](#current-compositing-and-blending-operator){#the-canvas-state:current-compositing-and-blending-operator},
  and [shadow
  color](#concept-canvasshadowstyles-shadow-color){#the-canvas-state:concept-canvasshadowstyles-shadow-color}.

- The current values of the following attributes:
  [`lineWidth`{#the-canvas-state:dom-context-2d-linewidth}](#dom-context-2d-linewidth),
  [`lineCap`{#the-canvas-state:dom-context-2d-linecap}](#dom-context-2d-linecap),
  [`lineJoin`{#the-canvas-state:dom-context-2d-linejoin}](#dom-context-2d-linejoin),
  [`miterLimit`{#the-canvas-state:dom-context-2d-miterlimit}](#dom-context-2d-miterlimit),
  [`lineDashOffset`{#the-canvas-state:dom-context-2d-linedashoffset}](#dom-context-2d-linedashoffset),
  [`shadowOffsetX`{#the-canvas-state:dom-context-2d-shadowoffsetx}](#dom-context-2d-shadowoffsetx),
  [`shadowOffsetY`{#the-canvas-state:dom-context-2d-shadowoffsety}](#dom-context-2d-shadowoffsety),
  [`shadowBlur`{#the-canvas-state:dom-context-2d-shadowblur}](#dom-context-2d-shadowblur),
  [`lang`{#the-canvas-state:dom-context-2d-lang}](#dom-context-2d-lang),
  [`font`{#the-canvas-state:dom-context-2d-font}](#dom-context-2d-font),
  [`textAlign`{#the-canvas-state:dom-context-2d-textalign}](#dom-context-2d-textalign),
  [`textBaseline`{#the-canvas-state:dom-context-2d-textbaseline}](#dom-context-2d-textbaseline),
  [`direction`{#the-canvas-state:dom-context-2d-direction}](#dom-context-2d-direction),
  [`fontKerning`{#the-canvas-state:dom-context-2d-fontkerning}](#dom-context-2d-fontkerning),
  [`fontStretch`{#the-canvas-state:dom-context-2d-fontstretch}](#dom-context-2d-fontstretch),
  [`fontVariantCaps`{#the-canvas-state:dom-context-2d-fontvariantcaps}](#dom-context-2d-fontvariantcaps),
  [`textRendering`{#the-canvas-state:dom-context-2d-textrendering}](#dom-context-2d-textrendering),
  [`imageSmoothingEnabled`{#the-canvas-state:dom-context-2d-imagesmoothingenabled}](#dom-context-2d-imagesmoothingenabled),
  [`imageSmoothingQuality`{#the-canvas-state:dom-context-2d-imagesmoothingquality}](#dom-context-2d-imagesmoothingquality).

- The current [dash list](#dash-list){#the-canvas-state:dash-list}.

The rendering context\'s bitmaps are not part of the drawing state, as
they depend on whether and how the rendering context is bound to a
[`canvas`{#the-canvas-state:the-canvas-element}](#the-canvas-element)
element.

Objects that implement the
[`CanvasState`{#the-canvas-state:canvasstate-2}](#canvasstate) mixin
have a [context lost]{#concept-canvas-context-lost .dfn} boolean, that
is initialized to false when the object is created. The [context
lost](#concept-canvas-context-lost){#the-canvas-state:concept-canvas-context-lost}
value is updated in the [context lost
steps](webappapis.html#context-lost-steps){#the-canvas-state:context-lost-steps}.

`context`{.variable}`.`[`save`](#dom-context-2d-save){#dom-context-2d-save-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/save](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/save "The CanvasRenderingContext2D.save() method of the Canvas 2D API saves the entire state of the canvas by pushing the current state onto a stack.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari2+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Pushes the current state onto the stack.

`context`{.variable}`.`[`restore`](#dom-context-2d-restore){#dom-context-2d-restore-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/restore](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/restore "The CanvasRenderingContext2D.restore() method of the Canvas 2D API restores the most recently saved canvas state by popping the top entry in the drawing state stack. If there is no saved state, this method does nothing.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari2+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Pops the top state on the stack, restoring the context to that state.

`context`{.variable}`.`[`reset`](#dom-context-2d-reset){#dom-context-2d-reset-dev}`()`

Resets the rendering context, which includes the backing buffer, the
drawing state stack, path, and styles.

`context`{.variable}`.`[`isContextLost`](#dom-context-2d-iscontextlost){#dom-context-2d-iscontextlost-dev}`()`

Returns true if the rendering context was lost. Context loss can occur
due to driver crashes, running out of memory, etc. In these cases, the
canvas loses its backing storage and takes steps to [reset the rendering
context to its default
state](#reset-the-rendering-context-to-its-default-state){#the-canvas-state:reset-the-rendering-context-to-its-default-state}.

The [`save()`]{#dom-context-2d-save .dfn dfn-for="CanvasState"
dfn-type="method"} method steps are to push a copy of the current
drawing state onto the drawing state stack.

The [`restore()`]{#dom-context-2d-restore .dfn dfn-for="CanvasState"
dfn-type="method"} method steps are to pop the top entry in the drawing
state stack, and reset the drawing state it describes. If there is no
saved state, then the method must do nothing.

::::::: {.mdn-anno .wrapped .before}
MDN

:::: feature
[CanvasRenderingContext2D/reset](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/reset "The CanvasRenderingContext2D.reset() method of the Canvas 2D API resets the rendering context to its default state, allowing it to be reused for drawing something else without having to explicitly reset all the properties.")

::: support
[Firefox113+]{.firefox .yes}[SafariNo]{.safari .no}[Chrome99+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge99+]{.edge_blink .yes}

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
[OffscreenCanvasRenderingContext2D#canvasrenderingcontext2d.reset](https://developer.mozilla.org/en-US/docs/Web/API/OffscreenCanvasRenderingContext2D#canvasrenderingcontext2d.reset "The OffscreenCanvasRenderingContext2D interface is a CanvasRenderingContext2D rendering context for drawing to the bitmap of an OffscreenCanvas object. It is similar to the CanvasRenderingContext2D object, with the following differences:")

::: support
[Firefox113+]{.firefox .yes}[SafariNo]{.safari .no}[Chrome99+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge99+]{.edge_blink .yes}

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

The [`reset()`]{#dom-context-2d-reset .dfn dfn-for="CanvasState"
dfn-type="method"} method steps are to [reset the rendering context to
its default
state](#reset-the-rendering-context-to-its-default-state){#the-canvas-state:reset-the-rendering-context-to-its-default-state-2}.

To [reset the rendering context to its default
state]{#reset-the-rendering-context-to-its-default-state .dfn}:

1.  Clear canvas\'s bitmap to [transparent
    black](https://drafts.csswg.org/css-color/#transparent-black){#the-canvas-state:transparent-black
    x-internal="transparent-black"}.

2.  Empty the list of subpaths in context\'s [current default
    path](#current-default-path){#the-canvas-state:current-default-path}.

3.  Clear the context\'s drawing state stack.

4.  Reset everything that [drawing
    state](#drawing-state){#the-canvas-state:drawing-state} consists of
    to their initial values.

::::: {.mdn-anno .wrapped .before}
**⚠**MDN

:::: feature
[CanvasRenderingContext2D/isContextLost](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/isContextLost "The CanvasRenderingContext2D.isContextLost() method of the Canvas 2D API returns true if the rendering context is lost (and has not yet been reset). This might occur due to driver crashes, running out of memory, and so on.")

Support in one engine only.

::: support
[FirefoxNo]{.firefox .no}[SafariNo]{.safari .no}[Chrome99+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge99+]{.edge_blink .yes}

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

The [`isContextLost()`]{#dom-context-2d-iscontextlost .dfn
dfn-for="CanvasState" dfn-type="method"} method steps are to return
[this](https://webidl.spec.whatwg.org/#this){#the-canvas-state:this
x-internal="this"}\'s [context
lost](#concept-canvas-context-lost){#the-canvas-state:concept-canvas-context-lost-2}.

###### [4.12.5.1.4]{.secno} Line styles[](#line-styles){.self-link}

`context`{.variable}`.`[`lineWidth`](#dom-context-2d-linewidth){#dom-context-2d-linewidth-dev}` [ = ``value`{.variable}` ]`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/lineWidth](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/lineWidth "The CanvasRenderingContext2D.lineWidth property of the Canvas 2D API sets the thickness of lines.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari2+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

`styles`{.variable}`.`[`lineWidth`](#dom-context-2d-linewidth){#line-styles:dom-context-2d-linewidth}` [ = ``value`{.variable}` ]`

Returns the current line width.

Can be set, to change the line width. Values that are not finite values
greater than zero are ignored.

`context`{.variable}`.`[`lineCap`](#dom-context-2d-linecap){#dom-context-2d-linecap-dev}` [ = ``value`{.variable}` ]`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/lineCap](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/lineCap "The CanvasRenderingContext2D.lineCap property of the Canvas 2D API determines the shape used to draw the end points of lines.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari2+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

`styles`{.variable}`.`[`lineCap`](#dom-context-2d-linecap){#line-styles:dom-context-2d-linecap}` [ = ``value`{.variable}` ]`

Returns the current line cap style.

Can be set, to change the line cap style.

The possible line cap styles are \"`butt`\", \"`round`\", and
\"`square`\". Other values are ignored.

`context`{.variable}`.`[`lineJoin`](#dom-context-2d-linejoin){#dom-context-2d-linejoin-dev}` [ = ``value`{.variable}` ]`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/lineJoin](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/lineJoin "The CanvasRenderingContext2D.lineJoin property of the Canvas 2D API determines the shape used to join two line segments where they meet.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari2+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

`styles`{.variable}`.`[`lineJoin`](#dom-context-2d-linejoin){#line-styles:dom-context-2d-linejoin}` [ = ``value`{.variable}` ]`

Returns the current line join style.

Can be set, to change the line join style.

The possible line join styles are \"`bevel`\", \"`round`\", and
\"`miter`\". Other values are ignored.

`context`{.variable}`.`[`miterLimit`](#dom-context-2d-miterlimit){#dom-context-2d-miterlimit-dev}` [ = ``value`{.variable}` ]`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/miterLimit](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/miterLimit "The CanvasRenderingContext2D.miterLimit property of the Canvas 2D API sets the miter limit ratio.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari2+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

`styles`{.variable}`.`[`miterLimit`](#dom-context-2d-miterlimit){#line-styles:dom-context-2d-miterlimit}` [ = ``value`{.variable}` ]`

Returns the current miter limit ratio.

Can be set, to change the miter limit ratio. Values that are not finite
values greater than zero are ignored.

`context`{.variable}`.`[`setLineDash`](#dom-context-2d-setlinedash){#dom-context-2d-setlinedash-dev}`(``segments`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/setLineDash](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/setLineDash "The setLineDash() method of the Canvas 2D API's CanvasRenderingContext2D interface sets the line dash pattern used when stroking lines. It uses an array of values that specify alternating lengths of lines and gaps which describe the pattern.")

Support in all current engines.

::: support
[Firefox27+]{.firefox .yes}[Safari7+]{.safari .yes}[Chrome23+]{.chrome
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

`styles`{.variable}`.`[`setLineDash`](#dom-context-2d-setlinedash){#line-styles:dom-context-2d-setlinedash}`(``segments`{.variable}`)`

Sets the current line dash pattern (as used when stroking). The argument
is a list of distances for which to alternately have the line on and the
line off.

`segments`{.variable}` = ``context`{.variable}`.`[`getLineDash`](#dom-context-2d-getlinedash){#dom-context-2d-getlinedash-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/getLineDash](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/getLineDash "The getLineDash() method of the Canvas 2D API's CanvasRenderingContext2D interface gets the current line dash pattern.")

Support in all current engines.

::: support
[Firefox27+]{.firefox .yes}[Safari7+]{.safari .yes}[Chrome23+]{.chrome
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

`segments`{.variable}` = ``styles`{.variable}`.`[`getLineDash`](#dom-context-2d-getlinedash){#line-styles:dom-context-2d-getlinedash}`()`

Returns a copy of the current line dash pattern. The array returned will
always have an even number of entries (i.e. the pattern is normalized).

`context`{.variable}`.`[`lineDashOffset`](#dom-context-2d-linedashoffset){#dom-context-2d-linedashoffset-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/lineDashOffset](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/lineDashOffset "The CanvasRenderingContext2D.lineDashOffset property of the Canvas 2D API sets the line dash offset, or "phase."")

Support in all current engines.

::: support
[Firefox27+]{.firefox .yes}[Safari7+]{.safari .yes}[Chrome23+]{.chrome
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

`styles`{.variable}`.`[`lineDashOffset`](#dom-context-2d-linedashoffset){#line-styles:dom-context-2d-linedashoffset}

Returns the phase offset (in the same units as the line dash pattern).

Can be set, to change the phase offset. Values that are not finite
values are ignored.

Objects that implement the
[`CanvasPathDrawingStyles`{#line-styles:canvaspathdrawingstyles}](#canvaspathdrawingstyles)
interface have attributes and methods (defined in this section) that
control how lines are treated by the object.

The [`lineWidth`]{#dom-context-2d-linewidth .dfn
dfn-for="CanvasPathDrawingStyles" dfn-type="attribute"} attribute gives
the width of lines, in coordinate space units. On getting, it must
return the current value. On setting, zero, negative, infinite, and NaN
values must be ignored, leaving the value unchanged; other values must
change the current value to the new value.

When the object implementing the
[`CanvasPathDrawingStyles`{#line-styles:canvaspathdrawingstyles-2}](#canvaspathdrawingstyles)
interface is created, the
[`lineWidth`{#line-styles:dom-context-2d-linewidth-2}](#dom-context-2d-linewidth)
attribute must initially have the value 1.0.

------------------------------------------------------------------------

The [`lineCap`]{#dom-context-2d-linecap .dfn
dfn-for="CanvasPathDrawingStyles" dfn-type="attribute"} attribute
defines the type of endings that UAs will place on the end of lines. The
three valid values are \"`butt`\", \"`round`\", and \"`square`\".

On getting, it must return the current value. On setting, the current
value must be changed to the new value.

When the object implementing the
[`CanvasPathDrawingStyles`{#line-styles:canvaspathdrawingstyles-3}](#canvaspathdrawingstyles)
interface is created, the
[`lineCap`{#line-styles:dom-context-2d-linecap-2}](#dom-context-2d-linecap)
attribute must initially have the value \"`butt`\".

------------------------------------------------------------------------

The [`lineJoin`]{#dom-context-2d-linejoin .dfn
dfn-for="CanvasPathDrawingStyles" dfn-type="attribute"} attribute
defines the type of corners that UAs will place where two lines meet.
The three valid values are \"`bevel`\", \"`round`\", and \"`miter`\".

On getting, it must return the current value. On setting, the current
value must be changed to the new value.

When the object implementing the
[`CanvasPathDrawingStyles`{#line-styles:canvaspathdrawingstyles-4}](#canvaspathdrawingstyles)
interface is created, the
[`lineJoin`{#line-styles:dom-context-2d-linejoin-2}](#dom-context-2d-linejoin)
attribute must initially have the value \"`miter`\".

------------------------------------------------------------------------

When the
[`lineJoin`{#line-styles:dom-context-2d-linejoin-3}](#dom-context-2d-linejoin)
attribute has the value \"`miter`\", strokes use the miter limit ratio
to decide how to render joins. The miter limit ratio can be explicitly
set using the [`miterLimit`]{#dom-context-2d-miterlimit .dfn
dfn-for="CanvasPathDrawingStyles" dfn-type="attribute"} attribute. On
getting, it must return the current value. On setting, zero, negative,
infinite, and NaN values must be ignored, leaving the value unchanged;
other values must change the current value to the new value.

When the object implementing the
[`CanvasPathDrawingStyles`{#line-styles:canvaspathdrawingstyles-5}](#canvaspathdrawingstyles)
interface is created, the
[`miterLimit`{#line-styles:dom-context-2d-miterlimit-2}](#dom-context-2d-miterlimit)
attribute must initially have the value 10.0.

------------------------------------------------------------------------

Each
[`CanvasPathDrawingStyles`{#line-styles:canvaspathdrawingstyles-6}](#canvaspathdrawingstyles)
object has a [dash list]{#dash-list .dfn}, which is either empty or
consists of an even number of non-negative numbers. Initially, the [dash
list](#dash-list){#line-styles:dash-list} must be empty.

The [`setLineDash(``segments`{.variable}`)`]{#dom-context-2d-setlinedash
.dfn dfn-for="CanvasPathDrawingStyles" dfn-type="method"} method, when
invoked, must run these steps:

1.  If any value in `segments`{.variable} is not finite (e.g. an
    Infinity or a NaN value), or if any value is negative (less than
    zero), then return (without throwing an exception; user agents could
    show a message on a developer console, though, as that would be
    helpful for debugging).

2.  If the number of elements in `segments`{.variable} is odd, then let
    `segments`{.variable} be the concatenation of two copies of
    `segments`{.variable}.

3.  Set the object\'s [dash list](#dash-list){#line-styles:dash-list-2}
    to `segments`{.variable}.

When the [`getLineDash()`]{#dom-context-2d-getlinedash .dfn
dfn-for="CanvasPathDrawingStyles" dfn-type="method"} method is invoked,
it must return a sequence whose values are the values of the object\'s
[dash list](#dash-list){#line-styles:dash-list-3}, in the same order.

It is sometimes useful to change the \"phase\" of the dash pattern, e.g.
to achieve a \"marching ants\" effect. The phase can be set using the
[`lineDashOffset`]{#dom-context-2d-linedashoffset .dfn
dfn-for="CanvasPathDrawingStyles" dfn-type="attribute"} attribute. On
getting, it must return the current value. On setting, infinite and NaN
values must be ignored, leaving the value unchanged; other values must
change the current value to the new value.

When the object implementing the
[`CanvasPathDrawingStyles`{#line-styles:canvaspathdrawingstyles-7}](#canvaspathdrawingstyles)
interface is created, the
[`lineDashOffset`{#line-styles:dom-context-2d-linedashoffset-2}](#dom-context-2d-linedashoffset)
attribute must initially have the value 0.0.

------------------------------------------------------------------------

When a user agent is to [trace a path]{#trace-a-path .dfn}, given an
object `style`{.variable} that implements the
[`CanvasPathDrawingStyles`{#line-styles:canvaspathdrawingstyles-8}](#canvaspathdrawingstyles)
interface, it must run the following algorithm. This algorithm returns a
new [path](#concept-path){#line-styles:concept-path}.

1.  Let `path`{.variable} be a copy of the path being traced.

2.  Prune all zero-length [line
    segments](#line-segments){#line-styles:line-segments} from
    `path`{.variable}.

3.  Remove from `path`{.variable} any subpaths containing no lines (i.e.
    subpaths with just one point).

4.  Replace each point in each subpath of `path`{.variable} other than
    the first point and the last point of each subpath by a *join* that
    joins the line leading to that point to the line leading out of that
    point, such that the subpaths all consist of two points (a starting
    point with a line leading out of it, and an ending point with a line
    leading into it), one or more lines (connecting the points and the
    joins), and zero or more joins (each connecting one line to
    another), connected together such that each subpath is a series of
    one or more lines with a join between each one and a point on each
    end.

5.  Add a straight closing line to each closed subpath in
    `path`{.variable} connecting the last point and the first point of
    that subpath; change the last point to a join (from the previously
    last line to the newly added closing line), and change the first
    point to a join (from the newly added closing line to the first
    line).

6.  If `style`{.variable}\'s [dash
    list](#dash-list){#line-styles:dash-list-4} is empty, then jump to
    the step labeled *convert*.

7.  Let `pattern width`{.variable} be the concatenation of all the
    entries of `style`{.variable}\'s [dash
    list](#dash-list){#line-styles:dash-list-5}, in coordinate space
    units.

8.  For each subpath `subpath`{.variable} in `path`{.variable}, run the
    following substeps. These substeps mutate the subpaths in
    `path`{.variable} *in vivo*.

    1.  Let `subpath width`{.variable} be the length of all the lines of
        `subpath`{.variable}, in coordinate space units.

    2.  Let `offset`{.variable} be the value of `style`{.variable}\'s
        [`lineDashOffset`{#line-styles:dom-context-2d-linedashoffset-3}](#dom-context-2d-linedashoffset),
        in coordinate space units.

    3.  While `offset`{.variable} is greater than
        `pattern width`{.variable}, decrement it by
        `pattern width`{.variable}.

        While `offset`{.variable} is less than zero, increment it by
        `pattern width`{.variable}.

    4.  Define `L`{.variable} to be a linear coordinate line defined
        along all lines in `subpath`{.variable}, such that the start of
        the first line in the subpath is defined as coordinate 0, and
        the end of the last line in the subpath is defined as coordinate
        `subpath width`{.variable}.

    5.  Let `position`{.variable} be zero minus `offset`{.variable}.

    6.  Let `index`{.variable} be 0.

    7.  Let `current state`{.variable} be *off* (the other states being
        *on* and *zero-on*).

    8.  *Dash on*: Let `segment length`{.variable} be the value of
        `style`{.variable}\'s [dash
        list](#dash-list){#line-styles:dash-list-6}\'s
        `index`{.variable}th entry.

    9.  Increment `position`{.variable} by `segment length`{.variable}.

    10. If `position`{.variable} is greater than
        `subpath width`{.variable}, then end these substeps for this
        subpath and start them again for the next subpath; if there are
        no more subpaths, then jump to the step labeled *convert*
        instead.

    11. If `segment length`{.variable} is nonzero, then let
        `current state`{.variable} be *on*.

    12. Increment `index`{.variable} by one.

    13. *Dash off*: Let `segment length`{.variable} be the value of
        `style`{.variable}\'s [dash
        list](#dash-list){#line-styles:dash-list-7}\'s
        `index`{.variable}th entry.

    14. Let `start`{.variable} be the offset `position`{.variable} on
        `L`{.variable}.

    15. Increment `position`{.variable} by `segment length`{.variable}.

    16. If `position`{.variable} is less than zero, then jump to the
        step labeled *post-cut*.

    17. If `start`{.variable} is less than zero, then let
        `start`{.variable} be zero.

    18. If `position`{.variable} is greater than
        `subpath width`{.variable}, then let `end`{.variable} be the
        offset `subpath width`{.variable} on `L`{.variable}. Otherwise,
        let `end`{.variable} be the offset `position`{.variable} on
        `L`{.variable}.

    19. Jump to the first appropriate step:

        If `segment length`{.variable} is zero and `current state`{.variable} is *off*

        :   Do nothing, just continue to the next step.

        If `current state`{.variable} is *off*

        :   Cut the line on which `end`{.variable} finds itself short at
            `end`{.variable} and place a point there, cutting in two the
            subpath that it was in; remove all line segments, joins,
            points, and subpaths that are between `start`{.variable} and
            `end`{.variable}; and finally place a single point at
            `start`{.variable} with no lines connecting to it.

            The point has a *directionality* for the purposes of drawing
            line caps (see below). The directionality is the direction
            that the original line had at that point (i.e. when
            `L`{.variable} was defined above).

        Otherwise

        :   Cut the line on which `start`{.variable} finds itself into
            two at `start`{.variable} and place a point there, cutting
            in two the subpath that it was in, and similarly cut the
            line on which `end`{.variable} finds itself short at
            `end`{.variable} and place a point there, cutting in two the
            subpath that *it* was in, and then remove all line segments,
            joins, points, and subpaths that are between
            `start`{.variable} and `end`{.variable}.

            If `start`{.variable} and `end`{.variable} are the same
            point, then this results in just the line being cut in two
            and two points being inserted there, with nothing being
            removed, unless a join also happens to be at that point, in
            which case the join must be removed.

    20. *Post-cut*: If `position`{.variable} is greater than
        `subpath width`{.variable}, then jump to the step labeled
        *convert*.

    21. Increment `index`{.variable} by one. If it is equal to the
        number of entries in `style`{.variable}\'s [dash
        list](#dash-list){#line-styles:dash-list-8}, then let
        `index`{.variable} be 0.

    22. Return to the step labeled *dash on*.

9.  *Convert*: This is the step that converts the path to a new path
    that represents its stroke.

    Create a new [path](#concept-path){#line-styles:concept-path-2} that
    describes the edge of the areas that would be covered if a straight
    line of length equal to `style`{.variable}\'s
    [`lineWidth`{#line-styles:dom-context-2d-linewidth-3}](#dom-context-2d-linewidth)
    was swept along each subpath in `path`{.variable} while being kept
    at an angle such that the line is orthogonal to the path being
    swept, replacing each point with the end cap necessary to satisfy
    `style`{.variable}\'s
    [`lineCap`{#line-styles:dom-context-2d-linecap-3}](#dom-context-2d-linecap)
    attribute as described previously and elaborated below, and
    replacing each join with the join necessary to satisfy
    `style`{.variable}\'s
    [`lineJoin`{#line-styles:dom-context-2d-linejoin-4}](#dom-context-2d-linejoin)
    type, as defined below.

    **Caps**: Each point has a flat edge perpendicular to the direction
    of the line coming out of it. This is then augmented according to
    the value of `style`{.variable}\'s
    [`lineCap`{#line-styles:dom-context-2d-linecap-4}](#dom-context-2d-linecap).
    The \"`butt`\" value means that no additional line cap is added. The
    \"`round`\" value means that a semi-circle with the diameter equal
    to `style`{.variable}\'s
    [`lineWidth`{#line-styles:dom-context-2d-linewidth-4}](#dom-context-2d-linewidth)
    width must additionally be placed on to the line coming out of each
    point. The \"`square`\" value means that a rectangle with the length
    of `style`{.variable}\'s
    [`lineWidth`{#line-styles:dom-context-2d-linewidth-5}](#dom-context-2d-linewidth)
    width and the width of half `style`{.variable}\'s
    [`lineWidth`{#line-styles:dom-context-2d-linewidth-6}](#dom-context-2d-linewidth)
    width, placed flat against the edge perpendicular to the direction
    of the line coming out of the point, must be added at each point.

    Points with no lines coming out of them must have two caps placed
    back-to-back as if it was really two points connected to each other
    by an infinitesimally short straight line in the direction of the
    point\'s *directionality* (as defined above).

    **Joins**: In addition to the point where a join occurs, two
    additional points are relevant to each join, one for each line: the
    two corners found half the line width away from the join point, one
    perpendicular to each line, each on the side furthest from the other
    line.

    A triangle connecting these two opposite corners with a straight
    line, with the third point of the triangle being the join point,
    must be added at all joins. The
    [`lineJoin`{#line-styles:dom-context-2d-linejoin-5}](#dom-context-2d-linejoin)
    attribute controls whether anything else is rendered. The three
    aforementioned values have the following meanings:

    The \"`bevel`\" value means that this is all that is rendered at
    joins.

    The \"`round`\" value means that an arc connecting the two
    aforementioned corners of the join, abutting (and not overlapping)
    the aforementioned triangle, with the diameter equal to the line
    width and the origin at the point of the join, must be added at
    joins.

    The \"`miter`\" value means that a second triangle must (if it can
    given the miter length) be added at the join, with one line being
    the line between the two aforementioned corners, abutting the first
    triangle, and the other two being continuations of the outside edges
    of the two joining lines, as long as required to intersect without
    going over the miter length.

    The miter length is the distance from the point where the join
    occurs to the intersection of the line edges on the outside of the
    join. The miter limit ratio is the maximum allowed ratio of the
    miter length to half the line width. If the miter length would cause
    the miter limit ratio (as set by `style`{.variable}\'s
    [`miterLimit`{#line-styles:dom-context-2d-miterlimit-3}](#dom-context-2d-miterlimit)
    attribute) to be exceeded, then this second triangle must not be
    added.

    The subpaths in the newly created path must be oriented such that
    for any point, the number of times a half-infinite straight line
    drawn from that point crosses a subpath is even if and only if the
    number of times a half-infinite straight line drawn from that same
    point crosses a subpath going in one direction is equal to the
    number of times it crosses a subpath going in the other direction.

10. Return the newly created path.

###### [4.12.5.1.5]{.secno} Text styles[](#text-styles){.self-link}

`context`{.variable}`.`[`lang`](#dom-context-2d-lang){#dom-context-2d-lang-dev}` [ = ``value`{.variable}` ]`

`styles`{.variable}`.`[`lang`](#dom-context-2d-lang){#text-styles:dom-context-2d-lang}` [ = ``value`{.variable}` ]`

Returns the current language setting.

Can be set to a BCP 47 language tag, the empty string, or
\"[`inherit`{#text-styles:dom-context-2d-lang-inherit}](#dom-context-2d-lang-inherit)\",
to change the language used when resolving fonts.
\"[`inherit`{#text-styles:dom-context-2d-lang-inherit-2}](#dom-context-2d-lang-inherit)\"
takes the language from the
[`canvas`{#text-styles:the-canvas-element}](#the-canvas-element)
element\'s language, or the associated [document
element](https://dom.spec.whatwg.org/#document-element){#text-styles:document-element
x-internal="document-element"} when there is no
[`canvas`{#text-styles:the-canvas-element-2}](#the-canvas-element)
element.

The default is
\"[`inherit`{#text-styles:dom-context-2d-lang-inherit-3}](#dom-context-2d-lang-inherit)\".

`context`{.variable}`.`[`font`](#dom-context-2d-font){#dom-context-2d-font-dev}` [ = ``value`{.variable}` ]`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/font](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/font "The CanvasRenderingContext2D.font property of the Canvas 2D API specifies the current text style to use when drawing text. This string uses the same syntax as the CSS font specifier.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome2+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

`styles`{.variable}`.`[`font`](#dom-context-2d-font){#text-styles:dom-context-2d-font}` [ = ``value`{.variable}` ]`

Returns the current font settings.

Can be set, to change the font. The syntax is the same as for the CSS
[\'font\'](https://drafts.csswg.org/css-fonts/#font-prop){#text-styles:'font'
x-internal="'font'"} property; values that cannot be parsed as CSS font
values are ignored. The default is \"10px sans-serif\".

Relative keywords and lengths are computed relative to the font of the
[`canvas`{#text-styles:the-canvas-element-3}](#the-canvas-element)
element.

`context`{.variable}`.`[`textAlign`](#dom-context-2d-textalign){#dom-context-2d-textalign-dev}` [ = ``value`{.variable}` ]`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/textAlign](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/textAlign "The CanvasRenderingContext2D.textAlign property of the Canvas 2D API specifies the current text alignment used when drawing text.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome2+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

`styles`{.variable}`.`[`textAlign`](#dom-context-2d-textalign){#text-styles:dom-context-2d-textalign}` [ = ``value`{.variable}` ]`

Returns the current text alignment settings.

Can be set, to change the alignment. The possible values are and their
meanings are given below. Other values are ignored. The default is
\"`start`\".

`context`{.variable}`.`[`textBaseline`](#dom-context-2d-textbaseline){#dom-context-2d-textbaseline-dev}` [ = ``value`{.variable}` ]`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/textBaseline](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/textBaseline "The CanvasRenderingContext2D.textBaseline property of the Canvas 2D API specifies the current text baseline used when drawing text.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome2+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

`styles`{.variable}`.`[`textBaseline`](#dom-context-2d-textbaseline){#text-styles:dom-context-2d-textbaseline}` [ = ``value`{.variable}` ]`

Returns the current baseline alignment settings.

Can be set, to change the baseline alignment. The possible values and
their meanings are given below. Other values are ignored. The default is
\"[`alphabetic`{#text-styles:dom-context-2d-textbaseline-alphabetic}](#dom-context-2d-textbaseline-alphabetic)\".

`context`{.variable}`.`[`direction`](#dom-context-2d-direction){#dom-context-2d-direction-dev}` [ = ``value`{.variable}` ]`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/direction](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/direction "The CanvasRenderingContext2D.direction property of the Canvas 2D API specifies the current text direction used to draw text.")

Support in all current engines.

::: support
[Firefox101+]{.firefox .yes}[Safari9+]{.safari .yes}[Chrome77+]{.chrome
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

`styles`{.variable}`.`[`direction`](#dom-context-2d-direction){#text-styles:dom-context-2d-direction}` [ = ``value`{.variable}` ]`

Returns the current directionality.

Can be set, to change the directionality. The possible values and their
meanings are given below. Other values are ignored. The default is
\"[`inherit`{#text-styles:dom-context-2d-direction-inherit}](#dom-context-2d-direction-inherit)\".

`context`{.variable}`.`[`letterSpacing`](#dom-context-2d-letterspacing){#dom-context-2d-letterspacing-dev}` [ = ``value`{.variable}` ]`

`styles`{.variable}`.`[`letterSpacing`](#dom-context-2d-letterspacing){#text-styles:dom-context-2d-letterspacing}` [ = ``value`{.variable}` ]`

Returns the current spacing between characters in the text.

Can be set, to change spacing between characters. Values that cannot be
parsed as a CSS
[\<length\>](https://drafts.csswg.org/css-values/#lengths){#text-styles:length-2
x-internal="length-2"} are ignored. The default is \"`0px`\".

`context`{.variable}`.`[`fontKerning`](#dom-context-2d-fontkerning){#dom-context-2d-fontkerning-dev}` [ = ``value`{.variable}` ]`

`styles`{.variable}`.`[`fontKerning`](#dom-context-2d-fontkerning){#text-styles:dom-context-2d-fontkerning}` [ = ``value`{.variable}` ]`

Returns the current font kerning settings.

Can be set, to change the font kerning. The possible values and their
meanings are given below. Other values are ignored. The default is
\"[`auto`{#text-styles:dom-context-2d-fontkerning-auto}](#dom-context-2d-fontkerning-auto)\".

`context`{.variable}`.`[`fontStretch`](#dom-context-2d-fontstretch){#dom-context-2d-fontstretch-dev}` [ = ``value`{.variable}` ]`

`styles`{.variable}`.`[`fontStretch`](#dom-context-2d-fontstretch){#text-styles:dom-context-2d-fontstretch}` [ = ``value`{.variable}` ]`

Returns the current font stretch settings.

Can be set, to change the font stretch. The possible values and their
meanings are given below. Other values are ignored. The default is
\"[`normal`{#text-styles:dom-context-2d-fontstretch-normal}](#dom-context-2d-fontstretch-normal)\".

`context`{.variable}`.`[`fontVariantCaps`](#dom-context-2d-fontvariantcaps){#dom-context-2d-fontvariantcaps-dev}` [ = ``value`{.variable}` ]`

`styles`{.variable}`.`[`fontVariantCaps`](#dom-context-2d-fontvariantcaps){#text-styles:dom-context-2d-fontvariantcaps}` [ = ``value`{.variable}` ]`

Returns the current font variant caps settings.

Can be set, to change the font variant caps. The possible values and
their meanings are given below. Other values are ignored. The default is
\"[`normal`{#text-styles:dom-context-2d-fontvariantcaps-normal}](#dom-context-2d-fontvariantcaps-normal)\".

`context`{.variable}`.`[`textRendering`](#dom-context-2d-textrendering){#dom-context-2d-textrendering-dev}` [ = ``value`{.variable}` ]`

`styles`{.variable}`.`[`textRendering`](#dom-context-2d-textrendering){#text-styles:dom-context-2d-textrendering}` [ = ``value`{.variable}` ]`

Returns the current text rendering settings.

Can be set, to change the text rendering. The possible values and their
meanings are given below. Other values are ignored. The default is
\"[`auto`{#text-styles:dom-context-2d-textrendering-auto}](#dom-context-2d-textrendering-auto)\".

`context`{.variable}`.`[`wordSpacing`](#dom-context-2d-wordspacing){#dom-context-2d-wordspacing-dev}` [ = ``value`{.variable}` ]`

`styles`{.variable}`.`[`wordSpacing`](#dom-context-2d-wordspacing){#text-styles:dom-context-2d-wordspacing}` [ = ``value`{.variable}` ]`

Returns the current spacing between words in the text.

Can be set, to change spacing between words. Values that cannot be
parsed as a CSS
[\<length\>](https://drafts.csswg.org/css-values/#lengths){#text-styles:length-2-2
x-internal="length-2"} are ignored. The default is \"`0px`\".

Objects that implement the
[`CanvasTextDrawingStyles`{#text-styles:canvastextdrawingstyles}](#canvastextdrawingstyles)
interface have attributes (defined in this section) that control how
text is laid out (rasterized or outlined) by the object. Such objects
can also have a [font style source object]{#font-style-source-object
.dfn}. For
[`CanvasRenderingContext2D`{#text-styles:canvasrenderingcontext2d}](#canvasrenderingcontext2d)
objects, this is the
[`canvas`{#text-styles:the-canvas-element-4}](#the-canvas-element)
element given by the value of the context\'s
[`canvas`{#text-styles:dom-context-2d-canvas}](#dom-context-2d-canvas)
attribute. For
[`OffscreenCanvasRenderingContext2D`{#text-styles:offscreencanvasrenderingcontext2d}](#offscreencanvasrenderingcontext2d)
objects, this is the [associated `OffscreenCanvas`
object](#associated-offscreencanvas-object){#text-styles:associated-offscreencanvas-object}.

Font resolution for the [font style source
object](#font-style-source-object){#text-styles:font-style-source-object}
requires a [font
source](https://drafts.csswg.org/css-font-loading/#font-source){#text-styles:font-source
x-internal="font-source"}. This is determined for a given
`object`{.variable} implementing
[`CanvasTextDrawingStyles`{#text-styles:canvastextdrawingstyles-2}](#canvastextdrawingstyles)
by the following steps:
[\[CSSFONTLOAD\]](references.html#refsCSSFONTLOAD)

1.  If `object`{.variable}\'s [font style source
    object](#font-style-source-object){#text-styles:font-style-source-object-2}
    is a
    [`canvas`{#text-styles:the-canvas-element-5}](#the-canvas-element)
    element, return the element\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#text-styles:node-document
    x-internal="node-document"}.

2.  Otherwise, `object`{.variable}\'s [font style source
    object](#font-style-source-object){#text-styles:font-style-source-object-3}
    is an
    [`OffscreenCanvas`{#text-styles:offscreencanvas}](#offscreencanvas)
    object:

    1.  Let `global`{.variable} be `object`{.variable}\'s [relevant
        global
        object](webappapis.html#concept-relevant-global){#text-styles:concept-relevant-global}.

    2.  If `global`{.variable} is a
        [`Window`{#text-styles:window}](nav-history-apis.html#window)
        object, then return `global`{.variable}\'s [associated
        `Document`](nav-history-apis.html#concept-document-window){#text-styles:concept-document-window}.

    3.  [Assert](https://infra.spec.whatwg.org/#assert){#text-styles:assert
        x-internal="assert"}: `global`{.variable} implements
        [`WorkerGlobalScope`{#text-styles:workerglobalscope}](workers.html#workerglobalscope).

    4.  Return `global`{.variable}.

::: example
This is an example of font resolution with a regular
[`canvas`{#text-styles:the-canvas-element-6}](#the-canvas-element)
element with ID `c1`.

``` js
const font = new FontFace("MyCanvasFont", "url(mycanvasfont.ttf)");
documents.fonts.add(font);

const context = document.getElementById("c1").getContext("2d");
document.fonts.ready.then(function() {
  context.font = "64px MyCanvasFont";
  context.fillText("hello", 0, 0);
});
```

In this example, the canvas will display text using `mycanvasfont.ttf`
as its font.
:::

::: example
This is an example of how font resolution can happen using
[`OffscreenCanvas`{#text-styles:offscreencanvas-2}](#offscreencanvas).
Assuming a
[`canvas`{#text-styles:the-canvas-element-7}](#the-canvas-element)
element with ID `c2` which is transferred to a worker like so:

``` js
const offscreenCanvas = document.getElementById("c2").transferControlToOffscreen();
worker.postMessage(offscreenCanvas, [offscreenCanvas]);
```

Then, in the worker:

``` js
self.onmessage = function(ev) {
  const transferredCanvas = ev.data;
  const context = transferredCanvas.getContext("2d");
  const font = new FontFace("MyFont", "url(myfont.ttf)");
  self.fonts.add(font);
  self.fonts.ready.then(function() {
    context.font = "64px MyFont";
    context.fillText("hello", 0, 0);
  });
};
```

In this example, the canvas will display a text using `myfont.ttf`.
Notice that the font is only loaded inside the worker, and not in the
document context.
:::

The [`font`]{#dom-context-2d-font .dfn dfn-for="CanvasTextDrawingStyles"
dfn-type="attribute"} IDL attribute, on setting, must be [parsed as a
CSS \<\'font\'\>
value](https://drafts.csswg.org/css-syntax/#parse-grammar){#text-styles:parse-something-according-to-a-css-grammar
x-internal="parse-something-according-to-a-css-grammar"} (but without
supporting property-independent style sheet syntax like \'inherit\'),
and the resulting font must be assigned to the context, with the
[\'line-height\'](https://drafts.csswg.org/css2/#propdef-line-height){#text-styles:'line-height'
x-internal="'line-height'"} component forced to \'normal\', with the
[\'font-size\'](https://drafts.csswg.org/css-fonts/#font-size-prop){#text-styles:'font-size'
x-internal="'font-size'"} component converted to [CSS
pixels](https://drafts.csswg.org/css-values/#px){#text-styles:'px'
x-internal="'px'"}, and with system fonts being computed to explicit
values. If the new value is syntactically incorrect (including using
property-independent style sheet syntax like \'inherit\' or
\'initial\'), then it must be ignored, without assigning a new font
value. [\[CSS\]](references.html#refsCSS)

Font family names must be interpreted in the context of the [font style
source
object](#font-style-source-object){#text-styles:font-style-source-object-4}
when the font is to be used; any fonts embedded using `@font-face` or
loaded using
[`FontFace`{#text-styles:fontface}](infrastructure.html#fontface)
objects that are visible to the [font style source
object](#font-style-source-object){#text-styles:font-style-source-object-5}
must therefore be available once they are loaded. (Each [font style
source
object](#font-style-source-object){#text-styles:font-style-source-object-6}
has a [font
source](https://drafts.csswg.org/css-font-loading/#font-source){#text-styles:font-source-2
x-internal="font-source"}, which determines what fonts are available.)
If a font is used before it is fully loaded, or if the [font style
source
object](#font-style-source-object){#text-styles:font-style-source-object-7}
does not have that font in scope at the time the font is to be used,
then it must be treated as if it was an unknown font, falling back to
another as described by the relevant CSS specifications.
[\[CSSFONTS\]](references.html#refsCSSFONTS)
[\[CSSFONTLOAD\]](references.html#refsCSSFONTLOAD)

On getting, the
[`font`{#text-styles:dom-context-2d-font-2}](#dom-context-2d-font)
attribute must return the [serialized
form](https://drafts.csswg.org/cssom/#serialize-a-css-value){#text-styles:serializing-a-css-value
x-internal="serializing-a-css-value"} of the current font of the context
(with no
[\'line-height\'](https://drafts.csswg.org/css2/#propdef-line-height){#text-styles:'line-height'-2
x-internal="'line-height'"} component).
[\[CSSOM\]](references.html#refsCSSOM)

::: example
For example, after the following statement:

``` js
context.font = 'italic 400 12px/2 Unknown Font, sans-serif';
```

\...the expression `context.font` would evaluate to the string
\"`italic 12px "Unknown Font", sans-serif`\". The \"400\" font-weight
doesn\'t appear because that is the default value. The line-height
doesn\'t appear because it is forced to \"normal\", the default value.
:::

When the object implementing the
[`CanvasTextDrawingStyles`{#text-styles:canvastextdrawingstyles-3}](#canvastextdrawingstyles)
interface is created, the font of the context must be set to 10px
sans-serif. When the
[\'font-size\'](https://drafts.csswg.org/css-fonts/#font-size-prop){#text-styles:'font-size'-2
x-internal="'font-size'"} component is set to lengths using percentages,
[\'em\'](https://drafts.csswg.org/css-values/#em){#text-styles:'em'
x-internal="'em'"} or
[\'ex\'](https://drafts.csswg.org/css-values/#ex){#text-styles:'ex'
x-internal="'ex'"} units, or the \'larger\' or \'smaller\' keywords,
these must be interpreted relative to the [computed
value](https://drafts.csswg.org/css-cascade/#computed-value){#text-styles:computed-value
x-internal="computed-value"} of the
[\'font-size\'](https://drafts.csswg.org/css-fonts/#font-size-prop){#text-styles:'font-size'-3
x-internal="'font-size'"} property of the [font style source
object](#font-style-source-object){#text-styles:font-style-source-object-8}
at the time that the attribute is set, if it is an element. When the
[\'font-weight\'](https://drafts.csswg.org/css-fonts/#font-weight-prop){#text-styles:'font-weight'
x-internal="'font-weight'"} component is set to the relative values
\'bolder\' and \'lighter\', these must be interpreted relative to the
[computed
value](https://drafts.csswg.org/css-cascade/#computed-value){#text-styles:computed-value-2
x-internal="computed-value"} of the
[\'font-weight\'](https://drafts.csswg.org/css-fonts/#font-weight-prop){#text-styles:'font-weight'-2
x-internal="'font-weight'"} property of the [font style source
object](#font-style-source-object){#text-styles:font-style-source-object-9}
at the time that the attribute is set, if it is an element. If the
[computed
values](https://drafts.csswg.org/css-cascade/#computed-value){#text-styles:computed-value-3
x-internal="computed-value"} are undefined for a particular case (e.g.
because the [font style source
object](#font-style-source-object){#text-styles:font-style-source-object-10}
is not an element or is not [being
rendered](rendering.html#being-rendered){#text-styles:being-rendered}),
then the relative keywords must be interpreted relative to the
normal-weight 10px sans-serif default.

The [`textAlign`]{#dom-context-2d-textalign .dfn
dfn-for="CanvasTextDrawingStyles" dfn-type="attribute"} IDL attribute,
on getting, must return the current value. On setting, the current value
must be changed to the new value. When the object implementing the
[`CanvasTextDrawingStyles`{#text-styles:canvastextdrawingstyles-4}](#canvastextdrawingstyles)
interface is created, the
[`textAlign`{#text-styles:dom-context-2d-textalign-2}](#dom-context-2d-textalign)
attribute must initially have the value
[`start`{#text-styles:dom-context-2d-textalign-start}](#dom-context-2d-textalign-start).

The [`textBaseline`]{#dom-context-2d-textbaseline .dfn
dfn-for="CanvasTextDrawingStyles" dfn-type="attribute"} IDL attribute,
on getting, must return the current value. On setting, the current value
must be changed to the new value. When the object implementing the
[`CanvasTextDrawingStyles`{#text-styles:canvastextdrawingstyles-5}](#canvastextdrawingstyles)
interface is created, the
[`textBaseline`{#text-styles:dom-context-2d-textbaseline-2}](#dom-context-2d-textbaseline)
attribute must initially have the value
[`alphabetic`{#text-styles:dom-context-2d-textbaseline-alphabetic-2}](#dom-context-2d-textbaseline-alphabetic).

Objects that implement the
[`CanvasTextDrawingStyles`{#text-styles:canvastextdrawingstyles-6}](#canvastextdrawingstyles)
interface have an associated
[language]{#concept-canvastextdrawingstyles-language .dfn} value used to
localize font rendering. Valid values are a BCP 47 language tag, the
empty string, or \"[inherit]{#dom-context-2d-lang-inherit .dfn}\" where
the language comes from the
[`canvas`{#text-styles:the-canvas-element-8}](#the-canvas-element)
element\'s language, or the associated [document
element](https://dom.spec.whatwg.org/#document-element){#text-styles:document-element-2
x-internal="document-element"} when there is no
[`canvas`{#text-styles:the-canvas-element-9}](#the-canvas-element)
element. Initially, the
[language](#concept-canvastextdrawingstyles-language){#text-styles:concept-canvastextdrawingstyles-language}
must be
\"[`inherit`{#text-styles:dom-context-2d-lang-inherit-4}](#dom-context-2d-lang-inherit)\".

The [`lang`]{#dom-context-2d-lang .dfn dfn-for="CanvasTextDrawingStyles"
dfn-type="attribute"} getter steps are to return
[this](https://webidl.spec.whatwg.org/#this){#text-styles:this
x-internal="this"}\'s
[language](#concept-canvastextdrawingstyles-language){#text-styles:concept-canvastextdrawingstyles-language-2}.

The [`lang`{#text-styles:dom-context-2d-lang-2}](#dom-context-2d-lang)
setter steps are to set
[this](https://webidl.spec.whatwg.org/#this){#text-styles:this-2
x-internal="this"}\'s
[language](#concept-canvastextdrawingstyles-language){#text-styles:concept-canvastextdrawingstyles-language-3}
to the given value.

The [`direction`]{#dom-context-2d-direction .dfn
dfn-for="CanvasTextDrawingStyles" dfn-type="attribute"} IDL attribute,
on getting, must return the current value. On setting, the current value
must be changed to the new value. When the object implementing the
[`CanvasTextDrawingStyles`{#text-styles:canvastextdrawingstyles-7}](#canvastextdrawingstyles)
interface is created, the
[`direction`{#text-styles:dom-context-2d-direction-2}](#dom-context-2d-direction)
attribute must initially have the value
\"[`inherit`{#text-styles:dom-context-2d-direction-inherit-2}](#dom-context-2d-direction-inherit)\".

Objects that implement the
[`CanvasTextDrawingStyles`{#text-styles:canvastextdrawingstyles-8}](#canvastextdrawingstyles)
interface have attributes that control the spacing between letters and
words. Such objects have associated [letter
spacing]{#concept-canvastextdrawingstyles-letter-spacing .dfn} and [word
spacing]{#concept-canvastextdrawingstyles-word-spacing .dfn} values,
which are CSS
[\<length\>](https://drafts.csswg.org/css-values/#lengths){#text-styles:length-2-3
x-internal="length-2"} values. Initially, both must be the result of
[parsing](https://drafts.csswg.org/css-syntax/#parse-grammar){#text-styles:parse-something-according-to-a-css-grammar-2
x-internal="parse-something-according-to-a-css-grammar"} \"`0px`\" as a
CSS
[\<length\>](https://drafts.csswg.org/css-values/#lengths){#text-styles:length-2-4
x-internal="length-2"}.

::::: {.mdn-anno .wrapped .before}
**⚠**MDN

:::: feature
[CanvasRenderingContext2D/letterSpacing](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/letterSpacing "The CanvasRenderingContext2D.letterSpacing property of the Canvas API specifies the spacing between letters when drawing text.")

Support in one engine only.

::: support
[FirefoxNo]{.firefox .no}[SafariNo]{.safari .no}[Chrome99+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge99+]{.edge_blink .yes}

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

The [`letterSpacing`]{#dom-context-2d-letterspacing .dfn
dfn-for="CanvasTextDrawingStyles" dfn-type="attribute"} getter steps are
to return the [serialized
form](https://drafts.csswg.org/cssom/#serialize-a-css-value){#text-styles:serializing-a-css-value-2
x-internal="serializing-a-css-value"} of
[this](https://webidl.spec.whatwg.org/#this){#text-styles:this-3
x-internal="this"}\'s [letter
spacing](#concept-canvastextdrawingstyles-letter-spacing){#text-styles:concept-canvastextdrawingstyles-letter-spacing}.

The
[`letterSpacing`{#text-styles:dom-context-2d-letterspacing-2}](#dom-context-2d-letterspacing)
setter steps are:

1.  Let `parsed`{.variable} be the result of
    [parsing](https://drafts.csswg.org/css-syntax/#parse-grammar){#text-styles:parse-something-according-to-a-css-grammar-3
    x-internal="parse-something-according-to-a-css-grammar"} the given
    value as a CSS
    [\<length\>](https://drafts.csswg.org/css-values/#lengths){#text-styles:length-2-5
    x-internal="length-2"}.

2.  If `parsed`{.variable} is failure, then return.

3.  Set [this](https://webidl.spec.whatwg.org/#this){#text-styles:this-4
    x-internal="this"}\'s [letter
    spacing](#concept-canvastextdrawingstyles-letter-spacing){#text-styles:concept-canvastextdrawingstyles-letter-spacing-2}
    to `parsed`{.variable}.

::::: {.mdn-anno .wrapped .before}
**⚠**MDN

:::: feature
[CanvasRenderingContext2D/wordSpacing](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/wordSpacing "The CanvasRenderingContext2D.wordSpacing property of the Canvas API specifies the spacing between words when drawing text.")

Support in one engine only.

::: support
[FirefoxNo]{.firefox .no}[SafariNo]{.safari .no}[Chrome99+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge99+]{.edge_blink .yes}

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

The [`wordSpacing`]{#dom-context-2d-wordspacing .dfn
dfn-for="CanvasTextDrawingStyles" dfn-type="attribute"} getter steps are
to return the [serialized
form](https://drafts.csswg.org/cssom/#serialize-a-css-value){#text-styles:serializing-a-css-value-3
x-internal="serializing-a-css-value"} of
[this](https://webidl.spec.whatwg.org/#this){#text-styles:this-5
x-internal="this"}\'s [word
spacing](#concept-canvastextdrawingstyles-word-spacing){#text-styles:concept-canvastextdrawingstyles-word-spacing}.

The
[`wordSpacing`{#text-styles:dom-context-2d-wordspacing-2}](#dom-context-2d-wordspacing)
setter steps are:

1.  Let `parsed`{.variable} be the result of
    [parsing](https://drafts.csswg.org/css-syntax/#parse-grammar){#text-styles:parse-something-according-to-a-css-grammar-4
    x-internal="parse-something-according-to-a-css-grammar"} the given
    value as a CSS
    [\<length\>](https://drafts.csswg.org/css-values/#lengths){#text-styles:length-2-6
    x-internal="length-2"}.

2.  If `parsed`{.variable} is failure, then return.

3.  Set [this](https://webidl.spec.whatwg.org/#this){#text-styles:this-6
    x-internal="this"}\'s [word
    spacing](#concept-canvastextdrawingstyles-word-spacing){#text-styles:concept-canvastextdrawingstyles-word-spacing-2}
    to `parsed`{.variable}.

::::: {.mdn-anno .wrapped .before}
MDN

:::: feature
[CanvasRenderingContext2D/fontKerning](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/fontKerning "The CanvasRenderingContext2D.fontKerning property of the Canvas API specifies how font kerning information is used.")

::: support
[Firefox104+]{.firefox .yes}[SafariNo]{.safari .no}[Chrome99+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge99+]{.edge_blink .yes}

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

The [`fontKerning`]{#dom-context-2d-fontkerning .dfn
dfn-for="CanvasTextDrawingStyles" dfn-type="attribute"} IDL attribute,
on getting, must return the current value. On setting, the current value
must be changed to the new value. When the object implementing the
[`CanvasTextDrawingStyles`{#text-styles:canvastextdrawingstyles-9}](#canvastextdrawingstyles)
interface is created, the
[`fontKerning`{#text-styles:dom-context-2d-fontkerning-2}](#dom-context-2d-fontkerning)
attribute must initially have the value
\"[`auto`{#text-styles:dom-context-2d-fontkerning-auto-2}](#dom-context-2d-fontkerning-auto)\".

::::: {.mdn-anno .wrapped .before}
**⚠**MDN

:::: feature
[CanvasRenderingContext2D/fontStretch](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/fontStretch "The CanvasRenderingContext2D.fontStretch property of the Canvas API specifies how the font may be expanded or condensed when drawing text.")

Support in one engine only.

::: support
[FirefoxNo]{.firefox .no}[SafariNo]{.safari .no}[Chrome99+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge99+]{.edge_blink .yes}

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

The [`fontStretch`]{#dom-context-2d-fontstretch .dfn
dfn-for="CanvasTextDrawingStyles" dfn-type="attribute"} IDL attribute,
on getting, must return the current value. On setting, the current value
must be changed to the new value. When the object implementing the
[`CanvasTextDrawingStyles`{#text-styles:canvastextdrawingstyles-10}](#canvastextdrawingstyles)
interface is created, the
[`fontStretch`{#text-styles:dom-context-2d-fontstretch-2}](#dom-context-2d-fontstretch)
attribute must initially have the value
\"[`normal`{#text-styles:dom-context-2d-fontstretch-normal-2}](#dom-context-2d-fontstretch-normal)\".

::::: {.mdn-anno .wrapped .before}
**⚠**MDN

:::: feature
[CanvasRenderingContext2D/fontVariantCaps](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/fontVariantCaps "The CanvasRenderingContext2D.fontVariantCaps property of the Canvas API specifies an alternative capitalization of the rendered text.")

Support in one engine only.

::: support
[FirefoxNo]{.firefox .no}[SafariNo]{.safari .no}[Chrome99+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge99+]{.edge_blink .yes}

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

The [`fontVariantCaps`]{#dom-context-2d-fontvariantcaps .dfn
dfn-for="CanvasTextDrawingStyles" dfn-type="attribute"} IDL attribute,
on getting, must return the current value. On setting, the current value
must be changed to the new value. When the object implementing the
[`CanvasTextDrawingStyles`{#text-styles:canvastextdrawingstyles-11}](#canvastextdrawingstyles)
interface is created, the
[`fontVariantCaps`{#text-styles:dom-context-2d-fontvariantcaps-2}](#dom-context-2d-fontvariantcaps)
attribute must initially have the value
\"[`normal`{#text-styles:dom-context-2d-fontvariantcaps-normal-2}](#dom-context-2d-fontvariantcaps-normal)\".

::::: {.mdn-anno .wrapped .before}
**⚠**MDN

:::: feature
[CanvasRenderingContext2D/textRendering](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/textRendering "The CanvasRenderingContext2D.textRendering property of the Canvas API provides information to the rendering engine about what to optimize for when rendering text.")

Support in one engine only.

::: support
[FirefoxNo]{.firefox .no}[SafariNo]{.safari .no}[Chrome99+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge99+]{.edge_blink .yes}

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

The [`textRendering`]{#dom-context-2d-textrendering .dfn
dfn-for="CanvasTextDrawingStyles" dfn-type="attribute"} IDL attribute,
on getting, must return the current value. On setting, the current value
must be changed to the new value. When the object implementing the
[`CanvasTextDrawingStyles`{#text-styles:canvastextdrawingstyles-12}](#canvastextdrawingstyles)
interface is created, the
[`textRendering`{#text-styles:dom-context-2d-textrendering-2}](#dom-context-2d-textrendering)
attribute must initially have the value
\"[`auto`{#text-styles:dom-context-2d-textrendering-auto-2}](#dom-context-2d-textrendering-auto)\".

The
[`textAlign`{#text-styles:dom-context-2d-textalign-3}](#dom-context-2d-textalign)
attribute\'s allowed keywords are as follows:

[`start`]{#dom-context-2d-textalign-start .dfn dfn-for="CanvasTextAlign" dfn-type="enum-value"}
:   Align to the start edge of the text (left side in left-to-right
    text, right side in right-to-left text).

[`end`]{#dom-context-2d-textalign-end .dfn dfn-for="CanvasTextAlign" dfn-type="enum-value"}
:   Align to the end edge of the text (right side in left-to-right text,
    left side in right-to-left text).

[`left`]{#dom-context-2d-textalign-left .dfn dfn-for="CanvasTextAlign" dfn-type="enum-value"}
:   Align to the left.

[`right`]{#dom-context-2d-textalign-right .dfn dfn-for="CanvasTextAlign" dfn-type="enum-value"}
:   Align to the right.

[`center`]{#dom-context-2d-textalign-center .dfn dfn-for="CanvasTextAlign" dfn-type="enum-value"}

:   Align to the center.

The
[`textBaseline`{#text-styles:dom-context-2d-textbaseline-3}](#dom-context-2d-textbaseline)
attribute\'s allowed keywords correspond to alignment points in the
font:

![The em-over baseline is roughly at the top of the glyphs in a font,
the hanging baseline is where some glyphs like आ are anchored, the
middle is half-way between the em-over and em-under baselines, the
alphabetic baseline is where characters like Á, ÿ, f, and Ω are
anchored, the ideographic-under baseline is where glyphs like 私 and 達
are anchored, and the em-under baseline is roughly at the bottom of the
glyphs in a font. The top and bottom of the bounding box can be far from
these baselines, due to glyphs extending far outside em-over and
em-under baselines.](/images/baselines.png){width="738" height="300"}

The keywords map to these alignment points as follows:

[`top`]{#dom-context-2d-textbaseline-top .dfn dfn-for="CanvasTextBaseline" dfn-type="enum-value"}
:   The [em-over
    baseline](https://drafts.csswg.org/css-inline/#em-over-baseline){#text-styles:em-over-baseline
    x-internal="em-over-baseline"}

[`hanging`]{#dom-context-2d-textbaseline-hanging .dfn dfn-for="CanvasTextBaseline" dfn-type="enum-value"}
:   The [hanging
    baseline](https://drafts.csswg.org/css-inline/#hanging-baseline){#text-styles:hanging-baseline
    x-internal="hanging-baseline"}

[`middle`]{#dom-context-2d-textbaseline-middle .dfn dfn-for="CanvasTextBaseline" dfn-type="enum-value"}
:   Halfway between the [em-over
    baseline](https://drafts.csswg.org/css-inline/#em-over-baseline){#text-styles:em-over-baseline-2
    x-internal="em-over-baseline"} and the [em-under
    baseline](https://drafts.csswg.org/css-inline/#em-under-baseline){#text-styles:em-under-baseline
    x-internal="em-under-baseline"}

[`alphabetic`]{#dom-context-2d-textbaseline-alphabetic .dfn dfn-for="CanvasTextBaseline" dfn-type="enum-value"}
:   The [alphabetic
    baseline](https://drafts.csswg.org/css-inline/#alphabetic-baseline){#text-styles:alphabetic-baseline
    x-internal="alphabetic-baseline"}

[`ideographic`]{#dom-context-2d-textbaseline-ideographic .dfn dfn-for="CanvasTextBaseline" dfn-type="enum-value"}
:   The [ideographic-under
    baseline](https://drafts.csswg.org/css-inline/#ideographic-under-baseline){#text-styles:ideographic-under-baseline
    x-internal="ideographic-under-baseline"}

[`bottom`]{#dom-context-2d-textbaseline-bottom .dfn dfn-for="CanvasTextBaseline" dfn-type="enum-value"}
:   The [em-under
    baseline](https://drafts.csswg.org/css-inline/#em-under-baseline){#text-styles:em-under-baseline-2
    x-internal="em-under-baseline"}

The
[`direction`{#text-styles:dom-context-2d-direction-3}](#dom-context-2d-direction)
attribute\'s allowed keywords are as follows:

[`ltr`]{#dom-context-2d-direction-ltr .dfn dfn-for="CanvasDirection" dfn-type="enum-value"}
:   Treat input to the [text preparation
    algorithm](#text-preparation-algorithm){#text-styles:text-preparation-algorithm}
    as left-to-right text.

[`rtl`]{#dom-context-2d-direction-rtl .dfn dfn-for="CanvasDirection" dfn-type="enum-value"}
:   Treat input to the [text preparation
    algorithm](#text-preparation-algorithm){#text-styles:text-preparation-algorithm-2}
    as right-to-left text.

[`inherit`]{#dom-context-2d-direction-inherit .dfn dfn-for="CanvasDirection" dfn-type="enum-value"}

:   Use the process in the [text preparation
    algorithm](#text-preparation-algorithm){#text-styles:text-preparation-algorithm-3}
    to obtain the text direction from the
    [`canvas`{#text-styles:the-canvas-element-10}](#the-canvas-element)
    element, [placeholder `canvas`
    element](#offscreencanvas-placeholder){#text-styles:offscreencanvas-placeholder},
    or [document
    element](https://dom.spec.whatwg.org/#document-element){#text-styles:document-element-3
    x-internal="document-element"}.

The
[`fontKerning`{#text-styles:dom-context-2d-fontkerning-3}](#dom-context-2d-fontkerning)
attribute\'s allowed keywords are as follows:

[`auto`]{#dom-context-2d-fontkerning-auto .dfn dfn-for="CanvasFontKerning" dfn-type="enum-value"}
:   Kerning is applied at the discretion of the user agent.

[`normal`]{#dom-context-2d-fontkerning-normal .dfn dfn-for="CanvasFontKerning" dfn-type="enum-value"}
:   Kerning is applied.

[`none`]{#dom-context-2d-fontkerning-none .dfn dfn-for="CanvasFontKerning" dfn-type="enum-value"}

:   Kerning is not applied.

The
[`fontStretch`{#text-styles:dom-context-2d-fontstretch-3}](#dom-context-2d-fontstretch)
attribute\'s allowed keywords are as follows:

[`ultra-condensed`]{#dom-context-2d-fontstretch-ultra-condensed .dfn dfn-for="CanvasFontStretch" dfn-type="enum-value"}
:   Same as CSS
    [\'font-stretch\'](https://drafts.csswg.org/css-fonts/#propdef-font-stretch){#text-styles:'font-stretch'
    x-internal="'font-stretch'"}
    [\'ultra-condensed\'](https://drafts.csswg.org/css-fonts/#valdef-font-stretch-ultra-condensed){#text-styles:'ultra-condensed'
    x-internal="'ultra-condensed'"} setting.

[`extra-condensed`]{#dom-context-2d-fontstretch-extra-condensed .dfn dfn-for="CanvasFontStretch" dfn-type="enum-value"}
:   Same as CSS
    [\'font-stretch\'](https://drafts.csswg.org/css-fonts/#propdef-font-stretch){#text-styles:'font-stretch'-2
    x-internal="'font-stretch'"}
    [\'extra-condensed\'](https://drafts.csswg.org/css-fonts/#valdef-font-stretch-extra-condensed){#text-styles:'extra-condensed'
    x-internal="'extra-condensed'"} setting.

[`condensed`]{#dom-context-2d-fontstretch-condensed .dfn dfn-for="CanvasFontStretch" dfn-type="enum-value"}
:   Same as CSS
    [\'font-stretch\'](https://drafts.csswg.org/css-fonts/#propdef-font-stretch){#text-styles:'font-stretch'-3
    x-internal="'font-stretch'"}
    [\'condensed\'](https://drafts.csswg.org/css-fonts/#valdef-font-stretch-condensed){#text-styles:'condensed'
    x-internal="'condensed'"} setting.

[`semi-condensed`]{#dom-context-2d-fontstretch-semi-condensed .dfn dfn-for="CanvasFontStretch" dfn-type="enum-value"}
:   Same as CSS
    [\'font-stretch\'](https://drafts.csswg.org/css-fonts/#propdef-font-stretch){#text-styles:'font-stretch'-4
    x-internal="'font-stretch'"}
    [\'semi-condensed\'](https://drafts.csswg.org/css-fonts/#valdef-font-stretch-semi-condensed){#text-styles:'semi-condensed'
    x-internal="'semi-condensed'"} setting.

[`normal`]{#dom-context-2d-fontstretch-normal .dfn dfn-for="CanvasFontStretch" dfn-type="enum-value"}
:   The default setting, where width of the glyphs is at 100%.

[`semi-expanded`]{#dom-context-2d-fontstretch-semi-expanded .dfn dfn-for="CanvasFontStretch" dfn-type="enum-value"}
:   Same as CSS
    [\'font-stretch\'](https://drafts.csswg.org/css-fonts/#propdef-font-stretch){#text-styles:'font-stretch'-5
    x-internal="'font-stretch'"}
    [\'semi-expanded\'](https://drafts.csswg.org/css-fonts/#valdef-font-stretch-semi-expanded){#text-styles:'semi-expanded'
    x-internal="'semi-expanded'"} setting.

[`expanded`]{#dom-context-2d-fontstretch-expanded .dfn dfn-for="CanvasFontStretch" dfn-type="enum-value"}
:   Same as CSS
    [\'font-stretch\'](https://drafts.csswg.org/css-fonts/#propdef-font-stretch){#text-styles:'font-stretch'-6
    x-internal="'font-stretch'"}
    [\'expanded\'](https://drafts.csswg.org/css-fonts/#valdef-font-stretch-expanded){#text-styles:'expanded'
    x-internal="'expanded'"} setting.

[`extra-expanded`]{#dom-context-2d-fontstretch-extra-expanded .dfn dfn-for="CanvasFontStretch" dfn-type="enum-value"}
:   Same as CSS
    [\'font-stretch\'](https://drafts.csswg.org/css-fonts/#propdef-font-stretch){#text-styles:'font-stretch'-7
    x-internal="'font-stretch'"}
    [\'extra-expanded\'](https://drafts.csswg.org/css-fonts/#valdef-font-stretch-extra-expanded){#text-styles:'extra-expanded'
    x-internal="'extra-expanded'"} setting.

[`ultra-expanded`]{#dom-context-2d-fontstretch-ultra-expanded .dfn dfn-for="CanvasFontStretch" dfn-type="enum-value"}

:   Same as CSS
    [\'font-stretch\'](https://drafts.csswg.org/css-fonts/#propdef-font-stretch){#text-styles:'font-stretch'-8
    x-internal="'font-stretch'"}
    [\'ultra-expanded\'](https://drafts.csswg.org/css-fonts/#valdef-font-stretch-ultra-expanded){#text-styles:'ultra-expanded'
    x-internal="'ultra-expanded'"} setting.

The
[`fontVariantCaps`{#text-styles:dom-context-2d-fontvariantcaps-3}](#dom-context-2d-fontvariantcaps)
attribute\'s allowed keywords are as follows:

[`normal`]{#dom-context-2d-fontvariantcaps-normal .dfn dfn-for="CanvasFontVariantCaps" dfn-type="enum-value"}
:   None of the features listed below are enabled.

[`small-caps`]{#dom-context-2d-fontvariantcaps-small-caps .dfn dfn-for="CanvasFontVariantCaps" dfn-type="enum-value"}
:   Same as CSS
    [\'font-variant-caps\'](https://drafts.csswg.org/css-fonts/#propdef-font-variant-caps){#text-styles:'font-variant-caps'
    x-internal="'font-variant-caps'"}
    [\'small-caps\'](https://drafts.csswg.org/css-fonts/#valdef-font-variant-caps-small-caps){#text-styles:'small-caps'
    x-internal="'small-caps'"} setting.

[`all-small-caps`]{#dom-context-2d-fontvariantcaps-all-small-caps .dfn dfn-for="CanvasFontVariantCaps" dfn-type="enum-value"}
:   Same as CSS
    [\'font-variant-caps\'](https://drafts.csswg.org/css-fonts/#propdef-font-variant-caps){#text-styles:'font-variant-caps'-2
    x-internal="'font-variant-caps'"}
    [\'all-small-caps\'](https://drafts.csswg.org/css-fonts/#valdef-font-variant-caps-all-small-caps){#text-styles:'all-small-caps'
    x-internal="'all-small-caps'"} setting.

[`petite-caps`]{#dom-context-2d-fontvariantcaps-petite-caps .dfn dfn-for="CanvasFontVariantCaps" dfn-type="enum-value"}
:   Same as CSS
    [\'font-variant-caps\'](https://drafts.csswg.org/css-fonts/#propdef-font-variant-caps){#text-styles:'font-variant-caps'-3
    x-internal="'font-variant-caps'"}
    [\'petite-caps\'](https://drafts.csswg.org/css-fonts/#valdef-font-variant-caps-petite-caps){#text-styles:'petite-caps'
    x-internal="'petite-caps'"} setting.

[`all-petite-caps`]{#dom-context-2d-fontvariantcaps-all-petite-caps .dfn dfn-for="CanvasFontVariantCaps" dfn-type="enum-value"}
:   Same as CSS
    [\'font-variant-caps\'](https://drafts.csswg.org/css-fonts/#propdef-font-variant-caps){#text-styles:'font-variant-caps'-4
    x-internal="'font-variant-caps'"}
    [\'all-petite-caps\'](https://drafts.csswg.org/css-fonts/#valdef-font-variant-caps-all-petite-caps){#text-styles:'all-petite-caps'
    x-internal="'all-petite-caps'"} setting.

[`unicase`]{#dom-context-2d-fontvariantcaps-unicase .dfn dfn-for="CanvasFontVariantCaps" dfn-type="enum-value"}
:   Same as CSS
    [\'font-variant-caps\'](https://drafts.csswg.org/css-fonts/#propdef-font-variant-caps){#text-styles:'font-variant-caps'-5
    x-internal="'font-variant-caps'"}
    [\'unicase\'](https://drafts.csswg.org/css-fonts/#valdef-font-variant-caps-unicase){#text-styles:'unicase'
    x-internal="'unicase'"} setting.

[`titling-caps`]{#dom-context-2d-fontvariantcaps-titling-caps .dfn dfn-for="CanvasFontVariantCaps" dfn-type="enum-value"}

:   Same as CSS
    [\'font-variant-caps\'](https://drafts.csswg.org/css-fonts/#propdef-font-variant-caps){#text-styles:'font-variant-caps'-6
    x-internal="'font-variant-caps'"}
    [\'titling-caps\'](https://drafts.csswg.org/css-fonts/#valdef-font-variant-caps-titling-caps){#text-styles:'titling-caps'
    x-internal="'titling-caps'"} setting.

The
[`textRendering`{#text-styles:dom-context-2d-textrendering-3}](#dom-context-2d-textrendering)
attribute\'s allowed keywords are as follows:

[`auto`]{#dom-context-2d-textrendering-auto .dfn dfn-for="CanvasTextRendering" dfn-type="enum-value"}
:   Same as \'auto\' in [SVG
    text-rendering](https://svgwg.org/svg2-draft/painting.html#TextRendering){#text-styles:svg-text-rendering
    x-internal="svg-text-rendering"} property.

[`optimizeSpeed`]{#dom-context-2d-textrendering-optimizespeed .dfn dfn-for="CanvasTextRendering" dfn-type="enum-value"}
:   Same as \'optimizeSpeed\' in [SVG
    text-rendering](https://svgwg.org/svg2-draft/painting.html#TextRendering){#text-styles:svg-text-rendering-2
    x-internal="svg-text-rendering"} property.

[`optimizeLegibility`]{#dom-context-2d-textrendering-optimizelegibility .dfn dfn-for="CanvasTextRendering" dfn-type="enum-value"}
:   Same as \'optimizeLegibility\' in [SVG
    text-rendering](https://svgwg.org/svg2-draft/painting.html#TextRendering){#text-styles:svg-text-rendering-3
    x-internal="svg-text-rendering"} property.

[`geometricPrecision`]{#dom-context-2d-textrendering-geometricprecision .dfn dfn-for="CanvasTextRendering" dfn-type="enum-value"}

:   Same as \'geometricPrecision\' in [SVG
    text-rendering](https://svgwg.org/svg2-draft/painting.html#TextRendering){#text-styles:svg-text-rendering-4
    x-internal="svg-text-rendering"} property.

The [text preparation algorithm]{#text-preparation-algorithm .dfn} is as
follows. It takes as input a string `text`{.variable}, a
[`CanvasTextDrawingStyles`{#text-styles:canvastextdrawingstyles-13}](#canvastextdrawingstyles)
object `target`{.variable}, and an optional length
`maxWidth`{.variable}. It returns an array of glyph shapes, each
positioned on a common coordinate space, a
`physical alignment`{.variable} whose value is one of *left*, *right*,
and *center*, and an [inline
box](https://drafts.csswg.org/css2/#inline-box){#text-styles:inline-box
x-internal="inline-box"}. (Most callers of this algorithm ignore the
`physical alignment`{.variable} and the [inline
box](https://drafts.csswg.org/css2/#inline-box){#text-styles:inline-box-2
x-internal="inline-box"}.)

1.  If `maxWidth`{.variable} was provided but is less than or equal to
    zero or equal to NaN, then return an empty array.

2.  Replace all [ASCII
    whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#text-styles:space-characters
    x-internal="space-characters"} in `text`{.variable} with U+0020
    SPACE characters.

3.  Let `font`{.variable} be the current font of `target`{.variable}, as
    given by that object\'s
    [`font`{#text-styles:dom-context-2d-font-3}](#dom-context-2d-font)
    attribute.

4.  Let `language`{.variable} be the `target`{.variable}\'s
    [language](#concept-canvastextdrawingstyles-language){#text-styles:concept-canvastextdrawingstyles-language-4}.

5.  If `language`{.variable} is
    \"[`inherit`{#text-styles:dom-context-2d-lang-inherit-5}](#dom-context-2d-lang-inherit)\":

    1.  Let `sourceObject`{.variable} be `object`{.variable}\'s [font
        style source
        object](#font-style-source-object){#text-styles:font-style-source-object-11}.

    2.  If `sourceObject`{.variable} is a
        [`canvas`{#text-styles:the-canvas-element-11}](#the-canvas-element)
        element, then set `language`{.variable} to the
        `sourceObject`{.variable}\'s
        [language](dom.html#language){#text-styles:language}.

    3.  Otherwise:

        1.  [Assert](https://infra.spec.whatwg.org/#assert){#text-styles:assert-2
            x-internal="assert"}: `sourceObject`{.variable} is an
            [`OffscreenCanvas`{#text-styles:offscreencanvas-3}](#offscreencanvas)
            object.

        2.  Set `language`{.variable} to the
            `sourceObject`{.variable}\'s [inherited
            language](#offscreencanvas-inherited-lang){#text-styles:offscreencanvas-inherited-lang}.

6.  If `language`{.variable} is the empty string, then set
    `language`{.variable} to [explicitly
    unknown](dom.html#concept-explicitly-unknown){#text-styles:concept-explicitly-unknown}.

7.  Apply the appropriate step from the following list to determine the
    value of `direction`{.variable}:

    If the `target`{.variable} object\'s [`direction`{#text-styles:dom-context-2d-direction-4}](#dom-context-2d-direction) attribute has the value \"[`ltr`{#text-styles:dom-context-2d-direction-ltr}](#dom-context-2d-direction-ltr)\"
    :   Let `direction`{.variable} be
        \'[ltr](dom.html#concept-ltr){#text-styles:concept-ltr}\'.

    If the `target`{.variable} object\'s [`direction`{#text-styles:dom-context-2d-direction-5}](#dom-context-2d-direction) attribute has the value \"[`rtl`{#text-styles:dom-context-2d-direction-rtl}](#dom-context-2d-direction-rtl)\"
    :   Let `direction`{.variable} be
        \'[rtl](dom.html#concept-rtl){#text-styles:concept-rtl}\'.

    If the `target`{.variable} object\'s [`direction`{#text-styles:dom-context-2d-direction-6}](#dom-context-2d-direction) attribute has the value \"[`inherit`{#text-styles:dom-context-2d-direction-inherit-3}](#dom-context-2d-direction-inherit)\"

    :   1.  Let `sourceObject`{.variable} be `object`{.variable}\'s
            [font style source
            object](#font-style-source-object){#text-styles:font-style-source-object-12}.

        2.  If `sourceObject`{.variable} is a
            [`canvas`{#text-styles:the-canvas-element-12}](#the-canvas-element)
            element, then let `direction`{.variable} be
            `sourceObject`{.variable}\'s
            [directionality](dom.html#the-directionality){#text-styles:the-directionality}.

        3.  Otherwise:

            1.  [Assert](https://infra.spec.whatwg.org/#assert){#text-styles:assert-3
                x-internal="assert"}: `sourceObject`{.variable} is an
                [`OffscreenCanvas`{#text-styles:offscreencanvas-4}](#offscreencanvas)
                object.

            2.  Let `direction`{.variable} be
                `sourceObject`{.variable}\'s [inherited
                direction](#offscreencanvas-inherited-direction){#text-styles:offscreencanvas-inherited-direction}.

8.  Form a hypothetical infinitely-wide CSS [line
    box](https://drafts.csswg.org/css2/#line-box){#text-styles:line-box
    x-internal="line-box"} containing a single [inline
    box](https://drafts.csswg.org/css2/#inline-box){#text-styles:inline-box-3
    x-internal="inline-box"} containing the text `text`{.variable}, with
    the CSS [content
    language](https://drafts.csswg.org/css-text-4/#content-language){#text-styles:content-language
    x-internal="content-language"} set to `language`{.variable}, and
    with its CSS properties set as follows:

    Property

    Source

    [\'direction\'](https://drafts.csswg.org/css-writing-modes/#direction){#text-styles:'direction'
    x-internal="'direction'"}

    `direction`{.variable}

    [\'font\'](https://drafts.csswg.org/css-fonts/#font-prop){#text-styles:'font'-2
    x-internal="'font'"}

    `font`{.variable}

    [\'font-kerning\'](https://drafts.csswg.org/css-fonts/#propdef-font-kerning){#text-styles:'font-kerning'
    x-internal="'font-kerning'"}

    `target`{.variable}\'s
    [`fontKerning`{#text-styles:dom-context-2d-fontkerning-4}](#dom-context-2d-fontkerning)

    [\'font-stretch\'](https://drafts.csswg.org/css-fonts/#propdef-font-stretch){#text-styles:'font-stretch'-9
    x-internal="'font-stretch'"}

    `target`{.variable}\'s
    [`fontStretch`{#text-styles:dom-context-2d-fontstretch-4}](#dom-context-2d-fontstretch)

    [\'font-variant-caps\'](https://drafts.csswg.org/css-fonts/#propdef-font-variant-caps){#text-styles:'font-variant-caps'-7
    x-internal="'font-variant-caps'"}

    `target`{.variable}\'s
    [`fontVariantCaps`{#text-styles:dom-context-2d-fontvariantcaps-4}](#dom-context-2d-fontvariantcaps)

    [\'letter-spacing\'](https://drafts.csswg.org/css-text/#letter-spacing-property){#text-styles:'letter-spacing'
    x-internal="'letter-spacing'"}

    `target`{.variable}\'s [letter
    spacing](#concept-canvastextdrawingstyles-letter-spacing){#text-styles:concept-canvastextdrawingstyles-letter-spacing-3}

    [SVG
    text-rendering](https://svgwg.org/svg2-draft/painting.html#TextRendering){#text-styles:svg-text-rendering-5
    x-internal="svg-text-rendering"}

    `target`{.variable}\'s
    [`textRendering`{#text-styles:dom-context-2d-textrendering-4}](#dom-context-2d-textrendering)

    [\'white-space\'](https://drafts.csswg.org/css-text/#white-space-property){#text-styles:'white-space'
    x-internal="'white-space'"}

    \'pre\'

    [\'word-spacing\'](https://drafts.csswg.org/css-text/#propdef-word-spacing){#text-styles:'word-spacing'
    x-internal="'word-spacing'"}

    `target`{.variable}\'s [word
    spacing](#concept-canvastextdrawingstyles-word-spacing){#text-styles:concept-canvastextdrawingstyles-word-spacing-3}

    and with all other properties set to their initial values.

9.  If `maxWidth`{.variable} was provided and the hypothetical width of
    the [inline
    box](https://drafts.csswg.org/css2/#inline-box){#text-styles:inline-box-4
    x-internal="inline-box"} in the hypothetical [line
    box](https://drafts.csswg.org/css2/#line-box){#text-styles:line-box-2
    x-internal="line-box"} is greater than `maxWidth`{.variable} [CSS
    pixels](https://drafts.csswg.org/css-values/#px){#text-styles:'px'-2
    x-internal="'px'"}, then change `font`{.variable} to have a more
    condensed font (if one is available or if a reasonably readable one
    can be synthesized by applying a horizontal scale factor to the
    font) or a smaller font, and return to the previous step.

10. The `anchor point`{.variable} is a point on the [inline
    box](https://drafts.csswg.org/css2/#inline-box){#text-styles:inline-box-5
    x-internal="inline-box"}, and the `physical alignment`{.variable} is
    one of the values *left*, *right*, and *center*. These variables are
    determined by the
    [`textAlign`{#text-styles:dom-context-2d-textalign-4}](#dom-context-2d-textalign)
    and
    [`textBaseline`{#text-styles:dom-context-2d-textbaseline-4}](#dom-context-2d-textbaseline)
    values as follows:

    Horizontal position:

    If [`textAlign`{#text-styles:dom-context-2d-textalign-5}](#dom-context-2d-textalign) is [`left`{#text-styles:dom-context-2d-textalign-left}](#dom-context-2d-textalign-left)\
    If [`textAlign`{#text-styles:dom-context-2d-textalign-6}](#dom-context-2d-textalign) is [`start`{#text-styles:dom-context-2d-textalign-start-2}](#dom-context-2d-textalign-start) and `direction`{.variable} is \'ltr\'\
    If [`textAlign`{#text-styles:dom-context-2d-textalign-7}](#dom-context-2d-textalign) is [`end`{#text-styles:dom-context-2d-textalign-end}](#dom-context-2d-textalign-end) and `direction`{.variable} is \'rtl\'
    :   Set the `anchor point`{.variable}\'s horizontal position to the
        left edge of the [inline
        box](https://drafts.csswg.org/css2/#inline-box){#text-styles:inline-box-6
        x-internal="inline-box"}, and let
        `physical alignment`{.variable} be *left*.

    If [`textAlign`{#text-styles:dom-context-2d-textalign-8}](#dom-context-2d-textalign) is [`right`{#text-styles:dom-context-2d-textalign-right}](#dom-context-2d-textalign-right)\
    If [`textAlign`{#text-styles:dom-context-2d-textalign-9}](#dom-context-2d-textalign) is [`end`{#text-styles:dom-context-2d-textalign-end-2}](#dom-context-2d-textalign-end) and `direction`{.variable} is \'ltr\'\
    If [`textAlign`{#text-styles:dom-context-2d-textalign-10}](#dom-context-2d-textalign) is [`start`{#text-styles:dom-context-2d-textalign-start-3}](#dom-context-2d-textalign-start) and `direction`{.variable} is \'rtl\'
    :   Set the `anchor point`{.variable}\'s horizontal position to the
        right edge of the [inline
        box](https://drafts.csswg.org/css2/#inline-box){#text-styles:inline-box-7
        x-internal="inline-box"}, and let
        `physical alignment`{.variable} be *right*.

    If [`textAlign`{#text-styles:dom-context-2d-textalign-11}](#dom-context-2d-textalign) is [`center`{#text-styles:dom-context-2d-textalign-center}](#dom-context-2d-textalign-center)
    :   Set the `anchor point`{.variable}\'s horizontal position to half
        way between the left and right edges of the [inline
        box](https://drafts.csswg.org/css2/#inline-box){#text-styles:inline-box-8
        x-internal="inline-box"}, and let
        `physical alignment`{.variable} be *center*.

    Vertical position:

    If [`textBaseline`{#text-styles:dom-context-2d-textbaseline-5}](#dom-context-2d-textbaseline) is [`top`{#text-styles:dom-context-2d-textbaseline-top}](#dom-context-2d-textbaseline-top)
    :   Set the `anchor point`{.variable}\'s vertical position to the
        top of the em box of the [first available
        font](https://drafts.csswg.org/css-fonts/#first-available-font){#text-styles:first-available-font
        x-internal="first-available-font"} of the [inline
        box](https://drafts.csswg.org/css2/#inline-box){#text-styles:inline-box-9
        x-internal="inline-box"}.

    If [`textBaseline`{#text-styles:dom-context-2d-textbaseline-6}](#dom-context-2d-textbaseline) is [`hanging`{#text-styles:dom-context-2d-textbaseline-hanging}](#dom-context-2d-textbaseline-hanging)
    :   Set the `anchor point`{.variable}\'s vertical position to the
        [hanging
        baseline](https://drafts.csswg.org/css-inline/#hanging-baseline){#text-styles:hanging-baseline-2
        x-internal="hanging-baseline"} of the [first available
        font](https://drafts.csswg.org/css-fonts/#first-available-font){#text-styles:first-available-font-2
        x-internal="first-available-font"} of the [inline
        box](https://drafts.csswg.org/css2/#inline-box){#text-styles:inline-box-10
        x-internal="inline-box"}.

    If [`textBaseline`{#text-styles:dom-context-2d-textbaseline-7}](#dom-context-2d-textbaseline) is [`middle`{#text-styles:dom-context-2d-textbaseline-middle}](#dom-context-2d-textbaseline-middle)
    :   Set the `anchor point`{.variable}\'s vertical position to half
        way between the bottom and the top of the em box of the [first
        available
        font](https://drafts.csswg.org/css-fonts/#first-available-font){#text-styles:first-available-font-3
        x-internal="first-available-font"} of the [inline
        box](https://drafts.csswg.org/css2/#inline-box){#text-styles:inline-box-11
        x-internal="inline-box"}.

    If [`textBaseline`{#text-styles:dom-context-2d-textbaseline-8}](#dom-context-2d-textbaseline) is [`alphabetic`{#text-styles:dom-context-2d-textbaseline-alphabetic-3}](#dom-context-2d-textbaseline-alphabetic)
    :   Set the `anchor point`{.variable}\'s vertical position to the
        [alphabetic
        baseline](https://drafts.csswg.org/css-inline/#alphabetic-baseline){#text-styles:alphabetic-baseline-2
        x-internal="alphabetic-baseline"} of the [first available
        font](https://drafts.csswg.org/css-fonts/#first-available-font){#text-styles:first-available-font-4
        x-internal="first-available-font"} of the [inline
        box](https://drafts.csswg.org/css2/#inline-box){#text-styles:inline-box-12
        x-internal="inline-box"}.

    If [`textBaseline`{#text-styles:dom-context-2d-textbaseline-9}](#dom-context-2d-textbaseline) is [`ideographic`{#text-styles:dom-context-2d-textbaseline-ideographic}](#dom-context-2d-textbaseline-ideographic)
    :   Set the `anchor point`{.variable}\'s vertical position to the
        [ideographic-under
        baseline](https://drafts.csswg.org/css-inline/#ideographic-under-baseline){#text-styles:ideographic-under-baseline-2
        x-internal="ideographic-under-baseline"} of the [first available
        font](https://drafts.csswg.org/css-fonts/#first-available-font){#text-styles:first-available-font-5
        x-internal="first-available-font"} of the [inline
        box](https://drafts.csswg.org/css2/#inline-box){#text-styles:inline-box-13
        x-internal="inline-box"}.

    If [`textBaseline`{#text-styles:dom-context-2d-textbaseline-10}](#dom-context-2d-textbaseline) is [`bottom`{#text-styles:dom-context-2d-textbaseline-bottom}](#dom-context-2d-textbaseline-bottom)
    :   Set the `anchor point`{.variable}\'s vertical position to the
        bottom of the em box of the [first available
        font](https://drafts.csswg.org/css-fonts/#first-available-font){#text-styles:first-available-font-6
        x-internal="first-available-font"} of the [inline
        box](https://drafts.csswg.org/css2/#inline-box){#text-styles:inline-box-14
        x-internal="inline-box"}.

11. Let `result`{.variable} be an array constructed by iterating over
    each glyph in the [inline
    box](https://drafts.csswg.org/css2/#inline-box){#text-styles:inline-box-15
    x-internal="inline-box"} from left to right (if any), adding to the
    array, for each glyph, the shape of the glyph as it is in the
    [inline
    box](https://drafts.csswg.org/css2/#inline-box){#text-styles:inline-box-16
    x-internal="inline-box"}, positioned on a coordinate space using
    [CSS
    pixels](https://drafts.csswg.org/css-values/#px){#text-styles:'px'-3
    x-internal="'px'"} with its origin at the `anchor point`{.variable}.

12. Return `result`{.variable}, `physical alignment`{.variable}, and the
    inline box.

###### [4.12.5.1.6]{.secno} Building paths[](#building-paths){.self-link}

Objects that implement the
[`CanvasPath`{#building-paths:canvaspath}](#canvaspath) interface have a
[path](#concept-path){#building-paths:concept-path}. A
[path]{#concept-path .dfn} has a list of zero or more subpaths. Each
subpath consists of a list of one or more points, connected by straight
or curved [line segments]{#line-segments .dfn}, and a flag indicating
whether the subpath is closed or not. A closed subpath is one where the
last point of the subpath is connected to the first point of the subpath
by a straight line. Subpaths with only one point are ignored when
painting the path.

[Paths](#concept-path){#building-paths:concept-path-2} have a [need new
subpath]{#need-new-subpath .dfn} flag. When this flag is set, certain
APIs create a new subpath rather than extending the previous one. When a
[path](#concept-path){#building-paths:concept-path-3} is created, its
[need new subpath](#need-new-subpath){#building-paths:need-new-subpath}
flag must be set.

When an object implementing the
[`CanvasPath`{#building-paths:canvaspath-2}](#canvaspath) interface is
created, its [path](#concept-path){#building-paths:concept-path-4} must
be initialized to zero subpaths.

`context`{.variable}`.`[`moveTo`](#dom-context-2d-moveto){#dom-context-2d-moveto-dev}`(``x`{.variable}`, ``y`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/moveTo](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/moveTo "The CanvasRenderingContext2D.moveTo() method of the Canvas 2D API begins a new sub-path at the point specified by the given (x, y) coordinates.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari2+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12+]{.opera_android .yes}
:::
::::
:::::

`path`{.variable}`.`[`moveTo`](#dom-context-2d-moveto){#building-paths:dom-context-2d-moveto}`(``x`{.variable}`, ``y`{.variable}`)`

Creates a new subpath with the given point.

`context`{.variable}`.`[`closePath`](#dom-context-2d-closepath){#dom-context-2d-closepath-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/closePath](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/closePath "The CanvasRenderingContext2D.closePath() method of the Canvas 2D API attempts to add a straight line from the current point to the start of the current sub-path. If the shape has already been closed or has only one point, this function does nothing.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari2+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12+]{.opera_android .yes}
:::
::::
:::::

`path`{.variable}`.`[`closePath`](#dom-context-2d-closepath){#building-paths:dom-context-2d-closepath}`()`

Marks the current subpath as closed, and starts a new subpath with a
point the same as the start and end of the newly closed subpath.

`context`{.variable}`.`[`lineTo`](#dom-context-2d-lineto){#dom-context-2d-lineto-dev}`(``x`{.variable}`, ``y`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/lineTo](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/lineTo "The CanvasRenderingContext2D method lineTo(), part of the Canvas 2D API, adds a straight line to the current sub-path by connecting the sub-path's last point to the specified (x, y) coordinates.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari2+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12+]{.opera_android .yes}
:::
::::
:::::

`path`{.variable}`.`[`lineTo`](#dom-context-2d-lineto){#building-paths:dom-context-2d-lineto}`(``x`{.variable}`, ``y`{.variable}`)`

Adds the given point to the current subpath, connected to the previous
one by a straight line.

`context`{.variable}`.`[`quadraticCurveTo`](#dom-context-2d-quadraticcurveto){#dom-context-2d-quadraticcurveto-dev}`(``cpx`{.variable}`, ``cpy`{.variable}`, ``x`{.variable}`, ``y`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/quadraticCurveTo](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/quadraticCurveTo "The CanvasRenderingContext2D.quadraticCurveTo() method of the Canvas 2D API adds a quadratic Bézier curve to the current sub-path. It requires two points: the first one is a control point and the second one is the end point. The starting point is the latest point in the current path, which can be changed using moveTo() before creating the quadratic Bézier curve.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12+]{.opera_android .yes}
:::
::::
:::::

`path`{.variable}`.`[`quadraticCurveTo`](#dom-context-2d-quadraticcurveto){#building-paths:dom-context-2d-quadraticcurveto}`(``cpx`{.variable}`, ``cpy`{.variable}`, ``x`{.variable}`, ``y`{.variable}`)`

Adds the given point to the current subpath, connected to the previous
one by a quadratic Bézier curve with the given control point.

`context`{.variable}`.`[`bezierCurveTo`](#dom-context-2d-beziercurveto){#dom-context-2d-beziercurveto-dev}`(``cp1x`{.variable}`, ``cp1y`{.variable}`, ``cp2x`{.variable}`, ``cp2y`{.variable}`, ``x`{.variable}`, ``y`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/bezierCurveTo](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/bezierCurveTo "The CanvasRenderingContext2D.bezierCurveTo() method of the Canvas 2D API adds a cubic Bézier curve to the current sub-path. It requires three points: the first two are control points and the third one is the end point. The starting point is the latest point in the current path, which can be changed using moveTo() before creating the Bézier curve.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari2+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12+]{.opera_android .yes}
:::
::::
:::::

`path`{.variable}`.`[`bezierCurveTo`](#dom-context-2d-beziercurveto){#building-paths:dom-context-2d-beziercurveto}`(``cp1x`{.variable}`, ``cp1y`{.variable}`, ``cp2x`{.variable}`, ``cp2y`{.variable}`, ``x`{.variable}`, ``y`{.variable}`)`

Adds the given point to the current subpath, connected to the previous
one by a cubic Bézier curve with the given control points.

`context`{.variable}`.`[`arcTo`](#dom-context-2d-arcto){#dom-context-2d-arcto-dev}`(``x1`{.variable}`, ``y1`{.variable}`, ``x2`{.variable}`, ``y2`{.variable}`, ``radius`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/arcTo](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/arcTo "The CanvasRenderingContext2D.arcTo() method of the Canvas 2D API adds a circular arc to the current sub-path, using the given control points and radius. The arc is automatically connected to the path's latest point with a straight line, if necessary for the specified parameters.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari2+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12+]{.opera_android .yes}
:::
::::
:::::

`path`{.variable}`.`[`arcTo`](#dom-context-2d-arcto){#building-paths:dom-context-2d-arcto}`(``x1`{.variable}`, ``y1`{.variable}`, ``x2`{.variable}`, ``y2`{.variable}`, ``radius`{.variable}`)`

Adds an arc with the given control points and radius to the current
subpath, connected to the previous point by a straight line.

Throws an
[\"`IndexSizeError`\"](https://webidl.spec.whatwg.org/#indexsizeerror){#building-paths:indexsizeerror
x-internal="indexsizeerror"}
[`DOMException`{#building-paths:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if the given radius is negative.

<figure class="diagrams">
<img src="/images/arcTo1.png" width="357" height="254" /> <img
src="/images/arcTo2.png" width="468" height="310" /> <img
src="/images/arcTo3.png" width="513" height="233" />
</figure>

`context`{.variable}`.`[`arc`](#dom-context-2d-arc){#dom-context-2d-arc-dev}`(``x`{.variable}`, ``y`{.variable}`, ``radius`{.variable}`, ``startAngle`{.variable}`, ``endAngle`{.variable}` [, ``counterclockwise`{.variable}` ])`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/arc](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/arc "The CanvasRenderingContext2D.arc() method of the Canvas 2D API adds a circular arc to the current sub-path.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12+]{.opera_android .yes}
:::
::::
:::::

`path`{.variable}`.`[`arc`](#dom-context-2d-arc){#building-paths:dom-context-2d-arc}`(``x`{.variable}`, ``y`{.variable}`, ``radius`{.variable}`, ``startAngle`{.variable}`, ``endAngle`{.variable}` [, ``counterclockwise`{.variable}` ])`

Adds points to the subpath such that the arc described by the
circumference of the circle described by the arguments, starting at the
given start angle and ending at the given end angle, going in the given
direction (defaulting to clockwise), is added to the path, connected to
the previous point by a straight line.

Throws an
[\"`IndexSizeError`\"](https://webidl.spec.whatwg.org/#indexsizeerror){#building-paths:indexsizeerror-2
x-internal="indexsizeerror"}
[`DOMException`{#building-paths:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if the given radius is negative.

<figure class="diagrams">
<img src="/images/arc1.png" width="590" height="255" />
</figure>

`context`{.variable}`.`[`ellipse`](#dom-context-2d-ellipse){#dom-context-2d-ellipse-dev}`(``x`{.variable}`, ``y`{.variable}`, ``radiusX`{.variable}`, ``radiusY`{.variable}`, ``rotation`{.variable}`, ``startAngle`{.variable}`, ``endAngle`{.variable}` [, ``counterclockwise`{.variable}`])`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/ellipse](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/ellipse "The CanvasRenderingContext2D.ellipse() method of the Canvas 2D API adds an elliptical arc to the current sub-path.")

Support in all current engines.

::: support
[Firefox48+]{.firefox .yes}[Safari9+]{.safari .yes}[Chrome31+]{.chrome
.yes}

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

`path`{.variable}`.`[`ellipse`](#dom-context-2d-ellipse){#building-paths:dom-context-2d-ellipse}`(``x`{.variable}`, ``y`{.variable}`, ``radiusX`{.variable}`, ``radiusY`{.variable}`, ``rotation`{.variable}`, ``startAngle`{.variable}`, ``endAngle`{.variable}` [, ``counterclockwise`{.variable}`])`

Adds points to the subpath such that the arc described by the
circumference of the ellipse described by the arguments, starting at the
given start angle and ending at the given end angle, going in the given
direction (defaulting to clockwise), is added to the path, connected to
the previous point by a straight line.

Throws an
[\"`IndexSizeError`\"](https://webidl.spec.whatwg.org/#indexsizeerror){#building-paths:indexsizeerror-3
x-internal="indexsizeerror"}
[`DOMException`{#building-paths:domexception-3}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if the given radius is negative.

`context`{.variable}`.`[`rect`](#dom-context-2d-rect){#dom-context-2d-rect-dev}`(``x`{.variable}`, ``y`{.variable}`, ``w`{.variable}`, ``h`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/rect](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/rect "The CanvasRenderingContext2D.rect() method of the Canvas 2D API adds a rectangle to the current path.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari2+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12+]{.opera_android .yes}
:::
::::
:::::

`path`{.variable}`.`[`rect`](#dom-context-2d-rect){#building-paths:dom-context-2d-rect}`(``x`{.variable}`, ``y`{.variable}`, ``w`{.variable}`, ``h`{.variable}`)`

Adds a new closed subpath to the path, representing the given rectangle.

`context`{.variable}`.`[`roundRect`](#dom-context-2d-roundrect){#dom-context-2d-roundrect-dev}`(``x`{.variable}`, ``y`{.variable}`, ``w`{.variable}`, ``h`{.variable}`, ``radii`{.variable}`)`

`path`{.variable}`.`[`roundRect`](#dom-context-2d-roundrect){#building-paths:dom-context-2d-roundrect}`(``x`{.variable}`, ``y`{.variable}`, ``w`{.variable}`, ``h`{.variable}`, ``radii`{.variable}`)`

Adds a new closed subpath to the path representing the given rounded
rectangle. `radii`{.variable} is either a list of radii or a single
radius representing the corners of the rectangle in pixels. If a list is
provided, the number and order of these radii function in the same way
as the CSS
[\'border-radius\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-radius){#building-paths:'border-radius'
x-internal="'border-radius'"} property. A single radius behaves the same
way as a list with a single element.

If `w`{.variable} and `h`{.variable} are both greater than or equal to
0, or if both are smaller than 0, then the path is drawn clockwise.
Otherwise, it is drawn counterclockwise.

When `w`{.variable} is negative, the rounded rectangle is flipped
horizontally, which means that the radius values that normally apply to
the left corners are used on the right and vice versa. Similarly, when
`h`{.variable} is negative, the rounded rect is flipped vertically.

When a value `r`{.variable} in `radii`{.variable} is a number, the
corresponding corner(s) are drawn as circular arcs of radius
`r`{.variable}.

When a value `r`{.variable} in `radii`{.variable} is an object with
`{ x, y }` properties, the corresponding corner(s) are drawn as
elliptical arcs whose `x`{.variable} and `y`{.variable} radii are equal
to `r.x`{.variable} and `r.y`{.variable}, respectively.

When the sum of the `radii`{.variable} of two corners of the same edge
is greater than the length of the edge, all the `radii`{.variable} of
the rounded rectangle are scaled by a factor of length /
(`r1`{.variable} + `r2`{.variable}). If multiple edges have this
property, the scale factor of the edge with the smallest scale factor is
used. This is consistent with CSS behavior.

Throws a
[`RangeError`{#building-paths:js-rangeerror}](https://tc39.es/ecma262/#sec-native-error-types-used-in-this-standard-rangeerror){x-internal="js-rangeerror"}
if `radii`{.variable} is a list whose size is not one, two, three, or
four.

Throws a
[`RangeError`{#building-paths:js-rangeerror-2}](https://tc39.es/ecma262/#sec-native-error-types-used-in-this-standard-rangeerror){x-internal="js-rangeerror"}
if a value in `radii`{.variable} is a negative number, or is an
`{ x, y }` object whose `x` or `y` properties are negative numbers.

The following methods allow authors to manipulate the
[paths](#concept-path){#building-paths:concept-path-5} of objects
implementing the
[`CanvasPath`{#building-paths:canvaspath-3}](#canvaspath) interface.

For objects implementing the
[`CanvasDrawPath`{#building-paths:canvasdrawpath}](#canvasdrawpath) and
[`CanvasTransform`{#building-paths:canvastransform}](#canvastransform)
interfaces, the points passed to the methods, and the resulting lines
added to [current default
path](#current-default-path){#building-paths:current-default-path} by
these methods, must be transformed according to the [current
transformation
matrix](#transformations){#building-paths:transformations} before being
added to the path.

The
[`moveTo(``x`{.variable}`, ``y`{.variable}`)`]{#dom-context-2d-moveto
.dfn dfn-for="CanvasPath" dfn-type="method"} method, when invoked, must
run these steps:

1.  If either of the arguments are infinite or NaN, then return.

2.  Create a new subpath with the specified point as its first (and
    only) point.

When the user agent is to [ensure there is a
subpath]{#ensure-there-is-a-subpath .dfn} for a coordinate
(`x`{.variable}, `y`{.variable}) on a
[path](#concept-path){#building-paths:concept-path-6}, the user agent
must check to see if the
[path](#concept-path){#building-paths:concept-path-7} has its [need new
subpath](#need-new-subpath){#building-paths:need-new-subpath-2} flag
set. If it does, then the user agent must create a new subpath with the
point (`x`{.variable}, `y`{.variable}) as its first (and only) point, as
if the
[`moveTo()`{#building-paths:dom-context-2d-moveto-2}](#dom-context-2d-moveto)
method had been called, and must then unset the
[path](#concept-path){#building-paths:concept-path-8}\'s [need new
subpath](#need-new-subpath){#building-paths:need-new-subpath-3} flag.

The [`closePath()`]{#dom-context-2d-closepath .dfn dfn-for="CanvasPath"
dfn-type="method"} method, when invoked, must do nothing if the
object\'s path has no subpaths. Otherwise, it must mark the last subpath
as closed, create a new subpath whose first point is the same as the
previous subpath\'s first point, and finally add this new subpath to the
path.

If the last subpath had more than one point in its list of points, then
this is equivalent to adding a straight line connecting the last point
back to the first point of the last subpath, thus \"closing\" the
subpath.

------------------------------------------------------------------------

New points and the lines connecting them are added to subpaths using the
methods described below. In all cases, the methods only modify the last
subpath in the object\'s path.

The
[`lineTo(``x`{.variable}`, ``y`{.variable}`)`]{#dom-context-2d-lineto
.dfn dfn-for="CanvasPath" dfn-type="method"} method, when invoked, must
run these steps:

1.  If either of the arguments are infinite or NaN, then return.

2.  If the object\'s path has no subpaths, then [ensure there is a
    subpath](#ensure-there-is-a-subpath){#building-paths:ensure-there-is-a-subpath}
    for (`x`{.variable}, `y`{.variable}).

3.  Otherwise, connect the last point in the subpath to the given point
    (`x`{.variable}, `y`{.variable}) using a straight line, and then add
    the given point (`x`{.variable}, `y`{.variable}) to the subpath.

The
[`quadraticCurveTo(``cpx`{.variable}`, ``cpy`{.variable}`, ``x`{.variable}`, ``y`{.variable}`)`]{#dom-context-2d-quadraticcurveto
.dfn dfn-for="CanvasPath" dfn-type="method"} method, when invoked, must
run these steps:

1.  If any of the arguments are infinite or NaN, then return.

2.  [Ensure there is a
    subpath](#ensure-there-is-a-subpath){#building-paths:ensure-there-is-a-subpath-2}
    for (`cpx`{.variable}, `cpy`{.variable}).

3.  Connect the last point in the subpath to the given point
    (`x`{.variable}, `y`{.variable}) using a quadratic Bézier curve with
    control point (`cpx`{.variable}, `cpy`{.variable}).
    [\[BEZIER\]](references.html#refsBEZIER)

4.  Add the given point (`x`{.variable}, `y`{.variable}) to the subpath.

The
[`bezierCurveTo(``cp1x`{.variable}`, ``cp1y`{.variable}`, ``cp2x`{.variable}`, ``cp2y`{.variable}`, ``x`{.variable}`, ``y`{.variable}`)`]{#dom-context-2d-beziercurveto
.dfn dfn-for="CanvasPath" dfn-type="method"} method, when invoked, must
run these steps:

1.  If any of the arguments are infinite or NaN, then return.

2.  [Ensure there is a
    subpath](#ensure-there-is-a-subpath){#building-paths:ensure-there-is-a-subpath-3}
    for (`cp1x`{.variable}, `cp1y`{.variable}).

3.  Connect the last point in the subpath to the given point
    (`x`{.variable}, `y`{.variable}) using a cubic Bézier curve with
    control points (`cp1x`{.variable}, `cp1y`{.variable}) and
    (`cp2x`{.variable}, `cp2y`{.variable}).
    [\[BEZIER\]](references.html#refsBEZIER)

4.  Add the point (`x`{.variable}, `y`{.variable}) to the subpath.

------------------------------------------------------------------------

The
[`arcTo(``x1`{.variable}`, ``y1`{.variable}`, ``x2`{.variable}`, ``y2`{.variable}`, ``radius`{.variable}`)`]{#dom-context-2d-arcto
.dfn dfn-for="CanvasPath" dfn-type="method"} method, when invoked, must
run these steps:

1.  If any of the arguments are infinite or NaN, then return.

2.  [Ensure there is a
    subpath](#ensure-there-is-a-subpath){#building-paths:ensure-there-is-a-subpath-4}
    for (`x1`{.variable}, `y1`{.variable}).

3.  If `radius`{.variable} is negative, then throw an
    [\"`IndexSizeError`\"](https://webidl.spec.whatwg.org/#indexsizeerror){#building-paths:indexsizeerror-4
    x-internal="indexsizeerror"}
    [`DOMException`{#building-paths:domexception-4}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

4.  Let the point (`x0`{.variable}, `y0`{.variable}) be the last point
    in the subpath, transformed by the inverse of the [current
    transformation
    matrix](#transformations){#building-paths:transformations-2} (so
    that it is in the same coordinate system as the points passed to the
    method).

5.  If the point (`x0`{.variable}, `y0`{.variable}) is equal to the
    point (`x1`{.variable}, `y1`{.variable}), or if the point
    (`x1`{.variable}, `y1`{.variable}) is equal to the point
    (`x2`{.variable}, `y2`{.variable}), or if `radius`{.variable} is
    zero, then add the point (`x1`{.variable}, `y1`{.variable}) to the
    subpath, and connect that point to the previous point
    (`x0`{.variable}, `y0`{.variable}) by a straight line.

6.  Otherwise, if the points (`x0`{.variable}, `y0`{.variable}),
    (`x1`{.variable}, `y1`{.variable}), and (`x2`{.variable},
    `y2`{.variable}) all lie on a single straight line, then add the
    point (`x1`{.variable}, `y1`{.variable}) to the subpath, and connect
    that point to the previous point (`x0`{.variable}, `y0`{.variable})
    by a straight line.

7.  Otherwise, let `The Arc`{.variable} be the shortest arc given by
    circumference of the circle that has radius `radius`{.variable}, and
    that has one point tangent to the half-infinite line that crosses
    the point (`x0`{.variable}, `y0`{.variable}) and ends at the point
    (`x1`{.variable}, `y1`{.variable}), and that has a different point
    tangent to the half-infinite line that ends at the point
    (`x1`{.variable}, `y1`{.variable}) and crosses the point
    (`x2`{.variable}, `y2`{.variable}). The points at which this circle
    touches these two lines are called the start and end tangent points
    respectively. Connect the point (`x0`{.variable}, `y0`{.variable})
    to the start tangent point by a straight line, adding the start
    tangent point to the subpath, and then connect the start tangent
    point to the end tangent point by `The Arc`{.variable}, adding the
    end tangent point to the subpath.

------------------------------------------------------------------------

The
[`arc(``x`{.variable}`, ``y`{.variable}`, ``radius`{.variable}`, ``startAngle`{.variable}`, ``endAngle`{.variable}`, ``counterclockwise`{.variable}`)`]{#dom-context-2d-arc
.dfn dfn-for="CanvasPath" dfn-type="method"} method, when invoked, must
run the [ellipse method
steps](#ellipse-method-steps){#building-paths:ellipse-method-steps} with
[this](https://webidl.spec.whatwg.org/#this){#building-paths:this
x-internal="this"}, `x`{.variable}, `y`{.variable}, `radius`{.variable},
`radius`{.variable}, 0, `startAngle`{.variable}, `endAngle`{.variable},
and `counterclockwise`{.variable}.

This makes it equivalent to
[`ellipse()`{#building-paths:dom-context-2d-ellipse-2}](#dom-context-2d-ellipse)
except that both radii are equal and `rotation`{.variable} is 0.

The
[`ellipse(``x`{.variable}`, ``y`{.variable}`, ``radiusX`{.variable}`, ``radiusY`{.variable}`, ``rotation`{.variable}`, ``startAngle`{.variable}`, ``endAngle`{.variable}`, ``counterclockwise`{.variable}`)`]{#dom-context-2d-ellipse
.dfn dfn-for="CanvasPath" dfn-type="method"} method, when invoked, must
run the [ellipse method
steps](#ellipse-method-steps){#building-paths:ellipse-method-steps-2}
with [this](https://webidl.spec.whatwg.org/#this){#building-paths:this-2
x-internal="this"}, `x`{.variable}, `y`{.variable},
`radiusX`{.variable}, `radiusY`{.variable}, `rotation`{.variable},
`startAngle`{.variable}, `endAngle`{.variable}, and
`counterclockwise`{.variable}.

The [determine the point on an ellipse
steps]{#determine-the-point-on-an-ellipse-steps .dfn}, given
`ellipse`{.variable}, and `angle`{.variable}, are:

1.  Let `eccentricCircle`{.variable} be the circle that shares its
    origin with `ellipse`{.variable}, with a radius equal to the
    semi-major axis of `ellipse`{.variable}.

2.  Let `outerPoint`{.variable} be the point on
    `eccentricCircle`{.variable}\'s circumference at `angle`{.variable}
    measured in radians clockwise from `ellipse`{.variable}\'s
    semi-major axis.

3.  Let `chord`{.variable} be the line perpendicular to
    `ellipse`{.variable}\'s major axis between this axis and
    `outerPoint`{.variable}.

4.  Return the point on `chord`{.variable} that crosses
    `ellipse`{.variable}\'s circumference.

The [ellipse method steps]{#ellipse-method-steps .dfn}, given
`canvasPath`{.variable}, `x`{.variable}, `y`{.variable},
`radiusX`{.variable}, `radiusY`{.variable}, `rotation`{.variable},
`startAngle`{.variable}, `endAngle`{.variable}, and
`counterclockwise`{.variable}, are:

1.  If any of the arguments are infinite or NaN, then return.

2.  If either `radiusX`{.variable} or `radiusY`{.variable} are negative,
    then throw an
    [\"`IndexSizeError`\"](https://webidl.spec.whatwg.org/#indexsizeerror){#building-paths:indexsizeerror-5
    x-internal="indexsizeerror"}
    [`DOMException`{#building-paths:domexception-5}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

3.  If `canvasPath`{.variable}\'s path has any subpaths, then add a
    straight line from the last point in the subpath to the start point
    of the arc.

4.  Add the start and end points of the arc to the subpath, and connect
    them with an arc. The arc and its start and end points are defined
    as follows:

    Consider an ellipse that has its origin at (`x`{.variable},
    `y`{.variable}), that has a major-axis radius `radiusX`{.variable}
    and a minor-axis radius `radiusY`{.variable}, and that is rotated
    about its origin such that its semi-major axis is inclined
    `rotation`{.variable} radians clockwise from the x-axis.

    If `counterclockwise`{.variable} is false and `endAngle`{.variable}
    − `startAngle`{.variable} is greater than or equal to 2π, or, if
    `counterclockwise`{.variable} is *true* and `startAngle`{.variable}
    − `endAngle`{.variable} is greater than or equal to 2π, then the arc
    is the whole circumference of this ellipse, and both the start point
    and the end point are the result of running the [determine the point
    on an ellipse
    steps](#determine-the-point-on-an-ellipse-steps){#building-paths:determine-the-point-on-an-ellipse-steps}
    given this ellipse and `startAngle`{.variable}.

    Otherwise, the start point is the result of running the [determine
    the point on an ellipse
    steps](#determine-the-point-on-an-ellipse-steps){#building-paths:determine-the-point-on-an-ellipse-steps-2}
    given this ellipse and `startAngle`{.variable}, the end point is the
    result of running the [determine the point on an ellipse
    steps](#determine-the-point-on-an-ellipse-steps){#building-paths:determine-the-point-on-an-ellipse-steps-3}
    given this ellipse and `endAngle`{.variable}, and the arc is the
    path along the circumference of this ellipse from the start point to
    the end point, going counterclockwise if
    `counterclockwise`{.variable} is true, and clockwise otherwise.
    Since the points are on the ellipse, as opposed to being simply
    angles from zero, the arc can never cover an angle greater than 2π
    radians.

    Even if the arc covers the entire circumference of the ellipse and
    there are no other points in the subpath, the path is not closed
    unless the
    [`closePath()`{#building-paths:dom-context-2d-closepath-2}](#dom-context-2d-closepath)
    method is appropriately invoked.

------------------------------------------------------------------------

The
[`rect(``x`{.variable}`, ``y`{.variable}`, ``w`{.variable}`, ``h`{.variable}`)`]{#dom-context-2d-rect
.dfn dfn-for="CanvasPath" dfn-type="method"} method, when invoked, must
run these steps:

1.  If any of the arguments are infinite or NaN, then return.

2.  Create a new subpath containing just the four points
    (`x`{.variable}, `y`{.variable}), (`x`{.variable}+`w`{.variable},
    `y`{.variable}), (`x`{.variable}+`w`{.variable},
    `y`{.variable}+`h`{.variable}), (`x`{.variable},
    `y`{.variable}+`h`{.variable}), in that order, with those four
    points connected by straight lines.

3.  Mark the subpath as closed.

4.  Create a new subpath with the point (`x`{.variable}, `y`{.variable})
    as the only point in the subpath.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/roundRect](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/roundRect "The CanvasRenderingContext2D.roundRect() method of the Canvas 2D API adds a rounded rectangle to the current path.")

Support in all current engines.

::: support
[Firefox112+]{.firefox .yes}[Safari16+]{.safari .yes}[Chrome99+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge99+]{.edge_blink .yes}

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
[`roundRect(``x`{.variable}`, ``y`{.variable}`, ``w`{.variable}`, ``h`{.variable}`, ``radii`{.variable}`)`]{#dom-context-2d-roundrect
.dfn dfn-for="CanvasPath" dfn-type="method"} method steps are:

1.  If any of `x`{.variable}, `y`{.variable}, `w`{.variable}, or
    `h`{.variable} are infinite or NaN, then return.

2.  If `radii`{.variable} is an
    [`unrestricted double`{#building-paths:idl-unrestricted-double}](https://webidl.spec.whatwg.org/#idl-unrestricted-double){x-internal="idl-unrestricted-double"}
    or
    [`DOMPointInit`{#building-paths:dompointinit}](https://drafts.fxtf.org/geometry/#dictdef-dompointinit){x-internal="dompointinit"},
    then set `radii`{.variable} to « `radii`{.variable} ».

3.  If `radii`{.variable} is not a list of size one, two, three, or
    four, then throw a
    [`RangeError`{#building-paths:js-rangeerror-3}](https://tc39.es/ecma262/#sec-native-error-types-used-in-this-standard-rangeerror){x-internal="js-rangeerror"}.

4.  Let `normalizedRadii`{.variable} be an empty list.

5.  For each `radius`{.variable} of `radii`{.variable}:

    1.  If `radius`{.variable} is a
        [`DOMPointInit`{#building-paths:dompointinit-2}](https://drafts.fxtf.org/geometry/#dictdef-dompointinit){x-internal="dompointinit"}:

        1.  If
            `radius`{.variable}\[\"[`x`{#building-paths:dompointinit-x}](https://drafts.fxtf.org/geometry/#dom-dompointinit-x){x-internal="dompointinit-x"}\"\]
            or
            `radius`{.variable}\[\"[`y`{#building-paths:dompointinit-y}](https://drafts.fxtf.org/geometry/#dom-dompointinit-y){x-internal="dompointinit-y"}\"\]
            is infinite or NaN, then return.

        2.  If
            `radius`{.variable}\[\"[`x`{#building-paths:dompointinit-x-2}](https://drafts.fxtf.org/geometry/#dom-dompointinit-x){x-internal="dompointinit-x"}\"\]
            or
            `radius`{.variable}\[\"[`y`{#building-paths:dompointinit-y-2}](https://drafts.fxtf.org/geometry/#dom-dompointinit-y){x-internal="dompointinit-y"}\"\]
            is negative, then throw a
            [`RangeError`{#building-paths:js-rangeerror-4}](https://tc39.es/ecma262/#sec-native-error-types-used-in-this-standard-rangeerror){x-internal="js-rangeerror"}.

        3.  Otherwise, append `radius`{.variable} to
            `normalizedRadii`{.variable}.

    2.  If `radius`{.variable} is a
        [`unrestricted double`{#building-paths:idl-unrestricted-double-2}](https://webidl.spec.whatwg.org/#idl-unrestricted-double){x-internal="idl-unrestricted-double"}:

        1.  If `radius`{.variable} is infinite or NaN, then return.

        2.  If `radius`{.variable} is negative, then throw a
            [`RangeError`{#building-paths:js-rangeerror-5}](https://tc39.es/ecma262/#sec-native-error-types-used-in-this-standard-rangeerror){x-internal="js-rangeerror"}.

        3.  Otherwise, append «\[
            \"[`x`{#building-paths:dompointinit-x-3}](https://drafts.fxtf.org/geometry/#dom-dompointinit-x){x-internal="dompointinit-x"}\"
            → `radius`{.variable},
            \"[`y`{#building-paths:dompointinit-x-4}](https://drafts.fxtf.org/geometry/#dom-dompointinit-x){x-internal="dompointinit-x"}\"
            → `radius`{.variable} \]» to `normalizedRadii`{.variable}.

6.  Let `upperLeft`{.variable}, `upperRight`{.variable},
    `lowerRight`{.variable}, and `lowerLeft`{.variable} be null.

7.  If `normalizedRadii`{.variable}\'s size is 4, then set
    `upperLeft`{.variable} to `normalizedRadii`{.variable}\[0\], set
    `upperRight`{.variable} to `normalizedRadii`{.variable}\[1\], set
    `lowerRight`{.variable} to `normalizedRadii`{.variable}\[2\], and
    set `lowerLeft`{.variable} to `normalizedRadii`{.variable}\[3\].

8.  If `normalizedRadii`{.variable}\'s size is 3, then set
    `upperLeft`{.variable} to `normalizedRadii`{.variable}\[0\], set
    `upperRight`{.variable} and `lowerLeft`{.variable} to
    `normalizedRadii`{.variable}\[1\], and set `lowerRight`{.variable}
    to `normalizedRadii`{.variable}\[2\].

9.  If `normalizedRadii`{.variable}\'s size is 2, then set
    `upperLeft`{.variable} and `lowerRight`{.variable} to
    `normalizedRadii`{.variable}\[0\] and set `upperRight`{.variable}
    and `lowerLeft`{.variable} to `normalizedRadii`{.variable}\[1\].

10. If `normalizedRadii`{.variable}\'s size is 1, then set
    `upperLeft`{.variable}, `upperRight`{.variable},
    `lowerRight`{.variable}, and `lowerLeft`{.variable} to
    `normalizedRadii`{.variable}\[0\].

11. Corner curves must not overlap. Scale all radii to prevent this:

    1.  Let `top`{.variable} be
        `upperLeft`{.variable}\[\"[`x`{#building-paths:dompointinit-x-5}](https://drafts.fxtf.org/geometry/#dom-dompointinit-x){x-internal="dompointinit-x"}\"\] +
        `upperRight`{.variable}\[\"[`x`{#building-paths:dompointinit-x-6}](https://drafts.fxtf.org/geometry/#dom-dompointinit-x){x-internal="dompointinit-x"}\"\].

    2.  Let `right`{.variable} be
        `upperRight`{.variable}\[\"[`y`{#building-paths:dompointinit-y-3}](https://drafts.fxtf.org/geometry/#dom-dompointinit-y){x-internal="dompointinit-y"}\"\] +
        `lowerRight`{.variable}\[\"[`y`{#building-paths:dompointinit-y-4}](https://drafts.fxtf.org/geometry/#dom-dompointinit-y){x-internal="dompointinit-y"}\"\].

    3.  Let `bottom`{.variable} be
        `lowerRight`{.variable}\[\"[`x`{#building-paths:dompointinit-x-7}](https://drafts.fxtf.org/geometry/#dom-dompointinit-x){x-internal="dompointinit-x"}\"\] +
        `lowerLeft`{.variable}\[\"[`x`{#building-paths:dompointinit-x-8}](https://drafts.fxtf.org/geometry/#dom-dompointinit-x){x-internal="dompointinit-x"}\"\].

    4.  Let `left`{.variable} be
        `upperLeft`{.variable}\[\"[`y`{#building-paths:dompointinit-y-5}](https://drafts.fxtf.org/geometry/#dom-dompointinit-y){x-internal="dompointinit-y"}\"\] +
        `lowerLeft`{.variable}\[\"[`y`{#building-paths:dompointinit-y-6}](https://drafts.fxtf.org/geometry/#dom-dompointinit-y){x-internal="dompointinit-y"}\"\].

    5.  Let `scale`{.variable} be the minimum value of the ratios
        `w`{.variable} / `top`{.variable}, `h`{.variable} /
        `right`{.variable}, `w`{.variable} / `bottom`{.variable},
        `h`{.variable} / `left`{.variable}.

    6.  If `scale`{.variable} is less than 1, then set the
        [`x`{#building-paths:dompointinit-x-9}](https://drafts.fxtf.org/geometry/#dom-dompointinit-x){x-internal="dompointinit-x"}
        and
        [`y`{#building-paths:dompointinit-y-7}](https://drafts.fxtf.org/geometry/#dom-dompointinit-y){x-internal="dompointinit-y"}
        members of `upperLeft`{.variable}, `upperRight`{.variable},
        `lowerLeft`{.variable}, and `lowerRight`{.variable} to their
        current values multiplied by `scale`{.variable}.

12. Create a new subpath:

    1.  Move to the point (`x`{.variable} +
        `upperLeft`{.variable}\[\"[`x`{#building-paths:dompointinit-x-10}](https://drafts.fxtf.org/geometry/#dom-dompointinit-x){x-internal="dompointinit-x"}\"\],
        `y`{.variable}).

    2.  Draw a straight line to the point (`x`{.variable} +
        `w`{.variable} −
        `upperRight`{.variable}\[\"[`x`{#building-paths:dompointinit-x-11}](https://drafts.fxtf.org/geometry/#dom-dompointinit-x){x-internal="dompointinit-x"}\"\],
        `y`{.variable}).

    3.  Draw an arc to the point (`x`{.variable} + `w`{.variable},
        `y`{.variable} +
        `upperRight`{.variable}\[\"[`y`{#building-paths:dompointinit-y-8}](https://drafts.fxtf.org/geometry/#dom-dompointinit-y){x-internal="dompointinit-y"}\"\]).

    4.  Draw a straight line to the point (`x`{.variable} +
        `w`{.variable}, `y`{.variable} + `h`{.variable} −
        `lowerRight`{.variable}\[\"[`y`{#building-paths:dompointinit-y-9}](https://drafts.fxtf.org/geometry/#dom-dompointinit-y){x-internal="dompointinit-y"}\"\]).

    5.  Draw an arc to the point (`x`{.variable} + `w`{.variable} −
        `lowerRight`{.variable}\[\"[`x`{#building-paths:dompointinit-x-12}](https://drafts.fxtf.org/geometry/#dom-dompointinit-x){x-internal="dompointinit-x"}\"\],
        `y`{.variable} + `h`{.variable}).

    6.  Draw a straight line to the point (`x`{.variable} +
        `lowerLeft`{.variable}\[\"[`x`{#building-paths:dompointinit-x-13}](https://drafts.fxtf.org/geometry/#dom-dompointinit-x){x-internal="dompointinit-x"}\"\],
        `y`{.variable} + `h`{.variable}).

    7.  Draw an arc to the point (`x`{.variable}, `y`{.variable} +
        `h`{.variable} −
        `lowerLeft`{.variable}\[\"[`y`{#building-paths:dompointinit-y-10}](https://drafts.fxtf.org/geometry/#dom-dompointinit-y){x-internal="dompointinit-y"}\"\]).

    8.  Draw a straight line to the point (`x`{.variable},
        `y`{.variable} +
        `upperLeft`{.variable}\[\"[`y`{#building-paths:dompointinit-y-11}](https://drafts.fxtf.org/geometry/#dom-dompointinit-y){x-internal="dompointinit-y"}\"\]).

    9.  Draw an arc to the point (`x`{.variable} +
        `upperLeft`{.variable}\[\"[`x`{#building-paths:dompointinit-x-14}](https://drafts.fxtf.org/geometry/#dom-dompointinit-x){x-internal="dompointinit-x"}\"\],
        `y`{.variable}).

13. Mark the subpath as closed.

14. Create a new subpath with the point (`x`{.variable}, `y`{.variable})
    as the only point in the subpath.

This is designed to behave similarly to the CSS
[\'border-radius\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-radius){#building-paths:'border-radius'-2
x-internal="'border-radius'"} property.

###### [4.12.5.1.7]{.secno} [`Path2D`{#path2d-objects:path2d}](#path2d) objects[](#path2d-objects){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Path2D](https://developer.mozilla.org/en-US/docs/Web/API/Path2D "The Path2D interface of the Canvas 2D API is used to declare a path that can then be used on a CanvasRenderingContext2D object. The path methods of the CanvasRenderingContext2D interface are also present on this interface, which gives you the convenience of being able to retain and replay your path whenever desired.")

Support in all current engines.

::: support
[Firefox31+]{.firefox .yes}[Safari8+]{.safari .yes}[Chrome36+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)14+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

[`Path2D`{#path2d-objects:path2d-2}](#path2d) objects can be used to
declare paths that are then later used on objects implementing the
[`CanvasDrawPath`{#path2d-objects:canvasdrawpath}](#canvasdrawpath)
interface. In addition to many of the APIs described in earlier
sections, [`Path2D`{#path2d-objects:path2d-3}](#path2d) objects have
methods to combine paths, and to add text to paths.

`path`{.variable}` = new `[`Path2D`](#dom-path2d){#dom-path2d-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Path2D/Path2D](https://developer.mozilla.org/en-US/docs/Web/API/Path2D/Path2D "The Path2D() constructor returns a newly instantiated Path2D object, optionally with another path as an argument (creates a copy), or optionally with a string consisting of SVG path data.")

Support in all current engines.

::: support
[Firefox31+]{.firefox .yes}[Safari8+]{.safari .yes}[Chrome36+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)14+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

Creates a new empty [`Path2D`{#path2d-objects:path2d-4}](#path2d)
object.

`path`{.variable}` = new `[`Path2D`](#dom-path2d){#path2d-objects:dom-path2d}`(``path`{.variable}`)`

When `path`{.variable} is a
[`Path2D`{#path2d-objects:path2d-5}](#path2d) object, returns a copy.

When `path`{.variable} is a string, creates the path described by the
argument, interpreted as SVG path data.
[\[SVG\]](references.html#refsSVG)

`path`{.variable}`.`[`addPath`](#dom-path2d-addpath){#dom-path2d-addpath-dev}`(``path`{.variable}` [, ``transform`{.variable}` ])`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Path2D/addPath](https://developer.mozilla.org/en-US/docs/Web/API/Path2D/addPath "The Path2D.addPath() method of the Canvas 2D API adds one Path2D object to another Path2D object.")

Support in all current engines.

::: support
[Firefox34+]{.firefox .yes}[Safari9+]{.safari .yes}[Chrome68+]{.chrome
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

Adds to the path the path given by the argument.

The [`Path2D(``path`{.variable}`)`]{#dom-path2d .dfn dfn-for="Path2D"
dfn-type="constructor"} constructor, when invoked, must run these steps:

1.  Let `output`{.variable} be a new
    [`Path2D`{#path2d-objects:path2d-6}](#path2d) object.

2.  If `path`{.variable} is not given, then return `output`{.variable}.

3.  If `path`{.variable} is a
    [`Path2D`{#path2d-objects:path2d-7}](#path2d) object, then add all
    subpaths of `path`{.variable} to `output`{.variable} and return
    `output`{.variable}. (In other words, it returns a copy of the
    argument.)

4.  Let `svgPath`{.variable} be the result of parsing and interpreting
    `path`{.variable} according to SVG 2\'s rules for path data.
    [\[SVG\]](references.html#refsSVG)

    The resulting path could be empty. SVG defines error handling rules
    for parsing and applying path data.

5.  Let (`x`{.variable}, `y`{.variable}) be the last point in
    `svgPath`{.variable}.

6.  Add all the subpaths, if any, from `svgPath`{.variable} to
    `output`{.variable}.

7.  Create a new subpath in `output`{.variable} with (`x`{.variable},
    `y`{.variable}) as the only point in the subpath.

8.  Return `output`{.variable}.

------------------------------------------------------------------------

The
[`addPath(``path`{.variable}`, ``transform`{.variable}`)`]{#dom-path2d-addpath
.dfn dfn-for="Path2D" dfn-type="method"} method, when invoked on a
[`Path2D`{#path2d-objects:path2d-8}](#path2d) object `a`{.variable},
must run these steps:

1.  If the [`Path2D`{#path2d-objects:path2d-9}](#path2d) object
    `path`{.variable} has no subpaths, then return.

2.  Let `matrix`{.variable} be the result of [creating a `DOMMatrix`
    from the 2D
    dictionary](https://drafts.fxtf.org/geometry/#create-a-dommatrix-from-the-2d-dictionary){#path2d-objects:create-a-dommatrix-from-a-2d-dictionary
    x-internal="create-a-dommatrix-from-a-2d-dictionary"}
    `transform`{.variable}.

3.  If one or more of `matrix`{.variable}\'s [m11
    element](https://drafts.fxtf.org/geometry/#matrix-m11-element){#path2d-objects:m11-element
    x-internal="m11-element"}, [m12
    element](https://drafts.fxtf.org/geometry/#matrix-m12-element){#path2d-objects:m12-element
    x-internal="m12-element"}, [m21
    element](https://drafts.fxtf.org/geometry/#matrix-m21-element){#path2d-objects:m21-element
    x-internal="m21-element"}, [m22
    element](https://drafts.fxtf.org/geometry/#matrix-m22-element){#path2d-objects:m22-element
    x-internal="m22-element"}, [m41
    element](https://drafts.fxtf.org/geometry/#matrix-m41-element){#path2d-objects:m41-element
    x-internal="m41-element"}, or [m42
    element](https://drafts.fxtf.org/geometry/#matrix-m42-element){#path2d-objects:m42-element
    x-internal="m42-element"} are infinite or NaN, then return.

4.  Create a copy of all the subpaths in `path`{.variable}. Let this
    copy be known as `c`{.variable}.

5.  Transform all the coordinates and lines in `c`{.variable} by the
    transform matrix `matrix`{.variable}.

6.  Let (`x`{.variable}, `y`{.variable}) be the last point in the last
    subpath of `c`{.variable}.

7.  Add all the subpaths in `c`{.variable} to `a`{.variable}.

8.  Create a new subpath in `a`{.variable} with (`x`{.variable},
    `y`{.variable}) as the only point in the subpath.

###### [4.12.5.1.8]{.secno} [Transformations]{.dfn}[](#transformations){.self-link}

Objects that implement the
[`CanvasTransform`{#transformations:canvastransform}](#canvastransform)
interface have a [current transformation
matrix]{#current-transformation-matrix .dfn}, as well as methods
(described in this section) to manipulate it. When an object
implementing the
[`CanvasTransform`{#transformations:canvastransform-2}](#canvastransform)
interface is created, its transformation matrix must be initialized to
the identity matrix.

The [current transformation
matrix](#current-transformation-matrix){#transformations:current-transformation-matrix}
is applied to coordinates when creating the [current default
path](#current-default-path){#transformations:current-default-path}, and
when painting text, shapes, and
[`Path2D`{#transformations:path2d}](#path2d) objects, on objects
implementing the
[`CanvasTransform`{#transformations:canvastransform-3}](#canvastransform)
interface.

The transformations must be performed in reverse order.

For instance, if a scale transformation that doubles the width is
applied to the canvas, followed by a rotation transformation that
rotates drawing operations by a quarter turn, and a rectangle twice as
wide as it is tall is then drawn on the canvas, the actual result will
be a square.

`context`{.variable}`.`[`scale`](#dom-context-2d-scale){#dom-context-2d-scale-dev}`(``x`{.variable}`, ``y`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/scale](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/scale "The CanvasRenderingContext2D.scale() method of the Canvas 2D API adds a scaling transformation to the canvas units horizontally and/or vertically.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari2+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Changes the [current transformation
matrix](#current-transformation-matrix){#transformations:current-transformation-matrix-2}
to apply a scaling transformation with the given characteristics.

`context`{.variable}`.`[`rotate`](#dom-context-2d-rotate){#dom-context-2d-rotate-dev}`(``angle`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/rotate](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/rotate "The CanvasRenderingContext2D.rotate() method of the Canvas 2D API adds a rotation to the transformation matrix.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari2+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Changes the [current transformation
matrix](#current-transformation-matrix){#transformations:current-transformation-matrix-3}
to apply a rotation transformation with the given characteristics. The
angle is in radians.

`context`{.variable}`.`[`translate`](#dom-context-2d-translate){#dom-context-2d-translate-dev}`(``x`{.variable}`, ``y`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/translate](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/translate "The CanvasRenderingContext2D.translate() method of the Canvas 2D API adds a translation transformation to the current matrix.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari2+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Changes the [current transformation
matrix](#current-transformation-matrix){#transformations:current-transformation-matrix-4}
to apply a translation transformation with the given characteristics.

`context`{.variable}`.`[`transform`](#dom-context-2d-transform){#dom-context-2d-transform-dev}`(``a`{.variable}`, ``b`{.variable}`, ``c`{.variable}`, ``d`{.variable}`, ``e`{.variable}`, ``f`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/transform](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/transform "The CanvasRenderingContext2D.transform() method of the Canvas 2D API multiplies the current transformation with the matrix described by the arguments of this method. This lets you scale, rotate, translate (move), and skew the context.")

Support in all current engines.

::: support
[Firefox3+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Changes the [current transformation
matrix](#current-transformation-matrix){#transformations:current-transformation-matrix-5}
to apply the matrix given by the arguments as described below.

`matrix`{.variable}` = ``context`{.variable}`.`[`getTransform`](#dom-context-2d-gettransform){#dom-context-2d-gettransform-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/getTransform](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/getTransform "The CanvasRenderingContext2D.getTransform() method of the Canvas 2D API retrieves the current transformation matrix being applied to the context.")

Support in all current engines.

::: support
[Firefox70+]{.firefox .yes}[Safari11.1+]{.safari
.yes}[Chrome68+]{.chrome .yes}

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

Returns a copy of the [current transformation
matrix](#current-transformation-matrix){#transformations:current-transformation-matrix-6},
as a newly created
[`DOMMatrix`{#transformations:dommatrix}](https://drafts.fxtf.org/geometry/#dommatrix){x-internal="dommatrix"}
object.

`context`{.variable}`.`[`setTransform`](#dom-context-2d-settransform){#dom-context-2d-settransform-dev}`(``a`{.variable}`, ``b`{.variable}`, ``c`{.variable}`, ``d`{.variable}`, ``e`{.variable}`, ``f`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/setTransform](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/setTransform "The CanvasRenderingContext2D.setTransform() method of the Canvas 2D API resets (overrides) the current transformation to the identity matrix, and then invokes a transformation described by the arguments of this method. This lets you scale, rotate, translate (move), and skew the context.")

Support in all current engines.

::: support
[Firefox3+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome2+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Changes the [current transformation
matrix](#current-transformation-matrix){#transformations:current-transformation-matrix-7}
*to* the matrix given by the arguments as described below.

`context`{.variable}`.`[`setTransform`](#dom-context-2d-settransform-matrix){#dom-context-2d-settransform-matrix-dev}`(``transform`{.variable}`)`

Changes the [current transformation
matrix](#current-transformation-matrix){#transformations:current-transformation-matrix-8}
*to* the matrix represented by the passed
[`DOMMatrix2DInit`{#transformations:dommatrix2dinit}](https://drafts.fxtf.org/geometry/#dictdef-dommatrix2dinit){x-internal="dommatrix2dinit"}
dictionary.

`context`{.variable}`.`[`resetTransform`](#dom-context-2d-resettransform){#dom-context-2d-resettransform-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/resetTransform](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/resetTransform "The CanvasRenderingContext2D.resetTransform() method of the Canvas 2D API resets the current transform to the identity matrix.")

Support in all current engines.

::: support
[Firefox36+]{.firefox .yes}[Safari10.1+]{.safari
.yes}[Chrome31+]{.chrome .yes}

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

Changes the [current transformation
matrix](#current-transformation-matrix){#transformations:current-transformation-matrix-9}
to the identity matrix.

The [`scale(``x`{.variable}`, ``y`{.variable}`)`]{#dom-context-2d-scale
.dfn dfn-for="CanvasTransform" dfn-type="method"} method, when invoked,
must run these steps:

1.  If either of the arguments are infinite or NaN, then return.

2.  Add the scaling transformation described by the arguments to the
    [current transformation
    matrix](#current-transformation-matrix){#transformations:current-transformation-matrix-10}.
    The `x`{.variable} argument represents the scale factor in the
    horizontal direction and the `y`{.variable} argument represents the
    scale factor in the vertical direction. The factors are multiples.

The [`rotate(``angle`{.variable}`)`]{#dom-context-2d-rotate .dfn
dfn-for="CanvasTransform" dfn-type="method"} method, when invoked, must
run these steps:

1.  If `angle`{.variable} is infinite or NaN, then return.

2.  Add the rotation transformation described by the argument to the
    [current transformation
    matrix](#current-transformation-matrix){#transformations:current-transformation-matrix-11}.
    The `angle`{.variable} argument represents a clockwise rotation
    angle expressed in radians.

The
[`translate(``x`{.variable}`, ``y`{.variable}`)`]{#dom-context-2d-translate
.dfn dfn-for="CanvasTransform" dfn-type="method"} method, when invoked,
must run these steps:

1.  If either of the arguments are infinite or NaN, then return.

2.  Add the translation transformation described by the arguments to the
    [current transformation
    matrix](#current-transformation-matrix){#transformations:current-transformation-matrix-12}.
    The `x`{.variable} argument represents the translation distance in
    the horizontal direction and the `y`{.variable} argument represents
    the translation distance in the vertical direction. The arguments
    are in coordinate space units.

The
[`transform(``a`{.variable}`, ``b`{.variable}`, ``c`{.variable}`, ``d`{.variable}`, ``e`{.variable}`, ``f`{.variable}`)`]{#dom-context-2d-transform
.dfn dfn-for="CanvasTransform" dfn-type="method"} method, when invoked,
must run these steps:

1.  If any of the arguments are infinite or NaN, then return.

2.  Replace the [current transformation
    matrix](#current-transformation-matrix){#transformations:current-transformation-matrix-13}
    with the result of
    [multiplying](https://drafts.fxtf.org/geometry/#matrix-multiply){#transformations:matrix-multiplication
    x-internal="matrix-multiplication"} the [current transformation
    matrix](#current-transformation-matrix){#transformations:current-transformation-matrix-14}
    with the matrix described by:

    `a`{.variable}

    `c`{.variable}

    `e`{.variable}

    `b`{.variable}

    `d`{.variable}

    `f`{.variable}

    0

    0

    1

The arguments `a`{.variable}, `b`{.variable}, `c`{.variable},
`d`{.variable}, `e`{.variable}, and `f`{.variable} are sometimes called
`m11`{.variable}, `m12`{.variable}, `m21`{.variable}, `m22`{.variable},
`dx`{.variable}, and `dy`{.variable} or `m11`{.variable},
`m21`{.variable}, `m12`{.variable}, `m22`{.variable}, `dx`{.variable},
and `dy`{.variable}. Care ought to be taken in particular with the order
of the second and third arguments (`b`{.variable} and `c`{.variable}) as
their order varies from API to API and APIs sometimes use the notation
`m12`{.variable}/`m21`{.variable} and sometimes
`m21`{.variable}/`m12`{.variable} for those positions.

The [`getTransform()`]{#dom-context-2d-gettransform .dfn
dfn-for="CanvasTransform" dfn-type="method"} method, when invoked, must
return a newly created
[`DOMMatrix`{#transformations:dommatrix-2}](https://drafts.fxtf.org/geometry/#dommatrix){x-internal="dommatrix"}
representing a copy of the [current transformation
matrix](#current-transformation-matrix){#transformations:current-transformation-matrix-15}
matrix of the context.

This returned object is not live, so updating it will not affect the
[current transformation
matrix](#current-transformation-matrix){#transformations:current-transformation-matrix-16},
and updating the [current transformation
matrix](#current-transformation-matrix){#transformations:current-transformation-matrix-17}
will not affect an already returned
[`DOMMatrix`{#transformations:dommatrix-3}](https://drafts.fxtf.org/geometry/#dommatrix){x-internal="dommatrix"}.

The
[`setTransform(``a`{.variable}`, ``b`{.variable}`, ``c`{.variable}`, ``d`{.variable}`, ``e`{.variable}`, ``f`{.variable}`)`]{#dom-context-2d-settransform
.dfn dfn-for="CanvasTransform" dfn-type="method"} method, when invoked,
must run these steps:

1.  If any of the arguments are infinite or NaN, then return.

2.  Reset the [current transformation
    matrix](#current-transformation-matrix){#transformations:current-transformation-matrix-18}
    to the matrix described by:

    `a`{.variable}

    `c`{.variable}

    `e`{.variable}

    `b`{.variable}

    `d`{.variable}

    `f`{.variable}

    0

    0

    1

The
[`setTransform(``transform`{.variable}`)`]{#dom-context-2d-settransform-matrix
.dfn dfn-for="CanvasTransform" dfn-type="method"} method, when invoked,
must run these steps:

1.  Let `matrix`{.variable} be the result of [creating a `DOMMatrix`
    from the 2D
    dictionary](https://drafts.fxtf.org/geometry/#create-a-dommatrix-from-the-2d-dictionary){#transformations:create-a-dommatrix-from-a-2d-dictionary
    x-internal="create-a-dommatrix-from-a-2d-dictionary"}
    `transform`{.variable}.

2.  If one or more of `matrix`{.variable}\'s [m11
    element](https://drafts.fxtf.org/geometry/#matrix-m11-element){#transformations:m11-element
    x-internal="m11-element"}, [m12
    element](https://drafts.fxtf.org/geometry/#matrix-m12-element){#transformations:m12-element
    x-internal="m12-element"}, [m21
    element](https://drafts.fxtf.org/geometry/#matrix-m21-element){#transformations:m21-element
    x-internal="m21-element"}, [m22
    element](https://drafts.fxtf.org/geometry/#matrix-m22-element){#transformations:m22-element
    x-internal="m22-element"}, [m41
    element](https://drafts.fxtf.org/geometry/#matrix-m41-element){#transformations:m41-element
    x-internal="m41-element"}, or [m42
    element](https://drafts.fxtf.org/geometry/#matrix-m42-element){#transformations:m42-element
    x-internal="m42-element"} are infinite or NaN, then return.

3.  Reset the [current transformation
    matrix](#current-transformation-matrix){#transformations:current-transformation-matrix-19}
    to `matrix`{.variable}.

The [`resetTransform()`]{#dom-context-2d-resettransform .dfn
dfn-for="CanvasTransform" dfn-type="method"} method, when invoked, must
reset the [current transformation
matrix](#current-transformation-matrix){#transformations:current-transformation-matrix-20}
to the identity matrix.

::: note
Given a matrix of the form created by the
[`transform()`{#transformations:dom-context-2d-transform}](#dom-context-2d-transform)
and
[`setTransform()`{#transformations:dom-context-2d-settransform}](#dom-context-2d-settransform)
methods, i.e.,

`a`{.variable}

`c`{.variable}

`e`{.variable}

`b`{.variable}

`d`{.variable}

`f`{.variable}

0

0

1

the resulting transformed coordinates after transform matrix
multiplication will be

`x`{.variable}~new~ = `a`{.variable} `x`{.variable} + `c`{.variable}
`y`{.variable} + `e`{.variable}\
`y`{.variable}~new~ = `b`{.variable} `x`{.variable} + `d`{.variable}
`y`{.variable} + `f`{.variable}
:::

###### [4.12.5.1.9]{.secno} Image sources for 2D rendering contexts[](#image-sources-for-2d-rendering-contexts){.self-link}

Some methods on the
[`CanvasDrawImage`{#image-sources-for-2d-rendering-contexts:canvasdrawimage}](#canvasdrawimage)
and
[`CanvasFillStrokeStyles`{#image-sources-for-2d-rendering-contexts:canvasfillstrokestyles}](#canvasfillstrokestyles)
interfaces take the union type
[`CanvasImageSource`{#image-sources-for-2d-rendering-contexts:canvasimagesource}](#canvasimagesource)
as an argument.

This union type allows objects implementing any of the following
interfaces to be used as image sources:

- [`HTMLOrSVGImageElement`{#image-sources-for-2d-rendering-contexts:htmlorsvgimageelement}](#htmlorsvgimageelement)
  ([`img`{#image-sources-for-2d-rendering-contexts:the-img-element}](embedded-content.html#the-img-element)
  or [SVG
  `image`](https://svgwg.org/svg2-draft/embedded.html#ImageElement){#image-sources-for-2d-rendering-contexts:svg-image
  x-internal="svg-image"} elements)
- [`HTMLVideoElement`{#image-sources-for-2d-rendering-contexts:htmlvideoelement}](media.html#htmlvideoelement)
  ([`video`{#image-sources-for-2d-rendering-contexts:the-video-element}](media.html#the-video-element)
  elements)
- [`HTMLCanvasElement`{#image-sources-for-2d-rendering-contexts:htmlcanvaselement}](#htmlcanvaselement)
  ([`canvas`{#image-sources-for-2d-rendering-contexts:the-canvas-element}](#the-canvas-element)
  elements)
- [`OffscreenCanvas`{#image-sources-for-2d-rendering-contexts:offscreencanvas}](#offscreencanvas)
- [`ImageBitmap`{#image-sources-for-2d-rendering-contexts:imagebitmap}](imagebitmap-and-animations.html#imagebitmap)
- [`VideoFrame`{#image-sources-for-2d-rendering-contexts:videoframe}](https://w3c.github.io/webcodecs/#videoframe-interface){x-internal="videoframe"}

Although not formally specified as such, [SVG
`image`](https://svgwg.org/svg2-draft/embedded.html#ImageElement){#image-sources-for-2d-rendering-contexts:svg-image-2
x-internal="svg-image"} elements are expected to be implemented nearly
identical to
[`img`{#image-sources-for-2d-rendering-contexts:the-img-element-2}](embedded-content.html#the-img-element)
elements. That is, [SVG
`image`](https://svgwg.org/svg2-draft/embedded.html#ImageElement){#image-sources-for-2d-rendering-contexts:svg-image-3
x-internal="svg-image"} elements share the fundamental concepts and
features of
[`img`{#image-sources-for-2d-rendering-contexts:the-img-element-3}](embedded-content.html#the-img-element)
elements.

The
[`ImageBitmap`{#image-sources-for-2d-rendering-contexts:imagebitmap-2}](imagebitmap-and-animations.html#imagebitmap)
interface can be created from a number of other image-representing
types, including
[`ImageData`{#image-sources-for-2d-rendering-contexts:imagedata}](imagebitmap-and-animations.html#imagedata).

To [check the usability of the `image`{.variable}
argument]{#check-the-usability-of-the-image-argument .dfn}, where
`image`{.variable} is a
[`CanvasImageSource`{#image-sources-for-2d-rendering-contexts:canvasimagesource-2}](#canvasimagesource)
object, run these steps:

1.  Switch on `image`{.variable}:

    [`HTMLOrSVGImageElement`{#image-sources-for-2d-rendering-contexts:htmlorsvgimageelement-2}](#htmlorsvgimageelement)

    :   If `image`{.variable}\'s [current
        request](images.html#current-request){#image-sources-for-2d-rendering-contexts:current-request}\'s
        [state](images.html#img-req-state){#image-sources-for-2d-rendering-contexts:img-req-state}
        is
        [broken](images.html#img-error){#image-sources-for-2d-rendering-contexts:img-error},
        then throw an
        [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#image-sources-for-2d-rendering-contexts:invalidstateerror
        x-internal="invalidstateerror"}
        [`DOMException`{#image-sources-for-2d-rendering-contexts:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

        If `image`{.variable} is not [fully
        decodable](images.html#img-good){#image-sources-for-2d-rendering-contexts:img-good},
        then return *bad*.

        If `image`{.variable} has a [natural
        width](https://drafts.csswg.org/css-images/#natural-width){#image-sources-for-2d-rendering-contexts:natural-width
        x-internal="natural-width"} or [natural
        height](https://drafts.csswg.org/css-images/#natural-height){#image-sources-for-2d-rendering-contexts:natural-height
        x-internal="natural-height"} (or both) equal to zero, then
        return *bad*.

    [`HTMLVideoElement`{#image-sources-for-2d-rendering-contexts:htmlvideoelement-2}](media.html#htmlvideoelement)
    :   If `image`{.variable}\'s
        [`readyState`{#image-sources-for-2d-rendering-contexts:dom-media-readystate}](media.html#dom-media-readystate)
        attribute is either
        [`HAVE_NOTHING`{#image-sources-for-2d-rendering-contexts:dom-media-have_nothing}](media.html#dom-media-have_nothing)
        or
        [`HAVE_METADATA`{#image-sources-for-2d-rendering-contexts:dom-media-have_metadata}](media.html#dom-media-have_metadata),
        then return *bad*.

    [`HTMLCanvasElement`{#image-sources-for-2d-rendering-contexts:htmlcanvaselement-2}](#htmlcanvaselement)\
    [`OffscreenCanvas`{#image-sources-for-2d-rendering-contexts:offscreencanvas-2}](#offscreencanvas)
    :   If `image`{.variable} has either a horizontal dimension or a
        vertical dimension equal to zero, then throw an
        [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#image-sources-for-2d-rendering-contexts:invalidstateerror-2
        x-internal="invalidstateerror"}
        [`DOMException`{#image-sources-for-2d-rendering-contexts:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    [`ImageBitmap`{#image-sources-for-2d-rendering-contexts:imagebitmap-3}](imagebitmap-and-animations.html#imagebitmap)\
    [`VideoFrame`{#image-sources-for-2d-rendering-contexts:videoframe-2}](https://w3c.github.io/webcodecs/#videoframe-interface){x-internal="videoframe"}

    :   If `image`{.variable}\'s
        [\[\[Detached\]\]](structured-data.html#detached){#image-sources-for-2d-rendering-contexts:detached}
        internal slot value is set to true, then throw an
        [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#image-sources-for-2d-rendering-contexts:invalidstateerror-3
        x-internal="invalidstateerror"}
        [`DOMException`{#image-sources-for-2d-rendering-contexts:domexception-3}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  Return *good*.

When a
[`CanvasImageSource`{#image-sources-for-2d-rendering-contexts:canvasimagesource-3}](#canvasimagesource)
object represents an
[`HTMLOrSVGImageElement`{#image-sources-for-2d-rendering-contexts:htmlorsvgimageelement-3}](#htmlorsvgimageelement),
the element\'s image must be used as the source image.

Specifically, when a
[`CanvasImageSource`{#image-sources-for-2d-rendering-contexts:canvasimagesource-4}](#canvasimagesource)
object represents an animated image in an
[`HTMLOrSVGImageElement`{#image-sources-for-2d-rendering-contexts:htmlorsvgimageelement-4}](#htmlorsvgimageelement),
the user agent must use the default image of the animation (the one that
the format defines is to be used when animation is not supported or is
disabled), or, if there is no such image, the first frame of the
animation, when rendering the image for
[`CanvasRenderingContext2D`{#image-sources-for-2d-rendering-contexts:canvasrenderingcontext2d}](#canvasrenderingcontext2d)
APIs.

When a
[`CanvasImageSource`{#image-sources-for-2d-rendering-contexts:canvasimagesource-5}](#canvasimagesource)
object represents an
[`HTMLVideoElement`{#image-sources-for-2d-rendering-contexts:htmlvideoelement-3}](media.html#htmlvideoelement),
then the frame at the [current playback
position](media.html#current-playback-position){#image-sources-for-2d-rendering-contexts:current-playback-position}
when the method with the argument is invoked must be used as the source
image when rendering the image for
[`CanvasRenderingContext2D`{#image-sources-for-2d-rendering-contexts:canvasrenderingcontext2d-2}](#canvasrenderingcontext2d)
APIs, and the source image\'s dimensions must be the [natural
width](media.html#concept-video-intrinsic-width){#image-sources-for-2d-rendering-contexts:concept-video-intrinsic-width}
and [natural
height](media.html#concept-video-intrinsic-height){#image-sources-for-2d-rendering-contexts:concept-video-intrinsic-height}
of the [media
resource](media.html#media-resource){#image-sources-for-2d-rendering-contexts:media-resource}
(i.e., after any aspect-ratio correction has been applied).

When a
[`CanvasImageSource`{#image-sources-for-2d-rendering-contexts:canvasimagesource-6}](#canvasimagesource)
object represents an
[`HTMLCanvasElement`{#image-sources-for-2d-rendering-contexts:htmlcanvaselement-3}](#htmlcanvaselement),
the element\'s bitmap must be used as the source image.

When a
[`CanvasImageSource`{#image-sources-for-2d-rendering-contexts:canvasimagesource-7}](#canvasimagesource)
object represents an element that is [being
rendered](rendering.html#being-rendered){#image-sources-for-2d-rendering-contexts:being-rendered}
and that element has been resized, the original image data of the source
image must be used, not the image as it is rendered (e.g.
[`width`{#image-sources-for-2d-rendering-contexts:attr-dim-width}](embedded-content-other.html#attr-dim-width)
and
[`height`{#image-sources-for-2d-rendering-contexts:attr-dim-height}](embedded-content-other.html#attr-dim-height)
attributes on the source element have no effect on how the object is
interpreted when rendering the image for
[`CanvasRenderingContext2D`{#image-sources-for-2d-rendering-contexts:canvasrenderingcontext2d-3}](#canvasrenderingcontext2d)
APIs).

When a
[`CanvasImageSource`{#image-sources-for-2d-rendering-contexts:canvasimagesource-8}](#canvasimagesource)
object represents an
[`ImageBitmap`{#image-sources-for-2d-rendering-contexts:imagebitmap-4}](imagebitmap-and-animations.html#imagebitmap),
the object\'s bitmap image data must be used as the source image.

When a
[`CanvasImageSource`{#image-sources-for-2d-rendering-contexts:canvasimagesource-9}](#canvasimagesource)
object represents an
[`OffscreenCanvas`{#image-sources-for-2d-rendering-contexts:offscreencanvas-3}](#offscreencanvas),
the object\'s bitmap must be used as the source image.

When a
[`CanvasImageSource`{#image-sources-for-2d-rendering-contexts:canvasimagesource-10}](#canvasimagesource)
object represents a
[`VideoFrame`{#image-sources-for-2d-rendering-contexts:videoframe-3}](https://w3c.github.io/webcodecs/#videoframe-interface){x-internal="videoframe"},
the object\'s pixel data must be used as the source image, and the
source image\'s dimensions must be the object\'s [\[\[display
width\]\]](https://w3c.github.io/webcodecs/#dom-videoframe-display-width-slot){#image-sources-for-2d-rendering-contexts:display-width
x-internal="display-width"} and [\[\[display
height\]\]](https://w3c.github.io/webcodecs/#dom-videoframe-display-height-slot){#image-sources-for-2d-rendering-contexts:display-height
x-internal="display-height"}.

An object `image`{.variable} [is not
origin-clean]{#the-image-argument-is-not-origin-clean .dfn} if,
switching on `image`{.variable}\'s type:

[`HTMLOrSVGImageElement`{#image-sources-for-2d-rendering-contexts:htmlorsvgimageelement-5}](#htmlorsvgimageelement)
:   `image`{.variable}\'s [current
    request](images.html#current-request){#image-sources-for-2d-rendering-contexts:current-request-2}\'s
    [image
    data](images.html#img-req-data){#image-sources-for-2d-rendering-contexts:img-req-data}
    is
    [CORS-cross-origin](urls-and-fetching.html#cors-cross-origin){#image-sources-for-2d-rendering-contexts:cors-cross-origin}.

[`HTMLVideoElement`{#image-sources-for-2d-rendering-contexts:htmlvideoelement-4}](media.html#htmlvideoelement)
:   `image`{.variable}\'s [media
    data](media.html#media-data){#image-sources-for-2d-rendering-contexts:media-data}
    is
    [CORS-cross-origin](urls-and-fetching.html#cors-cross-origin){#image-sources-for-2d-rendering-contexts:cors-cross-origin-2}.

[`HTMLCanvasElement`{#image-sources-for-2d-rendering-contexts:htmlcanvaselement-4}](#htmlcanvaselement)\
[`ImageBitmap`{#image-sources-for-2d-rendering-contexts:imagebitmap-5}](imagebitmap-and-animations.html#imagebitmap)\
[`OffscreenCanvas`{#image-sources-for-2d-rendering-contexts:offscreencanvas-4}](#offscreencanvas)

:   `image`{.variable}\'s bitmap\'s
    [origin-clean](#concept-canvas-origin-clean){#image-sources-for-2d-rendering-contexts:concept-canvas-origin-clean}
    flag is false.

###### [4.12.5.1.10]{.secno} Fill and stroke styles[](#fill-and-stroke-styles){.self-link}

`context`{.variable}`.`[`fillStyle`](#dom-context-2d-fillstyle){#dom-context-2d-fillstyle-dev}` [ = ``value`{.variable}` ]`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/fillStyle](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/fillStyle "The CanvasRenderingContext2D.fillStyle property of the Canvas 2D API specifies the color, gradient, or pattern to use inside shapes. The default style is #000 (black).")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari2+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns the current style used for filling shapes.

Can be set, to change the [fill
style](#concept-canvasfillstrokestyles-fill-style){#fill-and-stroke-styles:concept-canvasfillstrokestyles-fill-style}.

The style can be either a string containing a CSS color, or a
[`CanvasGradient`{#fill-and-stroke-styles:canvasgradient}](#canvasgradient)
or
[`CanvasPattern`{#fill-and-stroke-styles:canvaspattern}](#canvaspattern)
object. Invalid values are ignored.

`context`{.variable}`.`[`strokeStyle`](#dom-context-2d-strokestyle){#dom-context-2d-strokestyle-dev}` [ = ``value`{.variable}` ]`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/strokeStyle](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/strokeStyle "The CanvasRenderingContext2D.strokeStyle property of the Canvas 2D API specifies the color, gradient, or pattern to use for the strokes (outlines) around shapes. The default is #000 (black).")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari2+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns the current style used for stroking shapes.

Can be set, to change the [stroke
style.](#concept-canvasfillstrokestyles-stroke-style){#fill-and-stroke-styles:concept-canvasfillstrokestyles-stroke-style}

The style can be either a string containing a CSS color, or a
[`CanvasGradient`{#fill-and-stroke-styles:canvasgradient-2}](#canvasgradient)
or
[`CanvasPattern`{#fill-and-stroke-styles:canvaspattern-2}](#canvaspattern)
object. Invalid values are ignored.

Objects that implement the
[`CanvasFillStrokeStyles`{#fill-and-stroke-styles:canvasfillstrokestyles}](#canvasfillstrokestyles)
interface have attributes and methods (defined in this section) that
control how shapes are treated by the object.

Such objects have associated [fill
style]{#concept-canvasfillstrokestyles-fill-style .dfn} and [stroke
style]{#concept-canvasfillstrokestyles-stroke-style .dfn} values, which
are either CSS colors,
[`CanvasPattern`{#fill-and-stroke-styles:canvaspattern-3}](#canvaspattern)s,
or
[`CanvasGradient`{#fill-and-stroke-styles:canvasgradient-3}](#canvasgradient)s.
Initially, both must be the result of
[parsing](https://drafts.csswg.org/css-color/#parse-a-css-color-value){#fill-and-stroke-styles:parsed-as-a-css-color-value
x-internal="parsed-as-a-css-color-value"} the string \"`#000000`\".

When the value is a CSS color, it must not be affected by the
transformation matrix when used to draw on bitmaps.

When set to a
[`CanvasPattern`{#fill-and-stroke-styles:canvaspattern-4}](#canvaspattern)
or
[`CanvasGradient`{#fill-and-stroke-styles:canvasgradient-4}](#canvasgradient)
object, changes made to the object after the assignment do affect
subsequent stroking or filling of shapes.

The [`fillStyle`]{#dom-context-2d-fillstyle .dfn
dfn-for="CanvasFillStrokeStyles" dfn-type="attribute"} getter steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#fill-and-stroke-styles:this
    x-internal="this"}\'s [fill
    style](#concept-canvasfillstrokestyles-fill-style){#fill-and-stroke-styles:concept-canvasfillstrokestyles-fill-style-2}
    is a CSS color, then return the
    [serialization](https://drafts.csswg.org/css-color/#serializing-color-values){#fill-and-stroke-styles:serialisation-of-a-color
    x-internal="serialisation-of-a-color"} of that color with
    [HTML-compatible serialization
    requested](https://drafts.csswg.org/css-color/#color-serialization-html-compatible-serialization-is-requested){#fill-and-stroke-styles:html-compatible-serialization-is-requested
    x-internal="html-compatible-serialization-is-requested"}.

2.  Return
    [this](https://webidl.spec.whatwg.org/#this){#fill-and-stroke-styles:this-2
    x-internal="this"}\'s [fill
    style](#concept-canvasfillstrokestyles-fill-style){#fill-and-stroke-styles:concept-canvasfillstrokestyles-fill-style-3}.

The
[`fillStyle`{#fill-and-stroke-styles:dom-context-2d-fillstyle}](#dom-context-2d-fillstyle)
setter steps are:

1.  If the given value is a string, then:

    1.  Let `context`{.variable} be
        [this](https://webidl.spec.whatwg.org/#this){#fill-and-stroke-styles:this-3
        x-internal="this"}\'s
        [`canvas`{#fill-and-stroke-styles:dom-context-2d-canvas}](#dom-context-2d-canvas)
        attribute\'s value, if that is an element; otherwise null.

    2.  Let `parsedValue`{.variable} be the result of
        [parsing](https://drafts.csswg.org/css-color/#parse-a-css-color-value){#fill-and-stroke-styles:parsed-as-a-css-color-value-2
        x-internal="parsed-as-a-css-color-value"} the given value with
        `context`{.variable} if non-null.

    3.  If `parsedValue`{.variable} is failure, then return.

    4.  Set
        [this](https://webidl.spec.whatwg.org/#this){#fill-and-stroke-styles:this-4
        x-internal="this"}\'s [fill
        style](#concept-canvasfillstrokestyles-fill-style){#fill-and-stroke-styles:concept-canvasfillstrokestyles-fill-style-4}
        to `parsedValue`{.variable}.

    5.  Return.

2.  If the given value is a
    [`CanvasPattern`{#fill-and-stroke-styles:canvaspattern-5}](#canvaspattern)
    object that is marked as [not
    origin-clean](#concept-canvas-pattern-not-origin-clean){#fill-and-stroke-styles:concept-canvas-pattern-not-origin-clean},
    then set
    [this](https://webidl.spec.whatwg.org/#this){#fill-and-stroke-styles:this-5
    x-internal="this"}\'s
    [origin-clean](#concept-canvas-origin-clean){#fill-and-stroke-styles:concept-canvas-origin-clean}
    flag to false.

3.  Set
    [this](https://webidl.spec.whatwg.org/#this){#fill-and-stroke-styles:this-6
    x-internal="this"}\'s [fill
    style](#concept-canvasfillstrokestyles-fill-style){#fill-and-stroke-styles:concept-canvasfillstrokestyles-fill-style-5}
    to the given value.

The [`strokeStyle`]{#dom-context-2d-strokestyle .dfn
dfn-for="CanvasFillStrokeStyles" dfn-type="attribute"} getter steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#fill-and-stroke-styles:this-7
    x-internal="this"}\'s [stroke
    style](#concept-canvasfillstrokestyles-stroke-style){#fill-and-stroke-styles:concept-canvasfillstrokestyles-stroke-style-2}
    is a CSS color, then return the
    [serialization](https://drafts.csswg.org/css-color/#serializing-color-values){#fill-and-stroke-styles:serialisation-of-a-color-2
    x-internal="serialisation-of-a-color"} of that color with
    [HTML-compatible serialization
    requested](https://drafts.csswg.org/css-color/#color-serialization-html-compatible-serialization-is-requested){#fill-and-stroke-styles:html-compatible-serialization-is-requested-2
    x-internal="html-compatible-serialization-is-requested"}.

2.  Return
    [this](https://webidl.spec.whatwg.org/#this){#fill-and-stroke-styles:this-8
    x-internal="this"}\'s [stroke
    style](#concept-canvasfillstrokestyles-stroke-style){#fill-and-stroke-styles:concept-canvasfillstrokestyles-stroke-style-3}.

The
[`strokeStyle`{#fill-and-stroke-styles:dom-context-2d-strokestyle}](#dom-context-2d-strokestyle)
setter steps are:

1.  If the given value is a string, then:

    1.  Let `context`{.variable} be
        [this](https://webidl.spec.whatwg.org/#this){#fill-and-stroke-styles:this-9
        x-internal="this"}\'s
        [`canvas`{#fill-and-stroke-styles:dom-context-2d-canvas-2}](#dom-context-2d-canvas)
        attribute\'s value, if that is an element; otherwise null.

    2.  Let `parsedValue`{.variable} be the result of
        [parsing](https://drafts.csswg.org/css-color/#parse-a-css-color-value){#fill-and-stroke-styles:parsed-as-a-css-color-value-3
        x-internal="parsed-as-a-css-color-value"} the given value with
        `context`{.variable} if non-null.

    3.  If `parsedValue`{.variable} is failure, then return.

    4.  Set
        [this](https://webidl.spec.whatwg.org/#this){#fill-and-stroke-styles:this-10
        x-internal="this"}\'s [stroke
        style](#concept-canvasfillstrokestyles-stroke-style){#fill-and-stroke-styles:concept-canvasfillstrokestyles-stroke-style-4}
        to `parsedValue`{.variable}.

    5.  Return.

2.  If the given value is a
    [`CanvasPattern`{#fill-and-stroke-styles:canvaspattern-6}](#canvaspattern)
    object that is marked as [not
    origin-clean](#concept-canvas-pattern-not-origin-clean){#fill-and-stroke-styles:concept-canvas-pattern-not-origin-clean-2},
    then set
    [this](https://webidl.spec.whatwg.org/#this){#fill-and-stroke-styles:this-11
    x-internal="this"}\'s
    [origin-clean](#concept-canvas-origin-clean){#fill-and-stroke-styles:concept-canvas-origin-clean-2}
    flag to false.

3.  Set
    [this](https://webidl.spec.whatwg.org/#this){#fill-and-stroke-styles:this-12
    x-internal="this"}\'s [stroke
    style](#concept-canvasfillstrokestyles-stroke-style){#fill-and-stroke-styles:concept-canvasfillstrokestyles-stroke-style-5}
    to the given value.

------------------------------------------------------------------------

There are three types of gradients, linear gradients, radial gradients,
and conic gradients, represented by objects implementing the opaque
[`CanvasGradient`{#fill-and-stroke-styles:canvasgradient-5}](#canvasgradient)
interface.

Once a gradient has been created (see below), stops are placed along it
to define how the colors are distributed along the gradient. The color
of the gradient at each stop is the color specified for that stop.
Between each such stop, the colors and the alpha component must be
linearly interpolated over the RGBA space without premultiplying the
alpha value to find the color to use at that offset. Before the first
stop, the color must be the color of the first stop. After the last
stop, the color must be the color of the last stop. When there are no
stops, the gradient is [transparent
black](https://drafts.csswg.org/css-color/#transparent-black){#fill-and-stroke-styles:transparent-black
x-internal="transparent-black"}.

`gradient`{.variable}`.`[`addColorStop`](#dom-canvasgradient-addcolorstop){#dom-canvasgradient-addcolorstop-dev}`(``offset`{.variable}`, ``color`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasGradient/addColorStop](https://developer.mozilla.org/en-US/docs/Web/API/CanvasGradient/addColorStop "The CanvasGradient.addColorStop() method adds a new color stop, defined by an offset and a color, to a given canvas gradient.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari2+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera9+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

Adds a color stop with the given color to the gradient at the given
offset. 0.0 is the offset at one end of the gradient, 1.0 is the offset
at the other end.

Throws an
[\"`IndexSizeError`\"](https://webidl.spec.whatwg.org/#indexsizeerror){#fill-and-stroke-styles:indexsizeerror
x-internal="indexsizeerror"}
[`DOMException`{#fill-and-stroke-styles:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if the offset is out of range. Throws a
[\"`SyntaxError`\"](https://webidl.spec.whatwg.org/#syntaxerror){#fill-and-stroke-styles:syntaxerror
x-internal="syntaxerror"}
[`DOMException`{#fill-and-stroke-styles:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if the color cannot be parsed.

`gradient`{.variable}` = ``context`{.variable}`.`[`createLinearGradient`](#dom-context-2d-createlineargradient){#dom-context-2d-createlineargradient-dev}`(``x0`{.variable}`, ``y0`{.variable}`, ``x1`{.variable}`, ``y1`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/createLinearGradient](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/createLinearGradient "The CanvasRenderingContext2D.createLinearGradient() method of the Canvas 2D API creates a gradient along the line connecting two given coordinates.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari2+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns a
[`CanvasGradient`{#fill-and-stroke-styles:canvasgradient-6}](#canvasgradient)
object that represents a linear gradient that paints along the line
given by the coordinates represented by the arguments.

`gradient`{.variable}` = ``context`{.variable}`.`[`createRadialGradient`](#dom-context-2d-createradialgradient){#dom-context-2d-createradialgradient-dev}`(``x0`{.variable}`, ``y0`{.variable}`, ``r0`{.variable}`, ``x1`{.variable}`, ``y1`{.variable}`, ``r1`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/createRadialGradient](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/createRadialGradient "The CanvasRenderingContext2D.createRadialGradient() method of the Canvas 2D API creates a radial gradient using the size and coordinates of two circles.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari2+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns a
[`CanvasGradient`{#fill-and-stroke-styles:canvasgradient-7}](#canvasgradient)
object that represents a radial gradient that paints along the cone
given by the circles represented by the arguments.

If either of the radii are negative, throws an
[\"`IndexSizeError`\"](https://webidl.spec.whatwg.org/#indexsizeerror){#fill-and-stroke-styles:indexsizeerror-2
x-internal="indexsizeerror"}
[`DOMException`{#fill-and-stroke-styles:domexception-3}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

`gradient`{.variable}` = ``context`{.variable}`.`[`createConicGradient`](#dom-context-2d-createconicgradient){#dom-context-2d-createconicgradient-dev}`(``startAngle`{.variable}`, ``x`{.variable}`, ``y`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/createConicGradient](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/createConicGradient "The CanvasRenderingContext2D.createConicGradient() method of the Canvas 2D API creates a gradient around a point with given coordinates.")

Support in all current engines.

::: support
[Firefox112+]{.firefox .yes}[Safari16.1+]{.safari
.yes}[Chrome99+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge99+]{.edge_blink .yes}

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

Returns a
[`CanvasGradient`{#fill-and-stroke-styles:canvasgradient-8}](#canvasgradient)
object that represents a conic gradient that paints clockwise along the
rotation around the center represented by the arguments.

The
[`addColorStop(``offset`{.variable}`, ``color`{.variable}`)`]{#dom-canvasgradient-addcolorstop
.dfn dfn-for="CanvasGradient" dfn-type="method"} method on the
[`CanvasGradient`{#fill-and-stroke-styles:canvasgradient-9}](#canvasgradient),
when invoked, must run these steps:

1.  If the `offset`{.variable} is less than 0 or greater than 1, then
    throw an
    [\"`IndexSizeError`\"](https://webidl.spec.whatwg.org/#indexsizeerror){#fill-and-stroke-styles:indexsizeerror-3
    x-internal="indexsizeerror"}
    [`DOMException`{#fill-and-stroke-styles:domexception-4}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  Let `parsed color`{.variable} be the result of
    [parsing](https://drafts.csswg.org/css-color/#parse-a-css-color-value){#fill-and-stroke-styles:parsed-as-a-css-color-value-4
    x-internal="parsed-as-a-css-color-value"} `color`{.variable}.

    No element is passed to the parser because
    [`CanvasGradient`{#fill-and-stroke-styles:canvasgradient-10}](#canvasgradient)
    objects are
    [`canvas`{#fill-and-stroke-styles:the-canvas-element}](#the-canvas-element)-neutral
    --- a
    [`CanvasGradient`{#fill-and-stroke-styles:canvasgradient-11}](#canvasgradient)
    object created by one
    [`canvas`{#fill-and-stroke-styles:the-canvas-element-2}](#the-canvas-element)
    can be used by another, and there is therefore no way to know which
    is the \"element in question\" at the time that the color is
    specified.

3.  If `parsed color`{.variable} is failure, throw a
    [\"`SyntaxError`\"](https://webidl.spec.whatwg.org/#syntaxerror){#fill-and-stroke-styles:syntaxerror-2
    x-internal="syntaxerror"}
    [`DOMException`{#fill-and-stroke-styles:domexception-5}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

4.  Place a new stop on the gradient, at offset `offset`{.variable}
    relative to the whole gradient, and with the color
    `parsed color`{.variable}.

    If multiple stops are added at the same offset on a gradient, then
    they must be placed in the order added, with the first one closest
    to the start of the gradient, and each subsequent one
    infinitesimally further along towards the end point (in effect
    causing all but the first and last stop added at each point to be
    ignored).

The
[`createLinearGradient(``x0`{.variable}`, ``y0`{.variable}`, ``x1`{.variable}`, ``y1`{.variable}`)`]{#dom-context-2d-createlineargradient
.dfn dfn-for="CanvasFillStrokeStyles" dfn-type="method"} method takes
four arguments that represent the start point (`x0`{.variable},
`y0`{.variable}) and end point (`x1`{.variable}, `y1`{.variable}) of the
gradient. The method, when invoked, must return a linear
[`CanvasGradient`{#fill-and-stroke-styles:canvasgradient-12}](#canvasgradient)
initialized with the specified line.

Linear gradients must be rendered such that all points on a line
perpendicular to the line that crosses the start and end points have the
color at the point where those two lines cross (with the colors coming
from the [interpolation and extrapolation](#interpolation) described
above). The points in the linear gradient must be transformed as
described by the [current transformation
matrix](#transformations){#fill-and-stroke-styles:transformations} when
rendering.

If `x0`{.variable} = `x1`{.variable} and
`y0`{.variable} = `y1`{.variable}, then the linear gradient must paint
nothing.

The
[`createRadialGradient(``x0`{.variable}`, ``y0`{.variable}`, ``r0`{.variable}`, ``x1`{.variable}`, ``y1`{.variable}`, ``r1`{.variable}`)`]{#dom-context-2d-createradialgradient
.dfn dfn-for="CanvasFillStrokeStyles" dfn-type="method"} method takes
six arguments, the first three representing the start circle with origin
(`x0`{.variable}, `y0`{.variable}) and radius `r0`{.variable}, and the
last three representing the end circle with origin (`x1`{.variable},
`y1`{.variable}) and radius `r1`{.variable}. The values are in
coordinate space units. If either of `r0`{.variable} or `r1`{.variable}
are negative, then an
[\"`IndexSizeError`\"](https://webidl.spec.whatwg.org/#indexsizeerror){#fill-and-stroke-styles:indexsizeerror-4
x-internal="indexsizeerror"}
[`DOMException`{#fill-and-stroke-styles:domexception-6}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
must be thrown. Otherwise, the method, when invoked, must return a
radial
[`CanvasGradient`{#fill-and-stroke-styles:canvasgradient-13}](#canvasgradient)
initialized with the two specified circles.

Radial gradients must be rendered by following these steps:

1.  If `x`{.variable}~`0`{.variable}~ = `x`{.variable}~`1`{.variable}~
    and `y`{.variable}~`0`{.variable}~ = `y`{.variable}~`1`{.variable}~
    and `r`{.variable}~`0`{.variable}~ = `r`{.variable}~`1`{.variable}~,
    then the radial gradient must paint nothing. Return.

2.  Let
    x(`ω`{.variable}) = (`x`{.variable}~`1`{.variable}~-`x`{.variable}~`0`{.variable}~)`ω`{.variable} + `x`{.variable}~`0`{.variable}~

    Let
    y(`ω`{.variable}) = (`y`{.variable}~`1`{.variable}~-`y`{.variable}~`0`{.variable}~)`ω`{.variable} + `y`{.variable}~`0`{.variable}~

    Let
    r(`ω`{.variable}) = (`r`{.variable}~`1`{.variable}~-`r`{.variable}~`0`{.variable}~)`ω`{.variable} + `r`{.variable}~`0`{.variable}~

    Let the color at `ω`{.variable} be the color at that position on the
    gradient (with the colors coming from the [interpolation and
    extrapolation](#interpolation) described above).

3.  For all values of `ω`{.variable} where r(`ω`{.variable}) \> 0,
    starting with the value of `ω`{.variable} nearest to positive
    infinity and ending with the value of `ω`{.variable} nearest to
    negative infinity, draw the circumference of the circle with radius
    r(`ω`{.variable}) at position (x(`ω`{.variable}),
    y(`ω`{.variable})), with the color at `ω`{.variable}, but only
    painting on the parts of the bitmap that have not yet been painted
    on by earlier circles in this step for this rendering of the
    gradient.

This effectively creates a cone, touched by the two circles defined in
the creation of the gradient, with the part of the cone before the start
circle (0.0) using the color of the first offset, the part of the cone
after the end circle (1.0) using the color of the last offset, and areas
outside the cone untouched by the gradient ([transparent
black](https://drafts.csswg.org/css-color/#transparent-black){#fill-and-stroke-styles:transparent-black-2
x-internal="transparent-black"}).

The resulting radial gradient must then be transformed as described by
the [current transformation
matrix](#transformations){#fill-and-stroke-styles:transformations-2}
when rendering.

The
[`createConicGradient(``startAngle`{.variable}`, ``x`{.variable}`, ``y`{.variable}`)`]{#dom-context-2d-createconicgradient
.dfn dfn-for="CanvasFillStrokeStyles" dfn-type="method"} method takes
three arguments, the first argument, `startAngle`{.variable}, represents
the angle in radians at which the gradient begins, and the last two
arguments, (`x`{.variable}, `y`{.variable}), represent the center of the
gradient in [CSS
pixels](https://drafts.csswg.org/css-values/#px){#fill-and-stroke-styles:'px'
x-internal="'px'"}. The method, when invoked, must return a conic
[`CanvasGradient`{#fill-and-stroke-styles:canvasgradient-14}](#canvasgradient)
initialized with the specified center and angle.

It follows the same rendering rule as CSS
[\'conic-gradient\'](https://drafts.csswg.org/css-images-4/#funcdef-conic-gradient){#fill-and-stroke-styles:'conic-gradient'
x-internal="'conic-gradient'"} and it is equivalent to CSS
\'conic-gradient(from `adjustedStartAngle`{.variable}rad at
`x`{.variable}px `y`{.variable}px, `angularColorStopList`{.variable})\'.
Here:

- `adjustedStartAngle`{.variable} is given by `startAngle`{.variable} +
  π/2;

- `angularColorStopList`{.variable} is given by the color stops that
  have been added to the
  [`CanvasGradient`{#fill-and-stroke-styles:canvasgradient-15}](#canvasgradient)
  using
  [`addColorStop()`{#fill-and-stroke-styles:dom-canvasgradient-addcolorstop}](#dom-canvasgradient-addcolorstop),
  with the color stop offsets interpreted as percentages.

Gradients must be painted only where the relevant stroking or filling
effects requires that they be drawn.

------------------------------------------------------------------------

Patterns are represented by objects implementing the opaque
[`CanvasPattern`{#fill-and-stroke-styles:canvaspattern-7}](#canvaspattern)
interface.

`pattern`{.variable}` = ``context`{.variable}`.`[`createPattern`](#dom-context-2d-createpattern){#dom-context-2d-createpattern-dev}`(``image`{.variable}`, ``repetition`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/createPattern](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/createPattern "The CanvasRenderingContext2D.createPattern() method of the Canvas 2D API creates a pattern using the specified image and repetition. This method returns a CanvasPattern.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari2+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns a
[`CanvasPattern`{#fill-and-stroke-styles:canvaspattern-8}](#canvaspattern)
object that uses the given image and repeats in the direction(s) given
by the `repetition`{.variable} argument.

The allowed values for `repetition`{.variable} are `repeat` (both
directions), `repeat-x` (horizontal only), `repeat-y` (vertical only),
and `no-repeat` (neither). If the `repetition`{.variable} argument is
empty, the value `repeat` is used.

If the image isn\'t yet fully decoded, then nothing is drawn. If the
image is a canvas with no data, throws an
[\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#fill-and-stroke-styles:invalidstateerror
x-internal="invalidstateerror"}
[`DOMException`{#fill-and-stroke-styles:domexception-7}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

`pattern`{.variable}`.`[`setTransform`](#dom-canvaspattern-settransform){#dom-canvaspattern-settransform-dev}`(``transform`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasPattern/setTransform](https://developer.mozilla.org/en-US/docs/Web/API/CanvasPattern/setTransform "The CanvasPattern.setTransform() method uses a DOMMatrix object as the pattern's transformation matrix and invokes it on the pattern.")

Support in all current engines.

::: support
[Firefox33+]{.firefox .yes}[Safari11.1+]{.safari
.yes}[Chrome68+]{.chrome .yes}

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

Sets the transformation matrix that will be used when rendering the
pattern during a fill or stroke painting operation.

The
[`createPattern(``image`{.variable}`, ``repetition`{.variable}`)`]{#dom-context-2d-createpattern
.dfn dfn-for="CanvasFillStrokeStyles" dfn-type="method"} method, when
invoked, must run these steps:

1.  Let `usability`{.variable} be the result of [checking the usability
    of](#check-the-usability-of-the-image-argument){#fill-and-stroke-styles:check-the-usability-of-the-image-argument}
    `image`{.variable}.

2.  If `usability`{.variable} is *bad*, then return null.

3.  [Assert](https://infra.spec.whatwg.org/#assert){#fill-and-stroke-styles:assert
    x-internal="assert"}: `usability`{.variable} is *good*.

4.  If `repetition`{.variable} is the empty string, then set it to
    \"`repeat`\".

5.  If `repetition`{.variable} is not [identical
    to](https://infra.spec.whatwg.org/#string-is){#fill-and-stroke-styles:identical-to
    x-internal="identical-to"} one of \"`repeat`\", \"`repeat-x`\",
    \"`repeat-y`\", or \"`no-repeat`\", then throw a
    [\"`SyntaxError`\"](https://webidl.spec.whatwg.org/#syntaxerror){#fill-and-stroke-styles:syntaxerror-3
    x-internal="syntaxerror"}
    [`DOMException`{#fill-and-stroke-styles:domexception-8}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

6.  Let `pattern`{.variable} be a new
    [`CanvasPattern`{#fill-and-stroke-styles:canvaspattern-9}](#canvaspattern)
    object with the image `image`{.variable} and the repetition behavior
    given by `repetition`{.variable}.

7.  If `image`{.variable} [is not
    origin-clean](#the-image-argument-is-not-origin-clean){#fill-and-stroke-styles:the-image-argument-is-not-origin-clean},
    then mark `pattern`{.variable} as [not
    origin-clean]{#concept-canvas-pattern-not-origin-clean .dfn}.

8.  Return `pattern`{.variable}.

Modifying the `image`{.variable} used when creating a
[`CanvasPattern`{#fill-and-stroke-styles:canvaspattern-10}](#canvaspattern)
object after calling the
[`createPattern()`{#fill-and-stroke-styles:dom-context-2d-createpattern}](#dom-context-2d-createpattern)
method must not affect the pattern(s) rendered by the
[`CanvasPattern`{#fill-and-stroke-styles:canvaspattern-11}](#canvaspattern)
object.

Patterns have a transformation matrix, which controls how the pattern is
used when it is painted. Initially, a pattern\'s transformation matrix
must be the identity matrix.

The
[`setTransform(``transform`{.variable}`)`]{#dom-canvaspattern-settransform
.dfn dfn-for="CanvasPattern" dfn-type="method"} method, when invoked,
must run these steps:

1.  Let `matrix`{.variable} be the result of [creating a `DOMMatrix`
    from the 2D
    dictionary](https://drafts.fxtf.org/geometry/#create-a-dommatrix-from-the-2d-dictionary){#fill-and-stroke-styles:create-a-dommatrix-from-a-2d-dictionary
    x-internal="create-a-dommatrix-from-a-2d-dictionary"}
    `transform`{.variable}.

2.  If one or more of `matrix`{.variable}\'s [m11
    element](https://drafts.fxtf.org/geometry/#matrix-m11-element){#fill-and-stroke-styles:m11-element
    x-internal="m11-element"}, [m12
    element](https://drafts.fxtf.org/geometry/#matrix-m12-element){#fill-and-stroke-styles:m12-element
    x-internal="m12-element"}, [m21
    element](https://drafts.fxtf.org/geometry/#matrix-m21-element){#fill-and-stroke-styles:m21-element
    x-internal="m21-element"}, [m22
    element](https://drafts.fxtf.org/geometry/#matrix-m22-element){#fill-and-stroke-styles:m22-element
    x-internal="m22-element"}, [m41
    element](https://drafts.fxtf.org/geometry/#matrix-m41-element){#fill-and-stroke-styles:m41-element
    x-internal="m41-element"}, or [m42
    element](https://drafts.fxtf.org/geometry/#matrix-m42-element){#fill-and-stroke-styles:m42-element
    x-internal="m42-element"} are infinite or NaN, then return.

3.  Reset the pattern\'s transformation matrix to `matrix`{.variable}.

When a pattern is to be rendered within an area, the user agent must run
the following steps to determine what is rendered:

1.  Create an infinite [transparent
    black](https://drafts.csswg.org/css-color/#transparent-black){#fill-and-stroke-styles:transparent-black-3
    x-internal="transparent-black"} bitmap.

2.  Place a copy of the image on the bitmap, anchored such that its top
    left corner is at the origin of the coordinate space, with one
    coordinate space unit per [CSS
    pixel](https://drafts.csswg.org/css-values/#px){#fill-and-stroke-styles:'px'-2
    x-internal="'px'"} of the image, then place repeated copies of this
    image horizontally to the left and right, if the repetition behavior
    is \"`repeat-x`\", or vertically up and down, if the repetition
    behavior is \"`repeat-y`\", or in all four directions all over the
    bitmap, if the repetition behavior is \"`repeat`\".

    If the original image data is a bitmap image, then the value painted
    at a point in the area of the repetitions is computed by filtering
    the original image data. When scaling up, if the
    [`imageSmoothingEnabled`{#fill-and-stroke-styles:dom-context-2d-imagesmoothingenabled}](#dom-context-2d-imagesmoothingenabled)
    attribute is set to false, then the image must be rendered using
    nearest-neighbor interpolation. Otherwise, the user agent may use
    any filtering algorithm (for example bilinear interpolation or
    nearest-neighbor). User agents which support multiple filtering
    algorithms may use the value of the
    [`imageSmoothingQuality`{#fill-and-stroke-styles:dom-context-2d-imagesmoothingquality}](#dom-context-2d-imagesmoothingquality)
    attribute to guide the choice of filtering algorithm. When such a
    filtering algorithm requires a pixel value from outside the original
    image data, it must instead use the value from wrapping the pixel\'s
    coordinates to the original image\'s dimensions. (That is, the
    filter uses \'repeat\' behavior, regardless of the value of the
    pattern\'s repetition behavior.)

3.  Transform the resulting bitmap according to the pattern\'s
    transformation matrix.

4.  Transform the resulting bitmap again, this time according to the
    [current transformation
    matrix](#transformations){#fill-and-stroke-styles:transformations-3}.

5.  Replace any part of the image outside the area in which the pattern
    is to be rendered with [transparent
    black](https://drafts.csswg.org/css-color/#transparent-black){#fill-and-stroke-styles:transparent-black-4
    x-internal="transparent-black"}.

6.  The resulting bitmap is what is to be rendered, with the same origin
    and same scale.

------------------------------------------------------------------------

If a radial gradient or repeated pattern is used when the transformation
matrix is singular, then the resulting style must be [transparent
black](https://drafts.csswg.org/css-color/#transparent-black){#fill-and-stroke-styles:transparent-black-5
x-internal="transparent-black"} (otherwise the gradient or pattern would
be collapsed to a point or line, leaving the other pixels undefined).
Linear gradients and solid colors always define all points even with
singular transformation matrices.

###### [4.12.5.1.11]{.secno} Drawing rectangles to the bitmap[](#drawing-rectangles-to-the-bitmap){.self-link}

Objects that implement the
[`CanvasRect`{#drawing-rectangles-to-the-bitmap:canvasrect}](#canvasrect)
interface provide the following methods for immediately drawing
rectangles to the bitmap. The methods each take four arguments; the
first two give the `x`{.variable} and `y`{.variable} coordinates of the
top left of the rectangle, and the second two give the width
`w`{.variable} and height `h`{.variable} of the rectangle, respectively.

The [current transformation
matrix](#transformations){#drawing-rectangles-to-the-bitmap:transformations}
must be applied to the following four coordinates, which form the path
that must then be closed to get the specified rectangle:
(`x`{.variable}, `y`{.variable}), (`x`{.variable}+`w`{.variable},
`y`{.variable}), (`x`{.variable}+`w`{.variable},
`y`{.variable}+`h`{.variable}), (`x`{.variable},
`y`{.variable}+`h`{.variable}).

Shapes are painted without affecting the [current default
path](#current-default-path){#drawing-rectangles-to-the-bitmap:current-default-path},
and are subject to the [clipping
region](#clipping-region){#drawing-rectangles-to-the-bitmap:clipping-region},
and, with the exception of
[`clearRect()`{#drawing-rectangles-to-the-bitmap:dom-context-2d-clearrect}](#dom-context-2d-clearrect),
also [shadow
effects](#shadows){#drawing-rectangles-to-the-bitmap:shadows}, [global
alpha](#concept-canvas-global-alpha){#drawing-rectangles-to-the-bitmap:concept-canvas-global-alpha},
and the [current compositing and blending
operator](#current-compositing-and-blending-operator){#drawing-rectangles-to-the-bitmap:current-compositing-and-blending-operator}.

`context`{.variable}`.`[`clearRect`](#dom-context-2d-clearrect){#dom-context-2d-clearrect-dev}`(``x`{.variable}`, ``y`{.variable}`, ``w`{.variable}`, ``h`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/clearRect](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/clearRect "The CanvasRenderingContext2D.clearRect() method of the Canvas 2D API erases the pixels in a rectangular area by setting them to transparent black.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari2+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Clears all pixels on the bitmap in the given rectangle to [transparent
black](https://drafts.csswg.org/css-color/#transparent-black){#drawing-rectangles-to-the-bitmap:transparent-black
x-internal="transparent-black"}.

`context`{.variable}`.`[`fillRect`](#dom-context-2d-fillrect){#dom-context-2d-fillrect-dev}`(``x`{.variable}`, ``y`{.variable}`, ``w`{.variable}`, ``h`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/fillRect](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/fillRect "The CanvasRenderingContext2D.fillRect() method of the Canvas 2D API draws a rectangle that is filled according to the current fillStyle.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari2+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera9+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

Paints the given rectangle onto the bitmap, using the current fill
style.

`context`{.variable}`.`[`strokeRect`](#dom-context-2d-strokerect){#dom-context-2d-strokerect-dev}`(``x`{.variable}`, ``y`{.variable}`, ``w`{.variable}`, ``h`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/strokeRect](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/strokeRect "The CanvasRenderingContext2D.strokeRect() method of the Canvas 2D API draws a rectangle that is stroked (outlined) according to the current strokeStyle and other context settings.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari2+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Paints the box that outlines the given rectangle onto the bitmap, using
the current stroke style.

The
[`clearRect(``x`{.variable}`, ``y`{.variable}`, ``w`{.variable}`, ``h`{.variable}`)`]{#dom-context-2d-clearrect
.dfn dfn-for="CanvasRect" dfn-type="method"} method, when invoked, must
run these steps:

1.  If any of the arguments are infinite or NaN, then return.

2.  Let `pixels`{.variable} be the set of pixels in the specified
    rectangle that also intersect the current [clipping
    region](#clipping-region){#drawing-rectangles-to-the-bitmap:clipping-region-2}.

3.  Clear the pixels in `pixels`{.variable} to a [transparent
    black](https://drafts.csswg.org/css-color/#transparent-black){#drawing-rectangles-to-the-bitmap:transparent-black-2
    x-internal="transparent-black"}, erasing any previous image.

If either height or width are zero, this method has no effect, since the
set of pixels would be empty.

The
[`fillRect(``x`{.variable}`, ``y`{.variable}`, ``w`{.variable}`, ``h`{.variable}`)`]{#dom-context-2d-fillrect
.dfn dfn-for="CanvasRect" dfn-type="method"} method, when invoked, must
run these steps:

1.  If any of the arguments are infinite or NaN, then return.

2.  If either `w`{.variable} or `h`{.variable} are zero, then return.

3.  Paint the specified rectangular area using
    [this](https://webidl.spec.whatwg.org/#this){#drawing-rectangles-to-the-bitmap:this
    x-internal="this"}\'s [fill
    style](#concept-canvasfillstrokestyles-fill-style){#drawing-rectangles-to-the-bitmap:concept-canvasfillstrokestyles-fill-style}.

The
[`strokeRect(``x`{.variable}`, ``y`{.variable}`, ``w`{.variable}`, ``h`{.variable}`)`]{#dom-context-2d-strokerect
.dfn dfn-for="CanvasRect" dfn-type="method"} method, when invoked, must
run these steps:

1.  If any of the arguments are infinite or NaN, then return.

2.  Take the result of [tracing the
    path](#trace-a-path){#drawing-rectangles-to-the-bitmap:trace-a-path}
    described below, using the
    [`CanvasPathDrawingStyles`{#drawing-rectangles-to-the-bitmap:canvaspathdrawingstyles}](#canvaspathdrawingstyles)
    interface\'s line styles, and fill it with
    [this](https://webidl.spec.whatwg.org/#this){#drawing-rectangles-to-the-bitmap:this-2
    x-internal="this"}\'s [stroke
    style](#concept-canvasfillstrokestyles-stroke-style){#drawing-rectangles-to-the-bitmap:concept-canvasfillstrokestyles-stroke-style}.

If both `w`{.variable} and `h`{.variable} are zero, the path has a
single subpath with just one point (`x`{.variable}, `y`{.variable}), and
no lines, and this method thus has no effect (the [trace a
path](#trace-a-path){#drawing-rectangles-to-the-bitmap:trace-a-path-2}
algorithm returns an empty path in that case).

If just one of either `w`{.variable} or `h`{.variable} is zero, then the
path has a single subpath consisting of two points, with coordinates
(`x`{.variable}, `y`{.variable}) and (`x`{.variable}+`w`{.variable},
`y`{.variable}+`h`{.variable}), in that order, connected by a single
straight line.

Otherwise, the path has a single subpath consisting of four points, with
coordinates (`x`{.variable}, `y`{.variable}),
(`x`{.variable}+`w`{.variable}, `y`{.variable}),
(`x`{.variable}+`w`{.variable}, `y`{.variable}+`h`{.variable}), and
(`x`{.variable}, `y`{.variable}+`h`{.variable}), connected to each other
in that order by straight lines.

###### [4.12.5.1.12]{.secno} Drawing text to the bitmap[](#drawing-text-to-the-bitmap){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[CanvasRenderingContext2D](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D "The CanvasRenderingContext2D interface, part of the Canvas API, provides the 2D rendering context for the drawing surface of a <canvas> element. It is used for drawing shapes, text, images, and other objects.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari2+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera9+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

`context`{.variable}`.`[`fillText`](#dom-context-2d-filltext){#dom-context-2d-filltext-dev}`(``text`{.variable}`, ``x`{.variable}`, ``y`{.variable}` [, ``maxWidth`{.variable}` ])`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/fillText](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/fillText "The CanvasRenderingContext2D method fillText(), part of the Canvas 2D API, draws a text string at the specified coordinates, filling the string's characters with the current fillStyle. An optional parameter allows specifying a maximum width for the rendered text, which the user agent will achieve by condensing the text or by using a lower font size.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome2+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

`context`{.variable}`.`[`strokeText`](#dom-context-2d-stroketext){#dom-context-2d-stroketext-dev}`(``text`{.variable}`, ``x`{.variable}`, ``y`{.variable}` [, ``maxWidth`{.variable}` ])`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/strokeText](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/strokeText "The CanvasRenderingContext2D method strokeText(), part of the Canvas 2D API, strokes — that is, draws the outlines of — the characters of a text string at the specified coordinates. An optional parameter allows specifying a maximum width for the rendered text, which the user agent will achieve by condensing the text or by using a lower font size.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome2+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Fills or strokes (respectively) the given text at the given position. If
a maximum width is provided, the text will be scaled to fit that width
if necessary.

`metrics`{.variable}` = ``context`{.variable}`.`[`measureText`](#dom-context-2d-measuretext){#dom-context-2d-measuretext-dev}`(``text`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/measureText](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/measureText "The CanvasRenderingContext2D.measureText() method returns a TextMetrics object that contains information about the measured text (such as its width, for example).")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome2+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns a
[`TextMetrics`{#drawing-text-to-the-bitmap:textmetrics}](#textmetrics)
object with the metrics of the given text in the current font.

`metrics`{.variable}`.`[`width`](#dom-textmetrics-width){#dom-textmetrics-width-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[TextMetrics/width](https://developer.mozilla.org/en-US/docs/Web/API/TextMetrics/width "The read-only width property of the TextMetrics interface contains the text's advance width (the width of that inline box) in CSS pixels.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome2+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera9+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android31+]{.firefox_android .yes}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

`metrics`{.variable}`.`[`actualBoundingBoxLeft`](#dom-textmetrics-actualboundingboxleft){#dom-textmetrics-actualboundingboxleft-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[TextMetrics/actualBoundingBoxLeft](https://developer.mozilla.org/en-US/docs/Web/API/TextMetrics/actualBoundingBoxLeft "The read-only actualBoundingBoxLeft property of the TextMetrics interface is a double giving the distance parallel to the baseline from the alignment point given by the CanvasRenderingContext2D.textAlign property to the left side of the bounding rectangle of the given text, in CSS pixels; positive numbers indicating a distance going left from the given alignment point.")

Support in all current engines.

::: support
[Firefox74+]{.firefox .yes}[Safari11.1+]{.safari
.yes}[Chrome77+]{.chrome .yes}

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

`metrics`{.variable}`.`[`actualBoundingBoxRight`](#dom-textmetrics-actualboundingboxright){#dom-textmetrics-actualboundingboxright-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[TextMetrics/actualBoundingBoxRight](https://developer.mozilla.org/en-US/docs/Web/API/TextMetrics/actualBoundingBoxRight "The read-only actualBoundingBoxRight property of the TextMetrics interface is a double giving the distance parallel to the baseline from the alignment point given by the CanvasRenderingContext2D.textAlign property to the right side of the bounding rectangle of the given text, in CSS pixels.")

Support in all current engines.

::: support
[Firefox74+]{.firefox .yes}[Safari11.1+]{.safari
.yes}[Chrome77+]{.chrome .yes}

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

`metrics`{.variable}`.`[`fontBoundingBoxAscent`](#dom-textmetrics-fontboundingboxascent){#dom-textmetrics-fontboundingboxascent-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[TextMetrics/fontBoundingBoxAscent](https://developer.mozilla.org/en-US/docs/Web/API/TextMetrics/fontBoundingBoxAscent "The read-only fontBoundingBoxAscent property of the TextMetrics interface returns the distance from the horizontal line indicated by the CanvasRenderingContext2D.textBaseline attribute, to the top of the highest bounding rectangle of all the fonts used to render the text, in CSS pixels.")

Support in all current engines.

::: support
[Firefox116+]{.firefox .yes}[Safari11.1+]{.safari
.yes}[Chrome87+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge87+]{.edge_blink .yes}

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

`metrics`{.variable}`.`[`fontBoundingBoxDescent`](#dom-textmetrics-fontboundingboxdescent){#dom-textmetrics-fontboundingboxdescent-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[TextMetrics/fontBoundingBoxDescent](https://developer.mozilla.org/en-US/docs/Web/API/TextMetrics/fontBoundingBoxDescent "The read-only fontBoundingBoxDescent property of the TextMetrics interface returns the distance from the horizontal line indicated by the CanvasRenderingContext2D.textBaseline attribute to the bottom of the bounding rectangle of all the fonts used to render the text, in CSS pixels.")

Support in all current engines.

::: support
[Firefox116+]{.firefox .yes}[Safari11.1+]{.safari
.yes}[Chrome87+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge87+]{.edge_blink .yes}

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

`metrics`{.variable}`.`[`actualBoundingBoxAscent`](#dom-textmetrics-actualboundingboxascent){#dom-textmetrics-actualboundingboxascent-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[TextMetrics/actualBoundingBoxAscent](https://developer.mozilla.org/en-US/docs/Web/API/TextMetrics/actualBoundingBoxAscent "The read-only actualBoundingBoxAscent property of the TextMetrics interface is a double giving the distance from the horizontal line indicated by the CanvasRenderingContext2D.textBaseline attribute to the top of the bounding rectangle used to render the text, in CSS pixels.")

Support in all current engines.

::: support
[Firefox74+]{.firefox .yes}[Safari11.1+]{.safari
.yes}[Chrome77+]{.chrome .yes}

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

`metrics`{.variable}`.`[`actualBoundingBoxDescent`](#dom-textmetrics-actualboundingboxdescent){#dom-textmetrics-actualboundingboxdescent-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[TextMetrics/actualBoundingBoxDescent](https://developer.mozilla.org/en-US/docs/Web/API/TextMetrics/actualBoundingBoxDescent "The read-only actualBoundingBoxDescent property of the TextMetrics interface is a double giving the distance from the horizontal line indicated by the CanvasRenderingContext2D.textBaseline attribute to the bottom of the bounding rectangle used to render the text, in CSS pixels.")

Support in all current engines.

::: support
[Firefox74+]{.firefox .yes}[Safari11.1+]{.safari
.yes}[Chrome77+]{.chrome .yes}

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

`metrics`{.variable}`.`[`emHeightAscent`](#dom-textmetrics-emheightascent){#dom-textmetrics-emheightascent-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[TextMetrics/emHeightAscent](https://developer.mozilla.org/en-US/docs/Web/API/TextMetrics/emHeightAscent "The read-only emHeightAscent property of the TextMetrics interface is a double giving the distance from the horizontal line indicated by the CanvasRenderingContext2D.textBaseline property to the top of the em square in the line box, in CSS pixels.")

Support in all current engines.

::: support
[Firefox[🔰
74+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari11.1+]{.safari .yes}[Chrome[🔰
35+]{title="Requires setting a user preference or runtime flag."}]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge[🔰
79+]{title="Requires setting a user preference or runtime flag."}]{.edge_blink
.yes}

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

`metrics`{.variable}`.`[`emHeightDescent`](#dom-textmetrics-emheightdescent){#dom-textmetrics-emheightdescent-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[TextMetrics/emHeightDescent](https://developer.mozilla.org/en-US/docs/Web/API/TextMetrics/emHeightDescent "The read-only emHeightDescent property of the TextMetrics interface is a double giving the distance from the horizontal line indicated by the CanvasRenderingContext2D.textBaseline property to the bottom of the em square in the line box, in CSS pixels.")

Support in all current engines.

::: support
[Firefox[🔰
74+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari11.1+]{.safari .yes}[Chrome[🔰
35+]{title="Requires setting a user preference or runtime flag."}]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge[🔰
79+]{title="Requires setting a user preference or runtime flag."}]{.edge_blink
.yes}

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

`metrics`{.variable}`.`[`hangingBaseline`](#dom-textmetrics-hangingbaseline){#dom-textmetrics-hangingbaseline-dev}

::::: {.mdn-anno .wrapped .before}
MDN

:::: feature
[TextMetrics/hangingBaseline](https://developer.mozilla.org/en-US/docs/Web/API/TextMetrics/hangingBaseline "The read-only hangingBaseline property of the TextMetrics interface is a double giving the distance from the horizontal line indicated by the CanvasRenderingContext2D.textBaseline property to the hanging baseline of the line box, in CSS pixels.")

::: support
[Firefox[🔰
74+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari11.1+]{.safari .yes}[ChromeNo]{.chrome .no}

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

`metrics`{.variable}`.`[`alphabeticBaseline`](#dom-textmetrics-alphabeticbaseline){#dom-textmetrics-alphabeticbaseline-dev}

::::: {.mdn-anno .wrapped .before}
MDN

:::: feature
[TextMetrics/alphabeticBaseline](https://developer.mozilla.org/en-US/docs/Web/API/TextMetrics/alphabeticBaseline "The read-only alphabeticBaseline property of the TextMetrics interface is a double giving the distance from the horizontal line indicated by the CanvasRenderingContext2D.textBaseline property to the alphabetic baseline of the line box, in CSS pixels.")

::: support
[Firefox[🔰
74+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari11.1+]{.safari .yes}[ChromeNo]{.chrome .no}

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

`metrics`{.variable}`.`[`ideographicBaseline`](#dom-textmetrics-ideographicbaseline){#dom-textmetrics-ideographicbaseline-dev}

::::: {.mdn-anno .wrapped .before}
MDN

:::: feature
[TextMetrics/ideographicBaseline](https://developer.mozilla.org/en-US/docs/Web/API/TextMetrics/ideographicBaseline "The read-only ideographicBaseline property of the TextMetrics interface is a double giving the distance from the horizontal line indicated by the CanvasRenderingContext2D.textBaseline property to the ideographic baseline of the line box, in CSS pixels.")

::: support
[Firefox[🔰
74+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari11.1+]{.safari .yes}[ChromeNo]{.chrome .no}

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

Returns the measurement described below.

Objects that implement the
[`CanvasText`{#drawing-text-to-the-bitmap:canvastext}](#canvastext)
interface provide the following methods for rendering text.

The
[`fillText(``text`{.variable}`, ``x`{.variable}`, ``y`{.variable}`, ``maxWidth`{.variable}`)`]{#dom-context-2d-filltext
.dfn dfn-for="CanvasText" dfn-type="method"} and
[`strokeText(``text`{.variable}`, ``x`{.variable}`, ``y`{.variable}`, ``maxWidth`{.variable}`)`]{#dom-context-2d-stroketext
.dfn dfn-for="CanvasText" dfn-type="method"} methods render the given
`text`{.variable} at the given (`x`{.variable}, `y`{.variable})
coordinates ensuring that the text isn\'t wider than
`maxWidth`{.variable} if specified, using the current
[`font`{#drawing-text-to-the-bitmap:dom-context-2d-font}](#dom-context-2d-font),
[`textAlign`{#drawing-text-to-the-bitmap:dom-context-2d-textalign}](#dom-context-2d-textalign),
and
[`textBaseline`{#drawing-text-to-the-bitmap:dom-context-2d-textbaseline}](#dom-context-2d-textbaseline)
values. Specifically, when the methods are invoked, the user agent must
run these steps:

1.  If any of the arguments are infinite or NaN, then return.

2.  Run the [text preparation
    algorithm](#text-preparation-algorithm){#drawing-text-to-the-bitmap:text-preparation-algorithm},
    passing it `text`{.variable}, the object implementing the
    [`CanvasText`{#drawing-text-to-the-bitmap:canvastext-2}](#canvastext)
    interface, and, if the `maxWidth`{.variable} argument was provided,
    that argument. Let `glyphs`{.variable} be the result.

3.  Move all the shapes in `glyphs`{.variable} to the right by
    `x`{.variable} [CSS
    pixels](https://drafts.csswg.org/css-values/#px){#drawing-text-to-the-bitmap:'px'
    x-internal="'px'"} and down by `y`{.variable} [CSS
    pixels](https://drafts.csswg.org/css-values/#px){#drawing-text-to-the-bitmap:'px'-2
    x-internal="'px'"}.

4.  Paint the shapes given in `glyphs`{.variable}, as transformed by the
    [current transformation
    matrix](#transformations){#drawing-text-to-the-bitmap:transformations},
    with each [CSS
    pixel](https://drafts.csswg.org/css-values/#px){#drawing-text-to-the-bitmap:'px'-3
    x-internal="'px'"} in the coordinate space of `glyphs`{.variable}
    mapped to one coordinate space unit.

    For
    [`fillText()`{#drawing-text-to-the-bitmap:dom-context-2d-filltext}](#dom-context-2d-filltext),
    [this](https://webidl.spec.whatwg.org/#this){#drawing-text-to-the-bitmap:this
    x-internal="this"}\'s [fill
    style](#concept-canvasfillstrokestyles-fill-style){#drawing-text-to-the-bitmap:concept-canvasfillstrokestyles-fill-style}
    must be applied to the shapes and
    [this](https://webidl.spec.whatwg.org/#this){#drawing-text-to-the-bitmap:this-2
    x-internal="this"}\'s [stroke
    style](#concept-canvasfillstrokestyles-stroke-style){#drawing-text-to-the-bitmap:concept-canvasfillstrokestyles-stroke-style}
    must be ignored. For
    [`strokeText()`{#drawing-text-to-the-bitmap:dom-context-2d-stroketext}](#dom-context-2d-stroketext),
    the reverse holds:
    [this](https://webidl.spec.whatwg.org/#this){#drawing-text-to-the-bitmap:this-3
    x-internal="this"}\'s [stroke
    style](#concept-canvasfillstrokestyles-stroke-style){#drawing-text-to-the-bitmap:concept-canvasfillstrokestyles-stroke-style-2}
    must be applied to the result of
    [tracing](#trace-a-path){#drawing-text-to-the-bitmap:trace-a-path}
    the shapes using the object implementing the
    [`CanvasText`{#drawing-text-to-the-bitmap:canvastext-3}](#canvastext)
    interface for the line styles, and
    [this](https://webidl.spec.whatwg.org/#this){#drawing-text-to-the-bitmap:this-4
    x-internal="this"}\'s [fill
    style](#concept-canvasfillstrokestyles-fill-style){#drawing-text-to-the-bitmap:concept-canvasfillstrokestyles-fill-style-2}
    must be ignored.

    These shapes are painted without affecting the current path, and are
    subject to [shadow
    effects](#shadows){#drawing-text-to-the-bitmap:shadows}, [global
    alpha](#concept-canvas-global-alpha){#drawing-text-to-the-bitmap:concept-canvas-global-alpha},
    the [clipping
    region](#clipping-region){#drawing-text-to-the-bitmap:clipping-region},
    and the [current compositing and blending
    operator](#current-compositing-and-blending-operator){#drawing-text-to-the-bitmap:current-compositing-and-blending-operator}.

[![(This is a tracking
vector.)](https://resources.whatwg.org/tracking-vector.svg){width="46"
crossorigin=""
height="64"}](https://infra.spec.whatwg.org/#tracking-vector "There is a tracking vector here."){#drawing-text-to-the-bitmap:tracking-vector
.tracking-vector x-internal="tracking-vector"} The
[`measureText(``text`{.variable}`)`]{#dom-context-2d-measuretext .dfn
dfn-for="CanvasText" dfn-type="method"} method steps are to run the
[text preparation
algorithm](#text-preparation-algorithm){#drawing-text-to-the-bitmap:text-preparation-algorithm-2},
passing it `text`{.variable} and the object implementing the
[`CanvasText`{#drawing-text-to-the-bitmap:canvastext-4}](#canvastext)
interface, and then using the returned [inline
box](https://drafts.csswg.org/css2/#inline-box){#drawing-text-to-the-bitmap:inline-box
x-internal="inline-box"} return a new
[`TextMetrics`{#drawing-text-to-the-bitmap:textmetrics-2}](#textmetrics)
object with members behaving as described in the following list:
[\[CSS\]](references.html#refsCSS)

[`width`]{#dom-textmetrics-width .dfn dfn-for="TextMetrics" dfn-type="attribute"} attribute
:   The width of that [inline
    box](https://drafts.csswg.org/css2/#inline-box){#drawing-text-to-the-bitmap:inline-box-2
    x-internal="inline-box"}, in [CSS
    pixels](https://drafts.csswg.org/css-values/#px){#drawing-text-to-the-bitmap:'px'-4
    x-internal="'px'"}. (The text\'s advance width.)

[`actualBoundingBoxLeft`]{#dom-textmetrics-actualboundingboxleft .dfn dfn-for="TextMetrics" dfn-type="attribute"} attribute

:   The distance parallel to the baseline from the alignment point given
    by the
    [`textAlign`{#drawing-text-to-the-bitmap:dom-context-2d-textalign-2}](#dom-context-2d-textalign)
    attribute to the left side of the bounding rectangle of the given
    text, in [CSS
    pixels](https://drafts.csswg.org/css-values/#px){#drawing-text-to-the-bitmap:'px'-5
    x-internal="'px'"}; positive numbers indicating a distance going
    left from the given alignment point.

    The sum of this value and the next
    ([`actualBoundingBoxRight`{#drawing-text-to-the-bitmap:dom-textmetrics-actualboundingboxright}](#dom-textmetrics-actualboundingboxright))
    can be wider than the width of the [inline
    box](https://drafts.csswg.org/css2/#inline-box){#drawing-text-to-the-bitmap:inline-box-3
    x-internal="inline-box"}
    ([`width`{#drawing-text-to-the-bitmap:dom-textmetrics-width}](#dom-textmetrics-width)),
    in particular with slanted fonts where characters overhang their
    advance width.

[`actualBoundingBoxRight`]{#dom-textmetrics-actualboundingboxright .dfn dfn-for="TextMetrics" dfn-type="attribute"} attribute

:   The distance parallel to the baseline from the alignment point given
    by the
    [`textAlign`{#drawing-text-to-the-bitmap:dom-context-2d-textalign-3}](#dom-context-2d-textalign)
    attribute to the right side of the bounding rectangle of the given
    text, in [CSS
    pixels](https://drafts.csswg.org/css-values/#px){#drawing-text-to-the-bitmap:'px'-6
    x-internal="'px'"}; positive numbers indicating a distance going
    right from the given alignment point.

[`fontBoundingBoxAscent`]{#dom-textmetrics-fontboundingboxascent .dfn dfn-for="TextMetrics" dfn-type="attribute"} attribute

:   The distance from the horizontal line indicated by the
    [`textBaseline`{#drawing-text-to-the-bitmap:dom-context-2d-textbaseline-2}](#dom-context-2d-textbaseline)
    attribute to the [ascent
    metric](https://drafts.csswg.org/css-inline/#ascent-metric){#drawing-text-to-the-bitmap:ascent-metric
    x-internal="ascent-metric"} of the [first available
    font](https://drafts.csswg.org/css-fonts/#first-available-font){#drawing-text-to-the-bitmap:first-available-font
    x-internal="first-available-font"}, in [CSS
    pixels](https://drafts.csswg.org/css-values/#px){#drawing-text-to-the-bitmap:'px'-7
    x-internal="'px'"}; positive numbers indicating a distance going up
    from the given baseline.

    This value and the next are useful when rendering a background that
    have to have a consistent height even if the exact text being
    rendered changes. The
    [`actualBoundingBoxAscent`{#drawing-text-to-the-bitmap:dom-textmetrics-actualboundingboxascent}](#dom-textmetrics-actualboundingboxascent)
    attribute (and its corresponding attribute for the descent) are
    useful when drawing a bounding box around specific text.

[`fontBoundingBoxDescent`]{#dom-textmetrics-fontboundingboxdescent .dfn dfn-for="TextMetrics" dfn-type="attribute"} attribute
:   The distance from the horizontal line indicated by the
    [`textBaseline`{#drawing-text-to-the-bitmap:dom-context-2d-textbaseline-3}](#dom-context-2d-textbaseline)
    attribute to the [descent
    metric](https://drafts.csswg.org/css-inline/#descent-metric){#drawing-text-to-the-bitmap:descent-metric
    x-internal="descent-metric"} of the [first available
    font](https://drafts.csswg.org/css-fonts/#first-available-font){#drawing-text-to-the-bitmap:first-available-font-2
    x-internal="first-available-font"}, in [CSS
    pixels](https://drafts.csswg.org/css-values/#px){#drawing-text-to-the-bitmap:'px'-8
    x-internal="'px'"}; positive numbers indicating a distance going
    down from the given baseline.

[`actualBoundingBoxAscent`]{#dom-textmetrics-actualboundingboxascent .dfn dfn-for="TextMetrics" dfn-type="attribute"} attribute

:   The distance from the horizontal line indicated by the
    [`textBaseline`{#drawing-text-to-the-bitmap:dom-context-2d-textbaseline-4}](#dom-context-2d-textbaseline)
    attribute to the top of the bounding rectangle of the given text, in
    [CSS
    pixels](https://drafts.csswg.org/css-values/#px){#drawing-text-to-the-bitmap:'px'-9
    x-internal="'px'"}; positive numbers indicating a distance going up
    from the given baseline.

    This number can vary greatly based on the input text, even if the
    first font specified covers all the characters in the input. For
    example, the
    [`actualBoundingBoxAscent`{#drawing-text-to-the-bitmap:dom-textmetrics-actualboundingboxascent-2}](#dom-textmetrics-actualboundingboxascent)
    of a lowercase \"o\" from an [alphabetic
    baseline](https://drafts.csswg.org/css-inline/#alphabetic-baseline){#drawing-text-to-the-bitmap:alphabetic-baseline
    x-internal="alphabetic-baseline"} would be less than that of an
    uppercase \"F\". The value can easily be negative; for example, the
    distance from the top of the em box
    ([`textBaseline`{#drawing-text-to-the-bitmap:dom-context-2d-textbaseline-5}](#dom-context-2d-textbaseline)
    value
    \"[`top`{#drawing-text-to-the-bitmap:dom-context-2d-textbaseline-top}](#dom-context-2d-textbaseline-top)\")
    to the top of the bounding rectangle when the given text is just a
    single comma \"`,`\" would likely (unless the font is quite unusual)
    be negative.

[`actualBoundingBoxDescent`]{#dom-textmetrics-actualboundingboxdescent .dfn dfn-for="TextMetrics" dfn-type="attribute"} attribute
:   The distance from the horizontal line indicated by the
    [`textBaseline`{#drawing-text-to-the-bitmap:dom-context-2d-textbaseline-6}](#dom-context-2d-textbaseline)
    attribute to the bottom of the bounding rectangle of the given text,
    in [CSS
    pixels](https://drafts.csswg.org/css-values/#px){#drawing-text-to-the-bitmap:'px'-10
    x-internal="'px'"}; positive numbers indicating a distance going
    down from the given baseline.

[`emHeightAscent`]{#dom-textmetrics-emheightascent .dfn dfn-for="TextMetrics" dfn-type="attribute"} attribute
:   The distance from the horizontal line indicated by the
    [`textBaseline`{#drawing-text-to-the-bitmap:dom-context-2d-textbaseline-7}](#dom-context-2d-textbaseline)
    attribute to the [em-over
    baseline](https://drafts.csswg.org/css-inline/#em-over-baseline){#drawing-text-to-the-bitmap:em-over-baseline
    x-internal="em-over-baseline"} in the [inline
    box](https://drafts.csswg.org/css2/#inline-box){#drawing-text-to-the-bitmap:inline-box-4
    x-internal="inline-box"}, in [CSS
    pixels](https://drafts.csswg.org/css-values/#px){#drawing-text-to-the-bitmap:'px'-11
    x-internal="'px'"}; positive numbers indicating that the given
    baseline is below the [em-over
    baseline](https://drafts.csswg.org/css-inline/#em-over-baseline){#drawing-text-to-the-bitmap:em-over-baseline-2
    x-internal="em-over-baseline"} (so this value will usually be
    positive). Zero if the given baseline is the [em-over
    baseline](https://drafts.csswg.org/css-inline/#em-over-baseline){#drawing-text-to-the-bitmap:em-over-baseline-3
    x-internal="em-over-baseline"}; half the font size if the given
    baseline is halfway between the [em-over
    baseline](https://drafts.csswg.org/css-inline/#em-over-baseline){#drawing-text-to-the-bitmap:em-over-baseline-4
    x-internal="em-over-baseline"} and the [em-under
    baseline](https://drafts.csswg.org/css-inline/#em-under-baseline){#drawing-text-to-the-bitmap:em-under-baseline
    x-internal="em-under-baseline"}.

[`emHeightDescent`]{#dom-textmetrics-emheightdescent .dfn dfn-for="TextMetrics" dfn-type="attribute"} attribute
:   The distance from the horizontal line indicated by the
    [`textBaseline`{#drawing-text-to-the-bitmap:dom-context-2d-textbaseline-8}](#dom-context-2d-textbaseline)
    attribute to the [em-under
    baseline](https://drafts.csswg.org/css-inline/#em-under-baseline){#drawing-text-to-the-bitmap:em-under-baseline-2
    x-internal="em-under-baseline"} in the [inline
    box](https://drafts.csswg.org/css2/#inline-box){#drawing-text-to-the-bitmap:inline-box-5
    x-internal="inline-box"}, in [CSS
    pixels](https://drafts.csswg.org/css-values/#px){#drawing-text-to-the-bitmap:'px'-12
    x-internal="'px'"}; positive numbers indicating that the given
    baseline is above the [em-under
    baseline](https://drafts.csswg.org/css-inline/#em-under-baseline){#drawing-text-to-the-bitmap:em-under-baseline-3
    x-internal="em-under-baseline"}. (Zero if the given baseline is the
    [em-under
    baseline](https://drafts.csswg.org/css-inline/#em-under-baseline){#drawing-text-to-the-bitmap:em-under-baseline-4
    x-internal="em-under-baseline"}.)

[`hangingBaseline`]{#dom-textmetrics-hangingbaseline .dfn dfn-for="TextMetrics" dfn-type="attribute"} attribute
:   The distance from the horizontal line indicated by the
    [`textBaseline`{#drawing-text-to-the-bitmap:dom-context-2d-textbaseline-9}](#dom-context-2d-textbaseline)
    attribute to the [hanging
    baseline](https://drafts.csswg.org/css-inline/#hanging-baseline){#drawing-text-to-the-bitmap:hanging-baseline
    x-internal="hanging-baseline"} of the [inline
    box](https://drafts.csswg.org/css2/#inline-box){#drawing-text-to-the-bitmap:inline-box-6
    x-internal="inline-box"}, in [CSS
    pixels](https://drafts.csswg.org/css-values/#px){#drawing-text-to-the-bitmap:'px'-13
    x-internal="'px'"}; positive numbers indicating that the given
    baseline is below the [hanging
    baseline](https://drafts.csswg.org/css-inline/#hanging-baseline){#drawing-text-to-the-bitmap:hanging-baseline-2
    x-internal="hanging-baseline"}. (Zero if the given baseline is the
    [hanging
    baseline](https://drafts.csswg.org/css-inline/#hanging-baseline){#drawing-text-to-the-bitmap:hanging-baseline-3
    x-internal="hanging-baseline"}.)

[`alphabeticBaseline`]{#dom-textmetrics-alphabeticbaseline .dfn dfn-for="TextMetrics" dfn-type="attribute"} attribute
:   The distance from the horizontal line indicated by the
    [`textBaseline`{#drawing-text-to-the-bitmap:dom-context-2d-textbaseline-10}](#dom-context-2d-textbaseline)
    attribute to the [alphabetic
    baseline](https://drafts.csswg.org/css-inline/#alphabetic-baseline){#drawing-text-to-the-bitmap:alphabetic-baseline-2
    x-internal="alphabetic-baseline"} of the [inline
    box](https://drafts.csswg.org/css2/#inline-box){#drawing-text-to-the-bitmap:inline-box-7
    x-internal="inline-box"}, in [CSS
    pixels](https://drafts.csswg.org/css-values/#px){#drawing-text-to-the-bitmap:'px'-14
    x-internal="'px'"}; positive numbers indicating that the given
    baseline is below the [alphabetic
    baseline](https://drafts.csswg.org/css-inline/#alphabetic-baseline){#drawing-text-to-the-bitmap:alphabetic-baseline-3
    x-internal="alphabetic-baseline"}. (Zero if the given baseline is
    the [alphabetic
    baseline](https://drafts.csswg.org/css-inline/#alphabetic-baseline){#drawing-text-to-the-bitmap:alphabetic-baseline-4
    x-internal="alphabetic-baseline"}.)

[`ideographicBaseline`]{#dom-textmetrics-ideographicbaseline .dfn dfn-for="TextMetrics" dfn-type="attribute"} attribute

:   The distance from the horizontal line indicated by the
    [`textBaseline`{#drawing-text-to-the-bitmap:dom-context-2d-textbaseline-11}](#dom-context-2d-textbaseline)
    attribute to the [ideographic-under
    baseline](https://drafts.csswg.org/css-inline/#ideographic-under-baseline){#drawing-text-to-the-bitmap:ideographic-under-baseline
    x-internal="ideographic-under-baseline"} of the [inline
    box](https://drafts.csswg.org/css2/#inline-box){#drawing-text-to-the-bitmap:inline-box-8
    x-internal="inline-box"}, in [CSS
    pixels](https://drafts.csswg.org/css-values/#px){#drawing-text-to-the-bitmap:'px'-15
    x-internal="'px'"}; positive numbers indicating that the given
    baseline is below the [ideographic-under
    baseline](https://drafts.csswg.org/css-inline/#ideographic-under-baseline){#drawing-text-to-the-bitmap:ideographic-under-baseline-2
    x-internal="ideographic-under-baseline"}. (Zero if the given
    baseline is the [ideographic-under
    baseline](https://drafts.csswg.org/css-inline/#ideographic-under-baseline){#drawing-text-to-the-bitmap:ideographic-under-baseline-3
    x-internal="ideographic-under-baseline"}.)

Glyphs rendered using
[`fillText()`{#drawing-text-to-the-bitmap:dom-context-2d-filltext-2}](#dom-context-2d-filltext)
and
[`strokeText()`{#drawing-text-to-the-bitmap:dom-context-2d-stroketext-2}](#dom-context-2d-stroketext)
can spill out of the box given by the font size and the width returned
by
[`measureText()`{#drawing-text-to-the-bitmap:dom-context-2d-measuretext}](#dom-context-2d-measuretext)
(the text width). Authors are encouraged to use the bounding box values
described above if this is an issue.

A future version of the 2D context API might provide a way to render
fragments of documents, rendered using CSS, straight to the canvas. This
would be provided in preference to a dedicated way of doing multiline
layout.

###### [4.12.5.1.13]{.secno} Drawing paths to the canvas[](#drawing-paths-to-the-canvas){.self-link}

Objects that implement the
[`CanvasDrawPath`{#drawing-paths-to-the-canvas:canvasdrawpath}](#canvasdrawpath)
interface have a [current default path]{#current-default-path .dfn}.
There is only one [current default
path](#current-default-path){#drawing-paths-to-the-canvas:current-default-path},
it is not part of the [drawing
state](#drawing-state){#drawing-paths-to-the-canvas:drawing-state}. The
[current default
path](#current-default-path){#drawing-paths-to-the-canvas:current-default-path-2}
is a [path](#concept-path){#drawing-paths-to-the-canvas:concept-path},
as described above.

`context`{.variable}`.`[`beginPath`](#dom-context-2d-beginpath){#dom-context-2d-beginpath-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/beginPath](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/beginPath "The CanvasRenderingContext2D.beginPath() method of the Canvas 2D API starts a new path by emptying the list of sub-paths. Call this method when you want to create a new path.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari2+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Resets the [current default
path](#current-default-path){#drawing-paths-to-the-canvas:current-default-path-3}.

`context`{.variable}`.`[`fill`](#dom-context-2d-fill){#dom-context-2d-fill-dev}`([ ``fillRule`{.variable}` ])`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/fill](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/fill "The CanvasRenderingContext2D.fill() method of the Canvas 2D API fills the current or given path with the current fillStyle.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari2+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

`context`{.variable}`.`[`fill`](#dom-context-2d-fill){#drawing-paths-to-the-canvas:dom-context-2d-fill}`(``path`{.variable}` [, ``fillRule`{.variable}` ])`

Fills the subpaths of the [current default
path](#current-default-path){#drawing-paths-to-the-canvas:current-default-path-4}
or the given path with the current fill style, obeying the given fill
rule.

`context`{.variable}`.`[`stroke`](#dom-context-2d-stroke){#dom-context-2d-stroke-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/stroke](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/stroke "The CanvasRenderingContext2D.stroke() method of the Canvas 2D API strokes (outlines) the current or given path with the current stroke style.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari2+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

`context`{.variable}`.`[`stroke`](#dom-context-2d-stroke){#drawing-paths-to-the-canvas:dom-context-2d-stroke}`(``path`{.variable}`)`

Strokes the subpaths of the [current default
path](#current-default-path){#drawing-paths-to-the-canvas:current-default-path-5}
or the given path with the current stroke style.

`context`{.variable}`.`[`clip`](#dom-context-2d-clip){#dom-context-2d-clip-dev}`([ ``fillRule`{.variable}` ])`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/clip](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/clip "The CanvasRenderingContext2D.clip() method of the Canvas 2D API turns the current or given path into the current clipping region. The previous clipping region, if any, is intersected with the current or given path to create the new clipping region.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari2+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

`context`{.variable}`.`[`clip`](#dom-context-2d-clip){#drawing-paths-to-the-canvas:dom-context-2d-clip}`(``path`{.variable}` [, ``fillRule`{.variable}` ])`

Further constrains the clipping region to the [current default
path](#current-default-path){#drawing-paths-to-the-canvas:current-default-path-6}
or the given path, using the given fill rule to determine what points
are in the path.

`context`{.variable}`.`[`isPointInPath`](#dom-context-2d-ispointinpath){#dom-context-2d-ispointinpath-dev}`(``x`{.variable}`, ``y`{.variable}` [, ``fillRule`{.variable}` ])`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/isPointInPath](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/isPointInPath "The CanvasRenderingContext2D.isPointInPath() method of the Canvas 2D API reports whether or not the specified point is contained in the current path.")

Support in all current engines.

::: support
[Firefox2+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

`context`{.variable}`.`[`isPointInPath`](#dom-context-2d-ispointinpath){#drawing-paths-to-the-canvas:dom-context-2d-ispointinpath}`(``path`{.variable}`, ``x`{.variable}`, ``y`{.variable}` [, ``fillRule`{.variable}` ])`

Returns true if the given point is in the [current default
path](#current-default-path){#drawing-paths-to-the-canvas:current-default-path-7}
or the given path, using the given fill rule to determine what points
are in the path.

`context`{.variable}`.`[`isPointInStroke`](#dom-context-2d-ispointinstroke){#dom-context-2d-ispointinstroke-dev}`(``x`{.variable}`, ``y`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/isPointInStroke](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/isPointInStroke "The CanvasRenderingContext2D.isPointInStroke() method of the Canvas 2D API reports whether or not the specified point is inside the area contained by the stroking of a path.")

Support in all current engines.

::: support
[Firefox19+]{.firefox .yes}[Safari7+]{.safari .yes}[Chrome26+]{.chrome
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

`context`{.variable}`.`[`isPointInStroke`](#dom-context-2d-ispointinstroke){#drawing-paths-to-the-canvas:dom-context-2d-ispointinstroke}`(``path`{.variable}`, ``x`{.variable}`, ``y`{.variable}`)`

Returns true if the given point would be in the region covered by the
stroke of the [current default
path](#current-default-path){#drawing-paths-to-the-canvas:current-default-path-8}
or the given path, given the current stroke style.

The [`beginPath()`]{#dom-context-2d-beginpath .dfn
dfn-for="CanvasDrawPath" dfn-type="method"} method steps are to empty
the list of subpaths in
[this](https://webidl.spec.whatwg.org/#this){#drawing-paths-to-the-canvas:this
x-internal="this"}\'s [current default
path](#current-default-path){#drawing-paths-to-the-canvas:current-default-path-9}
so that it once again has zero subpaths.

Where the following method definitions use the term [intended
path]{#intended-path .dfn} for a
[`Path2D`{#drawing-paths-to-the-canvas:path2d}](#path2d)-or-null
`path`{.variable}, it means `path`{.variable} itself if it is a
[`Path2D`{#drawing-paths-to-the-canvas:path2d-2}](#path2d) object, or
the [current default
path](#current-default-path){#drawing-paths-to-the-canvas:current-default-path-10}
otherwise.

When the [intended
path](#intended-path){#drawing-paths-to-the-canvas:intended-path} is a
[`Path2D`{#drawing-paths-to-the-canvas:path2d-3}](#path2d) object, the
coordinates and lines of its subpaths must be transformed according to
the [current transformation
matrix](#transformations){#drawing-paths-to-the-canvas:transformations}
on the object implementing the
[`CanvasTransform`{#drawing-paths-to-the-canvas:canvastransform}](#canvastransform)
interface when used by these methods (without affecting the
[`Path2D`{#drawing-paths-to-the-canvas:path2d-4}](#path2d) object
itself). When the intended path is the [current default
path](#current-default-path){#drawing-paths-to-the-canvas:current-default-path-11},
it is not affected by the transform. (This is because transformations
already affect the [current default
path](#current-default-path){#drawing-paths-to-the-canvas:current-default-path-12}
when it is constructed, so applying it when it is painted as well would
result in a double transformation.)

The [`fill(``fillRule`{.variable}`)`]{#dom-context-2d-fill .dfn
dfn-for="CanvasDrawPath" dfn-type="method"} method steps are to run the
[fill steps](#fill-steps){#drawing-paths-to-the-canvas:fill-steps} given
[this](https://webidl.spec.whatwg.org/#this){#drawing-paths-to-the-canvas:this-2
x-internal="this"}, null, and `fillRule`{.variable}.

The
[`fill(``path`{.variable}`, ``fillRule`{.variable}`)`]{#dom-context-2d-fill-path
.dfn dfn-for="CanvasDrawPath" dfn-type="method"} method steps are to run
the [fill steps](#fill-steps){#drawing-paths-to-the-canvas:fill-steps-2}
given
[this](https://webidl.spec.whatwg.org/#this){#drawing-paths-to-the-canvas:this-3
x-internal="this"}, `path`{.variable}, and `fillRule`{.variable}.

The [fill steps]{#fill-steps .dfn}, given a
[`CanvasDrawPath`{#drawing-paths-to-the-canvas:canvasdrawpath-2}](#canvasdrawpath)
`context`{.variable}, a
[`Path2D`{#drawing-paths-to-the-canvas:path2d-5}](#path2d)-or-null
`path`{.variable}, and a [fill
rule](#fill-rule){#drawing-paths-to-the-canvas:fill-rule}
`fillRule`{.variable}, are to fill all the subpaths of the [intended
path](#intended-path){#drawing-paths-to-the-canvas:intended-path-2} for
`path`{.variable}, using `context`{.variable}\'s [fill
style](#concept-canvasfillstrokestyles-fill-style){#drawing-paths-to-the-canvas:concept-canvasfillstrokestyles-fill-style},
and using the [fill
rule](#fill-rule){#drawing-paths-to-the-canvas:fill-rule-2} indicated by
`fillRule`{.variable}. Open subpaths must be implicitly closed when
being filled (without affecting the actual subpaths).

The [`stroke()`]{#dom-context-2d-stroke .dfn dfn-for="CanvasDrawPath"
dfn-type="method"} method steps are to run the [stroke
steps](#stroke-steps){#drawing-paths-to-the-canvas:stroke-steps} given
[this](https://webidl.spec.whatwg.org/#this){#drawing-paths-to-the-canvas:this-4
x-internal="this"} and null.

The [`stroke(``path`{.variable}`)`]{#dom-context-2d-stroke-path .dfn
dfn-for="CanvasDrawPath" dfn-type="method"} method steps are to run the
[stroke
steps](#stroke-steps){#drawing-paths-to-the-canvas:stroke-steps-2} given
[this](https://webidl.spec.whatwg.org/#this){#drawing-paths-to-the-canvas:this-5
x-internal="this"} and `path`{.variable}.

The [stroke steps]{#stroke-steps .dfn}, given a
[`CanvasDrawPath`{#drawing-paths-to-the-canvas:canvasdrawpath-3}](#canvasdrawpath)
`context`{.variable} and a
[`Path2D`{#drawing-paths-to-the-canvas:path2d-6}](#path2d)-or-null
`path`{.variable}, are to
[trace](#trace-a-path){#drawing-paths-to-the-canvas:trace-a-path} the
[intended
path](#intended-path){#drawing-paths-to-the-canvas:intended-path-3} for
`path`{.variable}, using `context`{.variable}\'s line styles as set by
its
[`CanvasPathDrawingStyles`{#drawing-paths-to-the-canvas:canvaspathdrawingstyles}](#canvaspathdrawingstyles)
mixin, and then fill the resulting path using `context`{.variable}\'s
[stroke
style](#concept-canvasfillstrokestyles-stroke-style){#drawing-paths-to-the-canvas:concept-canvasfillstrokestyles-stroke-style},
using the [nonzero winding
rule](#dom-context-2d-fillrule-nonzero){#drawing-paths-to-the-canvas:dom-context-2d-fillrule-nonzero}.

As a result of how the algorithm to [trace a
path](#trace-a-path){#drawing-paths-to-the-canvas:trace-a-path-2} is
defined, overlapping parts of the paths in one stroke operation are
treated as if their union was what was painted.

The stroke *style* is affected by the transformation during painting,
even if the [current default
path](#current-default-path){#drawing-paths-to-the-canvas:current-default-path-13}
is used.

Paths, when filled or stroked, must be painted without affecting the
[current default
path](#current-default-path){#drawing-paths-to-the-canvas:current-default-path-14}
or any [`Path2D`{#drawing-paths-to-the-canvas:path2d-7}](#path2d)
objects, and must be subject to [shadow
effects](#shadows){#drawing-paths-to-the-canvas:shadows}, [global
alpha](#concept-canvas-global-alpha){#drawing-paths-to-the-canvas:concept-canvas-global-alpha},
the [clipping
region](#clipping-region){#drawing-paths-to-the-canvas:clipping-region},
and the [current compositing and blending
operator](#current-compositing-and-blending-operator){#drawing-paths-to-the-canvas:current-compositing-and-blending-operator}.
(The effect of transformations is described above and varies based on
which path is being used.)

------------------------------------------------------------------------

The [`clip(``fillRule`{.variable}`)`]{#dom-context-2d-clip .dfn
dfn-for="CanvasDrawPath" dfn-type="method"} method steps are to run the
[clip steps](#clip-steps){#drawing-paths-to-the-canvas:clip-steps} given
[this](https://webidl.spec.whatwg.org/#this){#drawing-paths-to-the-canvas:this-6
x-internal="this"}, null, and `fillRule`{.variable}.

The
[`clip(``path`{.variable}`, ``fillRule`{.variable}`)`]{#dom-context-2d-clip-path
.dfn dfn-for="CanvasDrawPath" dfn-type="method"} method steps are to run
the [clip steps](#clip-steps){#drawing-paths-to-the-canvas:clip-steps-2}
given
[this](https://webidl.spec.whatwg.org/#this){#drawing-paths-to-the-canvas:this-7
x-internal="this"}, `path`{.variable}, and `fillRule`{.variable}.

The [clip steps]{#clip-steps .dfn}, given a
[`CanvasDrawPath`{#drawing-paths-to-the-canvas:canvasdrawpath-4}](#canvasdrawpath)
`context`{.variable}, a
[`Path2D`{#drawing-paths-to-the-canvas:path2d-8}](#path2d)-or-null
`path`{.variable}, and a [fill
rule](#fill-rule){#drawing-paths-to-the-canvas:fill-rule-3}
`fillRule`{.variable}, are to create a new [clipping
region]{#clipping-region .dfn} by calculating the intersection of
`context`{.variable}\'s current clipping region and the area described
by the [intended
path](#intended-path){#drawing-paths-to-the-canvas:intended-path-4} for
`path`{.variable}, using the [fill
rule](#fill-rule){#drawing-paths-to-the-canvas:fill-rule-4} indicated by
`fillRule`{.variable}. Open subpaths must be implicitly closed when
computing the clipping region, without affecting the actual subpaths.
The new clipping region replaces the current clipping region.

When the context is initialized, its current clipping region must be set
to the largest infinite surface (i.e. by default, no clipping occurs).

------------------------------------------------------------------------

The
[`isPointInPath(``x`{.variable}`, ``y`{.variable}`, ``fillRule`{.variable}`)`]{#dom-context-2d-ispointinpath
.dfn dfn-for="CanvasDrawPath" dfn-type="method"} method steps are to
return the result of the [is point in path
steps](#is-point-in-path-steps){#drawing-paths-to-the-canvas:is-point-in-path-steps}
given
[this](https://webidl.spec.whatwg.org/#this){#drawing-paths-to-the-canvas:this-8
x-internal="this"}, null, `x`{.variable}, `y`{.variable}, and
`fillRule`{.variable}.

The
[`isPointInPath(``path`{.variable}`, ``x`{.variable}`, ``y`{.variable}`, ``fillRule`{.variable}`)`]{#dom-context-2d-ispointinpath-path
.dfn dfn-for="CanvasDrawPath" dfn-type="method"} method steps are to
return the result of the [is point in path
steps](#is-point-in-path-steps){#drawing-paths-to-the-canvas:is-point-in-path-steps-2}
given
[this](https://webidl.spec.whatwg.org/#this){#drawing-paths-to-the-canvas:this-9
x-internal="this"}, `path`{.variable}, `x`{.variable}, `y`{.variable},
and `fillRule`{.variable}.

The [is point in path steps]{#is-point-in-path-steps .dfn}, given a
[`CanvasDrawPath`{#drawing-paths-to-the-canvas:canvasdrawpath-5}](#canvasdrawpath)
`context`{.variable}, a
[`Path2D`{#drawing-paths-to-the-canvas:path2d-9}](#path2d)-or-null
`path`{.variable}, two numbers `x`{.variable} and `y`{.variable}, and a
[fill rule](#fill-rule){#drawing-paths-to-the-canvas:fill-rule-5}
`fillRule`{.variable}, are:

1.  If `x`{.variable} or `y`{.variable} are infinite or NaN, then return
    false.

2.  If the point given by the `x`{.variable} and `y`{.variable}
    coordinates, when treated as coordinates in the canvas coordinate
    space unaffected by the current transformation, is inside the
    [intended
    path](#intended-path){#drawing-paths-to-the-canvas:intended-path-5}
    for `path`{.variable} as determined by the [fill
    rule](#fill-rule){#drawing-paths-to-the-canvas:fill-rule-6}
    indicated by `fillRule`{.variable}, then return true. Open subpaths
    must be implicitly closed when computing the area inside the path,
    without affecting the actual subpaths. Points on the path itself
    must be considered to be inside the path.

3.  Return false.

------------------------------------------------------------------------

The
[`isPointInStroke(``x`{.variable}`, ``y`{.variable}`)`]{#dom-context-2d-ispointinstroke
.dfn dfn-for="CanvasDrawPath" dfn-type="method"} method steps are to
return the result of the [is point in stroke
steps](#is-point-in-stroke-steps){#drawing-paths-to-the-canvas:is-point-in-stroke-steps}
given
[this](https://webidl.spec.whatwg.org/#this){#drawing-paths-to-the-canvas:this-10
x-internal="this"}, null, `x`{.variable}, and `y`{.variable}.

The
[`isPointInStroke(``path`{.variable}`, ``x`{.variable}`, ``y`{.variable}`)`]{#dom-context-2d-ispointinstroke-path
.dfn dfn-for="CanvasDrawPath" dfn-type="method"} method steps are to
return the result of the [is point in stroke
steps](#is-point-in-stroke-steps){#drawing-paths-to-the-canvas:is-point-in-stroke-steps-2}
given
[this](https://webidl.spec.whatwg.org/#this){#drawing-paths-to-the-canvas:this-11
x-internal="this"}, `path`{.variable}, `x`{.variable}, and
`y`{.variable}.

The [is point in stroke steps]{#is-point-in-stroke-steps .dfn}, given a
[`CanvasDrawPath`{#drawing-paths-to-the-canvas:canvasdrawpath-6}](#canvasdrawpath)
`context`{.variable}, a
[`Path2D`{#drawing-paths-to-the-canvas:path2d-10}](#path2d)-or-null
`path`{.variable}, and two numbers `x`{.variable} and `y`{.variable},
are:

1.  If `x`{.variable} or `y`{.variable} are infinite or NaN, then return
    false.

2.  If the point given by the `x`{.variable} and `y`{.variable}
    coordinates, when treated as coordinates in the canvas coordinate
    space unaffected by the current transformation, is inside the path
    that results from
    [tracing](#trace-a-path){#drawing-paths-to-the-canvas:trace-a-path-3}
    the [intended
    path](#intended-path){#drawing-paths-to-the-canvas:intended-path-6}
    for `path`{.variable}, using the [nonzero winding
    rule](#dom-context-2d-fillrule-nonzero){#drawing-paths-to-the-canvas:dom-context-2d-fillrule-nonzero-2},
    and using `context`{.variable}\'s line styles as set by its
    [`CanvasPathDrawingStyles`{#drawing-paths-to-the-canvas:canvaspathdrawingstyles-2}](#canvaspathdrawingstyles)
    mixin, then return true. Points on the resulting path must be
    considered to be inside the path.

3.  Return false.

------------------------------------------------------------------------

::: {#drawCustomFocusRingExample .example}
This
[`canvas`{#drawing-paths-to-the-canvas:the-canvas-element}](#the-canvas-element)
element has a couple of checkboxes. The path-related commands are
highlighted:

``` html
<canvas height=400 width=750>
 <label><input type=checkbox id=showA> Show As</label>
 <label><input type=checkbox id=showB> Show Bs</label>
 <!-- ... -->
</canvas>
<script>
 function drawCheckbox(context, element, x, y, paint) {
   context.save();
   context.font = '10px sans-serif';
   context.textAlign = 'left';
   context.textBaseline = 'middle';
   var metrics = context.measureText(element.labels[0].textContent);
   if (paint) {
     context.beginPath();
     context.strokeStyle = 'black';
     context.rect(x-5, y-5, 10, 10);
     context.stroke();
     if (element.checked) {
       context.fillStyle = 'black';
       context.fill();
     }
     context.fillText(element.labels[0].textContent, x+5, y);
   }
   context.beginPath();
   context.rect(x-7, y-7, 12 + metrics.width+2, 14);

   context.drawFocusIfNeeded(element);
   context.restore();
 }
 function drawBase() { /* ... */ }
 function drawAs() { /* ... */ }
 function drawBs() { /* ... */ }
 function redraw() {
   var canvas = document.getElementsByTagName('canvas')[0];
   var context = canvas.getContext('2d');
   context.clearRect(0, 0, canvas.width, canvas.height);
   drawCheckbox(context, document.getElementById('showA'), 20, 40, true);
   drawCheckbox(context, document.getElementById('showB'), 20, 60, true);
   drawBase();
   if (document.getElementById('showA').checked)
     drawAs();
   if (document.getElementById('showB').checked)
     drawBs();
 }
 function processClick(event) {
   var canvas = document.getElementsByTagName('canvas')[0];
   var context = canvas.getContext('2d');
   var x = event.clientX;
   var y = event.clientY;
   var node = event.target;
   while (node) {
     x -= node.offsetLeft - node.scrollLeft;
     y -= node.offsetTop - node.scrollTop;
     node = node.offsetParent;
   }
   drawCheckbox(context, document.getElementById('showA'), 20, 40, false);
   if (context.isPointInPath(x, y))
     document.getElementById('showA').checked = !(document.getElementById('showA').checked);
   drawCheckbox(context, document.getElementById('showB'), 20, 60, false);
   if (context.isPointInPath(x, y))
     document.getElementById('showB').checked = !(document.getElementById('showB').checked);
   redraw();
 }
 document.getElementsByTagName('canvas')[0].addEventListener('focus', redraw, true);
 document.getElementsByTagName('canvas')[0].addEventListener('blur', redraw, true);
 document.getElementsByTagName('canvas')[0].addEventListener('change', redraw, true);
 document.getElementsByTagName('canvas')[0].addEventListener('click', processClick, false);
 redraw();
</script>
```
:::

###### [4.12.5.1.14]{.secno} Drawing focus rings[](#drawing-focus-rings-and-scrolling-paths-into-view){.self-link} {#drawing-focus-rings-and-scrolling-paths-into-view}

`context`{.variable}`.`[`drawFocusIfNeeded`](#dom-context-2d-drawfocusifneeded){#dom-context-2d-drawfocusifneeded-dev}`(``element`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/drawFocusIfNeeded](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/drawFocusIfNeeded "The CanvasRenderingContext2D.drawFocusIfNeeded() method of the Canvas 2D API draws a focus ring around the current or given path, if the specified element is focused.")

Support in all current engines.

::: support
[Firefox32+]{.firefox .yes}[Safari8+]{.safari .yes}[Chrome37+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)14+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

If `element`{.variable} is
[focused](interaction.html#focused){#drawing-focus-rings-and-scrolling-paths-into-view:focused},
draws a focus ring around the [current default
path](#current-default-path){#drawing-focus-rings-and-scrolling-paths-into-view:current-default-path},
following the platform conventions for focus rings.

`context`{.variable}`.`[`drawFocusIfNeeded`](#dom-context-2d-drawfocusifneeded-path-element){#dom-context-2d-drawfocusifneeded-path-element-dev}`(``path`{.variable}`, ``element`{.variable}`)`

If `element`{.variable} is
[focused](interaction.html#focused){#drawing-focus-rings-and-scrolling-paths-into-view:focused-2},
draws a focus ring around `path`{.variable}, following the platform
conventions for focus rings.

Objects that implement the
[`CanvasUserInterface`{#drawing-focus-rings-and-scrolling-paths-into-view:canvasuserinterface}](#canvasuserinterface)
interface provide the following methods to draw focus rings.

The
[`drawFocusIfNeeded(``element`{.variable}`)`]{#dom-context-2d-drawfocusifneeded
.dfn dfn-for="CanvasUserInterface" dfn-type="method"} method steps are
to [draw focus if
needed](#draw-focus-if-needed){#drawing-focus-rings-and-scrolling-paths-into-view:draw-focus-if-needed}
given
[this](https://webidl.spec.whatwg.org/#this){#drawing-focus-rings-and-scrolling-paths-into-view:this
x-internal="this"}, `element`{.variable}, and
[this](https://webidl.spec.whatwg.org/#this){#drawing-focus-rings-and-scrolling-paths-into-view:this-2
x-internal="this"}\'s [current default
path](#current-default-path){#drawing-focus-rings-and-scrolling-paths-into-view:current-default-path-2}.

The
[`drawFocusIfNeeded(``path`{.variable}`, ``element`{.variable}`)`]{#dom-context-2d-drawfocusifneeded-path-element
.dfn dfn-for="CanvasUserInterface" dfn-type="method"} method steps are
to [draw focus if
needed](#draw-focus-if-needed){#drawing-focus-rings-and-scrolling-paths-into-view:draw-focus-if-needed-2}
given
[this](https://webidl.spec.whatwg.org/#this){#drawing-focus-rings-and-scrolling-paths-into-view:this-3
x-internal="this"}, `element`{.variable}, and `path`{.variable}.

To [draw focus if needed]{#draw-focus-if-needed .dfn}, given an object
implementing
[`CanvasUserInterface`{#drawing-focus-rings-and-scrolling-paths-into-view:canvasuserinterface-2}](#canvasuserinterface)
`context`{.variable}, an element `element`{.variable}, and a
[path](#concept-path){#drawing-focus-rings-and-scrolling-paths-into-view:concept-path}
`path`{.variable}:

1.  If `element`{.variable} is not
    [focused](interaction.html#focused){#drawing-focus-rings-and-scrolling-paths-into-view:focused-3}
    or is not a descendant of `context`{.variable}\'s
    [`canvas`{#drawing-focus-rings-and-scrolling-paths-into-view:the-canvas-element}](#the-canvas-element)
    element, then return.

2.  Draw a focus ring of the appropriate style along `path`{.variable},
    following platform conventions.

    Some platforms only draw focus rings around elements that have been
    focused from the keyboard, and not those focused from the mouse.
    Other platforms simply don\'t draw focus rings around some elements
    at all unless relevant accessibility features are enabled. This API
    is intended to follow these conventions. User agents that implement
    distinctions based on the manner in which the element was focused
    are encouraged to classify focus driven by the
    [`focus()`{#drawing-focus-rings-and-scrolling-paths-into-view:dom-focus}](interaction.html#dom-focus)
    method based on the kind of user interaction event from which the
    call was triggered (if any).

    The focus ring should not be subject to the [shadow
    effects](#shadows){#drawing-focus-rings-and-scrolling-paths-into-view:shadows},
    the [global
    alpha](#concept-canvas-global-alpha){#drawing-focus-rings-and-scrolling-paths-into-view:concept-canvas-global-alpha},
    the [current compositing and blending
    operator](#current-compositing-and-blending-operator){#drawing-focus-rings-and-scrolling-paths-into-view:current-compositing-and-blending-operator},
    the [fill
    style](#concept-canvasfillstrokestyles-fill-style){#drawing-focus-rings-and-scrolling-paths-into-view:concept-canvasfillstrokestyles-fill-style},
    the [stroke
    style](#concept-canvasfillstrokestyles-stroke-style){#drawing-focus-rings-and-scrolling-paths-into-view:concept-canvasfillstrokestyles-stroke-style},
    or any of the members in the
    [`CanvasPathDrawingStyles`{#drawing-focus-rings-and-scrolling-paths-into-view:canvaspathdrawingstyles}](#canvaspathdrawingstyles),
    [`CanvasTextDrawingStyles`{#drawing-focus-rings-and-scrolling-paths-into-view:canvastextdrawingstyles}](#canvastextdrawingstyles)
    interfaces, but *should* be subject to the [clipping
    region](#clipping-region){#drawing-focus-rings-and-scrolling-paths-into-view:clipping-region}.
    (The effect of transformations is described above and varies based
    on which path is being used.)

3.  [Inform the user](#inform) that the focus is at the location given
    by the intended path. User agents may wait until the next time the
    [event
    loop](webappapis.html#event-loop){#drawing-focus-rings-and-scrolling-paths-into-view:event-loop}
    reaches its [update the
    rendering](webappapis.html#update-the-rendering){#drawing-focus-rings-and-scrolling-paths-into-view:update-the-rendering}
    step to optionally inform the user.

User agents should not implicitly close open subpaths in the intended
path when drawing the focus ring.

This might be a moot point, however. For example, if the focus ring is
drawn as an axis-aligned bounding rectangle around the points in the
intended path, then whether the subpaths are closed or not has no
effect. This specification intentionally does not specify precisely how
focus rings are to be drawn: user agents are expected to honor their
platform\'s native conventions.

\"Inform the user\", as used in this section, does not imply any
persistent state change. It could mean, for instance, calling a system
accessibility API to notify assistive technologies such as magnification
tools so that the user\'s magnifier moves to the given area of the
canvas. However, it does not associate the path with the element, or
provide a region for tactile feedback, etc.

###### [4.12.5.1.15]{.secno} Drawing images[](#drawing-images){.self-link}

Objects that implement the
[`CanvasDrawImage`{#drawing-images:canvasdrawimage}](#canvasdrawimage)
interface have the [`drawImage()`]{#dom-context-2d-drawimage .dfn
dfn-for="CanvasDrawImage" dfn-type="method"} method to draw images.

This method can be invoked with three different sets of arguments:

- `drawImage(``image`{.variable}`, ``dx`{.variable}`, ``dy`{.variable}`)`
- `drawImage(``image`{.variable}`, ``dx`{.variable}`, ``dy`{.variable}`, ``dw`{.variable}`, ``dh`{.variable}`)`
- `drawImage(``image`{.variable}`, ``sx`{.variable}`, ``sy`{.variable}`, ``sw`{.variable}`, ``sh`{.variable}`, ``dx`{.variable}`, ``dy`{.variable}`, ``dw`{.variable}`, ``dh`{.variable}`)`

`context`{.variable}`.`[`drawImage`](#dom-context-2d-drawimage){#dom-context-2d-drawimage-dev}`(``image`{.variable}`, ``dx`{.variable}`, ``dy`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/drawImage](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/drawImage "The CanvasRenderingContext2D.drawImage() method of the Canvas 2D API provides different ways to draw an image onto the canvas.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari2+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera9+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

`context`{.variable}`.`[`drawImage`](#dom-context-2d-drawimage){#drawing-images:dom-context-2d-drawimage}`(``image`{.variable}`, ``dx`{.variable}`, ``dy`{.variable}`, ``dw`{.variable}`, ``dh`{.variable}`)`

`context`{.variable}`.`[`drawImage`](#dom-context-2d-drawimage){#drawing-images:dom-context-2d-drawimage-2}`(``image`{.variable}`, ``sx`{.variable}`, ``sy`{.variable}`, ``sw`{.variable}`, ``sh`{.variable}`, ``dx`{.variable}`, ``dy`{.variable}`, ``dw`{.variable}`, ``dh`{.variable}`)`

Draws the given image onto the canvas. The arguments are interpreted as
follows:

![The sx and sy parameters give the x and y coordinates of the source
rectangle; the sw and sh arguments give the width and height of the
source rectangle; the dx and dy give the x and y coordinates of the
destination rectangle; and the dw and dh arguments give the width and
height of the destination rectangle.](/images/drawImage.png){width="356"
height="356"}

If the image isn\'t yet fully decoded, then nothing is drawn. If the
image is a canvas with no data, throws an
[\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#drawing-images:invalidstateerror
x-internal="invalidstateerror"}
[`DOMException`{#drawing-images:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

When the
[`drawImage()`{#drawing-images:dom-context-2d-drawimage-3}](#dom-context-2d-drawimage)
method is invoked, the user agent must run these steps:

1.  If any of the arguments are infinite or NaN, then return.

2.  Let `usability`{.variable} be the result of [checking the usability
    of
    `image`{.variable}](#check-the-usability-of-the-image-argument){#drawing-images:check-the-usability-of-the-image-argument}.

3.  If `usability`{.variable} is *bad*, then return (without drawing
    anything).

4.  Establish the source and destination rectangles as follows:

    If not specified, the `dw`{.variable} and `dh`{.variable} arguments
    must default to the values of `sw`{.variable} and `sh`{.variable},
    interpreted such that one [CSS
    pixel](https://drafts.csswg.org/css-values/#px){#drawing-images:'px'
    x-internal="'px'"} in the image is treated as one unit in the
    [output bitmap](#output-bitmap){#drawing-images:output-bitmap}\'s
    coordinate space. If the `sx`{.variable}, `sy`{.variable},
    `sw`{.variable}, and `sh`{.variable} arguments are omitted, then
    they must default to 0, 0, the image\'s [natural
    width](https://drafts.csswg.org/css-images/#natural-width){#drawing-images:natural-width
    x-internal="natural-width"} in image pixels, and the image\'s
    [natural
    height](https://drafts.csswg.org/css-images/#natural-height){#drawing-images:natural-height
    x-internal="natural-height"} in image pixels, respectively. If the
    image has no [natural
    dimensions](https://drafts.csswg.org/css-images/#natural-dimensions){#drawing-images:natural-dimensions
    x-internal="natural-dimensions"}, then the *concrete object size*
    must be used instead, as determined using the CSS \"[Concrete Object
    Size
    Resolution](https://drafts.csswg.org/css-images/#default-sizing)\"
    algorithm, with the *specified size* having neither a definite width
    nor height, nor any additional constraints, the object\'s natural
    properties being those of the `image`{.variable} argument, and the
    [default object
    size](https://drafts.csswg.org/css-images/#default-object-size){#drawing-images:default-object-size
    x-internal="default-object-size"} being the size of the [output
    bitmap](#output-bitmap){#drawing-images:output-bitmap-2}.
    [\[CSSIMAGES\]](references.html#refsCSSIMAGES)

    The source rectangle is the rectangle whose corners are the four
    points (`sx`{.variable}, `sy`{.variable}),
    (`sx`{.variable}+`sw`{.variable}, `sy`{.variable}),
    (`sx`{.variable}+`sw`{.variable}, `sy`{.variable}+`sh`{.variable}),
    (`sx`{.variable}, `sy`{.variable}+`sh`{.variable}).

    The destination rectangle is the rectangle whose corners are the
    four points (`dx`{.variable}, `dy`{.variable}),
    (`dx`{.variable}+`dw`{.variable}, `dy`{.variable}),
    (`dx`{.variable}+`dw`{.variable}, `dy`{.variable}+`dh`{.variable}),
    (`dx`{.variable}, `dy`{.variable}+`dh`{.variable}).

    When the source rectangle is outside the source image, the source
    rectangle must be clipped to the source image and the destination
    rectangle must be clipped in the same proportion.

    When the destination rectangle is outside the destination image (the
    [output bitmap](#output-bitmap){#drawing-images:output-bitmap-3}),
    the pixels that land outside the [output
    bitmap](#output-bitmap){#drawing-images:output-bitmap-4} are
    discarded, as if the destination was an infinite canvas whose
    rendering was clipped to the dimensions of the [output
    bitmap](#output-bitmap){#drawing-images:output-bitmap-5}.

5.  If one of the `sw`{.variable} or `sh`{.variable} arguments is zero,
    then return. Nothing is painted.

6.  Paint the region of the `image`{.variable} argument specified by the
    source rectangle on the region of the rendering context\'s [output
    bitmap](#output-bitmap){#drawing-images:output-bitmap-6} specified
    by the destination rectangle, after applying the [current
    transformation
    matrix](#transformations){#drawing-images:transformations} to the
    destination rectangle.

    The image data must be processed in the original direction, even if
    the dimensions given are negative.

    When scaling up, if the
    [`imageSmoothingEnabled`{#drawing-images:dom-context-2d-imagesmoothingenabled}](#dom-context-2d-imagesmoothingenabled)
    attribute is set to true, the user agent should attempt to apply a
    smoothing algorithm to the image data when it is scaled. User agents
    which support multiple filtering algorithms may use the value of the
    [`imageSmoothingQuality`{#drawing-images:dom-context-2d-imagesmoothingquality}](#dom-context-2d-imagesmoothingquality)
    attribute to guide the choice of filtering algorithm when the
    [`imageSmoothingEnabled`{#drawing-images:dom-context-2d-imagesmoothingenabled-2}](#dom-context-2d-imagesmoothingenabled)
    attribute is set to true. Otherwise, the image must be rendered
    using nearest-neighbor interpolation.

    This specification does not define the precise algorithm to use when
    scaling an image down, or when scaling an image up when the
    [`imageSmoothingEnabled`{#drawing-images:dom-context-2d-imagesmoothingenabled-3}](#dom-context-2d-imagesmoothingenabled)
    attribute is set to true.

    When a
    [`canvas`{#drawing-images:the-canvas-element}](#the-canvas-element)
    element is drawn onto itself, the [drawing
    model](#drawing-model){#drawing-images:drawing-model} requires the
    source to be copied before the image is drawn, so it is possible to
    copy parts of a
    [`canvas`{#drawing-images:the-canvas-element-2}](#the-canvas-element)
    element onto overlapping parts of itself.

    If the original image data is a bitmap image, then the value painted
    at a point in the destination rectangle is computed by filtering the
    original image data. The user agent may use any filtering algorithm
    (for example bilinear interpolation or nearest-neighbor). When the
    filtering algorithm requires a pixel value from outside the original
    image data, it must instead use the value from the nearest edge
    pixel. (That is, the filter uses \'clamp-to-edge\' behavior.) When
    the filtering algorithm requires a pixel value from outside the
    source rectangle but inside the original image data, then the value
    from the original image data must be used.

    Thus, scaling an image in parts or in whole will have the same
    effect. This does mean that when sprites coming from a single sprite
    sheet are to be scaled, adjacent images in the sprite sheet can
    interfere. This can be avoided by ensuring each sprite in the sheet
    is surrounded by a border of [transparent
    black](https://drafts.csswg.org/css-color/#transparent-black){#drawing-images:transparent-black
    x-internal="transparent-black"}, or by copying sprites to be scaled
    into temporary
    [`canvas`{#drawing-images:the-canvas-element-3}](#the-canvas-element)
    elements and drawing the scaled sprites from there.

    Images are painted without affecting the current path, and are
    subject to [shadow effects](#shadows){#drawing-images:shadows},
    [global
    alpha](#concept-canvas-global-alpha){#drawing-images:concept-canvas-global-alpha},
    the [clipping
    region](#clipping-region){#drawing-images:clipping-region}, and the
    [current compositing and blending
    operator](#current-compositing-and-blending-operator){#drawing-images:current-compositing-and-blending-operator}.

7.  If `image`{.variable} [is not
    origin-clean](#the-image-argument-is-not-origin-clean){#drawing-images:the-image-argument-is-not-origin-clean},
    then set the
    [`CanvasRenderingContext2D`{#drawing-images:canvasrenderingcontext2d}](#canvasrenderingcontext2d)\'s
    [origin-clean](#concept-canvas-origin-clean){#drawing-images:concept-canvas-origin-clean}
    flag to false.

###### [4.12.5.1.16]{.secno} [Pixel manipulation]{.dfn}[](#pixel-manipulation){.self-link}

`imageData`{.variable}` = ``context`{.variable}`.`[`createImageData`](#dom-context-2d-createimagedata-imagedata){#dom-context-2d-createimagedata-imagedata-dev}`(``imageData`{.variable}`)`

Returns an
[`ImageData`{#pixel-manipulation:imagedata}](imagebitmap-and-animations.html#imagedata)
object with the same dimensions and color space as the argument. All the
pixels in the returned object are [transparent
black](https://drafts.csswg.org/css-color/#transparent-black){#pixel-manipulation:transparent-black
x-internal="transparent-black"}.

`imageData`{.variable}` = ``context`{.variable}`.`[`createImageData`](#dom-context-2d-createimagedata){#dom-context-2d-createimagedata-dev}`(``sw`{.variable}`, ``sh`{.variable}` [, ``settings`{.variable}`])`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/createImageData](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/createImageData "The CanvasRenderingContext2D.createImageData() method of the Canvas 2D API creates a new, blank ImageData object with the specified dimensions. All of the pixels in the new object are transparent black.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome2+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns an
[`ImageData`{#pixel-manipulation:imagedata-2}](imagebitmap-and-animations.html#imagedata)
object with the given dimensions. The color space of the returned object
is the [color
space](#concept-canvas-color-space){#pixel-manipulation:concept-canvas-color-space}
of `context`{.variable} unless overridden by `settings`{.variable}. All
the pixels in the returned object are [transparent
black](https://drafts.csswg.org/css-color/#transparent-black){#pixel-manipulation:transparent-black-2
x-internal="transparent-black"}.

Throws an
[\"`IndexSizeError`\"](https://webidl.spec.whatwg.org/#indexsizeerror){#pixel-manipulation:indexsizeerror
x-internal="indexsizeerror"}
[`DOMException`{#pixel-manipulation:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if either of the width or height arguments are zero.

`imageData`{.variable}` = ``context`{.variable}`.`[`getImageData`](#dom-context-2d-getimagedata){#dom-context-2d-getimagedata-dev}`(``sx`{.variable}`, ``sy`{.variable}`, ``sw`{.variable}`, ``sh`{.variable}` [, ``settings`{.variable}`])`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/getImageData](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/getImageData "The CanvasRenderingContext2D method getImageData() of the Canvas 2D API returns an ImageData object representing the underlying pixel data for a specified portion of the canvas.")

Support in all current engines.

::: support
[Firefox2+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome2+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera9.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

Returns an
[`ImageData`{#pixel-manipulation:imagedata-3}](imagebitmap-and-animations.html#imagedata)
object containing the image data for the given rectangle of the bitmap.
The color space of the returned object is the [color
space](#concept-canvas-color-space){#pixel-manipulation:concept-canvas-color-space-2}
of `context`{.variable} unless overridden by `settings`{.variable}.

Throws an
[\"`IndexSizeError`\"](https://webidl.spec.whatwg.org/#indexsizeerror){#pixel-manipulation:indexsizeerror-2
x-internal="indexsizeerror"}
[`DOMException`{#pixel-manipulation:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if the either of the width or height arguments are zero.

`context`{.variable}`.`[`putImageData`](#dom-context-2d-putimagedata){#dom-context-2d-putimagedata-dev}`(``imageData`{.variable}`, ``dx`{.variable}`, ``dy`{.variable}` [, ``dirtyX`{.variable}`, ``dirtyY`{.variable}`, ``dirtyWidth`{.variable}`, ``dirtyHeight`{.variable}` ])`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/putImageData](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/putImageData "The CanvasRenderingContext2D.putImageData() method of the Canvas 2D API paints data from the given ImageData object onto the canvas. If a dirty rectangle is provided, only the pixels from that rectangle are painted. This method is not affected by the canvas transformation matrix.")

Support in all current engines.

::: support
[Firefox2+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome2+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Paints the data from the given
[`ImageData`{#pixel-manipulation:imagedata-4}](imagebitmap-and-animations.html#imagedata)
object onto the bitmap. If a dirty rectangle is provided, only the
pixels from that rectangle are painted.

The
[`globalAlpha`{#pixel-manipulation:dom-context-2d-globalalpha}](#dom-context-2d-globalalpha)
and
[`globalCompositeOperation`{#pixel-manipulation:dom-context-2d-globalcompositeoperation}](#dom-context-2d-globalcompositeoperation)
properties, as well as the [shadow
attributes](#shadows){#pixel-manipulation:shadows}, are ignored for the
purposes of this method call; pixels in the canvas are replaced
wholesale, with no composition, alpha blending, no shadows, etc.

Throws an
[\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#pixel-manipulation:invalidstateerror
x-internal="invalidstateerror"}
[`DOMException`{#pixel-manipulation:domexception-3}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if the `imageData`{.variable} object\'s
[`data`{#pixel-manipulation:dom-imagedata-data}](imagebitmap-and-animations.html#dom-imagedata-data)
attribute value\'s \[\[ViewedArrayBuffer\]\] internal slot is detached.

Objects that implement the
[`CanvasImageData`{#pixel-manipulation:canvasimagedata}](#canvasimagedata)
interface provide the following methods for reading and writing pixel
data to the bitmap.

The
[`createImageData(``sw`{.variable}`, ``sh`{.variable}`, ``settings`{.variable}`)`]{#dom-context-2d-createimagedata
.dfn dfn-for="CanvasImageData" dfn-type="method"} method steps are:

1.  If one or both of `sw`{.variable} and `sh`{.variable} are zero, then
    throw an
    [\"`IndexSizeError`\"](https://webidl.spec.whatwg.org/#indexsizeerror){#pixel-manipulation:indexsizeerror-3
    x-internal="indexsizeerror"}
    [`DOMException`{#pixel-manipulation:domexception-4}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  Let `newImageData`{.variable} be a
    [new](https://webidl.spec.whatwg.org/#new){#pixel-manipulation:new
    x-internal="new"}
    [`ImageData`{#pixel-manipulation:imagedata-5}](imagebitmap-and-animations.html#imagedata)
    object.

3.  [Initialize](imagebitmap-and-animations.html#initialize-an-imagedata-object){#pixel-manipulation:initialize-an-imagedata-object}
    `newImageData`{.variable} given the absolute magnitude of
    `sw`{.variable}, the absolute magnitude of `sh`{.variable},
    `settings`{.variable}, and
    *[defaultColorSpace](imagebitmap-and-animations.html#initialize-imagedata-defaultcolorspace){#pixel-manipulation:initialize-imagedata-defaultcolorspace}*
    set to
    [this](https://webidl.spec.whatwg.org/#this){#pixel-manipulation:this
    x-internal="this"}\'s [color
    space](#concept-canvas-color-space){#pixel-manipulation:concept-canvas-color-space-3}.

4.  Initialize the image data of `newImageData`{.variable} to
    [transparent
    black](https://drafts.csswg.org/css-color/#transparent-black){#pixel-manipulation:transparent-black-3
    x-internal="transparent-black"}.

5.  Return `newImageData`{.variable}.

The
[`createImageData(``imageData`{.variable}`)`]{#dom-context-2d-createimagedata-imagedata
.dfn dfn-for="CanvasImageData" dfn-type="method"} method steps are:

1.  Let `newImageData`{.variable} be a
    [new](https://webidl.spec.whatwg.org/#new){#pixel-manipulation:new-2
    x-internal="new"}
    [`ImageData`{#pixel-manipulation:imagedata-6}](imagebitmap-and-animations.html#imagedata)
    object.

2.  Let `settings`{.variable} be the
    [`ImageDataSettings`{#pixel-manipulation:imagedatasettings}](imagebitmap-and-animations.html#imagedatasettings)
    object «\[
    \"[`colorSpace`{#pixel-manipulation:dom-imagedatasettings-colorspace}](imagebitmap-and-animations.html#dom-imagedatasettings-colorspace)\"
    →
    [this](https://webidl.spec.whatwg.org/#this){#pixel-manipulation:this-2
    x-internal="this"}\'s
    [colorSpace](imagebitmap-and-animations.html#dom-imagedata-colorspace){#pixel-manipulation:dom-imagedata-colorspace},
    \"[`pixelFormat`{#pixel-manipulation:dom-imagedatasettings-pixelformat}](imagebitmap-and-animations.html#dom-imagedatasettings-pixelformat)\"
    →
    [this](https://webidl.spec.whatwg.org/#this){#pixel-manipulation:this-3
    x-internal="this"}\'s
    [pixelFormat](imagebitmap-and-animations.html#dom-imagedata-pixelformat){#pixel-manipulation:dom-imagedata-pixelformat}
    \]».

3.  [Initialize](imagebitmap-and-animations.html#initialize-an-imagedata-object){#pixel-manipulation:initialize-an-imagedata-object-2}
    `newImageData`{.variable} given the value of
    `imageData`{.variable}\'s
    [`width`{#pixel-manipulation:dom-imagedata-width}](imagebitmap-and-animations.html#dom-imagedata-width)
    attribute, the value of `imageData`{.variable}\'s
    [`height`{#pixel-manipulation:dom-imagedata-height}](imagebitmap-and-animations.html#dom-imagedata-height)
    attribute, and `settings`{.variable}.

4.  Initialize the image data of `newImageData`{.variable} to
    [transparent
    black](https://drafts.csswg.org/css-color/#transparent-black){#pixel-manipulation:transparent-black-4
    x-internal="transparent-black"}.

5.  Return `newImageData`{.variable}.

The
[`getImageData(``sx`{.variable}`, ``sy`{.variable}`, ``sw`{.variable}`, ``sh`{.variable}`, ``settings`{.variable}`)`]{#dom-context-2d-getimagedata
.dfn dfn-for="CanvasImageData" dfn-type="method"} method steps are:

1.  If either the `sw`{.variable} or `sh`{.variable} arguments are zero,
    then throw an
    [\"`IndexSizeError`\"](https://webidl.spec.whatwg.org/#indexsizeerror){#pixel-manipulation:indexsizeerror-4
    x-internal="indexsizeerror"}
    [`DOMException`{#pixel-manipulation:domexception-5}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  If the
    [`CanvasRenderingContext2D`{#pixel-manipulation:canvasrenderingcontext2d}](#canvasrenderingcontext2d)\'s
    [origin-clean](#concept-canvas-origin-clean){#pixel-manipulation:concept-canvas-origin-clean}
    flag is set to false, then throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#pixel-manipulation:securityerror
    x-internal="securityerror"}
    [`DOMException`{#pixel-manipulation:domexception-6}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

3.  Let `imageData`{.variable} be a
    [new](https://webidl.spec.whatwg.org/#new){#pixel-manipulation:new-3
    x-internal="new"}
    [`ImageData`{#pixel-manipulation:imagedata-7}](imagebitmap-and-animations.html#imagedata)
    object.

4.  [Initialize](imagebitmap-and-animations.html#initialize-an-imagedata-object){#pixel-manipulation:initialize-an-imagedata-object-3}
    `imageData`{.variable} given `sw`{.variable}, `sh`{.variable},
    `settings`{.variable}, and
    *[defaultColorSpace](imagebitmap-and-animations.html#initialize-imagedata-defaultcolorspace){#pixel-manipulation:initialize-imagedata-defaultcolorspace-2}*
    set to
    [this](https://webidl.spec.whatwg.org/#this){#pixel-manipulation:this-4
    x-internal="this"}\'s [color
    space](#concept-canvas-color-space){#pixel-manipulation:concept-canvas-color-space-4}.

5.  Let the source rectangle be the rectangle whose corners are the four
    points (`sx`{.variable}, `sy`{.variable}),
    (`sx`{.variable}+`sw`{.variable}, `sy`{.variable}),
    (`sx`{.variable}+`sw`{.variable}, `sy`{.variable}+`sh`{.variable}),
    (`sx`{.variable}, `sy`{.variable}+`sh`{.variable}).

6.  Set the pixel values of `imageData`{.variable} to be the pixels of
    [this](https://webidl.spec.whatwg.org/#this){#pixel-manipulation:this-5
    x-internal="this"}\'s [output
    bitmap](#output-bitmap){#pixel-manipulation:output-bitmap} in the
    area specified by the source rectangle in the bitmap\'s coordinate
    space units, converted from
    [this](https://webidl.spec.whatwg.org/#this){#pixel-manipulation:this-6
    x-internal="this"}\'s [color
    space](#concept-canvas-color-space){#pixel-manipulation:concept-canvas-color-space-5}
    to `imageData`{.variable}\'s
    [`colorSpace`{#pixel-manipulation:dom-imagedata-colorspace-2}](imagebitmap-and-animations.html#dom-imagedata-colorspace)
    using
    [\'relative-colorimetric\'](https://drafts.csswg.org/css-color-5/#valdef-color-profile-rendering-intent-relative-colorimetric){#pixel-manipulation:'relative-colorimetric'
    x-internal="'relative-colorimetric'"} rendering intent.

7.  Set the pixels values of `imageData`{.variable} for areas of the
    source rectangle that are outside of the [output
    bitmap](#output-bitmap){#pixel-manipulation:output-bitmap-2} to
    [transparent
    black](https://drafts.csswg.org/css-color/#transparent-black){#pixel-manipulation:transparent-black-5
    x-internal="transparent-black"}.

8.  Return `imageData`{.variable}.

The
[`putImageData(``imageData`{.variable}`, ``dx`{.variable}`, ``dy`{.variable}`)`]{#dom-context-2d-putimagedata-short
.dfn dfn-for="CanvasImageData" dfn-type="method"} method steps are to
[put pixels from an `ImageData` onto a
bitmap](#dom-context2d-putimagedata-common){#pixel-manipulation:dom-context2d-putimagedata-common},
given `imageData`{.variable},
[this](https://webidl.spec.whatwg.org/#this){#pixel-manipulation:this-7
x-internal="this"}\'s [output
bitmap](#output-bitmap){#pixel-manipulation:output-bitmap-3},
`dx`{.variable}, `dy`{.variable}, 0, 0, `imageData`{.variable}\'s
[`width`{#pixel-manipulation:dom-imagedata-width-2}](imagebitmap-and-animations.html#dom-imagedata-width),
and `imageData`{.variable}\'s
[`height`{#pixel-manipulation:dom-imagedata-height-2}](imagebitmap-and-animations.html#dom-imagedata-height).

The
[`putImageData(``imageData`{.variable}`, ``dx`{.variable}`, ``dy`{.variable}`, ``dirtyX`{.variable}`, ``dirtyY`{.variable}`, ``dirtyWidth`{.variable}`, ``dirtyHeight`{.variable}`)`]{#dom-context-2d-putimagedata
.dfn dfn-for="CanvasImageData" dfn-type="method"} method steps are to
[put pixels from an `ImageData` onto a
bitmap](#dom-context2d-putimagedata-common){#pixel-manipulation:dom-context2d-putimagedata-common-2},
given `imageData`{.variable},
[this](https://webidl.spec.whatwg.org/#this){#pixel-manipulation:this-8
x-internal="this"}\'s [output
bitmap](#output-bitmap){#pixel-manipulation:output-bitmap-4},
`dx`{.variable}, `dy`{.variable}, `dirtyX`{.variable},
`dirtyY`{.variable}, `dirtyWidth`{.variable}, and
`dirtyHeight`{.variable}.

To [put pixels from an `ImageData` onto a
bitmap]{#dom-context2d-putimagedata-common .dfn}, given an
[`ImageData`{#pixel-manipulation:imagedata-8}](imagebitmap-and-animations.html#imagedata)
`imageData`{.variable}, an [output
bitmap](#output-bitmap){#pixel-manipulation:output-bitmap-5}
`bitmap`{.variable}, and numbers `dx`{.variable}, `dy`{.variable},
`dirtyX`{.variable}, `dirtyY`{.variable}, `dirtyWidth`{.variable}, and
`dirtyHeight`{.variable}:

1.  Let `buffer`{.variable} be `imageData`{.variable}\'s
    [`data`{#pixel-manipulation:dom-imagedata-data-2}](imagebitmap-and-animations.html#dom-imagedata-data)
    attribute value\'s \[\[ViewedArrayBuffer\]\] internal slot.

2.  If
    [IsDetachedBuffer](https://tc39.es/ecma262/#sec-isdetachedbuffer){#pixel-manipulation:isdetachedbuffer
    x-internal="isdetachedbuffer"}(`buffer`{.variable}) is true, then
    throw an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#pixel-manipulation:invalidstateerror-2
    x-internal="invalidstateerror"}
    [`DOMException`{#pixel-manipulation:domexception-7}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

3.  If `dirtyWidth`{.variable} is negative, then let `dirtyX`{.variable}
    be `dirtyX`{.variable}+`dirtyWidth`{.variable}, and let
    `dirtyWidth`{.variable} be equal to the absolute magnitude of
    `dirtyWidth`{.variable}.

    If `dirtyHeight`{.variable} is negative, then let
    `dirtyY`{.variable} be `dirtyY`{.variable}+`dirtyHeight`{.variable},
    and let `dirtyHeight`{.variable} be equal to the absolute magnitude
    of `dirtyHeight`{.variable}.

4.  If `dirtyX`{.variable} is negative, then let `dirtyWidth`{.variable}
    be `dirtyWidth`{.variable}+`dirtyX`{.variable}, and let
    `dirtyX`{.variable} be zero.

    If `dirtyY`{.variable} is negative, then let
    `dirtyHeight`{.variable} be
    `dirtyHeight`{.variable}+`dirtyY`{.variable}, and let
    `dirtyY`{.variable} be zero.

5.  If `dirtyX`{.variable}+`dirtyWidth`{.variable} is greater than the
    [`width`{#pixel-manipulation:dom-imagedata-width-3}](imagebitmap-and-animations.html#dom-imagedata-width)
    attribute of the `imageData`{.variable} argument, then let
    `dirtyWidth`{.variable} be the value of that
    [`width`{#pixel-manipulation:dom-imagedata-width-4}](imagebitmap-and-animations.html#dom-imagedata-width)
    attribute, minus the value of `dirtyX`{.variable}.

    If `dirtyY`{.variable}+`dirtyHeight`{.variable} is greater than the
    [`height`{#pixel-manipulation:dom-imagedata-height-3}](imagebitmap-and-animations.html#dom-imagedata-height)
    attribute of the `imageData`{.variable} argument, then let
    `dirtyHeight`{.variable} be the value of that
    [`height`{#pixel-manipulation:dom-imagedata-height-4}](imagebitmap-and-animations.html#dom-imagedata-height)
    attribute, minus the value of `dirtyY`{.variable}.

6.  If, after those changes, either `dirtyWidth`{.variable} or
    `dirtyHeight`{.variable} are negative or zero, then return without
    affecting any bitmaps.

7.  For all integer values of `x`{.variable} and `y`{.variable} where
    `dirtyX`{.variable} ≤ `x`{.variable} \< `dirtyX`{.variable}+`dirtyWidth`{.variable}
    and
    `dirtyY`{.variable} ≤ `y`{.variable} \< `dirtyY`{.variable}+`dirtyHeight`{.variable},
    set the pixel with coordinate (`dx`{.variable}+`x`{.variable},
    `dy`{.variable}+`y`{.variable}) in `bitmap`{.variable} to the color
    of the pixel at coordinate (`x`{.variable}, `y`{.variable}) in the
    `imageData`{.variable} data structure\'s
    [bitmap](imagebitmap-and-animations.html#concept-imagedata-bitmap-representation){#pixel-manipulation:concept-imagedata-bitmap-representation},
    converted from `imageData`{.variable}\'s
    [`colorSpace`{#pixel-manipulation:dom-imagedata-colorspace-3}](imagebitmap-and-animations.html#dom-imagedata-colorspace)
    to the [color
    space](#concept-canvas-color-space){#pixel-manipulation:concept-canvas-color-space-6}
    of `bitmap`{.variable} using
    [\'relative-colorimetric\'](https://drafts.csswg.org/css-color-5/#valdef-color-profile-rendering-intent-relative-colorimetric){#pixel-manipulation:'relative-colorimetric'-2
    x-internal="'relative-colorimetric'"} rendering intent.

Due to the lossy nature of converting between color spaces and
converting to and from [premultiplied
alpha](#concept-premultiplied-alpha){#pixel-manipulation:concept-premultiplied-alpha}
color values, pixels that have just been set using
[`putImageData()`{#pixel-manipulation:dom-context-2d-putimagedata}](#dom-context-2d-putimagedata),
and are not completely opaque, might be returned to an equivalent
[`getImageData()`{#pixel-manipulation:dom-context-2d-getimagedata}](#dom-context-2d-getimagedata)
as different values.

The current path, [transformation
matrix](#transformations){#pixel-manipulation:transformations}, [shadow
attributes](#shadows){#pixel-manipulation:shadows-2}, [global
alpha](#concept-canvas-global-alpha){#pixel-manipulation:concept-canvas-global-alpha},
the [clipping
region](#clipping-region){#pixel-manipulation:clipping-region}, and
[current compositing and blending
operator](#current-compositing-and-blending-operator){#pixel-manipulation:current-compositing-and-blending-operator}
must not affect the methods described in this section.

::: example
In the following example, the script generates an
[`ImageData`{#pixel-manipulation:imagedata-9}](imagebitmap-and-animations.html#imagedata)
object so that it can draw onto it.

``` js
// canvas is a reference to a <canvas> element
var context = canvas.getContext('2d');

// create a blank slate
var data = context.createImageData(canvas.width, canvas.height);

// create some plasma
FillPlasma(data, 'green'); // green plasma

// add a cloud to the plasma
AddCloud(data, data.width/2, data.height/2); // put a cloud in the middle

// paint the plasma+cloud on the canvas
context.putImageData(data, 0, 0);

// support methods
function FillPlasma(data, color) { ... }
function AddCloud(data, x, y) { ... }
```
:::

::: example
Here is an example of using
[`getImageData()`{#pixel-manipulation:dom-context-2d-getimagedata-2}](#dom-context-2d-getimagedata)
and
[`putImageData()`{#pixel-manipulation:dom-context-2d-putimagedata-2}](#dom-context-2d-putimagedata)
to implement an edge detection filter.

``` html
<!DOCTYPE HTML>
<html lang="en">
 <head>
  <title>Edge detection demo</title>
  <script>
   var image = new Image();
   function init() {
     image.onload = demo;
     image.src = "image.jpeg";
   }
   function demo() {
     var canvas = document.getElementsByTagName('canvas')[0];
     var context = canvas.getContext('2d');

     // draw the image onto the canvas
     context.drawImage(image, 0, 0);

     // get the image data to manipulate
     var input = context.getImageData(0, 0, canvas.width, canvas.height);

     // get an empty slate to put the data into
     var output = context.createImageData(canvas.width, canvas.height);

     // alias some variables for convenience
     // In this case input.width and input.height
     // match canvas.width and canvas.height
     // but we'll use the former to keep the code generic.
     var w = input.width, h = input.height;
     var inputData = input.data;
     var outputData = output.data;

     // edge detection
     for (var y = 1; y < h-1; y += 1) {
       for (var x = 1; x < w-1; x += 1) {
         for (var c = 0; c < 3; c += 1) {
           var i = (y*w + x)*4 + c;
           outputData[i] = 127 + -inputData[i - w*4 - 4] -   inputData[i - w*4] - inputData[i - w*4 + 4] +
                                 -inputData[i - 4]       + 8*inputData[i]       - inputData[i + 4] +
                                 -inputData[i + w*4 - 4] -   inputData[i + w*4] - inputData[i + w*4 + 4];
         }
         outputData[(y*w + x)*4 + 3] = 255; // alpha
       }
     }

     // put the image data back after manipulation
     context.putImageData(output, 0, 0);
   }
  </script>
 </head>
 <body onload="init()">
  <canvas></canvas>
 </body>
</html>
```
:::

::: example
Here is an example of color space conversion applied when drawing a
solid color and reading the result back using and
[`getImageData()`{#pixel-manipulation:dom-context-2d-getimagedata-3}](#dom-context-2d-getimagedata).

``` html
<!DOCTYPE HTML>
<html lang="en">
<title>Color space image data demo</title>

<canvas></canvas>

<script>
const canvas = document.querySelector('canvas');
const context = canvas.getContext('2d', {colorSpace:'display-p3'});

// Draw a red rectangle. Note that the hex color notation
// specifies sRGB colors.
context.fillStyle = "#FF0000";
context.fillRect(0, 0, 64, 64);

// Get the image data.
const pixels = context.getImageData(0, 0, 1, 1);

// This will print 'display-p3', reflecting the default behavior
// of returning image data in the canvas's color space.
console.log(pixels.colorSpace);

// This will print the values 234, 51, and 35, reflecting the
// red fill color, converted to 'display-p3'.
console.log(pixels.data[0]);
console.log(pixels.data[1]);
console.log(pixels.data[2]);
</script>
```
:::

###### [4.12.5.1.17]{.secno} Compositing[](#compositing){.self-link}

`context`{.variable}`.`[`globalAlpha`](#dom-context-2d-globalalpha){#dom-context-2d-globalalpha-dev}` [ = ``value`{.variable}` ]`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/globalAlpha](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/globalAlpha "The CanvasRenderingContext2D.globalAlpha property of the Canvas 2D API specifies the alpha (transparency) value that is applied to shapes and images before they are drawn onto the canvas.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari2+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns the current [global
alpha](#concept-canvas-global-alpha){#compositing:concept-canvas-global-alpha}
value applied to rendering operations.

Can be set, to change the [global
alpha](#concept-canvas-global-alpha){#compositing:concept-canvas-global-alpha-2}
value. Values outside of the range 0.0 .. 1.0 are ignored.

`context`{.variable}`.`[`globalCompositeOperation`](#dom-context-2d-globalcompositeoperation){#dom-context-2d-globalcompositeoperation-dev}` [ = ``value`{.variable}` ]`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/globalCompositeOperation](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/globalCompositeOperation "The CanvasRenderingContext2D.globalCompositeOperation property of the Canvas 2D API sets the type of compositing operation to apply when drawing new shapes.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari2+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera9+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

Returns the [current compositing and blending
operator](#current-compositing-and-blending-operator){#compositing:current-compositing-and-blending-operator},
from the values defined in Compositing and Blending.
[\[COMPOSITE\]](references.html#refsCOMPOSITE)

Can be set, to change the [current compositing and blending
operator](#current-compositing-and-blending-operator){#compositing:current-compositing-and-blending-operator-2}.
Unknown values are ignored.

Objects that implement the
[`CanvasCompositing`{#compositing:canvascompositing}](#canvascompositing)
interface have a [global
alpha](#concept-canvas-global-alpha){#compositing:concept-canvas-global-alpha-3}
value and a [current compositing and blending
operator](#current-compositing-and-blending-operator){#compositing:current-compositing-and-blending-operator-3}
value that both affect all the drawing operations on this object.

The [global alpha]{#concept-canvas-global-alpha .dfn} value gives an
alpha value that is applied to shapes and images before they are
composited onto the [output
bitmap](#output-bitmap){#compositing:output-bitmap}. The value ranges
from 0.0 (fully transparent) to 1.0 (no additional transparency). It
must initially have the value 1.0.

The [`globalAlpha`]{#dom-context-2d-globalalpha .dfn
dfn-for="CanvasCompositing" dfn-type="attribute"} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#compositing:this
x-internal="this"}\'s [global
alpha](#concept-canvas-global-alpha){#compositing:concept-canvas-global-alpha-4}.

The
[`globalAlpha`{#compositing:dom-context-2d-globalalpha}](#dom-context-2d-globalalpha)
setter steps are:

1.  If the given value is either infinite, NaN, or not in the range 0.0
    to 1.0, then return.

2.  Otherwise, set
    [this](https://webidl.spec.whatwg.org/#this){#compositing:this-2
    x-internal="this"}\'s [global
    alpha](#concept-canvas-global-alpha){#compositing:concept-canvas-global-alpha-5}
    to the given value.

The [current compositing and blending
operator]{#current-compositing-and-blending-operator .dfn} value
controls how shapes and images are drawn onto the [output
bitmap](#output-bitmap){#compositing:output-bitmap-2}, once they have
had the [global
alpha](#concept-canvas-global-alpha){#compositing:concept-canvas-global-alpha-6}
and the [current transformation
matrix](#current-transformation-matrix){#compositing:current-transformation-matrix}
applied. Initially, it must be set to
\"[`source-over`{#compositing:gcop-source-over}](https://drafts.fxtf.org/compositing/#porterduffcompositingoperators_srcover){x-internal="gcop-source-over"}\".

The
[`globalCompositeOperation`]{#dom-context-2d-globalcompositeoperation
.dfn dfn-for="CanvasCompositing" dfn-type="attribute"} getter steps are
to return
[this](https://webidl.spec.whatwg.org/#this){#compositing:this-3
x-internal="this"}\'s [current compositing and blending
operator](#current-compositing-and-blending-operator){#compositing:current-compositing-and-blending-operator-4}.

The
[`globalCompositeOperation`{#compositing:dom-context-2d-globalcompositeoperation}](#dom-context-2d-globalcompositeoperation)
setter steps are:

1.  If the given value is not [identical
    to](https://infra.spec.whatwg.org/#string-is){#compositing:identical-to
    x-internal="identical-to"} any of the values that the
    [\<blend-mode\>](https://drafts.fxtf.org/compositing/#ltblendmodegt){#compositing:blend-mode
    x-internal="blend-mode"} or the
    [\<composite-mode\>](https://drafts.fxtf.org/compositing/#compositemode){#compositing:composite-mode
    x-internal="composite-mode"} properties are defined to take, then
    return. [\[COMPOSITE\]](references.html#refsCOMPOSITE)

2.  Otherwise, set
    [this](https://webidl.spec.whatwg.org/#this){#compositing:this-4
    x-internal="this"}\'s [current compositing and blending
    operator](#current-compositing-and-blending-operator){#compositing:current-compositing-and-blending-operator-5}
    to the given value.

###### [4.12.5.1.18]{.secno} Image smoothing[](#image-smoothing){.self-link}

`context`{.variable}`.`[`imageSmoothingEnabled`](#dom-context-2d-imagesmoothingenabled){#dom-context-2d-imagesmoothingenabled-dev}` [ = ``value`{.variable}` ]`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/imageSmoothingEnabled](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/imageSmoothingEnabled "The imageSmoothingEnabled property of the CanvasRenderingContext2D interface, part of the Canvas API, determines whether scaled images are smoothed (true, default) or not (false). On getting the imageSmoothingEnabled property, the last value it was set to is returned.")

Support in all current engines.

::: support
[Firefox51+]{.firefox .yes}[Safari9.1+]{.safari .yes}[Chrome30+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)15+]{.edge .yes}[Internet Explorer[🔰
11]{title="Requires a prefix or alternative name."}]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

Returns whether pattern fills and the
[`drawImage()`{#image-smoothing:dom-context-2d-drawimage}](#dom-context-2d-drawimage)
method will attempt to smooth images if their pixels don\'t line up
exactly with the display, when scaling images up.

Can be set, to change whether images are smoothed (true) or not (false).

`context`{.variable}`.`[`imageSmoothingQuality`](#dom-context-2d-imagesmoothingquality){#dom-context-2d-imagesmoothingquality-dev}` [ = ``value`{.variable}` ]`

::::: {.mdn-anno .wrapped .before}
MDN

:::: feature
[CanvasRenderingContext2D/imageSmoothingQuality](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/imageSmoothingQuality "The imageSmoothingQuality property of the CanvasRenderingContext2D interface, part of the Canvas API, lets you set the quality of image smoothing.")

::: support
[FirefoxNo]{.firefox .no}[Safari9.1+]{.safari .yes}[Chrome54+]{.chrome
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

Returns the current image-smoothing-quality preference.

Can be set, to change the preferred quality of image smoothing. The
possible values are
\"[`low`{#image-smoothing:dom-context-2d-imagesmoothingquality-low}](#dom-context-2d-imagesmoothingquality-low)\",
\"[`medium`{#image-smoothing:dom-context-2d-imagesmoothingquality-medium}](#dom-context-2d-imagesmoothingquality-medium)\"
and
\"[`high`{#image-smoothing:dom-context-2d-imagesmoothingquality-high}](#dom-context-2d-imagesmoothingquality-high)\".
Unknown values are ignored.

Objects that implement the
[`CanvasImageSmoothing`{#image-smoothing:canvasimagesmoothing}](#canvasimagesmoothing)
interface have attributes that control how image smoothing is performed.

The [`imageSmoothingEnabled`]{#dom-context-2d-imagesmoothingenabled .dfn
dfn-for="CanvasImageSmoothing" dfn-type="attribute"} attribute, on
getting, must return the last value it was set to. On setting, it must
be set to the new value. When the object implementing the
[`CanvasImageSmoothing`{#image-smoothing:canvasimagesmoothing-2}](#canvasimagesmoothing)
interface is created, the attribute must be set to true.

The [`imageSmoothingQuality`]{#dom-context-2d-imagesmoothingquality .dfn
dfn-for="CanvasImageSmoothing" dfn-type="attribute"} attribute, on
getting, must return the last value it was set to. On setting, it must
be set to the new value. When the object implementing the
[`CanvasImageSmoothing`{#image-smoothing:canvasimagesmoothing-3}](#canvasimagesmoothing)
interface is created, the attribute must be set to
\"[`low`{#image-smoothing:dom-context-2d-imagesmoothingquality-low-2}](#dom-context-2d-imagesmoothingquality-low)\".

###### [4.12.5.1.19]{.secno} [Shadows]{.dfn}[](#shadows){.self-link}

All drawing operations on an object which implements the
[`CanvasShadowStyles`{#shadows:canvasshadowstyles}](#canvasshadowstyles)
interface are affected by the four global shadow attributes.

`context`{.variable}`.`[`shadowColor`](#dom-context-2d-shadowcolor){#dom-context-2d-shadowcolor-dev}` [ = ``value`{.variable}` ]`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/shadowColor](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/shadowColor "The CanvasRenderingContext2D.shadowColor property of the Canvas 2D API specifies the color of shadows.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari2+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns the current shadow color.

Can be set, to change the shadow color. Values that cannot be parsed as
CSS colors are ignored.

`context`{.variable}`.`[`shadowOffsetX`](#dom-context-2d-shadowoffsetx){#dom-context-2d-shadowoffsetx-dev}` [ = ``value`{.variable}` ]`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/shadowOffsetX](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/shadowOffsetX "The CanvasRenderingContext2D.shadowOffsetX property of the Canvas 2D API specifies the distance that shadows will be offset horizontally.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari2+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

`context`{.variable}`.`[`shadowOffsetY`](#dom-context-2d-shadowoffsety){#dom-context-2d-shadowoffsety-dev}` [ = ``value`{.variable}` ]`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/shadowOffsetY](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/shadowOffsetY "The CanvasRenderingContext2D.shadowOffsetY property of the Canvas 2D API specifies the distance that shadows will be offset vertically.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari2+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns the current shadow offset.

Can be set, to change the shadow offset. Values that are not finite
numbers are ignored.

`context`{.variable}`.`[`shadowBlur`](#dom-context-2d-shadowblur){#dom-context-2d-shadowblur-dev}` [ = ``value`{.variable}` ]`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[CanvasRenderingContext2D/shadowBlur](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/shadowBlur "The CanvasRenderingContext2D.shadowBlur property of the Canvas 2D API specifies the amount of blur applied to shadows. The default is 0 (no blur).")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari2+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns the current level of blur applied to shadows.

Can be set, to change the blur level. Values that are not finite numbers
greater than or equal to zero are ignored.

Objects which implement the
[`CanvasShadowStyles`{#shadows:canvasshadowstyles-2}](#canvasshadowstyles)
interface have an associated [shadow
color]{#concept-canvasshadowstyles-shadow-color .dfn}, which is a CSS
color. Initially, it must be [transparent
black](https://drafts.csswg.org/css-color/#transparent-black){#shadows:transparent-black
x-internal="transparent-black"}.

The [`shadowColor`]{#dom-context-2d-shadowcolor .dfn
dfn-for="CanvasShadowStyles" dfn-type="attribute"} getter steps are to
return the
[serialization](https://drafts.csswg.org/css-color/#serializing-color-values){#shadows:serialisation-of-a-color
x-internal="serialisation-of-a-color"} of
[this](https://webidl.spec.whatwg.org/#this){#shadows:this
x-internal="this"}\'s [shadow
color](#concept-canvasshadowstyles-shadow-color){#shadows:concept-canvasshadowstyles-shadow-color}
with [HTML-compatible serialization
requested](https://drafts.csswg.org/css-color/#color-serialization-html-compatible-serialization-is-requested){#shadows:html-compatible-serialization-is-requested
x-internal="html-compatible-serialization-is-requested"}.

The
[`shadowColor`{#shadows:dom-context-2d-shadowcolor}](#dom-context-2d-shadowcolor)
setter steps are:

1.  Let `context`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#shadows:this-2
    x-internal="this"}\'s
    [`canvas`{#shadows:dom-context-2d-canvas}](#dom-context-2d-canvas)
    attribute\'s value, if that is an element; otherwise null.

2.  Let `parsedValue`{.variable} be the result of
    [parsing](https://drafts.csswg.org/css-color/#parse-a-css-color-value){#shadows:parsed-as-a-css-color-value
    x-internal="parsed-as-a-css-color-value"} the given value with
    `context`{.variable} if non-null.

3.  If `parsedValue`{.variable} is failure, then return.

4.  Set [this](https://webidl.spec.whatwg.org/#this){#shadows:this-3
    x-internal="this"}\'s [shadow
    color](#concept-canvasshadowstyles-shadow-color){#shadows:concept-canvasshadowstyles-shadow-color-2}
    to `parsedValue`{.variable}.

The [`shadowOffsetX`]{#dom-context-2d-shadowoffsetx .dfn
dfn-for="CanvasShadowStyles" dfn-type="attribute"} and
[`shadowOffsetY`]{#dom-context-2d-shadowoffsety .dfn
dfn-for="CanvasShadowStyles" dfn-type="attribute"} attributes specify
the distance that the shadow will be offset in the positive horizontal
and positive vertical distance respectively. Their values are in
coordinate space units. They are not affected by the current
transformation matrix.

When the context is created, the shadow offset attributes must initially
have the value 0.

On getting, they must return their current value. On setting, the
attribute being set must be set to the new value, except if the value is
infinite or NaN, in which case the new value must be ignored.

The [`shadowBlur`]{#dom-context-2d-shadowblur .dfn
dfn-for="CanvasShadowStyles" dfn-type="attribute"} attribute specifies
the level of the blurring effect. (The units do not map to coordinate
space units, and are not affected by the current transformation matrix.)

When the context is created, the
[`shadowBlur`{#shadows:dom-context-2d-shadowblur}](#dom-context-2d-shadowblur)
attribute must initially have the value 0.

On getting, the attribute must return its current value. On setting, the
attribute must be set to the new value, except if the value is negative,
infinite or NaN, in which case the new value must be ignored.

[Shadows are only drawn if]{#when-shadows-are-drawn .dfn} the opacity
component of the alpha component of the [shadow
color](#concept-canvasshadowstyles-shadow-color){#shadows:concept-canvasshadowstyles-shadow-color-3}
is nonzero and either the
[`shadowBlur`{#shadows:dom-context-2d-shadowblur-2}](#dom-context-2d-shadowblur)
is nonzero, or the
[`shadowOffsetX`{#shadows:dom-context-2d-shadowoffsetx}](#dom-context-2d-shadowoffsetx)
is nonzero, or the
[`shadowOffsetY`{#shadows:dom-context-2d-shadowoffsety}](#dom-context-2d-shadowoffsety)
is nonzero.

[When shadows are
drawn](#when-shadows-are-drawn){#shadows:when-shadows-are-drawn}, they
must be rendered as follows:

1.  Let `A`{.variable} be an infinite [transparent
    black](https://drafts.csswg.org/css-color/#transparent-black){#shadows:transparent-black-2
    x-internal="transparent-black"} bitmap on which the source image for
    which a shadow is being created has been rendered.

2.  Let `B`{.variable} be an infinite [transparent
    black](https://drafts.csswg.org/css-color/#transparent-black){#shadows:transparent-black-3
    x-internal="transparent-black"} bitmap, with a coordinate space and
    an origin identical to `A`{.variable}.

3.  Copy the alpha component of `A`{.variable} to `B`{.variable}, offset
    by
    [`shadowOffsetX`{#shadows:dom-context-2d-shadowoffsetx-2}](#dom-context-2d-shadowoffsetx)
    in the positive `x`{.variable} direction, and
    [`shadowOffsetY`{#shadows:dom-context-2d-shadowoffsety-2}](#dom-context-2d-shadowoffsety)
    in the positive `y`{.variable} direction.

4.  If
    [`shadowBlur`{#shadows:dom-context-2d-shadowblur-3}](#dom-context-2d-shadowblur)
    is greater than 0:

    1.  Let `σ`{.variable} be half the value of
        [`shadowBlur`{#shadows:dom-context-2d-shadowblur-4}](#dom-context-2d-shadowblur).

    2.  Perform a 2D Gaussian Blur on `B`{.variable}, using
        `σ`{.variable} as the standard deviation.

    User agents may limit values of `σ`{.variable} to an
    implementation-specific maximum value to avoid exceeding hardware
    limitations during the Gaussian blur operation.

5.  Set the red, green, and blue components of every pixel in
    `B`{.variable} to the red, green, and blue components (respectively)
    of the [shadow
    color](#concept-canvasshadowstyles-shadow-color){#shadows:concept-canvasshadowstyles-shadow-color-4}.

6.  Multiply the alpha component of every pixel in `B`{.variable} by the
    alpha component of the [shadow
    color](#concept-canvasshadowstyles-shadow-color){#shadows:concept-canvasshadowstyles-shadow-color-5}.

7.  The shadow is in the bitmap `B`{.variable}, and is rendered as part
    of the [drawing model](#drawing-model){#shadows:drawing-model}
    described below.

If the [current compositing and blending
operator](#current-compositing-and-blending-operator){#shadows:current-compositing-and-blending-operator}
is
\"[`copy`{#shadows:gcop-copy}](https://drafts.fxtf.org/compositing/#porterduffcompositingoperators_src){x-internal="gcop-copy"}\",
then shadows effectively won\'t render (since the shape will overwrite
the shadow).

###### [4.12.5.1.20]{.secno} Filters[](#filters){.self-link}

All drawing operations on an object which implements the
[`CanvasFilters`{#filters:canvasfilters}](#canvasfilters) interface are
affected by the global [`filter`]{#dom-context-2d-filter .dfn
dfn-for="CanvasFilters" dfn-type="attribute"} attribute.

`context`{.variable}`.`[`filter`](#dom-context-2d-filter){#dom-context-2d-filter-dev}` [ = ``value`{.variable}` ]`

::::: {.mdn-anno .wrapped .before}
MDN

:::: feature
[CanvasRenderingContext2D/filter](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/filter "The CanvasRenderingContext2D.filter property of the Canvas 2D API provides filter effects such as blurring and grayscaling. It is similar to the CSS filter property and accepts the same values.")

::: support
[Firefox49+]{.firefox .yes}[SafariNo]{.safari .no}[Chrome52+]{.chrome
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

Returns the current filter.

Can be set, to change the filter. Values can either be the string
\"`none`\" or a string parseable as a
[\<filter-value-list\>](https://drafts.fxtf.org/filter-effects/#typedef-filter-value-list){#filters:filter-value-list
x-internal="filter-value-list"}. Other values are ignored.

Such objects have an associated [current
filter]{#concept-canvas-current-filter .dfn}, which is a string.
Initially the [current
filter](#concept-canvas-current-filter){#filters:concept-canvas-current-filter}
is set to the string \"`none`\". Whenever the value of the [current
filter](#concept-canvas-current-filter){#filters:concept-canvas-current-filter-2}
is the string \"`none`\" filters will be disabled for the context.

The [`filter`{#filters:dom-context-2d-filter}](#dom-context-2d-filter)
getter steps are to return
[this](https://webidl.spec.whatwg.org/#this){#filters:this
x-internal="this"}\'s [current
filter](#concept-canvas-current-filter){#filters:concept-canvas-current-filter-3}.

The [`filter`{#filters:dom-context-2d-filter-2}](#dom-context-2d-filter)
setter steps are:

1.  If the given value is \"`none`\", then set
    [this](https://webidl.spec.whatwg.org/#this){#filters:this-2
    x-internal="this"}\'s [current
    filter](#concept-canvas-current-filter){#filters:concept-canvas-current-filter-4}
    to \"`none`\" and return.

2.  Let `parsedValue`{.variable} be the result of
    [parsing](https://drafts.csswg.org/css-syntax/#parse-grammar){#filters:parse-something-according-to-a-css-grammar
    x-internal="parse-something-according-to-a-css-grammar"} the given
    values as a
    [\<filter-value-list\>](https://drafts.fxtf.org/filter-effects/#typedef-filter-value-list){#filters:filter-value-list-2
    x-internal="filter-value-list"}. If any property-independent style
    sheet syntax like \'inherit\' or \'initial\' is present, then this
    parsing must return failure.

3.  If `parsedValue`{.variable} is failure, then return.

4.  Set [this](https://webidl.spec.whatwg.org/#this){#filters:this-3
    x-internal="this"}\'s [current
    filter](#concept-canvas-current-filter){#filters:concept-canvas-current-filter-5}
    to the given value.

Though
`context`{.variable}`.`[`filter`{#filters:dom-context-2d-filter-3}](#dom-context-2d-filter)` = "``none``"`
will disable filters for the context,
`context`{.variable}`.`[`filter`{#filters:dom-context-2d-filter-4}](#dom-context-2d-filter)` = ""`,
`context`{.variable}`.`[`filter`{#filters:dom-context-2d-filter-5}](#dom-context-2d-filter)` = null`,
and
`context`{.variable}`.`[`filter`{#filters:dom-context-2d-filter-6}](#dom-context-2d-filter)` = undefined`
are all treated as unparseable inputs and the value of the [current
filter](#concept-canvas-current-filter){#filters:concept-canvas-current-filter-6}
is left unchanged.

Coordinates used in the value of the [current
filter](#concept-canvas-current-filter){#filters:concept-canvas-current-filter-7}
are interpreted such that one pixel is equivalent to one SVG user space
unit and to one canvas coordinate space unit. Filter coordinates are not
affected by the [current transformation
matrix](#transformations){#filters:transformations}. The current
transformation matrix affects only the input to the filter. Filters are
applied in the [output
bitmap](#output-bitmap){#filters:output-bitmap}\'s coordinate space.

When the value of the [current
filter](#concept-canvas-current-filter){#filters:concept-canvas-current-filter-8}
is a string parsable as a
[\<filter-value-list\>](https://drafts.fxtf.org/filter-effects/#typedef-filter-value-list){#filters:filter-value-list-3
x-internal="filter-value-list"} which defines lengths using percentages
or using [\'em\'](https://drafts.csswg.org/css-values/#em){#filters:'em'
x-internal="'em'"} or
[\'ex\'](https://drafts.csswg.org/css-values/#ex){#filters:'ex'
x-internal="'ex'"} units, these must be interpreted relative to the
[computed
value](https://drafts.csswg.org/css-cascade/#computed-value){#filters:computed-value
x-internal="computed-value"} of the
[\'font-size\'](https://drafts.csswg.org/css-fonts/#font-size-prop){#filters:'font-size'
x-internal="'font-size'"} property of the [font style source
object](#font-style-source-object){#filters:font-style-source-object} at
the time that the attribute is set. If the [computed
values](https://drafts.csswg.org/css-cascade/#computed-value){#filters:computed-value-2
x-internal="computed-value"} are undefined for a particular case (e.g.
because the [font style source
object](#font-style-source-object){#filters:font-style-source-object-2}
is not an element or is not [being
rendered](rendering.html#being-rendered){#filters:being-rendered}), then
the relative keywords must be interpreted relative to the default value
of the [`font`{#filters:dom-context-2d-font}](#dom-context-2d-font)
attribute. The \'larger\' and \'smaller\' keywords are not supported.

If the value of the [current
filter](#concept-canvas-current-filter){#filters:concept-canvas-current-filter-9}
is a string parseable as a
[\<filter-value-list\>](https://drafts.fxtf.org/filter-effects/#typedef-filter-value-list){#filters:filter-value-list-4
x-internal="filter-value-list"} with a reference to an SVG filter in the
same document, and this SVG filter changes, then the changed filter is
used for the next draw operation.

If the value of the [current
filter](#concept-canvas-current-filter){#filters:concept-canvas-current-filter-10}
is a string parseable as a
[\<filter-value-list\>](https://drafts.fxtf.org/filter-effects/#typedef-filter-value-list){#filters:filter-value-list-5
x-internal="filter-value-list"} with a reference to an SVG filter in an
external resource document and that document is not loaded when a
drawing operation is invoked, then the drawing operation must proceed
with no filtering.

###### [4.12.5.1.21]{.secno} Working with externally-defined SVG filters[](#working-with-externally-defined-svg-filters){.self-link}

*This section is non-normative.*

Since drawing is performed using filter value \"`none`\" until an
externally-defined filter has finished loading, authors might wish to
determine whether such a filter has finished loading before proceeding
with a drawing operation. One way to accomplish this is to load the
externally-defined filter elsewhere within the same page in some element
that sends a `load` event (for example, an [SVG
`use`](https://svgwg.org/svg2-draft/struct.html#UseElement){#working-with-externally-defined-svg-filters:svg-use
x-internal="svg-use"} element), and wait for the `load` event to be
dispatched.

###### [4.12.5.1.22]{.secno} [Drawing model]{.dfn}[](#drawing-model){.self-link}

When a shape or image is painted, user agents must follow these steps,
in the order given (or act as if they do):

1.  Render the shape or image onto an infinite [transparent
    black](https://drafts.csswg.org/css-color/#transparent-black){#drawing-model:transparent-black
    x-internal="transparent-black"} bitmap, creating image
    `A`{.variable}, as described in the previous sections. For shapes,
    the current fill, stroke, and line styles must be honored, and the
    stroke must itself also be subjected to the current transformation
    matrix.

2.  Multiply the alpha component of every pixel in `A`{.variable} by
    [`global alpha`{#drawing-model:concept-canvas-global-alpha}](#concept-canvas-global-alpha).

3.  When the [current
    filter](#concept-canvas-current-filter){#drawing-model:concept-canvas-current-filter}
    is set to a value other than \"`none`\" and all the
    externally-defined filters it references, if any, are in documents
    that are currently loaded, then use image `A`{.variable} as the
    input to the [current
    filter](#concept-canvas-current-filter){#drawing-model:concept-canvas-current-filter-2},
    creating image `B`{.variable}. If the [current
    filter](#concept-canvas-current-filter){#drawing-model:concept-canvas-current-filter-3}
    is a string parseable as a
    [\<filter-value-list\>](https://drafts.fxtf.org/filter-effects/#typedef-filter-value-list){#drawing-model:filter-value-list
    x-internal="filter-value-list"}, then draw using the [current
    filter](#concept-canvas-current-filter){#drawing-model:concept-canvas-current-filter-4}
    in the same manner as SVG.

    Otherwise, let `B`{.variable} be an alias for `A`{.variable}.

4.  [When shadows are
    drawn](#when-shadows-are-drawn){#drawing-model:when-shadows-are-drawn},
    render the shadow from image `B`{.variable}, using the current
    shadow styles, creating image `C`{.variable}.

5.  [When shadows are
    drawn](#when-shadows-are-drawn){#drawing-model:when-shadows-are-drawn-2},
    composite `C`{.variable} within the [clipping
    region](#clipping-region){#drawing-model:clipping-region} over the
    current [output
    bitmap](#output-bitmap){#drawing-model:output-bitmap} using the
    [current compositing and blending
    operator](#current-compositing-and-blending-operator){#drawing-model:current-compositing-and-blending-operator}.

6.  Composite `B`{.variable} within the [clipping
    region](#clipping-region){#drawing-model:clipping-region-2} over the
    current [output
    bitmap](#output-bitmap){#drawing-model:output-bitmap-2} using the
    [current compositing and blending
    operator](#current-compositing-and-blending-operator){#drawing-model:current-compositing-and-blending-operator-2}.

When compositing onto the [output
bitmap](#output-bitmap){#drawing-model:output-bitmap-3}, pixels that
would fall outside of the [output
bitmap](#output-bitmap){#drawing-model:output-bitmap-4} must be
discarded.

###### [4.12.5.1.23]{.secno} Best practices[](#best-practices){.self-link}

When a canvas is interactive, authors should include
[focusable](interaction.html#focusable){#best-practices:focusable}
elements in the element\'s fallback content corresponding to each
[focusable](interaction.html#focusable){#best-practices:focusable-2}
part of the canvas, as in the [example
above](#drawCustomFocusRingExample).

When rendering focus rings, to ensure that focus rings have the
appearance of native focus rings, authors should use the
[`drawFocusIfNeeded()`{#best-practices:dom-context-2d-drawfocusifneeded}](#dom-context-2d-drawfocusifneeded)
method, passing it the element for which a ring is being drawn. This
method only draws the focus ring if the element is
[focused](interaction.html#focused){#best-practices:focused}, so that it
can simply be called whenever drawing the element, without checking
whether the element is focused or not first.

Authors should avoid implementing text editing controls using the
[`canvas`{#best-practices:the-canvas-element}](#the-canvas-element)
element. Doing so has a large number of disadvantages:

- Mouse placement of the caret has to be reimplemented.
- Keyboard movement of the caret has to be reimplemented (possibly
  across lines, for multiline text input).
- Scrolling of the text control has to be implemented (horizontally for
  long lines, vertically for multiline input).
- Native features such as copy-and-paste have to be reimplemented.
- Native features such as spell-checking have to be reimplemented.
- Native features such as drag-and-drop have to be reimplemented.
- Native features such as page-wide text search have to be
  reimplemented.
- Native features specific to the user, for example custom text
  services, have to be reimplemented. This is close to impossible since
  each user might have different services installed, and there is an
  unbounded set of possible such services.
- Bidirectional text editing has to be reimplemented.
- For multiline text editing, line wrapping has to be implemented for
  all relevant languages.
- Text selection has to be reimplemented.
- Dragging of bidirectional text selections has to be reimplemented.
- Platform-native keyboard shortcuts have to be reimplemented.
- Platform-native input method editors (IMEs) have to be reimplemented.
- Undo and redo functionality has to be reimplemented.
- Accessibility features such as magnification following the caret or
  selection have to be reimplemented.

This is a huge amount of work, and authors are most strongly encouraged
to avoid doing any of it by instead using the
[`input`{#best-practices:the-input-element}](input.html#the-input-element)
element, the
[`textarea`{#best-practices:the-textarea-element}](form-elements.html#the-textarea-element)
element, or the
[`contenteditable`{#best-practices:attr-contenteditable}](interaction.html#attr-contenteditable)
attribute.

###### [4.12.5.1.24]{.secno} Examples[](#examples){.self-link}

*This section is non-normative.*

::: example
Here is an example of a script that uses canvas to draw [pretty glowing
lines](data:text/html;charset=utf-8;base64,PCFET0NUWVBFIEhUTUw%2BDQo8aHRtbCBsYW5nPSJlbiI%2BDQogPGhlYWQ%2BDQogIDx0aXRsZT5QcmV0dHkgR2xvd2luZyBMaW5lczwvdGl0bGU%2BDQogPC9oZWFkPg0KIDxib2R5Pg0KPGNhbnZhcyB3aWR0aD0iODAwIiBoZWlnaHQ9IjQ1MCI%2BPC9jYW52YXM%2BDQo8c2NyaXB0Pg0KDQogdmFyIGNvbnRleHQgPSBkb2N1bWVudC5nZXRFbGVtZW50c0J5VGFnTmFtZSgnY2FudmFzJylbMF0uZ2V0Q29udGV4dCgnMmQnKTsNCg0KIHZhciBsYXN0WCA9IGNvbnRleHQuY2FudmFzLndpZHRoICogTWF0aC5yYW5kb20oKTsNCiB2YXIgbGFzdFkgPSBjb250ZXh0LmNhbnZhcy5oZWlnaHQgKiBNYXRoLnJhbmRvbSgpOw0KIHZhciBodWUgPSAwOw0KIGZ1bmN0aW9uIGxpbmUoKSB7DQogICBjb250ZXh0LnNhdmUoKTsNCiAgIGNvbnRleHQudHJhbnNsYXRlKGNvbnRleHQuY2FudmFzLndpZHRoLzIsIGNvbnRleHQuY2FudmFzLmhlaWdodC8yKTsNCiAgIGNvbnRleHQuc2NhbGUoMC45LCAwLjkpOw0KICAgY29udGV4dC50cmFuc2xhdGUoLWNvbnRleHQuY2FudmFzLndpZHRoLzIsIC1jb250ZXh0LmNhbnZhcy5oZWlnaHQvMik7DQogICBjb250ZXh0LmJlZ2luUGF0aCgpOw0KICAgY29udGV4dC5saW5lV2lkdGggPSA1ICsgTWF0aC5yYW5kb20oKSAqIDEwOw0KICAgY29udGV4dC5tb3ZlVG8obGFzdFgsIGxhc3RZKTsNCiAgIGxhc3RYID0gY29udGV4dC5jYW52YXMud2lkdGggKiBNYXRoLnJhbmRvbSgpOw0KICAgbGFzdFkgPSBjb250ZXh0LmNhbnZhcy5oZWlnaHQgKiBNYXRoLnJhbmRvbSgpOw0KICAgY29udGV4dC5iZXppZXJDdXJ2ZVRvKGNvbnRleHQuY2FudmFzLndpZHRoICogTWF0aC5yYW5kb20oKSwNCiAgICAgICAgICAgICAgICAgICAgICAgICBjb250ZXh0LmNhbnZhcy5oZWlnaHQgKiBNYXRoLnJhbmRvbSgpLA0KICAgICAgICAgICAgICAgICAgICAgICAgIGNvbnRleHQuY2FudmFzLndpZHRoICogTWF0aC5yYW5kb20oKSwNCiAgICAgICAgICAgICAgICAgICAgICAgICBjb250ZXh0LmNhbnZhcy5oZWlnaHQgKiBNYXRoLnJhbmRvbSgpLA0KICAgICAgICAgICAgICAgICAgICAgICAgIGxhc3RYLCBsYXN0WSk7DQoNCiAgIGh1ZSA9IGh1ZSArIDEwICogTWF0aC5yYW5kb20oKTsNCiAgIGNvbnRleHQuc3Ryb2tlU3R5bGUgPSAnaHNsKCcgKyBodWUgKyAnLCA1MCUsIDUwJSknOw0KICAgY29udGV4dC5zaGFkb3dDb2xvciA9ICd3aGl0ZSc7DQogICBjb250ZXh0LnNoYWRvd0JsdXIgPSAxMDsNCiAgIGNvbnRleHQuc3Ryb2tlKCk7DQogICBjb250ZXh0LnJlc3RvcmUoKTsNCiB9DQogc2V0SW50ZXJ2YWwobGluZSwgNTApOw0KDQogZnVuY3Rpb24gYmxhbmsoKSB7DQogICBjb250ZXh0LmZpbGxTdHlsZSA9ICdyZ2JhKDAsMCwwLDAuMSknOw0KICAgY29udGV4dC5maWxsUmVjdCgwLCAwLCBjb250ZXh0LmNhbnZhcy53aWR0aCwgY29udGV4dC5jYW52YXMuaGVpZ2h0KTsNCiB9DQogc2V0SW50ZXJ2YWwoYmxhbmssIDQwKTsNCg0KPC9zY3JpcHQ%2BDQogPC9ib2R5Pg0KPC9odG1sPg0K).

``` html
<canvas width="800" height="450"></canvas>
<script>

 var context = document.getElementsByTagName('canvas')[0].getContext('2d');

 var lastX = context.canvas.width * Math.random();
 var lastY = context.canvas.height * Math.random();
 var hue = 0;
 function line() {
   context.save();
   context.translate(context.canvas.width/2, context.canvas.height/2);
   context.scale(0.9, 0.9);
   context.translate(-context.canvas.width/2, -context.canvas.height/2);
   context.beginPath();
   context.lineWidth = 5 + Math.random() * 10;
   context.moveTo(lastX, lastY);
   lastX = context.canvas.width * Math.random();
   lastY = context.canvas.height * Math.random();
   context.bezierCurveTo(context.canvas.width * Math.random(),
                         context.canvas.height * Math.random(),
                         context.canvas.width * Math.random(),
                         context.canvas.height * Math.random(),
                         lastX, lastY);

   hue = hue + 10 * Math.random();
   context.strokeStyle = 'hsl(' + hue + ', 50%, 50%)';
   context.shadowColor = 'white';
   context.shadowBlur = 10;
   context.stroke();
   context.restore();
 }
 setInterval(line, 50);

 function blank() {
   context.fillStyle = 'rgba(0,0,0,0.1)';
   context.fillRect(0, 0, context.canvas.width, context.canvas.height);
 }
 setInterval(blank, 40);

</script>
```
:::

::: example
The 2D rendering context for
[`canvas`{#examples:the-canvas-element}](#the-canvas-element) is often
used for sprite-based games. The following example demonstrates this:

Here is the source for this example:

``` html
<!DOCTYPE HTML>
<html lang="en">
<meta charset="utf-8">
<title>Blue Robot Demo</title>
<style>
  html { overflow: hidden; min-height: 200px; min-width: 380px; }
  body { height: 200px; position: relative; margin: 8px; }
  .buttons { position: absolute; bottom: 0px; left: 0px; margin: 4px; }
</style>
<canvas width="380" height="200"></canvas>
<script>
 var Landscape = function (context, width, height) {
   this.offset = 0;
   this.width = width;
   this.advance = function (dx) {
     this.offset += dx;
   };
   this.horizon = height * 0.7;
   // This creates the sky gradient (from a darker blue to white at the bottom)
   this.sky = context.createLinearGradient(0, 0, 0, this.horizon);
   this.sky.addColorStop(0.0, 'rgb(55,121,179)');
   this.sky.addColorStop(0.7, 'rgb(121,194,245)');
   this.sky.addColorStop(1.0, 'rgb(164,200,214)');
   // this creates the grass gradient (from a darker green to a lighter green)
   this.earth = context.createLinearGradient(0, this.horizon, 0, height);
   this.earth.addColorStop(0.0, 'rgb(81,140,20)');
   this.earth.addColorStop(1.0, 'rgb(123,177,57)');
   this.paintBackground = function (context, width, height) {
     // first, paint the sky and grass rectangles
     context.fillStyle = this.sky;
     context.fillRect(0, 0, width, this.horizon);
     context.fillStyle = this.earth;
     context.fillRect(0, this.horizon, width, height-this.horizon);
     // then, draw the cloudy banner
     // we make it cloudy by having the draw text off the top of the
     // canvas, and just having the blurred shadow shown on the canvas
     context.save();
     context.translate(width-((this.offset+(this.width*3.2)) % (this.width*4.0))+0, 0);
     context.shadowColor = 'white';
     context.shadowOffsetY = 30+this.horizon/3; // offset down on canvas
     context.shadowBlur = '5';
     context.fillStyle = 'white';
     context.textAlign = 'left';
     context.textBaseline = 'top';
     context.font = '20px sans-serif';
     context.fillText('WHATWG ROCKS', 10, -30); // text up above canvas
     context.restore();
     // then, draw the background tree
     context.save();
     context.translate(width-((this.offset+(this.width*0.2)) % (this.width*1.5))+30, 0);
     context.beginPath();
     context.fillStyle = 'rgb(143,89,2)';
     context.lineStyle = 'rgb(10,10,10)';
     context.lineWidth = 2;
     context.rect(0, this.horizon+5, 10, -50); // trunk
     context.fill();
     context.stroke();
     context.beginPath();
     context.fillStyle = 'rgb(78,154,6)';
     context.arc(5, this.horizon-60, 30, 0, Math.PI*2); // leaves
     context.fill();
     context.stroke();
     context.restore();
   };
   this.paintForeground = function (context, width, height) {
     // draw the box that goes in front
     context.save();
     context.translate(width-((this.offset+(this.width*0.7)) % (this.width*1.1))+0, 0);
     context.beginPath();
     context.rect(0, this.horizon - 5, 25, 25);
     context.fillStyle = 'rgb(220,154,94)';
     context.lineStyle = 'rgb(10,10,10)';
     context.lineWidth = 2;
     context.fill();
     context.stroke();
     context.restore();
   };
 };
</script>
<script>
 var BlueRobot = function () {
   this.sprites = new Image();
   this.sprites.src = 'blue-robot.png'; // this sprite sheet has 8 cells
   this.targetMode = 'idle';
   this.walk = function () {
     this.targetMode = 'walk';
   };
   this.stop = function () {
     this.targetMode = 'idle';
   };
   this.frameIndex = {
     'idle': [0], // first cell is the idle frame
     'walk': [1,2,3,4,5,6], // the walking animation is cells 1-6
     'stop': [7], // last cell is the stopping animation
   };
   this.mode = 'idle';
   this.frame = 0; // index into frameIndex
   this.tick = function () {
     // this advances the frame and the robot
     // the return value is how many pixels the robot has moved
     this.frame += 1;
     if (this.frame >= this.frameIndex[this.mode].length) {
       // we've reached the end of this animation cycle
       this.frame = 0;
       if (this.mode != this.targetMode) {
         // switch to next cycle
         if (this.mode == 'walk') {
           // we need to stop walking before we decide what to do next
           this.mode = 'stop';
         } else if (this.mode == 'stop') {
           if (this.targetMode == 'walk')
             this.mode = 'walk';
           else
             this.mode = 'idle';
         } else if (this.mode == 'idle') {
           if (this.targetMode == 'walk')
             this.mode = 'walk';
         }
       }
     }
     if (this.mode == 'walk')
       return 8;
     return 0;
   },
   this.paint = function (context, x, y) {
     if (!this.sprites.complete) return;
     // draw the right frame out of the sprite sheet onto the canvas
     // we assume each frame is as high as the sprite sheet
     // the x,y coordinates give the position of the bottom center of the sprite
     context.drawImage(this.sprites,
                       this.frameIndex[this.mode][this.frame] * this.sprites.height, 0, this.sprites.height, this.sprites.height,
                       x-this.sprites.height/2, y-this.sprites.height, this.sprites.height, this.sprites.height);
   };
 };
</script>
<script>
 var canvas = document.getElementsByTagName('canvas')[0];
 var context = canvas.getContext('2d');
 var landscape = new Landscape(context, canvas.width, canvas.height);
 var blueRobot = new BlueRobot();
 // paint when the browser wants us to, using requestAnimationFrame()
 function paint() {
   context.clearRect(0, 0, canvas.width, canvas.height);
   landscape.paintBackground(context, canvas.width, canvas.height);
   blueRobot.paint(context, canvas.width/2, landscape.horizon*1.1);
   landscape.paintForeground(context, canvas.width, canvas.height);
   requestAnimationFrame(paint);
 }
 paint();
 // but tick every 100ms, so that we don't slow down when we don't paint
 setInterval(function () {
   var dx = blueRobot.tick();
   landscape.advance(dx);
 }, 100);
</script>
<p class="buttons">
 <input type=button value="Walk" onclick="blueRobot.walk()">
 <input type=button value="Stop" onclick="blueRobot.stop()">
<footer>
 <small> Blue Robot Player Sprite by <a href="https://johncolburn.deviantart.com/">JohnColburn</a>.
 Licensed under the terms of the Creative Commons Attribution Share-Alike 3.0 Unported license.</small>
 <small> This work is itself licensed under a <a rel="license" href="https://creativecommons.org/licenses/by-sa/3.0/">Creative
 Commons Attribution-ShareAlike 3.0 Unported License</a>.</small>
</footer>
```
:::

##### [4.12.5.2]{.secno} The [`ImageBitmap`{#the-imagebitmap-rendering-context:imagebitmap}](imagebitmap-and-animations.html#imagebitmap) rendering context[](#the-imagebitmap-rendering-context){.self-link}

###### [4.12.5.2.1]{.secno} Introduction[](#introduction-6){.self-link} {#introduction-6}

[`ImageBitmapRenderingContext`{#introduction-6:imagebitmaprenderingcontext}](#imagebitmaprenderingcontext)
is a performance-oriented interface that provides a low overhead method
for displaying the contents of
[`ImageBitmap`{#introduction-6:imagebitmap}](imagebitmap-and-animations.html#imagebitmap)
objects. It uses transfer semantics to reduce overall memory
consumption. It also streamlines performance by avoiding intermediate
compositing, unlike the
[`drawImage()`{#introduction-6:dom-context-2d-drawimage}](#dom-context-2d-drawimage)
method of
[`CanvasRenderingContext2D`{#introduction-6:canvasrenderingcontext2d}](#canvasrenderingcontext2d).

Using an
[`img`{#introduction-6:the-img-element}](embedded-content.html#the-img-element)
element as an intermediate for getting an image resource into a canvas,
for example, would result in two copies of the decoded image existing in
memory at the same time: the
[`img`{#introduction-6:the-img-element-2}](embedded-content.html#the-img-element)
element\'s copy, and the one in the canvas\'s backing store. This memory
cost can be prohibitive when dealing with extremely large images. This
can be avoided by using
[`ImageBitmapRenderingContext`{#introduction-6:imagebitmaprenderingcontext-2}](#imagebitmaprenderingcontext).

::: example
Using
[`ImageBitmapRenderingContext`{#introduction-6:imagebitmaprenderingcontext-3}](#imagebitmaprenderingcontext),
here is how to transcode an image to the JPEG format in a memory- and
CPU-efficient way:

``` js
createImageBitmap(inputImageBlob).then(image => {
  const canvas = document.createElement('canvas');
  const context = canvas.getContext('bitmaprenderer');
  context.transferFromImageBitmap(image);

  canvas.toBlob(outputJPEGBlob => {
    // Do something with outputJPEGBlob.
  }, 'image/jpeg');
});
```
:::

###### [4.12.5.2.2]{.secno} The [`ImageBitmapRenderingContext`{#the-imagebitmaprenderingcontext-interface:imagebitmaprenderingcontext}](#imagebitmaprenderingcontext) interface[](#the-imagebitmaprenderingcontext-interface){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[ImageBitmapRenderingContext](https://developer.mozilla.org/en-US/docs/Web/API/ImageBitmapRenderingContext "The ImageBitmapRenderingContext interface is a canvas rendering context that provides the functionality to replace the canvas's contents with the given ImageBitmap. Its context id (the first argument to HTMLCanvasElement.getContext() or OffscreenCanvas.getContext()) is "bitmaprenderer".")

Support in all current engines.

::: support
[Firefox46+]{.firefox .yes}[Safari11.1+]{.safari
.yes}[Chrome66+]{.chrome .yes}

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

``` idl
[Exposed=(Window,Worker)]
interface ImageBitmapRenderingContext {
  readonly attribute (HTMLCanvasElement or OffscreenCanvas) canvas;
  undefined transferFromImageBitmap(ImageBitmap? bitmap);
};

dictionary ImageBitmapRenderingContextSettings {
  boolean alpha = true;
};
```

`context`{.variable}` = ``canvas`{.variable}`.`[`getContext`](#dom-canvas-getcontext){#the-imagebitmaprenderingcontext-interface:dom-canvas-getcontext}`('bitmaprenderer' [, { [ `[`alpha`](#dom-imagebitmaprenderingcontextsettings-alpha){#the-imagebitmaprenderingcontext-interface:dom-imagebitmaprenderingcontextsettings-alpha-2}`: false ] } ])`

Returns an
[`ImageBitmapRenderingContext`{#the-imagebitmaprenderingcontext-interface:imagebitmaprenderingcontext-2}](#imagebitmaprenderingcontext)
object that is permanently bound to a particular
[`canvas`{#the-imagebitmaprenderingcontext-interface:the-canvas-element}](#the-canvas-element)
element.

If the
[`alpha`{#the-imagebitmaprenderingcontext-interface:dom-imagebitmaprenderingcontextsettings-alpha-3}](#dom-imagebitmaprenderingcontextsettings-alpha)
setting is provided and set to false, then the canvas is forced to
always be opaque.

`context`{.variable}`.`[`canvas`](#dom-imagebitmaprenderingcontext-canvas){#dom-imagebitmaprenderingcontext-canvas-dev}

Returns the
[`canvas`{#the-imagebitmaprenderingcontext-interface:the-canvas-element-2}](#the-canvas-element)
element that the context is bound to.

`context`{.variable}`.`[`transferFromImageBitmap`](#dom-imagebitmaprenderingcontext-transferfromimagebitmap){#dom-imagebitmaprenderingcontext-transferfromimagebitmap-dev}`(imageBitmap)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[ImageBitmapRenderingContext/transferFromImageBitmap](https://developer.mozilla.org/en-US/docs/Web/API/ImageBitmapRenderingContext/transferFromImageBitmap "The ImageBitmapRenderingContext.transferFromImageBitmap() method displays the given ImageBitmap in the canvas associated with this rendering context. The ownership of the ImageBitmap is transferred to the canvas as well.")

Support in all current engines.

::: support
[Firefox50+]{.firefox .yes}[Safari11.1+]{.safari
.yes}[Chrome66+]{.chrome .yes}

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

Transfers the underlying [bitmap
data](imagebitmap-and-animations.html#concept-imagebitmap-bitmap-data){#the-imagebitmaprenderingcontext-interface:concept-imagebitmap-bitmap-data}
from `imageBitmap`{.variable} to `context`{.variable}, and the bitmap
becomes the contents of the
[`canvas`{#the-imagebitmaprenderingcontext-interface:the-canvas-element-3}](#the-canvas-element)
element to which `context`{.variable} is bound.

`context`{.variable}`.`[`transferFromImageBitmap`](#dom-imagebitmaprenderingcontext-transferfromimagebitmap){#the-imagebitmaprenderingcontext-interface:dom-imagebitmaprenderingcontext-transferfromimagebitmap-2}`(null)`

Replaces contents of the
[`canvas`{#the-imagebitmaprenderingcontext-interface:the-canvas-element-4}](#the-canvas-element)
element to which `context`{.variable} is bound with a [transparent
black](https://drafts.csswg.org/css-color/#transparent-black){#the-imagebitmaprenderingcontext-interface:transparent-black
x-internal="transparent-black"} bitmap whose size corresponds to the
[`width`{#the-imagebitmaprenderingcontext-interface:attr-canvas-width}](#attr-canvas-width)
and
[`height`{#the-imagebitmaprenderingcontext-interface:attr-canvas-height}](#attr-canvas-height)
content attributes of the
[`canvas`{#the-imagebitmaprenderingcontext-interface:the-canvas-element-5}](#the-canvas-element)
element.

The [`canvas`]{#dom-imagebitmaprenderingcontext-canvas .dfn
dfn-for="ImageBitmapRenderingContext" dfn-type="attribute"} attribute
must return the value it was initialized to when the object was created.

An
[`ImageBitmapRenderingContext`{#the-imagebitmaprenderingcontext-interface:imagebitmaprenderingcontext-3}](#imagebitmaprenderingcontext)
object has an [output
bitmap]{#concept-imagebitmaprenderingcontext-output-bitmap .dfn}, which
is a reference to [bitmap
data](imagebitmap-and-animations.html#concept-imagebitmap-bitmap-data){#the-imagebitmaprenderingcontext-interface:concept-imagebitmap-bitmap-data-2}.

An
[`ImageBitmapRenderingContext`{#the-imagebitmaprenderingcontext-interface:imagebitmaprenderingcontext-4}](#imagebitmaprenderingcontext)
object has a [bitmap
mode]{#concept-imagebitmaprenderingcontext-bitmap-mode .dfn}, which can
be set to [valid]{#concept-imagebitmaprenderingcontext-valid .dfn} or
[blank]{#concept-imagebitmaprenderingcontext-blank .dfn}. A value of
[valid](#concept-imagebitmaprenderingcontext-valid){#the-imagebitmaprenderingcontext-interface:concept-imagebitmaprenderingcontext-valid}
indicates that the context\'s [output
bitmap](#concept-imagebitmaprenderingcontext-output-bitmap){#the-imagebitmaprenderingcontext-interface:concept-imagebitmaprenderingcontext-output-bitmap}
refers to [bitmap
data](imagebitmap-and-animations.html#concept-imagebitmap-bitmap-data){#the-imagebitmaprenderingcontext-interface:concept-imagebitmap-bitmap-data-3}
that was acquired via
[`transferFromImageBitmap()`{#the-imagebitmaprenderingcontext-interface:dom-imagebitmaprenderingcontext-transferfromimagebitmap-3}](#dom-imagebitmaprenderingcontext-transferfromimagebitmap).
A value
[blank](#concept-imagebitmaprenderingcontext-blank){#the-imagebitmaprenderingcontext-interface:concept-imagebitmaprenderingcontext-blank}
indicates that the context\'s [output
bitmap](#concept-imagebitmaprenderingcontext-output-bitmap){#the-imagebitmaprenderingcontext-interface:concept-imagebitmaprenderingcontext-output-bitmap-2}
is a default transparent bitmap.

An
[`ImageBitmapRenderingContext`{#the-imagebitmaprenderingcontext-interface:imagebitmaprenderingcontext-5}](#imagebitmaprenderingcontext)
object also has an [alpha]{#concept-imagebitmaprenderingcontext-alpha
.dfn} flag, which can be set to true or false. When an
[`ImageBitmapRenderingContext`{#the-imagebitmaprenderingcontext-interface:imagebitmaprenderingcontext-6}](#imagebitmaprenderingcontext)
object has its
[alpha](#concept-imagebitmaprenderingcontext-alpha){#the-imagebitmaprenderingcontext-interface:concept-imagebitmaprenderingcontext-alpha}
flag set to false, the contents of the
[`canvas`{#the-imagebitmaprenderingcontext-interface:the-canvas-element-6}](#the-canvas-element)
element to which the context is bound are obtained by compositing the
context\'s [output
bitmap](#concept-imagebitmaprenderingcontext-output-bitmap){#the-imagebitmaprenderingcontext-interface:concept-imagebitmaprenderingcontext-output-bitmap-3}
onto an [opaque
black](https://drafts.csswg.org/css-color/#opaque-black){#the-imagebitmaprenderingcontext-interface:opaque-black
x-internal="opaque-black"} bitmap of the same size using the
[source-over](https://drafts.fxtf.org/compositing/#porterduffcompositingoperators_srcover){#the-imagebitmaprenderingcontext-interface:gcop-source-over
x-internal="gcop-source-over"} compositing operator. If the
[alpha](#concept-imagebitmaprenderingcontext-alpha){#the-imagebitmaprenderingcontext-interface:concept-imagebitmaprenderingcontext-alpha-2}
flag is set to true, then the [output
bitmap](#concept-imagebitmaprenderingcontext-output-bitmap){#the-imagebitmaprenderingcontext-interface:concept-imagebitmaprenderingcontext-output-bitmap-4}
is used as the contents of the
[`canvas`{#the-imagebitmaprenderingcontext-interface:the-canvas-element-7}](#the-canvas-element)
element to which the context is bound.
[\[COMPOSITE\]](references.html#refsCOMPOSITE)

The step of compositing over an [opaque
black](https://drafts.csswg.org/css-color/#opaque-black){#the-imagebitmaprenderingcontext-interface:opaque-black-2
x-internal="opaque-black"} bitmap ought to be elided whenever equivalent
results can be obtained more efficiently by other means.

------------------------------------------------------------------------

When a user agent is required to [set an
`ImageBitmapRenderingContext`\'s output
bitmap]{#set-an-imagebitmaprenderingcontext's-output-bitmap .dfn}, with
a `context`{.variable} argument that is an
[`ImageBitmapRenderingContext`{#the-imagebitmaprenderingcontext-interface:imagebitmaprenderingcontext-7}](#imagebitmaprenderingcontext)
object and an optional argument `bitmap`{.variable} that refers to
[bitmap
data](imagebitmap-and-animations.html#concept-imagebitmap-bitmap-data){#the-imagebitmaprenderingcontext-interface:concept-imagebitmap-bitmap-data-4},
it must run these steps:

1.  If a `bitmap`{.variable} argument was not provided, then:

    1.  Set `context`{.variable}\'s [bitmap
        mode](#concept-imagebitmaprenderingcontext-bitmap-mode){#the-imagebitmaprenderingcontext-interface:concept-imagebitmaprenderingcontext-bitmap-mode}
        to
        [blank](#concept-imagebitmaprenderingcontext-blank){#the-imagebitmaprenderingcontext-interface:concept-imagebitmaprenderingcontext-blank-2}.

    2.  Let `canvas`{.variable} be the
        [`canvas`{#the-imagebitmaprenderingcontext-interface:the-canvas-element-8}](#the-canvas-element)
        element to which `context`{.variable} is bound.

    3.  Set `context`{.variable}\'s [output
        bitmap](#concept-imagebitmaprenderingcontext-output-bitmap){#the-imagebitmaprenderingcontext-interface:concept-imagebitmaprenderingcontext-output-bitmap-5}
        to be [transparent
        black](https://drafts.csswg.org/css-color/#transparent-black){#the-imagebitmaprenderingcontext-interface:transparent-black-2
        x-internal="transparent-black"} with a [natural
        width](https://drafts.csswg.org/css-images/#natural-width){#the-imagebitmaprenderingcontext-interface:natural-width
        x-internal="natural-width"} equal to [the numeric
        value](#obtain-numeric-values){#the-imagebitmaprenderingcontext-interface:obtain-numeric-values}
        of `canvas`{.variable}\'s
        [`width`{#the-imagebitmaprenderingcontext-interface:attr-canvas-width-2}](#attr-canvas-width)
        attribute and a [natural
        height](https://drafts.csswg.org/css-images/#natural-height){#the-imagebitmaprenderingcontext-interface:natural-height
        x-internal="natural-height"} equal to [the numeric
        value](#obtain-numeric-values){#the-imagebitmaprenderingcontext-interface:obtain-numeric-values-2}
        of `canvas`{.variable}\'s
        [`height`{#the-imagebitmaprenderingcontext-interface:attr-canvas-height-2}](#attr-canvas-height)
        attribute, those values being interpreted in [CSS
        pixels](https://drafts.csswg.org/css-values/#px){#the-imagebitmaprenderingcontext-interface:'px'
        x-internal="'px'"}.

    4.  Set the [output
        bitmap](#concept-imagebitmaprenderingcontext-output-bitmap){#the-imagebitmaprenderingcontext-interface:concept-imagebitmaprenderingcontext-output-bitmap-6}\'s
        [origin-clean](#concept-canvas-origin-clean){#the-imagebitmaprenderingcontext-interface:concept-canvas-origin-clean}
        flag to true.

2.  If a `bitmap`{.variable} argument was provided, then:

    1.  Set `context`{.variable}\'s [bitmap
        mode](#concept-imagebitmaprenderingcontext-bitmap-mode){#the-imagebitmaprenderingcontext-interface:concept-imagebitmaprenderingcontext-bitmap-mode-2}
        to
        [valid](#concept-imagebitmaprenderingcontext-valid){#the-imagebitmaprenderingcontext-interface:concept-imagebitmaprenderingcontext-valid-2}.

    2.  Set `context`{.variable}\'s [output
        bitmap](#concept-imagebitmaprenderingcontext-output-bitmap){#the-imagebitmaprenderingcontext-interface:concept-imagebitmaprenderingcontext-output-bitmap-7}
        to refer to the same underlying bitmap data as
        `bitmap`{.variable}, without making a copy.

        The
        [origin-clean](#concept-canvas-origin-clean){#the-imagebitmaprenderingcontext-interface:concept-canvas-origin-clean-2}
        flag of `bitmap`{.variable} is included in the bitmap data to be
        referenced by `context`{.variable}\'s [output
        bitmap](#concept-imagebitmaprenderingcontext-output-bitmap){#the-imagebitmaprenderingcontext-interface:concept-imagebitmaprenderingcontext-output-bitmap-8}.

------------------------------------------------------------------------

The [`ImageBitmapRenderingContext` creation
algorithm]{#imagebitmaprenderingcontext-creation-algorithm .dfn}, which
is passed a `target`{.variable} and `options`{.variable}, consists of
running these steps:

1.  Let `settings`{.variable} be the result of
    [converting](https://webidl.spec.whatwg.org/#es-type-mapping){#the-imagebitmaprenderingcontext-interface:concept-idl-convert
    x-internal="concept-idl-convert"} `options`{.variable} to the
    dictionary type
    [`ImageBitmapRenderingContextSettings`{#the-imagebitmaprenderingcontext-interface:imagebitmaprenderingcontextsettings}](#imagebitmaprenderingcontextsettings).
    (This can throw an exception.)

2.  Let `context`{.variable} be a new
    [`ImageBitmapRenderingContext`{#the-imagebitmaprenderingcontext-interface:imagebitmaprenderingcontext-8}](#imagebitmaprenderingcontext)
    object.

3.  Initialize `context`{.variable}\'s
    [`canvas`{#the-imagebitmaprenderingcontext-interface:dom-context-2d-canvas}](#dom-context-2d-canvas)
    attribute to point to `target`{.variable}.

4.  Set `context`{.variable}\'s [output
    bitmap](#concept-imagebitmaprenderingcontext-output-bitmap){#the-imagebitmaprenderingcontext-interface:concept-imagebitmaprenderingcontext-output-bitmap-9}
    to the same bitmap as `target`{.variable}\'s bitmap (so that they
    are shared).

5.  Run the steps to [set an `ImageBitmapRenderingContext`\'s output
    bitmap](#set-an-imagebitmaprenderingcontext's-output-bitmap){#the-imagebitmaprenderingcontext-interface:set-an-imagebitmaprenderingcontext's-output-bitmap}
    with `context`{.variable}.

6.  Initialize `context`{.variable}\'s
    [alpha](#concept-imagebitmaprenderingcontext-alpha){#the-imagebitmaprenderingcontext-interface:concept-imagebitmaprenderingcontext-alpha-3}
    flag to true.

7.  Process each of the members of `settings`{.variable} as follows:

    [`alpha`]{#dom-imagebitmaprenderingcontextsettings-alpha .dfn dfn-for="ImageBitmapRenderingContextSettings" dfn-type="dict-member"}
    :   If false, then set `context`{.variable}\'s
        [alpha](#concept-imagebitmaprenderingcontext-alpha){#the-imagebitmaprenderingcontext-interface:concept-imagebitmaprenderingcontext-alpha-4}
        flag to false.

8.  Return `context`{.variable}.

------------------------------------------------------------------------

The
[`transferFromImageBitmap(``bitmap`{.variable}`)`]{#dom-imagebitmaprenderingcontext-transferfromimagebitmap
.dfn dfn-for="ImageBitmapRenderingContext" dfn-type="method"} method,
when invoked, must run these steps:

1.  Let `bitmapContext`{.variable} be the
    [`ImageBitmapRenderingContext`{#the-imagebitmaprenderingcontext-interface:imagebitmaprenderingcontext-9}](#imagebitmaprenderingcontext)
    object on which the
    [`transferFromImageBitmap()`{#the-imagebitmaprenderingcontext-interface:dom-imagebitmaprenderingcontext-transferfromimagebitmap-4}](#dom-imagebitmaprenderingcontext-transferfromimagebitmap)
    method was called.

2.  If `bitmap`{.variable} is null, then run the steps to [set an
    ImageBitmapRenderingContext\'s output
    bitmap](#set-an-imagebitmaprenderingcontext's-output-bitmap){#the-imagebitmaprenderingcontext-interface:set-an-imagebitmaprenderingcontext's-output-bitmap-2},
    with `bitmapContext`{.variable} as the `context`{.variable} argument
    and no `bitmap`{.variable} argument, then return.

3.  If the value of `bitmap`{.variable}\'s
    [\[\[Detached\]\]](structured-data.html#detached){#the-imagebitmaprenderingcontext-interface:detached}
    internal slot is set to true, then throw an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-imagebitmaprenderingcontext-interface:invalidstateerror
    x-internal="invalidstateerror"}
    [`DOMException`{#the-imagebitmaprenderingcontext-interface:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

4.  Run the steps to [set an `ImageBitmapRenderingContext`\'s output
    bitmap](#set-an-imagebitmaprenderingcontext's-output-bitmap){#the-imagebitmaprenderingcontext-interface:set-an-imagebitmaprenderingcontext's-output-bitmap-3},
    with the `context`{.variable} argument equal to
    `bitmapContext`{.variable}, and the `bitmap`{.variable} argument
    referring to `bitmap`{.variable}\'s underlying [bitmap
    data](imagebitmap-and-animations.html#concept-imagebitmap-bitmap-data){#the-imagebitmaprenderingcontext-interface:concept-imagebitmap-bitmap-data-5}.

5.  Set the value of `bitmap`{.variable}\'s
    [\[\[Detached\]\]](structured-data.html#detached){#the-imagebitmaprenderingcontext-interface:detached-2}
    internal slot to true.

6.  Unset `bitmap`{.variable}\'s [bitmap
    data](imagebitmap-and-animations.html#concept-imagebitmap-bitmap-data){#the-imagebitmaprenderingcontext-interface:concept-imagebitmap-bitmap-data-6}.

##### [4.12.5.3]{.secno} The [`OffscreenCanvas`{#the-offscreencanvas-interface:offscreencanvas}](#offscreencanvas) interface[](#the-offscreencanvas-interface){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[OffscreenCanvas](https://developer.mozilla.org/en-US/docs/Web/API/OffscreenCanvas "When using the <canvas> element or the Canvas API, rendering, animation, and user interaction usually happen on the main execution thread of a web application. The computation relating to canvas animations and rendering can have a significant impact on application performance.")

Support in all current engines.

::: support
[Firefox105+]{.firefox .yes}[Safari16.4+]{.safari
.yes}[Chrome69+]{.chrome .yes}

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

``` idl
typedef (OffscreenCanvasRenderingContext2D or ImageBitmapRenderingContext or WebGLRenderingContext or WebGL2RenderingContext or GPUCanvasContext) OffscreenRenderingContext;

dictionary ImageEncodeOptions {
  DOMString type = "image/png";
  unrestricted double quality;
};

enum OffscreenRenderingContextId { "2d", "bitmaprenderer", "webgl", "webgl2", "webgpu" };

[Exposed=(Window,Worker), Transferable]
interface OffscreenCanvas : EventTarget {
  constructor([EnforceRange] unsigned long long width, [EnforceRange] unsigned long long height);

  attribute [EnforceRange] unsigned long long width;
  attribute [EnforceRange] unsigned long long height;

  OffscreenRenderingContext? getContext(OffscreenRenderingContextId contextId, optional any options = null);
  ImageBitmap transferToImageBitmap();
  Promise<Blob> convertToBlob(optional ImageEncodeOptions options = {});

  attribute EventHandler oncontextlost;
  attribute EventHandler oncontextrestored;
};
```

[`OffscreenCanvas`{#the-offscreencanvas-interface:offscreencanvas-2}](#offscreencanvas)
is an
[`EventTarget`{#the-offscreencanvas-interface:eventtarget-2}](https://dom.spec.whatwg.org/#interface-eventtarget){x-internal="eventtarget"},
so both
[`OffscreenCanvasRenderingContext2D`{#the-offscreencanvas-interface:offscreencanvasrenderingcontext2d-2}](#offscreencanvasrenderingcontext2d)
and WebGL can fire events at it.
[`OffscreenCanvasRenderingContext2D`{#the-offscreencanvas-interface:offscreencanvasrenderingcontext2d-3}](#offscreencanvasrenderingcontext2d)
can fire
[`contextlost`{#the-offscreencanvas-interface:event-contextlost}](indices.html#event-contextlost)
and
[`contextrestored`{#the-offscreencanvas-interface:event-contextrestored}](indices.html#event-contextrestored),
and WebGL can fire `webglcontextlost` and `webglcontextrestored`.
[\[WEBGL\]](references.html#refsWEBGL)

[`OffscreenCanvas`{#the-offscreencanvas-interface:offscreencanvas-3}](#offscreencanvas)
objects are used to create rendering contexts, much like an
[`HTMLCanvasElement`{#the-offscreencanvas-interface:htmlcanvaselement}](#htmlcanvaselement),
but with no connection to the DOM. This makes it possible to use canvas
rendering contexts in [workers](workers.html#workers).

An
[`OffscreenCanvas`{#the-offscreencanvas-interface:offscreencanvas-4}](#offscreencanvas)
object may hold a weak reference to a [placeholder `canvas`
element]{#offscreencanvas-placeholder .dfn}, which is typically in the
DOM, whose embedded content is provided by the
[`OffscreenCanvas`{#the-offscreencanvas-interface:offscreencanvas-5}](#offscreencanvas)
object. The bitmap of the
[`OffscreenCanvas`{#the-offscreencanvas-interface:offscreencanvas-6}](#offscreencanvas)
object is pushed to the [placeholder `canvas`
element](#offscreencanvas-placeholder){#the-offscreencanvas-interface:offscreencanvas-placeholder}
as part of the
[`OffscreenCanvas`{#the-offscreencanvas-interface:offscreencanvas-7}](#offscreencanvas)\'s
[relevant
agent](webappapis.html#relevant-agent){#the-offscreencanvas-interface:relevant-agent}\'s
[event
loop](webappapis.html#concept-agent-event-loop){#the-offscreencanvas-interface:concept-agent-event-loop}\'s
[update the
rendering](webappapis.html#update-the-rendering){#the-offscreencanvas-interface:update-the-rendering}
steps.

`offscreenCanvas`{.variable}` = new `[`OffscreenCanvas`](#dom-offscreencanvas){#dom-offscreencanvas-dev}`(``width`{.variable}`, ``height`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[OffscreenCanvas/OffscreenCanvas](https://developer.mozilla.org/en-US/docs/Web/API/OffscreenCanvas/OffscreenCanvas "The OffscreenCanvas() constructor returns a newly instantiated OffscreenCanvas object.")

Support in all current engines.

::: support
[Firefox105+]{.firefox .yes}[Safari16.4+]{.safari
.yes}[Chrome69+]{.chrome .yes}

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

Returns a new
[`OffscreenCanvas`{#the-offscreencanvas-interface:offscreencanvas-8}](#offscreencanvas)
object that is not linked to a [placeholder `canvas`
element](#offscreencanvas-placeholder){#the-offscreencanvas-interface:offscreencanvas-placeholder-2},
and whose bitmap\'s size is determined by the `width`{.variable} and
`height`{.variable} arguments.

`context`{.variable}` = ``offscreenCanvas`{.variable}`.`[`getContext`](#dom-offscreencanvas-getcontext){#dom-offscreencanvas-getcontext-dev}`(``contextId`{.variable}` [, ``options`{.variable}` ])`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[OffscreenCanvas/getContext](https://developer.mozilla.org/en-US/docs/Web/API/OffscreenCanvas/getContext "The OffscreenCanvas.getContext() method returns a drawing context for an offscreen canvas, or null if the context identifier is not supported.")

Support in all current engines.

::: support
[Firefox105+]{.firefox .yes}[Safari16.4+]{.safari
.yes}[Chrome69+]{.chrome .yes}

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

Returns an object that exposes an API for drawing on the
[`OffscreenCanvas`{#the-offscreencanvas-interface:offscreencanvas-9}](#offscreencanvas)
object. `contextId`{.variable} specifies the desired API:
\"[`2d`{#the-offscreencanvas-interface:offscreen-context-type-2d-2}](#offscreen-context-type-2d)\",
\"[`bitmaprenderer`{#the-offscreencanvas-interface:offscreen-context-type-bitmaprenderer-2}](#offscreen-context-type-bitmaprenderer)\",
\"[`webgl`{#the-offscreencanvas-interface:offscreen-context-type-webgl-2}](#offscreen-context-type-webgl)\",
\"[`webgl2`{#the-offscreencanvas-interface:offscreen-context-type-webgl2-2}](#offscreen-context-type-webgl2)\",
or
\"[`webgpu`{#the-offscreencanvas-interface:offscreen-context-type-webgpu-2}](#offscreen-context-type-webgpu)\".
`options`{.variable} is handled by that API.

This specification defines the
\"[`2d`{#the-offscreencanvas-interface:canvas-context-2d}](#canvas-context-2d)\"
context below, which is similar but distinct from the
\"[`2d`{#the-offscreencanvas-interface:offscreen-context-type-2d-3}](#offscreen-context-type-2d)\"
context that is created from a
[`canvas`{#the-offscreencanvas-interface:the-canvas-element}](#the-canvas-element)
element. The WebGL specifications define the
\"[`webgl`{#the-offscreencanvas-interface:offscreen-context-type-webgl-3}](#offscreen-context-type-webgl)\"
and
\"[`webgl2`{#the-offscreencanvas-interface:offscreen-context-type-webgl2-3}](#offscreen-context-type-webgl2)\"
contexts. WebGPU defines the
\"[`webgpu`{#the-offscreencanvas-interface:offscreen-context-type-webgpu-3}](#offscreen-context-type-webgpu)\"
context. [\[WEBGL\]](references.html#refsWEBGL)
[\[WEBGPU\]](references.html#refsWEBGPU)

Returns null if the canvas has already been initialized with another
context type (e.g., trying to get a
\"[`2d`{#the-offscreencanvas-interface:offscreen-context-type-2d-4}](#offscreen-context-type-2d)\"
context after getting a
\"[`webgl`{#the-offscreencanvas-interface:offscreen-context-type-webgl-4}](#offscreen-context-type-webgl)\"
context).

An
[`OffscreenCanvas`{#the-offscreencanvas-interface:offscreencanvas-10}](#offscreencanvas)
object has an internal [bitmap]{#offscreencanvas-bitmap .dfn} that is
initialized when the object is created. The width and height of the
[bitmap](#offscreencanvas-bitmap){#the-offscreencanvas-interface:offscreencanvas-bitmap}
are equal to the values of the
[`width`{#the-offscreencanvas-interface:dom-offscreencanvas-width-2}](#dom-offscreencanvas-width)
and
[`height`{#the-offscreencanvas-interface:dom-offscreencanvas-height-2}](#dom-offscreencanvas-height)
attributes of the
[`OffscreenCanvas`{#the-offscreencanvas-interface:offscreencanvas-11}](#offscreencanvas)
object. Initially, all the bitmap\'s pixels are [transparent
black](https://drafts.csswg.org/css-color/#transparent-black){#the-offscreencanvas-interface:transparent-black
x-internal="transparent-black"}.

An
[`OffscreenCanvas`{#the-offscreencanvas-interface:offscreencanvas-12}](#offscreencanvas)
object has an internal [inherited
language]{#offscreencanvas-inherited-lang .dfn} and [inherited
direction]{#offscreencanvas-inherited-direction .dfn} set when the
[`OffscreenCanvas`{#the-offscreencanvas-interface:offscreencanvas-13}](#offscreencanvas)
is created.

An
[`OffscreenCanvas`{#the-offscreencanvas-interface:offscreencanvas-14}](#offscreencanvas)
object can have a rendering context bound to it. Initially, it does not
have a bound rendering context. To keep track of whether it has a
rendering context or not, and what kind of rendering context it is, an
[`OffscreenCanvas`{#the-offscreencanvas-interface:offscreencanvas-15}](#offscreencanvas)
object also has a [context mode]{#offscreencanvas-context-mode .dfn},
which is initially [none]{#offscreencanvas-context-none .dfn} but can be
changed to either [2d]{#offscreencanvas-context-2d .dfn},
[bitmaprenderer]{#offscreencanvas-context-bitmaprenderer .dfn},
[webgl]{#offscreencanvas-context-webgl .dfn},
[webgl2]{#offscreencanvas-context-webgl2 .dfn},
[webgpu]{#offscreencanvas-context-webgpu .dfn}, or
[detached]{#offscreencanvas-context-detached .dfn} by algorithms defined
in this specification.

The
[`new OffscreenCanvas(``width`{.variable}`, ``height`{.variable}`)`]{#dom-offscreencanvas
.dfn dfn-for="OffscreenCanvas" dfn-type="constructor"} constructor steps
are:

1.  Initialize the
    [bitmap](#offscreencanvas-bitmap){#the-offscreencanvas-interface:offscreencanvas-bitmap-2}
    of
    [this](https://webidl.spec.whatwg.org/#this){#the-offscreencanvas-interface:this
    x-internal="this"} to a rectangular array of [transparent
    black](https://drafts.csswg.org/css-color/#transparent-black){#the-offscreencanvas-interface:transparent-black-2
    x-internal="transparent-black"} pixels of the dimensions specified
    by `width`{.variable} and `height`{.variable}.

2.  Initialize the
    [`width`{#the-offscreencanvas-interface:dom-offscreencanvas-width-3}](#dom-offscreencanvas-width)
    of
    [this](https://webidl.spec.whatwg.org/#this){#the-offscreencanvas-interface:this-2
    x-internal="this"} to `width`{.variable}.

3.  Initialize the
    [`height`{#the-offscreencanvas-interface:dom-offscreencanvas-height-3}](#dom-offscreencanvas-height)
    of
    [this](https://webidl.spec.whatwg.org/#this){#the-offscreencanvas-interface:this-3
    x-internal="this"} to `height`{.variable}.

4.  Set
    [this](https://webidl.spec.whatwg.org/#this){#the-offscreencanvas-interface:this-4
    x-internal="this"}\'s [inherited
    language](#offscreencanvas-inherited-lang){#the-offscreencanvas-interface:offscreencanvas-inherited-lang}
    to [explicitly
    unknown](dom.html#concept-explicitly-unknown){#the-offscreencanvas-interface:concept-explicitly-unknown}.

5.  Set
    [this](https://webidl.spec.whatwg.org/#this){#the-offscreencanvas-interface:this-5
    x-internal="this"}\'s [inherited
    direction](#offscreencanvas-inherited-direction){#the-offscreencanvas-interface:offscreencanvas-inherited-direction}
    to \"`ltr`\".

6.  Let `global`{.variable} be the [relevant global
    object](webappapis.html#concept-relevant-global){#the-offscreencanvas-interface:concept-relevant-global}
    of
    [this](https://webidl.spec.whatwg.org/#this){#the-offscreencanvas-interface:this-6
    x-internal="this"}.

7.  If `global`{.variable} is a
    [`Window`{#the-offscreencanvas-interface:window}](nav-history-apis.html#window)
    object:

    1.  Let `element`{.variable} be the [document
        element](https://dom.spec.whatwg.org/#document-element){#the-offscreencanvas-interface:document-element
        x-internal="document-element"} of `global`{.variable}\'s
        [associated
        `Document`](nav-history-apis.html#concept-document-window){#the-offscreencanvas-interface:concept-document-window}.

    2.  If `element`{.variable} is not null:

        1.  Set the [inherited
            language](#offscreencanvas-inherited-lang){#the-offscreencanvas-interface:offscreencanvas-inherited-lang-2}
            of
            [this](https://webidl.spec.whatwg.org/#this){#the-offscreencanvas-interface:this-7
            x-internal="this"} to `element`{.variable}\'s
            [language](dom.html#language){#the-offscreencanvas-interface:language}.

        2.  Set the [inherited
            direction](#offscreencanvas-inherited-direction){#the-offscreencanvas-interface:offscreencanvas-inherited-direction-2}
            of
            [this](https://webidl.spec.whatwg.org/#this){#the-offscreencanvas-interface:this-8
            x-internal="this"} to `element`{.variable}\'s
            [directionality](dom.html#the-directionality){#the-offscreencanvas-interface:the-directionality}.

[`OffscreenCanvas`{#the-offscreencanvas-interface:offscreencanvas-16}](#offscreencanvas)
objects are
[transferable](structured-data.html#transferable-objects){#the-offscreencanvas-interface:transferable-objects}.
Their [transfer
steps](structured-data.html#transfer-steps){#the-offscreencanvas-interface:transfer-steps},
given `value`{.variable} and `dataHolder`{.variable}, are as follows:

1.  If `value`{.variable}\'s [context
    mode](#offscreencanvas-context-mode){#the-offscreencanvas-interface:offscreencanvas-context-mode}
    is not equal to
    [none](#offscreencanvas-context-none){#the-offscreencanvas-interface:offscreencanvas-context-none},
    then throw an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-offscreencanvas-interface:invalidstateerror
    x-internal="invalidstateerror"}
    [`DOMException`{#the-offscreencanvas-interface:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  Set `value`{.variable}\'s [context
    mode](#offscreencanvas-context-mode){#the-offscreencanvas-interface:offscreencanvas-context-mode-2}
    to
    [detached](#offscreencanvas-context-detached){#the-offscreencanvas-interface:offscreencanvas-context-detached}.

3.  Let `width`{.variable} and `height`{.variable} be the dimensions of
    `value`{.variable}\'s
    [bitmap](#offscreencanvas-bitmap){#the-offscreencanvas-interface:offscreencanvas-bitmap-3}.

4.  Let `language`{.variable} and `direction`{.variable} be the values
    of `value`{.variable}\'s [inherited
    language](#offscreencanvas-inherited-lang){#the-offscreencanvas-interface:offscreencanvas-inherited-lang-3}
    and [inherited
    direction](#offscreencanvas-inherited-direction){#the-offscreencanvas-interface:offscreencanvas-inherited-direction-3}.

5.  Unset `value`{.variable}\'s
    [bitmap](#offscreencanvas-bitmap){#the-offscreencanvas-interface:offscreencanvas-bitmap-4}.

6.  Set `dataHolder`{.variable}.\[\[Width\]\] to `width`{.variable} and
    `dataHolder`{.variable}.\[\[Height\]\] to `height`{.variable}.

7.  Set `dataHolder`{.variable}.\[\[Language\]\] to
    `language`{.variable} and `dataHolder`{.variable}.\[\[Direction\]\]
    to `direction`{.variable}.

8.  Set `dataHolder`{.variable}.\[\[PlaceholderCanvas\]\] to be a weak
    reference to `value`{.variable}\'s [placeholder `canvas`
    element](#offscreencanvas-placeholder){#the-offscreencanvas-interface:offscreencanvas-placeholder-3},
    if `value`{.variable} has one, or null if it does not.

Their [transfer-receiving
steps](structured-data.html#transfer-receiving-steps){#the-offscreencanvas-interface:transfer-receiving-steps},
given `dataHolder`{.variable} and `value`{.variable}, are:

1.  Initialize `value`{.variable}\'s
    [bitmap](#offscreencanvas-bitmap){#the-offscreencanvas-interface:offscreencanvas-bitmap-5}
    to a rectangular array of [transparent
    black](https://drafts.csswg.org/css-color/#transparent-black){#the-offscreencanvas-interface:transparent-black-3
    x-internal="transparent-black"} pixels with width given by
    `dataHolder`{.variable}.\[\[Width\]\] and height given by
    `dataHolder`{.variable}.\[\[Height\]\].

2.  Set `value`{.variable}\'s [inherited
    language](#offscreencanvas-inherited-lang){#the-offscreencanvas-interface:offscreencanvas-inherited-lang-4}
    to `dataHolder`{.variable}.\[\[Language\]\] and its [inherited
    direction](#offscreencanvas-inherited-direction){#the-offscreencanvas-interface:offscreencanvas-inherited-direction-4}
    to `dataHolder`{.variable}.\[\[Direction\]\].

3.  If `dataHolder`{.variable}.\[\[PlaceholderCanvas\]\] is not null,
    set `value`{.variable}\'s [placeholder `canvas`
    element](#offscreencanvas-placeholder){#the-offscreencanvas-interface:offscreencanvas-placeholder-4}
    to `dataHolder`{.variable}.\[\[PlaceholderCanvas\]\] (while
    maintaining the weak reference semantics).

------------------------------------------------------------------------

The
[`getContext(``contextId`{.variable}`, ``options`{.variable}`)`]{#dom-offscreencanvas-getcontext
.dfn dfn-for="OffscreenCanvas" dfn-type="method"} method of an
[`OffscreenCanvas`{#the-offscreencanvas-interface:offscreencanvas-17}](#offscreencanvas)
object, when invoked, must run these steps:

1.  If `options`{.variable} is not an
    [object](https://webidl.spec.whatwg.org/#idl-object){#the-offscreencanvas-interface:idl-object
    x-internal="idl-object"}, then set `options`{.variable} to null.

2.  Set `options`{.variable} to the result of
    [converting](https://webidl.spec.whatwg.org/#es-type-mapping){#the-offscreencanvas-interface:concept-idl-convert
    x-internal="concept-idl-convert"} `options`{.variable} to a
    JavaScript value.

3.  Run the steps in the cell of the following table whose column header
    matches this
    [`OffscreenCanvas`{#the-offscreencanvas-interface:offscreencanvas-18}](#offscreencanvas)
    object\'s [context
    mode](#offscreencanvas-context-mode){#the-offscreencanvas-interface:offscreencanvas-context-mode-3}
    and whose row header matches `contextId`{.variable}:

    [none](#offscreencanvas-context-none){#the-offscreencanvas-interface:offscreencanvas-context-none-2}

    [2d](#offscreencanvas-context-2d){#the-offscreencanvas-interface:offscreencanvas-context-2d}

    [bitmaprenderer](#offscreencanvas-context-bitmaprenderer){#the-offscreencanvas-interface:offscreencanvas-context-bitmaprenderer}

    [webgl](#offscreencanvas-context-webgl){#the-offscreencanvas-interface:offscreencanvas-context-webgl}
    or
    [webgl2](#offscreencanvas-context-webgl2){#the-offscreencanvas-interface:offscreencanvas-context-webgl2}

    [webgpu](#offscreencanvas-context-webgpu){#the-offscreencanvas-interface:offscreencanvas-context-webgpu}

    [detached](#offscreencanvas-context-detached){#the-offscreencanvas-interface:offscreencanvas-context-detached-2}

    \"[`2d`]{#offscreen-context-type-2d .dfn}\"

    1.  Let `context`{.variable} be the result of running the [offscreen
        2D context creation
        algorithm](#offscreen-2d-context-creation-algorithm){#the-offscreencanvas-interface:offscreen-2d-context-creation-algorithm}
        given
        [this](https://webidl.spec.whatwg.org/#this){#the-offscreencanvas-interface:this-9
        x-internal="this"} and `options`{.variable}.

    2.  Set
        [this](https://webidl.spec.whatwg.org/#this){#the-offscreencanvas-interface:this-10
        x-internal="this"}\'s [context
        mode](#offscreencanvas-context-mode){#the-offscreencanvas-interface:offscreencanvas-context-mode-4}
        to
        [2d](#offscreencanvas-context-2d){#the-offscreencanvas-interface:offscreencanvas-context-2d-2}.

    3.  Return `context`{.variable}.

    Return the same object as was returned the last time the method was
    invoked with this same first argument.

    Return null.

    Return null.

    Return null.

    Throw an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-offscreencanvas-interface:invalidstateerror-2
    x-internal="invalidstateerror"}
    [`DOMException`{#the-offscreencanvas-interface:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    \"[`bitmaprenderer`]{#offscreen-context-type-bitmaprenderer .dfn}\"

    1.  Let `context`{.variable} be the result of running the
        [`ImageBitmapRenderingContext` creation
        algorithm](#imagebitmaprenderingcontext-creation-algorithm){#the-offscreencanvas-interface:imagebitmaprenderingcontext-creation-algorithm}
        given
        [this](https://webidl.spec.whatwg.org/#this){#the-offscreencanvas-interface:this-11
        x-internal="this"} and `options`{.variable}.

    2.  Set
        [this](https://webidl.spec.whatwg.org/#this){#the-offscreencanvas-interface:this-12
        x-internal="this"}\'s [context
        mode](#offscreencanvas-context-mode){#the-offscreencanvas-interface:offscreencanvas-context-mode-5}
        to
        [bitmaprenderer](#offscreencanvas-context-bitmaprenderer){#the-offscreencanvas-interface:offscreencanvas-context-bitmaprenderer-2}.

    3.  Return `context`{.variable}.

    Return null.

    Return the same object as was returned the last time the method was
    invoked with this same first argument.

    Return null.

    Return null.

    Throw an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-offscreencanvas-interface:invalidstateerror-3
    x-internal="invalidstateerror"}
    [`DOMException`{#the-offscreencanvas-interface:domexception-3}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    \"[`webgl`]{#offscreen-context-type-webgl .dfn}\" or
    \"[`webgl2`]{#offscreen-context-type-webgl2 .dfn}\"

    1.  Let `context`{.variable} be the result of following the
        instructions given in the WebGL specifications\' *Context
        Creation* sections. [\[WEBGL\]](references.html#refsWEBGL)

    2.  If `context`{.variable} is null, then return null; otherwise set
        [this](https://webidl.spec.whatwg.org/#this){#the-offscreencanvas-interface:this-13
        x-internal="this"}\'s [context
        mode](#offscreencanvas-context-mode){#the-offscreencanvas-interface:offscreencanvas-context-mode-6}
        to
        [webgl](#offscreencanvas-context-webgl){#the-offscreencanvas-interface:offscreencanvas-context-webgl-2}
        or
        [webgl2](#offscreencanvas-context-webgl2){#the-offscreencanvas-interface:offscreencanvas-context-webgl2-2}.

    3.  Return `context`{.variable}.

    Return null.

    Return null.

    Return the same value as was returned the last time the method was
    invoked with this same first argument.

    Return null.

    Throw an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-offscreencanvas-interface:invalidstateerror-4
    x-internal="invalidstateerror"}
    [`DOMException`{#the-offscreencanvas-interface:domexception-4}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    \"[`webgpu`]{#offscreen-context-type-webgpu .dfn}\"

    1.  Let `context`{.variable} be the result of following the
        instructions given in WebGPU\'s [Canvas
        Rendering](https://gpuweb.github.io/gpuweb/#canvas-rendering)
        section. [\[WEBGPU\]](references.html#refsWEBGPU)

    2.  If `context`{.variable} is null, then return null; otherwise set
        [this](https://webidl.spec.whatwg.org/#this){#the-offscreencanvas-interface:this-14
        x-internal="this"}\'s [context
        mode](#offscreencanvas-context-mode){#the-offscreencanvas-interface:offscreencanvas-context-mode-7}
        to
        [webgpu](#offscreencanvas-context-webgpu){#the-offscreencanvas-interface:offscreencanvas-context-webgpu-2}.

    3.  Return `context`{.variable}.

    Return null.

    Return null.

    Return null.

    Return the same value as was returned the last time the method was
    invoked with this same first argument.

    Throw an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-offscreencanvas-interface:invalidstateerror-5
    x-internal="invalidstateerror"}
    [`DOMException`{#the-offscreencanvas-interface:domexception-5}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

------------------------------------------------------------------------

`offscreenCanvas`{.variable}`.`[`width`](#dom-offscreencanvas-width){#dom-offscreencanvas-width-dev}` [ = ``value`{.variable}` ]`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[OffscreenCanvas/width](https://developer.mozilla.org/en-US/docs/Web/API/OffscreenCanvas/width "The width property returns and sets the width of an OffscreenCanvas object.")

Support in all current engines.

::: support
[Firefox105+]{.firefox .yes}[Safari16.4+]{.safari
.yes}[Chrome69+]{.chrome .yes}

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

`offscreenCanvas`{.variable}`.`[`height`](#dom-offscreencanvas-height){#dom-offscreencanvas-height-dev}` [ = ``value`{.variable}` ]`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[OffscreenCanvas/height](https://developer.mozilla.org/en-US/docs/Web/API/OffscreenCanvas/height "The height property returns and sets the height of an OffscreenCanvas object.")

Support in all current engines.

::: support
[Firefox105+]{.firefox .yes}[Safari16.4+]{.safari
.yes}[Chrome69+]{.chrome .yes}

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

These attributes return the dimensions of the
[`OffscreenCanvas`{#the-offscreencanvas-interface:offscreencanvas-19}](#offscreencanvas)
object\'s
[bitmap](#offscreencanvas-bitmap){#the-offscreencanvas-interface:offscreencanvas-bitmap-6}.

They can be set, to replace the
[bitmap](#offscreencanvas-bitmap){#the-offscreencanvas-interface:offscreencanvas-bitmap-7}
with a new, [transparent
black](https://drafts.csswg.org/css-color/#transparent-black){#the-offscreencanvas-interface:transparent-black-4
x-internal="transparent-black"} bitmap of the specified dimensions
(effectively resizing it).

If either the [`width`]{#dom-offscreencanvas-width .dfn
dfn-for="OffscreenCanvas" dfn-type="attribute"} or
[`height`]{#dom-offscreencanvas-height .dfn dfn-for="OffscreenCanvas"
dfn-type="attribute"} attributes of an
[`OffscreenCanvas`{#the-offscreencanvas-interface:offscreencanvas-20}](#offscreencanvas)
object are set (to a new value or to the same value as before) and the
[`OffscreenCanvas`{#the-offscreencanvas-interface:offscreencanvas-21}](#offscreencanvas)
object\'s [context
mode](#offscreencanvas-context-mode){#the-offscreencanvas-interface:offscreencanvas-context-mode-8}
is
[2d](#offscreencanvas-context-2d){#the-offscreencanvas-interface:offscreencanvas-context-2d-3},
then [reset the rendering context to its default
state](#reset-the-rendering-context-to-its-default-state){#the-offscreencanvas-interface:reset-the-rendering-context-to-its-default-state}
and resize the
[`OffscreenCanvas`{#the-offscreencanvas-interface:offscreencanvas-22}](#offscreencanvas)
object\'s
[bitmap](#offscreencanvas-bitmap){#the-offscreencanvas-interface:offscreencanvas-bitmap-8}
to the new values of the
[`width`{#the-offscreencanvas-interface:dom-offscreencanvas-width-4}](#dom-offscreencanvas-width)
and
[`height`{#the-offscreencanvas-interface:dom-offscreencanvas-height-4}](#dom-offscreencanvas-height)
attributes.

The resizing behavior for
\"[`webgl`{#the-offscreencanvas-interface:offscreen-context-type-webgl-5}](#offscreen-context-type-webgl)\"
and
\"[`webgl2`{#the-offscreencanvas-interface:offscreen-context-type-webgl2-4}](#offscreen-context-type-webgl2)\"
contexts is defined in the WebGL specifications.
[\[WEBGL\]](references.html#refsWEBGL)

The resizing behavior for
\"[`webgpu`{#the-offscreencanvas-interface:offscreen-context-type-webgpu-4}](#offscreen-context-type-webgpu)\"
context is defined in WebGPU. [\[WEBGPU\]](references.html#refsWEBGPU)

If an
[`OffscreenCanvas`{#the-offscreencanvas-interface:offscreencanvas-23}](#offscreencanvas)
object whose dimensions were changed has a [placeholder `canvas`
element](#offscreencanvas-placeholder){#the-offscreencanvas-interface:offscreencanvas-placeholder-5},
then the [placeholder `canvas`
element](#offscreencanvas-placeholder){#the-offscreencanvas-interface:offscreencanvas-placeholder-6}\'s
[natural
size](https://drafts.csswg.org/css-images/#natural-dimensions){#the-offscreencanvas-interface:natural-dimensions
x-internal="natural-dimensions"} will only be updated during the
[`OffscreenCanvas`{#the-offscreencanvas-interface:offscreencanvas-24}](#offscreencanvas)\'s
[relevant
agent](webappapis.html#relevant-agent){#the-offscreencanvas-interface:relevant-agent-2}\'s
[event
loop](webappapis.html#concept-agent-event-loop){#the-offscreencanvas-interface:concept-agent-event-loop-2}\'s
[update the
rendering](webappapis.html#update-the-rendering){#the-offscreencanvas-interface:update-the-rendering-2}
steps.

`promise`{.variable}` = ``offscreenCanvas`{.variable}`.`[`convertToBlob`](#dom-offscreencanvas-converttoblob){#dom-offscreencanvas-converttoblob-dev}`([``options`{.variable}`])`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[OffscreenCanvas/convertToBlob](https://developer.mozilla.org/en-US/docs/Web/API/OffscreenCanvas/convertToBlob "The OffscreenCanvas.convertToBlob() method creates a Blob object representing the image contained in the canvas.")

Support in all current engines.

::: support
[Firefox105+]{.firefox .yes}[Safari16.4+]{.safari
.yes}[Chrome69+]{.chrome .yes}

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

Returns a promise that will fulfill with a new
[`Blob`{#the-offscreencanvas-interface:blob-2}](https://w3c.github.io/FileAPI/#dfn-Blob){x-internal="blob"}
object representing a file containing the image in the
[`OffscreenCanvas`{#the-offscreencanvas-interface:offscreencanvas-25}](#offscreencanvas)
object.

The argument, if provided, is a dictionary that controls the encoding
options of the image file to be created. The
[`type`{#the-offscreencanvas-interface:image-encode-options-type-2}](#image-encode-options-type)
field specifies the file format and has a default value of
\"[`image/png`{#the-offscreencanvas-interface:image/png}](indices.html#image/png)\";
that type is also used if the requested type isn\'t supported. If the
image format supports variable quality (such as
\"[`image/jpeg`{#the-offscreencanvas-interface:image/jpeg}](indices.html#image/jpeg)\"),
then the
[`quality`{#the-offscreencanvas-interface:image-encode-options-quality-2}](#image-encode-options-quality)
field is a number in the range 0.0 to 1.0 inclusive indicating the
desired quality level for the resulting image.

`canvas`{.variable}`.`[`transferToImageBitmap`](#dom-offscreencanvas-transfertoimagebitmap){#dom-offscreencanvas-transfertoimagebitmap-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[OffscreenCanvas/transferToImageBitmap](https://developer.mozilla.org/en-US/docs/Web/API/OffscreenCanvas/transferToImageBitmap "The OffscreenCanvas.transferToImageBitmap() method creates an ImageBitmap object from the most recently rendered image of the OffscreenCanvas. The OffscreenCanvas allocates a new image for its subsequent rendering.")

Support in all current engines.

::: support
[Firefox105+]{.firefox .yes}[Safari16.4+]{.safari
.yes}[Chrome69+]{.chrome .yes}

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

Returns a newly created
[`ImageBitmap`{#the-offscreencanvas-interface:imagebitmap-2}](imagebitmap-and-animations.html#imagebitmap)
object with the image in the
[`OffscreenCanvas`{#the-offscreencanvas-interface:offscreencanvas-26}](#offscreencanvas)
object. The image in the
[`OffscreenCanvas`{#the-offscreencanvas-interface:offscreencanvas-27}](#offscreencanvas)
object is replaced with a new blank image.

The
[`convertToBlob(``options`{.variable}`)`]{#dom-offscreencanvas-converttoblob
.dfn dfn-for="OffscreenCanvas" dfn-type="method"} method steps are:

1.  If the value of
    [this](https://webidl.spec.whatwg.org/#this){#the-offscreencanvas-interface:this-15
    x-internal="this"}\'s
    [\[\[Detached\]\]](structured-data.html#detached){#the-offscreencanvas-interface:detached}
    internal slot is true, then return [a promise rejected
    with](https://webidl.spec.whatwg.org/#a-promise-rejected-with){#the-offscreencanvas-interface:a-promise-rejected-with
    x-internal="a-promise-rejected-with"} an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-offscreencanvas-interface:invalidstateerror-6
    x-internal="invalidstateerror"}
    [`DOMException`{#the-offscreencanvas-interface:domexception-6}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-offscreencanvas-interface:this-16
    x-internal="this"}\'s [context
    mode](#offscreencanvas-context-mode){#the-offscreencanvas-interface:offscreencanvas-context-mode-9}
    is
    [2d](#offscreencanvas-context-2d){#the-offscreencanvas-interface:offscreencanvas-context-2d-4}
    and the rendering context\'s [output
    bitmap](#output-bitmap){#the-offscreencanvas-interface:output-bitmap}\'s
    [origin-clean](#concept-canvas-origin-clean){#the-offscreencanvas-interface:concept-canvas-origin-clean}
    flag is set to false, then return [a promise rejected
    with](https://webidl.spec.whatwg.org/#a-promise-rejected-with){#the-offscreencanvas-interface:a-promise-rejected-with-2
    x-internal="a-promise-rejected-with"} a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#the-offscreencanvas-interface:securityerror
    x-internal="securityerror"}
    [`DOMException`{#the-offscreencanvas-interface:domexception-7}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

3.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-offscreencanvas-interface:this-17
    x-internal="this"}\'s
    [bitmap](#offscreencanvas-bitmap){#the-offscreencanvas-interface:offscreencanvas-bitmap-9}
    has no pixels (i.e., either its horizontal dimension or its vertical
    dimension is zero), then return [a promise rejected
    with](https://webidl.spec.whatwg.org/#a-promise-rejected-with){#the-offscreencanvas-interface:a-promise-rejected-with-3
    x-internal="a-promise-rejected-with"} an
    [\"`IndexSizeError`\"](https://webidl.spec.whatwg.org/#indexsizeerror){#the-offscreencanvas-interface:indexsizeerror
    x-internal="indexsizeerror"}
    [`DOMException`{#the-offscreencanvas-interface:domexception-8}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

4.  Let `bitmap`{.variable} be a copy of
    [this](https://webidl.spec.whatwg.org/#this){#the-offscreencanvas-interface:this-18
    x-internal="this"}\'s
    [bitmap](#offscreencanvas-bitmap){#the-offscreencanvas-interface:offscreencanvas-bitmap-10}.

5.  Let `result`{.variable} be a new promise object.

6.  Let `global`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#the-offscreencanvas-interface:this-19
    x-internal="this"}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#the-offscreencanvas-interface:concept-relevant-global-2}.

7.  Run these steps [in
    parallel](infrastructure.html#in-parallel){#the-offscreencanvas-interface:in-parallel}:

    1.  Let `file`{.variable} be [a serialization of `bitmap`{.variable}
        as a
        file](#a-serialisation-of-the-bitmap-as-a-file){#the-offscreencanvas-interface:a-serialisation-of-the-bitmap-as-a-file},
        with `options`{.variable}\'s [`type`]{#image-encode-options-type
        .dfn} and [`quality`]{#image-encode-options-quality .dfn} if
        present.

    2.  [Queue a global
        task](webappapis.html#queue-a-global-task){#the-offscreencanvas-interface:queue-a-global-task}
        on the [canvas blob serialization task
        source](#canvas-blob-serialisation-task-source){#the-offscreencanvas-interface:canvas-blob-serialisation-task-source}
        given `global`{.variable} to run these steps:

        1.  If `file`{.variable} is null, then reject
            `result`{.variable} with an
            [\"`EncodingError`\"](https://webidl.spec.whatwg.org/#encodingerror){#the-offscreencanvas-interface:encodingerror
            x-internal="encodingerror"}
            [`DOMException`{#the-offscreencanvas-interface:domexception-9}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

        2.  Otherwise, resolve `result`{.variable} with a new
            [`Blob`{#the-offscreencanvas-interface:blob-3}](https://w3c.github.io/FileAPI/#dfn-Blob){x-internal="blob"}
            object, created in `global`{.variable}\'s [relevant
            realm](webappapis.html#concept-relevant-realm){#the-offscreencanvas-interface:concept-relevant-realm},
            representing `file`{.variable}.
            [\[FILEAPI\]](references.html#refsFILEAPI)

8.  Return `result`{.variable}.

The
[`transferToImageBitmap()`]{#dom-offscreencanvas-transfertoimagebitmap
.dfn dfn-for="OffscreenCanvas" dfn-type="method"} method, when invoked,
must run the following steps:

1.  If the value of this
    [`OffscreenCanvas`{#the-offscreencanvas-interface:offscreencanvas-28}](#offscreencanvas)
    object\'s
    [\[\[Detached\]\]](structured-data.html#detached){#the-offscreencanvas-interface:detached-2}
    internal slot is set to true, then throw an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-offscreencanvas-interface:invalidstateerror-7
    x-internal="invalidstateerror"}
    [`DOMException`{#the-offscreencanvas-interface:domexception-10}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  If this
    [`OffscreenCanvas`{#the-offscreencanvas-interface:offscreencanvas-29}](#offscreencanvas)
    object\'s [context
    mode](#offscreencanvas-context-mode){#the-offscreencanvas-interface:offscreencanvas-context-mode-10}
    is set to
    [none](#offscreencanvas-context-none){#the-offscreencanvas-interface:offscreencanvas-context-none-3},
    then throw an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-offscreencanvas-interface:invalidstateerror-8
    x-internal="invalidstateerror"}
    [`DOMException`{#the-offscreencanvas-interface:domexception-11}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

3.  Let `image`{.variable} be a newly created
    [`ImageBitmap`{#the-offscreencanvas-interface:imagebitmap-3}](imagebitmap-and-animations.html#imagebitmap)
    object that references the same underlying bitmap data as this
    [`OffscreenCanvas`{#the-offscreencanvas-interface:offscreencanvas-30}](#offscreencanvas)
    object\'s
    [bitmap](#offscreencanvas-bitmap){#the-offscreencanvas-interface:offscreencanvas-bitmap-11}.

4.  Set this
    [`OffscreenCanvas`{#the-offscreencanvas-interface:offscreencanvas-31}](#offscreencanvas)
    object\'s
    [bitmap](#offscreencanvas-bitmap){#the-offscreencanvas-interface:offscreencanvas-bitmap-12}
    to reference a newly created bitmap of the same dimensions and color
    space as the previous bitmap, and with its pixels initialized to
    [transparent
    black](https://drafts.csswg.org/css-color/#transparent-black){#the-offscreencanvas-interface:transparent-black-5
    x-internal="transparent-black"}, or [opaque
    black](https://drafts.csswg.org/css-color/#opaque-black){#the-offscreencanvas-interface:opaque-black
    x-internal="opaque-black"} if the rendering context\'s
    [alpha](#concept-canvas-alpha){#the-offscreencanvas-interface:concept-canvas-alpha}
    is false.

    This means that if the rendering context of this
    [`OffscreenCanvas`{#the-offscreencanvas-interface:offscreencanvas-32}](#offscreencanvas)
    is a
    [`WebGLRenderingContext`{#the-offscreencanvas-interface:webglrenderingcontext-2}](https://www.khronos.org/registry/webgl/specs/latest/1.0/#WebGLRenderingContext){x-internal="webglrenderingcontext"},
    the value of
    [`preserveDrawingBuffer`{#the-offscreencanvas-interface:webglcontextattributes}](https://www.khronos.org/registry/webgl/specs/latest/1.0/#WebGLContextAttributes){x-internal="webglcontextattributes"}
    will have no effect. [\[WEBGL\]](references.html#refsWEBGL)

5.  Return `image`{.variable}.

The following are the [event
handlers](webappapis.html#event-handlers){#the-offscreencanvas-interface:event-handlers}
(and their corresponding [event handler event
types](webappapis.html#event-handler-event-type){#the-offscreencanvas-interface:event-handler-event-type})
that must be supported, as [event handler IDL
attributes](webappapis.html#event-handler-idl-attributes){#the-offscreencanvas-interface:event-handler-idl-attributes},
by all objects implementing the
[`OffscreenCanvas`{#the-offscreencanvas-interface:offscreencanvas-33}](#offscreencanvas)
interface:

[Event
handler](webappapis.html#event-handlers){#the-offscreencanvas-interface:event-handlers-2}

[Event handler event
type](webappapis.html#event-handler-event-type){#the-offscreencanvas-interface:event-handler-event-type-2}

[`oncontextlost`]{#handler-offscreencanvas-oncontextlost .dfn
dfn-for="OffscreenCanvas" dfn-type="attribute"}

[`contextlost`{#the-offscreencanvas-interface:event-contextlost-2}](indices.html#event-contextlost)

[`oncontextrestored`]{#handler-offscreencanvas-oncontextrestored .dfn
dfn-for="OffscreenCanvas" dfn-type="attribute"}

[`contextrestored`{#the-offscreencanvas-interface:event-contextrestored-2}](indices.html#event-contextrestored)

###### [4.12.5.3.1]{.secno} The offscreen 2D rendering context[](#the-offscreen-2d-rendering-context){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[OffscreenCanvasRenderingContext2D](https://developer.mozilla.org/en-US/docs/Web/API/OffscreenCanvasRenderingContext2D "The OffscreenCanvasRenderingContext2D interface is a CanvasRenderingContext2D rendering context for drawing to the bitmap of an OffscreenCanvas object. It is similar to the CanvasRenderingContext2D object, with the following differences:")

Support in all current engines.

::: support
[Firefox105+]{.firefox .yes}[Safari16.4+]{.safari
.yes}[Chrome69+]{.chrome .yes}

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

``` idl
[Exposed=(Window,Worker)]
interface OffscreenCanvasRenderingContext2D {
  readonly attribute OffscreenCanvas canvas;
};

OffscreenCanvasRenderingContext2D includes CanvasSettings;
OffscreenCanvasRenderingContext2D includes CanvasState;
OffscreenCanvasRenderingContext2D includes CanvasTransform;
OffscreenCanvasRenderingContext2D includes CanvasCompositing;
OffscreenCanvasRenderingContext2D includes CanvasImageSmoothing;
OffscreenCanvasRenderingContext2D includes CanvasFillStrokeStyles;
OffscreenCanvasRenderingContext2D includes CanvasShadowStyles;
OffscreenCanvasRenderingContext2D includes CanvasFilters;
OffscreenCanvasRenderingContext2D includes CanvasRect;
OffscreenCanvasRenderingContext2D includes CanvasDrawPath;
OffscreenCanvasRenderingContext2D includes CanvasText;
OffscreenCanvasRenderingContext2D includes CanvasDrawImage;
OffscreenCanvasRenderingContext2D includes CanvasImageData;
OffscreenCanvasRenderingContext2D includes CanvasPathDrawingStyles;
OffscreenCanvasRenderingContext2D includes CanvasTextDrawingStyles;
OffscreenCanvasRenderingContext2D includes CanvasPath;
```

The
[`OffscreenCanvasRenderingContext2D`{#the-offscreen-2d-rendering-context:offscreencanvasrenderingcontext2d-17}](#offscreencanvasrenderingcontext2d)
object is a rendering context for drawing to the
[bitmap](#offscreencanvas-bitmap){#the-offscreen-2d-rendering-context:offscreencanvas-bitmap}
of an
[`OffscreenCanvas`{#the-offscreen-2d-rendering-context:offscreencanvas-2}](#offscreencanvas)
object. It is similar to the
[`CanvasRenderingContext2D`{#the-offscreen-2d-rendering-context:canvasrenderingcontext2d}](#canvasrenderingcontext2d)
object, with the following differences:

- there is no support for [user
  interface](#canvasuserinterface){#the-offscreen-2d-rendering-context:canvasuserinterface}
  features;

- its
  [`canvas`{#the-offscreen-2d-rendering-context:offscreencontext2d-canvas-2}](#offscreencontext2d-canvas)
  attribute refers to an
  [`OffscreenCanvas`{#the-offscreen-2d-rendering-context:offscreencanvas-3}](#offscreencanvas)
  object rather than a
  [`canvas`{#the-offscreen-2d-rendering-context:the-canvas-element}](#the-canvas-element)
  element;

An
[`OffscreenCanvasRenderingContext2D`{#the-offscreen-2d-rendering-context:offscreencanvasrenderingcontext2d-18}](#offscreencanvasrenderingcontext2d)
object has an [associated `OffscreenCanvas`
object]{#associated-offscreencanvas-object .dfn}, which is the
[`OffscreenCanvas`{#the-offscreen-2d-rendering-context:offscreencanvas-4}](#offscreencanvas)
object from which the
[`OffscreenCanvasRenderingContext2D`{#the-offscreen-2d-rendering-context:offscreencanvasrenderingcontext2d-19}](#offscreencanvasrenderingcontext2d)
object was created.

`offscreenCanvas`{.variable}` = ``offscreenCanvasRenderingContext2D`{.variable}`.`[`canvas`](#offscreencontext2d-canvas){#offscreencontext2d-canvas-dev}

:   Returns the [associated `OffscreenCanvas`
    object](#associated-offscreencanvas-object){#the-offscreen-2d-rendering-context:associated-offscreencanvas-object}.

The [offscreen 2D context creation
algorithm]{#offscreen-2d-context-creation-algorithm .dfn}, which is
passed a `target`{.variable} (an
[`OffscreenCanvas`{#the-offscreen-2d-rendering-context:offscreencanvas-5}](#offscreencanvas)
object) and optionally some arguments, consists of running the following
steps:

1.  If the algorithm was passed some arguments, let `arg`{.variable} be
    the first such argument. Otherwise, let `arg`{.variable} be
    undefined.

2.  Let `settings`{.variable} be the result of
    [converting](https://webidl.spec.whatwg.org/#es-type-mapping){#the-offscreen-2d-rendering-context:concept-idl-convert
    x-internal="concept-idl-convert"} `arg`{.variable} to the dictionary
    type
    [`CanvasRenderingContext2DSettings`{#the-offscreen-2d-rendering-context:canvasrenderingcontext2dsettings}](#canvasrenderingcontext2dsettings).
    (This can throw an exception.)

3.  Let `context`{.variable} be a new
    [`OffscreenCanvasRenderingContext2D`{#the-offscreen-2d-rendering-context:offscreencanvasrenderingcontext2d-20}](#offscreencanvasrenderingcontext2d)
    object.

4.  Set `context`{.variable}\'s [associated `OffscreenCanvas`
    object](#associated-offscreencanvas-object){#the-offscreen-2d-rendering-context:associated-offscreencanvas-object-2}
    to `target`{.variable}.

5.  Run the [canvas settings output bitmap initialization
    algorithm](#canvas-setting-init-bitmap){#the-offscreen-2d-rendering-context:canvas-setting-init-bitmap},
    given `context`{.variable} and `settings`{.variable}.

6.  Set `context`{.variable}\'s [output
    bitmap](#output-bitmap){#the-offscreen-2d-rendering-context:output-bitmap}
    to a newly created bitmap with the dimensions specified by the
    [`width`{#the-offscreen-2d-rendering-context:dom-offscreencanvas-width}](#dom-offscreencanvas-width)
    and
    [`height`{#the-offscreen-2d-rendering-context:dom-offscreencanvas-height}](#dom-offscreencanvas-height)
    attributes of `target`{.variable}, and set `target`{.variable}\'s
    bitmap to the same bitmap (so that they are shared).

7.  If `context`{.variable}\'s
    [alpha](#concept-canvas-alpha){#the-offscreen-2d-rendering-context:concept-canvas-alpha}
    flag is set to true, initialize all the pixels of
    `context`{.variable}\'s [output
    bitmap](#output-bitmap){#the-offscreen-2d-rendering-context:output-bitmap-2}
    to [transparent
    black](https://drafts.csswg.org/css-color/#transparent-black){#the-offscreen-2d-rendering-context:transparent-black
    x-internal="transparent-black"}. Otherwise, initialize the pixels to
    [opaque
    black](https://drafts.csswg.org/css-color/#opaque-black){#the-offscreen-2d-rendering-context:opaque-black
    x-internal="opaque-black"}.

8.  Return `context`{.variable}.

Implementations are encouraged to short-circuit the graphics update
steps of the [window event
loop](webappapis.html#window-event-loop){#the-offscreen-2d-rendering-context:window-event-loop}
for the purposes of updating the contents of a [placeholder `canvas`
element](#offscreencanvas-placeholder){#the-offscreen-2d-rendering-context:offscreencanvas-placeholder}
to the display. This could mean, for example, that the bitmap contents
are copied directly to a graphics buffer that is mapped to the physical
display location of the [placeholder `canvas`
element](#offscreencanvas-placeholder){#the-offscreen-2d-rendering-context:offscreencanvas-placeholder-2}.
This or similar short-circuiting approaches can significantly reduce
display latency, especially in cases where the
[`OffscreenCanvas`{#the-offscreen-2d-rendering-context:offscreencanvas-6}](#offscreencanvas)
is updated from a [worker event
loop](webappapis.html#worker-event-loop-2){#the-offscreen-2d-rendering-context:worker-event-loop-2}
and the [window event
loop](webappapis.html#window-event-loop){#the-offscreen-2d-rendering-context:window-event-loop-2}
of the [placeholder `canvas`
element](#offscreencanvas-placeholder){#the-offscreen-2d-rendering-context:offscreencanvas-placeholder-3}
is busy. However, such shortcuts cannot have any script-observable
side-effects. This means that the committed bitmap still needs to be
sent to the [placeholder `canvas`
element](#offscreencanvas-placeholder){#the-offscreen-2d-rendering-context:offscreencanvas-placeholder-4},
in case the element is used as a
[`CanvasImageSource`{#the-offscreen-2d-rendering-context:canvasimagesource}](#canvasimagesource),
as an
[`ImageBitmapSource`{#the-offscreen-2d-rendering-context:imagebitmapsource}](imagebitmap-and-animations.html#imagebitmapsource),
or in case
[`toDataURL()`{#the-offscreen-2d-rendering-context:dom-canvas-todataurl}](#dom-canvas-todataurl)
or
[`toBlob()`{#the-offscreen-2d-rendering-context:dom-canvas-toblob}](#dom-canvas-toblob)
are called on it.

The [`canvas`]{#offscreencontext2d-canvas .dfn dfn-for="OffscreenCanvas"
dfn-type="attribute"} attribute, on getting, must return this
[`OffscreenCanvasRenderingContext2D`{#the-offscreen-2d-rendering-context:offscreencanvasrenderingcontext2d-21}](#offscreencanvasrenderingcontext2d)\'s
[associated `OffscreenCanvas`
object](#associated-offscreencanvas-object){#the-offscreen-2d-rendering-context:associated-offscreencanvas-object-3}.

##### [4.12.5.4]{.secno} Color spaces and color space conversion[](#colour-spaces-and-colour-correction){.self-link} {#colour-spaces-and-colour-correction}

The
[`canvas`{#colour-spaces-and-colour-correction:the-canvas-element}](#the-canvas-element)
APIs provide mechanisms for specifying the color space of the canvas\'s
backing store. The default backing store color space for all canvas APIs
is
[\'srgb\'](https://drafts.csswg.org/css-color/#valdef-color-srgb){#colour-spaces-and-colour-correction:'srgb'
x-internal="'srgb'"}.

[Color space
conversion](https://drafts.csswg.org/css-color/#color-conversion){#colour-spaces-and-colour-correction:converting-colors
x-internal="converting-colors"} must be applied to the canvas\'s backing
store when rendering the canvas to the output device. This color space
conversion must be identical to the color space conversion that would be
applied to an
[`img`{#colour-spaces-and-colour-correction:the-img-element}](embedded-content.html#the-img-element)
element with a color profile that specifies the same [color
space](#concept-canvas-color-space){#colour-spaces-and-colour-correction:concept-canvas-color-space}
as the canvas\'s backing store.

When drawing content to a 2D context, all inputs must be
[converted](https://drafts.csswg.org/css-color/#color-conversion){#colour-spaces-and-colour-correction:converting-colors-2
x-internal="converting-colors"} to the [context\'s color
space](#concept-canvas-color-space){#colour-spaces-and-colour-correction:concept-canvas-color-space-2}
before drawing. Interpolation of gradient color stops must be performed
on color values after conversion to the [context\'s color
space](#concept-canvas-color-space){#colour-spaces-and-colour-correction:concept-canvas-color-space-3}.
Alpha blending must be performed on values after conversion to the
[context\'s color
space](#concept-canvas-color-space){#colour-spaces-and-colour-correction:concept-canvas-color-space-4}.

There do not exist any inputs to a 2D context for which the color space
is undefined. The color space for CSS colors is defined in CSS Color.
The color space for images that specify no color profile information is
assumed to be
[\'srgb\'](https://drafts.csswg.org/css-color/#valdef-color-srgb){#colour-spaces-and-colour-correction:'srgb'-2
x-internal="'srgb'"}, as specified in the [Color Spaces of Untagged
Colors](https://drafts.csswg.org/css-color/#untagged) section of CSS
Color. [\[CSSCOLOR\]](references.html#refsCSSCOLOR)

##### [4.12.5.5]{.secno} Serializing bitmaps to a file[](#serialising-bitmaps-to-a-file){.self-link} {#serialising-bitmaps-to-a-file}

When a user agent is to create [a serialization of the bitmap as a
file]{#a-serialisation-of-the-bitmap-as-a-file .dfn}, given a
`type`{.variable} and an optional `quality`{.variable}, it must create
an image file in the format given by `type`{.variable}. If an error
occurs during the creation of the image file (e.g. an internal encoder
error), then the result of the serialization is null.
[\[PNG\]](references.html#refsPNG)

The image file\'s pixel data must be the bitmap\'s pixel data scaled to
one image pixel per coordinate space unit, and if the file format used
supports encoding resolution metadata, the resolution must be given as
96dpi (one image pixel per [CSS
pixel](https://drafts.csswg.org/css-values/#px){#serialising-bitmaps-to-a-file:'px'
x-internal="'px'"}).

If `type`{.variable} is supplied, then it must be interpreted as a [MIME
type](https://mimesniff.spec.whatwg.org/#mime-type){#serialising-bitmaps-to-a-file:mime-type
x-internal="mime-type"} giving the format to use. If the type has any
parameters, then it must be treated as not supported.

For example, the value
\"[`image/png`{#serialising-bitmaps-to-a-file:image/png}](indices.html#image/png)\"
would mean to generate a PNG image, the value
\"[`image/jpeg`{#serialising-bitmaps-to-a-file:image/jpeg}](indices.html#image/jpeg)\"
would mean to generate a JPEG image, and the value
\"[`image/svg+xml`{#serialising-bitmaps-to-a-file:image/svg+xml}](indices.html#image/svg+xml)\"
would mean to generate an SVG image (which would require that the user
agent track how the bitmap was generated, an unlikely, though
potentially awesome, feature).

User agents must support PNG
(\"[`image/png`{#serialising-bitmaps-to-a-file:image/png-2}](indices.html#image/png)\").
User agents may support other types. If the user agent does not support
the requested type, then it must create the file using the PNG format.
[\[PNG\]](references.html#refsPNG)

User agents must [convert the provided type to ASCII
lowercase](https://infra.spec.whatwg.org/#ascii-lowercase){#serialising-bitmaps-to-a-file:converted-to-ascii-lowercase
x-internal="converted-to-ascii-lowercase"} before establishing if they
support that type.

For image types that do not support an alpha component, the serialized
image must be the bitmap image composited onto an [opaque
black](https://drafts.csswg.org/css-color/#opaque-black){#serialising-bitmaps-to-a-file:opaque-black
x-internal="opaque-black"} background using the
[source-over](https://drafts.fxtf.org/compositing/#porterduffcompositingoperators_srcover){#serialising-bitmaps-to-a-file:gcop-source-over
x-internal="gcop-source-over"} compositing operator.

For image types that support color profiles, the serialized image must
include a color profile indicating the color space of the underlying
bitmap. For image types that do not support color profiles, the
serialized image must be
[converted](https://drafts.csswg.org/css-color/#color-conversion){#serialising-bitmaps-to-a-file:converting-colors
x-internal="converting-colors"} to the
[\'srgb\'](https://drafts.csswg.org/css-color/#valdef-color-srgb){#serialising-bitmaps-to-a-file:'srgb'
x-internal="'srgb'"} color space using
[\'relative-colorimetric\'](https://drafts.csswg.org/css-color-5/#valdef-color-profile-rendering-intent-relative-colorimetric){#serialising-bitmaps-to-a-file:'relative-colorimetric'
x-internal="'relative-colorimetric'"} rendering intent.

Thus, in the 2D context, calling the
[`drawImage()`{#serialising-bitmaps-to-a-file:dom-context-2d-drawimage}](#dom-context-2d-drawimage)
method to render the output of the
[`toDataURL()`{#serialising-bitmaps-to-a-file:dom-canvas-todataurl}](#dom-canvas-todataurl)
or
[`toBlob()`{#serialising-bitmaps-to-a-file:dom-canvas-toblob}](#dom-canvas-toblob)
method to the canvas, given the appropriate dimensions, has no visible
effect beyond, at most, clipping colors of the canvas to a more narrow
gamut.

For image types that support multiple bit depths, the serialized image
must use the bit depth that best preserves content of the underlying
bitmap.

For example, when serializing a 2D context that has [color
type](#concept-canvas-color-type){#serialising-bitmaps-to-a-file:concept-canvas-color-type}
of
[float16](#dom-canvascolortype-float16){#serialising-bitmaps-to-a-file:dom-canvascolortype-float16}
to `type`{.variable}
\"[`image/png`{#serialising-bitmaps-to-a-file:image/png-3}](indices.html#image/png)\",
the resulting image would have 16 bits per sample. This serialization
will still lose significant detail (all values less than 0.5/65535 would
be clamped to 0, and all values greater than 1 would be clamped to 1).

If `type`{.variable} is an image format that supports variable quality
(such as
\"[`image/jpeg`{#serialising-bitmaps-to-a-file:image/jpeg-2}](indices.html#image/jpeg)\"),
`quality`{.variable} is given, and `type`{.variable} is not
\"[`image/png`{#serialising-bitmaps-to-a-file:image/png-4}](indices.html#image/png)\",
then, if `quality`{.variable} [is a
Number](https://tc39.es/ecma262/#sec-ecmascript-language-types-number-type){#serialising-bitmaps-to-a-file:js-number
x-internal="js-number"} in the range 0.0 to 1.0 inclusive, the user
agent must treat `quality`{.variable} as the desired quality level.
Otherwise, the user agent must use its default quality value, as if the
`quality`{.variable} argument had not been given.

The use of type-testing here, instead of simply declaring
`quality`{.variable} as a Web IDL `double`, is a historical artifact.

Different implementations can have slightly different interpretations of
\"quality\". When the quality is not specified, an
implementation-specific default is used that represents a reasonable
compromise between compression ratio, image quality, and encoding time.

##### [4.12.5.6]{.secno} Security with [`canvas`{#security-with-canvas-elements:the-canvas-element}](#the-canvas-element) elements[](#security-with-canvas-elements){.self-link}

*This section is non-normative.*

**Information leakage** can occur if scripts from one
[origin](browsers.html#concept-origin){#security-with-canvas-elements:concept-origin}
can access information (e.g. read pixels) from images from another
origin (one that isn\'t the
[same](browsers.html#same-origin){#security-with-canvas-elements:same-origin}).

To mitigate this, bitmaps used with
[`canvas`{#security-with-canvas-elements:the-canvas-element-2}](#the-canvas-element)
elements,
[`OffscreenCanvas`{#security-with-canvas-elements:offscreencanvas}](#offscreencanvas)
objects, and
[`ImageBitmap`{#security-with-canvas-elements:imagebitmap}](imagebitmap-and-animations.html#imagebitmap)
objects are defined to have a flag indicating whether they are
[origin-clean](#concept-canvas-origin-clean){#security-with-canvas-elements:concept-canvas-origin-clean}.
All bitmaps start with their
[origin-clean](#concept-canvas-origin-clean){#security-with-canvas-elements:concept-canvas-origin-clean-2}
set to true. The flag is set to false when cross-origin images are used.

The
[`toDataURL()`{#security-with-canvas-elements:dom-canvas-todataurl}](#dom-canvas-todataurl),
[`toBlob()`{#security-with-canvas-elements:dom-canvas-toblob}](#dom-canvas-toblob),
and
[`getImageData()`{#security-with-canvas-elements:dom-context-2d-getimagedata}](#dom-context-2d-getimagedata)
methods check the flag and will throw a
[\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#security-with-canvas-elements:securityerror
x-internal="securityerror"}
[`DOMException`{#security-with-canvas-elements:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
rather than leak cross-origin data.

The value of the
[origin-clean](#concept-canvas-origin-clean){#security-with-canvas-elements:concept-canvas-origin-clean-3}
flag is propagated from a source\'s bitmap to a new
[`ImageBitmap`{#security-with-canvas-elements:imagebitmap-2}](imagebitmap-and-animations.html#imagebitmap)
object by
[`createImageBitmap()`{#security-with-canvas-elements:dom-createimagebitmap}](imagebitmap-and-animations.html#dom-createimagebitmap).
Conversely, a destination
[`canvas`{#security-with-canvas-elements:the-canvas-element-3}](#the-canvas-element)
element\'s bitmap will have its
[origin-clean](#concept-canvas-origin-clean){#security-with-canvas-elements:concept-canvas-origin-clean-4}
flags set to false by
[`drawImage`{#security-with-canvas-elements:dom-context-2d-drawimage}](#dom-context-2d-drawimage)
if the source image is an
[`ImageBitmap`{#security-with-canvas-elements:imagebitmap-3}](imagebitmap-and-animations.html#imagebitmap)
object whose bitmap has its
[origin-clean](#concept-canvas-origin-clean){#security-with-canvas-elements:concept-canvas-origin-clean-5}
flag set to false.

The flag can be reset in certain situations; for example, when changing
the value of the
[`width`{#security-with-canvas-elements:attr-canvas-width}](#attr-canvas-width)
or the
[`height`{#security-with-canvas-elements:attr-canvas-height}](#attr-canvas-height)
content attribute of the
[`canvas`{#security-with-canvas-elements:the-canvas-element-4}](#the-canvas-element)
element to which a
[`CanvasRenderingContext2D`{#security-with-canvas-elements:canvasrenderingcontext2d}](#canvasrenderingcontext2d)
is bound, the bitmap is cleared and its
[origin-clean](#concept-canvas-origin-clean){#security-with-canvas-elements:concept-canvas-origin-clean-6}
flag is reset.

When using an
[`ImageBitmapRenderingContext`{#security-with-canvas-elements:imagebitmaprenderingcontext}](#imagebitmaprenderingcontext),
the value of the
[origin-clean](#concept-canvas-origin-clean){#security-with-canvas-elements:concept-canvas-origin-clean-7}
flag is propagated from
[`ImageBitmap`{#security-with-canvas-elements:imagebitmap-4}](imagebitmap-and-animations.html#imagebitmap)
objects when they are transferred to the
[`canvas`{#security-with-canvas-elements:the-canvas-element-5}](#the-canvas-element)
via
[transferFromImageBitmap()](#dom-imagebitmaprenderingcontext-transferfromimagebitmap){#security-with-canvas-elements:dom-imagebitmaprenderingcontext-transferfromimagebitmap}.

##### [4.12.5.7]{.secno} Premultiplied alpha and the 2D rendering context[](#premultiplied-alpha-and-the-2d-rendering-context){.self-link}

[Premultiplied alpha]{#concept-premultiplied-alpha .dfn} refers to one
way of representing transparency in an image, the other being
non-premultiplied alpha.

Under non-premultiplied alpha, the red, green, and blue components of a
pixel represent that pixel\'s color, and its alpha component represents
that pixel\'s opacity.

Under premultiplied alpha, however, the red, green, and blue components
of a pixel represent the amounts of color that the pixel adds to the
image, and its alpha component represents the amount that the pixel
obscures whatever is behind it.

::: example
For instance, assuming the color components range from 0 (off) to 255
(full intensity), these example colors are represented in the following
ways:

CSS color representation

Premultiplied representation

Non-premultiplied representation

Description of color

Image of color blended above other content

rgba(255, 127, 0, 1)

255, 127, 0, 255

255, 127, 0, 255

Completely-opaque orange

![An opaque orange circle sits atop a
background](images/premultiplied-example-1.png){width="96" height="96"}

rgba(255, 255, 0, 0.5)

127, 127, 0, 127

255, 255, 0, 127

Halfway-opaque yellow

![A yellow circle, halfway transparent, sits atop a
background](images/premultiplied-example-2.png){width="96" height="96"}

Unrepresentable

255, 127, 0, 127

Unrepresentable

Additive halfway-opaque orange

![An orange circle somewhat brightens the background that it sits
atop](images/premultiplied-example-3.png){width="96" height="96"}

Unrepresentable

255, 127, 0, 0

Unrepresentable

Additive fully-transparent orange

![An orange circle completely brightens the background that it sits
atop](images/premultiplied-example-4.png){width="96" height="96"}

rgba(255, 127, 0, 0)

0, 0, 0, 0

255, 127, 0, 0

Fully-transparent (\"invisible\") orange

![An empty background with nothing atop
it](images/premultiplied-example-5.png){width="96" height="96"}

rgba(0, 127, 255, 0)

0, 0, 0, 0

255, 127, 0, 0

Fully-transparent (\"invisible\") turquoise

![An empty background with nothing atop
it](images/premultiplied-example-5.png){width="96" height="96"}
:::

[Converting a color value from a non-premultiplied representation to a
premultiplied one]{#convert-to-premultiplied .dfn} involves multiplying
the color\'s red, green, and blue components by its alpha component
(remapping the range of the alpha component such that \"fully
transparent\" is 0, and \"fully opaque\" is 1).

[Converting a color value from a premultiplied representation to a
non-premultiplied one]{#convert-from-premultiplied .dfn} involves the
inverse: dividing the color\'s red, green, and blue components by its
alpha component.

As certain colors can only be represented under premultiplied alpha (for
instance, additive colors), and others can only be represented under
non-premultiplied alpha (for instance, \"invisible\" colors which hold
certain red, green, and blue values even with no opacity); and division
and multiplication using finite precision entails a loss of accuracy,
converting between premultiplied and non-premultiplied alpha is a lossy
operation on colors that are not fully opaque.

A
[`CanvasRenderingContext2D`{#premultiplied-alpha-and-the-2d-rendering-context:canvasrenderingcontext2d}](#canvasrenderingcontext2d)\'s
[output
bitmap](#output-bitmap){#premultiplied-alpha-and-the-2d-rendering-context:output-bitmap}
and an
[`OffscreenCanvasRenderingContext2D`{#premultiplied-alpha-and-the-2d-rendering-context:offscreencanvasrenderingcontext2d}](#offscreencanvasrenderingcontext2d)\'s
[output
bitmap](#output-bitmap){#premultiplied-alpha-and-the-2d-rendering-context:output-bitmap-2}
must use premultiplied alpha to represent transparent colors.

It is important for canvas bitmaps to represent colors using
premultiplied alpha because it affects the range of representable
colors. While additive colors cannot currently be drawn onto canvases
directly because CSS colors are non-premultiplied and cannot represent
them, it is still possible to, for instance, draw additive colors onto a
WebGL canvas and then draw that WebGL canvas onto a 2D canvas via
[`drawImage()`{#premultiplied-alpha-and-the-2d-rendering-context:dom-context-2d-drawimage}](#dom-context-2d-drawimage).

[← 4.12 Scripting](scripting.html) --- [Table of Contents](./) --- [4.13
Custom elements →](custom-elements.html)
