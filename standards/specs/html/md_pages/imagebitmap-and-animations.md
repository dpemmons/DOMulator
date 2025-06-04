HTML Standard

::: {#head .head .with-buttons}
[![WHATWG](https://resources.whatwg.org/logo.svg){.darkmode-aware
width="100" crossorigin="" height="100"}](https://whatwg.org/){.logo}

# HTML {#html .allcaps}

Living Standard --- Last Updated [2 June 2025]{.pubdate}
:::

[← 8.9 System state and capabilities](system-state.html) --- [Table of
Contents](./) --- [9 Communication →](comms.html)

1.  ::: {#toc-webappapis}
    1.  [[8.10]{.secno}
        Images](imagebitmap-and-animations.html#images-2)
        1.  [[8.10.1]{.secno} The `ImageData`
            interface](imagebitmap-and-animations.html#the-imagedata-interface)
        2.  [[8.10.2]{.secno} The `ImageBitmap`
            interface](imagebitmap-and-animations.html#the-imagebitmap-interface)
    2.  [[8.11]{.secno} Animation
        frames](imagebitmap-and-animations.html#animation-frames)
    :::

### [8.10]{.secno} Images[](#images-2){.self-link} {#images-2}

#### [8.10.1]{.secno} The [`ImageData`{#the-imagedata-interface:imagedata}](#imagedata) interface[](#the-imagedata-interface){.self-link}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[ImageData](https://developer.mozilla.org/en-US/docs/Web/API/ImageData "The ImageData interface represents the underlying pixel data of an area of a <canvas> element.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera9+]{.opera .yes}[Edge79+]{.edge_blink .yes}

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

``` idl
typedef (Uint8ClampedArray or Float16Array) ImageDataArray;

enum ImageDataPixelFormat { "rgba-unorm8", "rgba-float16" };

dictionary ImageDataSettings {
  PredefinedColorSpace colorSpace;
  ImageDataPixelFormat pixelFormat = "rgba-unorm8";
};

[Exposed=(Window,Worker),
 Serializable]
interface ImageData {
  constructor(unsigned long sw, unsigned long sh, optional ImageDataSettings settings = {});
  constructor(ImageDataArray data, unsigned long sw, optional unsigned long sh, optional ImageDataSettings settings = {});

  readonly attribute unsigned long width;
  readonly attribute unsigned long height;
  readonly attribute ImageDataArray data;
  readonly attribute ImageDataPixelFormat pixelFormat;
  readonly attribute PredefinedColorSpace colorSpace;
};
```

An [`ImageData`{#the-imagedata-interface:imagedata-2}](#imagedata)
object [ represents a rectanglar
bitmap]{#concept-imagedata-bitmap-representation .dfn} with width equal
to the
` `{#the-imagedata-interface:dom-imagedata-width-2}[`width`{#the-imagedata-interface:dom-imagedata-width-2}](#dom-imagedata-width)
attribute and height equal to the
[`height`{#the-imagedata-interface:dom-imagedata-height-2}](#dom-imagedata-height)
attribute. The pixel values of this bitmap are stored in the
` `{#the-imagedata-interface:dom-imagedata-data-2}[`data`{#the-imagedata-interface:dom-imagedata-data-2}](#dom-imagedata-data)
attribute in left-to-right order, row by row from top to bottom,
starting with 0 for the top left pixel, with the order and numerical
representation of the color components of each pixel determined by the
[`pixelFormat`{#the-imagedata-interface:dom-imagedata-pixelformat-2}](#dom-imagedata-pixelformat)
attribute. The color space of the pixel values of the bitmap is
determined by the
[`colorSpace`{#the-imagedata-interface:dom-imagedata-colorspace-2}](#dom-imagedata-colorspace)
attribute.

`imageData`{.variable}` = new `[`ImageData`](#dom-imagedata){#dom-imagedata-dev}`(``sw`{.variable}`, ``sh`{.variable}` [, ``settings`{.variable}`])`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[ImageData/ImageData](https://developer.mozilla.org/en-US/docs/Web/API/ImageData/ImageData "The ImageData() constructor returns a newly instantiated ImageData object built from the typed array given and having the specified width and height.")

Support in all current engines.

::: support
[Firefox29+]{.firefox .yes}[Safari8+]{.safari .yes}[Chrome36+]{.chrome
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

Returns an
[`ImageData`{#the-imagedata-interface:imagedata-3}](#imagedata) object
with the given dimensions and the color space indicated by
`settings`{.variable}. All the pixels in the returned object are
[transparent
black](https://drafts.csswg.org/css-color/#transparent-black){#the-imagedata-interface:transparent-black
x-internal="transparent-black"}.

Throws an
[\"`IndexSizeError`\"](https://webidl.spec.whatwg.org/#indexsizeerror){#the-imagedata-interface:indexsizeerror
x-internal="indexsizeerror"}
[`DOMException`{#the-imagedata-interface:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if either of the width or height arguments are zero.

`imageData`{.variable}` = new `[`ImageData`](#dom-imagedata-with-data){#dom-imagedata-with-data-dev}`(``data`{.variable}`, ``sw`{.variable}` [, ``sh`{.variable}` [, ``settings`{.variable}` ] ])`

Returns an
[`ImageData`{#the-imagedata-interface:imagedata-4}](#imagedata) object
using the data provided in the
[`ImageDataArray`{#the-imagedata-interface:imagedataarray-3}](#imagedataarray)
argument, interpreted using the given dimensions and the color space
indicated by `settings`{.variable}.

The byte length of the data needs to be a multiple of the number of
bytes per pixel times the given width. If the height is provided as
well, then the length needs to be exactly the number of bytes per pixel
times the width times the height.

Throws an
[\"`IndexSizeError`\"](https://webidl.spec.whatwg.org/#indexsizeerror){#the-imagedata-interface:indexsizeerror-2
x-internal="indexsizeerror"}
[`DOMException`{#the-imagedata-interface:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if the given data and dimensions can\'t be interpreted consistently, or
if either dimension is zero.

`imageData`{.variable}`.`[`width`](#dom-imagedata-width){#dom-imagedata-width-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[ImageData/width](https://developer.mozilla.org/en-US/docs/Web/API/ImageData/width "The readonly ImageData.width property returns the number of pixels per row in the ImageData object.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera9+]{.opera .yes}[Edge79+]{.edge_blink .yes}

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

`imageData`{.variable}`.`[`height`](#dom-imagedata-height){#dom-imagedata-height-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[ImageData/height](https://developer.mozilla.org/en-US/docs/Web/API/ImageData/height "The readonly ImageData.height property returns the number of rows in the ImageData object.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera9+]{.opera .yes}[Edge79+]{.edge_blink .yes}

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

Returns the actual dimensions of the data in the
[`ImageData`{#the-imagedata-interface:imagedata-5}](#imagedata) object,
in pixels.

`imageData`{.variable}`.`[`data`](#dom-imagedata-data){#dom-imagedata-data-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[ImageData/data](https://developer.mozilla.org/en-US/docs/Web/API/ImageData/data "The readonly ImageData.data property returns a Uint8ClampedArray that contains the ImageData object's pixel data. Data is stored as a one-dimensional array in the RGBA order, with integer values between 0 and 255 (inclusive).")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera9+]{.opera .yes}[Edge79+]{.edge_blink .yes}

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

Returns the one-dimensional array containing the data in RGBA order, as
integers in the range 0 to 255.

`imageData`{.variable}`.`[`colorSpace`](#dom-imagedata-colorspace){#dom-imagedata-colorspace-dev}

Returns the color space of the pixels.

The
[`new ImageData(``sw`{.variable}`, ``sh`{.variable}`, ``settings`{.variable}`)`]{#dom-imagedata
.dfn dfn-for="ImageData" dfn-type="constructor"} constructor steps are:

1.  If one or both of `sw`{.variable} and `sh`{.variable} are zero, then
    throw an
    [\"`IndexSizeError`\"](https://webidl.spec.whatwg.org/#indexsizeerror){#the-imagedata-interface:indexsizeerror-3
    x-internal="indexsizeerror"}
    [`DOMException`{#the-imagedata-interface:domexception-3}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  [Initialize](#initialize-an-imagedata-object){#the-imagedata-interface:initialize-an-imagedata-object}
    [this](https://webidl.spec.whatwg.org/#this){#the-imagedata-interface:this
    x-internal="this"} given `sw`{.variable}, `sh`{.variable}, and
    `settings`{.variable}.

3.  Initialize the image data of
    [this](https://webidl.spec.whatwg.org/#this){#the-imagedata-interface:this-2
    x-internal="this"} to [transparent
    black](https://drafts.csswg.org/css-color/#transparent-black){#the-imagedata-interface:transparent-black-2
    x-internal="transparent-black"}.

The
[`new ImageData(``data`{.variable}`, ``sw`{.variable}`, ``sh`{.variable}`, ``settings`{.variable}`)`]{#dom-imagedata-with-data
.dfn dfn-for="ImageData" dfn-type="constructor"} constructor steps are:

1.  Let `bytesPerPixel`{.variable} be 4 if
    `settings`{.variable}\[\"[`pixelFormat`{#the-imagedata-interface:dom-imagedatasettings-pixelformat-2}](#dom-imagedatasettings-pixelformat)\"\]
    is
    \"[rgba-unorm8](#dom-imagedatapixelformat-rgba-unorm8){#the-imagedata-interface:dom-imagedatapixelformat-rgba-unorm8-3}\";
    otherwise 8.

2.  Let `length`{.variable} be the [buffer source byte
    length](https://webidl.spec.whatwg.org/#buffersource-byte-length){#the-imagedata-interface:buffer-source-byte-length
    x-internal="buffer-source-byte-length"} of `data`{.variable}.

3.  If `length`{.variable} is not a nonzero integral multiple of
    `bytesPerPixel`{.variable}, then throw an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-imagedata-interface:invalidstateerror
    x-internal="invalidstateerror"}
    [`DOMException`{#the-imagedata-interface:domexception-4}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

4.  Let `length`{.variable} be `length`{.variable} divided by
    `bytesPerPixel`{.variable}.

5.  If `length`{.variable} is not an integral multiple of
    `sw`{.variable}, then throw an
    [\"`IndexSizeError`\"](https://webidl.spec.whatwg.org/#indexsizeerror){#the-imagedata-interface:indexsizeerror-4
    x-internal="indexsizeerror"}
    [`DOMException`{#the-imagedata-interface:domexception-5}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    At this step, the length is guaranteed to be greater than zero
    (otherwise the second step above would have aborted the steps), so
    if `sw`{.variable} is zero, this step will throw the exception and
    return.

6.  Let `height`{.variable} be `length`{.variable} divided by
    `sw`{.variable}.

7.  If `sh`{.variable} was given and its value is not equal to
    `height`{.variable}, then throw an
    [\"`IndexSizeError`\"](https://webidl.spec.whatwg.org/#indexsizeerror){#the-imagedata-interface:indexsizeerror-5
    x-internal="indexsizeerror"}
    [`DOMException`{#the-imagedata-interface:domexception-6}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

8.  [Initialize](#initialize-an-imagedata-object){#the-imagedata-interface:initialize-an-imagedata-object-2}
    [this](https://webidl.spec.whatwg.org/#this){#the-imagedata-interface:this-3
    x-internal="this"} given `sw`{.variable}, `sh`{.variable},
    `settings`{.variable}, and
    *[source](#initialize-imagedata-source){#the-imagedata-interface:initialize-imagedata-source}*
    set to `data`{.variable}.

    This step does not set
    [this](https://webidl.spec.whatwg.org/#this){#the-imagedata-interface:this-4
    x-internal="this"}\'s data to a copy of `data`{.variable}. It sets
    it to the actual
    [`ImageDataArray`{#the-imagedata-interface:imagedataarray-4}](#imagedataarray)
    object passed as `data`{.variable}.

To [initialize an `ImageData` object]{#initialize-an-imagedata-object
.dfn} `imageData`{.variable}, given a positive integer number of pixels
per row `pixelsPerRow`{.variable}, a positive integer number of rows
`rows`{.variable}, an
[`ImageDataSettings`{#the-imagedata-interface:imagedatasettings-3}](#imagedatasettings)
`settings`{.variable}, an optional
[`ImageDataArray`{#the-imagedata-interface:imagedataarray-5}](#imagedataarray)
[`source`{.variable}]{#initialize-imagedata-source .dfn}, and an
optional
[`PredefinedColorSpace`{#the-imagedata-interface:predefinedcolorspace-3}](canvas.html#predefinedcolorspace)
[`defaultColorSpace`{.variable}]{#initialize-imagedata-defaultcolorspace
.dfn}:

::::: {.mdn-anno .wrapped .before}
MDN

:::: feature
[ImageData/colorSpace](https://developer.mozilla.org/en-US/docs/Web/API/ImageData/colorSpace "The read-only ImageData.colorSpace property is a string indicating the color space of the image data.")

::: support
[FirefoxNo]{.firefox .no}[Safari15.2+]{.safari .yes}[Chrome92+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge92+]{.edge_blink .yes}

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

1.  If `source`{.variable} was given:

    1.  If
        `settings`{.variable}\[\"[`pixelFormat`{#the-imagedata-interface:dom-imagedatasettings-pixelformat-3}](#dom-imagedatasettings-pixelformat)\"\]
        equals
        \"[rgba-unorm8](#dom-imagedatapixelformat-rgba-unorm8){#the-imagedata-interface:dom-imagedatapixelformat-rgba-unorm8-4}\"
        and `source`{.variable} is not a
        [`Uint8ClampedArray`{#the-imagedata-interface:idl-uint8clampedarray-2}](https://webidl.spec.whatwg.org/#idl-Uint8ClampedArray){x-internal="idl-uint8clampedarray"},
        then throw an
        [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-imagedata-interface:invalidstateerror-2
        x-internal="invalidstateerror"}
        [`DOMException`{#the-imagedata-interface:domexception-7}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    2.  If
        `settings`{.variable}\[\"[`pixelFormat`{#the-imagedata-interface:dom-imagedatasettings-pixelformat-4}](#dom-imagedatasettings-pixelformat)\"\]
        is
        \"[rgba-float16](#dom-imagedatapixelformat-rgba-float16){#the-imagedata-interface:dom-imagedatapixelformat-rgba-float16-2}\"
        and `source`{.variable} is not a
        [`Float16Array`{#the-imagedata-interface:idl-float16array-2}](https://webidl.spec.whatwg.org/#idl-Float16Array){x-internal="idl-float16array"},
        then throw an
        [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-imagedata-interface:invalidstateerror-3
        x-internal="invalidstateerror"}
        [`DOMException`{#the-imagedata-interface:domexception-8}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    3.  Initialize the [`data`]{#dom-imagedata-data .dfn
        dfn-for="ImageData" dfn-type="attribute"} attribute of
        `imageData`{.variable} to `source`{.variable}.

2.  Otherwise (`source`{.variable} was not given):

    1.  If
        `settings`{.variable}\[\"[`pixelFormat`{#the-imagedata-interface:dom-imagedatasettings-pixelformat-5}](#dom-imagedatasettings-pixelformat)\"\]
        is
        \"[rgba-unorm8](#dom-imagedatapixelformat-rgba-unorm8){#the-imagedata-interface:dom-imagedatapixelformat-rgba-unorm8-5}\",
        then initialize the
        [`data`{#the-imagedata-interface:dom-imagedata-data-3}](#dom-imagedata-data)
        attribute of `imageData`{.variable} to a new
        [`Uint8ClampedArray`{#the-imagedata-interface:idl-uint8clampedarray-3}](https://webidl.spec.whatwg.org/#idl-Uint8ClampedArray){x-internal="idl-uint8clampedarray"}
        object. The
        [`Uint8ClampedArray`{#the-imagedata-interface:idl-uint8clampedarray-4}](https://webidl.spec.whatwg.org/#idl-Uint8ClampedArray){x-internal="idl-uint8clampedarray"}
        object must use a new
        [`ArrayBuffer`{#the-imagedata-interface:idl-arraybuffer}](https://webidl.spec.whatwg.org/#idl-ArrayBuffer){x-internal="idl-arraybuffer"}
        for its storage, and must have a zero byte offset and byte
        length equal to the length of its storage, in bytes. The storage
        [`ArrayBuffer`{#the-imagedata-interface:idl-arraybuffer-2}](https://webidl.spec.whatwg.org/#idl-ArrayBuffer){x-internal="idl-arraybuffer"}
        must have a length of 4 × `rows`{.variable} ×
        `pixelsPerRow`{.variable} bytes.

    2.  Otherwise, if
        `settings`{.variable}\[\"[`pixelFormat`{#the-imagedata-interface:dom-imagedatasettings-pixelformat-6}](#dom-imagedatasettings-pixelformat)\"\]
        is
        \"[rgba-float16](#dom-imagedatapixelformat-rgba-unorm8){#the-imagedata-interface:dom-imagedatapixelformat-rgba-unorm8-6}\",
        then initialize the
        [`data`{#the-imagedata-interface:dom-imagedata-data-4}](#dom-imagedata-data)
        attribute of `imageData`{.variable} to a new
        [`Float16Array`{#the-imagedata-interface:idl-float16array-3}](https://webidl.spec.whatwg.org/#idl-Float16Array){x-internal="idl-float16array"}
        object. The
        [`Float16Array`{#the-imagedata-interface:idl-float16array-4}](https://webidl.spec.whatwg.org/#idl-Float16Array){x-internal="idl-float16array"}
        object must use a new
        [`ArrayBuffer`{#the-imagedata-interface:idl-arraybuffer-3}](https://webidl.spec.whatwg.org/#idl-ArrayBuffer){x-internal="idl-arraybuffer"}
        for its storage, and must have a zero byte offset and byte
        length equal to the length of its storage, in bytes. The storage
        [`ArrayBuffer`{#the-imagedata-interface:idl-arraybuffer-4}](https://webidl.spec.whatwg.org/#idl-ArrayBuffer){x-internal="idl-arraybuffer"}
        must have a length of 8 × `rows`{.variable} ×
        `pixelsPerRow`{.variable} bytes.

    3.  If the storage
        [`ArrayBuffer`{#the-imagedata-interface:idl-arraybuffer-5}](https://webidl.spec.whatwg.org/#idl-ArrayBuffer){x-internal="idl-arraybuffer"}
        could not be allocated, then rethrow the
        [`RangeError`{#the-imagedata-interface:js-rangeerror}](https://tc39.es/ecma262/#sec-native-error-types-used-in-this-standard-rangeerror){x-internal="js-rangeerror"}
        thrown by JavaScript, and return.

3.  Initialize the [`width`]{#dom-imagedata-width .dfn
    dfn-for="ImageData" dfn-type="attribute"} attribute of
    `imageData`{.variable} to `pixelsPerRow`{.variable}.

4.  Initialize the [`height`]{#dom-imagedata-height .dfn
    dfn-for="ImageData" dfn-type="attribute"} attribute of
    `imageData`{.variable} to `rows`{.variable}.

5.  Initialize the [`pixelFormat`]{#dom-imagedata-pixelformat .dfn
    dfn-for="ImageData" dfn-type="attribute"} attribute of
    `imageData`{.variable} to
    `settings`{.variable}\[\"[`pixelFormat`]{#dom-imagedatasettings-pixelformat
    .dfn dfn-for="ImageDataSettings" dfn-type="dict-member"}\"\].

6.  If
    `settings`{.variable}\[\"[`colorSpace`{#the-imagedata-interface:dom-imagedatasettings-colorspace}](#dom-imagedatasettings-colorspace)\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#the-imagedata-interface:map-exists
    x-internal="map-exists"}, then initialize the
    [`colorSpace`]{#dom-imagedata-colorspace .dfn dfn-for="ImageData"
    dfn-type="attribute"} attribute of `imageData`{.variable} to
    `settings`{.variable}\[\"[`colorSpace`]{#dom-imagedatasettings-colorspace
    .dfn dfn-for="ImageDataSettings" dfn-type="dict-member"}\"\].

7.  Otherwise, if `defaultColorSpace`{.variable} was given, then
    initialize the
    [`colorSpace`{#the-imagedata-interface:dom-imagedata-colorspace-3}](#dom-imagedata-colorspace)
    attribute of `imageData`{.variable} to
    `defaultColorSpace`{.variable}.

8.  Otherwise, initialize the
    [`colorSpace`{#the-imagedata-interface:dom-imagedata-colorspace-4}](#dom-imagedata-colorspace)
    attribute of `imageData`{.variable} to
    \"[srgb](canvas.html#dom-predefinedcolorspace-srgb){#the-imagedata-interface:dom-predefinedcolorspace-srgb-2}\".

[`ImageData`{#the-imagedata-interface:imagedata-6}](#imagedata) objects
are [serializable
objects](structured-data.html#serializable-objects){#the-imagedata-interface:serializable-objects}.
Their [serialization
steps](structured-data.html#serialization-steps){#the-imagedata-interface:serialization-steps},
given `value`{.variable} and `serialized`{.variable}, are:

1.  Set `serialized`{.variable}.\[\[Data\]\] to the
    [sub-serialization](structured-data.html#sub-serialization){#the-imagedata-interface:sub-serialization}
    of the value of `value`{.variable}\'s
    [`data`{#the-imagedata-interface:dom-imagedata-data-5}](#dom-imagedata-data)
    attribute.

2.  Set `serialized`{.variable}.\[\[Width\]\] to the value of
    `value`{.variable}\'s
    [`width`{#the-imagedata-interface:dom-imagedata-width-3}](#dom-imagedata-width)
    attribute.

3.  Set `serialized`{.variable}.\[\[Height\]\] to the value of
    `value`{.variable}\'s
    [`height`{#the-imagedata-interface:dom-imagedata-height-3}](#dom-imagedata-height)
    attribute.

4.  Set `serialized`{.variable}.\[\[ColorSpace\]\] to the value of
    `value`{.variable}\'s
    [`colorSpace`{#the-imagedata-interface:dom-imagedata-colorspace-5}](#dom-imagedata-colorspace)
    attribute.

5.  Set `serialized`{.variable}.\[\[PixelFormat\]\] to the value of
    `value`{.variable}\'s
    [`pixelFormat`{#the-imagedata-interface:dom-imagedata-pixelformat-3}](#dom-imagedata-pixelformat)
    attribute.

Their [deserialization
steps](structured-data.html#deserialization-steps){#the-imagedata-interface:deserialization-steps},
given `serialized`{.variable}, `value`{.variable}, and
`targetRealm`{.variable}, are:

1.  Initialize `value`{.variable}\'s
    [`data`{#the-imagedata-interface:dom-imagedata-data-6}](#dom-imagedata-data)
    attribute to the
    [sub-deserialization](structured-data.html#sub-deserialization){#the-imagedata-interface:sub-deserialization}
    of `serialized`{.variable}.\[\[Data\]\].

2.  Initialize `value`{.variable}\'s
    [`width`{#the-imagedata-interface:dom-imagedata-width-4}](#dom-imagedata-width)
    attribute to `serialized`{.variable}.\[\[Width\]\].

3.  Initialize `value`{.variable}\'s
    [`height`{#the-imagedata-interface:dom-imagedata-height-4}](#dom-imagedata-height)
    attribute to `serialized`{.variable}.\[\[Height\]\].

4.  Initialize `value`{.variable}\'s
    [`colorSpace`{#the-imagedata-interface:dom-imagedata-colorspace-6}](#dom-imagedata-colorspace)
    attribute to `serialized`{.variable}.\[\[ColorSpace\]\].

5.  Initialize `value`{.variable}\'s
    [`pixelFormat`{#the-imagedata-interface:dom-imagedata-pixelformat-4}](#dom-imagedata-pixelformat)
    attribute to `serialized`{.variable}.\[\[PixelFormat\]\].

The
[`ImageDataPixelFormat`{#the-imagedata-interface:imagedatapixelformat-3}](#imagedatapixelformat)
enumeration is used to specify type of the
[`data`{#the-imagedata-interface:dom-imagedata-data-7}](#dom-imagedata-data)
attribute of an
[`ImageData`{#the-imagedata-interface:imagedata-7}](#imagedata) and the
arrangement and numerical representation of the color components for
each pixel.

The \"[`rgba-unorm8`]{#dom-imagedatapixelformat-rgba-unorm8 .dfn
dfn-for="ImageDataPixelFormat" dfn-type="enum-value"}\" value indicates
that the
[`data`{#the-imagedata-interface:dom-imagedata-data-8}](#dom-imagedata-data)
attribute of an
[`ImageData`{#the-imagedata-interface:imagedata-8}](#imagedata) must be
of type
[Uint8ClampedArray](https://webidl.spec.whatwg.org/#idl-Uint8ClampedArray){#the-imagedata-interface:idl-uint8clampedarray-5
x-internal="idl-uint8clampedarray"}. The color components of each pixel
must be stored in four sequential elements in the order of red, green,
blue, and then alpha. Each element represents the 8-bit unsigned
normalized value for that component.

The \"[`rgba-float16`]{#dom-imagedatapixelformat-rgba-float16 .dfn
dfn-for="ImageDataPixelFormat" dfn-type="enum-value"}\" value indicates
that the
[`data`{#the-imagedata-interface:dom-imagedata-data-9}](#dom-imagedata-data)
attribute of an
[`ImageData`{#the-imagedata-interface:imagedata-9}](#imagedata) must be
of type
[Float16Array](https://webidl.spec.whatwg.org/#idl-Float16Array){#the-imagedata-interface:idl-float16array-5
x-internal="idl-float16array"}. The color components of each pixel must
be stored in four sequential elements in the order of red, green, blue,
and then alpha. Each element represents the value for that component.

#### [8.10.2]{.secno} The [`ImageBitmap`{#the-imagebitmap-interface:imagebitmap}](#imagebitmap) interface[](#the-imagebitmap-interface){.self-link}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[ImageBitmap](https://developer.mozilla.org/en-US/docs/Web/API/ImageBitmap "The ImageBitmap interface represents a bitmap image which can be drawn to a <canvas> without undue latency. It can be created from a variety of source objects using the createImageBitmap() factory method. ImageBitmap provides an asynchronous and resource efficient pathway to prepare textures for rendering in WebGL.")

Support in all current engines.

::: support
[Firefox42+]{.firefox .yes}[Safari15+]{.safari .yes}[Chrome50+]{.chrome
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

``` idl
[Exposed=(Window,Worker), Serializable, Transferable]
interface ImageBitmap {
  readonly attribute unsigned long width;
  readonly attribute unsigned long height;
  undefined close();
};

typedef (CanvasImageSource or
         Blob or
         ImageData) ImageBitmapSource;

enum ImageOrientation { "from-image", "flipY" };
enum PremultiplyAlpha { "none", "premultiply", "default" };
enum ColorSpaceConversion { "none", "default" };
enum ResizeQuality { "pixelated", "low", "medium", "high" };

dictionary ImageBitmapOptions {
  ImageOrientation imageOrientation = "from-image";
  PremultiplyAlpha premultiplyAlpha = "default";
  ColorSpaceConversion colorSpaceConversion = "default";
  [EnforceRange] unsigned long resizeWidth;
  [EnforceRange] unsigned long resizeHeight;
  ResizeQuality resizeQuality = "low";
};
```

An
[`ImageBitmap`{#the-imagebitmap-interface:imagebitmap-2}](#imagebitmap)
object represents a bitmap image that can be painted to a canvas without
undue latency.

The exact judgement of what is undue latency of this is left up to the
implementer, but in general if making use of the bitmap requires network
I/O, or even local disk I/O, then the latency is probably undue; whereas
if it only requires a blocking read from a GPU or system RAM, the
latency is probably acceptable.

`promise`{.variable}` = self.`[`createImageBitmap`](#dom-createimagebitmap){#dom-createimagebitmap-dev}`(``image`{.variable}` [, ``options`{.variable}` ])`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[createImageBitmap](https://developer.mozilla.org/en-US/docs/Web/API/createImageBitmap "The createImageBitmap() method creates a bitmap from a given source, optionally cropped to contain only a portion of that source. The method exists on the global scope in both windows and workers. It accepts a variety of different image sources, and returns a Promise which resolves to an ImageBitmap.")

Support in all current engines.

::: support
[Firefox42+]{.firefox .yes}[Safari15+]{.safari .yes}[Chrome50+]{.chrome
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

`promise`{.variable}` = self.`[`createImageBitmap`](#dom-createimagebitmap){#the-imagebitmap-interface:dom-createimagebitmap}`(``image`{.variable}`, ``sx`{.variable}`, ``sy`{.variable}`, ``sw`{.variable}`, ``sh`{.variable}` [, ``options`{.variable}` ])`

Takes `image`{.variable}, which can be an
[`img`{#the-imagebitmap-interface:the-img-element}](embedded-content.html#the-img-element)
element, an [SVG
`image`](https://svgwg.org/svg2-draft/embedded.html#ImageElement){#the-imagebitmap-interface:svg-image
x-internal="svg-image"} element, a
[`video`{#the-imagebitmap-interface:the-video-element}](media.html#the-video-element)
element, a
[`canvas`{#the-imagebitmap-interface:the-canvas-element}](canvas.html#the-canvas-element)
element, a
[`Blob`{#the-imagebitmap-interface:blob-2}](https://w3c.github.io/FileAPI/#dfn-Blob){x-internal="blob"}
object, an
[`ImageData`{#the-imagebitmap-interface:imagedata-2}](#imagedata)
object, or another
[`ImageBitmap`{#the-imagebitmap-interface:imagebitmap-3}](#imagebitmap)
object, and returns a promise that is resolved when a new
[`ImageBitmap`{#the-imagebitmap-interface:imagebitmap-4}](#imagebitmap)
is created.

If no
[`ImageBitmap`{#the-imagebitmap-interface:imagebitmap-5}](#imagebitmap)
object can be constructed, for example because the provided
`image`{.variable} data is not actually an image, then the promise is
rejected instead.

If `sx`{.variable}, `sy`{.variable}, `sw`{.variable}, and
`sh`{.variable} arguments are provided, the source image is cropped to
the given pixels, with any pixels missing in the original replaced by
[transparent
black](https://drafts.csswg.org/css-color/#transparent-black){#the-imagebitmap-interface:transparent-black
x-internal="transparent-black"}. These coordinates are in the source
image\'s pixel coordinate space, *not* in [CSS
pixels](https://drafts.csswg.org/css-values/#px){#the-imagebitmap-interface:'px'
x-internal="'px'"}.

If `options`{.variable} is provided, the
[`ImageBitmap`{#the-imagebitmap-interface:imagebitmap-6}](#imagebitmap)
object\'s bitmap data is modified according to `options`{.variable}. For
example, if the
[`premultiplyAlpha`{#the-imagebitmap-interface:dom-imagebitmapoptions-premultiplyalpha-2}](#dom-imagebitmapoptions-premultiplyalpha)
option is set to
\"[`premultiply`{#the-imagebitmap-interface:dom-premultiplyalpha-premultiply-2}](#dom-premultiplyalpha-premultiply)\",
the [bitmap
data](#concept-imagebitmap-bitmap-data){#the-imagebitmap-interface:concept-imagebitmap-bitmap-data}\'s
non-alpha color components are [premultiplied by the alpha
component](canvas.html#concept-premultiplied-alpha){#the-imagebitmap-interface:concept-premultiplied-alpha}.

Rejects the promise with an
[\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-imagebitmap-interface:invalidstateerror
x-internal="invalidstateerror"}
[`DOMException`{#the-imagebitmap-interface:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if the source image is not in a valid state (e.g., an
[`img`{#the-imagebitmap-interface:the-img-element-2}](embedded-content.html#the-img-element)
element that hasn\'t loaded successfully, an
[`ImageBitmap`{#the-imagebitmap-interface:imagebitmap-7}](#imagebitmap)
object whose
[\[\[Detached\]\]](structured-data.html#detached){#the-imagebitmap-interface:detached}
internal slot value is true, an
[`ImageData`{#the-imagebitmap-interface:imagedata-3}](#imagedata) object
whose
[`data`{#the-imagebitmap-interface:dom-imagedata-data}](#dom-imagedata-data)
attribute value\'s \[\[ViewedArrayBuffer\]\] internal slot is detached,
or a
[`Blob`{#the-imagebitmap-interface:blob-3}](https://w3c.github.io/FileAPI/#dfn-Blob){x-internal="blob"}
whose data cannot be interpreted as a bitmap image).

Rejects the promise with a
[\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#the-imagebitmap-interface:securityerror
x-internal="securityerror"}
[`DOMException`{#the-imagebitmap-interface:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if the script is not allowed to access the image data of the source
image (e.g. a
[`video`{#the-imagebitmap-interface:the-video-element-2}](media.html#the-video-element)
that is
[CORS-cross-origin](urls-and-fetching.html#cors-cross-origin){#the-imagebitmap-interface:cors-cross-origin},
or a
[`canvas`{#the-imagebitmap-interface:the-canvas-element-2}](canvas.html#the-canvas-element)
being drawn on by a script in a worker from another
[origin](browsers.html#concept-origin){#the-imagebitmap-interface:concept-origin}).

`imageBitmap`{.variable}`.`[`close`](#dom-imagebitmap-close){#dom-imagebitmap-close-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[ImageBitmap/close](https://developer.mozilla.org/en-US/docs/Web/API/ImageBitmap/close "The ImageBitmap.close() method disposes of all graphical resources associated with an ImageBitmap.")

Support in all current engines.

::: support
[Firefox46+]{.firefox .yes}[Safari15+]{.safari .yes}[Chrome52+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera37+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android37+]{.opera_android .yes}
:::
::::
:::::

Releases `imageBitmap`{.variable}\'s underlying [bitmap
data](#concept-imagebitmap-bitmap-data){#the-imagebitmap-interface:concept-imagebitmap-bitmap-data-2}.

`imageBitmap`{.variable}`.`[`width`](#dom-imagebitmap-width){#dom-imagebitmap-width-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[ImageBitmap/width](https://developer.mozilla.org/en-US/docs/Web/API/ImageBitmap/width "The read-only ImageBitmap.width property returns the ImageBitmap object's width in CSS pixels.")

Support in all current engines.

::: support
[Firefox42+]{.firefox .yes}[Safari15+]{.safari .yes}[Chrome50+]{.chrome
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

Returns the [natural
width](https://drafts.csswg.org/css-images/#natural-width){#the-imagebitmap-interface:natural-width
x-internal="natural-width"} of the image, in [CSS
pixels](https://drafts.csswg.org/css-values/#px){#the-imagebitmap-interface:'px'-2
x-internal="'px'"}.

`imageBitmap`{.variable}`.`[`height`](#dom-imagebitmap-height){#dom-imagebitmap-height-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[ImageBitmap/height](https://developer.mozilla.org/en-US/docs/Web/API/ImageBitmap/height "The read-only ImageBitmap.height property returns the ImageBitmap object's height in CSS pixels.")

Support in all current engines.

::: support
[Firefox42+]{.firefox .yes}[Safari15+]{.safari .yes}[Chrome50+]{.chrome
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

Returns the [natural
height](https://drafts.csswg.org/css-images/#natural-height){#the-imagebitmap-interface:natural-height
x-internal="natural-height"} of the image, in [CSS
pixels](https://drafts.csswg.org/css-values/#px){#the-imagebitmap-interface:'px'-3
x-internal="'px'"}.

An
[`ImageBitmap`{#the-imagebitmap-interface:imagebitmap-8}](#imagebitmap)
object whose
[\[\[Detached\]\]](structured-data.html#detached){#the-imagebitmap-interface:detached-2}
internal slot value is false always has associated [bitmap
data]{#concept-imagebitmap-bitmap-data .dfn}, with a width and a height.
However, it is possible for this data to be corrupted. If an
[`ImageBitmap`{#the-imagebitmap-interface:imagebitmap-9}](#imagebitmap)
object\'s media data can be decoded without errors, it is said to be
[fully decodable]{#concept-imagebitmap-good .dfn}.

An
[`ImageBitmap`{#the-imagebitmap-interface:imagebitmap-10}](#imagebitmap)
object\'s bitmap has an
[origin-clean](canvas.html#concept-canvas-origin-clean){#the-imagebitmap-interface:concept-canvas-origin-clean}
flag, which indicates whether the bitmap is tainted by content from a
different
[origin](browsers.html#concept-origin){#the-imagebitmap-interface:concept-origin-2}.
The flag is initially set to true and may be changed to false by the
steps of
[`createImageBitmap()`{#the-imagebitmap-interface:dom-createimagebitmap-2}](#dom-createimagebitmap).

------------------------------------------------------------------------

[`ImageBitmap`{#the-imagebitmap-interface:imagebitmap-11}](#imagebitmap)
objects are [serializable
objects](structured-data.html#serializable-objects){#the-imagebitmap-interface:serializable-objects}
and [transferable
objects](structured-data.html#transferable-objects){#the-imagebitmap-interface:transferable-objects}.

Their [serialization
steps](structured-data.html#serialization-steps){#the-imagebitmap-interface:serialization-steps},
given `value`{.variable} and `serialized`{.variable}, are:

1.  If `value`{.variable}\'s
    [origin-clean](canvas.html#concept-canvas-origin-clean){#the-imagebitmap-interface:concept-canvas-origin-clean-2}
    flag is not set, then throw a
    [\"`DataCloneError`\"](https://webidl.spec.whatwg.org/#datacloneerror){#the-imagebitmap-interface:datacloneerror
    x-internal="datacloneerror"}
    [`DOMException`{#the-imagebitmap-interface:domexception-3}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  Set `serialized`{.variable}.\[\[BitmapData\]\] to a copy of
    `value`{.variable}\'s [bitmap
    data](#concept-imagebitmap-bitmap-data){#the-imagebitmap-interface:concept-imagebitmap-bitmap-data-3}.

Their [deserialization
steps](structured-data.html#deserialization-steps){#the-imagebitmap-interface:deserialization-steps},
given `serialized`{.variable}, `value`{.variable}, and
`targetRealm`{.variable}, are:

1.  Set `value`{.variable}\'s [bitmap
    data](#concept-imagebitmap-bitmap-data){#the-imagebitmap-interface:concept-imagebitmap-bitmap-data-4}
    to `serialized`{.variable}.\[\[BitmapData\]\].

Their [transfer
steps](structured-data.html#transfer-steps){#the-imagebitmap-interface:transfer-steps},
given `value`{.variable} and `dataHolder`{.variable}, are:

1.  If `value`{.variable}\'s
    [origin-clean](canvas.html#concept-canvas-origin-clean){#the-imagebitmap-interface:concept-canvas-origin-clean-3}
    flag is not set, then throw a
    [\"`DataCloneError`\"](https://webidl.spec.whatwg.org/#datacloneerror){#the-imagebitmap-interface:datacloneerror-2
    x-internal="datacloneerror"}
    [`DOMException`{#the-imagebitmap-interface:domexception-4}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  Set `dataHolder`{.variable}.\[\[BitmapData\]\] to
    `value`{.variable}\'s [bitmap
    data](#concept-imagebitmap-bitmap-data){#the-imagebitmap-interface:concept-imagebitmap-bitmap-data-5}.

3.  Unset `value`{.variable}\'s [bitmap
    data](#concept-imagebitmap-bitmap-data){#the-imagebitmap-interface:concept-imagebitmap-bitmap-data-6}.

Their [transfer-receiving
steps](structured-data.html#transfer-receiving-steps){#the-imagebitmap-interface:transfer-receiving-steps},
given `dataHolder`{.variable} and `value`{.variable}, are:

1.  Set `value`{.variable}\'s [bitmap
    data](#concept-imagebitmap-bitmap-data){#the-imagebitmap-interface:concept-imagebitmap-bitmap-data-7}
    to `dataHolder`{.variable}.\[\[BitmapData\]\].

------------------------------------------------------------------------

The
[createImageBitmap](#dom-createimagebitmap){#the-imagebitmap-interface:dom-createimagebitmap-3}
method uses the [bitmap task source]{#bitmap-task-source .dfn} to settle
its returned Promise.

The
[`createImageBitmap(``image`{.variable}`, ``options`{.variable}`)`]{#dom-createimagebitmap
.dfn dfn-for="WindowOrWorkerGlobalScope" dfn-type="method"} and
`createImageBitmap(``image`{.variable}`, ``sx`{.variable}`, ``sy`{.variable}`, ``sw`{.variable}`, ``sh`{.variable}`, ``options`{.variable}`)`
methods, when invoked, must run these steps:

1.  If either `sw`{.variable} or `sh`{.variable} is given and is 0, then
    return [a promise rejected
    with](https://webidl.spec.whatwg.org/#a-promise-rejected-with){#the-imagebitmap-interface:a-promise-rejected-with
    x-internal="a-promise-rejected-with"} a
    [`RangeError`{#the-imagebitmap-interface:js-rangeerror}](https://tc39.es/ecma262/#sec-native-error-types-used-in-this-standard-rangeerror){x-internal="js-rangeerror"}.

2.  If either `options`{.variable}\'s
    [`resizeWidth`]{#dom-imagebitmapoptions-resizewidth .dfn
    dfn-for="ImageBitmapOptions" dfn-type="dict-member"} or
    `options`{.variable}\'s
    [`resizeHeight`]{#dom-imagebitmapoptions-resizeheight .dfn
    dfn-for="ImageBitmapOptions" dfn-type="dict-member"} is present and
    is 0, then return [a promise rejected
    with](https://webidl.spec.whatwg.org/#a-promise-rejected-with){#the-imagebitmap-interface:a-promise-rejected-with-2
    x-internal="a-promise-rejected-with"} an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-imagebitmap-interface:invalidstateerror-2
    x-internal="invalidstateerror"}
    [`DOMException`{#the-imagebitmap-interface:domexception-5}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

3.  [Check the usability of the `image`{.variable}
    argument](canvas.html#check-the-usability-of-the-image-argument){#the-imagebitmap-interface:check-the-usability-of-the-image-argument}.
    If this throws an exception or returns *bad*, then return [a promise
    rejected
    with](https://webidl.spec.whatwg.org/#a-promise-rejected-with){#the-imagebitmap-interface:a-promise-rejected-with-3
    x-internal="a-promise-rejected-with"} an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-imagebitmap-interface:invalidstateerror-3
    x-internal="invalidstateerror"}
    [`DOMException`{#the-imagebitmap-interface:domexception-6}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

4.  Let `promise`{.variable} be a new promise.

5.  Let `imageBitmap`{.variable} be a new
    [`ImageBitmap`{#the-imagebitmap-interface:imagebitmap-12}](#imagebitmap)
    object.

6.  Switch on `image`{.variable}:

    [`img`{#the-imagebitmap-interface:the-img-element-3}](embedded-content.html#the-img-element)\
    [SVG `image`](https://svgwg.org/svg2-draft/embedded.html#ImageElement){#the-imagebitmap-interface:svg-image-2 x-internal="svg-image"}

    :   1.  If `image`{.variable}\'s media data has no [natural
            dimensions](https://drafts.csswg.org/css-images/#natural-dimensions){#the-imagebitmap-interface:natural-dimensions
            x-internal="natural-dimensions"} (e.g., it\'s a vector
            graphic with no specified content size) and either
            `options`{.variable}\'s
            [`resizeWidth`{#the-imagebitmap-interface:dom-imagebitmapoptions-resizewidth-2}](#dom-imagebitmapoptions-resizewidth)
            or `options`{.variable}\'s
            [`resizeHeight`{#the-imagebitmap-interface:dom-imagebitmapoptions-resizeheight-2}](#dom-imagebitmapoptions-resizeheight)
            is not present, then return [a promise rejected
            with](https://webidl.spec.whatwg.org/#a-promise-rejected-with){#the-imagebitmap-interface:a-promise-rejected-with-4
            x-internal="a-promise-rejected-with"} an
            [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-imagebitmap-interface:invalidstateerror-4
            x-internal="invalidstateerror"}
            [`DOMException`{#the-imagebitmap-interface:domexception-7}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

        2.  If `image`{.variable}\'s media data has no [natural
            dimensions](https://drafts.csswg.org/css-images/#natural-dimensions){#the-imagebitmap-interface:natural-dimensions-2
            x-internal="natural-dimensions"} (e.g., it\'s a vector
            graphic with no specified content size), it should be
            rendered to a bitmap of the size specified by the
            [`resizeWidth`{#the-imagebitmap-interface:dom-imagebitmapoptions-resizewidth-3}](#dom-imagebitmapoptions-resizewidth)
            and the
            [`resizeHeight`{#the-imagebitmap-interface:dom-imagebitmapoptions-resizeheight-3}](#dom-imagebitmapoptions-resizeheight)
            options.

        3.  Set `imageBitmap`{.variable}\'s [bitmap
            data](#concept-imagebitmap-bitmap-data){#the-imagebitmap-interface:concept-imagebitmap-bitmap-data-8}
            to a copy of `image`{.variable}\'s media data, [cropped to
            the source rectangle with
            formatting](#cropped-to-the-source-rectangle-with-formatting){#the-imagebitmap-interface:cropped-to-the-source-rectangle-with-formatting}.
            If this is an animated image, `imageBitmap`{.variable}\'s
            [bitmap
            data](#concept-imagebitmap-bitmap-data){#the-imagebitmap-interface:concept-imagebitmap-bitmap-data-9}
            must only be taken from the default image of the animation
            (the one that the format defines is to be used when
            animation is not supported or is disabled), or, if there is
            no such image, the first frame of the animation.

        4.  If `image`{.variable} [is not
            origin-clean](canvas.html#the-image-argument-is-not-origin-clean){#the-imagebitmap-interface:the-image-argument-is-not-origin-clean},
            then set the
            [origin-clean](canvas.html#concept-canvas-origin-clean){#the-imagebitmap-interface:concept-canvas-origin-clean-4}
            flag of `imageBitmap`{.variable}\'s bitmap to false.

        5.  [Queue a global
            task](webappapis.html#queue-a-global-task){#the-imagebitmap-interface:queue-a-global-task},
            using the [bitmap task
            source](#bitmap-task-source){#the-imagebitmap-interface:bitmap-task-source},
            to resolve `promise`{.variable} with
            `imageBitmap`{.variable}.

    [`video`{#the-imagebitmap-interface:the-video-element-3}](media.html#the-video-element)

    :   1.  If `image`{.variable}\'s
            [`networkState`{#the-imagebitmap-interface:dom-media-networkstate}](media.html#dom-media-networkstate)
            attribute is
            [`NETWORK_EMPTY`{#the-imagebitmap-interface:dom-media-network_empty}](media.html#dom-media-network_empty),
            then return [a promise rejected
            with](https://webidl.spec.whatwg.org/#a-promise-rejected-with){#the-imagebitmap-interface:a-promise-rejected-with-5
            x-internal="a-promise-rejected-with"} an
            [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-imagebitmap-interface:invalidstateerror-5
            x-internal="invalidstateerror"}
            [`DOMException`{#the-imagebitmap-interface:domexception-8}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

        2.  Set `imageBitmap`{.variable}\'s [bitmap
            data](#concept-imagebitmap-bitmap-data){#the-imagebitmap-interface:concept-imagebitmap-bitmap-data-10}
            to a copy of the frame at the [current playback
            position](media.html#current-playback-position){#the-imagebitmap-interface:current-playback-position},
            at the [media
            resource](media.html#media-resource){#the-imagebitmap-interface:media-resource}\'s
            [natural
            width](media.html#concept-video-intrinsic-width){#the-imagebitmap-interface:concept-video-intrinsic-width}
            and [natural
            height](media.html#concept-video-intrinsic-height){#the-imagebitmap-interface:concept-video-intrinsic-height}
            (i.e., after any aspect-ratio correction has been applied),
            [cropped to the source rectangle with
            formatting](#cropped-to-the-source-rectangle-with-formatting){#the-imagebitmap-interface:cropped-to-the-source-rectangle-with-formatting-2}.

        3.  If `image`{.variable} [is not
            origin-clean](canvas.html#the-image-argument-is-not-origin-clean){#the-imagebitmap-interface:the-image-argument-is-not-origin-clean-2},
            then set the
            [origin-clean](canvas.html#concept-canvas-origin-clean){#the-imagebitmap-interface:concept-canvas-origin-clean-5}
            flag of `imageBitmap`{.variable}\'s bitmap to false.

        4.  [Queue a global
            task](webappapis.html#queue-a-global-task){#the-imagebitmap-interface:queue-a-global-task-2},
            using the [bitmap task
            source](#bitmap-task-source){#the-imagebitmap-interface:bitmap-task-source-2},
            to resolve `promise`{.variable} with
            `imageBitmap`{.variable}.

    [`canvas`{#the-imagebitmap-interface:the-canvas-element-3}](canvas.html#the-canvas-element)

    :   1.  Set `imageBitmap`{.variable}\'s [bitmap
            data](#concept-imagebitmap-bitmap-data){#the-imagebitmap-interface:concept-imagebitmap-bitmap-data-11}
            to a copy of `image`{.variable}\'s [bitmap
            data](#concept-imagebitmap-bitmap-data){#the-imagebitmap-interface:concept-imagebitmap-bitmap-data-12},
            [cropped to the source rectangle with
            formatting](#cropped-to-the-source-rectangle-with-formatting){#the-imagebitmap-interface:cropped-to-the-source-rectangle-with-formatting-3}.

        2.  Set the
            [origin-clean](canvas.html#concept-canvas-origin-clean){#the-imagebitmap-interface:concept-canvas-origin-clean-6}
            flag of the `imageBitmap`{.variable}\'s bitmap to the same
            value as the
            [origin-clean](canvas.html#concept-canvas-origin-clean){#the-imagebitmap-interface:concept-canvas-origin-clean-7}
            flag of `image`{.variable}\'s bitmap.

        3.  [Queue a global
            task](webappapis.html#queue-a-global-task){#the-imagebitmap-interface:queue-a-global-task-3},
            using the [bitmap task
            source](#bitmap-task-source){#the-imagebitmap-interface:bitmap-task-source-3},
            to resolve `promise`{.variable} with
            `imageBitmap`{.variable}.

    [`Blob`{#the-imagebitmap-interface:blob-4}](https://w3c.github.io/FileAPI/#dfn-Blob){x-internal="blob"}

    :   Run these steps [in
        parallel](infrastructure.html#in-parallel){#the-imagebitmap-interface:in-parallel}:

        1.  Let `imageData`{.variable} be the result of reading
            `image`{.variable}\'s data. If an [error occurs during
            reading of the
            object](infrastructure.html#file-error-read){#the-imagebitmap-interface:file-error-read},
            then [queue a global
            task](webappapis.html#queue-a-global-task){#the-imagebitmap-interface:queue-a-global-task-4},
            using the [bitmap task
            source](#bitmap-task-source){#the-imagebitmap-interface:bitmap-task-source-4},
            to reject `promise`{.variable} with an
            [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-imagebitmap-interface:invalidstateerror-6
            x-internal="invalidstateerror"}
            [`DOMException`{#the-imagebitmap-interface:domexception-9}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
            and abort these steps.

        2.  Apply the [image sniffing
            rules](https://mimesniff.spec.whatwg.org/#rules-for-sniffing-images-specifically){#the-imagebitmap-interface:content-type-sniffing:-image
            x-internal="content-type-sniffing:-image"} to determine the
            file format of `imageData`{.variable}, with MIME type of
            `image`{.variable} (as given by `image`{.variable}\'s
            [`type`{#the-imagebitmap-interface:dom-blob-type}](https://w3c.github.io/FileAPI/#dfn-type){x-internal="dom-blob-type"}
            attribute) giving the official type.

        3.  If `imageData`{.variable} is not in a supported image file
            format (e.g., it\'s not an image at all), or if
            `imageData`{.variable} is corrupted in some fatal way such
            that the image dimensions cannot be obtained (e.g., a vector
            graphic with no natural size), then [queue a global
            task](webappapis.html#queue-a-global-task){#the-imagebitmap-interface:queue-a-global-task-5},
            using the [bitmap task
            source](#bitmap-task-source){#the-imagebitmap-interface:bitmap-task-source-5},
            to reject `promise`{.variable} with an
            [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-imagebitmap-interface:invalidstateerror-7
            x-internal="invalidstateerror"}
            [`DOMException`{#the-imagebitmap-interface:domexception-10}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
            and abort these steps.

        4.  Set `imageBitmap`{.variable}\'s [bitmap
            data](#concept-imagebitmap-bitmap-data){#the-imagebitmap-interface:concept-imagebitmap-bitmap-data-13}
            to `imageData`{.variable}, [cropped to the source rectangle
            with
            formatting](#cropped-to-the-source-rectangle-with-formatting){#the-imagebitmap-interface:cropped-to-the-source-rectangle-with-formatting-4}.
            If this is an animated image, `imageBitmap`{.variable}\'s
            [bitmap
            data](#concept-imagebitmap-bitmap-data){#the-imagebitmap-interface:concept-imagebitmap-bitmap-data-14}
            must only be taken from the default image of the animation
            (the one that the format defines is to be used when
            animation is not supported or is disabled), or, if there is
            no such image, the first frame of the animation.

        5.  [Queue a global
            task](webappapis.html#queue-a-global-task){#the-imagebitmap-interface:queue-a-global-task-6},
            using the [bitmap task
            source](#bitmap-task-source){#the-imagebitmap-interface:bitmap-task-source-6},
            to resolve `promise`{.variable} with
            `imageBitmap`{.variable}.

    [`ImageData`{#the-imagebitmap-interface:imagedata-4}](#imagedata)

    :   1.  Let `buffer`{.variable} be `image`{.variable}\'s
            [`data`{#the-imagebitmap-interface:dom-imagedata-data-2}](#dom-imagedata-data)
            attribute value\'s \[\[ViewedArrayBuffer\]\] internal slot.

        2.  If
            [IsDetachedBuffer](https://tc39.es/ecma262/#sec-isdetachedbuffer){#the-imagebitmap-interface:isdetachedbuffer
            x-internal="isdetachedbuffer"}(`buffer`{.variable}) is true,
            then return [a promise rejected
            with](https://webidl.spec.whatwg.org/#a-promise-rejected-with){#the-imagebitmap-interface:a-promise-rejected-with-6
            x-internal="a-promise-rejected-with"} an
            [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-imagebitmap-interface:invalidstateerror-8
            x-internal="invalidstateerror"}
            [`DOMException`{#the-imagebitmap-interface:domexception-11}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

        3.  Set `imageBitmap`{.variable}\'s [bitmap
            data](#concept-imagebitmap-bitmap-data){#the-imagebitmap-interface:concept-imagebitmap-bitmap-data-15}
            to `image`{.variable}\'s image data, [cropped to the source
            rectangle with
            formatting](#cropped-to-the-source-rectangle-with-formatting){#the-imagebitmap-interface:cropped-to-the-source-rectangle-with-formatting-5}.

        4.  [Queue a global
            task](webappapis.html#queue-a-global-task){#the-imagebitmap-interface:queue-a-global-task-7},
            using the [bitmap task
            source](#bitmap-task-source){#the-imagebitmap-interface:bitmap-task-source-7},
            to resolve `promise`{.variable} with
            `imageBitmap`{.variable}.

    [`ImageBitmap`{#the-imagebitmap-interface:imagebitmap-13}](#imagebitmap)

    :   1.  Set `imageBitmap`{.variable}\'s [bitmap
            data](#concept-imagebitmap-bitmap-data){#the-imagebitmap-interface:concept-imagebitmap-bitmap-data-16}
            to a copy of `image`{.variable}\'s [bitmap
            data](#concept-imagebitmap-bitmap-data){#the-imagebitmap-interface:concept-imagebitmap-bitmap-data-17},
            [cropped to the source rectangle with
            formatting](#cropped-to-the-source-rectangle-with-formatting){#the-imagebitmap-interface:cropped-to-the-source-rectangle-with-formatting-6}.

        2.  Set the
            [origin-clean](canvas.html#concept-canvas-origin-clean){#the-imagebitmap-interface:concept-canvas-origin-clean-8}
            flag of `imageBitmap`{.variable}\'s bitmap to the same value
            as the
            [origin-clean](canvas.html#concept-canvas-origin-clean){#the-imagebitmap-interface:concept-canvas-origin-clean-9}
            flag of `image`{.variable}\'s bitmap.

        3.  [Queue a global
            task](webappapis.html#queue-a-global-task){#the-imagebitmap-interface:queue-a-global-task-8},
            using the [bitmap task
            source](#bitmap-task-source){#the-imagebitmap-interface:bitmap-task-source-8},
            to resolve `promise`{.variable} with
            `imageBitmap`{.variable}.

    [`VideoFrame`{#the-imagebitmap-interface:videoframe}](https://w3c.github.io/webcodecs/#videoframe-interface){x-internal="videoframe"}

    :   1.  Set `imageBitmap`{.variable}\'s [bitmap
            data](#concept-imagebitmap-bitmap-data){#the-imagebitmap-interface:concept-imagebitmap-bitmap-data-18}
            to a copy of `image`{.variable}\'s visible pixel data,
            [cropped to the source rectangle with
            formatting](#cropped-to-the-source-rectangle-with-formatting){#the-imagebitmap-interface:cropped-to-the-source-rectangle-with-formatting-7}.

        2.  [Queue a global
            task](webappapis.html#queue-a-global-task){#the-imagebitmap-interface:queue-a-global-task-9},
            using the [bitmap task
            source](#bitmap-task-source){#the-imagebitmap-interface:bitmap-task-source-9},
            to resolve `promise`{.variable} with
            `imageBitmap`{.variable}.

7.  Return `promise`{.variable}.

When the steps above require that the user agent [crop bitmap data to
the source rectangle with
formatting]{#cropped-to-the-source-rectangle-with-formatting .dfn}, the
user agent must run the following steps:

1.  Let `input`{.variable} be the [bitmap
    data](#concept-imagebitmap-bitmap-data){#the-imagebitmap-interface:concept-imagebitmap-bitmap-data-19}
    being transformed.

2.  If `sx`{.variable}, `sy`{.variable}, `sw`{.variable} and
    `sh`{.variable} are specified, let `sourceRectangle`{.variable} be a
    rectangle whose corners are the four points (`sx`{.variable},
    `sy`{.variable}), (`sx`{.variable}+`sw`{.variable},
    `sy`{.variable}), (`sx`{.variable}+`sw`{.variable},
    `sy`{.variable}+`sh`{.variable}), (`sx`{.variable},
    `sy`{.variable}+`sh`{.variable}). Otherwise, let
    `sourceRectangle`{.variable} be a rectangle whose corners are the
    four points (0, 0), (width of `input`{.variable}, 0), (width of
    `input`{.variable}, height of `input`{.variable}), (0, height of
    `input`{.variable}).

    If either `sw`{.variable} or `sh`{.variable} are negative, then the
    top-left corner of this rectangle will be to the left or above the
    (`sx`{.variable}, `sy`{.variable}) point.

3.  Let `outputWidth`{.variable} be determined as follows:

    If the [`resizeWidth`{#the-imagebitmap-interface:dom-imagebitmapoptions-resizewidth-4}](#dom-imagebitmapoptions-resizewidth) member of `options`{.variable} is specified
    :   the value of the
        [`resizeWidth`{#the-imagebitmap-interface:dom-imagebitmapoptions-resizewidth-5}](#dom-imagebitmapoptions-resizewidth)
        member of `options`{.variable}

    If the [`resizeWidth`{#the-imagebitmap-interface:dom-imagebitmapoptions-resizewidth-6}](#dom-imagebitmapoptions-resizewidth) member of `options`{.variable} is not specified, but the [`resizeHeight`{#the-imagebitmap-interface:dom-imagebitmapoptions-resizeheight-4}](#dom-imagebitmapoptions-resizeheight) member is specified
    :   the width of `sourceRectangle`{.variable}, times the value of
        the
        [`resizeHeight`{#the-imagebitmap-interface:dom-imagebitmapoptions-resizeheight-5}](#dom-imagebitmapoptions-resizeheight)
        member of `options`{.variable}, divided by the height of
        `sourceRectangle`{.variable}, rounded up to the nearest integer

    If neither [`resizeWidth`{#the-imagebitmap-interface:dom-imagebitmapoptions-resizewidth-7}](#dom-imagebitmapoptions-resizewidth) nor [`resizeHeight`{#the-imagebitmap-interface:dom-imagebitmapoptions-resizeheight-6}](#dom-imagebitmapoptions-resizeheight) are specified
    :   the width of `sourceRectangle`{.variable}

4.  Let `outputHeight`{.variable} be determined as follows:

    If the [`resizeHeight`{#the-imagebitmap-interface:dom-imagebitmapoptions-resizeheight-7}](#dom-imagebitmapoptions-resizeheight) member of `options`{.variable} is specified
    :   the value of the
        [`resizeHeight`{#the-imagebitmap-interface:dom-imagebitmapoptions-resizeheight-8}](#dom-imagebitmapoptions-resizeheight)
        member of `options`{.variable}

    If the [`resizeHeight`{#the-imagebitmap-interface:dom-imagebitmapoptions-resizeheight-9}](#dom-imagebitmapoptions-resizeheight) member of `options`{.variable} is not specified, but the [`resizeWidth`{#the-imagebitmap-interface:dom-imagebitmapoptions-resizewidth-8}](#dom-imagebitmapoptions-resizewidth) member is specified
    :   the height of `sourceRectangle`{.variable}, times the value of
        the
        [`resizeWidth`{#the-imagebitmap-interface:dom-imagebitmapoptions-resizewidth-9}](#dom-imagebitmapoptions-resizewidth)
        member of `options`{.variable}, divided by the width of
        `sourceRectangle`{.variable}, rounded up to the nearest integer

    If neither [`resizeWidth`{#the-imagebitmap-interface:dom-imagebitmapoptions-resizewidth-10}](#dom-imagebitmapoptions-resizewidth) nor [`resizeHeight`{#the-imagebitmap-interface:dom-imagebitmapoptions-resizeheight-10}](#dom-imagebitmapoptions-resizeheight) are specified
    :   the height of `sourceRectangle`{.variable}

5.  Place `input`{.variable} on an infinite [transparent
    black](https://drafts.csswg.org/css-color/#transparent-black){#the-imagebitmap-interface:transparent-black-2
    x-internal="transparent-black"} grid plane, positioned so that its
    top left corner is at the origin of the plane, with the
    `x`{.variable}-coordinate increasing to the right, and the
    `y`{.variable}-coordinate increasing down, and with each pixel in
    the `input`{.variable} image data occupying a cell on the plane\'s
    grid.

6.  Let `output`{.variable} be the rectangle on the plane denoted by
    `sourceRectangle`{.variable}.

7.  Scale `output`{.variable} to the size specified by
    `outputWidth`{.variable} and `outputHeight`{.variable}. The user
    agent should use the value of the
    [`resizeQuality`]{#dom-imagebitmapoptions-resizequality .dfn
    dfn-for="ImageBitmapOptions" dfn-type="dict-member"} option to guide
    the choice of scaling algorithm.

8.  If the value of the
    [`imageOrientation`]{#dom-imagebitmapoptions-imageorientation .dfn
    dfn-for="ImageBitmapOptions" dfn-type="dict-member"} member of
    `options`{.variable} is \"[`flipY`]{#dom-imageorientation-flipy .dfn
    dfn-for="ImageOrientation" dfn-type="enum-value"}\",
    `output`{.variable} must be flipped vertically, disregarding any
    image orientation metadata of the source (such as EXIF metadata), if
    any. [\[EXIF\]](references.html#refsEXIF)

    If the value is \"[`from-image`]{#dom-imageorientation-from-image
    .dfn dfn-for="ImageOrientation" dfn-type="enum-value"}\", no extra
    step is needed.

    There used to be a
    \"[`none`]{#dom-imagebitmapoptions-imageorientation-none .dfn}\"
    enum value. It was renamed to
    \"[`from-image`{#the-imagebitmap-interface:dom-imageorientation-from-image-3}](#dom-imageorientation-from-image)\".
    In the future,
    \"[`none`{#the-imagebitmap-interface:dom-imagebitmapoptions-imageorientation-none}](#dom-imagebitmapoptions-imageorientation-none)\"
    will be added back with a different meaning.

9.  If `image`{.variable} is an
    [`img`{#the-imagebitmap-interface:the-img-element-4}](embedded-content.html#the-img-element)
    element or a
    [`Blob`{#the-imagebitmap-interface:blob-5}](https://w3c.github.io/FileAPI/#dfn-Blob){x-internal="blob"}
    object, let `val`{.variable} be the value of the
    [`colorSpaceConversion`]{#dom-imagebitmapoptions-colorspaceconversion
    .dfn dfn-for="ImageBitmapOptions" dfn-type="dict-member"} member of
    `options`{.variable}, and then run these substeps:

    1.  If `val`{.variable} is
        \"[`default`]{#dom-colorspaceconversion-default .dfn
        dfn-for="ColorSpaceConversion" dfn-type="enum-value"}\", the
        color space conversion behavior is implementation-specific, and
        should be chosen according to the default color space that the
        implementation uses for drawing images onto the canvas.

    2.  If `val`{.variable} is \"[`none`]{#dom-colorspaceconversion-none
        .dfn dfn-for="ColorSpaceConversion" dfn-type="enum-value"}\",
        `output`{.variable} must be decoded without performing any color
        space conversions. This means that the image decoding algorithm
        must ignore color profile metadata embedded in the source data
        as well as the display device color profile.

10. Let `val`{.variable} be the value of
    [`premultiplyAlpha`]{#dom-imagebitmapoptions-premultiplyalpha .dfn
    dfn-for="ImageBitmapOptions" dfn-type="dict-member"} member of
    `options`{.variable}, and then run these substeps:

    1.  If `val`{.variable} is
        \"[`default`]{#dom-premultiplyalpha-default .dfn
        dfn-for="PremultiplyAlpha" dfn-type="enum-value"}\", the alpha
        premultiplication behavior is implementation-specific, and
        should be chosen according to implementation deems optimal for
        drawing images onto the canvas.

    2.  If `val`{.variable} is
        \"[`premultiply`]{#dom-premultiplyalpha-premultiply .dfn
        dfn-for="PremultiplyAlpha" dfn-type="enum-value"}\", the
        `output`{.variable} that is not premultiplied by alpha must have
        its color components [multiplied by
        alpha](canvas.html#convert-to-premultiplied){#the-imagebitmap-interface:convert-to-premultiplied}
        and that is premultiplied by alpha must be left untouched.

    3.  If `val`{.variable} is \"[`none`]{#dom-premultiplyalpha-none
        .dfn dfn-for="PremultiplyAlpha" dfn-type="enum-value"}\", the
        `output`{.variable} that is not premultiplied by alpha must be
        left untouched and that is premultiplied by alpha must have its
        color components [divided by
        alpha](canvas.html#convert-from-premultiplied){#the-imagebitmap-interface:convert-from-premultiplied}.

11. Return `output`{.variable}.

The [`close()`]{#dom-imagebitmap-close .dfn dfn-for="ImageBitmap"
dfn-type="method"} method steps are:

1.  Set
    [this](https://webidl.spec.whatwg.org/#this){#the-imagebitmap-interface:this
    x-internal="this"}\'s
    [\[\[Detached\]\]](structured-data.html#detached){#the-imagebitmap-interface:detached-3}
    internal slot value to true.

2.  Unset
    [this](https://webidl.spec.whatwg.org/#this){#the-imagebitmap-interface:this-2
    x-internal="this"}\'s [bitmap
    data](#concept-imagebitmap-bitmap-data){#the-imagebitmap-interface:concept-imagebitmap-bitmap-data-20}.

The [`width`]{#dom-imagebitmap-width .dfn dfn-for="ImageBitmap"
dfn-type="attribute"} getter steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-imagebitmap-interface:this-3
    x-internal="this"}\'s
    [\[\[Detached\]\]](structured-data.html#detached){#the-imagebitmap-interface:detached-4}
    internal slot\'s value is true, then return 0.

2.  Return
    [this](https://webidl.spec.whatwg.org/#this){#the-imagebitmap-interface:this-4
    x-internal="this"}\'s width, in [CSS
    pixels](https://drafts.csswg.org/css-values/#px){#the-imagebitmap-interface:'px'-4
    x-internal="'px'"}.

The [`height`]{#dom-imagebitmap-height .dfn dfn-for="ImageBitmap"
dfn-type="attribute"} getter steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-imagebitmap-interface:this-5
    x-internal="this"}\'s
    [\[\[Detached\]\]](structured-data.html#detached){#the-imagebitmap-interface:detached-5}
    internal slot\'s value is true, then return 0.

2.  Return
    [this](https://webidl.spec.whatwg.org/#this){#the-imagebitmap-interface:this-6
    x-internal="this"}\'s height, in [CSS
    pixels](https://drafts.csswg.org/css-values/#px){#the-imagebitmap-interface:'px'-5
    x-internal="'px'"}.

The
[`ResizeQuality`{#the-imagebitmap-interface:resizequality-2}](#resizequality)
enumeration is used to express a preference for the interpolation
quality to use when scaling images.

The \"[`pixelated`]{#dom-resizequality-pixelated .dfn
dfn-for="ResizeQuality" dfn-type="enum-value"}\" value indicates a
preference for scaling the image to preserve the pixelation of the
original as much as possible, with minor smoothing as necessary to avoid
distorting the image when the target size is not a clean multiple of the
original.

To implement
\"[`pixelated`{#the-imagebitmap-interface:dom-resizequality-pixelated-2}](#dom-resizequality-pixelated)\",
for each axis independently, first determine the integer multiple of its
natural size that puts it closest to the target size and is greater than
zero. Scale it to this integer-multiple-size using nearest neighbor,
then scale it the rest of the way to the target size using bilinear
interpolation.

The \"[`low`]{#dom-resizequality-low .dfn dfn-for="ResizeQuality"
dfn-type="enum-value"}\" value indicates a preference for a low level of
image interpolation quality. Low-quality image interpolation may be more
computationally efficient than higher settings.

The \"[`medium`]{#dom-resizequality-medium .dfn dfn-for="ResizeQuality"
dfn-type="enum-value"}\" value indicates a preference for a medium level
of image interpolation quality.

The \"[`high`]{#dom-resizequality-high .dfn dfn-for="ResizeQuality"
dfn-type="enum-value"}\" value indicates a preference for a high level
of image interpolation quality. High-quality image interpolation may be
more computationally expensive than lower settings.

Bilinear scaling is an example of a relatively fast, lower-quality
image-smoothing algorithm. Bicubic or Lanczos scaling are examples of
image-scaling algorithms that produce higher-quality output. This
specification does not mandate that specific interpolation algorithms be
used, except for
\"[`pixelated`{#the-imagebitmap-interface:dom-resizequality-pixelated-3}](#dom-resizequality-pixelated)\"
as described above.

::: example
Using this API, a sprite sheet can be precut and prepared:

``` js
var sprites = {};
function loadMySprites() {
  var image = new Image();
  image.src = 'mysprites.png';
  var resolver;
  var promise = new Promise(function (arg) { resolver = arg });
  image.onload = function () {
    resolver(Promise.all([
      createImageBitmap(image,  0,  0, 40, 40).then(function (image) { sprites.person = image }),
      createImageBitmap(image, 40,  0, 40, 40).then(function (image) { sprites.grass  = image }),
      createImageBitmap(image, 80,  0, 40, 40).then(function (image) { sprites.tree   = image }),
      createImageBitmap(image,  0, 40, 40, 40).then(function (image) { sprites.hut    = image }),
      createImageBitmap(image, 40, 40, 40, 40).then(function (image) { sprites.apple  = image }),
      createImageBitmap(image, 80, 40, 40, 40).then(function (image) { sprites.snake  = image })
    ]));
  };
  return promise;
}

function runDemo() {
  var canvas = document.querySelector('canvas#demo');
  var context = canvas.getContext('2d');
  context.drawImage(sprites.tree, 30, 10);
  context.drawImage(sprites.snake, 70, 10);
}

loadMySprites().then(runDemo);
```
:::

### [8.11]{.secno} Animation frames[](#animation-frames){.self-link}

Some objects include the
[`AnimationFrameProvider`{#animation-frames:animationframeprovider}](#animationframeprovider)
interface mixin.

``` idl
callback FrameRequestCallback = undefined (DOMHighResTimeStamp time);

interface mixin AnimationFrameProvider {
  unsigned long requestAnimationFrame(FrameRequestCallback callback);
  undefined cancelAnimationFrame(unsigned long handle);
};
Window includes AnimationFrameProvider;
DedicatedWorkerGlobalScope includes AnimationFrameProvider;
```

Each
[`AnimationFrameProvider`{#animation-frames:animationframeprovider-4}](#animationframeprovider)
object also has a [target
object]{#concept-animationframeprovider-target-object .dfn} that stores
the provider\'s internal state. It is defined as follows:

If the [`AnimationFrameProvider`{#animation-frames:animationframeprovider-5}](#animationframeprovider) is a [`Window`{#animation-frames:window-2}](nav-history-apis.html#window)
:   The
    [`Window`{#animation-frames:window-3}](nav-history-apis.html#window)\'s
    [associated
    `Document`](nav-history-apis.html#concept-document-window){#animation-frames:concept-document-window}

If the [`AnimationFrameProvider`{#animation-frames:animationframeprovider-6}](#animationframeprovider) is a [`DedicatedWorkerGlobalScope`{#animation-frames:dedicatedworkerglobalscope-2}](workers.html#dedicatedworkerglobalscope)
:   The
    [`DedicatedWorkerGlobalScope`{#animation-frames:dedicatedworkerglobalscope-3}](workers.html#dedicatedworkerglobalscope)

Each [target
object](#concept-animationframeprovider-target-object){#animation-frames:concept-animationframeprovider-target-object}
has a [map of animation frame
callbacks]{#list-of-animation-frame-callbacks .dfn}, which is an
[ordered
map](https://infra.spec.whatwg.org/#ordered-map){#animation-frames:ordered-map
x-internal="ordered-map"} that must be initially empty, and an
[animation frame callback
identifier]{#animation-frame-callback-identifier .dfn}, which is a
number that must initially be zero.

An
[`AnimationFrameProvider`{#animation-frames:animationframeprovider-7}](#animationframeprovider)
`provider`{.variable} is considered
[supported]{#concept-animationframeprovider-supported .dfn} if any of
the following are true:

- `provider`{.variable} is a
  [`Window`{#animation-frames:window-4}](nav-history-apis.html#window).

- `provider`{.variable}\'s [owner
  set](workers.html#concept-WorkerGlobalScope-owner-set){#animation-frames:concept-WorkerGlobalScope-owner-set}
  [contains](https://infra.spec.whatwg.org/#list-contain){#animation-frames:list-contains
  x-internal="list-contains"} a
  [`Document`{#animation-frames:document}](dom.html#document) object.

- Any of the
  [`DedicatedWorkerGlobalScope`{#animation-frames:dedicatedworkerglobalscope-4}](workers.html#dedicatedworkerglobalscope)
  objects in `provider`{.variable}\'s [owner
  set](workers.html#concept-WorkerGlobalScope-owner-set){#animation-frames:concept-WorkerGlobalScope-owner-set-2}
  are
  [supported](#concept-animationframeprovider-supported){#animation-frames:concept-animationframeprovider-supported}.

------------------------------------------------------------------------

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Window/requestAnimationFrame](https://developer.mozilla.org/en-US/docs/Web/API/Window/requestAnimationFrame "The window.requestAnimationFrame() method tells the browser that you wish to perform an animation and requests that the browser calls a specified function to update an animation right before the next repaint. The method takes a callback as an argument to be invoked before the repaint.")

Support in all current engines.

::: support
[Firefox23+]{.firefox .yes}[Safari7+]{.safari .yes}[Chrome24+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android23+]{.firefox_android .yes}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android4.4+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

The
[`requestAnimationFrame(``callback`{.variable}`)`]{#dom-animationframeprovider-requestanimationframe
.dfn dfn-for="AnimationFrameProvider" dfn-type="method"} method steps
are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#animation-frames:this
    x-internal="this"} is not
    [supported](#concept-animationframeprovider-supported){#animation-frames:concept-animationframeprovider-supported-2},
    then throw a
    [\"`NotSupportedError`\"](https://webidl.spec.whatwg.org/#notsupportederror){#animation-frames:notsupportederror
    x-internal="notsupportederror"}
    [`DOMException`{#animation-frames:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  Let `target`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#animation-frames:this-2
    x-internal="this"}\'s [target
    object](#concept-animationframeprovider-target-object){#animation-frames:concept-animationframeprovider-target-object-2}.

3.  Increment `target`{.variable}\'s [animation frame callback
    identifier](#animation-frame-callback-identifier){#animation-frames:animation-frame-callback-identifier}
    by one, and let `handle`{.variable} be the result.

4.  Let `callbacks`{.variable} be `target`{.variable}\'s [map of
    animation frame
    callbacks](#list-of-animation-frame-callbacks){#animation-frames:list-of-animation-frame-callbacks}.

5.  [Set](https://infra.spec.whatwg.org/#map-set){#animation-frames:map-set
    x-internal="map-set"} `callbacks`{.variable}\[`handle`{.variable}\]
    to `callback`{.variable}.

6.  Return `handle`{.variable}.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Window/cancelAnimationFrame](https://developer.mozilla.org/en-US/docs/Web/API/Window/cancelAnimationFrame "The window.cancelAnimationFrame() method cancels an animation frame request previously scheduled through a call to window.requestAnimationFrame().")

Support in all current engines.

::: support
[Firefox23+]{.firefox .yes}[Safari7+]{.safari .yes}[Chrome24+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

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
[`cancelAnimationFrame(``handle`{.variable}`)`]{#animationframeprovider-cancelanimationframe
.dfn dfn-for="AnimationFrameProvider" dfn-type="method"} method steps
are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#animation-frames:this-3
    x-internal="this"} is not
    [supported](#concept-animationframeprovider-supported){#animation-frames:concept-animationframeprovider-supported-3},
    then throw a
    [\"`NotSupportedError`\"](https://webidl.spec.whatwg.org/#notsupportederror){#animation-frames:notsupportederror-2
    x-internal="notsupportederror"}
    [`DOMException`{#animation-frames:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  Let `callbacks`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#animation-frames:this-4
    x-internal="this"}\'s [target
    object](#concept-animationframeprovider-target-object){#animation-frames:concept-animationframeprovider-target-object-3}\'s
    [map of animation frame
    callbacks](#list-of-animation-frame-callbacks){#animation-frames:list-of-animation-frame-callbacks-2}.

3.  [Remove](https://infra.spec.whatwg.org/#map-remove){#animation-frames:map-remove
    x-internal="map-remove"}
    `callbacks`{.variable}\[`handle`{.variable}\].

To [run the animation frame
callbacks]{#run-the-animation-frame-callbacks .dfn} for a [target
object](#concept-animationframeprovider-target-object){#animation-frames:concept-animationframeprovider-target-object-4}
`target`{.variable} with a timestamp `now`{.variable}:

1.  Let `callbacks`{.variable} be `target`{.variable}\'s [map of
    animation frame
    callbacks](#list-of-animation-frame-callbacks){#animation-frames:list-of-animation-frame-callbacks-3}.

2.  Let `callbackHandles`{.variable} be the result of [getting the
    keys](https://infra.spec.whatwg.org/#map-getting-the-keys){#animation-frames:map-get-the-keys
    x-internal="map-get-the-keys"} of `callbacks`{.variable}.

3.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#animation-frames:list-iterate
    x-internal="list-iterate"} `handle`{.variable} in
    `callbackHandles`{.variable}, if `handle`{.variable}
    [exists](https://infra.spec.whatwg.org/#map-exists){#animation-frames:map-exists
    x-internal="map-exists"} in `callbacks`{.variable}:

    1.  Let `callback`{.variable} be
        `callbacks`{.variable}\[`handle`{.variable}\].

    2.  [Remove](https://infra.spec.whatwg.org/#map-remove){#animation-frames:map-remove-2
        x-internal="map-remove"}
        `callbacks`{.variable}\[`handle`{.variable}\].

    3.  [Invoke](https://webidl.spec.whatwg.org/#invoke-a-callback-function){#animation-frames:es-invoking-callback-functions
        x-internal="es-invoking-callback-functions"}
        `callback`{.variable} with « `now`{.variable} » and
        \"`report`\".

::: example
Inside workers,
[`requestAnimationFrame()`{#animation-frames:dom-animationframeprovider-requestanimationframe-2}](#dom-animationframeprovider-requestanimationframe)
can be used together with an
[`OffscreenCanvas`{#animation-frames:offscreencanvas}](canvas.html#offscreencanvas)
transferred from a
[`canvas`{#animation-frames:the-canvas-element}](canvas.html#the-canvas-element)
element. First, in the document, transfer control to the worker:

``` js
const offscreenCanvas = document.getElementById("c").transferControlToOffscreen();
worker.postMessage(offscreenCanvas, [offscreenCanvas]);
```

Then, in the worker, the following code will draw a rectangle moving
from left to right:

``` js
let ctx, pos = 0;
function draw(dt) {
  ctx.clearRect(0, 0, 100, 100);
  ctx.fillRect(pos, 0, 10, 10);
  pos += 10 * dt;
  requestAnimationFrame(draw);
}

self.onmessage = function(ev) {
  const transferredCanvas = ev.data;
  ctx = transferredCanvas.getContext("2d");
  draw();
};
```
:::

[← 8.9 System state and capabilities](system-state.html) --- [Table of
Contents](./) --- [9 Communication →](comms.html)
