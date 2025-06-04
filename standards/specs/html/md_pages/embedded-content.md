HTML Standard

::: {#head .head .with-buttons}
[![WHATWG](https://resources.whatwg.org/logo.svg){.darkmode-aware
width="100" crossorigin="" height="100"}](https://whatwg.org/){.logo}

# HTML {#html .allcaps}

Living Standard --- Last Updated [2 June 2025]{.pubdate}
:::

[← 4.7 Edits](edits.html) --- [Table of Contents](./) --- [4.8.4 Images
→](images.html)

1.  ::: {#toc-semantics}
    1.  [[4.8]{.secno} Embedded
        content](embedded-content.html#embedded-content)
        1.  [[4.8.1]{.secno} The `picture`
            element](embedded-content.html#the-picture-element)
        2.  [[4.8.2]{.secno} The `source`
            element](embedded-content.html#the-source-element)
        3.  [[4.8.3]{.secno} The `img`
            element](embedded-content.html#the-img-element)
    :::

### [4.8]{.secno} Embedded content[](#embedded-content){.self-link}

#### [4.8.1]{.secno} The [`picture`]{.dfn dfn-type="element"} element[](#the-picture-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/picture](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/picture "The <picture> HTML element contains zero or more <source> elements and one <img> element to offer alternative versions of an image for different display/device scenarios.")

Support in all current engines.

::: support
[Firefox38+]{.firefox .yes}[Safari9.1+]{.safari .yes}[Chrome38+]{.chrome
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

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLPictureElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLPictureElement "The HTMLPictureElement interface represents a <picture> HTML element. It doesn't implement specific properties or methods.")

Support in all current engines.

::: support
[Firefox38+]{.firefox .yes}[Safari9.1+]{.safari .yes}[Chrome38+]{.chrome
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

[Categories](dom.html#concept-element-categories){#the-picture-element:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-picture-element:flow-content-2}.
:   [Phrasing
    content](dom.html#phrasing-content-2){#the-picture-element:phrasing-content-2}.
:   [Embedded
    content](dom.html#embedded-content-category){#the-picture-element:embedded-content-category}.
:   [Palpable
    content](dom.html#palpable-content-2){#the-picture-element:palpable-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-picture-element:concept-element-contexts}:
:   Where [embedded
    content](dom.html#embedded-content-category){#the-picture-element:embedded-content-category-2}
    is expected.

[Content model](dom.html#concept-element-content-model){#the-picture-element:concept-element-content-model}:
:   Zero or more
    [`source`{#the-picture-element:the-source-element}](#the-source-element)
    elements, followed by one
    [`img`{#the-picture-element:the-img-element}](#the-img-element)
    element, optionally intermixed with [script-supporting
    elements](dom.html#script-supporting-elements-2){#the-picture-element:script-supporting-elements-2}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-picture-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-picture-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-picture-element:global-attributes}

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-picture-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-picture).
:   [For implementers](https://w3c.github.io/html-aam/#el-picture).

[DOM interface](dom.html#concept-element-dom){#the-picture-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLPictureElement : HTMLElement {
      [HTMLConstructor] constructor();
    };
    ```

The
[`picture`{#the-picture-element:the-picture-element}](#the-picture-element)
element is a container which provides multiple sources to its contained
[`img`{#the-picture-element:the-img-element-2}](#the-img-element)
element to allow authors to declaratively control or give hints to the
user agent about which image resource to use, based on the screen pixel
density,
[viewport](https://drafts.csswg.org/css2/#viewport){#the-picture-element:viewport
x-internal="viewport"} size, image format, and other factors. It
[represents](dom.html#represents){#the-picture-element:represents} its
children.

The
[`picture`{#the-picture-element:the-picture-element-2}](#the-picture-element)
element is somewhat different from the similar-looking
[`video`{#the-picture-element:the-video-element}](media.html#the-video-element)
and
[`audio`{#the-picture-element:the-audio-element}](media.html#the-audio-element)
elements. While all of them contain
[`source`{#the-picture-element:the-source-element-2}](#the-source-element)
elements, the
[`source`{#the-picture-element:the-source-element-3}](#the-source-element)
element\'s
[`src`{#the-picture-element:attr-source-src}](#attr-source-src)
attribute has no meaning when the element is nested within a
[`picture`{#the-picture-element:the-picture-element-3}](#the-picture-element)
element, and the resource selection algorithm is different. Also, the
[`picture`{#the-picture-element:the-picture-element-4}](#the-picture-element)
element itself does not display anything; it merely provides a context
for its contained
[`img`{#the-picture-element:the-img-element-3}](#the-img-element)
element that enables it to choose from multiple
[URLs](https://url.spec.whatwg.org/#concept-url){#the-picture-element:url
x-internal="url"}.

#### [4.8.2]{.secno} The [`source`]{.dfn dfn-type="element"} element[](#the-source-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/source](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/source "The <source> HTML element specifies multiple media resources for the <picture>, the <audio> element, or the <video> element. It is a void element, meaning that it has no content and does not have a closing tag. It is commonly used to offer the same media content in multiple file formats in order to provide compatibility with a broad range of browsers given their differing support for image file formats and media file formats.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

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
[HTMLSourceElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLSourceElement "The HTMLSourceElement interface provides special properties (beyond the regular HTMLElement object interface it also has available to it by inheritance) for manipulating <source> elements.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
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

[Categories](dom.html#concept-element-categories){#the-source-element:concept-element-categories}:
:   None.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-source-element:concept-element-contexts}:
:   As a child of a
    [`picture`{#the-source-element:the-picture-element}](#the-picture-element)
    element, before the
    [`img`{#the-source-element:the-img-element}](#the-img-element)
    element.
:   As a child of a [media
    element](media.html#media-element){#the-source-element:media-element},
    before any [flow
    content](dom.html#flow-content-2){#the-source-element:flow-content-2}
    or
    [`track`{#the-source-element:the-track-element}](media.html#the-track-element)
    elements.

[Content model](dom.html#concept-element-content-model){#the-source-element:concept-element-content-model}:
:   [Nothing](dom.html#concept-content-nothing){#the-source-element:concept-content-nothing}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-source-element:concept-element-tag-omission}:
:   No [end
    tag](syntax.html#syntax-end-tag){#the-source-element:syntax-end-tag}.

[Content attributes](dom.html#concept-element-attributes){#the-source-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-source-element:global-attributes}
:   [`type`{#the-source-element:attr-source-type}](#attr-source-type)
    --- Type of embedded resource
:   [`media`{#the-source-element:attr-source-media}](#attr-source-media)
    --- Applicable media
:   [`src`{#the-source-element:attr-source-src}](#attr-source-src) (in
    [`audio`{#the-source-element:the-audio-element}](media.html#the-audio-element)
    or
    [`video`{#the-source-element:the-video-element}](media.html#the-video-element))
    --- Address of the resource
:   [`srcset`{#the-source-element:attr-source-srcset}](#attr-source-srcset)
    (in
    [`picture`{#the-source-element:the-picture-element-2}](#the-picture-element))
    --- Images to use in different situations, e.g., high-resolution
    displays, small monitors, etc.
:   [`sizes`{#the-source-element:attr-source-sizes}](#attr-source-sizes)
    (in
    [`picture`{#the-source-element:the-picture-element-3}](#the-picture-element))
    --- Image sizes for different page layouts
:   [`width`{#the-source-element:attr-dim-width}](embedded-content-other.html#attr-dim-width)
    (in
    [`picture`{#the-source-element:the-picture-element-4}](#the-picture-element))
    --- Horizontal dimension
:   [`height`{#the-source-element:attr-dim-height}](embedded-content-other.html#attr-dim-height)
    (in
    [`picture`{#the-source-element:the-picture-element-5}](#the-picture-element))
    --- Vertical dimension

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-source-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-source).
:   [For implementers](https://w3c.github.io/html-aam/#el-source).

[DOM interface](dom.html#concept-element-dom){#the-source-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLSourceElement : HTMLElement {
      [HTMLConstructor] constructor();

      [CEReactions] attribute USVString src;
      [CEReactions] attribute DOMString type;
      [CEReactions] attribute USVString srcset;
      [CEReactions] attribute DOMString sizes;
      [CEReactions] attribute DOMString media;
      [CEReactions] attribute unsigned long width;
      [CEReactions] attribute unsigned long height;
    };
    ```

The
[`source`{#the-source-element:the-source-element}](#the-source-element)
element allows authors to specify multiple alternative [source
sets](images.html#source-set){#the-source-element:source-set} for
[`img`{#the-source-element:the-img-element-2}](#the-img-element)
elements or multiple alternative [media
resources](media.html#media-resource){#the-source-element:media-resource}
for [media
elements](media.html#media-element){#the-source-element:media-element-2}.
It does not
[represent](dom.html#represents){#the-source-element:represents}
anything on its own.

The [`type`]{#attr-source-type .dfn dfn-for="source"
dfn-type="element-attr"} attribute may be present. If present, the value
must be a [valid MIME type
string](https://mimesniff.spec.whatwg.org/#valid-mime-type){#the-source-element:valid-mime-type-string
x-internal="valid-mime-type-string"}.

The [`media`]{#attr-source-media .dfn dfn-for="source"
dfn-type="element-attr"} attribute may also be present. If present, the
value must contain a [valid media query
list](common-microsyntaxes.html#valid-media-query-list){#the-source-element:valid-media-query-list}.
The user agent will skip to the next
[`source`{#the-source-element:the-source-element-2}](#the-source-element)
element if the value does not [match the
environment](common-microsyntaxes.html#matches-the-environment){#the-source-element:matches-the-environment}.

The
[`media`{#the-source-element:attr-source-media-2}](#attr-source-media)
attribute is only evaluated once during the [resource selection
algorithm](media.html#concept-media-load-algorithm){#the-source-element:concept-media-load-algorithm}
for [media
elements](media.html#media-element){#the-source-element:media-element-3}.
In contrast, when using the
[`picture`{#the-source-element:the-picture-element-6}](#the-picture-element)
element, the user agent will [react to changes in the
environment](images.html#img-environment-changes).

The remainder of the requirements depend on whether the parent is a
[`picture`{#the-source-element:the-picture-element-7}](#the-picture-element)
element or a [media
element](media.html#media-element){#the-source-element:media-element-4}:

The [`source`{#the-source-element:the-source-element-3}](#the-source-element) element\'s parent is a [`picture`{#the-source-element:the-picture-element-8}](#the-picture-element) element

:   The [`srcset`]{#attr-source-srcset .dfn dfn-for="source"
    dfn-type="element-attr"} attribute must be present, and is a [srcset
    attribute](images.html#srcset-attribute){#the-source-element:srcset-attribute}.

    The
    [`srcset`{#the-source-element:attr-source-srcset-2}](#attr-source-srcset)
    attribute contributes the [image
    sources](images.html#image-source){#the-source-element:image-source}
    to the [source
    set](images.html#source-set){#the-source-element:source-set-2}, if
    the
    [`source`{#the-source-element:the-source-element-4}](#the-source-element)
    element is selected.

    If the
    [`srcset`{#the-source-element:attr-source-srcset-3}](#attr-source-srcset)
    attribute has any [image candidate
    strings](images.html#image-candidate-string){#the-source-element:image-candidate-string}
    using a [width
    descriptor](images.html#width-descriptor){#the-source-element:width-descriptor},
    the
    [`sizes`{#the-source-element:attr-source-sizes-2}](#attr-source-sizes)
    attribute may also be present. If, additionally, the following
    sibling
    [`img`{#the-source-element:the-img-element-3}](#the-img-element)
    element does not [allow
    auto-sizes](#allows-auto-sizes){#the-source-element:allows-auto-sizes},
    the
    [`sizes`{#the-source-element:attr-source-sizes-3}](#attr-source-sizes)
    attribute must be present. The [`sizes`]{#attr-source-sizes .dfn
    dfn-for="source" dfn-type="element-attr"} attribute is a [sizes
    attribute](images.html#sizes-attribute){#the-source-element:sizes-attribute},
    which contributes the [source
    size](images.html#source-size-2){#the-source-element:source-size-2}
    to the [source
    set](images.html#source-set){#the-source-element:source-set-3}, if
    the
    [`source`{#the-source-element:the-source-element-5}](#the-source-element)
    element is selected.

    If the
    [`img`{#the-source-element:the-img-element-4}](#the-img-element)
    element [allows
    auto-sizes](#allows-auto-sizes){#the-source-element:allows-auto-sizes-2},
    then the
    [`sizes`{#the-source-element:attr-source-sizes-4}](#attr-source-sizes)
    attribute can be omitted on previous sibling
    [`source`{#the-source-element:the-source-element-6}](#the-source-element)
    elements. In such cases, it is equivalent to specifying
    [`auto`{#the-source-element:valdef-sizes-auto}](images.html#valdef-sizes-auto).

    The
    [`source`{#the-source-element:the-source-element-7}](#the-source-element)
    element supports [dimension
    attributes](embedded-content-other.html#dimension-attributes){#the-source-element:dimension-attributes}.
    The [`img`{#the-source-element:the-img-element-5}](#the-img-element)
    element can use the
    [`width`{#the-source-element:attr-dim-width-2}](embedded-content-other.html#attr-dim-width)
    and
    [`height`{#the-source-element:attr-dim-height-2}](embedded-content-other.html#attr-dim-height)
    attributes of a
    [`source`{#the-source-element:the-source-element-8}](#the-source-element)
    element, instead of those on the
    [`img`{#the-source-element:the-img-element-6}](#the-img-element)
    element itself, to determine its rendered dimensions and
    aspect-ratio, [as defined in the Rendering
    section](rendering.html#dimRendering).

    The
    [`type`{#the-source-element:attr-source-type-2}](#attr-source-type)
    attribute gives the type of the images in the [source
    set](images.html#source-set){#the-source-element:source-set-4}, to
    allow the user agent to skip to the next
    [`source`{#the-source-element:the-source-element-9}](#the-source-element)
    element if it does not support the given type.

    If the
    [`type`{#the-source-element:attr-source-type-3}](#attr-source-type)
    attribute is *not* specified, the user agent will not select a
    different
    [`source`{#the-source-element:the-source-element-10}](#the-source-element)
    element if it finds that it does not support the image format after
    fetching it.

    When a
    [`source`{#the-source-element:the-source-element-11}](#the-source-element)
    element has a following sibling
    [`source`{#the-source-element:the-source-element-12}](#the-source-element)
    element or
    [`img`{#the-source-element:the-img-element-7}](#the-img-element)
    element with a
    [`srcset`{#the-source-element:attr-img-srcset}](#attr-img-srcset)
    attribute specified, it must have at least one of the following:

    - A
      [`media`{#the-source-element:attr-source-media-3}](#attr-source-media)
      attribute specified with a value that, after [stripping leading
      and trailing ASCII
      whitespace](https://infra.spec.whatwg.org/#strip-leading-and-trailing-ascii-whitespace){#the-source-element:strip-leading-and-trailing-ascii-whitespace
      x-internal="strip-leading-and-trailing-ascii-whitespace"}, is not
      the empty string and is not an [ASCII
      case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#the-source-element:ascii-case-insensitive
      x-internal="ascii-case-insensitive"} match for the string
      \"`all`\".

    - A
      [`type`{#the-source-element:attr-source-type-4}](#attr-source-type)
      attribute specified.

    The [`src`{#the-source-element:attr-source-src-2}](#attr-source-src)
    attribute must not be present.

The [`source`{#the-source-element:the-source-element-13}](#the-source-element) element\'s parent is a [media element](media.html#media-element){#the-source-element:media-element-5}

:   The [`src`]{#attr-source-src .dfn dfn-for="source"
    dfn-type="element-attr"} attribute gives the
    [URL](https://url.spec.whatwg.org/#concept-url){#the-source-element:url
    x-internal="url"} of the [media
    resource](media.html#media-resource){#the-source-element:media-resource-2}.
    The value must be a [valid non-empty URL potentially surrounded by
    spaces](urls-and-fetching.html#valid-non-empty-url-potentially-surrounded-by-spaces){#the-source-element:valid-non-empty-url-potentially-surrounded-by-spaces}.
    This attribute must be present.

    The
    [`type`{#the-source-element:attr-source-type-5}](#attr-source-type)
    attribute gives the type of the [media
    resource](media.html#media-resource){#the-source-element:media-resource-3},
    to help the user agent determine if it can play this [media
    resource](media.html#media-resource){#the-source-element:media-resource-4}
    before fetching it. The `codecs` parameter, which certain MIME types
    define, might be necessary to specify exactly how the resource is
    encoded. [\[RFC6381\]](references.html#refsRFC6381)

    Dynamically modifying a
    [`source`{#the-source-element:the-source-element-14}](#the-source-element)
    element\'s
    [`src`{#the-source-element:attr-source-src-3}](#attr-source-src) or
    [`type`{#the-source-element:attr-source-type-6}](#attr-source-type)
    attribute when the element is already inserted in a
    [`video`{#the-source-element:the-video-element-2}](media.html#the-video-element)
    or
    [`audio`{#the-source-element:the-audio-element-2}](media.html#the-audio-element)
    element will have no effect. To change what is playing, just use the
    [`src`{#the-source-element:attr-media-src}](media.html#attr-media-src)
    attribute on the [media
    element](media.html#media-element){#the-source-element:media-element-6}
    directly, possibly making use of the
    [`canPlayType()`{#the-source-element:dom-navigator-canplaytype}](media.html#dom-navigator-canplaytype)
    method to pick from amongst available resources. Generally,
    manipulating
    [`source`{#the-source-element:the-source-element-15}](#the-source-element)
    elements manually after the document has been parsed is an
    unnecessarily complicated approach.

    ::: example
    The following list shows some examples of how to use the `codecs=`
    MIME parameter in the
    [`type`{#the-source-element:attr-source-type-7}](#attr-source-type)
    attribute.

    H.264 Constrained baseline profile video (main and extended video compatible) level 3 and Low-Complexity AAC audio in MP4 container

    :   ``` html
        <source src='video.mp4' type='video/mp4; codecs="avc1.42E01E, mp4a.40.2"'>
        ```

    H.264 Extended profile video (baseline-compatible) level 3 and Low-Complexity AAC audio in MP4 container

    :   ``` html
        <source src='video.mp4' type='video/mp4; codecs="avc1.58A01E, mp4a.40.2"'>
        ```

    H.264 Main profile video level 3 and Low-Complexity AAC audio in MP4 container

    :   ``` html
        <source src='video.mp4' type='video/mp4; codecs="avc1.4D401E, mp4a.40.2"'>
        ```

    H.264 \'High\' profile video (incompatible with main, baseline, or extended profiles) level 3 and Low-Complexity AAC audio in MP4 container

    :   ``` html
        <source src='video.mp4' type='video/mp4; codecs="avc1.64001E, mp4a.40.2"'>
        ```

    MPEG-4 Visual Simple Profile Level 0 video and Low-Complexity AAC audio in MP4 container

    :   ``` html
        <source src='video.mp4' type='video/mp4; codecs="mp4v.20.8, mp4a.40.2"'>
        ```

    MPEG-4 Advanced Simple Profile Level 0 video and Low-Complexity AAC audio in MP4 container

    :   ``` html
        <source src='video.mp4' type='video/mp4; codecs="mp4v.20.240, mp4a.40.2"'>
        ```

    MPEG-4 Visual Simple Profile Level 0 video and AMR audio in 3GPP container

    :   ``` html
        <source src='video.3gp' type='video/3gpp; codecs="mp4v.20.8, samr"'>
        ```

    Theora video and Vorbis audio in Ogg container

    :   ``` html
        <source src='video.ogv' type='video/ogg; codecs="theora, vorbis"'>
        ```

    Theora video and Speex audio in Ogg container

    :   ``` html
        <source src='video.ogv' type='video/ogg; codecs="theora, speex"'>
        ```

    Vorbis audio alone in Ogg container

    :   ``` html
        <source src='audio.ogg' type='audio/ogg; codecs=vorbis'>
        ```

    Speex audio alone in Ogg container

    :   ``` html
        <source src='audio.spx' type='audio/ogg; codecs=speex'>
        ```

    FLAC audio alone in Ogg container

    :   ``` html
        <source src='audio.oga' type='audio/ogg; codecs=flac'>
        ```

    Dirac video and Vorbis audio in Ogg container

    :   ``` html
        <source src='video.ogv' type='video/ogg; codecs="dirac, vorbis"'>
        ```
    :::

    The
    [`srcset`{#the-source-element:attr-source-srcset-4}](#attr-source-srcset)
    and
    [`sizes`{#the-source-element:attr-source-sizes-5}](#attr-source-sizes)
    attributes must not be present.

The
[`source`{#the-source-element:the-source-element-16}](#the-source-element)
[HTML element insertion
steps](infrastructure.html#html-element-insertion-steps){#the-source-element:html-element-insertion-steps},
given `insertedNode`{.variable}, are:

1.  Let `parent`{.variable} be `insertedNode`{.variable}\'s
    [parent](https://dom.spec.whatwg.org/#concept-tree-parent){#the-source-element:parent
    x-internal="parent"}.

2.  If `parent`{.variable} is a [media
    element](media.html#media-element){#the-source-element:media-element-7}
    that has no
    [`src`{#the-source-element:attr-media-src-2}](media.html#attr-media-src)
    attribute and whose
    [`networkState`{#the-source-element:dom-media-networkstate}](media.html#dom-media-networkstate)
    has the value
    [`NETWORK_EMPTY`{#the-source-element:dom-media-network_empty}](media.html#dom-media-network_empty),
    then invoke that [media
    element](media.html#media-element){#the-source-element:media-element-8}\'s
    [resource selection
    algorithm](media.html#concept-media-load-algorithm){#the-source-element:concept-media-load-algorithm-2}.

3.  If `parent`{.variable} is a
    [`picture`{#the-source-element:the-picture-element-9}](#the-picture-element)
    element, then [for
    each](https://infra.spec.whatwg.org/#list-iterate){#the-source-element:list-iterate
    x-internal="list-iterate"} `child`{.variable} of
    `parent`{.variable}\'s
    [children](https://dom.spec.whatwg.org/#concept-tree-child){#the-source-element:concept-tree-child
    x-internal="concept-tree-child"}, if `child`{.variable} is an
    [`img`{#the-source-element:the-img-element-8}](#the-img-element)
    element, then count this as a [relevant
    mutation](images.html#relevant-mutations){#the-source-element:relevant-mutations}
    for `child`{.variable}.

The
[`source`{#the-source-element:the-source-element-17}](#the-source-element)
[HTML element moving
steps](infrastructure.html#html-element-moving-steps){#the-source-element:html-element-moving-steps},
given `movedNode`{.variable} and `oldParent`{.variable}, are:

1.  If `oldParent`{.variable} is a
    [`picture`{#the-source-element:the-picture-element-10}](#the-picture-element)
    element, then [for
    each](https://infra.spec.whatwg.org/#list-iterate){#the-source-element:list-iterate-2
    x-internal="list-iterate"} `child`{.variable} of
    `oldParent`{.variable}\'s
    [children](https://dom.spec.whatwg.org/#concept-tree-child){#the-source-element:concept-tree-child-2
    x-internal="concept-tree-child"}, if `child`{.variable} is an
    [`img`{#the-source-element:the-img-element-9}](#the-img-element)
    element, then count this as a [relevant
    mutation](images.html#relevant-mutations){#the-source-element:relevant-mutations-2}
    for `child`{.variable}.

The
[`source`{#the-source-element:the-source-element-18}](#the-source-element)
[HTML element removing
steps](infrastructure.html#html-element-removing-steps){#the-source-element:html-element-removing-steps},
given `removedNode`{.variable} and `oldParent`{.variable}, are:

1.  If `oldParent`{.variable} is a
    [`picture`{#the-source-element:the-picture-element-11}](#the-picture-element)
    element, then [for
    each](https://infra.spec.whatwg.org/#list-iterate){#the-source-element:list-iterate-3
    x-internal="list-iterate"} `child`{.variable} of
    `oldParent`{.variable}\'s
    [children](https://dom.spec.whatwg.org/#concept-tree-child){#the-source-element:concept-tree-child-3
    x-internal="concept-tree-child"}, if `child`{.variable} is an
    [`img`{#the-source-element:the-img-element-10}](#the-img-element)
    element, then count this as a [relevant
    mutation](images.html#relevant-mutations){#the-source-element:relevant-mutations-3}
    for `child`{.variable}.

The IDL attributes [`src`]{#dom-source-src .dfn
dfn-for="HTMLSourceElement" dfn-type="attribute"},
[`type`]{#dom-source-type .dfn dfn-for="HTMLSourceElement"
dfn-type="attribute"}, [`srcset`]{#dom-source-srcset .dfn
dfn-for="HTMLSourceElement" dfn-type="attribute"},
[`sizes`]{#dom-source-sizes .dfn dfn-for="HTMLSourceElement"
dfn-type="attribute"}, and [`media`]{#dom-source-media .dfn
dfn-for="HTMLSourceElement" dfn-type="attribute"} must
[reflect](common-dom-interfaces.html#reflect){#the-source-element:reflect}
the respective content attributes of the same name.

::: example
If the author isn\'t sure if user agents will all be able to render the
media resources provided, the author can listen to the
[`error`{#the-source-element:event-error}](indices.html#event-error)
event on the last
[`source`{#the-source-element:the-source-element-19}](#the-source-element)
element and trigger fallback behavior:

``` html
<script>
 function fallback(video) {
   // replace <video> with its contents
   while (video.hasChildNodes()) {
     if (video.firstChild instanceof HTMLSourceElement)
       video.removeChild(video.firstChild);
     else
       video.parentNode.insertBefore(video.firstChild, video);
   }
   video.parentNode.removeChild(video);
 }
</script>
<video controls autoplay>
 <source src='video.mp4' type='video/mp4; codecs="avc1.42E01E, mp4a.40.2"'>
 <source src='video.ogv' type='video/ogg; codecs="theora, vorbis"'
         onerror="fallback(parentNode)">
 ...
</video>
```
:::

#### [4.8.3]{.secno} The [`img`]{.dfn dfn-type="element"} element[](#the-img-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/img](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/img "The <img> HTML element embeds an image into the document.")

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
[HTMLImageElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLImageElement "The HTMLImageElement interface represents an HTML <img> element, providing the properties and methods used to manipulate image elements.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera8+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

[Categories](dom.html#concept-element-categories){#the-img-element:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-img-element:flow-content-2}.
:   [Phrasing
    content](dom.html#phrasing-content-2){#the-img-element:phrasing-content-2}.
:   [Embedded
    content](dom.html#embedded-content-category){#the-img-element:embedded-content-category}.
:   [Form-associated
    element](forms.html#form-associated-element){#the-img-element:form-associated-element}.
:   If the element has a
    [`usemap`{#the-img-element:attr-hyperlink-usemap}](image-maps.html#attr-hyperlink-usemap)
    attribute: [Interactive
    content](dom.html#interactive-content-2){#the-img-element:interactive-content-2}.
:   [Palpable
    content](dom.html#palpable-content-2){#the-img-element:palpable-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-img-element:concept-element-contexts}:
:   Where [embedded
    content](dom.html#embedded-content-category){#the-img-element:embedded-content-category-2}
    is expected.
:   As a child of a
    [`picture`{#the-img-element:the-picture-element}](#the-picture-element)
    element, after all
    [`source`{#the-img-element:the-source-element}](#the-source-element)
    elements.

[Content model](dom.html#concept-element-content-model){#the-img-element:concept-element-content-model}:
:   [Nothing](dom.html#concept-content-nothing){#the-img-element:concept-content-nothing}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-img-element:concept-element-tag-omission}:
:   No [end
    tag](syntax.html#syntax-end-tag){#the-img-element:syntax-end-tag}.

[Content attributes](dom.html#concept-element-attributes){#the-img-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-img-element:global-attributes}
:   [`alt`{#the-img-element:attr-img-alt}](#attr-img-alt) ---
    Replacement text for use when images are not available
:   [`src`{#the-img-element:attr-img-src}](#attr-img-src) --- Address of
    the resource
:   [`srcset`{#the-img-element:attr-img-srcset}](#attr-img-srcset) ---
    Images to use in different situations, e.g., high-resolution
    displays, small monitors, etc.
:   [`sizes`{#the-img-element:attr-img-sizes}](#attr-img-sizes) ---
    Image sizes for different page layouts
:   [`crossorigin`{#the-img-element:attr-img-crossorigin}](#attr-img-crossorigin)
    --- How the element handles crossorigin requests
:   [`usemap`{#the-img-element:attr-hyperlink-usemap-2}](image-maps.html#attr-hyperlink-usemap)
    --- Name of [image
    map](image-maps.html#image-map){#the-img-element:image-map} to use
:   [`ismap`{#the-img-element:attr-img-ismap}](#attr-img-ismap) ---
    Whether the image is a server-side image map
:   [`width`{#the-img-element:attr-dim-width}](embedded-content-other.html#attr-dim-width)
    --- Horizontal dimension
:   [`height`{#the-img-element:attr-dim-height}](embedded-content-other.html#attr-dim-height)
    --- Vertical dimension
:   [`referrerpolicy`{#the-img-element:attr-img-referrerpolicy}](#attr-img-referrerpolicy)
    --- [Referrer
    policy](https://w3c.github.io/webappsec-referrer-policy/#referrer-policy){#the-img-element:referrer-policy
    x-internal="referrer-policy"} for
    [fetches](https://fetch.spec.whatwg.org/#concept-fetch){#the-img-element:concept-fetch
    x-internal="concept-fetch"} initiated by the element
:   [`decoding`{#the-img-element:attr-img-decoding}](#attr-img-decoding)
    --- Decoding hint to use when processing this image for presentation
:   [`loading`{#the-img-element:attr-img-loading}](#attr-img-loading)
    --- Used when determining loading deferral
:   [`fetchpriority`{#the-img-element:attr-img-fetchpriority}](#attr-img-fetchpriority)
    --- Sets the
    [priority](https://fetch.spec.whatwg.org/#request-priority){#the-img-element:concept-request-priority
    x-internal="concept-request-priority"} for
    [fetches](https://fetch.spec.whatwg.org/#concept-fetch){#the-img-element:concept-fetch-2
    x-internal="concept-fetch"} initiated by the element

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-img-element:concept-element-accessibility-considerations}:
:   If the element has a non-empty
    [`alt`{#the-img-element:attr-img-alt-2}](#attr-img-alt) attribute:
    [for authors](https://w3c.github.io/html-aria/#el-img); [for
    implementers](https://w3c.github.io/html-aam/#el-img).
:   Otherwise: [for
    authors](https://w3c.github.io/html-aria/#el-img-empty-alt); [for
    implementers](https://w3c.github.io/html-aam/#el-img-empty-alt).

[DOM interface](dom.html#concept-element-dom){#the-img-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window,
     LegacyFactoryFunction=Image(optional unsigned long width, optional unsigned long height)]
    interface HTMLImageElement : HTMLElement {
      [HTMLConstructor] constructor();

      [CEReactions] attribute DOMString alt;
      [CEReactions] attribute USVString src;
      [CEReactions] attribute USVString srcset;
      [CEReactions] attribute DOMString sizes;
      [CEReactions] attribute DOMString? crossOrigin;
      [CEReactions] attribute DOMString useMap;
      [CEReactions] attribute boolean isMap;
      [CEReactions] attribute unsigned long width;
      [CEReactions] attribute unsigned long height;
      readonly attribute unsigned long naturalWidth;
      readonly attribute unsigned long naturalHeight;
      readonly attribute boolean complete;
      readonly attribute USVString currentSrc;
      [CEReactions] attribute DOMString referrerPolicy;
      [CEReactions] attribute DOMString decoding;
      [CEReactions] attribute DOMString loading;
      [CEReactions] attribute DOMString fetchPriority;

      Promise<undefined> decode();

      // also has obsolete members
    };
    ```

An [`img`{#the-img-element:the-img-element}](#the-img-element) element
represents an image.

An [`img`{#the-img-element:the-img-element-2}](#the-img-element) element
has a [dimension attribute
source]{#concept-img-dimension-attribute-source .dfn}, initially set to
the element itself.

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLImageElement/src](https://developer.mozilla.org/en-US/docs/Web/API/HTMLImageElement/src "The HTMLImageElement property src, which reflects the HTML src attribute, specifies the image to display in the <img> element.")

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
[Element/img#attr-srcset](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/img#attr-srcset "The <img> HTML element embeds an image into the document.")

Support in all current engines.

::: support
[Firefox38+]{.firefox .yes}[Safari8+]{.safari .yes}[Chrome34+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)≤18+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::::

The image given by the [`src`]{#attr-img-src .dfn dfn-for="img"
dfn-type="element-attr"} and [`srcset`]{#attr-img-srcset .dfn
dfn-for="img" dfn-type="element-attr"} attributes, and any previous
sibling
[`source`{#the-img-element:the-source-element-2}](#the-source-element)
elements\'
[`srcset`{#the-img-element:attr-source-srcset}](#attr-source-srcset)
attributes if the parent is a
[`picture`{#the-img-element:the-picture-element-2}](#the-picture-element)
element, is the embedded content; the value of the [`alt`]{#attr-img-alt
.dfn dfn-for="img" dfn-type="element-attr"} attribute provides
equivalent content for those who cannot process images or who have image
loading disabled (i.e., it is the
[`img`{#the-img-element:the-img-element-3}](#the-img-element) element\'s
[fallback
content](dom.html#fallback-content){#the-img-element:fallback-content}).

The requirements on the
[`alt`{#the-img-element:attr-img-alt-3}](#attr-img-alt) attribute\'s
value are described [in a separate section](images.html#alt).

At least one of the
[`src`{#the-img-element:attr-img-src-2}](#attr-img-src) and
[`srcset`{#the-img-element:attr-img-srcset-2}](#attr-img-srcset)
attributes must be present.

If the [`src`{#the-img-element:attr-img-src-3}](#attr-img-src) attribute
is present, it must contain a [valid non-empty URL potentially
surrounded by
spaces](urls-and-fetching.html#valid-non-empty-url-potentially-surrounded-by-spaces){#the-img-element:valid-non-empty-url-potentially-surrounded-by-spaces}
referencing a non-interactive, optionally animated, image resource that
is neither paged nor scripted.

The requirements above imply that images can be static bitmaps (e.g.
PNGs, GIFs, JPEGs), single-page vector documents (single-page PDFs, XML
files with an SVG document element), animated bitmaps (APNGs, animated
GIFs), animated vector graphics (XML files with an SVG [document
element](https://dom.spec.whatwg.org/#document-element){#the-img-element:document-element
x-internal="document-element"} that use declarative SMIL animation), and
so forth. However, these definitions preclude SVG files with script,
multipage PDF files, interactive MNG files, HTML documents, plain text
documents, and the like. [\[PNG\]](references.html#refsPNG)
[\[GIF\]](references.html#refsGIF) [\[JPEG\]](references.html#refsJPEG)
[\[PDF\]](references.html#refsPDF) [\[XML\]](references.html#refsXML)
[\[APNG\]](references.html#refsAPNG) [\[SVG\]](references.html#refsSVG)
[\[MNG\]](references.html#refsMNG)

The [`srcset`{#the-img-element:attr-img-srcset-3}](#attr-img-srcset)
attribute is a [srcset
attribute](images.html#srcset-attribute){#the-img-element:srcset-attribute}.

The [`srcset`{#the-img-element:attr-img-srcset-4}](#attr-img-srcset)
attribute and the
[`src`{#the-img-element:attr-img-src-4}](#attr-img-src) attribute (if
[width
descriptors](images.html#width-descriptor){#the-img-element:width-descriptor}
are not used) contribute the [image
sources](images.html#image-source){#the-img-element:image-source} to the
[source set](images.html#source-set){#the-img-element:source-set} (if no
[`source`{#the-img-element:the-source-element-3}](#the-source-element)
element was selected).

If the [`srcset`{#the-img-element:attr-img-srcset-5}](#attr-img-srcset)
attribute is present and has any [image candidate
strings](images.html#image-candidate-string){#the-img-element:image-candidate-string}
using a [width
descriptor](images.html#width-descriptor){#the-img-element:width-descriptor-2},
the [`sizes`{#the-img-element:attr-img-sizes-2}](#attr-img-sizes)
attribute must also be present. If the
[`srcset`{#the-img-element:attr-img-srcset-6}](#attr-img-srcset)
attribute is *not* specified, and the
[`loading`{#the-img-element:attr-img-loading-2}](#attr-img-loading)
attribute is in the
[Lazy](urls-and-fetching.html#attr-loading-lazy-state){#the-img-element:attr-loading-lazy-state}
state, the [`sizes`{#the-img-element:attr-img-sizes-3}](#attr-img-sizes)
attribute may be specified with the value \"`auto`\" ([ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#the-img-element:ascii-case-insensitive
x-internal="ascii-case-insensitive"}). The [`sizes`]{#attr-img-sizes
.dfn dfn-for="img" dfn-type="element-attr"} attribute is a [sizes
attribute](images.html#sizes-attribute){#the-img-element:sizes-attribute},
which contributes the [source
size](images.html#source-size-2){#the-img-element:source-size-2} to the
[source set](images.html#source-set){#the-img-element:source-set-2} (if
no
[`source`{#the-img-element:the-source-element-4}](#the-source-element)
element was selected).

An [`img`{#the-img-element:the-img-element-4}](#the-img-element) element
[allows auto-sizes]{#allows-auto-sizes .dfn} if:

- its
  [`loading`{#the-img-element:attr-img-loading-3}](#attr-img-loading)
  attribute is in the
  [Lazy](urls-and-fetching.html#attr-loading-lazy-state){#the-img-element:attr-loading-lazy-state-2}
  state, and
- its [`sizes`{#the-img-element:attr-img-sizes-4}](#attr-img-sizes)
  attribute\'s value is \"`auto`\" ([ASCII
  case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#the-img-element:ascii-case-insensitive-2
  x-internal="ascii-case-insensitive"}), or starts with \"`auto,`\"
  ([ASCII
  case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#the-img-element:ascii-case-insensitive-3
  x-internal="ascii-case-insensitive"}).

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Attributes/crossorigin](https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/crossorigin "The crossorigin attribute, valid on the <audio>, <img>, <link>, <script>, and <video> elements, provides support for CORS, defining how the element handles cross-origin requests, thereby enabling the configuration of the CORS requests for the element's fetched data. Depending on the element, the attribute can be a CORS settings attribute.")

Support in all current engines.

::: support
[Firefox8+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome13+]{.chrome
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

The [`crossorigin`]{#attr-img-crossorigin .dfn dfn-for="img"
dfn-type="element-attr"} attribute is a [CORS settings
attribute](urls-and-fetching.html#cors-settings-attribute){#the-img-element:cors-settings-attribute}.
Its purpose is to allow images from third-party sites that allow
cross-origin access to be used with
[`canvas`{#the-img-element:the-canvas-element}](canvas.html#the-canvas-element).

The [`referrerpolicy`]{#attr-img-referrerpolicy .dfn dfn-for="img"
dfn-type="element-attr"} attribute is a [referrer policy
attribute](urls-and-fetching.html#referrer-policy-attribute){#the-img-element:referrer-policy-attribute}.
Its purpose is to set the [referrer
policy](https://w3c.github.io/webappsec-referrer-policy/#referrer-policy){#the-img-element:referrer-policy-2
x-internal="referrer-policy"} used when
[fetching](https://fetch.spec.whatwg.org/#concept-fetch){#the-img-element:concept-fetch-3
x-internal="concept-fetch"} the image.
[\[REFERRERPOLICY\]](references.html#refsREFERRERPOLICY)

The [`decoding`]{#attr-img-decoding .dfn dfn-for="img"
dfn-type="element-attr"} attribute indicates the preferred method to
[decode](images.html#img-decoding-process){#the-img-element:img-decoding-process}
this image. The attribute, if present, must be an [image decoding
hint](images.html#image-decoding-hint){#the-img-element:image-decoding-hint}.
This attribute\'s *[missing value
default](common-microsyntaxes.html#missing-value-default)* and *[invalid
value default](common-microsyntaxes.html#invalid-value-default)* are
both the
[Auto](images.html#attr-img-decoding-auto-state){#the-img-element:attr-img-decoding-auto-state}
state.

::::: {.mdn-anno .wrapped .before}
MDN

:::: feature
[HTMLImageElement/fetchPriority](https://developer.mozilla.org/en-US/docs/Web/API/HTMLImageElement/fetchPriority "The fetchPriority property of the HTMLImageElement interface represents a hint given to the browser on how it should prioritize the fetch of the image relative to other images.")

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

The [`fetchpriority`]{#attr-img-fetchpriority .dfn dfn-for="img"
dfn-type="element-attr"} attribute is a [fetch priority
attribute](urls-and-fetching.html#fetch-priority-attribute){#the-img-element:fetch-priority-attribute}.
Its purpose is to set the
[priority](https://fetch.spec.whatwg.org/#request-priority){#the-img-element:concept-request-priority-2
x-internal="concept-request-priority"} used when
[fetching](https://fetch.spec.whatwg.org/#concept-fetch){#the-img-element:concept-fetch-4
x-internal="concept-fetch"} the image.

The [`loading`]{#attr-img-loading .dfn dfn-for="img"
dfn-type="element-attr"} attribute is a [lazy loading
attribute](urls-and-fetching.html#lazy-loading-attribute){#the-img-element:lazy-loading-attribute}.
Its purpose is to indicate the policy for loading images that are
outside the viewport.

When the
[`loading`{#the-img-element:attr-img-loading-4}](#attr-img-loading)
attribute\'s state is changed to the
[Eager](urls-and-fetching.html#attr-loading-eager-state){#the-img-element:attr-loading-eager-state}
state, the user agent must run these steps:

1.  Let `resumptionSteps`{.variable} be the
    [`img`{#the-img-element:the-img-element-5}](#the-img-element)
    element\'s [lazy load resumption
    steps](urls-and-fetching.html#lazy-load-resumption-steps){#the-img-element:lazy-load-resumption-steps}.

2.  If `resumptionSteps`{.variable} is null, then return.

3.  Set the
    [`img`{#the-img-element:the-img-element-6}](#the-img-element)\'s
    [lazy load resumption
    steps](urls-and-fetching.html#lazy-load-resumption-steps){#the-img-element:lazy-load-resumption-steps-2}
    to null.

4.  Invoke `resumptionSteps`{.variable}.

::: example
``` html
<img src="1.jpeg" alt="1">
<img src="2.jpeg" loading=eager alt="2">
<img src="3.jpeg" loading=lazy alt="3">
<div id=very-large></div> <!-- Everything after this div is below the viewport -->
<img src="4.jpeg" alt="4">
<img src="5.jpeg" loading=lazy alt="5">
```

In the example above, the images load as follows:

`1.jpeg`, `2.jpeg`, `4.jpeg`
:   The images load eagerly and delay the window\'s load event.

`3.jpeg`
:   The image loads when layout is known, due to being in the viewport,
    however it does not delay the window\'s load event.

`5.jpeg`

:   The image loads only once scrolled into the viewport, and does not
    delay the window\'s load event.

Developers are encouraged to specify a preferred aspect ratio via
[`width`{#the-img-element:attr-dim-width-2}](embedded-content-other.html#attr-dim-width)
and
[`height`{#the-img-element:attr-dim-height-2}](embedded-content-other.html#attr-dim-height)
attributes on lazy loaded images, even if CSS sets the image\'s width
and height properties, to prevent the page layout from shifting around
after the image loads.
:::

The [`img`{#the-img-element:the-img-element-7}](#the-img-element) [HTML
element insertion
steps](infrastructure.html#html-element-insertion-steps){#the-img-element:html-element-insertion-steps},
given `insertedNode`{.variable}, are:

1.  If `insertedNode`{.variable}\'s parent is a
    [`picture`{#the-img-element:the-picture-element-3}](#the-picture-element)
    element, then, count this as a [relevant
    mutation](images.html#relevant-mutations){#the-img-element:relevant-mutations}
    for `insertedNode`{.variable}.

The [`img`{#the-img-element:the-img-element-8}](#the-img-element) [HTML
element moving
steps](infrastructure.html#html-element-moving-steps){#the-img-element:html-element-moving-steps},
given `movedNode`{.variable} and `oldParent`{.variable}, are:

1.  If `oldParent`{.variable} is a
    [`picture`{#the-img-element:the-picture-element-4}](#the-picture-element)
    element, then, count this as a [relevant
    mutation](images.html#relevant-mutations){#the-img-element:relevant-mutations-2}
    for `movedNode`{.variable}.

The [`img`{#the-img-element:the-img-element-9}](#the-img-element) [HTML
element removing
steps](infrastructure.html#html-element-removing-steps){#the-img-element:html-element-removing-steps},
given `removedNode`{.variable} and `oldParent`{.variable}, are:

1.  If `oldParent`{.variable} is a
    [`picture`{#the-img-element:the-picture-element-5}](#the-picture-element)
    element, then, count this as a [relevant
    mutation](images.html#relevant-mutations){#the-img-element:relevant-mutations-3}
    for `removedNode`{.variable}.

------------------------------------------------------------------------

The [`img`{#the-img-element:the-img-element-10}](#the-img-element)
element must not be used as a layout tool. In particular,
[`img`{#the-img-element:the-img-element-11}](#the-img-element) elements
should not be used to display transparent images, as such images rarely
convey meaning and rarely add anything useful to the document.

------------------------------------------------------------------------

What an [`img`{#the-img-element:the-img-element-12}](#the-img-element)
element represents depends on the
[`src`{#the-img-element:attr-img-src-5}](#attr-img-src) attribute and
the [`alt`{#the-img-element:attr-img-alt-4}](#attr-img-alt) attribute.

If the [`src`{#the-img-element:attr-img-src-6}](#attr-img-src) attribute is set and the [`alt`{#the-img-element:attr-img-alt-5}](#attr-img-alt) attribute is set to the empty string

:   The image is either decorative or supplemental to the rest of the
    content, redundant with some other information in the document.

    If the image is
    [available](images.html#img-available){#the-img-element:img-available}
    and the user agent is configured to display that image, then the
    element
    [represents](dom.html#represents){#the-img-element:represents} the
    element\'s image data.

    Otherwise, the element
    [represents](dom.html#represents){#the-img-element:represents-2}
    nothing, and may be omitted completely from the rendering. User
    agents may provide the user with a notification that an image is
    present but has been omitted from the rendering.

If the [`src`{#the-img-element:attr-img-src-7}](#attr-img-src) attribute is set and the [`alt`{#the-img-element:attr-img-alt-6}](#attr-img-alt) attribute is set to a value that isn\'t empty

:   The image is a key part of the content; the
    [`alt`{#the-img-element:attr-img-alt-7}](#attr-img-alt) attribute
    gives a textual equivalent or replacement for the image.

    If the image is
    [available](images.html#img-available){#the-img-element:img-available-2}
    and the user agent is configured to display that image, then the
    element
    [represents](dom.html#represents){#the-img-element:represents-3} the
    element\'s image data.

    Otherwise, the element
    [represents](dom.html#represents){#the-img-element:represents-4} the
    text given by the
    [`alt`{#the-img-element:attr-img-alt-8}](#attr-img-alt) attribute.
    User agents may provide the user with a notification that an image
    is present but has been omitted from the rendering.

If the [`src`{#the-img-element:attr-img-src-8}](#attr-img-src) attribute is set and the [`alt`{#the-img-element:attr-img-alt-9}](#attr-img-alt) attribute is not

:   The image might be a key part of the content, and there is no
    textual equivalent of the image available.

    In a conforming document, the absence of the
    [`alt`{#the-img-element:attr-img-alt-10}](#attr-img-alt) attribute
    indicates that the image is a key part of the content but that a
    textual replacement for the image was not available when the image
    was generated.

    If the image is
    [available](images.html#img-available){#the-img-element:img-available-3}
    and the user agent is configured to display that image, then the
    element
    [represents](dom.html#represents){#the-img-element:represents-5} the
    element\'s image data.

    If the image has a
    [`src`{#the-img-element:attr-img-src-9}](#attr-img-src) attribute
    whose value is the empty string, then the element
    [represents](dom.html#represents){#the-img-element:represents-6}
    nothing.

    Otherwise, the user agent should display some sort of indicator that
    there is an image that is not being rendered, and may, if requested
    by the user, or if so configured, or when required to provide
    contextual information in response to navigation, provide caption
    information for the image, derived as follows:

    1.  If the image has a
        [`title`{#the-img-element:attr-title}](dom.html#attr-title)
        attribute whose value is not the empty string, then return the
        value of that attribute.

    2.  If the image is a descendant of a
        [`figure`{#the-img-element:the-figure-element}](grouping-content.html#the-figure-element)
        element that has a child
        [`figcaption`{#the-img-element:the-figcaption-element}](grouping-content.html#the-figcaption-element)
        element, and, ignoring the
        [`figcaption`{#the-img-element:the-figcaption-element-2}](grouping-content.html#the-figcaption-element)
        element and its descendants, the
        [`figure`{#the-img-element:the-figure-element-2}](grouping-content.html#the-figure-element)
        element has no [flow
        content](dom.html#flow-content-2){#the-img-element:flow-content-2-2}
        descendants other than [inter-element
        whitespace](dom.html#inter-element-whitespace){#the-img-element:inter-element-whitespace}
        and the
        [`img`{#the-img-element:the-img-element-13}](#the-img-element)
        element, then return the contents of the first such
        [`figcaption`{#the-img-element:the-figcaption-element-3}](grouping-content.html#the-figcaption-element)
        element.

    3.  Return nothing. (There is no caption information.)

If the [`src`{#the-img-element:attr-img-src-10}](#attr-img-src) attribute is not set and either the [`alt`{#the-img-element:attr-img-alt-11}](#attr-img-alt) attribute is set to the empty string or the [`alt`{#the-img-element:attr-img-alt-12}](#attr-img-alt) attribute is not set at all

:   The element
    [represents](dom.html#represents){#the-img-element:represents-7}
    nothing.

Otherwise

:   The element
    [represents](dom.html#represents){#the-img-element:represents-8} the
    text given by the
    [`alt`{#the-img-element:attr-img-alt-13}](#attr-img-alt) attribute.

The [`alt`{#the-img-element:attr-img-alt-14}](#attr-img-alt) attribute
does not represent advisory information. User agents must not present
the contents of the
[`alt`{#the-img-element:attr-img-alt-15}](#attr-img-alt) attribute in
the same way as content of the
[`title`{#the-img-element:attr-title-2}](dom.html#attr-title) attribute.

User agents may always provide the user with the option to display any
image, or to prevent any image from being displayed. User agents may
also apply heuristics to help the user make use of the image when the
user is unable to see it, e.g. due to a visual disability or because
they are using a text terminal with no graphics capabilities. Such
heuristics could include, for instance, optical character recognition
(OCR) of text found within the image.

While user agents are encouraged to repair cases of missing
[`alt`{#the-img-element:attr-img-alt-16}](#attr-img-alt) attributes,
authors must not rely on such behavior. [Requirements for providing text
to act as an alternative for images](images.html#alt) are described in
detail below.

The *contents* of
[`img`{#the-img-element:the-img-element-14}](#the-img-element) elements,
if any, are ignored for the purposes of rendering.

------------------------------------------------------------------------

The
[`usemap`{#the-img-element:attr-hyperlink-usemap-3}](image-maps.html#attr-hyperlink-usemap)
attribute, if present, can indicate that the image has an associated
[image map](image-maps.html#image-map){#the-img-element:image-map-2}.

The [`ismap`]{#attr-img-ismap .dfn dfn-for="img"
dfn-type="element-attr"} attribute, when used on an element that is a
descendant of an
[`a`{#the-img-element:the-a-element}](text-level-semantics.html#the-a-element)
element with an
[`href`{#the-img-element:attr-hyperlink-href}](links.html#attr-hyperlink-href)
attribute, indicates by its presence that the element provides access to
a server-side image map. This affects how events are handled on the
corresponding
[`a`{#the-img-element:the-a-element-2}](text-level-semantics.html#the-a-element)
element.

The [`ismap`{#the-img-element:attr-img-ismap-2}](#attr-img-ismap)
attribute is a [boolean
attribute](common-microsyntaxes.html#boolean-attribute){#the-img-element:boolean-attribute}.
The attribute must not be specified on an element that does not have an
ancestor
[`a`{#the-img-element:the-a-element-3}](text-level-semantics.html#the-a-element)
element with an
[`href`{#the-img-element:attr-hyperlink-href-2}](links.html#attr-hyperlink-href)
attribute.

The
[`usemap`{#the-img-element:attr-hyperlink-usemap-4}](image-maps.html#attr-hyperlink-usemap)
and [`ismap`{#the-img-element:attr-img-ismap-3}](#attr-img-ismap)
attributes can result in confusing behavior when used together with
[`source`{#the-img-element:the-source-element-5}](#the-source-element)
elements with the
[`media`{#the-img-element:attr-source-media}](#attr-source-media)
attribute specified in a
[`picture`{#the-img-element:the-picture-element-6}](#the-picture-element)
element.

The [`img`{#the-img-element:the-img-element-15}](#the-img-element)
element supports [dimension
attributes](embedded-content-other.html#dimension-attributes){#the-img-element:dimension-attributes}.

::::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLImageElement/alt](https://developer.mozilla.org/en-US/docs/Web/API/HTMLImageElement/alt "The HTMLImageElement property alt provides fallback (alternate) text to display when the image specified by the <img> element is not loaded.")

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
[HTMLImageElement/srcset](https://developer.mozilla.org/en-US/docs/Web/API/HTMLImageElement/srcset "The HTMLImageElement property srcset is a string which identifies one or more image candidate strings, separated using commas (,) each specifying image resources to use under given circumstances.")

Support in all current engines.

::: support
[Firefox38+]{.firefox .yes}[Safari8+]{.safari .yes}[Chrome34+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

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

:::: feature
[HTMLImageElement/sizes](https://developer.mozilla.org/en-US/docs/Web/API/HTMLImageElement/sizes "The HTMLImageElement property sizes allows you to specify the layout width of the image for each of a list of media conditions. This provides the ability to automatically select among different images—even images of different orientations or aspect ratios—as the document state changes to match different media conditions.")

Support in all current engines.

::: support
[Firefox38+]{.firefox .yes}[Safari9.1+]{.safari .yes}[Chrome38+]{.chrome
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
:::::::::

The [`alt`]{#dom-img-alt .dfn dfn-for="HTMLImageElement"
dfn-type="attribute"}, [`src`]{#dom-img-src .dfn
dfn-for="HTMLImageElement" dfn-type="attribute"},
[`srcset`]{#dom-img-srcset .dfn dfn-for="HTMLImageElement"
dfn-type="attribute"}, and [`sizes`]{#dom-img-sizes .dfn
dfn-for="HTMLImageElement" dfn-type="attribute"} IDL attributes must
[reflect](common-dom-interfaces.html#reflect){#the-img-element:reflect}
the respective content attributes of the same name.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLImageElement/crossOrigin](https://developer.mozilla.org/en-US/docs/Web/API/HTMLImageElement/crossOrigin "The HTMLImageElement interface's crossOrigin attribute is a string which specifies the Cross-Origin Resource Sharing (CORS) setting to use when retrieving the image.")

Support in all current engines.

::: support
[Firefox8+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome13+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer11]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The [`crossOrigin`]{#dom-img-crossorigin .dfn dfn-for="HTMLImageElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-img-element:reflect-2}
the
[`crossorigin`{#the-img-element:attr-img-crossorigin-2}](#attr-img-crossorigin)
content attribute, [limited to only known
values](common-dom-interfaces.html#limited-to-only-known-values){#the-img-element:limited-to-only-known-values}.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLImageElement/useMap](https://developer.mozilla.org/en-US/docs/Web/API/HTMLImageElement/useMap "The useMap property on the HTMLImageElement interface reflects the value of the HTML usemap attribute, which is a string providing the name of the client-side image map to apply to the image.")

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

The [`useMap`]{#dom-img-usemap .dfn dfn-for="HTMLImageElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-img-element:reflect-3}
the
[`usemap`{#the-img-element:attr-hyperlink-usemap-5}](image-maps.html#attr-hyperlink-usemap)
content attribute.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLImageElement/isMap](https://developer.mozilla.org/en-US/docs/Web/API/HTMLImageElement/isMap "The HTMLImageElement property isMap is a Boolean value which indicates that the image is to be used by a server-side image map. This may only be used on images located within an <a> element.")

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

The [`isMap`]{#dom-img-ismap .dfn dfn-for="HTMLImageElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-img-element:reflect-4}
the [`ismap`{#the-img-element:attr-img-ismap-4}](#attr-img-ismap)
content attribute.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLImageElement/referrerPolicy](https://developer.mozilla.org/en-US/docs/Web/API/HTMLImageElement/referrerPolicy "The HTMLImageElement.referrerPolicy property reflects the HTML referrerpolicy attribute of the <img> element defining which referrer is sent when fetching the resource.")

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

The [`referrerPolicy`]{#dom-img-referrerpolicy .dfn
dfn-for="HTMLImageElement" dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-img-element:reflect-5}
the
[`referrerpolicy`{#the-img-element:attr-img-referrerpolicy-2}](#attr-img-referrerpolicy)
content attribute, [limited to only known
values](common-dom-interfaces.html#limited-to-only-known-values){#the-img-element:limited-to-only-known-values-2}.

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLImageElement/decoding](https://developer.mozilla.org/en-US/docs/Web/API/HTMLImageElement/decoding "The decoding property of the HTMLImageElement interface represents a hint given to the browser on how it should decode the image.")

Support in all current engines.

::: support
[Firefox63+]{.firefox .yes}[Safari11.1+]{.safari
.yes}[Chrome65+]{.chrome .yes}

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

:::: feature
[SVGImageElement/decoding](https://developer.mozilla.org/en-US/docs/Web/API/SVGImageElement/decoding "The decoding property of the SVGImageElement interface represents a hint given to the browser on how it should decode the image.")

::: support
[Firefox63+]{.firefox .yes}[SafariNo]{.safari .no}[Chrome65+]{.chrome
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
:::::::

The [`decoding`]{#dom-img-decoding .dfn dfn-for="HTMLImageElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-img-element:reflect-6}
the
[`decoding`{#the-img-element:attr-img-decoding-2}](#attr-img-decoding)
content attribute, [limited to only known
values](common-dom-interfaces.html#limited-to-only-known-values){#the-img-element:limited-to-only-known-values-3}.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLImageElement/loading](https://developer.mozilla.org/en-US/docs/Web/API/HTMLImageElement/loading "The HTMLImageElement property loading is a string whose value provides a hint to the user agent on how to handle the loading of the image which is currently outside the window's visual viewport.")

Support in all current engines.

::: support
[Firefox75+]{.firefox .yes}[Safari15.4+]{.safari
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

The [`loading`]{#dom-img-loading .dfn dfn-for="HTMLImageElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-img-element:reflect-7}
the [`loading`{#the-img-element:attr-img-loading-5}](#attr-img-loading)
content attribute, [limited to only known
values](common-dom-interfaces.html#limited-to-only-known-values){#the-img-element:limited-to-only-known-values-4}.

The [`fetchPriority`]{#dom-img-fetchpriority .dfn
dfn-for="HTMLImageElement" dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-img-element:reflect-8}
the
[`fetchpriority`{#the-img-element:attr-img-fetchpriority-2}](#attr-img-fetchpriority)
content attribute, [limited to only known
values](common-dom-interfaces.html#limited-to-only-known-values){#the-img-element:limited-to-only-known-values-5}.

`image`{.variable}`.`[`width`](#dom-img-width){#dom-img-width-dev}` [ = ``value`{.variable}` ]`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLImageElement/width](https://developer.mozilla.org/en-US/docs/Web/API/HTMLImageElement/width "The width property of the HTMLImageElement interface indicates the width at which an image is drawn in CSS pixels if it's being drawn or rendered to any visual medium such as a screen or printer. Otherwise, it's the natural, pixel density-corrected width of the image.")

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

`image`{.variable}`.`[`height`](#dom-img-height){#dom-img-height-dev}` [ = ``value`{.variable}` ]`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLImageElement/height](https://developer.mozilla.org/en-US/docs/Web/API/HTMLImageElement/height "The height property of the HTMLImageElement interface indicates the height at which the image is drawn, in CSS pixels if the image is being drawn or rendered to any visual medium such as the screen or a printer; otherwise, it's the natural, pixel density corrected height of the image.")

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

These attributes return the actual rendered dimensions of the image, or
0 if the dimensions are not known.

They can be set, to change the corresponding content attributes.

`image`{.variable}`.`[`naturalWidth`](#dom-img-naturalwidth){#dom-img-naturalwidth-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLImageElement/naturalWidth](https://developer.mozilla.org/en-US/docs/Web/API/HTMLImageElement/naturalWidth "The HTMLImageElement interface's read-only naturalWidth property returns the intrinsic (natural), density-corrected width of the image in CSS pixels.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
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

`image`{.variable}`.`[`naturalHeight`](#dom-img-naturalheight){#dom-img-naturalheight-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLImageElement/naturalHeight](https://developer.mozilla.org/en-US/docs/Web/API/HTMLImageElement/naturalHeight "The HTMLImageElement interface's naturalHeight property is a read-only value which returns the intrinsic (natural), density-corrected height of the image in CSS pixels.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
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

These attributes return the natural dimensions of the image, or 0 if the
dimensions are not known.

`image`{.variable}`.`[`complete`](#dom-img-complete){#dom-img-complete-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLImageElement/complete](https://developer.mozilla.org/en-US/docs/Web/API/HTMLImageElement/complete "The read-only HTMLImageElement interface's complete attribute is a Boolean value which indicates whether or not the image has completely loaded.")

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

Returns true if the image has been completely downloaded or if no image
is specified; otherwise, returns false.

`image`{.variable}`.`[`currentSrc`](#dom-img-currentsrc){#dom-img-currentsrc-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLImageElement/currentSrc](https://developer.mozilla.org/en-US/docs/Web/API/HTMLImageElement/currentSrc "The read-only HTMLImageElement property currentSrc indicates the URL of the image which is currently presented in the <img> element it represents.")

Support in all current engines.

::: support
[Firefox38+]{.firefox .yes}[Safari9.1+]{.safari .yes}[Chrome38+]{.chrome
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

Returns the image\'s [absolute
URL](https://url.spec.whatwg.org/#syntax-url-absolute){#the-img-element:absolute-url
x-internal="absolute-url"}.

`image`{.variable}`.`[`decode`](#dom-img-decode){#dom-img-decode-dev}`()`

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLImageElement/decode](https://developer.mozilla.org/en-US/docs/Web/API/HTMLImageElement/decode "The decode() method of the HTMLImageElement interface returns a Promise that resolves when the image is decoded and it is safe to append the image to the DOM.")

Support in all current engines.

::: support
[Firefox68+]{.firefox .yes}[Safari11.1+]{.safari
.yes}[Chrome64+]{.chrome .yes}

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

:::: feature
[SVGImageElement/decode](https://developer.mozilla.org/en-US/docs/Web/API/SVGImageElement/decode "The decode() method of the SVGImageElement interface initiates asynchronous decoding of an image, returning a Promise that resolves once the image data is ready for use.")

::: support
[Firefox68+]{.firefox .yes}[SafariNo]{.safari .no}[Chrome64+]{.chrome
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
:::::::

This method causes the user agent to
[decode](images.html#img-decoding-process){#the-img-element:img-decoding-process-2}
the image [in
parallel](infrastructure.html#in-parallel){#the-img-element:in-parallel},
returning a promise that fulfills when decoding is complete.

The promise will be rejected with an
[\"`EncodingError`\"](https://webidl.spec.whatwg.org/#encodingerror){#the-img-element:encodingerror
x-internal="encodingerror"}
[`DOMException`{#the-img-element:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if the image cannot be decoded.

`image`{.variable}` = new `[`Image`](#dom-image){#dom-image-dev}`([ ``width`{.variable}` [, ``height`{.variable}` ] ])`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLImageElement/Image](https://developer.mozilla.org/en-US/docs/Web/API/HTMLImageElement/Image "The Image() constructor creates a new HTMLImageElement instance. It is functionally equivalent to document.createElement('img').")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera8+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

Returns a new
[`img`{#the-img-element:the-img-element-16}](#the-img-element) element,
with the
[`width`{#the-img-element:attr-dim-width-3}](embedded-content-other.html#attr-dim-width)
and
[`height`{#the-img-element:attr-dim-height-3}](embedded-content-other.html#attr-dim-height)
attributes set to the values passed in the relevant arguments, if
applicable.

The IDL attributes [`width`]{#dom-img-width .dfn
dfn-for="HTMLImageElement" dfn-type="attribute"} and
[`height`]{#dom-img-height .dfn dfn-for="HTMLImageElement"
dfn-type="attribute"} must return the rendered width and height of the
image, in [CSS
pixels](https://drafts.csswg.org/css-values/#px){#the-img-element:'px'
x-internal="'px'"}, if the image is [being
rendered](rendering.html#being-rendered){#the-img-element:being-rendered};
or else the [density-corrected natural width and
height](images.html#density-corrected-intrinsic-width-and-height){#the-img-element:density-corrected-intrinsic-width-and-height}
of the image, in [CSS
pixels](https://drafts.csswg.org/css-values/#px){#the-img-element:'px'-2
x-internal="'px'"}, if the image has [density-corrected natural width
and
height](images.html#density-corrected-intrinsic-width-and-height){#the-img-element:density-corrected-intrinsic-width-and-height-2}
and is *[available](images.html#img-available)* but is not [being
rendered](rendering.html#being-rendered){#the-img-element:being-rendered-2};
or else 0, if the image is not *[available](images.html#img-available)*
or does not have [density-corrected natural width and
height](images.html#density-corrected-intrinsic-width-and-height){#the-img-element:density-corrected-intrinsic-width-and-height-3}.
[\[CSS\]](references.html#refsCSS)

On setting, they must act as if they
[reflected](common-dom-interfaces.html#reflect){#the-img-element:reflect-9}
the respective content attributes of the same name.

The IDL attributes [`naturalWidth`]{#dom-img-naturalwidth .dfn
dfn-for="HTMLImageElement" dfn-type="attribute"} and
[`naturalHeight`]{#dom-img-naturalheight .dfn dfn-for="HTMLImageElement"
dfn-type="attribute"} must return the [density-corrected natural width
and
height](images.html#density-corrected-intrinsic-width-and-height){#the-img-element:density-corrected-intrinsic-width-and-height-4}
of the image, in [CSS
pixels](https://drafts.csswg.org/css-values/#px){#the-img-element:'px'-3
x-internal="'px'"}, if the image has [density-corrected natural width
and
height](images.html#density-corrected-intrinsic-width-and-height){#the-img-element:density-corrected-intrinsic-width-and-height-5}
and is *[available](images.html#img-available)*, or else 0.
[\[CSS\]](references.html#refsCSS)

Since the [density-corrected natural width and
height](images.html#density-corrected-intrinsic-width-and-height){#the-img-element:density-corrected-intrinsic-width-and-height-6}
of an image take into account any orientation specified in its metadata,
[`naturalWidth`{#the-img-element:dom-img-naturalwidth-2}](#dom-img-naturalwidth)
and
[`naturalHeight`{#the-img-element:dom-img-naturalheight-2}](#dom-img-naturalheight)
reflect the dimensions after applying any rotation needed to correctly
orient the image, regardless of the value of the
[\'image-orientation\'](https://drafts.csswg.org/css-images-3/#the-image-orientation){#the-img-element:'image-orientation'
x-internal="'image-orientation'"} property.

The [`complete`]{#dom-img-complete .dfn dfn-for="HTMLImageElement"
dfn-type="attribute"} getter steps are:

1.  If any of the following are true:

    - both the [`src`{#the-img-element:attr-img-src-11}](#attr-img-src)
      attribute and the
      [`srcset`{#the-img-element:attr-img-srcset-7}](#attr-img-srcset)
      attribute are omitted;

    - the
      [`srcset`{#the-img-element:attr-img-srcset-8}](#attr-img-srcset)
      attribute is omitted and the
      [`src`{#the-img-element:attr-img-src-12}](#attr-img-src)
      attribute\'s value is the empty string;

    - the [`img`{#the-img-element:the-img-element-17}](#the-img-element)
      element\'s [current
      request](images.html#current-request){#the-img-element:current-request}\'s
      [state](images.html#img-req-state){#the-img-element:img-req-state}
      is [completely
      available](images.html#img-all){#the-img-element:img-all} and its
      [pending
      request](images.html#pending-request){#the-img-element:pending-request}
      is null; or

    - the [`img`{#the-img-element:the-img-element-18}](#the-img-element)
      element\'s [current
      request](images.html#current-request){#the-img-element:current-request-2}\'s
      [state](images.html#img-req-state){#the-img-element:img-req-state-2}
      is [broken](images.html#img-error){#the-img-element:img-error} and
      its [pending
      request](images.html#pending-request){#the-img-element:pending-request-2}
      is null,

    then return true.

2.  Return false.

The [`currentSrc`]{#dom-img-currentsrc .dfn dfn-for="HTMLImageElement"
dfn-type="attribute"} IDL attribute must return the
[`img`{#the-img-element:the-img-element-19}](#the-img-element)
element\'s [current
request](images.html#current-request){#the-img-element:current-request-3}\'s
[current URL](images.html#img-req-url){#the-img-element:img-req-url}.

The [`decode()`]{#dom-img-decode .dfn dfn-for="HTMLImageElement"
dfn-type="method"} method, when invoked, must perform the following
steps:

1.  Let `promise`{.variable} be a new promise.

2.  [Queue a
    microtask](webappapis.html#queue-a-microtask){#the-img-element:queue-a-microtask}
    to perform the following steps:

    ::: note
    This is done because [updating the image
    data](images.html#update-the-image-data){#the-img-element:update-the-image-data}
    takes place in a microtask as well. Thus, to make code such as

    ``` js
    img.src = "stars.jpg";
    img.decode();
    ```

    properly decode `stars.jpg`, we need to delay any processing by one
    microtask.
    :::

    1.  Let `global`{.variable} be
        [this](https://webidl.spec.whatwg.org/#this){#the-img-element:this
        x-internal="this"}\'s [relevant global
        object](webappapis.html#concept-relevant-global){#the-img-element:concept-relevant-global}.

    2.  If any of the following are true:

        - [this](https://webidl.spec.whatwg.org/#this){#the-img-element:this-2
          x-internal="this"}\'s [node
          document](https://dom.spec.whatwg.org/#concept-node-document){#the-img-element:node-document
          x-internal="node-document"} is not [fully
          active](document-sequences.html#fully-active){#the-img-element:fully-active};
          or

        - [this](https://webidl.spec.whatwg.org/#this){#the-img-element:this-3
          x-internal="this"}\'s [current
          request](images.html#current-request){#the-img-element:current-request-4}\'s
          [state](images.html#img-req-state){#the-img-element:img-req-state-3}
          is
          [broken](images.html#img-error){#the-img-element:img-error-2},

        then reject `promise`{.variable} with an
        [\"`EncodingError`\"](https://webidl.spec.whatwg.org/#encodingerror){#the-img-element:encodingerror-2
        x-internal="encodingerror"}
        [`DOMException`{#the-img-element:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    3.  Otherwise, [in
        parallel](infrastructure.html#in-parallel){#the-img-element:in-parallel-2},
        wait for one of the following cases to occur, and perform the
        corresponding actions:

        This [`img`{#the-img-element:the-img-element-20}](#the-img-element) element\'s [node document](https://dom.spec.whatwg.org/#concept-node-document){#the-img-element:node-document-2 x-internal="node-document"} stops being [fully active](document-sequences.html#fully-active){#the-img-element:fully-active-2}\
        This [`img`{#the-img-element:the-img-element-21}](#the-img-element) element\'s [current request](images.html#current-request){#the-img-element:current-request-5} changes or is mutated\
        This [`img`{#the-img-element:the-img-element-22}](#the-img-element) element\'s [current request](images.html#current-request){#the-img-element:current-request-6}\'s [state](images.html#img-req-state){#the-img-element:img-req-state-4} becomes [broken](images.html#img-error){#the-img-element:img-error-3}
        :   [Queue a global
            task](webappapis.html#queue-a-global-task){#the-img-element:queue-a-global-task}
            on the [DOM manipulation task
            source](webappapis.html#dom-manipulation-task-source){#the-img-element:dom-manipulation-task-source}
            with `global`{.variable} to reject `promise`{.variable} with
            an
            [\"`EncodingError`\"](https://webidl.spec.whatwg.org/#encodingerror){#the-img-element:encodingerror-3
            x-internal="encodingerror"}
            [`DOMException`{#the-img-element:domexception-3}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

        This [`img`{#the-img-element:the-img-element-23}](#the-img-element) element\'s [current request](images.html#current-request){#the-img-element:current-request-7}\'s [state](images.html#img-req-state){#the-img-element:img-req-state-5} becomes [completely available](images.html#img-all){#the-img-element:img-all-2}

        :   [Decode](images.html#img-decoding-process){#the-img-element:img-decoding-process-3}
            the image.

            If decoding does not need to be performed for this image
            (for example because it is a vector graphic) or the decoding
            process completes successfully, then [queue a global
            task](webappapis.html#queue-a-global-task){#the-img-element:queue-a-global-task-2}
            on the [DOM manipulation task
            source](webappapis.html#dom-manipulation-task-source){#the-img-element:dom-manipulation-task-source-2}
            with `global`{.variable} to resolve `promise`{.variable}
            with undefined.

            If decoding fails (for example due to invalid image data),
            then [queue a global
            task](webappapis.html#queue-a-global-task){#the-img-element:queue-a-global-task-3}
            on the [DOM manipulation task
            source](webappapis.html#dom-manipulation-task-source){#the-img-element:dom-manipulation-task-source-3}
            with `global`{.variable} to reject `promise`{.variable} with
            an
            [\"`EncodingError`\"](https://webidl.spec.whatwg.org/#encodingerror){#the-img-element:encodingerror-4
            x-internal="encodingerror"}
            [`DOMException`{#the-img-element:domexception-4}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

            User agents should ensure that the decoded media data stays
            readily available until at least the end of the next
            successful [update the
            rendering](webappapis.html#update-the-rendering){#the-img-element:update-the-rendering}
            step in the [event
            loop](webappapis.html#event-loop){#the-img-element:event-loop}.
            This is an important part of the API contract, and should
            not be broken if at all possible. (Typically, this would
            only be violated in low-memory situations that require
            evicting decoded image data, or when the image is too large
            to keep in decoded form for this period of time.)

        Animated images will become [completely
        available](images.html#img-all){#the-img-element:img-all-3} only
        after all their frames are loaded. Thus, even though an
        implementation could decode the first frame before that point,
        the above steps will not do so, instead waiting until all frames
        are available.

3.  Return `promise`{.variable}.

::: example
Without the
[`decode()`{#the-img-element:dom-img-decode-2}](#dom-img-decode) method,
the process of loading an
[`img`{#the-img-element:the-img-element-24}](#the-img-element) element
and then displaying it might look like the following:

``` js
const img = new Image();
img.src = "nebula.jpg";
img.onload = () => {
    document.body.appendChild(img);
};
img.onerror = () => {
    document.body.appendChild(new Text("Could not load the nebula :("));
};
```

However, this can cause notable dropped frames, as the paint that occurs
after inserting the image into the DOM causes a synchronous decode on
the main thread.

This can instead be rewritten using the
[`decode()`{#the-img-element:dom-img-decode-3}](#dom-img-decode) method:

``` js
const img = new Image();
img.src = "nebula.jpg";
img.decode().then(() => {
    document.body.appendChild(img);
}).catch(() => {
    document.body.appendChild(new Text("Could not load the nebula :("));
});
```

This latter form avoids the dropped frames of the original, by allowing
the user agent to decode the image [in
parallel](infrastructure.html#in-parallel){#the-img-element:in-parallel-3},
and only inserting it into the DOM (and thus causing it to be painted)
once the decoding process is complete.
:::

::: example
Because the
[`decode()`{#the-img-element:dom-img-decode-4}](#dom-img-decode) method
attempts to ensure that the decoded image data is available for at least
one frame, it can be combined with the
[`requestAnimationFrame()`{#the-img-element:dom-animationframeprovider-requestanimationframe}](imagebitmap-and-animations.html#dom-animationframeprovider-requestanimationframe)
API. This means it can be used with coding styles or frameworks that
ensure that all DOM modifications are batched together as [animation
frame
callbacks](imagebitmap-and-animations.html#list-of-animation-frame-callbacks){#the-img-element:list-of-animation-frame-callbacks}:

``` js
const container = document.querySelector("#container");

const { containerWidth, containerHeight } = computeDesiredSize();
requestAnimationFrame(() => {
 container.style.width = containerWidth;
 container.style.height = containerHeight;
});

// ...

const img = new Image();
img.src = "supernova.jpg";
img.decode().then(() => {
    requestAnimationFrame(() => container.appendChild(img));
});
```
:::

A legacy factory function is provided for creating
[`HTMLImageElement`{#the-img-element:htmlimageelement}](#htmlimageelement)
objects (in addition to the factory methods from DOM such as
[`createElement()`{#the-img-element:dom-document-createelement}](https://dom.spec.whatwg.org/#dom-document-createelement){x-internal="dom-document-createelement"}):
[`Image(``width`{.variable}`, ``height`{.variable}`)`]{#dom-image .dfn}.
When invoked, the legacy factory function must perform the following
steps:

1.  Let `document`{.variable} be the [current global
    object](webappapis.html#current-global-object){#the-img-element:current-global-object}\'s
    [associated
    `Document`](nav-history-apis.html#concept-document-window){#the-img-element:concept-document-window}.

2.  Let `img`{.variable} be the result of [creating an
    element](https://dom.spec.whatwg.org/#concept-create-element){#the-img-element:create-an-element
    x-internal="create-an-element"} given `document`{.variable},
    \"`img`\", and the [HTML
    namespace](https://infra.spec.whatwg.org/#html-namespace){#the-img-element:html-namespace-2
    x-internal="html-namespace-2"}.

3.  If `width`{.variable} is given, then [set an attribute
    value](https://dom.spec.whatwg.org/#concept-element-attributes-set-value){#the-img-element:concept-element-attributes-set-value
    x-internal="concept-element-attributes-set-value"} for
    `img`{.variable} using
    \"[`width`{#the-img-element:attr-dim-width-4}](embedded-content-other.html#attr-dim-width)\"
    and `width`{.variable}.

4.  If `height`{.variable} is given, then [set an attribute
    value](https://dom.spec.whatwg.org/#concept-element-attributes-set-value){#the-img-element:concept-element-attributes-set-value-2
    x-internal="concept-element-attributes-set-value"} for
    `img`{.variable} using
    \"[`height`{#the-img-element:attr-dim-height-4}](embedded-content-other.html#attr-dim-height)\"
    and `height`{.variable}.

5.  Return `img`{.variable}.

::: example
A single image can have different appropriate alternative text depending
on the context.

In each of the following cases, the same image is used, yet the
[`alt`{#the-img-element:attr-img-alt-17}](#attr-img-alt) text is
different each time. The image is the coat of arms of the Carouge
municipality in the canton Geneva in Switzerland.

Here it is used as a supplementary icon:

``` html
<p>I lived in <img src="carouge.svg" alt=""> Carouge.</p>
```

Here it is used as an icon representing the town:

``` html
<p>Home town: <img src="carouge.svg" alt="Carouge"></p>
```

Here it is used as part of a text on the town:

``` html
<p>Carouge has a coat of arms.</p>
<p><img src="carouge.svg" alt="The coat of arms depicts a lion, sitting in front of a tree."></p>
<p>It is used as decoration all over the town.</p>
```

Here it is used as a way to support a similar text where the description
is given as well as, instead of as an alternative to, the image:

``` html
<p>Carouge has a coat of arms.</p>
<p><img src="carouge.svg" alt=""></p>
<p>The coat of arms depicts a lion, sitting in front of a tree.
It is used as decoration all over the town.</p>
```

Here it is used as part of a story:

``` html
<p>She picked up the folder and a piece of paper fell out.</p>
<p><img src="carouge.svg" alt="Shaped like a shield, the paper had a
red background, a green tree, and a yellow lion with its tongue
hanging out and whose tail was shaped like an S."></p>
<p>She stared at the folder. S! The answer she had been looking for all
this time was simply the letter S! How had she not seen that before? It all
came together now. The phone call where Hector had referred to a lion's tail,
the time Maria had stuck her tongue out...</p>
```

Here it is not known at the time of publication what the image will be,
only that it will be a coat of arms of some kind, and thus no
replacement text can be provided, and instead only a brief caption for
the image is provided, in the
[`title`{#the-img-element:attr-title-3}](dom.html#attr-title) attribute:

``` html
<p>The last user to have uploaded a coat of arms uploaded this one:</p>
<p><img src="last-uploaded-coat-of-arms.cgi" title="User-uploaded coat of arms."></p>
```

Ideally, the author would find a way to provide real replacement text
even in this case, e.g. by asking the previous user. Not providing
replacement text makes the document more difficult to use for people who
are unable to view images, e.g. blind users, or users or very
low-bandwidth connections or who pay by the byte, or users who are
forced to use a text-only web browser.
:::

::: example
Here are some more examples showing the same picture used in different
contexts, with different appropriate alternate texts each time.

``` html
<article>
 <h1>My cats</h1>
 <h2>Fluffy</h2>
 <p>Fluffy is my favorite.</p>
 <img src="fluffy.jpg" alt="She likes playing with a ball of yarn.">
 <p>She's just too cute.</p>
 <h2>Miles</h2>
 <p>My other cat, Miles just eats and sleeps.</p>
</article>
```

``` html
<article>
 <h1>Photography</h1>
 <h2>Shooting moving targets indoors</h2>
 <p>The trick here is to know how to anticipate; to know at what speed and
 what distance the subject will pass by.</p>
 <img src="fluffy.jpg" alt="A cat flying by, chasing a ball of yarn, can be
 photographed quite nicely using this technique.">
 <h2>Nature by night</h2>
 <p>To achieve this, you'll need either an extremely sensitive film, or
 immense flash lights.</p>
</article>
```

``` html
<article>
 <h1>About me</h1>
 <h2>My pets</h2>
 <p>I've got a cat named Fluffy and a dog named Miles.</p>
 <img src="fluffy.jpg" alt="Fluffy, my cat, tends to keep itself busy.">
 <p>My dog Miles and I like go on long walks together.</p>
 <h2>music</h2>
 <p>After our walks, having emptied my mind, I like listening to Bach.</p>
</article>
```

``` html
<article>
 <h1>Fluffy and the Yarn</h1>
 <p>Fluffy was a cat who liked to play with yarn. She also liked to jump.</p>
 <aside><img src="fluffy.jpg" alt="" title="Fluffy"></aside>
 <p>She would play in the morning, she would play in the evening.</p>
</article>
```
:::

[← 4.7 Edits](edits.html) --- [Table of Contents](./) --- [4.8.4 Images
→](images.html)
