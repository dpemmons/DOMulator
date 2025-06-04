HTML Standard

::: {#head .head .with-buttons}
[![WHATWG](https://resources.whatwg.org/logo.svg){.darkmode-aware
width="100" crossorigin="" height="100"}](https://whatwg.org/){.logo}

# HTML {#html .allcaps}

Living Standard --- Last Updated [2 June 2025]{.pubdate}
:::

[← 4.8.5 The iframe element](iframe-embed-object.html) --- [Table of
Contents](./) --- [4.8.12 The map element →](image-maps.html)

1.  ::: {#toc-semantics}
    1.  1.  [[4.8.8]{.secno} The `video`
            element](media.html#the-video-element)
        2.  [[4.8.9]{.secno} The `audio`
            element](media.html#the-audio-element)
        3.  [[4.8.10]{.secno} The `track`
            element](media.html#the-track-element)
        4.  [[4.8.11]{.secno} Media elements](media.html#media-elements)
            1.  [[4.8.11.1]{.secno} Error codes](media.html#error-codes)
            2.  [[4.8.11.2]{.secno} Location of the media
                resource](media.html#location-of-the-media-resource)
            3.  [[4.8.11.3]{.secno} MIME types](media.html#mime-types)
            4.  [[4.8.11.4]{.secno} Network
                states](media.html#network-states)
            5.  [[4.8.11.5]{.secno} Loading the media
                resource](media.html#loading-the-media-resource)
            6.  [[4.8.11.6]{.secno} Offsets into the media
                resource](media.html#offsets-into-the-media-resource)
            7.  [[4.8.11.7]{.secno} Ready
                states](media.html#ready-states)
            8.  [[4.8.11.8]{.secno} Playing the media
                resource](media.html#playing-the-media-resource)
            9.  [[4.8.11.9]{.secno} Seeking](media.html#seeking)
            10. [[4.8.11.10]{.secno} Media resources with multiple media
                tracks](media.html#media-resources-with-multiple-media-tracks)
                1.  [[4.8.11.10.1]{.secno} `AudioTrackList` and
                    `VideoTrackList`
                    objects](media.html#audiotracklist-and-videotracklist-objects)
                2.  [[4.8.11.10.2]{.secno} Selecting specific audio and
                    video tracks
                    declaratively](media.html#selecting-specific-audio-and-video-tracks-declaratively)
            11. [[4.8.11.11]{.secno} Timed text
                tracks](media.html#timed-text-tracks)
                1.  [[4.8.11.11.1]{.secno} Text track
                    model](media.html#text-track-model)
                2.  [[4.8.11.11.2]{.secno} Sourcing in-band text
                    tracks](media.html#sourcing-in-band-text-tracks)
                3.  [[4.8.11.11.3]{.secno} Sourcing out-of-band text
                    tracks](media.html#sourcing-out-of-band-text-tracks)
                4.  [[4.8.11.11.4]{.secno} Guidelines for exposing cues
                    in various formats as text track
                    cues](media.html#guidelines-for-exposing-cues-in-various-formats-as-text-track-cues)
                5.  [[4.8.11.11.5]{.secno} Text track
                    API](media.html#text-track-api)
                6.  [[4.8.11.11.6]{.secno} Event handlers for objects of
                    the text track APIs](media.html#cue-events)
                7.  [[4.8.11.11.7]{.secno} Best practices for metadata
                    text
                    tracks](media.html#best-practices-for-metadata-text-tracks)
            12. [[4.8.11.12]{.secno} Identifying a track kind through a
                URL](media.html#identifying-a-track-kind-through-a-url)
            13. [[4.8.11.13]{.secno} User
                interface](media.html#user-interface)
            14. [[4.8.11.14]{.secno} Time
                ranges](media.html#time-ranges)
            15. [[4.8.11.15]{.secno} The `TrackEvent`
                interface](media.html#the-trackevent-interface)
            16. [[4.8.11.16]{.secno} Events
                summary](media.html#mediaevents)
            17. [[4.8.11.17]{.secno} Security and privacy
                considerations](media.html#security-and-privacy-considerations)
            18. [[4.8.11.18]{.secno} Best practices for authors using
                media
                elements](media.html#best-practices-for-authors-using-media-elements)
            19. [[4.8.11.19]{.secno} Best practices for implementers of
                media
                elements](media.html#best-practices-for-implementers-of-media-elements)
    :::

#### [4.8.8]{.secno} The [`video`]{#video .dfn dfn-type="element"} element[](#the-video-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/video](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/video "The <video> HTML element embeds a media player which supports video playback into the document. You can use <video> for audio content as well, but the <audio> element may provide a more appropriate user experience.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLVideoElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLVideoElement "Implemented by the <video> element, the HTMLVideoElement interface provides special properties and methods for manipulating video objects. It also inherits properties and methods of HTMLMediaElement and HTMLElement.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

[Categories](dom.html#concept-element-categories){#the-video-element:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-video-element:flow-content-2}.
:   [Phrasing
    content](dom.html#phrasing-content-2){#the-video-element:phrasing-content-2}.
:   [Embedded
    content](dom.html#embedded-content-category){#the-video-element:embedded-content-category}.
:   If the element has a
    [`controls`{#the-video-element:attr-media-controls}](#attr-media-controls)
    attribute: [Interactive
    content](dom.html#interactive-content-2){#the-video-element:interactive-content-2}.
:   [Palpable
    content](dom.html#palpable-content-2){#the-video-element:palpable-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-video-element:concept-element-contexts}:
:   Where [embedded
    content](dom.html#embedded-content-category){#the-video-element:embedded-content-category-2}
    is expected.

[Content model](dom.html#concept-element-content-model){#the-video-element:concept-element-content-model}:
:   If the element has a
    [`src`{#the-video-element:attr-media-src}](#attr-media-src)
    attribute: zero or more
    [`track`{#the-video-element:the-track-element}](#the-track-element)
    elements, then
    [transparent](dom.html#transparent){#the-video-element:transparent},
    but with no [media
    element](#media-element){#the-video-element:media-element}
    descendants.
:   If the element does not have a
    [`src`{#the-video-element:attr-media-src-2}](#attr-media-src)
    attribute: zero or more
    [`source`{#the-video-element:the-source-element}](embedded-content.html#the-source-element)
    elements, then zero or more
    [`track`{#the-video-element:the-track-element-2}](#the-track-element)
    elements, then
    [transparent](dom.html#transparent){#the-video-element:transparent-2},
    but with no [media
    element](#media-element){#the-video-element:media-element-2}
    descendants.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-video-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-video-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-video-element:global-attributes}
:   [`src`{#the-video-element:attr-media-src-3}](#attr-media-src) ---
    Address of the resource
:   [`crossorigin`{#the-video-element:attr-media-crossorigin}](#attr-media-crossorigin)
    --- How the element handles crossorigin requests
:   [`poster`{#the-video-element:attr-video-poster}](#attr-video-poster)
    --- Poster frame to show prior to video playback
:   [`preload`{#the-video-element:attr-media-preload}](#attr-media-preload)
    --- Hints how much buffering the [media
    resource](#media-resource){#the-video-element:media-resource} will
    likely need
:   [`autoplay`{#the-video-element:attr-media-autoplay}](#attr-media-autoplay)
    --- Hint that the [media
    resource](#media-resource){#the-video-element:media-resource-2} can
    be started automatically when the page is loaded
:   [`playsinline`{#the-video-element:attr-video-playsinline}](#attr-video-playsinline)
    --- Encourage the user agent to display video content within the
    element\'s playback area
:   [`loop`{#the-video-element:attr-media-loop}](#attr-media-loop) ---
    Whether to loop the [media
    resource](#media-resource){#the-video-element:media-resource-3}
:   [`muted`{#the-video-element:attr-media-muted}](#attr-media-muted)
    --- Whether to mute the [media
    resource](#media-resource){#the-video-element:media-resource-4} by
    default
:   [`controls`{#the-video-element:attr-media-controls-2}](#attr-media-controls)
    --- Show user agent controls
:   [`width`{#the-video-element:attr-dim-width}](embedded-content-other.html#attr-dim-width)
    --- Horizontal dimension
:   [`height`{#the-video-element:attr-dim-height}](embedded-content-other.html#attr-dim-height)
    --- Vertical dimension

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-video-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-video).
:   [For implementers](https://w3c.github.io/html-aam/#el-video).

[DOM interface](dom.html#concept-element-dom){#the-video-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLVideoElement : HTMLMediaElement {
      [HTMLConstructor] constructor();

      [CEReactions] attribute unsigned long width;
      [CEReactions] attribute unsigned long height;
      readonly attribute unsigned long videoWidth;
      readonly attribute unsigned long videoHeight;
      [CEReactions] attribute USVString poster;
      [CEReactions] attribute boolean playsInline;
    };
    ```

A [`video`{#the-video-element:the-video-element}](#the-video-element)
element is used for playing videos or movies, and audio files with
captions.

Content may be provided inside the
[`video`{#the-video-element:the-video-element-2}](#the-video-element)
element. User agents should not show this content to the user; it is
intended for older web browsers which do not support
[`video`{#the-video-element:the-video-element-3}](#the-video-element),
so that text can be shown to the users of these older browsers informing
them of how to access the video contents.

In particular, this content is not intended to address accessibility
concerns. To make video content accessible to the partially sighted, the
blind, the hard-of-hearing, the deaf, and those with other physical or
cognitive disabilities, a variety of features are available. Captions
can be provided, either embedded in the video stream or as external
files using the
[`track`{#the-video-element:the-track-element-3}](#the-track-element)
element. Sign-language tracks can be embedded in the video stream. Audio
descriptions can be embedded in the video stream or in text form using a
[WebVTT
file](https://w3c.github.io/webvtt/#webvtt-file){#the-video-element:webvtt-file
x-internal="webvtt-file"} referenced using the
[`track`{#the-video-element:the-track-element-4}](#the-track-element)
element and synthesized into speech by the user agent. WebVTT can also
be used to provide chapter titles. For users who would rather not use a
media element at all, transcripts or other textual alternatives can be
provided by simply linking to them in the prose near the
[`video`{#the-video-element:the-video-element-4}](#the-video-element)
element. [\[WEBVTT\]](references.html#refsWEBVTT)

The
[`video`{#the-video-element:the-video-element-5}](#the-video-element)
element is a [media
element](#media-element){#the-video-element:media-element-3} whose
[media data](#media-data){#the-video-element:media-data} is ostensibly
video data, possibly with associated audio data.

The [`src`{#the-video-element:attr-media-src-4}](#attr-media-src),
[`crossorigin`{#the-video-element:attr-media-crossorigin-2}](#attr-media-crossorigin),
[`preload`{#the-video-element:attr-media-preload-2}](#attr-media-preload),
[`autoplay`{#the-video-element:attr-media-autoplay-2}](#attr-media-autoplay),
[`loop`{#the-video-element:attr-media-loop-2}](#attr-media-loop),
[`muted`{#the-video-element:attr-media-muted-2}](#attr-media-muted), and
[`controls`{#the-video-element:attr-media-controls-3}](#attr-media-controls)
attributes are [the attributes common to all media
elements](#media-element-attributes){#the-video-element:media-element-attributes}.

The [`poster`]{#attr-video-poster .dfn dfn-for="video"
dfn-type="element-attr"} attribute gives the
[URL](https://url.spec.whatwg.org/#concept-url){#the-video-element:url
x-internal="url"} of an image file that the user agent can show while no
video data is available. The attribute, if present, must contain a
[valid non-empty URL potentially surrounded by
spaces](urls-and-fetching.html#valid-non-empty-url-potentially-surrounded-by-spaces){#the-video-element:valid-non-empty-url-potentially-surrounded-by-spaces}.

If the specified resource is to be used, then, when the element is
created or when the
[`poster`{#the-video-element:attr-video-poster-2}](#attr-video-poster)
attribute is set, changed, or removed, the user agent must run the
following steps to determine the element\'s [poster frame]{#poster-frame
.dfn} (regardless of the value of the element\'s [show poster
flag](#show-poster-flag){#the-video-element:show-poster-flag}):

1.  If there is an existing instance of this algorithm running for this
    [`video`{#the-video-element:the-video-element-6}](#the-video-element)
    element, abort that instance of this algorithm without changing the
    [poster frame](#poster-frame){#the-video-element:poster-frame}.

2.  If the
    [`poster`{#the-video-element:attr-video-poster-3}](#attr-video-poster)
    attribute\'s value is the empty string or if the attribute is
    absent, then there is no [poster
    frame](#poster-frame){#the-video-element:poster-frame-2}; return.

3.  Let `url`{.variable} be the result of [encoding-parsing a
    URL](urls-and-fetching.html#encoding-parsing-a-url){#the-video-element:encoding-parsing-a-url}
    given the
    [`poster`{#the-video-element:attr-video-poster-4}](#attr-video-poster)
    attribute\'s value, relative to the element\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-video-element:node-document
    x-internal="node-document"}.

4.  If `url`{.variable} is failure, then return. [There is no [poster
    frame](#poster-frame){#the-video-element:poster-frame-3}.]{.note}

5.  Let `request`{.variable} be a new
    [request](https://fetch.spec.whatwg.org/#concept-request){#the-video-element:concept-request
    x-internal="concept-request"} whose
    [URL](https://fetch.spec.whatwg.org/#concept-request-url){#the-video-element:concept-request-url
    x-internal="concept-request-url"} is `url`{.variable},
    [client](https://fetch.spec.whatwg.org/#concept-request-client){#the-video-element:concept-request-client
    x-internal="concept-request-client"} is the element\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-video-element:node-document-2
    x-internal="node-document"}\'s [relevant settings
    object](webappapis.html#relevant-settings-object){#the-video-element:relevant-settings-object},
    [destination](https://fetch.spec.whatwg.org/#concept-request-destination){#the-video-element:concept-request-destination
    x-internal="concept-request-destination"} is \"`image`\", [initiator
    type](https://fetch.spec.whatwg.org/#request-initiator-type){#the-video-element:concept-request-initiator-type
    x-internal="concept-request-initiator-type"} is \"`video`\",
    [credentials
    mode](https://fetch.spec.whatwg.org/#concept-request-credentials-mode){#the-video-element:concept-request-credentials-mode
    x-internal="concept-request-credentials-mode"} is \"`include`\", and
    whose [use-URL-credentials
    flag](https://fetch.spec.whatwg.org/#concept-request-use-url-credentials-flag){#the-video-element:use-url-credentials-flag
    x-internal="use-url-credentials-flag"} is set.

6.  [Fetch](https://fetch.spec.whatwg.org/#concept-fetch){#the-video-element:concept-fetch
    x-internal="concept-fetch"} `request`{.variable}. This must [delay
    the load
    event](parsing.html#delay-the-load-event){#the-video-element:delay-the-load-event}
    of the element\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-video-element:node-document-3
    x-internal="node-document"}.

7.  If an image is thus obtained, the [poster
    frame](#poster-frame){#the-video-element:poster-frame-4} is that
    image. Otherwise, there is no [poster
    frame](#poster-frame){#the-video-element:poster-frame-5}.

The image given by the
[`poster`{#the-video-element:attr-video-poster-5}](#attr-video-poster)
attribute, the *[poster frame](#poster-frame)*, is intended to be a
representative frame of the video (typically one of the first non-blank
frames) that gives the user an idea of what the video is like.

The [`playsinline`]{#attr-video-playsinline .dfn dfn-for="video"
dfn-type="element-attr"} attribute is a [boolean
attribute](common-microsyntaxes.html#boolean-attribute){#the-video-element:boolean-attribute}.
If present, it serves as a hint to the user agent that the video ought
to be displayed \"inline\" in the document by default, constrained to
the element\'s playback area, instead of being displayed fullscreen or
in an independent resizable window.

The absence of the
[`playsinline`{#the-video-element:attr-video-playsinline-2}](#attr-video-playsinline)
attribute does not imply that the video will display fullscreen by
default. Indeed, most user agents have chosen to play all videos inline
by default, and in such user agents the
[`playsinline`{#the-video-element:attr-video-playsinline-3}](#attr-video-playsinline)
attribute has no effect.

------------------------------------------------------------------------

A [`video`{#the-video-element:the-video-element-7}](#the-video-element)
element represents what is given for the first matching condition in the
list below:

When no video data is available (the element\'s [`readyState`{#the-video-element:dom-media-readystate}](#dom-media-readystate) attribute is either [`HAVE_NOTHING`{#the-video-element:dom-media-have_nothing}](#dom-media-have_nothing), or [`HAVE_METADATA`{#the-video-element:dom-media-have_metadata}](#dom-media-have_metadata) but no video data has yet been obtained at all, or the element\'s [`readyState`{#the-video-element:dom-media-readystate-2}](#dom-media-readystate) attribute is any subsequent value but the [media resource](#media-resource){#the-video-element:media-resource-5} does not have a video channel)
:   The
    [`video`{#the-video-element:the-video-element-8}](#the-video-element)
    element
    [represents](dom.html#represents){#the-video-element:represents} its
    [poster frame](#poster-frame){#the-video-element:poster-frame-7}, if
    any, or else [transparent
    black](https://drafts.csswg.org/css-color/#transparent-black){#the-video-element:transparent-black
    x-internal="transparent-black"} with no [natural
    dimensions](https://drafts.csswg.org/css-images/#natural-dimensions){#the-video-element:natural-dimensions
    x-internal="natural-dimensions"}.

When the [`video`{#the-video-element:the-video-element-9}](#the-video-element) element is [paused](#dom-media-paused){#the-video-element:dom-media-paused}, the [current playback position](#current-playback-position){#the-video-element:current-playback-position} is the first frame of video, and the element\'s [show poster flag](#show-poster-flag){#the-video-element:show-poster-flag-2} is set
:   The
    [`video`{#the-video-element:the-video-element-10}](#the-video-element)
    element
    [represents](dom.html#represents){#the-video-element:represents-2}
    its [poster
    frame](#poster-frame){#the-video-element:poster-frame-8}, if any, or
    else the first frame of the video.

When the [`video`{#the-video-element:the-video-element-11}](#the-video-element) element is [paused](#dom-media-paused){#the-video-element:dom-media-paused-2}, and the frame of video corresponding to the [current playback position](#current-playback-position){#the-video-element:current-playback-position-2} is not available (e.g. because the video is seeking or buffering)\
When the [`video`{#the-video-element:the-video-element-12}](#the-video-element) element is neither [potentially playing](#potentially-playing){#the-video-element:potentially-playing} nor [paused](#dom-media-paused){#the-video-element:dom-media-paused-3} (e.g. when seeking or stalled)
:   The
    [`video`{#the-video-element:the-video-element-13}](#the-video-element)
    element
    [represents](dom.html#represents){#the-video-element:represents-3}
    the last frame of the video to have been rendered.

When the [`video`{#the-video-element:the-video-element-14}](#the-video-element) element is [paused](#dom-media-paused){#the-video-element:dom-media-paused-4}
:   The
    [`video`{#the-video-element:the-video-element-15}](#the-video-element)
    element
    [represents](dom.html#represents){#the-video-element:represents-4}
    the frame of video corresponding to the [current playback
    position](#current-playback-position){#the-video-element:current-playback-position-3}.

Otherwise (the [`video`{#the-video-element:the-video-element-16}](#the-video-element) element has a video channel and is [potentially playing](#potentially-playing){#the-video-element:potentially-playing-2})
:   The
    [`video`{#the-video-element:the-video-element-17}](#the-video-element)
    element
    [represents](dom.html#represents){#the-video-element:represents-5}
    the frame of video at the continuously increasing [\"current\"
    position](#current-playback-position){#the-video-element:current-playback-position-4}.
    When the [current playback
    position](#current-playback-position){#the-video-element:current-playback-position-5}
    changes such that the last frame rendered is no longer the frame
    corresponding to the [current playback
    position](#current-playback-position){#the-video-element:current-playback-position-6}
    in the video, the new frame must be rendered.

Frames of video must be obtained from the video track that was
[selected](#dom-videotrack-selected){#the-video-element:dom-videotrack-selected}
when the [event
loop](webappapis.html#event-loop){#the-video-element:event-loop} last
reached [step 1](webappapis.html#step1).

Which frame in a video stream corresponds to a particular playback
position is defined by the video stream\'s format.

The
[`video`{#the-video-element:the-video-element-18}](#the-video-element)
element also
[represents](dom.html#represents){#the-video-element:represents-6} any
[text track cues](#text-track-cue){#the-video-element:text-track-cue}
whose [text track cue active
flag](#text-track-cue-active-flag){#the-video-element:text-track-cue-active-flag}
is set and whose [text
track](#text-track){#the-video-element:text-track} is in the
[showing](#text-track-showing){#the-video-element:text-track-showing}
mode, and any audio from the [media
resource](#media-resource){#the-video-element:media-resource-6}, at the
[current playback
position](#current-playback-position){#the-video-element:current-playback-position-7}.

Any audio associated with the [media
resource](#media-resource){#the-video-element:media-resource-7} must, if
played, be played synchronized with the [current playback
position](#current-playback-position){#the-video-element:current-playback-position-8},
at the element\'s [effective media
volume](#effective-media-volume){#the-video-element:effective-media-volume}.
The user agent must play the audio from audio tracks that were
[enabled](#dom-audiotrack-enabled){#the-video-element:dom-audiotrack-enabled}
when the [event
loop](webappapis.html#event-loop){#the-video-element:event-loop-2} last
reached step 1.

In addition to the above, the user agent may provide messages to the
user (such as \"buffering\", \"no video loaded\", \"error\", or more
detailed information) by overlaying text or icons on the video or other
areas of the element\'s playback area, or in another appropriate manner.

User agents that cannot render the video may instead make the element
[represent](dom.html#represents){#the-video-element:represents-7} a link
to an external video playback utility or to the video data itself.

When a
[`video`{#the-video-element:the-video-element-19}](#the-video-element)
element\'s [media
resource](#media-resource){#the-video-element:media-resource-8} has a
video channel, the element provides a [paint
source](https://drafts.csswg.org/css-images-4/#paint-source){#the-video-element:paint-source
x-internal="paint-source"} whose width is the [media
resource](#media-resource){#the-video-element:media-resource-9}\'s
[natural
width](#concept-video-intrinsic-width){#the-video-element:concept-video-intrinsic-width},
whose height is the [media
resource](#media-resource){#the-video-element:media-resource-10}\'s
[natural
height](#concept-video-intrinsic-height){#the-video-element:concept-video-intrinsic-height},
and whose appearance is the frame of video corresponding to the [current
playback
position](#current-playback-position){#the-video-element:current-playback-position-9},
if that is available, or else (e.g. when the video is seeking or
buffering) its previous appearance, if any, or else (e.g. because the
video is still loading the first frame) blackness.

------------------------------------------------------------------------

`video`{.variable}`.`[`videoWidth`](#dom-video-videowidth){#dom-video-videowidth-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLVideoElement/videoWidth](https://developer.mozilla.org/en-US/docs/Web/API/HTMLVideoElement/videoWidth "The HTMLVideoElement interface's read-only videoWidth property indicates the intrinsic width of the video, expressed in CSS pixels. In simple terms, this is the width of the media in its natural size.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

`video`{.variable}`.`[`videoHeight`](#dom-video-videoheight){#dom-video-videoheight-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLVideoElement/videoHeight](https://developer.mozilla.org/en-US/docs/Web/API/HTMLVideoElement/videoHeight "The HTMLVideoElement interface's read-only videoHeight property indicates the intrinsic height of the video, expressed in CSS pixels. In simple terms, this is the height of the media in its natural size.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

These attributes return the natural dimensions of the video, or 0 if the
dimensions are not known.

The [natural width]{#concept-video-intrinsic-width .dfn} and [natural
height]{#concept-video-intrinsic-height .dfn} of the [media
resource](#media-resource){#the-video-element:media-resource-11} are the
dimensions of the resource in [CSS
pixels](https://drafts.csswg.org/css-values/#px){#the-video-element:'px'
x-internal="'px'"} after taking into account the resource\'s dimensions,
aspect ratio, clean aperture, resolution, and so forth, as defined for
the format used by the resource. If an anamorphic format does not define
how to apply the aspect ratio to the video data\'s dimensions to obtain
the \"correct\" dimensions, then the user agent must apply the ratio by
increasing one dimension and leaving the other unchanged.

The [`videoWidth`]{#dom-video-videowidth .dfn dfn-for="HTMLVideoElement"
dfn-type="attribute"} IDL attribute must return the [natural
width](#concept-video-intrinsic-width){#the-video-element:concept-video-intrinsic-width-2}
of the video in [CSS
pixels](https://drafts.csswg.org/css-values/#px){#the-video-element:'px'-2
x-internal="'px'"}. The [`videoHeight`]{#dom-video-videoheight .dfn
dfn-for="HTMLVideoElement" dfn-type="attribute"} IDL attribute must
return the [natural
height](#concept-video-intrinsic-height){#the-video-element:concept-video-intrinsic-height-2}
of the video in [CSS
pixels](https://drafts.csswg.org/css-values/#px){#the-video-element:'px'-3
x-internal="'px'"}. If the element\'s
[`readyState`{#the-video-element:dom-media-readystate-3}](#dom-media-readystate)
attribute is
[`HAVE_NOTHING`{#the-video-element:dom-media-have_nothing-2}](#dom-media-have_nothing),
then the attributes must return 0.

Whenever the [natural
width](#concept-video-intrinsic-width){#the-video-element:concept-video-intrinsic-width-3}
or [natural
height](#concept-video-intrinsic-height){#the-video-element:concept-video-intrinsic-height-3}
of the video changes (including, for example, because the [selected
video
track](#dom-videotrack-selected){#the-video-element:dom-videotrack-selected-2}
was changed), if the element\'s
[`readyState`{#the-video-element:dom-media-readystate-4}](#dom-media-readystate)
attribute is not
[`HAVE_NOTHING`{#the-video-element:dom-media-have_nothing-3}](#dom-media-have_nothing),
the user agent must [queue a media element
task](#queue-a-media-element-task){#the-video-element:queue-a-media-element-task}
given the [media
element](#media-element){#the-video-element:media-element-4} to [fire an
event](https://dom.spec.whatwg.org/#concept-event-fire){#the-video-element:concept-event-fire
x-internal="concept-event-fire"} named
[`resize`{#the-video-element:event-media-resize}](#event-media-resize)
at the [media
element](#media-element){#the-video-element:media-element-5}.

The
[`video`{#the-video-element:the-video-element-20}](#the-video-element)
element supports [dimension
attributes](embedded-content-other.html#dimension-attributes){#the-video-element:dimension-attributes}.

In the absence of style rules to the contrary, video content should be
rendered inside the element\'s playback area such that the video content
is shown centered in the playback area at the largest possible size that
fits completely within it, with the video content\'s aspect ratio being
preserved. Thus, if the aspect ratio of the playback area does not match
the aspect ratio of the video, the video will be shown letterboxed or
pillarboxed. Areas of the element\'s playback area that do not contain
the video represent nothing.

In user agents that implement CSS, the above requirement can be
implemented by using the [style rule suggested in the Rendering
section](rendering.html#video-object-fit).

The [natural
width](https://drafts.csswg.org/css-images/#natural-width){#the-video-element:natural-width
x-internal="natural-width"} of a
[`video`{#the-video-element:the-video-element-21}](#the-video-element)
element\'s playback area is the [natural
width](https://drafts.csswg.org/css-images/#natural-width){#the-video-element:natural-width-2
x-internal="natural-width"} of the [poster
frame](#poster-frame){#the-video-element:poster-frame-9}, if that is
available and the element currently
[represents](dom.html#represents){#the-video-element:represents-8} its
poster frame; otherwise, it is the [natural
width](#concept-video-intrinsic-width){#the-video-element:concept-video-intrinsic-width-4}
of the video resource, if that is available; otherwise the [natural
width](https://drafts.csswg.org/css-images/#natural-width){#the-video-element:natural-width-3
x-internal="natural-width"} is missing.

The [natural
height](https://drafts.csswg.org/css-images/#natural-height){#the-video-element:natural-height
x-internal="natural-height"} of a
[`video`{#the-video-element:the-video-element-22}](#the-video-element)
element\'s playback area is the [natural
height](https://drafts.csswg.org/css-images/#natural-height){#the-video-element:natural-height-2
x-internal="natural-height"} of the [poster
frame](#poster-frame){#the-video-element:poster-frame-10}, if that is
available and the element currently
[represents](dom.html#represents){#the-video-element:represents-9} its
poster frame; otherwise it is the [natural
height](#concept-video-intrinsic-height){#the-video-element:concept-video-intrinsic-height-4}
of the video resource, if that is available; otherwise the [natural
height](https://drafts.csswg.org/css-images/#natural-height){#the-video-element:natural-height-3
x-internal="natural-height"} is missing.

The [default object
size](https://drafts.csswg.org/css-images/#default-object-size){#the-video-element:default-object-size
x-internal="default-object-size"} is a width of 300 [CSS
pixels](https://drafts.csswg.org/css-values/#px){#the-video-element:'px'-4
x-internal="'px'"} and a height of 150 [CSS
pixels](https://drafts.csswg.org/css-values/#px){#the-video-element:'px'-5
x-internal="'px'"}. [\[CSSIMAGES\]](references.html#refsCSSIMAGES)

------------------------------------------------------------------------

User agents should provide controls to enable or disable the display of
closed captions, audio description tracks, and other additional data
associated with the video stream, though such features should, again,
not interfere with the page\'s normal rendering.

User agents may allow users to view the video content in manners more
suitable to the user, such as fullscreen or in an independent resizable
window. User agents may even trigger such a viewing mode by default upon
playing a video, although they should not do so when the
[`playsinline`{#the-video-element:attr-video-playsinline-4}](#attr-video-playsinline)
attribute is specified. As with the other user interface features,
controls to enable this should not interfere with the page\'s normal
rendering unless the user agent is [exposing a user
interface](#expose-a-user-interface-to-the-user){#the-video-element:expose-a-user-interface-to-the-user}.
In such an independent viewing mode, however, user agents may make full
user interfaces visible, even if the
[`controls`{#the-video-element:attr-media-controls-4}](#attr-media-controls)
attribute is absent.

User agents may allow video playback to affect system features that
could interfere with the user\'s experience; for example, user agents
could disable screensavers while video playback is in progress.

------------------------------------------------------------------------

The [`poster`]{#dom-video-poster .dfn dfn-for="HTMLVideoElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-video-element:reflect}
the
[`poster`{#the-video-element:attr-video-poster-6}](#attr-video-poster)
content attribute.

The [`playsInline`]{#dom-video-playsinline .dfn
dfn-for="HTMLVideoElement" dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-video-element:reflect-2}
the
[`playsinline`{#the-video-element:attr-video-playsinline-5}](#attr-video-playsinline)
content attribute.

::: example
This example shows how to detect when a video has failed to play
correctly:

``` html
<script>
 function failed(e) {
   // video playback failed - show a message saying why
   switch (e.target.error.code) {
     case e.target.error.MEDIA_ERR_ABORTED:
       alert('You aborted the video playback.');
       break;
     case e.target.error.MEDIA_ERR_NETWORK:
       alert('A network error caused the video download to fail part-way.');
       break;
     case e.target.error.MEDIA_ERR_DECODE:
       alert('The video playback was aborted due to a corruption problem or because the video used features your browser did not support.');
       break;
     case e.target.error.MEDIA_ERR_SRC_NOT_SUPPORTED:
       alert('The video could not be loaded, either because the server or network failed or because the format is not supported.');
       break;
     default:
       alert('An unknown error occurred.');
       break;
   }
 }
</script>
<p><video src="tgif.vid" autoplay controls onerror="failed(event)"></video></p>
<p><a href="tgif.vid">Download the video file</a>.</p>
```
:::

#### [4.8.9]{.secno} The [`audio`]{#audio .dfn dfn-type="element"} element[](#the-audio-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/audio](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/audio "The <audio> HTML element is used to embed sound content in documents. It may contain one or more audio sources, represented using the src attribute or the <source> element: the browser will choose the most suitable one. It can also be the destination for streamed media, using a MediaStream.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android3+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLAudioElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLAudioElement "The HTMLAudioElement interface provides access to the properties of <audio> elements, as well as methods to manipulate them.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

[Categories](dom.html#concept-element-categories){#the-audio-element:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-audio-element:flow-content-2}.
:   [Phrasing
    content](dom.html#phrasing-content-2){#the-audio-element:phrasing-content-2}.
:   [Embedded
    content](dom.html#embedded-content-category){#the-audio-element:embedded-content-category}.
:   If the element has a
    [`controls`{#the-audio-element:attr-media-controls}](#attr-media-controls)
    attribute: [Interactive
    content](dom.html#interactive-content-2){#the-audio-element:interactive-content-2}.
:   If the element has a
    [`controls`{#the-audio-element:attr-media-controls-2}](#attr-media-controls)
    attribute: [Palpable
    content](dom.html#palpable-content-2){#the-audio-element:palpable-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-audio-element:concept-element-contexts}:
:   Where [embedded
    content](dom.html#embedded-content-category){#the-audio-element:embedded-content-category-2}
    is expected.

[Content model](dom.html#concept-element-content-model){#the-audio-element:concept-element-content-model}:
:   If the element has a
    [`src`{#the-audio-element:attr-media-src}](#attr-media-src)
    attribute: zero or more
    [`track`{#the-audio-element:the-track-element}](#the-track-element)
    elements, then
    [transparent](dom.html#transparent){#the-audio-element:transparent},
    but with no [media
    element](#media-element){#the-audio-element:media-element}
    descendants.
:   If the element does not have a
    [`src`{#the-audio-element:attr-media-src-2}](#attr-media-src)
    attribute: zero or more
    [`source`{#the-audio-element:the-source-element}](embedded-content.html#the-source-element)
    elements, then zero or more
    [`track`{#the-audio-element:the-track-element-2}](#the-track-element)
    elements, then
    [transparent](dom.html#transparent){#the-audio-element:transparent-2},
    but with no [media
    element](#media-element){#the-audio-element:media-element-2}
    descendants.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-audio-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-audio-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-audio-element:global-attributes}
:   [`src`{#the-audio-element:attr-media-src-3}](#attr-media-src) ---
    Address of the resource
:   [`crossorigin`{#the-audio-element:attr-media-crossorigin}](#attr-media-crossorigin)
    --- How the element handles crossorigin requests
:   [`preload`{#the-audio-element:attr-media-preload}](#attr-media-preload)
    --- Hints how much buffering the [media
    resource](#media-resource){#the-audio-element:media-resource} will
    likely need
:   [`autoplay`{#the-audio-element:attr-media-autoplay}](#attr-media-autoplay)
    --- Hint that the [media
    resource](#media-resource){#the-audio-element:media-resource-2} can
    be started automatically when the page is loaded
:   [`loop`{#the-audio-element:attr-media-loop}](#attr-media-loop) ---
    Whether to loop the [media
    resource](#media-resource){#the-audio-element:media-resource-3}
:   [`muted`{#the-audio-element:attr-media-muted}](#attr-media-muted)
    --- Whether to mute the [media
    resource](#media-resource){#the-audio-element:media-resource-4} by
    default
:   [`controls`{#the-audio-element:attr-media-controls-3}](#attr-media-controls)
    --- Show user agent controls

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-audio-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-audio).
:   [For implementers](https://w3c.github.io/html-aam/#el-audio).

[DOM interface](dom.html#concept-element-dom){#the-audio-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window,
     LegacyFactoryFunction=Audio(optional DOMString src)]
    interface HTMLAudioElement : HTMLMediaElement {
      [HTMLConstructor] constructor();
    };
    ```

An [`audio`{#the-audio-element:the-audio-element}](#the-audio-element)
element [represents](dom.html#represents){#the-audio-element:represents}
a sound or audio stream.

Content may be provided inside the
[`audio`{#the-audio-element:the-audio-element-2}](#the-audio-element)
element. User agents should not show this content to the user; it is
intended for older web browsers which do not support
[`audio`{#the-audio-element:the-audio-element-3}](#the-audio-element),
so that text can be shown to the users of these older browsers informing
them of how to access the audio contents.

In particular, this content is not intended to address accessibility
concerns. To make audio content accessible to the deaf or to those with
other physical or cognitive disabilities, a variety of features are
available. If captions or a sign language video are available, the
[`video`{#the-audio-element:the-video-element}](#the-video-element)
element can be used instead of the
[`audio`{#the-audio-element:the-audio-element-4}](#the-audio-element)
element to play the audio, allowing users to enable the visual
alternatives. Chapter titles can be provided to aid navigation, using
the
[`track`{#the-audio-element:the-track-element-3}](#the-track-element)
element and a [WebVTT
file](https://w3c.github.io/webvtt/#webvtt-file){#the-audio-element:webvtt-file
x-internal="webvtt-file"}. And, naturally, transcripts or other textual
alternatives can be provided by simply linking to them in the prose near
the
[`audio`{#the-audio-element:the-audio-element-5}](#the-audio-element)
element. [\[WEBVTT\]](references.html#refsWEBVTT)

The
[`audio`{#the-audio-element:the-audio-element-6}](#the-audio-element)
element is a [media
element](#media-element){#the-audio-element:media-element-3} whose
[media data](#media-data){#the-audio-element:media-data} is ostensibly
audio data.

The [`src`{#the-audio-element:attr-media-src-4}](#attr-media-src),
[`crossorigin`{#the-audio-element:attr-media-crossorigin-2}](#attr-media-crossorigin),
[`preload`{#the-audio-element:attr-media-preload-2}](#attr-media-preload),
[`autoplay`{#the-audio-element:attr-media-autoplay-2}](#attr-media-autoplay),
[`loop`{#the-audio-element:attr-media-loop-2}](#attr-media-loop),
[`muted`{#the-audio-element:attr-media-muted-2}](#attr-media-muted), and
[`controls`{#the-audio-element:attr-media-controls-4}](#attr-media-controls)
attributes are [the attributes common to all media
elements](#media-element-attributes){#the-audio-element:media-element-attributes}.

`audio`{.variable}` = new `[`Audio`](#dom-audio){#dom-audio-dev}`([ ``url`{.variable}` ])`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLAudioElement/Audio](https://developer.mozilla.org/en-US/docs/Web/API/HTMLAudioElement/Audio "The Audio() constructor creates and returns a new HTMLAudioElement which can be either attached to a document for the user to interact with and/or listen to, or can be used offscreen to manage and play audio.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns a new
[`audio`{#the-audio-element:the-audio-element-7}](#the-audio-element)
element, with the
[`src`{#the-audio-element:attr-media-src-5}](#attr-media-src) attribute
set to the value passed in the argument, if applicable.

A legacy factory function is provided for creating
[`HTMLAudioElement`{#the-audio-element:htmlaudioelement}](#htmlaudioelement)
objects (in addition to the factory methods from DOM such as
[`createElement()`{#the-audio-element:dom-document-createelement}](https://dom.spec.whatwg.org/#dom-document-createelement){x-internal="dom-document-createelement"}):
[`Audio(``src`{.variable}`)`]{#dom-audio .dfn}. When invoked, the legacy
factory function must perform the following steps:

1.  Let `document`{.variable} be the [current global
    object](webappapis.html#current-global-object){#the-audio-element:current-global-object}\'s
    [associated
    `Document`](nav-history-apis.html#concept-document-window){#the-audio-element:concept-document-window}.

2.  Let `audio`{.variable} be the result of [creating an
    element](https://dom.spec.whatwg.org/#concept-create-element){#the-audio-element:create-an-element
    x-internal="create-an-element"} given `document`{.variable},
    \"`audio`\", and the [HTML
    namespace](https://infra.spec.whatwg.org/#html-namespace){#the-audio-element:html-namespace-2
    x-internal="html-namespace-2"}.

3.  [Set an attribute
    value](https://dom.spec.whatwg.org/#concept-element-attributes-set-value){#the-audio-element:concept-element-attributes-set-value
    x-internal="concept-element-attributes-set-value"} for
    `audio`{.variable} using
    \"[`preload`{#the-audio-element:attr-media-preload-3}](#attr-media-preload)\"
    and
    \"[`auto`{#the-audio-element:attr-media-preload-auto}](#attr-media-preload-auto)\".

4.  If `src`{.variable} is given, then [set an attribute
    value](https://dom.spec.whatwg.org/#concept-element-attributes-set-value){#the-audio-element:concept-element-attributes-set-value-2
    x-internal="concept-element-attributes-set-value"} for
    `audio`{.variable} using
    \"[`src`{#the-audio-element:attr-media-src-6}](#attr-media-src)\"
    and `src`{.variable}. (This will [cause the user agent to
    invoke](#concept-media-load-algorithm-at-creation) the object\'s
    [resource selection
    algorithm](#concept-media-load-algorithm){#the-audio-element:concept-media-load-algorithm}
    before returning.)

5.  Return `audio`{.variable}.

#### [4.8.10]{.secno} The [`track`]{.dfn dfn-type="element"} element[](#the-track-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/track](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/track "The <track> HTML element is used as a child of the media elements, <audio> and <video>. It lets you specify timed text tracks (or time-based data), for example to automatically handle subtitles. The tracks are formatted in WebVTT format (.vtt files) — Web Video Text Tracks.")

Support in all current engines.

::: support
[Firefox31+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome23+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android25+]{.chrome_android .yes}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLTrackElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLTrackElement "The HTMLTrackElement interface represents an HTML <track> element within the DOM. This element can be used as a child of either <audio> or <video> to specify a text track containing information such as closed captions or subtitles.")

Support in all current engines.

::: support
[Firefox31+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome23+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12+]{.opera_android .yes}
:::
::::
:::::

[Categories](dom.html#concept-element-categories){#the-track-element:concept-element-categories}:
:   None.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-track-element:concept-element-contexts}:
:   As a child of a [media
    element](#media-element){#the-track-element:media-element}, before
    any [flow
    content](dom.html#flow-content-2){#the-track-element:flow-content-2}.

[Content model](dom.html#concept-element-content-model){#the-track-element:concept-element-content-model}:
:   [Nothing](dom.html#concept-content-nothing){#the-track-element:concept-content-nothing}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-track-element:concept-element-tag-omission}:
:   No [end
    tag](syntax.html#syntax-end-tag){#the-track-element:syntax-end-tag}.

[Content attributes](dom.html#concept-element-attributes){#the-track-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-track-element:global-attributes}
:   [`kind`{#the-track-element:attr-track-kind}](#attr-track-kind) ---
    The type of text track
:   [`src`{#the-track-element:attr-track-src}](#attr-track-src) ---
    Address of the resource
:   [`srclang`{#the-track-element:attr-track-srclang}](#attr-track-srclang)
    --- Language of the text track
:   [`label`{#the-track-element:attr-track-label}](#attr-track-label)
    --- User-visible label
:   [`default`{#the-track-element:attr-track-default}](#attr-track-default)
    --- Enable the track if no other [text
    track](#text-track){#the-track-element:text-track} is more suitable

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-track-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-track).
:   [For implementers](https://w3c.github.io/html-aam/#el-track).

[DOM interface](dom.html#concept-element-dom){#the-track-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLTrackElement : HTMLElement {
      [HTMLConstructor] constructor();

      [CEReactions] attribute DOMString kind;
      [CEReactions] attribute USVString src;
      [CEReactions] attribute DOMString srclang;
      [CEReactions] attribute DOMString label;
      [CEReactions] attribute boolean default;

      const unsigned short NONE = 0;
      const unsigned short LOADING = 1;
      const unsigned short LOADED = 2;
      const unsigned short ERROR = 3;
      readonly attribute unsigned short readyState;

      readonly attribute TextTrack track;
    };
    ```

The [`track`{#the-track-element:the-track-element}](#the-track-element)
element allows authors to specify explicit external timed [text
tracks](#text-track){#the-track-element:text-track-2} for [media
elements](#media-element){#the-track-element:media-element-2}. It does
not [represent](dom.html#represents){#the-track-element:represents}
anything on its own.

The [`kind`]{#attr-track-kind .dfn dfn-for="track"
dfn-type="element-attr"} attribute is an [enumerated
attribute](common-microsyntaxes.html#enumerated-attribute){#the-track-element:enumerated-attribute}
with the following keywords and states:

Keyword

State

Brief description

[`subtitles`]{#attr-track-kind-keyword-subtitles .dfn
dfn-for="track/kind" dfn-type="attr-value"}

[Subtitles]{#attr-track-kind-subtitles .dfn}

Transcription or translation of the dialogue, suitable for when the
sound is available but not understood (e.g. because the user does not
understand the language of the [media
resource](#media-resource){#the-track-element:media-resource}\'s audio
track). Overlaid on the video.

[`captions`]{#attr-track-kind-keyword-captions .dfn dfn-for="track/kind"
dfn-type="attr-value"}

[Captions]{#attr-track-kind-captions .dfn}

Transcription or translation of the dialogue, sound effects, relevant
musical cues, and other relevant audio information, suitable for when
sound is unavailable or not clearly audible (e.g. because it is muted,
drowned-out by ambient noise, or because the user is deaf). Overlaid on
the video; labeled as appropriate for the hard-of-hearing.

[`descriptions`]{#attr-track-kind-keyword-descriptions .dfn
dfn-for="track/kind" dfn-type="attr-value"}

[Descriptions]{#attr-track-kind-descriptions .dfn}

Textual descriptions of the video component of the [media
resource](#media-resource){#the-track-element:media-resource-2},
intended for audio synthesis when the visual component is obscured,
unavailable, or not usable (e.g. because the user is interacting with
the application without a screen while driving, or because the user is
blind). Synthesized as audio.

[`chapters`]{#attr-track-kind-keyword-chapters .dfn dfn-for="track/kind"
dfn-type="attr-value"}

[Chapters metadata]{#attr-track-kind-chapters .dfn}

Tracks intended for use from script. Not displayed by the user agent.

[`metadata`]{#attr-track-kind-keyword-metadata .dfn dfn-for="track/kind"
dfn-type="attr-value"}

[Metadata]{#attr-track-kind-metadata .dfn}

The attribute\'s *[missing value
default](common-microsyntaxes.html#missing-value-default)* is the
[subtitles](#attr-track-kind-subtitles){#the-track-element:attr-track-kind-subtitles}
state, and its *[invalid value
default](common-microsyntaxes.html#invalid-value-default)* is the
[metadata](#attr-track-kind-metadata){#the-track-element:attr-track-kind-metadata}
state.

The [`src`]{#attr-track-src .dfn dfn-for="track"
dfn-type="element-attr"} attribute gives the
[URL](https://url.spec.whatwg.org/#concept-url){#the-track-element:url
x-internal="url"} of the text track data. The value must be a [valid
non-empty URL potentially surrounded by
spaces](urls-and-fetching.html#valid-non-empty-url-potentially-surrounded-by-spaces){#the-track-element:valid-non-empty-url-potentially-surrounded-by-spaces}.
This attribute must be present.

The element has an associated [track URL]{#track-url .dfn} (a string),
initially the empty string.

When the element\'s
[`src`{#the-track-element:attr-track-src-2}](#attr-track-src) attribute
is set, run these steps:

1.  Let `trackURL`{.variable} be failure.

2.  Let `value`{.variable} be the element\'s
    [`src`{#the-track-element:attr-track-src-3}](#attr-track-src)
    attribute value.

3.  If `value`{.variable} is not the empty string, then set
    `trackURL`{.variable} to the result of
    [encoding-parsing-and-serializing a
    URL](urls-and-fetching.html#encoding-parsing-and-serializing-a-url){#the-track-element:encoding-parsing-and-serializing-a-url}
    given `value`{.variable}, relative to the element\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-track-element:node-document
    x-internal="node-document"}.

4.  Set the element\'s [track
    URL](#track-url){#the-track-element:track-url} to
    `trackURL`{.variable} if it is not failure; otherwise to the empty
    string.

If the element\'s [track
URL](#track-url){#the-track-element:track-url-2} identifies a WebVTT
resource, and the element\'s
[`kind`{#the-track-element:attr-track-kind-2}](#attr-track-kind)
attribute is not in the [chapters
metadata](#attr-track-kind-chapters){#the-track-element:attr-track-kind-chapters}
or
[metadata](#attr-track-kind-metadata){#the-track-element:attr-track-kind-metadata-2}
state, then the WebVTT file must be a [WebVTT file using cue
text](https://w3c.github.io/webvtt/#webvtt-file-using-cue-text){#the-track-element:webvtt-file-using-cue-text
x-internal="webvtt-file-using-cue-text"}.
[\[WEBVTT\]](references.html#refsWEBVTT)

The [`srclang`]{#attr-track-srclang .dfn dfn-for="track"
dfn-type="element-attr"} attribute gives the language of the text track
data. The value must be a valid BCP 47 language tag. This attribute must
be present if the element\'s
[`kind`{#the-track-element:attr-track-kind-3}](#attr-track-kind)
attribute is in the
[subtitles](#attr-track-kind-subtitles){#the-track-element:attr-track-kind-subtitles-2}
state. [\[BCP47\]](references.html#refsBCP47)

If the element has a
[`srclang`{#the-track-element:attr-track-srclang-2}](#attr-track-srclang)
attribute whose value is not the empty string, then the element\'s
[track language]{#track-language .dfn} is the value of the attribute.
Otherwise, the element has no [track
language](#track-language){#the-track-element:track-language}.

The [`label`]{#attr-track-label .dfn dfn-for="track"
dfn-type="element-attr"} attribute gives a user-readable title for the
track. This title is used by user agents when listing
[subtitle](#attr-track-kind-subtitles){#the-track-element:attr-track-kind-subtitles-3},
[caption](#attr-track-kind-captions){#the-track-element:attr-track-kind-captions},
and [audio
description](#attr-track-kind-descriptions){#the-track-element:attr-track-kind-descriptions}
tracks in their user interface.

The value of the
[`label`{#the-track-element:attr-track-label-2}](#attr-track-label)
attribute, if the attribute is present, must not be the empty string.
Furthermore, there must not be two
[`track`{#the-track-element:the-track-element-2}](#the-track-element)
element children of the same [media
element](#media-element){#the-track-element:media-element-3} whose
[`kind`{#the-track-element:attr-track-kind-4}](#attr-track-kind)
attributes are in the same state, whose
[`srclang`{#the-track-element:attr-track-srclang-3}](#attr-track-srclang)
attributes are both missing or have values that represent the same
language, and whose
[`label`{#the-track-element:attr-track-label-3}](#attr-track-label)
attributes are again both missing or both have the same value.

If the element has a
[`label`{#the-track-element:attr-track-label-4}](#attr-track-label)
attribute whose value is not the empty string, then the element\'s
[track label]{#track-label .dfn} is the value of the attribute.
Otherwise, the element\'s [track
label](#track-label){#the-track-element:track-label} is an empty string.

The [`default`]{#attr-track-default .dfn dfn-for="track"
dfn-type="element-attr"} attribute is a [boolean
attribute](common-microsyntaxes.html#boolean-attribute){#the-track-element:boolean-attribute},
which, if specified, indicates that the track is to be enabled if the
user\'s preferences do not indicate that another track would be more
appropriate.

Each [media element](#media-element){#the-track-element:media-element-4}
must have no more than one
[`track`{#the-track-element:the-track-element-3}](#the-track-element)
element child whose
[`kind`{#the-track-element:attr-track-kind-5}](#attr-track-kind)
attribute is in the
[subtitles](#attr-track-kind-subtitles){#the-track-element:attr-track-kind-subtitles-4}
or
[captions](#attr-track-kind-captions){#the-track-element:attr-track-kind-captions-2}
state and whose
[`default`{#the-track-element:attr-track-default-2}](#attr-track-default)
attribute is specified.

Each [media element](#media-element){#the-track-element:media-element-5}
must have no more than one
[`track`{#the-track-element:the-track-element-4}](#the-track-element)
element child whose
[`kind`{#the-track-element:attr-track-kind-6}](#attr-track-kind)
attribute is in the
[description](#attr-track-kind-descriptions){#the-track-element:attr-track-kind-descriptions-2}
state and whose
[`default`{#the-track-element:attr-track-default-3}](#attr-track-default)
attribute is specified.

Each [media element](#media-element){#the-track-element:media-element-6}
must have no more than one
[`track`{#the-track-element:the-track-element-5}](#the-track-element)
element child whose
[`kind`{#the-track-element:attr-track-kind-7}](#attr-track-kind)
attribute is in the [chapters
metadata](#attr-track-kind-chapters){#the-track-element:attr-track-kind-chapters-2}
state and whose
[`default`{#the-track-element:attr-track-default-4}](#attr-track-default)
attribute is specified.

There is no limit on the number of
[`track`{#the-track-element:the-track-element-6}](#the-track-element)
elements whose
[`kind`{#the-track-element:attr-track-kind-8}](#attr-track-kind)
attribute is in the
[metadata](#attr-track-kind-metadata){#the-track-element:attr-track-kind-metadata-3}
state and whose
[`default`{#the-track-element:attr-track-default-5}](#attr-track-default)
attribute is specified.

`track`{.variable}`.`[`readyState`](#dom-track-readystate){#dom-track-readystate-dev}

:   Returns the [text track readiness
    state](#text-track-readiness-state){#the-track-element:text-track-readiness-state},
    represented by a number from the following list:

    `track`{.variable}`.`[`NONE`](#dom-track-none){#dom-track-none-dev}` (0)`
    :   The [text track not
        loaded](#text-track-not-loaded){#the-track-element:text-track-not-loaded}
        state.

    `track`{.variable}`.`[`LOADING`](#dom-track-loading){#dom-track-loading-dev}` (1)`
    :   The [text track
        loading](#text-track-loading){#the-track-element:text-track-loading}
        state.

    `track`{.variable}`.`[`LOADED`](#dom-track-loaded){#dom-track-loaded-dev}` (2)`
    :   The [text track
        loaded](#text-track-loaded){#the-track-element:text-track-loaded}
        state.

    `track`{.variable}`.`[`ERROR`](#dom-track-error){#dom-track-error-dev}` (3)`

    :   The [text track failed to
        load](#text-track-failed-to-load){#the-track-element:text-track-failed-to-load}
        state.

`track`{.variable}`.`[`track`](#dom-track-track){#dom-track-track-dev}

:   Returns the
    [`TextTrack`{#the-track-element:texttrack-2}](#texttrack) object
    corresponding to the [text
    track](#text-track){#the-track-element:text-track-3} of the
    [`track`{#the-track-element:the-track-element-7}](#the-track-element)
    element.

The [`readyState`]{#dom-track-readystate .dfn dfn-for="HTMLTrackElement"
dfn-type="attribute"} attribute must return the numeric value
corresponding to the [text track readiness
state](#text-track-readiness-state){#the-track-element:text-track-readiness-state-2}
of the
[`track`{#the-track-element:the-track-element-8}](#the-track-element)
element\'s [text track](#text-track){#the-track-element:text-track-4},
as defined by the following list:

[`NONE`]{#dom-track-none .dfn dfn-for="HTMLTrackElement" dfn-type="const"} (numeric value 0)
:   The [text track not
    loaded](#text-track-not-loaded){#the-track-element:text-track-not-loaded-2}
    state.

[`LOADING`]{#dom-track-loading .dfn dfn-for="HTMLTrackElement" dfn-type="const"} (numeric value 1)
:   The [text track
    loading](#text-track-loading){#the-track-element:text-track-loading-2}
    state.

[`LOADED`]{#dom-track-loaded .dfn dfn-for="HTMLTrackElement" dfn-type="const"} (numeric value 2)
:   The [text track
    loaded](#text-track-loaded){#the-track-element:text-track-loaded-2}
    state.

[`ERROR`]{#dom-track-error .dfn dfn-for="HTMLTrackElement" dfn-type="const"} (numeric value 3)
:   The [text track failed to
    load](#text-track-failed-to-load){#the-track-element:text-track-failed-to-load-2}
    state.

The [`track`]{#dom-track-track .dfn dfn-for="HTMLTrackElement"
dfn-type="attribute"} IDL attribute must, on getting, return the
[`track`{#the-track-element:the-track-element-9}](#the-track-element)
element\'s [text track](#text-track){#the-track-element:text-track-5}\'s
corresponding [`TextTrack`{#the-track-element:texttrack-3}](#texttrack)
object.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLTrackElement/src](https://developer.mozilla.org/en-US/docs/Web/API/HTMLTrackElement/src "The HTMLTrackElement.src property reflects the value of the <track> element's src attribute, which indicates the URL of the text track's data.")

Support in all current engines.

::: support
[Firefox31+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome23+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12+]{.opera_android .yes}
:::
::::
:::::

The [`src`]{#dom-track-src .dfn dfn-for="HTMLTrackElement"
dfn-type="attribute"}, [`srclang`]{#dom-track-srclang .dfn
dfn-for="HTMLTrackElement" dfn-type="attribute"},
[`label`]{#dom-track-label .dfn dfn-for="HTMLTrackElement"
dfn-type="attribute"}, and [`default`]{#dom-track-default .dfn
dfn-for="HTMLTrackElement" dfn-type="attribute"} IDL attributes must
[reflect](common-dom-interfaces.html#reflect){#the-track-element:reflect}
the respective content attributes of the same name. The
[`kind`]{#dom-track-kind .dfn dfn-for="HTMLTrackElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-track-element:reflect-2}
the content attribute of the same name, [limited to only known
values](common-dom-interfaces.html#limited-to-only-known-values){#the-track-element:limited-to-only-known-values}.

::: example
This video has subtitles in several languages:

``` html
<video src="brave.webm">
 <track kind=subtitles src=brave.en.vtt srclang=en label="English">
 <track kind=captions src=brave.en.hoh.vtt srclang=en label="English for the Hard of Hearing">
 <track kind=subtitles src=brave.fr.vtt srclang=fr lang=fr label="Français">
 <track kind=subtitles src=brave.de.vtt srclang=de lang=de label="Deutsch">
</video>
```

(The [`lang`{#the-track-element:attr-lang}](dom.html#attr-lang)
attributes on the last two describe the language of the
[`label`{#the-track-element:attr-track-label-5}](#attr-track-label)
attribute, not the language of the subtitles themselves. The language of
the subtitles is given by the
[`srclang`{#the-track-element:attr-track-srclang-4}](#attr-track-srclang)
attribute.)
:::

#### [4.8.11]{.secno} Media elements[](#media-elements){.self-link}

[HTMLMediaElement](#htmlmediaelement){#media-elements:htmlmediaelement}
objects
([`audio`{#media-elements:the-audio-element}](#the-audio-element) and
[`video`{#media-elements:the-video-element}](#the-video-element), in
this specification) are simply known as [media elements]{#media-element
.dfn lt="media element" export=""}.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement "The HTMLMediaElement interface adds to HTMLElement the properties and methods needed to support basic media-related capabilities that are common to audio and video.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

``` idl
enum CanPlayTypeResult { "" /* empty string */, "maybe", "probably" };
typedef (MediaStream or MediaSource or Blob) MediaProvider;

[Exposed=Window]
interface HTMLMediaElement : HTMLElement {

  // error state
  readonly attribute MediaError? error;

  // network state
  [CEReactions] attribute USVString src;
  attribute MediaProvider? srcObject;
  readonly attribute USVString currentSrc;
  [CEReactions] attribute DOMString? crossOrigin;
  const unsigned short NETWORK_EMPTY = 0;
  const unsigned short NETWORK_IDLE = 1;
  const unsigned short NETWORK_LOADING = 2;
  const unsigned short NETWORK_NO_SOURCE = 3;
  readonly attribute unsigned short networkState;
  [CEReactions] attribute DOMString preload;
  readonly attribute TimeRanges buffered;
  undefined load();
  CanPlayTypeResult canPlayType(DOMString type);

  // ready state
  const unsigned short HAVE_NOTHING = 0;
  const unsigned short HAVE_METADATA = 1;
  const unsigned short HAVE_CURRENT_DATA = 2;
  const unsigned short HAVE_FUTURE_DATA = 3;
  const unsigned short HAVE_ENOUGH_DATA = 4;
  readonly attribute unsigned short readyState;
  readonly attribute boolean seeking;

  // playback state
  attribute double currentTime;
  undefined fastSeek(double time);
  readonly attribute unrestricted double duration;
  object getStartDate();
  readonly attribute boolean paused;
  attribute double defaultPlaybackRate;
  attribute double playbackRate;
  attribute boolean preservesPitch;
  readonly attribute TimeRanges played;
  readonly attribute TimeRanges seekable;
  readonly attribute boolean ended;
  [CEReactions] attribute boolean autoplay;
  [CEReactions] attribute boolean loop;
  Promise<undefined> play();
  undefined pause();

  // controls
  [CEReactions] attribute boolean controls;
  attribute double volume;
  attribute boolean muted;
  [CEReactions] attribute boolean defaultMuted;

  // tracks
  [SameObject] readonly attribute AudioTrackList audioTracks;
  [SameObject] readonly attribute VideoTrackList videoTracks;
  [SameObject] readonly attribute TextTrackList textTracks;
  TextTrack addTextTrack(TextTrackKind kind, optional DOMString label = "", optional DOMString language = "");
};
```

The [media element attributes]{#media-element-attributes .dfn},
[`src`{#media-elements:attr-media-src}](#attr-media-src),
[`crossorigin`{#media-elements:attr-media-crossorigin}](#attr-media-crossorigin),
[`preload`{#media-elements:attr-media-preload}](#attr-media-preload),
[`autoplay`{#media-elements:attr-media-autoplay}](#attr-media-autoplay),
[`loop`{#media-elements:attr-media-loop}](#attr-media-loop),
[`muted`{#media-elements:attr-media-muted}](#attr-media-muted), and
[`controls`{#media-elements:attr-media-controls}](#attr-media-controls),
apply to all [media
elements](#media-element){#media-elements:media-element}. They are
defined in this section.

[Media elements](#media-element){#media-elements:media-element-2} are
used to present audio data, or video and audio data, to the user. This
is referred to as [media data]{#media-data .dfn
dfn-for="HTMLMediaElement" export=""} in this section, since this
section applies equally to [media
elements](#media-element){#media-elements:media-element-3} for audio or
for video. The term [media resource]{#media-resource .dfn} is used to
refer to the complete set of media data, e.g. the complete video file,
or complete audio file.

A [media resource](#media-resource){#media-elements:media-resource} has
an associated [origin]{#media-resource-origin .dfn}, which is either
\"`none`\", \"`multiple`\", \"`rewritten`\", or an
[origin](browsers.html#concept-origin){#media-elements:concept-origin}.
It is initially set to \"`none`\".

A [media resource](#media-resource){#media-elements:media-resource-2}
can have multiple audio and video tracks. For the purposes of a [media
element](#media-element){#media-elements:media-element-4}, the video
data of the [media
resource](#media-resource){#media-elements:media-resource-3} is only
that of the currently selected track (if any) as given by the element\'s
[`videoTracks`{#media-elements:dom-media-videotracks-2}](#dom-media-videotracks)
attribute when the [event
loop](webappapis.html#event-loop){#media-elements:event-loop} last
reached [step 1](webappapis.html#step1), and the audio data of the
[media resource](#media-resource){#media-elements:media-resource-4} is
the result of mixing all the currently enabled tracks (if any) given by
the element\'s
[`audioTracks`{#media-elements:dom-media-audiotracks-2}](#dom-media-audiotracks)
attribute when the [event
loop](webappapis.html#event-loop){#media-elements:event-loop-2} last
reached [step 1](webappapis.html#step1).

Both [`audio`{#media-elements:the-audio-element-2}](#the-audio-element)
and [`video`{#media-elements:the-video-element-2}](#the-video-element)
elements can be used for both audio and video. The main difference
between the two is simply that the
[`audio`{#media-elements:the-audio-element-3}](#the-audio-element)
element has no playback area for visual content (such as video or
captions), whereas the
[`video`{#media-elements:the-video-element-3}](#the-video-element)
element does.

Each [media element](#media-element){#media-elements:media-element-5}
has a unique [media element event task
source]{#media-element-event-task-source .dfn}.

To [queue a media element task]{#queue-a-media-element-task .dfn
export=""} with a [media
element](#media-element){#media-elements:media-element-6}
`element`{.variable} and a series of steps `steps`{.variable}, [queue an
element
task](webappapis.html#queue-an-element-task){#media-elements:queue-an-element-task}
on the [media
element](#media-element){#media-elements:media-element-7}\'s [media
element event task
source](#media-element-event-task-source){#media-elements:media-element-event-task-source}
given `element`{.variable} and `steps`{.variable}.

##### [4.8.11.1]{.secno} Error codes[](#error-codes){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[MediaError](https://developer.mozilla.org/en-US/docs/Web/API/MediaError "The MediaError interface represents an error which occurred while handling media in an HTML media element based on HTMLMediaElement, such as <audio> or <video>.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
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

`media`{.variable}`.`[`error`](#dom-media-error){#dom-media-error-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/error](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/error "The HTMLMediaElement.error property is the MediaError object for the most recent error, or null if there has not been an error. When an error event is received by the element, you can determine details about what happened by examining this object.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns a [`MediaError`{#error-codes:mediaerror}](#mediaerror) object
representing the current error state of the element.

Returns null if there is no error.

All [media elements](#media-element){#error-codes:media-element} have an
associated error status, which records the last error the element
encountered since its [resource selection
algorithm](#concept-media-load-algorithm){#error-codes:concept-media-load-algorithm}
was last invoked. The [`error`]{#dom-media-error .dfn
dfn-for="HTMLMediaElement" dfn-type="attribute"} attribute, on getting,
must return the [`MediaError`{#error-codes:mediaerror-2}](#mediaerror)
object created for this last error, or null if there has not been an
error.

``` idl
[Exposed=Window]
interface MediaError {
  const unsigned short MEDIA_ERR_ABORTED = 1;
  const unsigned short MEDIA_ERR_NETWORK = 2;
  const unsigned short MEDIA_ERR_DECODE = 3;
  const unsigned short MEDIA_ERR_SRC_NOT_SUPPORTED = 4;

  readonly attribute unsigned short code;
  readonly attribute DOMString message;
};
```

`media`{.variable}`.`[`error`](#dom-media-error){#error-codes:dom-media-error}`.`[`code`](#dom-mediaerror-code){#dom-mediaerror-code-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[MediaError/code](https://developer.mozilla.org/en-US/docs/Web/API/MediaError/code "The read-only property MediaError.code returns a numeric value which represents the kind of error that occurred on a media element. To get a text string with specific diagnostic information, see MediaError.message.")

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

Returns the current error\'s error code, from the list below.

`media`{.variable}`.`[`error`](#dom-media-error){#error-codes:dom-media-error-2}`.`[`message`](#dom-mediaerror-message){#dom-mediaerror-message-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[MediaError/message](https://developer.mozilla.org/en-US/docs/Web/API/MediaError/message "The read-only property MediaError.message returns a human-readable string offering specific diagnostic details related to the error described by the MediaError object, or an empty string ("") if no diagnostic information can be determined or provided.")

Support in all current engines.

::: support
[Firefox52+]{.firefox .yes}[Safari15+]{.safari .yes}[Chrome59+]{.chrome
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

Returns a specific informative diagnostic message about the error
condition encountered. The message and message format are not generally
uniform across different user agents. If no such message is available,
then the empty string is returned.

Every [`MediaError`{#error-codes:mediaerror-3}](#mediaerror) object has
a [message]{#concept-mediaerror-message .dfn}, which is a string, and a
[code]{#concept-mediaerror-code .dfn}, which is one of the following:

[`MEDIA_ERR_ABORTED`]{#dom-mediaerror-media_err_aborted .dfn dfn-for="MediaError" dfn-type="const"} (numeric value 1)
:   The fetching process for the [media
    resource](#media-resource){#error-codes:media-resource} was aborted
    by the user agent at the user\'s request.

[`MEDIA_ERR_NETWORK`]{#dom-mediaerror-media_err_network .dfn dfn-for="MediaError" dfn-type="const"} (numeric value 2)
:   A network error of some description caused the user agent to stop
    fetching the [media
    resource](#media-resource){#error-codes:media-resource-2}, after the
    resource was established to be usable.

[`MEDIA_ERR_DECODE`]{#dom-mediaerror-media_err_decode .dfn dfn-for="MediaError" dfn-type="const"} (numeric value 3)
:   An error of some description occurred while decoding the [media
    resource](#media-resource){#error-codes:media-resource-3}, after the
    resource was established to be usable.

[`MEDIA_ERR_SRC_NOT_SUPPORTED`]{#dom-mediaerror-media_err_src_not_supported .dfn dfn-for="MediaError" dfn-type="const"} (numeric value 4)
:   The [media resource](#media-resource){#error-codes:media-resource-4}
    indicated by the
    [`src`{#error-codes:attr-media-src}](#attr-media-src) attribute or
    [assigned media provider
    object](#assigned-media-provider-object){#error-codes:assigned-media-provider-object}
    was not suitable.

To [create a `MediaError`]{#creating-a-mediaerror .dfn export=""}, given
an error code which is one of the above values, return a new
[`MediaError`{#error-codes:mediaerror-4}](#mediaerror) object whose
[code](#concept-mediaerror-code){#error-codes:concept-mediaerror-code}
is the given error code and whose
[message](#concept-mediaerror-message){#error-codes:concept-mediaerror-message}
is a string containing any details the user agent is able to supply
about the cause of the error condition, or the empty string if the user
agent is unable to supply such details. This message string must not
contain only the information already available via the supplied error
code; for example, it must not simply be a translation of the code into
a string format. If no additional information is available beyond that
provided by the error code, the
[message](#concept-mediaerror-message){#error-codes:concept-mediaerror-message-2}
must be set to the empty string.

The [`code`]{#dom-mediaerror-code .dfn dfn-for="MediaError"
dfn-type="attribute"} getter steps are to return
[this](https://webidl.spec.whatwg.org/#this){#error-codes:this
x-internal="this"}\'s
[code](#concept-mediaerror-code){#error-codes:concept-mediaerror-code-2}.

The [`message`]{#dom-mediaerror-message .dfn dfn-for="MediaError"
dfn-type="attribute"} getter steps are to return
[this](https://webidl.spec.whatwg.org/#this){#error-codes:this-2
x-internal="this"}\'s
[message](#concept-mediaerror-message){#error-codes:concept-mediaerror-message-3}.

##### [4.8.11.2]{.secno} Location of the media resource[](#location-of-the-media-resource){.self-link}

The [`src`]{#attr-media-src .dfn dfn-for="audio,video"
dfn-type="element-attr"} content attribute on [media
elements](#media-element){#location-of-the-media-resource:media-element}
gives the
[URL](https://url.spec.whatwg.org/#concept-url){#location-of-the-media-resource:url
x-internal="url"} of the media resource (video, audio) to show. The
attribute, if present, must contain a [valid non-empty URL potentially
surrounded by
spaces](urls-and-fetching.html#valid-non-empty-url-potentially-surrounded-by-spaces){#location-of-the-media-resource:valid-non-empty-url-potentially-surrounded-by-spaces}.

If the
[`itemprop`{#location-of-the-media-resource:names:-the-itemprop-attribute}](microdata.html#names:-the-itemprop-attribute)
attribute is specified on the [media
element](#media-element){#location-of-the-media-resource:media-element-2},
then the
[`src`{#location-of-the-media-resource:attr-media-src}](#attr-media-src)
attribute must also be specified.

The [`crossorigin`]{#attr-media-crossorigin .dfn dfn-for="audio,video"
dfn-type="element-attr"} content attribute on [media
elements](#media-element){#location-of-the-media-resource:media-element-3}
is a [CORS settings
attribute](urls-and-fetching.html#cors-settings-attribute){#location-of-the-media-resource:cors-settings-attribute}.

If a [media
element](#media-element){#location-of-the-media-resource:media-element-4}
is created with a
[`src`{#location-of-the-media-resource:attr-media-src-2}](#attr-media-src)
attribute, the user agent must
[immediately](infrastructure.html#immediately){#location-of-the-media-resource:immediately}
invoke the [media
element](#media-element){#location-of-the-media-resource:media-element-5}\'s
[resource selection
algorithm](#concept-media-load-algorithm){#location-of-the-media-resource:concept-media-load-algorithm}.

If a
[`src`{#location-of-the-media-resource:attr-media-src-3}](#attr-media-src)
attribute of a [media
element](#media-element){#location-of-the-media-resource:media-element-6}
is set or changed, the user agent must invoke the [media
element](#media-element){#location-of-the-media-resource:media-element-7}\'s
[media element load
algorithm](#media-element-load-algorithm){#location-of-the-media-resource:media-element-load-algorithm}.
(*Removing* the
[`src`{#location-of-the-media-resource:attr-media-src-4}](#attr-media-src)
attribute does not do this, even if there are
[`source`{#location-of-the-media-resource:the-source-element}](embedded-content.html#the-source-element)
elements present.)

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/src](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/src "The HTMLMediaElement.src property reflects the value of the HTML media element's src attribute, which indicates the URL of a media resource to use in the element.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The [`src`]{#dom-media-src .dfn dfn-for="HTMLMediaElement"
dfn-type="attribute"} IDL attribute on [media
elements](#media-element){#location-of-the-media-resource:media-element-8}
must
[reflect](common-dom-interfaces.html#reflect){#location-of-the-media-resource:reflect}
the content attribute of the same name.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/crossOrigin](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/crossOrigin "The HTMLMediaElement.crossOrigin property is the CORS setting for this media element. See CORS settings attributes for details.")

Support in all current engines.

::: support
[Firefox22+]{.firefox .yes}[Safari10+]{.safari .yes}[Chrome33+]{.chrome
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

The [`crossOrigin`]{#dom-media-crossorigin .dfn
dfn-for="HTMLMediaElement" dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#location-of-the-media-resource:reflect-2}
the
[`crossorigin`{#location-of-the-media-resource:attr-media-crossorigin}](#attr-media-crossorigin)
content attribute, [limited to only known
values](common-dom-interfaces.html#limited-to-only-known-values){#location-of-the-media-resource:limited-to-only-known-values}.

A [media provider object]{#media-provider-object .dfn} is an object that
can represent a [media
resource](#media-resource){#location-of-the-media-resource:media-resource},
separate from a
[URL](https://url.spec.whatwg.org/#concept-url){#location-of-the-media-resource:url-2
x-internal="url"}.
[`MediaStream`{#location-of-the-media-resource:mediastream}](https://w3c.github.io/mediacapture-main/getusermedia.html#idl-def-mediastream){x-internal="mediastream"}
objects,
[`MediaSource`{#location-of-the-media-resource:mediasource}](https://w3c.github.io/media-source/#idl-def-mediasource){x-internal="mediasource"}
objects, and
[`Blob`{#location-of-the-media-resource:blob}](https://w3c.github.io/FileAPI/#dfn-Blob){x-internal="blob"}
objects are all [media provider
objects](#media-provider-object){#location-of-the-media-resource:media-provider-object}.

Each [media
element](#media-element){#location-of-the-media-resource:media-element-9}
can have an [assigned media provider
object]{#assigned-media-provider-object .dfn}, which is a [media
provider
object](#media-provider-object){#location-of-the-media-resource:media-provider-object-2}.
When a [media
element](#media-element){#location-of-the-media-resource:media-element-10}
is created, it has no [assigned media provider
object](#assigned-media-provider-object){#location-of-the-media-resource:assigned-media-provider-object}.

`media`{.variable}`.`[`srcObject`](#dom-media-srcobject){#dom-media-srcobject-dev}` [ = ``source`{.variable}` ]`

::::: {.mdn-anno .wrapped .before}
**⚠**MDN

:::: feature
[HTMLMediaElement/srcObject](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/srcObject "The srcObject property of the HTMLMediaElement interface sets or returns the object which serves as the source of the media associated with the HTMLMediaElement.")

Support in one engine only.

::: support
[Firefox[🔰 42+]{title="Partial implementation."}]{.firefox
.yes}[Safari11+]{.safari .yes}[Chrome[🔰
108+]{title="Partial implementation."}]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge[🔰
108+]{title="Partial implementation."}]{.edge_blink .yes}

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

Allows the [media
element](#media-element){#location-of-the-media-resource:media-element-11}
to be assigned a [media provider
object](#media-provider-object){#location-of-the-media-resource:media-provider-object-3}.

`media`{.variable}`.`[`currentSrc`](#dom-media-currentsrc){#dom-media-currentsrc-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/currentSrc](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/currentSrc "The HTMLMediaElement.currentSrc property contains the absolute URL of the chosen media resource. This could happen, for example, if the web server selects a media file based on the resolution of the user's display. The value is an empty string if the networkState property is EMPTY.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns the
[URL](https://url.spec.whatwg.org/#concept-url){#location-of-the-media-resource:url-3
x-internal="url"} of the current [media
resource](#media-resource){#location-of-the-media-resource:media-resource-2},
if any.

Returns the empty string when there is no [media
resource](#media-resource){#location-of-the-media-resource:media-resource-3},
or it doesn\'t have a
[URL](https://url.spec.whatwg.org/#concept-url){#location-of-the-media-resource:url-4
x-internal="url"}.

The [`currentSrc`]{#dom-media-currentsrc .dfn dfn-for="HTMLMediaElement"
dfn-type="attribute"} IDL attribute must initially be set to the empty
string. Its value is changed by the [resource selection
algorithm](#concept-media-load-algorithm){#location-of-the-media-resource:concept-media-load-algorithm-2}
defined below.

The [`srcObject`]{#dom-media-srcobject .dfn dfn-for="HTMLMediaElement"
dfn-type="attribute"} IDL attribute, on getting, must return the
element\'s [assigned media provider
object](#assigned-media-provider-object){#location-of-the-media-resource:assigned-media-provider-object-2},
if any, or null otherwise. On setting, it must set the element\'s
[assigned media provider
object](#assigned-media-provider-object){#location-of-the-media-resource:assigned-media-provider-object-3}
to the new value, and then invoke the element\'s [media element load
algorithm](#media-element-load-algorithm){#location-of-the-media-resource:media-element-load-algorithm-2}.

There are three ways to specify a [media
resource](#media-resource){#location-of-the-media-resource:media-resource-4}:
the
[`srcObject`{#location-of-the-media-resource:dom-media-srcobject}](#dom-media-srcobject)
IDL attribute, the
[`src`{#location-of-the-media-resource:attr-media-src-5}](#attr-media-src)
content attribute, and
[`source`{#location-of-the-media-resource:the-source-element-2}](embedded-content.html#the-source-element)
elements. The IDL attribute takes priority, followed by the content
attribute, followed by the elements.

##### [4.8.11.3]{.secno} MIME types[](#mime-types){.self-link}

A [media resource](#media-resource){#mime-types:media-resource} can be
described in terms of its *type*, specifically a [MIME
type](https://mimesniff.spec.whatwg.org/#mime-type){#mime-types:mime-type
x-internal="mime-type"}, in some cases with a `codecs` parameter.
(Whether the `codecs` parameter is allowed or not depends on the MIME
type.) [\[RFC6381\]](references.html#refsRFC6381)

Types are usually somewhat incomplete descriptions; for example
\"`video/mpeg`\" doesn\'t say anything except what the container type
is, and even a type like
\"`video/mp4; codecs="avc1.42E01E, mp4a.40.2"`\" doesn\'t include
information like the actual bitrate (only the maximum bitrate). Thus,
given a type, a user agent can often only know whether it *might* be
able to play media of that type (with varying levels of confidence), or
whether it definitely *cannot* play media of that type.

[A type that the user agent knows it cannot
render]{#a-type-that-the-user-agent-knows-it-cannot-render .dfn} is one
that describes a resource that the user agent definitely does not
support, for example because it doesn\'t recognize the container type,
or it doesn\'t support the listed codecs.

The [MIME
type](https://mimesniff.spec.whatwg.org/#mime-type){#mime-types:mime-type-2
x-internal="mime-type"}
\"[`application/octet-stream`{#mime-types:application/octet-stream}](https://www.rfc-editor.org/rfc/rfc2046#section-4.5.1){x-internal="application/octet-stream"}\"
with no parameters is never [a type that the user agent knows it cannot
render](#a-type-that-the-user-agent-knows-it-cannot-render){#mime-types:a-type-that-the-user-agent-knows-it-cannot-render}.
User agents must treat that type as equivalent to the lack of any
explicit [Content-Type
metadata](urls-and-fetching.html#content-type){#mime-types:content-type}
when it is used to label a potential [media
resource](#media-resource){#mime-types:media-resource-2}.

Only the [MIME
type](https://mimesniff.spec.whatwg.org/#mime-type){#mime-types:mime-type-3
x-internal="mime-type"}
\"[`application/octet-stream`{#mime-types:application/octet-stream-2}](https://www.rfc-editor.org/rfc/rfc2046#section-4.5.1){x-internal="application/octet-stream"}\"
with no parameters is special-cased here; if any parameter appears with
it, it will be treated just like any other [MIME
type](https://mimesniff.spec.whatwg.org/#mime-type){#mime-types:mime-type-4
x-internal="mime-type"}. This is a deviation from the rule that unknown
[MIME
type](https://mimesniff.spec.whatwg.org/#mime-type){#mime-types:mime-type-5
x-internal="mime-type"} parameters should be ignored.

`media`{.variable}`.`[`canPlayType`](#dom-navigator-canplaytype){#dom-navigator-canplaytype-dev}`(``type`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/canPlayType](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/canPlayType "The HTMLMediaElement method canPlayType() reports how likely it is that the current browser will be able to play media of a given MIME type.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns the empty string (a negative response), \"maybe\", or
\"probably\" based on how confident the user agent is that it can play
media resources of the given type.

The [`canPlayType(``type`{.variable}`)`]{#dom-navigator-canplaytype .dfn
dfn-for="HTMLMediaElement" dfn-type="method"} method must return [the
empty string]{#dom-canplaytyperesult-nil .dfn} if `type`{.variable} is
[a type that the user agent knows it cannot
render](#a-type-that-the-user-agent-knows-it-cannot-render){#mime-types:a-type-that-the-user-agent-knows-it-cannot-render-2}
or is the type
\"[`application/octet-stream`{#mime-types:application/octet-stream-3}](https://www.rfc-editor.org/rfc/rfc2046#section-4.5.1){x-internal="application/octet-stream"}\";
it must return \"[`probably`]{#dom-canplaytyperesult-probably .dfn
dfn-for="CanPlayTypeResult" dfn-type="enum-value"}\" if the user agent
is confident that the type represents a [media
resource](#media-resource){#mime-types:media-resource-3} that it can
render if used in with this
[`audio`{#mime-types:the-audio-element}](#the-audio-element) or
[`video`{#mime-types:the-video-element}](#the-video-element) element;
and it must return \"[`maybe`]{#dom-canplaytyperesult-maybe .dfn
dfn-for="CanPlayTypeResult" dfn-type="enum-value"}\" otherwise.
Implementers are encouraged to return
\"[`maybe`{#mime-types:dom-canplaytyperesult-maybe}](#dom-canplaytyperesult-maybe)\"
unless the type can be confidently established as being supported or
not. Generally, a user agent should never return
\"[`probably`{#mime-types:dom-canplaytyperesult-probably}](#dom-canplaytyperesult-probably)\"
for a type that allows the `codecs` parameter if that parameter is not
present.

::: example
This script tests to see if the user agent supports a (fictional) new
format to dynamically decide whether to use a
[`video`{#mime-types:the-video-element-2}](#the-video-element) element:

``` html
<section id="video">
 <p><a href="playing-cats.nfv">Download video</a></p>
</section>
<script>
 const videoSection = document.getElementById('video');
 const videoElement = document.createElement('video');
 const support = videoElement.canPlayType('video/x-new-fictional-format;codecs="kittens,bunnies"');
 if (support === "probably") {
   videoElement.setAttribute("src", "playing-cats.nfv");
   videoSection.replaceChildren(videoElement);
 }
</script>
```
:::

The
[`type`{#mime-types:attr-source-type}](embedded-content.html#attr-source-type)
attribute of the
[`source`{#mime-types:the-source-element}](embedded-content.html#the-source-element)
element allows the user agent to avoid downloading resources that use
formats it cannot render.

##### [4.8.11.4]{.secno} Network states[](#network-states){.self-link}

`media`{.variable}`.`[`networkState`](#dom-media-networkstate){#dom-media-networkstate-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/networkState](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/networkState "The HTMLMediaElement.networkState property indicates the current state of the fetching of media over the network.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android4+]{.firefox_android .yes}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns the current state of network activity for the element, from the
codes in the list below.

As [media elements](#media-element){#network-states:media-element}
interact with the network, their current network activity is represented
by the [`networkState`]{#dom-media-networkstate .dfn
dfn-for="HTMLMediaElement" dfn-type="attribute"} attribute. On getting,
it must return the current network state of the element, which must be
one of the following values:

[`NETWORK_EMPTY`]{#dom-media-network_empty .dfn dfn-for="HTMLMediaElement" dfn-type="const"} (numeric value 0)
:   The element has not yet been initialized. All attributes are in
    their initial states.

[`NETWORK_IDLE`]{#dom-media-network_idle .dfn dfn-for="HTMLMediaElement" dfn-type="const"} (numeric value 1)
:   The element\'s [resource selection
    algorithm](#concept-media-load-algorithm){#network-states:concept-media-load-algorithm}
    is active and has selected a
    [resource](#media-resource){#network-states:media-resource}, but it
    is not actually using the network at this time.

[`NETWORK_LOADING`]{#dom-media-network_loading .dfn dfn-for="HTMLMediaElement" dfn-type="const"} (numeric value 2)
:   The user agent is actively trying to download data.

[`NETWORK_NO_SOURCE`]{#dom-media-network_no_source .dfn dfn-for="HTMLMediaElement" dfn-type="const"} (numeric value 3)
:   The element\'s [resource selection
    algorithm](#concept-media-load-algorithm){#network-states:concept-media-load-algorithm-2}
    is active, but it has not yet found a
    [resource](#media-resource){#network-states:media-resource-2} to
    use.

The [resource selection
algorithm](#concept-media-load-algorithm){#network-states:concept-media-load-algorithm-3}
defined below describes exactly when the
[`networkState`{#network-states:dom-media-networkstate}](#dom-media-networkstate)
attribute changes value and what events fire to indicate changes in this
state.

##### [4.8.11.5]{.secno} Loading the media resource[](#loading-the-media-resource){.self-link}

`media`{.variable}`.`[`load`](#dom-media-load){#dom-media-load-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/load](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/load "The HTMLMediaElement method load() resets the media element to its initial state and begins the process of selecting a media source and loading the media in preparation for playback to begin at the beginning.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Causes the element to reset and start selecting and loading a new [media
resource](#media-resource){#loading-the-media-resource:media-resource}
from scratch.

All [media
elements](#media-element){#loading-the-media-resource:media-element}
have a [can autoplay flag]{#can-autoplay-flag .dfn}, which must begin in
the true state, and a [delaying-the-load-event
flag]{#delaying-the-load-event-flag .dfn}, which must begin in the false
state. While the [delaying-the-load-event
flag](#delaying-the-load-event-flag){#loading-the-media-resource:delaying-the-load-event-flag}
is true, the element must [delay the load
event](parsing.html#delay-the-load-event){#loading-the-media-resource:delay-the-load-event}
of its document.

When the [`load()`]{#dom-media-load .dfn dfn-for="HTMLMediaElement"
dfn-type="method"} method on a [media
element](#media-element){#loading-the-media-resource:media-element-2} is
invoked, the user agent must run the [media element load
algorithm](#media-element-load-algorithm){#loading-the-media-resource:media-element-load-algorithm}.

A [media
element](#media-element){#loading-the-media-resource:media-element-3}
has an associated boolean [is currently stalled]{#is-currently-stalled
.dfn}, which is initially false.

The [media element load algorithm]{#media-element-load-algorithm .dfn}
consists of the following steps.

1.  Set this element\'s [is currently
    stalled](#is-currently-stalled){#loading-the-media-resource:is-currently-stalled}
    to false.

2.  Abort any already-running instance of the [resource selection
    algorithm](#concept-media-load-algorithm){#loading-the-media-resource:concept-media-load-algorithm}
    for this element.

3.  Let `pending tasks`{.variable} be a list of all
    [tasks](webappapis.html#concept-task){#loading-the-media-resource:concept-task}
    from the [media
    element](#media-element){#loading-the-media-resource:media-element-4}\'s
    [media element event task
    source](#media-element-event-task-source){#loading-the-media-resource:media-element-event-task-source}
    in one of the [task
    queues](webappapis.html#task-queue){#loading-the-media-resource:task-queue}.

4.  For each task in `pending tasks`{.variable} that would [resolve
    pending play
    promises](#resolve-pending-play-promises){#loading-the-media-resource:resolve-pending-play-promises}
    or [reject pending play
    promises](#reject-pending-play-promises){#loading-the-media-resource:reject-pending-play-promises},
    immediately resolve or reject those promises in the order the
    corresponding tasks were queued.

5.  Remove each
    [task](webappapis.html#concept-task){#loading-the-media-resource:concept-task-2}
    in `pending tasks`{.variable} from its [task
    queue](webappapis.html#task-queue){#loading-the-media-resource:task-queue-2}.

    Basically, pending events and callbacks are discarded and promises
    in-flight to be resolved/rejected are resolved/rejected immediately
    when the media element starts loading a new resource.

6.  If the [media
    element](#media-element){#loading-the-media-resource:media-element-5}\'s
    [`networkState`{#loading-the-media-resource:dom-media-networkstate}](#dom-media-networkstate)
    is set to
    [`NETWORK_LOADING`{#loading-the-media-resource:dom-media-network_loading}](#dom-media-network_loading)
    or
    [`NETWORK_IDLE`{#loading-the-media-resource:dom-media-network_idle}](#dom-media-network_idle),
    [queue a media element
    task](#queue-a-media-element-task){#loading-the-media-resource:queue-a-media-element-task}
    given the [media
    element](#media-element){#loading-the-media-resource:media-element-6}
    to [fire an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#loading-the-media-resource:concept-event-fire
    x-internal="concept-event-fire"} named
    [`abort`{#loading-the-media-resource:event-media-abort}](#event-media-abort)
    at the [media
    element](#media-element){#loading-the-media-resource:media-element-7}.

7.  If the [media
    element](#media-element){#loading-the-media-resource:media-element-8}\'s
    [`networkState`{#loading-the-media-resource:dom-media-networkstate-2}](#dom-media-networkstate)
    is not set to
    [`NETWORK_EMPTY`{#loading-the-media-resource:dom-media-network_empty}](#dom-media-network_empty),
    then:

    1.  [Queue a media element
        task](#queue-a-media-element-task){#loading-the-media-resource:queue-a-media-element-task-2}
        given the [media
        element](#media-element){#loading-the-media-resource:media-element-9}
        to [fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#loading-the-media-resource:concept-event-fire-2
        x-internal="concept-event-fire"} named
        [`emptied`{#loading-the-media-resource:event-media-emptied}](#event-media-emptied)
        at the [media
        element](#media-element){#loading-the-media-resource:media-element-10}.

    2.  If a fetching process is in progress for the [media
        element](#media-element){#loading-the-media-resource:media-element-11},
        the user agent should stop it.

    3.  If the [media
        element](#media-element){#loading-the-media-resource:media-element-12}\'s
        [assigned media provider
        object](#assigned-media-provider-object){#loading-the-media-resource:assigned-media-provider-object}
        is a
        [`MediaSource`{#loading-the-media-resource:mediasource}](https://w3c.github.io/media-source/#idl-def-mediasource){x-internal="mediasource"}
        object, then
        [detach](https://w3c.github.io/media-source/#mediasource-detach){#loading-the-media-resource:detaching-from-a-media-element
        x-internal="detaching-from-a-media-element"} it.

    4.  [Forget the media element\'s media-resource-specific
        tracks](#forget-the-media-element's-media-resource-specific-tracks){#loading-the-media-resource:forget-the-media-element's-media-resource-specific-tracks}.

    5.  If
        [`readyState`{#loading-the-media-resource:dom-media-readystate}](#dom-media-readystate)
        is not set to
        [`HAVE_NOTHING`{#loading-the-media-resource:dom-media-have_nothing}](#dom-media-have_nothing),
        then set it to that state.

    6.  If the
        [`paused`{#loading-the-media-resource:dom-media-paused}](#dom-media-paused)
        attribute is false, then:

        1.  Set the
            [`paused`{#loading-the-media-resource:dom-media-paused-2}](#dom-media-paused)
            attribute to true.

        2.  [Take pending play
            promises](#take-pending-play-promises){#loading-the-media-resource:take-pending-play-promises}
            and [reject pending play
            promises](#reject-pending-play-promises){#loading-the-media-resource:reject-pending-play-promises-2}
            with the result and an
            [\"`AbortError`\"](https://webidl.spec.whatwg.org/#aborterror){#loading-the-media-resource:aborterror
            x-internal="aborterror"}
            [`DOMException`{#loading-the-media-resource:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    7.  If
        [`seeking`{#loading-the-media-resource:dom-media-seeking}](#dom-media-seeking)
        is true, set it to false.

    8.  Set the [current playback
        position](#current-playback-position){#loading-the-media-resource:current-playback-position}
        to 0.

        Set the [official playback
        position](#official-playback-position){#loading-the-media-resource:official-playback-position}
        to 0.

        If this changed the [official playback
        position](#official-playback-position){#loading-the-media-resource:official-playback-position-2},
        then [queue a media element
        task](#queue-a-media-element-task){#loading-the-media-resource:queue-a-media-element-task-3}
        given the [media
        element](#media-element){#loading-the-media-resource:media-element-13}
        to [fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#loading-the-media-resource:concept-event-fire-3
        x-internal="concept-event-fire"} named
        [`timeupdate`{#loading-the-media-resource:event-media-timeupdate}](#event-media-timeupdate)
        at the [media
        element](#media-element){#loading-the-media-resource:media-element-14}.

    9.  Set the [timeline
        offset](#timeline-offset){#loading-the-media-resource:timeline-offset}
        to Not-a-Number (NaN).

    10. Update the
        [`duration`{#loading-the-media-resource:dom-media-duration}](#dom-media-duration)
        attribute to Not-a-Number (NaN).

        The user agent [will not](#durationChange) fire a
        [`durationchange`{#loading-the-media-resource:event-media-durationchange}](#event-media-durationchange)
        event for this particular change of the duration.

8.  Set the
    [`playbackRate`{#loading-the-media-resource:dom-media-playbackrate}](#dom-media-playbackrate)
    attribute to the value of the
    [`defaultPlaybackRate`{#loading-the-media-resource:dom-media-defaultplaybackrate}](#dom-media-defaultplaybackrate)
    attribute.

9.  Set the
    [`error`{#loading-the-media-resource:dom-media-error}](#dom-media-error)
    attribute to null and the [can autoplay
    flag](#can-autoplay-flag){#loading-the-media-resource:can-autoplay-flag}
    to true.

10. Invoke the [media
    element](#media-element){#loading-the-media-resource:media-element-15}\'s
    [resource selection
    algorithm](#concept-media-load-algorithm){#loading-the-media-resource:concept-media-load-algorithm-2}.

11. Playback of any previously playing [media
    resource](#media-resource){#loading-the-media-resource:media-resource-2}
    for this element stops.

The [resource selection algorithm]{#concept-media-load-algorithm .dfn}
for a [media
element](#media-element){#loading-the-media-resource:media-element-16}
is as follows. This algorithm is always invoked as part of a
[task](webappapis.html#concept-task){#loading-the-media-resource:concept-task-3},
but one of the first steps in the algorithm is to return and continue
running the remaining steps [in
parallel](infrastructure.html#in-parallel){#loading-the-media-resource:in-parallel}.
In addition, this algorithm interacts closely with the [event
loop](webappapis.html#event-loop){#loading-the-media-resource:event-loop}
mechanism; in particular, it has [synchronous
sections](webappapis.html#synchronous-section){#loading-the-media-resource:synchronous-section}
(which are triggered as part of the [event
loop](webappapis.html#event-loop){#loading-the-media-resource:event-loop-2}
algorithm). Steps in such sections are marked with ⌛.

1.  Set the element\'s
    [`networkState`{#loading-the-media-resource:dom-media-networkstate-3}](#dom-media-networkstate)
    attribute to the
    [`NETWORK_NO_SOURCE`{#loading-the-media-resource:dom-media-network_no_source}](#dom-media-network_no_source)
    value.

2.  Set the element\'s [show poster
    flag](#show-poster-flag){#loading-the-media-resource:show-poster-flag}
    to true.

3.  Set the [media
    element](#media-element){#loading-the-media-resource:media-element-17}\'s
    [delaying-the-load-event
    flag](#delaying-the-load-event-flag){#loading-the-media-resource:delaying-the-load-event-flag-2}
    to true (this [delays the load
    event](parsing.html#delay-the-load-event){#loading-the-media-resource:delay-the-load-event-2}).

4.  [Await a stable
    state](webappapis.html#await-a-stable-state){#loading-the-media-resource:await-a-stable-state},
    allowing the
    [task](webappapis.html#concept-task){#loading-the-media-resource:concept-task-4}
    that invoked this algorithm to continue. The [synchronous
    section](webappapis.html#synchronous-section){#loading-the-media-resource:synchronous-section-2}
    consists of all the remaining steps of this algorithm until the
    algorithm says the [synchronous
    section](webappapis.html#synchronous-section){#loading-the-media-resource:synchronous-section-3}
    has ended. (Steps in [synchronous
    sections](webappapis.html#synchronous-section){#loading-the-media-resource:synchronous-section-4}
    are marked with ⌛.)

5.  ⌛ If the [media
    element](#media-element){#loading-the-media-resource:media-element-18}\'s
    [blocked-on-parser](#blocked-on-parser){#loading-the-media-resource:blocked-on-parser}
    flag is false, then [populate the list of pending text
    tracks](#populate-the-list-of-pending-text-tracks){#loading-the-media-resource:populate-the-list-of-pending-text-tracks}.

6.  ⌛ If the [media
    element](#media-element){#loading-the-media-resource:media-element-19}
    has an [assigned media provider
    object](#assigned-media-provider-object){#loading-the-media-resource:assigned-media-provider-object-2},
    then let `mode`{.variable} be *object*.

    ⌛ Otherwise, if the [media
    element](#media-element){#loading-the-media-resource:media-element-20}
    has no [assigned media provider
    object](#assigned-media-provider-object){#loading-the-media-resource:assigned-media-provider-object-3}
    but has a
    [`src`{#loading-the-media-resource:attr-media-src}](#attr-media-src)
    attribute, then let `mode`{.variable} be *attribute*.

    ⌛ Otherwise, if the [media
    element](#media-element){#loading-the-media-resource:media-element-21}
    does not have an [assigned media provider
    object](#assigned-media-provider-object){#loading-the-media-resource:assigned-media-provider-object-4}
    and does not have a
    [`src`{#loading-the-media-resource:attr-media-src-2}](#attr-media-src)
    attribute, but does have a
    [`source`{#loading-the-media-resource:the-source-element}](embedded-content.html#the-source-element)
    element child, then let `mode`{.variable} be *children* and let
    `candidate`{.variable} be the first such
    [`source`{#loading-the-media-resource:the-source-element-2}](embedded-content.html#the-source-element)
    element child in [tree
    order](https://dom.spec.whatwg.org/#concept-tree-order){#loading-the-media-resource:tree-order
    x-internal="tree-order"}.

    ⌛ Otherwise, the [media
    element](#media-element){#loading-the-media-resource:media-element-22}
    has no [assigned media provider
    object](#assigned-media-provider-object){#loading-the-media-resource:assigned-media-provider-object-5}
    and has neither a
    [`src`{#loading-the-media-resource:attr-media-src-3}](#attr-media-src)
    attribute nor a
    [`source`{#loading-the-media-resource:the-source-element-3}](embedded-content.html#the-source-element)
    element child:

    1.  ⌛ Set the
        [`networkState`{#loading-the-media-resource:dom-media-networkstate-4}](#dom-media-networkstate)
        to
        [`NETWORK_EMPTY`{#loading-the-media-resource:dom-media-network_empty-2}](#dom-media-network_empty).

    2.  ⌛ Set the element\'s [delaying-the-load-event
        flag](#delaying-the-load-event-flag){#loading-the-media-resource:delaying-the-load-event-flag-3}
        to false. This stops [delaying the load
        event](parsing.html#delay-the-load-event){#loading-the-media-resource:delay-the-load-event-3}.

    3.  End the [synchronous
        section](webappapis.html#synchronous-section){#loading-the-media-resource:synchronous-section-5}
        and return.

7.  ⌛ Set the [media
    element](#media-element){#loading-the-media-resource:media-element-23}\'s
    [`networkState`{#loading-the-media-resource:dom-media-networkstate-5}](#dom-media-networkstate)
    to
    [`NETWORK_LOADING`{#loading-the-media-resource:dom-media-network_loading-2}](#dom-media-network_loading).

8.  ⌛ [Queue a media element
    task](#queue-a-media-element-task){#loading-the-media-resource:queue-a-media-element-task-4}
    given the [media
    element](#media-element){#loading-the-media-resource:media-element-24}
    to [fire an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#loading-the-media-resource:concept-event-fire-4
    x-internal="concept-event-fire"} named
    [`loadstart`{#loading-the-media-resource:event-media-loadstart}](#event-media-loadstart)
    at the [media
    element](#media-element){#loading-the-media-resource:media-element-25}.

9.  Run the appropriate steps from the following list:

    If `mode`{.variable} is *object*

    :   1.  ⌛ Set the
            [`currentSrc`{#loading-the-media-resource:dom-media-currentsrc}](#dom-media-currentsrc)
            attribute to the empty string.

        2.  End the [synchronous
            section](webappapis.html#synchronous-section){#loading-the-media-resource:synchronous-section-6},
            continuing the remaining steps [in
            parallel](infrastructure.html#in-parallel){#loading-the-media-resource:in-parallel-2}.

        3.  Run the [resource fetch
            algorithm](#concept-media-load-resource){#loading-the-media-resource:concept-media-load-resource}
            with the [assigned media provider
            object](#assigned-media-provider-object){#loading-the-media-resource:assigned-media-provider-object-6}.
            If that algorithm returns without aborting *this* one, then
            the load failed.

        4.  *Failed with media provider*: Reaching this step indicates
            that the media resource failed to load. [Take pending play
            promises](#take-pending-play-promises){#loading-the-media-resource:take-pending-play-promises-2}
            and [queue a media element
            task](#queue-a-media-element-task){#loading-the-media-resource:queue-a-media-element-task-5}
            given the [media
            element](#media-element){#loading-the-media-resource:media-element-26}
            to run the [dedicated media source failure
            steps](#dedicated-media-source-failure-steps){#loading-the-media-resource:dedicated-media-source-failure-steps}
            with the result.

        5.  Wait for the
            [task](webappapis.html#concept-task){#loading-the-media-resource:concept-task-5}
            queued by the previous step to have executed.

        6.  Return. The element won\'t attempt to load another resource
            until this algorithm is triggered again.

    If `mode`{.variable} is *attribute*

    :   1.  ⌛ If the
            [`src`{#loading-the-media-resource:attr-media-src-4}](#attr-media-src)
            attribute\'s value is the empty string, then end the
            [synchronous
            section](webappapis.html#synchronous-section){#loading-the-media-resource:synchronous-section-7},
            and jump down to the *failed with attribute* step below.

        2.  ⌛ Let `urlRecord`{.variable} be the result of
            [encoding-parsing a
            URL](urls-and-fetching.html#encoding-parsing-a-url){#loading-the-media-resource:encoding-parsing-a-url}
            given the
            [`src`{#loading-the-media-resource:attr-media-src-5}](#attr-media-src)
            attribute\'s value, relative to the [media
            element](#media-element){#loading-the-media-resource:media-element-27}\'s
            [node
            document](https://dom.spec.whatwg.org/#concept-node-document){#loading-the-media-resource:node-document
            x-internal="node-document"} when the
            [`src`{#loading-the-media-resource:attr-media-src-6}](#attr-media-src)
            attribute was last changed.

        3.  ⌛ If `urlRecord`{.variable} is not failure, then set the
            [`currentSrc`{#loading-the-media-resource:dom-media-currentsrc-2}](#dom-media-currentsrc)
            attribute to the result of applying the [URL
            serializer](https://url.spec.whatwg.org/#concept-url-serializer){#loading-the-media-resource:concept-url-serializer
            x-internal="concept-url-serializer"} to
            `urlRecord`{.variable}.

        4.  End the [synchronous
            section](webappapis.html#synchronous-section){#loading-the-media-resource:synchronous-section-8},
            continuing the remaining steps [in
            parallel](infrastructure.html#in-parallel){#loading-the-media-resource:in-parallel-3}.

        5.  If `urlRecord`{.variable} is not failure, then run the
            [resource fetch
            algorithm](#concept-media-load-resource){#loading-the-media-resource:concept-media-load-resource-2}
            with `urlRecord`{.variable}. If that algorithm returns
            without aborting *this* one, then the load failed.

        6.  *Failed with attribute*: Reaching this step indicates that
            the media resource failed to load or that
            `urlRecord`{.variable} is failure. [Take pending play
            promises](#take-pending-play-promises){#loading-the-media-resource:take-pending-play-promises-3}
            and [queue a media element
            task](#queue-a-media-element-task){#loading-the-media-resource:queue-a-media-element-task-6}
            given the [media
            element](#media-element){#loading-the-media-resource:media-element-28}
            to run the [dedicated media source failure
            steps](#dedicated-media-source-failure-steps){#loading-the-media-resource:dedicated-media-source-failure-steps-2}
            with the result.

        7.  Wait for the
            [task](webappapis.html#concept-task){#loading-the-media-resource:concept-task-6}
            queued by the previous step to have executed.

        8.  Return. The element won\'t attempt to load another resource
            until this algorithm is triggered again.

    Otherwise (`mode`{.variable} is *children*)

    :   1.  ⌛ Let `pointer`{.variable} be a position defined by two
            adjacent nodes in the [media
            element](#media-element){#loading-the-media-resource:media-element-29}\'s
            child list, treating the start of the list (before the first
            child in the list, if any) and end of the list (after the
            last child in the list, if any) as nodes in their own right.
            One node is the node before `pointer`{.variable}, and the
            other node is the node after `pointer`{.variable}.
            Initially, let `pointer`{.variable} be the position between
            the `candidate`{.variable} node and the next node, if there
            are any, or the end of the list, if it is the last node.

            As nodes are
            [inserted](https://dom.spec.whatwg.org/#concept-node-insert-ext){#loading-the-media-resource:concept-node-insert-ext
            x-internal="concept-node-insert-ext"},
            [removed](https://dom.spec.whatwg.org/#concept-node-remove-ext){#loading-the-media-resource:concept-node-remove-ext
            x-internal="concept-node-remove-ext"}, and
            [moved](https://dom.spec.whatwg.org/#concept-node-move-ext){#loading-the-media-resource:concept-node-move-ext
            x-internal="concept-node-move-ext"} into the [media
            element](#media-element){#loading-the-media-resource:media-element-30},
            `pointer`{.variable} must be updated as follows:

            If a new node is [inserted](https://dom.spec.whatwg.org/#concept-node-insert-ext){#loading-the-media-resource:concept-node-insert-ext-2 x-internal="concept-node-insert-ext"} or [moved](https://dom.spec.whatwg.org/#concept-node-move-ext){#loading-the-media-resource:concept-node-move-ext-2 x-internal="concept-node-move-ext"} between the two nodes that define `pointer`{.variable}
            :   Let `pointer`{.variable} be the point between the node
                before `pointer`{.variable} and the new node. In other
                words, insertions at `pointer`{.variable} go after
                `pointer`{.variable}.

            If the node before `pointer`{.variable} is removed
            :   Let `pointer`{.variable} be the point between the node
                after `pointer`{.variable} and the node before the node
                after `pointer`{.variable}. In other words,
                `pointer`{.variable} doesn\'t move relative to the
                remaining nodes.

            If the node after `pointer`{.variable} is removed
            :   Let `pointer`{.variable} be the point between the node
                before `pointer`{.variable} and the node after the node
                before `pointer`{.variable}. Just as with the previous
                case, `pointer`{.variable} doesn\'t move relative to the
                remaining nodes.

            Other changes don\'t affect `pointer`{.variable}.

        2.  ⌛ *Process candidate*: If `candidate`{.variable} does not
            have a
            [`src`{#loading-the-media-resource:attr-source-src}](embedded-content.html#attr-source-src)
            attribute, or if its
            [`src`{#loading-the-media-resource:attr-source-src-2}](embedded-content.html#attr-source-src)
            attribute\'s value is the empty string, then end the
            [synchronous
            section](webappapis.html#synchronous-section){#loading-the-media-resource:synchronous-section-9},
            and jump down to the *failed with elements* step below.

        3.  ⌛ If `candidate`{.variable} has a
            [`media`{#loading-the-media-resource:attr-source-media}](embedded-content.html#attr-source-media)
            attribute whose value does not [match the
            environment](common-microsyntaxes.html#matches-the-environment){#loading-the-media-resource:matches-the-environment},
            then end the [synchronous
            section](webappapis.html#synchronous-section){#loading-the-media-resource:synchronous-section-10},
            and jump down to the *failed with elements* step below.

        4.  ⌛ Let `urlRecord`{.variable} be the result of
            [encoding-parsing a
            URL](urls-and-fetching.html#encoding-parsing-a-url){#loading-the-media-resource:encoding-parsing-a-url-2}
            given `candidate`{.variable}\'s
            [`src`{#loading-the-media-resource:attr-media-src-7}](#attr-media-src)
            attribute\'s value, relative to `candidate`{.variable}\'s
            [node
            document](https://dom.spec.whatwg.org/#concept-node-document){#loading-the-media-resource:node-document-2
            x-internal="node-document"} when the
            [`src`{#loading-the-media-resource:attr-media-src-8}](#attr-media-src)
            attribute was last changed.

        5.  ⌛ If `urlRecord`{.variable} is failure, then end the
            [synchronous
            section](webappapis.html#synchronous-section){#loading-the-media-resource:synchronous-section-11},
            and jump down to the *failed with elements* step below.

        6.  ⌛ If `candidate`{.variable} has a
            [`type`{#loading-the-media-resource:attr-source-type}](embedded-content.html#attr-source-type)
            attribute whose value, when parsed as a [MIME
            type](https://mimesniff.spec.whatwg.org/#mime-type){#loading-the-media-resource:mime-type
            x-internal="mime-type"} (including any codecs described by
            the `codecs` parameter, for types that define that
            parameter), represents [a type that the user agent knows it
            cannot
            render](#a-type-that-the-user-agent-knows-it-cannot-render){#loading-the-media-resource:a-type-that-the-user-agent-knows-it-cannot-render},
            then end the [synchronous
            section](webappapis.html#synchronous-section){#loading-the-media-resource:synchronous-section-12},
            and jump down to the *failed with elements* step below.

        7.  ⌛ Set the
            [`currentSrc`{#loading-the-media-resource:dom-media-currentsrc-3}](#dom-media-currentsrc)
            attribute to the result of applying the [URL
            serializer](https://url.spec.whatwg.org/#concept-url-serializer){#loading-the-media-resource:concept-url-serializer-2
            x-internal="concept-url-serializer"} to
            `urlRecord`{.variable}.

        8.  End the [synchronous
            section](webappapis.html#synchronous-section){#loading-the-media-resource:synchronous-section-13},
            continuing the remaining steps [in
            parallel](infrastructure.html#in-parallel){#loading-the-media-resource:in-parallel-4}.

        9.  Run the [resource fetch
            algorithm](#concept-media-load-resource){#loading-the-media-resource:concept-media-load-resource-3}
            with `urlRecord`{.variable}. If that algorithm returns
            without aborting *this* one, then the load failed.

        10. *Failed with elements*: [Queue a media element
            task](#queue-a-media-element-task){#loading-the-media-resource:queue-a-media-element-task-7}
            given the [media
            element](#media-element){#loading-the-media-resource:media-element-31}
            to [fire an
            event](https://dom.spec.whatwg.org/#concept-event-fire){#loading-the-media-resource:concept-event-fire-5
            x-internal="concept-event-fire"} named
            [`error`{#loading-the-media-resource:event-source-error}](#event-source-error)
            at `candidate`{.variable}.

        11. [Await a stable
            state](webappapis.html#await-a-stable-state){#loading-the-media-resource:await-a-stable-state-2}.
            The [synchronous
            section](webappapis.html#synchronous-section){#loading-the-media-resource:synchronous-section-14}
            consists of all the remaining steps of this algorithm until
            the algorithm says the [synchronous
            section](webappapis.html#synchronous-section){#loading-the-media-resource:synchronous-section-15}
            has ended. (Steps in [synchronous
            sections](webappapis.html#synchronous-section){#loading-the-media-resource:synchronous-section-16}
            are marked with ⌛.)

        12. ⌛ [Forget the media element\'s media-resource-specific
            tracks](#forget-the-media-element's-media-resource-specific-tracks){#loading-the-media-resource:forget-the-media-element's-media-resource-specific-tracks-2}.

        13. ⌛ *Find next candidate*: Let `candidate`{.variable} be
            null.

        14. ⌛ *Search loop*: If the node after `pointer`{.variable} is
            the end of the list, then jump to the *waiting* step below.

        15. ⌛ If the node after `pointer`{.variable} is a
            [`source`{#loading-the-media-resource:the-source-element-4}](embedded-content.html#the-source-element)
            element, let `candidate`{.variable} be that element.

        16. ⌛ Advance `pointer`{.variable} so that the node before
            `pointer`{.variable} is now the node that was after
            `pointer`{.variable}, and the node after
            `pointer`{.variable} is the node after the node that used to
            be after `pointer`{.variable}, if any.

        17. ⌛ If `candidate`{.variable} is null, jump back to the
            *search loop* step. Otherwise, jump back to the *process
            candidate* step.

        18. ⌛ *Waiting*: Set the element\'s
            [`networkState`{#loading-the-media-resource:dom-media-networkstate-6}](#dom-media-networkstate)
            attribute to the
            [`NETWORK_NO_SOURCE`{#loading-the-media-resource:dom-media-network_no_source-2}](#dom-media-network_no_source)
            value.

        19. ⌛ Set the element\'s [show poster
            flag](#show-poster-flag){#loading-the-media-resource:show-poster-flag-2}
            to true.

        20. ⌛ [Queue a media element
            task](#queue-a-media-element-task){#loading-the-media-resource:queue-a-media-element-task-8}
            given the [media
            element](#media-element){#loading-the-media-resource:media-element-32}
            to set the element\'s [delaying-the-load-event
            flag](#delaying-the-load-event-flag){#loading-the-media-resource:delaying-the-load-event-flag-4}
            to false. This stops [delaying the load
            event](parsing.html#delay-the-load-event){#loading-the-media-resource:delay-the-load-event-4}.

        21. End the [synchronous
            section](webappapis.html#synchronous-section){#loading-the-media-resource:synchronous-section-17},
            continuing the remaining steps [in
            parallel](infrastructure.html#in-parallel){#loading-the-media-resource:in-parallel-5}.

        22. Wait until the node after `pointer`{.variable} is a node
            other than the end of the list. (This step might wait
            forever.)

        23. [Await a stable
            state](webappapis.html#await-a-stable-state){#loading-the-media-resource:await-a-stable-state-3}.
            The [synchronous
            section](webappapis.html#synchronous-section){#loading-the-media-resource:synchronous-section-18}
            consists of all the remaining steps of this algorithm until
            the algorithm says the [synchronous
            section](webappapis.html#synchronous-section){#loading-the-media-resource:synchronous-section-19}
            has ended. (Steps in [synchronous
            sections](webappapis.html#synchronous-section){#loading-the-media-resource:synchronous-section-20}
            are marked with ⌛.)

        24. ⌛ Set the element\'s [delaying-the-load-event
            flag](#delaying-the-load-event-flag){#loading-the-media-resource:delaying-the-load-event-flag-5}
            back to true (this [delays the load
            event](parsing.html#delay-the-load-event){#loading-the-media-resource:delay-the-load-event-5}
            again, in case it hasn\'t been fired yet).

        25. ⌛ Set the
            [`networkState`{#loading-the-media-resource:dom-media-networkstate-7}](#dom-media-networkstate)
            back to
            [`NETWORK_LOADING`{#loading-the-media-resource:dom-media-network_loading-3}](#dom-media-network_loading).

        26. ⌛ Jump back to the *find next candidate* step above.

    The [dedicated media source failure
    steps]{#dedicated-media-source-failure-steps .dfn} with a list of
    promises `promises`{.variable} are the following steps:

    1.  Set the
        [`error`{#loading-the-media-resource:dom-media-error-2}](#dom-media-error)
        attribute to the result of [creating a
        `MediaError`](#creating-a-mediaerror){#loading-the-media-resource:creating-a-mediaerror}
        with
        [`MEDIA_ERR_SRC_NOT_SUPPORTED`{#loading-the-media-resource:dom-mediaerror-media_err_src_not_supported}](#dom-mediaerror-media_err_src_not_supported).

    2.  [Forget the media element\'s media-resource-specific
        tracks](#forget-the-media-element's-media-resource-specific-tracks){#loading-the-media-resource:forget-the-media-element's-media-resource-specific-tracks-3}.

    3.  Set the element\'s
        [`networkState`{#loading-the-media-resource:dom-media-networkstate-8}](#dom-media-networkstate)
        attribute to the
        [`NETWORK_NO_SOURCE`{#loading-the-media-resource:dom-media-network_no_source-3}](#dom-media-network_no_source)
        value.

    4.  Set the element\'s [show poster
        flag](#show-poster-flag){#loading-the-media-resource:show-poster-flag-3}
        to true.

    5.  [Fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#loading-the-media-resource:concept-event-fire-6
        x-internal="concept-event-fire"} named
        [`error`{#loading-the-media-resource:event-media-error}](#event-media-error)
        at the [media
        element](#media-element){#loading-the-media-resource:media-element-33}.

    6.  [Reject pending play
        promises](#reject-pending-play-promises){#loading-the-media-resource:reject-pending-play-promises-3}
        with `promises`{.variable} and a
        [\"`NotSupportedError`\"](https://webidl.spec.whatwg.org/#notsupportederror){#loading-the-media-resource:notsupportederror
        x-internal="notsupportederror"}
        [`DOMException`{#loading-the-media-resource:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    7.  Set the element\'s [delaying-the-load-event
        flag](#delaying-the-load-event-flag){#loading-the-media-resource:delaying-the-load-event-flag-6}
        to false. This stops [delaying the load
        event](parsing.html#delay-the-load-event){#loading-the-media-resource:delay-the-load-event-6}.

To [verify a media response]{#verify-a-media-response .dfn} given a
[response](https://fetch.spec.whatwg.org/#concept-response){#loading-the-media-resource:concept-response
x-internal="concept-response"} `response`{.variable}, a [media
resource](#media-resource){#loading-the-media-resource:media-resource-3}
`resource`{.variable}, and \"`entire resource`\" or a (number, number or
\"`until end`\") tuple `byteRange`{.variable}:

1.  If `response`{.variable} is a [network
    error](https://fetch.spec.whatwg.org/#concept-network-error){#loading-the-media-resource:network-error
    x-internal="network-error"}, then return false.

2.  If `byteRange`{.variable} is \"`entire resource`\", then return
    true.

3.  Let `internalResponse`{.variable} be `response`{.variable}\'s
    [unsafe
    response](urls-and-fetching.html#unsafe-response){#loading-the-media-resource:unsafe-response}.

4.  If `internalResponse`{.variable}\'s
    [status](https://fetch.spec.whatwg.org/#concept-response-status){#loading-the-media-resource:concept-response-status
    x-internal="concept-response-status"} is 200, then return true.

5.  If `internalResponse`{.variable}\'s
    [status](https://fetch.spec.whatwg.org/#concept-response-status){#loading-the-media-resource:concept-response-status-2
    x-internal="concept-response-status"} is not 206, then return false.

6.  If the result of [extracting content-range
    values](https://wicg.github.io/background-fetch/#extract-content-range-values){#loading-the-media-resource:extract-content-range-values
    x-internal="extract-content-range-values"} from
    `internalResponse`{.variable} is failure, then return false.

    Note that the extracted values are not used, and in particular are
    not compared to `byteRange`{.variable}. So this step serves as
    syntactic validation of the
    \`[`Content-Range`{#loading-the-media-resource:http-content-range}](https://httpwg.org/specs/rfc7233.html#header.content-range){x-internal="http-content-range"}\`
    header, but if the
    \`[`Content-Range`{#loading-the-media-resource:http-content-range-2}](https://httpwg.org/specs/rfc7233.html#header.content-range){x-internal="http-content-range"}\`
    values on the response mismatch the
    \`[`Range`{#loading-the-media-resource:http-range}](https://httpwg.org/specs/rfc7233.html#header.range){x-internal="http-range"}\`
    values on the request, that is not considered a failure.

7.  Let `origin`{.variable} be \"`rewritten`\" if
    `internalResponse`{.variable}\'s
    [URL](https://fetch.spec.whatwg.org/#concept-response-url){#loading-the-media-resource:concept-response-url
    x-internal="concept-response-url"} is null; otherwise
    `internalResponse`{.variable}\'s
    [URL](https://fetch.spec.whatwg.org/#concept-response-url){#loading-the-media-resource:concept-response-url-2
    x-internal="concept-response-url"}\'s
    [origin](https://url.spec.whatwg.org/#concept-url-origin){#loading-the-media-resource:concept-url-origin
    x-internal="concept-url-origin"}.

8.  Let `previousOrigin`{.variable} be `resource`{.variable}\'s
    [origin](#media-resource-origin){#loading-the-media-resource:media-resource-origin}.

9.  If any of the following are true:

    - `previousOrigin`{.variable} is \"`none`\";

    - `origin`{.variable} and `previousOrigin`{.variable} are
      \"`rewritten`\"; or

    - `origin`{.variable} and `previousOrigin`{.variable} are
      [origins](browsers.html#concept-origin){#loading-the-media-resource:concept-origin},
      and `origin`{.variable} is [same
      origin](browsers.html#same-origin){#loading-the-media-resource:same-origin}
      with `previousOrigin`{.variable},

    then set `resource`{.variable}\'s
    [origin](#media-resource-origin){#loading-the-media-resource:media-resource-origin-2}
    to `origin`{.variable}.

    Otherwise, if `response`{.variable} is
    [CORS-cross-origin](urls-and-fetching.html#cors-cross-origin){#loading-the-media-resource:cors-cross-origin},
    then return false.

    Otherwise, set `resource`{.variable}\'s
    [origin](#media-resource-origin){#loading-the-media-resource:media-resource-origin-3}
    to \"`multiple`\".

    This ensures that opaque responses with range headers do not leak
    information by being patched together with other responses from
    different origins.

10. Return true.

The [resource fetch algorithm]{#concept-media-load-resource .dfn
export=""} for a [media
element](#media-element){#loading-the-media-resource:media-element-34}
and a given [URL
record](https://url.spec.whatwg.org/#concept-url){#loading-the-media-resource:url-record
x-internal="url-record"} or [media provider
object](#media-provider-object){#loading-the-media-resource:media-provider-object}
is as follows:

1.  Let `mode`{.variable} be *remote*.

2.  If the algorithm was invoked with [media provider
    object](#media-provider-object){#loading-the-media-resource:media-provider-object-2},
    then set `mode`{.variable} to *local*.

    Otherwise:

    1.  Let `object`{.variable} be the result of [obtaining a blob
        object](https://w3c.github.io/FileAPI/#blob-url-obtain-object){#loading-the-media-resource:blob-url-obtain-object
        x-internal="blob-url-obtain-object"} using the [URL
        record](https://url.spec.whatwg.org/#concept-url){#loading-the-media-resource:url-record-2
        x-internal="url-record"}\'s [blob URL
        entry](https://url.spec.whatwg.org/#concept-url-blob-entry){#loading-the-media-resource:concept-url-blob-entry
        x-internal="concept-url-blob-entry"} and the [media
        element](#media-element){#loading-the-media-resource:media-element-35}\'s
        [node
        document](https://dom.spec.whatwg.org/#concept-node-document){#loading-the-media-resource:node-document-3
        x-internal="node-document"}\'s [relevant settings
        object](webappapis.html#relevant-settings-object){#loading-the-media-resource:relevant-settings-object}.

    2.  If `object`{.variable} is a [media provider
        object](#media-provider-object){#loading-the-media-resource:media-provider-object-3},
        then set `mode`{.variable} to *local*.

3.  If `mode`{.variable} is *remote*, then let the
    `current media resource`{.variable} be the resource given by the
    [URL
    record](https://url.spec.whatwg.org/#concept-url){#loading-the-media-resource:url-record-3
    x-internal="url-record"} passed to this algorithm; otherwise, let
    the `current media resource`{.variable} be the resource given by the
    [media provider
    object](#media-provider-object){#loading-the-media-resource:media-provider-object-4}.
    Either way, the `current media resource`{.variable} is now the
    element\'s [media
    resource](#media-resource){#loading-the-media-resource:media-resource-4}.

4.  Remove all [media-resource-specific text
    tracks](#media-resource-specific-text-track){#loading-the-media-resource:media-resource-specific-text-track}
    from the [media
    element](#media-element){#loading-the-media-resource:media-element-36}\'s
    [list of pending text
    tracks](#list-of-pending-text-tracks){#loading-the-media-resource:list-of-pending-text-tracks},
    if any.

5.  Run the appropriate steps from the following list:

    If `mode`{.variable} is remote

    :   1.  Optionally, run the following substeps. This is the expected
            behavior if the user agent intends to not attempt to fetch
            the resource until the user requests it explicitly (e.g. as
            a way to implement the
            [`preload`{#loading-the-media-resource:attr-media-preload}](#attr-media-preload)
            attribute\'s
            [`none`{#loading-the-media-resource:attr-media-preload-none}](#attr-media-preload-none)
            keyword).

            1.  Set the
                [`networkState`{#loading-the-media-resource:dom-media-networkstate-9}](#dom-media-networkstate)
                to
                [`NETWORK_IDLE`{#loading-the-media-resource:dom-media-network_idle-2}](#dom-media-network_idle).

            2.  [Queue a media element
                task](#queue-a-media-element-task){#loading-the-media-resource:queue-a-media-element-task-9}
                given the [media
                element](#media-element){#loading-the-media-resource:media-element-37}
                to [fire an
                event](https://dom.spec.whatwg.org/#concept-event-fire){#loading-the-media-resource:concept-event-fire-7
                x-internal="concept-event-fire"} named
                [`suspend`{#loading-the-media-resource:event-media-suspend}](#event-media-suspend)
                at the element.

            3.  [Queue a media element
                task](#queue-a-media-element-task){#loading-the-media-resource:queue-a-media-element-task-10}
                given the [media
                element](#media-element){#loading-the-media-resource:media-element-38}
                to set the element\'s [delaying-the-load-event
                flag](#delaying-the-load-event-flag){#loading-the-media-resource:delaying-the-load-event-flag-7}
                to false. This stops [delaying the load
                event](parsing.html#delay-the-load-event){#loading-the-media-resource:delay-the-load-event-7}.

            4.  Wait for the task to be run.

            5.  Wait for an
                [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#loading-the-media-resource:implementation-defined
                x-internal="implementation-defined"} event (e.g., the
                user requesting that the media element begin playback).

            6.  Set the element\'s [delaying-the-load-event
                flag](#delaying-the-load-event-flag){#loading-the-media-resource:delaying-the-load-event-flag-8}
                back to true (this [delays the load
                event](parsing.html#delay-the-load-event){#loading-the-media-resource:delay-the-load-event-8}
                again, in case it hasn\'t been fired yet).

            7.  Set the
                [`networkState`{#loading-the-media-resource:dom-media-networkstate-10}](#dom-media-networkstate)
                to
                [`NETWORK_LOADING`{#loading-the-media-resource:dom-media-network_loading-4}](#dom-media-network_loading).

        2.  Let `destination`{.variable} be \"`audio`\" if the [media
            element](#media-element){#loading-the-media-resource:media-element-39}
            is an
            [`audio`{#loading-the-media-resource:the-audio-element}](#the-audio-element)
            element, or \"`video`\" otherwise.

        3.  Let `request`{.variable} be the result of [creating a
            potential-CORS
            request](urls-and-fetching.html#create-a-potential-cors-request){#loading-the-media-resource:create-a-potential-cors-request}
            given `current media resource`{.variable}\'s [URL
            record](https://url.spec.whatwg.org/#concept-url){#loading-the-media-resource:url-record-4
            x-internal="url-record"}, `destination`{.variable}, and the
            current state of the [media
            element](#media-element){#loading-the-media-resource:media-element-40}\'s
            [`crossorigin`{#loading-the-media-resource:attr-script-crossorigin}](scripting.html#attr-script-crossorigin)
            content attribute.

        4.  Set `request`{.variable}\'s
            [client](https://fetch.spec.whatwg.org/#concept-request-client){#loading-the-media-resource:concept-request-client
            x-internal="concept-request-client"} to the [media
            element](#media-element){#loading-the-media-resource:media-element-41}\'s
            [node
            document](https://dom.spec.whatwg.org/#concept-node-document){#loading-the-media-resource:node-document-4
            x-internal="node-document"}\'s [relevant settings
            object](webappapis.html#relevant-settings-object){#loading-the-media-resource:relevant-settings-object-2}.

        5.  Set `request`{.variable}\'s [initiator
            type](https://fetch.spec.whatwg.org/#request-initiator-type){#loading-the-media-resource:concept-request-initiator-type
            x-internal="concept-request-initiator-type"} to
            `destination`{.variable}.

        6.  Let `byteRange`{.variable}, which is \"`entire resource`\"
            or a (number, number or \"`until end`\") tuple, be the byte
            range required to satisfy missing data in [media
            data](#media-data){#loading-the-media-resource:media-data}.
            This value is
            [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#loading-the-media-resource:implementation-defined-2
            x-internal="implementation-defined"} and may rely on codec,
            network conditions or other heuristics. The user-agent may
            determine to fetch the resource in full, in which case
            `byteRange`{.variable} would be \"`entire resource`\", to
            fetch from a byte offset until the end, in which case
            `byteRange`{.variable} would be (number, \"`until end`\"),
            or to fetch a range between two byte offsets, in which case
            `byteRange`{.variable} would be a (number, number) tuple
            representing the two offsets.

        7.  If `byteRange`{.variable} is not \"`entire resource`\",
            then:

            1.  If `byteRange`{.variable}\[1\] is \"`until end`\", then
                [add a range
                header](https://fetch.spec.whatwg.org/#concept-request-add-range-header){#loading-the-media-resource:add-a-range-header
                x-internal="add-a-range-header"} to `request`{.variable}
                given `byteRange`{.variable}\[0\].

            2.  Otherwise, [add a range
                header](https://fetch.spec.whatwg.org/#concept-request-add-range-header){#loading-the-media-resource:add-a-range-header-2
                x-internal="add-a-range-header"} to `request`{.variable}
                given `byteRange`{.variable}\[0\] and
                `byteRange`{.variable}\[1\].

        8.  [Fetch](https://fetch.spec.whatwg.org/#concept-fetch){#loading-the-media-resource:concept-fetch
            x-internal="concept-fetch"} `request`{.variable}, with
            *[processResponse](https://fetch.spec.whatwg.org/#process-response){x-internal="processresponse"}*
            set to the following steps given
            [response](https://fetch.spec.whatwg.org/#concept-response){#loading-the-media-resource:concept-response-2
            x-internal="concept-response"} `response`{.variable}:

            1.  Let `global`{.variable} be the [media
                element](#media-element){#loading-the-media-resource:media-element-42}\'s
                [node
                document](https://dom.spec.whatwg.org/#concept-node-document){#loading-the-media-resource:node-document-5
                x-internal="node-document"}\'s [relevant global
                object](webappapis.html#concept-relevant-global){#loading-the-media-resource:concept-relevant-global}.

            2.  Let `updateMedia`{.variable} be to [queue a media
                element
                task](#queue-a-media-element-task){#loading-the-media-resource:queue-a-media-element-task-11}
                given the [media
                element](#media-element){#loading-the-media-resource:media-element-43}
                to run the first appropriate steps from the [media data
                processing steps
                list](#media-data-processing-steps-list){#loading-the-media-resource:media-data-processing-steps-list}
                below. (A new task is used for this so that the work
                described below occurs relative to the appropriate
                [media element event task
                source](#media-element-event-task-source){#loading-the-media-resource:media-element-event-task-source-2}
                rather than using the [networking task
                source](webappapis.html#networking-task-source){#loading-the-media-resource:networking-task-source}.)

            3.  Let `processEndOfMedia`{.variable} be the following
                step: If the fetching process has completed without
                errors, including decoding the media data, and if all of
                the data is available to the user agent without network
                access, then, the user agent must move on to the *final
                step* below. This might never happen, e.g. when
                streaming an infinite resource such as web radio, or if
                the resource is longer than the user agent\'s ability to
                cache data.

            4.  If the result of
                [verifying](#verify-a-media-response){#loading-the-media-resource:verify-a-media-response}
                `response`{.variable} given the
                `current media resource`{.variable} and
                `byteRange`{.variable} is false, then abort these steps.

            5.  Otherwise, [incrementally
                read](https://fetch.spec.whatwg.org/#body-incrementally-read){#loading-the-media-resource:body-incrementally-read
                x-internal="body-incrementally-read"}
                `response`{.variable}\'s
                [body](https://fetch.spec.whatwg.org/#concept-response-body){#loading-the-media-resource:concept-response-body
                x-internal="concept-response-body"} given
                `updateMedia`{.variable},
                `processEndOfMedia`{.variable}, an empty algorithm, and
                `global`{.variable}.

            6.  Update the [media
                data](#media-data){#loading-the-media-resource:media-data-2}
                with the contents of `response`{.variable}\'s [unsafe
                response](urls-and-fetching.html#unsafe-response){#loading-the-media-resource:unsafe-response-2}
                obtained in this fashion. `response`{.variable} can be
                [CORS-same-origin](urls-and-fetching.html#cors-same-origin){#loading-the-media-resource:cors-same-origin}
                or
                [CORS-cross-origin](urls-and-fetching.html#cors-cross-origin){#loading-the-media-resource:cors-cross-origin-2};
                this affects whether subtitles referenced in the [media
                data](#media-data){#loading-the-media-resource:media-data-3}
                are exposed in the API and, for
                [`video`{#loading-the-media-resource:the-video-element}](#the-video-element)
                elements, whether a
                [`canvas`{#loading-the-media-resource:the-canvas-element}](canvas.html#the-canvas-element)
                gets tainted when the video is drawn on it.

            The [media element stall timeout]{#stall-timeout .dfn
            export=""} is an
            [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#loading-the-media-resource:implementation-defined-3
            x-internal="implementation-defined"} length of time, which
            should be about three seconds. When a [media
            element](#media-element){#loading-the-media-resource:media-element-44}
            that is actively attempting to obtain [media
            data](#media-data){#loading-the-media-resource:media-data-4}
            has failed to receive any data for a duration equal to the
            [media element stall
            timeout](#stall-timeout){#loading-the-media-resource:stall-timeout},
            the user agent must [queue a media element
            task](#queue-a-media-element-task){#loading-the-media-resource:queue-a-media-element-task-12}
            given the [media
            element](#media-element){#loading-the-media-resource:media-element-45}
            to:

            1.  [Fire an
                event](https://dom.spec.whatwg.org/#concept-event-fire){#loading-the-media-resource:concept-event-fire-8
                x-internal="concept-event-fire"} named
                [`stalled`{#loading-the-media-resource:event-media-stalled}](#event-media-stalled)
                at the element.

            2.  Set the element\'s [is currently
                stalled](#is-currently-stalled){#loading-the-media-resource:is-currently-stalled-2}
                to true.

            User agents may allow users to selectively block or slow
            [media
            data](#media-data){#loading-the-media-resource:media-data-5}
            downloads. When a [media
            element](#media-element){#loading-the-media-resource:media-element-46}\'s
            download has been blocked altogether, the user agent must
            act as if it was stalled (as opposed to acting as if the
            connection was closed). The rate of the download may also be
            throttled automatically by the user agent, e.g. to balance
            the download with other connections sharing the same
            bandwidth.

            User agents may decide to not download more content at any
            time, e.g. after buffering five minutes of a one hour media
            resource, while waiting for the user to decide whether to
            play the resource or not, while waiting for user input in an
            interactive resource, or when the user navigates away from
            the page. When a [media
            element](#media-element){#loading-the-media-resource:media-element-47}\'s
            download has been suspended, the user agent must [queue a
            media element
            task](#queue-a-media-element-task){#loading-the-media-resource:queue-a-media-element-task-13}
            given the [media
            element](#media-element){#loading-the-media-resource:media-element-48}
            to set the
            [`networkState`{#loading-the-media-resource:dom-media-networkstate-11}](#dom-media-networkstate)
            to
            [`NETWORK_IDLE`{#loading-the-media-resource:dom-media-network_idle-3}](#dom-media-network_idle)
            and [fire an
            event](https://dom.spec.whatwg.org/#concept-event-fire){#loading-the-media-resource:concept-event-fire-9
            x-internal="concept-event-fire"} named
            [`suspend`{#loading-the-media-resource:event-media-suspend-2}](#event-media-suspend)
            at the element. If and when downloading of the resource
            resumes, the user agent must [queue a media element
            task](#queue-a-media-element-task){#loading-the-media-resource:queue-a-media-element-task-14}
            given the [media
            element](#media-element){#loading-the-media-resource:media-element-49}
            to set the
            [`networkState`{#loading-the-media-resource:dom-media-networkstate-12}](#dom-media-networkstate)
            to
            [`NETWORK_LOADING`{#loading-the-media-resource:dom-media-network_loading-5}](#dom-media-network_loading).
            Between the queuing of these tasks, the load is suspended
            (so
            [`progress`{#loading-the-media-resource:event-media-progress}](#event-media-progress)
            events don\'t fire, as described above).

            The
            [`preload`{#loading-the-media-resource:attr-media-preload-2}](#attr-media-preload)
            attribute provides a hint regarding how much buffering the
            author thinks is advisable, even in the absence of the
            [`autoplay`{#loading-the-media-resource:attr-media-autoplay}](#attr-media-autoplay)
            attribute.

            When a user agent decides to completely suspend a download,
            e.g., if it is waiting until the user starts playback before
            downloading any further content, the user agent must [queue
            a media element
            task](#queue-a-media-element-task){#loading-the-media-resource:queue-a-media-element-task-15}
            given the [media
            element](#media-element){#loading-the-media-resource:media-element-50}
            to set the element\'s [delaying-the-load-event
            flag](#delaying-the-load-event-flag){#loading-the-media-resource:delaying-the-load-event-flag-9}
            to false. This stops [delaying the load
            event](parsing.html#delay-the-load-event){#loading-the-media-resource:delay-the-load-event-9}.

            Although the above steps give an algorithm for issuing
            requests, the user agent may use other means besides those
            exact ones, especially in the face of error conditions. For
            example, the user agent may reconnect to the server or
            switch to a streaming protocol. The user agent must only
            consider the resource erroneous, and proceed into the error
            branches of the above steps, if the user agent has given up
            trying to fetch the resource.

            To determine the format of the [media
            resource](#media-resource){#loading-the-media-resource:media-resource-5},
            the user agent must use the [rules for sniffing audio and
            video
            specifically](https://mimesniff.spec.whatwg.org/#rules-for-sniffing-audio-and-video-specifically){#loading-the-media-resource:content-type-sniffing:-video
            x-internal="content-type-sniffing:-video"}.

            While the load is not suspended (see below), every 350ms
            (±200ms) or for every byte received, whichever is *least*
            frequent, [queue a media element
            task](#queue-a-media-element-task){#loading-the-media-resource:queue-a-media-element-task-16}
            given the [media
            element](#media-element){#loading-the-media-resource:media-element-51}
            to:

            1.  [Fire an
                event](https://dom.spec.whatwg.org/#concept-event-fire){#loading-the-media-resource:concept-event-fire-10
                x-internal="concept-event-fire"} named
                [`progress`{#loading-the-media-resource:event-media-progress-2}](#event-media-progress)
                at the element.

            2.  Set the element\'s [is currently
                stalled](#is-currently-stalled){#loading-the-media-resource:is-currently-stalled-3}
                to false.

            While the user agent might still need network access to
            obtain parts of the [media
            resource](#media-resource){#loading-the-media-resource:media-resource-6},
            the user agent must remain on this step.

            For example, if the user agent has discarded the first half
            of a video, the user agent will remain at this step even
            once the [playback has
            ended](#ended-playback){#loading-the-media-resource:ended-playback},
            because there is always the chance the user will seek back
            to the start. In fact, in this situation, once [playback has
            ended](#ended-playback){#loading-the-media-resource:ended-playback-2},
            the user agent will end up firing a
            [`suspend`{#loading-the-media-resource:event-media-suspend-3}](#event-media-suspend)
            event, as described earlier.

    Otherwise (`mode`{.variable} is *local*)

    :   The resource described by the
        `current media resource`{.variable}, if any, contains the [media
        data](#media-data){#loading-the-media-resource:media-data-6}. It
        is
        [CORS-same-origin](urls-and-fetching.html#cors-same-origin){#loading-the-media-resource:cors-same-origin-2}.

        If the `current media resource`{.variable} is a raw data stream
        (e.g. from a
        [`File`{#loading-the-media-resource:file}](https://w3c.github.io/FileAPI/#dfn-file){x-internal="file"}
        object), then to determine the format of the [media
        resource](#media-resource){#loading-the-media-resource:media-resource-7},
        the user agent must use the [rules for sniffing audio and video
        specifically](https://mimesniff.spec.whatwg.org/#rules-for-sniffing-audio-and-video-specifically){#loading-the-media-resource:content-type-sniffing:-video-2
        x-internal="content-type-sniffing:-video"}. Otherwise, if the
        data stream is pre-decoded, then the format is the format given
        by the relevant specification.

        Whenever new data for the `current media resource`{.variable}
        becomes available, [queue a media element
        task](#queue-a-media-element-task){#loading-the-media-resource:queue-a-media-element-task-17}
        given the [media
        element](#media-element){#loading-the-media-resource:media-element-52}
        to run the first appropriate steps from the [media data
        processing steps
        list](#media-data-processing-steps-list){#loading-the-media-resource:media-data-processing-steps-list-2}
        below.

        When the `current media resource`{.variable} is permanently
        exhausted (e.g. all the bytes of a
        [`Blob`{#loading-the-media-resource:blob}](https://w3c.github.io/FileAPI/#dfn-Blob){x-internal="blob"}
        have been processed), if there were no decoding errors, then the
        user agent must move on to the *final step* below. This might
        never happen, e.g. if the `current media resource`{.variable} is
        a
        [`MediaStream`{#loading-the-media-resource:mediastream}](https://w3c.github.io/mediacapture-main/getusermedia.html#idl-def-mediastream){x-internal="mediastream"}.

    The [media data processing steps
    list]{#media-data-processing-steps-list .dfn} is as follows:

    If the [media data](#media-data){#loading-the-media-resource:media-data-7} cannot be fetched at all, due to network errors, causing the user agent to give up trying to fetch the resource\
    If the [media data](#media-data){#loading-the-media-resource:media-data-8} can be fetched but is found by inspection to be in an unsupported format, or can otherwise not be rendered at all

    :   DNS errors, HTTP 4xx and 5xx errors (and equivalents in other
        protocols), and other fatal network errors that occur before the
        user agent has established whether the
        `current media resource`{.variable} is usable, as well as the
        file using an unsupported container format, or using unsupported
        codecs for all the data, must cause the user agent to execute
        the following steps:

        1.  The user agent should cancel the fetching process.

        2.  Abort this subalgorithm, returning to the [resource
            selection
            algorithm](#concept-media-load-algorithm){#loading-the-media-resource:concept-media-load-algorithm-3}.

    If the [media resource](#media-resource){#loading-the-media-resource:media-resource-8} is found to have an audio track

    :   1.  Create an
            [`AudioTrack`{#loading-the-media-resource:audiotrack}](#audiotrack)
            object to represent the audio track.

        2.  Update the [media
            element](#media-element){#loading-the-media-resource:media-element-53}\'s
            [`audioTracks`{#loading-the-media-resource:dom-media-audiotracks}](#dom-media-audiotracks)
            attribute\'s
            [`AudioTrackList`{#loading-the-media-resource:audiotracklist}](#audiotracklist)
            object with the new
            [`AudioTrack`{#loading-the-media-resource:audiotrack-2}](#audiotrack)
            object.

        3.  Let `enable`{.variable} be *unknown*.

        4.  If either the [media
            resource](#media-resource){#loading-the-media-resource:media-resource-9}
            or the
            [URL](https://url.spec.whatwg.org/#concept-url){#loading-the-media-resource:url
            x-internal="url"} of the `current media resource`{.variable}
            indicate a particular set of audio tracks to enable, or if
            the user agent has information that would facilitate the
            selection of specific audio tracks to improve the user\'s
            experience, then: if this audio track is one of the ones to
            enable, then set `enable`{.variable} to *true*, otherwise,
            set `enable`{.variable} to *false*.

            This could be triggered by [media fragment
            syntax](https://www.w3.org/TR/media-frags/#media-fragment-syntax){#loading-the-media-resource:media-fragment-syntax
            x-internal="media-fragment-syntax"}, but it could also be
            triggered e.g. by the user agent selecting a 5.1 surround
            sound audio track over a stereo audio track.

        5.  If `enable`{.variable} is still *unknown*, then, if the
            [media
            element](#media-element){#loading-the-media-resource:media-element-54}
            does not yet have an
            [enabled](#dom-audiotrack-enabled){#loading-the-media-resource:dom-audiotrack-enabled}
            audio track, then set `enable`{.variable} to *true*,
            otherwise, set `enable`{.variable} to *false*.

        6.  If `enable`{.variable} is *true*, then enable this audio
            track, otherwise, do not enable this audio track.

        7.  [Fire an
            event](https://dom.spec.whatwg.org/#concept-event-fire){#loading-the-media-resource:concept-event-fire-11
            x-internal="concept-event-fire"} named
            [`addtrack`{#loading-the-media-resource:event-media-addtrack}](#event-media-addtrack)
            at this
            [`AudioTrackList`{#loading-the-media-resource:audiotracklist-2}](#audiotracklist)
            object, using
            [`TrackEvent`{#loading-the-media-resource:trackevent}](#trackevent),
            with the
            [`track`{#loading-the-media-resource:dom-trackevent-track}](#dom-trackevent-track)
            attribute initialized to the new
            [`AudioTrack`{#loading-the-media-resource:audiotrack-3}](#audiotrack)
            object.

    If the [media resource](#media-resource){#loading-the-media-resource:media-resource-10} is found to have a video track

    :   1.  Create a
            [`VideoTrack`{#loading-the-media-resource:videotrack}](#videotrack)
            object to represent the video track.

        2.  Update the [media
            element](#media-element){#loading-the-media-resource:media-element-55}\'s
            [`videoTracks`{#loading-the-media-resource:dom-media-videotracks}](#dom-media-videotracks)
            attribute\'s
            [`VideoTrackList`{#loading-the-media-resource:videotracklist}](#videotracklist)
            object with the new
            [`VideoTrack`{#loading-the-media-resource:videotrack-2}](#videotrack)
            object.

        3.  Let `enable`{.variable} be *unknown*.

        4.  If either the [media
            resource](#media-resource){#loading-the-media-resource:media-resource-11}
            or the
            [URL](https://url.spec.whatwg.org/#concept-url){#loading-the-media-resource:url-2
            x-internal="url"} of the `current media resource`{.variable}
            indicate a particular set of video tracks to enable, or if
            the user agent has information that would facilitate the
            selection of specific video tracks to improve the user\'s
            experience, then: if this video track is the first such
            video track, then set `enable`{.variable} to *true*,
            otherwise, set `enable`{.variable} to *false*.

            This could again be triggered by [media fragment
            syntax](https://www.w3.org/TR/media-frags/#media-fragment-syntax){#loading-the-media-resource:media-fragment-syntax-2
            x-internal="media-fragment-syntax"}.

        5.  If `enable`{.variable} is still *unknown*, then, if the
            [media
            element](#media-element){#loading-the-media-resource:media-element-56}
            does not yet have a
            [selected](#dom-videotrack-selected){#loading-the-media-resource:dom-videotrack-selected}
            video track, then set `enable`{.variable} to *true*,
            otherwise, set `enable`{.variable} to *false*.

        6.  If `enable`{.variable} is *true*, then select this track and
            unselect any previously selected video tracks, otherwise, do
            not select this video track. If other tracks are unselected,
            then [a `change` event will be fired](#toggle-video-track).

        7.  [Fire an
            event](https://dom.spec.whatwg.org/#concept-event-fire){#loading-the-media-resource:concept-event-fire-12
            x-internal="concept-event-fire"} named
            [`addtrack`{#loading-the-media-resource:event-media-addtrack-2}](#event-media-addtrack)
            at this
            [`VideoTrackList`{#loading-the-media-resource:videotracklist-2}](#videotracklist)
            object, using
            [`TrackEvent`{#loading-the-media-resource:trackevent-2}](#trackevent),
            with the
            [`track`{#loading-the-media-resource:dom-trackevent-track-2}](#dom-trackevent-track)
            attribute initialized to the new
            [`VideoTrack`{#loading-the-media-resource:videotrack-3}](#videotrack)
            object.

    Once enough of the [media data](#media-data){#loading-the-media-resource:media-data-9} has been fetched to determine the duration of the [media resource](#media-resource){#loading-the-media-resource:media-resource-12}, its dimensions, and other metadata

    :   This indicates that the resource is usable. The user agent must
        follow these substeps:

        1.  [Establish the media
            timeline](#defineTimeline){#loading-the-media-resource:defineTimeline}
            for the purposes of the [current playback
            position](#current-playback-position){#loading-the-media-resource:current-playback-position-2}
            and the [earliest possible
            position](#earliest-possible-position){#loading-the-media-resource:earliest-possible-position},
            based on the [media
            data](#media-data){#loading-the-media-resource:media-data-10}.

        2.  Update the [timeline
            offset](#timeline-offset){#loading-the-media-resource:timeline-offset-2}
            to the date and time that corresponds to the zero time in
            the [media
            timeline](#media-timeline){#loading-the-media-resource:media-timeline}
            established in the previous step, if any. If no explicit
            time and date is given by the [media
            resource](#media-resource){#loading-the-media-resource:media-resource-13},
            the [timeline
            offset](#timeline-offset){#loading-the-media-resource:timeline-offset-3}
            must be set to Not-a-Number (NaN).

        3.  Set the [current playback
            position](#current-playback-position){#loading-the-media-resource:current-playback-position-3}
            and the [official playback
            position](#official-playback-position){#loading-the-media-resource:official-playback-position-3}
            to the [earliest possible
            position](#earliest-possible-position){#loading-the-media-resource:earliest-possible-position-2}.

        4.  Update the
            [`duration`{#loading-the-media-resource:dom-media-duration-2}](#dom-media-duration)
            attribute with the time of the last frame of the resource,
            if known, on the [media
            timeline](#media-timeline){#loading-the-media-resource:media-timeline-2}
            established above. If it is not known (e.g. a stream that is
            in principle infinite), update the
            [`duration`{#loading-the-media-resource:dom-media-duration-3}](#dom-media-duration)
            attribute to the value positive Infinity.

            The user agent [will](#durationChange) [queue a media
            element
            task](#queue-a-media-element-task){#loading-the-media-resource:queue-a-media-element-task-18}
            given the [media
            element](#media-element){#loading-the-media-resource:media-element-57}
            to [fire an
            event](https://dom.spec.whatwg.org/#concept-event-fire){#loading-the-media-resource:concept-event-fire-13
            x-internal="concept-event-fire"} named
            [`durationchange`{#loading-the-media-resource:event-media-durationchange-2}](#event-media-durationchange)
            at the element at this point.

        5.  For
            [`video`{#loading-the-media-resource:the-video-element-2}](#the-video-element)
            elements, set the
            [`videoWidth`{#loading-the-media-resource:dom-video-videowidth}](#dom-video-videowidth)
            and
            [`videoHeight`{#loading-the-media-resource:dom-video-videoheight}](#dom-video-videoheight)
            attributes, and [queue a media element
            task](#queue-a-media-element-task){#loading-the-media-resource:queue-a-media-element-task-19}
            given the [media
            element](#media-element){#loading-the-media-resource:media-element-58}
            to [fire an
            event](https://dom.spec.whatwg.org/#concept-event-fire){#loading-the-media-resource:concept-event-fire-14
            x-internal="concept-event-fire"} named
            [`resize`{#loading-the-media-resource:event-media-resize}](#event-media-resize)
            at the [media
            element](#media-element){#loading-the-media-resource:media-element-59}.

            Further
            [`resize`{#loading-the-media-resource:event-media-resize-2}](#event-media-resize)
            events will be fired if the dimensions subsequently change.

        6.  Set the
            [`readyState`{#loading-the-media-resource:dom-media-readystate-2}](#dom-media-readystate)
            attribute to
            [`HAVE_METADATA`{#loading-the-media-resource:dom-media-have_metadata}](#dom-media-have_metadata).

            A
            [`loadedmetadata`{#loading-the-media-resource:event-media-loadedmetadata}](#event-media-loadedmetadata)
            DOM event [will be fired](#fire-loadedmetadata) as part of
            setting the
            [`readyState`{#loading-the-media-resource:dom-media-readystate-3}](#dom-media-readystate)
            attribute to a new value.

        7.  Let `jumped`{.variable} be false.

        8.  If the [media
            element](#media-element){#loading-the-media-resource:media-element-60}\'s
            [default playback start
            position](#default-playback-start-position){#loading-the-media-resource:default-playback-start-position}
            is greater than zero, then
            [seek](#dom-media-seek){#loading-the-media-resource:dom-media-seek}
            to that time, and let `jumped`{.variable} be true.

        9.  Set the [media
            element](#media-element){#loading-the-media-resource:media-element-61}\'s
            [default playback start
            position](#default-playback-start-position){#loading-the-media-resource:default-playback-start-position-2}
            to zero.

        10. Let the `initial playback position`{.variable} be zero.

        11. If either the [media
            resource](#media-resource){#loading-the-media-resource:media-resource-14}
            or the
            [URL](https://url.spec.whatwg.org/#concept-url){#loading-the-media-resource:url-3
            x-internal="url"} of the `current media resource`{.variable}
            indicate a particular start time, then set the
            `initial playback position`{.variable} to that time and, if
            `jumped`{.variable} is still false,
            [seek](#dom-media-seek){#loading-the-media-resource:dom-media-seek-2}
            to that time.

            For example, with media formats that support [media fragment
            syntax](https://www.w3.org/TR/media-frags/#media-fragment-syntax){#loading-the-media-resource:media-fragment-syntax-3
            x-internal="media-fragment-syntax"}, the
            [fragment](https://url.spec.whatwg.org/#concept-url-fragment){#loading-the-media-resource:concept-url-fragment
            x-internal="concept-url-fragment"} can be used to indicate a
            start position.

        12. If there is no
            [enabled](#dom-audiotrack-enabled){#loading-the-media-resource:dom-audiotrack-enabled-2}
            audio track, then enable an audio track. This [will cause a
            `change` event to be fired](#toggle-audio-track).

        13. If there is no
            [selected](#dom-videotrack-selected){#loading-the-media-resource:dom-videotrack-selected-2}
            video track, then select a video track. This [will cause a
            `change` event to be fired](#toggle-video-track).

        Once the
        [`readyState`{#loading-the-media-resource:dom-media-readystate-4}](#dom-media-readystate)
        attribute reaches
        [`HAVE_CURRENT_DATA`{#loading-the-media-resource:dom-media-have_current_data}](#dom-media-have_current_data),
        [after the `loadeddata` event has been fired](#fire-loadeddata),
        set the element\'s [delaying-the-load-event
        flag](#delaying-the-load-event-flag){#loading-the-media-resource:delaying-the-load-event-flag-10}
        to false. This stops [delaying the load
        event](parsing.html#delay-the-load-event){#loading-the-media-resource:delay-the-load-event-10}.

        A user agent that is attempting to reduce network usage while
        still fetching the metadata for each [media
        resource](#media-resource){#loading-the-media-resource:media-resource-15}
        would also stop buffering at this point, following [the rules
        described previously](#resourceSuspend), which involve the
        [`networkState`{#loading-the-media-resource:dom-media-networkstate-13}](#dom-media-networkstate)
        attribute switching to the
        [`NETWORK_IDLE`{#loading-the-media-resource:dom-media-network_idle-4}](#dom-media-network_idle)
        value and a
        [`suspend`{#loading-the-media-resource:event-media-suspend-4}](#event-media-suspend)
        event firing.

        The user agent is *required* to determine the duration of the
        [media
        resource](#media-resource){#loading-the-media-resource:media-resource-16}
        and go through this step before playing.

    Once the entire [media resource](#media-resource){#loading-the-media-resource:media-resource-17} has been fetched (but potentially before any of it has been decoded)

    :   [Fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#loading-the-media-resource:concept-event-fire-15
        x-internal="concept-event-fire"} named
        [`progress`{#loading-the-media-resource:event-media-progress-3}](#event-media-progress)
        at the [media
        element](#media-element){#loading-the-media-resource:media-element-62}.

        Set the
        [`networkState`{#loading-the-media-resource:dom-media-networkstate-14}](#dom-media-networkstate)
        to
        [`NETWORK_IDLE`{#loading-the-media-resource:dom-media-network_idle-5}](#dom-media-network_idle)
        and [fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#loading-the-media-resource:concept-event-fire-16
        x-internal="concept-event-fire"} named
        [`suspend`{#loading-the-media-resource:event-media-suspend-5}](#event-media-suspend)
        at the [media
        element](#media-element){#loading-the-media-resource:media-element-63}.

        If the user agent ever discards any [media
        data](#media-data){#loading-the-media-resource:media-data-11}
        and then needs to resume the network activity to obtain it
        again, then it must [queue a media element
        task](#queue-a-media-element-task){#loading-the-media-resource:queue-a-media-element-task-20}
        given the [media
        element](#media-element){#loading-the-media-resource:media-element-64}
        to set the
        [`networkState`{#loading-the-media-resource:dom-media-networkstate-15}](#dom-media-networkstate)
        to
        [`NETWORK_LOADING`{#loading-the-media-resource:dom-media-network_loading-6}](#dom-media-network_loading).

        If the user agent can keep the [media
        resource](#media-resource){#loading-the-media-resource:media-resource-18}
        loaded, then the algorithm will continue to its *final step*
        below, which aborts the algorithm.

    If the connection is interrupted after some [media data](#media-data){#loading-the-media-resource:media-data-12} has been received, causing the user agent to give up trying to fetch the resource

    :   Fatal network errors that occur after the user agent has
        established whether the `current media resource`{.variable} is
        usable (i.e. once the [media
        element](#media-element){#loading-the-media-resource:media-element-65}\'s
        [`readyState`{#loading-the-media-resource:dom-media-readystate-5}](#dom-media-readystate)
        attribute is no longer
        [`HAVE_NOTHING`{#loading-the-media-resource:dom-media-have_nothing-2}](#dom-media-have_nothing))
        must cause the user agent to execute the following steps:

        1.  The user agent should cancel the fetching process.

        2.  Set the
            [`error`{#loading-the-media-resource:dom-media-error-3}](#dom-media-error)
            attribute to the result of [creating a
            `MediaError`](#creating-a-mediaerror){#loading-the-media-resource:creating-a-mediaerror-2}
            with
            [`MEDIA_ERR_NETWORK`{#loading-the-media-resource:dom-mediaerror-media_err_network}](#dom-mediaerror-media_err_network).

        3.  Set the element\'s
            [`networkState`{#loading-the-media-resource:dom-media-networkstate-16}](#dom-media-networkstate)
            attribute to the
            [`NETWORK_IDLE`{#loading-the-media-resource:dom-media-network_idle-6}](#dom-media-network_idle)
            value.

        4.  Set the element\'s [delaying-the-load-event
            flag](#delaying-the-load-event-flag){#loading-the-media-resource:delaying-the-load-event-flag-11}
            to false. This stops [delaying the load
            event](parsing.html#delay-the-load-event){#loading-the-media-resource:delay-the-load-event-11}.

        5.  [Fire an
            event](https://dom.spec.whatwg.org/#concept-event-fire){#loading-the-media-resource:concept-event-fire-17
            x-internal="concept-event-fire"} named
            [`error`{#loading-the-media-resource:event-media-error-2}](#event-media-error)
            at the [media
            element](#media-element){#loading-the-media-resource:media-element-66}.

        6.  Abort the overall [resource selection
            algorithm](#concept-media-load-algorithm){#loading-the-media-resource:concept-media-load-algorithm-4}.

    If the [media data](#media-data){#loading-the-media-resource:media-data-13} is corrupted

    :   Fatal errors in decoding the [media
        data](#media-data){#loading-the-media-resource:media-data-14}
        that occur after the user agent has established whether the
        `current media resource`{.variable} is usable (i.e. once the
        [media
        element](#media-element){#loading-the-media-resource:media-element-67}\'s
        [`readyState`{#loading-the-media-resource:dom-media-readystate-6}](#dom-media-readystate)
        attribute is no longer
        [`HAVE_NOTHING`{#loading-the-media-resource:dom-media-have_nothing-3}](#dom-media-have_nothing))
        must cause the user agent to execute the following steps:

        1.  The user agent should cancel the fetching process.

        2.  Set the
            [`error`{#loading-the-media-resource:dom-media-error-4}](#dom-media-error)
            attribute to the result of [creating a
            `MediaError`](#creating-a-mediaerror){#loading-the-media-resource:creating-a-mediaerror-3}
            with
            [`MEDIA_ERR_DECODE`{#loading-the-media-resource:dom-mediaerror-media_err_decode}](#dom-mediaerror-media_err_decode).

        3.  Set the element\'s
            [`networkState`{#loading-the-media-resource:dom-media-networkstate-17}](#dom-media-networkstate)
            attribute to the
            [`NETWORK_IDLE`{#loading-the-media-resource:dom-media-network_idle-7}](#dom-media-network_idle)
            value.

        4.  Set the element\'s [delaying-the-load-event
            flag](#delaying-the-load-event-flag){#loading-the-media-resource:delaying-the-load-event-flag-12}
            to false. This stops [delaying the load
            event](parsing.html#delay-the-load-event){#loading-the-media-resource:delay-the-load-event-12}.

        5.  [Fire an
            event](https://dom.spec.whatwg.org/#concept-event-fire){#loading-the-media-resource:concept-event-fire-18
            x-internal="concept-event-fire"} named
            [`error`{#loading-the-media-resource:event-media-error-3}](#event-media-error)
            at the [media
            element](#media-element){#loading-the-media-resource:media-element-68}.

        6.  Abort the overall [resource selection
            algorithm](#concept-media-load-algorithm){#loading-the-media-resource:concept-media-load-algorithm-5}.

    If the [media data](#media-data){#loading-the-media-resource:media-data-15} fetching process is aborted by the user

    :   The fetching process is aborted by the user, e.g. because the
        user pressed a \"stop\" button, the user agent must execute the
        following steps. These steps are not followed if the
        [`load()`{#loading-the-media-resource:dom-media-load}](#dom-media-load)
        method itself is invoked while these steps are running, as the
        steps above handle that particular kind of abort.

        1.  The user agent should cancel the fetching process.

        2.  Set the
            [`error`{#loading-the-media-resource:dom-media-error-5}](#dom-media-error)
            attribute to the result of [creating a
            `MediaError`](#creating-a-mediaerror){#loading-the-media-resource:creating-a-mediaerror-4}
            with
            [`MEDIA_ERR_ABORTED`{#loading-the-media-resource:dom-mediaerror-media_err_aborted}](#dom-mediaerror-media_err_aborted).

        3.  [Fire an
            event](https://dom.spec.whatwg.org/#concept-event-fire){#loading-the-media-resource:concept-event-fire-19
            x-internal="concept-event-fire"} named
            [`abort`{#loading-the-media-resource:event-media-abort-2}](#event-media-abort)
            at the [media
            element](#media-element){#loading-the-media-resource:media-element-69}.

        4.  If the [media
            element](#media-element){#loading-the-media-resource:media-element-70}\'s
            [`readyState`{#loading-the-media-resource:dom-media-readystate-7}](#dom-media-readystate)
            attribute has a value equal to
            [`HAVE_NOTHING`{#loading-the-media-resource:dom-media-have_nothing-4}](#dom-media-have_nothing),
            set the element\'s
            [`networkState`{#loading-the-media-resource:dom-media-networkstate-18}](#dom-media-networkstate)
            attribute to the
            [`NETWORK_EMPTY`{#loading-the-media-resource:dom-media-network_empty-3}](#dom-media-network_empty)
            value, set the element\'s [show poster
            flag](#show-poster-flag){#loading-the-media-resource:show-poster-flag-4}
            to true, and [fire an
            event](https://dom.spec.whatwg.org/#concept-event-fire){#loading-the-media-resource:concept-event-fire-20
            x-internal="concept-event-fire"} named
            [`emptied`{#loading-the-media-resource:event-media-emptied-2}](#event-media-emptied)
            at the element.

            Otherwise, set the element\'s
            [`networkState`{#loading-the-media-resource:dom-media-networkstate-19}](#dom-media-networkstate)
            attribute to the
            [`NETWORK_IDLE`{#loading-the-media-resource:dom-media-network_idle-8}](#dom-media-network_idle)
            value.

        5.  Set the element\'s [delaying-the-load-event
            flag](#delaying-the-load-event-flag){#loading-the-media-resource:delaying-the-load-event-flag-13}
            to false. This stops [delaying the load
            event](parsing.html#delay-the-load-event){#loading-the-media-resource:delay-the-load-event-13}.

        6.  Abort the overall [resource selection
            algorithm](#concept-media-load-algorithm){#loading-the-media-resource:concept-media-load-algorithm-6}.

    If the [media data](#media-data){#loading-the-media-resource:media-data-16} can be fetched but has non-fatal errors or uses, in part, codecs that are unsupported, preventing the user agent from rendering the content completely correctly but not preventing playback altogether

    :   The server returning data that is partially usable but cannot be
        optimally rendered must cause the user agent to render just the
        bits it can handle, and ignore the rest.

    If the [media resource](#media-resource){#loading-the-media-resource:media-resource-19} is found to declare a [media-resource-specific text track](#media-resource-specific-text-track){#loading-the-media-resource:media-resource-specific-text-track-2} that the user agent supports

    :   If the [media
        data](#media-data){#loading-the-media-resource:media-data-17} is
        [CORS-same-origin](urls-and-fetching.html#cors-same-origin){#loading-the-media-resource:cors-same-origin-3},
        run the [steps to expose a media-resource-specific text
        track](#steps-to-expose-a-media-resource-specific-text-track){#loading-the-media-resource:steps-to-expose-a-media-resource-specific-text-track}
        with the relevant data.

        Cross-origin videos do not expose their subtitles, since that
        would allow attacks such as hostile sites reading subtitles from
        confidential videos on a user\'s intranet.

6.  *Final step*: If the user agent ever reaches this step (which can
    only happen if the entire resource gets loaded and kept available):
    abort the overall [resource selection
    algorithm](#concept-media-load-algorithm){#loading-the-media-resource:concept-media-load-algorithm-7}.

When a [media
element](#media-element){#loading-the-media-resource:media-element-71}
is to [forget the media element\'s media-resource-specific
tracks]{#forget-the-media-element's-media-resource-specific-tracks
.dfn}, the user agent must remove from the [media
element](#media-element){#loading-the-media-resource:media-element-72}\'s
[list of text
tracks](#list-of-text-tracks){#loading-the-media-resource:list-of-text-tracks}
all the [media-resource-specific text
tracks](#media-resource-specific-text-track){#loading-the-media-resource:media-resource-specific-text-track-3},
then empty the [media
element](#media-element){#loading-the-media-resource:media-element-73}\'s
[`audioTracks`{#loading-the-media-resource:dom-media-audiotracks-2}](#dom-media-audiotracks)
attribute\'s
[`AudioTrackList`{#loading-the-media-resource:audiotracklist-3}](#audiotracklist)
object, then empty the [media
element](#media-element){#loading-the-media-resource:media-element-74}\'s
[`videoTracks`{#loading-the-media-resource:dom-media-videotracks-2}](#dom-media-videotracks)
attribute\'s
[`VideoTrackList`{#loading-the-media-resource:videotracklist-3}](#videotracklist)
object. No events (in particular, no
[`removetrack`{#loading-the-media-resource:event-media-removetrack}](#event-media-removetrack)
events) are fired as part of this; the
[`error`{#loading-the-media-resource:event-media-error-4}](#event-media-error)
and
[`emptied`{#loading-the-media-resource:event-media-emptied-3}](#event-media-emptied)
events, fired by the algorithms that invoke this one, can be used
instead.

------------------------------------------------------------------------

The [`preload`]{#attr-media-preload .dfn dfn-for="audio,video"
dfn-type="element-attr"} attribute is an [enumerated
attribute](common-microsyntaxes.html#enumerated-attribute){#loading-the-media-resource:enumerated-attribute}
with the following keywords and states:

Keyword

State

Brief description

[`auto`]{#attr-media-preload-auto .dfn
dfn-for="audio/preload,video/preload" dfn-type="attr-value"}

[Automatic]{#attr-media-preload-auto-state .dfn}

Hints to the user agent that the user agent can put the user\'s needs
first without risk to the server, up to and including optimistically
downloading the entire resource.

(the empty string)

[`none`]{#attr-media-preload-none .dfn
dfn-for="audio/preload,video/preload" dfn-type="attr-value"}

[None]{#attr-media-preload-none-state .dfn}

Hints to the user agent that either the author does not expect the user
to need the media resource, or that the server wants to minimize
unnecessary traffic. This state does not provide a hint regarding how
aggressively to actually download the media resource if buffering starts
anyway (e.g. once the user hits \"play\").

[`metadata`]{#attr-media-preload-metadata .dfn
dfn-for="audio/preload,video/preload" dfn-type="attr-value"}

[Metadata]{#attr-media-preload-metadata-state .dfn}

Hints to the user agent that the author does not expect the user to need
the media resource, but that fetching the resource metadata (dimensions,
track list, duration, etc.), and maybe even the first few frames, is
reasonable. If the user agent precisely fetches no more than the
metadata, then the [media
element](#media-element){#loading-the-media-resource:media-element-75}
will end up with its
[`readyState`{#loading-the-media-resource:dom-media-readystate-8}](#dom-media-readystate)
attribute set to
[`HAVE_METADATA`{#loading-the-media-resource:dom-media-have_metadata-2}](#dom-media-have_metadata);
typically though, some frames will be obtained as well and it will
probably be
[`HAVE_CURRENT_DATA`{#loading-the-media-resource:dom-media-have_current_data-2}](#dom-media-have_current_data)
or
[`HAVE_FUTURE_DATA`{#loading-the-media-resource:dom-media-have_future_data}](#dom-media-have_future_data).
When the media resource is playing, hints to the user agent that
bandwidth is to be considered scarce, e.g. suggesting throttling the
download so that the media data is obtained at the slowest possible rate
that still maintains consistent playback.

The attribute\'s *[missing value
default](common-microsyntaxes.html#missing-value-default)* and *[invalid
value default](common-microsyntaxes.html#invalid-value-default)* are
both
[implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#loading-the-media-resource:implementation-defined-4
x-internal="implementation-defined"}, though the
[Metadata](#attr-media-preload-metadata-state){#loading-the-media-resource:attr-media-preload-metadata-state}
state is suggested as a compromise between reducing server load and
providing an optimal user experience.

The attribute can be changed even once the [media
resource](#media-resource){#loading-the-media-resource:media-resource-20}
is being buffered or played; the descriptions in the table above are to
be interpreted with that in mind.

Authors might switch the attribute from
\"[`none`{#loading-the-media-resource:attr-media-preload-none-2}](#attr-media-preload-none)\"
or
\"[`metadata`{#loading-the-media-resource:attr-media-preload-metadata}](#attr-media-preload-metadata)\"
to
\"[`auto`{#loading-the-media-resource:attr-media-preload-auto}](#attr-media-preload-auto)\"
dynamically once the user begins playback. For example, on a page with
many videos this might be used to indicate that the many videos are not
to be downloaded unless requested, but that once one *is* requested it
is to be downloaded aggressively.

The
[`preload`{#loading-the-media-resource:attr-media-preload-3}](#attr-media-preload)
attribute is intended to provide a hint to the user agent about what the
author thinks will lead to the best user experience. The attribute may
be ignored altogether, for example based on explicit user preferences or
based on the available connectivity.

The [`preload`]{#dom-media-preload .dfn dfn-for="HTMLMediaElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#loading-the-media-resource:reflect}
the content attribute of the same name, [limited to only known
values](common-dom-interfaces.html#limited-to-only-known-values){#loading-the-media-resource:limited-to-only-known-values}.

The
[`autoplay`{#loading-the-media-resource:attr-media-autoplay-2}](#attr-media-autoplay)
attribute can override the
[`preload`{#loading-the-media-resource:attr-media-preload-4}](#attr-media-preload)
attribute (since if the media plays, it naturally has to buffer first,
regardless of the hint given by the
[`preload`{#loading-the-media-resource:attr-media-preload-5}](#attr-media-preload)
attribute). Including both is not an error, however.

------------------------------------------------------------------------

`media`{.variable}`.`[`buffered`](#dom-media-buffered){#dom-media-buffered-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/buffered](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/buffered "The buffered read-only property of HTMLMediaElement objects returns a new static normalized TimeRanges object that represents the ranges of the media resource, if any, that the user agent has buffered at the moment the buffered property is accessed.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns a
[`TimeRanges`{#loading-the-media-resource:timeranges}](#timeranges)
object that represents the ranges of the [media
resource](#media-resource){#loading-the-media-resource:media-resource-21}
that the user agent has buffered.

The [`buffered`]{#dom-media-buffered .dfn dfn-for="HTMLMediaElement"
dfn-type="attribute"} attribute must return a new static [normalized
`TimeRanges`
object](#normalised-timeranges-object){#loading-the-media-resource:normalised-timeranges-object}
that represents the ranges of the [media
resource](#media-resource){#loading-the-media-resource:media-resource-22},
if any, that the user agent has buffered, at the time the attribute is
evaluated. User agents must accurately determine the ranges available,
even for media streams where this can only be determined by tedious
inspection.

Typically this will be a single range anchored at the zero point, but
if, e.g. the user agent uses HTTP range requests in response to seeking,
then there could be multiple ranges.

User agents may discard previously buffered data.

Thus, a time position included within a range of the objects return by
the
[`buffered`{#loading-the-media-resource:dom-media-buffered}](#dom-media-buffered)
attribute at one time can end up being not included in the range(s) of
objects returned by the same attribute at later times.

Returning a new object each time is a bad pattern for attribute getters
and is only enshrined here as it would be costly to change it. It is not
to be copied to new APIs.

##### [4.8.11.6]{.secno} Offsets into the media resource[](#offsets-into-the-media-resource){.self-link}

`media`{.variable}`.`[`duration`](#dom-media-duration){#dom-media-duration-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/duration](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/duration "The read-only HTMLMediaElement property duration indicates the length of the element's media in seconds.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns the length of the [media
resource](#media-resource){#offsets-into-the-media-resource:media-resource},
in seconds, assuming that the start of the [media
resource](#media-resource){#offsets-into-the-media-resource:media-resource-2}
is at time zero.

Returns NaN if the duration isn\'t available.

Returns Infinity for unbounded streams.

`media`{.variable}`.`[`currentTime`](#dom-media-currenttime){#dom-media-currenttime-dev}` [ = ``value`{.variable}` ]`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/currentTime](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/currentTime "The HTMLMediaElement interface's currentTime property specifies the current playback time in seconds.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns the [official playback
position](#official-playback-position){#offsets-into-the-media-resource:official-playback-position},
in seconds.

Can be set, to seek to the given time.

A [media
resource](#media-resource){#offsets-into-the-media-resource:media-resource-3}
has a [media timeline]{#media-timeline .dfn} that maps times (in
seconds) to positions in the [media
resource](#media-resource){#offsets-into-the-media-resource:media-resource-4}.
The origin of a timeline is its earliest defined position. The duration
of a timeline is its last defined position.

[Establishing the media timeline]{#defineTimeline .dfn}: if the [media
resource](#media-resource){#offsets-into-the-media-resource:media-resource-5}
somehow specifies an explicit timeline whose origin is not negative
(i.e. gives each frame a specific time offset and gives the first frame
a zero or positive offset), then the [media
timeline](#media-timeline){#offsets-into-the-media-resource:media-timeline}
should be that timeline. (Whether the [media
resource](#media-resource){#offsets-into-the-media-resource:media-resource-6}
can specify a timeline or not depends on the [media
resource\'s](#media-resource){#offsets-into-the-media-resource:media-resource-7}
format.) If the [media
resource](#media-resource){#offsets-into-the-media-resource:media-resource-8}
specifies an explicit start time *and date*, then that time and date
should be considered the zero point in the [media
timeline](#media-timeline){#offsets-into-the-media-resource:media-timeline-2};
the [timeline
offset](#timeline-offset){#offsets-into-the-media-resource:timeline-offset}
will be the time and date, exposed using the
[`getStartDate()`{#offsets-into-the-media-resource:dom-media-getstartdate}](#dom-media-getstartdate)
method.

If the [media
resource](#media-resource){#offsets-into-the-media-resource:media-resource-9}
has a discontinuous timeline, the user agent must extend the timeline
used at the start of the resource across the entire resource, so that
the [media
timeline](#media-timeline){#offsets-into-the-media-resource:media-timeline-3}
of the [media
resource](#media-resource){#offsets-into-the-media-resource:media-resource-10}
increases linearly starting from the [earliest possible
position](#earliest-possible-position){#offsets-into-the-media-resource:earliest-possible-position}
(as defined below), even if the underlying [media
data](#media-data){#offsets-into-the-media-resource:media-data} has
out-of-order or even overlapping time codes.

For example, if two clips have been concatenated into one video file,
but the video format exposes the original times for the two clips, the
video data might expose a timeline that goes, say, 00:15..00:29 and then
00:05..00:38. However, the user agent would not expose those times; it
would instead expose the times as 00:15..00:29 and 00:29..01:02, as a
single video.

[![(This is a tracking
vector.)](https://resources.whatwg.org/tracking-vector.svg){width="46"
crossorigin=""
height="64"}](https://infra.spec.whatwg.org/#tracking-vector "There is a tracking vector here."){#offsets-into-the-media-resource:tracking-vector
.tracking-vector x-internal="tracking-vector"} In the rare case of a
[media
resource](#media-resource){#offsets-into-the-media-resource:media-resource-11}
that does not have an explicit timeline, the zero time on the [media
timeline](#media-timeline){#offsets-into-the-media-resource:media-timeline-4}
should correspond to the first frame of the [media
resource](#media-resource){#offsets-into-the-media-resource:media-resource-12}.
In the even rarer case of a [media
resource](#media-resource){#offsets-into-the-media-resource:media-resource-13}
with no explicit timings of any kind, not even frame durations, the user
agent must itself determine the time for each frame in an
[implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#offsets-into-the-media-resource:implementation-defined
x-internal="implementation-defined"} manner.

An example of a file format with no explicit timeline but with explicit
frame durations is the Animated GIF format. An example of a file format
with no explicit timings at all is the JPEG-push format
([`multipart/x-mixed-replace`{#offsets-into-the-media-resource:multipart/x-mixed-replace}](iana.html#multipart/x-mixed-replace)
with JPEG frames, often used as the format for MJPEG streams).

If, in the case of a resource with no timing information, the user agent
will nonetheless be able to seek to an earlier point than the first
frame originally provided by the server, then the zero time should
correspond to the earliest seekable time of the [media
resource](#media-resource){#offsets-into-the-media-resource:media-resource-14};
otherwise, it should correspond to the first frame received from the
server (the point in the [media
resource](#media-resource){#offsets-into-the-media-resource:media-resource-15}
at which the user agent began receiving the stream).

At the time of writing, there is no known format that lacks explicit
frame time offsets yet still supports seeking to a frame before the
first frame sent by the server.

::: example
Consider a stream from a TV broadcaster, which begins streaming on a
sunny Friday afternoon in October, and always sends connecting user
agents the media data on the same media timeline, with its zero time set
to the start of this stream. Months later, user agents connecting to
this stream will find that the first frame they receive has a time with
millions of seconds. The
[`getStartDate()`{#offsets-into-the-media-resource:dom-media-getstartdate-2}](#dom-media-getstartdate)
method would always return the date that the broadcast started; this
would allow controllers to display real times in their scrubber (e.g.
\"2:30pm\") rather than a time relative to when the broadcast began (\"8
months, 4 hours, 12 minutes, and 23 seconds\").

Consider a stream that carries a video with several concatenated
fragments, broadcast by a server that does not allow user agents to
request specific times but instead just streams the video data in a
predetermined order, with the first frame delivered always being
identified as the frame with time zero. If a user agent connects to this
stream and receives fragments defined as covering timestamps 2010-03-20
23:15:00 UTC to 2010-03-21 00:05:00 UTC and 2010-02-12 14:25:00 UTC to
2010-02-12 14:35:00 UTC, it would expose this with a [media
timeline](#media-timeline){#offsets-into-the-media-resource:media-timeline-5}
starting at 0s and extending to 3,600s (one hour). Assuming the
streaming server disconnected at the end of the second clip, the
[`duration`{#offsets-into-the-media-resource:dom-media-duration}](#dom-media-duration)
attribute would then return 3,600. The
[`getStartDate()`{#offsets-into-the-media-resource:dom-media-getstartdate-3}](#dom-media-getstartdate)
method would return a
[`Date`{#offsets-into-the-media-resource:date}](https://tc39.es/ecma262/#sec-date-objects){x-internal="date"}
object with a time corresponding to 2010-03-20 23:15:00 UTC. However, if
a different user agent connected five minutes later, *it* would
(presumably) receive fragments covering timestamps 2010-03-20 23:20:00
UTC to 2010-03-21 00:05:00 UTC and 2010-02-12 14:25:00 UTC to 2010-02-12
14:35:00 UTC, and would expose this with a [media
timeline](#media-timeline){#offsets-into-the-media-resource:media-timeline-6}
starting at 0s and extending to 3,300s (fifty five minutes). In this
case, the
[`getStartDate()`{#offsets-into-the-media-resource:dom-media-getstartdate-4}](#dom-media-getstartdate)
method would return a
[`Date`{#offsets-into-the-media-resource:date-2}](https://tc39.es/ecma262/#sec-date-objects){x-internal="date"}
object with a time corresponding to 2010-03-20 23:20:00 UTC.

In both of these examples, the
[`seekable`{#offsets-into-the-media-resource:dom-media-seekable}](#dom-media-seekable)
attribute would give the ranges that the controller would want to
actually display in its UI; typically, if the servers don\'t support
seeking to arbitrary times, this would be the range of time from the
moment the user agent connected to the stream up to the latest frame
that the user agent has obtained; however, if the user agent starts
discarding earlier information, the actual range might be shorter.
:::

In any case, the user agent must ensure that the [earliest possible
position](#earliest-possible-position){#offsets-into-the-media-resource:earliest-possible-position-2}
(as defined below) using the established [media
timeline](#media-timeline){#offsets-into-the-media-resource:media-timeline-7},
is greater than or equal to zero.

The [media
timeline](#media-timeline){#offsets-into-the-media-resource:media-timeline-8}
also has an associated clock. Which clock is used is user-agent defined,
and may be [media
resource](#media-resource){#offsets-into-the-media-resource:media-resource-16}-dependent,
but it should approximate the user\'s wall clock.

[Media
elements](#media-element){#offsets-into-the-media-resource:media-element}
have a [current playback position]{#current-playback-position .dfn},
which must initially (i.e. in the absence of [media
data](#media-data){#offsets-into-the-media-resource:media-data-2}) be
zero seconds. The [current playback
position](#current-playback-position){#offsets-into-the-media-resource:current-playback-position}
is a time on the [media
timeline](#media-timeline){#offsets-into-the-media-resource:media-timeline-9}.

[Media
elements](#media-element){#offsets-into-the-media-resource:media-element-2}
also have an [official playback position]{#official-playback-position
.dfn}, which must initially be set to zero seconds. The [official
playback
position](#official-playback-position){#offsets-into-the-media-resource:official-playback-position-2}
is an approximation of the [current playback
position](#current-playback-position){#offsets-into-the-media-resource:current-playback-position-2}
that is kept stable while scripts are running.

[Media
elements](#media-element){#offsets-into-the-media-resource:media-element-3}
also have a [default playback start
position]{#default-playback-start-position .dfn}, which must initially
be set to zero seconds. This time is used to allow the element to be
seeked even before the media is loaded.

Each [media
element](#media-element){#offsets-into-the-media-resource:media-element-4}
has a [show poster flag]{#show-poster-flag .dfn}. When a [media
element](#media-element){#offsets-into-the-media-resource:media-element-5}
is created, this flag must be set to true. This flag is used to control
when the user agent is to show a poster frame for a
[`video`{#offsets-into-the-media-resource:the-video-element}](#the-video-element)
element instead of showing the video contents.

The [`currentTime`]{#dom-media-currenttime .dfn
dfn-for="HTMLMediaElement" dfn-type="attribute"} attribute must, on
getting, return the [media
element](#media-element){#offsets-into-the-media-resource:media-element-6}\'s
[default playback start
position](#default-playback-start-position){#offsets-into-the-media-resource:default-playback-start-position},
unless that is zero, in which case it must return the element\'s
[official playback
position](#official-playback-position){#offsets-into-the-media-resource:official-playback-position-3}.
The returned value must be expressed in seconds. On setting, if the
[media
element](#media-element){#offsets-into-the-media-resource:media-element-7}\'s
[`readyState`{#offsets-into-the-media-resource:dom-media-readystate}](#dom-media-readystate)
is
[`HAVE_NOTHING`{#offsets-into-the-media-resource:dom-media-have_nothing}](#dom-media-have_nothing),
then it must set the [media
element](#media-element){#offsets-into-the-media-resource:media-element-8}\'s
[default playback start
position](#default-playback-start-position){#offsets-into-the-media-resource:default-playback-start-position-2}
to the new value; otherwise, it must set the [official playback
position](#official-playback-position){#offsets-into-the-media-resource:official-playback-position-4}
to the new value and then
[seek](#dom-media-seek){#offsets-into-the-media-resource:dom-media-seek}
to the new value. The new value must be interpreted as being in seconds.

If the [media
resource](#media-resource){#offsets-into-the-media-resource:media-resource-17}
is a streaming resource, then the user agent might be unable to obtain
certain parts of the resource after it has expired from its buffer.
Similarly, some [media
resources](#media-resource){#offsets-into-the-media-resource:media-resource-18}
might have a [media
timeline](#media-timeline){#offsets-into-the-media-resource:media-timeline-10}
that doesn\'t start at zero. The [earliest possible
position]{#earliest-possible-position .dfn} is the earliest position in
the stream or resource that the user agent can ever obtain again. It is
also a time on the [media
timeline](#media-timeline){#offsets-into-the-media-resource:media-timeline-11}.

The [earliest possible
position](#earliest-possible-position){#offsets-into-the-media-resource:earliest-possible-position-3}
is not explicitly exposed in the API; it corresponds to the start time
of the first range in the
[`seekable`{#offsets-into-the-media-resource:dom-media-seekable-2}](#dom-media-seekable)
attribute\'s
[`TimeRanges`{#offsets-into-the-media-resource:timeranges}](#timeranges)
object, if any, or the [current playback
position](#current-playback-position){#offsets-into-the-media-resource:current-playback-position-3}
otherwise.

When the [earliest possible
position](#earliest-possible-position){#offsets-into-the-media-resource:earliest-possible-position-4}
changes, then: if the [current playback
position](#current-playback-position){#offsets-into-the-media-resource:current-playback-position-4}
is before the [earliest possible
position](#earliest-possible-position){#offsets-into-the-media-resource:earliest-possible-position-5},
the user agent must
[seek](#dom-media-seek){#offsets-into-the-media-resource:dom-media-seek-2}
to the [earliest possible
position](#earliest-possible-position){#offsets-into-the-media-resource:earliest-possible-position-6};
otherwise, if the user agent has not fired a
[`timeupdate`{#offsets-into-the-media-resource:event-media-timeupdate}](#event-media-timeupdate)
event at the element in the past 15 to 250ms and is not still running
event handlers for such an event, then the user agent must [queue a
media element
task](#queue-a-media-element-task){#offsets-into-the-media-resource:queue-a-media-element-task}
given the [media
element](#media-element){#offsets-into-the-media-resource:media-element-9}
to [fire an
event](https://dom.spec.whatwg.org/#concept-event-fire){#offsets-into-the-media-resource:concept-event-fire
x-internal="concept-event-fire"} named
[`timeupdate`{#offsets-into-the-media-resource:event-media-timeupdate-2}](#event-media-timeupdate)
at the element.

Because of the above requirement and the requirement in the [resource
fetch
algorithm](#concept-media-load-resource){#offsets-into-the-media-resource:concept-media-load-resource}
that kicks in [when the metadata of the clip becomes
known](#getting-media-metadata), the [current playback
position](#current-playback-position){#offsets-into-the-media-resource:current-playback-position-5}
can never be less than the [earliest possible
position](#earliest-possible-position){#offsets-into-the-media-resource:earliest-possible-position-7}.

If at any time the user agent learns that an audio or video track has
ended and all [media
data](#media-data){#offsets-into-the-media-resource:media-data-3}
relating to that track corresponds to parts of the [media
timeline](#media-timeline){#offsets-into-the-media-resource:media-timeline-12}
that are *before* the [earliest possible
position](#earliest-possible-position){#offsets-into-the-media-resource:earliest-possible-position-8},
the user agent may [queue a media element
task](#queue-a-media-element-task){#offsets-into-the-media-resource:queue-a-media-element-task-2}
given the [media
element](#media-element){#offsets-into-the-media-resource:media-element-10}
to run these steps:

1.  Remove the track from the
    [`audioTracks`{#offsets-into-the-media-resource:dom-media-audiotracks}](#dom-media-audiotracks)
    attribute\'s
    [`AudioTrackList`{#offsets-into-the-media-resource:audiotracklist}](#audiotracklist)
    object or the
    [`videoTracks`{#offsets-into-the-media-resource:dom-media-videotracks}](#dom-media-videotracks)
    attribute\'s
    [`VideoTrackList`{#offsets-into-the-media-resource:videotracklist}](#videotracklist)
    object as appropriate.

2.  [Fire an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#offsets-into-the-media-resource:concept-event-fire-2
    x-internal="concept-event-fire"} named
    [`removetrack`{#offsets-into-the-media-resource:event-media-removetrack}](#event-media-removetrack)
    at the [media
    element](#media-element){#offsets-into-the-media-resource:media-element-11}\'s
    aforementioned
    [`AudioTrackList`{#offsets-into-the-media-resource:audiotracklist-2}](#audiotracklist)
    or
    [`VideoTrackList`{#offsets-into-the-media-resource:videotracklist-2}](#videotracklist)
    object, using
    [`TrackEvent`{#offsets-into-the-media-resource:trackevent}](#trackevent),
    with the
    [`track`{#offsets-into-the-media-resource:dom-trackevent-track}](#dom-trackevent-track)
    attribute initialized to the
    [`AudioTrack`{#offsets-into-the-media-resource:audiotrack}](#audiotrack)
    or
    [`VideoTrack`{#offsets-into-the-media-resource:videotrack}](#videotrack)
    object representing the track.

The [`duration`]{#dom-media-duration .dfn dfn-for="HTMLMediaElement"
dfn-type="attribute"} attribute must return the time of the end of the
[media
resource](#media-resource){#offsets-into-the-media-resource:media-resource-19},
in seconds, on the [media
timeline](#media-timeline){#offsets-into-the-media-resource:media-timeline-13}.
If no [media
data](#media-data){#offsets-into-the-media-resource:media-data-4} is
available, then the attributes must return the Not-a-Number (NaN) value.
If the [media
resource](#media-resource){#offsets-into-the-media-resource:media-resource-20}
is not known to be bounded (e.g. streaming radio, or a live event with
no announced end time), then the attribute must return the positive
Infinity value.

The user agent must determine the duration of the [media
resource](#media-resource){#offsets-into-the-media-resource:media-resource-21}
before playing any part of the [media
data](#media-data){#offsets-into-the-media-resource:media-data-5} and
before setting
[`readyState`{#offsets-into-the-media-resource:dom-media-readystate-2}](#dom-media-readystate)
to a value greater than or equal to
[`HAVE_METADATA`{#offsets-into-the-media-resource:dom-media-have_metadata}](#dom-media-have_metadata),
even if doing so requires fetching multiple parts of the resource.

When the length of the [media
resource](#media-resource){#offsets-into-the-media-resource:media-resource-22}
changes to a known value (e.g. from being unknown to known, or from a
previously established length to a new length) the user agent must
[queue a media element
task](#queue-a-media-element-task){#offsets-into-the-media-resource:queue-a-media-element-task-3}
given the [media
element](#media-element){#offsets-into-the-media-resource:media-element-12}
to [fire an
event](https://dom.spec.whatwg.org/#concept-event-fire){#offsets-into-the-media-resource:concept-event-fire-3
x-internal="concept-event-fire"} named
[`durationchange`{#offsets-into-the-media-resource:event-media-durationchange}](#event-media-durationchange)
at the [media
element](#media-element){#offsets-into-the-media-resource:media-element-13}.
(The event is not fired when the duration is reset as part of loading a
new media resource.) If the duration is changed such that the [current
playback
position](#current-playback-position){#offsets-into-the-media-resource:current-playback-position-6}
ends up being greater than the time of the end of the [media
resource](#media-resource){#offsets-into-the-media-resource:media-resource-23},
then the user agent must also
[seek](#dom-media-seek){#offsets-into-the-media-resource:dom-media-seek-3}
to the time of the end of the [media
resource](#media-resource){#offsets-into-the-media-resource:media-resource-24}.

If an \"infinite\" stream ends for some reason, then the duration would
change from positive Infinity to the time of the last frame or sample in
the stream, and the
[`durationchange`{#offsets-into-the-media-resource:event-media-durationchange-2}](#event-media-durationchange)
event would be fired. Similarly, if the user agent initially estimated
the [media
resource](#media-resource){#offsets-into-the-media-resource:media-resource-25}\'s
duration instead of determining it precisely, and later revises the
estimate based on new information, then the duration would change and
the
[`durationchange`{#offsets-into-the-media-resource:event-media-durationchange-3}](#event-media-durationchange)
event would be fired.

Some video files also have an explicit date and time corresponding to
the zero time in the [media
timeline](#media-timeline){#offsets-into-the-media-resource:media-timeline-14},
known as the [timeline offset]{#timeline-offset .dfn}. Initially, the
[timeline
offset](#timeline-offset){#offsets-into-the-media-resource:timeline-offset-2}
must be set to Not-a-Number (NaN).

The [`getStartDate()`]{#dom-media-getstartdate .dfn
dfn-for="HTMLMediaElement" dfn-type="method"} method must return [a new
`Date`
object](infrastructure.html#create-a-date-object){#offsets-into-the-media-resource:create-a-date-object}
representing the current [timeline
offset](#timeline-offset){#offsets-into-the-media-resource:timeline-offset-3}.

------------------------------------------------------------------------

The [`loop`]{#attr-media-loop .dfn dfn-for="audio,video"
dfn-type="element-attr"} attribute is a [boolean
attribute](common-microsyntaxes.html#boolean-attribute){#offsets-into-the-media-resource:boolean-attribute}
that, if specified, indicates that the [media
element](#media-element){#offsets-into-the-media-resource:media-element-14}
is to seek back to the start of the [media
resource](#media-resource){#offsets-into-the-media-resource:media-resource-26}
upon reaching the end.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/loop](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/loop "The HTMLMediaElement.loop property reflects the loop HTML attribute, which controls whether the media element should start over when it reaches the end.")

Support in all current engines.

::: support
[Firefox11+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The [`loop`]{#dom-media-loop .dfn dfn-for="HTMLMediaElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#offsets-into-the-media-resource:reflect}
the content attribute of the same name.

##### [4.8.11.7]{.secno} Ready states[](#ready-states){.self-link}

`media`{.variable}`.`[`readyState`](#dom-media-readystate){#dom-media-readystate-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/readyState](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/readyState "The HTMLMediaElement.readyState property indicates the readiness state of the media.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns a value that expresses the current state of the element with
respect to rendering the [current playback
position](#current-playback-position){#ready-states:current-playback-position},
from the codes in the list below.

[Media elements](#media-element){#ready-states:media-element} have a
*ready state*, which describes to what degree they are ready to be
rendered at the [current playback
position](#current-playback-position){#ready-states:current-playback-position-2}.
The possible values are as follows; the ready state of a media element
at any particular time is the greatest value describing the state of the
element:

[`HAVE_NOTHING`]{#dom-media-have_nothing .dfn dfn-for="HTMLMediaElement" dfn-type="const"} (numeric value 0)
:   No information regarding the [media
    resource](#media-resource){#ready-states:media-resource} is
    available. No data for the [current playback
    position](#current-playback-position){#ready-states:current-playback-position-3}
    is available. [Media
    elements](#media-element){#ready-states:media-element-2} whose
    [`networkState`{#ready-states:dom-media-networkstate}](#dom-media-networkstate)
    attribute are set to
    [`NETWORK_EMPTY`{#ready-states:dom-media-network_empty}](#dom-media-network_empty)
    are always in the
    [`HAVE_NOTHING`{#ready-states:dom-media-have_nothing}](#dom-media-have_nothing)
    state.

[`HAVE_METADATA`]{#dom-media-have_metadata .dfn dfn-for="HTMLMediaElement" dfn-type="const"} (numeric value 1)
:   Enough of the resource has been obtained that the duration of the
    resource is available. In the case of a
    [`video`{#ready-states:the-video-element}](#the-video-element)
    element, the dimensions of the video are also available. No [media
    data](#media-data){#ready-states:media-data} is available for the
    immediate [current playback
    position](#current-playback-position){#ready-states:current-playback-position-4}.

[`HAVE_CURRENT_DATA`]{#dom-media-have_current_data .dfn dfn-for="HTMLMediaElement" dfn-type="const"} (numeric value 2)
:   Data for the immediate [current playback
    position](#current-playback-position){#ready-states:current-playback-position-5}
    is available, but either not enough data is available that the user
    agent could successfully advance the [current playback
    position](#current-playback-position){#ready-states:current-playback-position-6}
    in the [direction of
    playback](#direction-of-playback){#ready-states:direction-of-playback}
    at all without immediately reverting to the
    [`HAVE_METADATA`{#ready-states:dom-media-have_metadata}](#dom-media-have_metadata)
    state, or there is no more data to obtain in the [direction of
    playback](#direction-of-playback){#ready-states:direction-of-playback-2}.
    For example, in video this corresponds to the user agent having data
    from the current frame, but not the next frame, when the [current
    playback
    position](#current-playback-position){#ready-states:current-playback-position-7}
    is at the end of the current frame; and to when [playback has
    ended](#ended-playback){#ready-states:ended-playback}.

[`HAVE_FUTURE_DATA`]{#dom-media-have_future_data .dfn dfn-for="HTMLMediaElement" dfn-type="const"} (numeric value 3)
:   Data for the immediate [current playback
    position](#current-playback-position){#ready-states:current-playback-position-8}
    is available, as well as enough data for the user agent to advance
    the [current playback
    position](#current-playback-position){#ready-states:current-playback-position-9}
    in the [direction of
    playback](#direction-of-playback){#ready-states:direction-of-playback-3}
    at least a little without immediately reverting to the
    [`HAVE_METADATA`{#ready-states:dom-media-have_metadata-2}](#dom-media-have_metadata)
    state, and [the text tracks are
    ready](#the-text-tracks-are-ready){#ready-states:the-text-tracks-are-ready}.
    For example, in video this corresponds to the user agent having data
    for at least the current frame and the next frame when the [current
    playback
    position](#current-playback-position){#ready-states:current-playback-position-10}
    is at the instant in time between the two frames, or to the user
    agent having the video data for the current frame and audio data to
    keep playing at least a little when the [current playback
    position](#current-playback-position){#ready-states:current-playback-position-11}
    is in the middle of a frame. The user agent cannot be in this state
    if [playback has
    ended](#ended-playback){#ready-states:ended-playback-2}, as the
    [current playback
    position](#current-playback-position){#ready-states:current-playback-position-12}
    can never advance in this case.

[`HAVE_ENOUGH_DATA`]{#dom-media-have_enough_data .dfn dfn-for="HTMLMediaElement" dfn-type="const"} (numeric value 4)

:   All the conditions described for the
    [`HAVE_FUTURE_DATA`{#ready-states:dom-media-have_future_data}](#dom-media-have_future_data)
    state are met, and, in addition, either of the following conditions
    is also true:

    - The user agent estimates that data is being fetched at a rate
      where the [current playback
      position](#current-playback-position){#ready-states:current-playback-position-13},
      if it were to advance at the element\'s
      [`playbackRate`{#ready-states:dom-media-playbackrate}](#dom-media-playbackrate),
      would not overtake the available data before playback reaches the
      end of the [media
      resource](#media-resource){#ready-states:media-resource-2}.
    - The user agent has entered a state where waiting longer will not
      result in further data being obtained, and therefore nothing would
      be gained by delaying playback any further. (For example, the
      buffer might be full.)

In practice, the difference between
[`HAVE_METADATA`{#ready-states:dom-media-have_metadata-3}](#dom-media-have_metadata)
and
[`HAVE_CURRENT_DATA`{#ready-states:dom-media-have_current_data}](#dom-media-have_current_data)
is negligible. Really the only time the difference is relevant is when
painting a
[`video`{#ready-states:the-video-element-2}](#the-video-element) element
onto a
[`canvas`{#ready-states:the-canvas-element}](canvas.html#the-canvas-element),
where it distinguishes the case where something will be drawn
([`HAVE_CURRENT_DATA`{#ready-states:dom-media-have_current_data-2}](#dom-media-have_current_data)
or greater) from the case where nothing is drawn
([`HAVE_METADATA`{#ready-states:dom-media-have_metadata-4}](#dom-media-have_metadata)
or less). Similarly, the difference between
[`HAVE_CURRENT_DATA`{#ready-states:dom-media-have_current_data-3}](#dom-media-have_current_data)
(only the current frame) and
[`HAVE_FUTURE_DATA`{#ready-states:dom-media-have_future_data-2}](#dom-media-have_future_data)
(at least this frame and the next) can be negligible (in the extreme,
only one frame). The only time that distinction really matters is when a
page provides an interface for \"frame-by-frame\" navigation.

When the ready state of a [media
element](#media-element){#ready-states:media-element-3} whose
[`networkState`{#ready-states:dom-media-networkstate-2}](#dom-media-networkstate)
is not
[`NETWORK_EMPTY`{#ready-states:dom-media-network_empty-2}](#dom-media-network_empty)
changes, the user agent must follow the steps given below:

1.  Apply the first applicable set of substeps from the following list:

    If the previous ready state was [`HAVE_NOTHING`{#ready-states:dom-media-have_nothing-2}](#dom-media-have_nothing), and the new ready state is [`HAVE_METADATA`{#ready-states:dom-media-have_metadata-5}](#dom-media-have_metadata)

    :   [Queue a media element
        task](#queue-a-media-element-task){#ready-states:queue-a-media-element-task}
        given the [media
        element](#media-element){#ready-states:media-element-4} to [fire
        an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#ready-states:concept-event-fire
        x-internal="concept-event-fire"} named
        [`loadedmetadata`{#ready-states:event-media-loadedmetadata}](#event-media-loadedmetadata)
        at the element.

        Before this task is run, as part of the [event
        loop](webappapis.html#event-loop){#ready-states:event-loop}
        mechanism, the rendering will have been updated to resize the
        [`video`{#ready-states:the-video-element-3}](#the-video-element)
        element if appropriate.

    If the previous ready state was [`HAVE_METADATA`{#ready-states:dom-media-have_metadata-6}](#dom-media-have_metadata) and the new ready state is [`HAVE_CURRENT_DATA`{#ready-states:dom-media-have_current_data-4}](#dom-media-have_current_data) or greater

    :   If this is the first time this occurs for this [media
        element](#media-element){#ready-states:media-element-5} since
        the [`load()`{#ready-states:dom-media-load}](#dom-media-load)
        algorithm was last invoked, the user agent must [queue a media
        element
        task](#queue-a-media-element-task){#ready-states:queue-a-media-element-task-2}
        given the [media
        element](#media-element){#ready-states:media-element-6} to [fire
        an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#ready-states:concept-event-fire-2
        x-internal="concept-event-fire"} named
        [`loadeddata`{#ready-states:event-media-loadeddata}](#event-media-loadeddata)
        at the element.

        If the new ready state is
        [`HAVE_FUTURE_DATA`{#ready-states:dom-media-have_future_data-3}](#dom-media-have_future_data)
        or
        [`HAVE_ENOUGH_DATA`{#ready-states:dom-media-have_enough_data}](#dom-media-have_enough_data),
        then the relevant steps below must then be run also.

    If the previous ready state was [`HAVE_FUTURE_DATA`{#ready-states:dom-media-have_future_data-4}](#dom-media-have_future_data) or more, and the new ready state is [`HAVE_CURRENT_DATA`{#ready-states:dom-media-have_current_data-5}](#dom-media-have_current_data) or less

    :   If the [media
        element](#media-element){#ready-states:media-element-7} was
        [potentially
        playing](#potentially-playing){#ready-states:potentially-playing}
        before its
        [`readyState`{#ready-states:dom-media-readystate}](#dom-media-readystate)
        attribute changed to a value lower than
        [`HAVE_FUTURE_DATA`{#ready-states:dom-media-have_future_data-5}](#dom-media-have_future_data),
        and the element has not [ended
        playback](#ended-playback){#ready-states:ended-playback-3}, and
        playback has not [stopped due to
        errors](#stopped-due-to-errors){#ready-states:stopped-due-to-errors},
        [paused for user
        interaction](#paused-for-user-interaction){#ready-states:paused-for-user-interaction},
        or [paused for in-band
        content](#paused-for-in-band-content){#ready-states:paused-for-in-band-content},
        the user agent must [queue a media element
        task](#queue-a-media-element-task){#ready-states:queue-a-media-element-task-3}
        given the [media
        element](#media-element){#ready-states:media-element-8} to [fire
        an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#ready-states:concept-event-fire-3
        x-internal="concept-event-fire"} named
        [`timeupdate`{#ready-states:event-media-timeupdate}](#event-media-timeupdate)
        at the element, and [queue a media element
        task](#queue-a-media-element-task){#ready-states:queue-a-media-element-task-4}
        given the [media
        element](#media-element){#ready-states:media-element-9} to [fire
        an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#ready-states:concept-event-fire-4
        x-internal="concept-event-fire"} named
        [`waiting`{#ready-states:event-media-waiting}](#event-media-waiting)
        at the element.

    If the previous ready state was [`HAVE_CURRENT_DATA`{#ready-states:dom-media-have_current_data-6}](#dom-media-have_current_data) or less, and the new ready state is [`HAVE_FUTURE_DATA`{#ready-states:dom-media-have_future_data-6}](#dom-media-have_future_data)

    :   The user agent must [queue a media element
        task](#queue-a-media-element-task){#ready-states:queue-a-media-element-task-5}
        given the [media
        element](#media-element){#ready-states:media-element-10} to
        [fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#ready-states:concept-event-fire-5
        x-internal="concept-event-fire"} named
        [`canplay`{#ready-states:event-media-canplay}](#event-media-canplay)
        at the element.

        If the element\'s
        [`paused`{#ready-states:dom-media-paused}](#dom-media-paused)
        attribute is false, the user agent must [notify about
        playing](#notify-about-playing){#ready-states:notify-about-playing}
        for the element.

    If the new ready state is [`HAVE_ENOUGH_DATA`{#ready-states:dom-media-have_enough_data-2}](#dom-media-have_enough_data)

    :   If the previous ready state was
        [`HAVE_CURRENT_DATA`{#ready-states:dom-media-have_current_data-7}](#dom-media-have_current_data)
        or less, the user agent must [queue a media element
        task](#queue-a-media-element-task){#ready-states:queue-a-media-element-task-6}
        given the [media
        element](#media-element){#ready-states:media-element-11} to
        [fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#ready-states:concept-event-fire-6
        x-internal="concept-event-fire"} named
        [`canplay`{#ready-states:event-media-canplay-2}](#event-media-canplay)
        at the element, and, if the element\'s
        [`paused`{#ready-states:dom-media-paused-2}](#dom-media-paused)
        attribute is false, [notify about
        playing](#notify-about-playing){#ready-states:notify-about-playing-2}
        for the element.

        The user agent must [queue a media element
        task](#queue-a-media-element-task){#ready-states:queue-a-media-element-task-7}
        given the [media
        element](#media-element){#ready-states:media-element-12} to
        [fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#ready-states:concept-event-fire-7
        x-internal="concept-event-fire"} named
        [`canplaythrough`{#ready-states:event-media-canplaythrough}](#event-media-canplaythrough)
        at the element.

        If the element is not [eligible for
        autoplay](#eligible-for-autoplay){#ready-states:eligible-for-autoplay},
        then the user agent must abort these substeps.

        The user agent may run the following substeps:

        1.  Set the
            [`paused`{#ready-states:dom-media-paused-3}](#dom-media-paused)
            attribute to false.
        2.  If the element\'s [show poster
            flag](#show-poster-flag){#ready-states:show-poster-flag} is
            true, set it to false and run the *[time marches
            on](#time-marches-on)* steps.
        3.  [Queue a media element
            task](#queue-a-media-element-task){#ready-states:queue-a-media-element-task-8}
            given the element to [fire an
            event](https://dom.spec.whatwg.org/#concept-event-fire){#ready-states:concept-event-fire-8
            x-internal="concept-event-fire"} named
            [`play`{#ready-states:event-media-play}](#event-media-play)
            at the element.
        4.  [Notify about
            playing](#notify-about-playing){#ready-states:notify-about-playing-3}
            for the element.

        Alternatively, if the element is a
        [`video`{#ready-states:the-video-element-4}](#the-video-element)
        element, the user agent may start observing whether the element
        [intersects the
        viewport](rendering.html#intersect-the-viewport){#ready-states:intersect-the-viewport}.
        When the element starts [intersecting the
        viewport](rendering.html#intersect-the-viewport){#ready-states:intersect-the-viewport-2},
        if the element is still [eligible for
        autoplay](#eligible-for-autoplay){#ready-states:eligible-for-autoplay-2},
        run the substeps above. Optionally, when the element stops
        [intersecting the
        viewport](rendering.html#intersect-the-viewport){#ready-states:intersect-the-viewport-3},
        if the [can autoplay
        flag](#can-autoplay-flag){#ready-states:can-autoplay-flag} is
        still true and the
        [`autoplay`{#ready-states:attr-media-autoplay}](#attr-media-autoplay)
        attribute is still specified, run the following substeps:

        1.  Run the [internal pause
            steps](#internal-pause-steps){#ready-states:internal-pause-steps}
            and set the [can autoplay
            flag](#can-autoplay-flag){#ready-states:can-autoplay-flag-2}
            to true.
        2.  [Queue a media element
            task](#queue-a-media-element-task){#ready-states:queue-a-media-element-task-9}
            given the element to [fire an
            event](https://dom.spec.whatwg.org/#concept-event-fire){#ready-states:concept-event-fire-9
            x-internal="concept-event-fire"} named
            [`pause`{#ready-states:event-media-pause}](#event-media-pause)
            at the element.

        The substeps for playing and pausing can run multiple times as
        the element starts or stops [intersecting the
        viewport](rendering.html#intersect-the-viewport){#ready-states:intersect-the-viewport-4},
        as long as the [can autoplay
        flag](#can-autoplay-flag){#ready-states:can-autoplay-flag-3} is
        true.

        User agents do not need to support autoplay, and it is suggested
        that user agents honor user preferences on the matter. Authors
        are urged to use the
        [`autoplay`{#ready-states:attr-media-autoplay-2}](#attr-media-autoplay)
        attribute rather than using script to force the video to play,
        so as to allow the user to override the behavior if so desired.

It is possible for the ready state of a media element to jump between
these states discontinuously. For example, the state of a media element
can jump straight from
[`HAVE_METADATA`{#ready-states:dom-media-have_metadata-7}](#dom-media-have_metadata)
to
[`HAVE_ENOUGH_DATA`{#ready-states:dom-media-have_enough_data-3}](#dom-media-have_enough_data)
without passing through the
[`HAVE_CURRENT_DATA`{#ready-states:dom-media-have_current_data-8}](#dom-media-have_current_data)
and
[`HAVE_FUTURE_DATA`{#ready-states:dom-media-have_future_data-7}](#dom-media-have_future_data)
states.

The [`readyState`]{#dom-media-readystate .dfn dfn-for="HTMLMediaElement"
dfn-type="attribute"} IDL attribute must, on getting, return the value
described above that describes the current ready state of the [media
element](#media-element){#ready-states:media-element-13}.

The [`autoplay`]{#attr-media-autoplay .dfn dfn-for="audio,video"
dfn-type="element-attr"} attribute is a [boolean
attribute](common-microsyntaxes.html#boolean-attribute){#ready-states:boolean-attribute}.
When present, the user agent (as described in the algorithm described
herein) will automatically begin playback of the [media
resource](#media-resource){#ready-states:media-resource-3} as soon as it
can do so without stopping.

Authors are urged to use the
[`autoplay`{#ready-states:attr-media-autoplay-3}](#attr-media-autoplay)
attribute rather than using script to trigger automatic playback, as
this allows the user to override the automatic playback when it is not
desired, e.g. when using a screen reader. Authors are also encouraged to
consider not using the automatic playback behavior at all, and instead
to let the user agent wait for the user to start playback explicitly.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/autoplay](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/autoplay "The HTMLMediaElement.autoplay property reflects the autoplay HTML attribute, indicating whether playback should automatically begin as soon as enough media is available to do so without interruption.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The [`autoplay`]{#dom-media-autoplay .dfn dfn-for="HTMLMediaElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#ready-states:reflect} the
content attribute of the same name.

##### [4.8.11.8]{.secno} Playing the media resource[](#playing-the-media-resource){.self-link}

`media`{.variable}`.`[`paused`](#dom-media-paused){#dom-media-paused-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/paused](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/paused "The read-only HTMLMediaElement.paused property tells whether the media element is paused.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns true if playback is paused; false otherwise.

`media`{.variable}`.`[`ended`](#dom-media-ended){#dom-media-ended-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/ended](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/ended "The HTMLMediaElement.ended property indicates whether the media element has ended playback.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns true if playback has reached the end of the [media
resource](#media-resource){#playing-the-media-resource:media-resource}.

`media`{.variable}`.`[`defaultPlaybackRate`](#dom-media-defaultplaybackrate){#dom-media-defaultplaybackrate-dev}` [ = ``value`{.variable}` ]`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/defaultPlaybackRate](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/defaultPlaybackRate "The HTMLMediaElement.defaultPlaybackRate property indicates the default playback rate for the media.")

Support in all current engines.

::: support
[Firefox20+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns the default rate of playback, for when the user is not
fast-forwarding or reversing through the [media
resource](#media-resource){#playing-the-media-resource:media-resource-2}.

Can be set, to change the default rate of playback.

The default rate has no direct effect on playback, but if the user
switches to a fast-forward mode, when they return to the normal playback
mode, it is expected that the rate of playback will be returned to the
default rate of playback.

`media`{.variable}`.`[`playbackRate`](#dom-media-playbackrate){#dom-media-playbackrate-dev}` [ = ``value`{.variable}` ]`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/playbackRate](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/playbackRate "The HTMLMediaElement.playbackRate property sets the rate at which the media is being played back. This is used to implement user controls for fast forward, slow motion, and so forth. The normal playback rate is multiplied by this value to obtain the current rate, so a value of 1.0 indicates normal speed.")

Support in all current engines.

::: support
[Firefox20+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns the current rate playback, where 1.0 is normal speed.

Can be set, to change the rate of playback.

`media`{.variable}`.`[`preservesPitch`](#dom-media-preservespitch){#dom-media-preservespitch-dev}

::::: {.mdn-anno .wrapped .before}
MDN

:::: feature
[HTMLMediaElement/preservesPitch](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/preservesPitch "The HTMLMediaElement.preservesPitch property determines whether or not the browser should adjust the pitch of the audio to compensate for changes to the playback rate made by setting HTMLMediaElement.playbackRate.")

::: support
[Firefox101+]{.firefox .yes}[Safari[🔰
4+]{title="Requires a prefix or alternative name."}]{.safari
.yes}[Chrome86+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge86+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS[🔰
4+]{title="Requires a prefix or alternative name."}]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

Returns true if pitch-preserving algorithms are used when the
[`playbackRate`{#playing-the-media-resource:dom-media-playbackrate}](#dom-media-playbackrate)
is not 1.0. The default value is true.

Can be set to false to have the [media
resource](#media-resource){#playing-the-media-resource:media-resource-3}\'s
audio pitch change up or down depending on the
[`playbackRate`{#playing-the-media-resource:dom-media-playbackrate-2}](#dom-media-playbackrate).
This is useful for aesthetic and performance reasons.

`media`{.variable}`.`[`played`](#dom-media-played){#dom-media-played-dev}

Returns a
[`TimeRanges`{#playing-the-media-resource:timeranges}](#timeranges)
object that represents the ranges of the [media
resource](#media-resource){#playing-the-media-resource:media-resource-4}
that the user agent has played.

`media`{.variable}`.`[`play`](#dom-media-play){#dom-media-play-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/play](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/play "The HTMLMediaElement play() method attempts to begin playback of the media. It returns a Promise which is resolved when playback has been successfully started.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

Sets the
[`paused`{#playing-the-media-resource:dom-media-paused}](#dom-media-paused)
attribute to false, loading the [media
resource](#media-resource){#playing-the-media-resource:media-resource-5}
and beginning playback if necessary. If the playback had ended, will
restart it from the start.

`media`{.variable}`.`[`pause`](#dom-media-pause){#dom-media-pause-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/pause](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/pause "The HTMLMediaElement.pause() method will pause playback of the media, if the media is already in a paused state this method will have no effect.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Sets the
[`paused`{#playing-the-media-resource:dom-media-paused-2}](#dom-media-paused)
attribute to true, loading the [media
resource](#media-resource){#playing-the-media-resource:media-resource-6}
if necessary.

The [`paused`]{#dom-media-paused .dfn dfn-for="HTMLMediaElement"
dfn-type="attribute"} attribute represents whether the [media
element](#media-element){#playing-the-media-resource:media-element} is
paused or not. The attribute must initially be true.

A [media
element](#media-element){#playing-the-media-resource:media-element-2} is
a [blocked media element]{#blocked-media-element .dfn} if its
[`readyState`{#playing-the-media-resource:dom-media-readystate}](#dom-media-readystate)
attribute is in the
[`HAVE_NOTHING`{#playing-the-media-resource:dom-media-have_nothing}](#dom-media-have_nothing)
state, the
[`HAVE_METADATA`{#playing-the-media-resource:dom-media-have_metadata}](#dom-media-have_metadata)
state, or the
[`HAVE_CURRENT_DATA`{#playing-the-media-resource:dom-media-have_current_data}](#dom-media-have_current_data)
state, or if the element has [paused for user
interaction](#paused-for-user-interaction){#playing-the-media-resource:paused-for-user-interaction}
or [paused for in-band
content](#paused-for-in-band-content){#playing-the-media-resource:paused-for-in-band-content}.

A [media
element](#media-element){#playing-the-media-resource:media-element-3} is
said to be [potentially playing]{#potentially-playing .dfn
dfn-for="media element" export=""} when its
[`paused`{#playing-the-media-resource:dom-media-paused-3}](#dom-media-paused)
attribute is false, the element has not [ended
playback](#ended-playback){#playing-the-media-resource:ended-playback},
playback has not [stopped due to
errors](#stopped-due-to-errors){#playing-the-media-resource:stopped-due-to-errors},
and the element is not a [blocked media
element](#blocked-media-element){#playing-the-media-resource:blocked-media-element}.

A
[`waiting`{#playing-the-media-resource:event-media-waiting}](#event-media-waiting)
DOM event [can be fired](#fire-waiting-when-waiting) as a result of an
element that is [potentially
playing](#potentially-playing){#playing-the-media-resource:potentially-playing}
stopping playback due to its
[`readyState`{#playing-the-media-resource:dom-media-readystate-2}](#dom-media-readystate)
attribute changing to a value lower than
[`HAVE_FUTURE_DATA`{#playing-the-media-resource:dom-media-have_future_data}](#dom-media-have_future_data).

A [media
element](#media-element){#playing-the-media-resource:media-element-4} is
said to be [eligible for autoplay]{#eligible-for-autoplay .dfn} when all
of the following are true:

- its [can autoplay
  flag](#can-autoplay-flag){#playing-the-media-resource:can-autoplay-flag}
  is true;

- its
  [`paused`{#playing-the-media-resource:dom-media-paused-4}](#dom-media-paused)
  attribute is true;

- it has an
  [`autoplay`{#playing-the-media-resource:attr-media-autoplay}](#attr-media-autoplay)
  attribute specified;

- its [node
  document](https://dom.spec.whatwg.org/#concept-node-document){#playing-the-media-resource:node-document
  x-internal="node-document"}\'s [active sandboxing flag
  set](browsers.html#active-sandboxing-flag-set){#playing-the-media-resource:active-sandboxing-flag-set}
  does not have the [sandboxed automatic features browsing context
  flag](browsers.html#sandboxed-automatic-features-browsing-context-flag){#playing-the-media-resource:sandboxed-automatic-features-browsing-context-flag}
  set; and

- its [node
  document](https://dom.spec.whatwg.org/#concept-node-document){#playing-the-media-resource:node-document-2
  x-internal="node-document"} is [allowed to
  use](iframe-embed-object.html#allowed-to-use){#playing-the-media-resource:allowed-to-use}
  the
  \"[`autoplay`{#playing-the-media-resource:autoplay-feature}](infrastructure.html#autoplay-feature)\"
  feature.

A [media
element](#media-element){#playing-the-media-resource:media-element-5} is
said to be [allowed to play]{#allowed-to-play .dfn} if the user agent
and the system allow media playback in the current context.

For example, a user agent could allow playback only when the [media
element](#media-element){#playing-the-media-resource:media-element-6}\'s
[`Window`{#playing-the-media-resource:window}](nav-history-apis.html#window)
object has [transient
activation](interaction.html#transient-activation){#playing-the-media-resource:transient-activation},
but an exception could be made to allow playback while
[muted](#concept-media-muted){#playing-the-media-resource:concept-media-muted}.

A [media
element](#media-element){#playing-the-media-resource:media-element-7} is
said to have [ended playback]{#ended-playback .dfn} when:

- The element\'s
  [`readyState`{#playing-the-media-resource:dom-media-readystate-3}](#dom-media-readystate)
  attribute is
  [`HAVE_METADATA`{#playing-the-media-resource:dom-media-have_metadata-2}](#dom-media-have_metadata)
  or greater, and

- Either:

  - The [current playback
    position](#current-playback-position){#playing-the-media-resource:current-playback-position}
    is the end of the [media
    resource](#media-resource){#playing-the-media-resource:media-resource-7},
    and

  - The [direction of
    playback](#direction-of-playback){#playing-the-media-resource:direction-of-playback}
    is forwards, and

  - The [media
    element](#media-element){#playing-the-media-resource:media-element-8}
    does not have a
    [`loop`{#playing-the-media-resource:attr-media-loop}](#attr-media-loop)
    attribute specified.

  Or:

  - The [current playback
    position](#current-playback-position){#playing-the-media-resource:current-playback-position-2}
    is the [earliest possible
    position](#earliest-possible-position){#playing-the-media-resource:earliest-possible-position},
    and

  - The [direction of
    playback](#direction-of-playback){#playing-the-media-resource:direction-of-playback-2}
    is backwards.

The [`ended`]{#dom-media-ended .dfn dfn-for="HTMLMediaElement"
dfn-type="attribute"} attribute must return true if, the last time the
[event
loop](webappapis.html#event-loop){#playing-the-media-resource:event-loop}
reached [step 1](webappapis.html#step1), the [media
element](#media-element){#playing-the-media-resource:media-element-9}
had [ended
playback](#ended-playback){#playing-the-media-resource:ended-playback-2}
and the [direction of
playback](#direction-of-playback){#playing-the-media-resource:direction-of-playback-3}
was forwards, and false otherwise.

A [media
element](#media-element){#playing-the-media-resource:media-element-10}
is said to have [stopped due to errors]{#stopped-due-to-errors .dfn}
when the element\'s
[`readyState`{#playing-the-media-resource:dom-media-readystate-4}](#dom-media-readystate)
attribute is
[`HAVE_METADATA`{#playing-the-media-resource:dom-media-have_metadata-3}](#dom-media-have_metadata)
or greater, and the user agent [encounters a non-fatal
error](#non-fatal-media-error) during the processing of the [media
data](#media-data){#playing-the-media-resource:media-data}, and due to
that error, is not able to play the content at the [current playback
position](#current-playback-position){#playing-the-media-resource:current-playback-position-3}.

A [media
element](#media-element){#playing-the-media-resource:media-element-11}
is said to have [paused for user
interaction]{#paused-for-user-interaction .dfn} when its
[`paused`{#playing-the-media-resource:dom-media-paused-5}](#dom-media-paused)
attribute is false, the
[`readyState`{#playing-the-media-resource:dom-media-readystate-5}](#dom-media-readystate)
attribute is either
[`HAVE_FUTURE_DATA`{#playing-the-media-resource:dom-media-have_future_data-2}](#dom-media-have_future_data)
or
[`HAVE_ENOUGH_DATA`{#playing-the-media-resource:dom-media-have_enough_data}](#dom-media-have_enough_data)
and the user agent has reached a point in the [media
resource](#media-resource){#playing-the-media-resource:media-resource-8}
where the user has to make a selection for the resource to continue.

It is possible for a [media
element](#media-element){#playing-the-media-resource:media-element-12}
to have both [ended
playback](#ended-playback){#playing-the-media-resource:ended-playback-3}
and [paused for user
interaction](#paused-for-user-interaction){#playing-the-media-resource:paused-for-user-interaction-2}
at the same time.

When a [media
element](#media-element){#playing-the-media-resource:media-element-13}
that is [potentially
playing](#potentially-playing){#playing-the-media-resource:potentially-playing-2}
stops playing because it has [paused for user
interaction](#paused-for-user-interaction){#playing-the-media-resource:paused-for-user-interaction-3},
the user agent must [queue a media element
task](#queue-a-media-element-task){#playing-the-media-resource:queue-a-media-element-task}
given the [media
element](#media-element){#playing-the-media-resource:media-element-14}
to [fire an
event](https://dom.spec.whatwg.org/#concept-event-fire){#playing-the-media-resource:concept-event-fire
x-internal="concept-event-fire"} named
[`timeupdate`{#playing-the-media-resource:event-media-timeupdate}](#event-media-timeupdate)
at the element.

A [media
element](#media-element){#playing-the-media-resource:media-element-15}
is said to have [paused for in-band content]{#paused-for-in-band-content
.dfn} when its
[`paused`{#playing-the-media-resource:dom-media-paused-6}](#dom-media-paused)
attribute is false, the
[`readyState`{#playing-the-media-resource:dom-media-readystate-6}](#dom-media-readystate)
attribute is either
[`HAVE_FUTURE_DATA`{#playing-the-media-resource:dom-media-have_future_data-3}](#dom-media-have_future_data)
or
[`HAVE_ENOUGH_DATA`{#playing-the-media-resource:dom-media-have_enough_data-2}](#dom-media-have_enough_data)
and the user agent has suspended playback of the [media
resource](#media-resource){#playing-the-media-resource:media-resource-9}
in order to play content that is temporally anchored to the [media
resource](#media-resource){#playing-the-media-resource:media-resource-10}
and has a nonzero length, or to play content that is temporally anchored
to a segment of the [media
resource](#media-resource){#playing-the-media-resource:media-resource-11}
but has a length longer than that segment.

One example of when a [media
element](#media-element){#playing-the-media-resource:media-element-16}
would be [paused for in-band
content](#paused-for-in-band-content){#playing-the-media-resource:paused-for-in-band-content-2}
is when the user agent is playing [audio
descriptions](#attr-track-kind-descriptions){#playing-the-media-resource:attr-track-kind-descriptions}
from an external WebVTT file, and the synthesized speech generated for a
cue is longer than the time between the [text track cue start
time](#text-track-cue-start-time){#playing-the-media-resource:text-track-cue-start-time}
and the [text track cue end
time](#text-track-cue-end-time){#playing-the-media-resource:text-track-cue-end-time}.

------------------------------------------------------------------------

When the [current playback
position](#current-playback-position){#playing-the-media-resource:current-playback-position-4}
reaches the end of the [media
resource](#media-resource){#playing-the-media-resource:media-resource-12}
when the [direction of
playback](#direction-of-playback){#playing-the-media-resource:direction-of-playback-4}
is forwards, then the user agent must follow these steps:

1.  If the [media
    element](#media-element){#playing-the-media-resource:media-element-17}
    has a
    [`loop`{#playing-the-media-resource:attr-media-loop-2}](#attr-media-loop)
    attribute specified, then
    [seek](#dom-media-seek){#playing-the-media-resource:dom-media-seek}
    to the [earliest possible
    position](#earliest-possible-position){#playing-the-media-resource:earliest-possible-position-2}
    of the [media
    resource](#media-resource){#playing-the-media-resource:media-resource-13}
    and return.

2.  As defined above, the
    [`ended`{#playing-the-media-resource:dom-media-ended}](#dom-media-ended)
    IDL attribute starts returning true once the [event
    loop](webappapis.html#event-loop){#playing-the-media-resource:event-loop-2}
    returns to [step 1](webappapis.html#step1).

3.  [Queue a media element
    task](#queue-a-media-element-task){#playing-the-media-resource:queue-a-media-element-task-2}
    given the [media
    element](#media-element){#playing-the-media-resource:media-element-18}
    and the following steps:

    1.  [Fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#playing-the-media-resource:concept-event-fire-2
        x-internal="concept-event-fire"} named
        [`timeupdate`{#playing-the-media-resource:event-media-timeupdate-2}](#event-media-timeupdate)
        at the [media
        element](#media-element){#playing-the-media-resource:media-element-19}.

    2.  If the [media
        element](#media-element){#playing-the-media-resource:media-element-20}
        has [ended
        playback](#ended-playback){#playing-the-media-resource:ended-playback-4},
        the [direction of
        playback](#direction-of-playback){#playing-the-media-resource:direction-of-playback-5}
        is forwards, and
        [paused](#dom-media-paused){#playing-the-media-resource:dom-media-paused-7}
        is false, then:

        1.  Set the
            [`paused`{#playing-the-media-resource:dom-media-paused-8}](#dom-media-paused)
            attribute to true.

        2.  [Fire an
            event](https://dom.spec.whatwg.org/#concept-event-fire){#playing-the-media-resource:concept-event-fire-3
            x-internal="concept-event-fire"} named
            [`pause`{#playing-the-media-resource:event-media-pause}](#event-media-pause)
            at the [media
            element](#media-element){#playing-the-media-resource:media-element-21}.

        3.  [Take pending play
            promises](#take-pending-play-promises){#playing-the-media-resource:take-pending-play-promises}
            and [reject pending play
            promises](#reject-pending-play-promises){#playing-the-media-resource:reject-pending-play-promises}
            with the result and an
            [\"`AbortError`\"](https://webidl.spec.whatwg.org/#aborterror){#playing-the-media-resource:aborterror
            x-internal="aborterror"}
            [`DOMException`{#playing-the-media-resource:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    3.  [Fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#playing-the-media-resource:concept-event-fire-4
        x-internal="concept-event-fire"} named
        [`ended`{#playing-the-media-resource:event-media-ended}](#event-media-ended)
        at the [media
        element](#media-element){#playing-the-media-resource:media-element-22}.

When the [current playback
position](#current-playback-position){#playing-the-media-resource:current-playback-position-5}
reaches the [earliest possible
position](#earliest-possible-position){#playing-the-media-resource:earliest-possible-position-3}
of the [media
resource](#media-resource){#playing-the-media-resource:media-resource-14}
when the [direction of
playback](#direction-of-playback){#playing-the-media-resource:direction-of-playback-6}
is backwards, then the user agent must only [queue a media element
task](#queue-a-media-element-task){#playing-the-media-resource:queue-a-media-element-task-3}
given the [media
element](#media-element){#playing-the-media-resource:media-element-23}
to [fire an
event](https://dom.spec.whatwg.org/#concept-event-fire){#playing-the-media-resource:concept-event-fire-5
x-internal="concept-event-fire"} named
[`timeupdate`{#playing-the-media-resource:event-media-timeupdate-3}](#event-media-timeupdate)
at the element.

The word \"reaches\" here does not imply that the [current playback
position](#current-playback-position){#playing-the-media-resource:current-playback-position-6}
needs to have changed during normal playback; it could be via
[seeking](#dom-media-seek){#playing-the-media-resource:dom-media-seek-2},
for instance.

------------------------------------------------------------------------

The [`defaultPlaybackRate`]{#dom-media-defaultplaybackrate .dfn
dfn-for="HTMLMediaElement" dfn-type="attribute"} attribute gives the
desired speed at which the [media
resource](#media-resource){#playing-the-media-resource:media-resource-15}
is to play, as a multiple of its intrinsic speed. The attribute is
mutable: on getting it must return the last value it was set to, or 1.0
if it hasn\'t yet been set; on setting the attribute must be set to the
new value.

The
[`defaultPlaybackRate`{#playing-the-media-resource:dom-media-defaultplaybackrate}](#dom-media-defaultplaybackrate)
is used by the user agent when it [exposes a user interface to the
user](#expose-a-user-interface-to-the-user){#playing-the-media-resource:expose-a-user-interface-to-the-user}.

The [`playbackRate`]{#dom-media-playbackrate .dfn
dfn-for="HTMLMediaElement" dfn-type="attribute"} attribute gives the
effective playback rate, which is the speed at which the [media
resource](#media-resource){#playing-the-media-resource:media-resource-16}
plays, as a multiple of its intrinsic speed. If it is not equal to the
[`defaultPlaybackRate`{#playing-the-media-resource:dom-media-defaultplaybackrate-2}](#dom-media-defaultplaybackrate),
then the implication is that the user is using a feature such as fast
forward or slow motion playback. The attribute is mutable: on getting it
must return the last value it was set to, or 1.0 if it hasn\'t yet been
set; on setting, the user agent must follow these steps:

1.  If the given value is not supported by the user agent, then throw a
    [\"`NotSupportedError`\"](https://webidl.spec.whatwg.org/#notsupportederror){#playing-the-media-resource:notsupportederror
    x-internal="notsupportederror"}
    [`DOMException`{#playing-the-media-resource:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  Set
    [`playbackRate`{#playing-the-media-resource:dom-media-playbackrate-3}](#dom-media-playbackrate)
    to the new value, and if the element is [potentially
    playing](#potentially-playing){#playing-the-media-resource:potentially-playing-3},
    change the playback speed.

When the
[`defaultPlaybackRate`{#playing-the-media-resource:dom-media-defaultplaybackrate-3}](#dom-media-defaultplaybackrate)
or
[`playbackRate`{#playing-the-media-resource:dom-media-playbackrate-4}](#dom-media-playbackrate)
attributes change value (either by being set by script or by being
changed directly by the user agent, e.g. in response to user control)
the user agent must [queue a media element
task](#queue-a-media-element-task){#playing-the-media-resource:queue-a-media-element-task-4}
given the [media
element](#media-element){#playing-the-media-resource:media-element-24}
to [fire an
event](https://dom.spec.whatwg.org/#concept-event-fire){#playing-the-media-resource:concept-event-fire-6
x-internal="concept-event-fire"} named
[`ratechange`{#playing-the-media-resource:event-media-ratechange}](#event-media-ratechange)
at the [media
element](#media-element){#playing-the-media-resource:media-element-25}.
The user agent must process attribute changes smoothly and must not
introduce any perceivable gaps or muting of playback in response.

The [`preservesPitch`]{#dom-media-preservespitch .dfn
dfn-for="HTMLMediaElement" dfn-type="attribute"} getter steps are to
return true if a pitch-preserving algorithm is in effect during
playback. The setter steps are to correspondingly switch the
pitch-preserving algorithm on or off, without any perceivable gaps or
muting of playback. By default, such a pitch-preserving algorithm must
be in effect (i.e., the getter will initially return true).

------------------------------------------------------------------------

The [`played`]{#dom-media-played .dfn dfn-for="HTMLMediaElement"
dfn-type="attribute"} attribute must return a new static [normalized
`TimeRanges`
object](#normalised-timeranges-object){#playing-the-media-resource:normalised-timeranges-object}
that represents the ranges of points on the [media
timeline](#media-timeline){#playing-the-media-resource:media-timeline}
of the [media
resource](#media-resource){#playing-the-media-resource:media-resource-17}
reached through the usual monotonic increase of the [current playback
position](#current-playback-position){#playing-the-media-resource:current-playback-position-7}
during normal playback, if any, at the time the attribute is evaluated.

Returning a new object each time is a bad pattern for attribute getters
and is only enshrined here as it would be costly to change it. It is not
to be copied to new APIs.

------------------------------------------------------------------------

Each [media
element](#media-element){#playing-the-media-resource:media-element-26}
has a [list of pending play promises]{#list-of-pending-play-promises
.dfn}, which must initially be empty.

To [take pending play promises]{#take-pending-play-promises .dfn} for a
[media
element](#media-element){#playing-the-media-resource:media-element-27},
the user agent must run the following steps:

1.  Let `promises`{.variable} be an empty list of promises.
2.  Copy the [media
    element](#media-element){#playing-the-media-resource:media-element-28}\'s
    [list of pending play
    promises](#list-of-pending-play-promises){#playing-the-media-resource:list-of-pending-play-promises}
    to `promises`{.variable}.
3.  Clear the [media
    element](#media-element){#playing-the-media-resource:media-element-29}\'s
    [list of pending play
    promises](#list-of-pending-play-promises){#playing-the-media-resource:list-of-pending-play-promises-2}.
4.  Return `promises`{.variable}.

To [resolve pending play promises]{#resolve-pending-play-promises .dfn}
for a [media
element](#media-element){#playing-the-media-resource:media-element-30}
with a list of promises `promises`{.variable}, the user agent must
resolve each promise in `promises`{.variable} with undefined.

To [reject pending play promises]{#reject-pending-play-promises .dfn}
for a [media
element](#media-element){#playing-the-media-resource:media-element-31}
with a list of promises `promises`{.variable} and an exception name
`error`{.variable}, the user agent must reject each promise in
`promises`{.variable} with `error`{.variable}.

To [notify about playing]{#notify-about-playing .dfn} for a [media
element](#media-element){#playing-the-media-resource:media-element-32},
the user agent must run the following steps:

1.  [Take pending play
    promises](#take-pending-play-promises){#playing-the-media-resource:take-pending-play-promises-2}
    and let `promises`{.variable} be the result.

2.  [Queue a media element
    task](#queue-a-media-element-task){#playing-the-media-resource:queue-a-media-element-task-5}
    given the element and the following steps:

    1.  [Fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#playing-the-media-resource:concept-event-fire-7
        x-internal="concept-event-fire"} named
        [`playing`{#playing-the-media-resource:event-media-playing}](#event-media-playing)
        at the element.

    2.  [Resolve pending play
        promises](#resolve-pending-play-promises){#playing-the-media-resource:resolve-pending-play-promises}
        with `promises`{.variable}.

When the [`play()`]{#dom-media-play .dfn dfn-for="HTMLMediaElement"
dfn-type="method"} method on a [media
element](#media-element){#playing-the-media-resource:media-element-33}
is invoked, the user agent must run the following steps.

1.  If the [media
    element](#media-element){#playing-the-media-resource:media-element-34}
    is not [allowed to
    play](#allowed-to-play){#playing-the-media-resource:allowed-to-play},
    then return [a promise rejected
    with](https://webidl.spec.whatwg.org/#a-promise-rejected-with){#playing-the-media-resource:a-promise-rejected-with
    x-internal="a-promise-rejected-with"} a
    [\"`NotAllowedError`\"](https://webidl.spec.whatwg.org/#notallowederror){#playing-the-media-resource:notallowederror
    x-internal="notallowederror"}
    [`DOMException`{#playing-the-media-resource:domexception-3}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  If the [media
    element](#media-element){#playing-the-media-resource:media-element-35}\'s
    [`error`{#playing-the-media-resource:dom-media-error}](#dom-media-error)
    attribute is not null and its
    [code](#concept-mediaerror-code){#playing-the-media-resource:concept-mediaerror-code}
    is
    [`MEDIA_ERR_SRC_NOT_SUPPORTED`{#playing-the-media-resource:dom-mediaerror-media_err_src_not_supported}](#dom-mediaerror-media_err_src_not_supported),
    then return [a promise rejected
    with](https://webidl.spec.whatwg.org/#a-promise-rejected-with){#playing-the-media-resource:a-promise-rejected-with-2
    x-internal="a-promise-rejected-with"} a
    [\"`NotSupportedError`\"](https://webidl.spec.whatwg.org/#notsupportederror){#playing-the-media-resource:notsupportederror-2
    x-internal="notsupportederror"}
    [`DOMException`{#playing-the-media-resource:domexception-4}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    This means that the [dedicated media source failure
    steps](#dedicated-media-source-failure-steps){#playing-the-media-resource:dedicated-media-source-failure-steps}
    have run. Playback is not possible until the [media element load
    algorithm](#media-element-load-algorithm){#playing-the-media-resource:media-element-load-algorithm}
    clears the
    [`error`{#playing-the-media-resource:dom-media-error-2}](#dom-media-error)
    attribute.

3.  Let `promise`{.variable} be a new promise and append
    `promise`{.variable} to the [list of pending play
    promises](#list-of-pending-play-promises){#playing-the-media-resource:list-of-pending-play-promises-3}.

4.  Run the [internal play
    steps](#internal-play-steps){#playing-the-media-resource:internal-play-steps}
    for the [media
    element](#media-element){#playing-the-media-resource:media-element-36}.

5.  Return `promise`{.variable}.

The [internal play steps]{#internal-play-steps .dfn} for a [media
element](#media-element){#playing-the-media-resource:media-element-37}
are as follows:

1.  If the [media
    element](#media-element){#playing-the-media-resource:media-element-38}\'s
    [`networkState`{#playing-the-media-resource:dom-media-networkstate}](#dom-media-networkstate)
    attribute has the value
    [`NETWORK_EMPTY`{#playing-the-media-resource:dom-media-network_empty}](#dom-media-network_empty),
    invoke the [media
    element](#media-element){#playing-the-media-resource:media-element-39}\'s
    [resource selection
    algorithm](#concept-media-load-algorithm){#playing-the-media-resource:concept-media-load-algorithm}.

2.  If the [playback has
    ended](#ended-playback){#playing-the-media-resource:ended-playback-5}
    and the [direction of
    playback](#direction-of-playback){#playing-the-media-resource:direction-of-playback-7}
    is forwards,
    [seek](#dom-media-seek){#playing-the-media-resource:dom-media-seek-3}
    to the [earliest possible
    position](#earliest-possible-position){#playing-the-media-resource:earliest-possible-position-4}
    of the [media
    resource](#media-resource){#playing-the-media-resource:media-resource-18}.

    This [will cause](#seekUpdate) the user agent to [queue a media
    element
    task](#queue-a-media-element-task){#playing-the-media-resource:queue-a-media-element-task-6}
    given the [media
    element](#media-element){#playing-the-media-resource:media-element-40}
    to [fire an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#playing-the-media-resource:concept-event-fire-8
    x-internal="concept-event-fire"} named
    [`timeupdate`{#playing-the-media-resource:event-media-timeupdate-4}](#event-media-timeupdate)
    at the [media
    element](#media-element){#playing-the-media-resource:media-element-41}.

3.  If the [media
    element](#media-element){#playing-the-media-resource:media-element-42}\'s
    [`paused`{#playing-the-media-resource:dom-media-paused-9}](#dom-media-paused)
    attribute is true, then:

    1.  Change the value of
        [`paused`{#playing-the-media-resource:dom-media-paused-10}](#dom-media-paused)
        to false.

    2.  If the [show poster
        flag](#show-poster-flag){#playing-the-media-resource:show-poster-flag}
        is true, set the element\'s [show poster
        flag](#show-poster-flag){#playing-the-media-resource:show-poster-flag-2}
        to false and run the *[time marches on](#time-marches-on)*
        steps.

    3.  [Queue a media element
        task](#queue-a-media-element-task){#playing-the-media-resource:queue-a-media-element-task-7}
        given the [media
        element](#media-element){#playing-the-media-resource:media-element-43}
        to [fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#playing-the-media-resource:concept-event-fire-9
        x-internal="concept-event-fire"} named
        [`play`{#playing-the-media-resource:event-media-play}](#event-media-play)
        at the element.

    4.  If the [media
        element](#media-element){#playing-the-media-resource:media-element-44}\'s
        [`readyState`{#playing-the-media-resource:dom-media-readystate-7}](#dom-media-readystate)
        attribute has the value
        [`HAVE_NOTHING`{#playing-the-media-resource:dom-media-have_nothing-2}](#dom-media-have_nothing),
        [`HAVE_METADATA`{#playing-the-media-resource:dom-media-have_metadata-4}](#dom-media-have_metadata),
        or
        [`HAVE_CURRENT_DATA`{#playing-the-media-resource:dom-media-have_current_data-2}](#dom-media-have_current_data),
        [queue a media element
        task](#queue-a-media-element-task){#playing-the-media-resource:queue-a-media-element-task-8}
        given the [media
        element](#media-element){#playing-the-media-resource:media-element-45}
        to [fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#playing-the-media-resource:concept-event-fire-10
        x-internal="concept-event-fire"} named
        [`waiting`{#playing-the-media-resource:event-media-waiting-2}](#event-media-waiting)
        at the element.

        Otherwise, the [media
        element](#media-element){#playing-the-media-resource:media-element-46}\'s
        [`readyState`{#playing-the-media-resource:dom-media-readystate-8}](#dom-media-readystate)
        attribute has the value
        [`HAVE_FUTURE_DATA`{#playing-the-media-resource:dom-media-have_future_data-4}](#dom-media-have_future_data)
        or
        [`HAVE_ENOUGH_DATA`{#playing-the-media-resource:dom-media-have_enough_data-3}](#dom-media-have_enough_data):
        [notify about
        playing](#notify-about-playing){#playing-the-media-resource:notify-about-playing}
        for the element.

4.  Otherwise, if the [media
    element](#media-element){#playing-the-media-resource:media-element-47}\'s
    [`readyState`{#playing-the-media-resource:dom-media-readystate-9}](#dom-media-readystate)
    attribute has the value
    [`HAVE_FUTURE_DATA`{#playing-the-media-resource:dom-media-have_future_data-5}](#dom-media-have_future_data)
    or
    [`HAVE_ENOUGH_DATA`{#playing-the-media-resource:dom-media-have_enough_data-4}](#dom-media-have_enough_data),
    [take pending play
    promises](#take-pending-play-promises){#playing-the-media-resource:take-pending-play-promises-3}
    and [queue a media element
    task](#queue-a-media-element-task){#playing-the-media-resource:queue-a-media-element-task-9}
    given the [media
    element](#media-element){#playing-the-media-resource:media-element-48}
    to [resolve pending play
    promises](#resolve-pending-play-promises){#playing-the-media-resource:resolve-pending-play-promises-2}
    with the result.

    The media element is already playing. However, it\'s possible that
    `promise`{.variable} will be
    [rejected](#reject-pending-play-promises){#playing-the-media-resource:reject-pending-play-promises-2}
    before the queued task is run.

5.  Set the [media
    element](#media-element){#playing-the-media-resource:media-element-49}\'s
    [can autoplay
    flag](#can-autoplay-flag){#playing-the-media-resource:can-autoplay-flag-2}
    to false.

------------------------------------------------------------------------

When the [`pause()`]{#dom-media-pause .dfn dfn-for="HTMLMediaElement"
dfn-type="method"} method is invoked, and when the user agent is
required to pause the [media
element](#media-element){#playing-the-media-resource:media-element-50},
the user agent must run the following steps:

1.  If the [media
    element](#media-element){#playing-the-media-resource:media-element-51}\'s
    [`networkState`{#playing-the-media-resource:dom-media-networkstate-2}](#dom-media-networkstate)
    attribute has the value
    [`NETWORK_EMPTY`{#playing-the-media-resource:dom-media-network_empty-2}](#dom-media-network_empty),
    invoke the [media
    element](#media-element){#playing-the-media-resource:media-element-52}\'s
    [resource selection
    algorithm](#concept-media-load-algorithm){#playing-the-media-resource:concept-media-load-algorithm-2}.

2.  Run the [internal pause
    steps](#internal-pause-steps){#playing-the-media-resource:internal-pause-steps}
    for the [media
    element](#media-element){#playing-the-media-resource:media-element-53}.

The [internal pause steps]{#internal-pause-steps .dfn} for a [media
element](#media-element){#playing-the-media-resource:media-element-54}
are as follows:

1.  Set the [media
    element](#media-element){#playing-the-media-resource:media-element-55}\'s
    [can autoplay
    flag](#can-autoplay-flag){#playing-the-media-resource:can-autoplay-flag-3}
    to false.

2.  If the [media
    element](#media-element){#playing-the-media-resource:media-element-56}\'s
    [`paused`{#playing-the-media-resource:dom-media-paused-11}](#dom-media-paused)
    attribute is false, run the following steps:

    1.  Change the value of
        [`paused`{#playing-the-media-resource:dom-media-paused-12}](#dom-media-paused)
        to true.

    2.  [Take pending play
        promises](#take-pending-play-promises){#playing-the-media-resource:take-pending-play-promises-4}
        and let `promises`{.variable} be the result.

    3.  [Queue a media element
        task](#queue-a-media-element-task){#playing-the-media-resource:queue-a-media-element-task-10}
        given the [media
        element](#media-element){#playing-the-media-resource:media-element-57}
        and the following steps:

        1.  [Fire an
            event](https://dom.spec.whatwg.org/#concept-event-fire){#playing-the-media-resource:concept-event-fire-11
            x-internal="concept-event-fire"} named
            [`timeupdate`{#playing-the-media-resource:event-media-timeupdate-5}](#event-media-timeupdate)
            at the element.

        2.  [Fire an
            event](https://dom.spec.whatwg.org/#concept-event-fire){#playing-the-media-resource:concept-event-fire-12
            x-internal="concept-event-fire"} named
            [`pause`{#playing-the-media-resource:event-media-pause-2}](#event-media-pause)
            at the element.

        3.  [Reject pending play
            promises](#reject-pending-play-promises){#playing-the-media-resource:reject-pending-play-promises-3}
            with `promises`{.variable} and an
            [\"`AbortError`\"](https://webidl.spec.whatwg.org/#aborterror){#playing-the-media-resource:aborterror-2
            x-internal="aborterror"}
            [`DOMException`{#playing-the-media-resource:domexception-5}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    4.  Set the [official playback
        position](#official-playback-position){#playing-the-media-resource:official-playback-position}
        to the [current playback
        position](#current-playback-position){#playing-the-media-resource:current-playback-position-8}.

------------------------------------------------------------------------

If the element\'s
[`playbackRate`{#playing-the-media-resource:dom-media-playbackrate-5}](#dom-media-playbackrate)
is positive or zero, then the [direction of
playback]{#direction-of-playback .dfn} is forwards. Otherwise, it is
backwards.

When a [media
element](#media-element){#playing-the-media-resource:media-element-58}
is [potentially
playing](#potentially-playing){#playing-the-media-resource:potentially-playing-4}
and its
[`Document`{#playing-the-media-resource:document}](dom.html#document) is
a [fully
active](document-sequences.html#fully-active){#playing-the-media-resource:fully-active}
[`Document`{#playing-the-media-resource:document-2}](dom.html#document),
its [current playback
position](#current-playback-position){#playing-the-media-resource:current-playback-position-9}
must increase monotonically at the element\'s
[`playbackRate`{#playing-the-media-resource:dom-media-playbackrate-6}](#dom-media-playbackrate)
units of media time per unit time of the [media
timeline](#media-timeline){#playing-the-media-resource:media-timeline-2}\'s
clock. (This specification always refers to this as an *increase*, but
that increase could actually be a *de*crease if the element\'s
[`playbackRate`{#playing-the-media-resource:dom-media-playbackrate-7}](#dom-media-playbackrate)
is negative.)

The element\'s
[`playbackRate`{#playing-the-media-resource:dom-media-playbackrate-8}](#dom-media-playbackrate)
can be 0.0, in which case the [current playback
position](#current-playback-position){#playing-the-media-resource:current-playback-position-10}
doesn\'t move, despite playback not being paused
([`paused`{#playing-the-media-resource:dom-media-paused-13}](#dom-media-paused)
doesn\'t become true, and the
[`pause`{#playing-the-media-resource:event-media-pause-3}](#event-media-pause)
event doesn\'t fire).

This specification doesn\'t define how the user agent achieves the
appropriate playback rate --- depending on the protocol and media
available, it is plausible that the user agent could negotiate with the
server to have the server provide the media data at the appropriate
rate, so that (except for the period between when the rate is changed
and when the server updates the stream\'s playback rate) the client
doesn\'t actually have to drop or interpolate any frames.

Any time the user agent [provides a stable
state](webappapis.html#await-a-stable-state){#playing-the-media-resource:await-a-stable-state},
the [official playback
position](#official-playback-position){#playing-the-media-resource:official-playback-position-2}
must be set to the [current playback
position](#current-playback-position){#playing-the-media-resource:current-playback-position-11}.

While the [direction of
playback](#direction-of-playback){#playing-the-media-resource:direction-of-playback-8}
is backwards, any corresponding audio must be
[muted](#concept-media-muted){#playing-the-media-resource:concept-media-muted-2}.
While the element\'s
[`playbackRate`{#playing-the-media-resource:dom-media-playbackrate-9}](#dom-media-playbackrate)
is so low or so high that the user agent cannot play audio usefully, the
corresponding audio must also be
[muted](#concept-media-muted){#playing-the-media-resource:concept-media-muted-3}.
If the element\'s
[`playbackRate`{#playing-the-media-resource:dom-media-playbackrate-10}](#dom-media-playbackrate)
is not 1.0 and
[`preservesPitch`{#playing-the-media-resource:dom-media-preservespitch}](#dom-media-preservespitch)
is true, the user agent must apply pitch adjustment to preserve the
original pitch of the audio. Otherwise, the user agent must speed up or
slow down the audio without any pitch adjustment.

When a [media
element](#media-element){#playing-the-media-resource:media-element-59}
is [potentially
playing](#potentially-playing){#playing-the-media-resource:potentially-playing-5},
its audio data played must be synchronized with the [current playback
position](#current-playback-position){#playing-the-media-resource:current-playback-position-12},
at the element\'s [effective media
volume](#effective-media-volume){#playing-the-media-resource:effective-media-volume}.
The user agent must play the audio from audio tracks that were enabled
when the [event
loop](webappapis.html#event-loop){#playing-the-media-resource:event-loop-3}
last reached [step 1](webappapis.html#step1).

When a [media
element](#media-element){#playing-the-media-resource:media-element-60}
is not [potentially
playing](#potentially-playing){#playing-the-media-resource:potentially-playing-6},
audio must not play for the element.

[Media
elements](#media-element){#playing-the-media-resource:media-element-61}
that are [potentially
playing](#potentially-playing){#playing-the-media-resource:potentially-playing-7}
while not [in a
document](https://dom.spec.whatwg.org/#in-a-document){#playing-the-media-resource:in-a-document
x-internal="in-a-document"} must not play any video, but should play any
audio component. Media elements must not stop playing just because all
references to them have been removed; only once a media element is in a
state where no further audio could ever be played by that element may
the element be garbage collected.

It is possible for an element to which no explicit references exist to
play audio, even if such an element is not still actively playing: for
instance, it could be unpaused but stalled waiting for content to
buffer, or it could be still buffering, but with a
[`suspend`{#playing-the-media-resource:event-media-suspend}](#event-media-suspend)
event listener that begins playback. Even a media element whose [media
resource](#media-resource){#playing-the-media-resource:media-resource-19}
has no audio tracks could eventually play audio again if it had an event
listener that changes the [media
resource](#media-resource){#playing-the-media-resource:media-resource-20}.

------------------------------------------------------------------------

Each [media
element](#media-element){#playing-the-media-resource:media-element-62}
has a [list of newly introduced cues]{#list-of-newly-introduced-cues
.dfn}, which must be initially empty. Whenever a [text track
cue](#text-track-cue){#playing-the-media-resource:text-track-cue} is
added to the [list of
cues](#text-track-list-of-cues){#playing-the-media-resource:text-track-list-of-cues}
of a [text track](#text-track){#playing-the-media-resource:text-track}
that is in the [list of text
tracks](#list-of-text-tracks){#playing-the-media-resource:list-of-text-tracks}
for a [media
element](#media-element){#playing-the-media-resource:media-element-63},
that
[cue](#text-track-cue){#playing-the-media-resource:text-track-cue-2}
must be added to the [media
element](#media-element){#playing-the-media-resource:media-element-64}\'s
[list of newly introduced
cues](#list-of-newly-introduced-cues){#playing-the-media-resource:list-of-newly-introduced-cues}.
Whenever a [text
track](#text-track){#playing-the-media-resource:text-track-2} is added
to the [list of text
tracks](#list-of-text-tracks){#playing-the-media-resource:list-of-text-tracks-2}
for a [media
element](#media-element){#playing-the-media-resource:media-element-65},
all of the
[cues](#text-track-cue){#playing-the-media-resource:text-track-cue-3} in
that [text
track](#text-track){#playing-the-media-resource:text-track-3}\'s [list
of
cues](#text-track-list-of-cues){#playing-the-media-resource:text-track-list-of-cues-2}
must be added to the [media
element](#media-element){#playing-the-media-resource:media-element-66}\'s
[list of newly introduced
cues](#list-of-newly-introduced-cues){#playing-the-media-resource:list-of-newly-introduced-cues-2}.
When a [media
element](#media-element){#playing-the-media-resource:media-element-67}\'s
[list of newly introduced
cues](#list-of-newly-introduced-cues){#playing-the-media-resource:list-of-newly-introduced-cues-3}
has new cues added while the [media
element](#media-element){#playing-the-media-resource:media-element-68}\'s
[show poster
flag](#show-poster-flag){#playing-the-media-resource:show-poster-flag-3}
is not set, then the user agent must run the *[time marches
on](#time-marches-on)* steps.

When a [text track
cue](#text-track-cue){#playing-the-media-resource:text-track-cue-4} is
removed from the [list of
cues](#text-track-list-of-cues){#playing-the-media-resource:text-track-list-of-cues-3}
of a [text track](#text-track){#playing-the-media-resource:text-track-4}
that is in the [list of text
tracks](#list-of-text-tracks){#playing-the-media-resource:list-of-text-tracks-3}
for a [media
element](#media-element){#playing-the-media-resource:media-element-69},
and whenever a [text
track](#text-track){#playing-the-media-resource:text-track-5} is removed
from the [list of text
tracks](#list-of-text-tracks){#playing-the-media-resource:list-of-text-tracks-4}
of a [media
element](#media-element){#playing-the-media-resource:media-element-70},
if the [media
element](#media-element){#playing-the-media-resource:media-element-71}\'s
[show poster
flag](#show-poster-flag){#playing-the-media-resource:show-poster-flag-4}
is not set, then the user agent must run the *[time marches
on](#time-marches-on)* steps.

When the [current playback
position](#current-playback-position){#playing-the-media-resource:current-playback-position-13}
of a [media
element](#media-element){#playing-the-media-resource:media-element-72}
changes (e.g. due to playback or seeking), the user agent must run the
*[time marches on](#time-marches-on)* steps. To support use cases that
depend on the timing accuracy of cue event firing, such as synchronizing
captions with shot changes in a video, user agents should fire cue
events as close as possible to their position on the media timeline, and
ideally within 20 milliseconds. If the [current playback
position](#current-playback-position){#playing-the-media-resource:current-playback-position-14}
changes while the steps are running, then the user agent must wait for
the steps to complete, and then must immediately rerun the steps. These
steps are thus run as often as possible or needed.

If one iteration takes a long time, this can cause short duration
[cues](#text-track-cue){#playing-the-media-resource:text-track-cue-5} to
be skipped over as the user agent rushes ahead to \"catch up\", so these
cues will not appear in the
[`activeCues`{#playing-the-media-resource:dom-texttrack-activecues}](#dom-texttrack-activecues)
list.

The [*time marches on*]{#time-marches-on .dfn} steps are as follows:

1.  Let `current cues`{.variable} be a list of
    [cues](#text-track-cue){#playing-the-media-resource:text-track-cue-6},
    initialized to contain all the
    [cues](#text-track-cue){#playing-the-media-resource:text-track-cue-7}
    of all the
    [hidden](#text-track-hidden){#playing-the-media-resource:text-track-hidden}
    or
    [showing](#text-track-showing){#playing-the-media-resource:text-track-showing}
    [text tracks](#text-track){#playing-the-media-resource:text-track-6}
    of the [media
    element](#media-element){#playing-the-media-resource:media-element-73}
    (not the
    [disabled](#text-track-disabled){#playing-the-media-resource:text-track-disabled}
    ones) whose [start
    times](#text-track-cue-start-time){#playing-the-media-resource:text-track-cue-start-time-2}
    are less than or equal to the [current playback
    position](#current-playback-position){#playing-the-media-resource:current-playback-position-15}
    and whose [end
    times](#text-track-cue-end-time){#playing-the-media-resource:text-track-cue-end-time-2}
    are greater than the [current playback
    position](#current-playback-position){#playing-the-media-resource:current-playback-position-16}.

2.  Let `other cues`{.variable} be a list of
    [cues](#text-track-cue){#playing-the-media-resource:text-track-cue-8},
    initialized to contain all the
    [cues](#text-track-cue){#playing-the-media-resource:text-track-cue-9}
    of
    [hidden](#text-track-hidden){#playing-the-media-resource:text-track-hidden-2}
    and
    [showing](#text-track-showing){#playing-the-media-resource:text-track-showing-2}
    [text tracks](#text-track){#playing-the-media-resource:text-track-7}
    of the [media
    element](#media-element){#playing-the-media-resource:media-element-74}
    that are not present in `current cues`{.variable}.

3.  Let `last time`{.variable} be the [current playback
    position](#current-playback-position){#playing-the-media-resource:current-playback-position-17}
    at the time this algorithm was last run for this [media
    element](#media-element){#playing-the-media-resource:media-element-75},
    if this is not the first time it has run.

4.  If the [current playback
    position](#current-playback-position){#playing-the-media-resource:current-playback-position-18}
    has, since the last time this algorithm was run, only changed
    through its usual monotonic increase during normal playback, then
    let `missed cues`{.variable} be the list of
    [cues](#text-track-cue){#playing-the-media-resource:text-track-cue-10}
    in `other cues`{.variable} whose [start
    times](#text-track-cue-start-time){#playing-the-media-resource:text-track-cue-start-time-3}
    are greater than or equal to `last time`{.variable} and whose [end
    times](#text-track-cue-end-time){#playing-the-media-resource:text-track-cue-end-time-3}
    are less than or equal to the [current playback
    position](#current-playback-position){#playing-the-media-resource:current-playback-position-19}.
    Otherwise, let `missed cues`{.variable} be an empty list.

5.  Remove all the
    [cues](#text-track-cue){#playing-the-media-resource:text-track-cue-11}
    in `missed cues`{.variable} that are also in the [media
    element](#media-element){#playing-the-media-resource:media-element-76}\'s
    [list of newly introduced
    cues](#list-of-newly-introduced-cues){#playing-the-media-resource:list-of-newly-introduced-cues-4},
    and then empty the element\'s [list of newly introduced
    cues](#list-of-newly-introduced-cues){#playing-the-media-resource:list-of-newly-introduced-cues-5}.

6.  If the time was reached through the usual monotonic increase of the
    [current playback
    position](#current-playback-position){#playing-the-media-resource:current-playback-position-20}
    during normal playback, and if the user agent has not fired a
    [`timeupdate`{#playing-the-media-resource:event-media-timeupdate-6}](#event-media-timeupdate)
    event at the element in the past 15 to 250ms and is not still
    running event handlers for such an event, then the user agent must
    [queue a media element
    task](#queue-a-media-element-task){#playing-the-media-resource:queue-a-media-element-task-11}
    given the [media
    element](#media-element){#playing-the-media-resource:media-element-77}
    to [fire an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#playing-the-media-resource:concept-event-fire-13
    x-internal="concept-event-fire"} named
    [`timeupdate`{#playing-the-media-resource:event-media-timeupdate-7}](#event-media-timeupdate)
    at the element. (In the other cases, such as explicit seeks,
    relevant events get fired as part of the overall process of changing
    the [current playback
    position](#current-playback-position){#playing-the-media-resource:current-playback-position-21}.)

    The event thus is not to be fired faster than about 66Hz or slower
    than 4Hz (assuming the event handlers don\'t take longer than 250ms
    to run). User agents are encouraged to vary the frequency of the
    event based on the system load and the average cost of processing
    the event each time, so that the UI updates are not any more
    frequent than the user agent can comfortably handle while decoding
    the video.

7.  If all of the
    [cues](#text-track-cue){#playing-the-media-resource:text-track-cue-12}
    in `current cues`{.variable} have their [text track cue active
    flag](#text-track-cue-active-flag){#playing-the-media-resource:text-track-cue-active-flag}
    set, none of the
    [cues](#text-track-cue){#playing-the-media-resource:text-track-cue-13}
    in `other cues`{.variable} have their [text track cue active
    flag](#text-track-cue-active-flag){#playing-the-media-resource:text-track-cue-active-flag-2}
    set, and `missed cues`{.variable} is empty, then return.

8.  If the time was reached through the usual monotonic increase of the
    [current playback
    position](#current-playback-position){#playing-the-media-resource:current-playback-position-22}
    during normal playback, and there are
    [cues](#text-track-cue){#playing-the-media-resource:text-track-cue-14}
    in `other cues`{.variable} that have their [text track cue
    pause-on-exit
    flag](#text-track-cue-pause-on-exit-flag){#playing-the-media-resource:text-track-cue-pause-on-exit-flag}
    set and that either have their [text track cue active
    flag](#text-track-cue-active-flag){#playing-the-media-resource:text-track-cue-active-flag-3}
    set or are also in `missed cues`{.variable}, then
    [immediately](infrastructure.html#immediately){#playing-the-media-resource:immediately}
    [pause](#dom-media-pause){#playing-the-media-resource:dom-media-pause}
    the [media
    element](#media-element){#playing-the-media-resource:media-element-78}.

    In the other cases, such as explicit seeks, playback is not paused
    by going past the end time of a
    [cue](#text-track-cue){#playing-the-media-resource:text-track-cue-15},
    even if that
    [cue](#text-track-cue){#playing-the-media-resource:text-track-cue-16}
    has its [text track cue pause-on-exit
    flag](#text-track-cue-pause-on-exit-flag){#playing-the-media-resource:text-track-cue-pause-on-exit-flag-2}
    set.

9.  Let `events`{.variable} be a list of
    [tasks](webappapis.html#concept-task){#playing-the-media-resource:concept-task},
    initially empty. Each
    [task](webappapis.html#concept-task){#playing-the-media-resource:concept-task-2}
    in this list will be associated with a [text
    track](#text-track){#playing-the-media-resource:text-track-8}, a
    [text track
    cue](#text-track-cue){#playing-the-media-resource:text-track-cue-17},
    and a time, which are used to sort the list before the
    [tasks](webappapis.html#concept-task){#playing-the-media-resource:concept-task-3}
    are queued.

    Let `affected tracks`{.variable} be a list of [text
    tracks](#text-track){#playing-the-media-resource:text-track-9},
    initially empty.

    When the steps below say to [prepare an event]{#prepare-an-event
    .dfn} named `event`{.variable} for a [text track
    cue](#text-track-cue){#playing-the-media-resource:text-track-cue-18}
    `target`{.variable} with a time `time`{.variable}, the user agent
    must run these steps:

    1.  Let `track`{.variable} be the [text
        track](#text-track){#playing-the-media-resource:text-track-10}
        with which the [text track
        cue](#text-track-cue){#playing-the-media-resource:text-track-cue-19}
        `target`{.variable} is associated.

    2.  Create a
        [task](webappapis.html#concept-task){#playing-the-media-resource:concept-task-4}
        to [fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#playing-the-media-resource:concept-event-fire-14
        x-internal="concept-event-fire"} named `event`{.variable} at
        `target`{.variable}.

    3.  Add the newly created
        [task](webappapis.html#concept-task){#playing-the-media-resource:concept-task-5}
        to `events`{.variable}, associated with the time
        `time`{.variable}, the [text
        track](#text-track){#playing-the-media-resource:text-track-11}
        `track`{.variable}, and the [text track
        cue](#text-track-cue){#playing-the-media-resource:text-track-cue-20}
        `target`{.variable}.

    4.  Add `track`{.variable} to `affected tracks`{.variable}.

10. For each [text track
    cue](#text-track-cue){#playing-the-media-resource:text-track-cue-21}
    in `missed cues`{.variable}, [prepare an
    event](#prepare-an-event){#playing-the-media-resource:prepare-an-event}
    named
    [`enter`{#playing-the-media-resource:event-media-enter}](#event-media-enter)
    for the
    [`TextTrackCue`{#playing-the-media-resource:texttrackcue}](#texttrackcue)
    object with the [text track cue start
    time](#text-track-cue-start-time){#playing-the-media-resource:text-track-cue-start-time-4}.

11. For each [text track
    cue](#text-track-cue){#playing-the-media-resource:text-track-cue-22}
    in `other cues`{.variable} that either has its [text track cue
    active
    flag](#text-track-cue-active-flag){#playing-the-media-resource:text-track-cue-active-flag-4}
    set or is in `missed cues`{.variable}, [prepare an
    event](#prepare-an-event){#playing-the-media-resource:prepare-an-event-2}
    named
    [`exit`{#playing-the-media-resource:event-media-exit}](#event-media-exit)
    for the
    [`TextTrackCue`{#playing-the-media-resource:texttrackcue-2}](#texttrackcue)
    object with the later of the [text track cue end
    time](#text-track-cue-end-time){#playing-the-media-resource:text-track-cue-end-time-4}
    and the [text track cue start
    time](#text-track-cue-start-time){#playing-the-media-resource:text-track-cue-start-time-5}.

12. For each [text track
    cue](#text-track-cue){#playing-the-media-resource:text-track-cue-23}
    in `current cues`{.variable} that does not have its [text track cue
    active
    flag](#text-track-cue-active-flag){#playing-the-media-resource:text-track-cue-active-flag-5}
    set, [prepare an
    event](#prepare-an-event){#playing-the-media-resource:prepare-an-event-3}
    named
    [`enter`{#playing-the-media-resource:event-media-enter-2}](#event-media-enter)
    for the
    [`TextTrackCue`{#playing-the-media-resource:texttrackcue-3}](#texttrackcue)
    object with the [text track cue start
    time](#text-track-cue-start-time){#playing-the-media-resource:text-track-cue-start-time-6}.

13. Sort the
    [tasks](webappapis.html#concept-task){#playing-the-media-resource:concept-task-6}
    in `events`{.variable} in ascending time order
    ([tasks](webappapis.html#concept-task){#playing-the-media-resource:concept-task-7}
    with earlier times first).

    Further sort
    [tasks](webappapis.html#concept-task){#playing-the-media-resource:concept-task-8}
    in `events`{.variable} that have the same time by the relative [text
    track cue
    order](#text-track-cue-order){#playing-the-media-resource:text-track-cue-order}
    of the [text track
    cues](#text-track-cue){#playing-the-media-resource:text-track-cue-24}
    associated with these
    [tasks](webappapis.html#concept-task){#playing-the-media-resource:concept-task-9}.

    Finally, sort
    [tasks](webappapis.html#concept-task){#playing-the-media-resource:concept-task-10}
    in `events`{.variable} that have the same time and same [text track
    cue
    order](#text-track-cue-order){#playing-the-media-resource:text-track-cue-order-2}
    by placing
    [tasks](webappapis.html#concept-task){#playing-the-media-resource:concept-task-11}
    that fire
    [`enter`{#playing-the-media-resource:event-media-enter-3}](#event-media-enter)
    events before those that fire
    [`exit`{#playing-the-media-resource:event-media-exit-2}](#event-media-exit)
    events.

14. [Queue a media element
    task](#queue-a-media-element-task){#playing-the-media-resource:queue-a-media-element-task-12}
    given the [media
    element](#media-element){#playing-the-media-resource:media-element-79}
    for each
    [task](webappapis.html#concept-task){#playing-the-media-resource:concept-task-12}
    in `events`{.variable}, in list order.

15. Sort `affected tracks`{.variable} in the same order as the [text
    tracks](#text-track){#playing-the-media-resource:text-track-12}
    appear in the [media
    element](#media-element){#playing-the-media-resource:media-element-80}\'s
    [list of text
    tracks](#list-of-text-tracks){#playing-the-media-resource:list-of-text-tracks-5},
    and remove duplicates.

16. For each [text
    track](#text-track){#playing-the-media-resource:text-track-13} in
    `affected tracks`{.variable}, in the list order, [queue a media
    element
    task](#queue-a-media-element-task){#playing-the-media-resource:queue-a-media-element-task-13}
    given the [media
    element](#media-element){#playing-the-media-resource:media-element-81}
    to [fire an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#playing-the-media-resource:concept-event-fire-15
    x-internal="concept-event-fire"} named
    [`cuechange`{#playing-the-media-resource:event-media-cuechange}](#event-media-cuechange)
    at the
    [`TextTrack`{#playing-the-media-resource:texttrack}](#texttrack)
    object, and, if the [text
    track](#text-track){#playing-the-media-resource:text-track-14} has a
    corresponding
    [`track`{#playing-the-media-resource:the-track-element}](#the-track-element)
    element, to then [fire an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#playing-the-media-resource:concept-event-fire-16
    x-internal="concept-event-fire"} named
    [`cuechange`{#playing-the-media-resource:event-media-cuechange-2}](#event-media-cuechange)
    at the
    [`track`{#playing-the-media-resource:the-track-element-2}](#the-track-element)
    element as well.

17. Set the [text track cue active
    flag](#text-track-cue-active-flag){#playing-the-media-resource:text-track-cue-active-flag-6}
    of all the
    [cues](#text-track-cue){#playing-the-media-resource:text-track-cue-25}
    in the `current cues`{.variable}, and unset the [text track cue
    active
    flag](#text-track-cue-active-flag){#playing-the-media-resource:text-track-cue-active-flag-7}
    of all the
    [cues](#text-track-cue){#playing-the-media-resource:text-track-cue-26}
    in the `other cues`{.variable}.

18. Run the [rules for updating the text track
    rendering](#rules-for-updating-the-text-track-rendering){#playing-the-media-resource:rules-for-updating-the-text-track-rendering}
    of each of the [text
    tracks](#text-track){#playing-the-media-resource:text-track-15} in
    `affected tracks`{.variable} that are
    [showing](#text-track-showing){#playing-the-media-resource:text-track-showing-3},
    providing the [text
    track](#text-track){#playing-the-media-resource:text-track-16}\'s
    [text track
    language](#text-track-language){#playing-the-media-resource:text-track-language}
    as the fallback language if it is not the empty string. For example,
    for [text
    tracks](#text-track){#playing-the-media-resource:text-track-17}
    based on WebVTT, the [rules for updating the display of WebVTT text
    tracks](https://w3c.github.io/webvtt/#rules-for-updating-the-display-of-webvtt-text-tracks){#playing-the-media-resource:rules-for-updating-the-display-of-webvtt-text-tracks
    x-internal="rules-for-updating-the-display-of-webvtt-text-tracks"}.
    [\[WEBVTT\]](references.html#refsWEBVTT)

For the purposes of the algorithm above, a [text track
cue](#text-track-cue){#playing-the-media-resource:text-track-cue-27} is
considered to be part of a [text
track](#text-track){#playing-the-media-resource:text-track-18} only if
it is listed in the [text track list of
cues](#text-track-list-of-cues){#playing-the-media-resource:text-track-list-of-cues-4},
not merely if it is associated with the [text
track](#text-track){#playing-the-media-resource:text-track-19}.

If the [media
element](#media-element){#playing-the-media-resource:media-element-82}\'s
[node
document](https://dom.spec.whatwg.org/#concept-node-document){#playing-the-media-resource:node-document-3
x-internal="node-document"} stops being a [fully
active](document-sequences.html#fully-active){#playing-the-media-resource:fully-active-2}
document, then the playback will [stop](#media-playback) until the
document is active again.

When a [media
element](#media-element){#playing-the-media-resource:media-element-83}
is [removed from a
`Document`](infrastructure.html#remove-an-element-from-a-document){#playing-the-media-resource:remove-an-element-from-a-document},
the user agent must run the following steps:

1.  [Await a stable
    state](webappapis.html#await-a-stable-state){#playing-the-media-resource:await-a-stable-state-2},
    allowing the
    [task](webappapis.html#concept-task){#playing-the-media-resource:concept-task-13}
    that removed the [media
    element](#media-element){#playing-the-media-resource:media-element-84}
    from the
    [`Document`{#playing-the-media-resource:document-3}](dom.html#document)
    to continue. The [synchronous
    section](webappapis.html#synchronous-section){#playing-the-media-resource:synchronous-section}
    consists of all the remaining steps of this algorithm. (Steps in the
    [synchronous
    section](webappapis.html#synchronous-section){#playing-the-media-resource:synchronous-section-2}
    are marked with ⌛.)

2.  ⌛ If the [media
    element](#media-element){#playing-the-media-resource:media-element-85}
    is [in a
    document](https://dom.spec.whatwg.org/#in-a-document){#playing-the-media-resource:in-a-document-2
    x-internal="in-a-document"}, return.

3.  ⌛ Run the [internal pause
    steps](#internal-pause-steps){#playing-the-media-resource:internal-pause-steps-2}
    for the [media
    element](#media-element){#playing-the-media-resource:media-element-86}.

##### [4.8.11.9]{.secno} Seeking[](#seeking){.self-link}

`media`{.variable}`.`[`seeking`](#dom-media-seeking){#dom-media-seeking-dev}

Returns true if the user agent is currently seeking.

`media`{.variable}`.`[`seekable`](#dom-media-seekable){#dom-media-seekable-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/seekable](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/seekable "The seekable read-only property of HTMLMediaElement objects returns a new static normalized TimeRanges object that represents the ranges of the media resource, if any, that the user agent is able to seek to at the time seekable property is accessed.")

Support in all current engines.

::: support
[Firefox8+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns a [`TimeRanges`{#seeking:timeranges}](#timeranges) object that
represents the ranges of the [media
resource](#media-resource){#seeking:media-resource} to which it is
possible for the user agent to seek.

`media`{.variable}`.`[`fastSeek`](#dom-media-fastseek){#dom-media-fastseek-dev}`(``time`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
MDN

:::: feature
[HTMLMediaElement/fastSeek](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/fastSeek "The HTMLMediaElement.fastSeek() method quickly seeks the media to the new time with precision tradeoff.")

::: support
[Firefox31+]{.firefox .yes}[Safari8+]{.safari .yes}[ChromeNo]{.chrome
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

Seeks to near the given `time`{.variable} as fast as possible, trading
precision for speed. (To seek to a precise time, use the
[`currentTime`{#seeking:dom-media-currenttime}](#dom-media-currenttime)
attribute.)

This does nothing if the media resource has not been loaded.

The [`seeking`]{#dom-media-seeking .dfn dfn-for="HTMLMediaElement"
dfn-type="attribute"} attribute must initially have the value false.

The [`fastSeek(``time`{.variable}`)`]{#dom-media-fastseek .dfn
dfn-for="HTMLMediaElement" dfn-type="method"} method must
[seek](#dom-media-seek){#seeking:dom-media-seek} to the time given by
`time`{.variable}, with the *approximate-for-speed* flag set.

When the user agent is required to [seek]{#dom-media-seek .dfn} to a
particular `new playback position`{.variable} in the [media
resource](#media-resource){#seeking:media-resource-2}, optionally with
the *approximate-for-speed* flag set, it means that the user agent must
run the following steps. This algorithm interacts closely with the
[event loop](webappapis.html#event-loop){#seeking:event-loop} mechanism;
in particular, it has a [synchronous
section](webappapis.html#synchronous-section){#seeking:synchronous-section}
(which is triggered as part of the [event
loop](webappapis.html#event-loop){#seeking:event-loop-2} algorithm).
Steps in that section are marked with ⌛.

1.  Set the [media element](#media-element){#seeking:media-element}\'s
    [show poster flag](#show-poster-flag){#seeking:show-poster-flag} to
    false.

2.  If the [media element](#media-element){#seeking:media-element-2}\'s
    [`readyState`{#seeking:dom-media-readystate}](#dom-media-readystate)
    is
    [`HAVE_NOTHING`{#seeking:dom-media-have_nothing}](#dom-media-have_nothing),
    return.

3.  If the element\'s
    [`seeking`{#seeking:dom-media-seeking}](#dom-media-seeking) IDL
    attribute is true, then another instance of this algorithm is
    already running. Abort that other instance of the algorithm without
    waiting for the step that it is running to complete.

4.  Set the
    [`seeking`{#seeking:dom-media-seeking-2}](#dom-media-seeking) IDL
    attribute to true.

5.  If the seek was in response to a DOM method call or setting of an
    IDL attribute, then continue the script. The remainder of these
    steps must be run [in
    parallel](infrastructure.html#in-parallel){#seeking:in-parallel}.
    With the exception of the steps marked with ⌛, they could be
    aborted at any time by another instance of this algorithm being
    invoked.

6.  If the `new playback position`{.variable} is later than the end of
    the [media resource](#media-resource){#seeking:media-resource-3},
    then let it be the end of the [media
    resource](#media-resource){#seeking:media-resource-4} instead.

7.  If the `new playback position`{.variable} is less than the [earliest
    possible
    position](#earliest-possible-position){#seeking:earliest-possible-position},
    let it be that position instead.

8.  If the (possibly now changed) `new playback position`{.variable} is
    not in one of the ranges given in the
    [`seekable`{#seeking:dom-media-seekable}](#dom-media-seekable)
    attribute, then let it be the position in one of the ranges given in
    the [`seekable`{#seeking:dom-media-seekable-2}](#dom-media-seekable)
    attribute that is the nearest to the
    `new playback position`{.variable}. If two positions both satisfy
    that constraint (i.e. the `new playback position`{.variable} is
    exactly in the middle between two ranges in the
    [`seekable`{#seeking:dom-media-seekable-3}](#dom-media-seekable)
    attribute), then use the position that is closest to the [current
    playback
    position](#current-playback-position){#seeking:current-playback-position}.
    If there are no ranges given in the
    [`seekable`{#seeking:dom-media-seekable-4}](#dom-media-seekable)
    attribute, then set the
    [`seeking`{#seeking:dom-media-seeking-3}](#dom-media-seeking) IDL
    attribute to false and return.

9.  If the *approximate-for-speed* flag is set, adjust the
    `new playback position`{.variable} to a value that will allow for
    playback to resume promptly. If `new playback position`{.variable}
    before this step is before [current playback
    position](#current-playback-position){#seeking:current-playback-position-2},
    then the adjusted `new playback position`{.variable} must also be
    before the [current playback
    position](#current-playback-position){#seeking:current-playback-position-3}.
    Similarly, if the `new playback position`{.variable} before this
    step is after [current playback
    position](#current-playback-position){#seeking:current-playback-position-4},
    then the adjusted `new playback position`{.variable} must also be
    after the [current playback
    position](#current-playback-position){#seeking:current-playback-position-5}.

    For example, the user agent could snap to a nearby key frame, so
    that it doesn\'t have to spend time decoding then discarding
    intermediate frames before resuming playback.

10. [Queue a media element
    task](#queue-a-media-element-task){#seeking:queue-a-media-element-task}
    given the [media element](#media-element){#seeking:media-element-3}
    to [fire an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#seeking:concept-event-fire
    x-internal="concept-event-fire"} named
    [`seeking`{#seeking:event-media-seeking}](#event-media-seeking) at
    the element.

11. Set the [current playback
    position](#current-playback-position){#seeking:current-playback-position-6}
    to the `new playback position`{.variable}.

    If the [media element](#media-element){#seeking:media-element-4} was
    [potentially
    playing](#potentially-playing){#seeking:potentially-playing}
    immediately before it started seeking, but seeking caused its
    [`readyState`{#seeking:dom-media-readystate-2}](#dom-media-readystate)
    attribute to change to a value lower than
    [`HAVE_FUTURE_DATA`{#seeking:dom-media-have_future_data}](#dom-media-have_future_data),
    then a
    [`waiting`{#seeking:event-media-waiting}](#event-media-waiting)
    event [will be fired](#fire-waiting-when-waiting) at the element.

    This step sets the [current playback
    position](#current-playback-position){#seeking:current-playback-position-7},
    and thus can immediately trigger other conditions, such as the rules
    regarding when playback \"[reaches the end of the media
    resource](#reaches-the-end)\" (part of the logic that handles
    looping), even before the user agent is actually able to render the
    media data for that position (as determined in the next step).

    The
    [`currentTime`{#seeking:dom-media-currenttime-2}](#dom-media-currenttime)
    attribute returns the [official playback
    position](#official-playback-position){#seeking:official-playback-position},
    not the [current playback
    position](#current-playback-position){#seeking:current-playback-position-8},
    and therefore gets updated before script execution, separate from
    this algorithm.

12. Wait until the user agent has established whether or not the [media
    data](#media-data){#seeking:media-data} for the
    `new playback position`{.variable} is available, and, if it is,
    until it has decoded enough data to play back that position.

13. [Await a stable
    state](webappapis.html#await-a-stable-state){#seeking:await-a-stable-state}.
    The [synchronous
    section](webappapis.html#synchronous-section){#seeking:synchronous-section-2}
    consists of all the remaining steps of this algorithm. (Steps in the
    [synchronous
    section](webappapis.html#synchronous-section){#seeking:synchronous-section-3}
    are marked with ⌛.)

14. ⌛ Set the
    [`seeking`{#seeking:dom-media-seeking-4}](#dom-media-seeking) IDL
    attribute to false.

15. ⌛ Run the [time marches
    on](#time-marches-on){#seeking:time-marches-on} steps.

16. [⌛ [Queue a media element
    task](#queue-a-media-element-task){#seeking:queue-a-media-element-task-2}
    given the [media element](#media-element){#seeking:media-element-5}
    to [fire an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#seeking:concept-event-fire-2
    x-internal="concept-event-fire"} named
    [`timeupdate`{#seeking:event-media-timeupdate}](#event-media-timeupdate)
    at the element.]{#seekUpdate}

17. ⌛ [Queue a media element
    task](#queue-a-media-element-task){#seeking:queue-a-media-element-task-3}
    given the [media element](#media-element){#seeking:media-element-6}
    to [fire an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#seeking:concept-event-fire-3
    x-internal="concept-event-fire"} named
    [`seeked`{#seeking:event-media-seeked}](#event-media-seeked) at the
    element.

------------------------------------------------------------------------

The [`seekable`]{#dom-media-seekable .dfn dfn-for="HTMLMediaElement"
dfn-type="attribute"} attribute must return a new static [normalized
`TimeRanges`
object](#normalised-timeranges-object){#seeking:normalised-timeranges-object}
that represents the ranges of the [media
resource](#media-resource){#seeking:media-resource-5}, if any, that the
user agent is able to seek to, at the time the attribute is evaluated.

If the user agent can seek to anywhere in the [media
resource](#media-resource){#seeking:media-resource-6}, e.g. because it
is a simple movie file and the user agent and the server support HTTP
Range requests, then the attribute would return an object with one
range, whose start is the time of the first frame (the [earliest
possible
position](#earliest-possible-position){#seeking:earliest-possible-position-2},
typically zero), and whose end is the same as the time of the first
frame plus the
[`duration`{#seeking:dom-media-duration}](#dom-media-duration)
attribute\'s value (which would equal the time of the last frame, and
might be positive Infinity).

The range might be continuously changing, e.g. if the user agent is
buffering a sliding window on an infinite stream. This is the behavior
seen with DVRs viewing live TV, for instance.

Returning a new object each time is a bad pattern for attribute getters
and is only enshrined here as it would be costly to change it. It is not
to be copied to new APIs.

User agents should adopt a very liberal and optimistic view of what is
seekable. User agents should also buffer recent content where possible
to enable seeking to be fast.

For instance, consider a large video file served on an HTTP server
without support for HTTP Range requests. A browser *could* implement
this by only buffering the current frame and data obtained for
subsequent frames, never allow seeking, except for seeking to the very
start by restarting the playback. However, this would be a poor
implementation. A high quality implementation would buffer the last few
minutes of content (or more, if sufficient storage space is available),
allowing the user to jump back and rewatch something surprising without
any latency, and would in addition allow arbitrary seeking by reloading
the file from the start if necessary, which would be slower but still
more convenient than having to literally restart the video and watch it
all the way through just to get to an earlier unbuffered spot.

[Media resources](#media-resource){#seeking:media-resource-7} might be
internally scripted or interactive. Thus, a [media
element](#media-element){#seeking:media-element-7} could play in a
non-linear fashion. If this happens, the user agent must act as if the
algorithm for [seeking](#dom-media-seek){#seeking:dom-media-seek-2} was
used whenever the [current playback
position](#current-playback-position){#seeking:current-playback-position-9}
changes in a discontinuous fashion (so that the relevant events fire).

##### [4.8.11.10]{.secno} Media resources with multiple media tracks[](#media-resources-with-multiple-media-tracks){.self-link}

A [media
resource](#media-resource){#media-resources-with-multiple-media-tracks:media-resource}
can have multiple embedded audio and video tracks. For example, in
addition to the primary video and audio tracks, a [media
resource](#media-resource){#media-resources-with-multiple-media-tracks:media-resource-2}
could have foreign-language dubbed dialogues, director\'s commentaries,
audio descriptions, alternative angles, or sign-language overlays.

`media`{.variable}`.`[`audioTracks`](#dom-media-audiotracks){#dom-media-audiotracks-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/audioTracks](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/audioTracks "The read-only audioTracks property on HTMLMediaElement objects returns an AudioTrackList object listing all of the AudioTrack objects representing the media element's audio tracks.")

Support in all current engines.

::: support
[Firefox[🔰
33+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari7+]{.safari .yes}[Chrome[🔰
37+]{title="Requires setting a user preference or runtime flag."}]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge[🔰
79+]{title="Requires setting a user preference or runtime flag."}]{.edge_blink
.yes}

------------------------------------------------------------------------

[Edge (Legacy)No]{.edge .no}[Internet Explorer10+]{.ie .yes}

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
[`AudioTrackList`{#media-resources-with-multiple-media-tracks:audiotracklist}](#audiotracklist)
object representing the audio tracks available in the [media
resource](#media-resource){#media-resources-with-multiple-media-tracks:media-resource-3}.

`media`{.variable}`.`[`videoTracks`](#dom-media-videotracks){#dom-media-videotracks-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/videoTracks](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/videoTracks "The read-only videoTracks property on HTMLMediaElement objects returns a VideoTrackList object listing all of the VideoTrack objects representing the media element's video tracks.")

Support in all current engines.

::: support
[Firefox[🔰
33+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari7+]{.safari .yes}[Chrome[🔰
37+]{title="Requires setting a user preference or runtime flag."}]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge[🔰
79+]{title="Requires setting a user preference or runtime flag."}]{.edge_blink
.yes}

------------------------------------------------------------------------

[Edge (Legacy)No]{.edge .no}[Internet ExplorerNo]{.ie .no}

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
[`VideoTrackList`{#media-resources-with-multiple-media-tracks:videotracklist}](#videotracklist)
object representing the video tracks available in the [media
resource](#media-resource){#media-resources-with-multiple-media-tracks:media-resource-4}.

The [`audioTracks`]{#dom-media-audiotracks .dfn
dfn-for="HTMLMediaElement" dfn-type="attribute"} attribute of a [media
element](#media-element){#media-resources-with-multiple-media-tracks:media-element}
must return a
[live](infrastructure.html#live){#media-resources-with-multiple-media-tracks:live}
[`AudioTrackList`{#media-resources-with-multiple-media-tracks:audiotracklist-2}](#audiotracklist)
object representing the audio tracks available in the [media
element](#media-element){#media-resources-with-multiple-media-tracks:media-element-2}\'s
[media
resource](#media-resource){#media-resources-with-multiple-media-tracks:media-resource-5}.

The [`videoTracks`]{#dom-media-videotracks .dfn
dfn-for="HTMLMediaElement" dfn-type="attribute"} attribute of a [media
element](#media-element){#media-resources-with-multiple-media-tracks:media-element-3}
must return a
[live](infrastructure.html#live){#media-resources-with-multiple-media-tracks:live-2}
[`VideoTrackList`{#media-resources-with-multiple-media-tracks:videotracklist-2}](#videotracklist)
object representing the video tracks available in the [media
element](#media-element){#media-resources-with-multiple-media-tracks:media-element-4}\'s
[media
resource](#media-resource){#media-resources-with-multiple-media-tracks:media-resource-6}.

There are only ever one
[`AudioTrackList`{#media-resources-with-multiple-media-tracks:audiotracklist-3}](#audiotracklist)
object and one
[`VideoTrackList`{#media-resources-with-multiple-media-tracks:videotracklist-3}](#videotracklist)
object per [media
element](#media-element){#media-resources-with-multiple-media-tracks:media-element-5},
even if another [media
resource](#media-resource){#media-resources-with-multiple-media-tracks:media-resource-7}
is loaded into the element: the objects are reused. (The
[`AudioTrack`{#media-resources-with-multiple-media-tracks:audiotrack}](#audiotrack)
and
[`VideoTrack`{#media-resources-with-multiple-media-tracks:videotrack}](#videotrack)
objects are not, though.)

###### [4.8.11.10.1]{.secno} [`AudioTrackList`{#audiotracklist-and-videotracklist-objects:audiotracklist}](#audiotracklist) and [`VideoTrackList`{#audiotracklist-and-videotracklist-objects:videotracklist}](#videotracklist) objects[](#audiotracklist-and-videotracklist-objects){.self-link}

::::::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[AudioTrackList](https://developer.mozilla.org/en-US/docs/Web/API/AudioTrackList "The AudioTrackList interface is used to represent a list of the audio tracks contained within a given HTML media element, with each track represented by a separate AudioTrack object in the list.")

Support in all current engines.

::: support
[Firefox[🔰
33+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari7+]{.safari .yes}[Chrome[🔰
37+]{title="Requires setting a user preference or runtime flag."}]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge[🔰
79+]{title="Requires setting a user preference or runtime flag."}]{.edge_blink
.yes}

------------------------------------------------------------------------

[Edge (Legacy)No]{.edge .no}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::

:::: feature
[VideoTrackList](https://developer.mozilla.org/en-US/docs/Web/API/VideoTrackList "The VideoTrackList interface is used to represent a list of the video tracks contained within a <video> element, with each track represented by a separate VideoTrack object in the list.")

Support in all current engines.

::: support
[Firefox[🔰
33+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari7+]{.safari .yes}[Chrome[🔰
37+]{title="Requires setting a user preference or runtime flag."}]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge[🔰
79+]{title="Requires setting a user preference or runtime flag."}]{.edge_blink
.yes}

------------------------------------------------------------------------

[Edge (Legacy)No]{.edge .no}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::

:::: feature
[AudioTrackList](https://developer.mozilla.org/en-US/docs/Web/API/AudioTrackList "The AudioTrackList interface is used to represent a list of the audio tracks contained within a given HTML media element, with each track represented by a separate AudioTrack object in the list.")

Support in all current engines.

::: support
[Firefox[🔰
33+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari7+]{.safari .yes}[Chrome[🔰
37+]{title="Requires setting a user preference or runtime flag."}]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge[🔰
79+]{title="Requires setting a user preference or runtime flag."}]{.edge_blink
.yes}

------------------------------------------------------------------------

[Edge (Legacy)No]{.edge .no}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}

------------------------------------------------------------------------

[[caniuse.com
table](https://caniuse.com/#feat=audiotracks "Audio Tracks")]{.caniuse}
:::
::::
:::::::::

The
[`AudioTrackList`{#audiotracklist-and-videotracklist-objects:audiotracklist-2}](#audiotracklist)
and
[`VideoTrackList`{#audiotracklist-and-videotracklist-objects:videotracklist-2}](#videotracklist)
interfaces are used by attributes defined in the previous section.

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[AudioTrack](https://developer.mozilla.org/en-US/docs/Web/API/AudioTrack "The AudioTrack interface represents a single audio track from one of the HTML media elements, <audio> or <video>.")

Support in all current engines.

::: support
[Firefox[🔰
33+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari8+]{.safari .yes}[Chrome[🔰
37+]{title="Requires setting a user preference or runtime flag."}]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge[🔰
79+]{title="Requires setting a user preference or runtime flag."}]{.edge_blink
.yes}

------------------------------------------------------------------------

[Edge (Legacy)No]{.edge .no}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::

:::: feature
[VideoTrack](https://developer.mozilla.org/en-US/docs/Web/API/VideoTrack "The VideoTrack interface represents a single video track from a <video> element.")

Support in all current engines.

::: support
[Firefox[🔰
33+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari7+]{.safari .yes}[Chrome[🔰
37+]{title="Requires setting a user preference or runtime flag."}]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge[🔰
79+]{title="Requires setting a user preference or runtime flag."}]{.edge_blink
.yes}

------------------------------------------------------------------------

[Edge (Legacy)No]{.edge .no}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::::

``` idl
[Exposed=Window]
interface AudioTrackList : EventTarget {
  readonly attribute unsigned long length;
  getter AudioTrack (unsigned long index);
  AudioTrack? getTrackById(DOMString id);

  attribute EventHandler onchange;
  attribute EventHandler onaddtrack;
  attribute EventHandler onremovetrack;
};

[Exposed=Window]
interface AudioTrack {
  readonly attribute DOMString id;
  readonly attribute DOMString kind;
  readonly attribute DOMString label;
  readonly attribute DOMString language;
  attribute boolean enabled;
};

[Exposed=Window]
interface VideoTrackList : EventTarget {
  readonly attribute unsigned long length;
  getter VideoTrack (unsigned long index);
  VideoTrack? getTrackById(DOMString id);
  readonly attribute long selectedIndex;

  attribute EventHandler onchange;
  attribute EventHandler onaddtrack;
  attribute EventHandler onremovetrack;
};

[Exposed=Window]
interface VideoTrack {
  readonly attribute DOMString id;
  readonly attribute DOMString kind;
  readonly attribute DOMString label;
  readonly attribute DOMString language;
  attribute boolean selected;
};
```

`media`{.variable}`.`[`audioTracks`](#dom-media-audiotracks){#audiotracklist-and-videotracklist-objects:dom-media-audiotracks}`.`[`length`](#dom-audiotracklist-length){#dom-audiotracklist-length-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[AudioTrackList/length](https://developer.mozilla.org/en-US/docs/Web/API/AudioTrackList/length "The read-only AudioTrackList property length returns the number of entries in the AudioTrackList, each of which is an AudioTrack representing one audio track in the media element. A value of 0 indicates that there are no audio tracks in the media.")

Support in all current engines.

::: support
[Firefox[🔰
33+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari7+]{.safari .yes}[Chrome[🔰
37+]{title="Requires setting a user preference or runtime flag."}]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge[🔰
79+]{title="Requires setting a user preference or runtime flag."}]{.edge_blink
.yes}

------------------------------------------------------------------------

[Edge (Legacy)No]{.edge .no}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

`media`{.variable}`.`[`videoTracks`](#dom-media-videotracks){#audiotracklist-and-videotracklist-objects:dom-media-videotracks}`.`[`length`](#dom-videotracklist-length){#dom-videotracklist-length-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[VideoTrackList/length](https://developer.mozilla.org/en-US/docs/Web/API/VideoTrackList/length "The read-only VideoTrackList property length returns the number of entries in the VideoTrackList, each of which is a VideoTrack representing one video track in the media element.")

Support in all current engines.

::: support
[Firefox[🔰
33+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari7+]{.safari .yes}[Chrome[🔰
37+]{title="Requires setting a user preference or runtime flag."}]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge[🔰
79+]{title="Requires setting a user preference or runtime flag."}]{.edge_blink
.yes}

------------------------------------------------------------------------

[Edge (Legacy)No]{.edge .no}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

Returns the number of tracks in the list.

`audioTrack`{.variable}` = ``media`{.variable}`.`[`audioTracks`](#dom-media-audiotracks){#audiotracklist-and-videotracklist-objects:dom-media-audiotracks-2}`[``index`{.variable}`]`

`videoTrack`{.variable}` = ``media`{.variable}`.`[`videoTracks`](#dom-media-videotracks){#audiotracklist-and-videotracklist-objects:dom-media-videotracks-2}`[``index`{.variable}`]`

Returns the specified
[`AudioTrack`{#audiotracklist-and-videotracklist-objects:audiotrack-3}](#audiotrack)
or
[`VideoTrack`{#audiotracklist-and-videotracklist-objects:videotrack-3}](#videotrack)
object.

`audioTrack`{.variable}` = ``media`{.variable}`.`[`audioTracks`](#dom-media-audiotracks){#audiotracklist-and-videotracklist-objects:dom-media-audiotracks-3}`.`[`getTrackById`](#dom-audiotracklist-gettrackbyid){#dom-audiotracklist-gettrackbyid-dev}`(``id`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[AudioTrackList/getTrackById](https://developer.mozilla.org/en-US/docs/Web/API/AudioTrackList/getTrackById "The AudioTrackList method getTrackById() returns the first AudioTrack object from the track list whose id matches the specified string. This lets you find a specified track if you know its ID string.")

Support in all current engines.

::: support
[Firefox[🔰
33+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari7+]{.safari .yes}[Chrome[🔰
37+]{title="Requires setting a user preference or runtime flag."}]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge[🔰
79+]{title="Requires setting a user preference or runtime flag."}]{.edge_blink
.yes}

------------------------------------------------------------------------

[Edge (Legacy)No]{.edge .no}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

`videoTrack`{.variable}` = ``media`{.variable}`.`[`videoTracks`](#dom-media-videotracks){#audiotracklist-and-videotracklist-objects:dom-media-videotracks-3}`.`[`getTrackById`](#dom-videotracklist-gettrackbyid){#dom-videotracklist-gettrackbyid-dev}`(``id`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[VideoTrackList/getTrackById](https://developer.mozilla.org/en-US/docs/Web/API/VideoTrackList/getTrackById "The VideoTrackList method getTrackById() returns the first VideoTrack object from the track list whose id matches the specified string.")

Support in all current engines.

::: support
[Firefox[🔰
33+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari7+]{.safari .yes}[Chrome[🔰
37+]{title="Requires setting a user preference or runtime flag."}]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge[🔰
79+]{title="Requires setting a user preference or runtime flag."}]{.edge_blink
.yes}

------------------------------------------------------------------------

[Edge (Legacy)No]{.edge .no}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

Returns the
[`AudioTrack`{#audiotracklist-and-videotracklist-objects:audiotrack-4}](#audiotrack)
or
[`VideoTrack`{#audiotracklist-and-videotracklist-objects:videotrack-4}](#videotrack)
object with the given identifier, or null if no track has that
identifier.

`audioTrack`{.variable}`.`[`id`](#dom-audiotrack-id){#dom-audiotrack-id-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[AudioTrack/id](https://developer.mozilla.org/en-US/docs/Web/API/AudioTrack/id "The id property contains a string which uniquely identifies the track represented by the AudioTrack.")

Support in all current engines.

::: support
[Firefox[🔰
33+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari8+]{.safari .yes}[Chrome[🔰
37+]{title="Requires setting a user preference or runtime flag."}]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge[🔰
79+]{title="Requires setting a user preference or runtime flag."}]{.edge_blink
.yes}

------------------------------------------------------------------------

[Edge (Legacy)No]{.edge .no}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

`videoTrack`{.variable}`.`[`id`](#dom-videotrack-id){#dom-videotrack-id-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[VideoTrack/id](https://developer.mozilla.org/en-US/docs/Web/API/VideoTrack/id "The id property contains a string which uniquely identifies the track represented by the VideoTrack.")

Support in all current engines.

::: support
[Firefox[🔰
33+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari7+]{.safari .yes}[Chrome[🔰
37+]{title="Requires setting a user preference or runtime flag."}]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge[🔰
79+]{title="Requires setting a user preference or runtime flag."}]{.edge_blink
.yes}

------------------------------------------------------------------------

[Edge (Legacy)No]{.edge .no}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

Returns the ID of the given track. This is the ID that can be used with
a
[fragment](https://url.spec.whatwg.org/#concept-url-fragment){#audiotracklist-and-videotracklist-objects:concept-url-fragment
x-internal="concept-url-fragment"} if the format supports [media
fragment
syntax](https://www.w3.org/TR/media-frags/#media-fragment-syntax){#audiotracklist-and-videotracklist-objects:media-fragment-syntax
x-internal="media-fragment-syntax"}, and that can be used with the
`getTrackById()` method.

`audioTrack`{.variable}`.`[`kind`](#dom-audiotrack-kind){#dom-audiotrack-kind-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[AudioTrack/kind](https://developer.mozilla.org/en-US/docs/Web/API/AudioTrack/kind "The kind property contains a string indicating the category of audio contained in the AudioTrack.")

Support in all current engines.

::: support
[Firefox[🔰
33+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari8+]{.safari .yes}[Chrome[🔰
37+]{title="Requires setting a user preference or runtime flag."}]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge[🔰
79+]{title="Requires setting a user preference or runtime flag."}]{.edge_blink
.yes}

------------------------------------------------------------------------

[Edge (Legacy)No]{.edge .no}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

`videoTrack`{.variable}`.`[`kind`](#dom-videotrack-kind){#dom-videotrack-kind-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[VideoTrack/kind](https://developer.mozilla.org/en-US/docs/Web/API/VideoTrack/kind "The kind property contains a string indicating the category of video contained in the VideoTrack.")

Support in all current engines.

::: support
[Firefox[🔰
33+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari7+]{.safari .yes}[Chrome[🔰
37+]{title="Requires setting a user preference or runtime flag."}]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge[🔰
79+]{title="Requires setting a user preference or runtime flag."}]{.edge_blink
.yes}

------------------------------------------------------------------------

[Edge (Legacy)No]{.edge .no}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

Returns the category the given track falls into. The [possible track
categories](#dom-TrackList-getKind-categories) are given below.

`audioTrack`{.variable}`.`[`label`](#dom-audiotrack-label){#dom-audiotrack-label-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[AudioTrack/label](https://developer.mozilla.org/en-US/docs/Web/API/AudioTrack/label "The read-only AudioTrack property label returns a string specifying the audio track's human-readable label, if one is available; otherwise, it returns an empty string.")

Support in all current engines.

::: support
[Firefox[🔰
33+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari8+]{.safari .yes}[Chrome[🔰
37+]{title="Requires setting a user preference or runtime flag."}]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge[🔰
79+]{title="Requires setting a user preference or runtime flag."}]{.edge_blink
.yes}

------------------------------------------------------------------------

[Edge (Legacy)No]{.edge .no}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

`videoTrack`{.variable}`.`[`label`](#dom-videotrack-label){#dom-videotrack-label-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[VideoTrack/label](https://developer.mozilla.org/en-US/docs/Web/API/VideoTrack/label "The read-only VideoTrack property label returns a string specifying the video track's human-readable label, if one is available; otherwise, it returns an empty string.")

Support in all current engines.

::: support
[Firefox[🔰
33+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari7+]{.safari .yes}[Chrome[🔰
37+]{title="Requires setting a user preference or runtime flag."}]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge[🔰
79+]{title="Requires setting a user preference or runtime flag."}]{.edge_blink
.yes}

------------------------------------------------------------------------

[Edge (Legacy)No]{.edge .no}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

Returns the label of the given track, if known, or the empty string
otherwise.

`audioTrack`{.variable}`.`[`language`](#dom-audiotrack-language){#dom-audiotrack-language-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[AudioTrack/language](https://developer.mozilla.org/en-US/docs/Web/API/AudioTrack/language "The read-only AudioTrack property language returns a string identifying the language used in the audio track.")

Support in all current engines.

::: support
[Firefox[🔰
33+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari8+]{.safari .yes}[Chrome[🔰
37+]{title="Requires setting a user preference or runtime flag."}]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge[🔰
79+]{title="Requires setting a user preference or runtime flag."}]{.edge_blink
.yes}

------------------------------------------------------------------------

[Edge (Legacy)No]{.edge .no}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

`videoTrack`{.variable}`.`[`language`](#dom-videotrack-language){#dom-videotrack-language-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[VideoTrack/language](https://developer.mozilla.org/en-US/docs/Web/API/VideoTrack/language "The read-only VideoTrack property language returns a string identifying the language used in the video track.")

Support in all current engines.

::: support
[Firefox[🔰
33+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari7+]{.safari .yes}[Chrome[🔰
37+]{title="Requires setting a user preference or runtime flag."}]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge[🔰
79+]{title="Requires setting a user preference or runtime flag."}]{.edge_blink
.yes}

------------------------------------------------------------------------

[Edge (Legacy)No]{.edge .no}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

Returns the language of the given track, if known, or the empty string
otherwise.

`audioTrack`{.variable}`.`[`enabled`](#dom-audiotrack-enabled){#dom-audiotrack-enabled-dev}` [ = ``value`{.variable}` ]`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[AudioTrack/enabled](https://developer.mozilla.org/en-US/docs/Web/API/AudioTrack/enabled "The AudioTrack property enabled specifies whether or not the described audio track is currently enabled for use. If the track is disabled by setting enabled to false, the track is muted and does not produce audio.")

Support in all current engines.

::: support
[Firefox[🔰
33+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari8+]{.safari .yes}[Chrome[🔰
37+]{title="Requires setting a user preference or runtime flag."}]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge[🔰
79+]{title="Requires setting a user preference or runtime flag."}]{.edge_blink
.yes}

------------------------------------------------------------------------

[Edge (Legacy)No]{.edge .no}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

Returns true if the given track is active, and false otherwise.

Can be set, to change whether the track is enabled or not. If multiple
audio tracks are enabled simultaneously, they are mixed.

`media`{.variable}`.`[`videoTracks`](#dom-media-videotracks){#audiotracklist-and-videotracklist-objects:dom-media-videotracks-4}`.`[`selectedIndex`](#dom-videotracklist-selectedindex){#dom-videotracklist-selectedindex-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[VideoTrackList/selectedIndex](https://developer.mozilla.org/en-US/docs/Web/API/VideoTrackList/selectedIndex "The read-only VideoTrackList property selectedIndex returns the index of the currently selected track, if any, or -1 otherwise.")

Support in all current engines.

::: support
[Firefox[🔰
33+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari7+]{.safari .yes}[Chrome[🔰
37+]{title="Requires setting a user preference or runtime flag."}]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge[🔰
79+]{title="Requires setting a user preference or runtime flag."}]{.edge_blink
.yes}

------------------------------------------------------------------------

[Edge (Legacy)No]{.edge .no}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

Returns the index of the currently selected track, if any, or −1
otherwise.

`videoTrack`{.variable}`.`[`selected`](#dom-videotrack-selected){#dom-videotrack-selected-dev}` [ = ``value`{.variable}` ]`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[VideoTrack/selected](https://developer.mozilla.org/en-US/docs/Web/API/VideoTrack/selected "The VideoTrack property selected controls whether or not a particular video track is active.")

Support in all current engines.

::: support
[Firefox[🔰
33+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari7+]{.safari .yes}[Chrome[🔰
37+]{title="Requires setting a user preference or runtime flag."}]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge[🔰
79+]{title="Requires setting a user preference or runtime flag."}]{.edge_blink
.yes}

------------------------------------------------------------------------

[Edge (Legacy)No]{.edge .no}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

Returns true if the given track is active, and false otherwise.

Can be set, to change whether the track is selected or not. Either zero
or one video track is selected; selecting a new track while a previous
one is selected will unselect the previous one.

An
[`AudioTrackList`{#audiotracklist-and-videotracklist-objects:audiotracklist-3}](#audiotracklist)
object represents a dynamic list of zero or more audio tracks, of which
zero or more can be enabled at a time. Each audio track is represented
by an
[`AudioTrack`{#audiotracklist-and-videotracklist-objects:audiotrack-5}](#audiotrack)
object.

A
[`VideoTrackList`{#audiotracklist-and-videotracklist-objects:videotracklist-3}](#videotracklist)
object represents a dynamic list of zero or more video tracks, of which
zero or one can be selected at a time. Each video track is represented
by a
[`VideoTrack`{#audiotracklist-and-videotracklist-objects:videotrack-5}](#videotrack)
object.

Tracks in
[`AudioTrackList`{#audiotracklist-and-videotracklist-objects:audiotracklist-4}](#audiotracklist)
and
[`VideoTrackList`{#audiotracklist-and-videotracklist-objects:videotracklist-4}](#videotracklist)
objects must be consistently ordered. If the [media
resource](#media-resource){#audiotracklist-and-videotracklist-objects:media-resource}
is in a format that defines an order, then that order must be used;
otherwise, the order must be the relative order in which the tracks are
declared in the [media
resource](#media-resource){#audiotracklist-and-videotracklist-objects:media-resource-2}.
The order used is called the *natural order* of the list.

Each track in one of these objects thus has an index; the first has the
index 0, and each subsequent track is numbered one higher than the
previous one. If a [media
resource](#media-resource){#audiotracklist-and-videotracklist-objects:media-resource-3}
dynamically adds or removes audio or video tracks, then the indices of
the tracks will change dynamically. If the [media
resource](#media-resource){#audiotracklist-and-videotracklist-objects:media-resource-4}
changes entirely, then all the previous tracks will be removed and
replaced with new tracks.

The
[`AudioTrackList`{#audiotracklist-and-videotracklist-objects:audiotracklist-5}](#audiotracklist)
[`length`]{#dom-audiotracklist-length .dfn dfn-for="AudioTrackList"
dfn-type="attribute"} and
[`VideoTrackList`{#audiotracklist-and-videotracklist-objects:videotracklist-5}](#videotracklist)
[`length`]{#dom-videotracklist-length .dfn dfn-for="VideoTrackList"
dfn-type="attribute"} attribute getters must return the number of tracks
represented by their objects at the time of getting.

The [supported property
indices](https://webidl.spec.whatwg.org/#dfn-supported-property-indices){#audiotracklist-and-videotracklist-objects:supported-property-indices
x-internal="supported-property-indices"} of
[`AudioTrackList`{#audiotracklist-and-videotracklist-objects:audiotracklist-6}](#audiotracklist)
and
[`VideoTrackList`{#audiotracklist-and-videotracklist-objects:videotracklist-6}](#videotracklist)
objects at any instant are the numbers from zero to the number of tracks
represented by the respective object minus one, if any tracks are
represented. If an
[`AudioTrackList`{#audiotracklist-and-videotracklist-objects:audiotracklist-7}](#audiotracklist)
or
[`VideoTrackList`{#audiotracklist-and-videotracklist-objects:videotracklist-7}](#videotracklist)
object represents no tracks, it has no [supported property
indices](https://webidl.spec.whatwg.org/#dfn-supported-property-indices){#audiotracklist-and-videotracklist-objects:supported-property-indices-2
x-internal="supported-property-indices"}.

To [determine the value of an indexed
property](https://webidl.spec.whatwg.org/#dfn-determine-the-value-of-an-indexed-property){#audiotracklist-and-videotracklist-objects:determine-the-value-of-an-indexed-property
x-internal="determine-the-value-of-an-indexed-property"} for a given
index `index`{.variable} in an
[`AudioTrackList`{#audiotracklist-and-videotracklist-objects:audiotracklist-8}](#audiotracklist)
or
[`VideoTrackList`{#audiotracklist-and-videotracklist-objects:videotracklist-8}](#videotracklist)
object `list`{.variable}, the user agent must return the
[`AudioTrack`{#audiotracklist-and-videotracklist-objects:audiotrack-6}](#audiotrack)
or
[`VideoTrack`{#audiotracklist-and-videotracklist-objects:videotrack-6}](#videotrack)
object that represents the `index`{.variable}th track in
`list`{.variable}.

The
[`AudioTrackList`{#audiotracklist-and-videotracklist-objects:audiotracklist-9}](#audiotracklist)
[`getTrackById(``id`{.variable}`)`]{#dom-audiotracklist-gettrackbyid
.dfn dfn-for="AudioTrackList" dfn-type="method"} and
[`VideoTrackList`{#audiotracklist-and-videotracklist-objects:videotracklist-9}](#videotracklist)
[`getTrackById(``id`{.variable}`)`]{#dom-videotracklist-gettrackbyid
.dfn dfn-for="VideoTrackList" dfn-type="method"} methods must return the
first
[`AudioTrack`{#audiotracklist-and-videotracklist-objects:audiotrack-7}](#audiotrack)
or
[`VideoTrack`{#audiotracklist-and-videotracklist-objects:videotrack-7}](#videotrack)
object (respectively) in the
[`AudioTrackList`{#audiotracklist-and-videotracklist-objects:audiotracklist-10}](#audiotracklist)
or
[`VideoTrackList`{#audiotracklist-and-videotracklist-objects:videotracklist-10}](#videotracklist)
object (respectively) whose identifier is equal to the value of the
`id`{.variable} argument (in the natural order of the list, as defined
above). When no tracks match the given argument, the methods must return
null.

The
[`AudioTrack`{#audiotracklist-and-videotracklist-objects:audiotrack-8}](#audiotrack)
and
[`VideoTrack`{#audiotracklist-and-videotracklist-objects:videotrack-8}](#videotrack)
objects represent specific tracks of a [media
resource](#media-resource){#audiotracklist-and-videotracklist-objects:media-resource-5}.
Each track can have an identifier, category, label, and language. These
aspects of a track are permanent for the lifetime of the track; even if
a track is removed from a [media
resource](#media-resource){#audiotracklist-and-videotracklist-objects:media-resource-6}\'s
[`AudioTrackList`{#audiotracklist-and-videotracklist-objects:audiotracklist-11}](#audiotracklist)
or
[`VideoTrackList`{#audiotracklist-and-videotracklist-objects:videotracklist-11}](#videotracklist)
objects, those aspects do not change.

In addition,
[`AudioTrack`{#audiotracklist-and-videotracklist-objects:audiotrack-9}](#audiotrack)
objects can each be enabled or disabled; this is the audio track\'s
*enabled state*. When an
[`AudioTrack`{#audiotracklist-and-videotracklist-objects:audiotrack-10}](#audiotrack)
is created, its *enabled state* must be set to false (disabled). The
[resource fetch
algorithm](#concept-media-load-resource){#audiotracklist-and-videotracklist-objects:concept-media-load-resource}
can override this.

Similarly, a single
[`VideoTrack`{#audiotracklist-and-videotracklist-objects:videotrack-9}](#videotrack)
object per
[`VideoTrackList`{#audiotracklist-and-videotracklist-objects:videotracklist-12}](#videotracklist)
object can be selected, this is the video track\'s *selection state*.
When a
[`VideoTrack`{#audiotracklist-and-videotracklist-objects:videotrack-10}](#videotrack)
is created, its *selection state* must be set to false (not selected).
The [resource fetch
algorithm](#concept-media-load-resource){#audiotracklist-and-videotracklist-objects:concept-media-load-resource-2}
can override this.

The
[`AudioTrack`{#audiotracklist-and-videotracklist-objects:audiotrack-11}](#audiotrack)
[`id`]{#dom-audiotrack-id .dfn dfn-for="AudioTrack"
dfn-type="attribute"} and
[`VideoTrack`{#audiotracklist-and-videotracklist-objects:videotrack-11}](#videotrack)
[`id`]{#dom-videotrack-id .dfn dfn-for="VideoTrack"
dfn-type="attribute"} attributes must return the identifier of the
track, if it has one, or the empty string otherwise. If the [media
resource](#media-resource){#audiotracklist-and-videotracklist-objects:media-resource-7}
is in a format that supports [media fragment
syntax](https://www.w3.org/TR/media-frags/#media-fragment-syntax){#audiotracklist-and-videotracklist-objects:media-fragment-syntax-2
x-internal="media-fragment-syntax"}, the identifier returned for a
particular track must be the same identifier that would enable the track
if used as the name of a track in the track dimension of such a
[fragment](https://url.spec.whatwg.org/#concept-url-fragment){#audiotracklist-and-videotracklist-objects:concept-url-fragment-2
x-internal="concept-url-fragment"}.
[\[INBAND\]](references.html#refsINBAND)

For example, in Ogg files, this would be the Name header field of the
track. [\[OGGSKELETONHEADERS\]](references.html#refsOGGSKELETONHEADERS)

The
[`AudioTrack`{#audiotracklist-and-videotracklist-objects:audiotrack-12}](#audiotrack)
[`kind`]{#dom-audiotrack-kind .dfn dfn-for="AudioTrack"
dfn-type="attribute"} and
[`VideoTrack`{#audiotracklist-and-videotracklist-objects:videotrack-12}](#videotrack)
[`kind`]{#dom-videotrack-kind .dfn dfn-for="VideoTrack"
dfn-type="attribute"} attributes must return the category of the track,
if it has one, or the empty string otherwise.

The category of a track is the string given in the first column of the
table below that is the most appropriate for the track based on the
definitions in the table\'s second and third columns, as determined by
the metadata included in the track in the [media
resource](#media-resource){#audiotracklist-and-videotracklist-objects:media-resource-8}.
The cell in the third column of a row says what the category given in
the cell in the first column of that row applies to; a category is only
appropriate for an audio track if it applies to audio tracks, and a
category is only appropriate for video tracks if it applies to video
tracks. Categories must only be returned for
[`AudioTrack`{#audiotracklist-and-videotracklist-objects:audiotrack-13}](#audiotrack)
objects if they are appropriate for audio, and must only be returned for
[`VideoTrack`{#audiotracklist-and-videotracklist-objects:videotrack-13}](#videotrack)
objects if they are appropriate for video.

For Ogg files, the Role header field of the track gives the relevant
metadata. For DASH media resources, the `Role` element conveys the
information. For WebM, only the `FlagDefault` element currently maps to
a value. Sourcing In-band Media Resource Tracks from Media Containers
into HTML has further details.
[\[OGGSKELETONHEADERS\]](references.html#refsOGGSKELETONHEADERS)
[\[DASH\]](references.html#refsDASH)
[\[WEBMCG\]](references.html#refsWEBMCG)
[\[INBAND\]](references.html#refsINBAND)

Return values for
[`AudioTrack`{#audiotracklist-and-videotracklist-objects:audiotrack-14}](#audiotrack)\'s
[`kind`{#audiotracklist-and-videotracklist-objects:dom-audiotrack-kind-2}](#dom-audiotrack-kind)
and
[`VideoTrack`{#audiotracklist-and-videotracklist-objects:videotrack-14}](#videotrack)\'s
[`kind`{#audiotracklist-and-videotracklist-objects:dom-videotrack-kind-2}](#dom-videotrack-kind)

Category

Definition

Applies to\...

Examples

\"[`alternative`]{#value-track-kind-alternate .dfn}\"

A possible alternative to the main track, e.g. a different take of a
song (audio), or a different angle (video).

Audio and video.

Ogg: \"audio/alternate\" or \"video/alternate\"; DASH: \"alternate\"
without \"main\" and \"commentary\" roles, and, for audio, without the
\"dub\" role (other roles ignored).

\"[`captions`]{#value-track-kind-caption .dfn}\"

A version of the main video track with captions burnt in. (For legacy
content; new content would use text tracks.)

Video only.

DASH: \"caption\" and \"main\" roles together (other roles ignored).

\"[`descriptions`]{#value-track-kind-descriptions .dfn}\"

An audio description of a video track.

Audio only.

Ogg: \"audio/audiodesc\".

\"[`main`]{#value-track-kind-main .dfn}\"

The primary audio or video track.

Audio and video.

Ogg: \"audio/main\" or \"video/main\"; WebM: the \"FlagDefault\" element
is set; DASH: \"main\" role without \"caption\", \"subtitle\", and
\"dub\" roles (other roles ignored).

\"[`main-desc`]{#value-track-kind-main-desc .dfn}\"

The primary audio track, mixed with audio descriptions.

Audio only.

AC3 audio in MPEG-2 TS: bsmod=2 and full_svc=1.

\"[`sign`]{#value-track-kind-sign .dfn}\"

A sign-language interpretation of an audio track.

Video only.

Ogg: \"video/sign\".

\"[`subtitles`]{#value-track-kind-subtitle .dfn}\"

A version of the main video track with subtitles burnt in. (For legacy
content; new content would use text tracks.)

Video only.

DASH: \"subtitle\" and \"main\" roles together (other roles ignored).

\"[`translation`]{#value-track-kind-translation .dfn}\"

A translated version of the main audio track.

Audio only.

Ogg: \"audio/dub\". DASH: \"dub\" and \"main\" roles together (other
roles ignored).

\"[`commentary`]{#value-track-kind-commentary .dfn}\"

Commentary on the primary audio or video track, e.g. a director\'s
commentary.

Audio and video.

DASH: \"commentary\" role without \"main\" role (other roles ignored).

\"[]{#value-track-kind-none .dfn}\" (empty string)

No explicit kind, or the kind given by the track\'s metadata is not
recognized by the user agent.

Audio and video.

The
[`AudioTrack`{#audiotracklist-and-videotracklist-objects:audiotrack-15}](#audiotrack)
[`label`]{#dom-audiotrack-label .dfn dfn-for="AudioTrack"
dfn-type="attribute"} and
[`VideoTrack`{#audiotracklist-and-videotracklist-objects:videotrack-15}](#videotrack)
[`label`]{#dom-videotrack-label .dfn dfn-for="VideoTrack"
dfn-type="attribute"} attributes must return the label of the track, if
it has one, or the empty string otherwise.
[\[INBAND\]](references.html#refsINBAND)

The
[`AudioTrack`{#audiotracklist-and-videotracklist-objects:audiotrack-16}](#audiotrack)
[`language`]{#dom-audiotrack-language .dfn dfn-for="AudioTrack"
dfn-type="attribute"} and
[`VideoTrack`{#audiotracklist-and-videotracklist-objects:videotrack-16}](#videotrack)
[`language`]{#dom-videotrack-language .dfn dfn-for="VideoTrack"
dfn-type="attribute"} attributes must return the BCP 47 language tag of
the language of the track, if it has one, or the empty string otherwise.
If the user agent is not able to express that language as a BCP 47
language tag (for example because the language information in the [media
resource](#media-resource){#audiotracklist-and-videotracklist-objects:media-resource-9}\'s
format is a free-form string without a defined interpretation), then the
method must return the empty string, as if the track had no language.
[\[INBAND\]](references.html#refsINBAND)

The
[`AudioTrack`{#audiotracklist-and-videotracklist-objects:audiotrack-17}](#audiotrack)
[`enabled`]{#dom-audiotrack-enabled .dfn dfn-for="AudioTrack"
dfn-type="attribute"} attribute, on getting, must return true if the
track is currently enabled, and false otherwise. On setting, it must
enable the track if the new value is true, and disable it otherwise. (If
the track is no longer in an
[`AudioTrackList`{#audiotracklist-and-videotracklist-objects:audiotracklist-12}](#audiotracklist)
object, then the track being enabled or disabled has no effect beyond
changing the value of the attribute on the
[`AudioTrack`{#audiotracklist-and-videotracklist-objects:audiotrack-18}](#audiotrack)
object.)

Whenever an audio track in an
[`AudioTrackList`{#audiotracklist-and-videotracklist-objects:audiotracklist-13}](#audiotracklist)
that was disabled is enabled, and whenever one that was enabled is
disabled, the user agent must [queue a media element
task](#queue-a-media-element-task){#audiotracklist-and-videotracklist-objects:queue-a-media-element-task}
given the [media
element](#media-element){#audiotracklist-and-videotracklist-objects:media-element}
to [fire an
event](https://dom.spec.whatwg.org/#concept-event-fire){#audiotracklist-and-videotracklist-objects:concept-event-fire
x-internal="concept-event-fire"} named
[`change`{#audiotracklist-and-videotracklist-objects:event-media-change}](#event-media-change)
at the
[`AudioTrackList`{#audiotracklist-and-videotracklist-objects:audiotracklist-14}](#audiotracklist)
object.

An audio track that has no data for a particular position on the [media
timeline](#media-timeline){#audiotracklist-and-videotracklist-objects:media-timeline},
or that does not exist at that position, must be interpreted as being
silent at that point on the timeline.

The
[`VideoTrackList`{#audiotracklist-and-videotracklist-objects:videotracklist-13}](#videotracklist)
[`selectedIndex`]{#dom-videotracklist-selectedindex .dfn
dfn-for="VideoTrackList" dfn-type="attribute"} attribute must return the
index of the currently selected track, if any. If the
[`VideoTrackList`{#audiotracklist-and-videotracklist-objects:videotracklist-14}](#videotracklist)
object does not currently represent any tracks, or if none of the tracks
are selected, it must instead return −1.

The
[`VideoTrack`{#audiotracklist-and-videotracklist-objects:videotrack-17}](#videotrack)
[`selected`]{#dom-videotrack-selected .dfn dfn-for="VideoTrack"
dfn-type="attribute"} attribute, on getting, must return true if the
track is currently selected, and false otherwise. On setting, it must
select the track if the new value is true, and unselect it otherwise. If
the track is in a
[`VideoTrackList`{#audiotracklist-and-videotracklist-objects:videotracklist-15}](#videotracklist),
then all the other
[`VideoTrack`{#audiotracklist-and-videotracklist-objects:videotrack-18}](#videotrack)
objects in that list must be unselected. (If the track is no longer in a
[`VideoTrackList`{#audiotracklist-and-videotracklist-objects:videotracklist-16}](#videotracklist)
object, then the track being selected or unselected has no effect beyond
changing the value of the attribute on the
[`VideoTrack`{#audiotracklist-and-videotracklist-objects:videotrack-19}](#videotrack)
object.)

Whenever a track in a
[`VideoTrackList`{#audiotracklist-and-videotracklist-objects:videotracklist-17}](#videotracklist)
that was previously not selected is selected, and whenever the selected
track in a
[`VideoTrackList`{#audiotracklist-and-videotracklist-objects:videotracklist-18}](#videotracklist)
is unselected without a new track being selected in its stead, the user
agent must [queue a media element
task](#queue-a-media-element-task){#audiotracklist-and-videotracklist-objects:queue-a-media-element-task-2}
given the [media
element](#media-element){#audiotracklist-and-videotracklist-objects:media-element-2}
to [fire an
event](https://dom.spec.whatwg.org/#concept-event-fire){#audiotracklist-and-videotracklist-objects:concept-event-fire-2
x-internal="concept-event-fire"} named
[`change`{#audiotracklist-and-videotracklist-objects:event-media-change-2}](#event-media-change)
at the
[`VideoTrackList`{#audiotracklist-and-videotracklist-objects:videotracklist-19}](#videotracklist)
object. This
[task](webappapis.html#concept-task){#audiotracklist-and-videotracklist-objects:concept-task}
must be
[queued](webappapis.html#queue-an-element-task){#audiotracklist-and-videotracklist-objects:queue-an-element-task}
before the
[task](webappapis.html#concept-task){#audiotracklist-and-videotracklist-objects:concept-task-2}
that fires the
[`resize`{#audiotracklist-and-videotracklist-objects:event-media-resize}](#event-media-resize)
event, if any.

A video track that has no data for a particular position on the [media
timeline](#media-timeline){#audiotracklist-and-videotracklist-objects:media-timeline-2}
must be interpreted as being [transparent
black](https://drafts.csswg.org/css-color/#transparent-black){#audiotracklist-and-videotracklist-objects:transparent-black
x-internal="transparent-black"} at that point on the timeline, with the
same dimensions as the last frame before that position, or, if the
position is before all the data for that track, the same dimensions as
the first frame for that track. A track that does not exist at all at
the current position must be treated as if it existed but had no data.

For instance, if a video has a track that is only introduced after one
hour of playback, and the user selects that track then goes back to the
start, then the user agent will act as if that track started at the
start of the [media
resource](#media-resource){#audiotracklist-and-videotracklist-objects:media-resource-10}
but was simply transparent until one hour in.

------------------------------------------------------------------------

The following are the [event
handlers](webappapis.html#event-handlers){#audiotracklist-and-videotracklist-objects:event-handlers}
(and their corresponding [event handler event
types](webappapis.html#event-handler-event-type){#audiotracklist-and-videotracklist-objects:event-handler-event-type})
that must be supported, as [event handler IDL
attributes](webappapis.html#event-handler-idl-attributes){#audiotracklist-and-videotracklist-objects:event-handler-idl-attributes},
by all objects implementing the
[`AudioTrackList`{#audiotracklist-and-videotracklist-objects:audiotracklist-15}](#audiotracklist)
and
[`VideoTrackList`{#audiotracklist-and-videotracklist-objects:videotracklist-20}](#videotracklist)
interfaces:

[Event
handler](webappapis.html#event-handlers){#audiotracklist-and-videotracklist-objects:event-handlers-2}

[Event handler event
type](webappapis.html#event-handler-event-type){#audiotracklist-and-videotracklist-objects:event-handler-event-type-2}

[`onchange`]{#handler-tracklist-onchange .dfn
dfn-for="AudioTrackList,VideoTrackList" dfn-type="attribute"}

::::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[AudioTrackList/change_event](https://developer.mozilla.org/en-US/docs/Web/API/AudioTrackList/change_event "The change event is fired when an audio track is enabled or disabled, for example by changing the track's enabled property.")

Support in all current engines.

::: support
[Firefox[🔰
33+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari7+]{.safari .yes}[Chrome[🔰
37+]{title="Requires setting a user preference or runtime flag."}]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge[🔰
79+]{title="Requires setting a user preference or runtime flag."}]{.edge_blink
.yes}

------------------------------------------------------------------------

[Edge (Legacy)No]{.edge .no}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::

:::: feature
[TextTrackList/change_event](https://developer.mozilla.org/en-US/docs/Web/API/TextTrackList/change_event "The change event is fired when a text track is made active or inactive, or a TextTrackList is otherwise changed.")

Support in all current engines.

::: support
[Firefox31+]{.firefox .yes}[Safari7+]{.safari .yes}[Chrome33+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)18]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android4.4+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::

:::: feature
[VideoTrackList/change_event](https://developer.mozilla.org/en-US/docs/Web/API/VideoTrackList/change_event "The change event is fired when a video track is made active or inactive, for example by changing the track's selected property.")

Support in all current engines.

::: support
[Firefox[🔰
33+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari7+]{.safari .yes}[Chrome[🔰
37+]{title="Requires setting a user preference or runtime flag."}]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge[🔰
79+]{title="Requires setting a user preference or runtime flag."}]{.edge_blink
.yes}

------------------------------------------------------------------------

[Edge (Legacy)No]{.edge .no}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::::::

[`change`{#audiotracklist-and-videotracklist-objects:event-media-change-3}](#event-media-change)

[`onaddtrack`]{#handler-tracklist-onaddtrack .dfn
dfn-for="AudioTrackList,VideoTrackList" dfn-type="attribute"}

::::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[AudioTrackList/addtrack_event](https://developer.mozilla.org/en-US/docs/Web/API/AudioTrackList/addtrack_event "The addtrack event is fired when a track is added to an AudioTrackList.")

Support in all current engines.

::: support
[Firefox[🔰
33+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari7+]{.safari .yes}[Chrome[🔰
37+]{title="Requires setting a user preference or runtime flag."}]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge[🔰
79+]{title="Requires setting a user preference or runtime flag."}]{.edge_blink
.yes}

------------------------------------------------------------------------

[Edge (Legacy)No]{.edge .no}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::

:::: feature
[TextTrackList/addtrack_event](https://developer.mozilla.org/en-US/docs/Web/API/TextTrackList/addtrack_event "The addtrack event is fired when a track is added to a TextTrackList.")

Support in all current engines.

::: support
[Firefox31+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome23+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer11]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS7+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::

:::: feature
[VideoTrackList/addtrack_event](https://developer.mozilla.org/en-US/docs/Web/API/VideoTrackList/addtrack_event "The addtrack event is fired when a video track is added to a VideoTrackList.")

Support in all current engines.

::: support
[Firefox[🔰
33+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari7+]{.safari .yes}[Chrome[🔰
37+]{title="Requires setting a user preference or runtime flag."}]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge[🔰
79+]{title="Requires setting a user preference or runtime flag."}]{.edge_blink
.yes}

------------------------------------------------------------------------

[Edge (Legacy)No]{.edge .no}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::::::

[`addtrack`{#audiotracklist-and-videotracklist-objects:event-media-addtrack}](#event-media-addtrack)

[`onremovetrack`]{#handler-tracklist-onremovetrack .dfn
dfn-for="AudioTrackList,VideoTrackList" dfn-type="attribute"}

::::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[AudioTrackList/removetrack_event](https://developer.mozilla.org/en-US/docs/Web/API/AudioTrackList/removetrack_event "The removetrack event is fired when a track is removed from an AudioTrackList.")

Support in all current engines.

::: support
[Firefox[🔰
33+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari7+]{.safari .yes}[Chrome[🔰
37+]{title="Requires setting a user preference or runtime flag."}]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge[🔰
79+]{title="Requires setting a user preference or runtime flag."}]{.edge_blink
.yes}

------------------------------------------------------------------------

[Edge (Legacy)No]{.edge .no}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::

:::: feature
[TextTrackList/removetrack_event](https://developer.mozilla.org/en-US/docs/Web/API/TextTrackList/removetrack_event "The removetrack event is fired when a track is removed from a TextTrackList.")

Support in all current engines.

::: support
[Firefox31+]{.firefox .yes}[Safari7+]{.safari .yes}[Chrome33+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera20+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)18]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android4.4+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android20+]{.opera_android .yes}
:::
::::

:::: feature
[VideoTrackList/removetrack_event](https://developer.mozilla.org/en-US/docs/Web/API/VideoTrackList/removetrack_event "The removetrack event is fired when a video track is removed from a VideoTrackList.")

Support in all current engines.

::: support
[Firefox[🔰
33+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari7+]{.safari .yes}[Chrome[🔰
37+]{title="Requires setting a user preference or runtime flag."}]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge[🔰
79+]{title="Requires setting a user preference or runtime flag."}]{.edge_blink
.yes}

------------------------------------------------------------------------

[Edge (Legacy)No]{.edge .no}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::::::

[`removetrack`{#audiotracklist-and-videotracklist-objects:event-media-removetrack}](#event-media-removetrack)

###### [4.8.11.10.2]{.secno} Selecting specific audio and video tracks declaratively[](#selecting-specific-audio-and-video-tracks-declaratively){.self-link}

The
[`audioTracks`{#selecting-specific-audio-and-video-tracks-declaratively:dom-media-audiotracks}](#dom-media-audiotracks)
and
[`videoTracks`{#selecting-specific-audio-and-video-tracks-declaratively:dom-media-videotracks}](#dom-media-videotracks)
attributes allow scripts to select which track should play, but it is
also possible to select specific tracks declaratively, by specifying
particular tracks in the
[fragment](https://url.spec.whatwg.org/#concept-url-fragment){#selecting-specific-audio-and-video-tracks-declaratively:concept-url-fragment
x-internal="concept-url-fragment"} of the
[URL](https://url.spec.whatwg.org/#concept-url){#selecting-specific-audio-and-video-tracks-declaratively:url
x-internal="url"} of the [media
resource](#media-resource){#selecting-specific-audio-and-video-tracks-declaratively:media-resource}.
The format of the
[fragment](https://url.spec.whatwg.org/#concept-url-fragment){#selecting-specific-audio-and-video-tracks-declaratively:concept-url-fragment-2
x-internal="concept-url-fragment"} depends on the [MIME
type](https://mimesniff.spec.whatwg.org/#mime-type){#selecting-specific-audio-and-video-tracks-declaratively:mime-type
x-internal="mime-type"} of the [media
resource](#media-resource){#selecting-specific-audio-and-video-tracks-declaratively:media-resource-2}.
[\[RFC2046\]](references.html#refsRFC2046)
[\[URL\]](references.html#refsURL)

::: example
In this example, a video that uses a format that supports [media
fragment
syntax](https://www.w3.org/TR/media-frags/#media-fragment-syntax){#selecting-specific-audio-and-video-tracks-declaratively:media-fragment-syntax
x-internal="media-fragment-syntax"} is embedded in such a way that the
alternative angles labeled \"Alternative\" are enabled instead of the
default video track.

``` html
<video src="myvideo#track=Alternative"></video>
```
:::

##### [4.8.11.11]{.secno} Timed text tracks[](#timed-text-tracks){.self-link}

###### [4.8.11.11.1]{.secno} Text track model[](#text-track-model){.self-link}

A [media element](#media-element){#text-track-model:media-element} can
have a group of associated [text tracks]{#text-track .dfn}, known as the
[media element](#media-element){#text-track-model:media-element-2}\'s
[list of text tracks]{#list-of-text-tracks .dfn}. The [text
tracks](#text-track){#text-track-model:text-track} are sorted as
follows:

1.  The [text tracks](#text-track){#text-track-model:text-track-2}
    corresponding to
    [`track`{#text-track-model:the-track-element}](#the-track-element)
    element children of the [media
    element](#media-element){#text-track-model:media-element-3}, in
    [tree
    order](https://dom.spec.whatwg.org/#concept-tree-order){#text-track-model:tree-order
    x-internal="tree-order"}.

2.  Any [text tracks](#text-track){#text-track-model:text-track-3} added
    using the
    [`addTextTrack()`{#text-track-model:dom-media-addtexttrack}](#dom-media-addtexttrack)
    method, in the order they were added, oldest first.

3.  Any [media-resource-specific text
    tracks](#media-resource-specific-text-track){#text-track-model:media-resource-specific-text-track}
    ([text tracks](#text-track){#text-track-model:text-track-4}
    corresponding to data in the [media
    resource](#media-resource){#text-track-model:media-resource}), in
    the order defined by the [media
    resource](#media-resource){#text-track-model:media-resource-2}\'s
    format specification.

A [text track](#text-track){#text-track-model:text-track-5} consists of:

[The kind of text track]{#text-track-kind .dfn}

:   This decides how the track is handled by the user agent. The kind is
    represented by a string. The possible strings are:

    - [`subtitles`]{#dom-texttrack-kind-subtitles .dfn}
    - [`captions`]{#dom-texttrack-kind-captions .dfn}
    - [`descriptions`]{#dom-texttrack-kind-descriptions .dfn}
    - [`chapters`]{#dom-texttrack-kind-chapters .dfn}
    - [`metadata`]{#dom-texttrack-kind-metadata .dfn}

    The [kind of
    track](#text-track-kind){#text-track-model:text-track-kind} can
    change dynamically, in the case of a [text
    track](#text-track){#text-track-model:text-track-6} corresponding to
    a
    [`track`{#text-track-model:the-track-element-2}](#the-track-element)
    element.

[A label]{#text-track-label .dfn}

:   This is a human-readable string intended to identify the track for
    the user.

    The [label of a
    track](#text-track-label){#text-track-model:text-track-label} can
    change dynamically, in the case of a [text
    track](#text-track){#text-track-model:text-track-7} corresponding to
    a
    [`track`{#text-track-model:the-track-element-3}](#the-track-element)
    element.

    When a [text track
    label](#text-track-label){#text-track-model:text-track-label-2} is
    the empty string, the user agent should automatically generate an
    appropriate label from the text track\'s other properties (e.g. the
    kind of text track and the text track\'s language) for use in its
    user interface. This automatically-generated label is not exposed in
    the API.

[An in-band metadata track dispatch type]{#text-track-in-band-metadata-track-dispatch-type .dfn}

:   This is a string extracted from the [media
    resource](#media-resource){#text-track-model:media-resource-3}
    specifically for in-band metadata tracks to enable such tracks to be
    dispatched to different scripts in the document.

    For example, a traditional TV station broadcast streamed on the web
    and augmented with web-specific interactive features could include
    text tracks with metadata for ad targeting, trivia game data during
    game shows, player states during sports games, recipe information
    during food programs, and so forth. As each program starts and ends,
    new tracks might be added or removed from the stream, and as each
    one is added, the user agent could bind them to dedicated script
    modules using the value of this attribute.

    Other than for in-band metadata text tracks, the [in-band metadata
    track dispatch
    type](#text-track-in-band-metadata-track-dispatch-type){#text-track-model:text-track-in-band-metadata-track-dispatch-type}
    is the empty string. How this value is populated for different media
    formats is described in [steps to expose a media-resource-specific
    text
    track](#steps-to-expose-a-media-resource-specific-text-track){#text-track-model:steps-to-expose-a-media-resource-specific-text-track}.

[A language]{#text-track-language .dfn}

:   This is a string (a BCP 47 language tag) representing the language
    of the text track\'s cues. [\[BCP47\]](references.html#refsBCP47)

    The [language of a text
    track](#text-track-language){#text-track-model:text-track-language}
    can change dynamically, in the case of a [text
    track](#text-track){#text-track-model:text-track-8} corresponding to
    a
    [`track`{#text-track-model:the-track-element-4}](#the-track-element)
    element.

[A readiness state]{#text-track-readiness-state .dfn}

:   One of the following:

    [Not loaded]{#text-track-not-loaded .dfn}

    :   Indicates that the text track\'s cues have not been obtained.

    [Loading]{#text-track-loading .dfn}

    :   Indicates that the text track is loading and there have been no
        fatal errors encountered so far. Further cues might still be
        added to the track by the parser.

    [Loaded]{#text-track-loaded .dfn}

    :   Indicates that the text track has been loaded with no fatal
        errors.

    [Failed to load]{#text-track-failed-to-load .dfn}

    :   Indicates that the text track was enabled, but when the user
        agent attempted to obtain it, this failed in some way (e.g.,
        [URL](https://url.spec.whatwg.org/#concept-url){#text-track-model:url
        x-internal="url"} could not be
        [parsed](urls-and-fetching.html#encoding-parsing-a-url){#text-track-model:encoding-parsing-a-url},
        network error, unknown text track format). Some or all of the
        cues are likely missing and will not be obtained.

    The [readiness
    state](#text-track-readiness-state){#text-track-model:text-track-readiness-state}
    of a [text track](#text-track){#text-track-model:text-track-9}
    changes dynamically as the track is obtained.

[A mode]{#text-track-mode .dfn}

:   One of the following:

    [Disabled]{#text-track-disabled .dfn}

    :   Indicates that the text track is not active. Other than for the
        purposes of exposing the track in the DOM, the user agent is
        ignoring the text track. No cues are active, no events are
        fired, and the user agent will not attempt to obtain the
        track\'s cues.

    [Hidden]{#text-track-hidden .dfn}

    :   Indicates that the text track is active, but that the user agent
        is not actively displaying the cues. If no attempt has yet been
        made to obtain the track\'s cues, the user agent will perform
        such an attempt momentarily. The user agent is maintaining a
        list of which cues are active, and events are being fired
        accordingly.

    [Showing]{#text-track-showing .dfn}

    :   Indicates that the text track is active. If no attempt has yet
        been made to obtain the track\'s cues, the user agent will
        perform such an attempt momentarily. The user agent is
        maintaining a list of which cues are active, and events are
        being fired accordingly. In addition, for text tracks whose
        [kind](#text-track-kind){#text-track-model:text-track-kind-2} is
        [`subtitles`{#text-track-model:dom-texttrack-kind-subtitles}](#dom-texttrack-kind-subtitles)
        or
        [`captions`{#text-track-model:dom-texttrack-kind-captions}](#dom-texttrack-kind-captions),
        the cues are being overlaid on the video as appropriate; for
        text tracks whose
        [kind](#text-track-kind){#text-track-model:text-track-kind-3} is
        [`descriptions`{#text-track-model:dom-texttrack-kind-descriptions}](#dom-texttrack-kind-descriptions),
        the user agent is making the cues available to the user in a
        non-visual fashion; and for text tracks whose
        [kind](#text-track-kind){#text-track-model:text-track-kind-4} is
        [`chapters`{#text-track-model:dom-texttrack-kind-chapters}](#dom-texttrack-kind-chapters),
        the user agent is making available to the user a mechanism by
        which the user can navigate to any point in the [media
        resource](#media-resource){#text-track-model:media-resource-4}
        by selecting a cue.

[A list of zero or more cues]{#text-track-list-of-cues .dfn}

:   A list of [text track
    cues](#text-track-cue){#text-track-model:text-track-cue}, along with
    [rules for updating the text track
    rendering]{#rules-for-updating-the-text-track-rendering .dfn}. For
    example, for WebVTT, the [rules for updating the display of WebVTT
    text
    tracks](https://w3c.github.io/webvtt/#rules-for-updating-the-display-of-webvtt-text-tracks){#text-track-model:rules-for-updating-the-display-of-webvtt-text-tracks
    x-internal="rules-for-updating-the-display-of-webvtt-text-tracks"}.
    [\[WEBVTT\]](references.html#refsWEBVTT)

    The [list of cues of a text
    track](#text-track-list-of-cues){#text-track-model:text-track-list-of-cues}
    can change dynamically, either because the [text
    track](#text-track){#text-track-model:text-track-10} has [not yet
    been
    loaded](#text-track-not-loaded){#text-track-model:text-track-not-loaded}
    or is still
    [loading](#text-track-loading){#text-track-model:text-track-loading},
    or due to DOM manipulation.

Each [text track](#text-track){#text-track-model:text-track-11} has a
corresponding [`TextTrack`{#text-track-model:texttrack}](#texttrack)
object.

------------------------------------------------------------------------

Each [media element](#media-element){#text-track-model:media-element-4}
has a [list of pending text tracks]{#list-of-pending-text-tracks .dfn},
which must initially be empty, a [blocked-on-parser]{#blocked-on-parser
.dfn} flag, which must initially be false, and a
[did-perform-automatic-track-selection]{#did-perform-automatic-track-selection
.dfn} flag, which must also initially be false.

When the user agent is required to [populate the list of pending text
tracks]{#populate-the-list-of-pending-text-tracks .dfn} of a [media
element](#media-element){#text-track-model:media-element-5}, the user
agent must add to the element\'s [list of pending text
tracks](#list-of-pending-text-tracks){#text-track-model:list-of-pending-text-tracks}
each [text track](#text-track){#text-track-model:text-track-12} in the
element\'s [list of text
tracks](#list-of-text-tracks){#text-track-model:list-of-text-tracks}
whose [text track
mode](#text-track-mode){#text-track-model:text-track-mode} is not
[disabled](#text-track-disabled){#text-track-model:text-track-disabled}
and whose [text track readiness
state](#text-track-readiness-state){#text-track-model:text-track-readiness-state-2}
is
[loading](#text-track-loading){#text-track-model:text-track-loading-2}.

Whenever a
[`track`{#text-track-model:the-track-element-5}](#the-track-element)
element\'s parent node changes, the user agent must remove the
corresponding [text track](#text-track){#text-track-model:text-track-13}
from any [list of pending text
tracks](#list-of-pending-text-tracks){#text-track-model:list-of-pending-text-tracks-2}
that it is in.

Whenever a [text track](#text-track){#text-track-model:text-track-14}\'s
[text track readiness
state](#text-track-readiness-state){#text-track-model:text-track-readiness-state-3}
changes to either
[loaded](#text-track-loaded){#text-track-model:text-track-loaded} or
[failed to
load](#text-track-failed-to-load){#text-track-model:text-track-failed-to-load},
the user agent must remove it from any [list of pending text
tracks](#list-of-pending-text-tracks){#text-track-model:list-of-pending-text-tracks-3}
that it is in.

When a [media
element](#media-element){#text-track-model:media-element-6} is created
by an [HTML
parser](parsing.html#html-parser){#text-track-model:html-parser} or [XML
parser](xhtml.html#xml-parser){#text-track-model:xml-parser}, the user
agent must set the element\'s
[blocked-on-parser](#blocked-on-parser){#text-track-model:blocked-on-parser}
flag to true. When a [media
element](#media-element){#text-track-model:media-element-7} is popped
off the [stack of open
elements](parsing.html#stack-of-open-elements){#text-track-model:stack-of-open-elements}
of an [HTML
parser](parsing.html#html-parser){#text-track-model:html-parser-2} or
[XML parser](xhtml.html#xml-parser){#text-track-model:xml-parser-2}, the
user agent must [honor user preferences for automatic text track
selection](#honor-user-preferences-for-automatic-text-track-selection){#text-track-model:honor-user-preferences-for-automatic-text-track-selection},
[populate the list of pending text
tracks](#populate-the-list-of-pending-text-tracks){#text-track-model:populate-the-list-of-pending-text-tracks},
and set the element\'s
[blocked-on-parser](#blocked-on-parser){#text-track-model:blocked-on-parser-2}
flag to false.

The [text tracks](#text-track){#text-track-model:text-track-15} of a
[media element](#media-element){#text-track-model:media-element-8} are
[ready]{#the-text-tracks-are-ready .dfn} when both the element\'s [list
of pending text
tracks](#list-of-pending-text-tracks){#text-track-model:list-of-pending-text-tracks-4}
is empty and the element\'s
[blocked-on-parser](#blocked-on-parser){#text-track-model:blocked-on-parser-3}
flag is false.

Each [media element](#media-element){#text-track-model:media-element-9}
has a [pending text track change notification
flag]{#pending-text-track-change-notification-flag .dfn}, which must
initially be unset.

Whenever a [text track](#text-track){#text-track-model:text-track-16}
that is in a [media
element](#media-element){#text-track-model:media-element-10}\'s [list of
text
tracks](#list-of-text-tracks){#text-track-model:list-of-text-tracks-2}
has its [text track
mode](#text-track-mode){#text-track-model:text-track-mode-2} change
value, the user agent must run the following steps for the [media
element](#media-element){#text-track-model:media-element-11}:

1.  If the [media
    element](#media-element){#text-track-model:media-element-12}\'s
    [pending text track change notification
    flag](#pending-text-track-change-notification-flag){#text-track-model:pending-text-track-change-notification-flag}
    is set, return.

2.  Set the [media
    element](#media-element){#text-track-model:media-element-13}\'s
    [pending text track change notification
    flag](#pending-text-track-change-notification-flag){#text-track-model:pending-text-track-change-notification-flag-2}.

3.  [Queue a media element
    task](#queue-a-media-element-task){#text-track-model:queue-a-media-element-task}
    given the [media
    element](#media-element){#text-track-model:media-element-14} to run
    these steps:

    1.  Unset the [media
        element](#media-element){#text-track-model:media-element-15}\'s
        [pending text track change notification
        flag](#pending-text-track-change-notification-flag){#text-track-model:pending-text-track-change-notification-flag-3}.

    2.  [Fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#text-track-model:concept-event-fire
        x-internal="concept-event-fire"} named
        [`change`{#text-track-model:event-media-change}](#event-media-change)
        at the [media
        element](#media-element){#text-track-model:media-element-16}\'s
        [`textTracks`{#text-track-model:dom-media-texttracks}](#dom-media-texttracks)
        attribute\'s
        [`TextTrackList`{#text-track-model:texttracklist}](#texttracklist)
        object.

4.  If the [media
    element](#media-element){#text-track-model:media-element-17}\'s
    [show poster
    flag](#show-poster-flag){#text-track-model:show-poster-flag} is not
    set, run the *[time marches on](#time-marches-on)* steps.

The [task
source](webappapis.html#task-source){#text-track-model:task-source} for
the
[tasks](webappapis.html#concept-task){#text-track-model:concept-task}
listed in this section is the [DOM manipulation task
source](webappapis.html#dom-manipulation-task-source){#text-track-model:dom-manipulation-task-source}.

------------------------------------------------------------------------

A [text track cue]{#text-track-cue .dfn} is the unit of time-sensitive
data in a [text track](#text-track){#text-track-model:text-track-17},
corresponding for instance for subtitles and captions to the text that
appears at a particular time and disappears at another time.

Each [text track
cue](#text-track-cue){#text-track-model:text-track-cue-2} consists of:

[An identifier]{#text-track-cue-identifier .dfn}

:   An arbitrary string.

[A start time]{#text-track-cue-start-time .dfn}

:   The time, in seconds and fractions of a second, that describes the
    beginning of the range of the [media
    data](#media-data){#text-track-model:media-data} to which the cue
    applies.

[An end time]{#text-track-cue-end-time .dfn}

:   The time, in seconds and fractions of a second, that describes the
    end of the range of the [media
    data](#media-data){#text-track-model:media-data-2} to which the cue
    applies, or positive Infinity for an [unbounded text track
    cue](#unbounded-text-track-cue){#text-track-model:unbounded-text-track-cue}.

[A pause-on-exit flag]{#text-track-cue-pause-on-exit-flag .dfn}

:   A boolean indicating whether playback of the [media
    resource](#media-resource){#text-track-model:media-resource-5} is to
    pause when the end of the range to which the cue applies is reached.

Some additional format-specific data

:   Additional fields, as needed for the format, including the actual
    data of the cue. For example, WebVTT has a [text track cue writing
    direction](https://w3c.github.io/webvtt/#webvtt-cue-writing-direction){#text-track-model:text-track-cue-writing-direction
    x-internal="text-track-cue-writing-direction"} and so forth.
    [\[WEBVTT\]](references.html#refsWEBVTT)

An [unbounded text track cue]{#unbounded-text-track-cue .dfn} is a text
track cue with a [text track cue end
time](#text-track-cue-end-time){#text-track-model:text-track-cue-end-time}
set to positive Infinity. An active [unbounded text track
cue](#unbounded-text-track-cue){#text-track-model:unbounded-text-track-cue-2}
cannot become inactive through the usual monotonic increase of the
[current playback
position](#current-playback-position){#text-track-model:current-playback-position}
during normal playback (e.g. a metadata cue for a chapter in a live
event with no announced end time.)

The [text track cue start
time](#text-track-cue-start-time){#text-track-model:text-track-cue-start-time}
and [text track cue end
time](#text-track-cue-end-time){#text-track-model:text-track-cue-end-time-2}
can be negative. (The [current playback
position](#current-playback-position){#text-track-model:current-playback-position-2}
can never be negative, though, so cues entirely before time zero cannot
be active.)

Each [text track
cue](#text-track-cue){#text-track-model:text-track-cue-3} has a
corresponding
[`TextTrackCue`{#text-track-model:texttrackcue}](#texttrackcue) object
(or more specifically, an object that inherits from
[`TextTrackCue`{#text-track-model:texttrackcue-2}](#texttrackcue) ---
for example, WebVTT cues use the
[`VTTCue`{#text-track-model:vttcue}](https://w3c.github.io/webvtt/#vttcue){x-internal="vttcue"}
interface). A [text track
cue](#text-track-cue){#text-track-model:text-track-cue-4}\'s in-memory
representation can be dynamically changed through this
[`TextTrackCue`{#text-track-model:texttrackcue-3}](#texttrackcue) API.
[\[WEBVTT\]](references.html#refsWEBVTT)

A [text track cue](#text-track-cue){#text-track-model:text-track-cue-5}
is associated with [rules for updating the text track
rendering](#rules-for-updating-the-text-track-rendering){#text-track-model:rules-for-updating-the-text-track-rendering},
as defined by the specification for the specific kind of [text track
cue](#text-track-cue){#text-track-model:text-track-cue-6}. These rules
are used specifically when the object representing the cue is added to a
[`TextTrack`{#text-track-model:texttrack-2}](#texttrack) object using
the
[`addCue()`{#text-track-model:dom-texttrack-addcue}](#dom-texttrack-addcue)
method.

In addition, each [text track
cue](#text-track-cue){#text-track-model:text-track-cue-7} has two pieces
of dynamic information:

The [active flag]{#text-track-cue-active-flag .dfn}

:   This flag must be initially unset. The flag is used to ensure events
    are fired appropriately when the cue becomes active or inactive, and
    to make sure the right cues are rendered.

    The user agent must synchronously unset this flag whenever the [text
    track cue](#text-track-cue){#text-track-model:text-track-cue-8} is
    removed from its [text
    track](#text-track){#text-track-model:text-track-18}\'s [text track
    list of
    cues](#text-track-list-of-cues){#text-track-model:text-track-list-of-cues-2};
    whenever the [text
    track](#text-track){#text-track-model:text-track-19} itself is
    removed from its [media
    element](#media-element){#text-track-model:media-element-18}\'s
    [list of text
    tracks](#list-of-text-tracks){#text-track-model:list-of-text-tracks-3}
    or has its [text track
    mode](#text-track-mode){#text-track-model:text-track-mode-3} changed
    to
    [disabled](#text-track-disabled){#text-track-model:text-track-disabled-2};
    and whenever the [media
    element](#media-element){#text-track-model:media-element-19}\'s
    [`readyState`{#text-track-model:dom-media-readystate}](#dom-media-readystate)
    is changed back to
    [`HAVE_NOTHING`{#text-track-model:dom-media-have_nothing}](#dom-media-have_nothing).
    When the flag is unset in this way for one or more cues in [text
    tracks](#text-track){#text-track-model:text-track-20} that were
    [showing](#text-track-showing){#text-track-model:text-track-showing}
    prior to the relevant incident, the user agent must, after having
    unset the flag for all the affected cues, apply the [rules for
    updating the text track
    rendering](#rules-for-updating-the-text-track-rendering){#text-track-model:rules-for-updating-the-text-track-rendering-2}
    of those [text
    tracks](#text-track){#text-track-model:text-track-21}. For example,
    for [text tracks](#text-track){#text-track-model:text-track-22}
    based on WebVTT, the [rules for updating the display of WebVTT text
    tracks](https://w3c.github.io/webvtt/#rules-for-updating-the-display-of-webvtt-text-tracks){#text-track-model:rules-for-updating-the-display-of-webvtt-text-tracks-2
    x-internal="rules-for-updating-the-display-of-webvtt-text-tracks"}.
    [\[WEBVTT\]](references.html#refsWEBVTT)

The [display state]{#text-track-cue-display-state .dfn}

:   This is used as part of the rendering model, to keep cues in a
    consistent position. It must initially be empty. Whenever the [text
    track cue active
    flag](#text-track-cue-active-flag){#text-track-model:text-track-cue-active-flag}
    is unset, the user agent must empty the [text track cue display
    state](#text-track-cue-display-state){#text-track-model:text-track-cue-display-state}.

The [text track
cues](#text-track-cue){#text-track-model:text-track-cue-9} of a [media
element](#media-element){#text-track-model:media-element-20}\'s [text
tracks](#text-track){#text-track-model:text-track-23} are ordered
relative to each other in the [text track cue
order]{#text-track-cue-order .dfn}, which is determined as follows:
first group the
[cues](#text-track-cue){#text-track-model:text-track-cue-10} by their
[text track](#text-track){#text-track-model:text-track-24}, with the
groups being sorted in the same order as their [text
tracks](#text-track){#text-track-model:text-track-25} appear in the
[media element](#media-element){#text-track-model:media-element-21}\'s
[list of text
tracks](#list-of-text-tracks){#text-track-model:list-of-text-tracks-4};
then, within each group,
[cues](#text-track-cue){#text-track-model:text-track-cue-11} must be
sorted by their [start
time](#text-track-cue-start-time){#text-track-model:text-track-cue-start-time-2},
earliest first; then, any
[cues](#text-track-cue){#text-track-model:text-track-cue-12} with the
same [start
time](#text-track-cue-start-time){#text-track-model:text-track-cue-start-time-3}
must be sorted by their [end
time](#text-track-cue-end-time){#text-track-model:text-track-cue-end-time-3},
latest first; and finally, any
[cues](#text-track-cue){#text-track-model:text-track-cue-13} with
identical [end
times](#text-track-cue-end-time){#text-track-model:text-track-cue-end-time-4}
must be sorted in the order they were last added to their respective
[text track list of
cues](#text-track-list-of-cues){#text-track-model:text-track-list-of-cues-3},
oldest first (so e.g. for cues from a WebVTT file, that would initially
be the order in which the cues were listed in the file).
[\[WEBVTT\]](references.html#refsWEBVTT)

###### [4.8.11.11.2]{.secno} Sourcing in-band text tracks[](#sourcing-in-band-text-tracks){.self-link}

A [media-resource-specific text
track]{#media-resource-specific-text-track .dfn} is a [text
track](#text-track){#sourcing-in-band-text-tracks:text-track} that
corresponds to data found in the [media
resource](#media-resource){#sourcing-in-band-text-tracks:media-resource}.

Rules for processing and rendering such data are defined by the relevant
specifications, e.g. the specification of the video format if the [media
resource](#media-resource){#sourcing-in-band-text-tracks:media-resource-2}
is a video. Details for some legacy formats can be found in Sourcing
In-band Media Resource Tracks from Media Containers into HTML.
[\[INBAND\]](references.html#refsINBAND)

When a [media
resource](#media-resource){#sourcing-in-band-text-tracks:media-resource-3}
contains data that the user agent recognizes and supports as being
equivalent to a [text
track](#text-track){#sourcing-in-band-text-tracks:text-track-2}, the
user agent [runs](#found-a-media-resource-specific-timed-track) the
[steps to expose a media-resource-specific text
track]{#steps-to-expose-a-media-resource-specific-text-track .dfn} with
the relevant data, as follows.

1.  Associate the relevant data with a new [text
    track](#text-track){#sourcing-in-band-text-tracks:text-track-3} and
    its corresponding new
    [`TextTrack`{#sourcing-in-band-text-tracks:texttrack}](#texttrack)
    object. The [text
    track](#text-track){#sourcing-in-band-text-tracks:text-track-4} is a
    [media-resource-specific text
    track](#media-resource-specific-text-track){#sourcing-in-band-text-tracks:media-resource-specific-text-track}.

2.  Set the new [text
    track](#text-track){#sourcing-in-band-text-tracks:text-track-5}\'s
    [kind](#text-track-kind){#sourcing-in-band-text-tracks:text-track-kind},
    [label](#text-track-label){#sourcing-in-band-text-tracks:text-track-label},
    and
    [language](#text-track-language){#sourcing-in-band-text-tracks:text-track-language}
    based on the semantics of the relevant data, as defined by the
    relevant specification. If there is no label in that data, then the
    [label](#text-track-label){#sourcing-in-band-text-tracks:text-track-label-2}
    must be set to the empty string.

3.  Associate the [text track list of
    cues](#text-track-list-of-cues){#sourcing-in-band-text-tracks:text-track-list-of-cues}
    with the [rules for updating the text track
    rendering](#rules-for-updating-the-text-track-rendering){#sourcing-in-band-text-tracks:rules-for-updating-the-text-track-rendering}
    appropriate for the format in question.

4.  If the new [text
    track](#text-track){#sourcing-in-band-text-tracks:text-track-6}\'s
    [kind](#text-track-kind){#sourcing-in-band-text-tracks:text-track-kind-2}
    is
    [`chapters`{#sourcing-in-band-text-tracks:dom-texttrack-kind-chapters}](#dom-texttrack-kind-chapters)
    or
    [`metadata`{#sourcing-in-band-text-tracks:dom-texttrack-kind-metadata}](#dom-texttrack-kind-metadata),
    then set the [text track in-band metadata track dispatch
    type](#text-track-in-band-metadata-track-dispatch-type){#sourcing-in-band-text-tracks:text-track-in-band-metadata-track-dispatch-type}
    as follows, based on the type of the [media
    resource](#media-resource){#sourcing-in-band-text-tracks:media-resource-4}:

    If the [media resource](#media-resource){#sourcing-in-band-text-tracks:media-resource-5} is an Ogg file
    :   The [text track in-band metadata track dispatch
        type](#text-track-in-band-metadata-track-dispatch-type){#sourcing-in-band-text-tracks:text-track-in-band-metadata-track-dispatch-type-2}
        must be set to the value of the Name header field.
        [\[OGGSKELETONHEADERS\]](references.html#refsOGGSKELETONHEADERS)

    If the [media resource](#media-resource){#sourcing-in-band-text-tracks:media-resource-6} is a WebM file
    :   The [text track in-band metadata track dispatch
        type](#text-track-in-band-metadata-track-dispatch-type){#sourcing-in-band-text-tracks:text-track-in-band-metadata-track-dispatch-type-3}
        must be set to the value of the `CodecID` element.
        [\[WEBMCG\]](references.html#refsWEBMCG)

    If the [media resource](#media-resource){#sourcing-in-band-text-tracks:media-resource-7} is an MPEG-2 file
    :   Let `stream type`{.variable} be the value of the \"stream_type\"
        field describing the text track\'s type in the file\'s program
        map section, interpreted as an 8-bit unsigned integer. Let
        `length`{.variable} be the value of the \"ES_info_length\" field
        for the track in the same part of the program map section,
        interpreted as an integer as defined by Generic coding of moving
        pictures and associated audio information. Let
        `descriptor bytes`{.variable} be the `length`{.variable} bytes
        following the \"ES_info_length\" field. The [text track in-band
        metadata track dispatch
        type](#text-track-in-band-metadata-track-dispatch-type){#sourcing-in-band-text-tracks:text-track-in-band-metadata-track-dispatch-type-4}
        must be set to the concatenation of the `stream type`{.variable}
        byte and the zero or more `descriptor bytes`{.variable} bytes,
        expressed in hexadecimal using [ASCII upper hex
        digits](https://infra.spec.whatwg.org/#ascii-upper-hex-digit){#sourcing-in-band-text-tracks:uppercase-ascii-hex-digits
        x-internal="uppercase-ascii-hex-digits"}.
        [\[MPEG2\]](references.html#refsMPEG2)

    If the [media resource](#media-resource){#sourcing-in-band-text-tracks:media-resource-8} is an MPEG-4 file
    :   Let the first `stsd` box of the first `stbl` box of the first
        `minf` box of the first `mdia` box of the [text
        track](#text-track){#sourcing-in-band-text-tracks:text-track-7}\'s
        `trak` box in the first `moov` box of the file be the *stsd
        box*, if any. If the file has no *stsd box*, or if the *stsd
        box* has neither a `mett` box nor a `metx` box, then the [text
        track in-band metadata track dispatch
        type](#text-track-in-band-metadata-track-dispatch-type){#sourcing-in-band-text-tracks:text-track-in-band-metadata-track-dispatch-type-5}
        must be set to the empty string. Otherwise, if the *stsd box*
        has a `mett` box then the [text track in-band metadata track
        dispatch
        type](#text-track-in-band-metadata-track-dispatch-type){#sourcing-in-band-text-tracks:text-track-in-band-metadata-track-dispatch-type-6}
        must be set to the concatenation of the string \"`mett`\", a
        U+0020 SPACE character, and the value of the first `mime_format`
        field of the first `mett` box of the *stsd box*, or the empty
        string if that field is absent in that box. Otherwise, if the
        *stsd box* has no `mett` box but has a `metx` box then the [text
        track in-band metadata track dispatch
        type](#text-track-in-band-metadata-track-dispatch-type){#sourcing-in-band-text-tracks:text-track-in-band-metadata-track-dispatch-type-7}
        must be set to the concatenation of the string \"`metx`\", a
        U+0020 SPACE character, and the value of the first `namespace`
        field of the first `metx` box of the *stsd box*, or the empty
        string if that field is absent in that box.
        [\[MPEG4\]](references.html#refsMPEG4)

5.  Populate the new [text
    track](#text-track){#sourcing-in-band-text-tracks:text-track-8}\'s
    [list of
    cues](#text-track-list-of-cues){#sourcing-in-band-text-tracks:text-track-list-of-cues-2}
    with the cues parsed so far, following the [guidelines for exposing
    cues](#guidelines-for-exposing-cues-in-various-formats-as-text-track-cues){#sourcing-in-band-text-tracks:guidelines-for-exposing-cues-in-various-formats-as-text-track-cues},
    and begin updating it dynamically as necessary.

6.  Set the new [text
    track](#text-track){#sourcing-in-band-text-tracks:text-track-9}\'s
    [readiness
    state](#text-track-readiness-state){#sourcing-in-band-text-tracks:text-track-readiness-state}
    to
    [loaded](#text-track-loaded){#sourcing-in-band-text-tracks:text-track-loaded}.

7.  Set the new [text
    track](#text-track){#sourcing-in-band-text-tracks:text-track-10}\'s
    [mode](#text-track-mode){#sourcing-in-band-text-tracks:text-track-mode}
    to the mode consistent with the user\'s preferences and the
    requirements of the relevant specification for the data.

    For instance, if there are no other active subtitles, and this is a
    forced subtitle track (a subtitle track giving subtitles in the
    audio track\'s primary language, but only for audio that is actually
    in another language), then those subtitles might be activated here.

8.  Add the new [text
    track](#text-track){#sourcing-in-band-text-tracks:text-track-11} to
    the [media
    element](#media-element){#sourcing-in-band-text-tracks:media-element}\'s
    [list of text
    tracks](#list-of-text-tracks){#sourcing-in-band-text-tracks:list-of-text-tracks}.

9.  [Fire an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#sourcing-in-band-text-tracks:concept-event-fire
    x-internal="concept-event-fire"} named
    [`addtrack`{#sourcing-in-band-text-tracks:event-media-addtrack}](#event-media-addtrack)
    at the [media
    element](#media-element){#sourcing-in-band-text-tracks:media-element-2}\'s
    [`textTracks`{#sourcing-in-band-text-tracks:dom-media-texttracks}](#dom-media-texttracks)
    attribute\'s
    [`TextTrackList`{#sourcing-in-band-text-tracks:texttracklist}](#texttracklist)
    object, using
    [`TrackEvent`{#sourcing-in-band-text-tracks:trackevent}](#trackevent),
    with the
    [`track`{#sourcing-in-band-text-tracks:dom-trackevent-track}](#dom-trackevent-track)
    attribute initialized to the [text
    track](#text-track){#sourcing-in-band-text-tracks:text-track-12}\'s
    [`TextTrack`{#sourcing-in-band-text-tracks:texttrack-2}](#texttrack)
    object.

###### [4.8.11.11.3]{.secno} Sourcing out-of-band text tracks[](#sourcing-out-of-band-text-tracks){.self-link}

When a
[`track`{#sourcing-out-of-band-text-tracks:the-track-element}](#the-track-element)
element is created, it must be associated with a new [text
track](#text-track){#sourcing-out-of-band-text-tracks:text-track} (with
its value set as defined below) and its corresponding new
[`TextTrack`{#sourcing-out-of-band-text-tracks:texttrack}](#texttrack)
object.

The [text track
kind](#text-track-kind){#sourcing-out-of-band-text-tracks:text-track-kind}
is determined from the state of the element\'s
[`kind`{#sourcing-out-of-band-text-tracks:attr-track-kind}](#attr-track-kind)
attribute according to the following table; for a state given in a cell
of the first column, the
[kind](#text-track-kind){#sourcing-out-of-band-text-tracks:text-track-kind-2}
is the string given in the second column:

State

String

[Subtitles](#attr-track-kind-subtitles){#sourcing-out-of-band-text-tracks:attr-track-kind-subtitles}

[`subtitles`{#sourcing-out-of-band-text-tracks:dom-texttrack-kind-subtitles}](#dom-texttrack-kind-subtitles)

[Captions](#attr-track-kind-captions){#sourcing-out-of-band-text-tracks:attr-track-kind-captions}

[`captions`{#sourcing-out-of-band-text-tracks:dom-texttrack-kind-captions}](#dom-texttrack-kind-captions)

[Descriptions](#attr-track-kind-descriptions){#sourcing-out-of-band-text-tracks:attr-track-kind-descriptions}

[`descriptions`{#sourcing-out-of-band-text-tracks:dom-texttrack-kind-descriptions}](#dom-texttrack-kind-descriptions)

[Chapters
metadata](#attr-track-kind-chapters){#sourcing-out-of-band-text-tracks:attr-track-kind-chapters}

[`chapters`{#sourcing-out-of-band-text-tracks:dom-texttrack-kind-chapters}](#dom-texttrack-kind-chapters)

[Metadata](#attr-track-kind-metadata){#sourcing-out-of-band-text-tracks:attr-track-kind-metadata}

[`metadata`{#sourcing-out-of-band-text-tracks:dom-texttrack-kind-metadata}](#dom-texttrack-kind-metadata)

The [text track
label](#text-track-label){#sourcing-out-of-band-text-tracks:text-track-label}
is the element\'s [track
label](#track-label){#sourcing-out-of-band-text-tracks:track-label}.

The [text track
language](#text-track-language){#sourcing-out-of-band-text-tracks:text-track-language}
is the element\'s [track
language](#track-language){#sourcing-out-of-band-text-tracks:track-language},
if any, or the empty string otherwise.

As the
[`kind`{#sourcing-out-of-band-text-tracks:attr-track-kind-2}](#attr-track-kind),
[`label`{#sourcing-out-of-band-text-tracks:attr-track-label}](#attr-track-label),
and
[`srclang`{#sourcing-out-of-band-text-tracks:attr-track-srclang}](#attr-track-srclang)
attributes are set, changed, or removed, the [text
track](#text-track){#sourcing-out-of-band-text-tracks:text-track-2} must
update accordingly, as per the definitions above.

Changes to the [track
URL](#track-url){#sourcing-out-of-band-text-tracks:track-url} are
handled in the algorithm below.

The [text track readiness
state](#text-track-readiness-state){#sourcing-out-of-band-text-tracks:text-track-readiness-state}
is initially [not
loaded](#text-track-not-loaded){#sourcing-out-of-band-text-tracks:text-track-not-loaded},
and the [text track
mode](#text-track-mode){#sourcing-out-of-band-text-tracks:text-track-mode}
is initially
[disabled](#text-track-disabled){#sourcing-out-of-band-text-tracks:text-track-disabled}.

The [text track list of
cues](#text-track-list-of-cues){#sourcing-out-of-band-text-tracks:text-track-list-of-cues}
is initially empty. It is dynamically modified when the referenced file
is parsed. Associated with the list are the [rules for updating the text
track
rendering](#rules-for-updating-the-text-track-rendering){#sourcing-out-of-band-text-tracks:rules-for-updating-the-text-track-rendering}
appropriate for the format in question; for WebVTT, this is the [rules
for updating the display of WebVTT text
tracks](https://w3c.github.io/webvtt/#rules-for-updating-the-display-of-webvtt-text-tracks){#sourcing-out-of-band-text-tracks:rules-for-updating-the-display-of-webvtt-text-tracks
x-internal="rules-for-updating-the-display-of-webvtt-text-tracks"}.
[\[WEBVTT\]](references.html#refsWEBVTT)

When a
[`track`{#sourcing-out-of-band-text-tracks:the-track-element-2}](#the-track-element)
element\'s parent element changes and the new parent is a [media
element](#media-element){#sourcing-out-of-band-text-tracks:media-element},
then the user agent must add the
[`track`{#sourcing-out-of-band-text-tracks:the-track-element-3}](#the-track-element)
element\'s corresponding [text
track](#text-track){#sourcing-out-of-band-text-tracks:text-track-3} to
the [media
element](#media-element){#sourcing-out-of-band-text-tracks:media-element-2}\'s
[list of text
tracks](#list-of-text-tracks){#sourcing-out-of-band-text-tracks:list-of-text-tracks},
and then [queue a media element
task](#queue-a-media-element-task){#sourcing-out-of-band-text-tracks:queue-a-media-element-task}
given the [media
element](#media-element){#sourcing-out-of-band-text-tracks:media-element-3}
to [fire an
event](https://dom.spec.whatwg.org/#concept-event-fire){#sourcing-out-of-band-text-tracks:concept-event-fire
x-internal="concept-event-fire"} named
[`addtrack`{#sourcing-out-of-band-text-tracks:event-media-addtrack}](#event-media-addtrack)
at the [media
element](#media-element){#sourcing-out-of-band-text-tracks:media-element-4}\'s
[`textTracks`{#sourcing-out-of-band-text-tracks:dom-media-texttracks}](#dom-media-texttracks)
attribute\'s
[`TextTrackList`{#sourcing-out-of-band-text-tracks:texttracklist}](#texttracklist)
object, using
[`TrackEvent`{#sourcing-out-of-band-text-tracks:trackevent}](#trackevent),
with the
[`track`{#sourcing-out-of-band-text-tracks:dom-trackevent-track}](#dom-trackevent-track)
attribute initialized to the [text
track](#text-track){#sourcing-out-of-band-text-tracks:text-track-4}\'s
[`TextTrack`{#sourcing-out-of-band-text-tracks:texttrack-2}](#texttrack)
object.

When a
[`track`{#sourcing-out-of-band-text-tracks:the-track-element-4}](#the-track-element)
element\'s parent element changes and the old parent was a [media
element](#media-element){#sourcing-out-of-band-text-tracks:media-element-5},
then the user agent must remove the
[`track`{#sourcing-out-of-band-text-tracks:the-track-element-5}](#the-track-element)
element\'s corresponding [text
track](#text-track){#sourcing-out-of-band-text-tracks:text-track-5} from
the [media
element](#media-element){#sourcing-out-of-band-text-tracks:media-element-6}\'s
[list of text
tracks](#list-of-text-tracks){#sourcing-out-of-band-text-tracks:list-of-text-tracks-2},
and then [queue a media element
task](#queue-a-media-element-task){#sourcing-out-of-band-text-tracks:queue-a-media-element-task-2}
given the [media
element](#media-element){#sourcing-out-of-band-text-tracks:media-element-7}
to [fire an
event](https://dom.spec.whatwg.org/#concept-event-fire){#sourcing-out-of-band-text-tracks:concept-event-fire-2
x-internal="concept-event-fire"} named
[`removetrack`{#sourcing-out-of-band-text-tracks:event-media-removetrack}](#event-media-removetrack)
at the [media
element](#media-element){#sourcing-out-of-band-text-tracks:media-element-8}\'s
[`textTracks`{#sourcing-out-of-band-text-tracks:dom-media-texttracks-2}](#dom-media-texttracks)
attribute\'s
[`TextTrackList`{#sourcing-out-of-band-text-tracks:texttracklist-2}](#texttracklist)
object, using
[`TrackEvent`{#sourcing-out-of-band-text-tracks:trackevent-2}](#trackevent),
with the
[`track`{#sourcing-out-of-band-text-tracks:dom-trackevent-track-2}](#dom-trackevent-track)
attribute initialized to the [text
track](#text-track){#sourcing-out-of-band-text-tracks:text-track-6}\'s
[`TextTrack`{#sourcing-out-of-band-text-tracks:texttrack-3}](#texttrack)
object.

------------------------------------------------------------------------

When a [text
track](#text-track){#sourcing-out-of-band-text-tracks:text-track-7}
corresponding to a
[`track`{#sourcing-out-of-band-text-tracks:the-track-element-6}](#the-track-element)
element is added to a [media
element](#media-element){#sourcing-out-of-band-text-tracks:media-element-9}\'s
[list of text
tracks](#list-of-text-tracks){#sourcing-out-of-band-text-tracks:list-of-text-tracks-3},
the user agent must [queue a media element
task](#queue-a-media-element-task){#sourcing-out-of-band-text-tracks:queue-a-media-element-task-3}
given the [media
element](#media-element){#sourcing-out-of-band-text-tracks:media-element-10}
to run the following steps for the [media
element](#media-element){#sourcing-out-of-band-text-tracks:media-element-11}:

1.  If the element\'s
    [blocked-on-parser](#blocked-on-parser){#sourcing-out-of-band-text-tracks:blocked-on-parser}
    flag is true, then return.

2.  If the element\'s
    [did-perform-automatic-track-selection](#did-perform-automatic-track-selection){#sourcing-out-of-band-text-tracks:did-perform-automatic-track-selection}
    flag is true, then return.

3.  [Honor user preferences for automatic text track
    selection](#honor-user-preferences-for-automatic-text-track-selection){#sourcing-out-of-band-text-tracks:honor-user-preferences-for-automatic-text-track-selection}
    for this element.

When the user agent is required to [honor user preferences for automatic
text track
selection]{#honor-user-preferences-for-automatic-text-track-selection
.dfn} for a [media
element](#media-element){#sourcing-out-of-band-text-tracks:media-element-12},
the user agent must run the following steps:

1.  [Perform automatic text track
    selection](#perform-automatic-text-track-selection){#sourcing-out-of-band-text-tracks:perform-automatic-text-track-selection}
    for
    [`subtitles`{#sourcing-out-of-band-text-tracks:dom-texttrack-kind-subtitles-2}](#dom-texttrack-kind-subtitles)
    and
    [`captions`{#sourcing-out-of-band-text-tracks:dom-texttrack-kind-captions-2}](#dom-texttrack-kind-captions).

2.  [Perform automatic text track
    selection](#perform-automatic-text-track-selection){#sourcing-out-of-band-text-tracks:perform-automatic-text-track-selection-2}
    for
    [`descriptions`{#sourcing-out-of-band-text-tracks:dom-texttrack-kind-descriptions-2}](#dom-texttrack-kind-descriptions).

3.  If there are any [text
    tracks](#text-track){#sourcing-out-of-band-text-tracks:text-track-8}
    in the [media
    element](#media-element){#sourcing-out-of-band-text-tracks:media-element-13}\'s
    [list of text
    tracks](#list-of-text-tracks){#sourcing-out-of-band-text-tracks:list-of-text-tracks-4}
    whose [text track
    kind](#text-track-kind){#sourcing-out-of-band-text-tracks:text-track-kind-3}
    is
    [`chapters`{#sourcing-out-of-band-text-tracks:dom-texttrack-kind-metadata-2}](#dom-texttrack-kind-metadata)
    or
    [`metadata`{#sourcing-out-of-band-text-tracks:dom-texttrack-kind-metadata-3}](#dom-texttrack-kind-metadata)
    that correspond to
    [`track`{#sourcing-out-of-band-text-tracks:the-track-element-7}](#the-track-element)
    elements with a
    [`default`{#sourcing-out-of-band-text-tracks:attr-track-default}](#attr-track-default)
    attribute set whose [text track
    mode](#text-track-mode){#sourcing-out-of-band-text-tracks:text-track-mode-2}
    is set to
    [disabled](#text-track-disabled){#sourcing-out-of-band-text-tracks:text-track-disabled-2},
    then set the [text track
    mode](#text-track-mode){#sourcing-out-of-band-text-tracks:text-track-mode-3}
    of all such tracks to
    [hidden](#text-track-hidden){#sourcing-out-of-band-text-tracks:text-track-hidden}.

4.  Set the element\'s
    [did-perform-automatic-track-selection](#did-perform-automatic-track-selection){#sourcing-out-of-band-text-tracks:did-perform-automatic-track-selection-2}
    flag to true.

When the steps above say to [perform automatic text track
selection]{#perform-automatic-text-track-selection .dfn} for one or more
[text track
kinds](#text-track-kind){#sourcing-out-of-band-text-tracks:text-track-kind-4},
it means to run the following steps:

1.  Let `candidates`{.variable} be a list consisting of the [text
    tracks](#text-track){#sourcing-out-of-band-text-tracks:text-track-9}
    in the [media
    element](#media-element){#sourcing-out-of-band-text-tracks:media-element-14}\'s
    [list of text
    tracks](#list-of-text-tracks){#sourcing-out-of-band-text-tracks:list-of-text-tracks-5}
    whose [text track
    kind](#text-track-kind){#sourcing-out-of-band-text-tracks:text-track-kind-5}
    is one of the kinds that were passed to the algorithm, if any, in
    the order given in the [list of text
    tracks](#list-of-text-tracks){#sourcing-out-of-band-text-tracks:list-of-text-tracks-6}.

2.  If `candidates`{.variable} is empty, then return.

3.  If any of the [text
    tracks](#text-track){#sourcing-out-of-band-text-tracks:text-track-10}
    in `candidates`{.variable} have a [text track
    mode](#text-track-mode){#sourcing-out-of-band-text-tracks:text-track-mode-4}
    set to
    [showing](#text-track-showing){#sourcing-out-of-band-text-tracks:text-track-showing},
    return.

4.  If the user has expressed an interest in having a track from
    `candidates`{.variable} enabled based on its [text track
    kind](#text-track-kind){#sourcing-out-of-band-text-tracks:text-track-kind-6},
    [text track
    language](#text-track-language){#sourcing-out-of-band-text-tracks:text-track-language-2},
    and [text track
    label](#text-track-label){#sourcing-out-of-band-text-tracks:text-track-label-2},
    then set its [text track
    mode](#text-track-mode){#sourcing-out-of-band-text-tracks:text-track-mode-5}
    to
    [showing](#text-track-showing){#sourcing-out-of-band-text-tracks:text-track-showing-2}.

    For example, the user could have set a browser preference to the
    effect of \"I want French captions whenever possible\", or \"If
    there is a subtitle track with \'Commentary\' in the title, enable
    it\", or \"If there are audio description tracks available, enable
    one, ideally in Swiss German, but failing that in Standard Swiss
    German or Standard German\".

    Otherwise, if there are any [text
    tracks](#text-track){#sourcing-out-of-band-text-tracks:text-track-11}
    in `candidates`{.variable} that correspond to
    [`track`{#sourcing-out-of-band-text-tracks:the-track-element-8}](#the-track-element)
    elements with a
    [`default`{#sourcing-out-of-band-text-tracks:attr-track-default-2}](#attr-track-default)
    attribute set whose [text track
    mode](#text-track-mode){#sourcing-out-of-band-text-tracks:text-track-mode-6}
    is set to
    [disabled](#text-track-disabled){#sourcing-out-of-band-text-tracks:text-track-disabled-3},
    then set the [text track
    mode](#text-track-mode){#sourcing-out-of-band-text-tracks:text-track-mode-7}
    of the first such track to
    [showing](#text-track-showing){#sourcing-out-of-band-text-tracks:text-track-showing-3}.

When a [text
track](#text-track){#sourcing-out-of-band-text-tracks:text-track-12}
corresponding to a
[`track`{#sourcing-out-of-band-text-tracks:the-track-element-9}](#the-track-element)
element experiences any of the following circumstances, the user agent
must [start the `track` processing
model](#start-the-track-processing-model){#sourcing-out-of-band-text-tracks:start-the-track-processing-model}
for that [text
track](#text-track){#sourcing-out-of-band-text-tracks:text-track-13} and
its
[`track`{#sourcing-out-of-band-text-tracks:the-track-element-10}](#the-track-element)
element:

- The
  [`track`{#sourcing-out-of-band-text-tracks:the-track-element-11}](#the-track-element)
  element is created.
- The [text
  track](#text-track){#sourcing-out-of-band-text-tracks:text-track-14}
  has its [text track
  mode](#text-track-mode){#sourcing-out-of-band-text-tracks:text-track-mode-8}
  changed.
- The
  [`track`{#sourcing-out-of-band-text-tracks:the-track-element-12}](#the-track-element)
  element\'s parent element changes and the new parent is a [media
  element](#media-element){#sourcing-out-of-band-text-tracks:media-element-15}.

When a user agent is to [start the `track` processing
model]{#start-the-track-processing-model .dfn} for a [text
track](#text-track){#sourcing-out-of-band-text-tracks:text-track-15} and
its
[`track`{#sourcing-out-of-band-text-tracks:the-track-element-13}](#the-track-element)
element, it must run the following algorithm. This algorithm interacts
closely with the [event
loop](webappapis.html#event-loop){#sourcing-out-of-band-text-tracks:event-loop}
mechanism; in particular, it has a [synchronous
section](webappapis.html#synchronous-section){#sourcing-out-of-band-text-tracks:synchronous-section}
(which is triggered as part of the [event
loop](webappapis.html#event-loop){#sourcing-out-of-band-text-tracks:event-loop-2}
algorithm). The steps in that section are marked with ⌛.

1.  If another occurrence of this algorithm is already running for this
    [text
    track](#text-track){#sourcing-out-of-band-text-tracks:text-track-16}
    and its
    [`track`{#sourcing-out-of-band-text-tracks:the-track-element-14}](#the-track-element)
    element, return, letting that other algorithm take care of this
    element.

2.  If the [text
    track](#text-track){#sourcing-out-of-band-text-tracks:text-track-17}\'s
    [text track
    mode](#text-track-mode){#sourcing-out-of-band-text-tracks:text-track-mode-9}
    is not set to one of
    [hidden](#text-track-hidden){#sourcing-out-of-band-text-tracks:text-track-hidden-2}
    or
    [showing](#text-track-showing){#sourcing-out-of-band-text-tracks:text-track-showing-4},
    then return.

3.  If the [text
    track](#text-track){#sourcing-out-of-band-text-tracks:text-track-18}\'s
    [`track`{#sourcing-out-of-band-text-tracks:the-track-element-15}](#the-track-element)
    element does not have a [media
    element](#media-element){#sourcing-out-of-band-text-tracks:media-element-16}
    as a parent, return.

4.  Run the remainder of these steps [in
    parallel](infrastructure.html#in-parallel){#sourcing-out-of-band-text-tracks:in-parallel},
    allowing whatever caused these steps to run to continue.

5.  *Top*: [Await a stable
    state](webappapis.html#await-a-stable-state){#sourcing-out-of-band-text-tracks:await-a-stable-state}.
    The [synchronous
    section](webappapis.html#synchronous-section){#sourcing-out-of-band-text-tracks:synchronous-section-2}
    consists of the following steps. (The steps in the [synchronous
    section](webappapis.html#synchronous-section){#sourcing-out-of-band-text-tracks:synchronous-section-3}
    are marked with ⌛.)

6.  ⌛ Set the [text track readiness
    state](#text-track-readiness-state){#sourcing-out-of-band-text-tracks:text-track-readiness-state-2}
    to
    [loading](#text-track-loading){#sourcing-out-of-band-text-tracks:text-track-loading}.

7.  ⌛ Let `URL`{.variable} be the [track
    URL](#track-url){#sourcing-out-of-band-text-tracks:track-url-2} of
    the
    [`track`{#sourcing-out-of-band-text-tracks:the-track-element-16}](#the-track-element)
    element.

8.  ⌛ If the
    [`track`{#sourcing-out-of-band-text-tracks:the-track-element-17}](#the-track-element)
    element\'s parent is a [media
    element](#media-element){#sourcing-out-of-band-text-tracks:media-element-17},
    then let `corsAttributeState`{.variable} be the state of the parent
    [media
    element](#media-element){#sourcing-out-of-band-text-tracks:media-element-18}\'s
    [`crossorigin`{#sourcing-out-of-band-text-tracks:attr-media-crossorigin}](#attr-media-crossorigin)
    content attribute. Otherwise, let `corsAttributeState`{.variable} be
    [No
    CORS](urls-and-fetching.html#attr-crossorigin-none){#sourcing-out-of-band-text-tracks:attr-crossorigin-none}.

9.  End the [synchronous
    section](webappapis.html#synchronous-section){#sourcing-out-of-band-text-tracks:synchronous-section-4},
    continuing the remaining steps [in
    parallel](infrastructure.html#in-parallel){#sourcing-out-of-band-text-tracks:in-parallel-2}.

10. If `URL`{.variable} is not the empty string, then:

    1.  Let `request`{.variable} be the result of [creating a
        potential-CORS
        request](urls-and-fetching.html#create-a-potential-cors-request){#sourcing-out-of-band-text-tracks:create-a-potential-cors-request}
        given `URL`{.variable}, \"`track`\", and
        `corsAttributeState`{.variable}, and with the *same-origin
        fallback flag* set.

    2.  Set `request`{.variable}\'s
        [client](https://fetch.spec.whatwg.org/#concept-request-client){#sourcing-out-of-band-text-tracks:concept-request-client
        x-internal="concept-request-client"} to the
        [`track`{#sourcing-out-of-band-text-tracks:the-track-element-18}](#the-track-element)
        element\'s [node
        document](https://dom.spec.whatwg.org/#concept-node-document){#sourcing-out-of-band-text-tracks:node-document
        x-internal="node-document"}\'s [relevant settings
        object](webappapis.html#relevant-settings-object){#sourcing-out-of-band-text-tracks:relevant-settings-object}.

    3.  Set `request`{.variable}\'s [initiator
        type](https://fetch.spec.whatwg.org/#request-initiator-type){#sourcing-out-of-band-text-tracks:concept-request-initiator-type
        x-internal="concept-request-initiator-type"} to \"`track`\".

    4.  [Fetch](https://fetch.spec.whatwg.org/#concept-fetch){#sourcing-out-of-band-text-tracks:concept-fetch
        x-internal="concept-fetch"} `request`{.variable}.

    The
    [tasks](webappapis.html#concept-task){#sourcing-out-of-band-text-tracks:concept-task}
    [queued](webappapis.html#queue-a-task){#sourcing-out-of-band-text-tracks:queue-a-task}
    by the fetching algorithm on the [networking task
    source](webappapis.html#networking-task-source){#sourcing-out-of-band-text-tracks:networking-task-source}
    to process the data as it is being fetched must determine the type
    of the resource. If the type of the resource is not a supported text
    track format, the load will fail, as described below. Otherwise, the
    resource\'s data must be passed to the appropriate parser (e.g., the
    [WebVTT
    parser](https://w3c.github.io/webvtt/#webvtt-parser){#sourcing-out-of-band-text-tracks:webvtt-parser
    x-internal="webvtt-parser"}) as it is received, with the [text track
    list of
    cues](#text-track-list-of-cues){#sourcing-out-of-band-text-tracks:text-track-list-of-cues-2}
    being used for that parser\'s output.
    [\[WEBVTT\]](references.html#refsWEBVTT)

    The appropriate parser will incrementally update the [text track
    list of
    cues](#text-track-list-of-cues){#sourcing-out-of-band-text-tracks:text-track-list-of-cues-3}
    during these [networking task
    source](webappapis.html#networking-task-source){#sourcing-out-of-band-text-tracks:networking-task-source-2}
    [tasks](webappapis.html#concept-task){#sourcing-out-of-band-text-tracks:concept-task-2},
    as each such task is run with whatever data has been received from
    the network).

    This specification does not currently say whether or how to check
    the MIME types of text tracks, or whether or how to perform file
    type sniffing using the actual file data. Implementers differ in
    their intentions on this matter and it is therefore unclear what the
    right solution is. In the absence of any requirement here, the HTTP
    specifications\' strict requirement to follow the Content-Type
    header prevails (\"Content-Type specifies the media type of the
    underlying data.\" \... \"If and only if the media type is not given
    by a Content-Type field, the recipient MAY attempt to guess the
    media type via inspection of its content and/or the name
    extension(s) of the URI used to identify the resource.\").

    If fetching fails for any reason (network error, the server returns
    an error code, CORS fails, etc.), or if `URL`{.variable} is the
    empty string, then [queue an element
    task](webappapis.html#queue-an-element-task){#sourcing-out-of-band-text-tracks:queue-an-element-task}
    on the [DOM manipulation task
    source](webappapis.html#dom-manipulation-task-source){#sourcing-out-of-band-text-tracks:dom-manipulation-task-source}
    given the [media
    element](#media-element){#sourcing-out-of-band-text-tracks:media-element-19}
    to first change the [text track readiness
    state](#text-track-readiness-state){#sourcing-out-of-band-text-tracks:text-track-readiness-state-3}
    to [failed to
    load](#text-track-failed-to-load){#sourcing-out-of-band-text-tracks:text-track-failed-to-load}
    and then [fire an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#sourcing-out-of-band-text-tracks:concept-event-fire-3
    x-internal="concept-event-fire"} named
    [`error`{#sourcing-out-of-band-text-tracks:event-track-error}](#event-track-error)
    at the
    [`track`{#sourcing-out-of-band-text-tracks:the-track-element-19}](#the-track-element)
    element.

    If fetching does not fail, but the type of the resource is not a
    supported text track format, or the file was not successfully
    processed (e.g., the format in question is an XML format and the
    file contained a well-formedness error that XML requires be detected
    and reported to the application), then the
    [task](webappapis.html#concept-task){#sourcing-out-of-band-text-tracks:concept-task-3}
    that is
    [queued](webappapis.html#queue-an-element-task){#sourcing-out-of-band-text-tracks:queue-an-element-task-2}
    on the [networking task
    source](webappapis.html#networking-task-source){#sourcing-out-of-band-text-tracks:networking-task-source-3}
    in which the aforementioned problem is found must change the [text
    track readiness
    state](#text-track-readiness-state){#sourcing-out-of-band-text-tracks:text-track-readiness-state-4}
    to [failed to
    load](#text-track-failed-to-load){#sourcing-out-of-band-text-tracks:text-track-failed-to-load-2}
    and [fire an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#sourcing-out-of-band-text-tracks:concept-event-fire-4
    x-internal="concept-event-fire"} named
    [`error`{#sourcing-out-of-band-text-tracks:event-track-error-2}](#event-track-error)
    at the
    [`track`{#sourcing-out-of-band-text-tracks:the-track-element-20}](#the-track-element)
    element.

    If fetching does not fail, and the file was successfully processed,
    then the final
    [task](webappapis.html#concept-task){#sourcing-out-of-band-text-tracks:concept-task-4}
    that is
    [queued](webappapis.html#queue-an-element-task){#sourcing-out-of-band-text-tracks:queue-an-element-task-3}
    by the [networking task
    source](webappapis.html#networking-task-source){#sourcing-out-of-band-text-tracks:networking-task-source-4},
    after it has finished parsing the data, must change the [text track
    readiness
    state](#text-track-readiness-state){#sourcing-out-of-band-text-tracks:text-track-readiness-state-5}
    to
    [loaded](#text-track-loaded){#sourcing-out-of-band-text-tracks:text-track-loaded},
    and [fire an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#sourcing-out-of-band-text-tracks:concept-event-fire-5
    x-internal="concept-event-fire"} named
    [`load`{#sourcing-out-of-band-text-tracks:event-track-load}](#event-track-load)
    at the
    [`track`{#sourcing-out-of-band-text-tracks:the-track-element-21}](#the-track-element)
    element.

    If, while fetching is ongoing, either:

    - the [track
      URL](#track-url){#sourcing-out-of-band-text-tracks:track-url-3}
      changes so that it is no longer equal to `URL`{.variable}, while
      the [text track
      mode](#text-track-mode){#sourcing-out-of-band-text-tracks:text-track-mode-10}
      is set to
      [hidden](#text-track-hidden){#sourcing-out-of-band-text-tracks:text-track-hidden-3}
      or
      [showing](#text-track-showing){#sourcing-out-of-band-text-tracks:text-track-showing-5};
      or
    - the [text track
      mode](#text-track-mode){#sourcing-out-of-band-text-tracks:text-track-mode-11}
      changes to
      [hidden](#text-track-hidden){#sourcing-out-of-band-text-tracks:text-track-hidden-4}
      or
      [showing](#text-track-showing){#sourcing-out-of-band-text-tracks:text-track-showing-6},
      while the [track
      URL](#track-url){#sourcing-out-of-band-text-tracks:track-url-4} is
      not equal to `URL`{.variable},

    \...then the user agent must abort
    [fetching](https://fetch.spec.whatwg.org/#concept-fetch){#sourcing-out-of-band-text-tracks:concept-fetch-2
    x-internal="concept-fetch"}, discarding any pending
    [tasks](webappapis.html#concept-task){#sourcing-out-of-band-text-tracks:concept-task-5}
    generated by that algorithm (and in particular, not adding any cues
    to the [text track list of
    cues](#text-track-list-of-cues){#sourcing-out-of-band-text-tracks:text-track-list-of-cues-4}
    after the moment the URL changed), and then [queue an element
    task](webappapis.html#queue-an-element-task){#sourcing-out-of-band-text-tracks:queue-an-element-task-4}
    on the [DOM manipulation task
    source](webappapis.html#dom-manipulation-task-source){#sourcing-out-of-band-text-tracks:dom-manipulation-task-source-2}
    given the
    [`track`{#sourcing-out-of-band-text-tracks:the-track-element-22}](#the-track-element)
    element that first changes the [text track readiness
    state](#text-track-readiness-state){#sourcing-out-of-band-text-tracks:text-track-readiness-state-6}
    to [failed to
    load](#text-track-failed-to-load){#sourcing-out-of-band-text-tracks:text-track-failed-to-load-3}
    and then [fires an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#sourcing-out-of-band-text-tracks:concept-event-fire-6
    x-internal="concept-event-fire"} named
    [`error`{#sourcing-out-of-band-text-tracks:event-track-error-3}](#event-track-error)
    at the
    [`track`{#sourcing-out-of-band-text-tracks:the-track-element-23}](#the-track-element)
    element.

11. Wait until the [text track readiness
    state](#text-track-readiness-state){#sourcing-out-of-band-text-tracks:text-track-readiness-state-7}
    is no longer set to
    [loading](#text-track-loading){#sourcing-out-of-band-text-tracks:text-track-loading-2}.

12. Wait until the [track
    URL](#track-url){#sourcing-out-of-band-text-tracks:track-url-5} is
    no longer equal to `URL`{.variable}, at the same time as the [text
    track
    mode](#text-track-mode){#sourcing-out-of-band-text-tracks:text-track-mode-12}
    is set to
    [hidden](#text-track-hidden){#sourcing-out-of-band-text-tracks:text-track-hidden-5}
    or
    [showing](#text-track-showing){#sourcing-out-of-band-text-tracks:text-track-showing-7}.

13. Jump to the step labeled *top*.

Whenever a
[`track`{#sourcing-out-of-band-text-tracks:the-track-element-24}](#the-track-element)
element has its
[`src`{#sourcing-out-of-band-text-tracks:attr-track-src}](#attr-track-src)
attribute set, changed, or removed, the user agent must
[immediately](infrastructure.html#immediately){#sourcing-out-of-band-text-tracks:immediately}
empty the element\'s [text
track](#text-track){#sourcing-out-of-band-text-tracks:text-track-19}\'s
[text track list of
cues](#text-track-list-of-cues){#sourcing-out-of-band-text-tracks:text-track-list-of-cues-5}.
(This also causes the algorithm above to stop adding cues from the
resource being obtained using the previously given URL, if any.)

###### [4.8.11.11.4]{.secno} [Guidelines for exposing cues]{.dfn} in various formats as [text track cues](#text-track-cue){#guidelines-for-exposing-cues-in-various-formats-as-text-track-cues:text-track-cue}[](#guidelines-for-exposing-cues-in-various-formats-as-text-track-cues){.self-link}

How a specific format\'s text track cues are to be interpreted for the
purposes of processing by an HTML user agent is defined by that format.
In the absence of such a specification, this section provides some
constraints within which implementations can attempt to consistently
expose such formats.

To support the [text
track](#text-track){#guidelines-for-exposing-cues-in-various-formats-as-text-track-cues:text-track}
model of HTML, each unit of timed data is converted to a [text track
cue](#text-track-cue){#guidelines-for-exposing-cues-in-various-formats-as-text-track-cues:text-track-cue-2}.
Where the mapping of the format\'s features to the aspects of a [text
track
cue](#text-track-cue){#guidelines-for-exposing-cues-in-various-formats-as-text-track-cues:text-track-cue-3}
as defined in this specification are not defined, implementations must
ensure that the mapping is consistent with the definitions of the
aspects of a [text track
cue](#text-track-cue){#guidelines-for-exposing-cues-in-various-formats-as-text-track-cues:text-track-cue-4}
as defined above, as well as with the following constraints:

The [text track cue identifier](#text-track-cue-identifier){#guidelines-for-exposing-cues-in-various-formats-as-text-track-cues:text-track-cue-identifier}

:   Should be set to the empty string if the format has no obvious
    analogue to a per-cue identifier.

The [text track cue pause-on-exit flag](#text-track-cue-pause-on-exit-flag){#guidelines-for-exposing-cues-in-various-formats-as-text-track-cues:text-track-cue-pause-on-exit-flag}

:   Should be set to false.

###### [4.8.11.11.5]{.secno} Text track API[](#text-track-api){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[TextTrackList](https://developer.mozilla.org/en-US/docs/Web/API/TextTrackList "The TextTrackList interface is used to represent a list of the text tracks defined by the <track> element, with each track represented by a separate textTrack object in the list.")

Support in all current engines.

::: support
[Firefox31+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome23+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS7+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

``` idl
[Exposed=Window]
interface TextTrackList : EventTarget {
  readonly attribute unsigned long length;
  getter TextTrack (unsigned long index);
  TextTrack? getTrackById(DOMString id);

  attribute EventHandler onchange;
  attribute EventHandler onaddtrack;
  attribute EventHandler onremovetrack;
};
```

`media`{.variable}`.`[`textTracks`](#dom-media-texttracks){#dom-media-texttracks-dev}`.``length`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/textTracks](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/textTracks "The read-only textTracks property on HTMLMediaElement objects returns a TextTrackList object listing all of the TextTrack objects representing the media element's text tracks, in the same order as in the list of text tracks.")

Support in all current engines.

::: support
[Firefox31+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome23+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet1.0+]{.samsunginternet_android .yes}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns the number of [text
tracks](#text-track){#text-track-api:text-track} associated with the
[media element](#media-element){#text-track-api:media-element} (e.g.
from [`track`{#text-track-api:the-track-element}](#the-track-element)
elements). This is the number of [text
tracks](#text-track){#text-track-api:text-track-2} in the [media
element](#media-element){#text-track-api:media-element-2}\'s [list of
text tracks](#list-of-text-tracks){#text-track-api:list-of-text-tracks}.

`media`{.variable}`.`[`textTracks[`](#dom-media-texttracks){#text-track-api:dom-media-texttracks}` ``n`{.variable}` ``]`

Returns the [`TextTrack`{#text-track-api:texttrack-3}](#texttrack)
object representing the `n`{.variable}th [text
track](#text-track){#text-track-api:text-track-3} in the [media
element](#media-element){#text-track-api:media-element-3}\'s [list of
text
tracks](#list-of-text-tracks){#text-track-api:list-of-text-tracks-2}.

`textTrack`{.variable}` = ``media`{.variable}`.`[`textTracks`](#dom-media-texttracks){#text-track-api:dom-media-texttracks-2}`.`[`getTrackById`](#dom-texttracklist-gettrackbyid){#dom-texttracklist-gettrackbyid-dev}`(``id`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[TextTrackList/getTrackById](https://developer.mozilla.org/en-US/docs/Web/API/TextTrackList/getTrackById "The TextTrackList method getTrackById() returns the first TextTrack object from the track list whose id matches the specified string. This lets you find a specified track if you know its ID string.")

Support in all current engines.

::: support
[Firefox31+]{.firefox .yes}[Safari8+]{.safari .yes}[Chrome33+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)18]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

Returns the [`TextTrack`{#text-track-api:texttrack-4}](#texttrack)
object with the given identifier, or null if no track has that
identifier.

A [`TextTrackList`{#text-track-api:texttracklist}](#texttracklist)
object represents a dynamically updating list of [text
tracks](#text-track){#text-track-api:text-track-4} in a given order.

The [`textTracks`]{#dom-media-texttracks .dfn dfn-for="HTMLMediaElement"
dfn-type="attribute"} attribute of [media
elements](#media-element){#text-track-api:media-element-4} must return a
[`TextTrackList`{#text-track-api:texttracklist-2}](#texttracklist)
object representing the
[`TextTrack`{#text-track-api:texttrack-5}](#texttrack) objects of the
[text tracks](#text-track){#text-track-api:text-track-5} in the [media
element](#media-element){#text-track-api:media-element-5}\'s [list of
text
tracks](#list-of-text-tracks){#text-track-api:list-of-text-tracks-3}, in
the same order as in the [list of text
tracks](#list-of-text-tracks){#text-track-api:list-of-text-tracks-4}.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[TextTrackList/length](https://developer.mozilla.org/en-US/docs/Web/API/TextTrackList/length "The read-only TextTrackList property length returns the number of entries in the TextTrackList, each of which is a TextTrack representing one track in the media element.")

Support in all current engines.

::: support
[Firefox31+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome23+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS7+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The [`length`]{#dom-texttracklist-length .dfn dfn-for="TextTrackList"
dfn-type="attribute"} attribute of a
[`TextTrackList`{#text-track-api:texttracklist-3}](#texttracklist)
object must return the number of [text
tracks](#text-track){#text-track-api:text-track-6} in the list
represented by the
[`TextTrackList`{#text-track-api:texttracklist-4}](#texttracklist)
object.

The [supported property
indices](https://webidl.spec.whatwg.org/#dfn-supported-property-indices){#text-track-api:supported-property-indices
x-internal="supported-property-indices"} of a
[`TextTrackList`{#text-track-api:texttracklist-5}](#texttracklist)
object at any instant are the numbers from zero to the number of [text
tracks](#text-track){#text-track-api:text-track-7} in the list
represented by the
[`TextTrackList`{#text-track-api:texttracklist-6}](#texttracklist)
object minus one, if any. If there are no [text
tracks](#text-track){#text-track-api:text-track-8} in the list, there
are no [supported property
indices](https://webidl.spec.whatwg.org/#dfn-supported-property-indices){#text-track-api:supported-property-indices-2
x-internal="supported-property-indices"}.

To [determine the value of an indexed
property](https://webidl.spec.whatwg.org/#dfn-determine-the-value-of-an-indexed-property){#text-track-api:determine-the-value-of-an-indexed-property
x-internal="determine-the-value-of-an-indexed-property"} of a
[`TextTrackList`{#text-track-api:texttracklist-7}](#texttracklist)
object for a given index `index`{.variable}, the user agent must return
the `index`{.variable}th [text
track](#text-track){#text-track-api:text-track-9} in the list
represented by the
[`TextTrackList`{#text-track-api:texttracklist-8}](#texttracklist)
object.

The [`getTrackById(``id`{.variable}`)`]{#dom-texttracklist-gettrackbyid
.dfn dfn-for="TextTrackList" dfn-type="method"} method must return the
first [`TextTrack`{#text-track-api:texttrack-6}](#texttrack) in the
[`TextTrackList`{#text-track-api:texttracklist-9}](#texttracklist)
object whose [`id`{#text-track-api:dom-texttrack-id}](#dom-texttrack-id)
IDL attribute would return a value equal to the value of the
`id`{.variable} argument. When no tracks match the given argument, the
method must return null.

------------------------------------------------------------------------

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[TextTrack](https://developer.mozilla.org/en-US/docs/Web/API/TextTrack "The TextTrack interface—part of the API for handling WebVTT (text tracks on media presentations)—describes and controls the text track associated with a particular <track> element.")

Support in all current engines.

::: support
[Firefox31+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome23+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android31+]{.firefox_android .yes}[Safari iOS7+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

``` idl
enum TextTrackMode { "disabled",  "hidden",  "showing" };
enum TextTrackKind { "subtitles",  "captions",  "descriptions",  "chapters",  "metadata" };

[Exposed=Window]
interface TextTrack : EventTarget {
  readonly attribute TextTrackKind kind;
  readonly attribute DOMString label;
  readonly attribute DOMString language;

  readonly attribute DOMString id;
  readonly attribute DOMString inBandMetadataTrackDispatchType;

  attribute TextTrackMode mode;

  readonly attribute TextTrackCueList? cues;
  readonly attribute TextTrackCueList? activeCues;

  undefined addCue(TextTrackCue cue);
  undefined removeCue(TextTrackCue cue);

  attribute EventHandler oncuechange;
};
```

`textTrack`{.variable}` = ``media`{.variable}`.`[`addTextTrack`](#dom-media-addtexttrack){#dom-media-addtexttrack-dev}`(``kind`{.variable}` [, ``label`{.variable}` [, ``language`{.variable}` ] ])`

Creates and returns a new
[`TextTrack`{#text-track-api:texttrack-7}](#texttrack) object, which is
also added to the [media
element](#media-element){#text-track-api:media-element-6}\'s [list of
text
tracks](#list-of-text-tracks){#text-track-api:list-of-text-tracks-5}.

`textTrack`{.variable}`.`[`kind`](#dom-texttrack-kind){#dom-texttrack-kind-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[TextTrack/kind](https://developer.mozilla.org/en-US/docs/Web/API/TextTrack/kind "The kind read-only property of the TextTrack interface returns the kind of text track this object represents. This decides how the track will be handled by a user agent.")

Support in all current engines.

::: support
[Firefox31+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome23+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS7+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns the [text track
kind](#text-track-kind){#text-track-api:text-track-kind} string.

`textTrack`{.variable}`.`[`label`](#dom-texttrack-label){#dom-texttrack-label-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[TextTrack/label](https://developer.mozilla.org/en-US/docs/Web/API/TextTrack/label "The label read-only property of the TextTrack interface returns a human-readable label for the text track, if it is available.")

Support in all current engines.

::: support
[Firefox31+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome23+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS7+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns the [text track
label](#text-track-label){#text-track-api:text-track-label}, if there is
one, or the empty string otherwise (indicating that a custom label
probably needs to be generated from the other attributes of the object
if the object is exposed to the user).

`textTrack`{.variable}`.`[`language`](#dom-texttrack-language){#dom-texttrack-language-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[TextTrack/language](https://developer.mozilla.org/en-US/docs/Web/API/TextTrack/language "The language read-only property of the TextTrack interface returns the language of the text track.")

Support in all current engines.

::: support
[Firefox31+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome23+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS7+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns the [text track
language](#text-track-language){#text-track-api:text-track-language}
string.

`textTrack`{.variable}`.`[`id`](#dom-texttrack-id){#dom-texttrack-id-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[TextTrack/id](https://developer.mozilla.org/en-US/docs/Web/API/TextTrack/id "The id read-only property of the TextTrack interface returns the ID of the track if it has one.")

Support in all current engines.

::: support
[Firefox31+]{.firefox .yes}[Safari8+]{.safari .yes}[Chrome33+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)18]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

Returns the ID of the given track.

For in-band tracks, this is the ID that can be used with a
[fragment](https://url.spec.whatwg.org/#concept-url-fragment){#text-track-api:concept-url-fragment
x-internal="concept-url-fragment"} if the format supports [media
fragment
syntax](https://www.w3.org/TR/media-frags/#media-fragment-syntax){#text-track-api:media-fragment-syntax
x-internal="media-fragment-syntax"}, and that can be used with the
[`getTrackById()`{#text-track-api:dom-texttracklist-gettrackbyid-2}](#dom-texttracklist-gettrackbyid)
method.

For [`TextTrack`{#text-track-api:texttrack-8}](#texttrack) objects
corresponding to
[`track`{#text-track-api:the-track-element-2}](#the-track-element)
elements, this is the ID of the
[`track`{#text-track-api:the-track-element-3}](#the-track-element)
element.

`textTrack`{.variable}`.`[`inBandMetadataTrackDispatchType`](#dom-texttrack-inbandmetadatatrackdispatchtype){#dom-texttrack-inbandmetadatatrackdispatchtype-dev}

::::: {.mdn-anno .wrapped .before}
MDN

:::: feature
[TextTrack/inBandMetadataTrackDispatchType](https://developer.mozilla.org/en-US/docs/Web/API/TextTrack/inBandMetadataTrackDispatchType "The inBandMetadataTrackDispatchType read-only property of the TextTrack interface returns the text track's in-band metadata dispatch type of the text track represented by the TextTrack object.")

::: support
[Firefox31+]{.firefox .yes}[Safari8+]{.safari .yes}[ChromeNo]{.chrome
.no}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[EdgeNo]{.edge_blink .no}

------------------------------------------------------------------------

[Edge (Legacy)No]{.edge .no}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

Returns the [text track in-band metadata track dispatch
type](#text-track-in-band-metadata-track-dispatch-type){#text-track-api:text-track-in-band-metadata-track-dispatch-type}
string.

`textTrack`{.variable}`.`[`mode`](#dom-texttrack-mode){#dom-texttrack-mode-dev}` [ = ``value`{.variable}` ]`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[TextTrack/mode](https://developer.mozilla.org/en-US/docs/Web/API/TextTrack/mode "The TextTrack interface's mode property is a string specifying and controlling the text track's mode: disabled, hidden, or showing. You can read this value to determine the current mode, and you can change this value to switch modes.")

Support in all current engines.

::: support
[Firefox31+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome23+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android31+]{.firefox_android .yes}[Safari iOS7+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns the [text track
mode](#text-track-mode){#text-track-api:text-track-mode}, represented by
a string from the following list:

\"[`disabled`{#text-track-api:dom-texttrack-disabled-2}](#dom-texttrack-disabled)\"
:   The [text track
    disabled](#text-track-disabled){#text-track-api:text-track-disabled}
    mode.

\"[`hidden`{#text-track-api:dom-texttrack-hidden-2}](#dom-texttrack-hidden)\"
:   The [text track
    hidden](#text-track-hidden){#text-track-api:text-track-hidden} mode.

\"[`showing`{#text-track-api:dom-texttrack-showing-2}](#dom-texttrack-showing)\"

:   The [text track
    showing](#text-track-showing){#text-track-api:text-track-showing}
    mode.

Can be set, to change the mode.

`textTrack`{.variable}`.`[`cues`](#dom-texttrack-cues){#dom-texttrack-cues-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[TextTrack/cues](https://developer.mozilla.org/en-US/docs/Web/API/TextTrack/cues "The cues read-only property of the TextTrack interface returns a TextTrackCueList object containing all of the track's cues.")

Support in all current engines.

::: support
[Firefox31+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome23+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android31+]{.firefox_android .yes}[Safari iOS7+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns the [text track list of
cues](#text-track-list-of-cues){#text-track-api:text-track-list-of-cues},
as a
[`TextTrackCueList`{#text-track-api:texttrackcuelist-3}](#texttrackcuelist)
object.

`textTrack`{.variable}`.`[`activeCues`](#dom-texttrack-activecues){#dom-texttrack-activecues-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[TextTrack/activeCues](https://developer.mozilla.org/en-US/docs/Web/API/TextTrack/activeCues "The activeCues read-only property of the TextTrack interface returns a TextTrackCueList object listing the currently active cues.")

Support in all current engines.

::: support
[Firefox31+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome23+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android31+]{.firefox_android .yes}[Safari iOS7+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns the [text track
cues](#text-track-cue){#text-track-api:text-track-cue} from the [text
track list of
cues](#text-track-list-of-cues){#text-track-api:text-track-list-of-cues-2}
that are currently active (i.e. that start before the [current playback
position](#current-playback-position){#text-track-api:current-playback-position}
and end after it), as a
[`TextTrackCueList`{#text-track-api:texttrackcuelist-4}](#texttrackcuelist)
object.

`textTrack`{.variable}`.`[`addCue`](#dom-texttrack-addcue){#dom-texttrack-addcue-dev}`(``cue`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[TextTrack/addCue](https://developer.mozilla.org/en-US/docs/Web/API/TextTrack/addCue "The addCue() method of the TextTrack interface adds a new cue to the list of cues.")

Support in all current engines.

::: support
[Firefox31+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome23+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS7+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Adds the given cue to `textTrack`{.variable}\'s [text track list of
cues](#text-track-list-of-cues){#text-track-api:text-track-list-of-cues-3}.

`textTrack`{.variable}`.`[`removeCue`](#dom-texttrack-removecue){#dom-texttrack-removecue-dev}`(``cue`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[TextTrack/removeCue](https://developer.mozilla.org/en-US/docs/Web/API/TextTrack/removeCue "The removeCue() method of the TextTrack interface removes a cue from the list of cues.")

Support in all current engines.

::: support
[Firefox31+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome23+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS7+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Removes the given cue from `textTrack`{.variable}\'s [text track list of
cues](#text-track-list-of-cues){#text-track-api:text-track-list-of-cues-4}.

The
[`addTextTrack(``kind`{.variable}`, ``label`{.variable}`, ``language`{.variable}`)`]{#dom-media-addtexttrack
.dfn dfn-for="HTMLMediaElement" dfn-type="method"} method of [media
elements](#media-element){#text-track-api:media-element-7}, when
invoked, must run the following steps:

1.  Create a new [`TextTrack`{#text-track-api:texttrack-9}](#texttrack)
    object.

2.  Create a new [text
    track](#text-track){#text-track-api:text-track-10} corresponding to
    the new object, and set its [text track
    kind](#text-track-kind){#text-track-api:text-track-kind-2} to
    `kind`{.variable}, its [text track
    label](#text-track-label){#text-track-api:text-track-label-2} to
    `label`{.variable}, its [text track
    language](#text-track-language){#text-track-api:text-track-language-2}
    to `language`{.variable}, its [text track readiness
    state](#text-track-readiness-state){#text-track-api:text-track-readiness-state}
    to the [text track
    loaded](#text-track-loaded){#text-track-api:text-track-loaded}
    state, its [text track
    mode](#text-track-mode){#text-track-api:text-track-mode-2} to the
    [text track
    hidden](#text-track-hidden){#text-track-api:text-track-hidden-2}
    mode, and its [text track list of
    cues](#text-track-list-of-cues){#text-track-api:text-track-list-of-cues-5}
    to an empty list.

    Initially, the [text track list of
    cues](#text-track-list-of-cues){#text-track-api:text-track-list-of-cues-6}
    is not associated with any [rules for updating the text track
    rendering](#rules-for-updating-the-text-track-rendering){#text-track-api:rules-for-updating-the-text-track-rendering}.
    When a [text track
    cue](#text-track-cue){#text-track-api:text-track-cue-2} is added to
    it, the [text track list of
    cues](#text-track-list-of-cues){#text-track-api:text-track-list-of-cues-7}
    has its rules permanently set accordingly.

3.  Add the new [text track](#text-track){#text-track-api:text-track-11}
    to the [media
    element](#media-element){#text-track-api:media-element-8}\'s [list
    of text
    tracks](#list-of-text-tracks){#text-track-api:list-of-text-tracks-6}.

4.  [Queue a media element
    task](#queue-a-media-element-task){#text-track-api:queue-a-media-element-task}
    given the [media
    element](#media-element){#text-track-api:media-element-9} to [fire
    an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#text-track-api:concept-event-fire
    x-internal="concept-event-fire"} named
    [`addtrack`{#text-track-api:event-media-addtrack}](#event-media-addtrack)
    at the [media
    element](#media-element){#text-track-api:media-element-10}\'s
    [`textTracks`{#text-track-api:dom-media-texttracks-3}](#dom-media-texttracks)
    attribute\'s
    [`TextTrackList`{#text-track-api:texttracklist-10}](#texttracklist)
    object, using
    [`TrackEvent`{#text-track-api:trackevent}](#trackevent), with the
    [`track`{#text-track-api:dom-trackevent-track}](#dom-trackevent-track)
    attribute initialized to the new [text
    track](#text-track){#text-track-api:text-track-12}\'s
    [`TextTrack`{#text-track-api:texttrack-10}](#texttrack) object.

5.  Return the new
    [`TextTrack`{#text-track-api:texttrack-11}](#texttrack) object.

------------------------------------------------------------------------

The [`kind`]{#dom-texttrack-kind .dfn dfn-for="TextTrack"
dfn-type="attribute"} attribute must return the [text track
kind](#text-track-kind){#text-track-api:text-track-kind-3} of the [text
track](#text-track){#text-track-api:text-track-13} that the
[`TextTrack`{#text-track-api:texttrack-12}](#texttrack) object
represents.

The [`label`]{#dom-texttrack-label .dfn dfn-for="TextTrack"
dfn-type="attribute"} attribute must return the [text track
label](#text-track-label){#text-track-api:text-track-label-3} of the
[text track](#text-track){#text-track-api:text-track-14} that the
[`TextTrack`{#text-track-api:texttrack-13}](#texttrack) object
represents.

The [`language`]{#dom-texttrack-language .dfn dfn-for="TextTrack"
dfn-type="attribute"} attribute must return the [text track
language](#text-track-language){#text-track-api:text-track-language-3}
of the [text track](#text-track){#text-track-api:text-track-15} that the
[`TextTrack`{#text-track-api:texttrack-14}](#texttrack) object
represents.

The [`id`]{#dom-texttrack-id .dfn dfn-for="TextTrack"
dfn-type="attribute"} attribute returns the track\'s identifier, if it
has one, or the empty string otherwise. For tracks that correspond to
[`track`{#text-track-api:the-track-element-4}](#the-track-element)
elements, the track\'s identifier is the value of the element\'s
[`id`{#text-track-api:the-id-attribute}](dom.html#the-id-attribute)
attribute, if any. For in-band tracks, the track\'s identifier is
specified by the [media
resource](#media-resource){#text-track-api:media-resource}. If the
[media resource](#media-resource){#text-track-api:media-resource-2} is
in a format that supports [media fragment
syntax](https://www.w3.org/TR/media-frags/#media-fragment-syntax){#text-track-api:media-fragment-syntax-2
x-internal="media-fragment-syntax"}, the identifier returned for a
particular track must be the same identifier that would enable the track
if used as the name of a track in the track dimension of such a
[fragment](https://url.spec.whatwg.org/#concept-url-fragment){#text-track-api:concept-url-fragment-2
x-internal="concept-url-fragment"}.

The
[`inBandMetadataTrackDispatchType`]{#dom-texttrack-inbandmetadatatrackdispatchtype
.dfn dfn-for="TextTrack" dfn-type="attribute"} attribute must return the
[text track in-band metadata track dispatch
type](#text-track-in-band-metadata-track-dispatch-type){#text-track-api:text-track-in-band-metadata-track-dispatch-type-2}
of the [text track](#text-track){#text-track-api:text-track-16} that the
[`TextTrack`{#text-track-api:texttrack-15}](#texttrack) object
represents.

The [`mode`]{#dom-texttrack-mode .dfn dfn-for="TextTrack"
dfn-type="attribute"} attribute, on getting, must return the string
corresponding to the [text track
mode](#text-track-mode){#text-track-api:text-track-mode-3} of the [text
track](#text-track){#text-track-api:text-track-17} that the
[`TextTrack`{#text-track-api:texttrack-16}](#texttrack) object
represents, as defined by the following list:

\"[`disabled`]{#dom-texttrack-disabled .dfn dfn-for="TextTrackMode" dfn-type="enum-value"}\"
:   The [text track
    disabled](#text-track-disabled){#text-track-api:text-track-disabled-2}
    mode.

\"[`hidden`]{#dom-texttrack-hidden .dfn dfn-for="TextTrackMode" dfn-type="enum-value"}\"
:   The [text track
    hidden](#text-track-hidden){#text-track-api:text-track-hidden-3}
    mode.

\"[`showing`]{#dom-texttrack-showing .dfn dfn-for="TextTrackMode" dfn-type="enum-value"}\"
:   The [text track
    showing](#text-track-showing){#text-track-api:text-track-showing-2}
    mode.

On setting, if the new value isn\'t equal to what the attribute would
currently return, the new value must be processed as follows:

If the new value is \"[`disabled`{#text-track-api:dom-texttrack-disabled-3}](#dom-texttrack-disabled)\"

:   Set the [text track
    mode](#text-track-mode){#text-track-api:text-track-mode-4} of the
    [text track](#text-track){#text-track-api:text-track-18} that the
    [`TextTrack`{#text-track-api:texttrack-17}](#texttrack) object
    represents to the [text track
    disabled](#text-track-disabled){#text-track-api:text-track-disabled-3}
    mode.

If the new value is \"[`hidden`{#text-track-api:dom-texttrack-hidden-3}](#dom-texttrack-hidden)\"

:   Set the [text track
    mode](#text-track-mode){#text-track-api:text-track-mode-5} of the
    [text track](#text-track){#text-track-api:text-track-19} that the
    [`TextTrack`{#text-track-api:texttrack-18}](#texttrack) object
    represents to the [text track
    hidden](#text-track-hidden){#text-track-api:text-track-hidden-4}
    mode.

If the new value is \"[`showing`{#text-track-api:dom-texttrack-showing-3}](#dom-texttrack-showing)\"

:   Set the [text track
    mode](#text-track-mode){#text-track-api:text-track-mode-6} of the
    [text track](#text-track){#text-track-api:text-track-20} that the
    [`TextTrack`{#text-track-api:texttrack-19}](#texttrack) object
    represents to the [text track
    showing](#text-track-showing){#text-track-api:text-track-showing-3}
    mode.

If the [text track
mode](#text-track-mode){#text-track-api:text-track-mode-7} of the [text
track](#text-track){#text-track-api:text-track-21} that the
[`TextTrack`{#text-track-api:texttrack-20}](#texttrack) object
represents is not the [text track
disabled](#text-track-disabled){#text-track-api:text-track-disabled-4}
mode, then the [`cues`]{#dom-texttrack-cues .dfn dfn-for="TextTrack"
dfn-type="attribute"} attribute must return a
[live](infrastructure.html#live){#text-track-api:live}
[`TextTrackCueList`{#text-track-api:texttrackcuelist-5}](#texttrackcuelist)
object that represents the subset of the [text track list of
cues](#text-track-list-of-cues){#text-track-api:text-track-list-of-cues-8}
of the [text track](#text-track){#text-track-api:text-track-22} that the
[`TextTrack`{#text-track-api:texttrack-21}](#texttrack) object
represents whose [end
times](#text-track-cue-end-time){#text-track-api:text-track-cue-end-time}
occur at or after the [earliest possible position when the script
started](#earliest-possible-position-when-the-script-started){#text-track-api:earliest-possible-position-when-the-script-started},
in [text track cue
order](#text-track-cue-order){#text-track-api:text-track-cue-order}.
Otherwise, it must return null. For each
[`TextTrack`{#text-track-api:texttrack-22}](#texttrack) object, when an
object is returned, the same
[`TextTrackCueList`{#text-track-api:texttrackcuelist-6}](#texttrackcuelist)
object must be returned each time.

The [earliest possible position when the script
started]{#earliest-possible-position-when-the-script-started .dfn} is
whatever the [earliest possible
position](#earliest-possible-position){#text-track-api:earliest-possible-position}
was the last time the [event
loop](webappapis.html#event-loop){#text-track-api:event-loop} reached
step 1.

If the [text track
mode](#text-track-mode){#text-track-api:text-track-mode-8} of the [text
track](#text-track){#text-track-api:text-track-23} that the
[`TextTrack`{#text-track-api:texttrack-23}](#texttrack) object
represents is not the [text track
disabled](#text-track-disabled){#text-track-api:text-track-disabled-5}
mode, then the [`activeCues`]{#dom-texttrack-activecues .dfn
dfn-for="TextTrack" dfn-type="attribute"} attribute must return a
[live](infrastructure.html#live){#text-track-api:live-2}
[`TextTrackCueList`{#text-track-api:texttrackcuelist-7}](#texttrackcuelist)
object that represents the subset of the [text track list of
cues](#text-track-list-of-cues){#text-track-api:text-track-list-of-cues-9}
of the [text track](#text-track){#text-track-api:text-track-24} that the
[`TextTrack`{#text-track-api:texttrack-24}](#texttrack) object
represents whose [active flag was set when the script
started](#active-flag-was-set-when-the-script-started){#text-track-api:active-flag-was-set-when-the-script-started},
in [text track cue
order](#text-track-cue-order){#text-track-api:text-track-cue-order-2}.
Otherwise, it must return null. For each
[`TextTrack`{#text-track-api:texttrack-25}](#texttrack) object, when an
object is returned, the same
[`TextTrackCueList`{#text-track-api:texttrackcuelist-8}](#texttrackcuelist)
object must be returned each time.

A [text track cue](#text-track-cue){#text-track-api:text-track-cue-3}\'s
[active flag was set when the script
started]{#active-flag-was-set-when-the-script-started .dfn} if its [text
track cue active
flag](#text-track-cue-active-flag){#text-track-api:text-track-cue-active-flag}
was set the last time the [event
loop](webappapis.html#event-loop){#text-track-api:event-loop-2} reached
[step 1](webappapis.html#step1).

------------------------------------------------------------------------

The [`addCue(``cue`{.variable}`)`]{#dom-texttrack-addcue .dfn
dfn-for="TextTrack" dfn-type="method"} method of
[`TextTrack`{#text-track-api:texttrack-26}](#texttrack) objects, when
invoked, must run the following steps:

1.  If the [text track list of
    cues](#text-track-list-of-cues){#text-track-api:text-track-list-of-cues-10}
    does not yet have any associated [rules for updating the text track
    rendering](#rules-for-updating-the-text-track-rendering){#text-track-api:rules-for-updating-the-text-track-rendering-2},
    then associate the [text track list of
    cues](#text-track-list-of-cues){#text-track-api:text-track-list-of-cues-11}
    with the [rules for updating the text track
    rendering](#rules-for-updating-the-text-track-rendering){#text-track-api:rules-for-updating-the-text-track-rendering-3}
    appropriate to `cue`{.variable}.

2.  If [text track list of
    cues](#text-track-list-of-cues){#text-track-api:text-track-list-of-cues-12}\'
    associated [rules for updating the text track
    rendering](#rules-for-updating-the-text-track-rendering){#text-track-api:rules-for-updating-the-text-track-rendering-4}
    are not the same [rules for updating the text track
    rendering](#rules-for-updating-the-text-track-rendering){#text-track-api:rules-for-updating-the-text-track-rendering-5}
    as appropriate for `cue`{.variable}, then throw an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#text-track-api:invalidstateerror
    x-internal="invalidstateerror"}
    [`DOMException`{#text-track-api:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

3.  If the given `cue`{.variable} is in a [text track list of
    cues](#text-track-list-of-cues){#text-track-api:text-track-list-of-cues-13},
    then remove `cue`{.variable} from that [text track list of
    cues](#text-track-list-of-cues){#text-track-api:text-track-list-of-cues-14}.

4.  Add `cue`{.variable} to the
    [`TextTrack`{#text-track-api:texttrack-27}](#texttrack) object\'s
    [text track](#text-track){#text-track-api:text-track-25}\'s [text
    track list of
    cues](#text-track-list-of-cues){#text-track-api:text-track-list-of-cues-15}.

The [`removeCue(``cue`{.variable}`)`]{#dom-texttrack-removecue .dfn
dfn-for="TextTrack" dfn-type="method"} method of
[`TextTrack`{#text-track-api:texttrack-28}](#texttrack) objects, when
invoked, must run the following steps:

1.  If the given `cue`{.variable} is not in the
    [`TextTrack`{#text-track-api:texttrack-29}](#texttrack) object\'s
    [text track](#text-track){#text-track-api:text-track-26}\'s [text
    track list of
    cues](#text-track-list-of-cues){#text-track-api:text-track-list-of-cues-16},
    then throw a
    [\"`NotFoundError`\"](https://webidl.spec.whatwg.org/#notfounderror){#text-track-api:notfounderror
    x-internal="notfounderror"}
    [`DOMException`{#text-track-api:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  Remove `cue`{.variable} from the
    [`TextTrack`{#text-track-api:texttrack-30}](#texttrack) object\'s
    [text track](#text-track){#text-track-api:text-track-27}\'s [text
    track list of
    cues](#text-track-list-of-cues){#text-track-api:text-track-list-of-cues-17}.

::: example
In this example, an
[`audio`{#text-track-api:the-audio-element}](#the-audio-element) element
is used to play a specific sound-effect from a sound file containing
many sound effects. A cue is used to pause the audio, so that it ends
exactly at the end of the clip, even if the browser is busy running some
script. If the page had relied on script to pause the audio, then the
start of the next clip might be heard if the browser was not able to run
the script at the exact time specified.

``` js
var sfx = new Audio('sfx.wav');
var sounds = sfx.addTextTrack('metadata');

// add sounds we care about
function addFX(start, end, name) {
  var cue = new VTTCue(start, end, '');
  cue.id = name;
  cue.pauseOnExit = true;
  sounds.addCue(cue);
}
addFX(12.783, 13.612, 'dog bark');
addFX(13.612, 15.091, 'kitten mew');

function playSound(id) {
  sfx.currentTime = sounds.getCueById(id).startTime;
  sfx.play();
}

// play a bark as soon as we can
sfx.oncanplaythrough = function () {
  playSound('dog bark');
}
// meow when the user tries to leave,
// and have the browser ask them to stay
window.onbeforeunload = function (e) {
  playSound('kitten mew');
  e.preventDefault();
}
```
:::

------------------------------------------------------------------------

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[TextTrackCueList](https://developer.mozilla.org/en-US/docs/Web/API/TextTrackCueList "The TextTrackCueList array-like object represents a dynamically updating list of TextTrackCue objects.")

Support in all current engines.

::: support
[Firefox31+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome23+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

``` idl
[Exposed=Window]
interface TextTrackCueList {
  readonly attribute unsigned long length;
  getter TextTrackCue (unsigned long index);
  TextTrackCue? getCueById(DOMString id);
};
```

`cuelist`{.variable}`.`[`length`](#dom-texttrackcuelist-length){#dom-texttrackcuelist-length-dev}
:   Returns the number of
    [cues](#text-track-cue){#text-track-api:text-track-cue-4} in the
    list.

`cuelist`{.variable}`[``index`{.variable}`]`

:   Returns the [text track
    cue](#text-track-cue){#text-track-api:text-track-cue-5} with index
    `index`{.variable} in the list. The cues are sorted in [text track
    cue
    order](#text-track-cue-order){#text-track-api:text-track-cue-order-3}.

`cuelist`{.variable}`.`[`getCueById`](#dom-texttrackcuelist-getcuebyid){#dom-texttrackcuelist-getcuebyid-dev}`(``id`{.variable}`)`

:   Returns the first [text track
    cue](#text-track-cue){#text-track-api:text-track-cue-6} (in [text
    track cue
    order](#text-track-cue-order){#text-track-api:text-track-cue-order-4})
    with [text track cue
    identifier](#text-track-cue-identifier){#text-track-api:text-track-cue-identifier}
    `id`{.variable}.

    Returns null if none of the cues have the given identifier or if the
    argument is the empty string.

A
[`TextTrackCueList`{#text-track-api:texttrackcuelist-9}](#texttrackcuelist)
object represents a dynamically updating list of [text track
cues](#text-track-cue){#text-track-api:text-track-cue-7} in a given
order.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[TextTrackCueList/length](https://developer.mozilla.org/en-US/docs/Web/API/TextTrackCueList/length "The length read-only property of the TextTrackCueList interface returns the number of cues in the list.")

Support in all current engines.

::: support
[Firefox31+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome23+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The [`length`]{#dom-texttrackcuelist-length .dfn
dfn-for="TextTrackCueList" dfn-type="attribute"} attribute must return
the number of [cues](#text-track-cue){#text-track-api:text-track-cue-8}
in the list represented by the
[`TextTrackCueList`{#text-track-api:texttrackcuelist-10}](#texttrackcuelist)
object.

The [supported property
indices](https://webidl.spec.whatwg.org/#dfn-supported-property-indices){#text-track-api:supported-property-indices-3
x-internal="supported-property-indices"} of a
[`TextTrackCueList`{#text-track-api:texttrackcuelist-11}](#texttrackcuelist)
object at any instant are the numbers from zero to the number of
[cues](#text-track-cue){#text-track-api:text-track-cue-9} in the list
represented by the
[`TextTrackCueList`{#text-track-api:texttrackcuelist-12}](#texttrackcuelist)
object minus one, if any. If there are no
[cues](#text-track-cue){#text-track-api:text-track-cue-10} in the list,
there are no [supported property
indices](https://webidl.spec.whatwg.org/#dfn-supported-property-indices){#text-track-api:supported-property-indices-4
x-internal="supported-property-indices"}.

To [determine the value of an indexed
property](https://webidl.spec.whatwg.org/#dfn-determine-the-value-of-an-indexed-property){#text-track-api:determine-the-value-of-an-indexed-property-2
x-internal="determine-the-value-of-an-indexed-property"} for a given
index `index`{.variable}, the user agent must return the
`index`{.variable}th [text track
cue](#text-track-cue){#text-track-api:text-track-cue-11} in the list
represented by the
[`TextTrackCueList`{#text-track-api:texttrackcuelist-13}](#texttrackcuelist)
object.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[TextTrackCueList/getCueById](https://developer.mozilla.org/en-US/docs/Web/API/TextTrackCueList/getCueById "The getCueById() method of the TextTrackCueList interface returns the first VTTCue in the list represented by the TextTrackCueList object whose identifier matches the value of id.")

Support in all current engines.

::: support
[Firefox31+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome23+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The [`getCueById(``id`{.variable}`)`]{#dom-texttrackcuelist-getcuebyid
.dfn dfn-for="TextTrackCueList" dfn-type="method"} method, when called
with an argument other than the empty string, must return the first
[text track cue](#text-track-cue){#text-track-api:text-track-cue-12} in
the list represented by the
[`TextTrackCueList`{#text-track-api:texttrackcuelist-14}](#texttrackcuelist)
object whose [text track cue
identifier](#text-track-cue-identifier){#text-track-api:text-track-cue-identifier-2}
is `id`{.variable}, if any, or null otherwise. If the argument is the
empty string, then the method must return null.

------------------------------------------------------------------------

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[TextTrackCue](https://developer.mozilla.org/en-US/docs/Web/API/TextTrackCue "TextTrackCue is an abstract class which is used as the basis for the various derived cue types, such as VTTCue; you will instead work with those derived types. These cues represent strings of text presented for some duration of time during the performance of a TextTrack. The cue includes the start time (the time at which the text will be displayed) and the end time (the time at which it will be removed from the display), as well as other information.")

Support in all current engines.

::: support
[Firefox31+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome23+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS7+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

``` idl
[Exposed=Window]
interface TextTrackCue : EventTarget {
  readonly attribute TextTrack? track;

  attribute DOMString id;
  attribute double startTime;
  attribute unrestricted double endTime;
  attribute boolean pauseOnExit;

  attribute EventHandler onenter;
  attribute EventHandler onexit;
};
```

`cue`{.variable}`.`[`track`](#dom-texttrackcue-track){#dom-texttrackcue-track-dev}

:   Returns the [`TextTrack`{#text-track-api:texttrack-32}](#texttrack)
    object to which this [text track
    cue](#text-track-cue){#text-track-api:text-track-cue-13} belongs, if
    any, or null otherwise.

`cue`{.variable}`.`[`id`](#dom-texttrackcue-id){#dom-texttrackcue-id-dev}` [ = ``value`{.variable}` ]`

:   Returns the [text track cue
    identifier](#text-track-cue-identifier){#text-track-api:text-track-cue-identifier-3}.

    Can be set.

`cue`{.variable}`.`[`startTime`](#dom-texttrackcue-starttime){#dom-texttrackcue-starttime-dev}` [ = ``value`{.variable}` ]`

:   Returns the [text track cue start
    time](#text-track-cue-start-time){#text-track-api:text-track-cue-start-time},
    in seconds.

    Can be set.

`cue`{.variable}`.`[`endTime`](#dom-texttrackcue-endtime){#dom-texttrackcue-endtime-dev}` [ = ``value`{.variable}` ]`

:   Returns the [text track cue end
    time](#text-track-cue-end-time){#text-track-api:text-track-cue-end-time-2},
    in seconds.

    Returns positive Infinity for an [unbounded text track
    cue](#unbounded-text-track-cue){#text-track-api:unbounded-text-track-cue}.

    Can be set.

`cue`{.variable}`.`[`pauseOnExit`](#dom-texttrackcue-pauseonexit){#dom-texttrackcue-pauseonexit-dev}` [ = ``value`{.variable}` ]`

:   Returns true if the [text track cue pause-on-exit
    flag](#text-track-cue-pause-on-exit-flag){#text-track-api:text-track-cue-pause-on-exit-flag}
    is set, false otherwise.

    Can be set.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[TextTrackCue/track](https://developer.mozilla.org/en-US/docs/Web/API/TextTrackCue/track "The track read-only property of the TextTrackCue interface returns the TextTrack object that this cue belongs to.")

Support in all current engines.

::: support
[Firefox31+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome23+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS7+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The [`track`]{#dom-texttrackcue-track .dfn dfn-for="TextTrackCue"
dfn-type="attribute"} attribute, on getting, must return the
[`TextTrack`{#text-track-api:texttrack-33}](#texttrack) object of the
[text track](#text-track){#text-track-api:text-track-28} in whose [list
of
cues](#text-track-list-of-cues){#text-track-api:text-track-list-of-cues-18}
the [text track cue](#text-track-cue){#text-track-api:text-track-cue-14}
that the [`TextTrackCue`{#text-track-api:texttrackcue-5}](#texttrackcue)
object represents finds itself, if any; or null otherwise.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[TextTrackCue/id](https://developer.mozilla.org/en-US/docs/Web/API/TextTrackCue/id "The id property of the TextTrackCue interface returns and sets the identifier for this cue.")

Support in all current engines.

::: support
[Firefox31+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome23+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS7+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The [`id`]{#dom-texttrackcue-id .dfn dfn-for="TextTrackCue"
dfn-type="attribute"} attribute, on getting, must return the [text track
cue
identifier](#text-track-cue-identifier){#text-track-api:text-track-cue-identifier-4}
of the [text track
cue](#text-track-cue){#text-track-api:text-track-cue-15} that the
[`TextTrackCue`{#text-track-api:texttrackcue-6}](#texttrackcue) object
represents. On setting, the [text track cue
identifier](#text-track-cue-identifier){#text-track-api:text-track-cue-identifier-5}
must be set to the new value.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[TextTrackCue/startTime](https://developer.mozilla.org/en-US/docs/Web/API/TextTrackCue/startTime "The startTime property of the TextTrackCue interface returns and sets the start time of the cue.")

Support in all current engines.

::: support
[Firefox31+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome23+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS7+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The [`startTime`]{#dom-texttrackcue-starttime .dfn
dfn-for="TextTrackCue" dfn-type="attribute"} attribute, on getting, must
return the [text track cue start
time](#text-track-cue-start-time){#text-track-api:text-track-cue-start-time-2}
of the [text track
cue](#text-track-cue){#text-track-api:text-track-cue-16} that the
[`TextTrackCue`{#text-track-api:texttrackcue-7}](#texttrackcue) object
represents, in seconds. On setting, the [text track cue start
time](#text-track-cue-start-time){#text-track-api:text-track-cue-start-time-3}
must be set to the new value, interpreted in seconds; then, if the
[`TextTrackCue`{#text-track-api:texttrackcue-8}](#texttrackcue)
object\'s [text track
cue](#text-track-cue){#text-track-api:text-track-cue-17} is in a [text
track](#text-track){#text-track-api:text-track-29}\'s [list of
cues](#text-track-list-of-cues){#text-track-api:text-track-list-of-cues-19},
and that [text track](#text-track){#text-track-api:text-track-30} is in
a [media element](#media-element){#text-track-api:media-element-11}\'s
[list of text
tracks](#list-of-text-tracks){#text-track-api:list-of-text-tracks-7},
and the [media
element](#media-element){#text-track-api:media-element-12}\'s [show
poster flag](#show-poster-flag){#text-track-api:show-poster-flag} is not
set, then run the *[time marches on](#time-marches-on)* steps for that
[media element](#media-element){#text-track-api:media-element-13}.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[TextTrackCue/endTime](https://developer.mozilla.org/en-US/docs/Web/API/TextTrackCue/endTime "The endTime property of the TextTrackCue interface returns and sets the end time of the cue.")

Support in all current engines.

::: support
[Firefox31+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome23+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS7+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The [`endTime`]{#dom-texttrackcue-endtime .dfn dfn-for="TextTrackCue"
dfn-type="attribute"} attribute, on getting, must return the [text track
cue end
time](#text-track-cue-end-time){#text-track-api:text-track-cue-end-time-3}
of the [text track
cue](#text-track-cue){#text-track-api:text-track-cue-18} that the
[`TextTrackCue`{#text-track-api:texttrackcue-9}](#texttrackcue) object
represents, in seconds or positive Infinity. On setting, if the new
value is negative Infinity or a Not-a-Number (NaN) value, then throw a
[TypeError](https://tc39.es/ecma262/#sec-native-error-types-used-in-this-standard-typeerror)
exception. Otherwise, the [text track cue end
time](#text-track-cue-end-time){#text-track-api:text-track-cue-end-time-4}
must be set to the new value. Then, if the
[`TextTrackCue`{#text-track-api:texttrackcue-10}](#texttrackcue)
object\'s [text track
cue](#text-track-cue){#text-track-api:text-track-cue-19} is in a [text
track](#text-track){#text-track-api:text-track-31}\'s [list of
cues](#text-track-list-of-cues){#text-track-api:text-track-list-of-cues-20},
and that [text track](#text-track){#text-track-api:text-track-32} is in
a [media element](#media-element){#text-track-api:media-element-14}\'s
[list of text
tracks](#list-of-text-tracks){#text-track-api:list-of-text-tracks-8},
and the [media
element](#media-element){#text-track-api:media-element-15}\'s [show
poster flag](#show-poster-flag){#text-track-api:show-poster-flag-2} is
not set, then run the *[time marches on](#time-marches-on)* steps for
that [media element](#media-element){#text-track-api:media-element-16}.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[TextTrackCue/pauseOnExit](https://developer.mozilla.org/en-US/docs/Web/API/TextTrackCue/pauseOnExit "The pauseOnExit property of the TextTrackCue interface returns or sets the flag indicating whether playback of the media should pause when the end of the range to which this cue applies is reached.")

Support in all current engines.

::: support
[Firefox31+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome23+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS7+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The [`pauseOnExit`]{#dom-texttrackcue-pauseonexit .dfn
dfn-for="TextTrackCue" dfn-type="attribute"} attribute, on getting, must
return true if the [text track cue pause-on-exit
flag](#text-track-cue-pause-on-exit-flag){#text-track-api:text-track-cue-pause-on-exit-flag-2}
of the [text track
cue](#text-track-cue){#text-track-api:text-track-cue-20} that the
[`TextTrackCue`{#text-track-api:texttrackcue-11}](#texttrackcue) object
represents is set; or false otherwise. On setting, the [text track cue
pause-on-exit
flag](#text-track-cue-pause-on-exit-flag){#text-track-api:text-track-cue-pause-on-exit-flag-3}
must be set if the new value is true, and must be unset otherwise.

###### [4.8.11.11.6]{.secno} Event handlers for objects of the text track APIs[](#cue-events){.self-link} {#cue-events}

The following are the [event
handlers](webappapis.html#event-handlers){#cue-events:event-handlers}
that (and their corresponding [event handler event
types](webappapis.html#event-handler-event-type){#cue-events:event-handler-event-type})
that must be supported, as [event handler IDL
attributes](webappapis.html#event-handler-idl-attributes){#cue-events:event-handler-idl-attributes},
by all objects implementing the
[`TextTrackList`{#cue-events:texttracklist}](#texttracklist) interface:

[Event
handler](webappapis.html#event-handlers){#cue-events:event-handlers-2}

[Event handler event
type](webappapis.html#event-handler-event-type){#cue-events:event-handler-event-type-2}

[`onchange`]{#handler-texttracklist-onchange .dfn
dfn-for="TextTrackList" dfn-type="attribute"}

[`change`{#cue-events:event-media-change}](#event-media-change)

[`onaddtrack`]{#handler-texttracklist-onaddtrack .dfn
dfn-for="TextTrackList" dfn-type="attribute"}

[`addtrack`{#cue-events:event-media-addtrack}](#event-media-addtrack)

[`onremovetrack`]{#handler-texttracklist-onremovetrack .dfn
dfn-for="TextTrackList" dfn-type="attribute"}

[`removetrack`{#cue-events:event-media-removetrack}](#event-media-removetrack)

The following are the [event
handlers](webappapis.html#event-handlers){#cue-events:event-handlers-3}
that (and their corresponding [event handler event
types](webappapis.html#event-handler-event-type){#cue-events:event-handler-event-type-3})
that must be supported, as [event handler IDL
attributes](webappapis.html#event-handler-idl-attributes){#cue-events:event-handler-idl-attributes-2},
by all objects implementing the
[`TextTrack`{#cue-events:texttrack}](#texttrack) interface:

[Event
handler](webappapis.html#event-handlers){#cue-events:event-handlers-4}

[Event handler event
type](webappapis.html#event-handler-event-type){#cue-events:event-handler-event-type-4}

[`oncuechange`]{#handler-texttrack-oncuechange .dfn dfn-for="TextTrack"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[TextTrack/cuechange_event](https://developer.mozilla.org/en-US/docs/Web/API/TextTrack/cuechange_event "The cuechange event fires when a TextTrack has changed the currently displaying cues. The event is fired on both the TextTrack and the HTMLTrackElement in which it's being presented, if any.")

Support in all current engines.

::: support
[Firefox31+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome23+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS7+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

[`cuechange`{#cue-events:event-media-cuechange}](#event-media-cuechange)

The following are the [event
handlers](webappapis.html#event-handlers){#cue-events:event-handlers-5}
(and their corresponding [event handler event
types](webappapis.html#event-handler-event-type){#cue-events:event-handler-event-type-5})
that must be supported, as [event handler IDL
attributes](webappapis.html#event-handler-idl-attributes){#cue-events:event-handler-idl-attributes-3},
by all objects implementing the
[`TextTrackCue`{#cue-events:texttrackcue}](#texttrackcue) interface:

[Event
handler](webappapis.html#event-handlers){#cue-events:event-handlers-6}

[Event handler event
type](webappapis.html#event-handler-event-type){#cue-events:event-handler-event-type-6}

[`onenter`]{#handler-texttrackcue-onenter .dfn dfn-for="TextTrackCue"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[TextTrackCue/enter_event](https://developer.mozilla.org/en-US/docs/Web/API/TextTrackCue/enter_event "The enter event fires when a cue becomes active. In the case of subtitles or a caption this is when it displays over the media.")

Support in all current engines.

::: support
[Firefox31+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome23+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS7+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

[`enter`{#cue-events:event-media-enter}](#event-media-enter)

[`onexit`]{#handler-texttrackcue-onexit .dfn dfn-for="TextTrackCue"
dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[TextTrackCue/exit_event](https://developer.mozilla.org/en-US/docs/Web/API/TextTrackCue/exit_event "The exit event fires when a cue stops being active.")

Support in all current engines.

::: support
[Firefox31+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome23+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS7+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

[`exit`{#cue-events:event-media-exit}](#event-media-exit)

###### [4.8.11.11.7]{.secno} Best practices for metadata text tracks[](#best-practices-for-metadata-text-tracks){.self-link}

*This section is non-normative.*

Text tracks can be used for storing data relating to the media data, for
interactive or augmented views.

For example, a page showing a sports broadcast could include information
about the current score. Suppose a robotics competition was being
streamed live. The image could be overlaid with the scores, as follows:

In order to make the score display render correctly whenever the user
seeks to an arbitrary point in the video, the metadata text track cues
need to be as long as is appropriate for the score. For example, in the
frame above, there would be maybe one cue that lasts the length of the
match that gives the match number, one cue that lasts until the blue
alliance\'s score changes, and one cue that lasts until the red
alliance\'s score changes. If the video is just a stream of the live
event, the time in the bottom right would presumably be automatically
derived from the current video time, rather than based on a cue.
However, if the video was just the highlights, then that might be given
in cues also.

The following shows what fragments of this could look like in a WebVTT
file:

    WEBVTT

    ...

    05:10:00.000 --> 05:12:15.000
    matchtype:qual
    matchnumber:37

    ...

    05:11:02.251 --> 05:11:17.198
    red:78

    05:11:03.672 --> 05:11:54.198
    blue:66

    05:11:17.198 --> 05:11:25.912
    red:80

    05:11:25.912 --> 05:11:26.522
    red:83

    05:11:26.522 --> 05:11:26.982
    red:86

    05:11:26.982 --> 05:11:27.499
    red:89

    ...

The key here is to notice that the information is given in cues that
span the length of time to which the relevant event applies. If,
instead, the scores were given as zero-length (or very brief, nearly
zero-length) cues when the score changes, for example saying \"red+2\"
at 05:11:17.198, \"red+3\" at 05:11:25.912, etc, problems arise:
primarily, seeking is much harder to implement, as the script has to
walk the entire list of cues to make sure that no notifications have
been missed; but also, if the cues are short it\'s possible the script
will never see that they are active unless it listens to them
specifically.

When using cues in this manner, authors are encouraged to use the
[`cuechange`{#best-practices-for-metadata-text-tracks:event-media-cuechange}](#event-media-cuechange)
event to update the current annotations. (In particular, using the
[`timeupdate`{#best-practices-for-metadata-text-tracks:event-media-timeupdate}](#event-media-timeupdate)
event would be less appropriate as it would require doing work even when
the cues haven\'t changed, and, more importantly, would introduce a
higher latency between when the metadata cues become active and when the
display is updated, since
[`timeupdate`{#best-practices-for-metadata-text-tracks:event-media-timeupdate-2}](#event-media-timeupdate)
events are rate-limited.)

##### [4.8.11.12]{.secno} Identifying a track kind through a URL[](#identifying-a-track-kind-through-a-url){.self-link}

Other specifications or formats that need a
[URL](https://url.spec.whatwg.org/#concept-url){#identifying-a-track-kind-through-a-url:url
x-internal="url"} to identify the return values of the
[`AudioTrack`{#identifying-a-track-kind-through-a-url:audiotrack}](#audiotrack)
[`kind`{#identifying-a-track-kind-through-a-url:dom-audiotrack-kind}](#dom-audiotrack-kind)
or
[`VideoTrack`{#identifying-a-track-kind-through-a-url:videotrack}](#videotrack)
[`kind`{#identifying-a-track-kind-through-a-url:dom-videotrack-kind}](#dom-videotrack-kind)
IDL attributes, or identify the [kind of text
track](#text-track-kind){#identifying-a-track-kind-through-a-url:text-track-kind},
must use the
[`about:html-kind`{#identifying-a-track-kind-through-a-url:about:html-kind}](urls-and-fetching.html#about:html-kind)
[URL](https://url.spec.whatwg.org/#concept-url){#identifying-a-track-kind-through-a-url:url-2
x-internal="url"}.

##### [4.8.11.13]{.secno} User interface[](#user-interface){.self-link}

The [`controls`]{#attr-media-controls .dfn dfn-for="audio,video"
dfn-type="element-attr"} attribute is a [boolean
attribute](common-microsyntaxes.html#boolean-attribute){#user-interface:boolean-attribute}.
If present, it indicates that the author has not provided a scripted
controller and would like the user agent to provide its own set of
controls.

If the attribute is present, or if [scripting is
disabled](webappapis.html#concept-n-noscript){#user-interface:concept-n-noscript}
for the [media element](#media-element){#user-interface:media-element},
then the user agent should [expose a user interface to the
user]{#expose-a-user-interface-to-the-user .dfn}. This user interface
should include features to begin playback, pause playback, seek to an
arbitrary position in the content (if the content supports arbitrary
seeking), change the volume, change the display of closed captions or
embedded sign-language tracks, select different audio tracks or turn on
audio descriptions, and show the media content in manners more suitable
to the user (e.g. fullscreen video or in an independent resizable
window). Other controls may also be made available.

Even when the attribute is absent, however, user agents may provide
controls to affect playback of the media resource (e.g. play, pause,
seeking, track selection, and volume controls), but such features should
not interfere with the page\'s normal rendering. For example, such
features could be exposed in the [media
element](#media-element){#user-interface:media-element-2}\'s context
menu, platform media keys, or a remote control. The user agent may
implement this simply by [exposing a user interface to the
user](#expose-a-user-interface-to-the-user){#user-interface:expose-a-user-interface-to-the-user}
as described above (as if the
[`controls`{#user-interface:attr-media-controls}](#attr-media-controls)
attribute was present).

If the user agent [exposes a user interface to the
user](#expose-a-user-interface-to-the-user){#user-interface:expose-a-user-interface-to-the-user-2}
by displaying controls over the [media
element](#media-element){#user-interface:media-element-3}, then the user
agent should suppress any user interaction events while the user agent
is interacting with this interface. (For example, if the user clicks on
a video\'s playback control,
[`mousedown`{#user-interface:event-mousedown}](https://w3c.github.io/uievents/#event-type-mousedown){x-internal="event-mousedown"}
events and so forth would not simultaneously be fired at elements on the
page.)

Where possible (specifically, for starting, stopping, pausing, and
unpausing playback, for seeking, for changing the rate of playback, for
fast-forwarding or rewinding, for listing, enabling, and disabling text
tracks, and for muting or changing the volume of the audio), user
interface features exposed by the user agent must be implemented in
terms of the DOM API described above, so that, e.g., all the same events
fire.

Features such as fast-forward or rewind must be implemented by only
changing the `playbackRate` attribute (and not the `defaultPlaybackRate`
attribute).

Seeking must be implemented in terms of
[seeking](#dom-media-seek){#user-interface:dom-media-seek} to the
requested position in the [media
element](#media-element){#user-interface:media-element-4}\'s [media
timeline](#media-timeline){#user-interface:media-timeline}. For media
resources where seeking to an arbitrary position would be slow, user
agents are encouraged to use the *approximate-for-speed* flag when
seeking in response to the user manipulating an approximate position
interface such as a seek bar.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/controls](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/controls "The HTMLMediaElement.controls property reflects the controls HTML attribute, which controls whether user interface controls for playing the media item will be displayed.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The [`controls`]{#dom-media-controls .dfn dfn-for="HTMLMediaElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#user-interface:reflect}
the content attribute of the same name.

------------------------------------------------------------------------

`media`{.variable}`.`[`volume`](#dom-media-volume){#dom-media-volume-dev}` [ = ``value`{.variable}` ]`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/volume](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/volume "The HTMLMediaElement.volume property sets the volume at which the media will be played.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS[🔰
3+]{title="Partial implementation."}]{.safari_ios .yes}[Chrome
Android?]{.chrome_android .unknown}[WebView Android37+]{.webview_android
.yes}[Samsung Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns the current playback volume, as a number in the range 0.0 to
1.0, where 0.0 is the quietest and 1.0 the loudest.

Can be set, to change the volume.

Throws an
[\"`IndexSizeError`\"](https://webidl.spec.whatwg.org/#indexsizeerror){#user-interface:indexsizeerror
x-internal="indexsizeerror"}
[`DOMException`{#user-interface:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if the new value is not in the range 0.0 .. 1.0.

`media`{.variable}`.`[`muted`](#dom-media-muted){#dom-media-muted-dev}` [ = ``value`{.variable}` ]`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/muted](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/muted "The HTMLMediaElement.muted property indicates whether the media element is muted.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns true if audio is muted, overriding the
[`volume`{#user-interface:dom-media-volume}](#dom-media-volume)
attribute, and false if the
[`volume`{#user-interface:dom-media-volume-2}](#dom-media-volume)
attribute is being honored.

Can be set, to change whether the audio is muted or not.

A [media element](#media-element){#user-interface:media-element-5} has a
[playback volume]{#concept-media-volume .dfn}, which is a fraction in
the range 0.0 (silent) to 1.0 (loudest). Initially, the volume should be
1.0, but user agents may remember the last set value across sessions, on
a per-site basis or otherwise, so the volume may start at other values.

The [`volume`]{#dom-media-volume .dfn dfn-for="HTMLMediaElement"
dfn-type="attribute"} IDL attribute must return the [playback
volume](#concept-media-volume){#user-interface:concept-media-volume} of
any audio portions of the [media
element](#media-element){#user-interface:media-element-6}. On setting,
if the new value is in the range 0.0 to 1.0 inclusive, the [media
element](#media-element){#user-interface:media-element-7}\'s [playback
volume](#concept-media-volume){#user-interface:concept-media-volume-2}
must be set to the new value. If the new value is outside the range 0.0
to 1.0 inclusive, then, on setting, an
[\"`IndexSizeError`\"](https://webidl.spec.whatwg.org/#indexsizeerror){#user-interface:indexsizeerror-2
x-internal="indexsizeerror"}
[`DOMException`{#user-interface:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
must be thrown instead.

A [media element](#media-element){#user-interface:media-element-8} can
also be [muted]{#concept-media-muted .dfn dfn-for="media element"
export=""}. If anything is muting the element, then it is muted. (For
example, when the [direction of
playback](#direction-of-playback){#user-interface:direction-of-playback}
is backwards, the element is muted.)

The [`muted`]{#dom-media-muted .dfn dfn-for="HTMLMediaElement"
dfn-type="attribute"} IDL attribute must return the value to which it
was last set. When a [media
element](#media-element){#user-interface:media-element-9} is created, if
the element has a
[`muted`{#user-interface:attr-media-muted}](#attr-media-muted) content
attribute specified, then the
[`muted`{#user-interface:dom-media-muted}](#dom-media-muted) IDL
attribute should be set to true; otherwise, the user agents may set the
value to the user\'s preferred value (e.g. remembering the last set
value across sessions, on a per-site basis or otherwise). While the
[`muted`{#user-interface:dom-media-muted-2}](#dom-media-muted) IDL
attribute is set to true, the [media
element](#media-element){#user-interface:media-element-10} must be
[muted](#concept-media-muted){#user-interface:concept-media-muted}.

Whenever either of the values that would be returned by the
[`volume`{#user-interface:dom-media-volume-3}](#dom-media-volume) and
[`muted`{#user-interface:dom-media-muted-3}](#dom-media-muted) IDL
attributes change, the user agent must [queue a media element
task](#queue-a-media-element-task){#user-interface:queue-a-media-element-task}
given the [media
element](#media-element){#user-interface:media-element-11} to [fire an
event](https://dom.spec.whatwg.org/#concept-event-fire){#user-interface:concept-event-fire
x-internal="concept-event-fire"} named
[`volumechange`{#user-interface:event-media-volumechange}](#event-media-volumechange)
at the [media
element](#media-element){#user-interface:media-element-12}. Then, if the
[media element](#media-element){#user-interface:media-element-13} is not
[allowed to play](#allowed-to-play){#user-interface:allowed-to-play},
the user agent must run the [internal pause
steps](#internal-pause-steps){#user-interface:internal-pause-steps} for
the [media element](#media-element){#user-interface:media-element-14}.

A user agent has an associated [volume locked]{#volume-locked .dfn} (a
boolean). Its value is
[implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#user-interface:implementation-defined
x-internal="implementation-defined"} and determines whether the
[playback
volume](#concept-media-volume){#user-interface:concept-media-volume-3}
takes effect.

An element\'s [effective media volume]{#effective-media-volume .dfn
dfn-for="HTMLMediaElement" export=""} is determined as follows:

1.  If the user has indicated that the user agent is to override the
    volume of the element, then return the volume desired by the user.

2.  If the user agent\'s [volume
    locked](#volume-locked){#user-interface:volume-locked} is true, then
    return the system volume.

3.  If the element\'s audio output is
    [muted](#concept-media-muted){#user-interface:concept-media-muted-2},
    then return zero.

4.  Let `volume`{.variable} be the [playback
    volume](#concept-media-volume){#user-interface:concept-media-volume-4}
    of the audio portions of the [media
    element](#media-element){#user-interface:media-element-15}, in range
    0.0 (silent) to 1.0 (loudest).

5.  Return `volume`{.variable}, interpreted relative to the range 0.0 to
    1.0, with 0.0 being silent, and 1.0 being the loudest setting,
    values in between increasing in loudness. The range need not be
    linear. The loudest setting may be lower than the system\'s loudest
    possible setting; for example the user could have set a maximum
    volume.

The [`muted`]{#attr-media-muted .dfn dfn-for="audio,video"
dfn-type="element-attr"} content attribute on [media
elements](#media-element){#user-interface:media-element-16} is a
[boolean
attribute](common-microsyntaxes.html#boolean-attribute){#user-interface:boolean-attribute-2}
that controls the default state of the audio output of the [media
resource](#media-resource){#user-interface:media-resource}, potentially
overriding user preferences.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/defaultMuted](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/defaultMuted "The HTMLMediaElement.defaultMuted property reflects the muted HTML attribute, which indicates whether the media element's audio output should be muted by default. This property has no dynamic effect. To mute and unmute the audio output, use the muted property.")

Support in all current engines.

::: support
[Firefox11+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome15+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The [`defaultMuted`]{#dom-media-defaultmuted .dfn
dfn-for="HTMLMediaElement" dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#user-interface:reflect-2}
the [`muted`{#user-interface:attr-media-muted-2}](#attr-media-muted)
content attribute.

This attribute has no dynamic effect (it only controls the default state
of the element).

::: example
This video (an advertisement) autoplays, but to avoid annoying users, it
does so without sound, and allows the user to turn the sound on. The
user agent can pause the video if it\'s unmuted without a user
interaction.

``` html
<video src="adverts.cgi?kind=video" controls autoplay loop muted></video>
```
:::

##### [4.8.11.14]{.secno} Time ranges[](#time-ranges){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[TimeRanges](https://developer.mozilla.org/en-US/docs/Web/API/TimeRanges "When loading a media resource for use by an <audio> or <video> element, the TimeRanges interface is used for representing the time ranges of the media resource that have been buffered, the time ranges that have been played, and the time ranges that are seekable.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome6+]{.chrome
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

Objects implementing the
[`TimeRanges`{#time-ranges:timeranges}](#timeranges) interface represent
a list of ranges (periods) of time.

``` idl
[Exposed=Window]
interface TimeRanges {
  readonly attribute unsigned long length;
  double start(unsigned long index);
  double end(unsigned long index);
};
```

`media`{.variable}`.`[`length`](#dom-timeranges-length){#dom-timeranges-length-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[TimeRanges/length](https://developer.mozilla.org/en-US/docs/Web/API/TimeRanges/length "The TimeRanges.length read-only property returns the number of ranges in the object.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome6+]{.chrome
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

Returns the number of ranges in the object.

`time`{.variable}` = ``media`{.variable}`.`[`start`](#dom-timeranges-start){#dom-timeranges-start-dev}`(``index`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[TimeRanges/start](https://developer.mozilla.org/en-US/docs/Web/API/TimeRanges/start "The start() method of the TimeRanges interface returns the time offset at which a specified time range begins.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome6+]{.chrome
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

Returns the time for the start of the range with the given index.

Throws an
[\"`IndexSizeError`\"](https://webidl.spec.whatwg.org/#indexsizeerror){#time-ranges:indexsizeerror
x-internal="indexsizeerror"}
[`DOMException`{#time-ranges:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if the index is out of range.

`time`{.variable}` = ``media`{.variable}`.`[`end`](#dom-timeranges-end){#dom-timeranges-end-dev}`(``index`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[TimeRanges/end](https://developer.mozilla.org/en-US/docs/Web/API/TimeRanges/end "The end() method of the TimeRanges interface returns the time offset at which a specified time range ends.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome6+]{.chrome
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

Returns the time for the end of the range with the given index.

Throws an
[\"`IndexSizeError`\"](https://webidl.spec.whatwg.org/#indexsizeerror){#time-ranges:indexsizeerror-2
x-internal="indexsizeerror"}
[`DOMException`{#time-ranges:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if the index is out of range.

The [`length`]{#dom-timeranges-length .dfn dfn-for="TimeRanges"
dfn-type="attribute"} IDL attribute must return the number of ranges
represented by the object.

The [`start(``index`{.variable}`)`]{#dom-timeranges-start .dfn
dfn-for="TimeRanges" dfn-type="method"} method must return the position
of the start of the `index`{.variable}th range represented by the
object, in seconds measured from the start of the timeline that the
object covers.

The [`end(``index`{.variable}`)`]{#dom-timeranges-end .dfn
dfn-for="TimeRanges" dfn-type="method"} method must return the position
of the end of the `index`{.variable}th range represented by the object,
in seconds measured from the start of the timeline that the object
covers.

These methods must throw
[\"`IndexSizeError`\"](https://webidl.spec.whatwg.org/#indexsizeerror){#time-ranges:indexsizeerror-3
x-internal="indexsizeerror"}
[`DOMException`{#time-ranges:domexception-3}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}s
if called with an `index`{.variable} argument greater than or equal to
the number of ranges represented by the object.

When a [`TimeRanges`{#time-ranges:timeranges-2}](#timeranges) object is
said to be a [normalized `TimeRanges`
object]{#normalised-timeranges-object .dfn export=""}, the ranges it
represents must obey the following criteria:

- The start of a range must be greater than the end of all earlier
  ranges.
- The start of a range must be less than or equal to the end of that
  same range.

In other words, the ranges in such an object are ordered, don\'t
overlap, and don\'t touch (adjacent ranges are folded into one bigger
range). A range can be empty (referencing just a single moment in time),
e.g. to indicate that only one frame is currently buffered in the case
that the user agent has discarded the entire [media
resource](#media-resource){#time-ranges:media-resource} except for the
current frame, when a [media
element](#media-element){#time-ranges:media-element} is paused.

Ranges in a [`TimeRanges`{#time-ranges:timeranges-3}](#timeranges)
object must be inclusive.

Thus, the end of a range would be equal to the start of a following
adjacent (touching but not overlapping) range. Similarly, a range
covering a whole timeline anchored at zero would have a start equal to
zero and an end equal to the duration of the timeline.

The timelines used by the objects returned by the
[`buffered`{#time-ranges:dom-media-buffered}](#dom-media-buffered),
[`seekable`{#time-ranges:dom-media-seekable}](#dom-media-seekable), and
[`played`{#time-ranges:dom-media-played}](#dom-media-played) IDL
attributes of [media
elements](#media-element){#time-ranges:media-element-2} must be that
element\'s [media
timeline](#media-timeline){#time-ranges:media-timeline}.

##### [4.8.11.15]{.secno} The [`TrackEvent`{#the-trackevent-interface:trackevent}](#trackevent) interface[](#the-trackevent-interface){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[TrackEvent](https://developer.mozilla.org/en-US/docs/Web/API/TrackEvent "The TrackEvent interface, which is part of the HTML DOM specification, is used for events which represent changes to a set of available tracks on an HTML media element; these events are addtrack and removetrack.")

Support in all current engines.

::: support
[Firefox27+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome23+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer11]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS7+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

``` idl
[Exposed=Window]
interface TrackEvent : Event {
  constructor(DOMString type, optional TrackEventInit eventInitDict = {});

  readonly attribute (VideoTrack or AudioTrack or TextTrack)? track;
};

dictionary TrackEventInit : EventInit {
  (VideoTrack or AudioTrack or TextTrack)? track = null;
};
```

`event`{.variable}`.`[`track`](#dom-trackevent-track){#dom-trackevent-track-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[TrackEvent/track](https://developer.mozilla.org/en-US/docs/Web/API/TrackEvent/track "The read-only track property of the TrackEvent interface specifies the media track object to which the event applies.")

Support in all current engines.

::: support
[Firefox27+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome23+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer11]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS7+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns the track object
([`TextTrack`{#the-trackevent-interface:texttrack-3}](#texttrack),
[`AudioTrack`{#the-trackevent-interface:audiotrack-3}](#audiotrack), or
[`VideoTrack`{#the-trackevent-interface:videotrack-3}](#videotrack)) to
which the event relates.

The [`track`]{#dom-trackevent-track .dfn dfn-for="TrackEvent"
dfn-type="attribute"} attribute must return the value it was initialized
to. It represents the context information for the event.

##### [4.8.11.16]{.secno} Events summary[](#mediaevents){.self-link} {#mediaevents}

*This section is non-normative.*

The following events fire on [media
elements](#media-element){#mediaevents:media-element} as part of the
processing model described above:

Event name

Interface

Fired when\...

Preconditions

[`loadstart`]{#event-media-loadstart .dfn dfn-for="HTMLMediaElement"
dfn-type="event"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/loadstart_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/loadstart_event "The loadstart event is fired when the browser has started to load a resource.")

Support in all current engines.

::: support
[Firefox6+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

[`Event`{#mediaevents:event}](https://dom.spec.whatwg.org/#interface-event){x-internal="event"}

The user agent begins looking for [media
data](#media-data){#mediaevents:media-data}, as part of the [resource
selection
algorithm](#concept-media-load-algorithm){#mediaevents:concept-media-load-algorithm}.

[`networkState`{#mediaevents:dom-media-networkstate}](#dom-media-networkstate)
equals
[`NETWORK_LOADING`{#mediaevents:dom-media-network_loading}](#dom-media-network_loading)

[`progress`]{#event-media-progress .dfn dfn-for="HTMLMediaElement"
dfn-type="event"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/progress_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/progress_event "The progress event is fired periodically as the browser loads a resource.")

Support in all current engines.

::: support
[Firefox6+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

[`Event`{#mediaevents:event-2}](https://dom.spec.whatwg.org/#interface-event){x-internal="event"}

The user agent is fetching [media
data](#media-data){#mediaevents:media-data-2}.

[`networkState`{#mediaevents:dom-media-networkstate-2}](#dom-media-networkstate)
equals
[`NETWORK_LOADING`{#mediaevents:dom-media-network_loading-2}](#dom-media-network_loading)

[`suspend`]{#event-media-suspend .dfn dfn-for="HTMLMediaElement"
dfn-type="event"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/suspend_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/suspend_event "The suspend event is fired when media data loading has been suspended.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

[`Event`{#mediaevents:event-3}](https://dom.spec.whatwg.org/#interface-event){x-internal="event"}

The user agent is intentionally not currently fetching [media
data](#media-data){#mediaevents:media-data-3}.

[`networkState`{#mediaevents:dom-media-networkstate-3}](#dom-media-networkstate)
equals
[`NETWORK_IDLE`{#mediaevents:dom-media-network_idle}](#dom-media-network_idle)

[`abort`]{#event-media-abort .dfn dfn-for="HTMLMediaElement"
dfn-type="event"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/abort_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/abort_event "The abort event is fired when the resource was not fully loaded, but not as the result of an error.")

Support in all current engines.

::: support
[Firefox9+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

[`Event`{#mediaevents:event-4}](https://dom.spec.whatwg.org/#interface-event){x-internal="event"}

The user agent stops fetching the [media
data](#media-data){#mediaevents:media-data-4} before it is completely
downloaded, but not due to an error.

[`error`{#mediaevents:dom-media-error}](#dom-media-error) is an object
with the code
[`MEDIA_ERR_ABORTED`{#mediaevents:dom-mediaerror-media_err_aborted}](#dom-mediaerror-media_err_aborted).
[`networkState`{#mediaevents:dom-media-networkstate-4}](#dom-media-networkstate)
equals either
[`NETWORK_EMPTY`{#mediaevents:dom-media-network_empty}](#dom-media-network_empty)
or
[`NETWORK_IDLE`{#mediaevents:dom-media-network_idle-2}](#dom-media-network_idle),
depending on when the download was aborted.

[`error`]{#event-media-error .dfn dfn-for="HTMLMediaElement"
dfn-type="event"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/error_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/error_event "The error event is fired when the resource could not be loaded due to an error (for example, a network connectivity problem).")

Support in all current engines.

::: support
[Firefox6+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12+]{.opera_android .yes}
:::
::::
:::::

[`Event`{#mediaevents:event-5}](https://dom.spec.whatwg.org/#interface-event){x-internal="event"}

An error occurs while fetching the [media
data](#media-data){#mediaevents:media-data-5} or the type of the
resource is not a supported media format.

[`error`{#mediaevents:dom-media-error-2}](#dom-media-error) is an object
with the code
[`MEDIA_ERR_NETWORK`{#mediaevents:dom-mediaerror-media_err_network}](#dom-mediaerror-media_err_network)
or higher.
[`networkState`{#mediaevents:dom-media-networkstate-5}](#dom-media-networkstate)
equals either
[`NETWORK_EMPTY`{#mediaevents:dom-media-network_empty-2}](#dom-media-network_empty)
or
[`NETWORK_IDLE`{#mediaevents:dom-media-network_idle-3}](#dom-media-network_idle),
depending on when the download was aborted.

[`emptied`]{#event-media-emptied .dfn dfn-for="HTMLMediaElement"
dfn-type="event"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/emptied_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/emptied_event "The emptied event is fired when the media has become empty; for example, this event is sent if the media has already been loaded (or partially loaded), and the load() method is called to reload it.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

[`Event`{#mediaevents:event-6}](https://dom.spec.whatwg.org/#interface-event){x-internal="event"}

A [media element](#media-element){#mediaevents:media-element-2} whose
[`networkState`{#mediaevents:dom-media-networkstate-6}](#dom-media-networkstate)
was previously not in the
[`NETWORK_EMPTY`{#mediaevents:dom-media-network_empty-3}](#dom-media-network_empty)
state has just switched to that state (either because of a fatal error
during load that\'s about to be reported, or because the
[`load()`{#mediaevents:dom-media-load}](#dom-media-load) method was
invoked while the [resource selection
algorithm](#concept-media-load-algorithm){#mediaevents:concept-media-load-algorithm-2}
was already running).

[`networkState`{#mediaevents:dom-media-networkstate-7}](#dom-media-networkstate)
is
[`NETWORK_EMPTY`{#mediaevents:dom-media-network_empty-4}](#dom-media-network_empty);
all the IDL attributes are in their initial states.

[`stalled`]{#event-media-stalled .dfn dfn-for="HTMLMediaElement"
dfn-type="event"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/stalled_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/stalled_event "The stalled event is fired when the user agent is trying to fetch media data, but data is unexpectedly not forthcoming.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

[`Event`{#mediaevents:event-7}](https://dom.spec.whatwg.org/#interface-event){x-internal="event"}

The user agent is trying to fetch [media
data](#media-data){#mediaevents:media-data-6}, but data is unexpectedly
not forthcoming.

[`networkState`{#mediaevents:dom-media-networkstate-8}](#dom-media-networkstate)
is
[`NETWORK_LOADING`{#mediaevents:dom-media-network_loading-3}](#dom-media-network_loading).

[`loadedmetadata`]{#event-media-loadedmetadata .dfn
dfn-for="HTMLMediaElement" dfn-type="event"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/loadedmetadata_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/loadedmetadata_event "The loadedmetadata event is fired when the metadata has been loaded.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

[`Event`{#mediaevents:event-8}](https://dom.spec.whatwg.org/#interface-event){x-internal="event"}

The user agent has just determined the duration and dimensions of the
[media resource](#media-resource){#mediaevents:media-resource} and [the
text tracks are
ready](#the-text-tracks-are-ready){#mediaevents:the-text-tracks-are-ready}.

[`readyState`{#mediaevents:dom-media-readystate}](#dom-media-readystate)
is newly equal to
[`HAVE_METADATA`{#mediaevents:dom-media-have_metadata}](#dom-media-have_metadata)
or greater for the first time.

[`loadeddata`]{#event-media-loadeddata .dfn dfn-for="HTMLMediaElement"
dfn-type="event"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/loadeddata_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/loadeddata_event "The loadeddata event is fired when the frame at the current playback position of the media has finished loading; often the first frame.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

[`Event`{#mediaevents:event-9}](https://dom.spec.whatwg.org/#interface-event){x-internal="event"}

The user agent can render the [media
data](#media-data){#mediaevents:media-data-7} at the [current playback
position](#current-playback-position){#mediaevents:current-playback-position}
for the first time.

[`readyState`{#mediaevents:dom-media-readystate-2}](#dom-media-readystate)
newly increased to
[`HAVE_CURRENT_DATA`{#mediaevents:dom-media-have_current_data}](#dom-media-have_current_data)
or greater for the first time.

[`canplay`]{#event-media-canplay .dfn dfn-for="HTMLMediaElement"
dfn-type="event"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/canplay_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/canplay_event "The canplay event is fired when the user agent can play the media, but estimates that not enough data has been loaded to play the media up to its end without having to stop for further buffering of content.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

[`Event`{#mediaevents:event-10}](https://dom.spec.whatwg.org/#interface-event){x-internal="event"}

The user agent can resume playback of the [media
data](#media-data){#mediaevents:media-data-8}, but estimates that if
playback were to be started now, the [media
resource](#media-resource){#mediaevents:media-resource-2} could not be
rendered at the current playback rate up to its end without having to
stop for further buffering of content.

[`readyState`{#mediaevents:dom-media-readystate-3}](#dom-media-readystate)
newly increased to
[`HAVE_FUTURE_DATA`{#mediaevents:dom-media-have_future_data}](#dom-media-have_future_data)
or greater.

[`canplaythrough`]{#event-media-canplaythrough .dfn
dfn-for="HTMLMediaElement" dfn-type="event"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/canplaythrough_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/canplaythrough_event "The canplaythrough event is fired when the user agent can play the media, and estimates that enough data has been loaded to play the media up to its end without having to stop for further buffering of content.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

[`Event`{#mediaevents:event-11}](https://dom.spec.whatwg.org/#interface-event){x-internal="event"}

The user agent estimates that if playback were to be started now, the
[media resource](#media-resource){#mediaevents:media-resource-3} could
be rendered at the current playback rate all the way to its end without
having to stop for further buffering.

[`readyState`{#mediaevents:dom-media-readystate-4}](#dom-media-readystate)
is newly equal to
[`HAVE_ENOUGH_DATA`{#mediaevents:dom-media-have_enough_data}](#dom-media-have_enough_data).

[`playing`]{#event-media-playing .dfn dfn-for="HTMLMediaElement"
dfn-type="event"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/playing_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/playing_event "The playing event is fired after playback is first started, and whenever it is restarted. For example it is fired when playback resumes after having been paused or delayed due to lack of data.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

[`Event`{#mediaevents:event-12}](https://dom.spec.whatwg.org/#interface-event){x-internal="event"}

Playback is ready to start after having been paused or delayed due to
lack of [media data](#media-data){#mediaevents:media-data-9}.

[`readyState`{#mediaevents:dom-media-readystate-5}](#dom-media-readystate)
is newly greater than or equal to
[`HAVE_FUTURE_DATA`{#mediaevents:dom-media-have_future_data-2}](#dom-media-have_future_data)
and [`paused`{#mediaevents:dom-media-paused}](#dom-media-paused) is
false, or [`paused`{#mediaevents:dom-media-paused-2}](#dom-media-paused)
is newly false and
[`readyState`{#mediaevents:dom-media-readystate-6}](#dom-media-readystate)
is greater than or equal to
[`HAVE_FUTURE_DATA`{#mediaevents:dom-media-have_future_data-3}](#dom-media-have_future_data).
Even if this event fires, the element might still not be [potentially
playing](#potentially-playing){#mediaevents:potentially-playing}, e.g.
if the element is [paused for user
interaction](#paused-for-user-interaction){#mediaevents:paused-for-user-interaction}
or [paused for in-band
content](#paused-for-in-band-content){#mediaevents:paused-for-in-band-content}.

[`waiting`]{#event-media-waiting .dfn dfn-for="HTMLMediaElement"
dfn-type="event"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/waiting_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/waiting_event "The waiting event is fired when playback has stopped because of a temporary lack of data.")

Support in all current engines.

::: support
[Firefox6+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

[`Event`{#mediaevents:event-13}](https://dom.spec.whatwg.org/#interface-event){x-internal="event"}

Playback has stopped because the next frame is not available, but the
user agent expects that frame to become available in due course.

[`readyState`{#mediaevents:dom-media-readystate-7}](#dom-media-readystate)
is less than or equal to
[`HAVE_CURRENT_DATA`{#mediaevents:dom-media-have_current_data-2}](#dom-media-have_current_data),
and [`paused`{#mediaevents:dom-media-paused-3}](#dom-media-paused) is
false. Either
[`seeking`{#mediaevents:dom-media-seeking}](#dom-media-seeking) is true,
or the [current playback
position](#current-playback-position){#mediaevents:current-playback-position-2}
is not contained in any of the ranges in
[`buffered`{#mediaevents:dom-media-buffered}](#dom-media-buffered). It
is possible for playback to stop for other reasons without
[`paused`{#mediaevents:dom-media-paused-4}](#dom-media-paused) being
false, but those reasons do not fire this event (and when those
situations resolve, a separate
[`playing`{#mediaevents:event-media-playing}](#event-media-playing)
event is not fired either): e.g., [playback has
ended](#ended-playback){#mediaevents:ended-playback}, or playback
[stopped due to
errors](#stopped-due-to-errors){#mediaevents:stopped-due-to-errors}, or
the element has [paused for user
interaction](#paused-for-user-interaction){#mediaevents:paused-for-user-interaction-2}
or [paused for in-band
content](#paused-for-in-band-content){#mediaevents:paused-for-in-band-content-2}.

[`seeking`]{#event-media-seeking .dfn dfn-for="HTMLMediaElement"
dfn-type="event"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/seeking_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/seeking_event "The seeking event is fired when a seek operation starts, meaning the Boolean seeking attribute has changed to true and the media is seeking a new position.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

[`Event`{#mediaevents:event-14}](https://dom.spec.whatwg.org/#interface-event){x-internal="event"}

The [`seeking`{#mediaevents:dom-media-seeking-2}](#dom-media-seeking)
IDL attribute changed to true, and the user agent has started seeking to
a new position.

[`seeked`]{#event-media-seeked .dfn dfn-for="HTMLMediaElement"
dfn-type="event"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/seeked_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/seeked_event "The seeked event is fired when a seek operation completed, the current playback position has changed, and the Boolean seeking attribute is changed to false.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

[`Event`{#mediaevents:event-15}](https://dom.spec.whatwg.org/#interface-event){x-internal="event"}

The [`seeking`{#mediaevents:dom-media-seeking-3}](#dom-media-seeking)
IDL attribute changed to false after the [current playback
position](#current-playback-position){#mediaevents:current-playback-position-3}
was changed.

[`ended`]{#event-media-ended .dfn dfn-for="HTMLMediaElement"
dfn-type="event"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/ended_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/ended_event "The ended event is fired when playback or streaming has stopped because the end of the media was reached or because no further data is available.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

[`Event`{#mediaevents:event-16}](https://dom.spec.whatwg.org/#interface-event){x-internal="event"}

Playback has stopped because the end of the [media
resource](#media-resource){#mediaevents:media-resource-4} was reached.

[`currentTime`{#mediaevents:dom-media-currenttime}](#dom-media-currenttime)
equals the end of the [media
resource](#media-resource){#mediaevents:media-resource-5};
[`ended`{#mediaevents:dom-media-ended}](#dom-media-ended) is true.

[`durationchange`]{#event-media-durationchange .dfn
dfn-for="HTMLMediaElement" dfn-type="event"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/durationchange_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/durationchange_event "The durationchange event is fired when the duration attribute has been updated.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

[`Event`{#mediaevents:event-17}](https://dom.spec.whatwg.org/#interface-event){x-internal="event"}

The [`duration`{#mediaevents:dom-media-duration}](#dom-media-duration)
attribute has just been updated.

[`timeupdate`]{#event-media-timeupdate .dfn dfn-for="HTMLMediaElement"
dfn-type="event"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/timeupdate_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/timeupdate_event "The timeupdate event is fired when the time indicated by the currentTime attribute has been updated.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

[`Event`{#mediaevents:event-18}](https://dom.spec.whatwg.org/#interface-event){x-internal="event"}

The [current playback
position](#current-playback-position){#mediaevents:current-playback-position-4}
changed as part of normal playback or in an especially interesting way,
for example discontinuously.

[`play`]{#event-media-play .dfn dfn-for="HTMLMediaElement"
dfn-type="event"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/play_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/play_event "The play event is fired when the paused property is changed from true to false, as a result of the play method, or the autoplay attribute.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

[`Event`{#mediaevents:event-19}](https://dom.spec.whatwg.org/#interface-event){x-internal="event"}

The element is no longer paused. Fired after the
[`play()`{#mediaevents:dom-media-play}](#dom-media-play) method has
returned, or when the
[`autoplay`{#mediaevents:attr-media-autoplay}](#attr-media-autoplay)
attribute has caused playback to begin.

[`paused`{#mediaevents:dom-media-paused-5}](#dom-media-paused) is newly
false.

[`pause`]{#event-media-pause .dfn dfn-for="HTMLMediaElement"
dfn-type="event"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/pause_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/pause_event "The pause event is sent when a request to pause an activity is handled and the activity has entered its paused state, most commonly after the media has been paused through a call to the element's pause() method.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

[`Event`{#mediaevents:event-20}](https://dom.spec.whatwg.org/#interface-event){x-internal="event"}

The element has been paused. Fired after the
[`pause()`{#mediaevents:dom-media-pause}](#dom-media-pause) method has
returned.

[`paused`{#mediaevents:dom-media-paused-6}](#dom-media-paused) is newly
true.

[`ratechange`]{#event-media-ratechange .dfn dfn-for="HTMLMediaElement"
dfn-type="event"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/ratechange_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/ratechange_event "The ratechange event is fired when the playback rate has changed.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

[`Event`{#mediaevents:event-21}](https://dom.spec.whatwg.org/#interface-event){x-internal="event"}

Either the
[`defaultPlaybackRate`{#mediaevents:dom-media-defaultplaybackrate}](#dom-media-defaultplaybackrate)
or the
[`playbackRate`{#mediaevents:dom-media-playbackrate}](#dom-media-playbackrate)
attribute has just been updated.

[`resize`]{#event-media-resize .dfn dfn-for="HTMLMediaElement"
dfn-type="event"}

[`Event`{#mediaevents:event-22}](https://dom.spec.whatwg.org/#interface-event){x-internal="event"}

One or both of the
[`videoWidth`{#mediaevents:dom-video-videowidth}](#dom-video-videowidth)
and
[`videoHeight`{#mediaevents:dom-video-videoheight}](#dom-video-videoheight)
attributes have just been updated.

[Media element](#media-element){#mediaevents:media-element-3} is a
[`video`{#mediaevents:the-video-element}](#the-video-element) element;
[`readyState`{#mediaevents:dom-media-readystate-8}](#dom-media-readystate)
is not
[`HAVE_NOTHING`{#mediaevents:dom-media-have_nothing}](#dom-media-have_nothing)

[`volumechange`]{#event-media-volumechange .dfn
dfn-for="HTMLMediaElement" dfn-type="event"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMediaElement/volumechange_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/volumechange_event "The volumechange event is fired when either the volume attribute or the muted attribute has changed.")

Support in all current engines.

::: support
[Firefox6+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

[`Event`{#mediaevents:event-23}](https://dom.spec.whatwg.org/#interface-event){x-internal="event"}

Either the [`volume`{#mediaevents:dom-media-volume}](#dom-media-volume)
attribute or the
[`muted`{#mediaevents:dom-media-muted}](#dom-media-muted) attribute has
changed. Fired after the relevant attribute\'s setter has returned.

The following event fires on
[`source`{#mediaevents:the-source-element}](embedded-content.html#the-source-element)
elements:

Event name

Interface

Fired when\...

[`error`]{#event-source-error .dfn dfn-for="HTMLSourceElement"
dfn-type="event"}

[`Event`{#mediaevents:event-24}](https://dom.spec.whatwg.org/#interface-event){x-internal="event"}

An error occurs while fetching the [media
data](#media-data){#mediaevents:media-data-10} or the type of the
resource is not a supported media format.

The following events fire on
[`AudioTrackList`{#mediaevents:audiotracklist}](#audiotracklist),
[`VideoTrackList`{#mediaevents:videotracklist}](#videotracklist), and
[`TextTrackList`{#mediaevents:texttracklist}](#texttracklist) objects:

Event name

Interface

Fired when\...

[`change`]{#event-media-change .dfn
dfn-for="AudioTrackList,VideoTrackList,TextTrackList" dfn-type="event"}

::::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[AudioTrackList/change_event](https://developer.mozilla.org/en-US/docs/Web/API/AudioTrackList/change_event "The change event is fired when an audio track is enabled or disabled, for example by changing the track's enabled property.")

Support in all current engines.

::: support
[Firefox[🔰
33+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari7+]{.safari .yes}[Chrome[🔰
37+]{title="Requires setting a user preference or runtime flag."}]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge[🔰
79+]{title="Requires setting a user preference or runtime flag."}]{.edge_blink
.yes}

------------------------------------------------------------------------

[Edge (Legacy)No]{.edge .no}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::

:::: feature
[TextTrackList/change_event](https://developer.mozilla.org/en-US/docs/Web/API/TextTrackList/change_event "The change event is fired when a text track is made active or inactive, or a TextTrackList is otherwise changed.")

Support in all current engines.

::: support
[Firefox31+]{.firefox .yes}[Safari7+]{.safari .yes}[Chrome33+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)18]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android4.4+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::

:::: feature
[VideoTrackList/change_event](https://developer.mozilla.org/en-US/docs/Web/API/VideoTrackList/change_event "The change event is fired when a video track is made active or inactive, for example by changing the track's selected property.")

Support in all current engines.

::: support
[Firefox[🔰
33+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari7+]{.safari .yes}[Chrome[🔰
37+]{title="Requires setting a user preference or runtime flag."}]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge[🔰
79+]{title="Requires setting a user preference or runtime flag."}]{.edge_blink
.yes}

------------------------------------------------------------------------

[Edge (Legacy)No]{.edge .no}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::::::

[`Event`{#mediaevents:event-25}](https://dom.spec.whatwg.org/#interface-event){x-internal="event"}

One or more tracks in the track list have been enabled or disabled.

[`addtrack`]{#event-media-addtrack .dfn
dfn-for="AudioTrackList,VideoTrackList,TextTrackList" dfn-type="event"}

::::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[AudioTrackList/addtrack_event](https://developer.mozilla.org/en-US/docs/Web/API/AudioTrackList/addtrack_event "The addtrack event is fired when a track is added to an AudioTrackList.")

Support in all current engines.

::: support
[Firefox[🔰
33+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari7+]{.safari .yes}[Chrome[🔰
37+]{title="Requires setting a user preference or runtime flag."}]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge[🔰
79+]{title="Requires setting a user preference or runtime flag."}]{.edge_blink
.yes}

------------------------------------------------------------------------

[Edge (Legacy)No]{.edge .no}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::

:::: feature
[TextTrackList/addtrack_event](https://developer.mozilla.org/en-US/docs/Web/API/TextTrackList/addtrack_event "The addtrack event is fired when a track is added to a TextTrackList.")

Support in all current engines.

::: support
[Firefox31+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome23+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer11]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS7+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::

:::: feature
[VideoTrackList/addtrack_event](https://developer.mozilla.org/en-US/docs/Web/API/VideoTrackList/addtrack_event "The addtrack event is fired when a video track is added to a VideoTrackList.")

Support in all current engines.

::: support
[Firefox[🔰
33+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari7+]{.safari .yes}[Chrome[🔰
37+]{title="Requires setting a user preference or runtime flag."}]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge[🔰
79+]{title="Requires setting a user preference or runtime flag."}]{.edge_blink
.yes}

------------------------------------------------------------------------

[Edge (Legacy)No]{.edge .no}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::::::

[`TrackEvent`{#mediaevents:trackevent}](#trackevent)

A track has been added to the track list.

[`removetrack`]{#event-media-removetrack .dfn
dfn-for="AudioTrackList,VideoTrackList,TextTrackList" dfn-type="event"}

::::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[AudioTrackList/removetrack_event](https://developer.mozilla.org/en-US/docs/Web/API/AudioTrackList/removetrack_event "The removetrack event is fired when a track is removed from an AudioTrackList.")

Support in all current engines.

::: support
[Firefox[🔰
33+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari7+]{.safari .yes}[Chrome[🔰
37+]{title="Requires setting a user preference or runtime flag."}]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge[🔰
79+]{title="Requires setting a user preference or runtime flag."}]{.edge_blink
.yes}

------------------------------------------------------------------------

[Edge (Legacy)No]{.edge .no}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::

:::: feature
[TextTrackList/removetrack_event](https://developer.mozilla.org/en-US/docs/Web/API/TextTrackList/removetrack_event "The removetrack event is fired when a track is removed from a TextTrackList.")

Support in all current engines.

::: support
[Firefox31+]{.firefox .yes}[Safari7+]{.safari .yes}[Chrome33+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera20+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)18]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android4.4+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android20+]{.opera_android .yes}
:::
::::

:::: feature
[VideoTrackList/removetrack_event](https://developer.mozilla.org/en-US/docs/Web/API/VideoTrackList/removetrack_event "The removetrack event is fired when a video track is removed from a VideoTrackList.")

Support in all current engines.

::: support
[Firefox[🔰
33+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari7+]{.safari .yes}[Chrome[🔰
37+]{title="Requires setting a user preference or runtime flag."}]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge[🔰
79+]{title="Requires setting a user preference or runtime flag."}]{.edge_blink
.yes}

------------------------------------------------------------------------

[Edge (Legacy)No]{.edge .no}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::::::

[`TrackEvent`{#mediaevents:trackevent-2}](#trackevent)

A track has been removed from the track list.

The following event fires on
[`TextTrack`{#mediaevents:texttrack}](#texttrack) objects and
[`track`{#mediaevents:the-track-element}](#the-track-element) elements:

Event name

Interface

Fired when\...

[`cuechange`]{#event-media-cuechange .dfn
dfn-for="TextTrack,HTMLTrackElement" dfn-type="event"}

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLTrackElement/cuechange_event](https://developer.mozilla.org/en-US/docs/Web/API/HTMLTrackElement/cuechange_event "The cuechange event fires when a TextTrack has changed the currently displaying cues. The event is fired on both the TextTrack and the HTMLTrackElement in which it's being presented, if any.")

Support in all current engines.

::: support
[Firefox68+]{.firefox .yes}[Safari10+]{.safari .yes}[Chrome32+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera19+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)14+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android4.4.3+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android19+]{.opera_android .yes}
:::
::::

:::: feature
[TextTrack/cuechange_event](https://developer.mozilla.org/en-US/docs/Web/API/TextTrack/cuechange_event "The cuechange event fires when a TextTrack has changed the currently displaying cues. The event is fired on both the TextTrack and the HTMLTrackElement in which it's being presented, if any.")

Support in all current engines.

::: support
[Firefox31+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome23+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS7+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::::

[`Event`{#mediaevents:event-26}](https://dom.spec.whatwg.org/#interface-event){x-internal="event"}

One or more cues in the track have become active or stopped being
active.

The following events fire on
[`track`{#mediaevents:the-track-element-2}](#the-track-element)
elements:

Event name

Interface

Fired when\...

[`error`]{#event-track-error .dfn dfn-for="HTMLTrackElement"
dfn-type="event"}

[`Event`{#mediaevents:event-27}](https://dom.spec.whatwg.org/#interface-event){x-internal="event"}

An error occurs while fetching the track data or the type of the
resource is not supported text track format.

[`load`]{#event-track-load .dfn dfn-for="HTMLTrackElement"
dfn-type="event"}

[`Event`{#mediaevents:event-28}](https://dom.spec.whatwg.org/#interface-event){x-internal="event"}

A track data has been fetched and successfully processed.

The following events fire on
[`TextTrackCue`{#mediaevents:texttrackcue}](#texttrackcue) objects:

Event name

Interface

Fired when\...

[`enter`]{#event-media-enter .dfn dfn-for="TextTrackCue"
dfn-type="event"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[TextTrackCue/enter_event](https://developer.mozilla.org/en-US/docs/Web/API/TextTrackCue/enter_event "The enter event fires when a cue becomes active. In the case of subtitles or a caption this is when it displays over the media.")

Support in all current engines.

::: support
[Firefox31+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome23+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS7+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

[`Event`{#mediaevents:event-29}](https://dom.spec.whatwg.org/#interface-event){x-internal="event"}

The cue has become active.

[`exit`]{#event-media-exit .dfn dfn-for="TextTrackCue" dfn-type="event"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[TextTrackCue/exit_event](https://developer.mozilla.org/en-US/docs/Web/API/TextTrackCue/exit_event "The exit event fires when a cue stops being active.")

Support in all current engines.

::: support
[Firefox31+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome23+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS7+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

[`Event`{#mediaevents:event-30}](https://dom.spec.whatwg.org/#interface-event){x-internal="event"}

The cue has stopped being active.

##### [4.8.11.17]{.secno} Security and privacy considerations[](#security-and-privacy-considerations){.self-link}

The main security and privacy implications of the
[`video`{#security-and-privacy-considerations:the-video-element}](#the-video-element)
and
[`audio`{#security-and-privacy-considerations:the-audio-element}](#the-audio-element)
elements come from the ability to embed media cross-origin. There are
two directions that threats can flow: from hostile content to a victim
page, and from a hostile page to victim content.

------------------------------------------------------------------------

If a victim page embeds hostile content, the threat is that the content
might contain scripted code that attempts to interact with the
[`Document`{#security-and-privacy-considerations:document}](dom.html#document)
that embeds the content. To avoid this, user agents must ensure that
there is no access from the content to the embedding page. In the case
of media content that uses DOM concepts, the embedded content must be
treated as if it was in its own unrelated [top-level
traversable](document-sequences.html#top-level-traversable){#security-and-privacy-considerations:top-level-traversable}.

For instance, if an SVG animation was embedded in a
[`video`{#security-and-privacy-considerations:the-video-element-2}](#the-video-element)
element, the user agent would not give it access to the DOM of the outer
page. From the perspective of scripts in the SVG resource, the SVG file
would appear to be in a lone top-level traversable with no parent.

------------------------------------------------------------------------

If a hostile page embeds victim content, the threat is that the
embedding page could obtain information from the content that it would
not otherwise have access to. The API does expose some information: the
existence of the media, its type, its duration, its size, and the
performance characteristics of its host. Such information is already
potentially problematic, but in practice the same information can more
or less be obtained using the
[`img`{#security-and-privacy-considerations:the-img-element}](embedded-content.html#the-img-element)
element, and so it has been deemed acceptable.

However, significantly more sensitive information could be obtained if
the user agent further exposes metadata within the content, such as
subtitles. That information is therefore only exposed if the video
resource uses CORS. The
[`crossorigin`{#security-and-privacy-considerations:attr-media-crossorigin}](#attr-media-crossorigin)
attribute allows authors to enable CORS.
[\[FETCH\]](references.html#refsFETCH)

Without this restriction, an attacker could trick a user running within
a corporate network into visiting a site that attempts to load a video
from a previously leaked location on the corporation\'s intranet. If
such a video included confidential plans for a new product, then being
able to read the subtitles would present a serious confidentiality
breach.

##### [4.8.11.18]{.secno} Best practices for authors using media elements[](#best-practices-for-authors-using-media-elements){.self-link}

*This section is non-normative.*

Playing audio and video resources on small devices such as set-top boxes
or mobile phones is often constrained by limited hardware resources in
the device. For example, a device might only support three simultaneous
videos. For this reason, it is a good practice to release resources held
by [media
elements](#media-element){#best-practices-for-authors-using-media-elements:media-element}
when they are done playing, either by being very careful about removing
all references to the element and allowing it to be garbage collected,
or, even better, by setting the element\'s
[`src`{#best-practices-for-authors-using-media-elements:attr-media-src}](#attr-media-src)
attribute to an empty string. In cases where
[`srcObject`{#best-practices-for-authors-using-media-elements:dom-media-srcobject}](#dom-media-srcobject)
was set, instead set the
[`srcObject`{#best-practices-for-authors-using-media-elements:dom-media-srcobject-2}](#dom-media-srcobject)
to null.

Similarly, when the playback rate is not exactly 1.0, hardware,
software, or format limitations can cause video frames to be dropped and
audio to be choppy or muted.

##### [4.8.11.19]{.secno} Best practices for implementers of media elements[](#best-practices-for-implementers-of-media-elements){.self-link}

*This section is non-normative.*

How accurately various aspects of the [media
element](#media-element){#best-practices-for-implementers-of-media-elements:media-element}
API are implemented is considered a quality-of-implementation issue.

For example, when implementing the
[`buffered`{#best-practices-for-implementers-of-media-elements:dom-media-buffered}](#dom-media-buffered)
attribute, how precise an implementation reports the ranges that have
been buffered depends on how carefully the user agent inspects the data.
Since the API reports ranges as times, but the data is obtained in byte
streams, a user agent receiving a variable-bitrate stream might only be
able to determine precise times by actually decoding all of the data.
User agents aren\'t required to do this, however; they can instead
return estimates (e.g. based on the average bitrate seen so far) which
get revised as more information becomes available.

As a general rule, user agents are urged to be conservative rather than
optimistic. For example, it would be bad to report that everything had
been buffered when it had not.

Another quality-of-implementation issue would be playing a video
backwards when the codec is designed only for forward playback (e.g.
there aren\'t many key frames, and they are far apart, and the
intervening frames only have deltas from the previous frame). User
agents could do a poor job, e.g. only showing key frames; however,
better implementations would do more work and thus do a better job, e.g.
actually decoding parts of the video forwards, storing the complete
frames, and then playing the frames backwards.

Similarly, while implementations are allowed to drop buffered data at
any time (there is no requirement that a user agent keep all the media
data obtained for the lifetime of the media element), it is again a
quality of implementation issue: user agents with sufficient resources
to keep all the data around are encouraged to do so, as this allows for
a better user experience. For example, if the user is watching a live
stream, a user agent could allow the user only to view the live video;
however, a better user agent would buffer everything and allow the user
to seek through the earlier material, pause it, play it forwards and
backwards, etc.

------------------------------------------------------------------------

When a [media
element](#media-element){#best-practices-for-implementers-of-media-elements:media-element-2}
that is paused is [removed from a
document](infrastructure.html#remove-an-element-from-a-document){#best-practices-for-implementers-of-media-elements:remove-an-element-from-a-document}
and not reinserted before the next time the [event
loop](webappapis.html#event-loop){#best-practices-for-implementers-of-media-elements:event-loop}
reaches [step 1](webappapis.html#step1), implementations that are
resource constrained are encouraged to take that opportunity to release
all hardware resources (like video planes, networking resources, and
data buffers) used by the [media
element](#media-element){#best-practices-for-implementers-of-media-elements:media-element-3}.
(User agents still have to keep track of the playback position and so
forth, though, in case playback is later restarted.)

[← 4.8.5 The iframe element](iframe-embed-object.html) --- [Table of
Contents](./) --- [4.8.12 The map element →](image-maps.html)
